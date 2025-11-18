package edgehostname

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateEdgeHostnameForTerraformRequest struct {
  // {"en":"Edge-Hostname comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"geoFence, data range: [global,inside_china_mainland,exclude_china_mainland].","zh_CN":"加速区域限定。取值范围：[global,inside_china_mainland,exclude_china_mainland]。"}
  GeoFence *string `json:"geoFence,omitempty" xml:"geoFence,omitempty" require:"true"`
  // {"en":"region configuration","zh_CN":"区域配置列表"}
  RegionConfigs []*UpdateEdgeHostnameForTerraformRequestRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEdgeHostnameForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformRequest) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameForTerraformRequest) SetComment(v string) *UpdateEdgeHostnameForTerraformRequest {
  s.Comment = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequest) SetGeoFence(v string) *UpdateEdgeHostnameForTerraformRequest {
  s.GeoFence = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequest) SetRegionConfigs(v []*UpdateEdgeHostnameForTerraformRequestRegionConfigs) *UpdateEdgeHostnameForTerraformRequest {
  s.RegionConfigs = v
  return s
}

type UpdateEdgeHostnameForTerraformRequestRegionConfigs struct     {
  // {"en":"region id","zh_CN":"区域ID"}
  RegionId *int `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
  // {"en":"action type","zh_CN":"调度方式。","exampleValue":"deliver,redirect,reject"}
  ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"config value","zh_CN":"记录值。"}
  ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
  // {"en":"ip protocol.","zh_CN":"ip协议。"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
  // {"en":"ttl","zh_CN":"生存时间。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty"`
  // {"en":"weight","zh_CN":"权重。"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s UpdateEdgeHostnameForTerraformRequestRegionConfigs) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformRequestRegionConfigs) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameForTerraformRequestRegionConfigs) SetRegionId(v int) *UpdateEdgeHostnameForTerraformRequestRegionConfigs {
  s.RegionId = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequestRegionConfigs) SetActionType(v string) *UpdateEdgeHostnameForTerraformRequestRegionConfigs {
  s.ActionType = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequestRegionConfigs) SetConfigValue(v string) *UpdateEdgeHostnameForTerraformRequestRegionConfigs {
  s.ConfigValue = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequestRegionConfigs) SetIpProtocol(v string) *UpdateEdgeHostnameForTerraformRequestRegionConfigs {
  s.IpProtocol = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequestRegionConfigs) SetTtl(v int) *UpdateEdgeHostnameForTerraformRequestRegionConfigs {
  s.Ttl = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformRequestRegionConfigs) SetWeight(v int) *UpdateEdgeHostnameForTerraformRequestRegionConfigs {
  s.Weight = &v
  return s
}

type UpdateEdgeHostnameForTerraformRequestHeader struct {
}

func (s UpdateEdgeHostnameForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformRequestHeader) GoString() string {
  return s.String()
}

type UpdateEdgeHostnameForTerraformPaths struct {
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s UpdateEdgeHostnameForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformPaths) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameForTerraformPaths) SetEdgeHostname(v string) *UpdateEdgeHostnameForTerraformPaths {
  s.EdgeHostname = &v
  return s
}

type UpdateEdgeHostnameForTerraformParameters struct {
}

func (s UpdateEdgeHostnameForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformParameters) GoString() string {
  return s.String()
}

type UpdateEdgeHostnameForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateEdgeHostnameForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformResponse) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameForTerraformResponse) SetCode(v string) *UpdateEdgeHostnameForTerraformResponse {
  s.Code = &v
  return s
}

func (s *UpdateEdgeHostnameForTerraformResponse) SetMessage(v string) *UpdateEdgeHostnameForTerraformResponse {
  s.Message = &v
  return s
}

type UpdateEdgeHostnameForTerraformResponseHeader struct {
}

func (s UpdateEdgeHostnameForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameForTerraformResponseHeader) GoString() string {
  return s.String()
}




type ListEdgeHostnamesRequest struct {
}

func (s ListEdgeHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesRequest) GoString() string {
  return s.String()
}

type ListEdgeHostnamesRequestHeader struct {
}

func (s ListEdgeHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesRequestHeader) GoString() string {
  return s.String()
}

type ListEdgeHostnamesPaths struct {
}

func (s ListEdgeHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesPaths) GoString() string {
  return s.String()
}

