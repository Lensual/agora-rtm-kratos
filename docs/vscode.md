# vscode

launch.json

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "example-basic",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/example/basic",
            "envFile": "${workspaceFolder}/.env",
            "env": {
                "LD_LIBRARY_PATH":"${env:LD_LIBRARY_PATH}:${workspaceFolder}/Agora-RTM-Server-SDK-Go/agora_sdk",
            },
        },
        {
            "name": "example-server",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/example/server",
            "envFile": "${workspaceFolder}/.env",
            "env": {
                "LD_LIBRARY_PATH":"${env:LD_LIBRARY_PATH}:${workspaceFolder}/Agora-RTM-Server-SDK-Go/agora_sdk",
            },
        },
        {
            "name": "sdk-example",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/Agora-RTM-Server-SDK-Go/cmd/example",
            "envFile": "${workspaceFolder}/.env",
            "env": {
                "LD_LIBRARY_PATH":"${env:LD_LIBRARY_PATH}:${workspaceFolder}/Agora-RTM-Server-SDK-Go/agora_sdk",
            },
        },
    ]
}
```

settings.json

```json
{
    "go.toolsEnvVars": {
        "CGO_CFLAGS": "-O0 -g -I${workspaceFolder}/Agora-RTM-Server-SDK-Go/agora_sdk/agora_rtm_sdk_c/include",
        "CGO_LDFLAGS": "-O0 -g -L${workspaceFolder}/Agora-RTM-Server-SDK-Go/agora_sdk",
    }
}
```

extensions.json

```json
{
    "recommendations": [
        "golang.go",
    ]
}
```
