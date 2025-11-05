//TODO 为代码添加注释
//TODO 为包编写文档注释

package main

import (
	"flag"
	"os"

	rtmTr "github.com/Lensual/agora-rtm-kratos/transport"
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
		// 这里使用默认的 RTM 实例创建方式。如果需要，可以传入自定义的 RTM 实例。
		rtmTr.DefaultRtmSupplier(appId, appCert, userId),
		rtmTr.Logger(logger),
		rtmTr.Middleware(
			// 这里可以添加 kratos 中间件
			logging.Server(logger),
		),
	)
	if err != nil {
		panic(err)
	}

	mySvc := &EchoService{}
	rtmServer.Subscribe(channelName, mySvc)

	app := kratos.New(
		kratos.Logger(logger),
		kratos.Server(
			rtmServer,
		),
	)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
