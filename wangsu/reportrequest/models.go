package reportrequest

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportRequestQuicServiceRequest struct {
  // {"en":"Start time:
  // 1. The time format is yyyy-MM-dd, for example, 2021-10-10;
  // 2. It cannot be greater than the current time
  // 3. Data for the most recent six months (183 days) can be obtained.", "zh_CN":"开始时间：
  //             时间格式为yyyy-MM-dd，例如，2021-10-10；
  //             不能大于当前时间
  //             最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-dd
  // 2. The end time must be greater than the start time. If the end time is greater than the current time, the current time is used.
  // 3. If both dateFrom and dateTo are not passed, the default query is the past day; if only one is not passed, an exception is thrown
  // 4. The maximum query time interval allowed is 7 days, that is, the difference between dateFrom and dateTo cannot exceed 7 days (you can contact technical support to adjust)", "zh_CN":"结束时间：
  //             时间格式为yyyy-MM-dd
  //             结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  //             dateFrom，dateTo二者都未传，默认查询过去的1天；如仅有一个未传，抛异常
  //             允许查询最大时间间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The upper limit of the number of domain names that can be passed is 20 by default (you can contact technical support to adjust it).
  // 2. Automatically filter out illegal domain names (if an illegal domain name is passed, it will be filtered out, and the query result will only return data for legal domain names).
  // 3. When this parameter is not passed, all domain names under the account are queried by default, but an error message will be prompted when the number of domain names under the account exceeds the upper limit.", "zh_CN":"域名：
  //             可传递域名数量上限默认为20个（可联系技术支持调整）。
  //             自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  //             未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s ReportRequestQuicServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportRequestQuicServiceRequest) SetDateFrom(v string) *ReportRequestQuicServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportRequestQuicServiceRequest) SetDateTo(v string) *ReportRequestQuicServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportRequestQuicServiceRequest) SetDomain(v []*string) *ReportRequestQuicServiceRequest {
  s.Domain = v
  return s
}

type ReportRequestQuicServiceResponse struct {
  // {"en":"code", "zh_CN":"code"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"message"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportRequestQuicServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestQuicServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportRequestQuicServiceResponse) SetCode(v string) *ReportRequestQuicServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportRequestQuicServiceResponse) SetMessage(v string) *ReportRequestQuicServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportRequestQuicServiceResponse) SetData(v []*ReportRequestQuicServiceResponseData) *ReportRequestQuicServiceResponse {
  s.Data = v
  return s
}

type ReportRequestQuicServiceResponseData struct     {
  // {"en":"Time, in the format of yyyy-MM-dd HH:mm:ss", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm:ss"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests, in units", "zh_CN":"请求个数，单位 个"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Number of requests, in units", "zh_CN":"请求个数，单位 个"}
  QuicRequest *string `json:"quicRequest,omitempty" xml:"quicRequest,omitempty" require:"true"`
}

func (s ReportRequestQuicServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportRequestQuicServiceResponseData) SetTimestamp(v string) *ReportRequestQuicServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportRequestQuicServiceResponseData) SetTotalRequest(v string) *ReportRequestQuicServiceResponseData {
  s.TotalRequest = &v
  return s
}

func (s *ReportRequestQuicServiceResponseData) SetQuicRequest(v string) *ReportRequestQuicServiceResponseData {
  s.QuicRequest = &v
  return s
}

type ReportRequestQuicServicePaths struct {
}

func (s ReportRequestQuicServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServicePaths) GoString() string {
  return s.String()
}

type ReportRequestQuicServiceParameters struct {
}

func (s ReportRequestQuicServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceParameters) GoString() string {
  return s.String()
}

type ReportRequestQuicServiceRequestHeader struct {
}

func (s ReportRequestQuicServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportRequestQuicServiceResponseHeader struct {
}

func (s ReportRequestQuicServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportHitServiceRequest struct {
}

func (s ReportHitServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServiceRequest) GoString() string {
  return s.String()
}

type ReportHitServiceResponse struct {
  // {"en":"Total requests", "zh_CN":"总请求数"}
  HitSummary *string `json:"hit-summary,omitempty" xml:"hit-summary,omitempty" require:"true"`
  // {'en':'hitData', 'zh_CN':'请求数数据'}
  HitData []*ReportHitServiceResponseHitData `json:"hit-data,omitempty" xml:"hit-data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportHitServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportHitServiceResponse) SetHitSummary(v string) *ReportHitServiceResponse {
  s.HitSummary = &v
  return s
}

func (s *ReportHitServiceResponse) SetHitData(v []*ReportHitServiceResponseHitData) *ReportHitServiceResponse {
  s.HitData = v
  return s
}

type ReportHitServiceResponseHitData struct     {
  // {"en":"Date
  // When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00;When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24;When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data;Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值；
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests", "zh_CN":"请求数"}
  Hit *string `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
}

func (s ReportHitServiceResponseHitData) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServiceResponseHitData) GoString() string {
  return s.String()
}

func (s *ReportHitServiceResponseHitData) SetTimestamp(v string) *ReportHitServiceResponseHitData {
  s.Timestamp = &v
  return s
}

func (s *ReportHitServiceResponseHitData) SetHit(v string) *ReportHitServiceResponseHitData {
  s.Hit = &v
  return s
}

type ReportHitServicePaths struct {
}

func (s ReportHitServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServicePaths) GoString() string {
  return s.String()
}

type ReportHitServiceParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00
  // ;2.And smaller than the current time and ‘dateTo’;
  // 3.Period between ‘dataFrom’ and ‘dateTo’ cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过31天（可联系技术支持调整）；
  // 4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than ‘dateFrom’;
  // 3.If it’s greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；
  // 3.如果大于当前时间，则重新赋值为当前时间；"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer", "zh_CN":"数据粒度
  // 1.fiveminutes：5分钟，hourly：1小时，daily：1天；
  // 2.不传递，默认为daily；
  // 3.传递fiveminutes时，若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ReportHitServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServiceParameters) GoString() string {
  return s.String()
}

func (s *ReportHitServiceParameters) SetDatefrom(v string) *ReportHitServiceParameters {
  s.Datefrom = &v
  return s
}

func (s *ReportHitServiceParameters) SetDateto(v string) *ReportHitServiceParameters {
  s.Dateto = &v
  return s
}

func (s *ReportHitServiceParameters) SetType(v string) *ReportHitServiceParameters {
  s.Type = &v
  return s
}

type ReportHitServiceRequestHeader struct {
}

func (s ReportHitServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportHitServiceResponseHeader struct {
}

func (s ReportHitServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportHitServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUserRequestCountryServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time;
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. The default query interval is 7 days, Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days.", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.默认查询间隔7天,允许查询最大间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1. Domain is not uploaded: Query all domain names of the account (More than 20 domains will error,you can contact technical support for adjustment);
  // 2. Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration area:
  // 1. Acceleration areaCode is not uploaded: Query all acceleration areas by default.", "zh_CN":"加速区域:
  // 未传递areaCode时,默认查询所有加速区域。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 5m: 5 minute granularity;
  // 1h: 1 hour granularity;
  // 1d: 1 day granularity; Default value is 1d.", "zh_CN":"数据粒度:
  // 5m:5分钟粒度;
  // 1h:1小时粒度;
  // 1d:1天粒度。不传默认1天粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Country area:
  // 1. countryCode is not uploaded: Query all country areas by default;
  // 2. countryCode is uploaded: Multiple can be uploaded, such as cn, in.  Please refer to the appendix description section of the overview page.", "zh_CN":"国家区域(含中国台湾、中国澳门、中国香港、中国大陆):
  // 1.未传递countryCode时:查询全部国家区域;
  // 2.有传递countryCode时:可传多个,如cn,in。可传递的值详见概览页附录说明章节"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Grouped dimension:
  // 1. The optional values are domain, countryCode;Multiple values can be uploaded;
  // 2. If no value is uploaded: Aggregate all data by default.", "zh_CN":"分组维度
  // 可选值为domain、countryCode,可传入多个值;
  // 有传入则按照该维度展示明细数据;
  // 没传默认全部聚合。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceRequest) SetDateFrom(v string) *ReportUserRequestCountryServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetDateTo(v string) *ReportUserRequestCountryServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetDomain(v []*string) *ReportUserRequestCountryServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetAreaCode(v []*string) *ReportUserRequestCountryServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetDataInterval(v string) *ReportUserRequestCountryServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetCountryCode(v []*string) *ReportUserRequestCountryServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetGroupBy(v []*string) *ReportUserRequestCountryServiceRequest {
  s.GroupBy = v
  return s
}

type ReportUserRequestCountryServiceResponse struct {
  Result []*ReportUserRequestCountryServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponse) SetResult(v []*ReportUserRequestCountryServiceResponseResult) *ReportUserRequestCountryServiceResponse {
  s.Result = v
  return s
}

type ReportUserRequestCountryServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportUserRequestCountryServiceResponseResultCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponseResult) SetDomain(v string) *ReportUserRequestCountryServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResult) SetCountryData(v []*ReportUserRequestCountryServiceResponseResultCountryData) *ReportUserRequestCountryServiceResponseResult {
  s.CountryData = v
  return s
}

