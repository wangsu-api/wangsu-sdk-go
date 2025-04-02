package reporturl

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportUrlOriginServiceRequest struct {
  // {'en':'Start date:
  // The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  // Cannot be greater than the current time
  // The most recent six-month (183 days) data are available.', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年1月1日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing Time, 1 January 2019: 0 minutes 0 seconds);
  // The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // Date from, dateTo, neither passed, default query past 24 hours; If there is only one unsent, throw an exception
  // Maximum allowed query time interval: 31 days, i.e., the difference between Date from and dateTo cannot exceed 31 days (adjusted by contact technology support).', 'zh_CN':'结束时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年1月1日10点0分0秒）；
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天（可联系技术支持调整）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain name:
  // 
  // The maximum number of TLDs that can be delivered is 20 by default (Technical Support Adjustments can be contacted).
  // Auto-filter illegal domain name (pass illegal domain name, will be filtered, query result only returns the data of legal domain name)
  // All domain names are queried by default when this entry is not delivered, but an error occurs when the number of domain names in the account exceeds the upper limit', 'zh_CN':'域名：
  // 
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）。
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 3.未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s ReportUrlOriginServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUrlOriginServiceRequest) SetDateFrom(v string) *ReportUrlOriginServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUrlOriginServiceRequest) SetDateTo(v string) *ReportUrlOriginServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUrlOriginServiceRequest) SetDomain(v []*string) *ReportUrlOriginServiceRequest {
  s.Domain = v
  return s
}

type ReportUrlOriginServiceResponse struct {
  // {'en':'Request result status code', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Request Result Information', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'Detailed data on results of requests', 'zh_CN':'请求结果的详细数据'}
  Data []*ReportUrlOriginServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlOriginServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUrlOriginServiceResponse) SetCode(v string) *ReportUrlOriginServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportUrlOriginServiceResponse) SetMessage(v string) *ReportUrlOriginServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportUrlOriginServiceResponse) SetData(v []*ReportUrlOriginServiceResponseData) *ReportUrlOriginServiceResponse {
  s.Data = v
  return s
}

type ReportUrlOriginServiceResponseData struct     {
  // {'en':'Domain name', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  UrlDatas []*ReportUrlOriginServiceResponseDataUrlDatas `json:"urlDatas,omitempty" xml:"urlDatas,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlOriginServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportUrlOriginServiceResponseData) SetDomain(v string) *ReportUrlOriginServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportUrlOriginServiceResponseData) SetUrlDatas(v []*ReportUrlOriginServiceResponseDataUrlDatas) *ReportUrlOriginServiceResponseData {
  s.UrlDatas = v
  return s
}

type ReportUrlOriginServiceResponseDataUrlDatas struct     {
  // {'en':'URL', 'zh_CN':'URL'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'Number of requests', 'zh_CN':'请求数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportUrlOriginServiceResponseDataUrlDatas) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceResponseDataUrlDatas) GoString() string {
  return s.String()
}

func (s *ReportUrlOriginServiceResponseDataUrlDatas) SetUrl(v string) *ReportUrlOriginServiceResponseDataUrlDatas {
  s.Url = &v
  return s
}

func (s *ReportUrlOriginServiceResponseDataUrlDatas) SetValue(v string) *ReportUrlOriginServiceResponseDataUrlDatas {
  s.Value = &v
  return s
}

type ReportUrlOriginServicePaths struct {
}

func (s ReportUrlOriginServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServicePaths) GoString() string {
  return s.String()
}

type ReportUrlOriginServiceParameters struct {
}

func (s ReportUrlOriginServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceParameters) GoString() string {
  return s.String()
}

type ReportUrlOriginServiceRequestHeader struct {
}