type ListEdgeHostnamesParameters struct {
  // {"en":"edgeHostnames.","zh_CN":"调度域名列表。多个以英文逗号分割"}
  EdgeHostnames *string `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty"`
  // {"en":"hostnames.","zh_CN":"域名列表。多个以英文逗号分割"}
  Hostnames *string `json:"hostnames,omitempty" xml:"hostnames,omitempty"`
  // {"en":"comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"The value can be 'inactive', or 'active' to filter the results based on dns service status.","zh_CN":"DNS服务状态。","exampleValue":"inactive, active"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty"`
  // {"en":"The value can be 'pending', or 'deploying', or 'success', or 'fail' to filter the results based on deploy status.remark:split with comma","zh_CN":"部署状态。多个以英文逗号分割"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
  // {"en":"fuzzy query.only used for edgeHostnames,hostnames and comment","zh_CN":"是否模糊匹配。只对edgeHostnames、hostnames与comment有效"}
  IsLike *string `json:"isLike,omitempty" xml:"isLike,omitempty"`
  // {"en":"The value can be '0', or '1' to filter the results based on allow china cdn .","zh_CN":"是否允许中国大陆加速。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty"`
  // {"defaultValue":"0","en":"Indicates the first item to return.","zh_CN":"查询起始位置。取值范围: >= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"defaultValue":"100","en":"Maximum number of properties to return.  Range: <= 200","zh_CN":"每次查询的最大条数。取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"defaultValue":"desc","en":"Order of properties to return.","zh_CN":"返回结果的顺序。默认按最后更新时间降序。","exampleValue":"asc,desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"defaultValue":"lastUpdateTime","en":"Returns results in sorted order.","zh_CN":"返回结果的排序依据。","exampleValue":"creationTime,lastUpdateTime,edgeHostname"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s ListEdgeHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesParameters) GoString() string {
  return s.String()
}

