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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgeai"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &edgeai.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &edgeai.{ActionName}Response{}
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
| Createragfile | Upload rag files through this interface | POST | /rag/file/create |
| Enabledisablefile | enable file or disable file | POST | /rag/file/operate |
| Updatefile | update file | POST | /rag/file/update |
| Listragfile | query file list | POST | /rag/file/list |
| Ragfiledeleteservice | delete rag file | POST | /rag/file/delete |
| Ragknowledgebasedeleteservice | delete rag knowledge base | POST | /rag/knowledgebase/delete |
| Ragknowledgebasecreateservice | create rag knowledge base | POST | /rag/knowledgebase/create |
| Ragknowledgebaselistservice | query rag knowledge base | POST | /rag/knowledgebase/list |
| Ragknowledgebaseupdateservice | update rag knowledge base | POST | /rag/knowledgebase/update |
| Createwrconnector | This API is used to create a Waiting Room connector. It supports configuring parameters such as QT account information, connector name, synchronization cycle, whether to verify user identity, and queue token validity period. Upon success, the instance ID of the created connector will be returned. | POST | /api/v2/waiting_room_connectors |
| Updatewrconnector | This interface is used to update the configuration information of an existing waiting room connector. Users need to specify the connector instance ID via the REST parameter, and update its corresponding configurations such as QT Account ID, API Key, whether to generate a queue token, synchronization cycle, etc., through the request body parameters. The interface response will return the operation result, response data, and message. | PUT | /api/v2/waiting_room_connectors/* |
| Deletewaitingroomconnector | This API is used to delete a waiting room connector by its specified ID. Users need to provide the unique ID of the connector as a REST parameter for the deletion operation. | DELETE | /api/v2/waiting_room_connectors/* |
| Listconnectorsbypage | This API is used to retrieve a paginated list of connectors. Users can filter by name (fuzzy search) and ID (exact search), and specify the page number and page size. The response includes the total number of records and the list of connectors for the current page. | GET | /api/v2/waiting_room_connectors |
| Viewwrconnectorinfo | This interface is used to view the detailed information of a waiting room connector. Users need to specify the connector instance via the RESTful path parameter `id`. The interface will return detailed information such as the connector's configuration content, synchronization cycle, queue token validity period, status, creation, and modification times. | GET | /api/v2/waiting_room_connectors/* |
| Addimageconfig | This interface is used to upload or modify image processing configuration. Users upload configuration details via the `property` and `policySets` parameters in the request body, and the system returns the response code and information of the operation. | POST | /api/v2/ivm_configs |