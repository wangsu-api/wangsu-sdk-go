package imagemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteImageTagRequest struct {
}

func (s DeleteImageTagRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteImageTagRequest) GoString() string {
  return s.String()
}

type DeleteImageTagResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"result: true/false", "zh_CN":"结果:true/false"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteImageTagResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteImageTagResponse) GoString() string {
  return s.String()
}

func (s *DeleteImageTagResponse) SetCode(v int64) *DeleteImageTagResponse {
  s.Code = &v
  return s
}

func (s *DeleteImageTagResponse) SetMsg(v string) *DeleteImageTagResponse {
  s.Msg = &v
  return s
}

func (s *DeleteImageTagResponse) SetRequestId(v string) *DeleteImageTagResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteImageTagResponse) SetData(v bool) *DeleteImageTagResponse {
  s.Data = &v
  return s
}

type DeleteImageTagPaths struct {
}

func (s DeleteImageTagPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteImageTagPaths) GoString() string {
  return s.String()
}

type DeleteImageTagParameters struct {
  // {"en":"project name", "zh_CN":"组织名称"}
  ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty" require:"true"`
  // {"en":"image tags,split with \",\"", "zh_CN":"镜像版本号,\",\"分隔"}
  Tags *string `json:"tags,omitempty" xml:"tags,omitempty" require:"true"`
  // {"en":"image name,like: project/image", "zh_CN":"镜像名称,格式:project/image"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteImageTagParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteImageTagParameters) GoString() string {
  return s.String()
}

func (s *DeleteImageTagParameters) SetProjectName(v string) *DeleteImageTagParameters {
  s.ProjectName = &v
  return s
}

func (s *DeleteImageTagParameters) SetTags(v string) *DeleteImageTagParameters {
  s.Tags = &v
  return s
}

func (s *DeleteImageTagParameters) SetName(v string) *DeleteImageTagParameters {
  s.Name = &v
  return s
}

type DeleteImageTagRequestHeader struct {
}

func (s DeleteImageTagRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteImageTagRequestHeader) GoString() string {
  return s.String()
}

type DeleteImageTagResponseHeader struct {
}

func (s DeleteImageTagResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteImageTagResponseHeader) GoString() string {
  return s.String()
}




type GetImageDetailRequest struct {
}

func (s GetImageDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailRequest) GoString() string {
  return s.String()
}

type GetImageDetailResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"tags list", "zh_CN":"tags列表"}
  Data *GetImageDetailImageInfo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetImageDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailResponse) GoString() string {
  return s.String()
}

func (s *GetImageDetailResponse) SetCode(v int64) *GetImageDetailResponse {
  s.Code = &v
  return s
}

func (s *GetImageDetailResponse) SetMsg(v string) *GetImageDetailResponse {
  s.Msg = &v
  return s
}

func (s *GetImageDetailResponse) SetRequestId(v string) *GetImageDetailResponse {
  s.RequestId = &v
  return s
}

func (s *GetImageDetailResponse) SetData(v *GetImageDetailImageInfo) *GetImageDetailResponse {
  s.Data = v
  return s
}

type GetImageDetailPaths struct {
}

func (s GetImageDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailPaths) GoString() string {
  return s.String()
}

type GetImageDetailParameters struct {
  // {"en":"image name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"project name", "zh_CN":"组织名称"}
  ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty" require:"true"`
}

func (s GetImageDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailParameters) GoString() string {
  return s.String()
}

func (s *GetImageDetailParameters) SetName(v string) *GetImageDetailParameters {
  s.Name = &v
  return s
}

func (s *GetImageDetailParameters) SetProjectName(v string) *GetImageDetailParameters {
  s.ProjectName = &v
  return s
}

type GetImageDetailRequestHeader struct {
}

func (s GetImageDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailRequestHeader) GoString() string {
  return s.String()
}

type GetImageDetailResponseHeader struct {
}

func (s GetImageDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailResponseHeader) GoString() string {
  return s.String()
}

type GetImageDetailImageInfo struct {
  // {"en":"image id", "zh_CN":"image id"}
  ImageId *uint64 `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"project id", "zh_CN":"组织 id"}
  ProjectId *uint64 `json:"projectId,omitempty" xml:"projectId,omitempty" require:"true"`
  // {"en":"image name", "zh_CN":"镜像名称"}
  ImageName *string `json:"imageName,omitempty" xml:"imageName,omitempty" require:"true"`
  // {"en":"project name", "zh_CN":"组织名称"}
  ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty" require:"true"`
  // {"en":"tag count", "zh_CN":"tag 数量"}
  TagCount *uint64 `json:"tagCount,omitempty" xml:"tagCount,omitempty" require:"true"`
  // {"en":"project type", "zh_CN":"组织类别"}
  Type *uint64 `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"repository_plus", "zh_CN":"repository_plus"}
  Repository_plus *string `json:"repository_plus,omitempty" xml:"repository_plus,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"创建时间"}
  CreateTime *uint64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"update time", "zh_CN":"更新时间"}
  UpdateTime *uint64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"tag list", "zh_CN":"tag列表"}
  Tags []*GetImageDetailTag `json:"tags,omitempty" xml:"tags,omitempty" require:"true" type:"Repeated"`
}

func (s GetImageDetailImageInfo) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailImageInfo) GoString() string {
  return s.String()
}

func (s *GetImageDetailImageInfo) SetImageId(v uint64) *GetImageDetailImageInfo {
  s.ImageId = &v
  return s
}

func (s *GetImageDetailImageInfo) SetProjectId(v uint64) *GetImageDetailImageInfo {
  s.ProjectId = &v
  return s
}

func (s *GetImageDetailImageInfo) SetImageName(v string) *GetImageDetailImageInfo {
  s.ImageName = &v
  return s
}

func (s *GetImageDetailImageInfo) SetProjectName(v string) *GetImageDetailImageInfo {
  s.ProjectName = &v
  return s
}

func (s *GetImageDetailImageInfo) SetTagCount(v uint64) *GetImageDetailImageInfo {
  s.TagCount = &v
  return s
}

func (s *GetImageDetailImageInfo) SetType(v uint64) *GetImageDetailImageInfo {
  s.Type = &v
  return s
}

func (s *GetImageDetailImageInfo) SetRepository_plus(v string) *GetImageDetailImageInfo {
  s.Repository_plus = &v
  return s
}

func (s *GetImageDetailImageInfo) SetCreateTime(v uint64) *GetImageDetailImageInfo {
  s.CreateTime = &v
  return s
}

func (s *GetImageDetailImageInfo) SetUpdateTime(v uint64) *GetImageDetailImageInfo {
  s.UpdateTime = &v
  return s
}

func (s *GetImageDetailImageInfo) SetTags(v []*GetImageDetailTag) *GetImageDetailImageInfo {
  s.Tags = v
  return s
}

type GetImageDetailTag struct {
  // {"en":"tag id", "zh_CN":"tag id"}
  TagId *string `json:"tagId,omitempty" xml:"tagId,omitempty" require:"true"`
  // {"en":"tag name", "zh_CN":"tag名称"}
  TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty" require:"true"`
  // {"en":"image size", "zh_CN":"镜像大小"}
  Size *string `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"description", "zh_CN":"描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"image address", "zh_CN":"镜像地址"}
  ImageAddr *string `json:"imageAddr,omitempty" xml:"imageAddr,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"创建时间"}
  CreateTime *uint64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"push time", "zh_CN":"推送时间"}
  PushTime *uint64 `json:"pushTime,omitempty" xml:"pushTime,omitempty" require:"true"`
}

func (s GetImageDetailTag) String() string {
  return tea.Prettify(s)
}

func (s GetImageDetailTag) GoString() string {
  return s.String()
}

func (s *GetImageDetailTag) SetTagId(v string) *GetImageDetailTag {
  s.TagId = &v
  return s
}

func (s *GetImageDetailTag) SetTagName(v string) *GetImageDetailTag {
  s.TagName = &v
  return s
}

func (s *GetImageDetailTag) SetSize(v string) *GetImageDetailTag {
  s.Size = &v
  return s
}

func (s *GetImageDetailTag) SetDescription(v string) *GetImageDetailTag {
  s.Description = &v
  return s
}

func (s *GetImageDetailTag) SetImageAddr(v string) *GetImageDetailTag {
  s.ImageAddr = &v
  return s
}

func (s *GetImageDetailTag) SetCreateTime(v uint64) *GetImageDetailTag {
  s.CreateTime = &v
  return s
}

func (s *GetImageDetailTag) SetPushTime(v uint64) *GetImageDetailTag {
  s.PushTime = &v
  return s
}




type PatchImagePullJobRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *PatchImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the PatchImagePullJobImagePullJob.", "zh_CN":"PatchImagePullJobImagePullJob 预期行为的规约。"}
  Spec *PatchImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s PatchImagePullJobRequest) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobRequest) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobRequest) SetApiVersion(v string) *PatchImagePullJobRequest {
  s.ApiVersion = &v
  return s
}

func (s *PatchImagePullJobRequest) SetKind(v string) *PatchImagePullJobRequest {
  s.Kind = &v
  return s
}

func (s *PatchImagePullJobRequest) SetMetadata(v *PatchImagePullJobObjectMeta) *PatchImagePullJobRequest {
  s.Metadata = v
  return s
}

func (s *PatchImagePullJobRequest) SetSpec(v *PatchImagePullJobImagePullJobSpec) *PatchImagePullJobRequest {
  s.Spec = v
  return s
}

type PatchImagePullJobResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"imagepulljob", "zh_CN":"imagepulljob"}
  Data *PatchImagePullJobImagePullJob `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PatchImagePullJobResponse) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobResponse) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobResponse) SetCode(v int64) *PatchImagePullJobResponse {
  s.Code = &v
  return s
}

func (s *PatchImagePullJobResponse) SetMsg(v string) *PatchImagePullJobResponse {
  s.Msg = &v
  return s
}

func (s *PatchImagePullJobResponse) SetRequestId(v string) *PatchImagePullJobResponse {
  s.RequestId = &v
  return s
}

func (s *PatchImagePullJobResponse) SetData(v *PatchImagePullJobImagePullJob) *PatchImagePullJobResponse {
  s.Data = v
  return s
}

type PatchImagePullJobPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"imagepulljob name", "zh_CN":"imagepulljob 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s PatchImagePullJobPaths) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobPaths) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobPaths) SetNamespace(v string) *PatchImagePullJobPaths {
  s.Namespace = &v
  return s
}

func (s *PatchImagePullJobPaths) SetName(v string) *PatchImagePullJobPaths {
  s.Name = &v
  return s
}

type PatchImagePullJobParameters struct {
}

func (s PatchImagePullJobParameters) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobParameters) GoString() string {
  return s.String()
}

type PatchImagePullJobRequestHeader struct {
}

func (s PatchImagePullJobRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobRequestHeader) GoString() string {
  return s.String()
}

type PatchImagePullJobResponseHeader struct {
}

func (s PatchImagePullJobResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobResponseHeader) GoString() string {
  return s.String()
}

type PatchImagePullJobImagePullJob struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *PatchImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the PatchImagePullJobImagePullJob.", "zh_CN":"PatchImagePullJobImagePullJob 预期行为的规约。"}
  Spec *PatchImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the PatchImagePullJobImagePullJob.", "zh_CN":"最近观测到的 PatchImagePullJobImagePullJob 状态。"}
  Status *PatchImagePullJobImagePullJobStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PatchImagePullJobImagePullJob) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobImagePullJob) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobImagePullJob) SetApiVersion(v string) *PatchImagePullJobImagePullJob {
  s.ApiVersion = &v
  return s
}

func (s *PatchImagePullJobImagePullJob) SetKind(v string) *PatchImagePullJobImagePullJob {
  s.Kind = &v
  return s
}

func (s *PatchImagePullJobImagePullJob) SetMetadata(v *PatchImagePullJobObjectMeta) *PatchImagePullJobImagePullJob {
  s.Metadata = v
  return s
}

func (s *PatchImagePullJobImagePullJob) SetSpec(v *PatchImagePullJobImagePullJobSpec) *PatchImagePullJobImagePullJob {
  s.Spec = v
  return s
}

func (s *PatchImagePullJobImagePullJob) SetStatus(v *PatchImagePullJobImagePullJobStatus) *PatchImagePullJobImagePullJob {
  s.Status = v
  return s
}

type PatchImagePullJobImagePullJobSpec struct {
  // {"en": "Image is the image to be pulled by the job", "zh_CN": "拉取镜像名"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en": "Parallelism is the requested parallelism, it can be set to any non-negative value. If it is unspecified, it defaults to 1. If it is specified as 0, then the Job is effectively paused until it is increased.The value range 0-10, +optional", "zh_CN": "并发拉取个数, 范围0-10"}
  Parallelism *PatchImagePullJobIntstrIntOrString `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
  // {"en": "PatchImagePullJobPullPolicy is an optional field to set parameters of the pulling task. If not specified, the system will use the default values.+optional", "zh_CN": "拉取策略"}
  PatchImagePullJobPullPolicy *PatchImagePullJobPullPolicy `json:"pullPolicy,omitempty" xml:"pullPolicy,omitempty"`
  // {"en": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling the image.If specified, these secrets will be passed to individual puller implementations for them to use.  For example,in the case of docker, only DockerConfig type secrets are honored.+optional", "zh_CN": "拉取镜像所需的密钥"}
  PullSecrets []*string `json:"pullSecrets,omitempty" xml:"pullSecrets,omitempty" type:"Repeated"`
}

func (s PatchImagePullJobImagePullJobSpec) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobImagePullJobSpec) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobImagePullJobSpec) SetImage(v string) *PatchImagePullJobImagePullJobSpec {
  s.Image = &v
  return s
}

func (s *PatchImagePullJobImagePullJobSpec) SetParallelism(v *PatchImagePullJobIntstrIntOrString) *PatchImagePullJobImagePullJobSpec {
  s.Parallelism = v
  return s
}

func (s *PatchImagePullJobImagePullJobSpec) SetPullPolicy(v *PatchImagePullJobPullPolicy) *PatchImagePullJobImagePullJobSpec {
  s.PatchImagePullJobPullPolicy = v
  return s
}

func (s *PatchImagePullJobImagePullJobSpec) SetPullSecrets(v []*string) *PatchImagePullJobImagePullJobSpec {
  s.PullSecrets = v
  return s
}

type PatchImagePullJobIntstrIntOrString struct {
  // {"en": "the integer value", "zh_CN": "整数值"}
  IntVal *int `json:"intVal,omitempty" xml:"intVal,omitempty"`
  // {"en": "the string value", "zh_CN": "字符串值"}
  StrVal *string `json:"strVal,omitempty" xml:"strVal,omitempty"`
  // {"en": "type", "zh_CN": "类型"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PatchImagePullJobIntstrIntOrString) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobIntstrIntOrString) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobIntstrIntOrString) SetIntVal(v int) *PatchImagePullJobIntstrIntOrString {
  s.IntVal = &v
  return s
}

func (s *PatchImagePullJobIntstrIntOrString) SetStrVal(v string) *PatchImagePullJobIntstrIntOrString {
  s.StrVal = &v
  return s
}

func (s *PatchImagePullJobIntstrIntOrString) SetType(v int) *PatchImagePullJobIntstrIntOrString {
  s.Type = &v
  return s
}

type PatchImagePullJobPullPolicy struct {
  // {"en": "Specifies the number of retries before marking the pulling task failed. Defaults to 3 +optional", "zh_CN": "backoff次数，默认3"}
  BackoffLimit *int `json:"backoffLimit,omitempty" xml:"backoffLimit,omitempty"`
  // {"en": "Specifies the timeout of the pulling task. Defaults to 600 +optional", "zh_CN": "拉取超时时间"}
  TimeoutSeconds *int `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s PatchImagePullJobPullPolicy) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobPullPolicy) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobPullPolicy) SetBackoffLimit(v int) *PatchImagePullJobPullPolicy {
  s.BackoffLimit = &v
  return s
}

func (s *PatchImagePullJobPullPolicy) SetTimeoutSeconds(v int) *PatchImagePullJobPullPolicy {
  s.TimeoutSeconds = &v
  return s
}

type PatchImagePullJobImagePullJobStatus struct {
  // {"en": "Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务完成时间"}
  CompletionTime *string `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
  // {"en": "The text prompt for job running status.+optional", "zh_CN": "状态消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // {"en": "Represents time when the job was acknowledged by the job controller.It is not guaranteed to be set in happens-before order across separate operations.It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en": "The rate of pulling tasks which reached phase Succeeded without not ready nodes. +optional", "zh_CN": "完成成功率"}
  SucceededRate *string `json:"succeededRate,omitempty" xml:"succeededRate,omitempty"`
}

func (s PatchImagePullJobImagePullJobStatus) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobImagePullJobStatus) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobImagePullJobStatus) SetCompletionTime(v string) *PatchImagePullJobImagePullJobStatus {
  s.CompletionTime = &v
  return s
}

func (s *PatchImagePullJobImagePullJobStatus) SetMessage(v string) *PatchImagePullJobImagePullJobStatus {
  s.Message = &v
  return s
}

func (s *PatchImagePullJobImagePullJobStatus) SetStartTime(v string) *PatchImagePullJobImagePullJobStatus {
  s.StartTime = &v
  return s
}

func (s *PatchImagePullJobImagePullJobStatus) SetSucceededRate(v string) *PatchImagePullJobImagePullJobStatus {
  s.SucceededRate = &v
  return s
}

type PatchImagePullJobObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*PatchImagePullJobOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*PatchImagePullJobManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s PatchImagePullJobObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobObjectMeta) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobObjectMeta) SetName(v string) *PatchImagePullJobObjectMeta {
  s.Name = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetGenerateName(v string) *PatchImagePullJobObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetNamespace(v string) *PatchImagePullJobObjectMeta {
  s.Namespace = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetSelfLink(v string) *PatchImagePullJobObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetUid(v string) *PatchImagePullJobObjectMeta {
  s.Uid = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetResourceVersion(v string) *PatchImagePullJobObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetGeneration(v int64) *PatchImagePullJobObjectMeta {
  s.Generation = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetCreationTimestamp(v string) *PatchImagePullJobObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetDeletionTimestamp(v string) *PatchImagePullJobObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetDeletionGracePeriodSeconds(v int64) *PatchImagePullJobObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetLabels(v map[string]*string) *PatchImagePullJobObjectMeta {
  s.Labels = v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetAnnotations(v map[string]*string) *PatchImagePullJobObjectMeta {
  s.Annotations = v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetOwnerReferences(v []*PatchImagePullJobOwnerReference) *PatchImagePullJobObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetFinalizers(v []*string) *PatchImagePullJobObjectMeta {
  s.Finalizers = v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetClusterName(v string) *PatchImagePullJobObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *PatchImagePullJobObjectMeta) SetManagedFields(v []*PatchImagePullJobManagedFieldsEntry) *PatchImagePullJobObjectMeta {
  s.ManagedFields = v
  return s
}

type PatchImagePullJobManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this PatchImagePullJobManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'PatchImagePullJobFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“PatchImagePullJobFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"PatchImagePullJobFieldsV1 holds the first JSON version format as described in the 'PatchImagePullJobFieldsV1' type", "zh_CN":"PatchImagePullJobFieldsV1 包含类型 “PatchImagePullJobFieldsV1” 中描述的第一个 JSON 版本格式"}
  PatchImagePullJobFieldsV1 *PatchImagePullJobFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s PatchImagePullJobManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobManagedFieldsEntry) SetManager(v string) *PatchImagePullJobManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *PatchImagePullJobManagedFieldsEntry) SetOperation(v string) *PatchImagePullJobManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *PatchImagePullJobManagedFieldsEntry) SetApiVersion(v string) *PatchImagePullJobManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *PatchImagePullJobManagedFieldsEntry) SetTime(v string) *PatchImagePullJobManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *PatchImagePullJobManagedFieldsEntry) SetFieldsType(v string) *PatchImagePullJobManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *PatchImagePullJobManagedFieldsEntry) SetFieldsV1(v *PatchImagePullJobFieldsV1) *PatchImagePullJobManagedFieldsEntry {
  s.PatchImagePullJobFieldsV1 = v
  return s
}

func (s *PatchImagePullJobManagedFieldsEntry) SetSubresource(v string) *PatchImagePullJobManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type PatchImagePullJobFieldsV1 struct {
}

func (s PatchImagePullJobFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobFieldsV1) GoString() string {
  return s.String()
}

type PatchImagePullJobOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s PatchImagePullJobOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s PatchImagePullJobOwnerReference) GoString() string {
  return s.String()
}

func (s *PatchImagePullJobOwnerReference) SetApiVersion(v string) *PatchImagePullJobOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *PatchImagePullJobOwnerReference) SetKind(v string) *PatchImagePullJobOwnerReference {
  s.Kind = &v
  return s
}

func (s *PatchImagePullJobOwnerReference) SetName(v string) *PatchImagePullJobOwnerReference {
  s.Name = &v
  return s
}

func (s *PatchImagePullJobOwnerReference) SetUid(v string) *PatchImagePullJobOwnerReference {
  s.Uid = &v
  return s
}

func (s *PatchImagePullJobOwnerReference) SetController(v bool) *PatchImagePullJobOwnerReference {
  s.Controller = &v
  return s
}

func (s *PatchImagePullJobOwnerReference) SetBlockOwnerDeletion(v bool) *PatchImagePullJobOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type DeployVmpImagePreheatingRequest struct {
  // {"en":"Image ID","zh_CN":"镜像id"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Name of preheating node","zh_CN":"预热节点"}
  NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" require:"true" type:"Repeated"`
}

func (s DeployVmpImagePreheatingRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingRequest) GoString() string {
  return s.String()
}

func (s *DeployVmpImagePreheatingRequest) SetImageId(v string) *DeployVmpImagePreheatingRequest {
  s.ImageId = &v
  return s
}

func (s *DeployVmpImagePreheatingRequest) SetNodeNames(v []*string) *DeployVmpImagePreheatingRequest {
  s.NodeNames = v
  return s
}

type DeployVmpImagePreheatingRequestHeader struct {
}

func (s DeployVmpImagePreheatingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingRequestHeader) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingPaths struct {
}

func (s DeployVmpImagePreheatingPaths) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingPaths) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingParameters struct {
}

func (s DeployVmpImagePreheatingParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingParameters) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeployVmpImagePreheatingResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingResponse) GoString() string {
  return s.String()
}

func (s *DeployVmpImagePreheatingResponse) SetCode(v string) *DeployVmpImagePreheatingResponse {
  s.Code = &v
  return s
}

func (s *DeployVmpImagePreheatingResponse) SetMessage(v string) *DeployVmpImagePreheatingResponse {
  s.Message = &v
  return s
}

type DeployVmpImagePreheatingResponseHeader struct {
}

func (s DeployVmpImagePreheatingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryImageRequest struct {
}

func (s VMPQueryImageRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageRequest) GoString() string {
  return s.String()
}

type VMPQueryImageResponse struct {
  // {"en":"Image information array", "zh_CN":"镜像信息数组"}
  Images []*string `json:"images,omitempty" xml:"images,omitempty" require:"true" type:"Repeated"`
  // {"en":"Image unique ID, global unique", "zh_CN":"镜像唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Mirror name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Mirror belongs to the Lord.", "zh_CN":"镜像属主"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Image creation time, such as: 2017-11-04 14:17:41", "zh_CN":"镜像创建时间，如：2017-11-04 14:17:41"}
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty" require:"true"`
  // {"en":"Image size in GB", "zh_CN":"镜像大小，单位是GB"}
  Size *string `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Mirror Status", "zh_CN":"镜像状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s VMPQueryImageResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryImageResponse) SetImages(v []*string) *VMPQueryImageResponse {
  s.Images = v
  return s
}

func (s *VMPQueryImageResponse) SetId(v string) *VMPQueryImageResponse {
  s.Id = &v
  return s
}

func (s *VMPQueryImageResponse) SetName(v string) *VMPQueryImageResponse {
  s.Name = &v
  return s
}

func (s *VMPQueryImageResponse) SetType(v string) *VMPQueryImageResponse {
  s.Type = &v
  return s
}

func (s *VMPQueryImageResponse) SetCreatedAt(v string) *VMPQueryImageResponse {
  s.CreatedAt = &v
  return s
}

func (s *VMPQueryImageResponse) SetSize(v string) *VMPQueryImageResponse {
  s.Size = &v
  return s
}

func (s *VMPQueryImageResponse) SetState(v string) *VMPQueryImageResponse {
  s.State = &v
  return s
}

type VMPQueryImagePaths struct {
}

func (s VMPQueryImagePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImagePaths) GoString() string {
  return s.String()
}

type VMPQueryImageParameters struct {
  // {"en":"There can be multiple field names for sorting. The values are:
  // Name, createdat, type, size, state", "zh_CN":"排序的字段名称，可以有多个，取值：
  // name、createdAt、type、size、state"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"Sorting direction must follow sortkey. Value:
  // Desc: descending, default
  // ASC: ascending order", "zh_CN":"排序方向，必须跟在sortKey后面出现，取值：
  // desc：降序，默认值
  // asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the image ID specified by the marker", "zh_CN":"从marker指定的镜像id开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Mirror ID. A maximum of 100 IDS can be queried at a time. The IDs are separated by a half angle comma character ','.", "zh_CN":"镜像 ID。单次最多查询 100 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"Mirror name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The image belongs to the master. Is it a public image or a user-defined image? Value:
  // 
  // Common: official image
  // 
  // Snapshot: user snapshot image
  // 
  // Custom: user defined image'", "zh_CN":"镜像属主，是公共镜像还是用户自定义镜像，取值：
  // COMMON：官方镜像
  // SNAPSHOT：用户快照镜像
  // CUSTOM：用户自定义镜像"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Image status, value:
  // 
  // Active: available status
  // 
  // Building: Creating
  // 
  // Inactive: not available (such as creation failure, etc.)'", "zh_CN":"镜像状态，取值：
  // ACTIVE：可用状态
  // BUILDING：创建中
  // INACTIVE：不可用（如创建失败等）"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s VMPQueryImageParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryImageParameters) SetSortKey(v string) *VMPQueryImageParameters {
  s.SortKey = &v
  return s
}

func (s *VMPQueryImageParameters) SetSortDir(v string) *VMPQueryImageParameters {
  s.SortDir = &v
  return s
}

func (s *VMPQueryImageParameters) SetLimit(v int) *VMPQueryImageParameters {
  s.Limit = &v
  return s
}

func (s *VMPQueryImageParameters) SetMarker(v string) *VMPQueryImageParameters {
  s.Marker = &v
  return s
}

func (s *VMPQueryImageParameters) SetIds(v string) *VMPQueryImageParameters {
  s.Ids = &v
  return s
}

func (s *VMPQueryImageParameters) SetName(v string) *VMPQueryImageParameters {
  s.Name = &v
  return s
}

func (s *VMPQueryImageParameters) SetType(v string) *VMPQueryImageParameters {
  s.Type = &v
  return s
}

func (s *VMPQueryImageParameters) SetState(v string) *VMPQueryImageParameters {
  s.State = &v
  return s
}

type VMPQueryImageRequestHeader struct {
}

func (s VMPQueryImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryImageResponseHeader struct {
}

func (s VMPQueryImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageResponseHeader) GoString() string {
  return s.String()
}




type CreateHarborProjectRequest struct {
  // {"en":"project name", "zh_CN":"项目名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"false: private, true: public", "zh_CN":"false: 私有,true:公开"}
  Pub *bool `json:"pub,omitempty" xml:"pub,omitempty" require:"true"`
  // {"en":"project describe", "zh_CN":"项目描述"}
  Describe *string `json:"describe,omitempty" xml:"describe,omitempty"`
  // {"en":"authorized user", "zh_CN":"授权用户"}
  AssignRoles []*CreateHarborProjectAssignRole `json:"assignRoles,omitempty" xml:"assignRoles,omitempty" type:"Repeated"`
}

func (s CreateHarborProjectRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectRequest) GoString() string {
  return s.String()
}

func (s *CreateHarborProjectRequest) SetName(v string) *CreateHarborProjectRequest {
  s.Name = &v
  return s
}

func (s *CreateHarborProjectRequest) SetPub(v bool) *CreateHarborProjectRequest {
  s.Pub = &v
  return s
}

func (s *CreateHarborProjectRequest) SetDescribe(v string) *CreateHarborProjectRequest {
  s.Describe = &v
  return s
}

func (s *CreateHarborProjectRequest) SetAssignRoles(v []*CreateHarborProjectAssignRole) *CreateHarborProjectRequest {
  s.AssignRoles = v
  return s
}

type CreateHarborProjectResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"project id", "zh_CN":"项目id"}
  Data *int64 `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateHarborProjectResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectResponse) GoString() string {
  return s.String()
}

func (s *CreateHarborProjectResponse) SetCode(v int64) *CreateHarborProjectResponse {
  s.Code = &v
  return s
}

func (s *CreateHarborProjectResponse) SetMsg(v string) *CreateHarborProjectResponse {
  s.Msg = &v
  return s
}

func (s *CreateHarborProjectResponse) SetRequestId(v string) *CreateHarborProjectResponse {
  s.RequestId = &v
  return s
}

func (s *CreateHarborProjectResponse) SetData(v int64) *CreateHarborProjectResponse {
  s.Data = &v
  return s
}

type CreateHarborProjectPaths struct {
}

func (s CreateHarborProjectPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectPaths) GoString() string {
  return s.String()
}

type CreateHarborProjectParameters struct {
}

func (s CreateHarborProjectParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectParameters) GoString() string {
  return s.String()
}

type CreateHarborProjectRequestHeader struct {
}

func (s CreateHarborProjectRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectRequestHeader) GoString() string {
  return s.String()
}

type CreateHarborProjectResponseHeader struct {
}

func (s CreateHarborProjectResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectResponseHeader) GoString() string {
  return s.String()
}

type CreateHarborProjectAssignRole struct {
  // {"en":"user name", "zh_CN":"用户名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"1: admin, 2: edit, 3: access", "zh_CN":"1: 管理,2: 编辑, 3：访问"}
  Role *int32 `json:"role,omitempty" xml:"role,omitempty" require:"true"`
}

func (s CreateHarborProjectAssignRole) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborProjectAssignRole) GoString() string {
  return s.String()
}

func (s *CreateHarborProjectAssignRole) SetName(v string) *CreateHarborProjectAssignRole {
  s.Name = &v
  return s
}

func (s *CreateHarborProjectAssignRole) SetRole(v int32) *CreateHarborProjectAssignRole {
  s.Role = &v
  return s
}




type ManageOemImageRequest struct {
  // {"en":"list of apps to operate on", "zh_CN":"操作的云手机应用数组"}
  OemImages []*ManageOemImageOperateObject `json:"oemImages,omitempty" xml:"oemImages,omitempty" require:"true" type:"Repeated"`
}

func (s ManageOemImageRequest) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageRequest) GoString() string {
  return s.String()
}

func (s *ManageOemImageRequest) SetOemImages(v []*ManageOemImageOperateObject) *ManageOemImageRequest {
  s.OemImages = v
  return s
}

type ManageOemImageOperateObject struct {
  // {"en":"ephone oem image id", "zh_CN":"云手机oem镜像id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"operate type", "zh_CN":"操作类型，可选值：delete"}
  Op *string `json:"op,omitempty" xml:"op,omitempty" require:"true"`
  // {"en":"oem image operate params", "zh_CN":"操作的参数"}
  Params *ManageOemImageAppParamsObject `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ManageOemImageOperateObject) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageOperateObject) GoString() string {
  return s.String()
}

func (s *ManageOemImageOperateObject) SetId(v string) *ManageOemImageOperateObject {
  s.Id = &v
  return s
}

func (s *ManageOemImageOperateObject) SetOp(v string) *ManageOemImageOperateObject {
  s.Op = &v
  return s
}

func (s *ManageOemImageOperateObject) SetParams(v *ManageOemImageAppParamsObject) *ManageOemImageOperateObject {
  s.Params = v
  return s
}

type ManageOemImageAppParamsObject struct {
}

func (s ManageOemImageAppParamsObject) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageAppParamsObject) GoString() string {
  return s.String()
}

type ManageOemImageResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Tasks []*ManageOemImageTask `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s ManageOemImageResponse) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageResponse) GoString() string {
  return s.String()
}

func (s *ManageOemImageResponse) SetTasks(v []*ManageOemImageTask) *ManageOemImageResponse {
  s.Tasks = v
  return s
}

func (s *ManageOemImageResponse) SetStatus(v int) *ManageOemImageResponse {
  s.Status = &v
  return s
}

func (s *ManageOemImageResponse) SetResult(v string) *ManageOemImageResponse {
  s.Result = &v
  return s
}

type ManageOemImageTask struct {
  // {"en":"task id", "zh_CN":"任务单id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"任务单消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ManageOemImageTask) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageTask) GoString() string {
  return s.String()
}

func (s *ManageOemImageTask) SetId(v string) *ManageOemImageTask {
  s.Id = &v
  return s
}

func (s *ManageOemImageTask) SetMessage(v string) *ManageOemImageTask {
  s.Message = &v
  return s
}

type ManageOemImagePaths struct {
}

func (s ManageOemImagePaths) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImagePaths) GoString() string {
  return s.String()
}

type ManageOemImageParameters struct {
}

func (s ManageOemImageParameters) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageParameters) GoString() string {
  return s.String()
}

type ManageOemImageRequestHeader struct {
}

func (s ManageOemImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageRequestHeader) GoString() string {
  return s.String()
}

type ManageOemImageResponseHeader struct {
}

func (s ManageOemImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageOemImageResponseHeader) GoString() string {
  return s.String()
}




type CreateOemImageRequest struct {
  // {"en":"list of ephones", "zh_CN":"云手机请求列表"}
  OemImages []*CreateOemImageOemImageReq `json:"oemImages,omitempty" xml:"oemImages,omitempty" require:"true" type:"Repeated"`
}

func (s CreateOemImageRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageRequest) GoString() string {
  return s.String()
}

func (s *CreateOemImageRequest) SetOemImages(v []*CreateOemImageOemImageReq) *CreateOemImageRequest {
  s.OemImages = v
  return s
}

type CreateOemImageOemImageReq struct {
  // {"en":"ephone instance id", "zh_CN":"云手机实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"oem alias name", "zh_CN":"oem镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreateOemImageOemImageReq) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageOemImageReq) GoString() string {
  return s.String()
}

func (s *CreateOemImageOemImageReq) SetInstanceId(v string) *CreateOemImageOemImageReq {
  s.InstanceId = &v
  return s
}

func (s *CreateOemImageOemImageReq) SetName(v string) *CreateOemImageOemImageReq {
  s.Name = &v
  return s
}

type CreateOemImageResponse struct {
  // {"en":"oem image detail", "zh_CN":"Oem镜像详情"}
  OemImages []*CreateOemImageOemImageInfo `json:"oemImages,omitempty" xml:"oemImages,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s CreateOemImageResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageResponse) GoString() string {
  return s.String()
}

func (s *CreateOemImageResponse) SetOemImages(v []*CreateOemImageOemImageInfo) *CreateOemImageResponse {
  s.OemImages = v
  return s
}

func (s *CreateOemImageResponse) SetStatus(v int) *CreateOemImageResponse {
  s.Status = &v
  return s
}

func (s *CreateOemImageResponse) SetResult(v string) *CreateOemImageResponse {
  s.Result = &v
  return s
}

type CreateOemImageOemImageInfo struct {
  // {"en":"oem image id", "zh_CN":"oem 镜像ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"oem alias name", "zh_CN":"OEM名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"ephone instance id", "zh_CN":"云手机实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
}

func (s CreateOemImageOemImageInfo) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageOemImageInfo) GoString() string {
  return s.String()
}

func (s *CreateOemImageOemImageInfo) SetId(v string) *CreateOemImageOemImageInfo {
  s.Id = &v
  return s
}

func (s *CreateOemImageOemImageInfo) SetName(v string) *CreateOemImageOemImageInfo {
  s.Name = &v
  return s
}

func (s *CreateOemImageOemImageInfo) SetInstanceId(v string) *CreateOemImageOemImageInfo {
  s.InstanceId = &v
  return s
}

type CreateOemImagePaths struct {
}

func (s CreateOemImagePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImagePaths) GoString() string {
  return s.String()
}

type CreateOemImageParameters struct {
}

func (s CreateOemImageParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageParameters) GoString() string {
  return s.String()
}

type CreateOemImageRequestHeader struct {
}

func (s CreateOemImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageRequestHeader) GoString() string {
  return s.String()
}

type CreateOemImageResponseHeader struct {
}

func (s CreateOemImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateOemImageResponseHeader) GoString() string {
  return s.String()
}




type UpdateImagePullJobRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *UpdateImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the UpdateImagePullJobImagePullJob.", "zh_CN":"UpdateImagePullJobImagePullJob 预期行为的规约。"}
  Spec *UpdateImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the UpdateImagePullJobImagePullJob.", "zh_CN":"最近观测到的 UpdateImagePullJobImagePullJob 状态。"}
  Status *UpdateImagePullJobImagePullJobStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateImagePullJobRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobRequest) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobRequest) SetApiVersion(v string) *UpdateImagePullJobRequest {
  s.ApiVersion = &v
  return s
}

func (s *UpdateImagePullJobRequest) SetKind(v string) *UpdateImagePullJobRequest {
  s.Kind = &v
  return s
}

func (s *UpdateImagePullJobRequest) SetMetadata(v *UpdateImagePullJobObjectMeta) *UpdateImagePullJobRequest {
  s.Metadata = v
  return s
}

func (s *UpdateImagePullJobRequest) SetSpec(v *UpdateImagePullJobImagePullJobSpec) *UpdateImagePullJobRequest {
  s.Spec = v
  return s
}

func (s *UpdateImagePullJobRequest) SetStatus(v *UpdateImagePullJobImagePullJobStatus) *UpdateImagePullJobRequest {
  s.Status = v
  return s
}

type UpdateImagePullJobResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"imagepulljob", "zh_CN":"imagepulljob"}
  Data *UpdateImagePullJobImagePullJob `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateImagePullJobResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobResponse) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobResponse) SetCode(v int64) *UpdateImagePullJobResponse {
  s.Code = &v
  return s
}

