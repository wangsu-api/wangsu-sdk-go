# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/edgekv
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/edgekv"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &edgekv.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgekv.{ActionName}Response{}
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
| Deletekeyvalue | Delete multiple KV pairs from the namespace.  | DELETE | /edgekv/kv |
| Setkeyvalue | Write kv pairs to the specified namespace | PUT | /edgekv/kv |
| Getkeyvalue | get key value  from the specified namespace | POST | /edgekv/kv |
| Createshorturl | Short Url Create | POST | /short-urls/create |
| Getshorturl | query long url by short url | POST | /short-urls/query |
| Delshorturl | delete  short url | POST | /short-urls/del |
| Ecakvinfo | Query edge KV storage information, including: storage capacity, read request count, write request count, delete request count. | POST | /myview/Ecakvinfo |
| Sharkletvisit | Number of edge application requests, including basic application requests, medium-sized application requests, and special application requests | POST | /myview/sharkletVisit |