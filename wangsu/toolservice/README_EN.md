# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/toolservice
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/toolservice"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &toolservice.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := toolservice.{ActionName}Response{}
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
| Querybandwidthlimittasklistservice | This interface is used to query the bandwidth limit task list under the account and return detailed information of all tasks with bandwidth. The returned content includes the domain name, task name, and the set maximum bandwidth value. This interface is suitable for scenarios where traffic control policies need to be evaluated and managed, helping users quickly identify and manage the currently set bandwidth limit tasks. | POST | /api/tools/queryBandwidthLimitTaskList |
| Icpqueryservice | Query whether specified domain has been registered in MIIT of Mainland China.<br><br>Limit: call frequency of the interface cannot exceed 50 times/day. | GET | /api/icp |
| Akamaiipforbiddenservice | This interface is used to block specific IP addresses from accessing AKAMAI services. Users can do this by entering a list of IP addresses to be blocked. This interface helps users block access from untrusted or harmful IP addresses and improve network security. | POST | /api/tools/ip-forbid/akamai |
| Reportserveripcountrycodeservice | Query the CDN service IP list for specific domains in different countries. Users provide the domain and country to get the IP list of coverage nodes for that domain. It's useful for viewing global CDN service IP distribution. | POST | /api/report/service-ip/country |
| Akamaiippermitservice | Akamaiippermitservice | POST | /api/tools/ip-permit/akamai |
| Forbidorresumevisitoripsbydomainservice | Forbid and Resume guest IP addresses from accessing the domain. | POST | /api/spider/ip-forbid |
| Queryforbiddingvisitoripsbydomainservice | Query Forbidding IPs by domain, support paging query. | POST | /api/spider/ip-forbid/query |
| Ipdomainservice | This interface is used to query the domain names that are using the IP address. The user enters the IP address to obtain the list of domain names associated with the IP. The information returned by the interface includes the current usage status of the IP and the list of domain names that use the IP. In actual applications, this interface can help users detect the domain name usage of a specific IP, which is suitable for network monitoring and management. | POST | /api/tools/ip/domain-list |
| Queryallbandwidthlimittasklistservice | This interface is used to query all bandwidth restriction tasks configured under the user account. When calling, the user can choose whether to include all customer domain names involved in the task and decide whether to return information on all task status. The returned data displays detailed information on each bandwidth restriction task in a list format, including task name, category, status, and related control policies and parameters. This interface helps users manage bandwidth settings and can effectively control and process specific traffic and requests in a timely manner. | POST | /api/tools/queryAllBandwidthLimitTaskList |
| Queryforbiddingvisitoripsbylabelcodeservice | Query the forbidding IPs by LabelCode, and support pagination.<br><br> | POST | /api/spider/label-ip-forbid/query |
| Forbidorresumevisitoripsbylabelcodeservice | A customer can create a label that can be associated with multiple domains and used to forbid or resume visitor IPs through that label . This has the same effect as forbidding or resuming the designated visitor IPs for all domains associated with the label. | POST | /api/spider/label-ip-forbid/operate |
| Addorremoveforbiddingipwhitelistservice | Provide the ability to add or remove IP whitelist. Depending on its granularity, IPs in the whitelist will be filtered for "Label Granularity" and "Domain Granularity" forbidding requests.<br>Supports distributing IP whitelist at three granularities. They are as follows:<br>"Customer Granularity": IP whitelists managed at customer affect all effective domains and effective labels under that customer.It will be filtered for both "Label Granularity" and "Domain Granularity" forbidding requests.<br>"Domain Granularity": IP whitelists managed at domain affect that effective domains.It will be filtered for only 'Domain Granularity' forbidding requests.<br>"Label Granularity": IP whitelists managed at label affect that effective labels.It will be filtered for only 'Label  Granularity' forbidding requests.<br>When adding, simultaneously resume IPs that are still forbidding at the corresponding granularity. | POST | /api/spider/ip-whitelist/operate |
| Queryforbiddingipwhitelistservice | Provide the ability to query the list of IPs on the forbidding whitelist. | POST | /api/spider/ip-whitelist/query |