func (s *ListEdgeHostnamesParameters) SetEdgeHostnames(v string) *ListEdgeHostnamesParameters {
  s.EdgeHostnames = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetHostnames(v string) *ListEdgeHostnamesParameters {
  s.Hostnames = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetComment(v string) *ListEdgeHostnamesParameters {
  s.Comment = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetDnsServiceStatus(v string) *ListEdgeHostnamesParameters {
  s.DnsServiceStatus = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetDeployStatus(v string) *ListEdgeHostnamesParameters {
  s.DeployStatus = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetIsLike(v string) *ListEdgeHostnamesParameters {
  s.IsLike = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetAllowChinaCdn(v string) *ListEdgeHostnamesParameters {
  s.AllowChinaCdn = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetOffset(v int) *ListEdgeHostnamesParameters {
  s.Offset = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetLimit(v int) *ListEdgeHostnamesParameters {
  s.Limit = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetSortOrder(v string) *ListEdgeHostnamesParameters {
  s.SortOrder = &v
  return s
}

func (s *ListEdgeHostnamesParameters) SetSortBy(v string) *ListEdgeHostnamesParameters {
  s.SortBy = &v
  return s
}

type ListEdgeHostnamesResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *ListEdgeHostnamesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListEdgeHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesResponse) GoString() string {
  return s.String()
}

func (s *ListEdgeHostnamesResponse) SetCode(v string) *ListEdgeHostnamesResponse {
  s.Code = &v
  return s
}

func (s *ListEdgeHostnamesResponse) SetMessage(v string) *ListEdgeHostnamesResponse {
  s.Message = &v
  return s
}

func (s *ListEdgeHostnamesResponse) SetData(v *ListEdgeHostnamesResponseData) *ListEdgeHostnamesResponse {
  s.Data = v
  return s
}

type ListEdgeHostnamesResponseData struct {
  // {"en":"Number of properties.","zh_CN":"调度域名数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of properties.","zh_CN":"项目列表。"}
  EdgeHostnames []*ListEdgeHostnamesResponseDataEdgeHostnames `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListEdgeHostnamesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesResponseData) GoString() string {
  return s.String()
}

func (s *ListEdgeHostnamesResponseData) SetCount(v int) *ListEdgeHostnamesResponseData {
  s.Count = &v
  return s
}

func (s *ListEdgeHostnamesResponseData) SetEdgeHostnames(v []*ListEdgeHostnamesResponseDataEdgeHostnames) *ListEdgeHostnamesResponseData {
  s.EdgeHostnames = v
  return s
}

type ListEdgeHostnamesResponseDataEdgeHostnames struct     {
  // {"en":"Edge-Hostname ID.","zh_CN":"调度域名标识"}
  EdgeHostnameId *int64 `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
  // {"en":"Edge-Hostname.","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en":"dns service status.","zh_CN":"DNS服务状态。","exampleValue":"inactive, active"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"deploy status.","zh_CN":"部署状态","exampleValue":"pending, deploying, success, fail"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
  // {"en":"allow china cdn. remark: 0 no, 1 yes.","zh_CN":"是否允许中国大陆加速。备注：0 否，1 是。","exampleValue":"0，1"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" require:"true"`
  // {"en":"Edge-Hostname comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the Edge-Hostname was created.","zh_CN":"RFC 3339格式的日期，表示创建调度域名的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the Edge-Hostname was last updated.","zh_CN":"RFC 3339格式的日期，表示调度域名的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*ListEdgeHostnamesResponseDataEdgeHostnamesHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListEdgeHostnamesResponseDataEdgeHostnames) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesResponseDataEdgeHostnames) GoString() string {
  return s.String()
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetEdgeHostnameId(v int64) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.EdgeHostnameId = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetEdgeHostname(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.EdgeHostname = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetDnsServiceStatus(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.DnsServiceStatus = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetDeployStatus(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.DeployStatus = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetAllowChinaCdn(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.AllowChinaCdn = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetComment(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.Comment = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetCreationTime(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.CreationTime = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetLastUpdateTime(v string) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.LastUpdateTime = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnames) SetHostnames(v []*ListEdgeHostnamesResponseDataEdgeHostnamesHostnames) *ListEdgeHostnamesResponseDataEdgeHostnames {
  s.Hostnames = v
  return s
}

type ListEdgeHostnamesResponseDataEdgeHostnamesHostnames struct     {
  // {"en":"name of the domain which use this edge-hostname.","zh_CN":"域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"the deploy target of this hostname.","zh_CN":"部署环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
}

func (s ListEdgeHostnamesResponseDataEdgeHostnamesHostnames) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesResponseDataEdgeHostnamesHostnames) GoString() string {
  return s.String()
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnamesHostnames) SetHostname(v string) *ListEdgeHostnamesResponseDataEdgeHostnamesHostnames {
  s.Hostname = &v
  return s
}

func (s *ListEdgeHostnamesResponseDataEdgeHostnamesHostnames) SetTarget(v string) *ListEdgeHostnamesResponseDataEdgeHostnamesHostnames {
  s.Target = &v
  return s
}

type ListEdgeHostnamesResponseHeader struct {
}

func (s ListEdgeHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListEdgeHostnamesResponseHeader) GoString() string {
  return s.String()
}




type DeployEdgeHostnameForTerraformRequest struct {
}

func (s DeployEdgeHostnameForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameForTerraformRequest) GoString() string {
  return s.String()
}

type DeployEdgeHostnameForTerraformRequestHeader struct {
}

func (s DeployEdgeHostnameForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameForTerraformRequestHeader) GoString() string {
  return s.String()
}

type DeployEdgeHostnameForTerraformPaths struct {
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DeployEdgeHostnameForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameForTerraformPaths) GoString() string {
  return s.String()
}

func (s *DeployEdgeHostnameForTerraformPaths) SetEdgeHostname(v string) *DeployEdgeHostnameForTerraformPaths {
  s.EdgeHostname = &v
  return s
}

type DeployEdgeHostnameForTerraformParameters struct {
}

func (s DeployEdgeHostnameForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameForTerraformParameters) GoString() string {
  return s.String()
}

type DeployEdgeHostnameForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeployEdgeHostnameForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameForTerraformResponse) GoString() string {
  return s.String()
}

func (s *DeployEdgeHostnameForTerraformResponse) SetCode(v string) *DeployEdgeHostnameForTerraformResponse {
  s.Code = &v
  return s
}

func (s *DeployEdgeHostnameForTerraformResponse) SetMessage(v string) *DeployEdgeHostnameForTerraformResponse {
  s.Message = &v
  return s
}

type DeployEdgeHostnameForTerraformResponseHeader struct {
}

func (s DeployEdgeHostnameForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameForTerraformResponseHeader) GoString() string {
  return s.String()
}




type DeleteEdgeHostnameRequest struct {
}

func (s DeleteEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameRequest) GoString() string {
  return s.String()
}

type DeleteEdgeHostnameRequestHeader struct {
}

func (s DeleteEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type DeleteEdgeHostnamePaths struct {
  // {"en":"edgeHostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DeleteEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *DeleteEdgeHostnamePaths) SetEdgeHostname(v string) *DeleteEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type DeleteEdgeHostnameParameters struct {
}

func (s DeleteEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameParameters) GoString() string {
  return s.String()
}

type DeleteEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *DeleteEdgeHostnameResponse) SetCode(v string) *DeleteEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *DeleteEdgeHostnameResponse) SetMessage(v string) *DeleteEdgeHostnameResponse {
  s.Message = &v
  return s
}

type DeleteEdgeHostnameResponseHeader struct {
}

func (s DeleteEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type DeployEdgeHostnameDNSRequest struct {
}

func (s DeployEdgeHostnameDNSRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameDNSRequest) GoString() string {
  return s.String()
}

type DeployEdgeHostnameDNSRequestHeader struct {
}

func (s DeployEdgeHostnameDNSRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameDNSRequestHeader) GoString() string {
  return s.String()
}

type DeployEdgeHostnameDNSPaths struct {
  // {"en":"edgeHostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DeployEdgeHostnameDNSPaths) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameDNSPaths) GoString() string {
  return s.String()
}

func (s *DeployEdgeHostnameDNSPaths) SetEdgeHostname(v string) *DeployEdgeHostnameDNSPaths {
  s.EdgeHostname = &v
  return s
}

type DeployEdgeHostnameDNSParameters struct {
}

func (s DeployEdgeHostnameDNSParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameDNSParameters) GoString() string {
  return s.String()
}

type DeployEdgeHostnameDNSResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeployEdgeHostnameDNSResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameDNSResponse) GoString() string {
  return s.String()
}

func (s *DeployEdgeHostnameDNSResponse) SetCode(v string) *DeployEdgeHostnameDNSResponse {
  s.Code = &v
  return s
}

func (s *DeployEdgeHostnameDNSResponse) SetMessage(v string) *DeployEdgeHostnameDNSResponse {
  s.Message = &v
  return s
}

type DeployEdgeHostnameDNSResponseHeader struct {
}

func (s DeployEdgeHostnameDNSResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployEdgeHostnameDNSResponseHeader) GoString() string {
  return s.String()
}




type GetEdgeHostnameRequest struct {
}

func (s GetEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameRequest) GoString() string {
  return s.String()
}

type GetEdgeHostnameRequestHeader struct {
}

func (s GetEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeHostnamePaths struct {
  // {"en":"edgeHostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s GetEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnamePaths) SetEdgeHostname(v string) *GetEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type GetEdgeHostnameParameters struct {
}

func (s GetEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameParameters) GoString() string {
  return s.String()
}

type GetEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetEdgeHostnameResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameResponse) SetCode(v string) *GetEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *GetEdgeHostnameResponse) SetMessage(v string) *GetEdgeHostnameResponse {
  s.Message = &v
  return s
}

func (s *GetEdgeHostnameResponse) SetData(v *GetEdgeHostnameResponseData) *GetEdgeHostnameResponse {
  s.Data = v
  return s
}

type GetEdgeHostnameResponseData struct {
  // {"en":"Edge-Hostname ID","zh_CN":"调度域名ID"}
  EdgeHostnameId *int64 `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
  // {"en":"Edge-Hostname Name","zh_CN":"调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en":"Edge-Hostname comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"dns service status.","zh_CN":"DNS服务状态","exampleValue":"inactive, active"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"deploy status.","zh_CN":"部署状态","exampleValue":"pending,deploying,success,fail"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
  // {"en":"allow china cdn. remark: 0 no, 1 yes.","zh_CN":"是否允许中国大陆加速。备注：0 否，1 是。","exampleValue":"0,1"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" require:"true"`
  // {"en":"gdpr compliant","zh_CN":"遵循GDPR","exampleValue":"0,1,2"}
  GdprCompliant *string `json:"gdprCompliant,omitempty" xml:"gdprCompliant,omitempty" require:"true"`
  // {"en":"geoFence","zh_CN":"加速区域限定","exampleValue":"global,inside_china_mainland,exclude_china_mainland]。"}
  GeoFence *string `json:"geoFence,omitempty" xml:"geoFence,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the edge-hostname was created.","zh_CN":"RFC 3339格式的日期，表示创建edge-hostname的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the edge-hostname was last updated.","zh_CN":"RFC 3339格式的日期，表示edge-hostname的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames","zh_CN":"关联的加速域名"}
  Hostnames []*GetEdgeHostnameResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"region configuration","zh_CN":"区域配置列表"}
  RegionConfigs []*GetEdgeHostnameResponseDataRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeHostnameResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameResponseData) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameResponseData) SetEdgeHostnameId(v int64) *GetEdgeHostnameResponseData {
  s.EdgeHostnameId = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetEdgeHostname(v string) *GetEdgeHostnameResponseData {
  s.EdgeHostname = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetComment(v string) *GetEdgeHostnameResponseData {
  s.Comment = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetDnsServiceStatus(v string) *GetEdgeHostnameResponseData {
  s.DnsServiceStatus = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetDeployStatus(v string) *GetEdgeHostnameResponseData {
  s.DeployStatus = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetAllowChinaCdn(v string) *GetEdgeHostnameResponseData {
  s.AllowChinaCdn = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetGdprCompliant(v string) *GetEdgeHostnameResponseData {
  s.GdprCompliant = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetGeoFence(v string) *GetEdgeHostnameResponseData {
  s.GeoFence = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetCreationTime(v string) *GetEdgeHostnameResponseData {
  s.CreationTime = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetLastUpdateTime(v string) *GetEdgeHostnameResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *GetEdgeHostnameResponseData) SetHostnames(v []*GetEdgeHostnameResponseDataHostnames) *GetEdgeHostnameResponseData {
  s.Hostnames = v
  return s
}

func (s *GetEdgeHostnameResponseData) SetRegionConfigs(v []*GetEdgeHostnameResponseDataRegionConfigs) *GetEdgeHostnameResponseData {
  s.RegionConfigs = v
  return s
}

type GetEdgeHostnameResponseDataHostnames struct     {
  // {"en":"hostname","zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"deploy target","zh_CN":"部署环境"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"property id","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"property version","zh_CN":"项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"property name","zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
}

func (s GetEdgeHostnameResponseDataHostnames) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameResponseDataHostnames) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameResponseDataHostnames) SetHostname(v string) *GetEdgeHostnameResponseDataHostnames {
  s.Hostname = &v
  return s
}

func (s *GetEdgeHostnameResponseDataHostnames) SetTarget(v string) *GetEdgeHostnameResponseDataHostnames {
  s.Target = &v
  return s
}

func (s *GetEdgeHostnameResponseDataHostnames) SetPropertyId(v int) *GetEdgeHostnameResponseDataHostnames {
  s.PropertyId = &v
  return s
}

func (s *GetEdgeHostnameResponseDataHostnames) SetPropertyVersion(v int) *GetEdgeHostnameResponseDataHostnames {
  s.PropertyVersion = &v
  return s
}

func (s *GetEdgeHostnameResponseDataHostnames) SetPropertyName(v string) *GetEdgeHostnameResponseDataHostnames {
  s.PropertyName = &v
  return s
}

type GetEdgeHostnameResponseDataRegionConfigs struct     {
  // {"en":"region id","zh_CN":"区域ID"}
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
  // {"en":"action type","zh_CN":"调度方式。"}
  ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"config type","zh_CN":"记录类型。"}
  ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty" require:"true"`
  // {"en":"config value","zh_CN":"记录值。"}
  ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty" require:"true"`
  // {"en":"ip protocol.","zh_CN":"ip协议。"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty" require:"true"`
  // {"en":"ttl","zh_CN":"生存时间。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重。"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s GetEdgeHostnameResponseDataRegionConfigs) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameResponseDataRegionConfigs) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetRegionId(v string) *GetEdgeHostnameResponseDataRegionConfigs {
  s.RegionId = &v
  return s
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetActionType(v string) *GetEdgeHostnameResponseDataRegionConfigs {
  s.ActionType = &v
  return s
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetConfigType(v string) *GetEdgeHostnameResponseDataRegionConfigs {
  s.ConfigType = &v
  return s
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetConfigValue(v string) *GetEdgeHostnameResponseDataRegionConfigs {
  s.ConfigValue = &v
  return s
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetIpProtocol(v string) *GetEdgeHostnameResponseDataRegionConfigs {
  s.IpProtocol = &v
  return s
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetTtl(v int) *GetEdgeHostnameResponseDataRegionConfigs {
  s.Ttl = &v
  return s
}

func (s *GetEdgeHostnameResponseDataRegionConfigs) SetWeight(v int) *GetEdgeHostnameResponseDataRegionConfigs {
  s.Weight = &v
  return s
}

type GetEdgeHostnameResponseHeader struct {
}

func (s GetEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type UndeployEdgeHostnameDNSRequest struct {
}

func (s UndeployEdgeHostnameDNSRequest) String() string {
  return tea.Prettify(s)
}

func (s UndeployEdgeHostnameDNSRequest) GoString() string {
  return s.String()
}

type UndeployEdgeHostnameDNSRequestHeader struct {
}

func (s UndeployEdgeHostnameDNSRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UndeployEdgeHostnameDNSRequestHeader) GoString() string {
  return s.String()
}

type UndeployEdgeHostnameDNSPaths struct {
  // {"en":"edgeHostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s UndeployEdgeHostnameDNSPaths) String() string {
  return tea.Prettify(s)
}

func (s UndeployEdgeHostnameDNSPaths) GoString() string {
  return s.String()
}

func (s *UndeployEdgeHostnameDNSPaths) SetEdgeHostname(v string) *UndeployEdgeHostnameDNSPaths {
  s.EdgeHostname = &v
  return s
}

type UndeployEdgeHostnameDNSParameters struct {
}

func (s UndeployEdgeHostnameDNSParameters) String() string {
  return tea.Prettify(s)
}

func (s UndeployEdgeHostnameDNSParameters) GoString() string {
  return s.String()
}

type UndeployEdgeHostnameDNSResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UndeployEdgeHostnameDNSResponse) String() string {
  return tea.Prettify(s)
}

func (s UndeployEdgeHostnameDNSResponse) GoString() string {
  return s.String()
}

func (s *UndeployEdgeHostnameDNSResponse) SetCode(v string) *UndeployEdgeHostnameDNSResponse {
  s.Code = &v
  return s
}

func (s *UndeployEdgeHostnameDNSResponse) SetMessage(v string) *UndeployEdgeHostnameDNSResponse {
  s.Message = &v
  return s
}

type UndeployEdgeHostnameDNSResponseHeader struct {
}

func (s UndeployEdgeHostnameDNSResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UndeployEdgeHostnameDNSResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeHostnameForTerraformRequest struct {
}

func (s QueryEdgeHostnameForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformRequest) GoString() string {
  return s.String()
}

type QueryEdgeHostnameForTerraformRequestHeader struct {
}

func (s QueryEdgeHostnameForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeHostnameForTerraformPaths struct {
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s QueryEdgeHostnameForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformPaths) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameForTerraformPaths) SetEdgeHostname(v string) *QueryEdgeHostnameForTerraformPaths {
  s.EdgeHostname = &v
  return s
}

type QueryEdgeHostnameForTerraformParameters struct {
}

func (s QueryEdgeHostnameForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformParameters) GoString() string {
  return s.String()
}

type QueryEdgeHostnameForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryEdgeHostnameForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeHostnameForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameForTerraformResponse) SetCode(v string) *QueryEdgeHostnameForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponse) SetMessage(v string) *QueryEdgeHostnameForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponse) SetData(v *QueryEdgeHostnameForTerraformResponseData) *QueryEdgeHostnameForTerraformResponse {
  s.Data = v
  return s
}

type QueryEdgeHostnameForTerraformResponseData struct {
  // {"en":"Edge-Hostname ID","zh_CN":"调度域名ID"}
  EdgeHostnameId *int `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
  // {"en":"Edge-Hostname Name","zh_CN":"调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en":"Edge-Hostname comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"dns service status.data range:[inactive, active]","zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"deploy status.data range:[pending, deploying, success, fail]","zh_CN":"部署状态。取值范围：[pending, deploying, success, fail]"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
  // {"en":"allow china cdn.,data range:[0,1].remark: 0 no, 1 yes.","zh_CN":"是否允许中国大陆加速。取值范围：[0,1]。备注：0 否，1 是。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" require:"true"`
  // {"en":"gdpr compliant,data range:[0,1,2].","zh_CN":"遵循GDPR,取值范围：[0,1,2]。"}
  GdprCompliant *string `json:"gdprCompliant,omitempty" xml:"gdprCompliant,omitempty" require:"true"`
  // {"en":"geoFence, data range: [global,inside_china_mainland,exclude_china_mainland].","zh_CN":"加速区域限定。取值范围：[global,inside_china_mainland,exclude_china_mainland]。"}
  GeoFence *string `json:"geoFence,omitempty" xml:"geoFence,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the edge-hostname was created.","zh_CN":"RFC 3339格式的日期，表示创建edge-hostname的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the edge-hostname was last updated.","zh_CN":"RFC 3339格式的日期，表示edge-hostname的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames","zh_CN":"关联的加速域名"}
  Hostnames []*QueryEdgeHostnameForTerraformResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"region configuration","zh_CN":"区域配置列表"}
  RegionConfigs []*QueryEdgeHostnameForTerraformResponseDataRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeHostnameForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetEdgeHostnameId(v int) *QueryEdgeHostnameForTerraformResponseData {
  s.EdgeHostnameId = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetEdgeHostname(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.EdgeHostname = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetComment(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.Comment = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetDnsServiceStatus(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetDeployStatus(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.DeployStatus = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetAllowChinaCdn(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.AllowChinaCdn = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetGdprCompliant(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.GdprCompliant = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetGeoFence(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.GeoFence = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetCreationTime(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.CreationTime = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetLastUpdateTime(v string) *QueryEdgeHostnameForTerraformResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetHostnames(v []*QueryEdgeHostnameForTerraformResponseDataHostnames) *QueryEdgeHostnameForTerraformResponseData {
  s.Hostnames = v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseData) SetRegionConfigs(v []*QueryEdgeHostnameForTerraformResponseDataRegionConfigs) *QueryEdgeHostnameForTerraformResponseData {
  s.RegionConfigs = v
  return s
}

type QueryEdgeHostnameForTerraformResponseDataHostnames struct     {
  // {"en":"hostname","zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"deploy target","zh_CN":"部署环境"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"property id","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"property version","zh_CN":"项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"property name","zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
}

func (s QueryEdgeHostnameForTerraformResponseDataHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformResponseDataHostnames) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameForTerraformResponseDataHostnames) SetHostname(v string) *QueryEdgeHostnameForTerraformResponseDataHostnames {
  s.Hostname = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataHostnames) SetTarget(v string) *QueryEdgeHostnameForTerraformResponseDataHostnames {
  s.Target = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataHostnames) SetPropertyId(v int) *QueryEdgeHostnameForTerraformResponseDataHostnames {
  s.PropertyId = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataHostnames) SetPropertyVersion(v int) *QueryEdgeHostnameForTerraformResponseDataHostnames {
  s.PropertyVersion = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataHostnames) SetPropertyName(v string) *QueryEdgeHostnameForTerraformResponseDataHostnames {
  s.PropertyName = &v
  return s
}

type QueryEdgeHostnameForTerraformResponseDataRegionConfigs struct     {
  // {"en":"region id","zh_CN":"区域ID"}
  RegionId *int `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
  // {"en":"action type","zh_CN":"调度方式。"}
  ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"config type","zh_CN":"记录类型。"}
  ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty" require:"true"`
  // {"en":"config value","zh_CN":"记录值。"}
  ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty" require:"true"`
  // {"en":"ip protocol.","zh_CN":"ip协议。"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty" require:"true"`
  // {"en":"ttl","zh_CN":"生存时间。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重。"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s QueryEdgeHostnameForTerraformResponseDataRegionConfigs) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformResponseDataRegionConfigs) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetRegionId(v int) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.RegionId = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetActionType(v string) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.ActionType = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetConfigType(v string) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.ConfigType = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetConfigValue(v string) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.ConfigValue = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetIpProtocol(v string) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.IpProtocol = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetTtl(v int) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.Ttl = &v
  return s
}

func (s *QueryEdgeHostnameForTerraformResponseDataRegionConfigs) SetWeight(v int) *QueryEdgeHostnameForTerraformResponseDataRegionConfigs {
  s.Weight = &v
  return s
}

type QueryEdgeHostnameForTerraformResponseHeader struct {
}

func (s QueryEdgeHostnameForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameForTerraformResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeHostnamesForTerraformRequest struct {
}

func (s QueryEdgeHostnamesForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformRequest) GoString() string {
  return s.String()
}

type QueryEdgeHostnamesForTerraformRequestHeader struct {
}

func (s QueryEdgeHostnamesForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeHostnamesForTerraformPaths struct {
}

func (s QueryEdgeHostnamesForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformPaths) GoString() string {
  return s.String()
}

type QueryEdgeHostnamesForTerraformParameters struct {
  // {"en":"edgeHostnames.","zh_CN":"调度域名列表。多个以英文逗号分割"}
  EdgeHostnames *string `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty"`
  // {"en":"hostnames.","zh_CN":"域名列表。多个以英文逗号分割"}
  Hostnames *string `json:"hostnames,omitempty" xml:"hostnames,omitempty"`
  // {"en":"comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"The value can be 'inactive', or 'active' to filter the results based on dns service status.","zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty"`
  // {"en":"The value can be 'pending', or 'deploying', or 'success', or 'fail' to filter the results based on deploy status.remark:split with comma","zh_CN":"部署状态。多个以英文逗号分割"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
  // {"en":"fuzzy query.only used for edgeHostnames,hostnames and comment","zh_CN":"是否模糊匹配。只对edgeHostnames、hostnames与comment有效"}
  IsLike *string `json:"isLike,omitempty" xml:"isLike,omitempty"`
  // {"en":"The value can be '0', or '1' to filter the results based on allow china cdn .","zh_CN":"是否允许中国大陆加速。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty"`
  // {"en":"Indicates the first item to return. The default is '0'.","zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200","zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Order of properties to return. Enum: asc,desc Default: desc","zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en":"Returns results in sorted order. Enum: creationTime,lastUpdateTime,edgeHostname Default: lastUpdateTime","zh_CN":"返回结果的排序依据。取值范围: creationTime,lastUpdateTime,edgeHostname 默认值: lastUpdateTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s QueryEdgeHostnamesForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformParameters) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetEdgeHostnames(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.EdgeHostnames = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetHostnames(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.Hostnames = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetComment(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.Comment = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetDnsServiceStatus(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetDeployStatus(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.DeployStatus = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetIsLike(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.IsLike = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetAllowChinaCdn(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.AllowChinaCdn = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetOffset(v int) *QueryEdgeHostnamesForTerraformParameters {
  s.Offset = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetLimit(v int) *QueryEdgeHostnamesForTerraformParameters {
  s.Limit = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetSortOrder(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.SortOrder = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformParameters) SetSortBy(v string) *QueryEdgeHostnamesForTerraformParameters {
  s.SortBy = &v
  return s
}

type QueryEdgeHostnamesForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryEdgeHostnamesForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeHostnamesForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesForTerraformResponse) SetCode(v string) *QueryEdgeHostnamesForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponse) SetMessage(v string) *QueryEdgeHostnamesForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponse) SetData(v *QueryEdgeHostnamesForTerraformResponseData) *QueryEdgeHostnamesForTerraformResponse {
  s.Data = v
  return s
}

type QueryEdgeHostnamesForTerraformResponseData struct {
  // {"en":"Number of edge-hostnames.","zh_CN":"调度域名数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of edge-hostnames.","zh_CN":"调度域名列表。"}
  EdgeHostnames []*QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeHostnamesForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesForTerraformResponseData) SetCount(v int) *QueryEdgeHostnamesForTerraformResponseData {
  s.Count = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseData) SetEdgeHostnames(v []*QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) *QueryEdgeHostnamesForTerraformResponseData {
  s.EdgeHostnames = v
  return s
}

type QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames struct     {
  // {"en":"Edge-Hostname ID.","zh_CN":"调度域名标识"}
  EdgeHostnameId *int `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
  // {"en":"Edge-Hostname.","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en":"dns service status.data range:[inactive, active]","zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"deploy status.data range:[pending, deploying, success, fail].","zh_CN":"部署状态。取值范围：[pending, deploying, success, fail]"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
  // {"en":"allow china cdn.,data range:[0,1].remark: 0 no, 1 yes.","zh_CN":"是否允许中国大陆加速。取值范围：[0,1]。备注：0 否，1 是。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" require:"true"`
  // {"en":"Edge-Hostname comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the Edge-Hostname was created.","zh_CN":"RFC 3339格式的日期，表示创建调度域名的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the Edge-Hostname was last updated.","zh_CN":"RFC 3339格式的日期，表示调度域名的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetEdgeHostnameId(v int) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.EdgeHostnameId = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetEdgeHostname(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.EdgeHostname = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetDnsServiceStatus(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetDeployStatus(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.DeployStatus = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetAllowChinaCdn(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.AllowChinaCdn = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetComment(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.Comment = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetCreationTime(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.CreationTime = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetLastUpdateTime(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetHostnames(v []*QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
  s.Hostnames = v
  return s
}

type QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames struct     {
  // {"en":"name of the domain which use this edge-hostname.","zh_CN":"域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"the deploy target of this hostname.","zh_CN":"部署环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
}

func (s QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames) SetHostname(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames {
  s.Hostname = &v
  return s
}

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames) SetTarget(v string) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames {
  s.Target = &v
  return s
}

type QueryEdgeHostnamesForTerraformResponseHeader struct {
}

func (s QueryEdgeHostnamesForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesForTerraformResponseHeader) GoString() string {
  return s.String()
}




type DeleteEdgeHostnameForTerraformRequest struct {
}

func (s DeleteEdgeHostnameForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameForTerraformRequest) GoString() string {
  return s.String()
}

type DeleteEdgeHostnameForTerraformRequestHeader struct {
}

func (s DeleteEdgeHostnameForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameForTerraformRequestHeader) GoString() string {
  return s.String()
}

type DeleteEdgeHostnameForTerraformPaths struct {
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DeleteEdgeHostnameForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameForTerraformPaths) GoString() string {
  return s.String()
}

func (s *DeleteEdgeHostnameForTerraformPaths) SetEdgeHostname(v string) *DeleteEdgeHostnameForTerraformPaths {
  s.EdgeHostname = &v
  return s
}

type DeleteEdgeHostnameForTerraformParameters struct {
}

func (s DeleteEdgeHostnameForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameForTerraformParameters) GoString() string {
  return s.String()
}

type DeleteEdgeHostnameForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteEdgeHostnameForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameForTerraformResponse) GoString() string {
  return s.String()
}

func (s *DeleteEdgeHostnameForTerraformResponse) SetCode(v string) *DeleteEdgeHostnameForTerraformResponse {
  s.Code = &v
  return s
}

func (s *DeleteEdgeHostnameForTerraformResponse) SetMessage(v string) *DeleteEdgeHostnameForTerraformResponse {
  s.Message = &v
  return s
}

type DeleteEdgeHostnameForTerraformResponseHeader struct {
}

func (s DeleteEdgeHostnameForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameForTerraformResponseHeader) GoString() string {
  return s.String()
}




type UpdateEdgeHostnameRequest struct {
  // {"en":"Edge-Hostname comment.","zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"geoFence","zh_CN":"加速区域限定","exampleValue":"global,inside_china_mainland,exclude_china_mainland"}
  GeoFence *string `json:"geoFence,omitempty" xml:"geoFence,omitempty"`
  // {"en":"region configuration","zh_CN":"区域配置列表"}
  RegionConfigs []*UpdateEdgeHostnameRequestRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameRequest) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameRequest) SetComment(v string) *UpdateEdgeHostnameRequest {
  s.Comment = &v
  return s
}

func (s *UpdateEdgeHostnameRequest) SetGeoFence(v string) *UpdateEdgeHostnameRequest {
  s.GeoFence = &v
  return s
}

func (s *UpdateEdgeHostnameRequest) SetRegionConfigs(v []*UpdateEdgeHostnameRequestRegionConfigs) *UpdateEdgeHostnameRequest {
  s.RegionConfigs = v
  return s
}

type UpdateEdgeHostnameRequestRegionConfigs struct     {
  // {"en":"region id","zh_CN":"区域ID"}
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
  // {"en":"action type","zh_CN":"调度方式。","exampleValue":"deliver,redirect,reject"}
  ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"config value","zh_CN":"记录值。"}
  ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
  // {"en":"ip protocol.","zh_CN":"ip协议。"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
  // {"en":"ttl","zh_CN":"生存时间。"}
  Ttl *string `json:"ttl,omitempty" xml:"ttl,omitempty"`
  // {"en":"weight","zh_CN":"权重。"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s UpdateEdgeHostnameRequestRegionConfigs) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameRequestRegionConfigs) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameRequestRegionConfigs) SetRegionId(v string) *UpdateEdgeHostnameRequestRegionConfigs {
  s.RegionId = &v
  return s
}

func (s *UpdateEdgeHostnameRequestRegionConfigs) SetActionType(v string) *UpdateEdgeHostnameRequestRegionConfigs {
  s.ActionType = &v
  return s
}

func (s *UpdateEdgeHostnameRequestRegionConfigs) SetConfigValue(v string) *UpdateEdgeHostnameRequestRegionConfigs {
  s.ConfigValue = &v
  return s
}

func (s *UpdateEdgeHostnameRequestRegionConfigs) SetIpProtocol(v string) *UpdateEdgeHostnameRequestRegionConfigs {
  s.IpProtocol = &v
  return s
}

func (s *UpdateEdgeHostnameRequestRegionConfigs) SetTtl(v string) *UpdateEdgeHostnameRequestRegionConfigs {
  s.Ttl = &v
  return s
}

func (s *UpdateEdgeHostnameRequestRegionConfigs) SetWeight(v int) *UpdateEdgeHostnameRequestRegionConfigs {
  s.Weight = &v
  return s
}

type UpdateEdgeHostnameRequestHeader struct {
}

func (s UpdateEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type UpdateEdgeHostnamePaths struct {
  // {"en":"edgeHostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s UpdateEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnamePaths) SetEdgeHostname(v string) *UpdateEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type UpdateEdgeHostnameParameters struct {
}

func (s UpdateEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameParameters) GoString() string {
  return s.String()
}

type UpdateEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameResponse) SetCode(v string) *UpdateEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *UpdateEdgeHostnameResponse) SetMessage(v string) *UpdateEdgeHostnameResponse {
  s.Message = &v
  return s
}

type UpdateEdgeHostnameResponseHeader struct {
}

func (s UpdateEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




