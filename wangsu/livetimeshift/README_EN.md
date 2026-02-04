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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/livetimeshift"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &livetimeshift.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &livetimeshift.{ActionName}Response{}
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
| Querytimeshiftparametertemplate | This interface is used to query all time-shift parameter templates created under a customer's account. Users can query specific template parameters by providing an optional template ID, or filter templates by region. A successful call will return a list containing all time-shift parameter templates or the details of the specified template. | GET | /api/v2/streams/time-shift-parameter-template |
| Modifytimeshiftparametertemplate | This interface is used to edit and modify the parameters of a time-shift template. To avoid impacting ongoing services, modification is prohibited when the template is in use. The target template is identified by its ID (`templateId`) in the URL path, and details like the template name, storage configuration, callback URL, and time-shift parameter list are updated from the request body. | PUT | /api/v2/streams/time-shift-parameter-template/* |
| Addtimeshiftparametertemplate | This interface allows you to create live stream time-shift related parameter templates in advance. These templates primarily store time-shift parameters, callback notification parameters, and other storage parameters. Upon successful creation, the interface will return the generated template ID. | POST | /api/v2/streams/time-shift-parameter-template |
| Deletetimeshiftparametertemplate | This interface is used to delete a specified time-shift parameter template. Users must provide the unique identifier `templateId` of the template to be deleted in the path parameters. Deletion is prohibited if the template is currently in use by any process or already associated with time-shift rules. Upon successful deletion, the response body will return a `code` and `message`. | DELETE | /api/v2/streams/time-shift-parameter-template/* |
| Querytimeshiftrule | This API allows you to query created time-shift rules. It supports filtering by rule ID, domain, publishing point, and stream name. If no filtering conditions are specified, it will query all rules by default. A successful query will return a list of matching rules, including their detailed information such as creation time and enabled status. | GET | /api/v2/streams/time-shift-rules |
| Stoprealtimetimeshiftrecording | This interface is used to terminate an ongoing real-time time shift recording task. Users need to provide the unique identifier `persistentId` to specify the task. Upon successful invocation, the API returns the overall response code and message. | PUT | /api/v2/streams/time-shift-recordings/* |
| Startrealtimetimeshiftrecording | This API is used to perform real-time time-shift recording on live streams. Users can select the time shift recording target by specifying the pull domain, publishing point, stream name, and time shift recording template. Upon successful invocation, the API will return a recording task ID and the time shift recording status for each stream. If premature termination of time shift recording is required, the "Stop Real-Time Time-Shift Recording" API can be called; otherwise, the recording task will automatically terminate after the live stream concludes. | POST | /api/v2/streams/time-shift-recordings |
| Addtimeshiftrules | This interface allows you to add time-shift rules, which are used to specify live streams for recording and enable time-shift functionality. Users must provide rule parameters such as the push domain, publishing point, and stream name, and can also specify the associated time-shift template ID and the rule's enabled status. When a corresponding rule is triggered, the system will record video based on the associated time-shift template. | POST | /api/v2/streams/time-shift-rules |
| Deletetimeshiftrules | This interface is used to delete a previously created time-shift rule. Deletion of a rule that is currently executing is prohibited. Users must specify the time-shift rule to be deleted via its rule ID. The interface returns information on the success or failure of the deletion operation. | DELETE | /api/v2/streams/time-shift-rules/* |
| Modifytimeshiftrules | This interface is used to modify existing time-shift rules. Users need to provide the unique identifier (ruleId) of the rule to be modified and specify the rule attributes to be updated in the request body, such as publish point, domain, enabled status, template ID, and stream-related parameters. | PUT | /api/v2/streams/time-shift-rules/* |