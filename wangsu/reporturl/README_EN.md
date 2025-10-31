# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/reporturl
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reporturl"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reporturl.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reporturl.{ActionName}Response{}
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
| Querytoprankingurl | This interface is used to count the number of URL requests and traffic of a domain name within each hour, and generate a ranking list of up to TOP500. Users need to provide the time return and domain name to obtain the total number of URL requests, total traffic, number of hit requests, number of failed requests, and details of different status codes. The interface supports sorting by total number of requests or total traffic to quickly identify high-frequency access or high-traffic URLs. This interface helps users monitor URL access frequency and traffic to adjust corresponding network policies. | POST | /api/report/url/top |
| Reporturloriginservice | Query URL information for Back-to-origin requests across multiple domains. By providing a list of domains and a time range, users can obtain URLs with Back-to-origin requests and the request count for each domain. It's useful for monitoring and analyzing these requests. | POST | /api/report/url/origin |
| Reporturlcustomtopdailyservice | Query the URL ranking for specified domains within a given time range. Users can specify the ranking criteria by either traffic or request count, with the default being sorted by request count. The API returns data that includes the ranking of each URL, the respective URL, the total traffic, and the total request count for each URL, assisting users in identifying URLs with concentrated traffic and request volumes. | POST | /api/report/url/custom-top/daily |
| Reportdomainrefererurlservice | This interface is used to query the URL source rankings of multiple domain names, It only supports data query in the East 8th District. Users can set the query time range and domain name list to obtain the corresponding ranking information. The interface will return the top 10 URL sources ranked by the number of requests or traffic value, and a maximum of the top 200 can be queried. This helps users identify URL source addresses with high traffic or high request counts for data analysis and performance optimization. | POST | /api/report/domain/referer-url |
| Reportdomainrefererwebsiteservice | Query the website referer ranking of multiple domains | POST | /api/report/domain/referer-website |