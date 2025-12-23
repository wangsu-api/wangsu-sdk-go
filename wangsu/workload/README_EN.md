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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/workload"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &workload.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := workload.{ActionName}Response{}
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
| Listdeployment | list all Deployments of user | GET | /apis/apps/v1/namespaces/*/deployments |
| Getdeployment | query details of Deployment | GET | /apis/apps/v1/namespaces/*/deployments/* |
| Createdeployment | create Deployment | POST | /apis/apps/v1/namespaces/*/deployments |
| Updatedeployment | replace Deployment | PUT | /apis/apps/v1/namespaces/*/deployments/* |
| Deletedeployment | delete Deployment | DELETE | /apis/apps/v1/namespaces/*/deployments/* |
| Listcrdresource | query list of crd resources | GET | /apis/custom/v1/namespaces/*/crdresources |
| Getcrdresource | query details of crd resource | GET | /apis/custom/v1/namespaces/*/crdresources/* |
| Createcrdresource | create crd resource | POST | /apis/custom/v1/namespaces/*/crdresources |
| Updatecrdresource | update crd resource | PUT | /apis/custom/v1/namespaces/*/crdresources/* |
| Deletecrdresource | delete crd resource | DELETE | /apis/custom/v1/namespaces/*/crdresources/* |
| Putpatchdeployment | partial update Deployment | PUT | /apis/apps/v1/namespaces/*/deployments/*/ws/patch |
| Listclustercrdresource | query crd resource list of cluster | GET | /apis/custom/v1/clustercrdresources |
| Createclustercrdresource | create crd resource of cluster | POST | /apis/custom/v1/clustercrdresources |
| Getclustercrdresource | query crd resource details of cluster | GET | /apis/custom/v1/clustercrdresources/* |
| Updateclustercrdresource | update crd resource of cluster | PUT | /apis/custom/v1/clustercrdresources/* |
| Deleteclustercrdresource | delete crd resource of cluster | DELETE | /apis/custom/v1/clustercrdresources/* |
| Putpatchclustercrdresource | partial update crd resource of cluster | PUT | /apis/custom/v1/clustercrdresources/*/ws/patch |
| Putpatchcrdresource | partial update crd resource | PUT | /apis/custom/v1/namespaces/*/crdresources/*/ws/patch |
| Getwsproxy | query edge cluster resources | GET | /apis/custom/v1/namespaces/*/wsproxy/* |
| Listpod | listPod | GET | /openapi/custom/api/v1/pods/list |
| Getcontainerlog | Query the list of poolclusters for which the user has permissions | GET | /api/v1/namespaces/*/pods/*/log |
| Liststatefulset | list statefulset | GET | /apis/apps/v1/namespaces/*/statefulsets |
| Getstatefulset | get statefulset detail | GET | /apis/apps/v1/namespaces/*/statefulsets/* |
| Createstatefulset | create statefulset | POST | /apis/apps/v1/namespaces/*/statefulsets |
| Updatestatefulset | update statefulset | PUT | /apis/apps/v1/namespaces/*/statefulsets/* |
| Deletestatefulset | delete statefulset | DELETE | /apis/apps/v1/namespaces/*/statefulsets/* |
| Listevents | query list of events | GET | /openapi/custom/api/v1/events |
| Getresourcestatus | get resource status | GET | /openapi/custom/api/v1/common/resource/status |
| Deletepod | delete  pod | POST | /openapi/custom/api/v1/pods/delete |
| Listjob | list job on namespace | GET | /apis/batch/v1/namespaces/*/jobs |
| Getjob | get job detail on namespace | GET | /apis/batch/v1/namespaces/*/jobs/* |
| Createjob | create job on namespace | POST | /apis/batch/v1/namespaces/*/jobs |
| Updatejob | update job on namespace | PUT | /apis/batch/v1/namespaces/*/jobs/* |
| Deletejob | delete job on namespace | DELETE | /apis/batch/v1/namespaces/*/jobs/* |
| Listpodfromedge | listPodFromEdge | GET | /openapi/custom/api/v1/pods/edge |
| Getpodfromedge | Get pod list information from the edge | GET | /openapi/custom/api/v1/namespaces/*/pods/* |
| Listpodevent | listPodEvent | GET | /openapi/custom/api/v1/pods/events |
| Patchjob | patch job on namespace | PATCH | /apis/batch/v1/namespaces/*/jobs/* |
| Patchstatefulset | patch statefulset on namespace | PATCH | /apis/apps/v1/namespaces/*/statefulsets/* |
| Getcronjob | get cronjob detail on namespace | GET | /apis/batch/v1/namespaces/*/cronjobs/* |
| Listcronjob | list cronjob on namespace | GET | /apis/batch/v1/namespaces/*/cronjobs |
| Createcronjob | create cronjob on namespace | POST | /apis/batch/v1/namespaces/*/cronjobs |
| Updatecronjob | update cronjob on namespace | PUT | /apis/batch/v1/namespaces/*/cronjobs/* |
| Patchcronjob | patch cronjob on namespace | PATCH | /apis/batch/v1/namespaces/*/cronjobs/* |
| Deletecronjob | delete cronjob on namespace | DELETE | /apis/batch/v1/namespaces/*/cronjobs/* |
| Getephemeralcontainers | query details of Pod's EphemeralContainers | GET | /api/v1/namespaces/*/pods/*/ephemeralcontainers |
| Updateephemeralcontainers | update EphemeralContainers of Pod | PUT | /api/v1/namespaces/*/pods/*/ephemeralcontainers |