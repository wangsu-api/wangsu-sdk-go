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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgehostname"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &edgehostname.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgehostname.{ActionName}Response{}
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
| Deleteedgehostname | This API is used to delete a specified Edge Hostname. If the Edge Hostname has never been deployed, calling the API directly deletes the record. If the Edge Hostname is in a deployed state, calling this API will trigger the removal of the deployment and return the deployment task ID. | DELETE | /api/edge-hostnames/* |
| Deployedgehostnamedns | Deploy edge hostname is used to distribute the configuration of the edge hostname to global edge DNS servers. After deployment, the edge DNS server will execute traffic routing based on the DNS policy. | POST | /api/edge-hostnames/*/deploy |
| Undeployedgehostnamedns | Undeploy Edge Hostname is used to remove the configuration of the edge hostname from global edge DNS servers. After undeployment, the edge DNS server will no longer perform traffic routing for this hostname. | POST | /api/edge-hostnames/*/disable |
| Getedgehostname | Query the configuration details of the specified EdgeHostname, including:  name, description, DNS service status, DNS deployment status, whether acceleration in Mainland China is allowed, list of associated accelerated domains, DNS scheduling rules, etc. | GET | /api/edge-hostnames/* |
| Listedgehostnames | This API is used to query the list of Edge Hostnames, returning all Edge Hostnames under the account along with their CNAME status and deployment status, with support for pagination. | GET | /api/edge-hostnames |
| Updateedgehostname | This interface is used to modify the basic configuration of the Edge Hostname, enabling self-service control over the routing effects, including remarks, acceleration region restrictions, and region configuration. | PUT | /api/edge-hostnames/* |
| Deployedgehostnameforterraform | This API is used for deploying the specified edge-hostname in Terraform scenarios. Deployment refers to issuing the edge-hostname configuration and making its DNS service effective. After deployment, the accelerated domain name using this scheduling domain can normally use CDN services. Additionally, when your edge-hostname deployment status is pending or failed, this API can be used to redeploy the edge-hostname. | POST | /api/terraform/edge-hostnames/*/deploy |
| Deleteedgehostnameforterraform | This API is used to delete a specific edge-hostname in the Terraform scenario. Deletion refers to uninstalling the DNS deployment and deleting data. Only edge-hostname not associated with any acceleration domain can be deleted. Note: Deletion cannot be undone, please perform deletion with caution. | DELETE | /api/terraform/edge-hostnames/* |
| Queryedgehostnameforterraform | This API is used to query edge-hostname details in Terraform scenarios. Users need to specify the edge-hostname for the query. | GET | /api/terraform/edge-hostnames/* |
| Updateedgehostnameforterraform | This API is used to modify the configuration of a specified edge-hostname in a Terraform scenario, supporting configuration of partition scheduling rules. Users can specify modifications to the edge-hostname configuration and partition scheduling rules. | PUT | /api/terraform/edge-hostnames/* |
| Queryedgehostnamesforterraform | This API is used in Terraform scenarios to query the list of edge-hostname. Returning all edge-hostnames under the account along with their CNAME status and deployment status, with support for pagination. | GET | /api/terraform/edge-hostnames |