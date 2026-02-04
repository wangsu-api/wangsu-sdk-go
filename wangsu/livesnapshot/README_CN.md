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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/livesnapshot"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &livesnapshot.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &livesnapshot.{ActionName}Response{}
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
| Querysnapshotparametertemplate | 通过该接口可以查询客户账号下已经创建的所有截图参数模板。用户可以选择通过模板ID查询指定模板参数，或不带参数查询所有模板。接口返回所有符合条件的模板详情，包括模板名称、截图参数列表、存储配置等信息。 | GET | /api/v2/streams/snapshot-parameter-templates |
| Addsnapshotparametertemplate | 该接口用于创建直播截图相关的参数模板，以便提前配置截图、回调通知和存储等参数。用户可通过请求体参数指定模板名称、截图参数列表、存储配置以及回调通知地址等信息。接口成功后将返回创建的模板ID。 | POST | /api/v2/streams/snapshot-parameter-templates |
| Deletesnapshotparametertemplate | 该接口用于删除指定的截图参数模版。用户需提供待删除模版的ID。请注意，如果模版正在被使用或已关联截图规则，则无法直接删除。 | DELETE | /api/v2/streams/snapshot-parameter-templates/* |
| Modifysnapshotparametertemplate | 该接口用于编辑修改截图参数模版内容。用户可通过`templateId`指定待修改的模版，并通过请求体参数`templateName`、`snapshotParams`等更新模版配置，包括截图的宽度、高度、间隔、文件格式、存储桶、存储路径、拉流超时及回调通知等信息。当模版正在被使用时，禁止进行修改操作。 | PUT | /api/v2/streams/snapshot-parameter-templates/* |
| Modifysnapshotrules | 该接口用于修改已存在的录制规则。用户需通过规则ID指定要修改的规则，并可提供发布点、域名、流名、流名扩展参数、是否启用、模板ID等信息来更新规则内容。 | PUT | /api/v2/streams/snapshot-rules/* |
| Querysnapshotrule | 该接口用于查询已创建的截图规则。用户可以通过截图规则ID、域名、发布点或流名进行筛选，若未指定筛选条件，则默认查询所有规则。接口将返回符合条件的截图规则列表，每条规则包含拉流域名、创建时间、发布点、域名、是否启用状态、规则ID、模版ID、流名和流名扩展参数等信息。 | GET | /api/v2/streams/snapshot-rules |
| Stoprealtimesnapshot | 该接口用于结束正在进行的实时截图任务。用户需提供截图任务的唯一标识persistentId来指定要结束的任务。接口调用成功后，会返回API调用的整体响应码和响应信息。 | PUT | /api/v2/streams/snapshotings/* |
| Deletescreenshotrules | 该接口用于删除已经创建的截图规则，需指定待删除规则的唯一ID。请注意，禁止删除正在执行的规则。接口成功时返回成功状态码，失败时返回对应的错误码。 | DELETE | /api/v2/streams/snapshot-rules/* |
| Startrealtimesnapshot | 该接口用于对正在直播的流进行实时截图，用户可以通过指定拉流域名、发布点、流名及录制模板来选择截图目标。接口成功调用后，将返回截图任务ID及各流的截图状态。若需提前结束截图，可调用“结束实时截图”接口，否则截图任务将在直播结束后自动终止。 | POST | /api/v2/streams/snapshotings |
| Addsnapshotrules | 通过新增截图规则接口，您可以指定直播流进行截图。该接口支持设置推流域名、发布点和流名粒度的规则。用户需要提供如拉流域名、发布点、域名、是否启用、模板ID和流名等参数。当触发到对应规则后，系统会根据规则关联的截图模板进行视频截图，并返回生成的规则ID。 | POST | /api/v2/streams/snapshot-rules |