type ReportUserRequestCountryServiceResponseResultCountryData struct     {
  // {"en":"Country area", "zh_CN":"国家区域"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Total number of requests.", "zh_CN":"国家的请求总数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  RequestData []*ReportUserRequestCountryServiceResponseResultCountryDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceResponseResultCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseResultCountryData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponseResultCountryData) SetCountryCode(v string) *ReportUserRequestCountryServiceResponseResultCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResultCountryData) SetTotalRequest(v string) *ReportUserRequestCountryServiceResponseResultCountryData {
  s.TotalRequest = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResultCountryData) SetRequestData(v []*ReportUserRequestCountryServiceResponseResultCountryDataRequestData) *ReportUserRequestCountryServiceResponseResultCountryData {
  s.RequestData = v
  return s
}

type ReportUserRequestCountryServiceResponseResultCountryDataRequestData struct     {
  // {"en":"Time:
  // 
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; ach time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00;
  // 3. When the data query granularity is 1d, the format is yyyy-MM-dd; Each time slice value represents the value of the day;
  // 4. Return the time slices that contained in start time and in end time.", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 查询的数据粒度为1d时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests.", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s ReportUserRequestCountryServiceResponseResultCountryDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseResultCountryDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponseResultCountryDataRequestData) SetTimestamp(v string) *ReportUserRequestCountryServiceResponseResultCountryDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResultCountryDataRequestData) SetRequest(v string) *ReportUserRequestCountryServiceResponseResultCountryDataRequestData {
  s.Request = &v
  return s
}

type ReportUserRequestCountryServicePaths struct {
}

func (s ReportUserRequestCountryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServicePaths) GoString() string {
  return s.String()
}

type ReportUserRequestCountryServiceParameters struct {
}

func (s ReportUserRequestCountryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceParameters) GoString() string {
  return s.String()
}

type ReportUserRequestCountryServiceRequestHeader struct {
}

func (s ReportUserRequestCountryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUserRequestCountryServiceResponseHeader struct {
}

func (s ReportUserRequestCountryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportRequestHttpHttpsServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH:mm:SS +08:00, for example, 2021-05-19T10:00:00+08:00 (10:00:00 Beijing time on May 19, 2021);
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
  // 4. Maximum query interval allowed: 1 day, that is, the difference between dateFrom and dateTo can not exceed 1 day.", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support).
  // 	2.Domain is not uploaded: Query all domain names of the account", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"1. Optional value: domain
  // 	2. No value is passed, and all domain name data is aggregated by default", "zh_CN":"1.可选值:domain
  // 2.不传默认聚合所有频道数据"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportRequestHttpHttpsServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportRequestHttpHttpsServiceRequest) SetDateFrom(v string) *ReportRequestHttpHttpsServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceRequest) SetDateTo(v string) *ReportRequestHttpHttpsServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceRequest) SetDomain(v []*string) *ReportRequestHttpHttpsServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportRequestHttpHttpsServiceRequest) SetGroupBy(v []*string) *ReportRequestHttpHttpsServiceRequest {
  s.GroupBy = v
  return s
}

type ReportRequestHttpHttpsServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportRequestHttpHttpsServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestHttpHttpsServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportRequestHttpHttpsServiceResponse) SetCode(v string) *ReportRequestHttpHttpsServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceResponse) SetMessage(v string) *ReportRequestHttpHttpsServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceResponse) SetData(v []*ReportRequestHttpHttpsServiceResponseData) *ReportRequestHttpHttpsServiceResponse {
  s.Data = v
  return s
}

type ReportRequestHttpHttpsServiceResponseData struct     {
  // {"en":"Domain name, this field is not returned when summarizing all domain name data", "zh_CN":"域名,聚合全部域名数据不返回该字段"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*ReportRequestHttpHttpsServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestHttpHttpsServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportRequestHttpHttpsServiceResponseData) SetDomain(v string) *ReportRequestHttpHttpsServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceResponseData) SetDetailList(v []*ReportRequestHttpHttpsServiceResponseDataDetailList) *ReportRequestHttpHttpsServiceResponseData {
  s.DetailList = v
  return s
}

type ReportRequestHttpHttpsServiceResponseDataDetailList struct     {
  // {"en":"TimeStamp, returns the time slice containing the start time and end time. Time format: 1 minute: yyyy-MM-dd HH: mm", "zh_CN":"时间片,返回开始时间和结束时间包含的时间片。时间格式:1分钟:yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"HTTPS requests", "zh_CN":"https请求数"}
  HttpsValue *string `json:"httpsValue,omitempty" xml:"httpsValue,omitempty" require:"true"`
  // {"en":"HTTP requests", "zh_CN":"http请求数"}
  HttpValue *string `json:"httpValue,omitempty" xml:"httpValue,omitempty" require:"true"`
}

func (s ReportRequestHttpHttpsServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportRequestHttpHttpsServiceResponseDataDetailList) SetTimestamp(v string) *ReportRequestHttpHttpsServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceResponseDataDetailList) SetHttpsValue(v string) *ReportRequestHttpHttpsServiceResponseDataDetailList {
  s.HttpsValue = &v
  return s
}

func (s *ReportRequestHttpHttpsServiceResponseDataDetailList) SetHttpValue(v string) *ReportRequestHttpHttpsServiceResponseDataDetailList {
  s.HttpValue = &v
  return s
}

type ReportRequestHttpHttpsServicePaths struct {
}

func (s ReportRequestHttpHttpsServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServicePaths) GoString() string {
  return s.String()
}

type ReportRequestHttpHttpsServiceParameters struct {
}

func (s ReportRequestHttpHttpsServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceParameters) GoString() string {
  return s.String()
}

type ReportRequestHttpHttpsServiceRequestHeader struct {
}

func (s ReportRequestHttpHttpsServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportRequestHttpHttpsServiceResponseHeader struct {
}

func (s ReportRequestHttpHttpsServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHttpHttpsServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUserRequestIspProvinceServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. The default query interval upper limit is 7 days (can be adjusted by contacting technical support), the maximum allowable query interval: 31 days, that is, the difference between dateFrom and dateTo cannot exceed 31 days.", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.默认查询间隔上限为7天(可联系技术支持调整),允许查询最大间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1.Domain is not uploaded: Query all domain names of the account(More than 20 domains will error,you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration area:
  // Acceleration areaCode is not uploaded: Query all acceleration areas by default.", "zh_CN":"加速区域:
  // 未传递areaCode时,默认查询所有加速区域。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {"en":"By default, all provinces are queried, and the Chinese name of the province is passed. The province information code table is detailed in the appendix of the overview page.", "zh_CN":"默认查询全部省份,传递省份中文名称,省份信息码表详见概览页附录说明章节"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"By default, all operators are queried, and the operator's Chinese name is transmitted. The operator information code table is detailed in the appendix of the overview page.", "zh_CN":"默认查询全部运营商,传递运营商中文名称,运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Grouped dimension:
  // 1. The optional values are domain, province, isp; Multiple values can be uploaded;
  // 2. If no value is uploaded: Aggregate all data by default.", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值;
  // 2.有传入则按照该维度展示明细数据,,没传默认全部聚合。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportUserRequestIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetDateFrom(v string) *ReportUserRequestIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetDateTo(v string) *ReportUserRequestIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetDomain(v []*string) *ReportUserRequestIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetAreaCode(v []*string) *ReportUserRequestIspProvinceServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetProvince(v []*string) *ReportUserRequestIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetIsp(v []*string) *ReportUserRequestIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportUserRequestIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportUserRequestIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportUserRequestIspProvinceServiceResponse struct {
  Result []*ReportUserRequestIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUserRequestIspProvinceServiceResponse) SetResult(v []*ReportUserRequestIspProvinceServiceResponseResult) *ReportUserRequestIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportUserRequestIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportUserRequestIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportUserRequestIspProvinceServiceResponseResult) SetDomain(v string) *ReportUserRequestIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceResponseResult) SetIspData(v []*ReportUserRequestIspProvinceServiceResponseResultIspData) *ReportUserRequestIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportUserRequestIspProvinceServiceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportUserRequestIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData) *ReportUserRequestIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Total number of requests.", "zh_CN":"省份运营商的请求总数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  RequestData []*ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData) SetTotalRequest(v string) *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData {
  s.TotalRequest = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData) SetRequestData(v []*ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceData {
  s.RequestData = v
  return s
}

type ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData struct     {
  // {"en":"Time:
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; ach time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00;
  // 3. When the data query granularity is 1d, the format is yyyy-MM-dd; Each time slice value represents the value of the day;
  // 4. Return the time slices that contained in start time and in end time.", "zh_CN":"时间:
  // 1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 3.查询的数据粒度为1d时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests.", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) SetTimestamp(v string) *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) SetRequest(v string) *ReportUserRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData {
  s.Request = &v
  return s
}

type ReportUserRequestIspProvinceServicePaths struct {
}

func (s ReportUserRequestIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportUserRequestIspProvinceServiceParameters struct {
}

func (s ReportUserRequestIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportUserRequestIspProvinceServiceRequestHeader struct {
}

func (s ReportUserRequestIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUserRequestIspProvinceServiceResponseHeader struct {
}

func (s ReportUserRequestIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportProtocolOriginRequestServiceRequest struct {
  // {'en':'Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.', 'zh_CN':'开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (you can contact technical support for adjustment).. ', 'zh_CN':'结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.默认允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).', 'zh_CN':'域名：
  // 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)；
  // 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'Data granularity:
  // 
  // 1m: 1 minute granularity; 
  // 5m: 5 minute granularity; Default value is 5m.
  // 1h: 1 hour granularity;
  // 1d: 1 day granularity.', 'zh_CN':'数据粒度：
  // 1m：1分钟粒度。
  // 5m：5分钟粒度。不传默认5分钟粒度
  // 1h：1小时粒度;
  // 1d：1天粒度'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Grouped dimension:
  // 1.When the "groupBy" parameter has a value, the value can only be "domain". It will group and return the detailed based on domain;
  // 2.When groupBy does not have a value, aggregate all domain by default.', 'zh_CN':'分组维度
  // 1.有传，只能传domain。按照domain分组展示明细数据;
  // 2.不传，则默认所有域名聚合。'}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s ReportProtocolOriginRequestServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportProtocolOriginRequestServiceRequest) SetDateFrom(v string) *ReportProtocolOriginRequestServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceRequest) SetDateTo(v string) *ReportProtocolOriginRequestServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceRequest) SetDomain(v []*string) *ReportProtocolOriginRequestServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportProtocolOriginRequestServiceRequest) SetDataInterval(v string) *ReportProtocolOriginRequestServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceRequest) SetGroupBy(v string) *ReportProtocolOriginRequestServiceRequest {
  s.GroupBy = &v
  return s
}

type ReportProtocolOriginRequestServiceResponse struct {
  // {'en':'data', 'zh_CN':'请求结果'}
  Data []*ReportProtocolOriginRequestServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportProtocolOriginRequestServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportProtocolOriginRequestServiceResponse) SetData(v []*ReportProtocolOriginRequestServiceResponseData) *ReportProtocolOriginRequestServiceResponse {
  s.Data = v
  return s
}

type ReportProtocolOriginRequestServiceResponseData struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'details', 'zh_CN':'请求结果的详细数据'}
  Details []*ReportProtocolOriginRequestServiceResponseDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportProtocolOriginRequestServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportProtocolOriginRequestServiceResponseData) SetDomain(v string) *ReportProtocolOriginRequestServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceResponseData) SetDetails(v []*ReportProtocolOriginRequestServiceResponseDataDetails) *ReportProtocolOriginRequestServiceResponseData {
  s.Details = v
  return s
}

type ReportProtocolOriginRequestServiceResponseDataDetails struct     {
  // {'en':'Time:
  // 
  // 1.When the data query granularity is 1m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00.
  // 2.When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00.
  // 3. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00.
  // 4. When the data query granularity is 1d, the format is yyyy-MM-dd; Each time slice value represents the value within the previous time granularity range.
  // 5. Return the time slices that contained in start time and in end time.', 'zh_CN':'时间，
  // 查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  // 查询的数据粒度为1d时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值；
  // 返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Number of all back-to-source requests.', 'zh_CN':'全部回源请求数'}
  AllRequests *string `json:"allRequests,omitempty" xml:"allRequests,omitempty" require:"true"`
  // {'en':'Number of http back-to-source requests.', 'zh_CN':'http回源请求数'}
  HttpRequests *string `json:"httpRequests,omitempty" xml:"httpRequests,omitempty" require:"true"`
  // {'en':'Number of https back-to-source requests.', 'zh_CN':'https回源请求数'}
  HttpsRequests *string `json:"httpsRequests,omitempty" xml:"httpsRequests,omitempty" require:"true"`
}

func (s ReportProtocolOriginRequestServiceResponseDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceResponseDataDetails) GoString() string {
  return s.String()
}

func (s *ReportProtocolOriginRequestServiceResponseDataDetails) SetTimestamp(v string) *ReportProtocolOriginRequestServiceResponseDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceResponseDataDetails) SetAllRequests(v string) *ReportProtocolOriginRequestServiceResponseDataDetails {
  s.AllRequests = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceResponseDataDetails) SetHttpRequests(v string) *ReportProtocolOriginRequestServiceResponseDataDetails {
  s.HttpRequests = &v
  return s
}

func (s *ReportProtocolOriginRequestServiceResponseDataDetails) SetHttpsRequests(v string) *ReportProtocolOriginRequestServiceResponseDataDetails {
  s.HttpsRequests = &v
  return s
}

type ReportProtocolOriginRequestServicePaths struct {
}

func (s ReportProtocolOriginRequestServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServicePaths) GoString() string {
  return s.String()
}

type ReportProtocolOriginRequestServiceParameters struct {
}

func (s ReportProtocolOriginRequestServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceParameters) GoString() string {
  return s.String()
}

type ReportProtocolOriginRequestServiceRequestHeader struct {
}

func (s ReportProtocolOriginRequestServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportProtocolOriginRequestServiceResponseHeader struct {
}

func (s ReportProtocolOriginRequestServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolOriginRequestServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportProtocolRequestServiceRequest struct {
  // {'en':'Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.', 'zh_CN':'开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (you can contact technical support for adjustment).;
  // 6.When querying daily granularity data, 00:00:00 on the following day represents the data of the current day, so the time range needs to include 00:00:00 on the following day.', 'zh_CN':'结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.默认允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整）。
  // 6.查询天粒度数据时，次日的00:00:00表示当天的数据，所以时间范围需包含次日的00:00:00。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).', 'zh_CN':'域名：
  // 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)；
  // 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'Data granularity:
  // 
  // 1m: 1 minute granularity; 
  // 5m: 5 minute granularity; Default value is 5m.
  // 1h: 1 hour granularity;
  // 1d: 1 day granularity.', 'zh_CN':'数据粒度：
  // 1m：1分钟粒度。
  // 5m：5分钟粒度。不传默认5分钟粒度
  // 1h：1小时粒度;
  // 1d：1天粒度'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Grouped dimension:
  // 
  // 1.When the "groupBy" parameter has a value, the value can only be "domain". It will group and return the detailed based on domain;
  // 2.When groupBy does not have a value, aggregate all domain by default.', 'zh_CN':'分组维度
  // 1.有传，只能传domain。按照domain分组展示明细数据;
  // 2.不传，则默认所有域名聚合。'}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s ReportProtocolRequestServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportProtocolRequestServiceRequest) SetDateFrom(v string) *ReportProtocolRequestServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportProtocolRequestServiceRequest) SetDateTo(v string) *ReportProtocolRequestServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportProtocolRequestServiceRequest) SetDomain(v []*string) *ReportProtocolRequestServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportProtocolRequestServiceRequest) SetDataInterval(v string) *ReportProtocolRequestServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportProtocolRequestServiceRequest) SetGroupBy(v string) *ReportProtocolRequestServiceRequest {
  s.GroupBy = &v
  return s
}

type ReportProtocolRequestServiceResponse struct {
  // {'en':'data', 'zh_CN':'请求结果'}
  Data []*ReportProtocolRequestServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportProtocolRequestServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportProtocolRequestServiceResponse) SetData(v []*ReportProtocolRequestServiceResponseData) *ReportProtocolRequestServiceResponse {
  s.Data = v
  return s
}

type ReportProtocolRequestServiceResponseData struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'details', 'zh_CN':'请求结果的详细数据'}
  Details []*ReportProtocolRequestServiceResponseDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportProtocolRequestServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportProtocolRequestServiceResponseData) SetDomain(v string) *ReportProtocolRequestServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportProtocolRequestServiceResponseData) SetDetails(v []*ReportProtocolRequestServiceResponseDataDetails) *ReportProtocolRequestServiceResponseData {
  s.Details = v
  return s
}

type ReportProtocolRequestServiceResponseDataDetails struct     {
  // {'en':'Time:
  // 
  // 1.When the data query granularity is 1m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00.
  // 2.When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00.
  // 3. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00.
  // 4. When the data query granularity is 1d, the format is yyyy-MM-dd; Each time slice value represents the value within the previous time granularity range.
  // 5. Return the time slices that contained in start time and in end time.', 'zh_CN':'时间，
  // 查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  // 查询的数据粒度为1d时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值；
  // 返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Number of all requests.', 'zh_CN':'全部请求数'}
  AllRequests *string `json:"allRequests,omitempty" xml:"allRequests,omitempty" require:"true"`
  // {'en':'Number of httpe requests.', 'zh_CN':'http请求数'}
  HttpRequests *string `json:"httpRequests,omitempty" xml:"httpRequests,omitempty" require:"true"`
  // {'en':'Number of https requests.', 'zh_CN':'https请求数'}
  HttpsRequests *string `json:"httpsRequests,omitempty" xml:"httpsRequests,omitempty" require:"true"`
}

func (s ReportProtocolRequestServiceResponseDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceResponseDataDetails) GoString() string {
  return s.String()
}

func (s *ReportProtocolRequestServiceResponseDataDetails) SetTimestamp(v string) *ReportProtocolRequestServiceResponseDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportProtocolRequestServiceResponseDataDetails) SetAllRequests(v string) *ReportProtocolRequestServiceResponseDataDetails {
  s.AllRequests = &v
  return s
}

func (s *ReportProtocolRequestServiceResponseDataDetails) SetHttpRequests(v string) *ReportProtocolRequestServiceResponseDataDetails {
  s.HttpRequests = &v
  return s
}

func (s *ReportProtocolRequestServiceResponseDataDetails) SetHttpsRequests(v string) *ReportProtocolRequestServiceResponseDataDetails {
  s.HttpsRequests = &v
  return s
}

