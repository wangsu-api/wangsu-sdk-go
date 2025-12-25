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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/storagemanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &storagemanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := storagemanagement.{ActionName}Response{}
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
| Listpvcs | 获取pvc列表 | GET | /api/v1/namespaces/*/persistentvolumeclaims |
| Createpvcs | 创建pvc | POST | /api/v1/namespaces/*/persistentvolumeclaims |
| Getpvcs | 获取pvc的详细信息 | GET | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Updatepvcs | 更新pvc | PUT | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Deletepvcs | 删除pvc | DELETE | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Putpatchpvcs | 部分更新pvc | PUT | /api/v1/namespaces/*/persistentvolumeclaims/*/ws/patch |
| Pagingpvcs | pvc列表分页查询 | GET | /openapi/custom/api/v1/persistentvolumeclaims |
| Pvcinnamespace | 获取namespace下的pvc列表 | GET | /openapi/custom/api/v1/namespaces/*/persistentvolumeclaims |
| Deletepvcfromedge | 直接从边缘删除pvc | DELETE | /openapi/custom/api/v1/namespaces/*/persistentvolumeclaims/* |
| Liststorageclass | 获取storageClass列表 | GET | /openapi/custom/api/v1/storageclasses |
| Querychannelrecordfiles | 查询通道录制文件 | POST | /ivcs/report/origin/query-channel-record |
| Querydevicerecordfiles | 用于查询gb28181设备端录像文件。接口一次最长只能查询7天数据，避免文件太多设备返回失败，最好是按天查询。如果返回内容为空说明设备没有录像文件或者设备端有异常，建议咨询设备生产服务厂商。 | POST | /ivcs/report/origin/query-device-record |
| Playbackcontrol | 通过该接口可以设置回放进度、回放倍率、回放暂停、回放继续播放。为了提高用户体验，回放控制时尽量使用webrtc拉流协议，降低数据延迟。 | POST | /ivcs/report/origin/playback-control |
| Startplayback | 开始回放指定设备端录制文件，该接口是下发指令给设备端，让设备端开始推录制文件视频流。为了保证设备推流带宽稳定，当前一个通道只能支持一路流进行回放，如果要回放该通道其他文件，需要结束已经在回放的那路流，才能开始回放新的文件 | POST | /ivcs/report/origin/start-playback |
| Stopplayback | 结束回放，下发信令让设备停止推录像文件视频流。是否断流成功可以通过回调信息进行判断。 | POST | /ivcs/report/origin/stop-playback |
| Queryplaybacklist | 通过调用查询该接口可以查询正在进行设备录像回放的流，可以通过查询接口对不需要回放的流进行结束回放，避免占用设备带宽资源。 | POST | /ivcs/report/origin/query-playback-list |