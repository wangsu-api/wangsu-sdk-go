# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```


## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/configcenter"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &configcenter.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := configcenter.{ActionName}Response{}
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
| Createconfigmap | create configmap | POST | /api/v1/namespaces/*/configmaps |
| Listconfigmap | query list of configmap | GET | /api/v1/namespaces/*/configmaps |
| Getconfigmap | query details of configmap | GET | /api/v1/namespaces/*/configmaps/* |
| Updateconfigmap | update configmap | PUT | /api/v1/namespaces/*/configmaps/* |
| Deleteconfigmap | delete configmap | DELETE | /api/v1/namespaces/*/configmaps/* |
| Listsecret | query list of secret | GET | /api/v1/namespaces/*/secrets |
| Getsecret | query details of secret | GET | /api/v1/namespaces/*/secrets/* |
| Createsecret | create secret | POST | /api/v1/namespaces/*/secrets |
| Updatesecret | update secret | PUT | /api/v1/namespaces/*/secrets/* |
| Deletesecret | delete secret | DELETE | /api/v1/namespaces/*/secrets/* |
| Putpatchconfigmap | partial update configmap | PUT | /api/v1/namespaces/*/configmaps/*/ws/patch |
| Putpatchsecret | partial update secret | PUT | /api/v1/namespaces/*/secrets/*/ws/patch |
| Gettoken | get authentication token | GET | /openapi/custom/api/v1/token |