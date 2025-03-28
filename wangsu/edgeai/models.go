package edgeai

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

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
  Data []*RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *RagKnowledgeBaseListServiceResponse) SetData(v []*RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) *RagKnowledgeBaseListServiceResponse {
  s.Data = v
  return s
}

func (s *RagKnowledgeBaseListServiceResponse) SetMessage(v string) *RagKnowledgeBaseListServiceResponse {
  s.Message = &v
  return s
}

type RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData struct     {
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

func (s RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) GoString() string {
  return s.String()
}

func (s *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) SetLastUpdated(v string) *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData {
  s.LastUpdated = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) SetSize(v int) *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData {
  s.Size = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) SetCreateTime(v string) *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData {
  s.CreateTime = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) SetName(v string) *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData {
  s.Name = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) SetDescription(v string) *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData {
  s.Description = &v
  return s
}

func (s *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData) SetId(v string) *RagKnowledgeBaseListServiceRagKnowledgeBaseListServiceResponseData {
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




