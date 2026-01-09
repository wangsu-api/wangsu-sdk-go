package edgeai

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateWRConnectorRequest struct {
  // {"en":"QueueIT platform account ID, obtained from the 'Customer Account ID' value on the 'Company Profile' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT账号ID：QueueIT 平台账号，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Company Profile\" 页面的 \"Customer Account ID\" 值"}
  QtAccountId *string `json:"qtAccountId,omitempty" xml:"qtAccountId,omitempty"`
  // {"en":"QueueIT platform Secret Key, obtained from the 'KnownUser secret key' value in the 'Integration' tab at the top of the 'Settings' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT Secret Key：QueueIT 平台的 SecretKey，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Settings\" 页面顶部 \"Integration\" 标签中的 \"KnownUser secret key\" 值"}
  QtSK *string `json:"qtSK,omitempty" xml:"qtSK,omitempty"`
  // {"en":"QueueIT platform API Key, obtainable from the 'API Keys' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT API Key：QueueIT 平台的 APIKey，可从go.queue-it.net左侧导航栏 \"Account\" 下 \"API Keys\" 页面获取"}
  QtApiKey *string `json:"qtApiKey,omitempty" xml:"qtApiKey,omitempty"`
  // {"en":"Synchronization period (unit: seconds)","zh_CN":"同步周期：从 QueueIT 平台自动同步连接器配置的周期（单位：秒）"}
  SynCycle *int `json:"synCycle,omitempty" xml:"synCycle,omitempty"`
  // {"en":"Whether to verify the user identity secret key. If enabled, only invited users are allowed to access.","zh_CN":"验证用户身份：是否验证用户身份标识秘钥，开启时仅允许受邀用户访问"}
  VerifyUser *bool `json:"verifyUser,omitempty" xml:"verifyUser,omitempty"`
  // {"en":"Whether to generate a queue token","zh_CN":"生成队列令牌：是否生成访问等候室的队列令牌（ENTOKEN）"}
  GenerateQueueToken *bool `json:"generateQueueToken,omitempty" xml:"generateQueueToken,omitempty"`
  // {"en":"Queue token validity period (unit: seconds)","zh_CN":"队列令牌有效期：队列令牌（ENTOKEN）的有效期（单位：秒）"}
  QueueTokenValidityPeriod *int `json:"queueTokenValidityPeriod,omitempty" xml:"queueTokenValidityPeriod,omitempty"`
}

func (s UpdateWRConnectorRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorRequest) GoString() string {
  return s.String()
}

func (s *UpdateWRConnectorRequest) SetQtAccountId(v string) *UpdateWRConnectorRequest {
  s.QtAccountId = &v
  return s
}

func (s *UpdateWRConnectorRequest) SetQtSK(v string) *UpdateWRConnectorRequest {
  s.QtSK = &v
  return s
}

func (s *UpdateWRConnectorRequest) SetQtApiKey(v string) *UpdateWRConnectorRequest {
  s.QtApiKey = &v
  return s
}

func (s *UpdateWRConnectorRequest) SetSynCycle(v int) *UpdateWRConnectorRequest {
  s.SynCycle = &v
  return s
}

func (s *UpdateWRConnectorRequest) SetVerifyUser(v bool) *UpdateWRConnectorRequest {
  s.VerifyUser = &v
  return s
}

func (s *UpdateWRConnectorRequest) SetGenerateQueueToken(v bool) *UpdateWRConnectorRequest {
  s.GenerateQueueToken = &v
  return s
}

func (s *UpdateWRConnectorRequest) SetQueueTokenValidityPeriod(v int) *UpdateWRConnectorRequest {
  s.QueueTokenValidityPeriod = &v
  return s
}

type UpdateWRConnectorRequestHeader struct {
}

func (s UpdateWRConnectorRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorRequestHeader) GoString() string {
  return s.String()
}

type UpdateWRConnectorPaths struct {
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateWRConnectorPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorPaths) GoString() string {
  return s.String()
}

func (s *UpdateWRConnectorPaths) SetId(v string) *UpdateWRConnectorPaths {
  s.Id = &v
  return s
}

type UpdateWRConnectorParameters struct {
}

func (s UpdateWRConnectorParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorParameters) GoString() string {
  return s.String()
}

type UpdateWRConnectorResponse struct {
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"响应数据"}
  Data *UpdateWRConnectorResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s UpdateWRConnectorResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorResponse) GoString() string {
  return s.String()
}

func (s *UpdateWRConnectorResponse) SetCode(v string) *UpdateWRConnectorResponse {
  s.Code = &v
  return s
}

func (s *UpdateWRConnectorResponse) SetMessage(v string) *UpdateWRConnectorResponse {
  s.Message = &v
  return s
}

func (s *UpdateWRConnectorResponse) SetData(v *UpdateWRConnectorResponseData) *UpdateWRConnectorResponse {
  s.Data = v
  return s
}

type UpdateWRConnectorResponseData struct {
}

func (s UpdateWRConnectorResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorResponseData) GoString() string {
  return s.String()
}

type UpdateWRConnectorResponseHeader struct {
}

func (s UpdateWRConnectorResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWRConnectorResponseHeader) GoString() string {
  return s.String()
}




type ViewWRConnectorInfoRequest struct {
}

func (s ViewWRConnectorInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoRequest) GoString() string {
  return s.String()
}

type ViewWRConnectorInfoRequestHeader struct {
}

func (s ViewWRConnectorInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoRequestHeader) GoString() string {
  return s.String()
}

type ViewWRConnectorInfoPaths struct {
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s ViewWRConnectorInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoPaths) GoString() string {
  return s.String()
}

func (s *ViewWRConnectorInfoPaths) SetId(v string) *ViewWRConnectorInfoPaths {
  s.Id = &v
  return s
}

type ViewWRConnectorInfoParameters struct {
}

func (s ViewWRConnectorInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoParameters) GoString() string {
  return s.String()
}

type ViewWRConnectorInfoResponse struct {
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"响应数据"}
  Data *ViewWRConnectorInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ViewWRConnectorInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoResponse) GoString() string {
  return s.String()
}

func (s *ViewWRConnectorInfoResponse) SetCode(v string) *ViewWRConnectorInfoResponse {
  s.Code = &v
  return s
}

func (s *ViewWRConnectorInfoResponse) SetMessage(v string) *ViewWRConnectorInfoResponse {
  s.Message = &v
  return s
}

func (s *ViewWRConnectorInfoResponse) SetData(v *ViewWRConnectorInfoResponseData) *ViewWRConnectorInfoResponse {
  s.Data = v
  return s
}

type ViewWRConnectorInfoResponseData struct {
  // {"en":"connector","zh_CN":"连接器"}
  Connector *ViewWRConnectorInfoResponseDataConnector `json:"connector,omitempty" xml:"connector,omitempty" require:"true" type:"Struct"`
}

func (s ViewWRConnectorInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoResponseData) GoString() string {
  return s.String()
}

func (s *ViewWRConnectorInfoResponseData) SetConnector(v *ViewWRConnectorInfoResponseDataConnector) *ViewWRConnectorInfoResponseData {
  s.Connector = v
  return s
}