func (s *UpdateImagePullJobResponse) SetMsg(v string) *UpdateImagePullJobResponse {
  s.Msg = &v
  return s
}

func (s *UpdateImagePullJobResponse) SetRequestId(v string) *UpdateImagePullJobResponse {
  s.RequestId = &v
  return s
}

func (s *UpdateImagePullJobResponse) SetData(v *UpdateImagePullJobImagePullJob) *UpdateImagePullJobResponse {
  s.Data = v
  return s
}

type UpdateImagePullJobPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"imagepulljob name", "zh_CN":"imagepulljob 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateImagePullJobPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobPaths) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobPaths) SetNamespace(v string) *UpdateImagePullJobPaths {
  s.Namespace = &v
  return s
}

func (s *UpdateImagePullJobPaths) SetName(v string) *UpdateImagePullJobPaths {
  s.Name = &v
  return s
}

type UpdateImagePullJobParameters struct {
}

func (s UpdateImagePullJobParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobParameters) GoString() string {
  return s.String()
}

type UpdateImagePullJobRequestHeader struct {
}

func (s UpdateImagePullJobRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobRequestHeader) GoString() string {
  return s.String()
}

type UpdateImagePullJobResponseHeader struct {
}

func (s UpdateImagePullJobResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobResponseHeader) GoString() string {
  return s.String()
}

type UpdateImagePullJobImagePullJob struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *UpdateImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the UpdateImagePullJobImagePullJob.", "zh_CN":"UpdateImagePullJobImagePullJob 预期行为的规约。"}
  Spec *UpdateImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the UpdateImagePullJobImagePullJob.", "zh_CN":"最近观测到的 UpdateImagePullJobImagePullJob 状态。"}
  Status *UpdateImagePullJobImagePullJobStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateImagePullJobImagePullJob) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobImagePullJob) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobImagePullJob) SetApiVersion(v string) *UpdateImagePullJobImagePullJob {
  s.ApiVersion = &v
  return s
}

func (s *UpdateImagePullJobImagePullJob) SetKind(v string) *UpdateImagePullJobImagePullJob {
  s.Kind = &v
  return s
}

func (s *UpdateImagePullJobImagePullJob) SetMetadata(v *UpdateImagePullJobObjectMeta) *UpdateImagePullJobImagePullJob {
  s.Metadata = v
  return s
}

func (s *UpdateImagePullJobImagePullJob) SetSpec(v *UpdateImagePullJobImagePullJobSpec) *UpdateImagePullJobImagePullJob {
  s.Spec = v
  return s
}

func (s *UpdateImagePullJobImagePullJob) SetStatus(v *UpdateImagePullJobImagePullJobStatus) *UpdateImagePullJobImagePullJob {
  s.Status = v
  return s
}

type UpdateImagePullJobImagePullJobSpec struct {
  // {"en": "Image is the image to be pulled by the job", "zh_CN": "拉取镜像名"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en": "Parallelism is the requested parallelism, it can be set to any non-negative value. If it is unspecified, it defaults to 1. If it is specified as 0, then the Job is effectively paused until it is increased.The value range 0-10, +optional", "zh_CN": "并发拉取个数, 范围0-10"}
  Parallelism *UpdateImagePullJobIntstrIntOrString `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
  // {"en": "UpdateImagePullJobPullPolicy is an optional field to set parameters of the pulling task. If not specified, the system will use the default values.+optional", "zh_CN": "拉取策略"}
  UpdateImagePullJobPullPolicy *UpdateImagePullJobPullPolicy `json:"pullPolicy,omitempty" xml:"pullPolicy,omitempty"`
  // {"en": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling the image.If specified, these secrets will be passed to individual puller implementations for them to use.  For example,in the case of docker, only DockerConfig type secrets are honored.+optional", "zh_CN": "拉取镜像所需的密钥"}
  PullSecrets []*string `json:"pullSecrets,omitempty" xml:"pullSecrets,omitempty" type:"Repeated"`
}

func (s UpdateImagePullJobImagePullJobSpec) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobImagePullJobSpec) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobImagePullJobSpec) SetImage(v string) *UpdateImagePullJobImagePullJobSpec {
  s.Image = &v
  return s
}

func (s *UpdateImagePullJobImagePullJobSpec) SetParallelism(v *UpdateImagePullJobIntstrIntOrString) *UpdateImagePullJobImagePullJobSpec {
  s.Parallelism = v
  return s
}

func (s *UpdateImagePullJobImagePullJobSpec) SetPullPolicy(v *UpdateImagePullJobPullPolicy) *UpdateImagePullJobImagePullJobSpec {
  s.UpdateImagePullJobPullPolicy = v
  return s
}

func (s *UpdateImagePullJobImagePullJobSpec) SetPullSecrets(v []*string) *UpdateImagePullJobImagePullJobSpec {
  s.PullSecrets = v
  return s
}

type UpdateImagePullJobIntstrIntOrString struct {
  // {"en": "the integer value", "zh_CN": "整数值"}
  IntVal *int `json:"intVal,omitempty" xml:"intVal,omitempty"`
  // {"en": "the string value", "zh_CN": "字符串值"}
  StrVal *string `json:"strVal,omitempty" xml:"strVal,omitempty"`
  // {"en": "type", "zh_CN": "类型"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateImagePullJobIntstrIntOrString) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobIntstrIntOrString) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobIntstrIntOrString) SetIntVal(v int) *UpdateImagePullJobIntstrIntOrString {
  s.IntVal = &v
  return s
}

func (s *UpdateImagePullJobIntstrIntOrString) SetStrVal(v string) *UpdateImagePullJobIntstrIntOrString {
  s.StrVal = &v
  return s
}

func (s *UpdateImagePullJobIntstrIntOrString) SetType(v int) *UpdateImagePullJobIntstrIntOrString {
  s.Type = &v
  return s
}

type UpdateImagePullJobPullPolicy struct {
  // {"en": "Specifies the number of retries before marking the pulling task failed. Defaults to 3 +optional", "zh_CN": "backoff次数，默认3"}
  BackoffLimit *int `json:"backoffLimit,omitempty" xml:"backoffLimit,omitempty"`
  // {"en": "Specifies the timeout of the pulling task. Defaults to 600 +optional", "zh_CN": "拉取超时时间"}
  TimeoutSeconds *int `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s UpdateImagePullJobPullPolicy) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobPullPolicy) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobPullPolicy) SetBackoffLimit(v int) *UpdateImagePullJobPullPolicy {
  s.BackoffLimit = &v
  return s
}

func (s *UpdateImagePullJobPullPolicy) SetTimeoutSeconds(v int) *UpdateImagePullJobPullPolicy {
  s.TimeoutSeconds = &v
  return s
}

type UpdateImagePullJobImagePullJobStatus struct {
  // {"en": "Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务完成时间"}
  CompletionTime *string `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
  // {"en": "The text prompt for job running status.+optional", "zh_CN": "状态消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // {"en": "Represents time when the job was acknowledged by the job controller.It is not guaranteed to be set in happens-before order across separate operations.It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en": "The rate of pulling tasks which reached phase Succeeded without not ready nodes. +optional", "zh_CN": "完成成功率"}
  SucceededRate *string `json:"succeededRate,omitempty" xml:"succeededRate,omitempty"`
}

func (s UpdateImagePullJobImagePullJobStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobImagePullJobStatus) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobImagePullJobStatus) SetCompletionTime(v string) *UpdateImagePullJobImagePullJobStatus {
  s.CompletionTime = &v
  return s
}

func (s *UpdateImagePullJobImagePullJobStatus) SetMessage(v string) *UpdateImagePullJobImagePullJobStatus {
  s.Message = &v
  return s
}

func (s *UpdateImagePullJobImagePullJobStatus) SetStartTime(v string) *UpdateImagePullJobImagePullJobStatus {
  s.StartTime = &v
  return s
}

func (s *UpdateImagePullJobImagePullJobStatus) SetSucceededRate(v string) *UpdateImagePullJobImagePullJobStatus {
  s.SucceededRate = &v
  return s
}

type UpdateImagePullJobObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*UpdateImagePullJobOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*UpdateImagePullJobManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s UpdateImagePullJobObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobObjectMeta) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobObjectMeta) SetName(v string) *UpdateImagePullJobObjectMeta {
  s.Name = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetGenerateName(v string) *UpdateImagePullJobObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetNamespace(v string) *UpdateImagePullJobObjectMeta {
  s.Namespace = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetSelfLink(v string) *UpdateImagePullJobObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetUid(v string) *UpdateImagePullJobObjectMeta {
  s.Uid = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetResourceVersion(v string) *UpdateImagePullJobObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetGeneration(v int64) *UpdateImagePullJobObjectMeta {
  s.Generation = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetCreationTimestamp(v string) *UpdateImagePullJobObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetDeletionTimestamp(v string) *UpdateImagePullJobObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetDeletionGracePeriodSeconds(v int64) *UpdateImagePullJobObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetLabels(v map[string]*string) *UpdateImagePullJobObjectMeta {
  s.Labels = v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetAnnotations(v map[string]*string) *UpdateImagePullJobObjectMeta {
  s.Annotations = v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetOwnerReferences(v []*UpdateImagePullJobOwnerReference) *UpdateImagePullJobObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetFinalizers(v []*string) *UpdateImagePullJobObjectMeta {
  s.Finalizers = v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetClusterName(v string) *UpdateImagePullJobObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *UpdateImagePullJobObjectMeta) SetManagedFields(v []*UpdateImagePullJobManagedFieldsEntry) *UpdateImagePullJobObjectMeta {
  s.ManagedFields = v
  return s
}

type UpdateImagePullJobManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this UpdateImagePullJobManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'UpdateImagePullJobFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“UpdateImagePullJobFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"UpdateImagePullJobFieldsV1 holds the first JSON version format as described in the 'UpdateImagePullJobFieldsV1' type", "zh_CN":"UpdateImagePullJobFieldsV1 包含类型 “UpdateImagePullJobFieldsV1” 中描述的第一个 JSON 版本格式"}
  UpdateImagePullJobFieldsV1 *UpdateImagePullJobFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s UpdateImagePullJobManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetManager(v string) *UpdateImagePullJobManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetOperation(v string) *UpdateImagePullJobManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetApiVersion(v string) *UpdateImagePullJobManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetTime(v string) *UpdateImagePullJobManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetFieldsType(v string) *UpdateImagePullJobManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetFieldsV1(v *UpdateImagePullJobFieldsV1) *UpdateImagePullJobManagedFieldsEntry {
  s.UpdateImagePullJobFieldsV1 = v
  return s
}

func (s *UpdateImagePullJobManagedFieldsEntry) SetSubresource(v string) *UpdateImagePullJobManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type UpdateImagePullJobFieldsV1 struct {
}

func (s UpdateImagePullJobFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobFieldsV1) GoString() string {
  return s.String()
}

type UpdateImagePullJobOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s UpdateImagePullJobOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s UpdateImagePullJobOwnerReference) GoString() string {
  return s.String()
}

func (s *UpdateImagePullJobOwnerReference) SetApiVersion(v string) *UpdateImagePullJobOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *UpdateImagePullJobOwnerReference) SetKind(v string) *UpdateImagePullJobOwnerReference {
  s.Kind = &v
  return s
}

func (s *UpdateImagePullJobOwnerReference) SetName(v string) *UpdateImagePullJobOwnerReference {
  s.Name = &v
  return s
}

func (s *UpdateImagePullJobOwnerReference) SetUid(v string) *UpdateImagePullJobOwnerReference {
  s.Uid = &v
  return s
}

func (s *UpdateImagePullJobOwnerReference) SetController(v bool) *UpdateImagePullJobOwnerReference {
  s.Controller = &v
  return s
}

func (s *UpdateImagePullJobOwnerReference) SetBlockOwnerDeletion(v bool) *UpdateImagePullJobOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type DeleteHarborUserRequest struct {
}

func (s DeleteHarborUserRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborUserRequest) GoString() string {
  return s.String()
}

type DeleteHarborUserResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"none", "zh_CN":"无"}
  Data *int64 `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteHarborUserResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborUserResponse) GoString() string {
  return s.String()
}

func (s *DeleteHarborUserResponse) SetCode(v int64) *DeleteHarborUserResponse {
  s.Code = &v
  return s
}

func (s *DeleteHarborUserResponse) SetMsg(v string) *DeleteHarborUserResponse {
  s.Msg = &v
  return s
}

func (s *DeleteHarborUserResponse) SetRequestId(v string) *DeleteHarborUserResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteHarborUserResponse) SetData(v int64) *DeleteHarborUserResponse {
  s.Data = &v
  return s
}

type DeleteHarborUserPaths struct {
  // {"en":"user name", "zh_CN":"用户名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteHarborUserPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborUserPaths) GoString() string {
  return s.String()
}

func (s *DeleteHarborUserPaths) SetName(v string) *DeleteHarborUserPaths {
  s.Name = &v
  return s
}

type DeleteHarborUserParameters struct {
}

func (s DeleteHarborUserParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborUserParameters) GoString() string {
  return s.String()
}

type DeleteHarborUserRequestHeader struct {
}

func (s DeleteHarborUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborUserRequestHeader) GoString() string {
  return s.String()
}

type DeleteHarborUserResponseHeader struct {
}

func (s DeleteHarborUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborUserResponseHeader) GoString() string {
  return s.String()
}




type GetImagePullJobRequest struct {
}

func (s GetImagePullJobRequest) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobRequest) GoString() string {
  return s.String()
}

type GetImagePullJobResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"imagepulljob", "zh_CN":"imagepulljob"}
  Data *GetImagePullJobImagePullJob `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetImagePullJobResponse) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobResponse) GoString() string {
  return s.String()
}

func (s *GetImagePullJobResponse) SetCode(v int64) *GetImagePullJobResponse {
  s.Code = &v
  return s
}

func (s *GetImagePullJobResponse) SetMsg(v string) *GetImagePullJobResponse {
  s.Msg = &v
  return s
}

func (s *GetImagePullJobResponse) SetRequestId(v string) *GetImagePullJobResponse {
  s.RequestId = &v
  return s
}

func (s *GetImagePullJobResponse) SetData(v *GetImagePullJobImagePullJob) *GetImagePullJobResponse {
  s.Data = v
  return s
}

type GetImagePullJobPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"imagepulljob name", "zh_CN":"imagepulljob 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetImagePullJobPaths) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobPaths) GoString() string {
  return s.String()
}

func (s *GetImagePullJobPaths) SetNamespace(v string) *GetImagePullJobPaths {
  s.Namespace = &v
  return s
}

func (s *GetImagePullJobPaths) SetName(v string) *GetImagePullJobPaths {
  s.Name = &v
  return s
}

type GetImagePullJobParameters struct {
}

func (s GetImagePullJobParameters) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobParameters) GoString() string {
  return s.String()
}

type GetImagePullJobRequestHeader struct {
}

func (s GetImagePullJobRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobRequestHeader) GoString() string {
  return s.String()
}

type GetImagePullJobResponseHeader struct {
}

func (s GetImagePullJobResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobResponseHeader) GoString() string {
  return s.String()
}

type GetImagePullJobImagePullJob struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *GetImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the GetImagePullJobImagePullJob.", "zh_CN":"GetImagePullJobImagePullJob 预期行为的规约。"}
  Spec *GetImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the GetImagePullJobImagePullJob.", "zh_CN":"最近观测到的 GetImagePullJobImagePullJob 状态。"}
  Status *GetImagePullJobImagePullJobStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetImagePullJobImagePullJob) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobImagePullJob) GoString() string {
  return s.String()
}

func (s *GetImagePullJobImagePullJob) SetApiVersion(v string) *GetImagePullJobImagePullJob {
  s.ApiVersion = &v
  return s
}

func (s *GetImagePullJobImagePullJob) SetKind(v string) *GetImagePullJobImagePullJob {
  s.Kind = &v
  return s
}

func (s *GetImagePullJobImagePullJob) SetMetadata(v *GetImagePullJobObjectMeta) *GetImagePullJobImagePullJob {
  s.Metadata = v
  return s
}

func (s *GetImagePullJobImagePullJob) SetSpec(v *GetImagePullJobImagePullJobSpec) *GetImagePullJobImagePullJob {
  s.Spec = v
  return s
}

func (s *GetImagePullJobImagePullJob) SetStatus(v *GetImagePullJobImagePullJobStatus) *GetImagePullJobImagePullJob {
  s.Status = v
  return s
}

type GetImagePullJobImagePullJobSpec struct {
  // {"en": "Image is the image to be pulled by the job", "zh_CN": "拉取镜像名"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en": "Parallelism is the requested parallelism, it can be set to any non-negative value. If it is unspecified, it defaults to 1. If it is specified as 0, then the Job is effectively paused until it is increased.The value range 0-10, +optional", "zh_CN": "并发拉取个数, 范围0-10"}
  Parallelism *GetImagePullJobIntstrIntOrString `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
  // {"en": "GetImagePullJobPullPolicy is an optional field to set parameters of the pulling task. If not specified, the system will use the default values.+optional", "zh_CN": "拉取策略"}
  GetImagePullJobPullPolicy *GetImagePullJobPullPolicy `json:"pullPolicy,omitempty" xml:"pullPolicy,omitempty"`
  // {"en": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling the image.If specified, these secrets will be passed to individual puller implementations for them to use.  For example,in the case of docker, only DockerConfig type secrets are honored.+optional", "zh_CN": "拉取镜像所需的密钥"}
  PullSecrets []*string `json:"pullSecrets,omitempty" xml:"pullSecrets,omitempty" type:"Repeated"`
}

