package utils

import (
	"context"
	"log"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmTokenBuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtmtokenbuilder2"
	rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
)

type TokenSupplier func() (string, error)

func NewTokenSupplier(appId, appCert, userId string) TokenSupplier {
	return func() (string, error) {
		return rtmTokenBuilder.BuildToken(appId, appCert, userId, 60*10) // 60s * 10
	}
}

func NewAutoRenewTokenHandler(rtmClientSupplier func() *agrtm.IRtmClient, tokenSupplier TokenSupplier) *agrtm.RtmEventHandler {
	return &agrtm.RtmEventHandler{
		OnTokenPrivilegeWillExpire: func(channelName string) {
			// Build token
			rtmToken, err := tokenSupplier()
			if err != nil {
				log.Printf("Failed to build token: %v", err)
				return
			}

			// Renew Token
			_, err = rtmAsync.Call[int](func() (ret int, reqId uint64) {
				return rtmClientSupplier().RenewToken(rtmToken)
			}).Await(context.TODO())
			if err != nil {
				log.Printf("Failed to renew token: %v", err)
				return
			}
		},
	}
}