type ViewWRConnectorInfoResponseDataConnector struct {
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"The unique name of the QueueIT connector across the platform. Supports Chinese characters, English letters, numbers, hyphens (-), and underscores (_). The length must not exceed 50 characters.","zh_CN":"QueueIT连接器名称：全平台唯一，支持输入中文、英文、数字、中划线(-)、下划线(_)，长度不超过50个字符。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"QueueIT platform account ID, obtained from the 'Customer Account ID' value on the 'Company Profile' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT账号ID：QueueIT 平台账号，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Company Profile\" 页面的 \"Customer Account ID\" 值"}
  QtAccountId *string `json:"qtAccountId,omitempty" xml:"qtAccountId,omitempty" require:"true"`
  // {"en":"QueueIT platform Secret Key, obtained from the 'KnownUser secret key' value in the 'Integration' tab at the top of the 'Settings' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT Secret Key：QueueIT 平台的 SecretKey，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Settings\" 页面顶部 \"Integration\" 标签中的 \"KnownUser secret key\" 值"}
  QtSK *string `json:"qtSK,omitempty" xml:"qtSK,omitempty" require:"true"`
  // {"en":"QueueIT platform API Key, obtainable from the 'API Keys' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT API Key：QueueIT 平台的 APIKey，可从go.queue-it.net左侧导航栏 \"Account\" 下 \"API Keys\" 页面获取"}
  QtApiKey *string `json:"qtApiKey,omitempty" xml:"qtApiKey,omitempty" require:"true"`
  // {"en":"The period (in seconds) for automatically synchronizing connector configurations from the QueueIT platform.","zh_CN":"同步周期：从 QueueIT 平台自动同步连接器配置的周期（单位：秒）"}
  SynCycle *int `json:"synCycle,omitempty" xml:"synCycle,omitempty" require:"true"`
  // {"en":"Configuration content","zh_CN":"配置内容"}
  ConfigContent *string `json:"configContent,omitempty" xml:"configContent,omitempty" require:"true"`
  // {"en":"Whether to verify the user identity secret key. If enabled, only invited users are allowed to access.","zh_CN":"验证用户身份：是否验证用户身份标识秘钥，开启时仅允许受邀用户访问"}
  VerifyUser *bool `json:"verifyUser,omitempty" xml:"verifyUser,omitempty" require:"true"`
  // {"en":"Whether to generate a queue token (ENTOKEN) for accessing the waiting room.","zh_CN":"生成队列令牌：是否生成访问等候室的队列令牌（ENTOKEN）"}
  GenerateQueueToken *bool `json:"generateQueueToken,omitempty" xml:"generateQueueToken,omitempty" require:"true"`
  // {"en":"The validity period of the queue token (ENTOKEN) in seconds.","zh_CN":"队列令牌有效期：队列令牌（ENTOKEN）的有效期（单位：秒）"}
  QueueTokenValidityPeriod *int `json:"queueTokenValidityPeriod,omitempty" xml:"queueTokenValidityPeriod,omitempty" require:"true"`
  // {"en":"Status: 0 Not deployed 1 Deployed 2 Effective","zh_CN":"状态：0未部署 1已部署 2已生效"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Creation time (timestamp)","zh_CN":"创建时间（时间戳）"}
  CreateTime *int `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modification time (timestamp)","zh_CN":"修改时间（时间戳）"}
  UpdateTime *int `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"Whether the configuration has been modified or not deployed","zh_CN":"是否配置修改，未部署"}
  ConfigChanged *bool `json:"configChanged,omitempty" xml:"configChanged,omitempty" require:"true"`
}

