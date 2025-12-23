package reportbandwidth

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type PerzoneBillingRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope. 2.With format yyyy-mm-dd. 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope. 2.With format yyyy-mm-dd 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
  // {"en":"In Greenwich time zone, the value GMT+09:00 means East Nine, GMT-09:00 means West Nine, and the default is East Eight.", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"Whether it is a DWA product, 1: yes; 0: no, if not fill in, the default is 0", "zh_CN":"是否DWA产品，1:是;0:否,不填默认0"}
  Dwa *string `json:"dwa,omitempty" xml:"dwa,omitempty"`
  // {"en":"Data type (1: bandwidth|2: traffic) default 1", "zh_CN":"数据类型（1:带宽|2:流量）默认1"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PerzoneBillingRequest) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingRequest) GoString() string {
  return s.String()
}

func (s *PerzoneBillingRequest) SetCust(v string) *PerzoneBillingRequest {
  s.Cust = &v
  return s
}

func (s *PerzoneBillingRequest) SetDate(v string) *PerzoneBillingRequest {
  s.Date = &v
  return s
}

func (s *PerzoneBillingRequest) SetStartDate(v string) *PerzoneBillingRequest {
  s.StartDate = &v
  return s
}

func (s *PerzoneBillingRequest) SetEndDate(v string) *PerzoneBillingRequest {
  s.EndDate = &v
  return s
}

func (s *PerzoneBillingRequest) SetTimezone(v string) *PerzoneBillingRequest {
  s.Timezone = &v
  return s
}

func (s *PerzoneBillingRequest) SetChannel(v string) *PerzoneBillingRequest {
  s.Channel = &v
  return s
}

func (s *PerzoneBillingRequest) SetAccetype(v string) *PerzoneBillingRequest {
  s.Accetype = &v
  return s
}

func (s *PerzoneBillingRequest) SetDwa(v string) *PerzoneBillingRequest {
  s.Dwa = &v
  return s
}

func (s *PerzoneBillingRequest) SetType(v string) *PerzoneBillingRequest {
  s.Type = &v
  return s
}

type PerzoneBillingResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *PerzoneBillingResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponse) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponse) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponse) SetProvider(v *PerzoneBillingResponseProvider) *PerzoneBillingResponse {
  s.Provider = v
  return s
}

type PerzoneBillingResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'perzone带宽数据'}
  Date *PerzoneBillingResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProvider) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProvider) SetName(v string) *PerzoneBillingResponseProvider {
  s.Name = &v
  return s
}

func (s *PerzoneBillingResponseProvider) SetType(v string) *PerzoneBillingResponseProvider {
  s.Type = &v
  return s
}

func (s *PerzoneBillingResponseProvider) SetResultType(v string) *PerzoneBillingResponseProvider {
  s.ResultType = &v
  return s
}

func (s *PerzoneBillingResponseProvider) SetDate(v *PerzoneBillingResponseProviderDate) *PerzoneBillingResponseProvider {
  s.Date = v
  return s
}

type PerzoneBillingResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'数据类型1:带宽|2:流量'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  AreaMap *PerzoneBillingResponseProviderDateAreaMap `json:"areaMap,omitempty" xml:"areaMap,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDate) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDate) SetStartDate(v string) *PerzoneBillingResponseProviderDate {
  s.StartDate = &v
  return s
}

func (s *PerzoneBillingResponseProviderDate) SetEndDate(v string) *PerzoneBillingResponseProviderDate {
  s.EndDate = &v
  return s
}

func (s *PerzoneBillingResponseProviderDate) SetType(v string) *PerzoneBillingResponseProviderDate {
  s.Type = &v
  return s
}

func (s *PerzoneBillingResponseProviderDate) SetAreaMap(v *PerzoneBillingResponseProviderDateAreaMap) *PerzoneBillingResponseProviderDate {
  s.AreaMap = v
  return s
}

type PerzoneBillingResponseProviderDateAreaMap struct {
  // {'en':'perzone', 'zh_CN':'perzone数据'}
  Perzone *PerzoneBillingResponseProviderDateAreaMapPerzone `json:"perzone,omitempty" xml:"perzone,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponseProviderDateAreaMap) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDateAreaMap) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDateAreaMap) SetPerzone(v *PerzoneBillingResponseProviderDateAreaMapPerzone) *PerzoneBillingResponseProviderDateAreaMap {
  s.Perzone = v
  return s
}

type PerzoneBillingResponseProviderDateAreaMapPerzone struct {
  // {'en':'name', 'zh_CN':'perzone区域'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'peakValue', 'zh_CN':'峰值'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'peakTime', 'zh_CN':'峰值时间点'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'totalFlow', 'zh_CN':'总流量'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Detail []*PerzoneBillingResponseProviderDateAreaMapPerzoneDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzone) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzone) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetName(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.Name = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetPeakValue(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.PeakValue = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetPeakTime(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.PeakTime = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetTotalFlow(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.TotalFlow = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetDetail(v []*PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.Detail = v
  return s
}

type PerzoneBillingResponseProviderDateAreaMapPerzoneDetail struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) SetTime(v string) *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail {
  s.Time = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) SetText(v string) *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail {
  s.Text = &v
  return s
}

type PerzoneBillingPaths struct {
}

func (s PerzoneBillingPaths) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingPaths) GoString() string {
  return s.String()
}

type PerzoneBillingParameters struct {
}

func (s PerzoneBillingParameters) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingParameters) GoString() string {
  return s.String()
}

type PerzoneBillingRequestHeader struct {
}

func (s PerzoneBillingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingRequestHeader) GoString() string {
  return s.String()
}

type PerzoneBillingResponseHeader struct {
}

func (s PerzoneBillingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseHeader) GoString() string {
  return s.String()
}




type ReportAppaFlowDomainCountryServiceRequest struct {
  // {"en":"The time format for the start time is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm represents the time zone offset, which can be adjusted according to your data requirements. For example, +00:00 stands for UTC time, +08:00 for UTC+8, and -05:00 for UTC-5. For instance, 2024-01-15T10:30:45+00:00 represents 10:30:45 AM, January 15, 2024 in UTC. The start time cannot be later than the current time. You can retrieve data for up to the most recent six months (183 days).","zh_CN":"开始时间\n1. 时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒。\n2. 不能大于当前时间。\n3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"The time format for the end time is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted based on your data requirements. For example, +00:00 represents UTC time, +08:00 represents UTC+8, and -05:00 represents UTC-5. For instance, 2024-01-15T10:30:45+00:00 represents 10:30:45 AM on January 15, 2024, in UTC. The end time must be greater than the start time. When the data granularity is set to '1d', it should be set to the end of the day at 23:59:59. If the end time is later than the current time, the current time will be used as the end time. if only one of them is missing, an exception will be thrown. The maximum allowed query time interval is 7 days; that is, the difference between dateFrom and dateTo cannot exceed 7 days (this limit can be adjusted by contacting technical support, up to a maximum of 31 days).","zh_CN":"结束时间\n1. 时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n2. 当数据粒度为天粒度时，需设置为当天结束时间的23:59:59\n3. 结束时间需大于开始时间，结束时间如果大于当前时间,取当前时间。\n4. 允许查询最大时间间隔7天，即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最长31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain Name: The maximum number of domain names that can be submitted is 20 by default (this limit can be adjusted by contacting technical support); Note: Authentication is required to verify whether the account has permission for these domain names during the specified time period.","zh_CN":"域名:\n可传递域名数量上限默认为20个(可联系技术支持调整);\n特别注意：需要鉴权，判断账号这个时间这些域名是否有权限"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data Granularity:  Supports 5m (5-minute granularity) and 1d (daily granularity). If not specified, the default is 5m.","zh_CN":"数据粒度:\n1. 支持5m(5分钟粒度),1d(天粒度)\n2. 不传默认为5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Country/Region Code:  If not specified, all countries and regions will be queried by default.","zh_CN":"国家地区代号: 不传默认查询全部国家地区;"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Optional values: domain, country, and aggregatedOversea. You can provide one or multiple values, but aggregatedOversea and country cannot be specified at the same time. If provided, the data will be displayed in detail based on the specified dimensions: domain: Data will be grouped and displayed by domain. country: Data will be grouped and displayed by country within each domain. aggregatedOversea: Data will be grouped and displayed by domestic and overseas dimensions. The order of the levels in the returned results is fixed; the order of parameters in the request does not affect the order of the results. For example, groupBy: [domain, country] and groupBy: [country, domain] will return results in the same order.","zh_CN":"1.不传该值,则不分组; 可选值domain、country、aggregatedOversea，可传入单个或多个值，其中不能同时传aggregatedOversea 和 country；\n2.有传入则按照该维度展示明细数据: \n(1) domain:按照域名维度进行分组展示;\n(2) country:按照国家维度进行分组展示;\n(3) aggregatedOversea:按照国内和oversea维度进行分组展示 \n3.返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:groupBy: [domain,country]与groupBy: [country,domain]返回结果一样。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"Optional values, single selection. If not specified, the default is total, upstream (uplink), downstream (downlink), and total (uplink + downlink).","zh_CN":"单选，不传默认total,可选值有: upstream（上行）、downstream（下行）、total（上行+下行）"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
  // {"en":"Set the rounding method for dateFrom, dateTo, and return timestamp, rounding up or down. Only valid when the granularity is less than 1d (not including 1 day)\n1. up-rounds up, e.g.: granularity=5m, 00:00:00-00:04:49 will be displayed as 00:05:00\n2. down-rounds down, e.g.: granularity=5m, 00:00:00-00:04:49 will be displayed as 00:00:00\n3. If not specified, the result will be rounded up (up).","zh_CN":"设置dateFrom和dateTo和返回时间的归整方式，进行向上或向下取整。仅当粒度小于1天（不包含1天）时有效。\n1. up – 向上取整。例如：granularity=5m，00:00:00-00:04:49 将会显示为 00:05:00\n2. down-向下取整，例如：granularity=5m时，00:00:00-00:04:49将显示为00:00:00\n3. 如未传值，结果将进行向上取整(up)"}
  TimeRounding *string `json:"timeRounding,omitempty" xml:"timeRounding,omitempty"`
}

func (s ReportAppaFlowDomainCountryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetDateFrom(v string) *ReportAppaFlowDomainCountryServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetDateTo(v string) *ReportAppaFlowDomainCountryServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetDomain(v []*string) *ReportAppaFlowDomainCountryServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetGranularity(v string) *ReportAppaFlowDomainCountryServiceRequest {
  s.Granularity = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetCountryCode(v []*string) *ReportAppaFlowDomainCountryServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetGroupBy(v []*string) *ReportAppaFlowDomainCountryServiceRequest {
  s.GroupBy = v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetDataType(v string) *ReportAppaFlowDomainCountryServiceRequest {
  s.DataType = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceRequest) SetTimeRounding(v string) *ReportAppaFlowDomainCountryServiceRequest {
  s.TimeRounding = &v
  return s
}

type ReportAppaFlowDomainCountryServiceRequestHeader struct {
}

func (s ReportAppaFlowDomainCountryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportAppaFlowDomainCountryServicePaths struct {
}

func (s ReportAppaFlowDomainCountryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServicePaths) GoString() string {
  return s.String()
}

type ReportAppaFlowDomainCountryServiceParameters struct {
}

func (s ReportAppaFlowDomainCountryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceParameters) GoString() string {
  return s.String()
}

type ReportAppaFlowDomainCountryServiceResponse struct {
  // {"en":"Response code","zh_CN":"响应状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Reponse message","zh_CN":"响应结果"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*ReportAppaFlowDomainCountryServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAppaFlowDomainCountryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportAppaFlowDomainCountryServiceResponse) SetCode(v string) *ReportAppaFlowDomainCountryServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponse) SetMessage(v string) *ReportAppaFlowDomainCountryServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponse) SetData(v []*ReportAppaFlowDomainCountryServiceResponseData) *ReportAppaFlowDomainCountryServiceResponse {
  s.Data = v
  return s
}

type ReportAppaFlowDomainCountryServiceResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  CountryData []*ReportAppaFlowDomainCountryServiceResponseDataCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAppaFlowDomainCountryServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportAppaFlowDomainCountryServiceResponseData) SetDomain(v string) *ReportAppaFlowDomainCountryServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseData) SetCountryData(v []*ReportAppaFlowDomainCountryServiceResponseDataCountryData) *ReportAppaFlowDomainCountryServiceResponseData {
  s.CountryData = v
  return s
}

type ReportAppaFlowDomainCountryServiceResponseDataCountryData struct     {
  // {"en":"Country code","zh_CN":"国家地区代号"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Country name","zh_CN":"国家地区名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  // {"en":"Summary of traffic in national regions: Summary of traffic flow in a single country region during the query period, unit of measure MB, retaining 2 decimal places","zh_CN":"国家地区流量汇总:单个国家地区流量在查询时段内的流量汇总值,计量单位MB,保留2位小数"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"National regional traffic ratio: the proportion (percentage) of traffic value in a single country region during the query period, 2 decimal places.","zh_CN":"国家地区流量占比:单个国家地区流量在查询时段内的流量值的占比(百分比),保留2位小数"}
  TrafficPercentage *string `json:"trafficPercentage,omitempty" xml:"trafficPercentage,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAppaFlowDomainCountryServiceResponseDataCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceResponseDataCountryData) GoString() string {
  return s.String()
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryData) SetCountryCode(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryData) SetCountryName(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryData {
  s.CountryName = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryData) SetTotalTraffic(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryData {
  s.TotalTraffic = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryData) SetTrafficPercentage(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryData {
  s.TrafficPercentage = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryData) SetDataSeries(v []*ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries) *ReportAppaFlowDomainCountryServiceResponseDataCountryData {
  s.DataSeries = v
  return s
}

type ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries struct     {
  // {"en":"The data timestamp uses the time zone specified by the input dataFrom and dataTo parameter. By default, the timestamp is rounded up, but you can set the rounding direction (up or down) via the timeRounding parameter.","zh_CN":"数据时间，时区根据入参决定，默认向上取整，可根据timeRounding调整向上或向下取整"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth value, unit of measure Mbps, 2 decimal places.","zh_CN":"带宽值,计量单位Mbps,保留2位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Flow value, unit of measure MB, 2 decimal places","zh_CN":"流量值,计量单位MB,保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries) GoString() string {
  return s.String()
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries) SetTimestamp(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries) SetBandwidth(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries {
  s.Bandwidth = &v
  return s
}

func (s *ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries) SetTraffic(v string) *ReportAppaFlowDomainCountryServiceResponseDataCountryDataDataSeries {
  s.Traffic = &v
  return s
}

type ReportAppaFlowDomainCountryServiceResponseHeader struct {
}

func (s ReportAppaFlowDomainCountryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportAppaFlowDomainCountryServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportDirBandwidthInfoServiceRequest struct {
  // {'en':'Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00:00 Beijing time on December 2, 2016);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.', 'zh_CN':'开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. The end time is greater than the start time.
  // 
  // 3. If the end time is greater than the current time, the current time is taken.
  // 
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days. ', 'zh_CN':'结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：31天，即dateFrom和dateTo相差不能超过31天。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Acceleration area:
  // 
  // 1.Acceleration areaCode is not uploaded: Query all acceleration areas by default.
  // 
  // 2.Acceleration areaCode is uploaded: Multiple can be uploaded, such as cn, af.  Please refer to the appendix description section of the overview page.', 'zh_CN':'加速区域：
  // 
  // 1.未传递areaCode时，默认查询所有加速区域；
  // 
  // 2.有传递areaCode时：可传多个，如cn, af。可传递的值详见概览页附录说明章节'}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {'en':'Directory levels, value range 1-4. Only one vlaue can be submitted', 'zh_CN':'目录层级,取值范围1~4,只能提交单个值'}
  DirHierarchy *string `json:"dirHierarchy,omitempty" xml:"dirHierarchy,omitempty" require:"true"`
  DomainDir []*ReportDirBandwidthInfoServiceRequestDomainDir `json:"domainDir,omitempty" xml:"domainDir,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDirBandwidthInfoServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDirBandwidthInfoServiceRequest) SetDateFrom(v string) *ReportDirBandwidthInfoServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceRequest) SetDateTo(v string) *ReportDirBandwidthInfoServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceRequest) SetAreaCode(v []*string) *ReportDirBandwidthInfoServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportDirBandwidthInfoServiceRequest) SetDirHierarchy(v string) *ReportDirBandwidthInfoServiceRequest {
  s.DirHierarchy = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceRequest) SetDomainDir(v []*ReportDirBandwidthInfoServiceRequestDomainDir) *ReportDirBandwidthInfoServiceRequest {
  s.DomainDir = v
  return s
}

type ReportDirBandwidthInfoServiceRequestDomainDir struct     {
  // {'en':'Domains.', 'zh_CN':'域名：
  // 
  // 1.域名个数限制根据账号可调,默认为1个'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Table of contents
  // 
  // 1.Directory number limits can be adjusted depending on different accounts. The default value is 200;
  // 
  // 2.Empty value means to query all directories. Number of directories shall not exceed set limit;
  // 
  // 3.Invalid directories are not returned', 'zh_CN':'目录
  // 
  // 1.目录个数限制根据账号可调,默认为200个。更多找技术支持调整;
  // 
  // 2.不传代表查询该域名下的所有目录,同时接受目录个数限制;
  // 
  // 3.无效的目录不返回'}
  Dir []*string `json:"dir,omitempty" xml:"dir,omitempty" type:"Repeated"`
}

func (s ReportDirBandwidthInfoServiceRequestDomainDir) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceRequestDomainDir) GoString() string {
  return s.String()
}

func (s *ReportDirBandwidthInfoServiceRequestDomainDir) SetDomain(v string) *ReportDirBandwidthInfoServiceRequestDomainDir {
  s.Domain = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceRequestDomainDir) SetDir(v []*string) *ReportDirBandwidthInfoServiceRequestDomainDir {
  s.Dir = v
  return s
}

type ReportDirBandwidthInfoServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*ReportDirBandwidthInfoServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDirBandwidthInfoServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDirBandwidthInfoServiceResponse) SetResult(v []*ReportDirBandwidthInfoServiceResponseResult) *ReportDirBandwidthInfoServiceResponse {
  s.Result = v
  return s
}

type ReportDirBandwidthInfoServiceResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'The total flow under the domain name, in MB, with 2 decimal places retained.', 'zh_CN':'域名下的总流量,单位MB,保留2位小数'}
  DomainTotalFlow *string `json:"domainTotalFlow,omitempty" xml:"domainTotalFlow,omitempty" require:"true"`
  // {'en':'Directory details under the domain.', 'zh_CN':'域名下的目录详情'}
  Details []*ReportDirBandwidthInfoServiceResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDirBandwidthInfoServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportDirBandwidthInfoServiceResponseResult) SetDomain(v string) *ReportDirBandwidthInfoServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceResponseResult) SetDomainTotalFlow(v string) *ReportDirBandwidthInfoServiceResponseResult {
  s.DomainTotalFlow = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceResponseResult) SetDetails(v []*ReportDirBandwidthInfoServiceResponseResultDetails) *ReportDirBandwidthInfoServiceResponseResult {
  s.Details = v
  return s
}

type ReportDirBandwidthInfoServiceResponseResultDetails struct     {
  // {'en':'The directory.l', 'zh_CN':'具体目录'}
  Dir *string `json:"dir,omitempty" xml:"dir,omitempty" require:"true"`
  // {'en':'Total flow under the directory.', 'zh_CN':'目录下的总流量'}
  DirTotalFlow *string `json:"dirTotalFlow,omitempty" xml:"dirTotalFlow,omitempty" require:"true"`
  // {'en':'Time segment details under specific directory.', 'zh_CN':'具体目录下的时间片段明细'}
  Details []*ReportDirBandwidthInfoServiceResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDirBandwidthInfoServiceResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceResponseResultDetails) GoString() string {
  return s.String()
}

func (s *ReportDirBandwidthInfoServiceResponseResultDetails) SetDir(v string) *ReportDirBandwidthInfoServiceResponseResultDetails {
  s.Dir = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceResponseResultDetails) SetDirTotalFlow(v string) *ReportDirBandwidthInfoServiceResponseResultDetails {
  s.DirTotalFlow = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceResponseResultDetails) SetDetails(v []*ReportDirBandwidthInfoServiceResponseResultDetailsDetails) *ReportDirBandwidthInfoServiceResponseResultDetails {
  s.Details = v
  return s
}

type ReportDirBandwidthInfoServiceResponseResultDetailsDetails struct     {
  // {'en':'Time:
  // 
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  // 2. Return the time slices that contained in start time and in end time.', 'zh_CN':'时间，
  // 查询的数据粒度为5m，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Bandwidth value. Unit is Mbps and 2 digits of decimals are allowed.', 'zh_CN':'带宽，单位Mbps。保留2位小数'}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {'en':'Flow value. Unit is MB and 2 digits of decimals are allowed.', 'zh_CN':'流量，单位MB。保留2位小数'}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s ReportDirBandwidthInfoServiceResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *ReportDirBandwidthInfoServiceResponseResultDetailsDetails) SetTimestamp(v string) *ReportDirBandwidthInfoServiceResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceResponseResultDetailsDetails) SetBandwidth(v string) *ReportDirBandwidthInfoServiceResponseResultDetailsDetails {
  s.Bandwidth = &v
  return s
}

func (s *ReportDirBandwidthInfoServiceResponseResultDetailsDetails) SetFlow(v string) *ReportDirBandwidthInfoServiceResponseResultDetailsDetails {
  s.Flow = &v
  return s
}

type ReportDirBandwidthInfoServicePaths struct {
}

func (s ReportDirBandwidthInfoServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServicePaths) GoString() string {
  return s.String()
}

type ReportDirBandwidthInfoServiceParameters struct {
}

func (s ReportDirBandwidthInfoServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceParameters) GoString() string {
  return s.String()
}

type ReportDirBandwidthInfoServiceRequestHeader struct {
}

func (s ReportDirBandwidthInfoServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDirBandwidthInfoServiceResponseHeader struct {
}

func (s ReportDirBandwidthInfoServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDirBandwidthInfoServiceResponseHeader) GoString() string {
  return s.String()
}




type BandwidthChannelProtocolRequest struct {
  // {"en":"Specifies the query date:\n1.With format yyyy-mm-dd.\n2.If not Specifies,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.\n2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.\n2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1.If there are multiple inputs,use  ';' as separator.\n2.If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:\n1.'true' as default.\n2. If not 'true',it will query data of channels that ends with any item of input 'channel's.","zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2.If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type:\n1.If there are multiple inputs,use ';' as separator.\n2.If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1.optional values:xml, json.\n2.'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way:\n1.If specified 1,get the merged result.\n2.If specified 2,get the separate result.\n3.If specified 3,get both merged result and separate result.\n4.If not specified,means '1'.","zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"data type:\n1.If specified 1,get the bandwidth data of both http and https protocals.\n2.If specified 2,get the bandwidth data of http protocal.\n3.If specified 3,get the bandwidth data of https protocal.","zh_CN":"查询的数据类型。填写1时：输出http+https的带宽数据；填写2时：输出http的带宽数据；填写3时：输出https的带宽数据。多个值请用英文分号';'.不选或者为空时默认为'1'。"}
  Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
}

func (s BandwidthChannelProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolRequest) GoString() string {
  return s.String()
}

func (s *BandwidthChannelProtocolRequest) SetDate(v string) *BandwidthChannelProtocolRequest {
  s.Date = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetStartdate(v string) *BandwidthChannelProtocolRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetEnddate(v string) *BandwidthChannelProtocolRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetChannel(v string) *BandwidthChannelProtocolRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetIsExactMatch(v string) *BandwidthChannelProtocolRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetRegion(v string) *BandwidthChannelProtocolRequest {
  s.Region = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetAccetype(v string) *BandwidthChannelProtocolRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetDataformat(v string) *BandwidthChannelProtocolRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetResultType(v string) *BandwidthChannelProtocolRequest {
  s.ResultType = &v
  return s
}

func (s *BandwidthChannelProtocolRequest) SetDatatype(v string) *BandwidthChannelProtocolRequest {
  s.Datatype = &v
  return s
}

type BandwidthChannelProtocolRequestHeader struct {
}

func (s BandwidthChannelProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolRequestHeader) GoString() string {
  return s.String()
}

type BandwidthChannelProtocolPaths struct {
}

func (s BandwidthChannelProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolPaths) GoString() string {
  return s.String()
}

type BandwidthChannelProtocolParameters struct {
}

func (s BandwidthChannelProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolParameters) GoString() string {
  return s.String()
}

type BandwidthChannelProtocolResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *BandwidthChannelProtocolResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolResponse) GoString() string {
  return s.String()
}

func (s *BandwidthChannelProtocolResponse) SetProvider(v *BandwidthChannelProtocolResponseProvider) *BandwidthChannelProtocolResponse {
  s.Provider = v
  return s
}

type BandwidthChannelProtocolResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"resultType","zh_CN":"统计类型"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {"en":"data","zh_CN":"频道带宽区分协议带宽数据"}
  Date *BandwidthChannelProtocolResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelProtocolResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthChannelProtocolResponseProvider) SetName(v string) *BandwidthChannelProtocolResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProvider) SetType(v string) *BandwidthChannelProtocolResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProvider) SetResultType(v string) *BandwidthChannelProtocolResponseProvider {
  s.ResultType = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProvider) SetDate(v *BandwidthChannelProtocolResponseProviderDate) *BandwidthChannelProtocolResponseProvider {
  s.Date = v
  return s
}

type BandwidthChannelProtocolResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"channel","zh_CN":"频道"}
  Channel *BandwidthChannelProtocolResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelProtocolResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthChannelProtocolResponseProviderDate) SetStartdate(v string) *BandwidthChannelProtocolResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProviderDate) SetEnddate(v string) *BandwidthChannelProtocolResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProviderDate) SetChannel(v *BandwidthChannelProtocolResponseProviderDateChannel) *BandwidthChannelProtocolResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthChannelProtocolResponseProviderDateChannel struct {
  // {"en":"channel","zh_CN":"频道"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"频道带宽区分协议带宽数据"}
  Bandwidth []*BandwidthChannelProtocolResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthChannelProtocolResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthChannelProtocolResponseProviderDateChannel) SetName(v string) *BandwidthChannelProtocolResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProviderDateChannel) SetBandwidth(v []*BandwidthChannelProtocolResponseProviderDateChannelBandwidth) *BandwidthChannelProtocolResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthChannelProtocolResponseProviderDateChannelBandwidth struct     {
  // {"en":"time","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"total bandwidth of http and https,unit Mbps.This value only displays When parameter 'datatype' is not specified or includes '1'","zh_CN":"http+https的带宽，单位：Mbps（当入参datatype为空或者包含1时有值）"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"http bandwidth,unit Mbps.This value only displays when parameter 'datatype' includes '2'","zh_CN":"http类型的带宽，单位：Mbps(当入参datatype包含2时有值)"}
  Http *string `json:"http,omitempty" xml:"http,omitempty" require:"true"`
  // {"en":"https bandwidth,unit Mbps.This value only displays when parameter 'datatype' includes '3'","zh_CN":"https类型的带宽，单位：Mbps（当入参datatype包含3时有值）"}
  Https *string `json:"https,omitempty" xml:"https,omitempty" require:"true"`
}

func (s BandwidthChannelProtocolResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthChannelProtocolResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthChannelProtocolResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProviderDateChannelBandwidth) SetTotal(v string) *BandwidthChannelProtocolResponseProviderDateChannelBandwidth {
  s.Total = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProviderDateChannelBandwidth) SetHttp(v string) *BandwidthChannelProtocolResponseProviderDateChannelBandwidth {
  s.Http = &v
  return s
}

func (s *BandwidthChannelProtocolResponseProviderDateChannelBandwidth) SetHttps(v string) *BandwidthChannelProtocolResponseProviderDateChannelBandwidth {
  s.Https = &v
  return s
}

type BandwidthChannelProtocolResponseHeader struct {
}

func (s BandwidthChannelProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelProtocolResponseHeader) GoString() string {
  return s.String()
}




type ReportP2pBandwidthDomainServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH:mm:SS+08:00, for example, 2021-05-19T10:00:00+08:00 (10:00:00 Beijing time on May 19, 2021);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒)
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 1 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days. ", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support).", "zh_CN":"1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.如未填,则默认查询此账号下所有S-P2P的加速域名"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s ReportP2pBandwidthDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportP2pBandwidthDomainServiceRequest) SetDateFrom(v string) *ReportP2pBandwidthDomainServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceRequest) SetDateTo(v string) *ReportP2pBandwidthDomainServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceRequest) SetDomain(v []*string) *ReportP2pBandwidthDomainServiceRequest {
  s.Domain = v
  return s
}

type ReportP2pBandwidthDomainServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportP2pBandwidthDomainServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportP2pBandwidthDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportP2pBandwidthDomainServiceResponse) SetCode(v string) *ReportP2pBandwidthDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceResponse) SetMessage(v string) *ReportP2pBandwidthDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceResponse) SetData(v []*ReportP2pBandwidthDomainServiceResponseData) *ReportP2pBandwidthDomainServiceResponse {
  s.Data = v
  return s
}

type ReportP2pBandwidthDomainServiceResponseData struct     {
  DomainList []*ReportP2pBandwidthDomainServiceResponseDataDomainList `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportP2pBandwidthDomainServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportP2pBandwidthDomainServiceResponseData) SetDomainList(v []*ReportP2pBandwidthDomainServiceResponseDataDomainList) *ReportP2pBandwidthDomainServiceResponseData {
  s.DomainList = v
  return s
}

type ReportP2pBandwidthDomainServiceResponseDataDomainList struct     {
  // {"en":"1. If the accelerated domain name is a generic domain name, this field is the detailed domain name of the generic domain name (a generic domain name may have many detailed domain names, and detailed data of each detailed domain name will be returned).
  // 						  2. If the accelerated domain name queried is not a generic domain name but a precise domain name, then this field displays the same domain name as the accelerated domain name.", "zh_CN":"1.如果加速域名是泛域名,则此字段为泛域名的明细域名(一个泛域名可能会有很多个明细域名,则会返回每个明细域名的详细数据)。
  // 						  2.如果查询的加速域名非泛域名,而是精确域名,则此字段展示同加速域名。"}
  VDomain *string `json:"vDomain,omitempty" xml:"vDomain,omitempty" require:"true"`
  BandwidthList []*ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList `json:"bandwidthList,omitempty" xml:"bandwidthList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportP2pBandwidthDomainServiceResponseDataDomainList) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceResponseDataDomainList) GoString() string {
  return s.String()
}

func (s *ReportP2pBandwidthDomainServiceResponseDataDomainList) SetVDomain(v string) *ReportP2pBandwidthDomainServiceResponseDataDomainList {
  s.VDomain = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceResponseDataDomainList) SetBandwidthList(v []*ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) *ReportP2pBandwidthDomainServiceResponseDataDomainList {
  s.BandwidthList = v
  return s
}

type ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList struct     {
  // {"en":"Timestamp, returns the time slice containing the start time and end time. Time format: yyyy-MM-dd HH:mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。时间格式:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"The bandwidth value of the CDN requested by the SDK, in Mbps, with 2 decimal places", "zh_CN":"SDK请求CDN的带宽值,单位Mbps,保留2位小数。"}
  CdnValue *string `json:"cdnValue,omitempty" xml:"cdnValue,omitempty" require:"true"`
  // {"en":"P2P bandwidth value, unit Mbps, keep 2 decimal places", "zh_CN":"P2P带宽值,单位Mbps,保留2位小数。"}
  P2pValue *string `json:"p2pValue,omitempty" xml:"p2pValue,omitempty" require:"true"`
  // {"en":"IPv6 bandwidth in P2P bandwidth, unit Mbps, keep 2 decimal places", "zh_CN":"P2P带宽中的IPv6带宽,单位Mbps,保留2位小数。"}
  P2pIpv6Value *string `json:"p2pIpv6Value,omitempty" xml:"p2pIpv6Value,omitempty" require:"true"`
}

func (s ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) GoString() string {
  return s.String()
}

func (s *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) SetTimestamp(v string) *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList {
  s.Timestamp = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) SetCdnValue(v string) *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList {
  s.CdnValue = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) SetP2pValue(v string) *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList {
  s.P2pValue = &v
  return s
}

func (s *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList) SetP2pIpv6Value(v string) *ReportP2pBandwidthDomainServiceResponseDataDomainListBandwidthList {
  s.P2pIpv6Value = &v
  return s
}

type ReportP2pBandwidthDomainServicePaths struct {
}

func (s ReportP2pBandwidthDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServicePaths) GoString() string {
  return s.String()
}

type ReportP2pBandwidthDomainServiceParameters struct {
}

func (s ReportP2pBandwidthDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceParameters) GoString() string {
  return s.String()
}

type ReportP2pBandwidthDomainServiceRequestHeader struct {
}

