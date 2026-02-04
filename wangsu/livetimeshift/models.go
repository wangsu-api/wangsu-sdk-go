package livetimeshift

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteTimeShiftParameterTemplateRequest struct {
}

func (s DeleteTimeShiftParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplateRequest) GoString() string {
  return s.String()
}

type DeleteTimeShiftParameterTemplateRequestHeader struct {
}

func (s DeleteTimeShiftParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type DeleteTimeShiftParameterTemplatePaths struct {
  // {"en":"The unique identifier of the time-shift parameter template to be deleted.","zh_CN":"要删除的时移参数模版的唯一标识符。"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s DeleteTimeShiftParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *DeleteTimeShiftParameterTemplatePaths) SetTemplateId(v string) *DeleteTimeShiftParameterTemplatePaths {
  s.TemplateId = &v
  return s
}

type DeleteTimeShiftParameterTemplateParameters struct {
}

func (s DeleteTimeShiftParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplateParameters) GoString() string {
  return s.String()
}

type DeleteTimeShiftParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteTimeShiftParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *DeleteTimeShiftParameterTemplateResponse) SetCode(v int) *DeleteTimeShiftParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *DeleteTimeShiftParameterTemplateResponse) SetMessage(v string) *DeleteTimeShiftParameterTemplateResponse {
  s.Message = &v
  return s
}

type DeleteTimeShiftParameterTemplateResponseHeader struct {
}

func (s DeleteTimeShiftParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type StopRealTimeTimeShiftRecordingRequest struct {
}

func (s StopRealTimeTimeShiftRecordingRequest) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeTimeShiftRecordingRequest) GoString() string {
  return s.String()
}

type StopRealTimeTimeShiftRecordingRequestHeader struct {
}

func (s StopRealTimeTimeShiftRecordingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeTimeShiftRecordingRequestHeader) GoString() string {
  return s.String()
}

type StopRealTimeTimeShiftRecordingPaths struct {
  // {"en":"The unique identifier for the time shift recording task.","zh_CN":"时移录制任务的唯一标识符"}
  PersistentId *string `json:"persistentId,omitempty" xml:"persistentId,omitempty" require:"true"`
}

func (s StopRealTimeTimeShiftRecordingPaths) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeTimeShiftRecordingPaths) GoString() string {
  return s.String()
}

func (s *StopRealTimeTimeShiftRecordingPaths) SetPersistentId(v string) *StopRealTimeTimeShiftRecordingPaths {
  s.PersistentId = &v
  return s
}

type StopRealTimeTimeShiftRecordingParameters struct {
}

func (s StopRealTimeTimeShiftRecordingParameters) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeTimeShiftRecordingParameters) GoString() string {
  return s.String()
}

type StopRealTimeTimeShiftRecordingResponse struct {
  // {"en":"The overall response code of the API call","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The overall response message of the API call","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s StopRealTimeTimeShiftRecordingResponse) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeTimeShiftRecordingResponse) GoString() string {
  return s.String()
}

func (s *StopRealTimeTimeShiftRecordingResponse) SetCode(v int) *StopRealTimeTimeShiftRecordingResponse {
  s.Code = &v
  return s
}

func (s *StopRealTimeTimeShiftRecordingResponse) SetMessage(v string) *StopRealTimeTimeShiftRecordingResponse {
  s.Message = &v
  return s
}

type StopRealTimeTimeShiftRecordingResponseHeader struct {
}

func (s StopRealTimeTimeShiftRecordingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeTimeShiftRecordingResponseHeader) GoString() string {
  return s.String()
}




type ModifyTimeShiftParameterTemplateRequest struct {
  // {"en":"Name of the time-shift parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"List of time shift parameters.","zh_CN":"时移参数列表"}
  TimeShiftParams []*ModifyTimeShiftParameterTemplateRequestTimeShiftParams `json:"timeShiftParams,omitempty" xml:"timeShiftParams,omitempty" type:"Repeated"`
}

