# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/devicemanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "ggithub.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/devicemanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &devicemanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := devicemanagement.{ActionName}Response{}
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
| Getstreamurl | Get the channel playback URL | POST | /ivcs/devicemanage/getStreamUrl |
| Editdevice | Edit device info | POST | /ivcs/devicemanage/editDevice |
| Getdeviceslist | Obtain the list of devices | POST | /ivcs/devicemanage/getDevicesList |
| Invitepush | Invitepush | POST | /ivcs/devicemanage/invitePush |
| Stoppush | Stoppush | POST | /ivcs/devicemanage/stopPush |
| Createdevice | Add new device | POST | /ivcs/devicemanage/createDevice |
| Deletedevice | Deleting a device | POST | /ivcs/devicemanage/deleteDevice |
| Querychannellist | Get channel list | POST | /ivcs/devicemanage/getChnList |
| Editchannelinfo | edit channel info | POST | /ivcs/devicemanage/editChnInfo |
| Ptzcontrol | ptz control | POST | /ivcs/devicemanage/Ptzcontrol |
| ChannelControl | Channel disable/enable | POST | /ivcs/devicemanage/channelControl |
| ChannelStatusUpdate | This interface is used to update channel information. You can use this interface to actively update the channel list and channel status of the device. It is mainly used when the device does not report to the platform when adding, deleting, or going online or offline, or the report fails. Users can use this interface to actively let the device update and report again. | POST | /ivcs/devicemanage/channelStatusUpdate |