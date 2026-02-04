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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgefunc"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &edgefunc.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &edgefunc.{ActionName}Response{}
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
| Uploadfunccode | Function code upload interface, for the scenario that does not use CloudIDE, but directly calls the interface to upload function code | POST | /edgefunc/upload |
| Createedgefunctrigger | This interface is used to create function triggers. Users are required to provide the domain and a list of routing rules as request parameters. Upon successful creation, the interface will return a status code and the list of successfully created triggers. | POST | /api/v2/cdn_triggers |
| Queryedgefunctrigger | This interface is used to query a list of function triggers. Users can filter by domain or function name and use pagination. The response will return a detailed list of triggers. | GET | /api/v2/cdn_triggers |
| Deletefuncdomaintrigger | This interface is used to remove an existing function trigger. The user needs to specify the unique identifier of the function trigger to be deleted via the REST parameter 'id'. | DELETE | /api/v2/cdn_triggers/* |
| Queryfunctionlist | This interface is used to query the function list. Users can perform exact queries by function ID or fuzzy queries by function name. It also supports paginated queries, requiring the specification of page number and page size. The interface returns the function list, total count, and operation status information. | GET | /api/v2/edge_funcs |
| Getedgefunctioninfo | This interface is used to query information about a specified edge function. Users need to provide the function ID (which can be passed via Query or REST path parameters), and the interface will return detailed information about the function, including the function name, creation time, test domain, memo, and update time. | GET | /api/v2/edge_funcs/* |
| Edgefuncgetdebuglog | This interface is used to query the debug logs of a specified function. Users need to provide the function ID and debug session ID to obtain the corresponding log content. | GET | /api/v2/edge_funcs/*/logs/* |
| Queryfunctioncode | This interface is used to query the latest or currently deployed code for a function. Users need to provide the function ID as a parameter. | GET | /api/v2/edge_funcs/*/files |
| Savefunctioncode | This API is used to save the code for a specified function. Users need to provide the function ID. | POST | /api/v2/edge_funcs/*/files |
| Deleteedgefunction | This interface is used to delete a function based on its ID. | DELETE | /api/v2/edge_funcs/* |
| Listfunctiontemplates | his interface is used to query the list of edge function templates. Users can filter and paginate queries by specifying the page number `pageNo`, page size `pageSize`, template name `name` (supports fuzzy queries), and whether the template name is in English `isEnglish`. The interface response will return the list of matching function templates `templates`, along with the total number of records `total`, current page number `pageNum`, page size `pageSize`, operation status code `code`, and description message `message`. | GET | /api/v2/edge_funcs/templates |
| Createedgefunc | This API is used to create an edge function. Users need to provide parameters such as function name, template name, test domain, memo, and function alias. Upon successful creation, the function ID will be returned | POST | /api/v2/edge_funcs |
| Edgefuncgetdebugurl | This interface is used to retrieve the test URL for a specified edge function. Users need to provide the function ID, and the interface will return the operation result and relevant status information. | GET | /api/v2/edge_funcs/*/debug_url |
| Updateedgefunc | This API is used to update an edge function. Users need to provide parameters such as function ID, function name, template name, test domain, memo, and function alias. | PUT | /api/v2/edge_funcs/* |