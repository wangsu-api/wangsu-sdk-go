package reportvisitor

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportVisitorCustomTopDailyServiceRequest struct {
  // {"en":"Start time:
  // The time format is yyyy-MM-dd, for example, 2021-10-10;
  // It cannot be greater than the current time
  // You can get data for the last three months (90 days) at most.", "zh_CN":"开始时间：
  // 时间格式为yyyy-MM-dd，例如，2021-10-10；
  // 不能大于当前时间
  // 最多可获取最近三个月（90天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-dd, the end time must be greater than the start time. If the end time is greater than the current time, the current time is used.
  // 2. If both dateFrom and dateTo are not transmitted, the default query is the past 7 days. If only one is not transmitted, an exception is thrown;
  // 3. The maximum query time interval allowed is 31 days, that is, the difference between dateFrom and dateTo cannot exceed 31 days (you can contact technical support to adjust)", "zh_CN":"结束时间：
  // 时间格式为yyyy-MM-dd
  // 结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // dateFrom，dateTo二者都未传，默认查询过去的7天；如仅有一个未传，抛异常
  // 允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The upper limit of the number of domain names that can be passed is 100 by default (you can contact technical support to adjust it);
  // 2. Automatically filter out illegal domain names (if an illegal domain name is passed, it will be filtered out, and the query result will only return data for legal domain names);
  // 3. When this parameter is not passed, all domain names under the account will be queried by default, but an error will be prompted when the number of domain names under the account exceeds the upper limit.", "zh_CN":"域名：
  // 可传递域名数量上限默认为100个（可联系技术支持调整）。
  // 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Number of TOPs:
  // 1. If not specified, the default is TOP 10;
  // 2. The maximum number of TOPs is 100.", "zh_CN":"TOP个数：
  // 不传默认TOP 10
  // 最大TOP 100"}
  Top *string `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Sorting:
  // 1. Optional values: request, flow
  // 2. Do not pass the default request", "zh_CN":"排序：
  // 1、可选值为：request, flow
  // 2、不传默认request"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ReportVisitorCustomTopDailyServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetDateFrom(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetDateTo(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetDomain(v []*string) *ReportVisitorCustomTopDailyServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetTop(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.Top = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetOrderBy(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.OrderBy = &v
  return s
}

type ReportVisitorCustomTopDailyServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  ReportVisitorCustomTopDailyServiceData []*ReportVisitorCustomTopDailyServiceData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportVisitorCustomTopDailyServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportVisitorCustomTopDailyServiceResponse) SetCode(v string) *ReportVisitorCustomTopDailyServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceResponse) SetMessage(v string) *ReportVisitorCustomTopDailyServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceResponse) SetData(v []*ReportVisitorCustomTopDailyServiceData) *ReportVisitorCustomTopDailyServiceResponse {
  s.ReportVisitorCustomTopDailyServiceData = v
  return s
}

type ReportVisitorCustomTopDailyServiceData struct {
  // {"en":"Top", "zh_CN":"top排名"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Total traffic: The unit of measurement is MB, with 2 decimal places.", "zh_CN":"总流量：计量单位MB，保留2位小数"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Total requests", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
}

func (s ReportVisitorCustomTopDailyServiceData) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceData) GoString() string {
  return s.String()
}

func (s *ReportVisitorCustomTopDailyServiceData) SetTop(v string) *ReportVisitorCustomTopDailyServiceData {
  s.Top = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceData) SetIp(v string) *ReportVisitorCustomTopDailyServiceData {
  s.Ip = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceData) SetTotalFlow(v string) *ReportVisitorCustomTopDailyServiceData {
  s.TotalFlow = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceData) SetTotalRequest(v string) *ReportVisitorCustomTopDailyServiceData {
  s.TotalRequest = &v
  return s
}

type ReportVisitorCustomTopDailyServicePaths struct {
}

func (s ReportVisitorCustomTopDailyServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServicePaths) GoString() string {
  return s.String()
}

type ReportVisitorCustomTopDailyServiceParameters struct {
}

func (s ReportVisitorCustomTopDailyServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceParameters) GoString() string {
  return s.String()
}

type ReportVisitorCustomTopDailyServiceRequestHeader struct {
}

func (s ReportVisitorCustomTopDailyServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportVisitorCustomTopDailyServiceResponseHeader struct {
}

func (s ReportVisitorCustomTopDailyServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUvIspProvinceServiceRequest struct {
  // {"en":"Start time
  //         1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  //         2. Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  //         3. Period between dataFrom and dateTo cannot be longer than 183 days;", "zh_CN":"开始时间:
  //         1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  //         2.不能大于当前时间;
  //         3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  //         1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  //         2. the end time is larger than the start time.
  //         3. if the end time is greater than the current time, take the current time.
  //         4. DateFrom and dateTo are not uploaded, default query for the past 24 hours; if only one is not uploaded, throw an exception;
  //         5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust). ", "zh_CN":"结束时间:
  //         1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  //         2.结束时间需大于开始时间;
  //         3.结束时间如果大于当前时间,取当前时间;
  //         4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  //         5.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递该入参时查询账号下所有域名,但当账号下域名数量超过限制时不可查询(报错)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5-minute granularity", "zh_CN":"数据粒度,5m:5分钟粒度"}
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
  // {"en":"Group dimension
  //         1. Options are domain, province, isp, and more than one value can be entered;
  //         2. The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  //         1.可选值为domain、province、isp,可传入多个值;
  //         2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportUvIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUvIspProvinceServiceRequest) SetDateFrom(v string) *ReportUvIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUvIspProvinceServiceRequest) SetDateTo(v string) *ReportUvIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUvIspProvinceServiceRequest) SetDomain(v []*string) *ReportUvIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUvIspProvinceServiceRequest) SetDataInterval(v string) *ReportUvIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportUvIspProvinceServiceRequest) SetProvince(v []*string) *ReportUvIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportUvIspProvinceServiceRequest) SetIsp(v []*string) *ReportUvIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportUvIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportUvIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportUvIspProvinceServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUvIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUvIspProvinceServiceResponse) SetResult(v []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult) *ReportUvIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"ISP数据"}
  IspData []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult) SetDomain(v string) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult) SetIspData(v []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份数据"}
  ProvinceData []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Number of requests from unique IP addresses", "zh_CN":"独立IP数"}
  UvData []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData `json:"uvData,omitempty" xml:"uvData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData) SetUvData(v []*ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceData {
  s.UvData = v
  return s
}

type ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData struct     {
  // {"en":"Date
  //         1. When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  //         2. When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;
  //         3. Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间
  //         1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  //         2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;
  //         3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests from unique IP addresses", "zh_CN":"独立IP数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData) GoString() string {
  return s.String()
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData) SetTimestamp(v string) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData {
  s.Timestamp = &v
  return s
}

func (s *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData) SetValue(v string) *ReportUvIspProvinceServiceReportUvIspProvinceServiceResponseResultIspDataProvinceDataUvData {
  s.Value = &v
  return s
}

type ReportUvIspProvinceServicePaths struct {
}

func (s ReportUvIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportUvIspProvinceServiceParameters struct {
}

func (s ReportUvIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportUvIspProvinceServiceRequestHeader struct {
}

func (s ReportUvIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUvIspProvinceServiceResponseHeader struct {
}

func (s ReportUvIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUvIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportIpTopDetailsServiceRequest struct {
  // {"en":"Start date:
  //         1.The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  //         2.Cannot be greater than the current time
  //         3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间：
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年1月1日10点0分0秒);
  //         2.不能大于当前时间
  //         3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  //         1.The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  //         2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  //         3.Date from, dateTo, neither passed, default query past 24 hours; If there is only one unsent, throw an exception
  //         4.The default query interval is 7 days, maximum allowed query time interval: 31 days, i.e., the difference between Date from and dateTo cannot exceed 31 days (adjusted by contact technology support).", "zh_CN":"结束时间：
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年1月1日10点0分0秒);
  //         2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  //         3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常
  //         4.默认查询时间间隔7天，允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  //         1.The maximum number of TLDs that can be delivered is 20 by default (Technical Support Adjustments can be contacted).
  //         2.Auto-filter illegal domain name (pass illegal domain name, will be filtered, query result only returns the data of legal domain name)
  //         3.All domain names are queried by default when this entry is not delivered, but an error occurs when the number of domain names in the account exceeds the upper limit", "zh_CN":"域名：
  //         1.可传递域名数量上限默认为20个(可联系技术支持调整)。
  //         2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  //         3.未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Ordering:
  //         1.Optional values are: bandwidth, request, flow
  //         2.No default bandwidth", "zh_CN":"排序：
  //         1.可选值为：bandwidth 、request、flow
  //         2.不传默认 bandwidth"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ReportIpTopDetailsServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportIpTopDetailsServiceRequest) SetDateFrom(v string) *ReportIpTopDetailsServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportIpTopDetailsServiceRequest) SetDateTo(v string) *ReportIpTopDetailsServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportIpTopDetailsServiceRequest) SetDomain(v []*string) *ReportIpTopDetailsServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportIpTopDetailsServiceRequest) SetOrderBy(v string) *ReportIpTopDetailsServiceRequest {
  s.OrderBy = &v
  return s
}

type ReportIpTopDetailsServiceResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request Result Information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the requests", "zh_CN":"请求结果的详细数据"}
  Data []*ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportIpTopDetailsServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportIpTopDetailsServiceResponse) SetCode(v string) *ReportIpTopDetailsServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportIpTopDetailsServiceResponse) SetMessage(v string) *ReportIpTopDetailsServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportIpTopDetailsServiceResponse) SetData(v []*ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData) *ReportIpTopDetailsServiceResponse {
  s.Data = v
  return s
}

type ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData struct     {
  // {"en":"timestamp:
  //         The format is yyyy-MM-dd HH:MM:ss; Each time slice data value represents the data value in the previous time-granularity range, such as yyyy-MM-dd 00:05, Data in the range 00:00:00 to 00:05.", "zh_CN":"timestamp"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"domainData", "zh_CN":"域名数据"}
  DomainData []*ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData `json:"domainData,omitempty" xml:"domainData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData) SetTimestamp(v string) *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData) SetDomainData(v []*ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData) *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseData {
  s.DomainData = v
  return s
}

type ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData struct     {
  // {"en":"Domain name", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IpData []*ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData `json:"ipData,omitempty" xml:"ipData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData) GoString() string {
  return s.String()
}

func (s *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData) SetDomain(v string) *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData {
  s.Domain = &v
  return s
}

func (s *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData) SetIpData(v []*ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData) *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainData {
  s.IpData = v
  return s
}

type ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData struct     {
  // {"en":"Ip, default TOP 100", "zh_CN":"ip，默认 TOP 100"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"
  // 		1.Flow: Unit of measure MB, keeping 2 decimal places
  //         2.Bandwidth: Unit of measure Mbps, keeping 2 decimal places
  //         3.Request", "zh_CN":"
  // 		1.流量：计量单位MB，保留2位小数
  //         2.带宽：计量单位Mbps，保留2位小数
  //         3.请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData) GoString() string {
  return s.String()
}

func (s *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData) SetIp(v string) *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData {
  s.Ip = &v
  return s
}

func (s *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData) SetValue(v string) *ReportIpTopDetailsServiceReportIpTopDetailsServiceResponseDataDomainDataIpData {
  s.Value = &v
  return s
}

type ReportIpTopDetailsServicePaths struct {
}

func (s ReportIpTopDetailsServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServicePaths) GoString() string {
  return s.String()
}

