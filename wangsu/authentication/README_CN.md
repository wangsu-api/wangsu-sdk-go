# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/authentication
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/authentication"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &authentication.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := authentication.{ActionName}Response{}
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
| Createuser | 创建一个本地用户 | POST | /api/securelink/idaas/user |
| Deletegroup | 删除一个本地用户组信息，删除组及组内的用户与子组。 | DELETE | /api/securelink/idaas/usergroup |
| Creategroup | 创建一个本地用户组 | POST | /api/securelink/idaas/usergroup |
| Modifygroup | 修改一个本地用户组信息 | PUT | /api/securelink/idaas/usergroup |
| Listgroups | 获取指定用户组下的子组列表 | GET | /api/securelink/idaas/usergroup/list |
| Listusers | 查询指定用户组的用户列表 | GET | /api/securelink/idaas/user/list |
| Deleteuser | 删除一个本地用户信息 | DELETE | /api/securelink/idaas/user |
| Modifyuser | 修改一个本地用户的信息 | PUT | /api/securelink/idaas/user |
| Describeuserinfo | 查询一个用户的信息 | GET | /api/securelink/idaas/user |
| Syncauthconfig | 手动同步身份源 | POST | /api/securelink/idaas/authconfig/sync |