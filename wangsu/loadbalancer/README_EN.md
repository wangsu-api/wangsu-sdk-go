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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/loadbalancer"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &loadbalancer.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &loadbalancer.{ActionName}Response{}
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
| Createloadbalancerlistener | This interface is used to create a listener for a load balancer instance. Users need to specify parameters such as the listener protocol, port, load balancer ID, server pool ID, and load balancing algorithm. Additionally, users can enable health checks and session persistence as needed, configuring detailed health check policies (e.g., check protocol, port, path, timeout) and session persistence timeout. Upon successful creation, the interface will return the ID of the newly created listener. | POST | /vmp/slb/listeners |
| Queryloadbalancerlisteners | This API is used to query the list of load balancer listeners. Users can filter based on listener ID, load balancer ID, listener status, server pool ID, etc. The API returns detailed information about the listeners, including health check configuration, session persistence settings, etc. | GET | /vmp/slb/listeners |
| Updateloadbalancerlistener | This API is used to update the configuration of a specified load balancer listener. By providing the `listenerId`, users can modify its protocol, port, name, session persistence settings, server pool ID, load balancing algorithm, and detailed health check configurations, including whether to enable health check, its protocol, port, path, timeout, interval, retry count, and normal status codes. | PUT | /vmp/slb/listeners/* |
| Updateloadbalancerserverpool | This API is used to update an existing load balancer server pool. Users must specify the `poolId` of the server pool to be updated. The API allows updating server pool members (including RS instance ID, weight, and service port), the server pool name, and remark information. | PUT | /vmp/slb/serverpools/* |
| Createloadbalancerserverpool | This API is used to create a load balancing server pool. Users need to specify the Load Balancer ID (`lbId`), server pool members (`members`), server pool name (`name`), and optional remark information (`remark`) in the request body. Server pool members can include RS instance ID (`instanceId`), weight (`weight`), and service port (`servicePort`). Upon successful creation, the API will return detailed information about the new server pool, including its ID (`id`) and a `data` object containing specific configurations such as node name, VIP information, network type, IP protocol, VIP address, creation/modification time, load balancing name, and status. | POST | /vmp/slb/serverpools |
| Createloadbalancer | This interface is used to create a load balancer object. Users need to provide the node name and instance specification, and can optionally specify the load balancer's carrier line, network type, load balancer name, and IP protocol. Upon successful creation, the interface will return detailed information about the load balancer, including the node name, VIP information, network type, IP protocol, VIP address, creation and modification times, load balancer name, state description, ID, and specification type. | POST | /vmp/slb/loadbalancers |
| Queryloadbalancer | This interface is used to query the load balancer list. Users can filter queries by specifying parameters such as node name, load balancer ID, state, and specification. All parameters support multiple values. The interface will return detailed information of the compliant load balancers, including their node name, VIP information, network type, IP protocol, VIP address, modification time, creation time, name, state description, ID, state, and specification type. | GET | /vmp/slb/loadbalancers |
| Deleteloadbalancerserverpools | This interface is used to delete specified load balancer server pools. Users need to provide the ID of the server pool to be deleted. Batch deletion is supported by providing multiple IDs separated by commas. | DELETE | /vmp/slb/serverpools/* |
| Queryloadbalancerserverpools | This API is used to query a list of load balancer server pools based on specified criteria (e.g., Load Balancer Server Pool ID or Load Balancer ID). The results include detailed information for each server pool, such as its associated RS instances. | GET | /vmp/slb/serverpools |
| Deleteloadbalancerlisteners | This API is used to delete specified load balancer listeners. Users need to provide the IDs of the load balancer listeners to be deleted, supporting multiple IDs separated by commas for batch deletion. Upon successful operation, the system will return a success message. | DELETE | /vmp/slb/listeners/* |
| Startloadbalancerlistener | This API is used to enable specified load balancer listeners. Users need to provide a list of IDs of one or more load balancer listeners, and upon successful operation, the API will return the enablement result. | PUT | /vmp/slb/listeners/action/start |
| Enableloadbalancer | This API is used to enable specified load balancer instances. Users need to provide a list of one or more load balancer IDs, and upon successful operation, the API will return the enablement result. | PUT | /vmp/slb/loadbalancers/action/start |
| Disableloadbalancer | This API is used to disable specified load balancer instances. Users need to provide a list of one or more load balancer IDs, and upon successful operation, the API will return the disablement result. | PUT | /vmp/slb/loadbalancers/action/stop |
| Stoploadbalancerlistener | This API is used to stop specified load balancer listeners. Users need to provide a list of one or more load balancer listener IDs. Upon success, the API will return the stop result. | PUT | /vmp/slb/listeners/action/stop |
| Deleteloadbalancer | This API is used to delete the specified load balancer instance. The user needs to provide the ID of the load balancer to be deleted, and the API will return the operation result. | DELETE | /vmp/slb/loadbalancers/* |