type ReportIpTopDetailsServiceParameters struct {
}

func (s ReportIpTopDetailsServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceParameters) GoString() string {
  return s.String()
}

type ReportIpTopDetailsServiceRequestHeader struct {
}

func (s ReportIpTopDetailsServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportIpTopDetailsServiceResponseHeader struct {
}

func (s ReportIpTopDetailsServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportIpTopDetailsServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportReferrerTopDetailsServiceRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  // 2.Cannot be greater than the current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年1月1日10点0分0秒)；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, dateTo, neither passed, default query past 24 hours; If there is only one unsent, throw an exception
  // 4.Maximum allowed query time interval: 31 days, i.e., the difference between Date from and dateTo cannot exceed 31 days (adjusted by contact technology support).", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年1月1日10点0分0秒)；
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔:31天，即dateFrom和dateTo相差不能超过31天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1.The maximum number of TLDs that can be delivered is 20 by default (Technical Support Adjustments can be contacted).
  // 2.Auto-filter illegal domain name (pass illegal domain name, will be filtered, query result only returns the data of legal domain name)
  // 3.All domain names are queried by default when this entry is not delivered, but an error occurs when the number of domain names in the account exceeds the upper limit", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)。
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Ordering:
  // 1.Optional values are: bandwidth, request, flow
  // 2.No default bandwidth", "zh_CN":"排序:
  // 1.可选值为:bandwidth 、request、flow
  // 2.不传默认 bandwidth"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ReportReferrerTopDetailsServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportReferrerTopDetailsServiceRequest) SetDateFrom(v string) *ReportReferrerTopDetailsServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceRequest) SetDateTo(v string) *ReportReferrerTopDetailsServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceRequest) SetDomain(v []*string) *ReportReferrerTopDetailsServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportReferrerTopDetailsServiceRequest) SetOrderBy(v string) *ReportReferrerTopDetailsServiceRequest {
  s.OrderBy = &v
  return s
}

type ReportReferrerTopDetailsServiceResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request Result Information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on results of requests", "zh_CN":"请求结果的详细数据"}
  Data []*ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportReferrerTopDetailsServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportReferrerTopDetailsServiceResponse) SetCode(v string) *ReportReferrerTopDetailsServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceResponse) SetMessage(v string) *ReportReferrerTopDetailsServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceResponse) SetData(v []*ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData) *ReportReferrerTopDetailsServiceResponse {
  s.Data = v
  return s
}

type ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData struct     {
  // {"en":"timestamp:
  // The format is yyyy-MM-dd HH:MM:ss; Each time slice data value represents the data value in the previous time-granularity range, such as yyyy-MM-dd 00:05, Data in the range 00:00:00 to 00:05.", "zh_CN":"timestamp"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  DomainData []*ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData `json:"domainData,omitempty" xml:"domainData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData) SetTimestamp(v string) *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData) SetDomainData(v []*ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData) *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseData {
  s.DomainData = v
  return s
}

type ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData struct     {
  // {"en":"Domain name", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  ReferData []*ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData `json:"referData,omitempty" xml:"referData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData) GoString() string {
  return s.String()
}

func (s *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData) SetDomain(v string) *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData {
  s.Domain = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData) SetReferData(v []*ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData) *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainData {
  s.ReferData = v
  return s
}

type ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData struct     {
  // {"en":"Refer, default TOP 100", "zh_CN":"refer，默认 TOP 100"}
  Refer *string `json:"refer,omitempty" xml:"refer,omitempty" require:"true"`
  // {"en":"
  // 									  1.Flow: Unit of measure MB, keeping 2 decimal places
  //                                       2.Bandwidth: Unit of measure Mbps, keeping 2 decimal places
  //                                       3.Request", "zh_CN":"
  // 									  1.流量:计量单位MB，保留2位小数
  //                                       2.带宽:计量单位Mbps，保留2位小数
  //                                       3.请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData) GoString() string {
  return s.String()
}

