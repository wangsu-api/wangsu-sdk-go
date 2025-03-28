package reportflow

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportFlowDomainCountryServiceRequest struct {
  // {"en":"Starting time
  // 
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2. Cannot be greater than the current time
  // 3. Get up to the last six months (183 days) of data.", "zh_CN":"开始时间
  // 1. 时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2. 不能大于当前时间
  // 3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 1. Time format 2016-12-02T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception
  // 4. Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo can&rsquo;t exceed 7 days (can contact technical support adjustment, up to 31 days).", "zh_CN":"结束时间:
  // 1. 时间格式2016-12-02T10:00:00+08:00
  // 2. 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3. dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4. 允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最长31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity:
  // 1. Support 5m (5 minutes granularity),1d (1 day granularity)
  // 2. Do not pass the default to 5m", "zh_CN":"数据粒度:
  // 1. 支持5m(5分钟粒度),1d(天粒度)
  // 2. 不传默认为5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"domain name:
  // 
  // 1. The maximum number of deliverable domain names is 20 by default (can be contacted by technical support);
  // 2. Automatically filter out illegal domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of legitimate domain names).", "zh_CN":"域名:
  // 1. 可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2. 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Country code:
  // 
  // 1. Do not pass the default query for all countries and regions;
  // 2. The values that can be passed are detailed in the Countrycode list of appendix table on the API Overview page", "zh_CN":"国家地区代号:
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
  // 3. The result hierarchy is fixed in order, and the order of the parameters does not affect the order of the returned results. For example: 'groupBy': ['domain','country'] and 'groupBy': ['country','domain'] return the same result.",
  // "zh_CN":"1. 可选值domain、country、aggregatedOversea,可传入单个或多个值,其中不能同时传aggregatedOversea 和 country;
  // 2. 有传入则按照该维度展示明细数据:
  //   1.domain:按照域名维度进行分组展示;
  //   2.domain:country:按照国家维度进行分组展示;
  //   3.aggregatedOversea:按照国内 和 海外维度进行分组展示
  // 3. 返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:groupBy: [domain,country]与groupBy: [country,domain]返回结果一样。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceRequest) SetDateFrom(v string) *ReportFlowDomainCountryServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetDateTo(v string) *ReportFlowDomainCountryServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetDataInterval(v string) *ReportFlowDomainCountryServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetDomain(v []*string) *ReportFlowDomainCountryServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetCountryCode(v []*string) *ReportFlowDomainCountryServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetGroupBy(v []*string) *ReportFlowDomainCountryServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowDomainCountryServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceResponse) SetResult(v []*ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult) *ReportFlowDomainCountryServiceResponse {
  s.Result = v
  return s
}

type ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult) SetDomain(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult) SetCountryData(v []*ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResult {
  s.CountryData = v
  return s
}

type ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData struct     {
  // {"en":"Country code", "zh_CN":"国家地区代号"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Country name", "zh_CN":"国家地区名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  // {"en":"Summary of traffic in national regions: Summary of traffic flow in a single country region during the query period, unit of measure MB, retaining 2 decimal places", "zh_CN":"国家地区流量汇总:单个国家地区流量在查询时段内的流量汇总值,计量单位MB,保留2位小数"}
  FlowSum *string `json:"flowSum,omitempty" xml:"flowSum,omitempty" require:"true"`
  // {"en":"National regional traffic ratio: the proportion (percentage) of traffic value in a single country region during the query period, retaining 2 decimal places", "zh_CN":"国家地区流量占比:单个国家地区流量在查询时段内的流量值的占比(百分比),保留2位小数"}
  FlowPercentage *string `json:"flowPercentage,omitempty" xml:"flowPercentage,omitempty" require:"true"`
  FlowData []*ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) SetCountryCode(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) SetCountryName(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData {
  s.CountryName = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) SetFlowSum(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData {
  s.FlowSum = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) SetFlowPercentage(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData {
  s.FlowPercentage = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData) SetFlowData(v []*ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryData {
  s.FlowData = v
  return s
}

type ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData struct     {
  // {"en":"1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. Returns the time slice contained in the start time and end time.", "zh_CN":"1. 查询的数据粒度为5m时,格式为yyyy-MM-dd  HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。2. 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Flow value, in megabytes, 2 decimal places", "zh_CN":"流量值,计量单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Bandwidth value. Unit is Mbps and 2 digits of decimals are allowed.", "zh_CN":"带宽值,单位Mbps,保留2位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData) SetTimestamp(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData) SetValue(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData {
  s.Value = &v
  return s
}

func (s *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData) SetBandwidth(v string) *ReportFlowDomainCountryServiceReportFlowDomainCountryServiceResponseResultCountryDataFlowData {
  s.Bandwidth = &v
  return s
}

type ReportFlowDomainCountryServicePaths struct {
}

func (s ReportFlowDomainCountryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServicePaths) GoString() string {
  return s.String()
}

type ReportFlowDomainCountryServiceParameters struct {
}

func (s ReportFlowDomainCountryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowDomainCountryServiceRequestHeader struct {
}

func (s ReportFlowDomainCountryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowDomainCountryServiceResponseHeader struct {
}

func (s ReportFlowDomainCountryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponseHeader) GoString() string {
  return s.String()
}




type QuerySumUpTrafficUnderAccountRequest struct {
}

func (s QuerySumUpTrafficUnderAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountRequest) GoString() string {
  return s.String()
}

type QuerySumUpTrafficUnderAccountResponse struct {
  // {"en":"Total traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *string `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量数据"}
  FlowData []*QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySumUpTrafficUnderAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountResponse) GoString() string {
  return s.String()
}

func (s *QuerySumUpTrafficUnderAccountResponse) SetFlowSummary(v string) *QuerySumUpTrafficUnderAccountResponse {
  s.FlowSummary = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountResponse) SetFlowData(v []*QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData) *QuerySumUpTrafficUnderAccountResponse {
  s.FlowData = v
  return s
}

type QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData struct     {
  // {"en":"Date
  // 1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.
  // 2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.
  // 3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.
  // 4.Return the time slice contained in start time and the time slice contained in end time", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"流量,保留2位小数,单位为MB"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData) GoString() string {
  return s.String()
}

func (s *QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData) SetTimestamp(v string) *QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData) SetFlow(v string) *QuerySumUpTrafficUnderAccountQuerySumUpTrafficUnderAccountResponseFlowData {
  s.Flow = &v
  return s
}

type QuerySumUpTrafficUnderAccountPaths struct {
}

func (s QuerySumUpTrafficUnderAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountPaths) GoString() string {
  return s.String()
}

type QuerySumUpTrafficUnderAccountParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and 'dateto';
  // 3.Period between 'datafrom' and 'dateto' cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须小于当前时间和dateto;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than 'datefrom'; 
  // 3.if it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于datefrom;如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is a specific configuration on data collecting granularity for the custome", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QuerySumUpTrafficUnderAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountParameters) GoString() string {
  return s.String()
}

func (s *QuerySumUpTrafficUnderAccountParameters) SetDatefrom(v string) *QuerySumUpTrafficUnderAccountParameters {
  s.Datefrom = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountParameters) SetDateto(v string) *QuerySumUpTrafficUnderAccountParameters {
  s.Dateto = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountParameters) SetType(v string) *QuerySumUpTrafficUnderAccountParameters {
  s.Type = &v
  return s
}

type QuerySumUpTrafficUnderAccountRequestHeader struct {
}

func (s QuerySumUpTrafficUnderAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountRequestHeader) GoString() string {
  return s.String()
}

type QuerySumUpTrafficUnderAccountResponseHeader struct {
}

func (s QuerySumUpTrafficUnderAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowDirInfoServiceRequest struct {
  // {"en":"Start time
  // 				1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 				2.Must be smaller than the current time and dateTo;
  // 				3.Period between dataFrom and dateTo cannot be longer than 31 days(technical support can be contacted to adjust);
  // 				4.You can only query data for the last 6 months.", "zh_CN":"开始时间
  // 				1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 				2.必须小于当前时间和dateTo;
  // 				3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);
  // 				4.只能查询最近半年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 				1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 				2.Must be greater than dateFrom;
  // 				3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 				1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 				2.必须大于dateFrom;
  // 				3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Directory levels, value range 1-4. Only one vlaue can be submitted", "zh_CN":"目录层级,取值范围1~4,只能提交单个值"}
  DirHierarchy *string `json:"dirHierarchy,omitempty" xml:"dirHierarchy,omitempty" require:"true"`
  DomainDir []*ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir `json:"domainDir,omitempty" xml:"domainDir,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDirInfoServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowDirInfoServiceRequest) SetDateFrom(v string) *ReportFlowDirInfoServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowDirInfoServiceRequest) SetDateTo(v string) *ReportFlowDirInfoServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowDirInfoServiceRequest) SetDirHierarchy(v string) *ReportFlowDirInfoServiceRequest {
  s.DirHierarchy = &v
  return s
}

func (s *ReportFlowDirInfoServiceRequest) SetDomainDir(v []*ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir) *ReportFlowDirInfoServiceRequest {
  s.DomainDir = v
  return s
}

type ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir struct     {
  // {"en":"Domain
  // 					1.Need to meet the regular expression rules that are used to validate domains;
  // 					2.Domain number limits can be adjusted depending on different accounts. The default value is 1;", "zh_CN":"域名
  // 					1.需要满足域名的正则校验;
  // 					2.域名个数限制根据账号可调,默认为1个;"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Table of contents
  // 					1.Directory number limits can be adjusted depending on different accounts. The default value is 200;
  // 					2.Empty value means to query all directories. Number of directories shall not exceed set limit;
  // 					3.Invalid directories are not returned", "zh_CN":"目录
  // 					1.目录个数限制根据账号可调,默认为200个;
  // 					2.不传代表查询该域名下的所有目录,同时接受目录个数限制;
  // 					3.无效的目录不返回"}
  Dir []*string `json:"dir,omitempty" xml:"dir,omitempty" type:"Repeated"`
}

func (s ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir) GoString() string {
  return s.String()
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir) SetDomain(v string) *ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir {
  s.Domain = &v
  return s
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir) SetDir(v []*string) *ReportFlowDirInfoServiceReportFlowDirInfoServiceRequestDomainDir {
  s.Dir = v
  return s
}

type ReportFlowDirInfoServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDirInfoServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowDirInfoServiceResponse) SetResult(v []*ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult) *ReportFlowDirInfoServiceResponse {
  s.Result = v
  return s
}

type ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"details", "zh_CN":"详情数据"}
  Details []*ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult) SetDomain(v string) *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult) SetDetails(v []*ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails) *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResult {
  s.Details = v
  return s
}

type ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails struct     {
  // {"en":"Directory name of corresponding level", "zh_CN":"对应层级的目录名称"}
  Dir *string `json:"dir,omitempty" xml:"dir,omitempty" require:"true"`
  // {"en":"Total traffic, unit is MB and 2  digits of decimals allowed", "zh_CN":"总流量,单位MB,保留2位小数"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Bandwidth peak value with granularity of 5 minutes. Unit Mbps, two decimal digits allowed", "zh_CN":"带宽峰值,单位Mbps,保留2位小数"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty" require:"true"`
}

func (s ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails) GoString() string {
  return s.String()
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails) SetDir(v string) *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails {
  s.Dir = &v
  return s
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails) SetTotalFlow(v string) *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails) SetBandwidthPeakValue(v string) *ReportFlowDirInfoServiceReportFlowDirInfoServiceResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

type ReportFlowDirInfoServicePaths struct {
}

func (s ReportFlowDirInfoServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServicePaths) GoString() string {
  return s.String()
}

type ReportFlowDirInfoServiceParameters struct {
}

func (s ReportFlowDirInfoServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowDirInfoServiceRequestHeader struct {
}

func (s ReportFlowDirInfoServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowDirInfoServiceResponseHeader struct {
}

func (s ReportFlowDirInfoServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDirInfoServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowOriginIspProvinceByteServiceRequest struct {
  // {"en":"Starting time:
  //         1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  //         2. Must be greater than the current time -183 days, and less than the current time and dateTo;
  //         3. The difference between dateFrom and dateTo cannot exceed 7 days(technical support can be contacted to adjust);
  //         4. dateFrom and dateTo are either passed or not passed;
  //         5. dateFrom and dateTo are not passed, the default query data for the past 24 hours", "zh_CN":"开始时间
  //         1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  //         2.必须大于当前时间-183天，并且小于当前时间和dateTo;
  //         3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);
  //         4.dateFrom和dateTo要么都传递，要么都不传递;
  //         5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  //         1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  //         2. Must be greater than dateFrom; if greater than the current time, re-assign the current time;", "zh_CN":"结束时间
  //         1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  //         2.必须大于dateFrom;如果大于当前时间，则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name: 
  // 		1.The maximum number of deliverable domain names is 20 by default (can be contacted by technical support adjustment). 
  // 		2.All domain names under the account are not passed when the entry is passed, but it cannot be queried when the number of domain names under the account exceeds the limit (error).", "zh_CN":"域名:
  // 		1.可传递域名数量上限默认为20个(可联系技术支持调整)
  // 		2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询(报错)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5 minutes granularity, 1h: 1 hour granularity", "zh_CN":"数据粒度，5m:5分钟粒度，1h:1小时粒度"}
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
  // {"en":"Grouping dimension:
  //         1. The optional values are domain,province, and isp, which can pass multiple values.
  //         2. If there is an incoming, the detailed data will be displayed according to the dimension;", "zh_CN":"分组维度:
  //         1.可选值为domain、province、isp，可传入多个值;
  //         2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceByteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetDateFrom(v string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetDateTo(v string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetDomain(v []*string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetDataInterval(v string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetProvince(v []*string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetIsp(v []*string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceRequest) SetGroupBy(v []*string) *ReportFlowOriginIspProvinceByteServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowOriginIspProvinceByteServiceResponse struct {
  Result []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceByteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceByteServiceResponse) SetResult(v []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult) *ReportFlowOriginIspProvinceByteServiceResponse {
  s.Result = v
  return s
}

type ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult) SetDomain(v string) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult) SetIspData(v []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData) SetIsp(v string) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData) SetProvinceData(v []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  OriginFlowData []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData `json:"originFlowData,omitempty" xml:"originFlowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData) SetOriginFlowData(v []*ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceData {
  s.OriginFlowData = v
  return s
}

type ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData struct     {
  // {"en":"time,
  //         1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00;
  //         2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00;
  //         3. Return to the time slice included in the start time and end time.", "zh_CN":"时间，
  //         1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05，最后一个时间片是(yyyy-MM-dd+1)00:00;
  //         2.查询的数据粒度为1h时，格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 01，最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;
  //         3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Return source traffic, unit Byte", "zh_CN":"回源流量，单位Byte"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData) SetTimestamp(v string) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData) SetValue(v string) *ReportFlowOriginIspProvinceByteServiceReportFlowOriginIspProvinceByteServiceResponseResultIspDataProvinceDataOriginFlowData {
  s.Value = &v
  return s
}

type ReportFlowOriginIspProvinceByteServicePaths struct {
}

func (s ReportFlowOriginIspProvinceByteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServicePaths) GoString() string {
  return s.String()
}

type ReportFlowOriginIspProvinceByteServiceParameters struct {
}

func (s ReportFlowOriginIspProvinceByteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowOriginIspProvinceByteServiceRequestHeader struct {
}

func (s ReportFlowOriginIspProvinceByteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowOriginIspProvinceByteServiceResponseHeader struct {
}

func (s ReportFlowOriginIspProvinceByteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceByteServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowOriginIspProvinceServiceRequest struct {
  // {"en":"Start time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days;
  // 
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天;(可联系技术支持调整)
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
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5-minute granularity, 1h: 1-hour granularity, The default granularity is 5 minutes.", "zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度, 默认粒度5分钟"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province:
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
  // 
  // 2.The data is displayed according to the specified dimension, if not transmitted, aggregate according to all dimensions.", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据, 没传则按照所有维度聚合;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetDateFrom(v string) *ReportFlowOriginIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetDateTo(v string) *ReportFlowOriginIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetDomain(v []*string) *ReportFlowOriginIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetDataInterval(v string) *ReportFlowOriginIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetProvince(v []*string) *ReportFlowOriginIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetIsp(v []*string) *ReportFlowOriginIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportFlowOriginIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowOriginIspProvinceServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceServiceResponse) SetResult(v []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult) *ReportFlowOriginIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ispData", "zh_CN":"ISP数据"}
  IspData []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult) SetDomain(v string) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult) SetIspData(v []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"provinceData", "zh_CN":"省份数据"}
  ProvinceData []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"originFlowData", "zh_CN":"流量数据"}
  OriginFlowData []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData `json:"originFlowData,omitempty" xml:"originFlowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData) SetOriginFlowData(v []*ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceData {
  s.OriginFlowData = v
  return s
}

type ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData struct     {
  // {"en":"Time,
  // 
  //    1.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  //    2.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;
  // 
  //    3.Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间,
  //    1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  //    2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;
  //    3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Origin traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"回源流量,保留2位小数,单位MB"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Average origin traffic. Keep two decimals. Unit: Mbps", "zh_CN":"回源平均带宽,保留2位小数,单位Mbps"}
  OriginBandwidth *string `json:"originBandwidth,omitempty" xml:"originBandwidth,omitempty" require:"true"`
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData) SetTimestamp(v string) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData) SetValue(v string) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData {
  s.Value = &v
  return s
}

func (s *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData) SetOriginBandwidth(v string) *ReportFlowOriginIspProvinceServiceReportFlowOriginIspProvinceServiceResponseResultIspDataProvinceDataOriginFlowData {
  s.OriginBandwidth = &v
  return s
}

type ReportFlowOriginIspProvinceServicePaths struct {
}

func (s ReportFlowOriginIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportFlowOriginIspProvinceServiceParameters struct {
}

func (s ReportFlowOriginIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowOriginIspProvinceServiceRequestHeader struct {
}

func (s ReportFlowOriginIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowOriginIspProvinceServiceResponseHeader struct {
}

func (s ReportFlowOriginIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowOriginIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowIspProvinceHitRateDetailServiceRequest struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 24 hours,if the dataInterval is 1m,period between dataFrom and dateTo cannot be longer than 6 hours(you can contact technical support to adjust it);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 30 minutes is queried", "zh_CN":"开始时间
  //         1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  //         2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  //         3.dateFrom和dateTo相差不能超过24小时;一分钟粒度dateFrom和dateTo相差不能超过6小时(可联系技术支持调整)
  //         4.dateFrom和dateTo要么都传递,要么都不传递;
  //         5.dateFrom和dateTo都未传递,则默认查询过去30分钟的数据;"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  //         1/格式为yyyy-MM-ddTHH:mm:ss+08:00;
  //         2/必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain
  // 1.domain number limits can be adjusted depending on different accounts. The default value is 20.
  // 2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)", "zh_CN":"域名:
  //         1.可传递域名数量上限默认为20个(可联系技术支持调整);
  //         2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  //         3.域名超过上限,报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. Support 1m (1 minute granularity),5m (5 minutes granularity)
  // 2. The default value is 5m", "zh_CN":"数据粒度:
  //         1.1m:1分钟粒度, 5m:5分钟粒度
  //         2.不传默认查询 5m"}
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
  // {"en":"Query Dimension
  // 1.Optional values: flow, request
  // 2.The default value is flow
  // 3.Flow: flow, two decimal places reserved;
  // 4.Request: Number of requests", "zh_CN":"查询维度
  //         1.可选值 flow、request
  //         2.不传默认 flow
  //         3.flow:流量,保留两位小数;
  //         4.request:请求数"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional values: domain,province,isp. Multiple values can be transferred in;
  // Display detailed data according to this dimension if it is transferred in", "zh_CN":"分组维度
  //         可选值为domain、province、isp,可传入多个值;
  //         有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowIspProvinceHitRateDetailServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetDateFrom(v string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetDateTo(v string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetDomain(v []*string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetDataInterval(v string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetProvince(v []*string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetIsp(v []*string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetQueryBy(v string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.QueryBy = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceRequest) SetGroupBy(v []*string) *ReportFlowIspProvinceHitRateDetailServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowIspProvinceHitRateDetailServiceResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result infotmation", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detail data of request result", "zh_CN":"请求结果的详细数据"}
  Data []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceHitRateDetailServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceHitRateDetailServiceResponse) SetCode(v string) *ReportFlowIspProvinceHitRateDetailServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceResponse) SetMessage(v string) *ReportFlowIspProvinceHitRateDetailServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceResponse) SetData(v []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData) *ReportFlowIspProvinceHitRateDetailServiceResponse {
  s.Data = v
  return s
}

type ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ISP data", "zh_CN":"运营商数据"}
  IspData []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData) SetDomain(v string) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData) SetIspData(v []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseData {
  s.IspData = v
  return s
}

type ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"Province data", "zh_CN":"省份数据"}
  ProvinceData []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData) SetIsp(v string) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData) SetProvinceData(v []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"-", "zh_CN":""}
  Details []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData) SetProvince(v string) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData) SetDetails(v []*ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceData {
  s.Details = v
  return s
}

type ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails struct     {
  // {"en":"Time,the format is yyyy-MM-dd HH:mm", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Hit data incording to query by(flow/requests)", "zh_CN":"命中数据 flow:流量,保留两位小数; request:请求数"}
  HitValue *string `json:"hitValue,omitempty" xml:"hitValue,omitempty" require:"true"`
  // {"en":"Hit rate,keep four decimal places", "zh_CN":"命中率,保留四位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails) SetTimestamp(v string) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails) SetHitValue(v string) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails {
  s.HitValue = &v
  return s
}

func (s *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails) SetHitRate(v string) *ReportFlowIspProvinceHitRateDetailServiceReportFlowIspProvinceHitRateDetailServiceResponseDataIspDataProvinceDataDetails {
  s.HitRate = &v
  return s
}

type ReportFlowIspProvinceHitRateDetailServicePaths struct {
}

func (s ReportFlowIspProvinceHitRateDetailServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServicePaths) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceHitRateDetailServiceParameters struct {
}

func (s ReportFlowIspProvinceHitRateDetailServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceHitRateDetailServiceRequestHeader struct {
}

func (s ReportFlowIspProvinceHitRateDetailServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceHitRateDetailServiceResponseHeader struct {
}

func (s ReportFlowIspProvinceHitRateDetailServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceHitRateDetailServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryTrafficRequestInTotalAndPeakValueRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time);
  // 2. No bigger than the current time;
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. the time format is 2016-12-02T10:00:00+08:00;
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used;
  // 3.  If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default, if only one field is filled in and one is left empty, then exception will be occur;
  // 4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days).", "zh_CN":"结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00;
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间;
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时，如仅有一个未传，抛异常;
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大不超过31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The default upper limit to domains that can be entered is 200 (Contact technical support to adjust, the upper limit is 500);
  // 2. All domains under the account are queried if this input parameter is not specified, but if the number of domains under the account exceeds limits, no query will be done (Error).", "zh_CN":"域名：
  // 1.可传递域名数量上限默认200个（可联系技术支持调整，最高上限500）;
  // 2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Group keywords:
  // 1. By default, group data will be displayed;
  // 2. If there are keywords entered, value details shall be displayed by keywords;
  // If domain is specified to groupBy, it means results are returned according to domains;
  // 3. Only domain can be specified.", "zh_CN":"分组关键词：
  // 1.默认聚合展示；
  // 2.传入关键词则代表需要按照关键词对应的值展示明细；
  // 例如groupBy传domain，则代表返回按照domain明细展开。
  // 3.只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryTrafficRequestInTotalAndPeakValueRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueRequest) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetDateFrom(v string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetDateTo(v string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetDomain(v []*string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.Domain = v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetGroupBy(v []*string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.GroupBy = v
  return s
}

type QueryTrafficRequestInTotalAndPeakValueResponse struct {
  // {"en":"", "zh_CN":""}
  Result []*QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficRequestInTotalAndPeakValueResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueResponse) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponse) SetResult(v []*QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) *QueryTrafficRequestInTotalAndPeakValueResponse {
  s.Result = v
  return s
}

type QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"totalFlow, unit: MB, 2 decimal places reserved, example (74099.92)", "zh_CN":"总流量，单位:MB ，保留2位小数，示例 ( 74099.92 )"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"totalRequest", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"peakBandwidth, unit: Mbps, 2 decimal places reserved, example (74099.92)", "zh_CN":"峰值带宽，单位: Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"peakTime", "zh_CN":"峰值时间，示例（2019-02-13 18:01）"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak request", "zh_CN":"请求数峰值"}
  PeakRequest *string `json:"peakRequest,omitempty" xml:"peakRequest,omitempty" require:"true"`
  // {"en":"Peak time of request", "zh_CN":"请求数峰值时间"}
  PeakRequestTime *string `json:"peakRequestTime,omitempty" xml:"peakRequestTime,omitempty" require:"true"`
  // {"en":"", "zh_CN":""}
  FlowRequestData []*QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData `json:"flowRequestData,omitempty" xml:"flowRequestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetDomain(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetTotalFlow(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetTotalRequest(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakBandwidth(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakTime(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakTime = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakRequest(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakRequest = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakRequestTime(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakRequestTime = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult) SetFlowRequestData(v []*QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.FlowRequestData = v
  return s
}

type QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData struct     {
  // {"en":"timestamp
  // 1. When the data granularity of query is 5m, the format is yyyy-mm-dd HH: MM;Each time slice data value represents the data value in the previous time granularity range.The time slice at the beginning of the day is yyyy-mm-dd 00:05, and the last time slice is (yyyy-mm-dd +1) 00:00;
  // 2. Return the time slice of start time and end time.", "zh_CN":"时间片
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00;
  // 2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量值，单位MB，保留2位小数；"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {"en":"Bandwidth, unit: Mbps, 2 decimal places reserved, example (931556.21)", "zh_CN":"带宽值，单位: Mbps，保留2位小数，示例 （931556.21）"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"request", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetTimestamp(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetFlow(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Flow = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetBandwidth(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Bandwidth = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetRequest(v string) *QueryTrafficRequestInTotalAndPeakValueQueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Request = &v
  return s
}

type QueryTrafficRequestInTotalAndPeakValuePaths struct {
}

func (s QueryTrafficRequestInTotalAndPeakValuePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValuePaths) GoString() string {
  return s.String()
}

type QueryTrafficRequestInTotalAndPeakValueParameters struct {
}

func (s QueryTrafficRequestInTotalAndPeakValueParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueParameters) GoString() string {
  return s.String()
}

type QueryTrafficRequestInTotalAndPeakValueRequestHeader struct {
}

func (s QueryTrafficRequestInTotalAndPeakValueRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueRequestHeader) GoString() string {
  return s.String()
}

type QueryTrafficRequestInTotalAndPeakValueResponseHeader struct {
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowIspProvinceByteServiceRequest struct {
  // {'en':'Starting time
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. Must be greater than the current time -183 days, and less than the current time and dateTo;
  // 
  // 3. The difference between dateFrom and dateTo cannot exceed 7 days (You can contact technical support for adjustment, up to 31 days);
  // 
  // 4.dateFrom and dateTo are either passed or not passed;
  // 
  // 5. DateFrom and dateTo are not passed, the default query data for the past 24 hours;', 'zh_CN':'开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于当前时间-183天，并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过7天（可联系技术支持调整，最大31天）；
  // 4.dateFrom和dateTo要么都传递，要么都不传递；
  // 5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据；'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End Time
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. Must be greater than dateFrom; if greater than the current time, re-assign the current time;', 'zh_CN':'结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；如果大于当前时间，则重新赋值为当前时间；'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'The domain name and the number of domain names are adjustable according to the account number. The default is 20', 'zh_CN':'域名，域名个数限制根据账号可调，默认为20个'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'Data granularity, 5m: 5 minutes granularity, 1h: 1 hour granularity', 'zh_CN':'数据粒度，5m：5分钟粒度，1h：1小时粒度'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Province
  // 
  // 1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.', 'zh_CN':'省份
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。'}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {'en':'ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.', 'zh_CN':'运营商：
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节'}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {'en':'Grouping dimension
  // 
  // 1. The optional values are domain,province, and isp, which can pass multiple values.
  // 
  // 2. If there is an incoming, the detailed data will be displayed according to the dimension;', 'zh_CN':'分组维度
  // 1.可选值为domain、province、isp，可传入多个值；
  // 2.有传入则按照该维度展示明细数据；'}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowIspProvinceByteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetDateFrom(v string) *ReportFlowIspProvinceByteServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetDateTo(v string) *ReportFlowIspProvinceByteServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetDomain(v []*string) *ReportFlowIspProvinceByteServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetDataInterval(v string) *ReportFlowIspProvinceByteServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetProvince(v []*string) *ReportFlowIspProvinceByteServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetIsp(v []*string) *ReportFlowIspProvinceByteServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowIspProvinceByteServiceRequest) SetGroupBy(v []*string) *ReportFlowIspProvinceByteServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowIspProvinceByteServiceResponse struct {
  Result []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceByteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceByteServiceResponse) SetResult(v []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult) *ReportFlowIspProvinceByteServiceResponse {
  s.Result = v
  return s
}

type ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult) SetDomain(v string) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult) SetIspData(v []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData struct     {
  // {'en':'Internet service providers', 'zh_CN':'运营商'}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData) SetIsp(v string) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData) SetProvinceData(v []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData struct     {
  // {'en':'province', 'zh_CN':'省份'}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  FlowData []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData) SetFlowData(v []*ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceData {
  s.FlowData = v
  return s
}

type ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData struct     {
  // {'en':'time
  // 1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00.
  // 3. Return to the time slice included in the start time and end time.', 'zh_CN':'时间
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05，最后一个时间片是（yyyy-MM-dd+1）00:00。
  // 2.查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 01，最后一个时间片是（yyyy-MM-dd+1）&nbsp;00。
  // 3.返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Flow value in Byte', 'zh_CN':'流量值，单位为Byte'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData) SetTimestamp(v string) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData) SetValue(v string) *ReportFlowIspProvinceByteServiceReportFlowIspProvinceByteServiceResponseResultIspDataProvinceDataFlowData {
  s.Value = &v
  return s
}

type ReportFlowIspProvinceByteServicePaths struct {
}

func (s ReportFlowIspProvinceByteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServicePaths) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceByteServiceParameters struct {
}

func (s ReportFlowIspProvinceByteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceByteServiceRequestHeader struct {
}

func (s ReportFlowIspProvinceByteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceByteServiceResponseHeader struct {
}

func (s ReportFlowIspProvinceByteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceByteServiceResponseHeader) GoString() string {
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
  Result []*QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponse) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponse) SetResult(v []*QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult) *QueryOutputTrafficUnderShieldPoPResponse {
  s.Result = v
  return s
}

type QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic", "zh_CN":"总流量"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量值数据"}
  FlowData []*QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult) SetDomain(v string) *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult) SetTotalFlow(v string) *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult) SetFlowData(v []*QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData) *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResult {
  s.FlowData = v
  return s
}

type QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits &nbsp; of decimals allowed", "zh_CN":"流量值,单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData) SetTimestamp(v string) *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData) SetValue(v string) *QueryOutputTrafficUnderShieldPoPQueryOutputTrafficUnderShieldPoPResponseResultFlowData {
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




type ReportFlowHitRateIspProvinceServiceRequest struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must   be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days;
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last   24 hour is queried", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天;(可联系技术支持调整)
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it&rsquo;s greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number   limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5-minute granularity, 1h: 1-hour granularity", "zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度"}
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
  // 1.Options are domain, province, isp, and more than one value can be entered;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetDateFrom(v string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetDateTo(v string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetDomain(v []*string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetDataInterval(v string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetProvince(v []*string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetIsp(v []*string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportFlowHitRateIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowHitRateIspProvinceServiceResponse struct {
  Result []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceServiceResponse) SetResult(v []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult) *ReportFlowHitRateIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult) SetDomain(v string) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult) SetIspData(v []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  HitRateData []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData `json:"hitRateData,omitempty" xml:"hitRateData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData) SetHitRateData(v []*ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceData {
  s.HitRateData = v
  return s
}

type ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData struct     {
  // {"en":"Time,
  // 1.When   the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;
  // 3.Return the time slice contained in start time and the time slice contained in end   time.", "zh_CN":"时间,
  // 1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Byte hit rate. Four digits of &nbsp; decimals are allowed", "zh_CN":"字节命中率,保留4位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Hit traffic. Unit is MB and 2 &nbsp; digits of decimals are allowed", "zh_CN":"命中流量值,单位MB,保留2位小数;
  // 默认不返回该字段,需要返回的请联系技术支持"}
  HitFlow *string `json:"hitFlow,omitempty" xml:"hitFlow,omitempty" require:"true"`
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) SetTimestamp(v string) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) SetValue(v string) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData {
  s.Value = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) SetHitFlow(v string) *ReportFlowHitRateIspProvinceServiceReportFlowHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData {
  s.HitFlow = &v
  return s
}

type ReportFlowHitRateIspProvinceServicePaths struct {
}

func (s ReportFlowHitRateIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportFlowHitRateIspProvinceServiceParameters struct {
}

func (s ReportFlowHitRateIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowHitRateIspProvinceServiceRequestHeader struct {
}

func (s ReportFlowHitRateIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowHitRateIspProvinceServiceResponseHeader struct {
}

func (s ReportFlowHitRateIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceServiceResponseHeader) GoString() string {
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
  Result []*ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowExactDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceResponse) SetResult(v []*ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult) *ReportFlowExactDomainServiceResponse {
  s.Result = v
  return s
}

type ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult struct     {
  // {'en':'Exact domain name', 'zh_CN':'精确域名'}
  ExactDomain *string `json:"exactDomain,omitempty" xml:"exactDomain,omitempty" require:"true"`
  FlowData []*ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult) SetExactDomain(v string) *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult {
  s.ExactDomain = &v
  return s
}

func (s *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult) SetFlowData(v []*ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData) *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResult {
  s.FlowData = v
  return s
}

type ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData struct     {
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

func (s ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData) SetTimestamp(v string) *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData) SetValue(v string) *ReportFlowExactDomainServiceReportFlowExactDomainServiceResponseResultFlowData {
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




type ReportFlowHitRateIspProvinceByteServiceRequest struct {
  // {"en":"Starting time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.Must be greater than the current time -183 days, and less than the current time and dateTo;
  // 
  // 3.The difference between dateFrom and dateTo cannot exceed 7 days(technical support can be contacted to adjust);
  // 
  // 4.DateFrom and dateTo are either passed or not passed;
  // 
  // 5.DateFrom and dateTo are not passed, the default query data for the past 24 hours", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"Starting time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.Must be greater than the current time -183 days, and less than the current time and dateTo;
  // 
  // 3.The difference between dateFrom and dateTo cannot exceed 7 days;
  // 
  // 4.DateFrom and dateTo are either passed or not passed;
  // 
  // 5.DateFrom and dateTo are not passed, the default query data for the past 24 hours", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"The domain name and the number of domain names are adjustable according to the account number. The default is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5 minutes granularity, 1h: 1 hour granularity", "zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度"}
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
  // {"en":"Grouping dimension
  // 
  // 1. The optional values are domain,province, and isp, which can pass multiple values.
  // 
  // 2. If there is an incoming, the detailed data will be displayed according to the dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceByteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetDateFrom(v string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetDateTo(v string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetDomain(v []*string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetDataInterval(v string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetProvince(v []*string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetIsp(v []*string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceRequest) SetGroupBy(v []*string) *ReportFlowHitRateIspProvinceByteServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowHitRateIspProvinceByteServiceResponse struct {
  Result []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceByteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceByteServiceResponse) SetResult(v []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult) *ReportFlowHitRateIspProvinceByteServiceResponse {
  s.Result = v
  return s
}

type ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult) SetDomain(v string) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult) SetIspData(v []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData) SetIsp(v string) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData) SetProvinceData(v []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  HitRateData []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData `json:"hitRateData,omitempty" xml:"hitRateData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData) SetHitRateData(v []*ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceData {
  s.HitRateData = v
  return s
}

type ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData struct     {
  // {"en":"time,
  // 
  // 1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00;
  // 
  // 2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1)00;
  // 
  // 3. Return to the time slice included in the start time and end time.", "zh_CN":"时间,
  // 1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00;
  // 2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 01,最后一个时间片是(yyyy-MM-dd+1)00;
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Byte hit ratio, retain 4 decimal places", "zh_CN":"字节命中率,保留4位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Hit traffic value, unit Byte", "zh_CN":"命中流量值,单位Byte"}
  HitFlow *string `json:"hitFlow,omitempty" xml:"hitFlow,omitempty" require:"true"`
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData) GoString() string {
  return s.String()
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData) SetTimestamp(v string) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData) SetValue(v string) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData {
  s.Value = &v
  return s
}

func (s *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData) SetHitFlow(v string) *ReportFlowHitRateIspProvinceByteServiceReportFlowHitRateIspProvinceByteServiceResponseResultIspDataProvinceDataHitRateData {
  s.HitFlow = &v
  return s
}

type ReportFlowHitRateIspProvinceByteServicePaths struct {
}

func (s ReportFlowHitRateIspProvinceByteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServicePaths) GoString() string {
  return s.String()
}

type ReportFlowHitRateIspProvinceByteServiceParameters struct {
}

func (s ReportFlowHitRateIspProvinceByteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowHitRateIspProvinceByteServiceRequestHeader struct {
}

func (s ReportFlowHitRateIspProvinceByteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowHitRateIspProvinceByteServiceResponseHeader struct {
}

func (s ReportFlowHitRateIspProvinceByteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowHitRateIspProvinceByteServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainTotalTrafficRequest struct {
  // {"en":"Domain list.
  //   Domain number limits can be adjusted depending on different accounts. The default value is 20(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表
  //   1.域名个数限制根据账号可调,默认为20个(可联系技术支持下单调整);"}
  QueryDomainTotalTrafficDomainList *QueryDomainTotalTrafficDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryDomainTotalTrafficRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficRequest) SetDomainList(v *QueryDomainTotalTrafficDomainList) *QueryDomainTotalTrafficRequest {
  s.QueryDomainTotalTrafficDomainList = v
  return s
}

type QueryDomainTotalTrafficDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalTrafficDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficDomainList) SetDomainName(v []*string) *QueryDomainTotalTrafficDomainList {
  s.DomainName = v
  return s
}

type QueryDomainTotalTrafficResponse struct {
  // {"en":"Total traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *int `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量数据"}
  FlowData []*QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalTrafficResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficResponse) SetFlowSummary(v int) *QueryDomainTotalTrafficResponse {
  s.FlowSummary = &v
  return s
}

func (s *QueryDomainTotalTrafficResponse) SetFlowData(v []*QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData) *QueryDomainTotalTrafficResponse {
  s.FlowData = v
  return s
}

type QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData struct     {
  // {"en":"Date
  //   1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.
  //   2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.
  //   3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.Return the time slice contained in start time and the time slice contained in end time", "zh_CN":"时间
  //   1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  //   2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  //   3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。
  //   4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"流量,保留2位小数,单位为MB"}
  Flow *int `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData) SetTimestamp(v string) *QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData) SetFlow(v int) *QueryDomainTotalTrafficQueryDomainTotalTrafficResponseFlowData {
  s.Flow = &v
  return s
}

type QueryDomainTotalTrafficPaths struct {
}

func (s QueryDomainTotalTrafficPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficPaths) GoString() string {
  return s.String()
}

type QueryDomainTotalTrafficParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and 'dateTo';
  // 3.Period between 'dataFrom' and 'dateTo' cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天;4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than 'dateFrom';
  // 3.If it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDomainTotalTrafficParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficParameters) SetDateFrom(v string) *QueryDomainTotalTrafficParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainTotalTrafficParameters) SetDateTo(v string) *QueryDomainTotalTrafficParameters {
  s.DateTo = &v
  return s
}

func (s *QueryDomainTotalTrafficParameters) SetType(v string) *QueryDomainTotalTrafficParameters {
  s.Type = &v
  return s
}

type QueryDomainTotalTrafficRequestHeader struct {
}

func (s QueryDomainTotalTrafficRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainTotalTrafficResponseHeader struct {
}

func (s QueryDomainTotalTrafficResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficResponseHeader) GoString() string {
  return s.String()
}




type FlowChannelRequest struct {
  // {"en":"cust_en_name ", "zh_CN":"客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
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
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1.If there isp multiple inputs,use ';' as demimeter.
  // 2.optional values of isp: refers to the ISP-section of appendix.
  // 3.If not specified,means all the isp.", "zh_CN":"&nbsp;要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"flowInfo:displays bandwidth peak, peak time,and total flow information;", "zh_CN":"flowInfo ：展示带宽峰值、峰值时间、总流量信息；"}
  OptionalFields *string `json:"optionalFields,omitempty" xml:"optionalFields,omitempty"`
}

func (s FlowChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelRequest) GoString() string {
  return s.String()
}

func (s *FlowChannelRequest) SetCust(v string) *FlowChannelRequest {
  s.Cust = &v
  return s
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

func (s *FlowChannelRequest) SetOptionalFields(v string) *FlowChannelRequest {
  s.OptionalFields = &v
  return s
}

type FlowChannelResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *FlowChannelFlowChannelResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s FlowChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponse) GoString() string {
  return s.String()
}

func (s *FlowChannelResponse) SetProvider(v *FlowChannelFlowChannelResponseProvider) *FlowChannelResponse {
  s.Provider = v
  return s
}

type FlowChannelFlowChannelResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道流量数据'}
  Date *FlowChannelFlowChannelResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s FlowChannelFlowChannelResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelFlowChannelResponseProvider) GoString() string {
  return s.String()
}

func (s *FlowChannelFlowChannelResponseProvider) SetName(v string) *FlowChannelFlowChannelResponseProvider {
  s.Name = &v
  return s
}

func (s *FlowChannelFlowChannelResponseProvider) SetType(v string) *FlowChannelFlowChannelResponseProvider {
  s.Type = &v
  return s
}

func (s *FlowChannelFlowChannelResponseProvider) SetResultType(v string) *FlowChannelFlowChannelResponseProvider {
  s.ResultType = &v
  return s
}

func (s *FlowChannelFlowChannelResponseProvider) SetDate(v *FlowChannelFlowChannelResponseProviderDate) *FlowChannelFlowChannelResponseProvider {
  s.Date = v
  return s
}

type FlowChannelFlowChannelResponseProviderDate struct {
  // {'en':'name', 'zh_CN':'日期'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *FlowChannelFlowChannelResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s FlowChannelFlowChannelResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelFlowChannelResponseProviderDate) GoString() string {
  return s.String()
}

func (s *FlowChannelFlowChannelResponseProviderDate) SetName(v string) *FlowChannelFlowChannelResponseProviderDate {
  s.Name = &v
  return s
}

func (s *FlowChannelFlowChannelResponseProviderDate) SetChannel(v *FlowChannelFlowChannelResponseProviderDateChannel) *FlowChannelFlowChannelResponseProviderDate {
  s.Channel = v
  return s
}

type FlowChannelFlowChannelResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'流量数据'}
  Flow []*FlowChannelFlowChannelResponseProviderDateChannelFlow `json:"flow,omitempty" xml:"flow,omitempty" require:"true" type:"Repeated"`
}

func (s FlowChannelFlowChannelResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelFlowChannelResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *FlowChannelFlowChannelResponseProviderDateChannel) SetName(v string) *FlowChannelFlowChannelResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *FlowChannelFlowChannelResponseProviderDateChannel) SetFlow(v []*FlowChannelFlowChannelResponseProviderDateChannelFlow) *FlowChannelFlowChannelResponseProviderDateChannel {
  s.Flow = v
  return s
}

type FlowChannelFlowChannelResponseProviderDateChannelFlow struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'流量'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s FlowChannelFlowChannelResponseProviderDateChannelFlow) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelFlowChannelResponseProviderDateChannelFlow) GoString() string {
  return s.String()
}

func (s *FlowChannelFlowChannelResponseProviderDateChannelFlow) SetTime(v string) *FlowChannelFlowChannelResponseProviderDateChannelFlow {
  s.Time = &v
  return s
}

func (s *FlowChannelFlowChannelResponseProviderDateChannelFlow) SetText(v string) *FlowChannelFlowChannelResponseProviderDateChannelFlow {
  s.Text = &v
  return s
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

type FlowChannelRequestHeader struct {
}

func (s FlowChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelRequestHeader) GoString() string {
  return s.String()
}

type FlowChannelResponseHeader struct {
}

func (s FlowChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowChannelResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowIspProvinceTotalServiceRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如 2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔:24小时(可联系技术支持调整)，即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3. Domain name exceeding limit, misstatement", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.域名超过上限，提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity: 
  // 1. default 1m
  // 2. 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度:
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
  // 1. Optional value flow, request,bandwidth
  // 2. Default flow
  // 3. Flow: Flow, keep two decimal places;
  // 4. Request: number of Request;
  // 5. bandwidth: bandwidth, keep two decimal places", "zh_CN":"查询维度:
  // 1.可选值 flow、request、bandwidth
  // 2.传默认 flow
  // 3.flow:流量，保留两位小数;
  // 4.request:请求数;
  // 5.bandwidth:带宽，保留两位小数"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional: domain, all, If it is empty, it defaults to returning by domain dimension;
  // If all is passed, merge and return according to the query domain name.", "zh_CN":"可选项：domain、all, 为空则默认为按domain维度返回;
  // 若传递all，则按查询域名合并返回"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s ReportFlowIspProvinceTotalServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetDateFrom(v string) *ReportFlowIspProvinceTotalServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetDateTo(v string) *ReportFlowIspProvinceTotalServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetDomain(v []*string) *ReportFlowIspProvinceTotalServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetDataInterval(v string) *ReportFlowIspProvinceTotalServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetProvince(v []*string) *ReportFlowIspProvinceTotalServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetIsp(v []*string) *ReportFlowIspProvinceTotalServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetQueryBy(v string) *ReportFlowIspProvinceTotalServiceRequest {
  s.QueryBy = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceRequest) SetGroupBy(v string) *ReportFlowIspProvinceTotalServiceRequest {
  s.GroupBy = &v
  return s
}

type ReportFlowIspProvinceTotalServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"数据结果"}
  Data []*ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceTotalServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceTotalServiceResponse) SetCode(v string) *ReportFlowIspProvinceTotalServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceResponse) SetMessage(v string) *ReportFlowIspProvinceTotalServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceResponse) SetData(v []*ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData) *ReportFlowIspProvinceTotalServiceResponse {
  s.Data = v
  return s
}

type ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"detailList", "zh_CN":"数据明细"}
  DetailList []*ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData) SetDomain(v string) *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData) SetDetailList(v []*ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList) *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseData {
  s.DetailList = v
  return s
}

type ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"1. Flow: Flow, keep two decimal places;
  //             2. Bandwidth: Bandwidth, keep two decimals;
  //             3. Request: number of Request", "zh_CN":"1.flow：流量，保留两位小数；
  //             2.bandwidth: 带宽，保留两位小数；
  //             3.request：请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList) SetTimestamp(v string) *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList) SetValue(v string) *ReportFlowIspProvinceTotalServiceReportFlowIspProvinceTotalServiceResponseDataDetailList {
  s.Value = &v
  return s
}

type ReportFlowIspProvinceTotalServicePaths struct {
}

func (s ReportFlowIspProvinceTotalServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServicePaths) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceTotalServiceParameters struct {
}

func (s ReportFlowIspProvinceTotalServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceTotalServiceRequestHeader struct {
}

func (s ReportFlowIspProvinceTotalServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceTotalServiceResponseHeader struct {
}

func (s ReportFlowIspProvinceTotalServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceTotalServiceResponseHeader) GoString() string {
  return s.String()
}




type WctQueryRequest struct {
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
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"trans;hevc;audio_trans;encap_trans;vframe;file_op;trans_nbhd;hevc_nbhd;video_enhance;drm.", "zh_CN":"转码操作类型：trans(H.264视频转码);hevc(H.265视频转码);audio_trans(音频转码);encap_trans(转封装);vframe(截图);file_op(文件处理);trans_nbhd(H.264智控高清视频转码);hevc_nbhd(H.265智控高清视频转码);video_enhance(AI视频增强);drm(DRM加密)"}
  Transcoding *string `json:"transcoding,omitempty" xml:"transcoding,omitempty"`
  // {"en":"sd240;sd480;sd720;hd1080;2k;4k.", "zh_CN":"梯度: sd240;sd480;sd720;hd1080;2k;4k"}
  TranscodingType *string `json:"transcodingType,omitempty" xml:"transcodingType,omitempty"`
  // {"en":"space .", "zh_CN":"空间，多个值请用英文分号“;”分割"}
  Space *string `json:"space,omitempty" xml:"space,omitempty"`
  // {"en":"provider .", "zh_CN":"功能来源，提供方。多个值请用英文分号“;”分割"}
  Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
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

type WctQueryResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *WctQueryWctQueryResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponse) GoString() string {
  return s.String()
}

func (s *WctQueryResponse) SetProvider(v *WctQueryWctQueryResponseProvider) *WctQueryResponse {
  s.Provider = v
  return s
}

type WctQueryWctQueryResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'转码数据'}
  Date *WctQueryWctQueryResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryWctQueryResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s WctQueryWctQueryResponseProvider) GoString() string {
  return s.String()
}

func (s *WctQueryWctQueryResponseProvider) SetName(v string) *WctQueryWctQueryResponseProvider {
  s.Name = &v
  return s
}

func (s *WctQueryWctQueryResponseProvider) SetType(v string) *WctQueryWctQueryResponseProvider {
  s.Type = &v
  return s
}

func (s *WctQueryWctQueryResponseProvider) SetDate(v *WctQueryWctQueryResponseProviderDate) *WctQueryWctQueryResponseProvider {
  s.Date = v
  return s
}

type WctQueryWctQueryResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'customerId', 'zh_CN':'客户id'}
  CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *string `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true"`
  // {'en':'transcodingType', 'zh_CN':'梯度'}
  TranscodingType *string `json:"transcodingType,omitempty" xml:"transcodingType,omitempty" require:"true"`
  // {'en':'unit', 'zh_CN':'单位'}
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'总数'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'wct', 'zh_CN':'转码数据'}
  Wct *WctQueryWctQueryResponseProviderDateWct `json:"wct,omitempty" xml:"wct,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryWctQueryResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s WctQueryWctQueryResponseProviderDate) GoString() string {
  return s.String()
}

func (s *WctQueryWctQueryResponseProviderDate) SetStartdate(v string) *WctQueryWctQueryResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetEnddate(v string) *WctQueryWctQueryResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetCustomerId(v string) *WctQueryWctQueryResponseProviderDate {
  s.CustomerId = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetTranscoding(v string) *WctQueryWctQueryResponseProviderDate {
  s.Transcoding = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetTranscodingType(v string) *WctQueryWctQueryResponseProviderDate {
  s.TranscodingType = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetUnit(v string) *WctQueryWctQueryResponseProviderDate {
  s.Unit = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetTotal(v string) *WctQueryWctQueryResponseProviderDate {
  s.Total = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDate) SetWct(v *WctQueryWctQueryResponseProviderDateWct) *WctQueryWctQueryResponseProviderDate {
  s.Wct = v
  return s
}

type WctQueryWctQueryResponseProviderDateWct struct {
  // {'en':'type', 'zh_CN':'类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'总数'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Data []*WctQueryWctQueryResponseProviderDateWctData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s WctQueryWctQueryResponseProviderDateWct) String() string {
  return tea.Prettify(s)
}

func (s WctQueryWctQueryResponseProviderDateWct) GoString() string {
  return s.String()
}

func (s *WctQueryWctQueryResponseProviderDateWct) SetType(v string) *WctQueryWctQueryResponseProviderDateWct {
  s.Type = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDateWct) SetTotal(v string) *WctQueryWctQueryResponseProviderDateWct {
  s.Total = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDateWct) SetData(v []*WctQueryWctQueryResponseProviderDateWctData) *WctQueryWctQueryResponseProviderDateWct {
  s.Data = v
  return s
}

type WctQueryWctQueryResponseProviderDateWctData struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s WctQueryWctQueryResponseProviderDateWctData) String() string {
  return tea.Prettify(s)
}

func (s WctQueryWctQueryResponseProviderDateWctData) GoString() string {
  return s.String()
}

func (s *WctQueryWctQueryResponseProviderDateWctData) SetTime(v string) *WctQueryWctQueryResponseProviderDateWctData {
  s.Time = &v
  return s
}

func (s *WctQueryWctQueryResponseProviderDateWctData) SetText(v string) *WctQueryWctQueryResponseProviderDateWctData {
  s.Text = &v
  return s
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

type WctQueryRequestHeader struct {
}

func (s WctQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s WctQueryRequestHeader) GoString() string {
  return s.String()
}

type WctQueryResponseHeader struct {
}

func (s WctQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainTrafficUnderSpecificSettlementAreaRequest struct {
  // {"en":"Start time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 31 days;4.You can only query data for the last 2 years(technical support can be contacted to adjust).", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain list, and all domains are queried if this field is not specified;", "zh_CN":"域名列表,不传递则查询全部域名;"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity
  // 
  // 1.Options: 5m(5 minutes), 1h(1 hour) and 1d(1 day);
  // 2.Default value of 1d is used if the field is not specified;
  // 3.If 5m is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.", "zh_CN":"数据粒度
  // 1.可选值为:5m(5分钟)、1h(1小时)、1d(1天);
  // 2.不传时默认为1d;
  // 3.传递5m时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Queried area
  // 1.Use  English semicolon to separate two areas;
  // 2.Options: cn, nc, ov, apac, am, euna, emea, sa, af, au, hk, tw 
  //  ...;", "zh_CN":"查询区域
  // 1.多个区域使用英文分号分隔
  // 2.可选值为:cn、nc、ov、apac、am、euna、emea、sa、af、au、hk、tw等;"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaRequest) SetDateFrom(v string) *QueryDomainTrafficUnderSpecificSettlementAreaRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaRequest) SetDateTo(v string) *QueryDomainTrafficUnderSpecificSettlementAreaRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaRequest) SetDomain(v []*string) *QueryDomainTrafficUnderSpecificSettlementAreaRequest {
  s.Domain = v
  return s
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaRequest) SetDataInterval(v string) *QueryDomainTrafficUnderSpecificSettlementAreaRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaRequest) SetAreaCode(v string) *QueryDomainTrafficUnderSpecificSettlementAreaRequest {
  s.AreaCode = &v
  return s
}

type QueryDomainTrafficUnderSpecificSettlementAreaResponse struct {
  // {"en":"Total traffic, 2 decimal places, in MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *string `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  FlowData []*QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaResponse) SetFlowSummary(v string) *QueryDomainTrafficUnderSpecificSettlementAreaResponse {
  s.FlowSummary = &v
  return s
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaResponse) SetFlowData(v []*QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData) *QueryDomainTrafficUnderSpecificSettlementAreaResponse {
  s.FlowData = v
  return s
}

type QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData struct     {
  // {"en":"time
  // 1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH: MM; Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is yyyy-MM-dd 24:00.
  // 2. When the data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 01, and the last time slice is yyyy-MM-dd 24.
  // 3. when the data granularity of query is daily, the format is yyyy-MM-dd; The value of data for each time slice represents the value of the data within that day;
  // 4. return the time slice contained in the start and end times.", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic, 2 decimal places, in MB", "zh_CN":"流量,保留2位小数,单位为MB"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData) GoString() string {
  return s.String()
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData) SetTimestamp(v string) *QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData) SetFlow(v string) *QueryDomainTrafficUnderSpecificSettlementAreaQueryDomainTrafficUnderSpecificSettlementAreaResponseFlowData {
  s.Flow = &v
  return s
}

type QueryDomainTrafficUnderSpecificSettlementAreaPaths struct {
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaPaths) GoString() string {
  return s.String()
}

type QueryDomainTrafficUnderSpecificSettlementAreaParameters struct {
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaParameters) GoString() string {
  return s.String()
}

type QueryDomainTrafficUnderSpecificSettlementAreaRequestHeader struct {
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainTrafficUnderSpecificSettlementAreaResponseHeader struct {
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTrafficUnderSpecificSettlementAreaResponseHeader) GoString() string {
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
  Result []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceResponse) SetResult(v []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult) *ReportFlowIspProvinceShortTimeServiceResponse {
  s.Result = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"ISP数据"}
  IspData []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult) SetDomain(v string) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult) SetIspData(v []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份数据"}
  ProvinceData []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData) SetIsp(v string) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData) SetProvinceData(v []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"flowdata", "zh_CN":"流量数据"}
  FlowData []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData) SetFlowData(v []*ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceData {
  s.FlowData = v
  return s
}

type ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits   of decimals allowed", "zh_CN":"流量值,单位为MB,保留两位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Bandwidth value. Unit is Mbps and   2 digits of decimals allowed", "zh_CN":"带宽值,单位为Mbps,保留两位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) SetTimestamp(v string) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) SetValue(v string) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData {
  s.Value = &v
  return s
}

