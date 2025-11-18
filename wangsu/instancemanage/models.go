package instancemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type LECHQueryInstanceRequest struct {
}

func (s LECHQueryInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceRequest) GoString() string {
  return s.String()
}

type LECHQueryInstanceRequestHeader struct {
}

func (s LECHQueryInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryInstancePaths struct {
}

func (s LECHQueryInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstancePaths) GoString() string {
  return s.String()
}

type LECHQueryInstanceParameters struct {
  // {"en":"Sort field name, can have more than one, value:\nName, regionName, createdAt, type, state, nodeName, etc","zh_CN":"排序的字段名称，可以有多个，取值：\nname、regionName、createdAt、province、state、nodeName等"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"There can be multiple field names for sorting, with values:\n\nName, regionname, createdat, province, state, nodeName, etc.'","zh_CN":"排序方向，必须跟在sortKey后面出现，取值：\ndesc：降序，默认值\nasc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default","zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the virtual machine ID specified by the marker","zh_CN":"从marker指定的实例id开始查询，升序查询（若要分页查询，则不能指定sortKey参数）"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Virtual machine ID. A maximum of 100 IDS can be queried at a time. The IDs are separated by a half angle comma character ','.","zh_CN":"实例ID。单次最多查询 100 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"Name of the region, for example, South China region is Huanan, refer to Appendix 3","zh_CN":"区域名称（区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Province of virtual machine","zh_CN":"实例所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Operator of virtual machine","zh_CN":"实例所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Virtual machine image identity","zh_CN":"实例镜像标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
  // {"en":"Virtual machine specification ID","zh_CN":"实例规格标识"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty"`
  // {"en":"Virtual machine name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Virtual machine ips. A maximum of 100 IPS can be queried at a time. The IPs are separated by a half angle comma character ','.","zh_CN":"实例IP。单次最多查询 100 条 IP，IP 之间用半角逗号字符','隔开。"}
  Ips *string `json:"ips,omitempty" xml:"ips,omitempty"`
  // {"en":"Virtual machine status, value (meaning to query virtual machines in the following status)\n\nRunning status\n\nBuilding new status\n\nStopped stop\n\nERROR error\n\nDeleting destroying\n\nRestarting\n\nStarting\n\nStopping","zh_CN":"实例状态，取值（意为查询处于下述状态的虚拟机）\nRUNNING 运行状态\nBUILDING 新建状态\nSTOPPED 停机\nERROR 错误\nDELETING 销毁中\nRESTARTING  重启中\nSTARTING  启动中\nSTOPPING  停止中"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Whether it is a free instance, value:\nYes free instances, NO billed instances","zh_CN":"是否免费实例，取值:\nYES 免费实例，NO 计费实例"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty"`
  // {"en":"1 means to query only bare-metal instances,\n-1 means that only virtual machine instances are queried,\nWithout this parameter all cloud hosts are queried","zh_CN":"1表示只查询裸机实例，\n-1表示只查询虚拟机实例，\n不带这个参数表示查询所有云主机"}
  IsBm *string `json:"isBm,omitempty" xml:"isBm,omitempty"`
  // {"en":"Cloud Host Label\nMultiple values are separated by a half corner comma, and the relationship between multiple values is or, that is, the instance label equals any one of these multiple values","zh_CN":"云主机标签\n多个值用半角逗号隔开，多个值是或者的关系，即实例标签等于这多个中的任意一个就满足条件"}
  Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
  // {"en":"ipv6 formate: 1: Zero Compressed(default); 6: With Leading Zero Suppression; 3: Full Address","zh_CN":"ipv6格式：1：省略零压缩格式(默认)；2：省略前导零格式; 3: 完整格式"}
  Ipv6Format *string `json:"ipv6Format,omitempty" xml:"ipv6Format,omitempty"`
}

func (s LECHQueryInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryInstanceParameters) SetSortKey(v string) *LECHQueryInstanceParameters {
  s.SortKey = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetSortDir(v string) *LECHQueryInstanceParameters {
  s.SortDir = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetLimit(v int) *LECHQueryInstanceParameters {
  s.Limit = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetMarker(v string) *LECHQueryInstanceParameters {
  s.Marker = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetIds(v string) *LECHQueryInstanceParameters {
  s.Ids = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetRegionName(v string) *LECHQueryInstanceParameters {
  s.RegionName = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetProvince(v string) *LECHQueryInstanceParameters {
  s.Province = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetCarrier(v string) *LECHQueryInstanceParameters {
  s.Carrier = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetImageId(v string) *LECHQueryInstanceParameters {
  s.ImageId = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetFlavorId(v string) *LECHQueryInstanceParameters {
  s.FlavorId = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetName(v string) *LECHQueryInstanceParameters {
  s.Name = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetIps(v string) *LECHQueryInstanceParameters {
  s.Ips = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetState(v string) *LECHQueryInstanceParameters {
  s.State = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetIsFree(v string) *LECHQueryInstanceParameters {
  s.IsFree = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetIsBm(v string) *LECHQueryInstanceParameters {
  s.IsBm = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetTags(v string) *LECHQueryInstanceParameters {
  s.Tags = &v
  return s
}

func (s *LECHQueryInstanceParameters) SetIpv6Format(v string) *LECHQueryInstanceParameters {
  s.Ipv6Format = &v
  return s
}

type LECHQueryInstanceResponse struct {
  // {"en":"Virtual machine information array","zh_CN":"实例信息数组"}
  Servers []*LECHQueryInstanceResponseServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryInstanceResponse) SetServers(v []*LECHQueryInstanceResponseServers) *LECHQueryInstanceResponse {
  s.Servers = v
  return s
}

type LECHQueryInstanceResponseServers struct     {
  // {"en":"Unique identity of virtual machine","zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Virtual Machine Area","zh_CN":"实例所属区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"Province of virtual machine","zh_CN":"实例所属省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Operator of virtual machine","zh_CN":"实例所属运营商"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Virtual machine image information","zh_CN":"实例镜像信息"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Virtual machine specifications","zh_CN":"实例规格"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Virtual machine name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Virtual machine status","zh_CN":"实例状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Virtual machine creation time","zh_CN":"实例创建时间"}
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty" require:"true"`
  // {"en":"Virtual machine IP address","zh_CN":"实例IP地址"}
  AccessIPv4 *string `json:"accessIPv4,omitempty" xml:"accessIPv4,omitempty" require:"true"`
  // {"en":"IPv4 address of virtual machine intranet. If intranet is specified when creating virtual machine, intranet IP is returned. Otherwise, it is empty.","zh_CN":"实例内网IPv4地址，如果创建实例时指定了需要内网，则返回内网IPv4，否则是空"}
  PrivateIPv4 *string `json:"privateIPv4,omitempty" xml:"privateIPv4,omitempty" require:"true"`
  // {"en":"IPv6 address of virtual machine intranet. If intranet is specified when creating virtual machine, intranet IP is returned. Otherwise, it is empty.","zh_CN":"实例内网IPv6地址，如果创建实例时指定了需要内网，则返回内网IPv6，否则是空"}
  PrivateIPv6 *string `json:"privateIPv6,omitempty" xml:"privateIPv6,omitempty" require:"true"`
  // {"en":"Virtual machine login SSH key pair name","zh_CN":"实例登录SSH秘钥对名称"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty" require:"true"`
  // {"en":"Node name, the name of the node where the virtual machine is located. Through this node name, you can query the real-time redundant bandwidth of each node.","zh_CN":"节点名称，实例所在节点名称，通过这个节点名称可以查询每个节点的实时冗余带宽情况"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Virtual machine IP address","zh_CN":"实例ip地址"}
  AccessIP []*LECHQueryInstanceResponseServersAccessIP `json:"accessIP,omitempty" xml:"accessIP,omitempty" require:"true" type:"Repeated"`
  // {"en":"If it is free instance, value: YES free instance, NO billing instance","zh_CN":"是否免费实例，取值：YES 免费实例，NO 计费实例"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty" require:"true"`
  // {"en":"1 indicates that the instance is bare metal\n-1 indicates that the instance is a virtual machine","zh_CN":"1表示该实例是裸机\n-1表示该实例是虚拟机"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"A list of security group IDs for instance bindings","zh_CN":"实例绑定的安全组id列表"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" require:"true" type:"Repeated"`
  // {"en":"Disk information","zh_CN":"磁盘信息"}
  DiskInfo []*LECHQueryInstanceResponseServersDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
  // {"en":"instance tag","zh_CN":"实例标签"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty" require:"true"`
}

func (s LECHQueryInstanceResponseServers) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceResponseServers) GoString() string {
  return s.String()
}

func (s *LECHQueryInstanceResponseServers) SetId(v string) *LECHQueryInstanceResponseServers {
  s.Id = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetRegionName(v string) *LECHQueryInstanceResponseServers {
  s.RegionName = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetProvince(v string) *LECHQueryInstanceResponseServers {
  s.Province = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetCarrier(v string) *LECHQueryInstanceResponseServers {
  s.Carrier = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetImageId(v string) *LECHQueryInstanceResponseServers {
  s.ImageId = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetFlavorId(v string) *LECHQueryInstanceResponseServers {
  s.FlavorId = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetName(v string) *LECHQueryInstanceResponseServers {
  s.Name = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetState(v string) *LECHQueryInstanceResponseServers {
  s.State = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetCreatedAt(v string) *LECHQueryInstanceResponseServers {
  s.CreatedAt = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetAccessIPv4(v string) *LECHQueryInstanceResponseServers {
  s.AccessIPv4 = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetPrivateIPv4(v string) *LECHQueryInstanceResponseServers {
  s.PrivateIPv4 = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetPrivateIPv6(v string) *LECHQueryInstanceResponseServers {
  s.PrivateIPv6 = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetKeyName(v string) *LECHQueryInstanceResponseServers {
  s.KeyName = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetNodeName(v string) *LECHQueryInstanceResponseServers {
  s.NodeName = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetAccessIP(v []*LECHQueryInstanceResponseServersAccessIP) *LECHQueryInstanceResponseServers {
  s.AccessIP = v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetIsFree(v string) *LECHQueryInstanceResponseServers {
  s.IsFree = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetIsBm(v int) *LECHQueryInstanceResponseServers {
  s.IsBm = &v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetSecurityGroupIds(v []*string) *LECHQueryInstanceResponseServers {
  s.SecurityGroupIds = v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetDiskInfo(v []*LECHQueryInstanceResponseServersDiskInfo) *LECHQueryInstanceResponseServers {
  s.DiskInfo = v
  return s
}

func (s *LECHQueryInstanceResponseServers) SetTag(v string) *LECHQueryInstanceResponseServers {
  s.Tag = &v
  return s
}

type LECHQueryInstanceResponseServersAccessIP struct     {
  // {"en":"Ip address","zh_CN":"Ip地址"}
  Address *string `json:"address,omitempty" xml:"address,omitempty" require:"true"`
  // {"en":"IP address operator","zh_CN":"Ip地址所属运营商"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Protocol type: 4: IPv4 address; 6: IPv6 address","zh_CN":"协议类型：4：ipv4地址；6：ipv6地址"}
  Protocol *int `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s LECHQueryInstanceResponseServersAccessIP) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceResponseServersAccessIP) GoString() string {
  return s.String()
}

func (s *LECHQueryInstanceResponseServersAccessIP) SetAddress(v string) *LECHQueryInstanceResponseServersAccessIP {
  s.Address = &v
  return s
}

func (s *LECHQueryInstanceResponseServersAccessIP) SetCarrier(v string) *LECHQueryInstanceResponseServersAccessIP {
  s.Carrier = &v
  return s
}

func (s *LECHQueryInstanceResponseServersAccessIP) SetProtocol(v int) *LECHQueryInstanceResponseServersAccessIP {
  s.Protocol = &v
  return s
}

type LECHQueryInstanceResponseServersDiskInfo struct     {
  // {"en":"Disk size (GB)","zh_CN":"磁盘大小（GB）"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk Purpose: System - System disk;DATA - DATA plate","zh_CN":"磁盘用途：SYSTEM-系统盘；DATA-数据盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk type: HDD/SSD","zh_CN":"磁盘类型：HDD/SSD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s LECHQueryInstanceResponseServersDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceResponseServersDiskInfo) GoString() string {
  return s.String()
}

func (s *LECHQueryInstanceResponseServersDiskInfo) SetSize(v int) *LECHQueryInstanceResponseServersDiskInfo {
  s.Size = &v
  return s
}

func (s *LECHQueryInstanceResponseServersDiskInfo) SetType(v string) *LECHQueryInstanceResponseServersDiskInfo {
  s.Type = &v
  return s
}

func (s *LECHQueryInstanceResponseServersDiskInfo) SetCategory(v string) *LECHQueryInstanceResponseServersDiskInfo {
  s.Category = &v
  return s
}

type LECHQueryInstanceResponseHeader struct {
}

func (s LECHQueryInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryInstanceResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceDiskScalingRequest struct {
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DiskInfo []*LECHInstanceDiskScalingRequestDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
}

func (s LECHInstanceDiskScalingRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingRequest) GoString() string {
  return s.String()
}

func (s *LECHInstanceDiskScalingRequest) SetId(v string) *LECHInstanceDiskScalingRequest {
  s.Id = &v
  return s
}

func (s *LECHInstanceDiskScalingRequest) SetDiskInfo(v []*LECHInstanceDiskScalingRequestDiskInfo) *LECHInstanceDiskScalingRequest {
  s.DiskInfo = v
  return s
}

type LECHInstanceDiskScalingRequestDiskInfo struct     {
  // {"en":"disk size GB","zh_CN":"磁盘大小，单位GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"disk type ,HDD/SDD","zh_CN":"磁盘类型，取值：\nHDD：普通硬盘\nSSD：固态硬盘"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s LECHInstanceDiskScalingRequestDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingRequestDiskInfo) GoString() string {
  return s.String()
}

func (s *LECHInstanceDiskScalingRequestDiskInfo) SetSize(v int) *LECHInstanceDiskScalingRequestDiskInfo {
  s.Size = &v
  return s
}

func (s *LECHInstanceDiskScalingRequestDiskInfo) SetCategory(v string) *LECHInstanceDiskScalingRequestDiskInfo {
  s.Category = &v
  return s
}

type LECHInstanceDiskScalingRequestHeader struct {
}

func (s LECHInstanceDiskScalingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceDiskScalingPaths struct {
}

func (s LECHInstanceDiskScalingPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingPaths) GoString() string {
  return s.String()
}

type LECHInstanceDiskScalingParameters struct {
}

func (s LECHInstanceDiskScalingParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingParameters) GoString() string {
  return s.String()
}

type LECHInstanceDiskScalingResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHInstanceDiskScalingResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHInstanceDiskScalingResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingResponse) GoString() string {
  return s.String()
}

func (s *LECHInstanceDiskScalingResponse) SetCode(v string) *LECHInstanceDiskScalingResponse {
  s.Code = &v
  return s
}

func (s *LECHInstanceDiskScalingResponse) SetMessage(v string) *LECHInstanceDiskScalingResponse {
  s.Message = &v
  return s
}

func (s *LECHInstanceDiskScalingResponse) SetData(v *LECHInstanceDiskScalingResponseData) *LECHInstanceDiskScalingResponse {
  s.Data = v
  return s
}

type LECHInstanceDiskScalingResponseData struct {
  // {"en":"Instance info","zh_CN":"实例信息"}
  Server *LECHInstanceDiskScalingResponseDataServer `json:"server,omitempty" xml:"server,omitempty" require:"true" type:"Struct"`
}

func (s LECHInstanceDiskScalingResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingResponseData) GoString() string {
  return s.String()
}

func (s *LECHInstanceDiskScalingResponseData) SetServer(v *LECHInstanceDiskScalingResponseDataServer) *LECHInstanceDiskScalingResponseData {
  s.Server = v
  return s
}

type LECHInstanceDiskScalingResponseDataServer struct {
  // {"en":"Instance ID","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Disk info","zh_CN":"磁盘信息"}
  DiskInfo []*LECHInstanceDiskScalingResponseDataServerDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
}

func (s LECHInstanceDiskScalingResponseDataServer) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingResponseDataServer) GoString() string {
  return s.String()
}

func (s *LECHInstanceDiskScalingResponseDataServer) SetId(v string) *LECHInstanceDiskScalingResponseDataServer {
  s.Id = &v
  return s
}

func (s *LECHInstanceDiskScalingResponseDataServer) SetDiskInfo(v []*LECHInstanceDiskScalingResponseDataServerDiskInfo) *LECHInstanceDiskScalingResponseDataServer {
  s.DiskInfo = v
  return s
}

type LECHInstanceDiskScalingResponseDataServerDiskInfo struct     {
  // {"en":"Disk size","zh_CN":"磁盘大小，单位GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk type","zh_CN":"磁盘类型,DATA：数据盘, SYSTEM：系统盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk category","zh_CN":"磁盘类型，HDD/SDD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s LECHInstanceDiskScalingResponseDataServerDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingResponseDataServerDiskInfo) GoString() string {
  return s.String()
}

func (s *LECHInstanceDiskScalingResponseDataServerDiskInfo) SetSize(v int) *LECHInstanceDiskScalingResponseDataServerDiskInfo {
  s.Size = &v
  return s
}

func (s *LECHInstanceDiskScalingResponseDataServerDiskInfo) SetType(v string) *LECHInstanceDiskScalingResponseDataServerDiskInfo {
  s.Type = &v
  return s
}

func (s *LECHInstanceDiskScalingResponseDataServerDiskInfo) SetCategory(v string) *LECHInstanceDiskScalingResponseDataServerDiskInfo {
  s.Category = &v
  return s
}

type LECHInstanceDiskScalingResponseHeader struct {
}

func (s LECHInstanceDiskScalingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceDiskScalingResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceBandwidth5MinQueryRequest struct {
}

func (s LECHInstanceBandwidth5MinQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryRequest) GoString() string {
  return s.String()
}

type LECHInstanceBandwidth5MinQueryRequestHeader struct {
}

func (s LECHInstanceBandwidth5MinQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceBandwidth5MinQueryPaths struct {
}

func (s LECHInstanceBandwidth5MinQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryPaths) GoString() string {
  return s.String()
}

type LECHInstanceBandwidth5MinQueryParameters struct {
  // {"en":"The ID of the virtual machine to query. At most 30 queries can be made at a time, ids\nare separated by character  ','","zh_CN":"云主机ID。单次最多查询 20 条 ID，ID 之间用半角逗号字符','隔开。支持查询已销毁实例的带宽数据"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Query time range in format: yyyymmddhhmm-yyyymmddhhmm\nQuery data for nearly 90 days, a single query no more than 3 days.\nFor example: 202001201730-202001201930 means to query 2020-01-20 17:30 to 19:30 monitoring data.","zh_CN":"查询时间范围，格式：yyyyMMddHHmm- yyyyMMddHHmm\n查询近90天的数据，单次查询不超过10天。\n例如：202001201730-202001201930表示查询2020-01-20 17:30到19:30的带宽数据。"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"The ISP code","zh_CN":"主要用于多线实例流量拆分时过滤所属运营商：dx-电信；wt-网通；yd-移动。一次仅允许传入一个运营商参数"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
}

func (s LECHInstanceBandwidth5MinQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryParameters) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidth5MinQueryParameters) SetIds(v string) *LECHInstanceBandwidth5MinQueryParameters {
  s.Ids = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryParameters) SetStatTime(v string) *LECHInstanceBandwidth5MinQueryParameters {
  s.StatTime = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryParameters) SetCarrier(v string) *LECHInstanceBandwidth5MinQueryParameters {
  s.Carrier = &v
  return s
}

type LECHInstanceBandwidth5MinQueryResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *LECHInstanceBandwidth5MinQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHInstanceBandwidth5MinQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryResponse) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidth5MinQueryResponse) SetCode(v string) *LECHInstanceBandwidth5MinQueryResponse {
  s.Code = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponse) SetData(v *LECHInstanceBandwidth5MinQueryResponseData) *LECHInstanceBandwidth5MinQueryResponse {
  s.Data = v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponse) SetMessage(v string) *LECHInstanceBandwidth5MinQueryResponse {
  s.Message = &v
  return s
}

type LECHInstanceBandwidth5MinQueryResponseData struct {
  // {"en":"Instance information array","zh_CN":"实例信息数组"}
  Servers []*LECHInstanceBandwidth5MinQueryResponseDataServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s LECHInstanceBandwidth5MinQueryResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryResponseData) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidth5MinQueryResponseData) SetServers(v []*LECHInstanceBandwidth5MinQueryResponseDataServers) *LECHInstanceBandwidth5MinQueryResponseData {
  s.Servers = v
  return s
}

type LECHInstanceBandwidth5MinQueryResponseDataServers struct     {
  // {"en":"Bandwidth information","zh_CN":"带宽信息"}
  Bandwidths []*LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths `json:"bandwidths,omitempty" xml:"bandwidths,omitempty" require:"true" type:"Repeated"`
  // {"en":"instance id","zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHInstanceBandwidth5MinQueryResponseDataServers) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryResponseDataServers) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServers) SetBandwidths(v []*LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) *LECHInstanceBandwidth5MinQueryResponseDataServers {
  s.Bandwidths = v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServers) SetId(v string) *LECHInstanceBandwidth5MinQueryResponseDataServers {
  s.Id = &v
  return s
}

type LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths struct     {
  // {"en":"ISP","zh_CN":"运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Inflow bandwidth of external network (not including intra-node traffic), unit of Mbps","zh_CN":"外网流入带宽（不含节点内流量），单位Mbps"}
  ExtIn *string `json:"extIn,omitempty" xml:"extIn,omitempty" require:"true"`
  // {"en":"Outflow bandwidth of the external network (excluding intra-node traffic), unit of Mbps","zh_CN":"外网流出带宽（不含节点内流量），单位Mbps"}
  ExtOut *string `json:"extOut,omitempty" xml:"extOut,omitempty" require:"true"`
  // {"en":"Total inflow bandwidth (including intra-node traffic) in Mbps","zh_CN":"总流入带宽（含节点内流量），单位Mbps"}
  In *string `json:"in,omitempty" xml:"in,omitempty" require:"true"`
  // {"en":"instance ip","zh_CN":"实例IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Total outflow bandwidth (including intra-node flow) in Mbps","zh_CN":"总流出带宽（含节点内流量），单位Mbps"}
  Out *string `json:"out,omitempty" xml:"out,omitempty" require:"true"`
  // {"en":"The bandwidth collection time, formatted as YYYYMMDDHHMM, is 5 minutes granularity","zh_CN":"带宽采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
}

func (s LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetCarrier(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.Carrier = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetExtIn(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.ExtIn = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetExtOut(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.ExtOut = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetIn(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.In = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetIp(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.Ip = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetOut(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.Out = &v
  return s
}

func (s *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetStatTime(v string) *LECHInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.StatTime = &v
  return s
}

type LECHInstanceBandwidth5MinQueryResponseHeader struct {
}

func (s LECHInstanceBandwidth5MinQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidth5MinQueryResponseHeader) GoString() string {
  return s.String()
}




type InstanceReplaceIpRequest struct {
  // {"en":"vm id","zh_CN":"云主机ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"server old replace ips","zh_CN":"虚拟机待更换ip列表"}
  InstanceIps []*string `json:"instanceIps,omitempty" xml:"instanceIps,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceReplaceIpRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpRequest) GoString() string {
  return s.String()
}

func (s *InstanceReplaceIpRequest) SetInstanceId(v string) *InstanceReplaceIpRequest {
  s.InstanceId = &v
  return s
}

func (s *InstanceReplaceIpRequest) SetInstanceIps(v []*string) *InstanceReplaceIpRequest {
  s.InstanceIps = v
  return s
}

type InstanceReplaceIpRequestHeader struct {
}

func (s InstanceReplaceIpRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpRequestHeader) GoString() string {
  return s.String()
}

type InstanceReplaceIpPaths struct {
}

func (s InstanceReplaceIpPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpPaths) GoString() string {
  return s.String()
}

type InstanceReplaceIpParameters struct {
}

func (s InstanceReplaceIpParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpParameters) GoString() string {
  return s.String()
}

type InstanceReplaceIpResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *InstanceReplaceIpResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s InstanceReplaceIpResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpResponse) GoString() string {
  return s.String()
}

func (s *InstanceReplaceIpResponse) SetCode(v string) *InstanceReplaceIpResponse {
  s.Code = &v
  return s
}

func (s *InstanceReplaceIpResponse) SetMessage(v string) *InstanceReplaceIpResponse {
  s.Message = &v
  return s
}

func (s *InstanceReplaceIpResponse) SetData(v *InstanceReplaceIpResponseData) *InstanceReplaceIpResponse {
  s.Data = v
  return s
}

type InstanceReplaceIpResponseData struct {
}

func (s InstanceReplaceIpResponseData) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpResponseData) GoString() string {
  return s.String()
}

type InstanceReplaceIpResponseHeader struct {
}

func (s InstanceReplaceIpResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceReplaceIpResponseHeader) GoString() string {
  return s.String()
}




type SetTrafficLimitRequest struct {
  // {"en":"instance id list","zh_CN":"实例id列表，单次最多填写50个"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" require:"true" type:"Repeated"`
  // {"en":"Natural Month Traffic limit(G), range 1-65535. Traffic limit cannot exceed order constraints.. Not filling in means no limitation.","zh_CN":"自然月流量上限值，范围1-65535。流量上限值不可超过订单约束的上限值。不填意为不限制。"}
  TrafficLimit *int `json:"trafficLimit,omitempty" xml:"trafficLimit,omitempty"`
}

func (s SetTrafficLimitRequest) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitRequest) GoString() string {
  return s.String()
}

func (s *SetTrafficLimitRequest) SetIds(v []*string) *SetTrafficLimitRequest {
  s.Ids = v
  return s
}

func (s *SetTrafficLimitRequest) SetTrafficLimit(v int) *SetTrafficLimitRequest {
  s.TrafficLimit = &v
  return s
}

type SetTrafficLimitRequestHeader struct {
}

func (s SetTrafficLimitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitRequestHeader) GoString() string {
  return s.String()
}

type SetTrafficLimitPaths struct {
}

func (s SetTrafficLimitPaths) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitPaths) GoString() string {
  return s.String()
}

type SetTrafficLimitParameters struct {
}

func (s SetTrafficLimitParameters) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitParameters) GoString() string {
  return s.String()
}

type SetTrafficLimitResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data []*SetTrafficLimitResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s SetTrafficLimitResponse) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitResponse) GoString() string {
  return s.String()
}

func (s *SetTrafficLimitResponse) SetCode(v string) *SetTrafficLimitResponse {
  s.Code = &v
  return s
}

func (s *SetTrafficLimitResponse) SetData(v []*SetTrafficLimitResponseData) *SetTrafficLimitResponse {
  s.Data = v
  return s
}

func (s *SetTrafficLimitResponse) SetMessage(v string) *SetTrafficLimitResponse {
  s.Message = &v
  return s
}

type SetTrafficLimitResponseData struct     {
  // {"en":"result code","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"result message","zh_CN":"结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"instance name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s SetTrafficLimitResponseData) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitResponseData) GoString() string {
  return s.String()
}

func (s *SetTrafficLimitResponseData) SetCode(v int) *SetTrafficLimitResponseData {
  s.Code = &v
  return s
}

func (s *SetTrafficLimitResponseData) SetId(v string) *SetTrafficLimitResponseData {
  s.Id = &v
  return s
}

func (s *SetTrafficLimitResponseData) SetMessage(v string) *SetTrafficLimitResponseData {
  s.Message = &v
  return s
}

func (s *SetTrafficLimitResponseData) SetName(v string) *SetTrafficLimitResponseData {
  s.Name = &v
  return s
}

type SetTrafficLimitResponseHeader struct {
}

func (s SetTrafficLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SetTrafficLimitResponseHeader) GoString() string {
  return s.String()
}




type ConvertFreeTypeInstanceToChargeTypeRequest struct {
  // {"en":"Unique cloud host identity.Up to 100 IDs can be sent at a time, separated by the half comma character ', '.","zh_CN":"云主机唯一标识。单次最多可发送100 条ID，ID 之间用半角逗号字符','隔开。"}
  Servers *string `json:"servers,omitempty" xml:"servers,omitempty" require:"true"`
}

func (s ConvertFreeTypeInstanceToChargeTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeRequest) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeRequest) SetServers(v string) *ConvertFreeTypeInstanceToChargeTypeRequest {
  s.Servers = &v
  return s
}

type ConvertFreeTypeInstanceToChargeTypeRequestHeader struct {
}

func (s ConvertFreeTypeInstanceToChargeTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeRequestHeader) GoString() string {
  return s.String()
}

type ConvertFreeTypeInstanceToChargeTypePaths struct {
}

func (s ConvertFreeTypeInstanceToChargeTypePaths) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypePaths) GoString() string {
  return s.String()
}

type ConvertFreeTypeInstanceToChargeTypeParameters struct {
}

func (s ConvertFreeTypeInstanceToChargeTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeParameters) GoString() string {
  return s.String()
}

type ConvertFreeTypeInstanceToChargeTypeResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *ConvertFreeTypeInstanceToChargeTypeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ConvertFreeTypeInstanceToChargeTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeResponse) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponse) SetCode(v string) *ConvertFreeTypeInstanceToChargeTypeResponse {
  s.Code = &v
  return s
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponse) SetData(v *ConvertFreeTypeInstanceToChargeTypeResponseData) *ConvertFreeTypeInstanceToChargeTypeResponse {
  s.Data = v
  return s
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponse) SetMessage(v string) *ConvertFreeTypeInstanceToChargeTypeResponse {
  s.Message = &v
  return s
}

type ConvertFreeTypeInstanceToChargeTypeResponseData struct {
  // {"en":"Batch error","zh_CN":"批量操作失败信息"}
  BatchErrorMsg []*ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseData) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseData) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponseData) SetBatchErrorMsg(v []*ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) *ConvertFreeTypeInstanceToChargeTypeResponseData {
  s.BatchErrorMsg = v
  return s
}

type ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg struct     {
  // {"en":"Error code","zh_CN":"错误编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Instance ID","zh_CN":"实例id"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) SetCode(v string) *ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) SetKey(v string) *ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) SetMsg(v string) *ConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type ConvertFreeTypeInstanceToChargeTypeResponseHeader struct {
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseHeader) GoString() string {
  return s.String()
}




type LECHRemoveInstanceRequest struct {
  // {"en":"Instance ID","zh_CN":"实例唯一标识。单次最多可发送100 条ID，ID 之间用半角逗号字符“,”隔开。"}
  Servers *string `json:"servers,omitempty" xml:"servers,omitempty" require:"true"`
}

func (s LECHRemoveInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveInstanceRequest) GoString() string {
  return s.String()
}

func (s *LECHRemoveInstanceRequest) SetServers(v string) *LECHRemoveInstanceRequest {
  s.Servers = &v
  return s
}

type LECHRemoveInstanceRequestHeader struct {
}

func (s LECHRemoveInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveInstanceRequestHeader) GoString() string {
  return s.String()
}

type LECHRemoveInstancePaths struct {
}

func (s LECHRemoveInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveInstancePaths) GoString() string {
  return s.String()
}

type LECHRemoveInstanceParameters struct {
}

func (s LECHRemoveInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveInstanceParameters) GoString() string {
  return s.String()
}

type LECHRemoveInstanceResponse struct {
}

func (s LECHRemoveInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveInstanceResponse) GoString() string {
  return s.String()
}

type LECHRemoveInstanceResponseHeader struct {
}

func (s LECHRemoveInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveInstanceResponseHeader) GoString() string {
  return s.String()
}




type ManageInstanceTagsRequest struct {
  // {"en":"The cloud host ID, with multiple IDs separated by the half corner comma character ', '.","zh_CN":"云主机id，多个ID 之间用半角逗号字符','隔开。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"New instance label value;If the value is null, the instance label of the specified cloud host is deleted.","zh_CN":"新的实例标签值；如果值为空，则表示删除指定云主机的实例标签。"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty" require:"true"`
}

func (s ManageInstanceTagsRequest) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsRequest) GoString() string {
  return s.String()
}

func (s *ManageInstanceTagsRequest) SetId(v string) *ManageInstanceTagsRequest {
  s.Id = &v
  return s
}

func (s *ManageInstanceTagsRequest) SetTag(v string) *ManageInstanceTagsRequest {
  s.Tag = &v
  return s
}

type ManageInstanceTagsRequestHeader struct {
}

func (s ManageInstanceTagsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsRequestHeader) GoString() string {
  return s.String()
}

type ManageInstanceTagsPaths struct {
}

func (s ManageInstanceTagsPaths) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsPaths) GoString() string {
  return s.String()
}

type ManageInstanceTagsParameters struct {
}

func (s ManageInstanceTagsParameters) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsParameters) GoString() string {
  return s.String()
}

type ManageInstanceTagsResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *ManageInstanceTagsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ManageInstanceTagsResponse) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsResponse) GoString() string {
  return s.String()
}

func (s *ManageInstanceTagsResponse) SetCode(v string) *ManageInstanceTagsResponse {
  s.Code = &v
  return s
}

func (s *ManageInstanceTagsResponse) SetMessage(v string) *ManageInstanceTagsResponse {
  s.Message = &v
  return s
}

func (s *ManageInstanceTagsResponse) SetData(v *ManageInstanceTagsResponseData) *ManageInstanceTagsResponse {
  s.Data = v
  return s
}

type ManageInstanceTagsResponseData struct {
  // {"en":"Batch error","zh_CN":"批量操作失败信息"}
  BatchErrorMsg []*ManageInstanceTagsResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s ManageInstanceTagsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsResponseData) GoString() string {
  return s.String()
}

func (s *ManageInstanceTagsResponseData) SetBatchErrorMsg(v []*ManageInstanceTagsResponseDataBatchErrorMsg) *ManageInstanceTagsResponseData {
  s.BatchErrorMsg = v
  return s
}

type ManageInstanceTagsResponseDataBatchErrorMsg struct     {
  // {"en":"Instance ID","zh_CN":"实例ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Error code","zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ManageInstanceTagsResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *ManageInstanceTagsResponseDataBatchErrorMsg) SetKey(v string) *ManageInstanceTagsResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *ManageInstanceTagsResponseDataBatchErrorMsg) SetCode(v string) *ManageInstanceTagsResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *ManageInstanceTagsResponseDataBatchErrorMsg) SetMsg(v string) *ManageInstanceTagsResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type ManageInstanceTagsResponseHeader struct {
}

func (s ManageInstanceTagsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsResponseHeader) GoString() string {
  return s.String()
}




type VmpInstanceBandwidth5MinQueryRequest struct {
}

func (s VmpInstanceBandwidth5MinQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryRequest) GoString() string {
  return s.String()
}

type VmpInstanceBandwidth5MinQueryRequestHeader struct {
}

func (s VmpInstanceBandwidth5MinQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryRequestHeader) GoString() string {
  return s.String()
}

type VmpInstanceBandwidth5MinQueryPaths struct {
}

func (s VmpInstanceBandwidth5MinQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryPaths) GoString() string {
  return s.String()
}

type VmpInstanceBandwidth5MinQueryParameters struct {
  // {"en":"The ID of the virtual machine to query. At most 30 queries can be made at a time, ids\nare separated by character  ','","zh_CN":"云主机ID。单次最多查询 20 条 ID，ID 之间用半角逗号字符','隔开。支持查询已销毁实例的带宽数据"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Query time range in format: yyyymmddhhmm-yyyymmddhhmm\nQuery data for nearly 90 days, a single query no more than 3 days.\nFor example: 202001201730-202001201930 means to query 2020-01-20 17:30 to 19:30 monitoring data.","zh_CN":"查询时间范围，格式：yyyyMMddHHmm- yyyyMMddHHmm\n查询近90天的数据，单次查询不超过10天。\n例如：202001201730-202001201930表示查询2020-01-20 17:30到19:30的带宽数据。"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"The ISP code","zh_CN":"主要用于多线实例流量拆分时过滤所属运营商：dx-电信；wt-网通；yd-移动。一次仅允许传入一个运营商参数"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
}

func (s VmpInstanceBandwidth5MinQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryParameters) GoString() string {
  return s.String()
}

func (s *VmpInstanceBandwidth5MinQueryParameters) SetIds(v string) *VmpInstanceBandwidth5MinQueryParameters {
  s.Ids = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryParameters) SetStatTime(v string) *VmpInstanceBandwidth5MinQueryParameters {
  s.StatTime = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryParameters) SetCarrier(v string) *VmpInstanceBandwidth5MinQueryParameters {
  s.Carrier = &v
  return s
}

type VmpInstanceBandwidth5MinQueryResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *VmpInstanceBandwidth5MinQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s VmpInstanceBandwidth5MinQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryResponse) GoString() string {
  return s.String()
}

func (s *VmpInstanceBandwidth5MinQueryResponse) SetCode(v string) *VmpInstanceBandwidth5MinQueryResponse {
  s.Code = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponse) SetData(v *VmpInstanceBandwidth5MinQueryResponseData) *VmpInstanceBandwidth5MinQueryResponse {
  s.Data = v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponse) SetMessage(v string) *VmpInstanceBandwidth5MinQueryResponse {
  s.Message = &v
  return s
}

type VmpInstanceBandwidth5MinQueryResponseData struct {
  // {"en":"Instance information array","zh_CN":"实例信息数组"}
  Servers []*VmpInstanceBandwidth5MinQueryResponseDataServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VmpInstanceBandwidth5MinQueryResponseData) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryResponseData) GoString() string {
  return s.String()
}

func (s *VmpInstanceBandwidth5MinQueryResponseData) SetServers(v []*VmpInstanceBandwidth5MinQueryResponseDataServers) *VmpInstanceBandwidth5MinQueryResponseData {
  s.Servers = v
  return s
}

type VmpInstanceBandwidth5MinQueryResponseDataServers struct     {
  // {"en":"Bandwidth information","zh_CN":"带宽信息"}
  Bandwidths []*VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths `json:"bandwidths,omitempty" xml:"bandwidths,omitempty" require:"true" type:"Repeated"`
  // {"en":"instance id","zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s VmpInstanceBandwidth5MinQueryResponseDataServers) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryResponseDataServers) GoString() string {
  return s.String()
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServers) SetBandwidths(v []*VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) *VmpInstanceBandwidth5MinQueryResponseDataServers {
  s.Bandwidths = v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServers) SetId(v string) *VmpInstanceBandwidth5MinQueryResponseDataServers {
  s.Id = &v
  return s
}

type VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths struct     {
  // {"en":"ISP","zh_CN":"运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Inflow bandwidth of external network (not including intra-node traffic), unit of Mbps","zh_CN":"外网流入带宽（不含节点内流量），单位Mbps"}
  ExtIn *string `json:"extIn,omitempty" xml:"extIn,omitempty" require:"true"`
  // {"en":"Outflow bandwidth of the external network (excluding intra-node traffic), unit of Mbps","zh_CN":"外网流出带宽（不含节点内流量），单位Mbps"}
  ExtOut *string `json:"extOut,omitempty" xml:"extOut,omitempty" require:"true"`
  // {"en":"Total inflow bandwidth (including intra-node traffic) in Mbps","zh_CN":"总流入带宽（含节点内流量），单位Mbps"}
  In *string `json:"in,omitempty" xml:"in,omitempty" require:"true"`
  // {"en":"instance ip","zh_CN":"实例IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Total outflow bandwidth (including intra-node flow) in Mbps","zh_CN":"总流出带宽（含节点内流量），单位Mbps"}
  Out *string `json:"out,omitempty" xml:"out,omitempty" require:"true"`
  // {"en":"The bandwidth collection time, formatted as YYYYMMDDHHMM, is 5 minutes granularity","zh_CN":"带宽采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
}

func (s VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) GoString() string {
  return s.String()
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetCarrier(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.Carrier = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetExtIn(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.ExtIn = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetExtOut(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.ExtOut = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetIn(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.In = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetIp(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.Ip = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetOut(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.Out = &v
  return s
}

func (s *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths) SetStatTime(v string) *VmpInstanceBandwidth5MinQueryResponseDataServersBandwidths {
  s.StatTime = &v
  return s
}

type VmpInstanceBandwidth5MinQueryResponseHeader struct {
}

func (s VmpInstanceBandwidth5MinQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VmpInstanceBandwidth5MinQueryResponseHeader) GoString() string {
  return s.String()
}




type UpdateEcciInstanceRequest struct {
  // {"en":"instance distribution", "zh_CN":"实例分发属性"}
  Distributions *UpdateEcciInstanceInstanceDistribution `json:"distributions,omitempty" xml:"distributions,omitempty" require:"true"`
  // {"en":"Whether to assign public IP", "zh_CN":"是否分配公网ip"}
  AllocatePublicIp *bool `json:"allocatePublicIp,omitempty" xml:"allocatePublicIp,omitempty" require:"true"`
  // {"en":"Specification of the desired behavior of the pod.", "zh_CN":"Pod 预期行为的规约。"}
  UpdateEcciInstancePodSpec *UpdateEcciInstancePodSpec `json:"podSpec,omitempty" xml:"podSpec,omitempty" require:"true"`
  // {"en":"assigned ip attributes", "zh_CN":"分配ip属性"}
  AssignIps []*UpdateEcciInstanceAssignIp `json:"assignIps,omitempty" xml:"assignIps,omitempty" type:"Repeated"`
}

func (s UpdateEcciInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceRequest) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceRequest) SetDistributions(v *UpdateEcciInstanceInstanceDistribution) *UpdateEcciInstanceRequest {
  s.Distributions = v
  return s
}

func (s *UpdateEcciInstanceRequest) SetAllocatePublicIp(v bool) *UpdateEcciInstanceRequest {
  s.AllocatePublicIp = &v
  return s
}

func (s *UpdateEcciInstanceRequest) SetPodSpec(v *UpdateEcciInstancePodSpec) *UpdateEcciInstanceRequest {
  s.UpdateEcciInstancePodSpec = v
  return s
}

func (s *UpdateEcciInstanceRequest) SetAssignIps(v []*UpdateEcciInstanceAssignIp) *UpdateEcciInstanceRequest {
  s.AssignIps = v
  return s
}

type UpdateEcciInstanceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s UpdateEcciInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceResponse) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceResponse) SetCode(v int64) *UpdateEcciInstanceResponse {
  s.Code = &v
  return s
}

func (s *UpdateEcciInstanceResponse) SetMsg(v string) *UpdateEcciInstanceResponse {
  s.Msg = &v
  return s
}

func (s *UpdateEcciInstanceResponse) SetRequestId(v string) *UpdateEcciInstanceResponse {
  s.RequestId = &v
  return s
}

type UpdateEcciInstancePaths struct {
  // {"en":"ecci instance name", "zh_CN":"ecci实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateEcciInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePaths) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePaths) SetName(v string) *UpdateEcciInstancePaths {
  s.Name = &v
  return s
}

type UpdateEcciInstanceParameters struct {
}

func (s UpdateEcciInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceParameters) GoString() string {
  return s.String()
}

type UpdateEcciInstanceRequestHeader struct {
}

func (s UpdateEcciInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceRequestHeader) GoString() string {
  return s.String()
}

type UpdateEcciInstanceResponseHeader struct {
}

func (s UpdateEcciInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceResponseHeader) GoString() string {
  return s.String()
}

type UpdateEcciInstanceInstanceDistribution struct {
  // {"en":"distribution cluster name", "zh_CN":"部署集群名称"}
  Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty" require:"true"`
  // {"en":"Number of desired pods", "zh_CN":"预期 Pod 的数量"}
  Replicas *int64 `json:"replicas,omitempty" xml:"replicas,omitempty" require:"true"`
  // {"en":"assigned ip attributes", "zh_CN":"分配ip属性"}
  AssignIps []*UpdateEcciInstanceAssignIp `json:"assignIps,omitempty" xml:"assignIps,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEcciInstanceInstanceDistribution) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceInstanceDistribution) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceInstanceDistribution) SetCluster(v string) *UpdateEcciInstanceInstanceDistribution {
  s.Cluster = &v
  return s
}

func (s *UpdateEcciInstanceInstanceDistribution) SetReplicas(v int64) *UpdateEcciInstanceInstanceDistribution {
  s.Replicas = &v
  return s
}

func (s *UpdateEcciInstanceInstanceDistribution) SetAssignIps(v []*UpdateEcciInstanceAssignIp) *UpdateEcciInstanceInstanceDistribution {
  s.AssignIps = v
  return s
}

type UpdateEcciInstanceAssignIp struct {
  // {"en":"ip operator id", "zh_CN":"ip所属运营商id"}
  IspId *int64 `json:"ispId,omitempty" xml:"ispId,omitempty"`
  // {"en":"Whether to assign ipv4 ip, ipv4 is assigned by default", "zh_CN":"是否分配ipv4的ip，默认分配ipv4"}
  Ipv4 *bool `json:"ipv4,omitempty" xml:"ipv4,omitempty"`
  // {"en":"Whether to assign ipv6 ip", "zh_CN":"是否分配ipv6的ip"}
  Ipv6 *bool `json:"ipv6,omitempty" xml:"ipv6,omitempty"`
}

func (s UpdateEcciInstanceAssignIp) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceAssignIp) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceAssignIp) SetIspId(v int64) *UpdateEcciInstanceAssignIp {
  s.IspId = &v
  return s
}

func (s *UpdateEcciInstanceAssignIp) SetIpv4(v bool) *UpdateEcciInstanceAssignIp {
  s.Ipv4 = &v
  return s
}

func (s *UpdateEcciInstanceAssignIp) SetIpv6(v bool) *UpdateEcciInstanceAssignIp {
  s.Ipv6 = &v
  return s
}

type UpdateEcciInstanceContainer struct {
  // {"en":"UpdateEcciInstanceContainer cpu limit, the sum of all container cpu limits cannot exceed the instance cpu limit", "zh_CN":"容器cpu限制,所有容器cpu限制总和不能超过实例cpu限制"}
  CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty" require:"true"`
  // {"en":"UpdateEcciInstanceContainer memory limit, the sum of all container memory limits cannot exceed the instance memory limit", "zh_CN":"容器内存限制,所有容器内存限制总和不能超过实例内存限制"}
  MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty" require:"true"`
  // {"en":"UpdateEcciInstanceContainer name", "zh_CN":"容器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"UpdateEcciInstanceContainer image", "zh_CN":"容器镜像"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en":"Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided.", "zh_CN":"入口点数组。不在 Shell 中执行。如果未提供，则使用容器镜像的 ENTRYPOINT"}
  Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
  // {"en":"Arguments to the entrypoint. The container image's CMD is used if this is not provided", "zh_CN":"entrypoint 的参数。如果未提供，则使用容器镜像的 CMD 设置"}
  Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
  // {"en":"UpdateEcciInstanceContainer's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image", "zh_CN":"容器的工作目录。如果未指定，将使用容器运行时的默认值，默认值可能在容器镜像中配置"}
  WorkingDir *string `json:"workingDir,omitempty" xml:"workingDir,omitempty"`
  // {"en":"ist of environment variables to set in the container", "zh_CN":"要在容器中设置的环境变量列表"}
  Env map[string]*string `json:"env,omitempty" xml:"env,omitempty"`
  // {"en":"downwardapi type sensitive environment variables, authorization required", "zh_CN":"downwardapi 类型敏感环境变量,需授权"}
  EnvRef []*UpdateEcciInstanceContainerEnvRef `json:"envRef,omitempty" xml:"envRef,omitempty" type:"Repeated"`
  // {"en":"Pod volumes to mount into the container's filesystem. Cannot be updated", "zh_CN":"要挂载到容器文件系统中的 Pod 卷。无法更新"}
  VolumeMounts []*UpdateEcciInstanceVolumeMount `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty" type:"Repeated"`
}

func (s UpdateEcciInstanceContainer) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceContainer) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceContainer) SetCpuLimit(v string) *UpdateEcciInstanceContainer {
  s.CpuLimit = &v
  return s
}

func (s *UpdateEcciInstanceContainer) SetMemoryLimit(v string) *UpdateEcciInstanceContainer {
  s.MemoryLimit = &v
  return s
}

func (s *UpdateEcciInstanceContainer) SetName(v string) *UpdateEcciInstanceContainer {
  s.Name = &v
  return s
}

func (s *UpdateEcciInstanceContainer) SetImage(v string) *UpdateEcciInstanceContainer {
  s.Image = &v
  return s
}

func (s *UpdateEcciInstanceContainer) SetCommand(v []*string) *UpdateEcciInstanceContainer {
  s.Command = v
  return s
}

func (s *UpdateEcciInstanceContainer) SetArgs(v []*string) *UpdateEcciInstanceContainer {
  s.Args = v
  return s
}

func (s *UpdateEcciInstanceContainer) SetWorkingDir(v string) *UpdateEcciInstanceContainer {
  s.WorkingDir = &v
  return s
}

func (s *UpdateEcciInstanceContainer) SetEnv(v map[string]*string) *UpdateEcciInstanceContainer {
  s.Env = v
  return s
}

func (s *UpdateEcciInstanceContainer) SetEnvRef(v []*UpdateEcciInstanceContainerEnvRef) *UpdateEcciInstanceContainer {
  s.EnvRef = v
  return s
}

func (s *UpdateEcciInstanceContainer) SetVolumeMounts(v []*UpdateEcciInstanceVolumeMount) *UpdateEcciInstanceContainer {
  s.VolumeMounts = v
  return s
}

type UpdateEcciInstanceContainerPort struct {
  // {"en":"If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services", "zh_CN":"如果设置此字段，这必须是 IANA_SVC_NAME 并且在 Pod 中唯一。 Pod 中的每个命名端口都必须具有唯一的名称。服务可以引用的端口的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536", "zh_CN":"要在 Pod 的 IP 地址上公开的端口号。这必须是有效的端口号，0 < x < 65536"}
  UpdateEcciInstanceContainerPort *int32 `json:"containerPort,omitempty" xml:"containerPort,omitempty"`
  // {"en":"Protocol for port. Must be UDP, TCP, or SCTP. Defaults to \"TCP\"", "zh_CN":"端口协议。必须是 UDP、TCP 或 SCTP。默认为 TCP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateEcciInstanceContainerPort) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceContainerPort) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceContainerPort) SetName(v string) *UpdateEcciInstanceContainerPort {
  s.Name = &v
  return s
}

func (s *UpdateEcciInstanceContainerPort) SetContainerPort(v int32) *UpdateEcciInstanceContainerPort {
  s.UpdateEcciInstanceContainerPort = &v
  return s
}

func (s *UpdateEcciInstanceContainerPort) SetProtocol(v string) *UpdateEcciInstanceContainerPort {
  s.Protocol = &v
  return s
}

type UpdateEcciInstancePodSpec struct {
  // {"en":"List of containers belonging to the pod. There must be at least one container in a Pod. ", "zh_CN":"属于 Pod 的容器列表。Pod 中必须至少有一个容器。"}
  Containers []*UpdateEcciInstanceContainer `json:"containers,omitempty" xml:"containers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always.", "zh_CN":"Pod 内所有容器的重启策略。Always、OnFailure、Never 之一。默认为 Always。"}
  RestartPolicy *string `json:"restartPolicy,omitempty" xml:"restartPolicy,omitempty" require:"true"`
  // {"en":"Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.", "zh_CN":"可选字段，表示 Pod 需要体面终止的所需的时长（以秒为单位）。字段值可以在删除请求中减少。 字段值必须是非负整数。零值表示收到 kill 信号则立即停止（没有机会关闭）。 如果此值为 nil，则将使用默认宽限期。 宽限期是从 Pod 中运行的进程收到终止信号后，到进程被 kill 信号强制停止之前，Pod 可以继续存在的时间（以秒为单位）。 应该将此值设置为比你的进程的预期清理时间更长。默认为 30 秒。"}
  TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" xml:"terminationGracePeriodSeconds,omitempty" require:"true"`
  // {"en":"List of volumes that can be mounted by containers belonging to the pod.", "zh_CN":"可以由属于 Pod 的容器挂载的卷列表。"}
  Volumes []*UpdateEcciInstancePodVolume `json:"volumes,omitempty" xml:"volumes,omitempty" type:"Repeated"`
  // {"en":"If specified, the pod's scheduling constraints", "zh_CN":"如果指定了，则作为 Pod 的调度约束。"}
  UpdateEcciInstanceAffinity *UpdateEcciInstanceAffinity `json:"affinity,omitempty" xml:"affinity,omitempty" require:"true"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
}

func (s UpdateEcciInstancePodSpec) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePodSpec) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePodSpec) SetContainers(v []*UpdateEcciInstanceContainer) *UpdateEcciInstancePodSpec {
  s.Containers = v
  return s
}

func (s *UpdateEcciInstancePodSpec) SetRestartPolicy(v string) *UpdateEcciInstancePodSpec {
  s.RestartPolicy = &v
  return s
}

func (s *UpdateEcciInstancePodSpec) SetTerminationGracePeriodSeconds(v int64) *UpdateEcciInstancePodSpec {
  s.TerminationGracePeriodSeconds = &v
  return s
}

func (s *UpdateEcciInstancePodSpec) SetVolumes(v []*UpdateEcciInstancePodVolume) *UpdateEcciInstancePodSpec {
  s.Volumes = v
  return s
}

func (s *UpdateEcciInstancePodSpec) SetAffinity(v *UpdateEcciInstanceAffinity) *UpdateEcciInstancePodSpec {
  s.UpdateEcciInstanceAffinity = v
  return s
}

func (s *UpdateEcciInstancePodSpec) SetAnnotations(v map[string]*string) *UpdateEcciInstancePodSpec {
  s.Annotations = v
  return s
}

func (s *UpdateEcciInstancePodSpec) SetLabels(v map[string]*string) *UpdateEcciInstancePodSpec {
  s.Labels = v
  return s
}

type UpdateEcciInstanceVolumeMount struct {
  // {"en":"This must match the Name of a Volume", "zh_CN":"此字段必须与卷的名称匹配"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Path within the container at which the volume should be mounted. Must not contain ':'", "zh_CN":"容器内卷的挂载路径。不得包含 ':'"}
  MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty" require:"true"`
  // {"en":"Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false", "zh_CN":"如果为 true，则以只读方式挂载，否则（false 或未设置）以读写方式挂载。默认为 false"}
  ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty" require:"true"`
}

func (s UpdateEcciInstanceVolumeMount) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceVolumeMount) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceVolumeMount) SetName(v string) *UpdateEcciInstanceVolumeMount {
  s.Name = &v
  return s
}

func (s *UpdateEcciInstanceVolumeMount) SetMountPath(v string) *UpdateEcciInstanceVolumeMount {
  s.MountPath = &v
  return s
}

func (s *UpdateEcciInstanceVolumeMount) SetReadOnly(v bool) *UpdateEcciInstanceVolumeMount {
  s.ReadOnly = &v
  return s
}

type UpdateEcciInstanceHostPathVolume struct {
  // {"en":"path of the directory on the host. If the path is a symlink, it will follow the link to the real path", "zh_CN":"目录在主机上的路径。如果该路径是一个符号链接，则它将沿着链接指向真实路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"type for HostPath Volume Defaults to ''", "zh_CN":"卷的类型。默认为 ''"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s UpdateEcciInstanceHostPathVolume) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceHostPathVolume) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceHostPathVolume) SetPath(v string) *UpdateEcciInstanceHostPathVolume {
  s.Path = &v
  return s
}

func (s *UpdateEcciInstanceHostPathVolume) SetType(v string) *UpdateEcciInstanceHostPathVolume {
  s.Type = &v
  return s
}

type UpdateEcciInstanceContainerEnvRef struct {
  // {"en":"Environment variable name", "zh_CN":"环境变量名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Path to the field to be selected by downwardapi", "zh_CN":"downwardapi要选择的字段的路径"}
  FieldPath *string `json:"fieldPath,omitempty" xml:"fieldPath,omitempty" require:"true"`
}

func (s UpdateEcciInstanceContainerEnvRef) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceContainerEnvRef) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceContainerEnvRef) SetName(v string) *UpdateEcciInstanceContainerEnvRef {
  s.Name = &v
  return s
}

func (s *UpdateEcciInstanceContainerEnvRef) SetFieldPath(v string) *UpdateEcciInstanceContainerEnvRef {
  s.FieldPath = &v
  return s
}

type UpdateEcciInstancePodVolume struct {
  // {"en":"name of the volume. Must be a DNS_LABEL and unique within the pod.", "zh_CN":"卷的名称。必须是 DNS_LABEL 且在 Pod 内是唯一的。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"represents a reference to a PersistentVolumeClaim in the same namespace", "zh_CN":"表示对同一名字空间中 PersistentVolumeClaim 的引用"}
  PersistentVolumeClaim *UpdateEcciInstancePodVolumeClaim `json:"persistentVolumeClaim,omitempty" xml:"persistentVolumeClaim,omitempty" require:"true"`
  // {"en":"Represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this.", "zh_CN":"表示主机上预先存在的文件或目录，它们将被直接暴露给容器。 这种卷通常用于系统代理或允许查看主机的其他特权操作。大多数容器不需要这种卷。"}
  HostPath *UpdateEcciInstanceHostPathVolume `json:"hostPath,omitempty" xml:"hostPath,omitempty" require:"true"`
}

func (s UpdateEcciInstancePodVolume) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePodVolume) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePodVolume) SetName(v string) *UpdateEcciInstancePodVolume {
  s.Name = &v
  return s
}

func (s *UpdateEcciInstancePodVolume) SetPersistentVolumeClaim(v *UpdateEcciInstancePodVolumeClaim) *UpdateEcciInstancePodVolume {
  s.PersistentVolumeClaim = v
  return s
}

func (s *UpdateEcciInstancePodVolume) SetHostPath(v *UpdateEcciInstanceHostPathVolume) *UpdateEcciInstancePodVolume {
  s.HostPath = v
  return s
}

type UpdateEcciInstancePodVolumeClaim struct {
  // {"en":"the name of a PersistentVolumeClaim in the same namespace as the pod using this volume", "zh_CN":"与使用此卷的 Pod 位于同一名字空间中的 PersistentVolumeClaim 的名称"}
  ClaimName *string `json:"claimName,omitempty" xml:"claimName,omitempty" require:"true"`
}

func (s UpdateEcciInstancePodVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePodVolumeClaim) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePodVolumeClaim) SetClaimName(v string) *UpdateEcciInstancePodVolumeClaim {
  s.ClaimName = &v
  return s
}

type UpdateEcciInstanceAffinity struct {
  // {"en":"A group of inter pod affinity scheduling rules.", "zh_CN":"一组 Pod 间亲和性调度规则。"}
  UpdateEcciInstancePodAffinity *UpdateEcciInstancePodAffinity `json:"podAffinity,omitempty" xml:"podAffinity,omitempty" require:"true"`
  // {"en":"A group of node affinity scheduling rules.", "zh_CN":"一组节点亲和性调度规则。"}
  UpdateEcciInstancePodAntiAffinity *UpdateEcciInstancePodAntiAffinity `json:"podAntiAffinity,omitempty" xml:"podAntiAffinity,omitempty" require:"true"`
}

func (s UpdateEcciInstanceAffinity) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceAffinity) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceAffinity) SetPodAffinity(v *UpdateEcciInstancePodAffinity) *UpdateEcciInstanceAffinity {
  s.UpdateEcciInstancePodAffinity = v
  return s
}

func (s *UpdateEcciInstanceAffinity) SetPodAntiAffinity(v *UpdateEcciInstancePodAntiAffinity) *UpdateEcciInstanceAffinity {
  s.UpdateEcciInstancePodAntiAffinity = v
  return s
}

type UpdateEcciInstancePodAntiAffinity struct {
  // {"en":"If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.", "zh_CN":"如果在调度时不满足该字段指定的反亲和性要求，则该 Pod 不会被调度到该节点上。 如果在 Pod 执行期间的某个时间点不再满足此字段指定的反亲和性要求（例如：由于 Pod 标签更新）， 系统可能会或可能不会尝试最终将 Pod 从其节点中逐出。 当有多个元素时，每个 podAffinityTerm 对应的节点列表是取其交集的，即必须满足所有条件。"}
  RequiredDuringSchedulingIgnoredDuringExecution []*UpdateEcciInstancePodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
  // {"en":"The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.", "zh_CN":"调度器更倾向于将 Pod 调度到满足该字段指定的反亲和性表达式的节点， 但它可能会选择违反一个或多个表达式的节点。 最优选的节点是权重总和最大的节点，即对于满足所有调度要求（资源请求、requiredDuringScheduling 反亲和性表达式等）的每个节点，通过遍历元素来计算总和如果节点具有与相应 podAffinityTerm 匹配的 Pod，则此字段并在总和中添加\"权重\"；具有最高加和的节点是最优选的。"}
  PreferredDuringSchedulingIgnoredDuringExecution []*UpdateEcciInstanceWeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEcciInstancePodAntiAffinity) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePodAntiAffinity) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePodAntiAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []*UpdateEcciInstancePodAffinityTerm) *UpdateEcciInstancePodAntiAffinity {
  s.RequiredDuringSchedulingIgnoredDuringExecution = v
  return s
}

func (s *UpdateEcciInstancePodAntiAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []*UpdateEcciInstanceWeightedPodAffinityTerm) *UpdateEcciInstancePodAntiAffinity {
  s.PreferredDuringSchedulingIgnoredDuringExecution = v
  return s
}

type UpdateEcciInstancePodAffinity struct {
  // {"en":"If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.", "zh_CN":"如果在调度时不满足该字段指定的亲和性要求，则该 Pod 不会被调度到该节点上。 如果在 Pod 执行期间的某个时间点不再满足此字段指定的亲和性要求（例如：由于 Pod 标签更新）， 系统可能会也可能不会尝试最终将 Pod 从其节点中逐出。 当此列表中有多个元素时，每个 podAffinityTerm 对应的节点列表是取其交集的，即必须满足所有条件。"}
  RequiredDuringSchedulingIgnoredDuringExecution []*UpdateEcciInstancePodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
  // {"en":"The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.", "zh_CN":"调度器会更倾向于将 Pod 调度到满足该字段指定的亲和性表达式的节点， 但它可能会选择违反一个或多个表达式的节点。最优选择是权重总和最大的节点， 即对于满足所有调度要求（资源请求、requiredDuringScheduling 亲和表达式等）的每个节点， 通过迭代该字段的元素来计算总和，如果节点具有与相应 podAffinityTerm 匹配的 Pod，则将“权重”添加到总和中； 具有最高总和的节点是最优选的。"}
  PreferredDuringSchedulingIgnoredDuringExecution []*UpdateEcciInstanceWeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEcciInstancePodAffinity) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePodAffinity) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePodAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []*UpdateEcciInstancePodAffinityTerm) *UpdateEcciInstancePodAffinity {
  s.RequiredDuringSchedulingIgnoredDuringExecution = v
  return s
}

func (s *UpdateEcciInstancePodAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []*UpdateEcciInstanceWeightedPodAffinityTerm) *UpdateEcciInstancePodAffinity {
  s.PreferredDuringSchedulingIgnoredDuringExecution = v
  return s
}

type UpdateEcciInstancePodAffinityTerm struct {
  // {"en":"A label query over a set of resources, in this case pods.", "zh_CN":"对一组资源的标签查询，在这里资源为 Pod。"}
  UpdateEcciInstanceLabelSelector *UpdateEcciInstanceLabelSelector `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" require:"true"`
  // {"en":"pod's namespace", "zh_CN":" Pod 的名字空间"}
  Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" require:"true" type:"Repeated"`
  // {"en":"This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.", "zh_CN":"此 Pod 应与指定名字空间中与标签选择算符匹配的 Pod 集合位于同一位置（亲和性） 或位于不同位置（反亲和性），这里的“在同一位置”意味着运行在一个节点上，其键名为 topologyKey 的标签值与运行所选 Pod 集合中的某 Pod 的任何节点上的标签值匹配。 不允许使用空的 topologyKey。"}
  TopologyKey *string `json:"topologyKey,omitempty" xml:"topologyKey,omitempty" require:"true"`
}

func (s UpdateEcciInstancePodAffinityTerm) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstancePodAffinityTerm) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstancePodAffinityTerm) SetLabelSelector(v *UpdateEcciInstanceLabelSelector) *UpdateEcciInstancePodAffinityTerm {
  s.UpdateEcciInstanceLabelSelector = v
  return s
}

