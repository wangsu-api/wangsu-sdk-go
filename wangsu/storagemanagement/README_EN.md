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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/storagemanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &storagemanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &storagemanagement.{ActionName}Response{}
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
| Listpvcs | query list of pvc | GET | /api/v1/namespaces/*/persistentvolumeclaims |
| Createpvcs | create pvc | POST | /api/v1/namespaces/*/persistentvolumeclaims |
| Getpvcs | query details of pvc | GET | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Updatepvcs | update pvc | PUT | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Deletepvcs | delete pvc | DELETE | /api/v1/namespaces/*/persistentvolumeclaims/* |
| Putpatchpvcs | partial update pvc | PUT | /api/v1/namespaces/*/persistentvolumeclaims/*/ws/patch |
| Pagingpvcs | Get pvc paging list | GET | /openapi/custom/api/v1/persistentvolumeclaims |
| Pvcinnamespace | Get the pvc under the namespace | GET | /openapi/custom/api/v1/namespaces/*/persistentvolumeclaims |
| Deletepvcfromedge | Remove pvc directly from edge cluster | DELETE | /openapi/custom/api/v1/namespaces/*/persistentvolumeclaims/* |
| Liststorageclass | Get the storageClass list | GET | /openapi/custom/api/v1/storageclasses |
| Querychannelrecordfiles | Query channel recording files | POST | /ivcs/report/origin/query-channel-record |
| Querydevicerecordfiles | Used to query the video files on the gb28181 device. The interface can only query data for up to 7 days at a time. To avoid device failure due to too many files, it is best to query by day. If the returned content is empty, it means that the device has no video files or there is an abnormality on the device. It is recommended to consult the device manufacturer. | POST | /ivcs/report/origin/query-device-record |
| Playbackcontrol | This interface can be used to set playback progress, playback ratio, playback pause, and playback resume. In order to improve user experience, try to use the webrtc stream pulling protocol when controlling playback to reduce data latency. | POST | /ivcs/report/origin/playback-control |
| Startplayback | Start playback of the specified device recording file. This interface sends a command to the device to start pushing the recorded file video stream. In order to ensure the stability of the device stream pushing bandwidth, a channel can only support one stream for playback. If you want to play back other files on the channel, you need to end the stream that is already playing before you can start playing back the new file. | POST | /ivcs/report/origin/start-playback |
| Stopplayback | End playback and send a signal to the device to stop pushing the video stream of the recorded file. Whether the stream is disconnected successfully can be judged through the callback information. | POST | /ivcs/report/origin/stop-playback |
| Queryplaybacklist | By Invoke this interface, you can query the stream that is currently playing back the device video. You can use the query interface to end the playback of the stream that does not need to be played back to avoid occupying the device bandwidth resources. | POST | /ivcs/report/origin/query-playback-list |