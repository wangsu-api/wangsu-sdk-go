package networkmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryEdgePrivateIPRequest struct {
}

func (s QueryEdgePrivateIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPRequest) GoString() string {
  return s.String()
}

type QueryEdgePrivateIPRequestHeader struct {
}

func (s QueryEdgePrivateIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgePrivateIPPaths struct {
}

func (s QueryEdgePrivateIPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPPaths) GoString() string {
  return s.String()
}

type QueryEdgePrivateIPParameters struct {
  // {"en":"node name","zh_CN":"可选\n节点名称，多个用英文逗号分隔"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"virtual machine ID","zh_CN":"可选\n虚拟机ID，多个用英文逗号分隔"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
  // {"en":"extra Ip","zh_CN":"可选\n额外Ip，多个用英文逗号分隔"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty"`
  // {"en":"extra Ip","zh_CN":"可选\n额外公网Ip，多个用英文逗号分隔"}
  PublicEdgeIp *string `json:"publicEdgeIp,omitempty" xml:"publicEdgeIp,omitempty"`
  // {"en":"IP state","zh_CN":"可选\nIP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s QueryEdgePrivateIPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPParameters) GoString() string {
  return s.String()
}

func (s *QueryEdgePrivateIPParameters) SetNodeName(v string) *QueryEdgePrivateIPParameters {
  s.NodeName = &v
  return s
}

func (s *QueryEdgePrivateIPParameters) SetInstanceId(v string) *QueryEdgePrivateIPParameters {
  s.InstanceId = &v
  return s
}

func (s *QueryEdgePrivateIPParameters) SetEdgeIp(v string) *QueryEdgePrivateIPParameters {
  s.EdgeIp = &v
  return s
}

func (s *QueryEdgePrivateIPParameters) SetPublicEdgeIp(v string) *QueryEdgePrivateIPParameters {
  s.PublicEdgeIp = &v
  return s
}

func (s *QueryEdgePrivateIPParameters) SetState(v string) *QueryEdgePrivateIPParameters {
  s.State = &v
  return s
}

type QueryEdgePrivateIPResponse struct {
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *QueryEdgePrivateIPResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgePrivateIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgePrivateIPResponse) SetCode(v string) *QueryEdgePrivateIPResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgePrivateIPResponse) SetMessage(v string) *QueryEdgePrivateIPResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgePrivateIPResponse) SetData(v *QueryEdgePrivateIPResponseData) *QueryEdgePrivateIPResponse {
  s.Data = v
  return s
}

type QueryEdgePrivateIPResponseData struct {
  // {"en":"additional IP details","zh_CN":"额外内网Ip详细信息"}
  EdgeIps []*QueryEdgePrivateIPResponseDataEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgePrivateIPResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgePrivateIPResponseData) SetEdgeIps(v []*QueryEdgePrivateIPResponseDataEdgeIps) *QueryEdgePrivateIPResponseData {
  s.EdgeIps = v
  return s
}

type QueryEdgePrivateIPResponseDataEdgeIps struct     {
  // {"en":"extra Ip","zh_CN":"额外内网Ip"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty" require:"true"`
  // {"en":"IP state","zh_CN":"IP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"binding virtual machine ID","zh_CN":"绑定的虚拟机ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"binding edge public IP","zh_CN":"绑定的额外公网IP"}
  PublicEdgeIp *string `json:"publicEdgeIp,omitempty" xml:"publicEdgeIp,omitempty" require:"true"`
  // {"en":"node name","zh_CN":"所属节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
}

func (s QueryEdgePrivateIPResponseDataEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPResponseDataEdgeIps) GoString() string {
  return s.String()
}

func (s *QueryEdgePrivateIPResponseDataEdgeIps) SetEdgeIp(v string) *QueryEdgePrivateIPResponseDataEdgeIps {
  s.EdgeIp = &v
  return s
}

func (s *QueryEdgePrivateIPResponseDataEdgeIps) SetState(v string) *QueryEdgePrivateIPResponseDataEdgeIps {
  s.State = &v
  return s
}

func (s *QueryEdgePrivateIPResponseDataEdgeIps) SetInstanceId(v string) *QueryEdgePrivateIPResponseDataEdgeIps {
  s.InstanceId = &v
  return s
}

func (s *QueryEdgePrivateIPResponseDataEdgeIps) SetPublicEdgeIp(v string) *QueryEdgePrivateIPResponseDataEdgeIps {
  s.PublicEdgeIp = &v
  return s
}

func (s *QueryEdgePrivateIPResponseDataEdgeIps) SetNodeName(v string) *QueryEdgePrivateIPResponseDataEdgeIps {
  s.NodeName = &v
  return s
}

type QueryEdgePrivateIPResponseHeader struct {
}

func (s QueryEdgePrivateIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgePrivateIPResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryAvailableCidrsDetailRequest struct {
}

func (s VMPQueryAvailableCidrsDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailRequest) GoString() string {
  return s.String()
}

type VMPQueryAvailableCidrsDetailRequestHeader struct {
}

func (s VMPQueryAvailableCidrsDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryAvailableCidrsDetailPaths struct {
}

func (s VMPQueryAvailableCidrsDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailPaths) GoString() string {
  return s.String()
}

type VMPQueryAvailableCidrsDetailParameters struct {
  // {"en":"Node name.","zh_CN":"节点名称，多个节点用英文逗号分隔，最多填写20个"}
  Node *string `json:"node,omitempty" xml:"node,omitempty" require:"true"`
  // {"en":"CIDR","zh_CN":"网段 CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"IP native attribute.(-1: Native, 1: Non-native)","zh_CN":"原生属性. (-1: 原生, 1: 非原生)"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"IPv6 segment(1: yes, -1: no)","zh_CN":"是否返回IPV6网段(1: 是, -1: 否)"}
  NeedIPv6 *string `json:"needIPv6,omitempty" xml:"needIPv6,omitempty"`
}

func (s VMPQueryAvailableCidrsDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsDetailParameters) SetNode(v string) *VMPQueryAvailableCidrsDetailParameters {
  s.Node = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailParameters) SetCidr(v string) *VMPQueryAvailableCidrsDetailParameters {
  s.Cidr = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailParameters) SetNativeAttribute(v string) *VMPQueryAvailableCidrsDetailParameters {
  s.NativeAttribute = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailParameters) SetNeedIPv6(v string) *VMPQueryAvailableCidrsDetailParameters {
  s.NeedIPv6 = &v
  return s
}

type VMPQueryAvailableCidrsDetailResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *VMPQueryAvailableCidrsDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s VMPQueryAvailableCidrsDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsDetailResponse) SetCode(v string) *VMPQueryAvailableCidrsDetailResponse {
  s.Code = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponse) SetData(v *VMPQueryAvailableCidrsDetailResponseData) *VMPQueryAvailableCidrsDetailResponse {
  s.Data = v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponse) SetMessage(v string) *VMPQueryAvailableCidrsDetailResponse {
  s.Message = &v
  return s
}

type VMPQueryAvailableCidrsDetailResponseData struct {
  // {"en":"available cidrs","zh_CN":"可用的cidr列表"}
  Nodes []*VMPQueryAvailableCidrsDetailResponseDataNodes `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryAvailableCidrsDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailResponseData) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsDetailResponseData) SetNodes(v []*VMPQueryAvailableCidrsDetailResponseDataNodes) *VMPQueryAvailableCidrsDetailResponseData {
  s.Nodes = v
  return s
}

type VMPQueryAvailableCidrsDetailResponseDataNodes struct     {
  // {"en":"CIDR detail.","zh_CN":"网段详情"}
  Cidrs []*VMPQueryAvailableCidrsDetailResponseDataNodesCidrs `json:"cidrs,omitempty" xml:"cidrs,omitempty" require:"true" type:"Repeated"`
  // {"en":"Node name.","zh_CN":"节点名称"}
  Node *string `json:"node,omitempty" xml:"node,omitempty" require:"true"`
}

func (s VMPQueryAvailableCidrsDetailResponseDataNodes) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailResponseDataNodes) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodes) SetCidrs(v []*VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) *VMPQueryAvailableCidrsDetailResponseDataNodes {
  s.Cidrs = v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodes) SetNode(v string) *VMPQueryAvailableCidrsDetailResponseDataNodes {
  s.Node = &v
  return s
}

type VMPQueryAvailableCidrsDetailResponseDataNodesCidrs struct     {
  // {"en":"CIDR","zh_CN":"网段 CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en":"Number of free and available IPs","zh_CN":"空闲可用IP数"}
  FreeIps *int `json:"freeIps,omitempty" xml:"freeIps,omitempty" require:"true"`
  // {"en":"Freezing IP number","zh_CN":"冷却IP数"}
  FreezingIps *int `json:"freezingIps,omitempty" xml:"freezingIps,omitempty" require:"true"`
  // {"en":"IP native attribute.(-1: Native, 1: Non-native)","zh_CN":"原生属性. (-1: 原生, 1: 非原生)"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty" require:"true"`
  // {"en":"Number of IPs already used","zh_CN":"已用IP数"}
  UsedIps *int `json:"usedIps,omitempty" xml:"usedIps,omitempty" require:"true"`
}

func (s VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) SetCidr(v string) *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.Cidr = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) SetFreeIps(v int) *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.FreeIps = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) SetFreezingIps(v int) *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.FreezingIps = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) SetNativeAttribute(v string) *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.NativeAttribute = &v
  return s
}

func (s *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs) SetUsedIps(v int) *VMPQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.UsedIps = &v
  return s
}

type VMPQueryAvailableCidrsDetailResponseHeader struct {
}

func (s VMPQueryAvailableCidrsDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsDetailResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryEdgeIPRequest struct {
}

func (s LECHQueryEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPRequest) GoString() string {
  return s.String()
}

type LECHQueryEdgeIPRequestHeader struct {
}

func (s LECHQueryEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryEdgeIPPaths struct {
}

func (s LECHQueryEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPPaths) GoString() string {
  return s.String()
}

type LECHQueryEdgeIPParameters struct {
  // {"en":"node name","zh_CN":"可选\n节点名称，多个用英文逗号分隔"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"virtual machine ID","zh_CN":"可选\n虚拟机ID，多个用英文逗号分隔"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty"`
  // {"en":"virtual machine master IP","zh_CN":"可选\n虚拟机主IP，多个用英文逗号分隔"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty"`
  // {"en":"extra Ip","zh_CN":"可选\n额外Ip，多个用英文逗号分隔"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty"`
  // {"en":"IP state","zh_CN":"可选\nIP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s LECHQueryEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryEdgeIPParameters) SetNodeName(v string) *LECHQueryEdgeIPParameters {
  s.NodeName = &v
  return s
}

func (s *LECHQueryEdgeIPParameters) SetServerId(v string) *LECHQueryEdgeIPParameters {
  s.ServerId = &v
  return s
}

func (s *LECHQueryEdgeIPParameters) SetServerIp(v string) *LECHQueryEdgeIPParameters {
  s.ServerIp = &v
  return s
}

func (s *LECHQueryEdgeIPParameters) SetEdgeIp(v string) *LECHQueryEdgeIPParameters {
  s.EdgeIp = &v
  return s
}

func (s *LECHQueryEdgeIPParameters) SetState(v string) *LECHQueryEdgeIPParameters {
  s.State = &v
  return s
}

type LECHQueryEdgeIPResponse struct {
  // {"en":"additional IP details","zh_CN":"额外Ip详细信息"}
  EdgeIps []*LECHQueryEdgeIPResponseEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryEdgeIPResponse) SetEdgeIps(v []*LECHQueryEdgeIPResponseEdgeIps) *LECHQueryEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type LECHQueryEdgeIPResponseEdgeIps struct     {
  // {"en":"extra Ip","zh_CN":"额外Ip"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty" require:"true"`
  // {"en":"IP state","zh_CN":"IP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"binding virtual machine ID","zh_CN":"绑定的虚拟机ID"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"binding virtual machine extranet IP","zh_CN":"绑定的虚拟机外网IP"}
  ServerIp []*string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true" type:"Repeated"`
  // {"en":"node name","zh_CN":"所属节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Is it exclusive","zh_CN":"是否独占"}
  OccupancyFlag *bool `json:"occupancyFlag,omitempty" xml:"occupancyFlag,omitempty" require:"true"`
  // {"en":"Net Mask","zh_CN":"子网掩码"}
  Netmask *string `json:"netmask,omitempty" xml:"netmask,omitempty" require:"true"`
}

func (s LECHQueryEdgeIPResponseEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPResponseEdgeIps) GoString() string {
  return s.String()
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetEdgeIp(v string) *LECHQueryEdgeIPResponseEdgeIps {
  s.EdgeIp = &v
  return s
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetState(v string) *LECHQueryEdgeIPResponseEdgeIps {
  s.State = &v
  return s
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetServerId(v string) *LECHQueryEdgeIPResponseEdgeIps {
  s.ServerId = &v
  return s
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetServerIp(v []*string) *LECHQueryEdgeIPResponseEdgeIps {
  s.ServerIp = v
  return s
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetNodeName(v string) *LECHQueryEdgeIPResponseEdgeIps {
  s.NodeName = &v
  return s
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetOccupancyFlag(v bool) *LECHQueryEdgeIPResponseEdgeIps {
  s.OccupancyFlag = &v
  return s
}

func (s *LECHQueryEdgeIPResponseEdgeIps) SetNetmask(v string) *LECHQueryEdgeIPResponseEdgeIps {
  s.Netmask = &v
  return s
}

type LECHQueryEdgeIPResponseHeader struct {
}

func (s LECHQueryEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type DeleteServiceRequest struct {
}

func (s DeleteServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
  return s.String()
}

type DeleteServiceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"status"}
  Data *DeleteServiceStatus `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
  return s.String()
}

func (s *DeleteServiceResponse) SetCode(v int64) *DeleteServiceResponse {
  s.Code = &v
  return s
}

func (s *DeleteServiceResponse) SetMsg(v string) *DeleteServiceResponse {
  s.Msg = &v
  return s
}

func (s *DeleteServiceResponse) SetRequestId(v string) *DeleteServiceResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteServiceResponse) SetData(v *DeleteServiceStatus) *DeleteServiceResponse {
  s.Data = v
  return s
}

type DeleteServicePaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"service name", "zh_CN":"service 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteServicePaths) GoString() string {
  return s.String()
}

func (s *DeleteServicePaths) SetNamespace(v string) *DeleteServicePaths {
  s.Namespace = &v
  return s
}

func (s *DeleteServicePaths) SetName(v string) *DeleteServicePaths {
  s.Name = &v
  return s
}

type DeleteServiceParameters struct {
}

func (s DeleteServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceParameters) GoString() string {
  return s.String()
}

type DeleteServiceRequestHeader struct {
}

func (s DeleteServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceRequestHeader) GoString() string {
  return s.String()
}

type DeleteServiceResponseHeader struct {
}

func (s DeleteServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceResponseHeader) GoString() string {
  return s.String()
}

type DeleteServiceStatus struct {
  // {"en":"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values", "zh_CN":"APIVersion 定义对象表示的版本化模式。 服务器应将已识别的模式转换为最新的内部值，并可能拒绝无法识别的值"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase", "zh_CN":"Kind 是一个字符串值，表示此对象表示的 REST 资源。 服务器可以从客户端提交请求的端点推断出这一点。 无法更新。驼峰式规则"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"DeleteServiceStatus of the operation. One of: 'Success' or 'Failure'", "zh_CN":"操作状态。“Success”或“Failure” 之一"}
  DeleteServiceStatus *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Suggested HTTP return code for this status, 0 if not set", "zh_CN":"此状态的建议 HTTP 返回代码，如果未设置，则为 0"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Extended data associated with the reason. Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type", "zh_CN":"与原因（Reason）相关的扩展数据。每个原因都可以定义自己的扩展细节。 此字段是可选的，并且不保证返回的数据符合任何模式，除非由原因类型定义"}
  Details *DeleteServiceStatusDetails `json:"details,omitempty" xml:"details,omitempty" require:"true"`
}

func (s DeleteServiceStatus) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceStatus) GoString() string {
  return s.String()
}

func (s *DeleteServiceStatus) SetApiVersion(v string) *DeleteServiceStatus {
  s.ApiVersion = &v
  return s
}

func (s *DeleteServiceStatus) SetKind(v string) *DeleteServiceStatus {
  s.Kind = &v
  return s
}

func (s *DeleteServiceStatus) SetStatus(v string) *DeleteServiceStatus {
  s.DeleteServiceStatus = &v
  return s
}

func (s *DeleteServiceStatus) SetCode(v int32) *DeleteServiceStatus {
  s.Code = &v
  return s
}

func (s *DeleteServiceStatus) SetDetails(v *DeleteServiceStatusDetails) *DeleteServiceStatus {
  s.Details = v
  return s
}

type DeleteServiceStatusDetails struct {
  // {"en":"The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)", "zh_CN":"与状态 StatusReason 关联的资源的名称属性（当有一个可以描述的名称时）"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind", "zh_CN":"与状态 StatusReason 关联的资源的种类属性"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"The group attribute of the resource associated with the status StatusReason", "zh_CN":"与状态 StatusReason 关联的资源的组属性"}
  Group *string `json:"group,omitempty" xml:"group,omitempty" require:"true"`
  // {"en":"UID of the resource. (when there is a single resource which can be described)", "zh_CN":"资源的 UID（当有单个可以描述的资源时）"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty" require:"true"`
}

func (s DeleteServiceStatusDetails) String() string {
  return tea.Prettify(s)
}

func (s DeleteServiceStatusDetails) GoString() string {
  return s.String()
}

func (s *DeleteServiceStatusDetails) SetName(v string) *DeleteServiceStatusDetails {
  s.Name = &v
  return s
}

func (s *DeleteServiceStatusDetails) SetKind(v string) *DeleteServiceStatusDetails {
  s.Kind = &v
  return s
}

func (s *DeleteServiceStatusDetails) SetGroup(v string) *DeleteServiceStatusDetails {
  s.Group = &v
  return s
}

func (s *DeleteServiceStatusDetails) SetUid(v string) *DeleteServiceStatusDetails {
  s.Uid = &v
  return s
}




type CreateServiceRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreateServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 CreateServiceService 的行为"}
  Spec *CreateServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
  return s.String()
}

func (s *CreateServiceRequest) SetApiVersion(v string) *CreateServiceRequest {
  s.ApiVersion = &v
  return s
}

func (s *CreateServiceRequest) SetKind(v string) *CreateServiceRequest {
  s.Kind = &v
  return s
}

func (s *CreateServiceRequest) SetMetadata(v *CreateServiceObjectMeta) *CreateServiceRequest {
  s.Metadata = v
  return s
}

func (s *CreateServiceRequest) SetSpec(v *CreateServiceServiceSpec) *CreateServiceRequest {
  s.Spec = v
  return s
}

type CreateServiceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"service", "zh_CN":"service"}
  Data *CreateServiceService `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
  return s.String()
}

func (s *CreateServiceResponse) SetCode(v int64) *CreateServiceResponse {
  s.Code = &v
  return s
}

func (s *CreateServiceResponse) SetMsg(v string) *CreateServiceResponse {
  s.Msg = &v
  return s
}

func (s *CreateServiceResponse) SetRequestId(v string) *CreateServiceResponse {
  s.RequestId = &v
  return s
}

func (s *CreateServiceResponse) SetData(v *CreateServiceService) *CreateServiceResponse {
  s.Data = v
  return s
}

type CreateServicePaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s CreateServicePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateServicePaths) GoString() string {
  return s.String()
}

func (s *CreateServicePaths) SetNamespace(v string) *CreateServicePaths {
  s.Namespace = &v
  return s
}

type CreateServiceParameters struct {
}

func (s CreateServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceParameters) GoString() string {
  return s.String()
}

type CreateServiceRequestHeader struct {
}

func (s CreateServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceRequestHeader) GoString() string {
  return s.String()
}

type CreateServiceResponseHeader struct {
}

func (s CreateServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceResponseHeader) GoString() string {
  return s.String()
}

type CreateServiceService struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreateServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 CreateServiceService 的行为"}
  Spec *CreateServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the service. Populated by the system. Read-only", "zh_CN":"最近观察到的 CreateServiceService 状态。由系统填充。只读"}
  Status *CreateServiceServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateServiceService) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceService) GoString() string {
  return s.String()
}

func (s *CreateServiceService) SetApiVersion(v string) *CreateServiceService {
  s.ApiVersion = &v
  return s
}

func (s *CreateServiceService) SetKind(v string) *CreateServiceService {
  s.Kind = &v
  return s
}

func (s *CreateServiceService) SetMetadata(v *CreateServiceObjectMeta) *CreateServiceService {
  s.Metadata = v
  return s
}

func (s *CreateServiceService) SetSpec(v *CreateServiceServiceSpec) *CreateServiceService {
  s.Spec = v
  return s
}

func (s *CreateServiceService) SetStatus(v *CreateServiceServiceStatus) *CreateServiceService {
  s.Status = v
  return s
}

type CreateServiceServiceStatus struct {
  // {"en":"Current service state", "zh_CN":"loadBalancer 包含负载均衡器的当前状态（如果存在）"}
  LoadBalancer *CreateServiceLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
  // {"en":"LoadBalancer contains the current status of the load-balancer, if one is present", "zh_CN":"服务的当前状态"}
  Conditions []*CreateServiceMetaV1Condition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s CreateServiceServiceStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceServiceStatus) GoString() string {
  return s.String()
}

func (s *CreateServiceServiceStatus) SetLoadBalancer(v *CreateServiceLoadBalancerStatus) *CreateServiceServiceStatus {
  s.LoadBalancer = v
  return s
}

func (s *CreateServiceServiceStatus) SetConditions(v []*CreateServiceMetaV1Condition) *CreateServiceServiceStatus {
  s.Conditions = v
  return s
}

type CreateServiceMetaV1Condition struct {
  // {"en":"type of condition in CamelCase or in foo.example.com/CamelCase", "zh_CN":"CamelCase 或 foo.example.com/CamelCase 中的条件类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status of the condition, one of True, False, Unknown", "zh_CN":"condition 的状态，True、False、Unknown 之一"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance", "zh_CN":"表示设置 condition 基于的 .metadata.generation 的过期次数。 例如，如果 .metadata.generation 当前为 12，但 .status.conditions[x].observedGeneration 为 9， 则 condition 相对于实例的当前状态已过期"}
  ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
  // {"en":"lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable", "zh_CN":"状况最近一次状态转化的时间。 变化应该发生在下层状况发生变化的时候。如果不知道下层状况发生变化的时间， 那么使用 API 字段更改的时间是可以接受的"}
  LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
  // {"en":"reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty", "zh_CN":"reason 包含一个程序标识符，指示 condition 最后一次转换的原因。 特定条件类型的生产者可以定义该字段的预期值和含义，以及这些值是否被视为有保证的 API。 该值应该是 CamelCase 字符串且不能为空"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":"message is a human readable message indicating details about the transition. This may be an empty string", "zh_CN":"message 是人类可读的消息，有关转换的详细信息，可以是空字符串"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateServiceMetaV1Condition) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceMetaV1Condition) GoString() string {
  return s.String()
}

func (s *CreateServiceMetaV1Condition) SetType(v string) *CreateServiceMetaV1Condition {
  s.Type = &v
  return s
}

func (s *CreateServiceMetaV1Condition) SetStatus(v string) *CreateServiceMetaV1Condition {
  s.Status = &v
  return s
}

func (s *CreateServiceMetaV1Condition) SetObservedGeneration(v int64) *CreateServiceMetaV1Condition {
  s.ObservedGeneration = &v
  return s
}

func (s *CreateServiceMetaV1Condition) SetLastTransitionTime(v string) *CreateServiceMetaV1Condition {
  s.LastTransitionTime = &v
  return s
}

func (s *CreateServiceMetaV1Condition) SetReason(v string) *CreateServiceMetaV1Condition {
  s.Reason = &v
  return s
}

func (s *CreateServiceMetaV1Condition) SetMessage(v string) *CreateServiceMetaV1Condition {
  s.Message = &v
  return s
}

type CreateServiceLoadBalancerStatus struct {
  // {"en":"Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points", "zh_CN":"ingress 是一个包含负载均衡器 Ingress 点的列表。CreateServiceService 的流量需要被发送到这些 Ingress 点"}
  Ingress []*CreateServiceLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s CreateServiceLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *CreateServiceLoadBalancerStatus) SetIngress(v []*CreateServiceLoadBalancerIngress) *CreateServiceLoadBalancerStatus {
  s.Ingress = v
  return s
}

type CreateServiceLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)", "zh_CN":"ip 是为基于 IP 的负载均衡器 Ingress 点（通常是 GCE 或 OpenStack 负载均衡器）设置的"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)", "zh_CN":"hostname 是为基于 DNS 的负载均衡器 Ingress 点（通常是 AWS 负载均衡器）设置的"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
  // {"en":"Ports is a list of records of service ports If used, every port defined in the service should have an entry in it", "zh_CN":"ports 是 CreateServiceService 的端口列表。如果设置了此字段，CreateServiceService 中定义的每个端口都应该在此列表中"}
  Ports []*CreateServicePortStatus `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s CreateServiceLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *CreateServiceLoadBalancerIngress) SetIp(v string) *CreateServiceLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *CreateServiceLoadBalancerIngress) SetHostName(v string) *CreateServiceLoadBalancerIngress {
  s.HostName = &v
  return s
}

func (s *CreateServiceLoadBalancerIngress) SetPorts(v []*CreateServicePortStatus) *CreateServiceLoadBalancerIngress {
  s.Ports = v
  return s
}

type CreateServicePortStatus struct {
  // {"en":"the port number of the service port of which status is recorded here", "zh_CN":"port 是所记录的服务端口状态的端口号"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"the protocol of the service port of which status is recorded here The supported values are: 'TCP', 'UDP', 'SCTP'", "zh_CN":"protocol 是所记录的服务端口状态的协议。支持的值为：“TCP”、”UDP”、“SCTP”"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase", "zh_CN":"error 是记录 CreateServiceService 端口的问题。 错误的格式应符合以下规则:内置错误原因应在此文件中指定，应使用 CamelCase 名称。云提供商特定错误原因的名称必须符合格式 foo.example.com/CamelCase"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s CreateServicePortStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateServicePortStatus) GoString() string {
  return s.String()
}

func (s *CreateServicePortStatus) SetPort(v int32) *CreateServicePortStatus {
  s.Port = &v
  return s
}

func (s *CreateServicePortStatus) SetProtocol(v string) *CreateServicePortStatus {
  s.Protocol = &v
  return s
}

func (s *CreateServicePortStatus) SetError(v string) *CreateServicePortStatus {
  s.Error = &v
  return s
}

type CreateServiceServiceSpec struct {
  // {"en":"The list of ports that are exposed by this service", "zh_CN":"此 CreateServiceService 公开的端口列表"}
  Ports []*CreateServiceServicePort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en":"Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName", "zh_CN":"将 CreateServiceService 流量路由到具有与此 selector 匹配的标签键值对的 Pod。 如果为空或不存在，则假定该服务有一个外部进程管理其端点，Kubernetes 不会修改该端点。 仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型。如果类型为 ExternalName，则忽略"}
  Selector map[string]*string `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a CreateServiceService of type ExternalName, creation will fail. This field will be wiped when updating a CreateServiceService to type ExternalName", "zh_CN":"clusterIP 是服务的 IP 地址，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给服务，否则创建服务将失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIP 为空）或 type 已经是 ExternalName 时，可以更改 clusterIP（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIP 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 仅适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 CreateServiceService 时指定了 clusterIP，则创建将失败。 更新 CreateServiceService type 为 ExternalName 时，clusterIP 会被移除"}
  ClusterIP *string `json:"clusterIP,omitempty" xml:"clusterIP,omitempty"`
  // {"en":"ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a CreateServiceService of type ExternalName, creation will fail. This field will be wiped when updating a CreateServiceService to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"clusterIPs 是分配给该 CreateServiceService 的 IP 地址列表，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给 CreateServiceService；否则创建 CreateServiceService 失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIPs 为空）或 type 已经是 ExternalName 时，可以更改 clusterIPs（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIPs 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 CreateServiceService 时指定了 clusterIPs，则会创建失败。 更新 CreateServiceService type 为 ExternalName 时，该字段将被移除。如果未指定此字段，则将从 clusterIP 字段初始化。 如果指定 clusterIPs，客户端必须确保 clusterIPs[0] 和 clusterIP 一致。clusterIPs 最多可包含两个条目（双栈系列，按任意顺序）。 这些 IP 必须与 ipFamilies 的值相对应。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 管理"}
  ClusterIPs []*string `json:"clusterIPs,omitempty" xml:"clusterIPs,omitempty" type:"Repeated"`
  // {"en":"type determines how the CreateServiceService is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services", "zh_CN":"type 确定 CreateServiceService 的公开方式。默认为 ClusterIP。 有效选项为 ExternalName、ClusterIP、NodePort 和 LoadBalancer。 “ClusterIP” 为端点分配一个集群内部 IP 地址用于负载均衡。 Endpoints 由 selector 确定，如果未设置 selector，则需要通过手动构造 Endpoints 或 EndpointSlice 的对象来确定。 如果 clusterIP 为 “None”，则不分配虚拟 IP，并且 Endpoints 作为一组端点而不是虚拟 IP 发布。 “NodePort” 建立在 ClusterIP 之上，并在每个节点上分配一个端口，该端口路由到与 clusterIP 相同的 Endpoints。 “LoadBalancer” 基于 NodePort 构建并创建一个外部负载均衡器（如果当前云支持），该负载均衡器路由到与 clusterIP 相同的 Endpoints。 “externalName” 将此 CreateServiceService 别名为指定的 externalName。其他几个字段不适用于 ExternalName CreateServiceService"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system", "zh_CN":"externalIPs 是一个 IP 列表，集群中的节点会为此 CreateServiceService 接收针对这些 IP 地址的流量。 这些 IP 不被 Kubernetes 管理。用户需要确保流量可以到达具有此 IP 的节点。 一个常见的例子是不属于 Kubernetes 系统的外部负载均衡器"}
  ExternalIPs []*string `json:"externalIPs,omitempty" xml:"externalIPs,omitempty" type:"Repeated"`
  // {"en":"Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None", "zh_CN":"支持 “ClientIP” 和 “None”。用于维护会话亲和性。 启用基于客户端 IP 的会话亲和性。必须是 ClientIP 或 None。默认为 None"}
  SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
  // {"en":"Only applies to CreateServiceService Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version", "zh_CN":"仅适用于服务类型: LoadBalancer。此功能取决于底层云提供商是否支持负载均衡器。 如果云提供商不支持该功能，该字段将被忽略。 已弃用: 该字段信息不足，且其含义因实现而异，而且不支持双栈。 从 Kubernetes v1.24 开始，鼓励用户在可用时使用特定于实现的注释。在未来的 API 版本中可能会删除此字段"}
  LoadBalancerIP *string `json:"loadBalancerIP,omitempty" xml:"loadBalancerIP,omitempty"`
  // {"en":"If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature", "zh_CN":"如果设置了此字段并且被平台支持，将限制通过云厂商的负载均衡器的流量到指定的客户端 IP。 如果云提供商不支持该功能，该字段将被忽略"}
  LoadBalancerSourceRanges []*string `json:"loadBalancerSourceRanges,omitempty" xml:"loadBalancerSourceRanges,omitempty" type:"Repeated"`
  // {"en":"externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires type to be 'ExternalName'", "zh_CN":"externalName 是发现机制将返回的外部引用，作为此服务的别名（例如 DNS CNAME 记录）。 不涉及代理。必须是小写的 RFC-1123 主机名 (https://tools.ietf.org/html/rfc1123)， 并且要求 type 为 “ExternalName”"}
  ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
  // {"en":"externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the CreateServiceService's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node", "zh_CN":"externalTrafficPolicy 描述了节点如何分发它们在 CreateServiceService 的“外部访问”地址 （NodePort、ExternalIP 和 LoadBalancer IP）接收到的服务流量。 如果设置为 “Local”，代理将以一种假设外部负载均衡器将负责在节点之间服务流量负载均衡， 因此每个节点将仅向服务的节点本地端点传递流量，而不会伪装客户端源 IP。 （将丢弃错误发送到没有端点的节点的流量。） “Cluster” 默认值使用负载均衡路由到所有端点的策略（可能会根据拓扑和其他特性进行修改）。 请注意，从集群内部发送到 External IP 或 LoadBalancer IP 的流量始终具有 “Cluster” 语义， 但是从集群内部发送到 NodePort 的客户端需要在选择节点时考虑流量路由策略"}
  ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" xml:"externalTrafficPolicy,omitempty"`
  // {"en":"healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a CreateServiceService which does not need it, creation will fail. This field will be wiped when updating a CreateServiceService to no longer need it (e.g. changing type). This field cannot be updated once set", "zh_CN":"healthCheckNodePort 指定 CreateServiceService 的健康检查节点端口。 仅适用于 type 为 LoadBalancer 且 externalTrafficPolicy 设置为 Local 的情况。 如果为此字段设定了一个值，该值在合法范围内且没有被使用，则使用所指定的值。 如果未设置此字段，则自动分配字段值。外部系统（例如负载平衡器）可以使用此端口来确定给定节点是否拥有此服务的端点。 在创建不需要 healthCheckNodePort 的 CreateServiceService 时指定了此字段，则 CreateServiceService 创建会失败。 要移除 healthCheckNodePort，需要更改 CreateServiceService 的 type。 该字段一旦设置就无法更改"}
  HealthCheckNodePort *int32 `json:"healthCheckNodePort,omitempty" xml:"healthCheckNodePort,omitempty"`
  // {"en":"publishNotReadyAddresses indicates that any agent which deals with endpoints for this CreateServiceService should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless CreateServiceService to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior", "zh_CN":"publishNotReadyAddresses 表示任何处理此 CreateServiceService 端点的代理都应忽略任何准备就绪/未准备就绪的指示。 设置此字段的主要场景是为 StatefulSet 的服务提供支持，使之能够为其 Pod 传播 SRV DNS 记录，以实现对等发现。 为 CreateServiceService 生成 Endpoints 和 EndpointSlice 资源的 Kubernetes 控制器对字段的解读是， 即使 Pod 本身还没有准备好，所有端点都可被视为 “已就绪”。 对于代理而言，如果仅使用 Kubernetes 通过 Endpoints 或 EndpointSlice 资源所生成的端点， 则可以安全地假设这种行为"}
  PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" xml:"publishNotReadyAddresses,omitempty"`
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"sessionAffinityConfig 包含会话亲和性的配置"}
  CreateServiceSessionAffinityConfig *CreateServiceSessionAffinityConfig `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
  // {"en":"IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the CreateServiceService. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a CreateServiceService to type ExternalName.This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"iPFamilies 是分配给此服务的 IP 协议（例如 IPv4、IPv6）的列表。 该字段通常根据集群配置和 ipFamilyPolicy 字段自动设置。 如果手动指定该字段，且请求的协议在集群中可用，且 ipFamilyPolicy 允许，则使用；否则服务创建将失败。 该字段修改是有条件的：它允许添加或删除辅助 IP 协议，但不允许更改服务的主要 IP 协议。 有效值为 “IPv4” 和 “IPv6”。 该字段仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型的服务，并且确实可用于“无头”服务。 更新服务设置类型为 ExternalName 时，该字段将被擦除。该字段最多可以包含两个条目（双栈系列，按任意顺序）。 如果指定，这些协议栈必须对应于 clusterIPs 字段的值。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 字段管理"}
  IpFamilies []*string `json:"ipFamilies,omitempty" xml:"ipFamilies,omitempty" type:"Repeated"`
  // {"en":"IPFamilyPolicy represents the dual-stack-ness requested or required by this CreateServiceService. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName", "zh_CN":"iPFamilyPolicy 表示此服务请求或要求的双栈特性。 如果没有提供值，则此字段将被设置为 SingleStack。 服务可以是 “SingleStack”（单个 IP 协议）、 “PreferDualStack”（双栈配置集群上的两个 IP 协议或单栈集群上的单个 IP 协议） 或 “RequireDualStack”（双栈上的两个 IP 协议配置的集群，否则失败）。 ipFamilies 和 clusterIPs 字段取决于此字段的值。 更新服务设置类型为 ExternalName 时，此字段将被擦除"}
  IpFamilyPolicy *string `json:"ipFamilyPolicy,omitempty" xml:"ipFamilyPolicy,omitempty"`
  // {"en":"allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type", "zh_CN":"allocateLoadBalancerNodePorts 定义了是否会自动为 LoadBalancer 类型的 CreateServiceService 分配 NodePort。默认为 true。 如果集群负载均衡器不依赖 NodePort，则可以设置此字段为 false。 如果调用者（通过指定一个值）请求特定的 NodePort，则无论此字段如何，都会接受这些请求。 该字段只能设置在 type 为 LoadBalancer 的 CreateServiceService 上，如果 type 更改为任何其他类型，该字段将被移除"}
  AllocateLoadBalancerNodePorts *bool `json:"allocateLoadBalancerNodePorts,omitempty" xml:"allocateLoadBalancerNodePorts,omitempty"`
  // {"en":"loadBalancerClass is the class of the load balancer implementation this CreateServiceService belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the CreateServiceService type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a CreateServiceService to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type", "zh_CN":"loadBalancerClass 是此 CreateServiceService 所属的负载均衡器实现的类。 如果设置了此字段，则字段值必须是标签风格的标识符，带有可选前缀，例如 ”internal-vip” 或 “example.com/internal-vip”。 无前缀名称是为最终用户保留的。该字段只能在 CreateServiceService 类型为 “LoadBalancer” 时设置。 如果未设置此字段，则使用默认负载均衡器实现。默认负载均衡器现在通常通过云提供商集成完成，但应适用于任何默认实现。 如果设置了此字段，则假定负载均衡器实现正在监测具有对应负载均衡器类的 CreateServiceService。 任何默认负载均衡器实现（例如云提供商）都应忽略设置此字段的 CreateServiceService。 只有在创建或更新的 CreateServiceService 的 type 为 “LoadBalancer” 时，才可设置此字段。 一经设定，不可更改。当 CreateServiceService 的 type 更新为 “LoadBalancer” 之外的其他类型时，此字段将被移除"}
  LoadBalancerClass *string `json:"loadBalancerClass,omitempty" xml:"loadBalancerClass,omitempty"`
  // {"en":"InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features)", "zh_CN":"InternalTrafficPolicy 描述节点如何分发它们在 ClusterIP 上接收到的服务流量。 如果设置为 “Local”，代理将假定 Pod 只想与在同一节点上的服务端点通信，如果没有本地端点，它将丢弃流量。 “Cluster” 默认将流量路由到所有端点（可能会根据拓扑和其他特性进行修改）"}
  InternalTrafficPolicy *string `json:"internalTrafficPolicy,omitempty" xml:"internalTrafficPolicy,omitempty"`
}

func (s CreateServiceServiceSpec) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceServiceSpec) GoString() string {
  return s.String()
}

func (s *CreateServiceServiceSpec) SetPorts(v []*CreateServiceServicePort) *CreateServiceServiceSpec {
  s.Ports = v
  return s
}

func (s *CreateServiceServiceSpec) SetSelector(v map[string]*string) *CreateServiceServiceSpec {
  s.Selector = v
  return s
}

func (s *CreateServiceServiceSpec) SetClusterIP(v string) *CreateServiceServiceSpec {
  s.ClusterIP = &v
  return s
}

func (s *CreateServiceServiceSpec) SetClusterIPs(v []*string) *CreateServiceServiceSpec {
  s.ClusterIPs = v
  return s
}

func (s *CreateServiceServiceSpec) SetType(v string) *CreateServiceServiceSpec {
  s.Type = &v
  return s
}

func (s *CreateServiceServiceSpec) SetExternalIPs(v []*string) *CreateServiceServiceSpec {
  s.ExternalIPs = v
  return s
}

func (s *CreateServiceServiceSpec) SetSessionAffinity(v string) *CreateServiceServiceSpec {
  s.SessionAffinity = &v
  return s
}

func (s *CreateServiceServiceSpec) SetLoadBalancerIP(v string) *CreateServiceServiceSpec {
  s.LoadBalancerIP = &v
  return s
}

func (s *CreateServiceServiceSpec) SetLoadBalancerSourceRanges(v []*string) *CreateServiceServiceSpec {
  s.LoadBalancerSourceRanges = v
  return s
}

func (s *CreateServiceServiceSpec) SetExternalName(v string) *CreateServiceServiceSpec {
  s.ExternalName = &v
  return s
}

func (s *CreateServiceServiceSpec) SetExternalTrafficPolicy(v string) *CreateServiceServiceSpec {
  s.ExternalTrafficPolicy = &v
  return s
}

func (s *CreateServiceServiceSpec) SetHealthCheckNodePort(v int32) *CreateServiceServiceSpec {
  s.HealthCheckNodePort = &v
  return s
}

func (s *CreateServiceServiceSpec) SetPublishNotReadyAddresses(v bool) *CreateServiceServiceSpec {
  s.PublishNotReadyAddresses = &v
  return s
}

func (s *CreateServiceServiceSpec) SetSessionAffinityConfig(v *CreateServiceSessionAffinityConfig) *CreateServiceServiceSpec {
  s.CreateServiceSessionAffinityConfig = v
  return s
}

func (s *CreateServiceServiceSpec) SetIpFamilies(v []*string) *CreateServiceServiceSpec {
  s.IpFamilies = v
  return s
}

func (s *CreateServiceServiceSpec) SetIpFamilyPolicy(v string) *CreateServiceServiceSpec {
  s.IpFamilyPolicy = &v
  return s
}

func (s *CreateServiceServiceSpec) SetAllocateLoadBalancerNodePorts(v bool) *CreateServiceServiceSpec {
  s.AllocateLoadBalancerNodePorts = &v
  return s
}

func (s *CreateServiceServiceSpec) SetLoadBalancerClass(v string) *CreateServiceServiceSpec {
  s.LoadBalancerClass = &v
  return s
}

func (s *CreateServiceServiceSpec) SetInternalTrafficPolicy(v string) *CreateServiceServiceSpec {
  s.InternalTrafficPolicy = &v
  return s
}

type CreateServiceServicePort struct {
  // {"en":"The name of this port within the service. This must be a DNS_LABEL. All ports within a CreateServiceServiceSpec must have unique names. When considering the endpoints for a CreateServiceService, this must match the 'name' field in the EndpointPort. Optional if only one CreateServiceServicePort is defined on this service", "zh_CN":"CreateServiceService 中此端口的名称。这必须是 DNS_LABEL。 CreateServiceServiceSpec 中的所有端口的名称都必须唯一。 在考虑 CreateServiceService 的端点时，这一字段值必须与 EndpointPort 中的 name 字段相同。 如果此服务上仅定义一个 CreateServiceServicePort，则为此字段为可选"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The IP protocol for this port. Supports TCP, UDP, and SCTP. Default is TCP", "zh_CN":"此端口的 IP 协议。支持 “TCP”、“UDP” 和 “SCTP”。默认为 TCP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol", "zh_CN":"此端口的应用协议，遵循标准的 Kubernetes 标签语法，无前缀名称按照 IANA 标准服务名称 （参见 RFC-6335 和 https://www.iana.org/assignments/service-names）。 非标准协议应该使用前缀名称，如 mycompany.com/my-custom-protocol"}
  AppProtocol *string `json:"appProtocol,omitempty" xml:"appProtocol,omitempty"`
  // {"en":"The port that will be exposed by this service", "zh_CN":"CreateServiceService 将公开的端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field", "zh_CN":"在 CreateServiceService 所针对的 Pod 上要访问的端口号或名称。 编号必须在 1 到 65535 的范围内。名称必须是 IANA_SVC_NAME。 如果此值是一个字符串，将在目标 Pod 的容器端口中作为命名端口进行查找。 如果未指定字段，则使用 “port” 字段的值（直接映射）。 对于 clusterIP 为 None 的服务，此字段将被忽略， 应忽略不设或设置为 “port” 字段的取值"}
  TargetPort *int32 `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
  // {"en":"The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this CreateServiceService requires one. If this field is specified when creating a CreateServiceService which does not need it, creation will fail. This field will be wiped when updating a CreateServiceService to no longer need it (e.g. changing type from NodePort to ClusterIP).", "zh_CN":"当类型为 NodePort 或 LoadBalancer 时，CreateServiceService 公开在节点上的端口， 通常由系统分配。如果指定了一个在范围内且未使用的值，则将使用该值，否则操作将失败。 如果在创建的 CreateServiceService 需要该端口时未指定该字段，则会分配端口。 如果在创建不需要该端口的 Service时指定了该字段，则会创建失败。 当更新 CreateServiceService 时，如果不再需要此字段（例如，将类型从 NodePort 更改为 ClusterIP），这个字段将被擦除"}
  NodePort *int32 `json:"nodePort,omitempty" xml:"nodePort,omitempty"`
}

func (s CreateServiceServicePort) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceServicePort) GoString() string {
  return s.String()
}

func (s *CreateServiceServicePort) SetName(v string) *CreateServiceServicePort {
  s.Name = &v
  return s
}

func (s *CreateServiceServicePort) SetProtocol(v string) *CreateServiceServicePort {
  s.Protocol = &v
  return s
}

func (s *CreateServiceServicePort) SetAppProtocol(v string) *CreateServiceServicePort {
  s.AppProtocol = &v
  return s
}

func (s *CreateServiceServicePort) SetPort(v int32) *CreateServiceServicePort {
  s.Port = &v
  return s
}

func (s *CreateServiceServicePort) SetTargetPort(v int32) *CreateServiceServicePort {
  s.TargetPort = &v
  return s
}

func (s *CreateServiceServicePort) SetNodePort(v int32) *CreateServiceServicePort {
  s.NodePort = &v
  return s
}

type CreateServiceSessionAffinityConfig struct {
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"clientIP 包含基于客户端 IP 的会话亲和性的配置"}
  ClientIP *CreateServiceClientIPConfig `json:"clientIP,omitempty" xml:"clientIP,omitempty"`
}

func (s CreateServiceSessionAffinityConfig) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceSessionAffinityConfig) GoString() string {
  return s.String()
}

func (s *CreateServiceSessionAffinityConfig) SetClientIP(v *CreateServiceClientIPConfig) *CreateServiceSessionAffinityConfig {
  s.ClientIP = v
  return s
}

type CreateServiceClientIPConfig struct {
  // {"en":"timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).", "zh_CN":"timeoutSeconds 指定 ClientIP 类型会话的维系时间秒数。 如果 ServiceAffinity == 'ClientIP'，则该值必须 >0 && <=86400（1 天）。默认值为 10800（3 小时）"}
  TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s CreateServiceClientIPConfig) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceClientIPConfig) GoString() string {
  return s.String()
}

func (s *CreateServiceClientIPConfig) SetTimeoutSeconds(v int32) *CreateServiceClientIPConfig {
  s.TimeoutSeconds = &v
  return s
}

type CreateServiceObjectMeta struct {
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
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 CreateServiceService 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*CreateServiceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*CreateServiceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s CreateServiceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceObjectMeta) GoString() string {
  return s.String()
}

func (s *CreateServiceObjectMeta) SetName(v string) *CreateServiceObjectMeta {
  s.Name = &v
  return s
}

func (s *CreateServiceObjectMeta) SetGenerateName(v string) *CreateServiceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *CreateServiceObjectMeta) SetNamespace(v string) *CreateServiceObjectMeta {
  s.Namespace = &v
  return s
}

func (s *CreateServiceObjectMeta) SetSelfLink(v string) *CreateServiceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *CreateServiceObjectMeta) SetUid(v string) *CreateServiceObjectMeta {
  s.Uid = &v
  return s
}

func (s *CreateServiceObjectMeta) SetResourceVersion(v string) *CreateServiceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *CreateServiceObjectMeta) SetGeneration(v int64) *CreateServiceObjectMeta {
  s.Generation = &v
  return s
}

func (s *CreateServiceObjectMeta) SetCreationTimestamp(v string) *CreateServiceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *CreateServiceObjectMeta) SetDeletionTimestamp(v string) *CreateServiceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *CreateServiceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *CreateServiceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *CreateServiceObjectMeta) SetLabels(v map[string]*string) *CreateServiceObjectMeta {
  s.Labels = v
  return s
}

func (s *CreateServiceObjectMeta) SetAnnotations(v map[string]*string) *CreateServiceObjectMeta {
  s.Annotations = v
  return s
}

func (s *CreateServiceObjectMeta) SetOwnerReferences(v []*CreateServiceOwnerReference) *CreateServiceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *CreateServiceObjectMeta) SetFinalizers(v []*string) *CreateServiceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *CreateServiceObjectMeta) SetClusterName(v string) *CreateServiceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *CreateServiceObjectMeta) SetManagedFields(v []*CreateServiceManagedFieldsEntry) *CreateServiceObjectMeta {
  s.ManagedFields = v
  return s
}

type CreateServiceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this CreateServiceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'CreateServiceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“CreateServiceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"CreateServiceFieldsV1 holds the first JSON version format as described in the 'CreateServiceFieldsV1' type", "zh_CN":"CreateServiceFieldsV1 包含类型 “CreateServiceFieldsV1” 中描述的第一个 JSON 版本格式"}
  CreateServiceFieldsV1 *CreateServiceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s CreateServiceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *CreateServiceManagedFieldsEntry) SetManager(v string) *CreateServiceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *CreateServiceManagedFieldsEntry) SetOperation(v string) *CreateServiceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *CreateServiceManagedFieldsEntry) SetApiVersion(v string) *CreateServiceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *CreateServiceManagedFieldsEntry) SetTime(v string) *CreateServiceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *CreateServiceManagedFieldsEntry) SetFieldsType(v string) *CreateServiceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *CreateServiceManagedFieldsEntry) SetFieldsV1(v *CreateServiceFieldsV1) *CreateServiceManagedFieldsEntry {
  s.CreateServiceFieldsV1 = v
  return s
}

func (s *CreateServiceManagedFieldsEntry) SetSubresource(v string) *CreateServiceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type CreateServiceFieldsV1 struct {
}

func (s CreateServiceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceFieldsV1) GoString() string {
  return s.String()
}

type CreateServiceOwnerReference struct {
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

func (s CreateServiceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s CreateServiceOwnerReference) GoString() string {
  return s.String()
}

func (s *CreateServiceOwnerReference) SetApiVersion(v string) *CreateServiceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *CreateServiceOwnerReference) SetKind(v string) *CreateServiceOwnerReference {
  s.Kind = &v
  return s
}

func (s *CreateServiceOwnerReference) SetName(v string) *CreateServiceOwnerReference {
  s.Name = &v
  return s
}

func (s *CreateServiceOwnerReference) SetUid(v string) *CreateServiceOwnerReference {
  s.Uid = &v
  return s
}

func (s *CreateServiceOwnerReference) SetController(v bool) *CreateServiceOwnerReference {
  s.Controller = &v
  return s
}

func (s *CreateServiceOwnerReference) SetBlockOwnerDeletion(v bool) *CreateServiceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type VMPAllocateEdgeIPRequest struct {
  // {"en":"node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"IP protocol:4-ipv4(default);6-ipv6(temporary unsupported)","zh_CN":"可选\nIP协议：4-ipv4(默认)；6-ipv6"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"IP native attribute, 1: non-native;-1: native;","zh_CN":"IPv4原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"cidr","zh_CN":"CIDR，一次只能传入一个CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"number of applications IP  (the single upper limit is 50)","zh_CN":"申请IP数（单次申请Ip数上限为50个）"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Allocate IP randomly","zh_CN":"是否需要随机分配IP（仅对ipv4生效）\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
}

func (s VMPAllocateEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPAllocateEdgeIPRequest) SetNodeName(v string) *VMPAllocateEdgeIPRequest {
  s.NodeName = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetProtocol(v string) *VMPAllocateEdgeIPRequest {
  s.Protocol = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetNativeAttribute(v string) *VMPAllocateEdgeIPRequest {
  s.NativeAttribute = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetCidr(v string) *VMPAllocateEdgeIPRequest {
  s.Cidr = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetCount(v int) *VMPAllocateEdgeIPRequest {
  s.Count = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetRandomAllocateIp(v int) *VMPAllocateEdgeIPRequest {
  s.RandomAllocateIp = &v
  return s
}

type VMPAllocateEdgeIPRequestHeader struct {
}

func (s VMPAllocateEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPAllocateEdgeIPPaths struct {
}

func (s VMPAllocateEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPAllocateEdgeIPParameters struct {
}

func (s VMPAllocateEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPAllocateEdgeIPResponse struct {
  // {"en":"successful application for all or part of IP","zh_CN":"成功申请到的全部或部分IP\n说明：不同场景的响应说明如下\nA、所有IP都申请成功，返回申请到的所有IP\nB、只申请到部分IP，返回申请到的那部分IP\nC、未申请到任何IP，返回失败信息\nD、若出现申请失败的情况，请间隔10S之后再次申请"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPAllocateEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPAllocateEdgeIPResponse) SetEdgeIps(v []*string) *VMPAllocateEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPAllocateEdgeIPResponseHeader struct {
}

func (s VMPAllocateEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type AssignEdgePrivateIPRequest struct {
  // {"en":"target virtual machine id","zh_CN":"目标实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"additional IP to bind to the virtual machine","zh_CN":"要绑定到目标实例的额外IP"}
  EdgeIps []*AssignEdgePrivateIPRequestEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s AssignEdgePrivateIPRequest) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPRequest) GoString() string {
  return s.String()
}

func (s *AssignEdgePrivateIPRequest) SetInstanceId(v string) *AssignEdgePrivateIPRequest {
  s.InstanceId = &v
  return s
}

func (s *AssignEdgePrivateIPRequest) SetEdgeIps(v []*AssignEdgePrivateIPRequestEdgeIps) *AssignEdgePrivateIPRequest {
  s.EdgeIps = v
  return s
}

type AssignEdgePrivateIPRequestEdgeIps struct     {
  // {"en":"edge private ip","zh_CN":"额外内网IP"}
  EdgePrivateIp *string `json:"edgePrivateIp,omitempty" xml:"edgePrivateIp,omitempty" require:"true"`
  // {"en":"edge public IP","zh_CN":"额外公网IP"}
  EdgePublicIp *string `json:"edgePublicIp,omitempty" xml:"edgePublicIp,omitempty" require:"true"`
}

func (s AssignEdgePrivateIPRequestEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPRequestEdgeIps) GoString() string {
  return s.String()
}

func (s *AssignEdgePrivateIPRequestEdgeIps) SetEdgePrivateIp(v string) *AssignEdgePrivateIPRequestEdgeIps {
  s.EdgePrivateIp = &v
  return s
}

func (s *AssignEdgePrivateIPRequestEdgeIps) SetEdgePublicIp(v string) *AssignEdgePrivateIPRequestEdgeIps {
  s.EdgePublicIp = &v
  return s
}

type AssignEdgePrivateIPRequestHeader struct {
}

func (s AssignEdgePrivateIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPRequestHeader) GoString() string {
  return s.String()
}

type AssignEdgePrivateIPPaths struct {
}

func (s AssignEdgePrivateIPPaths) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPPaths) GoString() string {
  return s.String()
}

type AssignEdgePrivateIPParameters struct {
}

func (s AssignEdgePrivateIPParameters) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPParameters) GoString() string {
  return s.String()
}

type AssignEdgePrivateIPResponse struct {
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *AssignEdgePrivateIPResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AssignEdgePrivateIPResponse) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPResponse) GoString() string {
  return s.String()
}

func (s *AssignEdgePrivateIPResponse) SetCode(v string) *AssignEdgePrivateIPResponse {
  s.Code = &v
  return s
}

func (s *AssignEdgePrivateIPResponse) SetMessage(v string) *AssignEdgePrivateIPResponse {
  s.Message = &v
  return s
}

func (s *AssignEdgePrivateIPResponse) SetData(v *AssignEdgePrivateIPResponseData) *AssignEdgePrivateIPResponse {
  s.Data = v
  return s
}

type AssignEdgePrivateIPResponseData struct {
  // {"en":"target virtual machine id","zh_CN":"目标实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"Additional IP that is bound to the instance","zh_CN":"已绑定到目标实例的额外IP"}
  EdgeIps []*AssignEdgePrivateIPResponseDataEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s AssignEdgePrivateIPResponseData) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPResponseData) GoString() string {
  return s.String()
}

func (s *AssignEdgePrivateIPResponseData) SetInstanceId(v string) *AssignEdgePrivateIPResponseData {
  s.InstanceId = &v
  return s
}

func (s *AssignEdgePrivateIPResponseData) SetEdgeIps(v []*AssignEdgePrivateIPResponseDataEdgeIps) *AssignEdgePrivateIPResponseData {
  s.EdgeIps = v
  return s
}

type AssignEdgePrivateIPResponseDataEdgeIps struct     {
  // {"en":"edge private ip","zh_CN":"额外内网IP"}
  EdgePrivateIp *string `json:"edgePrivateIp,omitempty" xml:"edgePrivateIp,omitempty" require:"true"`
  // {"en":"edge public IP","zh_CN":"额外公网IP"}
  EdgePublicIp *string `json:"edgePublicIp,omitempty" xml:"edgePublicIp,omitempty" require:"true"`
}

func (s AssignEdgePrivateIPResponseDataEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPResponseDataEdgeIps) GoString() string {
  return s.String()
}

func (s *AssignEdgePrivateIPResponseDataEdgeIps) SetEdgePrivateIp(v string) *AssignEdgePrivateIPResponseDataEdgeIps {
  s.EdgePrivateIp = &v
  return s
}

func (s *AssignEdgePrivateIPResponseDataEdgeIps) SetEdgePublicIp(v string) *AssignEdgePrivateIPResponseDataEdgeIps {
  s.EdgePublicIp = &v
  return s
}

type AssignEdgePrivateIPResponseHeader struct {
}

func (s AssignEdgePrivateIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssignEdgePrivateIPResponseHeader) GoString() string {
  return s.String()
}




type ListNetworkPolicyRequest struct {
}

func (s ListNetworkPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyRequest) GoString() string {
  return s.String()
}

type ListNetworkPolicyResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"networkPolicy", "zh_CN":"网络策略"}
  Data *ListNetworkPolicyNetworkPolicyList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListNetworkPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyResponse) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyResponse) SetCode(v int64) *ListNetworkPolicyResponse {
  s.Code = &v
  return s
}

func (s *ListNetworkPolicyResponse) SetMsg(v string) *ListNetworkPolicyResponse {
  s.Msg = &v
  return s
}

func (s *ListNetworkPolicyResponse) SetRequestId(v string) *ListNetworkPolicyResponse {
  s.RequestId = &v
  return s
}

func (s *ListNetworkPolicyResponse) SetData(v *ListNetworkPolicyNetworkPolicyList) *ListNetworkPolicyResponse {
  s.Data = v
  return s
}

type ListNetworkPolicyPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s ListNetworkPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyPaths) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyPaths) SetNamespace(v string) *ListNetworkPolicyPaths {
  s.Namespace = &v
  return s
}

type ListNetworkPolicyParameters struct {
  // {"en":"networkPolicy name", "zh_CN":"网络策略名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s ListNetworkPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyParameters) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyParameters) SetName(v string) *ListNetworkPolicyParameters {
  s.Name = &v
  return s
}

type ListNetworkPolicyRequestHeader struct {
}

func (s ListNetworkPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyRequestHeader) GoString() string {
  return s.String()
}

type ListNetworkPolicyResponseHeader struct {
}

func (s ListNetworkPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyResponseHeader) GoString() string {
  return s.String()
}

type ListNetworkPolicyNetworkPolicyList struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *ListNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "list of networkPolicy", "zh_CN": "网络策略列表"}
  Items []*ListNetworkPolicyNetworkPolicy `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s ListNetworkPolicyNetworkPolicyList) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicyList) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicyList) SetApiVersion(v string) *ListNetworkPolicyNetworkPolicyList {
  s.ApiVersion = &v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyList) SetKind(v string) *ListNetworkPolicyNetworkPolicyList {
  s.Kind = &v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyList) SetMetadata(v *ListNetworkPolicyObjectMeta) *ListNetworkPolicyNetworkPolicyList {
  s.Metadata = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyList) SetItems(v []*ListNetworkPolicyNetworkPolicy) *ListNetworkPolicyNetworkPolicyList {
  s.Items = v
  return s
}

type ListNetworkPolicyNetworkPolicy struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *ListNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this ListNetworkPolicyNetworkPolicy.", "zh_CN": "ListNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *ListNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s ListNetworkPolicyNetworkPolicy) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicy) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicy) SetApiVersion(v string) *ListNetworkPolicyNetworkPolicy {
  s.ApiVersion = &v
  return s
}

func (s *ListNetworkPolicyNetworkPolicy) SetKind(v string) *ListNetworkPolicyNetworkPolicy {
  s.Kind = &v
  return s
}

func (s *ListNetworkPolicyNetworkPolicy) SetMetadata(v *ListNetworkPolicyObjectMeta) *ListNetworkPolicyNetworkPolicy {
  s.Metadata = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicy) SetSpec(v *ListNetworkPolicyNetworkPolicySpec) *ListNetworkPolicyNetworkPolicy {
  s.Spec = v
  return s
}

type ListNetworkPolicyNetworkPolicySpec struct {
  // {"en": "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the ListNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this ListNetworkPolicyNetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8", "zh_CN": "出网规则"}
  Egress []*ListNetworkPolicyNetworkPolicyEgressRule `json:"egress,omitempty" xml:"egress,omitempty" type:"Repeated"`
  // {"en": "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the ListNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this ListNetworkPolicyNetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)", "zh_CN": "入网规则"}
  Ingress []*ListNetworkPolicyNetworkPolicyIngressRule `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
  // {"en": "Selects the pods to which this ListNetworkPolicyNetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.", "zh_CN": "限制pod的选择器"}
  PodSelector *ListNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty" require:"true"`
  // {"en": "List of rule types that the ListNetworkPolicyNetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ Egress ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include Egress (since such a policy would not include an Egress section and would otherwise default to just [ Ingress ]). This field is beta-level in 1.8", "zh_CN": "策略类型"}
  PolicyTypes []*string `json:"policyTypes,omitempty" xml:"policyTypes,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyNetworkPolicySpec) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicySpec) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicySpec) SetEgress(v []*ListNetworkPolicyNetworkPolicyEgressRule) *ListNetworkPolicyNetworkPolicySpec {
  s.Egress = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicySpec) SetIngress(v []*ListNetworkPolicyNetworkPolicyIngressRule) *ListNetworkPolicyNetworkPolicySpec {
  s.Ingress = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicySpec) SetPodSelector(v *ListNetworkPolicyPodLabelSelector) *ListNetworkPolicyNetworkPolicySpec {
  s.PodSelector = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicySpec) SetPolicyTypes(v []*string) *ListNetworkPolicyNetworkPolicySpec {
  s.PolicyTypes = v
  return s
}

type ListNetworkPolicyNetworkPolicyIngressRule struct {
  // {"en": "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.", "zh_CN": "入网规则信息"}
  From []*ListNetworkPolicyNetworkPolicyPeer `json:"from,omitempty" xml:"from,omitempty" type:"Repeated"`
  // {"en": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*ListNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyNetworkPolicyIngressRule) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicyIngressRule) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicyIngressRule) SetFrom(v []*ListNetworkPolicyNetworkPolicyPeer) *ListNetworkPolicyNetworkPolicyIngressRule {
  s.From = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyIngressRule) SetPorts(v []*ListNetworkPolicyNetworkPolicyPort) *ListNetworkPolicyNetworkPolicyIngressRule {
  s.Ports = v
  return s
}

type ListNetworkPolicyNetworkPolicyEgressRule struct {
  // {"en": "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*ListNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en": "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.", "zh_CN": "出网规则信息"}
  To []*ListNetworkPolicyNetworkPolicyPeer `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyNetworkPolicyEgressRule) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicyEgressRule) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicyEgressRule) SetPorts(v []*ListNetworkPolicyNetworkPolicyPort) *ListNetworkPolicyNetworkPolicyEgressRule {
  s.Ports = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyEgressRule) SetTo(v []*ListNetworkPolicyNetworkPolicyPeer) *ListNetworkPolicyNetworkPolicyEgressRule {
  s.To = v
  return s
}

type ListNetworkPolicyNetworkPolicyPeer struct {
  // {"en": "ListNetworkPolicyIPBlock defines policy on a particular ListNetworkPolicyIPBlock. If this field is set then neither of the other fields can be.", "zh_CN": "IP规则"}
  IpBlock *ListNetworkPolicyIPBlock `json:"ipBlock,omitempty" xml:"ipBlock,omitempty"`
  // {"en": "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.If PodSelector is also set, then the ListNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.", "zh_CN": "namespace选择器"}
  NamespaceSelector *ListNetworkPolicyNsLabelSelector `json:"namespaceSelector,omitempty" xml:"namespaceSelector,omitempty"`
  // {"en": "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.If NamespaceSelector is also set, then the ListNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.", "zh_CN": "pod选择器"}
  PodSelector *ListNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty"`
}

func (s ListNetworkPolicyNetworkPolicyPeer) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicyPeer) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicyPeer) SetIpBlock(v *ListNetworkPolicyIPBlock) *ListNetworkPolicyNetworkPolicyPeer {
  s.IpBlock = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyPeer) SetNamespaceSelector(v *ListNetworkPolicyNsLabelSelector) *ListNetworkPolicyNetworkPolicyPeer {
  s.NamespaceSelector = v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyPeer) SetPodSelector(v *ListNetworkPolicyPodLabelSelector) *ListNetworkPolicyNetworkPolicyPeer {
  s.PodSelector = v
  return s
}

type ListNetworkPolicyIPBlock struct {
  // {"en": "CIDR is a string representing the IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64", "zh_CN": "生效IP网段"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en": "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64 Except values will be rejected if they are outside the CIDR range", "zh_CN": "例外IP网段"}
  Except []*string `json:"except,omitempty" xml:"except,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyIPBlock) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyIPBlock) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyIPBlock) SetCidr(v string) *ListNetworkPolicyIPBlock {
  s.Cidr = &v
  return s
}

func (s *ListNetworkPolicyIPBlock) SetExcept(v []*string) *ListNetworkPolicyIPBlock {
  s.Except = v
  return s
}

type ListNetworkPolicyNetworkPolicyPort struct {
  // {"en": "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.", "zh_CN": "端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
  // {"en": "The protocol (TCP, UDP) which traffic must match. If not specified, this field defaults to TCP.", "zh_CN": "协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ListNetworkPolicyNetworkPolicyPort) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNetworkPolicyPort) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNetworkPolicyPort) SetPort(v string) *ListNetworkPolicyNetworkPolicyPort {
  s.Port = &v
  return s
}

func (s *ListNetworkPolicyNetworkPolicyPort) SetProtocol(v string) *ListNetworkPolicyNetworkPolicyPort {
  s.Protocol = &v
  return s
}

type ListNetworkPolicyPodLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
}

func (s ListNetworkPolicyPodLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyPodLabelSelector) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyPodLabelSelector) SetMatchLabels(v map[string]*string) *ListNetworkPolicyPodLabelSelector {
  s.MatchLabels = v
  return s
}

type ListNetworkPolicyNsLabelSelector struct {
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*ListNetworkPolicyLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyNsLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyNsLabelSelector) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyNsLabelSelector) SetMatchExpressions(v []*ListNetworkPolicyLabelSelectorRequirement) *ListNetworkPolicyNsLabelSelector {
  s.MatchExpressions = v
  return s
}

type ListNetworkPolicyLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyLabelSelectorRequirement) SetKey(v string) *ListNetworkPolicyLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *ListNetworkPolicyLabelSelectorRequirement) SetOperator(v string) *ListNetworkPolicyLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *ListNetworkPolicyLabelSelectorRequirement) SetValues(v []*string) *ListNetworkPolicyLabelSelectorRequirement {
  s.Values = v
  return s
}

type ListNetworkPolicyObjectMeta struct {
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
  OwnerReferences []*ListNetworkPolicyOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*ListNetworkPolicyManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s ListNetworkPolicyObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyObjectMeta) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyObjectMeta) SetName(v string) *ListNetworkPolicyObjectMeta {
  s.Name = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetGenerateName(v string) *ListNetworkPolicyObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetNamespace(v string) *ListNetworkPolicyObjectMeta {
  s.Namespace = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetSelfLink(v string) *ListNetworkPolicyObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetUid(v string) *ListNetworkPolicyObjectMeta {
  s.Uid = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetResourceVersion(v string) *ListNetworkPolicyObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetGeneration(v int64) *ListNetworkPolicyObjectMeta {
  s.Generation = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetCreationTimestamp(v string) *ListNetworkPolicyObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetDeletionTimestamp(v string) *ListNetworkPolicyObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetDeletionGracePeriodSeconds(v int64) *ListNetworkPolicyObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetLabels(v map[string]*string) *ListNetworkPolicyObjectMeta {
  s.Labels = v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetAnnotations(v map[string]*string) *ListNetworkPolicyObjectMeta {
  s.Annotations = v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetOwnerReferences(v []*ListNetworkPolicyOwnerReference) *ListNetworkPolicyObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetFinalizers(v []*string) *ListNetworkPolicyObjectMeta {
  s.Finalizers = v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetClusterName(v string) *ListNetworkPolicyObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *ListNetworkPolicyObjectMeta) SetManagedFields(v []*ListNetworkPolicyManagedFieldsEntry) *ListNetworkPolicyObjectMeta {
  s.ManagedFields = v
  return s
}

type ListNetworkPolicyOwnerReference struct {
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

func (s ListNetworkPolicyOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyOwnerReference) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyOwnerReference) SetApiVersion(v string) *ListNetworkPolicyOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *ListNetworkPolicyOwnerReference) SetKind(v string) *ListNetworkPolicyOwnerReference {
  s.Kind = &v
  return s
}

func (s *ListNetworkPolicyOwnerReference) SetName(v string) *ListNetworkPolicyOwnerReference {
  s.Name = &v
  return s
}

func (s *ListNetworkPolicyOwnerReference) SetUid(v string) *ListNetworkPolicyOwnerReference {
  s.Uid = &v
  return s
}

func (s *ListNetworkPolicyOwnerReference) SetController(v bool) *ListNetworkPolicyOwnerReference {
  s.Controller = &v
  return s
}

func (s *ListNetworkPolicyOwnerReference) SetBlockOwnerDeletion(v bool) *ListNetworkPolicyOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type ListNetworkPolicyManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this ListNetworkPolicyManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'ListNetworkPolicyFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“ListNetworkPolicyFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"ListNetworkPolicyFieldsV1 holds the first JSON version format as described in the 'ListNetworkPolicyFieldsV1' type", "zh_CN":"ListNetworkPolicyFieldsV1 包含类型 “ListNetworkPolicyFieldsV1” 中描述的第一个 JSON 版本格式"}
  ListNetworkPolicyFieldsV1 *ListNetworkPolicyFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s ListNetworkPolicyManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetManager(v string) *ListNetworkPolicyManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetOperation(v string) *ListNetworkPolicyManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetApiVersion(v string) *ListNetworkPolicyManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetTime(v string) *ListNetworkPolicyManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetFieldsType(v string) *ListNetworkPolicyManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetFieldsV1(v *ListNetworkPolicyFieldsV1) *ListNetworkPolicyManagedFieldsEntry {
  s.ListNetworkPolicyFieldsV1 = v
  return s
}

func (s *ListNetworkPolicyManagedFieldsEntry) SetSubresource(v string) *ListNetworkPolicyManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type ListNetworkPolicyFieldsV1 struct {
}

func (s ListNetworkPolicyFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s ListNetworkPolicyFieldsV1) GoString() string {
  return s.String()
}




type LECHAssignEdgeIPRequest struct {
  // {"en":"target virtual machine external network Ip","zh_CN":"目标实例的公网Ip"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional IP to bind to the virtual machine","zh_CN":"要绑定到目标实例的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHAssignEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHAssignEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *LECHAssignEdgeIPRequest) SetServerIp(v string) *LECHAssignEdgeIPRequest {
  s.ServerIp = &v
  return s
}

func (s *LECHAssignEdgeIPRequest) SetEdgeIps(v []*string) *LECHAssignEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type LECHAssignEdgeIPRequestHeader struct {
}

func (s LECHAssignEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAssignEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type LECHAssignEdgeIPPaths struct {
}

func (s LECHAssignEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHAssignEdgeIPPaths) GoString() string {
  return s.String()
}

type LECHAssignEdgeIPParameters struct {
}

func (s LECHAssignEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHAssignEdgeIPParameters) GoString() string {
  return s.String()
}

type LECHAssignEdgeIPResponse struct {
  // {"en":"target virtual machine IP","zh_CN":"目标实例公网IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"Additional IP that is bound to the instance","zh_CN":"已绑定到实例的额外公网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHAssignEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHAssignEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *LECHAssignEdgeIPResponse) SetServerIp(v string) *LECHAssignEdgeIPResponse {
  s.ServerIp = &v
  return s
}

func (s *LECHAssignEdgeIPResponse) SetEdgeIps(v []*string) *LECHAssignEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type LECHAssignEdgeIPResponseHeader struct {
}

func (s LECHAssignEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAssignEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryEdgeIPRequest struct {
}

func (s VMPQueryEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPRequest) GoString() string {
  return s.String()
}

type VMPQueryEdgeIPRequestHeader struct {
}

func (s VMPQueryEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryEdgeIPPaths struct {
}

func (s VMPQueryEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPQueryEdgeIPParameters struct {
  // {"en":"node name","zh_CN":"可选\n节点名称，多个用英文逗号分隔"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"virtual machine ID","zh_CN":"可选\n虚拟机ID，多个用英文逗号分隔"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty"`
  // {"en":"virtual machine master IP","zh_CN":"可选\n虚拟机主IP，多个用英文逗号分隔"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty"`
  // {"en":"extra Ip","zh_CN":"可选\n额外Ip，多个用英文逗号分隔"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty"`
  // {"en":"IP state","zh_CN":"可选\nIP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s VMPQueryEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryEdgeIPParameters) SetNodeName(v string) *VMPQueryEdgeIPParameters {
  s.NodeName = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetServerId(v string) *VMPQueryEdgeIPParameters {
  s.ServerId = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetServerIp(v string) *VMPQueryEdgeIPParameters {
  s.ServerIp = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetEdgeIp(v string) *VMPQueryEdgeIPParameters {
  s.EdgeIp = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetState(v string) *VMPQueryEdgeIPParameters {
  s.State = &v
  return s
}

type VMPQueryEdgeIPResponse struct {
  // {"en":"additional IP details","zh_CN":"额外Ip详细信息"}
  EdgeIps []*VMPQueryEdgeIPResponseEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryEdgeIPResponse) SetEdgeIps(v []*VMPQueryEdgeIPResponseEdgeIps) *VMPQueryEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPQueryEdgeIPResponseEdgeIps struct     {
  // {"en":"extra Ip","zh_CN":"额外Ip"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty" require:"true"`
  // {"en":"IP state","zh_CN":"IP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"binding virtual machine ID","zh_CN":"绑定的虚拟机ID"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"binding virtual machine extranet IP","zh_CN":"绑定的虚拟机外网IP"}
  ServerIp []*string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true" type:"Repeated"`
  // {"en":"node name","zh_CN":"所属节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Is it exclusive","zh_CN":"是否独占"}
  OccupancyFlag *bool `json:"occupancyFlag,omitempty" xml:"occupancyFlag,omitempty" require:"true"`
  // {"en":"Net Mask","zh_CN":"子网掩码"}
  Netmask *string `json:"netmask,omitempty" xml:"netmask,omitempty" require:"true"`
}

func (s VMPQueryEdgeIPResponseEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPResponseEdgeIps) GoString() string {
  return s.String()
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetEdgeIp(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.EdgeIp = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetState(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.State = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetServerId(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.ServerId = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetServerIp(v []*string) *VMPQueryEdgeIPResponseEdgeIps {
  s.ServerIp = v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetNodeName(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.NodeName = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetOccupancyFlag(v bool) *VMPQueryEdgeIPResponseEdgeIps {
  s.OccupancyFlag = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetNetmask(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.Netmask = &v
  return s
}

type VMPQueryEdgeIPResponseHeader struct {
}

func (s VMPQueryEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type EdgeIPrivatepAllocateRequest struct {
  // {"en":"node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"number of applications IP  (the single upper limit is 50)","zh_CN":"申请IP数（单次申请Ip数上限为50个）"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s EdgeIPrivatepAllocateRequest) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocateRequest) GoString() string {
  return s.String()
}

func (s *EdgeIPrivatepAllocateRequest) SetNodeName(v string) *EdgeIPrivatepAllocateRequest {
  s.NodeName = &v
  return s
}

func (s *EdgeIPrivatepAllocateRequest) SetCount(v int) *EdgeIPrivatepAllocateRequest {
  s.Count = &v
  return s
}

type EdgeIPrivatepAllocateRequestHeader struct {
}

func (s EdgeIPrivatepAllocateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocateRequestHeader) GoString() string {
  return s.String()
}

type EdgeIPrivatepAllocatePaths struct {
}

func (s EdgeIPrivatepAllocatePaths) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocatePaths) GoString() string {
  return s.String()
}

type EdgeIPrivatepAllocateParameters struct {
}

func (s EdgeIPrivatepAllocateParameters) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocateParameters) GoString() string {
  return s.String()
}

type EdgeIPrivatepAllocateResponse struct {
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *EdgeIPrivatepAllocateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s EdgeIPrivatepAllocateResponse) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocateResponse) GoString() string {
  return s.String()
}

func (s *EdgeIPrivatepAllocateResponse) SetCode(v string) *EdgeIPrivatepAllocateResponse {
  s.Code = &v
  return s
}

func (s *EdgeIPrivatepAllocateResponse) SetMessage(v string) *EdgeIPrivatepAllocateResponse {
  s.Message = &v
  return s
}

func (s *EdgeIPrivatepAllocateResponse) SetData(v *EdgeIPrivatepAllocateResponseData) *EdgeIPrivatepAllocateResponse {
  s.Data = v
  return s
}

type EdgeIPrivatepAllocateResponseData struct {
  // {"en":"successful application for all or part of IP","zh_CN":"成功申请到的全部或部分IP说明：不同场景的响应说明如下A、所有IP都申请成功，返回申请到的所有IPB、只申请到部分IP，返回申请到的那部分IPC、未申请到任何IP，返回失败信息D、若出现申请失败的情况，请间隔10S之后再次申请"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s EdgeIPrivatepAllocateResponseData) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocateResponseData) GoString() string {
  return s.String()
}

func (s *EdgeIPrivatepAllocateResponseData) SetEdgeIps(v []*string) *EdgeIPrivatepAllocateResponseData {
  s.EdgeIps = v
  return s
}

type EdgeIPrivatepAllocateResponseHeader struct {
}

func (s EdgeIPrivatepAllocateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EdgeIPrivatepAllocateResponseHeader) GoString() string {
  return s.String()
}




type GetNetworkPolicyRequest struct {
}

func (s GetNetworkPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyRequest) GoString() string {
  return s.String()
}

type GetNetworkPolicyResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"networkPolicy", "zh_CN":"网络策略"}
  Data *GetNetworkPolicyNetworkPolicy `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetNetworkPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyResponse) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyResponse) SetCode(v int64) *GetNetworkPolicyResponse {
  s.Code = &v
  return s
}

func (s *GetNetworkPolicyResponse) SetMsg(v string) *GetNetworkPolicyResponse {
  s.Msg = &v
  return s
}

func (s *GetNetworkPolicyResponse) SetRequestId(v string) *GetNetworkPolicyResponse {
  s.RequestId = &v
  return s
}

func (s *GetNetworkPolicyResponse) SetData(v *GetNetworkPolicyNetworkPolicy) *GetNetworkPolicyResponse {
  s.Data = v
  return s
}

type GetNetworkPolicyPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"networkPolicy name", "zh_CN":"networkPolicy 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetNetworkPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyPaths) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyPaths) SetNamespace(v string) *GetNetworkPolicyPaths {
  s.Namespace = &v
  return s
}

func (s *GetNetworkPolicyPaths) SetName(v string) *GetNetworkPolicyPaths {
  s.Name = &v
  return s
}

type GetNetworkPolicyParameters struct {
}

func (s GetNetworkPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyParameters) GoString() string {
  return s.String()
}

type GetNetworkPolicyRequestHeader struct {
}

func (s GetNetworkPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyRequestHeader) GoString() string {
  return s.String()
}

type GetNetworkPolicyResponseHeader struct {
}

func (s GetNetworkPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyResponseHeader) GoString() string {
  return s.String()
}

type GetNetworkPolicyNetworkPolicy struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *GetNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this GetNetworkPolicyNetworkPolicy.", "zh_CN": "GetNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *GetNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s GetNetworkPolicyNetworkPolicy) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNetworkPolicy) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNetworkPolicy) SetApiVersion(v string) *GetNetworkPolicyNetworkPolicy {
  s.ApiVersion = &v
  return s
}

func (s *GetNetworkPolicyNetworkPolicy) SetKind(v string) *GetNetworkPolicyNetworkPolicy {
  s.Kind = &v
  return s
}

func (s *GetNetworkPolicyNetworkPolicy) SetMetadata(v *GetNetworkPolicyObjectMeta) *GetNetworkPolicyNetworkPolicy {
  s.Metadata = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicy) SetSpec(v *GetNetworkPolicyNetworkPolicySpec) *GetNetworkPolicyNetworkPolicy {
  s.Spec = v
  return s
}

type GetNetworkPolicyNetworkPolicySpec struct {
  // {"en": "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the GetNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this GetNetworkPolicyNetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8", "zh_CN": "出网规则"}
  Egress []*GetNetworkPolicyNetworkPolicyEgressRule `json:"egress,omitempty" xml:"egress,omitempty" type:"Repeated"`
  // {"en": "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the GetNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this GetNetworkPolicyNetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)", "zh_CN": "入网规则"}
  Ingress []*GetNetworkPolicyNetworkPolicyIngressRule `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
  // {"en": "Selects the pods to which this GetNetworkPolicyNetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.", "zh_CN": "限制pod的选择器"}
  PodSelector *GetNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty" require:"true"`
  // {"en": "List of rule types that the GetNetworkPolicyNetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ Egress ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include Egress (since such a policy would not include an Egress section and would otherwise default to just [ Ingress ]). This field is beta-level in 1.8", "zh_CN": "策略类型"}
  PolicyTypes []*string `json:"policyTypes,omitempty" xml:"policyTypes,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyNetworkPolicySpec) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNetworkPolicySpec) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNetworkPolicySpec) SetEgress(v []*GetNetworkPolicyNetworkPolicyEgressRule) *GetNetworkPolicyNetworkPolicySpec {
  s.Egress = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicySpec) SetIngress(v []*GetNetworkPolicyNetworkPolicyIngressRule) *GetNetworkPolicyNetworkPolicySpec {
  s.Ingress = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicySpec) SetPodSelector(v *GetNetworkPolicyPodLabelSelector) *GetNetworkPolicyNetworkPolicySpec {
  s.PodSelector = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicySpec) SetPolicyTypes(v []*string) *GetNetworkPolicyNetworkPolicySpec {
  s.PolicyTypes = v
  return s
}

type GetNetworkPolicyNetworkPolicyIngressRule struct {
  // {"en": "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.", "zh_CN": "入网规则信息"}
  From []*GetNetworkPolicyNetworkPolicyPeer `json:"from,omitempty" xml:"from,omitempty" type:"Repeated"`
  // {"en": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*GetNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyNetworkPolicyIngressRule) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNetworkPolicyIngressRule) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNetworkPolicyIngressRule) SetFrom(v []*GetNetworkPolicyNetworkPolicyPeer) *GetNetworkPolicyNetworkPolicyIngressRule {
  s.From = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicyIngressRule) SetPorts(v []*GetNetworkPolicyNetworkPolicyPort) *GetNetworkPolicyNetworkPolicyIngressRule {
  s.Ports = v
  return s
}

type GetNetworkPolicyNetworkPolicyEgressRule struct {
  // {"en": "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*GetNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en": "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.", "zh_CN": "出网规则信息"}
  To []*GetNetworkPolicyNetworkPolicyPeer `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyNetworkPolicyEgressRule) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNetworkPolicyEgressRule) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNetworkPolicyEgressRule) SetPorts(v []*GetNetworkPolicyNetworkPolicyPort) *GetNetworkPolicyNetworkPolicyEgressRule {
  s.Ports = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicyEgressRule) SetTo(v []*GetNetworkPolicyNetworkPolicyPeer) *GetNetworkPolicyNetworkPolicyEgressRule {
  s.To = v
  return s
}

type GetNetworkPolicyNetworkPolicyPeer struct {
  // {"en": "GetNetworkPolicyIPBlock defines policy on a particular GetNetworkPolicyIPBlock. If this field is set then neither of the other fields can be.", "zh_CN": "IP规则"}
  IpBlock *GetNetworkPolicyIPBlock `json:"ipBlock,omitempty" xml:"ipBlock,omitempty"`
  // {"en": "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.If PodSelector is also set, then the GetNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.", "zh_CN": "namespace选择器"}
  NamespaceSelector *GetNetworkPolicyNsLabelSelector `json:"namespaceSelector,omitempty" xml:"namespaceSelector,omitempty"`
  // {"en": "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.If NamespaceSelector is also set, then the GetNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.", "zh_CN": "pod选择器"}
  PodSelector *GetNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty"`
}

func (s GetNetworkPolicyNetworkPolicyPeer) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNetworkPolicyPeer) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNetworkPolicyPeer) SetIpBlock(v *GetNetworkPolicyIPBlock) *GetNetworkPolicyNetworkPolicyPeer {
  s.IpBlock = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicyPeer) SetNamespaceSelector(v *GetNetworkPolicyNsLabelSelector) *GetNetworkPolicyNetworkPolicyPeer {
  s.NamespaceSelector = v
  return s
}

func (s *GetNetworkPolicyNetworkPolicyPeer) SetPodSelector(v *GetNetworkPolicyPodLabelSelector) *GetNetworkPolicyNetworkPolicyPeer {
  s.PodSelector = v
  return s
}

type GetNetworkPolicyIPBlock struct {
  // {"en": "CIDR is a string representing the IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64", "zh_CN": "生效IP网段"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en": "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64 Except values will be rejected if they are outside the CIDR range", "zh_CN": "例外IP网段"}
  Except []*string `json:"except,omitempty" xml:"except,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyIPBlock) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyIPBlock) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyIPBlock) SetCidr(v string) *GetNetworkPolicyIPBlock {
  s.Cidr = &v
  return s
}

func (s *GetNetworkPolicyIPBlock) SetExcept(v []*string) *GetNetworkPolicyIPBlock {
  s.Except = v
  return s
}

type GetNetworkPolicyNetworkPolicyPort struct {
  // {"en": "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.", "zh_CN": "端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
  // {"en": "The protocol (TCP, UDP) which traffic must match. If not specified, this field defaults to TCP.", "zh_CN": "协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetNetworkPolicyNetworkPolicyPort) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNetworkPolicyPort) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNetworkPolicyPort) SetPort(v string) *GetNetworkPolicyNetworkPolicyPort {
  s.Port = &v
  return s
}

func (s *GetNetworkPolicyNetworkPolicyPort) SetProtocol(v string) *GetNetworkPolicyNetworkPolicyPort {
  s.Protocol = &v
  return s
}

type GetNetworkPolicyPodLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
}

func (s GetNetworkPolicyPodLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyPodLabelSelector) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyPodLabelSelector) SetMatchLabels(v map[string]*string) *GetNetworkPolicyPodLabelSelector {
  s.MatchLabels = v
  return s
}

type GetNetworkPolicyNsLabelSelector struct {
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*GetNetworkPolicyLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyNsLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyNsLabelSelector) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyNsLabelSelector) SetMatchExpressions(v []*GetNetworkPolicyLabelSelectorRequirement) *GetNetworkPolicyNsLabelSelector {
  s.MatchExpressions = v
  return s
}

type GetNetworkPolicyLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyLabelSelectorRequirement) SetKey(v string) *GetNetworkPolicyLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *GetNetworkPolicyLabelSelectorRequirement) SetOperator(v string) *GetNetworkPolicyLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *GetNetworkPolicyLabelSelectorRequirement) SetValues(v []*string) *GetNetworkPolicyLabelSelectorRequirement {
  s.Values = v
  return s
}

type GetNetworkPolicyObjectMeta struct {
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
  OwnerReferences []*GetNetworkPolicyOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*GetNetworkPolicyManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s GetNetworkPolicyObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyObjectMeta) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyObjectMeta) SetName(v string) *GetNetworkPolicyObjectMeta {
  s.Name = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetGenerateName(v string) *GetNetworkPolicyObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetNamespace(v string) *GetNetworkPolicyObjectMeta {
  s.Namespace = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetSelfLink(v string) *GetNetworkPolicyObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetUid(v string) *GetNetworkPolicyObjectMeta {
  s.Uid = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetResourceVersion(v string) *GetNetworkPolicyObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetGeneration(v int64) *GetNetworkPolicyObjectMeta {
  s.Generation = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetCreationTimestamp(v string) *GetNetworkPolicyObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetDeletionTimestamp(v string) *GetNetworkPolicyObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetDeletionGracePeriodSeconds(v int64) *GetNetworkPolicyObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetLabels(v map[string]*string) *GetNetworkPolicyObjectMeta {
  s.Labels = v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetAnnotations(v map[string]*string) *GetNetworkPolicyObjectMeta {
  s.Annotations = v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetOwnerReferences(v []*GetNetworkPolicyOwnerReference) *GetNetworkPolicyObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetFinalizers(v []*string) *GetNetworkPolicyObjectMeta {
  s.Finalizers = v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetClusterName(v string) *GetNetworkPolicyObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *GetNetworkPolicyObjectMeta) SetManagedFields(v []*GetNetworkPolicyManagedFieldsEntry) *GetNetworkPolicyObjectMeta {
  s.ManagedFields = v
  return s
}

type GetNetworkPolicyOwnerReference struct {
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

func (s GetNetworkPolicyOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyOwnerReference) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyOwnerReference) SetApiVersion(v string) *GetNetworkPolicyOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *GetNetworkPolicyOwnerReference) SetKind(v string) *GetNetworkPolicyOwnerReference {
  s.Kind = &v
  return s
}

func (s *GetNetworkPolicyOwnerReference) SetName(v string) *GetNetworkPolicyOwnerReference {
  s.Name = &v
  return s
}

func (s *GetNetworkPolicyOwnerReference) SetUid(v string) *GetNetworkPolicyOwnerReference {
  s.Uid = &v
  return s
}

func (s *GetNetworkPolicyOwnerReference) SetController(v bool) *GetNetworkPolicyOwnerReference {
  s.Controller = &v
  return s
}

func (s *GetNetworkPolicyOwnerReference) SetBlockOwnerDeletion(v bool) *GetNetworkPolicyOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type GetNetworkPolicyManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this GetNetworkPolicyManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'GetNetworkPolicyFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“GetNetworkPolicyFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"GetNetworkPolicyFieldsV1 holds the first JSON version format as described in the 'GetNetworkPolicyFieldsV1' type", "zh_CN":"GetNetworkPolicyFieldsV1 包含类型 “GetNetworkPolicyFieldsV1” 中描述的第一个 JSON 版本格式"}
  GetNetworkPolicyFieldsV1 *GetNetworkPolicyFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s GetNetworkPolicyManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetManager(v string) *GetNetworkPolicyManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetOperation(v string) *GetNetworkPolicyManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetApiVersion(v string) *GetNetworkPolicyManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetTime(v string) *GetNetworkPolicyManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetFieldsType(v string) *GetNetworkPolicyManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetFieldsV1(v *GetNetworkPolicyFieldsV1) *GetNetworkPolicyManagedFieldsEntry {
  s.GetNetworkPolicyFieldsV1 = v
  return s
}

func (s *GetNetworkPolicyManagedFieldsEntry) SetSubresource(v string) *GetNetworkPolicyManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type GetNetworkPolicyFieldsV1 struct {
}

func (s GetNetworkPolicyFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s GetNetworkPolicyFieldsV1) GoString() string {
  return s.String()
}




type LECHEdgeIpAllocate4OccupancyRequest struct {
  // {"en":"Instance IP to be bound","zh_CN":"额外IP要绑定的实例IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"IP protocol: 4-ipv4(default); 6-ipv6","zh_CN":"IP协议：4-ipv4（默认）；6-ipv6"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"IP native attribute, 1: non-native;-1: native;","zh_CN":"指定IPv4原生属性。可选值：1：非原生，-1：原生。不指定默认随机分配原生属性"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"CIDR","zh_CN":"指定CIDR申请IP"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"A. When an instance has multiple carrier IPs, you can specify an additional IP carrier. B. If not specified: For instances with a single carrier IP, the additional IP will use the same carrier as the instance. For instances with cross-ISP IPs, For the additional IP, one of the carrier BGP nodes will be selected; this field will not take effect.","zh_CN":"A、实例有多个运营商IP时，可指定额外IP运营商\nB、不指定时：\n实例单运营商IP，额外IP运营商一致\n实例多运营商IP，额外IP选其中一个运营商\nBGP节点该字段不生效。"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Allocate ip random","zh_CN":"IPv4是否随机分配\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
  // {"en":"Number of applications IP","zh_CN":"申请个数"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s LECHEdgeIpAllocate4OccupancyRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyRequest) GoString() string {
  return s.String()
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetServerIp(v string) *LECHEdgeIpAllocate4OccupancyRequest {
  s.ServerIp = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetProtocol(v string) *LECHEdgeIpAllocate4OccupancyRequest {
  s.Protocol = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetNativeAttribute(v string) *LECHEdgeIpAllocate4OccupancyRequest {
  s.NativeAttribute = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetCidr(v string) *LECHEdgeIpAllocate4OccupancyRequest {
  s.Cidr = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetCarrier(v string) *LECHEdgeIpAllocate4OccupancyRequest {
  s.Carrier = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetRandomAllocateIp(v int) *LECHEdgeIpAllocate4OccupancyRequest {
  s.RandomAllocateIp = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyRequest) SetCount(v int) *LECHEdgeIpAllocate4OccupancyRequest {
  s.Count = &v
  return s
}

type LECHEdgeIpAllocate4OccupancyRequestHeader struct {
}

func (s LECHEdgeIpAllocate4OccupancyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyRequestHeader) GoString() string {
  return s.String()
}

type LECHEdgeIpAllocate4OccupancyPaths struct {
}

func (s LECHEdgeIpAllocate4OccupancyPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyPaths) GoString() string {
  return s.String()
}

type LECHEdgeIpAllocate4OccupancyParameters struct {
}

func (s LECHEdgeIpAllocate4OccupancyParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyParameters) GoString() string {
  return s.String()
}

type LECHEdgeIpAllocate4OccupancyResponse struct {
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *LECHEdgeIpAllocate4OccupancyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s LECHEdgeIpAllocate4OccupancyResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyResponse) GoString() string {
  return s.String()
}

func (s *LECHEdgeIpAllocate4OccupancyResponse) SetMessage(v string) *LECHEdgeIpAllocate4OccupancyResponse {
  s.Message = &v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyResponse) SetData(v *LECHEdgeIpAllocate4OccupancyResponseData) *LECHEdgeIpAllocate4OccupancyResponse {
  s.Data = v
  return s
}

func (s *LECHEdgeIpAllocate4OccupancyResponse) SetCode(v string) *LECHEdgeIpAllocate4OccupancyResponse {
  s.Code = &v
  return s
}

type LECHEdgeIpAllocate4OccupancyResponseData struct {
  // {"en":"Successful application for all or part of IP","zh_CN":"成功申请到的全部或部分IP说明：不同场景的响应说明如下 A、所有IP都申请成功，返回申请到的所有IP B、只申请到部分IP，返回申请到的那部分IP C、未申请到任何IP，返回失败信息 D、若出现申请失败的情况，请间隔10S之后再次申请"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHEdgeIpAllocate4OccupancyResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyResponseData) GoString() string {
  return s.String()
}

func (s *LECHEdgeIpAllocate4OccupancyResponseData) SetEdgeIps(v []*string) *LECHEdgeIpAllocate4OccupancyResponseData {
  s.EdgeIps = v
  return s
}

type LECHEdgeIpAllocate4OccupancyResponseHeader struct {
}

func (s LECHEdgeIpAllocate4OccupancyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEdgeIpAllocate4OccupancyResponseHeader) GoString() string {
  return s.String()
}




type GetIngressControllerRequest struct {
}

func (s GetIngressControllerRequest) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerRequest) GoString() string {
  return s.String()
}

type GetIngressControllerResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress controller", "zh_CN":"路由控制器"}
  Data *GetIngressControllerIngressController `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetIngressControllerResponse) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerResponse) GoString() string {
  return s.String()
}

func (s *GetIngressControllerResponse) SetCode(v int64) *GetIngressControllerResponse {
  s.Code = &v
  return s
}

func (s *GetIngressControllerResponse) SetMsg(v string) *GetIngressControllerResponse {
  s.Msg = &v
  return s
}

func (s *GetIngressControllerResponse) SetRequestId(v string) *GetIngressControllerResponse {
  s.RequestId = &v
  return s
}

func (s *GetIngressControllerResponse) SetData(v *GetIngressControllerIngressController) *GetIngressControllerResponse {
  s.Data = v
  return s
}

type GetIngressControllerPaths struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetIngressControllerPaths) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerPaths) GoString() string {
  return s.String()
}

func (s *GetIngressControllerPaths) SetName(v string) *GetIngressControllerPaths {
  s.Name = &v
  return s
}

type GetIngressControllerParameters struct {
}

func (s GetIngressControllerParameters) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerParameters) GoString() string {
  return s.String()
}

type GetIngressControllerRequestHeader struct {
}

func (s GetIngressControllerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerRequestHeader) GoString() string {
  return s.String()
}

type GetIngressControllerResponseHeader struct {
}

func (s GetIngressControllerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerResponseHeader) GoString() string {
  return s.String()
}

type GetIngressControllerIngressController struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"resource limit", "zh_CN":"资源限制"}
  Limit map[string]*string `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"cluster and replicate", "zh_CN":"部署集群和副本数"}
  Clusters []*GetIngressControllerIngressCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
}

func (s GetIngressControllerIngressController) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerIngressController) GoString() string {
  return s.String()
}

func (s *GetIngressControllerIngressController) SetName(v string) *GetIngressControllerIngressController {
  s.Name = &v
  return s
}

func (s *GetIngressControllerIngressController) SetLimit(v map[string]*string) *GetIngressControllerIngressController {
  s.Limit = v
  return s
}

func (s *GetIngressControllerIngressController) SetClusters(v []*GetIngressControllerIngressCluster) *GetIngressControllerIngressController {
  s.Clusters = v
  return s
}

type GetIngressControllerIngressCluster struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"replicates", "zh_CN":"副本数"}
  Replicate *int32 `json:"replicate,omitempty" xml:"replicate,omitempty" require:"true"`
}

func (s GetIngressControllerIngressCluster) String() string {
  return tea.Prettify(s)
}

func (s GetIngressControllerIngressCluster) GoString() string {
  return s.String()
}

func (s *GetIngressControllerIngressCluster) SetName(v string) *GetIngressControllerIngressCluster {
  s.Name = &v
  return s
}

func (s *GetIngressControllerIngressCluster) SetReplicate(v int32) *GetIngressControllerIngressCluster {
  s.Replicate = &v
  return s
}




type LECHReleaseEdgeIPRequest struct {
  // {"en":"additional IP to be released","zh_CN":"要释放的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHReleaseEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *LECHReleaseEdgeIPRequest) SetEdgeIps(v []*string) *LECHReleaseEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type LECHReleaseEdgeIPRequestHeader struct {
}

func (s LECHReleaseEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type LECHReleaseEdgeIPPaths struct {
}

func (s LECHReleaseEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPPaths) GoString() string {
  return s.String()
}

type LECHReleaseEdgeIPParameters struct {
}

func (s LECHReleaseEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPParameters) GoString() string {
  return s.String()
}

type LECHReleaseEdgeIPResponse struct {
  // {"en":"Error Message Set","zh_CN":"错误信息集"}
  BatchErrorMsg []*LECHReleaseEdgeIPResponseBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s LECHReleaseEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *LECHReleaseEdgeIPResponse) SetBatchErrorMsg(v []*LECHReleaseEdgeIPResponseBatchErrorMsg) *LECHReleaseEdgeIPResponse {
  s.BatchErrorMsg = v
  return s
}

type LECHReleaseEdgeIPResponseBatchErrorMsg struct     {
  // {"en":"Ip","zh_CN":"Ip"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s LECHReleaseEdgeIPResponseBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPResponseBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *LECHReleaseEdgeIPResponseBatchErrorMsg) SetKey(v string) *LECHReleaseEdgeIPResponseBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *LECHReleaseEdgeIPResponseBatchErrorMsg) SetMsg(v string) *LECHReleaseEdgeIPResponseBatchErrorMsg {
  s.Msg = &v
  return s
}

type LECHReleaseEdgeIPResponseHeader struct {
}

func (s LECHReleaseEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHReleaseEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type VMPAssignEdgeIPRequest struct {
  // {"en":"target virtual machine external network Ip","zh_CN":"目标实例的公网Ip"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional IP to bind to the virtual machine","zh_CN":"要绑定到目标实例的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPAssignEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPAssignEdgeIPRequest) SetServerIp(v string) *VMPAssignEdgeIPRequest {
  s.ServerIp = &v
  return s
}

func (s *VMPAssignEdgeIPRequest) SetEdgeIps(v []*string) *VMPAssignEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type VMPAssignEdgeIPRequestHeader struct {
}

func (s VMPAssignEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPAssignEdgeIPPaths struct {
}

func (s VMPAssignEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPAssignEdgeIPParameters struct {
}

func (s VMPAssignEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPAssignEdgeIPResponse struct {
  // {"en":"target virtual machine IP","zh_CN":"目标实例公网IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"Additional IP that is bound to the instance","zh_CN":"已绑定到实例的额外公网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPAssignEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPAssignEdgeIPResponse) SetServerIp(v string) *VMPAssignEdgeIPResponse {
  s.ServerIp = &v
  return s
}

func (s *VMPAssignEdgeIPResponse) SetEdgeIps(v []*string) *VMPAssignEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPAssignEdgeIPResponseHeader struct {
}

func (s VMPAssignEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type LECHUnassignEdgeIPRequest struct {
  // {"en":"target virtual machine external network Ip","zh_CN":"目标实例外网Ip"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional Ip to unbind","zh_CN":"要解除绑定的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHUnassignEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHUnassignEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *LECHUnassignEdgeIPRequest) SetServerIp(v string) *LECHUnassignEdgeIPRequest {
  s.ServerIp = &v
  return s
}

func (s *LECHUnassignEdgeIPRequest) SetEdgeIps(v []*string) *LECHUnassignEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type LECHUnassignEdgeIPRequestHeader struct {
}

func (s LECHUnassignEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHUnassignEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type LECHUnassignEdgeIPPaths struct {
}

func (s LECHUnassignEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHUnassignEdgeIPPaths) GoString() string {
  return s.String()
}

type LECHUnassignEdgeIPParameters struct {
}

func (s LECHUnassignEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHUnassignEdgeIPParameters) GoString() string {
  return s.String()
}

type LECHUnassignEdgeIPResponse struct {
  // {"en":"target virtual machine IP","zh_CN":"目标实例公网IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional Ip that has been bound to the virtual machine","zh_CN":"已绑定到实例的额外公网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHUnassignEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHUnassignEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *LECHUnassignEdgeIPResponse) SetServerIp(v string) *LECHUnassignEdgeIPResponse {
  s.ServerIp = &v
  return s
}

func (s *LECHUnassignEdgeIPResponse) SetEdgeIps(v []*string) *LECHUnassignEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type LECHUnassignEdgeIPResponseHeader struct {
}

func (s LECHUnassignEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHUnassignEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type UpdateIngressRequest struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  ControllerName *string `json:"controllerName,omitempty" xml:"controllerName,omitempty" require:"true"`
  // {"en":"clusters", "zh_CN":"集群列表"}
  Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"ingress", "zh_CN":"路由"}
  UpdateIngressIngress *UpdateIngressIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true"`
}

func (s UpdateIngressRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressRequest) GoString() string {
  return s.String()
}

func (s *UpdateIngressRequest) SetControllerName(v string) *UpdateIngressRequest {
  s.ControllerName = &v
  return s
}

func (s *UpdateIngressRequest) SetClusters(v []*string) *UpdateIngressRequest {
  s.Clusters = v
  return s
}

func (s *UpdateIngressRequest) SetIngress(v *UpdateIngressIngress) *UpdateIngressRequest {
  s.UpdateIngressIngress = v
  return s
}

type UpdateIngressResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress object", "zh_CN":"路由对象"}
  Data *UpdateIngressCustomIngress `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateIngressResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressResponse) GoString() string {
  return s.String()
}

func (s *UpdateIngressResponse) SetCode(v int64) *UpdateIngressResponse {
  s.Code = &v
  return s
}

func (s *UpdateIngressResponse) SetMsg(v string) *UpdateIngressResponse {
  s.Msg = &v
  return s
}

func (s *UpdateIngressResponse) SetRequestId(v string) *UpdateIngressResponse {
  s.RequestId = &v
  return s
}

func (s *UpdateIngressResponse) SetData(v *UpdateIngressCustomIngress) *UpdateIngressResponse {
  s.Data = v
  return s
}

type UpdateIngressPaths struct {
  // {"en":"ingress name", "zh_CN":"路由名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateIngressPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressPaths) GoString() string {
  return s.String()
}

func (s *UpdateIngressPaths) SetName(v string) *UpdateIngressPaths {
  s.Name = &v
  return s
}

type UpdateIngressParameters struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s UpdateIngressParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressParameters) GoString() string {
  return s.String()
}

func (s *UpdateIngressParameters) SetNamespace(v string) *UpdateIngressParameters {
  s.Namespace = &v
  return s
}

type UpdateIngressRequestHeader struct {
}

func (s UpdateIngressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressRequestHeader) GoString() string {
  return s.String()
}

type UpdateIngressResponseHeader struct {
}

func (s UpdateIngressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressResponseHeader) GoString() string {
  return s.String()
}

type UpdateIngressCustomIngress struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  ControllerName *string `json:"controllerName,omitempty" xml:"controllerName,omitempty" require:"true"`
  // {"en":"clusters", "zh_CN":"集群列表"}
  Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"ingress", "zh_CN":"路由"}
  UpdateIngressIngress *UpdateIngressIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true"`
}

func (s UpdateIngressCustomIngress) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressCustomIngress) GoString() string {
  return s.String()
}

func (s *UpdateIngressCustomIngress) SetControllerName(v string) *UpdateIngressCustomIngress {
  s.ControllerName = &v
  return s
}

func (s *UpdateIngressCustomIngress) SetClusters(v []*string) *UpdateIngressCustomIngress {
  s.Clusters = v
  return s
}

func (s *UpdateIngressCustomIngress) SetIngress(v *UpdateIngressIngress) *UpdateIngressCustomIngress {
  s.UpdateIngressIngress = v
  return s
}

type UpdateIngressIngress struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdateIngressObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"ingress desired", "zh_CN":"路由期望属性"}
  Spec *UpdateIngressIngressSpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s UpdateIngressIngress) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngress) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngress) SetApiVersion(v string) *UpdateIngressIngress {
  s.ApiVersion = &v
  return s
}

func (s *UpdateIngressIngress) SetKind(v string) *UpdateIngressIngress {
  s.Kind = &v
  return s
}

func (s *UpdateIngressIngress) SetMetadata(v *UpdateIngressObjectMeta) *UpdateIngressIngress {
  s.Metadata = v
  return s
}

func (s *UpdateIngressIngress) SetSpec(v *UpdateIngressIngressSpec) *UpdateIngressIngress {
  s.Spec = v
  return s
}

type UpdateIngressIngressSpec struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  IngressClassName *string `json:"ingressClassName,omitempty" xml:"ingressClassName,omitempty"`
  // {"en":"DefaultBackend is the backend that should handle requests that don't match any rule", "zh_CN":"默认后端,当请求不匹配任何规则时调用"}
  DefaultBackend *UpdateIngressIngressBackend `json:"defaultBackend,omitempty" xml:"defaultBackend,omitempty"`
  Tls []*UpdateIngressIngressTLS `json:"tls,omitempty" xml:"tls,omitempty" type:"Repeated"`
  // {"en":"A list of host rules used to configure the UpdateIngressIngress", "zh_CN":"路由规则列表"}
  Rules []*UpdateIngressIngressRule `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s UpdateIngressIngressSpec) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngressSpec) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngressSpec) SetIngressClassName(v string) *UpdateIngressIngressSpec {
  s.IngressClassName = &v
  return s
}

func (s *UpdateIngressIngressSpec) SetDefaultBackend(v *UpdateIngressIngressBackend) *UpdateIngressIngressSpec {
  s.DefaultBackend = v
  return s
}

func (s *UpdateIngressIngressSpec) SetTls(v []*UpdateIngressIngressTLS) *UpdateIngressIngressSpec {
  s.Tls = v
  return s
}

func (s *UpdateIngressIngressSpec) SetRules(v []*UpdateIngressIngressRule) *UpdateIngressIngressSpec {
  s.Rules = v
  return s
}

type UpdateIngressIngressRule struct {
  // {"en":"Host is the fully qualified domain name of a network host", "zh_CN":"域名"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  Http *UpdateIngressHTTPIngressRuleValue `json:"http,omitempty" xml:"http,omitempty"`
}

func (s UpdateIngressIngressRule) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngressRule) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngressRule) SetHost(v string) *UpdateIngressIngressRule {
  s.Host = &v
  return s
}

func (s *UpdateIngressIngressRule) SetHttp(v *UpdateIngressHTTPIngressRuleValue) *UpdateIngressIngressRule {
  s.Http = v
  return s
}

type UpdateIngressHTTPIngressRuleValue struct {
  // {"en":"A collection of paths that map requests to backends", "zh_CN":"请求路径匹配规则"}
  UpdateIngressPaths []*UpdateIngressHTTPIngressPath `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
}

func (s UpdateIngressHTTPIngressRuleValue) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressHTTPIngressRuleValue) GoString() string {
  return s.String()
}

func (s *UpdateIngressHTTPIngressRuleValue) SetUpdateIngressPaths(v []*UpdateIngressHTTPIngressPath) *UpdateIngressHTTPIngressRuleValue {
  s.UpdateIngressPaths = v
  return s
}

type UpdateIngressHTTPIngressPath struct {
  // {"en":"Path is matched against the path of an incoming request", "zh_CN":"请求路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty"`
  // {"en":"PathType determines the interpretation of the Path matching,PathType can be one of the following values: Exact,Prefix,ImplementationSpecific", "zh_CN":"路径匹配类型: Exact,Prefix,ImplementationSpecific"}
  PathType *string `json:"pathType,omitempty" xml:"pathType,omitempty"`
  // {"en":"Backend defines the referenced service endpoint to which the traffic will be forwarded to", "zh_CN":"指定后端服务"}
  Backend *UpdateIngressIngressBackend `json:"backend,omitempty" xml:"backend,omitempty"`
}

func (s UpdateIngressHTTPIngressPath) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressHTTPIngressPath) GoString() string {
  return s.String()
}

func (s *UpdateIngressHTTPIngressPath) SetPath(v string) *UpdateIngressHTTPIngressPath {
  s.Path = &v
  return s
}

func (s *UpdateIngressHTTPIngressPath) SetPathType(v string) *UpdateIngressHTTPIngressPath {
  s.PathType = &v
  return s
}

func (s *UpdateIngressHTTPIngressPath) SetBackend(v *UpdateIngressIngressBackend) *UpdateIngressHTTPIngressPath {
  s.Backend = v
  return s
}

type UpdateIngressIngressTLS struct {
  // {"en":"Hosts are a list of hosts included in the TLS certificate", "zh_CN":"tls证书包含域名"}
  Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
  // {"en":"SecretName is the name of the secret used to terminate TLS traffic on port 443", "zh_CN":"tls秘钥名称"}
  SecretName *string `json:"secretName,omitempty" xml:"secretName,omitempty"`
}

func (s UpdateIngressIngressTLS) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngressTLS) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngressTLS) SetHosts(v []*string) *UpdateIngressIngressTLS {
  s.Hosts = v
  return s
}

func (s *UpdateIngressIngressTLS) SetSecretName(v string) *UpdateIngressIngressTLS {
  s.SecretName = &v
  return s
}

type UpdateIngressIngressBackend struct {
  // {"en":"Service references a Service as a Backend", "zh_CN":"指定后端服务"}
  Service *UpdateIngressIngressServiceBackend `json:"service,omitempty" xml:"service,omitempty"`
  // {"en":"Resource is an ObjectRef to another Kubernetes resource in the namespace of the UpdateIngressIngress object", "zh_CN":"路由指定后端资源"}
  Resource *UpdateIngressTypedLocalObjectReference `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s UpdateIngressIngressBackend) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngressBackend) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngressBackend) SetService(v *UpdateIngressIngressServiceBackend) *UpdateIngressIngressBackend {
  s.Service = v
  return s
}

func (s *UpdateIngressIngressBackend) SetResource(v *UpdateIngressTypedLocalObjectReference) *UpdateIngressIngressBackend {
  s.Resource = v
  return s
}

type UpdateIngressTypedLocalObjectReference struct {
  // {"en":"Name is the name of resource being referenced", "zh_CN":"资源名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Kind is the type of resource being referenced", "zh_CN":"资源类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"APIGroup is the group for the resource being referenced", "zh_CN":"资源分组"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
}

func (s UpdateIngressTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *UpdateIngressTypedLocalObjectReference) SetName(v string) *UpdateIngressTypedLocalObjectReference {
  s.Name = &v
  return s
}

func (s *UpdateIngressTypedLocalObjectReference) SetKind(v string) *UpdateIngressTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *UpdateIngressTypedLocalObjectReference) SetApiGroup(v string) *UpdateIngressTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

type UpdateIngressIngressServiceBackend struct {
  // {"en":"Name is the referenced service", "zh_CN":"服务名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Port of the referenced service. A port name or port number", "zh_CN":"服务端口或端口名称"}
  Port *UpdateIngressServiceBackendPort `json:"port,omitempty" xml:"port,omitempty"`
}

func (s UpdateIngressIngressServiceBackend) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngressServiceBackend) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngressServiceBackend) SetName(v string) *UpdateIngressIngressServiceBackend {
  s.Name = &v
  return s
}

func (s *UpdateIngressIngressServiceBackend) SetPort(v *UpdateIngressServiceBackendPort) *UpdateIngressIngressServiceBackend {
  s.Port = v
  return s
}

type UpdateIngressServiceBackendPort struct {
  // {"en":"Name is the name of the port on the Service", "zh_CN":"服务端口名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Number is the numerical port number on the Service", "zh_CN":"服务数字端口"}
  Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
}

func (s UpdateIngressServiceBackendPort) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressServiceBackendPort) GoString() string {
  return s.String()
}

func (s *UpdateIngressServiceBackendPort) SetName(v string) *UpdateIngressServiceBackendPort {
  s.Name = &v
  return s
}

func (s *UpdateIngressServiceBackendPort) SetNumber(v int32) *UpdateIngressServiceBackendPort {
  s.Number = &v
  return s
}

type UpdateIngressIngressStatus struct {
  // {"en":"LoadBalancer contains the current status of the load-balancer", "zh_CN":"包含当前负载均衡服务的状态"}
  LoadBalancer *UpdateIngressLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
}

func (s UpdateIngressIngressStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressIngressStatus) GoString() string {
  return s.String()
}

func (s *UpdateIngressIngressStatus) SetLoadBalancer(v *UpdateIngressLoadBalancerStatus) *UpdateIngressIngressStatus {
  s.LoadBalancer = v
  return s
}

type UpdateIngressLoadBalancerStatus struct {
  UpdateIngressIngress []*UpdateIngressLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s UpdateIngressLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *UpdateIngressLoadBalancerStatus) SetIngress(v []*UpdateIngressLoadBalancerIngress) *UpdateIngressLoadBalancerStatus {
  s.UpdateIngressIngress = v
  return s
}

type UpdateIngressLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based", "zh_CN":"负载均衡服务ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based", "zh_CN":"负载均衡类型服务dns"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Ports is a list of records of service ports", "zh_CN":"服务端口状态列表"}
  Ports []*UpdateIngressPortStatus `json:"ports,omitempty" xml:"ports,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateIngressLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *UpdateIngressLoadBalancerIngress) SetIp(v string) *UpdateIngressLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *UpdateIngressLoadBalancerIngress) SetHostname(v string) *UpdateIngressLoadBalancerIngress {
  s.Hostname = &v
  return s
}

func (s *UpdateIngressLoadBalancerIngress) SetPorts(v []*UpdateIngressPortStatus) *UpdateIngressLoadBalancerIngress {
  s.Ports = v
  return s
}

type UpdateIngressPortStatus struct {
  // {"en":"Port is the port number of the service port of which status is recorded here", "zh_CN":"服务端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Protocol is the protocol of the service port of which status is recorded here,The supported values are: TCP, UDP, SCTP", "zh_CN":"服务支持类型: TCP,UDP,SCTP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port", "zh_CN":"记录服务端口错误信息"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s UpdateIngressPortStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressPortStatus) GoString() string {
  return s.String()
}

func (s *UpdateIngressPortStatus) SetPort(v int32) *UpdateIngressPortStatus {
  s.Port = &v
  return s
}

func (s *UpdateIngressPortStatus) SetProtocol(v string) *UpdateIngressPortStatus {
  s.Protocol = &v
  return s
}

func (s *UpdateIngressPortStatus) SetError(v string) *UpdateIngressPortStatus {
  s.Error = &v
  return s
}

type UpdateIngressObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
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
  OwnerReferences []*UpdateIngressOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*UpdateIngressManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s UpdateIngressObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressObjectMeta) GoString() string {
  return s.String()
}

func (s *UpdateIngressObjectMeta) SetName(v string) *UpdateIngressObjectMeta {
  s.Name = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetGenerateName(v string) *UpdateIngressObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetNamespace(v string) *UpdateIngressObjectMeta {
  s.Namespace = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetSelfLink(v string) *UpdateIngressObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetUid(v string) *UpdateIngressObjectMeta {
  s.Uid = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetResourceVersion(v string) *UpdateIngressObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetGeneration(v int64) *UpdateIngressObjectMeta {
  s.Generation = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetCreationTimestamp(v string) *UpdateIngressObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetDeletionTimestamp(v string) *UpdateIngressObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetDeletionGracePeriodSeconds(v int64) *UpdateIngressObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetLabels(v map[string]*string) *UpdateIngressObjectMeta {
  s.Labels = v
  return s
}

func (s *UpdateIngressObjectMeta) SetAnnotations(v map[string]*string) *UpdateIngressObjectMeta {
  s.Annotations = v
  return s
}

func (s *UpdateIngressObjectMeta) SetOwnerReferences(v []*UpdateIngressOwnerReference) *UpdateIngressObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *UpdateIngressObjectMeta) SetFinalizers(v []*string) *UpdateIngressObjectMeta {
  s.Finalizers = v
  return s
}

func (s *UpdateIngressObjectMeta) SetClusterName(v string) *UpdateIngressObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *UpdateIngressObjectMeta) SetManagedFields(v []*UpdateIngressManagedFieldsEntry) *UpdateIngressObjectMeta {
  s.ManagedFields = v
  return s
}

type UpdateIngressManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this UpdateIngressManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'UpdateIngressFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“UpdateIngressFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"UpdateIngressFieldsV1 holds the first JSON version format as described in the 'UpdateIngressFieldsV1' type", "zh_CN":"UpdateIngressFieldsV1 包含类型 “UpdateIngressFieldsV1” 中描述的第一个 JSON 版本格式"}
  UpdateIngressFieldsV1 *UpdateIngressFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s UpdateIngressManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *UpdateIngressManagedFieldsEntry) SetManager(v string) *UpdateIngressManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *UpdateIngressManagedFieldsEntry) SetOperation(v string) *UpdateIngressManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *UpdateIngressManagedFieldsEntry) SetApiVersion(v string) *UpdateIngressManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *UpdateIngressManagedFieldsEntry) SetTime(v string) *UpdateIngressManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *UpdateIngressManagedFieldsEntry) SetFieldsType(v string) *UpdateIngressManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *UpdateIngressManagedFieldsEntry) SetFieldsV1(v *UpdateIngressFieldsV1) *UpdateIngressManagedFieldsEntry {
  s.UpdateIngressFieldsV1 = v
  return s
}

func (s *UpdateIngressManagedFieldsEntry) SetSubresource(v string) *UpdateIngressManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type UpdateIngressFieldsV1 struct {
}

func (s UpdateIngressFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressFieldsV1) GoString() string {
  return s.String()
}

type UpdateIngressOwnerReference struct {
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

func (s UpdateIngressOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressOwnerReference) GoString() string {
  return s.String()
}

func (s *UpdateIngressOwnerReference) SetApiVersion(v string) *UpdateIngressOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *UpdateIngressOwnerReference) SetKind(v string) *UpdateIngressOwnerReference {
  s.Kind = &v
  return s
}

func (s *UpdateIngressOwnerReference) SetName(v string) *UpdateIngressOwnerReference {
  s.Name = &v
  return s
}

func (s *UpdateIngressOwnerReference) SetUid(v string) *UpdateIngressOwnerReference {
  s.Uid = &v
  return s
}

func (s *UpdateIngressOwnerReference) SetController(v bool) *UpdateIngressOwnerReference {
  s.Controller = &v
  return s
}

func (s *UpdateIngressOwnerReference) SetBlockOwnerDeletion(v bool) *UpdateIngressOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type ListIngressRequest struct {
}

func (s ListIngressRequest) String() string {
  return tea.Prettify(s)
}

func (s ListIngressRequest) GoString() string {
  return s.String()
}

type ListIngressResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress object", "zh_CN":"路由对象"}
  Data *ListIngressCustomIngressList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListIngressResponse) String() string {
  return tea.Prettify(s)
}

func (s ListIngressResponse) GoString() string {
  return s.String()
}

func (s *ListIngressResponse) SetCode(v int64) *ListIngressResponse {
  s.Code = &v
  return s
}

func (s *ListIngressResponse) SetMsg(v string) *ListIngressResponse {
  s.Msg = &v
  return s
}

func (s *ListIngressResponse) SetRequestId(v string) *ListIngressResponse {
  s.RequestId = &v
  return s
}

func (s *ListIngressResponse) SetData(v *ListIngressCustomIngressList) *ListIngressResponse {
  s.Data = v
  return s
}

type ListIngressPaths struct {
}

func (s ListIngressPaths) String() string {
  return tea.Prettify(s)
}

func (s ListIngressPaths) GoString() string {
  return s.String()
}

type ListIngressParameters struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"keyword", "zh_CN":"关键字"}
  Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
  // {"en":"page index", "zh_CN":"页数"}
  PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"page size", "zh_CN":"每页数量"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListIngressParameters) String() string {
  return tea.Prettify(s)
}

func (s ListIngressParameters) GoString() string {
  return s.String()
}

func (s *ListIngressParameters) SetNamespace(v string) *ListIngressParameters {
  s.Namespace = &v
  return s
}

func (s *ListIngressParameters) SetKeyword(v string) *ListIngressParameters {
  s.Keyword = &v
  return s
}

func (s *ListIngressParameters) SetPageIndex(v string) *ListIngressParameters {
  s.PageIndex = &v
  return s
}

func (s *ListIngressParameters) SetPageSize(v string) *ListIngressParameters {
  s.PageSize = &v
  return s
}

type ListIngressRequestHeader struct {
}

func (s ListIngressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListIngressRequestHeader) GoString() string {
  return s.String()
}

type ListIngressResponseHeader struct {
}

func (s ListIngressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListIngressResponseHeader) GoString() string {
  return s.String()
}

type ListIngressCustomIngressList struct {
  // {"en":"total  count", "zh_CN":"总数"}
  Total *int32 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  Items []*ListIngressCustomIngressItem `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressCustomIngressList) String() string {
  return tea.Prettify(s)
}

func (s ListIngressCustomIngressList) GoString() string {
  return s.String()
}

func (s *ListIngressCustomIngressList) SetTotal(v int32) *ListIngressCustomIngressList {
  s.Total = &v
  return s
}

func (s *ListIngressCustomIngressList) SetItems(v []*ListIngressCustomIngressItem) *ListIngressCustomIngressList {
  s.Items = v
  return s
}

type ListIngressCustomIngressItem struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  ControllerName *string `json:"controllerName,omitempty" xml:"controllerName,omitempty" require:"true"`
  // {"en":"clusters", "zh_CN":"集群列表"}
  Clusters []*ListIngressCustomCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"ingress", "zh_CN":"路由"}
  ListIngressIngress *ListIngressIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListIngressCustomIngressItem) String() string {
  return tea.Prettify(s)
}

func (s ListIngressCustomIngressItem) GoString() string {
  return s.String()
}

func (s *ListIngressCustomIngressItem) SetControllerName(v string) *ListIngressCustomIngressItem {
  s.ControllerName = &v
  return s
}

func (s *ListIngressCustomIngressItem) SetClusters(v []*ListIngressCustomCluster) *ListIngressCustomIngressItem {
  s.Clusters = v
  return s
}

func (s *ListIngressCustomIngressItem) SetIngress(v *ListIngressIngress) *ListIngressCustomIngressItem {
  s.ListIngressIngress = v
  return s
}

func (s *ListIngressCustomIngressItem) SetStatus(v string) *ListIngressCustomIngressItem {
  s.Status = &v
  return s
}

func (s *ListIngressCustomIngressItem) SetUpdateTime(v int64) *ListIngressCustomIngressItem {
  s.UpdateTime = &v
  return s
}

type ListIngressCustomCluster struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"cluster cn name", "zh_CN":"集群中文名"}
  NameCn *string `json:"nameCn,omitempty" xml:"nameCn,omitempty" require:"true"`
}

func (s ListIngressCustomCluster) String() string {
  return tea.Prettify(s)
}

func (s ListIngressCustomCluster) GoString() string {
  return s.String()
}

func (s *ListIngressCustomCluster) SetName(v string) *ListIngressCustomCluster {
  s.Name = &v
  return s
}

func (s *ListIngressCustomCluster) SetNameCn(v string) *ListIngressCustomCluster {
  s.NameCn = &v
  return s
}

type ListIngressIngress struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *ListIngressObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"ingress desired", "zh_CN":"路由期望属性"}
  Spec *ListIngressIngressSpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
  // {"en":"Status is the current state of the ListIngressIngress", "zh_CN":"路由状态"}
  Status *ListIngressIngressStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIngressIngress) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngress) GoString() string {
  return s.String()
}

func (s *ListIngressIngress) SetApiVersion(v string) *ListIngressIngress {
  s.ApiVersion = &v
  return s
}

func (s *ListIngressIngress) SetKind(v string) *ListIngressIngress {
  s.Kind = &v
  return s
}

func (s *ListIngressIngress) SetMetadata(v *ListIngressObjectMeta) *ListIngressIngress {
  s.Metadata = v
  return s
}

func (s *ListIngressIngress) SetSpec(v *ListIngressIngressSpec) *ListIngressIngress {
  s.Spec = v
  return s
}

func (s *ListIngressIngress) SetStatus(v *ListIngressIngressStatus) *ListIngressIngress {
  s.Status = v
  return s
}

type ListIngressIngressSpec struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  IngressClassName *string `json:"ingressClassName,omitempty" xml:"ingressClassName,omitempty" require:"true"`
  // {"en":"DefaultBackend is the backend that should handle requests that don't match any rule", "zh_CN":"默认后端,当请求不匹配任何规则时调用"}
  DefaultBackend *ListIngressIngressBackend `json:"defaultBackend,omitempty" xml:"defaultBackend,omitempty" require:"true"`
  Tls []*ListIngressIngressTLS `json:"tls,omitempty" xml:"tls,omitempty" require:"true" type:"Repeated"`
  // {"en":"A list of host rules used to configure the ListIngressIngress", "zh_CN":"路由规则列表"}
  Rules []*ListIngressIngressRule `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressIngressSpec) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngressSpec) GoString() string {
  return s.String()
}

func (s *ListIngressIngressSpec) SetIngressClassName(v string) *ListIngressIngressSpec {
  s.IngressClassName = &v
  return s
}

func (s *ListIngressIngressSpec) SetDefaultBackend(v *ListIngressIngressBackend) *ListIngressIngressSpec {
  s.DefaultBackend = v
  return s
}

func (s *ListIngressIngressSpec) SetTls(v []*ListIngressIngressTLS) *ListIngressIngressSpec {
  s.Tls = v
  return s
}

func (s *ListIngressIngressSpec) SetRules(v []*ListIngressIngressRule) *ListIngressIngressSpec {
  s.Rules = v
  return s
}

type ListIngressIngressRule struct {
  // {"en":"Host is the fully qualified domain name of a network host", "zh_CN":"域名"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  Http *ListIngressHTTPIngressRuleValue `json:"http,omitempty" xml:"http,omitempty" require:"true"`
}

func (s ListIngressIngressRule) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngressRule) GoString() string {
  return s.String()
}

func (s *ListIngressIngressRule) SetHost(v string) *ListIngressIngressRule {
  s.Host = &v
  return s
}

func (s *ListIngressIngressRule) SetHttp(v *ListIngressHTTPIngressRuleValue) *ListIngressIngressRule {
  s.Http = v
  return s
}

type ListIngressHTTPIngressRuleValue struct {
  // {"en":"A collection of paths that map requests to backends", "zh_CN":"请求路径匹配规则"}
  ListIngressPaths []*ListIngressHTTPIngressPath `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressHTTPIngressRuleValue) String() string {
  return tea.Prettify(s)
}

func (s ListIngressHTTPIngressRuleValue) GoString() string {
  return s.String()
}

func (s *ListIngressHTTPIngressRuleValue) SetListIngressPaths(v []*ListIngressHTTPIngressPath) *ListIngressHTTPIngressRuleValue {
  s.ListIngressPaths = v
  return s
}

type ListIngressHTTPIngressPath struct {
  // {"en":"Path is matched against the path of an incoming request", "zh_CN":"请求路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"PathType determines the interpretation of the Path matching,PathType can be one of the following values: Exact,Prefix,ImplementationSpecific", "zh_CN":"路径匹配类型: Exact,Prefix,ImplementationSpecific"}
  PathType *string `json:"pathType,omitempty" xml:"pathType,omitempty" require:"true"`
  // {"en":"Backend defines the referenced service endpoint to which the traffic will be forwarded to", "zh_CN":"指定后端服务"}
  Backend *ListIngressIngressBackend `json:"backend,omitempty" xml:"backend,omitempty" require:"true"`
}

func (s ListIngressHTTPIngressPath) String() string {
  return tea.Prettify(s)
}

func (s ListIngressHTTPIngressPath) GoString() string {
  return s.String()
}

func (s *ListIngressHTTPIngressPath) SetPath(v string) *ListIngressHTTPIngressPath {
  s.Path = &v
  return s
}

func (s *ListIngressHTTPIngressPath) SetPathType(v string) *ListIngressHTTPIngressPath {
  s.PathType = &v
  return s
}

func (s *ListIngressHTTPIngressPath) SetBackend(v *ListIngressIngressBackend) *ListIngressHTTPIngressPath {
  s.Backend = v
  return s
}

type ListIngressIngressTLS struct {
  // {"en":"Hosts are a list of hosts included in the TLS certificate", "zh_CN":"tls证书包含域名"}
  Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" require:"true" type:"Repeated"`
  // {"en":"SecretName is the name of the secret used to terminate TLS traffic on port 443", "zh_CN":"tls秘钥名称"}
  SecretName *string `json:"secretName,omitempty" xml:"secretName,omitempty" require:"true"`
}

func (s ListIngressIngressTLS) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngressTLS) GoString() string {
  return s.String()
}

func (s *ListIngressIngressTLS) SetHosts(v []*string) *ListIngressIngressTLS {
  s.Hosts = v
  return s
}

func (s *ListIngressIngressTLS) SetSecretName(v string) *ListIngressIngressTLS {
  s.SecretName = &v
  return s
}

type ListIngressIngressBackend struct {
  // {"en":"Service references a Service as a Backend", "zh_CN":"指定后端服务"}
  Service *ListIngressIngressServiceBackend `json:"service,omitempty" xml:"service,omitempty" require:"true"`
  // {"en":"Resource is an ObjectRef to another Kubernetes resource in the namespace of the ListIngressIngress object", "zh_CN":"路由指定后端资源"}
  Resource *ListIngressTypedLocalObjectReference `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
}

func (s ListIngressIngressBackend) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngressBackend) GoString() string {
  return s.String()
}

func (s *ListIngressIngressBackend) SetService(v *ListIngressIngressServiceBackend) *ListIngressIngressBackend {
  s.Service = v
  return s
}

func (s *ListIngressIngressBackend) SetResource(v *ListIngressTypedLocalObjectReference) *ListIngressIngressBackend {
  s.Resource = v
  return s
}

type ListIngressTypedLocalObjectReference struct {
  // {"en":"Name is the name of resource being referenced", "zh_CN":"资源名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Kind is the type of resource being referenced", "zh_CN":"资源类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"APIGroup is the group for the resource being referenced", "zh_CN":"资源分组"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty" require:"true"`
}

func (s ListIngressTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s ListIngressTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *ListIngressTypedLocalObjectReference) SetName(v string) *ListIngressTypedLocalObjectReference {
  s.Name = &v
  return s
}

func (s *ListIngressTypedLocalObjectReference) SetKind(v string) *ListIngressTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *ListIngressTypedLocalObjectReference) SetApiGroup(v string) *ListIngressTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

type ListIngressIngressServiceBackend struct {
  // {"en":"Name is the referenced service", "zh_CN":"服务名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Port of the referenced service. A port name or port number", "zh_CN":"服务端口或端口名称"}
  Port *ListIngressServiceBackendPort `json:"port,omitempty" xml:"port,omitempty" require:"true"`
}

func (s ListIngressIngressServiceBackend) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngressServiceBackend) GoString() string {
  return s.String()
}

func (s *ListIngressIngressServiceBackend) SetName(v string) *ListIngressIngressServiceBackend {
  s.Name = &v
  return s
}

func (s *ListIngressIngressServiceBackend) SetPort(v *ListIngressServiceBackendPort) *ListIngressIngressServiceBackend {
  s.Port = v
  return s
}

type ListIngressServiceBackendPort struct {
  // {"en":"Name is the name of the port on the Service", "zh_CN":"服务端口名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Number is the numerical port number on the Service", "zh_CN":"服务数字端口"}
  Number *int32 `json:"number,omitempty" xml:"number,omitempty" require:"true"`
}

func (s ListIngressServiceBackendPort) String() string {
  return tea.Prettify(s)
}

func (s ListIngressServiceBackendPort) GoString() string {
  return s.String()
}

func (s *ListIngressServiceBackendPort) SetName(v string) *ListIngressServiceBackendPort {
  s.Name = &v
  return s
}

func (s *ListIngressServiceBackendPort) SetNumber(v int32) *ListIngressServiceBackendPort {
  s.Number = &v
  return s
}

type ListIngressIngressStatus struct {
  // {"en":"LoadBalancer contains the current status of the load-balancer", "zh_CN":"包含当前负载均衡服务的状态"}
  LoadBalancer *ListIngressLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty" require:"true"`
}

func (s ListIngressIngressStatus) String() string {
  return tea.Prettify(s)
}

func (s ListIngressIngressStatus) GoString() string {
  return s.String()
}

func (s *ListIngressIngressStatus) SetLoadBalancer(v *ListIngressLoadBalancerStatus) *ListIngressIngressStatus {
  s.LoadBalancer = v
  return s
}

type ListIngressLoadBalancerStatus struct {
  ListIngressIngress []*ListIngressLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s ListIngressLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *ListIngressLoadBalancerStatus) SetIngress(v []*ListIngressLoadBalancerIngress) *ListIngressLoadBalancerStatus {
  s.ListIngressIngress = v
  return s
}

type ListIngressLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based", "zh_CN":"负载均衡服务ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based", "zh_CN":"负载均衡类型服务dns"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Ports is a list of records of service ports", "zh_CN":"服务端口状态列表"}
  Ports []*ListIngressPortStatus `json:"ports,omitempty" xml:"ports,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s ListIngressLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *ListIngressLoadBalancerIngress) SetIp(v string) *ListIngressLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *ListIngressLoadBalancerIngress) SetHostname(v string) *ListIngressLoadBalancerIngress {
  s.Hostname = &v
  return s
}

func (s *ListIngressLoadBalancerIngress) SetPorts(v []*ListIngressPortStatus) *ListIngressLoadBalancerIngress {
  s.Ports = v
  return s
}

type ListIngressPortStatus struct {
  // {"en":"Port is the port number of the service port of which status is recorded here", "zh_CN":"服务端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Protocol is the protocol of the service port of which status is recorded here,The supported values are: TCP, UDP, SCTP", "zh_CN":"服务支持类型: TCP,UDP,SCTP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"Error is to record the problem with the service port", "zh_CN":"记录服务端口错误信息"}
  Error *string `json:"error,omitempty" xml:"error,omitempty" require:"true"`
}

func (s ListIngressPortStatus) String() string {
  return tea.Prettify(s)
}

func (s ListIngressPortStatus) GoString() string {
  return s.String()
}

func (s *ListIngressPortStatus) SetPort(v int32) *ListIngressPortStatus {
  s.Port = &v
  return s
}

func (s *ListIngressPortStatus) SetProtocol(v string) *ListIngressPortStatus {
  s.Protocol = &v
  return s
}

func (s *ListIngressPortStatus) SetError(v string) *ListIngressPortStatus {
  s.Error = &v
  return s
}

type ListIngressObjectMeta struct {
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
  OwnerReferences []*ListIngressOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*ListIngressManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s ListIngressObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s ListIngressObjectMeta) GoString() string {
  return s.String()
}

func (s *ListIngressObjectMeta) SetName(v string) *ListIngressObjectMeta {
  s.Name = &v
  return s
}

func (s *ListIngressObjectMeta) SetGenerateName(v string) *ListIngressObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *ListIngressObjectMeta) SetNamespace(v string) *ListIngressObjectMeta {
  s.Namespace = &v
  return s
}

func (s *ListIngressObjectMeta) SetSelfLink(v string) *ListIngressObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *ListIngressObjectMeta) SetUid(v string) *ListIngressObjectMeta {
  s.Uid = &v
  return s
}

func (s *ListIngressObjectMeta) SetResourceVersion(v string) *ListIngressObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListIngressObjectMeta) SetGeneration(v int64) *ListIngressObjectMeta {
  s.Generation = &v
  return s
}

func (s *ListIngressObjectMeta) SetCreationTimestamp(v string) *ListIngressObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *ListIngressObjectMeta) SetDeletionTimestamp(v string) *ListIngressObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *ListIngressObjectMeta) SetDeletionGracePeriodSeconds(v int64) *ListIngressObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *ListIngressObjectMeta) SetLabels(v map[string]*string) *ListIngressObjectMeta {
  s.Labels = v
  return s
}

func (s *ListIngressObjectMeta) SetAnnotations(v map[string]*string) *ListIngressObjectMeta {
  s.Annotations = v
  return s
}

func (s *ListIngressObjectMeta) SetOwnerReferences(v []*ListIngressOwnerReference) *ListIngressObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *ListIngressObjectMeta) SetFinalizers(v []*string) *ListIngressObjectMeta {
  s.Finalizers = v
  return s
}

func (s *ListIngressObjectMeta) SetClusterName(v string) *ListIngressObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *ListIngressObjectMeta) SetManagedFields(v []*ListIngressManagedFieldsEntry) *ListIngressObjectMeta {
  s.ManagedFields = v
  return s
}

type ListIngressManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this ListIngressManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'ListIngressFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“ListIngressFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"ListIngressFieldsV1 holds the first JSON version format as described in the 'ListIngressFieldsV1' type", "zh_CN":"ListIngressFieldsV1 包含类型 “ListIngressFieldsV1” 中描述的第一个 JSON 版本格式"}
  ListIngressFieldsV1 *ListIngressFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s ListIngressManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s ListIngressManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *ListIngressManagedFieldsEntry) SetManager(v string) *ListIngressManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *ListIngressManagedFieldsEntry) SetOperation(v string) *ListIngressManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *ListIngressManagedFieldsEntry) SetApiVersion(v string) *ListIngressManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *ListIngressManagedFieldsEntry) SetTime(v string) *ListIngressManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *ListIngressManagedFieldsEntry) SetFieldsType(v string) *ListIngressManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *ListIngressManagedFieldsEntry) SetFieldsV1(v *ListIngressFieldsV1) *ListIngressManagedFieldsEntry {
  s.ListIngressFieldsV1 = v
  return s
}

func (s *ListIngressManagedFieldsEntry) SetSubresource(v string) *ListIngressManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type ListIngressFieldsV1 struct {
}

func (s ListIngressFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s ListIngressFieldsV1) GoString() string {
  return s.String()
}

type ListIngressOwnerReference struct {
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

func (s ListIngressOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s ListIngressOwnerReference) GoString() string {
  return s.String()
}

func (s *ListIngressOwnerReference) SetApiVersion(v string) *ListIngressOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *ListIngressOwnerReference) SetKind(v string) *ListIngressOwnerReference {
  s.Kind = &v
  return s
}

func (s *ListIngressOwnerReference) SetName(v string) *ListIngressOwnerReference {
  s.Name = &v
  return s
}

func (s *ListIngressOwnerReference) SetUid(v string) *ListIngressOwnerReference {
  s.Uid = &v
  return s
}

func (s *ListIngressOwnerReference) SetController(v bool) *ListIngressOwnerReference {
  s.Controller = &v
  return s
}

func (s *ListIngressOwnerReference) SetBlockOwnerDeletion(v bool) *ListIngressOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type LECHAllocateEdgeIPRequest struct {
  // {"en":"node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"IP protocol:4-ipv4(default);6-ipv6(temporary unsupported)","zh_CN":"可选\nIP协议：4-ipv4(默认)；6-ipv6"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"IP native attribute, 1: non-native;-1: native;","zh_CN":"IPv4原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"cidr","zh_CN":"CIDR，一次只能传入一个CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"number of applications IP  (the single upper limit is 50)","zh_CN":"申请IP数（单次申请Ip数上限为50个）"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Allocate IP randomly","zh_CN":"是否需要随机分配IP（仅对ipv4生效）\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
}

func (s LECHAllocateEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHAllocateEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *LECHAllocateEdgeIPRequest) SetNodeName(v string) *LECHAllocateEdgeIPRequest {
  s.NodeName = &v
  return s
}

func (s *LECHAllocateEdgeIPRequest) SetProtocol(v string) *LECHAllocateEdgeIPRequest {
  s.Protocol = &v
  return s
}

func (s *LECHAllocateEdgeIPRequest) SetNativeAttribute(v string) *LECHAllocateEdgeIPRequest {
  s.NativeAttribute = &v
  return s
}

func (s *LECHAllocateEdgeIPRequest) SetCidr(v string) *LECHAllocateEdgeIPRequest {
  s.Cidr = &v
  return s
}

func (s *LECHAllocateEdgeIPRequest) SetCount(v int) *LECHAllocateEdgeIPRequest {
  s.Count = &v
  return s
}

func (s *LECHAllocateEdgeIPRequest) SetRandomAllocateIp(v int) *LECHAllocateEdgeIPRequest {
  s.RandomAllocateIp = &v
  return s
}

type LECHAllocateEdgeIPRequestHeader struct {
}

func (s LECHAllocateEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAllocateEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type LECHAllocateEdgeIPPaths struct {
}

func (s LECHAllocateEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHAllocateEdgeIPPaths) GoString() string {
  return s.String()
}

type LECHAllocateEdgeIPParameters struct {
}

func (s LECHAllocateEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHAllocateEdgeIPParameters) GoString() string {
  return s.String()
}

type LECHAllocateEdgeIPResponse struct {
  // {"en":"successful application for all or part of IP","zh_CN":"成功申请到的全部或部分IP\n说明：不同场景的响应说明如下\nA、所有IP都申请成功，返回申请到的所有IP\nB、只申请到部分IP，返回申请到的那部分IP\nC、未申请到任何IP，返回失败信息\nD、若出现申请失败的情况，请间隔10S之后再次申请"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHAllocateEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHAllocateEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *LECHAllocateEdgeIPResponse) SetEdgeIps(v []*string) *LECHAllocateEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type LECHAllocateEdgeIPResponseHeader struct {
}

func (s LECHAllocateEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAllocateEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type DeleteIngressRequest struct {
}

func (s DeleteIngressRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressRequest) GoString() string {
  return s.String()
}

type DeleteIngressResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s DeleteIngressResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressResponse) GoString() string {
  return s.String()
}

func (s *DeleteIngressResponse) SetCode(v int64) *DeleteIngressResponse {
  s.Code = &v
  return s
}

func (s *DeleteIngressResponse) SetMsg(v string) *DeleteIngressResponse {
  s.Msg = &v
  return s
}

func (s *DeleteIngressResponse) SetRequestId(v string) *DeleteIngressResponse {
  s.RequestId = &v
  return s
}

type DeleteIngressPaths struct {
  // {"en":"ingress name", "zh_CN":"路由名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteIngressPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressPaths) GoString() string {
  return s.String()
}

func (s *DeleteIngressPaths) SetName(v string) *DeleteIngressPaths {
  s.Name = &v
  return s
}

type DeleteIngressParameters struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s DeleteIngressParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressParameters) GoString() string {
  return s.String()
}

func (s *DeleteIngressParameters) SetNamespace(v string) *DeleteIngressParameters {
  s.Namespace = &v
  return s
}

type DeleteIngressRequestHeader struct {
}

func (s DeleteIngressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressRequestHeader) GoString() string {
  return s.String()
}

type DeleteIngressResponseHeader struct {
}

func (s DeleteIngressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryAvailableCidrsRequest struct {
}

func (s LECHQueryAvailableCidrsRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsRequest) GoString() string {
  return s.String()
}

type LECHQueryAvailableCidrsRequestHeader struct {
}

func (s LECHQueryAvailableCidrsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryAvailableCidrsPaths struct {
}

func (s LECHQueryAvailableCidrsPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsPaths) GoString() string {
  return s.String()
}

type LECHQueryAvailableCidrsParameters struct {
  // {"en":"node name","zh_CN":"节点英文名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
}

func (s LECHQueryAvailableCidrsParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsParameters) SetNodeName(v string) *LECHQueryAvailableCidrsParameters {
  s.NodeName = &v
  return s
}

type LECHQueryAvailableCidrsResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHQueryAvailableCidrsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHQueryAvailableCidrsResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsResponse) SetCode(v string) *LECHQueryAvailableCidrsResponse {
  s.Code = &v
  return s
}

func (s *LECHQueryAvailableCidrsResponse) SetData(v *LECHQueryAvailableCidrsResponseData) *LECHQueryAvailableCidrsResponse {
  s.Data = v
  return s
}

func (s *LECHQueryAvailableCidrsResponse) SetMessage(v string) *LECHQueryAvailableCidrsResponse {
  s.Message = &v
  return s
}

type LECHQueryAvailableCidrsResponseData struct {
  // {"en":"available cidrs","zh_CN":"可用的cidr列表"}
  Cidrs []*string `json:"cidrs,omitempty" xml:"cidrs,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryAvailableCidrsResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsResponseData) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsResponseData) SetCidrs(v []*string) *LECHQueryAvailableCidrsResponseData {
  s.Cidrs = v
  return s
}

type LECHQueryAvailableCidrsResponseHeader struct {
}

func (s LECHQueryAvailableCidrsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsResponseHeader) GoString() string {
  return s.String()
}




type CreateNetworkPolicyRequest struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *CreateNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en": "Specification of the desired behavior for this CreateNetworkPolicyNetworkPolicy.", "zh_CN": "CreateNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *CreateNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateNetworkPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyRequest) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyRequest) SetApiVersion(v string) *CreateNetworkPolicyRequest {
  s.ApiVersion = &v
  return s
}

func (s *CreateNetworkPolicyRequest) SetKind(v string) *CreateNetworkPolicyRequest {
  s.Kind = &v
  return s
}

func (s *CreateNetworkPolicyRequest) SetMetadata(v *CreateNetworkPolicyObjectMeta) *CreateNetworkPolicyRequest {
  s.Metadata = v
  return s
}

func (s *CreateNetworkPolicyRequest) SetSpec(v *CreateNetworkPolicyNetworkPolicySpec) *CreateNetworkPolicyRequest {
  s.Spec = v
  return s
}

type CreateNetworkPolicyResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"networkPolicy", "zh_CN":"网络策略"}
  Data *CreateNetworkPolicyNetworkPolicy `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateNetworkPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyResponse) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyResponse) SetCode(v int64) *CreateNetworkPolicyResponse {
  s.Code = &v
  return s
}

func (s *CreateNetworkPolicyResponse) SetMsg(v string) *CreateNetworkPolicyResponse {
  s.Msg = &v
  return s
}

func (s *CreateNetworkPolicyResponse) SetRequestId(v string) *CreateNetworkPolicyResponse {
  s.RequestId = &v
  return s
}

func (s *CreateNetworkPolicyResponse) SetData(v *CreateNetworkPolicyNetworkPolicy) *CreateNetworkPolicyResponse {
  s.Data = v
  return s
}

type CreateNetworkPolicyPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s CreateNetworkPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyPaths) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyPaths) SetNamespace(v string) *CreateNetworkPolicyPaths {
  s.Namespace = &v
  return s
}

type CreateNetworkPolicyParameters struct {
}

func (s CreateNetworkPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyParameters) GoString() string {
  return s.String()
}

type CreateNetworkPolicyRequestHeader struct {
}

func (s CreateNetworkPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyRequestHeader) GoString() string {
  return s.String()
}

type CreateNetworkPolicyResponseHeader struct {
}

func (s CreateNetworkPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyResponseHeader) GoString() string {
  return s.String()
}

type CreateNetworkPolicyNetworkPolicy struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *CreateNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this CreateNetworkPolicyNetworkPolicy.", "zh_CN": "CreateNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *CreateNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s CreateNetworkPolicyNetworkPolicy) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNetworkPolicy) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNetworkPolicy) SetApiVersion(v string) *CreateNetworkPolicyNetworkPolicy {
  s.ApiVersion = &v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicy) SetKind(v string) *CreateNetworkPolicyNetworkPolicy {
  s.Kind = &v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicy) SetMetadata(v *CreateNetworkPolicyObjectMeta) *CreateNetworkPolicyNetworkPolicy {
  s.Metadata = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicy) SetSpec(v *CreateNetworkPolicyNetworkPolicySpec) *CreateNetworkPolicyNetworkPolicy {
  s.Spec = v
  return s
}

type CreateNetworkPolicyNetworkPolicySpec struct {
  // {"en": "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the CreateNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this CreateNetworkPolicyNetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8", "zh_CN": "出网规则"}
  Egress []*CreateNetworkPolicyNetworkPolicyEgressRule `json:"egress,omitempty" xml:"egress,omitempty" type:"Repeated"`
  // {"en": "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the CreateNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this CreateNetworkPolicyNetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)", "zh_CN": "入网规则"}
  Ingress []*CreateNetworkPolicyNetworkPolicyIngressRule `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
  // {"en": "Selects the pods to which this CreateNetworkPolicyNetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.", "zh_CN": "限制pod的选择器"}
  PodSelector *CreateNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty" require:"true"`
  // {"en": "List of rule types that the CreateNetworkPolicyNetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ Egress ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include Egress (since such a policy would not include an Egress section and would otherwise default to just [ Ingress ]). This field is beta-level in 1.8", "zh_CN": "策略类型"}
  PolicyTypes []*string `json:"policyTypes,omitempty" xml:"policyTypes,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyNetworkPolicySpec) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNetworkPolicySpec) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNetworkPolicySpec) SetEgress(v []*CreateNetworkPolicyNetworkPolicyEgressRule) *CreateNetworkPolicyNetworkPolicySpec {
  s.Egress = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicySpec) SetIngress(v []*CreateNetworkPolicyNetworkPolicyIngressRule) *CreateNetworkPolicyNetworkPolicySpec {
  s.Ingress = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicySpec) SetPodSelector(v *CreateNetworkPolicyPodLabelSelector) *CreateNetworkPolicyNetworkPolicySpec {
  s.PodSelector = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicySpec) SetPolicyTypes(v []*string) *CreateNetworkPolicyNetworkPolicySpec {
  s.PolicyTypes = v
  return s
}

type CreateNetworkPolicyNetworkPolicyIngressRule struct {
  // {"en": "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.", "zh_CN": "入网规则信息"}
  From []*CreateNetworkPolicyNetworkPolicyPeer `json:"from,omitempty" xml:"from,omitempty" type:"Repeated"`
  // {"en": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*CreateNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyNetworkPolicyIngressRule) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNetworkPolicyIngressRule) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNetworkPolicyIngressRule) SetFrom(v []*CreateNetworkPolicyNetworkPolicyPeer) *CreateNetworkPolicyNetworkPolicyIngressRule {
  s.From = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicyIngressRule) SetPorts(v []*CreateNetworkPolicyNetworkPolicyPort) *CreateNetworkPolicyNetworkPolicyIngressRule {
  s.Ports = v
  return s
}

type CreateNetworkPolicyNetworkPolicyEgressRule struct {
  // {"en": "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*CreateNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en": "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.", "zh_CN": "出网规则信息"}
  To []*CreateNetworkPolicyNetworkPolicyPeer `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyNetworkPolicyEgressRule) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNetworkPolicyEgressRule) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNetworkPolicyEgressRule) SetPorts(v []*CreateNetworkPolicyNetworkPolicyPort) *CreateNetworkPolicyNetworkPolicyEgressRule {
  s.Ports = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicyEgressRule) SetTo(v []*CreateNetworkPolicyNetworkPolicyPeer) *CreateNetworkPolicyNetworkPolicyEgressRule {
  s.To = v
  return s
}

type CreateNetworkPolicyNetworkPolicyPeer struct {
  // {"en": "CreateNetworkPolicyIPBlock defines policy on a particular CreateNetworkPolicyIPBlock. If this field is set then neither of the other fields can be.", "zh_CN": "IP规则"}
  IpBlock *CreateNetworkPolicyIPBlock `json:"ipBlock,omitempty" xml:"ipBlock,omitempty"`
  // {"en": "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.If PodSelector is also set, then the CreateNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.", "zh_CN": "namespace选择器"}
  NamespaceSelector *CreateNetworkPolicyNsLabelSelector `json:"namespaceSelector,omitempty" xml:"namespaceSelector,omitempty"`
  // {"en": "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.If NamespaceSelector is also set, then the CreateNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.", "zh_CN": "pod选择器"}
  PodSelector *CreateNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty"`
}

func (s CreateNetworkPolicyNetworkPolicyPeer) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNetworkPolicyPeer) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNetworkPolicyPeer) SetIpBlock(v *CreateNetworkPolicyIPBlock) *CreateNetworkPolicyNetworkPolicyPeer {
  s.IpBlock = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicyPeer) SetNamespaceSelector(v *CreateNetworkPolicyNsLabelSelector) *CreateNetworkPolicyNetworkPolicyPeer {
  s.NamespaceSelector = v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicyPeer) SetPodSelector(v *CreateNetworkPolicyPodLabelSelector) *CreateNetworkPolicyNetworkPolicyPeer {
  s.PodSelector = v
  return s
}

type CreateNetworkPolicyIPBlock struct {
  // {"en": "CIDR is a string representing the IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64", "zh_CN": "生效IP网段"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en": "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64 Except values will be rejected if they are outside the CIDR range", "zh_CN": "例外IP网段"}
  Except []*string `json:"except,omitempty" xml:"except,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyIPBlock) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyIPBlock) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyIPBlock) SetCidr(v string) *CreateNetworkPolicyIPBlock {
  s.Cidr = &v
  return s
}

func (s *CreateNetworkPolicyIPBlock) SetExcept(v []*string) *CreateNetworkPolicyIPBlock {
  s.Except = v
  return s
}

type CreateNetworkPolicyNetworkPolicyPort struct {
  // {"en": "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.", "zh_CN": "端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
  // {"en": "The protocol (TCP, UDP) which traffic must match. If not specified, this field defaults to TCP.", "zh_CN": "协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s CreateNetworkPolicyNetworkPolicyPort) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNetworkPolicyPort) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNetworkPolicyPort) SetPort(v string) *CreateNetworkPolicyNetworkPolicyPort {
  s.Port = &v
  return s
}

func (s *CreateNetworkPolicyNetworkPolicyPort) SetProtocol(v string) *CreateNetworkPolicyNetworkPolicyPort {
  s.Protocol = &v
  return s
}

type CreateNetworkPolicyPodLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
}

func (s CreateNetworkPolicyPodLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyPodLabelSelector) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyPodLabelSelector) SetMatchLabels(v map[string]*string) *CreateNetworkPolicyPodLabelSelector {
  s.MatchLabels = v
  return s
}

type CreateNetworkPolicyNsLabelSelector struct {
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*CreateNetworkPolicyLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyNsLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyNsLabelSelector) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyNsLabelSelector) SetMatchExpressions(v []*CreateNetworkPolicyLabelSelectorRequirement) *CreateNetworkPolicyNsLabelSelector {
  s.MatchExpressions = v
  return s
}

type CreateNetworkPolicyLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyLabelSelectorRequirement) SetKey(v string) *CreateNetworkPolicyLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *CreateNetworkPolicyLabelSelectorRequirement) SetOperator(v string) *CreateNetworkPolicyLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *CreateNetworkPolicyLabelSelectorRequirement) SetValues(v []*string) *CreateNetworkPolicyLabelSelectorRequirement {
  s.Values = v
  return s
}

type CreateNetworkPolicyObjectMeta struct {
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
  OwnerReferences []*CreateNetworkPolicyOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*CreateNetworkPolicyManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s CreateNetworkPolicyObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyObjectMeta) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyObjectMeta) SetName(v string) *CreateNetworkPolicyObjectMeta {
  s.Name = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetGenerateName(v string) *CreateNetworkPolicyObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetNamespace(v string) *CreateNetworkPolicyObjectMeta {
  s.Namespace = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetSelfLink(v string) *CreateNetworkPolicyObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetUid(v string) *CreateNetworkPolicyObjectMeta {
  s.Uid = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetResourceVersion(v string) *CreateNetworkPolicyObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetGeneration(v int64) *CreateNetworkPolicyObjectMeta {
  s.Generation = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetCreationTimestamp(v string) *CreateNetworkPolicyObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetDeletionTimestamp(v string) *CreateNetworkPolicyObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetDeletionGracePeriodSeconds(v int64) *CreateNetworkPolicyObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetLabels(v map[string]*string) *CreateNetworkPolicyObjectMeta {
  s.Labels = v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetAnnotations(v map[string]*string) *CreateNetworkPolicyObjectMeta {
  s.Annotations = v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetOwnerReferences(v []*CreateNetworkPolicyOwnerReference) *CreateNetworkPolicyObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetFinalizers(v []*string) *CreateNetworkPolicyObjectMeta {
  s.Finalizers = v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetClusterName(v string) *CreateNetworkPolicyObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *CreateNetworkPolicyObjectMeta) SetManagedFields(v []*CreateNetworkPolicyManagedFieldsEntry) *CreateNetworkPolicyObjectMeta {
  s.ManagedFields = v
  return s
}

type CreateNetworkPolicyOwnerReference struct {
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

func (s CreateNetworkPolicyOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyOwnerReference) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyOwnerReference) SetApiVersion(v string) *CreateNetworkPolicyOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *CreateNetworkPolicyOwnerReference) SetKind(v string) *CreateNetworkPolicyOwnerReference {
  s.Kind = &v
  return s
}

func (s *CreateNetworkPolicyOwnerReference) SetName(v string) *CreateNetworkPolicyOwnerReference {
  s.Name = &v
  return s
}

func (s *CreateNetworkPolicyOwnerReference) SetUid(v string) *CreateNetworkPolicyOwnerReference {
  s.Uid = &v
  return s
}

func (s *CreateNetworkPolicyOwnerReference) SetController(v bool) *CreateNetworkPolicyOwnerReference {
  s.Controller = &v
  return s
}

func (s *CreateNetworkPolicyOwnerReference) SetBlockOwnerDeletion(v bool) *CreateNetworkPolicyOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type CreateNetworkPolicyManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this CreateNetworkPolicyManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'CreateNetworkPolicyFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“CreateNetworkPolicyFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"CreateNetworkPolicyFieldsV1 holds the first JSON version format as described in the 'CreateNetworkPolicyFieldsV1' type", "zh_CN":"CreateNetworkPolicyFieldsV1 包含类型 “CreateNetworkPolicyFieldsV1” 中描述的第一个 JSON 版本格式"}
  CreateNetworkPolicyFieldsV1 *CreateNetworkPolicyFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s CreateNetworkPolicyManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetManager(v string) *CreateNetworkPolicyManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetOperation(v string) *CreateNetworkPolicyManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetApiVersion(v string) *CreateNetworkPolicyManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetTime(v string) *CreateNetworkPolicyManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetFieldsType(v string) *CreateNetworkPolicyManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetFieldsV1(v *CreateNetworkPolicyFieldsV1) *CreateNetworkPolicyManagedFieldsEntry {
  s.CreateNetworkPolicyFieldsV1 = v
  return s
}

func (s *CreateNetworkPolicyManagedFieldsEntry) SetSubresource(v string) *CreateNetworkPolicyManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type CreateNetworkPolicyFieldsV1 struct {
}

func (s CreateNetworkPolicyFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s CreateNetworkPolicyFieldsV1) GoString() string {
  return s.String()
}




type UpdateServiceRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdateServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 UpdateServiceService 的行为"}
  Spec *UpdateServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the service. Populated by the system. Read-only", "zh_CN":"最近观察到的 UpdateServiceService 状态。由系统填充。只读"}
  Status *UpdateServiceServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
  return s.String()
}

func (s *UpdateServiceRequest) SetApiVersion(v string) *UpdateServiceRequest {
  s.ApiVersion = &v
  return s
}

func (s *UpdateServiceRequest) SetKind(v string) *UpdateServiceRequest {
  s.Kind = &v
  return s
}

func (s *UpdateServiceRequest) SetMetadata(v *UpdateServiceObjectMeta) *UpdateServiceRequest {
  s.Metadata = v
  return s
}

func (s *UpdateServiceRequest) SetSpec(v *UpdateServiceServiceSpec) *UpdateServiceRequest {
  s.Spec = v
  return s
}

func (s *UpdateServiceRequest) SetStatus(v *UpdateServiceServiceStatus) *UpdateServiceRequest {
  s.Status = v
  return s
}

type UpdateServiceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"service", "zh_CN":"service"}
  Data *UpdateServiceService `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
  return s.String()
}

func (s *UpdateServiceResponse) SetCode(v int64) *UpdateServiceResponse {
  s.Code = &v
  return s
}

func (s *UpdateServiceResponse) SetMsg(v string) *UpdateServiceResponse {
  s.Msg = &v
  return s
}

func (s *UpdateServiceResponse) SetRequestId(v string) *UpdateServiceResponse {
  s.RequestId = &v
  return s
}

func (s *UpdateServiceResponse) SetData(v *UpdateServiceService) *UpdateServiceResponse {
  s.Data = v
  return s
}

type UpdateServicePaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"service name", "zh_CN":"service 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateServicePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateServicePaths) GoString() string {
  return s.String()
}

func (s *UpdateServicePaths) SetNamespace(v string) *UpdateServicePaths {
  s.Namespace = &v
  return s
}

func (s *UpdateServicePaths) SetName(v string) *UpdateServicePaths {
  s.Name = &v
  return s
}

type UpdateServiceParameters struct {
}

func (s UpdateServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceParameters) GoString() string {
  return s.String()
}

type UpdateServiceRequestHeader struct {
}

func (s UpdateServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceRequestHeader) GoString() string {
  return s.String()
}

type UpdateServiceResponseHeader struct {
}

func (s UpdateServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceResponseHeader) GoString() string {
  return s.String()
}

type UpdateServiceService struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdateServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 UpdateServiceService 的行为"}
  Spec *UpdateServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the service. Populated by the system. Read-only", "zh_CN":"最近观察到的 UpdateServiceService 状态。由系统填充。只读"}
  Status *UpdateServiceServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateServiceService) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceService) GoString() string {
  return s.String()
}

func (s *UpdateServiceService) SetApiVersion(v string) *UpdateServiceService {
  s.ApiVersion = &v
  return s
}

func (s *UpdateServiceService) SetKind(v string) *UpdateServiceService {
  s.Kind = &v
  return s
}

func (s *UpdateServiceService) SetMetadata(v *UpdateServiceObjectMeta) *UpdateServiceService {
  s.Metadata = v
  return s
}

func (s *UpdateServiceService) SetSpec(v *UpdateServiceServiceSpec) *UpdateServiceService {
  s.Spec = v
  return s
}

func (s *UpdateServiceService) SetStatus(v *UpdateServiceServiceStatus) *UpdateServiceService {
  s.Status = v
  return s
}

type UpdateServiceServiceStatus struct {
  // {"en":"Current service state", "zh_CN":"loadBalancer 包含负载均衡器的当前状态（如果存在）"}
  LoadBalancer *UpdateServiceLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
  // {"en":"LoadBalancer contains the current status of the load-balancer, if one is present", "zh_CN":"服务的当前状态"}
  Conditions []*UpdateServiceMetaV1Condition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s UpdateServiceServiceStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceServiceStatus) GoString() string {
  return s.String()
}

func (s *UpdateServiceServiceStatus) SetLoadBalancer(v *UpdateServiceLoadBalancerStatus) *UpdateServiceServiceStatus {
  s.LoadBalancer = v
  return s
}

func (s *UpdateServiceServiceStatus) SetConditions(v []*UpdateServiceMetaV1Condition) *UpdateServiceServiceStatus {
  s.Conditions = v
  return s
}

type UpdateServiceMetaV1Condition struct {
  // {"en":"type of condition in CamelCase or in foo.example.com/CamelCase", "zh_CN":"CamelCase 或 foo.example.com/CamelCase 中的条件类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status of the condition, one of True, False, Unknown", "zh_CN":"condition 的状态，True、False、Unknown 之一"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance", "zh_CN":"表示设置 condition 基于的 .metadata.generation 的过期次数。 例如，如果 .metadata.generation 当前为 12，但 .status.conditions[x].observedGeneration 为 9， 则 condition 相对于实例的当前状态已过期"}
  ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
  // {"en":"lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable", "zh_CN":"状况最近一次状态转化的时间。 变化应该发生在下层状况发生变化的时候。如果不知道下层状况发生变化的时间， 那么使用 API 字段更改的时间是可以接受的"}
  LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
  // {"en":"reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty", "zh_CN":"reason 包含一个程序标识符，指示 condition 最后一次转换的原因。 特定条件类型的生产者可以定义该字段的预期值和含义，以及这些值是否被视为有保证的 API。 该值应该是 CamelCase 字符串且不能为空"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":"message is a human readable message indicating details about the transition. This may be an empty string", "zh_CN":"message 是人类可读的消息，有关转换的详细信息，可以是空字符串"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateServiceMetaV1Condition) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceMetaV1Condition) GoString() string {
  return s.String()
}

func (s *UpdateServiceMetaV1Condition) SetType(v string) *UpdateServiceMetaV1Condition {
  s.Type = &v
  return s
}

func (s *UpdateServiceMetaV1Condition) SetStatus(v string) *UpdateServiceMetaV1Condition {
  s.Status = &v
  return s
}

func (s *UpdateServiceMetaV1Condition) SetObservedGeneration(v int64) *UpdateServiceMetaV1Condition {
  s.ObservedGeneration = &v
  return s
}

func (s *UpdateServiceMetaV1Condition) SetLastTransitionTime(v string) *UpdateServiceMetaV1Condition {
  s.LastTransitionTime = &v
  return s
}

func (s *UpdateServiceMetaV1Condition) SetReason(v string) *UpdateServiceMetaV1Condition {
  s.Reason = &v
  return s
}

func (s *UpdateServiceMetaV1Condition) SetMessage(v string) *UpdateServiceMetaV1Condition {
  s.Message = &v
  return s
}

type UpdateServiceLoadBalancerStatus struct {
  // {"en":"Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points", "zh_CN":"ingress 是一个包含负载均衡器 Ingress 点的列表。UpdateServiceService 的流量需要被发送到这些 Ingress 点"}
  Ingress []*UpdateServiceLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s UpdateServiceLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *UpdateServiceLoadBalancerStatus) SetIngress(v []*UpdateServiceLoadBalancerIngress) *UpdateServiceLoadBalancerStatus {
  s.Ingress = v
  return s
}

type UpdateServiceLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)", "zh_CN":"ip 是为基于 IP 的负载均衡器 Ingress 点（通常是 GCE 或 OpenStack 负载均衡器）设置的"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)", "zh_CN":"hostname 是为基于 DNS 的负载均衡器 Ingress 点（通常是 AWS 负载均衡器）设置的"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
  // {"en":"Ports is a list of records of service ports If used, every port defined in the service should have an entry in it", "zh_CN":"ports 是 UpdateServiceService 的端口列表。如果设置了此字段，UpdateServiceService 中定义的每个端口都应该在此列表中"}
  Ports []*UpdateServicePortStatus `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s UpdateServiceLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *UpdateServiceLoadBalancerIngress) SetIp(v string) *UpdateServiceLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *UpdateServiceLoadBalancerIngress) SetHostName(v string) *UpdateServiceLoadBalancerIngress {
  s.HostName = &v
  return s
}

func (s *UpdateServiceLoadBalancerIngress) SetPorts(v []*UpdateServicePortStatus) *UpdateServiceLoadBalancerIngress {
  s.Ports = v
  return s
}

type UpdateServicePortStatus struct {
  // {"en":"the port number of the service port of which status is recorded here", "zh_CN":"port 是所记录的服务端口状态的端口号"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"the protocol of the service port of which status is recorded here The supported values are: 'TCP', 'UDP', 'SCTP'", "zh_CN":"protocol 是所记录的服务端口状态的协议。支持的值为：“TCP”、”UDP”、“SCTP”"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase", "zh_CN":"error 是记录 UpdateServiceService 端口的问题。 错误的格式应符合以下规则:内置错误原因应在此文件中指定，应使用 CamelCase 名称。云提供商特定错误原因的名称必须符合格式 foo.example.com/CamelCase"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s UpdateServicePortStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdateServicePortStatus) GoString() string {
  return s.String()
}

func (s *UpdateServicePortStatus) SetPort(v int32) *UpdateServicePortStatus {
  s.Port = &v
  return s
}

func (s *UpdateServicePortStatus) SetProtocol(v string) *UpdateServicePortStatus {
  s.Protocol = &v
  return s
}

func (s *UpdateServicePortStatus) SetError(v string) *UpdateServicePortStatus {
  s.Error = &v
  return s
}

type UpdateServiceServiceSpec struct {
  // {"en":"The list of ports that are exposed by this service", "zh_CN":"此 UpdateServiceService 公开的端口列表"}
  Ports []*UpdateServiceServicePort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en":"Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName", "zh_CN":"将 UpdateServiceService 流量路由到具有与此 selector 匹配的标签键值对的 Pod。 如果为空或不存在，则假定该服务有一个外部进程管理其端点，Kubernetes 不会修改该端点。 仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型。如果类型为 ExternalName，则忽略"}
  Selector map[string]*string `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a UpdateServiceService of type ExternalName, creation will fail. This field will be wiped when updating a UpdateServiceService to type ExternalName", "zh_CN":"clusterIP 是服务的 IP 地址，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给服务，否则创建服务将失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIP 为空）或 type 已经是 ExternalName 时，可以更改 clusterIP（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIP 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 仅适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 UpdateServiceService 时指定了 clusterIP，则创建将失败。 更新 UpdateServiceService type 为 ExternalName 时，clusterIP 会被移除"}
  ClusterIP *string `json:"clusterIP,omitempty" xml:"clusterIP,omitempty"`
  // {"en":"ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a UpdateServiceService of type ExternalName, creation will fail. This field will be wiped when updating a UpdateServiceService to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"clusterIPs 是分配给该 UpdateServiceService 的 IP 地址列表，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给 UpdateServiceService；否则创建 UpdateServiceService 失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIPs 为空）或 type 已经是 ExternalName 时，可以更改 clusterIPs（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIPs 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 UpdateServiceService 时指定了 clusterIPs，则会创建失败。 更新 UpdateServiceService type 为 ExternalName 时，该字段将被移除。如果未指定此字段，则将从 clusterIP 字段初始化。 如果指定 clusterIPs，客户端必须确保 clusterIPs[0] 和 clusterIP 一致。clusterIPs 最多可包含两个条目（双栈系列，按任意顺序）。 这些 IP 必须与 ipFamilies 的值相对应。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 管理"}
  ClusterIPs []*string `json:"clusterIPs,omitempty" xml:"clusterIPs,omitempty" type:"Repeated"`
  // {"en":"type determines how the UpdateServiceService is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services", "zh_CN":"type 确定 UpdateServiceService 的公开方式。默认为 ClusterIP。 有效选项为 ExternalName、ClusterIP、NodePort 和 LoadBalancer。 “ClusterIP” 为端点分配一个集群内部 IP 地址用于负载均衡。 Endpoints 由 selector 确定，如果未设置 selector，则需要通过手动构造 Endpoints 或 EndpointSlice 的对象来确定。 如果 clusterIP 为 “None”，则不分配虚拟 IP，并且 Endpoints 作为一组端点而不是虚拟 IP 发布。 “NodePort” 建立在 ClusterIP 之上，并在每个节点上分配一个端口，该端口路由到与 clusterIP 相同的 Endpoints。 “LoadBalancer” 基于 NodePort 构建并创建一个外部负载均衡器（如果当前云支持），该负载均衡器路由到与 clusterIP 相同的 Endpoints。 “externalName” 将此 UpdateServiceService 别名为指定的 externalName。其他几个字段不适用于 ExternalName UpdateServiceService"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system", "zh_CN":"externalIPs 是一个 IP 列表，集群中的节点会为此 UpdateServiceService 接收针对这些 IP 地址的流量。 这些 IP 不被 Kubernetes 管理。用户需要确保流量可以到达具有此 IP 的节点。 一个常见的例子是不属于 Kubernetes 系统的外部负载均衡器"}
  ExternalIPs []*string `json:"externalIPs,omitempty" xml:"externalIPs,omitempty" type:"Repeated"`
  // {"en":"Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None", "zh_CN":"支持 “ClientIP” 和 “None”。用于维护会话亲和性。 启用基于客户端 IP 的会话亲和性。必须是 ClientIP 或 None。默认为 None"}
  SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
  // {"en":"Only applies to UpdateServiceService Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version", "zh_CN":"仅适用于服务类型: LoadBalancer。此功能取决于底层云提供商是否支持负载均衡器。 如果云提供商不支持该功能，该字段将被忽略。 已弃用: 该字段信息不足，且其含义因实现而异，而且不支持双栈。 从 Kubernetes v1.24 开始，鼓励用户在可用时使用特定于实现的注释。在未来的 API 版本中可能会删除此字段"}
  LoadBalancerIP *string `json:"loadBalancerIP,omitempty" xml:"loadBalancerIP,omitempty"`
  // {"en":"If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature", "zh_CN":"如果设置了此字段并且被平台支持，将限制通过云厂商的负载均衡器的流量到指定的客户端 IP。 如果云提供商不支持该功能，该字段将被忽略"}
  LoadBalancerSourceRanges []*string `json:"loadBalancerSourceRanges,omitempty" xml:"loadBalancerSourceRanges,omitempty" type:"Repeated"`
  // {"en":"externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires type to be 'ExternalName'", "zh_CN":"externalName 是发现机制将返回的外部引用，作为此服务的别名（例如 DNS CNAME 记录）。 不涉及代理。必须是小写的 RFC-1123 主机名 (https://tools.ietf.org/html/rfc1123)， 并且要求 type 为 “ExternalName”"}
  ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
  // {"en":"externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the UpdateServiceService's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node", "zh_CN":"externalTrafficPolicy 描述了节点如何分发它们在 UpdateServiceService 的“外部访问”地址 （NodePort、ExternalIP 和 LoadBalancer IP）接收到的服务流量。 如果设置为 “Local”，代理将以一种假设外部负载均衡器将负责在节点之间服务流量负载均衡， 因此每个节点将仅向服务的节点本地端点传递流量，而不会伪装客户端源 IP。 （将丢弃错误发送到没有端点的节点的流量。） “Cluster” 默认值使用负载均衡路由到所有端点的策略（可能会根据拓扑和其他特性进行修改）。 请注意，从集群内部发送到 External IP 或 LoadBalancer IP 的流量始终具有 “Cluster” 语义， 但是从集群内部发送到 NodePort 的客户端需要在选择节点时考虑流量路由策略"}
  ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" xml:"externalTrafficPolicy,omitempty"`
  // {"en":"healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a UpdateServiceService which does not need it, creation will fail. This field will be wiped when updating a UpdateServiceService to no longer need it (e.g. changing type). This field cannot be updated once set", "zh_CN":"healthCheckNodePort 指定 UpdateServiceService 的健康检查节点端口。 仅适用于 type 为 LoadBalancer 且 externalTrafficPolicy 设置为 Local 的情况。 如果为此字段设定了一个值，该值在合法范围内且没有被使用，则使用所指定的值。 如果未设置此字段，则自动分配字段值。外部系统（例如负载平衡器）可以使用此端口来确定给定节点是否拥有此服务的端点。 在创建不需要 healthCheckNodePort 的 UpdateServiceService 时指定了此字段，则 UpdateServiceService 创建会失败。 要移除 healthCheckNodePort，需要更改 UpdateServiceService 的 type。 该字段一旦设置就无法更改"}
  HealthCheckNodePort *int32 `json:"healthCheckNodePort,omitempty" xml:"healthCheckNodePort,omitempty"`
  // {"en":"publishNotReadyAddresses indicates that any agent which deals with endpoints for this UpdateServiceService should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless UpdateServiceService to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior", "zh_CN":"publishNotReadyAddresses 表示任何处理此 UpdateServiceService 端点的代理都应忽略任何准备就绪/未准备就绪的指示。 设置此字段的主要场景是为 StatefulSet 的服务提供支持，使之能够为其 Pod 传播 SRV DNS 记录，以实现对等发现。 为 UpdateServiceService 生成 Endpoints 和 EndpointSlice 资源的 Kubernetes 控制器对字段的解读是， 即使 Pod 本身还没有准备好，所有端点都可被视为 “已就绪”。 对于代理而言，如果仅使用 Kubernetes 通过 Endpoints 或 EndpointSlice 资源所生成的端点， 则可以安全地假设这种行为"}
  PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" xml:"publishNotReadyAddresses,omitempty"`
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"sessionAffinityConfig 包含会话亲和性的配置"}
  UpdateServiceSessionAffinityConfig *UpdateServiceSessionAffinityConfig `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
  // {"en":"IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the UpdateServiceService. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a UpdateServiceService to type ExternalName.This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"iPFamilies 是分配给此服务的 IP 协议（例如 IPv4、IPv6）的列表。 该字段通常根据集群配置和 ipFamilyPolicy 字段自动设置。 如果手动指定该字段，且请求的协议在集群中可用，且 ipFamilyPolicy 允许，则使用；否则服务创建将失败。 该字段修改是有条件的：它允许添加或删除辅助 IP 协议，但不允许更改服务的主要 IP 协议。 有效值为 “IPv4” 和 “IPv6”。 该字段仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型的服务，并且确实可用于“无头”服务。 更新服务设置类型为 ExternalName 时，该字段将被擦除。该字段最多可以包含两个条目（双栈系列，按任意顺序）。 如果指定，这些协议栈必须对应于 clusterIPs 字段的值。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 字段管理"}
  IpFamilies []*string `json:"ipFamilies,omitempty" xml:"ipFamilies,omitempty" type:"Repeated"`
  // {"en":"IPFamilyPolicy represents the dual-stack-ness requested or required by this UpdateServiceService. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName", "zh_CN":"iPFamilyPolicy 表示此服务请求或要求的双栈特性。 如果没有提供值，则此字段将被设置为 SingleStack。 服务可以是 “SingleStack”（单个 IP 协议）、 “PreferDualStack”（双栈配置集群上的两个 IP 协议或单栈集群上的单个 IP 协议） 或 “RequireDualStack”（双栈上的两个 IP 协议配置的集群，否则失败）。 ipFamilies 和 clusterIPs 字段取决于此字段的值。 更新服务设置类型为 ExternalName 时，此字段将被擦除"}
  IpFamilyPolicy *string `json:"ipFamilyPolicy,omitempty" xml:"ipFamilyPolicy,omitempty"`
  // {"en":"allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type", "zh_CN":"allocateLoadBalancerNodePorts 定义了是否会自动为 LoadBalancer 类型的 UpdateServiceService 分配 NodePort。默认为 true。 如果集群负载均衡器不依赖 NodePort，则可以设置此字段为 false。 如果调用者（通过指定一个值）请求特定的 NodePort，则无论此字段如何，都会接受这些请求。 该字段只能设置在 type 为 LoadBalancer 的 UpdateServiceService 上，如果 type 更改为任何其他类型，该字段将被移除"}
  AllocateLoadBalancerNodePorts *bool `json:"allocateLoadBalancerNodePorts,omitempty" xml:"allocateLoadBalancerNodePorts,omitempty"`
  // {"en":"loadBalancerClass is the class of the load balancer implementation this UpdateServiceService belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the UpdateServiceService type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a UpdateServiceService to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type", "zh_CN":"loadBalancerClass 是此 UpdateServiceService 所属的负载均衡器实现的类。 如果设置了此字段，则字段值必须是标签风格的标识符，带有可选前缀，例如 ”internal-vip” 或 “example.com/internal-vip”。 无前缀名称是为最终用户保留的。该字段只能在 UpdateServiceService 类型为 “LoadBalancer” 时设置。 如果未设置此字段，则使用默认负载均衡器实现。默认负载均衡器现在通常通过云提供商集成完成，但应适用于任何默认实现。 如果设置了此字段，则假定负载均衡器实现正在监测具有对应负载均衡器类的 UpdateServiceService。 任何默认负载均衡器实现（例如云提供商）都应忽略设置此字段的 UpdateServiceService。 只有在创建或更新的 UpdateServiceService 的 type 为 “LoadBalancer” 时，才可设置此字段。 一经设定，不可更改。当 UpdateServiceService 的 type 更新为 “LoadBalancer” 之外的其他类型时，此字段将被移除"}
  LoadBalancerClass *string `json:"loadBalancerClass,omitempty" xml:"loadBalancerClass,omitempty"`
  // {"en":"InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features)", "zh_CN":"InternalTrafficPolicy 描述节点如何分发它们在 ClusterIP 上接收到的服务流量。 如果设置为 “Local”，代理将假定 Pod 只想与在同一节点上的服务端点通信，如果没有本地端点，它将丢弃流量。 “Cluster” 默认将流量路由到所有端点（可能会根据拓扑和其他特性进行修改）"}
  InternalTrafficPolicy *string `json:"internalTrafficPolicy,omitempty" xml:"internalTrafficPolicy,omitempty"`
}

func (s UpdateServiceServiceSpec) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceServiceSpec) GoString() string {
  return s.String()
}

func (s *UpdateServiceServiceSpec) SetPorts(v []*UpdateServiceServicePort) *UpdateServiceServiceSpec {
  s.Ports = v
  return s
}

func (s *UpdateServiceServiceSpec) SetSelector(v map[string]*string) *UpdateServiceServiceSpec {
  s.Selector = v
  return s
}

func (s *UpdateServiceServiceSpec) SetClusterIP(v string) *UpdateServiceServiceSpec {
  s.ClusterIP = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetClusterIPs(v []*string) *UpdateServiceServiceSpec {
  s.ClusterIPs = v
  return s
}

func (s *UpdateServiceServiceSpec) SetType(v string) *UpdateServiceServiceSpec {
  s.Type = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetExternalIPs(v []*string) *UpdateServiceServiceSpec {
  s.ExternalIPs = v
  return s
}

func (s *UpdateServiceServiceSpec) SetSessionAffinity(v string) *UpdateServiceServiceSpec {
  s.SessionAffinity = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetLoadBalancerIP(v string) *UpdateServiceServiceSpec {
  s.LoadBalancerIP = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetLoadBalancerSourceRanges(v []*string) *UpdateServiceServiceSpec {
  s.LoadBalancerSourceRanges = v
  return s
}

func (s *UpdateServiceServiceSpec) SetExternalName(v string) *UpdateServiceServiceSpec {
  s.ExternalName = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetExternalTrafficPolicy(v string) *UpdateServiceServiceSpec {
  s.ExternalTrafficPolicy = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetHealthCheckNodePort(v int32) *UpdateServiceServiceSpec {
  s.HealthCheckNodePort = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetPublishNotReadyAddresses(v bool) *UpdateServiceServiceSpec {
  s.PublishNotReadyAddresses = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetSessionAffinityConfig(v *UpdateServiceSessionAffinityConfig) *UpdateServiceServiceSpec {
  s.UpdateServiceSessionAffinityConfig = v
  return s
}

func (s *UpdateServiceServiceSpec) SetIpFamilies(v []*string) *UpdateServiceServiceSpec {
  s.IpFamilies = v
  return s
}

func (s *UpdateServiceServiceSpec) SetIpFamilyPolicy(v string) *UpdateServiceServiceSpec {
  s.IpFamilyPolicy = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetAllocateLoadBalancerNodePorts(v bool) *UpdateServiceServiceSpec {
  s.AllocateLoadBalancerNodePorts = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetLoadBalancerClass(v string) *UpdateServiceServiceSpec {
  s.LoadBalancerClass = &v
  return s
}

func (s *UpdateServiceServiceSpec) SetInternalTrafficPolicy(v string) *UpdateServiceServiceSpec {
  s.InternalTrafficPolicy = &v
  return s
}

type UpdateServiceServicePort struct {
  // {"en":"The name of this port within the service. This must be a DNS_LABEL. All ports within a UpdateServiceServiceSpec must have unique names. When considering the endpoints for a UpdateServiceService, this must match the 'name' field in the EndpointPort. Optional if only one UpdateServiceServicePort is defined on this service", "zh_CN":"UpdateServiceService 中此端口的名称。这必须是 DNS_LABEL。 UpdateServiceServiceSpec 中的所有端口的名称都必须唯一。 在考虑 UpdateServiceService 的端点时，这一字段值必须与 EndpointPort 中的 name 字段相同。 如果此服务上仅定义一个 UpdateServiceServicePort，则为此字段为可选"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The IP protocol for this port. Supports TCP, UDP, and SCTP. Default is TCP", "zh_CN":"此端口的 IP 协议。支持 “TCP”、“UDP” 和 “SCTP”。默认为 TCP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol", "zh_CN":"此端口的应用协议，遵循标准的 Kubernetes 标签语法，无前缀名称按照 IANA 标准服务名称 （参见 RFC-6335 和 https://www.iana.org/assignments/service-names）。 非标准协议应该使用前缀名称，如 mycompany.com/my-custom-protocol"}
  AppProtocol *string `json:"appProtocol,omitempty" xml:"appProtocol,omitempty"`
  // {"en":"The port that will be exposed by this service", "zh_CN":"UpdateServiceService 将公开的端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field", "zh_CN":"在 UpdateServiceService 所针对的 Pod 上要访问的端口号或名称。 编号必须在 1 到 65535 的范围内。名称必须是 IANA_SVC_NAME。 如果此值是一个字符串，将在目标 Pod 的容器端口中作为命名端口进行查找。 如果未指定字段，则使用 “port” 字段的值（直接映射）。 对于 clusterIP 为 None 的服务，此字段将被忽略， 应忽略不设或设置为 “port” 字段的取值"}
  TargetPort *int32 `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
  // {"en":"The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this UpdateServiceService requires one. If this field is specified when creating a UpdateServiceService which does not need it, creation will fail. This field will be wiped when updating a UpdateServiceService to no longer need it (e.g. changing type from NodePort to ClusterIP).", "zh_CN":"当类型为 NodePort 或 LoadBalancer 时，UpdateServiceService 公开在节点上的端口， 通常由系统分配。如果指定了一个在范围内且未使用的值，则将使用该值，否则操作将失败。 如果在创建的 UpdateServiceService 需要该端口时未指定该字段，则会分配端口。 如果在创建不需要该端口的 Service时指定了该字段，则会创建失败。 当更新 UpdateServiceService 时，如果不再需要此字段（例如，将类型从 NodePort 更改为 ClusterIP），这个字段将被擦除"}
  NodePort *int32 `json:"nodePort,omitempty" xml:"nodePort,omitempty"`
}

func (s UpdateServiceServicePort) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceServicePort) GoString() string {
  return s.String()
}

func (s *UpdateServiceServicePort) SetName(v string) *UpdateServiceServicePort {
  s.Name = &v
  return s
}

func (s *UpdateServiceServicePort) SetProtocol(v string) *UpdateServiceServicePort {
  s.Protocol = &v
  return s
}

func (s *UpdateServiceServicePort) SetAppProtocol(v string) *UpdateServiceServicePort {
  s.AppProtocol = &v
  return s
}

func (s *UpdateServiceServicePort) SetPort(v int32) *UpdateServiceServicePort {
  s.Port = &v
  return s
}

func (s *UpdateServiceServicePort) SetTargetPort(v int32) *UpdateServiceServicePort {
  s.TargetPort = &v
  return s
}

func (s *UpdateServiceServicePort) SetNodePort(v int32) *UpdateServiceServicePort {
  s.NodePort = &v
  return s
}

type UpdateServiceSessionAffinityConfig struct {
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"clientIP 包含基于客户端 IP 的会话亲和性的配置"}
  ClientIP *UpdateServiceClientIPConfig `json:"clientIP,omitempty" xml:"clientIP,omitempty"`
}

func (s UpdateServiceSessionAffinityConfig) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceSessionAffinityConfig) GoString() string {
  return s.String()
}

func (s *UpdateServiceSessionAffinityConfig) SetClientIP(v *UpdateServiceClientIPConfig) *UpdateServiceSessionAffinityConfig {
  s.ClientIP = v
  return s
}

type UpdateServiceClientIPConfig struct {
  // {"en":"timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).", "zh_CN":"timeoutSeconds 指定 ClientIP 类型会话的维系时间秒数。 如果 ServiceAffinity == 'ClientIP'，则该值必须 >0 && <=86400（1 天）。默认值为 10800（3 小时）"}
  TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s UpdateServiceClientIPConfig) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceClientIPConfig) GoString() string {
  return s.String()
}

func (s *UpdateServiceClientIPConfig) SetTimeoutSeconds(v int32) *UpdateServiceClientIPConfig {
  s.TimeoutSeconds = &v
  return s
}

type UpdateServiceObjectMeta struct {
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
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 UpdateServiceService 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*UpdateServiceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*UpdateServiceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s UpdateServiceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceObjectMeta) GoString() string {
  return s.String()
}

func (s *UpdateServiceObjectMeta) SetName(v string) *UpdateServiceObjectMeta {
  s.Name = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetGenerateName(v string) *UpdateServiceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetNamespace(v string) *UpdateServiceObjectMeta {
  s.Namespace = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetSelfLink(v string) *UpdateServiceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetUid(v string) *UpdateServiceObjectMeta {
  s.Uid = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetResourceVersion(v string) *UpdateServiceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetGeneration(v int64) *UpdateServiceObjectMeta {
  s.Generation = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetCreationTimestamp(v string) *UpdateServiceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetDeletionTimestamp(v string) *UpdateServiceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *UpdateServiceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetLabels(v map[string]*string) *UpdateServiceObjectMeta {
  s.Labels = v
  return s
}

func (s *UpdateServiceObjectMeta) SetAnnotations(v map[string]*string) *UpdateServiceObjectMeta {
  s.Annotations = v
  return s
}

func (s *UpdateServiceObjectMeta) SetOwnerReferences(v []*UpdateServiceOwnerReference) *UpdateServiceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *UpdateServiceObjectMeta) SetFinalizers(v []*string) *UpdateServiceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *UpdateServiceObjectMeta) SetClusterName(v string) *UpdateServiceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *UpdateServiceObjectMeta) SetManagedFields(v []*UpdateServiceManagedFieldsEntry) *UpdateServiceObjectMeta {
  s.ManagedFields = v
  return s
}

type UpdateServiceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this UpdateServiceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'UpdateServiceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“UpdateServiceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"UpdateServiceFieldsV1 holds the first JSON version format as described in the 'UpdateServiceFieldsV1' type", "zh_CN":"UpdateServiceFieldsV1 包含类型 “UpdateServiceFieldsV1” 中描述的第一个 JSON 版本格式"}
  UpdateServiceFieldsV1 *UpdateServiceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s UpdateServiceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *UpdateServiceManagedFieldsEntry) SetManager(v string) *UpdateServiceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *UpdateServiceManagedFieldsEntry) SetOperation(v string) *UpdateServiceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *UpdateServiceManagedFieldsEntry) SetApiVersion(v string) *UpdateServiceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *UpdateServiceManagedFieldsEntry) SetTime(v string) *UpdateServiceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *UpdateServiceManagedFieldsEntry) SetFieldsType(v string) *UpdateServiceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *UpdateServiceManagedFieldsEntry) SetFieldsV1(v *UpdateServiceFieldsV1) *UpdateServiceManagedFieldsEntry {
  s.UpdateServiceFieldsV1 = v
  return s
}

func (s *UpdateServiceManagedFieldsEntry) SetSubresource(v string) *UpdateServiceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type UpdateServiceFieldsV1 struct {
}

func (s UpdateServiceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceFieldsV1) GoString() string {
  return s.String()
}

type UpdateServiceOwnerReference struct {
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

func (s UpdateServiceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s UpdateServiceOwnerReference) GoString() string {
  return s.String()
}

func (s *UpdateServiceOwnerReference) SetApiVersion(v string) *UpdateServiceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *UpdateServiceOwnerReference) SetKind(v string) *UpdateServiceOwnerReference {
  s.Kind = &v
  return s
}

func (s *UpdateServiceOwnerReference) SetName(v string) *UpdateServiceOwnerReference {
  s.Name = &v
  return s
}

func (s *UpdateServiceOwnerReference) SetUid(v string) *UpdateServiceOwnerReference {
  s.Uid = &v
  return s
}

func (s *UpdateServiceOwnerReference) SetController(v bool) *UpdateServiceOwnerReference {
  s.Controller = &v
  return s
}

func (s *UpdateServiceOwnerReference) SetBlockOwnerDeletion(v bool) *UpdateServiceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type ListServiceRequest struct {
}

func (s ListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ListServiceRequest) GoString() string {
  return s.String()
}

type ListServiceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"service list", "zh_CN":"service 列表"}
  Data *ListServiceServiceList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ListServiceResponse) GoString() string {
  return s.String()
}

func (s *ListServiceResponse) SetCode(v int64) *ListServiceResponse {
  s.Code = &v
  return s
}

func (s *ListServiceResponse) SetMsg(v string) *ListServiceResponse {
  s.Msg = &v
  return s
}

func (s *ListServiceResponse) SetRequestId(v string) *ListServiceResponse {
  s.RequestId = &v
  return s
}

func (s *ListServiceResponse) SetData(v *ListServiceServiceList) *ListServiceResponse {
  s.Data = v
  return s
}

type ListServicePaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s ListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ListServicePaths) GoString() string {
  return s.String()
}

func (s *ListServicePaths) SetNamespace(v string) *ListServicePaths {
  s.Namespace = &v
  return s
}

type ListServiceParameters struct {
  // {"en":"service name", "zh_CN":"service 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"labelSelector", "zh_CN":"labelSelector"}
  LabelSelector *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ListServiceParameters) GoString() string {
  return s.String()
}

func (s *ListServiceParameters) SetName(v string) *ListServiceParameters {
  s.Name = &v
  return s
}

func (s *ListServiceParameters) SetLabelSelector(v string) *ListServiceParameters {
  s.LabelSelector = &v
  return s
}

type ListServiceRequestHeader struct {
}

func (s ListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListServiceRequestHeader) GoString() string {
  return s.String()
}

type ListServiceResponseHeader struct {
}

func (s ListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListServiceResponseHeader) GoString() string {
  return s.String()
}

type ListServiceServiceList struct {
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Standard list metadata", "zh_CN":"标准列表元数据"}
  Metadata *ListServiceListMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"List of services", "zh_CN":"ListServiceService 列表"}
  Items []*ListServiceService `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListServiceServiceList) String() string {
  return tea.Prettify(s)
}

func (s ListServiceServiceList) GoString() string {
  return s.String()
}

func (s *ListServiceServiceList) SetKind(v string) *ListServiceServiceList {
  s.Kind = &v
  return s
}

func (s *ListServiceServiceList) SetApiVersion(v string) *ListServiceServiceList {
  s.ApiVersion = &v
  return s
}

func (s *ListServiceServiceList) SetMetadata(v *ListServiceListMeta) *ListServiceServiceList {
  s.Metadata = v
  return s
}

func (s *ListServiceServiceList) SetItems(v []*ListServiceService) *ListServiceServiceList {
  s.Items = v
  return s
}

type ListServiceListMeta struct {
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system", "zh_CN":"selfLink 表示此对象的 URL，由系统填充，只读。已弃用：selfLink 是一个遗留的只读字段，不再由系统填充。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only", "zh_CN":"标识该对象的服务器内部版本的字符串，客户端可以用该字段来确定对象何时被更改。 该值对客户端是不透明的，并且应该原样传回给服务器。该值由系统填充，只读"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message", "zh_CN":"如果用户对返回的条目数量设置了限制，则 continue 可能被设置，表示服务器有更多可用的数据。 该值是不透明的，可用于向提供此列表服务的端点发出另一个请求，以检索下一组可用的对象。 如果服务器配置已更改或时间已过去几分钟，则可能无法继续提供一致的列表。 除非你在错误消息中收到此令牌（token），否则使用此 continue 值时返回的 resourceVersion 字段应该和第一个响应中的值是相同的"}
  Continue *string `json:"continue,omitempty" xml:"continue,omitempty"`
  // {"en":"remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is estimating the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact", "zh_CN":"remainingItemCount 是列表中未包含在此列表响应中的后续项目的数量。 如果列表请求包含标签或字段选择器，则剩余项目的数量是未知的，并且在序列化期间该字段将保持未设置和省略。 如果列表是完整的（因为它没有分块或者这是最后一个块），那么就没有剩余的项目，并且在序列化过程中该字段将保持未设置和省略。 早于 v1.15 的服务器不设置此字段。remainingItemCount 的预期用途是估计集合的大小。 客户端不应依赖于设置准确的 remainingItemCount"}
  RemainingItemCount *int64 `json:"remainingItemCount,omitempty" xml:"remainingItemCount,omitempty"`
}

func (s ListServiceListMeta) String() string {
  return tea.Prettify(s)
}

func (s ListServiceListMeta) GoString() string {
  return s.String()
}

func (s *ListServiceListMeta) SetSelfLink(v string) *ListServiceListMeta {
  s.SelfLink = &v
  return s
}

func (s *ListServiceListMeta) SetResourceVersion(v string) *ListServiceListMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListServiceListMeta) SetContinue(v string) *ListServiceListMeta {
  s.Continue = &v
  return s
}

func (s *ListServiceListMeta) SetRemainingItemCount(v int64) *ListServiceListMeta {
  s.RemainingItemCount = &v
  return s
}

type ListServiceService struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *ListServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 ListServiceService 的行为"}
  Spec *ListServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the service. Populated by the system. Read-only", "zh_CN":"最近观察到的 ListServiceService 状态。由系统填充。只读"}
  Status *ListServiceServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceService) String() string {
  return tea.Prettify(s)
}

func (s ListServiceService) GoString() string {
  return s.String()
}

func (s *ListServiceService) SetApiVersion(v string) *ListServiceService {
  s.ApiVersion = &v
  return s
}

func (s *ListServiceService) SetKind(v string) *ListServiceService {
  s.Kind = &v
  return s
}

func (s *ListServiceService) SetMetadata(v *ListServiceObjectMeta) *ListServiceService {
  s.Metadata = v
  return s
}

func (s *ListServiceService) SetSpec(v *ListServiceServiceSpec) *ListServiceService {
  s.Spec = v
  return s
}

func (s *ListServiceService) SetStatus(v *ListServiceServiceStatus) *ListServiceService {
  s.Status = v
  return s
}

type ListServiceServiceStatus struct {
  // {"en":"Current service state", "zh_CN":"loadBalancer 包含负载均衡器的当前状态（如果存在）"}
  LoadBalancer *ListServiceLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
  // {"en":"LoadBalancer contains the current status of the load-balancer, if one is present", "zh_CN":"服务的当前状态"}
  Conditions []*ListServiceMetaV1Condition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s ListServiceServiceStatus) String() string {
  return tea.Prettify(s)
}

func (s ListServiceServiceStatus) GoString() string {
  return s.String()
}

func (s *ListServiceServiceStatus) SetLoadBalancer(v *ListServiceLoadBalancerStatus) *ListServiceServiceStatus {
  s.LoadBalancer = v
  return s
}

func (s *ListServiceServiceStatus) SetConditions(v []*ListServiceMetaV1Condition) *ListServiceServiceStatus {
  s.Conditions = v
  return s
}

type ListServiceMetaV1Condition struct {
  // {"en":"type of condition in CamelCase or in foo.example.com/CamelCase", "zh_CN":"CamelCase 或 foo.example.com/CamelCase 中的条件类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status of the condition, one of True, False, Unknown", "zh_CN":"condition 的状态，True、False、Unknown 之一"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance", "zh_CN":"表示设置 condition 基于的 .metadata.generation 的过期次数。 例如，如果 .metadata.generation 当前为 12，但 .status.conditions[x].observedGeneration 为 9， 则 condition 相对于实例的当前状态已过期"}
  ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
  // {"en":"lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable", "zh_CN":"状况最近一次状态转化的时间。 变化应该发生在下层状况发生变化的时候。如果不知道下层状况发生变化的时间， 那么使用 API 字段更改的时间是可以接受的"}
  LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
  // {"en":"reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty", "zh_CN":"reason 包含一个程序标识符，指示 condition 最后一次转换的原因。 特定条件类型的生产者可以定义该字段的预期值和含义，以及这些值是否被视为有保证的 API。 该值应该是 CamelCase 字符串且不能为空"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":"message is a human readable message indicating details about the transition. This may be an empty string", "zh_CN":"message 是人类可读的消息，有关转换的详细信息，可以是空字符串"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListServiceMetaV1Condition) String() string {
  return tea.Prettify(s)
}

func (s ListServiceMetaV1Condition) GoString() string {
  return s.String()
}

func (s *ListServiceMetaV1Condition) SetType(v string) *ListServiceMetaV1Condition {
  s.Type = &v
  return s
}

func (s *ListServiceMetaV1Condition) SetStatus(v string) *ListServiceMetaV1Condition {
  s.Status = &v
  return s
}

func (s *ListServiceMetaV1Condition) SetObservedGeneration(v int64) *ListServiceMetaV1Condition {
  s.ObservedGeneration = &v
  return s
}

func (s *ListServiceMetaV1Condition) SetLastTransitionTime(v string) *ListServiceMetaV1Condition {
  s.LastTransitionTime = &v
  return s
}

func (s *ListServiceMetaV1Condition) SetReason(v string) *ListServiceMetaV1Condition {
  s.Reason = &v
  return s
}

func (s *ListServiceMetaV1Condition) SetMessage(v string) *ListServiceMetaV1Condition {
  s.Message = &v
  return s
}

type ListServiceLoadBalancerStatus struct {
  // {"en":"Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points", "zh_CN":"ingress 是一个包含负载均衡器 Ingress 点的列表。ListServiceService 的流量需要被发送到这些 Ingress 点"}
  Ingress []*ListServiceLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s ListServiceLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s ListServiceLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *ListServiceLoadBalancerStatus) SetIngress(v []*ListServiceLoadBalancerIngress) *ListServiceLoadBalancerStatus {
  s.Ingress = v
  return s
}

type ListServiceLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)", "zh_CN":"ip 是为基于 IP 的负载均衡器 Ingress 点（通常是 GCE 或 OpenStack 负载均衡器）设置的"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)", "zh_CN":"hostname 是为基于 DNS 的负载均衡器 Ingress 点（通常是 AWS 负载均衡器）设置的"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
  // {"en":"Ports is a list of records of service ports If used, every port defined in the service should have an entry in it", "zh_CN":"ports 是 ListServiceService 的端口列表。如果设置了此字段，ListServiceService 中定义的每个端口都应该在此列表中"}
  Ports []*ListServicePortStatus `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s ListServiceLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s ListServiceLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *ListServiceLoadBalancerIngress) SetIp(v string) *ListServiceLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *ListServiceLoadBalancerIngress) SetHostName(v string) *ListServiceLoadBalancerIngress {
  s.HostName = &v
  return s
}

func (s *ListServiceLoadBalancerIngress) SetPorts(v []*ListServicePortStatus) *ListServiceLoadBalancerIngress {
  s.Ports = v
  return s
}

type ListServicePortStatus struct {
  // {"en":"the port number of the service port of which status is recorded here", "zh_CN":"port 是所记录的服务端口状态的端口号"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"the protocol of the service port of which status is recorded here The supported values are: 'TCP', 'UDP', 'SCTP'", "zh_CN":"protocol 是所记录的服务端口状态的协议。支持的值为：“TCP”、”UDP”、“SCTP”"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase", "zh_CN":"error 是记录 ListServiceService 端口的问题。 错误的格式应符合以下规则:内置错误原因应在此文件中指定，应使用 CamelCase 名称。云提供商特定错误原因的名称必须符合格式 foo.example.com/CamelCase"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s ListServicePortStatus) String() string {
  return tea.Prettify(s)
}

func (s ListServicePortStatus) GoString() string {
  return s.String()
}

func (s *ListServicePortStatus) SetPort(v int32) *ListServicePortStatus {
  s.Port = &v
  return s
}

func (s *ListServicePortStatus) SetProtocol(v string) *ListServicePortStatus {
  s.Protocol = &v
  return s
}

func (s *ListServicePortStatus) SetError(v string) *ListServicePortStatus {
  s.Error = &v
  return s
}

type ListServiceServiceSpec struct {
  // {"en":"The list of ports that are exposed by this service", "zh_CN":"此 ListServiceService 公开的端口列表"}
  Ports []*ListServiceServicePort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en":"Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName", "zh_CN":"将 ListServiceService 流量路由到具有与此 selector 匹配的标签键值对的 Pod。 如果为空或不存在，则假定该服务有一个外部进程管理其端点，Kubernetes 不会修改该端点。 仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型。如果类型为 ExternalName，则忽略"}
  Selector map[string]*string `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a ListServiceService of type ExternalName, creation will fail. This field will be wiped when updating a ListServiceService to type ExternalName", "zh_CN":"clusterIP 是服务的 IP 地址，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给服务，否则创建服务将失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIP 为空）或 type 已经是 ExternalName 时，可以更改 clusterIP（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIP 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 仅适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 ListServiceService 时指定了 clusterIP，则创建将失败。 更新 ListServiceService type 为 ExternalName 时，clusterIP 会被移除"}
  ClusterIP *string `json:"clusterIP,omitempty" xml:"clusterIP,omitempty"`
  // {"en":"ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a ListServiceService of type ExternalName, creation will fail. This field will be wiped when updating a ListServiceService to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"clusterIPs 是分配给该 ListServiceService 的 IP 地址列表，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给 ListServiceService；否则创建 ListServiceService 失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIPs 为空）或 type 已经是 ExternalName 时，可以更改 clusterIPs（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIPs 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 ListServiceService 时指定了 clusterIPs，则会创建失败。 更新 ListServiceService type 为 ExternalName 时，该字段将被移除。如果未指定此字段，则将从 clusterIP 字段初始化。 如果指定 clusterIPs，客户端必须确保 clusterIPs[0] 和 clusterIP 一致。clusterIPs 最多可包含两个条目（双栈系列，按任意顺序）。 这些 IP 必须与 ipFamilies 的值相对应。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 管理"}
  ClusterIPs []*string `json:"clusterIPs,omitempty" xml:"clusterIPs,omitempty" type:"Repeated"`
  // {"en":"type determines how the ListServiceService is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services", "zh_CN":"type 确定 ListServiceService 的公开方式。默认为 ClusterIP。 有效选项为 ExternalName、ClusterIP、NodePort 和 LoadBalancer。 “ClusterIP” 为端点分配一个集群内部 IP 地址用于负载均衡。 Endpoints 由 selector 确定，如果未设置 selector，则需要通过手动构造 Endpoints 或 EndpointSlice 的对象来确定。 如果 clusterIP 为 “None”，则不分配虚拟 IP，并且 Endpoints 作为一组端点而不是虚拟 IP 发布。 “NodePort” 建立在 ClusterIP 之上，并在每个节点上分配一个端口，该端口路由到与 clusterIP 相同的 Endpoints。 “LoadBalancer” 基于 NodePort 构建并创建一个外部负载均衡器（如果当前云支持），该负载均衡器路由到与 clusterIP 相同的 Endpoints。 “externalName” 将此 ListServiceService 别名为指定的 externalName。其他几个字段不适用于 ExternalName ListServiceService"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system", "zh_CN":"externalIPs 是一个 IP 列表，集群中的节点会为此 ListServiceService 接收针对这些 IP 地址的流量。 这些 IP 不被 Kubernetes 管理。用户需要确保流量可以到达具有此 IP 的节点。 一个常见的例子是不属于 Kubernetes 系统的外部负载均衡器"}
  ExternalIPs []*string `json:"externalIPs,omitempty" xml:"externalIPs,omitempty" type:"Repeated"`
  // {"en":"Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None", "zh_CN":"支持 “ClientIP” 和 “None”。用于维护会话亲和性。 启用基于客户端 IP 的会话亲和性。必须是 ClientIP 或 None。默认为 None"}
  SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
  // {"en":"Only applies to ListServiceService Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version", "zh_CN":"仅适用于服务类型: LoadBalancer。此功能取决于底层云提供商是否支持负载均衡器。 如果云提供商不支持该功能，该字段将被忽略。 已弃用: 该字段信息不足，且其含义因实现而异，而且不支持双栈。 从 Kubernetes v1.24 开始，鼓励用户在可用时使用特定于实现的注释。在未来的 API 版本中可能会删除此字段"}
  LoadBalancerIP *string `json:"loadBalancerIP,omitempty" xml:"loadBalancerIP,omitempty"`
  // {"en":"If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature", "zh_CN":"如果设置了此字段并且被平台支持，将限制通过云厂商的负载均衡器的流量到指定的客户端 IP。 如果云提供商不支持该功能，该字段将被忽略"}
  LoadBalancerSourceRanges []*string `json:"loadBalancerSourceRanges,omitempty" xml:"loadBalancerSourceRanges,omitempty" type:"Repeated"`
  // {"en":"externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires type to be 'ExternalName'", "zh_CN":"externalName 是发现机制将返回的外部引用，作为此服务的别名（例如 DNS CNAME 记录）。 不涉及代理。必须是小写的 RFC-1123 主机名 (https://tools.ietf.org/html/rfc1123)， 并且要求 type 为 “ExternalName”"}
  ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
  // {"en":"externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the ListServiceService's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node", "zh_CN":"externalTrafficPolicy 描述了节点如何分发它们在 ListServiceService 的“外部访问”地址 （NodePort、ExternalIP 和 LoadBalancer IP）接收到的服务流量。 如果设置为 “Local”，代理将以一种假设外部负载均衡器将负责在节点之间服务流量负载均衡， 因此每个节点将仅向服务的节点本地端点传递流量，而不会伪装客户端源 IP。 （将丢弃错误发送到没有端点的节点的流量。） “Cluster” 默认值使用负载均衡路由到所有端点的策略（可能会根据拓扑和其他特性进行修改）。 请注意，从集群内部发送到 External IP 或 LoadBalancer IP 的流量始终具有 “Cluster” 语义， 但是从集群内部发送到 NodePort 的客户端需要在选择节点时考虑流量路由策略"}
  ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" xml:"externalTrafficPolicy,omitempty"`
  // {"en":"healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a ListServiceService which does not need it, creation will fail. This field will be wiped when updating a ListServiceService to no longer need it (e.g. changing type). This field cannot be updated once set", "zh_CN":"healthCheckNodePort 指定 ListServiceService 的健康检查节点端口。 仅适用于 type 为 LoadBalancer 且 externalTrafficPolicy 设置为 Local 的情况。 如果为此字段设定了一个值，该值在合法范围内且没有被使用，则使用所指定的值。 如果未设置此字段，则自动分配字段值。外部系统（例如负载平衡器）可以使用此端口来确定给定节点是否拥有此服务的端点。 在创建不需要 healthCheckNodePort 的 ListServiceService 时指定了此字段，则 ListServiceService 创建会失败。 要移除 healthCheckNodePort，需要更改 ListServiceService 的 type。 该字段一旦设置就无法更改"}
  HealthCheckNodePort *int32 `json:"healthCheckNodePort,omitempty" xml:"healthCheckNodePort,omitempty"`
  // {"en":"publishNotReadyAddresses indicates that any agent which deals with endpoints for this ListServiceService should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless ListServiceService to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior", "zh_CN":"publishNotReadyAddresses 表示任何处理此 ListServiceService 端点的代理都应忽略任何准备就绪/未准备就绪的指示。 设置此字段的主要场景是为 StatefulSet 的服务提供支持，使之能够为其 Pod 传播 SRV DNS 记录，以实现对等发现。 为 ListServiceService 生成 Endpoints 和 EndpointSlice 资源的 Kubernetes 控制器对字段的解读是， 即使 Pod 本身还没有准备好，所有端点都可被视为 “已就绪”。 对于代理而言，如果仅使用 Kubernetes 通过 Endpoints 或 EndpointSlice 资源所生成的端点， 则可以安全地假设这种行为"}
  PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" xml:"publishNotReadyAddresses,omitempty"`
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"sessionAffinityConfig 包含会话亲和性的配置"}
  ListServiceSessionAffinityConfig *ListServiceSessionAffinityConfig `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
  // {"en":"IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the ListServiceService. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a ListServiceService to type ExternalName.This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"iPFamilies 是分配给此服务的 IP 协议（例如 IPv4、IPv6）的列表。 该字段通常根据集群配置和 ipFamilyPolicy 字段自动设置。 如果手动指定该字段，且请求的协议在集群中可用，且 ipFamilyPolicy 允许，则使用；否则服务创建将失败。 该字段修改是有条件的：它允许添加或删除辅助 IP 协议，但不允许更改服务的主要 IP 协议。 有效值为 “IPv4” 和 “IPv6”。 该字段仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型的服务，并且确实可用于“无头”服务。 更新服务设置类型为 ExternalName 时，该字段将被擦除。该字段最多可以包含两个条目（双栈系列，按任意顺序）。 如果指定，这些协议栈必须对应于 clusterIPs 字段的值。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 字段管理"}
  IpFamilies []*string `json:"ipFamilies,omitempty" xml:"ipFamilies,omitempty" type:"Repeated"`
  // {"en":"IPFamilyPolicy represents the dual-stack-ness requested or required by this ListServiceService. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName", "zh_CN":"iPFamilyPolicy 表示此服务请求或要求的双栈特性。 如果没有提供值，则此字段将被设置为 SingleStack。 服务可以是 “SingleStack”（单个 IP 协议）、 “PreferDualStack”（双栈配置集群上的两个 IP 协议或单栈集群上的单个 IP 协议） 或 “RequireDualStack”（双栈上的两个 IP 协议配置的集群，否则失败）。 ipFamilies 和 clusterIPs 字段取决于此字段的值。 更新服务设置类型为 ExternalName 时，此字段将被擦除"}
  IpFamilyPolicy *string `json:"ipFamilyPolicy,omitempty" xml:"ipFamilyPolicy,omitempty"`
  // {"en":"allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type", "zh_CN":"allocateLoadBalancerNodePorts 定义了是否会自动为 LoadBalancer 类型的 ListServiceService 分配 NodePort。默认为 true。 如果集群负载均衡器不依赖 NodePort，则可以设置此字段为 false。 如果调用者（通过指定一个值）请求特定的 NodePort，则无论此字段如何，都会接受这些请求。 该字段只能设置在 type 为 LoadBalancer 的 ListServiceService 上，如果 type 更改为任何其他类型，该字段将被移除"}
  AllocateLoadBalancerNodePorts *bool `json:"allocateLoadBalancerNodePorts,omitempty" xml:"allocateLoadBalancerNodePorts,omitempty"`
  // {"en":"loadBalancerClass is the class of the load balancer implementation this ListServiceService belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the ListServiceService type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a ListServiceService to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type", "zh_CN":"loadBalancerClass 是此 ListServiceService 所属的负载均衡器实现的类。 如果设置了此字段，则字段值必须是标签风格的标识符，带有可选前缀，例如 ”internal-vip” 或 “example.com/internal-vip”。 无前缀名称是为最终用户保留的。该字段只能在 ListServiceService 类型为 “LoadBalancer” 时设置。 如果未设置此字段，则使用默认负载均衡器实现。默认负载均衡器现在通常通过云提供商集成完成，但应适用于任何默认实现。 如果设置了此字段，则假定负载均衡器实现正在监测具有对应负载均衡器类的 ListServiceService。 任何默认负载均衡器实现（例如云提供商）都应忽略设置此字段的 ListServiceService。 只有在创建或更新的 ListServiceService 的 type 为 “LoadBalancer” 时，才可设置此字段。 一经设定，不可更改。当 ListServiceService 的 type 更新为 “LoadBalancer” 之外的其他类型时，此字段将被移除"}
  LoadBalancerClass *string `json:"loadBalancerClass,omitempty" xml:"loadBalancerClass,omitempty"`
  // {"en":"InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features)", "zh_CN":"InternalTrafficPolicy 描述节点如何分发它们在 ClusterIP 上接收到的服务流量。 如果设置为 “Local”，代理将假定 Pod 只想与在同一节点上的服务端点通信，如果没有本地端点，它将丢弃流量。 “Cluster” 默认将流量路由到所有端点（可能会根据拓扑和其他特性进行修改）"}
  InternalTrafficPolicy *string `json:"internalTrafficPolicy,omitempty" xml:"internalTrafficPolicy,omitempty"`
}

func (s ListServiceServiceSpec) String() string {
  return tea.Prettify(s)
}

func (s ListServiceServiceSpec) GoString() string {
  return s.String()
}

func (s *ListServiceServiceSpec) SetPorts(v []*ListServiceServicePort) *ListServiceServiceSpec {
  s.Ports = v
  return s
}

func (s *ListServiceServiceSpec) SetSelector(v map[string]*string) *ListServiceServiceSpec {
  s.Selector = v
  return s
}

func (s *ListServiceServiceSpec) SetClusterIP(v string) *ListServiceServiceSpec {
  s.ClusterIP = &v
  return s
}

func (s *ListServiceServiceSpec) SetClusterIPs(v []*string) *ListServiceServiceSpec {
  s.ClusterIPs = v
  return s
}

func (s *ListServiceServiceSpec) SetType(v string) *ListServiceServiceSpec {
  s.Type = &v
  return s
}

func (s *ListServiceServiceSpec) SetExternalIPs(v []*string) *ListServiceServiceSpec {
  s.ExternalIPs = v
  return s
}

func (s *ListServiceServiceSpec) SetSessionAffinity(v string) *ListServiceServiceSpec {
  s.SessionAffinity = &v
  return s
}

func (s *ListServiceServiceSpec) SetLoadBalancerIP(v string) *ListServiceServiceSpec {
  s.LoadBalancerIP = &v
  return s
}

func (s *ListServiceServiceSpec) SetLoadBalancerSourceRanges(v []*string) *ListServiceServiceSpec {
  s.LoadBalancerSourceRanges = v
  return s
}

func (s *ListServiceServiceSpec) SetExternalName(v string) *ListServiceServiceSpec {
  s.ExternalName = &v
  return s
}

func (s *ListServiceServiceSpec) SetExternalTrafficPolicy(v string) *ListServiceServiceSpec {
  s.ExternalTrafficPolicy = &v
  return s
}

func (s *ListServiceServiceSpec) SetHealthCheckNodePort(v int32) *ListServiceServiceSpec {
  s.HealthCheckNodePort = &v
  return s
}

func (s *ListServiceServiceSpec) SetPublishNotReadyAddresses(v bool) *ListServiceServiceSpec {
  s.PublishNotReadyAddresses = &v
  return s
}

func (s *ListServiceServiceSpec) SetSessionAffinityConfig(v *ListServiceSessionAffinityConfig) *ListServiceServiceSpec {
  s.ListServiceSessionAffinityConfig = v
  return s
}

func (s *ListServiceServiceSpec) SetIpFamilies(v []*string) *ListServiceServiceSpec {
  s.IpFamilies = v
  return s
}

func (s *ListServiceServiceSpec) SetIpFamilyPolicy(v string) *ListServiceServiceSpec {
  s.IpFamilyPolicy = &v
  return s
}

func (s *ListServiceServiceSpec) SetAllocateLoadBalancerNodePorts(v bool) *ListServiceServiceSpec {
  s.AllocateLoadBalancerNodePorts = &v
  return s
}

func (s *ListServiceServiceSpec) SetLoadBalancerClass(v string) *ListServiceServiceSpec {
  s.LoadBalancerClass = &v
  return s
}

func (s *ListServiceServiceSpec) SetInternalTrafficPolicy(v string) *ListServiceServiceSpec {
  s.InternalTrafficPolicy = &v
  return s
}

type ListServiceServicePort struct {
  // {"en":"The name of this port within the service. This must be a DNS_LABEL. All ports within a ListServiceServiceSpec must have unique names. When considering the endpoints for a ListServiceService, this must match the 'name' field in the EndpointPort. Optional if only one ListServiceServicePort is defined on this service", "zh_CN":"ListServiceService 中此端口的名称。这必须是 DNS_LABEL。 ListServiceServiceSpec 中的所有端口的名称都必须唯一。 在考虑 ListServiceService 的端点时，这一字段值必须与 EndpointPort 中的 name 字段相同。 如果此服务上仅定义一个 ListServiceServicePort，则为此字段为可选"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The IP protocol for this port. Supports TCP, UDP, and SCTP. Default is TCP", "zh_CN":"此端口的 IP 协议。支持 “TCP”、“UDP” 和 “SCTP”。默认为 TCP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol", "zh_CN":"此端口的应用协议，遵循标准的 Kubernetes 标签语法，无前缀名称按照 IANA 标准服务名称 （参见 RFC-6335 和 https://www.iana.org/assignments/service-names）。 非标准协议应该使用前缀名称，如 mycompany.com/my-custom-protocol"}
  AppProtocol *string `json:"appProtocol,omitempty" xml:"appProtocol,omitempty"`
  // {"en":"The port that will be exposed by this service", "zh_CN":"ListServiceService 将公开的端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field", "zh_CN":"在 ListServiceService 所针对的 Pod 上要访问的端口号或名称。 编号必须在 1 到 65535 的范围内。名称必须是 IANA_SVC_NAME。 如果此值是一个字符串，将在目标 Pod 的容器端口中作为命名端口进行查找。 如果未指定字段，则使用 “port” 字段的值（直接映射）。 对于 clusterIP 为 None 的服务，此字段将被忽略， 应忽略不设或设置为 “port” 字段的取值"}
  TargetPort *int32 `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
  // {"en":"The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this ListServiceService requires one. If this field is specified when creating a ListServiceService which does not need it, creation will fail. This field will be wiped when updating a ListServiceService to no longer need it (e.g. changing type from NodePort to ClusterIP).", "zh_CN":"当类型为 NodePort 或 LoadBalancer 时，ListServiceService 公开在节点上的端口， 通常由系统分配。如果指定了一个在范围内且未使用的值，则将使用该值，否则操作将失败。 如果在创建的 ListServiceService 需要该端口时未指定该字段，则会分配端口。 如果在创建不需要该端口的 Service时指定了该字段，则会创建失败。 当更新 ListServiceService 时，如果不再需要此字段（例如，将类型从 NodePort 更改为 ClusterIP），这个字段将被擦除"}
  NodePort *int32 `json:"nodePort,omitempty" xml:"nodePort,omitempty"`
}

func (s ListServiceServicePort) String() string {
  return tea.Prettify(s)
}

func (s ListServiceServicePort) GoString() string {
  return s.String()
}

func (s *ListServiceServicePort) SetName(v string) *ListServiceServicePort {
  s.Name = &v
  return s
}

func (s *ListServiceServicePort) SetProtocol(v string) *ListServiceServicePort {
  s.Protocol = &v
  return s
}

func (s *ListServiceServicePort) SetAppProtocol(v string) *ListServiceServicePort {
  s.AppProtocol = &v
  return s
}

func (s *ListServiceServicePort) SetPort(v int32) *ListServiceServicePort {
  s.Port = &v
  return s
}

func (s *ListServiceServicePort) SetTargetPort(v int32) *ListServiceServicePort {
  s.TargetPort = &v
  return s
}

func (s *ListServiceServicePort) SetNodePort(v int32) *ListServiceServicePort {
  s.NodePort = &v
  return s
}

type ListServiceSessionAffinityConfig struct {
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"clientIP 包含基于客户端 IP 的会话亲和性的配置"}
  ClientIP *ListServiceClientIPConfig `json:"clientIP,omitempty" xml:"clientIP,omitempty"`
}

func (s ListServiceSessionAffinityConfig) String() string {
  return tea.Prettify(s)
}

func (s ListServiceSessionAffinityConfig) GoString() string {
  return s.String()
}

func (s *ListServiceSessionAffinityConfig) SetClientIP(v *ListServiceClientIPConfig) *ListServiceSessionAffinityConfig {
  s.ClientIP = v
  return s
}

type ListServiceClientIPConfig struct {
  // {"en":"timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).", "zh_CN":"timeoutSeconds 指定 ClientIP 类型会话的维系时间秒数。 如果 ServiceAffinity == 'ClientIP'，则该值必须 >0 && <=86400（1 天）。默认值为 10800（3 小时）"}
  TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s ListServiceClientIPConfig) String() string {
  return tea.Prettify(s)
}

func (s ListServiceClientIPConfig) GoString() string {
  return s.String()
}

func (s *ListServiceClientIPConfig) SetTimeoutSeconds(v int32) *ListServiceClientIPConfig {
  s.TimeoutSeconds = &v
  return s
}

type ListServiceObjectMeta struct {
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
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 ListServiceService 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*ListServiceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*ListServiceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s ListServiceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s ListServiceObjectMeta) GoString() string {
  return s.String()
}

func (s *ListServiceObjectMeta) SetName(v string) *ListServiceObjectMeta {
  s.Name = &v
  return s
}

func (s *ListServiceObjectMeta) SetGenerateName(v string) *ListServiceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *ListServiceObjectMeta) SetNamespace(v string) *ListServiceObjectMeta {
  s.Namespace = &v
  return s
}

func (s *ListServiceObjectMeta) SetSelfLink(v string) *ListServiceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *ListServiceObjectMeta) SetUid(v string) *ListServiceObjectMeta {
  s.Uid = &v
  return s
}

func (s *ListServiceObjectMeta) SetResourceVersion(v string) *ListServiceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListServiceObjectMeta) SetGeneration(v int64) *ListServiceObjectMeta {
  s.Generation = &v
  return s
}

func (s *ListServiceObjectMeta) SetCreationTimestamp(v string) *ListServiceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *ListServiceObjectMeta) SetDeletionTimestamp(v string) *ListServiceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *ListServiceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *ListServiceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *ListServiceObjectMeta) SetLabels(v map[string]*string) *ListServiceObjectMeta {
  s.Labels = v
  return s
}

func (s *ListServiceObjectMeta) SetAnnotations(v map[string]*string) *ListServiceObjectMeta {
  s.Annotations = v
  return s
}

func (s *ListServiceObjectMeta) SetOwnerReferences(v []*ListServiceOwnerReference) *ListServiceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *ListServiceObjectMeta) SetFinalizers(v []*string) *ListServiceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *ListServiceObjectMeta) SetClusterName(v string) *ListServiceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *ListServiceObjectMeta) SetManagedFields(v []*ListServiceManagedFieldsEntry) *ListServiceObjectMeta {
  s.ManagedFields = v
  return s
}

type ListServiceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this ListServiceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'ListServiceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“ListServiceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"ListServiceFieldsV1 holds the first JSON version format as described in the 'ListServiceFieldsV1' type", "zh_CN":"ListServiceFieldsV1 包含类型 “ListServiceFieldsV1” 中描述的第一个 JSON 版本格式"}
  ListServiceFieldsV1 *ListServiceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s ListServiceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s ListServiceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *ListServiceManagedFieldsEntry) SetManager(v string) *ListServiceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *ListServiceManagedFieldsEntry) SetOperation(v string) *ListServiceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *ListServiceManagedFieldsEntry) SetApiVersion(v string) *ListServiceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *ListServiceManagedFieldsEntry) SetTime(v string) *ListServiceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *ListServiceManagedFieldsEntry) SetFieldsType(v string) *ListServiceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *ListServiceManagedFieldsEntry) SetFieldsV1(v *ListServiceFieldsV1) *ListServiceManagedFieldsEntry {
  s.ListServiceFieldsV1 = v
  return s
}

func (s *ListServiceManagedFieldsEntry) SetSubresource(v string) *ListServiceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type ListServiceFieldsV1 struct {
}

func (s ListServiceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s ListServiceFieldsV1) GoString() string {
  return s.String()
}

type ListServiceOwnerReference struct {
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

func (s ListServiceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s ListServiceOwnerReference) GoString() string {
  return s.String()
}

func (s *ListServiceOwnerReference) SetApiVersion(v string) *ListServiceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *ListServiceOwnerReference) SetKind(v string) *ListServiceOwnerReference {
  s.Kind = &v
  return s
}

func (s *ListServiceOwnerReference) SetName(v string) *ListServiceOwnerReference {
  s.Name = &v
  return s
}

func (s *ListServiceOwnerReference) SetUid(v string) *ListServiceOwnerReference {
  s.Uid = &v
  return s
}

func (s *ListServiceOwnerReference) SetController(v bool) *ListServiceOwnerReference {
  s.Controller = &v
  return s
}

func (s *ListServiceOwnerReference) SetBlockOwnerDeletion(v bool) *ListServiceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type VMPQueryAvailableCidrsRequest struct {
}

func (s VMPQueryAvailableCidrsRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsRequest) GoString() string {
  return s.String()
}

type VMPQueryAvailableCidrsRequestHeader struct {
}

func (s VMPQueryAvailableCidrsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryAvailableCidrsPaths struct {
}

func (s VMPQueryAvailableCidrsPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsPaths) GoString() string {
  return s.String()
}

type VMPQueryAvailableCidrsParameters struct {
  // {"en":"node name","zh_CN":"节点英文名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
}

func (s VMPQueryAvailableCidrsParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsParameters) SetNodeName(v string) *VMPQueryAvailableCidrsParameters {
  s.NodeName = &v
  return s
}

type VMPQueryAvailableCidrsResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *VMPQueryAvailableCidrsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s VMPQueryAvailableCidrsResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsResponse) SetCode(v string) *VMPQueryAvailableCidrsResponse {
  s.Code = &v
  return s
}

func (s *VMPQueryAvailableCidrsResponse) SetData(v *VMPQueryAvailableCidrsResponseData) *VMPQueryAvailableCidrsResponse {
  s.Data = v
  return s
}

func (s *VMPQueryAvailableCidrsResponse) SetMessage(v string) *VMPQueryAvailableCidrsResponse {
  s.Message = &v
  return s
}

type VMPQueryAvailableCidrsResponseData struct {
  // {"en":"available cidrs","zh_CN":"可用的cidr列表"}
  Cidrs []*string `json:"cidrs,omitempty" xml:"cidrs,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryAvailableCidrsResponseData) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsResponseData) GoString() string {
  return s.String()
}

func (s *VMPQueryAvailableCidrsResponseData) SetCidrs(v []*string) *VMPQueryAvailableCidrsResponseData {
  s.Cidrs = v
  return s
}

type VMPQueryAvailableCidrsResponseHeader struct {
}

func (s VMPQueryAvailableCidrsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryAvailableCidrsResponseHeader) GoString() string {
  return s.String()
}




type UnassignEdgePrivateIPRequest struct {
  // {"en":"Target virtual machine id","zh_CN":"目标实例id"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"Additional Ip to unbind","zh_CN":"要解除绑定的额外内网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s UnassignEdgePrivateIPRequest) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPRequest) GoString() string {
  return s.String()
}

func (s *UnassignEdgePrivateIPRequest) SetInstanceId(v string) *UnassignEdgePrivateIPRequest {
  s.InstanceId = &v
  return s
}

func (s *UnassignEdgePrivateIPRequest) SetEdgeIps(v []*string) *UnassignEdgePrivateIPRequest {
  s.EdgeIps = v
  return s
}

type UnassignEdgePrivateIPRequestHeader struct {
}

func (s UnassignEdgePrivateIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPRequestHeader) GoString() string {
  return s.String()
}

type UnassignEdgePrivateIPPaths struct {
}

func (s UnassignEdgePrivateIPPaths) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPPaths) GoString() string {
  return s.String()
}

type UnassignEdgePrivateIPParameters struct {
}

func (s UnassignEdgePrivateIPParameters) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPParameters) GoString() string {
  return s.String()
}

type UnassignEdgePrivateIPResponse struct {
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *UnassignEdgePrivateIPResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s UnassignEdgePrivateIPResponse) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPResponse) GoString() string {
  return s.String()
}

func (s *UnassignEdgePrivateIPResponse) SetCode(v string) *UnassignEdgePrivateIPResponse {
  s.Code = &v
  return s
}

func (s *UnassignEdgePrivateIPResponse) SetMessage(v string) *UnassignEdgePrivateIPResponse {
  s.Message = &v
  return s
}

func (s *UnassignEdgePrivateIPResponse) SetData(v *UnassignEdgePrivateIPResponseData) *UnassignEdgePrivateIPResponse {
  s.Data = v
  return s
}

type UnassignEdgePrivateIPResponseData struct {
  // {"en":"target virtual machine id","zh_CN":"目标实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"Additional IP that is bound to the instance","zh_CN":"已绑定到目标实例的额外IP"}
  EdgeIps []*UnassignEdgePrivateIPResponseDataEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s UnassignEdgePrivateIPResponseData) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPResponseData) GoString() string {
  return s.String()
}

func (s *UnassignEdgePrivateIPResponseData) SetInstanceId(v string) *UnassignEdgePrivateIPResponseData {
  s.InstanceId = &v
  return s
}

func (s *UnassignEdgePrivateIPResponseData) SetEdgeIps(v []*UnassignEdgePrivateIPResponseDataEdgeIps) *UnassignEdgePrivateIPResponseData {
  s.EdgeIps = v
  return s
}

type UnassignEdgePrivateIPResponseDataEdgeIps struct     {
  // {"en":"Edge private ip","zh_CN":"额外内网IP"}
  EdgePrivateIp *string `json:"edgePrivateIp,omitempty" xml:"edgePrivateIp,omitempty" require:"true"`
  // {"en":"Edge public IP","zh_CN":"额外公网IP"}
  EdgePublicIp *string `json:"edgePublicIp,omitempty" xml:"edgePublicIp,omitempty" require:"true"`
}

func (s UnassignEdgePrivateIPResponseDataEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPResponseDataEdgeIps) GoString() string {
  return s.String()
}

func (s *UnassignEdgePrivateIPResponseDataEdgeIps) SetEdgePrivateIp(v string) *UnassignEdgePrivateIPResponseDataEdgeIps {
  s.EdgePrivateIp = &v
  return s
}

func (s *UnassignEdgePrivateIPResponseDataEdgeIps) SetEdgePublicIp(v string) *UnassignEdgePrivateIPResponseDataEdgeIps {
  s.EdgePublicIp = &v
  return s
}

type UnassignEdgePrivateIPResponseHeader struct {
}

func (s UnassignEdgePrivateIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UnassignEdgePrivateIPResponseHeader) GoString() string {
  return s.String()
}




type VMPReleaseEdgeIPRequest struct {
  // {"en":"additional IP to be released","zh_CN":"要释放的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPReleaseEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPReleaseEdgeIPRequest) SetEdgeIps(v []*string) *VMPReleaseEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type VMPReleaseEdgeIPRequestHeader struct {
}

func (s VMPReleaseEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPReleaseEdgeIPPaths struct {
}

func (s VMPReleaseEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPReleaseEdgeIPParameters struct {
}

func (s VMPReleaseEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPReleaseEdgeIPResponse struct {
  // {"en":"Error Message Set","zh_CN":"错误信息集"}
  BatchErrorMsg []*VMPReleaseEdgeIPResponseBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s VMPReleaseEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPReleaseEdgeIPResponse) SetBatchErrorMsg(v []*VMPReleaseEdgeIPResponseBatchErrorMsg) *VMPReleaseEdgeIPResponse {
  s.BatchErrorMsg = v
  return s
}

type VMPReleaseEdgeIPResponseBatchErrorMsg struct     {
  // {"en":"Ip","zh_CN":"Ip"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s VMPReleaseEdgeIPResponseBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPResponseBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *VMPReleaseEdgeIPResponseBatchErrorMsg) SetKey(v string) *VMPReleaseEdgeIPResponseBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *VMPReleaseEdgeIPResponseBatchErrorMsg) SetMsg(v string) *VMPReleaseEdgeIPResponseBatchErrorMsg {
  s.Msg = &v
  return s
}

type VMPReleaseEdgeIPResponseHeader struct {
}

func (s VMPReleaseEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type UpdateNetworkPolicyRequest struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *UpdateNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this UpdateNetworkPolicyNetworkPolicy.", "zh_CN": "UpdateNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *UpdateNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s UpdateNetworkPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyRequest) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyRequest) SetApiVersion(v string) *UpdateNetworkPolicyRequest {
  s.ApiVersion = &v
  return s
}

func (s *UpdateNetworkPolicyRequest) SetKind(v string) *UpdateNetworkPolicyRequest {
  s.Kind = &v
  return s
}

func (s *UpdateNetworkPolicyRequest) SetMetadata(v *UpdateNetworkPolicyObjectMeta) *UpdateNetworkPolicyRequest {
  s.Metadata = v
  return s
}

func (s *UpdateNetworkPolicyRequest) SetSpec(v *UpdateNetworkPolicyNetworkPolicySpec) *UpdateNetworkPolicyRequest {
  s.Spec = v
  return s
}

type UpdateNetworkPolicyResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"networkPolicy", "zh_CN":"网络策略"}
  Data *UpdateNetworkPolicyNetworkPolicy `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateNetworkPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyResponse) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyResponse) SetCode(v int64) *UpdateNetworkPolicyResponse {
  s.Code = &v
  return s
}

func (s *UpdateNetworkPolicyResponse) SetMsg(v string) *UpdateNetworkPolicyResponse {
  s.Msg = &v
  return s
}

func (s *UpdateNetworkPolicyResponse) SetRequestId(v string) *UpdateNetworkPolicyResponse {
  s.RequestId = &v
  return s
}

func (s *UpdateNetworkPolicyResponse) SetData(v *UpdateNetworkPolicyNetworkPolicy) *UpdateNetworkPolicyResponse {
  s.Data = v
  return s
}

type UpdateNetworkPolicyPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"networkPolicy name", "zh_CN":"networkPolicy 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateNetworkPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyPaths) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyPaths) SetNamespace(v string) *UpdateNetworkPolicyPaths {
  s.Namespace = &v
  return s
}

func (s *UpdateNetworkPolicyPaths) SetName(v string) *UpdateNetworkPolicyPaths {
  s.Name = &v
  return s
}

type UpdateNetworkPolicyParameters struct {
}

func (s UpdateNetworkPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyParameters) GoString() string {
  return s.String()
}

type UpdateNetworkPolicyRequestHeader struct {
}

func (s UpdateNetworkPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyRequestHeader) GoString() string {
  return s.String()
}

type UpdateNetworkPolicyResponseHeader struct {
}

func (s UpdateNetworkPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyResponseHeader) GoString() string {
  return s.String()
}

type UpdateNetworkPolicyNetworkPolicy struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *UpdateNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this UpdateNetworkPolicyNetworkPolicy.", "zh_CN": "UpdateNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *UpdateNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s UpdateNetworkPolicyNetworkPolicy) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNetworkPolicy) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNetworkPolicy) SetApiVersion(v string) *UpdateNetworkPolicyNetworkPolicy {
  s.ApiVersion = &v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicy) SetKind(v string) *UpdateNetworkPolicyNetworkPolicy {
  s.Kind = &v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicy) SetMetadata(v *UpdateNetworkPolicyObjectMeta) *UpdateNetworkPolicyNetworkPolicy {
  s.Metadata = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicy) SetSpec(v *UpdateNetworkPolicyNetworkPolicySpec) *UpdateNetworkPolicyNetworkPolicy {
  s.Spec = v
  return s
}

type UpdateNetworkPolicyNetworkPolicySpec struct {
  // {"en": "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the UpdateNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this UpdateNetworkPolicyNetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8", "zh_CN": "出网规则"}
  Egress []*UpdateNetworkPolicyNetworkPolicyEgressRule `json:"egress,omitempty" xml:"egress,omitempty" type:"Repeated"`
  // {"en": "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the UpdateNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this UpdateNetworkPolicyNetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)", "zh_CN": "入网规则"}
  Ingress []*UpdateNetworkPolicyNetworkPolicyIngressRule `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
  // {"en": "Selects the pods to which this UpdateNetworkPolicyNetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.", "zh_CN": "限制pod的选择器"}
  PodSelector *UpdateNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty" require:"true"`
  // {"en": "List of rule types that the UpdateNetworkPolicyNetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ Egress ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include Egress (since such a policy would not include an Egress section and would otherwise default to just [ Ingress ]). This field is beta-level in 1.8", "zh_CN": "策略类型"}
  PolicyTypes []*string `json:"policyTypes,omitempty" xml:"policyTypes,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyNetworkPolicySpec) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNetworkPolicySpec) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNetworkPolicySpec) SetEgress(v []*UpdateNetworkPolicyNetworkPolicyEgressRule) *UpdateNetworkPolicyNetworkPolicySpec {
  s.Egress = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicySpec) SetIngress(v []*UpdateNetworkPolicyNetworkPolicyIngressRule) *UpdateNetworkPolicyNetworkPolicySpec {
  s.Ingress = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicySpec) SetPodSelector(v *UpdateNetworkPolicyPodLabelSelector) *UpdateNetworkPolicyNetworkPolicySpec {
  s.PodSelector = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicySpec) SetPolicyTypes(v []*string) *UpdateNetworkPolicyNetworkPolicySpec {
  s.PolicyTypes = v
  return s
}

type UpdateNetworkPolicyNetworkPolicyIngressRule struct {
  // {"en": "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.", "zh_CN": "入网规则信息"}
  From []*UpdateNetworkPolicyNetworkPolicyPeer `json:"from,omitempty" xml:"from,omitempty" type:"Repeated"`
  // {"en": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*UpdateNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyNetworkPolicyIngressRule) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNetworkPolicyIngressRule) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNetworkPolicyIngressRule) SetFrom(v []*UpdateNetworkPolicyNetworkPolicyPeer) *UpdateNetworkPolicyNetworkPolicyIngressRule {
  s.From = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicyIngressRule) SetPorts(v []*UpdateNetworkPolicyNetworkPolicyPort) *UpdateNetworkPolicyNetworkPolicyIngressRule {
  s.Ports = v
  return s
}

type UpdateNetworkPolicyNetworkPolicyEgressRule struct {
  // {"en": "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*UpdateNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en": "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.", "zh_CN": "出网规则信息"}
  To []*UpdateNetworkPolicyNetworkPolicyPeer `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyNetworkPolicyEgressRule) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNetworkPolicyEgressRule) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNetworkPolicyEgressRule) SetPorts(v []*UpdateNetworkPolicyNetworkPolicyPort) *UpdateNetworkPolicyNetworkPolicyEgressRule {
  s.Ports = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicyEgressRule) SetTo(v []*UpdateNetworkPolicyNetworkPolicyPeer) *UpdateNetworkPolicyNetworkPolicyEgressRule {
  s.To = v
  return s
}

type UpdateNetworkPolicyNetworkPolicyPeer struct {
  // {"en": "UpdateNetworkPolicyIPBlock defines policy on a particular UpdateNetworkPolicyIPBlock. If this field is set then neither of the other fields can be.", "zh_CN": "IP规则"}
  IpBlock *UpdateNetworkPolicyIPBlock `json:"ipBlock,omitempty" xml:"ipBlock,omitempty"`
  // {"en": "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.If PodSelector is also set, then the UpdateNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.", "zh_CN": "namespace选择器"}
  NamespaceSelector *UpdateNetworkPolicyNsLabelSelector `json:"namespaceSelector,omitempty" xml:"namespaceSelector,omitempty"`
  // {"en": "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.If NamespaceSelector is also set, then the UpdateNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.", "zh_CN": "pod选择器"}
  PodSelector *UpdateNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty"`
}

func (s UpdateNetworkPolicyNetworkPolicyPeer) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNetworkPolicyPeer) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNetworkPolicyPeer) SetIpBlock(v *UpdateNetworkPolicyIPBlock) *UpdateNetworkPolicyNetworkPolicyPeer {
  s.IpBlock = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicyPeer) SetNamespaceSelector(v *UpdateNetworkPolicyNsLabelSelector) *UpdateNetworkPolicyNetworkPolicyPeer {
  s.NamespaceSelector = v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicyPeer) SetPodSelector(v *UpdateNetworkPolicyPodLabelSelector) *UpdateNetworkPolicyNetworkPolicyPeer {
  s.PodSelector = v
  return s
}

type UpdateNetworkPolicyIPBlock struct {
  // {"en": "CIDR is a string representing the IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64", "zh_CN": "生效IP网段"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en": "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64 Except values will be rejected if they are outside the CIDR range", "zh_CN": "例外IP网段"}
  Except []*string `json:"except,omitempty" xml:"except,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyIPBlock) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyIPBlock) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyIPBlock) SetCidr(v string) *UpdateNetworkPolicyIPBlock {
  s.Cidr = &v
  return s
}

func (s *UpdateNetworkPolicyIPBlock) SetExcept(v []*string) *UpdateNetworkPolicyIPBlock {
  s.Except = v
  return s
}

type UpdateNetworkPolicyNetworkPolicyPort struct {
  // {"en": "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.", "zh_CN": "端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
  // {"en": "The protocol (TCP, UDP) which traffic must match. If not specified, this field defaults to TCP.", "zh_CN": "协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateNetworkPolicyNetworkPolicyPort) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNetworkPolicyPort) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNetworkPolicyPort) SetPort(v string) *UpdateNetworkPolicyNetworkPolicyPort {
  s.Port = &v
  return s
}

func (s *UpdateNetworkPolicyNetworkPolicyPort) SetProtocol(v string) *UpdateNetworkPolicyNetworkPolicyPort {
  s.Protocol = &v
  return s
}

type UpdateNetworkPolicyPodLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
}

func (s UpdateNetworkPolicyPodLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyPodLabelSelector) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyPodLabelSelector) SetMatchLabels(v map[string]*string) *UpdateNetworkPolicyPodLabelSelector {
  s.MatchLabels = v
  return s
}

type UpdateNetworkPolicyNsLabelSelector struct {
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*UpdateNetworkPolicyLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyNsLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyNsLabelSelector) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyNsLabelSelector) SetMatchExpressions(v []*UpdateNetworkPolicyLabelSelectorRequirement) *UpdateNetworkPolicyNsLabelSelector {
  s.MatchExpressions = v
  return s
}

type UpdateNetworkPolicyLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyLabelSelectorRequirement) SetKey(v string) *UpdateNetworkPolicyLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *UpdateNetworkPolicyLabelSelectorRequirement) SetOperator(v string) *UpdateNetworkPolicyLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *UpdateNetworkPolicyLabelSelectorRequirement) SetValues(v []*string) *UpdateNetworkPolicyLabelSelectorRequirement {
  s.Values = v
  return s
}

type UpdateNetworkPolicyObjectMeta struct {
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
  OwnerReferences []*UpdateNetworkPolicyOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*UpdateNetworkPolicyManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s UpdateNetworkPolicyObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyObjectMeta) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyObjectMeta) SetName(v string) *UpdateNetworkPolicyObjectMeta {
  s.Name = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetGenerateName(v string) *UpdateNetworkPolicyObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetNamespace(v string) *UpdateNetworkPolicyObjectMeta {
  s.Namespace = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetSelfLink(v string) *UpdateNetworkPolicyObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetUid(v string) *UpdateNetworkPolicyObjectMeta {
  s.Uid = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetResourceVersion(v string) *UpdateNetworkPolicyObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetGeneration(v int64) *UpdateNetworkPolicyObjectMeta {
  s.Generation = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetCreationTimestamp(v string) *UpdateNetworkPolicyObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetDeletionTimestamp(v string) *UpdateNetworkPolicyObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetDeletionGracePeriodSeconds(v int64) *UpdateNetworkPolicyObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetLabels(v map[string]*string) *UpdateNetworkPolicyObjectMeta {
  s.Labels = v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetAnnotations(v map[string]*string) *UpdateNetworkPolicyObjectMeta {
  s.Annotations = v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetOwnerReferences(v []*UpdateNetworkPolicyOwnerReference) *UpdateNetworkPolicyObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetFinalizers(v []*string) *UpdateNetworkPolicyObjectMeta {
  s.Finalizers = v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetClusterName(v string) *UpdateNetworkPolicyObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *UpdateNetworkPolicyObjectMeta) SetManagedFields(v []*UpdateNetworkPolicyManagedFieldsEntry) *UpdateNetworkPolicyObjectMeta {
  s.ManagedFields = v
  return s
}

type UpdateNetworkPolicyOwnerReference struct {
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

func (s UpdateNetworkPolicyOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyOwnerReference) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyOwnerReference) SetApiVersion(v string) *UpdateNetworkPolicyOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *UpdateNetworkPolicyOwnerReference) SetKind(v string) *UpdateNetworkPolicyOwnerReference {
  s.Kind = &v
  return s
}

func (s *UpdateNetworkPolicyOwnerReference) SetName(v string) *UpdateNetworkPolicyOwnerReference {
  s.Name = &v
  return s
}

func (s *UpdateNetworkPolicyOwnerReference) SetUid(v string) *UpdateNetworkPolicyOwnerReference {
  s.Uid = &v
  return s
}

func (s *UpdateNetworkPolicyOwnerReference) SetController(v bool) *UpdateNetworkPolicyOwnerReference {
  s.Controller = &v
  return s
}

func (s *UpdateNetworkPolicyOwnerReference) SetBlockOwnerDeletion(v bool) *UpdateNetworkPolicyOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type UpdateNetworkPolicyManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this UpdateNetworkPolicyManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'UpdateNetworkPolicyFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“UpdateNetworkPolicyFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"UpdateNetworkPolicyFieldsV1 holds the first JSON version format as described in the 'UpdateNetworkPolicyFieldsV1' type", "zh_CN":"UpdateNetworkPolicyFieldsV1 包含类型 “UpdateNetworkPolicyFieldsV1” 中描述的第一个 JSON 版本格式"}
  UpdateNetworkPolicyFieldsV1 *UpdateNetworkPolicyFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s UpdateNetworkPolicyManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetManager(v string) *UpdateNetworkPolicyManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetOperation(v string) *UpdateNetworkPolicyManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetApiVersion(v string) *UpdateNetworkPolicyManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetTime(v string) *UpdateNetworkPolicyManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetFieldsType(v string) *UpdateNetworkPolicyManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetFieldsV1(v *UpdateNetworkPolicyFieldsV1) *UpdateNetworkPolicyManagedFieldsEntry {
  s.UpdateNetworkPolicyFieldsV1 = v
  return s
}

func (s *UpdateNetworkPolicyManagedFieldsEntry) SetSubresource(v string) *UpdateNetworkPolicyManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type UpdateNetworkPolicyFieldsV1 struct {
}

func (s UpdateNetworkPolicyFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s UpdateNetworkPolicyFieldsV1) GoString() string {
  return s.String()
}




type ReleaseEdgePrivateIPRequest struct {
  // {"en":"node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"additional IP to be released","zh_CN":"要释放的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s ReleaseEdgePrivateIPRequest) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPRequest) GoString() string {
  return s.String()
}

func (s *ReleaseEdgePrivateIPRequest) SetNodeName(v string) *ReleaseEdgePrivateIPRequest {
  s.NodeName = &v
  return s
}

func (s *ReleaseEdgePrivateIPRequest) SetEdgeIps(v []*string) *ReleaseEdgePrivateIPRequest {
  s.EdgeIps = v
  return s
}

type ReleaseEdgePrivateIPRequestHeader struct {
}

func (s ReleaseEdgePrivateIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPRequestHeader) GoString() string {
  return s.String()
}

type ReleaseEdgePrivateIPPaths struct {
}

func (s ReleaseEdgePrivateIPPaths) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPPaths) GoString() string {
  return s.String()
}

type ReleaseEdgePrivateIPParameters struct {
}

func (s ReleaseEdgePrivateIPParameters) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPParameters) GoString() string {
  return s.String()
}

type ReleaseEdgePrivateIPResponse struct {
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *ReleaseEdgePrivateIPResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ReleaseEdgePrivateIPResponse) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPResponse) GoString() string {
  return s.String()
}

func (s *ReleaseEdgePrivateIPResponse) SetCode(v string) *ReleaseEdgePrivateIPResponse {
  s.Code = &v
  return s
}

func (s *ReleaseEdgePrivateIPResponse) SetMessage(v string) *ReleaseEdgePrivateIPResponse {
  s.Message = &v
  return s
}

func (s *ReleaseEdgePrivateIPResponse) SetData(v *ReleaseEdgePrivateIPResponseData) *ReleaseEdgePrivateIPResponse {
  s.Data = v
  return s
}

type ReleaseEdgePrivateIPResponseData struct {
  // {"en":"Release IP error message","zh_CN":"释放Ip错误信息"}
  BatchErrorMsg []*ReleaseEdgePrivateIPResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s ReleaseEdgePrivateIPResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPResponseData) GoString() string {
  return s.String()
}

func (s *ReleaseEdgePrivateIPResponseData) SetBatchErrorMsg(v []*ReleaseEdgePrivateIPResponseDataBatchErrorMsg) *ReleaseEdgePrivateIPResponseData {
  s.BatchErrorMsg = v
  return s
}

type ReleaseEdgePrivateIPResponseDataBatchErrorMsg struct     {
  // {"en":"Ip","zh_CN":"Ip"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Ip Error message","zh_CN":"Ip错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ReleaseEdgePrivateIPResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *ReleaseEdgePrivateIPResponseDataBatchErrorMsg) SetKey(v string) *ReleaseEdgePrivateIPResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *ReleaseEdgePrivateIPResponseDataBatchErrorMsg) SetMsg(v string) *ReleaseEdgePrivateIPResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type ReleaseEdgePrivateIPResponseHeader struct {
}

func (s ReleaseEdgePrivateIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReleaseEdgePrivateIPResponseHeader) GoString() string {
  return s.String()
}




type DeleteNetworkPolicyRequest struct {
}

func (s DeleteNetworkPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyRequest) GoString() string {
  return s.String()
}

type DeleteNetworkPolicyResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"status"}
  Data *DeleteNetworkPolicyStatus `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteNetworkPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyResponse) GoString() string {
  return s.String()
}

func (s *DeleteNetworkPolicyResponse) SetCode(v int64) *DeleteNetworkPolicyResponse {
  s.Code = &v
  return s
}

func (s *DeleteNetworkPolicyResponse) SetMsg(v string) *DeleteNetworkPolicyResponse {
  s.Msg = &v
  return s
}

func (s *DeleteNetworkPolicyResponse) SetRequestId(v string) *DeleteNetworkPolicyResponse {
  s.RequestId = &v
  return s
}

func (s *DeleteNetworkPolicyResponse) SetData(v *DeleteNetworkPolicyStatus) *DeleteNetworkPolicyResponse {
  s.Data = v
  return s
}

type DeleteNetworkPolicyPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"networkPolicy name", "zh_CN":"networkPolicy 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteNetworkPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyPaths) GoString() string {
  return s.String()
}

func (s *DeleteNetworkPolicyPaths) SetNamespace(v string) *DeleteNetworkPolicyPaths {
  s.Namespace = &v
  return s
}

func (s *DeleteNetworkPolicyPaths) SetName(v string) *DeleteNetworkPolicyPaths {
  s.Name = &v
  return s
}

type DeleteNetworkPolicyParameters struct {
}

func (s DeleteNetworkPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyParameters) GoString() string {
  return s.String()
}

type DeleteNetworkPolicyRequestHeader struct {
}

func (s DeleteNetworkPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyRequestHeader) GoString() string {
  return s.String()
}

type DeleteNetworkPolicyResponseHeader struct {
}

func (s DeleteNetworkPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyResponseHeader) GoString() string {
  return s.String()
}

type DeleteNetworkPolicyStatus struct {
  // {"en":"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values", "zh_CN":"APIVersion 定义对象表示的版本化模式。 服务器应将已识别的模式转换为最新的内部值，并可能拒绝无法识别的值"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase", "zh_CN":"Kind 是一个字符串值，表示此对象表示的 REST 资源。 服务器可以从客户端提交请求的端点推断出这一点。 无法更新。驼峰式规则"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"DeleteNetworkPolicyStatus of the operation. One of: 'Success' or 'Failure'", "zh_CN":"操作状态。“Success”或“Failure” 之一"}
  DeleteNetworkPolicyStatus *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Suggested HTTP return code for this status, 0 if not set", "zh_CN":"此状态的建议 HTTP 返回代码，如果未设置，则为 0"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Extended data associated with the reason. Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type", "zh_CN":"与原因（Reason）相关的扩展数据。每个原因都可以定义自己的扩展细节。 此字段是可选的，并且不保证返回的数据符合任何模式，除非由原因类型定义"}
  Details *DeleteNetworkPolicyStatusDetails `json:"details,omitempty" xml:"details,omitempty" require:"true"`
}

func (s DeleteNetworkPolicyStatus) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyStatus) GoString() string {
  return s.String()
}

func (s *DeleteNetworkPolicyStatus) SetApiVersion(v string) *DeleteNetworkPolicyStatus {
  s.ApiVersion = &v
  return s
}

func (s *DeleteNetworkPolicyStatus) SetKind(v string) *DeleteNetworkPolicyStatus {
  s.Kind = &v
  return s
}

func (s *DeleteNetworkPolicyStatus) SetStatus(v string) *DeleteNetworkPolicyStatus {
  s.DeleteNetworkPolicyStatus = &v
  return s
}

func (s *DeleteNetworkPolicyStatus) SetCode(v int32) *DeleteNetworkPolicyStatus {
  s.Code = &v
  return s
}

func (s *DeleteNetworkPolicyStatus) SetDetails(v *DeleteNetworkPolicyStatusDetails) *DeleteNetworkPolicyStatus {
  s.Details = v
  return s
}

type DeleteNetworkPolicyStatusDetails struct {
  // {"en":"The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)", "zh_CN":"与状态 StatusReason 关联的资源的名称属性（当有一个可以描述的名称时）"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind", "zh_CN":"与状态 StatusReason 关联的资源的种类属性"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"The group attribute of the resource associated with the status StatusReason", "zh_CN":"与状态 StatusReason 关联的资源的组属性"}
  Group *string `json:"group,omitempty" xml:"group,omitempty" require:"true"`
  // {"en":"UID of the resource. (when there is a single resource which can be described)", "zh_CN":"资源的 UID（当有单个可以描述的资源时）"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty" require:"true"`
}

func (s DeleteNetworkPolicyStatusDetails) String() string {
  return tea.Prettify(s)
}

func (s DeleteNetworkPolicyStatusDetails) GoString() string {
  return s.String()
}

func (s *DeleteNetworkPolicyStatusDetails) SetName(v string) *DeleteNetworkPolicyStatusDetails {
  s.Name = &v
  return s
}

func (s *DeleteNetworkPolicyStatusDetails) SetKind(v string) *DeleteNetworkPolicyStatusDetails {
  s.Kind = &v
  return s
}

func (s *DeleteNetworkPolicyStatusDetails) SetGroup(v string) *DeleteNetworkPolicyStatusDetails {
  s.Group = &v
  return s
}

func (s *DeleteNetworkPolicyStatusDetails) SetUid(v string) *DeleteNetworkPolicyStatusDetails {
  s.Uid = &v
  return s
}




type CreateIngressRequest struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  ControllerName *string `json:"controllerName,omitempty" xml:"controllerName,omitempty" require:"true"`
  // {"en":"clusters", "zh_CN":"集群列表"}
  Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"ingress", "zh_CN":"路由"}
  CreateIngressIngress *CreateIngressIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true"`
}

func (s CreateIngressRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressRequest) GoString() string {
  return s.String()
}

func (s *CreateIngressRequest) SetControllerName(v string) *CreateIngressRequest {
  s.ControllerName = &v
  return s
}

func (s *CreateIngressRequest) SetClusters(v []*string) *CreateIngressRequest {
  s.Clusters = v
  return s
}

func (s *CreateIngressRequest) SetIngress(v *CreateIngressIngress) *CreateIngressRequest {
  s.CreateIngressIngress = v
  return s
}

type CreateIngressResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress object", "zh_CN":"路由对象"}
  Data *CreateIngressCustomIngress `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateIngressResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressResponse) GoString() string {
  return s.String()
}

func (s *CreateIngressResponse) SetCode(v int64) *CreateIngressResponse {
  s.Code = &v
  return s
}

func (s *CreateIngressResponse) SetMsg(v string) *CreateIngressResponse {
  s.Msg = &v
  return s
}

func (s *CreateIngressResponse) SetRequestId(v string) *CreateIngressResponse {
  s.RequestId = &v
  return s
}

func (s *CreateIngressResponse) SetData(v *CreateIngressCustomIngress) *CreateIngressResponse {
  s.Data = v
  return s
}

type CreateIngressPaths struct {
}

func (s CreateIngressPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressPaths) GoString() string {
  return s.String()
}

type CreateIngressParameters struct {
}

func (s CreateIngressParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressParameters) GoString() string {
  return s.String()
}

type CreateIngressRequestHeader struct {
}

func (s CreateIngressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressRequestHeader) GoString() string {
  return s.String()
}

type CreateIngressResponseHeader struct {
}

func (s CreateIngressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressResponseHeader) GoString() string {
  return s.String()
}

type CreateIngressCustomIngress struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  ControllerName *string `json:"controllerName,omitempty" xml:"controllerName,omitempty" require:"true"`
  // {"en":"clusters", "zh_CN":"集群列表"}
  Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"ingress", "zh_CN":"路由"}
  CreateIngressIngress *CreateIngressIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true"`
}

func (s CreateIngressCustomIngress) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressCustomIngress) GoString() string {
  return s.String()
}

func (s *CreateIngressCustomIngress) SetControllerName(v string) *CreateIngressCustomIngress {
  s.ControllerName = &v
  return s
}

func (s *CreateIngressCustomIngress) SetClusters(v []*string) *CreateIngressCustomIngress {
  s.Clusters = v
  return s
}

func (s *CreateIngressCustomIngress) SetIngress(v *CreateIngressIngress) *CreateIngressCustomIngress {
  s.CreateIngressIngress = v
  return s
}

type CreateIngressIngress struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreateIngressObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"ingress desired", "zh_CN":"路由期望属性"}
  Spec *CreateIngressIngressSpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s CreateIngressIngress) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngress) GoString() string {
  return s.String()
}

func (s *CreateIngressIngress) SetApiVersion(v string) *CreateIngressIngress {
  s.ApiVersion = &v
  return s
}

func (s *CreateIngressIngress) SetKind(v string) *CreateIngressIngress {
  s.Kind = &v
  return s
}

func (s *CreateIngressIngress) SetMetadata(v *CreateIngressObjectMeta) *CreateIngressIngress {
  s.Metadata = v
  return s
}

func (s *CreateIngressIngress) SetSpec(v *CreateIngressIngressSpec) *CreateIngressIngress {
  s.Spec = v
  return s
}

type CreateIngressIngressSpec struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  IngressClassName *string `json:"ingressClassName,omitempty" xml:"ingressClassName,omitempty"`
  // {"en":"DefaultBackend is the backend that should handle requests that don't match any rule", "zh_CN":"默认后端,当请求不匹配任何规则时调用"}
  DefaultBackend *CreateIngressIngressBackend `json:"defaultBackend,omitempty" xml:"defaultBackend,omitempty"`
  Tls []*CreateIngressIngressTLS `json:"tls,omitempty" xml:"tls,omitempty" type:"Repeated"`
  // {"en":"A list of host rules used to configure the CreateIngressIngress", "zh_CN":"路由规则列表"}
  Rules []*CreateIngressIngressRule `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s CreateIngressIngressSpec) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngressSpec) GoString() string {
  return s.String()
}

func (s *CreateIngressIngressSpec) SetIngressClassName(v string) *CreateIngressIngressSpec {
  s.IngressClassName = &v
  return s
}

func (s *CreateIngressIngressSpec) SetDefaultBackend(v *CreateIngressIngressBackend) *CreateIngressIngressSpec {
  s.DefaultBackend = v
  return s
}

func (s *CreateIngressIngressSpec) SetTls(v []*CreateIngressIngressTLS) *CreateIngressIngressSpec {
  s.Tls = v
  return s
}

func (s *CreateIngressIngressSpec) SetRules(v []*CreateIngressIngressRule) *CreateIngressIngressSpec {
  s.Rules = v
  return s
}

type CreateIngressIngressRule struct {
  // {"en":"Host is the fully qualified domain name of a network host", "zh_CN":"域名"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  Http *CreateIngressHTTPIngressRuleValue `json:"http,omitempty" xml:"http,omitempty"`
}

func (s CreateIngressIngressRule) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngressRule) GoString() string {
  return s.String()
}

func (s *CreateIngressIngressRule) SetHost(v string) *CreateIngressIngressRule {
  s.Host = &v
  return s
}

func (s *CreateIngressIngressRule) SetHttp(v *CreateIngressHTTPIngressRuleValue) *CreateIngressIngressRule {
  s.Http = v
  return s
}

type CreateIngressHTTPIngressRuleValue struct {
  // {"en":"A collection of paths that map requests to backends", "zh_CN":"请求路径匹配规则"}
  CreateIngressPaths []*CreateIngressHTTPIngressPath `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
}

func (s CreateIngressHTTPIngressRuleValue) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressHTTPIngressRuleValue) GoString() string {
  return s.String()
}

func (s *CreateIngressHTTPIngressRuleValue) SetCreateIngressPaths(v []*CreateIngressHTTPIngressPath) *CreateIngressHTTPIngressRuleValue {
  s.CreateIngressPaths = v
  return s
}

type CreateIngressHTTPIngressPath struct {
  // {"en":"Path is matched against the path of an incoming request", "zh_CN":"请求路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty"`
  // {"en":"PathType determines the interpretation of the Path matching,PathType can be one of the following values: Exact,Prefix,ImplementationSpecific", "zh_CN":"路径匹配类型: Exact,Prefix,ImplementationSpecific"}
  PathType *string `json:"pathType,omitempty" xml:"pathType,omitempty"`
  // {"en":"Backend defines the referenced service endpoint to which the traffic will be forwarded to", "zh_CN":"指定后端服务"}
  Backend *CreateIngressIngressBackend `json:"backend,omitempty" xml:"backend,omitempty"`
}

func (s CreateIngressHTTPIngressPath) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressHTTPIngressPath) GoString() string {
  return s.String()
}

func (s *CreateIngressHTTPIngressPath) SetPath(v string) *CreateIngressHTTPIngressPath {
  s.Path = &v
  return s
}

func (s *CreateIngressHTTPIngressPath) SetPathType(v string) *CreateIngressHTTPIngressPath {
  s.PathType = &v
  return s
}

func (s *CreateIngressHTTPIngressPath) SetBackend(v *CreateIngressIngressBackend) *CreateIngressHTTPIngressPath {
  s.Backend = v
  return s
}

type CreateIngressIngressTLS struct {
  // {"en":"Hosts are a list of hosts included in the TLS certificate", "zh_CN":"tls证书包含域名"}
  Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
  // {"en":"SecretName is the name of the secret used to terminate TLS traffic on port 443", "zh_CN":"tls秘钥名称"}
  SecretName *string `json:"secretName,omitempty" xml:"secretName,omitempty"`
}

func (s CreateIngressIngressTLS) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngressTLS) GoString() string {
  return s.String()
}

func (s *CreateIngressIngressTLS) SetHosts(v []*string) *CreateIngressIngressTLS {
  s.Hosts = v
  return s
}

func (s *CreateIngressIngressTLS) SetSecretName(v string) *CreateIngressIngressTLS {
  s.SecretName = &v
  return s
}

type CreateIngressIngressBackend struct {
  // {"en":"Service references a Service as a Backend", "zh_CN":"指定后端服务"}
  Service *CreateIngressIngressServiceBackend `json:"service,omitempty" xml:"service,omitempty"`
  // {"en":"Resource is an ObjectRef to another Kubernetes resource in the namespace of the CreateIngressIngress object", "zh_CN":"路由指定后端资源"}
  Resource *CreateIngressTypedLocalObjectReference `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s CreateIngressIngressBackend) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngressBackend) GoString() string {
  return s.String()
}

func (s *CreateIngressIngressBackend) SetService(v *CreateIngressIngressServiceBackend) *CreateIngressIngressBackend {
  s.Service = v
  return s
}

func (s *CreateIngressIngressBackend) SetResource(v *CreateIngressTypedLocalObjectReference) *CreateIngressIngressBackend {
  s.Resource = v
  return s
}

type CreateIngressTypedLocalObjectReference struct {
  // {"en":"Name is the name of resource being referenced", "zh_CN":"资源名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Kind is the type of resource being referenced", "zh_CN":"资源类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"APIGroup is the group for the resource being referenced", "zh_CN":"资源分组"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
}

func (s CreateIngressTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *CreateIngressTypedLocalObjectReference) SetName(v string) *CreateIngressTypedLocalObjectReference {
  s.Name = &v
  return s
}

func (s *CreateIngressTypedLocalObjectReference) SetKind(v string) *CreateIngressTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *CreateIngressTypedLocalObjectReference) SetApiGroup(v string) *CreateIngressTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

type CreateIngressIngressServiceBackend struct {
  // {"en":"Name is the referenced service", "zh_CN":"服务名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Port of the referenced service. A port name or port number", "zh_CN":"服务端口或端口名称"}
  Port *CreateIngressServiceBackendPort `json:"port,omitempty" xml:"port,omitempty"`
}

func (s CreateIngressIngressServiceBackend) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngressServiceBackend) GoString() string {
  return s.String()
}

func (s *CreateIngressIngressServiceBackend) SetName(v string) *CreateIngressIngressServiceBackend {
  s.Name = &v
  return s
}

func (s *CreateIngressIngressServiceBackend) SetPort(v *CreateIngressServiceBackendPort) *CreateIngressIngressServiceBackend {
  s.Port = v
  return s
}

type CreateIngressServiceBackendPort struct {
  // {"en":"Name is the name of the port on the Service", "zh_CN":"服务端口名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Number is the numerical port number on the Service", "zh_CN":"服务数字端口"}
  Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
}

func (s CreateIngressServiceBackendPort) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressServiceBackendPort) GoString() string {
  return s.String()
}

func (s *CreateIngressServiceBackendPort) SetName(v string) *CreateIngressServiceBackendPort {
  s.Name = &v
  return s
}

func (s *CreateIngressServiceBackendPort) SetNumber(v int32) *CreateIngressServiceBackendPort {
  s.Number = &v
  return s
}

type CreateIngressIngressStatus struct {
  // {"en":"LoadBalancer contains the current status of the load-balancer", "zh_CN":"包含当前负载均衡服务的状态"}
  LoadBalancer *CreateIngressLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
}

func (s CreateIngressIngressStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressIngressStatus) GoString() string {
  return s.String()
}

func (s *CreateIngressIngressStatus) SetLoadBalancer(v *CreateIngressLoadBalancerStatus) *CreateIngressIngressStatus {
  s.LoadBalancer = v
  return s
}

type CreateIngressLoadBalancerStatus struct {
  CreateIngressIngress []*CreateIngressLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s CreateIngressLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *CreateIngressLoadBalancerStatus) SetIngress(v []*CreateIngressLoadBalancerIngress) *CreateIngressLoadBalancerStatus {
  s.CreateIngressIngress = v
  return s
}

type CreateIngressLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based", "zh_CN":"负载均衡服务ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based", "zh_CN":"负载均衡类型服务dns"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Ports is a list of records of service ports", "zh_CN":"服务端口状态列表"}
  Ports []*CreateIngressPortStatus `json:"ports,omitempty" xml:"ports,omitempty" require:"true" type:"Repeated"`
}

func (s CreateIngressLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *CreateIngressLoadBalancerIngress) SetIp(v string) *CreateIngressLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *CreateIngressLoadBalancerIngress) SetHostname(v string) *CreateIngressLoadBalancerIngress {
  s.Hostname = &v
  return s
}

func (s *CreateIngressLoadBalancerIngress) SetPorts(v []*CreateIngressPortStatus) *CreateIngressLoadBalancerIngress {
  s.Ports = v
  return s
}

type CreateIngressPortStatus struct {
  // {"en":"Port is the port number of the service port of which status is recorded here", "zh_CN":"服务端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Protocol is the protocol of the service port of which status is recorded here,The supported values are: TCP, UDP, SCTP", "zh_CN":"服务支持类型: TCP,UDP,SCTP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port", "zh_CN":"记录服务端口错误信息"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s CreateIngressPortStatus) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressPortStatus) GoString() string {
  return s.String()
}

func (s *CreateIngressPortStatus) SetPort(v int32) *CreateIngressPortStatus {
  s.Port = &v
  return s
}

func (s *CreateIngressPortStatus) SetProtocol(v string) *CreateIngressPortStatus {
  s.Protocol = &v
  return s
}

func (s *CreateIngressPortStatus) SetError(v string) *CreateIngressPortStatus {
  s.Error = &v
  return s
}

type CreateIngressObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
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
  OwnerReferences []*CreateIngressOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*CreateIngressManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s CreateIngressObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressObjectMeta) GoString() string {
  return s.String()
}

func (s *CreateIngressObjectMeta) SetName(v string) *CreateIngressObjectMeta {
  s.Name = &v
  return s
}

func (s *CreateIngressObjectMeta) SetGenerateName(v string) *CreateIngressObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *CreateIngressObjectMeta) SetNamespace(v string) *CreateIngressObjectMeta {
  s.Namespace = &v
  return s
}

func (s *CreateIngressObjectMeta) SetSelfLink(v string) *CreateIngressObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *CreateIngressObjectMeta) SetUid(v string) *CreateIngressObjectMeta {
  s.Uid = &v
  return s
}

func (s *CreateIngressObjectMeta) SetResourceVersion(v string) *CreateIngressObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *CreateIngressObjectMeta) SetGeneration(v int64) *CreateIngressObjectMeta {
  s.Generation = &v
  return s
}

func (s *CreateIngressObjectMeta) SetCreationTimestamp(v string) *CreateIngressObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *CreateIngressObjectMeta) SetDeletionTimestamp(v string) *CreateIngressObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *CreateIngressObjectMeta) SetDeletionGracePeriodSeconds(v int64) *CreateIngressObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *CreateIngressObjectMeta) SetLabels(v map[string]*string) *CreateIngressObjectMeta {
  s.Labels = v
  return s
}

func (s *CreateIngressObjectMeta) SetAnnotations(v map[string]*string) *CreateIngressObjectMeta {
  s.Annotations = v
  return s
}

func (s *CreateIngressObjectMeta) SetOwnerReferences(v []*CreateIngressOwnerReference) *CreateIngressObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *CreateIngressObjectMeta) SetFinalizers(v []*string) *CreateIngressObjectMeta {
  s.Finalizers = v
  return s
}

func (s *CreateIngressObjectMeta) SetClusterName(v string) *CreateIngressObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *CreateIngressObjectMeta) SetManagedFields(v []*CreateIngressManagedFieldsEntry) *CreateIngressObjectMeta {
  s.ManagedFields = v
  return s
}

type CreateIngressManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this CreateIngressManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'CreateIngressFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“CreateIngressFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"CreateIngressFieldsV1 holds the first JSON version format as described in the 'CreateIngressFieldsV1' type", "zh_CN":"CreateIngressFieldsV1 包含类型 “CreateIngressFieldsV1” 中描述的第一个 JSON 版本格式"}
  CreateIngressFieldsV1 *CreateIngressFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s CreateIngressManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *CreateIngressManagedFieldsEntry) SetManager(v string) *CreateIngressManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *CreateIngressManagedFieldsEntry) SetOperation(v string) *CreateIngressManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *CreateIngressManagedFieldsEntry) SetApiVersion(v string) *CreateIngressManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *CreateIngressManagedFieldsEntry) SetTime(v string) *CreateIngressManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *CreateIngressManagedFieldsEntry) SetFieldsType(v string) *CreateIngressManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *CreateIngressManagedFieldsEntry) SetFieldsV1(v *CreateIngressFieldsV1) *CreateIngressManagedFieldsEntry {
  s.CreateIngressFieldsV1 = v
  return s
}

func (s *CreateIngressManagedFieldsEntry) SetSubresource(v string) *CreateIngressManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type CreateIngressFieldsV1 struct {
}

func (s CreateIngressFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressFieldsV1) GoString() string {
  return s.String()
}

type CreateIngressOwnerReference struct {
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

func (s CreateIngressOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressOwnerReference) GoString() string {
  return s.String()
}

func (s *CreateIngressOwnerReference) SetApiVersion(v string) *CreateIngressOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *CreateIngressOwnerReference) SetKind(v string) *CreateIngressOwnerReference {
  s.Kind = &v
  return s
}

func (s *CreateIngressOwnerReference) SetName(v string) *CreateIngressOwnerReference {
  s.Name = &v
  return s
}

func (s *CreateIngressOwnerReference) SetUid(v string) *CreateIngressOwnerReference {
  s.Uid = &v
  return s
}

func (s *CreateIngressOwnerReference) SetController(v bool) *CreateIngressOwnerReference {
  s.Controller = &v
  return s
}

func (s *CreateIngressOwnerReference) SetBlockOwnerDeletion(v bool) *CreateIngressOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type UpdateIngressControllerRequest struct {
  // {"en":"resource limit", "zh_CN":"资源限制"}
  Limit map[string]*string `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"resource request", "zh_CN":"所需资源"}
  Request map[string]*string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"cluster and replicate", "zh_CN":"部署集群和副本数"}
  Clusters []*UpdateIngressControllerIngressCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateIngressControllerRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerRequest) GoString() string {
  return s.String()
}

func (s *UpdateIngressControllerRequest) SetLimit(v map[string]*string) *UpdateIngressControllerRequest {
  s.Limit = v
  return s
}

func (s *UpdateIngressControllerRequest) SetRequest(v map[string]*string) *UpdateIngressControllerRequest {
  s.Request = v
  return s
}

func (s *UpdateIngressControllerRequest) SetClusters(v []*UpdateIngressControllerIngressCluster) *UpdateIngressControllerRequest {
  s.Clusters = v
  return s
}

type UpdateIngressControllerResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress controller", "zh_CN":"路由控制器"}
  Data *UpdateIngressControllerIngressController `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateIngressControllerResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerResponse) GoString() string {
  return s.String()
}

func (s *UpdateIngressControllerResponse) SetCode(v int64) *UpdateIngressControllerResponse {
  s.Code = &v
  return s
}

func (s *UpdateIngressControllerResponse) SetMsg(v string) *UpdateIngressControllerResponse {
  s.Msg = &v
  return s
}

func (s *UpdateIngressControllerResponse) SetRequestId(v string) *UpdateIngressControllerResponse {
  s.RequestId = &v
  return s
}

func (s *UpdateIngressControllerResponse) SetData(v *UpdateIngressControllerIngressController) *UpdateIngressControllerResponse {
  s.Data = v
  return s
}

type UpdateIngressControllerPaths struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateIngressControllerPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerPaths) GoString() string {
  return s.String()
}

func (s *UpdateIngressControllerPaths) SetName(v string) *UpdateIngressControllerPaths {
  s.Name = &v
  return s
}

type UpdateIngressControllerParameters struct {
}

func (s UpdateIngressControllerParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerParameters) GoString() string {
  return s.String()
}

type UpdateIngressControllerRequestHeader struct {
}

func (s UpdateIngressControllerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerRequestHeader) GoString() string {
  return s.String()
}

type UpdateIngressControllerResponseHeader struct {
}

func (s UpdateIngressControllerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerResponseHeader) GoString() string {
  return s.String()
}

type UpdateIngressControllerIngressController struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"resource limit", "zh_CN":"资源限制"}
  Limit map[string]*string `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"resource request", "zh_CN":"所需资源"}
  Request map[string]*string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"cluster and replicate", "zh_CN":"部署集群和副本数"}
  Clusters []*UpdateIngressControllerIngressCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateIngressControllerIngressController) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerIngressController) GoString() string {
  return s.String()
}

func (s *UpdateIngressControllerIngressController) SetName(v string) *UpdateIngressControllerIngressController {
  s.Name = &v
  return s
}

func (s *UpdateIngressControllerIngressController) SetLimit(v map[string]*string) *UpdateIngressControllerIngressController {
  s.Limit = v
  return s
}

func (s *UpdateIngressControllerIngressController) SetRequest(v map[string]*string) *UpdateIngressControllerIngressController {
  s.Request = v
  return s
}

func (s *UpdateIngressControllerIngressController) SetClusters(v []*UpdateIngressControllerIngressCluster) *UpdateIngressControllerIngressController {
  s.Clusters = v
  return s
}

type UpdateIngressControllerIngressCluster struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"replicates", "zh_CN":"副本数"}
  Replicate *int32 `json:"replicate,omitempty" xml:"replicate,omitempty" require:"true"`
}

func (s UpdateIngressControllerIngressCluster) String() string {
  return tea.Prettify(s)
}

func (s UpdateIngressControllerIngressCluster) GoString() string {
  return s.String()
}

func (s *UpdateIngressControllerIngressCluster) SetName(v string) *UpdateIngressControllerIngressCluster {
  s.Name = &v
  return s
}

func (s *UpdateIngressControllerIngressCluster) SetReplicate(v int32) *UpdateIngressControllerIngressCluster {
  s.Replicate = &v
  return s
}




type LECHQueryAvailableCidrsDetailRequest struct {
}

func (s LECHQueryAvailableCidrsDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailRequest) GoString() string {
  return s.String()
}

type LECHQueryAvailableCidrsDetailRequestHeader struct {
}

func (s LECHQueryAvailableCidrsDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryAvailableCidrsDetailPaths struct {
}

func (s LECHQueryAvailableCidrsDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailPaths) GoString() string {
  return s.String()
}

type LECHQueryAvailableCidrsDetailParameters struct {
  // {"en":"Node name.","zh_CN":"节点名称，多个节点用英文逗号分隔，最多填写20个"}
  Node *string `json:"node,omitempty" xml:"node,omitempty" require:"true"`
  // {"en":"CIDR","zh_CN":"网段 CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"IP native attribute.(-1: Native, 1: Non-native)","zh_CN":"原生属性. (-1: 原生, 1: 非原生)"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"IPv6 segment(1: yes, -1: no)","zh_CN":"是否返回IPV6网段(1: 是, -1: 否)"}
  NeedIPv6 *string `json:"needIPv6,omitempty" xml:"needIPv6,omitempty"`
}

func (s LECHQueryAvailableCidrsDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsDetailParameters) SetNode(v string) *LECHQueryAvailableCidrsDetailParameters {
  s.Node = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailParameters) SetCidr(v string) *LECHQueryAvailableCidrsDetailParameters {
  s.Cidr = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailParameters) SetNativeAttribute(v string) *LECHQueryAvailableCidrsDetailParameters {
  s.NativeAttribute = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailParameters) SetNeedIPv6(v string) *LECHQueryAvailableCidrsDetailParameters {
  s.NeedIPv6 = &v
  return s
}

type LECHQueryAvailableCidrsDetailResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHQueryAvailableCidrsDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHQueryAvailableCidrsDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsDetailResponse) SetCode(v string) *LECHQueryAvailableCidrsDetailResponse {
  s.Code = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponse) SetData(v *LECHQueryAvailableCidrsDetailResponseData) *LECHQueryAvailableCidrsDetailResponse {
  s.Data = v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponse) SetMessage(v string) *LECHQueryAvailableCidrsDetailResponse {
  s.Message = &v
  return s
}

type LECHQueryAvailableCidrsDetailResponseData struct {
  // {"en":"available cidrs","zh_CN":"可用的cidr列表"}
  Nodes []*LECHQueryAvailableCidrsDetailResponseDataNodes `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryAvailableCidrsDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailResponseData) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsDetailResponseData) SetNodes(v []*LECHQueryAvailableCidrsDetailResponseDataNodes) *LECHQueryAvailableCidrsDetailResponseData {
  s.Nodes = v
  return s
}

type LECHQueryAvailableCidrsDetailResponseDataNodes struct     {
  // {"en":"CIDR detail.","zh_CN":"网段详情"}
  Cidrs []*LECHQueryAvailableCidrsDetailResponseDataNodesCidrs `json:"cidrs,omitempty" xml:"cidrs,omitempty" require:"true" type:"Repeated"`
  // {"en":"Node name.","zh_CN":"节点名称"}
  Node *string `json:"node,omitempty" xml:"node,omitempty" require:"true"`
}

func (s LECHQueryAvailableCidrsDetailResponseDataNodes) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailResponseDataNodes) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodes) SetCidrs(v []*LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) *LECHQueryAvailableCidrsDetailResponseDataNodes {
  s.Cidrs = v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodes) SetNode(v string) *LECHQueryAvailableCidrsDetailResponseDataNodes {
  s.Node = &v
  return s
}

type LECHQueryAvailableCidrsDetailResponseDataNodesCidrs struct     {
  // {"en":"CIDR","zh_CN":"网段 CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en":"Number of free and available IPs","zh_CN":"空闲可用IP数"}
  FreeIps *int `json:"freeIps,omitempty" xml:"freeIps,omitempty" require:"true"`
  // {"en":"Freezing IP number","zh_CN":"冷却IP数"}
  FreezingIps *int `json:"freezingIps,omitempty" xml:"freezingIps,omitempty" require:"true"`
  // {"en":"IP native attribute.(-1: Native, 1: Non-native)","zh_CN":"原生属性. (-1: 原生, 1: 非原生)"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty" require:"true"`
  // {"en":"Number of IPs already used","zh_CN":"已用IP数"}
  UsedIps *int `json:"usedIps,omitempty" xml:"usedIps,omitempty" require:"true"`
}

func (s LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) GoString() string {
  return s.String()
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) SetCidr(v string) *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.Cidr = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) SetFreeIps(v int) *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.FreeIps = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) SetFreezingIps(v int) *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.FreezingIps = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) SetNativeAttribute(v string) *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.NativeAttribute = &v
  return s
}

func (s *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs) SetUsedIps(v int) *LECHQueryAvailableCidrsDetailResponseDataNodesCidrs {
  s.UsedIps = &v
  return s
}

type LECHQueryAvailableCidrsDetailResponseHeader struct {
}

func (s LECHQueryAvailableCidrsDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryAvailableCidrsDetailResponseHeader) GoString() string {
  return s.String()
}




type GetServiceRequest struct {
}

func (s GetServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
  return s.String()
}

type GetServiceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"service", "zh_CN":"service"}
  Data *GetServiceService `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
  return s.String()
}

func (s *GetServiceResponse) SetCode(v int64) *GetServiceResponse {
  s.Code = &v
  return s
}

func (s *GetServiceResponse) SetMsg(v string) *GetServiceResponse {
  s.Msg = &v
  return s
}

func (s *GetServiceResponse) SetRequestId(v string) *GetServiceResponse {
  s.RequestId = &v
  return s
}

func (s *GetServiceResponse) SetData(v *GetServiceService) *GetServiceResponse {
  s.Data = v
  return s
}

type GetServicePaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"service name", "zh_CN":"service 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetServicePaths) String() string {
  return tea.Prettify(s)
}

func (s GetServicePaths) GoString() string {
  return s.String()
}

func (s *GetServicePaths) SetNamespace(v string) *GetServicePaths {
  s.Namespace = &v
  return s
}

func (s *GetServicePaths) SetName(v string) *GetServicePaths {
  s.Name = &v
  return s
}

type GetServiceParameters struct {
}

func (s GetServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s GetServiceParameters) GoString() string {
  return s.String()
}

type GetServiceRequestHeader struct {
}

func (s GetServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetServiceRequestHeader) GoString() string {
  return s.String()
}

type GetServiceResponseHeader struct {
}

func (s GetServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetServiceResponseHeader) GoString() string {
  return s.String()
}

type GetServiceService struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *GetServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 GetServiceService 的行为"}
  Spec *GetServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the service. Populated by the system. Read-only", "zh_CN":"最近观察到的 GetServiceService 状态。由系统填充。只读"}
  Status *GetServiceServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceService) String() string {
  return tea.Prettify(s)
}

func (s GetServiceService) GoString() string {
  return s.String()
}

func (s *GetServiceService) SetApiVersion(v string) *GetServiceService {
  s.ApiVersion = &v
  return s
}

func (s *GetServiceService) SetKind(v string) *GetServiceService {
  s.Kind = &v
  return s
}

func (s *GetServiceService) SetMetadata(v *GetServiceObjectMeta) *GetServiceService {
  s.Metadata = v
  return s
}

func (s *GetServiceService) SetSpec(v *GetServiceServiceSpec) *GetServiceService {
  s.Spec = v
  return s
}

func (s *GetServiceService) SetStatus(v *GetServiceServiceStatus) *GetServiceService {
  s.Status = v
  return s
}

type GetServiceServiceStatus struct {
  // {"en":"Current service state", "zh_CN":"loadBalancer 包含负载均衡器的当前状态（如果存在）"}
  LoadBalancer *GetServiceLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
  // {"en":"LoadBalancer contains the current status of the load-balancer, if one is present", "zh_CN":"服务的当前状态"}
  Conditions []*GetServiceMetaV1Condition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s GetServiceServiceStatus) String() string {
  return tea.Prettify(s)
}

func (s GetServiceServiceStatus) GoString() string {
  return s.String()
}

func (s *GetServiceServiceStatus) SetLoadBalancer(v *GetServiceLoadBalancerStatus) *GetServiceServiceStatus {
  s.LoadBalancer = v
  return s
}

func (s *GetServiceServiceStatus) SetConditions(v []*GetServiceMetaV1Condition) *GetServiceServiceStatus {
  s.Conditions = v
  return s
}

type GetServiceMetaV1Condition struct {
  // {"en":"type of condition in CamelCase or in foo.example.com/CamelCase", "zh_CN":"CamelCase 或 foo.example.com/CamelCase 中的条件类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status of the condition, one of True, False, Unknown", "zh_CN":"condition 的状态，True、False、Unknown 之一"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance", "zh_CN":"表示设置 condition 基于的 .metadata.generation 的过期次数。 例如，如果 .metadata.generation 当前为 12，但 .status.conditions[x].observedGeneration 为 9， 则 condition 相对于实例的当前状态已过期"}
  ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
  // {"en":"lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable", "zh_CN":"状况最近一次状态转化的时间。 变化应该发生在下层状况发生变化的时候。如果不知道下层状况发生变化的时间， 那么使用 API 字段更改的时间是可以接受的"}
  LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
  // {"en":"reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty", "zh_CN":"reason 包含一个程序标识符，指示 condition 最后一次转换的原因。 特定条件类型的生产者可以定义该字段的预期值和含义，以及这些值是否被视为有保证的 API。 该值应该是 CamelCase 字符串且不能为空"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":"message is a human readable message indicating details about the transition. This may be an empty string", "zh_CN":"message 是人类可读的消息，有关转换的详细信息，可以是空字符串"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetServiceMetaV1Condition) String() string {
  return tea.Prettify(s)
}

func (s GetServiceMetaV1Condition) GoString() string {
  return s.String()
}

func (s *GetServiceMetaV1Condition) SetType(v string) *GetServiceMetaV1Condition {
  s.Type = &v
  return s
}

func (s *GetServiceMetaV1Condition) SetStatus(v string) *GetServiceMetaV1Condition {
  s.Status = &v
  return s
}

func (s *GetServiceMetaV1Condition) SetObservedGeneration(v int64) *GetServiceMetaV1Condition {
  s.ObservedGeneration = &v
  return s
}

func (s *GetServiceMetaV1Condition) SetLastTransitionTime(v string) *GetServiceMetaV1Condition {
  s.LastTransitionTime = &v
  return s
}

func (s *GetServiceMetaV1Condition) SetReason(v string) *GetServiceMetaV1Condition {
  s.Reason = &v
  return s
}

func (s *GetServiceMetaV1Condition) SetMessage(v string) *GetServiceMetaV1Condition {
  s.Message = &v
  return s
}

type GetServiceLoadBalancerStatus struct {
  // {"en":"Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points", "zh_CN":"ingress 是一个包含负载均衡器 Ingress 点的列表。GetServiceService 的流量需要被发送到这些 Ingress 点"}
  Ingress []*GetServiceLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s GetServiceLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s GetServiceLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *GetServiceLoadBalancerStatus) SetIngress(v []*GetServiceLoadBalancerIngress) *GetServiceLoadBalancerStatus {
  s.Ingress = v
  return s
}

type GetServiceLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)", "zh_CN":"ip 是为基于 IP 的负载均衡器 Ingress 点（通常是 GCE 或 OpenStack 负载均衡器）设置的"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)", "zh_CN":"hostname 是为基于 DNS 的负载均衡器 Ingress 点（通常是 AWS 负载均衡器）设置的"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
  // {"en":"Ports is a list of records of service ports If used, every port defined in the service should have an entry in it", "zh_CN":"ports 是 GetServiceService 的端口列表。如果设置了此字段，GetServiceService 中定义的每个端口都应该在此列表中"}
  Ports []*GetServicePortStatus `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s GetServiceLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s GetServiceLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *GetServiceLoadBalancerIngress) SetIp(v string) *GetServiceLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *GetServiceLoadBalancerIngress) SetHostName(v string) *GetServiceLoadBalancerIngress {
  s.HostName = &v
  return s
}

func (s *GetServiceLoadBalancerIngress) SetPorts(v []*GetServicePortStatus) *GetServiceLoadBalancerIngress {
  s.Ports = v
  return s
}

type GetServicePortStatus struct {
  // {"en":"the port number of the service port of which status is recorded here", "zh_CN":"port 是所记录的服务端口状态的端口号"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"the protocol of the service port of which status is recorded here The supported values are: 'TCP', 'UDP', 'SCTP'", "zh_CN":"protocol 是所记录的服务端口状态的协议。支持的值为：“TCP”、”UDP”、“SCTP”"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase", "zh_CN":"error 是记录 GetServiceService 端口的问题。 错误的格式应符合以下规则:内置错误原因应在此文件中指定，应使用 CamelCase 名称。云提供商特定错误原因的名称必须符合格式 foo.example.com/CamelCase"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s GetServicePortStatus) String() string {
  return tea.Prettify(s)
}

func (s GetServicePortStatus) GoString() string {
  return s.String()
}

func (s *GetServicePortStatus) SetPort(v int32) *GetServicePortStatus {
  s.Port = &v
  return s
}

func (s *GetServicePortStatus) SetProtocol(v string) *GetServicePortStatus {
  s.Protocol = &v
  return s
}

func (s *GetServicePortStatus) SetError(v string) *GetServicePortStatus {
  s.Error = &v
  return s
}

type GetServiceServiceSpec struct {
  // {"en":"The list of ports that are exposed by this service", "zh_CN":"此 GetServiceService 公开的端口列表"}
  Ports []*GetServiceServicePort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en":"Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName", "zh_CN":"将 GetServiceService 流量路由到具有与此 selector 匹配的标签键值对的 Pod。 如果为空或不存在，则假定该服务有一个外部进程管理其端点，Kubernetes 不会修改该端点。 仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型。如果类型为 ExternalName，则忽略"}
  Selector map[string]*string `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a GetServiceService of type ExternalName, creation will fail. This field will be wiped when updating a GetServiceService to type ExternalName", "zh_CN":"clusterIP 是服务的 IP 地址，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给服务，否则创建服务将失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIP 为空）或 type 已经是 ExternalName 时，可以更改 clusterIP（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIP 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 仅适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 GetServiceService 时指定了 clusterIP，则创建将失败。 更新 GetServiceService type 为 ExternalName 时，clusterIP 会被移除"}
  ClusterIP *string `json:"clusterIP,omitempty" xml:"clusterIP,omitempty"`
  // {"en":"ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a GetServiceService of type ExternalName, creation will fail. This field will be wiped when updating a GetServiceService to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"clusterIPs 是分配给该 GetServiceService 的 IP 地址列表，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给 GetServiceService；否则创建 GetServiceService 失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIPs 为空）或 type 已经是 ExternalName 时，可以更改 clusterIPs（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIPs 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 GetServiceService 时指定了 clusterIPs，则会创建失败。 更新 GetServiceService type 为 ExternalName 时，该字段将被移除。如果未指定此字段，则将从 clusterIP 字段初始化。 如果指定 clusterIPs，客户端必须确保 clusterIPs[0] 和 clusterIP 一致。clusterIPs 最多可包含两个条目（双栈系列，按任意顺序）。 这些 IP 必须与 ipFamilies 的值相对应。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 管理"}
  ClusterIPs []*string `json:"clusterIPs,omitempty" xml:"clusterIPs,omitempty" type:"Repeated"`
  // {"en":"type determines how the GetServiceService is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services", "zh_CN":"type 确定 GetServiceService 的公开方式。默认为 ClusterIP。 有效选项为 ExternalName、ClusterIP、NodePort 和 LoadBalancer。 “ClusterIP” 为端点分配一个集群内部 IP 地址用于负载均衡。 Endpoints 由 selector 确定，如果未设置 selector，则需要通过手动构造 Endpoints 或 EndpointSlice 的对象来确定。 如果 clusterIP 为 “None”，则不分配虚拟 IP，并且 Endpoints 作为一组端点而不是虚拟 IP 发布。 “NodePort” 建立在 ClusterIP 之上，并在每个节点上分配一个端口，该端口路由到与 clusterIP 相同的 Endpoints。 “LoadBalancer” 基于 NodePort 构建并创建一个外部负载均衡器（如果当前云支持），该负载均衡器路由到与 clusterIP 相同的 Endpoints。 “externalName” 将此 GetServiceService 别名为指定的 externalName。其他几个字段不适用于 ExternalName GetServiceService"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system", "zh_CN":"externalIPs 是一个 IP 列表，集群中的节点会为此 GetServiceService 接收针对这些 IP 地址的流量。 这些 IP 不被 Kubernetes 管理。用户需要确保流量可以到达具有此 IP 的节点。 一个常见的例子是不属于 Kubernetes 系统的外部负载均衡器"}
  ExternalIPs []*string `json:"externalIPs,omitempty" xml:"externalIPs,omitempty" type:"Repeated"`
  // {"en":"Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None", "zh_CN":"支持 “ClientIP” 和 “None”。用于维护会话亲和性。 启用基于客户端 IP 的会话亲和性。必须是 ClientIP 或 None。默认为 None"}
  SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
  // {"en":"Only applies to GetServiceService Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version", "zh_CN":"仅适用于服务类型: LoadBalancer。此功能取决于底层云提供商是否支持负载均衡器。 如果云提供商不支持该功能，该字段将被忽略。 已弃用: 该字段信息不足，且其含义因实现而异，而且不支持双栈。 从 Kubernetes v1.24 开始，鼓励用户在可用时使用特定于实现的注释。在未来的 API 版本中可能会删除此字段"}
  LoadBalancerIP *string `json:"loadBalancerIP,omitempty" xml:"loadBalancerIP,omitempty"`
  // {"en":"If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature", "zh_CN":"如果设置了此字段并且被平台支持，将限制通过云厂商的负载均衡器的流量到指定的客户端 IP。 如果云提供商不支持该功能，该字段将被忽略"}
  LoadBalancerSourceRanges []*string `json:"loadBalancerSourceRanges,omitempty" xml:"loadBalancerSourceRanges,omitempty" type:"Repeated"`
  // {"en":"externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires type to be 'ExternalName'", "zh_CN":"externalName 是发现机制将返回的外部引用，作为此服务的别名（例如 DNS CNAME 记录）。 不涉及代理。必须是小写的 RFC-1123 主机名 (https://tools.ietf.org/html/rfc1123)， 并且要求 type 为 “ExternalName”"}
  ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
  // {"en":"externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the GetServiceService's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node", "zh_CN":"externalTrafficPolicy 描述了节点如何分发它们在 GetServiceService 的“外部访问”地址 （NodePort、ExternalIP 和 LoadBalancer IP）接收到的服务流量。 如果设置为 “Local”，代理将以一种假设外部负载均衡器将负责在节点之间服务流量负载均衡， 因此每个节点将仅向服务的节点本地端点传递流量，而不会伪装客户端源 IP。 （将丢弃错误发送到没有端点的节点的流量。） “Cluster” 默认值使用负载均衡路由到所有端点的策略（可能会根据拓扑和其他特性进行修改）。 请注意，从集群内部发送到 External IP 或 LoadBalancer IP 的流量始终具有 “Cluster” 语义， 但是从集群内部发送到 NodePort 的客户端需要在选择节点时考虑流量路由策略"}
  ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" xml:"externalTrafficPolicy,omitempty"`
  // {"en":"healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a GetServiceService which does not need it, creation will fail. This field will be wiped when updating a GetServiceService to no longer need it (e.g. changing type). This field cannot be updated once set", "zh_CN":"healthCheckNodePort 指定 GetServiceService 的健康检查节点端口。 仅适用于 type 为 LoadBalancer 且 externalTrafficPolicy 设置为 Local 的情况。 如果为此字段设定了一个值，该值在合法范围内且没有被使用，则使用所指定的值。 如果未设置此字段，则自动分配字段值。外部系统（例如负载平衡器）可以使用此端口来确定给定节点是否拥有此服务的端点。 在创建不需要 healthCheckNodePort 的 GetServiceService 时指定了此字段，则 GetServiceService 创建会失败。 要移除 healthCheckNodePort，需要更改 GetServiceService 的 type。 该字段一旦设置就无法更改"}
  HealthCheckNodePort *int32 `json:"healthCheckNodePort,omitempty" xml:"healthCheckNodePort,omitempty"`
  // {"en":"publishNotReadyAddresses indicates that any agent which deals with endpoints for this GetServiceService should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless GetServiceService to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior", "zh_CN":"publishNotReadyAddresses 表示任何处理此 GetServiceService 端点的代理都应忽略任何准备就绪/未准备就绪的指示。 设置此字段的主要场景是为 StatefulSet 的服务提供支持，使之能够为其 Pod 传播 SRV DNS 记录，以实现对等发现。 为 GetServiceService 生成 Endpoints 和 EndpointSlice 资源的 Kubernetes 控制器对字段的解读是， 即使 Pod 本身还没有准备好，所有端点都可被视为 “已就绪”。 对于代理而言，如果仅使用 Kubernetes 通过 Endpoints 或 EndpointSlice 资源所生成的端点， 则可以安全地假设这种行为"}
  PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" xml:"publishNotReadyAddresses,omitempty"`
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"sessionAffinityConfig 包含会话亲和性的配置"}
  GetServiceSessionAffinityConfig *GetServiceSessionAffinityConfig `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
  // {"en":"IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the GetServiceService. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a GetServiceService to type ExternalName.This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"iPFamilies 是分配给此服务的 IP 协议（例如 IPv4、IPv6）的列表。 该字段通常根据集群配置和 ipFamilyPolicy 字段自动设置。 如果手动指定该字段，且请求的协议在集群中可用，且 ipFamilyPolicy 允许，则使用；否则服务创建将失败。 该字段修改是有条件的：它允许添加或删除辅助 IP 协议，但不允许更改服务的主要 IP 协议。 有效值为 “IPv4” 和 “IPv6”。 该字段仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型的服务，并且确实可用于“无头”服务。 更新服务设置类型为 ExternalName 时，该字段将被擦除。该字段最多可以包含两个条目（双栈系列，按任意顺序）。 如果指定，这些协议栈必须对应于 clusterIPs 字段的值。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 字段管理"}
  IpFamilies []*string `json:"ipFamilies,omitempty" xml:"ipFamilies,omitempty" type:"Repeated"`
  // {"en":"IPFamilyPolicy represents the dual-stack-ness requested or required by this GetServiceService. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName", "zh_CN":"iPFamilyPolicy 表示此服务请求或要求的双栈特性。 如果没有提供值，则此字段将被设置为 SingleStack。 服务可以是 “SingleStack”（单个 IP 协议）、 “PreferDualStack”（双栈配置集群上的两个 IP 协议或单栈集群上的单个 IP 协议） 或 “RequireDualStack”（双栈上的两个 IP 协议配置的集群，否则失败）。 ipFamilies 和 clusterIPs 字段取决于此字段的值。 更新服务设置类型为 ExternalName 时，此字段将被擦除"}
  IpFamilyPolicy *string `json:"ipFamilyPolicy,omitempty" xml:"ipFamilyPolicy,omitempty"`
  // {"en":"allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type", "zh_CN":"allocateLoadBalancerNodePorts 定义了是否会自动为 LoadBalancer 类型的 GetServiceService 分配 NodePort。默认为 true。 如果集群负载均衡器不依赖 NodePort，则可以设置此字段为 false。 如果调用者（通过指定一个值）请求特定的 NodePort，则无论此字段如何，都会接受这些请求。 该字段只能设置在 type 为 LoadBalancer 的 GetServiceService 上，如果 type 更改为任何其他类型，该字段将被移除"}
  AllocateLoadBalancerNodePorts *bool `json:"allocateLoadBalancerNodePorts,omitempty" xml:"allocateLoadBalancerNodePorts,omitempty"`
  // {"en":"loadBalancerClass is the class of the load balancer implementation this GetServiceService belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the GetServiceService type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a GetServiceService to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type", "zh_CN":"loadBalancerClass 是此 GetServiceService 所属的负载均衡器实现的类。 如果设置了此字段，则字段值必须是标签风格的标识符，带有可选前缀，例如 ”internal-vip” 或 “example.com/internal-vip”。 无前缀名称是为最终用户保留的。该字段只能在 GetServiceService 类型为 “LoadBalancer” 时设置。 如果未设置此字段，则使用默认负载均衡器实现。默认负载均衡器现在通常通过云提供商集成完成，但应适用于任何默认实现。 如果设置了此字段，则假定负载均衡器实现正在监测具有对应负载均衡器类的 GetServiceService。 任何默认负载均衡器实现（例如云提供商）都应忽略设置此字段的 GetServiceService。 只有在创建或更新的 GetServiceService 的 type 为 “LoadBalancer” 时，才可设置此字段。 一经设定，不可更改。当 GetServiceService 的 type 更新为 “LoadBalancer” 之外的其他类型时，此字段将被移除"}
  LoadBalancerClass *string `json:"loadBalancerClass,omitempty" xml:"loadBalancerClass,omitempty"`
  // {"en":"InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features)", "zh_CN":"InternalTrafficPolicy 描述节点如何分发它们在 ClusterIP 上接收到的服务流量。 如果设置为 “Local”，代理将假定 Pod 只想与在同一节点上的服务端点通信，如果没有本地端点，它将丢弃流量。 “Cluster” 默认将流量路由到所有端点（可能会根据拓扑和其他特性进行修改）"}
  InternalTrafficPolicy *string `json:"internalTrafficPolicy,omitempty" xml:"internalTrafficPolicy,omitempty"`
}

func (s GetServiceServiceSpec) String() string {
  return tea.Prettify(s)
}

func (s GetServiceServiceSpec) GoString() string {
  return s.String()
}

func (s *GetServiceServiceSpec) SetPorts(v []*GetServiceServicePort) *GetServiceServiceSpec {
  s.Ports = v
  return s
}

func (s *GetServiceServiceSpec) SetSelector(v map[string]*string) *GetServiceServiceSpec {
  s.Selector = v
  return s
}

func (s *GetServiceServiceSpec) SetClusterIP(v string) *GetServiceServiceSpec {
  s.ClusterIP = &v
  return s
}

func (s *GetServiceServiceSpec) SetClusterIPs(v []*string) *GetServiceServiceSpec {
  s.ClusterIPs = v
  return s
}

func (s *GetServiceServiceSpec) SetType(v string) *GetServiceServiceSpec {
  s.Type = &v
  return s
}

func (s *GetServiceServiceSpec) SetExternalIPs(v []*string) *GetServiceServiceSpec {
  s.ExternalIPs = v
  return s
}

func (s *GetServiceServiceSpec) SetSessionAffinity(v string) *GetServiceServiceSpec {
  s.SessionAffinity = &v
  return s
}

func (s *GetServiceServiceSpec) SetLoadBalancerIP(v string) *GetServiceServiceSpec {
  s.LoadBalancerIP = &v
  return s
}

func (s *GetServiceServiceSpec) SetLoadBalancerSourceRanges(v []*string) *GetServiceServiceSpec {
  s.LoadBalancerSourceRanges = v
  return s
}

func (s *GetServiceServiceSpec) SetExternalName(v string) *GetServiceServiceSpec {
  s.ExternalName = &v
  return s
}

func (s *GetServiceServiceSpec) SetExternalTrafficPolicy(v string) *GetServiceServiceSpec {
  s.ExternalTrafficPolicy = &v
  return s
}

func (s *GetServiceServiceSpec) SetHealthCheckNodePort(v int32) *GetServiceServiceSpec {
  s.HealthCheckNodePort = &v
  return s
}

func (s *GetServiceServiceSpec) SetPublishNotReadyAddresses(v bool) *GetServiceServiceSpec {
  s.PublishNotReadyAddresses = &v
  return s
}

func (s *GetServiceServiceSpec) SetSessionAffinityConfig(v *GetServiceSessionAffinityConfig) *GetServiceServiceSpec {
  s.GetServiceSessionAffinityConfig = v
  return s
}

func (s *GetServiceServiceSpec) SetIpFamilies(v []*string) *GetServiceServiceSpec {
  s.IpFamilies = v
  return s
}

func (s *GetServiceServiceSpec) SetIpFamilyPolicy(v string) *GetServiceServiceSpec {
  s.IpFamilyPolicy = &v
  return s
}

func (s *GetServiceServiceSpec) SetAllocateLoadBalancerNodePorts(v bool) *GetServiceServiceSpec {
  s.AllocateLoadBalancerNodePorts = &v
  return s
}

func (s *GetServiceServiceSpec) SetLoadBalancerClass(v string) *GetServiceServiceSpec {
  s.LoadBalancerClass = &v
  return s
}

func (s *GetServiceServiceSpec) SetInternalTrafficPolicy(v string) *GetServiceServiceSpec {
  s.InternalTrafficPolicy = &v
  return s
}

type GetServiceServicePort struct {
  // {"en":"The name of this port within the service. This must be a DNS_LABEL. All ports within a GetServiceServiceSpec must have unique names. When considering the endpoints for a GetServiceService, this must match the 'name' field in the EndpointPort. Optional if only one GetServiceServicePort is defined on this service", "zh_CN":"GetServiceService 中此端口的名称。这必须是 DNS_LABEL。 GetServiceServiceSpec 中的所有端口的名称都必须唯一。 在考虑 GetServiceService 的端点时，这一字段值必须与 EndpointPort 中的 name 字段相同。 如果此服务上仅定义一个 GetServiceServicePort，则为此字段为可选"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The IP protocol for this port. Supports TCP, UDP, and SCTP. Default is TCP", "zh_CN":"此端口的 IP 协议。支持 “TCP”、“UDP” 和 “SCTP”。默认为 TCP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol", "zh_CN":"此端口的应用协议，遵循标准的 Kubernetes 标签语法，无前缀名称按照 IANA 标准服务名称 （参见 RFC-6335 和 https://www.iana.org/assignments/service-names）。 非标准协议应该使用前缀名称，如 mycompany.com/my-custom-protocol"}
  AppProtocol *string `json:"appProtocol,omitempty" xml:"appProtocol,omitempty"`
  // {"en":"The port that will be exposed by this service", "zh_CN":"GetServiceService 将公开的端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field", "zh_CN":"在 GetServiceService 所针对的 Pod 上要访问的端口号或名称。 编号必须在 1 到 65535 的范围内。名称必须是 IANA_SVC_NAME。 如果此值是一个字符串，将在目标 Pod 的容器端口中作为命名端口进行查找。 如果未指定字段，则使用 “port” 字段的值（直接映射）。 对于 clusterIP 为 None 的服务，此字段将被忽略， 应忽略不设或设置为 “port” 字段的取值"}
  TargetPort *int32 `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
  // {"en":"The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this GetServiceService requires one. If this field is specified when creating a GetServiceService which does not need it, creation will fail. This field will be wiped when updating a GetServiceService to no longer need it (e.g. changing type from NodePort to ClusterIP).", "zh_CN":"当类型为 NodePort 或 LoadBalancer 时，GetServiceService 公开在节点上的端口， 通常由系统分配。如果指定了一个在范围内且未使用的值，则将使用该值，否则操作将失败。 如果在创建的 GetServiceService 需要该端口时未指定该字段，则会分配端口。 如果在创建不需要该端口的 Service时指定了该字段，则会创建失败。 当更新 GetServiceService 时，如果不再需要此字段（例如，将类型从 NodePort 更改为 ClusterIP），这个字段将被擦除"}
  NodePort *int32 `json:"nodePort,omitempty" xml:"nodePort,omitempty"`
}

func (s GetServiceServicePort) String() string {
  return tea.Prettify(s)
}

func (s GetServiceServicePort) GoString() string {
  return s.String()
}

func (s *GetServiceServicePort) SetName(v string) *GetServiceServicePort {
  s.Name = &v
  return s
}

func (s *GetServiceServicePort) SetProtocol(v string) *GetServiceServicePort {
  s.Protocol = &v
  return s
}

func (s *GetServiceServicePort) SetAppProtocol(v string) *GetServiceServicePort {
  s.AppProtocol = &v
  return s
}

func (s *GetServiceServicePort) SetPort(v int32) *GetServiceServicePort {
  s.Port = &v
  return s
}

func (s *GetServiceServicePort) SetTargetPort(v int32) *GetServiceServicePort {
  s.TargetPort = &v
  return s
}

func (s *GetServiceServicePort) SetNodePort(v int32) *GetServiceServicePort {
  s.NodePort = &v
  return s
}

type GetServiceSessionAffinityConfig struct {
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"clientIP 包含基于客户端 IP 的会话亲和性的配置"}
  ClientIP *GetServiceClientIPConfig `json:"clientIP,omitempty" xml:"clientIP,omitempty"`
}

func (s GetServiceSessionAffinityConfig) String() string {
  return tea.Prettify(s)
}

func (s GetServiceSessionAffinityConfig) GoString() string {
  return s.String()
}

func (s *GetServiceSessionAffinityConfig) SetClientIP(v *GetServiceClientIPConfig) *GetServiceSessionAffinityConfig {
  s.ClientIP = v
  return s
}

type GetServiceClientIPConfig struct {
  // {"en":"timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).", "zh_CN":"timeoutSeconds 指定 ClientIP 类型会话的维系时间秒数。 如果 ServiceAffinity == 'ClientIP'，则该值必须 >0 && <=86400（1 天）。默认值为 10800（3 小时）"}
  TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s GetServiceClientIPConfig) String() string {
  return tea.Prettify(s)
}

func (s GetServiceClientIPConfig) GoString() string {
  return s.String()
}

func (s *GetServiceClientIPConfig) SetTimeoutSeconds(v int32) *GetServiceClientIPConfig {
  s.TimeoutSeconds = &v
  return s
}

type GetServiceObjectMeta struct {
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
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 GetServiceService 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*GetServiceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*GetServiceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s GetServiceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s GetServiceObjectMeta) GoString() string {
  return s.String()
}

func (s *GetServiceObjectMeta) SetName(v string) *GetServiceObjectMeta {
  s.Name = &v
  return s
}

func (s *GetServiceObjectMeta) SetGenerateName(v string) *GetServiceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *GetServiceObjectMeta) SetNamespace(v string) *GetServiceObjectMeta {
  s.Namespace = &v
  return s
}

func (s *GetServiceObjectMeta) SetSelfLink(v string) *GetServiceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *GetServiceObjectMeta) SetUid(v string) *GetServiceObjectMeta {
  s.Uid = &v
  return s
}

func (s *GetServiceObjectMeta) SetResourceVersion(v string) *GetServiceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *GetServiceObjectMeta) SetGeneration(v int64) *GetServiceObjectMeta {
  s.Generation = &v
  return s
}

func (s *GetServiceObjectMeta) SetCreationTimestamp(v string) *GetServiceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *GetServiceObjectMeta) SetDeletionTimestamp(v string) *GetServiceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *GetServiceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *GetServiceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *GetServiceObjectMeta) SetLabels(v map[string]*string) *GetServiceObjectMeta {
  s.Labels = v
  return s
}

func (s *GetServiceObjectMeta) SetAnnotations(v map[string]*string) *GetServiceObjectMeta {
  s.Annotations = v
  return s
}

func (s *GetServiceObjectMeta) SetOwnerReferences(v []*GetServiceOwnerReference) *GetServiceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *GetServiceObjectMeta) SetFinalizers(v []*string) *GetServiceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *GetServiceObjectMeta) SetClusterName(v string) *GetServiceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *GetServiceObjectMeta) SetManagedFields(v []*GetServiceManagedFieldsEntry) *GetServiceObjectMeta {
  s.ManagedFields = v
  return s
}

type GetServiceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this GetServiceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'GetServiceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“GetServiceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"GetServiceFieldsV1 holds the first JSON version format as described in the 'GetServiceFieldsV1' type", "zh_CN":"GetServiceFieldsV1 包含类型 “GetServiceFieldsV1” 中描述的第一个 JSON 版本格式"}
  GetServiceFieldsV1 *GetServiceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s GetServiceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s GetServiceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *GetServiceManagedFieldsEntry) SetManager(v string) *GetServiceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *GetServiceManagedFieldsEntry) SetOperation(v string) *GetServiceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *GetServiceManagedFieldsEntry) SetApiVersion(v string) *GetServiceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *GetServiceManagedFieldsEntry) SetTime(v string) *GetServiceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *GetServiceManagedFieldsEntry) SetFieldsType(v string) *GetServiceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *GetServiceManagedFieldsEntry) SetFieldsV1(v *GetServiceFieldsV1) *GetServiceManagedFieldsEntry {
  s.GetServiceFieldsV1 = v
  return s
}

func (s *GetServiceManagedFieldsEntry) SetSubresource(v string) *GetServiceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type GetServiceFieldsV1 struct {
}

func (s GetServiceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s GetServiceFieldsV1) GoString() string {
  return s.String()
}

type GetServiceOwnerReference struct {
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

func (s GetServiceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s GetServiceOwnerReference) GoString() string {
  return s.String()
}

func (s *GetServiceOwnerReference) SetApiVersion(v string) *GetServiceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *GetServiceOwnerReference) SetKind(v string) *GetServiceOwnerReference {
  s.Kind = &v
  return s
}

func (s *GetServiceOwnerReference) SetName(v string) *GetServiceOwnerReference {
  s.Name = &v
  return s
}

func (s *GetServiceOwnerReference) SetUid(v string) *GetServiceOwnerReference {
  s.Uid = &v
  return s
}

func (s *GetServiceOwnerReference) SetController(v bool) *GetServiceOwnerReference {
  s.Controller = &v
  return s
}

func (s *GetServiceOwnerReference) SetBlockOwnerDeletion(v bool) *GetServiceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type PutPatchServiceRequest struct {
}

func (s PutPatchServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceRequest) GoString() string {
  return s.String()
}

type PutPatchServiceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"service", "zh_CN":"service"}
  Data *PutPatchServiceService `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PutPatchServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceResponse) GoString() string {
  return s.String()
}

func (s *PutPatchServiceResponse) SetCode(v int64) *PutPatchServiceResponse {
  s.Code = &v
  return s
}

func (s *PutPatchServiceResponse) SetMsg(v string) *PutPatchServiceResponse {
  s.Msg = &v
  return s
}

func (s *PutPatchServiceResponse) SetRequestId(v string) *PutPatchServiceResponse {
  s.RequestId = &v
  return s
}

func (s *PutPatchServiceResponse) SetData(v *PutPatchServiceService) *PutPatchServiceResponse {
  s.Data = v
  return s
}

type PutPatchServicePaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"service name", "zh_CN":"service 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s PutPatchServicePaths) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServicePaths) GoString() string {
  return s.String()
}

func (s *PutPatchServicePaths) SetNamespace(v string) *PutPatchServicePaths {
  s.Namespace = &v
  return s
}

func (s *PutPatchServicePaths) SetName(v string) *PutPatchServicePaths {
  s.Name = &v
  return s
}

type PutPatchServiceParameters struct {
}

func (s PutPatchServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceParameters) GoString() string {
  return s.String()
}

type PutPatchServiceRequestHeader struct {
}

func (s PutPatchServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceRequestHeader) GoString() string {
  return s.String()
}

type PutPatchServiceResponseHeader struct {
}

func (s PutPatchServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceResponseHeader) GoString() string {
  return s.String()
}

type PutPatchServiceService struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *PutPatchServiceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"Spec defines the behavior of a service", "zh_CN":"spec 定义 PutPatchServiceService 的行为"}
  Spec *PutPatchServiceServiceSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Most recently observed status of the service. Populated by the system. Read-only", "zh_CN":"最近观察到的 PutPatchServiceService 状态。由系统填充。只读"}
  Status *PutPatchServiceServiceStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutPatchServiceService) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceService) GoString() string {
  return s.String()
}

func (s *PutPatchServiceService) SetApiVersion(v string) *PutPatchServiceService {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchServiceService) SetKind(v string) *PutPatchServiceService {
  s.Kind = &v
  return s
}

func (s *PutPatchServiceService) SetMetadata(v *PutPatchServiceObjectMeta) *PutPatchServiceService {
  s.Metadata = v
  return s
}

func (s *PutPatchServiceService) SetSpec(v *PutPatchServiceServiceSpec) *PutPatchServiceService {
  s.Spec = v
  return s
}

func (s *PutPatchServiceService) SetStatus(v *PutPatchServiceServiceStatus) *PutPatchServiceService {
  s.Status = v
  return s
}

type PutPatchServiceServiceStatus struct {
  // {"en":"Current service state", "zh_CN":"loadBalancer 包含负载均衡器的当前状态（如果存在）"}
  LoadBalancer *PutPatchServiceLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty"`
  // {"en":"LoadBalancer contains the current status of the load-balancer, if one is present", "zh_CN":"服务的当前状态"}
  Conditions []*PutPatchServiceMetaV1Condition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s PutPatchServiceServiceStatus) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceServiceStatus) GoString() string {
  return s.String()
}

func (s *PutPatchServiceServiceStatus) SetLoadBalancer(v *PutPatchServiceLoadBalancerStatus) *PutPatchServiceServiceStatus {
  s.LoadBalancer = v
  return s
}

func (s *PutPatchServiceServiceStatus) SetConditions(v []*PutPatchServiceMetaV1Condition) *PutPatchServiceServiceStatus {
  s.Conditions = v
  return s
}

type PutPatchServiceMetaV1Condition struct {
  // {"en":"type of condition in CamelCase or in foo.example.com/CamelCase", "zh_CN":"CamelCase 或 foo.example.com/CamelCase 中的条件类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status of the condition, one of True, False, Unknown", "zh_CN":"condition 的状态，True、False、Unknown 之一"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance", "zh_CN":"表示设置 condition 基于的 .metadata.generation 的过期次数。 例如，如果 .metadata.generation 当前为 12，但 .status.conditions[x].observedGeneration 为 9， 则 condition 相对于实例的当前状态已过期"}
  ObservedGeneration *int64 `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
  // {"en":"lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable", "zh_CN":"状况最近一次状态转化的时间。 变化应该发生在下层状况发生变化的时候。如果不知道下层状况发生变化的时间， 那么使用 API 字段更改的时间是可以接受的"}
  LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
  // {"en":"reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty", "zh_CN":"reason 包含一个程序标识符，指示 condition 最后一次转换的原因。 特定条件类型的生产者可以定义该字段的预期值和含义，以及这些值是否被视为有保证的 API。 该值应该是 CamelCase 字符串且不能为空"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":"message is a human readable message indicating details about the transition. This may be an empty string", "zh_CN":"message 是人类可读的消息，有关转换的详细信息，可以是空字符串"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s PutPatchServiceMetaV1Condition) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceMetaV1Condition) GoString() string {
  return s.String()
}

func (s *PutPatchServiceMetaV1Condition) SetType(v string) *PutPatchServiceMetaV1Condition {
  s.Type = &v
  return s
}

func (s *PutPatchServiceMetaV1Condition) SetStatus(v string) *PutPatchServiceMetaV1Condition {
  s.Status = &v
  return s
}

func (s *PutPatchServiceMetaV1Condition) SetObservedGeneration(v int64) *PutPatchServiceMetaV1Condition {
  s.ObservedGeneration = &v
  return s
}

func (s *PutPatchServiceMetaV1Condition) SetLastTransitionTime(v string) *PutPatchServiceMetaV1Condition {
  s.LastTransitionTime = &v
  return s
}

func (s *PutPatchServiceMetaV1Condition) SetReason(v string) *PutPatchServiceMetaV1Condition {
  s.Reason = &v
  return s
}

func (s *PutPatchServiceMetaV1Condition) SetMessage(v string) *PutPatchServiceMetaV1Condition {
  s.Message = &v
  return s
}

type PutPatchServiceLoadBalancerStatus struct {
  // {"en":"Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points", "zh_CN":"ingress 是一个包含负载均衡器 Ingress 点的列表。PutPatchServiceService 的流量需要被发送到这些 Ingress 点"}
  Ingress []*PutPatchServiceLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
}

func (s PutPatchServiceLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *PutPatchServiceLoadBalancerStatus) SetIngress(v []*PutPatchServiceLoadBalancerIngress) *PutPatchServiceLoadBalancerStatus {
  s.Ingress = v
  return s
}

type PutPatchServiceLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)", "zh_CN":"ip 是为基于 IP 的负载均衡器 Ingress 点（通常是 GCE 或 OpenStack 负载均衡器）设置的"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)", "zh_CN":"hostname 是为基于 DNS 的负载均衡器 Ingress 点（通常是 AWS 负载均衡器）设置的"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
  // {"en":"Ports is a list of records of service ports If used, every port defined in the service should have an entry in it", "zh_CN":"ports 是 PutPatchServiceService 的端口列表。如果设置了此字段，PutPatchServiceService 中定义的每个端口都应该在此列表中"}
  Ports []*PutPatchServicePortStatus `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s PutPatchServiceLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *PutPatchServiceLoadBalancerIngress) SetIp(v string) *PutPatchServiceLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *PutPatchServiceLoadBalancerIngress) SetHostName(v string) *PutPatchServiceLoadBalancerIngress {
  s.HostName = &v
  return s
}

func (s *PutPatchServiceLoadBalancerIngress) SetPorts(v []*PutPatchServicePortStatus) *PutPatchServiceLoadBalancerIngress {
  s.Ports = v
  return s
}

type PutPatchServicePortStatus struct {
  // {"en":"the port number of the service port of which status is recorded here", "zh_CN":"port 是所记录的服务端口状态的端口号"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"the protocol of the service port of which status is recorded here The supported values are: 'TCP', 'UDP', 'SCTP'", "zh_CN":"protocol 是所记录的服务端口状态的协议。支持的值为：“TCP”、”UDP”、“SCTP”"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use CamelCase names - cloud provider specific error values must have names that comply with the format foo.example.com/CamelCase", "zh_CN":"error 是记录 PutPatchServiceService 端口的问题。 错误的格式应符合以下规则:内置错误原因应在此文件中指定，应使用 CamelCase 名称。云提供商特定错误原因的名称必须符合格式 foo.example.com/CamelCase"}
  Error *string `json:"error,omitempty" xml:"error,omitempty"`
}

func (s PutPatchServicePortStatus) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServicePortStatus) GoString() string {
  return s.String()
}

func (s *PutPatchServicePortStatus) SetPort(v int32) *PutPatchServicePortStatus {
  s.Port = &v
  return s
}

func (s *PutPatchServicePortStatus) SetProtocol(v string) *PutPatchServicePortStatus {
  s.Protocol = &v
  return s
}

func (s *PutPatchServicePortStatus) SetError(v string) *PutPatchServicePortStatus {
  s.Error = &v
  return s
}

type PutPatchServiceServiceSpec struct {
  // {"en":"The list of ports that are exposed by this service", "zh_CN":"此 PutPatchServiceService 公开的端口列表"}
  Ports []*PutPatchServiceServicePort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en":"Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName", "zh_CN":"将 PutPatchServiceService 流量路由到具有与此 selector 匹配的标签键值对的 Pod。 如果为空或不存在，则假定该服务有一个外部进程管理其端点，Kubernetes 不会修改该端点。 仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型。如果类型为 ExternalName，则忽略"}
  Selector map[string]*string `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a PutPatchServiceService of type ExternalName, creation will fail. This field will be wiped when updating a PutPatchServiceService to type ExternalName", "zh_CN":"clusterIP 是服务的 IP 地址，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给服务，否则创建服务将失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIP 为空）或 type 已经是 ExternalName 时，可以更改 clusterIP（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIP 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 仅适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 PutPatchServiceService 时指定了 clusterIP，则创建将失败。 更新 PutPatchServiceService type 为 ExternalName 时，clusterIP 会被移除"}
  ClusterIP *string `json:"clusterIP,omitempty" xml:"clusterIP,omitempty"`
  // {"en":"ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a PutPatchServiceService of type ExternalName, creation will fail. This field will be wiped when updating a PutPatchServiceService to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value.This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"clusterIPs 是分配给该 PutPatchServiceService 的 IP 地址列表，通常是随机分配的。 如果地址是手动指定的，在范围内（根据系统配置），且没有被使用，它将被分配给 PutPatchServiceService；否则创建 PutPatchServiceService 失败。 clusterIP 一般不会被更改，除非 type 被更改为 ExternalName （ExternalName 需要 clusterIPs 为空）或 type 已经是 ExternalName 时，可以更改 clusterIPs（在这种情况下，可以选择指定此字段）。 可选值 “None”、空字符串 (“”) 或有效的 IP 地址。 clusterIPs 为 “None” 时会生成“无头服务”（无虚拟 IP），这在首选直接 Endpoint 连接且不需要代理时很有用。 适用于 ClusterIP、NodePort、和 LoadBalancer 类型的服务。 如果在创建 ExternalName 类型的 PutPatchServiceService 时指定了 clusterIPs，则会创建失败。 更新 PutPatchServiceService type 为 ExternalName 时，该字段将被移除。如果未指定此字段，则将从 clusterIP 字段初始化。 如果指定 clusterIPs，客户端必须确保 clusterIPs[0] 和 clusterIP 一致。clusterIPs 最多可包含两个条目（双栈系列，按任意顺序）。 这些 IP 必须与 ipFamilies 的值相对应。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 管理"}
  ClusterIPs []*string `json:"clusterIPs,omitempty" xml:"clusterIPs,omitempty" type:"Repeated"`
  // {"en":"type determines how the PutPatchServiceService is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services", "zh_CN":"type 确定 PutPatchServiceService 的公开方式。默认为 ClusterIP。 有效选项为 ExternalName、ClusterIP、NodePort 和 LoadBalancer。 “ClusterIP” 为端点分配一个集群内部 IP 地址用于负载均衡。 Endpoints 由 selector 确定，如果未设置 selector，则需要通过手动构造 Endpoints 或 EndpointSlice 的对象来确定。 如果 clusterIP 为 “None”，则不分配虚拟 IP，并且 Endpoints 作为一组端点而不是虚拟 IP 发布。 “NodePort” 建立在 ClusterIP 之上，并在每个节点上分配一个端口，该端口路由到与 clusterIP 相同的 Endpoints。 “LoadBalancer” 基于 NodePort 构建并创建一个外部负载均衡器（如果当前云支持），该负载均衡器路由到与 clusterIP 相同的 Endpoints。 “externalName” 将此 PutPatchServiceService 别名为指定的 externalName。其他几个字段不适用于 ExternalName PutPatchServiceService"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system", "zh_CN":"externalIPs 是一个 IP 列表，集群中的节点会为此 PutPatchServiceService 接收针对这些 IP 地址的流量。 这些 IP 不被 Kubernetes 管理。用户需要确保流量可以到达具有此 IP 的节点。 一个常见的例子是不属于 Kubernetes 系统的外部负载均衡器"}
  ExternalIPs []*string `json:"externalIPs,omitempty" xml:"externalIPs,omitempty" type:"Repeated"`
  // {"en":"Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None", "zh_CN":"支持 “ClientIP” 和 “None”。用于维护会话亲和性。 启用基于客户端 IP 的会话亲和性。必须是 ClientIP 或 None。默认为 None"}
  SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
  // {"en":"Only applies to PutPatchServiceService Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations, and it cannot support dual-stack. As of Kubernetes v1.24, users are encouraged to use implementation-specific annotations when available. This field may be removed in a future API version", "zh_CN":"仅适用于服务类型: LoadBalancer。此功能取决于底层云提供商是否支持负载均衡器。 如果云提供商不支持该功能，该字段将被忽略。 已弃用: 该字段信息不足，且其含义因实现而异，而且不支持双栈。 从 Kubernetes v1.24 开始，鼓励用户在可用时使用特定于实现的注释。在未来的 API 版本中可能会删除此字段"}
  LoadBalancerIP *string `json:"loadBalancerIP,omitempty" xml:"loadBalancerIP,omitempty"`
  // {"en":"If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature", "zh_CN":"如果设置了此字段并且被平台支持，将限制通过云厂商的负载均衡器的流量到指定的客户端 IP。 如果云提供商不支持该功能，该字段将被忽略"}
  LoadBalancerSourceRanges []*string `json:"loadBalancerSourceRanges,omitempty" xml:"loadBalancerSourceRanges,omitempty" type:"Repeated"`
  // {"en":"externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires type to be 'ExternalName'", "zh_CN":"externalName 是发现机制将返回的外部引用，作为此服务的别名（例如 DNS CNAME 记录）。 不涉及代理。必须是小写的 RFC-1123 主机名 (https://tools.ietf.org/html/rfc1123)， 并且要求 type 为 “ExternalName”"}
  ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
  // {"en":"externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the PutPatchServiceService's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node", "zh_CN":"externalTrafficPolicy 描述了节点如何分发它们在 PutPatchServiceService 的“外部访问”地址 （NodePort、ExternalIP 和 LoadBalancer IP）接收到的服务流量。 如果设置为 “Local”，代理将以一种假设外部负载均衡器将负责在节点之间服务流量负载均衡， 因此每个节点将仅向服务的节点本地端点传递流量，而不会伪装客户端源 IP。 （将丢弃错误发送到没有端点的节点的流量。） “Cluster” 默认值使用负载均衡路由到所有端点的策略（可能会根据拓扑和其他特性进行修改）。 请注意，从集群内部发送到 External IP 或 LoadBalancer IP 的流量始终具有 “Cluster” 语义， 但是从集群内部发送到 NodePort 的客户端需要在选择节点时考虑流量路由策略"}
  ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" xml:"externalTrafficPolicy,omitempty"`
  // {"en":"healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a PutPatchServiceService which does not need it, creation will fail. This field will be wiped when updating a PutPatchServiceService to no longer need it (e.g. changing type). This field cannot be updated once set", "zh_CN":"healthCheckNodePort 指定 PutPatchServiceService 的健康检查节点端口。 仅适用于 type 为 LoadBalancer 且 externalTrafficPolicy 设置为 Local 的情况。 如果为此字段设定了一个值，该值在合法范围内且没有被使用，则使用所指定的值。 如果未设置此字段，则自动分配字段值。外部系统（例如负载平衡器）可以使用此端口来确定给定节点是否拥有此服务的端点。 在创建不需要 healthCheckNodePort 的 PutPatchServiceService 时指定了此字段，则 PutPatchServiceService 创建会失败。 要移除 healthCheckNodePort，需要更改 PutPatchServiceService 的 type。 该字段一旦设置就无法更改"}
  HealthCheckNodePort *int32 `json:"healthCheckNodePort,omitempty" xml:"healthCheckNodePort,omitempty"`
  // {"en":"publishNotReadyAddresses indicates that any agent which deals with endpoints for this PutPatchServiceService should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless PutPatchServiceService to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior", "zh_CN":"publishNotReadyAddresses 表示任何处理此 PutPatchServiceService 端点的代理都应忽略任何准备就绪/未准备就绪的指示。 设置此字段的主要场景是为 StatefulSet 的服务提供支持，使之能够为其 Pod 传播 SRV DNS 记录，以实现对等发现。 为 PutPatchServiceService 生成 Endpoints 和 EndpointSlice 资源的 Kubernetes 控制器对字段的解读是， 即使 Pod 本身还没有准备好，所有端点都可被视为 “已就绪”。 对于代理而言，如果仅使用 Kubernetes 通过 Endpoints 或 EndpointSlice 资源所生成的端点， 则可以安全地假设这种行为"}
  PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty" xml:"publishNotReadyAddresses,omitempty"`
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"sessionAffinityConfig 包含会话亲和性的配置"}
  PutPatchServiceSessionAffinityConfig *PutPatchServiceSessionAffinityConfig `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
  // {"en":"IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the PutPatchServiceService. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a PutPatchServiceService to type ExternalName.This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field", "zh_CN":"iPFamilies 是分配给此服务的 IP 协议（例如 IPv4、IPv6）的列表。 该字段通常根据集群配置和 ipFamilyPolicy 字段自动设置。 如果手动指定该字段，且请求的协议在集群中可用，且 ipFamilyPolicy 允许，则使用；否则服务创建将失败。 该字段修改是有条件的：它允许添加或删除辅助 IP 协议，但不允许更改服务的主要 IP 协议。 有效值为 “IPv4” 和 “IPv6”。 该字段仅适用于 ClusterIP、NodePort 和 LoadBalancer 类型的服务，并且确实可用于“无头”服务。 更新服务设置类型为 ExternalName 时，该字段将被擦除。该字段最多可以包含两个条目（双栈系列，按任意顺序）。 如果指定，这些协议栈必须对应于 clusterIPs 字段的值。 clusterIP 和 ipFamilies 都由 ipFamilyPolicy 字段管理"}
  IpFamilies []*string `json:"ipFamilies,omitempty" xml:"ipFamilies,omitempty" type:"Repeated"`
  // {"en":"IPFamilyPolicy represents the dual-stack-ness requested or required by this PutPatchServiceService. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName", "zh_CN":"iPFamilyPolicy 表示此服务请求或要求的双栈特性。 如果没有提供值，则此字段将被设置为 SingleStack。 服务可以是 “SingleStack”（单个 IP 协议）、 “PreferDualStack”（双栈配置集群上的两个 IP 协议或单栈集群上的单个 IP 协议） 或 “RequireDualStack”（双栈上的两个 IP 协议配置的集群，否则失败）。 ipFamilies 和 clusterIPs 字段取决于此字段的值。 更新服务设置类型为 ExternalName 时，此字段将被擦除"}
  IpFamilyPolicy *string `json:"ipFamilyPolicy,omitempty" xml:"ipFamilyPolicy,omitempty"`
  // {"en":"allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type", "zh_CN":"allocateLoadBalancerNodePorts 定义了是否会自动为 LoadBalancer 类型的 PutPatchServiceService 分配 NodePort。默认为 true。 如果集群负载均衡器不依赖 NodePort，则可以设置此字段为 false。 如果调用者（通过指定一个值）请求特定的 NodePort，则无论此字段如何，都会接受这些请求。 该字段只能设置在 type 为 LoadBalancer 的 PutPatchServiceService 上，如果 type 更改为任何其他类型，该字段将被移除"}
  AllocateLoadBalancerNodePorts *bool `json:"allocateLoadBalancerNodePorts,omitempty" xml:"allocateLoadBalancerNodePorts,omitempty"`
  // {"en":"loadBalancerClass is the class of the load balancer implementation this PutPatchServiceService belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the PutPatchServiceService type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a PutPatchServiceService to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type", "zh_CN":"loadBalancerClass 是此 PutPatchServiceService 所属的负载均衡器实现的类。 如果设置了此字段，则字段值必须是标签风格的标识符，带有可选前缀，例如 ”internal-vip” 或 “example.com/internal-vip”。 无前缀名称是为最终用户保留的。该字段只能在 PutPatchServiceService 类型为 “LoadBalancer” 时设置。 如果未设置此字段，则使用默认负载均衡器实现。默认负载均衡器现在通常通过云提供商集成完成，但应适用于任何默认实现。 如果设置了此字段，则假定负载均衡器实现正在监测具有对应负载均衡器类的 PutPatchServiceService。 任何默认负载均衡器实现（例如云提供商）都应忽略设置此字段的 PutPatchServiceService。 只有在创建或更新的 PutPatchServiceService 的 type 为 “LoadBalancer” 时，才可设置此字段。 一经设定，不可更改。当 PutPatchServiceService 的 type 更新为 “LoadBalancer” 之外的其他类型时，此字段将被移除"}
  LoadBalancerClass *string `json:"loadBalancerClass,omitempty" xml:"loadBalancerClass,omitempty"`
  // {"en":"InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features)", "zh_CN":"InternalTrafficPolicy 描述节点如何分发它们在 ClusterIP 上接收到的服务流量。 如果设置为 “Local”，代理将假定 Pod 只想与在同一节点上的服务端点通信，如果没有本地端点，它将丢弃流量。 “Cluster” 默认将流量路由到所有端点（可能会根据拓扑和其他特性进行修改）"}
  InternalTrafficPolicy *string `json:"internalTrafficPolicy,omitempty" xml:"internalTrafficPolicy,omitempty"`
}

func (s PutPatchServiceServiceSpec) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceServiceSpec) GoString() string {
  return s.String()
}

func (s *PutPatchServiceServiceSpec) SetPorts(v []*PutPatchServiceServicePort) *PutPatchServiceServiceSpec {
  s.Ports = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetSelector(v map[string]*string) *PutPatchServiceServiceSpec {
  s.Selector = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetClusterIP(v string) *PutPatchServiceServiceSpec {
  s.ClusterIP = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetClusterIPs(v []*string) *PutPatchServiceServiceSpec {
  s.ClusterIPs = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetType(v string) *PutPatchServiceServiceSpec {
  s.Type = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetExternalIPs(v []*string) *PutPatchServiceServiceSpec {
  s.ExternalIPs = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetSessionAffinity(v string) *PutPatchServiceServiceSpec {
  s.SessionAffinity = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetLoadBalancerIP(v string) *PutPatchServiceServiceSpec {
  s.LoadBalancerIP = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetLoadBalancerSourceRanges(v []*string) *PutPatchServiceServiceSpec {
  s.LoadBalancerSourceRanges = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetExternalName(v string) *PutPatchServiceServiceSpec {
  s.ExternalName = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetExternalTrafficPolicy(v string) *PutPatchServiceServiceSpec {
  s.ExternalTrafficPolicy = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetHealthCheckNodePort(v int32) *PutPatchServiceServiceSpec {
  s.HealthCheckNodePort = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetPublishNotReadyAddresses(v bool) *PutPatchServiceServiceSpec {
  s.PublishNotReadyAddresses = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetSessionAffinityConfig(v *PutPatchServiceSessionAffinityConfig) *PutPatchServiceServiceSpec {
  s.PutPatchServiceSessionAffinityConfig = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetIpFamilies(v []*string) *PutPatchServiceServiceSpec {
  s.IpFamilies = v
  return s
}

func (s *PutPatchServiceServiceSpec) SetIpFamilyPolicy(v string) *PutPatchServiceServiceSpec {
  s.IpFamilyPolicy = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetAllocateLoadBalancerNodePorts(v bool) *PutPatchServiceServiceSpec {
  s.AllocateLoadBalancerNodePorts = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetLoadBalancerClass(v string) *PutPatchServiceServiceSpec {
  s.LoadBalancerClass = &v
  return s
}

func (s *PutPatchServiceServiceSpec) SetInternalTrafficPolicy(v string) *PutPatchServiceServiceSpec {
  s.InternalTrafficPolicy = &v
  return s
}

type PutPatchServiceServicePort struct {
  // {"en":"The name of this port within the service. This must be a DNS_LABEL. All ports within a PutPatchServiceServiceSpec must have unique names. When considering the endpoints for a PutPatchServiceService, this must match the 'name' field in the EndpointPort. Optional if only one PutPatchServiceServicePort is defined on this service", "zh_CN":"PutPatchServiceService 中此端口的名称。这必须是 DNS_LABEL。 PutPatchServiceServiceSpec 中的所有端口的名称都必须唯一。 在考虑 PutPatchServiceService 的端点时，这一字段值必须与 EndpointPort 中的 name 字段相同。 如果此服务上仅定义一个 PutPatchServiceServicePort，则为此字段为可选"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The IP protocol for this port. Supports TCP, UDP, and SCTP. Default is TCP", "zh_CN":"此端口的 IP 协议。支持 “TCP”、“UDP” 和 “SCTP”。默认为 TCP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol", "zh_CN":"此端口的应用协议，遵循标准的 Kubernetes 标签语法，无前缀名称按照 IANA 标准服务名称 （参见 RFC-6335 和 https://www.iana.org/assignments/service-names）。 非标准协议应该使用前缀名称，如 mycompany.com/my-custom-protocol"}
  AppProtocol *string `json:"appProtocol,omitempty" xml:"appProtocol,omitempty"`
  // {"en":"The port that will be exposed by this service", "zh_CN":"PutPatchServiceService 将公开的端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field", "zh_CN":"在 PutPatchServiceService 所针对的 Pod 上要访问的端口号或名称。 编号必须在 1 到 65535 的范围内。名称必须是 IANA_SVC_NAME。 如果此值是一个字符串，将在目标 Pod 的容器端口中作为命名端口进行查找。 如果未指定字段，则使用 “port” 字段的值（直接映射）。 对于 clusterIP 为 None 的服务，此字段将被忽略， 应忽略不设或设置为 “port” 字段的取值"}
  TargetPort *int32 `json:"targetPort,omitempty" xml:"targetPort,omitempty"`
  // {"en":"The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this PutPatchServiceService requires one. If this field is specified when creating a PutPatchServiceService which does not need it, creation will fail. This field will be wiped when updating a PutPatchServiceService to no longer need it (e.g. changing type from NodePort to ClusterIP).", "zh_CN":"当类型为 NodePort 或 LoadBalancer 时，PutPatchServiceService 公开在节点上的端口， 通常由系统分配。如果指定了一个在范围内且未使用的值，则将使用该值，否则操作将失败。 如果在创建的 PutPatchServiceService 需要该端口时未指定该字段，则会分配端口。 如果在创建不需要该端口的 Service时指定了该字段，则会创建失败。 当更新 PutPatchServiceService 时，如果不再需要此字段（例如，将类型从 NodePort 更改为 ClusterIP），这个字段将被擦除"}
  NodePort *int32 `json:"nodePort,omitempty" xml:"nodePort,omitempty"`
}

func (s PutPatchServiceServicePort) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceServicePort) GoString() string {
  return s.String()
}

func (s *PutPatchServiceServicePort) SetName(v string) *PutPatchServiceServicePort {
  s.Name = &v
  return s
}

func (s *PutPatchServiceServicePort) SetProtocol(v string) *PutPatchServiceServicePort {
  s.Protocol = &v
  return s
}

func (s *PutPatchServiceServicePort) SetAppProtocol(v string) *PutPatchServiceServicePort {
  s.AppProtocol = &v
  return s
}

func (s *PutPatchServiceServicePort) SetPort(v int32) *PutPatchServiceServicePort {
  s.Port = &v
  return s
}

func (s *PutPatchServiceServicePort) SetTargetPort(v int32) *PutPatchServiceServicePort {
  s.TargetPort = &v
  return s
}

func (s *PutPatchServiceServicePort) SetNodePort(v int32) *PutPatchServiceServicePort {
  s.NodePort = &v
  return s
}

type PutPatchServiceSessionAffinityConfig struct {
  // {"en":"sessionAffinityConfig contains the configurations of session affinity", "zh_CN":"clientIP 包含基于客户端 IP 的会话亲和性的配置"}
  ClientIP *PutPatchServiceClientIPConfig `json:"clientIP,omitempty" xml:"clientIP,omitempty"`
}

func (s PutPatchServiceSessionAffinityConfig) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceSessionAffinityConfig) GoString() string {
  return s.String()
}

func (s *PutPatchServiceSessionAffinityConfig) SetClientIP(v *PutPatchServiceClientIPConfig) *PutPatchServiceSessionAffinityConfig {
  s.ClientIP = v
  return s
}

type PutPatchServiceClientIPConfig struct {
  // {"en":"timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).", "zh_CN":"timeoutSeconds 指定 ClientIP 类型会话的维系时间秒数。 如果 ServiceAffinity == 'ClientIP'，则该值必须 >0 && <=86400（1 天）。默认值为 10800（3 小时）"}
  TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s PutPatchServiceClientIPConfig) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceClientIPConfig) GoString() string {
  return s.String()
}

func (s *PutPatchServiceClientIPConfig) SetTimeoutSeconds(v int32) *PutPatchServiceClientIPConfig {
  s.TimeoutSeconds = &v
  return s
}

type PutPatchServiceObjectMeta struct {
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
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 PutPatchServiceService 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*PutPatchServiceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*PutPatchServiceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s PutPatchServiceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceObjectMeta) GoString() string {
  return s.String()
}

func (s *PutPatchServiceObjectMeta) SetName(v string) *PutPatchServiceObjectMeta {
  s.Name = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetGenerateName(v string) *PutPatchServiceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetNamespace(v string) *PutPatchServiceObjectMeta {
  s.Namespace = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetSelfLink(v string) *PutPatchServiceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetUid(v string) *PutPatchServiceObjectMeta {
  s.Uid = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetResourceVersion(v string) *PutPatchServiceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetGeneration(v int64) *PutPatchServiceObjectMeta {
  s.Generation = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetCreationTimestamp(v string) *PutPatchServiceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetDeletionTimestamp(v string) *PutPatchServiceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *PutPatchServiceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetLabels(v map[string]*string) *PutPatchServiceObjectMeta {
  s.Labels = v
  return s
}

func (s *PutPatchServiceObjectMeta) SetAnnotations(v map[string]*string) *PutPatchServiceObjectMeta {
  s.Annotations = v
  return s
}

func (s *PutPatchServiceObjectMeta) SetOwnerReferences(v []*PutPatchServiceOwnerReference) *PutPatchServiceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *PutPatchServiceObjectMeta) SetFinalizers(v []*string) *PutPatchServiceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *PutPatchServiceObjectMeta) SetClusterName(v string) *PutPatchServiceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *PutPatchServiceObjectMeta) SetManagedFields(v []*PutPatchServiceManagedFieldsEntry) *PutPatchServiceObjectMeta {
  s.ManagedFields = v
  return s
}

type PutPatchServiceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this PutPatchServiceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'PutPatchServiceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“PutPatchServiceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"PutPatchServiceFieldsV1 holds the first JSON version format as described in the 'PutPatchServiceFieldsV1' type", "zh_CN":"PutPatchServiceFieldsV1 包含类型 “PutPatchServiceFieldsV1” 中描述的第一个 JSON 版本格式"}
  PutPatchServiceFieldsV1 *PutPatchServiceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s PutPatchServiceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *PutPatchServiceManagedFieldsEntry) SetManager(v string) *PutPatchServiceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *PutPatchServiceManagedFieldsEntry) SetOperation(v string) *PutPatchServiceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *PutPatchServiceManagedFieldsEntry) SetApiVersion(v string) *PutPatchServiceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchServiceManagedFieldsEntry) SetTime(v string) *PutPatchServiceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *PutPatchServiceManagedFieldsEntry) SetFieldsType(v string) *PutPatchServiceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *PutPatchServiceManagedFieldsEntry) SetFieldsV1(v *PutPatchServiceFieldsV1) *PutPatchServiceManagedFieldsEntry {
  s.PutPatchServiceFieldsV1 = v
  return s
}

func (s *PutPatchServiceManagedFieldsEntry) SetSubresource(v string) *PutPatchServiceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type PutPatchServiceFieldsV1 struct {
}

func (s PutPatchServiceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceFieldsV1) GoString() string {
  return s.String()
}

type PutPatchServiceOwnerReference struct {
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

func (s PutPatchServiceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s PutPatchServiceOwnerReference) GoString() string {
  return s.String()
}

func (s *PutPatchServiceOwnerReference) SetApiVersion(v string) *PutPatchServiceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchServiceOwnerReference) SetKind(v string) *PutPatchServiceOwnerReference {
  s.Kind = &v
  return s
}

func (s *PutPatchServiceOwnerReference) SetName(v string) *PutPatchServiceOwnerReference {
  s.Name = &v
  return s
}

func (s *PutPatchServiceOwnerReference) SetUid(v string) *PutPatchServiceOwnerReference {
  s.Uid = &v
  return s
}

func (s *PutPatchServiceOwnerReference) SetController(v bool) *PutPatchServiceOwnerReference {
  s.Controller = &v
  return s
}

func (s *PutPatchServiceOwnerReference) SetBlockOwnerDeletion(v bool) *PutPatchServiceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type PutNetworkPolicyRequest struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *PutNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this PutNetworkPolicyNetworkPolicy.", "zh_CN": "PutNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *PutNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s PutNetworkPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyRequest) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyRequest) SetApiVersion(v string) *PutNetworkPolicyRequest {
  s.ApiVersion = &v
  return s
}

func (s *PutNetworkPolicyRequest) SetKind(v string) *PutNetworkPolicyRequest {
  s.Kind = &v
  return s
}

func (s *PutNetworkPolicyRequest) SetMetadata(v *PutNetworkPolicyObjectMeta) *PutNetworkPolicyRequest {
  s.Metadata = v
  return s
}

func (s *PutNetworkPolicyRequest) SetSpec(v *PutNetworkPolicyNetworkPolicySpec) *PutNetworkPolicyRequest {
  s.Spec = v
  return s
}

type PutNetworkPolicyResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"networkPolicy", "zh_CN":"网络策略"}
  Data *PutNetworkPolicyNetworkPolicy `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PutNetworkPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyResponse) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyResponse) SetCode(v int64) *PutNetworkPolicyResponse {
  s.Code = &v
  return s
}

func (s *PutNetworkPolicyResponse) SetMsg(v string) *PutNetworkPolicyResponse {
  s.Msg = &v
  return s
}

func (s *PutNetworkPolicyResponse) SetRequestId(v string) *PutNetworkPolicyResponse {
  s.RequestId = &v
  return s
}

func (s *PutNetworkPolicyResponse) SetData(v *PutNetworkPolicyNetworkPolicy) *PutNetworkPolicyResponse {
  s.Data = v
  return s
}

type PutNetworkPolicyPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"networkPolicy name", "zh_CN":"networkPolicy 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s PutNetworkPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyPaths) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyPaths) SetNamespace(v string) *PutNetworkPolicyPaths {
  s.Namespace = &v
  return s
}

func (s *PutNetworkPolicyPaths) SetName(v string) *PutNetworkPolicyPaths {
  s.Name = &v
  return s
}

type PutNetworkPolicyParameters struct {
}

func (s PutNetworkPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyParameters) GoString() string {
  return s.String()
}

type PutNetworkPolicyRequestHeader struct {
}

func (s PutNetworkPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyRequestHeader) GoString() string {
  return s.String()
}

type PutNetworkPolicyResponseHeader struct {
}

func (s PutNetworkPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyResponseHeader) GoString() string {
  return s.String()
}

type PutNetworkPolicyNetworkPolicy struct {
  // {"en": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources", "zh_CN": "版本号"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds", "zh_CN": "被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata", "zh_CN": "标准的对象元数据"}
  Metadata *PutNetworkPolicyObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true"`
  // {"en": "Specification of the desired behavior for this PutNetworkPolicyNetworkPolicy.", "zh_CN": "PutNetworkPolicyNetworkPolicy 预期行为的规约。"}
  Spec *PutNetworkPolicyNetworkPolicySpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
}

func (s PutNetworkPolicyNetworkPolicy) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNetworkPolicy) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNetworkPolicy) SetApiVersion(v string) *PutNetworkPolicyNetworkPolicy {
  s.ApiVersion = &v
  return s
}

func (s *PutNetworkPolicyNetworkPolicy) SetKind(v string) *PutNetworkPolicyNetworkPolicy {
  s.Kind = &v
  return s
}

func (s *PutNetworkPolicyNetworkPolicy) SetMetadata(v *PutNetworkPolicyObjectMeta) *PutNetworkPolicyNetworkPolicy {
  s.Metadata = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicy) SetSpec(v *PutNetworkPolicyNetworkPolicySpec) *PutNetworkPolicyNetworkPolicy {
  s.Spec = v
  return s
}

type PutNetworkPolicyNetworkPolicySpec struct {
  // {"en": "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the PutNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this PutNetworkPolicyNetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8", "zh_CN": "出网规则"}
  Egress []*PutNetworkPolicyNetworkPolicyEgressRule `json:"egress,omitempty" xml:"egress,omitempty" type:"Repeated"`
  // {"en": "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the PutNetworkPolicyNetworkPolicy objects whose podSelector matches the pod. If this field is empty then this PutNetworkPolicyNetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)", "zh_CN": "入网规则"}
  Ingress []*PutNetworkPolicyNetworkPolicyIngressRule `json:"ingress,omitempty" xml:"ingress,omitempty" type:"Repeated"`
  // {"en": "Selects the pods to which this PutNetworkPolicyNetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.", "zh_CN": "限制pod的选择器"}
  PodSelector *PutNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty" require:"true"`
  // {"en": "List of rule types that the PutNetworkPolicyNetworkPolicy relates to. Valid options are Ingress, Egress, or Ingress,Egress. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ Egress ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include Egress (since such a policy would not include an Egress section and would otherwise default to just [ Ingress ]). This field is beta-level in 1.8", "zh_CN": "策略类型"}
  PolicyTypes []*string `json:"policyTypes,omitempty" xml:"policyTypes,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyNetworkPolicySpec) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNetworkPolicySpec) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNetworkPolicySpec) SetEgress(v []*PutNetworkPolicyNetworkPolicyEgressRule) *PutNetworkPolicyNetworkPolicySpec {
  s.Egress = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicySpec) SetIngress(v []*PutNetworkPolicyNetworkPolicyIngressRule) *PutNetworkPolicyNetworkPolicySpec {
  s.Ingress = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicySpec) SetPodSelector(v *PutNetworkPolicyPodLabelSelector) *PutNetworkPolicyNetworkPolicySpec {
  s.PodSelector = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicySpec) SetPolicyTypes(v []*string) *PutNetworkPolicyNetworkPolicySpec {
  s.PolicyTypes = v
  return s
}

type PutNetworkPolicyNetworkPolicyIngressRule struct {
  // {"en": "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.", "zh_CN": "入网规则信息"}
  From []*PutNetworkPolicyNetworkPolicyPeer `json:"from,omitempty" xml:"from,omitempty" type:"Repeated"`
  // {"en": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*PutNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyNetworkPolicyIngressRule) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNetworkPolicyIngressRule) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNetworkPolicyIngressRule) SetFrom(v []*PutNetworkPolicyNetworkPolicyPeer) *PutNetworkPolicyNetworkPolicyIngressRule {
  s.From = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicyIngressRule) SetPorts(v []*PutNetworkPolicyNetworkPolicyPort) *PutNetworkPolicyNetworkPolicyIngressRule {
  s.Ports = v
  return s
}

type PutNetworkPolicyNetworkPolicyEgressRule struct {
  // {"en": "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.", "zh_CN": "限制端口"}
  Ports []*PutNetworkPolicyNetworkPolicyPort `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
  // {"en": "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.", "zh_CN": "出网规则信息"}
  To []*PutNetworkPolicyNetworkPolicyPeer `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyNetworkPolicyEgressRule) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNetworkPolicyEgressRule) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNetworkPolicyEgressRule) SetPorts(v []*PutNetworkPolicyNetworkPolicyPort) *PutNetworkPolicyNetworkPolicyEgressRule {
  s.Ports = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicyEgressRule) SetTo(v []*PutNetworkPolicyNetworkPolicyPeer) *PutNetworkPolicyNetworkPolicyEgressRule {
  s.To = v
  return s
}

type PutNetworkPolicyNetworkPolicyPeer struct {
  // {"en": "PutNetworkPolicyIPBlock defines policy on a particular PutNetworkPolicyIPBlock. If this field is set then neither of the other fields can be.", "zh_CN": "IP规则"}
  IpBlock *PutNetworkPolicyIPBlock `json:"ipBlock,omitempty" xml:"ipBlock,omitempty"`
  // {"en": "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.If PodSelector is also set, then the PutNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.", "zh_CN": "namespace选择器"}
  NamespaceSelector *PutNetworkPolicyNsLabelSelector `json:"namespaceSelector,omitempty" xml:"namespaceSelector,omitempty"`
  // {"en": "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.If NamespaceSelector is also set, then the PutNetworkPolicyNetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.", "zh_CN": "pod选择器"}
  PodSelector *PutNetworkPolicyPodLabelSelector `json:"podSelector,omitempty" xml:"podSelector,omitempty"`
}

func (s PutNetworkPolicyNetworkPolicyPeer) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNetworkPolicyPeer) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNetworkPolicyPeer) SetIpBlock(v *PutNetworkPolicyIPBlock) *PutNetworkPolicyNetworkPolicyPeer {
  s.IpBlock = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicyPeer) SetNamespaceSelector(v *PutNetworkPolicyNsLabelSelector) *PutNetworkPolicyNetworkPolicyPeer {
  s.NamespaceSelector = v
  return s
}

func (s *PutNetworkPolicyNetworkPolicyPeer) SetPodSelector(v *PutNetworkPolicyPodLabelSelector) *PutNetworkPolicyNetworkPolicyPeer {
  s.PodSelector = v
  return s
}

type PutNetworkPolicyIPBlock struct {
  // {"en": "CIDR is a string representing the IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64", "zh_CN": "生效IP网段"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty" require:"true"`
  // {"en": "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are 192.168.1.1/24 or 2001:db9::/64 Except values will be rejected if they are outside the CIDR range", "zh_CN": "例外IP网段"}
  Except []*string `json:"except,omitempty" xml:"except,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyIPBlock) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyIPBlock) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyIPBlock) SetCidr(v string) *PutNetworkPolicyIPBlock {
  s.Cidr = &v
  return s
}

func (s *PutNetworkPolicyIPBlock) SetExcept(v []*string) *PutNetworkPolicyIPBlock {
  s.Except = v
  return s
}

type PutNetworkPolicyNetworkPolicyPort struct {
  // {"en": "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.", "zh_CN": "端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
  // {"en": "The protocol (TCP, UDP) which traffic must match. If not specified, this field defaults to TCP.", "zh_CN": "协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s PutNetworkPolicyNetworkPolicyPort) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNetworkPolicyPort) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNetworkPolicyPort) SetPort(v string) *PutNetworkPolicyNetworkPolicyPort {
  s.Port = &v
  return s
}

func (s *PutNetworkPolicyNetworkPolicyPort) SetProtocol(v string) *PutNetworkPolicyNetworkPolicyPort {
  s.Protocol = &v
  return s
}

type PutNetworkPolicyPodLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
}

func (s PutNetworkPolicyPodLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyPodLabelSelector) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyPodLabelSelector) SetMatchLabels(v map[string]*string) *PutNetworkPolicyPodLabelSelector {
  s.MatchLabels = v
  return s
}

type PutNetworkPolicyNsLabelSelector struct {
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*PutNetworkPolicyLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyNsLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyNsLabelSelector) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyNsLabelSelector) SetMatchExpressions(v []*PutNetworkPolicyLabelSelectorRequirement) *PutNetworkPolicyNsLabelSelector {
  s.MatchExpressions = v
  return s
}

type PutNetworkPolicyLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyLabelSelectorRequirement) SetKey(v string) *PutNetworkPolicyLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *PutNetworkPolicyLabelSelectorRequirement) SetOperator(v string) *PutNetworkPolicyLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *PutNetworkPolicyLabelSelectorRequirement) SetValues(v []*string) *PutNetworkPolicyLabelSelectorRequirement {
  s.Values = v
  return s
}

type PutNetworkPolicyObjectMeta struct {
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
  OwnerReferences []*PutNetworkPolicyOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*PutNetworkPolicyManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s PutNetworkPolicyObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyObjectMeta) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyObjectMeta) SetName(v string) *PutNetworkPolicyObjectMeta {
  s.Name = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetGenerateName(v string) *PutNetworkPolicyObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetNamespace(v string) *PutNetworkPolicyObjectMeta {
  s.Namespace = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetSelfLink(v string) *PutNetworkPolicyObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetUid(v string) *PutNetworkPolicyObjectMeta {
  s.Uid = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetResourceVersion(v string) *PutNetworkPolicyObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetGeneration(v int64) *PutNetworkPolicyObjectMeta {
  s.Generation = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetCreationTimestamp(v string) *PutNetworkPolicyObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetDeletionTimestamp(v string) *PutNetworkPolicyObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetDeletionGracePeriodSeconds(v int64) *PutNetworkPolicyObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetLabels(v map[string]*string) *PutNetworkPolicyObjectMeta {
  s.Labels = v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetAnnotations(v map[string]*string) *PutNetworkPolicyObjectMeta {
  s.Annotations = v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetOwnerReferences(v []*PutNetworkPolicyOwnerReference) *PutNetworkPolicyObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetFinalizers(v []*string) *PutNetworkPolicyObjectMeta {
  s.Finalizers = v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetClusterName(v string) *PutNetworkPolicyObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *PutNetworkPolicyObjectMeta) SetManagedFields(v []*PutNetworkPolicyManagedFieldsEntry) *PutNetworkPolicyObjectMeta {
  s.ManagedFields = v
  return s
}

type PutNetworkPolicyOwnerReference struct {
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

func (s PutNetworkPolicyOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyOwnerReference) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyOwnerReference) SetApiVersion(v string) *PutNetworkPolicyOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *PutNetworkPolicyOwnerReference) SetKind(v string) *PutNetworkPolicyOwnerReference {
  s.Kind = &v
  return s
}

func (s *PutNetworkPolicyOwnerReference) SetName(v string) *PutNetworkPolicyOwnerReference {
  s.Name = &v
  return s
}

func (s *PutNetworkPolicyOwnerReference) SetUid(v string) *PutNetworkPolicyOwnerReference {
  s.Uid = &v
  return s
}

func (s *PutNetworkPolicyOwnerReference) SetController(v bool) *PutNetworkPolicyOwnerReference {
  s.Controller = &v
  return s
}

func (s *PutNetworkPolicyOwnerReference) SetBlockOwnerDeletion(v bool) *PutNetworkPolicyOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type PutNetworkPolicyManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this PutNetworkPolicyManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'PutNetworkPolicyFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“PutNetworkPolicyFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"PutNetworkPolicyFieldsV1 holds the first JSON version format as described in the 'PutNetworkPolicyFieldsV1' type", "zh_CN":"PutNetworkPolicyFieldsV1 包含类型 “PutNetworkPolicyFieldsV1” 中描述的第一个 JSON 版本格式"}
  PutNetworkPolicyFieldsV1 *PutNetworkPolicyFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s PutNetworkPolicyManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetManager(v string) *PutNetworkPolicyManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetOperation(v string) *PutNetworkPolicyManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetApiVersion(v string) *PutNetworkPolicyManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetTime(v string) *PutNetworkPolicyManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetFieldsType(v string) *PutNetworkPolicyManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetFieldsV1(v *PutNetworkPolicyFieldsV1) *PutNetworkPolicyManagedFieldsEntry {
  s.PutNetworkPolicyFieldsV1 = v
  return s
}

func (s *PutNetworkPolicyManagedFieldsEntry) SetSubresource(v string) *PutNetworkPolicyManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type PutNetworkPolicyFieldsV1 struct {
}

func (s PutNetworkPolicyFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s PutNetworkPolicyFieldsV1) GoString() string {
  return s.String()
}




type VMPUnassignEdgeIPRequest struct {
  // {"en":"target virtual machine external network Ip","zh_CN":"目标实例外网Ip"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional Ip to unbind","zh_CN":"要解除绑定的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPUnassignEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPUnassignEdgeIPRequest) SetServerIp(v string) *VMPUnassignEdgeIPRequest {
  s.ServerIp = &v
  return s
}

func (s *VMPUnassignEdgeIPRequest) SetEdgeIps(v []*string) *VMPUnassignEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type VMPUnassignEdgeIPRequestHeader struct {
}

func (s VMPUnassignEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPUnassignEdgeIPPaths struct {
}

func (s VMPUnassignEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPUnassignEdgeIPParameters struct {
}

func (s VMPUnassignEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPUnassignEdgeIPResponse struct {
  // {"en":"target virtual machine IP","zh_CN":"目标实例公网IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional Ip that has been bound to the virtual machine","zh_CN":"已绑定到实例的额外公网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPUnassignEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPUnassignEdgeIPResponse) SetServerIp(v string) *VMPUnassignEdgeIPResponse {
  s.ServerIp = &v
  return s
}

func (s *VMPUnassignEdgeIPResponse) SetEdgeIps(v []*string) *VMPUnassignEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPUnassignEdgeIPResponseHeader struct {
}

func (s VMPUnassignEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type GetIngressRequest struct {
}

func (s GetIngressRequest) String() string {
  return tea.Prettify(s)
}

func (s GetIngressRequest) GoString() string {
  return s.String()
}

type GetIngressResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress object", "zh_CN":"路由对象"}
  Data *GetIngressCustomIngressDetail `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetIngressResponse) String() string {
  return tea.Prettify(s)
}

func (s GetIngressResponse) GoString() string {
  return s.String()
}

func (s *GetIngressResponse) SetCode(v int64) *GetIngressResponse {
  s.Code = &v
  return s
}

func (s *GetIngressResponse) SetMsg(v string) *GetIngressResponse {
  s.Msg = &v
  return s
}

func (s *GetIngressResponse) SetRequestId(v string) *GetIngressResponse {
  s.RequestId = &v
  return s
}

func (s *GetIngressResponse) SetData(v *GetIngressCustomIngressDetail) *GetIngressResponse {
  s.Data = v
  return s
}

type GetIngressPaths struct {
  // {"en":"ingress name", "zh_CN":"路由名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetIngressPaths) String() string {
  return tea.Prettify(s)
}

func (s GetIngressPaths) GoString() string {
  return s.String()
}

func (s *GetIngressPaths) SetName(v string) *GetIngressPaths {
  s.Name = &v
  return s
}

type GetIngressParameters struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s GetIngressParameters) String() string {
  return tea.Prettify(s)
}

func (s GetIngressParameters) GoString() string {
  return s.String()
}

func (s *GetIngressParameters) SetNamespace(v string) *GetIngressParameters {
  s.Namespace = &v
  return s
}

type GetIngressRequestHeader struct {
}

func (s GetIngressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIngressRequestHeader) GoString() string {
  return s.String()
}

type GetIngressResponseHeader struct {
}

func (s GetIngressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIngressResponseHeader) GoString() string {
  return s.String()
}

type GetIngressCustomIngressDetail struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  ControllerName *string `json:"controllerName,omitempty" xml:"controllerName,omitempty" require:"true"`
  // {"en":"clusters", "zh_CN":"集群列表"}
  Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"ingress", "zh_CN":"路由"}
  GetIngressIngress *GetIngressIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s GetIngressCustomIngressDetail) String() string {
  return tea.Prettify(s)
}

func (s GetIngressCustomIngressDetail) GoString() string {
  return s.String()
}

func (s *GetIngressCustomIngressDetail) SetControllerName(v string) *GetIngressCustomIngressDetail {
  s.ControllerName = &v
  return s
}

func (s *GetIngressCustomIngressDetail) SetClusters(v []*string) *GetIngressCustomIngressDetail {
  s.Clusters = v
  return s
}

func (s *GetIngressCustomIngressDetail) SetIngress(v *GetIngressIngress) *GetIngressCustomIngressDetail {
  s.GetIngressIngress = v
  return s
}

func (s *GetIngressCustomIngressDetail) SetCreateTime(v int64) *GetIngressCustomIngressDetail {
  s.CreateTime = &v
  return s
}

func (s *GetIngressCustomIngressDetail) SetUpdateTime(v int64) *GetIngressCustomIngressDetail {
  s.UpdateTime = &v
  return s
}

type GetIngressIngress struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *GetIngressObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"ingress desired", "zh_CN":"路由期望属性"}
  Spec *GetIngressIngressSpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
  // {"en":"Status is the current state of the GetIngressIngress", "zh_CN":"路由状态"}
  Status *GetIngressIngressStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetIngressIngress) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngress) GoString() string {
  return s.String()
}

func (s *GetIngressIngress) SetApiVersion(v string) *GetIngressIngress {
  s.ApiVersion = &v
  return s
}

func (s *GetIngressIngress) SetKind(v string) *GetIngressIngress {
  s.Kind = &v
  return s
}

func (s *GetIngressIngress) SetMetadata(v *GetIngressObjectMeta) *GetIngressIngress {
  s.Metadata = v
  return s
}

func (s *GetIngressIngress) SetSpec(v *GetIngressIngressSpec) *GetIngressIngress {
  s.Spec = v
  return s
}

func (s *GetIngressIngress) SetStatus(v *GetIngressIngressStatus) *GetIngressIngress {
  s.Status = v
  return s
}

type GetIngressIngressSpec struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  IngressClassName *string `json:"ingressClassName,omitempty" xml:"ingressClassName,omitempty" require:"true"`
  // {"en":"DefaultBackend is the backend that should handle requests that don't match any rule", "zh_CN":"默认后端,当请求不匹配任何规则时调用"}
  DefaultBackend *GetIngressIngressBackend `json:"defaultBackend,omitempty" xml:"defaultBackend,omitempty" require:"true"`
  Tls []*GetIngressIngressTLS `json:"tls,omitempty" xml:"tls,omitempty" require:"true" type:"Repeated"`
  // {"en":"A list of host rules used to configure the GetIngressIngress", "zh_CN":"路由规则列表"}
  Rules []*GetIngressIngressRule `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Repeated"`
}

func (s GetIngressIngressSpec) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngressSpec) GoString() string {
  return s.String()
}

func (s *GetIngressIngressSpec) SetIngressClassName(v string) *GetIngressIngressSpec {
  s.IngressClassName = &v
  return s
}

func (s *GetIngressIngressSpec) SetDefaultBackend(v *GetIngressIngressBackend) *GetIngressIngressSpec {
  s.DefaultBackend = v
  return s
}

func (s *GetIngressIngressSpec) SetTls(v []*GetIngressIngressTLS) *GetIngressIngressSpec {
  s.Tls = v
  return s
}

func (s *GetIngressIngressSpec) SetRules(v []*GetIngressIngressRule) *GetIngressIngressSpec {
  s.Rules = v
  return s
}

type GetIngressIngressRule struct {
  // {"en":"Host is the fully qualified domain name of a network host", "zh_CN":"域名"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  Http *GetIngressHTTPIngressRuleValue `json:"http,omitempty" xml:"http,omitempty" require:"true"`
}

func (s GetIngressIngressRule) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngressRule) GoString() string {
  return s.String()
}

func (s *GetIngressIngressRule) SetHost(v string) *GetIngressIngressRule {
  s.Host = &v
  return s
}

func (s *GetIngressIngressRule) SetHttp(v *GetIngressHTTPIngressRuleValue) *GetIngressIngressRule {
  s.Http = v
  return s
}

type GetIngressHTTPIngressRuleValue struct {
  // {"en":"A collection of paths that map requests to backends", "zh_CN":"请求路径匹配规则"}
  GetIngressPaths []*GetIngressHTTPIngressPath `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s GetIngressHTTPIngressRuleValue) String() string {
  return tea.Prettify(s)
}

func (s GetIngressHTTPIngressRuleValue) GoString() string {
  return s.String()
}

func (s *GetIngressHTTPIngressRuleValue) SetGetIngressPaths(v []*GetIngressHTTPIngressPath) *GetIngressHTTPIngressRuleValue {
  s.GetIngressPaths = v
  return s
}

type GetIngressHTTPIngressPath struct {
  // {"en":"Path is matched against the path of an incoming request", "zh_CN":"请求路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"PathType determines the interpretation of the Path matching,PathType can be one of the following values: Exact,Prefix,ImplementationSpecific", "zh_CN":"路径匹配类型: Exact,Prefix,ImplementationSpecific"}
  PathType *string `json:"pathType,omitempty" xml:"pathType,omitempty" require:"true"`
  // {"en":"Backend defines the referenced service endpoint to which the traffic will be forwarded to", "zh_CN":"指定后端服务"}
  Backend *GetIngressIngressBackend `json:"backend,omitempty" xml:"backend,omitempty" require:"true"`
}

func (s GetIngressHTTPIngressPath) String() string {
  return tea.Prettify(s)
}

func (s GetIngressHTTPIngressPath) GoString() string {
  return s.String()
}

func (s *GetIngressHTTPIngressPath) SetPath(v string) *GetIngressHTTPIngressPath {
  s.Path = &v
  return s
}

func (s *GetIngressHTTPIngressPath) SetPathType(v string) *GetIngressHTTPIngressPath {
  s.PathType = &v
  return s
}

func (s *GetIngressHTTPIngressPath) SetBackend(v *GetIngressIngressBackend) *GetIngressHTTPIngressPath {
  s.Backend = v
  return s
}

type GetIngressIngressTLS struct {
  // {"en":"Hosts are a list of hosts included in the TLS certificate", "zh_CN":"tls证书包含域名"}
  Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" require:"true" type:"Repeated"`
  // {"en":"SecretName is the name of the secret used to terminate TLS traffic on port 443", "zh_CN":"tls秘钥名称"}
  SecretName *string `json:"secretName,omitempty" xml:"secretName,omitempty" require:"true"`
}

func (s GetIngressIngressTLS) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngressTLS) GoString() string {
  return s.String()
}

func (s *GetIngressIngressTLS) SetHosts(v []*string) *GetIngressIngressTLS {
  s.Hosts = v
  return s
}

func (s *GetIngressIngressTLS) SetSecretName(v string) *GetIngressIngressTLS {
  s.SecretName = &v
  return s
}

type GetIngressIngressBackend struct {
  // {"en":"Service references a Service as a Backend", "zh_CN":"指定后端服务"}
  Service *GetIngressIngressServiceBackend `json:"service,omitempty" xml:"service,omitempty" require:"true"`
  // {"en":"Resource is an ObjectRef to another Kubernetes resource in the namespace of the GetIngressIngress object", "zh_CN":"路由指定后端资源"}
  Resource *GetIngressTypedLocalObjectReference `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
}

func (s GetIngressIngressBackend) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngressBackend) GoString() string {
  return s.String()
}

func (s *GetIngressIngressBackend) SetService(v *GetIngressIngressServiceBackend) *GetIngressIngressBackend {
  s.Service = v
  return s
}

func (s *GetIngressIngressBackend) SetResource(v *GetIngressTypedLocalObjectReference) *GetIngressIngressBackend {
  s.Resource = v
  return s
}

type GetIngressTypedLocalObjectReference struct {
  // {"en":"Name is the name of resource being referenced", "zh_CN":"资源名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Kind is the type of resource being referenced", "zh_CN":"资源类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"APIGroup is the group for the resource being referenced", "zh_CN":"资源分组"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty" require:"true"`
}

func (s GetIngressTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s GetIngressTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *GetIngressTypedLocalObjectReference) SetName(v string) *GetIngressTypedLocalObjectReference {
  s.Name = &v
  return s
}

func (s *GetIngressTypedLocalObjectReference) SetKind(v string) *GetIngressTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *GetIngressTypedLocalObjectReference) SetApiGroup(v string) *GetIngressTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

type GetIngressIngressServiceBackend struct {
  // {"en":"Name is the referenced service", "zh_CN":"服务名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Port of the referenced service. A port name or port number", "zh_CN":"服务端口或端口名称"}
  Port *GetIngressServiceBackendPort `json:"port,omitempty" xml:"port,omitempty" require:"true"`
}

func (s GetIngressIngressServiceBackend) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngressServiceBackend) GoString() string {
  return s.String()
}

func (s *GetIngressIngressServiceBackend) SetName(v string) *GetIngressIngressServiceBackend {
  s.Name = &v
  return s
}

func (s *GetIngressIngressServiceBackend) SetPort(v *GetIngressServiceBackendPort) *GetIngressIngressServiceBackend {
  s.Port = v
  return s
}

type GetIngressServiceBackendPort struct {
  // {"en":"Name is the name of the port on the Service", "zh_CN":"服务端口名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Number is the numerical port number on the Service", "zh_CN":"服务数字端口"}
  Number *int32 `json:"number,omitempty" xml:"number,omitempty" require:"true"`
}

func (s GetIngressServiceBackendPort) String() string {
  return tea.Prettify(s)
}

func (s GetIngressServiceBackendPort) GoString() string {
  return s.String()
}

func (s *GetIngressServiceBackendPort) SetName(v string) *GetIngressServiceBackendPort {
  s.Name = &v
  return s
}

func (s *GetIngressServiceBackendPort) SetNumber(v int32) *GetIngressServiceBackendPort {
  s.Number = &v
  return s
}

type GetIngressIngressStatus struct {
  // {"en":"LoadBalancer contains the current status of the load-balancer", "zh_CN":"包含当前负载均衡服务的状态"}
  LoadBalancer *GetIngressLoadBalancerStatus `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty" require:"true"`
}

func (s GetIngressIngressStatus) String() string {
  return tea.Prettify(s)
}

func (s GetIngressIngressStatus) GoString() string {
  return s.String()
}

func (s *GetIngressIngressStatus) SetLoadBalancer(v *GetIngressLoadBalancerStatus) *GetIngressIngressStatus {
  s.LoadBalancer = v
  return s
}

type GetIngressLoadBalancerStatus struct {
  GetIngressIngress []*GetIngressLoadBalancerIngress `json:"ingress,omitempty" xml:"ingress,omitempty" require:"true" type:"Repeated"`
}

func (s GetIngressLoadBalancerStatus) String() string {
  return tea.Prettify(s)
}

func (s GetIngressLoadBalancerStatus) GoString() string {
  return s.String()
}

func (s *GetIngressLoadBalancerStatus) SetIngress(v []*GetIngressLoadBalancerIngress) *GetIngressLoadBalancerStatus {
  s.GetIngressIngress = v
  return s
}

type GetIngressLoadBalancerIngress struct {
  // {"en":"IP is set for load-balancer ingress points that are IP based", "zh_CN":"负载均衡服务ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Hostname is set for load-balancer ingress points that are DNS based", "zh_CN":"负载均衡类型服务dns"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Ports is a list of records of service ports", "zh_CN":"服务端口状态列表"}
  Ports []*GetIngressPortStatus `json:"ports,omitempty" xml:"ports,omitempty" require:"true" type:"Repeated"`
}

func (s GetIngressLoadBalancerIngress) String() string {
  return tea.Prettify(s)
}

func (s GetIngressLoadBalancerIngress) GoString() string {
  return s.String()
}

func (s *GetIngressLoadBalancerIngress) SetIp(v string) *GetIngressLoadBalancerIngress {
  s.Ip = &v
  return s
}

func (s *GetIngressLoadBalancerIngress) SetHostname(v string) *GetIngressLoadBalancerIngress {
  s.Hostname = &v
  return s
}

func (s *GetIngressLoadBalancerIngress) SetPorts(v []*GetIngressPortStatus) *GetIngressLoadBalancerIngress {
  s.Ports = v
  return s
}

type GetIngressPortStatus struct {
  // {"en":"Port is the port number of the service port of which status is recorded here", "zh_CN":"服务端口"}
  Port *int32 `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Protocol is the protocol of the service port of which status is recorded here,The supported values are: TCP, UDP, SCTP", "zh_CN":"服务支持类型: TCP,UDP,SCTP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"Error is to record the problem with the service port", "zh_CN":"记录服务端口错误信息"}
  Error *string `json:"error,omitempty" xml:"error,omitempty" require:"true"`
}

func (s GetIngressPortStatus) String() string {
  return tea.Prettify(s)
}

func (s GetIngressPortStatus) GoString() string {
  return s.String()
}

func (s *GetIngressPortStatus) SetPort(v int32) *GetIngressPortStatus {
  s.Port = &v
  return s
}

func (s *GetIngressPortStatus) SetProtocol(v string) *GetIngressPortStatus {
  s.Protocol = &v
  return s
}

func (s *GetIngressPortStatus) SetError(v string) *GetIngressPortStatus {
  s.Error = &v
  return s
}

type GetIngressObjectMeta struct {
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
  OwnerReferences []*GetIngressOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*GetIngressManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s GetIngressObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s GetIngressObjectMeta) GoString() string {
  return s.String()
}

func (s *GetIngressObjectMeta) SetName(v string) *GetIngressObjectMeta {
  s.Name = &v
  return s
}

func (s *GetIngressObjectMeta) SetGenerateName(v string) *GetIngressObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *GetIngressObjectMeta) SetNamespace(v string) *GetIngressObjectMeta {
  s.Namespace = &v
  return s
}

func (s *GetIngressObjectMeta) SetSelfLink(v string) *GetIngressObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *GetIngressObjectMeta) SetUid(v string) *GetIngressObjectMeta {
  s.Uid = &v
  return s
}

func (s *GetIngressObjectMeta) SetResourceVersion(v string) *GetIngressObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *GetIngressObjectMeta) SetGeneration(v int64) *GetIngressObjectMeta {
  s.Generation = &v
  return s
}

func (s *GetIngressObjectMeta) SetCreationTimestamp(v string) *GetIngressObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *GetIngressObjectMeta) SetDeletionTimestamp(v string) *GetIngressObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *GetIngressObjectMeta) SetDeletionGracePeriodSeconds(v int64) *GetIngressObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *GetIngressObjectMeta) SetLabels(v map[string]*string) *GetIngressObjectMeta {
  s.Labels = v
  return s
}

func (s *GetIngressObjectMeta) SetAnnotations(v map[string]*string) *GetIngressObjectMeta {
  s.Annotations = v
  return s
}

func (s *GetIngressObjectMeta) SetOwnerReferences(v []*GetIngressOwnerReference) *GetIngressObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *GetIngressObjectMeta) SetFinalizers(v []*string) *GetIngressObjectMeta {
  s.Finalizers = v
  return s
}

func (s *GetIngressObjectMeta) SetClusterName(v string) *GetIngressObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *GetIngressObjectMeta) SetManagedFields(v []*GetIngressManagedFieldsEntry) *GetIngressObjectMeta {
  s.ManagedFields = v
  return s
}

type GetIngressManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this GetIngressManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'GetIngressFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“GetIngressFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"GetIngressFieldsV1 holds the first JSON version format as described in the 'GetIngressFieldsV1' type", "zh_CN":"GetIngressFieldsV1 包含类型 “GetIngressFieldsV1” 中描述的第一个 JSON 版本格式"}
  GetIngressFieldsV1 *GetIngressFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s GetIngressManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s GetIngressManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *GetIngressManagedFieldsEntry) SetManager(v string) *GetIngressManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *GetIngressManagedFieldsEntry) SetOperation(v string) *GetIngressManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *GetIngressManagedFieldsEntry) SetApiVersion(v string) *GetIngressManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *GetIngressManagedFieldsEntry) SetTime(v string) *GetIngressManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *GetIngressManagedFieldsEntry) SetFieldsType(v string) *GetIngressManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *GetIngressManagedFieldsEntry) SetFieldsV1(v *GetIngressFieldsV1) *GetIngressManagedFieldsEntry {
  s.GetIngressFieldsV1 = v
  return s
}

func (s *GetIngressManagedFieldsEntry) SetSubresource(v string) *GetIngressManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type GetIngressFieldsV1 struct {
}

func (s GetIngressFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s GetIngressFieldsV1) GoString() string {
  return s.String()
}

type GetIngressOwnerReference struct {
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

func (s GetIngressOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s GetIngressOwnerReference) GoString() string {
  return s.String()
}

func (s *GetIngressOwnerReference) SetApiVersion(v string) *GetIngressOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *GetIngressOwnerReference) SetKind(v string) *GetIngressOwnerReference {
  s.Kind = &v
  return s
}

func (s *GetIngressOwnerReference) SetName(v string) *GetIngressOwnerReference {
  s.Name = &v
  return s
}

func (s *GetIngressOwnerReference) SetUid(v string) *GetIngressOwnerReference {
  s.Uid = &v
  return s
}

func (s *GetIngressOwnerReference) SetController(v bool) *GetIngressOwnerReference {
  s.Controller = &v
  return s
}

func (s *GetIngressOwnerReference) SetBlockOwnerDeletion(v bool) *GetIngressOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type DeleteIngressControllerRequest struct {
}

func (s DeleteIngressControllerRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressControllerRequest) GoString() string {
  return s.String()
}

type DeleteIngressControllerResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s DeleteIngressControllerResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressControllerResponse) GoString() string {
  return s.String()
}

func (s *DeleteIngressControllerResponse) SetCode(v int64) *DeleteIngressControllerResponse {
  s.Code = &v
  return s
}

func (s *DeleteIngressControllerResponse) SetMsg(v string) *DeleteIngressControllerResponse {
  s.Msg = &v
  return s
}

func (s *DeleteIngressControllerResponse) SetRequestId(v string) *DeleteIngressControllerResponse {
  s.RequestId = &v
  return s
}

type DeleteIngressControllerPaths struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteIngressControllerPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressControllerPaths) GoString() string {
  return s.String()
}

func (s *DeleteIngressControllerPaths) SetName(v string) *DeleteIngressControllerPaths {
  s.Name = &v
  return s
}

type DeleteIngressControllerParameters struct {
}

func (s DeleteIngressControllerParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressControllerParameters) GoString() string {
  return s.String()
}

type DeleteIngressControllerRequestHeader struct {
}

func (s DeleteIngressControllerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressControllerRequestHeader) GoString() string {
  return s.String()
}

type DeleteIngressControllerResponseHeader struct {
}

func (s DeleteIngressControllerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteIngressControllerResponseHeader) GoString() string {
  return s.String()
}




type CreateIngressControllerRequest struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"resource limit", "zh_CN":"资源限制"}
  Limit map[string]*string `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"resource request", "zh_CN":"所需资源"}
  Request map[string]*string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"cluster and replicate", "zh_CN":"部署集群和副本数"}
  Clusters []*CreateIngressControllerIngressCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
}

func (s CreateIngressControllerRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerRequest) GoString() string {
  return s.String()
}

func (s *CreateIngressControllerRequest) SetName(v string) *CreateIngressControllerRequest {
  s.Name = &v
  return s
}

func (s *CreateIngressControllerRequest) SetLimit(v map[string]*string) *CreateIngressControllerRequest {
  s.Limit = v
  return s
}

func (s *CreateIngressControllerRequest) SetRequest(v map[string]*string) *CreateIngressControllerRequest {
  s.Request = v
  return s
}

func (s *CreateIngressControllerRequest) SetClusters(v []*CreateIngressControllerIngressCluster) *CreateIngressControllerRequest {
  s.Clusters = v
  return s
}

type CreateIngressControllerResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress controller", "zh_CN":"路由控制器"}
  Data *CreateIngressControllerIngressController `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateIngressControllerResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerResponse) GoString() string {
  return s.String()
}

func (s *CreateIngressControllerResponse) SetCode(v int64) *CreateIngressControllerResponse {
  s.Code = &v
  return s
}

func (s *CreateIngressControllerResponse) SetMsg(v string) *CreateIngressControllerResponse {
  s.Msg = &v
  return s
}

func (s *CreateIngressControllerResponse) SetRequestId(v string) *CreateIngressControllerResponse {
  s.RequestId = &v
  return s
}

func (s *CreateIngressControllerResponse) SetData(v *CreateIngressControllerIngressController) *CreateIngressControllerResponse {
  s.Data = v
  return s
}

type CreateIngressControllerPaths struct {
}

func (s CreateIngressControllerPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerPaths) GoString() string {
  return s.String()
}

type CreateIngressControllerParameters struct {
}

func (s CreateIngressControllerParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerParameters) GoString() string {
  return s.String()
}

type CreateIngressControllerRequestHeader struct {
}

func (s CreateIngressControllerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerRequestHeader) GoString() string {
  return s.String()
}

type CreateIngressControllerResponseHeader struct {
}

func (s CreateIngressControllerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerResponseHeader) GoString() string {
  return s.String()
}

type CreateIngressControllerIngressController struct {
  // {"en":"ingress controller name", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"resource limit", "zh_CN":"资源限制"}
  Limit map[string]*string `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"resource request", "zh_CN":"所需资源"}
  Request map[string]*string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"cluster and replicate", "zh_CN":"部署集群和副本数"}
  Clusters []*CreateIngressControllerIngressCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
}

func (s CreateIngressControllerIngressController) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerIngressController) GoString() string {
  return s.String()
}

func (s *CreateIngressControllerIngressController) SetName(v string) *CreateIngressControllerIngressController {
  s.Name = &v
  return s
}

func (s *CreateIngressControllerIngressController) SetLimit(v map[string]*string) *CreateIngressControllerIngressController {
  s.Limit = v
  return s
}

func (s *CreateIngressControllerIngressController) SetRequest(v map[string]*string) *CreateIngressControllerIngressController {
  s.Request = v
  return s
}

func (s *CreateIngressControllerIngressController) SetClusters(v []*CreateIngressControllerIngressCluster) *CreateIngressControllerIngressController {
  s.Clusters = v
  return s
}

type CreateIngressControllerIngressCluster struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"replicates", "zh_CN":"副本数"}
  Replicate *int32 `json:"replicate,omitempty" xml:"replicate,omitempty" require:"true"`
}

func (s CreateIngressControllerIngressCluster) String() string {
  return tea.Prettify(s)
}

func (s CreateIngressControllerIngressCluster) GoString() string {
  return s.String()
}

func (s *CreateIngressControllerIngressCluster) SetName(v string) *CreateIngressControllerIngressCluster {
  s.Name = &v
  return s
}

func (s *CreateIngressControllerIngressCluster) SetReplicate(v int32) *CreateIngressControllerIngressCluster {
  s.Replicate = &v
  return s
}




type ListIngressControllerRequest struct {
}

func (s ListIngressControllerRequest) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerRequest) GoString() string {
  return s.String()
}

type ListIngressControllerResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ingress controller list", "zh_CN":"路由控制器列表"}
  Data *ListIngressControllerIngressControllerList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListIngressControllerResponse) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerResponse) GoString() string {
  return s.String()
}

func (s *ListIngressControllerResponse) SetCode(v int64) *ListIngressControllerResponse {
  s.Code = &v
  return s
}

func (s *ListIngressControllerResponse) SetMsg(v string) *ListIngressControllerResponse {
  s.Msg = &v
  return s
}

func (s *ListIngressControllerResponse) SetRequestId(v string) *ListIngressControllerResponse {
  s.RequestId = &v
  return s
}

func (s *ListIngressControllerResponse) SetData(v *ListIngressControllerIngressControllerList) *ListIngressControllerResponse {
  s.Data = v
  return s
}

type ListIngressControllerPaths struct {
}

func (s ListIngressControllerPaths) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerPaths) GoString() string {
  return s.String()
}

type ListIngressControllerParameters struct {
  // {"en":"keyword", "zh_CN":"关键字"}
  Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
  // {"en":"page index", "zh_CN":"页数"}
  PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"page size", "zh_CN":"每页数量"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListIngressControllerParameters) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerParameters) GoString() string {
  return s.String()
}

func (s *ListIngressControllerParameters) SetKeyword(v string) *ListIngressControllerParameters {
  s.Keyword = &v
  return s
}

func (s *ListIngressControllerParameters) SetPageIndex(v string) *ListIngressControllerParameters {
  s.PageIndex = &v
  return s
}

func (s *ListIngressControllerParameters) SetPageSize(v string) *ListIngressControllerParameters {
  s.PageSize = &v
  return s
}

type ListIngressControllerRequestHeader struct {
}

func (s ListIngressControllerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerRequestHeader) GoString() string {
  return s.String()
}

type ListIngressControllerResponseHeader struct {
}

func (s ListIngressControllerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerResponseHeader) GoString() string {
  return s.String()
}

type ListIngressControllerIngressControllerList struct {
  // {"en":"total  count", "zh_CN":"总数"}
  Total *int32 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  Items []*ListIngressControllerIngressControllerItem `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressControllerIngressControllerList) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerIngressControllerList) GoString() string {
  return s.String()
}

func (s *ListIngressControllerIngressControllerList) SetTotal(v int32) *ListIngressControllerIngressControllerList {
  s.Total = &v
  return s
}

func (s *ListIngressControllerIngressControllerList) SetItems(v []*ListIngressControllerIngressControllerItem) *ListIngressControllerIngressControllerList {
  s.Items = v
  return s
}

type ListIngressControllerIngressControllerItem struct {
  // {"en":"ingress controller name ", "zh_CN":"路由控制器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"ingress controller status", "zh_CN":"路由控制器状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"create time", "zh_CN":"创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"update time", "zh_CN":"更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  Clusters []*ListIngressControllerClusterAttr `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
}

func (s ListIngressControllerIngressControllerItem) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerIngressControllerItem) GoString() string {
  return s.String()
}

func (s *ListIngressControllerIngressControllerItem) SetName(v string) *ListIngressControllerIngressControllerItem {
  s.Name = &v
  return s
}

func (s *ListIngressControllerIngressControllerItem) SetStatus(v string) *ListIngressControllerIngressControllerItem {
  s.Status = &v
  return s
}

func (s *ListIngressControllerIngressControllerItem) SetCreateTime(v int64) *ListIngressControllerIngressControllerItem {
  s.CreateTime = &v
  return s
}

func (s *ListIngressControllerIngressControllerItem) SetUpdateTime(v int64) *ListIngressControllerIngressControllerItem {
  s.UpdateTime = &v
  return s
}

func (s *ListIngressControllerIngressControllerItem) SetClusters(v []*ListIngressControllerClusterAttr) *ListIngressControllerIngressControllerItem {
  s.Clusters = v
  return s
}

type ListIngressControllerClusterAttr struct {
  // {"en":"cluster name ", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty" require:"true"`
  // {"en":"cluster cn name", "zh_CN":"集群中文名"}
  ClusterNameCn *string `json:"clusterNameCn,omitempty" xml:"clusterNameCn,omitempty" require:"true"`
  // {"en":"service ip ", "zh_CN":"服务ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Total number of non-terminated pods targeted by this deployment (their labels match the selector).", "zh_CN":"期望副本数"}
  Replicas *int32 `json:"replicas,omitempty" xml:"replicas,omitempty" require:"true"`
  // {"en":"Total number of ready pods targeted by this deployment", "zh_CN":"就绪副本数"}
  ReadyReplicas *int32 `json:"readyReplicas,omitempty" xml:"readyReplicas,omitempty" require:"true"`
  // {"en":"Total number of available pods (ready for at least minReadySeconds) targeted by this deployment", "zh_CN":"一段时间内(minReadySeconds)就绪副本数"}
  AvailableReplicas *int32 `json:"availableReplicas,omitempty" xml:"availableReplicas,omitempty" require:"true"`
  // {"en":"Total number of non-terminated pods targeted by this deployment that have the desired template spec.", "zh_CN":"可访问副本数"}
  UpdatedReplicas *string `json:"updatedReplicas,omitempty" xml:"updatedReplicas,omitempty" require:"true"`
}

func (s ListIngressControllerClusterAttr) String() string {
  return tea.Prettify(s)
}

func (s ListIngressControllerClusterAttr) GoString() string {
  return s.String()
}

func (s *ListIngressControllerClusterAttr) SetClusterName(v string) *ListIngressControllerClusterAttr {
  s.ClusterName = &v
  return s
}

func (s *ListIngressControllerClusterAttr) SetClusterNameCn(v string) *ListIngressControllerClusterAttr {
  s.ClusterNameCn = &v
  return s
}

func (s *ListIngressControllerClusterAttr) SetIp(v string) *ListIngressControllerClusterAttr {
  s.Ip = &v
  return s
}

func (s *ListIngressControllerClusterAttr) SetReplicas(v int32) *ListIngressControllerClusterAttr {
  s.Replicas = &v
  return s
}

func (s *ListIngressControllerClusterAttr) SetReadyReplicas(v int32) *ListIngressControllerClusterAttr {
  s.ReadyReplicas = &v
  return s
}

func (s *ListIngressControllerClusterAttr) SetAvailableReplicas(v int32) *ListIngressControllerClusterAttr {
  s.AvailableReplicas = &v
  return s
}

func (s *ListIngressControllerClusterAttr) SetUpdatedReplicas(v string) *ListIngressControllerClusterAttr {
  s.UpdatedReplicas = &v
  return s
}




type VmpEdgeIpAllocate4OccupancyRequest struct {
  // {"en":"Instance IP to be bound","zh_CN":"额外IP要绑定的实例IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"IP protocol: 4-ipv4(default); 6-ipv6","zh_CN":"IP协议：4-ipv4（默认）；6-ipv6"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"IP native attribute, 1: non-native;-1: native;","zh_CN":"指定IPv4原生属性。可选值：1：非原生，-1：原生。不指定默认随机分配原生属性"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"CIDR","zh_CN":"指定CIDR申请IP"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"A. When an instance has multiple carrier IPs, you can specify an additional IP carrier. B. If not specified: For instances with a single carrier IP, the additional IP will use the same carrier as the instance. For instances with cross-ISP IPs, For the additional IP, one of the carrier BGP nodes will be selected; this field will not take effect.","zh_CN":"A、实例有多个运营商IP时，可指定额外IP运营商\nB、不指定时：\n实例单运营商IP，额外IP运营商一致\n实例多运营商IP，额外IP选其中一个运营商\nBGP节点该字段不生效。"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Allocate ip random","zh_CN":"IPv4是否随机分配\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
  // {"en":"Number of applications IP","zh_CN":"申请个数"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s VmpEdgeIpAllocate4OccupancyRequest) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyRequest) GoString() string {
  return s.String()
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetServerIp(v string) *VmpEdgeIpAllocate4OccupancyRequest {
  s.ServerIp = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetProtocol(v string) *VmpEdgeIpAllocate4OccupancyRequest {
  s.Protocol = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetNativeAttribute(v string) *VmpEdgeIpAllocate4OccupancyRequest {
  s.NativeAttribute = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetCidr(v string) *VmpEdgeIpAllocate4OccupancyRequest {
  s.Cidr = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetCarrier(v string) *VmpEdgeIpAllocate4OccupancyRequest {
  s.Carrier = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetRandomAllocateIp(v int) *VmpEdgeIpAllocate4OccupancyRequest {
  s.RandomAllocateIp = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyRequest) SetCount(v int) *VmpEdgeIpAllocate4OccupancyRequest {
  s.Count = &v
  return s
}

type VmpEdgeIpAllocate4OccupancyRequestHeader struct {
}

func (s VmpEdgeIpAllocate4OccupancyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyRequestHeader) GoString() string {
  return s.String()
}

type VmpEdgeIpAllocate4OccupancyPaths struct {
}

func (s VmpEdgeIpAllocate4OccupancyPaths) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyPaths) GoString() string {
  return s.String()
}

type VmpEdgeIpAllocate4OccupancyParameters struct {
}

func (s VmpEdgeIpAllocate4OccupancyParameters) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyParameters) GoString() string {
  return s.String()
}

type VmpEdgeIpAllocate4OccupancyResponse struct {
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *VmpEdgeIpAllocate4OccupancyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Return Code","zh_CN":"返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s VmpEdgeIpAllocate4OccupancyResponse) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyResponse) GoString() string {
  return s.String()
}

func (s *VmpEdgeIpAllocate4OccupancyResponse) SetMessage(v string) *VmpEdgeIpAllocate4OccupancyResponse {
  s.Message = &v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyResponse) SetData(v *VmpEdgeIpAllocate4OccupancyResponseData) *VmpEdgeIpAllocate4OccupancyResponse {
  s.Data = v
  return s
}

func (s *VmpEdgeIpAllocate4OccupancyResponse) SetCode(v string) *VmpEdgeIpAllocate4OccupancyResponse {
  s.Code = &v
  return s
}

type VmpEdgeIpAllocate4OccupancyResponseData struct {
  // {"en":"Successful application for all or part of IP","zh_CN":"成功申请到的全部或部分IP说明：不同场景的响应说明如下 A、所有IP都申请成功，返回申请到的所有IP B、只申请到部分IP，返回申请到的那部分IP C、未申请到任何IP，返回失败信息 D、若出现申请失败的情况，请间隔10S之后再次申请"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VmpEdgeIpAllocate4OccupancyResponseData) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyResponseData) GoString() string {
  return s.String()
}

func (s *VmpEdgeIpAllocate4OccupancyResponseData) SetEdgeIps(v []*string) *VmpEdgeIpAllocate4OccupancyResponseData {
  s.EdgeIps = v
  return s
}

type VmpEdgeIpAllocate4OccupancyResponseHeader struct {
}

func (s VmpEdgeIpAllocate4OccupancyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VmpEdgeIpAllocate4OccupancyResponseHeader) GoString() string {
  return s.String()
}