type ReportProtocolRequestServicePaths struct {
}

func (s ReportProtocolRequestServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServicePaths) GoString() string {
  return s.String()
}

type ReportProtocolRequestServiceParameters struct {
}

func (s ReportProtocolRequestServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceParameters) GoString() string {
  return s.String()
}

type ReportProtocolRequestServiceRequestHeader struct {
}

func (s ReportProtocolRequestServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportProtocolRequestServiceResponseHeader struct {
}

func (s ReportProtocolRequestServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportProtocolRequestServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportStreamTaskNumberServiceRequest struct {
  // {'en':'Start Time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing time on December 2, 2018 at 10:00 am to 0 seconds);
  // 2. No more than the current time
  // 3. Up to 30 days of data available.', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2018年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近 30 天的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End Time:
  // Time Format 2019-01-02T10:00:00+08:00
  // The end time should be greater than the start time. If the end time is greater than the current time, take the current time.
  // DateFrom, dateTo both not passed, the default query past 30 minutes; If only one is not sent, throw exception
  // Maximum time interval allowed for queries: 30 minutes', 'zh_CN':'结束时间：
  // 1.时间格式2019-01-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的 30 分钟；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：30分钟'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain name:
  // The maximum number of transitive domain names by default is 200
  // Automatically filter out illegal domain names (such as passing illegal domain names, will be filtered out, the query results only return the data of the legal domain name).', 'zh_CN':'域名：
  // 1.可传递域名数量上限默认为200个
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamTaskNumberServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStreamTaskNumberServiceRequest) SetDateFrom(v string) *ReportStreamTaskNumberServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStreamTaskNumberServiceRequest) SetDateTo(v string) *ReportStreamTaskNumberServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStreamTaskNumberServiceRequest) SetDomain(v []*string) *ReportStreamTaskNumberServiceRequest {
  s.Domain = v
  return s
}

type ReportStreamTaskNumberServiceResponse struct {
  // {'en':'code', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'请求结果的详细数据'}
  Data []*ReportStreamTaskNumberServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamTaskNumberServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStreamTaskNumberServiceResponse) SetCode(v string) *ReportStreamTaskNumberServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportStreamTaskNumberServiceResponse) SetMessage(v string) *ReportStreamTaskNumberServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportStreamTaskNumberServiceResponse) SetData(v []*ReportStreamTaskNumberServiceResponseData) *ReportStreamTaskNumberServiceResponse {
  s.Data = v
  return s
}

type ReportStreamTaskNumberServiceResponseData struct     {
  // {'en':'timestamp', 'zh_CN':'时间片
  // 每一个时间片数据值代表的是前一个时间粒度范围内的数据值，比如yyyy-MM-dd 00:01，代表00:00到00:51范围内的数据。
  // 返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  DataDetail []*ReportStreamTaskNumberServiceResponseDataDataDetail `json:"dataDetail,omitempty" xml:"dataDetail,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamTaskNumberServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportStreamTaskNumberServiceResponseData) SetTimestamp(v string) *ReportStreamTaskNumberServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportStreamTaskNumberServiceResponseData) SetDataDetail(v []*ReportStreamTaskNumberServiceResponseDataDataDetail) *ReportStreamTaskNumberServiceResponseData {
  s.DataDetail = v
  return s
}

type ReportStreamTaskNumberServiceResponseDataDataDetail struct     {
  // {'en':'domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'h265', 'zh_CN':'1 or 0 , 1 代表 H265 转码，0 代表 H264 转码。'}
  H265 *string `json:"h265,omitempty" xml:"h265,omitempty" require:"true"`
  // {'en':'transcodeNum', 'zh_CN':'转码数量'}
  TranscodeNum *string `json:"transcodeNum,omitempty" xml:"transcodeNum,omitempty" require:"true"`
  // {'en':'duration', 'zh_CN':'转码时长'}
  Duration *string `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
}

func (s ReportStreamTaskNumberServiceResponseDataDataDetail) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceResponseDataDataDetail) GoString() string {
  return s.String()
}

func (s *ReportStreamTaskNumberServiceResponseDataDataDetail) SetDomain(v string) *ReportStreamTaskNumberServiceResponseDataDataDetail {
  s.Domain = &v
  return s
}

func (s *ReportStreamTaskNumberServiceResponseDataDataDetail) SetH265(v string) *ReportStreamTaskNumberServiceResponseDataDataDetail {
  s.H265 = &v
  return s
}

func (s *ReportStreamTaskNumberServiceResponseDataDataDetail) SetTranscodeNum(v string) *ReportStreamTaskNumberServiceResponseDataDataDetail {
  s.TranscodeNum = &v
  return s
}

func (s *ReportStreamTaskNumberServiceResponseDataDataDetail) SetDuration(v string) *ReportStreamTaskNumberServiceResponseDataDataDetail {
  s.Duration = &v
  return s
}

type ReportStreamTaskNumberServicePaths struct {
}

func (s ReportStreamTaskNumberServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServicePaths) GoString() string {
  return s.String()
}

type ReportStreamTaskNumberServiceParameters struct {
}

func (s ReportStreamTaskNumberServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceParameters) GoString() string {
  return s.String()
}

type ReportStreamTaskNumberServiceRequestHeader struct {
}

func (s ReportStreamTaskNumberServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStreamTaskNumberServiceResponseHeader struct {
}

func (s ReportStreamTaskNumberServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamTaskNumberServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainTotalRequestRequest struct {
  // {"en":"Domain list
  // Domain number limits can be adjusted depending on different accounts. The default value is 20(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表
  // 1.域名个数限制根据账号可调，默认为20个（可联系技术支持下单调整）；"}
  QueryDomainTotalRequestDomainList *QueryDomainTotalRequestDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryDomainTotalRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestRequest) SetDomainList(v *QueryDomainTotalRequestDomainList) *QueryDomainTotalRequestRequest {
  s.QueryDomainTotalRequestDomainList = v
  return s
}

type QueryDomainTotalRequestDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalRequestDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestDomainList) SetDomainName(v []*string) *QueryDomainTotalRequestDomainList {
  s.DomainName = v
  return s
}

type QueryDomainTotalRequestResponse struct {
  // {"en":"Total requests", "zh_CN":"总请求数"}
  HitSummary *int `json:"hit-summary,omitempty" xml:"hit-summary,omitempty" require:"true"`
  // {'en':'hitData', 'zh_CN':'请求数据'}
  HitData []*QueryDomainTotalRequestResponseHitData `json:"hit-data,omitempty" xml:"hit-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestResponse) SetHitSummary(v int) *QueryDomainTotalRequestResponse {
  s.HitSummary = &v
  return s
}

func (s *QueryDomainTotalRequestResponse) SetHitData(v []*QueryDomainTotalRequestResponseHitData) *QueryDomainTotalRequestResponse {
  s.HitData = v
  return s
}

