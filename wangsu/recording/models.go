package recording

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ModifyRecordingParameterTemplateRequest struct {
  // {"en":"Name of the recording parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"List of recording parameters.","zh_CN":"录制参数列表"}
  RecordParams []*ModifyRecordingParameterTemplateRequestRecordParams `json:"recordParams,omitempty" xml:"recordParams,omitempty" type:"Repeated"`
}

func (s ModifyRecordingParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *ModifyRecordingParameterTemplateRequest) SetTemplateName(v string) *ModifyRecordingParameterTemplateRequest {
  s.TemplateName = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequest) SetAk(v string) *ModifyRecordingParameterTemplateRequest {
  s.Ak = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequest) SetSk(v string) *ModifyRecordingParameterTemplateRequest {
  s.Sk = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequest) SetMgrUrl(v string) *ModifyRecordingParameterTemplateRequest {
  s.MgrUrl = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequest) SetNotifyUrl(v string) *ModifyRecordingParameterTemplateRequest {
  s.NotifyUrl = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequest) SetRecordParams(v []*ModifyRecordingParameterTemplateRequestRecordParams) *ModifyRecordingParameterTemplateRequest {
  s.RecordParams = v
  return s
}

type ModifyRecordingParameterTemplateRequestRecordParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
  // {"en":"Recording operation parameters.","zh_CN":"录制参数"}
  Fops *ModifyRecordingParameterTemplateRequestRecordParamsFops `json:"fops,omitempty" xml:"fops,omitempty" type:"Struct"`
}

func (s ModifyRecordingParameterTemplateRequestRecordParams) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateRequestRecordParams) GoString() string {
  return s.String()
}

func (s *ModifyRecordingParameterTemplateRequestRecordParams) SetBucketName(v string) *ModifyRecordingParameterTemplateRequestRecordParams {
  s.BucketName = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParams) SetStorageTime(v int) *ModifyRecordingParameterTemplateRequestRecordParams {
  s.StorageTime = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParams) SetFilePath(v string) *ModifyRecordingParameterTemplateRequestRecordParams {
  s.FilePath = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParams) SetFops(v *ModifyRecordingParameterTemplateRequestRecordParamsFops) *ModifyRecordingParameterTemplateRequestRecordParams {
  s.Fops = v
  return s
}

type ModifyRecordingParameterTemplateRequestRecordParamsFops struct {
  // {"en":"Format of the recording file (e.g., mp4, flv).","zh_CN":"文件格式"}
  FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
  // {"en":"Maximum file size in MB.","zh_CN":"文件大小"}
  FileSize *int `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
  // {"en":"Duration of recording segments in seconds.","zh_CN":"录制分段时长"}
  Interval *int `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"Whether to concatenate recorded segment files (0: No, 1: Yes).","zh_CN":"录制分段文件是否进行合并"}
  Concat *int `json:"concat,omitempty" xml:"concat,omitempty"`
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

