# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/reportvisitor
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportvisitor"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportvisitor.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportvisitor.{ActionName}Response{}
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
| Querytotalnumberofuniqueipundersingledomain | This interface is used to query the number of independent IP addresses for each stream name within a single domain name. The user provides a specific domain name and time to obtain information about the specified stream during this period (in days). The results returned by the interface include statistics on independent IP addresses for the domain name and its corresponding stream name. This helps users understand the distribution of visitors to each stream, thereby optimizing traffic configuration or evaluating the use of streaming media services. | POST | /api/report/visitor/total/stream |
| Reportuvispprovinceservice | Query unique visitor IPs for multiple domains by province and ISP. Users can specify time range, domain, province, ISP, and choose grouping by domain, province, or ISP. Results aid in analyzing access trends and optimizing content distribution. Supports Accept-Language (zh-CN, en-US), default is zh-CN. With en-US, province and ISP are in codes; otherwise, in Chinese. | POST | /api/report/uv/isp-province |
| Reportiptopdetailsservice | Query 5 minute details of multiple domain names. TOP IP | POST | /api/report/ip/top-details |
| Reportreferrertopdetailsservice | This interface is used to query the top visitor source details of a domain name every 5 minutes. The user needs to provide the domain name and time range, and the returned content includes detailed data of each referral source. The optional values include traffic, bandwidth, and number of requests. This interface helps users analyze site traffic and make corresponding business decisions. | POST | /api/report/referrer/top-details |
| Reporturltopdetailsservice | This interface is used to count the top URL information of multiple domain names within 5 minutes. The user needs to provide the time range and domain name. The returned content includes the traffic, bandwidth or number of requests of the top URL of the domain name. It helps users understand the popular access links and optimize resource allocation. | POST | /api/report/url/top-details |
| Reportvisitorcustomtopdailyservice | Reportvisitorcustomtopdailyservice | POST | /api/report/visitor/custom-top/daily |