type QueryDomainTotalRequestResponseHitData struct     {
  // {"en":"Date
  // When the querying data granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00;When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24;When the querying data granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the data;Return the time slices contained in start time and in end time", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是yyyy-MM-dd 24
  // 3.查询的数据粒度为daily时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值；
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests, More than 6 digits are displayed in scientific notation,e.g.:1642565=1.642565E6", "zh_CN":"请求数,超过6位数的以科学计数展示，例：1642565=1.642565E6"}
  Hit *int `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
}

func (s QueryDomainTotalRequestResponseHitData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestResponseHitData) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestResponseHitData) SetTimestamp(v string) *QueryDomainTotalRequestResponseHitData {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainTotalRequestResponseHitData) SetHit(v int) *QueryDomainTotalRequestResponseHitData {
  s.Hit = &v
  return s
}

type QueryDomainTotalRequestPaths struct {
}

func (s QueryDomainTotalRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestPaths) GoString() string {
  return s.String()
}

type QueryDomainTotalRequestParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo should not be longer than 31 days;", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过31天；
  // 4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2.Must be greater than dateFrom;
  // 3.If it is greater than the current time, then the current time will be assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；
  // 3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in the granularity of actual configuration when there is specific configuration of the data collecting granularity for the customer", "zh_CN":"数据粒度
  // 1.fiveminutes：5分钟，hourly：1小时，daily：1天；
  // 2.不传递，默认为daily；
  // 3.传递fiveminutes时，若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDomainTotalRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestParameters) SetDateFrom(v string) *QueryDomainTotalRequestParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainTotalRequestParameters) SetDateTo(v string) *QueryDomainTotalRequestParameters {
  s.DateTo = &v
  return s
}

func (s *QueryDomainTotalRequestParameters) SetType(v string) *QueryDomainTotalRequestParameters {
  s.Type = &v
  return s
}

type QueryDomainTotalRequestRequestHeader struct {
}

func (s QueryDomainTotalRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainTotalRequestResponseHeader struct {
}

func (s QueryDomainTotalRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestResponseHeader) GoString() string {
  return s.String()
}




type HitRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"This parameter specifies if the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
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
  // {"en":"Display statistic result in merged or separate way
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s HitRequest) String() string {
  return tea.Prettify(s)
}

func (s HitRequest) GoString() string {
  return s.String()
}

func (s *HitRequest) SetCust(v string) *HitRequest {
  s.Cust = &v
  return s
}

func (s *HitRequest) SetDate(v string) *HitRequest {
  s.Date = &v
  return s
}

func (s *HitRequest) SetStartdate(v string) *HitRequest {
  s.Startdate = &v
  return s
}

func (s *HitRequest) SetEnddate(v string) *HitRequest {
  s.Enddate = &v
  return s
}

func (s *HitRequest) SetTimezone(v string) *HitRequest {
  s.Timezone = &v
  return s
}

func (s *HitRequest) SetChannel(v string) *HitRequest {
  s.Channel = &v
  return s
}

func (s *HitRequest) SetIsExactMatch(v string) *HitRequest {
  s.IsExactMatch = &v
  return s
}

func (s *HitRequest) SetRegion(v string) *HitRequest {
  s.Region = &v
  return s
}

func (s *HitRequest) SetAccetype(v string) *HitRequest {
  s.Accetype = &v
  return s
}

func (s *HitRequest) SetDataformat(v string) *HitRequest {
  s.Dataformat = &v
  return s
}

func (s *HitRequest) SetResultType(v string) *HitRequest {
  s.ResultType = &v
  return s
}

type HitResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *HitResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s HitResponse) String() string {
  return tea.Prettify(s)
}

func (s HitResponse) GoString() string {
  return s.String()
}

func (s *HitResponse) SetProvider(v *HitResponseProvider) *HitResponse {
  s.Provider = v
  return s
}

type HitResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'请求数数据'}
  Date *HitResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s HitResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s HitResponseProvider) GoString() string {
  return s.String()
}

func (s *HitResponseProvider) SetName(v string) *HitResponseProvider {
  s.Name = &v
  return s
}

func (s *HitResponseProvider) SetType(v string) *HitResponseProvider {
  s.Type = &v
  return s
}

func (s *HitResponseProvider) SetResultType(v string) *HitResponseProvider {
  s.ResultType = &v
  return s
}

func (s *HitResponseProvider) SetDate(v *HitResponseProviderDate) *HitResponseProvider {
  s.Date = v
  return s
}

type HitResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *HitResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s HitResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s HitResponseProviderDate) GoString() string {
  return s.String()
}

func (s *HitResponseProviderDate) SetStartdate(v string) *HitResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *HitResponseProviderDate) SetEnddate(v string) *HitResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *HitResponseProviderDate) SetChannel(v *HitResponseProviderDateChannel) *HitResponseProviderDate {
  s.Channel = v
  return s
}

type HitResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'请求数数据'}
  Bandwidth []*HitResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s HitResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s HitResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *HitResponseProviderDateChannel) SetName(v string) *HitResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *HitResponseProviderDateChannel) SetBandwidth(v []*HitResponseProviderDateChannelBandwidth) *HitResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type HitResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'请求数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s HitResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s HitResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *HitResponseProviderDateChannelBandwidth) SetTime(v string) *HitResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *HitResponseProviderDateChannelBandwidth) SetText(v string) *HitResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type HitPaths struct {
}

func (s HitPaths) String() string {
  return tea.Prettify(s)
}

func (s HitPaths) GoString() string {
  return s.String()
}

type HitParameters struct {
}

func (s HitParameters) String() string {
  return tea.Prettify(s)
}

func (s HitParameters) GoString() string {
  return s.String()
}

type HitRequestHeader struct {
}

func (s HitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s HitRequestHeader) GoString() string {
  return s.String()
}

type HitResponseHeader struct {
}

func (s HitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s HitResponseHeader) GoString() string {
  return s.String()
}




type ReportMinuteUserRequestIspProvinceServiceRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:00:00 Beijing time on December 2, 2016);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
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
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days. ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment)
  // 
  // 2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment).", "zh_CN":"域名：
  // 
  // 1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)；
  // 
  // 2.有传递domain时：域名最多支持传20个（可联系技术支持调整）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Province:
  // 
  // 1. Province is not upload: Query all provinces and aggregate the returned data according to all provinces;
  // 
  // 2. Province is upload: Support upload multiple provinces. Upload the Chinese name of the province, such as: beijing, shanghai. Please refer to the appendix description section of the overview page for the provincial information code table.", "zh_CN":"省份：
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：支持传多个。传递省份code，如：beijing、shanghai。省份信息码表详见概览页附录说明章节"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 
  // 1. ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;
  // 
  // 2. ISPs is upload: Support upload multiple ISPs. Upload the Chinese name of the ISPs, such as: dx, tt. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商：
  // 
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。
  // 
  // 2.有传递isp时：支持传多个。传递运营商code，如：dx、tt。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  // 1m: 1 minute granularity.
  // 
  // 5m: 5 minute granularity; Default value is 5m.
  // 
  // 1h: 1 hour granularity.", "zh_CN":"数据粒度：
  // 
  // 1m：1分钟粒度。
  // 
  // 5m：5分钟粒度。不传默认5分钟粒度
  // 
  // 1h：1小时粒度;"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Grouped dimension:
  // 
  // 1. The optional values are province, isp,domain;Multiple values can be uploaded;
  // 
  // 2. If no value is uploaded: Aggregate all data by default.", "zh_CN":"分组维度
  // 可选值为province、isp、domain，可传入多个值；
  // 有传入则按照该维度展示明细数据；
  // 没传默认全部聚合。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportMinuteUserRequestIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetDateFrom(v string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetDateTo(v string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetDomain(v []*string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetProvince(v []*string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetIsp(v []*string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetDataInterval(v string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportMinuteUserRequestIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportMinuteUserRequestIspProvinceServiceResponse struct {
  Data []*ReportMinuteUserRequestIspProvinceServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportMinuteUserRequestIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponse) SetData(v []*ReportMinuteUserRequestIspProvinceServiceResponseData) *ReportMinuteUserRequestIspProvinceServiceResponse {
  s.Data = v
  return s
}

type ReportMinuteUserRequestIspProvinceServiceResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportMinuteUserRequestIspProvinceServiceResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseData) SetDomain(v string) *ReportMinuteUserRequestIspProvinceServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseData) SetIspData(v []*ReportMinuteUserRequestIspProvinceServiceResponseDataIspData) *ReportMinuteUserRequestIspProvinceServiceResponseData {
  s.IspData = v
  return s
}

type ReportMinuteUserRequestIspProvinceServiceResponseDataIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseDataIspData) GoString() string {
  return s.String()
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspData) SetIsp(v string) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspData {
  s.Isp = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspData) SetProvinceData(v []*ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspData {
  s.ProvinceData = v
  return s
}

type ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Total number of requests.", "zh_CN":"省份运营商的请求总数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  Details []*ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData) SetProvince(v string) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData) SetTotalRequest(v string) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData {
  s.TotalRequest = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData) SetDetails(v []*ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceData {
  s.Details = v
  return s
}

type ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails struct     {
  // {"en":"ime:
  // 
  // 1. When the data query granularity is 1m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  // 2. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  // 3. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00;
  // 
  // 4. Return the time slices that contained in start time and in end time.", "zh_CN":"时间，
  // 查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests.", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetTimestamp(v string) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetRequest(v string) *ReportMinuteUserRequestIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Request = &v
  return s
}

type ReportMinuteUserRequestIspProvinceServicePaths struct {
}

func (s ReportMinuteUserRequestIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportMinuteUserRequestIspProvinceServiceParameters struct {
}

func (s ReportMinuteUserRequestIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportMinuteUserRequestIspProvinceServiceRequestHeader struct {
}

func (s ReportMinuteUserRequestIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportMinuteUserRequestIspProvinceServiceResponseHeader struct {
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportMinuteUserRequestIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestHitRatioRequest struct {
  // {"en":"From date:
  //         1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example: 2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8)
  //         2.Cannot exceed current time
  //         3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2018年12月2日10点0分0秒)；
  //         2.不能大于当前时间
  //         3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"To time:
  //         1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example: 2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8)
  //         2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  //         3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  //         4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days", "zh_CN":"结束时间:
  //         1.时间格式2019-01-02T10:00:00+08:00
  //         2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  //         3.dateFrom,dateTo二者都未传,默认查询过去的24小时；如仅有一个未传,抛异常
  //         4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  //         1.The maximum number of deliverable domain names is 200 by default
  // 		2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  //         3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.", "zh_CN":"域名:
  //         1.可传递域名数量上限默认为200个
  //         2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  //         3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  // 1m: 1 minute granularity.
  // 
  // 5m: 5 minute granularity.
  // 
  // 1h: 1 hour granularity.
  // 
  // 1d: 1 day granularity. Default value is 1d.", "zh_CN":"数据粒度：
  // 
  // 支持1m（1分钟）、5m（5分钟）、1h（1小时）、1d（天）
  // 不传默认1d。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryRequestHitRatioRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioRequest) SetDateFrom(v string) *QueryRequestHitRatioRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestHitRatioRequest) SetDateTo(v string) *QueryRequestHitRatioRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestHitRatioRequest) SetDomain(v []*string) *QueryRequestHitRatioRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestHitRatioRequest) SetDataInterval(v string) *QueryRequestHitRatioRequest {
  s.DataInterval = &v
  return s
}

type QueryRequestHitRatioResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryRequestHitRatioResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestHitRatioResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioResponse) SetCode(v string) *QueryRequestHitRatioResponse {
  s.Code = &v
  return s
}

func (s *QueryRequestHitRatioResponse) SetMessage(v string) *QueryRequestHitRatioResponse {
  s.Message = &v
  return s
}

func (s *QueryRequestHitRatioResponse) SetData(v []*QueryRequestHitRatioResponseData) *QueryRequestHitRatioResponse {
  s.Data = v
  return s
}

type QueryRequestHitRatioResponseData struct     {
  // {"en":"Actually processed time. yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间,格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total hit ratio.", "zh_CN":"总命中率的平均值"}
  TotalAvg *float64 `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"hitRatioDatas", "zh_CN":"缓存命中率数据"}
  HitRatioDatas []*QueryRequestHitRatioResponseDataHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestHitRatioResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponseData) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioResponseData) SetRealDate(v string) *QueryRequestHitRatioResponseData {
  s.RealDate = &v
  return s
}

