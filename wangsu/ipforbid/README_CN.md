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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/ipforbid"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &ipforbid.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &ipforbid.{ActionName}Response{}
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
| Forbidorresumevisitoripsbydomainservice | 对访问指定域名的IP，进行封禁解禁操作。 | POST | /api/spider/ip-forbid |
| Queryforbiddingvisitoripsbydomainservice | 查询域名粒度的封禁IP信息，支持分页查询。 | POST | /api/spider/ip-forbid/query |
| Queryforbiddingvisitoripsbylabelcodeservice | 查询标签粒度的封禁IP信息，支持分页查询。 | POST | /api/spider/label-ip-forbid/query |
| Forbidorresumevisitoripsbylabelcodeservice | 一个客户可以创建一个标签，该标签可以关联若干域名，通过该标签进行封禁解禁访问IP的操作，效果等同于对该标签关联的所有域名，进行封禁解禁指定的访客IP。 | POST | /api/spider/label-ip-forbid/operate |
| Addorremoveforbiddingipwhitelistservice | IP白名单功能支持三个粒度级别：<br>（1）客户粒度<br>范围：对该客户所有域名和标签生效<br>作用：过滤该客户下所有"标签粒度"和"域名粒度"的封禁请求<br>（2）域名粒度<br>范围：仅对指定域名生效<br>作用：仅过滤该域名的"域名粒度"封禁请求<br>（3）标签粒度<br>范围：仅对指定标签生效<br>作用：仅过滤该标签的"标签粒度"封禁请求<br><br>说明：新增白名单时，可选择自动解禁对应粒度下已被封禁的IP。 | POST | /api/spider/ip-whitelist/operate |
| Queryforbiddingipwhitelistservice | 提供查询封禁IP白名单列表的功能。 | POST | /api/spider/ip-whitelist/query |