func (s ModifyTimeShiftParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftParameterTemplateRequest) SetTemplateName(v string) *ModifyTimeShiftParameterTemplateRequest {
  s.TemplateName = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequest) SetAk(v string) *ModifyTimeShiftParameterTemplateRequest {
  s.Ak = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequest) SetSk(v string) *ModifyTimeShiftParameterTemplateRequest {
  s.Sk = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequest) SetMgrUrl(v string) *ModifyTimeShiftParameterTemplateRequest {
  s.MgrUrl = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequest) SetNotifyUrl(v string) *ModifyTimeShiftParameterTemplateRequest {
  s.NotifyUrl = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequest) SetTimeShiftParams(v []*ModifyTimeShiftParameterTemplateRequestTimeShiftParams) *ModifyTimeShiftParameterTemplateRequest {
  s.TimeShiftParams = v
  return s
}

type ModifyTimeShiftParameterTemplateRequestTimeShiftParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
  // {"en":"Time shift operation parameters.","zh_CN":"时移参数"}
  Fops *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops `json:"fops,omitempty" xml:"fops,omitempty" type:"Struct"`
}

func (s ModifyTimeShiftParameterTemplateRequestTimeShiftParams) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateRequestTimeShiftParams) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParams) SetBucketName(v string) *ModifyTimeShiftParameterTemplateRequestTimeShiftParams {
  s.BucketName = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParams) SetStorageTime(v int) *ModifyTimeShiftParameterTemplateRequestTimeShiftParams {
  s.StorageTime = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParams) SetFilePath(v string) *ModifyTimeShiftParameterTemplateRequestTimeShiftParams {
  s.FilePath = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParams) SetFops(v *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) *ModifyTimeShiftParameterTemplateRequestTimeShiftParams {
  s.Fops = v
  return s
}

type ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops struct {
  // {"en":"Duration of each TS segment in m3u8 in seconds.","zh_CN":"指定m3u8的分段ts时长"}
  Segtime *int `json:"segtime,omitempty" xml:"segtime,omitempty"`
  // {"en":"Whether to remove the audio stream (0: No, 1: Yes).","zh_CN":"是否去除音频流"}
  RemoveAudio *int `json:"removeAudio,omitempty" xml:"removeAudio,omitempty"`
  // {"en":"Whether to remove the video stream (0: No, 1: Yes).","zh_CN":"是否去除视频流"}
  RemoveVideo *int `json:"removeVideo,omitempty" xml:"removeVideo,omitempty"`
  // {"en":"Timeout duration for stream pull in milliseconds.","zh_CN":"拉流超时时间"}
  TimeOut *int `json:"timeOut,omitempty" xml:"timeOut,omitempty"`
  // {"en":"Number of retries for stream pull timeout.","zh_CN":"拉流超时重试次数"}
  RetryTimes *int `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
}

func (s ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetSegtime(v int) *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.Segtime = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetRemoveAudio(v int) *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.RemoveAudio = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetRemoveVideo(v int) *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.RemoveVideo = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetTimeOut(v int) *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.TimeOut = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetRetryTimes(v int) *ModifyTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.RetryTimes = &v
  return s
}

type ModifyTimeShiftParameterTemplateRequestHeader struct {
}

func (s ModifyTimeShiftParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type ModifyTimeShiftParameterTemplatePaths struct {
  // {"en":"The unique identifier of the time-shift parameter template to be modified.","zh_CN":"模版ID，用于唯一标识需要修改的时移参数模版。"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s ModifyTimeShiftParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftParameterTemplatePaths) SetTemplateId(v string) *ModifyTimeShiftParameterTemplatePaths {
  s.TemplateId = &v
  return s
}

type ModifyTimeShiftParameterTemplateParameters struct {
}

func (s ModifyTimeShiftParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateParameters) GoString() string {
  return s.String()
}

type ModifyTimeShiftParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifyTimeShiftParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftParameterTemplateResponse) SetCode(v int) *ModifyTimeShiftParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *ModifyTimeShiftParameterTemplateResponse) SetMessage(v string) *ModifyTimeShiftParameterTemplateResponse {
  s.Message = &v
  return s
}

type ModifyTimeShiftParameterTemplateResponseHeader struct {
}

func (s ModifyTimeShiftParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type ModifyTimeShiftRulesRequest struct {
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

func (s ModifyTimeShiftRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftRulesRequest) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftRulesRequest) SetTemplateId(v string) *ModifyTimeShiftRulesRequest {
  s.TemplateId = &v
  return s
}

func (s *ModifyTimeShiftRulesRequest) SetDomain(v string) *ModifyTimeShiftRulesRequest {
  s.Domain = &v
  return s
}

func (s *ModifyTimeShiftRulesRequest) SetAppName(v string) *ModifyTimeShiftRulesRequest {
  s.AppName = &v
  return s
}

func (s *ModifyTimeShiftRulesRequest) SetStreamName(v string) *ModifyTimeShiftRulesRequest {
  s.StreamName = &v
  return s
}

func (s *ModifyTimeShiftRulesRequest) SetStreamParams(v string) *ModifyTimeShiftRulesRequest {
  s.StreamParams = &v
  return s
}

func (s *ModifyTimeShiftRulesRequest) SetIsEnabled(v int) *ModifyTimeShiftRulesRequest {
  s.IsEnabled = &v
  return s
}

func (s *ModifyTimeShiftRulesRequest) SetPullDomain(v string) *ModifyTimeShiftRulesRequest {
  s.PullDomain = &v
  return s
}

type ModifyTimeShiftRulesRequestHeader struct {
}

func (s ModifyTimeShiftRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftRulesRequestHeader) GoString() string {
  return s.String()
}

type ModifyTimeShiftRulesPaths struct {
  // {"en":"The unique identifier of the rule to be modified.","zh_CN":"时移规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s ModifyTimeShiftRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftRulesPaths) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftRulesPaths) SetRuleId(v string) *ModifyTimeShiftRulesPaths {
  s.RuleId = &v
  return s
}

type ModifyTimeShiftRulesParameters struct {
}

func (s ModifyTimeShiftRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftRulesParameters) GoString() string {
  return s.String()
}

type ModifyTimeShiftRulesResponse struct {
  // {"en":"The status code of the API response.","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed message describing the outcome of the API call.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifyTimeShiftRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftRulesResponse) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftRulesResponse) SetCode(v int) *ModifyTimeShiftRulesResponse {
  s.Code = &v
  return s
}

func (s *ModifyTimeShiftRulesResponse) SetMessage(v string) *ModifyTimeShiftRulesResponse {
  s.Message = &v
  return s
}

type ModifyTimeShiftRulesResponseHeader struct {
}

func (s ModifyTimeShiftRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftRulesResponseHeader) GoString() string {
  return s.String()
}




type StartRealTimeTimeShiftRecordingRequest struct {
  // {"en":"The ID of the time shifting template to use","zh_CN":"模版ID"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
  // {"en":"The pull domain of the live stream","zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"The application name or publishing point for the stream","zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {"en":"The name of the live stream to be time shifted","zh_CN":"流名，支持多个流名，多个流名用英文逗号分隔。最多5个  示例：stream1,stream2,stream3"}
  StreamNames *string `json:"streamNames,omitempty" xml:"streamNames,omitempty" require:"true"`
}

func (s StartRealTimeTimeShiftRecordingRequest) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingRequest) GoString() string {
  return s.String()
}

func (s *StartRealTimeTimeShiftRecordingRequest) SetTemplateId(v string) *StartRealTimeTimeShiftRecordingRequest {
  s.TemplateId = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingRequest) SetPullDomain(v string) *StartRealTimeTimeShiftRecordingRequest {
  s.PullDomain = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingRequest) SetAppName(v string) *StartRealTimeTimeShiftRecordingRequest {
  s.AppName = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingRequest) SetStreamNames(v string) *StartRealTimeTimeShiftRecordingRequest {
  s.StreamNames = &v
  return s
}

type StartRealTimeTimeShiftRecordingRequestHeader struct {
}

func (s StartRealTimeTimeShiftRecordingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingRequestHeader) GoString() string {
  return s.String()
}

type StartRealTimeTimeShiftRecordingPaths struct {
}

func (s StartRealTimeTimeShiftRecordingPaths) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingPaths) GoString() string {
  return s.String()
}

type StartRealTimeTimeShiftRecordingParameters struct {
}

func (s StartRealTimeTimeShiftRecordingParameters) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingParameters) GoString() string {
  return s.String()
}

type StartRealTimeTimeShiftRecordingResponse struct {
  // {"en":"The overall response code of the API call","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The overall response message of the API call","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The detailed response data containing time shifting statuses","zh_CN":"返回数据"}
  Data []*StartRealTimeTimeShiftRecordingResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s StartRealTimeTimeShiftRecordingResponse) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingResponse) GoString() string {
  return s.String()
}

func (s *StartRealTimeTimeShiftRecordingResponse) SetCode(v int) *StartRealTimeTimeShiftRecordingResponse {
  s.Code = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingResponse) SetMessage(v string) *StartRealTimeTimeShiftRecordingResponse {
  s.Message = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingResponse) SetData(v []*StartRealTimeTimeShiftRecordingResponseData) *StartRealTimeTimeShiftRecordingResponse {
  s.Data = v
  return s
}

type StartRealTimeTimeShiftRecordingResponseData struct     {
  // {"en":"Stream-level response code","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Stream-level response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The name of the stream this status refers to","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty" require:"true"`
  // {"en":"The unique ID of the real-time time shifting task","zh_CN":"时移录制任务的id"}
  PersistentId *string `json:"persistentId,omitempty" xml:"persistentId,omitempty" require:"true"`
}

