# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/edgehostname
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgehostname"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &edgehostname.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgehostname.{ActionName}Response{}
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
| Deleteedgehostname | This API is used to delete a specified Edge Hostname. If the Edge Hostname has never been deployed, calling the API directly deletes the record. If the Edge Hostname is in a deployed state, calling this API will trigger the removal of the deployment and return the deployment task ID. | DELETE | /api/edge-hostnames/* |
| Deployakcdnedgehostname | This API is used to deploy the specified Edge Hostname and returns the deployment task ID. | POST | /api/edge-hostnames/*/deploy |
| Disableedgehostname | This API is used to disable a specified edge hostname. Only edge hostnames that are active are eligible for disabling. Disabling will trigger the deletion of the edge hostname deployment but will not remove the edge hostname record. The API returns the deployment task ID. | POST | /api/edge-hostnames/*/disable |
| Enableedgehostname | This API is used to enable a specific Edge Hostname. Only Edge Hostnames that are in a disabled state can be enabled. Enabling it will trigger an Edge Hostname deployment, and the API will return the deployment task ID. | POST | /api/edge-hostnames/*/enable |
| Queryedgehostname | This API is used to enable a specific Edge Hostname. Only Edge Hostnames that are in a disabled state can be enabled. Enabling it will trigger an Edge Hostname deployment, and the API will return the deployment task ID. | GET | /api/edge-hostnames/* |
| Queryedgehostnames | This API is used to query the list of Edge Hostnames, returning all Edge Hostnames under the account along with their CNAME status and deployment status, with support for pagination. | GET | /api/edge-hostnames |
| Updateedgehostname | This interface is used to modify the basic configuration of the Edge Hostname, enabling self-service control over the routing effects, including remarks, acceleration region restrictions, and region configuration. | PUT | /api/edge-hostnames/* |