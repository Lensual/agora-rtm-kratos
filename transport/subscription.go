package transport

import (
	"context"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

type IRtmMessageChannelService interface {
	OnMessageEvent(ctx context.Context, event *agrtm.MessageEvent) error
}

// SMessage Channel 订阅事件处理器
func newMessageChannelSubscription(s *Server, svc IRtmMessageChannelService, channelName string) agrtm.RtmEventHandler {
	return agrtm.RtmEventHandler{
		OnMessageEvent: func(event *agrtm.MessageEvent) {
			ctx := s.newServerContext(channelName, "OnMessageEvent")

			wrappedHandler := s.middlewareChain(func(ctx context.Context, req any) (any, error) {
				//call real handler
				err := svc.OnMessageEvent(ctx, req.(*agrtm.MessageEvent))
				return nil, err
			})
			_, err := wrappedHandler(ctx, event)
			if err != nil {
				s.logh.Errorw("err", err)
			}
		},
	}
}
