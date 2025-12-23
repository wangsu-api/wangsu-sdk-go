# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/devicemanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &devicemanagement.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := devicemanagement.{ActionName}Response{}
    _, err := auth.Invoke(config, request, response)

    // Handle response
    if err != nil {
        log.Printf("error: %s\n", err)
        return
    }

    log.Printf("response body: %s\n", response.String())
}
```

## 错误处理

始终检查 API 调用返回的错误：

```go
_, err := auth.Invoke(config, request, response)
if err != nil {
    log.Printf("error: %s\n", err)
    // Handle the error appropriately
    return
}
```

## API列表
有关详细的 API 文档和可用方法，请参阅[官方 Wangsu API 文档](https://www.wangsu.com/document/api-doc/Overview?productType=all)。

| ActionName | description | client_methods | uri |
| --- | --- | --- | --- |
| Getstreamurl | 获取通道播放url | POST | /ivcs/devicemanage/getStreamUrl |
| Editdevice | 编辑设备信息 | POST | /ivcs/devicemanage/editDevice |
| Getdeviceslist | 获取设备列表 | POST | /ivcs/devicemanage/getDevicesList |
| Invitepush | 邀请设备通道推流 | POST | /ivcs/devicemanage/invitePush |
| Stoppush | 停止设备通道推流 | POST | /ivcs/devicemanage/stopPush |
| Createdevice | 新增设备 | POST | /ivcs/devicemanage/createDevice |
| Deletedevice | 删除设备 | POST | /ivcs/devicemanage/deleteDevice |
| Querychannellist | 获取通道列表 | POST | /ivcs/devicemanage/getChnList |
| Editchannelinfo | 编辑通道信息 | POST | /ivcs/devicemanage/editChnInfo |
| Ptzcontrol | 云台控制 | POST | /ivcs/devicemanage/Ptzcontrol |
| ChannelControl | 通道禁用/启用 | POST | /ivcs/devicemanage/channelControl |
| ChannelStatusUpdate | 该接口用于更新通道信息，可以通过该接口主动更新设备下的通道列表和通道状态。主要用于设备在通道新增、删除、或者上下线时没有向平台汇报，或者汇报失败，用户可以通过改接口主动让设备再重新更新汇报一次。 | POST | /ivcs/devicemanage/channelStatusUpdate |