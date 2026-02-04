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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/livesnapshot"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &livesnapshot.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &livesnapshot.{ActionName}Response{}
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
| Querysnapshotparametertemplate | This interface allows you to query all screenshot parameter templates created under a customer account. You can query specific template parameters by template ID, or query all templates without specifying parameters. The interface returns details of all matching templates, including the template name, screenshot parameter list, and storage configuration. | GET | /api/v2/streams/snapshot-parameter-templates |
| Addsnapshotparametertemplate | This interface is used to create templates for live screenshot parameters, allowing users to pre-configure settings such as screenshot details, callback notifications, and storage options. Users can specify parameters like template name, screenshot parameter list, storage configuration, and callback notification URL in the request body. Upon successful creation, the interface returns the ID of the newly created template. | POST | /api/v2/streams/snapshot-parameter-templates |
| Deletesnapshotparametertemplate | This API is used to delete the specified screenshot parameter template. Users need to provide the ID of the template to be deleted. Please note that if the template is in use or already associated with screenshot rules, direct deletion is prohibited. | DELETE | /api/v2/streams/snapshot-parameter-templates/* |
| Modifysnapshotparametertemplate | This interface is used to edit and modify the content of a screenshot parameter template. Users can specify the template to be modified using `templateId` and update its configuration via request body parameters such as `templateName` and `snapshotParams`, including details like screenshot width, height, interval, file format, storage bucket, storage path, stream pull timeout, and callback notifications. Modification is prohibited when the template is currently in use. | PUT | /api/v2/streams/snapshot-parameter-templates/* |
| Modifysnapshotrules | This interface is used to modify existing snapshot rules. Users specify the rule to be modified by its rule ID and can update the rule's content by providing information such as app name, domain, stream name, stream extension parameters, enable status, and template ID. | PUT | /api/v2/streams/snapshot-rules/* |
| Querysnapshotrule | This API is used to query already created screenshot rules. Users can filter by screenshot rule ID, domain, publishing point, or stream name. If no filtering criteria are specified, all rules will be queried by default. The interface will return a list of matching screenshot rules, each containing information such as pull stream domain, creation time, publishing point, domain, enabled status, rule ID, template ID, stream name, and extended stream parameters. | GET | /api/v2/streams/snapshot-rules |
| Stoprealtimesnapshot | This interface is used to terminate an ongoing real-time snapshotting task. Users need to provide the unique identifier `persistentId` to specify the task. Upon successful invocation, the API returns the overall response code and message. | PUT | /api/v2/streams/snapshotings/* |
| Deletescreenshotrules | This interface is used to delete existing screenshot rules. You must specify the unique ID of the rule to be deleted. Please note that deleting currently executing rules is prohibited. The interface returns a success status code upon success and corresponding error codes otherwise. | DELETE | /api/v2/streams/snapshot-rules/* |
| Startrealtimesnapshot | This API is used to snapshot live streams in real-time. Users can select the snapshotting target by specifying the pull domain, publishing point, stream name, and snapshot template. Upon successful invocation, the API will return a snapshotting task ID and the snapshotting status for each stream. If premature termination of snapshotting is required, the "Stop Real-Time Snapshotting" API can be called; otherwise, the snapshotting task will automatically terminate after the live stream concludes. | POST | /api/v2/streams/snapshotings |
| Addsnapshotrules | This API allows you to add screenshot rules for live streams. It supports setting rules at the granularity of push domain, publishing point, and stream name. Users need to provide parameters such as pull stream domain, publishing point, domain, whether enabled, template ID, and stream name. When the corresponding rule is triggered, the system will take a screenshot of the video based on the associated screenshot template and return the generated rule ID. | POST | /api/v2/streams/snapshot-rules |