func (s GetImagePullJobImagePullJobSpec) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobImagePullJobSpec) GoString() string {
  return s.String()
}

func (s *GetImagePullJobImagePullJobSpec) SetImage(v string) *GetImagePullJobImagePullJobSpec {
  s.Image = &v
  return s
}

func (s *GetImagePullJobImagePullJobSpec) SetParallelism(v *GetImagePullJobIntstrIntOrString) *GetImagePullJobImagePullJobSpec {
  s.Parallelism = v
  return s
}

func (s *GetImagePullJobImagePullJobSpec) SetPullPolicy(v *GetImagePullJobPullPolicy) *GetImagePullJobImagePullJobSpec {
  s.GetImagePullJobPullPolicy = v
  return s
}

func (s *GetImagePullJobImagePullJobSpec) SetPullSecrets(v []*string) *GetImagePullJobImagePullJobSpec {
  s.PullSecrets = v
  return s
}

type GetImagePullJobIntstrIntOrString struct {
  // {"en": "the integer value", "zh_CN": "整数值"}
  IntVal *int `json:"intVal,omitempty" xml:"intVal,omitempty"`
  // {"en": "the string value", "zh_CN": "字符串值"}
  StrVal *string `json:"strVal,omitempty" xml:"strVal,omitempty"`
  // {"en": "type", "zh_CN": "类型"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetImagePullJobIntstrIntOrString) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobIntstrIntOrString) GoString() string {
  return s.String()
}

func (s *GetImagePullJobIntstrIntOrString) SetIntVal(v int) *GetImagePullJobIntstrIntOrString {
  s.IntVal = &v
  return s
}

func (s *GetImagePullJobIntstrIntOrString) SetStrVal(v string) *GetImagePullJobIntstrIntOrString {
  s.StrVal = &v
  return s
}

func (s *GetImagePullJobIntstrIntOrString) SetType(v int) *GetImagePullJobIntstrIntOrString {
  s.Type = &v
  return s
}

type GetImagePullJobPullPolicy struct {
  // {"en": "Specifies the number of retries before marking the pulling task failed. Defaults to 3 +optional", "zh_CN": "backoff次数，默认3"}
  BackoffLimit *int `json:"backoffLimit,omitempty" xml:"backoffLimit,omitempty"`
  // {"en": "Specifies the timeout of the pulling task. Defaults to 600 +optional", "zh_CN": "拉取超时时间"}
  TimeoutSeconds *int `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s GetImagePullJobPullPolicy) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobPullPolicy) GoString() string {
  return s.String()
}

func (s *GetImagePullJobPullPolicy) SetBackoffLimit(v int) *GetImagePullJobPullPolicy {
  s.BackoffLimit = &v
  return s
}

func (s *GetImagePullJobPullPolicy) SetTimeoutSeconds(v int) *GetImagePullJobPullPolicy {
  s.TimeoutSeconds = &v
  return s
}

type GetImagePullJobImagePullJobStatus struct {
  // {"en": "Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务完成时间"}
  CompletionTime *string `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
  // {"en": "The text prompt for job running status.+optional", "zh_CN": "状态消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // {"en": "Represents time when the job was acknowledged by the job controller.It is not guaranteed to be set in happens-before order across separate operations.It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en": "The rate of pulling tasks which reached phase Succeeded without not ready nodes. +optional", "zh_CN": "完成成功率"}
  SucceededRate *string `json:"succeededRate,omitempty" xml:"succeededRate,omitempty"`
}

func (s GetImagePullJobImagePullJobStatus) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobImagePullJobStatus) GoString() string {
  return s.String()
}

func (s *GetImagePullJobImagePullJobStatus) SetCompletionTime(v string) *GetImagePullJobImagePullJobStatus {
  s.CompletionTime = &v
  return s
}

func (s *GetImagePullJobImagePullJobStatus) SetMessage(v string) *GetImagePullJobImagePullJobStatus {
  s.Message = &v
  return s
}

func (s *GetImagePullJobImagePullJobStatus) SetStartTime(v string) *GetImagePullJobImagePullJobStatus {
  s.StartTime = &v
  return s
}

func (s *GetImagePullJobImagePullJobStatus) SetSucceededRate(v string) *GetImagePullJobImagePullJobStatus {
  s.SucceededRate = &v
  return s
}

type GetImagePullJobObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*GetImagePullJobOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*GetImagePullJobManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s GetImagePullJobObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobObjectMeta) GoString() string {
  return s.String()
}

func (s *GetImagePullJobObjectMeta) SetName(v string) *GetImagePullJobObjectMeta {
  s.Name = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetGenerateName(v string) *GetImagePullJobObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetNamespace(v string) *GetImagePullJobObjectMeta {
  s.Namespace = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetSelfLink(v string) *GetImagePullJobObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetUid(v string) *GetImagePullJobObjectMeta {
  s.Uid = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetResourceVersion(v string) *GetImagePullJobObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetGeneration(v int64) *GetImagePullJobObjectMeta {
  s.Generation = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetCreationTimestamp(v string) *GetImagePullJobObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetDeletionTimestamp(v string) *GetImagePullJobObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetDeletionGracePeriodSeconds(v int64) *GetImagePullJobObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetLabels(v map[string]*string) *GetImagePullJobObjectMeta {
  s.Labels = v
  return s
}

func (s *GetImagePullJobObjectMeta) SetAnnotations(v map[string]*string) *GetImagePullJobObjectMeta {
  s.Annotations = v
  return s
}

func (s *GetImagePullJobObjectMeta) SetOwnerReferences(v []*GetImagePullJobOwnerReference) *GetImagePullJobObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *GetImagePullJobObjectMeta) SetFinalizers(v []*string) *GetImagePullJobObjectMeta {
  s.Finalizers = v
  return s
}

func (s *GetImagePullJobObjectMeta) SetClusterName(v string) *GetImagePullJobObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *GetImagePullJobObjectMeta) SetManagedFields(v []*GetImagePullJobManagedFieldsEntry) *GetImagePullJobObjectMeta {
  s.ManagedFields = v
  return s
}

type GetImagePullJobManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this GetImagePullJobManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'GetImagePullJobFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“GetImagePullJobFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"GetImagePullJobFieldsV1 holds the first JSON version format as described in the 'GetImagePullJobFieldsV1' type", "zh_CN":"GetImagePullJobFieldsV1 包含类型 “GetImagePullJobFieldsV1” 中描述的第一个 JSON 版本格式"}
  GetImagePullJobFieldsV1 *GetImagePullJobFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s GetImagePullJobManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *GetImagePullJobManagedFieldsEntry) SetManager(v string) *GetImagePullJobManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *GetImagePullJobManagedFieldsEntry) SetOperation(v string) *GetImagePullJobManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *GetImagePullJobManagedFieldsEntry) SetApiVersion(v string) *GetImagePullJobManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *GetImagePullJobManagedFieldsEntry) SetTime(v string) *GetImagePullJobManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *GetImagePullJobManagedFieldsEntry) SetFieldsType(v string) *GetImagePullJobManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *GetImagePullJobManagedFieldsEntry) SetFieldsV1(v *GetImagePullJobFieldsV1) *GetImagePullJobManagedFieldsEntry {
  s.GetImagePullJobFieldsV1 = v
  return s
}

func (s *GetImagePullJobManagedFieldsEntry) SetSubresource(v string) *GetImagePullJobManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type GetImagePullJobFieldsV1 struct {
}

func (s GetImagePullJobFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobFieldsV1) GoString() string {
  return s.String()
}

type GetImagePullJobOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s GetImagePullJobOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s GetImagePullJobOwnerReference) GoString() string {
  return s.String()
}

func (s *GetImagePullJobOwnerReference) SetApiVersion(v string) *GetImagePullJobOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *GetImagePullJobOwnerReference) SetKind(v string) *GetImagePullJobOwnerReference {
  s.Kind = &v
  return s
}

func (s *GetImagePullJobOwnerReference) SetName(v string) *GetImagePullJobOwnerReference {
  s.Name = &v
  return s
}

func (s *GetImagePullJobOwnerReference) SetUid(v string) *GetImagePullJobOwnerReference {
  s.Uid = &v
  return s
}

func (s *GetImagePullJobOwnerReference) SetController(v bool) *GetImagePullJobOwnerReference {
  s.Controller = &v
  return s
}

func (s *GetImagePullJobOwnerReference) SetBlockOwnerDeletion(v bool) *GetImagePullJobOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type QueryVmpImagePreheatingStateRequest struct {
}

func (s QueryVmpImagePreheatingStateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateRequest) GoString() string {
  return s.String()
}

type QueryVmpImagePreheatingStateRequestHeader struct {
}

func (s QueryVmpImagePreheatingStateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateRequestHeader) GoString() string {
  return s.String()
}

type QueryVmpImagePreheatingStatePaths struct {
  // {"en":"no","zh_CN":"镜像id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryVmpImagePreheatingStatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStatePaths) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStatePaths) SetId(v string) *QueryVmpImagePreheatingStatePaths {
  s.Id = &v
  return s
}

type QueryVmpImagePreheatingStateParameters struct {
}

func (s QueryVmpImagePreheatingStateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateParameters) GoString() string {
  return s.String()
}

type QueryVmpImagePreheatingStateResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *QueryVmpImagePreheatingStateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryVmpImagePreheatingStateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateResponse) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStateResponse) SetCode(v string) *QueryVmpImagePreheatingStateResponse {
  s.Code = &v
  return s
}

func (s *QueryVmpImagePreheatingStateResponse) SetMessage(v string) *QueryVmpImagePreheatingStateResponse {
  s.Message = &v
  return s
}

func (s *QueryVmpImagePreheatingStateResponse) SetData(v *QueryVmpImagePreheatingStateResponseData) *QueryVmpImagePreheatingStateResponse {
  s.Data = v
  return s
}

type QueryVmpImagePreheatingStateResponseData struct {
  // {"en":"Preheating info","zh_CN":"预热信息"}
  PreHeatingInfo []*QueryVmpImagePreheatingStateResponseDataPreHeatingInfo `json:"preHeatingInfo,omitempty" xml:"preHeatingInfo,omitempty" require:"true" type:"Repeated"`
}

func (s QueryVmpImagePreheatingStateResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateResponseData) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStateResponseData) SetPreHeatingInfo(v []*QueryVmpImagePreheatingStateResponseDataPreHeatingInfo) *QueryVmpImagePreheatingStateResponseData {
  s.PreHeatingInfo = v
  return s
}

type QueryVmpImagePreheatingStateResponseDataPreHeatingInfo struct     {
  // {"en":"Node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"State","zh_CN":"状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s QueryVmpImagePreheatingStateResponseDataPreHeatingInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateResponseDataPreHeatingInfo) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStateResponseDataPreHeatingInfo) SetNodeName(v string) *QueryVmpImagePreheatingStateResponseDataPreHeatingInfo {
  s.NodeName = &v
  return s
}

func (s *QueryVmpImagePreheatingStateResponseDataPreHeatingInfo) SetState(v string) *QueryVmpImagePreheatingStateResponseDataPreHeatingInfo {
  s.State = &v
  return s
}

type QueryVmpImagePreheatingStateResponseHeader struct {
}

func (s QueryVmpImagePreheatingStateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateResponseHeader) GoString() string {
  return s.String()
}




type LECHDeployImagePreheatingRequest struct {
  // {"en":"Image ID","zh_CN":"镜像id"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Name of preheating node","zh_CN":"预热节点"}
  NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" require:"true" type:"Repeated"`
}

func (s LECHDeployImagePreheatingRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHDeployImagePreheatingRequest) GoString() string {
  return s.String()
}

func (s *LECHDeployImagePreheatingRequest) SetImageId(v string) *LECHDeployImagePreheatingRequest {
  s.ImageId = &v
  return s
}

func (s *LECHDeployImagePreheatingRequest) SetNodeNames(v []*string) *LECHDeployImagePreheatingRequest {
  s.NodeNames = v
  return s
}

type LECHDeployImagePreheatingRequestHeader struct {
}

func (s LECHDeployImagePreheatingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHDeployImagePreheatingRequestHeader) GoString() string {
  return s.String()
}

type LECHDeployImagePreheatingPaths struct {
}

func (s LECHDeployImagePreheatingPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHDeployImagePreheatingPaths) GoString() string {
  return s.String()
}

type LECHDeployImagePreheatingParameters struct {
}

func (s LECHDeployImagePreheatingParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHDeployImagePreheatingParameters) GoString() string {
  return s.String()
}

type LECHDeployImagePreheatingResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHDeployImagePreheatingResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHDeployImagePreheatingResponse) GoString() string {
  return s.String()
}

func (s *LECHDeployImagePreheatingResponse) SetCode(v string) *LECHDeployImagePreheatingResponse {
  s.Code = &v
  return s
}

func (s *LECHDeployImagePreheatingResponse) SetMessage(v string) *LECHDeployImagePreheatingResponse {
  s.Message = &v
  return s
}

type LECHDeployImagePreheatingResponseHeader struct {
}

func (s LECHDeployImagePreheatingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHDeployImagePreheatingResponseHeader) GoString() string {
  return s.String()
}




type VMPCreateImageRequest struct {
  // {"en":"Mirror name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Virtual machine instance ID, optional parameter. This parameter and imagesrcurl must be one of two choices", "zh_CN":"虚拟机实例标识，可选参数，该参数与imageSrcUrl必须二选一"}
  InstanceUuid *string `json:"instanceUuid,omitempty" xml:"instanceUuid,omitempty"`
  // {"en":"Virtual machine image URL address, optional parameter. The parameter and instanceuuid must be one of two choices.", "zh_CN":"虚拟机镜像Url地址，可选参数，该参数与instanceUuid必须二选一"}
  ImageSrcUrl *string `json:"imageSrcUrl,omitempty" xml:"imageSrcUrl,omitempty"`
  // {"en":"MD5 value of virtual machine image, used with imagesrcurl", "zh_CN":"虚拟机镜像的md5值，与imageSrcUrl配合使用"}
  Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
  // {"en":"Operating system type, if the URL address is carried, this parameter is required; if the virtual machine is specified to be created, the ostype of the virtual machine shall prevail.
  // 
  // There are two values: windows and Linux'", "zh_CN":"操作系统类型，如果携带的是url地址时，该参数必填；如果是指定虚拟机创建，则以虚拟机的ostype为准。
  // 有2种取值：windows、linux"}
  Ostype *string `json:"ostype,omitempty" xml:"ostype,omitempty"`
  // {"en":"Minimum requirements for system disk, unit: GB. The system disk size of the selected template must be greater than or equal to this value, otherwise virtual machine creation fails", "zh_CN":"系统盘最小要求，单位是GB。选择的模板的系统盘大小必须大于等于该值，否则虚拟机创建失败"}
  MinDisk *int `json:"minDisk,omitempty" xml:"minDisk,omitempty"`
  // {"en":"Whether QEMU guest agent is enabled for the image, and password reset is supported for the opened image
  // 
  // There are two values: true and false'", "zh_CN":"镜像是否开启了qemu guest agent，有开启的镜像支持密码重置
  // 有2种取值：TRUE、FALSE"}
  QgaEnabled *string `json:"qgaEnabled,omitempty" xml:"qgaEnabled,omitempty"`
}

func (s VMPCreateImageRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageRequest) GoString() string {
  return s.String()
}

func (s *VMPCreateImageRequest) SetName(v string) *VMPCreateImageRequest {
  s.Name = &v
  return s
}

func (s *VMPCreateImageRequest) SetInstanceUuid(v string) *VMPCreateImageRequest {
  s.InstanceUuid = &v
  return s
}

func (s *VMPCreateImageRequest) SetImageSrcUrl(v string) *VMPCreateImageRequest {
  s.ImageSrcUrl = &v
  return s
}

func (s *VMPCreateImageRequest) SetMd5(v string) *VMPCreateImageRequest {
  s.Md5 = &v
  return s
}

func (s *VMPCreateImageRequest) SetOstype(v string) *VMPCreateImageRequest {
  s.Ostype = &v
  return s
}

func (s *VMPCreateImageRequest) SetMinDisk(v int) *VMPCreateImageRequest {
  s.MinDisk = &v
  return s
}

func (s *VMPCreateImageRequest) SetQgaEnabled(v string) *VMPCreateImageRequest {
  s.QgaEnabled = &v
  return s
}

type VMPCreateImageResponse struct {
  // {"en":"Image unique ID, global unique", "zh_CN":"镜像唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s VMPCreateImageResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageResponse) GoString() string {
  return s.String()
}

func (s *VMPCreateImageResponse) SetId(v string) *VMPCreateImageResponse {
  s.Id = &v
  return s
}

type VMPCreateImagePaths struct {
}

func (s VMPCreateImagePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImagePaths) GoString() string {
  return s.String()
}

type VMPCreateImageParameters struct {
}

func (s VMPCreateImageParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageParameters) GoString() string {
  return s.String()
}

type VMPCreateImageRequestHeader struct {
}

func (s VMPCreateImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageRequestHeader) GoString() string {
  return s.String()
}

type VMPCreateImageResponseHeader struct {
}

func (s VMPCreateImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageResponseHeader) GoString() string {
  return s.String()
}




type ListHarborUserRequest struct {
}

func (s ListHarborUserRequest) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserRequest) GoString() string {
  return s.String()
}

type ListHarborUserResponse struct {
  // {"zh_CN":"请求返回码","en":"Request return code"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"zh_CN":"请求返回信息","en":"Request return information"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"zh_CN":"请求识别码","en":"Request ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"zh_CN":"用户列表","en":"user list"}
  Data *ListHarborUserUserList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListHarborUserResponse) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserResponse) GoString() string {
  return s.String()
}

func (s *ListHarborUserResponse) SetCode(v int64) *ListHarborUserResponse {
  s.Code = &v
  return s
}

func (s *ListHarborUserResponse) SetMsg(v string) *ListHarborUserResponse {
  s.Msg = &v
  return s
}