func (s StartRealTimeTimeShiftRecordingResponseData) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingResponseData) GoString() string {
  return s.String()
}

func (s *StartRealTimeTimeShiftRecordingResponseData) SetCode(v int) *StartRealTimeTimeShiftRecordingResponseData {
  s.Code = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingResponseData) SetMessage(v string) *StartRealTimeTimeShiftRecordingResponseData {
  s.Message = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingResponseData) SetStreamName(v string) *StartRealTimeTimeShiftRecordingResponseData {
  s.StreamName = &v
  return s
}

func (s *StartRealTimeTimeShiftRecordingResponseData) SetPersistentId(v string) *StartRealTimeTimeShiftRecordingResponseData {
  s.PersistentId = &v
  return s
}

type StartRealTimeTimeShiftRecordingResponseHeader struct {
}

func (s StartRealTimeTimeShiftRecordingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeTimeShiftRecordingResponseHeader) GoString() string {
  return s.String()
}




type AddTimeShiftParameterTemplateRequest struct {
  // {"en":"Geographical region of the cloud storage.","zh_CN":"区域"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Name of the time-shift parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"List of time shift parameters.","zh_CN":"时移参数列表"}
  TimeShiftParams []*AddTimeShiftParameterTemplateRequestTimeShiftParams `json:"timeShiftParams,omitempty" xml:"timeShiftParams,omitempty" type:"Repeated"`
}

func (s AddTimeShiftParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *AddTimeShiftParameterTemplateRequest) SetRegion(v string) *AddTimeShiftParameterTemplateRequest {
  s.Region = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequest) SetTemplateName(v string) *AddTimeShiftParameterTemplateRequest {
  s.TemplateName = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequest) SetAk(v string) *AddTimeShiftParameterTemplateRequest {
  s.Ak = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequest) SetSk(v string) *AddTimeShiftParameterTemplateRequest {
  s.Sk = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequest) SetMgrUrl(v string) *AddTimeShiftParameterTemplateRequest {
  s.MgrUrl = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequest) SetNotifyUrl(v string) *AddTimeShiftParameterTemplateRequest {
  s.NotifyUrl = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequest) SetTimeShiftParams(v []*AddTimeShiftParameterTemplateRequestTimeShiftParams) *AddTimeShiftParameterTemplateRequest {
  s.TimeShiftParams = v
  return s
}

type AddTimeShiftParameterTemplateRequestTimeShiftParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
  // {"en":"Time shift operation parameters.","zh_CN":"时移参数"}
  Fops *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops `json:"fops,omitempty" xml:"fops,omitempty" type:"Struct"`
}

func (s AddTimeShiftParameterTemplateRequestTimeShiftParams) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateRequestTimeShiftParams) GoString() string {
  return s.String()
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParams) SetBucketName(v string) *AddTimeShiftParameterTemplateRequestTimeShiftParams {
  s.BucketName = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParams) SetStorageTime(v int) *AddTimeShiftParameterTemplateRequestTimeShiftParams {
  s.StorageTime = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParams) SetFilePath(v string) *AddTimeShiftParameterTemplateRequestTimeShiftParams {
  s.FilePath = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParams) SetFops(v *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) *AddTimeShiftParameterTemplateRequestTimeShiftParams {
  s.Fops = v
  return s
}

