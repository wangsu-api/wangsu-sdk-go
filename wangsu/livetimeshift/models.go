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
  TemplateId *int `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s DeleteTimeShiftParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteTimeShiftParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *DeleteTimeShiftParameterTemplatePaths) SetTemplateId(v int) *DeleteTimeShiftParameterTemplatePaths {
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
  TemplateId *int `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s ModifyTimeShiftParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s ModifyTimeShiftParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *ModifyTimeShiftParameterTemplatePaths) SetTemplateId(v int) *ModifyTimeShiftParameterTemplatePaths {
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
  TemplateId *int `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QueryTimeShiftParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTimeShiftParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *QueryTimeShiftParameterTemplateRequest) SetTemplateId(v int) *QueryTimeShiftParameterTemplateRequest {
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




