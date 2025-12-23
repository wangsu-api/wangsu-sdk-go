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
    "github.com/wangsu-api/wangsu-sdk-go/wangsu/reportflow"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportflow.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportflow.{ActionName}Response{}
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
| Flowchannel | Query the minute granularity traffic value of the channel | POST | /myview/Flowchannel |
| Wctquery | Query the 5-minute granularity statistics of various types and gradients of customer on-demand transcoding | POST | /myview/Wctquery |
| Flowday | Display channel traffic data by day, output date, peak time, peak bandwidth, total traffic. | POST | /myview/flow-day |
| Flowtype | Query the proportion of each type of traffic, including edge node traffic, origin traffic. | POST | /myview/flow-type |
| Querydirectorybandwidthtrafficunderlivestreamdomain | Query traffic or bandwidth statistics for each directory under the live push and pull streaming domains. Users need to provide the start and end time, domain name, and can specify a directory for querying. By default, the data is aggregated every 5 minutes. The returned results include the total traffic or bandwidth peak for each directory, along with detailed data for each time segment. | POST | /api/report/flow/dir/detail |
| Querystreamtrafficundermalbdomain | This interface is used to query the traffic or bandwidth data of live streaming under a domain name within a specified time period. The user needs to provide the time range, domain name and stream name, as well as data type and granularity options. The interface returns the traffic or bandwidth details of each stream under each domain name, including total traffic and peak bandwidth. It helps users monitor live streaming traffic to optimize resource allocation. | POST | /api/report/flow/stream/detail |
| Reportflowexactdomainservice | This interface is used to query the traffic data of precise domain names, especially for the traffic analysis of specific domain names corresponding to wildcard domain names. Users need to provide the time range and domain name to obtain traffic data. The returned results include the traffic data of each precise domain name in different time slices. It helps users understand the traffic situation of domain names and helps optimize domain name configuration and network resources. Note that the data is delayed by 3 to 5 hours, and it is recommended to query the data 24 hours ago. | POST | /api/report/flow/exact-domain |
| Queryoutputtrafficundershieldpop | Query the downstream traffic data of the Relay node of multiple domain . Users can obtain detailed Traffic data reports by providing the start time, end time, and domain list. The return value includes the total Traffic of each domain and provides the Traffic value of each time slice. This interface is suitable for monitoring and analyzing the Traffic profile and change trend of CDN Relay node. It's recommended that the call frequency is no higher than 30/5min. | POST | /api/report/flow/parent-node |
| Querymultiplestreamtrafficandbandwidthunderthedomain | This interface is used to query the traffic and bandwidth information of live streams under multiple domain names within a specified time period. Users can enter the start time, end time, and domain name to obtain detailed data, including the total traffic of each stream under each domain name, peak bandwidth, traffic or bandwidth data at each time point, and (optional) the number of online users at a time point. This is suitable for monitoring live stream network performance and user interaction, optimizing network configuration and resource allocation through data analysis, and improving the quality of live broadcast services. | POST | /api/report/flow-bandwidth/stream/detail |
| Reportflowp2pshareratioservice | Query the P2P share rate for multiple domains. Users provide a start time, end time, and domain list to get share rate data for each domain or aggregated statistics. Data can be returned in minute or five-minute granularity, including share rates for each time slice. It's ideal for monitoring and analyzing domain P2P usage efficiency. | POST | /api/report/flow/p2p/share-ratio |
| Querydatatransferforalldomainbyte | ReportFlowAllForByteService | POST | /api/report/flow/byte |
| Reportflowispprovinceshorttimeservice | Query of the traffic of multiple domains of each ISP in each province, only supports queries with a short time span. | POST | /api/report/flow/isp-province/shorttime |
| Queryedgebytehitratioservice | Query the hit rate of cache traffic at the edge node for multiple domains at the minute level | POST | /api/report/flow/edge-hit-ratio/total |
| Reportupflowdomaincountryservice | Query the upload traffic bandwidth distribution of a domain across countries and regions. Users input a time range and domain list to view data by domain, country, or domestic vs. foreign. Results include total region traffic, its percentage, and traffic and bandwidth per time segment, aiding in global website traffic analysis and management. Query data from 24 hours prior is recommended. | POST | /api/report/up-flow/domain-country |
| Gettotaltrafficforalldomains | This API is used to query the traffic summary information for all domains under a user's account. Users provide start and end time parameters to obtain detailed traffic data. Through this API, users can understand and analyze the CDN traffic usage of all domains within the specified time period. | GET | /api/report/traffic |
| Getbandwidthandtrafficbydirectoriesmainlandchinaonly | This API is used to query the bandwidth and total traffic for multiple subdirectories under domains, limited to the acceleration region of mainland China. Users need to provide the query time range and directory level information. The returned results include the total traffic and bandwidth peak value for each directory of the domain. This API helps users analyze website traffic distribution and bandwidth usage, aiding in resource optimization and service performance enhancement. | POST | /api/report/traffic/dir/info |
| Gettrafficandbandwidthbyclientcountry | Query the traffic bandwidth distribution of a domain across countries and regions. Users input a time range and domain list to view data by domain, country, or domestic vs. foreign. Results include total regional traffic, its percentage, and traffic and bandwidth per time segment, aiding in global website traffic analysis and management. Query data from 24 hours prior is recommended. | POST | /api/report/traffic/domain-country |
| Gettrafficandrequestshitratiobyispprovince1mingranularity | This API is used to query the minute-level traffic hit rate and request hit rate for specific provinces and ISPs. Users can obtain traffic or request hit rate data per minute by providing information such as start time, end time, domain name, province, and operator. This API is applicable for scenarios where network performance needs to be analyzed for different provinces and operators within a specified time period. Users can monitor the hit conditions of traffic and requests and optimize their network resource configuration through this API. | POST | /api/report/traffic/isp-province/hit-rate/total |
| Gettrafficbyispprovince1mingranularity | This API is used to query the summary of minute-level traffic for each province and ISP within a specified time period. Users can obtain detailed data summaries by specifying multiple domains, time intervals, ISPs, and other conditions. The API returns data including traffic, bandwidth, and request count. It is applicable for monitoring website traffic and bandwidth usage in specified provinces and ISPs to optimize network resources or identify access trends. | POST | /api/report/traffic/isp-province/total |
| Gettrafficrequestshitratiobyclientispprovince1mingranularity | This interface is used to query the minute-level traffic and request hit rate of the province and ISP dimensions, and supports returning Chinese or English information based on the request header Accept-Language. The user provides the time range, domain name, ISP, and province information. The returned results include the detailed traffic or request hit rate of the ISP in each province. It helps users analyze the network performance of each region, thereby optimizing content distribution strategies and improving network service quality. | POST | /api/report/traffic/isp-province/hit-rate/detail |
| Gettrafficbybillingregionformultidomains | This API is used to obtain the traffic summary statistics for accelerated domain names in specified billing areas. Users need to provide the query time range, area code, and optionally domain names and data granularity. The returned results include total traffic summary and traffic data for each time period. This API is suitable for users who need to analyze traffic distribution and optimize network resource allocation, helping in billing area traffic management and cost control. | POST | /api/report/traffic/area |
| Gettraffichitratiobyispprovinceformultidomains | This interface is used to query the byte hit rate of multiple domain names between different operators and provinces. Users can obtain detailed data distribution by specifying query time range, domain name list, ISP, province and other parameters. The data includes the byte hit rate at each moment and related hit traffic information. This interface is suitable for users to monitor and analyze cache performance in different regions and different ISPs, so as to optimize network resource configuration and improve user experience. | POST | /api/report/traffic/hit-rate/isp-province |
| Gettrafficandbandwidthbyispprovinceformultidomains | This API is used to query the traffic and bandwidth data of multiple domain names in different ISPs and provinces. Users need to provide the query time range and domain name list, and can group by ISP and province. The returned content includes the traffic and bandwidth data of each domain name in each province and ISP. This API can help users analyze the network performance of different regions and ISPs, thereby optimizing resource allocation and improving service quality. | POST | /api/report/traffic/isp-province |
| Gettrafficbyispprovinceformultidomainsbytes | This API allows users to query the traffic data (in Bytes) for multiple domain service nodes belonging to various ISPs and provinces. Users need to provide the query time range and domain list, and can group results by ISP and province. The results include detailed information on traffic and bandwidth for each domain, suitable for analyzing network performance across different regions and ISPs to optimize resources and enhance service quality. | POST | /api/report/traffic/isp-province/byte |
| Getorigintrafficbyispprovince | This API provides back-to-origin traffic data for multiple domains by province and ISP. Users can group queries by domain, ISP, and province, and specify time range and data granularity. Results include detailed traffic and bandwidth information to analyze regional and ISP performance. | POST | /api/report/traffic/origin/isp-province |
| Gettrafficbyprotocol | This API is used to query traffic data for multiple domains under a specified transmission protocol, covering data from all edge nodes. Users must specify parameters such as domain, start and end time, and transmission protocol, and can select data granularity and grouping dimensions. By default, the query returns HTTPS protocol traffic data, with output showing traffic values for each domain at different time points. This interface is suitable for scenarios requiring the query and analysis of multi-domain traffic across various transmission protocols. | POST | /api/report/traffic/protocol |
| Gettrafficrequestsandpeakbandwidthformultidomains | This interface is used to query the traffic request count and peak bandwidth of multiple domain names. Users need to provide a time range and a domain name list to obtain the total traffic, total request count, peak bandwidth and its occurrence time of each domain name, as well as detailed traffic and request count for each time period. This helps users understand the access status of the website and effectively adjust strategies to meet different traffic requirements. | POST | /api/report/traffic-request |
| Querytotaltrafficformultidomains | This interface is used to obtain traffic summary information of multiple domain names. Users need to provide a time range and domain name, and the data granularity can be selected as five minutes, one hour, or one day. The returned content includes the total traffic and the traffic value of each time node. This interface helps users monitor and analyze the traffic of multiple domain names, so as to grasp the overall traffic trend and changes. | POST | /api/report/domaintraffic |
| Getorigintrafficandrequestsformultidomains | This interface is designed to query the back-to-source traffic and number of requests for multiple domain names. Users can obtain this data by entering information such as the start time, end time, and specified domain name. The interface returns the total traffic, total number of requests, request and bandwidth peaks for each domain name and their corresponding time points, and supports providing detailed traffic, request, and bandwidth data at a minute granularity. This feature helps users monitor and analyze website back-to-source traffic. | POST | /api/report/traffic-request/origin |
| Getbandwidthandrequestsbyipversionformultidomains | This API is used to query IPv6 and IPv4 network bandwidth or request count for multiple domains. Users can specify the start time, end time, and domain list to obtain the data. The API returns detailed statistics for each domain, including IPv4 and IPv6 bandwidth or request count within the specified time period. This API is suitable for users who need to monitor traffic request conditions for multiple domains. By analyzing the returned data, network resource allocation can be optimized and website performance can be improved. | POST | /api/report/traffic-request/multi-ip-version |
| Getbandwidthsavingratio | This API is used to query the amount of bandwidth savings over a specified time period. Users need to provide the start time, end time, and domain information to retrieve data. The returned data includes the average total bandwidth savings and the actual bandwidth savings per time slice. This API is particularly suitable for scenarios requiring monitoring and analysis of bandwidth savings resulting from CDN cache hits. | POST | /api/report/bandwidth/saving-bandwidth/total |
| Getcdnorigintraffic | This interface is used to query CDN origin traffic data for specified dimensions. Users can use this interface to obtain detailed channel origin traffic reports for specific customers, including output date, peak time, peak bandwidth, total traffic, and more. This provides customers with direct insight into their own origin traffic usage. | POST | /cdn/traffic/origin |