type AddTimeShiftParameterTemplateRequestTimeShiftParamsFops struct {
  // {"en":"Duration of each TS segment in m3u8 in seconds.","zh_CN":"指定m3u8的分段ts时长"}
  Segtime *int `json:"segtime,omitempty" xml:"segtime,omitempty"`
  // {"en":"Whether to remove the audio stream (0: No, 1: Yes).","zh_CN":"是否去除音频流"}
  RemoveAudio *int `json:"removeAudio,omitempty" xml:"removeAudio,omitempty"`
  // {"en":"Whether to remove the video stream (0: No, 1: Yes).","zh_CN":"是否去除视频流"}
  RemoveVideo *int `json:"removeVideo,omitempty" xml:"removeVideo,omitempty"`
  // {"en":"Timeout duration for stream pull in milliseconds.","zh_CN":"拉流超时时间"}
  TimeOut *int `json:"timeOut,omitempty" xml:"timeOut,omitempty"`
  // {"en":"Number of retries for stream pull timeout.","zh_CN":"拉流超时重试次数"}
  RetryTimes *int `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
}

func (s AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) GoString() string {
  return s.String()
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetSegtime(v int) *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.Segtime = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetRemoveAudio(v int) *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.RemoveAudio = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetRemoveVideo(v int) *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.RemoveVideo = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetTimeOut(v int) *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.TimeOut = &v
  return s
}

func (s *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops) SetRetryTimes(v int) *AddTimeShiftParameterTemplateRequestTimeShiftParamsFops {
  s.RetryTimes = &v
  return s
}

type AddTimeShiftParameterTemplateRequestHeader struct {
}

func (s AddTimeShiftParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type AddTimeShiftParameterTemplatePaths struct {
}

func (s AddTimeShiftParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplatePaths) GoString() string {
  return s.String()
}

type AddTimeShiftParameterTemplateParameters struct {
}

func (s AddTimeShiftParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateParameters) GoString() string {
  return s.String()
}

type AddTimeShiftParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data payload.","zh_CN":"响应数据"}
  Data *AddTimeShiftParameterTemplateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddTimeShiftParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *AddTimeShiftParameterTemplateResponse) SetCode(v int) *AddTimeShiftParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *AddTimeShiftParameterTemplateResponse) SetMessage(v string) *AddTimeShiftParameterTemplateResponse {
  s.Message = &v
  return s
}

func (s *AddTimeShiftParameterTemplateResponse) SetData(v *AddTimeShiftParameterTemplateResponseData) *AddTimeShiftParameterTemplateResponse {
  s.Data = v
  return s
}

type AddTimeShiftParameterTemplateResponseData struct {
  // {"en":"ID of the created template.","zh_CN":"模版id"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s AddTimeShiftParameterTemplateResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateResponseData) GoString() string {
  return s.String()
}

func (s *AddTimeShiftParameterTemplateResponseData) SetTemplateId(v string) *AddTimeShiftParameterTemplateResponseData {
  s.TemplateId = &v
  return s
}

type AddTimeShiftParameterTemplateResponseHeader struct {
}

func (s AddTimeShiftParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type QueryTimeShiftParameterTemplateRequest struct {
  // {"en":"The ID of the time-shift parameter template to query.","zh_CN":"要查询的时移参数模板ID"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QueryTimeShiftParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateRequest) SetTemplateId(v string) *QueryTimeShiftParameterTemplateRequest {
  s.TemplateId = &v
  return s
}

type QueryTimeShiftParameterTemplateRequestHeader struct {
}

func (s QueryTimeShiftParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type QueryTimeShiftParameterTemplatePaths struct {
}

func (s QueryTimeShiftParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplatePaths) GoString() string {
  return s.String()
}

type QueryTimeShiftParameterTemplateParameters struct {
}

func (s QueryTimeShiftParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateParameters) GoString() string {
  return s.String()
}

type QueryTimeShiftParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryTimeShiftParameterTemplateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryTimeShiftParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateResponse) SetCode(v int) *QueryTimeShiftParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponse) SetMessage(v string) *QueryTimeShiftParameterTemplateResponse {
  s.Message = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponse) SetData(v *QueryTimeShiftParameterTemplateResponseData) *QueryTimeShiftParameterTemplateResponse {
  s.Data = v
  return s
}

type QueryTimeShiftParameterTemplateResponseData struct {
  // {"en":"data size.","zh_CN":"总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of time-shift parameter template data.","zh_CN":"时移参数模板数据列表"}
  List []*QueryTimeShiftParameterTemplateResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTimeShiftParameterTemplateResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateResponseData) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateResponseData) SetTotal(v int) *QueryTimeShiftParameterTemplateResponseData {
  s.Total = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseData) SetList(v []*QueryTimeShiftParameterTemplateResponseDataList) *QueryTimeShiftParameterTemplateResponseData {
  s.List = v
  return s
}

type QueryTimeShiftParameterTemplateResponseDataList struct     {
  // {"en":"The unique identifier of the time-shift parameter template.","zh_CN":"时移参数模板的唯一标识符。"}
  TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
  // {"en":"Geographical region of the time-shift parameter template","zh_CN":"时移参数模板所在的地理区域"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"Name of the time-shift parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty" require:"true"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty" require:"true"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty" require:"true"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty" require:"true"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty" require:"true"`
  // {"en":"List of time-shift parameter templates.","zh_CN":"时移参数模板数据列表。"}
  TimeShiftParams []*QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams `json:"timeShiftParams,omitempty" xml:"timeShiftParams,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTimeShiftParameterTemplateResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateResponseDataList) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetTemplateId(v int64) *QueryTimeShiftParameterTemplateResponseDataList {
  s.TemplateId = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetRegion(v string) *QueryTimeShiftParameterTemplateResponseDataList {
  s.Region = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetTemplateName(v string) *QueryTimeShiftParameterTemplateResponseDataList {
  s.TemplateName = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetAk(v string) *QueryTimeShiftParameterTemplateResponseDataList {
  s.Ak = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetSk(v string) *QueryTimeShiftParameterTemplateResponseDataList {
  s.Sk = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetMgrUrl(v string) *QueryTimeShiftParameterTemplateResponseDataList {
  s.MgrUrl = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetNotifyUrl(v string) *QueryTimeShiftParameterTemplateResponseDataList {
  s.NotifyUrl = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataList) SetTimeShiftParams(v []*QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) *QueryTimeShiftParameterTemplateResponseDataList {
  s.TimeShiftParams = v
  return s
}

type QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty" require:"true"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty" require:"true"`
  // {"en":"Time-shift operation configuration.","zh_CN":"时移参数"}
  Fops *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops `json:"fops,omitempty" xml:"fops,omitempty" require:"true" type:"Struct"`
}

