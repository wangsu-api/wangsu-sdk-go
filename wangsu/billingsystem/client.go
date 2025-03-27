package billingsystem

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryUsedNodesRequest struct {
}

func (s QueryUsedNodesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryUsedNodesRequest) GoString() string {
  return s.String()
}

type QueryUsedNodesResponse struct {
  // {"en":"the nodes that have been used", "zh_CN":"使用过的节点列表"}
  Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
}

func (s QueryUsedNodesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryUsedNodesResponse) GoString() string {
  return s.String()
}

func (s *QueryUsedNodesResponse) SetNodes(v []*string) *QueryUsedNodesResponse {
  s.Nodes = v
  return s
}

type QueryUsedNodesPaths struct {
}

func (s QueryUsedNodesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryUsedNodesPaths) GoString() string {
  return s.String()
}

type QueryUsedNodesParameters struct {
  // {"en":"startTime", "zh_CN":"起始时间，格式为：yyyy-MM-dd"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"endTime", "zh_CN":"截止时间，格式为：yyyy-MM-dd"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s QueryUsedNodesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryUsedNodesParameters) GoString() string {
  return s.String()
}

func (s *QueryUsedNodesParameters) SetStartTime(v string) *QueryUsedNodesParameters {
  s.StartTime = &v
  return s
}

func (s *QueryUsedNodesParameters) SetEndTime(v string) *QueryUsedNodesParameters {
  s.EndTime = &v
  return s
}

type QueryUsedNodesRequestHeader struct {
}

func (s QueryUsedNodesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryUsedNodesRequestHeader) GoString() string {
  return s.String()
}

type QueryUsedNodesResponseHeader struct {
}

func (s QueryUsedNodesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryUsedNodesResponseHeader) GoString() string {
  return s.String()
}




type QueryBillingDetailsOfComputingServiceRequest struct {
}

func (s QueryBillingDetailsOfComputingServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceRequest) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServiceServerBill struct {
  // {"en":"Instance id", "zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance IPv4", "zh_CN":"实例IPv4"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance IPv6", "zh_CN":"实例IPv6"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance type:1---Bare metal instance;-1:Virtual machine instance", "zh_CN":"实例类型
  // 1：裸机实例,-1：虚拟机实例"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"Instance Status:
  // RUNNING
  // STOPPED
  // ERROR
  // DELETED
  // RESTARTING
  // STARTING
  // STOPPING
  // SNAPSHOTTING
  // REBUILDING
  // MIGRATING
  // ", "zh_CN":"实例状态
  // RUNNING	运行状态
  // STOPPED	停机
  // ERROR	错误
  // DELETED	已销毁
  // RESTARTING  重启中
  // STARTING  启动中
  // STOPPING  停止中
  // SNAPSHOTTING  制作快照镜像中
  // REBUILDING 重建中
  // MIGRATING 迁移中"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Node name", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Node Chinese Name", "zh_CN":"节点中文名称"}
  NodeNameCn *string `json:"nodeNameCn,omitempty" xml:"nodeNameCn,omitempty" require:"true"`
  // {"en":"Charge region", "zh_CN":"计费区域，14个取值：
  // 中国大陆
  // 亚太
  // 香港
  // 台湾
  // 美洲
  // 欧洲
  // 中东
  // 非洲
  // 台湾 
  // 香港
  // 非洲
  // 南美
  // 澳大利亚
  // 印度"}
  ChargeRegion *string `json:"chargeRegion,omitempty" xml:"chargeRegion,omitempty" require:"true"`
  // {"en":"Instance spec", "zh_CN":"实例规格"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty" require:"true"`
  // {"en":"Instance configuration, including CPU, memory, and disk specifications information of the virtual machine", "zh_CN":"实例配置，包含虚拟机的CPU、内存、磁盘规格信息"}
  ServerFeature *string `json:"serverFeature,omitempty" xml:"serverFeature,omitempty" require:"true"`
  // {"en":"Starting time of use within the accounting period", "zh_CN":"账期内开始使用时间，格式yyyy-MM-dd"}
  PeriodStartTime *string `json:"periodStartTime,omitempty" xml:"periodStartTime,omitempty" require:"true"`
  // {"en":"End of usage time within the accounting period", "zh_CN":"账期内结束时间，格式yyyy-MM-dd"}
  PeriodEndTime *string `json:"periodEndTime,omitempty" xml:"periodEndTime,omitempty" require:"true"`
  // {"en":"Usage days within the accounting period", "zh_CN":"账期内使用天数"}
  PeriodDays *int32 `json:"periodDays,omitempty" xml:"periodDays,omitempty" require:"true"`
  // {"en":"Instance usage during the accounting period", "zh_CN":"账期内实例用量"}
  PeriodCount *float32 `json:"periodCount,omitempty" xml:"periodCount,omitempty" require:"true"`
}

