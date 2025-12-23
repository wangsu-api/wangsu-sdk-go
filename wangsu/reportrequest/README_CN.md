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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportrequest"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportrequest.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportrequest.{ActionName}Response{}
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
| Getrequestsbydomain | 查询请求数，输出5分钟粒度的请求数数据 | POST | /myview/hit |
| Querydomaintotalrequest | 该接口用于查询多个域名的总请求数。用户通过提供时间，域名等信息，接口将返回每个时间片段内的请求数，支持不同粒度的数据（如每五分钟、每小时或每天的数据）。该接口有助于用户监控网站访问情况，及时对异常情况进行跟进和优化。 | POST | /api/report/domainhit |
| Reportrequestispprovinceservice | 查询多域名服务节点归属各ISP各省份请求数及QPS；<br>数据延迟在5~15分钟左右；<br>支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。 | POST | /api/report/request/isp-province |
| Reportrequesthitrateispprovinceservice | 该接口允许用户查询域名在不同运营商和省份的请求数命中率，支持根据请求头Accept-Language返回中文或者英文数据。用户需提供时间范围和域名等参数获取不同时间维度的数据。返回内容包括每个域名在不同运营商及省份的请求数命中率。有助于用户分析网络请求的分布情况。 | POST | /api/report/request/hit-rate/isp-province |
| Queryrequestnumbersundershieldpop | 该接口用于查询多个域名在父节点的请求数。用户提供查询时间范围和域名列表，并可选择是否按域名分组输出。返回结果包括总请求数、请求数峰值，以及每5分钟的请求数。此接口有助于用户需要分析域名在父节点上的访问量，以调整缓存策略。 | POST | /api/report/request/parent-node |
| Queryrequestbyspecificprotocol | 查询多域名的指定传输协议的请求数，针对的是所有边缘节点的数据。 | POST | /api/report/request/protocol |
| Queryipv6requestofeachispandprovince | 根据访客访问日志，查询多域名访客IP归属各ISP各省份不同IP类型的请求数<br>支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。 | POST | /api/report/request/isp-province/ipv6 |
| Reporthitservice | 该接口用于查询账号下所有加速域名的请求数汇总信息。用户可以通过调用此接口获取指定时间段内的总请求数及每个时间片段的请求数详细数据。通过利用该接口，用户能够方便地监控和分析域名在不同时间段的请求量，以优化资源分配和提升性能表现。 | GET | /api/report/hit |
| Queryrequesthitratio | 该接口用于查询特定时间段内的分钟级别请求数缓存命中率，用户需提供开始时间、结束时间和域名信息，以获取请求命中率的数据。可选数据粒度1分钟、5分钟、1小时和1天。接口返回的数据信息包括每分钟的缓存命中率，帮助用户评估其域名命中CDN缓存的情况。 | POST | /api/report/request/hit-ratio/total |
| Reportstreamtasknumberservice | 该接口用于查询指定时间段内的直播转码任务数。用户需提供查询时间范围和域名获取转码任务详细数据。接口返回内容包括使用的编码格式、转码数量、每次转码的时长等。该接口有助于用户监控直播转码任务的数量和效率，从而优化直播服务的资源分配与性能提升。 | POST | /api/report/stream/task-number |
| Dispatchdomainareatimereport | 获取调度域名解析量 | POST | /clouddns/Dispatchdomainareatimereport |
| Policydetailreport | 获取调度报告api | POST | /clouddns/Policydetailreport |
| Reportrequesthttphttpsservice | 该接口用于查询多个域名在指定时间段内每分钟的HTTP和HTTPS请求数量统计。用户输入开始和结束时间及域名，返回数据可选按域名分组或合并数据，接口将返回域名的每分钟的HTTP和HTTPS请求数。 | POST | /api/report/request/http/https |
| Reportrequestquicservice | 该接口用于查询多个域名的QUIC请求数量和总请求数。用户需要提供开始时间、结束时间以及域名列表来获取数据。接口返回的数据包括每个时间戳的QUIC请求数和总请求数。有助于用户分析QUIC的使用情况，调整策略以提升网站的响应速度和用户体验。<br>建议调用频率：300次 / 5min | POST | /api/report/request/quic |
| Reportuserrequestispprovinceservice | 该接口用于查询多域名在中国大陆各省、各运营商的请求数，支持根据请求头Accept-Language返回中文或英文数据。用户需提供域名和查询时间范围，可以指定ISP和省份进行详细分析。返回结果包括每个域名在各省各ISP的请求数及其总数。有助于用户分析地域和运营商请求分布，优化流量管理和提升服务效率。<br> | POST | /api/report/domain/request/isp-province |
| Reportuserrequestcountryservice | 查询多域名访客IP归属各国家区域请求数 | POST | /api/report/request/country |
| Reportprotocolrequestservice | 该接口用于查询多个域名在特定时间段内按协议分组的请求数量。用户需提供查询时间范围、域名及数据查询粒度（如1分钟、5分钟、1小时、1天）。返回结果包括每个域名在各时间片内的总请求数，以及HTTP和HTTPS对应协议的请求数。有助于用户分析不同协议下的请求情况，为网站性能优化和流量分析提供数据支持。 | POST | /api/report/protocol/request |
| Reportprotocoloriginrequestservice | 该接口用于查询多个域名按照协议回源的请求数信息。用户可以通过指定时间范围和域名来进行查询。返回内容包含每个域名的每个时间片不同协议的回源请求数数据。有助于用户分析网络访问情况，回源数据流量，从而进行网站性能优化和资源配置。 | POST | /api/report/protocol/origin/request |
| Reportminuteuserrequestispprovinceservice | 查询分钟级别访客IP归属各省各ISP请求数 | POST | /api/minute/report/domain/request/isp-province |
| Queryedgerequesthitratioservice | 查询多域名分钟级别的在边缘节点命中缓存的请求数命中率 | POST | /api/report/request/edge-hit-ratio/total |