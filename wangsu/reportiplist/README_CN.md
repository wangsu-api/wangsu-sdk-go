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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportiplist"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportiplist.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportiplist.{ActionName}Response{}
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
| Reportserveripispprovinceservice | 查询CDN在各ISP各省份的服务IP列表 | POST | /api/report/server-ip/isp-province |
| Reportserveripexistflowservice | 该接口可以通过提供CDN的域名来获取对应的有流量的CDN服务IP列表。此功能适用于需要了解域名实际加速使用服务IP的场景。 | POST | /api/report/server-list/exist-flow |
| Querycdniplist | 该接口用于查询CDN服务中每个指定的域名对应的覆盖节点IP列表。用户可以通过提交相关请求获取其账号下多域名的节点IP信息，主要用于流量调度和优化管理。返回的数据显示了每个域名的详细覆盖节点IP，帮助用户了解和监控CDN服务的节点分布。 | POST | /api/report/server-list |
| Reportserverlistservice | 该接口用于查询特定域名在CDN服务中的IP列表及其归属运营商和区域信息。通过该接口，用户可以获取与域名相关联的CDN服务节点的详细信息。 | GET | /api/report/server-list/ip-isp-area |