func (s ModifyRecordingParameterTemplateRequestRecordParamsFops) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateRequestRecordParamsFops) GoString() string {
  return s.String()
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetFileFormat(v string) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.FileFormat = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetFileSize(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.FileSize = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetInterval(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.Interval = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetConcat(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.Concat = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetSegtime(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.Segtime = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetRemoveAudio(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.RemoveAudio = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetRemoveVideo(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.RemoveVideo = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetTimeOut(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.TimeOut = &v
  return s
}

func (s *ModifyRecordingParameterTemplateRequestRecordParamsFops) SetRetryTimes(v int) *ModifyRecordingParameterTemplateRequestRecordParamsFops {
  s.RetryTimes = &v
  return s
}

type ModifyRecordingParameterTemplateRequestHeader struct {
}

func (s ModifyRecordingParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type ModifyRecordingParameterTemplatePaths struct {
  // {"en":"ID of the template to be modified.","zh_CN":"要修改的模版ID。"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s ModifyRecordingParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *ModifyRecordingParameterTemplatePaths) SetTemplateId(v string) *ModifyRecordingParameterTemplatePaths {
  s.TemplateId = &v
  return s
}

type ModifyRecordingParameterTemplateParameters struct {
}

func (s ModifyRecordingParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateParameters) GoString() string {
  return s.String()
}

type ModifyRecordingParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifyRecordingParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *ModifyRecordingParameterTemplateResponse) SetCode(v int) *ModifyRecordingParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *ModifyRecordingParameterTemplateResponse) SetMessage(v string) *ModifyRecordingParameterTemplateResponse {
  s.Message = &v
  return s
}

type ModifyRecordingParameterTemplateResponseHeader struct {
}

func (s ModifyRecordingParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type QueryRecordingRuleRequest struct {
  // {"en":"rule id","zh_CN":"录制规则ID"}
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

func (s QueryRecordingRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleRequest) GoString() string {
  return s.String()
}

func (s *QueryRecordingRuleRequest) SetRuleId(v string) *QueryRecordingRuleRequest {
  s.RuleId = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetDomain(v string) *QueryRecordingRuleRequest {
  s.Domain = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetAppName(v string) *QueryRecordingRuleRequest {
  s.AppName = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetStreamName(v string) *QueryRecordingRuleRequest {
  s.StreamName = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetTemplateId(v string) *QueryRecordingRuleRequest {
  s.TemplateId = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetPullDomain(v string) *QueryRecordingRuleRequest {
  s.PullDomain = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetPageNum(v int) *QueryRecordingRuleRequest {
  s.PageNum = &v
  return s
}

func (s *QueryRecordingRuleRequest) SetPageSize(v int) *QueryRecordingRuleRequest {
  s.PageSize = &v
  return s
}

type QueryRecordingRuleRequestHeader struct {
}

func (s QueryRecordingRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleRequestHeader) GoString() string {
  return s.String()
}

type QueryRecordingRulePaths struct {
}

func (s QueryRecordingRulePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRulePaths) GoString() string {
  return s.String()
}

type QueryRecordingRuleParameters struct {
}

func (s QueryRecordingRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleParameters) GoString() string {
  return s.String()
}

type QueryRecordingRuleResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"返回数据"}
  Data *QueryRecordingRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryRecordingRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleResponse) GoString() string {
  return s.String()
}

func (s *QueryRecordingRuleResponse) SetCode(v int) *QueryRecordingRuleResponse {
  s.Code = &v
  return s
}

func (s *QueryRecordingRuleResponse) SetMessage(v string) *QueryRecordingRuleResponse {
  s.Message = &v
  return s
}

func (s *QueryRecordingRuleResponse) SetData(v *QueryRecordingRuleResponseData) *QueryRecordingRuleResponse {
  s.Data = v
  return s
}

type QueryRecordingRuleResponseData struct {
  // {"en":"data size","zh_CN":"符合查询条件总数量"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of recording rule data.","zh_CN":"规则列表"}
  List []*QueryRecordingRuleResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRecordingRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleResponseData) GoString() string {
  return s.String()
}

func (s *QueryRecordingRuleResponseData) SetTotal(v int) *QueryRecordingRuleResponseData {
  s.Total = &v
  return s
}

func (s *QueryRecordingRuleResponseData) SetList(v []*QueryRecordingRuleResponseDataList) *QueryRecordingRuleResponseData {
  s.List = v
  return s
}

type QueryRecordingRuleResponseDataList struct     {
  // {"en":"rule id","zh_CN":"规则ID"}
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

func (s QueryRecordingRuleResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleResponseDataList) GoString() string {
  return s.String()
}

func (s *QueryRecordingRuleResponseDataList) SetRuleId(v string) *QueryRecordingRuleResponseDataList {
  s.RuleId = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetTemplateId(v string) *QueryRecordingRuleResponseDataList {
  s.TemplateId = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetDomain(v string) *QueryRecordingRuleResponseDataList {
  s.Domain = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetAppName(v string) *QueryRecordingRuleResponseDataList {
  s.AppName = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetStreamName(v string) *QueryRecordingRuleResponseDataList {
  s.StreamName = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetStreamParams(v string) *QueryRecordingRuleResponseDataList {
  s.StreamParams = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetPullDomain(v string) *QueryRecordingRuleResponseDataList {
  s.PullDomain = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetIsEnabled(v int) *QueryRecordingRuleResponseDataList {
  s.IsEnabled = &v
  return s
}

func (s *QueryRecordingRuleResponseDataList) SetCreateTime(v int64) *QueryRecordingRuleResponseDataList {
  s.CreateTime = &v
  return s
}

type QueryRecordingRuleResponseHeader struct {
}

func (s QueryRecordingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingRuleResponseHeader) GoString() string {
  return s.String()
}




type QueryRecordingParameterTemplateRequest struct {
  // {"en":"The ID of the recording parameter template to query. If not provided, all templates under the account will be returned.","zh_CN":"要查询的录制参数模板ID。如果未提供，将返回账户下所有模板。"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QueryRecordingParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *QueryRecordingParameterTemplateRequest) SetTemplateId(v string) *QueryRecordingParameterTemplateRequest {
  s.TemplateId = &v
  return s
}

type QueryRecordingParameterTemplateRequestHeader struct {
}

func (s QueryRecordingParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type QueryRecordingParameterTemplatePaths struct {
}

func (s QueryRecordingParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplatePaths) GoString() string {
  return s.String()
}

type QueryRecordingParameterTemplateParameters struct {
}

func (s QueryRecordingParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateParameters) GoString() string {
  return s.String()
}

type QueryRecordingParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryRecordingParameterTemplateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryRecordingParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *QueryRecordingParameterTemplateResponse) SetCode(v int) *QueryRecordingParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponse) SetMessage(v string) *QueryRecordingParameterTemplateResponse {
  s.Message = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponse) SetData(v *QueryRecordingParameterTemplateResponseData) *QueryRecordingParameterTemplateResponse {
  s.Data = v
  return s
}

type QueryRecordingParameterTemplateResponseData struct {
  // {"en":"data size.","zh_CN":"总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of screenshot parameter template data.","zh_CN":"截图参数模板数据列表"}
  List []*QueryRecordingParameterTemplateResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRecordingParameterTemplateResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateResponseData) GoString() string {
  return s.String()
}

func (s *QueryRecordingParameterTemplateResponseData) SetTotal(v int) *QueryRecordingParameterTemplateResponseData {
  s.Total = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseData) SetList(v []*QueryRecordingParameterTemplateResponseDataList) *QueryRecordingParameterTemplateResponseData {
  s.List = v
  return s
}

type QueryRecordingParameterTemplateResponseDataList struct     {
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
  RecordParams []*QueryRecordingParameterTemplateResponseDataListRecordParams `json:"recordParams,omitempty" xml:"recordParams,omitempty" require:"true" type:"Repeated"`
  // {"en":"The unique identifier of the screenshot parameter template.","zh_CN":"模版id"}
  TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s QueryRecordingParameterTemplateResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateResponseDataList) GoString() string {
  return s.String()
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetRegion(v string) *QueryRecordingParameterTemplateResponseDataList {
  s.Region = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetTemplateName(v string) *QueryRecordingParameterTemplateResponseDataList {
  s.TemplateName = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetAk(v string) *QueryRecordingParameterTemplateResponseDataList {
  s.Ak = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetSk(v string) *QueryRecordingParameterTemplateResponseDataList {
  s.Sk = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetMgrUrl(v string) *QueryRecordingParameterTemplateResponseDataList {
  s.MgrUrl = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetNotifyUrl(v string) *QueryRecordingParameterTemplateResponseDataList {
  s.NotifyUrl = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetRecordParams(v []*QueryRecordingParameterTemplateResponseDataListRecordParams) *QueryRecordingParameterTemplateResponseDataList {
  s.RecordParams = v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataList) SetTemplateId(v int64) *QueryRecordingParameterTemplateResponseDataList {
  s.TemplateId = &v
  return s
}

type QueryRecordingParameterTemplateResponseDataListRecordParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty" require:"true"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty" require:"true"`
  // {"en":"Screenshot operation parameters.","zh_CN":"截图参数"}
  Fops *QueryRecordingParameterTemplateResponseDataListRecordParamsFops `json:"fops,omitempty" xml:"fops,omitempty" require:"true" type:"Struct"`
}

func (s QueryRecordingParameterTemplateResponseDataListRecordParams) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateResponseDataListRecordParams) GoString() string {
  return s.String()
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParams) SetBucketName(v string) *QueryRecordingParameterTemplateResponseDataListRecordParams {
  s.BucketName = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParams) SetStorageTime(v int) *QueryRecordingParameterTemplateResponseDataListRecordParams {
  s.StorageTime = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParams) SetFilePath(v string) *QueryRecordingParameterTemplateResponseDataListRecordParams {
  s.FilePath = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParams) SetFops(v *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) *QueryRecordingParameterTemplateResponseDataListRecordParams {
  s.Fops = v
  return s
}

type QueryRecordingParameterTemplateResponseDataListRecordParamsFops struct {
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

func (s QueryRecordingParameterTemplateResponseDataListRecordParamsFops) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateResponseDataListRecordParamsFops) GoString() string {
  return s.String()
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetFileFormat(v string) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.FileFormat = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetInterval(v int) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.Interval = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetWidth(v int) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.Width = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetHeight(v int) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.Height = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetLongside(v int) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.Longside = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetTimeOut(v int) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.TimeOut = &v
  return s
}

func (s *QueryRecordingParameterTemplateResponseDataListRecordParamsFops) SetRetryTimes(v int) *QueryRecordingParameterTemplateResponseDataListRecordParamsFops {
  s.RetryTimes = &v
  return s
}

type QueryRecordingParameterTemplateResponseHeader struct {
}

func (s QueryRecordingParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordingParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type DeleteRecordingRulesRequest struct {
}

func (s DeleteRecordingRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingRulesRequest) GoString() string {
  return s.String()
}

type DeleteRecordingRulesRequestHeader struct {
}

func (s DeleteRecordingRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingRulesRequestHeader) GoString() string {
  return s.String()
}

type DeleteRecordingRulesPaths struct {
  // {"en":"The unique identifier ID of the recording rule to be deleted","zh_CN":"规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s DeleteRecordingRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingRulesPaths) GoString() string {
  return s.String()
}

func (s *DeleteRecordingRulesPaths) SetRuleId(v string) *DeleteRecordingRulesPaths {
  s.RuleId = &v
  return s
}

type DeleteRecordingRulesParameters struct {
}

func (s DeleteRecordingRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingRulesParameters) GoString() string {
  return s.String()
}

type DeleteRecordingRulesResponse struct {
  // {"en":"The status code of the API response.","zh_CN":"接口响应的状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed message describing the outcome of the API call.","zh_CN":"描述API调用结果的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteRecordingRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingRulesResponse) GoString() string {
  return s.String()
}

func (s *DeleteRecordingRulesResponse) SetCode(v int) *DeleteRecordingRulesResponse {
  s.Code = &v
  return s
}

func (s *DeleteRecordingRulesResponse) SetMessage(v string) *DeleteRecordingRulesResponse {
  s.Message = &v
  return s
}

type DeleteRecordingRulesResponseHeader struct {
}

func (s DeleteRecordingRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingRulesResponseHeader) GoString() string {
  return s.String()
}




type AddRecordingRulesRequest struct {
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
  // {"en":"Pull stream domain","zh_CN":"拉流域名，禁止传空字符串"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"Whether the rule is enabled. 0 for disabled, 1 for enabled. Defaults to 1.","zh_CN":"0：不启用，1：启用 默认为1"}
  IsEnabled *int `json:"isEnabled,omitempty" xml:"isEnabled,omitempty"`
}

func (s AddRecordingRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesRequest) GoString() string {
  return s.String()
}

func (s *AddRecordingRulesRequest) SetTemplateId(v string) *AddRecordingRulesRequest {
  s.TemplateId = &v
  return s
}

func (s *AddRecordingRulesRequest) SetDomain(v string) *AddRecordingRulesRequest {
  s.Domain = &v
  return s
}

func (s *AddRecordingRulesRequest) SetAppName(v string) *AddRecordingRulesRequest {
  s.AppName = &v
  return s
}

func (s *AddRecordingRulesRequest) SetStreamName(v string) *AddRecordingRulesRequest {
  s.StreamName = &v
  return s
}

func (s *AddRecordingRulesRequest) SetStreamParams(v string) *AddRecordingRulesRequest {
  s.StreamParams = &v
  return s
}

func (s *AddRecordingRulesRequest) SetPullDomain(v string) *AddRecordingRulesRequest {
  s.PullDomain = &v
  return s
}

func (s *AddRecordingRulesRequest) SetIsEnabled(v int) *AddRecordingRulesRequest {
  s.IsEnabled = &v
  return s
}

type AddRecordingRulesRequestHeader struct {
}

func (s AddRecordingRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesRequestHeader) GoString() string {
  return s.String()
}

type AddRecordingRulesPaths struct {
}

func (s AddRecordingRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesPaths) GoString() string {
  return s.String()
}

type AddRecordingRulesParameters struct {
}

func (s AddRecordingRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesParameters) GoString() string {
  return s.String()
}

type AddRecordingRulesResponse struct {
  // {"en":"code","zh_CN":"结果状态码，200为成功"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回数据"}
  Data *AddRecordingRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddRecordingRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesResponse) GoString() string {
  return s.String()
}

func (s *AddRecordingRulesResponse) SetCode(v int) *AddRecordingRulesResponse {
  s.Code = &v
  return s
}

func (s *AddRecordingRulesResponse) SetMessage(v string) *AddRecordingRulesResponse {
  s.Message = &v
  return s
}

func (s *AddRecordingRulesResponse) SetData(v *AddRecordingRulesResponseData) *AddRecordingRulesResponse {
  s.Data = v
  return s
}

type AddRecordingRulesResponseData struct {
  // {"en":"rule id","zh_CN":"录制规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s AddRecordingRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesResponseData) GoString() string {
  return s.String()
}

func (s *AddRecordingRulesResponseData) SetRuleId(v string) *AddRecordingRulesResponseData {
  s.RuleId = &v
  return s
}

type AddRecordingRulesResponseHeader struct {
}

func (s AddRecordingRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingRulesResponseHeader) GoString() string {
  return s.String()
}




type AddRecordingParameterTemplateRequest struct {
  // {"en":"Geographical region of the cloud storage.","zh_CN":"区域"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"Name of the recording parameter template.","zh_CN":"模版名称"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty" require:"true"`
  // {"en":"Access Key for cloud storage access.","zh_CN":"云存储访问的Access Key"}
  Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
  // {"en":"Secret Key for cloud storage access.","zh_CN":"云存储访问的Secret Key"}
  Sk *string `json:"sk,omitempty" xml:"sk,omitempty"`
  // {"en":"Cloud storage space management domain name (Management URL).","zh_CN":"云存储空间管理域名"}
  MgrUrl *string `json:"mgrUrl,omitempty" xml:"mgrUrl,omitempty" require:"true"`
  // {"en":"Callback notification URL.","zh_CN":"回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"List of recording parameters.","zh_CN":"录制参数列表"}
  RecordParams []*AddRecordingParameterTemplateRequestRecordParams `json:"recordParams,omitempty" xml:"recordParams,omitempty" require:"true" type:"Repeated"`
}

func (s AddRecordingParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *AddRecordingParameterTemplateRequest) SetRegion(v string) *AddRecordingParameterTemplateRequest {
  s.Region = &v
  return s
}

func (s *AddRecordingParameterTemplateRequest) SetTemplateName(v string) *AddRecordingParameterTemplateRequest {
  s.TemplateName = &v
  return s
}

func (s *AddRecordingParameterTemplateRequest) SetAk(v string) *AddRecordingParameterTemplateRequest {
  s.Ak = &v
  return s
}

func (s *AddRecordingParameterTemplateRequest) SetSk(v string) *AddRecordingParameterTemplateRequest {
  s.Sk = &v
  return s
}

func (s *AddRecordingParameterTemplateRequest) SetMgrUrl(v string) *AddRecordingParameterTemplateRequest {
  s.MgrUrl = &v
  return s
}

func (s *AddRecordingParameterTemplateRequest) SetNotifyUrl(v string) *AddRecordingParameterTemplateRequest {
  s.NotifyUrl = &v
  return s
}

func (s *AddRecordingParameterTemplateRequest) SetRecordParams(v []*AddRecordingParameterTemplateRequestRecordParams) *AddRecordingParameterTemplateRequest {
  s.RecordParams = v
  return s
}

type AddRecordingParameterTemplateRequestRecordParams struct     {
  // {"en":"Name of the storage bucket.","zh_CN":"存储空间名"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
  // {"en":"Storage duration in seconds.","zh_CN":"存储时间"}
  StorageTime *int `json:"storageTime,omitempty" xml:"storageTime,omitempty"`
  // {"en":"Name of the stored file.","zh_CN":"存储文件名称"}
  FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
  // {"en":"Recording operation parameters.","zh_CN":"录制参数"}
  Fops *AddRecordingParameterTemplateRequestRecordParamsFops `json:"fops,omitempty" xml:"fops,omitempty" type:"Struct"`
}

func (s AddRecordingParameterTemplateRequestRecordParams) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateRequestRecordParams) GoString() string {
  return s.String()
}

func (s *AddRecordingParameterTemplateRequestRecordParams) SetBucketName(v string) *AddRecordingParameterTemplateRequestRecordParams {
  s.BucketName = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParams) SetStorageTime(v int) *AddRecordingParameterTemplateRequestRecordParams {
  s.StorageTime = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParams) SetFilePath(v string) *AddRecordingParameterTemplateRequestRecordParams {
  s.FilePath = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParams) SetFops(v *AddRecordingParameterTemplateRequestRecordParamsFops) *AddRecordingParameterTemplateRequestRecordParams {
  s.Fops = v
  return s
}

type AddRecordingParameterTemplateRequestRecordParamsFops struct {
  // {"en":"Format of the recording file (e.g., mp4, flv).","zh_CN":"文件格式"}
  FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty" require:"true"`
  // {"en":"Maximum file size in MB.","zh_CN":"文件大小"}
  FileSize *int `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
  // {"en":"Duration of recording segments in seconds.","zh_CN":"录制分段时长"}
  Interval *int `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"Whether to concatenate recorded segment files (0: No, 1: Yes).","zh_CN":"录制分段文件是否进行合并"}
  Concat *int `json:"concat,omitempty" xml:"concat,omitempty"`
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

func (s AddRecordingParameterTemplateRequestRecordParamsFops) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateRequestRecordParamsFops) GoString() string {
  return s.String()
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetFileFormat(v string) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.FileFormat = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetFileSize(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.FileSize = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetInterval(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.Interval = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetConcat(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.Concat = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetSegtime(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.Segtime = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetRemoveAudio(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.RemoveAudio = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetRemoveVideo(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.RemoveVideo = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetTimeOut(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.TimeOut = &v
  return s
}

func (s *AddRecordingParameterTemplateRequestRecordParamsFops) SetRetryTimes(v int) *AddRecordingParameterTemplateRequestRecordParamsFops {
  s.RetryTimes = &v
  return s
}

type AddRecordingParameterTemplateRequestHeader struct {
}

func (s AddRecordingParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type AddRecordingParameterTemplatePaths struct {
}

func (s AddRecordingParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplatePaths) GoString() string {
  return s.String()
}

type AddRecordingParameterTemplateParameters struct {
}

func (s AddRecordingParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateParameters) GoString() string {
  return s.String()
}

type AddRecordingParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data payload.","zh_CN":"响应数据"}
  Data *AddRecordingParameterTemplateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddRecordingParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *AddRecordingParameterTemplateResponse) SetCode(v int) *AddRecordingParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *AddRecordingParameterTemplateResponse) SetMessage(v string) *AddRecordingParameterTemplateResponse {
  s.Message = &v
  return s
}

func (s *AddRecordingParameterTemplateResponse) SetData(v *AddRecordingParameterTemplateResponseData) *AddRecordingParameterTemplateResponse {
  s.Data = v
  return s
}

type AddRecordingParameterTemplateResponseData struct {
  // {"en":"ID of the created template.","zh_CN":"模版id"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s AddRecordingParameterTemplateResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateResponseData) GoString() string {
  return s.String()
}

func (s *AddRecordingParameterTemplateResponseData) SetTemplateId(v string) *AddRecordingParameterTemplateResponseData {
  s.TemplateId = &v
  return s
}

type AddRecordingParameterTemplateResponseHeader struct {
}

func (s AddRecordingParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddRecordingParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type ModifyRecordingRulesRequest struct {
  // {"en":"The unique identifier of the template.","zh_CN":"模版ID,禁止传空字符串。  不传：不修改原来的值"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
  // {"en":"domain","zh_CN":"推流域名，如果是推拉架构，推流域名必填，不能传空字符串  不传：不修改原来的值  空字符串：删除配置"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Publishing point name","zh_CN":"发布点  不传：不修改原来的值  空字符串：删除配置"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
  // {"en":"stream name","zh_CN":"流名  不传：不修改原来的值  空字符串：删除配置"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
  // {"en":"Extension parameters for the stream name.","zh_CN":"流名扩展参数。  不传：不修改原来的值  空字符串：删除配置"}
  StreamParams *string `json:"streamParams,omitempty" xml:"streamParams,omitempty"`
  // {"en":"Whether the rule is enabled. 0 for disabled, 1 for enabled. Defaults to 1.","zh_CN":"0：不启用，1：启用 默认为1"}
  IsEnabled *int `json:"isEnabled,omitempty" xml:"isEnabled,omitempty"`
  // {"en":"pull domain","zh_CN":"拉流域名,禁止传空字符串  不传：不修改原来的值"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty"`
}

func (s ModifyRecordingRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingRulesRequest) GoString() string {
  return s.String()
}

func (s *ModifyRecordingRulesRequest) SetTemplateId(v string) *ModifyRecordingRulesRequest {
  s.TemplateId = &v
  return s
}

func (s *ModifyRecordingRulesRequest) SetDomain(v string) *ModifyRecordingRulesRequest {
  s.Domain = &v
  return s
}

func (s *ModifyRecordingRulesRequest) SetAppName(v string) *ModifyRecordingRulesRequest {
  s.AppName = &v
  return s
}

func (s *ModifyRecordingRulesRequest) SetStreamName(v string) *ModifyRecordingRulesRequest {
  s.StreamName = &v
  return s
}

func (s *ModifyRecordingRulesRequest) SetStreamParams(v string) *ModifyRecordingRulesRequest {
  s.StreamParams = &v
  return s
}

func (s *ModifyRecordingRulesRequest) SetIsEnabled(v int) *ModifyRecordingRulesRequest {
  s.IsEnabled = &v
  return s
}

func (s *ModifyRecordingRulesRequest) SetPullDomain(v string) *ModifyRecordingRulesRequest {
  s.PullDomain = &v
  return s
}

type ModifyRecordingRulesRequestHeader struct {
}

func (s ModifyRecordingRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingRulesRequestHeader) GoString() string {
  return s.String()
}

type ModifyRecordingRulesPaths struct {
  // {"en":"The unique identifier of the recording rule.","zh_CN":"录制规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s ModifyRecordingRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingRulesPaths) GoString() string {
  return s.String()
}

func (s *ModifyRecordingRulesPaths) SetRuleId(v string) *ModifyRecordingRulesPaths {
  s.RuleId = &v
  return s
}

type ModifyRecordingRulesParameters struct {
}

func (s ModifyRecordingRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingRulesParameters) GoString() string {
  return s.String()
}

type ModifyRecordingRulesResponse struct {
  // {"en":"The status code of the API response.","zh_CN":"接口响应的状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed message describing the outcome of the API call.","zh_CN":"描述API调用结果的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifyRecordingRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingRulesResponse) GoString() string {
  return s.String()
}

func (s *ModifyRecordingRulesResponse) SetCode(v int) *ModifyRecordingRulesResponse {
  s.Code = &v
  return s
}

func (s *ModifyRecordingRulesResponse) SetMessage(v string) *ModifyRecordingRulesResponse {
  s.Message = &v
  return s
}

type ModifyRecordingRulesResponseHeader struct {
}

func (s ModifyRecordingRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordingRulesResponseHeader) GoString() string {
  return s.String()
}




type LiveVideoConcatRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"The service ID must be unique. You are advised to use a 32-bit UUID and the value can be a string of up to 32 characters", "zh_CN":"业务ID，需用户自己控制唯一性，建议使用32位UUID，并且最长为32位字符串"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"Start time, unix time stamp. Default is the first broadcast start time", "zh_CN":"开始时间，unix时间戳，默认为第一次直播开始时间"}
  Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
  // {"en":"End time, unix timestamp, default current time. End time Future time is prohibited", "zh_CN":"结束时间，unix时间戳，默认当前时间。结束时间禁止填未来时间"}
  End *int64 `json:"end,omitempty" xml:"end,omitempty"`
  // {"en":"File name. If it is empty, the system automatically generates a file name (stream name _ start time _ end time).", "zh_CN":"文件名。若为空则系统自动生成一个文件名（流名_开始时间_结束时间）"}
  Fname *string `json:"fname,omitempty" xml:"fname,omitempty"`
  // {"en":"The optional file format is:
  // 
  // flv: FLV format, which combines multiple recorded videos into a single flv file. Default format
  // 
  // mp4: MP4 format, which combines multiple recorded videos into a single mp4 file.
  // 
  // Format not supported:
  // 
  // m3u8: indicates the HLS format", "zh_CN":"文件格式，可选文件格式为：
  // 
  // flv：FLV格式，将多个录制视频合并成单个flv文件。默认格式
  // 
  // mp4：MP4格式，将多个录制视频合并成单个mp4文件。
  // 
  // 不支持格式：
  // 
  // m3u8：HLS格式"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
  // {"en":"Callback address. When the task is complete, the callback address is not specified. If the address is not specified, the callback is not performed", "zh_CN":"回调地址。完成任务后回调通知地址，不指定表示不做回调"}
  Notify *string `json:"notify,omitempty" xml:"notify,omitempty"`
  // {"en":"Start forced transcoding, 1: start. 0: Do not start, the default is 0", "zh_CN":"启动强制转码，1:启动。0:不启动 默认是0"}
  EnableVideoTranscode *int32 `json:"enableVideoTranscode,omitempty" xml:"enableVideoTranscode,omitempty"`
}

func (s LiveVideoConcatRequest) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatRequest) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatRequest) SetPullId(v string) *LiveVideoConcatRequest {
  s.PullId = &v
  return s
}

func (s *LiveVideoConcatRequest) SetTransNo(v string) *LiveVideoConcatRequest {
  s.TransNo = &v
  return s
}

func (s *LiveVideoConcatRequest) SetStart(v int64) *LiveVideoConcatRequest {
  s.Start = &v
  return s
}

func (s *LiveVideoConcatRequest) SetEnd(v int64) *LiveVideoConcatRequest {
  s.End = &v
  return s
}

func (s *LiveVideoConcatRequest) SetFname(v string) *LiveVideoConcatRequest {
  s.Fname = &v
  return s
}

func (s *LiveVideoConcatRequest) SetSuffix(v string) *LiveVideoConcatRequest {
  s.Suffix = &v
  return s
}

func (s *LiveVideoConcatRequest) SetNotify(v string) *LiveVideoConcatRequest {
  s.Notify = &v
  return s
}

func (s *LiveVideoConcatRequest) SetEnableVideoTranscode(v int32) *LiveVideoConcatRequest {
  s.EnableVideoTranscode = &v
  return s
}

type LiveVideoConcatResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational infomation", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  LiveVideoConcatData *LiveVideoConcatData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LiveVideoConcatResponse) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatResponse) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatResponse) SetCode(v int32) *LiveVideoConcatResponse {
  s.Code = &v
  return s
}

func (s *LiveVideoConcatResponse) SetMessage(v string) *LiveVideoConcatResponse {
  s.Message = &v
  return s
}

func (s *LiveVideoConcatResponse) SetData(v *LiveVideoConcatData) *LiveVideoConcatResponse {
  s.LiveVideoConcatData = v
  return s
}

type LiveVideoConcatData struct {
  // {"en":"transNo", "zh_CN":"业务ID"}
  TranNo *string `json:"tranNo,omitempty" xml:"tranNo,omitempty" require:"true"`
}

func (s LiveVideoConcatData) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatData) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatData) SetTranNo(v string) *LiveVideoConcatData {
  s.TranNo = &v
  return s
}

type LiveVideoConcatPaths struct {
}

func (s LiveVideoConcatPaths) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatPaths) GoString() string {
  return s.String()
}

type LiveVideoConcatParameters struct {
}

func (s LiveVideoConcatParameters) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatParameters) GoString() string {
  return s.String()
}

type LiveVideoConcatRequestHeader struct {
}

func (s LiveVideoConcatRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatRequestHeader) GoString() string {
  return s.String()
}

type LiveVideoConcatResponseHeader struct {
}

func (s LiveVideoConcatResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatResponseHeader) GoString() string {
  return s.String()
}




type DeleteRecordingParameterTemplateRequest struct {
}

func (s DeleteRecordingParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingParameterTemplateRequest) GoString() string {
  return s.String()
}

type DeleteRecordingParameterTemplateRequestHeader struct {
}

func (s DeleteRecordingParameterTemplateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingParameterTemplateRequestHeader) GoString() string {
  return s.String()
}

type DeleteRecordingParameterTemplatePaths struct {
  // {"en":"The unique identifier of the template to be deleted.","zh_CN":"模版id"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s DeleteRecordingParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *DeleteRecordingParameterTemplatePaths) SetTemplateId(v string) *DeleteRecordingParameterTemplatePaths {
  s.TemplateId = &v
  return s
}

type DeleteRecordingParameterTemplateParameters struct {
}

func (s DeleteRecordingParameterTemplateParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingParameterTemplateParameters) GoString() string {
  return s.String()
}

type DeleteRecordingParameterTemplateResponse struct {
  // {"en":"Response status code.","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message indicating success or failure.","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteRecordingParameterTemplateResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingParameterTemplateResponse) GoString() string {
  return s.String()
}

func (s *DeleteRecordingParameterTemplateResponse) SetCode(v int) *DeleteRecordingParameterTemplateResponse {
  s.Code = &v
  return s
}

func (s *DeleteRecordingParameterTemplateResponse) SetMessage(v string) *DeleteRecordingParameterTemplateResponse {
  s.Message = &v
  return s
}

type DeleteRecordingParameterTemplateResponseHeader struct {
}

func (s DeleteRecordingParameterTemplateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordingParameterTemplateResponseHeader) GoString() string {
  return s.String()
}




type GetRecordTaskListQueryRequest struct {
  // {"en":"Channel pull ID. Multiple streaming IDs are separated by ,; if not filled in, recording tasks for all channels will be returned by default", "zh_CN":"频道拉流ID。多个拉流id用,隔开；未填写默认返回所有频道的录制任务"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty"`
  // {"en":"Business ID needs to be uniquely controlled by the user. Use , to separate multiple ones.", "zh_CN":"业务ID，需用户自己控制唯一性。多个用,隔开。"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"The status of the recording task. If not filled in, all will be queried. 0 has not started, 1 has started, 2 has ended normally, and 3 has terminated.", "zh_CN":"录制任务的状态， 不填则查询所有，0未开始，1已开始，2正常结束，3终止"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"List order, value range: <br>
  // 0(in descending order by creation time)<br>
  // 1(in ascending order of creation time) The default value is 0", "zh_CN":"列表排列顺序，取值范围 ：<br>
  // 0(按创建时间降序排列)<br>
  // 1(按创建时间升序排列)默认为0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"The page number in the task paging list starts from 1; the default is 1", "zh_CN":"task分页列表中第几页，从1开始取值；默认1"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"The average number of channels per page, the default is 10, the value is between 1-50.", "zh_CN":"平均每页频道数量，默认为10，取值在1-50之间"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Query starting time, the time format is, 2016-01-01 12:00:00; used to query recording tasks according to the creation time period; if not filled in, query all queries all", "zh_CN":"查询起始时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询录制任务；如果不填则查询所有查询所有"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"Query deadline, the time format is, 2016-01-01 12:00:00; used to query recording tasks according to the creation time period, which is less than the current query time;. If not filled in, query all query all", "zh_CN":"查询截止时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询录制任务，小于当前查询时间；。如果不填则查询所有查询所有"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s GetRecordTaskListQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRequest) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryRequest) SetPullId(v string) *GetRecordTaskListQueryRequest {
  s.PullId = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetTransNo(v string) *GetRecordTaskListQueryRequest {
  s.TransNo = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetStatus(v int32) *GetRecordTaskListQueryRequest {
  s.Status = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetListOrder(v int32) *GetRecordTaskListQueryRequest {
  s.ListOrder = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetPageIndex(v int32) *GetRecordTaskListQueryRequest {
  s.PageIndex = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetPageSize(v int32) *GetRecordTaskListQueryRequest {
  s.PageSize = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetStartTime(v string) *GetRecordTaskListQueryRequest {
  s.StartTime = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetEndTime(v string) *GetRecordTaskListQueryRequest {
  s.EndTime = &v
  return s
}

type GetRecordTaskListQueryResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *GetRecordTaskListQueryRecordTask `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetRecordTaskListQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryResponse) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryResponse) SetCode(v int32) *GetRecordTaskListQueryResponse {
  s.Code = &v
  return s
}

func (s *GetRecordTaskListQueryResponse) SetMessage(v string) *GetRecordTaskListQueryResponse {
  s.Message = &v
  return s
}

func (s *GetRecordTaskListQueryResponse) SetData(v *GetRecordTaskListQueryRecordTask) *GetRecordTaskListQueryResponse {
  s.Data = v
  return s
}

type GetRecordTaskListQueryRecordTask struct {
  // {"en":"The number of records of the recording task list information currently returned. Note that the number of records returned here is only the number of records on the current page.", "zh_CN":"当前返回的录制任务列表信息的记录数，注意这里返回的记录数只是当前页的记录数"}
  RecordTaskTotal *int32 `json:"recordTaskTotal,omitempty" xml:"recordTaskTotal,omitempty" require:"true"`
  // {"en":"Recording task list", "zh_CN":"录制任务列表"}
  RecordTaskQueryResponseList []*GetRecordTaskListQueryRecordTaskQuery `json:"recordTaskQueryResponseList,omitempty" xml:"recordTaskQueryResponseList,omitempty" require:"true" type:"Repeated"`
}

func (s GetRecordTaskListQueryRecordTask) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRecordTask) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryRecordTask) SetRecordTaskTotal(v int32) *GetRecordTaskListQueryRecordTask {
  s.RecordTaskTotal = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTask) SetRecordTaskQueryResponseList(v []*GetRecordTaskListQueryRecordTaskQuery) *GetRecordTaskListQueryRecordTask {
  s.RecordTaskQueryResponseList = v
  return s
}

type GetRecordTaskListQueryRecordTaskQuery struct {
  // {"en":"Task Id", "zh_CN":"直播录制任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"Channel streaming ID", "zh_CN":"频道拉流ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"The status of the recording task, 0 has not started, 1 has started, 2 has ended normally, and 3 has terminated", "zh_CN":"录制任务的状态， 0未开始，1已开始，2正常结束，3终止"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Custom task number passed in by the customer", "zh_CN":"客户传入的自定义任务编号"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"Recording file format", "zh_CN":"录制文件的格式"}
  FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty" require:"true"`
  // {"en":"Whether to merge into one video", "zh_CN":"是否合并成一个视频"}
  IsConcat *bool `json:"isConcat,omitempty" xml:"isConcat,omitempty" require:"true"`
  // {"en":"Task start time", "zh_CN":"任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Task end time", "zh_CN":"任务结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s GetRecordTaskListQueryRecordTaskQuery) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRecordTaskQuery) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetTaskId(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.TaskId = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetPullId(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.PullId = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetStatus(v int32) *GetRecordTaskListQueryRecordTaskQuery {
  s.Status = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetTransNo(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.TransNo = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetFileType(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.FileType = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetIsConcat(v bool) *GetRecordTaskListQueryRecordTaskQuery {
  s.IsConcat = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetStartTime(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.StartTime = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetEndTime(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.EndTime = &v
  return s
}

type GetRecordTaskListQueryPaths struct {
}

func (s GetRecordTaskListQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryPaths) GoString() string {
  return s.String()
}

type GetRecordTaskListQueryParameters struct {
}

func (s GetRecordTaskListQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryParameters) GoString() string {
  return s.String()
}

type GetRecordTaskListQueryRequestHeader struct {
}

func (s GetRecordTaskListQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRequestHeader) GoString() string {
  return s.String()
}

type GetRecordTaskListQueryResponseHeader struct {
}

func (s GetRecordTaskListQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryResponseHeader) GoString() string {
  return s.String()
}




type StartRealTimeRecordRequest struct {
  // {"en":"The ID of the recording template to use","zh_CN":"模版ID"}
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
  // {"en":"The pull domain of the live stream","zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"The application name or publishing point for the stream","zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {"en":"The name of the live stream to be recorded","zh_CN":"流名，支持多个流名，多个流名用英文逗号分隔。最多5个  示例：stream1,stream2,stream3"}
  StreamNames *string `json:"streamNames,omitempty" xml:"streamNames,omitempty" require:"true"`
}

func (s StartRealTimeRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordRequest) GoString() string {
  return s.String()
}

func (s *StartRealTimeRecordRequest) SetTemplateId(v string) *StartRealTimeRecordRequest {
  s.TemplateId = &v
  return s
}

func (s *StartRealTimeRecordRequest) SetPullDomain(v string) *StartRealTimeRecordRequest {
  s.PullDomain = &v
  return s
}

func (s *StartRealTimeRecordRequest) SetAppName(v string) *StartRealTimeRecordRequest {
  s.AppName = &v
  return s
}

func (s *StartRealTimeRecordRequest) SetStreamNames(v string) *StartRealTimeRecordRequest {
  s.StreamNames = &v
  return s
}

type StartRealTimeRecordRequestHeader struct {
}

func (s StartRealTimeRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordRequestHeader) GoString() string {
  return s.String()
}

type StartRealTimeRecordPaths struct {
}

func (s StartRealTimeRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordPaths) GoString() string {
  return s.String()
}

type StartRealTimeRecordParameters struct {
}

func (s StartRealTimeRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordParameters) GoString() string {
  return s.String()
}

type StartRealTimeRecordResponse struct {
  // {"en":"The overall response code of the API call","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The overall response message of the API call","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The detailed response data containing recording statuses","zh_CN":"返回数据"}
  Data []*StartRealTimeRecordResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s StartRealTimeRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordResponse) GoString() string {
  return s.String()
}

func (s *StartRealTimeRecordResponse) SetCode(v int) *StartRealTimeRecordResponse {
  s.Code = &v
  return s
}

func (s *StartRealTimeRecordResponse) SetMessage(v string) *StartRealTimeRecordResponse {
  s.Message = &v
  return s
}

func (s *StartRealTimeRecordResponse) SetData(v []*StartRealTimeRecordResponseData) *StartRealTimeRecordResponse {
  s.Data = v
  return s
}

type StartRealTimeRecordResponseData struct     {
  // {"en":"Stream-level response code","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Stream-level response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The name of the stream this status refers to","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty" require:"true"`
  // {"en":"The unique ID of the real-time recording task","zh_CN":"录制任务的id"}
  PersistentId *string `json:"persistentId,omitempty" xml:"persistentId,omitempty" require:"true"`
}

func (s StartRealTimeRecordResponseData) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordResponseData) GoString() string {
  return s.String()
}

func (s *StartRealTimeRecordResponseData) SetCode(v int) *StartRealTimeRecordResponseData {
  s.Code = &v
  return s
}

func (s *StartRealTimeRecordResponseData) SetMessage(v string) *StartRealTimeRecordResponseData {
  s.Message = &v
  return s
}

func (s *StartRealTimeRecordResponseData) SetStreamName(v string) *StartRealTimeRecordResponseData {
  s.StreamName = &v
  return s
}

func (s *StartRealTimeRecordResponseData) SetPersistentId(v string) *StartRealTimeRecordResponseData {
  s.PersistentId = &v
  return s
}

type StartRealTimeRecordResponseHeader struct {
}

func (s StartRealTimeRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StartRealTimeRecordResponseHeader) GoString() string {
  return s.String()
}




type StopRealTimeRecordRequest struct {
}

func (s StopRealTimeRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeRecordRequest) GoString() string {
  return s.String()
}

type StopRealTimeRecordRequestHeader struct {
}

func (s StopRealTimeRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeRecordRequestHeader) GoString() string {
  return s.String()
}

type StopRealTimeRecordPaths struct {
  // {"en":"The unique identifier for the recording task.","zh_CN":"录制任务的唯一标识符"}
  PersistentId *string `json:"persistentId,omitempty" xml:"persistentId,omitempty" require:"true"`
}

func (s StopRealTimeRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeRecordPaths) GoString() string {
  return s.String()
}

func (s *StopRealTimeRecordPaths) SetPersistentId(v string) *StopRealTimeRecordPaths {
  s.PersistentId = &v
  return s
}

type StopRealTimeRecordParameters struct {
}

func (s StopRealTimeRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeRecordParameters) GoString() string {
  return s.String()
}

type StopRealTimeRecordResponse struct {
  // {"en":"The overall response code of the API call","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The overall response message of the API call","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s StopRealTimeRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeRecordResponse) GoString() string {
  return s.String()
}

func (s *StopRealTimeRecordResponse) SetCode(v int) *StopRealTimeRecordResponse {
  s.Code = &v
  return s
}

func (s *StopRealTimeRecordResponse) SetMessage(v string) *StopRealTimeRecordResponse {
  s.Message = &v
  return s
}

type StopRealTimeRecordResponseHeader struct {
}

func (s StopRealTimeRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StopRealTimeRecordResponseHeader) GoString() string {
  return s.String()
}




type LiveVideoConcatQueryRequest struct {
  // {"en":"Service ID, service ID and channel pull id cannot be empty at the same time", "zh_CN":"业务ID，业务ID和频道拉流id不可以同时为空"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"Channel pull id, service ID, and channel pull id cannot be empty at the same time", "zh_CN":"频道拉流id，业务ID和频道拉流id不可以同时为空"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty"`
  // {"en":"List order, value range:
  // 0(in descending order of creation time)
  // 1(in ascending order of creation time)
  // Default is 0", "zh_CN":"列表排列顺序，取值范围 ：
  // 0(按创建时间降序排列)
  // 1(按创建时间升序排列)
  // 默认为0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"On the page of the scenario list, the value starts from 1. The default value is 1. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"取场景列表第几页，从1开始取值,默认为1。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average Number of scenarios per page. The value ranges from 1 to 50. The default value is 10. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"平均每页场景数量，取值范围1-50，默认为10。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s LiveVideoConcatQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryRequest) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryRequest) SetTransNo(v string) *LiveVideoConcatQueryRequest {
  s.TransNo = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetPullId(v string) *LiveVideoConcatQueryRequest {
  s.PullId = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetListOrder(v int32) *LiveVideoConcatQueryRequest {
  s.ListOrder = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetPageIndex(v int32) *LiveVideoConcatQueryRequest {
  s.PageIndex = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetPageSize(v int32) *LiveVideoConcatQueryRequest {
  s.PageSize = &v
  return s
}

type LiveVideoConcatQueryResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  LiveVideoConcatQueryData []*LiveVideoConcatQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s LiveVideoConcatQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryResponse) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryResponse) SetCode(v int32) *LiveVideoConcatQueryResponse {
  s.Code = &v
  return s
}

func (s *LiveVideoConcatQueryResponse) SetMessage(v string) *LiveVideoConcatQueryResponse {
  s.Message = &v
  return s
}

func (s *LiveVideoConcatQueryResponse) SetData(v []*LiveVideoConcatQueryData) *LiveVideoConcatQueryResponse {
  s.LiveVideoConcatQueryData = v
  return s
}

type LiveVideoConcatQueryData struct {
  // {"en":"Task id", "zh_CN":"任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"transNo", "zh_CN":"业务ID"}
  TranNo *string `json:"tranNo,omitempty" xml:"tranNo,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"状态
  // 取值范围：
  // 1：处理中
  // 2：成功
  // 3：失败"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Video details; This is only returned if statu is equal to 2", "zh_CN":"视频详情；只要statu = 2的时候才会返回这个值"}
  LiveVideoConcatQueryVideoInfo *LiveVideoConcatQueryVideoInfo `json:"videoInfo,omitempty" xml:"videoInfo,omitempty" require:"true"`
}

func (s LiveVideoConcatQueryData) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryData) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryData) SetTaskId(v string) *LiveVideoConcatQueryData {
  s.TaskId = &v
  return s
}

func (s *LiveVideoConcatQueryData) SetTranNo(v string) *LiveVideoConcatQueryData {
  s.TranNo = &v
  return s
}

func (s *LiveVideoConcatQueryData) SetStatus(v int32) *LiveVideoConcatQueryData {
  s.Status = &v
  return s
}

func (s *LiveVideoConcatQueryData) SetVideoInfo(v *LiveVideoConcatQueryVideoInfo) *LiveVideoConcatQueryData {
  s.LiveVideoConcatQueryVideoInfo = v
  return s
}

type LiveVideoConcatQueryVideoInfo struct {
  // {"en":"Video name", "zh_CN":"视频名称"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty" require:"true"`
  // {"en":"Video id", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Whether to encrypt transcoding files
  // Value range: 0(unencrypted), 1(encrypted)", "zh_CN":"是否加密转码文件
  // 取值范围 ：0(不加密)，1(加密)"}
  Encrypt *int32 `json:"encrypt,omitempty" xml:"encrypt,omitempty" require:"true"`
  // {"en":"The space occupied by the video, and the total space used by the video and its transcoding", "zh_CN":"视频占用空间大小，视频及其转码后视频的总空间使用量"}
  VideoSize *string `json:"videoSize,omitempty" xml:"videoSize,omitempty" require:"true"`
  // {"en":"Video duration", "zh_CN":"视频时长"}
  VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"Video upload time", "zh_CN":"视频上传时间"}
  UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty" require:"true"`
  // {"en":"Video modification time", "zh_CN":"视频修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"Video description", "zh_CN":"视频描述"}
  VideoDescription *string `json:"videoDescription,omitempty" xml:"videoDescription,omitempty" require:"true"`
  // {"en":"Video classification", "zh_CN":"视频分类"}
  VideoClassification *string `json:"videoClassification,omitempty" xml:"videoClassification,omitempty" require:"true"`
  // {"en":"Video cover URL", "zh_CN":"视频封面URL"}
  ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" require:"true"`
  // {"en":"The domain name of the video", "zh_CN":"视频的发布域名"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty" require:"true"`
  // {"en":"Name of the player used by the video", "zh_CN":"视频使用的播放器名称"}
  PlayerName *string `json:"playerName,omitempty" xml:"playerName,omitempty" require:"true"`
  // {"en":"The player ID used by the video", "zh_CN":"视频使用的播放器ID"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty" require:"true"`
  // {"en":"Video state
  // Value range: 0(normal), 1(masked)", "zh_CN":"视频状态
  // 取值范围：0(正常)，1(屏蔽)"}
  VideoState *string `json:"videoState,omitempty" xml:"videoState,omitempty" require:"true"`
  // {"en":"If authorized play is not enabled, the video transcoding status ranges from:
  // 1(transcoding), 2(not transcoding), 3(transcoding), 4(transcoding failed)
  // Value range of transcoding status when the Authorized Play (video encryption) function is enabled:
  // 1(encrypted transcoding), 2(unencrypted transcoding), 3(in transcoding), 4(transcoding failed), 5(not transcoding)", "zh_CN":"未开启授权播放，视频的转码状态的取值范围 ：
  // 1(已转码)，2(未转码)，3(转码中)，4(转码失败)
  // 开启授权播放（视频加密）功能时的转码状态的取值范围 ：
  // 1(已加密转码)，2(非加密转码)，3(转码中)，4(转码失败)，5(未转码)"}
  TranscodeState *string `json:"transcodeState,omitempty" xml:"transcodeState,omitempty" require:"true"`
  // {"en":"Video source
  // Value range:
  // 0(other), 1(upload), 2 (live to record), 3 (video pull), 4 (video cut), 5 (video splicing), 6 (edge pull to record), 10 (general version live to record), 11 (upload SDK), 12 (upload tool)", "zh_CN":"视频来源
  // 取值范围：
  // 0(其他)，1(上传)，2（直播转录制），3（视频拉取），4（视频剪切），5（视频拼接），6（边缘拉流录制），10（通用版直播转录制），11（上传SDK），12（上传工具）"}
  VideoSourceCode *int32 `json:"videoSourceCode,omitempty" xml:"videoSourceCode,omitempty" require:"true"`
  // {"en":"Video resolution and other information", "zh_CN":"视频分辨率等信息"}
  LiveVideoConcatQueryVideoResolutions []*LiveVideoConcatQueryVideoResolutions `json:"videoResolutions,omitempty" xml:"videoResolutions,omitempty" require:"true" type:"Repeated"`
}

func (s LiveVideoConcatQueryVideoInfo) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryVideoInfo) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoName(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoName = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoId(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoId = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetEncrypt(v int32) *LiveVideoConcatQueryVideoInfo {
  s.Encrypt = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoSize(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoSize = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoDuration(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoDuration = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetUploadTime(v string) *LiveVideoConcatQueryVideoInfo {
  s.UploadTime = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetUpdateTime(v string) *LiveVideoConcatQueryVideoInfo {
  s.UpdateTime = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoDescription(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoDescription = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoClassification(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoClassification = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetImageUrl(v string) *LiveVideoConcatQueryVideoInfo {
  s.ImageUrl = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetPublishDomain(v string) *LiveVideoConcatQueryVideoInfo {
  s.PublishDomain = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetPlayerName(v string) *LiveVideoConcatQueryVideoInfo {
  s.PlayerName = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetPlayerId(v string) *LiveVideoConcatQueryVideoInfo {
  s.PlayerId = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoState(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoState = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetTranscodeState(v string) *LiveVideoConcatQueryVideoInfo {
  s.TranscodeState = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoSourceCode(v int32) *LiveVideoConcatQueryVideoInfo {
  s.VideoSourceCode = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoResolutions(v []*LiveVideoConcatQueryVideoResolutions) *LiveVideoConcatQueryVideoInfo {
  s.LiveVideoConcatQueryVideoResolutions = v
  return s
}

type LiveVideoConcatQueryVideoResolutions struct {
  // {"en":"Clarity. Value range: 1(original painting), 2(smooth), 3(standard definition), 4(HD), 5(ultra HD), -99(recorded file)", "zh_CN":"清晰度。取值范围 ：1(原画)，2(流畅)，3(标清)，4(高清)，5(超清)，-99(录制文件)"}
  Clarity *int32 `json:"clarity,omitempty" xml:"clarity,omitempty" require:"true"`
  // {"en":"Terminal type. Value range: 0(PC), 1(original video)", "zh_CN":"终端类型。取值范围 ：0(PC)，1(原视频)"}
  ServerType *int32 `json:"serverType,omitempty" xml:"serverType,omitempty" require:"true"`
  // {"en":"Height", "zh_CN":"高度"}
  Height *int32 `json:"height,omitempty" xml:"height,omitempty" require:"true"`
  // {"en":"Width", "zh_CN":"宽度"}
  Width *int32 `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"File size(bit)", "zh_CN":"文件大小(单位为bit)"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s LiveVideoConcatQueryVideoResolutions) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryVideoResolutions) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryVideoResolutions) SetClarity(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.Clarity = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetServerType(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.ServerType = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetHeight(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.Height = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetWidth(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.Width = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetFileSize(v int64) *LiveVideoConcatQueryVideoResolutions {
  s.FileSize = &v
  return s
}

type LiveVideoConcatQueryPaths struct {
}

func (s LiveVideoConcatQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryPaths) GoString() string {
  return s.String()
}

type LiveVideoConcatQueryParameters struct {
}

func (s LiveVideoConcatQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryParameters) GoString() string {
  return s.String()
}

type LiveVideoConcatQueryRequestHeader struct {
}

func (s LiveVideoConcatQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryRequestHeader) GoString() string {
  return s.String()
}

type LiveVideoConcatQueryResponseHeader struct {
}

func (s LiveVideoConcatQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryResponseHeader) GoString() string {
  return s.String()
}




