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
	Data *QueryEdgeHostnamesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *QueryEdgeHostnamesResponse) SetData(v *QueryEdgeHostnamesResponseData) *QueryEdgeHostnamesResponse {
	s.Data = v
	return s
}

type QueryEdgeHostnamesResponseData struct {
	// {"en":"Number of properties.", "zh_CN":"调度域名数量。"}
	Count *int32 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	// {"en":"List of properties.", "zh_CN":"项目列表。"}
	EdgeHostnames []*QueryEdgeHostnamesResponseDataEdgeHostnames `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeHostnamesResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryEdgeHostnamesResponseData) GoString() string {
	return s.String()
}

func (s *QueryEdgeHostnamesResponseData) SetCount(v int32) *QueryEdgeHostnamesResponseData {
	s.Count = &v
	return s
}

func (s *QueryEdgeHostnamesResponseData) SetEdgeHostnames(v []*QueryEdgeHostnamesResponseDataEdgeHostnames) *QueryEdgeHostnamesResponseData {
	s.EdgeHostnames = v
	return s
}

type QueryEdgeHostnamesResponseDataEdgeHostnames struct {
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
	Hostnames []*QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s QueryEdgeHostnamesResponseDataEdgeHostnames) String() string {
	return tea.Prettify(s)
}

func (s QueryEdgeHostnamesResponseDataEdgeHostnames) GoString() string {
	return s.String()
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetEdgeHostnameId(v int64) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.EdgeHostnameId = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetEdgeHostname(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.EdgeHostname = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetDnsServiceStatus(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.DnsServiceStatus = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetDeployStatus(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.DeployStatus = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetAllowChinaCdn(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.AllowChinaCdn = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetComment(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.Comment = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetCreationTime(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.CreationTime = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetLastUpdateTime(v string) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnames) SetHostnames(v []*QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) *QueryEdgeHostnamesResponseDataEdgeHostnames {
	s.Hostnames = v
	return s
}

type QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames struct {
	// {"en":"name of the domain which use this edge-hostname.", "zh_CN":"域名。"}
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
	// {"en":"the deploy target of this hostname.", "zh_CN":"部署环境。"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
}

func (s QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) String() string {
	return tea.Prettify(s)
}

func (s QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) GoString() string {
	return s.String()
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) SetHostname(v string) *QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames {
	s.Hostname = &v
	return s
}

func (s *QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames) SetTarget(v string) *QueryEdgeHostnamesResponseDataEdgeHostnamesHostnames {
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
	Data *QueryEdgeHostnameResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *QueryEdgeHostnameResponse) SetData(v *QueryEdgeHostnameResponseData) *QueryEdgeHostnameResponse {
	s.Data = v
	return s
}

type QueryEdgeHostnameResponseData struct {
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
	Hostnames []*QueryEdgeHostnameResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
	// {"en":"region configuration", "zh_CN":"区域配置列表 "}
	RegionConfigs *QueryEdgeHostnameResponseDataRegionConfigs `json:"regionConfigs,omitempty" xml:"regionConfigs,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeHostnameResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryEdgeHostnameResponseData) GoString() string {
	return s.String()
}

