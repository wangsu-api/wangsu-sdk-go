package livesnapshot

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

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




type QuerySnapshotParameterTemplateRequest struct {
  // {"en":"The unique identifier of the screenshot parameter template.","zh_CN":"截图参数模板的唯一标识符"}
  TemplateId *int `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s QuerySnapshotParameterTemplateRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySnapshotParameterTemplateRequest) GoString() string {
  return s.String()
}

func (s *QuerySnapshotParameterTemplateRequest) SetTemplateId(v int) *QuerySnapshotParameterTemplateRequest {
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
  TemplateId *int `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s DeleteSnapshotParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSnapshotParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *DeleteSnapshotParameterTemplatePaths) SetTemplateId(v int) *DeleteSnapshotParameterTemplatePaths {
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
  TemplateId *int `json:"templateId,omitempty" xml:"templateId,omitempty" require:"true"`
}

func (s ModifySnapshotParameterTemplatePaths) String() string {
  return tea.Prettify(s)
}

func (s ModifySnapshotParameterTemplatePaths) GoString() string {
  return s.String()
}

func (s *ModifySnapshotParameterTemplatePaths) SetTemplateId(v int) *ModifySnapshotParameterTemplatePaths {
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




