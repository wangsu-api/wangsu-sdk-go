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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/basicmonitor"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &basicmonitor.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &basicmonitor.{ActionName}Response{}
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
| Reportflowdomainispprovinceiaasservice | This interface is used to query the traffic data of edge servers in various ISPs and provinces, and supports returning Chinese or English data according to the request header Accept-Language. Users need to provide the query time range, domain name, ISP and province information, and the returned content includes the traffic details of each ISP and province, in MB, and displayed at a time granularity of 5 minutes. It is suitable for users who need to analyze the traffic distribution of different regions and operators to help optimize network resources and improve service efficiency. | POST | /api/report/flow/domain-isp-province/iaas |
| Vmpqueryservermetric | Query the CPU usage, memory usage, and bandwidth usage of the virtual machine instance in order to access its own monitoring system and control the operation of the virtual machine. Provide monitoring data queries for nearly 90 days, with a single query range of no more than 3 days and a data granularity of 5 minutes. If it is a bare metal instance, currently only bandwidth query is supported, and CPU and memory query is not supported. | GET | /vmp/servers/metric |
| Lechqueryservermetric | Query the CPU usage, memory usage, and bandwidth usage of the virtual machine instance in order to access its own monitoring system and control the operation of the virtual machine. Provide monitoring data queries for nearly 90 days, with a single query range of no more than 3 days and a data granularity of 5 minutes. If it is a bare metal instance, currently only bandwidth query is supported, and CPU and memory query is not supported. | GET | /lech/servers/metric |