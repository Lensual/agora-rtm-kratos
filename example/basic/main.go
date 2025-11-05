//TODO 为代码添加注释
//TODO 为包编写文档注释

package main

import (
	"context"
	"flag"
	"os"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
	rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
	rtmTr "github.com/Lensual/agora-rtm-kratos/transport"
	"github.com/go-kratos/kratos/v2/log"
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

	// Create RTM Client
	rtmClient, err := rtmTr.DefaultRtmSupplier(appId, appCert, userId)()
	if err != nil {
		panic(err)
	}
	defer rtmClient.Logout()
	defer rtmClient.Release()

	ctx := context.Background()

	// Example for async/await with Subscribe Message Channel
	subscribeRes, err := rtmAsync.Call[int](func() (ret int, reqId uint64) {
		return rtmClient.Subscribe(channelName, agrtm.NewSubscribeOptions())
	}).Await(ctx)
	if err != nil {
		panic(err)
	}
	log.Infof("Subscribe result: %+v", subscribeRes)

	// Example for async/await with Publish on Message Channel
	publishRes, err := rtmAsync.Call[int](func() (ret int, reqId uint64) {
		return rtmClient.Publish(channelName, []byte("hello world!"), &agrtm.PublishOptions{
			ChannelType:    agrtm.RtmChannelTypeMESSAGE,
			MessageType:    agrtm.RtmMessageTypeSTRING,
			CustomType:     "",
			StoreInHistory: false,
		})
	}).Await(ctx)
	if err != nil {
		panic(err)
	}
	log.Infof("Publish result: %+v", publishRes)

	// Example for async/await with WhoNow API
	whoNowRes, err := rtmAsync.Call[int](func() (ret int, reqId uint64) {
		ret = rtmClient.GetPresence().WhoNow(channelName,
			agrtm.RtmChannelTypeMESSAGE,
			agrtm.NewPresenceOptions(),
			&reqId,
		)
		return ret, reqId
	}).Await(ctx)
	if err != nil {
		panic(err)
	}
	log.Infof("WhoNow result: %+v", whoNowRes)
}
