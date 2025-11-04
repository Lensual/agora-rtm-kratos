package transport

import (
	"context"
	"errors"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmTokenBuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtmtokenbuilder2"
	rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
	"github.com/Lensual/agora-rtm-kratos/utils"
)

// RtmClientSupplier is a function that returns an IRtmClient instance.
//
//	NOTE: 由于 [agrtm.IRtmClient] 为单例模式，[transport.Server] 不应该主导全局的 [agrtm.IRtmClient] 实例创建过程。
//	所以这里提供了 [transport.RtmClientSupplier] 抽象接口。
//	[transport.RtmClientSupplier] 负责初始化 [agrtm.IRtmClient] 并完成相应的登录操作
//
// 1. 返回的 [agrtm.IRtmClient] 实例需要自行实现登录和 Token 逻辑
// 2. 返回的 [agrtm.IRtmClient] 需要为已登录状态
// TODO：确认Server需要的订阅操作是否真的是需要为已登录状态
type RtmClientSupplier func() (*agrtm.IRtmClient, error)

// DefaultRtmSupplier is a default implementation of RtmClientSupplier.
func DefaultRtmSupplier(appId, appCert, userId string) RtmClientSupplier {
	return func() (*agrtm.IRtmClient, error) {
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
			)

		// Create new [agrtm.IRtmClient] instance.
		rtmClient := agrtm.NewRtmClient(rtmConfig)

		ctx := context.Background()
		loginTimeoutCtx, _ := context.WithTimeoutCause(ctx, time.Second*3,
			errors.New("login timeout"))

		// Build token
		rtmToken, err := rtmTokenBuilder.BuildToken(appId, appCert, userId, 60*10) // 60s * 10
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
