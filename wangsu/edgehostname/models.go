package edgehostname

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryEdgeHostnamesRequest struct {
}

func (s QueryEdgeHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesRequest) GoString() string {
  return s.String()
}

type QueryEdgeHostnamesResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.", "zh_CN":"接口响应数据"}
  Data *QueryEdgeHostnamesQueryEdgeHostnamesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesResponse) SetCode(v string) *QueryEdgeHostnamesResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeHostnamesResponse) SetMessage(v string) *QueryEdgeHostnamesResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeHostnamesResponse) SetData(v *QueryEdgeHostnamesQueryEdgeHostnamesResponseData) *QueryEdgeHostnamesResponse {
  s.Data = v
  return s
}

type QueryEdgeHostnamesQueryEdgeHostnamesResponseData struct {
  // {"en":"Number of properties.", "zh_CN":"调度域名数量。"}
  Count *int32 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of properties.", "zh_CN":"项目列表。"}
  EdgeHostnames []*QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeHostnamesQueryEdgeHostnamesResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesQueryEdgeHostnamesResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseData) SetCount(v int32) *QueryEdgeHostnamesQueryEdgeHostnamesResponseData {
  s.Count = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseData) SetEdgeHostnames(v []*QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) *QueryEdgeHostnamesQueryEdgeHostnamesResponseData {
  s.EdgeHostnames = v
  return s
}

type QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames struct     {
  // {"en":"Edge-Hostname ID.", "zh_CN":"调度域名标识"}
  EdgeHostnameId *int64 `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
  // {"en":"Edge-Hostname.", "zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en":"dns service status.data range:[inactive, active]", "zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"deploy status.data range:[pending, deploying, success, fail].", "zh_CN":"部署状态。取值范围：[pending, deploying, success, fail]"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
  // {"en":"allow china cdn.,data range:[0,1].remark: 0 no, 1 yes.", "zh_CN":"是否允许中国大陆加速。取值范围：[0,1]。备注：0 否，1 是。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" require:"true"`
  // {"en":"Edge-Hostname comment.", "zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the Edge-Hostname was created.", "zh_CN":"RFC 3339格式的日期，表示创建调度域名的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the Edge-Hostname was last updated.", "zh_CN":"RFC 3339格式的日期，表示调度域名的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames.", "zh_CN":"域名列表"}
  Hostnames []*QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetEdgeHostnameId(v int64) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.EdgeHostnameId = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetEdgeHostname(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.EdgeHostname = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetDnsServiceStatus(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetDeployStatus(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.DeployStatus = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetAllowChinaCdn(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.AllowChinaCdn = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetComment(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.Comment = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetCreationTime(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.CreationTime = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetLastUpdateTime(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames) SetHostnames(v []*QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnames {
  s.Hostnames = v
  return s
}

type QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames struct     {
  // {"en":"name of the domain which use this edge-hostname.", "zh_CN":"域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"the deploy target of this hostname.", "zh_CN":"部署环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
}

func (s QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) SetHostname(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames {
  s.Hostname = &v
  return s
}

func (s *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) SetTarget(v string) *QueryEdgeHostnamesQueryEdgeHostnamesResponseDataEdgeHostnamesHostnames {
  s.Target = &v
  return s
}

type QueryEdgeHostnamesPaths struct {
}

func (s QueryEdgeHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesPaths) GoString() string {
  return s.String()
}

type QueryEdgeHostnamesParameters struct {
  // {"en":"edgeHostnames.", "zh_CN":"调度域名列表。多个以英文逗号分割"}
  EdgeHostnames *string `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty"`
  // {"en":"hostnames.", "zh_CN":"域名列表。多个以英文逗号分割"}
  Hostnames *string `json:"hostnames,omitempty" xml:"hostnames,omitempty"`
  // {"en":"comment.", "zh_CN":"调度域名描述	。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"The value can be 'inactive', or 'active' to filter the results based on dns service status.", "zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty"`
  // {"en":"The value can be 'pending', or 'deploying', or 'success', or 'fail' to filter the results based on deploy status.remark:split with comma", "zh_CN":"部署状态。多个以英文逗号分割"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
  // {"en":"fuzzy query.only used for edgeHostnames,hostnames and comment", "zh_CN":"是否模糊匹配。只对edgeHostnames、hostnames与comment有效"}
  IsLike *bool `json:"isLike,omitempty" xml:"isLike,omitempty"`
  // {"en":"The value can be '0', or '1' to filter the results based on allow china cdn .", "zh_CN":"是否允许中国大陆加速。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty"`
  // {"en":"Indicates the first item to return. The default is '0'.", "zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
  Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200", "zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
  Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Order of properties to return. Enum: asc,desc Default: desc", "zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en":"Returns results in sorted order. Enum: creationTime,lastUpdateTime,edgeHostname Default: lastUpdateTime", "zh_CN":"返回结果的排序依据。取值范围: creationTime,lastUpdateTime,edgeHostname 默认值: lastUpdateTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s QueryEdgeHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesParameters) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamesParameters) SetEdgeHostnames(v string) *QueryEdgeHostnamesParameters {
  s.EdgeHostnames = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetHostnames(v string) *QueryEdgeHostnamesParameters {
  s.Hostnames = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetComment(v string) *QueryEdgeHostnamesParameters {
  s.Comment = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetDnsServiceStatus(v string) *QueryEdgeHostnamesParameters {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetDeployStatus(v string) *QueryEdgeHostnamesParameters {
  s.DeployStatus = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetIsLike(v bool) *QueryEdgeHostnamesParameters {
  s.IsLike = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetAllowChinaCdn(v string) *QueryEdgeHostnamesParameters {
  s.AllowChinaCdn = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetOffset(v int32) *QueryEdgeHostnamesParameters {
  s.Offset = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetLimit(v int32) *QueryEdgeHostnamesParameters {
  s.Limit = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetSortOrder(v string) *QueryEdgeHostnamesParameters {
  s.SortOrder = &v
  return s
}

func (s *QueryEdgeHostnamesParameters) SetSortBy(v string) *QueryEdgeHostnamesParameters {
  s.SortBy = &v
  return s
}

type QueryEdgeHostnamesRequestHeader struct {
}

func (s QueryEdgeHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeHostnamesResponseHeader struct {
}

func (s QueryEdgeHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamesResponseHeader) GoString() string {
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

type DeleteEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
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

type DeleteEdgeHostnamePaths struct {
  // {"en":"edgeHostname", "zh_CN":"调度域名"}
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

type DeleteEdgeHostnameRequestHeader struct {
}

func (s DeleteEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type DeleteEdgeHostnameResponseHeader struct {
}

func (s DeleteEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type EnableEdgeHostnameRequest struct {
}

func (s EnableEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableEdgeHostnameRequest) GoString() string {
  return s.String()
}

type EnableEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EnableEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *EnableEdgeHostnameResponse) SetCode(v string) *EnableEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *EnableEdgeHostnameResponse) SetMessage(v string) *EnableEdgeHostnameResponse {
  s.Message = &v
  return s
}

type EnableEdgeHostnamePaths struct {
  // {"en":"edgeHostname", "zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s EnableEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s EnableEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *EnableEdgeHostnamePaths) SetEdgeHostname(v string) *EnableEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type EnableEdgeHostnameParameters struct {
}

func (s EnableEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableEdgeHostnameParameters) GoString() string {
  return s.String()
}

type EnableEdgeHostnameRequestHeader struct {
}

func (s EnableEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type EnableEdgeHostnameResponseHeader struct {
}

func (s EnableEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type DeployAkcdnEdgeHostnameRequest struct {
}

func (s DeployAkcdnEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployAkcdnEdgeHostnameRequest) GoString() string {
  return s.String()
}

type DeployAkcdnEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeployAkcdnEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployAkcdnEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *DeployAkcdnEdgeHostnameResponse) SetCode(v string) *DeployAkcdnEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *DeployAkcdnEdgeHostnameResponse) SetMessage(v string) *DeployAkcdnEdgeHostnameResponse {
  s.Message = &v
  return s
}

type DeployAkcdnEdgeHostnamePaths struct {
  // {"en":"edgeHostname", "zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DeployAkcdnEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s DeployAkcdnEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *DeployAkcdnEdgeHostnamePaths) SetEdgeHostname(v string) *DeployAkcdnEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type DeployAkcdnEdgeHostnameParameters struct {
}

func (s DeployAkcdnEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployAkcdnEdgeHostnameParameters) GoString() string {
  return s.String()
}

type DeployAkcdnEdgeHostnameRequestHeader struct {
}

func (s DeployAkcdnEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployAkcdnEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type DeployAkcdnEdgeHostnameResponseHeader struct {
}

func (s DeployAkcdnEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployAkcdnEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeHostnameRequest struct {
}

func (s QueryEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameRequest) GoString() string {
  return s.String()
}

type QueryEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.", "zh_CN":"接口响应数据"}
  Data *QueryEdgeHostnameQueryEdgeHostnameResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameResponse) SetCode(v string) *QueryEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeHostnameResponse) SetMessage(v string) *QueryEdgeHostnameResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeHostnameResponse) SetData(v *QueryEdgeHostnameQueryEdgeHostnameResponseData) *QueryEdgeHostnameResponse {
  s.Data = v
  return s
}

type QueryEdgeHostnameQueryEdgeHostnameResponseData struct {
  // {"en":"Edge-Hostname ID", "zh_CN":"调度域名ID"}
  EdgeHostnameId *int64 `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
  // {"en":"Edge-Hostname Name", "zh_CN":"调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en":"Edge-Hostname comment.", "zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"dns service status.data range:[inactive, active]", "zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"deploy status.data range:[pending, deploying, success, fail]", "zh_CN":"部署状态。取值范围：[pending, deploying, success, fail]"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
  // {"en":"allow china cdn.,data range:[0,1].remark: 0 no, 1 yes.", "zh_CN":"是否允许中国大陆加速。取值范围：[0,1]。备注：0 否，1 是。"}
  AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" require:"true"`
  // {"en":"gdpr compliant,data range:[0,1,2].", "zh_CN":"遵循GDPR,取值范围：[0,1,2]。"}
  GdprCompliant *string `json:"gdprCompliant,omitempty" xml:"gdprCompliant,omitempty"`
  // {"en":"geoFence, data range: [global,inside_china_mainland,exclude_china_mainland].", "zh_CN":"加速区域限定。取值范围：[global,inside_china_mainland,exclude_china_mainland]。"}
  GeoFence *string `json:"geoFence,omitempty" xml:"geoFence,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the edge-hostname was created.", "zh_CN":"RFC 3339格式的日期，表示创建edge-hostname的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the edge-hostname was last updated.", "zh_CN":"RFC 3339格式的日期，表示edge-hostname的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames", "zh_CN":"关联的加速域名"}
  Hostnames []*QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en":"region configuration", "zh_CN":"区域配置列表 "}
  RegionConfigs *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeHostnameQueryEdgeHostnameResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameQueryEdgeHostnameResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetEdgeHostnameId(v int64) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.EdgeHostnameId = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetEdgeHostname(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.EdgeHostname = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetComment(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.Comment = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetDnsServiceStatus(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetDeployStatus(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.DeployStatus = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetAllowChinaCdn(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.AllowChinaCdn = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetGdprCompliant(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.GdprCompliant = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetGeoFence(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.GeoFence = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetCreationTime(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.CreationTime = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetLastUpdateTime(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetHostnames(v []*QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.Hostnames = v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseData) SetRegionConfigs(v *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) *QueryEdgeHostnameQueryEdgeHostnameResponseData {
  s.RegionConfigs = v
  return s
}

type QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames struct     {
  // {"en":"hostname", "zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"deploy target", "zh_CN":"部署环境"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"property id", "zh_CN":"项目ID"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"property version", "zh_CN":"项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"property name", "zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
}

func (s QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) SetHostname(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames {
  s.Hostname = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) SetTarget(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames {
  s.Target = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) SetPropertyId(v int64) *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames {
  s.PropertyId = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) SetPropertyVersion(v int) *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames {
  s.PropertyVersion = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames) SetPropertyName(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataHostnames {
  s.PropertyName = &v
  return s
}

type QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs struct {
  // {"en":"region id", "zh_CN":"区域ID"}
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
  // {"en":"action type", "zh_CN":"调度方式。"}
  ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"config type", "zh_CN":"记录类型。"}
  ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty" require:"true"`
  // {"en":"config value", "zh_CN":"记录值。"}
  ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty" require:"true"`
  // {"en":"ip protocol.", "zh_CN":"ip协议。"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"生存时间。"}
  Ttl *string `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"weight", "zh_CN":"权重。"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetRegionId(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.RegionId = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetActionType(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.ActionType = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetConfigType(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.ConfigType = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetConfigValue(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.ConfigValue = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetIpProtocol(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.IpProtocol = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetTtl(v string) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.Ttl = &v
  return s
}

func (s *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs) SetWeight(v int) *QueryEdgeHostnameQueryEdgeHostnameResponseDataRegionConfigs {
  s.Weight = &v
  return s
}

type QueryEdgeHostnamePaths struct {
  // {"en":"edgeHostname", "zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s QueryEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *QueryEdgeHostnamePaths) SetEdgeHostname(v string) *QueryEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type QueryEdgeHostnameParameters struct {
}

func (s QueryEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameParameters) GoString() string {
  return s.String()
}

type QueryEdgeHostnameRequestHeader struct {
}

func (s QueryEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeHostnameResponseHeader struct {
}

func (s QueryEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type DisableEdgeHostnameRequest struct {
}

func (s DisableEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableEdgeHostnameRequest) GoString() string {
  return s.String()
}

type DisableEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DisableEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *DisableEdgeHostnameResponse) SetCode(v string) *DisableEdgeHostnameResponse {
  s.Code = &v
  return s
}

func (s *DisableEdgeHostnameResponse) SetMessage(v string) *DisableEdgeHostnameResponse {
  s.Message = &v
  return s
}

type DisableEdgeHostnamePaths struct {
  // {"en":"edgeHostname", "zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DisableEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s DisableEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *DisableEdgeHostnamePaths) SetEdgeHostname(v string) *DisableEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type DisableEdgeHostnameParameters struct {
}

func (s DisableEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableEdgeHostnameParameters) GoString() string {
  return s.String()
}

type DisableEdgeHostnameRequestHeader struct {
}

func (s DisableEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type DisableEdgeHostnameResponseHeader struct {
}

func (s DisableEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




type UpdateEdgeHostnameRequest struct {
  // {"en":"Edge-Hostname comment.", "zh_CN":"调度域名描述。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"geoFence, data range: [global,inside_china_mainland,exclude_china_mainland].", "zh_CN":"加速区域限定。取值范围：[global,inside_china_mainland,exclude_china_mainland]。"}
  GeoFence *string `json:"geoFence,omitempty" xml:"geoFence,omitempty"`
  // {"en":"region configuration", "zh_CN":"区域配置列表 "}
  RegionConfigs []*UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Repeated"`
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

func (s *UpdateEdgeHostnameRequest) SetRegionConfigs(v []*UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) *UpdateEdgeHostnameRequest {
  s.RegionConfigs = v
  return s
}

type UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs struct     {
  // {"en":"region id", "zh_CN":"区域ID"}
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
  // {"en":"action type", "zh_CN":"调度方式。"}
  ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"config value", "zh_CN":"记录值。"}
  ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
  // {"en":"ip protocol.", "zh_CN":"ip协议。"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
  // {"en":"ttl", "zh_CN":"生存时间。"}
  Ttl *string `json:"ttl,omitempty" xml:"ttl,omitempty"`
  // {"en":"weight", "zh_CN":"权重。"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) GoString() string {
  return s.String()
}

func (s *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) SetRegionId(v string) *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs {
  s.RegionId = &v
  return s
}

func (s *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) SetActionType(v string) *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs {
  s.ActionType = &v
  return s
}

func (s *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) SetConfigValue(v string) *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs {
  s.ConfigValue = &v
  return s
}

func (s *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) SetIpProtocol(v string) *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs {
  s.IpProtocol = &v
  return s
}

func (s *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) SetTtl(v string) *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs {
  s.Ttl = &v
  return s
}

func (s *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs) SetWeight(v int) *UpdateEdgeHostnameUpdateEdgeHostnameRequestRegionConfigs {
  s.Weight = &v
  return s
}

type UpdateEdgeHostnameResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
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

type UpdateEdgeHostnamePaths struct {
  // {"en":"edgeHostname", "zh_CN":"调度域名"}
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

type UpdateEdgeHostnameRequestHeader struct {
}

func (s UpdateEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type UpdateEdgeHostnameResponseHeader struct {
}

func (s UpdateEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}