func (s *ListHarborUserResponse) SetRequestId(v string) *ListHarborUserResponse {
  s.RequestId = &v
  return s
}

func (s *ListHarborUserResponse) SetData(v *ListHarborUserUserList) *ListHarborUserResponse {
  s.Data = v
  return s
}

type ListHarborUserPaths struct {
}

func (s ListHarborUserPaths) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserPaths) GoString() string {
  return s.String()
}

type ListHarborUserParameters struct {
  // {"zh_CN":"搜索关键字","en":"search for the keyword"}
  Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
  // {"zh_CN":"分页页数: 默认0","en":"Number of pagination pages: Default 0"}
  PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"zh_CN":"每页记录数: 默认10","en":"Number of records per page: Default 10"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListHarborUserParameters) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserParameters) GoString() string {
  return s.String()
}

func (s *ListHarborUserParameters) SetKeyword(v string) *ListHarborUserParameters {
  s.Keyword = &v
  return s
}

func (s *ListHarborUserParameters) SetPageIndex(v string) *ListHarborUserParameters {
  s.PageIndex = &v
  return s
}

func (s *ListHarborUserParameters) SetPageSize(v string) *ListHarborUserParameters {
  s.PageSize = &v
  return s
}

type ListHarborUserRequestHeader struct {
}

func (s ListHarborUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserRequestHeader) GoString() string {
  return s.String()
}

type ListHarborUserResponseHeader struct {
}

func (s ListHarborUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserResponseHeader) GoString() string {
  return s.String()
}

type ListHarborUserUserList struct {
  // {"zh_CN":"用户列表","en":"user list"}
  Items []*ListHarborUserUser `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"记录总数","en":"total number of records"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s ListHarborUserUserList) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserUserList) GoString() string {
  return s.String()
}

func (s *ListHarborUserUserList) SetItems(v []*ListHarborUserUser) *ListHarborUserUserList {
  s.Items = v
  return s
}

func (s *ListHarborUserUserList) SetTotal(v int64) *ListHarborUserUserList {
  s.Total = &v
  return s
}

type ListHarborUserUser struct {
  // {"zh_CN":"用户id","en":"user id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"zh_CN":"用户名","en":"username"}
  Username *string `json:"username,omitempty" xml:"username,omitempty" require:"true"`
  // {"zh_CN":"创建时间","en":"creation time"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"zh_CN":"更新时间","en":"Update time"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"zh_CN":"用户所属项目","en":"ListHarborUserUser's project"}
  Projects []*string `json:"projects,omitempty" xml:"projects,omitempty" require:"true" type:"Repeated"`
}

func (s ListHarborUserUser) String() string {
  return tea.Prettify(s)
}

func (s ListHarborUserUser) GoString() string {
  return s.String()
}

func (s *ListHarborUserUser) SetId(v int64) *ListHarborUserUser {
  s.Id = &v
  return s
}

func (s *ListHarborUserUser) SetUsername(v string) *ListHarborUserUser {
  s.Username = &v
  return s
}

func (s *ListHarborUserUser) SetCreateTime(v int64) *ListHarborUserUser {
  s.CreateTime = &v
  return s
}

func (s *ListHarborUserUser) SetUpdateTime(v int64) *ListHarborUserUser {
  s.UpdateTime = &v
  return s
}

func (s *ListHarborUserUser) SetProjects(v []*string) *ListHarborUserUser {
  s.Projects = v
  return s
}




type QueryOemImageRequest struct {
}

func (s QueryOemImageRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImageRequest) GoString() string {
  return s.String()
}

type QueryOemImageOemImageInfo struct {
  // {"en":"oem image id", "zh_CN":"oem 镜像ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"oem alias name", "zh_CN":"OEM名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"ephone instance id", "zh_CN":"云手机实例id"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"oem image status", "zh_CN":"OEM镜像可用状态。1-Available，代表可用，2-Unavailable，代表不可用"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"创建时间"}
  CreateAt *string `json:"createAt,omitempty" xml:"createAt,omitempty" require:"true"`
  // {"en":"available time", "zh_CN":"可用时间"}
  AvailableAt *string `json:"availableAt,omitempty" xml:"availableAt,omitempty" require:"true"`
}

func (s QueryOemImageOemImageInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImageOemImageInfo) GoString() string {
  return s.String()
}

func (s *QueryOemImageOemImageInfo) SetId(v string) *QueryOemImageOemImageInfo {
  s.Id = &v
  return s
}

func (s *QueryOemImageOemImageInfo) SetName(v string) *QueryOemImageOemImageInfo {
  s.Name = &v
  return s
}

func (s *QueryOemImageOemImageInfo) SetInstanceId(v string) *QueryOemImageOemImageInfo {
  s.InstanceId = &v
  return s
}

func (s *QueryOemImageOemImageInfo) SetStatus(v string) *QueryOemImageOemImageInfo {
  s.Status = &v
  return s
}

func (s *QueryOemImageOemImageInfo) SetCreateAt(v string) *QueryOemImageOemImageInfo {
  s.CreateAt = &v
  return s
}

func (s *QueryOemImageOemImageInfo) SetAvailableAt(v string) *QueryOemImageOemImageInfo {
  s.AvailableAt = &v
  return s
}

type QueryOemImageResponse struct {
  // {"en":"list of ome image", "zh_CN":"云手机oem镜像列表"}
  OemImages []*QueryOemImageOemImageInfo `json:"oemImages,omitempty" xml:"oemImages,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s QueryOemImageResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImageResponse) GoString() string {
  return s.String()
}

func (s *QueryOemImageResponse) SetOemImages(v []*QueryOemImageOemImageInfo) *QueryOemImageResponse {
  s.OemImages = v
  return s
}

func (s *QueryOemImageResponse) SetStatus(v int) *QueryOemImageResponse {
  s.Status = &v
  return s
}

func (s *QueryOemImageResponse) SetResult(v string) *QueryOemImageResponse {
  s.Result = &v
  return s
}

type QueryOemImagePaths struct {
}

func (s QueryOemImagePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImagePaths) GoString() string {
  return s.String()
}

type QueryOemImageParameters struct {
  // {"en":"oem image id to be queried", "zh_CN":"要查询的oem镜像id"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"oem image name bo be queried", "zh_CN":"要查询的oem镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryOemImageParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImageParameters) GoString() string {
  return s.String()
}

func (s *QueryOemImageParameters) SetIds(v string) *QueryOemImageParameters {
  s.Ids = &v
  return s
}

func (s *QueryOemImageParameters) SetName(v string) *QueryOemImageParameters {
  s.Name = &v
  return s
}

type QueryOemImageRequestHeader struct {
}

func (s QueryOemImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImageRequestHeader) GoString() string {
  return s.String()
}

type QueryOemImageResponseHeader struct {
}

func (s QueryOemImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOemImageResponseHeader) GoString() string {
  return s.String()
}




type DeleteImagePullJobRequest struct {
}

func (s DeleteImagePullJobRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobRequest) GoString() string {
  return s.String()
}

type DeleteImagePullJobResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"status"}
  Data *DeleteImagePullJobStatus `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteImagePullJobResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobResponse) GoString() string {
  return s.String()
}

func (s *DeleteImagePullJobResponse) SetCode(v int64) *DeleteImagePullJobResponse {
  s.Code = &v
  return s
}

func (s *DeleteImagePullJobResponse) SetMsg(v string) *DeleteImagePullJobResponse {
  s.Msg = &v
  return s
}

func (s *DeleteImagePullJobResponse) SetRequestId(v string) *DeleteImagePullJobResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteImagePullJobResponse) SetData(v *DeleteImagePullJobStatus) *DeleteImagePullJobResponse {
  s.Data = v
  return s
}

type DeleteImagePullJobPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"imagepulljob name", "zh_CN":"imagepulljob 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteImagePullJobPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobPaths) GoString() string {
  return s.String()
}

func (s *DeleteImagePullJobPaths) SetNamespace(v string) *DeleteImagePullJobPaths {
  s.Namespace = &v
  return s
}

func (s *DeleteImagePullJobPaths) SetName(v string) *DeleteImagePullJobPaths {
  s.Name = &v
  return s
}

type DeleteImagePullJobParameters struct {
}

func (s DeleteImagePullJobParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobParameters) GoString() string {
  return s.String()
}

type DeleteImagePullJobRequestHeader struct {
}

func (s DeleteImagePullJobRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobRequestHeader) GoString() string {
  return s.String()
}

type DeleteImagePullJobResponseHeader struct {
}

func (s DeleteImagePullJobResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobResponseHeader) GoString() string {
  return s.String()
}

type DeleteImagePullJobStatus struct {
  // {"en":"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values", "zh_CN":"APIVersion 定义对象表示的版本化模式。 服务器应将已识别的模式转换为最新的内部值，并可能拒绝无法识别的值"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase", "zh_CN":"Kind 是一个字符串值，表示此对象表示的 REST 资源。 服务器可以从客户端提交请求的端点推断出这一点。 无法更新。驼峰式规则"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"DeleteImagePullJobStatus of the operation. One of: 'Success' or 'Failure'", "zh_CN":"操作状态。“Success”或“Failure” 之一"}
  DeleteImagePullJobStatus *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Suggested HTTP return code for this status, 0 if not set", "zh_CN":"此状态的建议 HTTP 返回代码，如果未设置，则为 0"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Extended data associated with the reason. Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type", "zh_CN":"与原因（Reason）相关的扩展数据。每个原因都可以定义自己的扩展细节。 此字段是可选的，并且不保证返回的数据符合任何模式，除非由原因类型定义"}
  Details *DeleteImagePullJobStatusDetails `json:"details,omitempty" xml:"details,omitempty" require:"true"`
}

func (s DeleteImagePullJobStatus) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobStatus) GoString() string {
  return s.String()
}

func (s *DeleteImagePullJobStatus) SetApiVersion(v string) *DeleteImagePullJobStatus {
  s.ApiVersion = &v
  return s
}

func (s *DeleteImagePullJobStatus) SetKind(v string) *DeleteImagePullJobStatus {
  s.Kind = &v
  return s
}

func (s *DeleteImagePullJobStatus) SetStatus(v string) *DeleteImagePullJobStatus {
  s.DeleteImagePullJobStatus = &v
  return s
}

func (s *DeleteImagePullJobStatus) SetCode(v int32) *DeleteImagePullJobStatus {
  s.Code = &v
  return s
}

func (s *DeleteImagePullJobStatus) SetDetails(v *DeleteImagePullJobStatusDetails) *DeleteImagePullJobStatus {
  s.Details = v
  return s
}

type DeleteImagePullJobStatusDetails struct {
  // {"en":"The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)", "zh_CN":"与状态 StatusReason 关联的资源的名称属性（当有一个可以描述的名称时）"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind", "zh_CN":"与状态 StatusReason 关联的资源的种类属性"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"The group attribute of the resource associated with the status StatusReason", "zh_CN":"与状态 StatusReason 关联的资源的组属性"}
  Group *string `json:"group,omitempty" xml:"group,omitempty" require:"true"`
  // {"en":"UID of the resource. (when there is a single resource which can be described)", "zh_CN":"资源的 UID（当有单个可以描述的资源时）"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty" require:"true"`
}

func (s DeleteImagePullJobStatusDetails) String() string {
  return tea.Prettify(s)
}

func (s DeleteImagePullJobStatusDetails) GoString() string {
  return s.String()
}

func (s *DeleteImagePullJobStatusDetails) SetName(v string) *DeleteImagePullJobStatusDetails {
  s.Name = &v
  return s
}

func (s *DeleteImagePullJobStatusDetails) SetKind(v string) *DeleteImagePullJobStatusDetails {
  s.Kind = &v
  return s
}

func (s *DeleteImagePullJobStatusDetails) SetGroup(v string) *DeleteImagePullJobStatusDetails {
  s.Group = &v
  return s
}

func (s *DeleteImagePullJobStatusDetails) SetUid(v string) *DeleteImagePullJobStatusDetails {
  s.Uid = &v
  return s
}




type GetHarborProjectRequest struct {
}

func (s GetHarborProjectRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectRequest) GoString() string {
  return s.String()
}

type GetHarborProjectResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"project detal", "zh_CN":"项目详情"}
  Data *GetHarborProjectProjectDetail `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetHarborProjectResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectResponse) GoString() string {
  return s.String()
}

func (s *GetHarborProjectResponse) SetCode(v int64) *GetHarborProjectResponse {
  s.Code = &v
  return s
}

func (s *GetHarborProjectResponse) SetMsg(v string) *GetHarborProjectResponse {
  s.Msg = &v
  return s
}

func (s *GetHarborProjectResponse) SetRequestId(v string) *GetHarborProjectResponse {
  s.RequestId = &v
  return s
}

func (s *GetHarborProjectResponse) SetData(v *GetHarborProjectProjectDetail) *GetHarborProjectResponse {
  s.Data = v
  return s
}

type GetHarborProjectPaths struct {
  // {"en":"project name", "zh_CN":"项目名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetHarborProjectPaths) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectPaths) GoString() string {
  return s.String()
}

func (s *GetHarborProjectPaths) SetName(v string) *GetHarborProjectPaths {
  s.Name = &v
  return s
}

type GetHarborProjectParameters struct {
}

func (s GetHarborProjectParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectParameters) GoString() string {
  return s.String()
}

type GetHarborProjectRequestHeader struct {
}

func (s GetHarborProjectRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectRequestHeader) GoString() string {
  return s.String()
}

type GetHarborProjectResponseHeader struct {
}

func (s GetHarborProjectResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectResponseHeader) GoString() string {
  return s.String()
}

type GetHarborProjectProjectDetail struct {
  // {"en":"project id", "zh_CN":"项目id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"project name", "zh_CN":"项目名称"}
  ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty" require:"true"`
  // {"en":"0: private, 1: public", "zh_CN":"0: 私有, 1: 公开"}
  Pub *bool `json:"pub,omitempty" xml:"pub,omitempty" require:"true"`
  // {"en":"project describe", "zh_CN":"项目描述"}
  Describe *string `json:"describe,omitempty" xml:"describe,omitempty" require:"true"`
  // {"en":"project createTime", "zh_CN":"项目创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"project updateTime", "zh_CN":"项目更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"authorized user", "zh_CN":"授权用户"}
  AssignRoles []*GetHarborProjectAssignRole `json:"assignRoles,omitempty" xml:"assignRoles,omitempty" require:"true" type:"Repeated"`
}

func (s GetHarborProjectProjectDetail) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectProjectDetail) GoString() string {
  return s.String()
}

func (s *GetHarborProjectProjectDetail) SetId(v int64) *GetHarborProjectProjectDetail {
  s.Id = &v
  return s
}

func (s *GetHarborProjectProjectDetail) SetProjectName(v string) *GetHarborProjectProjectDetail {
  s.ProjectName = &v
  return s
}

func (s *GetHarborProjectProjectDetail) SetPub(v bool) *GetHarborProjectProjectDetail {
  s.Pub = &v
  return s
}

func (s *GetHarborProjectProjectDetail) SetDescribe(v string) *GetHarborProjectProjectDetail {
  s.Describe = &v
  return s
}

func (s *GetHarborProjectProjectDetail) SetCreateTime(v int64) *GetHarborProjectProjectDetail {
  s.CreateTime = &v
  return s
}

func (s *GetHarborProjectProjectDetail) SetUpdateTime(v int64) *GetHarborProjectProjectDetail {
  s.UpdateTime = &v
  return s
}

func (s *GetHarborProjectProjectDetail) SetAssignRoles(v []*GetHarborProjectAssignRole) *GetHarborProjectProjectDetail {
  s.AssignRoles = v
  return s
}

type GetHarborProjectAssignRole struct {
  // {"en":"user id", "zh_CN":"用户id"}
  RepoUserId *int64 `json:"repoUserId,omitempty" xml:"repoUserId,omitempty" require:"true"`
  // {"en":"user name", "zh_CN":"用户名称"}
  RepoUserName *string `json:"repoUserName,omitempty" xml:"repoUserName,omitempty" require:"true"`
  // {"en":"1: admin, 2: edit, 3: access", "zh_CN":"1: 管理,2: 编辑, 3：访问"}
  Role *int32 `json:"role,omitempty" xml:"role,omitempty" require:"true"`
  // {"en":"project createTime", "zh_CN":"项目创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"project updateTime", "zh_CN":"项目更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s GetHarborProjectAssignRole) String() string {
  return tea.Prettify(s)
}

func (s GetHarborProjectAssignRole) GoString() string {
  return s.String()
}

func (s *GetHarborProjectAssignRole) SetRepoUserId(v int64) *GetHarborProjectAssignRole {
  s.RepoUserId = &v
  return s
}

func (s *GetHarborProjectAssignRole) SetRepoUserName(v string) *GetHarborProjectAssignRole {
  s.RepoUserName = &v
  return s
}

func (s *GetHarborProjectAssignRole) SetRole(v int32) *GetHarborProjectAssignRole {
  s.Role = &v
  return s
}

func (s *GetHarborProjectAssignRole) SetCreateTime(v int64) *GetHarborProjectAssignRole {
  s.CreateTime = &v
  return s
}

func (s *GetHarborProjectAssignRole) SetUpdateTime(v int64) *GetHarborProjectAssignRole {
  s.UpdateTime = &v
  return s
}




type ResetHarborUserPwdRequest struct {
  // {"en":"user name", "zh_CN":"用户名称"}
  Username *string `json:"username,omitempty" xml:"username,omitempty" require:"true"`
}

func (s ResetHarborUserPwdRequest) String() string {
  return tea.Prettify(s)
}

func (s ResetHarborUserPwdRequest) GoString() string {
  return s.String()
}

func (s *ResetHarborUserPwdRequest) SetUsername(v string) *ResetHarborUserPwdRequest {
  s.Username = &v
  return s
}

type ResetHarborUserPwdResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"password", "zh_CN":"密码"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ResetHarborUserPwdResponse) String() string {
  return tea.Prettify(s)
}

func (s ResetHarborUserPwdResponse) GoString() string {
  return s.String()
}

func (s *ResetHarborUserPwdResponse) SetCode(v int64) *ResetHarborUserPwdResponse {
  s.Code = &v
  return s
}

func (s *ResetHarborUserPwdResponse) SetMsg(v string) *ResetHarborUserPwdResponse {
  s.Msg = &v
  return s
}

func (s *ResetHarborUserPwdResponse) SetRequestId(v string) *ResetHarborUserPwdResponse {
  s.RequestId = &v
  return s
}