func (s ReportP2pBandwidthDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportP2pBandwidthDomainServiceResponseHeader struct {
}

func (s ReportP2pBandwidthDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportP2pBandwidthDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthofOriginminutelyRequest struct {
  // {"en":"Start Time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. Not later than the current time;\n3. Up to 6 months (183 days) of data are available.","zh_CN":"开始时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.不能大于当前时间；\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. the end time should be greater than the start time. if the end time is greater than the current time, take the current time;\n3. dateFrom, dateTo, both are not sent, default query past 24 hours; If only one is not sent, throw exception;\n4. Allow maximum query interval: 7 days, i.e., 7 days between dateFrom and dateTo. Do not exceed 7 days.","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间；\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；\n4.允许查询最大时间间隔：7天，即dateFrom和dateTo相差不能超过7天，不支持调整"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"defaultValue":"5m","en":"Data granularity:\n1. Support for 1m (1 minute), 5m (5 minutes), 1h (1 hour);\n2. do not pass default 5m.\nData granularity, default to 5m","zh_CN":"数据粒度：\n1.支持1m（1分钟）、5m（5分钟）、1h（1小时）；\n2.不传默认5m。\n\n数据粒度，默认为5m","exampleValue":"1m,5m,1h"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Domain name:\n1. the maximum number of transitive domain names is 50 by default;\n2. Automatically filter out illegal domain names (e.g. passing illegal domain names will be filtered out, and the search results will only return the legal domain name data)\n3. If left blank, all domain names will be obtained. If the total number of domain names exceeds the upper limit, an error will be reported.","zh_CN":"域名：\n1.可传递域名数量上限默认为50（可联系技术支持调整）；\n2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。\n3. 若未填写默认查询全部域名，全部域名超出域名上限报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration zone:\n1. do not pass the default query to all areas;\n2. currently only the leaflet area is supported externally;\n3. optional values: cn, apac, am, emea.","zh_CN":"加速区域：\n1.不传默认查询全部区域；\n2.目前对外只支持传单区域；\n3.可选值：cn、apac、am、emea。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"grouped dimension:\n1. the optional value is domain;\n2. If incoming data is shown in accordance with the dimension.","zh_CN":"分组维度：\n1.可选值为domain；\n2.有传入则按照该维度展示明细数据。"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryBandwidthofOriginminutelyRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyRequest) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDateFrom(v string) *QueryBandwidthofOriginminutelyRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDateTo(v string) *QueryBandwidthofOriginminutelyRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDataInterval(v string) *QueryBandwidthofOriginminutelyRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDomain(v []*string) *QueryBandwidthofOriginminutelyRequest {
  s.Domain = v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetRegion(v string) *QueryBandwidthofOriginminutelyRequest {
  s.Region = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetGroupBy(v string) *QueryBandwidthofOriginminutelyRequest {
  s.GroupBy = &v
  return s
}

type QueryBandwidthofOriginminutelyRequestHeader struct {
}

func (s QueryBandwidthofOriginminutelyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthofOriginminutelyPaths struct {
}

func (s QueryBandwidthofOriginminutelyPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyPaths) GoString() string {
  return s.String()
}

type QueryBandwidthofOriginminutelyParameters struct {
}

func (s QueryBandwidthofOriginminutelyParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyParameters) GoString() string {
  return s.String()
}

type QueryBandwidthofOriginminutelyResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on request results","zh_CN":"请求结果的详细数据"}
  Data []*QueryBandwidthofOriginminutelyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthofOriginminutelyResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyResponse) SetCode(v string) *QueryBandwidthofOriginminutelyResponse {
  s.Code = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponse) SetMessage(v string) *QueryBandwidthofOriginminutelyResponse {
  s.Message = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponse) SetData(v []*QueryBandwidthofOriginminutelyResponseData) *QueryBandwidthofOriginminutelyResponse {
  s.Data = v
  return s
}

type QueryBandwidthofOriginminutelyResponseData struct     {
  // {"en":"Domain","zh_CN":"域名信息"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak bandwidth Mbps, example (931556.21 Mbps)","zh_CN":"峰值带宽 Mbps，示例 （931556.21 Mbps）"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Peak Time, Example (2019-02-13 18:01)","zh_CN":"峰值时间，示例(2019-02-13 18:01)"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Total return flow, example ( 74099.92 MB )","zh_CN":"回源总流量，示例 ( 74099.92 MB )"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  OriginBandwidthData []*QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData `json:"originBandwidthData,omitempty" xml:"originBandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthofOriginminutelyResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponseData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetDomain(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.Domain = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetPeakValue(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetPeakTime(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetTotal(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.Total = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetOriginBandwidthData(v []*QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) *QueryBandwidthofOriginminutelyResponseData {
  s.OriginBandwidthData = v
  return s
}

type QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData struct     {
  // {"en":"The granularity of data is 1 minute, and the format is yyyy-MM-dd HH:MM.","zh_CN":"数据粒度为1分钟，格式为yyyy-MM-dd HH:mm。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Return the source bandwidth value, in Mbps, 2 decimal places reserved.","zh_CN":"回源带宽值，单位Mbps，保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) SetTimestamp(v string) *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) SetValue(v string) *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData {
  s.Value = &v
  return s
}

type QueryBandwidthofOriginminutelyResponseHeader struct {
}

func (s QueryBandwidthofOriginminutelyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponseHeader) GoString() string {
  return s.String()
}




type GetCdnRelayTrafficRequest struct {
  // {"en":"Start Time\n1. The time format is 'yyyy-MM-dd'.\n2. Cannot exceed the current date. \n3. Up to the past 730 days of data can be obtained.","zh_CN":"开始时间\n1. 时间格式为'yyyy-MM-dd'；\n2. 不能大于当前日期；\n3. 最多可获取最近730天的数据；"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time\n1. The time format is 'yyyy-MM-dd'.\n2. The end time must be greater than the start time.\n3. If the end time greater than the current time, the current time is taken.\n4. DateFrom and dateTo are both not provided, defaulting to query the current time; if only one is not provided, throw an exception.\n5. Maximum query interval allowed: 30 days, that is the range between dateFrom and dateTo can not exceed 30 days.","zh_CN":"结束时间\n1. 时间格式为'yyyy-MM-dd'；\n2. 结束时间需大于等于开始时间；\n3. 结束时间如果大于当前时间，取当前时间；\n4. dateFrom，dateTo二者都未传，默认查询当天；如仅有一个未传，抛异常；\n5. 允许查询最大间隔：30天，即dateFrom和dateTo相差不能超过30天；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"1. Specify the time zone for both the query time range(dateFrom/dateTo) and the returned data. \n2. Greenwich Mean Time Zone, the default time zone is GMT+08:00.\n3. If you wish to specify a different time zone, you can append a time zone identifier, i.e. 'GMT+09:00' or 'GMT-10:00'.","zh_CN":"1. 指定查询时间(dateFrom/dateTo)和返回数据的时区；\n2. 格林尼治时区，默认时区是GMT+08:00；\n3. 如果您希望指定不同的时区，可以附加时区标识，即'GMT+09:00'或'GMT-10:00'；"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"defaultValue":"up","en":"Rounds up or down a timestamp by a given time period. Only valid when the granularity is less than 1d(not including 1day) \n1. 'up' denotes rounds up, eg: granularity=5m then 00:00:00 ~ 00:05:00 will be displayed as 00:05:00.\n2. 'down' denotes rounds down, eg: granularity=5m then 00:00:00 ~ 00:05:00 will be displayed as 00:00:00.\n3. if not specified, the result will be rounded up.","zh_CN":"根据指定的时间周期对时间戳进行向上或向下取整。仅当粒度小于1天（不包含1天）时有效。 \n1. 'up'表示向上取整。例如：granularity=5m，00:00:00~00:05:00将会显示为 00:05:00；\n2. 'down'表示向下取整，例如：granularity=5m时，00:00:00~00:05:00将显示为00:00:00； \n3. 如未传值，结果将进行向上取整(up)；","exampleValue":"up,down"}
  TimeRounding *string `json:"timeRounding,omitempty" xml:"timeRounding,omitempty"`
  // {"en":"Domains\n1. Domain is not uploaded: Query all domain names of the account (More than 2000 domains will encounter error, you can contact technical support for adjustment). \n2. Domain is uploaded: Up to 2000 domains are supported (you can contact technical support for adjustment).\n3. For multiple domain, please separate them with a semicolon ';'.","zh_CN":"域名\n1. 未传递domain时：查询账号下所有全部域名(域名超过2000个则报错，可联系技术支持调整)；\n2. 有传递domain时：域名最多支持传2000个（可联系技术支持调整）；\n3. 多个域名用英文逗号';'分割；"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Service type of the acceleration domain \n1. If not specified, it is considered as 'no restriction on service type.\n2. For multiple Application server types, please separate them with a semicolon ';'. eg: 'web,web-https'","zh_CN":"加速域名的服务类型\n1. 未传递视为不限服务类型；\n2. 多个服务类型请使用英文分号';'分隔；eg: 'web;web-https'"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"billing region of the Acceleration domain\n1. If not specified, it is considered as 'no restriction on billing region.\n2. For multiple billing regions, please separate them with a semicolon ';'. eg: 'cn;kr'","zh_CN":"加速域名的计费区域\n1. 未传递视为不限计费区域；\n2. 多个计费区域请用英文分号';'分隔；如：'cn;hk'"}
  BillingRegion *string `json:"billingRegion,omitempty" xml:"billingRegion,omitempty"`
  // {"defaultValue":"1d","en":"time granularity: the default granularity is one day. \n- 5m: five minutes \n- 1h: one hour \n- 1d: one day","zh_CN":"数据粒度：默认1天粒度 \n- 5m：5分钟粒度 \n- 1h：1小时粒度 \n- 1d：1天粒度","exampleValue":"5m,1h,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Group keywords\n1. If not specified, result will be aggregated as default.\n2. Passing in a keyword means that the details need to be displayed according to the values corresponding to the keyword dimension grouping (for example, passing in 'domain' means that the details are expanded according to domain). If an invalid value is specified, it will encounter an error.\n3. Support passing multiple values, for multiple value please separate them with a semicolon ';', only support 'domain' grouping currently.","zh_CN":"分组关键词\n1. 未传递时，默认聚合展示；\n2. 传入关键词则代表需要按照关键词维度分组对应的值展示明细（例如传domain，则代表返回按照domain明细展开），如传入不支持的关键词，返回相应错误提示；\n3. 支持传多个值，传多个请用英文分号';'分隔，当前只支持'domain'；","exampleValue":"domain"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s GetCdnRelayTrafficRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficRequest) GoString() string {
  return s.String()
}

func (s *GetCdnRelayTrafficRequest) SetDateFrom(v string) *GetCdnRelayTrafficRequest {
  s.DateFrom = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetDateTo(v string) *GetCdnRelayTrafficRequest {
  s.DateTo = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetTimezone(v string) *GetCdnRelayTrafficRequest {
  s.Timezone = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetTimeRounding(v string) *GetCdnRelayTrafficRequest {
  s.TimeRounding = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetDomain(v string) *GetCdnRelayTrafficRequest {
  s.Domain = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetServiceType(v string) *GetCdnRelayTrafficRequest {
  s.ServiceType = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetBillingRegion(v string) *GetCdnRelayTrafficRequest {
  s.BillingRegion = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetGranularity(v string) *GetCdnRelayTrafficRequest {
  s.Granularity = &v
  return s
}

func (s *GetCdnRelayTrafficRequest) SetGroupBy(v string) *GetCdnRelayTrafficRequest {
  s.GroupBy = &v
  return s
}

type GetCdnRelayTrafficRequestHeader struct {
}

func (s GetCdnRelayTrafficRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficRequestHeader) GoString() string {
  return s.String()
}

type GetCdnRelayTrafficPaths struct {
}

func (s GetCdnRelayTrafficPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficPaths) GoString() string {
  return s.String()
}

type GetCdnRelayTrafficParameters struct {
}

func (s GetCdnRelayTrafficParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficParameters) GoString() string {
  return s.String()
}

type GetCdnRelayTrafficResponse struct {
  // {"en":"request status code","zh_CN":"请求状态"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request status description","zh_CN":"请求状态描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"具体响应内容"}
  Data *GetCdnRelayTrafficResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCdnRelayTrafficResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficResponse) GoString() string {
  return s.String()
}

func (s *GetCdnRelayTrafficResponse) SetCode(v string) *GetCdnRelayTrafficResponse {
  s.Code = &v
  return s
}

func (s *GetCdnRelayTrafficResponse) SetMessage(v string) *GetCdnRelayTrafficResponse {
  s.Message = &v
  return s
}

func (s *GetCdnRelayTrafficResponse) SetData(v *GetCdnRelayTrafficResponseData) *GetCdnRelayTrafficResponse {
  s.Data = v
  return s
}

type GetCdnRelayTrafficResponseData struct {
  // {"en":"container for data information.","zh_CN":"数据信息的容器"}
  Report *GetCdnRelayTrafficResponseDataReport `json:"report,omitempty" xml:"report,omitempty" require:"true" type:"Struct"`
}

func (s GetCdnRelayTrafficResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficResponseData) GoString() string {
  return s.String()
}

func (s *GetCdnRelayTrafficResponseData) SetReport(v *GetCdnRelayTrafficResponseDataReport) *GetCdnRelayTrafficResponseData {
  s.Report = v
  return s
}

type GetCdnRelayTrafficResponseDataReport struct {
  // {"defaultValue":"cdn-traffic-report-data","en":"The metric name","zh_CN":"接口数据名称"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Indicates the starting time of response data. The time format is yyyy-MM-dd. ","zh_CN":"标记返回数据的开始时间，格式为 yyyy-MM-dd."}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"Indicates the ending time of response data. The time format is yyyy-MM-dd. ","zh_CN":"标记返回数据的结束时间，格式为 yyyy-MM-dd."}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"defaultValue":"GMT+08:00","en":"Indicates the time zone for the returned data. Greenwich Mean Time Zone, the format is GMT+08:00","zh_CN":"标记返回数据的时区，格林尼治时区，格式是GMT+08:00"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true"`
  // {"en":"List of data items. Each item contains fields as following: 'domain', 'flow', 'peakBandwidth', 'peakTime' ","zh_CN":"数据项集合，每个数据项包含：domain，flow，peakBandwidth，peakTime字段"}
  Groups []*GetCdnRelayTrafficResponseDataReportGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnRelayTrafficResponseDataReport) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficResponseDataReport) GoString() string {
  return s.String()
}

func (s *GetCdnRelayTrafficResponseDataReport) SetType(v string) *GetCdnRelayTrafficResponseDataReport {
  s.Type = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReport) SetDateFrom(v string) *GetCdnRelayTrafficResponseDataReport {
  s.DateFrom = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReport) SetDateTo(v string) *GetCdnRelayTrafficResponseDataReport {
  s.DateTo = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReport) SetTimezone(v string) *GetCdnRelayTrafficResponseDataReport {
  s.Timezone = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReport) SetGroups(v []*GetCdnRelayTrafficResponseDataReportGroups) *GetCdnRelayTrafficResponseDataReport {
  s.Groups = v
  return s
}

type GetCdnRelayTrafficResponseDataReportGroups struct     {
  // {"en":"the name of domain","zh_CN":"频道"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Indicates the CDN-to-Client traffic volume in gigabytes. The value is accurate to 3 decimal places.","zh_CN":"CDN到客户端的流量，单位：GB，精确到小数点后3位。"}
  TotalRelayTraffic *string `json:"totalRelayTraffic,omitempty" xml:"totalRelayTraffic,omitempty" require:"true"`
  // {"en":"Indicates the peak bandwidth(5-minute granularity) of CDN-to-Client traffic volume in Mbps. 1.The value is accurate to 3 decimal places. 2.If granularity<=5m, no display in groups container objects","zh_CN":"带宽峰值(5分钟粒度），单位Mbps 1、精确到小数点后3位 2、如果granularity<=5m，则在[groups]分组容器对象中不显示"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"Indicates the time of peak bandwidth(5-minute granularity). ","zh_CN":"峰值时间(5分钟粒度）"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"List of data items.Each item contains fields as following: 'time', 'flow'","zh_CN":"明细数据列表，每项包含time和flow字段"}
  Item []*GetCdnRelayTrafficResponseDataReportGroupsItem `json:"item,omitempty" xml:"item,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnRelayTrafficResponseDataReportGroups) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficResponseDataReportGroups) GoString() string {
  return s.String()
}

func (s *GetCdnRelayTrafficResponseDataReportGroups) SetDomain(v string) *GetCdnRelayTrafficResponseDataReportGroups {
  s.Domain = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReportGroups) SetTotalRelayTraffic(v string) *GetCdnRelayTrafficResponseDataReportGroups {
  s.TotalRelayTraffic = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReportGroups) SetPeakBandwidth(v string) *GetCdnRelayTrafficResponseDataReportGroups {
  s.PeakBandwidth = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReportGroups) SetPeakTime(v string) *GetCdnRelayTrafficResponseDataReportGroups {
  s.PeakTime = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReportGroups) SetItem(v []*GetCdnRelayTrafficResponseDataReportGroupsItem) *GetCdnRelayTrafficResponseDataReportGroups {
  s.Item = v
  return s
}

type GetCdnRelayTrafficResponseDataReportGroupsItem struct     {
  // {"en":"Indicates the date and time 1. If the type is \"5m\", the timestamp format is \"YYYY-MM-DD hh:mm:00\", that means the period's end time. 2. If the type is \"1h\", the timestamp format is \"YYYY-MM-DD hh:00:00\", that means the period's start time. 3. If the type is \"1d\", the timestamp format is \"YYYY-MM-DD 00:00:00\", that means the period's start time.","zh_CN":"表示日期和时间 1、如果类型为“5m”，则时间戳格式为“YYYY-MM-DD hh:mm:00”，表示该时间段的结束时间。 2、如果类型为“1h”，则时间戳格式为“YYYY-MM-DD hh:00:00”，表示该时间段的开始时间。 3、如果类型为“1d”，则时间戳格式为“YYYY-MM-DD 00:00:00”，表示该时间段的开始时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Indicates the CDN-to-Client traffic volume in gigabytes. The value is accurate to 3 decimal places.","zh_CN":"CDN到客户端的流量，单位：GB，精确到小数点后3位。"}
  RelayTraffic *string `json:"relayTraffic,omitempty" xml:"relayTraffic,omitempty" require:"true"`
}

func (s GetCdnRelayTrafficResponseDataReportGroupsItem) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficResponseDataReportGroupsItem) GoString() string {
  return s.String()
}

func (s *GetCdnRelayTrafficResponseDataReportGroupsItem) SetTime(v string) *GetCdnRelayTrafficResponseDataReportGroupsItem {
  s.Time = &v
  return s
}

func (s *GetCdnRelayTrafficResponseDataReportGroupsItem) SetRelayTraffic(v string) *GetCdnRelayTrafficResponseDataReportGroupsItem {
  s.RelayTraffic = &v
  return s
}

type GetCdnRelayTrafficResponseHeader struct {
}

func (s GetCdnRelayTrafficResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCdnRelayTrafficResponseHeader) GoString() string {
  return s.String()
}




type BandwidthChannelRequest struct {
  // {"en":"Specifies the query date:\n1.With format yyyy-mm-dd.\n2.If not Specifies,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.\n2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.\n2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1.If there are multiple inputs,use  ';' as separator.\n2.If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2.If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.\n1.If there are multiple inputs,use ';' as separator.\n2.If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1.optional values:xml, json.\n2.'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:\n1.'true' as default.\n2. If not 'true',it will query data of channels that ends with any item of input 'channel's.","zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=ispCode","en":"1.If there isp multiple inputs,use ';' as demimeter.\n2.optional values of isp: refers to the ISP-section of appendix.\n3. If not specified,means all the isp.","zh_CN":"&nbsp;要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"Display statistic result in merged or separate way.\n1.If specified 1,get the merged result.\n2.If specified 2,get the separate result.\n3.If specified 3,get both merged result and separate result.\n4.If not specified,means '1'.","zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"If return the flow details:\nchoose 1: Return\nchoose 0: Not return\nthe default is 0","zh_CN":"needFlow 是否需要返回流量明细，1：需要；0：不需要。默认为0."}
  NeedFlow *string `json:"needFlow,omitempty" xml:"needFlow,omitempty"`
  // {"en":"provide 'flowInfo' will displays bandwidth peak, peak time, and total flow information;","zh_CN":"填写'flowInfo'时，展示带宽峰值,峰值时间,总流量信息."}
  OptionalFields *string `json:"optionalFields,omitempty" xml:"optionalFields,omitempty"`
  // {"en":"Greenwich Mean Time, parameter format GMT%2b09:00 represents Eastern 9th Zone, GMT-09:00 represents Western 9th Zone, if not specified, it defaults to local time zone (Eastern 8th Zone).","zh_CN":"格林尼治时区，参数格式 GMT%2b09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区(东八区)"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Whether to aggregate according to a specific method. Format: number_day|hour. For example, '3_hour' means aggregation by 3 hours; '2_day' means aggregation by 2 days.","zh_CN":"是否按照特定方式聚合,格式 :  数字_day|hour .例如 3_hour表示按照3小时聚合; 2_day表示按照2天聚合"}
  ReturnType *string `json:"returnType,omitempty" xml:"returnType,omitempty"`
}

func (s BandwidthChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelRequest) GoString() string {
  return s.String()
}

func (s *BandwidthChannelRequest) SetDate(v string) *BandwidthChannelRequest {
  s.Date = &v
  return s
}

func (s *BandwidthChannelRequest) SetStartdate(v string) *BandwidthChannelRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelRequest) SetEnddate(v string) *BandwidthChannelRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelRequest) SetChannel(v string) *BandwidthChannelRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthChannelRequest) SetRegion(v string) *BandwidthChannelRequest {
  s.Region = &v
  return s
}

func (s *BandwidthChannelRequest) SetAccetype(v string) *BandwidthChannelRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthChannelRequest) SetDataformat(v string) *BandwidthChannelRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthChannelRequest) SetIsExactMatch(v string) *BandwidthChannelRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthChannelRequest) SetIsp(v string) *BandwidthChannelRequest {
  s.Isp = &v
  return s
}

func (s *BandwidthChannelRequest) SetResultType(v string) *BandwidthChannelRequest {
  s.ResultType = &v
  return s
}

func (s *BandwidthChannelRequest) SetNeedFlow(v string) *BandwidthChannelRequest {
  s.NeedFlow = &v
  return s
}

func (s *BandwidthChannelRequest) SetOptionalFields(v string) *BandwidthChannelRequest {
  s.OptionalFields = &v
  return s
}

func (s *BandwidthChannelRequest) SetTimezone(v string) *BandwidthChannelRequest {
  s.Timezone = &v
  return s
}

func (s *BandwidthChannelRequest) SetReturnType(v string) *BandwidthChannelRequest {
  s.ReturnType = &v
  return s
}

type BandwidthChannelRequestHeader struct {
}

func (s BandwidthChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelRequestHeader) GoString() string {
  return s.String()
}

type BandwidthChannelPaths struct {
}

func (s BandwidthChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelPaths) GoString() string {
  return s.String()
}

type BandwidthChannelParameters struct {
}

func (s BandwidthChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelParameters) GoString() string {
  return s.String()
}

type BandwidthChannelResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *BandwidthChannelResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponse) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponse) SetProvider(v *BandwidthChannelResponseProvider) *BandwidthChannelResponse {
  s.Provider = v
  return s
}

type BandwidthChannelResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"resultType","zh_CN":"统计类型"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {"en":"data","zh_CN":"频道带宽数据"}
  Date *BandwidthChannelResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProvider) SetName(v string) *BandwidthChannelResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthChannelResponseProvider) SetType(v string) *BandwidthChannelResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthChannelResponseProvider) SetResultType(v string) *BandwidthChannelResponseProvider {
  s.ResultType = &v
  return s
}

func (s *BandwidthChannelResponseProvider) SetDate(v *BandwidthChannelResponseProviderDate) *BandwidthChannelResponseProvider {
  s.Date = v
  return s
}

type BandwidthChannelResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"peakValue","zh_CN":"峰值带宽,单位Mbps"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"peakTime","zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"totalFlow","zh_CN":"总流量,单位GB"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"channel","zh_CN":"频道"}
  Channel *BandwidthChannelResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDate) SetStartdate(v string) *BandwidthChannelResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetEnddate(v string) *BandwidthChannelResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetPeakValue(v string) *BandwidthChannelResponseProviderDate {
  s.PeakValue = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetPeakTime(v string) *BandwidthChannelResponseProviderDate {
  s.PeakTime = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetTotalFlow(v string) *BandwidthChannelResponseProviderDate {
  s.TotalFlow = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetChannel(v *BandwidthChannelResponseProviderDateChannel) *BandwidthChannelResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthChannelResponseProviderDateChannel struct {
  // {"en":"channel","zh_CN":"频道"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"带宽数据"}
  Bandwidth []*BandwidthChannelResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
  // {"en":"Traffic data, provided only when the parameter needFlow=1 is specified.","zh_CN":"流量数据，当入参needFlow=1时才会提供"}
  Flow []*BandwidthChannelResponseProviderDateChannelFlow `json:"flow,omitempty" xml:"flow,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthChannelResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDateChannel) SetName(v string) *BandwidthChannelResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthChannelResponseProviderDateChannel) SetBandwidth(v []*BandwidthChannelResponseProviderDateChannelBandwidth) *BandwidthChannelResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

func (s *BandwidthChannelResponseProviderDateChannel) SetFlow(v []*BandwidthChannelResponseProviderDateChannelFlow) *BandwidthChannelResponseProviderDateChannel {
  s.Flow = v
  return s
}

type BandwidthChannelResponseProviderDateChannelBandwidth struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"带宽,单位Mbps"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthChannelResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthChannelResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthChannelResponseProviderDateChannelBandwidth) SetText(v string) *BandwidthChannelResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type BandwidthChannelResponseProviderDateChannelFlow struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Traffic, measured in MB.","zh_CN":"流量，单位MB"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthChannelResponseProviderDateChannelFlow) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDateChannelFlow) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDateChannelFlow) SetTime(v string) *BandwidthChannelResponseProviderDateChannelFlow {
  s.Time = &v
  return s
}

func (s *BandwidthChannelResponseProviderDateChannelFlow) SetText(v string) *BandwidthChannelResponseProviderDateChannelFlow {
  s.Text = &v
  return s
}

type BandwidthChannelResponseHeader struct {
}

func (s BandwidthChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseHeader) GoString() string {
  return s.String()
}




type BillingOrderRequest struct {
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not specified,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope.\n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope.\n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1)If there are multiple inputs,use  ';' as separator.\n2)If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号“;”分隔开，如查询大陆及亚太区域，参数填写为：“region=cn;apac”。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Contract Code. If there are multiple inputs,use ';' as separator.","zh_CN":"合同号，如contracCode=SR190001，多个请用英文分号';'分隔开"}
  ContractCode *string `json:"contractCode,omitempty" xml:"contractCode,omitempty"`
  // {"en":"Show the details of order or not. Optional valuse: true, false.\n'false' as defaullt.","zh_CN":"showDetail=true 显示订单详情，showDetail=false不显示订单详情，默认值是false。"}
  ShowDetail *string `json:"showDetail,omitempty" xml:"showDetail,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Filter the billing start and end time of formal orders","zh_CN":"是否对正式订单的计费起止时间进行过滤。payDateFilter=true表示要过滤，即正式订单只返回计费起止时间与查询时间有交集的结果；payDateFilter=false表示不过滤，默认false；测试订单不受此参数影响"}
  PayDateFilter *string `json:"payDateFilter,omitempty" xml:"payDateFilter,omitempty"`
}

func (s BillingOrderRequest) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderRequest) GoString() string {
  return s.String()
}

func (s *BillingOrderRequest) SetDate(v string) *BillingOrderRequest {
  s.Date = &v
  return s
}

func (s *BillingOrderRequest) SetStartdate(v string) *BillingOrderRequest {
  s.Startdate = &v
  return s
}

func (s *BillingOrderRequest) SetEnddate(v string) *BillingOrderRequest {
  s.Enddate = &v
  return s
}

func (s *BillingOrderRequest) SetChannel(v string) *BillingOrderRequest {
  s.Channel = &v
  return s
}

func (s *BillingOrderRequest) SetRegion(v string) *BillingOrderRequest {
  s.Region = &v
  return s
}

func (s *BillingOrderRequest) SetContractCode(v string) *BillingOrderRequest {
  s.ContractCode = &v
  return s
}

func (s *BillingOrderRequest) SetShowDetail(v string) *BillingOrderRequest {
  s.ShowDetail = &v
  return s
}

func (s *BillingOrderRequest) SetDataformat(v string) *BillingOrderRequest {
  s.Dataformat = &v
  return s
}

func (s *BillingOrderRequest) SetPayDateFilter(v string) *BillingOrderRequest {
  s.PayDateFilter = &v
  return s
}

type BillingOrderRequestHeader struct {
}

func (s BillingOrderRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderRequestHeader) GoString() string {
  return s.String()
}

type BillingOrderPaths struct {
}

func (s BillingOrderPaths) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderPaths) GoString() string {
  return s.String()
}

type BillingOrderParameters struct {
}

func (s BillingOrderParameters) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderParameters) GoString() string {
  return s.String()
}

type BillingOrderResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *BillingOrderResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BillingOrderResponse) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponse) GoString() string {
  return s.String()
}

func (s *BillingOrderResponse) SetProvider(v *BillingOrderResponseProvider) *BillingOrderResponse {
  s.Provider = v
  return s
}

type BillingOrderResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"订单计费值"}
  Date *BillingOrderResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BillingOrderResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponseProvider) GoString() string {
  return s.String()
}

func (s *BillingOrderResponseProvider) SetName(v string) *BillingOrderResponseProvider {
  s.Name = &v
  return s
}

func (s *BillingOrderResponseProvider) SetType(v string) *BillingOrderResponseProvider {
  s.Type = &v
  return s
}

func (s *BillingOrderResponseProvider) SetDate(v *BillingOrderResponseProviderDate) *BillingOrderResponseProvider {
  s.Date = v
  return s
}

type BillingOrderResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"channel","zh_CN":"订单计费值"}
  Order []*BillingOrderResponseProviderDateOrder `json:"order,omitempty" xml:"order,omitempty" require:"true" type:"Repeated"`
}

func (s BillingOrderResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BillingOrderResponseProviderDate) SetStartdate(v string) *BillingOrderResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BillingOrderResponseProviderDate) SetEnddate(v string) *BillingOrderResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BillingOrderResponseProviderDate) SetOrder(v []*BillingOrderResponseProviderDateOrder) *BillingOrderResponseProviderDate {
  s.Order = v
  return s
}

type BillingOrderResponseProviderDateOrder struct     {
  // {"en":"id","zh_CN":"订单ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"orderInfo","zh_CN":"订单信息"}
  OrderInfo *string `json:"orderInfo,omitempty" xml:"orderInfo,omitempty" require:"true"`
  // {"en":"chargeMethod","zh_CN":"计费方式"}
  ChargeMethod *string `json:"chargeMethod,omitempty" xml:"chargeMethod,omitempty" require:"true"`
  // {"en":"chargeExplanation","zh_CN":"计费说明"}
  ChargeExplanation *string `json:"chargeExplanation,omitempty" xml:"chargeExplanation,omitempty" require:"true"`
  // {"en":"regionAndIsp","zh_CN":"区域和ISP"}
  RegionAndIsp *string `json:"regionAndIsp,omitempty" xml:"regionAndIsp,omitempty" require:"true"`
  // {"en":"acceType","zh_CN":"加速类型"}
  AcceType *string `json:"acceType,omitempty" xml:"acceType,omitempty" require:"true"`
  // {"en":"chargeValue","zh_CN":"计费值"}
  ChargeValue *string `json:"chargeValue,omitempty" xml:"chargeValue,omitempty" require:"true"`
  // {"en":"chargeCategory","zh_CN":"费用类别"}
  ChargeCategory *string `json:"chargeCategory,omitempty" xml:"chargeCategory,omitempty" require:"true"`
  // {"en":"detailInfo","zh_CN":"计费单位"}
  DetailInfo *string `json:"detailInfo,omitempty" xml:"detailInfo,omitempty" require:"true"`
  // {"en":"channel","zh_CN":"频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"带宽数据"}
  Type *BillingOrderResponseProviderDateOrderType `json:"type,omitempty" xml:"type,omitempty" require:"true" type:"Struct"`
}

func (s BillingOrderResponseProviderDateOrder) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponseProviderDateOrder) GoString() string {
  return s.String()
}

func (s *BillingOrderResponseProviderDateOrder) SetId(v string) *BillingOrderResponseProviderDateOrder {
  s.Id = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetOrderInfo(v string) *BillingOrderResponseProviderDateOrder {
  s.OrderInfo = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetChargeMethod(v string) *BillingOrderResponseProviderDateOrder {
  s.ChargeMethod = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetChargeExplanation(v string) *BillingOrderResponseProviderDateOrder {
  s.ChargeExplanation = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetRegionAndIsp(v string) *BillingOrderResponseProviderDateOrder {
  s.RegionAndIsp = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetAcceType(v string) *BillingOrderResponseProviderDateOrder {
  s.AcceType = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetChargeValue(v string) *BillingOrderResponseProviderDateOrder {
  s.ChargeValue = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetChargeCategory(v string) *BillingOrderResponseProviderDateOrder {
  s.ChargeCategory = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetDetailInfo(v string) *BillingOrderResponseProviderDateOrder {
  s.DetailInfo = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetChannel(v string) *BillingOrderResponseProviderDateOrder {
  s.Channel = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrder) SetType(v *BillingOrderResponseProviderDateOrderType) *BillingOrderResponseProviderDateOrder {
  s.Type = v
  return s
}

type BillingOrderResponseProviderDateOrderType struct {
  // {"en":"name","zh_CN":"流量"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"带宽数据"}
  Detail []*BillingOrderResponseProviderDateOrderTypeDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
}

func (s BillingOrderResponseProviderDateOrderType) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponseProviderDateOrderType) GoString() string {
  return s.String()
}

func (s *BillingOrderResponseProviderDateOrderType) SetName(v string) *BillingOrderResponseProviderDateOrderType {
  s.Name = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrderType) SetDetail(v []*BillingOrderResponseProviderDateOrderTypeDetail) *BillingOrderResponseProviderDateOrderType {
  s.Detail = v
  return s
}

type BillingOrderResponseProviderDateOrderTypeDetail struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"带宽"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BillingOrderResponseProviderDateOrderTypeDetail) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponseProviderDateOrderTypeDetail) GoString() string {
  return s.String()
}

func (s *BillingOrderResponseProviderDateOrderTypeDetail) SetTime(v string) *BillingOrderResponseProviderDateOrderTypeDetail {
  s.Time = &v
  return s
}

func (s *BillingOrderResponseProviderDateOrderTypeDetail) SetText(v string) *BillingOrderResponseProviderDateOrderTypeDetail {
  s.Text = &v
  return s
}

type BillingOrderResponseHeader struct {
}

func (s BillingOrderResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BillingOrderResponseHeader) GoString() string {
  return s.String()
}




type GetBandwidthLogRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way.
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"The unit of the flow value in the details, the default is Mbps. Optional byte (byte) or bps", "zh_CN":"明细中的流量值单位，默认为Mbps。可选 byte (字节)或者bps"}
  FlowUnit *string `json:"flowUnit,omitempty" xml:"flowUnit,omitempty"`
  // {"en":"If specified 0,the time return as 00:05--24:00,If specified 1,the time return as 00:00--23:55,If not specified, the default value is 0", "zh_CN":"当值为0:返回00:05--24:00；当值为1:返回00:00--23:55。不传默认为:0"}
  TimeFromZero *string `json:"timeFromZero,omitempty" xml:"timeFromZero,omitempty"`
}

func (s GetBandwidthLogRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogRequest) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogRequest) SetCust(v string) *GetBandwidthLogRequest {
  s.Cust = &v
  return s
}

func (s *GetBandwidthLogRequest) SetDate(v string) *GetBandwidthLogRequest {
  s.Date = &v
  return s
}

func (s *GetBandwidthLogRequest) SetStartdate(v string) *GetBandwidthLogRequest {
  s.Startdate = &v
  return s
}

func (s *GetBandwidthLogRequest) SetEnddate(v string) *GetBandwidthLogRequest {
  s.Enddate = &v
  return s
}

func (s *GetBandwidthLogRequest) SetTimezone(v string) *GetBandwidthLogRequest {
  s.Timezone = &v
  return s
}

func (s *GetBandwidthLogRequest) SetChannel(v string) *GetBandwidthLogRequest {
  s.Channel = &v
  return s
}

func (s *GetBandwidthLogRequest) SetRegion(v string) *GetBandwidthLogRequest {
  s.Region = &v
  return s
}

func (s *GetBandwidthLogRequest) SetAccetype(v string) *GetBandwidthLogRequest {
  s.Accetype = &v
  return s
}

func (s *GetBandwidthLogRequest) SetDataformat(v string) *GetBandwidthLogRequest {
  s.Dataformat = &v
  return s
}

func (s *GetBandwidthLogRequest) SetResultType(v string) *GetBandwidthLogRequest {
  s.ResultType = &v
  return s
}

func (s *GetBandwidthLogRequest) SetFlowUnit(v string) *GetBandwidthLogRequest {
  s.FlowUnit = &v
  return s
}

func (s *GetBandwidthLogRequest) SetTimeFromZero(v string) *GetBandwidthLogRequest {
  s.TimeFromZero = &v
  return s
}

type GetBandwidthLogResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *GetBandwidthLogResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthLogResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponse) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponse) SetProvider(v *GetBandwidthLogResponseProvider) *GetBandwidthLogResponse {
  s.Provider = v
  return s
}

type GetBandwidthLogResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *GetBandwidthLogResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthLogResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProvider) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProvider) SetName(v string) *GetBandwidthLogResponseProvider {
  s.Name = &v
  return s
}

func (s *GetBandwidthLogResponseProvider) SetType(v string) *GetBandwidthLogResponseProvider {
  s.Type = &v
  return s
}

func (s *GetBandwidthLogResponseProvider) SetResultType(v string) *GetBandwidthLogResponseProvider {
  s.ResultType = &v
  return s
}

func (s *GetBandwidthLogResponseProvider) SetDate(v *GetBandwidthLogResponseProviderDate) *GetBandwidthLogResponseProvider {
  s.Date = v
  return s
}

type GetBandwidthLogResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *GetBandwidthLogResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthLogResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProviderDate) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProviderDate) SetStartdate(v string) *GetBandwidthLogResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDate) SetEnddate(v string) *GetBandwidthLogResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDate) SetChannel(v *GetBandwidthLogResponseProviderDateChannel) *GetBandwidthLogResponseProviderDate {
  s.Channel = v
  return s
}

type GetBandwidthLogResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*GetBandwidthLogResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthLogResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProviderDateChannel) SetName(v string) *GetBandwidthLogResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDateChannel) SetBandwidth(v []*GetBandwidthLogResponseProviderDateChannelBandwidth) *GetBandwidthLogResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type GetBandwidthLogResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetBandwidthLogResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProviderDateChannelBandwidth) SetTime(v string) *GetBandwidthLogResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDateChannelBandwidth) SetText(v string) *GetBandwidthLogResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type GetBandwidthLogPaths struct {
}

func (s GetBandwidthLogPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogPaths) GoString() string {
  return s.String()
}

type GetBandwidthLogParameters struct {
}

func (s GetBandwidthLogParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogParameters) GoString() string {
  return s.String()
}

type GetBandwidthLogRequestHeader struct {
}

func (s GetBandwidthLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogRequestHeader) GoString() string {
  return s.String()
}

type GetBandwidthLogResponseHeader struct {
}

func (s GetBandwidthLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseHeader) GoString() string {
  return s.String()
}




type QueryRealTimeBandwidthForMultiDomainRequest struct {
  // {'en':'Start time: 
  //     1.Start time: time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (December 2rd, 2016, 10:00 a.m., Beijing Time); 
  //     2.Not greater than the current time;
  //     3.The most recent half-year (183 days) data can be obtained
  // ', 'zh_CN':'开始时间：
  // 
  // 时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年1月1日10点0分0秒）；
  // 不能大于当前时间
  // 最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time: 
  //     1.The time format is 2016-12-02T10:00:00+08:00 
  //     2.End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  //     3.If both fields of dataFrom and dateTo are left empty, then data in the last 1 hours will be queried by default; if one field is filled and one is left empty, then exception will occur. 
  //     4.Maximum time range allowable for query: The default value is 1 hour, that is, the difference between dateFrom and dateTo cannot exceed 1 hour (you can contact technical support to adjust it, the maximum is 31 days)
  // ', 'zh_CN':'结束时间：
  // 
  // 时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常
  // 允许查询最大时间间隔：默认1小时，即dateFrom和dateTo相差不能超过1小时（可联系技术支持调整，最长31天）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity: 
  //     1. Support 1m (1 minute granularity),5m (5 minutes granularity) 
  //     2. The default value is 1m
  // ', 'zh_CN':'数据粒度：不传默认1m
  // 
  // 支持1m（1分钟）、5m（5分钟）'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Domain: 
  //     1.The default upper limit of domains that can be entered is 20 (if you want to adjust, please, contact technical support); 
  //     2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)
  // ', 'zh_CN':'域名：
  // 
  // 可传递域名数量上限默认为20（可联系技术支持调整）；
  // 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'The optional value of the grouping dimension is domain; if it is passed in, detailed data will be displayed according to this dimension;', 'zh_CN':'分组维度
  // 
  // 可选值为domain；
  // 有传入则按照该维度展示明细数据；'}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryRealTimeBandwidthForMultiDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDateFrom(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDateTo(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDataInterval(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDomain(v []*string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.Domain = v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetGroupBy(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.GroupBy = &v
  return s
}

type QueryRealTimeBandwidthForMultiDomainResponse struct {
  // {'en':'request result status code', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'request result information', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryRealTimeBandwidthForMultiDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRealTimeBandwidthForMultiDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainResponse) SetCode(v string) *QueryRealTimeBandwidthForMultiDomainResponse {
  s.Code = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponse) SetMessage(v string) *QueryRealTimeBandwidthForMultiDomainResponse {
  s.Message = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponse) SetData(v []*QueryRealTimeBandwidthForMultiDomainResponseData) *QueryRealTimeBandwidthForMultiDomainResponse {
  s.Data = v
  return s
}

type QueryRealTimeBandwidthForMultiDomainResponseData struct     {
  // {'en':' Domain name. If you do not select domain name group Dimension, this field is a semicolon-separated string of all domain names.', 'zh_CN':'域名，如果不选择域名分组维度，该字段为所有域名以分号分隔的字符串'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Peak Bandwidth,unit is Mbps,example(9811.21Mbps)', 'zh_CN':'峰值带宽，单位Mbps，示例 （9811.21Mbps)'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'Time of peak bandwidth,example(2019-02-13 18:01)', 'zh_CN':'峰值时间，示例（2019-02-13 18:01）'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'Edge total traffic,example(74099.91MB)', 'zh_CN':'边缘总流量，单位MB，示例 ( 74099.91MB)'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  BandwidthData []*QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRealTimeBandwidthForMultiDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponseData) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetDomain(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.Domain = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetPeakValue(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetPeakTime(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetTotal(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.Total = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetBandwidthData(v []*QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.BandwidthData = v
  return s
}

type QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData struct     {
  // {'en':'Time: 1. When the data query granularity is 1m, then the format is yyyy-MM-dd HH:mm; 
  //     Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00; 
  //     2. Return the time slices that contained in start time and in end time.
  // ', 'zh_CN':'格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。
  // 
  // 一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是第二天（yyyy-MM-dd） 00:00。
  // 
  // 返回开始时间和结束时间包含的时间片'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Edge bandwidth,the unit is Mbps,keep 2 decimal places', 'zh_CN':'带宽值，单位Mbps，保留2位小数。'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) SetTimestamp(v string) *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) SetValue(v string) *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData {
  s.Value = &v
  return s
}

type QueryRealTimeBandwidthForMultiDomainPaths struct {
}

func (s QueryRealTimeBandwidthForMultiDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainPaths) GoString() string {
  return s.String()
}

type QueryRealTimeBandwidthForMultiDomainParameters struct {
}

func (s QueryRealTimeBandwidthForMultiDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainParameters) GoString() string {
  return s.String()
}

type QueryRealTimeBandwidthForMultiDomainRequestHeader struct {
}

func (s QueryRealTimeBandwidthForMultiDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryRealTimeBandwidthForMultiDomainResponseHeader struct {
}

func (s QueryRealTimeBandwidthForMultiDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponseHeader) GoString() string {
  return s.String()
}




type BandwidthChannelValidRequest struct {
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not Specifies,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"domains that been queried:\n1)If there are multiple inputs,use  ';' as separator.\n2)If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. \n2)With format yyyy-mm-dd hh:MM.If 'hh:MM'","zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. \n2)With format yyyy-mm-dd hh:MM.If 'hh:MM'","zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
}

func (s BandwidthChannelValidRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelValidRequest) GoString() string {
  return s.String()
}

func (s *BandwidthChannelValidRequest) SetDate(v string) *BandwidthChannelValidRequest {
  s.Date = &v
  return s
}

func (s *BandwidthChannelValidRequest) SetChannel(v string) *BandwidthChannelValidRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthChannelValidRequest) SetEnddate(v string) *BandwidthChannelValidRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelValidRequest) SetStartdate(v string) *BandwidthChannelValidRequest {
  s.Startdate = &v
  return s
}

type BandwidthChannelValidRequestHeader struct {
}

func (s BandwidthChannelValidRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelValidRequestHeader) GoString() string {
  return s.String()
}

type BandwidthChannelValidPaths struct {
}

func (s BandwidthChannelValidPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelValidPaths) GoString() string {
  return s.String()
}

type BandwidthChannelValidParameters struct {
}

func (s BandwidthChannelValidParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelValidParameters) GoString() string {
  return s.String()
}

type BandwidthChannelValidResponse struct {
  // {"en":"startdate,with format yyyy-mm-dd hh:MM","zh_CN":"开始时间，yyyy-mm-dd hh:MM"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate,with format yyyy-mm-dd hh:MM","zh_CN":"结束时间，yyyy-mm-dd hh:MM"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"domain name.","zh_CN":"频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"time","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"text,displaying the bandwidth with unit Mbps","zh_CN":"时间点对应的带宽值，单位Mbps，带单位"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthChannelValidResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelValidResponse) GoString() string {
  return s.String()
}

func (s *BandwidthChannelValidResponse) SetStartdate(v string) *BandwidthChannelValidResponse {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelValidResponse) SetEnddate(v string) *BandwidthChannelValidResponse {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelValidResponse) SetChannel(v string) *BandwidthChannelValidResponse {
  s.Channel = &v
  return s
}

func (s *BandwidthChannelValidResponse) SetTime(v string) *BandwidthChannelValidResponse {
  s.Time = &v
  return s
}

func (s *BandwidthChannelValidResponse) SetText(v string) *BandwidthChannelValidResponse {
  s.Text = &v
  return s
}

type BandwidthChannelValidResponseHeader struct {
}

func (s BandwidthChannelValidResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelValidResponseHeader) GoString() string {
  return s.String()
}




type QueryCPSBandwidthRequest struct {
  // {"en":"cust_en_name.If there are multiple inputs,use ';' as separator","zh_CN":"客户的英文名，多个';'隔开"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope.\n2.With format yyyy-mm-dd.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；\n\n此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope.\n\n2.With format yyyy-mm-dd 3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；\n\n此参数需与startdate参数配合,若存在date参数,则该参数无效"}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
  // {"en":"domains that been queried:\n1.If there are multiple inputs,use ';' as separator.\n2.If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，\n\n返回多个频道的汇总值，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，\n\nGMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Input 'flowInfo', which will show the bandwidth peak, peak time, total flow. not selected or null&nbsp;will not be showed","zh_CN":"入参flowInfo，将展示带宽峰值、峰值时间、总流量,不选或者为空默认不展示"}
  OptionalFields *string `json:"optionalFields,omitempty" xml:"optionalFields,omitempty"`
  // {"en":"This parameter represents the type of the dedicated line. It can have multiple values. Use the semicolon as a separator if there are multiple values. hk: represents China Premium Service; jp: represents China Premium Service-Basic. eu: not applicable to CDNW. all: not applicable to CDNW.","zh_CN":"专线类型:hk;jp，支持多个，用分号隔开。hk:中港；jp:中日；eu:中欧；all:所有，默认为：hk"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty"`
}

func (s QueryCPSBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthRequest) SetCust(v string) *QueryCPSBandwidthRequest {
  s.Cust = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetStartDate(v string) *QueryCPSBandwidthRequest {
  s.StartDate = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetEndDate(v string) *QueryCPSBandwidthRequest {
  s.EndDate = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetChannel(v string) *QueryCPSBandwidthRequest {
  s.Channel = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetTimezone(v string) *QueryCPSBandwidthRequest {
  s.Timezone = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetOptionalFields(v string) *QueryCPSBandwidthRequest {
  s.OptionalFields = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetLineType(v string) *QueryCPSBandwidthRequest {
  s.LineType = &v
  return s
}

type QueryCPSBandwidthRequestHeader struct {
}

func (s QueryCPSBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryCPSBandwidthPaths struct {
}

func (s QueryCPSBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthPaths) GoString() string {
  return s.String()
}

type QueryCPSBandwidthParameters struct {
}

func (s QueryCPSBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthParameters) GoString() string {
  return s.String()
}

type QueryCPSBandwidthResponse struct {
  // {"en":"code","zh_CN":"返回编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"频道带宽数据"}
  Data *QueryCPSBandwidthResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCPSBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthResponse) SetCode(v string) *QueryCPSBandwidthResponse {
  s.Code = &v
  return s
}

func (s *QueryCPSBandwidthResponse) SetMessage(v string) *QueryCPSBandwidthResponse {
  s.Message = &v
  return s
}

func (s *QueryCPSBandwidthResponse) SetData(v *QueryCPSBandwidthResponseData) *QueryCPSBandwidthResponse {
  s.Data = v
  return s
}

type QueryCPSBandwidthResponseData struct {
  // {"en":"totalFlow","zh_CN":"总流量，单位GB"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"peakTime","zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"peakvalue","zh_CN":"带宽峰值，单位Mbps"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"detail","zh_CN":"时点带宽数据"}
  Detail *QueryCPSBandwidthResponseDataDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Struct"`
}

func (s QueryCPSBandwidthResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponseData) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthResponseData) SetTotalFlow(v string) *QueryCPSBandwidthResponseData {
  s.TotalFlow = &v
  return s
}

func (s *QueryCPSBandwidthResponseData) SetPeakTime(v string) *QueryCPSBandwidthResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryCPSBandwidthResponseData) SetPeakValue(v string) *QueryCPSBandwidthResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryCPSBandwidthResponseData) SetDetail(v *QueryCPSBandwidthResponseDataDetail) *QueryCPSBandwidthResponseData {
  s.Detail = v
  return s
}

type QueryCPSBandwidthResponseDataDetail struct {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"带宽"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QueryCPSBandwidthResponseDataDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponseDataDetail) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthResponseDataDetail) SetTime(v string) *QueryCPSBandwidthResponseDataDetail {
  s.Time = &v
  return s
}

func (s *QueryCPSBandwidthResponseDataDetail) SetText(v string) *QueryCPSBandwidthResponseDataDetail {
  s.Text = &v
  return s
}

type QueryCPSBandwidthResponseHeader struct {
}

func (s QueryCPSBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponseHeader) GoString() string {
  return s.String()
}




type QueryP2PBandwidthRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date: 1.With format yyyy-mm-dd. 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope. 2.With format yyyy-mm-dd. 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope.
  // 2.With format yyyy-mm-dd
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Whether the channel is completely matched. When it is true, the complete domain name must be filled in (invalid or duplicate channels entered by the user will be filtered at this time, and 403 will be returned if all input channels are invalid. If it is not true, the display will end with the channel entered by the user All channels. Defaults to true", "zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"A value of 1 means log bandwidth is returned, 0 means cdn bandwidth is returned, and the default value is 0", "zh_CN":"值为1，表示返回的是log带宽，0表示返回的是cdn带宽，默认值为0"}
  IsLog *string `json:"isLog,omitempty" xml:"isLog,omitempty"`
  // {"en":"Control whether to return box bandwidth, the default is 0, when box=1, return cdn(log) bandwidth+p2p bandwidth+box bandwidth; when box=0, return cdn(log) bandwidth+p2p bandwidth", "zh_CN":"控制是否返回盒子带宽,默认为0,box=1时，返回cdn(log)带宽+p2p带宽+box带宽;box=0时，返回cdn(log)带宽+p2p带宽"}
  Box *string `json:"box,omitempty" xml:"box,omitempty"`
  // {"en":"Control whether to return flow details, the default is 0 and no return, return flow details when passing 1", "zh_CN":"控制是否返回流量明细，默认为0不返回，传1时返回流量明细"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s QueryP2PBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthRequest) SetCust(v string) *QueryP2PBandwidthRequest {
  s.Cust = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetDate(v string) *QueryP2PBandwidthRequest {
  s.Date = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetStartdate(v string) *QueryP2PBandwidthRequest {
  s.Startdate = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetEnddate(v string) *QueryP2PBandwidthRequest {
  s.Enddate = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetChannel(v string) *QueryP2PBandwidthRequest {
  s.Channel = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetAccetype(v string) *QueryP2PBandwidthRequest {
  s.Accetype = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetDataformat(v string) *QueryP2PBandwidthRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetIsExactMatch(v string) *QueryP2PBandwidthRequest {
  s.IsExactMatch = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetIsLog(v string) *QueryP2PBandwidthRequest {
  s.IsLog = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetBox(v string) *QueryP2PBandwidthRequest {
  s.Box = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetFlow(v string) *QueryP2PBandwidthRequest {
  s.Flow = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetRegion(v string) *QueryP2PBandwidthRequest {
  s.Region = &v
  return s
}

type QueryP2PBandwidthResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryP2PBandwidthResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryP2PBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponse) SetProvider(v *QueryP2PBandwidthResponseProvider) *QueryP2PBandwidthResponse {
  s.Provider = v
  return s
}

type QueryP2PBandwidthResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'p2p带宽数据'}
  Date *QueryP2PBandwidthResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryP2PBandwidthResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProvider) SetName(v string) *QueryP2PBandwidthResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryP2PBandwidthResponseProvider) SetType(v string) *QueryP2PBandwidthResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryP2PBandwidthResponseProvider) SetDate(v *QueryP2PBandwidthResponseProviderDate) *QueryP2PBandwidthResponseProvider {
  s.Date = v
  return s
}

type QueryP2PBandwidthResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *QueryP2PBandwidthResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s QueryP2PBandwidthResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProviderDate) SetStartdate(v string) *QueryP2PBandwidthResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDate) SetEnddate(v string) *QueryP2PBandwidthResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDate) SetChannel(v *QueryP2PBandwidthResponseProviderDateChannel) *QueryP2PBandwidthResponseProviderDate {
  s.Channel = v
  return s
}

type QueryP2PBandwidthResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'p2p带宽数据'}
  Bandwidth []*QueryP2PBandwidthResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s QueryP2PBandwidthResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProviderDateChannel) SetName(v string) *QueryP2PBandwidthResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannel) SetBandwidth(v []*QueryP2PBandwidthResponseProviderDateChannelBandwidth) *QueryP2PBandwidthResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type QueryP2PBandwidthResponseProviderDateChannelBandwidth struct     {
  // {'en':'time', 'zh_CN':'时间点，格式 yyyy-MM-dd hh:mm:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'p2pBandWidth', 'zh_CN':'cdn带宽，单位 Mbps'}
  P2p *string `json:"p2p,omitempty" xml:"p2p,omitempty" require:"true"`
  // {'en':'cdnBandWidth', 'zh_CN':'cdn带宽，单位 Mbps'}
  Cdn *string `json:"cdn,omitempty" xml:"cdn,omitempty" require:"true"`
  // {'en':'boxBandWidth', 'zh_CN':'p2sp带宽，单位 Mbps'}
  Box *string `json:"box,omitempty" xml:"box,omitempty" require:"true"`
  // {'en':'totalBandWidth', 'zh_CN':'总带宽，单位 Mbps'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryP2PBandwidthResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetTime(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetP2p(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.P2p = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetCdn(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Cdn = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetBox(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Box = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetTotal(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Total = &v
  return s
}

type QueryP2PBandwidthPaths struct {
}

func (s QueryP2PBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthPaths) GoString() string {
  return s.String()
}

type QueryP2PBandwidthParameters struct {
}

func (s QueryP2PBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthParameters) GoString() string {
  return s.String()
}

type QueryP2PBandwidthRequestHeader struct {
}

func (s QueryP2PBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryP2PBandwidthResponseHeader struct {
}

func (s QueryP2PBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseHeader) GoString() string {
  return s.String()
}




type BandwidthTotalRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1)If there isp multiple inputs,use ';' as demimeter.
  // 2)optional values of isp: refers to the ISP-section of appendix.
  // 3) If not specified,means all the isp.", "zh_CN":"&nbsp;要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s BandwidthTotalRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalRequest) GoString() string {
  return s.String()
}

func (s *BandwidthTotalRequest) SetCust(v string) *BandwidthTotalRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthTotalRequest) SetDate(v string) *BandwidthTotalRequest {
  s.Date = &v
  return s
}

func (s *BandwidthTotalRequest) SetStartdate(v string) *BandwidthTotalRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthTotalRequest) SetEnddate(v string) *BandwidthTotalRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthTotalRequest) SetChannel(v string) *BandwidthTotalRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthTotalRequest) SetRegion(v string) *BandwidthTotalRequest {
  s.Region = &v
  return s
}

func (s *BandwidthTotalRequest) SetAccetype(v string) *BandwidthTotalRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthTotalRequest) SetDataformat(v string) *BandwidthTotalRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthTotalRequest) SetIsExactMatch(v string) *BandwidthTotalRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthTotalRequest) SetIsp(v string) *BandwidthTotalRequest {
  s.Isp = &v
  return s
}

type BandwidthTotalResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthTotalResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthTotalResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalResponse) GoString() string {
  return s.String()
}

func (s *BandwidthTotalResponse) SetProvider(v *BandwidthTotalResponseProvider) *BandwidthTotalResponse {
  s.Provider = v
  return s
}

type BandwidthTotalResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'时点数据'}
  Date *BandwidthTotalResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthTotalResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthTotalResponseProvider) SetName(v string) *BandwidthTotalResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthTotalResponseProvider) SetType(v string) *BandwidthTotalResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthTotalResponseProvider) SetResultType(v string) *BandwidthTotalResponseProvider {
  s.ResultType = &v
  return s
}

func (s *BandwidthTotalResponseProvider) SetDate(v *BandwidthTotalResponseProviderDate) *BandwidthTotalResponseProvider {
  s.Date = v
  return s
}

type BandwidthTotalResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthTotalResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthTotalResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthTotalResponseProviderDate) SetStartdate(v string) *BandwidthTotalResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthTotalResponseProviderDate) SetEnddate(v string) *BandwidthTotalResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthTotalResponseProviderDate) SetChannel(v *BandwidthTotalResponseProviderDateChannel) *BandwidthTotalResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthTotalResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'totalFlow', 'zh_CN':'总流量'}
  TotalFlow []*BandwidthTotalResponseProviderDateChannelTotalFlow `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthTotalResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthTotalResponseProviderDateChannel) SetName(v string) *BandwidthTotalResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthTotalResponseProviderDateChannel) SetTotalFlow(v []*BandwidthTotalResponseProviderDateChannelTotalFlow) *BandwidthTotalResponseProviderDateChannel {
  s.TotalFlow = v
  return s
}

type BandwidthTotalResponseProviderDateChannelTotalFlow struct     {
  // {'en':'the total flow(GB)', 'zh_CN':'总流量（单位GB）'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthTotalResponseProviderDateChannelTotalFlow) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalResponseProviderDateChannelTotalFlow) GoString() string {
  return s.String()
}

func (s *BandwidthTotalResponseProviderDateChannelTotalFlow) SetText(v string) *BandwidthTotalResponseProviderDateChannelTotalFlow {
  s.Text = &v
  return s
}

type BandwidthTotalPaths struct {
}

func (s BandwidthTotalPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalPaths) GoString() string {
  return s.String()
}

type BandwidthTotalParameters struct {
}

func (s BandwidthTotalParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalParameters) GoString() string {
  return s.String()
}

type BandwidthTotalRequestHeader struct {
}

func (s BandwidthTotalRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalRequestHeader) GoString() string {
  return s.String()
}

type BandwidthTotalResponseHeader struct {
}

func (s BandwidthTotalResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthTotalResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainOriginResponseTimeServiceRequest struct {
  // {'en':'Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2024-01-23T10:00 + 08:00 (10:00:00 Beijing time on January 23, 2024);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.', 'zh_CN':'开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2024-01-23T10:00:00+08:00（为北京时间2024年01月23日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. The end time is greater than the start time.
  // 
  // 3. If the end time is greater than the current time, the current time is taken.
  // 
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days. ', 'zh_CN':'结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment);
  // 
  // 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).', 'zh_CN':'域名：
  // 
  // 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)；
  // 
  // 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s ReportDomainOriginResponseTimeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainOriginResponseTimeServiceRequest) SetDateFrom(v string) *ReportDomainOriginResponseTimeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainOriginResponseTimeServiceRequest) SetDateTo(v string) *ReportDomainOriginResponseTimeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainOriginResponseTimeServiceRequest) SetDomain(v []*string) *ReportDomainOriginResponseTimeServiceRequest {
  s.Domain = v
  return s
}

type ReportDomainOriginResponseTimeServiceResponse struct {
  // {'en':'data', 'zh_CN':'请求结果'}
  Data []*ReportDomainOriginResponseTimeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainOriginResponseTimeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainOriginResponseTimeServiceResponse) SetData(v []*ReportDomainOriginResponseTimeServiceResponseData) *ReportDomainOriginResponseTimeServiceResponse {
  s.Data = v
  return s
}

type ReportDomainOriginResponseTimeServiceResponseData struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'details', 'zh_CN':'请求结果的详细数据'}
  DetailList []*ReportDomainOriginResponseTimeServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainOriginResponseTimeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainOriginResponseTimeServiceResponseData) SetDomain(v string) *ReportDomainOriginResponseTimeServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportDomainOriginResponseTimeServiceResponseData) SetDetailList(v []*ReportDomainOriginResponseTimeServiceResponseDataDetailList) *ReportDomainOriginResponseTimeServiceResponseData {
  s.DetailList = v
  return s
}

type ReportDomainOriginResponseTimeServiceResponseDataDetailList struct     {
  // {'en':'Time:
  // 
  // 1. When the data query granularity is 1m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  // 2. Return the time slices that contained in start time and in end time.', 'zh_CN':'时间，
  // 查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Return to origin response time. Unit is ms and 2 digits of decimals are allowed.
  // ', 'zh_CN':'回源响应时间，单位ms。保留2位小数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportDomainOriginResponseTimeServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportDomainOriginResponseTimeServiceResponseDataDetailList) SetTimestamp(v string) *ReportDomainOriginResponseTimeServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportDomainOriginResponseTimeServiceResponseDataDetailList) SetValue(v string) *ReportDomainOriginResponseTimeServiceResponseDataDetailList {
  s.Value = &v
  return s
}

type ReportDomainOriginResponseTimeServicePaths struct {
}

func (s ReportDomainOriginResponseTimeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServicePaths) GoString() string {
  return s.String()
}

type ReportDomainOriginResponseTimeServiceParameters struct {
}

func (s ReportDomainOriginResponseTimeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainOriginResponseTimeServiceRequestHeader struct {
}

func (s ReportDomainOriginResponseTimeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainOriginResponseTimeServiceResponseHeader struct {
}

func (s ReportDomainOriginResponseTimeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainOriginResponseTimeServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportCountryServerBandwidthServiceRequest struct {
  // {"en":"Start time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days.  ", "zh_CN":"结束时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error,you can contact technical support for adjustment);
  // 
  // 2.Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Country area:
  // 
  // 1. countryCode is not uploaded: Query all country areas by default;
  // 
  // 2. countryCode is uploaded: Multiple can be uploaded, such as cn, in.  Please refer to the appendix description section of the overview page.", "zh_CN":"国家区域(含中国台湾、中国澳门、中国香港、中国大陆):
  // 
  // 1.未传递countryCode时:查询全部国家区域;
  // 
  // 2.有传递countryCode时:可传多个,如cn,in。可传递的值详见概览页附录说明章节"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  // 5m: 5 minute granularity; Default value is 5m.
  // 
  // 1h: 1 hour granularity.", "zh_CN":"数据粒度:
  // 
  // 5m:5分钟粒度。不传默认5分钟粒度
  // 
  // 1h:1小时粒度;"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportCountryServerBandwidthServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDateFrom(v string) *ReportCountryServerBandwidthServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDateTo(v string) *ReportCountryServerBandwidthServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDomain(v []*string) *ReportCountryServerBandwidthServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetCountryCode(v []*string) *ReportCountryServerBandwidthServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDataInterval(v string) *ReportCountryServerBandwidthServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportCountryServerBandwidthServiceResponse struct {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportCountryServerBandwidthServiceResponseCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportCountryServerBandwidthServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceResponse) SetDomain(v string) *ReportCountryServerBandwidthServiceResponse {
  s.Domain = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponse) SetCountryData(v []*ReportCountryServerBandwidthServiceResponseCountryData) *ReportCountryServerBandwidthServiceResponse {
  s.CountryData = v
  return s
}

type ReportCountryServerBandwidthServiceResponseCountryData struct     {
  // {"en":"Country area", "zh_CN":"国家地区"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  DetailList []*ReportCountryServerBandwidthServiceResponseCountryDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportCountryServerBandwidthServiceResponseCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponseCountryData) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceResponseCountryData) SetCountryCode(v string) *ReportCountryServerBandwidthServiceResponseCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponseCountryData) SetDetailList(v []*ReportCountryServerBandwidthServiceResponseCountryDataDetailList) *ReportCountryServerBandwidthServiceResponseCountryData {
  s.DetailList = v
  return s
}

type ReportCountryServerBandwidthServiceResponseCountryDataDetailList struct     {
  // {"en":"Time:
  // 
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; ach time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  // 2. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00;
  // 
  // 3. Return the time slices that contained in start time and in end time.", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth value, unit Mbps", "zh_CN":"时间点对应的边缘带宽数据,单位Mbps"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Flow, unit MB", "zh_CN":"时间点对应的边缘流量,单位MB"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s ReportCountryServerBandwidthServiceResponseCountryDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponseCountryDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceResponseCountryDataDetailList) SetTimestamp(v string) *ReportCountryServerBandwidthServiceResponseCountryDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponseCountryDataDetailList) SetValue(v string) *ReportCountryServerBandwidthServiceResponseCountryDataDetailList {
  s.Value = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponseCountryDataDetailList) SetFlow(v string) *ReportCountryServerBandwidthServiceResponseCountryDataDetailList {
  s.Flow = &v
  return s
}

type ReportCountryServerBandwidthServicePaths struct {
}

func (s ReportCountryServerBandwidthServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServicePaths) GoString() string {
  return s.String()
}

type ReportCountryServerBandwidthServiceParameters struct {
}

func (s ReportCountryServerBandwidthServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceParameters) GoString() string {
  return s.String()
}

type ReportCountryServerBandwidthServiceRequestHeader struct {
}

func (s ReportCountryServerBandwidthServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportCountryServerBandwidthServiceResponseHeader struct {
}

func (s ReportCountryServerBandwidthServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportBandwidthRequestByIPIspProvinceRequest struct {
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
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days.  ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains: 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment); 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).", "zh_CN":"域名： 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)； 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity: 1m:1 mintues5m: 5 minute granularity; Default value is 5m. ", "zh_CN":"数据粒度：1m：1分钟粒度 5m：5分钟粒度。不传默认5分钟粒度 "}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province: 1. Province is empty: Query all provinces and aggregate the returned data according to all provinces; 2. Province : please send province code. ", "zh_CN":"省份： 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。 2.有传递province时：传递省份code，可传多个"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP: 1. ISP  is empty: Query all ISPs and aggregate the returned data according to all ISPs; 2. ISP: please send ISP code.", "zh_CN":"运营商： 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 2.有传递isp时：传递运营商code，可传多个"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension 1.Options are domain, province, isp, and more than one value can be entered; 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度 可选值为domain、province、isp，可传入多个值； 有传入则按照该维度展示明细数据； 没传默认全部聚合。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportBandwidthRequestByIPIspProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceRequest) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetDateFrom(v string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetDateTo(v string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetDomain(v []*string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.Domain = v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetDataInterval(v string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetProvince(v []*string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.Province = v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetIsp(v []*string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.Isp = v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceRequest) SetGroupBy(v []*string) *ReportBandwidthRequestByIPIspProvinceRequest {
  s.GroupBy = v
  return s
}

type ReportBandwidthRequestByIPIspProvinceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"result", "zh_CN":"结果"}
  Data []*ReportBandwidthRequestByIPIspProvinceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthRequestByIPIspProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceResponse) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRequestByIPIspProvinceResponse) SetCode(v string) *ReportBandwidthRequestByIPIspProvinceResponse {
  s.Code = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponse) SetMessage(v string) *ReportBandwidthRequestByIPIspProvinceResponse {
  s.Message = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponse) SetData(v []*ReportBandwidthRequestByIPIspProvinceResponseData) *ReportBandwidthRequestByIPIspProvinceResponse {
  s.Data = v
  return s
}

type ReportBandwidthRequestByIPIspProvinceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ISP code", "zh_CN":"运营商code"}
  IspData []*ReportBandwidthRequestByIPIspProvinceResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthRequestByIPIspProvinceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceResponseData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseData) SetDomain(v string) *ReportBandwidthRequestByIPIspProvinceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseData) SetIspData(v []*ReportBandwidthRequestByIPIspProvinceResponseDataIspData) *ReportBandwidthRequestByIPIspProvinceResponseData {
  s.IspData = v
  return s
}

type ReportBandwidthRequestByIPIspProvinceResponseDataIspData struct     {
  // {"en":"Province code", "zh_CN":"省份code"}
  ProvinceData []*ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthRequestByIPIspProvinceResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceResponseDataIspData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspData) SetProvinceData(v []*ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData) *ReportBandwidthRequestByIPIspProvinceResponseDataIspData {
  s.ProvinceData = v
  return s
}

type ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData struct     {
  // {"en":"", "zh_CN":""}
  DetailList []*ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData) SetDetailList(v []*ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceData {
  s.DetailList = v
  return s
}

type ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList struct     {
  // {"en":"Time: 1.the format is yyyy-MM-dd HH:mm;   ach time slice value represents the value within the previous time granularity range. When the data query granularity is 5m,The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00; 2. When the data query granularity is 1m, The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间， 格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。查询的数据粒度为5m时，一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00； 查询的数据粒度为1m时，一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of IPv4 requests.", "zh_CN":"IPv4请求数"}
  V4Request *string `json:"v4Request,omitempty" xml:"v4Request,omitempty" require:"true"`
  // {"en":"Bandwidth of IPv4", "zh_CN":"IPv4带宽 单位Mbps。保留2位小数"}
  V4Bandwidth *string `json:"v4Bandwidth,omitempty" xml:"v4Bandwidth,omitempty" require:"true"`
  // {"en":"Number of IPv6 requests.", "zh_CN":"IPv6请求数"}
  V6Request *string `json:"v6Request,omitempty" xml:"v6Request,omitempty" require:"true"`
  // {"en":"Bandwidth of IPv6", "zh_CN":"IPv6带宽 单位Mbps。保留2位小数"}
  V6Bandwidth *string `json:"v6Bandwidth,omitempty" xml:"v6Bandwidth,omitempty" require:"true"`
}

func (s ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) SetTimestamp(v string) *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) SetV4Request(v string) *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList {
  s.V4Request = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) SetV4Bandwidth(v string) *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList {
  s.V4Bandwidth = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) SetV6Request(v string) *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList {
  s.V6Request = &v
  return s
}

func (s *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList) SetV6Bandwidth(v string) *ReportBandwidthRequestByIPIspProvinceResponseDataIspDataProvinceDataDetailList {
  s.V6Bandwidth = &v
  return s
}

type ReportBandwidthRequestByIPIspProvincePaths struct {
}

func (s ReportBandwidthRequestByIPIspProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvincePaths) GoString() string {
  return s.String()
}

type ReportBandwidthRequestByIPIspProvinceParameters struct {
}

func (s ReportBandwidthRequestByIPIspProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceParameters) GoString() string {
  return s.String()
}

type ReportBandwidthRequestByIPIspProvinceRequestHeader struct {
}

func (s ReportBandwidthRequestByIPIspProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceRequestHeader) GoString() string {
  return s.String()
}

type ReportBandwidthRequestByIPIspProvinceResponseHeader struct {
}

func (s ReportBandwidthRequestByIPIspProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRequestByIPIspProvinceResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthMinutelyRequest struct {
  // {"en":"Start Time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, December 2, 2016 10:00-02T10:00:00+08:00 (Beijing time on December 2, 2016 10:00 am 0 seconds);
  // 2. No more than the current time
  // 3. Only available for the last 30 minutes.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.只能查询最近30分钟。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 1. Time Format 2016-12-02T10:00:00+08:00
  // 2. the end time should be greater than the start time. if the end time is greater than the current time, take the current time.
  // 3. dateFrom, dateTo both not upload, default query past 30 minutes; If only one is not sent, throw exception
  // 4. Allow maximum query interval: 30 minutes, i.e. the difference between dateFrom and dateTo should not exceed 30 minutes.", "zh_CN":"结束时间：
  // 
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的30分钟；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：30分钟，即dateFrom和dateTo相差不能超过30分钟。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Multiple domain names:
  // 1. The maximum number of incoming domain names is 50 (contact technical support for adjustment).
  // 2. filter out invalid domain name; A domain name that is not in the historical domain name table.", "zh_CN":"多域名：
  // 1.单次传入域名数量最多50个（可联系技术支持调整）。
  // 2.过滤掉无效域名；即不在历史域名表中的域名。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration zone:
  // 1. do not pass the default query to all areas.
  // 2. currently only the leaflet area is supported externally.
  // 3. optional values: cn, apac, am, emea.", "zh_CN":"加速区域：
  // 1.不传默认查询全部区域。
  // 2.目前对外只支持传单区域。
  // 3.可选值：cn、apac、am、emea；"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"grouped dimension:
  // 1. the optional value is domain.
  // 2. If incoming data is shown in accordance with the dimension;", "zh_CN":"分组维度:
  // 1.可选值为domain；
  // 2.有传入则按照该维度展示明细数据；"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryBandwidthMinutelyRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyRequest) GoString() string {
  return s.String()
}

func (s *QueryBandwidthMinutelyRequest) SetDateFrom(v string) *QueryBandwidthMinutelyRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBandwidthMinutelyRequest) SetDateTo(v string) *QueryBandwidthMinutelyRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBandwidthMinutelyRequest) SetDomain(v []*string) *QueryBandwidthMinutelyRequest {
  s.Domain = v
  return s
}

func (s *QueryBandwidthMinutelyRequest) SetRegion(v string) *QueryBandwidthMinutelyRequest {
  s.Region = &v
  return s
}

func (s *QueryBandwidthMinutelyRequest) SetGroupBy(v string) *QueryBandwidthMinutelyRequest {
  s.GroupBy = &v
  return s
}

type QueryBandwidthMinutelyResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryBandwidthMinutelyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthMinutelyResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthMinutelyResponse) SetCode(v string) *QueryBandwidthMinutelyResponse {
  s.Code = &v
  return s
}

func (s *QueryBandwidthMinutelyResponse) SetMessage(v string) *QueryBandwidthMinutelyResponse {
  s.Message = &v
  return s
}

func (s *QueryBandwidthMinutelyResponse) SetData(v []*QueryBandwidthMinutelyResponseData) *QueryBandwidthMinutelyResponse {
  s.Data = v
  return s
}

type QueryBandwidthMinutelyResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak Time", "zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak Bandwidth", "zh_CN":"带宽峰值"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"The granularity of data is 1 minute, and the format is yyyy-MM-dd HH:MM. Each time slice data value represents the data value in the previous time granularity range. Returns the time slice included with the start and end times.", "zh_CN":"数据粒度为1分钟，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。
  //                 一天开始的时间片是yyyy-MM-dd 00:01:00，最后一个时间片是（yyyy-MM-dd） 24:00:00。
  //                 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth value in Mbps, with 2 decimal places reserved.", "zh_CN":"带宽值，单位Mbps，保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryBandwidthMinutelyResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyResponseData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthMinutelyResponseData) SetDomain(v string) *QueryBandwidthMinutelyResponseData {
  s.Domain = &v
  return s
}

func (s *QueryBandwidthMinutelyResponseData) SetPeakTime(v string) *QueryBandwidthMinutelyResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryBandwidthMinutelyResponseData) SetPeakValue(v string) *QueryBandwidthMinutelyResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryBandwidthMinutelyResponseData) SetTimestamp(v string) *QueryBandwidthMinutelyResponseData {
  s.Timestamp = &v
  return s
}

func (s *QueryBandwidthMinutelyResponseData) SetValue(v string) *QueryBandwidthMinutelyResponseData {
  s.Value = &v
  return s
}

type QueryBandwidthMinutelyPaths struct {
}

func (s QueryBandwidthMinutelyPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyPaths) GoString() string {
  return s.String()
}

type QueryBandwidthMinutelyParameters struct {
}

func (s QueryBandwidthMinutelyParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyParameters) GoString() string {
  return s.String()
}

type QueryBandwidthMinutelyRequestHeader struct {
}

func (s QueryBandwidthMinutelyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthMinutelyResponseHeader struct {
}

func (s QueryBandwidthMinutelyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthMinutelyResponseHeader) GoString() string {
  return s.String()
}




type ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 90 days, that is, the difference between dateFrom and dateTo can not exceed 90 days.", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：31天，即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The default upper limit of domains that can be entered is 20 (if you want to adjust, please, contact technical support);
  // 2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 3.域名超过上限，报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"region:
  // 1. All areas are queried by default if the system does not transmit data.
  // 2. Multiple regions can be passed. The value can be cn, hk, ov, apac, am, emea, or fg.", "zh_CN":"区域：
  // 1.不传默认查询全部区域。
  // 2.支持传递多个区域，可选值为：cn、hk、ov、apac、am、emea、fg。"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
  // {"en":"Grouping dimension:
  // 1. By default, it is not transmitted and displayed in aggregation mode;
  // 2. The optional value supports only domain, and displays detailed data by domain.", "zh_CN":"分组维度
  // 1.默认不传，聚合展示；
  // 2.可选值仅支持domain，传入时按域名展示明细数据。"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) SetDateFrom(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) SetDateTo(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) SetDomain(v []*string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) SetRegion(v []*string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest {
  s.Region = v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest) SetGroupBy(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceRequest {
  s.GroupBy = &v
  return s
}

type ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse) SetCode(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse) SetMessage(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse) SetData(v []*ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponse {
  s.Data = v
  return s
}

type ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData struct     {
  // {"en":"Domain name. If you do not select domain name group Dimension, this field is a semicolon-separated string of all domain names.", "zh_CN":"域名，如果不选择域名分组维度，该字段为所有域名以分号分隔的字符串。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak Bandwidth,unit is Mbps,example(9811.21Mbps)", "zh_CN":"峰值带宽 Mbps，示例 （931556.21 Mbps）"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Time of peak bandwidth,example(2019-02-13 18:01)", "zh_CN":"峰值时间，示例（2019-02-13 18:01）;"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Edge total traffic,example(74099.91MB)", "zh_CN":"边缘总流量，示例 ( 74099.92 MB )"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  BandwidthData []*ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) SetDomain(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) SetPeakValue(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData {
  s.PeakValue = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) SetPeakTime(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData {
  s.PeakTime = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) SetTotal(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData {
  s.Total = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData) SetBandwidthData(v []*ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseData {
  s.BandwidthData = v
  return s
}

type ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData struct     {
  // {"en":"The data granularity is 1 minute,the format is yyyy-MM-dd HH:mm", "zh_CN":"数据粒度为5分钟，格式为yyyy-MM-dd HH:mm；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Edge bandwidth,the unit is Mbps,keep 2 decimal places", "zh_CN":"边缘带宽值，单位Mbps，保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData) GoString() string {
  return s.String()
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData) SetTimestamp(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData) SetValue(v string) *ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseDataBandwidthData {
  s.Value = &v
  return s
}

type ReportLogBandwidthMultiDomainEcdnEdgeServicePaths struct {
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServicePaths) GoString() string {
  return s.String()
}

type ReportLogBandwidthMultiDomainEcdnEdgeServiceParameters struct {
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceParameters) GoString() string {
  return s.String()
}

type ReportLogBandwidthMultiDomainEcdnEdgeServiceRequestHeader struct {
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseHeader struct {
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportLogBandwidthMultiDomainEcdnEdgeServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportP2SPOriginBandwidthWildcardDomainServiceRequest struct {
  // {'en':'Start time:
  // 1. The format is yyyy-MM-ddTHH:mm:SS+08:00, for example, 2016-12-02T10:00 + 08:00 (10:00:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.', 'zh_CN':'开始时间:
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒);
  //         2.不能大于当前时间;
  //         3.最多可获取最近半年(183天)的数据'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time;
  // 3. If the end time is greater than the current time, the current time is taken;
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours, if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 31 days, that is the difference between dateFrom and dateTo can not exceed 31 days. ', 'zh_CN':'结束时间:
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00;
  //         2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间;
  //         3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  //         4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天.'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Wildcard Domain:
  // 1. The upper limit of the number of Wildcard Domain names that can be passed is 20 by default (contact technical support for adjustment);
  // 2. Automatically filter out invalid domain names (if a non-common domain name is passed in, it will be filtered out);
  // 3. If it is not filled, all pan domain names under the account will be queried by default.', 'zh_CN':'泛域名:
  //         1.可传递泛域名数量上限默认为20个(可联系技术支持调整);
  //         2.自动过滤掉无效域名(如传递非泛域名,会被过滤掉);
  //         3.如未填,则默认查询此账号下所有泛域名.'}
  WildcardDomain []*string `json:"wildcardDomain,omitempty" xml:"wildcardDomain,omitempty" type:"Repeated"`
  // {'en':'Grouped dimension:
  // 1.When the "groupBy" parameter has a value, the value can only be "domain". It will group and return the detailed back-to-origin bandwidth based on wildcard domains and exact domains.
  // 2.When groupBy does not have a value, return the total origin bandwidth details for the wildcard domain. ', 'zh_CN':'	
  // 分组维度:
  // 1.有传值时,只能选domain,按照泛域名及精确域名分组返回回源带宽明细;
  // 2.不传时:返回查询泛域名下的总回源带宽明细.'}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceRequest) SetDateFrom(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceRequest) SetDateTo(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceRequest) SetWildcardDomain(v []*string) *ReportP2SPOriginBandwidthWildcardDomainServiceRequest {
  s.WildcardDomain = v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceRequest) SetGroupBy(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceRequest {
  s.GroupBy = &v
  return s
}

type ReportP2SPOriginBandwidthWildcardDomainServiceResponse struct {
  // {'en':'Request result status code', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Request result information', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportP2SPOriginBandwidthWildcardDomainServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponse) SetCode(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponse) SetMessage(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponse) SetData(v []*ReportP2SPOriginBandwidthWildcardDomainServiceResponseData) *ReportP2SPOriginBandwidthWildcardDomainServiceResponse {
  s.Data = v
  return s
}

type ReportP2SPOriginBandwidthWildcardDomainServiceResponseData struct     {
  // {'en':'Wildcard domain', 'zh_CN':'泛域名'}
  WildcardDomain *string `json:"wildcardDomain,omitempty" xml:"wildcardDomain,omitempty" require:"true"`
  DomainList []*ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponseData) SetWildcardDomain(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceResponseData {
  s.WildcardDomain = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponseData) SetDomainList(v []*ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList) *ReportP2SPOriginBandwidthWildcardDomainServiceResponseData {
  s.DomainList = v
  return s
}

type ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList struct     {
  // {'en':'Domain', 'zh_CN':'明细域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  OriginBandwidthList []*ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList `json:"originBandwidthList,omitempty" xml:"originBandwidthList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList) GoString() string {
  return s.String()
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList) SetDomain(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList {
  s.Domain = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList) SetOriginBandwidthList(v []*ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList) *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainList {
  s.OriginBandwidthList = v
  return s
}

type ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList struct     {
  // {'en':'Time:
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2. Return the time slices that contained in start time and in end time.', 'zh_CN':'时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值.一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 返回开始时间和结束时间包含的时间片.'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Origin bandwidth value. Unit is Mbps and 2 digits of decimals are allowed', 'zh_CN':'回源带宽数据,单位Mbps,保留2位小数.'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList) GoString() string {
  return s.String()
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList) SetTimestamp(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList {
  s.Timestamp = &v
  return s
}

func (s *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList) SetValue(v string) *ReportP2SPOriginBandwidthWildcardDomainServiceResponseDataDomainListOriginBandwidthList {
  s.Value = &v
  return s
}

type ReportP2SPOriginBandwidthWildcardDomainServicePaths struct {
}

func (s ReportP2SPOriginBandwidthWildcardDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServicePaths) GoString() string {
  return s.String()
}

type ReportP2SPOriginBandwidthWildcardDomainServiceParameters struct {
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceParameters) GoString() string {
  return s.String()
}

type ReportP2SPOriginBandwidthWildcardDomainServiceRequestHeader struct {
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportP2SPOriginBandwidthWildcardDomainServiceResponseHeader struct {
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportP2SPOriginBandwidthWildcardDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportIPV6BandwidthServiceRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据."}
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
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days. ", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天."}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error,you can contact technical support for adjustment) ;
  // 
  // 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).", "zh_CN":"域名:
  // 
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)."}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  // 5m: 5 minute granularity; Default value is 5m.
  // 
  // 1h: 1 hour granularity.", "zh_CN":"数据粒度:
  // 
  // 5m:5分钟粒度.不传默认5分钟粒度
  // 
  // 1h:1小时粒度;"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportIPV6BandwidthServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportIPV6BandwidthServiceRequest) SetDateFrom(v string) *ReportIPV6BandwidthServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportIPV6BandwidthServiceRequest) SetDateTo(v string) *ReportIPV6BandwidthServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportIPV6BandwidthServiceRequest) SetDomain(v []*string) *ReportIPV6BandwidthServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportIPV6BandwidthServiceRequest) SetDataInterval(v string) *ReportIPV6BandwidthServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportIPV6BandwidthServiceResponse struct {
  Data []*ReportIPV6BandwidthServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportIPV6BandwidthServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportIPV6BandwidthServiceResponse) SetData(v []*ReportIPV6BandwidthServiceResponseData) *ReportIPV6BandwidthServiceResponse {
  s.Data = v
  return s
}

