# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/basicpermission
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/basicpermission"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &basicpermission.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := basicpermission.{ActionName}Response{}
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
| Addterminalauth | Add a terminal authority. | POST | /api/securelink/idaas/terminalauth |
| Updateterminalauth | edit a terminal authority. | PUT | /api/securelink/idaas/terminalauth |
| Deleteterminalauth | delete a terminal auth by terminalAuthId. | DELETE | /api/securelink/idaas/terminalauth |
| Addterminalauthuserorgroup | Basic permissions are associated with users or user groups. | POST | /api/securelink/idaas/terminalauth/user-group |
| Queryterminalauthlist | Get a list of basic permissions. | GET | /api/securelink/idaas/terminalauth |
| Queryterminalauthinfo | Query basic permission information | GET | /api/securelink/idaas/terminalauth/query |
| Removeterminalauthresource | Remove apps associated with basic permissions | DELETE | /api/securelink/idaas/terminalauth/related-resources |
| Removeterminalauthuserorgroup | Remove users or user groups associated with basic permissions | DELETE | /api/securelink/idaas/terminalauth/user-group |
| Queryterminalauthresourcelist | Get the list of applications with basic permissions | GET | /api/securelink/idaas/terminalauth/resources |
| Updateterminalauthrelatedresources | Modify applications associated with basic permissions | PUT | /api/securelink/idaas/terminalauth/related-resources |
| Queryterminalusergroupbyauthconfig | query terminal user group list by auth config. | GET | /api/securelink/idaas/usergroup/list-by-authconfig |
| Associaterightsgroupstouserorusergroup | This API is used to directly associate a permission group with specified users/user groups. If multiple users need to be assigned the same set of application permissions, the administrator can first associate the applications with a single permission group and then use this API to assign it to the users/user groups. | POST | /api/securelink/idaas/terminalauth/user-bind |