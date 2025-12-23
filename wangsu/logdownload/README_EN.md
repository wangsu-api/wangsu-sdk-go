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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/logdownload"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &logdownload.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := logdownload.{ActionName}Response{}
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
| Querydomainlogdownloadaddress | Report Log Multi-Domain Service Interface description Query the log download address of multiple domains. <br>The default granularity of log file is 24 hours and data are returned according to the actual configurations.<br>Around 3-5 hours of data delay | POST | /api/report/log/downloadLink |
| Querytranscodingdurationlogdownloadaddress | This interface is used to provide the download address of the transcoding duration log file of multiple domain names. Users can obtain the transcoding log information related to each domain name by specifying the time range and domain name list. The returned content includes the start and end time of the log file, download address, file name, file size, and expiration time of the download address. This interface helps users accurately grasp the transcoding details of each domain name within the specified time and make corresponding optimizations. | POST | /api/report/log/download-file/transcoding |
| Querycataloguetrafficandbroadcastcountdownload | This interface is used to query the download address for directory traffic statistics and playback count statistics of multiple domain names within a specified time range. Users need to specify the domain name and start/end time. If no domain name is specified, all domain names under the account will be queried by default (limited by quantity). The response will return information such as the download address, file size, expiration time, etc. of the log file. | POST | /api/report/catalogue/broadcast/download-file |