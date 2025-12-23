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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportflow"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportflow.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportflow.{ActionName}Response{}
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
| Flowchannel | 查询频道的分钟粒度流量值 | POST | /myview/Flowchannel |
| Wctquery | 查询客户点播转码各种类型和梯度的5分钟粒度的统计信息 | POST | /myview/Wctquery |
| Flowday | 按天展示频道的流量数据，输出日期、峰值时间、带宽峰值、总流量 | POST | /myview/flow-day |
| Flowtype | 查询流量类型各自占比，包括边缘节点流量、回源流量。 | POST | /myview/flow-type |
| Querydirectorybandwidthtrafficunderlivestreamdomain | 该接口用于查询直播推拉流域名下各个目录的流量或带宽统计数据。用户需提供起止时间、域名，可指定目录查询。默认情况下，数据以5分钟为粒度进行统计。返回结果包括每个目录的总流量或带宽峰值，以及时间分片的详细数据。 | POST | /api/report/flow/dir/detail |
| Querystreamtrafficundermalbdomain | 该接口用于查询指定时间段内域名下的直播流流量或带宽数据。用户需提供时间范围、域名和流名，以及数据类型和粒度选项。接口返回每个域名下各流的流量或带宽详情，包括总流量和峰值带宽。有助于用户监控直播流量情况以优化资源配置。不用于计费核对。 | POST | /api/report/flow/stream/detail |
| Reportflowexactdomainservice | 该接口用于查询精确域名的流量数据，特别适用于泛域名对应的具体域名流量分析。用户需提供时间范围和域名获取流量数据。返回结果包含每个精确域名在不同时间片的流量数据。有助于用户了解域名的流量情况，帮助优化域名配置和网络资源。注意数据有3至5小时延迟，建议查询24小时前的数据。 | POST | /api/report/flow/exact-domain |
| Queryoutputtrafficundershieldpop | 该接口的主要功能是查询多个域名的父节点下行流量数据。用户可以通过提供开始时间、结束时间以及域名列表来获得详细的流量数据报告。返回值包括每个域名的总流量，并提供每个时间片的流量值。此接口适用于需要监测和分析CDN父节点的流量概况及变化趋势。接口调用频率勿高于30/5min。 | POST | /api/report/flow/parent-node |
| Querymultiplestreamtrafficandbandwidthunderthedomain | 该接口用于查询指定时间段内多个域名下直播流的流量和带宽信息。用户输入开始时间、结束时间及域名等，即可获取详细数据，包含每个域名下各流的总流量、峰值带宽、各时间点的流量或带宽数据，以及（可选的）时间点在线人数。这适用于监控直播流网络性能和用户交互，通过数据分析优化网络配置和资源分配，提高直播服务质量。 | POST | /api/report/flow-bandwidth/stream/detail |
| Reportflowp2pshareratioservice | 该接口用于查询多个域名的P2P分享率，用户可以通过提供开始时间、结束时间和域名列表来获取每个域名或域名合并统计的分享率数据。数据可选按分钟或五分钟粒度返回。接口返回的数据包括域名的每个时间片对应的分享率。此接口适用于需要监控和分析域名P2P使用效率的场景。 | POST | /api/report/flow/p2p/share-ratio |
| Querydatatransferforalldomainbyte | 查询账号下所有域名的流量汇总，单位字节。 | POST | /api/report/flow/byte |
| Reportflowispprovinceshorttimeservice | 查询多域名访客IP归属在各ISP各省份流量及带宽，仅支持短时间跨度的查询。 | POST | /api/report/flow/isp-province/shorttime |
| Queryedgebytehitratioservice | 查询多域名分钟级别的在边缘节点命中缓存的字节命中率 | POST | /api/report/flow/edge-hit-ratio/total |
| Reportupflowdomaincountryservice | 该接口查询指定域名在各国家和地区的访客上行流量带宽分布。用户输入时间范围和域名列表，可按域名、国家或国内外查询数据。返回结果包括各地区的上行流量总和及其百分比，以及每个时间片段的上行流量和带宽，帮助分析和管理全球范围内的网站流量带宽。建议查询24小时前的数据。 | POST | /api/report/up-flow/domain-country |
| Gettotaltrafficforalldomains | 该接口用于查询用户账号下所有域名的流量汇总信息。用户提供开始结束时间参数来获取详细的流量数据。通过该接口，用户可了解和分析所有域名在指定时间段内的CDN流量使用情况。 | GET | /api/report/traffic |
| Getbandwidthandtrafficbydirectoriesmainlandchinaonly | 该接口用于查询多个域名下多级目录的带宽和总流量，仅限于中国大陆加速区域。用户需提供查询时间范围和目录层级信息。返回结果包括每个域名的目录总流量和带宽峰值。此接口有助于用户分析网站流量分布和带宽使用情况，帮助优化资源配置和提升服务性能。 | POST | /api/report/traffic/dir/info |
| Gettrafficandbandwidthbyclientcountry | 该接口查询指定域名在各国家和地区的访客流量带宽分布。用户输入时间范围和域名列表，可按域名、国家或国内外查询数据。返回结果包括各地区的流量总和及其百分比，以及每个时间片段的流量和带宽，帮助分析和管理全球范围内的网站流量带宽。建议查询24小时前的数据。 | POST | /api/report/traffic/domain-country |
| Gettrafficandrequestshitratiobyispprovince1mingranularity | 该接口用于查询特定省份和ISP的分钟级流量和请求数命中率，用户通过提供开始和结束时间、域名、省份及运营商等信息，获得每分钟的流量或请求命中率数据。该接口适用于需要分析不同省份和运营商在指定时间段内的网络性能的场景。用户可以通过此接口监控流量和请求的命中情况，并优化其网络资源配置。 | POST | /api/report/traffic/isp-province/hit-rate/total |
| Gettrafficbyispprovince1mingranularity | 该接口用于查询省份和运营商在指定时间段内的每分钟流量汇总。用户可以通过指定多个域名、时间区间、ISP等条件来获取详细的数据汇总。接口返回数据包括流量、带宽和请求数量。接口适用于在指定省份和ISP中监控网站的流量和带宽使用情况，以便于优化网络资源或识别访问趋势。 | POST | /api/report/traffic/isp-province/total |
| Gettrafficrequestshitratiobyclientispprovince1mingranularity | 该接口用于查询省份和ISP维度的分钟级别流量和请求数命中率，支持根据请求头Accept-Language返回中文或英文信息。用户提供时间范围、域名和ISP、省份信息。返回结果包含每个省份ISP的详细流量或请求数命中率。有助于用户分析各地区网络性能，从而优化内容分发策略和提高网络服务质量。<br> | POST | /api/report/traffic/isp-province/hit-rate/detail |
| Gettrafficbybillingregionformultidomains | 该接口用于获取加速域名在指定结算区域的流量汇总统计数据。用户需提供查询的时间范围、区域代码和可选的域名和数据粒度。返回结果包括总流量汇总和各时间段的流量数据。此接口适合需要分析流量分布和优化网络资源配置的用户，帮助进行结算区域流量管理和成本控制。 | POST | /api/report/traffic/area |
| Gettraffichitratiobyispprovinceformultidomains | 该接口用于查询多个域名在不同运营商和省份中的字节命中率。用户可以通过指定查询时间范围、域名列表、运营商和省份等参数来获取详细的数据分布。数据包括每个时刻的字节命中率和相关的命中流量信息。该接口适用于需要监控和分析不同区域和运营商下缓存性能的用户，可以帮助优化网络资源分配和提升用户体验。 | POST | /api/report/traffic/hit-rate/isp-province |
| Gettrafficandbandwidthbyispprovinceformultidomains | 该接口用于查询多个域名在各运营商和省份的流量和带宽数据。用户需提供查询时间范围和域名列表，可选择按ISP和省份进行分组。返回内容包括每个域名在各省份各运营商的流量和带宽数据。此接口有助于用户分析不同地区和运营商下网络性能，帮助优化资源配置和提升服务质量。<br> | POST | /api/report/traffic/isp-province |
| Gettrafficbyispprovinceformultidomainsbytes | 该接口允许用户查询多域名服务节点所属的各ISP和省份的流量数据（单位：Byte）。用户需提供查询时间范围和域名列表，可按ISP和省份分组。结果包含每个域名的流量和带宽详细信息，适用于分析各地区和运营商的网络性能，从而优化资源和提升服务质量。 | POST | /api/report/traffic/isp-province/byte |
| Getorigintrafficbyispprovince | 该接口查询多个域名在各省份和运营商的回源流量。用户可按域名、运营商、省份分组查询，并指定时间范围和数据粒度。返回结果提供详细的回源流量和带宽信息，帮助分析区域和运营商的回源性能。 | POST | /api/report/traffic/origin/isp-province |
| Gettrafficbyprotocol | 该接口用于查询多个域名在指定传输协议下的流量数据，查询的是所有边缘节点的数据。用户需指定域名、起止时间、传输协议等参数，并可选择数据粒度和分组维度。默认情况下，查询会返回HTTPS协议的流量数据，输出结果为每个域名在不同时间点的流量值。此接口适用于需要查询和分析多域名不同传输协议流量的场景。 | POST | /api/report/traffic/protocol |
| Gettrafficrequestsandpeakbandwidthformultidomains | 该接口用于查询多个域名的流量请求数和峰值带宽。用户需提供时间范围和域名列表，可获取各域名的总流量、总请求数、峰值带宽及其发生时间，以及每个时间段的详细流量与请求数。有助于用户了解网站的访问情况，从而有效调整策略以满足不同流量需求。 | POST | /api/report/traffic-request |
| Querytotaltrafficformultidomains | 此接口用于获取多个域名的流量汇总信息。用户需提供时间范围和域名，数据粒度可选五分钟、一小时或一天。返回内容括总流量和各时间节点的流量值。此接口有助于用户监控和分析多个域名流量情况，从而掌握整体流量趋势及变化。 | POST | /api/report/domaintraffic |
| Getorigintrafficandrequestsformultidomains | 该接口旨在查询多个域名的回源流量和请求数，用户可以通过输入开始时间、结束时间、指定的域名等信息来获取这些数据。接口会返回每个域名的总流量、总请求数、请求和带宽峰值及其对应的时间点，并支持以分钟粒度提供详细流量、请求和带宽数据。这一功能有助于用户监控和分析网站回源流量。 | POST | /api/report/traffic-request/origin |
| Getbandwidthandrequestsbyipversionformultidomains | 该接口用于查询多个域名的IPv6和IPv4网络带宽或请求数，用户可以通过指定开始时间、结束时间、域名列表来获取数据。接口返回每个域名的详细统计信息，包括指定时间段内的IPv4和IPv6的带宽或请求数。该接口适用于用户需要监控多个域名流量请求情况，通过分析返回的数据，可以优化网络资源分配和提高网站性能。 | POST | /api/report/traffic-request/multi-ip-version |
| Getbandwidthsavingratio | 该接口用于查询指定时间段内的带宽节省量，用户需提供开始和结束时间，域名信息获取数据。返回的数据包括总节省带宽平均值，按时间片对应的实际带宽节省量。该接口特别适用于需要监控分析CDN缓存命中带来的带宽节省量的场景。 | POST | /api/report/bandwidth/saving-bandwidth/total |
| Getcdnorigintraffic | 用于查询指定维度的cdn回源站流量数据，用户可以通过该接口来查询对应客户的详细频道回源站流量报表，包括输出日期、峰值时间、带宽峰值、总流量等。这对客户了解自身的回源站流量使用情况有很直接的帮助。 | POST | /cdn/traffic/origin |