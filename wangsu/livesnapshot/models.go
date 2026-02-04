package livesnapshot

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QuerySnapshotRuleRequest struct {
  // {"en":"rule id","zh_CN":"截图规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
  // {"en":"domain","zh_CN":"推流域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Publishing point name","zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
  // {"en":"stream name","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
  // {"en":"template id","zh_CN":"根据模版ID，查询关联该模版的规则"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
  // {"en":"pull domain","zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty"`
  // {"en":"page number","zh_CN":"分页编号，从1开始。默认为1"}
  PageNum *int `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
  // {"en":"page size","zh_CN":"分页大小，取值范围[1,200]。默认50"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QuerySnapshotRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleRequest) GoString() string {
  return s.String()
}

func (s *QuerySnapshotRuleRequest) SetRuleId(v string) *QuerySnapshotRuleRequest {
  s.RuleId = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetDomain(v string) *QuerySnapshotRuleRequest {
  s.Domain = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetAppName(v string) *QuerySnapshotRuleRequest {
  s.AppName = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetStreamName(v string) *QuerySnapshotRuleRequest {
  s.StreamName = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetTemplateId(v string) *QuerySnapshotRuleRequest {
  s.TemplateId = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetPullDomain(v string) *QuerySnapshotRuleRequest {
  s.PullDomain = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetPageNum(v int) *QuerySnapshotRuleRequest {
  s.PageNum = &v
  return s
}

func (s *QuerySnapshotRuleRequest) SetPageSize(v int) *QuerySnapshotRuleRequest {
  s.PageSize = &v
  return s
}

type QuerySnapshotRuleRequestHeader struct {
}

func (s QuerySnapshotRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleRequestHeader) GoString() string {
  return s.String()
}

type QuerySnapshotRulePaths struct {
}

func (s QuerySnapshotRulePaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRulePaths) GoString() string {
  return s.String()
}

type QuerySnapshotRuleParameters struct {
}

func (s QuerySnapshotRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleParameters) GoString() string {
  return s.String()
}

type QuerySnapshotRuleResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QuerySnapshotRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QuerySnapshotRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleResponse) GoString() string {
  return s.String()
}

func (s *QuerySnapshotRuleResponse) SetCode(v int) *QuerySnapshotRuleResponse {
  s.Code = &v
  return s
}

func (s *QuerySnapshotRuleResponse) SetMessage(v string) *QuerySnapshotRuleResponse {
  s.Message = &v
  return s
}

func (s *QuerySnapshotRuleResponse) SetData(v *QuerySnapshotRuleResponseData) *QuerySnapshotRuleResponse {
  s.Data = v
  return s
}

type QuerySnapshotRuleResponseData struct {
  // {"en":"data size.","zh_CN":"符合查询条件总数量"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of snapshot rule data.","zh_CN":"规则列表"}
  List []*QuerySnapshotRuleResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySnapshotRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleResponseData) GoString() string {
  return s.String()
}

func (s *QuerySnapshotRuleResponseData) SetTotal(v int) *QuerySnapshotRuleResponseData {
  s.Total = &v
  return s
}

func (s *QuerySnapshotRuleResponseData) SetList(v []*QuerySnapshotRuleResponseDataList) *QuerySnapshotRuleResponseData {
  s.List = v
  return s
}

type QuerySnapshotRuleResponseDataList struct     {
  // {"en":"rule id","zh_CN":"截图规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"template id","zh_CN":"模版ID"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
  // {"en":"domain","zh_CN":"推流域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Publishing point name","zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {"en":"stream name","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty" require:"true"`
  // {"en":"Extension parameters for the stream name.","zh_CN":"流名扩展参数"}
  StreamParams *string `json:"streamParams,omitempty" xml:"streamParams,omitempty" require:"true"`
  // {"en":"Pull stream domain","zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"Whether the rule is enabled. 0 for disabled, 1 for enabled. Defaults to 1.","zh_CN":"0：不启用，1：启用 默认为1"}
  IsEnabled *int `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"Creation timestamp","zh_CN":"创建时间，时间戳"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s QuerySnapshotRuleResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleResponseDataList) GoString() string {
  return s.String()
}

func (s *QuerySnapshotRuleResponseDataList) SetRuleId(v string) *QuerySnapshotRuleResponseDataList {
  s.RuleId = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetTemplateId(v string) *QuerySnapshotRuleResponseDataList {
  s.TemplateId = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetDomain(v string) *QuerySnapshotRuleResponseDataList {
  s.Domain = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetAppName(v string) *QuerySnapshotRuleResponseDataList {
  s.AppName = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetStreamName(v string) *QuerySnapshotRuleResponseDataList {
  s.StreamName = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetStreamParams(v string) *QuerySnapshotRuleResponseDataList {
  s.StreamParams = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetPullDomain(v string) *QuerySnapshotRuleResponseDataList {
  s.PullDomain = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetIsEnabled(v int) *QuerySnapshotRuleResponseDataList {
  s.IsEnabled = &v
  return s
}

func (s *QuerySnapshotRuleResponseDataList) SetCreateTime(v int64) *QuerySnapshotRuleResponseDataList {
  s.CreateTime = &v
  return s
}

type QuerySnapshotRuleResponseHeader struct {
}

func (s QuerySnapshotRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotRuleResponseHeader) GoString() string {
  return s.String()
}




type AddSnapshotParameterTemplateRequest struct {
  // {"en":"Geographical region of the cloud storage.","zh_CN":"区域"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Name of the screenshot parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"List of snapshot parameters.","zh_CN":"截图参数列表"}
  SnapshotParams []*AddSnapshotParameterTemplateRequestSnapshotParams `json:"snapshotParams,omitempty" xml:"snapshotParams,omitempty" type:"Repeated"`
}

func (s AddSnapshotParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *AddSnapshotParameterTemplateRequest) SetRegion(v string) *AddSnapshotParameterTemplateRequest {
  s.Region = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequest) SetTemplateName(v string) *AddSnapshotParameterTemplateRequest {
  s.TemplateName = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequest) SetAk(v string) *AddSnapshotParameterTemplateRequest {
  s.Ak = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequest) SetSk(v string) *AddSnapshotParameterTemplateRequest {
  s.Sk = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequest) SetMgrUrl(v string) *AddSnapshotParameterTemplateRequest {
  s.MgrUrl = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequest) SetNotifyUrl(v string) *AddSnapshotParameterTemplateRequest {
  s.NotifyUrl = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequest) SetSnapshotParams(v []*AddSnapshotParameterTemplateRequestSnapshotParams) *AddSnapshotParameterTemplateRequest {
  s.SnapshotParams = v
  return s
}

type AddSnapshotParameterTemplateRequestSnapshotParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
  // {"en":"Screenshot operation parameters.","zh_CN":"截图参数"}
  Fops *AddSnapshotParameterTemplateRequestSnapshotParamsFops `json:"fops,omitempty" xml:"fops,omitempty" type:"Struct"`
}

func (s AddSnapshotParameterTemplateRequestSnapshotParams) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateRequestSnapshotParams) GoString() string {
  return s.String()
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParams) SetBucketName(v string) *AddSnapshotParameterTemplateRequestSnapshotParams {
  s.BucketName = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParams) SetStorageTime(v int) *AddSnapshotParameterTemplateRequestSnapshotParams {
  s.StorageTime = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParams) SetFilePath(v string) *AddSnapshotParameterTemplateRequestSnapshotParams {
  s.FilePath = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParams) SetFops(v *AddSnapshotParameterTemplateRequestSnapshotParamsFops) *AddSnapshotParameterTemplateRequestSnapshotParams {
  s.Fops = v
  return s
}

type AddSnapshotParameterTemplateRequestSnapshotParamsFops struct {
  // {"en":"Format of the screenshot file (e.g., jpg, png).","zh_CN":"文件格式（如：jpg, png）"}
  FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
  // {"en":"Screenshot interval duration in seconds.","zh_CN":"截图间隔时长"}
  Interval *int `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"Screenshot width, in pixels (px).","zh_CN":"截图宽度，单位：像素（px）"}
  Width *int `json:"width,omitempty" xml:"width,omitempty"`
  // {"en":"Screenshot height, in pixels (px).","zh_CN":"截图高度，单位：像素（px）"}
  Height *int `json:"height,omitempty" xml:"height,omitempty"`
  // {"en":"The long side value for adaptive screenshot, in pixels (px).","zh_CN":"自适应截图的图像长边值，单位：像素（px）"}
  Longside *int `json:"longside,omitempty" xml:"longside,omitempty"`
  // {"en":"Timeout duration for stream pull in milliseconds.","zh_CN":"拉流超时时间"}
  TimeOut *int `json:"timeOut,omitempty" xml:"timeOut,omitempty"`
  // {"en":"Number of retries for stream pull timeout.","zh_CN":"拉流超时重试次数"}
  RetryTimes *int `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
}

func (s AddSnapshotParameterTemplateRequestSnapshotParamsFops) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateRequestSnapshotParamsFops) GoString() string {
  return s.String()
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetFileFormat(v string) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.FileFormat = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetInterval(v int) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Interval = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetWidth(v int) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Width = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetHeight(v int) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Height = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetLongside(v int) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Longside = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetTimeOut(v int) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.TimeOut = &v
  return s
}

func (s *AddSnapshotParameterTemplateRequestSnapshotParamsFops) SetRetryTimes(v int) *AddSnapshotParameterTemplateRequestSnapshotParamsFops {
  s.RetryTimes = &v
  return s
}

type AddSnapshotParameterTemplateRequestHeader struct {
}

func (s AddSnapshotParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type AddSnapshotParameterTemplatePaths struct {
}

func (s AddSnapshotParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplatePaths) GoString() string {
  return s.String()
}

type AddSnapshotParameterTemplateParameters struct {
}

func (s AddSnapshotParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateParameters) GoString() string {
  return s.String()
}

type AddSnapshotParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data payload.","zh_CN":"响应数据"}
  Data *AddSnapshotParameterTemplateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddSnapshotParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *AddSnapshotParameterTemplateResponse) SetCode(v int) *AddSnapshotParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *AddSnapshotParameterTemplateResponse) SetMessage(v string) *AddSnapshotParameterTemplateResponse {
  s.Message = &v
  return s
}

func (s *AddSnapshotParameterTemplateResponse) SetData(v *AddSnapshotParameterTemplateResponseData) *AddSnapshotParameterTemplateResponse {
  s.Data = v
  return s
}

type AddSnapshotParameterTemplateResponseData struct {
  // {"en":"ID of the created template.","zh_CN":"模版id"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s AddSnapshotParameterTemplateResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateResponseData) GoString() string {
  return s.String()
}

func (s *AddSnapshotParameterTemplateResponseData) SetTemplateId(v string) *AddSnapshotParameterTemplateResponseData {
  s.TemplateId = &v
  return s
}

type AddSnapshotParameterTemplateResponseHeader struct {
}

func (s AddSnapshotParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type AddSnapshotRulesRequest struct {
  // {"en":"template id","zh_CN":"模版ID"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
  // {"en":"domain","zh_CN":"推流域名，禁止传空字符串。  如果是推拉架构，推流域名必填"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Publishing point name","zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
  // {"en":"stream name","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
  // {"en":"Extension parameters for the stream name.","zh_CN":"流名扩展参数"}
  StreamParams *string `json:"streamParams,omitempty" xml:"streamParams,omitempty"`
  // {"en":"Pull stream domain","zh_CN":"拉流域名,禁止传空字符串"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"Whether the rule is enabled. 0 for disabled, 1 for enabled. Defaults to 1.","zh_CN":"0：不启用，1：启用 默认为1"}
  IsEnabled *int `json:"isEnabled,omitempty" xml:"isEnabled,omitempty"`
}

func (s AddSnapshotRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesRequest) GoString() string {
  return s.String()
}

func (s *AddSnapshotRulesRequest) SetTemplateId(v string) *AddSnapshotRulesRequest {
  s.TemplateId = &v
  return s
}

func (s *AddSnapshotRulesRequest) SetDomain(v string) *AddSnapshotRulesRequest {
  s.Domain = &v
  return s
}

func (s *AddSnapshotRulesRequest) SetAppName(v string) *AddSnapshotRulesRequest {
  s.AppName = &v
  return s
}

func (s *AddSnapshotRulesRequest) SetStreamName(v string) *AddSnapshotRulesRequest {
  s.StreamName = &v
  return s
}

func (s *AddSnapshotRulesRequest) SetStreamParams(v string) *AddSnapshotRulesRequest {
  s.StreamParams = &v
  return s
}

func (s *AddSnapshotRulesRequest) SetPullDomain(v string) *AddSnapshotRulesRequest {
  s.PullDomain = &v
  return s
}

func (s *AddSnapshotRulesRequest) SetIsEnabled(v int) *AddSnapshotRulesRequest {
  s.IsEnabled = &v
  return s
}

type AddSnapshotRulesRequestHeader struct {
}

func (s AddSnapshotRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesRequestHeader) GoString() string {
  return s.String()
}

type AddSnapshotRulesPaths struct {
}

func (s AddSnapshotRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesPaths) GoString() string {
  return s.String()
}

type AddSnapshotRulesParameters struct {
}

func (s AddSnapshotRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesParameters) GoString() string {
  return s.String()
}

type AddSnapshotRulesResponse struct {
  // {"en":"The response status code.","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed information or error message for the interface response.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The specific business data returned by the interface.","zh_CN":"返回数据"}
  Data *AddSnapshotRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddSnapshotRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesResponse) GoString() string {
  return s.String()
}

func (s *AddSnapshotRulesResponse) SetCode(v int) *AddSnapshotRulesResponse {
  s.Code = &v
  return s
}

func (s *AddSnapshotRulesResponse) SetMessage(v string) *AddSnapshotRulesResponse {
  s.Message = &v
  return s
}

func (s *AddSnapshotRulesResponse) SetData(v *AddSnapshotRulesResponseData) *AddSnapshotRulesResponse {
  s.Data = v
  return s
}

type AddSnapshotRulesResponseData struct {
  // {"en":"The ID of the newly created screenshot rule","zh_CN":"截图规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s AddSnapshotRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesResponseData) GoString() string {
  return s.String()
}

func (s *AddSnapshotRulesResponseData) SetRuleId(v string) *AddSnapshotRulesResponseData {
  s.RuleId = &v
  return s
}

type AddSnapshotRulesResponseHeader struct {
}

func (s AddSnapshotRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSnapshotRulesResponseHeader) GoString() string {
  return s.String()
}




type StopRealTimeSnapshotRequest struct {
}

func (s StopRealTimeSnapshotRequest) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeSnapshotRequest) GoString() string {
  return s.String()
}

type StopRealTimeSnapshotRequestHeader struct {
}

func (s StopRealTimeSnapshotRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeSnapshotRequestHeader) GoString() string {
  return s.String()
}

type StopRealTimeSnapshotPaths struct {
  // {"en":"The unique identifier for the snapshotting task.","zh_CN":"截图任务的唯一标识符"}
  PersistentId *string `json:"persistentId,omitempty" xml:"persistentId,omitempty" require:"true"`
}

func (s StopRealTimeSnapshotPaths) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeSnapshotPaths) GoString() string {
  return s.String()
}

func (s *StopRealTimeSnapshotPaths) SetPersistentId(v string) *StopRealTimeSnapshotPaths {
  s.PersistentId = &v
  return s
}

type StopRealTimeSnapshotParameters struct {
}

func (s StopRealTimeSnapshotParameters) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeSnapshotParameters) GoString() string {
  return s.String()
}

type StopRealTimeSnapshotResponse struct {
  // {"en":"The overall response code of the API call","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The overall response message of the API call","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s StopRealTimeSnapshotResponse) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeSnapshotResponse) GoString() string {
  return s.String()
}

func (s *StopRealTimeSnapshotResponse) SetCode(v int) *StopRealTimeSnapshotResponse {
  s.Code = &v
  return s
}

func (s *StopRealTimeSnapshotResponse) SetMessage(v string) *StopRealTimeSnapshotResponse {
  s.Message = &v
  return s
}

type StopRealTimeSnapshotResponseHeader struct {
}

func (s StopRealTimeSnapshotResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeSnapshotResponseHeader) GoString() string {
  return s.String()
}




type QuerySnapshotParameterTemplateRequest struct {
  // {"en":"The unique identifier of the screenshot parameter template.","zh_CN":"截图参数模板的唯一标识符"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QuerySnapshotParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateRequest) SetTemplateId(v string) *QuerySnapshotParameterTemplateRequest {
  s.TemplateId = &v
  return s
}

type QuerySnapshotParameterTemplateRequestHeader struct {
}

func (s QuerySnapshotParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type QuerySnapshotParameterTemplatePaths struct {
}

func (s QuerySnapshotParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplatePaths) GoString() string {
  return s.String()
}

type QuerySnapshotParameterTemplateParameters struct {
}

func (s QuerySnapshotParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateParameters) GoString() string {
  return s.String()
}

type QuerySnapshotParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QuerySnapshotParameterTemplateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QuerySnapshotParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateResponse) SetCode(v int) *QuerySnapshotParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponse) SetMessage(v string) *QuerySnapshotParameterTemplateResponse {
  s.Message = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponse) SetData(v *QuerySnapshotParameterTemplateResponseData) *QuerySnapshotParameterTemplateResponse {
  s.Data = v
  return s
}

type QuerySnapshotParameterTemplateResponseData struct {
  // {"en":"data size.","zh_CN":"总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of screenshot parameter template data.","zh_CN":"截图参数模板数据列表"}
  List []*QuerySnapshotParameterTemplateResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySnapshotParameterTemplateResponseData) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateResponseData) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateResponseData) SetTotal(v int) *QuerySnapshotParameterTemplateResponseData {
  s.Total = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseData) SetList(v []*QuerySnapshotParameterTemplateResponseDataList) *QuerySnapshotParameterTemplateResponseData {
  s.List = v
  return s
}

type QuerySnapshotParameterTemplateResponseDataList struct     {
  // {"en":"Geographical region of the cloud storage.","zh_CN":"区域"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"Name of the screenshot parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty" require:"true"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty" require:"true"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty" require:"true"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty" require:"true"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty" require:"true"`
  // {"en":"List of snapshot parameters.","zh_CN":"截图参数列表"}
  SnapshotParams []*QuerySnapshotParameterTemplateResponseDataListSnapshotParams `json:"snapshotParams,omitempty" xml:"snapshotParams,omitempty" require:"true" type:"Repeated"`
  // {"en":"The unique identifier of the screenshot parameter template.","zh_CN":"模版id"}
  TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s QuerySnapshotParameterTemplateResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateResponseDataList) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetRegion(v string) *QuerySnapshotParameterTemplateResponseDataList {
  s.Region = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetTemplateName(v string) *QuerySnapshotParameterTemplateResponseDataList {
  s.TemplateName = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetAk(v string) *QuerySnapshotParameterTemplateResponseDataList {
  s.Ak = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetSk(v string) *QuerySnapshotParameterTemplateResponseDataList {
  s.Sk = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetMgrUrl(v string) *QuerySnapshotParameterTemplateResponseDataList {
  s.MgrUrl = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetNotifyUrl(v string) *QuerySnapshotParameterTemplateResponseDataList {
  s.NotifyUrl = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetSnapshotParams(v []*QuerySnapshotParameterTemplateResponseDataListSnapshotParams) *QuerySnapshotParameterTemplateResponseDataList {
  s.SnapshotParams = v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataList) SetTemplateId(v int64) *QuerySnapshotParameterTemplateResponseDataList {
  s.TemplateId = &v
  return s
}

type QuerySnapshotParameterTemplateResponseDataListSnapshotParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty" require:"true"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty" require:"true"`
  // {"en":"Screenshot operation parameters.","zh_CN":"截图参数"}
  Fops *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops `json:"fops,omitempty" xml:"fops,omitempty" require:"true" type:"Struct"`
}

func (s QuerySnapshotParameterTemplateResponseDataListSnapshotParams) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateResponseDataListSnapshotParams) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParams) SetBucketName(v string) *QuerySnapshotParameterTemplateResponseDataListSnapshotParams {
  s.BucketName = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParams) SetStorageTime(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParams {
  s.StorageTime = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParams) SetFilePath(v string) *QuerySnapshotParameterTemplateResponseDataListSnapshotParams {
  s.FilePath = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParams) SetFops(v *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) *QuerySnapshotParameterTemplateResponseDataListSnapshotParams {
  s.Fops = v
  return s
}

type QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops struct {
  // {"en":"Format of the screenshot file (e.g., jpg, png).","zh_CN":"文件格式（如：jpg, png）"}
  FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty" require:"true"`
  // {"en":"Screenshot interval duration in seconds.","zh_CN":"截图间隔时长"}
  Interval *int `json:"interval,omitempty" xml:"interval,omitempty" require:"true"`
  // {"en":"Screenshot width, in pixels (px).","zh_CN":"截图宽度，单位：像素（px）"}
  Width *int `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"Screenshot height, in pixels (px).","zh_CN":"截图高度，单位：像素（px）"}
  Height *int `json:"height,omitempty" xml:"height,omitempty" require:"true"`
  // {"en":"The long side value for adaptive screenshot, in pixels (px).","zh_CN":"自适应截图的图像长边值，单位：像素（px）"}
  Longside *int `json:"longside,omitempty" xml:"longside,omitempty" require:"true"`
  // {"en":"Timeout duration for stream pull in milliseconds.","zh_CN":"拉流超时时间"}
  TimeOut *int `json:"timeOut,omitempty" xml:"timeOut,omitempty" require:"true"`
  // {"en":"Number of retries for stream pull timeout.","zh_CN":"拉流超时重试次数"}
  RetryTimes *int `json:"retryTimes,omitempty" xml:"retryTimes,omitempty" require:"true"`
}

func (s QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetFileFormat(v string) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.FileFormat = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetInterval(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.Interval = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetWidth(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.Width = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetHeight(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.Height = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetLongside(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.Longside = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetTimeOut(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.TimeOut = &v
  return s
}

func (s *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops) SetRetryTimes(v int) *QuerySnapshotParameterTemplateResponseDataListSnapshotParamsFops {
  s.RetryTimes = &v
  return s
}

type QuerySnapshotParameterTemplateResponseHeader struct {
}

func (s QuerySnapshotParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type DeleteSnapshotParameterTemplateRequest struct {
}

func (s DeleteSnapshotParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplateRequest) GoString() string {
  return s.String()
}

type DeleteSnapshotParameterTemplateRequestHeader struct {
}

func (s DeleteSnapshotParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type DeleteSnapshotParameterTemplatePaths struct {
  // {"en":"The unique identifier of the template to be deleted.","zh_CN":"待删除模版的唯一标识符"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s DeleteSnapshotParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *DeleteSnapshotParameterTemplatePaths) SetTemplateId(v string) *DeleteSnapshotParameterTemplatePaths {
  s.TemplateId = &v
  return s
}

type DeleteSnapshotParameterTemplateParameters struct {
}

func (s DeleteSnapshotParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplateParameters) GoString() string {
  return s.String()
}

type DeleteSnapshotParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteSnapshotParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *DeleteSnapshotParameterTemplateResponse) SetCode(v int) *DeleteSnapshotParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *DeleteSnapshotParameterTemplateResponse) SetMessage(v string) *DeleteSnapshotParameterTemplateResponse {
  s.Message = &v
  return s
}

type DeleteSnapshotParameterTemplateResponseHeader struct {
}

func (s DeleteSnapshotParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type DeleteScreenshotRulesRequest struct {
}

func (s DeleteScreenshotRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteScreenshotRulesRequest) GoString() string {
  return s.String()
}

type DeleteScreenshotRulesRequestHeader struct {
}

func (s DeleteScreenshotRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteScreenshotRulesRequestHeader) GoString() string {
  return s.String()
}

type DeleteScreenshotRulesPaths struct {
  // {"en":"The unique identifier of the screenshot rule to be deleted.","zh_CN":"规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s DeleteScreenshotRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteScreenshotRulesPaths) GoString() string {
  return s.String()
}

func (s *DeleteScreenshotRulesPaths) SetRuleId(v string) *DeleteScreenshotRulesPaths {
  s.RuleId = &v
  return s
}

type DeleteScreenshotRulesParameters struct {
}

func (s DeleteScreenshotRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteScreenshotRulesParameters) GoString() string {
  return s.String()
}

type DeleteScreenshotRulesResponse struct {
  // {"en":"The status code of the API response.","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed message describing the outcome of the API call.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteScreenshotRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteScreenshotRulesResponse) GoString() string {
  return s.String()
}

func (s *DeleteScreenshotRulesResponse) SetCode(v int) *DeleteScreenshotRulesResponse {
  s.Code = &v
  return s
}

func (s *DeleteScreenshotRulesResponse) SetMessage(v string) *DeleteScreenshotRulesResponse {
  s.Message = &v
  return s
}

type DeleteScreenshotRulesResponseHeader struct {
}

func (s DeleteScreenshotRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteScreenshotRulesResponseHeader) GoString() string {
  return s.String()
}




type ModifySnapshotParameterTemplateRequest struct {
  // {"en":"Name of the screenshot parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"List of snapshot parameters.","zh_CN":"截图参数列表"}
  SnapshotParams []*ModifySnapshotParameterTemplateRequestSnapshotParams `json:"snapshotParams,omitempty" xml:"snapshotParams,omitempty" type:"Repeated"`
}

func (s ModifySnapshotParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *ModifySnapshotParameterTemplateRequest) SetTemplateName(v string) *ModifySnapshotParameterTemplateRequest {
  s.TemplateName = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequest) SetAk(v string) *ModifySnapshotParameterTemplateRequest {
  s.Ak = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequest) SetSk(v string) *ModifySnapshotParameterTemplateRequest {
  s.Sk = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequest) SetMgrUrl(v string) *ModifySnapshotParameterTemplateRequest {
  s.MgrUrl = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequest) SetNotifyUrl(v string) *ModifySnapshotParameterTemplateRequest {
  s.NotifyUrl = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequest) SetSnapshotParams(v []*ModifySnapshotParameterTemplateRequestSnapshotParams) *ModifySnapshotParameterTemplateRequest {
  s.SnapshotParams = v
  return s
}

type ModifySnapshotParameterTemplateRequestSnapshotParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
  // {"en":"Screenshot operation parameters.","zh_CN":"截图参数"}
  Fops *ModifySnapshotParameterTemplateRequestSnapshotParamsFops `json:"fops,omitempty" xml:"fops,omitempty" type:"Struct"`
}

func (s ModifySnapshotParameterTemplateRequestSnapshotParams) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateRequestSnapshotParams) GoString() string {
  return s.String()
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParams) SetBucketName(v string) *ModifySnapshotParameterTemplateRequestSnapshotParams {
  s.BucketName = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParams) SetStorageTime(v int) *ModifySnapshotParameterTemplateRequestSnapshotParams {
  s.StorageTime = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParams) SetFilePath(v string) *ModifySnapshotParameterTemplateRequestSnapshotParams {
  s.FilePath = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParams) SetFops(v *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) *ModifySnapshotParameterTemplateRequestSnapshotParams {
  s.Fops = v
  return s
}

type ModifySnapshotParameterTemplateRequestSnapshotParamsFops struct {
  // {"en":"Format of the screenshot file (e.g., jpg, png).","zh_CN":"文件格式（如：jpg, png）"}
  FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
  // {"en":"Screenshot interval duration in seconds.","zh_CN":"截图间隔时长"}
  Interval *int `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"Screenshot width, in pixels (px).","zh_CN":"截图宽度，单位：像素（px）"}
  Width *int `json:"width,omitempty" xml:"width,omitempty"`
  // {"en":"Screenshot height, in pixels (px).","zh_CN":"截图高度，单位：像素（px）"}
  Height *int `json:"height,omitempty" xml:"height,omitempty"`
  // {"en":"The long side value for adaptive screenshot, in pixels (px).","zh_CN":"自适应截图的图像长边值，单位：像素（px）"}
  Longside *int `json:"longside,omitempty" xml:"longside,omitempty"`
  // {"en":"Timeout duration for stream pull in milliseconds.","zh_CN":"拉流超时时间"}
  TimeOut *int `json:"timeOut,omitempty" xml:"timeOut,omitempty"`
  // {"en":"Number of retries for stream pull timeout.","zh_CN":"拉流超时重试次数"}
  RetryTimes *int `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
}

func (s ModifySnapshotParameterTemplateRequestSnapshotParamsFops) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateRequestSnapshotParamsFops) GoString() string {
  return s.String()
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetFileFormat(v string) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.FileFormat = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetInterval(v int) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Interval = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetWidth(v int) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Width = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetHeight(v int) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Height = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetLongside(v int) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.Longside = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetTimeOut(v int) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.TimeOut = &v
  return s
}

func (s *ModifySnapshotParameterTemplateRequestSnapshotParamsFops) SetRetryTimes(v int) *ModifySnapshotParameterTemplateRequestSnapshotParamsFops {
  s.RetryTimes = &v
  return s
}

type ModifySnapshotParameterTemplateRequestHeader struct {
}

func (s ModifySnapshotParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type ModifySnapshotParameterTemplatePaths struct {
  // {"en":"Unique identifier of the screenshot parameter template to be modified.","zh_CN":"待修改的截图参数模板的唯一标识符。"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s ModifySnapshotParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *ModifySnapshotParameterTemplatePaths) SetTemplateId(v string) *ModifySnapshotParameterTemplatePaths {
  s.TemplateId = &v
  return s
}

type ModifySnapshotParameterTemplateParameters struct {
}

func (s ModifySnapshotParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateParameters) GoString() string {
  return s.String()
}

type ModifySnapshotParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifySnapshotParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *ModifySnapshotParameterTemplateResponse) SetCode(v int) *ModifySnapshotParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *ModifySnapshotParameterTemplateResponse) SetMessage(v string) *ModifySnapshotParameterTemplateResponse {
  s.Message = &v
  return s
}

type ModifySnapshotParameterTemplateResponseHeader struct {
}

func (s ModifySnapshotParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type StartRealTimeSnapshotRequest struct {
  // {"en":"The ID of the snapshot template to use for the real-time snapshot task","zh_CN":"模版ID"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
  // {"en":"The pull domain of the live stream","zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"The application name or publishing point for the stream","zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {"en":"The name of the live stream for which to take a real-time snapshot","zh_CN":"流名，支持多个流名，多个流名用英文逗号分隔。最多5个  示例：stream1,stream2,stream3"}
  StreamNames *string `json:"streamNames,omitempty" xml:"streamNames,omitempty" require:"true"`
}

func (s StartRealTimeSnapshotRequest) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotRequest) GoString() string {
  return s.String()
}

func (s *StartRealTimeSnapshotRequest) SetTemplateId(v string) *StartRealTimeSnapshotRequest {
  s.TemplateId = &v
  return s
}

func (s *StartRealTimeSnapshotRequest) SetPullDomain(v string) *StartRealTimeSnapshotRequest {
  s.PullDomain = &v
  return s
}

func (s *StartRealTimeSnapshotRequest) SetAppName(v string) *StartRealTimeSnapshotRequest {
  s.AppName = &v
  return s
}

func (s *StartRealTimeSnapshotRequest) SetStreamNames(v string) *StartRealTimeSnapshotRequest {
  s.StreamNames = &v
  return s
}

type StartRealTimeSnapshotRequestHeader struct {
}

func (s StartRealTimeSnapshotRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotRequestHeader) GoString() string {
  return s.String()
}

type StartRealTimeSnapshotPaths struct {
}

func (s StartRealTimeSnapshotPaths) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotPaths) GoString() string {
  return s.String()
}

type StartRealTimeSnapshotParameters struct {
}

func (s StartRealTimeSnapshotParameters) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotParameters) GoString() string {
  return s.String()
}

type StartRealTimeSnapshotResponse struct {
  // {"en":"The overall response code of the API call","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The overall response message of the API call","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The detailed response data containing snapshotting statuses","zh_CN":"返回数据"}
  Data []*StartRealTimeSnapshotResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s StartRealTimeSnapshotResponse) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotResponse) GoString() string {
  return s.String()
}

func (s *StartRealTimeSnapshotResponse) SetCode(v int) *StartRealTimeSnapshotResponse {
  s.Code = &v
  return s
}

func (s *StartRealTimeSnapshotResponse) SetMessage(v string) *StartRealTimeSnapshotResponse {
  s.Message = &v
  return s
}

func (s *StartRealTimeSnapshotResponse) SetData(v []*StartRealTimeSnapshotResponseData) *StartRealTimeSnapshotResponse {
  s.Data = v
  return s
}

type StartRealTimeSnapshotResponseData struct     {
  // {"en":"Stream-level response code","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Stream-level response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The name of the stream this snapshotting status refers to","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty" require:"true"`
  // {"en":"The unique ID of the real-time snapshotting task","zh_CN":"截图任务的id"}
  PersistentId *string `json:"persistentId,omitempty" xml:"persistentId,omitempty" require:"true"`
}

func (s StartRealTimeSnapshotResponseData) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotResponseData) GoString() string {
  return s.String()
}

func (s *StartRealTimeSnapshotResponseData) SetCode(v int) *StartRealTimeSnapshotResponseData {
  s.Code = &v
  return s
}

func (s *StartRealTimeSnapshotResponseData) SetMessage(v string) *StartRealTimeSnapshotResponseData {
  s.Message = &v
  return s
}

func (s *StartRealTimeSnapshotResponseData) SetStreamName(v string) *StartRealTimeSnapshotResponseData {
  s.StreamName = &v
  return s
}

func (s *StartRealTimeSnapshotResponseData) SetPersistentId(v string) *StartRealTimeSnapshotResponseData {
  s.PersistentId = &v
  return s
}

type StartRealTimeSnapshotResponseHeader struct {
}

func (s StartRealTimeSnapshotResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeSnapshotResponseHeader) GoString() string {
  return s.String()
}




type ModifySnapshotRulesRequest struct {
  // {"en":"The unique identifier of the template.","zh_CN":"模版ID,禁止传空字符串。  不传：不修改原来的值"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
  // {"en":"The domain name.","zh_CN":"推流域名，如果是推拉架构，推流域名必填，不能传空字符串  不传：不修改原来的值  空字符串：删除配置"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Publishing point name","zh_CN":"发布点  不传：不修改原来的值  空字符串：删除配置"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
  // {"en":"The name of the stream.","zh_CN":"流名  不传：不修改原来的值  空字符串：删除配置"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
  // {"en":"Extension parameters for the stream name.","zh_CN":"流名扩展参数。  不传：不修改原来的值  空字符串：删除配置"}
  StreamParams *string `json:"streamParams,omitempty" xml:"streamParams,omitempty"`
  // {"en":"Whether the rule is enabled. 0 for disabled, 1 for enabled. Defaults to 1.","zh_CN":"0：不启用，1：启用 默认为1"}
  IsEnabled *int `json:"isEnabled,omitempty" xml:"isEnabled,omitempty"`
  // {"en":"pull domain","zh_CN":"拉流域名,禁止传空字符串  不传：不修改原来的值"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty"`
}

func (s ModifySnapshotRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotRulesRequest) GoString() string {
  return s.String()
}

func (s *ModifySnapshotRulesRequest) SetTemplateId(v string) *ModifySnapshotRulesRequest {
  s.TemplateId = &v
  return s
}

func (s *ModifySnapshotRulesRequest) SetDomain(v string) *ModifySnapshotRulesRequest {
  s.Domain = &v
  return s
}

func (s *ModifySnapshotRulesRequest) SetAppName(v string) *ModifySnapshotRulesRequest {
  s.AppName = &v
  return s
}

func (s *ModifySnapshotRulesRequest) SetStreamName(v string) *ModifySnapshotRulesRequest {
  s.StreamName = &v
  return s
}

func (s *ModifySnapshotRulesRequest) SetStreamParams(v string) *ModifySnapshotRulesRequest {
  s.StreamParams = &v
  return s
}

func (s *ModifySnapshotRulesRequest) SetIsEnabled(v int) *ModifySnapshotRulesRequest {
  s.IsEnabled = &v
  return s
}

func (s *ModifySnapshotRulesRequest) SetPullDomain(v string) *ModifySnapshotRulesRequest {
  s.PullDomain = &v
  return s
}

type ModifySnapshotRulesRequestHeader struct {
}

func (s ModifySnapshotRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotRulesRequestHeader) GoString() string {
  return s.String()
}

type ModifySnapshotRulesPaths struct {
  // {"en":"The unique identifier of the rule to be modified.","zh_CN":"截图规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s ModifySnapshotRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotRulesPaths) GoString() string {
  return s.String()
}

func (s *ModifySnapshotRulesPaths) SetRuleId(v string) *ModifySnapshotRulesPaths {
  s.RuleId = &v
  return s
}

type ModifySnapshotRulesParameters struct {
}

func (s ModifySnapshotRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotRulesParameters) GoString() string {
  return s.String()
}

type ModifySnapshotRulesResponse struct {
  // {"en":"The status code of the API response.","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed message describing the outcome of the API call.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifySnapshotRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotRulesResponse) GoString() string {
  return s.String()
}

func (s *ModifySnapshotRulesResponse) SetCode(v int) *ModifySnapshotRulesResponse {
  s.Code = &v
  return s
}

func (s *ModifySnapshotRulesResponse) SetMessage(v string) *ModifySnapshotRulesResponse {
  s.Message = &v
  return s
}

type ModifySnapshotRulesResponseHeader struct {
}

func (s ModifySnapshotRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotRulesResponseHeader) GoString() string {
  return s.String()
}




