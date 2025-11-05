module github.com/Lensual/agora-rtm-kratos/transport

go 1.22.4

replace github.com/Lensual/agora-rtm-kratos => ../

// agora rtm
replace github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go => ../Agora-RTM-Server-SDK-Go

require (
	github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go v0.0.0-00010101000000-000000000000
	github.com/Lensual/agora-rtm-kratos v0.0.0-00010101000000-000000000000
	github.com/go-kratos/kratos/v2 v2.9.1
)

require (
	github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src v0.0.0-20250825033728-374cd21f5220 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
