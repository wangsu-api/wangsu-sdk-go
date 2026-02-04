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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/recording"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &recording.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &recording.{ActionName}Response{}
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
| Getrecordtasklistquery | 通过录制任务查询接口，可以查询用户下的录制任务，查询的仅是客户从cloudv下发的录制任务，不包含即推记录的任务 | POST | /live/channelManage/recordTaskListQuery |
| Livevideoconcat | 录制文件合并 | POST | /live/channelManage/liveVideoConcat |
| Livevideoconcatquery | 查询录制合并任务 | POST | /live/channelManage/liveVideoConcatQuery |
| Addrecordingparametertemplate | 该接口用于新增录制参数模板。用户可以提前创建直播录制相关的参数模板，其中包含录制参数、回调通知参数、存储参数等。接口成功后将返回模板ID。 | POST | /api/v2/streams/recording-parameter-templates |
| Queryrecordingparametertemplate | 该接口用于查询客户账号下已创建的所有录制参数模板，或通过模板ID查询指定的模板参数。查询结果将返回模板的详细配置信息，包括模板名称、存储配置、录制参数等。 | GET | /api/v2/streams/recording-parameter-templates |
| Modifyrecordingparametertemplate | 该接口用于编辑修改录制模版参数内容。用户需指定模版ID，并可提交模版名称、云存储配置（如管理域名、AK/SK、存储空间名、文件名、存储时间）、回调通知地址以及具体的录制参数（如拉流超时设置、文件大小、分段时长、音视频流处理、文件格式等）。当模版正在被使用时，禁止修改。 | PUT | /api/v2/streams/recording-parameter-templates/* |
| Deleterecordingparametertemplate | 该接口用于删除录制参数模版。用户需通过路径参数 `templateId` 指定要删除的模版。如果模版正在被使用或已关联录制规则，则禁止删除。接口成功删除模版后将返回操作结果。 | DELETE | /api/v2/streams/recording-parameter-templates/* |
| Modifyrecordingrules | 该接口用于修改已存在的录制规则内容。用户需指定规则ID，并可选择性地更新发布点、域名、启用状态、模板ID或流名称及扩展参数。接口响应将返回操作结果。 | PUT | /api/v2/streams/recording-rules/* |
| Stoprealtimerecord | 该接口用于结束正在进行的实时录制任务。用户需提供录制任务的唯一标识persistentId来指定要结束的任务。接口调用成功后，会返回API调用的整体响应码和响应信息。 | PUT | /api/v2/streams/recordings/* |
| Addrecordingrules | 该接口用于新增直播流的录制规则。用户可指定推流域名、发布点、流名粒度规则等参数，并选择是否启用该规则及关联的录制模版。当规则被触发时，系统将根据关联的录制模板进行直播录制。成功新增规则后，将返回对应的规则ID。 | POST | /api/v2/streams/recording-rules |
| Queryrecordingrule | 通过该接口可以查询已经创建的录制规则，支持通过录制规则ID、域名、发布点、流名信息进行过滤筛选，若未指定任何筛选条件，将默认查询所有规则。响应将返回符合条件的录制规则列表，每条规则包含其ID、创建时间、发布点、域名、是否启用状态等详细信息。 | GET | /api/v2/streams/recording-rules |
| Deleterecordingrules | 该接口用于删除已创建的录制规则。用户需通过规则ID（`ruleId`）指定要删除的规则。请注意，禁止删除正在执行的规则。接口成功后将返回相应的响应码和响应信息。 | DELETE | /api/v2/streams/recording-rules/* |
| Startrealtimerecord | 该接口用于对正在直播的流进行实时录制，用户可以通过指定拉流域名、发布点、流名及录制模板来选择录制目标。接口成功调用后，将返回录制任务ID及各流的录制状态。若需提前结束录制，可调用“结束实时录制”接口，否则录制任务将在直播结束后自动终止。 | POST | /api/v2/streams/recordings |