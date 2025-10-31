# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/strategymanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/strategymanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &strategymanagement.{ActionName}Request{}

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

## Error Handling

Always check for errors returned by the API calls:

```go
_, err := auth.Invoke(config, request, response)
if err != nil {
    log.Printf("error: %s\n", err)
    // Handle the error appropriately
    return
}
```

## API List
For detailed API documentation and available methods, please refer to the [official Wangsu API documentation](https://www.wangsu.com/document/api-doc/Overview?productType=all).

| ActionName | enDescription | client_methods | uri |
| --- | --- | --- | --- |
| Listpropagationpolicies | query list of propagation | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies |
| Createpropagationpolicies | create propagation | POST | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies |
| Deletepropagationpolicies | delete propagation | DELETE | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/* |
| Getpropagationpolicies | query details of propagationpolicy | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/* |
| Updatepropagationpolicies | update propagation | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/* |
| Putpatchpropagationpolicies | partial update propagationpolicy | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/propagationpolicies/*/ws/patch |
| Updatehorizontalpodautoscaler | update hpa | PUT | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Patchhorizontalpodautoscaler | patch hpa | PATCH | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Deletehorizontalpodautoscaler | delete hpa | DELETE | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Listhorizontalpodautoscaler | list  hpa | GET | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers |
| Gethorizontalpodautoscaler | get hpa | GET | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers/* |
| Createhorizontalpodautoscaler | create  hpa | POST | /apis/autoscaling/v2beta2/namespaces/*/horizontalpodautoscalers |
| Listoverridepolicy | list Overridepolicies | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies |
| Createoverridepolicy | Create Overridepolicyes | POST | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies |
| Getoverridepolicy | Get OverridePolicy | GET | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/* |
| Putoverridepolicy | Update Overridepolicy | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/* |
| Deleteoverridepolicy | Delete Overridepolicy | DELETE | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/* |
| Wspatchoverridepolicy | Patch Overridepolicy | PUT | /apis/policy.karmada.io/v1alpha1/namespaces/*/overridepolicies/*/ws/patch |