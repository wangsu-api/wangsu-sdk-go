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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportiplist"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportiplist.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportiplist.{ActionName}Response{}
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
| Reportserveripispprovinceservice | Query the service IP list of CDN in each ISP and each province | POST | /api/report/server-ip/isp-province |
| Reportserveripexistflowservice | This interface can obtain the corresponding CDN service IP list with Traffic by providing the CDN domain . This function is suitable for scenarios where you need to understand the actual accelerated use of the domain service IP. | POST | /api/report/server-list/exist-flow |
| Querycdniplist | This interface is used to query the coverage node IP list corresponding to each specified domain name in the CDN service. Users can obtain the node IP information of multiple domain names under their account by submitting relevant requests, which is mainly used for traffic scheduling and optimization management. The returned data shows the detailed coverage node IP of each domain name, helping users understand and monitor the node distribution of CDN services. | POST | /api/report/server-list |
| Reportserverlistservice | Query the IP list of a specific domain in the CDN service and its affiliated ISP and region information. Through this interface, users can obtain detailed information about the CDN service node associated with the domain. | GET | /api/report/server-list/ip-isp-area |