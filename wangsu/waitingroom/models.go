package waitingroom

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryDeploymentRecordsRequest struct {
}

func (s QueryDeploymentRecordsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsRequest) GoString() string {
  return s.String()
}

type QueryDeploymentRecordsRequestHeader struct {
}

func (s QueryDeploymentRecordsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsRequestHeader) GoString() string {
  return s.String()
}

type QueryDeploymentRecordsPaths struct {
}

func (s QueryDeploymentRecordsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsPaths) GoString() string {
  return s.String()
}

type QueryDeploymentRecordsParameters struct {
  // {"defaultValue":"1","en":"page number","zh_CN":"页码"}
  PageNo *string `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"defaultValue":"10","en":"page size","zh_CN":"页大小"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"instance type","zh_CN":"实例类型"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
}

func (s QueryDeploymentRecordsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsParameters) GoString() string {
  return s.String()
}

func (s *QueryDeploymentRecordsParameters) SetPageNo(v string) *QueryDeploymentRecordsParameters {
  s.PageNo = &v
  return s
}

func (s *QueryDeploymentRecordsParameters) SetPageSize(v string) *QueryDeploymentRecordsParameters {
  s.PageSize = &v
  return s
}

func (s *QueryDeploymentRecordsParameters) SetInstanceType(v string) *QueryDeploymentRecordsParameters {
  s.InstanceType = &v
  return s
}

type QueryDeploymentRecordsResponse struct {
  // {"en":"response code","zh_CN":"响应编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"数据"}
  Data *QueryDeploymentRecordsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentRecordsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsResponse) GoString() string {
  return s.String()
}

func (s *QueryDeploymentRecordsResponse) SetCode(v string) *QueryDeploymentRecordsResponse {
  s.Code = &v
  return s
}

func (s *QueryDeploymentRecordsResponse) SetMessage(v string) *QueryDeploymentRecordsResponse {
  s.Message = &v
  return s
}

func (s *QueryDeploymentRecordsResponse) SetData(v *QueryDeploymentRecordsResponseData) *QueryDeploymentRecordsResponse {
  s.Data = v
  return s
}

type QueryDeploymentRecordsResponseData struct {
  // {"en":"total records","zh_CN":"总记录数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"deployment task list","zh_CN":"部署任务列表"}
  Deployments []*QueryDeploymentRecordsResponseDataDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDeploymentRecordsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsResponseData) GoString() string {
  return s.String()
}

func (s *QueryDeploymentRecordsResponseData) SetTotal(v int) *QueryDeploymentRecordsResponseData {
  s.Total = &v
  return s
}

func (s *QueryDeploymentRecordsResponseData) SetDeployments(v []*QueryDeploymentRecordsResponseDataDeployments) *QueryDeploymentRecordsResponseData {
  s.Deployments = v
  return s
}

type QueryDeploymentRecordsResponseDataDeployments struct     {
  // {"en":"deploy task id","zh_CN":"部署任务ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"deploy time","zh_CN":"部署时间（时间戳）"}
  DeployTime *int `json:"deployTime,omitempty" xml:"deployTime,omitempty" require:"true"`
  // {"en":"deploy version","zh_CN":"部署版本"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"remark","zh_CN":"备注说明"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"deploy status","zh_CN":"部署状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"instance type","zh_CN":"实例类型"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty" require:"true"`
  // {"en":"deploy content","zh_CN":"部署内容详情"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s QueryDeploymentRecordsResponseDataDeployments) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsResponseDataDeployments) GoString() string {
  return s.String()
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetId(v string) *QueryDeploymentRecordsResponseDataDeployments {
  s.Id = &v
  return s
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetDeployTime(v int) *QueryDeploymentRecordsResponseDataDeployments {
  s.DeployTime = &v
  return s
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetVersion(v string) *QueryDeploymentRecordsResponseDataDeployments {
  s.Version = &v
  return s
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetRemark(v string) *QueryDeploymentRecordsResponseDataDeployments {
  s.Remark = &v
  return s
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetStatus(v int) *QueryDeploymentRecordsResponseDataDeployments {
  s.Status = &v
  return s
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetInstanceType(v string) *QueryDeploymentRecordsResponseDataDeployments {
  s.InstanceType = &v
  return s
}

func (s *QueryDeploymentRecordsResponseDataDeployments) SetContent(v string) *QueryDeploymentRecordsResponseDataDeployments {
  s.Content = &v
  return s
}

type QueryDeploymentRecordsResponseHeader struct {
}

func (s QueryDeploymentRecordsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentRecordsResponseHeader) GoString() string {
  return s.String()
}




type QueryDeploymentDetailRequest struct {
}

func (s QueryDeploymentDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailRequest) GoString() string {
  return s.String()
}

type QueryDeploymentDetailRequestHeader struct {
}

func (s QueryDeploymentDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailRequestHeader) GoString() string {
  return s.String()
}

type QueryDeploymentDetailPaths struct {
  // {"en":"The unique identifier of the deployment task to query","zh_CN":"需要查询的部署任务的唯一标识符"}
  DeployId *string `json:"deployId,omitempty" xml:"deployId,omitempty" require:"true"`
}

func (s QueryDeploymentDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailPaths) GoString() string {
  return s.String()
}

func (s *QueryDeploymentDetailPaths) SetDeployId(v string) *QueryDeploymentDetailPaths {
  s.DeployId = &v
  return s
}

type QueryDeploymentDetailParameters struct {
}

func (s QueryDeploymentDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailParameters) GoString() string {
  return s.String()
}

type QueryDeploymentDetailResponse struct {
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Business data object","zh_CN":"业务数据"}
  Data *QueryDeploymentDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryDeploymentDetailResponse) SetCode(v string) *QueryDeploymentDetailResponse {
  s.Code = &v
  return s
}

func (s *QueryDeploymentDetailResponse) SetMessage(v string) *QueryDeploymentDetailResponse {
  s.Message = &v
  return s
}

func (s *QueryDeploymentDetailResponse) SetData(v *QueryDeploymentDetailResponseData) *QueryDeploymentDetailResponse {
  s.Data = v
  return s
}

type QueryDeploymentDetailResponseData struct {
  // {"en":"deployment task detail","zh_CN":"部署任务详情对象"}
  Deployment *QueryDeploymentDetailResponseDataDeployment `json:"deployment,omitempty" xml:"deployment,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailResponseData) GoString() string {
  return s.String()
}

func (s *QueryDeploymentDetailResponseData) SetDeployment(v *QueryDeploymentDetailResponseDataDeployment) *QueryDeploymentDetailResponseData {
  s.Deployment = v
  return s
}

type QueryDeploymentDetailResponseDataDeployment struct {
  // {"en":"deployment id","zh_CN":"部署任务ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Deployment time (timestamp)","zh_CN":"部署时间（时间戳）"}
  DeployTime *int `json:"deployTime,omitempty" xml:"deployTime,omitempty" require:"true"`
  // {"en":"deployment version","zh_CN":"部署版本"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"remark","zh_CN":"备注说明"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"deployment status","zh_CN":"部署状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"instance type","zh_CN":"实例类型"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty" require:"true"`
  // {"en":"deployment content detail","zh_CN":"部署内容详情"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s QueryDeploymentDetailResponseDataDeployment) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailResponseDataDeployment) GoString() string {
  return s.String()
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetId(v string) *QueryDeploymentDetailResponseDataDeployment {
  s.Id = &v
  return s
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetDeployTime(v int) *QueryDeploymentDetailResponseDataDeployment {
  s.DeployTime = &v
  return s
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetVersion(v string) *QueryDeploymentDetailResponseDataDeployment {
  s.Version = &v
  return s
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetRemark(v string) *QueryDeploymentDetailResponseDataDeployment {
  s.Remark = &v
  return s
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetStatus(v int) *QueryDeploymentDetailResponseDataDeployment {
  s.Status = &v
  return s
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetInstanceType(v string) *QueryDeploymentDetailResponseDataDeployment {
  s.InstanceType = &v
  return s
}

func (s *QueryDeploymentDetailResponseDataDeployment) SetContent(v string) *QueryDeploymentDetailResponseDataDeployment {
  s.Content = &v
  return s
}

type QueryDeploymentDetailResponseHeader struct {
}

func (s QueryDeploymentDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentDetailResponseHeader) GoString() string {
  return s.String()
}




type DeployInstanceRequest struct {
  // {"en":"The ID of the entity to deploy","zh_CN":"要部署的实体ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"The entity type to deploy, currently only WaitingRoomConnector","zh_CN":"要部署的实体类型,当前只有 WaitingRoomConnector"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty" require:"true"`
}

func (s DeployInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployInstanceRequest) GoString() string {
  return s.String()
}

func (s *DeployInstanceRequest) SetInstanceId(v string) *DeployInstanceRequest {
  s.InstanceId = &v
  return s
}

func (s *DeployInstanceRequest) SetInstanceType(v string) *DeployInstanceRequest {
  s.InstanceType = &v
  return s
}

type DeployInstanceRequestHeader struct {
}

func (s DeployInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployInstanceRequestHeader) GoString() string {
  return s.String()
}

type DeployInstancePaths struct {
}

func (s DeployInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s DeployInstancePaths) GoString() string {
  return s.String()
}

type DeployInstanceParameters struct {
}

func (s DeployInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployInstanceParameters) GoString() string {
  return s.String()
}

type DeployInstanceResponse struct {
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data","zh_CN":"响应数据"}
  Data *DeployInstanceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s DeployInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployInstanceResponse) GoString() string {
  return s.String()
}

func (s *DeployInstanceResponse) SetCode(v string) *DeployInstanceResponse {
  s.Code = &v
  return s
}

func (s *DeployInstanceResponse) SetMessage(v string) *DeployInstanceResponse {
  s.Message = &v
  return s
}

func (s *DeployInstanceResponse) SetData(v *DeployInstanceResponseData) *DeployInstanceResponse {
  s.Data = v
  return s
}

type DeployInstanceResponseData struct {
  // {"en":"Deployment ID","zh_CN":"部署id"}
  DeployId *string `json:"deployId,omitempty" xml:"deployId,omitempty" require:"true"`
}

func (s DeployInstanceResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeployInstanceResponseData) GoString() string {
  return s.String()
}

func (s *DeployInstanceResponseData) SetDeployId(v string) *DeployInstanceResponseData {
  s.DeployId = &v
  return s
}

type DeployInstanceResponseHeader struct {
}

func (s DeployInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployInstanceResponseHeader) GoString() string {
  return s.String()
}




