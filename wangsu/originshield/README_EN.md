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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/originshield"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &originshield.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := originshield.{ActionName}Response{}
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
| Queryshieldroutemaps | This API is used to query the list of available route map for the origin shield. Users need to call this API to query and select available route map before creating an origin shield. | GET | /api/origin_shields/maps |
| Queryoriginshield | This API is used to query the Origin Shield. Users specify the Origin Shield ID to query the Origin Shield. | GET | /api/origin_shields/* |
| Disableoriginshield | This API is used to disable the origin shield. The user specifies the origin shield ID to disable it. Before disabling, ensure that the origin shield is not referenced by any deployed property. | POST | /api/origin_shields/*/disable |
| Enableoriginshield | This API is used to enable the origin shield. The user specifies the origin shield ID to enable the origin shield. Before enabling, ensure that all IP segments associated with the origin shield code are whitelisted. | POST | /api/origin_shields/*/enable |
| Deleteoriginshield | This API is used to delete the origin site protection shield. The user specifies the origin site protection shield ID to delete the origin site protection shield. | DELETE | /api/origin_shields/* |
| Updateoriginshield | This API is used to modify the origin shield. Users specify the name and description of the origin shield to modify it. | PUT | /api/origin_shields/* |
| Createoriginshield | This API is used to create an origin shield. Users select an available route map and specify the name of the origin shield to create it. | POST | /api/origin_shields |
| Confirmoriginshieldipsegmentchange | This API is used to confirm the change of IP segments under the origin shield. The user specifies the IP segment record ID under the origin shield to confirm the change of IP segments under the origin shield. | POST | /api/origin_shields/*/ip_segments/confirm |
| Queryoriginshields | This API is used to query the list of origin shields. Users can specify the origin shield ID, route map code, or origin shield name for filtered queries. | GET | /api/origin_shields |
| Queryavailableoriginshields | This API is used to query the list of available origin shields. Users need to specify the service type or property ID for the query. | GET | /api/origin_shields/available |
| Queryoriginshieldpendingipsegments | This API is used to query the list of IP segments pending confirmation for the origin shield. Users can specify the origin shield ID, action, and IP type for filtered queries. | GET | /api/origin_shields/*/ip_segments/pending |
| Queryoriginshieldactiveipsegments | This API is used to query the list of active IP segments of the origin shield. Users need to specify the origin shield ID for the query. | GET | /api/origin_shields/*/ip_segments/active |
| Queryoriginshieldrelatedinfos | This API is used to query the origin and domain information associated with the Origin Shield. Users need to specify the Origin Shield ID for the query. | GET | /api/origin_shields/*/related_infos |