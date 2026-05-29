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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportlog"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportlog.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &reportlog.{ActionName}Response{}
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
| Queryccllogs | Query cc logs | POST | /soc/api/LogManagement/queryCCAttackLogs |
| Getbotattackincidentlogdata | Get bot attack event log data | POST | /api/bot/report/event-log |
| GetNgRelayAppserviceIp | Provided to the ng CDN component to query the IP segment status of the ng parent appservice whitelist | POST | /api/GetNgRelayAppserviceIp |
| GetNgRelayWhiteIp | Querying the IP segment of the whitelist of ng transfer schemes on the task | POST | /api/GetNgRelayWhiteIp |
| Checkiswhiteip | This interface is used to supplement the function of querying whether CDN accelerates IP. It is mainly provided to customers to check whether a certain IP is the IP that actually provides acceleration to customers.<br>Working principle:<br>1 Technical support configures whitelist for customers on task and generates API address and parameters<br>2 The client forwards the w+ api request to the task's api address to check whether the specified ip is a cdn acceleration ip. If yes, it returns ip: yes, otherwise it returns ip: no | GET | /task/api/customers/whitelist-check |