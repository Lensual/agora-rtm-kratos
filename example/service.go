package example

import (
	"context"

	agorartm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmTr "github.com/Lensual/agora-rtm-kratos/transport"
	"github.com/go-kratos/kratos/v2/log"
)

var _ rtmTr.IRtmMessageChannelService = (*MyService)(nil)

type MyService struct {
}

// OnMessageEvent implements transport.IRtmMessageChannelService.
func (m *MyService) OnMessageEvent(ctx context.Context, event *agorartm.MessageEvent) error {
	log.Info(event)
	return nil
}
