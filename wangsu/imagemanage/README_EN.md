# Wangsu SDK for Go

This README provides documentation for using the Wangsu SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/wangsu-api/wangsu-sdk-go
```

## Product Single Installation

```bash
go get github.com/wangsu-api/wangsu-sdk-go/wangsu/imagemanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/imagemanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &imagemanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := imagemanage.{ActionName}Response{}
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
| Vmpqueryimage | Query the list of images that users can use. The list of image resources displayed includes user-defined images and public images provided by the edge computing platform. | GET | /vmp/images |
| Vmpcreateimage | You can create a virtual machine system disk as a mirror, and then use it to create a new virtual machine. It is recommended to shut down the virtual machine or stop applications or services on the virtual machine during the production of the image to avoid affecting the integrity of the image data. After the image production is completed, restart the virtual machine and its applications. The image created by this type of operation is returned as SNAPSHOT in the image query interface, representing the image created using a virtual machine snapshot. | POST | /vmp/images |
| Vmpremoveimage | Delete your custom image.<br>Deleting the image does not affect the already created virtual machine, but it cannot be used to create new virtual machines in the future. <br>You can only delete custom images created by yourself, and other customers' custom images and public images cannot be deleted. | DELETE | /vmp/images/* |
| Deployvmpimagepreheating | Used to extract private images of preheating customers. This interface is asynchronous, and the image preheating results need to be queried separately. | PUT | /vmp/images/preHeating |
| Queryvmpimagepreheatingstate | Used to query the image preheating status. | GET | /vmp/images/preHeatingInfo/* |
| Createharborproject | create harbor project | POST | /openapi/custom/v1/harbors/project |
| Updateharborproject | update harbor project | PUT | /openapi/custom/v1/harbors/project |
| Deleteharborproject | delete harbor project | DELETE | /openapi/custom/v1/harbors/project/* |
| Getharborproject | get harbor project detail | GET | /openapi/custom/v1/harbors/project/* |
| Listharborproject | list harbor projects | GET | /openapi/custom/v1/harbors/project |
| Createharboruser | create harbor user | POST | /openapi/custom/v1/harbors/user |
| Resetharboruserpwd | reset harbor user password | PUT | /openapi/custom/v1/harbors/user |
| Deleteharboruser | delete harbor user | DELETE | /openapi/custom/v1/harbors/user/* |
| Getharboruser | get harbor user detail | GET | /openapi/custom/v1/harbors/user/* |
| Listharboruser | list harbor user | GET | /openapi/custom/v1/harbors/user |
| Getimagedetail | get imageidetail and tags list | GET | /openapi/custom/v1/harbors/images/detail |
| Listimagepulljob | list all ImagePullJobs of user | GET | /apis/apps.kruise.io/v1alpha1/namespaces/*/imagepulljobs |
| Getimagepulljob | query details of ImagePullJob | GET | /apis/apps.kruise.io/v1alpha1/namespaces/*/imagepulljobs/* |
| Createimagepulljob | create ImagePullJob | POST | /apis/apps.kruise.io/v1alpha1/namespaces/*/imagepulljobs |
| Updateimagepulljob | replace ImagePullJob | PUT | /apis/apps.kruise.io/v1alpha1/namespaces/*/imagepulljobs/* |
| Deleteimagepulljob | delete ImagePullJob | DELETE | /apis/apps.kruise.io/v1alpha1/namespaces/*/imagepulljobs/* |
| Patchimagepulljob | partial update ImagePullJob | PATCH | /apis/apps.kruise.io/v1alpha1/namespaces/*/imagepulljobs/* |
| Deleteimagetag | delete image tag | DELETE | /openapi/custom/v1/harbors/images/tag |
| Createoemimage | Create OEM Image | POST | /ephone/oemImages/create |
| Queryoemimage | Queries the list of oem image. | GET | /ephone/oemImages/list |
| Manageoemimage | Specify the EdgeCloud phone OEM image to operate, currently only supports the delete operation | POST | /ephone/oemImages/action |