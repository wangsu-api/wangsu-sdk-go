# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/gtmmanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/gtmmanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &gtmmanage.{ActionName}Request{}

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
| Deldispatchpolicy | q | POST | /clouddns/Deldispatchpolicy |
| Controldispatchpolicy | Use to batch enable/disable dispatch policies. Takes one minute to take effect. | POST | /clouddns/Controldispatchpolicy |
| Querydispatchpolicydetail | Used to query the detailed information of dispatch policy. Takes one minute to take effect. | POST | /clouddns/Querydispatchpolicydetail |
| Savedispatchpolicy | Used to add dispatch policy, and the interface used is the same to the one to midify dispatch policy, but the two use different json formats. Takes one minute to take effect. | POST | /clouddns/Savedispatchpolicy |
| Querydispatchpolicies | Used to query dispatch policy information by page. Takes one minute to take effect. | POST | /clouddns/Querydispatchpolicies |
| Controldispatchresource | Start or stop scheduling resources | POST | /clouddns/Controldispatchresource |
| Controlresourcecluster | Scheduling policy enable/disable primary source | POST | /clouddns/Controlresourcecluster |