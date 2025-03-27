package logdownload

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryDomainLogDownloadAddressRequest struct {
  // {"en":"Domain list:
  //     1.All domains will be queried if this field is not specified;
  //     2.The default upper limit of the number of domain names is 20 (you can contact technical support for adjustment), The maximum recommended cap is 500 .", "zh_CN":"域名列表:
  //     1.不传递则查询全部域名;
  //     2.域名数量默认上限为20个(可联系技术支持调整), 最大上限为500."}
  QueryDomainLogDownloadAddressDomainList *QueryDomainLogDownloadAddressDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty"`
  // {"en":"Start time:
  //     1.The format is yyyy-MM-ddTHH:mm:ss+08:00 ;
  //     2.Must be smaller than the current time and dateTo ;
  //     3.Period between dataFrom and dateTo cannot be longer than 31 days .", "zh_CN":"开始时间:
  //     1.格式为yyyy-MM-ddTHH:mm:ss+08:00 ;
  //     2.必须小于当前时间和dateTo ;
  //     3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整) ;
  //     4.只能查询最近2年内数据.(实际可查询的日志范围,取决于域名配置的日志保留天数) ."}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time, format is yyyy-MM-ddTHH:mm:ss+08:00(The actual range of logs that can be queried depends on the number of days of log retention configured by the domain name).", "zh_CN":"结束时间,格式为yyyy-MM-ddTHH:mm:ss+08:00(实际可查询的日志范围,取决于域名配置的日志保留天数)."}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Log type, optional values:cdn,bot,ddos;Multiple are separated by English commas. If they are not transmitted, the data will be queried by logtype = cdn,bot by default.", "zh_CN":"日志类型,可选值: cdn,bot,ddos ;
  // 多个通过英文逗号分隔,若未传则默认按logtype=cdn,bot查询数据"}
  LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
  // {"en":"areaCode", "zh_CN":"加速区域"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
}

func (s QueryDomainLogDownloadAddressRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressRequest) SetDomainList(v *QueryDomainLogDownloadAddressDomainList) *QueryDomainLogDownloadAddressRequest {
  s.QueryDomainLogDownloadAddressDomainList = v
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

type QueryDomainLogDownloadAddressDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressDomainList) SetDomainName(v []*string) *QueryDomainLogDownloadAddressDomainList {
  s.DomainName = v
  return s
}

type QueryDomainLogDownloadAddressResponse struct {
  Logs []*QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs `json:"logs,omitempty" xml:"logs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponse) SetLogs(v []*QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs) *QueryDomainLogDownloadAddressResponse {
  s.Logs = v
  return s
}

type QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"areaCode", "zh_CN":"加速区域"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
  Files []*QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles `json:"files,omitempty" xml:"files,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs) SetDomain(v string) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs {
  s.Domain = &v
  return s
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs) SetAreaCode(v string) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs {
  s.AreaCode = &v
  return s
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs) SetFiles(v []*QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogs {
  s.Files = v
  return s
}

type QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles struct     {
  // {"en":"The start time of log file, format is yyyy-MM-dd-HHmm", "zh_CN":"日志文件的开始时间,格式为yyyy-MM-dd-HHmm"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"The end time of log file, format is yyyy-MM-dd-HHmm", "zh_CN":"日志文件的结束时间,格式为yyyy-MM-dd-HHmm"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Download address of log file", "zh_CN":"日志文件下载地址"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Size of log file", "zh_CN":"日志文件大小"}
  FileSize *int32 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) SetDateFrom(v string) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) SetDateTo(v string) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles {
  s.DateTo = &v
  return s
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) SetLogUrl(v string) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles {
  s.LogUrl = &v
  return s
}

func (s *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles) SetFileSize(v int32) *QueryDomainLogDownloadAddressQueryDomainLogDownloadAddressResponseLogsFiles {
  s.FileSize = &v
  return s
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

type QueryDomainLogDownloadAddressRequestHeader struct {
}

func (s QueryDomainLogDownloadAddressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequestHeader) GoString() string {
  return s.String()
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
  Result []*QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTranscodingDurationLogDownloadAddressResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressResponse) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressResponse) SetResult(v []*QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult) *QueryTranscodingDurationLogDownloadAddressResponse {
  s.Result = v
  return s
}

type QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  FileData []*QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData `json:"fileData,omitempty" xml:"fileData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult) SetDomain(v string) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult) SetFileData(v []*QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResult {
  s.FileData = v
  return s
}

type QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData struct     {
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

func (s QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) String() string {
  return tea.Prettify(s)
}

func (s QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) GoString() string {
  return s.String()
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetDateFrom(v string) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.DateFrom = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetDateTo(v string) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.DateTo = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetLogUrl(v string) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.LogUrl = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetUrlExpireTime(v string) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.UrlExpireTime = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetFileName(v string) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData {
  s.FileName = &v
  return s
}

func (s *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData) SetFileSize(v int64) *QueryTranscodingDurationLogDownloadAddressQueryTranscodingDurationLogDownloadAddressResponseResultFileData {
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




