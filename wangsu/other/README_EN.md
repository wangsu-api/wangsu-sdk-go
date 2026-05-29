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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/other"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &other.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := &other.{ActionName}Response{}
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
| Attoverview | Query the request number and deny number. | POST | /myview/Attoverview |
| Bandwidthappa | Query the bandwidth of appa. | POST | /myview/Bandwidthappa |
| Existflowchannel | Query channels with traffic. | POST | /myview/Existflowchannel |
| Querydailylivetranscodingduration | Query the live transcoding length at 1 day granularity. | POST | /myview/LiveTranscodingPerDayV2 |
| Query5minlivetranscodingduration | Query the live transcoding length of 5 minutes granularity. | POST | /myview/LiveTranscodingV2 |
| Virtualchannel | Query the virtual domain list of Pan domains. | POST | /myview/Virtualchannel |
| Wafhit | Query the waf request number at 5 minutes granularity. | POST | /myview/Wafhit |
| BandwidthUploadPeakRanking | Query the upload bandwidth ranking by different channels' peak value | POST | /myview/BandwidthUploadPeakRanking |
| HitDeny | Query the denied request number at minute granularity. | POST | /myview/HitDeny |
| Reportdomainlistexistflowservice | Query of domains that have generated traffic within certain period. Around 5-15 minutes of data delay. It's recommended that the call frequency is no higher than 30/5min. | POST | /api/report/domainList/existFlow |
| Reportstreamlistservice | This interface is used to query the stream name list under the specified domain name, helping users obtain all the stream names of the domain name in mainland China within a certain period of time. Users need to provide the query time range and domain name information, and the maximum time span is 3 days. The returned result contains the stream name list of the queried domain name, which helps users manage and analyze live streams. | POST | /api/report/streamList |
| Reporturldlfinishservice | This interface is used to query the top 500 URLs from multiple domain names by the number of completed downloads. The user needs to provide the start time, end time, and a list of domain names to be queried. The interface will return each URL and its corresponding number of successful downloads, helping users understand which URLs were accessed in a specific period of time. By calling this interface, users can effectively identify popular content or abnormal behavior, and make corresponding decisions. | POST | /api/report/url/dl-finish |
| Queryliverecordingduration | Query the live recording length of minute granularity. | POST | /myview/RecordingTime |
| Channelvalue | To inquire about the total traffic of a channel, provide the necessary details such as the channel name, billing method, acceleration type, billing value, billing unit, and peak value. | POST | /myview/Channelvalue |
| Flowappaday | Query the appa traffic at day granularity. | POST | /myview/Flowappaday |
| Flowappachannel | Query appa traffic by domain, and output the total edge upstream, edge downstream, and the sum of edge upstream and edge downstream traffic as well as the peak bandwidth of each within the query period for every domain. | POST | /myview/Flowappachannel |
| Concurrentsession | Query the concurrency number of appa | POST | /myview/Concurrentsession |
| Clouddirectduration | Query the cloud directing duration. Output the total directing duration and duration at 1-minute granularity. | POST | /myview/Clouddirectduration |
| Queryddosmitigatedbandwidthbysuiteorproduct | DDoS mitigatedBandwidth for query | POST | /soc/api/report/QueryMitigatedBandwidth |
| Picprocessstatistics | Output the image processing times by domain at 5-minute granularity. | POST | /myview/PicFlowHit |
| Httpdnsstatistics | Output the customers's resolution times at 5-minute granularity | POST | /myview/HttpdnsStatistics |
| Bandwidthupload | Query the upload bandwidth. | POST | /myview/bandwidth-upload |
| Reportavgspeeddomainispprovinceservice | This API is used to query the average download speed of a specific domain in different time periods. Users can specify the time range, ISP, domain, and grouping dimension to obtain accurate download speed data. The returned data includes the average download speed, average response time, average first packet response time, and total response time every five minutes within the specified time, helping users understand the network performance of different regions and ISP. This interface is suitable for scenarios where you want to analyze website performance, diagnose network problems, or optimize user experience. | POST | /api/report/avg-speed/domain-isp-province |
| Querylivestreamstatus | The current interface non-real-time data content includes: stream status information, Bitrate information, frame rate information, number of online users, header information, and Domain bandwidth; the real-time data content includes: stream status information, Bitrate information, frame rate information, number of online users (excluding HLS), and header information. | POST | /api/quality/stream-status-statistic |
| Getlivestreampushingstatus | Get the frame rate, bit rate, frame loss rate and other information of all channels under the specified domain name (supports multiple domain names and stream names) at the specified time. | POST | /api/quality/frame-rate |
| Queryonlineviewercount | This interface provides access to the online count of live streaming domain names or channels. It supports only the GET request method. | POST | /QOSS/api/onlineViewers |
| Livebacksourcequality | Provides a live domain name or channel back to source streaming quality interface, the interface only supports GET request methods | GET | /QOSS/api/backSourceQuality |
| Streambanning | Query the channels' banning info. | POST | /api/Streambanning |
| Reportbandwidthlowdelayp2pservice | This interface is used to query the minute-level low-latency P2P bandwidth data of a domain name within a specified time period. Users need to provide the query time range and domain name to obtain bandwidth information. The returned content is detailed P2P bandwidth data per minute. It helps users monitor the domain name P2P bandwidth in real time. | POST | /api/report/bandwidth/low-delay/p2p |
| Reportonlinenumispprovinceservice | This interface queries the number of online users of live broadcasts in different operators and provinces, and supports returning Chinese or English results based on the request header Accept-Language. Users need to provide domain name, time range, operator and province information, and the returned data includes the number of online users in each time slice. This interface is suitable for network management and data analysis, helping enterprises understand user dynamics in different regions and operators to optimize network resources and performance. | POST | /api/report/online-num/isp-province/detail |
| Reportdomainstreamdurationservice | Query the push duration of multi-domain names and multi-stream names | POST | /api/report/domain/stream/duration |
| Reportdomainstreamhlsonlineservice | This interface is used to query the number of online users under a domain name and its stream name within a specified time period. Users can enter the start time, end time, domain name, and stream name to obtain detailed information. The returned content includes the number of online users for each stream, the number of streams, and the total number of online users. This has important practical value for real-time understanding of the viewing situation of online live broadcast platforms and user participation in streaming media services. | POST | /api/report/hls/online |
| Reportdomainuseragentservice | This interface is used to query the traffic and number of requests of different types of UAs for domain names in different time periods. Users can specify the query time range, domain name list, and query type (such as browser, mobile device brand, or operating system) to obtain relevant data. The data returned by the interface includes the traffic (in MB) and number of requests for each UA. This interface helps users analyze the usage of websites on different devices and optimize resource allocation. | POST | /api/report/domain/useragent |
| HttpTest | HTTP Detect | POST | /wstApi/tools/http-test |
| DnsTest | DNS Detect | POST | /wstApi/tools/dns-test |
| MtrTest | MTR Detect | POST | /wstApi/tools/mtr-test |
| Streamtrafficservice | This interface is used to query stream name bandwidth. Users can query data by specifying time, domain, region, publishing point, protocol, etc.<br>Data is returned by aggregating traffic dimensions based on region, domain, stream name, publishing point, protocol, etc. | POST | /api/gdp/stream-traffic-service |
| Querydomainresourcegroup | This API is used to query the resource group to which a domain belongs. Users can query related resource group information through the domain, including resource group name and resource group ID. | POST | /api/v1/dms/domainResource |
| Getcdnoriginiprequest | This interface is used to query the number of back-to-origin requests for a customer's origin IP address. Users can use this interface to query the number of back-to-origin requests for a specific origin , including output date, status code, and number of requests. This helps customers understand their own back-to-origin origin request usage.<br> | POST | /cdn/request/originip_request |