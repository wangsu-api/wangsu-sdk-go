# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/reportvisitor
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportvisitor"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportvisitor.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportvisitor.{ActionName}Response{}
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
| Querytotalnumberofuniqueipundersingledomain | 该接口用于查询单个域名内的各流名的独立IP数量。用户提供具体的域名和时间，获取在此期间指定流的信息（以天为粒度）。接口返回的结果包括域名及其下对应流名的独立IP统计信息。有助于用户了解每路流的访问者分布，从而优化流量配置或评估流媒体服务的使用情况。 | POST | /api/report/visitor/total/stream |
| Reportuvispprovinceservice | 该接口用于查询多域名在不同省份和运营商的独立访客IP数。用户可指定时间范围、域名、省份、运营商，并选择按域名、省份或运营商分组。结果帮助分析区域访问趋势，优化内容分发。支持Accept-Language头（仅zh-CN和en-US），默认为zh-CN。Accept-Language为en-US时，省份和运营商参数及返回为代码，否则为中文。 | POST | /api/report/uv/isp-province |
| Reportiptopdetailsservice | 查询多域名的5分钟明细的访客 TOP IP | POST | /api/report/ip/top-details |
| Reportreferrertopdetailsservice | 这个接口用于查询域名每5分钟的TOP访客来源详情。用户需提供域名以及时间范围，返回内容包括每个refer来源的详细数据，可选值包含流量、带宽和请求数。该接口有助于用户分析站点流量，从而做出相应的业务决策。<br> | POST | /api/report/referrer/top-details |
| Reporturltopdetailsservice | 该接口用于统计多域名5分钟内的TOP URL信息。用户需提供时间范围和域名。返回内容包括域名的TOP URL的流量或带宽 或请求数信息。有助于用户了解热门访问链接，从而进行资源配置优化。 | POST | /api/report/url/top-details |
| Reportvisitorcustomtopdailyservice | 查询天粒度访客IP的自定义TOP排行 | POST | /api/report/visitor/custom-top/daily |