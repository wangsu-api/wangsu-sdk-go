package templatemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VMPQueryNodeServersRequest struct {
}

func (s VMPQueryNodeServersRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersRequest) GoString() string {
  return s.String()
}

type VMPQueryNodeServersRequestHeader struct {
}

func (s VMPQueryNodeServersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryNodeServersPaths struct {
}

func (s VMPQueryNodeServersPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersPaths) GoString() string {
  return s.String()
}

type VMPQueryNodeServersParameters struct {
  // {"en":"Province","zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Carrier","zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node Name","zh_CN":"节点名称，表示指定节点创建实例"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"FlavorId","zh_CN":"实例规格标识（可通过实例规格查询接口进行查询：https://www.wangsu.com/document/api-doc/22811?rsr=ws）"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Protocols:\n4:only ipv4\n6:only ipv6\n0:both ipv4 & ipv6","zh_CN":"是否需要多ip协议地址\n4：只需要ipv4地址，默认值\n6：只需要ipv6地址\n0：ipv4、ipv6都需要"}
  Protocols *int `json:"protocols,omitempty" xml:"protocols,omitempty"`
  // {"en":"Virtual machine network type:\nANY - the default;\nDPDK - DPDK network.","zh_CN":"虚拟机网络类型：\nANY-默认；\nDPDK-DPDK网络。"}
  NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
  // {"en":"Use unique IP segment","zh_CN":"是否使用唯一网段\n1：是\n-1：否"}
  UseUniqueIpSegment *int `json:"useUniqueIpSegment,omitempty" xml:"useUniqueIpSegment,omitempty"`
  // {"en":"Dual-line, multi-line, BGP cluster specifies part of the isp, separated by commas. Example: dx,wt","zh_CN":"双线,多线,bgp集群指定部分isp,逗号分割.例子: dx,wt"}
  MultiCarriers *string `json:"multiCarriers,omitempty" xml:"multiCarriers,omitempty"`
}

func (s VMPQueryNodeServersParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeServersParameters) SetProvince(v string) *VMPQueryNodeServersParameters {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetCarrier(v string) *VMPQueryNodeServersParameters {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetNodeName(v string) *VMPQueryNodeServersParameters {
  s.NodeName = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetFlavorId(v string) *VMPQueryNodeServersParameters {
  s.FlavorId = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetProtocols(v int) *VMPQueryNodeServersParameters {
  s.Protocols = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetNetworkType(v string) *VMPQueryNodeServersParameters {
  s.NetworkType = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetUseUniqueIpSegment(v int) *VMPQueryNodeServersParameters {
  s.UseUniqueIpSegment = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetMultiCarriers(v string) *VMPQueryNodeServersParameters {
  s.MultiCarriers = &v
  return s
}

type VMPQueryNodeServersResponse struct {
  // {"en":"Instance quantity information","zh_CN":"实例数量信息列表"}
  Servers []*VMPQueryNodeServersResponseServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryNodeServersResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeServersResponse) SetServers(v []*VMPQueryNodeServersResponseServers) *VMPQueryNodeServersResponse {
  s.Servers = v
  return s
}

type VMPQueryNodeServersResponseServers struct     {
  // {"en":"Province","zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Carrier","zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Node Name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Number of instances that can be created","zh_CN":"可创建的实例的数量"}
  Count *string `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s VMPQueryNodeServersResponseServers) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersResponseServers) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeServersResponseServers) SetProvince(v string) *VMPQueryNodeServersResponseServers {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeServersResponseServers) SetCarrier(v string) *VMPQueryNodeServersResponseServers {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeServersResponseServers) SetNodeName(v string) *VMPQueryNodeServersResponseServers {
  s.NodeName = &v
  return s
}

func (s *VMPQueryNodeServersResponseServers) SetCount(v string) *VMPQueryNodeServersResponseServers {
  s.Count = &v
  return s
}

type VMPQueryNodeServersResponseHeader struct {
}

func (s VMPQueryNodeServersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryNodeServersRequest struct {
}

func (s LECHQueryNodeServersRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersRequest) GoString() string {
  return s.String()
}

type LECHQueryNodeServersRequestHeader struct {
}

func (s LECHQueryNodeServersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryNodeServersPaths struct {
}

func (s LECHQueryNodeServersPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersPaths) GoString() string {
  return s.String()
}

type LECHQueryNodeServersParameters struct {
  // {"en":"Province","zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Carrier","zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node Name","zh_CN":"节点名称，表示指定节点创建实例"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"FlavorId","zh_CN":"实例规格标识（可通过实例规格查询接口进行查询：https://www.wangsu.com/document/api-doc/22811?rsr=ws）"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Protocols:\n4:only ipv4\n6:only ipv6\n0:both ipv4 & ipv6","zh_CN":"是否需要多ip协议地址\n4：只需要ipv4地址，默认值\n6：只需要ipv6地址\n0：ipv4、ipv6都需要"}
  Protocols *int `json:"protocols,omitempty" xml:"protocols,omitempty"`
  // {"en":"Virtual machine network type:\nANY - the default;\nDPDK - DPDK network.","zh_CN":"虚拟机网络类型：\nANY-默认；\nDPDK-DPDK网络。"}
  NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
  // {"en":"Use unique IP segment","zh_CN":"是否使用唯一网段\n1：是\n-1：否"}
  UseUniqueIpSegment *int `json:"useUniqueIpSegment,omitempty" xml:"useUniqueIpSegment,omitempty"`
  // {"en":"Dual-line, multi-line, BGP cluster specifies part of the isp, separated by commas. Example: dx,wt","zh_CN":"双线,多线,bgp集群指定部分isp,逗号分割.例子: dx,wt"}
  MultiCarriers *string `json:"multiCarriers,omitempty" xml:"multiCarriers,omitempty"`
}

func (s LECHQueryNodeServersParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryNodeServersParameters) SetProvince(v string) *LECHQueryNodeServersParameters {
  s.Province = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetCarrier(v string) *LECHQueryNodeServersParameters {
  s.Carrier = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetNodeName(v string) *LECHQueryNodeServersParameters {
  s.NodeName = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetFlavorId(v string) *LECHQueryNodeServersParameters {
  s.FlavorId = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetProtocols(v int) *LECHQueryNodeServersParameters {
  s.Protocols = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetNetworkType(v string) *LECHQueryNodeServersParameters {
  s.NetworkType = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetUseUniqueIpSegment(v int) *LECHQueryNodeServersParameters {
  s.UseUniqueIpSegment = &v
  return s
}

func (s *LECHQueryNodeServersParameters) SetMultiCarriers(v string) *LECHQueryNodeServersParameters {
  s.MultiCarriers = &v
  return s
}

type LECHQueryNodeServersResponse struct {
  // {"en":"Instance quantity information","zh_CN":"实例数量信息列表"}
  Servers []*LECHQueryNodeServersResponseServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryNodeServersResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryNodeServersResponse) SetServers(v []*LECHQueryNodeServersResponseServers) *LECHQueryNodeServersResponse {
  s.Servers = v
  return s
}

type LECHQueryNodeServersResponseServers struct     {
  // {"en":"Province","zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Carrier","zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Node Name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Number of instances that can be created","zh_CN":"可创建的实例的数量"}
  Count *string `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s LECHQueryNodeServersResponseServers) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersResponseServers) GoString() string {
  return s.String()
}

func (s *LECHQueryNodeServersResponseServers) SetProvince(v string) *LECHQueryNodeServersResponseServers {
  s.Province = &v
  return s
}

func (s *LECHQueryNodeServersResponseServers) SetCarrier(v string) *LECHQueryNodeServersResponseServers {
  s.Carrier = &v
  return s
}

func (s *LECHQueryNodeServersResponseServers) SetNodeName(v string) *LECHQueryNodeServersResponseServers {
  s.NodeName = &v
  return s
}

func (s *LECHQueryNodeServersResponseServers) SetCount(v string) *LECHQueryNodeServersResponseServers {
  s.Count = &v
  return s
}

type LECHQueryNodeServersResponseHeader struct {
}

func (s LECHQueryNodeServersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeServersResponseHeader) GoString() string {
  return s.String()
}