func (s QueryBillingDetailsOfComputingServiceServerBill) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceServerBill) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetId(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Id = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetName(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Name = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetIpv4(v []*string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Ipv4 = v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetIpv6(v []*string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Ipv6 = v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetIsBm(v int) *QueryBillingDetailsOfComputingServiceServerBill {
  s.IsBm = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetState(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.State = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetNodeName(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.NodeName = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetNodeNameCn(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.NodeNameCn = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetChargeRegion(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.ChargeRegion = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetInstanceType(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.InstanceType = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetServerFeature(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.ServerFeature = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodStartTime(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodStartTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodEndTime(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodEndTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodDays(v int32) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodDays = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodCount(v float32) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodCount = &v
  return s
}

type QueryBillingDetailsOfComputingServiceResponse struct {
  // {"en":"Instance computing power information array", "zh_CN":"实例算力信息数组"}
  Servers []*QueryBillingDetailsOfComputingServiceServerBill `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBillingDetailsOfComputingServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceResponse) SetServers(v []*QueryBillingDetailsOfComputingServiceServerBill) *QueryBillingDetailsOfComputingServiceResponse {
  s.Servers = v
  return s
}

type QueryBillingDetailsOfComputingServicePaths struct {
}

func (s QueryBillingDetailsOfComputingServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServicePaths) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServiceParameters struct {
  // {"en":"Starting time of use within the accounting period,format:yyyy-MM-dd", "zh_CN":"账期开始时间，格式yyyy-MM-dd"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"The end time of the accounting period, in the format yyyy-MM-dd. If not filled in, it will be the current time.", "zh_CN":"账期结束时间，格式yyyy-MM-dd，如果未填写则为当前时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s QueryBillingDetailsOfComputingServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceParameters) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceParameters) SetStartTime(v string) *QueryBillingDetailsOfComputingServiceParameters {
  s.StartTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceParameters) SetEndTime(v string) *QueryBillingDetailsOfComputingServiceParameters {
  s.EndTime = &v
  return s
}

type QueryBillingDetailsOfComputingServiceRequestHeader struct {
}

func (s QueryBillingDetailsOfComputingServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServiceResponseHeader struct {
}

func (s QueryBillingDetailsOfComputingServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponseHeader) GoString() string {
  return s.String()
}




type VmNodeInfoForMaxOutAndInRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到天,日期格式为yyyy-mm-dd 此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到天,日期格式为yyyy-mm-dd 此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried: 1)If there are multiple inputs,use ';' as separator. 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道。"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号“;”分隔开，如查询大陆及亚太区域，参数填写为：“region=cn;apac”。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type. 1)If there are multiple inputs,use ';' as separator. 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"Return result format, supported formats are XML and JSON, default is XML.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"1)If there area multiple inputs,use ';' as demimeter. 2)optional values of isp: refers to the ISP-section of appendix. 3) If not specified,means all the isp.", "zh_CN":"要查询的运营商的缩写，多个isp请用英文分号“;”分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了“cn”时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  IspId *string `json:"ispId,omitempty" xml:"ispId,omitempty"`
  // {"en":"IPv4, IPv6, if not specified, search all by default.", "zh_CN":"ipv4、ipv6，不填默认查全部"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
  // {"en":"Node names, please separate multiple values with a semicolon ';'", "zh_CN":"节点名称,多个值请用英文分号;分隔"}
  Node *string `json:"node,omitempty" xml:"node,omitempty"`
  // {"en":"Instance specifications. Please separate multiple values with a semicolon ';'. If not selected or left blank, it defaults to querying all.", "zh_CN":"实例规格。多个值请用英文分号“;”，不选或者为空时默认为查询所有。"}
  Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"4 types. 0: Storage capacity; 1: Read request count; 2: Write request count; 3: Delete request count. Default is: 0.", "zh_CN":"4种类型。0：存储量；1：读请求数；2：写请求数；3：删请求数。默认取：0"}
  LineMode *string `json:"lineMode,omitempty" xml:"lineMode,omitempty"`
  // {"en":"Greenwich Mean Time (GMT) zone, the parameter format is GMT+09:00 to indicate East 9th Zone, and GMT-09:00 to indicate West 9th Zone. If not provided, the default is the local time zone (East 8th Zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s VmNodeInfoForMaxOutAndInRequest) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInRequest) GoString() string {
  return s.String()
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetCust(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Cust = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetDate(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Date = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetStartdate(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Startdate = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetEnddate(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Enddate = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetChannel(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Channel = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetRegion(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Region = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetAccetype(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Accetype = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetDataformat(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Dataformat = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetIspId(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.IspId = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetIpProtocol(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.IpProtocol = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetNode(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Node = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetSpec(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Spec = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetLineMode(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.LineMode = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInRequest) SetTimezone(v string) *VmNodeInfoForMaxOutAndInRequest {
  s.Timezone = &v
  return s
}

type VmNodeInfoForMaxOutAndInResponse struct {
  // {'en':'status code', 'zh_CN':'状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'统计数据'}
  Data []*VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s VmNodeInfoForMaxOutAndInResponse) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInResponse) GoString() string {
  return s.String()
}

func (s *VmNodeInfoForMaxOutAndInResponse) SetCode(v string) *VmNodeInfoForMaxOutAndInResponse {
  s.Code = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInResponse) SetMessage(v string) *VmNodeInfoForMaxOutAndInResponse {
  s.Message = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInResponse) SetData(v []*VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) *VmNodeInfoForMaxOutAndInResponse {
  s.Data = v
  return s
}

type VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData struct     {
  // {"en":"node", "zh_CN":"节点"}
  Node *string `json:"node,omitempty" xml:"node,omitempty" require:"true"`
  // {"en":"95th billing value, unit: Mbps", "zh_CN":"95计费值, 单位Mbps"}
  Charge95Value *string `json:"charge95Value,omitempty" xml:"charge95Value,omitempty" require:"true"`
  // {"en":"95th billing time", "zh_CN":"95计费值时刻"}
  Charge95Time *string `json:"charge95Time,omitempty" xml:"charge95Time,omitempty" require:"true"`
  // {"en":"Data types corresponding to the 95th percentile value: ein: external inbound bandwidth, eout: external outbound bandwidth", "zh_CN":"95值对应的数据类型：ein:外网流入带宽,eout:外网流出带宽"}
  DataTypeFor95 *string `json:"dataTypeFor95,omitempty" xml:"dataTypeFor95,omitempty" require:"true"`
  // {"en":"Peak average billing value, unit: Mbps", "zh_CN":"峰值平均计费值, 单位Mbps"}
  PeakAvgValue *string `json:"peakAvgValue,omitempty" xml:"peakAvgValue,omitempty" require:"true"`
  // {"en":"Peak value, unit: Mbps", "zh_CN":"第一峰值, 单位Mbps"}
  Peak1stValue *string `json:"peak1stValue,omitempty" xml:"peak1stValue,omitempty" require:"true"`
  // {"en":"Peak time", "zh_CN":"第一峰值时刻"}
  Peak1stTime *string `json:"peak1stTime,omitempty" xml:"peak1stTime,omitempty" require:"true"`
  // {"en":"Data types corresponding to the peak value: ein: external inbound bandwidth, eout: external outbound bandwidth", "zh_CN":"峰值对应的数据类型：ein:外网流入带宽,eout:外网流出带宽"}
  DataTypeForPeak *string `json:"dataTypeForPeak,omitempty" xml:"dataTypeForPeak,omitempty" require:"true"`
  // {"en":"Total traffic, unit: GB", "zh_CN":"总流量, 单位GB"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
}

func (s VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) GoString() string {
  return s.String()
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetNode(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.Node = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetCharge95Value(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.Charge95Value = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetCharge95Time(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.Charge95Time = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetDataTypeFor95(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.DataTypeFor95 = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetPeakAvgValue(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.PeakAvgValue = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetPeak1stValue(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.Peak1stValue = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetPeak1stTime(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.Peak1stTime = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetDataTypeForPeak(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.DataTypeForPeak = &v
  return s
}

func (s *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData) SetTotalFlow(v string) *VmNodeInfoForMaxOutAndInVmNodeInfoForMaxOutAndInResponseData {
  s.TotalFlow = &v
  return s
}

type VmNodeInfoForMaxOutAndInPaths struct {
}

func (s VmNodeInfoForMaxOutAndInPaths) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInPaths) GoString() string {
  return s.String()
}

type VmNodeInfoForMaxOutAndInParameters struct {
}

func (s VmNodeInfoForMaxOutAndInParameters) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInParameters) GoString() string {
  return s.String()
}

type VmNodeInfoForMaxOutAndInRequestHeader struct {
}

func (s VmNodeInfoForMaxOutAndInRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInRequestHeader) GoString() string {
  return s.String()
}

type VmNodeInfoForMaxOutAndInResponseHeader struct {
}

func (s VmNodeInfoForMaxOutAndInResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VmNodeInfoForMaxOutAndInResponseHeader) GoString() string {
  return s.String()
}




