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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/propertyconfig"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &propertyconfig.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := propertyconfig.{ActionName}Response{}
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
| Createproperty | Creates a property to define the configuration of one or more hostnames (domains) to deploy to the CDN servers. | POST | /api/properties |
| Listproperties | Get List of Properties | GET | /api/properties |
| Deleteproperty | Delete a property | DELETE | /api/properties/* |
| Getproperty | Query a property | GET | /api/properties/* |
| Listpropertyversions | Get list of property versions | GET | /api/properties/*/versions |
| Createpropertyversion | Create a new version of a property. | POST | /api/properties/*/versions |
| Getpropertyversion | Get the detailed configuration of a property version. | GET | /api/properties/*/versions/* |
| Updatepropertyversion | Update a property version | PUT | /api/properties/*/versions/* |
| Createdeploymenttask | Create a deployment task to deploy your property to the staging or production environments or to remove the property.  | POST | /api/properties/deployments |
| Listdeploymenttasks | This interface is used to query a list of deployment tasks that meet specified criteria. Users can filter by property ID, task status, deployment environment, and control the query start position, maximum number of items, sort order, and sort field. The interface returns the total number of deployment tasks and a detailed list of tasks. | GET | /api/properties/deployments |
| Getdeploymenttask | Get deployment task | GET | /api/properties/deployments/* |
| Createpropertyforakamaimigration | Creates a property to define the configuration of one or more hostnames (domains) to deploy to the Akamai China CDN servers. | POST | /api/properties/migration |
| Createdeploymenttaskforterraform | This API is used to create deployment tasks in Terraform scenarios to deploy properties to staging or production environments, or to uninstall already deployed properties. The user needs to specify the property ID, version, and operation type. | POST | /api/terraform/properties/deployments |
| Createpropertyforterraform | This API is used to create a property in the Terraform scenario. Users need to specify domain configuration, rule configuration, origin configuration, variable configuration, and other information. | POST | /api/terraform/properties |
| Deletepropertyforterraform | This API is used to delete property configuration in Terraform scenarios. The user needs to specify the property ID to delete. | DELETE | /api/terraform/properties/* |
| Updatepropertyforterraform | This API is used to update a property in the Terraform scenario. Users need to specify domain configuration, rule configuration, origin configuration, variable configuration, and other information. | PUT | /api/terraform/properties/* |
| Querydeploymentforterraform | This API is used to query the details of deployment tasks in Terraform scenarios. Users need to specify the deployment task ID to query. | GET | /api/terraform/properties/deployments/* |
| Querypropertiesforterraform | This API is used for querying the property list in Terraform scenarios. Users can specify the service type and deployment environment for filtered queries. | GET | /api/terraform/properties |
| Querypropertyversionconfigforterrform | This API is used to query the property version configuration for Terraform scenarios. The user needs to specify the property ID and the property version to query. | GET | /api/terraform/properties/*/versions/* |
| Querydeploymentsforterraform | This API is used to query the list of deployment tasks in Terraform scenarios. Users can filter by specifying property ID, deployment environment, etc. | GET | /api/terraform/properties/deployments |
| Querypropertyconfigforterrform | This API is used for querying property configurations in Terraform scenarios. Users need to specify a project ID to query. | GET | /api/terraform/properties/* |
| Querytierroutemaps | This API is used to query the list of available tier route map in the property configuration. Users need to specify the service type or property id for the query. | GET | /api/maps |
| Queryipsegmentbyroutemapcode | This API is used to query the IP segments associated with the route map. Users specify the route map code to query all associated IP segments. | GET | /api/origin_shields/maps/ip_segments |