func (s *UpdateEcciInstancePodAffinityTerm) SetNamespaces(v []*string) *UpdateEcciInstancePodAffinityTerm {
  s.Namespaces = v
  return s
}

func (s *UpdateEcciInstancePodAffinityTerm) SetTopologyKey(v string) *UpdateEcciInstancePodAffinityTerm {
  s.TopologyKey = &v
  return s
}

type UpdateEcciInstanceLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty" require:"true"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*UpdateEcciInstanceLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEcciInstanceLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceLabelSelector) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceLabelSelector) SetMatchLabels(v map[string]*string) *UpdateEcciInstanceLabelSelector {
  s.MatchLabels = v
  return s
}

func (s *UpdateEcciInstanceLabelSelector) SetMatchExpressions(v []*UpdateEcciInstanceLabelSelectorRequirement) *UpdateEcciInstanceLabelSelector {
  s.MatchExpressions = v
  return s
}

type UpdateEcciInstanceLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" require:"true" type:"Repeated"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
}

func (s UpdateEcciInstanceLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceLabelSelectorRequirement) SetKey(v string) *UpdateEcciInstanceLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *UpdateEcciInstanceLabelSelectorRequirement) SetValues(v []*string) *UpdateEcciInstanceLabelSelectorRequirement {
  s.Values = v
  return s
}

func (s *UpdateEcciInstanceLabelSelectorRequirement) SetOperator(v string) *UpdateEcciInstanceLabelSelectorRequirement {
  s.Operator = &v
  return s
}

type UpdateEcciInstanceWeightedPodAffinityTerm struct {
  // {"en":"associated with matching the corresponding podAffinityTerm, in the range 1-100.", "zh_CN":"匹配相应 podAffinityTerm 条件的权重，范围为 1-100。"}
  Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"Required. A pod affinity term, associated with the corresponding weight.", "zh_CN":"必需的字段。一个 Pod 亲和性条件，对应一个与相应的权重值。"}
  UpdateEcciInstancePodAffinityTerm *UpdateEcciInstancePodAffinityTerm `json:"podAffinityTerm,omitempty" xml:"podAffinityTerm,omitempty" require:"true"`
}

