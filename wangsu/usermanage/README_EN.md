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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/usermanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &usermanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &usermanage.{ActionName}Response{}
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
| Usercustomerupdateservice | Usercustomerupdateservice | POST | /api/user/customer/update |
| Addsubaccount | Create a sub account | POST | /sub-account |
| Querysubaccountinfo | Get sub account | GET | /sub-account/* |
| Getsubaccountlist | Get the list of sub-users by the main account | POST | /sub-account/list |
| Updatesubaccount | Update sub account.  | PUT | /sub-account |
| Deletesubaccount | Delete Specified sub-users | DELETE | /sub-account/* |
| Querypolicyattachedmainaccountorsubaccount | This interface is used to query the list of permission policies associated with a specified user. By entering their login name, users can obtain information about the policies linked to them, including the policy ID, name, description, type (system policy or custom policy), and support for multilingual descriptions. The results returned by the interface include status codes and related information to help users understand the specific attributes and classifications of the policies. This is very useful for managing user permissions and creating or modifying permission policies, enabling system administrators to configure and adjust permissions more efficiently. | GET | /user/policy-attached/* |
| Batchaddorrevokepolicytosubaccount | This interface is used to batch add or revoke permission policies for a specified sub-user. By inputting the sub-user's login name and a list of permission policy identifiers, you can choose to perform add or revoke operations. When adding permissions, the system will add the corresponding policies to the sub-account to expand its permission scope; when revoking permissions, it will remove the specified policies, reducing the sub-account's access permissions. The return value includes a request status code and operation information prompts, allowing users to confirm the success of the batch operation. This interface is suitable for scenarios where centralized management of sub-account permissions is needed, simplifying the process of batch permission adjustments.<br><br><br> | POST | /user/policies |
| Checkloginnamelegal | Check login name legal | POST | /login-name/check |
| Addaccountident | Add account aksk<br> | POST | /account-ident |
| Deleteaccountident | Delete account aksk | DELETE | /account-ident/* |
| Updateaccountident | Update account aksk<br> | PUT | /account-ident |
| Listaccountident | List account aksk<br> | POST | /account-ident/list |
| Queryagentassociatedmainaccountservice | This interface is used to query the relationship between multi-level agent accounts and their associated main accounts. Through this interface, users can understand detailed information about each agent account's main account, including the main account's display name, login name, and the login name of its corresponding parent main account. When calling this interface, users receive status codes and messages about the request, ensuring transparency of the query execution. This interface is suitable for scenarios requiring the management and querying of complex hierarchical agent account relationships, helping to enhance overall understanding and control of account management. | POST | /user/multilevel-agent/main-accounts |
| Listgroupusers | This interface is used to query the users of a group. You can query by providing either the group ID or the group name. | POST | /user/groups/list/users |
| Modifygrouppolicy | This API is used to modify the permission policy of a user group. You can add or revoke permission policies by providing either the user group ID or the user group name as input parameters. | POST | /user/groups/policies |
| Listgroupattachedpolicy | This API is used to query the policy list (i.e., permission list) associated with a specified user group. You must specify either the user group ID or user group name, and the API will return detailed policy information associated with the user group, including policy ID, name, type, and description. | POST | /user/groups/attached/policies |
| Editgroupbasicinfo | This API is used to modify the basic information of user groups. By using the user group ID or user group name, the display name and notes for the user group can be modified. | POST | /user/groups/edit/basicinfo |
| Creategroup | This API is used to create a user group. Users need to provide the user group name, user group display name, and user group note through the request body. After successful creation, the interface will return the ID of the new user group. | POST | /user/groups/create |
| Deletegroup | This API is used to delete the specified user group. You can delete a user group by either groupId or groupName. | POST | /user/groups/delete |
| Modifygroupuser | The API is used to modify user group users. You can update user group users by specifying a user group ID or name, and providing the type of operation (such as adding or deleting) and a specific user list. | POST | /user/groups/edit/user |
| Listgroupbasicinfo | This API is used to query basic information about user groups. When a specific group ID or group name is provided, it returns information for the corresponding user group. If no parameters are provided, it queries all user groups under the main account. | POST | /user/groups/list/basicinfo |