package main

import (
	"context"

	agorartm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmTr "github.com/Lensual/agora-rtm-kratos/transport"
	"github.com/go-kratos/kratos/v2/log"
)

var _ rtmTr.IRtmMessageChannelService = (*EchoService)(nil)

type EchoService struct {
}

// OnMessageEvent implements transport.IRtmMessageChannelService.
func (m *EchoService) OnMessageEvent(ctx context.Context, event *agorartm.MessageEvent) error {
	log.Info(event)
	return nil
}
