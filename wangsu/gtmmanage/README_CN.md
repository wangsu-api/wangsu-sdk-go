# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/gtmmanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/gtmmanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &gtmmanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := gtmmanage.{ActionName}Response{}
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
| Deldispatchpolicy | 删除调度策略是将域名下的策略配置全部删除。 | POST | /clouddns/Deldispatchpolicy |
| Controldispatchpolicy | 对调度策略进行控制，启用调度策略即下发策略配置并部署生效；停用调度策略时，策略配置不删除，但配置不生效。 | POST | /clouddns/Controldispatchpolicy |
| Querydispatchpolicydetail | 根据ID查询调度详情，查询条件：输入域名ID、策略ID | POST | /clouddns/Querydispatchpolicydetail |
| Savedispatchpolicy | 为域名添加新的调度策略，调度策略分成2种类型：负载均衡；主备+负载均衡。<br>新增、修改调度策略为同一个接口，但json格式部分不一样。 | POST | /clouddns/Savedispatchpolicy |
| Querydispatchpolicies | 分页查询策略信息，包括策略类型、线路信息、策略状态（解析至主源还是备源）、监控配置。 | POST | /clouddns/Querydispatchpolicies |
| Controldispatchresource | 启停调度资源 | POST | /clouddns/Controldispatchresource |
| Controlresourcecluster | 调度策略启用/停用主源 | POST | /clouddns/Controlresourcecluster |