func (s *QueryEdgeHostnameResponseData) SetEdgeHostnameId(v int64) *QueryEdgeHostnameResponseData {
	s.EdgeHostnameId = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetEdgeHostname(v string) *QueryEdgeHostnameResponseData {
	s.EdgeHostname = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetComment(v string) *QueryEdgeHostnameResponseData {
	s.Comment = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetDnsServiceStatus(v string) *QueryEdgeHostnameResponseData {
	s.DnsServiceStatus = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetDeployStatus(v string) *QueryEdgeHostnameResponseData {
	s.DeployStatus = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetAllowChinaCdn(v string) *QueryEdgeHostnameResponseData {
	s.AllowChinaCdn = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetGdprCompliant(v string) *QueryEdgeHostnameResponseData {
	s.GdprCompliant = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetGeoFence(v string) *QueryEdgeHostnameResponseData {
	s.GeoFence = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetCreationTime(v string) *QueryEdgeHostnameResponseData {
	s.CreationTime = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetLastUpdateTime(v string) *QueryEdgeHostnameResponseData {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetHostnames(v []*QueryEdgeHostnameResponseDataHostnames) *QueryEdgeHostnameResponseData {
	s.Hostnames = v
	return s
}

func (s *QueryEdgeHostnameResponseData) SetRegionConfigs(v *QueryEdgeHostnameResponseDataRegionConfigs) *QueryEdgeHostnameResponseData {
	s.RegionConfigs = v
	return s
}

type QueryEdgeHostnameResponseDataHostnames struct {
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

func (s QueryEdgeHostnameResponseDataHostnames) String() string {
	return tea.Prettify(s)
}

func (s QueryEdgeHostnameResponseDataHostnames) GoString() string {
	return s.String()
}

func (s *QueryEdgeHostnameResponseDataHostnames) SetHostname(v string) *QueryEdgeHostnameResponseDataHostnames {
	s.Hostname = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataHostnames) SetTarget(v string) *QueryEdgeHostnameResponseDataHostnames {
	s.Target = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataHostnames) SetPropertyId(v int64) *QueryEdgeHostnameResponseDataHostnames {
	s.PropertyId = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataHostnames) SetPropertyVersion(v int) *QueryEdgeHostnameResponseDataHostnames {
	s.PropertyVersion = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataHostnames) SetPropertyName(v string) *QueryEdgeHostnameResponseDataHostnames {
	s.PropertyName = &v
	return s
}

type QueryEdgeHostnameResponseDataRegionConfigs struct {
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

func (s QueryEdgeHostnameResponseDataRegionConfigs) String() string {
	return tea.Prettify(s)
}

func (s QueryEdgeHostnameResponseDataRegionConfigs) GoString() string {
	return s.String()
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetRegionId(v string) *QueryEdgeHostnameResponseDataRegionConfigs {
	s.RegionId = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetActionType(v string) *QueryEdgeHostnameResponseDataRegionConfigs {
	s.ActionType = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetConfigType(v string) *QueryEdgeHostnameResponseDataRegionConfigs {
	s.ConfigType = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetConfigValue(v string) *QueryEdgeHostnameResponseDataRegionConfigs {
	s.ConfigValue = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetIpProtocol(v string) *QueryEdgeHostnameResponseDataRegionConfigs {
	s.IpProtocol = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetTtl(v string) *QueryEdgeHostnameResponseDataRegionConfigs {
	s.Ttl = &v
	return s
}

func (s *QueryEdgeHostnameResponseDataRegionConfigs) SetWeight(v int) *QueryEdgeHostnameResponseDataRegionConfigs {
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

type UpdateEdgeHostnameRequestRegionConfigs struct {
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
	EdgeHostnames *string `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" url:"edgeHostnames"`
	// {"en":"hostnames.","zh_CN":"域名列表。多个以英文逗号分割"}
	Hostnames *string `json:"hostnames,omitempty" xml:"hostnames,omitempty" url:"hostnames"`
	// {"en":"comment.","zh_CN":"调度域名描述。"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" url:"comment"`
	// {"en":"The value can be 'inactive', or 'active' to filter the results based on dns service status.","zh_CN":"DNS服务状态。取值范围：[inactive, active]"}
	DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" url:"dnsServiceStatus"`
	// {"en":"The value can be 'pending', or 'deploying', or 'success', or 'fail' to filter the results based on deploy status.remark:split with comma","zh_CN":"部署状态。多个以英文逗号分割"}
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" url:"deployStatus"`
	// {"en":"fuzzy query.only used for edgeHostnames,hostnames and comment","zh_CN":"是否模糊匹配。只对edgeHostnames、hostnames与comment有效"}
	IsLike *string `json:"isLike,omitempty" xml:"isLike,omitempty" url:"isLike"`
	// {"en":"The value can be '0', or '1' to filter the results based on allow china cdn .","zh_CN":"是否允许中国大陆加速。"}
	AllowChinaCdn *string `json:"allowChinaCdn,omitempty" xml:"allowChinaCdn,omitempty" url:"allowChinaCdn"`
	// {"en":"Indicates the first item to return. The default is '0'.","zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
	Offset *int `json:"offset,omitempty" xml:"offset,omitempty" url:"offset"`
	// {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200","zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
	Limit *int `json:"limit,omitempty" xml:"limit,omitempty" url:"limit"`
	// {"en":"Order of properties to return. Enum: asc,desc Default: desc","zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty" url:"sortOrder"`
	// {"en":"Returns results in sorted order. Enum: creationTime,lastUpdateTime,edgeHostname Default: lastUpdateTime","zh_CN":"返回结果的排序依据。取值范围: creationTime,lastUpdateTime,edgeHostname 默认值: lastUpdateTime"}
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty" url:"sortBy"`
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
	// {"en":"Number of properties.","zh_CN":"调度域名数量。"}
	Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	// {"en":"List of properties.","zh_CN":"项目列表。"}
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

type QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames struct {
	// {"en":"Edge-Hostname ID.","zh_CN":"调度域名标识"}
	EdgeHostnameId *int64 `json:"edgeHostnameId,omitempty" xml:"edgeHostnameId,omitempty" require:"true"`
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

func (s *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames) SetEdgeHostnameId(v int64) *QueryEdgeHostnamesForTerraformResponseDataEdgeHostnames {
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

type QueryEdgeHostnamesForTerraformResponseDataEdgeHostnamesHostnames struct {
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

type UpdateEdgeHostnameForTerraformRequestRegionConfigs struct {
	// {"en":"region id","zh_CN":"区域ID"}
	RegionId *int `json:"regionId,omitempty" xml:"regionId,omitempty" require:"true"`
	// {"en":"action type","zh_CN":"调度方式。"}
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

type QueryEdgeHostnameForTerraformResponseDataHostnames struct {
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

type QueryEdgeHostnameForTerraformResponseDataRegionConfigs struct {
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
