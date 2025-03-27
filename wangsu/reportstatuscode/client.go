package reportstatuscode

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryIPV6StatusOfeachISPandProvinceRequest struct {
  // {"en":"Start time: 
  // 	1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time); 
  //     2. No bigger than the current time. 
  //     3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 	1. the time format is 2016-12-02T10:00:00+08:00 
  //     2. End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  //     3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur. 
  //     4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support).", "zh_CN":"结束时间：
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name: 
  // 	1. The maximum number of domains is 200 by default (Technical Support Adjustment can be contacted); 
  //     2. Automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为200个（可联系技术支持调整）；
  // 2.自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  //     1. Support 5m (5 minute granularity), 1h (1 hour granularity); 
  // 	2. The default is 5m", "zh_CN":"数据粒度：
  // 1.支持5m（5分钟）、1h（1小时）
  // 2.不传默认5m。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province
  // 
  // 1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.", "zh_CN":"省份
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商：
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"IP type: 
  // 	1.The optional values are IPv6 and IPv4 
  // 	2.If let this parameter empty,it will query all IP type", "zh_CN":"IP类型：
  // 1.可选值为 IPV6、IPV4
  // 2.不传默认查询全部"}
  IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
  // {"en":"Grouped dimension: 
  //     1.Aggregation date by default. 
  //     2.the optional value is domain,province,isp,allow to send multi option 
  //     3.send the Grouped dimension represent the need to display details by their corresponding values.For example, when groupBy is isp, the ISP dimension needs to be displayed in detail. When an ISP is not passed, it represents an aggregate date and would not return the ISP node. Provinces and domains have the same logic. For example, by passing 'groupBy': ['domain', 'province'], the ISP node under ispData does not need to return. {domain:'www.aaaa.com','ispData': [{'isp','China Telecom','provinceData': [...]}]}", "zh_CN":"分组关键词：
  // 1.默认聚合展示；
  // 2.可选值为domain.province.isp，可传入多个值；
  // 3.传入关键词则代表需要按照关键词对应的值展示明细； 例如groupBy传入isp，则isp维度需要明细展示；当没有传递isp，则代表isp聚合展示，同时isp节点则不返回。其他province和domain相同逻辑。 例如：传递'groupBy':   ['domain','province']，则ispData下的isp节点无需返回。 { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp':   '中国电信', 'provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDateTo(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetIPType(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.IPType = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponse struct {
  Result []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponse) SetResult(v []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult) *QueryIPV6StatusOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult) SetIspData(v []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  StatusCodeData []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) SetStatusCodeData(v []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.StatusCodeData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData struct     {
  // {"en":"Status Code", "zh_CN":"状态码类型"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetStatusCode(v string) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetRequestData(v []*QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData struct     {
  // {"en":"Time granularity is 5m, the format is yyyy-MM-dd HH:mm", "zh_CN":"数据粒度为5分钟，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Requests of status code", "zh_CN":"状态码请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetValue(v string) *QueryIPV6StatusOfeachISPandProvinceQueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryIPV6StatusOfeachISPandProvincePaths struct {
}

func (s QueryIPV6StatusOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryIPV6StatusOfeachISPandProvinceParameters struct {
}

func (s QueryIPV6StatusOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryIPV6StatusOfeachISPandProvinceRequestHeader struct {
}

func (s QueryIPV6StatusOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryIPV6StatusOfeachISPandProvinceResponseHeader struct {
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowIpVersionStatusCodeServiceRequest struct {
  // {"en":"Starting time:
  // 	1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00); 
  // 	2. Cannot be greater than the current time 
  // 	3. Get up to the last six months (183 days) of data.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒)
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time: 
  // 	1. Time format 2016-12-02T10:00:00+08:00 
  // 	2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken. 
  // 	3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception 
  // 	4. Allow query maximum time interval: 1 days, that is, the difference between dateFrom and dateTo can't exceed 1 days (can contact technical support adjustment).", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"domain name: 
  // 	1. The maximum number of deliverable domain names is 20 by default (can be contacted by technical support); 
  // 	2. Automatically filter out illegal domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of legitimate domain names).", "zh_CN":"域名:
  // 1、可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2、自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"1.Selection:domain 2.If groupBy left empty, merge date of all domains", "zh_CN":"可选值:domain
  // 不传默认聚合所有频道数据"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowIpVersionStatusCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowIpVersionStatusCodeServiceRequest) SetDateFrom(v string) *ReportFlowIpVersionStatusCodeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceRequest) SetDateTo(v string) *ReportFlowIpVersionStatusCodeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceRequest) SetDomain(v []*string) *ReportFlowIpVersionStatusCodeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceRequest) SetGroupBy(v []*string) *ReportFlowIpVersionStatusCodeServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowIpVersionStatusCodeServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIpVersionStatusCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIpVersionStatusCodeServiceResponse) SetCode(v string) *ReportFlowIpVersionStatusCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceResponse) SetMessage(v string) *ReportFlowIpVersionStatusCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceResponse) SetData(v []*ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData) *ReportFlowIpVersionStatusCodeServiceResponse {
  s.Data = v
  return s
}

type ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData struct     {
  // {"en":"Domain. If merge date of all domains will not return domain", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData) SetDomain(v string) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData) SetStatusCodeDataList(v []*ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseData {
  s.StatusCodeDataList = v
  return s
}

type ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList) SetStatusCode(v string) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList) SetDetailList(v []*ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList struct     {
  // {"en":"timestamp,Returns the timestamp between the start time and end time. Time format: Hours: yyyy MM DD hh:00:00", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。
  // 时间格式:5分钟:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"IPv4 data", "zh_CN":"IPv4数据"}
  V4Value *string `json:"v4Value,omitempty" xml:"v4Value,omitempty" require:"true"`
  // {"en":"IPv6 data", "zh_CN":"IPv6数据"}
  V6Value *string `json:"v6Value,omitempty" xml:"v6Value,omitempty" require:"true"`
}

func (s ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetV4Value(v string) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.V4Value = &v
  return s
}

func (s *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetV6Value(v string) *ReportFlowIpVersionStatusCodeServiceReportFlowIpVersionStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.V6Value = &v
  return s
}

type ReportFlowIpVersionStatusCodeServicePaths struct {
}

func (s ReportFlowIpVersionStatusCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServicePaths) GoString() string {
  return s.String()
}

type ReportFlowIpVersionStatusCodeServiceParameters struct {
}

func (s ReportFlowIpVersionStatusCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowIpVersionStatusCodeServiceRequestHeader struct {
}

func (s ReportFlowIpVersionStatusCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowIpVersionStatusCodeServiceResponseHeader struct {
}

func (s ReportFlowIpVersionStatusCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIpVersionStatusCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeNodeOriginServiceRequest struct {
  // {"en":"Start time
  // 
  // 1.The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2.Cannot be greater than the current time
  // 3.Get up to the last six months (183 days) of data.", "zh_CN":"开始时间
  // 
  //         格式为yyyy-MM-ddTHH:mm:ss+08:00；
  //         必须大于当前时间-183 天，并且小于当前时间和dateTo；
  //         dateFrom和dateTo相差不能超过7天；
  //         dateFrom和dateTo要么都传递，要么都不传递；
  //         dateFrom和dateTo都未传递，则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;
  // 3.If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default;
  // 4.Allowable maximum time range for query: 7 day, means the period between dateFrom to dateTo should not exceed 7 day", "zh_CN":"结束时间
  //         格式为yyyy-MM-ddTHH:mm:ss+08:00；
  //         必须大于dateFrom；如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The default upper limit to domains that can be entered is 20 (Contact technical support to adjust);", "zh_CN":"域名，域名个数限制根据账号可调，默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:5m (5 minutes granularity)", "zh_CN":"数据粒度，5m：5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension:
  // Optional values:domain.The detailed data will be displayed according to the dimension", "zh_CN":"分组维度
  //         可选值为domain；
  //         有传入则按照该维度展示明细数据；"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportStatusCodeNodeOriginServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeNodeOriginServiceRequest) SetDateFrom(v string) *ReportStatusCodeNodeOriginServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceRequest) SetDateTo(v string) *ReportStatusCodeNodeOriginServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceRequest) SetDomain(v []*string) *ReportStatusCodeNodeOriginServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceRequest) SetDataInterval(v string) *ReportStatusCodeNodeOriginServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceRequest) SetGroupBy(v []*string) *ReportStatusCodeNodeOriginServiceRequest {
  s.GroupBy = v
  return s
}

type ReportStatusCodeNodeOriginServiceResponse struct {
  // {"en":"Result", "zh_CN":"结果"}
  Result []*ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeNodeOriginServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeNodeOriginServiceResponse) SetResult(v []*ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult) *ReportStatusCodeNodeOriginServiceResponse {
  s.Result = v
  return s
}

type ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Status code data", "zh_CN":"状态码数据"}
  StatusCodeOriginData []*ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData `json:"statusCodeOriginData,omitempty" xml:"statusCodeOriginData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult) SetDomain(v string) *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult) SetStatusCodeOriginData(v []*ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData) *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResult {
  s.StatusCodeOriginData = v
  return s
}

type ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Request data", "zh_CN":"请求数据"}
  RequestData []*ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData) SetStatusCode(v string) *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData) SetRequestData(v []*ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData) *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginData {
  s.RequestData = v
  return s
}

type ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData struct     {
  // {"en":"The format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range.The first time of the day was yyyy-MM-dd 00:05, and the last time was yyyy-MM-dd 24:00", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是yyyy-MM-dd 24:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Requests of status code", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData) SetTimestamp(v string) *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData) SetValue(v string) *ReportStatusCodeNodeOriginServiceReportStatusCodeNodeOriginServiceResponseResultStatusCodeOriginDataRequestData {
  s.Value = &v
  return s
}

type ReportStatusCodeNodeOriginServicePaths struct {
}

func (s ReportStatusCodeNodeOriginServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeNodeOriginServiceParameters struct {
}

func (s ReportStatusCodeNodeOriginServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeNodeOriginServiceRequestHeader struct {
}

func (s ReportStatusCodeNodeOriginServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeNodeOriginServiceResponseHeader struct {
}

func (s ReportStatusCodeNodeOriginServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeNodeOriginServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeOriginFailRateServiceRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.Cannot exceed current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4.Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.", "zh_CN":"结束时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：24小时(可联系技术支持调整)，即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1.The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3.Domain name exceeding limit, misstatement", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.域名超过上限，报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:  
  // 1.The default is 1m;
  // 2.Support 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度：
  // 1.不传默认1m
  // 2.支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportStatusCodeOriginFailRateServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeOriginFailRateServiceRequest) SetDateFrom(v string) *ReportStatusCodeOriginFailRateServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceRequest) SetDateTo(v string) *ReportStatusCodeOriginFailRateServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceRequest) SetDomain(v []*string) *ReportStatusCodeOriginFailRateServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceRequest) SetDataInterval(v string) *ReportStatusCodeOriginFailRateServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportStatusCodeOriginFailRateServiceResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeOriginFailRateServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeOriginFailRateServiceResponse) SetCode(v string) *ReportStatusCodeOriginFailRateServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceResponse) SetMessage(v string) *ReportStatusCodeOriginFailRateServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceResponse) SetData(v []*ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) *ReportStatusCodeOriginFailRateServiceResponse {
  s.Data = v
  return s
}

type ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData struct     {
  // {"en":"Time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of return requests", "zh_CN":"回源总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Number of source failure requests (non-2XX status codes)", "zh_CN":"回源失败请求数(非2XX状态码)"}
  FailRequest *string `json:"failRequest,omitempty" xml:"failRequest,omitempty" require:"true"`
  // {"en":"Return source failure rate, reserved 4 decimal places", "zh_CN":"回源失败率，保留4位小数"}
  FailRate *string `json:"failRate,omitempty" xml:"failRate,omitempty" require:"true"`
}

func (s ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) SetTimestamp(v string) *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) SetTotalRequest(v string) *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData {
  s.TotalRequest = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) SetFailRequest(v string) *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData {
  s.FailRequest = &v
  return s
}

func (s *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData) SetFailRate(v string) *ReportStatusCodeOriginFailRateServiceReportStatusCodeOriginFailRateServiceResponseData {
  s.FailRate = &v
  return s
}

type ReportStatusCodeOriginFailRateServicePaths struct {
}

func (s ReportStatusCodeOriginFailRateServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeOriginFailRateServiceParameters struct {
}

func (s ReportStatusCodeOriginFailRateServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeOriginFailRateServiceRequestHeader struct {
}

func (s ReportStatusCodeOriginFailRateServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeOriginFailRateServiceResponseHeader struct {
}

func (s ReportStatusCodeOriginFailRateServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeOriginFailRateServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributioninCountriesRequest struct {
  // {"en":"Start time
  // 
  // 1. The format is yyyyy-MM-ddTHH:mm:SS+08:00, for example, 2016-12-02T10:00+08:00 (10:00:00 Beijing time on December 2, 2016);
  // 2. can not exceed the current time;
  // 3. the latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 4. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust). ", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间;
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 4.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 1m: 1-minute  5m: 5-minute  granularity, 1h: 1-hour granularity, 1d: 1-day granularity", "zh_CN":"数据粒度,1m: 1分钟粒度, 5m:5分钟粒度,1h:1小时粒度,1d:1天粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Country code 1) By default, all countries  are queried; 2) For specific country and region names, please refer to the appendix chapter on the overview page).","zh_CN":"国家代号: 1) 没传返回所有国家; 2)具体国家和地区名称请参考概览页面的附录章节."}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Group dimension
  // 
  // 1.        Options   are domain, province, isp, and more than one value can be entered;
  // 
  // 2.        The   data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、country,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDateFrom(v string) *QueryStatusCodeDistributioninCountriesRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDateTo(v string) *QueryStatusCodeDistributioninCountriesRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDomain(v []*string) *QueryStatusCodeDistributioninCountriesRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDataInterval(v string) *QueryStatusCodeDistributioninCountriesRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetCountryCode(v []*string) *QueryStatusCodeDistributioninCountriesRequest {
  s.CountryCode = v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributioninCountriesRequest {
  s.GroupBy = v
  return s
}

type QueryStatusCodeDistributioninCountriesResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesResponse) SetCode(v string) *QueryStatusCodeDistributioninCountriesResponse {
  s.Code = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponse) SetMessage(v string) *QueryStatusCodeDistributioninCountriesResponse {
  s.Message = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponse) SetData(v []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData) *QueryStatusCodeDistributioninCountriesResponse {
  s.Data = v
  return s
}

type QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData) SetDomain(v string) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData) SetCountryData(v []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseData {
  s.CountryData = v
  return s
}

type QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData struct     {
  // {"en":"country code", "zh_CN":"国家代码"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"country name", "zh_CN":"国家名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  StatusCodeData []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData) SetCountryCode(v string) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData {
  s.CountryCode = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData) SetCountryName(v string) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData {
  s.CountryName = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData) SetStatusCodeData(v []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryData {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData struct     {
  // {"en":"status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData struct     {
  // {"en":"Time,
  // 1.        When   the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data   value of every time slice represents the data value within the previous time   granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM,   and the last one is (yyyy-MM-dd+1) 00:00;
  // 2.        When   the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value   of every time slice represents the data value within the previous time   granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and   the last one is (yyyy-MM-dd+1) 00;
  // 3.        Return   the time slice contained in start time and the time slice contained in end   time.", "zh_CN":"时间,
  // 1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status   code", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributioninCountriesQueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributioninCountriesPaths struct {
}

func (s QueryStatusCodeDistributioninCountriesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesPaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributioninCountriesParameters struct {
}

func (s QueryStatusCodeDistributioninCountriesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributioninCountriesRequestHeader struct {
}

func (s QueryStatusCodeDistributioninCountriesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributioninCountriesResponseHeader struct {
}

func (s QueryStatusCodeDistributioninCountriesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginStatusCodeDistributionRequest struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days;
  // 4dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天;
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is  20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: granularity of 5 minutes", "zh_CN":"数据粒度,5m:5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"Optional values 0, 1. Default is 0
  // Input parameter 1 returns the total number of requests to the source, and 0 only returns the number of requests to the source station", "zh_CN":"可选值 0, 1 。默认为 0
  // 入参 1 则返回全部回源请求数,入参 0 则只返回回源站请求数"}
  BacksrcOnly *int `json:"backsrcOnly,omitempty" xml:"backsrcOnly,omitempty"`
  // {"en":"Query dimension. Optional values: statusCode , statusCodeType. Default value is statuscode.
  // 1.statusCode: returns the status code details;
  // 2.statusCodeType: returns the requests of each status code type (such as the number of requests corresponding to success, redirect, not modified, permission, not found, server error, and other)", "zh_CN":"查询维度,可选值:statusCode, statusCodeType;不传默认statusCode
  // 1.statusCode :返回状态码明细;
  // 2.statusCodeType:返回状态码类型对应明细(如Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other对应的请求数)"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s QueryOriginStatusCodeDistributionRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionRequest) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDateFrom(v string) *QueryOriginStatusCodeDistributionRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDateTo(v string) *QueryOriginStatusCodeDistributionRequest {
  s.DateTo = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDomain(v []*string) *QueryOriginStatusCodeDistributionRequest {
  s.Domain = v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDataInterval(v string) *QueryOriginStatusCodeDistributionRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetGroupBy(v []*string) *QueryOriginStatusCodeDistributionRequest {
  s.GroupBy = v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetBacksrcOnly(v int) *QueryOriginStatusCodeDistributionRequest {
  s.BacksrcOnly = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetQueryBy(v string) *QueryOriginStatusCodeDistributionRequest {
  s.QueryBy = &v
  return s
}

type QueryOriginStatusCodeDistributionResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginStatusCodeDistributionResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionResponse) SetResult(v []*QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult) *QueryOriginStatusCodeDistributionResponse {
  s.Result = v
  return s
}

type QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other", "zh_CN":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other"}
  StatusCodeType *string `json:"statusCodeType,omitempty" xml:"statusCodeType,omitempty" require:"true"`
  // {"en":"statusCodeOriginData", "zh_CN":"回源状态码数据"}
  StatusCodeOriginData []*QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData `json:"statusCodeOriginData,omitempty" xml:"statusCodeOriginData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult) SetDomain(v string) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult) SetStatusCodeType(v string) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult {
  s.StatusCodeType = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult) SetStatusCodeOriginData(v []*QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResult {
  s.StatusCodeOriginData = v
  return s
}

type QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"requestData", "zh_CN":"数据"}
  RequestData []*QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) SetStatusCode(v string) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData {
  s.StatusCode = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) SetRequestData(v []*QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData {
  s.RequestData = v
  return s
}

type QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status  code", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) SetTimestamp(v string) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) SetValue(v string) *QueryOriginStatusCodeDistributionQueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData {
  s.Value = &v
  return s
}

type QueryOriginStatusCodeDistributionPaths struct {
}

func (s QueryOriginStatusCodeDistributionPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionPaths) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionParameters struct {
}

func (s QueryOriginStatusCodeDistributionParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionParameters) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionRequestHeader struct {
}

func (s QueryOriginStatusCodeDistributionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionResponseHeader struct {
}

func (s QueryOriginStatusCodeDistributionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponseHeader) GoString() string {
  return s.String()
}




type ReportFlvStatusCodeRealTimeOriginTotalServiceRequest struct {
  // {"en":"Start time: 
  // 	1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 	2.Cannot exceed current time 
  // 	3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 	1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 	2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time. 
  // 	3.If both fields of dataFrom and dateTo are left empty, the default query past 24 hours; If there is only one unsent, throw an exception 
  // 	4.Maximum allowed query time interval: 24 hours, Date from and dateTo, not more than 24 hours", "zh_CN":"结束时间:
  // 时间格式2016-12-02T10:00:00+08:00
  // 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 允许查询最大时间间隔:24小时(可联系技术支持调整),即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support). 
  // 	2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 可传递域名数量上限默认为20个(可联系技术支持调整);
  // 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 域名超过上限,报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity: 1. Default 1m 2. 1m (1 minute), 5m (5 minutes)
  // ", "zh_CN":"数据粒度:不传默认1m
  // 支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Query dimension:
  // Optional values: statusCode, 2XX, 3XX, 4XX, 5XX;
  // Default value is statusCode if not passed;
  // statusCode: Returns the detailed origin status code;
  // 2XX: Returns the detailed data of origin status codes starting with 2;
  // 3XX: Returns the detailed data of origin status codes starting with 3;
  // 4XX: Returns the detailed data of origin status codes starting with 4;
  // 5XX: Returns the detailed data of origin status codes starting with 5.", "zh_CN":"查询维度:
  // 可选值 statusCode, 2XX, 3XX, 4XX, 5XX;
  // 不传默认 statusCode;
  // statusCode :返回回源状态码明细;
  // 2XX : 返回各 2 开头回源状态码明细数据;
  // 3XX : 返回各 3 开头回源状态码明细数据;
  // 4XX : 返回各 4 开头回源状态码明细数据;
  // 5XX : 返回各 5 开头回源状态码明细数据;"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) SetDateFrom(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) SetDateTo(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) SetDomain(v []*string) *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) SetDataInterval(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest) SetQueryBy(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceRequest {
  s.QueryBy = &v
  return s
}

type ReportFlvStatusCodeRealTimeOriginTotalServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceResponse) SetCode(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceResponse) SetMessage(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceResponse) SetData(v []*ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData) *ReportFlvStatusCodeRealTimeOriginTotalServiceResponse {
  s.Data = v
  return s
}

type ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData) SetStatusCode(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData {
  s.StatusCode = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData) SetDetailList(v []*ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList) *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseData {
  s.DetailList = v
  return s
}

type ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList struct     {
  // {"en":" Time format: yyyy-MM-dd HH:mm", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"The total number of back-to-source requests", "zh_CN":"回源总请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList) SetTimestamp(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList) SetValue(v string) *ReportFlvStatusCodeRealTimeOriginTotalServiceReportFlvStatusCodeRealTimeOriginTotalServiceResponseDataDetailList {
  s.Value = &v
  return s
}

type ReportFlvStatusCodeRealTimeOriginTotalServicePaths struct {
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServicePaths) GoString() string {
  return s.String()
}

type ReportFlvStatusCodeRealTimeOriginTotalServiceParameters struct {
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceParameters) GoString() string {
  return s.String()
}

type ReportFlvStatusCodeRealTimeOriginTotalServiceRequestHeader struct {
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlvStatusCodeRealTimeOriginTotalServiceResponseHeader struct {
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlvStatusCodeRealTimeOriginTotalServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryISPProvinceStatusCodeRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.Cannot exceed current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4.Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.", "zh_CN":"结束时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：24小时(可联系技术支持调整)，即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1.The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3.If left blank, all domain names will be obtained. If the total number of domain names exceeds the upper limit, an error will be reported.", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.若未填写默认查询全部域名，全部域名超出域名上限报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:  
  // 1.The default is 1m;
  // 2.Support 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度：
  // 1.不传默认1m
  // 2.支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province
  // 
  // 1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.", "zh_CN":"省份
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商：
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"query dimensionality:
  // 1.Optional value StatusCode, 2XX, 3XX, 4XX, 5XX, No default statusCode
  // 2.StatusCode: returns the status code details;
  // 3.2XX:Returns the 2XX return status code summary and the 2 start return status code detail data;
  // 4.3XX:Returns the 2XX return status code summary and the 3 start return status code detail data;
  // 5.4XX:Returns the 2XX return status code summary and the 4 start return status code detail data;
  // 6.5XX:Returns the 2XX return status code summary and the 5 start return status code detail data;", "zh_CN":"查询维度:
  // 1.可选值 statusCode 2XX 3XX 4XX 5XX, 不传默认 statusCode
  // 2.statusCode ：返回请求状态码明细；
  // 3.2XX:返回 2XX 状态码汇总及各 2 开头状态码明细数据；
  // 4.3XX:返回 3XX 状态码汇总及各 2 开头状态码明细数据；
  // 5.4XX:返回 4XX 状态码汇总及各 4 开头状态码明细数据；
  // 6.5XX:返回 5XX 状态码汇总及各 5 开头状态码明细数据；"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional: domain, all, If it is empty, it defaults to returning by domain dimension;
  // If all is passed, merge and return according to the query domain name.", "zh_CN":"可选项：domain、all, 为空则默认为按domain维度返回;
  // 若传递all，则按查询域名合并返回"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryISPProvinceStatusCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeRequest) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeRequest) SetDateFrom(v string) *QueryISPProvinceStatusCodeRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetDateTo(v string) *QueryISPProvinceStatusCodeRequest {
  s.DateTo = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetDomain(v []*string) *QueryISPProvinceStatusCodeRequest {
  s.Domain = v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetDataInterval(v string) *QueryISPProvinceStatusCodeRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetProvince(v []*string) *QueryISPProvinceStatusCodeRequest {
  s.Province = v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetIsp(v []*string) *QueryISPProvinceStatusCodeRequest {
  s.Isp = v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetQueryBy(v string) *QueryISPProvinceStatusCodeRequest {
  s.QueryBy = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetGroupBy(v string) *QueryISPProvinceStatusCodeRequest {
  s.GroupBy = &v
  return s
}

type QueryISPProvinceStatusCodeResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceStatusCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponse) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeResponse) SetCode(v string) *QueryISPProvinceStatusCodeResponse {
  s.Code = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponse) SetMessage(v string) *QueryISPProvinceStatusCodeResponse {
  s.Message = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponse) SetData(v []*QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData) *QueryISPProvinceStatusCodeResponse {
  s.Data = v
  return s
}

type QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData) SetDomain(v string) *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData {
  s.Domain = &v
  return s
}

func (s *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData) SetStatusCodeDataList(v []*QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList) *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseData {
  s.StatusCodeDataList = v
  return s
}

type QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList struct     {
  // {"en":"StatusCode", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList) SetStatusCode(v string) *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList) SetDetailList(v []*QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of  requests", "zh_CN":"总请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) SetValue(v string) *QueryISPProvinceStatusCodeQueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList {
  s.Value = &v
  return s
}

type QueryISPProvinceStatusCodePaths struct {
}

func (s QueryISPProvinceStatusCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodePaths) GoString() string {
  return s.String()
}

type QueryISPProvinceStatusCodeParameters struct {
}

func (s QueryISPProvinceStatusCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeParameters) GoString() string {
  return s.String()
}

type QueryISPProvinceStatusCodeRequestHeader struct {
}

func (s QueryISPProvinceStatusCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryISPProvinceStatusCodeResponseHeader struct {
}

func (s QueryISPProvinceStatusCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeRealTimeEdgeServiceRequest struct {
  // {'en':'Start time:
  // 1. Start time: time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (December 2rd, 2016, 10:00 a.m., Beijing Time);
  // 2. Not greater than the current time
  // 3. The most recent half-year (183 days) data can be obtained', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年01月01日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 1. The time format is 2016-12-02T10:00:00+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1 hours will be queried by default; if one field is filled and one is left empty, then exception will occur.
  // 4. Maximum time range allowable for query: 1 hour, means the period between dateFrom to dateTo should not exceed 1 hour', 'zh_CN':'结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：1小时（可联系技术支持调整），即dateFrom和dateTo相差不能超过1小时。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain:
  // 1. The default upper limit of domains that can be entered is 20 (if you want to adjust, please, contact technical support);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3. Domain name exceeding limit, misstatement', 'zh_CN':'域名：
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 3.域名超过上限，报错'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'Data granularity:
  // 1. default 1m
  // 2. 1m (1 minute), 5m (5 minutes)', 'zh_CN':'数据粒度：不传默认1m
  // 1.支持1m（1分钟）、5m（5分钟）'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportStatusCodeRealTimeEdgeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDateFrom(v string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDateTo(v string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDomain(v []*string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDataInterval(v string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeEdgeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponse) SetResult(v []*ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult) *ReportStatusCodeRealTimeEdgeServiceResponse {
  s.Result = v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult struct     {
  // {'en':'statusCodeData', 'zh_CN':'状态码数据'}
  StatusCodeData []*ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult) SetStatusCodeData(v []*ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResult {
  s.StatusCodeData = v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData struct     {
  // {'en':'Status code', 'zh_CN':'状态码'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'requestData', 'zh_CN':'请求数数据'}
  RequestData []*ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) SetStatusCode(v string) *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) SetRequestData(v []*ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData {
  s.RequestData = v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData struct     {
  // {'en':'DateTime, the format is   yyyy-MM-dd HH:mm; the data value of every time slice represents the data   value within the previous time granularity range.', 'zh_CN':'时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Number of requests for status codes', 'zh_CN':'状态码对应的请求数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) SetTimestamp(v string) *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) SetValue(v string) *ReportStatusCodeRealTimeEdgeServiceReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type ReportStatusCodeRealTimeEdgeServicePaths struct {
}

func (s ReportStatusCodeRealTimeEdgeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeEdgeServiceParameters struct {
}

func (s ReportStatusCodeRealTimeEdgeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeEdgeServiceRequestHeader struct {
}

func (s ReportStatusCodeRealTimeEdgeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeEdgeServiceResponseHeader struct {
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowProtocolStatusCodeServiceRequest struct {
  // {"en":"Start time: 
  // 	1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 	2.Cannot exceed current time 
  // 	3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒)
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 	1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 	2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time. 
  // 	3.If both fields of dataFrom and dateTo are left empty, the default query past 1 day; If there is only one unsent, throw an exception 
  // 	4.Maximum allowed query time interval: 1 days, Date from and dateTo, not more than 1 days", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support). 
  // 	2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 1、可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2、自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"1.Selection:domain 2.If groupBy left empty, merge date of all domains", "zh_CN":"可选值:domain
  // 不传默认聚合所有频道数据"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowProtocolStatusCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolStatusCodeServiceRequest) SetDateFrom(v string) *ReportFlowProtocolStatusCodeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceRequest) SetDateTo(v string) *ReportFlowProtocolStatusCodeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceRequest) SetDomain(v []*string) *ReportFlowProtocolStatusCodeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceRequest) SetGroupBy(v []*string) *ReportFlowProtocolStatusCodeServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowProtocolStatusCodeServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowProtocolStatusCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolStatusCodeServiceResponse) SetCode(v string) *ReportFlowProtocolStatusCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceResponse) SetMessage(v string) *ReportFlowProtocolStatusCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceResponse) SetData(v []*ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData) *ReportFlowProtocolStatusCodeServiceResponse {
  s.Data = v
  return s
}

type ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData struct     {
  // {"en":"Domain. If merge date of all domains will not return domain", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData) SetDomain(v string) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData) SetStatusCodeDataList(v []*ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseData {
  s.StatusCodeDataList = v
  return s
}

type ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList) SetStatusCode(v string) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList) SetDetailList(v []*ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList struct     {
  // {"en":"timestamp,Returns the timestamp between the start time and end time. Time format: yyyy-MM-dd HH:mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。
  // 时间格式:
  // 5分钟:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"https", "zh_CN":"https数据"}
  HttpsValue *string `json:"httpsValue,omitempty" xml:"httpsValue,omitempty" require:"true"`
  // {"en":"http", "zh_CN":"http数据"}
  HttpValue *string `json:"httpValue,omitempty" xml:"httpValue,omitempty" require:"true"`
}

func (s ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetHttpsValue(v string) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.HttpsValue = &v
  return s
}

func (s *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetHttpValue(v string) *ReportFlowProtocolStatusCodeServiceReportFlowProtocolStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.HttpValue = &v
  return s
}

type ReportFlowProtocolStatusCodeServicePaths struct {
}

func (s ReportFlowProtocolStatusCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServicePaths) GoString() string {
  return s.String()
}

type ReportFlowProtocolStatusCodeServiceParameters struct {
}

func (s ReportFlowProtocolStatusCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowProtocolStatusCodeServiceRequestHeader struct {
}

func (s ReportFlowProtocolStatusCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowProtocolStatusCodeServiceResponseHeader struct {
}

func (s ReportFlowProtocolStatusCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolStatusCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributionOfeachISPandProvinceRequest struct {
  // {"en":"Start time
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. can not exceed the current time;
  // 3. the latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. the end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 4. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust). ", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间;
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 4.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: when domain is not passed: default is all domain names, maximum supported domain names are 20 (can be adjusted by contacting technical support)", "zh_CN":"域名：当domain没有传时:默认为全部域名,最大域名支持20个(可联系技术支持调整)	"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 1m: 1-minute 5m: 5-minute granularity, 1h: 1-hour granularity", "zh_CN":"数据粒度,1m: 1分钟粒度, 5m:5分钟粒度,1h:1小时粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province
  // 
  // 1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.", "zh_CN":"省份
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。
  // 
  // "}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商：
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension
  // 
  // 1.Options are domain, province, isp, and more than one value can be entered;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDateTo(v string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponse struct {
  Result []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponse) SetResult(v []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult) *QueryStatusCodeDistributionOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult) SetIspData(v []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  StatusCodeData []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) SetStatusCodeData(v []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData struct     {
  // {"en":"Return specific status code details such as 200, 201, 500, as well as aggregated 1XX, 2XX, 3XX, 4XX, 5XX, all, OTHERS. Return when values are available.", "zh_CN":"返回具体状态码明细如200,201,500，及聚合的1XX，2XX，3XX，4XX，5XX，all，OTHERS。有值时返回"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData struct     {
  // {"en":"Time,
  //                                                                     1.When the data query granularity is 1m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 																	2.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 																	3.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;
  // 																	4.Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间,
  //                                                                     1.查询的数据粒度为1m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 																	2.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 																	3.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)00;
  // 																	4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status code", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributionOfeachISPandProvinceQueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvincePaths struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceParameters struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceRequestHeader struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseHeader struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeLogServiceRequest struct {
  // {'en':'From date:
  // 
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00.
  // 
  //  2.Cannot exceed current time
  // 
  // 3.The most recent six-month (183 days) data are available.', 'zh_CN':'开始时间：
  // 
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 
  // 2.不能大于当前时间
  // 
  // 3.最多可获取最近半年（183天）的数据'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 
  // 3.If both fields of dataFrom and dateTo are left empty,  the default query past 1 day; If there is only one unsent, throw an exception
  // 
  // 4.Maximum allowed query time interval: 30 days, Date from and dateTo, not more than 30 days', 'zh_CN':'结束时间：
  // 
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间
  // 
  // 3.dateFrom，dateTo二者都未传，默认查询过去的1天；如仅有一个未传，抛异常
  // 
  // 4.允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天。（可联系技术支持调整）'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain:
  // 
  // 1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support).
  // 
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)', 'zh_CN':'域名：
  // 
  // 1、可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 
  // 2、自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'1.Selection:domain
  // 2.If groupBy left empty, merge date of all domains', 'zh_CN':'1. 可选值：domain
  // 2. 不传默认聚合所有频道数据'}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceRequest) SetDateFrom(v string) *ReportStatusCodeLogServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeLogServiceRequest) SetDateTo(v string) *ReportStatusCodeLogServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeLogServiceRequest) SetDomain(v []*string) *ReportStatusCodeLogServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeLogServiceRequest) SetGroupBy(v []*string) *ReportStatusCodeLogServiceRequest {
  s.GroupBy = v
  return s
}

type ReportStatusCodeLogServiceResponse struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceResponse) SetCode(v string) *ReportStatusCodeLogServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponse) SetMessage(v string) *ReportStatusCodeLogServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponse) SetData(v []*ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData) *ReportStatusCodeLogServiceResponse {
  s.Data = v
  return s
}

type ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData struct     {
  // {'en':'Domain. If merge date of all domains will not return domain', 'zh_CN':'域名，聚合全部域名数据不返回该字段'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData) SetDomain(v string) *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData) SetStatusCodeDataList(v []*ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList) *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseData {
  s.StatusCodeDataList = v
  return s
}

type ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList struct     {
  // {'en':'Status code', 'zh_CN':'状态码'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList) SetStatusCode(v string) *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList) SetDetailList(v []*ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList struct     {
  // {'en':'timestamp,Returns the timestamp between the start time and end time.Time format:
  //                                                                                  Hours: yyyy MM DD hh:00:00', 
  //                                                                                  'zh_CN':'时间片,返回开始时间和结束时间包含的时间片。
  //                                                                                  时间格式:
  //                                                                                  小时：yyyy-MM-dd HH:00:00'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Status code times', 'zh_CN':'状态码次数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) SetValue(v string) *ReportStatusCodeLogServiceReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList {
  s.Value = &v
  return s
}

type ReportStatusCodeLogServicePaths struct {
}

func (s ReportStatusCodeLogServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeLogServiceParameters struct {
}

func (s ReportStatusCodeLogServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeLogServiceRequestHeader struct {
}

func (s ReportStatusCodeLogServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeLogServiceResponseHeader struct {
}

func (s ReportStatusCodeLogServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest struct {
  // {"en":"Start time
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. can not exceed the current time;
  // 3. the latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. the end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 4. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust). ", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间;
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 4.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递视为全部域名，全部域名超出域名数量上限将报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 1m: 1-minute 5m: 5-minute granularity, 1h: 1-hour granularity", "zh_CN":"数据粒度,1m: 1分钟粒度, 5m:5分钟粒度,1h:1小时粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province
  // 1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;
  // 2.Province is upload: Send province code, multiple can be sent. Please refer to the appendix description section of the overview page for the provincial information code table.", "zh_CN":"省份
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 2.有传递province时：省份传code，可传多个。省份信息码表详见概览页附录说明章节"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;
  // 2.ISPs is upload: Send isp code, multiple can be sent. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商： 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 2.有传递isp时：运营商 传code，可传多个。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension
  // 
  // 1.Options are domain, province, isp, and more than one value can be entered;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"Dimension of status code query:
  // 1. true: Status code aggregation return, return values are 2XX, 3XX, 4XX, 5XX, ALL, OTHER.
  // 2. false: Return detailed data based on specific status codes. isStatusMerge defaults to false when not passed.", "zh_CN":"状态码查询维度:
  // 1.true：状态码聚合返回，返回2XX,3XX,4XX,5XX,ALL,OTHER
  // 2.false：按具体状态码明细数据返回。不传默认为false"}
  IsStatusMerge *string `json:"isStatusMerge,omitempty" xml:"isStatusMerge,omitempty"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetDateFrom(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetDateTo(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetDomain(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetDataInterval(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetProvince(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.Province = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetIsp(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.Isp = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.GroupBy = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest) SetIsStatusMerge(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequest {
  s.IsStatusMerge = &v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponse struct {
  Result []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponse) SetResult(v []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponse {
  s.Result = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult) SetDomain(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult) SetIspData(v []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResult {
  s.IspData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData) SetIsp(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData) SetProvinceData(v []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  StatusCodeData []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData) SetProvince(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData) SetStatusCodeData(v []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceData {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData struct     {
  // {"en":"Status codes, with options of: '200', '500' and 'Others'", "zh_CN":"状态码,取值可能为:'200'、'500'、'其他'等"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData struct     {
  // {"en":"Time,
  // 																	1.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 																	2.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;
  // 																	3.Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间,
  // 																	1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 																	2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)00;
  // 																	3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status code", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributionOfeachISPandProvinceByUserIPQueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPPaths struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPPaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPParameters struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequestHeader struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseHeader struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceByUserIPResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributionRequest struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于当前时间-183天，并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整)；
  // 4.dateFrom和dateTo要么都传递，要么都不传递；
  // 5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names:
  // 1.Domain number limits can be adjusted depending on different accounts. The default value is 20
  // 2.Query all domain names under account when this entry is not passed", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.未传递该入参时查询账号下所有域名"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 5m: granularity of 5 minutes", "zh_CN":"数据粒度，5m：5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain；
  // 2.有传入则按照该维度展示明细数据；"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"1.If 0 is added, the value is true or false, which is false by default
  // 2.When the value of data adding is true, data is added to time points without data
  // 3.When the dataPadding value is false, no treatment will be done", "zh_CN":"是否补0，取值为true或false，默认为false
  // 当dataPadding取值为true时，对没有数据的时间点填充数据，取值为0
  // 当dataPadding取值为false时，不做处理"}
  DataPadding *bool `json:"dataPadding,omitempty" xml:"dataPadding,omitempty"`
  // {"en":"Query dimension. Optional values: statusCode, statusCodeType. Default value is statuscode.
  // 1.statusCode: returns the status code details;
  // 2.statusCodeType: returns the requests of eachstatus code type (such as the number of requests corresponding to success, redirect, not modified, permission, not found, server error, and other)", "zh_CN":"查询维度，可选值：statusCode、statusCodeType；不传默认statusCode
  // 1.statusCode ：返回状态码明细；
  // 2.statusCodeType：返回状态码类型对应明细(如Success、Redirect、Not-Modified、Permission、Not-Found、Server Error、Other对应的请求数)"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s QueryStatusCodeDistributionRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionRequest) SetDateFrom(v string) *QueryStatusCodeDistributionRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDateTo(v string) *QueryStatusCodeDistributionRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDomain(v []*string) *QueryStatusCodeDistributionRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDataInterval(v string) *QueryStatusCodeDistributionRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributionRequest {
  s.GroupBy = v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDataPadding(v bool) *QueryStatusCodeDistributionRequest {
  s.DataPadding = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetQueryBy(v string) *QueryStatusCodeDistributionRequest {
  s.QueryBy = &v
  return s
}

type QueryStatusCodeDistributionResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionResponse) SetResult(v []*QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult) *QueryStatusCodeDistributionResponse {
  s.Result = v
  return s
}

type QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"statusCodeData", "zh_CN":"状态码数据"}
  StatusCodeData []*QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult) SetDomain(v string) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult) SetStatusCodeData(v []*QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResult {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"totalRequest", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other", "zh_CN":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other"}
  StatusCodeType *string `json:"statusCodeType,omitempty" xml:"statusCodeType,omitempty" require:"true"`
  // {"en":"requestData", "zh_CN":"请求数数据"}
  RequestData []*QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) SetTotalRequest(v string) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData {
  s.TotalRequest = &v
  return s
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) SetStatusCodeType(v string) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData {
  s.StatusCodeType = &v
  return s
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range.", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests for status codes", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributionQueryStatusCodeDistributionResponseResultStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributionPaths struct {
}

func (s QueryStatusCodeDistributionPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionPaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionParameters struct {
}

func (s QueryStatusCodeDistributionParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionRequestHeader struct {
}

func (s QueryStatusCodeDistributionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionResponseHeader struct {
}

func (s QueryStatusCodeDistributionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusAllServiceRequest struct {
}

func (s ReportStatusAllServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServiceRequest) GoString() string {
  return s.String()
}

type ReportStatusAllServiceResponse struct {
  // {'en':'statusCodeData', 'zh_CN':'状态码汇总'}
  StatusCodeDataList []*ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList `json:"status-code-data,omitempty" xml:"status-code-data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusAllServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusAllServiceResponse) SetStatusCodeDataList(v []*ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList) *ReportStatusAllServiceResponse {
  s.StatusCodeDataList = v
  return s
}

type ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"status-code,omitempty" xml:"status-code,omitempty" require:"true"`
  // {"en":"Number of requests of the status   code", "zh_CN":"状态码对应的请求数"}
  Hit *string `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
}

func (s ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList) SetStatusCode(v string) *ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList) SetHit(v string) *ReportStatusAllServiceReportStatusAllServiceResponseStatusCodeDataList {
  s.Hit = &v
  return s
}

type ReportStatusAllServicePaths struct {
}

func (s ReportStatusAllServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServicePaths) GoString() string {
  return s.String()
}

type ReportStatusAllServiceParameters struct {
  // {"en":"start time
  // 
  // The 1. format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. must be less than the current time and dateto;
  // 
  // The difference between 3.dateFrom and dateTo should not exceed 31 days (with technical support adjustment).
  // 
  // 4. we can only query data in the latest 2 years.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须小于当前时间和dateto；
  // 3.dateFrom和dateTo相差不能超过31天（可联系技术支持调整）；
  // 4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 
  // 1.        The   format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.        Must   be greater than datefrom; if it’s greater than the current time, then the   current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于datefrom；如果大于当前时间，则重新赋值为当前时间；"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
}

func (s ReportStatusAllServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServiceParameters) GoString() string {
  return s.String()
}

func (s *ReportStatusAllServiceParameters) SetDatefrom(v string) *ReportStatusAllServiceParameters {
  s.Datefrom = &v
  return s
}

func (s *ReportStatusAllServiceParameters) SetDateto(v string) *ReportStatusAllServiceParameters {
  s.Dateto = &v
  return s
}

type ReportStatusAllServiceRequestHeader struct {
}

func (s ReportStatusAllServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusAllServiceResponseHeader struct {
}

func (s ReportStatusAllServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusAllServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeUrlTopServiceRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2024-12-23T10:00 + 08:00 (10:00:00 Beijing time on December 23, 2024);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2024-12-23T10:00:00+08:00（为北京时间2024年12月23日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. The end time is greater than the start time.
  // 
  // 3. If the end time is greater than the current time, the current time is taken.
  // 
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days.", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain
  // 
  // Can only transfer one domain.", "zh_CN":"域名：
  // 
  // 只能传一个域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Status code
  // 
  // Support transmitting specific status codes, such as 201/202, or summarizing status codes 1XX/2XX/3XX/4XX/5XX", "zh_CN":"状态码
  // 
  // 支持传具体状态码，如201，202、或者是汇总状态码1XX/2XX/3XX/4XX/5XX"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Default top 200, maximum input 1000", "zh_CN":"不传默认top200，最多输入1000"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
}

func (s ReportStatusCodeUrlTopServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeUrlTopServiceRequest) SetDateFrom(v string) *ReportStatusCodeUrlTopServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceRequest) SetDateTo(v string) *ReportStatusCodeUrlTopServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceRequest) SetDomain(v string) *ReportStatusCodeUrlTopServiceRequest {
  s.Domain = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceRequest) SetStatusCode(v string) *ReportStatusCodeUrlTopServiceRequest {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceRequest) SetTop(v int) *ReportStatusCodeUrlTopServiceRequest {
  s.Top = &v
  return s
}

type ReportStatusCodeUrlTopServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeUrlTopServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeUrlTopServiceResponse) SetCode(v string) *ReportStatusCodeUrlTopServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceResponse) SetMessage(v string) *ReportStatusCodeUrlTopServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceResponse) SetData(v []*ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData) *ReportStatusCodeUrlTopServiceResponse {
  s.Data = v
  return s
}

type ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData struct     {
  // {"en":"url", "zh_CN":"URL"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Number of requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData) SetUrl(v string) *ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData {
  s.Url = &v
  return s
}

func (s *ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData) SetValue(v string) *ReportStatusCodeUrlTopServiceReportStatusCodeUrlTopServiceResponseData {
  s.Value = &v
  return s
}

type ReportStatusCodeUrlTopServicePaths struct {
}

func (s ReportStatusCodeUrlTopServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeUrlTopServiceParameters struct {
}

func (s ReportStatusCodeUrlTopServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeUrlTopServiceRequestHeader struct {
}

func (s ReportStatusCodeUrlTopServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeUrlTopServiceResponseHeader struct {
}

func (s ReportStatusCodeUrlTopServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeUrlTopServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowProtocolOneMinStatusCodeServiceRequest struct {
  // {"en":"Start time: 
  // 	1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 	2.Cannot exceed current time 
  // 	3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒)
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 	1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. 
  // 	2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time. 
  // 	3.If both fields of dataFrom and dateTo are left empty, the default query past 1 day; If there is only one unsent, throw an exception 
  // 	4.Maximum allowed query time interval: 1 days, Date from and dateTo, not more than 1 days", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support). 
  // 	2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"1.Selection:domain 
  // 	2.If groupBy left empty, merge date of all domains", "zh_CN":"1.可选值:domain
  // 2.不传默认聚合所有频道数据"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowProtocolOneMinStatusCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceRequest) SetDateFrom(v string) *ReportFlowProtocolOneMinStatusCodeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceRequest) SetDateTo(v string) *ReportFlowProtocolOneMinStatusCodeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceRequest) SetDomain(v []*string) *ReportFlowProtocolOneMinStatusCodeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceRequest) SetGroupBy(v []*string) *ReportFlowProtocolOneMinStatusCodeServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowProtocolOneMinStatusCodeServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowProtocolOneMinStatusCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceResponse) SetCode(v string) *ReportFlowProtocolOneMinStatusCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceResponse) SetMessage(v string) *ReportFlowProtocolOneMinStatusCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceResponse) SetData(v []*ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData) *ReportFlowProtocolOneMinStatusCodeServiceResponse {
  s.Data = v
  return s
}

type ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData struct     {
  // {"en":"Domain. If merge date of all domains will not return domain", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData) SetDomain(v string) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData) SetStatusCodeDataList(v []*ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseData {
  s.StatusCodeDataList = v
  return s
}

type ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList) SetStatusCode(v string) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList) SetDetailList(v []*ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList struct     {
  // {"en":"timestamp,Returns the timestamp between the start time and end time. Time format: yyyy-MM-dd HH:mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。时间格式:1分钟:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"https", "zh_CN":"https数据"}
  HttpsValue *string `json:"httpsValue,omitempty" xml:"httpsValue,omitempty" require:"true"`
  // {"en":"http", "zh_CN":"http数据"}
  HttpValue *string `json:"httpValue,omitempty" xml:"httpValue,omitempty" require:"true"`
}

func (s ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetHttpsValue(v string) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.HttpsValue = &v
  return s
}

func (s *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList) SetHttpValue(v string) *ReportFlowProtocolOneMinStatusCodeServiceReportFlowProtocolOneMinStatusCodeServiceResponseDataStatusCodeDataListDetailList {
  s.HttpValue = &v
  return s
}

type ReportFlowProtocolOneMinStatusCodeServicePaths struct {
}

func (s ReportFlowProtocolOneMinStatusCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServicePaths) GoString() string {
  return s.String()
}

type ReportFlowProtocolOneMinStatusCodeServiceParameters struct {
}

func (s ReportFlowProtocolOneMinStatusCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowProtocolOneMinStatusCodeServiceRequestHeader struct {
}

func (s ReportFlowProtocolOneMinStatusCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowProtocolOneMinStatusCodeServiceResponseHeader struct {
}

func (s ReportFlowProtocolOneMinStatusCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowProtocolOneMinStatusCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryRealTimeOriginStatusCodeRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.Cannot exceed current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00.For example, 2019-01-01T10:00:00+08:00
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4.The default maximum query time interval is 24 hours (you can contact technical support to adjust it, up to 31 days)", "zh_CN":"结束时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.默认允许查询最大时间间隔：24小时(可联系技术支持调整，最大31天)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1.The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3.Domain name exceeding limit, misstatement", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.域名超过上限，报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:  
  // 1.The default is 1m ;
  // 2.Support 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度：
  // 1.不传默认1m
  // 2.支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Query dimension:
  // 1.Optional value statusCode, statusCodeType, 2XX, 3XX, 4XX, 5XX. Default value is statusCode;
  // 2.StatusCode: returns the source status code details
  // 3.statusCodeType: returns the requests of eachstatus code type (such as the number of requests corresponding to success, redirect, not modified, permission, not found, server error, and other)
  // 4.2XX:Returns the 2XX return source status code summary and the 2 start return source status code detail data;
  // 5.3XX:Returns the 2XX return source status code summary and the 3 start return source status code detail data;
  // 6.4XX:Returns the 2XX return source status code summary and the 4 start return source status code detail data;
  // 7.5XX:Returns the 2XX return source status code summary and the 5 start return source status code detail data;", "zh_CN":"查询维度:
  // 1.可选值 statusCode，statusCodeType， 2XX ，3XX， 4XX， 5XX，不传默认 statusCode
  // 2.statusCode: 返回回源状态码明细；
  // 3.statusCodeType：返回状态码类型对应明细(如Success、Redirect、Not-Modified、Permission、Not-Found、Server Error、Other对应的请求数)
  // 4.2XX : 返回 2XX 回源状态码汇总及各 2 开头回源状态码明细数据；
  // 5.3XX : 返回 3XX 回源状态码汇总及各 2 开头回源状态码明细数据；
  // 6.4XX : 返回 4XX 回源状态码汇总及各 4 开头回源状态码明细数据；
  // 7.5XX : 返回 5XX回源状态码汇总及各 5 开头回源状态码明细数据；"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s QueryRealTimeOriginStatusCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeRequest) GoString() string {
  return s.String()
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDateFrom(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDateTo(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDomain(v []*string) *QueryRealTimeOriginStatusCodeRequest {
  s.Domain = v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDataInterval(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetQueryBy(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.QueryBy = &v
  return s
}

type QueryRealTimeOriginStatusCodeResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"StatusCode", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other", "zh_CN":"Success、Redirect、Not-Modified、Permission、Not-Found、Server Error、Other"}
  StatusCodeType *string `json:"statusCodeType,omitempty" xml:"statusCodeType,omitempty" require:"true"`
}

func (s QueryRealTimeOriginStatusCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeResponse) GoString() string {
  return s.String()
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetCode(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.Code = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetMessage(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.Message = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetData(v []*QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData) *QueryRealTimeOriginStatusCodeResponse {
  s.Data = v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetStatusCode(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.StatusCode = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetStatusCodeType(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.StatusCodeType = &v
  return s
}

type QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of return requests", "zh_CN":"回源总请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData) GoString() string {
  return s.String()
}

func (s *QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData) SetTimestamp(v string) *QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData {
  s.Timestamp = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData) SetValue(v string) *QueryRealTimeOriginStatusCodeQueryRealTimeOriginStatusCodeResponseData {
  s.Value = &v
  return s
}

type QueryRealTimeOriginStatusCodePaths struct {
}

func (s QueryRealTimeOriginStatusCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodePaths) GoString() string {
  return s.String()
}

type QueryRealTimeOriginStatusCodeParameters struct {
}

func (s QueryRealTimeOriginStatusCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeParameters) GoString() string {
  return s.String()
}

type QueryRealTimeOriginStatusCodeRequestHeader struct {
}

func (s QueryRealTimeOriginStatusCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryRealTimeOriginStatusCodeResponseHeader struct {
}

func (s QueryRealTimeOriginStatusCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeRealTimeOriginServiceRequest struct {
  // {"en":"Start time
  // 1.The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2.Cannot be greater than the current time
  // 3.Get up to the last six months (183 days) of data.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;
  // 3.If both fields of dataFrom and dateTo are left empty, then data in the last 1 hours will be queried by default;
  // 4.Allowable maximum time range for query: 1 hour, means the period between dateFrom to dateTo should not exceed 1 hour (can be adjusted by contacting technical support up to 31 days)", "zh_CN":"结束时间:
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:1小时(可联系技术支持调整),即dateFrom和dateTo相差不能超过1小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The default upper limit to domains that can be entered is 200 (Contact technical support to adjust);
  // 2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 3.域名超过上限,报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. Support 1m (1 minute granularity),5m (5 minutes granularity)
  // 2. The default value is 1m", "zh_CN":"数据粒度:不传默认1m
  // 1. 支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportStatusCodeRealTimeOriginServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeOriginServiceRequest) SetDateFrom(v string) *ReportStatusCodeRealTimeOriginServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeRealTimeOriginServiceRequest) SetDateTo(v string) *ReportStatusCodeRealTimeOriginServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeRealTimeOriginServiceRequest) SetDomain(v []*string) *ReportStatusCodeRealTimeOriginServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeRealTimeOriginServiceRequest) SetDataInterval(v string) *ReportStatusCodeRealTimeOriginServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportStatusCodeRealTimeOriginServiceResponse struct {
  Result []*ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeOriginServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeOriginServiceResponse) SetResult(v []*ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult) *ReportStatusCodeRealTimeOriginServiceResponse {
  s.Result = v
  return s
}

type ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult struct     {
  StatusCodeOriginData []*ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData `json:"statusCodeOriginData,omitempty" xml:"statusCodeOriginData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult) SetStatusCodeOriginData(v []*ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData) *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResult {
  s.StatusCodeOriginData = v
  return s
}

type ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData struct     {
  // {"en":"Back to origin status code type", "zh_CN":"回源状态码类型"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData) SetStatusCode(v string) *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData) SetRequestData(v []*ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData) *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginData {
  s.RequestData = v
  return s
}

type ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData struct     {
  // {"en":"The data granularity is 1 minute, and the format is yyyy MM dd HH: mm", "zh_CN":"数据粒度为1分钟,格式为yyyy-MM-dd HH:mm;"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Requests of back to origin status code", "zh_CN":"回源状态码请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData) SetTimestamp(v string) *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData) SetValue(v string) *ReportStatusCodeRealTimeOriginServiceReportStatusCodeRealTimeOriginServiceResponseResultStatusCodeOriginDataRequestData {
  s.Value = &v
  return s
}

type ReportStatusCodeRealTimeOriginServicePaths struct {
}

func (s ReportStatusCodeRealTimeOriginServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeOriginServiceParameters struct {
}

func (s ReportStatusCodeRealTimeOriginServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeOriginServiceRequestHeader struct {
}

func (s ReportStatusCodeRealTimeOriginServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeOriginServiceResponseHeader struct {
}

func (s ReportStatusCodeRealTimeOriginServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeOriginServiceResponseHeader) GoString() string {
  return s.String()
}