type ReportIPV6BandwidthServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*ReportIPV6BandwidthServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportIPV6BandwidthServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportIPV6BandwidthServiceResponseData) SetDomain(v string) *ReportIPV6BandwidthServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportIPV6BandwidthServiceResponseData) SetDetailList(v []*ReportIPV6BandwidthServiceResponseDataDetailList) *ReportIPV6BandwidthServiceResponseData {
  s.DetailList = v
  return s
}

type ReportIPV6BandwidthServiceResponseDataDetailList struct     {
  // {"en":"Time:
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; ach time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值.一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值.一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 返回开始时间和结束时间包含的时间片."}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"IPV6 bandwidth data, unit Mbps. Keep 2 decimal places", "zh_CN":"IPV6带宽数据,单位Mbps.保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"IPV6 traffic, unit MB. Keep 2 decimal places", "zh_CN":"IPV6流量,单位MB.保留2位小数"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s ReportIPV6BandwidthServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportIPV6BandwidthServiceResponseDataDetailList) SetTimestamp(v string) *ReportIPV6BandwidthServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportIPV6BandwidthServiceResponseDataDetailList) SetValue(v string) *ReportIPV6BandwidthServiceResponseDataDetailList {
  s.Value = &v
  return s
}

func (s *ReportIPV6BandwidthServiceResponseDataDetailList) SetFlow(v string) *ReportIPV6BandwidthServiceResponseDataDetailList {
  s.Flow = &v
  return s
}

type ReportIPV6BandwidthServicePaths struct {
}

func (s ReportIPV6BandwidthServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServicePaths) GoString() string {
  return s.String()
}

type ReportIPV6BandwidthServiceParameters struct {
}

func (s ReportIPV6BandwidthServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceParameters) GoString() string {
  return s.String()
}

type ReportIPV6BandwidthServiceRequestHeader struct {
}

