package transport

import (
	"context"
	"fmt"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	ktransport "github.com/go-kratos/kratos/v2/transport"
)

var _ ktransport.Server = (*Server)(nil)

// Server implements [ktransport.Server].
type Server struct {
	serverCtx       context.Context
	serverCancel    context.CancelFunc
	logger          log.Logger
	logh            *log.Helper
	middlewareChain middleware.Middleware
	timeout         time.Duration

	rtmClient *agrtm.IRtmClient

	messageChannelSubscription map[string][]agrtm.RtmEventHandler
}

func NewServer(rtmSupplier RtmClientSupplier, opts ...ServerOption) (*Server, error) {
	s := &Server{
		serverCtx:       nil,
		serverCancel:    nil,
		logger:          log.GetLogger(),
		logh:            log.NewHelper(log.GetLogger()),
		middlewareChain: nil,
		timeout:         0,

		rtmClient: nil,

		messageChannelSubscription: map[string][]agrtm.RtmEventHandler{},
	}

	for _, opt := range opts {
		opt(s)
	}

	rtmClient, err := rtmSupplier()
	if err != nil {
		return nil, err
	}
	s.rtmClient = rtmClient

	return s, nil
}

// Start implements [ktransport.Server].
func (s *Server) Start(ctx context.Context) error {
	s.serverCtx, s.serverCancel = context.WithCancel(ctx)

	// 订阅 Message Channel
	for cn := range s.messageChannelSubscription {
		_, err := rtmAsync.Call[int](func() (ret int, reqId uint64) {
			return s.rtmClient.Subscribe(cn, &agrtm.SubscribeOptions{
				WithMessage:  true,
				WithMetadata: true,
				WithPresence: true,
				WithLock:     true,
				BeQuiet:      false,
			})
		}).Await(s.serverCtx)
		if err != nil {
			return err
		}
	}

	// TODO 订阅 Stream Channel

	return nil
}

// Stop implements [ktransport.Server].
func (s *Server) Stop(ctx context.Context) error {
	s.serverCancel()

	return nil
}

// Subscribe a Message Channel
func (s *Server) Subscribe(channelName string, svc IRtmMessageChannelService) {
	// Find or New Slice
	subs := s.messageChannelSubscription[channelName]
	if subs == nil {
		s.messageChannelSubscription[channelName] = []agrtm.RtmEventHandler{}
	}

	// Subscribe and Append to the Slices
	sub := newMessageChannelSubscription(s, svc, channelName)
	subs = append(subs, sub)
	s.messageChannelSubscription[channelName] = subs
}

func (s *Server) newServerContext(channelName string, topic string) context.Context {
	tr := &Transport{
		endpoint:    fmt.Sprintf("%s://%s/%s/%s", KindRtm, SubKindMessageChannel, channelName, topic),
		reqHeader:   headerCarrier{}, //TODO
		replyHeader: headerCarrier{},

		topic:   topic,
		subKind: SubKindMessageChannel,
		server:  s,
	}

	var ctx context.Context
	if s.timeout > 0 {
		ctx, _ = context.WithTimeoutCause(s.serverCtx, s.timeout, ErrServerContextTimeout)
	}
	ctx = ktransport.NewServerContext(ctx, tr)

	return ctx
}
