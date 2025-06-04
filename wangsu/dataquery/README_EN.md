# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/dataquery
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/dataquery"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &dataquery.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := dataquery.{ActionName}Response{}
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
| Sitetotalflowchart | Query the allocated bandwidth and actual bandwidth trends of the specified site. | GET | /api/siteFlow/siteTotalFlowChart |
| Sitesessioninfo | Query the session logs of the enterprise site. | GET | /api/siteSession/info |
| Sitequalitychart | Query the link quality trend data of the enterprise site. | GET | /api/siteTunnel/qualityChart |
| Sitemsdevice | Query the currently used CPE at a specified enterprise's dual-device site. | GET | /api/siteDevice/ms |
| Sitelogicchart | Site connection quality | GET | /api/siteLogic/qualityChart |
| Allsitesumflow | Query the total traffic of all sites under a specified enterprise. | GET | /api/siteFlow/orgSumFlow |
| Saassession | query site saas-session data. | GET | /api/saasSession/info |