func (s *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData) SetRefer(v string) *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData {
  s.Refer = &v
  return s
}

func (s *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData) SetValue(v string) *ReportReferrerTopDetailsServiceReportReferrerTopDetailsServiceResponseDataDomainDataReferData {
  s.Value = &v
  return s
}

type ReportReferrerTopDetailsServicePaths struct {
}

func (s ReportReferrerTopDetailsServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServicePaths) GoString() string {
  return s.String()
}

type ReportReferrerTopDetailsServiceParameters struct {
}

func (s ReportReferrerTopDetailsServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceParameters) GoString() string {
  return s.String()
}

type ReportReferrerTopDetailsServiceRequestHeader struct {
}

func (s ReportReferrerTopDetailsServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportReferrerTopDetailsServiceResponseHeader struct {
}

func (s ReportReferrerTopDetailsServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportReferrerTopDetailsServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUrlTopDetailsServiceRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  // 2. Cannot be greater than the current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2019年1月1日10点0分0秒)；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, dateTo, neither passed, default query past 24 hours; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 7 days, i.e., the difference between Date from and dateTo cannot exceed 7 days (adjusted by contact technology support, the maximum interval is 31 days).", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2019年1月1日10点0分0秒)；
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时；如仅有一个未传,抛异常
  // 4.默认最大查询时间间隔:7天，即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最大间隔31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of TLDs that can be delivered is 20 by default (Technical Support Adjustments can be contacted).
  // 2. Auto-filter illegal domain name (pass illegal domain name, will be filtered, query result only returns the data of legal domain name)
  // 3. All domain names are queried by default when this entry is not delivered, but an error occurs when the number of domain names in the account exceeds the upper limit", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)。
  // 2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Ordering:
  // 1. Optional values are: bandwidth, request, flow
  // 2. No default bandwidth", "zh_CN":"排序:
  // 1.可选值为:bandwidth 、request、flow
  // 2.不传默认 bandwidth"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ReportUrlTopDetailsServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUrlTopDetailsServiceRequest) SetDateFrom(v string) *ReportUrlTopDetailsServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUrlTopDetailsServiceRequest) SetDateTo(v string) *ReportUrlTopDetailsServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUrlTopDetailsServiceRequest) SetDomain(v []*string) *ReportUrlTopDetailsServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUrlTopDetailsServiceRequest) SetOrderBy(v string) *ReportUrlTopDetailsServiceRequest {
  s.OrderBy = &v
  return s
}

type ReportUrlTopDetailsServiceResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request Result Information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on results of requests", "zh_CN":"请求结果的详细数据"}
  Data []*ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlTopDetailsServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUrlTopDetailsServiceResponse) SetCode(v string) *ReportUrlTopDetailsServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportUrlTopDetailsServiceResponse) SetMessage(v string) *ReportUrlTopDetailsServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportUrlTopDetailsServiceResponse) SetData(v []*ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData) *ReportUrlTopDetailsServiceResponse {
  s.Data = v
  return s
}

type ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData struct     {
  // {"en":"timestamp
  // 1. the format is yyyy-MM-dd HH:MM:ss; Each time slice data value represents the data value in the previous time-granularity range, such as yyyy-MM-dd 00:05, Data in the range 00:00:00 to 00:05.", "zh_CN":"timestamp"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  DomainData []*ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData `json:"domainData,omitempty" xml:"domainData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData) SetTimestamp(v string) *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData) SetDomainData(v []*ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData) *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseData {
  s.DomainData = v
  return s
}

type ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData struct     {
  // {"en":"Domain name", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  UrlData []*ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData `json:"urlData,omitempty" xml:"urlData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData) GoString() string {
  return s.String()
}

func (s *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData) SetDomain(v string) *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData {
  s.Domain = &v
  return s
}

func (s *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData) SetUrlData(v []*ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData) *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainData {
  s.UrlData = v
  return s
}

type ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData struct     {
  // {"en":"URL, default TOP 100", "zh_CN":"url,默认 TOP 100"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"1. Flow: Unit of measure MB, keeping 2 decimal places;
  // 									2. Bandwidth: Unit of measure Mbps, keeping 2 decimal places;
  // 									3. Request", "zh_CN":"1.流量:计量单位MB,保留2位小数; 
  // 									2.带宽:计量单位Mbps,保留2位小数
  // 									3.请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData) GoString() string {
  return s.String()
}

func (s *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData) SetUrl(v string) *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData {
  s.Url = &v
  return s
}

func (s *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData) SetValue(v string) *ReportUrlTopDetailsServiceReportUrlTopDetailsServiceResponseDataDomainDataUrlData {
  s.Value = &v
  return s
}

type ReportUrlTopDetailsServicePaths struct {
}

func (s ReportUrlTopDetailsServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServicePaths) GoString() string {
  return s.String()
}

type ReportUrlTopDetailsServiceParameters struct {
}

func (s ReportUrlTopDetailsServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceParameters) GoString() string {
  return s.String()
}

type ReportUrlTopDetailsServiceRequestHeader struct {
}

func (s ReportUrlTopDetailsServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUrlTopDetailsServiceResponseHeader struct {
}

func (s ReportUrlTopDetailsServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlTopDetailsServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryTotalNumberofUniqueIPUnderSingleDomainRequest struct {
  // {"en":"Domain, maximum number is 1", "zh_CN":"域名，数量上限1个"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Stream name:
  // 1. Limit to the number of streams can be adjusted depending on different accounts. The default value is 20;
  // 2. All streams are queried by default is this field is left empty and at the same time, the upper limit of the number of streams can be set.", "zh_CN":"流名：
  // 1.流名个数限制根据账号可调，默认为20个；
  // 2.不传时默认查询域名下所有流名，同时受流名数量上限限制；"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainRequest) SetDomain(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainRequest {
  s.Domain = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainRequest) SetStream(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainRequest {
  s.Stream = &v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainResponse struct {
  Result []*QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponse) SetResult(v []*QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) *QueryTotalNumberofUniqueIPUnderSingleDomainResponse {
  s.Result = v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) SetDomain(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) SetDetails(v []*QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResult {
  s.Details = v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails struct     {
  // {"en":"Stream name", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"Time, format is yyyy-MM-dd", "zh_CN":"时间，格式为yyyy-MM-dd"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of unique IP addresses", "zh_CN":"独立IP数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) SetStream(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails {
  s.Stream = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) SetTimestamp(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) SetTotal(v int) *QueryTotalNumberofUniqueIPUnderSingleDomainQueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails {
  s.Total = &v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainPaths struct {
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainPaths) GoString() string {
  return s.String()
}

type QueryTotalNumberofUniqueIPUnderSingleDomainParameters struct {
  // {"en":"Start time
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. Must be smaller than the current time and dateTo;
  // 3. Period between dataFrom and dateTo cannot be longer than 3 days(technical support can be contacted to adjust);
  // 4. You can only query data for the last 2 years.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过3天（可联系技术支持调整）；
  // 4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. Must be greater than dateFrom;
  // 3. The query range must include 00:00:00 of a certain day to query the data of that day. For example, if the query range includes 2017-11-07 00:00:00, the data of 2017-11-07 can be queried.", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；
  // 3.查询范围必须包含某一天的00:00:00，才能查询到当天数据，例如查询范围包含 2017-11-07 00:00:00 可以查询到2017-11-07当天数据。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainParameters) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainParameters) SetDateFrom(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainParameters) SetDateTo(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainParameters {
  s.DateTo = &v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainRequestHeader struct {
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryTotalNumberofUniqueIPUnderSingleDomainResponseHeader struct {
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseHeader) GoString() string {
  return s.String()
}




