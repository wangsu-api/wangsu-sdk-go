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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/configcenter"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &configcenter.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &configcenter.{ActionName}Response{}
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
| Createconfigmap | 创建configmap | POST | /api/v1/namespaces/*/configmaps |
| Listconfigmap | 获取configmap列表 | GET | /api/v1/namespaces/*/configmaps |
| Getconfigmap | 获取configmap的详细信息 | GET | /api/v1/namespaces/*/configmaps/* |
| Updateconfigmap | 更新configmap | PUT | /api/v1/namespaces/*/configmaps/* |
| Deleteconfigmap | 删除configmap | DELETE | /api/v1/namespaces/*/configmaps/* |
| Listsecret | 获取secret列表 | GET | /api/v1/namespaces/*/secrets |
| Getsecret | 获取secret的详细信息 | GET | /api/v1/namespaces/*/secrets/* |
| Createsecret | 创建secret | POST | /api/v1/namespaces/*/secrets |
| Updatesecret | 更新secret | PUT | /api/v1/namespaces/*/secrets/* |
| Deletesecret | 删除secret | DELETE | /api/v1/namespaces/*/secrets/* |
| Putpatchconfigmap | 部分更新configmap | PUT | /api/v1/namespaces/*/configmaps/*/ws/patch |
| Putpatchsecret | 部分更新secret | PUT | /api/v1/namespaces/*/secrets/*/ws/patch |
| Gettoken | 获取认证token | GET | /openapi/custom/api/v1/token |