func (s UpdateEcciInstanceWeightedPodAffinityTerm) String() string {
  return tea.Prettify(s)
}

func (s UpdateEcciInstanceWeightedPodAffinityTerm) GoString() string {
  return s.String()
}

func (s *UpdateEcciInstanceWeightedPodAffinityTerm) SetWeight(v int32) *UpdateEcciInstanceWeightedPodAffinityTerm {
  s.Weight = &v
  return s
}

func (s *UpdateEcciInstanceWeightedPodAffinityTerm) SetPodAffinityTerm(v *UpdateEcciInstancePodAffinityTerm) *UpdateEcciInstanceWeightedPodAffinityTerm {
  s.UpdateEcciInstancePodAffinityTerm = v
  return s
}




type GetEcciInstanceRequest struct {
}

func (s GetEcciInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceRequest) GoString() string {
  return s.String()
}

type GetEcciInstanceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"instanceDetail", "zh_CN":"实例详情"}
  Data *GetEcciInstanceEcciInstanceDetail `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetEcciInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceResponse) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceResponse) SetCode(v int64) *GetEcciInstanceResponse {
  s.Code = &v
  return s
}

func (s *GetEcciInstanceResponse) SetMsg(v string) *GetEcciInstanceResponse {
  s.Msg = &v
  return s
}

func (s *GetEcciInstanceResponse) SetRequestId(v string) *GetEcciInstanceResponse {
  s.RequestId = &v
  return s
}

func (s *GetEcciInstanceResponse) SetData(v *GetEcciInstanceEcciInstanceDetail) *GetEcciInstanceResponse {
  s.Data = v
  return s
}

type GetEcciInstancePaths struct {
  // {"en":"instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetEcciInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePaths) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePaths) SetName(v string) *GetEcciInstancePaths {
  s.Name = &v
  return s
}

type GetEcciInstanceParameters struct {
}

func (s GetEcciInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceParameters) GoString() string {
  return s.String()
}

type GetEcciInstanceRequestHeader struct {
}

func (s GetEcciInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceRequestHeader) GoString() string {
  return s.String()
}

type GetEcciInstanceResponseHeader struct {
}

func (s GetEcciInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceResponseHeader) GoString() string {
  return s.String()
}

type GetEcciInstanceContainer struct {
  // {"en":"GetEcciInstanceContainer cpu limit, the sum of all container cpu limits cannot exceed the instance cpu limit", "zh_CN":"容器cpu限制,所有容器cpu限制总和不能超过实例cpu限制"}
  CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty" require:"true"`
  // {"en":"GetEcciInstanceContainer memory limit, the sum of all container memory limits cannot exceed the instance memory limit", "zh_CN":"容器内存限制,所有容器内存限制总和不能超过实例内存限制"}
  MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty" require:"true"`
  // {"en":"GetEcciInstanceContainer name", "zh_CN":"容器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"GetEcciInstanceContainer image", "zh_CN":"容器镜像"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en":"Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided.", "zh_CN":"入口点数组。不在 Shell 中执行。如果未提供，则使用容器镜像的 ENTRYPOINT"}
  Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
  // {"en":"Arguments to the entrypoint. The container image's CMD is used if this is not provided", "zh_CN":"entrypoint 的参数。如果未提供，则使用容器镜像的 CMD 设置"}
  Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
  // {"en":"GetEcciInstanceContainer's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image", "zh_CN":"容器的工作目录。如果未指定，将使用容器运行时的默认值，默认值可能在容器镜像中配置"}
  WorkingDir *string `json:"workingDir,omitempty" xml:"workingDir,omitempty"`
  // {"en":"ist of environment variables to set in the container", "zh_CN":"要在容器中设置的环境变量列表"}
  Env map[string]*string `json:"env,omitempty" xml:"env,omitempty"`
  // {"en":"downwardapi type sensitive environment variables, authorization required", "zh_CN":"downwardapi 类型敏感环境变量,需授权"}
  EnvRef []*GetEcciInstanceContainerEnvRef `json:"envRef,omitempty" xml:"envRef,omitempty" type:"Repeated"`
  // {"en":"Pod volumes to mount into the container's filesystem. Cannot be updated", "zh_CN":"要挂载到容器文件系统中的 Pod 卷。无法更新"}
  VolumeMounts []*GetEcciInstanceVolumeMount `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty" type:"Repeated"`
}

func (s GetEcciInstanceContainer) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceContainer) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceContainer) SetCpuLimit(v string) *GetEcciInstanceContainer {
  s.CpuLimit = &v
  return s
}

func (s *GetEcciInstanceContainer) SetMemoryLimit(v string) *GetEcciInstanceContainer {
  s.MemoryLimit = &v
  return s
}

func (s *GetEcciInstanceContainer) SetName(v string) *GetEcciInstanceContainer {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceContainer) SetImage(v string) *GetEcciInstanceContainer {
  s.Image = &v
  return s
}

func (s *GetEcciInstanceContainer) SetCommand(v []*string) *GetEcciInstanceContainer {
  s.Command = v
  return s
}

func (s *GetEcciInstanceContainer) SetArgs(v []*string) *GetEcciInstanceContainer {
  s.Args = v
  return s
}

func (s *GetEcciInstanceContainer) SetWorkingDir(v string) *GetEcciInstanceContainer {
  s.WorkingDir = &v
  return s
}

func (s *GetEcciInstanceContainer) SetEnv(v map[string]*string) *GetEcciInstanceContainer {
  s.Env = v
  return s
}

func (s *GetEcciInstanceContainer) SetEnvRef(v []*GetEcciInstanceContainerEnvRef) *GetEcciInstanceContainer {
  s.EnvRef = v
  return s
}

func (s *GetEcciInstanceContainer) SetVolumeMounts(v []*GetEcciInstanceVolumeMount) *GetEcciInstanceContainer {
  s.VolumeMounts = v
  return s
}

type GetEcciInstancePodSpec struct {
  // {"en":"List of containers belonging to the pod. There must be at least one container in a Pod. ", "zh_CN":"属于 Pod 的容器列表。Pod 中必须至少有一个容器。"}
  Containers []*GetEcciInstanceContainer `json:"containers,omitempty" xml:"containers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always.", "zh_CN":"Pod 内所有容器的重启策略。Always、OnFailure、Never 之一。默认为 Always。"}
  RestartPolicy *string `json:"restartPolicy,omitempty" xml:"restartPolicy,omitempty" require:"true"`
  // {"en":"Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.", "zh_CN":"可选字段，表示 Pod 需要体面终止的所需的时长（以秒为单位）。字段值可以在删除请求中减少。 字段值必须是非负整数。零值表示收到 kill 信号则立即停止（没有机会关闭）。 如果此值为 nil，则将使用默认宽限期。 宽限期是从 Pod 中运行的进程收到终止信号后，到进程被 kill 信号强制停止之前，Pod 可以继续存在的时间（以秒为单位）。 应该将此值设置为比你的进程的预期清理时间更长。默认为 30 秒。"}
  TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" xml:"terminationGracePeriodSeconds,omitempty" require:"true"`
  // {"en":"List of volumes that can be mounted by containers belonging to the pod.", "zh_CN":"可以由属于 Pod 的容器挂载的卷列表。"}
  Volumes []*GetEcciInstancePodVolume `json:"volumes,omitempty" xml:"volumes,omitempty" type:"Repeated"`
  // {"en":"If specified, the pod's scheduling constraints", "zh_CN":"如果指定了，则作为 Pod 的调度约束。"}
  GetEcciInstanceAffinity *GetEcciInstanceAffinity `json:"affinity,omitempty" xml:"affinity,omitempty" require:"true"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Whether to use kata runtime, Use kata by default", "zh_CN":"是否使用kata运行时,默认使用kata"}
  KataRuntime *bool `json:"kataRuntime,omitempty" xml:"kataRuntime,omitempty"`
}

func (s GetEcciInstancePodSpec) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePodSpec) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePodSpec) SetContainers(v []*GetEcciInstanceContainer) *GetEcciInstancePodSpec {
  s.Containers = v
  return s
}

func (s *GetEcciInstancePodSpec) SetRestartPolicy(v string) *GetEcciInstancePodSpec {
  s.RestartPolicy = &v
  return s
}

func (s *GetEcciInstancePodSpec) SetTerminationGracePeriodSeconds(v int64) *GetEcciInstancePodSpec {
  s.TerminationGracePeriodSeconds = &v
  return s
}

func (s *GetEcciInstancePodSpec) SetVolumes(v []*GetEcciInstancePodVolume) *GetEcciInstancePodSpec {
  s.Volumes = v
  return s
}

func (s *GetEcciInstancePodSpec) SetAffinity(v *GetEcciInstanceAffinity) *GetEcciInstancePodSpec {
  s.GetEcciInstanceAffinity = v
  return s
}

func (s *GetEcciInstancePodSpec) SetAnnotations(v map[string]*string) *GetEcciInstancePodSpec {
  s.Annotations = v
  return s
}

func (s *GetEcciInstancePodSpec) SetLabels(v map[string]*string) *GetEcciInstancePodSpec {
  s.Labels = v
  return s
}

func (s *GetEcciInstancePodSpec) SetKataRuntime(v bool) *GetEcciInstancePodSpec {
  s.KataRuntime = &v
  return s
}

type GetEcciInstanceEcciInstanceDetail struct {
  // {"en":"Instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance display name", "zh_CN":"实例展示名称"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"private ip list", "zh_CN":"内网ip列表"}
  PrivateIps []*string `json:"privateIps,omitempty" xml:"privateIps,omitempty" require:"true" type:"Repeated"`
  // {"en":"cluster name", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty" require:"true"`
  // {"en":"instance cpu limit", "zh_CN":"实例cpu限制"}
  CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty" require:"true"`
  // {"en":"instance memory limit", "zh_CN":"实例内存限制"}
  MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty" require:"true"`
  // {"en":"Minimum amount of cpu resources required", "zh_CN":"cpu 要求最小资源量"}
  CpuRequest *string `json:"cpuRequest,omitempty" xml:"cpuRequest,omitempty" require:"true"`
  // {"en":"Minimum amount of memory resources required", "zh_CN":"内存要求最小资源量"}
  MemoryRequest *string `json:"memoryRequest,omitempty" xml:"memoryRequest,omitempty" require:"true"`
  // {"en":"pod status", "zh_CN":"pod状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Number of containers ", "zh_CN":"容器数量"}
  Containers *int64 `json:"containers,omitempty" xml:"containers,omitempty" require:"true"`
  // {"en":"Number of running containers ", "zh_CN":"运行中容器数量"}
  RunningContainers *int64 `json:"runningContainers,omitempty" xml:"runningContainers,omitempty" require:"true"`
  // {"en":"description ", "zh_CN":"描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Specification of the desired behavior of the pod.", "zh_CN":"Pod 预期行为的规约。"}
  GetEcciInstancePodSpec *GetEcciInstancePodSpec `json:"podSpec,omitempty" xml:"podSpec,omitempty" require:"true"`
  // {"en":"create time ", "zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"update time ", "zh_CN":"更新时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"area", "zh_CN":"区域"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"country", "zh_CN":"国家"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"city", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"Whether to assign public IP", "zh_CN":"是否分配公网ip"}
  AllocatePublicIp *bool `json:"allocatePublicIp,omitempty" xml:"allocatePublicIp,omitempty" require:"true"`
  // {"en":"IP operator list", "zh_CN":"ip运营商列表"}
  Isps []*GetEcciInstanceIspIp `json:"isps,omitempty" xml:"isps,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether to enable fault scheduling", "zh_CN":"是否开启故障调度"}
  Failover *bool `json:"failover,omitempty" xml:"failover,omitempty" require:"true"`
  // {"en":"is colocation resource or not", "zh_CN":"是否在离线混部资源"}
  Colocation *bool `json:"colocation,omitempty" xml:"colocation,omitempty" require:"true"`
  // {"en":"the cluster behavior", "zh_CN":"集群策略配置"}
  GetEcciInstanceClusterBehavior *GetEcciInstanceClusterBehavior `json:"clusterBehavior,omitempty" xml:"clusterBehavior,omitempty" require:"true"`
}

func (s GetEcciInstanceEcciInstanceDetail) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceEcciInstanceDetail) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceEcciInstanceDetail) SetName(v string) *GetEcciInstanceEcciInstanceDetail {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetDisplayName(v string) *GetEcciInstanceEcciInstanceDetail {
  s.DisplayName = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetPrivateIps(v []*string) *GetEcciInstanceEcciInstanceDetail {
  s.PrivateIps = v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetClusterName(v string) *GetEcciInstanceEcciInstanceDetail {
  s.ClusterName = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetCpuLimit(v string) *GetEcciInstanceEcciInstanceDetail {
  s.CpuLimit = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetMemoryLimit(v string) *GetEcciInstanceEcciInstanceDetail {
  s.MemoryLimit = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetCpuRequest(v string) *GetEcciInstanceEcciInstanceDetail {
  s.CpuRequest = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetMemoryRequest(v string) *GetEcciInstanceEcciInstanceDetail {
  s.MemoryRequest = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetStatus(v string) *GetEcciInstanceEcciInstanceDetail {
  s.Status = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetContainers(v int64) *GetEcciInstanceEcciInstanceDetail {
  s.Containers = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetRunningContainers(v int64) *GetEcciInstanceEcciInstanceDetail {
  s.RunningContainers = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetDescription(v string) *GetEcciInstanceEcciInstanceDetail {
  s.Description = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetPodSpec(v *GetEcciInstancePodSpec) *GetEcciInstanceEcciInstanceDetail {
  s.GetEcciInstancePodSpec = v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetCreateTime(v string) *GetEcciInstanceEcciInstanceDetail {
  s.CreateTime = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetUpdateTime(v string) *GetEcciInstanceEcciInstanceDetail {
  s.UpdateTime = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetArea(v string) *GetEcciInstanceEcciInstanceDetail {
  s.Area = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetCountry(v string) *GetEcciInstanceEcciInstanceDetail {
  s.Country = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetProvince(v string) *GetEcciInstanceEcciInstanceDetail {
  s.Province = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetCity(v string) *GetEcciInstanceEcciInstanceDetail {
  s.City = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetAllocatePublicIp(v bool) *GetEcciInstanceEcciInstanceDetail {
  s.AllocatePublicIp = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetIsps(v []*GetEcciInstanceIspIp) *GetEcciInstanceEcciInstanceDetail {
  s.Isps = v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetFailover(v bool) *GetEcciInstanceEcciInstanceDetail {
  s.Failover = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetColocation(v bool) *GetEcciInstanceEcciInstanceDetail {
  s.Colocation = &v
  return s
}

func (s *GetEcciInstanceEcciInstanceDetail) SetClusterBehavior(v *GetEcciInstanceClusterBehavior) *GetEcciInstanceEcciInstanceDetail {
  s.GetEcciInstanceClusterBehavior = v
  return s
}

type GetEcciInstanceClusterBehavior struct {
  // {"en":"a list of the scheduling constraints", "zh_CN":"调度策略列表"}
  GetEcciInstanceSpreadConstraint []*GetEcciInstanceSpreadConstraint `json:"spreadConstraint,omitempty" xml:"spreadConstraint,omitempty" require:"true" type:"Repeated"`
  // {"en":"cluster faileover", "zh_CN":"集群重调度"}
  GetEcciInstanceClusterFailover *GetEcciInstanceClusterFailover `json:"clusterFailover,omitempty" xml:"clusterFailover,omitempty" require:"true"`
}

func (s GetEcciInstanceClusterBehavior) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceClusterBehavior) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceClusterBehavior) SetSpreadConstraint(v []*GetEcciInstanceSpreadConstraint) *GetEcciInstanceClusterBehavior {
  s.GetEcciInstanceSpreadConstraint = v
  return s
}

func (s *GetEcciInstanceClusterBehavior) SetClusterFailover(v *GetEcciInstanceClusterFailover) *GetEcciInstanceClusterBehavior {
  s.GetEcciInstanceClusterFailover = v
  return s
}

type GetEcciInstanceSpreadConstraint struct {
  // {"en":"spread field support: group or cluster", "zh_CN":"分组策略,当前支持: group, cluster"}
  SpreadByField *string `json:"spreadByField,omitempty" xml:"spreadByField,omitempty" require:"true"`
  // {"en":"max groups in field. by default 1", "zh_CN":"策略需要的最大分组个数.默认: 1"}
  MaxGroups *int32 `json:"maxGroups,omitempty" xml:"maxGroups,omitempty" require:"true"`
  // {"en":"min groups in field. by default 1", "zh_CN":"策略需要的最小分组个数.默认: 1"}
  MinGroups *int32 `json:"minGroups,omitempty" xml:"minGroups,omitempty" require:"true"`
}

func (s GetEcciInstanceSpreadConstraint) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceSpreadConstraint) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceSpreadConstraint) SetSpreadByField(v string) *GetEcciInstanceSpreadConstraint {
  s.SpreadByField = &v
  return s
}

func (s *GetEcciInstanceSpreadConstraint) SetMaxGroups(v int32) *GetEcciInstanceSpreadConstraint {
  s.MaxGroups = &v
  return s
}

func (s *GetEcciInstanceSpreadConstraint) SetMinGroups(v int32) *GetEcciInstanceSpreadConstraint {
  s.MinGroups = &v
  return s
}

type GetEcciInstanceClusterFailover struct {
  // {"en":"the cluser faileover. 0: off, 1: on. default 0", "zh_CN":"集群故障转移开关.0: 关闭,1打开. 默认:0 "}
  GetEcciInstanceClusterFailover *int32 `json:"clusterFailover,omitempty" xml:"clusterFailover,omitempty" require:"true"`
  // {"en":"the tolerationSeconds of the workload unhealthy. default 300", "zh_CN":"应用创建后多久时间没有runing,超过认为不健康.默认300秒"}
  TolerationSeconds *int32 `json:"tolerationSeconds,omitempty" xml:"tolerationSeconds,omitempty" require:"true"`
  // {"en":"PurgeMode represents how to deal with the legacy applications on the cluster from which the application is migrated.only support Graciously", "zh_CN":"应用漂移后,旧应用的删除模式.只支持 Graciously 平滑删除"}
  PurgeMode *string `json:"purgeMode,omitempty" xml:"purgeMode,omitempty" require:"true"`
  // {"en":"GracePeriodSeconds is the maximum waiting duration in seconds before application on the migrated cluster should be deleted. default:300", "zh_CN":"平滑删除时间. 默认:300秒 "}
  GracePeriodSeconds *int32 `json:"gracePeriodSeconds,omitempty" xml:"gracePeriodSeconds,omitempty" require:"true"`
}

func (s GetEcciInstanceClusterFailover) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceClusterFailover) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceClusterFailover) SetClusterFailover(v int32) *GetEcciInstanceClusterFailover {
  s.GetEcciInstanceClusterFailover = &v
  return s
}

func (s *GetEcciInstanceClusterFailover) SetTolerationSeconds(v int32) *GetEcciInstanceClusterFailover {
  s.TolerationSeconds = &v
  return s
}

func (s *GetEcciInstanceClusterFailover) SetPurgeMode(v string) *GetEcciInstanceClusterFailover {
  s.PurgeMode = &v
  return s
}

func (s *GetEcciInstanceClusterFailover) SetGracePeriodSeconds(v int32) *GetEcciInstanceClusterFailover {
  s.GracePeriodSeconds = &v
  return s
}

type GetEcciInstanceIspIp struct {
  // {"en":"operator id", "zh_CN":"运营商id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Operator Chinese name", "zh_CN":"运营商中文名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Operator english name", "zh_CN":"运营商英文名"}
  NameEn *string `json:"nameEn,omitempty" xml:"nameEn,omitempty" require:"true"`
  // {"en":"ipv4 ip", "zh_CN":"ipv4 ip"}
  Ipv4 *string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true"`
  // {"en":"ipv6 ip", "zh_CN":"ipv6 ip"}
  Ipv6 *string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true"`
}

func (s GetEcciInstanceIspIp) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceIspIp) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceIspIp) SetId(v int64) *GetEcciInstanceIspIp {
  s.Id = &v
  return s
}

func (s *GetEcciInstanceIspIp) SetName(v string) *GetEcciInstanceIspIp {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceIspIp) SetNameEn(v string) *GetEcciInstanceIspIp {
  s.NameEn = &v
  return s
}

func (s *GetEcciInstanceIspIp) SetIpv4(v string) *GetEcciInstanceIspIp {
  s.Ipv4 = &v
  return s
}

func (s *GetEcciInstanceIspIp) SetIpv6(v string) *GetEcciInstanceIspIp {
  s.Ipv6 = &v
  return s
}

type GetEcciInstancePodVolume struct {
  // {"en":"name of the volume. Must be a DNS_LABEL and unique within the pod.", "zh_CN":"卷的名称。必须是 DNS_LABEL 且在 Pod 内是唯一的。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"represents a reference to a PersistentVolumeClaim in the same namespace", "zh_CN":"表示对同一名字空间中 PersistentVolumeClaim 的引用"}
  PersistentVolumeClaim *GetEcciInstancePodVolumeClaim `json:"persistentVolumeClaim,omitempty" xml:"persistentVolumeClaim,omitempty" require:"true"`
  // {"en":"Represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this.", "zh_CN":"表示主机上预先存在的文件或目录，它们将被直接暴露给容器。 这种卷通常用于系统代理或允许查看主机的其他特权操作。大多数容器不需要这种卷。"}
  HostPath *GetEcciInstanceHostPathVolume `json:"hostPath,omitempty" xml:"hostPath,omitempty" require:"true"`
}

func (s GetEcciInstancePodVolume) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePodVolume) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePodVolume) SetName(v string) *GetEcciInstancePodVolume {
  s.Name = &v
  return s
}

func (s *GetEcciInstancePodVolume) SetPersistentVolumeClaim(v *GetEcciInstancePodVolumeClaim) *GetEcciInstancePodVolume {
  s.PersistentVolumeClaim = v
  return s
}

func (s *GetEcciInstancePodVolume) SetHostPath(v *GetEcciInstanceHostPathVolume) *GetEcciInstancePodVolume {
  s.HostPath = v
  return s
}

type GetEcciInstanceHostPathVolume struct {
  // {"en":"path of the directory on the host. If the path is a symlink, it will follow the link to the real path", "zh_CN":"目录在主机上的路径。如果该路径是一个符号链接，则它将沿着链接指向真实路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"type for HostPath Volume Defaults to ''", "zh_CN":"卷的类型。默认为 ''"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s GetEcciInstanceHostPathVolume) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceHostPathVolume) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceHostPathVolume) SetPath(v string) *GetEcciInstanceHostPathVolume {
  s.Path = &v
  return s
}

func (s *GetEcciInstanceHostPathVolume) SetType(v string) *GetEcciInstanceHostPathVolume {
  s.Type = &v
  return s
}

type GetEcciInstanceContainerEnvRef struct {
  // {"en":"Environment variable name", "zh_CN":"环境变量名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Path to the field to be selected by downwardapi", "zh_CN":"downwardapi要选择的字段的路径"}
  FieldPath *string `json:"fieldPath,omitempty" xml:"fieldPath,omitempty" require:"true"`
}

func (s GetEcciInstanceContainerEnvRef) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceContainerEnvRef) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceContainerEnvRef) SetName(v string) *GetEcciInstanceContainerEnvRef {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceContainerEnvRef) SetFieldPath(v string) *GetEcciInstanceContainerEnvRef {
  s.FieldPath = &v
  return s
}

type GetEcciInstancePodVolumeClaim struct {
  // {"en":"the name of a PersistentVolumeClaim in the same namespace as the pod using this volume", "zh_CN":"与使用此卷的 Pod 位于同一名字空间中的 PersistentVolumeClaim 的名称"}
  ClaimName *string `json:"claimName,omitempty" xml:"claimName,omitempty" require:"true"`
}

func (s GetEcciInstancePodVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePodVolumeClaim) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePodVolumeClaim) SetClaimName(v string) *GetEcciInstancePodVolumeClaim {
  s.ClaimName = &v
  return s
}

type GetEcciInstanceVolumeMount struct {
  // {"en":"This must match the Name of a Volume", "zh_CN":"此字段必须与卷的名称匹配"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Path within the container at which the volume should be mounted. Must not contain ':'", "zh_CN":"容器内卷的挂载路径。不得包含 ':'"}
  MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty" require:"true"`
  // {"en":"Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false", "zh_CN":"如果为 true，则以只读方式挂载，否则（false 或未设置）以读写方式挂载。默认为 false"}
  ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty" require:"true"`
}

func (s GetEcciInstanceVolumeMount) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceVolumeMount) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceVolumeMount) SetName(v string) *GetEcciInstanceVolumeMount {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceVolumeMount) SetMountPath(v string) *GetEcciInstanceVolumeMount {
  s.MountPath = &v
  return s
}

func (s *GetEcciInstanceVolumeMount) SetReadOnly(v bool) *GetEcciInstanceVolumeMount {
  s.ReadOnly = &v
  return s
}

type GetEcciInstanceAffinity struct {
  // {"en":"A group of inter pod affinity scheduling rules.", "zh_CN":"一组 Pod 间亲和性调度规则。"}
  GetEcciInstancePodAffinity *GetEcciInstancePodAffinity `json:"podAffinity,omitempty" xml:"podAffinity,omitempty" require:"true"`
  // {"en":"A group of node affinity scheduling rules.", "zh_CN":"一组节点亲和性调度规则。"}
  GetEcciInstancePodAntiAffinity *GetEcciInstancePodAntiAffinity `json:"podAntiAffinity,omitempty" xml:"podAntiAffinity,omitempty" require:"true"`
}

func (s GetEcciInstanceAffinity) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceAffinity) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceAffinity) SetPodAffinity(v *GetEcciInstancePodAffinity) *GetEcciInstanceAffinity {
  s.GetEcciInstancePodAffinity = v
  return s
}

func (s *GetEcciInstanceAffinity) SetPodAntiAffinity(v *GetEcciInstancePodAntiAffinity) *GetEcciInstanceAffinity {
  s.GetEcciInstancePodAntiAffinity = v
  return s
}

type GetEcciInstancePodAntiAffinity struct {
  // {"en":"If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.", "zh_CN":"如果在调度时不满足该字段指定的反亲和性要求，则该 Pod 不会被调度到该节点上。 如果在 Pod 执行期间的某个时间点不再满足此字段指定的反亲和性要求（例如：由于 Pod 标签更新）， 系统可能会或可能不会尝试最终将 Pod 从其节点中逐出。 当有多个元素时，每个 podAffinityTerm 对应的节点列表是取其交集的，即必须满足所有条件。"}
  RequiredDuringSchedulingIgnoredDuringExecution []*GetEcciInstancePodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
  // {"en":"The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.", "zh_CN":"调度器更倾向于将 Pod 调度到满足该字段指定的反亲和性表达式的节点， 但它可能会选择违反一个或多个表达式的节点。 最优选的节点是权重总和最大的节点，即对于满足所有调度要求（资源请求、requiredDuringScheduling 反亲和性表达式等）的每个节点，通过遍历元素来计算总和如果节点具有与相应 podAffinityTerm 匹配的 Pod，则此字段并在总和中添加\"权重\"；具有最高加和的节点是最优选的。"}
  PreferredDuringSchedulingIgnoredDuringExecution []*GetEcciInstanceWeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
}

func (s GetEcciInstancePodAntiAffinity) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePodAntiAffinity) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePodAntiAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []*GetEcciInstancePodAffinityTerm) *GetEcciInstancePodAntiAffinity {
  s.RequiredDuringSchedulingIgnoredDuringExecution = v
  return s
}

func (s *GetEcciInstancePodAntiAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []*GetEcciInstanceWeightedPodAffinityTerm) *GetEcciInstancePodAntiAffinity {
  s.PreferredDuringSchedulingIgnoredDuringExecution = v
  return s
}

type GetEcciInstancePodAffinity struct {
  // {"en":"If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.", "zh_CN":"如果在调度时不满足该字段指定的亲和性要求，则该 Pod 不会被调度到该节点上。 如果在 Pod 执行期间的某个时间点不再满足此字段指定的亲和性要求（例如：由于 Pod 标签更新）， 系统可能会也可能不会尝试最终将 Pod 从其节点中逐出。 当此列表中有多个元素时，每个 podAffinityTerm 对应的节点列表是取其交集的，即必须满足所有条件。"}
  RequiredDuringSchedulingIgnoredDuringExecution []*GetEcciInstancePodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
  // {"en":"The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.", "zh_CN":"调度器会更倾向于将 Pod 调度到满足该字段指定的亲和性表达式的节点， 但它可能会选择违反一个或多个表达式的节点。最优选择是权重总和最大的节点， 即对于满足所有调度要求（资源请求、requiredDuringScheduling 亲和表达式等）的每个节点， 通过迭代该字段的元素来计算总和，如果节点具有与相应 podAffinityTerm 匹配的 Pod，则将“权重”添加到总和中； 具有最高总和的节点是最优选的。"}
  PreferredDuringSchedulingIgnoredDuringExecution []*GetEcciInstanceWeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
}

func (s GetEcciInstancePodAffinity) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePodAffinity) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePodAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []*GetEcciInstancePodAffinityTerm) *GetEcciInstancePodAffinity {
  s.RequiredDuringSchedulingIgnoredDuringExecution = v
  return s
}

func (s *GetEcciInstancePodAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []*GetEcciInstanceWeightedPodAffinityTerm) *GetEcciInstancePodAffinity {
  s.PreferredDuringSchedulingIgnoredDuringExecution = v
  return s
}

type GetEcciInstancePodAffinityTerm struct {
  // {"en":"A label query over a set of resources, in this case pods.", "zh_CN":"对一组资源的标签查询，在这里资源为 Pod。"}
  GetEcciInstanceLabelSelector *GetEcciInstanceLabelSelector `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" require:"true"`
  // {"en":"pod's namespace", "zh_CN":" Pod 的名字空间"}
  Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" require:"true" type:"Repeated"`
  // {"en":"This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.", "zh_CN":"此 Pod 应与指定名字空间中与标签选择算符匹配的 Pod 集合位于同一位置（亲和性） 或位于不同位置（反亲和性），这里的“在同一位置”意味着运行在一个节点上，其键名为 topologyKey 的标签值与运行所选 Pod 集合中的某 Pod 的任何节点上的标签值匹配。 不允许使用空的 topologyKey。"}
  TopologyKey *string `json:"topologyKey,omitempty" xml:"topologyKey,omitempty" require:"true"`
}

func (s GetEcciInstancePodAffinityTerm) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstancePodAffinityTerm) GoString() string {
  return s.String()
}

func (s *GetEcciInstancePodAffinityTerm) SetLabelSelector(v *GetEcciInstanceLabelSelector) *GetEcciInstancePodAffinityTerm {
  s.GetEcciInstanceLabelSelector = v
  return s
}

func (s *GetEcciInstancePodAffinityTerm) SetNamespaces(v []*string) *GetEcciInstancePodAffinityTerm {
  s.Namespaces = v
  return s
}

func (s *GetEcciInstancePodAffinityTerm) SetTopologyKey(v string) *GetEcciInstancePodAffinityTerm {
  s.TopologyKey = &v
  return s
}

type GetEcciInstanceLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty" require:"true"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*GetEcciInstanceLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" require:"true" type:"Repeated"`
}

func (s GetEcciInstanceLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceLabelSelector) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceLabelSelector) SetMatchLabels(v map[string]*string) *GetEcciInstanceLabelSelector {
  s.MatchLabels = v
  return s
}

func (s *GetEcciInstanceLabelSelector) SetMatchExpressions(v []*GetEcciInstanceLabelSelectorRequirement) *GetEcciInstanceLabelSelector {
  s.MatchExpressions = v
  return s
}

type GetEcciInstanceLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" require:"true" type:"Repeated"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
}

func (s GetEcciInstanceLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceLabelSelectorRequirement) SetKey(v string) *GetEcciInstanceLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *GetEcciInstanceLabelSelectorRequirement) SetValues(v []*string) *GetEcciInstanceLabelSelectorRequirement {
  s.Values = v
  return s
}

func (s *GetEcciInstanceLabelSelectorRequirement) SetOperator(v string) *GetEcciInstanceLabelSelectorRequirement {
  s.Operator = &v
  return s
}

type GetEcciInstanceWeightedPodAffinityTerm struct {
  // {"en":"associated with matching the corresponding podAffinityTerm, in the range 1-100.", "zh_CN":"匹配相应 podAffinityTerm 条件的权重，范围为 1-100。"}
  Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"Required. A pod affinity term, associated with the corresponding weight.", "zh_CN":"必需的字段。一个 Pod 亲和性条件，对应一个与相应的权重值。"}
  GetEcciInstancePodAffinityTerm *GetEcciInstancePodAffinityTerm `json:"podAffinityTerm,omitempty" xml:"podAffinityTerm,omitempty" require:"true"`
}

func (s GetEcciInstanceWeightedPodAffinityTerm) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceWeightedPodAffinityTerm) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceWeightedPodAffinityTerm) SetWeight(v int32) *GetEcciInstanceWeightedPodAffinityTerm {
  s.Weight = &v
  return s
}

func (s *GetEcciInstanceWeightedPodAffinityTerm) SetPodAffinityTerm(v *GetEcciInstancePodAffinityTerm) *GetEcciInstanceWeightedPodAffinityTerm {
  s.GetEcciInstancePodAffinityTerm = v
  return s
}




type DeleteEcciInstanceRequest struct {
}

func (s DeleteEcciInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteEcciInstanceRequest) GoString() string {
  return s.String()
}

type DeleteEcciInstanceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s DeleteEcciInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteEcciInstanceResponse) GoString() string {
  return s.String()
}

func (s *DeleteEcciInstanceResponse) SetCode(v int64) *DeleteEcciInstanceResponse {
  s.Code = &v
  return s
}

func (s *DeleteEcciInstanceResponse) SetMsg(v string) *DeleteEcciInstanceResponse {
  s.Msg = &v
  return s
}

func (s *DeleteEcciInstanceResponse) SetRequestId(v string) *DeleteEcciInstanceResponse {
  s.RequestId = &v
  return s
}

type DeleteEcciInstancePaths struct {
  // {"en":"instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteEcciInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteEcciInstancePaths) GoString() string {
  return s.String()
}

func (s *DeleteEcciInstancePaths) SetName(v string) *DeleteEcciInstancePaths {
  s.Name = &v
  return s
}

type DeleteEcciInstanceParameters struct {
}

func (s DeleteEcciInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteEcciInstanceParameters) GoString() string {
  return s.String()
}

type DeleteEcciInstanceRequestHeader struct {
}

func (s DeleteEcciInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEcciInstanceRequestHeader) GoString() string {
  return s.String()
}

type DeleteEcciInstanceResponseHeader struct {
}

func (s DeleteEcciInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEcciInstanceResponseHeader) GoString() string {
  return s.String()
}




type InstanceRebuildRequest struct {
  // {"en":"vm id","zh_CN":"云主机ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Image ID","zh_CN":"镜像ID（指定了镜像ID则使用指定的镜像重装，否则使用原镜像重装）"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
  // {"en":"password","zh_CN":"密码（使用公共镜像重装必须指定密码，使用自定义镜像可不指定）"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"Retain Data Disk","zh_CN":"是否保留数据盘（1：是；-1：否）"}
  RetainDataDisk *int `json:"retainDataDisk,omitempty" xml:"retainDataDisk,omitempty"`
}

func (s InstanceRebuildRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildRequest) GoString() string {
  return s.String()
}

func (s *InstanceRebuildRequest) SetId(v string) *InstanceRebuildRequest {
  s.Id = &v
  return s
}

func (s *InstanceRebuildRequest) SetImageId(v string) *InstanceRebuildRequest {
  s.ImageId = &v
  return s
}

func (s *InstanceRebuildRequest) SetPassword(v string) *InstanceRebuildRequest {
  s.Password = &v
  return s
}

func (s *InstanceRebuildRequest) SetRetainDataDisk(v int) *InstanceRebuildRequest {
  s.RetainDataDisk = &v
  return s
}

type InstanceRebuildRequestHeader struct {
}

func (s InstanceRebuildRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildRequestHeader) GoString() string {
  return s.String()
}

type InstanceRebuildPaths struct {
}

func (s InstanceRebuildPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildPaths) GoString() string {
  return s.String()
}

type InstanceRebuildParameters struct {
}

func (s InstanceRebuildParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildParameters) GoString() string {
  return s.String()
}

type InstanceRebuildResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s InstanceRebuildResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildResponse) GoString() string {
  return s.String()
}

func (s *InstanceRebuildResponse) SetCode(v string) *InstanceRebuildResponse {
  s.Code = &v
  return s
}

func (s *InstanceRebuildResponse) SetMessage(v string) *InstanceRebuildResponse {
  s.Message = &v
  return s
}

type InstanceRebuildResponseHeader struct {
}

func (s InstanceRebuildResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceOperationRequest struct {
  // {"en":"Instance action.\nValues:\nSTART: START the\nSHUTDOWN: SHUTDOWN\nREBOOT: Force a REBOOT","zh_CN":"实例操作动作。\n取值：\nSTART：启动\nSHUTDOWN：正常关机\nREBOOT：强制重启"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty" require:"true"`
}

func (s LECHInstanceOperationRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceOperationRequest) GoString() string {
  return s.String()
}

func (s *LECHInstanceOperationRequest) SetOperation(v string) *LECHInstanceOperationRequest {
  s.Operation = &v
  return s
}

type LECHInstanceOperationRequestHeader struct {
}

func (s LECHInstanceOperationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceOperationRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceOperationPaths struct {
  // {"en":"Unique identity of virtual machine","zh_CN":"实例唯一标识"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s LECHInstanceOperationPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceOperationPaths) GoString() string {
  return s.String()
}

func (s *LECHInstanceOperationPaths) SetServerId(v string) *LECHInstanceOperationPaths {
  s.ServerId = &v
  return s
}

type LECHInstanceOperationParameters struct {
}

func (s LECHInstanceOperationParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceOperationParameters) GoString() string {
  return s.String()
}

type LECHInstanceOperationResponse struct {
}

func (s LECHInstanceOperationResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceOperationResponse) GoString() string {
  return s.String()
}

type LECHInstanceOperationResponseHeader struct {
}

func (s LECHInstanceOperationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceOperationResponseHeader) GoString() string {
  return s.String()
}




type GetTokenRequest struct {
}

func (s GetTokenRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
  return s.String()
}

type GetTokenResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"token", "zh_CN":"token"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
  return s.String()
}

func (s *GetTokenResponse) SetCode(v int64) *GetTokenResponse {
  s.Code = &v
  return s
}

func (s *GetTokenResponse) SetMsg(v string) *GetTokenResponse {
  s.Msg = &v
  return s
}

func (s *GetTokenResponse) SetRequestId(v string) *GetTokenResponse {
  s.RequestId = &v
  return s
}

func (s *GetTokenResponse) SetData(v string) *GetTokenResponse {
  s.Data = &v
  return s
}

type GetTokenPaths struct {
}

func (s GetTokenPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTokenPaths) GoString() string {
  return s.String()
}

type GetTokenParameters struct {
}

func (s GetTokenParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTokenParameters) GoString() string {
  return s.String()
}

type GetTokenRequestHeader struct {
}

func (s GetTokenRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTokenRequestHeader) GoString() string {
  return s.String()
}

type GetTokenResponseHeader struct {
}

func (s GetTokenResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTokenResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceBandwidthAggregationQueryRequest struct {
}

func (s LECHInstanceBandwidthAggregationQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryRequest) GoString() string {
  return s.String()
}

type LECHInstanceBandwidthAggregationQueryRequestHeader struct {
}

func (s LECHInstanceBandwidthAggregationQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceBandwidthAggregationQueryPaths struct {
}

func (s LECHInstanceBandwidthAggregationQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryPaths) GoString() string {
  return s.String()
}

type LECHInstanceBandwidthAggregationQueryParameters struct {
  // {"en":"Instance ID: A maximum of 20 IDs can be sent at a time, and IDs are separated by a comma character ','.","zh_CN":"实例ID，单次最多可发送20条ID，ID之间用半角逗号字符“,”隔开。已销毁的实例不支持查询。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Start time,format: YYYY-MM-DD. You can only query the flow data within the last 90 days, and the query range cannot exceed 31 days at a time.","zh_CN":"查询开始时间，格式yyyy-MM-dd。最多只能查询90天内的流量数据，且单次查询时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time: YYYY-MM-DD,and the query range cannot exceed 31 days at a time.","zh_CN":"查询结束时间，格式yyyy-MM-dd，单次查询时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s LECHInstanceBandwidthAggregationQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryParameters) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidthAggregationQueryParameters) SetIds(v string) *LECHInstanceBandwidthAggregationQueryParameters {
  s.Ids = &v
  return s
}

func (s *LECHInstanceBandwidthAggregationQueryParameters) SetStartTime(v string) *LECHInstanceBandwidthAggregationQueryParameters {
  s.StartTime = &v
  return s
}

func (s *LECHInstanceBandwidthAggregationQueryParameters) SetEndTime(v string) *LECHInstanceBandwidthAggregationQueryParameters {
  s.EndTime = &v
  return s
}

type LECHInstanceBandwidthAggregationQueryResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *LECHInstanceBandwidthAggregationQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHInstanceBandwidthAggregationQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryResponse) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidthAggregationQueryResponse) SetCode(v string) *LECHInstanceBandwidthAggregationQueryResponse {
  s.Code = &v
  return s
}

func (s *LECHInstanceBandwidthAggregationQueryResponse) SetData(v *LECHInstanceBandwidthAggregationQueryResponseData) *LECHInstanceBandwidthAggregationQueryResponse {
  s.Data = v
  return s
}

func (s *LECHInstanceBandwidthAggregationQueryResponse) SetMessage(v string) *LECHInstanceBandwidthAggregationQueryResponse {
  s.Message = &v
  return s
}

type LECHInstanceBandwidthAggregationQueryResponseData struct {
  // {"en":"instances detail info","zh_CN":"虚拟机详细信息"}
  Servers []*LECHInstanceBandwidthAggregationQueryResponseDataServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s LECHInstanceBandwidthAggregationQueryResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryResponseData) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidthAggregationQueryResponseData) SetServers(v []*LECHInstanceBandwidthAggregationQueryResponseDataServers) *LECHInstanceBandwidthAggregationQueryResponseData {
  s.Servers = v
  return s
}

type LECHInstanceBandwidthAggregationQueryResponseDataServers struct     {
  // {"en":"External Inbound Traffic Summary (MB)","zh_CN":"外网流入流量汇总值（MB）"}
  ExtTrafficIn *int64 `json:"extTrafficIn,omitempty" xml:"extTrafficIn,omitempty" require:"true"`
  // {"en":"External outbound Traffic Summary (MB)","zh_CN":"外网流出流量汇总值（MB）"}
  ExtTrafficOut *int64 `json:"extTrafficOut,omitempty" xml:"extTrafficOut,omitempty" require:"true"`
  // {"en":"instance Id","zh_CN":"实例Id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHInstanceBandwidthAggregationQueryResponseDataServers) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryResponseDataServers) GoString() string {
  return s.String()
}

func (s *LECHInstanceBandwidthAggregationQueryResponseDataServers) SetExtTrafficIn(v int64) *LECHInstanceBandwidthAggregationQueryResponseDataServers {
  s.ExtTrafficIn = &v
  return s
}

func (s *LECHInstanceBandwidthAggregationQueryResponseDataServers) SetExtTrafficOut(v int64) *LECHInstanceBandwidthAggregationQueryResponseDataServers {
  s.ExtTrafficOut = &v
  return s
}

func (s *LECHInstanceBandwidthAggregationQueryResponseDataServers) SetId(v string) *LECHInstanceBandwidthAggregationQueryResponseDataServers {
  s.Id = &v
  return s
}

type LECHInstanceBandwidthAggregationQueryResponseHeader struct {
}

func (s LECHInstanceBandwidthAggregationQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceBandwidthAggregationQueryResponseHeader) GoString() string {
  return s.String()
}




type VMPRemoveInstanceRequest struct {
  // {"en":"Instance ID", "zh_CN":"实例唯一标识。单次最多可发送100 条ID，ID 之间用半角逗号字符“,”隔开。"}
  Servers *string `json:"servers,omitempty" xml:"servers,omitempty" require:"true"`
}

func (s VMPRemoveInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceRequest) GoString() string {
  return s.String()
}

func (s *VMPRemoveInstanceRequest) SetServers(v string) *VMPRemoveInstanceRequest {
  s.Servers = &v
  return s
}

type VMPRemoveInstanceResponse struct {
}

func (s VMPRemoveInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceResponse) GoString() string {
  return s.String()
}

type VMPRemoveInstancePaths struct {
}

func (s VMPRemoveInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstancePaths) GoString() string {
  return s.String()
}

type VMPRemoveInstanceParameters struct {
}

func (s VMPRemoveInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceParameters) GoString() string {
  return s.String()
}

type VMPRemoveInstanceRequestHeader struct {
}

func (s VMPRemoveInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceRequestHeader) GoString() string {
  return s.String()
}

type VMPRemoveInstanceResponseHeader struct {
}

func (s VMPRemoveInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceResponseHeader) GoString() string {
  return s.String()
}




type QueryGpsRequest struct {
}

func (s QueryGpsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsRequest) GoString() string {
  return s.String()
}

type QueryGpsResponse struct {
  // {"en":"list of ephone instances", "zh_CN":"云手机实例列表"}
  Data []*QueryGpsGpsInfo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s QueryGpsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsResponse) GoString() string {
  return s.String()
}

func (s *QueryGpsResponse) SetData(v []*QueryGpsGpsInfo) *QueryGpsResponse {
  s.Data = v
  return s
}

func (s *QueryGpsResponse) SetStatus(v int) *QueryGpsResponse {
  s.Status = &v
  return s
}

func (s *QueryGpsResponse) SetResult(v string) *QueryGpsResponse {
  s.Result = &v
  return s
}

type QueryGpsGpsInfo struct {
  // {"en":"longitude", "zh_CN":"经度"}
  Longitude *float32 `json:"longitude,omitempty" xml:"longitude,omitempty" require:"true"`
  // {"en":"latitude", "zh_CN":"纬度"}
  Latitude *float32 `json:"latitude,omitempty" xml:"latitude,omitempty" require:"true"`
  // {"en":"instance id", "zh_CN":"实例id"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"query status", "zh_CN":"查询状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QueryGpsGpsInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsGpsInfo) GoString() string {
  return s.String()
}

func (s *QueryGpsGpsInfo) SetLongitude(v float32) *QueryGpsGpsInfo {
  s.Longitude = &v
  return s
}

func (s *QueryGpsGpsInfo) SetLatitude(v float32) *QueryGpsGpsInfo {
  s.Latitude = &v
  return s
}

func (s *QueryGpsGpsInfo) SetInstanceId(v string) *QueryGpsGpsInfo {
  s.InstanceId = &v
  return s
}

func (s *QueryGpsGpsInfo) SetStatus(v string) *QueryGpsGpsInfo {
  s.Status = &v
  return s
}

func (s *QueryGpsGpsInfo) SetMessage(v string) *QueryGpsGpsInfo {
  s.Message = &v
  return s
}

type QueryGpsPaths struct {
}

func (s QueryGpsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsPaths) GoString() string {
  return s.String()
}

type QueryGpsParameters struct {
  // {"en":"instance ID to be queried", "zh_CN":"要查询的实例id"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"instance name bo be queried", "zh_CN":"要查询的实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryGpsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsParameters) GoString() string {
  return s.String()
}

func (s *QueryGpsParameters) SetIds(v string) *QueryGpsParameters {
  s.Ids = &v
  return s
}

func (s *QueryGpsParameters) SetName(v string) *QueryGpsParameters {
  s.Name = &v
  return s
}

type QueryGpsRequestHeader struct {
}

func (s QueryGpsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsRequestHeader) GoString() string {
  return s.String()
}

type QueryGpsResponseHeader struct {
}

func (s QueryGpsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryGpsResponseHeader) GoString() string {
  return s.String()
}




type ScreenshotEphoneInstanceRequest struct {
  // {"en":"list of instances to operate on", "zh_CN":"操作实例的数组对象"}
  Ephones []*ScreenshotEphoneInstanceScreenshotObject `json:"ephones,omitempty" xml:"ephones,omitempty" require:"true" type:"Repeated"`
}

func (s ScreenshotEphoneInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceRequest) GoString() string {
  return s.String()
}

func (s *ScreenshotEphoneInstanceRequest) SetEphones(v []*ScreenshotEphoneInstanceScreenshotObject) *ScreenshotEphoneInstanceRequest {
  s.Ephones = v
  return s
}

type ScreenshotEphoneInstanceScreenshotObject struct {
  // {"en":"ephone instance id", "zh_CN":"云手机实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"screenshot image size, optional value:  tiny, small, medium, large", "zh_CN":"截图大小，可选值有: tiny, small, medium, large"}
  Size *string `json:"size,omitempty" xml:"size,omitempty" require:"true"`
}

func (s ScreenshotEphoneInstanceScreenshotObject) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceScreenshotObject) GoString() string {
  return s.String()
}

func (s *ScreenshotEphoneInstanceScreenshotObject) SetId(v string) *ScreenshotEphoneInstanceScreenshotObject {
  s.Id = &v
  return s
}

func (s *ScreenshotEphoneInstanceScreenshotObject) SetSize(v string) *ScreenshotEphoneInstanceScreenshotObject {
  s.Size = &v
  return s
}

type ScreenshotEphoneInstanceResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Ephones []*ScreenshotEphoneInstanceScreenshotData `json:"ephones,omitempty" xml:"ephones,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s ScreenshotEphoneInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceResponse) GoString() string {
  return s.String()
}

func (s *ScreenshotEphoneInstanceResponse) SetEphones(v []*ScreenshotEphoneInstanceScreenshotData) *ScreenshotEphoneInstanceResponse {
  s.Ephones = v
  return s
}

func (s *ScreenshotEphoneInstanceResponse) SetStatus(v int) *ScreenshotEphoneInstanceResponse {
  s.Status = &v
  return s
}

func (s *ScreenshotEphoneInstanceResponse) SetResult(v string) *ScreenshotEphoneInstanceResponse {
  s.Result = &v
  return s
}

type ScreenshotEphoneInstanceScreenshotData struct {
  // {"en":"ephone instance id", "zh_CN":"云手机实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"data encoded via base64 for screenshot image", "zh_CN":"base64编码过的截图数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ScreenshotEphoneInstanceScreenshotData) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceScreenshotData) GoString() string {
  return s.String()
}

func (s *ScreenshotEphoneInstanceScreenshotData) SetId(v string) *ScreenshotEphoneInstanceScreenshotData {
  s.Id = &v
  return s
}

func (s *ScreenshotEphoneInstanceScreenshotData) SetData(v string) *ScreenshotEphoneInstanceScreenshotData {
  s.Data = &v
  return s
}

func (s *ScreenshotEphoneInstanceScreenshotData) SetMessage(v string) *ScreenshotEphoneInstanceScreenshotData {
  s.Message = &v
  return s
}

type ScreenshotEphoneInstancePaths struct {
}

func (s ScreenshotEphoneInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstancePaths) GoString() string {
  return s.String()
}

type ScreenshotEphoneInstanceParameters struct {
}

func (s ScreenshotEphoneInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceParameters) GoString() string {
  return s.String()
}

type ScreenshotEphoneInstanceRequestHeader struct {
}

func (s ScreenshotEphoneInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceRequestHeader) GoString() string {
  return s.String()
}

type ScreenshotEphoneInstanceResponseHeader struct {
}

func (s ScreenshotEphoneInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ScreenshotEphoneInstanceResponseHeader) GoString() string {
  return s.String()
}




type VMPInstanceOperationRequest struct {
  // {"en":"Instance action.
  // Values:
  // START: START the
  // SHUTDOWN: SHUTDOWN
  // REBOOT: Force a REBOOT", "zh_CN":"实例操作动作。
  // 取值：
  // START：启动
  // SHUTDOWN：正常关机
  // REBOOT：强制重启"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty" require:"true"`
}

func (s VMPInstanceOperationRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationRequest) GoString() string {
  return s.String()
}

func (s *VMPInstanceOperationRequest) SetOperation(v string) *VMPInstanceOperationRequest {
  s.Operation = &v
  return s
}

type VMPInstanceOperationResponse struct {
}

func (s VMPInstanceOperationResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationResponse) GoString() string {
  return s.String()
}

type VMPInstanceOperationPaths struct {
  // {"en":"Unique identity of virtual machine", "zh_CN":"实例唯一标识"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s VMPInstanceOperationPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationPaths) GoString() string {
  return s.String()
}

func (s *VMPInstanceOperationPaths) SetServerId(v string) *VMPInstanceOperationPaths {
  s.ServerId = &v
  return s
}

type VMPInstanceOperationParameters struct {
}

func (s VMPInstanceOperationParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationParameters) GoString() string {
  return s.String()
}

type VMPInstanceOperationRequestHeader struct {
}

func (s VMPInstanceOperationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationRequestHeader) GoString() string {
  return s.String()
}

type VMPInstanceOperationResponseHeader struct {
}

func (s VMPInstanceOperationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationResponseHeader) GoString() string {
  return s.String()
}




type LECHSetTrafficLimitRequest struct {
  // {"en":"instance id list","zh_CN":"实例id列表，单次最多填写50个"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" require:"true" type:"Repeated"`
  // {"en":"Natural Month Traffic limit(G), range 1-65535. Traffic limit cannot exceed order constraints.. Not filling in means no limitation.","zh_CN":"自然月流量上限值，范围1-65535。流量上限值不可超过订单约束的上限值。不填意为不限制。"}
  TrafficLimit *int `json:"trafficLimit,omitempty" xml:"trafficLimit,omitempty"`
}

func (s LECHSetTrafficLimitRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitRequest) GoString() string {
  return s.String()
}

func (s *LECHSetTrafficLimitRequest) SetIds(v []*string) *LECHSetTrafficLimitRequest {
  s.Ids = v
  return s
}

func (s *LECHSetTrafficLimitRequest) SetTrafficLimit(v int) *LECHSetTrafficLimitRequest {
  s.TrafficLimit = &v
  return s
}

type LECHSetTrafficLimitRequestHeader struct {
}

func (s LECHSetTrafficLimitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitRequestHeader) GoString() string {
  return s.String()
}

type LECHSetTrafficLimitPaths struct {
}

func (s LECHSetTrafficLimitPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitPaths) GoString() string {
  return s.String()
}

type LECHSetTrafficLimitParameters struct {
}

func (s LECHSetTrafficLimitParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitParameters) GoString() string {
  return s.String()
}

type LECHSetTrafficLimitResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data []*LECHSetTrafficLimitResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHSetTrafficLimitResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitResponse) GoString() string {
  return s.String()
}

func (s *LECHSetTrafficLimitResponse) SetCode(v string) *LECHSetTrafficLimitResponse {
  s.Code = &v
  return s
}

func (s *LECHSetTrafficLimitResponse) SetData(v []*LECHSetTrafficLimitResponseData) *LECHSetTrafficLimitResponse {
  s.Data = v
  return s
}

func (s *LECHSetTrafficLimitResponse) SetMessage(v string) *LECHSetTrafficLimitResponse {
  s.Message = &v
  return s
}

type LECHSetTrafficLimitResponseData struct     {
  // {"en":"result code","zh_CN":"结果状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"result message","zh_CN":"结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"instance name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s LECHSetTrafficLimitResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitResponseData) GoString() string {
  return s.String()
}

func (s *LECHSetTrafficLimitResponseData) SetCode(v int) *LECHSetTrafficLimitResponseData {
  s.Code = &v
  return s
}

func (s *LECHSetTrafficLimitResponseData) SetId(v string) *LECHSetTrafficLimitResponseData {
  s.Id = &v
  return s
}

func (s *LECHSetTrafficLimitResponseData) SetMessage(v string) *LECHSetTrafficLimitResponseData {
  s.Message = &v
  return s
}

func (s *LECHSetTrafficLimitResponseData) SetName(v string) *LECHSetTrafficLimitResponseData {
  s.Name = &v
  return s
}

type LECHSetTrafficLimitResponseHeader struct {
}

func (s LECHSetTrafficLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHSetTrafficLimitResponseHeader) GoString() string {
  return s.String()
}




type ManageEphoneInstanceRequest struct {
  // {"en":"list of instances to operate on", "zh_CN":"操作实例的数组对象"}
  Ephones []*ManageEphoneInstanceOperateObject `json:"ephones,omitempty" xml:"ephones,omitempty" require:"true" type:"Repeated"`
}

func (s ManageEphoneInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceRequest) GoString() string {
  return s.String()
}

func (s *ManageEphoneInstanceRequest) SetEphones(v []*ManageEphoneInstanceOperateObject) *ManageEphoneInstanceRequest {
  s.Ephones = v
  return s
}

type ManageEphoneInstanceOperateObject struct {
  // {"en":"ephone instance id", "zh_CN":"云手机实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"operate type", "zh_CN":"操作类型，可选值：delete, reboot, renew, clear, setGps, createH5"}
  Op *string `json:"op,omitempty" xml:"op,omitempty" require:"true"`
  // {"en":"app operaate params", "zh_CN":"app 操作的参数"}
  Params *ManageEphoneInstanceAppParamsObject `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ManageEphoneInstanceOperateObject) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceOperateObject) GoString() string {
  return s.String()
}

func (s *ManageEphoneInstanceOperateObject) SetId(v string) *ManageEphoneInstanceOperateObject {
  s.Id = &v
  return s
}

func (s *ManageEphoneInstanceOperateObject) SetOp(v string) *ManageEphoneInstanceOperateObject {
  s.Op = &v
  return s
}

func (s *ManageEphoneInstanceOperateObject) SetParams(v *ManageEphoneInstanceAppParamsObject) *ManageEphoneInstanceOperateObject {
  s.Params = v
  return s
}

type ManageEphoneInstanceAppParamsObject struct {
  // {"en":"oem image id", "zh_CN":"oem镜像id"}
  OemImage *string `json:"oemImage,omitempty" xml:"oemImage,omitempty"`
  // {"en":"longitude", "zh_CN":"经度"}
  Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
  // {"en":"latitude", "zh_CN":"纬度"}
  Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
  // {"en":"hardware info in json format", "zh_CN":"json 格式的硬件信息"}
  Hardware *string `json:"hardware,omitempty" xml:"hardware,omitempty"`
  // {"en":"retain user data", "zh_CN":"renew重装系统时生效，指定true可保留用户数据，不指定或者指定false会清空用户数据"}
  RetainData *string `json:"retainData,omitempty" xml:"retainData,omitempty"`
}

func (s ManageEphoneInstanceAppParamsObject) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceAppParamsObject) GoString() string {
  return s.String()
}

func (s *ManageEphoneInstanceAppParamsObject) SetOemImage(v string) *ManageEphoneInstanceAppParamsObject {
  s.OemImage = &v
  return s
}

func (s *ManageEphoneInstanceAppParamsObject) SetLongitude(v string) *ManageEphoneInstanceAppParamsObject {
  s.Longitude = &v
  return s
}

func (s *ManageEphoneInstanceAppParamsObject) SetLatitude(v string) *ManageEphoneInstanceAppParamsObject {
  s.Latitude = &v
  return s
}

func (s *ManageEphoneInstanceAppParamsObject) SetHardware(v string) *ManageEphoneInstanceAppParamsObject {
  s.Hardware = &v
  return s
}

func (s *ManageEphoneInstanceAppParamsObject) SetRetainData(v string) *ManageEphoneInstanceAppParamsObject {
  s.RetainData = &v
  return s
}

type ManageEphoneInstanceResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Tasks []*ManageEphoneInstanceTask `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s ManageEphoneInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceResponse) GoString() string {
  return s.String()
}

func (s *ManageEphoneInstanceResponse) SetTasks(v []*ManageEphoneInstanceTask) *ManageEphoneInstanceResponse {
  s.Tasks = v
  return s
}

func (s *ManageEphoneInstanceResponse) SetStatus(v int) *ManageEphoneInstanceResponse {
  s.Status = &v
  return s
}

func (s *ManageEphoneInstanceResponse) SetResult(v string) *ManageEphoneInstanceResponse {
  s.Result = &v
  return s
}

type ManageEphoneInstanceTask struct {
  // {"en":"task id", "zh_CN":"任务单id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"任务单消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ManageEphoneInstanceTask) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceTask) GoString() string {
  return s.String()
}

func (s *ManageEphoneInstanceTask) SetId(v string) *ManageEphoneInstanceTask {
  s.Id = &v
  return s
}

func (s *ManageEphoneInstanceTask) SetMessage(v string) *ManageEphoneInstanceTask {
  s.Message = &v
  return s
}

type ManageEphoneInstancePaths struct {
}

func (s ManageEphoneInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstancePaths) GoString() string {
  return s.String()
}

type ManageEphoneInstanceParameters struct {
}

func (s ManageEphoneInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceParameters) GoString() string {
  return s.String()
}

type ManageEphoneInstanceRequestHeader struct {
}

func (s ManageEphoneInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceRequestHeader) GoString() string {
  return s.String()
}

type ManageEphoneInstanceResponseHeader struct {
}

func (s ManageEphoneInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneInstanceResponseHeader) GoString() string {
  return s.String()
}




type RunEphoneAdbShellRequest struct {
  // {"en":"list of instances to run adb shell on", "zh_CN":"要执行adb shell的实例数组对象"}
  Ephones []*RunEphoneAdbShellAdbObject `json:"ephones,omitempty" xml:"ephones,omitempty" require:"true" type:"Repeated"`
}

func (s RunEphoneAdbShellRequest) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellRequest) GoString() string {
  return s.String()
}

func (s *RunEphoneAdbShellRequest) SetEphones(v []*RunEphoneAdbShellAdbObject) *RunEphoneAdbShellRequest {
  s.Ephones = v
  return s
}

type RunEphoneAdbShellAdbObject struct {
  // {"en":"ephone instance id", "zh_CN":"云手机实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"adb shell command content", "zh_CN":"adb shell 命令的内容"}
  Command *string `json:"command,omitempty" xml:"command,omitempty" require:"true"`
}

func (s RunEphoneAdbShellAdbObject) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellAdbObject) GoString() string {
  return s.String()
}

func (s *RunEphoneAdbShellAdbObject) SetId(v string) *RunEphoneAdbShellAdbObject {
  s.Id = &v
  return s
}

func (s *RunEphoneAdbShellAdbObject) SetCommand(v string) *RunEphoneAdbShellAdbObject {
  s.Command = &v
  return s
}

type RunEphoneAdbShellResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Tasks []*RunEphoneAdbShellTask `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s RunEphoneAdbShellResponse) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellResponse) GoString() string {
  return s.String()
}

func (s *RunEphoneAdbShellResponse) SetTasks(v []*RunEphoneAdbShellTask) *RunEphoneAdbShellResponse {
  s.Tasks = v
  return s
}

func (s *RunEphoneAdbShellResponse) SetStatus(v int) *RunEphoneAdbShellResponse {
  s.Status = &v
  return s
}

func (s *RunEphoneAdbShellResponse) SetResult(v string) *RunEphoneAdbShellResponse {
  s.Result = &v
  return s
}

type RunEphoneAdbShellTask struct {
  // {"en":"task id", "zh_CN":"任务单id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"任务单消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RunEphoneAdbShellTask) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellTask) GoString() string {
  return s.String()
}

func (s *RunEphoneAdbShellTask) SetId(v string) *RunEphoneAdbShellTask {
  s.Id = &v
  return s
}

func (s *RunEphoneAdbShellTask) SetMessage(v string) *RunEphoneAdbShellTask {
  s.Message = &v
  return s
}

type RunEphoneAdbShellPaths struct {
}

func (s RunEphoneAdbShellPaths) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellPaths) GoString() string {
  return s.String()
}

type RunEphoneAdbShellParameters struct {
}

func (s RunEphoneAdbShellParameters) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellParameters) GoString() string {
  return s.String()
}

type RunEphoneAdbShellRequestHeader struct {
}

func (s RunEphoneAdbShellRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellRequestHeader) GoString() string {
  return s.String()
}

type RunEphoneAdbShellResponseHeader struct {
}

func (s RunEphoneAdbShellResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RunEphoneAdbShellResponseHeader) GoString() string {
  return s.String()
}




type LECHEditInstanceRequest struct {
  // {"en":"server","zh_CN":"实例信息对象"}
  Server []*LECHEditInstanceRequestServer `json:"server,omitempty" xml:"server,omitempty" require:"true" type:"Repeated"`
}

func (s LECHEditInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstanceRequest) GoString() string {
  return s.String()
}

func (s *LECHEditInstanceRequest) SetServer(v []*LECHEditInstanceRequestServer) *LECHEditInstanceRequest {
  s.Server = v
  return s
}

type LECHEditInstanceRequestServer struct     {
  // {"en":"Server id","zh_CN":"要更新的实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"New instance name\nConstraints:\n1. Length of 2-128 characters\n2. Must start with a letter, and can only contain letters, numbers, underlines, lines, and dots","zh_CN":"新的实例名称\n约束：\n1. 长度2-128个字符\n2. 必须以字母开头，且只能包含字母、数字、下划线、横线、点号"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s LECHEditInstanceRequestServer) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstanceRequestServer) GoString() string {
  return s.String()
}

func (s *LECHEditInstanceRequestServer) SetId(v string) *LECHEditInstanceRequestServer {
  s.Id = &v
  return s
}

func (s *LECHEditInstanceRequestServer) SetName(v string) *LECHEditInstanceRequestServer {
  s.Name = &v
  return s
}

type LECHEditInstanceRequestHeader struct {
}

func (s LECHEditInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstanceRequestHeader) GoString() string {
  return s.String()
}

type LECHEditInstancePaths struct {
}

func (s LECHEditInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstancePaths) GoString() string {
  return s.String()
}

type LECHEditInstanceParameters struct {
}

func (s LECHEditInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstanceParameters) GoString() string {
  return s.String()
}

type LECHEditInstanceResponse struct {
}

func (s LECHEditInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstanceResponse) GoString() string {
  return s.String()
}

type LECHEditInstanceResponseHeader struct {
}

func (s LECHEditInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEditInstanceResponseHeader) GoString() string {
  return s.String()
}




type LECHCreateInstanceRequest struct {
  // {"en":"Creating array objects for virtual machines","zh_CN":"创建实例的数组对象"}
  Servers []*LECHCreateInstanceRequestServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s LECHCreateInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequest) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceRequest) SetServers(v []*LECHCreateInstanceRequestServers) *LECHCreateInstanceRequest {
  s.Servers = v
  return s
}

type LECHCreateInstanceRequestServers struct     {
  // {"en":"Virtual machine area (see Appendix for details)","zh_CN":"实例所属区域（节点名称nodeName和区域regionName至少需要上传一个。\n区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Province of virtual machine (see Appendix for details)","zh_CN":"实例所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"If the operator of the virtual machine (see the appendix for details) carries this parameter, please keep it consistent with the carrier returned from the '3.4 node list query' interface.","zh_CN":"实例所属运营商（dx-电信；wt-网通；yd-移动）如果携带了该参数，请与'3.4节点列表查询'接口返回的carrier保持一致"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node name, indicating that the specified node creates a virtual machine (the node name returned by interface 3.4)","zh_CN":"节点名称，表示指定节点创建实例（节点名称可通过资源管理-节点列表查询接口获取）"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"Virtual machine image identity","zh_CN":"实例镜像标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Virtual machine specification ID","zh_CN":"实例规格标识"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Virtual machine name. If the created quantity is greater than 1, the real name is spliced with 3 digits after the parameter. For example, instance 0001, instance 0002","zh_CN":"实例名称，如果创建数量大于1，则真实名称是在该参数后拼接3位数字。如instance_0001，instance_0002"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Inject user data, support to inject text, text file or gzip file. The maximum length of injected content is 32KB. For content injection, Base64 format encoding is required.","zh_CN":"注入用户数据，支持注入文本、文本文件或gzip文件。注入内容最大长度32KB。注入内容，需要进行base64格式编码。"}
  UserData *string `json:"userData,omitempty" xml:"userData,omitempty"`
  // {"en":"Number of virtual machines applied","zh_CN":"申请实例数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty"`
  // {"en":"Virtual machine root login password","zh_CN":"实例root用户登录密码（如果选择的是公共镜像，则密码password必填）。\n密码规则：大写字母 小写字母 数字 特殊字符，四种包括三种，长度8~30"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"The name of the SSH secret key pair for virtual machine login. If this parameter is specified, the password login mode is disabled by default, and the password parameter is invalid at the same time.","zh_CN":"实例登录SSH秘钥对名称，如果指定该参数，默认禁用密码登录方式，password参数同时失效"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
  // {"en":"Whether the virtual machine needs intranet, value:\nYes: intranet required\nNo: no intranet is required, default value'","zh_CN":"实例是否需要内网网络，取值：\nYES：需要内网\nNO：不需要内网，默认值"}
  InnerNet *string `json:"innerNet,omitempty" xml:"innerNet,omitempty"`
  // {"en":"CIDR of virtual machine intranet is meaningful only when innernet = yes","zh_CN":"实例内网的cidr，只有innerNet=YES时才有意义"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"If IP address is specified, it must be within the scope of CIDR, otherwise creation fails.","zh_CN":"实例内网ip地址，如果指定了ip，必须在cidr的范围内，否则创建失败"}
  PrivateIPv4 *string `json:"privateIPv4,omitempty" xml:"privateIPv4,omitempty"`
  // {"en":"Whether the virtual machine needs intranet2, value:\nYes: intranet2 required\nNo: no intranet2 is required, default value'","zh_CN":"实例是否需要内网2网络，取值：\nYES：需要内网2\nNO：不需要内网2，默认值"}
  InnerNet2 *string `json:"innerNet2,omitempty" xml:"innerNet2,omitempty"`
  // {"en":"CIDR of virtual machine intranet2 is meaningful only when innernet = yes","zh_CN":"实例内网2的cidr，只有innerNet2=YES时才有意义"}
  Cidr2 *string `json:"cidr2,omitempty" xml:"cidr2,omitempty"`
  // {"en":"If IP address is specified, it must be within the scope of CIDR2, otherwise creation fails.","zh_CN":"实例内网2ip地址，如果指定了ip，必须在cidr2的范围内，否则创建失败"}
  PrivateIPv42 *string `json:"privateIPv42,omitempty" xml:"privateIPv42,omitempty"`
  // {"en":"Inner ipv6 info","zh_CN":"内网IPv6信息"}
  PrivateIpv6Info []*LECHCreateInstanceRequestServersPrivateIpv6Info `json:"privateIpv6Info,omitempty" xml:"privateIpv6Info,omitempty" type:"Repeated"`
  // {"en":"Whether multiple IP protocol addresses are required\n\n4: only IPv4 address is required, default value\n\n0: both IPv4 and IPv6 need'","zh_CN":"是否需要多ip协议地址\n4：只需要ipv4地址，默认值\n0：ipv4、ipv6都需要"}
  Protocols *int `json:"protocols,omitempty" xml:"protocols,omitempty"`
  // {"en":"IPv4 native attribute, 1: non-native;-1: native;","zh_CN":"IPv4原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  Ipv4NativeAttribute *string `json:"ipv4NativeAttribute,omitempty" xml:"ipv4NativeAttribute,omitempty"`
  // {"en":"IPv6 native attribute, 1: non-native;-1: native;","zh_CN":"IPv6原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  Ipv6NativeAttribute *string `json:"ipv6NativeAttribute,omitempty" xml:"ipv6NativeAttribute,omitempty"`
  // {"en":"Whether the instance is free or not, the default billing instance, and the bare machine instance cannot be free, values are as follows:\nYes: Free instances\nNo: Billing instance\nIf you are using a free instance, you need to configure permissions in advance","zh_CN":"是否免费实例，默认计费实例，裸机实例不能免费，取值：\nYES：免费实例\nNO：计费实例\n如果使用免费实例，需要提前配置权限"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty"`
  // {"en":"Specify a security group ID to create multiple security groups separated by commas, up to 5\nIf you are creating a bare machine, you cannot specify a security group","zh_CN":"指定安全组id进行创建，多个安全组以逗号分隔，最多指定5个\n如果是创建裸机，不能指定安全组"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
  // {"en":"Disk information\nIf this information is carried, the disk definition on the template will be ignored and the instance disk will be created with this information, not for bare-metal instance creation","zh_CN":"磁盘信息\n如果携带该信息，将忽略模板上的磁盘定义，以该信息创建实例磁盘，不适用于裸机实例创建"}
  DiskInfo []*LECHCreateInstanceRequestServersDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" type:"Repeated"`
  // {"en":"Anti-affinity group name\nVirtual machines with the same ServerGroup are created on different hosts","zh_CN":"反亲和性组名称\n拥有相同serverGroup的虚拟机会被创建在不同的宿主机上"}
  ServerGroup *string `json:"serverGroup,omitempty" xml:"serverGroup,omitempty"`
  // {"en":"Instance Tag","zh_CN":"实例标签"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
  // {"en":"Use  unique ip segment","zh_CN":"是否使用唯一网段\n1：是\n-1：否"}
  UseUniqueIpSegment *int `json:"useUniqueIpSegment,omitempty" xml:"useUniqueIpSegment,omitempty"`
  // {"en":"Allocate IP randomly","zh_CN":"是否需要随机分配IPv4\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
  // {"en":"Default Gateway","zh_CN":"默认网关运营商如：dx-电信；yd-移动；wt-网通"}
  DefaultGateway *string `json:"defaultGateway,omitempty" xml:"defaultGateway,omitempty"`
  // {"en":"Policy routing type","zh_CN":"策略路由类型：0-目的地址策略路由（默认）；1-源地址策略路由"}
  PolicyRoutingType *int `json:"policyRoutingType,omitempty" xml:"policyRoutingType,omitempty"`
  // {"en":"Nic allocate type","zh_CN":"实例网卡分配方式：0-多个ip共用一张网卡（默认）；1-每个ip独立一张网卡；2-V4V6混合，同协议IP同网卡，不同线路IP不同网卡"}
  NicAllocateType *int `json:"nicAllocateType,omitempty" xml:"nicAllocateType,omitempty"`
  // {"en":"Ipv4 cidr","zh_CN":"指定外网IPv4网段CIDR(不支持多线)"}
  SinglePublicIpv4Cidr *string `json:"singlePublicIpv4Cidr,omitempty" xml:"singlePublicIpv4Cidr,omitempty"`
  // {"en":"Specify certain public IPv4 ISPs; effective for multi-line nodes; if not specified, an instance with IPs from all carriers is created by default.","zh_CN":"指定部分公网ipv4运营商, 仅多线节点生效。多线节点未指定该参数时，默认创建包含所有运营商IP的实例"}
  PublicIpv4Info []*LECHCreateInstanceRequestServersPublicIpv4Info `json:"publicIpv4Info,omitempty" xml:"publicIpv4Info,omitempty" type:"Repeated"`
  // {"en":"Specify certain public IPv6 ISPs; effective for multi-line nodes; if not specified, an instance with IPs from all carriers is created by default.","zh_CN":"指定部分公网ipv6运营商, 仅多线节点生效。多线节点指定需要ipv6未指定该参数时，默认创建包含所有运营商IP的实例"}
  PublicIpv6Info []*LECHCreateInstanceRequestServersPublicIpv6Info `json:"publicIpv6Info,omitempty" xml:"publicIpv6Info,omitempty" type:"Repeated"`
}