func (s QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) SetBucketName(v string) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams {
  s.BucketName = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) SetStorageTime(v int) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams {
  s.StorageTime = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) SetFilePath(v string) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams {
  s.FilePath = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams) SetFops(v *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParams {
  s.Fops = v
  return s
}

type QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops struct {
  // {"en":"Duration of each TS segment in m3u8 in seconds.","zh_CN":"指定m3u8的分段ts时长"}
  Segtime *int `json:"segtime,omitempty" xml:"segtime,omitempty" require:"true"`
  // {"en":"Whether to remove the audio stream (0: No, 1: Yes).","zh_CN":"是否去除音频流"}
  RemoveAudio *int `json:"removeAudio,omitempty" xml:"removeAudio,omitempty" require:"true"`
  // {"en":"Whether to remove the video stream (0: No, 1: Yes).","zh_CN":"是否去除视频流"}
  RemoveVideo *int `json:"removeVideo,omitempty" xml:"removeVideo,omitempty" require:"true"`
  // {"en":"Timeout duration for stream pull in milliseconds.","zh_CN":"拉流超时时间"}
  TimeOut *int `json:"timeOut,omitempty" xml:"timeOut,omitempty" require:"true"`
  // {"en":"Number of retries for stream pull timeout.","zh_CN":"拉流超时重试次数"}
  RetryTimes *int `json:"retryTimes,omitempty" xml:"retryTimes,omitempty" require:"true"`
}

func (s QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) SetSegtime(v int) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops {
  s.Segtime = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) SetRemoveAudio(v int) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops {
  s.RemoveAudio = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) SetRemoveVideo(v int) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops {
  s.RemoveVideo = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) SetTimeOut(v int) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops {
  s.TimeOut = &v
  return s
}

func (s *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops) SetRetryTimes(v int) *QueryTimeShiftParameterTemplateResponseDataListTimeShiftParamsFops {
  s.RetryTimes = &v
  return s
}

type QueryTimeShiftParameterTemplateResponseHeader struct {
}

func (s QueryTimeShiftParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type QueryTimeShiftRuleRequest struct {
  // {"en":"rule id","zh_CN":"时移规则ID"}
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

func (s QueryTimeShiftRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleRequest) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftRuleRequest) SetRuleId(v string) *QueryTimeShiftRuleRequest {
  s.RuleId = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetDomain(v string) *QueryTimeShiftRuleRequest {
  s.Domain = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetAppName(v string) *QueryTimeShiftRuleRequest {
  s.AppName = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetStreamName(v string) *QueryTimeShiftRuleRequest {
  s.StreamName = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetTemplateId(v string) *QueryTimeShiftRuleRequest {
  s.TemplateId = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetPullDomain(v string) *QueryTimeShiftRuleRequest {
  s.PullDomain = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetPageNum(v int) *QueryTimeShiftRuleRequest {
  s.PageNum = &v
  return s
}

func (s *QueryTimeShiftRuleRequest) SetPageSize(v int) *QueryTimeShiftRuleRequest {
  s.PageSize = &v
  return s
}

type QueryTimeShiftRuleRequestHeader struct {
}

func (s QueryTimeShiftRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleRequestHeader) GoString() string {
  return s.String()
}

type QueryTimeShiftRulePaths struct {
}

func (s QueryTimeShiftRulePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRulePaths) GoString() string {
  return s.String()
}

type QueryTimeShiftRuleParameters struct {
}

func (s QueryTimeShiftRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleParameters) GoString() string {
  return s.String()
}

type QueryTimeShiftRuleResponse struct {
  // {"en":"Response status code.","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryTimeShiftRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryTimeShiftRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleResponse) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftRuleResponse) SetCode(v int) *QueryTimeShiftRuleResponse {
  s.Code = &v
  return s
}

func (s *QueryTimeShiftRuleResponse) SetMessage(v string) *QueryTimeShiftRuleResponse {
  s.Message = &v
  return s
}

func (s *QueryTimeShiftRuleResponse) SetData(v *QueryTimeShiftRuleResponseData) *QueryTimeShiftRuleResponse {
  s.Data = v
  return s
}

type QueryTimeShiftRuleResponseData struct {
  // {"en":"data size.","zh_CN":"符合查询条件总数量"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of time-shift rule data.","zh_CN":"规则列表"}
  List []*QueryTimeShiftRuleResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTimeShiftRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleResponseData) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftRuleResponseData) SetTotal(v int) *QueryTimeShiftRuleResponseData {
  s.Total = &v
  return s
}

func (s *QueryTimeShiftRuleResponseData) SetList(v []*QueryTimeShiftRuleResponseDataList) *QueryTimeShiftRuleResponseData {
  s.List = v
  return s
}

type QueryTimeShiftRuleResponseDataList struct     {
  // {"en":"The ID of the rule.","zh_CN":"时移规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"The ID of the template.","zh_CN":"模版ID"}
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
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s QueryTimeShiftRuleResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleResponseDataList) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftRuleResponseDataList) SetRuleId(v string) *QueryTimeShiftRuleResponseDataList {
  s.RuleId = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetTemplateId(v string) *QueryTimeShiftRuleResponseDataList {
  s.TemplateId = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetDomain(v string) *QueryTimeShiftRuleResponseDataList {
  s.Domain = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetAppName(v string) *QueryTimeShiftRuleResponseDataList {
  s.AppName = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetStreamName(v string) *QueryTimeShiftRuleResponseDataList {
  s.StreamName = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetStreamParams(v string) *QueryTimeShiftRuleResponseDataList {
  s.StreamParams = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetPullDomain(v string) *QueryTimeShiftRuleResponseDataList {
  s.PullDomain = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetIsEnabled(v int) *QueryTimeShiftRuleResponseDataList {
  s.IsEnabled = &v
  return s
}

func (s *QueryTimeShiftRuleResponseDataList) SetCreateTime(v string) *QueryTimeShiftRuleResponseDataList {
  s.CreateTime = &v
  return s
}

type QueryTimeShiftRuleResponseHeader struct {
}

func (s QueryTimeShiftRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftRuleResponseHeader) GoString() string {
  return s.String()
}




type AddTimeShiftRulesRequest struct {
  // {"en":"The ID of the associated time-shift template.","zh_CN":"模版ID"}
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

func (s AddTimeShiftRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesRequest) GoString() string {
  return s.String()
}

func (s *AddTimeShiftRulesRequest) SetTemplateId(v string) *AddTimeShiftRulesRequest {
  s.TemplateId = &v
  return s
}

func (s *AddTimeShiftRulesRequest) SetDomain(v string) *AddTimeShiftRulesRequest {
  s.Domain = &v
  return s
}

func (s *AddTimeShiftRulesRequest) SetAppName(v string) *AddTimeShiftRulesRequest {
  s.AppName = &v
  return s
}

func (s *AddTimeShiftRulesRequest) SetStreamName(v string) *AddTimeShiftRulesRequest {
  s.StreamName = &v
  return s
}

func (s *AddTimeShiftRulesRequest) SetStreamParams(v string) *AddTimeShiftRulesRequest {
  s.StreamParams = &v
  return s
}

func (s *AddTimeShiftRulesRequest) SetPullDomain(v string) *AddTimeShiftRulesRequest {
  s.PullDomain = &v
  return s
}

func (s *AddTimeShiftRulesRequest) SetIsEnabled(v int) *AddTimeShiftRulesRequest {
  s.IsEnabled = &v
  return s
}

type AddTimeShiftRulesRequestHeader struct {
}

func (s AddTimeShiftRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesRequestHeader) GoString() string {
  return s.String()
}

type AddTimeShiftRulesPaths struct {
}

func (s AddTimeShiftRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesPaths) GoString() string {
  return s.String()
}

type AddTimeShiftRulesParameters struct {
}

func (s AddTimeShiftRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesParameters) GoString() string {
  return s.String()
}

type AddTimeShiftRulesResponse struct {
  // {"en":"The response status code.","zh_CN":"响应状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed information or error message for the interface response.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The specific business data returned by the interface.","zh_CN":"返回数据"}
  Data *AddTimeShiftRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddTimeShiftRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesResponse) GoString() string {
  return s.String()
}

func (s *AddTimeShiftRulesResponse) SetCode(v int) *AddTimeShiftRulesResponse {
  s.Code = &v
  return s
}

func (s *AddTimeShiftRulesResponse) SetMessage(v string) *AddTimeShiftRulesResponse {
  s.Message = &v
  return s
}

func (s *AddTimeShiftRulesResponse) SetData(v *AddTimeShiftRulesResponseData) *AddTimeShiftRulesResponse {
  s.Data = v
  return s
}

type AddTimeShiftRulesResponseData struct {
  // {"en":"The ID of the newly created time-shift rule","zh_CN":"录制规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s AddTimeShiftRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesResponseData) GoString() string {
  return s.String()
}

func (s *AddTimeShiftRulesResponseData) SetRuleId(v string) *AddTimeShiftRulesResponseData {
  s.RuleId = &v
  return s
}

type AddTimeShiftRulesResponseHeader struct {
}

func (s AddTimeShiftRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTimeShiftRulesResponseHeader) GoString() string {
  return s.String()
}




type DeleteTimeShiftRulesRequest struct {
}

func (s DeleteTimeShiftRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftRulesRequest) GoString() string {
  return s.String()
}

type DeleteTimeShiftRulesRequestHeader struct {
}

func (s DeleteTimeShiftRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftRulesRequestHeader) GoString() string {
  return s.String()
}

type DeleteTimeShiftRulesPaths struct {
  // {"en":"The ID of the time-shift rule.","zh_CN":"时移规则的唯一标识符"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s DeleteTimeShiftRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftRulesPaths) GoString() string {
  return s.String()
}

func (s *DeleteTimeShiftRulesPaths) SetRuleId(v string) *DeleteTimeShiftRulesPaths {
  s.RuleId = &v
  return s
}

type DeleteTimeShiftRulesParameters struct {
}

func (s DeleteTimeShiftRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftRulesParameters) GoString() string {
  return s.String()
}

type DeleteTimeShiftRulesResponse struct {
  // {"en":"The status code of the API response.","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed message describing the outcome of the API call.","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteTimeShiftRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftRulesResponse) GoString() string {
  return s.String()
}

func (s *DeleteTimeShiftRulesResponse) SetCode(v int) *DeleteTimeShiftRulesResponse {
  s.Code = &v
  return s
}

func (s *DeleteTimeShiftRulesResponse) SetMessage(v string) *DeleteTimeShiftRulesResponse {
  s.Message = &v
  return s
}

type DeleteTimeShiftRulesResponseHeader struct {
}

func (s DeleteTimeShiftRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftRulesResponseHeader) GoString() string {
  return s.String()
}