func (s ReportUrlOriginServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUrlOriginServiceResponseHeader struct {
}

func (s ReportUrlOriginServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlOriginServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainRefererUrlServiceRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days.", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account and aggregate the returned data according to all domain names.
  // 2.Domain is uploaded: Up to 20 domains are supported.", "zh_CN":"域名:
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整)。
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整);
  // 3.自动过滤掉无效域名,(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration areaCode is not uploaded: Query all acceleration areas by default.", "zh_CN":"加速区域:
  // 
  // 未传递areaCode时,默认查询所有加速区域。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {"en":"Order by: 
  // 	1.The optional values are flow and request; 
  // 	2.If not passed, the default is request; 
  // 	3.When the value is flow, the query results are sorted in descending order of traffic, and when the value is request, they are sorted in descending order of the number of requests.
  // ", "zh_CN":"排序:
  // 
  // 1.request或flow;
  // 2.不传默认request;
  // 3.值为request时查询结果按照请求数降序排列;值为flow时查询结果按照流量值降序排列。"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
  // {"en":"Top quantity:
  // If no parameter is passed, the default is TOP 10, and the maximum is TOP 200", "zh_CN":"排行:
  // 不传默认top10,有传最大top200。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
}

func (s ReportDomainRefererUrlServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainRefererUrlServiceRequest) SetDateFrom(v string) *ReportDomainRefererUrlServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainRefererUrlServiceRequest) SetDateTo(v string) *ReportDomainRefererUrlServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainRefererUrlServiceRequest) SetDomain(v []*string) *ReportDomainRefererUrlServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportDomainRefererUrlServiceRequest) SetAreaCode(v []*string) *ReportDomainRefererUrlServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportDomainRefererUrlServiceRequest) SetOrderBy(v string) *ReportDomainRefererUrlServiceRequest {
  s.OrderBy = &v
  return s
}

func (s *ReportDomainRefererUrlServiceRequest) SetTop(v int) *ReportDomainRefererUrlServiceRequest {
  s.Top = &v
  return s
}

type ReportDomainRefererUrlServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*ReportDomainRefererUrlServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainRefererUrlServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainRefererUrlServiceResponse) SetCode(v string) *ReportDomainRefererUrlServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportDomainRefererUrlServiceResponse) SetMessage(v string) *ReportDomainRefererUrlServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportDomainRefererUrlServiceResponse) SetData(v []*ReportDomainRefererUrlServiceResponseData) *ReportDomainRefererUrlServiceResponse {
  s.Data = v
  return s
}

type ReportDomainRefererUrlServiceResponseData struct     {
  // {"en":"URL", "zh_CN":"URL"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Request", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"Flow", "zh_CN":"流量值,单位MB,保留2位小数"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s ReportDomainRefererUrlServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainRefererUrlServiceResponseData) SetUrl(v string) *ReportDomainRefererUrlServiceResponseData {
  s.Url = &v
  return s
}

func (s *ReportDomainRefererUrlServiceResponseData) SetRequest(v string) *ReportDomainRefererUrlServiceResponseData {
  s.Request = &v
  return s
}

func (s *ReportDomainRefererUrlServiceResponseData) SetFlow(v string) *ReportDomainRefererUrlServiceResponseData {
  s.Flow = &v
  return s
}

type ReportDomainRefererUrlServicePaths struct {
}

func (s ReportDomainRefererUrlServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServicePaths) GoString() string {
  return s.String()
}

type ReportDomainRefererUrlServiceParameters struct {
}

func (s ReportDomainRefererUrlServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainRefererUrlServiceRequestHeader struct {
}

func (s ReportDomainRefererUrlServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainRefererUrlServiceResponseHeader struct {
}

func (s ReportDomainRefererUrlServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererUrlServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryTopRankingUrlRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (Beijing Time 2 December 2016 10:0 min 0 seconds);
  // 2. Not Greater Than The Current Time;
  // 3. Data for the last six months (183 days) are available at most.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. Time format yyyy-MM-ddTHH:MM:ss+08:00
  // 2. The end time must be greater than the start time. if the end time is greater than the current time, take the current time.
  // 3. Date from, dateTo, neither passed, default query day; If there is only one unsent, throw an exception
  // 4. Maximum time interval allowed for queries: 31 days, i.e. the difference between Date from and dateTo cannot exceed 31 days. (Could contact technical support adjustment)", "zh_CN":"结束时间：
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询当天；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The default number of domain names is 100, which can be adjusted by contacting technical support. The maximum limit is 500;
  // 2. Automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).
  // 3. Did not deliver this entry, the default query account under all domain names, but when the number of domain names under the account number exceeds the upper limit, prompt an error", "zh_CN":"域名：
  // 1.可传递域名数量默认为100个，可联系技术支持调整，最大上限500；
  // 2.自动过滤掉无效域名（如传递非法域名，会被过滤，查询结果只返回有效域名的数据）。
  // 3.未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Ordering:
  // 1. The optional values are: totalRequest, totalFlow
  // 2. No default totalRequest", "zh_CN":"排序：
  // 1.可选值为：totalRequest、totalFlow
  // 2.不传则默认取值 totalRequest"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s QueryTopRankingUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTopRankingUrlRequest) GoString() string {
  return s.String()
}

func (s *QueryTopRankingUrlRequest) SetDateFrom(v string) *QueryTopRankingUrlRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTopRankingUrlRequest) SetDateTo(v string) *QueryTopRankingUrlRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTopRankingUrlRequest) SetDomain(v []*string) *QueryTopRankingUrlRequest {
  s.Domain = v
  return s
}

func (s *QueryTopRankingUrlRequest) SetOrderBy(v string) *QueryTopRankingUrlRequest {
  s.OrderBy = &v
  return s
}

type QueryTopRankingUrlResponse struct {
  // {"en":"Top ranking", "zh_CN":"top排名"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"url", "zh_CN":"url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Total flow: unit of measure MB, keeping 2 decimal places", "zh_CN":"总流量：计量单位MB，保留2位小数"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Total request", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Total hit request", "zh_CN":"请求数hit值"}
  HitRequest *string `json:"hitRequest,omitempty" xml:"hitRequest,omitempty" require:"true"`
  // {"en":"Total miss request", "zh_CN":"请求数miss值"}
  MissRequest *string `json:"missRequest,omitempty" xml:"missRequest,omitempty" require:"true"`
  // {"en":"Number of requests for status code 200", "zh_CN":"状态码200的请求数"}
  Sc200 *string `json:"sc200,omitempty" xml:"sc200,omitempty" require:"true"`
  // {"en":"Number of requests for status code 206", "zh_CN":"状态码206的请求数"}
  Sc206 *string `json:"sc206,omitempty" xml:"sc206,omitempty" require:"true"`
  // {"en":"Number of requests for status code 302", "zh_CN":"状态码302的请求数"}
  Sc302 *string `json:"sc302,omitempty" xml:"sc302,omitempty" require:"true"`
  // {"en":"Number of requests for status code 304", "zh_CN":"状态码304的请求数"}
  Sc304 *string `json:"sc304,omitempty" xml:"sc304,omitempty" require:"true"`
  // {"en":"Number of requests for status code 404", "zh_CN":"状态码404的请求数"}
  Sc404 *string `json:"sc404,omitempty" xml:"sc404,omitempty" require:"true"`
}

func (s QueryTopRankingUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTopRankingUrlResponse) GoString() string {
  return s.String()
}

