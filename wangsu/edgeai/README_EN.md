# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/edgeai
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
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
    response := edgeai.{ActionName}Response{}
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