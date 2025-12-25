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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/propertyconfig"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &propertyconfig.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := propertyconfig.{ActionName}Response{}
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
| Createproperty | 创建项目 | POST | /api/properties |
| Listproperties | 查询项目列表 | GET | /api/properties |
| Deleteproperty | 删除项目 | DELETE | /api/properties/* |
| Getproperty | 查询某个项目的信息 | GET | /api/properties/* |
| Listpropertyversions | 查询项目版本列表 | GET | /api/properties/*/versions |
| Createpropertyversion | 创建一个新的项目版本。 | POST | /api/properties/*/versions |
| Getpropertyversion | 获取项目版本的详细配置信息。 | GET | /api/properties/*/versions/* |
| Updatepropertyversion | 更新项目版本的配置信息 | PUT | /api/properties/*/versions/* |
| Createdeploymenttask | 创建部署任务将项目部署到演练或生产环境，或将已部署的项目卸载。 | POST | /api/properties/deployments |
| Listdeploymenttasks | 该接口用于查询符合条件的部署任务列表。用户可以根据加速项目标识、任务状态、部署环境等条件进行筛选，并可控制查询结果的起始位置、最大条数、排序方式和排序字段。接口返回部署任务的总数以及详细的任务列表。 | GET | /api/properties/deployments |
| Getdeploymenttask | 获取部署任务详细信息 | GET | /api/properties/deployments/* |
| Createpropertyforakamaimigration | 创建Akamai China CDN迁移项目 | POST | /api/properties/migration |
| Createdeploymenttaskforterraform | 该接口用于Terraform场景创建部署任务将项目部署到演练或生产环境，或将已部署的项目卸载。用户需指定项目ID，版本与操作类型。 | POST | /api/terraform/properties/deployments |
| Createpropertyforterraform | 该接口用于Terraform场景创建项目。用户需指定域名配置、规则配置、回源配置、变量配置等信息。 | POST | /api/terraform/properties |
| Deletepropertyforterraform | 该接口用于Terraform场景删除项目配置。用户需指定项目ID来删除。 | DELETE | /api/terraform/properties/* |
| Updatepropertyforterraform | 该接口用于Terraform场景修改项目。用户需指定域名配置、规则配置、回源配置、变量配置等信息。 | PUT | /api/terraform/properties/* |
| Querydeploymentforterraform | 该接口用于Terraform场景查询部署任务详情。用户需指定部署任务ID来查询。 | GET | /api/terraform/properties/deployments/* |
| Querypropertiesforterraform | 该接口用于Terraform场景查询项目列表。用户可以指定服务类型、部署环境进行过滤查询。 | GET | /api/terraform/properties |
| Querypropertyversionconfigforterrform | 该接口用于Terraform场景查询项目版本配置。用户需指定项目ID与项目版本来查询。 | GET | /api/terraform/properties/*/versions/* |
| Querydeploymentsforterraform | 该接口用于Terraform场景查询部署任务列表。用户可指定项目ID，部署环境等进行过滤。 | GET | /api/terraform/properties/deployments |
| Querypropertyconfigforterrform | 该接口用于Terraform场景查询项目配置。用户需指定项目ID来查询。 | GET | /api/terraform/properties/* |
| Querytierroutemaps | 该接口用于查询项目配置中可用的层级回源路由列表。用户需指定服务类型或者项目ID进行查询。 | GET | /api/maps |
| Queryipsegmentbyroutemapcode | 该接口用于查询回源路由关联的IP段。用户指定回源路由编码来查询关联的所有IP段。 | GET | /api/origin_shields/maps/ip_segments |