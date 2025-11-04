//TODO 为代码添加注释
//TODO 为包编写文档注释

package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmTokenBuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtmtokenbuilder2"
	rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
	"github.com/Lensual/agora-rtm-kratos/example"
	rtmTr "github.com/Lensual/agora-rtm-kratos/transport"
	"github.com/Lensual/agora-rtm-kratos/utils"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
)

var (
	appId       string
	appCert     string
	userId      string
	channelName string
)

func init() {
	flag.StringVar(&appId, "appId", os.Getenv("APP_ID"), "Agora RTM App ID")
	flag.StringVar(&appCert, "appCert", os.Getenv("APP_CERT"), "Agora RTM App ID")
	flag.StringVar(&userId, "userId", os.Getenv("USER_ID"), "Agora RTM User ID")
	flag.StringVar(&channelName, "channelName", os.Getenv("CHANNEL_NAME"), "Agora RTM Channel Name")
}

func createRtmClient() (*agrtm.IRtmClient, error) {
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
			NewMyRtmEventHandler(), //////////////////
		)

	// Create new [agrtm.IRtmClient] instance.
	rtmClient := agrtm.NewRtmClient(rtmConfig)

	ctx := context.Background()
	loginTimeoutCtx, _ := context.WithTimeoutCause(ctx, time.Second*3,
		errors.New("login timeout"))

	// Build token
	rtmToken, err := rtmTokenBuilder.BuildToken(appId, appCert, userId, 60*10) // 60s * 10

	// RTM Login
	_, err = rtmAsync.Call[int](func() (int, uint64) {
		return rtmClient.Login(rtmToken)
	}).Await(loginTimeoutCtx)
	if err != nil {
		return nil, err
	}

	return rtmClient, nil
}

func main() {
	flag.Parse()
	if appId == "" || appCert == "" || userId == "" || channelName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	logger := log.With(log.GetLogger(),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	log.SetLogger(logger)

	rtmServer, err := rtmTr.NewServer(
		createRtmClient,
		rtmTr.Logger(logger),
		rtmTr.Middleware(
			logging.Server(log.GetLogger()),
		),
	)
	if err != nil {
		panic(err)
	}

	mySvc := &example.MyService{}
	rtmServer.Subscribe(channelName, mySvc)

	app := kratos.New(
		kratos.Server(
			rtmServer,
		),
	)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