func (s ReportIPV6BandwidthServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportIPV6BandwidthServiceResponseHeader struct {
}

func (s ReportIPV6BandwidthServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportIPV6BandwidthServiceResponseHeader) GoString() string {
  return s.String()
}




type ChannelValueSumRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
}

func (s ChannelValueSumRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumRequest) GoString() string {
  return s.String()
}

func (s *ChannelValueSumRequest) SetCust(v string) *ChannelValueSumRequest {
  s.Cust = &v
  return s
}

func (s *ChannelValueSumRequest) SetDate(v string) *ChannelValueSumRequest {
  s.Date = &v
  return s
}

func (s *ChannelValueSumRequest) SetStartdate(v string) *ChannelValueSumRequest {
  s.Startdate = &v
  return s
}

func (s *ChannelValueSumRequest) SetEnddate(v string) *ChannelValueSumRequest {
  s.Enddate = &v
  return s
}

func (s *ChannelValueSumRequest) SetChannel(v string) *ChannelValueSumRequest {
  s.Channel = &v
  return s
}

func (s *ChannelValueSumRequest) SetRegion(v string) *ChannelValueSumRequest {
  s.Region = &v
  return s
}

func (s *ChannelValueSumRequest) SetIsExactMatch(v string) *ChannelValueSumRequest {
  s.IsExactMatch = &v
  return s
}

func (s *ChannelValueSumRequest) SetAccetype(v string) *ChannelValueSumRequest {
  s.Accetype = &v
  return s
}

func (s *ChannelValueSumRequest) SetDataformat(v string) *ChannelValueSumRequest {
  s.Dataformat = &v
  return s
}

type ChannelValueSumResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *ChannelValueSumResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueSumResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumResponse) GoString() string {
  return s.String()
}

func (s *ChannelValueSumResponse) SetProvider(v *ChannelValueSumResponseProvider) *ChannelValueSumResponse {
  s.Provider = v
  return s
}

type ChannelValueSumResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'多域名总计费带宽数据'}
  Date *ChannelValueSumResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueSumResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumResponseProvider) GoString() string {
  return s.String()
}

func (s *ChannelValueSumResponseProvider) SetName(v string) *ChannelValueSumResponseProvider {
  s.Name = &v
  return s
}

func (s *ChannelValueSumResponseProvider) SetType(v string) *ChannelValueSumResponseProvider {
  s.Type = &v
  return s
}

func (s *ChannelValueSumResponseProvider) SetResultType(v string) *ChannelValueSumResponseProvider {
  s.ResultType = &v
  return s
}

func (s *ChannelValueSumResponseProvider) SetDate(v *ChannelValueSumResponseProviderDate) *ChannelValueSumResponseProvider {
  s.Date = v
  return s
}

type ChannelValueSumResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'information', 'zh_CN':'频道计费数据'}
  Information []*ChannelValueSumResponseProviderDateInformation `json:"information,omitempty" xml:"information,omitempty" require:"true" type:"Repeated"`
}

func (s ChannelValueSumResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumResponseProviderDate) GoString() string {
  return s.String()
}

func (s *ChannelValueSumResponseProviderDate) SetStartdate(v string) *ChannelValueSumResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *ChannelValueSumResponseProviderDate) SetEnddate(v string) *ChannelValueSumResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *ChannelValueSumResponseProviderDate) SetInformation(v []*ChannelValueSumResponseProviderDateInformation) *ChannelValueSumResponseProviderDate {
  s.Information = v
  return s
}

type ChannelValueSumResponseProviderDateInformation struct     {
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*ChannelValueSumResponseProviderDateInformationBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s ChannelValueSumResponseProviderDateInformation) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumResponseProviderDateInformation) GoString() string {
  return s.String()
}

func (s *ChannelValueSumResponseProviderDateInformation) SetChannel(v string) *ChannelValueSumResponseProviderDateInformation {
  s.Channel = &v
  return s
}

func (s *ChannelValueSumResponseProviderDateInformation) SetBandwidth(v []*ChannelValueSumResponseProviderDateInformationBandwidth) *ChannelValueSumResponseProviderDateInformation {
  s.Bandwidth = v
  return s
}

type ChannelValueSumResponseProviderDateInformationBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'charge method', 'zh_CN':'计费方式'}
  ChargeMethod *string `json:"charge-method,omitempty" xml:"charge-method,omitempty" require:"true"`
  // {'en':'charge value', 'zh_CN':'计费值'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'charge unit', 'zh_CN':'计费单位'}
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty" require:"true"`
}

func (s ChannelValueSumResponseProviderDateInformationBandwidth) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumResponseProviderDateInformationBandwidth) GoString() string {
  return s.String()
}

func (s *ChannelValueSumResponseProviderDateInformationBandwidth) SetTime(v string) *ChannelValueSumResponseProviderDateInformationBandwidth {
  s.Time = &v
  return s
}

func (s *ChannelValueSumResponseProviderDateInformationBandwidth) SetChargeMethod(v string) *ChannelValueSumResponseProviderDateInformationBandwidth {
  s.ChargeMethod = &v
  return s
}

func (s *ChannelValueSumResponseProviderDateInformationBandwidth) SetValue(v string) *ChannelValueSumResponseProviderDateInformationBandwidth {
  s.Value = &v
  return s
}

func (s *ChannelValueSumResponseProviderDateInformationBandwidth) SetUnit(v string) *ChannelValueSumResponseProviderDateInformationBandwidth {
  s.Unit = &v
  return s
}

type ChannelValueSumPaths struct {
}

func (s ChannelValueSumPaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumPaths) GoString() string {
  return s.String()
}

type ChannelValueSumParameters struct {
}

func (s ChannelValueSumParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumParameters) GoString() string {
  return s.String()
}

type ChannelValueSumRequestHeader struct {
}

func (s ChannelValueSumRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumRequestHeader) GoString() string {
  return s.String()
}

type ChannelValueSumResponseHeader struct {
}

func (s ChannelValueSumResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueSumResponseHeader) GoString() string {
  return s.String()
}




type ReportLowDelayCountryBandwidthServiceRequest struct {
  // {"en":"Start time:
  // 	1.The format is yyyy-MM-ddTHH:mm:ss+08:00; 
  // 	2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo; 
  // 	3.Period between dataFrom and dateTo cannot be longer than 7 days; 
  // 	4.dateFrom and dateTo can be either both are specified or neither is specifies; 
  // 	5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天;(可联系技术支持调整)
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00.
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 250 (can be adjusted by contacting technical support).
  // 	2.Domain is not uploaded: Query all domain names of the account", "zh_CN":"域名:可传递域名数量上限默认为250个(可联系技术支持调整),未传递该入参时查询账号下所有域名,但当账号下域名数量超过限制时不可查询(报错)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration zone:
  // 1. If not passed or empty, it defaults to all regions.
  // 2. Multiple areas are separated by English semicolons;
  // 3. The optional values are: cn, apac, etc., see the overview page [Appendix 1] for details", "zh_CN":"加速区域:
  // 1.不选或者为空时默认为全部区域。
  // 2.多个区域使用英文分号分隔;
  // 3.可选值为:cn、apac…等,详见概览页【附表1】说明"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {"en":"Country code:
  // 1.If no value is passed, all countries and regions will be queried by default;
  // 2.For the values that can be passed, see the overview page [Appendix 1] for details.", "zh_CN":"国家地区代号:
  // 1.不传默认查询全部国家地区;
  // 2.可传递的值详见概览页【附表1】说明。"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
}

func (s ReportLowDelayCountryBandwidthServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportLowDelayCountryBandwidthServiceRequest) SetDateFrom(v string) *ReportLowDelayCountryBandwidthServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceRequest) SetDateTo(v string) *ReportLowDelayCountryBandwidthServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceRequest) SetDomain(v []*string) *ReportLowDelayCountryBandwidthServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceRequest) SetAreaCode(v []*string) *ReportLowDelayCountryBandwidthServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceRequest) SetCountryCode(v []*string) *ReportLowDelayCountryBandwidthServiceRequest {
  s.CountryCode = v
  return s
}

type ReportLowDelayCountryBandwidthServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportLowDelayCountryBandwidthServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLowDelayCountryBandwidthServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportLowDelayCountryBandwidthServiceResponse) SetCode(v string) *ReportLowDelayCountryBandwidthServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceResponse) SetMessage(v string) *ReportLowDelayCountryBandwidthServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceResponse) SetData(v []*ReportLowDelayCountryBandwidthServiceResponseData) *ReportLowDelayCountryBandwidthServiceResponse {
  s.Data = v
  return s
}

type ReportLowDelayCountryBandwidthServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportLowDelayCountryBandwidthServiceResponseDataCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLowDelayCountryBandwidthServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportLowDelayCountryBandwidthServiceResponseData) SetDomain(v string) *ReportLowDelayCountryBandwidthServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceResponseData) SetCountryData(v []*ReportLowDelayCountryBandwidthServiceResponseDataCountryData) *ReportLowDelayCountryBandwidthServiceResponseData {
  s.CountryData = v
  return s
}

type ReportLowDelayCountryBandwidthServiceResponseDataCountryData struct     {
  // {"en":"CountryCode", "zh_CN":"国家地区"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  BandwidthList []*ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList `json:"bandwidthList,omitempty" xml:"bandwidthList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLowDelayCountryBandwidthServiceResponseDataCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceResponseDataCountryData) GoString() string {
  return s.String()
}

func (s *ReportLowDelayCountryBandwidthServiceResponseDataCountryData) SetCountryCode(v string) *ReportLowDelayCountryBandwidthServiceResponseDataCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceResponseDataCountryData) SetBandwidthList(v []*ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList) *ReportLowDelayCountryBandwidthServiceResponseDataCountryData {
  s.BandwidthList = v
  return s
}

type ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList struct     {
  // {"en":"time:
  // 																	1.When the query data granularity is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value within the previous time granularity range. 
  // 																	2.The time slice at the beginning of a day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00; return the time slice included in the start time and end time.", "zh_CN":"时间:
  // 																	1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。
  // 																	2.一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"The bandwidth value corresponding to the time point, in Mbps", "zh_CN":"时间点对应的带宽值,单位Mbps"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList) GoString() string {
  return s.String()
}

func (s *ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList) SetTimestamp(v string) *ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList {
  s.Timestamp = &v
  return s
}

func (s *ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList) SetValue(v string) *ReportLowDelayCountryBandwidthServiceResponseDataCountryDataBandwidthList {
  s.Value = &v
  return s
}

type ReportLowDelayCountryBandwidthServicePaths struct {
}

func (s ReportLowDelayCountryBandwidthServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServicePaths) GoString() string {
  return s.String()
}

type ReportLowDelayCountryBandwidthServiceParameters struct {
}

func (s ReportLowDelayCountryBandwidthServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceParameters) GoString() string {
  return s.String()
}

type ReportLowDelayCountryBandwidthServiceRequestHeader struct {
}

func (s ReportLowDelayCountryBandwidthServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportLowDelayCountryBandwidthServiceResponseHeader struct {
}

func (s ReportLowDelayCountryBandwidthServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportLowDelayCountryBandwidthServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportLogFlowIspProvinceServiceRequest struct {
  // {"en":"Start time 
  // 	1.The format is yyyy-MM-ddTHH:mm:ss+08:00; 
  // 	2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo; 
  // 	3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust); 
  // 	4.dateFrom and dateTo can be either both are specified or neither is specifies; 
  // 	5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天;(可联系技术支持调整)
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据;"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time 
  // 	1.The format is yyyy-MM-ddTHH:mm:ss+08:00; 
  // 	2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value.", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity, 5m: 5-minute granularity, 1h: 1-hour granularity", "zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度。默认5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个(可联系技术支持调整)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
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
  // 	1.Options are domain, province, isp, and more than one value can be entered; 
  // 	2.The data is displayed according to the specified dimension", "zh_CN":"分组维度
  //     1.可选值为domain、province、isp,可传入多个值;
  //     2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportLogFlowIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetDateFrom(v string) *ReportLogFlowIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetDateTo(v string) *ReportLogFlowIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetDataInterval(v string) *ReportLogFlowIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetDomain(v []*string) *ReportLogFlowIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetProvince(v []*string) *ReportLogFlowIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetIsp(v []*string) *ReportLogFlowIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportLogFlowIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportLogFlowIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportLogFlowIspProvinceServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportLogFlowIspProvinceServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLogFlowIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportLogFlowIspProvinceServiceResponse) SetCode(v string) *ReportLogFlowIspProvinceServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponse) SetMessage(v string) *ReportLogFlowIspProvinceServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponse) SetData(v []*ReportLogFlowIspProvinceServiceResponseData) *ReportLogFlowIspProvinceServiceResponse {
  s.Data = v
  return s
}

type ReportLogFlowIspProvinceServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportLogFlowIspProvinceServiceResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLogFlowIspProvinceServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportLogFlowIspProvinceServiceResponseData) SetDomain(v string) *ReportLogFlowIspProvinceServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponseData) SetIspData(v []*ReportLogFlowIspProvinceServiceResponseDataIspData) *ReportLogFlowIspProvinceServiceResponseData {
  s.IspData = v
  return s
}

type ReportLogFlowIspProvinceServiceResponseDataIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLogFlowIspProvinceServiceResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceResponseDataIspData) GoString() string {
  return s.String()
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspData) SetIsp(v string) *ReportLogFlowIspProvinceServiceResponseDataIspData {
  s.Isp = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspData) SetProvinceData(v []*ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData) *ReportLogFlowIspProvinceServiceResponseDataIspData {
  s.ProvinceData = v
  return s
}

type ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  Details []*ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData) SetProvince(v string) *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData) SetDetails(v []*ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails) *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceData {
  s.Details = v
  return s
}

type ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails struct     {
  // {"en":"Time:
  // 				1.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 				2.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;
  // 				3.Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间:
  // 
  // 1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。
  // 2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00。
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits of decimals allowed", "zh_CN":"流量值,单位为MB,保留两位小数"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {"en":"Bandwidth value. Unit is Mbps and 2 digits of decimals allowed", "zh_CN":"带宽值,单位为Mbps,保留两位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetTimestamp(v string) *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetFlow(v string) *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Flow = &v
  return s
}

func (s *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetBandwidth(v string) *ReportLogFlowIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Bandwidth = &v
  return s
}

type ReportLogFlowIspProvinceServicePaths struct {
}

func (s ReportLogFlowIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportLogFlowIspProvinceServiceParameters struct {
}

func (s ReportLogFlowIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportLogFlowIspProvinceServiceRequestHeader struct {
}

func (s ReportLogFlowIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportLogFlowIspProvinceServiceResponseHeader struct {
}

func (s ReportLogFlowIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportLogFlowIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type BandwidthVmRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到天,日期格式为yyyy-mm-dd 此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到天,日期格式为yyyy-mm-dd 此参数需与startdate参数配合,若存在date参数,则该参数无效"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type:
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1.If there isp multiple inputs,use ';' as demimeter.
  // 2.optional values of isp: refers to the ISP-section of appendix.
  // 3. If not specified,means all the isp.", "zh_CN":"要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  IspId *string `json:"ispId,omitempty" xml:"ispId,omitempty"`
  // {"en":"choose the ipv4 or ipv6.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"ipv4、ipv6，不填默认查全部"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
  // {"en":"the type of bandwidth:
  // 1. in : the in's bandwidth.
  // 2. out: the out's bandwidth.
  // 3. ein: the exin's bandwidth.
  // 4. eout: the exout's bandwidth.", "zh_CN":"带宽类型：in:流入带宽,out:流出带宽;ein:外网流入带宽,eout:外网流出带宽。不传或者空时默认为：in:流入带宽,out:流出带宽;ein:外网流入带宽,eout:外网流出带宽"}
  BandwidthType *string `json:"bandwidthType,omitempty" xml:"bandwidthType,omitempty"`
  // {"en":"The name of the node", "zh_CN":"节点名称"}
  Node *string `json:"node,omitempty" xml:"node,omitempty"`
  // {"en":"chargeInfo:show the charge method,charge value,the 95 of charge value", "zh_CN":"chargeInfo：计费时间计费方式、计费值、95值"}
  OptionalFields *string `json:"optionalFields,omitempty" xml:"optionalFields,omitempty"`
  // {"en":"method of data aggregation, can be left blank or filled with 'isp'", "zh_CN":"groupBy：数据聚合方式，可不填或填'isp'"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s BandwidthVmRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmRequest) GoString() string {
  return s.String()
}

func (s *BandwidthVmRequest) SetCust(v string) *BandwidthVmRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthVmRequest) SetDate(v string) *BandwidthVmRequest {
  s.Date = &v
  return s
}

func (s *BandwidthVmRequest) SetStartdate(v string) *BandwidthVmRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthVmRequest) SetEnddate(v string) *BandwidthVmRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthVmRequest) SetChannel(v string) *BandwidthVmRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthVmRequest) SetRegion(v string) *BandwidthVmRequest {
  s.Region = &v
  return s
}

func (s *BandwidthVmRequest) SetAccetype(v string) *BandwidthVmRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthVmRequest) SetDataformat(v string) *BandwidthVmRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthVmRequest) SetIsExactMatch(v string) *BandwidthVmRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthVmRequest) SetIspId(v string) *BandwidthVmRequest {
  s.IspId = &v
  return s
}

func (s *BandwidthVmRequest) SetIpProtocol(v string) *BandwidthVmRequest {
  s.IpProtocol = &v
  return s
}

func (s *BandwidthVmRequest) SetBandwidthType(v string) *BandwidthVmRequest {
  s.BandwidthType = &v
  return s
}

func (s *BandwidthVmRequest) SetNode(v string) *BandwidthVmRequest {
  s.Node = &v
  return s
}

func (s *BandwidthVmRequest) SetOptionalFields(v string) *BandwidthVmRequest {
  s.OptionalFields = &v
  return s
}

func (s *BandwidthVmRequest) SetGroupBy(v string) *BandwidthVmRequest {
  s.GroupBy = &v
  return s
}

type BandwidthVmResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthVmResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthVmResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponse) GoString() string {
  return s.String()
}

func (s *BandwidthVmResponse) SetProvider(v *BandwidthVmResponseProvider) *BandwidthVmResponse {
  s.Provider = v
  return s
}

type BandwidthVmResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *BandwidthVmResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthVmResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthVmResponseProvider) SetName(v string) *BandwidthVmResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthVmResponseProvider) SetType(v string) *BandwidthVmResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthVmResponseProvider) SetDate(v *BandwidthVmResponseProviderDate) *BandwidthVmResponseProvider {
  s.Date = v
  return s
}

type BandwidthVmResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthVmResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
  // {'en':'isp', 'zh_CN':'运营商'}
  Isp *BandwidthVmResponseProviderDateIsp `json:"isp,omitempty" xml:"isp,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthVmResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthVmResponseProviderDate) SetStartdate(v string) *BandwidthVmResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthVmResponseProviderDate) SetEnddate(v string) *BandwidthVmResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthVmResponseProviderDate) SetChannel(v *BandwidthVmResponseProviderDateChannel) *BandwidthVmResponseProviderDate {
  s.Channel = v
  return s
}

func (s *BandwidthVmResponseProviderDate) SetIsp(v *BandwidthVmResponseProviderDateIsp) *BandwidthVmResponseProviderDate {
  s.Isp = v
  return s
}

type BandwidthVmResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*BandwidthVmResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthVmResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthVmResponseProviderDateChannel) SetName(v string) *BandwidthVmResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthVmResponseProviderDateChannel) SetBandwidth(v []*BandwidthVmResponseProviderDateChannelBandwidth) *BandwidthVmResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthVmResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthVmResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthVmResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthVmResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthVmResponseProviderDateChannelBandwidth) SetText(v string) *BandwidthVmResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type BandwidthVmResponseProviderDateIsp struct {
  // {'en':'isp id', 'zh_CN':'isp id'}
  Ispid *string `json:"ispid,omitempty" xml:"ispid,omitempty" require:"true"`
  // {'en':'node', 'zh_CN':'节点名称'}
  Node *string `json:"node,omitempty" xml:"node,omitempty" require:"true"`
  // {'en':'peakTime', 'zh_CN':'带宽峰值时间'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'peakValue', 'zh_CN':'带宽峰值'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
}

func (s BandwidthVmResponseProviderDateIsp) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponseProviderDateIsp) GoString() string {
  return s.String()
}

func (s *BandwidthVmResponseProviderDateIsp) SetIspid(v string) *BandwidthVmResponseProviderDateIsp {
  s.Ispid = &v
  return s
}

func (s *BandwidthVmResponseProviderDateIsp) SetNode(v string) *BandwidthVmResponseProviderDateIsp {
  s.Node = &v
  return s
}

func (s *BandwidthVmResponseProviderDateIsp) SetPeakTime(v string) *BandwidthVmResponseProviderDateIsp {
  s.PeakTime = &v
  return s
}

func (s *BandwidthVmResponseProviderDateIsp) SetPeakValue(v string) *BandwidthVmResponseProviderDateIsp {
  s.PeakValue = &v
  return s
}

type BandwidthVmPaths struct {
}

func (s BandwidthVmPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmPaths) GoString() string {
  return s.String()
}

type BandwidthVmParameters struct {
}

func (s BandwidthVmParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmParameters) GoString() string {
  return s.String()
}

type BandwidthVmRequestHeader struct {
}

func (s BandwidthVmRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmRequestHeader) GoString() string {
  return s.String()
}

type BandwidthVmResponseHeader struct {
}

func (s BandwidthVmResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthVmResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthbyISPProvinceRequest struct {
  // {'en':'Start time', 'zh_CN':'开始时间：
  // &nbsp; 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // &nbsp; 2.不能大于当前时间；
  // &nbsp; 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // The 1. format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. the end time is larger than the start time.
  // 
  // 3. if the end time is greater than the current time, take the current time.
  // 
  // 4. DateFrom and dateTo are not uploaded, default query for the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days (technical support can be contacted to adjust).', 'zh_CN':'结束时间：
  // &nbsp; 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // &nbsp; 2.结束时间需大于开始时间；
  // &nbsp; 3.结束时间如果大于当前时间，取当前时间；
  // &nbsp; 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // &nbsp; 5.允许查询最大间隔：31天，即dateFrom和dateTo相差不能超过31天（可联系技术支持调整）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity
  // 
  // 1.        fiveminutes:   five minutes, hourly: one hour, daily: one day;
  // 
  // 2.        If   not specified, daily is set as the default value;
  // 
  // 3.        If   fiveminutes is specified as the value, then data is returned in actual   configured granularity when there is specific configuration to data collecting   granularity for the customer.', 'zh_CN':'数据粒度：
  // 1.支持5m（5分钟）
  // 2.不传默认5m。'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'area:
  // 
  // 1. Do not pass the default query all areas.
  // 
  // 2. Support to pass multiple regions, the optional values are: cn, hk, ov, apac, am, emea, fg.', 'zh_CN':'区域：
  // 1.不传默认查询全部区域。
  // 2.支持传递多个区域，可选值为：cn、hk、ov、apac、am、emea、fg。'}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {'en':'domain:
  // 
  // 1. The maximum number of passable domain names is 20 by default (you can contact technical support adjustment).
  // 
  // 2. Query all the domain names under the account when the entry is not passed, but you cannot query (error) when the number of domain names under the account exceeds the limit.', 'zh_CN':'域名：
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整），
  // 2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'By default, all provinces are queried, and the provinces are transferred. The province information code table is detailed in the appendix of the overview page.', 'zh_CN':'默认查询全部省份，传递省份，省份信息码表详见概览页附录说明章节'}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {'en':'The default query is for all operators, the operators and operator information codes are detailed in the appendix of the overview page.', 'zh_CN':'默认查询全部运营商，传递运营商，运营商信息码表详见概览页附录说明章节'}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
}

func (s QueryBandwidthbyISPProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDateFrom(v string) *QueryBandwidthbyISPProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDateTo(v string) *QueryBandwidthbyISPProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDataInterval(v string) *QueryBandwidthbyISPProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetAreaCode(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.AreaCode = v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDomain(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetProvince(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetIsp(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.Isp = v
  return s
}

type QueryBandwidthbyISPProvinceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*QueryBandwidthbyISPProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthbyISPProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceResponse) SetResult(v []*QueryBandwidthbyISPProvinceResponseResult) *QueryBandwidthbyISPProvinceResponse {
  s.Result = v
  return s
}

type QueryBandwidthbyISPProvinceResponseResult struct     {
  BandwidthData []*QueryBandwidthbyISPProvinceResponseResultBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthbyISPProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceResponseResult) SetBandwidthData(v []*QueryBandwidthbyISPProvinceResponseResultBandwidthData) *QueryBandwidthbyISPProvinceResponseResult {
  s.BandwidthData = v
  return s
}

type QueryBandwidthbyISPProvinceResponseResultBandwidthData struct     {
  // {'en':'Date
  // 
  // 1.        When   the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm;   the data value of every time slice represents the data value within the   previous time granularity range. 
  // 2.        Return   the time slice contained in start time and the time slice contained in end   time.', 'zh_CN':'时间
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd &nbsp; HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值，比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。
  // 2.返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'The bandwidth value, the corresponding value value of each time slice, shows the cumulative data value represented by this time slice.
  // 
  // 
  // 
  // Unit Mbps, keep 2 decimal digits.', 'zh_CN':'带宽值，每个时间片的对应value值都展示这个时间片代表的完整数据累计值。
  // 单位Mbps，保留2位小数。'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryBandwidthbyISPProvinceResponseResultBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponseResultBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceResponseResultBandwidthData) SetTimestamp(v string) *QueryBandwidthbyISPProvinceResponseResultBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceResponseResultBandwidthData) SetValue(v string) *QueryBandwidthbyISPProvinceResponseResultBandwidthData {
  s.Value = &v
  return s
}

type QueryBandwidthbyISPProvincePaths struct {
}

func (s QueryBandwidthbyISPProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvincePaths) GoString() string {
  return s.String()
}

type QueryBandwidthbyISPProvinceParameters struct {
}

func (s QueryBandwidthbyISPProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceParameters) GoString() string {
  return s.String()
}

type QueryBandwidthbyISPProvinceRequestHeader struct {
}

func (s QueryBandwidthbyISPProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthbyISPProvinceResponseHeader struct {
}

func (s QueryBandwidthbyISPProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponseHeader) GoString() string {
  return s.String()
}




type BandwidthPeakRankingRequest struct {
  // {"en":"Specifies the query date:\n1.With format yyyy-mm-dd.\n2.If not specified,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.\n2.With format yyyy-mm-dd.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.\n2.With format yyyy-mm-dd.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1.If there are multiple inputs,use  ';' as separator.\n2.If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2.If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=ispCode","en":"The abbreviations of the ISPs to be queried. For multiple ISPs, please separate them with a semicolon ';'. Note: Only when the region is specified as 'cn' does the ISP information take effect. If not selected or left blank, all ISPs will be included by default.","zh_CN":"要查询的运营商的缩写，多个isp请用英文分号';'分隔开。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"acceleration type.\n1.If there are multiple inputs,use ';' as separator.\n2.If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1.optional values:xml, json.\n2.'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:\n1.'true' as default.\n2. If not 'true',it will query data of channels that ends with any item of input 'channel's.","zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"Different data types.\n1.optional values:1,2,3.\n2.'2' means bandwidth of http.'3' means bandwidth of https.'1' mean the total bandwidth.\n3.If specified 2 or 3, ISP parameter is not supported.","zh_CN":"datatype=1时，输出总带宽；datatype=2时输出http的带宽；datatype=3时，输出https的带宽。默认datatype=1。当datatype=2或者3时，不支持isp入参。"}
  Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
  // {"en":"Greenwich Mean Time, parameter format GMT%2b09:00 represents Eastern 9th Zone, GMT-09:00 represents Western 9th Zone, if not specified, it defaults to local time zone (Eastern 8th Zone).","zh_CN":"格林尼治时区，参数格式 GMT%2b09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区(东八区)"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s BandwidthPeakRankingRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingRequest) GoString() string {
  return s.String()
}

func (s *BandwidthPeakRankingRequest) SetDate(v string) *BandwidthPeakRankingRequest {
  s.Date = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetStartdate(v string) *BandwidthPeakRankingRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetEnddate(v string) *BandwidthPeakRankingRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetChannel(v string) *BandwidthPeakRankingRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetRegion(v string) *BandwidthPeakRankingRequest {
  s.Region = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetIsp(v string) *BandwidthPeakRankingRequest {
  s.Isp = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetAccetype(v string) *BandwidthPeakRankingRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetDataformat(v string) *BandwidthPeakRankingRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetIsExactMatch(v string) *BandwidthPeakRankingRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetDatatype(v string) *BandwidthPeakRankingRequest {
  s.Datatype = &v
  return s
}

func (s *BandwidthPeakRankingRequest) SetTimezone(v string) *BandwidthPeakRankingRequest {
  s.Timezone = &v
  return s
}

type BandwidthPeakRankingRequestHeader struct {
}

func (s BandwidthPeakRankingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingRequestHeader) GoString() string {
  return s.String()
}

type BandwidthPeakRankingPaths struct {
}

func (s BandwidthPeakRankingPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingPaths) GoString() string {
  return s.String()
}

type BandwidthPeakRankingParameters struct {
}

func (s BandwidthPeakRankingParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingParameters) GoString() string {
  return s.String()
}

type BandwidthPeakRankingResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *BandwidthPeakRankingResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthPeakRankingResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingResponse) GoString() string {
  return s.String()
}

func (s *BandwidthPeakRankingResponse) SetProvider(v *BandwidthPeakRankingResponseProvider) *BandwidthPeakRankingResponse {
  s.Provider = v
  return s
}

type BandwidthPeakRankingResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"数据"}
  Date *BandwidthPeakRankingResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthPeakRankingResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthPeakRankingResponseProvider) SetName(v string) *BandwidthPeakRankingResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthPeakRankingResponseProvider) SetType(v string) *BandwidthPeakRankingResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthPeakRankingResponseProvider) SetDate(v *BandwidthPeakRankingResponseProviderDate) *BandwidthPeakRankingResponseProvider {
  s.Date = v
  return s
}

type BandwidthPeakRankingResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始日期"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束日期"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"channelPeak","zh_CN":"频道峰值数据"}
  ChannelPeak *BandwidthPeakRankingResponseProviderDateChannelPeak `json:"channelPeak,omitempty" xml:"channelPeak,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthPeakRankingResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthPeakRankingResponseProviderDate) SetStartdate(v string) *BandwidthPeakRankingResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthPeakRankingResponseProviderDate) SetEnddate(v string) *BandwidthPeakRankingResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthPeakRankingResponseProviderDate) SetChannelPeak(v *BandwidthPeakRankingResponseProviderDateChannelPeak) *BandwidthPeakRankingResponseProviderDate {
  s.ChannelPeak = v
  return s
}

type BandwidthPeakRankingResponseProviderDateChannelPeak struct {
  // {"en":"channel","zh_CN":"频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"peakTime","zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"peakvalue(Mbps)","zh_CN":"带宽峰值，单位Mbps"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"total traffic,unit GB","zh_CN":"总流量，单位GB"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
}

func (s BandwidthPeakRankingResponseProviderDateChannelPeak) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingResponseProviderDateChannelPeak) GoString() string {
  return s.String()
}

func (s *BandwidthPeakRankingResponseProviderDateChannelPeak) SetChannel(v string) *BandwidthPeakRankingResponseProviderDateChannelPeak {
  s.Channel = &v
  return s
}

func (s *BandwidthPeakRankingResponseProviderDateChannelPeak) SetPeakTime(v string) *BandwidthPeakRankingResponseProviderDateChannelPeak {
  s.PeakTime = &v
  return s
}

func (s *BandwidthPeakRankingResponseProviderDateChannelPeak) SetPeakValue(v string) *BandwidthPeakRankingResponseProviderDateChannelPeak {
  s.PeakValue = &v
  return s
}

func (s *BandwidthPeakRankingResponseProviderDateChannelPeak) SetTotalFlow(v string) *BandwidthPeakRankingResponseProviderDateChannelPeak {
  s.TotalFlow = &v
  return s
}

type BandwidthPeakRankingResponseHeader struct {
}

func (s BandwidthPeakRankingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthPeakRankingResponseHeader) GoString() string {
  return s.String()
}




type BandwidthEcdnRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"The option without specifying a value will return all types of data:
  // 1)ecdnBandwidth: will only return ECDN data.
  // 2)cdnBandwidth: will only return CDN details data.
  // 3)cdnPlusEcdnBandwidth: will only return ECDN+CDN details dat.", "zh_CN":"不传返回所有类型数据。ecdnBandwidth：只返回ECDN数据；cdnBandwidth：只返回CDN明细数据；cdnPlusEcdnBandwidth：只返回ECDN+cdn明细数据"}
  Datasource *string `json:"datasource,omitempty" xml:"datasource,omitempty"`
}

func (s BandwidthEcdnRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnRequest) GoString() string {
  return s.String()
}

func (s *BandwidthEcdnRequest) SetCust(v string) *BandwidthEcdnRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthEcdnRequest) SetDate(v string) *BandwidthEcdnRequest {
  s.Date = &v
  return s
}

func (s *BandwidthEcdnRequest) SetStartdate(v string) *BandwidthEcdnRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthEcdnRequest) SetEnddate(v string) *BandwidthEcdnRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthEcdnRequest) SetChannel(v string) *BandwidthEcdnRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthEcdnRequest) SetTimezone(v string) *BandwidthEcdnRequest {
  s.Timezone = &v
  return s
}

func (s *BandwidthEcdnRequest) SetRegion(v string) *BandwidthEcdnRequest {
  s.Region = &v
  return s
}

func (s *BandwidthEcdnRequest) SetAccetype(v string) *BandwidthEcdnRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthEcdnRequest) SetDataformat(v string) *BandwidthEcdnRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthEcdnRequest) SetDatasource(v string) *BandwidthEcdnRequest {
  s.Datasource = &v
  return s
}

type BandwidthEcdnResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthEcdnResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthEcdnResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnResponse) GoString() string {
  return s.String()
}

func (s *BandwidthEcdnResponse) SetProvider(v *BandwidthEcdnResponseProvider) *BandwidthEcdnResponse {
  s.Provider = v
  return s
}

type BandwidthEcdnResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'ecdn数据'}
  Date *BandwidthEcdnResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthEcdnResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthEcdnResponseProvider) SetName(v string) *BandwidthEcdnResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthEcdnResponseProvider) SetType(v string) *BandwidthEcdnResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthEcdnResponseProvider) SetDate(v *BandwidthEcdnResponseProviderDate) *BandwidthEcdnResponseProvider {
  s.Date = v
  return s
}

type BandwidthEcdnResponseProviderDate struct {
  // {'en':'chartDataList', 'zh_CN':'带宽明细'}
  ChartDataList []*BandwidthEcdnResponseProviderDateChartDataList `json:"chartDataList,omitempty" xml:"chartDataList,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthEcdnResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthEcdnResponseProviderDate) SetChartDataList(v []*BandwidthEcdnResponseProviderDateChartDataList) *BandwidthEcdnResponseProviderDate {
  s.ChartDataList = v
  return s
}

type BandwidthEcdnResponseProviderDateChartDataList struct     {
  // {'en':'ecdnBandwidth', 'zh_CN':'支持：[ecdnBandwidth],[cdnBandwidth],[cdnPlusEcdnBandwidth] 3种类型带宽数据'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽明细'}
  Data []*BandwidthEcdnResponseProviderDateChartDataListData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthEcdnResponseProviderDateChartDataList) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnResponseProviderDateChartDataList) GoString() string {
  return s.String()
}

func (s *BandwidthEcdnResponseProviderDateChartDataList) SetName(v string) *BandwidthEcdnResponseProviderDateChartDataList {
  s.Name = &v
  return s
}

func (s *BandwidthEcdnResponseProviderDateChartDataList) SetData(v []*BandwidthEcdnResponseProviderDateChartDataListData) *BandwidthEcdnResponseProviderDateChartDataList {
  s.Data = v
  return s
}

type BandwidthEcdnResponseProviderDateChartDataListData struct     {
  // {'en':'time of every 5 duration,with format yyyy-mmm-dd hh:MM:ss', 'zh_CN':'ecdn带宽5分钟粒度时间，格式yyyy-mm-dd hh:MM:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'displaying the Bandwidth(Mbps)', 'zh_CN':'ecdn带宽(Mbps)'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthEcdnResponseProviderDateChartDataListData) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnResponseProviderDateChartDataListData) GoString() string {
  return s.String()
}

func (s *BandwidthEcdnResponseProviderDateChartDataListData) SetTime(v string) *BandwidthEcdnResponseProviderDateChartDataListData {
  s.Time = &v
  return s
}

func (s *BandwidthEcdnResponseProviderDateChartDataListData) SetText(v string) *BandwidthEcdnResponseProviderDateChartDataListData {
  s.Text = &v
  return s
}

type BandwidthEcdnPaths struct {
}

func (s BandwidthEcdnPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnPaths) GoString() string {
  return s.String()
}

type BandwidthEcdnParameters struct {
}

func (s BandwidthEcdnParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnParameters) GoString() string {
  return s.String()
}

type BandwidthEcdnRequestHeader struct {
}

func (s BandwidthEcdnRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnRequestHeader) GoString() string {
  return s.String()
}

type BandwidthEcdnResponseHeader struct {
}

func (s BandwidthEcdnResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthEcdnResponseHeader) GoString() string {
  return s.String()
}




type QuicLogBandwidthRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"2 types. 0: QUIC bandwidth; 1: Log bandwidth deduction QUIC. If not filled, the default query is: 0.", "zh_CN":"2种类型。0：QUIC带宽；1：日志带宽扣减QUIC。不填默认查询：0。"}
  Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
}

func (s QuicLogBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QuicLogBandwidthRequest) SetCust(v string) *QuicLogBandwidthRequest {
  s.Cust = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetDate(v string) *QuicLogBandwidthRequest {
  s.Date = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetStartdate(v string) *QuicLogBandwidthRequest {
  s.Startdate = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetEnddate(v string) *QuicLogBandwidthRequest {
  s.Enddate = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetChannel(v string) *QuicLogBandwidthRequest {
  s.Channel = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetTimezone(v string) *QuicLogBandwidthRequest {
  s.Timezone = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetRegion(v string) *QuicLogBandwidthRequest {
  s.Region = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetDataformat(v string) *QuicLogBandwidthRequest {
  s.Dataformat = &v
  return s
}

func (s *QuicLogBandwidthRequest) SetDatatype(v string) *QuicLogBandwidthRequest {
  s.Datatype = &v
  return s
}

type QuicLogBandwidthResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QuicLogBandwidthResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QuicLogBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QuicLogBandwidthResponse) SetProvider(v *QuicLogBandwidthResponseProvider) *QuicLogBandwidthResponse {
  s.Provider = v
  return s
}

type QuicLogBandwidthResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *QuicLogBandwidthResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QuicLogBandwidthResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthResponseProvider) GoString() string {
  return s.String()
}

func (s *QuicLogBandwidthResponseProvider) SetName(v string) *QuicLogBandwidthResponseProvider {
  s.Name = &v
  return s
}

func (s *QuicLogBandwidthResponseProvider) SetType(v string) *QuicLogBandwidthResponseProvider {
  s.Type = &v
  return s
}

func (s *QuicLogBandwidthResponseProvider) SetResultType(v string) *QuicLogBandwidthResponseProvider {
  s.ResultType = &v
  return s
}

func (s *QuicLogBandwidthResponseProvider) SetDate(v *QuicLogBandwidthResponseProviderDate) *QuicLogBandwidthResponseProvider {
  s.Date = v
  return s
}

type QuicLogBandwidthResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *QuicLogBandwidthResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s QuicLogBandwidthResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QuicLogBandwidthResponseProviderDate) SetStartdate(v string) *QuicLogBandwidthResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QuicLogBandwidthResponseProviderDate) SetEnddate(v string) *QuicLogBandwidthResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QuicLogBandwidthResponseProviderDate) SetChannel(v *QuicLogBandwidthResponseProviderDateChannel) *QuicLogBandwidthResponseProviderDate {
  s.Channel = v
  return s
}

type QuicLogBandwidthResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*QuicLogBandwidthResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s QuicLogBandwidthResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *QuicLogBandwidthResponseProviderDateChannel) SetName(v string) *QuicLogBandwidthResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *QuicLogBandwidthResponseProviderDateChannel) SetBandwidth(v []*QuicLogBandwidthResponseProviderDateChannelBandwidth) *QuicLogBandwidthResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type QuicLogBandwidthResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽Mbps'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QuicLogBandwidthResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *QuicLogBandwidthResponseProviderDateChannelBandwidth) SetTime(v string) *QuicLogBandwidthResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *QuicLogBandwidthResponseProviderDateChannelBandwidth) SetText(v string) *QuicLogBandwidthResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type QuicLogBandwidthPaths struct {
}

func (s QuicLogBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthPaths) GoString() string {
  return s.String()
}

type QuicLogBandwidthParameters struct {
}

func (s QuicLogBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthParameters) GoString() string {
  return s.String()
}

type QuicLogBandwidthRequestHeader struct {
}

func (s QuicLogBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QuicLogBandwidthResponseHeader struct {
}

func (s QuicLogBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuicLogBandwidthResponseHeader) GoString() string {
  return s.String()
}




type BandwidthLogEcdnRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"The option without specifying a value will return all types of data:
  // 1)ecdnBandwidth: will only return ECDN data.
  // 2)cdnBandwidth: will only return CDN details data.
  // 3)cdnPlusEcdnBandwidth: will only return ECDN+CDN details dat.", "zh_CN":"不传返回所有类型数据。ecdnBandwidth：只返回ECDN数据；cdnBandwidth：只返回CDN明细数据；cdnPlusEcdnBandwidth：只返回ECDN+cdn明细数据"}
  Datasource *string `json:"datasource,omitempty" xml:"datasource,omitempty"`
}

func (s BandwidthLogEcdnRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnRequest) GoString() string {
  return s.String()
}

func (s *BandwidthLogEcdnRequest) SetCust(v string) *BandwidthLogEcdnRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetDate(v string) *BandwidthLogEcdnRequest {
  s.Date = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetStartdate(v string) *BandwidthLogEcdnRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetEnddate(v string) *BandwidthLogEcdnRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetChannel(v string) *BandwidthLogEcdnRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetTimezone(v string) *BandwidthLogEcdnRequest {
  s.Timezone = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetRegion(v string) *BandwidthLogEcdnRequest {
  s.Region = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetAccetype(v string) *BandwidthLogEcdnRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetDataformat(v string) *BandwidthLogEcdnRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthLogEcdnRequest) SetDatasource(v string) *BandwidthLogEcdnRequest {
  s.Datasource = &v
  return s
}

type BandwidthLogEcdnResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthLogEcdnResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthLogEcdnResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnResponse) GoString() string {
  return s.String()
}

func (s *BandwidthLogEcdnResponse) SetProvider(v *BandwidthLogEcdnResponseProvider) *BandwidthLogEcdnResponse {
  s.Provider = v
  return s
}

type BandwidthLogEcdnResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'ecdn数据'}
  Date *BandwidthLogEcdnResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthLogEcdnResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthLogEcdnResponseProvider) SetName(v string) *BandwidthLogEcdnResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthLogEcdnResponseProvider) SetType(v string) *BandwidthLogEcdnResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthLogEcdnResponseProvider) SetDate(v *BandwidthLogEcdnResponseProviderDate) *BandwidthLogEcdnResponseProvider {
  s.Date = v
  return s
}

type BandwidthLogEcdnResponseProviderDate struct {
  // {'en':'chartDataList', 'zh_CN':'带宽明细'}
  ChartDataList []*BandwidthLogEcdnResponseProviderDateChartDataList `json:"chartDataList,omitempty" xml:"chartDataList,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthLogEcdnResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthLogEcdnResponseProviderDate) SetChartDataList(v []*BandwidthLogEcdnResponseProviderDateChartDataList) *BandwidthLogEcdnResponseProviderDate {
  s.ChartDataList = v
  return s
}

type BandwidthLogEcdnResponseProviderDateChartDataList struct     {
  // {'en':'ecdnBandwidth', 'zh_CN':'支持：[ecdnBandwidth],[cdnBandwidth],[cdnPlusEcdnBandwidth] 3种类型带宽数据'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽明细'}
  Data []*BandwidthLogEcdnResponseProviderDateChartDataListData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthLogEcdnResponseProviderDateChartDataList) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnResponseProviderDateChartDataList) GoString() string {
  return s.String()
}

func (s *BandwidthLogEcdnResponseProviderDateChartDataList) SetName(v string) *BandwidthLogEcdnResponseProviderDateChartDataList {
  s.Name = &v
  return s
}

func (s *BandwidthLogEcdnResponseProviderDateChartDataList) SetData(v []*BandwidthLogEcdnResponseProviderDateChartDataListData) *BandwidthLogEcdnResponseProviderDateChartDataList {
  s.Data = v
  return s
}

type BandwidthLogEcdnResponseProviderDateChartDataListData struct     {
  // {'en':'time of every 5 duration,with format yyyy-mmm-dd hh:MM:ss', 'zh_CN':'ecdn带宽5分钟粒度时间，格式yyyy-mm-dd hh:MM:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'displaying the Bandwidth(Mbps)', 'zh_CN':'ecdn带宽(Mbps)'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthLogEcdnResponseProviderDateChartDataListData) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnResponseProviderDateChartDataListData) GoString() string {
  return s.String()
}

func (s *BandwidthLogEcdnResponseProviderDateChartDataListData) SetTime(v string) *BandwidthLogEcdnResponseProviderDateChartDataListData {
  s.Time = &v
  return s
}

func (s *BandwidthLogEcdnResponseProviderDateChartDataListData) SetText(v string) *BandwidthLogEcdnResponseProviderDateChartDataListData {
  s.Text = &v
  return s
}

type BandwidthLogEcdnPaths struct {
}

func (s BandwidthLogEcdnPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnPaths) GoString() string {
  return s.String()
}

type BandwidthLogEcdnParameters struct {
}

func (s BandwidthLogEcdnParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnParameters) GoString() string {
  return s.String()
}

type BandwidthLogEcdnRequestHeader struct {
}

func (s BandwidthLogEcdnRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnRequestHeader) GoString() string {
  return s.String()
}

type BandwidthLogEcdnResponseHeader struct {
}

func (s BandwidthLogEcdnResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLogEcdnResponseHeader) GoString() string {
  return s.String()
}




