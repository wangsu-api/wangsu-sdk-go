# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/edgeai
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgeai"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &edgeai.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgeai.{ActionName}Response{}
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