func (s *QueryRequestHitRatioResponseData) SetTotalAvg(v float64) *QueryRequestHitRatioResponseData {
  s.TotalAvg = &v
  return s
}

func (s *QueryRequestHitRatioResponseData) SetHitRatioDatas(v []*QueryRequestHitRatioResponseDataHitRatioDatas) *QueryRequestHitRatioResponseData {
  s.HitRatioDatas = v
  return s
}

type QueryRequestHitRatioResponseDataHitRatioDatas struct     {
  // {"en":"timetamp
  // 1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range, For example yyyy-MM-dd 00:05 represents data in the range from 00:00 to 00:05.
  // 2.The data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents data values in the previous time granularity range such as yyyy-MM-dd 01 that represent data from 00 to 01.
  // 3. the data granularity of the query is daily, the format is yyyy-MM-dd; Each time slice data value represents the data value for that day.
  // 4.Returns the timetamp contained in start time and end time.", "zh_CN":"时间，
  // 
  // 查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Cache hit ratio,keep 4 decimal places", "zh_CN":"缓存命中率,保留4位小数"}
  HitRatio *string `json:"hitRatio,omitempty" xml:"hitRatio,omitempty" require:"true"`
}

func (s QueryRequestHitRatioResponseDataHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponseDataHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioResponseDataHitRatioDatas) SetTimestamp(v string) *QueryRequestHitRatioResponseDataHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestHitRatioResponseDataHitRatioDatas) SetHitRatio(v string) *QueryRequestHitRatioResponseDataHitRatioDatas {
  s.HitRatio = &v
  return s
}

type QueryRequestHitRatioPaths struct {
}

func (s QueryRequestHitRatioPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioPaths) GoString() string {
  return s.String()
}

type QueryRequestHitRatioParameters struct {
}

func (s QueryRequestHitRatioParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioParameters) GoString() string {
  return s.String()
}

type QueryRequestHitRatioRequestHeader struct {
}

