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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgefunc"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &edgefunc.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &edgefunc.{ActionName}Response{}
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
| Uploadfunccode | 函数代码上传接口，针对不使用 CloudIDE ，而是直接调用接口上传函数代码的场景 | POST | /edgefunc/upload |
| Createedgefunctrigger | 该接口用于创建函数触发器。用户需要提供域名和路由列表作为请求参数。成功创建后，接口将返回响应码和创建成功的触发器列表。 | POST | /api/v2/cdn_triggers |
| Queryedgefunctrigger | 该接口用于查询函数触发器列表。用户可以根据域名或函数名称进行筛选，并支持分页查询。响应中将返回触发器的详细信息列表。 | GET | /api/v2/cdn_triggers |
| Deletefuncdomaintrigger | 该接口用于删除一个已存在的函数触发器。用户需要通过REST参数 `id`指定待删除的函数触发器的唯一标识。 | DELETE | /api/v2/cdn_triggers/* |
| Queryfunctionlist | 该接口用于查询函数列表。用户可以通过函数ID进行精确查询，或者通过函数名称进行模糊查询。同时支持分页查询，需要指定页码和每页大小。接口返回函数列表、总数以及操作结果状态信息。 | GET | /api/v2/edge_funcs |
| Getedgefunctioninfo | 该接口用于查询指定边缘函数的信息。用户需要提供函数 ID (可通过Query或REST路径参数传递)，接口将返回该函数的详细信息，包括函数名称、创建时间、测试域名、备注、更新时间等。 | GET | /api/v2/edge_funcs/* |
| Edgefuncgetdebuglog | 该接口用于查询指定函数的调试日志。用户需要提供函数ID和调试会话ID以获取相应的日志内容。 | GET | /api/v2/edge_funcs/*/logs/* |
| Queryfunctioncode | 该接口用于查询函数的最新代码或当前部署的代码。用户需提供函数ID作为参数。 | GET | /api/v2/edge_funcs/*/files |
| Savefunctioncode | 该接口用于保存指定函数的代码。用户需提供函数ID。 | POST | /api/v2/edge_funcs/*/files |
| Deleteedgefunction | 该接口用于根据指定的函数ID删除函数。 | DELETE | /api/v2/edge_funcs/* |
| Listfunctiontemplates | 该接口用于查询边缘函数模板列表。用户可以通过指定页码 `pageNo`、每页大小 `pageSize`、模板名称 `name`（支持模糊查询）以及模板名称是否为英文 `isEnglish` 来进行分页和过滤查询。接口响应将返回符合条件的函数模板列表 `templates`，以及总记录数 `total`、当前页码 `pageNum`、每页大小 `pageSize`、操作状态码 `code` 和描述信息 `message` | GET | /api/v2/edge_funcs/templates |
| Createedgefunc | 该接口用于创建边缘函数。用户需要提供函数名称、模板名称、测试域名、备注以及函数别名等参数。成功创建后将返回函数ID。 | POST | /api/v2/edge_funcs |
| Edgefuncgetdebugurl | 该接口用于获取指定边缘函数的测试URL。用户需传入函数ID，接口将返回操作结果及相关状态信息。 | GET | /api/v2/edge_funcs/*/debug_url |
| Updateedgefunc | 该接口用于更新边缘函数。用户需要提供函数ID、函数名称、模板名称、测试域名、备注以及函数别名等参数。 | PUT | /api/v2/edge_funcs/* |