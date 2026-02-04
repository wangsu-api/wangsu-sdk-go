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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/recording"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &recording.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &recording.{ActionName}Response{}
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
| Getrecordtasklistquery | Through the recording task query interface, you can query the recording tasks issued by the user. The query only includes the recording tasks issued by the customer from cloudv, and does not include the tasks of instant push recording | POST | /live/channelManage/recordTaskListQuery |
| Livevideoconcat | Record file merging | POST | /live/channelManage/liveVideoConcat |
| Livevideoconcatquery | Example Query record merging tasks | POST | /live/channelManage/liveVideoConcatQuery |
| Addrecordingparametertemplate | This interface is used to add a recording parameter template. Users can pre-create live recording-related parameter templates, which include recording parameters, callback notification parameters, and storage parameters. Upon successful creation, the interface will return the template ID. | POST | /api/v2/streams/recording-parameter-templates |
| Queryrecordingparametertemplate | This API is used to query all recording parameter templates created under a customer's account, or to query a specific template by its ID. The query result will return detailed configuration information of the template, including template name, storage configuration, and recording parameters. | GET | /api/v2/streams/recording-parameter-templates |
| Modifyrecordingparametertemplate | This interface is used to edit and modify the content of a recording parameter template. Users must specify the template ID and can submit parameters such as template name, cloud storage configuration (e.g., management URL, AK/SK, bucket name, file name, storage duration), callback notification URL, and specific recording parameters (e.g., stream pull timeout settings, file size, segment duration, audio/video stream handling, file format). Modification is prohibited when the template is currently in use. | PUT | /api/v2/streams/recording-parameter-templates/* |
| Deleterecordingparametertemplate | This interface is used to delete a recording parameter template. Users must specify the template to be deleted via the `templateId` path parameter. Deletion is prohibited if the template is currently in use or associated with recording rules. Upon successful deletion, the interface will return the operation result. | DELETE | /api/v2/streams/recording-parameter-templates/* |
| Modifyrecordingrules | This interface is used to modify existing recording rules. Users need to specify the rule ID and can selectively update the publishing point, domain name, enabled status, template ID or stream name, and extended parameters. The interface response will return the operation result. | PUT | /api/v2/streams/recording-rules/* |
| Stoprealtimerecord | This interface is used to terminate an ongoing real-time recording task. Users need to provide the unique identifier `persistentId` to specify the task. Upon successful invocation, the API returns the overall response code and message. | PUT | /api/v2/streams/recordings/* |
| Addrecordingrules | This interface allows you to add recording rules for live streams. It supports specifying push domains, publishing points, and stream name granularity rules. When a rule is triggered, the system will perform live recording based on the associated recording template. Upon successful creation, the corresponding rule ID will be returned. | POST | /api/v2/streams/recording-rules |
| Queryrecordingrule | This API allows you to query already created recording rules. It supports filtering by recording rule ID, domain, publishing point, and stream name. If no filtering conditions are specified, all rules will be queried by default. The response will return a list of matching recording rules, with each rule including detailed information such as its ID, creation time, publishing point, domain, and enabled status. | GET | /api/v2/streams/recording-rules |
| Deleterecordingrules | This interface is used to delete created recording rules. Users must specify the rule to be deleted using its rule ID (`ruleId`). Please note that deleting a rule that is currently being executed is prohibited. Upon successful execution, the interface will return the corresponding response code and information. | DELETE | /api/v2/streams/recording-rules/* |
| Startrealtimerecord | This API is used to record live streams in real-time. Users can select the recording target by specifying the pull domain, publishing point, stream name, and recording template. Upon successful invocation, the API will return a recording task ID and the recording status for each stream. If premature termination of recording is required, the "Stop Real-Time Recording" API can be called; otherwise, the recording task will automatically terminate after the live stream concludes. | POST | /api/v2/streams/recordings |