package logdownload

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCatalogueTrafficAndBroadcastCountDownloadRequest struct {
  // {"en":"Start Time:\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time.\n\n2.The start time cannot be later than the current time.\n\n3.The time span between dateFrom and dateTo cannot exceed 90 days.\n\n4.Data can only be queried for the past 1 year (366 days).","zh_CN":"开始时间: \n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n\n2.不能大于当前时间\n\n3.dateFrom和dateTo相差不能超过90天 \n\n4.只能查询最近1年（366天）数据."}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents UTC+08:00, and -05:00 represents UTC-05:00. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM.\n\n2. The end time must be later than the start time.\n\n3. If dateFrom and dateTo parameters are not passed, the default query is the past day.","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n2.结束时间需大于开始时间\n\n3.如果不传dateFrom和dateTo参数，则默认查询过去一天的数据"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:\n\n1.If no domain is provided: Query all domain names of the account (an error will be returned if more than 200 domains are provided, you can contact technical support for adjustment);\n\n2.If domains are provided: Up to 200 domains are supported (you can contact technical support for adjustment).","zh_CN":"域名:\n\n1.未传递domain时:查询账号下所有全部域名(域名超过200个则报错,可联系技术支持调整);\n\n2.有传递domain时:域名最多支持传200个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadRequest) GoString() string {
  return s.String()
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadRequest) SetDateFrom(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadRequest) SetDateTo(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadRequest {
  s.DateTo = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadRequest) SetDomain(v []*string) *QueryCatalogueTrafficAndBroadcastCountDownloadRequest {
  s.Domain = v
  return s
}

type QueryCatalogueTrafficAndBroadcastCountDownloadRequestHeader struct {
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadRequestHeader) GoString() string {
  return s.String()
}

type QueryCatalogueTrafficAndBroadcastCountDownloadPaths struct {
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadPaths) GoString() string {
  return s.String()
}

type QueryCatalogueTrafficAndBroadcastCountDownloadParameters struct {
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadParameters) GoString() string {
  return s.String()
}

type QueryCatalogueTrafficAndBroadcastCountDownloadResponse struct {
  // {"en":"Response code","zh_CN":"业务处理返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returned data","zh_CN":"接口返回的具体数据内容"}
  Data []*QueryCatalogueTrafficAndBroadcastCountDownloadResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponse) GoString() string {
  return s.String()
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponse) SetCode(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponse {
  s.Code = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponse) SetMessage(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponse {
  s.Message = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponse) SetData(v []*QueryCatalogueTrafficAndBroadcastCountDownloadResponseData) *QueryCatalogueTrafficAndBroadcastCountDownloadResponse {
  s.Data = v
  return s
}

type QueryCatalogueTrafficAndBroadcastCountDownloadResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"List of file data","zh_CN":"文件数据列表”"}
  FileData []*QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData `json:"fileData,omitempty" xml:"fileData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponseData) GoString() string {
  return s.String()
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseData) SetDomain(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseData {
  s.Domain = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseData) SetFileData(v []*QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseData {
  s.FileData = v
  return s
}

type QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData struct     {
  // {"en":"The start time of log file, format is yyyy-MM-dd-HHmm.","zh_CN":"日志文件的开始时间，格式为yyyy-MM-dd-HHmm"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"The end time of log file, format is yyyy-MM-dd-HHmm.","zh_CN":"日志文件的结束时间，格式为yyyy-MM-dd-HHmm"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Download address of log file.","zh_CN":"日志文件下载地址"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"The expiration time of the log file download address.","zh_CN":"日志文件下载地址过期时间"}
  UrlExpireTime *string `json:"urlExpireTime,omitempty" xml:"urlExpireTime,omitempty" require:"true"`
  // {"en":"File name.","zh_CN":"文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"The size of the log file in bytes.","zh_CN":"日志文件大小，单位：Byte"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) GoString() string {
  return s.String()
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) SetDateFrom(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData {
  s.DateFrom = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) SetDateTo(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData {
  s.DateTo = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) SetLogUrl(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData {
  s.LogUrl = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) SetUrlExpireTime(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData {
  s.UrlExpireTime = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) SetFileName(v string) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData {
  s.FileName = &v
  return s
}

func (s *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData) SetFileSize(v int64) *QueryCatalogueTrafficAndBroadcastCountDownloadResponseDataFileData {
  s.FileSize = &v
  return s
}

type QueryCatalogueTrafficAndBroadcastCountDownloadResponseHeader struct {
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCatalogueTrafficAndBroadcastCountDownloadResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainLogDownloadAddressRequest struct {
  // {"en":"Domain list:\n1.The default upper limit of the number of domain names is 20 (you can contact technical support for adjustment), The maximum recommended cap is 500 .","zh_CN":"域名列表:\n1.域名数量默认上限为20个(可联系技术支持调整), 最大上限为500."}
  DomainList *QueryDomainLogDownloadAddressRequestDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" type:"Struct"`
  // {"en":"Start time:\n1.The format is yyyy-MM-ddTHH:mm:ss+00:00 ;\n2.Must be smaller than the current time and dateTo ;\n3.Period between dataFrom and dateTo cannot be longer than 31 days .","zh_CN":"开始时间:\n1.格式为yyyy-MM-ddTHH:mm:ss+00:00 ;\n2.必须小于当前时间和dateTo ;\n3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整) ;\n4.只能查询最近2年内数据.(实际可查询的日志范围,取决于域名配置的日志保留天数) ."}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time, format is yyyy-MM-ddTHH:mm:ss+00:00(The actual range of logs that can be queried depends on the number of days of log retention configured by the domain name).","zh_CN":"结束时间,格式为yyyy-MM-ddTHH:mm:ss+00:00(实际可查询的日志范围,取决于域名配置的日志保留天数)."}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"defaultValue":"cdn,bot","en":"Log type, optional values:cdn,bot,ddos,waap,waf;Multiple are separated by English commas. If they are not transmitted, the data will be queried by logtype = cdn,bot by default.","zh_CN":"日志类型,可选值: cdn,bot,ddos,waap,waf; 多个通过英文逗号分隔,若未传则默认按logtype=cdn,bot查询数据","exampleValue":"cdn,bot,ddos"}
  LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"areaCode","zh_CN":"加速区域"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
}

func (s QueryDomainLogDownloadAddressRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressRequest) SetDomainList(v *QueryDomainLogDownloadAddressRequestDomainList) *QueryDomainLogDownloadAddressRequest {
  s.DomainList = v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetDatefrom(v string) *QueryDomainLogDownloadAddressRequest {
  s.Datefrom = &v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetDateto(v string) *QueryDomainLogDownloadAddressRequest {
  s.Dateto = &v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetLogType(v string) *QueryDomainLogDownloadAddressRequest {
  s.LogType = &v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetAreaCode(v string) *QueryDomainLogDownloadAddressRequest {
  s.AreaCode = &v
  return s
}

type QueryDomainLogDownloadAddressRequestDomainList struct {
  // {"en":"Domain","zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressRequestDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequestDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressRequestDomainList) SetDomainName(v []*string) *QueryDomainLogDownloadAddressRequestDomainList {
  s.DomainName = v
  return s
}

type QueryDomainLogDownloadAddressRequestHeader struct {
}

func (s QueryDomainLogDownloadAddressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainLogDownloadAddressPaths struct {
}

func (s QueryDomainLogDownloadAddressPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressPaths) GoString() string {
  return s.String()
}

type QueryDomainLogDownloadAddressParameters struct {
}

func (s QueryDomainLogDownloadAddressParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressParameters) GoString() string {
  return s.String()
}

type QueryDomainLogDownloadAddressResponse struct {
  // {"en":"","zh_CN":""}
  Logs []*QueryDomainLogDownloadAddressResponseLogs `json:"logs,omitempty" xml:"logs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponse) SetLogs(v []*QueryDomainLogDownloadAddressResponseLogs) *QueryDomainLogDownloadAddressResponse {
  s.Logs = v
  return s
}

type QueryDomainLogDownloadAddressResponseLogs struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"areaCode","zh_CN":"加速区域"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Files []*QueryDomainLogDownloadAddressResponseLogsFiles `json:"files,omitempty" xml:"files,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressResponseLogs) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponseLogs) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponseLogs) SetDomain(v string) *QueryDomainLogDownloadAddressResponseLogs {
  s.Domain = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogs) SetAreaCode(v string) *QueryDomainLogDownloadAddressResponseLogs {
  s.AreaCode = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogs) SetFiles(v []*QueryDomainLogDownloadAddressResponseLogsFiles) *QueryDomainLogDownloadAddressResponseLogs {
  s.Files = v
  return s
}

type QueryDomainLogDownloadAddressResponseLogsFiles struct     {
  // {"en":"The start time of log file, format is yyyy-MM-dd-HHmm","zh_CN":"日志文件的开始时间,格式为yyyy-MM-dd-HHmm"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"The end time of log file, format is yyyy-MM-dd-HHmm","zh_CN":"日志文件的结束时间,格式为yyyy-MM-dd-HHmm"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Download address of log file","zh_CN":"日志文件下载地址"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Size of log file","zh_CN":"日志文件大小"}
  FileSize *int `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s QueryDomainLogDownloadAddressResponseLogsFiles) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponseLogsFiles) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetDateFrom(v string) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetDateTo(v string) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.DateTo = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetLogUrl(v string) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.LogUrl = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetFileSize(v int) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.FileSize = &v
  return s
}

type QueryDomainLogDownloadAddressResponseHeader struct {
}

func (s QueryDomainLogDownloadAddressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponseHeader) GoString() string {
  return s.String()
}




type QueryTranscodingDurationLogDownloadAddressRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days. ", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 100 domains will error,you can contact technical support for adjustment);
  // 
  // 2.Domain is uploaded: Up to 100 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryTranscodingDurationLogDownloadAddressRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressRequest) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressRequest) SetDateFrom(v string) *QueryTranscodingDurationLogDownloadAddressRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressRequest) SetDateTo(v string) *QueryTranscodingDurationLogDownloadAddressRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressRequest) SetDomain(v []*string) *QueryTranscodingDurationLogDownloadAddressRequest {
  s.Domain = v
  return s
}

type QueryTranscodingDurationLogDownloadAddressResponse struct {
  Result []*QueryTranscodingDurationLogDownloadAddressResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTranscodingDurationLogDownloadAddressResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressResponse) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressResponse) SetResult(v []*QueryTranscodingDurationLogDownloadAddressResponseResult) *QueryTranscodingDurationLogDownloadAddressResponse {
  s.Result = v
  return s
}

type QueryTranscodingDurationLogDownloadAddressResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  FileData []*QueryTranscodingDurationLogDownloadAddressResponseResultFileData `json:"fileData,omitempty" xml:"fileData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTranscodingDurationLogDownloadAddressResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResult) SetDomain(v string) *QueryTranscodingDurationLogDownloadAddressResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResult) SetFileData(v []*QueryTranscodingDurationLogDownloadAddressResponseResultFileData) *QueryTranscodingDurationLogDownloadAddressResponseResult {
  s.FileData = v
  return s
}

type QueryTranscodingDurationLogDownloadAddressResponseResultFileData struct     {
  // {"en":"The start time of log file, format is yyyy-MM-dd-HHmm", "zh_CN":"日志文件的开始时间,格式为yyyy-MM-dd-HHmm"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"The end time of log file, format is yyyy-MM-dd-HHmm", "zh_CN":"日志文件的结束时间,格式为yyyy-MM-dd-HHmm"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Download address of log file", "zh_CN":"日志文件下载地址"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Expire time", "zh_CN":"日志文件下载地址过期时间"}
  UrlExpireTime *string `json:"urlExpireTime,omitempty" xml:"urlExpireTime,omitempty" require:"true"`
  // {"en":"File name", "zh_CN":"文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"Size of log file", "zh_CN":"日志文件大小,单位:Byte"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s QueryTranscodingDurationLogDownloadAddressResponseResultFileData) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressResponseResultFileData) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetDateFrom(v string) *QueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.DateFrom = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetDateTo(v string) *QueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.DateTo = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetLogUrl(v string) *QueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.LogUrl = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetUrlExpireTime(v string) *QueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.UrlExpireTime = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetFileName(v string) *QueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.FileName = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetFileSize(v int64) *QueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.FileSize = &v
  return s
}

type QueryTranscodingDurationLogDownloadAddressPaths struct {
}

func (s QueryTranscodingDurationLogDownloadAddressPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressPaths) GoString() string {
  return s.String()
}

type QueryTranscodingDurationLogDownloadAddressParameters struct {
}

func (s QueryTranscodingDurationLogDownloadAddressParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressParameters) GoString() string {
  return s.String()
}

type QueryTranscodingDurationLogDownloadAddressRequestHeader struct {
}

func (s QueryTranscodingDurationLogDownloadAddressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressRequestHeader) GoString() string {
  return s.String()
}

type QueryTranscodingDurationLogDownloadAddressResponseHeader struct {
}

func (s QueryTranscodingDurationLogDownloadAddressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressResponseHeader) GoString() string {
  return s.String()
}