func (s *ResetHarborUserPwdResponse) SetData(v string) *ResetHarborUserPwdResponse {
  s.Data = &v
  return s
}

type ResetHarborUserPwdPaths struct {
}

func (s ResetHarborUserPwdPaths) String() string {
  return tea.Prettify(s)
}

func (s ResetHarborUserPwdPaths) GoString() string {
  return s.String()
}

type ResetHarborUserPwdParameters struct {
}

func (s ResetHarborUserPwdParameters) String() string {
  return tea.Prettify(s)
}

func (s ResetHarborUserPwdParameters) GoString() string {
  return s.String()
}

type ResetHarborUserPwdRequestHeader struct {
}

func (s ResetHarborUserPwdRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ResetHarborUserPwdRequestHeader) GoString() string {
  return s.String()
}

type ResetHarborUserPwdResponseHeader struct {
}

func (s ResetHarborUserPwdResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ResetHarborUserPwdResponseHeader) GoString() string {
  return s.String()
}




type VMPRemoveImageRequest struct {
  // {"en":"Image unique identification", "zh_CN":"镜像唯一标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
}

func (s VMPRemoveImageRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageRequest) GoString() string {
  return s.String()
}

func (s *VMPRemoveImageRequest) SetImageId(v string) *VMPRemoveImageRequest {
  s.ImageId = &v
  return s
}

type VMPRemoveImageResponse struct {
}

func (s VMPRemoveImageResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageResponse) GoString() string {
  return s.String()
}

type VMPRemoveImagePaths struct {
}

func (s VMPRemoveImagePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImagePaths) GoString() string {
  return s.String()
}

type VMPRemoveImageParameters struct {
}

func (s VMPRemoveImageParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageParameters) GoString() string {
  return s.String()
}

type VMPRemoveImageRequestHeader struct {
}

func (s VMPRemoveImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageRequestHeader) GoString() string {
  return s.String()
}

type VMPRemoveImageResponseHeader struct {
}

func (s VMPRemoveImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryImagePreheatingStateRequest struct {
}

func (s LECHQueryImagePreheatingStateRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateRequest) GoString() string {
  return s.String()
}

type LECHQueryImagePreheatingStateRequestHeader struct {
}

func (s LECHQueryImagePreheatingStateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryImagePreheatingStatePaths struct {
  // {"en":"no","zh_CN":"镜像id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHQueryImagePreheatingStatePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStatePaths) GoString() string {
  return s.String()
}

func (s *LECHQueryImagePreheatingStatePaths) SetId(v string) *LECHQueryImagePreheatingStatePaths {
  s.Id = &v
  return s
}

type LECHQueryImagePreheatingStateParameters struct {
}

func (s LECHQueryImagePreheatingStateParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateParameters) GoString() string {
  return s.String()
}

type LECHQueryImagePreheatingStateResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHQueryImagePreheatingStateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHQueryImagePreheatingStateResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryImagePreheatingStateResponse) SetCode(v string) *LECHQueryImagePreheatingStateResponse {
  s.Code = &v
  return s
}

func (s *LECHQueryImagePreheatingStateResponse) SetMessage(v string) *LECHQueryImagePreheatingStateResponse {
  s.Message = &v
  return s
}

func (s *LECHQueryImagePreheatingStateResponse) SetData(v *LECHQueryImagePreheatingStateResponseData) *LECHQueryImagePreheatingStateResponse {
  s.Data = v
  return s
}

type LECHQueryImagePreheatingStateResponseData struct {
  // {"en":"Preheating info","zh_CN":"预热信息"}
  PreHeatingInfo []*LECHQueryImagePreheatingStateResponseDataPreHeatingInfo `json:"preHeatingInfo,omitempty" xml:"preHeatingInfo,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryImagePreheatingStateResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateResponseData) GoString() string {
  return s.String()
}

func (s *LECHQueryImagePreheatingStateResponseData) SetPreHeatingInfo(v []*LECHQueryImagePreheatingStateResponseDataPreHeatingInfo) *LECHQueryImagePreheatingStateResponseData {
  s.PreHeatingInfo = v
  return s
}

type LECHQueryImagePreheatingStateResponseDataPreHeatingInfo struct     {
  // {"en":"Node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"State","zh_CN":"状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s LECHQueryImagePreheatingStateResponseDataPreHeatingInfo) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateResponseDataPreHeatingInfo) GoString() string {
  return s.String()
}

func (s *LECHQueryImagePreheatingStateResponseDataPreHeatingInfo) SetNodeName(v string) *LECHQueryImagePreheatingStateResponseDataPreHeatingInfo {
  s.NodeName = &v
  return s
}

func (s *LECHQueryImagePreheatingStateResponseDataPreHeatingInfo) SetState(v string) *LECHQueryImagePreheatingStateResponseDataPreHeatingInfo {
  s.State = &v
  return s
}

type LECHQueryImagePreheatingStateResponseHeader struct {
}

func (s LECHQueryImagePreheatingStateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePreheatingStateResponseHeader) GoString() string {
  return s.String()
}




type LECHRemoveImageRequest struct {
  // {"en":"Image unique identification","zh_CN":"镜像唯一标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
}

func (s LECHRemoveImageRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveImageRequest) GoString() string {
  return s.String()
}

func (s *LECHRemoveImageRequest) SetImageId(v string) *LECHRemoveImageRequest {
  s.ImageId = &v
  return s
}

type LECHRemoveImageRequestHeader struct {
}

func (s LECHRemoveImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveImageRequestHeader) GoString() string {
  return s.String()
}

type LECHRemoveImagePaths struct {
}

func (s LECHRemoveImagePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveImagePaths) GoString() string {
  return s.String()
}

type LECHRemoveImageParameters struct {
}

func (s LECHRemoveImageParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveImageParameters) GoString() string {
  return s.String()
}

type LECHRemoveImageResponse struct {
}

func (s LECHRemoveImageResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveImageResponse) GoString() string {
  return s.String()
}

type LECHRemoveImageResponseHeader struct {
}

func (s LECHRemoveImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveImageResponseHeader) GoString() string {
  return s.String()
}




type CreateHarborUserRequest struct {
  // {"en":"user name", "zh_CN":"用户名称"}
  Username *string `json:"username,omitempty" xml:"username,omitempty" require:"true"`
}

func (s CreateHarborUserRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborUserRequest) GoString() string {
  return s.String()
}

func (s *CreateHarborUserRequest) SetUsername(v string) *CreateHarborUserRequest {
  s.Username = &v
  return s
}

type CreateHarborUserResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"password", "zh_CN":"密码"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateHarborUserResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborUserResponse) GoString() string {
  return s.String()
}

func (s *CreateHarborUserResponse) SetCode(v int64) *CreateHarborUserResponse {
  s.Code = &v
  return s
}

func (s *CreateHarborUserResponse) SetMsg(v string) *CreateHarborUserResponse {
  s.Msg = &v
  return s
}

func (s *CreateHarborUserResponse) SetRequestId(v string) *CreateHarborUserResponse {
  s.RequestId = &v
  return s
}

func (s *CreateHarborUserResponse) SetData(v string) *CreateHarborUserResponse {
  s.Data = &v
  return s
}

type CreateHarborUserPaths struct {
}

func (s CreateHarborUserPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborUserPaths) GoString() string {
  return s.String()
}

type CreateHarborUserParameters struct {
}

func (s CreateHarborUserParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborUserParameters) GoString() string {
  return s.String()
}

type CreateHarborUserRequestHeader struct {
}

func (s CreateHarborUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborUserRequestHeader) GoString() string {
  return s.String()
}

type CreateHarborUserResponseHeader struct {
}

func (s CreateHarborUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateHarborUserResponseHeader) GoString() string {
  return s.String()
}




type ListImagePullJobRequest struct {
}

func (s ListImagePullJobRequest) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobRequest) GoString() string {
  return s.String()
}

type ListImagePullJobResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"imagepulljob list", "zh_CN":"imagepulljob列表"}
  Data *ListImagePullJobImagePullJobList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListImagePullJobResponse) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobResponse) GoString() string {
  return s.String()
}

func (s *ListImagePullJobResponse) SetCode(v int64) *ListImagePullJobResponse {
  s.Code = &v
  return s
}

func (s *ListImagePullJobResponse) SetMsg(v string) *ListImagePullJobResponse {
  s.Msg = &v
  return s
}

func (s *ListImagePullJobResponse) SetRequestId(v string) *ListImagePullJobResponse {
  s.RequestId = &v
  return s
}

func (s *ListImagePullJobResponse) SetData(v *ListImagePullJobImagePullJobList) *ListImagePullJobResponse {
  s.Data = v
  return s
}

type ListImagePullJobPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s ListImagePullJobPaths) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobPaths) GoString() string {
  return s.String()
}

func (s *ListImagePullJobPaths) SetNamespace(v string) *ListImagePullJobPaths {
  s.Namespace = &v
  return s
}

type ListImagePullJobParameters struct {
}

func (s ListImagePullJobParameters) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobParameters) GoString() string {
  return s.String()
}

type ListImagePullJobRequestHeader struct {
}

func (s ListImagePullJobRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobRequestHeader) GoString() string {
  return s.String()
}

type ListImagePullJobResponseHeader struct {
}

func (s ListImagePullJobResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobResponseHeader) GoString() string {
  return s.String()
}

type ListImagePullJobImagePullJobList struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard list metadata", "zh_CN":"标准列表元数据"}
  Metadata *ListImagePullJobListMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"List of ListImagePullJobImagePullJob", "zh_CN":"ListImagePullJobImagePullJob 列表"}
  Items []*ListImagePullJobImagePullJob `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListImagePullJobImagePullJobList) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobImagePullJobList) GoString() string {
  return s.String()
}

func (s *ListImagePullJobImagePullJobList) SetApiVersion(v string) *ListImagePullJobImagePullJobList {
  s.ApiVersion = &v
  return s
}

func (s *ListImagePullJobImagePullJobList) SetKind(v string) *ListImagePullJobImagePullJobList {
  s.Kind = &v
  return s
}

func (s *ListImagePullJobImagePullJobList) SetMetadata(v *ListImagePullJobListMeta) *ListImagePullJobImagePullJobList {
  s.Metadata = v
  return s
}

func (s *ListImagePullJobImagePullJobList) SetItems(v []*ListImagePullJobImagePullJob) *ListImagePullJobImagePullJobList {
  s.Items = v
  return s
}

type ListImagePullJobListMeta struct {
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system", "zh_CN":"selfLink 表示此对象的 URL，由系统填充，只读。已弃用：selfLink 是一个遗留的只读字段，不再由系统填充。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only", "zh_CN":"标识该对象的服务器内部版本的字符串，客户端可以用该字段来确定对象何时被更改。 该值对客户端是不透明的，并且应该原样传回给服务器。该值由系统填充，只读"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message", "zh_CN":"如果用户对返回的条目数量设置了限制，则 continue 可能被设置，表示服务器有更多可用的数据。 该值是不透明的，可用于向提供此列表服务的端点发出另一个请求，以检索下一组可用的对象。 如果服务器配置已更改或时间已过去几分钟，则可能无法继续提供一致的列表。 除非你在错误消息中收到此令牌（token），否则使用此 continue 值时返回的 resourceVersion 字段应该和第一个响应中的值是相同的"}
  Continue *string `json:"continue,omitempty" xml:"continue,omitempty"`
  // {"en":"remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is estimating the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact", "zh_CN":"remainingItemCount 是列表中未包含在此列表响应中的后续项目的数量。 如果列表请求包含标签或字段选择器，则剩余项目的数量是未知的，并且在序列化期间该字段将保持未设置和省略。 如果列表是完整的（因为它没有分块或者这是最后一个块），那么就没有剩余的项目，并且在序列化过程中该字段将保持未设置和省略。 早于 v1.15 的服务器不设置此字段。remainingItemCount 的预期用途是估计集合的大小。 客户端不应依赖于设置准确的 remainingItemCount"}
  RemainingItemCount *int64 `json:"remainingItemCount,omitempty" xml:"remainingItemCount,omitempty"`
}

func (s ListImagePullJobListMeta) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobListMeta) GoString() string {
  return s.String()
}

func (s *ListImagePullJobListMeta) SetSelfLink(v string) *ListImagePullJobListMeta {
  s.SelfLink = &v
  return s
}

func (s *ListImagePullJobListMeta) SetResourceVersion(v string) *ListImagePullJobListMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListImagePullJobListMeta) SetContinue(v string) *ListImagePullJobListMeta {
  s.Continue = &v
  return s
}

func (s *ListImagePullJobListMeta) SetRemainingItemCount(v int64) *ListImagePullJobListMeta {
  s.RemainingItemCount = &v
  return s
}

type ListImagePullJobImagePullJob struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *ListImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the ListImagePullJobImagePullJob.", "zh_CN":"ListImagePullJobImagePullJob 预期行为的规约。"}
  Spec *ListImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListImagePullJobImagePullJob) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobImagePullJob) GoString() string {
  return s.String()
}

func (s *ListImagePullJobImagePullJob) SetApiVersion(v string) *ListImagePullJobImagePullJob {
  s.ApiVersion = &v
  return s
}

func (s *ListImagePullJobImagePullJob) SetKind(v string) *ListImagePullJobImagePullJob {
  s.Kind = &v
  return s
}

func (s *ListImagePullJobImagePullJob) SetMetadata(v *ListImagePullJobObjectMeta) *ListImagePullJobImagePullJob {
  s.Metadata = v
  return s
}

func (s *ListImagePullJobImagePullJob) SetSpec(v *ListImagePullJobImagePullJobSpec) *ListImagePullJobImagePullJob {
  s.Spec = v
  return s
}

type ListImagePullJobImagePullJobSpec struct {
  // {"en": "Image is the image to be pulled by the job", "zh_CN": "拉取镜像名"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en": "Parallelism is the requested parallelism, it can be set to any non-negative value. If it is unspecified, it defaults to 1. If it is specified as 0, then the Job is effectively paused until it is increased.The value range 0-10, +optional", "zh_CN": "并发拉取个数, 范围0-10"}
  Parallelism *ListImagePullJobIntstrIntOrString `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
  // {"en": "ListImagePullJobPullPolicy is an optional field to set parameters of the pulling task. If not specified, the system will use the default values.+optional", "zh_CN": "拉取策略"}
  ListImagePullJobPullPolicy *ListImagePullJobPullPolicy `json:"pullPolicy,omitempty" xml:"pullPolicy,omitempty"`
  // {"en": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling the image.If specified, these secrets will be passed to individual puller implementations for them to use.  For example,in the case of docker, only DockerConfig type secrets are honored.+optional", "zh_CN": "拉取镜像所需的密钥"}
  PullSecrets []*string `json:"pullSecrets,omitempty" xml:"pullSecrets,omitempty" type:"Repeated"`
}

func (s ListImagePullJobImagePullJobSpec) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobImagePullJobSpec) GoString() string {
  return s.String()
}

func (s *ListImagePullJobImagePullJobSpec) SetImage(v string) *ListImagePullJobImagePullJobSpec {
  s.Image = &v
  return s
}

func (s *ListImagePullJobImagePullJobSpec) SetParallelism(v *ListImagePullJobIntstrIntOrString) *ListImagePullJobImagePullJobSpec {
  s.Parallelism = v
  return s
}

func (s *ListImagePullJobImagePullJobSpec) SetPullPolicy(v *ListImagePullJobPullPolicy) *ListImagePullJobImagePullJobSpec {
  s.ListImagePullJobPullPolicy = v
  return s
}

func (s *ListImagePullJobImagePullJobSpec) SetPullSecrets(v []*string) *ListImagePullJobImagePullJobSpec {
  s.PullSecrets = v
  return s
}

type ListImagePullJobIntstrIntOrString struct {
  // {"en": "the integer value", "zh_CN": "整数值"}
  IntVal *int `json:"intVal,omitempty" xml:"intVal,omitempty"`
  // {"en": "the string value", "zh_CN": "字符串值"}
  StrVal *string `json:"strVal,omitempty" xml:"strVal,omitempty"`
  // {"en": "type", "zh_CN": "类型"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListImagePullJobIntstrIntOrString) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobIntstrIntOrString) GoString() string {
  return s.String()
}

func (s *ListImagePullJobIntstrIntOrString) SetIntVal(v int) *ListImagePullJobIntstrIntOrString {
  s.IntVal = &v
  return s
}

func (s *ListImagePullJobIntstrIntOrString) SetStrVal(v string) *ListImagePullJobIntstrIntOrString {
  s.StrVal = &v
  return s
}

func (s *ListImagePullJobIntstrIntOrString) SetType(v int) *ListImagePullJobIntstrIntOrString {
  s.Type = &v
  return s
}

type ListImagePullJobPullPolicy struct {
  // {"en": "Specifies the number of retries before marking the pulling task failed. Defaults to 3 +optional", "zh_CN": "backoff次数，默认3"}
  BackoffLimit *int `json:"backoffLimit,omitempty" xml:"backoffLimit,omitempty"`
  // {"en": "Specifies the timeout of the pulling task. Defaults to 600 +optional", "zh_CN": "拉取超时时间"}
  TimeoutSeconds *int `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s ListImagePullJobPullPolicy) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobPullPolicy) GoString() string {
  return s.String()
}

func (s *ListImagePullJobPullPolicy) SetBackoffLimit(v int) *ListImagePullJobPullPolicy {
  s.BackoffLimit = &v
  return s
}

func (s *ListImagePullJobPullPolicy) SetTimeoutSeconds(v int) *ListImagePullJobPullPolicy {
  s.TimeoutSeconds = &v
  return s
}

type ListImagePullJobObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*ListImagePullJobOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*ListImagePullJobManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s ListImagePullJobObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobObjectMeta) GoString() string {
  return s.String()
}

