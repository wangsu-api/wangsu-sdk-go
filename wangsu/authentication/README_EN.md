# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/authentication
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/authentication"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &authentication.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := authentication.{ActionName}Response{}
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
| Createuser | Create a User for SecureLink | POST | /api/securelink/idaas/user |
| Deletegroup | To delete a local user group, delete the group, users and subgroups in the group. | DELETE | /api/securelink/idaas/usergroup |
| Creategroup | Create a new group for SecureLink | POST | /api/securelink/idaas/usergroup |
| Modifygroup | Modify the specific Group | PUT | /api/securelink/idaas/usergroup |
| Listgroups | Query the groups of specific parent group | GET | /api/securelink/idaas/usergroup/list |
| Listusers | Example Query the list of users in a specified user group | GET | /api/securelink/idaas/user/list |
| Deleteuser | Delete a specific User | DELETE | /api/securelink/idaas/user |
| Modifyuser | Modify the information of the specific user. | PUT | /api/securelink/idaas/user |
| Describeuserinfo | Describe User | GET | /api/securelink/idaas/user |
| Syncauthconfig | sync auth config  | POST | /api/securelink/idaas/authconfig/sync |