func (s *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData) SetBandwidth(v string) *ReportFlowIspProvinceShortTimeServiceReportFlowIspProvinceShortTimeServiceResponseResultIspDataProvinceDataFlowData {
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
  Result []*QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponse) SetResult(v []*QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult) *QueryEdgeByteHitRatioServiceResponse {
  s.Result = v
  return s
}

type QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间，格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total edge node hit ratio", "zh_CN":"总边缘节点命中率的平均值,2位小数"}
  TotalAvg *string `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  HitRatioDatas []*QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult) SetRealDate(v string) *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult {
  s.RealDate = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult) SetTotalAvg(v string) *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult {
  s.TotalAvg = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult) SetHitRatioDatas(v []*QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResult {
  s.HitRatioDatas = v
  return s
}

type QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas struct     {
  // {"en":"timestamp", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"edge node hit ratio,keep 4 decimal places", "zh_CN":"边缘节点缓存字节命中率，保留4位小数"}
  EdgeHitRatio *string `json:"edgeHitRatio,omitempty" xml:"edgeHitRatio,omitempty" require:"true"`
}

func (s QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) SetTimestamp(v string) *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) SetEdgeHitRatio(v string) *QueryEdgeByteHitRatioServiceQueryEdgeByteHitRatioServiceResponseResultHitRatioDatas {
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




type QueryBacktoOriginTrafficAndRequestRequest struct {
  // {'en':'Start time: 
  // 1. Start time: time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (December 2rd, 2016, 10:00 a.m., Beijing Time); 
  // 2. Not greater than the current time 
  // 3. The data can be queried is the last week (183 days).', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time: 
  // 1. The time format is 2016-12-02T10:00:00+08:00 
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if one field is filled and one is left empty, then exception will occur. 
  // 4. Maximum query time interval allowed: 1 day by default, that is, the difference between dateFrom and dateTo cannot exceed 1 day (you can contact technical support to adjust it, the maximum adjustment is 31 days)', 'zh_CN':'结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 5.允许查询最大时间间隔：默认1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大调整到31天）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity: 
  // 1. Support 5m (granularity of 5 minutes) and 1m
  // 2. 5m by default if the value is empty', 'zh_CN':'数据粒度：
  // 1、支持5m（5分钟）。和 1m（1分钟）
  // 2、不传默认5m。'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Domain name: 
  // 1. The default upper limit of domains that can be entered is 200 (if you want to adjust, please, contact technical support); 
  // 2. All domains under the account will be queried if this input parameter is not specified. But if the number of domains under the account exceeds the limits, no query will be executed (Error)', 'zh_CN':'域名：
  // 1、可传递域名数量上限默认为200个（可联系技术支持调整）；
  // 2、未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'Grouping keywords:
  // 1. By default, data will be displayed by group;
  // 2. If there are keywords entered, value details shall be displayed by keywords; If groupBy is specified as domain, it means the results are returned according to domains. 
  // 3. Only domain can be specified', 'zh_CN':'分组关键词：
  // 1、默认聚合展示；
  // &nbsp; 2、传入关键词则代表需要按照关键词对应的值展示明细；
  // 例如groupBy传domain，则代表返回按照domain明细展开。
  // 3、只能传递domain。'}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {'en':'Optional values 0, 1. 
  // Input 0 returns all data, input 1 only returns source site data, default is0&nbsp;', 'zh_CN':'可选值 0、1 。入参 0 则返回全部数据，入参 1 则只返回回源站数据，默认为 0'}
  BacksrcOnly *int `json:"backsrcOnly,omitempty" xml:"backsrcOnly,omitempty"`
}

func (s QueryBacktoOriginTrafficAndRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestRequest) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDateFrom(v string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDateTo(v string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDataInterval(v string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDomain(v []*string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.Domain = v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetGroupBy(v []*string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.GroupBy = v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetBacksrcOnly(v int) *QueryBacktoOriginTrafficAndRequestRequest {
  s.BacksrcOnly = &v
  return s
}

type QueryBacktoOriginTrafficAndRequestResponse struct {
  Result []*QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBacktoOriginTrafficAndRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestResponse) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestResponse) SetResult(v []*QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) *QueryBacktoOriginTrafficAndRequestResponse {
  s.Result = v
  return s
}

type QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'totalFlow, unit :MB, 2 decimal places reserved, example (74099.92)', 'zh_CN':'总流量，单位MB，保留两位小时'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {'en':'totalRequest', 'zh_CN':'总请求数'}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {'en':'peak of Request', 'zh_CN':'请求数峰值'}
  PeakRequest *string `json:"peakRequest,omitempty" xml:"peakRequest,omitempty" require:"true"`
  // {'en':'peaktime of Request', 'zh_CN':'请求数峰值时间'}
  PeakRequestTime *string `json:"peakRequestTime,omitempty" xml:"peakRequestTime,omitempty" require:"true"`
  // {'en':'peakBandwidth, unit :Mbps, 2 decimal places reserved, example (74099.92)', 'zh_CN':'带宽峰值，单位Mbps，保留2位小数，示例 （931556.21）'}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {'en':'Peak time of Bandwidth', 'zh_CN':'带宽峰值时间'}
  PeakBandwidthTime *string `json:"peakBandwidthTime,omitempty" xml:"peakBandwidthTime,omitempty" require:"true"`
  FlowRequestOriginData []*QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData `json:"flowRequestOriginData,omitempty" xml:"flowRequestOriginData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetDomain(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetTotalFlow(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetTotalRequest(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetPeakRequest(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakRequest = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetPeakRequestTime(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakRequestTime = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetPeakBandwidth(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetPeakBandwidthTime(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakBandwidthTime = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult) SetFlowRequestOriginData(v []*QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResult {
  s.FlowRequestOriginData = v
  return s
}

type QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData struct     {
  // {'en':'1. When the querying data granularity is 5m, then the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00. 
  //         2. Return the time slices that contained in start time and in end time.', 'zh_CN':'1.查询的数据粒度为5m时，格式为yyyy-MM-dd &nbsp; HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） &nbsp; 00:00。
  //         2.返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Traffic unit is MB and keep the number to two decimal places', 'zh_CN':'流量值，单位MB，保留2位小数；'}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {'en':'Bandwidth, unit: Mbps, 2 decimal places reserved, example (931556.21)', 'zh_CN':'带宽，单位：Mbps，保留两位小数'}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {'en':'Total number of requests', 'zh_CN':'请求数'}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetTimestamp(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Timestamp = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetFlow(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Flow = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetBandwidth(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Bandwidth = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetRequest(v string) *QueryBacktoOriginTrafficAndRequestQueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Request = &v
  return s
}

type QueryBacktoOriginTrafficAndRequestPaths struct {
}

func (s QueryBacktoOriginTrafficAndRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestPaths) GoString() string {
  return s.String()
}

type QueryBacktoOriginTrafficAndRequestParameters struct {
}

func (s QueryBacktoOriginTrafficAndRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestParameters) GoString() string {
  return s.String()
}

type QueryBacktoOriginTrafficAndRequestRequestHeader struct {
}

func (s QueryBacktoOriginTrafficAndRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestRequestHeader) GoString() string {
  return s.String()
}

type QueryBacktoOriginTrafficAndRequestResponseHeader struct {
}

func (s QueryBacktoOriginTrafficAndRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowRequestMultiIPVersionServiceRequest struct {
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
  // {"en":"Data granularity: 
  // 	1.Support for 1m(1 minutes), 5m (5 minutes), 1h (1 hour) 
  // 	2.Default 5m
  // ", "zh_CN":"数据粒度:
  // 1、支持1m(1分钟)、5m(5分钟)、1h(1小时)
  // 2、不传默认5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"1. Optional value:bandwidth, request 2. Default bandwidth", "zh_CN":"1.可选值:bandwidth、request
  // 2.不传默认bandwidth"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"1.Selection:domain 2.If groupBy left empty, merge date of all domains", "zh_CN":"可选值:domain
  // 不传默认聚合所有频道数据"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowRequestMultiIPVersionServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowRequestMultiIPVersionServiceRequest) SetDateFrom(v string) *ReportFlowRequestMultiIPVersionServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceRequest) SetDateTo(v string) *ReportFlowRequestMultiIPVersionServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceRequest) SetDomain(v []*string) *ReportFlowRequestMultiIPVersionServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceRequest) SetDataInterval(v string) *ReportFlowRequestMultiIPVersionServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceRequest) SetQueryBy(v string) *ReportFlowRequestMultiIPVersionServiceRequest {
  s.QueryBy = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceRequest) SetGroupBy(v []*string) *ReportFlowRequestMultiIPVersionServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowRequestMultiIPVersionServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowRequestMultiIPVersionServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowRequestMultiIPVersionServiceResponse) SetCode(v string) *ReportFlowRequestMultiIPVersionServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceResponse) SetMessage(v string) *ReportFlowRequestMultiIPVersionServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceResponse) SetData(v []*ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData) *ReportFlowRequestMultiIPVersionServiceResponse {
  s.Data = v
  return s
}

type ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData struct     {
  // {"en":"Domain. If merge date of all domains will not return domain", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData) SetDomain(v string) *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData) SetDetailList(v []*ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList) *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseData {
  s.DetailList = v
  return s
}

type ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList struct     {
  // {"en":"timestamp,Returns the timestamp between the start time and end time. 
  // 				 Time format: 
  // 				 Minutes: yyyy-MM-dd HH:mm 
  // 				 Hours: yyyy-MM-dd HH", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。
  // 时间格式:
  // 1/5分钟:yyyy-MM-dd HH:mm
  // 1小时:yyyy-MM-dd HH"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Ipv4 (Bandwidth unit is Mbps, number of requests is in units", "zh_CN":"IPv4数据（带宽单位 Mbps，请求数单位 个）"}
  V4Value *string `json:"v4Value,omitempty" xml:"v4Value,omitempty" require:"true"`
  // {"en":"Ipv6 (Bandwidth unit is Mbps, number of requests is in units", "zh_CN":"IPv6数据（带宽单位 Mbps，请求数单位 个）"}
  V6Value *string `json:"v6Value,omitempty" xml:"v6Value,omitempty" require:"true"`
}

func (s ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList) SetTimestamp(v string) *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList) SetV4Value(v string) *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList {
  s.V4Value = &v
  return s
}

func (s *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList) SetV6Value(v string) *ReportFlowRequestMultiIPVersionServiceReportFlowRequestMultiIPVersionServiceResponseDataDetailList {
  s.V6Value = &v
  return s
}

type ReportFlowRequestMultiIPVersionServicePaths struct {
}

func (s ReportFlowRequestMultiIPVersionServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServicePaths) GoString() string {
  return s.String()
}

type ReportFlowRequestMultiIPVersionServiceParameters struct {
}

func (s ReportFlowRequestMultiIPVersionServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowRequestMultiIPVersionServiceRequestHeader struct {
}

func (s ReportFlowRequestMultiIPVersionServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowRequestMultiIPVersionServiceResponseHeader struct {
}

func (s ReportFlowRequestMultiIPVersionServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowRequestMultiIPVersionServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportFlowIspProvinceServiceRequest struct {
  // {"en":"Start time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried;", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术人员调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据;"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递该入参时查询账号下所有域名,但当账号下域名数量超过限制时不可查询(报错)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5-minute  granularity, 1h: 1-hour granularity", "zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度"}
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
  // 1.Options are domain, province, isp, and more than one value can be entered;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceServiceRequest) SetDateFrom(v string) *ReportFlowIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowIspProvinceServiceRequest) SetDateTo(v string) *ReportFlowIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowIspProvinceServiceRequest) SetDomain(v []*string) *ReportFlowIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowIspProvinceServiceRequest) SetDataInterval(v string) *ReportFlowIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowIspProvinceServiceRequest) SetProvince(v []*string) *ReportFlowIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowIspProvinceServiceRequest) SetIsp(v []*string) *ReportFlowIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportFlowIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowIspProvinceServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceServiceResponse) SetResult(v []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult) *ReportFlowIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"ISP数据"}
  IspData []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult) SetDomain(v string) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult) SetIspData(v []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份数据"}
  ProvinceData []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"flowdata", "zh_CN":"流量数据"}
  FlowData []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData) SetFlowData(v []*ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceData {
  s.FlowData = v
  return s
}

type ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits   of decimals allowed", "zh_CN":"流量值,单位为MB,保留两位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Bandwidth value. Unit is Mbps and   2 digits of decimals allowed", "zh_CN":"带宽值,单位为Mbps,保留两位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData) SetTimestamp(v string) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData) SetValue(v string) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData {
  s.Value = &v
  return s
}

func (s *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData) SetBandwidth(v string) *ReportFlowIspProvinceServiceReportFlowIspProvinceServiceResponseResultIspDataProvinceDataFlowData {
  s.Bandwidth = &v
  return s
}

type ReportFlowIspProvinceServicePaths struct {
}

func (s ReportFlowIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceServiceParameters struct {
}

func (s ReportFlowIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceServiceRequestHeader struct {
}

func (s ReportFlowIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowIspProvinceServiceResponseHeader struct {
}

func (s ReportFlowIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowIspProvinceServiceResponseHeader) GoString() string {
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
  DomainDir []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir `json:"domainDir,omitempty" xml:"domainDir,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDomainDir(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
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

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir struct     {
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

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) SetDomain(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir {
  s.Domain = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) SetDir(v []*string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir {
  s.Dir = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) SetResult(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse {
  s.Result = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) SetDomain(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) SetDetails(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult {
  s.Details = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails struct     {
  // {"en":"Directory", "zh_CN":"目录"}
  Dir *string `json:"dir,omitempty" xml:"dir,omitempty" require:"true"`
  // {"en":"Return when the dataType is flow; Total flow of every stream. Unit is MB and 2 digits of decimals allowed;", "zh_CN":"当dataType为flow时返回;每路流的总流量单位MB,保留2位小数;"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Return when the dataType is bandwidth;
  // 
  // Bandwidth peak value of every stream within specified time. Unit is Mbps, two decimals digits;", "zh_CN":"当dataType为bandwidth时返回;
  // 每路流在该时间段内的带宽峰值单位Mbps,保留2位小数;"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty" require:"true"`
  Details []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetDir(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.Dir = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetTotalFlow(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetBandwidthPeakValue(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetDetails(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.Details = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails struct     {
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

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) SetTimestamp(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) SetValue(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainQueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails {
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




type QueryRequesBandwidthtSavingRatioRequest struct {
  // {"en":"From date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8);
  // 2.Cannot exceed current time;
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2018年12月2日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"To time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8);
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time;
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception;
  // 4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days", "zh_CN":"结束时间:
  // 1.时间格式2019-01-02T10:00:00+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The maximum number of deliverable domain names is 200 by default;
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names);
  // 3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为200个
  // 2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Time interval of data: 5m (5 min), 1h (1 hour), 1d (1 day); 
  // 	The default is 1d.", "zh_CN":"数据粒度:
  // 1.支持5m(5分钟)、1h(1小时)、1d(天)
  // 2.不传默认1d。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryRequesBandwidthtSavingRatioRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioRequest) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDateFrom(v string) *QueryRequesBandwidthtSavingRatioRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDateTo(v string) *QueryRequesBandwidthtSavingRatioRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDomain(v []*string) *QueryRequesBandwidthtSavingRatioRequest {
  s.Domain = v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDataInterval(v string) *QueryRequesBandwidthtSavingRatioRequest {
  s.DataInterval = &v
  return s
}

type QueryRequesBandwidthtSavingRatioResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequesBandwidthtSavingRatioResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioResponse) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioResponse) SetCode(v string) *QueryRequesBandwidthtSavingRatioResponse {
  s.Code = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponse) SetMessage(v string) *QueryRequesBandwidthtSavingRatioResponse {
  s.Message = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponse) SetData(v []*QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData) *QueryRequesBandwidthtSavingRatioResponse {
  s.Data = v
  return s
}

type QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间,格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total saving of bandwidth.", "zh_CN":"总节省带宽的平均值"}
  TotalAvg *int `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  SavingBandwidthDatas []*QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas `json:"savingBandwidthDatas,omitempty" xml:"savingBandwidthDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData) SetRealDate(v string) *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData {
  s.RealDate = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData) SetTotalAvg(v int) *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData {
  s.TotalAvg = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData) SetSavingBandwidthDatas(v []*QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseData {
  s.SavingBandwidthDatas = v
  return s
}

type QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas struct     {
  // {"en":"timetamp
  // 1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range, For example yyyy-MM-dd 00:05 represents data in the range from 00:00 to 00:05.
  // 2.The data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents data values in the previous time granularity range such as yyyy-MM-dd 01 that represent data from 00 to 01.
  // 3. the data granularity of the query is daily, the format is yyyy-MM-dd; Each time slice data value represents the data value for that day.
  // 4.Returns the timetamp contained in start time and end time.", "zh_CN":"时间片
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 01,代表00到01之间的数据。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"saving of bandwidth,keep 4 decimal places", "zh_CN":"节省带宽,保留4位小数"}
  SavingBandwidth *string `json:"savingBandwidth,omitempty" xml:"savingBandwidth,omitempty" require:"true"`
}

func (s QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) SetTimestamp(v string) *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) SetSavingBandwidth(v string) *QueryRequesBandwidthtSavingRatioQueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas {
  s.SavingBandwidth = &v
  return s
}

type QueryRequesBandwidthtSavingRatioPaths struct {
}

func (s QueryRequesBandwidthtSavingRatioPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioPaths) GoString() string {
  return s.String()
}

type QueryRequesBandwidthtSavingRatioParameters struct {
}

func (s QueryRequesBandwidthtSavingRatioParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioParameters) GoString() string {
  return s.String()
}

type QueryRequesBandwidthtSavingRatioRequestHeader struct {
}

func (s QueryRequesBandwidthtSavingRatioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioRequestHeader) GoString() string {
  return s.String()
}

type QueryRequesBandwidthtSavingRatioResponseHeader struct {
}

func (s QueryRequesBandwidthtSavingRatioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioResponseHeader) GoString() string {
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
  DomainStream []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDomainStream(v []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream) *QueryStreamTrafficUnderMALBDomainRequest {
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

type QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream struct     {
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"可传递域名数量上限默认为20个(可联系技术支持调整)"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Stream name:
  // 				1. If this field is not specified, it means all domains under the domain are queried;
  // 				2. Number of streams can be adjusted depending on different accounts. The default value is 2000(this limit applies to the empty value);", "zh_CN":"流名：'发布点'+'流名'。例如：live/test-20180101-test ,其中live是发布点，test-20180101-test是流名。"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream) SetDomain(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream) SetStream(v []*string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainRequestDomainStream {
  s.Stream = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponse struct {
  Result []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponse) SetResult(v []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult) *QueryStreamTrafficUnderMALBDomainResponse {
  s.Result = v
  return s
}

type QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult) SetDomain(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult) SetDetails(v []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResult {
  s.Details = v
  return s
}

type QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails struct     {
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
  Details []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) SetStream(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.Stream = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) SetTotalFlow(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) SetBandwidthPeakValue(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails) SetDetails(v []*QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.Details = v
  return s
}

type QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails struct     {
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

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) SetTimestamp(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) SetValue(v string) *QueryStreamTrafficUnderMALBDomainQueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails {
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




type FlowDayRequest struct {
  // {"en":"main account cust_en_name, cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"主账号客户英文名，合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
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
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions. Common regions: cn: Mainland China, hk: Hong Kong, ov: Oversea average, tw: Taiwan, euna: Europe and North America, apac: Asia-Pacific, sa: South America, af: Africa, am: Americas, emea: Europe/Middle East/Africa, kr: South Korea, au: Australia, in: India, jp: Japan, ru: Russia, indo: Indonesia, me: Middle East, eu: Europe, ph: Philippines", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。常用区域：cn:中国大陆,hk:香港,ov:海外平均,tw:台湾,euna:欧美,apac:亚太,sa:南美,af:非洲,am:美洲,emea:欧洲/中东/非洲,kr:韩国,au:澳大利亚,in:印度,jp:日本,ru:俄罗斯,indo:印尼,me:中东,eu:欧洲,ph:菲律宾 "}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"1)If there isp multiple inputs,use ';' as demimeter.
  // 2)optional values of isp: refers to the ISP-section of appendix.
  // 3) If not specified,means all the isp.", "zh_CN":"要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes. Common Acceleration Types: Web Acceleration: web, Web-HTTPS: web-https, Whole Site Acceleration: wsa, Whole Site Acceleration-HTTPS: wsa-https, HTTP Download: download, Livestream Media: livestream, Livestream-HTTPS: live-https, VOD-HTTPS: vod-https, Cloud Video on Demand: cloudv-vod, VOD Streaming: vodstream, Financial Security Acceleration Solution: fsa, Government and Enterprise Security Acceleration Solution: gess, E-commerce Security Acceleration Solution: esa, Mobile Acceleration (MAA): maa, Application Security Acceleration Solution: s-appa, Application Acceleration: appa, WAF: waf, Upload Acceleration: upload, Upload Acceleration-HTTPS: upa-https, IPv6 Integration Solution: osv6, Live P2P: livep2p", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型. 常用加速类型：网页加速:web,网页-HTTPS:web-https ,全站加速:wsa,全站加速-https:wsa-https,HTTP下载:download,流媒体直播:livestream,直播-https:live-https,点播-HTTPS:vod-https,云点播:cloudv-vod,流媒体点播:vodstream,金融安全加速解决方案:fsa,政企安全加速解决方案:gess,电商安全加速解决方案:esa,移动加速(MAA):maa,应用安全加速解决方案:s-appa,应用加速:appa,WAF:waf,上传加速:upload,上传加速-https:upa-https,IPv6一体化解决方案:osv6,直播P2P:livep2p"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Greenwich time zone, the parameter format GMT+09:00 means East Nine District, GMT-09:00 means West Nine District, if not passed, the default is the local time zone (East Eight District)", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
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

type FlowDayResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *FlowDayFlowDayResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s FlowDayResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowDayResponse) GoString() string {
  return s.String()
}

func (s *FlowDayResponse) SetProvider(v *FlowDayFlowDayResponseProvider) *FlowDayResponse {
  s.Provider = v
  return s
}

type FlowDayFlowDayResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'数据'}
  Date *FlowDayFlowDayResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s FlowDayFlowDayResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s FlowDayFlowDayResponseProvider) GoString() string {
  return s.String()
}

func (s *FlowDayFlowDayResponseProvider) SetName(v string) *FlowDayFlowDayResponseProvider {
  s.Name = &v
  return s
}

func (s *FlowDayFlowDayResponseProvider) SetType(v string) *FlowDayFlowDayResponseProvider {
  s.Type = &v
  return s
}

func (s *FlowDayFlowDayResponseProvider) SetDate(v *FlowDayFlowDayResponseProviderDate) *FlowDayFlowDayResponseProvider {
  s.Date = v
  return s
}

type FlowDayFlowDayResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始日期'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束日期'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *FlowDayFlowDayResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s FlowDayFlowDayResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s FlowDayFlowDayResponseProviderDate) GoString() string {
  return s.String()
}

func (s *FlowDayFlowDayResponseProviderDate) SetStartdate(v string) *FlowDayFlowDayResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *FlowDayFlowDayResponseProviderDate) SetEnddate(v string) *FlowDayFlowDayResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *FlowDayFlowDayResponseProviderDate) SetChannel(v *FlowDayFlowDayResponseProviderDateChannel) *FlowDayFlowDayResponseProviderDate {
  s.Channel = v
  return s
}

type FlowDayFlowDayResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'channel-peak', 'zh_CN':'频道峰值数据'}
  ChannelPeak []*FlowDayFlowDayResponseProviderDateChannelChannelPeak `json:"channel-peak,omitempty" xml:"channel-peak,omitempty" require:"true" type:"Repeated"`
}

func (s FlowDayFlowDayResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s FlowDayFlowDayResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *FlowDayFlowDayResponseProviderDateChannel) SetName(v string) *FlowDayFlowDayResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *FlowDayFlowDayResponseProviderDateChannel) SetChannelPeak(v []*FlowDayFlowDayResponseProviderDateChannelChannelPeak) *FlowDayFlowDayResponseProviderDateChannel {
  s.ChannelPeak = v
  return s
}

type FlowDayFlowDayResponseProviderDateChannelChannelPeak struct     {
  // {'en':'date', 'zh_CN':'日期'}
  Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
  // {'en':'peakTime', 'zh_CN':'峰值时间'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'peakvalue(Mbps)', 'zh_CN':'带宽峰值（单位Mbps）'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'the total flow', 'zh_CN':'总流量'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
}

func (s FlowDayFlowDayResponseProviderDateChannelChannelPeak) String() string {
  return tea.Prettify(s)
}

func (s FlowDayFlowDayResponseProviderDateChannelChannelPeak) GoString() string {
  return s.String()
}

func (s *FlowDayFlowDayResponseProviderDateChannelChannelPeak) SetDate(v string) *FlowDayFlowDayResponseProviderDateChannelChannelPeak {
  s.Date = &v
  return s
}

func (s *FlowDayFlowDayResponseProviderDateChannelChannelPeak) SetPeakTime(v string) *FlowDayFlowDayResponseProviderDateChannelChannelPeak {
  s.PeakTime = &v
  return s
}

func (s *FlowDayFlowDayResponseProviderDateChannelChannelPeak) SetPeakValue(v string) *FlowDayFlowDayResponseProviderDateChannelChannelPeak {
  s.PeakValue = &v
  return s
}

func (s *FlowDayFlowDayResponseProviderDateChannelChannelPeak) SetTotalFlow(v string) *FlowDayFlowDayResponseProviderDateChannelChannelPeak {
  s.TotalFlow = &v
  return s
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

type FlowDayRequestHeader struct {
}

func (s FlowDayRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowDayRequestHeader) GoString() string {
  return s.String()
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
  Data []*ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportFlowP2pShareRatioServiceResponse) SetData(v []*ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData) *ReportFlowP2pShareRatioServiceResponse {
  s.Data = v
  return s
}

type ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData struct     {
  // {"en":"Domain. If merge date of all domains will not return domain", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData) SetDomain(v string) *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData) SetDetailList(v []*ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList) *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseData {
  s.DetailList = v
  return s
}

type ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList struct     {
  // {"en":"timestamp,Returns the timestamp between the start time and end time. Time format:  yyyy-MM-dd HH:mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。
  // 时间格式:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"share rate, unit %", "zh_CN":"分享率,单位%"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList) SetTimestamp(v string) *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList) SetValue(v string) *ReportFlowP2pShareRatioServiceReportFlowP2pShareRatioServiceResponseDataDetailList {
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
  Data []*QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryDataTransferForAllDomainByteResponse) SetData(v []*QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData) *QueryDataTransferForAllDomainByteResponse {
  s.Data = v
  return s
}

type QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData struct     {
  // {"en":"flowSummary", "zh_CN":"总流量"}
  FlowSummary *string `json:"flowSummary,omitempty" xml:"flowSummary,omitempty" require:"true"`
  FlowData []*QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData) SetFlowSummary(v string) *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData {
  s.FlowSummary = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData) SetFlowData(v []*QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData) *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseData {
  s.FlowData = v
  return s
}

type QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData) SetTimestamp(v string) *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData) SetFlow(v string) *QueryDataTransferForAllDomainByteQueryDataTransferForAllDomainByteResponseDataFlowData {
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




type QueryTrafficBySpecificProtocolRequest struct {
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
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // Support for 1m(1 minutes), 5m (5 minutes), 1h (1 hour), 1d (1 day)", "zh_CN":"数据粒度:
  // 支持1m(1分钟),5m(5分钟),1h(1小时),1d(1天)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Transmission protocol:
  // 1.Options: http, https;
  // 2.https is used as the default value is no value specified;
  // 3.httpFlowData is displayed if http is queried, and httpsFlowData is displayed if https is queried;", "zh_CN":"传输协议
  // 1.可选值为http、https;
  // 2.不传默认查询https;
  // 3.查询http时出参展示httpFlowData,查询https时出参展示httpsFlowData;"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group dimension:
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryTrafficBySpecificProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolRequest) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDateFrom(v string) *QueryTrafficBySpecificProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDateTo(v string) *QueryTrafficBySpecificProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDomain(v []*string) *QueryTrafficBySpecificProtocolRequest {
  s.Domain = v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDataInterval(v string) *QueryTrafficBySpecificProtocolRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetProtocolType(v string) *QueryTrafficBySpecificProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetGroupBy(v []*string) *QueryTrafficBySpecificProtocolRequest {
  s.GroupBy = v
  return s
}

type QueryTrafficBySpecificProtocolResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficBySpecificProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolResponse) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolResponse) SetResult(v []*QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult) *QueryTrafficBySpecificProtocolResponse {
  s.Result = v
  return s
}

type QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  HttpsFlowData []*QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData `json:"httpsFlowData,omitempty" xml:"httpsFlowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult) SetDomain(v string) *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult) SetHttpsFlowData(v []*QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData) *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResult {
  s.HttpsFlowData = v
  return s
}

type QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData struct     {
  // {"en":"DateTime: the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits  of decimals allowed", "zh_CN":"流量值,单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData) SetTimestamp(v string) *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData) SetValue(v string) *QueryTrafficBySpecificProtocolQueryTrafficBySpecificProtocolResponseResultHttpsFlowData {
  s.Value = &v
  return s
}

type QueryTrafficBySpecificProtocolPaths struct {
}

func (s QueryTrafficBySpecificProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolPaths) GoString() string {
  return s.String()
}

type QueryTrafficBySpecificProtocolParameters struct {
}

func (s QueryTrafficBySpecificProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolParameters) GoString() string {
  return s.String()
}

type QueryTrafficBySpecificProtocolRequestHeader struct {
}

func (s QueryTrafficBySpecificProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolRequestHeader) GoString() string {
  return s.String()
}

type QueryTrafficBySpecificProtocolResponseHeader struct {
}

func (s QueryTrafficBySpecificProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolResponseHeader) GoString() string {
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
  DomainStream []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDomainStream(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
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

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream struct     {
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

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) SetDomain(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) SetStream(v []*string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream {
  s.Stream = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetData(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Data = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DomainOfStreamList []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList `json:"domainOfStreamList,omitempty" xml:"domainOfStreamList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) SetDomain(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData {
  s.Domain = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) SetDomainOfStreamList(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData {
  s.DomainOfStreamList = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList struct     {
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
  DetailList []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetStream(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.Stream = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetTotalFlow(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.TotalFlow = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetBandwidthPeakValue(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetDetailList(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.DetailList = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList struct     {
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

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetTimestamp(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetValue(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.Value = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetOnlineUser(v int) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainQueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
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




type QueryISPProvinceHitRateRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年01月01日10点0分0秒）;
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如 2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔:24小时（可联系技术支持调整），即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3. Domain name exceeding limit, misstatement", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）;
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 3.域名超过上限，提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity: 
  // 1. default 1m
  // 2. 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度:
  // 1.不传默认1m
  // 2.支持1m（1分钟）、5m（5分钟）"}
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
  // 1. Optional value flow, request
  // 2. Default flow
  // 3. Flow: Flow, keep two decimal places;
  // 4. Request: number of Request", "zh_CN":"查询维度:
  // 1.可选值 flow、request
  // 2.传默认 flow
  // 3.flow:流量，保留两位小数;
  // 4.request:请求数"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional: domain, all, If it is empty, it defaults to returning by domain dimension;
  // If all is passed, merge and return according to the query domain name.", "zh_CN":"可选项：domain、all, 为空则默认为按domain维度返回;
  // 若传递all，则按查询域名合并返回"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryISPProvinceHitRateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateRequest) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateRequest) SetDateFrom(v string) *QueryISPProvinceHitRateRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetDateTo(v string) *QueryISPProvinceHitRateRequest {
  s.DateTo = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetDomain(v []*string) *QueryISPProvinceHitRateRequest {
  s.Domain = v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetDataInterval(v string) *QueryISPProvinceHitRateRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetProvince(v []*string) *QueryISPProvinceHitRateRequest {
  s.Province = v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetIsp(v []*string) *QueryISPProvinceHitRateRequest {
  s.Isp = v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetQueryBy(v string) *QueryISPProvinceHitRateRequest {
  s.QueryBy = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetGroupBy(v string) *QueryISPProvinceHitRateRequest {
  s.GroupBy = &v
  return s
}

type QueryISPProvinceHitRateResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceHitRateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateResponse) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateResponse) SetCode(v string) *QueryISPProvinceHitRateResponse {
  s.Code = &v
  return s
}

func (s *QueryISPProvinceHitRateResponse) SetMessage(v string) *QueryISPProvinceHitRateResponse {
  s.Message = &v
  return s
}

func (s *QueryISPProvinceHitRateResponse) SetData(v []*QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData) *QueryISPProvinceHitRateResponse {
  s.Data = v
  return s
}

type QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData) SetDomain(v string) *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData {
  s.Domain = &v
  return s
}

func (s *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData) SetDetailList(v []*QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList) *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseData {
  s.DetailList = v
  return s
}

type QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Hit data:
  // Flow: Flow, keep two decimal places;
  // Request: number of Request", "zh_CN":"命中数据:
  // 1.flow:流量，保留两位小数;
  // 2.request:请求数"}
  HitValue *string `json:"hitValue,omitempty" xml:"hitValue,omitempty" require:"true"`
  // {"en":"Hit rate, keep four decimal places", "zh_CN":"命中率，保留四位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
}

func (s QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList) SetTimestamp(v string) *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList) SetHitValue(v string) *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList {
  s.HitValue = &v
  return s
}

func (s *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList) SetHitRate(v string) *QueryISPProvinceHitRateQueryISPProvinceHitRateResponseDataDetailList {
  s.HitRate = &v
  return s
}

type QueryISPProvinceHitRatePaths struct {
}

func (s QueryISPProvinceHitRatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRatePaths) GoString() string {
  return s.String()
}

type QueryISPProvinceHitRateParameters struct {
}

func (s QueryISPProvinceHitRateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateParameters) GoString() string {
  return s.String()
}

type QueryISPProvinceHitRateRequestHeader struct {
}

func (s QueryISPProvinceHitRateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateRequestHeader) GoString() string {
  return s.String()
}

type QueryISPProvinceHitRateResponseHeader struct {
}

func (s QueryISPProvinceHitRateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateResponseHeader) GoString() string {
  return s.String()
}