func (s LECHCreateInstanceRequestServers) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequestServers) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceRequestServers) SetRegionName(v string) *LECHCreateInstanceRequestServers {
  s.RegionName = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetProvince(v string) *LECHCreateInstanceRequestServers {
  s.Province = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetCarrier(v string) *LECHCreateInstanceRequestServers {
  s.Carrier = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetNodeName(v string) *LECHCreateInstanceRequestServers {
  s.NodeName = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetImageId(v string) *LECHCreateInstanceRequestServers {
  s.ImageId = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetFlavorId(v string) *LECHCreateInstanceRequestServers {
  s.FlavorId = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetName(v string) *LECHCreateInstanceRequestServers {
  s.Name = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetUserData(v string) *LECHCreateInstanceRequestServers {
  s.UserData = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetCount(v int) *LECHCreateInstanceRequestServers {
  s.Count = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPassword(v string) *LECHCreateInstanceRequestServers {
  s.Password = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetKeyName(v string) *LECHCreateInstanceRequestServers {
  s.KeyName = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetInnerNet(v string) *LECHCreateInstanceRequestServers {
  s.InnerNet = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetCidr(v string) *LECHCreateInstanceRequestServers {
  s.Cidr = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPrivateIPv4(v string) *LECHCreateInstanceRequestServers {
  s.PrivateIPv4 = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetInnerNet2(v string) *LECHCreateInstanceRequestServers {
  s.InnerNet2 = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetCidr2(v string) *LECHCreateInstanceRequestServers {
  s.Cidr2 = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPrivateIPv42(v string) *LECHCreateInstanceRequestServers {
  s.PrivateIPv42 = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPrivateIpv6Info(v []*LECHCreateInstanceRequestServersPrivateIpv6Info) *LECHCreateInstanceRequestServers {
  s.PrivateIpv6Info = v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetProtocols(v int) *LECHCreateInstanceRequestServers {
  s.Protocols = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetIpv4NativeAttribute(v string) *LECHCreateInstanceRequestServers {
  s.Ipv4NativeAttribute = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetIpv6NativeAttribute(v string) *LECHCreateInstanceRequestServers {
  s.Ipv6NativeAttribute = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetIsFree(v string) *LECHCreateInstanceRequestServers {
  s.IsFree = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetSecurityGroupIds(v []*string) *LECHCreateInstanceRequestServers {
  s.SecurityGroupIds = v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetDiskInfo(v []*LECHCreateInstanceRequestServersDiskInfo) *LECHCreateInstanceRequestServers {
  s.DiskInfo = v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetServerGroup(v string) *LECHCreateInstanceRequestServers {
  s.ServerGroup = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetTag(v string) *LECHCreateInstanceRequestServers {
  s.Tag = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetUseUniqueIpSegment(v int) *LECHCreateInstanceRequestServers {
  s.UseUniqueIpSegment = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetRandomAllocateIp(v int) *LECHCreateInstanceRequestServers {
  s.RandomAllocateIp = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetDefaultGateway(v string) *LECHCreateInstanceRequestServers {
  s.DefaultGateway = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPolicyRoutingType(v int) *LECHCreateInstanceRequestServers {
  s.PolicyRoutingType = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetNicAllocateType(v int) *LECHCreateInstanceRequestServers {
  s.NicAllocateType = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetSinglePublicIpv4Cidr(v string) *LECHCreateInstanceRequestServers {
  s.SinglePublicIpv4Cidr = &v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPublicIpv4Info(v []*LECHCreateInstanceRequestServersPublicIpv4Info) *LECHCreateInstanceRequestServers {
  s.PublicIpv4Info = v
  return s
}

func (s *LECHCreateInstanceRequestServers) SetPublicIpv6Info(v []*LECHCreateInstanceRequestServersPublicIpv6Info) *LECHCreateInstanceRequestServers {
  s.PublicIpv6Info = v
  return s
}

type LECHCreateInstanceRequestServersPrivateIpv6Info struct     {
  // {"en":"Inner network number","zh_CN":"内网编号（1-对应v4的内网1；2-对应v4的内网2）"}
  NetNo *int `json:"netNo,omitempty" xml:"netNo,omitempty"`
  // {"en":"Inner network ipv6 cidr","zh_CN":"指定内网IPv6 CIDR"}
  PrivateCidr *string `json:"privateCidr,omitempty" xml:"privateCidr,omitempty"`
  // {"en":"Inner network ipv6 address:","zh_CN":"指定内网IPv6地址"}
  PrivateIp *string `json:"privateIp,omitempty" xml:"privateIp,omitempty"`
}

func (s LECHCreateInstanceRequestServersPrivateIpv6Info) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequestServersPrivateIpv6Info) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceRequestServersPrivateIpv6Info) SetNetNo(v int) *LECHCreateInstanceRequestServersPrivateIpv6Info {
  s.NetNo = &v
  return s
}

func (s *LECHCreateInstanceRequestServersPrivateIpv6Info) SetPrivateCidr(v string) *LECHCreateInstanceRequestServersPrivateIpv6Info {
  s.PrivateCidr = &v
  return s
}

func (s *LECHCreateInstanceRequestServersPrivateIpv6Info) SetPrivateIp(v string) *LECHCreateInstanceRequestServersPrivateIpv6Info {
  s.PrivateIp = &v
  return s
}

type LECHCreateInstanceRequestServersDiskInfo struct     {
  // {"en":"Disk size (GB)","zh_CN":"磁盘大小（GB）"}
  Size *int `json:"size,omitempty" xml:"size,omitempty"`
  // {"en":"Disk Purpose:\nSystem - System disk;\nDATA - DATA plate","zh_CN":"磁盘用途：\nSYSTEM-系统盘；\nDATA-数据盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Disk type: HDD/SSD","zh_CN":"磁盘类型：HDD/SSD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // {"en":"Is isolate: 1(Yes) / -1(No)","zh_CN":"是否独立盘：1(是) / -1(否)"}
  IsIndependent *string `json:"isIndependent,omitempty" xml:"isIndependent,omitempty"`
}

func (s LECHCreateInstanceRequestServersDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequestServersDiskInfo) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceRequestServersDiskInfo) SetSize(v int) *LECHCreateInstanceRequestServersDiskInfo {
  s.Size = &v
  return s
}

func (s *LECHCreateInstanceRequestServersDiskInfo) SetType(v string) *LECHCreateInstanceRequestServersDiskInfo {
  s.Type = &v
  return s
}

func (s *LECHCreateInstanceRequestServersDiskInfo) SetCategory(v string) *LECHCreateInstanceRequestServersDiskInfo {
  s.Category = &v
  return s
}

func (s *LECHCreateInstanceRequestServersDiskInfo) SetIsIndependent(v string) *LECHCreateInstanceRequestServersDiskInfo {
  s.IsIndependent = &v
  return s
}

type LECHCreateInstanceRequestServersPublicIpv4Info struct     {
  // {"en":"ISP abbreviation format, eg: dx,wt,yd","zh_CN":"运营商缩写格式：dx-电信；yd-移动；wt-网通"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
}

func (s LECHCreateInstanceRequestServersPublicIpv4Info) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequestServersPublicIpv4Info) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceRequestServersPublicIpv4Info) SetCarrier(v string) *LECHCreateInstanceRequestServersPublicIpv4Info {
  s.Carrier = &v
  return s
}

type LECHCreateInstanceRequestServersPublicIpv6Info struct     {
  // {"en":"ISP abbreviation format, eg: dx,wt,yd","zh_CN":"运营商缩写格式：dx-电信；yd-移动；wt-网通"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
}

func (s LECHCreateInstanceRequestServersPublicIpv6Info) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequestServersPublicIpv6Info) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceRequestServersPublicIpv6Info) SetCarrier(v string) *LECHCreateInstanceRequestServersPublicIpv6Info {
  s.Carrier = &v
  return s
}

type LECHCreateInstanceRequestHeader struct {
}

func (s LECHCreateInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceRequestHeader) GoString() string {
  return s.String()
}

type LECHCreateInstancePaths struct {
}

func (s LECHCreateInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstancePaths) GoString() string {
  return s.String()
}

type LECHCreateInstanceParameters struct {
}

func (s LECHCreateInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceParameters) GoString() string {
  return s.String()
}

type LECHCreateInstanceResponse struct {
  // {"en":"Virtual machine identity list","zh_CN":"实例标识列表"}
  Id []*string `json:"id,omitempty" xml:"id,omitempty" require:"true" type:"Repeated"`
}

func (s LECHCreateInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceResponse) GoString() string {
  return s.String()
}

func (s *LECHCreateInstanceResponse) SetId(v []*string) *LECHCreateInstanceResponse {
  s.Id = v
  return s
}

type LECHCreateInstanceResponseHeader struct {
}

func (s LECHCreateInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateInstanceResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryInstanceRequest struct {
}

func (s VMPQueryInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceRequest) GoString() string {
  return s.String()
}

type VMPQueryInstanceAccessIP struct {
  // {"en":"Ip address", "zh_CN":"Ip地址"}
  Address *string `json:"address,omitempty" xml:"address,omitempty" require:"true"`
  // {"en":"IP address operator", "zh_CN":"Ip地址所属运营商"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Protocol type: 4: IPv4 address; 6: IPv6 address", "zh_CN":"协议类型：4：ipv4地址；6：ipv6地址"}
  Protocol *int `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s VMPQueryInstanceAccessIP) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceAccessIP) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceAccessIP) SetAddress(v string) *VMPQueryInstanceAccessIP {
  s.Address = &v
  return s
}

func (s *VMPQueryInstanceAccessIP) SetCarrier(v string) *VMPQueryInstanceAccessIP {
  s.Carrier = &v
  return s
}

func (s *VMPQueryInstanceAccessIP) SetProtocol(v int) *VMPQueryInstanceAccessIP {
  s.Protocol = &v
  return s
}

type VMPQueryInstanceDiskInfo struct {
  // {"en":"Disk size (GB)", "zh_CN":"磁盘大小（GB）"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk Purpose: System - System disk;DATA - DATA plate", "zh_CN":"磁盘用途：SYSTEM-系统盘；DATA-数据盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk type: HDD/SSD", "zh_CN":"磁盘类型：HDD/SSD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s VMPQueryInstanceDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceDiskInfo) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceDiskInfo) SetSize(v int) *VMPQueryInstanceDiskInfo {
  s.Size = &v
  return s
}

func (s *VMPQueryInstanceDiskInfo) SetType(v string) *VMPQueryInstanceDiskInfo {
  s.Type = &v
  return s
}

func (s *VMPQueryInstanceDiskInfo) SetCategory(v string) *VMPQueryInstanceDiskInfo {
  s.Category = &v
  return s
}

type VMPQueryInstanceServer struct {
  // {"en":"Unique identity of virtual machine", "zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Virtual Machine Area", "zh_CN":"实例所属区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"Province of virtual machine", "zh_CN":"实例所属省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Operator of virtual machine", "zh_CN":"实例所属运营商"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Virtual machine image information", "zh_CN":"实例镜像信息"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Virtual machine specifications", "zh_CN":"实例规格"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Virtual machine name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Virtual machine status", "zh_CN":"实例状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Virtual machine creation time", "zh_CN":"实例创建时间"}
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty" require:"true"`
  // {"en":"Virtual machine IP address", "zh_CN":"实例IP地址"}
  AccessIPv4 *string `json:"accessIPv4,omitempty" xml:"accessIPv4,omitempty" require:"true"`
  // {"en":"IPv4 address of virtual machine intranet. If intranet is specified when creating virtual machine, intranet IP is returned. Otherwise, it is empty.", "zh_CN":"实例内网IPv4地址，如果创建实例时指定了需要内网，则返回内网IPv4，否则是空"}
  PrivateIPv4 *string `json:"privateIPv4,omitempty" xml:"privateIPv4,omitempty" require:"true"`
  // {"en":"IPv6 address of virtual machine intranet. If intranet is specified when creating virtual machine, intranet IP is returned. Otherwise, it is empty.", "zh_CN":"实例内网IPv6地址，如果创建实例时指定了需要内网，则返回内网IPv6，否则是空"}
  PrivateIPv6 *string `json:"privateIPv6,omitempty" xml:"privateIPv6,omitempty" require:"true"`
  // {"en":"Virtual machine login SSH key pair name", "zh_CN":"实例登录SSH秘钥对名称"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty" require:"true"`
  // {"en":"Node name, the name of the node where the virtual machine is located. Through this node name, you can query the real-time redundant bandwidth of each node.", "zh_CN":"节点名称，实例所在节点名称，通过这个节点名称可以查询每个节点的实时冗余带宽情况"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Virtual machine IP address", "zh_CN":"实例ip地址"}
  VMPQueryInstanceAccessIP []*VMPQueryInstanceAccessIP `json:"accessIP,omitempty" xml:"accessIP,omitempty" require:"true" type:"Repeated"`
  // {"en":"If it is free instance, value: YES free instance, NO billing instance", "zh_CN":"是否免费实例，取值：YES 免费实例，NO 计费实例"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty" require:"true"`
  // {"en":"1 indicates that the instance is bare metal
  // -1 indicates that the instance is a virtual machine", "zh_CN":"1表示该实例是裸机
  // -1表示该实例是虚拟机"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"A list of security group IDs for instance bindings", "zh_CN":"实例绑定的安全组id列表"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" require:"true" type:"Repeated"`
  // {"en":"Disk information", "zh_CN":"磁盘信息"}
  VMPQueryInstanceDiskInfo []*VMPQueryInstanceDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
  // {"en":"instance tag", "zh_CN":"实例标签"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty" require:"true"`
}

func (s VMPQueryInstanceServer) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceServer) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceServer) SetId(v string) *VMPQueryInstanceServer {
  s.Id = &v
  return s
}

func (s *VMPQueryInstanceServer) SetRegionName(v string) *VMPQueryInstanceServer {
  s.RegionName = &v
  return s
}

func (s *VMPQueryInstanceServer) SetProvince(v string) *VMPQueryInstanceServer {
  s.Province = &v
  return s
}

func (s *VMPQueryInstanceServer) SetCarrier(v string) *VMPQueryInstanceServer {
  s.Carrier = &v
  return s
}

func (s *VMPQueryInstanceServer) SetImageId(v string) *VMPQueryInstanceServer {
  s.ImageId = &v
  return s
}

func (s *VMPQueryInstanceServer) SetFlavorId(v string) *VMPQueryInstanceServer {
  s.FlavorId = &v
  return s
}

func (s *VMPQueryInstanceServer) SetName(v string) *VMPQueryInstanceServer {
  s.Name = &v
  return s
}

func (s *VMPQueryInstanceServer) SetState(v string) *VMPQueryInstanceServer {
  s.State = &v
  return s
}

func (s *VMPQueryInstanceServer) SetCreatedAt(v string) *VMPQueryInstanceServer {
  s.CreatedAt = &v
  return s
}

func (s *VMPQueryInstanceServer) SetAccessIPv4(v string) *VMPQueryInstanceServer {
  s.AccessIPv4 = &v
  return s
}

func (s *VMPQueryInstanceServer) SetPrivateIPv4(v string) *VMPQueryInstanceServer {
  s.PrivateIPv4 = &v
  return s
}

func (s *VMPQueryInstanceServer) SetPrivateIPv6(v string) *VMPQueryInstanceServer {
  s.PrivateIPv6 = &v
  return s
}

func (s *VMPQueryInstanceServer) SetKeyName(v string) *VMPQueryInstanceServer {
  s.KeyName = &v
  return s
}

func (s *VMPQueryInstanceServer) SetNodeName(v string) *VMPQueryInstanceServer {
  s.NodeName = &v
  return s
}

func (s *VMPQueryInstanceServer) SetAccessIP(v []*VMPQueryInstanceAccessIP) *VMPQueryInstanceServer {
  s.VMPQueryInstanceAccessIP = v
  return s
}

func (s *VMPQueryInstanceServer) SetIsFree(v string) *VMPQueryInstanceServer {
  s.IsFree = &v
  return s
}

func (s *VMPQueryInstanceServer) SetIsBm(v int) *VMPQueryInstanceServer {
  s.IsBm = &v
  return s
}

func (s *VMPQueryInstanceServer) SetSecurityGroupIds(v []*string) *VMPQueryInstanceServer {
  s.SecurityGroupIds = v
  return s
}

func (s *VMPQueryInstanceServer) SetDiskInfo(v []*VMPQueryInstanceDiskInfo) *VMPQueryInstanceServer {
  s.VMPQueryInstanceDiskInfo = v
  return s
}

func (s *VMPQueryInstanceServer) SetTag(v string) *VMPQueryInstanceServer {
  s.Tag = &v
  return s
}

type VMPQueryInstanceResponse struct {
  // {"en":"Virtual machine information array", "zh_CN":"实例信息数组"}
  Servers []*VMPQueryInstanceServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceResponse) SetServers(v []*VMPQueryInstanceServer) *VMPQueryInstanceResponse {
  s.Servers = v
  return s
}

type VMPQueryInstancePaths struct {
}

func (s VMPQueryInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstancePaths) GoString() string {
  return s.String()
}

type VMPQueryInstanceParameters struct {
  // {"en":"Sort field name, can have more than one, value:
  // Name, regionName, createdAt, type, state, nodeName, etc", "zh_CN":"排序的字段名称，可以有多个，取值：
  // name、regionName、createdAt、province、state、nodeName等"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"There can be multiple field names for sorting, with values:
  // 
  // Name, regionname, createdat, province, state, nodeName, etc.'", "zh_CN":"排序方向，必须跟在sortKey后面出现，取值：
  // desc：降序，默认值
  // asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the virtual machine ID specified by the marker", "zh_CN":"从marker指定的实例id开始查询，升序查询（若要分页查询，则不能指定sortKey参数）"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Virtual machine ID. A maximum of 100 IDS can be queried at a time. The IDs are separated by a half angle comma character ','.", "zh_CN":"实例ID。单次最多查询 100 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"Name of the region, for example, South China region is Huanan, refer to Appendix 3", "zh_CN":"区域名称（区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Province of virtual machine", "zh_CN":"实例所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Operator of virtual machine", "zh_CN":"实例所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Virtual machine image identity", "zh_CN":"实例镜像标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
  // {"en":"Virtual machine specification ID", "zh_CN":"实例规格标识"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty"`
  // {"en":"Virtual machine name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Virtual machine ips. A maximum of 100 IPS can be queried at a time. The IPs are separated by a half angle comma character ','.", "zh_CN":"实例IP。单次最多查询 100 条 IP，IP 之间用半角逗号字符','隔开。"}
  Ips *string `json:"ips,omitempty" xml:"ips,omitempty"`
  // {"en":"Virtual machine status, value (meaning to query virtual machines in the following status)
  // 
  // Running status
  // 
  // Building new status
  // 
  // Stopped stop
  // 
  // ERROR error
  // 
  // Deleting destroying
  // 
  // Restarting
  // 
  // Starting
  // 
  // Stopping", "zh_CN":"实例状态，取值（意为查询处于下述状态的虚拟机）
  // RUNNING 运行状态
  // BUILDING 新建状态
  // STOPPED 停机
  // ERROR 错误
  // DELETING 销毁中
  // RESTARTING  重启中
  // STARTING  启动中
  // STOPPING  停止中"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Whether it is a free instance, value:
  // Yes free instances, NO billed instances", "zh_CN":"是否免费实例，取值:
  // YES 免费实例，NO 计费实例"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty"`
  // {"en":"1 means to query only bare-metal instances,
  // -1 means that only virtual machine instances are queried,
  // Without this parameter all cloud hosts are queried", "zh_CN":"1表示只查询裸机实例，
  // -1表示只查询虚拟机实例，
  // 不带这个参数表示查询所有云主机"}
  IsBm *string `json:"isBm,omitempty" xml:"isBm,omitempty"`
  // {"en":"Cloud Host Label
  // Multiple values are separated by a half corner comma, and the relationship between multiple values is or, that is, the instance label equals any one of these multiple values", "zh_CN":"云主机标签
  // 多个值用半角逗号隔开，多个值是或者的关系，即实例标签等于这多个中的任意一个就满足条件"}
  Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
  // {"en":"ipv6 formate: 1: Zero Compressed(default); 6: With Leading Zero Suppression; 3: Full Address", "zh_CN":"ipv6格式：1：省略零压缩格式(默认)；2：省略前导零格式; 3: 完整格式"}
  Ipv6Format *string `json:"ipv6Format,omitempty" xml:"ipv6Format,omitempty"`
}

func (s VMPQueryInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceParameters) SetSortKey(v string) *VMPQueryInstanceParameters {
  s.SortKey = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetSortDir(v string) *VMPQueryInstanceParameters {
  s.SortDir = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetLimit(v int) *VMPQueryInstanceParameters {
  s.Limit = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetMarker(v string) *VMPQueryInstanceParameters {
  s.Marker = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIds(v string) *VMPQueryInstanceParameters {
  s.Ids = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetRegionName(v string) *VMPQueryInstanceParameters {
  s.RegionName = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetProvince(v string) *VMPQueryInstanceParameters {
  s.Province = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetCarrier(v string) *VMPQueryInstanceParameters {
  s.Carrier = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetImageId(v string) *VMPQueryInstanceParameters {
  s.ImageId = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetFlavorId(v string) *VMPQueryInstanceParameters {
  s.FlavorId = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetName(v string) *VMPQueryInstanceParameters {
  s.Name = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIps(v string) *VMPQueryInstanceParameters {
  s.Ips = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetState(v string) *VMPQueryInstanceParameters {
  s.State = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIsFree(v string) *VMPQueryInstanceParameters {
  s.IsFree = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIsBm(v string) *VMPQueryInstanceParameters {
  s.IsBm = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetTags(v string) *VMPQueryInstanceParameters {
  s.Tags = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIpv6Format(v string) *VMPQueryInstanceParameters {
  s.Ipv6Format = &v
  return s
}

type VMPQueryInstanceRequestHeader struct {
}

func (s VMPQueryInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryInstanceResponseHeader struct {
}

func (s VMPQueryInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceResponseHeader) GoString() string {
  return s.String()
}




type InstanceIpv6ManagementRequest struct {
  // {"en":"Operation:\nALLOCATION - ipv6 application\nREMOVE - ipv6 is removed","zh_CN":"操作：\nALLOCATION-ipv6申请\nREMOVE-ipv6移除"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Instance ID (single only)","zh_CN":"实例id（只支持单个）"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
}

func (s InstanceIpv6ManagementRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementRequest) GoString() string {
  return s.String()
}

func (s *InstanceIpv6ManagementRequest) SetAction(v string) *InstanceIpv6ManagementRequest {
  s.Action = &v
  return s
}

func (s *InstanceIpv6ManagementRequest) SetInstanceId(v string) *InstanceIpv6ManagementRequest {
  s.InstanceId = &v
  return s
}

type InstanceIpv6ManagementRequestHeader struct {
}

func (s InstanceIpv6ManagementRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementRequestHeader) GoString() string {
  return s.String()
}

type InstanceIpv6ManagementPaths struct {
}

func (s InstanceIpv6ManagementPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementPaths) GoString() string {
  return s.String()
}

type InstanceIpv6ManagementParameters struct {
}

func (s InstanceIpv6ManagementParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementParameters) GoString() string {
  return s.String()
}

type InstanceIpv6ManagementResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *InstanceIpv6ManagementResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s InstanceIpv6ManagementResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementResponse) GoString() string {
  return s.String()
}

func (s *InstanceIpv6ManagementResponse) SetCode(v string) *InstanceIpv6ManagementResponse {
  s.Code = &v
  return s
}

func (s *InstanceIpv6ManagementResponse) SetMessage(v string) *InstanceIpv6ManagementResponse {
  s.Message = &v
  return s
}

func (s *InstanceIpv6ManagementResponse) SetData(v *InstanceIpv6ManagementResponseData) *InstanceIpv6ManagementResponse {
  s.Data = v
  return s
}

type InstanceIpv6ManagementResponseData struct {
  // {"en":"Instance ID","zh_CN":"实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"IPv6","zh_CN":"IPv6"}
  AccessIPv6 []*string `json:"accessIPv6,omitempty" xml:"accessIPv6,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceIpv6ManagementResponseData) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementResponseData) GoString() string {
  return s.String()
}

func (s *InstanceIpv6ManagementResponseData) SetInstanceId(v string) *InstanceIpv6ManagementResponseData {
  s.InstanceId = &v
  return s
}

func (s *InstanceIpv6ManagementResponseData) SetAccessIPv6(v []*string) *InstanceIpv6ManagementResponseData {
  s.AccessIPv6 = v
  return s
}

type InstanceIpv6ManagementResponseHeader struct {
}

func (s InstanceIpv6ManagementResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceIpv6ManagementResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceRebuildRequest struct {
  // {"en":"vm id","zh_CN":"云主机ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Image ID","zh_CN":"镜像ID（指定了镜像ID则使用指定的镜像重装，否则使用原镜像重装）"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
  // {"en":"password","zh_CN":"密码（使用公共镜像重装必须指定密码，使用自定义镜像可不指定）"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"Retain Data Disk","zh_CN":"是否保留数据盘（1：是；-1：否）"}
  RetainDataDisk *int `json:"retainDataDisk,omitempty" xml:"retainDataDisk,omitempty"`
}

func (s LECHInstanceRebuildRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceRebuildRequest) GoString() string {
  return s.String()
}

func (s *LECHInstanceRebuildRequest) SetId(v string) *LECHInstanceRebuildRequest {
  s.Id = &v
  return s
}

func (s *LECHInstanceRebuildRequest) SetImageId(v string) *LECHInstanceRebuildRequest {
  s.ImageId = &v
  return s
}

func (s *LECHInstanceRebuildRequest) SetPassword(v string) *LECHInstanceRebuildRequest {
  s.Password = &v
  return s
}

func (s *LECHInstanceRebuildRequest) SetRetainDataDisk(v int) *LECHInstanceRebuildRequest {
  s.RetainDataDisk = &v
  return s
}

type LECHInstanceRebuildRequestHeader struct {
}

func (s LECHInstanceRebuildRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceRebuildRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceRebuildPaths struct {
}

func (s LECHInstanceRebuildPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceRebuildPaths) GoString() string {
  return s.String()
}

type LECHInstanceRebuildParameters struct {
}

func (s LECHInstanceRebuildParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceRebuildParameters) GoString() string {
  return s.String()
}

type LECHInstanceRebuildResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHInstanceRebuildResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceRebuildResponse) GoString() string {
  return s.String()
}

func (s *LECHInstanceRebuildResponse) SetCode(v string) *LECHInstanceRebuildResponse {
  s.Code = &v
  return s
}

func (s *LECHInstanceRebuildResponse) SetMessage(v string) *LECHInstanceRebuildResponse {
  s.Message = &v
  return s
}

type LECHInstanceRebuildResponseHeader struct {
}

func (s LECHInstanceRebuildResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceRebuildResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceReplaceIpRequest struct {
  // {"en":"vm id","zh_CN":"云主机ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"server old replace ips","zh_CN":"虚拟机待更换ip列表"}
  InstanceIps []*string `json:"instanceIps,omitempty" xml:"instanceIps,omitempty" require:"true" type:"Repeated"`
}

func (s LECHInstanceReplaceIpRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpRequest) GoString() string {
  return s.String()
}

func (s *LECHInstanceReplaceIpRequest) SetInstanceId(v string) *LECHInstanceReplaceIpRequest {
  s.InstanceId = &v
  return s
}

func (s *LECHInstanceReplaceIpRequest) SetInstanceIps(v []*string) *LECHInstanceReplaceIpRequest {
  s.InstanceIps = v
  return s
}

type LECHInstanceReplaceIpRequestHeader struct {
}

func (s LECHInstanceReplaceIpRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceReplaceIpPaths struct {
}

func (s LECHInstanceReplaceIpPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpPaths) GoString() string {
  return s.String()
}

type LECHInstanceReplaceIpParameters struct {
}

func (s LECHInstanceReplaceIpParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpParameters) GoString() string {
  return s.String()
}

type LECHInstanceReplaceIpResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHInstanceReplaceIpResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHInstanceReplaceIpResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpResponse) GoString() string {
  return s.String()
}

func (s *LECHInstanceReplaceIpResponse) SetCode(v string) *LECHInstanceReplaceIpResponse {
  s.Code = &v
  return s
}

func (s *LECHInstanceReplaceIpResponse) SetMessage(v string) *LECHInstanceReplaceIpResponse {
  s.Message = &v
  return s
}

func (s *LECHInstanceReplaceIpResponse) SetData(v *LECHInstanceReplaceIpResponseData) *LECHInstanceReplaceIpResponse {
  s.Data = v
  return s
}

type LECHInstanceReplaceIpResponseData struct {
}

func (s LECHInstanceReplaceIpResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpResponseData) GoString() string {
  return s.String()
}

type LECHInstanceReplaceIpResponseHeader struct {
}

func (s LECHInstanceReplaceIpResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceReplaceIpResponseHeader) GoString() string {
  return s.String()
}




type GetEcciInstanceListRequest struct {
}

func (s GetEcciInstanceListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListRequest) GoString() string {
  return s.String()
}

type GetEcciInstanceListResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ecci instance list", "zh_CN":"ecci实例列表"}
  Data *GetEcciInstanceListInstances `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetEcciInstanceListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListResponse) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceListResponse) SetCode(v int64) *GetEcciInstanceListResponse {
  s.Code = &v
  return s
}

func (s *GetEcciInstanceListResponse) SetMsg(v string) *GetEcciInstanceListResponse {
  s.Msg = &v
  return s
}

func (s *GetEcciInstanceListResponse) SetRequestId(v string) *GetEcciInstanceListResponse {
  s.RequestId = &v
  return s
}

func (s *GetEcciInstanceListResponse) SetData(v *GetEcciInstanceListInstances) *GetEcciInstanceListResponse {
  s.Data = v
  return s
}

type GetEcciInstanceListPaths struct {
}

func (s GetEcciInstanceListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListPaths) GoString() string {
  return s.String()
}

type GetEcciInstanceListParameters struct {
  // {"en":"Keyword: instance name/id/ip", "zh_CN":"关键字: 实例名称/id/ip"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"page index", "zh_CN":"页数"}
  PageIndex *int64 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Number per page", "zh_CN":"每页个数"}
  PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetEcciInstanceListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListParameters) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceListParameters) SetKey(v string) *GetEcciInstanceListParameters {
  s.Key = &v
  return s
}

func (s *GetEcciInstanceListParameters) SetPageIndex(v int64) *GetEcciInstanceListParameters {
  s.PageIndex = &v
  return s
}

func (s *GetEcciInstanceListParameters) SetPageSize(v int64) *GetEcciInstanceListParameters {
  s.PageSize = &v
  return s
}

type GetEcciInstanceListRequestHeader struct {
}

func (s GetEcciInstanceListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListRequestHeader) GoString() string {
  return s.String()
}

type GetEcciInstanceListResponseHeader struct {
}

func (s GetEcciInstanceListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListResponseHeader) GoString() string {
  return s.String()
}

type GetEcciInstanceListInstances struct {
  // {"en":"ecci instance", "zh_CN":"ecci 实例"}
  GetEcciInstanceListInstances []*GetEcciInstanceListEcciInstanceSummary `json:"instances,omitempty" xml:"instances,omitempty" require:"true" type:"Repeated"`
  // {"en":"total instance", "zh_CN":"实例总数量"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s GetEcciInstanceListInstances) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListInstances) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceListInstances) SetInstances(v []*GetEcciInstanceListEcciInstanceSummary) *GetEcciInstanceListInstances {
  s.GetEcciInstanceListInstances = v
  return s
}

func (s *GetEcciInstanceListInstances) SetTotal(v int64) *GetEcciInstanceListInstances {
  s.Total = &v
  return s
}

type GetEcciInstanceListEcciInstanceSummary struct {
  // {"en":"Instance id", "zh_CN":"实例d"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance display name", "zh_CN":"实例展示名称"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"private ip list", "zh_CN":"内网ip列表"}
  PrivateIps []*string `json:"privateIps,omitempty" xml:"privateIps,omitempty" require:"true" type:"Repeated"`
  // {"en":"cluster name", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty" require:"true"`
  // {"en":"instance cpu limit", "zh_CN":"实例cpu限制"}
  CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty" require:"true"`
  // {"en":"Minimum amount of cpu resources required", "zh_CN":"cpu 要求最小资源量"}
  CpuRequest *string `json:"cpuRequest,omitempty" xml:"cpuRequest,omitempty" require:"true"`
  // {"en":"instance memory limit", "zh_CN":"实例内存限制"}
  MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty" require:"true"`
  // {"en":"Minimum amount of memory resources required", "zh_CN":"内存要求最小资源量"}
  MemoryRequest *string `json:"memoryRequest,omitempty" xml:"memoryRequest,omitempty" require:"true"`
  // {"en":"pod status", "zh_CN":"pod状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"create time ", "zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"area", "zh_CN":"区域"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"country", "zh_CN":"国家"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"city", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"pod name", "zh_CN":"pod名称"}
  PodName *string `json:"podName,omitempty" xml:"podName,omitempty" require:"true"`
  // {"en":"IP operator list", "zh_CN":"ip运营商列表"}
  Isps []*GetEcciInstanceListIspIp `json:"isps,omitempty" xml:"isps,omitempty" require:"true" type:"Repeated"`
  // {"en":"is colocation resource or not", "zh_CN":"是否在离线混部资源"}
  Colocation *bool `json:"colocation,omitempty" xml:"colocation,omitempty" require:"true"`
}

func (s GetEcciInstanceListEcciInstanceSummary) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListEcciInstanceSummary) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetId(v int64) *GetEcciInstanceListEcciInstanceSummary {
  s.Id = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetName(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetDisplayName(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.DisplayName = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetPrivateIps(v []*string) *GetEcciInstanceListEcciInstanceSummary {
  s.PrivateIps = v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetClusterName(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.ClusterName = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetCpuLimit(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.CpuLimit = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetCpuRequest(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.CpuRequest = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetMemoryLimit(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.MemoryLimit = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetMemoryRequest(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.MemoryRequest = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetStatus(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.Status = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetCreateTime(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.CreateTime = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetArea(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.Area = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetCountry(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.Country = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetProvince(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.Province = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetCity(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.City = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetPodName(v string) *GetEcciInstanceListEcciInstanceSummary {
  s.PodName = &v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetIsps(v []*GetEcciInstanceListIspIp) *GetEcciInstanceListEcciInstanceSummary {
  s.Isps = v
  return s
}

func (s *GetEcciInstanceListEcciInstanceSummary) SetColocation(v bool) *GetEcciInstanceListEcciInstanceSummary {
  s.Colocation = &v
  return s
}

type GetEcciInstanceListIspIp struct {
  // {"en":"operator id", "zh_CN":"运营商id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Operator Chinese name", "zh_CN":"运营商中文名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Operator english name", "zh_CN":"运营商英文名"}
  NameEn *string `json:"nameEn,omitempty" xml:"nameEn,omitempty" require:"true"`
  // {"en":"ipv4 ip", "zh_CN":"ipv4 ip"}
  Ipv4 *string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true"`
  // {"en":"ipv6 ip", "zh_CN":"ipv6 ip"}
  Ipv6 *string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true"`
}

func (s GetEcciInstanceListIspIp) String() string {
  return tea.Prettify(s)
}

func (s GetEcciInstanceListIspIp) GoString() string {
  return s.String()
}

func (s *GetEcciInstanceListIspIp) SetId(v int64) *GetEcciInstanceListIspIp {
  s.Id = &v
  return s
}

func (s *GetEcciInstanceListIspIp) SetName(v string) *GetEcciInstanceListIspIp {
  s.Name = &v
  return s
}

func (s *GetEcciInstanceListIspIp) SetNameEn(v string) *GetEcciInstanceListIspIp {
  s.NameEn = &v
  return s
}

func (s *GetEcciInstanceListIspIp) SetIpv4(v string) *GetEcciInstanceListIspIp {
  s.Ipv4 = &v
  return s
}

func (s *GetEcciInstanceListIspIp) SetIpv6(v string) *GetEcciInstanceListIspIp {
  s.Ipv6 = &v
  return s
}




type VMPCreateInstanceRequest struct {
  // {"en":"Creating array objects for virtual machines","zh_CN":"创建实例的数组对象"}
  Servers []*VMPCreateInstanceRequestServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequest) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequest) SetServers(v []*VMPCreateInstanceRequestServers) *VMPCreateInstanceRequest {
  s.Servers = v
  return s
}

type VMPCreateInstanceRequestServers struct     {
  // {"en":"Virtual machine area (see Appendix for details)","zh_CN":"实例所属区域（节点名称nodeName和区域regionName至少需要上传一个。\n区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Province of virtual machine (see Appendix for details)","zh_CN":"实例所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"If the operator of the virtual machine (see the appendix for details) carries this parameter, please keep it consistent with the carrier returned from the '3.4 node list query' interface.","zh_CN":"实例所属运营商（dx-电信；wt-网通；yd-移动）如果携带了该参数，请与'3.4节点列表查询'接口返回的carrier保持一致"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node name, indicating that the specified node creates a virtual machine (the node name returned by interface 3.4)","zh_CN":"节点名称，表示指定节点创建实例（节点名称可通过资源管理-节点列表查询接口获取）"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"Virtual machine image identity","zh_CN":"实例镜像标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Virtual machine specification ID","zh_CN":"实例规格标识"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Virtual machine name. If the created quantity is greater than 1, the real name is spliced with 3 digits after the parameter. For example, instance 0001, instance 0002","zh_CN":"实例名称，如果创建数量大于1，则真实名称是在该参数后拼接3位数字。如instance_0001，instance_0002"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Inject user data, support to inject text, text file or gzip file. The maximum length of injected content is 32KB. For content injection, Base64 format encoding is required.","zh_CN":"注入用户数据，支持注入文本、文本文件或gzip文件。注入内容最大长度32KB。注入内容，需要进行base64格式编码。"}
  UserData *string `json:"userData,omitempty" xml:"userData,omitempty"`
  // {"en":"Number of virtual machines applied","zh_CN":"申请实例数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty"`
  // {"en":"Virtual machine root login password","zh_CN":"实例root用户登录密码（如果选择的是公共镜像，则密码password必填）。\n密码规则：大写字母 小写字母 数字 特殊字符，四种包括三种，长度8~30"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"The name of the SSH secret key pair for virtual machine login. If this parameter is specified, the password login mode is disabled by default, and the password parameter is invalid at the same time.","zh_CN":"实例登录SSH秘钥对名称，如果指定该参数，默认禁用密码登录方式，password参数同时失效"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
  // {"en":"Whether the virtual machine needs intranet, value:\nYes: intranet required\nNo: no intranet is required, default value'","zh_CN":"实例是否需要内网网络，取值：\nYES：需要内网\nNO：不需要内网，默认值"}
  InnerNet *string `json:"innerNet,omitempty" xml:"innerNet,omitempty"`
  // {"en":"CIDR of virtual machine intranet is meaningful only when innernet = yes","zh_CN":"实例内网的cidr，只有innerNet=YES时才有意义"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"If IP address is specified, it must be within the scope of CIDR, otherwise creation fails.","zh_CN":"实例内网ip地址，如果指定了ip，必须在cidr的范围内，否则创建失败"}
  PrivateIPv4 *string `json:"privateIPv4,omitempty" xml:"privateIPv4,omitempty"`
  // {"en":"Whether the virtual machine needs intranet2, value:\nYes: intranet2 required\nNo: no intranet2 is required, default value'","zh_CN":"实例是否需要内网2网络，取值：\nYES：需要内网2\nNO：不需要内网2，默认值"}
  InnerNet2 *string `json:"innerNet2,omitempty" xml:"innerNet2,omitempty"`
  // {"en":"CIDR of virtual machine intranet2 is meaningful only when innernet = yes","zh_CN":"实例内网2的cidr，只有innerNet2=YES时才有意义"}
  Cidr2 *string `json:"cidr2,omitempty" xml:"cidr2,omitempty"`
  // {"en":"If IP address is specified, it must be within the scope of CIDR2, otherwise creation fails.","zh_CN":"实例内网2ip地址，如果指定了ip，必须在cidr2的范围内，否则创建失败"}
  PrivateIPv42 *string `json:"privateIPv42,omitempty" xml:"privateIPv42,omitempty"`
  // {"en":"Inner ipv6 info","zh_CN":"内网IPv6信息"}
  PrivateIpv6Info []*VMPCreateInstanceRequestServersPrivateIpv6Info `json:"privateIpv6Info,omitempty" xml:"privateIpv6Info,omitempty" type:"Repeated"`
  // {"en":"Whether multiple IP protocol addresses are required\n\n4: only IPv4 address is required, default value\n\n0: both IPv4 and IPv6 need'","zh_CN":"是否需要多ip协议地址\n4：只需要ipv4地址，默认值\n0：ipv4、ipv6都需要"}
  Protocols *int `json:"protocols,omitempty" xml:"protocols,omitempty"`
  // {"en":"IPv4 native attribute, 1: non-native;-1: native;","zh_CN":"IPv4原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  Ipv4NativeAttribute *string `json:"ipv4NativeAttribute,omitempty" xml:"ipv4NativeAttribute,omitempty"`
  // {"en":"IPv6 native attribute, 1: non-native;-1: native;","zh_CN":"IPv6原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  Ipv6NativeAttribute *string `json:"ipv6NativeAttribute,omitempty" xml:"ipv6NativeAttribute,omitempty"`
  // {"en":"Whether the instance is free or not, the default billing instance, and the bare machine instance cannot be free, values are as follows:\nYes: Free instances\nNo: Billing instance\nIf you are using a free instance, you need to configure permissions in advance","zh_CN":"是否免费实例，默认计费实例，裸机实例不能免费，取值：\nYES：免费实例\nNO：计费实例\n如果使用免费实例，需要提前配置权限"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty"`
  // {"en":"Specify a security group ID to create multiple security groups separated by commas, up to 5\nIf you are creating a bare machine, you cannot specify a security group","zh_CN":"指定安全组id进行创建，多个安全组以逗号分隔，最多指定5个\n如果是创建裸机，不能指定安全组"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
  // {"en":"Disk information\nIf this information is carried, the disk definition on the template will be ignored and the instance disk will be created with this information, not for bare-metal instance creation","zh_CN":"磁盘信息\n如果携带该信息，将忽略模板上的磁盘定义，以该信息创建实例磁盘，不适用于裸机实例创建"}
  DiskInfo []*VMPCreateInstanceRequestServersDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" type:"Repeated"`
  // {"en":"Anti-affinity group name\nVirtual machines with the same ServerGroup are created on different hosts","zh_CN":"反亲和性组名称\n拥有相同serverGroup的虚拟机会被创建在不同的宿主机上"}
  ServerGroup *string `json:"serverGroup,omitempty" xml:"serverGroup,omitempty"`
  // {"en":"Instance Tag","zh_CN":"实例标签"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
  // {"en":"Use  unique ip segment","zh_CN":"是否使用唯一网段\n1：是\n-1：否"}
  UseUniqueIpSegment *int `json:"useUniqueIpSegment,omitempty" xml:"useUniqueIpSegment,omitempty"`
  // {"en":"Allocate IP randomly","zh_CN":"是否需要随机分配IPv4\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
  // {"en":"Default Gateway","zh_CN":"默认网关运营商如：dx-电信；yd-移动；wt-网通"}
  DefaultGateway *string `json:"defaultGateway,omitempty" xml:"defaultGateway,omitempty"`
  // {"en":"Policy routing type","zh_CN":"策略路由类型：0-目的地址策略路由（默认）；1-源地址策略路由"}
  PolicyRoutingType *int `json:"policyRoutingType,omitempty" xml:"policyRoutingType,omitempty"`
  // {"en":"Private gateway flag","zh_CN":"内网网关标识：1-分配内网网关"}
  PrivateGatewayFlag *int `json:"privateGatewayFlag,omitempty" xml:"privateGatewayFlag,omitempty"`
  // {"en":"Nic allocate type","zh_CN":"实例网卡分配方式：0-多个ip共用一张网卡（默认）；1-每个ip独立一张网卡；2-V4V6混合，同协议IP同网卡，不同线路IP不同网卡"}
  NicAllocateType *int `json:"nicAllocateType,omitempty" xml:"nicAllocateType,omitempty"`
  // {"en":"Ipv4 cidr","zh_CN":"指定外网IPv4网段CIDR(不支持多线)"}
  SinglePublicIpv4Cidr *string `json:"singlePublicIpv4Cidr,omitempty" xml:"singlePublicIpv4Cidr,omitempty"`
  // {"en":"Specify certain public IPv4 ISPs; effective for multi-line nodes; if not specified, an instance with IPs from all carriers is created by default.","zh_CN":"指定部分公网ipv4运营商, 仅多线节点生效。多线节点未指定该参数时，默认创建包含所有运营商IP的实例"}
  PublicIpv4Info []*VMPCreateInstanceRequestServersPublicIpv4Info `json:"publicIpv4Info,omitempty" xml:"publicIpv4Info,omitempty" type:"Repeated"`
  // {"en":"Specify certain public IPv6 ISPs; effective for multi-line nodes; if not specified, an instance with IPs from all carriers is created by default.","zh_CN":"指定部分公网ipv6运营商, 仅多线节点生效。多线节点指定需要ipv6未指定该参数时，默认创建包含所有运营商IP的实例"}
  PublicIpv6Info []*VMPCreateInstanceRequestServersPublicIpv6Info `json:"publicIpv6Info,omitempty" xml:"publicIpv6Info,omitempty" type:"Repeated"`
}

func (s VMPCreateInstanceRequestServers) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestServers) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequestServers) SetRegionName(v string) *VMPCreateInstanceRequestServers {
  s.RegionName = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetProvince(v string) *VMPCreateInstanceRequestServers {
  s.Province = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetCarrier(v string) *VMPCreateInstanceRequestServers {
  s.Carrier = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetNodeName(v string) *VMPCreateInstanceRequestServers {
  s.NodeName = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetImageId(v string) *VMPCreateInstanceRequestServers {
  s.ImageId = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetFlavorId(v string) *VMPCreateInstanceRequestServers {
  s.FlavorId = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetName(v string) *VMPCreateInstanceRequestServers {
  s.Name = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetUserData(v string) *VMPCreateInstanceRequestServers {
  s.UserData = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetCount(v int) *VMPCreateInstanceRequestServers {
  s.Count = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPassword(v string) *VMPCreateInstanceRequestServers {
  s.Password = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetKeyName(v string) *VMPCreateInstanceRequestServers {
  s.KeyName = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetInnerNet(v string) *VMPCreateInstanceRequestServers {
  s.InnerNet = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetCidr(v string) *VMPCreateInstanceRequestServers {
  s.Cidr = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPrivateIPv4(v string) *VMPCreateInstanceRequestServers {
  s.PrivateIPv4 = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetInnerNet2(v string) *VMPCreateInstanceRequestServers {
  s.InnerNet2 = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetCidr2(v string) *VMPCreateInstanceRequestServers {
  s.Cidr2 = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPrivateIPv42(v string) *VMPCreateInstanceRequestServers {
  s.PrivateIPv42 = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPrivateIpv6Info(v []*VMPCreateInstanceRequestServersPrivateIpv6Info) *VMPCreateInstanceRequestServers {
  s.PrivateIpv6Info = v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetProtocols(v int) *VMPCreateInstanceRequestServers {
  s.Protocols = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetIpv4NativeAttribute(v string) *VMPCreateInstanceRequestServers {
  s.Ipv4NativeAttribute = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetIpv6NativeAttribute(v string) *VMPCreateInstanceRequestServers {
  s.Ipv6NativeAttribute = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetIsFree(v string) *VMPCreateInstanceRequestServers {
  s.IsFree = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetSecurityGroupIds(v []*string) *VMPCreateInstanceRequestServers {
  s.SecurityGroupIds = v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetDiskInfo(v []*VMPCreateInstanceRequestServersDiskInfo) *VMPCreateInstanceRequestServers {
  s.DiskInfo = v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetServerGroup(v string) *VMPCreateInstanceRequestServers {
  s.ServerGroup = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetTag(v string) *VMPCreateInstanceRequestServers {
  s.Tag = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetUseUniqueIpSegment(v int) *VMPCreateInstanceRequestServers {
  s.UseUniqueIpSegment = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetRandomAllocateIp(v int) *VMPCreateInstanceRequestServers {
  s.RandomAllocateIp = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetDefaultGateway(v string) *VMPCreateInstanceRequestServers {
  s.DefaultGateway = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPolicyRoutingType(v int) *VMPCreateInstanceRequestServers {
  s.PolicyRoutingType = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPrivateGatewayFlag(v int) *VMPCreateInstanceRequestServers {
  s.PrivateGatewayFlag = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetNicAllocateType(v int) *VMPCreateInstanceRequestServers {
  s.NicAllocateType = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetSinglePublicIpv4Cidr(v string) *VMPCreateInstanceRequestServers {
  s.SinglePublicIpv4Cidr = &v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPublicIpv4Info(v []*VMPCreateInstanceRequestServersPublicIpv4Info) *VMPCreateInstanceRequestServers {
  s.PublicIpv4Info = v
  return s
}

func (s *VMPCreateInstanceRequestServers) SetPublicIpv6Info(v []*VMPCreateInstanceRequestServersPublicIpv6Info) *VMPCreateInstanceRequestServers {
  s.PublicIpv6Info = v
  return s
}

type VMPCreateInstanceRequestServersPrivateIpv6Info struct     {
  // {"en":"Inner network number","zh_CN":"内网编号（1-对应v4的内网1；2-对应v4的内网2）"}
  NetNo *int `json:"netNo,omitempty" xml:"netNo,omitempty"`
  // {"en":"Inner network ipv6 cidr","zh_CN":"指定内网IPv6 CIDR"}
  PrivateCidr *string `json:"privateCidr,omitempty" xml:"privateCidr,omitempty"`
  // {"en":"Inner network ipv6 address:","zh_CN":"指定内网IPv6地址"}
  PrivateIp *string `json:"privateIp,omitempty" xml:"privateIp,omitempty"`
}

func (s VMPCreateInstanceRequestServersPrivateIpv6Info) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestServersPrivateIpv6Info) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequestServersPrivateIpv6Info) SetNetNo(v int) *VMPCreateInstanceRequestServersPrivateIpv6Info {
  s.NetNo = &v
  return s
}

func (s *VMPCreateInstanceRequestServersPrivateIpv6Info) SetPrivateCidr(v string) *VMPCreateInstanceRequestServersPrivateIpv6Info {
  s.PrivateCidr = &v
  return s
}

func (s *VMPCreateInstanceRequestServersPrivateIpv6Info) SetPrivateIp(v string) *VMPCreateInstanceRequestServersPrivateIpv6Info {
  s.PrivateIp = &v
  return s
}

type VMPCreateInstanceRequestServersDiskInfo struct     {
  // {"en":"Disk size (GB)","zh_CN":"磁盘大小（GB）"}
  Size *int `json:"size,omitempty" xml:"size,omitempty"`
  // {"en":"Disk Purpose:\nSystem - System disk;\nDATA - DATA plate","zh_CN":"磁盘用途：\nSYSTEM-系统盘；\nDATA-数据盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Disk type: HDD/SSD","zh_CN":"磁盘类型：HDD/SSD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // {"en":"Is isolate: 1(Yes) / -1(No)","zh_CN":"是否独立盘：1(是) / -1(否)"}
  IsIndependent *string `json:"isIndependent,omitempty" xml:"isIndependent,omitempty"`
}

func (s VMPCreateInstanceRequestServersDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestServersDiskInfo) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequestServersDiskInfo) SetSize(v int) *VMPCreateInstanceRequestServersDiskInfo {
  s.Size = &v
  return s
}

func (s *VMPCreateInstanceRequestServersDiskInfo) SetType(v string) *VMPCreateInstanceRequestServersDiskInfo {
  s.Type = &v
  return s
}

func (s *VMPCreateInstanceRequestServersDiskInfo) SetCategory(v string) *VMPCreateInstanceRequestServersDiskInfo {
  s.Category = &v
  return s
}

func (s *VMPCreateInstanceRequestServersDiskInfo) SetIsIndependent(v string) *VMPCreateInstanceRequestServersDiskInfo {
  s.IsIndependent = &v
  return s
}

type VMPCreateInstanceRequestServersPublicIpv4Info struct     {
  // {"en":"ISP abbreviation format, eg: dx,wt,yd","zh_CN":"运营商缩写格式：dx-电信；yd-移动；wt-网通"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
}

func (s VMPCreateInstanceRequestServersPublicIpv4Info) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestServersPublicIpv4Info) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequestServersPublicIpv4Info) SetCarrier(v string) *VMPCreateInstanceRequestServersPublicIpv4Info {
  s.Carrier = &v
  return s
}

type VMPCreateInstanceRequestServersPublicIpv6Info struct     {
  // {"en":"ISP abbreviation format, eg: dx,wt,yd","zh_CN":"运营商缩写格式：dx-电信；yd-移动；wt-网通"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
}

func (s VMPCreateInstanceRequestServersPublicIpv6Info) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestServersPublicIpv6Info) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequestServersPublicIpv6Info) SetCarrier(v string) *VMPCreateInstanceRequestServersPublicIpv6Info {
  s.Carrier = &v
  return s
}

type VMPCreateInstanceRequestHeader struct {
}

func (s VMPCreateInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestHeader) GoString() string {
  return s.String()
}

type VMPCreateInstancePaths struct {
}

func (s VMPCreateInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstancePaths) GoString() string {
  return s.String()
}

type VMPCreateInstanceParameters struct {
}

func (s VMPCreateInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceParameters) GoString() string {
  return s.String()
}

type VMPCreateInstanceResponse struct {
  // {"en":"Virtual machine identity list","zh_CN":"实例id"}
  Id []*string `json:"id,omitempty" xml:"id,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceResponse) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceResponse) SetId(v []*string) *VMPCreateInstanceResponse {
  s.Id = v
  return s
}

type VMPCreateInstanceResponseHeader struct {
}

func (s VMPCreateInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceResponseHeader) GoString() string {
  return s.String()
}




type CreateDeploymentRequest struct {
  // {"en":"ecci instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"specification of the desired behavior of the Ecci.", "zh_CN":"Ecci 预期行为的规约"}
  Spec *CreateDeploymentEcciInstanceSpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
  // {"en":"instance distribution", "zh_CN":"实例分发属性"}
  Distributions []*CreateDeploymentInstanceDistribution `json:"distributions,omitempty" xml:"distributions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether to assign public IP", "zh_CN":"是否分配公网ip"}
  AllocatePublicIp *bool `json:"allocatePublicIp,omitempty" xml:"allocatePublicIp,omitempty" require:"true"`
}

func (s CreateDeploymentRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentRequest) GoString() string {
  return s.String()
}

func (s *CreateDeploymentRequest) SetName(v string) *CreateDeploymentRequest {
  s.Name = &v
  return s
}

func (s *CreateDeploymentRequest) SetSpec(v *CreateDeploymentEcciInstanceSpec) *CreateDeploymentRequest {
  s.Spec = v
  return s
}

func (s *CreateDeploymentRequest) SetDistributions(v []*CreateDeploymentInstanceDistribution) *CreateDeploymentRequest {
  s.Distributions = v
  return s
}

func (s *CreateDeploymentRequest) SetAllocatePublicIp(v bool) *CreateDeploymentRequest {
  s.AllocatePublicIp = &v
  return s
}

type CreateDeploymentResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ecci object", "zh_CN":"ecci对象"}
  Data *CreateDeploymentSummary `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateDeploymentResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentResponse) GoString() string {
  return s.String()
}

func (s *CreateDeploymentResponse) SetCode(v int64) *CreateDeploymentResponse {
  s.Code = &v
  return s
}

func (s *CreateDeploymentResponse) SetMsg(v string) *CreateDeploymentResponse {
  s.Msg = &v
  return s
}

func (s *CreateDeploymentResponse) SetRequestId(v string) *CreateDeploymentResponse {
  s.RequestId = &v
  return s
}

func (s *CreateDeploymentResponse) SetData(v *CreateDeploymentSummary) *CreateDeploymentResponse {
  s.Data = v
  return s
}

type CreateDeploymentPaths struct {
}

func (s CreateDeploymentPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPaths) GoString() string {
  return s.String()
}

type CreateDeploymentParameters struct {
}

func (s CreateDeploymentParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentParameters) GoString() string {
  return s.String()
}

type CreateDeploymentRequestHeader struct {
}

func (s CreateDeploymentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentRequestHeader) GoString() string {
  return s.String()
}

type CreateDeploymentResponseHeader struct {
}

func (s CreateDeploymentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentResponseHeader) GoString() string {
  return s.String()
}

type CreateDeploymentContainer struct {
  // {"en":"CreateDeploymentContainer cpu limit, the sum of all container cpu limits cannot exceed the instance cpu limit", "zh_CN":"容器cpu限制,所有容器cpu限制总和不能超过实例cpu限制"}
  CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty" require:"true"`
  // {"en":"CreateDeploymentContainer memory limit, the sum of all container memory limits cannot exceed the instance memory limit", "zh_CN":"容器内存限制,所有容器内存限制总和不能超过实例内存限制"}
  MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty" require:"true"`
  // {"en":"CreateDeploymentContainer name", "zh_CN":"容器名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"CreateDeploymentContainer image", "zh_CN":"容器镜像"}
  Image *string `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en":"Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided.", "zh_CN":"入口点数组。不在 Shell 中执行。如果未提供，则使用容器镜像的 ENTRYPOINT"}
  Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
  // {"en":"Arguments to the entrypoint. The container image's CMD is used if this is not provided", "zh_CN":"entrypoint 的参数。如果未提供，则使用容器镜像的 CMD 设置"}
  Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
  // {"en":"CreateDeploymentContainer's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image", "zh_CN":"容器的工作目录。如果未指定，将使用容器运行时的默认值，默认值可能在容器镜像中配置"}
  WorkingDir *string `json:"workingDir,omitempty" xml:"workingDir,omitempty"`
  // {"en":"ist of environment variables to set in the container", "zh_CN":"要在容器中设置的环境变量列表"}
  Env map[string]*string `json:"env,omitempty" xml:"env,omitempty"`
  // {"en":"downwardapi type sensitive environment variables, authorization required", "zh_CN":"downwardapi 类型敏感环境变量,需授权"}
  EnvRef []*CreateDeploymentContainerEnvRef `json:"envRef,omitempty" xml:"envRef,omitempty" type:"Repeated"`
  // {"en":"Pod volumes to mount into the container's filesystem. Cannot be updated", "zh_CN":"要挂载到容器文件系统中的 Pod 卷。无法更新"}
  VolumeMounts []*CreateDeploymentVolumeMount `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty" type:"Repeated"`
}

func (s CreateDeploymentContainer) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentContainer) GoString() string {
  return s.String()
}

func (s *CreateDeploymentContainer) SetCpuLimit(v string) *CreateDeploymentContainer {
  s.CpuLimit = &v
  return s
}

func (s *CreateDeploymentContainer) SetMemoryLimit(v string) *CreateDeploymentContainer {
  s.MemoryLimit = &v
  return s
}

func (s *CreateDeploymentContainer) SetName(v string) *CreateDeploymentContainer {
  s.Name = &v
  return s
}

func (s *CreateDeploymentContainer) SetImage(v string) *CreateDeploymentContainer {
  s.Image = &v
  return s
}

func (s *CreateDeploymentContainer) SetCommand(v []*string) *CreateDeploymentContainer {
  s.Command = v
  return s
}

func (s *CreateDeploymentContainer) SetArgs(v []*string) *CreateDeploymentContainer {
  s.Args = v
  return s
}

func (s *CreateDeploymentContainer) SetWorkingDir(v string) *CreateDeploymentContainer {
  s.WorkingDir = &v
  return s
}

func (s *CreateDeploymentContainer) SetEnv(v map[string]*string) *CreateDeploymentContainer {
  s.Env = v
  return s
}

func (s *CreateDeploymentContainer) SetEnvRef(v []*CreateDeploymentContainerEnvRef) *CreateDeploymentContainer {
  s.EnvRef = v
  return s
}

func (s *CreateDeploymentContainer) SetVolumeMounts(v []*CreateDeploymentVolumeMount) *CreateDeploymentContainer {
  s.VolumeMounts = v
  return s
}

type CreateDeploymentPodSpec struct {
  // {"en":"List of containers belonging to the pod. There must be at least one container in a Pod. ", "zh_CN":"属于 Pod 的容器列表。Pod 中必须至少有一个容器。"}
  Containers []*CreateDeploymentContainer `json:"containers,omitempty" xml:"containers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always.", "zh_CN":"Pod 内所有容器的重启策略。Always、OnFailure、Never 之一。默认为 Always。"}
  RestartPolicy *string `json:"restartPolicy,omitempty" xml:"restartPolicy,omitempty" require:"true"`
  // {"en":"Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.", "zh_CN":"可选字段，表示 Pod 需要体面终止的所需的时长（以秒为单位）。字段值可以在删除请求中减少。 字段值必须是非负整数。零值表示收到 kill 信号则立即停止（没有机会关闭）。 如果此值为 nil，则将使用默认宽限期。 宽限期是从 Pod 中运行的进程收到终止信号后，到进程被 kill 信号强制停止之前，Pod 可以继续存在的时间（以秒为单位）。 应该将此值设置为比你的进程的预期清理时间更长。默认为 30 秒。"}
  TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" xml:"terminationGracePeriodSeconds,omitempty" require:"true"`
  // {"en":"List of volumes that can be mounted by containers belonging to the pod.", "zh_CN":"可以由属于 Pod 的容器挂载的卷列表。"}
  Volumes []*CreateDeploymentPodVolume `json:"volumes,omitempty" xml:"volumes,omitempty" type:"Repeated"`
  // {"en":"If specified, the pod's scheduling constraints", "zh_CN":"如果指定了，则作为 Pod 的调度约束。"}
  CreateDeploymentAffinity *CreateDeploymentAffinity `json:"affinity,omitempty" xml:"affinity,omitempty" require:"true"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Whether to use kata runtime, Use kata by default", "zh_CN":"是否使用kata运行时,默认使用kata"}
  KataRuntime *bool `json:"kataRuntime,omitempty" xml:"kataRuntime,omitempty"`
  // {"en":"is colocation resource or not", "zh_CN":"是否在离线混部资源"}
  Colocation *bool `json:"colocation,omitempty" xml:"colocation,omitempty"`
  // {"en":"the cluster behavior", "zh_CN":"集群策略配置"}
  CreateDeploymentClusterBehavior *CreateDeploymentClusterBehavior `json:"clusterBehavior,omitempty" xml:"clusterBehavior,omitempty"`
}

func (s CreateDeploymentPodSpec) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPodSpec) GoString() string {
  return s.String()
}

func (s *CreateDeploymentPodSpec) SetContainers(v []*CreateDeploymentContainer) *CreateDeploymentPodSpec {
  s.Containers = v
  return s
}

func (s *CreateDeploymentPodSpec) SetRestartPolicy(v string) *CreateDeploymentPodSpec {
  s.RestartPolicy = &v
  return s
}

func (s *CreateDeploymentPodSpec) SetTerminationGracePeriodSeconds(v int64) *CreateDeploymentPodSpec {
  s.TerminationGracePeriodSeconds = &v
  return s
}

func (s *CreateDeploymentPodSpec) SetVolumes(v []*CreateDeploymentPodVolume) *CreateDeploymentPodSpec {
  s.Volumes = v
  return s
}

func (s *CreateDeploymentPodSpec) SetAffinity(v *CreateDeploymentAffinity) *CreateDeploymentPodSpec {
  s.CreateDeploymentAffinity = v
  return s
}

func (s *CreateDeploymentPodSpec) SetAnnotations(v map[string]*string) *CreateDeploymentPodSpec {
  s.Annotations = v
  return s
}

func (s *CreateDeploymentPodSpec) SetLabels(v map[string]*string) *CreateDeploymentPodSpec {
  s.Labels = v
  return s
}

func (s *CreateDeploymentPodSpec) SetKataRuntime(v bool) *CreateDeploymentPodSpec {
  s.KataRuntime = &v
  return s
}

func (s *CreateDeploymentPodSpec) SetColocation(v bool) *CreateDeploymentPodSpec {
  s.Colocation = &v
  return s
}

func (s *CreateDeploymentPodSpec) SetClusterBehavior(v *CreateDeploymentClusterBehavior) *CreateDeploymentPodSpec {
  s.CreateDeploymentClusterBehavior = v
  return s
}

type CreateDeploymentClusterBehavior struct {
  // {"en":"a list of the scheduling constraints", "zh_CN":"调度策略列表"}
  CreateDeploymentSpreadConstraint []*CreateDeploymentSpreadConstraint `json:"spreadConstraint,omitempty" xml:"spreadConstraint,omitempty" type:"Repeated"`
  // {"en":"cluster faileover", "zh_CN":"集群重调度"}
  CreateDeploymentClusterFailover *CreateDeploymentClusterFailover `json:"clusterFailover,omitempty" xml:"clusterFailover,omitempty"`
}

func (s CreateDeploymentClusterBehavior) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentClusterBehavior) GoString() string {
  return s.String()
}

func (s *CreateDeploymentClusterBehavior) SetSpreadConstraint(v []*CreateDeploymentSpreadConstraint) *CreateDeploymentClusterBehavior {
  s.CreateDeploymentSpreadConstraint = v
  return s
}

func (s *CreateDeploymentClusterBehavior) SetClusterFailover(v *CreateDeploymentClusterFailover) *CreateDeploymentClusterBehavior {
  s.CreateDeploymentClusterFailover = v
  return s
}

type CreateDeploymentSpreadConstraint struct {
  // {"en":"spread field support: group or cluster", "zh_CN":"分组策略,当前支持: group, cluster"}
  SpreadByField *string `json:"spreadByField,omitempty" xml:"spreadByField,omitempty"`
  // {"en":"max groups in field. by default 1", "zh_CN":"策略需要的最大分组个数.默认: 1"}
  MaxGroups *int32 `json:"maxGroups,omitempty" xml:"maxGroups,omitempty"`
  // {"en":"min groups in field. by default 1", "zh_CN":"策略需要的最小分组个数.默认: 1"}
  MinGroups *int32 `json:"minGroups,omitempty" xml:"minGroups,omitempty"`
}

func (s CreateDeploymentSpreadConstraint) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentSpreadConstraint) GoString() string {
  return s.String()
}

func (s *CreateDeploymentSpreadConstraint) SetSpreadByField(v string) *CreateDeploymentSpreadConstraint {
  s.SpreadByField = &v
  return s
}

func (s *CreateDeploymentSpreadConstraint) SetMaxGroups(v int32) *CreateDeploymentSpreadConstraint {
  s.MaxGroups = &v
  return s
}

func (s *CreateDeploymentSpreadConstraint) SetMinGroups(v int32) *CreateDeploymentSpreadConstraint {
  s.MinGroups = &v
  return s
}

type CreateDeploymentClusterFailover struct {
  // {"en":"the cluser faileover. 0: off, 1: on. default 0", "zh_CN":"集群故障转移开关.0: 关闭,1打开. 默认:0 "}
  CreateDeploymentClusterFailover *int32 `json:"clusterFailover,omitempty" xml:"clusterFailover,omitempty" require:"true"`
  // {"en":"the tolerationSeconds of the workload unhealthy. default 300", "zh_CN":"应用创建后多久时间没有runing,超过认为不健康.默认300秒"}
  TolerationSeconds *int32 `json:"tolerationSeconds,omitempty" xml:"tolerationSeconds,omitempty"`
  // {"en":"PurgeMode represents how to deal with the legacy applications on the cluster from which the application is migrated.only support Graciously", "zh_CN":"应用漂移后,旧应用的删除模式.只支持 Graciously 平滑删除"}
  PurgeMode *string `json:"purgeMode,omitempty" xml:"purgeMode,omitempty"`
  // {"en":"GracePeriodSeconds is the maximum waiting duration in seconds before application on the migrated cluster should be deleted. default:300", "zh_CN":"平滑删除时间. 默认:300秒 "}
  GracePeriodSeconds *int32 `json:"gracePeriodSeconds,omitempty" xml:"gracePeriodSeconds,omitempty"`
}

func (s CreateDeploymentClusterFailover) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentClusterFailover) GoString() string {
  return s.String()
}

func (s *CreateDeploymentClusterFailover) SetClusterFailover(v int32) *CreateDeploymentClusterFailover {
  s.CreateDeploymentClusterFailover = &v
  return s
}

func (s *CreateDeploymentClusterFailover) SetTolerationSeconds(v int32) *CreateDeploymentClusterFailover {
  s.TolerationSeconds = &v
  return s
}

func (s *CreateDeploymentClusterFailover) SetPurgeMode(v string) *CreateDeploymentClusterFailover {
  s.PurgeMode = &v
  return s
}

func (s *CreateDeploymentClusterFailover) SetGracePeriodSeconds(v int32) *CreateDeploymentClusterFailover {
  s.GracePeriodSeconds = &v
  return s
}

type CreateDeploymentEcciInstanceSpec struct {
  // {"en":"Instance cpu limit", "zh_CN":"实例cpu限制"}
  CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty" require:"true"`
  // {"en":"Instance memory limit", "zh_CN":"实例内存限制"}
  MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty" require:"true"`
  // {"en":"Instance description", "zh_CN":"实例描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Specification of the desired behavior of the pod.", "zh_CN":"Pod 预期行为的规约。"}
  CreateDeploymentPodSpec *CreateDeploymentPodSpec `json:"podSpec,omitempty" xml:"podSpec,omitempty" require:"true"`
  // {"en":"Failover, whether the pod is rescheduled after a node exception, the default is false", "zh_CN":"故障转移,node异常后pod是否重新调度,默认false"}
  Failover *bool `json:"failover,omitempty" xml:"failover,omitempty" require:"true"`
}

func (s CreateDeploymentEcciInstanceSpec) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentEcciInstanceSpec) GoString() string {
  return s.String()
}

func (s *CreateDeploymentEcciInstanceSpec) SetCpuLimit(v string) *CreateDeploymentEcciInstanceSpec {
  s.CpuLimit = &v
  return s
}

func (s *CreateDeploymentEcciInstanceSpec) SetMemoryLimit(v string) *CreateDeploymentEcciInstanceSpec {
  s.MemoryLimit = &v
  return s
}

func (s *CreateDeploymentEcciInstanceSpec) SetDescription(v string) *CreateDeploymentEcciInstanceSpec {
  s.Description = &v
  return s
}

func (s *CreateDeploymentEcciInstanceSpec) SetPodSpec(v *CreateDeploymentPodSpec) *CreateDeploymentEcciInstanceSpec {
  s.CreateDeploymentPodSpec = v
  return s
}

func (s *CreateDeploymentEcciInstanceSpec) SetFailover(v bool) *CreateDeploymentEcciInstanceSpec {
  s.Failover = &v
  return s
}

type CreateDeploymentInstanceDistribution struct {
  // {"en":"distribution cluster name", "zh_CN":"部署集群名称"}
  Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
  // {"en":"the cluster selector", "zh_CN":"集群标签选择. 这字段和集群名称二选一"}
  ClusterSelector map[string]*string `json:"clusterSelector,omitempty" xml:"clusterSelector,omitempty"`
  // {"en":"Number of desired pods", "zh_CN":"预期 Pod 的数量"}
  Replicas *int64 `json:"replicas,omitempty" xml:"replicas,omitempty" require:"true"`
  // {"en":"assigned ip attributes", "zh_CN":"分配ip属性"}
  AssignIps []*CreateDeploymentAssignIp `json:"assignIps,omitempty" xml:"assignIps,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentInstanceDistribution) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentInstanceDistribution) GoString() string {
  return s.String()
}

func (s *CreateDeploymentInstanceDistribution) SetCluster(v string) *CreateDeploymentInstanceDistribution {
  s.Cluster = &v
  return s
}

func (s *CreateDeploymentInstanceDistribution) SetClusterSelector(v map[string]*string) *CreateDeploymentInstanceDistribution {
  s.ClusterSelector = v
  return s
}

func (s *CreateDeploymentInstanceDistribution) SetReplicas(v int64) *CreateDeploymentInstanceDistribution {
  s.Replicas = &v
  return s
}

func (s *CreateDeploymentInstanceDistribution) SetAssignIps(v []*CreateDeploymentAssignIp) *CreateDeploymentInstanceDistribution {
  s.AssignIps = v
  return s
}

type CreateDeploymentAssignIp struct {
  // {"en":"ip operator id", "zh_CN":"ip所属运营商id"}
  IspId *int64 `json:"ispId,omitempty" xml:"ispId,omitempty" require:"true"`
  // {"en":"Whether to assign ipv4 ip, ipv4 is assigned by default", "zh_CN":"是否分配ipv4的ip，默认分配ipv4"}
  Ipv4 *bool `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true"`
  // {"en":"Whether to assign ipv6 ip", "zh_CN":"是否分配ipv6的ip"}
  Ipv6 *bool `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true"`
}

func (s CreateDeploymentAssignIp) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentAssignIp) GoString() string {
  return s.String()
}

func (s *CreateDeploymentAssignIp) SetIspId(v int64) *CreateDeploymentAssignIp {
  s.IspId = &v
  return s
}

func (s *CreateDeploymentAssignIp) SetIpv4(v bool) *CreateDeploymentAssignIp {
  s.Ipv4 = &v
  return s
}

func (s *CreateDeploymentAssignIp) SetIpv6(v bool) *CreateDeploymentAssignIp {
  s.Ipv6 = &v
  return s
}

type CreateDeploymentDeploySummary struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty" require:"true"`
  // {"en":"Number of desired pods", "zh_CN":"预期 Pod 的数量"}
  Replicas *int64 `json:"replicas,omitempty" xml:"replicas,omitempty" require:"true"`
  // {"en":"Number of created pods", "zh_CN":"已创建 Pod 的数量"}
  CreatedReplicas *int64 `json:"createdReplicas,omitempty" xml:"createdReplicas,omitempty" require:"true"`
  // {"en":"The number of IP addresses assigned to the instance", "zh_CN":"实例分配ip数量"}
  IpAllocatedReplicas *int64 `json:"ipAllocatedReplicas,omitempty" xml:"ipAllocatedReplicas,omitempty" require:"true"`
  // {"en":"Instance name list", "zh_CN":"实例名称列表"}
  InstanceNames []*CreateDeploymentInstanceName `json:"instanceNames,omitempty" xml:"instanceNames,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentDeploySummary) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentDeploySummary) GoString() string {
  return s.String()
}

func (s *CreateDeploymentDeploySummary) SetClusterName(v string) *CreateDeploymentDeploySummary {
  s.ClusterName = &v
  return s
}

func (s *CreateDeploymentDeploySummary) SetReplicas(v int64) *CreateDeploymentDeploySummary {
  s.Replicas = &v
  return s
}

func (s *CreateDeploymentDeploySummary) SetCreatedReplicas(v int64) *CreateDeploymentDeploySummary {
  s.CreatedReplicas = &v
  return s
}

func (s *CreateDeploymentDeploySummary) SetIpAllocatedReplicas(v int64) *CreateDeploymentDeploySummary {
  s.IpAllocatedReplicas = &v
  return s
}

func (s *CreateDeploymentDeploySummary) SetInstanceNames(v []*CreateDeploymentInstanceName) *CreateDeploymentDeploySummary {
  s.InstanceNames = v
  return s
}

type CreateDeploymentInstanceName struct {
  // {"en":"Instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance display name", "zh_CN":"实例展示名称"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
}

func (s CreateDeploymentInstanceName) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentInstanceName) GoString() string {
  return s.String()
}

func (s *CreateDeploymentInstanceName) SetName(v string) *CreateDeploymentInstanceName {
  s.Name = &v
  return s
}

func (s *CreateDeploymentInstanceName) SetDisplayName(v string) *CreateDeploymentInstanceName {
  s.DisplayName = &v
  return s
}

type CreateDeploymentSummary struct {
  // {"en":"Instance properties", "zh_CN":"实例属性"}
  DeploySummaries []*CreateDeploymentDeploySummary `json:"deploySummaries,omitempty" xml:"deploySummaries,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentSummary) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentSummary) GoString() string {
  return s.String()
}

func (s *CreateDeploymentSummary) SetDeploySummaries(v []*CreateDeploymentDeploySummary) *CreateDeploymentSummary {
  s.DeploySummaries = v
  return s
}

type CreateDeploymentVolumeMount struct {
  // {"en":"This must match the Name of a Volume", "zh_CN":"此字段必须与卷的名称匹配"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Path within the container at which the volume should be mounted. Must not contain ':'", "zh_CN":"容器内卷的挂载路径。不得包含 ':'"}
  MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty" require:"true"`
  // {"en":"Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false", "zh_CN":"如果为 true，则以只读方式挂载，否则（false 或未设置）以读写方式挂载。默认为 false"}
  ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty" require:"true"`
}

func (s CreateDeploymentVolumeMount) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentVolumeMount) GoString() string {
  return s.String()
}

func (s *CreateDeploymentVolumeMount) SetName(v string) *CreateDeploymentVolumeMount {
  s.Name = &v
  return s
}

func (s *CreateDeploymentVolumeMount) SetMountPath(v string) *CreateDeploymentVolumeMount {
  s.MountPath = &v
  return s
}

func (s *CreateDeploymentVolumeMount) SetReadOnly(v bool) *CreateDeploymentVolumeMount {
  s.ReadOnly = &v
  return s
}

type CreateDeploymentContainerEnvRef struct {
  // {"en":"Environment variable name", "zh_CN":"环境变量名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Path to the field to be selected by downwardapi", "zh_CN":"downwardapi要选择的字段的路径"}
  FieldPath *string `json:"fieldPath,omitempty" xml:"fieldPath,omitempty" require:"true"`
}

func (s CreateDeploymentContainerEnvRef) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentContainerEnvRef) GoString() string {
  return s.String()
}

func (s *CreateDeploymentContainerEnvRef) SetName(v string) *CreateDeploymentContainerEnvRef {
  s.Name = &v
  return s
}

func (s *CreateDeploymentContainerEnvRef) SetFieldPath(v string) *CreateDeploymentContainerEnvRef {
  s.FieldPath = &v
  return s
}

type CreateDeploymentHostPathVolume struct {
  // {"en":"path of the directory on the host. If the path is a symlink, it will follow the link to the real path", "zh_CN":"目录在主机上的路径。如果该路径是一个符号链接，则它将沿着链接指向真实路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"type for HostPath Volume Defaults to ''", "zh_CN":"卷的类型。默认为 ''"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s CreateDeploymentHostPathVolume) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentHostPathVolume) GoString() string {
  return s.String()
}

func (s *CreateDeploymentHostPathVolume) SetPath(v string) *CreateDeploymentHostPathVolume {
  s.Path = &v
  return s
}

func (s *CreateDeploymentHostPathVolume) SetType(v string) *CreateDeploymentHostPathVolume {
  s.Type = &v
  return s
}

type CreateDeploymentPodVolume struct {
  // {"en":"name of the volume. Must be a DNS_LABEL and unique within the pod.", "zh_CN":"卷的名称。必须是 DNS_LABEL 且在 Pod 内是唯一的。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"represents a reference to a PersistentVolumeClaim in the same namespace", "zh_CN":"表示对同一名字空间中 PersistentVolumeClaim 的引用"}
  PersistentVolumeClaim *CreateDeploymentPodVolumeClaim `json:"persistentVolumeClaim,omitempty" xml:"persistentVolumeClaim,omitempty" require:"true"`
  // {"en":"Represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this.", "zh_CN":"表示主机上预先存在的文件或目录，它们将被直接暴露给容器。 这种卷通常用于系统代理或允许查看主机的其他特权操作。大多数容器不需要这种卷。"}
  HostPath *CreateDeploymentHostPathVolume `json:"hostPath,omitempty" xml:"hostPath,omitempty" require:"true"`
}

func (s CreateDeploymentPodVolume) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPodVolume) GoString() string {
  return s.String()
}

func (s *CreateDeploymentPodVolume) SetName(v string) *CreateDeploymentPodVolume {
  s.Name = &v
  return s
}

func (s *CreateDeploymentPodVolume) SetPersistentVolumeClaim(v *CreateDeploymentPodVolumeClaim) *CreateDeploymentPodVolume {
  s.PersistentVolumeClaim = v
  return s
}

func (s *CreateDeploymentPodVolume) SetHostPath(v *CreateDeploymentHostPathVolume) *CreateDeploymentPodVolume {
  s.HostPath = v
  return s
}

type CreateDeploymentPodVolumeClaim struct {
  // {"en":"the name of a PersistentVolumeClaim in the same namespace as the pod using this volume", "zh_CN":"与使用此卷的 Pod 位于同一名字空间中的 PersistentVolumeClaim 的名称"}
  ClaimName *string `json:"claimName,omitempty" xml:"claimName,omitempty" require:"true"`
}

func (s CreateDeploymentPodVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPodVolumeClaim) GoString() string {
  return s.String()
}

func (s *CreateDeploymentPodVolumeClaim) SetClaimName(v string) *CreateDeploymentPodVolumeClaim {
  s.ClaimName = &v
  return s
}

type CreateDeploymentAffinity struct {
  // {"en":"A group of inter pod affinity scheduling rules.", "zh_CN":"一组 Pod 间亲和性调度规则。"}
  CreateDeploymentPodAffinity *CreateDeploymentPodAffinity `json:"podAffinity,omitempty" xml:"podAffinity,omitempty" require:"true"`
  // {"en":"A group of node affinity scheduling rules.", "zh_CN":"一组节点亲和性调度规则。"}
  CreateDeploymentPodAntiAffinity *CreateDeploymentPodAntiAffinity `json:"podAntiAffinity,omitempty" xml:"podAntiAffinity,omitempty" require:"true"`
}

func (s CreateDeploymentAffinity) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentAffinity) GoString() string {
  return s.String()
}

func (s *CreateDeploymentAffinity) SetPodAffinity(v *CreateDeploymentPodAffinity) *CreateDeploymentAffinity {
  s.CreateDeploymentPodAffinity = v
  return s
}

func (s *CreateDeploymentAffinity) SetPodAntiAffinity(v *CreateDeploymentPodAntiAffinity) *CreateDeploymentAffinity {
  s.CreateDeploymentPodAntiAffinity = v
  return s
}

type CreateDeploymentPodAntiAffinity struct {
  // {"en":"If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.", "zh_CN":"如果在调度时不满足该字段指定的反亲和性要求，则该 Pod 不会被调度到该节点上。 如果在 Pod 执行期间的某个时间点不再满足此字段指定的反亲和性要求（例如：由于 Pod 标签更新）， 系统可能会或可能不会尝试最终将 Pod 从其节点中逐出。 当有多个元素时，每个 podAffinityTerm 对应的节点列表是取其交集的，即必须满足所有条件。"}
  RequiredDuringSchedulingIgnoredDuringExecution []*CreateDeploymentPodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
  // {"en":"The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.", "zh_CN":"调度器更倾向于将 Pod 调度到满足该字段指定的反亲和性表达式的节点， 但它可能会选择违反一个或多个表达式的节点。 最优选的节点是权重总和最大的节点，即对于满足所有调度要求（资源请求、requiredDuringScheduling 反亲和性表达式等）的每个节点，通过遍历元素来计算总和如果节点具有与相应 podAffinityTerm 匹配的 Pod，则此字段并在总和中添加\"权重\"；具有最高加和的节点是最优选的。"}
  PreferredDuringSchedulingIgnoredDuringExecution []*CreateDeploymentWeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentPodAntiAffinity) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPodAntiAffinity) GoString() string {
  return s.String()
}

func (s *CreateDeploymentPodAntiAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []*CreateDeploymentPodAffinityTerm) *CreateDeploymentPodAntiAffinity {
  s.RequiredDuringSchedulingIgnoredDuringExecution = v
  return s
}

func (s *CreateDeploymentPodAntiAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []*CreateDeploymentWeightedPodAffinityTerm) *CreateDeploymentPodAntiAffinity {
  s.PreferredDuringSchedulingIgnoredDuringExecution = v
  return s
}

type CreateDeploymentPodAffinity struct {
  // {"en":"If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.", "zh_CN":"如果在调度时不满足该字段指定的亲和性要求，则该 Pod 不会被调度到该节点上。 如果在 Pod 执行期间的某个时间点不再满足此字段指定的亲和性要求（例如：由于 Pod 标签更新）， 系统可能会也可能不会尝试最终将 Pod 从其节点中逐出。 当此列表中有多个元素时，每个 podAffinityTerm 对应的节点列表是取其交集的，即必须满足所有条件。"}
  RequiredDuringSchedulingIgnoredDuringExecution []*CreateDeploymentPodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
  // {"en":"The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.", "zh_CN":"调度器会更倾向于将 Pod 调度到满足该字段指定的亲和性表达式的节点， 但它可能会选择违反一个或多个表达式的节点。最优选择是权重总和最大的节点， 即对于满足所有调度要求（资源请求、requiredDuringScheduling 亲和表达式等）的每个节点， 通过迭代该字段的元素来计算总和，如果节点具有与相应 podAffinityTerm 匹配的 Pod，则将“权重”添加到总和中； 具有最高总和的节点是最优选的。"}
  PreferredDuringSchedulingIgnoredDuringExecution []*CreateDeploymentWeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" xml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentPodAffinity) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPodAffinity) GoString() string {
  return s.String()
}

func (s *CreateDeploymentPodAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []*CreateDeploymentPodAffinityTerm) *CreateDeploymentPodAffinity {
  s.RequiredDuringSchedulingIgnoredDuringExecution = v
  return s
}

func (s *CreateDeploymentPodAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []*CreateDeploymentWeightedPodAffinityTerm) *CreateDeploymentPodAffinity {
  s.PreferredDuringSchedulingIgnoredDuringExecution = v
  return s
}

type CreateDeploymentPodAffinityTerm struct {
  // {"en":"A label query over a set of resources, in this case pods.", "zh_CN":"对一组资源的标签查询，在这里资源为 Pod。"}
  CreateDeploymentLabelSelector *CreateDeploymentLabelSelector `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" require:"true"`
  // {"en":"pod's namespace", "zh_CN":" Pod 的名字空间"}
  Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" require:"true" type:"Repeated"`
  // {"en":"This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.", "zh_CN":"此 Pod 应与指定名字空间中与标签选择算符匹配的 Pod 集合位于同一位置（亲和性） 或位于不同位置（反亲和性），这里的“在同一位置”意味着运行在一个节点上，其键名为 topologyKey 的标签值与运行所选 Pod 集合中的某 Pod 的任何节点上的标签值匹配。 不允许使用空的 topologyKey。"}
  TopologyKey *string `json:"topologyKey,omitempty" xml:"topologyKey,omitempty" require:"true"`
}

func (s CreateDeploymentPodAffinityTerm) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentPodAffinityTerm) GoString() string {
  return s.String()
}

func (s *CreateDeploymentPodAffinityTerm) SetLabelSelector(v *CreateDeploymentLabelSelector) *CreateDeploymentPodAffinityTerm {
  s.CreateDeploymentLabelSelector = v
  return s
}

func (s *CreateDeploymentPodAffinityTerm) SetNamespaces(v []*string) *CreateDeploymentPodAffinityTerm {
  s.Namespaces = v
  return s
}

func (s *CreateDeploymentPodAffinityTerm) SetTopologyKey(v string) *CreateDeploymentPodAffinityTerm {
  s.TopologyKey = &v
  return s
}

type CreateDeploymentLabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty" require:"true"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*CreateDeploymentLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentLabelSelector) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentLabelSelector) GoString() string {
  return s.String()
}

func (s *CreateDeploymentLabelSelector) SetMatchLabels(v map[string]*string) *CreateDeploymentLabelSelector {
  s.MatchLabels = v
  return s
}

func (s *CreateDeploymentLabelSelector) SetMatchExpressions(v []*CreateDeploymentLabelSelectorRequirement) *CreateDeploymentLabelSelector {
  s.MatchExpressions = v
  return s
}

type CreateDeploymentLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" require:"true" type:"Repeated"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
}

func (s CreateDeploymentLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *CreateDeploymentLabelSelectorRequirement) SetKey(v string) *CreateDeploymentLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *CreateDeploymentLabelSelectorRequirement) SetValues(v []*string) *CreateDeploymentLabelSelectorRequirement {
  s.Values = v
  return s
}

func (s *CreateDeploymentLabelSelectorRequirement) SetOperator(v string) *CreateDeploymentLabelSelectorRequirement {
  s.Operator = &v
  return s
}

type CreateDeploymentWeightedPodAffinityTerm struct {
  // {"en":"associated with matching the corresponding podAffinityTerm, in the range 1-100.", "zh_CN":"匹配相应 podAffinityTerm 条件的权重，范围为 1-100。"}
  Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"Required. A pod affinity term, associated with the corresponding weight.", "zh_CN":"必需的字段。一个 Pod 亲和性条件，对应一个与相应的权重值。"}
  CreateDeploymentPodAffinityTerm *CreateDeploymentPodAffinityTerm `json:"podAffinityTerm,omitempty" xml:"podAffinityTerm,omitempty" require:"true"`
}

func (s CreateDeploymentWeightedPodAffinityTerm) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentWeightedPodAffinityTerm) GoString() string {
  return s.String()
}

func (s *CreateDeploymentWeightedPodAffinityTerm) SetWeight(v int32) *CreateDeploymentWeightedPodAffinityTerm {
  s.Weight = &v
  return s
}

func (s *CreateDeploymentWeightedPodAffinityTerm) SetPodAffinityTerm(v *CreateDeploymentPodAffinityTerm) *CreateDeploymentWeightedPodAffinityTerm {
  s.CreateDeploymentPodAffinityTerm = v
  return s
}




type GetClusterListRequest struct {
}

func (s GetClusterListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListRequest) GoString() string {
  return s.String()
}

type GetClusterListResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"cluster list", "zh_CN":"集群列表"}
  Data *GetClusterListClusterList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetClusterListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListResponse) GoString() string {
  return s.String()
}

func (s *GetClusterListResponse) SetCode(v int64) *GetClusterListResponse {
  s.Code = &v
  return s
}

func (s *GetClusterListResponse) SetMsg(v string) *GetClusterListResponse {
  s.Msg = &v
  return s
}

func (s *GetClusterListResponse) SetRequestId(v string) *GetClusterListResponse {
  s.RequestId = &v
  return s
}

func (s *GetClusterListResponse) SetData(v *GetClusterListClusterList) *GetClusterListResponse {
  s.Data = v
  return s
}

type GetClusterListClusterList struct {
  // {"en":"cluster list", "zh_CN":"集群列表"}
  Clusters []*GetClusterListClusterSummary `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"total clusters", "zh_CN":"集群总数量"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s GetClusterListClusterList) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListClusterList) GoString() string {
  return s.String()
}

func (s *GetClusterListClusterList) SetClusters(v []*GetClusterListClusterSummary) *GetClusterListClusterList {
  s.Clusters = v
  return s
}

func (s *GetClusterListClusterList) SetTotal(v int64) *GetClusterListClusterList {
  s.Total = &v
  return s
}

type GetClusterListClusterSummary struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"cluster chinese name", "zh_CN":"集群中文名称"}
  NameCn *string `json:"nameCn,omitempty" xml:"nameCn,omitempty" require:"true"`
  // {"en":"area", "zh_CN":"区域"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"country", "zh_CN":"国家"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"city", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"IP operator list", "zh_CN":"ip运营商列表"}
  Isps []*GetClusterListIsp `json:"isps,omitempty" xml:"isps,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether to support kata runtime", "zh_CN":"是否支持kata运行时"}
  KataRuntime *bool `json:"kataRuntime,omitempty" xml:"kataRuntime,omitempty" require:"true"`
}

func (s GetClusterListClusterSummary) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListClusterSummary) GoString() string {
  return s.String()
}

func (s *GetClusterListClusterSummary) SetName(v string) *GetClusterListClusterSummary {
  s.Name = &v
  return s
}

func (s *GetClusterListClusterSummary) SetNameCn(v string) *GetClusterListClusterSummary {
  s.NameCn = &v
  return s
}

func (s *GetClusterListClusterSummary) SetArea(v string) *GetClusterListClusterSummary {
  s.Area = &v
  return s
}

func (s *GetClusterListClusterSummary) SetCountry(v string) *GetClusterListClusterSummary {
  s.Country = &v
  return s
}

func (s *GetClusterListClusterSummary) SetProvince(v string) *GetClusterListClusterSummary {
  s.Province = &v
  return s
}

func (s *GetClusterListClusterSummary) SetCity(v string) *GetClusterListClusterSummary {
  s.City = &v
  return s
}

func (s *GetClusterListClusterSummary) SetIsps(v []*GetClusterListIsp) *GetClusterListClusterSummary {
  s.Isps = v
  return s
}

func (s *GetClusterListClusterSummary) SetKataRuntime(v bool) *GetClusterListClusterSummary {
  s.KataRuntime = &v
  return s
}

type GetClusterListIsp struct {
  // {"en":"operator id", "zh_CN":"运营商id"}
  IspId *int64 `json:"ispId,omitempty" xml:"ispId,omitempty" require:"true"`
  // {"en":"operator english name", "zh_CN":"运营商英文名"}
  IspEn *string `json:"ispEn,omitempty" xml:"ispEn,omitempty" require:"true"`
  // {"en":"operator chinese name", "zh_CN":"运营商中文名"}
  IspCn *string `json:"ispCn,omitempty" xml:"ispCn,omitempty" require:"true"`
}

func (s GetClusterListIsp) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListIsp) GoString() string {
  return s.String()
}

func (s *GetClusterListIsp) SetIspId(v int64) *GetClusterListIsp {
  s.IspId = &v
  return s
}

func (s *GetClusterListIsp) SetIspEn(v string) *GetClusterListIsp {
  s.IspEn = &v
  return s
}

func (s *GetClusterListIsp) SetIspCn(v string) *GetClusterListIsp {
  s.IspCn = &v
  return s
}

type GetClusterListPaths struct {
}

func (s GetClusterListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListPaths) GoString() string {
  return s.String()
}

type GetClusterListParameters struct {
  // {"en":"node ids", "zh_CN":"节点ids"}
  NodeIds []*string `json:"nodeIds,omitempty" xml:"nodeIds,omitempty" type:"Repeated"`
  // {"en":"operator ids", "zh_CN":"运营商ids"}
  IspIds []*string `json:"ispIds,omitempty" xml:"ispIds,omitempty" type:"Repeated"`
  // {"en":"page index", "zh_CN":"页数"}
  PageIndex *int64 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"page size", "zh_CN":"每页个数"}
  PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"0: does not support kataruntime, 1: supports kataruntime, others: all", "zh_CN":"0: 不支持kataruntime, 1: 支持kataruntime, 其他: 全部"}
  KataRuntime *string `json:"kataRuntime,omitempty" xml:"kataRuntime,omitempty"`
}

func (s GetClusterListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListParameters) GoString() string {
  return s.String()
}

func (s *GetClusterListParameters) SetNodeIds(v []*string) *GetClusterListParameters {
  s.NodeIds = v
  return s
}

func (s *GetClusterListParameters) SetIspIds(v []*string) *GetClusterListParameters {
  s.IspIds = v
  return s
}

func (s *GetClusterListParameters) SetPageIndex(v int64) *GetClusterListParameters {
  s.PageIndex = &v
  return s
}

func (s *GetClusterListParameters) SetPageSize(v int64) *GetClusterListParameters {
  s.PageSize = &v
  return s
}

func (s *GetClusterListParameters) SetKataRuntime(v string) *GetClusterListParameters {
  s.KataRuntime = &v
  return s
}

type GetClusterListRequestHeader struct {
}

func (s GetClusterListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListRequestHeader) GoString() string {
  return s.String()
}

type GetClusterListResponseHeader struct {
}

func (s GetClusterListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClusterListResponseHeader) GoString() string {
  return s.String()
}




type InstanceDiskScalingRequest struct {
  // {"en":"instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DiskInfo []*InstanceDiskScalingRequestDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceDiskScalingRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingRequest) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingRequest) SetId(v string) *InstanceDiskScalingRequest {
  s.Id = &v
  return s
}

func (s *InstanceDiskScalingRequest) SetDiskInfo(v []*InstanceDiskScalingRequestDiskInfo) *InstanceDiskScalingRequest {
  s.DiskInfo = v
  return s
}

type InstanceDiskScalingRequestDiskInfo struct     {
  // {"en":"disk size GB","zh_CN":"磁盘大小，单位GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"disk type ,HDD/SDD","zh_CN":"磁盘类型，取值：\nHDD：普通硬盘\nSSD：固态硬盘"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s InstanceDiskScalingRequestDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingRequestDiskInfo) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingRequestDiskInfo) SetSize(v int) *InstanceDiskScalingRequestDiskInfo {
  s.Size = &v
  return s
}

func (s *InstanceDiskScalingRequestDiskInfo) SetCategory(v string) *InstanceDiskScalingRequestDiskInfo {
  s.Category = &v
  return s
}

type InstanceDiskScalingRequestHeader struct {
}

func (s InstanceDiskScalingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingRequestHeader) GoString() string {
  return s.String()
}

type InstanceDiskScalingPaths struct {
}

func (s InstanceDiskScalingPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingPaths) GoString() string {
  return s.String()
}

type InstanceDiskScalingParameters struct {
}

func (s InstanceDiskScalingParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingParameters) GoString() string {
  return s.String()
}

type InstanceDiskScalingResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *InstanceDiskScalingResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s InstanceDiskScalingResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponse) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingResponse) SetCode(v string) *InstanceDiskScalingResponse {
  s.Code = &v
  return s
}

func (s *InstanceDiskScalingResponse) SetMessage(v string) *InstanceDiskScalingResponse {
  s.Message = &v
  return s
}

func (s *InstanceDiskScalingResponse) SetData(v *InstanceDiskScalingResponseData) *InstanceDiskScalingResponse {
  s.Data = v
  return s
}

type InstanceDiskScalingResponseData struct {
  // {"en":"Instance info","zh_CN":"实例信息"}
  Server *InstanceDiskScalingResponseDataServer `json:"server,omitempty" xml:"server,omitempty" require:"true" type:"Struct"`
}

func (s InstanceDiskScalingResponseData) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponseData) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingResponseData) SetServer(v *InstanceDiskScalingResponseDataServer) *InstanceDiskScalingResponseData {
  s.Server = v
  return s
}

type InstanceDiskScalingResponseDataServer struct {
  // {"en":"Instance ID","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Disk info","zh_CN":"磁盘信息"}
  DiskInfo []*InstanceDiskScalingResponseDataServerDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceDiskScalingResponseDataServer) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponseDataServer) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingResponseDataServer) SetId(v string) *InstanceDiskScalingResponseDataServer {
  s.Id = &v
  return s
}

func (s *InstanceDiskScalingResponseDataServer) SetDiskInfo(v []*InstanceDiskScalingResponseDataServerDiskInfo) *InstanceDiskScalingResponseDataServer {
  s.DiskInfo = v
  return s
}

type InstanceDiskScalingResponseDataServerDiskInfo struct     {
  // {"en":"Disk size","zh_CN":"磁盘大小，单位GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk type","zh_CN":"磁盘类型,DATA：数据盘, SYSTEM：系统盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk category","zh_CN":"磁盘类型，HDD/SDD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s InstanceDiskScalingResponseDataServerDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponseDataServerDiskInfo) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingResponseDataServerDiskInfo) SetSize(v int) *InstanceDiskScalingResponseDataServerDiskInfo {
  s.Size = &v
  return s
}

func (s *InstanceDiskScalingResponseDataServerDiskInfo) SetType(v string) *InstanceDiskScalingResponseDataServerDiskInfo {
  s.Type = &v
  return s
}

func (s *InstanceDiskScalingResponseDataServerDiskInfo) SetCategory(v string) *InstanceDiskScalingResponseDataServerDiskInfo {
  s.Category = &v
  return s
}

type InstanceDiskScalingResponseHeader struct {
}

func (s InstanceDiskScalingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponseHeader) GoString() string {
  return s.String()
}




type LECHConvertFreeTypeInstanceToChargeTypeRequest struct {
  // {"en":"Unique cloud host identity.Up to 100 IDs can be sent at a time, separated by the half comma character ', '.","zh_CN":"云主机唯一标识。单次最多可发送100 条ID，ID 之间用半角逗号字符','隔开。"}
  Servers *string `json:"servers,omitempty" xml:"servers,omitempty" require:"true"`
}

func (s LECHConvertFreeTypeInstanceToChargeTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeRequest) GoString() string {
  return s.String()
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeRequest) SetServers(v string) *LECHConvertFreeTypeInstanceToChargeTypeRequest {
  s.Servers = &v
  return s
}

type LECHConvertFreeTypeInstanceToChargeTypeRequestHeader struct {
}

func (s LECHConvertFreeTypeInstanceToChargeTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeRequestHeader) GoString() string {
  return s.String()
}

type LECHConvertFreeTypeInstanceToChargeTypePaths struct {
}

func (s LECHConvertFreeTypeInstanceToChargeTypePaths) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypePaths) GoString() string {
  return s.String()
}

type LECHConvertFreeTypeInstanceToChargeTypeParameters struct {
}

func (s LECHConvertFreeTypeInstanceToChargeTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeParameters) GoString() string {
  return s.String()
}

type LECHConvertFreeTypeInstanceToChargeTypeResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHConvertFreeTypeInstanceToChargeTypeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponse) GoString() string {
  return s.String()
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponse) SetCode(v string) *LECHConvertFreeTypeInstanceToChargeTypeResponse {
  s.Code = &v
  return s
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponse) SetData(v *LECHConvertFreeTypeInstanceToChargeTypeResponseData) *LECHConvertFreeTypeInstanceToChargeTypeResponse {
  s.Data = v
  return s
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponse) SetMessage(v string) *LECHConvertFreeTypeInstanceToChargeTypeResponse {
  s.Message = &v
  return s
}

type LECHConvertFreeTypeInstanceToChargeTypeResponseData struct {
  // {"en":"Batch error","zh_CN":"批量操作失败信息"}
  BatchErrorMsg []*LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponseData) GoString() string {
  return s.String()
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponseData) SetBatchErrorMsg(v []*LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) *LECHConvertFreeTypeInstanceToChargeTypeResponseData {
  s.BatchErrorMsg = v
  return s
}

type LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg struct     {
  // {"en":"Error code","zh_CN":"错误编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Instance ID","zh_CN":"实例id"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) SetCode(v string) *LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) SetKey(v string) *LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg) SetMsg(v string) *LECHConvertFreeTypeInstanceToChargeTypeResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type LECHConvertFreeTypeInstanceToChargeTypeResponseHeader struct {
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHConvertFreeTypeInstanceToChargeTypeResponseHeader) GoString() string {
  return s.String()
}




type LECHInstanceIpv6ManagementRequest struct {
  // {"en":"Operation:\nALLOCATION - ipv6 application\nREMOVE - ipv6 is removed","zh_CN":"操作：\nALLOCATION-ipv6申请\nREMOVE-ipv6移除"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Instance ID (single only)","zh_CN":"实例id（只支持单个）"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
}

func (s LECHInstanceIpv6ManagementRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementRequest) GoString() string {
  return s.String()
}

func (s *LECHInstanceIpv6ManagementRequest) SetAction(v string) *LECHInstanceIpv6ManagementRequest {
  s.Action = &v
  return s
}

func (s *LECHInstanceIpv6ManagementRequest) SetInstanceId(v string) *LECHInstanceIpv6ManagementRequest {
  s.InstanceId = &v
  return s
}

type LECHInstanceIpv6ManagementRequestHeader struct {
}

func (s LECHInstanceIpv6ManagementRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementRequestHeader) GoString() string {
  return s.String()
}

type LECHInstanceIpv6ManagementPaths struct {
}

func (s LECHInstanceIpv6ManagementPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementPaths) GoString() string {
  return s.String()
}

type LECHInstanceIpv6ManagementParameters struct {
}

func (s LECHInstanceIpv6ManagementParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementParameters) GoString() string {
  return s.String()
}

type LECHInstanceIpv6ManagementResponse struct {
  // {"en":"Response code","zh_CN":"响应码，0为成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *LECHInstanceIpv6ManagementResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHInstanceIpv6ManagementResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementResponse) GoString() string {
  return s.String()
}

func (s *LECHInstanceIpv6ManagementResponse) SetCode(v string) *LECHInstanceIpv6ManagementResponse {
  s.Code = &v
  return s
}

func (s *LECHInstanceIpv6ManagementResponse) SetMessage(v string) *LECHInstanceIpv6ManagementResponse {
  s.Message = &v
  return s
}

func (s *LECHInstanceIpv6ManagementResponse) SetData(v *LECHInstanceIpv6ManagementResponseData) *LECHInstanceIpv6ManagementResponse {
  s.Data = v
  return s
}

type LECHInstanceIpv6ManagementResponseData struct {
  // {"en":"Instance ID","zh_CN":"实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"IPv6","zh_CN":"IPv6"}
  AccessIPv6 []*string `json:"accessIPv6,omitempty" xml:"accessIPv6,omitempty" require:"true" type:"Repeated"`
}

func (s LECHInstanceIpv6ManagementResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementResponseData) GoString() string {
  return s.String()
}

func (s *LECHInstanceIpv6ManagementResponseData) SetInstanceId(v string) *LECHInstanceIpv6ManagementResponseData {
  s.InstanceId = &v
  return s
}

func (s *LECHInstanceIpv6ManagementResponseData) SetAccessIPv6(v []*string) *LECHInstanceIpv6ManagementResponseData {
  s.AccessIPv6 = v
  return s
}

type LECHInstanceIpv6ManagementResponseHeader struct {
}

func (s LECHInstanceIpv6ManagementResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHInstanceIpv6ManagementResponseHeader) GoString() string {
  return s.String()
}




type InstanceBandwidthAggregationQueryRequest struct {
}

func (s InstanceBandwidthAggregationQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryRequest) GoString() string {
  return s.String()
}

type InstanceBandwidthAggregationQueryRequestHeader struct {
}

func (s InstanceBandwidthAggregationQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryRequestHeader) GoString() string {
  return s.String()
}

type InstanceBandwidthAggregationQueryPaths struct {
}

func (s InstanceBandwidthAggregationQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryPaths) GoString() string {
  return s.String()
}

type InstanceBandwidthAggregationQueryParameters struct {
  // {"en":"Instance ID: A maximum of 20 IDs can be sent at a time, and IDs are separated by a comma character ','.","zh_CN":"实例ID，单次最多可发送20条ID，ID之间用半角逗号字符“,”隔开。已销毁的实例不支持查询。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Start time,format: YYYY-MM-DD. You can only query the flow data within the last 90 days, and the query range cannot exceed 31 days at a time.","zh_CN":"查询开始时间，格式yyyy-MM-dd。最多只能查询90天内的流量数据，且单次查询时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time: YYYY-MM-DD,and the query range cannot exceed 31 days at a time.","zh_CN":"查询结束时间，格式yyyy-MM-dd，单次查询时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s InstanceBandwidthAggregationQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryParameters) GoString() string {
  return s.String()
}

func (s *InstanceBandwidthAggregationQueryParameters) SetIds(v string) *InstanceBandwidthAggregationQueryParameters {
  s.Ids = &v
  return s
}

func (s *InstanceBandwidthAggregationQueryParameters) SetStartTime(v string) *InstanceBandwidthAggregationQueryParameters {
  s.StartTime = &v
  return s
}

func (s *InstanceBandwidthAggregationQueryParameters) SetEndTime(v string) *InstanceBandwidthAggregationQueryParameters {
  s.EndTime = &v
  return s
}

type InstanceBandwidthAggregationQueryResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *InstanceBandwidthAggregationQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s InstanceBandwidthAggregationQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryResponse) GoString() string {
  return s.String()
}

func (s *InstanceBandwidthAggregationQueryResponse) SetCode(v string) *InstanceBandwidthAggregationQueryResponse {
  s.Code = &v
  return s
}

func (s *InstanceBandwidthAggregationQueryResponse) SetData(v *InstanceBandwidthAggregationQueryResponseData) *InstanceBandwidthAggregationQueryResponse {
  s.Data = v
  return s
}

func (s *InstanceBandwidthAggregationQueryResponse) SetMessage(v string) *InstanceBandwidthAggregationQueryResponse {
  s.Message = &v
  return s
}

type InstanceBandwidthAggregationQueryResponseData struct {
  // {"en":"instances detail info","zh_CN":"虚拟机详细信息"}
  Servers []*InstanceBandwidthAggregationQueryResponseDataServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceBandwidthAggregationQueryResponseData) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryResponseData) GoString() string {
  return s.String()
}

func (s *InstanceBandwidthAggregationQueryResponseData) SetServers(v []*InstanceBandwidthAggregationQueryResponseDataServers) *InstanceBandwidthAggregationQueryResponseData {
  s.Servers = v
  return s
}

type InstanceBandwidthAggregationQueryResponseDataServers struct     {
  // {"en":"External Inbound Traffic Summary (MB)","zh_CN":"外网流入流量汇总值（MB）"}
  ExtTrafficIn *int64 `json:"extTrafficIn,omitempty" xml:"extTrafficIn,omitempty" require:"true"`
  // {"en":"External outbound Traffic Summary (MB)","zh_CN":"外网流出流量汇总值（MB）"}
  ExtTrafficOut *int64 `json:"extTrafficOut,omitempty" xml:"extTrafficOut,omitempty" require:"true"`
  // {"en":"instance Id","zh_CN":"实例Id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s InstanceBandwidthAggregationQueryResponseDataServers) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryResponseDataServers) GoString() string {
  return s.String()
}

func (s *InstanceBandwidthAggregationQueryResponseDataServers) SetExtTrafficIn(v int64) *InstanceBandwidthAggregationQueryResponseDataServers {
  s.ExtTrafficIn = &v
  return s
}

func (s *InstanceBandwidthAggregationQueryResponseDataServers) SetExtTrafficOut(v int64) *InstanceBandwidthAggregationQueryResponseDataServers {
  s.ExtTrafficOut = &v
  return s
}

func (s *InstanceBandwidthAggregationQueryResponseDataServers) SetId(v string) *InstanceBandwidthAggregationQueryResponseDataServers {
  s.Id = &v
  return s
}

type InstanceBandwidthAggregationQueryResponseHeader struct {
}

func (s InstanceBandwidthAggregationQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceBandwidthAggregationQueryResponseHeader) GoString() string {
  return s.String()
}




type QueryInstanceListRequest struct {
}

func (s QueryInstanceListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListRequest) GoString() string {
  return s.String()
}

type QueryInstanceListResponse struct {
  // {"en":"list of ephone instances", "zh_CN":"云手机实例列表"}
  Ephones []*QueryInstanceListEphone `json:"ephones,omitempty" xml:"ephones,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s QueryInstanceListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListResponse) GoString() string {
  return s.String()
}

func (s *QueryInstanceListResponse) SetEphones(v []*QueryInstanceListEphone) *QueryInstanceListResponse {
  s.Ephones = v
  return s
}

func (s *QueryInstanceListResponse) SetStatus(v int) *QueryInstanceListResponse {
  s.Status = &v
  return s
}

func (s *QueryInstanceListResponse) SetResult(v string) *QueryInstanceListResponse {
  s.Result = &v
  return s
}

type QueryInstanceListEphone struct {
  // {"en":"id of instance", "zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"name of instance", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"image info", "zh_CN":"镜像信息"}
  QueryInstanceListImage *QueryInstanceListImage `json:"image,omitempty" xml:"image,omitempty" require:"true"`
  // {"en":"resource of instance", "zh_CN":"实例资源详情"}
  QueryInstanceListResource *QueryInstanceListResource `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
  // {"en":"resolution", "zh_CN":"手机分辨率"}
  QueryInstanceListResolution *QueryInstanceListResolution `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {"en":"instance adb address", "zh_CN":"实例adb连接地址"}
  AdbAddress *string `json:"adbAddress,omitempty" xml:"adbAddress,omitempty" require:"true"`
  // {"en":"instance h5 address", "zh_CN":"实例h5连接地址"}
  H5Address *string `json:"h5Address,omitempty" xml:"h5Address,omitempty" require:"true"`
  // {"en":"state of instance", "zh_CN":"实例状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s QueryInstanceListEphone) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListEphone) GoString() string {
  return s.String()
}

func (s *QueryInstanceListEphone) SetId(v string) *QueryInstanceListEphone {
  s.Id = &v
  return s
}

func (s *QueryInstanceListEphone) SetName(v string) *QueryInstanceListEphone {
  s.Name = &v
  return s
}

func (s *QueryInstanceListEphone) SetImage(v *QueryInstanceListImage) *QueryInstanceListEphone {
  s.QueryInstanceListImage = v
  return s
}

func (s *QueryInstanceListEphone) SetResource(v *QueryInstanceListResource) *QueryInstanceListEphone {
  s.QueryInstanceListResource = v
  return s
}

func (s *QueryInstanceListEphone) SetResolution(v *QueryInstanceListResolution) *QueryInstanceListEphone {
  s.QueryInstanceListResolution = v
  return s
}

func (s *QueryInstanceListEphone) SetAdbAddress(v string) *QueryInstanceListEphone {
  s.AdbAddress = &v
  return s
}

func (s *QueryInstanceListEphone) SetH5Address(v string) *QueryInstanceListEphone {
  s.H5Address = &v
  return s
}

func (s *QueryInstanceListEphone) SetState(v string) *QueryInstanceListEphone {
  s.State = &v
  return s
}

type QueryInstanceListImage struct {
  // {"en":"os version", "zh_CN":"操作系统版本"}
  OsVersion *int `json:"osVersion,omitempty" xml:"osVersion,omitempty" require:"true"`
}

func (s QueryInstanceListImage) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListImage) GoString() string {
  return s.String()
}

func (s *QueryInstanceListImage) SetOsVersion(v int) *QueryInstanceListImage {
  s.OsVersion = &v
  return s
}

type QueryInstanceListResource struct {
  // {"en":"cpu", "zh_CN":"cpu核数"}
  Cpu *int `json:"cpu,omitempty" xml:"cpu,omitempty" require:"true"`
  // {"en":"memory", "zh_CN":"内存"}
  Memory *int `json:"memory,omitempty" xml:"memory,omitempty" require:"true"`
  // {"en":"disk storage", "zh_CN":"存储容量"}
  Storage *int `json:"storage,omitempty" xml:"storage,omitempty" require:"true"`
}

func (s QueryInstanceListResource) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListResource) GoString() string {
  return s.String()
}

func (s *QueryInstanceListResource) SetCpu(v int) *QueryInstanceListResource {
  s.Cpu = &v
  return s
}

func (s *QueryInstanceListResource) SetMemory(v int) *QueryInstanceListResource {
  s.Memory = &v
  return s
}

func (s *QueryInstanceListResource) SetStorage(v int) *QueryInstanceListResource {
  s.Storage = &v
  return s
}

type QueryInstanceListResolution struct {
  // {"en":"screen width", "zh_CN":"屏幕宽"}
  Width *int `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"screen height", "zh_CN":"屏幕高"}
  Height *int `json:"height,omitempty" xml:"height,omitempty" require:"true"`
}

func (s QueryInstanceListResolution) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListResolution) GoString() string {
  return s.String()
}

func (s *QueryInstanceListResolution) SetWidth(v int) *QueryInstanceListResolution {
  s.Width = &v
  return s
}

func (s *QueryInstanceListResolution) SetHeight(v int) *QueryInstanceListResolution {
  s.Height = &v
  return s
}

type QueryInstanceListPaths struct {
}

func (s QueryInstanceListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListPaths) GoString() string {
  return s.String()
}

type QueryInstanceListParameters struct {
  // {"en":"instance ID to be queried", "zh_CN":"要查询的实例id"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"instance name bo be queried", "zh_CN":"要查询的实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryInstanceListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListParameters) GoString() string {
  return s.String()
}

func (s *QueryInstanceListParameters) SetIds(v string) *QueryInstanceListParameters {
  s.Ids = &v
  return s
}

func (s *QueryInstanceListParameters) SetName(v string) *QueryInstanceListParameters {
  s.Name = &v
  return s
}

type QueryInstanceListRequestHeader struct {
}

func (s QueryInstanceListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListRequestHeader) GoString() string {
  return s.String()
}

type QueryInstanceListResponseHeader struct {
}

func (s QueryInstanceListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceListResponseHeader) GoString() string {
  return s.String()
}




type CreateEphoneInstanceRequest struct {
  // {"en":"list of ephones", "zh_CN":"云手机请求列表"}
  Ephones []*CreateEphoneInstanceEphone `json:"ephones,omitempty" xml:"ephones,omitempty" require:"true" type:"Repeated"`
}

func (s CreateEphoneInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceRequest) GoString() string {
  return s.String()
}

func (s *CreateEphoneInstanceRequest) SetEphones(v []*CreateEphoneInstanceEphone) *CreateEphoneInstanceRequest {
  s.Ephones = v
  return s
}

type CreateEphoneInstanceEphone struct {
  // {"en":"ephone spec", "zh_CN":"云手机详细配置"}
  CreateEphoneInstanceSpec *CreateEphoneInstanceSpec `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
  // {"en":"ephone resolution", "zh_CN":"云手机分辨率"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
  // {"en":"ephone alias name", "zh_CN":"云手机名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"nodeName", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"osVersion", "zh_CN":"操作系统版本，如Android10，Android11"}
  OsVersion *string `json:"osVersion,omitempty" xml:"osVersion,omitempty"`
  // {"en":"oem image id", "zh_CN":"oem镜像id"}
  OemImage *string `json:"oemImage,omitempty" xml:"oemImage,omitempty"`
  // {"en":"ephone counts", "zh_CN":"申请实例数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty"`
}

func (s CreateEphoneInstanceEphone) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceEphone) GoString() string {
  return s.String()
}

func (s *CreateEphoneInstanceEphone) SetSpec(v *CreateEphoneInstanceSpec) *CreateEphoneInstanceEphone {
  s.CreateEphoneInstanceSpec = v
  return s
}

func (s *CreateEphoneInstanceEphone) SetResolution(v string) *CreateEphoneInstanceEphone {
  s.Resolution = &v
  return s
}

func (s *CreateEphoneInstanceEphone) SetName(v string) *CreateEphoneInstanceEphone {
  s.Name = &v
  return s
}

func (s *CreateEphoneInstanceEphone) SetNodeName(v string) *CreateEphoneInstanceEphone {
  s.NodeName = &v
  return s
}

func (s *CreateEphoneInstanceEphone) SetOsVersion(v string) *CreateEphoneInstanceEphone {
  s.OsVersion = &v
  return s
}

func (s *CreateEphoneInstanceEphone) SetOemImage(v string) *CreateEphoneInstanceEphone {
  s.OemImage = &v
  return s
}

func (s *CreateEphoneInstanceEphone) SetCount(v int) *CreateEphoneInstanceEphone {
  s.Count = &v
  return s
}

type CreateEphoneInstanceSpec struct {
  // {"en":"ephone spec type", "zh_CN":"边缘云手机配置类型，可选值：basic-基础型，standard-通用型，professional-旗舰型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"cpu resource", "zh_CN":"边缘云手机CPU核数"}
  Cpu *int `json:"cpu,omitempty" xml:"cpu,omitempty" require:"true"`
  // {"en":"memory resource", "zh_CN":"边缘云手机内存大小"}
  Memory *int `json:"memory,omitempty" xml:"memory,omitempty" require:"true"`
}

func (s CreateEphoneInstanceSpec) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceSpec) GoString() string {
  return s.String()
}

func (s *CreateEphoneInstanceSpec) SetType(v string) *CreateEphoneInstanceSpec {
  s.Type = &v
  return s
}

func (s *CreateEphoneInstanceSpec) SetCpu(v int) *CreateEphoneInstanceSpec {
  s.Cpu = &v
  return s
}

func (s *CreateEphoneInstanceSpec) SetMemory(v int) *CreateEphoneInstanceSpec {
  s.Memory = &v
  return s
}

type CreateEphoneInstanceResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Tasks []*CreateEphoneInstanceTask `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s CreateEphoneInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceResponse) GoString() string {
  return s.String()
}

func (s *CreateEphoneInstanceResponse) SetTasks(v []*CreateEphoneInstanceTask) *CreateEphoneInstanceResponse {
  s.Tasks = v
  return s
}

func (s *CreateEphoneInstanceResponse) SetStatus(v int) *CreateEphoneInstanceResponse {
  s.Status = &v
  return s
}

func (s *CreateEphoneInstanceResponse) SetResult(v string) *CreateEphoneInstanceResponse {
  s.Result = &v
  return s
}

type CreateEphoneInstanceTask struct {
  // {"en":"task id", "zh_CN":"任务单id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"任务单消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateEphoneInstanceTask) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceTask) GoString() string {
  return s.String()
}

func (s *CreateEphoneInstanceTask) SetId(v string) *CreateEphoneInstanceTask {
  s.Id = &v
  return s
}

func (s *CreateEphoneInstanceTask) SetMessage(v string) *CreateEphoneInstanceTask {
  s.Message = &v
  return s
}

type CreateEphoneInstancePaths struct {
}

func (s CreateEphoneInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstancePaths) GoString() string {
  return s.String()
}

type CreateEphoneInstanceParameters struct {
}

func (s CreateEphoneInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceParameters) GoString() string {
  return s.String()
}

type CreateEphoneInstanceRequestHeader struct {
}

func (s CreateEphoneInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceRequestHeader) GoString() string {
  return s.String()
}

type CreateEphoneInstanceResponseHeader struct {
}

func (s CreateEphoneInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEphoneInstanceResponseHeader) GoString() string {
  return s.String()
}




type EditInstanceRequest struct {
  // {"en":"server","zh_CN":"实例信息对象"}
  Server *EditInstanceRequestServer `json:"server,omitempty" xml:"server,omitempty" require:"true" type:"Struct"`
}

func (s EditInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceRequest) GoString() string {
  return s.String()
}

func (s *EditInstanceRequest) SetServer(v *EditInstanceRequestServer) *EditInstanceRequest {
  s.Server = v
  return s
}

type EditInstanceRequestServer struct {
  // {"en":"Instance ID","zh_CN":"实例ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance Name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s EditInstanceRequestServer) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceRequestServer) GoString() string {
  return s.String()
}

func (s *EditInstanceRequestServer) SetId(v string) *EditInstanceRequestServer {
  s.Id = &v
  return s
}

func (s *EditInstanceRequestServer) SetName(v string) *EditInstanceRequestServer {
  s.Name = &v
  return s
}

type EditInstanceRequestHeader struct {
}

func (s EditInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceRequestHeader) GoString() string {
  return s.String()
}

type EditInstancePaths struct {
}

func (s EditInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s EditInstancePaths) GoString() string {
  return s.String()
}

type EditInstanceParameters struct {
}

func (s EditInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceParameters) GoString() string {
  return s.String()
}

type EditInstanceResponse struct {
}

func (s EditInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceResponse) GoString() string {
  return s.String()
}

type EditInstanceResponseHeader struct {
}

func (s EditInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceResponseHeader) GoString() string {
  return s.String()
}




