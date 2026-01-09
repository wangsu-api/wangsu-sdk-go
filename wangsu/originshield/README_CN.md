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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/originshield"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &originshield.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &originshield.{ActionName}Response{}
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
| Queryshieldroutemaps | 该接口用于查询源站防护盾可用的回源路由列表。用户创建源站防护盾前需调用该接口查询并选择可用的回源路由。<br> | GET | /api/origin_shields/maps |
| Queryoriginshield | 该接口用于查询源站防护盾。用户指定源站防护盾ID查询源站防护盾。 | GET | /api/origin_shields/* |
| Disableoriginshield | 该接口用于禁用源站防护盾。用户指定源站防护盾ID禁用源站防护盾，禁用前需确认源站防护盾没有被已部署的项目引用。 | POST | /api/origin_shields/*/disable |
| Enableoriginshield | 该接口用于启用源站防护盾。用户指定源站防护盾ID启用源站防护盾，启用前需确认源站防护盾编码关联的所有IP段都已加入白名单中。 | POST | /api/origin_shields/*/enable |
| Deleteoriginshield | 该接口用于删除源站防护盾。用户指定源站防护盾ID删除源站防护盾。 | DELETE | /api/origin_shields/* |
| Updateoriginshield | 该接口用于修改源站防护盾。用户指定源站防护盾名称与描述来修改源站防护盾。 | PUT | /api/origin_shields/* |
| Createoriginshield | 该接口用于创建源站防护盾。用户选择可用的回源路由并指定源站防护盾名称创建源站防护盾。<br> | POST | /api/origin_shields |
| Confirmoriginshieldipsegmentchange | 该接口用于确认源站防护盾IP段变更。用户指定源站防护盾下IP段记录ID确认源站防护盾下IP段变更。 | POST | /api/origin_shields/*/ip_segments/confirm |
| Queryoriginshields | 该接口用于查询源站防护盾列表。用户可以指定源站防护盾ID、源站防护盾编码、源站防护盾名称进行过滤查询。 | GET | /api/origin_shields |
| Queryavailableoriginshields | 该接口用于查询可用的源站防护盾列表。用户需指定商品或者项目ID进行查询。 | GET | /api/origin_shields/available |
| Queryoriginshieldpendingipsegments | 该接口用于查询源站防护盾待确认IP段列表。用户可以指定源站防护盾ID、操作、IP类型进行过滤查询。 | GET | /api/origin_shields/*/ip_segments/pending |
| Queryoriginshieldactiveipsegments | 该接口用于查询源站防护盾已生效IP段列表。用户需指定源站防护盾ID进行查询。<br><br> | GET | /api/origin_shields/*/ip_segments/active |
| Queryoriginshieldrelatedinfos | 该接口用于查询源站防护盾关联的源站与域名信息。用户需指定源站防护盾ID进行查询。 | GET | /api/origin_shields/*/related_infos |