# Wangsu SDK for Go

文档提供了使用 Wangsu SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## 单独安装

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/basicpermission
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/basicpermission"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &basicpermission.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := basicpermission.{ActionName}Response{}
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
| Addterminalauth | 新增基础权限 | POST | /api/securelink/idaas/terminalauth |
| Updateterminalauth | 编辑基础权限. | PUT | /api/securelink/idaas/terminalauth |
| Deleteterminalauth | 根据ID删除基础权限 | DELETE | /api/securelink/idaas/terminalauth |
| Addterminalauthuserorgroup | 新增基础权限关联的用户/用户组 | POST | /api/securelink/idaas/terminalauth/user-group |
| Queryterminalauthlist | 获取基础权限列表 | GET | /api/securelink/idaas/terminalauth |
| Queryterminalauthinfo | 查询基础权限信息 | GET | /api/securelink/idaas/terminalauth/query |
| Removeterminalauthresource | 移除基础权限关联的应用 | DELETE | /api/securelink/idaas/terminalauth/related-resources |
| Removeterminalauthuserorgroup | 移除基础权限关联的用户/用户组 | DELETE | /api/securelink/idaas/terminalauth/user-group |
| Queryterminalauthresourcelist | 获取基础权限可关联的应用列表 | GET | /api/securelink/idaas/terminalauth/resources |
| Updateterminalauthrelatedresources | 修改基础权限关联的应用 | PUT | /api/securelink/idaas/terminalauth/related-resources |
| Queryterminalusergroupbyauthconfig | 根据身份源名称获取用户组列表. | GET | /api/securelink/idaas/usergroup/list-by-authconfig |
| Associaterightsgroupstouserorusergroup | 该接口用来直接为指定用户/用户组关联权限组。如果有多个用户/用户需要分配的多个一样的应用权限时，管理员可以先把多个应用关联到同一个权限组后，再使用该接口分配给用户/用户组。 | POST | /api/securelink/idaas/terminalauth/user-bind |