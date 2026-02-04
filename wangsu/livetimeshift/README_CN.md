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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/livetimeshift"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &livetimeshift.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &livetimeshift.{ActionName}Response{}
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
| Querytimeshiftparametertemplate | 该接口用于查询客户账号下已创建的所有时移参数模板。用户可通过可选的模板ID查询指定模板参数，也可通过区域参数对模板进行筛选。成功调用将返回包含所有时移参数模板或指定模板详情的列表。 | GET | /api/v2/streams/time-shift-parameter-template |
| Modifytimeshiftparametertemplate | 该接口用于编辑和修改时移模版参数内容。为避免影响正在进行的业务，当模版处于使用中时，禁止进行修改操作。接口通过URL路径中的模版ID (`templateId`) 来识别目标模版，并接收请求体中传递的模版名称、存储配置、回调地址以及时移参数列表等详细信息进行更新。 | PUT | /api/v2/streams/time-shift-parameter-template/* |
| Addtimeshiftparametertemplate | 该接口用于提前创建直播时移相关的参数模板。这些模板主要用于保存时移参数、回调通知参数、以及存储参数等。成功创建后，接口将返回生成的模板ID。 | POST | /api/v2/streams/time-shift-parameter-template |
| Deletetimeshiftparametertemplate | 该接口用于删除指定的时移参数模版。用户需在路径参数中提供待删除模版的唯一标识符`templateId`。如果模版**正在被使用**或已关联时移规则，则禁止直接删除。接口成功删除后，响应体将返回状态码`code`和响应信息`message`。 | DELETE | /api/v2/streams/time-shift-parameter-template/* |
| Querytimeshiftrule | 该接口用于查询已经创建的时移规则。支持通过录制规则ID、域名、发布点和流名进行筛选。如果未指定任何筛选条件，将默认查询所有时移规则。成功查询将返回符合条件的规则列表，包含各规则的详细信息，如创建时间、启用状态等。 | GET | /api/v2/streams/time-shift-rules |
| Stoprealtimetimeshiftrecording | 该接口用于结束正在进行的实时时移录制任务。用户需提供时移录制任务的唯一标识persistentId来指定要结束的任务。接口调用成功后，会返回API调用的整体响应码和响应信息。 | PUT | /api/v2/streams/time-shift-recordings/* |
| Startrealtimetimeshiftrecording | 该接口用于对正在直播的流进行实时时移录制，用户可以通过指定拉流域名、发布点、流名及时移录制模板来选择录制目标。接口成功调用后，将返回录制任务ID及各流的时移录制状态。若需提前结束录制，可调用“结束实时时移录制”接口，否则时移录制任务将在直播结束后自动终止。 | POST | /api/v2/streams/time-shift-recordings |
| Addtimeshiftrules | 通过该接口可以新增时移规则，用于指定直播流进行录制并实现时移功能。用户需提供包括推流域名、发布点、流名等规则参数，并可指定关联的时移模版ID和规则启用状态。当对应的规则被触发后，系统将根据关联的时移模版进行视频录制。 | POST | /api/v2/streams/time-shift-rules |
| Deletetimeshiftrules | 该接口用于删除已经创建的时移规则，但禁止删除正在执行中的规则。用户需通过规则ID指定待删除的时移规则，接口将返回删除操作的成功或失败信息。 | DELETE | /api/v2/streams/time-shift-rules/* |
| Modifytimeshiftrules | 该接口用于修改已存在的时移规则。用户需要提供要修改规则的唯一标识符（ruleId），并通过请求体指定需要更新的规则属性，例如发布点、域名、启用状态、模板ID和流相关参数等。 | PUT | /api/v2/streams/time-shift-rules/* |