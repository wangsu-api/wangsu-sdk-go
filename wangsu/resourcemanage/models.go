package resourcemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VMPQueryFlavorRequest struct {
}

func (s VMPQueryFlavorRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorRequest) GoString() string {
  return s.String()
}

type VMPQueryFlavorRequestHeader struct {
}

func (s VMPQueryFlavorRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryFlavorPaths struct {
}

func (s VMPQueryFlavorPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorPaths) GoString() string {
  return s.String()
}

type VMPQueryFlavorParameters struct {
  // {"en":"The virtual machine specification is a unique identifier. Multiple values are separated by commas.Can be left blank, and when left blank, all available template specifications will be returned.","zh_CN":"实例规格唯一标识，多个值用英文逗号分隔。可放空不填，不填时返回所有可用模板规格。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
}

func (s VMPQueryFlavorParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryFlavorParameters) SetIds(v string) *VMPQueryFlavorParameters {
  s.Ids = &v
  return s
}

type VMPQueryFlavorResponse struct {
  // {"en":"flavors","zh_CN":"规格"}
  Flavors []*string `json:"flavors,omitempty" xml:"flavors,omitempty" require:"true" type:"Repeated"`
  // {"en":"Unique identification of virtual machine specification, global unique","zh_CN":"实例规格唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"form name","zh_CN":"规格名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Number of CPUs of virtual machine","zh_CN":"实例的cpu数"}
  Vcpus *int `json:"vcpus,omitempty" xml:"vcpus,omitempty" require:"true"`
  // {"en":"Virtual machine memory in GB","zh_CN":"实例内存,单位是GB"}
  Ram *int `json:"ram,omitempty" xml:"ram,omitempty" require:"true"`
  // {"en":"Disk information of virtual machine","zh_CN":"实例的磁盘信息"}
  Disks []*VMPQueryFlavorResponseDisks `json:"disks,omitempty" xml:"disks,omitempty" require:"true" type:"Repeated"`
  // {"en":"Bearable bandwidth, Mbps","zh_CN":"可承载带宽，单位是Mbps"}
  Bandwidth *int `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"1: Yes, -1: No,1 means the template is bare metal template;,-1 indicates that the template is a cloud host template;","zh_CN":"1：是，-1：否,1表示该模板是裸机模板；-1表示该模板是云主机模板；"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"SSD system disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template","zh_CN":"SSD系统盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  SysSsdLimit *int `json:"sysSsdLimit,omitempty" xml:"sysSsdLimit,omitempty" require:"true"`
  // {"en":"HDD system disk quota (GB). If it is a bare-metal template or a stand-alone template, this parameter has no meaning","zh_CN":"HDD系统盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  SysHddLimit *int `json:"sysHddLimit,omitempty" xml:"sysHddLimit,omitempty" require:"true"`
  // {"en":"SSD disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template","zh_CN":"SSD数据盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  DataSsdLimit *int `json:"dataSsdLimit,omitempty" xml:"dataSsdLimit,omitempty" require:"true"`
  // {"en":"HDD disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template","zh_CN":"HDD数据盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  DataHddLimit *int `json:"dataHddLimit,omitempty" xml:"dataHddLimit,omitempty" require:"true"`
  // {"en":"Template type,Values: 201- public template, 202- custom template","zh_CN":"模板类型,取值：201-公共模板、202-自定义模板"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Virtual Machine Area (see Appendix for details)","zh_CN":"【仅为预留字段，实际未使用】实例所属区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
}

func (s VMPQueryFlavorResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryFlavorResponse) SetFlavors(v []*string) *VMPQueryFlavorResponse {
  s.Flavors = v
  return s
}

func (s *VMPQueryFlavorResponse) SetId(v string) *VMPQueryFlavorResponse {
  s.Id = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetName(v string) *VMPQueryFlavorResponse {
  s.Name = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetVcpus(v int) *VMPQueryFlavorResponse {
  s.Vcpus = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetRam(v int) *VMPQueryFlavorResponse {
  s.Ram = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetDisks(v []*VMPQueryFlavorResponseDisks) *VMPQueryFlavorResponse {
  s.Disks = v
  return s
}

func (s *VMPQueryFlavorResponse) SetBandwidth(v int) *VMPQueryFlavorResponse {
  s.Bandwidth = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetIsBm(v int) *VMPQueryFlavorResponse {
  s.IsBm = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetSysSsdLimit(v int) *VMPQueryFlavorResponse {
  s.SysSsdLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetSysHddLimit(v int) *VMPQueryFlavorResponse {
  s.SysHddLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetDataSsdLimit(v int) *VMPQueryFlavorResponse {
  s.DataSsdLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetDataHddLimit(v int) *VMPQueryFlavorResponse {
  s.DataHddLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetType(v string) *VMPQueryFlavorResponse {
  s.Type = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetRegionName(v string) *VMPQueryFlavorResponse {
  s.RegionName = &v
  return s
}

type VMPQueryFlavorResponseDisks struct     {
  // {"en":"disk type ,system disk or data disk","zh_CN":"磁盘类型,数据盘或者系统盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk space size in GB","zh_CN":"磁盘空间大小，单位是GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk type, value:HDD: ordinary hard disk,SSD: solid state drive,The default is HDD","zh_CN":"磁盘类型，取值：HDD：普通硬盘,SSD：固态硬盘,默认是HDD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s VMPQueryFlavorResponseDisks) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorResponseDisks) GoString() string {
  return s.String()
}

func (s *VMPQueryFlavorResponseDisks) SetType(v string) *VMPQueryFlavorResponseDisks {
  s.Type = &v
  return s
}

func (s *VMPQueryFlavorResponseDisks) SetSize(v int) *VMPQueryFlavorResponseDisks {
  s.Size = &v
  return s
}

func (s *VMPQueryFlavorResponseDisks) SetCategory(v string) *VMPQueryFlavorResponseDisks {
  s.Category = &v
  return s
}

type VMPQueryFlavorResponseHeader struct {
}

func (s VMPQueryFlavorResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryNodeRequest struct {
}

func (s LECHQueryNodeRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeRequest) GoString() string {
  return s.String()
}

type LECHQueryNodeRequestHeader struct {
}

func (s LECHQueryNodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryNodePaths struct {
}

func (s LECHQueryNodePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodePaths) GoString() string {
  return s.String()
}

type LECHQueryNodeParameters struct {
  // {"en":"The sorted field name can have multiple values: name, regionname, province","zh_CN":"排序的字段名称，可以有多个，取值：name、regionName、province"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"Sorting direction must follow sortkey, value: desc: descending, default value: ASC: ascending","zh_CN":"排序方向，必须跟在sortKey后面出现，取值：desc：降序，默认值 asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default","zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the name specified by marker","zh_CN":"从marker指定的名称开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Node area (see Appendix for details)","zh_CN":"节点所属区域（区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Node province (see Appendix for details)","zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Node carrier (see Appendix for details)","zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Line type: single -- single line node; double -- double line node; triple -- three line node; BGP -- BGP node","zh_CN":"线路类型：single -- 单线节点；double -- 双线节点；triple -- 三线节点；bgp -- BGP节点"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty"`
  // {"en":"IPv6 supported: true: IPv6 supported false: IPv6 not supported","zh_CN":"是否支持ipv6：True：支持ipv6 False：不支持ipv6"}
  Ipv6Supported *string `json:"ipv6Supported,omitempty" xml:"ipv6Supported,omitempty"`
  // {"en":"Whether the node has bare metal resources\nTrue: There are bare metal resources\nFalse: No bare metal resources, only virtual machine resources","zh_CN":"该节点是否有裸机资源\nTrue：有裸机资源\nFalse：没有裸机资源，只有虚拟机资源"}
  BmSupported *string `json:"bmSupported,omitempty" xml:"bmSupported,omitempty"`
}

func (s LECHQueryNodeParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryNodeParameters) SetSortKey(v string) *LECHQueryNodeParameters {
  s.SortKey = &v
  return s
}

func (s *LECHQueryNodeParameters) SetSortDir(v string) *LECHQueryNodeParameters {
  s.SortDir = &v
  return s
}

func (s *LECHQueryNodeParameters) SetLimit(v int) *LECHQueryNodeParameters {
  s.Limit = &v
  return s
}

func (s *LECHQueryNodeParameters) SetMarker(v string) *LECHQueryNodeParameters {
  s.Marker = &v
  return s
}

func (s *LECHQueryNodeParameters) SetRegionName(v string) *LECHQueryNodeParameters {
  s.RegionName = &v
  return s
}

func (s *LECHQueryNodeParameters) SetProvince(v string) *LECHQueryNodeParameters {
  s.Province = &v
  return s
}

func (s *LECHQueryNodeParameters) SetCarrier(v string) *LECHQueryNodeParameters {
  s.Carrier = &v
  return s
}

func (s *LECHQueryNodeParameters) SetLineType(v string) *LECHQueryNodeParameters {
  s.LineType = &v
  return s
}

func (s *LECHQueryNodeParameters) SetIpv6Supported(v string) *LECHQueryNodeParameters {
  s.Ipv6Supported = &v
  return s
}

func (s *LECHQueryNodeParameters) SetBmSupported(v string) *LECHQueryNodeParameters {
  s.BmSupported = &v
  return s
}

type LECHQueryNodeResponse struct {
  // {"en":"Node information array","zh_CN":"节点信息数组"}
  Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Node name, unique","zh_CN":"节点名称，唯一"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Node area","zh_CN":"节点所在区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"Province of node","zh_CN":"节点所在省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"If the node is a multi line node, multiple operators will be returned, separated by '/'","zh_CN":"节点所在运营商，如果是多线节点，则返回多个运营商，以'/'分隔"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Node status: running - node available; maintenance - node in maintenance, temporarily unavailable","zh_CN":"节点状态：RUNNING ---节点可用；MAINTENANCE ---节点维护中，暂时不可用"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Line typenode","zh_CN":"线路类型"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty" require:"true"`
  // {"en":"IPv6 supported","zh_CN":"是否支持ipv6"}
  Ipv6Supported *string `json:"ipv6Supported,omitempty" xml:"ipv6Supported,omitempty" require:"true"`
  // {"en":"Whether the node has bare metal resources","zh_CN":"该节点是否有裸机资源"}
  BmSupported *string `json:"bmSupported,omitempty" xml:"bmSupported,omitempty" require:"true"`
}

func (s LECHQueryNodeResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryNodeResponse) SetNodes(v []*string) *LECHQueryNodeResponse {
  s.Nodes = v
  return s
}

func (s *LECHQueryNodeResponse) SetName(v string) *LECHQueryNodeResponse {
  s.Name = &v
  return s
}

func (s *LECHQueryNodeResponse) SetRegionName(v string) *LECHQueryNodeResponse {
  s.RegionName = &v
  return s
}

func (s *LECHQueryNodeResponse) SetProvince(v string) *LECHQueryNodeResponse {
  s.Province = &v
  return s
}

func (s *LECHQueryNodeResponse) SetCarrier(v string) *LECHQueryNodeResponse {
  s.Carrier = &v
  return s
}

func (s *LECHQueryNodeResponse) SetState(v string) *LECHQueryNodeResponse {
  s.State = &v
  return s
}

func (s *LECHQueryNodeResponse) SetLineType(v string) *LECHQueryNodeResponse {
  s.LineType = &v
  return s
}

func (s *LECHQueryNodeResponse) SetIpv6Supported(v string) *LECHQueryNodeResponse {
  s.Ipv6Supported = &v
  return s
}

func (s *LECHQueryNodeResponse) SetBmSupported(v string) *LECHQueryNodeResponse {
  s.BmSupported = &v
  return s
}

type LECHQueryNodeResponseHeader struct {
}

func (s LECHQueryNodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryNodeResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryFlavorRequest struct {
}

func (s LECHQueryFlavorRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorRequest) GoString() string {
  return s.String()
}

type LECHQueryFlavorRequestHeader struct {
}

func (s LECHQueryFlavorRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryFlavorPaths struct {
}

func (s LECHQueryFlavorPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorPaths) GoString() string {
  return s.String()
}

type LECHQueryFlavorParameters struct {
  // {"en":"The virtual machine specification is a unique identifier. Multiple values are separated by commas.Can be left blank, and when left blank, all available template specifications will be returned.","zh_CN":"实例规格唯一标识，多个值用英文逗号分隔。可放空不填，不填时返回所有可用模板规格。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
}

func (s LECHQueryFlavorParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryFlavorParameters) SetIds(v string) *LECHQueryFlavorParameters {
  s.Ids = &v
  return s
}

type LECHQueryFlavorResponse struct {
  // {"en":"flavors","zh_CN":"规格"}
  Flavors []*string `json:"flavors,omitempty" xml:"flavors,omitempty" require:"true" type:"Repeated"`
  // {"en":"Unique identification of virtual machine specification, global unique","zh_CN":"实例规格唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"form name","zh_CN":"规格名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Number of CPUs of virtual machine","zh_CN":"实例的cpu数"}
  Vcpus *int `json:"vcpus,omitempty" xml:"vcpus,omitempty" require:"true"`
  // {"en":"Virtual machine memory in GB","zh_CN":"实例内存,单位是GB"}
  Ram *int `json:"ram,omitempty" xml:"ram,omitempty" require:"true"`
  // {"en":"Disk information of virtual machine","zh_CN":"实例的磁盘信息"}
  Disks []*LECHQueryFlavorResponseDisks `json:"disks,omitempty" xml:"disks,omitempty" require:"true" type:"Repeated"`
  // {"en":"Bearable bandwidth, Mbps","zh_CN":"可承载带宽，单位是Mbps"}
  Bandwidth *int `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"1: Yes, -1: No,1 means the template is bare metal template;,-1 indicates that the template is a cloud host template;","zh_CN":"1：是，-1：否,1表示该模板是裸机模板；-1表示该模板是云主机模板；"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"SSD system disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template","zh_CN":"SSD系统盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  SysSsdLimit *int `json:"sysSsdLimit,omitempty" xml:"sysSsdLimit,omitempty" require:"true"`
  // {"en":"HDD system disk quota (GB). If it is a bare-metal template or a stand-alone template, this parameter has no meaning","zh_CN":"HDD系统盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  SysHddLimit *int `json:"sysHddLimit,omitempty" xml:"sysHddLimit,omitempty" require:"true"`
  // {"en":"SSD disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template","zh_CN":"SSD数据盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  DataSsdLimit *int `json:"dataSsdLimit,omitempty" xml:"dataSsdLimit,omitempty" require:"true"`
  // {"en":"HDD disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template","zh_CN":"HDD数据盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  DataHddLimit *int `json:"dataHddLimit,omitempty" xml:"dataHddLimit,omitempty" require:"true"`
  // {"en":"Template type,Values: 201- public template, 202- custom template","zh_CN":"模板类型,取值：201-公共模板、202-自定义模板"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Virtual Machine Area (see Appendix for details)","zh_CN":"【仅为预留字段，实际未使用】实例所属区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
}

func (s LECHQueryFlavorResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryFlavorResponse) SetFlavors(v []*string) *LECHQueryFlavorResponse {
  s.Flavors = v
  return s
}

func (s *LECHQueryFlavorResponse) SetId(v string) *LECHQueryFlavorResponse {
  s.Id = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetName(v string) *LECHQueryFlavorResponse {
  s.Name = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetVcpus(v int) *LECHQueryFlavorResponse {
  s.Vcpus = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetRam(v int) *LECHQueryFlavorResponse {
  s.Ram = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetDisks(v []*LECHQueryFlavorResponseDisks) *LECHQueryFlavorResponse {
  s.Disks = v
  return s
}

func (s *LECHQueryFlavorResponse) SetBandwidth(v int) *LECHQueryFlavorResponse {
  s.Bandwidth = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetIsBm(v int) *LECHQueryFlavorResponse {
  s.IsBm = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetSysSsdLimit(v int) *LECHQueryFlavorResponse {
  s.SysSsdLimit = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetSysHddLimit(v int) *LECHQueryFlavorResponse {
  s.SysHddLimit = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetDataSsdLimit(v int) *LECHQueryFlavorResponse {
  s.DataSsdLimit = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetDataHddLimit(v int) *LECHQueryFlavorResponse {
  s.DataHddLimit = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetType(v string) *LECHQueryFlavorResponse {
  s.Type = &v
  return s
}

func (s *LECHQueryFlavorResponse) SetRegionName(v string) *LECHQueryFlavorResponse {
  s.RegionName = &v
  return s
}

type LECHQueryFlavorResponseDisks struct     {
  // {"en":"disk type ,system disk or data disk","zh_CN":"磁盘类型,数据盘或者系统盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk space size in GB","zh_CN":"磁盘空间大小，单位是GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk type, value:HDD: ordinary hard disk,SSD: solid state drive,The default is HDD","zh_CN":"磁盘类型，取值：HDD：普通硬盘,SSD：固态硬盘,默认是HDD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s LECHQueryFlavorResponseDisks) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorResponseDisks) GoString() string {
  return s.String()
}

func (s *LECHQueryFlavorResponseDisks) SetType(v string) *LECHQueryFlavorResponseDisks {
  s.Type = &v
  return s
}

func (s *LECHQueryFlavorResponseDisks) SetSize(v int) *LECHQueryFlavorResponseDisks {
  s.Size = &v
  return s
}

func (s *LECHQueryFlavorResponseDisks) SetCategory(v string) *LECHQueryFlavorResponseDisks {
  s.Category = &v
  return s
}

type LECHQueryFlavorResponseHeader struct {
}

func (s LECHQueryFlavorResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryFlavorResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryNodeRequest struct {
}

func (s VMPQueryNodeRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeRequest) GoString() string {
  return s.String()
}

type VMPQueryNodeResponse struct {
  // {"en":"Node information array", "zh_CN":"节点信息数组"}
  Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Node name, unique", "zh_CN":"节点名称，唯一"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Node area", "zh_CN":"节点所在区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"Province of node", "zh_CN":"节点所在省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"If the node is a multi line node, multiple operators will be returned, separated by '/'", "zh_CN":"节点所在运营商，如果是多线节点，则返回多个运营商，以'/'分隔"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Node status: running - node available; maintenance - node in maintenance, temporarily unavailable", "zh_CN":"节点状态：RUNNING ---节点可用；MAINTENANCE ---节点维护中，暂时不可用"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Line typenode", "zh_CN":"线路类型"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty" require:"true"`
  // {"en":"IPv6 supported", "zh_CN":"是否支持ipv6"}
  Ipv6Supported *string `json:"ipv6Supported,omitempty" xml:"ipv6Supported,omitempty" require:"true"`
  // {"en":"Whether the node has bare metal resources", "zh_CN":"该节点是否有裸机资源"}
  BmSupported *string `json:"bmSupported,omitempty" xml:"bmSupported,omitempty" require:"true"`
}

func (s VMPQueryNodeResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeResponse) SetNodes(v []*string) *VMPQueryNodeResponse {
  s.Nodes = v
  return s
}

func (s *VMPQueryNodeResponse) SetName(v string) *VMPQueryNodeResponse {
  s.Name = &v
  return s
}

func (s *VMPQueryNodeResponse) SetRegionName(v string) *VMPQueryNodeResponse {
  s.RegionName = &v
  return s
}

func (s *VMPQueryNodeResponse) SetProvince(v string) *VMPQueryNodeResponse {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeResponse) SetCarrier(v string) *VMPQueryNodeResponse {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeResponse) SetState(v string) *VMPQueryNodeResponse {
  s.State = &v
  return s
}

func (s *VMPQueryNodeResponse) SetLineType(v string) *VMPQueryNodeResponse {
  s.LineType = &v
  return s
}

func (s *VMPQueryNodeResponse) SetIpv6Supported(v string) *VMPQueryNodeResponse {
  s.Ipv6Supported = &v
  return s
}

func (s *VMPQueryNodeResponse) SetBmSupported(v string) *VMPQueryNodeResponse {
  s.BmSupported = &v
  return s
}

type VMPQueryNodePaths struct {
}

func (s VMPQueryNodePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodePaths) GoString() string {
  return s.String()
}

type VMPQueryNodeParameters struct {
  // {"en":"The sorted field name can have multiple values: name, regionname, province", "zh_CN":"排序的字段名称，可以有多个，取值：name、regionName、province"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"Sorting direction must follow sortkey, value: desc: descending, default value: ASC: ascending", "zh_CN":"排序方向，必须跟在sortKey后面出现，取值：desc：降序，默认值 asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the name specified by marker", "zh_CN":"从marker指定的名称开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Node area (see Appendix for details)", "zh_CN":"节点所属区域（区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Node province (see Appendix for details)", "zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Node carrier (see Appendix for details)", "zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Line type: single -- single line node; double -- double line node; triple -- three line node; BGP -- BGP node", "zh_CN":"线路类型：single -- 单线节点；double -- 双线节点；triple -- 三线节点；bgp -- BGP节点"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty"`
  // {"en":"IPv6 supported: true: IPv6 supported false: IPv6 not supported", "zh_CN":"是否支持ipv6：True：支持ipv6 False：不支持ipv6"}
  Ipv6Supported *string `json:"ipv6Supported,omitempty" xml:"ipv6Supported,omitempty"`
  // {"en":"Whether the node has bare metal resources
  // True: There are bare metal resources
  // False: No bare metal resources, only virtual machine resources", "zh_CN":"该节点是否有裸机资源
  // True：有裸机资源
  // False：没有裸机资源，只有虚拟机资源"}
  BmSupported *string `json:"bmSupported,omitempty" xml:"bmSupported,omitempty"`
}

func (s VMPQueryNodeParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeParameters) SetSortKey(v string) *VMPQueryNodeParameters {
  s.SortKey = &v
  return s
}

func (s *VMPQueryNodeParameters) SetSortDir(v string) *VMPQueryNodeParameters {
  s.SortDir = &v
  return s
}

func (s *VMPQueryNodeParameters) SetLimit(v int) *VMPQueryNodeParameters {
  s.Limit = &v
  return s
}

func (s *VMPQueryNodeParameters) SetMarker(v string) *VMPQueryNodeParameters {
  s.Marker = &v
  return s
}

func (s *VMPQueryNodeParameters) SetRegionName(v string) *VMPQueryNodeParameters {
  s.RegionName = &v
  return s
}

func (s *VMPQueryNodeParameters) SetProvince(v string) *VMPQueryNodeParameters {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeParameters) SetCarrier(v string) *VMPQueryNodeParameters {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeParameters) SetLineType(v string) *VMPQueryNodeParameters {
  s.LineType = &v
  return s
}

func (s *VMPQueryNodeParameters) SetIpv6Supported(v string) *VMPQueryNodeParameters {
  s.Ipv6Supported = &v
  return s
}

func (s *VMPQueryNodeParameters) SetBmSupported(v string) *VMPQueryNodeParameters {
  s.BmSupported = &v
  return s
}

type VMPQueryNodeRequestHeader struct {
}

func (s VMPQueryNodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryNodeResponseHeader struct {
}

func (s VMPQueryNodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryBandwidthRequest struct {
  // {"en":"VMPQueryBandwidthNode array", "zh_CN":"节点数组"}
  NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" type:"Repeated"`
}

func (s VMPQueryBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthRequest) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthRequest) SetNodeNames(v []*string) *VMPQueryBandwidthRequest {
  s.NodeNames = v
  return s
}

type VMPQueryBandwidthResponse struct {
  // {"en":"node", "zh_CN":"节点"}
  Nodes []*VMPQueryBandwidthNode `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthResponse) SetNodes(v []*VMPQueryBandwidthNode) *VMPQueryBandwidthResponse {
  s.Nodes = v
  return s
}

type VMPQueryBandwidthNode struct {
  // {"en":"VMPQueryBandwidthNode name", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"VMPQueryBandwidthNode bandwidth", "zh_CN":"节点带宽"}
  VMPQueryBandwidthNodeBw []*VMPQueryBandwidthNodeBw `json:"nodeBw,omitempty" xml:"nodeBw,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryBandwidthNode) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthNode) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthNode) SetNodeName(v string) *VMPQueryBandwidthNode {
  s.NodeName = &v
  return s
}

func (s *VMPQueryBandwidthNode) SetNodeBw(v []*VMPQueryBandwidthNodeBw) *VMPQueryBandwidthNode {
  s.VMPQueryBandwidthNodeBw = v
  return s
}

type VMPQueryBandwidthNodeBw struct {
  // {"en":"Operator code", "zh_CN":"运营商代码"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Whether the node has redundant bandwidth
  // True: redundant bandwidth
  // False: no redundant bandwidth
  // Undefined: node bandwidth statistics are not available. Whether there is redundant bandwidth is unknown. It is recommended to query again later.", "zh_CN":"节点是否有冗余带宽
  // True：有冗余带宽
  // False：无冗余带宽
  // Undefined：节点带宽统计数据不可用，是否有冗余带宽未知，建议稍后重新查询"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s VMPQueryBandwidthNodeBw) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthNodeBw) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthNodeBw) SetCarrier(v string) *VMPQueryBandwidthNodeBw {
  s.Carrier = &v
  return s
}

func (s *VMPQueryBandwidthNodeBw) SetResult(v string) *VMPQueryBandwidthNodeBw {
  s.Result = &v
  return s
}

type VMPQueryBandwidthPaths struct {
}

func (s VMPQueryBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthPaths) GoString() string {
  return s.String()
}

type VMPQueryBandwidthParameters struct {
}

func (s VMPQueryBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthParameters) GoString() string {
  return s.String()
}

type VMPQueryBandwidthRequestHeader struct {
}

func (s VMPQueryBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryBandwidthResponseHeader struct {
}

func (s VMPQueryBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthResponseHeader) GoString() string {
  return s.String()
}




type NodeRedundantBandwidth4PstatpRequest struct {
}

func (s NodeRedundantBandwidth4PstatpRequest) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpRequest) GoString() string {
  return s.String()
}

type NodeRedundantBandwidth4PstatpRequestHeader struct {
}

func (s NodeRedundantBandwidth4PstatpRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpRequestHeader) GoString() string {
  return s.String()
}

type NodeRedundantBandwidth4PstatpPaths struct {
}

func (s NodeRedundantBandwidth4PstatpPaths) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpPaths) GoString() string {
  return s.String()
}

type NodeRedundantBandwidth4PstatpParameters struct {
  // {"en":"Query account","zh_CN":"需要查询的账号"}
  AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
  // {"en":"Node ID","zh_CN":"厂商侧节点id，不填默认查询账号实际有在用的节点"}
  RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
  // {"en":"Page Number","zh_CN":"分页页数，不填默认返回全部数据"}
  PageNum *int `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
  // {"en":"Page size","zh_CN":"每页条数，不填时默认10条"}
  PageSize *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s NodeRedundantBandwidth4PstatpParameters) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpParameters) GoString() string {
  return s.String()
}

func (s *NodeRedundantBandwidth4PstatpParameters) SetAccountId(v string) *NodeRedundantBandwidth4PstatpParameters {
  s.AccountId = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpParameters) SetRoomId(v string) *NodeRedundantBandwidth4PstatpParameters {
  s.RoomId = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpParameters) SetPageNum(v int) *NodeRedundantBandwidth4PstatpParameters {
  s.PageNum = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpParameters) SetPageSize(v int) *NodeRedundantBandwidth4PstatpParameters {
  s.PageSize = &v
  return s
}

type NodeRedundantBandwidth4PstatpResponse struct {
  // {"en":"Code","zh_CN":"接口是否成功"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Datas","zh_CN":"带宽列表"}
  Data []*NodeRedundantBandwidth4PstatpResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s NodeRedundantBandwidth4PstatpResponse) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpResponse) GoString() string {
  return s.String()
}

func (s *NodeRedundantBandwidth4PstatpResponse) SetCode(v int) *NodeRedundantBandwidth4PstatpResponse {
  s.Code = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpResponse) SetData(v []*NodeRedundantBandwidth4PstatpResponseData) *NodeRedundantBandwidth4PstatpResponse {
  s.Data = v
  return s
}

type NodeRedundantBandwidth4PstatpResponseData struct     {
  // {"en":"Room ID","zh_CN":"节点ID"}
  RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
  // {"en":"Isp Bandwidth","zh_CN":"带宽列表"}
  IspBandwidth []*NodeRedundantBandwidth4PstatpResponseDataIspBandwidth `json:"IspBandwidth,omitempty" xml:"IspBandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s NodeRedundantBandwidth4PstatpResponseData) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpResponseData) GoString() string {
  return s.String()
}

func (s *NodeRedundantBandwidth4PstatpResponseData) SetRoomId(v string) *NodeRedundantBandwidth4PstatpResponseData {
  s.RoomId = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpResponseData) SetIspBandwidth(v []*NodeRedundantBandwidth4PstatpResponseDataIspBandwidth) *NodeRedundantBandwidth4PstatpResponseData {
  s.IspBandwidth = v
  return s
}

type NodeRedundantBandwidth4PstatpResponseDataIspBandwidth struct     {
  // {"en":"Isp","zh_CN":"运营商"}
  Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty" require:"true"`
  // {"en":"Current Bandwidth","zh_CN":"节点保障带宽"}
  CurrentBandwidth *int `json:"CurrentBandwidth,omitempty" xml:"CurrentBandwidth,omitempty" require:"true"`
  // {"en":"Available Elastic Bandwidth","zh_CN":"节点冗余带宽"}
  AvailableElasticBandwidth *int `json:"AvailableElasticBandwidth,omitempty" xml:"AvailableElasticBandwidth,omitempty" require:"true"`
}

func (s NodeRedundantBandwidth4PstatpResponseDataIspBandwidth) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpResponseDataIspBandwidth) GoString() string {
  return s.String()
}

func (s *NodeRedundantBandwidth4PstatpResponseDataIspBandwidth) SetIsp(v string) *NodeRedundantBandwidth4PstatpResponseDataIspBandwidth {
  s.Isp = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpResponseDataIspBandwidth) SetCurrentBandwidth(v int) *NodeRedundantBandwidth4PstatpResponseDataIspBandwidth {
  s.CurrentBandwidth = &v
  return s
}

func (s *NodeRedundantBandwidth4PstatpResponseDataIspBandwidth) SetAvailableElasticBandwidth(v int) *NodeRedundantBandwidth4PstatpResponseDataIspBandwidth {
  s.AvailableElasticBandwidth = &v
  return s
}

type NodeRedundantBandwidth4PstatpResponseHeader struct {
}

func (s NodeRedundantBandwidth4PstatpResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s NodeRedundantBandwidth4PstatpResponseHeader) GoString() string {
  return s.String()
}




