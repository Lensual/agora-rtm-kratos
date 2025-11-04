package transport

import (
	"context"
	"net/http"

	ktransport "github.com/go-kratos/kratos/v2/transport"
)

const KindRtm ktransport.Kind = "rtm"
const SubKindMessageChannel ktransport.Kind = "MessageChannel"
const SubKindStreamChannel ktransport.Kind = "StreamChannel"

var _ ktransport.Transporter = (*Transport)(nil)

// Transport implements [ktransport.Transporter].
type Transport struct {
	endpoint    string
	reqHeader   headerCarrier
	replyHeader headerCarrier

	topic   string
	subKind ktransport.Kind
	server  *Server
}

// Kind implements [ktransport.Transporter].
func (tr *Transport) Kind() ktransport.Kind {
	return KindRtm
}

// Endpoint implements [ktransport.Transporter].
func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

// Operation implements [ktransport.Transporter].
func (tr *Transport) Operation() string {
	return tr.topic
}

// RequestHeader implements [ktransport.Transporter].
func (tr *Transport) RequestHeader() ktransport.Header {
	return tr.reqHeader
}

// ReplyHeader implements [ktransport.Transporter].
func (tr *Transport) ReplyHeader() ktransport.Header {
	return tr.replyHeader
}

func (tr *Transport) Topic() string {
	return tr.topic
}

func (tr *Transport) SubKind() ktransport.Kind {
	return tr.subKind
}

func (tr *Transport) Server() *Server {
	return tr.server
}

// ServerFromServerContext returns [Server] from context.
func ServerFromServerContext(ctx context.Context) (*Server, bool) {
	if tr, ok := ktransport.FromServerContext(ctx); ok {
		if tr, ok := tr.(*Transport); ok {
			return tr.Server(), true
		}
	}
	return nil, false
}

type headerCarrier http.Header

// Get returns the value associated with the passed key.
func (hc headerCarrier) Get(key string) string {
	return http.Header(hc).Get(key)
}

// Set stores the key-value pair.
func (hc headerCarrier) Set(key string, value string) {
	http.Header(hc).Set(key, value)
}

// Add append value to key-values pair.
func (hc headerCarrier) Add(key string, value string) {
	http.Header(hc).Add(key, value)
}

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of values associated with the passed key.
func (hc headerCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}
