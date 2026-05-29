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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/loadbalancer"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &loadbalancer.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &loadbalancer.{ActionName}Response{}
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
| Createloadbalancerlistener | 该接口用于为负载均衡实例创建监听。用户需指定监听协议、端口、所属负载均衡ID、服务器池ID、负载均衡算法，并根据需要开启健康检查及会话保持功能，同时可配置详细的健康检查策略（如检查协议、端口、路径、超时时间等）和会话保持超时时间。成功创建后，接口将返回新创建监听的ID。 | POST | /vmp/slb/listeners |
| Queryloadbalancerlisteners | 该接口用于查询负载均衡的监听列表。用户可根据监听ID、负载均衡ID、监听状态、服务器池ID等条件进行筛选。接口返回监听的详细信息，包括健康检查配置、会话保持设置等。 | GET | /vmp/slb/listeners |
| Updateloadbalancerlistener | 该接口用于更新指定负载均衡监听的配置。通过`listenerId`定位监听器，用户可以修改其协议、端口、名称、会话保持设置、服务器池ID、负载均衡算法，以及健康检查的详细配置，包括是否开启健康检查、健康检查的协议、端口、路径、超时时间、间隔、重试次数和正常状态码等。 | PUT | /vmp/slb/listeners/* |
| Updateloadbalancerserverpool | 该接口用于更新一个负载均衡服务器池。用户需通过 `poolId` 指定要更新的服务器池。可以更新服务器池的成员（members，包含RS实例ID、权重、服务端口）、服务器池名称以及备注信息。 | PUT | /vmp/slb/serverpools/* |
| Createloadbalancerserverpool | 该接口用于创建一个负载均衡服务器池。用户需在请求体中指定负载均衡ID (`lbId`)、服务器池成员 (`members`)、服务器池名称 (`name`) 和可选的备注信息 (`remark`)。服务器池成员可包含RS实例ID (`instanceId`)、权重 (`weight`) 和服务端口 (`servicePort`)。接口成功后，将返回新创建服务器池的详细信息，包括其ID (`id`) 以及包含节点名称、VIP信息、网络类型、IP协议、VIP地址、创建/修改时间、负载均衡名称、状态等具体配置的`data`对象。 | POST | /vmp/slb/serverpools |
| Createloadbalancer | 该接口用于创建一个负载均衡对象。用户需要提供节点名称和实例规格，并可选择性指定负载均衡的线路运营商、网络类型、负载均衡名称和IP协议。成功创建后，接口将返回负载均衡的详细信息，包括节点名称、VIP信息、网络类型、IP协议、VIP地址、创建和修改时间、负载均衡名称、状态描述、ID和规格类型等。 | POST | /vmp/slb/loadbalancers |
| Queryloadbalancer | 该接口用于查询负载均衡列表。用户可以通过指定节点名称、负载均衡ID、状态和规格等参数进行筛选查询，所有参数均支持多值查询。接口将返回符合条件的负载均衡详细信息，包括其节点名称、VIP信息、网络类型、IP协议、VIP地址、修改时间、创建时间、名称、状态描述、ID、状态和规格类型等。 | GET | /vmp/slb/loadbalancers |
| Deleteloadbalancerserverpools | 接口用于删除指定的负载均衡服务器池。用户需提供待删除负载均衡服务器池的ID，支持通过逗号分隔提供多个ID进行批量删除。操作成功后，系统会返回删除成功的提示信息。 | DELETE | /vmp/slb/serverpools/* |
| Queryloadbalancerserverpools | 该接口用于根据指定的条件（如负载均衡服务器池ID或负载均衡ID）查询负载均衡服务器池的列表。用户可通过传入ID或lbId参数进行筛选。返回结果包含每个服务器池的详细信息，包括其关联的RS实例列表。 | GET | /vmp/slb/serverpools |
| Deleteloadbalancerlisteners | 接口用于删除指定的负载均衡监听。用户需提供待删除负载均衡监听的ID，支持通过逗号分隔提供多个ID进行批量删除。操作成功后，系统会返回删除成功的提示信息。 | DELETE | /vmp/slb/listeners/* |
| Startloadbalancerlistener | 该接口用于启用指定的负载均衡监听。用户需要提供一个或多个负载均衡监听的ID列表，成功后接口将返回启用结果。 | PUT | /vmp/slb/listeners/action/start |
| Enableloadbalancer | 该接口用于启用指定的负载均衡实例。用户需要提供一个或多个负载均衡的ID列表，成功后接口将返回启用结果。 | PUT | /vmp/slb/loadbalancers/action/start |
| Disableloadbalancer | 该接口用于停用指定的负载均衡实例。用户需要提供一个或多个负载均衡的ID列表，成功后接口将返回停用结果。 | PUT | /vmp/slb/loadbalancers/action/stop |
| Stoploadbalancerlistener | 该接口用于停用指定的负载均衡监听。用户需要提供一个或多个负载均衡监听的ID列表，成功后接口将返回停用结果。 | PUT | /vmp/slb/listeners/action/stop |
| Deleteloadbalancer | 该接口用于删除指定的负载均衡实例。用户需提供待删除负载均衡的ID，接口将返回操作结果。 | DELETE | /vmp/slb/loadbalancers/* |