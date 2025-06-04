# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/basicmonitor
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/basicmonitor"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &basicmonitor.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := basicmonitor.{ActionName}Response{}
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
| Reportflowdomainispprovinceiaasservice | 该接口用于查询边缘服务器在各ISP和省份的流量数据，支持根据请求头Accept-Language返回中文或英文数据。用户需提供查询的时间范围、域名、ISP和省份等信息，返回内容包含每个ISP和省份的流量明细，以MB为单位，并按5分钟的时间粒度显示。适用于需要分析不同地区和运营商流量分布的用户，帮助优化网络资源和提升服务效率。 | POST | /api/report/flow/domain-isp-province/iaas |
| Vmpqueryservermetric | 查询云主机实例的CPU使用率、内存使用情况和带宽使用情况，以便接入自身的监控系统，掌控云主机的运行情况。提供近90天的监控数据查询，单次查询范围不超过3天，数据粒度为5分钟。如果是裸机实例，当前仅支持查询带宽，不支持查询CPU和内存。 | GET | /vmp/servers/metric |