func (s *ListImagePullJobObjectMeta) SetName(v string) *ListImagePullJobObjectMeta {
  s.Name = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetGenerateName(v string) *ListImagePullJobObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetNamespace(v string) *ListImagePullJobObjectMeta {
  s.Namespace = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetSelfLink(v string) *ListImagePullJobObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetUid(v string) *ListImagePullJobObjectMeta {
  s.Uid = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetResourceVersion(v string) *ListImagePullJobObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetGeneration(v int64) *ListImagePullJobObjectMeta {
  s.Generation = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetCreationTimestamp(v string) *ListImagePullJobObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetDeletionTimestamp(v string) *ListImagePullJobObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetDeletionGracePeriodSeconds(v int64) *ListImagePullJobObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetLabels(v map[string]*string) *ListImagePullJobObjectMeta {
  s.Labels = v
  return s
}

func (s *ListImagePullJobObjectMeta) SetAnnotations(v map[string]*string) *ListImagePullJobObjectMeta {
  s.Annotations = v
  return s
}

func (s *ListImagePullJobObjectMeta) SetOwnerReferences(v []*ListImagePullJobOwnerReference) *ListImagePullJobObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *ListImagePullJobObjectMeta) SetFinalizers(v []*string) *ListImagePullJobObjectMeta {
  s.Finalizers = v
  return s
}

func (s *ListImagePullJobObjectMeta) SetClusterName(v string) *ListImagePullJobObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *ListImagePullJobObjectMeta) SetManagedFields(v []*ListImagePullJobManagedFieldsEntry) *ListImagePullJobObjectMeta {
  s.ManagedFields = v
  return s
}

type ListImagePullJobManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this ListImagePullJobManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'ListImagePullJobFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“ListImagePullJobFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"ListImagePullJobFieldsV1 holds the first JSON version format as described in the 'ListImagePullJobFieldsV1' type", "zh_CN":"ListImagePullJobFieldsV1 包含类型 “ListImagePullJobFieldsV1” 中描述的第一个 JSON 版本格式"}
  ListImagePullJobFieldsV1 *ListImagePullJobFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s ListImagePullJobManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *ListImagePullJobManagedFieldsEntry) SetManager(v string) *ListImagePullJobManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *ListImagePullJobManagedFieldsEntry) SetOperation(v string) *ListImagePullJobManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *ListImagePullJobManagedFieldsEntry) SetApiVersion(v string) *ListImagePullJobManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *ListImagePullJobManagedFieldsEntry) SetTime(v string) *ListImagePullJobManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *ListImagePullJobManagedFieldsEntry) SetFieldsType(v string) *ListImagePullJobManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *ListImagePullJobManagedFieldsEntry) SetFieldsV1(v *ListImagePullJobFieldsV1) *ListImagePullJobManagedFieldsEntry {
  s.ListImagePullJobFieldsV1 = v
  return s
}

func (s *ListImagePullJobManagedFieldsEntry) SetSubresource(v string) *ListImagePullJobManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type ListImagePullJobFieldsV1 struct {
}

func (s ListImagePullJobFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobFieldsV1) GoString() string {
  return s.String()
}

type ListImagePullJobOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s ListImagePullJobOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s ListImagePullJobOwnerReference) GoString() string {
  return s.String()
}

func (s *ListImagePullJobOwnerReference) SetApiVersion(v string) *ListImagePullJobOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *ListImagePullJobOwnerReference) SetKind(v string) *ListImagePullJobOwnerReference {
  s.Kind = &v
  return s
}

func (s *ListImagePullJobOwnerReference) SetName(v string) *ListImagePullJobOwnerReference {
  s.Name = &v
  return s
}

func (s *ListImagePullJobOwnerReference) SetUid(v string) *ListImagePullJobOwnerReference {
  s.Uid = &v
  return s
}

func (s *ListImagePullJobOwnerReference) SetController(v bool) *ListImagePullJobOwnerReference {
  s.Controller = &v
  return s
}

func (s *ListImagePullJobOwnerReference) SetBlockOwnerDeletion(v bool) *ListImagePullJobOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type DeleteHarborProjectRequest struct {
}

func (s DeleteHarborProjectRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborProjectRequest) GoString() string {
  return s.String()
}

type DeleteHarborProjectResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"none", "zh_CN":"无"}
  Data *int64 `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteHarborProjectResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborProjectResponse) GoString() string {
  return s.String()
}

func (s *DeleteHarborProjectResponse) SetCode(v int64) *DeleteHarborProjectResponse {
  s.Code = &v
  return s
}

func (s *DeleteHarborProjectResponse) SetMsg(v string) *DeleteHarborProjectResponse {
  s.Msg = &v
  return s
}

func (s *DeleteHarborProjectResponse) SetRequestId(v string) *DeleteHarborProjectResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteHarborProjectResponse) SetData(v int64) *DeleteHarborProjectResponse {
  s.Data = &v
  return s
}

type DeleteHarborProjectPaths struct {
  // {"en":"project name", "zh_CN":"项目名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteHarborProjectPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborProjectPaths) GoString() string {
  return s.String()
}

func (s *DeleteHarborProjectPaths) SetName(v string) *DeleteHarborProjectPaths {
  s.Name = &v
  return s
}

type DeleteHarborProjectParameters struct {
}

func (s DeleteHarborProjectParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborProjectParameters) GoString() string {
  return s.String()
}

type DeleteHarborProjectRequestHeader struct {
}

func (s DeleteHarborProjectRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborProjectRequestHeader) GoString() string {
  return s.String()
}

type DeleteHarborProjectResponseHeader struct {
}

func (s DeleteHarborProjectResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteHarborProjectResponseHeader) GoString() string {
  return s.String()
}




type LECHCreateImageRequest struct {
  // {"en":"Mirror name","zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Virtual machine instance ID, optional parameter. This parameter and imagesrcurl must be one of two choices","zh_CN":"虚拟机实例标识，可选参数，该参数与imageSrcUrl必须二选一"}
  InstanceUuid *string `json:"instanceUuid,omitempty" xml:"instanceUuid,omitempty"`
  // {"en":"Virtual machine image URL address, optional parameter. The parameter and instanceuuid must be one of two choices.","zh_CN":"虚拟机镜像Url地址，可选参数，该参数与instanceUuid必须二选一"}
  ImageSrcUrl *string `json:"imageSrcUrl,omitempty" xml:"imageSrcUrl,omitempty"`
  // {"en":"MD5 value of virtual machine image, used with imagesrcurl","zh_CN":"虚拟机镜像的md5值，与imageSrcUrl配合使用"}
  Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
  // {"en":"Operating system type, if the URL address is carried, this parameter is required; if the virtual machine is specified to be created, the ostype of the virtual machine shall prevail.\n\nThere are two values: windows and Linux'","zh_CN":"操作系统类型，如果携带的是url地址时，该参数必填；如果是指定虚拟机创建，则以虚拟机的ostype为准。\n有2种取值：windows、linux"}
  Ostype *string `json:"ostype,omitempty" xml:"ostype,omitempty"`
  // {"en":"Minimum requirements for system disk, unit: GB. The system disk size of the selected template must be greater than or equal to this value, otherwise virtual machine creation fails","zh_CN":"系统盘最小要求，单位是GB。选择的模板的系统盘大小必须大于等于该值，否则虚拟机创建失败"}
  MinDisk *int `json:"minDisk,omitempty" xml:"minDisk,omitempty"`
  // {"en":"Whether QEMU guest agent is enabled for the image, and password reset is supported for the opened image\n\nThere are two values: true and false'","zh_CN":"镜像是否开启了qemu guest agent，有开启的镜像支持密码重置\n有2种取值：TRUE、FALSE"}
  QgaEnabled *string `json:"qgaEnabled,omitempty" xml:"qgaEnabled,omitempty"`
}

func (s LECHCreateImageRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateImageRequest) GoString() string {
  return s.String()
}

func (s *LECHCreateImageRequest) SetName(v string) *LECHCreateImageRequest {
  s.Name = &v
  return s
}

func (s *LECHCreateImageRequest) SetInstanceUuid(v string) *LECHCreateImageRequest {
  s.InstanceUuid = &v
  return s
}

func (s *LECHCreateImageRequest) SetImageSrcUrl(v string) *LECHCreateImageRequest {
  s.ImageSrcUrl = &v
  return s
}

func (s *LECHCreateImageRequest) SetMd5(v string) *LECHCreateImageRequest {
  s.Md5 = &v
  return s
}

func (s *LECHCreateImageRequest) SetOstype(v string) *LECHCreateImageRequest {
  s.Ostype = &v
  return s
}

func (s *LECHCreateImageRequest) SetMinDisk(v int) *LECHCreateImageRequest {
  s.MinDisk = &v
  return s
}

func (s *LECHCreateImageRequest) SetQgaEnabled(v string) *LECHCreateImageRequest {
  s.QgaEnabled = &v
  return s
}

type LECHCreateImageRequestHeader struct {
}

func (s LECHCreateImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateImageRequestHeader) GoString() string {
  return s.String()
}

type LECHCreateImagePaths struct {
}

func (s LECHCreateImagePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateImagePaths) GoString() string {
  return s.String()
}

type LECHCreateImageParameters struct {
}

func (s LECHCreateImageParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateImageParameters) GoString() string {
  return s.String()
}

type LECHCreateImageResponse struct {
  // {"en":"Image unique ID, global unique","zh_CN":"镜像唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHCreateImageResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateImageResponse) GoString() string {
  return s.String()
}

func (s *LECHCreateImageResponse) SetId(v string) *LECHCreateImageResponse {
  s.Id = &v
  return s
}

type LECHCreateImageResponseHeader struct {
}

func (s LECHCreateImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateImageResponseHeader) GoString() string {
  return s.String()
}




type ListHarborProjectRequest struct {
}

func (s ListHarborProjectRequest) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectRequest) GoString() string {
  return s.String()
}

type ListHarborProjectResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"project list", "zh_CN":"项目列表"}
  Data *ListHarborProjectProjectList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListHarborProjectResponse) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectResponse) GoString() string {
  return s.String()
}

func (s *ListHarborProjectResponse) SetCode(v int64) *ListHarborProjectResponse {
  s.Code = &v
  return s
}

func (s *ListHarborProjectResponse) SetMsg(v string) *ListHarborProjectResponse {
  s.Msg = &v
  return s
}

func (s *ListHarborProjectResponse) SetRequestId(v string) *ListHarborProjectResponse {
  s.RequestId = &v
  return s
}

func (s *ListHarborProjectResponse) SetData(v *ListHarborProjectProjectList) *ListHarborProjectResponse {
  s.Data = v
  return s
}

type ListHarborProjectPaths struct {
}

func (s ListHarborProjectPaths) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectPaths) GoString() string {
  return s.String()
}

type ListHarborProjectParameters struct {
  // {"en":"search keyword", "zh_CN":"搜索关键字"}
  Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
  // {"en":"page index: default 0", "zh_CN":"分页页数: 默认0"}
  PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"page size: default 10", "zh_CN":"每页记录数: 默认10"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListHarborProjectParameters) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectParameters) GoString() string {
  return s.String()
}

func (s *ListHarborProjectParameters) SetKeyword(v string) *ListHarborProjectParameters {
  s.Keyword = &v
  return s
}

func (s *ListHarborProjectParameters) SetPageIndex(v string) *ListHarborProjectParameters {
  s.PageIndex = &v
  return s
}

func (s *ListHarborProjectParameters) SetPageSize(v string) *ListHarborProjectParameters {
  s.PageSize = &v
  return s
}

type ListHarborProjectRequestHeader struct {
}

func (s ListHarborProjectRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectRequestHeader) GoString() string {
  return s.String()
}

type ListHarborProjectResponseHeader struct {
}

func (s ListHarborProjectResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectResponseHeader) GoString() string {
  return s.String()
}

type ListHarborProjectProjectList struct {
  // {"en":"projects", "zh_CN":"项目列表"}
  Items []*ListHarborProjectProject `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
  // {"en":"total records count", "zh_CN":"记录总数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s ListHarborProjectProjectList) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectProjectList) GoString() string {
  return s.String()
}

func (s *ListHarborProjectProjectList) SetItems(v []*ListHarborProjectProject) *ListHarborProjectProjectList {
  s.Items = v
  return s
}

func (s *ListHarborProjectProjectList) SetTotal(v int64) *ListHarborProjectProjectList {
  s.Total = &v
  return s
}

type ListHarborProjectProject struct {
  // {"en":"project id", "zh_CN":"项目id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"project name", "zh_CN":"项目名称"}
  ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty" require:"true"`
  // {"en":"0: private, 1: public", "zh_CN":"0: 私有, 1: 公开"}
  Pub *bool `json:"pub,omitempty" xml:"pub,omitempty" require:"true"`
  // {"en":"project describe", "zh_CN":"项目描述"}
  Describe *string `json:"describe,omitempty" xml:"describe,omitempty" require:"true"`
  // {"en":"project createTime", "zh_CN":"项目创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"project updateTime", "zh_CN":"项目更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"project members", "zh_CN":"项目成员"}
  RepoUsers []*string `json:"repoUsers,omitempty" xml:"repoUsers,omitempty" require:"true" type:"Repeated"`
}

func (s ListHarborProjectProject) String() string {
  return tea.Prettify(s)
}

func (s ListHarborProjectProject) GoString() string {
  return s.String()
}

func (s *ListHarborProjectProject) SetId(v int64) *ListHarborProjectProject {
  s.Id = &v
  return s
}

func (s *ListHarborProjectProject) SetProjectName(v string) *ListHarborProjectProject {
  s.ProjectName = &v
  return s
}

func (s *ListHarborProjectProject) SetPub(v bool) *ListHarborProjectProject {
  s.Pub = &v
  return s
}

func (s *ListHarborProjectProject) SetDescribe(v string) *ListHarborProjectProject {
  s.Describe = &v
  return s
}

func (s *ListHarborProjectProject) SetCreateTime(v int64) *ListHarborProjectProject {
  s.CreateTime = &v
  return s
}

func (s *ListHarborProjectProject) SetUpdateTime(v int64) *ListHarborProjectProject {
  s.UpdateTime = &v
  return s
}

func (s *ListHarborProjectProject) SetRepoUsers(v []*string) *ListHarborProjectProject {
  s.RepoUsers = v
  return s
}




type LECHQueryImageRequest struct {
}

func (s LECHQueryImageRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImageRequest) GoString() string {
  return s.String()
}

type LECHQueryImageRequestHeader struct {
}

func (s LECHQueryImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImageRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryImagePaths struct {
}

func (s LECHQueryImagePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImagePaths) GoString() string {
  return s.String()
}

type LECHQueryImageParameters struct {
  // {"en":"There can be multiple field names for sorting. The values are:\nName, createdat, type, size, state","zh_CN":"排序的字段名称，可以有多个，取值：\nname、createdAt、type、size、state"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"Sorting direction must follow sortkey. Value:\nDesc: descending, default\nASC: ascending order","zh_CN":"排序方向，必须跟在sortKey后面出现，取值：\ndesc：降序，默认值\nasc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default","zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the image ID specified by the marker","zh_CN":"从marker指定的镜像id开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Mirror ID. A maximum of 100 IDS can be queried at a time. The IDs are separated by a half angle comma character ','.","zh_CN":"镜像 ID。单次最多查询 100 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"Mirror name","zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The image belongs to the master. Is it a public image or a user-defined image? Value:\n\nCommon: official image\n\nSnapshot: user snapshot image\n\nCustom: user defined image'","zh_CN":"镜像属主，是公共镜像还是用户自定义镜像，取值：\nCOMMON：官方镜像\nSNAPSHOT：用户快照镜像\nCUSTOM：用户自定义镜像"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Image status, value:\n\nActive: available status\n\nBuilding: Creating\n\nInactive: not available (such as creation failure, etc.)'","zh_CN":"镜像状态，取值：\nACTIVE：可用状态\nBUILDING：创建中\nINACTIVE：不可用（如创建失败等）"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s LECHQueryImageParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImageParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryImageParameters) SetSortKey(v string) *LECHQueryImageParameters {
  s.SortKey = &v
  return s
}

func (s *LECHQueryImageParameters) SetSortDir(v string) *LECHQueryImageParameters {
  s.SortDir = &v
  return s
}

func (s *LECHQueryImageParameters) SetLimit(v int) *LECHQueryImageParameters {
  s.Limit = &v
  return s
}

func (s *LECHQueryImageParameters) SetMarker(v string) *LECHQueryImageParameters {
  s.Marker = &v
  return s
}

func (s *LECHQueryImageParameters) SetIds(v string) *LECHQueryImageParameters {
  s.Ids = &v
  return s
}

func (s *LECHQueryImageParameters) SetName(v string) *LECHQueryImageParameters {
  s.Name = &v
  return s
}

func (s *LECHQueryImageParameters) SetType(v string) *LECHQueryImageParameters {
  s.Type = &v
  return s
}

func (s *LECHQueryImageParameters) SetState(v string) *LECHQueryImageParameters {
  s.State = &v
  return s
}

type LECHQueryImageResponse struct {
  // {"en":"Image information array","zh_CN":"镜像信息数组"}
  Images []*string `json:"images,omitempty" xml:"images,omitempty" require:"true" type:"Repeated"`
  // {"en":"Image unique ID, global unique","zh_CN":"镜像唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Mirror name","zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Mirror belongs to the Lord.","zh_CN":"镜像属主"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Image creation time, such as: 2017-11-04 14:17:41","zh_CN":"镜像创建时间，如：2017-11-04 14:17:41"}
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty" require:"true"`
  // {"en":"Image size in GB","zh_CN":"镜像大小，单位是GB"}
  Size *string `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Mirror Status","zh_CN":"镜像状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s LECHQueryImageResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImageResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryImageResponse) SetImages(v []*string) *LECHQueryImageResponse {
  s.Images = v
  return s
}

func (s *LECHQueryImageResponse) SetId(v string) *LECHQueryImageResponse {
  s.Id = &v
  return s
}

func (s *LECHQueryImageResponse) SetName(v string) *LECHQueryImageResponse {
  s.Name = &v
  return s
}

func (s *LECHQueryImageResponse) SetType(v string) *LECHQueryImageResponse {
  s.Type = &v
  return s
}

func (s *LECHQueryImageResponse) SetCreatedAt(v string) *LECHQueryImageResponse {
  s.CreatedAt = &v
  return s
}

func (s *LECHQueryImageResponse) SetSize(v string) *LECHQueryImageResponse {
  s.Size = &v
  return s
}

func (s *LECHQueryImageResponse) SetState(v string) *LECHQueryImageResponse {
  s.State = &v
  return s
}

type LECHQueryImageResponseHeader struct {
}

func (s LECHQueryImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryImageResponseHeader) GoString() string {
  return s.String()
}




type CreateImagePullJobRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *CreateImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the CreateImagePullJobImagePullJob.", "zh_CN":"CreateImagePullJobImagePullJob 预期行为的规约。"}
  Spec *CreateImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateImagePullJobRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobRequest) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobRequest) SetApiVersion(v string) *CreateImagePullJobRequest {
  s.ApiVersion = &v
  return s
}

func (s *CreateImagePullJobRequest) SetKind(v string) *CreateImagePullJobRequest {
  s.Kind = &v
  return s
}

func (s *CreateImagePullJobRequest) SetMetadata(v *CreateImagePullJobObjectMeta) *CreateImagePullJobRequest {
  s.Metadata = v
  return s
}

func (s *CreateImagePullJobRequest) SetSpec(v *CreateImagePullJobImagePullJobSpec) *CreateImagePullJobRequest {
  s.Spec = v
  return s
}

type CreateImagePullJobResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"imagepulljob object", "zh_CN":"imagepulljob对象"}
  Data *CreateImagePullJobImagePullJob `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateImagePullJobResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobResponse) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobResponse) SetCode(v int64) *CreateImagePullJobResponse {
  s.Code = &v
  return s
}