type QueryIPV6BandwidthOfeachISPandProvinceRequest struct {
  // {"en":"Start time
  // 
  // 1.The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2.Cannot be greater than the current time
  // 3.Get up to the last six months (183 days) of data.", "zh_CN":"开始时间:
  // 
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 
  // 2.不能大于当前时间
  // 
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;
  // 3.If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default;
  // 4.Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days)", "zh_CN":"结束时间:
  // 
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 
  // 4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The default upper limit to domains that can be entered is 200 (Contact technical support to adjust);
  // 2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为200个(可联系技术支持调整);
  // 
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. Support 5m (5 minutes granularity),1h (1 hour granularity)
  // 2. The default value is 5m", "zh_CN":"数据粒度:
  // 
  // 1.支持5m(5分钟)、1h(1小时)
  // 
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
  // {"en":"IP type
  // 1.Optional values:IPV6,IPV4
  // 2.Query all IP type by default.", "zh_CN":"IP类型:
  // 
  // 1.可选值为 IPV6.IPV4
  // 2.不传默认查询全部"}
  IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
  // {"en":"Group dimension:
  // 1.Default response aggregation data without group
  // 2.Optional values:domain,province,isp,you can pass in single or multiple values
  // 3.The detailed data will be displayed according to the dimension.For example,if the dimension is isp,the detail data will be group by each isp.", "zh_CN":"分组关键词:
  // 1.默认聚合展示;
  // 2.可选值为domain.province.isp,可传入多个值;
  // 3.传入关键词则代表需要按照关键词对应的值展示明细; 例如groupBy传入isp,则isp维度需要明细展示;当没有传递isp,则代表isp聚合展示,同时isp节点则不返回。其他province和domain相同逻辑。 例如:传递'groupBy':   ['domain','province'],则ispData下的isp节点无需返回。 { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp':   '中国电信', 'provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDateTo(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetIPType(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.IPType = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryIPV6BandwidthOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponse) SetResult(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResult) *QueryIPV6BandwidthOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResult) SetIspData(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) *QueryIPV6BandwidthOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  BandwidthData []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) SetBandwidthData(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.BandwidthData = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData struct     {
  // {"en":"Time
  // 1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00.
  // 3. Returns the time slice contained in the start time and end time.", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth,unit is Mbps,Keep 2 decimal places", "zh_CN":"带宽值,单位Mbps,保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) SetTimestamp(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) SetValue(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData {
  s.Value = &v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvincePaths struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryIPV6BandwidthOfeachISPandProvinceParameters struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryIPV6BandwidthOfeachISPandProvinceRequestHeader struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseHeader struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type ReportBandwidthRealTimeEdgeServiceRequest struct {
  // {"en":"Start time:
  // 1.Start time: time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (December 2rd, 2016, 10:00 a.m., Beijing Time);
  // 2.Not greater than the current time
  // 3.The most recent half-year (183 days) data can be obtained", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is 2016-12-02T10:00:00+08:00
  // 2.End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3.If both fields of dataFrom and dateTo are left empty, then data in the last 1 hours will be queried by default; if one field is filled and one is left empty, then exception will occur.
  // 4.Maximum time range allowable for query: The default value is 1 hour, that is, the difference between dateFrom and dateTo cannot exceed 1 hour (you can contact technical support to adjust it, the maximum is 31 days).", "zh_CN":"结束时间:
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:默认1小时，即dateFrom和dateTo相差不能超过1小时（可联系技术支持调整，最长31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The default upper limit of domains that can be entered is 20 (if you want to adjust, please, contact technical support);
  // 2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 3.域名超过上限,报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. Support 1m (1 minute granularity),5m (5 minutes granularity)
  // 2. The default value is 1m", "zh_CN":"数据粒度:不传默认1m
  // 1.支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportBandwidthRealTimeEdgeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRealTimeEdgeServiceRequest) SetDateFrom(v string) *ReportBandwidthRealTimeEdgeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceRequest) SetDateTo(v string) *ReportBandwidthRealTimeEdgeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceRequest) SetDomain(v []*string) *ReportBandwidthRealTimeEdgeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceRequest) SetDataInterval(v string) *ReportBandwidthRealTimeEdgeServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportBandwidthRealTimeEdgeServiceResponse struct {
  Result []*ReportBandwidthRealTimeEdgeServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthRealTimeEdgeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRealTimeEdgeServiceResponse) SetResult(v []*ReportBandwidthRealTimeEdgeServiceResponseResult) *ReportBandwidthRealTimeEdgeServiceResponse {
  s.Result = v
  return s
}

type ReportBandwidthRealTimeEdgeServiceResponseResult struct     {
  // {"en":"Peak Bandwidth,unit is Mbps,example(9811.21Mbps)", "zh_CN":"峰值带宽 Mbps,示例 (931556.21 Mbps)"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Time of peak bandwidth,example(2019-02-13 18:01)", "zh_CN":"峰值时间,示例(2019-02-13 18:01);"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Edge total traffic,example(74099.91MB)", "zh_CN":"边缘总流量,示例 ( 74099.92 MB )"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  BandwidthData []*ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthRealTimeEdgeServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRealTimeEdgeServiceResponseResult) SetPeakValue(v string) *ReportBandwidthRealTimeEdgeServiceResponseResult {
  s.PeakValue = &v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceResponseResult) SetPeakTime(v string) *ReportBandwidthRealTimeEdgeServiceResponseResult {
  s.PeakTime = &v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceResponseResult) SetTotal(v string) *ReportBandwidthRealTimeEdgeServiceResponseResult {
  s.Total = &v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceResponseResult) SetBandwidthData(v []*ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData) *ReportBandwidthRealTimeEdgeServiceResponseResult {
  s.BandwidthData = v
  return s
}

type ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData struct     {
  // {"en":"The data granularity is 1 minute,the format is yyyy-MM-dd HH:mm", "zh_CN":"数据粒度为1分钟,格式为yyyy-MM-dd HH:mm;"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Edge bandwidth,the unit is Mbps,keep 2 decimal places", "zh_CN":"边缘带宽值,单位Mbps,保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData) SetTimestamp(v string) *ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData) SetValue(v string) *ReportBandwidthRealTimeEdgeServiceResponseResultBandwidthData {
  s.Value = &v
  return s
}

type ReportBandwidthRealTimeEdgeServicePaths struct {
}

func (s ReportBandwidthRealTimeEdgeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServicePaths) GoString() string {
  return s.String()
}

type ReportBandwidthRealTimeEdgeServiceParameters struct {
}

func (s ReportBandwidthRealTimeEdgeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceParameters) GoString() string {
  return s.String()
}

type ReportBandwidthRealTimeEdgeServiceRequestHeader struct {
}

func (s ReportBandwidthRealTimeEdgeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportBandwidthRealTimeEdgeServiceResponseHeader struct {
}

func (s ReportBandwidthRealTimeEdgeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthRealTimeEdgeServiceResponseHeader) GoString() string {
  return s.String()
}




type BandwidthMiddleRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号“;”分隔开，如查询大陆及亚太区域，参数填写为：“region=cn;apac”。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type:
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":" 频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1, get the merged result.
  // 2.If  specified 2,get the separate result.
  // 3.If specifed 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":" 结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为“1”。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"Different data types:
  // 1.optional values: hit,miss.
  // 2.If  there are multiple inputs,use ';' as separator.
  // 3.If not specified,it means the merged value of all types.", "zh_CN":"可选值：hit（hit带宽），miss（miss带宽）。多个以英文分号分隔，不选或为空默认展示合计值。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Different father types:
  // 1.optional values: stafu,dynfu,all
  // 2.If there are multiple inputs,use ';' as separator.
  // 3.If not specified,it means 'all'.", "zh_CN":"可选值：stafu（静态父），dynfu（动态父），all(全选)；默认all（全选）。"}
  FatherType *string `json:"fatherType,omitempty" xml:"fatherType,omitempty"`
  // {"en":"Return traffic details, 1: mandatory; 0: optional. The default value is 0.", "zh_CN":"返回流量明细，1：需要；0：不需要。默认为0."}
  NeedFlow *string `json:"needFlow,omitempty" xml:"needFlow,omitempty"`
}

func (s BandwidthMiddleRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleRequest) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleRequest) SetCust(v string) *BandwidthMiddleRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthMiddleRequest) SetDate(v string) *BandwidthMiddleRequest {
  s.Date = &v
  return s
}

func (s *BandwidthMiddleRequest) SetStartdate(v string) *BandwidthMiddleRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthMiddleRequest) SetEnddate(v string) *BandwidthMiddleRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthMiddleRequest) SetChannel(v string) *BandwidthMiddleRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthMiddleRequest) SetRegion(v string) *BandwidthMiddleRequest {
  s.Region = &v
  return s
}

func (s *BandwidthMiddleRequest) SetAccetype(v string) *BandwidthMiddleRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthMiddleRequest) SetDataformat(v string) *BandwidthMiddleRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthMiddleRequest) SetIsExactMatch(v string) *BandwidthMiddleRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthMiddleRequest) SetResultType(v string) *BandwidthMiddleRequest {
  s.ResultType = &v
  return s
}

func (s *BandwidthMiddleRequest) SetType(v string) *BandwidthMiddleRequest {
  s.Type = &v
  return s
}

func (s *BandwidthMiddleRequest) SetFatherType(v string) *BandwidthMiddleRequest {
  s.FatherType = &v
  return s
}

func (s *BandwidthMiddleRequest) SetNeedFlow(v string) *BandwidthMiddleRequest {
  s.NeedFlow = &v
  return s
}

type BandwidthMiddleResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthMiddleResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthMiddleResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponse) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponse) SetProvider(v *BandwidthMiddleResponseProvider) *BandwidthMiddleResponse {
  s.Provider = v
  return s
}

type BandwidthMiddleResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'中间缓存带宽数据'}
  Date *BandwidthMiddleResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthMiddleResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProvider) SetName(v string) *BandwidthMiddleResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthMiddleResponseProvider) SetType(v string) *BandwidthMiddleResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthMiddleResponseProvider) SetResultType(v string) *BandwidthMiddleResponseProvider {
  s.ResultType = &v
  return s
}

func (s *BandwidthMiddleResponseProvider) SetDate(v *BandwidthMiddleResponseProviderDate) *BandwidthMiddleResponseProvider {
  s.Date = v
  return s
}

type BandwidthMiddleResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthMiddleResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthMiddleResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProviderDate) SetStartdate(v string) *BandwidthMiddleResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDate) SetEnddate(v string) *BandwidthMiddleResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDate) SetChannel(v *BandwidthMiddleResponseProviderDateChannel) *BandwidthMiddleResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthMiddleResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'中间缓存带宽数据'}
  Bandwidth []*BandwidthMiddleResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthMiddleResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProviderDateChannel) SetName(v string) *BandwidthMiddleResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannel) SetBandwidth(v []*BandwidthMiddleResponseProviderDateChannelBandwidth) *BandwidthMiddleResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthMiddleResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'summary', 'zh_CN':'带宽合计'}
  Summary *string `json:"summary,omitempty" xml:"summary,omitempty" require:"true"`
  // {'en':'stafuHit', 'zh_CN':'静态Hit带宽'}
  StafuHit *string `json:"stafuHit,omitempty" xml:"stafuHit,omitempty" require:"true"`
  // {'en':'dynfuHit', 'zh_CN':'动态Hit带宽'}
  DynfuHit *string `json:"dynfuHit,omitempty" xml:"dynfuHit,omitempty" require:"true"`
  // {'en':'stafuMiss', 'zh_CN':'静态Miss带宽'}
  StafuMiss *string `json:"stafuMiss,omitempty" xml:"stafuMiss,omitempty" require:"true"`
  // {'en':'dynfuMiss', 'zh_CN':'动态Miss带宽'}
  DynfuMiss *string `json:"dynfuMiss,omitempty" xml:"dynfuMiss,omitempty" require:"true"`
  // {'en':'thirdWs', 'zh_CN':'带宽'}
  ThirdWs *string `json:"thirdWs,omitempty" xml:"thirdWs,omitempty" require:"true"`
}

func (s BandwidthMiddleResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetSummary(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.Summary = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetStafuHit(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.StafuHit = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetDynfuHit(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.DynfuHit = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetStafuMiss(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.StafuMiss = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetDynfuMiss(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.DynfuMiss = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetThirdWs(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.ThirdWs = &v
  return s
}

type BandwidthMiddlePaths struct {
}

func (s BandwidthMiddlePaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddlePaths) GoString() string {
  return s.String()
}

type BandwidthMiddleParameters struct {
}

func (s BandwidthMiddleParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleParameters) GoString() string {
  return s.String()
}

type BandwidthMiddleRequestHeader struct {
}

func (s BandwidthMiddleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleRequestHeader) GoString() string {
  return s.String()
}

type BandwidthMiddleResponseHeader struct {
}

func (s BandwidthMiddleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseHeader) GoString() string {
  return s.String()
}




type WsiInfoRequest struct {
  // {"en":"cust_en_name of sub-client. When a merged-account wants to view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"rbi mapping account. If there is a pass, use this account to query; if it is empty or not, use w+authentication account to query.", "zh_CN":"rbi映射账号。有传就使用此账号查询；为空或者没传，则使用w+鉴权账号查询。"}
  Relateaccount *string `json:"relateaccount,omitempty" xml:"relateaccount,omitempty"`
  // {"en":"Specifies the query date: 1.With format yyyy-mm-dd. 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope. 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'. 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope. 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'. 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried: 1.If there are multiple inputs,use ';' as separator. 2.If not specified, it means all the domains of the account .When datatype=0, when querying the average number of domain names, this parameter is invalid.", "zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道。当datatype=0，查询域名数均值的时候，此参数无效。"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Greenwich Mean Time Zone, the parameter format GMT+09:00 means East 9th District, GMT-09:00 means West 9th District, if not passed, it defaults to the local time zone (East 8th District). When datatype=0, when querying the average number of domain names, this parameter is invalid.", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）。当datatype=0，查询域名数均值的时候，此参数无效。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"The response format: 1.optional values:xml, json. 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"3 types. 0: The average number of query domain names; 1: Query bandwidth; 2: The number of query requests. Default query: 0.", "zh_CN":"3种类型。0：查询域名数均值；1：查询带宽；2：查询请求数。默认查询：0。"}
  Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
}

func (s WsiInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoRequest) GoString() string {
  return s.String()
}

func (s *WsiInfoRequest) SetCust(v string) *WsiInfoRequest {
  s.Cust = &v
  return s
}

func (s *WsiInfoRequest) SetRelateaccount(v string) *WsiInfoRequest {
  s.Relateaccount = &v
  return s
}

func (s *WsiInfoRequest) SetDate(v string) *WsiInfoRequest {
  s.Date = &v
  return s
}

func (s *WsiInfoRequest) SetStartdate(v string) *WsiInfoRequest {
  s.Startdate = &v
  return s
}

func (s *WsiInfoRequest) SetEnddate(v string) *WsiInfoRequest {
  s.Enddate = &v
  return s
}

func (s *WsiInfoRequest) SetChannel(v string) *WsiInfoRequest {
  s.Channel = &v
  return s
}

func (s *WsiInfoRequest) SetTimezone(v string) *WsiInfoRequest {
  s.Timezone = &v
  return s
}

func (s *WsiInfoRequest) SetDataformat(v string) *WsiInfoRequest {
  s.Dataformat = &v
  return s
}

func (s *WsiInfoRequest) SetDatatype(v string) *WsiInfoRequest {
  s.Datatype = &v
  return s
}

type WsiInfoResponse struct {
  // {"en":"peakTime", "zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"peakvalue(Mbps)", "zh_CN":"带宽峰值（单位Mbps）"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"chargeMethod", "zh_CN":"计费方式"}
  ChargeMethod *string `json:"chargeMethod,omitempty" xml:"chargeMethod,omitempty" require:"true"`
  // {"en":"chargeValue", "zh_CN":"计费值"}
  ChargeValue *string `json:"chargeValue,omitempty" xml:"chargeValue,omitempty" require:"true"`
  // {"en":"the total flow(GB)", "zh_CN":"总流量（单位GB）"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"the total hits", "zh_CN":"总请求数"}
  TotalHit *string `json:"totalHit,omitempty" xml:"totalHit,omitempty" require:"true"`
  // {"en":"Average number of domain names", "zh_CN":"域名数均值"}
  ChannelAvg *string `json:"channelAvg,omitempty" xml:"channelAvg,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'channel'}
  Channel *WsiInfoResponseChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s WsiInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoResponse) GoString() string {
  return s.String()
}

func (s *WsiInfoResponse) SetPeakTime(v string) *WsiInfoResponse {
  s.PeakTime = &v
  return s
}

func (s *WsiInfoResponse) SetPeakValue(v string) *WsiInfoResponse {
  s.PeakValue = &v
  return s
}

func (s *WsiInfoResponse) SetChargeMethod(v string) *WsiInfoResponse {
  s.ChargeMethod = &v
  return s
}

func (s *WsiInfoResponse) SetChargeValue(v string) *WsiInfoResponse {
  s.ChargeValue = &v
  return s
}

func (s *WsiInfoResponse) SetTotalFlow(v string) *WsiInfoResponse {
  s.TotalFlow = &v
  return s
}

func (s *WsiInfoResponse) SetTotalHit(v string) *WsiInfoResponse {
  s.TotalHit = &v
  return s
}

func (s *WsiInfoResponse) SetChannelAvg(v string) *WsiInfoResponse {
  s.ChannelAvg = &v
  return s
}

func (s *WsiInfoResponse) SetChannel(v *WsiInfoResponseChannel) *WsiInfoResponse {
  s.Channel = v
  return s
}

type WsiInfoResponseChannel struct {
  // {'en':'detail', 'zh_CN':'detail'}
  Detail []*WsiInfoResponseChannelDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
}

func (s WsiInfoResponseChannel) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoResponseChannel) GoString() string {
  return s.String()
}

func (s *WsiInfoResponseChannel) SetDetail(v []*WsiInfoResponseChannelDetail) *WsiInfoResponseChannel {
  s.Detail = v
  return s
}

type WsiInfoResponseChannelDetail struct     {
  // {"en":"time", "zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"The bandwidth value corresponding to the time point, the default unit is Mbps,
  // Return the value of the corresponding unit according to flowUnit", "zh_CN":"时间点对应的带宽值，默认单位Mbps，根据flowUnit返回对应单位的数值"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s WsiInfoResponseChannelDetail) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoResponseChannelDetail) GoString() string {
  return s.String()
}

func (s *WsiInfoResponseChannelDetail) SetTime(v string) *WsiInfoResponseChannelDetail {
  s.Time = &v
  return s
}

func (s *WsiInfoResponseChannelDetail) SetText(v string) *WsiInfoResponseChannelDetail {
  s.Text = &v
  return s
}

type WsiInfoPaths struct {
}

func (s WsiInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoPaths) GoString() string {
  return s.String()
}

type WsiInfoParameters struct {
}

func (s WsiInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoParameters) GoString() string {
  return s.String()
}

type WsiInfoRequestHeader struct {
}

func (s WsiInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoRequestHeader) GoString() string {
  return s.String()
}

type WsiInfoResponseHeader struct {
}

func (s WsiInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s WsiInfoResponseHeader) GoString() string {
  return s.String()
}




type ReportBandwidthWildcardDomainServiceRequest struct {
  // {"en":"Start time: 
  // 		1. The format is yyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00:00+08:00 (10:00:00 Beijing time on December 2, 2016); 
  // 		2. Can not exceed the current time; 
  // 		3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月02日10点0分0秒)
  //         2.不能大于当前时间
  //         3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 		1. The format is yyyy-MM-ddTHH:mm:ss+08:00; 
  // 		2. The end time is greater than the start time. If the end time is greater than the current time, the current time is taken. 
  // 		3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception; 
  // 		4. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days.", "zh_CN":"结束时间:
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  //         2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  //         3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  //         4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Wildcard Domain:
  //         1. The upper limit of the number of Wildcard Domain names that can be passed is 20 by default (contact technical support for adjustment);
  //         2. Automatically filter out invalid domain names (if a non-common domain name is passed in, it will be filtered out).
  //         3. If it is not filled, all pan domain names under the account will be queried by default", "zh_CN":"泛域名:
  //         1. 可传递泛域名数量上限默认为20个(可联系技术支持调整);
  //         2. 自动过滤掉无效域名(如传递非泛域名,会被过滤掉)。
  //         3. 如未填,则默认查询此账号下所有泛域名"}
  WildcardDomain []*string `json:"wildcardDomain,omitempty" xml:"wildcardDomain,omitempty" type:"Repeated"`
}

func (s ReportBandwidthWildcardDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportBandwidthWildcardDomainServiceRequest) SetDateFrom(v string) *ReportBandwidthWildcardDomainServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceRequest) SetDateTo(v string) *ReportBandwidthWildcardDomainServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceRequest) SetWildcardDomain(v []*string) *ReportBandwidthWildcardDomainServiceRequest {
  s.WildcardDomain = v
  return s
}

type ReportBandwidthWildcardDomainServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportBandwidthWildcardDomainServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthWildcardDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportBandwidthWildcardDomainServiceResponse) SetCode(v string) *ReportBandwidthWildcardDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceResponse) SetMessage(v string) *ReportBandwidthWildcardDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceResponse) SetData(v []*ReportBandwidthWildcardDomainServiceResponseData) *ReportBandwidthWildcardDomainServiceResponse {
  s.Data = v
  return s
}

type ReportBandwidthWildcardDomainServiceResponseData struct     {
  // {"en":"wildcard Domain", "zh_CN":"泛域名"}
  WildcardDomain *string `json:"wildcardDomain,omitempty" xml:"wildcardDomain,omitempty" require:"true"`
  DomainList []*ReportBandwidthWildcardDomainServiceResponseDataDomainList `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthWildcardDomainServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthWildcardDomainServiceResponseData) SetWildcardDomain(v string) *ReportBandwidthWildcardDomainServiceResponseData {
  s.WildcardDomain = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceResponseData) SetDomainList(v []*ReportBandwidthWildcardDomainServiceResponseDataDomainList) *ReportBandwidthWildcardDomainServiceResponseData {
  s.DomainList = v
  return s
}

type ReportBandwidthWildcardDomainServiceResponseDataDomainList struct     {
  // {"en":"domain", "zh_CN":"明细域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  BandwidthList []*ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList `json:"bandwidthList,omitempty" xml:"bandwidthList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthWildcardDomainServiceResponseDataDomainList) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceResponseDataDomainList) GoString() string {
  return s.String()
}

func (s *ReportBandwidthWildcardDomainServiceResponseDataDomainList) SetDomain(v string) *ReportBandwidthWildcardDomainServiceResponseDataDomainList {
  s.Domain = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceResponseDataDomainList) SetBandwidthList(v []*ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList) *ReportBandwidthWildcardDomainServiceResponseDataDomainList {
  s.BandwidthList = v
  return s
}

type ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList struct     {
  // {"en":"Time slice, returns the time slice containing the start time and end time. Time format: yyyy-MM-dd HH:mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。时间格式:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth value, unit Mbps, keep 2 decimal places", "zh_CN":"带宽值,单位Mbps,保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList) GoString() string {
  return s.String()
}

func (s *ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList) SetTimestamp(v string) *ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList {
  s.Timestamp = &v
  return s
}

func (s *ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList) SetValue(v string) *ReportBandwidthWildcardDomainServiceResponseDataDomainListBandwidthList {
  s.Value = &v
  return s
}

type ReportBandwidthWildcardDomainServicePaths struct {
}

func (s ReportBandwidthWildcardDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServicePaths) GoString() string {
  return s.String()
}

type ReportBandwidthWildcardDomainServiceParameters struct {
}

func (s ReportBandwidthWildcardDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceParameters) GoString() string {
  return s.String()
}

type ReportBandwidthWildcardDomainServiceRequestHeader struct {
}

func (s ReportBandwidthWildcardDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportBandwidthWildcardDomainServiceResponseHeader struct {
}

func (s ReportBandwidthWildcardDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthWildcardDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainBandwidthRequest struct {
  // {"en":"Domain list.
  // Domain number limits can be adjusted depending on different accounts. The default value is 1000(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表
  // 1.域名个数限制根据账号可调,默认为1000个（可联系技术支持下单调整）;"}
  QueryDomainBandwidthDomainList *QueryDomainBandwidthDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryDomainBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthRequest) SetDomainList(v *QueryDomainBandwidthDomainList) *QueryDomainBandwidthRequest {
  s.QueryDomainBandwidthDomainList = v
  return s
}

type QueryDomainBandwidthDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainBandwidthDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthDomainList) SetDomainName(v []*string) *QueryDomainBandwidthDomainList {
  s.DomainName = v
  return s
}

type QueryDomainBandwidthResponse struct {
  BandwidthReport []*QueryDomainBandwidthResponseBandwidthReport `json:"bandwidthReport,omitempty" xml:"bandwidthReport,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthResponse) SetBandwidthReport(v []*QueryDomainBandwidthResponseBandwidthReport) *QueryDomainBandwidthResponse {
  s.BandwidthReport = v
  return s
}

type QueryDomainBandwidthResponseBandwidthReport struct     {
  // {"en":"Date
  // When the querying data granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00;When the querying data granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is yyyy-MM-dd 24;When the querying data granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data;Return the time slice contained in start time and in end time", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth, keep the number to four decimal places. Unit: Mbps", "zh_CN":"带宽,保留4位小数,单位为Mbps"}
  Bandwidth *int `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s QueryDomainBandwidthResponseBandwidthReport) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthResponseBandwidthReport) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthResponseBandwidthReport) SetTimestamp(v string) *QueryDomainBandwidthResponseBandwidthReport {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainBandwidthResponseBandwidthReport) SetBandwidth(v int) *QueryDomainBandwidthResponseBandwidthReport {
  s.Bandwidth = &v
  return s
}

type QueryDomainBandwidthPaths struct {
}

func (s QueryDomainBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthPaths) GoString() string {
  return s.String()
}

type QueryDomainBandwidthParameters struct {
  // {"en":"Start time
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.And smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天（可联系技术支持调整）;4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time will be assigned as the value", "zh_CN":"结束时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Data granularity
  // fiveminutes: five minutes, hourly: one hour, daily: one day;If not specified, daily is set as the default value", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDomainBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthParameters) SetDateFrom(v string) *QueryDomainBandwidthParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainBandwidthParameters) SetDateTo(v string) *QueryDomainBandwidthParameters {
  s.DateTo = &v
  return s
}

func (s *QueryDomainBandwidthParameters) SetType(v string) *QueryDomainBandwidthParameters {
  s.Type = &v
  return s
}

type QueryDomainBandwidthRequestHeader struct {
}

func (s QueryDomainBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainBandwidthResponseHeader struct {
}

func (s QueryDomainBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthResponseHeader) GoString() string {
  return s.String()
}




