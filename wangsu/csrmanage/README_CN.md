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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/csrmanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &csrmanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := csrmanage.{ActionName}Response{}
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
| CreateTheCsr | 创建CSR | POST | /api/csr-manage/csr |
| DeleteCsrRecord | 删除csr记录接口 | DELETE | /api/csr-manage/csr/* |
| QueryCsrService | 该接口用于根据CSR ID获取单个CSR的详细信息。用户需提供一个CSR ID作为路径参数，系统将返回该CSR的创建时间、修改时间、主域名、备份域名列表、密钥算法、密钥强度、CSR内容等详细属性。 | GET | /api/csr-manage/csr/* |
| QueryCsrList | 查询csr列表 | GET | /api/csr-manage/list |
| Updatecsr | 该接口用于更新指定的CSR（Certificate Signing Request）记录。用户需提供CSR的名称，并可选择性添加备注信息。接口响应成功时会返回响应代码和响应信息。 | PUT | /api/csr-manage/update/csr/* |