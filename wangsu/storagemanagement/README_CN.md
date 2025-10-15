# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/storagemanagement
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
    request := &storagemanagement.ActionNameRequest{}

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