func (s *QueryTopRankingUrlResponse) SetTop(v string) *QueryTopRankingUrlResponse {
  s.Top = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetUrl(v string) *QueryTopRankingUrlResponse {
  s.Url = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetTotalFlow(v string) *QueryTopRankingUrlResponse {
  s.TotalFlow = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetTotalRequest(v string) *QueryTopRankingUrlResponse {
  s.TotalRequest = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetHitRequest(v string) *QueryTopRankingUrlResponse {
  s.HitRequest = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetMissRequest(v string) *QueryTopRankingUrlResponse {
  s.MissRequest = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetSc200(v string) *QueryTopRankingUrlResponse {
  s.Sc200 = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetSc206(v string) *QueryTopRankingUrlResponse {
  s.Sc206 = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetSc302(v string) *QueryTopRankingUrlResponse {
  s.Sc302 = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetSc304(v string) *QueryTopRankingUrlResponse {
  s.Sc304 = &v
  return s
}

func (s *QueryTopRankingUrlResponse) SetSc404(v string) *QueryTopRankingUrlResponse {
  s.Sc404 = &v
  return s
}

type QueryTopRankingUrlPaths struct {
}

func (s QueryTopRankingUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTopRankingUrlPaths) GoString() string {
  return s.String()
}

type QueryTopRankingUrlParameters struct {
}

func (s QueryTopRankingUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTopRankingUrlParameters) GoString() string {
  return s.String()
}

type QueryTopRankingUrlRequestHeader struct {
}

func (s QueryTopRankingUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTopRankingUrlRequestHeader) GoString() string {
  return s.String()
}

type QueryTopRankingUrlResponseHeader struct {
}

func (s QueryTopRankingUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTopRankingUrlResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainRefererWebsiteServiceRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2024-01-23T10:00 + 08:00 (10:00:00 Beijing time on January 23, 2024);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2024-01-23T10:00:00+08:00（为北京时间2024年01月23日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. The end time is greater than the start time.
  // 
  // 3. If the end time is greater than the current time, the current time is taken.
  // 
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days. ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：31天，即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 1.The maximum number of deliverable domain names is 200 by default 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names) 3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为200个
  // 
  // 2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 
  // 3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Order by: 1.The optional values are flow and request; 2.the default value is request; 3.When the value is flow, the query results are sorted in descending order of traffic, and when the value is request, they are sorted in descending order of the number of requests.", "zh_CN":"排序： 1.可选值为：request、flow 2.不传则默认取值 request
  // 
  // 3.值为request时查询结果按照请求数降序排列;值为flow时查询结果按照流量值降序排列。"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
  // {"en":"op quantity: the default is TOP 10, and the maximum is TOP 500
  // 
  // ", "zh_CN":"排行: 不传默认top10,有传最大top500。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
}

func (s ReportDomainRefererWebsiteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainRefererWebsiteServiceRequest) SetDateFrom(v string) *ReportDomainRefererWebsiteServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceRequest) SetDateTo(v string) *ReportDomainRefererWebsiteServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceRequest) SetDomain(v []*string) *ReportDomainRefererWebsiteServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportDomainRefererWebsiteServiceRequest) SetOrderBy(v string) *ReportDomainRefererWebsiteServiceRequest {
  s.OrderBy = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceRequest) SetTop(v int) *ReportDomainRefererWebsiteServiceRequest {
  s.Top = &v
  return s
}

type ReportDomainRefererWebsiteServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*ReportDomainRefererWebsiteServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainRefererWebsiteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainRefererWebsiteServiceResponse) SetCode(v string) *ReportDomainRefererWebsiteServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceResponse) SetMessage(v string) *ReportDomainRefererWebsiteServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceResponse) SetData(v []*ReportDomainRefererWebsiteServiceResponseData) *ReportDomainRefererWebsiteServiceResponse {
  s.Data = v
  return s
}

type ReportDomainRefererWebsiteServiceResponseData struct     {
  // {"en":"Website", "zh_CN":"网站"}
  Website *string `json:"website,omitempty" xml:"website,omitempty" require:"true"`
  // {"en":"Requests", "zh_CN":"对应请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"Traffic data, in MB, rounded to two decimal places", "zh_CN":"对应流量数据，单位：MB,保留两位小数"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s ReportDomainRefererWebsiteServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainRefererWebsiteServiceResponseData) SetWebsite(v string) *ReportDomainRefererWebsiteServiceResponseData {
  s.Website = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceResponseData) SetRequest(v string) *ReportDomainRefererWebsiteServiceResponseData {
  s.Request = &v
  return s
}

func (s *ReportDomainRefererWebsiteServiceResponseData) SetFlow(v string) *ReportDomainRefererWebsiteServiceResponseData {
  s.Flow = &v
  return s
}

type ReportDomainRefererWebsiteServicePaths struct {
}

func (s ReportDomainRefererWebsiteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServicePaths) GoString() string {
  return s.String()
}

type ReportDomainRefererWebsiteServiceParameters struct {
}

func (s ReportDomainRefererWebsiteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainRefererWebsiteServiceRequestHeader struct {
}

func (s ReportDomainRefererWebsiteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainRefererWebsiteServiceResponseHeader struct {
}

func (s ReportDomainRefererWebsiteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainRefererWebsiteServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUrlCustomTopDailyServiceRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 7 days will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 31days, means the period between dateFrom to dateTo should not exceed 31days.", "zh_CN":"结束时间:
  // 1.时间格式2019-01-02T10:00:00+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的7天;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1. The maximum number of domain is 100 by default (the upper limit is 500 according to technical support adjustment);
  // 2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).
  // 3. If the input parameter is not inputed, all domains under the account will be queried by default, but an error will be prompted when the number of domains under the account exceeds the upper limit", "zh_CN":"域名:
  // 可传递域名数量上限默认为100个(可联系技术支持调整)。
  // 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"highest number:
  // Default TOP 100;
  // Maximum TOP 500.", "zh_CN":"TOP个数:
  // 不传默认TOP 100
  // 最大TOP 500"}
  Top *string `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"sort by:
  // 1. Optional values are: request, flow
  // 2. Default value request", "zh_CN":"排序:
  // 1. 可选值为:request, flow
  // 2. 不传默认request"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ReportUrlCustomTopDailyServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUrlCustomTopDailyServiceRequest) SetDateFrom(v string) *ReportUrlCustomTopDailyServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceRequest) SetDateTo(v string) *ReportUrlCustomTopDailyServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceRequest) SetDomain(v []*string) *ReportUrlCustomTopDailyServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUrlCustomTopDailyServiceRequest) SetTop(v string) *ReportUrlCustomTopDailyServiceRequest {
  s.Top = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceRequest) SetOrderBy(v string) *ReportUrlCustomTopDailyServiceRequest {
  s.OrderBy = &v
  return s
}

type ReportUrlCustomTopDailyServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*ReportUrlCustomTopDailyServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlCustomTopDailyServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUrlCustomTopDailyServiceResponse) SetCode(v string) *ReportUrlCustomTopDailyServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceResponse) SetMessage(v string) *ReportUrlCustomTopDailyServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceResponse) SetData(v []*ReportUrlCustomTopDailyServiceResponseData) *ReportUrlCustomTopDailyServiceResponse {
  s.Data = v
  return s
}

type ReportUrlCustomTopDailyServiceResponseData struct     {
  // {"en":"top", "zh_CN":"top排名"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"url", "zh_CN":"url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Total traffic: unit MB, with 2 decimal places", "zh_CN":"总流量:计量单位MB,保留2位小数"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Total requests", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
}

func (s ReportUrlCustomTopDailyServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportUrlCustomTopDailyServiceResponseData) SetTop(v string) *ReportUrlCustomTopDailyServiceResponseData {
  s.Top = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceResponseData) SetUrl(v string) *ReportUrlCustomTopDailyServiceResponseData {
  s.Url = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceResponseData) SetTotalFlow(v string) *ReportUrlCustomTopDailyServiceResponseData {
  s.TotalFlow = &v
  return s
}

func (s *ReportUrlCustomTopDailyServiceResponseData) SetTotalRequest(v string) *ReportUrlCustomTopDailyServiceResponseData {
  s.TotalRequest = &v
  return s
}

type ReportUrlCustomTopDailyServicePaths struct {
}

func (s ReportUrlCustomTopDailyServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServicePaths) GoString() string {
  return s.String()
}

type ReportUrlCustomTopDailyServiceParameters struct {
}

func (s ReportUrlCustomTopDailyServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServiceParameters) GoString() string {
  return s.String()
}

type ReportUrlCustomTopDailyServiceRequestHeader struct {
}

func (s ReportUrlCustomTopDailyServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUrlCustomTopDailyServiceResponseHeader struct {
}

func (s ReportUrlCustomTopDailyServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlCustomTopDailyServiceResponseHeader) GoString() string {
  return s.String()
}




