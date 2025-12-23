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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/strategymanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &strategymanagement.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := strategymanagement.{ActionName}Response{}
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
| Listpropagationpolicies | 获取propagation列表 | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies |
| Createpropagationpolicies | 创建Propagation | POST | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies |
| Deletepropagationpolicies | 删除propagation | DELETE | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/* |
| Getpropagationpolicies | 获取propagation的详细信息 | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/* |
| Updatepropagationpolicies | 更新propagation | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/* |
| Putpatchpropagationpolicies | 部分更新PropagationPolicy | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/*/ws/patch |
| Updatehorizontalpodautoscaler | 修改HPA | PUT | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Patchhorizontalpodautoscaler | 部分修改HPA | PATCH | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Deletehorizontalpodautoscaler | 删除HPA | DELETE | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Listhorizontalpodautoscaler | 查询HPA列表 | GET | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers |
| Gethorizontalpodautoscaler | 查询HPA | GET | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Createhorizontalpodautoscaler | 新增HPA | POST | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers |
| Listoverridepolicy | Overridepolicies列表 | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies |
| Createoverridepolicy | 创建Overridepolicies | POST | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies |
| Getoverridepolicy | 获取OverridePolicy的详细信息 | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/* |
| Putoverridepolicy | 更新OverridePolicy | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/* |
| Deleteoverridepolicy | 删除OverridePolicy | DELETE | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/* |
| Wspatchoverridepolicy | 部分更新OverridePolicy | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/*/ws/patch |