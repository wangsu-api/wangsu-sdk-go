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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/certificateapplication"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &certificateapplication.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &certificateapplication.{ActionName}Response{}
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
| Createcertificateapplicationorder | Create the certificate applying order. | POST | /api/certificate/order/create |
| Listcertificateapplicationorders | Get List of Certificate Applying Orders. | POST | /api/certificate/order/list |
| Getdcvcontent | Get the validate content for domain control validation. | POST | /api/certificate/order/domain/validate/info |
| Cancelcertificateapplicationorder | Cancel the certificate applying order. | POST | /api/certificate/order/cancel |
| Getcertificateapplicationorder | query certificate sale order detail info | POST | /api/certificate/order/detail |
| Reissuecertificateforwplus | This interface is used for reissuing certificates. You can reissue the certificate by providing the certificate ID, certificate description, certificate algorithm, verification method, whether it is automatically verified, whether it is automatically deployed, common name, and subject alternate name. When the call is successful, the interface will return the sales order ID. | POST | /api/certificate/reissue |
| Createcertificateapplicationorderforterraform | Creates a certificate application order specifically for the Terraform scenario. | POST | /api/terraform/certificate/order/create |
| Listcertificateapplicationordersforterraform | Retrieves a list of certificate application orders for the Terraform scenario. | POST | /api/terraform/certificate/order/list |
| Getcertificateapplicationorderforterraform | This interface is used to query certificate application order details in the Terraform scenario. | POST | /api/terraform/certificate/order/detail |
| Cancelcertificateapplicationorderforterraform | Cancel the certificate applying order in the Terraform scenario. | POST | /api/terraform/certificate/order/cancel |
| Batchgetdcvcontent | This interface is used to batch query the validation content required by CA organizations for domain control validation (DCV). Users need to provide a list of purchase record IDs or domain names for the query, and the interface will return the corresponding certificate domain validation information. | POST | /api/certificate/orders/domain/validate/info |
| Querycertificateapplicationandissuancerecorddetail | This API is used to query the detailed information of certificate application and issuance records based on the certificate order ID, including basic attributes of the certificate, order status, contact information, etc. | POST | /api/certificate/order/record/detail |
| Querycertificateapplicationissuancerecordslist | This API is used to query all issuance records of certificate applications through the sales order ID. Users need to provide the order ID and can optionally fill in pagination parameters (such as page number and page size) to obtain a list of records within the specified range. | POST | /api/certificate/order/record/list |