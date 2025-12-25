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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgehostname"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &edgehostname.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgehostname.{ActionName}Response{}
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
| Deleteedgehostname | 删除指定的调度域名。删除指的是卸载DNS部署，同时删除数据。仅未关联任何加速域名的调度域名才允许删除。注意：删除后无法恢复，请谨慎执行删除操作。 | DELETE | /api/edge-hostnames/* |
| Deployedgehostnamedns | 部署调度域名DNS用于将调度域名的配置下发到全球边缘调度服务器上，部署完成后边缘调度服务将根据DNS配置策略执行流量调度。 | POST | /api/edge-hostnames/*/deploy |
| Undeployedgehostnamedns | 卸载调度域名用于将调度域名的配置从全球边缘调度服务器上清除，卸载完成后边缘调度服务不再执行流量调度。 | POST | /api/edge-hostnames/*/disable |
| Getedgehostname | 查询指定的调度域名的配置详情，包括：调度域名名称、调度域名描述、DNS服务状态、DNS部署状态、是否允许中国大陆加速、关联的加速域名列表、分区调度规则等。 | GET | /api/edge-hostnames/* |
| Listedgehostnames | 查询调度域名列表，返回参数包括：调度域名名称、调度域名描述、DNS服务状态、DNS部署状态、是否允许中国大陆加速、关联的加速域名列表等。 | GET | /api/edge-hostnames |
| Updateedgehostname | 修改指定的调度域名的配置，支持配置分区调度规则。注意，此接口禁修改配置数据，修改完后您需再调用部署调度域名接口触发部署。 | PUT | /api/edge-hostnames/* |
| Deployedgehostnameforterraform | 该接口用于Terraform场景部署指定的调度域名。部署指的是下发调度域名配置并使其DNS服务生效。部署后，使用此调度域名的加速域名可以正常使用CDN服务。同时，当您的调度域名部署状态为待部署、部署失败时，可使用此接口重新部署调度域名。 | POST | /api/terraform/edge-hostnames/*/deploy |
| Deleteedgehostnameforterraform | 该接口用于Terraform场景删除指定的调度域名。删除指的是卸载DNS部署，同时删除数据。仅未关联任何加速域名的调度域名才允许删除。注意：删除后无法恢复，请谨慎执行删除操作。 | DELETE | /api/terraform/edge-hostnames/* |
| Queryedgehostnameforterraform | 该接口用于Terraform场景查询调度域名详情。用户需指定调度域名进行查询。 | GET | /api/terraform/edge-hostnames/* |
| Updateedgehostnameforterraform | 该接口用于Terraform场景修改指定的调度域名的配置，支持配置分区调度规则。用户可指定调度域名配置与分区调度规则进行修改。 | PUT | /api/terraform/edge-hostnames/* |
| Queryedgehostnamesforterraform | 该接口用于Terraform场景查询调度域名列表，返回参数包括：调度域名名称、调度域名描述、DNS服务状态、DNS部署状态、是否允许中国大陆加速、关联的加速域名列表等。 | GET | /api/terraform/edge-hostnames |