package reportflow

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest struct {
  // {"en":"Start time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. No bigger than the current time;\n3. Data in the last 183 days at most can be queried.","zh_CN":"开始时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.不能大于当前时间;\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. End time should be greater than start time. If the end time is greater than current time, current time will be used;\n3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default, if only one field is filled in and one is left empty, then exception will be occur;\n4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days).","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间;\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时，如仅有一个未传，抛异常;\n4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大不超过31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:\n1. The default upper limit to domains that can be entered is 1000 (Contact technical support to adjust);\n2. All domains under the account are queried if this input parameter is not specified, but if the number of domains under the account exceeds limits, no query will be done (Error).","zh_CN":"域名：\n1.可传递域名数量上限默认1000个（可联系技术支持调整）;\n2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Group keywords:\n1. By default, group data will be displayed;\n2. If there are keywords entered, value details shall be displayed by keywords;\nIf domain is specified to groupBy, it means results are returned according to domains;\n3. Only domain can be specified.","zh_CN":"分组关键词：\n1.默认聚合展示；\n2.传入关键词则代表需要按照关键词对应的值展示明细；\n例如groupBy传domain，则代表返回按照domain明细展开。\n3.只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetDateFrom(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetDateTo(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetDomain(v []*string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetGroupBy(v []*string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequestHeader struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsPaths struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsParameters struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse struct {
  // {"en":"","zh_CN":""}
  Result []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse) SetResult(v []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse {
  s.Result = v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic, unit: MB, retain two decimals, example (74099.92)","zh_CN":"总流量，单位:MB ，保留2位小数，示例 ( 74099.92 )"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Total requests","zh_CN":"总请求数"}
  TotalRequests *string `json:"totalRequests,omitempty" xml:"totalRequests,omitempty" require:"true"`
  // {"en":"Peak bandwidth, unit: Mbps, retain two decimals, example (74099.92)","zh_CN":"峰值带宽，单位: Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"Peak time of bandwidth, example (2019-02-13 18:01)","zh_CN":"峰值时间，示例（2019-02-13 18:01）"}
  PeakBandwidthTimestamp *string `json:"peakBandwidthTimestamp,omitempty" xml:"peakBandwidthTimestamp,omitempty" require:"true"`
  // {"en":"Peak requests","zh_CN":"请求数峰值"}
  PeakRequests *string `json:"peakRequests,omitempty" xml:"peakRequests,omitempty" require:"true"`
  // {"en":"Peak time of requests","zh_CN":"请求数峰值时间"}
  PeakRequestsTimestamp *string `json:"peakRequestsTimestamp,omitempty" xml:"peakRequestsTimestamp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetDomain(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetTotalTraffic(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.TotalTraffic = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetTotalRequests(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.TotalRequests = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakBandwidth(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakBandwidthTimestamp(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakBandwidthTimestamp = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakRequests(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakRequests = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakRequestsTimestamp(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakRequestsTimestamp = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetDataSeries(v []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.DataSeries = v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries struct     {
  // {"en":"Timestamp1. When the data granularity of query is 5m, the format is yyyy-mm-dd HH: MM;Each time slice data value represents the data value in the previous time granularity range.The time slice at the beginning of the day is yyyy-mm-dd 00:05, and the last time slice is (yyyy-mm-dd +1) 00:00;2. Return the time slice of start time and end time.","zh_CN":"时间片\n1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00;\n2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit: MB, retain two decimals,example (34099.92)","zh_CN":"流量值：\n\n单位MB，保留2位小数，示例 (34099.92)"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth, unit: Mbps, retain two decimals, example (931556.21)","zh_CN":"带宽值，单位: Mbps，保留2位小数，示例 （931556.21）"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Requests","zh_CN":"请求数"}
  Requests *string `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetTimestamp(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetTraffic(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetBandwidth(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Bandwidth = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetRequests(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Requests = &v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseHeader struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetOriginTrafficAndRequestsForMultiDomainsRequest struct {
  // {"en":"Start time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. Not greater than the current time;\n3. The data can be queried is the last week (183 days).","zh_CN":"开始时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.不能大于当前时间；\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. End time should be greater than start time. If the end time is greater than current time, current time will be used.\n3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if one field is filled and one is left empty, then exception will occur.\n4. Maximum query time interval allowed: 1 day by default, that is, the difference between dateFrom and dateTo cannot exceed 1 day (you can contact technical support to adjust it, the maximum adjustment is 31 days)","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常\n4.允许查询最大时间间隔：默认1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大调整到31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"defaultValue":"5m","en":"Data granularity:\n1. Support 5m (granularity of 5 minutes) and 1m","zh_CN":"数据粒度：\n1、支持5m（5分钟）、1m（1分钟）","exampleValue":"1m,5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Domain name:\n1. The default upper limit of domains that can be entered is 200 (if you want to adjust, please, contact technical support);\n2. All domains under the account will be queried if this input parameter is not specified. But if the number of domains under the account exceeds the limits, no query will be executed (Error)","zh_CN":"域名：\n1、可传递域名数量上限默认为200个（可联系技术支持调整）；\n2、未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Grouping keywords:\n1. By default, data will be displayed by group;\n2. If there are keywords entered, value details shall be displayed by keywords; If groupBy is specified as domain, it means the results are returned according to domains.\n3. Only domain can be specified","zh_CN":"分组关键词：\n1、默认聚合展示；\n2、传入关键词则代表需要按照关键词对应的值展示明细；\n例如groupBy传domain，则代表返回按照domain明细展开。\n3、只能传递domain。","exampleValue":"domain"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"defaultValue":"0","en":"Input 0 returns all data, input 1 only returns source site data","zh_CN":"入参 0 则返回全部数据，入参 1 则只返回回源站数据","exampleValue":"0,1"}
  OriginOnly *int `json:"originOnly,omitempty" xml:"originOnly,omitempty"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetDateFrom(v string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetDateTo(v string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetGranularity(v string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.Granularity = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetDomain(v []*string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetGroupBy(v []*string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetOriginOnly(v int) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.OriginOnly = &v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsRequestHeader struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetOriginTrafficAndRequestsForMultiDomainsPaths struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetOriginTrafficAndRequestsForMultiDomainsParameters struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetOriginTrafficAndRequestsForMultiDomainsResponse struct {
  // {"en":"","zh_CN":""}
  Result []*GetOriginTrafficAndRequestsForMultiDomainsResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponse) SetResult(v []*GetOriginTrafficAndRequestsForMultiDomainsResponseResult) *GetOriginTrafficAndRequestsForMultiDomainsResponse {
  s.Result = v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic, Unit:  MB, retain two decimals, example (74099.92)","zh_CN":"总流量，单位MB，保留2位小数，示例：(74099.92)"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Total requests","zh_CN":"总请求数"}
  TotalRequests *string `json:"totalRequests,omitempty" xml:"totalRequests,omitempty" require:"true"`
  // {"en":"Peak of Requests","zh_CN":"请求数峰值"}
  PeakRequests *string `json:"peakRequests,omitempty" xml:"peakRequests,omitempty" require:"true"`
  // {"en":"Peak time of Request","zh_CN":"请求数峰值时间"}
  PeakRequestsTimestamp *string `json:"peakRequestsTimestamp,omitempty" xml:"peakRequestsTimestamp,omitempty" require:"true"`
  // {"en":"Peak bandwidth, Unit: Mbps, retain two decimals, example (74099.92)","zh_CN":"带宽峰值，单位Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"Peak time of Bandwidth","zh_CN":"带宽峰值时间"}
  PeakBandwidthTimestamp *string `json:"peakBandwidthTimestamp,omitempty" xml:"peakBandwidthTimestamp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResult) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetDomain(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.Domain = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetTotalTraffic(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.TotalTraffic = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetTotalRequests(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.TotalRequests = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakRequests(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakRequests = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakRequestsTimestamp(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakRequestsTimestamp = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakBandwidth(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakBandwidthTimestamp(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakBandwidthTimestamp = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetDataSeries(v []*GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.DataSeries = v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries struct     {
  // {"en":"1. When the querying data granularity is 5m, then the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.\n2. Return the time slices that contained in start time and in end time.","zh_CN":"1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1）00:00。\n2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit:  MB, retain two decimals, example (34099.92)","zh_CN":"流量值：\n\n单位MB，保留2位小数，示例：(34099.92)"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth value:\n\nUnit: Mbps, retain two decimals, example: (31556.21)","zh_CN":"带宽值：\n\n单位Mbps，保留2位小数, 示例：(31556.21)"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Total number of requests","zh_CN":"请求数"}
  Requests *string `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetTimestamp(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetTraffic(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetBandwidth(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Bandwidth = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetRequests(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Requests = &v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsResponseHeader struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest struct {
  // {"en":"Start Time:\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, e.g., +00:00 represents UTC time, +08:00 represents East 8th District. For example, 2024-01-15T10:30:45+00:00 indicates January 15, 2024, 10:30:45 AM UTC time.\n2.Cannot be later than the current time.\n3.Data for up to the past six months (183 days) can be retrieved.","zh_CN":"开始时间：\n\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n\n2.不能大于当前时间\n\n3.最多可获取最近半年（183天）的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM.\n2. The end time must be greater than the start time.\n3. If the end time is set later than the current time, the current time will be used instead.\n4. The maximum allowed query interval is 1 day (i.e., the difference between `dateFrom` and `dateTo` cannot exceed 1 day). This limit can be adjusted by contacting technical support.","zh_CN":"结束时间： 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.结束时间需大于开始时间  \n3.结束时间如果大于当前时间，取当前时间  \n4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:\n1. The maximum number of domains allowed is 20 (can be adjusted by contacting technical support).\n2. Invalid domain names (e.g., illegal domain names passed in the request) will be automatically filtered out, and query results will only return data for valid domain names.","zh_CN":"域名:\n1、可传递域名数量上限默认为20个(可联系技术支持调整);\n2、自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"5m","en":"Data granularity: \n1. Supports 1m (1 minute), 5m (5 minutes), 1h (1 hour).\n2. Defaults to 5m if not specified.","zh_CN":"数据粒度:\n1、支持1m(1分钟)、5m(5分钟)、1h(1小时)\n2、不传默认5m","exampleValue":"1m,5m,1h"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"defaultValue":"bandwidth","en":"Query data types:\nbandwidth and request","zh_CN":"查询数据类型：\n带宽(bandwidth)，请求数(request)","exampleValue":"bandwidth,request"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"If `groupBy` is left empty, data for all domains will be aggregated.","zh_CN":"不传默认聚合所有频道数据","exampleValue":"domain"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) SetDateFrom(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) SetDateTo(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) SetDomain(v []*string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) SetGranularity(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest {
  s.Granularity = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) SetQueryBy(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest {
  s.QueryBy = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest) SetGroupBy(v []*string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsRequestHeader struct {
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsPaths struct {
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsParameters struct {
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"List of response data","zh_CN":"响应数据列表"}
  Data []*GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse) SetCode(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse {
  s.Code = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse) SetMessage(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse {
  s.Message = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse) SetData(v []*GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponse {
  s.Data = v
  return s
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData struct     {
  // {"en":"Domain. This field is not returned when aggregating all domain data.","zh_CN":"域名, 聚合全部域名数据时不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Time-series data list.","zh_CN":"时间序列数据列表"}
  DataSeries []*GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData) SetDomain(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData {
  s.Domain = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData) SetDataSeries(v []*GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseData {
  s.DataSeries = v
  return s
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries struct     {
  // {"en":"Timestamp. Returns the timestamps within the specified start and end times. Time format: 1/5 minutes: yyyy-MM-dd HH:mm; 1 hour: yyyy-MM-dd HH.","zh_CN":"时间片, 返回开始时间和结束时间包含的时间片。时间格式:1/5分钟:yyyy-MM-dd HH:mm, 1小时:yyyy-MM-dd HH"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"IPv4 data (bandwidth unit: Mbps, number of requests unit: count)","zh_CN":"IPv4数据（带宽单位 Mbps，请求数单位 个）"}
  V4Value *string `json:"v4Value,omitempty" xml:"v4Value,omitempty" require:"true"`
  // {"en":"IPv6 data (bandwidth unit: Mbps, number of requests unit: count)","zh_CN":"IPv6数据（带宽单位 Mbps，请求数单位 个）"}
  V6Value *string `json:"v6Value,omitempty" xml:"v6Value,omitempty" require:"true"`
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries) SetTimestamp(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries) SetV4Value(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries {
  s.V4Value = &v
  return s
}

func (s *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries) SetV6Value(v string) *GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseDataDataSeries {
  s.V6Value = &v
  return s
}

type GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseHeader struct {
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndRequestsByIpVersionForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficHitRatioByIspProvinceForMultiDomainsRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.The data can be obtained for the latest six months (183 days) at most, and the start time cannot be greater than the current time and end time.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n2.最多可获取最近半年（183天）的数据，并且开始时间不能大于当前时间和结束时间"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n2. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception\n3. The default query span is up to 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (you can contact technical support for adjustment, up to 31 days);","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常  \n3.默认允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整，最大31天）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number   limits can be adjusted depending on different accounts. The default value is 20","zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"5m","en":"Data granularity, 5m: 5-minute granularity, 1h: 1-hour granularity","zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度","exampleValue":"5m,1h"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province\n1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;\n2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"省份\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;\n2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。\n2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension\n1.Options are domain, province, isp, and more than one value can be entered;\n2.The data is displayed according to the specified dimension;","zh_CN":"分组维度\n1.可选值为domain、province、isp,可传入多个值;\n2.有传入则按照该维度展示明细数据;","exampleValue":"domain,province,isp"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetDateFrom(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetDateTo(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetDomain(v []*string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetGranularity(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetProvince(v []*string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.Province = v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetIsp(v []*string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest) SetGroupBy(v []*string) *GetTrafficHitRatioByIspProvinceForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsRequestHeader struct {
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsPaths struct {
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsParameters struct {
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsResponse struct {
  // {"en":"","zh_CN":""}
  Result []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponse) SetResult(v []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponse {
  s.Result = v
  return s
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  IspData []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult) SetDomain(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult) SetIspData(v []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResult {
  s.IspData = v
  return s
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData struct     {
  // {"en":"ISP","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  ProvinceData []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData) GoString() string {
  return s.String()
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData) SetIsp(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData) SetProvinceData(v []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspData {
  s.ProvinceData = v
  return s
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData struct     {
  // {"en":"Province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) SetProvince(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) SetDataSeries(v []*GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceData {
  s.DataSeries = v
  return s
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries struct     {
  // {"en":"Time,\n1.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;\n2.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;\n3.Return the time slice contained in start time and the time slice contained in end time.","zh_CN":"时间,\n1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;\n2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;\n3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Byte hit rate. retain four decimals","zh_CN":"字节命中率,保留4位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
  // {"en":"Hit traffic. \nUnit is MB ,retain two decimals","zh_CN":"命中流量值,单位MB,保留2位小数;\n默认不返回该字段,需要返回的请联系技术支持"}
  HitTraffic *string `json:"hitTraffic,omitempty" xml:"hitTraffic,omitempty" require:"true"`
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) SetTimestamp(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) SetHitRate(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries {
  s.HitRate = &v
  return s
}

func (s *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) SetHitTraffic(v string) *GetTrafficHitRatioByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries {
  s.HitTraffic = &v
  return s
}

type GetTrafficHitRatioByIspProvinceForMultiDomainsResponseHeader struct {
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficHitRatioByIspProvinceForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;\n3.Period between dataFrom and dateTo cannot be longer than 24 hours,if the dataInterval is 1m,period between dataFrom and dateTo cannot be longer than 6 hours(you can contact technical support to adjust it);\n4.dateFrom and dateTo can be either both are specified or neither is specifies;\n5.If neither dateFrom nor dateTo is specified, then by default, data in the last 30 minutes is queried","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.必须大于当前时间-183天,并且小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过24小时;一分钟粒度dateFrom和dateTo相差不能超过6小时(可联系技术支持调整)\n4.dateFrom和dateTo要么都传递,要么都不传递;\n5.dateFrom和dateTo都未传递,则默认查询过去30分钟的数据;"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain\n1.domain number limits can be adjusted depending on different accounts. The default value is 20.\n2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)","zh_CN":"域名:\n1.可传递域名数量上限默认为20个(可联系技术支持调整);\n2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)\n3.域名超过上限,报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:\n1. Support 1m (1 minute granularity),5m (5 minutes granularity)\n2. The default value is 5m","zh_CN":"数据粒度:\n1.1m:1分钟粒度, 5m:5分钟粒度\n2.不传默认查询 5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Province\n\n1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;\n2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.\n\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"省份\n\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n\n2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节\n\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:\n1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;\n2.ISP is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。\n2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Query Dimension\n1.Optional values: flow, request\n2.The default value is flow\n3.Flow: flow,Unit MB, two decimal places reserved;\n4.Request: Number of requests","zh_CN":"查询维度\n1.可选值 flow、request\n2.不传默认 flow\n3.flow:流量,单位 MB,保留两位小数;\n4.request:请求数"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional values: \n1.domain,province,isp. Multiple values can be transferred in;\n2.Display detailed data according to this dimension if it is transferred in","zh_CN":"分组维度:\n1.可选值为domain、province、isp,可传入多个值;\n2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetDateFrom(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetDateTo(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetDomain(v []*string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetGranularity(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetProvince(v []*string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.Province = v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetIsp(v []*string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetQueryBy(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.QueryBy = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest) SetGroupBy(v []*string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequest {
  s.GroupBy = v
  return s
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequestHeader struct {
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityPaths struct {
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityPaths) GoString() string {
  return s.String()
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityParameters struct {
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityParameters) GoString() string {
  return s.String()
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result infotmation","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detail data of request result","zh_CN":"请求结果的详细数据"}
  Data []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse) SetCode(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse {
  s.Code = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse) SetMessage(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse {
  s.Message = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse) SetData(v []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponse {
  s.Data = v
  return s
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ISP data","zh_CN":"运营商数据"}
  IspData []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData) SetDomain(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData {
  s.Domain = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData) SetIspData(v []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseData {
  s.IspData = v
  return s
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData struct     {
  // {"en":"ISP","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"Province data","zh_CN":"省份数据"}
  ProvinceData []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData) SetIsp(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData {
  s.Isp = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData) SetProvinceData(v []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspData {
  s.ProvinceData = v
  return s
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData struct     {
  // {"en":"Province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"-","zh_CN":""}
  DataSeries []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData) SetProvince(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData) SetDataSeries(v []*GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceData {
  s.DataSeries = v
  return s
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries struct     {
  // {"en":"Time,the format is yyyy-MM-dd HH:mm","zh_CN":"时间,格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Hit data incording to query by(flow/requests)","zh_CN":"命中数据 flow:流量,保留两位小数; request:请求数"}
  HitValue *string `json:"hitValue,omitempty" xml:"hitValue,omitempty" require:"true"`
  // {"en":"Hit rate,keep four decimal places","zh_CN":"命中率,保留四位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries) SetTimestamp(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries) SetHitValue(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries {
  s.HitValue = &v
  return s
}

func (s *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries) SetHitRate(v string) *GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseDataIspDataProvinceDataDataSeries {
  s.HitRate = &v
  return s
}

type GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseHeader struct {
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsHitRatioByClientIspProvince1minGranularityResponseHeader) GoString() string {
  return s.String()
}




type QueryOutputTrafficUnderShieldPoPRequest struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it&rsquo;s greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, support 5m (granularity of 5 minutes)", "zh_CN":"数据粒度,支持5m:5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPRequest) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDateFrom(v string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDateTo(v string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.DateTo = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDomain(v []*string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.Domain = v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDataInterval(v string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetGroupBy(v []*string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.GroupBy = v
  return s
}

type QueryOutputTrafficUnderShieldPoPResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryOutputTrafficUnderShieldPoPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponse) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponse) SetResult(v []*QueryOutputTrafficUnderShieldPoPResponseResult) *QueryOutputTrafficUnderShieldPoPResponse {
  s.Result = v
  return s
}

type QueryOutputTrafficUnderShieldPoPResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic", "zh_CN":"总流量"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量值数据"}
  FlowData []*QueryOutputTrafficUnderShieldPoPResponseResultFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResult) SetDomain(v string) *QueryOutputTrafficUnderShieldPoPResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResult) SetTotalFlow(v string) *QueryOutputTrafficUnderShieldPoPResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResult) SetFlowData(v []*QueryOutputTrafficUnderShieldPoPResponseResultFlowData) *QueryOutputTrafficUnderShieldPoPResponseResult {
  s.FlowData = v
  return s
}

type QueryOutputTrafficUnderShieldPoPResponseResultFlowData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits &nbsp; of decimals allowed", "zh_CN":"流量值,单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryOutputTrafficUnderShieldPoPResponseResultFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponseResultFlowData) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResultFlowData) SetTimestamp(v string) *QueryOutputTrafficUnderShieldPoPResponseResultFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResultFlowData) SetValue(v string) *QueryOutputTrafficUnderShieldPoPResponseResultFlowData {
  s.Value = &v
  return s
}

type QueryOutputTrafficUnderShieldPoPPaths struct {
}

func (s QueryOutputTrafficUnderShieldPoPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPPaths) GoString() string {
  return s.String()
}

type QueryOutputTrafficUnderShieldPoPParameters struct {
}

func (s QueryOutputTrafficUnderShieldPoPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPParameters) GoString() string {
  return s.String()
}

type QueryOutputTrafficUnderShieldPoPRequestHeader struct {
}

func (s QueryOutputTrafficUnderShieldPoPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPRequestHeader) GoString() string {
  return s.String()
}

type QueryOutputTrafficUnderShieldPoPResponseHeader struct {
}

func (s QueryOutputTrafficUnderShieldPoPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficByProtocolRequest struct {
  // {"en":"Start time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be no more than 183 days prior to the current time, and must be earlier than both the current time and `dateTo`.\n3.The period between `dateFrom` and `dateTo` cannot exceed 7 days (contact technical support for adjustments).\n4.Either both `dateFrom` and `dateTo` must be specified, or neither should be specified.\n5.If neither dateFrom nor dateTo is specified, then by default, data for the last 24 hours is queried.","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于当前时间-183天,并且小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);\n4.dateFrom和dateTo要么都传递,要么都不传递;\n5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM.\n2.Must be greater than dateFrom.\n3.If it is greater than the current time, it will be reset to the current time.","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于dateFrom;\n3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names. The number of allowed domains can be adjusted based on the account type, with a default limit of 20.","zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"5m","en":"**Data granularity**\n1m: 1minute\n5m: 5minutes\n1h: 1hour\n1d: 1day","zh_CN":"数据粒度：\n1m：1分钟粒度\n5m：5分钟粒度\n1h：1小时粒度\n1d：1天粒度","exampleValue":"1m,5m,1h,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"defaultValue":"https","en":"Transmission protocol:\n1.Options: http, https.\n2.If no value is specified, https is used as the default.\n3.If http is queried, `httpFlowData` is displayed in the response; if https is queried, `httpsFlowData` is displayed.","zh_CN":"传输协议\n1.可选值为http、https;\n2.不传默认查询https;\n3.查询http时出参展示httpFlowData,查询https时出参展示httpsFlowData;","exampleValue":"http,https"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group dimension:\n1.Only 'domain' is a valid value.\n2.If provided, detailed data will be displayed according to this dimension.","zh_CN":"分组维度\n1.可选值为domain;\n2.有传入则按照该维度展示明细数据;","exampleValue":"domain"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficByProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolRequest) SetDateFrom(v string) *GetTrafficByProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetDateTo(v string) *GetTrafficByProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetDomain(v []*string) *GetTrafficByProtocolRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficByProtocolRequest) SetGranularity(v string) *GetTrafficByProtocolRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetProtocolType(v string) *GetTrafficByProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetGroupBy(v []*string) *GetTrafficByProtocolRequest {
  s.GroupBy = v
  return s
}

type GetTrafficByProtocolRequestHeader struct {
}

func (s GetTrafficByProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficByProtocolPaths struct {
}

func (s GetTrafficByProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolPaths) GoString() string {
  return s.String()
}

type GetTrafficByProtocolParameters struct {
}

func (s GetTrafficByProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolParameters) GoString() string {
  return s.String()
}

type GetTrafficByProtocolResponse struct {
  // {"en":"The query result set.","zh_CN":"结果"}
  Result []*GetTrafficByProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolResponse) SetResult(v []*GetTrafficByProtocolResponseResult) *GetTrafficByProtocolResponse {
  s.Result = v
  return s
}

type GetTrafficByProtocolResponseResult struct     {
  // {"en":"The domain name.","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Time-series data list.","zh_CN":"时间序列数据列表"}
  DataSeries []*GetTrafficByProtocolResponseResultDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolResponseResult) SetDomain(v string) *GetTrafficByProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficByProtocolResponseResult) SetDataSeries(v []*GetTrafficByProtocolResponseResultDataSeries) *GetTrafficByProtocolResponseResult {
  s.DataSeries = v
  return s
}

type GetTrafficByProtocolResponseResultDataSeries struct     {
  // {"en":"Timestamp, in yyyy-MM-dd HH:mm format. Each timestamp's data value represents the aggregated data within the preceding time granularity interval. The first timestamp for a day is yyyy-MM-dd 00:05, and the last is (yyyy-MM-dd+1) 00:00.","zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit: MB. Retains two decimal places.","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s GetTrafficByProtocolResponseResultDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponseResultDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolResponseResultDataSeries) SetTimestamp(v string) *GetTrafficByProtocolResponseResultDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficByProtocolResponseResultDataSeries) SetTraffic(v string) *GetTrafficByProtocolResponseResultDataSeries {
  s.Traffic = &v
  return s
}

type GetTrafficByProtocolResponseHeader struct {
}

func (s GetTrafficByProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponseHeader) GoString() string {
  return s.String()
}




type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM.\n2.Must be smaller than the current time and dateTo.\n3.Period between dateFrom and dateTo cannot be longer than 31 days.\n4.You can only query data for the last 6 months.","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.必须小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过31天;\n4.只能查询最近半年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM.\n2.Must be greater than dateFrom.\n3.If it is greater than the current time, it will be reset to the current time.","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.必须大于dateFrom;\n3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Directory levels, value range 1-4. Only one value can be submitted.","zh_CN":"目录层级,取值范围1~4,只能提交单个值"}
  DirHierarchy *string `json:"dirHierarchy,omitempty" xml:"dirHierarchy,omitempty" require:"true"`
  // {"en":"Domain and directory details","zh_CN":"域名目录详情"}
  DomainDir []*GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir `json:"domainDir,omitempty" xml:"domainDir,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest) SetDateFrom(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest {
  s.DateFrom = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest) SetDateTo(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest {
  s.DateTo = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest) SetDirHierarchy(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest {
  s.DirHierarchy = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest) SetDomainDir(v []*GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequest {
  s.DomainDir = v
  return s
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir struct     {
  // {"en":"Domain\n1.Must comply with domain name regular expression validation rules.\n2.Domain number limits can be adjusted depending on different accounts. The default value is 1.","zh_CN":"域名\n1.需要满足域名的正则校验;\n2.域名个数限制根据账号可调,默认为1个;"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Directory\n1.Directory number limits can be adjusted depending on different accounts. The default value is 200.\n2.Empty value means to query all directories. The number of directories must not exceed the set limit.\n3.Invalid directories are not returned.","zh_CN":"目录\n1.目录个数限制根据账号可调,默认为200个;\n2.不传代表查询该域名下的所有目录,同时接受目录个数限制;\n3.无效的目录不返回"}
  Dir []*string `json:"dir,omitempty" xml:"dir,omitempty" type:"Repeated"`
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir) SetDomain(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir {
  s.Domain = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir) SetDir(v []*string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestDomainDir {
  s.Dir = v
  return s
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestHeader struct {
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyRequestHeader) GoString() string {
  return s.String()
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyPaths struct {
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyPaths) GoString() string {
  return s.String()
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyParameters struct {
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyParameters) GoString() string {
  return s.String()
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponse struct {
  // {"en":"List of query results","zh_CN":"查询结果列表"}
  Result []*GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponse) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponse) SetResult(v []*GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponse {
  s.Result = v
  return s
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Detailed data","zh_CN":"详情数据"}
  Details []*GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult) SetDomain(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult {
  s.Domain = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult) SetDetails(v []*GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResult {
  s.Details = v
  return s
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails struct     {
  // {"en":"Directory name of corresponding level","zh_CN":"对应层级的目录名称"}
  Dir *string `json:"dir,omitempty" xml:"dir,omitempty" require:"true"`
  // {"en":"Total traffic, unit MB, retained to two decimal places","zh_CN":"总流量,单位MB,保留2位小数"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Peak bandwidth, unit Mbps, retained to two decimal places","zh_CN":"带宽峰值,单位Mbps,保留2位小数"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails) GoString() string {
  return s.String()
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails) SetDir(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails {
  s.Dir = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails) SetTotalTraffic(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails {
  s.TotalTraffic = &v
  return s
}

func (s *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails) SetPeakBandwidth(v string) *GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseResultDetails {
  s.PeakBandwidth = &v
  return s
}

type GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseHeader struct {
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthAndTrafficByDirectoriesMainlandChinaOnlyResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowExactDomainServiceRequest struct {
  // {'en':'Starting time
  // 
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00) );
  // 
  // 2. Cannot be greater than the current time
  // 
  // 3. Get up to the last six months (183 days) of data.', 'zh_CN':'开始时间
  // 1.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;不能大于当前时间
  // 3.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End Time
  // 
  // 1. Time format yyyy-MM-ddTHH:mm:ss+08:00, for example: 2016-12-02T10:00:00+08:00
  // 
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 
  // 3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception
  // 
  // 4. Allow query maximum time interval: default 1 day, that is, the difference between dateFrom and dateTo can't exceed 1 day (can contact technical support adjustment).', 'zh_CN':'结束时间
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00，例如：2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：默认1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity:
  // 
  // 1. Support 5m (5 minutes granularity), 1d (1 day granularity)
  // 
  // 2. Do not pass the default to 5m', 'zh_CN':'数据粒度：
  // 1.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;支持5m（5分钟粒度）、1d（1天粒度）
  // 2.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;不传默认为5m'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Domain name:
  // 
  // 1. The maximum number of passable domain names is 20 by default (you can contact technical support adjustment).
  // 
  // 2. When the total number of precise domain names under multiple generic domain names exceeds the number of transmittables configured by exactDomain, it is not possible to check (error prompt).
  // 
  // 3. Automatically filter out illegal domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of legitimate domain names).', 'zh_CN':'域名（即泛域名）：
  // 1.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;可传递域名数量上限默认为20个（可联系技术支持调整）。
  // 2.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;当传入的多个泛域名下的精确域名总数量超过exactDomain配置的可传数量时不可查（报错提示）。
  // 3.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'Exact domain name
  // 
  // 1. The maximum number of passable precise domain names is 200 by default (can be contacted for technical support adjustment)
  // 
  // 2. Can be an exact domain name under different pan-domain names.', 'zh_CN':'精确域名
  // 1.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;可传递精确域名数量上限默认为200个（可联系技术支持调整）
  // 2.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;可以是不同泛域名下的精确域名。'}
  ExactDomain []*string `json:"exactDomain,omitempty" xml:"exactDomain,omitempty" type:"Repeated"`
}

func (s ReportFlowExactDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceRequest) SetDateFrom(v string) *ReportFlowExactDomainServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowExactDomainServiceRequest) SetDateTo(v string) *ReportFlowExactDomainServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowExactDomainServiceRequest) SetDataInterval(v string) *ReportFlowExactDomainServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowExactDomainServiceRequest) SetDomain(v []*string) *ReportFlowExactDomainServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowExactDomainServiceRequest) SetExactDomain(v []*string) *ReportFlowExactDomainServiceRequest {
  s.ExactDomain = v
  return s
}

type ReportFlowExactDomainServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*ReportFlowExactDomainServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowExactDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceResponse) SetResult(v []*ReportFlowExactDomainServiceResponseResult) *ReportFlowExactDomainServiceResponse {
  s.Result = v
  return s
}

type ReportFlowExactDomainServiceResponseResult struct     {
  // {'en':'Exact domain name', 'zh_CN':'精确域名'}
  ExactDomain *string `json:"exactDomain,omitempty" xml:"exactDomain,omitempty" require:"true"`
  FlowData []*ReportFlowExactDomainServiceResponseResultFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowExactDomainServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceResponseResult) SetExactDomain(v string) *ReportFlowExactDomainServiceResponseResult {
  s.ExactDomain = &v
  return s
}

func (s *ReportFlowExactDomainServiceResponseResult) SetFlowData(v []*ReportFlowExactDomainServiceResponseResultFlowData) *ReportFlowExactDomainServiceResponseResult {
  s.FlowData = v
  return s
}

type ReportFlowExactDomainServiceResponseResultFlowData struct     {
  // {'en':'time
  // 1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. When the data granularity of the query is 1d, the format is yyyy-MM-dd; the data value of the day represented by each time slice data value.
  // 3. Returns the time slice contained in the start time and end time.', 'zh_CN':'时间
  // 1.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;查询的数据粒度为5m时，格式为yyyy-MM-dd &nbsp; HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00。
  // 2.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;查询的数据粒度为1d时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值。
  // 3.&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Flow, unit of measure MB, retain 2 decimal places', 'zh_CN':'流量，计量单位MB，保留2位小数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowExactDomainServiceResponseResultFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceResponseResultFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceResponseResultFlowData) SetTimestamp(v string) *ReportFlowExactDomainServiceResponseResultFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowExactDomainServiceResponseResultFlowData) SetValue(v string) *ReportFlowExactDomainServiceResponseResultFlowData {
  s.Value = &v
  return s
}

type ReportFlowExactDomainServicePaths struct {
}

func (s ReportFlowExactDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServicePaths) GoString() string {
  return s.String()
}

type ReportFlowExactDomainServiceParameters struct {
}

func (s ReportFlowExactDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowExactDomainServiceRequestHeader struct {
}

func (s ReportFlowExactDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowExactDomainServiceResponseHeader struct {
}

func (s ReportFlowExactDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type GetOriginTrafficByIspProvinceRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.The data can be obtained for the latest six months (183 days) at most, and the start time cannot be greater than the current time and end time.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n2.最多可获取最近半年（183天）的数据，并且开始时间不能大于当前时间和结束时间"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n2. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception\n3. The default query span is up to 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (you can contact technical support for adjustment, up to 31 days);","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常  \n3.默认允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整，最大31天）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20","zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"5m","en":"Data granularity, 5m: 5-minute granularity, 1h: 1-hour granularity, The default granularity is 5 minutes.","zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度, 默认粒度5分钟","exampleValue":"5m,1h"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province:\n1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;\n2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"省份\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;\n2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。\n2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension\n1.Options are domain, province, isp, and more than one value can be entered;\n2.The data is displayed according to the specified dimension, if not transmitted, aggregate according to all dimensions.","zh_CN":"分组维度\n1.可选值为domain、province、isp,可传入多个值;\n2.有传入则按照该维度展示明细数据, 没传则按照所有维度聚合;","exampleValue":"domain,province,isp"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetOriginTrafficByIspProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceRequest) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficByIspProvinceRequest) SetDateFrom(v string) *GetOriginTrafficByIspProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceRequest) SetDateTo(v string) *GetOriginTrafficByIspProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceRequest) SetDomain(v []*string) *GetOriginTrafficByIspProvinceRequest {
  s.Domain = v
  return s
}

func (s *GetOriginTrafficByIspProvinceRequest) SetGranularity(v string) *GetOriginTrafficByIspProvinceRequest {
  s.Granularity = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceRequest) SetProvince(v []*string) *GetOriginTrafficByIspProvinceRequest {
  s.Province = v
  return s
}

func (s *GetOriginTrafficByIspProvinceRequest) SetIsp(v []*string) *GetOriginTrafficByIspProvinceRequest {
  s.Isp = v
  return s
}

func (s *GetOriginTrafficByIspProvinceRequest) SetGroupBy(v []*string) *GetOriginTrafficByIspProvinceRequest {
  s.GroupBy = v
  return s
}

type GetOriginTrafficByIspProvinceRequestHeader struct {
}

func (s GetOriginTrafficByIspProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceRequestHeader) GoString() string {
  return s.String()
}

type GetOriginTrafficByIspProvincePaths struct {
}

func (s GetOriginTrafficByIspProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvincePaths) GoString() string {
  return s.String()
}

type GetOriginTrafficByIspProvinceParameters struct {
}

func (s GetOriginTrafficByIspProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceParameters) GoString() string {
  return s.String()
}

type GetOriginTrafficByIspProvinceResponse struct {
  // {"en":"Result","zh_CN":"结果"}
  Result []*GetOriginTrafficByIspProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficByIspProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceResponse) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficByIspProvinceResponse) SetResult(v []*GetOriginTrafficByIspProvinceResponseResult) *GetOriginTrafficByIspProvinceResponse {
  s.Result = v
  return s
}

type GetOriginTrafficByIspProvinceResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Isp data","zh_CN":"ISP数据"}
  IspData []*GetOriginTrafficByIspProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficByIspProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficByIspProvinceResponseResult) SetDomain(v string) *GetOriginTrafficByIspProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceResponseResult) SetIspData(v []*GetOriginTrafficByIspProvinceResponseResultIspData) *GetOriginTrafficByIspProvinceResponseResult {
  s.IspData = v
  return s
}

type GetOriginTrafficByIspProvinceResponseResultIspData struct     {
  // {"en":"Internet service providers","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"Province data","zh_CN":"省份数据"}
  ProvinceData []*GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficByIspProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspData) SetIsp(v string) *GetOriginTrafficByIspProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspData) SetProvinceData(v []*GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData) *GetOriginTrafficByIspProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData) SetProvince(v string) *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData) SetDataSeries(v []*GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries) *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceData {
  s.DataSeries = v
  return s
}

type GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries struct     {
  // {"en":"Time,\n\n1.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;\n\n2.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;\n\n3.Return the time slice contained in start time and the time slice contained in end time.","zh_CN":"时间,\n1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;\n2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;\n3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Origin traffic. Unit: MB, retain two decimals","zh_CN":"回源流量,保留2位小数,单位MB"}
  OriginTraffic *string `json:"originTraffic,omitempty" xml:"originTraffic,omitempty" require:"true"`
  // {"en":"Average origin traffic.  Unit: Mbps, retain two decimals","zh_CN":"回源平均带宽,保留2位小数,单位Mbps"}
  OriginBandwidth *string `json:"originBandwidth,omitempty" xml:"originBandwidth,omitempty" require:"true"`
}

func (s GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries) SetTimestamp(v string) *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries) SetOriginTraffic(v string) *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries {
  s.OriginTraffic = &v
  return s
}

func (s *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries) SetOriginBandwidth(v string) *GetOriginTrafficByIspProvinceResponseResultIspDataProvinceDataDataSeries {
  s.OriginBandwidth = &v
  return s
}

type GetOriginTrafficByIspProvinceResponseHeader struct {
}

func (s GetOriginTrafficByIspProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficByIspProvinceResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.The data can be obtained for the latest six months (183 days) at most, and the start time cannot be greater than the current time and end time.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n2.最多可获取最近半年（183天）的数据，并且开始时间不能大于当前时间和结束时间"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n2. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception\n3. The default query span is up to 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (you can contact technical support for adjustment, up to 31 days);","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常  \n3.默认允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整，最大31天）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20","zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递该入参时查询账号下所有域名,但当账号下域名数量超过限制时不可查询(报错)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"5m","en":"Data granularity, 5m: 5-minute  granularity, 1h: 1-hour granularity","zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度","exampleValue":"5m,1h"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province\n1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;\n2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"省份\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;\n2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。\n2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension\n1.Options are domain, province, isp, and more than one value can be entered;\n2.The data is displayed according to the specified dimension;","zh_CN":"分组维度\n1.可选值为domain、province、isp,可传入多个值;\n2.有传入则按照该维度展示明细数据;","exampleValue":"domain,province,isp"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetDateFrom(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetDateTo(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetDomain(v []*string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetGranularity(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetProvince(v []*string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.Province = v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetIsp(v []*string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest) SetGroupBy(v []*string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequestHeader struct {
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsPaths struct {
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsParameters struct {
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponse struct {
  // {"en":"Result","zh_CN":"结果"}
  Result []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponse) SetResult(v []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponse {
  s.Result = v
  return s
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  IspData []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult) SetDomain(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult) SetIspData(v []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResult {
  s.IspData = v
  return s
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData struct     {
  // {"en":"ISP","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  ProvinceData []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData) SetIsp(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData) SetProvinceData(v []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspData {
  s.ProvinceData = v
  return s
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData struct     {
  // {"en":"Province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) SetProvince(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData) SetDataSeries(v []*GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceData {
  s.DataSeries = v
  return s
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries struct     {
  // {"en":"Timestamp","zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit:  MB, retain two decimals","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth value. \nUnit is Mbps, retain two decimals","zh_CN":"带宽值：单位为Mbps, 保留两位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) SetTimestamp(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) SetTraffic(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries) SetBandwidth(v string) *GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseResultIspDataProvinceDataDataSeries {
  s.Bandwidth = &v
  return s
}

type GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseHeader struct {
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByIspProvinceForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficAndBandwidthByClientCountryRequest struct {
  // {"en":"Starting time\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. Cannot be greater than the current time\n3. Get up to the last six months (183 days) of data.","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2. 不能大于当前时间\n3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.\n3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception\n4. Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo can&rsquo;t exceed 7 days (can contact technical support adjustment, up to 31 days).","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2. 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。\n3. dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常\n4. 允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最长31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"defaultValue":"5m","en":"Data granularity:\n1.Support 5m (5 minutes granularity),1d (1 day granularity)\n2.Do not pass the default to 5m","zh_CN":"数据粒度:\n1.支持5m(5分钟粒度),1d(天粒度)\n2.不传默认为5m","exampleValue":"5m,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Domains:\n1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment)\n2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment)","zh_CN":"域名：\n1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)\n2.有传递domain时：域名最多支持传20个（可联系技术支持调整）"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"dictionary":"belong=BCS-CC-API|dict=regionCodeList","en":"Country code:\n\n1. If not specified, all countries and regions will be queried by default\n2. Support language request header Accept-Language, only supports zh-CN and en-US, the default is zh-CN. When Accept-Language is en-US, the country and region returned are all in English, otherwise they are returned in Chinese.","zh_CN":"国家地区代号:\n1.不传默认查询全部国家地区\n2. 支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，国家地区返回都为英文，否则返回的为中文。"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Grouping dimension:\n1.Can pass in single or multiple values. The aggregatedOversea and country cannot be passed at the same time;\n2.If there is an incoming, the detailed data will be displayed according to the dimension:\ndomain: Group display by domain name dimension;\ncountry: Group display by country dimension;\naggregatedOversea: Group display according to domestic and overseas dimensions.\n3.The result hierarchy is fixed in order, and the order of the parameters does not affect the order of the returned results. For example: 'groupBy': ['domain','country'] and 'groupBy': ['country','domain'] return the same result.","zh_CN":"分组维度：\n1.可传入单个或多个值,其中不能同时传aggregatedOversea 和 country;\n2.有传入则按照该维度展示明细数据:\n3.domain:按照域名维度进行分组展示;\n4.domain:country:按照国家维度进行分组展示;\n5.aggregatedOversea:按照国内 和 海外维度进行分组展示\n6.返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:groupBy: [domain,country]与groupBy: [country,domain]返回结果一样。","exampleValue":"domain,country,aggregatedOverse"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetDateFrom(v string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetDateTo(v string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetGranularity(v string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetDomain(v []*string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetCountryCode(v []*string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.CountryCode = v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetGroupBy(v []*string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.GroupBy = v
  return s
}

type GetTrafficAndBandwidthByClientCountryRequestHeader struct {
}

func (s GetTrafficAndBandwidthByClientCountryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByClientCountryPaths struct {
}

func (s GetTrafficAndBandwidthByClientCountryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryPaths) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByClientCountryParameters struct {
}

func (s GetTrafficAndBandwidthByClientCountryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryParameters) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByClientCountryResponse struct {
  // {"en":"Result","zh_CN":"结果"}
  Result []*GetTrafficAndBandwidthByClientCountryResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponse) SetResult(v []*GetTrafficAndBandwidthByClientCountryResponseResult) *GetTrafficAndBandwidthByClientCountryResponse {
  s.Result = v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  CountryData []*GetTrafficAndBandwidthByClientCountryResponseResultCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResult) SetDomain(v string) *GetTrafficAndBandwidthByClientCountryResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResult) SetCountryData(v []*GetTrafficAndBandwidthByClientCountryResponseResultCountryData) *GetTrafficAndBandwidthByClientCountryResponseResult {
  s.CountryData = v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseResultCountryData struct     {
  // {"en":"Country code","zh_CN":"国家地区代号"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Country name","zh_CN":"国家地区名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  // {"en":"Summary of traffic in national regions: Summary of traffic flow in a single country region during the query period, unit of measure MB, retain two decimals","zh_CN":"国家地区流量汇总:单个国家地区流量在查询时段内的流量汇总值,计量单位MB,保留2位小数"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"National regional traffic ratio: the proportion (percentage) of traffic value in a single country region during the query period, retain two decimals","zh_CN":"国家地区流量占比:单个国家地区流量在查询时段内的流量值的占比(百分比),保留2位小数"}
  TrafficPercentage *string `json:"trafficPercentage,omitempty" xml:"trafficPercentage,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryData) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetCountryCode(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.CountryCode = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetCountryName(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.CountryName = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetTotalTraffic(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.TotalTraffic = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetTrafficPercentage(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.TrafficPercentage = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetDataSeries(v []*GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.DataSeries = v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries struct     {
  // {"en":"1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.\n2. Returns the time slice contained in the start time and end time.","zh_CN":"1. 查询的数据粒度为5m时,格式为yyyy-MM-dd  HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。\n2. 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\n\nUnit:  MB, retain two decimals","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth value. \nUnit: Mbps , retain two decimals","zh_CN":"带宽值,单位Mbps,保留2位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) SetTimestamp(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) SetTraffic(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) SetBandwidth(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries {
  s.Bandwidth = &v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseHeader struct {
}

func (s GetTrafficAndBandwidthByClientCountryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficByIspProvince1minGranularityRequest struct {
  // {"en":"Start date:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. Cannot exceed current time\n3. The most recent six-month (183 days) data are available.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.不能大于当前时间\n3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.\n3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception\n4. Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常\n4.允许查询最大时间间隔:24小时(可联系技术支持调整)，即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:\n1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);\n2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)\n3. Domain name exceeding limit, misstatement","zh_CN":"域名:\n1.可传递域名数量上限默认为20个(可联系技术支持调整);\n2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)\n3.域名超过上限，提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"1m","en":"Data granularity:\n1. default 1m;\n2. 1m (1 minute), 5m (5 minutes)","zh_CN":"数据粒度:\n1.不传默认1m;\n2.支持1m(1分钟)、5m(5分钟)","exampleValue":"1m,5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province:\n1.If no province  is specified: Query all provinces and aggregate the returned data according to all provinces\n2.Province specified: Send province code, multiple codes can be sent.","zh_CN":"省份：\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份传code，可传多个。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1. If no ISP is specified: all ISPs will be queried, and the returned data will be aggregated by all ISPs\n2. ISP specified: Send ISP code, multiple codes can be provided","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合\n2.有传递isp时：传递运营商code，可传多个"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"defaultValue":"flow","en":"query dimensionality:\n1. Optional value flow, request,bandwidth\n2. Default flow\n3. Flow: Flow, keep two decimal places;\n4. Request: number of Request;\n5. bandwidth: bandwidth, keep two decimal places","zh_CN":"查询维度:\n1.可选值 flow、request、bandwidth\n2.传默认 flow\n3.flow:流量，保留两位小数;\n4.request:请求数;\n5.bandwidth:带宽，保留两位小数","exampleValue":"flow,request,bandwidth"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"defaultValue":"domain","en":"Optional: domain, all, If it is empty, it defaults to returning by domain dimension;\nIf all is passed, merge and return according to the query domain name.","zh_CN":"可选项：domain、all, 为空则默认为按domain维度返回;\n若传递all，则按查询域名合并返回","exampleValue":"domain,all"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s GetTrafficByIspProvince1minGranularityRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetDateFrom(v string) *GetTrafficByIspProvince1minGranularityRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetDateTo(v string) *GetTrafficByIspProvince1minGranularityRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetDomain(v []*string) *GetTrafficByIspProvince1minGranularityRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetGranularity(v string) *GetTrafficByIspProvince1minGranularityRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetProvince(v []*string) *GetTrafficByIspProvince1minGranularityRequest {
  s.Province = v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetIsp(v []*string) *GetTrafficByIspProvince1minGranularityRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetQueryBy(v string) *GetTrafficByIspProvince1minGranularityRequest {
  s.QueryBy = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityRequest) SetGroupBy(v string) *GetTrafficByIspProvince1minGranularityRequest {
  s.GroupBy = &v
  return s
}

type GetTrafficByIspProvince1minGranularityRequestHeader struct {
}

func (s GetTrafficByIspProvince1minGranularityRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficByIspProvince1minGranularityPaths struct {
}

func (s GetTrafficByIspProvince1minGranularityPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityPaths) GoString() string {
  return s.String()
}

type GetTrafficByIspProvince1minGranularityParameters struct {
}

func (s GetTrafficByIspProvince1minGranularityParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityParameters) GoString() string {
  return s.String()
}

type GetTrafficByIspProvince1minGranularityResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"数据结果"}
  Data []*GetTrafficByIspProvince1minGranularityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByIspProvince1minGranularityResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvince1minGranularityResponse) SetCode(v string) *GetTrafficByIspProvince1minGranularityResponse {
  s.Code = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityResponse) SetMessage(v string) *GetTrafficByIspProvince1minGranularityResponse {
  s.Message = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityResponse) SetData(v []*GetTrafficByIspProvince1minGranularityResponseData) *GetTrafficByIspProvince1minGranularityResponse {
  s.Data = v
  return s
}

type GetTrafficByIspProvince1minGranularityResponseData struct     {
  // {"en":"domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"detailList","zh_CN":"数据明细"}
  DataSeries []*GetTrafficByIspProvince1minGranularityResponseDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByIspProvince1minGranularityResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityResponseData) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvince1minGranularityResponseData) SetDomain(v string) *GetTrafficByIspProvince1minGranularityResponseData {
  s.Domain = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityResponseData) SetDataSeries(v []*GetTrafficByIspProvince1minGranularityResponseDataDataSeries) *GetTrafficByIspProvince1minGranularityResponseData {
  s.DataSeries = v
  return s
}

type GetTrafficByIspProvince1minGranularityResponseDataDataSeries struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM","zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"1. Flow: Flow, keep two decimal places;\n2. Bandwidth: Bandwidth, keep two decimals;\n3. Request: number of Request","zh_CN":"1.flow：流量，保留两位小数；\n2.bandwidth: 带宽，保留两位小数；\n3.request：请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTrafficByIspProvince1minGranularityResponseDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityResponseDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvince1minGranularityResponseDataDataSeries) SetTimestamp(v string) *GetTrafficByIspProvince1minGranularityResponseDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficByIspProvince1minGranularityResponseDataDataSeries) SetValue(v string) *GetTrafficByIspProvince1minGranularityResponseDataDataSeries {
  s.Value = &v
  return s
}

type GetTrafficByIspProvince1minGranularityResponseHeader struct {
}

func (s GetTrafficByIspProvince1minGranularityResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvince1minGranularityResponseHeader) GoString() string {
  return s.String()
}




type FlowChannelRequest struct {
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
  // {"en":"acceleration type:\n1.If there are multiple inputs,use ';' as separator.\n2.If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1.optional values:xml, json.\n2.'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:\n1.'true' as default.\n2. If not 'true',it will query data of channels that ends with any item of input 'channel's.","zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=ispCode","en":"The abbreviations of the ISPs to be queried. For multiple ISPs, please separate them with a semicolon ';'. Note: Only when the region is specified as 'cn' does the ISP information take effect. If not selected or left blank, all ISPs will be included by default.","zh_CN":"要查询的运营商的缩写，多个isp请用英文分号';'分隔开。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"Display statistic result in merged or separate way:\n1.If specified 1,get the merged result.\n2.If specified 2,get the separate result.\n3.If specified 3,get both merged result and separate result.\n4.If not specified,means '1'.","zh_CN":"结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s FlowChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelRequest) GoString() string {
  return s.String()
}

func (s *FlowChannelRequest) SetDate(v string) *FlowChannelRequest {
  s.Date = &v
  return s
}

func (s *FlowChannelRequest) SetStartdate(v string) *FlowChannelRequest {
  s.Startdate = &v
  return s
}

func (s *FlowChannelRequest) SetEnddate(v string) *FlowChannelRequest {
  s.Enddate = &v
  return s
}

func (s *FlowChannelRequest) SetChannel(v string) *FlowChannelRequest {
  s.Channel = &v
  return s
}

func (s *FlowChannelRequest) SetRegion(v string) *FlowChannelRequest {
  s.Region = &v
  return s
}

func (s *FlowChannelRequest) SetAccetype(v string) *FlowChannelRequest {
  s.Accetype = &v
  return s
}

func (s *FlowChannelRequest) SetDataformat(v string) *FlowChannelRequest {
  s.Dataformat = &v
  return s
}

func (s *FlowChannelRequest) SetIsExactMatch(v string) *FlowChannelRequest {
  s.IsExactMatch = &v
  return s
}

func (s *FlowChannelRequest) SetIsp(v string) *FlowChannelRequest {
  s.Isp = &v
  return s
}

func (s *FlowChannelRequest) SetResultType(v string) *FlowChannelRequest {
  s.ResultType = &v
  return s
}

type FlowChannelRequestHeader struct {
}

func (s FlowChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelRequestHeader) GoString() string {
  return s.String()
}

type FlowChannelPaths struct {
}

func (s FlowChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelPaths) GoString() string {
  return s.String()
}

type FlowChannelParameters struct {
}

func (s FlowChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelParameters) GoString() string {
  return s.String()
}

type FlowChannelResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *FlowChannelResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s FlowChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponse) GoString() string {
  return s.String()
}

func (s *FlowChannelResponse) SetProvider(v *FlowChannelResponseProvider) *FlowChannelResponse {
  s.Provider = v
  return s
}

type FlowChannelResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"resultType","zh_CN":"统计类型"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {"en":"data","zh_CN":"频道流量数据"}
  Date *FlowChannelResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s FlowChannelResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponseProvider) GoString() string {
  return s.String()
}

func (s *FlowChannelResponseProvider) SetName(v string) *FlowChannelResponseProvider {
  s.Name = &v
  return s
}

func (s *FlowChannelResponseProvider) SetType(v string) *FlowChannelResponseProvider {
  s.Type = &v
  return s
}

func (s *FlowChannelResponseProvider) SetResultType(v string) *FlowChannelResponseProvider {
  s.ResultType = &v
  return s
}

func (s *FlowChannelResponseProvider) SetDate(v *FlowChannelResponseProviderDate) *FlowChannelResponseProvider {
  s.Date = v
  return s
}

type FlowChannelResponseProviderDate struct {
  // {"en":"name","zh_CN":"日期"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"channel","zh_CN":"频道"}
  Channel *FlowChannelResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s FlowChannelResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponseProviderDate) GoString() string {
  return s.String()
}

func (s *FlowChannelResponseProviderDate) SetName(v string) *FlowChannelResponseProviderDate {
  s.Name = &v
  return s
}

func (s *FlowChannelResponseProviderDate) SetChannel(v *FlowChannelResponseProviderDateChannel) *FlowChannelResponseProviderDate {
  s.Channel = v
  return s
}

type FlowChannelResponseProviderDateChannel struct {
  // {"en":"channel","zh_CN":"频道"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"流量数据"}
  Flow []*FlowChannelResponseProviderDateChannelFlow `json:"flow,omitempty" xml:"flow,omitempty" require:"true" type:"Repeated"`
}

func (s FlowChannelResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *FlowChannelResponseProviderDateChannel) SetName(v string) *FlowChannelResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *FlowChannelResponseProviderDateChannel) SetFlow(v []*FlowChannelResponseProviderDateChannelFlow) *FlowChannelResponseProviderDateChannel {
  s.Flow = v
  return s
}

type FlowChannelResponseProviderDateChannelFlow struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"bandwidth","zh_CN":"流量,单位Byte"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s FlowChannelResponseProviderDateChannelFlow) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponseProviderDateChannelFlow) GoString() string {
  return s.String()
}

func (s *FlowChannelResponseProviderDateChannelFlow) SetTime(v string) *FlowChannelResponseProviderDateChannelFlow {
  s.Time = &v
  return s
}

func (s *FlowChannelResponseProviderDateChannelFlow) SetText(v string) *FlowChannelResponseProviderDateChannelFlow {
  s.Text = &v
  return s
}

type FlowChannelResponseHeader struct {
}

func (s FlowChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponseHeader) GoString() string {
  return s.String()
}




type WctQueryRequest struct {
  // {"en":"cust_en_name of sub-client.\nWhen a merged-account wants to  view the information of the subclient,the cust_en_name is required.","zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not specified,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope.\n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope.\n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"trans;hevc;audio_trans;encap_trans;vframe;file_op;trans_nbhd;hevc_nbhd;video_enhance;drm.","zh_CN":"转码操作类型：trans(H.264视频转码);hevc(H.265视频转码);audio_trans(音频转码);encap_trans(转封装);vframe(截图);file_op(文件处理);trans_nbhd(H.264智控高清视频转码);hevc_nbhd(H.265智控高清视频转码);video_enhance(AI视频增强);drm(DRM加密)"}
  Transcoding *string `json:"transcoding,omitempty" xml:"transcoding,omitempty"`
  // {"en":"sd240;sd480;sd720;hd1080;2k;4k.","zh_CN":"梯度: sd240;sd480;sd720;hd1080;2k;4k"}
  TranscodingType *string `json:"transcodingType,omitempty" xml:"transcodingType,omitempty"`
  // {"en":"Space; for multiple values, please use a semicolon ';' to separate them.","zh_CN":"空间，多个值请用英文分号“;”分割"}
  Space *string `json:"space,omitempty" xml:"space,omitempty"`
  // {"en":"Source provider of the feature. If there are multiple values, please separate them with a semicolon ';'.","zh_CN":"功能来源，提供方。多个值请用英文分号“;”分割"}
  Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
}

func (s WctQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s WctQueryRequest) GoString() string {
  return s.String()
}

func (s *WctQueryRequest) SetCust(v string) *WctQueryRequest {
  s.Cust = &v
  return s
}

func (s *WctQueryRequest) SetDate(v string) *WctQueryRequest {
  s.Date = &v
  return s
}

func (s *WctQueryRequest) SetStartdate(v string) *WctQueryRequest {
  s.Startdate = &v
  return s
}

func (s *WctQueryRequest) SetEnddate(v string) *WctQueryRequest {
  s.Enddate = &v
  return s
}

func (s *WctQueryRequest) SetTimezone(v string) *WctQueryRequest {
  s.Timezone = &v
  return s
}

func (s *WctQueryRequest) SetTranscoding(v string) *WctQueryRequest {
  s.Transcoding = &v
  return s
}

func (s *WctQueryRequest) SetTranscodingType(v string) *WctQueryRequest {
  s.TranscodingType = &v
  return s
}

func (s *WctQueryRequest) SetSpace(v string) *WctQueryRequest {
  s.Space = &v
  return s
}

func (s *WctQueryRequest) SetProvider(v string) *WctQueryRequest {
  s.Provider = &v
  return s
}

func (s *WctQueryRequest) SetRegion(v string) *WctQueryRequest {
  s.Region = &v
  return s
}

func (s *WctQueryRequest) SetDataformat(v string) *WctQueryRequest {
  s.Dataformat = &v
  return s
}

type WctQueryRequestHeader struct {
}

func (s WctQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s WctQueryRequestHeader) GoString() string {
  return s.String()
}

type WctQueryPaths struct {
}

func (s WctQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s WctQueryPaths) GoString() string {
  return s.String()
}

type WctQueryParameters struct {
}

func (s WctQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s WctQueryParameters) GoString() string {
  return s.String()
}

type WctQueryResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *WctQueryResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponse) GoString() string {
  return s.String()
}

func (s *WctQueryResponse) SetProvider(v *WctQueryResponseProvider) *WctQueryResponse {
  s.Provider = v
  return s
}

type WctQueryResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"转码数据"}
  Date *WctQueryResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProvider) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProvider) SetName(v string) *WctQueryResponseProvider {
  s.Name = &v
  return s
}

func (s *WctQueryResponseProvider) SetType(v string) *WctQueryResponseProvider {
  s.Type = &v
  return s
}

func (s *WctQueryResponseProvider) SetDate(v *WctQueryResponseProviderDate) *WctQueryResponseProvider {
  s.Date = v
  return s
}

type WctQueryResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"customerId","zh_CN":"客户id"}
  CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty" require:"true"`
  // {"en":"transcoding","zh_CN":"转码类型"}
  Transcoding *string `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true"`
  // {"en":"transcodingType","zh_CN":"梯度"}
  TranscodingType *string `json:"transcodingType,omitempty" xml:"transcodingType,omitempty" require:"true"`
  // {"en":"unit","zh_CN":"单位"}
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty" require:"true"`
  // {"en":"total","zh_CN":"总数"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"wct","zh_CN":"转码数据"}
  Wct *WctQueryResponseProviderDateWct `json:"wct,omitempty" xml:"wct,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProviderDate) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProviderDate) SetStartdate(v string) *WctQueryResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetEnddate(v string) *WctQueryResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetCustomerId(v string) *WctQueryResponseProviderDate {
  s.CustomerId = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetTranscoding(v string) *WctQueryResponseProviderDate {
  s.Transcoding = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetTranscodingType(v string) *WctQueryResponseProviderDate {
  s.TranscodingType = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetUnit(v string) *WctQueryResponseProviderDate {
  s.Unit = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetTotal(v string) *WctQueryResponseProviderDate {
  s.Total = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetWct(v *WctQueryResponseProviderDateWct) *WctQueryResponseProviderDate {
  s.Wct = v
  return s
}

type WctQueryResponseProviderDateWct struct {
  // {"en":"type","zh_CN":"类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"total","zh_CN":"总数"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"data","zh_CN":"明细数据"}
  Data []*WctQueryResponseProviderDateWctData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s WctQueryResponseProviderDateWct) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProviderDateWct) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProviderDateWct) SetType(v string) *WctQueryResponseProviderDateWct {
  s.Type = &v
  return s
}

func (s *WctQueryResponseProviderDateWct) SetTotal(v string) *WctQueryResponseProviderDateWct {
  s.Total = &v
  return s
}

func (s *WctQueryResponseProviderDateWct) SetData(v []*WctQueryResponseProviderDateWctData) *WctQueryResponseProviderDateWct {
  s.Data = v
  return s
}

type WctQueryResponseProviderDateWctData struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"hit count","zh_CN":"明细数"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s WctQueryResponseProviderDateWctData) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProviderDateWctData) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProviderDateWctData) SetTime(v string) *WctQueryResponseProviderDateWctData {
  s.Time = &v
  return s
}

func (s *WctQueryResponseProviderDateWctData) SetText(v string) *WctQueryResponseProviderDateWctData {
  s.Text = &v
  return s
}

type WctQueryResponseHeader struct {
}

func (s WctQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseHeader) GoString() string {
  return s.String()
}




type GetTotalTrafficForAllDomainsRequest struct {
}

func (s GetTotalTrafficForAllDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsRequest) GoString() string {
  return s.String()
}

type GetTotalTrafficForAllDomainsRequestHeader struct {
}

func (s GetTotalTrafficForAllDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTotalTrafficForAllDomainsPaths struct {
}

func (s GetTotalTrafficForAllDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsPaths) GoString() string {
  return s.String()
}

type GetTotalTrafficForAllDomainsParameters struct {
  // {"en":"Start time 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; 2.Must be smaller than the current time and dateto; 3.Period between datafrom and dateto cannot be longer than 31 days","zh_CN":"开始时间1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;2.必须小于当前时间和dateto;3.dateFrom和dateTo相差不能超过31天;4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be greater than 'datefrom';\n3.if it's greater than the current time, then the current time is assigned as the value","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于datefrom;如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"defaultValue":"daily","en":"Data granularity\n1.fiveminutes: five minutes, hourly: one hour, daily: one day;\n2.If not specified, daily is set as the default value;\n3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is a specific configuration on data collecting granularity for the custome","zh_CN":"数据粒度\n1.fiveminutes:5分钟,hourly:1小时,daily:1天;\n2.不传递,默认为daily;\n3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。","exampleValue":"fiveminutes,hourly,daily"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
}

func (s GetTotalTrafficForAllDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsParameters) GoString() string {
  return s.String()
}

func (s *GetTotalTrafficForAllDomainsParameters) SetDatefrom(v string) *GetTotalTrafficForAllDomainsParameters {
  s.Datefrom = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsParameters) SetDateto(v string) *GetTotalTrafficForAllDomainsParameters {
  s.Dateto = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsParameters) SetGranularity(v string) *GetTotalTrafficForAllDomainsParameters {
  s.Granularity = &v
  return s
}

type GetTotalTrafficForAllDomainsResponse struct {
  // {"en":"Total traffic. Unit: MB, retain two decimals","zh_CN":"总流量,保留2位小数,单位为MB"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Traffic data","zh_CN":"流量数据"}
  DataSeries []*GetTotalTrafficForAllDomainsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTotalTrafficForAllDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTotalTrafficForAllDomainsResponse) SetTotalTraffic(v string) *GetTotalTrafficForAllDomainsResponse {
  s.TotalTraffic = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsResponse) SetDataSeries(v []*GetTotalTrafficForAllDomainsResponseDataSeries) *GetTotalTrafficForAllDomainsResponse {
  s.DataSeries = v
  return s
}

type GetTotalTrafficForAllDomainsResponseDataSeries struct     {
  // {"en":"Time:\n1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.\n2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.\n3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.\n4.Return the time slice contained in start time and the time slice contained in end time","zh_CN":"时间: \n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic.  Unit: MB ,retain two decimals","zh_CN":"流量,保留2位小数,单位为MB"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s GetTotalTrafficForAllDomainsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetTotalTrafficForAllDomainsResponseDataSeries) SetTimestamp(v string) *GetTotalTrafficForAllDomainsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsResponseDataSeries) SetTraffic(v string) *GetTotalTrafficForAllDomainsResponseDataSeries {
  s.Traffic = &v
  return s
}

type GetTotalTrafficForAllDomainsResponseHeader struct {
}

func (s GetTotalTrafficForAllDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetCdnOriginTrafficRequest struct {
  // {"en":"Start Time:\n1.The time format is yyyy-MM-dd\n2.Cannot exceed the current date\n3.Up to the past 183 days of data can be obtained","zh_CN":"开始时间：\n1.时间格式为 yyyy-MM-dd\n2.不能大于当前日期\n3.最多可获取最近183天的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:\n1.The time format is yyyy-MM-dd\n2.The end time must be greater than the start time\n3.If the end time greater than the current time, the current time is taken\n4. Time granularity 5m maximum query interval allowed: 7 days; Time granularity 1h or 1d maximum query interval allowed: 31 days.","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-dd\n2.结束时间需大于等于开始时间\n3.结束时间如果大于当前时间，取当前时间\n4.5分钟粒度允许查询最大间隔：7天；1小时或天粒度允许查询最大间隔：31天；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"defaultValue":"GMT+08:00","en":"Specify the time zone for both the query time range(dateFrom/dateTo) and the returned data.\nGreenwich Mean Time Zone, the default time zone is GMT+08:00.\nIf you wish to specify a different time zone, you can append a time zone identifier, i.e. GMT+09:00 or GMT-10:00.","zh_CN":"指定查询时间(dateFrom/dateTo)和返回数据的时区。格林尼治时区，默认时区是GMT+08:00。如果您希望指定不同的时区，可以附加时区标识，即GMT+09:00或GMT-10:00。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"defaultValue":"up","en":"Rounds up or down a timestamp by a given time period. Only valid when the granularity is less than 1d(not including 1day)\n1. up-rounds up, e.g.: granularity=5m, 00:00:00-00:04:49 will be displayed as 00:05:00\n2. down-rounds down, e.g.: granularity=5m, 00:00:00-00:04:49 will be displayed as 00:00:00\n3. If not specified, the result will be rounded up (up).","zh_CN":"根据指定的时间周期对时间戳进行向上或向下取整。仅当粒度小于1天（不包含1天）时有效。\n1. up-向上取整。例如：granularity=5m，00:00:00-00:04:49 将会显示为 00:05:00\n2. down-向下取整，例如：granularity=5m时，00:00:00-00:04:49将显示为00:00:00\n3. 如未传值，结果将进行向上取整(up)"}
  TimeRounding *string `json:"timeRounding,omitempty" xml:"timeRounding,omitempty"`
  // {"en":"Domains:\n1.Domain is not uploaded: Query all domain names of the account \n2.Domain is uploaded: Up to 2000 domains are supported (you can contact technical support for adjustment)\n3.For multiple domain, please separate them with an English semicolon \";\"","zh_CN":"域名：\n1.未传递domain时：查询账号下所有全部域名\n2.有传递domain时：域名最多支持传2000个（可联系技术支持调整）\n3.多个域名用英文分号\";\"分隔"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"Billing region of the Acceleration domain:\n1. If not specified, it is considered as no restriction on billing region.\n2. For multiple billing regions, please separate them with an English semicolon \";\".eg: cn;kr\n3. Available Common regions: cn: Mainland China, hk: Hong Kong, ov: Oversea average, tw: Taiwan, euna: Europe and North America, apac: Asia-Pacific, sa: South America, af: Africa, am: Americas, emea: Europe/Middle East/Africa, kr: South Korea, au: Australia, in: India, jp: Japan, ru: Russia, indo: Indonesia, me: Middle East, eu: Europe, ph: Philippines","zh_CN":"加速域名的计费区域：\n1.未传递视为不限计费区域\n2.多个计费区域请用英文分号\";\"分隔。如：cn;hk\n3.计费常用区域：cn:中国大陆,hk:香港,ov:海外平均,tw:台湾,euna:欧美,apac:亚太,sa:南美,af:非洲,am:美洲,emea:欧洲/中东/非洲,kr:韩国,au:澳大利亚,in:印度,jp:日本,ru:俄罗斯,indo:印尼,me:中东,eu:欧洲,ph:菲律宾"}
  BillingRegion *string `json:"billingRegion,omitempty" xml:"billingRegion,omitempty"`
  // {"defaultValue":"1d","en":"Time granularityZ: the default granularity is 1day.5m: 5minutes;1h: 1hour;1d: 1day;","zh_CN":"数据粒度：默认1天粒度。\n5m：5分钟粒度；\n1h：1小时粒度；\n1d：1天粒度；"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Group keywords:\n1.If not specified, result will be aggregated as default. \n2.Passing in a keyword means that the details need to be displayed according to the values corresponding to the keyword dimension grouping (for example, passing in domain means that the details are expanded according to domain). If an invalid value is specified, it will return error.\n3.Support passing multiple values, for multiple value please separate them with an English semicolon \";\",only support domain grouping currently","zh_CN":"分组关键词：\n1.未传递时，默认聚合展示\n2.传入关键词则代表需要按照关键词维度分组对应的值展示明细（例如传domain，则代表返回按照domain明细展开）,如传入不支持的关键词，返回相应的错误提示\n3.支持传多个值，传多个请用英文分号\";\"分隔，当前只支持domain"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s GetCdnOriginTrafficRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficRequest) GoString() string {
  return s.String()
}

func (s *GetCdnOriginTrafficRequest) SetDateFrom(v string) *GetCdnOriginTrafficRequest {
  s.DateFrom = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetDateTo(v string) *GetCdnOriginTrafficRequest {
  s.DateTo = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetTimezone(v string) *GetCdnOriginTrafficRequest {
  s.Timezone = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetTimeRounding(v string) *GetCdnOriginTrafficRequest {
  s.TimeRounding = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetDomain(v string) *GetCdnOriginTrafficRequest {
  s.Domain = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetBillingRegion(v string) *GetCdnOriginTrafficRequest {
  s.BillingRegion = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetGranularity(v string) *GetCdnOriginTrafficRequest {
  s.Granularity = &v
  return s
}

func (s *GetCdnOriginTrafficRequest) SetGroupBy(v string) *GetCdnOriginTrafficRequest {
  s.GroupBy = &v
  return s
}

type GetCdnOriginTrafficRequestHeader struct {
}

func (s GetCdnOriginTrafficRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficRequestHeader) GoString() string {
  return s.String()
}

type GetCdnOriginTrafficPaths struct {
}

func (s GetCdnOriginTrafficPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficPaths) GoString() string {
  return s.String()
}

type GetCdnOriginTrafficParameters struct {
}

func (s GetCdnOriginTrafficParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficParameters) GoString() string {
  return s.String()
}

type GetCdnOriginTrafficResponse struct {
  // {"en":"Response code","zh_CN":"响应状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Reponse message","zh_CN":"响应状态码"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *GetCdnOriginTrafficResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCdnOriginTrafficResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficResponse) GoString() string {
  return s.String()
}

func (s *GetCdnOriginTrafficResponse) SetCode(v string) *GetCdnOriginTrafficResponse {
  s.Code = &v
  return s
}

func (s *GetCdnOriginTrafficResponse) SetMessage(v string) *GetCdnOriginTrafficResponse {
  s.Message = &v
  return s
}

func (s *GetCdnOriginTrafficResponse) SetData(v *GetCdnOriginTrafficResponseData) *GetCdnOriginTrafficResponse {
  s.Data = v
  return s
}

type GetCdnOriginTrafficResponseData struct {
  // {"en":"Container for data information. It will contain multiple <data-report> child element.","zh_CN":"数据信息的容器。它将包含多个子元素"}
  Report *GetCdnOriginTrafficResponseDataReport `json:"report,omitempty" xml:"report,omitempty" require:"true" type:"Struct"`
}

func (s GetCdnOriginTrafficResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficResponseData) GoString() string {
  return s.String()
}

func (s *GetCdnOriginTrafficResponseData) SetReport(v *GetCdnOriginTrafficResponseDataReport) *GetCdnOriginTrafficResponseData {
  s.Report = v
  return s
}

type GetCdnOriginTrafficResponseDataReport struct {
  // {"en":"The metric name","zh_CN":"接口数据名称"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Indicates the starting time of response data, the time format is yyyy-MM-dd","zh_CN":"标记返回数据的开始时间,格式为 yyyy-MM-dd"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"Indicates the ending time of response data, the time format is yyyy-MM-dd","zh_CN":"标记返回数据的结束时间,格式为 yyyy-MM-dd"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Indicates the time zone for the returned data. Greenwich Mean Time Zone, the format is GMT+08:00","zh_CN":"标记返回数据的时区,格林尼治时区，格式是GMT+08:00。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true"`
  // {"en":"List of data items. Each item contains fields as following: 'domain', 'totalOriginTraffic', 'peakBandwidth', 'peakTime' ","zh_CN":"数据项集合，每个数据项包含：domain，totalOriginTraffic，peakBandwidth，peakTime字段"}
  Groups []*GetCdnOriginTrafficResponseDataReportGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnOriginTrafficResponseDataReport) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficResponseDataReport) GoString() string {
  return s.String()
}

func (s *GetCdnOriginTrafficResponseDataReport) SetType(v string) *GetCdnOriginTrafficResponseDataReport {
  s.Type = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReport) SetDateFrom(v string) *GetCdnOriginTrafficResponseDataReport {
  s.DateFrom = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReport) SetDateTo(v string) *GetCdnOriginTrafficResponseDataReport {
  s.DateTo = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReport) SetTimezone(v string) *GetCdnOriginTrafficResponseDataReport {
  s.Timezone = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReport) SetGroups(v []*GetCdnOriginTrafficResponseDataReportGroups) *GetCdnOriginTrafficResponseDataReport {
  s.Groups = v
  return s
}

type GetCdnOriginTrafficResponseDataReportGroups struct     {
  // {"en":"Domain","zh_CN":"频道"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Indicates the CDN node-to-origin traffic. Measurement unit GB, keep 3 decimal places.","zh_CN":"CDN节点到源站的汇总流量，计量单位GB，保留3位小数。"}
  TotalOriginTraffic *string `json:"totalOriginTraffic,omitempty" xml:"totalOriginTraffic,omitempty" require:"true"`
  // {"en":"Indicates the peak bandwidth(5-minute granularity) of CDN-to-Client traffic. Measurement unit GB, keep 3 decimal places. If granularity<5m, no display in groups container objects","zh_CN":"带宽峰值(5分钟粒度），计量单位Mbps，保留3位小数，如果granularity<5m，则在[groups]分组容器对象中不显示"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"Indicates the time of peak bandwidth(5-minute granularity). ","zh_CN":"峰值时间(5分钟粒度）"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"List of data items. Each item contains fields as following: 'time', 'originTraffic'","zh_CN":"明细数据列表，每项包含time和originTraffic字段"}
  Item []*GetCdnOriginTrafficResponseDataReportGroupsItem `json:"item,omitempty" xml:"item,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnOriginTrafficResponseDataReportGroups) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficResponseDataReportGroups) GoString() string {
  return s.String()
}

func (s *GetCdnOriginTrafficResponseDataReportGroups) SetDomain(v string) *GetCdnOriginTrafficResponseDataReportGroups {
  s.Domain = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReportGroups) SetTotalOriginTraffic(v string) *GetCdnOriginTrafficResponseDataReportGroups {
  s.TotalOriginTraffic = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReportGroups) SetPeakBandwidth(v string) *GetCdnOriginTrafficResponseDataReportGroups {
  s.PeakBandwidth = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReportGroups) SetPeakTime(v string) *GetCdnOriginTrafficResponseDataReportGroups {
  s.PeakTime = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReportGroups) SetItem(v []*GetCdnOriginTrafficResponseDataReportGroupsItem) *GetCdnOriginTrafficResponseDataReportGroups {
  s.Item = v
  return s
}

type GetCdnOriginTrafficResponseDataReportGroupsItem struct     {
  // {"en":"Indicates the date and time\n1. If the type is \"5m\", the timestamp format is \"YYYY-MM-DD hh:mm:00\", that means the period's end time.\n2. If the type is \"1h\", the timestamp format is \"YYYY-MM-DD hh:00:00\", that means the period's start time.\n3. If the type is \"1d\", the timestamp format is \"YYYY-MM-DD 00:00:00\", that means the period's start time.","zh_CN":"表示日期和时间\n1. 如果类型为“5m”，则时间戳格式为“YYYY-MM-DD hh:mm:00”，表示该时间段的结束时间。\n2. 如果类型为“1h”，则时间戳格式为“YYYY-MM-DD hh:00:00”，表示该时间段的开始时间。\n3. 如果类型为“1d”，则时间戳格式为“YYYY-MM-DD 00:00:00”，表示该时间段的开始时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Indicates the CDN node-to-Origin traffic. Measurement unit GB, keep 3 decimal places.","zh_CN":"CDN节点到源站的流量，计量单位GB，保留3位小数。"}
  OriginTraffic *string `json:"originTraffic,omitempty" xml:"originTraffic,omitempty" require:"true"`
}

func (s GetCdnOriginTrafficResponseDataReportGroupsItem) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficResponseDataReportGroupsItem) GoString() string {
  return s.String()
}

func (s *GetCdnOriginTrafficResponseDataReportGroupsItem) SetTime(v string) *GetCdnOriginTrafficResponseDataReportGroupsItem {
  s.Time = &v
  return s
}

func (s *GetCdnOriginTrafficResponseDataReportGroupsItem) SetOriginTraffic(v string) *GetCdnOriginTrafficResponseDataReportGroupsItem {
  s.OriginTraffic = &v
  return s
}

type GetCdnOriginTrafficResponseHeader struct {
}

func (s GetCdnOriginTrafficResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginTrafficResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowIspProvinceShortTimeServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:00:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.;", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 1 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 hours, that is, the difference between dateFrom and dateTo can not exceed 2 hours ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：2小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error , you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).", "zh_CN":"域名：
  // 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)；
  // 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 1m: 1 minute granularity;
  // 5m: 5 minute granularity, Default value is 5m;
  // 1h: 1 hour granularity.", "zh_CN":"数据粒度：
  // 1m：1分钟粒度；
  // 5m：5分钟粒度。不传默认5分钟粒度；
  // 1h：1小时粒度。"}
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
  // 1.Options are domain, province, isp, and more than one value can be entered;
  // 2.The data is displayed according to the specified dimension.", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetDateFrom(v string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetDateTo(v string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetDomain(v []*string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetDataInterval(v string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetProvince(v []*string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetIsp(v []*string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceRequest) SetGroupBy(v []*string) *ReportFlowIspProvinceShortTimeServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportFlowIspProvinceShortTimeServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceResponse) SetResult(v []*ReportFlowIspProvinceShortTimeServiceResponseResult) *ReportFlowIspProvinceShortTimeServiceResponse {
  s.Result = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"ISP数据"}
  IspData []*ReportFlowIspProvinceShortTimeServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResult) SetDomain(v string) *ReportFlowIspProvinceShortTimeServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResult) SetIspData(v []*ReportFlowIspProvinceShortTimeServiceResponseResultIspData) *ReportFlowIspProvinceShortTimeServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份数据"}
  ProvinceData []*ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspData) SetIsp(v string) *ReportFlowIspProvinceShortTimeServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspData) SetProvinceData(v []*ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) *ReportFlowIspProvinceShortTimeServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"flowdata", "zh_CN":"流量数据"}
  FlowData []*ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) SetFlowData(v []*ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData {
  s.FlowData = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits   of decimals allowed", "zh_CN":"流量值,单位为MB,保留两位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Bandwidth value. Unit is Mbps and   2 digits of decimals allowed", "zh_CN":"带宽值,单位为Mbps,保留两位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) SetTimestamp(v string) *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) SetValue(v string) *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData {
  s.Value = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) SetBandwidth(v string) *ReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData {
  s.Bandwidth = &v
  return s
}

type ReportFlowIspProvinceShortTimeServicePaths struct {
}

func (s ReportFlowIspProvinceShortTimeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServicePaths) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceShortTimeServiceParameters struct {
}

func (s ReportFlowIspProvinceShortTimeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceShortTimeServiceRequestHeader struct {
}

func (s ReportFlowIspProvinceShortTimeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceShortTimeServiceResponseHeader struct {
}

func (s ReportFlowIspProvinceShortTimeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest struct {
  // {"en":"Start date:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. Cannot exceed current time\n3. The most recent six-month (183 days) data are available.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.不能大于当前时间\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.\n3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception\n4. Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常\n4.允许查询最大时间间隔:24小时（可联系技术支持调整），即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:\n1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);\n2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)\n3. Domain name exceeding limit, misstatement","zh_CN":"域名:\n1.可传递域名数量上限默认为20个（可联系技术支持调整）;\n2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）\n3.域名超过上限，提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"1m","en":"Data granularity:\n1. default 1m;\n2. 1m (1 minute), 5m (5 minutes)","zh_CN":"数据粒度:\n1.不传默认1m;\n2.支持1m（1分钟）、5m（5分钟）","exampleValue":"1m,5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province:\n1.If no province  is specified: Query all provinces and aggregate the returned data according to all provinces\n2.Province specified: Send province code, multiple codes can be sent.","zh_CN":"省份：\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份传code，可传多个。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1. If no ISP is specified: all ISPs will be queried, and the returned data will be aggregated by all ISPs\n2. ISP specified: Send ISP code, multiple codes can be provided","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合\n2.有传递isp时：传递运营商code，可传多个"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"defaultValue":"flow","en":"query dimensionality:\n1. Optional value flow, request\n2. Default flow\n3. Flow: Flow Unit MB, keep two decimal places;\n4. Request: number of Request","zh_CN":"查询维度:\n1.可选值 flow、request\n2.传默认 flow\n3.flow:流量，单位MB，保留两位小数;\n4.request:请求数","exampleValue":"flow,request"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"defaultValue":"domain","en":"Optional: \ndomain, all, If it is empty, it defaults to returning by domain dimension;\nIf all is passed, merge and return according to the query domain name.","zh_CN":"可选项：\ndomain、all, 为空则默认为按domain维度返回;\n若传递all，则按查询域名合并返回","exampleValue":"domain,all"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetDateFrom(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetDateTo(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetDomain(v []*string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetGranularity(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetProvince(v []*string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Province = v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetIsp(v []*string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetQueryBy(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.QueryBy = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetGroupBy(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.GroupBy = &v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequestHeader struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityPaths struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityPaths) GoString() string {
  return s.String()
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityParameters struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityParameters) GoString() string {
  return s.String()
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request","zh_CN":"请求结果的详细数据"}
  Data []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) SetCode(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse {
  s.Code = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) SetMessage(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse {
  s.Message = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) SetData(v []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse {
  s.Data = v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) SetDomain(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData {
  s.Domain = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) SetDataSeries(v []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData {
  s.DataSeries = v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries struct     {
  // {"en":"Time, in yyyy-MM-dd HH:MM","zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Hit data:\nFlow: Flow, keep two decimal places;\nRequest: number of Request","zh_CN":"命中数据:\n1.flow:流量，保留两位小数;\n2.request:请求数"}
  HitValue *string `json:"hitValue,omitempty" xml:"hitValue,omitempty" require:"true"`
  // {"en":"Hit rate, keep four decimal places","zh_CN":"命中率，保留四位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) SetTimestamp(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) SetHitValue(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries {
  s.HitValue = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) SetHitRate(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries {
  s.HitRate = &v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseHeader struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeByteHitRatioServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:00:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.;", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1.Domain is not uploaded: Query all domain names of the account (More than 200 domains will error , you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 200 domains are supported (you can contact technical support for adjustment).", "zh_CN":"域名：
  // 1.未传递domain时：查询账号下所有全部域名(域名超过200个则报错，可联系技术支持调整)；
  // 2.有传递domain时：域名最多支持传200个（可联系技术支持调整）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceRequest) SetDateFrom(v string) *QueryEdgeByteHitRatioServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceRequest) SetDateTo(v string) *QueryEdgeByteHitRatioServiceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceRequest) SetDomain(v []*string) *QueryEdgeByteHitRatioServiceRequest {
  s.Domain = v
  return s
}

type QueryEdgeByteHitRatioServiceResponse struct {
  Result []*QueryEdgeByteHitRatioServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponse) SetResult(v []*QueryEdgeByteHitRatioServiceResponseResult) *QueryEdgeByteHitRatioServiceResponse {
  s.Result = v
  return s
}

type QueryEdgeByteHitRatioServiceResponseResult struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间，格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total edge node hit ratio", "zh_CN":"总边缘节点命中率的平均值,2位小数"}
  TotalAvg *string `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  HitRatioDatas []*QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponseResult) SetRealDate(v string) *QueryEdgeByteHitRatioServiceResponseResult {
  s.RealDate = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceResponseResult) SetTotalAvg(v string) *QueryEdgeByteHitRatioServiceResponseResult {
  s.TotalAvg = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceResponseResult) SetHitRatioDatas(v []*QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) *QueryEdgeByteHitRatioServiceResponseResult {
  s.HitRatioDatas = v
  return s
}

type QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas struct     {
  // {"en":"timestamp", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"edge node hit ratio,keep 4 decimal places", "zh_CN":"边缘节点缓存字节命中率，保留4位小数"}
  EdgeHitRatio *string `json:"edgeHitRatio,omitempty" xml:"edgeHitRatio,omitempty" require:"true"`
}

func (s QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) SetTimestamp(v string) *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) SetEdgeHitRatio(v string) *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas {
  s.EdgeHitRatio = &v
  return s
}

type QueryEdgeByteHitRatioServicePaths struct {
}

func (s QueryEdgeByteHitRatioServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServicePaths) GoString() string {
  return s.String()
}

type QueryEdgeByteHitRatioServiceParameters struct {
}

func (s QueryEdgeByteHitRatioServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceParameters) GoString() string {
  return s.String()
}

type QueryEdgeByteHitRatioServiceRequestHeader struct {
}

func (s QueryEdgeByteHitRatioServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeByteHitRatioServiceResponseHeader struct {
}

func (s QueryEdgeByteHitRatioServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUpFlowDomainCountryServiceRequest struct {
  // {"en":"Starting time
  // 
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2. Cannot be greater than the current time
  // 3. Get up to the last six months (183 days) of data.","zh_CN":"开始时间
  // 1. 时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2. 不能大于当前时间
  // 3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 1. Time format 2016-12-02T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception
  // 4. Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo can&rsquo;t exceed 7 days (can contact technical support adjustment, up to 31 days).","zh_CN":"结束时间:
  // 1. 时间格式2016-12-02T10:00:00+08:00
  // 2. 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3. dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4. 允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最长31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains: 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment); 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).","zh_CN":"域名： 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)； 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. Support 5m (5 minutes granularity),1d (1 day granularity)
  // 2. Do not pass the default to 5m","zh_CN":"数据粒度:
  // 1. 支持5m(5分钟粒度),1d(天粒度)
  // 2. 不传默认为5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Country code:
  // 
  // 1. Do not pass the default query for all countries and regions;
  // 2. The values that can be passed are detailed in the Countrycode list of appendix table on the API Overview page","zh_CN":"国家地区代号:
  // 1. 不传默认查询全部国家地区;
  // 2. 可传递的值详见概览页国家地区列表说明。"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Grouping dimension:
  // 
  // 1. Optional values domain, country and aggregatedOversea, can pass in single or multiple values. The aggregatedOversea and country cannot be passed at the same time;
  // 
  // 2. If there is an incoming, the detailed data will be displayed according to the dimension:
  // domain: Group display by domain name dimension;
  // country: Group display by country dimension;
  // aggregatedOversea: Group display according to domestic and overseas dimensions.
  // 3. The result hierarchy is fixed in order, and the order of the parameters does not affect the order of the returned results. For example: 'groupBy': ['domain','country'] and 'groupBy': ['country','domain'] return the same result.","zh_CN":"1. 可选值domain、country、aggregatedOversea,可传入单个或多个值,其中不能同时传aggregatedOversea 和 country;2. 有传入则按照该维度展示明细数据:1.domain:按照域名维度进行分组展示;2.domain:country:按照国家维度进行分组展示;3.aggregatedOversea:按照国内 和 海外维度进行分组展示3. 返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:groupBy: [domain,country]与groupBy: [country,domain]返回结果一样。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportUpFlowDomainCountryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUpFlowDomainCountryServiceRequest) SetDateFrom(v string) *ReportUpFlowDomainCountryServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceRequest) SetDateTo(v string) *ReportUpFlowDomainCountryServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceRequest) SetDomain(v []*string) *ReportUpFlowDomainCountryServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUpFlowDomainCountryServiceRequest) SetDataInterval(v string) *ReportUpFlowDomainCountryServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceRequest) SetCountryCode(v []*string) *ReportUpFlowDomainCountryServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportUpFlowDomainCountryServiceRequest) SetGroupBy(v []*string) *ReportUpFlowDomainCountryServiceRequest {
  s.GroupBy = v
  return s
}

type ReportUpFlowDomainCountryServiceRequestHeader struct {
}

func (s ReportUpFlowDomainCountryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUpFlowDomainCountryServicePaths struct {
}

func (s ReportUpFlowDomainCountryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServicePaths) GoString() string {
  return s.String()
}

type ReportUpFlowDomainCountryServiceParameters struct {
}

func (s ReportUpFlowDomainCountryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceParameters) GoString() string {
  return s.String()
}

type ReportUpFlowDomainCountryServiceResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"结果"}
  Data []*ReportUpFlowDomainCountryServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ReportUpFlowDomainCountryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUpFlowDomainCountryServiceResponse) SetCode(v string) *ReportUpFlowDomainCountryServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponse) SetData(v []*ReportUpFlowDomainCountryServiceResponseData) *ReportUpFlowDomainCountryServiceResponse {
  s.Data = v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponse) SetMessage(v string) *ReportUpFlowDomainCountryServiceResponse {
  s.Message = &v
  return s
}

type ReportUpFlowDomainCountryServiceResponseData struct     {
  // {"en":"","zh_CN":""}
  CountryData []*ReportUpFlowDomainCountryServiceResponseDataCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s ReportUpFlowDomainCountryServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportUpFlowDomainCountryServiceResponseData) SetCountryData(v []*ReportUpFlowDomainCountryServiceResponseDataCountryData) *ReportUpFlowDomainCountryServiceResponseData {
  s.CountryData = v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseData) SetDomain(v string) *ReportUpFlowDomainCountryServiceResponseData {
  s.Domain = &v
  return s
}

type ReportUpFlowDomainCountryServiceResponseDataCountryData struct     {
  // {"en":"National regional traffic ratio: the proportion (percentage) of traffic value in a single country region during the query period, retaining 2 decimal places","zh_CN":"国家地区流量占比:单个国家地区流量在查询时段内的流量值的占比(百分比),保留2位小数"}
  FlowPercentage *string `json:"flowPercentage,omitempty" xml:"flowPercentage,omitempty" require:"true"`
  // {"en":"Country code","zh_CN":"国家地区代号"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Summary of traffic in national regions: Summary of traffic flow in a single country region during the query period, unit of measure MB, retaining 2 decimal places","zh_CN":"国家地区流量汇总:单个国家地区流量在查询时段内的流量汇总值,计量单位MB,保留2位小数"}
  FlowSum *string `json:"flowSum,omitempty" xml:"flowSum,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  FlowData []*ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
  // {"en":"Country name","zh_CN":"国家地区名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
}

func (s ReportUpFlowDomainCountryServiceResponseDataCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceResponseDataCountryData) GoString() string {
  return s.String()
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryData) SetFlowPercentage(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryData {
  s.FlowPercentage = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryData) SetCountryCode(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryData) SetFlowSum(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryData {
  s.FlowSum = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryData) SetFlowData(v []*ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData) *ReportUpFlowDomainCountryServiceResponseDataCountryData {
  s.FlowData = v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryData) SetCountryName(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryData {
  s.CountryName = &v
  return s
}

type ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData struct     {
  // {"en":"Bandwidth value. Unit is Mbps and 2 digits of decimals are allowed.","zh_CN":"带宽值,单位Mbps,保留2位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Flow value, in megabytes, 2 decimal places","zh_CN":"流量值,计量单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. Returns the time slice contained in the start time and end time.","zh_CN":"1. 查询的数据粒度为5m时,格式为yyyy-MM-dd  HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。2. 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
}

func (s ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData) SetBandwidth(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData {
  s.Bandwidth = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData) SetValue(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData {
  s.Value = &v
  return s
}

func (s *ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData) SetTimestamp(v string) *ReportUpFlowDomainCountryServiceResponseDataCountryDataFlowData {
  s.Timestamp = &v
  return s
}

type ReportUpFlowDomainCountryServiceResponseHeader struct {
}

func (s ReportUpFlowDomainCountryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUpFlowDomainCountryServiceResponseHeader) GoString() string {
  return s.String()
}




type FlowTypeRequest struct {
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not Specifies,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号“;”分隔开，如查询大陆及亚太区域，参数填写为：“region=cn;apac”。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"domains that been queried:\n1)If there are multiple inputs,use  ';' as separator.\n2)If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. \n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"acceleration type.\n1)If there are multiple inputs,use ';' as separator.\n2)If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. \n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:\n1)'true' as default.\n2) If not 'true',it will query data of channels that ends with any item of input 'channel's.","zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
}

func (s FlowTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s FlowTypeRequest) GoString() string {
  return s.String()
}

func (s *FlowTypeRequest) SetDate(v string) *FlowTypeRequest {
  s.Date = &v
  return s
}

func (s *FlowTypeRequest) SetRegion(v string) *FlowTypeRequest {
  s.Region = &v
  return s
}

func (s *FlowTypeRequest) SetChannel(v string) *FlowTypeRequest {
  s.Channel = &v
  return s
}

func (s *FlowTypeRequest) SetEnddate(v string) *FlowTypeRequest {
  s.Enddate = &v
  return s
}

func (s *FlowTypeRequest) SetAccetype(v string) *FlowTypeRequest {
  s.Accetype = &v
  return s
}

func (s *FlowTypeRequest) SetStartdate(v string) *FlowTypeRequest {
  s.Startdate = &v
  return s
}

func (s *FlowTypeRequest) SetDataformat(v string) *FlowTypeRequest {
  s.Dataformat = &v
  return s
}

func (s *FlowTypeRequest) SetIsExactMatch(v string) *FlowTypeRequest {
  s.IsExactMatch = &v
  return s
}

type FlowTypeRequestHeader struct {
}

func (s FlowTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowTypeRequestHeader) GoString() string {
  return s.String()
}

type FlowTypePaths struct {
}

func (s FlowTypePaths) String() string {
  return tea.Prettify(s)
}

func (s FlowTypePaths) GoString() string {
  return s.String()
}

type FlowTypeParameters struct {
}

func (s FlowTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s FlowTypeParameters) GoString() string {
  return s.String()
}

type FlowTypeResponse struct {
  // {"en":"startdate,with format yyyy-mm-dd","zh_CN":"开始时间，yyyy-mm-dd"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate,with format yyyy-mm-dd","zh_CN":"结束时间，yyyy-mm-dd"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"domain name.","zh_CN":"频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"type code","zh_CN":"类型代码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"type name","zh_CN":"类型名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type English name","zh_CN":"英文名"}
  Enname *string `json:"enname,omitempty" xml:"enname,omitempty" require:"true"`
  // {"en":"Percentage of traffic","zh_CN":"流量占比"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s FlowTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowTypeResponse) GoString() string {
  return s.String()
}

func (s *FlowTypeResponse) SetStartdate(v string) *FlowTypeResponse {
  s.Startdate = &v
  return s
}

func (s *FlowTypeResponse) SetEnddate(v string) *FlowTypeResponse {
  s.Enddate = &v
  return s
}

func (s *FlowTypeResponse) SetChannel(v string) *FlowTypeResponse {
  s.Channel = &v
  return s
}

func (s *FlowTypeResponse) SetCode(v string) *FlowTypeResponse {
  s.Code = &v
  return s
}

func (s *FlowTypeResponse) SetName(v string) *FlowTypeResponse {
  s.Name = &v
  return s
}

func (s *FlowTypeResponse) SetEnname(v string) *FlowTypeResponse {
  s.Enname = &v
  return s
}

func (s *FlowTypeResponse) SetText(v string) *FlowTypeResponse {
  s.Text = &v
  return s
}

type FlowTypeResponseHeader struct {
}

func (s FlowTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowTypeResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficByIspProvinceForMultiDomainsBytesRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.The data can be obtained for the latest six months (183 days) at most, and the start time cannot be greater than the current time and end time.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n2.最多可获取最近半年（183天）的数据，并且开始时间不能大于当前时间和结束时间"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n2. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception\n3. The default query span is up to 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (you can contact technical support for adjustment, up to 31 days);","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常  \n3.默认允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整，最大31天）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"The domain name and the number of domain names are adjustable according to the account number. The default is 20","zh_CN":"域名，域名个数限制根据账号可调，默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"5m","en":"Data granularity, 5m: 5 minutes granularity, 1h: 1 hour granularity","zh_CN":"数据粒度，5m：5分钟粒度，1h：1小时粒度","exampleValue":"5m,1h"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province:\n1.If no province  is specified: Query all provinces and aggregate the returned data according to all provinces\n2.Province specified: Send province code, multiple codes can be sent.","zh_CN":"省份：\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份传code，可传多个。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1. If no ISP is specified: all ISPs will be queried, and the returned data will be aggregated by all ISPs\n2. ISP specified: Send ISP code, multiple codes can be provided","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合\n2.有传递isp时：传递运营商code，可传多个"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension\n1.Options are domain, province, isp, and more than one value can be entered;\n2.The data is displayed according to the specified dimension;","zh_CN":"分组维度\n1.可选值为domain、province、isp,可传入多个值;\n2.有传入则按照该维度展示明细数据;","exampleValue":"domain,province,isp"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetDateFrom(v string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetDateTo(v string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetDomain(v []*string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetGranularity(v string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetProvince(v []*string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.Province = v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetIsp(v []*string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesRequest) SetGroupBy(v []*string) *GetTrafficByIspProvinceForMultiDomainsBytesRequest {
  s.GroupBy = v
  return s
}

type GetTrafficByIspProvinceForMultiDomainsBytesRequestHeader struct {
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficByIspProvinceForMultiDomainsBytesPaths struct {
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesPaths) GoString() string {
  return s.String()
}

type GetTrafficByIspProvinceForMultiDomainsBytesParameters struct {
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesParameters) GoString() string {
  return s.String()
}

type GetTrafficByIspProvinceForMultiDomainsBytesResponse struct {
  // {"en":"","zh_CN":""}
  Result []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponse) SetResult(v []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResult) *GetTrafficByIspProvinceForMultiDomainsBytesResponse {
  s.Result = v
  return s
}

type GetTrafficByIspProvinceForMultiDomainsBytesResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  IspData []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResult) SetDomain(v string) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResult) SetIspData(v []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResult {
  s.IspData = v
  return s
}

type GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData struct     {
  // {"en":"Internet service providers","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  ProvinceData []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData) SetIsp(v string) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData) SetProvinceData(v []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspData {
  s.ProvinceData = v
  return s
}

type GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData struct     {
  // {"en":"province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData) SetProvince(v string) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData) SetDataSeries(v []*GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceData {
  s.DataSeries = v
  return s
}

type GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries struct     {
  // {"en":"time\n1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.\n2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00.\n3. Return to the time slice included in the start time and end time.","zh_CN":"时间\n1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1）00:00。\n2.查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1）00。\n3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:Unit:  Byte","zh_CN":"流量值：单位Byte"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries) SetTimestamp(v string) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries) SetTraffic(v string) *GetTrafficByIspProvinceForMultiDomainsBytesResponseResultIspDataProvinceDataDataSeries {
  s.Traffic = &v
  return s
}

type GetTrafficByIspProvinceForMultiDomainsBytesResponseHeader struct {
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByIspProvinceForMultiDomainsBytesResponseHeader) GoString() string {
  return s.String()
}




type QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);4.You can only query data for the last 2 years.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  DomainDir []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir `json:"domainDir,omitempty" xml:"domainDir,omitempty" require:"true" type:"Repeated"`
  // {"en":"Query protocol:
  // 1.Means query statistics with http protocol or with https protocol;
  // 2.Values can be selected: http or https;
  // 3.Empty means both http and https;", "zh_CN":"查询协议:
  // 1.代表统计http协议或https协议;
  // 2.可选值:http 或 https;
  // 3.不传默认代表不区分http和https;"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Query granularity:
  // 1.Empty value means 5 minutes of granularity by default;
  // 2.Values can be selected: 5m or 1m;
  // 3.Option of 1m can be configured in flowDirDataIntervalConfig depending on differrent accounts;", "zh_CN":"查询粒度:
  // 1.不传默认代表5分钟粒度;
  // 2.可选值:5m 或 1m;
  // 3.1m选项根据账号在数据字典flowDirDataIntervalConfig中可配;"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Query type
  // 
  // 1.Values can be selected: flow or bandwidth;
  // 
  // 2.bandwidth means to get bandwidth value, flow means to get the flow vaule. By default, data in all regions is queried;", "zh_CN":"查询类型
  // 1.可选值:flow 或bandwidth;
  // 2.bandwidth代表获取带宽值,flow代表获取流量值,默认查询所有区域的数据;"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty" require:"true"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDateFrom(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDateTo(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDomainDir(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DomainDir = v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetProtocolType(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDataInterval(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDataType(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DataType = &v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir struct     {
  // {"en":"Domain:
  // 1.Need to meet the regular expression rules that are used to validate domains;
  // 2.Domain number limits can be adjusted depending on different accounts. The default value is 20;", "zh_CN":"域名:
  // 1.需要满足域名的正则校验;
  // 2.域名个数限制根据账号可调,默认为20个;"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Table of contents:
  // 1.Empty value means to query all directories under the domain
  // 2.Directory limits can be adjusted depending on different accounts. The default value is 200(this limit applies to the empty value);
  // 3.Invalid directories are not returned", "zh_CN":"目录:
  // 1.目录个数限制根据账号可调,默认为200个;
  // 2.不传代表查询该域名下的所有目录,同时接受目录个数限制;
  // 3.无效的目录不返回"}
  Dir []*string `json:"dir,omitempty" xml:"dir,omitempty" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) SetDomain(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir {
  s.Domain = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) SetDir(v []*string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir {
  s.Dir = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) SetResult(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse {
  s.Result = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) SetDomain(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) SetDetails(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult {
  s.Details = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails struct     {
  // {"en":"Directory", "zh_CN":"目录"}
  Dir *string `json:"dir,omitempty" xml:"dir,omitempty" require:"true"`
  // {"en":"Return when the dataType is flow; Total flow of every stream. Unit is MB and 2 digits of decimals allowed;", "zh_CN":"当dataType为flow时返回;每路流的总流量单位MB,保留2位小数;"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Return when the dataType is bandwidth;
  // 
  // Bandwidth peak value of every stream within specified time. Unit is Mbps, two decimals digits;", "zh_CN":"当dataType为bandwidth时返回;
  // 每路流在该时间段内的带宽峰值单位Mbps,保留2位小数;"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty" require:"true"`
  Details []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetDir(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.Dir = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetTotalFlow(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetBandwidthPeakValue(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetDetails(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.Details = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails struct     {
  // {"en":"Date
  // 1.When the data query granularity is 1m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 3.Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间
  // 1.查询的数据粒度为1m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:01,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 2.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Data of the time point, two  decimal digits", "zh_CN":"时间点的数据,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) SetTimestamp(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) SetValue(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails {
  s.Value = &v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainPaths struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainPaths) GoString() string {
  return s.String()
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainParameters struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainParameters) GoString() string {
  return s.String()
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestHeader struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseHeader struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryTotalTrafficForMultiDomainsRequest struct {
  // {"en":"Domain list.\nDomain number limits can be adjusted depending on different accounts. The default value is 1000(if you want to adjust,please, contact technical support)","zh_CN":"域名列表\n1.域名个数限制根据账号可调,默认为1000个(可联系技术支持下单调整);"}
  DomainList *QueryTotalTrafficForMultiDomainsRequestDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" type:"Struct"`
}

func (s QueryTotalTrafficForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsRequest) SetDomainList(v *QueryTotalTrafficForMultiDomainsRequestDomainList) *QueryTotalTrafficForMultiDomainsRequest {
  s.DomainList = v
  return s
}

type QueryTotalTrafficForMultiDomainsRequestDomainList struct {
  // {"en":"Domain","zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" type:"Repeated"`
}

func (s QueryTotalTrafficForMultiDomainsRequestDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsRequestDomainList) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsRequestDomainList) SetDomainName(v []*string) *QueryTotalTrafficForMultiDomainsRequestDomainList {
  s.DomainName = v
  return s
}

type QueryTotalTrafficForMultiDomainsRequestHeader struct {
}

func (s QueryTotalTrafficForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type QueryTotalTrafficForMultiDomainsPaths struct {
}

func (s QueryTotalTrafficForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsPaths) GoString() string {
  return s.String()
}

type QueryTotalTrafficForMultiDomainsParameters struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.And smaller than the current time and 'dateTo';\n3.Period between 'dataFrom' and 'dateTo' cannot be longer than 31 days","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.并且小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过31天;4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be greater than 'dateFrom';\n3.If it's greater than the current time, then the current time is assigned as the value","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于dateFrom;\n3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"defaultValue":"daily","en":"Data granularity\n1.fiveminutes: five minutes, hourly: one hour, daily: one day;\n2.If not specified, daily is set as the default value;\n3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.","zh_CN":"数据粒度\n1.fiveminutes:5分钟,hourly:1小时,daily:1天;\n2.不传递,默认为daily;\n3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。","exampleValue":"fiveminutes,hourly,daily"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
}

func (s QueryTotalTrafficForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsParameters) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsParameters) SetDateFrom(v string) *QueryTotalTrafficForMultiDomainsParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsParameters) SetDateTo(v string) *QueryTotalTrafficForMultiDomainsParameters {
  s.DateTo = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsParameters) SetGranularity(v string) *QueryTotalTrafficForMultiDomainsParameters {
  s.Granularity = &v
  return s
}

type QueryTotalTrafficForMultiDomainsResponse struct {
  // {"en":"Total traffic. Unit:  MB, retain two decimals","zh_CN":"总流量, 单位MB，保留2位小数"}
  TotalTraffic *int `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"dataSeries","zh_CN":"流量数据"}
  DataSeries []*QueryTotalTrafficForMultiDomainsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalTrafficForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsResponse) SetTotalTraffic(v int) *QueryTotalTrafficForMultiDomainsResponse {
  s.TotalTraffic = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsResponse) SetDataSeries(v []*QueryTotalTrafficForMultiDomainsResponseDataSeries) *QueryTotalTrafficForMultiDomainsResponse {
  s.DataSeries = v
  return s
}

type QueryTotalTrafficForMultiDomainsResponseDataSeries struct     {
  // {"en":"Date\n1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.\n2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.\n3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.Return the time slice contained in start time and the time slice contained in end time","zh_CN":"时间\n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit:  MB, retain two decimals","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s QueryTotalTrafficForMultiDomainsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsResponseDataSeries) SetTimestamp(v string) *QueryTotalTrafficForMultiDomainsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsResponseDataSeries) SetTraffic(v string) *QueryTotalTrafficForMultiDomainsResponseDataSeries {
  s.Traffic = &v
  return s
}

type QueryTotalTrafficForMultiDomainsResponseHeader struct {
}

func (s QueryTotalTrafficForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type QueryStreamTrafficUnderMALBDomainRequest struct {
  // {"en":"Starting time
  // 			1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 			2. Must be less than the current time and dateTo;
  // 			3. The difference between dateFrom and dateTo cannot exceed 1 day(technical support can be contacted to adjust);
  // 			4. Only data within the last 6 months can be queried.", "zh_CN":"开始时间
  // 			1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 			2.必须小于当前时间和dateTo；
  // 			3.dateFrom和dateTo相差不能超过1天(可联系技术支持调整)；
  // 			4.只能查询最近6个月内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 			1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 			2. Must be greater than dateFrom;
  // 			3. If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 			1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 			2.必须大于dateFrom；
  // 			3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain and stream name list", "zh_CN":"域名和流名组：
  // 						1.可传递的域名和流名组数量上限默认为20组(可联系技术支持调整)；
  // 						2.域名domain：一组域名和流名组中只能传递单个域名，且域名必须传递；
  // 						3.流名stream：一组域名和流名组下(即单个域名下)可传递的流名数量上限默认为2000个(可联系技术支持调整)，流名未传递时默认查询域名下所有流名，但当域名下流名数量超过限制时不可查询(报错)。"}
  DomainStream []*QueryStreamTrafficUnderMALBDomainRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 				1. 1m: 1-minute granularity, 5m: 5-minute granularity;
  // 				2. Data with granularity of 5 minutes is queried by default, if you need to query the granularity of 1 minute, please contact technical support for configuration", "zh_CN":"数据粒度
  // 				1.1m：1分钟粒度，5m：5分钟粒度；
  // 				2.默认查询5分钟粒度数据，若需要查询1m粒度请联系技术支持进行特殊配置"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Data type: flow, bandwidth", "zh_CN":"数据类型，flow：流量，bandwidth：带宽"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty" require:"true"`
  // {"en":"Filter nullname stream:
  // 			1.Enter optional value '0' or '1';
  // 			2. The parameter 0 does not filter the data whose data is null, and the parameter 1 filter the data whose data is null.
  // 			3. the default value is 0;", "zh_CN":"是否过滤空流名:
  // 			1.入参可选值 '0' 或 '1' ；
  // 			2.传参 0 不过滤流名为空的数据，传参 1 则过滤流名为空的数据 ；
  // 			3.默认值为 0 ；"}
  FilterEmptyStream *int `json:"filterEmptyStream,omitempty" xml:"filterEmptyStream,omitempty"`
}

func (s QueryStreamTrafficUnderMALBDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDateFrom(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDateTo(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDomainStream(v []*QueryStreamTrafficUnderMALBDomainRequestDomainStream) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DomainStream = v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDataInterval(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDataType(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DataType = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetFilterEmptyStream(v int) *QueryStreamTrafficUnderMALBDomainRequest {
  s.FilterEmptyStream = &v
  return s
}

type QueryStreamTrafficUnderMALBDomainRequestDomainStream struct     {
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"可传递域名数量上限默认为20个(可联系技术支持调整)"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Stream name:
  // 				1. If this field is not specified, it means all domains under the domain are queried;
  // 				2. Number of streams can be adjusted depending on different accounts. The default value is 2000(this limit applies to the empty value);", "zh_CN":"流名：'发布点'+'流名'。例如：live/test-20180101-test ,其中live是发布点，test-20180101-test是流名。"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainRequestDomainStream) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainRequestDomainStream) SetDomain(v string) *QueryStreamTrafficUnderMALBDomainRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequestDomainStream) SetStream(v []*string) *QueryStreamTrafficUnderMALBDomainRequestDomainStream {
  s.Stream = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponse struct {
  Result []*QueryStreamTrafficUnderMALBDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponse) SetResult(v []*QueryStreamTrafficUnderMALBDomainResponseResult) *QueryStreamTrafficUnderMALBDomainResponse {
  s.Result = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryStreamTrafficUnderMALBDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResult) SetDomain(v string) *QueryStreamTrafficUnderMALBDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResult) SetDetails(v []*QueryStreamTrafficUnderMALBDomainResponseResultDetails) *QueryStreamTrafficUnderMALBDomainResponseResult {
  s.Details = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponseResultDetails struct     {
  // {"en":"Stream name", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"Total traffic:
  // 					1. Keep two digits of decimals. Unit: MB;
  // 					2. Return when the input parameter of datatype is flow.", "zh_CN":"总流量
  // 					1.保留2位小数，单位为MB;
  // 					2.当入参dataType为flow时，返回。"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty"`
  // {"en":"Peak value of bandwidth
  // 					1. Keep two digits of decimals. Unit: Mbps
  // 					2. Return when the input dataType is bandwidth.", "zh_CN":"峰值带宽
  // 					1.保留2位小数，单位Mbps
  // 					2.当入参dataType为bandwidth时，返回。"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty"`
  Details []*QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetStream(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.Stream = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetTotalFlow(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetBandwidthPeakValue(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetDetails(v []*QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.Details = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails struct     {
  // {"en":"Date:
  // 						1. When the data query granularity is 1m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00.
  // 						2. When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间
  // 						1.查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。
  // 						2.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Data of the time point:
  // 						1. Two decimal digits allowed;
  // 						2. When the input parameter of dataType is flow, the value is the flow/traffic value, with MB as the unit;
  // 						3. When the input parameter of dataType is bandwidth, the value is the bandwidth value, with Mbps as the unit;", "zh_CN":"时间点的数据
  // 						1.保留2位小数；
  // 						2.当入参dataType为flow时，value值为流量，单位为MB；
  // 						3.当入参dataType为bandwidth时，value值为带宽值，单位为Mbps；"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) SetTimestamp(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) SetValue(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails {
  s.Value = &v
  return s
}

type QueryStreamTrafficUnderMALBDomainPaths struct {
}

func (s QueryStreamTrafficUnderMALBDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainPaths) GoString() string {
  return s.String()
}

type QueryStreamTrafficUnderMALBDomainParameters struct {
}

func (s QueryStreamTrafficUnderMALBDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainParameters) GoString() string {
  return s.String()
}

type QueryStreamTrafficUnderMALBDomainRequestHeader struct {
}

func (s QueryStreamTrafficUnderMALBDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryStreamTrafficUnderMALBDomainResponseHeader struct {
}

func (s QueryStreamTrafficUnderMALBDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficByBillingRegionForMultiDomainsRequest struct {
  // {"en":"Start time\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.And smaller than the current time and dateTo;\n3.Period between dataFrom and dateTo cannot be longer than 31 days;4.You can only query data for the last 2 years(technical support can be contacted to adjust).","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.并且小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.Must be greater than dateFrom;\n3.If it's greater than the current time, then the current time is assigned as the value;","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于dateFrom;\n3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain name list, if not passed, all domain names will be queried, the number is unlimited; when domain names are passed, the default limit is 1000 (you can contact technical support to place an order to adjust)","zh_CN":"域名列表,不传递则查询全部域名,数量数量不受限; 有传递域名时,默认限制为1000个(可联系技术支持下单调整)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"1d","en":"Data granularity\n\n1.Options: 5m(5 minutes), 1h(1 hour) and 1d(1 day);\n2.Default value of 1d is used if the field is not specified;\n3.If 5m is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.","zh_CN":"数据粒度\n1.可选值为:5m(5分钟)、1h(1小时)、1d(1天);\n2.不传时默认为1d;\n3.传递5m时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。","exampleValue":"5m,1h,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"Queried area\n1.Use  English semicolon to separate two areas;\n2.Options: cn, nc, ov, apac, am, euna, emea, sa, af, au, hk, tw\n...;","zh_CN":"查询区域\n1.多个区域使用英文分号分隔\n2.可选值为:cn、nc、ov、apac、am、euna、emea、sa、af、au、hk、tw等;"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
}

func (s GetTrafficByBillingRegionForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficByBillingRegionForMultiDomainsRequest) SetDateFrom(v string) *GetTrafficByBillingRegionForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficByBillingRegionForMultiDomainsRequest) SetDateTo(v string) *GetTrafficByBillingRegionForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficByBillingRegionForMultiDomainsRequest) SetDomain(v []*string) *GetTrafficByBillingRegionForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficByBillingRegionForMultiDomainsRequest) SetGranularity(v string) *GetTrafficByBillingRegionForMultiDomainsRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficByBillingRegionForMultiDomainsRequest) SetAreaCode(v string) *GetTrafficByBillingRegionForMultiDomainsRequest {
  s.AreaCode = &v
  return s
}

type GetTrafficByBillingRegionForMultiDomainsRequestHeader struct {
}

func (s GetTrafficByBillingRegionForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficByBillingRegionForMultiDomainsPaths struct {
}

func (s GetTrafficByBillingRegionForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetTrafficByBillingRegionForMultiDomainsParameters struct {
}

func (s GetTrafficByBillingRegionForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetTrafficByBillingRegionForMultiDomainsResponse struct {
  // {"en":"Total traffic: \nUnit:  MB, retain two decimals","zh_CN":"总流量, 单位MB，保留2位小数"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficByBillingRegionForMultiDomainsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByBillingRegionForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficByBillingRegionForMultiDomainsResponse) SetTotalTraffic(v string) *GetTrafficByBillingRegionForMultiDomainsResponse {
  s.TotalTraffic = &v
  return s
}

func (s *GetTrafficByBillingRegionForMultiDomainsResponse) SetDataSeries(v []*GetTrafficByBillingRegionForMultiDomainsResponseDataSeries) *GetTrafficByBillingRegionForMultiDomainsResponse {
  s.DataSeries = v
  return s
}

type GetTrafficByBillingRegionForMultiDomainsResponseDataSeries struct     {
  // {"en":"time\n1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH: MM; Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is yyyy-MM-dd 24:00.\n2. When the data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 01, and the last time slice is yyyy-MM-dd 24.\n3. when the data granularity of query is daily, the format is yyyy-MM-dd; The value of data for each time slice represents the value of the data within that day;\n4. return the time slice contained in the start and end times.","zh_CN":"时间\n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit:  MB, retain two decimals","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s GetTrafficByBillingRegionForMultiDomainsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficByBillingRegionForMultiDomainsResponseDataSeries) SetTimestamp(v string) *GetTrafficByBillingRegionForMultiDomainsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficByBillingRegionForMultiDomainsResponseDataSeries) SetTraffic(v string) *GetTrafficByBillingRegionForMultiDomainsResponseDataSeries {
  s.Traffic = &v
  return s
}

type GetTrafficByBillingRegionForMultiDomainsResponseHeader struct {
}

func (s GetTrafficByBillingRegionForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByBillingRegionForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type FlowDayRequest struct {
  // {"en":"main account cust_en_name, cust_en_name of sub-client.\nWhen a merged-account wants to  view the information of the subclient,the cust_en_name is required.","zh_CN":"主账号客户英文名，合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not Specifies,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope.\n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope.\n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1)If there are multiple inputs,use  ';' as separator.\n2)If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:\n1)'true' as default.\n2) If not 'true',it will query data of channels that ends with any item of input 'channel's.","zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions. Common regions: cn: Mainland China, hk: Hong Kong, ov: Oversea average, tw: Taiwan, euna: Europe and North America, apac: Asia-Pacific, sa: South America, af: Africa, am: Americas, emea: Europe/Middle East/Africa, kr: South Korea, au: Australia, in: India, jp: Japan, ru: Russia, indo: Indonesia, me: Middle East, eu: Europe, ph: Philippines","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。常用区域：cn:中国大陆,hk:香港,ov:海外平均,tw:台湾,euna:欧美,apac:亚太,sa:南美,af:非洲,am:美洲,emea:欧洲/中东/非洲,kr:韩国,au:澳大利亚,in:印度,jp:日本,ru:俄罗斯,indo:印尼,me:中东,eu:欧洲,ph:菲律宾"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"1)If there isp multiple inputs,use ';' as demimeter.\n2)optional values of isp: refers to the ISP-section of appendix.\n3) If not specified,means all the isp.","zh_CN":"要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"acceleration type.\n1)If there are multiple inputs,use ';' as separator.\n2)If not specified or specified as 'all', it means all the accetypes. Common Acceleration Types: Web Acceleration: web, Web-HTTPS: web-https, Whole Site Acceleration: wsa, Whole Site Acceleration-HTTPS: wsa-https, HTTP Download: download, Livestream Media: livestream, Livestream-HTTPS: live-https, VOD-HTTPS: vod-https, Cloud Video on Demand: cloudv-vod, VOD Streaming: vodstream, Financial Security Acceleration Solution: fsa, Government and Enterprise Security Acceleration Solution: gess, E-commerce Security Acceleration Solution: esa, Mobile Acceleration (MAA): maa, Application Security Acceleration Solution: s-appa, Application Acceleration: appa, WAF: waf, Upload Acceleration: upload, Upload Acceleration-HTTPS: upa-https, IPv6 Integration Solution: osv6, Live P2P: livep2p","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型. 常用加速类型：网页加速:web,网页-HTTPS:web-https ,全站加速:wsa,全站加速-https:wsa-https,HTTP下载:download,流媒体直播:livestream,直播-https:live-https,点播-HTTPS:vod-https,云点播:cloudv-vod,流媒体点播:vodstream,金融安全加速解决方案:fsa,政企安全加速解决方案:gess,电商安全加速解决方案:esa,移动加速(MAA):maa,应用安全加速解决方案:s-appa,应用加速:appa,WAF:waf,上传加速:upload,上传加速-https:upa-https,IPv6一体化解决方案:osv6,直播P2P:livep2p"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Greenwich time zone, the parameter format GMT+09:00 means East Nine District, GMT-09:00 means West Nine District, if not passed, the default is the local time zone (East Eight District)","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Display statistic result in merged or separate way. 1.If specified 1,get the merged result. 2.If specified 2,get the separate result. 3.If specified 3,get both merged result and separate result. 4.If not specified,means '1'.","zh_CN":"结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为“1”"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s FlowDayRequest) String() string {
  return tea.Prettify(s)
}

func (s FlowDayRequest) GoString() string {
  return s.String()
}

func (s *FlowDayRequest) SetCust(v string) *FlowDayRequest {
  s.Cust = &v
  return s
}

func (s *FlowDayRequest) SetDate(v string) *FlowDayRequest {
  s.Date = &v
  return s
}

func (s *FlowDayRequest) SetStartdate(v string) *FlowDayRequest {
  s.Startdate = &v
  return s
}

func (s *FlowDayRequest) SetEnddate(v string) *FlowDayRequest {
  s.Enddate = &v
  return s
}

func (s *FlowDayRequest) SetChannel(v string) *FlowDayRequest {
  s.Channel = &v
  return s
}

func (s *FlowDayRequest) SetIsExactMatch(v string) *FlowDayRequest {
  s.IsExactMatch = &v
  return s
}

func (s *FlowDayRequest) SetRegion(v string) *FlowDayRequest {
  s.Region = &v
  return s
}

func (s *FlowDayRequest) SetIsp(v string) *FlowDayRequest {
  s.Isp = &v
  return s
}

func (s *FlowDayRequest) SetAccetype(v string) *FlowDayRequest {
  s.Accetype = &v
  return s
}

func (s *FlowDayRequest) SetDataformat(v string) *FlowDayRequest {
  s.Dataformat = &v
  return s
}

func (s *FlowDayRequest) SetTimezone(v string) *FlowDayRequest {
  s.Timezone = &v
  return s
}

func (s *FlowDayRequest) SetResultType(v string) *FlowDayRequest {
  s.ResultType = &v
  return s
}

type FlowDayRequestHeader struct {
}

func (s FlowDayRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowDayRequestHeader) GoString() string {
  return s.String()
}

type FlowDayPaths struct {
}

func (s FlowDayPaths) String() string {
  return tea.Prettify(s)
}

func (s FlowDayPaths) GoString() string {
  return s.String()
}

type FlowDayParameters struct {
}

func (s FlowDayParameters) String() string {
  return tea.Prettify(s)
}

func (s FlowDayParameters) GoString() string {
  return s.String()
}

type FlowDayResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *FlowDayResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s FlowDayResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponse) GoString() string {
  return s.String()
}

func (s *FlowDayResponse) SetProvider(v *FlowDayResponseProvider) *FlowDayResponse {
  s.Provider = v
  return s
}

type FlowDayResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"数据"}
  Date *FlowDayResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s FlowDayResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponseProvider) GoString() string {
  return s.String()
}

func (s *FlowDayResponseProvider) SetName(v string) *FlowDayResponseProvider {
  s.Name = &v
  return s
}

func (s *FlowDayResponseProvider) SetType(v string) *FlowDayResponseProvider {
  s.Type = &v
  return s
}

func (s *FlowDayResponseProvider) SetDate(v *FlowDayResponseProviderDate) *FlowDayResponseProvider {
  s.Date = v
  return s
}

type FlowDayResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始日期"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束日期"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"channel","zh_CN":"频道"}
  Channel *FlowDayResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s FlowDayResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponseProviderDate) GoString() string {
  return s.String()
}

func (s *FlowDayResponseProviderDate) SetStartdate(v string) *FlowDayResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *FlowDayResponseProviderDate) SetEnddate(v string) *FlowDayResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *FlowDayResponseProviderDate) SetChannel(v *FlowDayResponseProviderDateChannel) *FlowDayResponseProviderDate {
  s.Channel = v
  return s
}

type FlowDayResponseProviderDateChannel struct {
  // {"en":"channel","zh_CN":"频道"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"channel-peak","zh_CN":"频道峰值数据"}
  ChannelPeak []*FlowDayResponseProviderDateChannelChannelPeak `json:"channel-peak,omitempty" xml:"channel-peak,omitempty" require:"true" type:"Repeated"`
}

func (s FlowDayResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *FlowDayResponseProviderDateChannel) SetName(v string) *FlowDayResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *FlowDayResponseProviderDateChannel) SetChannelPeak(v []*FlowDayResponseProviderDateChannelChannelPeak) *FlowDayResponseProviderDateChannel {
  s.ChannelPeak = v
  return s
}

type FlowDayResponseProviderDateChannelChannelPeak struct     {
  // {"en":"date","zh_CN":"日期"}
  Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
  // {"en":"peakTime","zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"peakvalue(Mbps)","zh_CN":"带宽峰值（单位Mbps）"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"the total flow","zh_CN":"总流量"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"peak traffic value, unit: Byte, granularity: determined by the granularity enabled by the customer.","zh_CN":"流量峰值，单位：Byte，粒度：根据本身客户开启的粒度决定"}
  PeakValueByte *string `json:"peakValueByte,omitempty" xml:"peakValueByte,omitempty" require:"true"`
}

func (s FlowDayResponseProviderDateChannelChannelPeak) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponseProviderDateChannelChannelPeak) GoString() string {
  return s.String()
}

func (s *FlowDayResponseProviderDateChannelChannelPeak) SetDate(v string) *FlowDayResponseProviderDateChannelChannelPeak {
  s.Date = &v
  return s
}

func (s *FlowDayResponseProviderDateChannelChannelPeak) SetPeakTime(v string) *FlowDayResponseProviderDateChannelChannelPeak {
  s.PeakTime = &v
  return s
}

func (s *FlowDayResponseProviderDateChannelChannelPeak) SetPeakValue(v string) *FlowDayResponseProviderDateChannelChannelPeak {
  s.PeakValue = &v
  return s
}

func (s *FlowDayResponseProviderDateChannelChannelPeak) SetTotalFlow(v string) *FlowDayResponseProviderDateChannelChannelPeak {
  s.TotalFlow = &v
  return s
}

func (s *FlowDayResponseProviderDateChannelChannelPeak) SetPeakValueByte(v string) *FlowDayResponseProviderDateChannelChannelPeak {
  s.PeakValueByte = &v
  return s
}

type FlowDayResponseHeader struct {
}

func (s FlowDayResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowP2pShareRatioServiceRequest struct {
  // {"en":"Start date: 
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
  // {"en":"Data granularity: 
  // 	1.Support for 1m(1 minutes), 5m (5 minutes)
  // 	2.Default 5m
  // ", "zh_CN":"数据粒度:
  // 1、支持1m(1分钟)、5m(5分钟)
  // 2、不传默认5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"1. Selection: flash, android, ios, h5, windows 2.Default all platform types", "zh_CN":"1.可选值:flash、android、ios、h5、windows
  // 2.不传默认全部平台类型"}
  Platform []*string `json:"platform,omitempty" xml:"platform,omitempty" type:"Repeated"`
  // {"en":"1.Selection:domain 2.If groupBy left empty, merge date of all domains", "zh_CN":"可选值:domain
  // 不传默认聚合所有频道数据"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowP2pShareRatioServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowP2pShareRatioServiceRequest) SetDateFrom(v string) *ReportFlowP2pShareRatioServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceRequest) SetDateTo(v string) *ReportFlowP2pShareRatioServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceRequest) SetDomain(v []*string) *ReportFlowP2pShareRatioServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowP2pShareRatioServiceRequest) SetDataInterval(v string) *ReportFlowP2pShareRatioServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceRequest) SetPlatform(v []*string) *ReportFlowP2pShareRatioServiceRequest {
  s.Platform = v
  return s
}

func (s *ReportFlowP2pShareRatioServiceRequest) SetGroupBy(v []*string) *ReportFlowP2pShareRatioServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowP2pShareRatioServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportFlowP2pShareRatioServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowP2pShareRatioServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowP2pShareRatioServiceResponse) SetCode(v string) *ReportFlowP2pShareRatioServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceResponse) SetMessage(v string) *ReportFlowP2pShareRatioServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceResponse) SetData(v []*ReportFlowP2pShareRatioServiceResponseData) *ReportFlowP2pShareRatioServiceResponse {
  s.Data = v
  return s
}

type ReportFlowP2pShareRatioServiceResponseData struct     {
  // {"en":"Domain. If merge date of all domains will not return domain", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*ReportFlowP2pShareRatioServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowP2pShareRatioServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowP2pShareRatioServiceResponseData) SetDomain(v string) *ReportFlowP2pShareRatioServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceResponseData) SetDetailList(v []*ReportFlowP2pShareRatioServiceResponseDataDetailList) *ReportFlowP2pShareRatioServiceResponseData {
  s.DetailList = v
  return s
}

type ReportFlowP2pShareRatioServiceResponseDataDetailList struct     {
  // {"en":"timestamp,Returns the timestamp between the start time and end time. Time format:  yyyy-MM-dd HH:mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。
  // 时间格式:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"share rate, unit %", "zh_CN":"分享率,单位%"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowP2pShareRatioServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowP2pShareRatioServiceResponseDataDetailList) SetTimestamp(v string) *ReportFlowP2pShareRatioServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceResponseDataDetailList) SetValue(v string) *ReportFlowP2pShareRatioServiceResponseDataDetailList {
  s.Value = &v
  return s
}

type ReportFlowP2pShareRatioServicePaths struct {
}

func (s ReportFlowP2pShareRatioServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServicePaths) GoString() string {
  return s.String()
}

type ReportFlowP2pShareRatioServiceParameters struct {
}

func (s ReportFlowP2pShareRatioServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowP2pShareRatioServiceRequestHeader struct {
}

func (s ReportFlowP2pShareRatioServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowP2pShareRatioServiceResponseHeader struct {
}

func (s ReportFlowP2pShareRatioServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDataTransferForAllDomainByteRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 2 years at most can be queried.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.不能大于当前时间
  // 3.最多可获取最近2年内的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1day will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 31days, means the period between dateFrom to dateTo should not exceed 31days (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间,
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 4.允许查询最大时间间隔31天:,即dateFrom和dateTo相差不能超过31天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity, 5m: 5-minute   granularity, 1h: 1-hour granularity, 1d: 1-day granularity ", "zh_CN":"数据粒度, 5m:5分钟粒度,1h:1小时粒度, 1d: 1天粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryDataTransferForAllDomainByteRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteRequest) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteRequest) SetDateFrom(v string) *QueryDataTransferForAllDomainByteRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteRequest) SetDateTo(v string) *QueryDataTransferForAllDomainByteRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteRequest) SetDataInterval(v string) *QueryDataTransferForAllDomainByteRequest {
  s.DataInterval = &v
  return s
}

type QueryDataTransferForAllDomainByteResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryDataTransferForAllDomainByteResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDataTransferForAllDomainByteResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponse) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteResponse) SetCode(v string) *QueryDataTransferForAllDomainByteResponse {
  s.Code = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponse) SetMessage(v string) *QueryDataTransferForAllDomainByteResponse {
  s.Message = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponse) SetData(v []*QueryDataTransferForAllDomainByteResponseData) *QueryDataTransferForAllDomainByteResponse {
  s.Data = v
  return s
}

type QueryDataTransferForAllDomainByteResponseData struct     {
  // {"en":"flowSummary", "zh_CN":"总流量"}
  FlowSummary *string `json:"flowSummary,omitempty" xml:"flowSummary,omitempty" require:"true"`
  FlowData []*QueryDataTransferForAllDomainByteResponseDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDataTransferForAllDomainByteResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponseData) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteResponseData) SetFlowSummary(v string) *QueryDataTransferForAllDomainByteResponseData {
  s.FlowSummary = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponseData) SetFlowData(v []*QueryDataTransferForAllDomainByteResponseDataFlowData) *QueryDataTransferForAllDomainByteResponseData {
  s.FlowData = v
  return s
}

type QueryDataTransferForAllDomainByteResponseDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryDataTransferForAllDomainByteResponseDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponseDataFlowData) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteResponseDataFlowData) SetTimestamp(v string) *QueryDataTransferForAllDomainByteResponseDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponseDataFlowData) SetFlow(v string) *QueryDataTransferForAllDomainByteResponseDataFlowData {
  s.Flow = &v
  return s
}

type QueryDataTransferForAllDomainBytePaths struct {
}

func (s QueryDataTransferForAllDomainBytePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainBytePaths) GoString() string {
  return s.String()
}

type QueryDataTransferForAllDomainByteParameters struct {
}

func (s QueryDataTransferForAllDomainByteParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteParameters) GoString() string {
  return s.String()
}

type QueryDataTransferForAllDomainByteRequestHeader struct {
}

func (s QueryDataTransferForAllDomainByteRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteRequestHeader) GoString() string {
  return s.String()
}

type QueryDataTransferForAllDomainByteResponseHeader struct {
}

func (s QueryDataTransferForAllDomainByteResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponseHeader) GoString() string {
  return s.String()
}




type GetBandwidthSavingRatioRequest struct {
  // {"en":"Start time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.Cannot exceed current time;\n3.The most recent six-month (183 days) data are available.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.不能大于当前时间\n3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time;\n3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception;\n4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。\n3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常\n4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:\n1.The maximum number of deliverable domain names is 200 by default;\n2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names);\n3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.","zh_CN":"域名:\n1.可传递域名数量上限默认为200个\n2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)\n3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"1d","en":"Time interval of data: 5m (5 min), 1h (1 hour), 1d (1 day);\nThe default is 1d.","zh_CN":"数据粒度:\n1.支持5m(5分钟)、1h(1小时)、1d(天)\n2.不传默认1d。","exampleValue":"5m,1h,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
}

func (s GetBandwidthSavingRatioRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioRequest) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioRequest) SetDateFrom(v string) *GetBandwidthSavingRatioRequest {
  s.DateFrom = &v
  return s
}

func (s *GetBandwidthSavingRatioRequest) SetDateTo(v string) *GetBandwidthSavingRatioRequest {
  s.DateTo = &v
  return s
}

func (s *GetBandwidthSavingRatioRequest) SetDomain(v []*string) *GetBandwidthSavingRatioRequest {
  s.Domain = v
  return s
}

func (s *GetBandwidthSavingRatioRequest) SetGranularity(v string) *GetBandwidthSavingRatioRequest {
  s.Granularity = &v
  return s
}

type GetBandwidthSavingRatioRequestHeader struct {
}

func (s GetBandwidthSavingRatioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioRequestHeader) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioPaths struct {
}

func (s GetBandwidthSavingRatioPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioPaths) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioParameters struct {
}

func (s GetBandwidthSavingRatioParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioParameters) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*GetBandwidthSavingRatioResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthSavingRatioResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponse) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioResponse) SetCode(v string) *GetBandwidthSavingRatioResponse {
  s.Code = &v
  return s
}

func (s *GetBandwidthSavingRatioResponse) SetMessage(v string) *GetBandwidthSavingRatioResponse {
  s.Message = &v
  return s
}

func (s *GetBandwidthSavingRatioResponse) SetData(v []*GetBandwidthSavingRatioResponseData) *GetBandwidthSavingRatioResponse {
  s.Data = v
  return s
}

type GetBandwidthSavingRatioResponseData struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format","zh_CN":"实际查询时间,格式 yyyy-MM-dd HH:mm"}
  QueryTime *string `json:"queryTime,omitempty" xml:"queryTime,omitempty" require:"true"`
  // {"en":"Average of total saving of bandwidth.","zh_CN":"总节省带宽的平均值"}
  AvgSavedBandwidthRatio *GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio `json:"avgSavedBandwidthRatio,omitempty" xml:"avgSavedBandwidthRatio,omitempty" require:"true" type:"Struct"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetBandwidthSavingRatioResponseDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthSavingRatioResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseData) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioResponseData) SetQueryTime(v string) *GetBandwidthSavingRatioResponseData {
  s.QueryTime = &v
  return s
}

func (s *GetBandwidthSavingRatioResponseData) SetAvgSavedBandwidthRatio(v *GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio) *GetBandwidthSavingRatioResponseData {
  s.AvgSavedBandwidthRatio = v
  return s
}

func (s *GetBandwidthSavingRatioResponseData) SetDataSeries(v []*GetBandwidthSavingRatioResponseDataDataSeries) *GetBandwidthSavingRatioResponseData {
  s.DataSeries = v
  return s
}

type GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio struct {
}

func (s GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioResponseDataDataSeries struct     {
  // {"en":"Timetamp\n1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range, For example yyyy-MM-dd 00:05 represents data in the range from 00:00 to 00:05.\n2.The data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents data values in the previous time granularity range such as yyyy-MM-dd 01 that represent data from 00 to 01.\n3. the data granularity of the query is daily, the format is yyyy-MM-dd; Each time slice data value represents the data value for that day.\n4.Returns the timetamp contained in start time and end time.","zh_CN":"时间片\n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 01,代表00到01之间的数据。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Saving of bandwidth, retain four decimals","zh_CN":"节省带宽,保留4位小数"}
  SavedBandwidthRatio *string `json:"savedBandwidthRatio,omitempty" xml:"savedBandwidthRatio,omitempty" require:"true"`
}

func (s GetBandwidthSavingRatioResponseDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioResponseDataDataSeries) SetTimestamp(v string) *GetBandwidthSavingRatioResponseDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetBandwidthSavingRatioResponseDataDataSeries) SetSavedBandwidthRatio(v string) *GetBandwidthSavingRatioResponseDataDataSeries {
  s.SavedBandwidthRatio = &v
  return s
}

type GetBandwidthSavingRatioResponseHeader struct {
}

func (s GetBandwidthSavingRatioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseHeader) GoString() string {
  return s.String()
}




type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest struct {
  // {"en":"From date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8)
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过 10 分钟 ；
  // 4.只能查询最近半年内数据。
  // 5.dateFrom 和 dateTo 都不填则默认查询过去 10 分钟的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"To time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8)
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception", "zh_CN":"
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；
  // 3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain and stream group:
  // 1. the maximum number of transmissible domain names and stream name groups is 20 by default 
  // 3. stream name stream : the maximum number of stream names that can be passed under a group of domain names and stream name groups ( that is, under a single domain name ) defaults to 2000 ( which can be adjusted with technical support ).", "zh_CN":"域名和流名组:
  // 
  // 1.可传递的域名和流名组数量上限默认为20组(可联系技术支持调整)；
  // 2.域名domain:一组域名和流名组中只能传递单个域名，且域名必须传递；
  // 3.流名stream:一组域名和流名组下(即单个域名下)可传递的流名数量上限默认为2000个(可联系技术支持调整)，流名未传递时默认查询域名下所有流名，但当域名下流名数量超过限制时不可查询(报错)。"}
  DomainStream []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
  // {"en":"data granularity
  // 1m: 1 min granularity, 5m: 5 min granularity
  // Default 5-minute granularity data query", "zh_CN":"数据粒度
  // 1.1m: 1分钟粒度，5m: 5分钟粒度
  // 2.默认查询5分钟粒度数据"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Data type, flow: traffic, bandwidth: bandwidth", "zh_CN":"数据类型，flow:流量，bandwidth:带宽"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
  // {"en":"Filter empty stream name:
  // 1. The input optional value '0' or '1';
  // 2. The data whose parameter 0 does not filter flow is null, and the data whose parameter 1 filter stream is null.
  // 3. The default value is 0;", "zh_CN":"是否过滤空流名
  // 1.入参可选值 '0' 或 '1' ；
  // 2.传参 0 不过滤流名为空的数据，传参 1 则过滤流名为空的数据 ；
  // 3.默认值为 0；"}
  FilterEmptyStream *int `json:"filterEmptyStream,omitempty" xml:"filterEmptyStream,omitempty"`
  // {"en":"Query string, multiple queries are not supported, and fuzzy matching is used by default, allowing regular expression queries", "zh_CN":"查询字符串，不支持多次查询，默认使用模糊匹配，允许正则表达式查询"}
  Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
  // {"en":"Whether to return the number of online users:
  //     1.If true is passed, the number of online users will only be returned when dataInterval=1m. Does not return when dataInterval=5min;
  //     2.If no or false is passed,the default is not to return the number of online users.", "zh_CN":"是否返回在线人数: 
  //     1.传true, 仅当dataInterval=1m时,返回在线人数,5min时不返回;
  //     2.不传或传false，则默认不返回在线人数."}
  IsReturnOnline *string `json:"isReturnOnline,omitempty" xml:"isReturnOnline,omitempty"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDateFrom(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDateTo(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDomainStream(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DomainStream = v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDataInterval(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDataType(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DataType = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetFilterEmptyStream(v int) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.FilterEmptyStream = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetKeyword(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.Keyword = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetIsReturnOnline(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.IsReturnOnline = &v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream struct     {
  // {"en":"Domain:
  // 1. The maximum number of deliverable domain names is 200 by default
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为200个
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Celebrity: Publishing Point  Stream Name. Example: live/test-20180101-test where live is a publishing point and test-20180101-test is a stream name
  // No, the default queries all stream data in the specified domain name", "zh_CN":"
  // 流名:'发布点'+'流名'。例如:live/test-20180101-test,其中live是发布点,test-20180101-test是流名;
  // 不传，默认查询指定域名下的所有流的数据"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) SetDomain(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) SetStream(v []*string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream {
  s.Stream = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetCode(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Code = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetMessage(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Message = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetData(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Data = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DomainOfStreamList []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList `json:"domainOfStreamList,omitempty" xml:"domainOfStreamList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) SetDomain(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData {
  s.Domain = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) SetDomainOfStreamList(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData {
  s.DomainOfStreamList = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList struct     {
  // {"en":"stream name", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"aggregate flow
  // 1. Keep 2 decimal places, in MB;
  // 2. When entering dataType is flow.
  // 3. Only data with traffic greater than0;is returned.", "zh_CN":"总流量
  // 1.保留2位小数，单位为MB;
  // 2.当入参dataType为flow时，返回。
  // 3.仅返回流量大于 0  的数据。"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"peak bandwidth
  // 1. keep 2 decimal places in Mbps
  // 2. when entering dataType is bandwidth, return.
  // 3. only returns data with bandwidth greater than 0.", "zh_CN":"峰值带宽
  // 1.保留2位小数，单位Mbps
  // 2.当入参dataType为bandwidth时，返回。
  // 3.仅返回带宽大于 0 的数据。"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty" require:"true"`
  DetailList []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetStream(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.Stream = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetTotalFlow(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.TotalFlow = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetBandwidthPeakValue(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetDetailList(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.DetailList = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList struct     {
  // {"en":"time
  // 1. Each time slice data value represents the data value in the previous time granularity range. 
  // 2. The data granularity of the query is 5m, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range.", "zh_CN":"时间
  // 1.查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是(yyyy-MM-dd+1) 00:00。
  // 2.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是(yyyy-MM-dd+1) 00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Data at the point in time:
  // 1. keep 2 decimal places;
  // 2. When the dataType is flow, value is flow in MB;
  // 3. value is bandwidth value in Mbps when dataType is bandwidth.", "zh_CN":"时间点的数据
  // 1.保留2位小数；
  // 2.当入参dataType为flow时，value值为流量，单位为MB；
  // 3.当入参dataType为bandwidth时，value值为带宽值，单位为Mbps；"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Number of online users at a time point.", "zh_CN":"时间点的在线人数"}
  OnlineUser *int `json:"onlineUser,omitempty" xml:"onlineUser,omitempty" require:"true"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetTimestamp(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetValue(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.Value = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetOnlineUser(v int) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.OnlineUser = &v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainPaths struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainPaths) GoString() string {
  return s.String()
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainParameters struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainParameters) GoString() string {
  return s.String()
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestHeader struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseHeader struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseHeader) GoString() string {
  return s.String()
}




