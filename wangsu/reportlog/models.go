package reportlog

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CheckIsWhiteIpRequest struct {
  // {"en":"entername", "zh_CN":"客户英文名"}
  White_name *string `json:"white_name,omitempty" xml:"white_name,omitempty" require:"true"`
  // {"en":"ip_white type:
  // comm_white
  // relay_white", "zh_CN":"白名单类型 
  // 普通白名单：comm_white
  // 中转白名单：relay_white"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"ip white hash verify value
  // format hash1:hash2", "zh_CN":"白名单hash值
  // hash1:hash2
  // 如果ipv6独立
  // 则：hash1:hash2:y"}
  White_hash *string `json:"white_hash,omitempty" xml:"white_hash,omitempty" require:"true"`
  // {"en":"check ips,split by ,", "zh_CN":"检查的ip，多个,隔开"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s CheckIsWhiteIpRequest) String() string {
  return tea.Prettify(s)
}

func (s CheckIsWhiteIpRequest) GoString() string {
  return s.String()
}

func (s *CheckIsWhiteIpRequest) SetWhite_name(v string) *CheckIsWhiteIpRequest {
  s.White_name = &v
  return s
}

func (s *CheckIsWhiteIpRequest) SetType(v string) *CheckIsWhiteIpRequest {
  s.Type = &v
  return s
}

func (s *CheckIsWhiteIpRequest) SetWhite_hash(v string) *CheckIsWhiteIpRequest {
  s.White_hash = &v
  return s
}

func (s *CheckIsWhiteIpRequest) SetIp(v string) *CheckIsWhiteIpRequest {
  s.Ip = &v
  return s
}

type CheckIsWhiteIpResponse struct {
  // {"en":"Return status code
  // success represents normal
  // fail  represents abnormality", "zh_CN":"返回状态码 
  // success 代表正常
  // fail 代表异常"}
  Ret_code *string `json:"ret_code,omitempty" xml:"ret_code,omitempty" require:"true"`
  // {"en":"return is every ip in use with yes or no", "zh_CN":"返回每个ip是否是在用白名单IP，格式为yes或者no"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s CheckIsWhiteIpResponse) String() string {
  return tea.Prettify(s)
}

func (s CheckIsWhiteIpResponse) GoString() string {
  return s.String()
}

func (s *CheckIsWhiteIpResponse) SetRet_code(v string) *CheckIsWhiteIpResponse {
  s.Ret_code = &v
  return s
}

func (s *CheckIsWhiteIpResponse) SetData(v []*string) *CheckIsWhiteIpResponse {
  s.Data = v
  return s
}

type CheckIsWhiteIpPaths struct {
}

func (s CheckIsWhiteIpPaths) String() string {
  return tea.Prettify(s)
}

func (s CheckIsWhiteIpPaths) GoString() string {
  return s.String()
}

type CheckIsWhiteIpParameters struct {
}

func (s CheckIsWhiteIpParameters) String() string {
  return tea.Prettify(s)
}

func (s CheckIsWhiteIpParameters) GoString() string {
  return s.String()
}

type CheckIsWhiteIpRequestHeader struct {
}

func (s CheckIsWhiteIpRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIsWhiteIpRequestHeader) GoString() string {
  return s.String()
}

type CheckIsWhiteIpResponseHeader struct {
}

func (s CheckIsWhiteIpResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIsWhiteIpResponseHeader) GoString() string {
  return s.String()
}




type GetBotAttackIncidentLogDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'.", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
  // {"en":"Attack Type,separate by ';'.", "zh_CN":"攻击类型。多个以;隔开。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty"`
  // {"en":"Action. 
  //  1:Interception 
  //  2:Log 
  //  7:Flag 
  //  8:Captcha", "zh_CN":"处理动作。 
  //  1：拦截 
  //  2：监控 
  //  7：攻击标记 
  //  8：验证码"}
  Act *string `json:"act,omitempty" xml:"act,omitempty"`
  // {"en":"IP location.", "zh_CN":"IP地理位置。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty"`
  // {"en":"Client IP.", "zh_CN":"客户端IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"URI.", "zh_CN":"URI。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en":"Referer.", "zh_CN":"Referer。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty"`
  // {"en":"Status code.", "zh_CN":"状态码。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  // {"en":"User-Agent.", "zh_CN":"User-Agent。"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
  // {"en":"Event ID.", "zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
  // {"en":"The number of entries displayed per page.Maximum limit 10,000.", "zh_CN":"每页显示的条目数。最大限制10,000。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Current page number.", "zh_CN":"当前页码。"}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
  // {"en":"Client ID.", "zh_CN":"客户端ID。"}
  ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
  // {"en":"Browser Fingerprint.", "zh_CN":"浏览器指纹。"}
  BrowserFp *string `json:"browserFp,omitempty" xml:"browserFp,omitempty"`
  // {"en":"Rule name.", "zh_CN":"规则名。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Bot rule name.", "zh_CN":"Bot规则名。"}
  BotRuleName *string `json:"botRuleName,omitempty" xml:"botRuleName,omitempty"`
  // {"en":"Query criteria matching method of 'Client IP'.Default value: 3. 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端IP'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  IpCondition *int `json:"ipCondition,omitempty" xml:"ipCondition,omitempty"`
  // {"en":"Query criteria matching method of 'URI'. Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端URL'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  UrlCondition *int `json:"urlCondition,omitempty" xml:"urlCondition,omitempty"`
  // {"en":"Query criteria matching method of 'Referer'. Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端Rerfer'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  RefererCondition *int `json:"refererCondition,omitempty" xml:"refererCondition,omitempty"`
  // {"en":"Query criteria matching method of 'Response code'.Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端状态码'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  StatusCodeConditon *int `json:"statusCodeConditon,omitempty" xml:"statusCodeConditon,omitempty"`
  // {"en":"Query criteria matching method of 'User-Agent'.Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include. ", "zh_CN":"'客户端UA'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  UserAgentCondition *int `json:"userAgentCondition,omitempty" xml:"userAgentCondition,omitempty"`
}

func (s GetBotAttackIncidentLogDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataRequest) SetDomain(v string) *GetBotAttackIncidentLogDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetStartTime(v string) *GetBotAttackIncidentLogDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetEndTime(v string) *GetBotAttackIncidentLogDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetTimeZone(v int) *GetBotAttackIncidentLogDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetLang(v string) *GetBotAttackIncidentLogDataRequest {
  s.Lang = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetAttackType(v string) *GetBotAttackIncidentLogDataRequest {
  s.AttackType = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetAct(v string) *GetBotAttackIncidentLogDataRequest {
  s.Act = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetLocation(v string) *GetBotAttackIncidentLogDataRequest {
  s.Location = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetIp(v string) *GetBotAttackIncidentLogDataRequest {
  s.Ip = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUrl(v string) *GetBotAttackIncidentLogDataRequest {
  s.Url = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetReferer(v string) *GetBotAttackIncidentLogDataRequest {
  s.Referer = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetStatusCode(v string) *GetBotAttackIncidentLogDataRequest {
  s.StatusCode = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUserAgent(v string) *GetBotAttackIncidentLogDataRequest {
  s.UserAgent = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUuid(v string) *GetBotAttackIncidentLogDataRequest {
  s.Uuid = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetPageSize(v int) *GetBotAttackIncidentLogDataRequest {
  s.PageSize = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetCurrentPage(v int) *GetBotAttackIncidentLogDataRequest {
  s.CurrentPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetClientId(v string) *GetBotAttackIncidentLogDataRequest {
  s.ClientId = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetBrowserFp(v string) *GetBotAttackIncidentLogDataRequest {
  s.BrowserFp = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetRuleName(v string) *GetBotAttackIncidentLogDataRequest {
  s.RuleName = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetBotRuleName(v string) *GetBotAttackIncidentLogDataRequest {
  s.BotRuleName = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetIpCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.IpCondition = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUrlCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.UrlCondition = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetRefererCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.RefererCondition = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetStatusCodeConditon(v int) *GetBotAttackIncidentLogDataRequest {
  s.StatusCodeConditon = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUserAgentCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.UserAgentCondition = &v
  return s
}

type GetBotAttackIncidentLogDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据"}
  Data *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetBotAttackIncidentLogDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataResponse) SetCode(v string) *GetBotAttackIncidentLogDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponse) SetMessage(v string) *GetBotAttackIncidentLogDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponse) SetData(v *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) *GetBotAttackIncidentLogDataResponse {
  s.Data = v
  return s
}

type GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData struct {
  // {"en":"Rule name.", "zh_CN":"当前页码。"}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty" require:"true"`
  // {"en":"Current page number.", "zh_CN":"首页页码。"}
  FirstPage *int `json:"firstPage,omitempty" xml:"firstPage,omitempty" require:"true"`
  // {"en":"last page number.", "zh_CN":"末页页码。"}
  LastPage *int `json:"lastPage,omitempty" xml:"lastPage,omitempty" require:"true"`
  // {"en":"The number of entries displayed per page.Maximum limit 10,000.", "zh_CN":"每页显示的条目数。最大限制10,000。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total entries.Maximum limit 1,000,000.", "zh_CN":"总条目数。最大限制1,000,000。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"Total page count.", "zh_CN":"总页数。"}
  TotalPageCount *int `json:"totalPageCount,omitempty" xml:"totalPageCount,omitempty" require:"true"`
  // {"en":"Data List", "zh_CN":"数据列表"}
  List []*GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetCurrentPage(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.CurrentPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetFirstPage(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.FirstPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetLastPage(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.LastPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetPageSize(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.PageSize = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetTotalCount(v int64) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.TotalCount = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetTotalPageCount(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.TotalPageCount = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData) SetList(v []*GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseData {
  s.List = v
  return s
}

type GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList struct     {
  // {"en":"Referer.", "zh_CN":"Referer。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {"en":"Browser Fingerprint.", "zh_CN":"浏览器指纹。"}
  Browser_fp *string `json:"browser_fp,omitempty" xml:"browser_fp,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  Attack_type *string `json:"attack_type,omitempty" xml:"attack_type,omitempty" require:"true"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  Rule_name *string `json:"rule_name,omitempty" xml:"rule_name,omitempty" require:"true"`
  // {"en":"IP.", "zh_CN":"IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Event id.", "zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {"en":"Version.", "zh_CN":"版本号。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Client id.", "zh_CN":"客户端ID。"}
  Client_id *string `json:"client_id,omitempty" xml:"client_id,omitempty" require:"true"`
  // {"en":"URI.", "zh_CN":"URI。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Block id.", "zh_CN":"Block id。"}
  Block_id *string `json:"block_id,omitempty" xml:"block_id,omitempty" require:"true"`
  // {"en":"Content.", "zh_CN":"内容。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Request method.", "zh_CN":"请求方法。"}
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty" require:"true"`
  // {"en":"Rule id.", "zh_CN":"规则id。"}
  Final_rule_id *int `json:"final_rule_id,omitempty" xml:"final_rule_id,omitempty" require:"true"`
  // {"en":"Event type.", "zh_CN":"事件类型。"}
  Event_type *string `json:"event_type,omitempty" xml:"event_type,omitempty" require:"true"`
  // {"en":"Processing action.", "zh_CN":"处理动作。"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Zone.", "zh_CN":"时区。"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
  // {"en":"Attack time.", "zh_CN":"攻击时间"}
  Attack_time *string `json:"attack_time,omitempty" xml:"attack_time,omitempty" require:"true"`
  // {"en":"Strategy description.", "zh_CN":"策略描述。。"}
  Strategy_desc *string `json:"strategy_desc,omitempty" xml:"strategy_desc,omitempty" require:"true"`
  // {"en":"Domain.", "zh_CN":"域名。"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"IP geographical location.", "zh_CN":"IP地理位置。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Strategy name.", "zh_CN":"策略名称。"}
  Strategy_name *string `json:"strategy_name,omitempty" xml:"strategy_name,omitempty" require:"true"`
  // {"en":"User-Agent.", "zh_CN":"User-Agent"}
  User_agent *string `json:"user_agent,omitempty" xml:"user_agent,omitempty" require:"true"`
  // {"en":"Domain detail.", "zh_CN":"域名详情。"}
  Detail_host *string `json:"detail_host,omitempty" xml:"detail_host,omitempty" require:"true"`
  // {"en":"Status code.", "zh_CN":"状态码。"}
  StatusCode *int `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetReferer(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Referer = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetBrowser_fp(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Browser_fp = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetAttack_type(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Attack_type = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetRule_name(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Rule_name = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetIp(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Ip = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetUuid(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Uuid = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetVersion(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Version = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetClient_id(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Client_id = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetUrl(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Url = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetBlock_id(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Block_id = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetContent(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Content = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetMode(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Mode = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetFinal_rule_id(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Final_rule_id = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetEvent_type(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Event_type = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetAct(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Act = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetZone(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Zone = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetAttack_time(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Attack_time = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetStrategy_desc(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Strategy_desc = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetHost(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Host = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetLocation(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Location = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetStrategy_name(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Strategy_name = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetUser_agent(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.User_agent = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetDetail_host(v string) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.Detail_host = &v
  return s
}

func (s *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList) SetStatusCode(v int) *GetBotAttackIncidentLogDataGetBotAttackIncidentLogDataResponseDataList {
  s.StatusCode = &v
  return s
}

type GetBotAttackIncidentLogDataPaths struct {
}

func (s GetBotAttackIncidentLogDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataPaths) GoString() string {
  return s.String()
}

type GetBotAttackIncidentLogDataParameters struct {
}

func (s GetBotAttackIncidentLogDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataParameters) GoString() string {
  return s.String()
}

type GetBotAttackIncidentLogDataRequestHeader struct {
}

func (s GetBotAttackIncidentLogDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotAttackIncidentLogDataResponseHeader struct {
}

func (s GetBotAttackIncidentLogDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataResponseHeader) GoString() string {
  return s.String()
}




