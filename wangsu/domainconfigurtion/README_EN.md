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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/domainconfigurtion"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &domainconfigurtion.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &domainconfigurtion.{ActionName}Response{}
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
| Batchupdateapidomainservice | Batch update domains' config. | POST | /api/domain/batchupdate |
| Updatedomainsrcstrategyforwplus | Set the back-to-origin configuration policy of the domain, including: advanced source, Rangel back-to-origin, 301 and 302 back-to-origin follow. | POST | /api/domain/setsrcconfig |
| Querydomain | query domain configuration. | GET | /api/si/domain |
| Editdomainconfig | To modify configuration of the specified domain; both domain name and domain name ID are supported. | PUT | /api/domain/* |
| Predeployapidomainservice | To pre-deploy configuration of certain acceleration domain name; both domain name and domain name ID are supported. | POST | /api/domain/predeploy/* |
| Querycdnwcontractdomains | This interface is used to query the association between Korea Division contracts and domains. The response includes a list of contract items, where each item contains the order item ID, contract ID, associated domain list, and customer code. | GET | /api/cdnw/contract/domains |
| Querycdnwcontractdomainsbycustomer | This interface is used to query the contract information and associated domain names for a specific customer under the Korea branch. By providing the item ID, contract ID, and customer code, users can retrieve a list of all domain names under the contract and related contract details. | POST | /api/cdnw/contract/domain |
| Queryhwantihotlinkingconfig | Queryhwantihotlinkingconfig | GET | /api/config/huaweicontrol/* |
| Updatetimecontrolservice | Modify timestamp anti-theft chain custom configuration item content | PUT | /api/config/timecontrol/* |
| Querytimecontrolservice | Query Time anti-theft chain custom configuration item content | GET | /api/config/timecontrol/* |
| Editdomainredirectconfig | Self-service through interface, url redirection function | PUT | /api/config/InnerRedirect/* |
| Edithttpheaderconfig | Through the interface self-implementation http header additions and deletions to the function, can be achieved in the cdn layer of personalized http header control, so that customers do not need to modify the source stationin this case, a custom http header and speedup are implemented. The interface * can be domain name or domain id. | PUT | /api/config/headermodify/* |
| Queryhttpheaderconfig | Query http header configuration via interface self - service. The interface * can be domain name or domain id. | GET | /api/config/headermodify/* |
| Editcachetimeconfig | The interface self-help implementation modifies the domain name cache time configuration, realizes the custom cache function according to the customer's request. Node cache is divided into regular cache and query string URl cache, where you can set the cache time and ignore certain headers that affect the cache, and whether to cache empty files, etc., can the query string Url be set to multiple or to cache the Url after removing the question mark (increasing hit rate) | PUT | /api/config/cachetime/* |
| Editantihotlinkingconfig | Modify the IP whitelist configuration under the domain name anti-theft chain through the interface self-service, and implement whitelist or blacklist control on the specific access IP. The interface * can be domain name or domain id. | PUT | /api/config/visitcontrol/* |
| Queryantihotlinkingconfig | Self-checking IP black and white list anti-theft chain configuration through interface. The interface * can be domain name or domain id. | GET | /api/config/visitcontrol/* |
| Updatesourceverificationconfig | Retweet back to the source with parameter authentication through the interface self-service, realize rtmp retweet back to the source, the edge can simulate the client with timestamp anti-leech parameters according to the encryption rules to retweet back to the source. | PUT | /api/config/sourceverification/* |
| Querysourceverificationconfig | Query retweet back to the source with parameter authentication configuration. | GET | /api/config/sourceverification/* |
| Updatedomainsrcinfo | Self-modify advanced source configurations through interfaces.Allows customers to specify regional configurations, different IP backsources, and when multiple sources are available, policy backsources can be specified. | PUT | /api/basicconfig/advancedsource/* |
| Updatecachebyresponseheaderconfig | The self-service implementation through the interface follows the source site caching rules. When there is a cache head, it is cached according to the source site caching head. When there is no cache head, it is not cached. | PUT | /api/config/cachebyrespheader/* |
| Querycachebyresponseheaderconfig | Self-query response header content caching rules through the interface. | GET | /api/config/cachebyrespheader/* |
| Querybanurlunderdomain | Get domain ban urls. | GET | /api/basicconfig/illegalinformation/* |
| Addbanurltodomian | The feature of URL blocking addition via API is now available. Customers can call the API to add URL blocking entries according to their own requirements. | PUT | /api/basicconfig/illegalinformation |
| Querypredeployresultfordomain | Query predeploy result for domain. You can query by the predeploy task id returned by the pre-deployment interface. The details of the pre-deployment result including machine IP, geography, and ISP information. After selecting the appropriate IP, configure the local host, host the domain to the corresponding machine IP, and perform configuration verification. | GET | /api/domain/predeploy/* |
| Predeploycachetimeconfig | Self-query node cache configuration configuration through interface | PUT | /api/predeploy/cachetime/* |
| Predeploysrcinfo | Self-modify advanced source configurations through interfaces.Allows customers to specify regional configurations, different IP backsources, and when multiple sources are available, policy backsources can be specified. | PUT | /api/predeploy/advancedsource/* |
| Predeployredirectconfig | View internal redirect configuration | PUT | /api/predeploy/InnerRedirect/* |
| Queryapideployservice | Requests for new, modified, enabled, disabled, cancelled, restored, deleted, etc. for domain names | POST | /api/request/* |
| Updatecompressionconfig | Modify compress setting configurations. | PUT | /api/config/compresssetting/* |
| Querycompressionconfig | Get compress setting configurations. | GET | /api/config/compresssetting/* |
| Predeploycompressionconfig | Get compress setting configurations. | PUT | /api/predeploy/compresssetting/* |
| Queryhttpcodecasheconfig | Get http code cache configurations. | GET | /api/config/httpcodecache/* |
| Queryipv6config | Get whether a domain uses ipv6 resources. | GET | /api/domain/ipv6/* |
| Queryaccessspeedconfig | Get access speed limitation configurations. | GET | /api/config/accessspeed/* |
| Queryafterredirect | Query the file after the pull jump | GET | /api/config/afterredirect/* |
| Updateafterredirect | Modify the pull jump after the file | PUT | /api/config/afterredirect/* |
| Updatetimecontrolservices | Modify Time anti-theft chain custom configuration item content | PUT | /api/config/timecontrols/* |
| Querysrcinfo | Query the customer advanced source configuration through the interface | GET | /api/basicconfig/advancedsource/* |
| Queryinnerredirect | View internal redirect configuration | GET | /api/config/InnerRedirect/* |
| Querycachetime | Self-service query of node cache configuration configuration through the interface. The * of the interface calling url can be the domain name or domain id. | GET | /api/config/cachetime/* |
| Editquerystringurlconfig | Modify query string setting configurations. | PUT | /api/config/querystring/* |
| Queryquerystringurlconfig | Through the interface, the query string settings can be modified by self-service, realizing the customized cache function. With the query string URL, you can set whether to cache multiple copies or cache the URL after removing the question mark (to increase the hit rate), and you can set whether to use the original request to return to the source, etc. The * of the interface calling url can be the domain name or domain id. | GET | /api/config/querystring/* |
| Updatetosauthorizationconfig | Modify tos access authorization configurations. The interface * can be domain name or domain id. | PUT | /api/config/tosaccess/* |
| Querytosauthorizationconfig | Get tos access authorization  configurations. The interface * can be domain name or domain id. | GET | /api/config/tosaccess/* |
| Edithttpcodecache | Modify http code cache configurations. | PUT | /api/config/httpcodecache/* |
| Queryamazons3authorizationconfig | Get amazon s3 access authorization  configurations. The interface * can be domain name or domain id. | GET | /api/config/amazons3access/* |
| Updateamazons3authorizationconfig | Modify amazon s3 access authorization configurations. The interface * can be domain name or domain id. | PUT | /api/config/amazons3access/* |
| Updatealiyunossauthorizationconfig | Modify Aliyun OSS access authorization configurations. The interface * can be domain name or domain id. | PUT | /api/config/ossaccess/* |
| Queryaliyunossauthorizationconfig | Get Aliyun OSS access authorization  configurations. The interface * can be domain name or domain id. | GET | /api/config/ossaccess/* |
| Editdomainproperty | Modify domain name properties, such as back source IP, back source host, and back source port. | POST | /api/domain/property/* |
| Queryignoreprotocol | Protocol caching and push configuration are ignored through interface self-service queries | GET | /api/config/ignoreprotocol/* |
| Editignoreprotocol | Modify ignore protocol cache and push configuration. | PUT | /api/config/ignoreprotocol/* |
| Editoriginuriandhost | Modify the source URI and host rewrite | PUT | /api/config/originrulesrewrites/* |
| Queryoriginuriandhost | Query back source URI and host rewrite | GET | /api/config/originrulesrewrites/* |
| Editwebsocketconfig | Self-modify the websocket switch configuration through the interface. The interface * can be domain name or domain id. | PUT | /api/config/websocket/* |
| Querywebsocketconfig | Self-query the websocket switch configuration via the interface. The interface * can be domain name or domain id. | GET | /api/config/websocket/* |
| Deletebanurls | the interface of delete the illegal masking | DELETE | /api/basicconfig/illegalinformation |
| Batchaddillegalinformation | Batch add domain's illegal information url. | PUT | /api/config/Batchaddillegalinformation |
| Batchdelillegalinformation | Batch delete domain's illegal information url. | DELETE | /api/config/Batchdelillegalinformation |
| Updatesingletranscodingconfigforwplus | Update Single Transcoding Configuration. The interface * can be domain name or domain id. | PUT | /api/config/transcodes/* |
| Querysingletranscodingconfigforwplus | Query Single Transcoding Configuration. The interface * can be domain name or domain id. | GET | /api/config/transcodes/* |
| Updateglobaltranscodingconfigforwplus | Update Global Transcoding Configuration. The interface * can be domain name or domain id. | PUT | /api/config/transcodeswitch/* |
| Queryglobaltranscodingconfigforwplus | Query Global Transcoding Configuration. The interface * can be domain name or domain id. | GET | /api/config/transcodeswitch/* |
| Querylivestreamingantihotlinkingconfig | Self-query the live-visitcontrol configuration via the interface. The interface * can be domain name or domain id. | GET | /api/config/live-visitcontrol/* |
| Updatelivevisitcontrolconfig | Self-modify the live-visitcontrol configuration through the interface. The interface * can be domain name or domain id. | PUT | /api/config/live-visitcontrol/* |
| Querylivestreamingtimestampantihotlinkingconfig | Query streaming media timestamp visit control. The interface * can be domain name or domain id. | GET | /api/config/live-timestampvisitcontrol/* |
| Editlivestreamingtimestampantihotlinkingconfig | Update streaming media timestamp visit control. The interface * can be domain name or domain id. | PUT | /api/config/live-timestampvisitcontrol/* |
| PredeployedStreamingTimestampVisitControl | PredeployedStreamingTimestampVisitControl. The interface * can be domain name or domain id. | PUT | /api/predeploy/live-timestampvisitcontrol/* |
| QueryCloudStorageBasicConfiguration | QueryCloudStorageBasicConfiguration.The interface * can be domain name or domain id. | GET | /api/config/cloudstorage/* |
| UpdateCloudStorageBasicConfiguration | Update recording basic configuration.The interface * can be domain name or domain id. | PUT | /api/config/cloudstorage/* |
| QueryRecordingBasicConfiguration | QueryRecordingBasicConfiguration.The interface * can be domain name or domain id. | GET | /api/config/recording/* |
| UpdateRecordingBasicConfiguration | UpdateRecordingBasicConfiguration.The interface * can be domain name or domain id. | PUT | /api/config/recording/* |
| QueryScreenshotConfiguration | Query screenshot basic configuration.The interface * can be domain name or domain id. | GET | /api/config/screenshot/* |
| UpdateScreenshotConfiguration | Modify screenshot basic configuration.The interface * can be domain name or domain id. | PUT | /api/config/screenshot/* |
| Getbasicconfigurationofdomain | View the configuration of the specified domain. Both domain name and domain name ID are supported. | GET | /api/domain/* |
| Queryappadomainportinfoforwplus | Get APPA domain config. | GET | /api/domainlist/appa |
| QueryBanUrlByDomain | query the blocked url list under a specified account, whick calls this api | POST | /api/config/banurl |
| Queryhttp2settingsconfigforwplus | Queryhttp2settingsconfigforwplus. The interface * can be domain name or domain id. | GET | /api/config/http2/* |
| Updatehttp2settingsconfigforwplus | Updatehttp2settingsconfigforwplus. The interface * can be domain name or domain id. | PUT | /api/config/http2/* |
| Enableordisablewafprotection | Enable or disable WAF protection | PUT | /api/domain/wafsecurity |
| Enableordisablebotprotection | Enable or disable Bot protection | PUT | /api/domain/botsecurity |
| UpdateCacheKeyConfiguration | UpdateCacheKeyConfiguration.The interface * can be domain name or domain id. | PUT | /api/config/cachekey/* |
| QueryCacheKeyConfiguration | QueryCacheKeyConfiguration.The interface * can be domain name or domain id. | GET | /api/config/cachekey/* |
| Editaccessspeedlimit | Editaccessspeedlimit | POST | /api/config/accessspeed/* |
| Enableordisabledmsprotection | UpdateDomainDmsStatusServiceForWplus | PUT | /api/domain/dmssecurity |
| Editback2originprotocolrewriteconfig | Update Back2Origin Protocol Rewrite Config	 | PUT | /api/config/back2originrewrite/* |
| Queryback2originprotocolrewriteconfig | Update Back2Origin Protocol Rewrite Config. | GET | /api/config/back2originrewrite/* |
| Updatestreamnotificationconfig | Update Stream Notification Config (Normal) | PUT | /api/config/streamnotification/* |
| Querystreamnotificationconfig | Query Stream Notification Config (Normal) | GET | /api/config/streamnotification/* |
| Querydomainbillingarea | Get domain's billing areas. | GET | /api/domain/billingarea/* |
| Updatedomaincertconfig | Set domain's certificate config, support SNI only. | PUT | /api/config/certificate/* |
| Addappadomain | Add APPA domain. | POST | /api/domain/appa |
| Updateappadomain | Update APPA domain. | PUT | /api/domain/appa/* |
| Getappadomainconfig | Get APPA domain config. | GET | /api/domain/appa/* |
| Querydomaincertconfig | Get domain's certificate config, support SNI only. | GET | /api/config/certificate/* |
| Pushnotifyaddress | Push notify address | PUT | /api/notifyaddress |
| Queryappacarryclientipconfig | Query APPA carry clientip config for specified domain. | GET | /api/config/appacarryclientip/* |
| Updatedomainmulticertconfig | Set domain's certificate config, support SNI and mutil SNI. | PUT | /api/config/certificate/v2/* |
| Updateipversionconfig | Set IPv4 and IPv6 resource for your domain. | PUT | /api/config/ipversion/* |
| Querylivedomainorigins | Query Live Domain Origins | GET | /live/domains/*/origins |
| Updatelivedomainorigins | Update Live Domain Origins | PUT | /live/domains/*/origins |
| Querylivedomainpagination | Query Live Domain Pagination | GET | /live/domains |
| Querylivedomaintls | Query Live Domain TLS | GET | /live/domains/*/tls |
| Updatelivedomaintls | Update Live Domain TLS | PUT | /live/domains/*/tls |
| Querylivedomainheaderrules | Query Live Domain Header Rules | GET | /live/domains/*/headerRules |
| Updatelivedomainheaderrules | Update Live Domain Header Rules | PUT | /live/domains/*/headerRules |
| Querylivedomainaccessctrls | Query Live Domain Access Ctrls | GET | /live/domains/*/accessCtrls |
| Updatelivedomainaccessctrls | Update Live Domain Access Ctrls | PUT | /live/domains/*/accessCtrls |
| Querylivedomainhls | Query Live Domain HLS Config | GET | /live/domains/*/hls |
| Updatelivedomainhls | Update Live Domain HLS Config | PUT | /live/domains/*/hls |
| Querylivedomainmisc | Query Live Domain Misc Configuration | GET | /live/domains/*/misc |
| Updatelivedomainmisc | Update Live Domain Misc Configuration | PUT | /live/domains/*/misc |
| Querylivedomainpublishreports | Query Live Domain Publish Reports Configuration | GET | /live/domains/*/publishReports |
| Updatelivedomainpublishreports | Update Live Domain Publishing Reports Configuration | PUT | /live/domains/*/publishReports |
| Querylivedomaindetail | Query Live Domain Detail Config | GET | /live/domains/* |
| Addlivedomain | Add Live Domain | POST | /live/domains |
| Updatelivedomain | Update Live Domain Configurations | PUT | /live/domains/* |
| Querydomainrangefollowconfig | Query the back-to-origin configuration policy of the domain, including: angel back-to-origin, 301 and 302 back-to-origin follow. | GET | /api/domain/range-follow-config/* |
| Updatevariableconfig | Update variable config. The interface * can be domain name or domain id.<br> | PUT | /api/config/variable/* |
| Queryvariableconfig | Query variable config. The interface * can be domain name or domain id. | GET | /api/config/variable/* |
| Querydomaindraftconfig | The interface self-help implementation query the domain name draft configuration. The user obtains all valid draft configurations, including the operations of adding, modifying, and deleting draft configurations, by specifying the domain/domain ID.Users can know the previously set draft configuration, which is convenient for subsequent updates or deployments. | GET | /api/domain/draft/config/* |
| Canceldomaindraftconfig | The interface self-help implementation cancel the domain name draft configuration. The user cancels the draft configuration of the specified domain/domain ID, including all additions, modifications, and deletions of the previous configuration. It is used to clear unnecessary draft configurations. | PUT | /api/domain/draft/cancel/* |
| Deploydomaindraftconfig | The interface self-help implementation deploy the domain name draft configuration. The user deploys the draft configuration of the specified domain/domain ID, including all the additions, modifications, and deletions of the previous configuration. It is used to initiate the deployment of the draft configuration. The x-cdn-request-id returned by the successful call can be used to query the deployment result. | PUT | /api/domain/draft/deploy/* |
| Updatedomaindraftconfig | The interface self-help implementation modifies the domain name draft configuration. The user sets the draft configuration for the specified domain/domain ID, including the addition, modification, and deletion of the configuration. The draft configuration will not be deployed immediately; it will be deployed through the deployment draft configuration interface after the user confirms it. | POST | /api/domain/draft/config/* |
| Batchupdatedomaincertconfig | The interface is used to modify the configuration certificates of multiple CDN acceleration domains. Users can modify the certificates associated with the domains by providing the certificate ID and the list of domain names to be modified. | PUT | /api/config/certificate/batch |
| Predeploydomaindraftconfig | The interface self-help implementation pre-deploy the domain name draft configuration. The user pre-deploys the draft configuration of the specified domain, including all the additions, modifications, and deletions of the previous configuration. It is used to initiate the pre-deployment of the draft configuration. The preDeployI returned by the successful call can be used to query the pre-deployment result. | PUT | /api/domain/draft/predeploy/* |
| Querytranscodetemplatevideocodecs | This API allows querying supported transcoding video codec formats. Users can use this API to get information about supported transcoding video codec formats, including default values and supported video codecs. | GET | /api/template/transcodes/video-codecs |
| Querytranscodetemplatevideopresets | The API is used to query the supported transcoding preset modes. Users can obtain the supported transcoding preset modes via this API, including default values and Supported video presets. | GET | /api/template/transcodes/video-presets |
| Querytranscodetemplateprofiles | The API is used to query the supported transcoding profile modes. Users can obtain the supported transcoding profile modes through this API, including default values and supported profiles. | GET | /api/template/transcodes/profiles |
| Querytranscodetemplateaudiocodecs | This API facilitates querying the supported transcoding audio codecs. Users can retrieve the supported transcoding audio codecs through this interface, including default values and supported audio codecs. | GET | /api/template/transcodes/audio-codecs |
| Querydomainversionconfig | Query the domain configuration for a specified domain and version. Users can input the domain and version to query specific configurations, including domain name, cname, and configuration details. This mainly allows users to understand the configuration of a specified version. | POST | /api/domain-config/version |
| Querydomainversion | Query the deployment version of a specified domain. Users can input a domain to query the deployment version, including version list information. This mainly allows users to obtain the versions they have deployed. | GET | /api/domain-version/* |
| Rollbackdomainconfig | Rollback the domain configuration of the specified domain to the specified version. Users can input the domain and version rollback configuration. The main purpose is to allow users to rollback the configuration to a specified version. | POST | /api/rollback-domain-config |
| Batchupdateapidomainforwplus | This API is used to batch modify the basic configuration of specified acceleration domains. Users need to pass a list of domainConfigs in the request body, each element containing the `domain-name` of the domain to be modified, as well as optional other domain configurations. | PUT | /api/batch/domain |
| Querydomainconfigbydomainnames | This interface is used to query detailed configuration information for a list of domain names provided by the user. Its main functions include retrieving the domain's top-level domain, status, SNI certificate configuration, and origin protocol settings. Users need to provide the list of domain names to be queried in the request body. | POST | /api/domain/config |