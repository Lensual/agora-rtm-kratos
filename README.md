# agora-rtm-kratos

[![GoDoc](https://godoc.org/github.com/Lensual/agora-rtm-kratos?status.svg)](https://pkg.go.dev/github.com/Lensual/agora-rtm-kratos)

Agora RTM SDK for Kratos.

当然也可以直接使用 utils 和 async 包！

## Feature

- 支持 kratos transport
- 回调事件增强，易用
  - 可组合注册多个回调监听 `RtmEventHandler`
  - 对回调接口抽象成 interface
- 更完善的异步调用机制

## Notice

- RTM UID 只能在一处登录
- 同一个进程空间下，RTM 实例只能有一个。不要重复初始化。
-

## Usage

Example code can be found [here](./example/basic/).

Install

```sh
go get -u github.com/Lensual/agora-rtm-kratos
```

Transport Server

```go
import (
  rtmTr "github.com/Lensual/agora-rtm-kratos/transport"
  "github.com/go-kratos/kratos/v2"
)
// ...

rtmClientSupplier := rtmTr.DefaultRtmSupplier(appId, appCert, userId)
rtmTrSrv := rtmTr.NewServer(rtmClientSupplier)
app := kratos.New(
  kratos.Server(
    rtmTrSrv,
  ),
)
app.Run()
```

Async Call

```go
import rtmAsync "github.com/Lensual/agora-rtm-kratos/async"
// ...

// RTM Login
promise := rtmAsync.Call[int](func() (int, uint64) {
  return rtmClient.Login(rtmToken)
})

err := promise.Await()
if err != nil {
  return nil, err
}
```

组合多个回调监听

```go
import (
  agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
  "github.com/Lensual/agora-rtm-kratos/utils"
)
// ...

rtmConfig := agrtm.NewRtmConfig()
rtmConfig.EventHandler = utils.NewCompositeRtmEventHandler(
  // Handle A
  // Handle B
  // ...
)
```

## 自定义 RTM 实例

默认情况下，由 `transport.Server` 负责初始化 RTM 实例。如果需要自定义 RTM 实例，请自行实现 `transport.RtmClientSupplier`。通常来说你需要：

- 实现 `agrtm.RtmConfig` 的创建过程
- 实现 `agrtm.IRtmClient` 的创建过程
- 实现 `agrtm.IRtmEventHandler` 的创建过程
- 实现 RTM 登录、Token 更新、掉线处理逻辑。
- 保证返回的 `agrtm.IRtmClient` 实例为已登录状态。 // TODO：确认Server需要的订阅操作是否真的是需要为已登录状态

## TODO

- [x] token 更新逻辑
- [ ] async，需要将多个回调参数合并成一个结构体。
- [ ] 提供 Stream Channel
- [ ] 提供 Storage UserMetadata 事件订阅
- [ ] Message Channel
  - [ ] 支持以点对点方式回复消息
  - [ ] 支持 encoding DTO
- [ ] 为 transport 的 header 增加一些 metadata
- [ ] 完善 Example 注释