func (s QueryRequestHitRatioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestHitRatioResponseHeader struct {
}

func (s QueryRequestHitRatioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestBySpecificProtocolRequest struct {
  // {"en":"Start time:
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于当前时间-183天，并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过7天（可联系技术支持调整）；
  // 4.dateFrom和dateTo要么都传递，要么都不传递；
  // 5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；
  // 3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名，域名个数限制根据账号可调，默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity: 1m , 5m, support 5m", "zh_CN":"数据粒度 : 1m  , 5m ，默认5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Transmission protocol
  // 
  // 1.Options: http, https;
  // 2.https is used as the default value is no value specified;
  // 3.httpRequestData is displayed if http is queried, and httpsRequestData is displayed if https is queried;", "zh_CN":"传输协议
  // 1.可选值为http、https；
  // 2.不传默认查询https；
  // 3.查询http时出参展示httpRequestData，查询https时出参展示httpsRequestData；"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain；
  // 2.有传入则按照该维度展示明细数据；"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryRequestBySpecificProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolRequest) SetDateFrom(v string) *QueryRequestBySpecificProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetDateTo(v string) *QueryRequestBySpecificProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetDomain(v []*string) *QueryRequestBySpecificProtocolRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetDataInterval(v string) *QueryRequestBySpecificProtocolRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetProtocolType(v string) *QueryRequestBySpecificProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetGroupBy(v []*string) *QueryRequestBySpecificProtocolRequest {
  s.GroupBy = v
  return s
}

type QueryRequestBySpecificProtocolResponse struct {
  Result []*QueryRequestBySpecificProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestBySpecificProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolResponse) SetResult(v []*QueryRequestBySpecificProtocolResponseResult) *QueryRequestBySpecificProtocolResponse {
  s.Result = v
  return s
}

type QueryRequestBySpecificProtocolResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  HttpsRequestData []*QueryRequestBySpecificProtocolResponseResultHttpsRequestData `json:"httpsRequestData,omitempty" xml:"httpsRequestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestBySpecificProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolResponseResult) SetDomain(v string) *QueryRequestBySpecificProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryRequestBySpecificProtocolResponseResult) SetHttpsRequestData(v []*QueryRequestBySpecificProtocolResponseResultHttpsRequestData) *QueryRequestBySpecificProtocolResponseResult {
  s.HttpsRequestData = v
  return s
}

type QueryRequestBySpecificProtocolResponseResultHttpsRequestData struct     {
  // {"en":"DateTime, the format is  yyyy-MM-dd HH:mm; the data value of every time slice represents the data  value within the previous time granularity range.", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRequestBySpecificProtocolResponseResultHttpsRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponseResultHttpsRequestData) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolResponseResultHttpsRequestData) SetTimestamp(v string) *QueryRequestBySpecificProtocolResponseResultHttpsRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestBySpecificProtocolResponseResultHttpsRequestData) SetValue(v string) *QueryRequestBySpecificProtocolResponseResultHttpsRequestData {
  s.Value = &v
  return s
}

type QueryRequestBySpecificProtocolPaths struct {
}

func (s QueryRequestBySpecificProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolPaths) GoString() string {
  return s.String()
}

type QueryRequestBySpecificProtocolParameters struct {
}

func (s QueryRequestBySpecificProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolParameters) GoString() string {
  return s.String()
}

type QueryRequestBySpecificProtocolRequestHeader struct {
}

func (s QueryRequestBySpecificProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestBySpecificProtocolResponseHeader struct {
}

func (s QueryRequestBySpecificProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponseHeader) GoString() string {
  return s.String()
}




type QueryIPV6RequestOfeachISPandProvinceRequest struct {
  // {"en":"Start time: 
  //     1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time); 
  //     2. No bigger than the current time;
  //     3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）;
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 1. the time format is 2016-12-02T10:00:00+08:00;
  //     2.End time should be greater than start time. If the end time is greater than current time, current time will be used;
  //     3.If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur;
  //     4.Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support).", "zh_CN":"结束时间：
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间;
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时，如仅有一个未传，抛异常;
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name: 1. The maximum number of TLDs is 20 by default (Technical Support Adjustment can be contacted); 
  //     2. automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为200个（可联系技术支持调整）;
  // 2.自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity 1. Support 5m (5 minute granularity), 1h (1 hour granularity); 2. The default is 5m;", "zh_CN":"数据粒度：
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
  // 	1.The optional values are IPv6 and IPv4;
  // 	2.If let this parameter empty,it will query all IP type.", "zh_CN":"IP类型：
  // 1.可选值为 IPV6、IPV4;
  // 2.不传默认查询全部"}
  IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
  // {"en":"Grouped dimension:
  //     1.Aggregation date by default;
  //     2.the optional value is domain,province,isp,allow to send multi option ;
  //     3.send the Grouped dimension represent the need to display details by their corresponding values.For example, when groupBy is isp, the ISP dimension needs to be displayed in detail. When an ISP is not passed, it represents an aggregate date and would not return the ISP node.Provinces and domains have the same logic. For example, by passing 'groupBy': ['domain', 'province'], the ISP node under ispData does not need to return. {domain:'www.aaaa.com','ispData': [{'isp','China Telecom','provinceData': [...]}}",
  //     "zh_CN":"分组关键词：
  // 1.默认聚合展示;
  // 2.可选值为domain.province.isp，可传入多个值;
  // 3.传入关键词则代表需要按照关键词对应的值展示明细; 例如groupBy传入isp，则isp维度需要明细展示;当没有传递isp，则代表isp聚合展示，同时isp节点则不返回。其他province和domain相同逻辑。 例如：传递'groupBy':   ['domain','province']，则ispData下的isp节点无需返回。 { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp':   '中国电信', 'provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDateTo(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetIPType(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.IPType = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponse struct {
  // {"en":"", "zh_CN":""}
  Result []*QueryIPV6RequestOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponse) SetResult(v []*QueryIPV6RequestOfeachISPandProvinceResponseResult) *QueryIPV6RequestOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResult struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResult) SetIspData(v []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) *QueryIPV6RequestOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"isp", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  RequestData []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) SetRequestData(v []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.RequestData = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData struct     {
  // {"en":"1.When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm (the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00 ); 
  //                     2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24;
  //                     3.Return the time slices that contained in start time and in end time.", "zh_CN":"时间，
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 2.查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) SetTimestamp(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) SetValue(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData {
  s.Value = &v
  return s
}

type QueryIPV6RequestOfeachISPandProvincePaths struct {
}

func (s QueryIPV6RequestOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryIPV6RequestOfeachISPandProvinceParameters struct {
}

func (s QueryIPV6RequestOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryIPV6RequestOfeachISPandProvinceRequestHeader struct {
}

func (s QueryIPV6RequestOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryIPV6RequestOfeachISPandProvinceResponseHeader struct {
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestNumbersUnderShieldPoPRequest struct {
  // {"en":"Start time:
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, support 5m (granularity of 5 minutes)", "zh_CN":"数据粒度,支持5m: 5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryRequestNumbersUnderShieldPoPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDateFrom(v string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDateTo(v string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDomain(v []*string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDataInterval(v string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetGroupBy(v []*string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.GroupBy = v
  return s
}

type QueryRequestNumbersUnderShieldPoPResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryRequestNumbersUnderShieldPoPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestNumbersUnderShieldPoPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPResponse) SetResult(v []*QueryRequestNumbersUnderShieldPoPResponseResult) *QueryRequestNumbersUnderShieldPoPResponse {
  s.Result = v
  return s
}

type QueryRequestNumbersUnderShieldPoPResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak Time", "zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak Request", "zh_CN":"请求数峰值"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Total Request", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  RequestData []*QueryRequestNumbersUnderShieldPoPResponseResultRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestNumbersUnderShieldPoPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetDomain(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetPeakTime(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.PeakTime = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetPeakValue(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.PeakValue = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetTotalRequest(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetRequestData(v []*QueryRequestNumbersUnderShieldPoPResponseResultRequestData) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.RequestData = v
  return s
}

type QueryRequestNumbersUnderShieldPoPResponseResultRequestData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data  value within the previous time granularity range.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRequestNumbersUnderShieldPoPResponseResultRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponseResultRequestData) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResultRequestData) SetTimestamp(v string) *QueryRequestNumbersUnderShieldPoPResponseResultRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResultRequestData) SetValue(v string) *QueryRequestNumbersUnderShieldPoPResponseResultRequestData {
  s.Value = &v
  return s
}

type QueryRequestNumbersUnderShieldPoPPaths struct {
}

func (s QueryRequestNumbersUnderShieldPoPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPPaths) GoString() string {
  return s.String()
}

type QueryRequestNumbersUnderShieldPoPParameters struct {
}

func (s QueryRequestNumbersUnderShieldPoPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPParameters) GoString() string {
  return s.String()
}

type QueryRequestNumbersUnderShieldPoPRequestHeader struct {
}

func (s QueryRequestNumbersUnderShieldPoPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestNumbersUnderShieldPoPResponseHeader struct {
}

func (s QueryRequestNumbersUnderShieldPoPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponseHeader) GoString() string {
  return s.String()
}




type ReportRequestIspProvinceServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. can not exceed the current time;
  // 3. the latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒)；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. the end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 4. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust).", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间；
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时；如仅有一个未传,抛异常；
  // 4.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递该入参时查询账号下所有域名,但当账号下域名数量超过限制时不可查询(报错)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: 5-minute; granularity, 1h: 1-hour granularity", "zh_CN":"数据粒度,5m:5分钟粒度,1h:1小时粒度"}
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
  // 1. Options are domain, province, isp, and more than one value can be entered;
  // 2. The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、province、isp,可传入多个值；
  // 2.有传入则按照该维度展示明细数据；"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportRequestIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportRequestIspProvinceServiceRequest) SetDateFrom(v string) *ReportRequestIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportRequestIspProvinceServiceRequest) SetDateTo(v string) *ReportRequestIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportRequestIspProvinceServiceRequest) SetDomain(v []*string) *ReportRequestIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportRequestIspProvinceServiceRequest) SetDataInterval(v string) *ReportRequestIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportRequestIspProvinceServiceRequest) SetProvince(v []*string) *ReportRequestIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportRequestIspProvinceServiceRequest) SetIsp(v []*string) *ReportRequestIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportRequestIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportRequestIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportRequestIspProvinceServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportRequestIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportRequestIspProvinceServiceResponse) SetResult(v []*ReportRequestIspProvinceServiceResponseResult) *ReportRequestIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportRequestIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ispData", "zh_CN":"isp数据"}
  IspData []*ReportRequestIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportRequestIspProvinceServiceResponseResult) SetDomain(v string) *ReportRequestIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportRequestIspProvinceServiceResponseResult) SetIspData(v []*ReportRequestIspProvinceServiceResponseResultIspData) *ReportRequestIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportRequestIspProvinceServiceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"provinceData", "zh_CN":"省份数据"}
  ProvinceData []*ReportRequestIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportRequestIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportRequestIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportRequestIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportRequestIspProvinceServiceResponseResultIspDataProvinceData) *ReportRequestIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportRequestIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"qps", "zh_CN":"qps数据"}
  RequestData []*ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportRequestIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportRequestIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportRequestIspProvinceServiceResponseResultIspDataProvinceData) SetRequestData(v []*ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) *ReportRequestIspProvinceServiceResponseResultIspDataProvinceData {
  s.RequestData = v
  return s
}

type ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData struct     {
  // {"en":"Time,", "zh_CN":"时间,
  //                1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00；
  //                2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1);00；
  //                3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Number of requests within unit ; time. Keep two digits of decimals.", "zh_CN":"单位时间内的请求数,保留2位小数"}
  Qps *string `json:"qps,omitempty" xml:"qps,omitempty" require:"true"`
}

func (s ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) SetTimestamp(v string) *ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) SetValue(v string) *ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData {
  s.Value = &v
  return s
}

func (s *ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData) SetQps(v string) *ReportRequestIspProvinceServiceResponseResultIspDataProvinceDataRequestData {
  s.Qps = &v
  return s
}

type ReportRequestIspProvinceServicePaths struct {
}

func (s ReportRequestIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportRequestIspProvinceServiceParameters struct {
}

func (s ReportRequestIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportRequestIspProvinceServiceRequestHeader struct {
}

func (s ReportRequestIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportRequestIspProvinceServiceResponseHeader struct {
}

func (s ReportRequestIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type PolicyDetailReportRequest struct {
  // {"en":"domain id", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"policy Id", "zh_CN":"策略id"}
  PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"begin time", "zh_CN":"开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"end time", "zh_CN":"结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s PolicyDetailReportRequest) String() string {
  return tea.Prettify(s)
}

func (s PolicyDetailReportRequest) GoString() string {
  return s.String()
}

func (s *PolicyDetailReportRequest) SetDomainId(v int) *PolicyDetailReportRequest {
  s.DomainId = &v
  return s
}

func (s *PolicyDetailReportRequest) SetPolicyId(v string) *PolicyDetailReportRequest {
  s.PolicyId = &v
  return s
}

func (s *PolicyDetailReportRequest) SetStartTime(v string) *PolicyDetailReportRequest {
  s.StartTime = &v
  return s
}

func (s *PolicyDetailReportRequest) SetEndTime(v string) *PolicyDetailReportRequest {
  s.EndTime = &v
  return s
}

type PolicyDetailReportResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"detail", "zh_CN":"详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s PolicyDetailReportResponse) String() string {
  return tea.Prettify(s)
}

func (s PolicyDetailReportResponse) GoString() string {
  return s.String()
}

func (s *PolicyDetailReportResponse) SetResCode(v string) *PolicyDetailReportResponse {
  s.ResCode = &v
  return s
}

func (s *PolicyDetailReportResponse) SetMsg(v string) *PolicyDetailReportResponse {
  s.Msg = &v
  return s
}

type PolicyDetailReportPaths struct {
}

func (s PolicyDetailReportPaths) String() string {
  return tea.Prettify(s)
}

func (s PolicyDetailReportPaths) GoString() string {
  return s.String()
}

type PolicyDetailReportParameters struct {
}

func (s PolicyDetailReportParameters) String() string {
  return tea.Prettify(s)
}

func (s PolicyDetailReportParameters) GoString() string {
  return s.String()
}

type PolicyDetailReportRequestHeader struct {
}

func (s PolicyDetailReportRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PolicyDetailReportRequestHeader) GoString() string {
  return s.String()
}

type PolicyDetailReportResponseHeader struct {
}

func (s PolicyDetailReportResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PolicyDetailReportResponseHeader) GoString() string {
  return s.String()
}




type ReportRequestHitRateIspProvinceServiceRequest struct {
  // {"en":"Starting time:
  //         1.The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00) );
  //         2.Cannot be greater than the current time
  //         3.Up to the data for the last six months (183 days) can be obtained.", "zh_CN":"开始时间:
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  //         2.不能大于当前时间
  //         3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  //         1.Time format yyyy-MM-ddTHH:mm:ss+08:00
  //         2.The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  //         3.dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one undelivered, throw an exception
  //         4.Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo cannot exceed 7 days","zh_CN":"结束时间:
  //         1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  //         2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  //         3.dateFrom,dateTo二者都未传,默认查询过去的24小时；如仅有一个未传,抛异常
  //         4.允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天。
  //         （可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"domain name:
  // 
  //         1.The maximum number of deliverable domain names is 20 by default (can be contacted by technical support);
  //         2.Automatically filter out invalid domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of valid domain names).", "zh_CN":"域名:
  //         1.可传递域名数量上限默认为20个（可联系技术支持调整）；
  //         2.自动过滤掉无效域名（如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  //         1.support 5m (5 minutes), 1h (1 hour)
  // 
  //         2.do not pass the default 5m.", "zh_CN":"数据粒度:
  //         1.支持5m（5分钟）、1h（1小时）
  //         2.不传默认5m。"}
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
  // {"en":"Group keywords:
  // 
  //         1. the default aggregate display;
  //         2. The optional values are domain, province, and isp, which can pass multiple values.
  //         3. The incoming keyword means that the details need to be displayed according to the value corresponding to the keyword;
  //         For example, when groupBy is passed into isp, the isp dimension needs to be displayed in detail; when it is not passed, it represents the isp aggregation display, and the isp node does not return. Other provinces and domains have the same logic.
  //         For example: pass 'groupBy': ['domain', 'province'], then the isp node under ispData does not need to return.
  //         { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp': 'China Telecom', 'provinceData': [....] }]}", "zh_CN":"分组关键词:
  //         1.默认聚合展示；
  //         2.可选值为domain、province、isp,可传入多个值；
  //         3.传入关键词则代表需要按照关键词对应的值展示明细；
  //         例如groupBy传入isp,则isp维度需要明细展示；当没有传递isp,则代表isp聚合展示,同时isp节点则不返回。其他province和domain相同逻辑。
  //         例如:传递'groupBy':  ['domain','province'],则ispData下的isp节点无需返回。
  //         { 'domain': 'www.aaaa.com', 'ispData': [ {'isp':  '中国电信','provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportRequestHitRateIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetDateFrom(v string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetDateTo(v string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetDomain(v []*string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetDataInterval(v string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetProvince(v []*string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetIsp(v []*string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportRequestHitRateIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportRequestHitRateIspProvinceServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportRequestHitRateIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestHitRateIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportRequestHitRateIspProvinceServiceResponse) SetResult(v []*ReportRequestHitRateIspProvinceServiceResponseResult) *ReportRequestHitRateIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportRequestHitRateIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"ISP数据"}
  IspData []*ReportRequestHitRateIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestHitRateIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResult) SetDomain(v string) *ReportRequestHitRateIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResult) SetIspData(v []*ReportRequestHitRateIspProvinceServiceResponseResultIspData) *ReportRequestHitRateIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportRequestHitRateIspProvinceServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份数据"}
  ProvinceData []*ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestHitRateIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportRequestHitRateIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData) *ReportRequestHitRateIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Request number hit rate, retain 4 decimal places", "zh_CN":"求数命中率数据"}
  HitRateData []*ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData `json:"hitRateData,omitempty" xml:"hitRateData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData) SetHitRateData(v []*ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceData {
  s.HitRateData = v
  return s
}

type ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData struct     {
  // {"en":"time,
  //         1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00;
  //         2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00;
  //         3. Return to the time slice included in the start time and end time.", "zh_CN":"时间,
  //         1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是（yyyy-MM-dd+1）00:00；
  //         2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是（yyyy-MM-dd+1）00；
  //         3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Request number hit rate, retain 4 decimal places", "zh_CN":"请求数命中率,保留4位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Number of hit requests. (This field is not provided by default, if you need this field, you can contact configuration support)", "zh_CN":"命中请求数
  //         (该字段默认不提供,需要该字段可联系配置支持)"}
  HitRequest *string `json:"hitRequest,omitempty" xml:"hitRequest,omitempty" require:"true"`
}

func (s ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) GoString() string {
  return s.String()
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) SetTimestamp(v string) *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData {
  s.Timestamp = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) SetValue(v string) *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData {
  s.Value = &v
  return s
}

func (s *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData) SetHitRequest(v string) *ReportRequestHitRateIspProvinceServiceResponseResultIspDataProvinceDataHitRateData {
  s.HitRequest = &v
  return s
}

type ReportRequestHitRateIspProvinceServicePaths struct {
}

func (s ReportRequestHitRateIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportRequestHitRateIspProvinceServiceParameters struct {
}

func (s ReportRequestHitRateIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportRequestHitRateIspProvinceServiceRequestHeader struct {
}

func (s ReportRequestHitRateIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportRequestHitRateIspProvinceServiceResponseHeader struct {
}

func (s ReportRequestHitRateIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestHitRateIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type DispatchDomainAreaTimeReportRequest struct {
  // {"en":"domain id", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"view id", "zh_CN":"线路id 多个用英文逗号,隔开。为空表示查所有。"}
  ViewIds *string `json:"viewIds,omitempty" xml:"viewIds,omitempty"`
  // {"en":"begin time", "zh_CN":"开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"end time", "zh_CN":"结算时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s DispatchDomainAreaTimeReportRequest) String() string {
  return tea.Prettify(s)
}

func (s DispatchDomainAreaTimeReportRequest) GoString() string {
  return s.String()
}

func (s *DispatchDomainAreaTimeReportRequest) SetDomainId(v int) *DispatchDomainAreaTimeReportRequest {
  s.DomainId = &v
  return s
}

func (s *DispatchDomainAreaTimeReportRequest) SetViewIds(v string) *DispatchDomainAreaTimeReportRequest {
  s.ViewIds = &v
  return s
}

func (s *DispatchDomainAreaTimeReportRequest) SetStartTime(v string) *DispatchDomainAreaTimeReportRequest {
  s.StartTime = &v
  return s
}

func (s *DispatchDomainAreaTimeReportRequest) SetEndTime(v string) *DispatchDomainAreaTimeReportRequest {
  s.EndTime = &v
  return s
}

type DispatchDomainAreaTimeReportResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"detail", "zh_CN":"详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DispatchDomainAreaTimeReportResponse) String() string {
  return tea.Prettify(s)
}

func (s DispatchDomainAreaTimeReportResponse) GoString() string {
  return s.String()
}

func (s *DispatchDomainAreaTimeReportResponse) SetResCode(v string) *DispatchDomainAreaTimeReportResponse {
  s.ResCode = &v
  return s
}

func (s *DispatchDomainAreaTimeReportResponse) SetMsg(v string) *DispatchDomainAreaTimeReportResponse {
  s.Msg = &v
  return s
}

type DispatchDomainAreaTimeReportPaths struct {
}

func (s DispatchDomainAreaTimeReportPaths) String() string {
  return tea.Prettify(s)
}

func (s DispatchDomainAreaTimeReportPaths) GoString() string {
  return s.String()
}

type DispatchDomainAreaTimeReportParameters struct {
}

func (s DispatchDomainAreaTimeReportParameters) String() string {
  return tea.Prettify(s)
}

func (s DispatchDomainAreaTimeReportParameters) GoString() string {
  return s.String()
}

type DispatchDomainAreaTimeReportRequestHeader struct {
}

func (s DispatchDomainAreaTimeReportRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DispatchDomainAreaTimeReportRequestHeader) GoString() string {
  return s.String()
}

type DispatchDomainAreaTimeReportResponseHeader struct {
}

func (s DispatchDomainAreaTimeReportResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DispatchDomainAreaTimeReportResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeRequestHitRatioServiceRequest struct {
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

func (s QueryEdgeRequestHitRatioServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceRequest) SetDateFrom(v string) *QueryEdgeRequestHitRatioServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceRequest) SetDateTo(v string) *QueryEdgeRequestHitRatioServiceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceRequest) SetDomain(v []*string) *QueryEdgeRequestHitRatioServiceRequest {
  s.Domain = v
  return s
}

type QueryEdgeRequestHitRatioServiceResponse struct {
  Result []*QueryEdgeRequestHitRatioServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeRequestHitRatioServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceResponse) SetResult(v []*QueryEdgeRequestHitRatioServiceResponseResult) *QueryEdgeRequestHitRatioServiceResponse {
  s.Result = v
  return s
}

type QueryEdgeRequestHitRatioServiceResponseResult struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间，格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total edge node hit ratio", "zh_CN":"总边缘节点命中率的平均值,2位小数"}
  TotalAvg *string `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  HitRatioDatas []*QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeRequestHitRatioServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceResponseResult) SetRealDate(v string) *QueryEdgeRequestHitRatioServiceResponseResult {
  s.RealDate = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceResponseResult) SetTotalAvg(v string) *QueryEdgeRequestHitRatioServiceResponseResult {
  s.TotalAvg = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceResponseResult) SetHitRatioDatas(v []*QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) *QueryEdgeRequestHitRatioServiceResponseResult {
  s.HitRatioDatas = v
  return s
}

type QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas struct     {
  // {"en":"timestamp", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"edge node hit ratio,keep 4 decimal places", "zh_CN":"边缘节点请求数命中率，保留4位小数"}
  EdgeHitRatio *string `json:"edgeHitRatio,omitempty" xml:"edgeHitRatio,omitempty" require:"true"`
}

func (s QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) SetTimestamp(v string) *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) SetEdgeHitRatio(v string) *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas {
  s.EdgeHitRatio = &v
  return s
}

type QueryEdgeRequestHitRatioServicePaths struct {
}

func (s QueryEdgeRequestHitRatioServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServicePaths) GoString() string {
  return s.String()
}

type QueryEdgeRequestHitRatioServiceParameters struct {
}

func (s QueryEdgeRequestHitRatioServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceParameters) GoString() string {
  return s.String()
}

type QueryEdgeRequestHitRatioServiceRequestHeader struct {
}

func (s QueryEdgeRequestHitRatioServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeRequestHitRatioServiceResponseHeader struct {
}

func (s QueryEdgeRequestHitRatioServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponseHeader) GoString() string {
  return s.String()
}




