# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/other
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/other"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &other.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := other.{ActionName}Response{}
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
| Attoverview | 查询安全防护的请求次数和被拦截次数 | POST | /myview/Attoverview |
| Bandwidthappa | 查询客户的appa加速的带宽 | POST | /myview/Bandwidthappa |
| Querydailylivetranscodingduration | 直播转码时长每日统计。 | POST | /myview/LiveTranscodingPerDayV2 |
| Query5minlivetranscodingduration | 输出客户分钟粒度的直播转码时长。 | POST | /myview/LiveTranscodingV2 |
| Wafhit | 查询5分钟粒度的waf请求数 | POST | /myview/Wafhit |
| Reportdomainlistexistflowservice | 查询指定时间段内有量的域名列表。<br>数据延迟5~15分钟左右。建议接口调用频率不高于30次/5分钟。 | POST | /api/report/domainList/existFlow |
| Reportstreamlistservice | 该接口用于查询指定域名下的流名列表，帮助用户获取域名在某段时间内在中国大陆的所有流名。用户需要提供查询时间范围和域名信息，时间跨度最大为3天。返回结果包含所查询域名的流名列表，有助于用户进行直播流的管理和分析。 | POST | /api/report/streamList |
| Reporturldlfinishservice | 该接口用于按下载完成数从多个域名中查询排名前500的URL。用户需要提供开始时间、结束时间和待查询的域名列表。接口将返回每个URL及其对应的下载成功数，帮助用户了解哪些URL在特定时段内的访问情况。通过调用该接口，用户可以有效地识别热度高的内容或识别异常行为，从而进行相应决策。 | POST | /api/report/url/dl-finish |
| Queryliverecordingduration | 输出客户分钟粒度的直播录制时长 | POST | /myview/RecordingTime |
| Flowappachannel | 查询分频道appa流量，输出每个频道查询周期内的边缘上行、边缘下行、总流量值 | POST | /myview/Flowappachannel |
| Concurrentsession | 查询appa并发数情况，输出计费值和分钟粒度的并发数 | POST | /myview/Concurrentsession |
| Clouddirectduration | 查询云导播时长。输出1分钟粒度的导播时长与总时长。 | POST | /myview/Clouddirectduration |
| Picprocessstatistics | 查询域名5分钟粒度图片处理次数。 | POST | /myview/PicFlowHit |
| Httpdnsstatistics | 查询5分钟粒度httpdns解析量。 | POST | /myview/HttpdnsStatistics |
| Bandwidthupload | 查询上传带宽，输出计费方式、计费值、总流量，以及分钟粒度的带宽值（可指定边缘上行、中间上行） | POST | /myview/bandwidth-upload |
| Reportavgspeeddomainispprovinceservice | 该接口用于查询特定域名在不同时段的平均下载速度，用户可以指定时间范围、运营商、域名以及分组维度来获取精确的下载速度数据。返回的数据包括在指定时间内的每五分钟平均下载速度、平均响应时间、平均首包响应时间及总响应时间，帮助用户了解不同地区和运营商的网络性能情况。此接口适用于想要分析网站性能、诊断网络问题或优化用户体验的场景。 | POST | /api/report/avg-speed/domain-isp-province |
| Querylivestreamstatus | 该API 具体提供包括码率、帧率、在线人数、头部信息四类直播流状态信息<br>四类信息可单独开通，供选择性查询 | POST | /api/quality/stream-status-statistic |
| Getlivestreampushingstatus | 获取指定时刻，指定域名下（支持多个域名和流名）的所有频道的帧率、码率、丢帧率等信息 | POST | /api/quality/frame-rate |
| Queryonlineviewercount | 提供直播域名或者频道的在线人数接口，接口只支持GET请求方式 | POST | /QOSS/api/onlineViewers |
| Queryddosmitigatedbandwidthbysuiteorproduct | 查看套餐或产品下所有域名及转发规则所清洗的攻击带宽 | POST | /soc/api/report/QueryMitigatedBandwidth |
| Reportbandwidthlowdelayp2pservice | 该接口用于查询指定时间段内域名的分钟级别低时延P2P带宽数据。用户需提供查询时间范围和域名来获取带宽信息。返回内容为每分钟的详细P2P带宽数据。有助于用户实时监控域名P2P带宽。<br> | POST | /api/report/bandwidth/low-delay/p2p |
| Reportonlinenumispprovinceservice | 该接口查询直播在不同运营商和省份的在线人数，支持根据请求头Accept-Language返回中文或英文结果。用户需提供域名、时间范围、运营商和省份信息，返回数据包括各时间片的在线人数。此接口适用于网络管理和数据分析，帮助企业了解不同地域和运营商的用户动态，以优化网络资源和性能。 | POST | /api/report/online-num/isp-province/detail |
| Reportdomainstreamdurationservice | 查询多域名多流名的推流时长 | POST | /api/report/domain/stream/duration |
| Channelvalue | 查询客户每个频道的计费值情况，输出频道、计费方法、加速类型、计费值、计费单位、峰值。 | POST | /myview/Channelvalue |
| Reportdomainstreamhlsonlineservice | 该接口用于查询指定时间段内域名及其流名下的在线人数。用户输入开始时间、结束时间、域名和流名即可获取详细信息。返回内容包括每个流的在线人数、流数量以及总在线人数。这对于实时了解在线直播平台的观看情况、流媒体服务的用户参与度等具有重要的实用价值。 | POST | /api/report/hls/online |
| Reportdomainuseragentservice | 该接口用于查询域名在不同时间段内的不同类型的UA流量和请求数。用户可以指定查询的时间范围、域名列表和查询类型（如浏览器、移动设备品牌或操作系统）来获取相关数据。接口返回的数据包括每个UA的流量（以MB为单位）和请求数。该接口有助于用户分析网站在不同设备上的使用情况，优化资源分配。 | POST | /api/report/domain/useragent |
| HttpTest | HTTP 探测 | POST | /wstApi/tools/http-test |
| DnsTest | DNS 探测 | POST | /wstApi/tools/dns-test |
| MtrTest | MTR 探测 | POST | /wstApi/tools/mtr-test |