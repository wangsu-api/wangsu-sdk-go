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
    response := &toolservice.{ActionName}Response{}
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
| Bandwidthlimitservice | Set/cancel bandwidth limits to certain specified domain. It's recommended that the call frequency is no higher than 1/5min.<br><br>Limit: no more than 3 domains can be operated on at a time. To set the limit higher, please contact technological customer service to evaluate. (The interface must be under strict control because it directly affects bandwidth) | POST | /api/tools/setBandwidthLimit |
| Querybandwidthlimittasklistservice | This interface is used to query the bandwidth limit task list under the account and return detailed information of all tasks with bandwidth. The returned content includes the domain name, task name, and the set maximum bandwidth value. This interface is suitable for scenarios where traffic control policies need to be evaluated and managed, helping users quickly identify and manage the currently set bandwidth limit tasks. | POST | /api/tools/queryBandwidthLimitTaskList |
| Icpqueryservice | Query whether specified domain has been registered in MIIT of Mainland China.<br><br>Limit: call frequency of the interface cannot exceed 50 times/day. | GET | /api/icp |
| Akamaiipforbiddenservice | This interface is used to block specific IP addresses from accessing AKAMAI services. Users can do this by entering a list of IP addresses to be blocked. This interface helps users block access from untrusted or harmful IP addresses and improve network security. | POST | /api/tools/ip-forbid/akamai |
| Reportserveripcountrycodeservice | Query the CDN service IP list for specific domains in different countries. Users provide the domain and country to get the IP list of coverage nodes for that domain. It's useful for viewing global CDN service IP distribution. | POST | /api/report/service-ip/country |
| Akamaiippermitservice | Akamaiippermitservice | POST | /api/tools/ip-permit/akamai |
| Supplyregisterservice |  supply register service<br> | POST | /sr/supply/register |
| Supplyincludeservice |  supply include service<br> | POST | /sr/supply/include |
| Ipdomainservice | This interface is used to query the domain names that are using the IP address. The user enters the IP address to obtain the list of domain names associated with the IP. The information returned by the interface includes the current usage status of the IP and the list of domain names that use the IP. In actual applications, this interface can help users detect the domain name usage of a specific IP, which is suitable for network monitoring and management. | POST | /api/tools/ip/domain-list |
| Queryallbandwidthlimittasklistservice | This interface is used to query all bandwidth restriction tasks configured under the user account. When calling, the user can choose whether to include all customer domain names involved in the task and decide whether to return information on all task status. The returned data displays detailed information on each bandwidth restriction task in a list format, including task name, category, status, and related control policies and parameters. This interface helps users manage bandwidth settings and can effectively control and process specific traffic and requests in a timely manner. | POST | /api/tools/queryAllBandwidthLimitTaskList |
| Queryconversiontaskdetail | This interface is used to query the detailed information of a specified conversion task, including task ID, configuration file type, conversion result, file name, creation time, domain list, and task status. Users need to provide the task ID and configuration file type as input parameters. | POST | /api/v1/akamai/get-task-detail |
| Queryconversiontasklist | This interface is used to query a list of conversion tasks. Users can filter tasks by conditions such as domain, task ID, start time, and end time. The interface returns detailed information about conversion tasks, including task ID, creation time, type, and status. | POST | /api/v1/akamai/wplus/get-task-list |
| Createwebsecurityconfigurationtask | This API is used to create a Web security configuration task. Users can create a task by uploading a file name and base64 encrypted file content, and providing a list of domain names. After the task is created, a task ID will be returned for subsequent task status queries. | POST | /api/v1/akamai/wplus/domain/move |
| Publishconverteddomain | This interface is used to publish domain names that have undergone security configuration transformation. Users can publish these domain names by providing a list of domains. The interface will return corresponding response messages and status codes, indicating whether the operation was successful. | POST | /api/v1/akamai/domain/complete |
| Createclientnetworklisttask | This API is used to create a client or network list task. The user needs to specify parameters such as task type and list content. Upon successful creation, the API will return the task ID and other related information. | POST | /api/v1/akamai/wplus/client-list/move |
| Querydeviceoperationalinfoservice | This interface is used to query device operational information, specifically whether it meets the conditions for quality inspection. Users can obtain the device status by providing the device serial number (SN) and query type (e.g., to query for qualifiable status). The response will return a query result code indicating whether the device exists, is qualifiable, or if there's an SN conflict, along with the device serial number and related information. | POST | /sr/supply/query |