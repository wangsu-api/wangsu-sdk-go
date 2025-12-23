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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/networkmanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &networkmanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := networkmanage.{ActionName}Response{}
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
| Vmpreleaseedgeip | Used to release exclusive and drifting additional public IP addresses.<br>Explanation:<br>1) Drift type additional IP needs to be unbound with the cloud host before it can be released. If it is not unbound, the release will fail;<br>2) If multiple drifting IPs are released in bulk, some IPs have been unbound and some IPs have not been unbound, the unbound ones will be released normally, while the unbound ones will fail to be released.<br>3) Exclusive IP is released without unbinding. | PUT | /vmp/edgeIp/release |
| Vmpallocateedgeip | Used to apply for drift type additional public IP. Drift mode supports the simultaneous use of multiple instances of the same IP, and is commonly used in master-slave switching scenarios, such as LVS. | POST | /vmp/edgeIp/allocate |
| Vmpassignedgeip | Bind the requested drifting additional IP to the specified instance.<br>Drift type additional IP supports the simultaneous use of multiple instances of the same IP, commonly used in master-slave switching scenarios, such as LVS. | PUT | /vmp/edgeIp/assign |
| Vmpunassignedgeip | Can be used to unbind drifting additional public network IPs and instances. Drift type additional public IP supports the simultaneous use of multiple instances of the same IP, and is commonly used in master-slave switching scenarios, such as LVS.<br>Explanation:<br>1) The drifting additional public IP to unbind must be an IP bound to the specified instance, and cannot be an IP bound to other virtual machines;<br>2) If multiple drifting additional public network IPs are unbound in bulk, and some IPs are bound to other instances, all IPs fail to be unbound, and the interface returns an error message. | PUT | /vmp/edgeIp/unassign |
| Vmpqueryedgeip | Used to query edge public IP addresses that have been applied for. | GET | /vmp/edgeIp |
| Listservice | query list of service | GET | /api/v1/namespaces/*/services |
| Getservice | query details of service | GET | /api/v1/namespaces/*/services/* |
| Createservice | create service | POST | /api/v1/namespaces/*/services |
| Updateservice | update service | PUT | /api/v1/namespaces/*/services/* |
| Deleteservice | delete service | DELETE | /api/v1/namespaces/*/services/* |
| Vmpedgeipallocate4occupancy | Used to apply for exclusive additional public IP addresses. Once the IP is applied, it will be bound to the specified instance simultaneously. One IP can only be bound to one instance. | POST | /vmp/edgeIp/allocate4Occupancy |
| Putpatchservice | partial update service | PUT | /api/v1/namespaces/*/services/*/ws/patch |
| Getnetworkpolicy | get network policy details info | GET | /apis/networking.k8s.io/v1/namespaces/*/networkpolicies/* |
| Listnetworkpolicy | list network policy | GET | /apis/networking.k8s.io/v1/namespaces/*/networkpolicies |
| Createnetworkpolicy | create network policy | POST | /apis/networking.k8s.io/v1/namespaces/*/networkpolicies |
| Deletenetworkpolicy | delete network policy | DELETE | /apis/networking.k8s.io/v1/namespaces/*/networkpolicies/* |
| Updatenetworkpolicy | update network policy | PATCH | /apis/networking.k8s.io/v1/namespaces/*/networkpolicies/* |
| Putnetworkpolicy | put network policy | PUT | /apis/networking.k8s.io/v1/namespaces/*/networkpolicies/* |
| Createingresscontroller | create ingress controller | POST | /openapi/custom/api/v1/ingress-contrl |
| Updateingresscontroller | update ingress controller | PUT | /openapi/custom/api/v1/ingress-contrl/* |
| Deleteingresscontroller | delete ingress controller | DELETE | /openapi/custom/api/v1/ingress-contrl/* |
| Getingresscontroller | query ingress controller detail | GET | /openapi/custom/api/v1/ingress-contrl/* |
| Listingresscontroller | query ingress controller list | GET | /openapi/custom/api/v1/ingress-contrl |
| Createingress | create ingress | POST | /openapi/custom/api/v1/ingress |
| Updateingress | update ingress | PUT | /openapi/custom/api/v1/ingress/* |
| Deleteingress | delete ingress | DELETE | /openapi/custom/api/v1/ingress/* |
| Getingress | query ingress detail | GET | /openapi/custom/api/v1/ingress/* |
| Listingress | query ingress list | GET | /openapi/custom/api/v1/ingress |
| Edgeiprivatepallocate | Used to apply for drift type additional private IP | POST | /vmp/edgeIp/private/allocate |
| Releaseedgeprivateip | Release edge private IP | PUT | /vmp/edgeIp/private/release |
| Assignedgeprivateip | Assign edge private IP | PUT | /vmp/edgeIp/private/assign |
| Unassignedgeprivateip | Unassign edge private IP | PUT | /vmp/edgeIp/private/unassign |
| Queryedgeprivateip | Query edge private IP | GET | /vmp/edgeIp/private |
| Vmpqueryavailablecidrs | Query available cidrs in a node | GET | /vmp/cidrs |
| Vmpqueryavailablecidrsdetail | This interface fetches available subnet details, including free IP count and attributes, within a specified node by its name, for IP allocation or instance creation. | GET | /vmp/cidrs/detail |
| Lechqueryavailablecidrs | Query available cidrs in a node | GET | /lech/cidrs |
| Lechqueryavailablecidrsdetail | This interface fetches available subnet details, including free IP count and attributes, within a specified node by its name, for IP allocation or instance creation. | GET | /lech/cidrs/detail |
| Lechedgeipallocate4occupancy | Used to apply for exclusive additional public IP addresses. Once the IP is applied, it will be bound to the specified instance simultaneously. One IP can only be bound to one instance. | POST | /lech/edgeIp/allocate4Occupancy |
| Lechallocateedgeip | Used to apply for drift type additional public IP. Drift mode supports the simultaneous use of multiple instances of the same IP, and is commonly used in master-slave switching scenarios, such as LVS. | POST | /lech/edgeIp/allocate |
| Lechassignedgeip | Bind the requested drifting additional IP to the specified instance.<br>Drift type additional IP supports the simultaneous use of multiple instances of the same IP, commonly used in master-slave switching scenarios, such as LVS. | PUT | /lech/edgeIp/assign |
| Lechqueryedgeip | Used to query edge public IP addresses that have been applied for. | GET | /lech/edgeIp |
| Lechreleaseedgeip | Used to release exclusive and drifting additional public IP addresses.<br>Explanation:<br>1) Drift type additional IP needs to be unbound with the cloud host before it can be released. If it is not unbound, the release will fail;<br>2) If multiple drifting IPs are released in bulk, some IPs have been unbound and some IPs have not been unbound, the unbound ones will be released normally, while the unbound ones will fail to be released.<br>3) Exclusive IP is released without unbinding. | PUT | /lech/edgeIp/release |
| Lechunassignedgeip | Can be used to unbind drifting additional public network IPs and instances. Drift type additional public IP supports the simultaneous use of multiple instances of the same IP, and is commonly used in master-slave switching scenarios, such as LVS.<br>Explanation:<br>1) The drifting additional public IP to unbind must be an IP bound to the specified instance, and cannot be an IP bound to other virtual machines;<br>2) If multiple drifting additional public network IPs are unbound in bulk, and some IPs are bound to other instances, all IPs fail to be unbound, and the interface returns an error message. | PUT | /lech/edgeIp/unassign |