func (s *CreateImagePullJobResponse) SetMsg(v string) *CreateImagePullJobResponse {
  s.Msg = &v
  return s
}

func (s *CreateImagePullJobResponse) SetRequestId(v string) *CreateImagePullJobResponse {
  s.RequestId = &v
  return s
}

func (s *CreateImagePullJobResponse) SetData(v *CreateImagePullJobImagePullJob) *CreateImagePullJobResponse {
  s.Data = v
  return s
}

type CreateImagePullJobPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s CreateImagePullJobPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobPaths) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobPaths) SetNamespace(v string) *CreateImagePullJobPaths {
  s.Namespace = &v
  return s
}

type CreateImagePullJobParameters struct {
}

func (s CreateImagePullJobParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobParameters) GoString() string {
  return s.String()
}

type CreateImagePullJobRequestHeader struct {
}

func (s CreateImagePullJobRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobRequestHeader) GoString() string {
  return s.String()
}

type CreateImagePullJobResponseHeader struct {
}

func (s CreateImagePullJobResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobResponseHeader) GoString() string {
  return s.String()
}

type CreateImagePullJobImagePullJob struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard object metadata.", "zh_CN":"标准的对象元数据。"}
  Metadata *CreateImagePullJobObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Specification of the desired behavior of the CreateImagePullJobImagePullJob.", "zh_CN":"CreateImagePullJobImagePullJob 预期行为的规约。"}
  Spec *CreateImagePullJobImagePullJobSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the CreateImagePullJobImagePullJob.", "zh_CN":"最近观测到的 CreateImagePullJobImagePullJob 状态。"}
  Status *CreateImagePullJobImagePullJobStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateImagePullJobImagePullJob) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobImagePullJob) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobImagePullJob) SetApiVersion(v string) *CreateImagePullJobImagePullJob {
  s.ApiVersion = &v
  return s
}

func (s *CreateImagePullJobImagePullJob) SetKind(v string) *CreateImagePullJobImagePullJob {
  s.Kind = &v
  return s
}

func (s *CreateImagePullJobImagePullJob) SetMetadata(v *CreateImagePullJobObjectMeta) *CreateImagePullJobImagePullJob {
  s.Metadata = v
  return s
}

func (s *CreateImagePullJobImagePullJob) SetSpec(v *CreateImagePullJobImagePullJobSpec) *CreateImagePullJobImagePullJob {
  s.Spec = v
  return s
}

func (s *CreateImagePullJobImagePullJob) SetStatus(v *CreateImagePullJobImagePullJobStatus) *CreateImagePullJobImagePullJob {
  s.Status = v
  return s
}

type CreateImagePullJobImagePullJobSpec struct {
  // {"en": "Image is the image to be pulled by the job", "zh_CN": "拉取镜像名"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en": "Parallelism is the requested parallelism, it can be set to any non-negative value. If it is unspecified, it defaults to 1. If it is specified as 0, then the Job is effectively paused until it is increased.The value range 0-10, +optional", "zh_CN": "并发拉取个数, 范围0-10"}
  Parallelism *CreateImagePullJobIntstrIntOrString `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
  // {"en": "CreateImagePullJobPullPolicy is an optional field to set parameters of the pulling task. If not specified, the system will use the default values.+optional", "zh_CN": "拉取策略"}
  CreateImagePullJobPullPolicy *CreateImagePullJobPullPolicy `json:"pullPolicy,omitempty" xml:"pullPolicy,omitempty"`
  // {"en": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling the image.If specified, these secrets will be passed to individual puller implementations for them to use.  For example,in the case of docker, only DockerConfig type secrets are honored.+optional", "zh_CN": "拉取镜像所需的密钥"}
  PullSecrets []*string `json:"pullSecrets,omitempty" xml:"pullSecrets,omitempty" type:"Repeated"`
}

func (s CreateImagePullJobImagePullJobSpec) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobImagePullJobSpec) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobImagePullJobSpec) SetImage(v string) *CreateImagePullJobImagePullJobSpec {
  s.Image = &v
  return s
}

func (s *CreateImagePullJobImagePullJobSpec) SetParallelism(v *CreateImagePullJobIntstrIntOrString) *CreateImagePullJobImagePullJobSpec {
  s.Parallelism = v
  return s
}

func (s *CreateImagePullJobImagePullJobSpec) SetPullPolicy(v *CreateImagePullJobPullPolicy) *CreateImagePullJobImagePullJobSpec {
  s.CreateImagePullJobPullPolicy = v
  return s
}

func (s *CreateImagePullJobImagePullJobSpec) SetPullSecrets(v []*string) *CreateImagePullJobImagePullJobSpec {
  s.PullSecrets = v
  return s
}

type CreateImagePullJobIntstrIntOrString struct {
  // {"en": "the integer value", "zh_CN": "整数值"}
  IntVal *int `json:"intVal,omitempty" xml:"intVal,omitempty"`
  // {"en": "the string value", "zh_CN": "字符串值"}
  StrVal *string `json:"strVal,omitempty" xml:"strVal,omitempty"`
  // {"en": "type", "zh_CN": "类型"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateImagePullJobIntstrIntOrString) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobIntstrIntOrString) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobIntstrIntOrString) SetIntVal(v int) *CreateImagePullJobIntstrIntOrString {
  s.IntVal = &v
  return s
}

func (s *CreateImagePullJobIntstrIntOrString) SetStrVal(v string) *CreateImagePullJobIntstrIntOrString {
  s.StrVal = &v
  return s
}

func (s *CreateImagePullJobIntstrIntOrString) SetType(v int) *CreateImagePullJobIntstrIntOrString {
  s.Type = &v
  return s
}

type CreateImagePullJobPullPolicy struct {
  // {"en": "Specifies the number of retries before marking the pulling task failed. Defaults to 3 +optional", "zh_CN": "backoff次数，默认3"}
  BackoffLimit *int `json:"backoffLimit,omitempty" xml:"backoffLimit,omitempty"`
  // {"en": "Specifies the timeout of the pulling task. Defaults to 600 +optional", "zh_CN": "拉取超时时间"}
  TimeoutSeconds *int `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s CreateImagePullJobPullPolicy) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobPullPolicy) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobPullPolicy) SetBackoffLimit(v int) *CreateImagePullJobPullPolicy {
  s.BackoffLimit = &v
  return s
}

func (s *CreateImagePullJobPullPolicy) SetTimeoutSeconds(v int) *CreateImagePullJobPullPolicy {
  s.TimeoutSeconds = &v
  return s
}

type CreateImagePullJobImagePullJobStatus struct {
  // {"en": "Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务完成时间"}
  CompletionTime *string `json:"completionTime,omitempty" xml:"completionTime,omitempty"`
  // {"en": "The text prompt for job running status.+optional", "zh_CN": "状态消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // {"en": "Represents time when the job was acknowledged by the job controller.It is not guaranteed to be set in happens-before order across separate operations.It is represented in RFC3339 form and is in UTC.+optional", "zh_CN": "任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en": "The rate of pulling tasks which reached phase Succeeded without not ready nodes. +optional", "zh_CN": "完成成功率"}
  SucceededRate *string `json:"succeededRate,omitempty" xml:"succeededRate,omitempty"`
}

func (s CreateImagePullJobImagePullJobStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobImagePullJobStatus) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobImagePullJobStatus) SetCompletionTime(v string) *CreateImagePullJobImagePullJobStatus {
  s.CompletionTime = &v
  return s
}

func (s *CreateImagePullJobImagePullJobStatus) SetMessage(v string) *CreateImagePullJobImagePullJobStatus {
  s.Message = &v
  return s
}

func (s *CreateImagePullJobImagePullJobStatus) SetStartTime(v string) *CreateImagePullJobImagePullJobStatus {
  s.StartTime = &v
  return s
}

func (s *CreateImagePullJobImagePullJobStatus) SetSucceededRate(v string) *CreateImagePullJobImagePullJobStatus {
  s.SucceededRate = &v
  return s
}

type CreateImagePullJobObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*CreateImagePullJobOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*CreateImagePullJobManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s CreateImagePullJobObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobObjectMeta) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobObjectMeta) SetName(v string) *CreateImagePullJobObjectMeta {
  s.Name = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetGenerateName(v string) *CreateImagePullJobObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetNamespace(v string) *CreateImagePullJobObjectMeta {
  s.Namespace = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetSelfLink(v string) *CreateImagePullJobObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetUid(v string) *CreateImagePullJobObjectMeta {
  s.Uid = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetResourceVersion(v string) *CreateImagePullJobObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetGeneration(v int64) *CreateImagePullJobObjectMeta {
  s.Generation = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetCreationTimestamp(v string) *CreateImagePullJobObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetDeletionTimestamp(v string) *CreateImagePullJobObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetDeletionGracePeriodSeconds(v int64) *CreateImagePullJobObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetLabels(v map[string]*string) *CreateImagePullJobObjectMeta {
  s.Labels = v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetAnnotations(v map[string]*string) *CreateImagePullJobObjectMeta {
  s.Annotations = v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetOwnerReferences(v []*CreateImagePullJobOwnerReference) *CreateImagePullJobObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetFinalizers(v []*string) *CreateImagePullJobObjectMeta {
  s.Finalizers = v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetClusterName(v string) *CreateImagePullJobObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *CreateImagePullJobObjectMeta) SetManagedFields(v []*CreateImagePullJobManagedFieldsEntry) *CreateImagePullJobObjectMeta {
  s.ManagedFields = v
  return s
}

type CreateImagePullJobManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this CreateImagePullJobManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'CreateImagePullJobFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“CreateImagePullJobFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"CreateImagePullJobFieldsV1 holds the first JSON version format as described in the 'CreateImagePullJobFieldsV1' type", "zh_CN":"CreateImagePullJobFieldsV1 包含类型 “CreateImagePullJobFieldsV1” 中描述的第一个 JSON 版本格式"}
  CreateImagePullJobFieldsV1 *CreateImagePullJobFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s CreateImagePullJobManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobManagedFieldsEntry) SetManager(v string) *CreateImagePullJobManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *CreateImagePullJobManagedFieldsEntry) SetOperation(v string) *CreateImagePullJobManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *CreateImagePullJobManagedFieldsEntry) SetApiVersion(v string) *CreateImagePullJobManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *CreateImagePullJobManagedFieldsEntry) SetTime(v string) *CreateImagePullJobManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *CreateImagePullJobManagedFieldsEntry) SetFieldsType(v string) *CreateImagePullJobManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *CreateImagePullJobManagedFieldsEntry) SetFieldsV1(v *CreateImagePullJobFieldsV1) *CreateImagePullJobManagedFieldsEntry {
  s.CreateImagePullJobFieldsV1 = v
  return s
}

func (s *CreateImagePullJobManagedFieldsEntry) SetSubresource(v string) *CreateImagePullJobManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type CreateImagePullJobFieldsV1 struct {
}

func (s CreateImagePullJobFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobFieldsV1) GoString() string {
  return s.String()
}

type CreateImagePullJobOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s CreateImagePullJobOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s CreateImagePullJobOwnerReference) GoString() string {
  return s.String()
}

func (s *CreateImagePullJobOwnerReference) SetApiVersion(v string) *CreateImagePullJobOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *CreateImagePullJobOwnerReference) SetKind(v string) *CreateImagePullJobOwnerReference {
  s.Kind = &v
  return s
}

func (s *CreateImagePullJobOwnerReference) SetName(v string) *CreateImagePullJobOwnerReference {
  s.Name = &v
  return s
}

func (s *CreateImagePullJobOwnerReference) SetUid(v string) *CreateImagePullJobOwnerReference {
  s.Uid = &v
  return s
}

func (s *CreateImagePullJobOwnerReference) SetController(v bool) *CreateImagePullJobOwnerReference {
  s.Controller = &v
  return s
}

func (s *CreateImagePullJobOwnerReference) SetBlockOwnerDeletion(v bool) *CreateImagePullJobOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type UpdateHarborProjectRequest struct {
  // {"en":"project name", "zh_CN":"项目名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"false: private, true: public", "zh_CN":"false: 私有,true:公开"}
  Pub *bool `json:"pub,omitempty" xml:"pub,omitempty" require:"true"`
  // {"en":"project describe", "zh_CN":"项目描述"}
  Describe *string `json:"describe,omitempty" xml:"describe,omitempty"`
  // {"en":"authorized user", "zh_CN":"授权用户"}
  AssignRoles []*UpdateHarborProjectAssignRole `json:"assignRoles,omitempty" xml:"assignRoles,omitempty" type:"Repeated"`
}

func (s UpdateHarborProjectRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectRequest) GoString() string {
  return s.String()
}

func (s *UpdateHarborProjectRequest) SetName(v string) *UpdateHarborProjectRequest {
  s.Name = &v
  return s
}

func (s *UpdateHarborProjectRequest) SetPub(v bool) *UpdateHarborProjectRequest {
  s.Pub = &v
  return s
}

func (s *UpdateHarborProjectRequest) SetDescribe(v string) *UpdateHarborProjectRequest {
  s.Describe = &v
  return s
}

func (s *UpdateHarborProjectRequest) SetAssignRoles(v []*UpdateHarborProjectAssignRole) *UpdateHarborProjectRequest {
  s.AssignRoles = v
  return s
}

type UpdateHarborProjectResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"none", "zh_CN":"无"}
  Data *int64 `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateHarborProjectResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectResponse) GoString() string {
  return s.String()
}

func (s *UpdateHarborProjectResponse) SetCode(v int64) *UpdateHarborProjectResponse {
  s.Code = &v
  return s
}

func (s *UpdateHarborProjectResponse) SetMsg(v string) *UpdateHarborProjectResponse {
  s.Msg = &v
  return s
}

func (s *UpdateHarborProjectResponse) SetRequestId(v string) *UpdateHarborProjectResponse {
  s.RequestId = &v
  return s
}

func (s *UpdateHarborProjectResponse) SetData(v int64) *UpdateHarborProjectResponse {
  s.Data = &v
  return s
}

type UpdateHarborProjectPaths struct {
}

func (s UpdateHarborProjectPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectPaths) GoString() string {
  return s.String()
}

type UpdateHarborProjectParameters struct {
}

func (s UpdateHarborProjectParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectParameters) GoString() string {
  return s.String()
}

type UpdateHarborProjectRequestHeader struct {
}

func (s UpdateHarborProjectRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectRequestHeader) GoString() string {
  return s.String()
}

type UpdateHarborProjectResponseHeader struct {
}

func (s UpdateHarborProjectResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectResponseHeader) GoString() string {
  return s.String()
}

type UpdateHarborProjectAssignRole struct {
  // {"en":"user name", "zh_CN":"用户名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"1: admin, 2: edit, 3: access", "zh_CN":"1: 管理,2: 编辑, 3：访问"}
  Role *int32 `json:"role,omitempty" xml:"role,omitempty" require:"true"`
}

func (s UpdateHarborProjectAssignRole) String() string {
  return tea.Prettify(s)
}

func (s UpdateHarborProjectAssignRole) GoString() string {
  return s.String()
}

func (s *UpdateHarborProjectAssignRole) SetName(v string) *UpdateHarborProjectAssignRole {
  s.Name = &v
  return s
}

func (s *UpdateHarborProjectAssignRole) SetRole(v int32) *UpdateHarborProjectAssignRole {
  s.Role = &v
  return s
}




type GetHarborUserRequest struct {
}

func (s GetHarborUserRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserRequest) GoString() string {
  return s.String()
}

type GetHarborUserResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"user detail", "zh_CN":"用户详情"}
  Data *GetHarborUserUser `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetHarborUserResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserResponse) GoString() string {
  return s.String()
}

func (s *GetHarborUserResponse) SetCode(v int64) *GetHarborUserResponse {
  s.Code = &v
  return s
}

func (s *GetHarborUserResponse) SetMsg(v string) *GetHarborUserResponse {
  s.Msg = &v
  return s
}

func (s *GetHarborUserResponse) SetRequestId(v string) *GetHarborUserResponse {
  s.RequestId = &v
  return s
}

func (s *GetHarborUserResponse) SetData(v *GetHarborUserUser) *GetHarborUserResponse {
  s.Data = v
  return s
}

type GetHarborUserPaths struct {
  // {"en":"user name", "zh_CN":"用户名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetHarborUserPaths) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserPaths) GoString() string {
  return s.String()
}

func (s *GetHarborUserPaths) SetName(v string) *GetHarborUserPaths {
  s.Name = &v
  return s
}

type GetHarborUserParameters struct {
}

func (s GetHarborUserParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserParameters) GoString() string {
  return s.String()
}

type GetHarborUserRequestHeader struct {
}

func (s GetHarborUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserRequestHeader) GoString() string {
  return s.String()
}

type GetHarborUserResponseHeader struct {
}

func (s GetHarborUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserResponseHeader) GoString() string {
  return s.String()
}

type GetHarborUserUser struct {
  // {"en":"user id", "zh_CN":"用户id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"user name", "zh_CN":"用户名"}
  Username *string `json:"username,omitempty" xml:"username,omitempty" require:"true"`
  // {"en":"createTime", "zh_CN":"创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"updateTime", "zh_CN":"更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s GetHarborUserUser) String() string {
  return tea.Prettify(s)
}

func (s GetHarborUserUser) GoString() string {
  return s.String()
}

func (s *GetHarborUserUser) SetId(v int64) *GetHarborUserUser {
  s.Id = &v
  return s
}

func (s *GetHarborUserUser) SetUsername(v string) *GetHarborUserUser {
  s.Username = &v
  return s
}

func (s *GetHarborUserUser) SetCreateTime(v int64) *GetHarborUserUser {
  s.CreateTime = &v
  return s
}

func (s *GetHarborUserUser) SetUpdateTime(v int64) *GetHarborUserUser {
  s.UpdateTime = &v
  return s
}




