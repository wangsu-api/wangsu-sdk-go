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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/toolservice"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &toolservice.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := toolservice.{ActionName}Response{}
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
| Bandwidthlimitservice | 设置/取消指定域名的带宽限制。 | POST | /api/tools/setBandwidthLimit |
| Querybandwidthlimittasklistservice | 此接口用于查询账号下的带宽限制任务列表，返回所有有带宽的任务详细信息。返回内容包括域名、任务名称以及设置的最大带宽值。该接口适用于需要评估和管理流量控制策略的场景，帮助用户快速识别并管理当前设置的带宽限制任务。 | POST | /api/tools/queryBandwidthLimitTaskList |
| Icpqueryservice | 查询所指定的域名是否在中国大陆工信部进行备案。 | GET | /api/icp |
| Akamaiipforbiddenservice | 该接口用于封禁特定IP地址对AKAMAI服务的访问，用户可通过输入要封禁的IP列表来实现操作。此接口有助于用户阻挡不可信或有害IP地址的访问，提升网络安全性。<br> | POST | /api/tools/ip-forbid/akamai |
| Reportserveripcountrycodeservice | 该接口用于查询特定域名在不同国家的CDN服务IP列表。用户提供域名以及国家来获取信息。返回的数据包括每个域名在指定国家覆盖节点的IP列表。适用于用户有效查看CDN在全球范围内的覆盖和服务节点分布。 | POST | /api/report/service-ip/country |
| Akamaiippermitservice | 解封AKAMAI访问IP | POST | /api/tools/ip-permit/akamai |
| Supplyregisterservice | 第三方设备注册接口<br> | POST | /sr/supply/register |
| Supplyincludeservice | 第三方设备纳管申请接口<br> | POST | /sr/supply/include |
| Ipdomainservice | 该接口用于根据IP地址查询正在使用该IP的域名。用户输入IP地址来获取与该IP相关联的域名列表。接口返回的信息包含IP当前使用状态，以及使用该IP的域名列表。在实际应用中，此接口可帮助用户检测特定IP的域名使用情况，适用于网络监控和管理。 | POST | /api/tools/ip/domain-list |
| Queryallbandwidthlimittasklistservice | 该接口用于查询用户账号下配置的所有带宽限制任务。用户在调用时可以选择是否包含任务所涉及的所有客户域名以及决定是否返回所有任务状态的信息。返回的数据以清单方式展示每个带宽限制任务的详细信息，包括任务名称、类别、状态以及相关控制策略和参数。此接口有助于用户管理带宽设置，可以及时对特定的流量和请求进行有效的控制和处理。 | POST | /api/tools/queryAllBandwidthLimitTaskList |
| Queryconversiontaskdetail | 该接口用于查询指定转换任务的详细信息，包括任务ID、配置文件类型、转换结果、文件名称、创建时间、域名列表、任务状态等。用户需要提供任务ID和配置文件类型作为入参。 | POST | /api/v1/akamai/get-task-detail |
| Queryconversiontasklist | 该接口用于查询转换任务的列表。用户可以根据域名、任务ID、起始时间、结束时间等条件进行筛选。接口返回转换任务的详细信息，包括任务ID、创建时间、类型和状态等。 | POST | /api/v1/akamai/wplus/get-task-list |
| Createwebsecurityconfigurationtask | 该接口用于创建一个Web安全配置任务。用户可以通过上传文件名和base64加密的文件内容和提供域名列表来创建任务。任务创建后将返回任务ID，用于后续查询任务状态。 | POST | /api/v1/akamai/wplus/domain/move |
| Publishconverteddomain | 该接口用于发布通过安全配置转换后的域名。用户可以通过传入域名列表，将这些域名进行发布操作。接口会返回相应的响应信息和状态码，指示操作是否成功。 | POST | /api/v1/akamai/domain/complete |
| Createclientnetworklisttask | 该接口用于创建客户端或网络列表任务。用户需指定任务类型、列表内容等参数。成功创建后，接口将返回任务ID及其他相关信息。 | POST | /api/v1/akamai/wplus/client-list/move |
| Querydeviceoperationalinfoservice | 该接口用于支持各种查询信息，目前实现了查询设备的运营信息，特别是其是否符合质检条件。用户可以通过提供设备序列号（SN）和查询类型（例如，查询可质检状态）来获取设备的状态。响应将返回查询结果代码，指示设备是否存在、是否可质检、SN是否冲突，以及设备序列号和相关信息。 | POST | /sr/supply/query |