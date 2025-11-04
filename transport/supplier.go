package transport

import (
	"context"
	"errors"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
	"github.com/Lensual/agora-rtm-kratos/utils"
)

// RtmClientSupplier defines a function that returns an initialized and logged-in [agrtm.IRtmClient].
//
// 该函数负责创建并初始化 [agrtm.IRtmClient] 实例，并在返回前完成登录。
//
// Before returning, make sure that:
//  1. Login and RenewToken are properly implemented;
//  2. The returned client is already in a logged-in state.
//
// See [NewServer].
type RtmClientSupplier func() (*agrtm.IRtmClient, error)

// DefaultRtmSupplier is a default implementation of [RtmClientSupplier].
func DefaultRtmSupplier(appId, appCert, userId string) RtmClientSupplier {
	return func() (*agrtm.IRtmClient, error) {
		tokenSupplier := utils.NewTokenSupplier(appId, appCert, userId)

		var rtmClient *agrtm.IRtmClient

		// Create new [agrtm.RtmConfig]
		rtmConfig := agrtm.NewRtmConfig()
		rtmConfig.AppId = appId
		rtmConfig.UserId = userId
		rtmConfig.EventHandler =
			// Combine multiple handlers
			utils.NewCompositeRtmEventHandler(
				// AsyncCallRtmEventHandler provides enhanced async/await capabilities.
				// See [rtmAsync] for more details.
				rtmAsync.NewAsyncCallRtmEventHandler(),

				// AutoRenewTokenHandler renews the token when it is about to expire.
				// See [utils.NewAutoRenewTokenHandler] for more details.
				utils.NewAutoRenewTokenHandler(func() *agrtm.IRtmClient {
					return rtmClient
				}, tokenSupplier),

				// For Kratos, you can use the `logging.Server()` middleware to handle logging.
				// Uncomment the following line to debug logging.
				// utils.NewLoggerHandler(nil),
			)

		// Create new [agrtm.IRtmClient] instance.
		rtmClient = agrtm.NewRtmClient(rtmConfig)

		ctx := context.Background()
		loginTimeoutCtx, _ := context.WithTimeoutCause(ctx, time.Second*3,
			errors.New("login timeout"))

		// Build token
		rtmToken, err := tokenSupplier()
		if err != nil {
			return nil, err
		}

		// RTM Login
		_, err = rtmAsync.Call[int](func() (int, uint64) {
			return rtmClient.Login(rtmToken)
		}).Await(loginTimeoutCtx)
		if err != nil {
			return nil, err
		}

		return rtmClient, nil
	}
}