func (s ViewWRConnectorInfoResponseDataConnector) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoResponseDataConnector) GoString() string {
  return s.String()
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetId(v string) *ViewWRConnectorInfoResponseDataConnector {
  s.Id = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetName(v string) *ViewWRConnectorInfoResponseDataConnector {
  s.Name = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetQtAccountId(v string) *ViewWRConnectorInfoResponseDataConnector {
  s.QtAccountId = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetQtSK(v string) *ViewWRConnectorInfoResponseDataConnector {
  s.QtSK = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetQtApiKey(v string) *ViewWRConnectorInfoResponseDataConnector {
  s.QtApiKey = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetSynCycle(v int) *ViewWRConnectorInfoResponseDataConnector {
  s.SynCycle = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetConfigContent(v string) *ViewWRConnectorInfoResponseDataConnector {
  s.ConfigContent = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetVerifyUser(v bool) *ViewWRConnectorInfoResponseDataConnector {
  s.VerifyUser = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetGenerateQueueToken(v bool) *ViewWRConnectorInfoResponseDataConnector {
  s.GenerateQueueToken = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetQueueTokenValidityPeriod(v int) *ViewWRConnectorInfoResponseDataConnector {
  s.QueueTokenValidityPeriod = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetStatus(v int) *ViewWRConnectorInfoResponseDataConnector {
  s.Status = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetCreateTime(v int) *ViewWRConnectorInfoResponseDataConnector {
  s.CreateTime = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetUpdateTime(v int) *ViewWRConnectorInfoResponseDataConnector {
  s.UpdateTime = &v
  return s
}

func (s *ViewWRConnectorInfoResponseDataConnector) SetConfigChanged(v bool) *ViewWRConnectorInfoResponseDataConnector {
  s.ConfigChanged = &v
  return s
}

type ViewWRConnectorInfoResponseHeader struct {
}

func (s ViewWRConnectorInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ViewWRConnectorInfoResponseHeader) GoString() string {
  return s.String()
}




type EnableDisableFileRequest struct {
  // {"en":"knowledgeId", "zh_CN":"知识库id"}
  KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty" require:"true"`
  // {"en":"fileIds", "zh_CN":"文件id列表"}
  FileIds []*string `json:"fileIds,omitempty" xml:"fileIds,omitempty" require:"true" type:"Repeated"`
  // {"en":"isEnable", "zh_CN":"是否启用"}
  Enable *bool `json:"enable,omitempty" xml:"enable,omitempty" require:"true"`
}

func (s EnableDisableFileRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableFileRequest) GoString() string {
  return s.String()
}

func (s *EnableDisableFileRequest) SetKnowledgeBaseId(v string) *EnableDisableFileRequest {
  s.KnowledgeBaseId = &v
  return s
}

func (s *EnableDisableFileRequest) SetFileIds(v []*string) *EnableDisableFileRequest {
  s.FileIds = v
  return s
}

func (s *EnableDisableFileRequest) SetEnable(v bool) *EnableDisableFileRequest {
  s.Enable = &v
  return s
}

type EnableDisableFileResponse struct {
  // {"en":"code", "zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"响应信息"}
  Messgae *string `json:"messgae,omitempty" xml:"messgae,omitempty" require:"true"`
}

func (s EnableDisableFileResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableFileResponse) GoString() string {
  return s.String()
}

func (s *EnableDisableFileResponse) SetCode(v int) *EnableDisableFileResponse {
  s.Code = &v
  return s
}

func (s *EnableDisableFileResponse) SetMessgae(v string) *EnableDisableFileResponse {
  s.Messgae = &v
  return s
}

type EnableDisableFilePaths struct {
}

func (s EnableDisableFilePaths) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableFilePaths) GoString() string {
  return s.String()
}

type EnableDisableFileParameters struct {
}

func (s EnableDisableFileParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableFileParameters) GoString() string {
  return s.String()
}

type EnableDisableFileRequestHeader struct {
}

func (s EnableDisableFileRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableFileRequestHeader) GoString() string {
  return s.String()
}

type EnableDisableFileResponseHeader struct {
}

func (s EnableDisableFileResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableFileResponseHeader) GoString() string {
  return s.String()
}




type RagKnowledgeBaseListServiceRequest struct {
  // {"en":"knowledge base name","zh_CN":"知识库名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"page num","zh_CN":"页数"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"page size","zh_CN":"页大小"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
}

func (s RagKnowledgeBaseListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceRequest) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseListServiceRequest) SetName(v string) *RagKnowledgeBaseListServiceRequest {
  s.Name = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRequest) SetPageNo(v int) *RagKnowledgeBaseListServiceRequest {
  s.PageNo = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRequest) SetPageSize(v int) *RagKnowledgeBaseListServiceRequest {
  s.PageSize = &v
  return s
}

type RagKnowledgeBaseListServiceRequestHeader struct {
}

func (s RagKnowledgeBaseListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceRequestHeader) GoString() string {
  return s.String()
}

type RagKnowledgeBaseListServicePaths struct {
}

func (s RagKnowledgeBaseListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServicePaths) GoString() string {
  return s.String()
}

type RagKnowledgeBaseListServiceParameters struct {
}

func (s RagKnowledgeBaseListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceParameters) GoString() string {
  return s.String()
}

type RagKnowledgeBaseListServiceResponse struct {
  // {"en":"total","zh_CN":"总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"code","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"array","zh_CN":"数组"}
  Data []*RagKnowledgeBaseListServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"message","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RagKnowledgeBaseListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceResponse) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseListServiceResponse) SetTotal(v int) *RagKnowledgeBaseListServiceResponse {
  s.Total = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponse) SetCode(v int) *RagKnowledgeBaseListServiceResponse {
  s.Code = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponse) SetData(v []*RagKnowledgeBaseListServiceResponseData) *RagKnowledgeBaseListServiceResponse {
  s.Data = v
  return s
}

func (s *RagKnowledgeBaseListServiceResponse) SetMessage(v string) *RagKnowledgeBaseListServiceResponse {
  s.Message = &v
  return s
}

type RagKnowledgeBaseListServiceResponseData struct     {
  // {"en":"last update time","zh_CN":"最近更新时间"}
  LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty" require:"true"`
  // {"en":"size","zh_CN":"知识库大小"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"create time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"knowledge base name","zh_CN":"知识库名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"knowledge base desc","zh_CN":"知识库描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"knowledge base id","zh_CN":"知识库 id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s RagKnowledgeBaseListServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceResponseData) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseListServiceResponseData) SetLastUpdated(v string) *RagKnowledgeBaseListServiceResponseData {
  s.LastUpdated = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponseData) SetSize(v int) *RagKnowledgeBaseListServiceResponseData {
  s.Size = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponseData) SetCreateTime(v string) *RagKnowledgeBaseListServiceResponseData {
  s.CreateTime = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponseData) SetName(v string) *RagKnowledgeBaseListServiceResponseData {
  s.Name = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponseData) SetDescription(v string) *RagKnowledgeBaseListServiceResponseData {
  s.Description = &v
  return s
}

func (s *RagKnowledgeBaseListServiceResponseData) SetId(v string) *RagKnowledgeBaseListServiceResponseData {
  s.Id = &v
  return s
}

type RagKnowledgeBaseListServiceResponseHeader struct {
}

func (s RagKnowledgeBaseListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceResponseHeader) GoString() string {
  return s.String()
}




type RagFileDeleteServiceRequest struct {
  // {"en":"knowledge base id","zh_CN":"知识库 ID"}
  KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty" require:"true"`
  // {"en":"file id list","zh_CN":"文件 id列表"}
  FileIds []*string `json:"fileIds,omitempty" xml:"fileIds,omitempty" require:"true" type:"Repeated"`
}

func (s RagFileDeleteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s RagFileDeleteServiceRequest) GoString() string {
  return s.String()
}

func (s *RagFileDeleteServiceRequest) SetKnowledgeBaseId(v string) *RagFileDeleteServiceRequest {
  s.KnowledgeBaseId = &v
  return s
}

func (s *RagFileDeleteServiceRequest) SetFileIds(v []*string) *RagFileDeleteServiceRequest {
  s.FileIds = v
  return s
}

type RagFileDeleteServiceRequestHeader struct {
}

func (s RagFileDeleteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RagFileDeleteServiceRequestHeader) GoString() string {
  return s.String()
}

type RagFileDeleteServicePaths struct {
}

func (s RagFileDeleteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s RagFileDeleteServicePaths) GoString() string {
  return s.String()
}

type RagFileDeleteServiceParameters struct {
}

func (s RagFileDeleteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s RagFileDeleteServiceParameters) GoString() string {
  return s.String()
}

type RagFileDeleteServiceResponse struct {
  // {"en":"code","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RagFileDeleteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s RagFileDeleteServiceResponse) GoString() string {
  return s.String()
}

func (s *RagFileDeleteServiceResponse) SetCode(v int) *RagFileDeleteServiceResponse {
  s.Code = &v
  return s
}

func (s *RagFileDeleteServiceResponse) SetMessage(v string) *RagFileDeleteServiceResponse {
  s.Message = &v
  return s
}

type RagFileDeleteServiceResponseHeader struct {
}

func (s RagFileDeleteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RagFileDeleteServiceResponseHeader) GoString() string {
  return s.String()
}




type RagKnowledgeBaseDeleteServiceRequest struct {
  // {"en":"knowledge base id","zh_CN":"知识库 id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s RagKnowledgeBaseDeleteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseDeleteServiceRequest) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseDeleteServiceRequest) SetId(v string) *RagKnowledgeBaseDeleteServiceRequest {
  s.Id = &v
  return s
}

type RagKnowledgeBaseDeleteServiceRequestHeader struct {
}

func (s RagKnowledgeBaseDeleteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseDeleteServiceRequestHeader) GoString() string {
  return s.String()
}

type RagKnowledgeBaseDeleteServicePaths struct {
}

func (s RagKnowledgeBaseDeleteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseDeleteServicePaths) GoString() string {
  return s.String()
}

type RagKnowledgeBaseDeleteServiceParameters struct {
}

func (s RagKnowledgeBaseDeleteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseDeleteServiceParameters) GoString() string {
  return s.String()
}

type RagKnowledgeBaseDeleteServiceResponse struct {
  // {"en":"code","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RagKnowledgeBaseDeleteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseDeleteServiceResponse) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseDeleteServiceResponse) SetCode(v int) *RagKnowledgeBaseDeleteServiceResponse {
  s.Code = &v
  return s
}

func (s *RagKnowledgeBaseDeleteServiceResponse) SetMessage(v string) *RagKnowledgeBaseDeleteServiceResponse {
  s.Message = &v
  return s
}

type RagKnowledgeBaseDeleteServiceResponseHeader struct {
}

func (s RagKnowledgeBaseDeleteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseDeleteServiceResponseHeader) GoString() string {
  return s.String()
}




type ListRagFileRequest struct {
  // {"en":"knowledge id", "zh_CN":"知识库id"}
  KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty" require:"true"`
  // {"en":"fileName", "zh_CN":"文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
  // {"en":"pageNo", "zh_CN":"第几页"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"pageSize", "zh_CN":"每页大小"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
}

func (s ListRagFileRequest) String() string {
  return tea.Prettify(s)
}

func (s ListRagFileRequest) GoString() string {
  return s.String()
}

func (s *ListRagFileRequest) SetKnowledgeBaseId(v string) *ListRagFileRequest {
  s.KnowledgeBaseId = &v
  return s
}

func (s *ListRagFileRequest) SetFileName(v string) *ListRagFileRequest {
  s.FileName = &v
  return s
}

func (s *ListRagFileRequest) SetPageNo(v int) *ListRagFileRequest {
  s.PageNo = &v
  return s
}

func (s *ListRagFileRequest) SetPageSize(v int) *ListRagFileRequest {
  s.PageSize = &v
  return s
}

type ListRagFileResponse struct {
  // {"en":"total", "zh_CN":"总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"code", "zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"", "zh_CN":""}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"update time", "zh_CN":"更新时间"}
  LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty" require:"true"`
  // {"en":"file name", "zh_CN":"文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"file size", "zh_CN":"文件大小"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"isEnabled", "zh_CN":"是否启用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"fileId", "zh_CN":"文件id"}
  FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"0:失败,1:成功,2:处理中"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ListRagFileResponse) String() string {
  return tea.Prettify(s)
}

func (s ListRagFileResponse) GoString() string {
  return s.String()
}

func (s *ListRagFileResponse) SetTotal(v int) *ListRagFileResponse {
  s.Total = &v
  return s
}

func (s *ListRagFileResponse) SetCode(v int) *ListRagFileResponse {
  s.Code = &v
  return s
}

func (s *ListRagFileResponse) SetData(v []*string) *ListRagFileResponse {
  s.Data = v
  return s
}

func (s *ListRagFileResponse) SetLastUpdated(v string) *ListRagFileResponse {
  s.LastUpdated = &v
  return s
}

func (s *ListRagFileResponse) SetFileName(v string) *ListRagFileResponse {
  s.FileName = &v
  return s
}

func (s *ListRagFileResponse) SetSize(v int) *ListRagFileResponse {
  s.Size = &v
  return s
}

func (s *ListRagFileResponse) SetIsEnabled(v bool) *ListRagFileResponse {
  s.IsEnabled = &v
  return s
}

func (s *ListRagFileResponse) SetFileId(v string) *ListRagFileResponse {
  s.FileId = &v
  return s
}

func (s *ListRagFileResponse) SetStatus(v int) *ListRagFileResponse {
  s.Status = &v
  return s
}

func (s *ListRagFileResponse) SetMessage(v string) *ListRagFileResponse {
  s.Message = &v
  return s
}

type ListRagFilePaths struct {
}

func (s ListRagFilePaths) String() string {
  return tea.Prettify(s)
}

func (s ListRagFilePaths) GoString() string {
  return s.String()
}

type ListRagFileParameters struct {
}

func (s ListRagFileParameters) String() string {
  return tea.Prettify(s)
}

func (s ListRagFileParameters) GoString() string {
  return s.String()
}

type ListRagFileRequestHeader struct {
}

func (s ListRagFileRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListRagFileRequestHeader) GoString() string {
  return s.String()
}

type ListRagFileResponseHeader struct {
}

func (s ListRagFileResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListRagFileResponseHeader) GoString() string {
  return s.String()
}




type ListConnectorsByPageRequest struct {
}

func (s ListConnectorsByPageRequest) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageRequest) GoString() string {
  return s.String()
}

type ListConnectorsByPageRequestHeader struct {
}

func (s ListConnectorsByPageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageRequestHeader) GoString() string {
  return s.String()
}

type ListConnectorsByPagePaths struct {
}

func (s ListConnectorsByPagePaths) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPagePaths) GoString() string {
  return s.String()
}

type ListConnectorsByPageParameters struct {
  // {"en":"Name for fuzzy matching.","zh_CN":"名称（模糊查询）"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"page number","zh_CN":"页码"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"page size","zh_CN":"每页大小"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Connector ID for exact matching.","zh_CN":"连接器ID（精确查询）"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListConnectorsByPageParameters) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageParameters) GoString() string {
  return s.String()
}

func (s *ListConnectorsByPageParameters) SetName(v string) *ListConnectorsByPageParameters {
  s.Name = &v
  return s
}

func (s *ListConnectorsByPageParameters) SetPageNo(v int) *ListConnectorsByPageParameters {
  s.PageNo = &v
  return s
}

func (s *ListConnectorsByPageParameters) SetPageSize(v int) *ListConnectorsByPageParameters {
  s.PageSize = &v
  return s
}

func (s *ListConnectorsByPageParameters) SetId(v string) *ListConnectorsByPageParameters {
  s.Id = &v
  return s
}

type ListConnectorsByPageResponse struct {
  // {"en":"response code","zh_CN":"响应编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"响应数据"}
  Data *ListConnectorsByPageResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListConnectorsByPageResponse) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageResponse) GoString() string {
  return s.String()
}

func (s *ListConnectorsByPageResponse) SetCode(v string) *ListConnectorsByPageResponse {
  s.Code = &v
  return s
}

func (s *ListConnectorsByPageResponse) SetMessage(v string) *ListConnectorsByPageResponse {
  s.Message = &v
  return s
}

func (s *ListConnectorsByPageResponse) SetData(v *ListConnectorsByPageResponseData) *ListConnectorsByPageResponse {
  s.Data = v
  return s
}

type ListConnectorsByPageResponseData struct {
  // {"en":"Total number of records","zh_CN":"总记录数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of connectors","zh_CN":"连接器列表"}
  Connectors []*ListConnectorsByPageResponseDataConnectors `json:"connectors,omitempty" xml:"connectors,omitempty" require:"true" type:"Repeated"`
}

func (s ListConnectorsByPageResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageResponseData) GoString() string {
  return s.String()
}

func (s *ListConnectorsByPageResponseData) SetTotal(v int) *ListConnectorsByPageResponseData {
  s.Total = &v
  return s
}

func (s *ListConnectorsByPageResponseData) SetConnectors(v []*ListConnectorsByPageResponseDataConnectors) *ListConnectorsByPageResponseData {
  s.Connectors = v
  return s
}

type ListConnectorsByPageResponseDataConnectors struct     {
  // {"en":"Connector ID","zh_CN":"连接器ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"The unique name of the QueueIT connector across the platform. Supports Chinese characters, English letters, numbers, hyphens (-), and underscores (_). The length must not exceed 50 characters.","zh_CN":"QueueIT连接器名称：全平台唯一，支持输入中文、英文、数字、中划线(-)、下划线(_)，长度不超过50个字符。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"QueueIT platform account ID, obtained from the 'Customer Account ID' value on the 'Company Profile' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT账号ID：QueueIT 平台账号，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Company Profile\" 页面的 \"Customer Account ID\" 值"}
  QtAccountId *string `json:"qtAccountId,omitempty" xml:"qtAccountId,omitempty" require:"true"`
  // {"en":"QueueIT platform Secret Key, obtained from the 'KnownUser secret key' value in the 'Integration' tab at the top of the 'Settings' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT Secret Key：QueueIT 平台的 SecretKey，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Settings\" 页面顶部 \"Integration\" 标签中的 \"KnownUser secret key\" 值"}
  QtSK *string `json:"qtSK,omitempty" xml:"qtSK,omitempty" require:"true"`
  // {"en":"QueueIT platform API Key, obtainable from the 'API Keys' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT API Key：QueueIT 平台的 APIKey，可从go.queue-it.net左侧导航栏 \"Account\" 下 \"API Keys\" 页面获取"}
  QtApiKey *string `json:"qtApiKey,omitempty" xml:"qtApiKey,omitempty" require:"true"`
  // {"en":"Synchronization period (unit: seconds)","zh_CN":"同步周期：从 QueueIT 平台自动同步连接器配置的周期（单位：秒）"}
  SynCycle *int `json:"synCycle,omitempty" xml:"synCycle,omitempty" require:"true"`
  // {"en":"Configuration content","zh_CN":"配置内容"}
  ConfigContent *string `json:"configContent,omitempty" xml:"configContent,omitempty" require:"true"`
  // {"en":"Whether to verify the user identity secret key. If enabled, only invited users are allowed to access.","zh_CN":"验证用户身份：是否验证用户身份标识秘钥，开启时仅允许受邀用户访问"}
  VerifyUser *bool `json:"verifyUser,omitempty" xml:"verifyUser,omitempty" require:"true"`
  // {"en":"Whether to generate a queue token","zh_CN":"生成队列令牌：是否生成访问等候室的队列令牌（ENTOKEN）"}
  GenerateQueueToken *bool `json:"generateQueueToken,omitempty" xml:"generateQueueToken,omitempty" require:"true"`
  // {"en":"Queue token validity period (unit: seconds)","zh_CN":"队列令牌有效期：队列令牌（ENTOKEN）的有效期（单位：秒）"}
  QueueTokenValidityPeriod *int `json:"queueTokenValidityPeriod,omitempty" xml:"queueTokenValidityPeriod,omitempty" require:"true"`
  // {"en":"Status: 0 Not deployed 1 Deployed 2 Effective","zh_CN":"状态：0未部署 1已部署 2已生效"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Creation time (timestamp)","zh_CN":"创建时间（时间戳）"}
  CreateTime *int `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modification time (timestamp)","zh_CN":"修改时间（时间戳）"}
  UpdateTime *int `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"Whether the configuration has been modified or not deployed","zh_CN":"是否配置修改，未部署"}
  ConfigChanged *bool `json:"configChanged,omitempty" xml:"configChanged,omitempty" require:"true"`
}

func (s ListConnectorsByPageResponseDataConnectors) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageResponseDataConnectors) GoString() string {
  return s.String()
}

func (s *ListConnectorsByPageResponseDataConnectors) SetId(v string) *ListConnectorsByPageResponseDataConnectors {
  s.Id = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetName(v string) *ListConnectorsByPageResponseDataConnectors {
  s.Name = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetQtAccountId(v string) *ListConnectorsByPageResponseDataConnectors {
  s.QtAccountId = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetQtSK(v string) *ListConnectorsByPageResponseDataConnectors {
  s.QtSK = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetQtApiKey(v string) *ListConnectorsByPageResponseDataConnectors {
  s.QtApiKey = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetSynCycle(v int) *ListConnectorsByPageResponseDataConnectors {
  s.SynCycle = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetConfigContent(v string) *ListConnectorsByPageResponseDataConnectors {
  s.ConfigContent = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetVerifyUser(v bool) *ListConnectorsByPageResponseDataConnectors {
  s.VerifyUser = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetGenerateQueueToken(v bool) *ListConnectorsByPageResponseDataConnectors {
  s.GenerateQueueToken = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetQueueTokenValidityPeriod(v int) *ListConnectorsByPageResponseDataConnectors {
  s.QueueTokenValidityPeriod = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetStatus(v int) *ListConnectorsByPageResponseDataConnectors {
  s.Status = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetCreateTime(v int) *ListConnectorsByPageResponseDataConnectors {
  s.CreateTime = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetUpdateTime(v int) *ListConnectorsByPageResponseDataConnectors {
  s.UpdateTime = &v
  return s
}

func (s *ListConnectorsByPageResponseDataConnectors) SetConfigChanged(v bool) *ListConnectorsByPageResponseDataConnectors {
  s.ConfigChanged = &v
  return s
}

type ListConnectorsByPageResponseHeader struct {
}

func (s ListConnectorsByPageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListConnectorsByPageResponseHeader) GoString() string {
  return s.String()
}




type CreateRagFileRequest struct {
  // {"en":"knowledge id", "zh_CN":"知识库id"}
  KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty" require:"true"`
  // {"en":"file name,Maximum Length 255", "zh_CN":"文件名，最大长度 255"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"file content after Base64 encode", "zh_CN":"文件内容Base64 encode "}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s CreateRagFileRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateRagFileRequest) GoString() string {
  return s.String()
}

func (s *CreateRagFileRequest) SetKnowledgeBaseId(v string) *CreateRagFileRequest {
  s.KnowledgeBaseId = &v
  return s
}

func (s *CreateRagFileRequest) SetFileName(v string) *CreateRagFileRequest {
  s.FileName = &v
  return s
}

func (s *CreateRagFileRequest) SetContent(v string) *CreateRagFileRequest {
  s.Content = &v
  return s
}

type CreateRagFileResponse struct {
  // {"en":"code", "zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"file id", "zh_CN":"文件id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateRagFileResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateRagFileResponse) GoString() string {
  return s.String()
}

func (s *CreateRagFileResponse) SetCode(v int) *CreateRagFileResponse {
  s.Code = &v
  return s
}

func (s *CreateRagFileResponse) SetId(v string) *CreateRagFileResponse {
  s.Id = &v
  return s
}

func (s *CreateRagFileResponse) SetMessage(v string) *CreateRagFileResponse {
  s.Message = &v
  return s
}

type CreateRagFilePaths struct {
}

func (s CreateRagFilePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateRagFilePaths) GoString() string {
  return s.String()
}

type CreateRagFileParameters struct {
}

func (s CreateRagFileParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateRagFileParameters) GoString() string {
  return s.String()
}

type CreateRagFileRequestHeader struct {
}

func (s CreateRagFileRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateRagFileRequestHeader) GoString() string {
  return s.String()
}

type CreateRagFileResponseHeader struct {
}

func (s CreateRagFileResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateRagFileResponseHeader) GoString() string {
  return s.String()
}




type RagKnowledgeBaseCreateServiceRequest struct {
  // {"en":"name,Maximum Length 100","zh_CN":"知识库名称，最大长度 100"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"description,Maximum Length 500","zh_CN":"知识库描述，最大长度 500"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s RagKnowledgeBaseCreateServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseCreateServiceRequest) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseCreateServiceRequest) SetName(v string) *RagKnowledgeBaseCreateServiceRequest {
  s.Name = &v
  return s
}

func (s *RagKnowledgeBaseCreateServiceRequest) SetDescription(v string) *RagKnowledgeBaseCreateServiceRequest {
  s.Description = &v
  return s
}

type RagKnowledgeBaseCreateServiceRequestHeader struct {
}

func (s RagKnowledgeBaseCreateServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseCreateServiceRequestHeader) GoString() string {
  return s.String()
}

type RagKnowledgeBaseCreateServicePaths struct {
}

func (s RagKnowledgeBaseCreateServicePaths) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseCreateServicePaths) GoString() string {
  return s.String()
}

type RagKnowledgeBaseCreateServiceParameters struct {
}

func (s RagKnowledgeBaseCreateServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseCreateServiceParameters) GoString() string {
  return s.String()
}

type RagKnowledgeBaseCreateServiceResponse struct {
  // {"en":"code","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"knowledge base id","zh_CN":"知识库 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RagKnowledgeBaseCreateServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseCreateServiceResponse) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseCreateServiceResponse) SetCode(v int) *RagKnowledgeBaseCreateServiceResponse {
  s.Code = &v
  return s
}

func (s *RagKnowledgeBaseCreateServiceResponse) SetId(v string) *RagKnowledgeBaseCreateServiceResponse {
  s.Id = &v
  return s
}

func (s *RagKnowledgeBaseCreateServiceResponse) SetMessage(v string) *RagKnowledgeBaseCreateServiceResponse {
  s.Message = &v
  return s
}

type RagKnowledgeBaseCreateServiceResponseHeader struct {
}

func (s RagKnowledgeBaseCreateServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseCreateServiceResponseHeader) GoString() string {
  return s.String()
}




type CreateWRConnectorRequest struct {
  // {"en":"The unique name of the QueueIT connector across the platform. Supports Chinese characters, English letters, numbers, hyphens (-), and underscores (_). The length must not exceed 50 characters.","zh_CN":"QueueIT连接器名称：全平台唯一，支持输入中文、英文、数字、中划线(-)、下划线(_)，长度不超过50个字符。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"QueueIT platform account ID, obtained from the 'Customer Account ID' value on the 'Company Profile' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT账号ID：QueueIT 平台账号，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Company Profile\" 页面的 \"Customer Account ID\" 值"}
  QtAccountId *string `json:"qtAccountId,omitempty" xml:"qtAccountId,omitempty" require:"true"`
  // {"en":"QueueIT platform Secret Key, obtained from the 'KnownUser secret key' value in the 'Integration' tab at the top of the 'Settings' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT Secret Key：QueueIT 平台的 SecretKey，取自go.queue-it.net左侧导航栏 \"Account\" 下 \"Settings\" 页面顶部 \"Integration\" 标签中的 \"KnownUser secret key\" 值"}
  QtSK *string `json:"qtSK,omitempty" xml:"qtSK,omitempty"`
  // {"en":"QueueIT platform API Key, obtainable from the 'API Keys' page under 'Account' in the left navigation bar of go.queue-it.net.","zh_CN":"QueueIT API Key：QueueIT 平台的 APIKey，可从go.queue-it.net左侧导航栏 \"Account\" 下 \"API Keys\" 页面获取"}
  QtApiKey *string `json:"qtApiKey,omitempty" xml:"qtApiKey,omitempty"`
  // {"en":"Synchronization period (unit: seconds)","zh_CN":"同步周期：从 QueueIT 平台自动同步连接器配置的周期（单位：秒）"}
  SynCycle *int `json:"synCycle,omitempty" xml:"synCycle,omitempty"`
  // {"en":"Whether to verify the user identity secret key. If enabled, only invited users are allowed to access.","zh_CN":"验证用户身份：是否验证用户身份标识秘钥，开启时仅允许受邀用户访问"}
  VerifyUser *bool `json:"verifyUser,omitempty" xml:"verifyUser,omitempty"`
  // {"en":"Whether to generate a queue token","zh_CN":"生成队列令牌：是否生成访问等候室的队列令牌（ENTOKEN）"}
  GenerateQueueToken *bool `json:"generateQueueToken,omitempty" xml:"generateQueueToken,omitempty"`
  // {"en":"Queue token validity period (unit: seconds)","zh_CN":"队列令牌有效期：队列令牌（ENTOKEN）的有效期（单位：秒）"}
  QueueTokenValidityPeriod *int `json:"queueTokenValidityPeriod,omitempty" xml:"queueTokenValidityPeriod,omitempty"`
}

func (s CreateWRConnectorRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorRequest) GoString() string {
  return s.String()
}

func (s *CreateWRConnectorRequest) SetName(v string) *CreateWRConnectorRequest {
  s.Name = &v
  return s
}

func (s *CreateWRConnectorRequest) SetQtAccountId(v string) *CreateWRConnectorRequest {
  s.QtAccountId = &v
  return s
}

func (s *CreateWRConnectorRequest) SetQtSK(v string) *CreateWRConnectorRequest {
  s.QtSK = &v
  return s
}

func (s *CreateWRConnectorRequest) SetQtApiKey(v string) *CreateWRConnectorRequest {
  s.QtApiKey = &v
  return s
}

func (s *CreateWRConnectorRequest) SetSynCycle(v int) *CreateWRConnectorRequest {
  s.SynCycle = &v
  return s
}

func (s *CreateWRConnectorRequest) SetVerifyUser(v bool) *CreateWRConnectorRequest {
  s.VerifyUser = &v
  return s
}

func (s *CreateWRConnectorRequest) SetGenerateQueueToken(v bool) *CreateWRConnectorRequest {
  s.GenerateQueueToken = &v
  return s
}

func (s *CreateWRConnectorRequest) SetQueueTokenValidityPeriod(v int) *CreateWRConnectorRequest {
  s.QueueTokenValidityPeriod = &v
  return s
}

type CreateWRConnectorRequestHeader struct {
}

func (s CreateWRConnectorRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorRequestHeader) GoString() string {
  return s.String()
}

type CreateWRConnectorPaths struct {
}

func (s CreateWRConnectorPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorPaths) GoString() string {
  return s.String()
}

type CreateWRConnectorParameters struct {
}

func (s CreateWRConnectorParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorParameters) GoString() string {
  return s.String()
}

type CreateWRConnectorResponse struct {
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"响应数据对象"}
  Data *CreateWRConnectorResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateWRConnectorResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorResponse) GoString() string {
  return s.String()
}

func (s *CreateWRConnectorResponse) SetCode(v string) *CreateWRConnectorResponse {
  s.Code = &v
  return s
}

func (s *CreateWRConnectorResponse) SetMessage(v string) *CreateWRConnectorResponse {
  s.Message = &v
  return s
}

func (s *CreateWRConnectorResponse) SetData(v *CreateWRConnectorResponseData) *CreateWRConnectorResponse {
  s.Data = v
  return s
}

type CreateWRConnectorResponseData struct {
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateWRConnectorResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorResponseData) GoString() string {
  return s.String()
}

func (s *CreateWRConnectorResponseData) SetId(v string) *CreateWRConnectorResponseData {
  s.Id = &v
  return s
}

type CreateWRConnectorResponseHeader struct {
}

func (s CreateWRConnectorResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWRConnectorResponseHeader) GoString() string {
  return s.String()
}




type AddImageConfigRequest struct {
  // {"en":"Properties of the image processing configuration, including basic information.","zh_CN":"图片处理配置的属性，包含基本信息等。"}
  Property *AddImageConfigRequestProperty `json:"property,omitempty" xml:"property,omitempty" type:"Struct"`
  // {"en":"Collection of image processing policies, containing specific processing rules.","zh_CN":"图片处理策略的集合，包含具体的处理规则。"}
  PolicySets []*AddImageConfigRequestPolicySets `json:"policySets,omitempty" xml:"policySets,omitempty" type:"Repeated"`
}

func (s AddImageConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigRequest) GoString() string {
  return s.String()
}

func (s *AddImageConfigRequest) SetProperty(v *AddImageConfigRequestProperty) *AddImageConfigRequest {
  s.Property = v
  return s
}

func (s *AddImageConfigRequest) SetPolicySets(v []*AddImageConfigRequestPolicySets) *AddImageConfigRequest {
  s.PolicySets = v
  return s
}

type AddImageConfigRequestProperty struct {
}

func (s AddImageConfigRequestProperty) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigRequestProperty) GoString() string {
  return s.String()
}

type AddImageConfigRequestPolicySets struct     {
}

func (s AddImageConfigRequestPolicySets) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigRequestPolicySets) GoString() string {
  return s.String()
}

type AddImageConfigRequestHeader struct {
}

func (s AddImageConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigRequestHeader) GoString() string {
  return s.String()
}

type AddImageConfigPaths struct {
}

func (s AddImageConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigPaths) GoString() string {
  return s.String()
}

type AddImageConfigParameters struct {
}

func (s AddImageConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigParameters) GoString() string {
  return s.String()
}

type AddImageConfigResponse struct {
  // {"en":"The response code indicating the result of the operation.","zh_CN":"操作结果的响应码。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The detailed message or error description for the operation result.","zh_CN":"操作结果的详细信息或错误提示。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddImageConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigResponse) GoString() string {
  return s.String()
}

func (s *AddImageConfigResponse) SetCode(v int) *AddImageConfigResponse {
  s.Code = &v
  return s
}

func (s *AddImageConfigResponse) SetMessage(v string) *AddImageConfigResponse {
  s.Message = &v
  return s
}

type AddImageConfigResponseHeader struct {
}

func (s AddImageConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddImageConfigResponseHeader) GoString() string {
  return s.String()
}




type RagKnowledgeBaseUpdateServiceRequest struct {
  // {"en":"knowledge base id","zh_CN":"知识库 id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"knowledge base name, Maximum Length 100","zh_CN":"知识库名称，最大长度100"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"knowledge base description, Maximum Length 500","zh_CN":"知识库描述，最大长度 500"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s RagKnowledgeBaseUpdateServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseUpdateServiceRequest) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseUpdateServiceRequest) SetId(v string) *RagKnowledgeBaseUpdateServiceRequest {
  s.Id = &v
  return s
}

func (s *RagKnowledgeBaseUpdateServiceRequest) SetName(v string) *RagKnowledgeBaseUpdateServiceRequest {
  s.Name = &v
  return s
}

func (s *RagKnowledgeBaseUpdateServiceRequest) SetDescription(v string) *RagKnowledgeBaseUpdateServiceRequest {
  s.Description = &v
  return s
}

type RagKnowledgeBaseUpdateServiceRequestHeader struct {
}

func (s RagKnowledgeBaseUpdateServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseUpdateServiceRequestHeader) GoString() string {
  return s.String()
}

type RagKnowledgeBaseUpdateServicePaths struct {
}

func (s RagKnowledgeBaseUpdateServicePaths) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseUpdateServicePaths) GoString() string {
  return s.String()
}

type RagKnowledgeBaseUpdateServiceParameters struct {
}

func (s RagKnowledgeBaseUpdateServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseUpdateServiceParameters) GoString() string {
  return s.String()
}

type RagKnowledgeBaseUpdateServiceResponse struct {
  // {"en":"code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RagKnowledgeBaseUpdateServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseUpdateServiceResponse) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseUpdateServiceResponse) SetCode(v string) *RagKnowledgeBaseUpdateServiceResponse {
  s.Code = &v
  return s
}

func (s *RagKnowledgeBaseUpdateServiceResponse) SetMessage(v string) *RagKnowledgeBaseUpdateServiceResponse {
  s.Message = &v
  return s
}

type RagKnowledgeBaseUpdateServiceResponseHeader struct {
}

func (s RagKnowledgeBaseUpdateServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseUpdateServiceResponseHeader) GoString() string {
  return s.String()
}




type UpdatefileRequest struct {
  // {"en":"knowledge id", "zh_CN":"知识库id"}
  KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty" require:"true"`
  // {"en":"file id", "zh_CN":"文件id"}
  FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty" require:"true"`
  // {"en":"file name", "zh_CN":"文件名文件名，最大长度255"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"file content after Base64 encode", "zh_CN":"文件内容Base64 encode"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s UpdatefileRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatefileRequest) GoString() string {
  return s.String()
}

func (s *UpdatefileRequest) SetKnowledgeBaseId(v string) *UpdatefileRequest {
  s.KnowledgeBaseId = &v
  return s
}

func (s *UpdatefileRequest) SetFileId(v string) *UpdatefileRequest {
  s.FileId = &v
  return s
}

func (s *UpdatefileRequest) SetFileName(v string) *UpdatefileRequest {
  s.FileName = &v
  return s
}

func (s *UpdatefileRequest) SetContent(v string) *UpdatefileRequest {
  s.Content = &v
  return s
}

type UpdatefileResponse struct {
  // {"en":"code", "zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdatefileResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatefileResponse) GoString() string {
  return s.String()
}

func (s *UpdatefileResponse) SetCode(v int) *UpdatefileResponse {
  s.Code = &v
  return s
}

func (s *UpdatefileResponse) SetMessage(v string) *UpdatefileResponse {
  s.Message = &v
  return s
}

type UpdatefilePaths struct {
}

func (s UpdatefilePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatefilePaths) GoString() string {
  return s.String()
}

type UpdatefileParameters struct {
}

func (s UpdatefileParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatefileParameters) GoString() string {
  return s.String()
}

type UpdatefileRequestHeader struct {
}

func (s UpdatefileRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatefileRequestHeader) GoString() string {
  return s.String()
}

type UpdatefileResponseHeader struct {
}

func (s UpdatefileResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatefileResponseHeader) GoString() string {
  return s.String()
}




type DeleteWaitingRoomConnectorRequest struct {
}

func (s DeleteWaitingRoomConnectorRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorRequest) GoString() string {
  return s.String()
}

type DeleteWaitingRoomConnectorRequestHeader struct {
}

func (s DeleteWaitingRoomConnectorRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorRequestHeader) GoString() string {
  return s.String()
}

type DeleteWaitingRoomConnectorPaths struct {
  // {"en":"The ID of the connector instance to be deleted","zh_CN":"待删除的连接器实例ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteWaitingRoomConnectorPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorPaths) GoString() string {
  return s.String()
}

func (s *DeleteWaitingRoomConnectorPaths) SetId(v string) *DeleteWaitingRoomConnectorPaths {
  s.Id = &v
  return s
}

type DeleteWaitingRoomConnectorParameters struct {
}

func (s DeleteWaitingRoomConnectorParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorParameters) GoString() string {
  return s.String()
}

type DeleteWaitingRoomConnectorResponse struct {
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"响应数据"}
  Data *DeleteWaitingRoomConnectorResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s DeleteWaitingRoomConnectorResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorResponse) GoString() string {
  return s.String()
}

func (s *DeleteWaitingRoomConnectorResponse) SetCode(v string) *DeleteWaitingRoomConnectorResponse {
  s.Code = &v
  return s
}

func (s *DeleteWaitingRoomConnectorResponse) SetMessage(v string) *DeleteWaitingRoomConnectorResponse {
  s.Message = &v
  return s
}

func (s *DeleteWaitingRoomConnectorResponse) SetData(v *DeleteWaitingRoomConnectorResponseData) *DeleteWaitingRoomConnectorResponse {
  s.Data = v
  return s
}

type DeleteWaitingRoomConnectorResponseData struct {
}

func (s DeleteWaitingRoomConnectorResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorResponseData) GoString() string {
  return s.String()
}

type DeleteWaitingRoomConnectorResponseHeader struct {
}

func (s DeleteWaitingRoomConnectorResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWaitingRoomConnectorResponseHeader) GoString() string {
  return s.String()
}




type QueryAiGatewayUsageRequest struct {
  // {"en":"The English name (cust_en_name) of a sub-client. When a merged account needs to view a sub-client's information, this parameter is required.","zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名。"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date. The date format is yyyy-mm-dd. If not provided or empty, it defaults to the current day.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"The start date for the query. Format: yyyy-mm-dd. This parameter must be used with 'enddate'. If the 'date' parameter is also provided, this parameter will be ignored.","zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合，若存在date参数，则该参数无效。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"The end date for the query. Format: yyyy-mm-dd. This parameter must be used with 'startdate'. If the 'date' parameter is also provided, this parameter will be ignored.","zh_CN":"查询的结束日期，日期格式为yyyy-mm-dd；此参数需与startdate参数配合，若存在date参数，则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"The abbreviation for the accelerated region to query. If there are multiple regions, use ';' as a separator. For example, 'region=cn;apac' queries data for the CN and APAC regions. If not specified or empty, it defaults to all regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The acceleration type, for example, 'web'. If multiple types are provided, use ';' as a separator. If not specified or if the value is 'all', it refers to all acceleration types.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型。"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The format of the response. Supported values are 'xml' and 'json', with 'xml' as the default.","zh_CN":"返回结果格式，支持格式为xml和json，默认为xml。"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"AI service provider. For multiple values, please separate them with an English semicolon ';'. If not filled in or set to 'all', it refers to all providers.","zh_CN":"AI服务提供商。多个请用英文分号“;”分隔开，不填或值为all表示所有提供商。"}
  Aiprovider *string `json:"aiprovider,omitempty" xml:"aiprovider,omitempty"`
  // {"en":"AI Gateway. For multiple types, please separate them with a semicolon ';'. If left blank or set to 'all', it will indicate all gateways.","zh_CN":"AI网关。多个请用英文分号“;”分隔开，不填或值为all表示所有网关。"}
  Aigateway *string `json:"aigateway,omitempty" xml:"aigateway,omitempty"`
  // {"en":"Model. If specifying multiple types, please separate them with an English semicolon ';'. Leaving this field empty or setting it to 'all' indicates all models.","zh_CN":"模型，多个请用英文分号“;”分隔开，不填或值为all表示所有模型。"}
  Model *string `json:"model,omitempty" xml:"model,omitempty"`
  // {"en":"Deep thinking status: '1' indicates yes, '0' indicates no, '-' represents historical data. For multiple entries, please separate using a semicolon ';'. Leave blank or use 'all' to indicate all types.","zh_CN":"是否深度思考，'1'表示是，'0'表示否，'-'表示历史数据，多个请用英文分号“;”分隔开，不填或值为all表示所有类型。"}
  Reasonflag *string `json:"reasonflag,omitempty" xml:"reasonflag,omitempty"`
  // {"en":"The GMT time zone. The parameter format is GMT+09:00 for the East Ninth Zone and GMT-09:00 for the West Ninth Zone. If not provided, it defaults to the local time zone (GMT+08:00, East Eighth Zone).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Whether to aggregate in a specific way. Format: number_day|hour. For example, '2_day' means aggregation by every 2 days. If not specified, data will be returned in its original granularity (5_min) by default.","zh_CN":"是否按照特定方式聚合,格式 :  数字_day|hour，例如‘2_day’表示按照2天聚合, 不填默认以底层数据粒度返回 （5_min）。"}
  ReturnType *string `json:"returnType,omitempty" xml:"returnType,omitempty"`
  // {"en":"Datatype\n- 0: Total number of requests\n- 1: Total number of tokens\n- 2: Number of input tokens\n- 3: Number of output tokens\n- 4: Number of generated images\n- 5: Input audio duration\n- 6: Number of input characters\n- 7: Number of input images\n- Default: all (return all)","zh_CN":"数据类型\n- 0：总请求数\n- 1：总Tokens数\n- 2：输入Tokens数\n- 3：输出Tokens数\n- 4：生成图片数\n- 5：输入语音时长\n- 6：输入字符数 \n- 7：输入图片数\n- 默认：all（全部都返回）"}
  Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
}

func (s QueryAiGatewayUsageRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageRequest) GoString() string {
  return s.String()
}

func (s *QueryAiGatewayUsageRequest) SetCust(v string) *QueryAiGatewayUsageRequest {
  s.Cust = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetDate(v string) *QueryAiGatewayUsageRequest {
  s.Date = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetStartdate(v string) *QueryAiGatewayUsageRequest {
  s.Startdate = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetEnddate(v string) *QueryAiGatewayUsageRequest {
  s.Enddate = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetRegion(v string) *QueryAiGatewayUsageRequest {
  s.Region = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetAccetype(v string) *QueryAiGatewayUsageRequest {
  s.Accetype = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetDataformat(v string) *QueryAiGatewayUsageRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetAiprovider(v string) *QueryAiGatewayUsageRequest {
  s.Aiprovider = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetAigateway(v string) *QueryAiGatewayUsageRequest {
  s.Aigateway = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetModel(v string) *QueryAiGatewayUsageRequest {
  s.Model = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetReasonflag(v string) *QueryAiGatewayUsageRequest {
  s.Reasonflag = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetTimezone(v string) *QueryAiGatewayUsageRequest {
  s.Timezone = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetReturnType(v string) *QueryAiGatewayUsageRequest {
  s.ReturnType = &v
  return s
}

func (s *QueryAiGatewayUsageRequest) SetDatatype(v string) *QueryAiGatewayUsageRequest {
  s.Datatype = &v
  return s
}

type QueryAiGatewayUsageRequestHeader struct {
}

func (s QueryAiGatewayUsageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageRequestHeader) GoString() string {
  return s.String()
}

type QueryAiGatewayUsagePaths struct {
}

func (s QueryAiGatewayUsagePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsagePaths) GoString() string {
  return s.String()
}

type QueryAiGatewayUsageParameters struct {
}

func (s QueryAiGatewayUsageParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageParameters) GoString() string {
  return s.String()
}

type QueryAiGatewayUsageResponse struct {
  // {"en":"Result data from the provider.","zh_CN":"提供方返回的结果数据。"}
  Provider *QueryAiGatewayUsageResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryAiGatewayUsageResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageResponse) GoString() string {
  return s.String()
}

func (s *QueryAiGatewayUsageResponse) SetProvider(v *QueryAiGatewayUsageResponseProvider) *QueryAiGatewayUsageResponse {
  s.Provider = v
  return s
}

type QueryAiGatewayUsageResponseProvider struct {
  // {"en":"Tenant name.","zh_CN":"租户名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"defaultValue":"sharkletAIgateway","en":"The business type of the interface.","zh_CN":"接口的业务类型。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"AI Gateway Data","zh_CN":"AI网关数据"}
  Date *QueryAiGatewayUsageResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryAiGatewayUsageResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryAiGatewayUsageResponseProvider) SetName(v string) *QueryAiGatewayUsageResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProvider) SetType(v string) *QueryAiGatewayUsageResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProvider) SetDate(v *QueryAiGatewayUsageResponseProviderDate) *QueryAiGatewayUsageResponseProvider {
  s.Date = v
  return s
}

type QueryAiGatewayUsageResponseProviderDate struct {
  // {"en":"Start date.","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"End date.","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"Total number of requests","zh_CN":"总请求数"}
  TotalVisit *string `json:"totalVisit,omitempty" xml:"totalVisit,omitempty" require:"true"`
  // {"en":"Total number of tokens","zh_CN":"总Tokens数"}
  TotalTokens *string `json:"totalTokens,omitempty" xml:"totalTokens,omitempty" require:"true"`
  // {"en":"Total number of input tokens","zh_CN":"总输入Tokens数"}
  TotalPromptTokens *string `json:"totalPromptTokens,omitempty" xml:"totalPromptTokens,omitempty" require:"true"`
  // {"en":"Total number of output tokens","zh_CN":"总输出Tokens"}
  TotalCompleteTokens *string `json:"totalCompleteTokens,omitempty" xml:"totalCompleteTokens,omitempty" require:"true"`
  // {"en":"Total number of generated images","zh_CN":"总生成图片数"}
  TotalItemNum *string `json:"totalItemNum,omitempty" xml:"totalItemNum,omitempty" require:"true"`
  // {"en":"Total input audio duration","zh_CN":"总输入语音时长"}
  TotalDuration *string `json:"totalDuration,omitempty" xml:"totalDuration,omitempty" require:"true"`
  // {"en":"Total number of input characters","zh_CN":"总输入字符数"}
  TotalInputCharacters *string `json:"totalInputCharacters,omitempty" xml:"totalInputCharacters,omitempty" require:"true"`
  // {"en":"Total number of input images","zh_CN":"总输入图片数"}
  TotalInputImageNum *string `json:"totalInputImageNum,omitempty" xml:"totalInputImageNum,omitempty" require:"true"`
  // {"en":"Detailed data","zh_CN":"明细数据"}
  ChartDataList []*QueryAiGatewayUsageResponseProviderDateChartDataList `json:"chartDataList,omitempty" xml:"chartDataList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAiGatewayUsageResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetStartdate(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetEnddate(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalVisit(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalVisit = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalTokens(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalTokens = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalPromptTokens(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalPromptTokens = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalCompleteTokens(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalCompleteTokens = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalItemNum(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalItemNum = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalDuration(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalDuration = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalInputCharacters(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalInputCharacters = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetTotalInputImageNum(v string) *QueryAiGatewayUsageResponseProviderDate {
  s.TotalInputImageNum = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDate) SetChartDataList(v []*QueryAiGatewayUsageResponseProviderDateChartDataList) *QueryAiGatewayUsageResponseProviderDate {
  s.ChartDataList = v
  return s
}

type QueryAiGatewayUsageResponseProviderDateChartDataList struct     {
  // {"en":"Detailed Data Type\n- totalVisit: Total number of requests\n- totalTokens: Total number of Tokens\n- promptTokens: Number of input Tokens\n- completeTokens: Number of output Tokens\n- itemNum: Number of generated images\n- duration: Duration of input audio (milliseconds)\n- inputCharacters: Number of input characters\n- inputImageNum: Number of input images","zh_CN":"明细数据类型\n- totalVisit：总请求数\n- totalTokens：总Tokens数\n- promptTokens：输入Tokens数\n- completeTokens：输出Tokens数\n- itemNum：生成图片数\n- duration：输入语音时长（毫秒）\n- inputCharacters：输入字符数\n- inputImageNum：输入图片数","exampleValue":"totalVisit,totalTokens,promptTokens,completeTokens,itemNum,duration,inputCharacters,inputImageNum"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Detailed data","zh_CN":"明细数据"}
  Data []*QueryAiGatewayUsageResponseProviderDateChartDataListData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAiGatewayUsageResponseProviderDateChartDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageResponseProviderDateChartDataList) GoString() string {
  return s.String()
}

func (s *QueryAiGatewayUsageResponseProviderDateChartDataList) SetName(v string) *QueryAiGatewayUsageResponseProviderDateChartDataList {
  s.Name = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDateChartDataList) SetData(v []*QueryAiGatewayUsageResponseProviderDateChartDataListData) *QueryAiGatewayUsageResponseProviderDateChartDataList {
  s.Data = v
  return s
}

type QueryAiGatewayUsageResponseProviderDateChartDataListData struct     {
  // {"en":"time","zh_CN":"时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"value","zh_CN":"明细值"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QueryAiGatewayUsageResponseProviderDateChartDataListData) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageResponseProviderDateChartDataListData) GoString() string {
  return s.String()
}

func (s *QueryAiGatewayUsageResponseProviderDateChartDataListData) SetTime(v string) *QueryAiGatewayUsageResponseProviderDateChartDataListData {
  s.Time = &v
  return s
}

func (s *QueryAiGatewayUsageResponseProviderDateChartDataListData) SetText(v string) *QueryAiGatewayUsageResponseProviderDateChartDataListData {
  s.Text = &v
  return s
}

type QueryAiGatewayUsageResponseHeader struct {
}

func (s QueryAiGatewayUsageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAiGatewayUsageResponseHeader) GoString() string {
  return s.String()
}




