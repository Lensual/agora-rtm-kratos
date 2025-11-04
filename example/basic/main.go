//TODO 为代码添加注释
//TODO 为包编写文档注释

package main

import (
	"flag"
	"os"

	"github.com/Lensual/agora-rtm-kratos/example"
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
		// NOTE: 由于 [agrtm.IRtmClient] 为单例模式，[rtmTr.Server] 不应该主导全局的 [agrtm.IRtmClient] 实例创建过程。
		// 所以这里提供了 [transport.RtmClientSupplier] 抽象接口。
		// [transport.RtmClientSupplier] 负责初始化 [agrtm.IRtmClient] 并完成相应的登录操作。
		rtmTr.DefaultRtmSupplier(appId, appCert, userId),
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
		kratos.Logger(log.GetLogger()),
		kratos.Server(
			rtmServer,
		),
	)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
