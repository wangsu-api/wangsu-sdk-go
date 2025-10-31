# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/storagemanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/storagemanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &storagemanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := storagemanagement.{ActionName}Response{}
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
| Listpvcs | query list of pvc | GET | /api/v1/namespaces/*/persistentvolumeclaims |
| Createpvcs | create pvc | POST | /api/v1/namespaces/*/persistentvolumeclaims |
| Getpvcs | query details of pvc | GET | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Updatepvcs | update pvc | PUT | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Deletepvcs | delete pvc | DELETE | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Putpatchpvcs | partial update pvc | PUT | /api/v1/namespaces/*/persistentvolumeclaims/*/ws/patch |
| Pagingpvcs | Get pvc paging list | GET | /openapi/custom/api/v1/persistentvolumeclaims |
| Pvcinnamespace | Get the pvc under the namespace | GET | /openapi/custom/api/v1/namespaces/*/persistentvolumeclaims |
| Deletepvcfromedge | Remove pvc directly from edge cluster | DELETE | /openapi/custom/api/v1/namespaces/*/persistentvolumeclaims/* |
| Liststorageclass | Get the storageClass list | GET | /openapi/custom/api/v1/storageclasses |
| Querychannelrecordfiles | Query channel recording files | POST | /ivcs/report/origin/query-channel-record |