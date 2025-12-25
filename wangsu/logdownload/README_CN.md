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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/logdownload"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &logdownload.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := logdownload.{ActionName}Response{}
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
| Querydomainlogdownloadaddress | 查询多个域名的日志下载地址。日志文件粒度默认为24小时，返回的数据以实际配置为准。 | POST | /api/report/log/downloadLink |
| Querytranscodingdurationlogdownloadaddress | 该接口用于提供多域名的转码时长日志文件的下载地址。用户通过指定时间范围和域名列表，可以获取每个域名相关的转码日志信息，返回内容包括日志文件的开始和结束时间、下载地址、文件名、文件大小及下载地址的过期时间。此接口有助于用户精确掌握各域名在指定时间内的转码详情并做出相应优化。 | POST | /api/report/log/download-file/transcoding |
| Querycataloguetrafficandbroadcastcountdownload | 该接口用于查询多个域名在指定时间范围内的目录流量统计及播放次数统计数据的下载地址。用户需指定域名和起止时间。如果未指定域名，默认查询账号下所有域名（受数量限制）。响应将返回日志文件下载地址、文件大小、过期时间等信息。 | POST | /api/report/catalogue/broadcast/download-file |