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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgeai"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &edgeai.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &edgeai.{ActionName}Response{}
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
| Createragfile | 通过该接口上传rag文件 | POST | /rag/file/create |
| Enabledisablefile | 启用文件或者禁用文件 | POST | /rag/file/operate |
| Updatefile | 更新文件。用户通过上传文件，和相关的文件属性，来更新知识库中已有的文件。 | POST | /rag/file/update |
| Listragfile | 查询文件列表 | POST | /rag/file/list |
| Ragfiledeleteservice | RAG文件删除接口 | POST | /rag/file/delete |
| Ragknowledgebasedeleteservice | 知识库删除接口，用户可以通过提供知识库 ID 来删除整个知识库。请谨慎操作。 | POST | /rag/knowledgebase/delete |
| Ragknowledgebasecreateservice | 知识库创建接口 | POST | /rag/knowledgebase/create |
| Ragknowledgebaselistservice | 知识库查询接口 | POST | /rag/knowledgebase/list |
| Ragknowledgebaseupdateservice | 知识库更新接口 | POST | /rag/knowledgebase/update |
| Createwrconnector | 该接口用于创建一个等候室连接器，支持配置QT账号信息、连接器名称、同步周期、是否验证用户身份以及队列令牌有效期等参数。成功后返回连接器的实例ID。 | POST | /api/v2/waiting_room_connectors |
| Updatewrconnector | 该接口用于更新现有等候室连接器的配置信息。用户需通过REST参数指定要更新的连接器实例ID，并通过请求体参数更新如QT账号ID、API密钥、是否生成队列令牌、同步周期等配置。接口响应将返回操作结果、响应数据和消息。 | PUT | /api/v2/waiting_room_connectors/* |
| Deletewaitingroomconnector | 该接口用于根据指定的连接器ID删除等候室连接器。用户需提供连接器的唯一ID作为REST参数进行删除操作。 | DELETE | /api/v2/waiting_room_connectors/* |
| Listconnectorsbypage | 该接口用于分页查询连接器列表。用户可以根据名称（支持模糊查询）、ID（支持精确查询）等条件进行筛选，并指定页码和每页大小。响应结果包含连接器总数及当前页的连接器详细信息。 | GET | /api/v2/waiting_room_connectors |
| Viewwrconnectorinfo | 该接口用于查看等候室连接器的详细信息。用户需通过RESTful路径参数`id`指定连接器实例，接口将返回该连接器的配置内容、同步周期、队列令牌有效期、状态、创建及修改时间等详细信息。 | GET | /api/v2/waiting_room_connectors/* |
| Addimageconfig | 该接口用于上传或修改图片处理配置。用户通过请求体中的property和policySets参数上传配置详情，系统将返回操作结果的响应码和信息。 | POST | /api/v2/ivm_configs |