package dataquery

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type AllSiteSumFlowRequest struct {
}

func (s AllSiteSumFlowRequest) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowRequest) GoString() string {
  return s.String()
}

type AllSiteSumFlowResponse struct {
  // {'en':'Interface status codes', 'zh_CN':'接口状态码'}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {'en':'Interface information', 'zh_CN':'接口信息'}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {'en':'response content', 'zh_CN':'响应内容'}
  Content *AllSiteSumFlowResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s AllSiteSumFlowResponse) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowResponse) GoString() string {
  return s.String()
}

func (s *AllSiteSumFlowResponse) SetReturnCode(v int) *AllSiteSumFlowResponse {
  s.ReturnCode = &v
  return s
}

func (s *AllSiteSumFlowResponse) SetReturnMsg(v string) *AllSiteSumFlowResponse {
  s.ReturnMsg = &v
  return s
}

func (s *AllSiteSumFlowResponse) SetContent(v *AllSiteSumFlowResponseContent) *AllSiteSumFlowResponse {
  s.Content = v
  return s
}

type AllSiteSumFlowResponseContent struct {
  // {'en':'Total outgoing traffic of the enterprise site', 'zh_CN':'企业站点流出流量总和'}
  SumSiteOutFlow *string `json:"sumSiteOutFlow,omitempty" xml:"sumSiteOutFlow,omitempty" require:"true"`
  // {'en':'The sum of the inbound traffic of the enterprise site', 'zh_CN':'企业站点流入流量总和'}
  SumSiteInFlow *string `json:"sumSiteInFlow,omitempty" xml:"sumSiteInFlow,omitempty" require:"true"`
}

func (s AllSiteSumFlowResponseContent) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowResponseContent) GoString() string {
  return s.String()
}

func (s *AllSiteSumFlowResponseContent) SetSumSiteOutFlow(v string) *AllSiteSumFlowResponseContent {
  s.SumSiteOutFlow = &v
  return s
}

func (s *AllSiteSumFlowResponseContent) SetSumSiteInFlow(v string) *AllSiteSumFlowResponseContent {
  s.SumSiteInFlow = &v
  return s
}

type AllSiteSumFlowPaths struct {
}

func (s AllSiteSumFlowPaths) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowPaths) GoString() string {
  return s.String()
}

type AllSiteSumFlowParameters struct {
  // {"en":"start time", "zh_CN":"1、查询的开始时间，格式：yyyy-mm-dd
  // 2、必须小于当前时间和endtime；
  // 3、startTime和endTime相差不能超过31天；（可联系技术支持调整）
  // 4、只能查询最近半年内数据。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"end time", "zh_CN":"1、查询的结束时间，格式：yyyy-mm-dd
  // 2、必须小于当前时间，大于或等于starttime
  // 3、startTime和endTime相差不能超过31天（可联系技术支持调整）；
  // 4、如果起始时间、结束时间是当天，且不满一天，那就查询从00:00到查询时刻的站点累计流量。
  // 4、只能查询最近半年内数据。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s AllSiteSumFlowParameters) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowParameters) GoString() string {
  return s.String()
}

func (s *AllSiteSumFlowParameters) SetStartTime(v string) *AllSiteSumFlowParameters {
  s.StartTime = &v
  return s
}

func (s *AllSiteSumFlowParameters) SetEndTime(v string) *AllSiteSumFlowParameters {
  s.EndTime = &v
  return s
}

type AllSiteSumFlowRequestHeader struct {
}

func (s AllSiteSumFlowRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowRequestHeader) GoString() string {
  return s.String()
}

type AllSiteSumFlowResponseHeader struct {
}

func (s AllSiteSumFlowResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AllSiteSumFlowResponseHeader) GoString() string {
  return s.String()
}




type SiteMsDeviceRequest struct {
}

func (s SiteMsDeviceRequest) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceRequest) GoString() string {
  return s.String()
}

type SiteMsDeviceResponse struct {
  // {'en':'Interface status codes', 'zh_CN':'接口状态码'}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {'en':'Interface information', 'zh_CN':'接口信息'}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {'en':'response content', 'zh_CN':'响应内容'}
  Content []*SiteMsDeviceResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s SiteMsDeviceResponse) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceResponse) GoString() string {
  return s.String()
}

func (s *SiteMsDeviceResponse) SetReturnCode(v int) *SiteMsDeviceResponse {
  s.ReturnCode = &v
  return s
}

func (s *SiteMsDeviceResponse) SetReturnMsg(v string) *SiteMsDeviceResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SiteMsDeviceResponse) SetContent(v []*SiteMsDeviceResponseContent) *SiteMsDeviceResponse {
  s.Content = v
  return s
}

type SiteMsDeviceResponseContent struct     {
  // {'en':'query time', 'zh_CN':'查询时间'}
  TimeStamp *string `json:"timeStamp,omitempty" xml:"timeStamp,omitempty" require:"true"`
  // {'en':'Use of site equipment', 'zh_CN':'站点设备的使用情况'}
  Details *SiteMsDeviceResponseContentDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Struct"`
}

func (s SiteMsDeviceResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceResponseContent) GoString() string {
  return s.String()
}

func (s *SiteMsDeviceResponseContent) SetTimeStamp(v string) *SiteMsDeviceResponseContent {
  s.TimeStamp = &v
  return s
}

func (s *SiteMsDeviceResponseContent) SetDetails(v *SiteMsDeviceResponseContentDetails) *SiteMsDeviceResponseContent {
  s.Details = v
  return s
}

type SiteMsDeviceResponseContentDetails struct {
  // {'en':'site name', 'zh_CN':'站点名称'}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {'en':'site Id', 'zh_CN':'站点 ID'}
  SiteId *int64 `json:"siteId,omitempty" xml:"siteId,omitempty" require:"true"`
  // {'en':'deviceNum num', 'zh_CN':'设备数量'}
  DeviceNum *int `json:"deviceNum,omitempty" xml:"deviceNum,omitempty" require:"true"`
  // {'en':'master site name', 'zh_CN':'主用设备名称'}
  Master *string `json:"master,omitempty" xml:"master,omitempty" require:"true"`
  // {'en':'slave site name', 'zh_CN':'备用设备名称'}
  Slave *string `json:"slave,omitempty" xml:"slave,omitempty" require:"true"`
}

func (s SiteMsDeviceResponseContentDetails) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceResponseContentDetails) GoString() string {
  return s.String()
}

func (s *SiteMsDeviceResponseContentDetails) SetSiteName(v string) *SiteMsDeviceResponseContentDetails {
  s.SiteName = &v
  return s
}

func (s *SiteMsDeviceResponseContentDetails) SetSiteId(v int64) *SiteMsDeviceResponseContentDetails {
  s.SiteId = &v
  return s
}

func (s *SiteMsDeviceResponseContentDetails) SetDeviceNum(v int) *SiteMsDeviceResponseContentDetails {
  s.DeviceNum = &v
  return s
}

func (s *SiteMsDeviceResponseContentDetails) SetMaster(v string) *SiteMsDeviceResponseContentDetails {
  s.Master = &v
  return s
}

func (s *SiteMsDeviceResponseContentDetails) SetSlave(v string) *SiteMsDeviceResponseContentDetails {
  s.Slave = &v
  return s
}

type SiteMsDevicePaths struct {
}

func (s SiteMsDevicePaths) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDevicePaths) GoString() string {
  return s.String()
}

type SiteMsDeviceParameters struct {
  // {'en':'Site name(Only support querying CPE sites)
  // This parameter can be unspecified; if not specified, it queries all CPE sites of the enterprise for their primary and backup equipment', 'zh_CN':'站点名称(只支持查询CPE站点)
  // 该参数可不指定，不指定时，查询企业的所有CPE站点的主用、备用设备情况'}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty"`
}

func (s SiteMsDeviceParameters) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceParameters) GoString() string {
  return s.String()
}

func (s *SiteMsDeviceParameters) SetSiteName(v string) *SiteMsDeviceParameters {
  s.SiteName = &v
  return s
}

type SiteMsDeviceRequestHeader struct {
}

func (s SiteMsDeviceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceRequestHeader) GoString() string {
  return s.String()
}

type SiteMsDeviceResponseHeader struct {
}

func (s SiteMsDeviceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteMsDeviceResponseHeader) GoString() string {
  return s.String()
}




type SiteSessionInfoRequest struct {
}

func (s SiteSessionInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoRequest) GoString() string {
  return s.String()
}

type SiteSessionInfoRequestHeader struct {
}

func (s SiteSessionInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoRequestHeader) GoString() string {
  return s.String()
}

type SiteSessionInfoPaths struct {
}

func (s SiteSessionInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoPaths) GoString() string {
  return s.String()
}

type SiteSessionInfoParameters struct {
  // {"en":"Site name\nWhen this parameter must be specified, query the specified site data","zh_CN":"站点名称\n该参数必须指定时，查询指定的站点数据"}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {"en":"1.the start time of the query, in the format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm\n2.must be less than the current time and endtime\n3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)\n4.can only query the last six months of data","zh_CN":"1、查询的开始时间，格式：yyyy-mm-ddThh:mm:ss 或 yyyy-mm-ddThh:mm\n2、必须小于当前时间和endtime\n3、startTime和endTime相差不能超过31天（可联系技术支持调整）\n4、只能查询最近半年内数据"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"1.the end time of the query, format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm\n2.must be less than the current time, more than starttime\n3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)\n4.can only query the last six months of data","zh_CN":"1、查询的结束时间，格式：yyyy-mm-ddThh:mm:ss或 yyyy-mm-ddThh:mm\n2、必须小于当前时间，大于starttime\n3、startTime和endTime相差不能超过31天（可联系技术支持调整）\n4、只能查询最近半年内数据"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s SiteSessionInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoParameters) GoString() string {
  return s.String()
}

func (s *SiteSessionInfoParameters) SetSiteName(v string) *SiteSessionInfoParameters {
  s.SiteName = &v
  return s
}

func (s *SiteSessionInfoParameters) SetStartTime(v string) *SiteSessionInfoParameters {
  s.StartTime = &v
  return s
}

func (s *SiteSessionInfoParameters) SetEndTime(v string) *SiteSessionInfoParameters {
  s.EndTime = &v
  return s
}

type SiteSessionInfoResponse struct {
  // {"en":"Interface status codes","zh_CN":"接口状态码"}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Interface information","zh_CN":"接口信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"response content","zh_CN":"响应内容"}
  Content *SiteSessionInfoResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s SiteSessionInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoResponse) GoString() string {
  return s.String()
}

func (s *SiteSessionInfoResponse) SetReturnCode(v int) *SiteSessionInfoResponse {
  s.ReturnCode = &v
  return s
}

func (s *SiteSessionInfoResponse) SetReturnMsg(v string) *SiteSessionInfoResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SiteSessionInfoResponse) SetContent(v *SiteSessionInfoResponseContent) *SiteSessionInfoResponse {
  s.Content = v
  return s
}

type SiteSessionInfoResponseContent struct {
  // {"en":"site name","zh_CN":"站点名称"}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {"en":"site Id","zh_CN":"站点 ID"}
  SiteId *int64 `json:"siteId,omitempty" xml:"siteId,omitempty" require:"true"`
  // {"en":"Session data","zh_CN":"会话数据"}
  Content []*SiteSessionInfoResponseContentContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s SiteSessionInfoResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoResponseContent) GoString() string {
  return s.String()
}

func (s *SiteSessionInfoResponseContent) SetSiteName(v string) *SiteSessionInfoResponseContent {
  s.SiteName = &v
  return s
}

func (s *SiteSessionInfoResponseContent) SetSiteId(v int64) *SiteSessionInfoResponseContent {
  s.SiteId = &v
  return s
}

func (s *SiteSessionInfoResponseContent) SetContent(v []*SiteSessionInfoResponseContentContent) *SiteSessionInfoResponseContent {
  s.Content = v
  return s
}

type SiteSessionInfoResponseContentContent struct     {
  // {"en":"Original IP","zh_CN":"原始IP"}
  SrcIp *string `json:"srcIp,omitempty" xml:"srcIp,omitempty" require:"true"`
  // {"en":"Original Port","zh_CN":"原始端口"}
  SrcPort *int `json:"srcPort,omitempty" xml:"srcPort,omitempty" require:"true"`
  // {"en":"Destination IP","zh_CN":"目的IP"}
  DestIp *string `json:"destIp,omitempty" xml:"destIp,omitempty" require:"true"`
  // {"en":"Destination Port","zh_CN":"目的端口"}
  DestPort *int `json:"destPort,omitempty" xml:"destPort,omitempty" require:"true"`
  // {"en":"protocol","zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"Flow size in byte","zh_CN":"流量大小，单位byte"}
  Flow *SiteSessionInfoResponseContentContentFlow `json:"flow,omitempty" xml:"flow,omitempty" require:"true" type:"Struct"`
}

func (s SiteSessionInfoResponseContentContent) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoResponseContentContent) GoString() string {
  return s.String()
}

func (s *SiteSessionInfoResponseContentContent) SetSrcIp(v string) *SiteSessionInfoResponseContentContent {
  s.SrcIp = &v
  return s
}

func (s *SiteSessionInfoResponseContentContent) SetSrcPort(v int) *SiteSessionInfoResponseContentContent {
  s.SrcPort = &v
  return s
}

func (s *SiteSessionInfoResponseContentContent) SetDestIp(v string) *SiteSessionInfoResponseContentContent {
  s.DestIp = &v
  return s
}

func (s *SiteSessionInfoResponseContentContent) SetDestPort(v int) *SiteSessionInfoResponseContentContent {
  s.DestPort = &v
  return s
}

func (s *SiteSessionInfoResponseContentContent) SetProtocol(v string) *SiteSessionInfoResponseContentContent {
  s.Protocol = &v
  return s
}

func (s *SiteSessionInfoResponseContentContent) SetFlow(v *SiteSessionInfoResponseContentContentFlow) *SiteSessionInfoResponseContentContent {
  s.Flow = v
  return s
}

type SiteSessionInfoResponseContentContentFlow struct {
}

func (s SiteSessionInfoResponseContentContentFlow) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoResponseContentContentFlow) GoString() string {
  return s.String()
}

type SiteSessionInfoResponseHeader struct {
}

func (s SiteSessionInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteSessionInfoResponseHeader) GoString() string {
  return s.String()
}




type SiteTotalFlowChartRequest struct {
}

func (s SiteTotalFlowChartRequest) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartRequest) GoString() string {
  return s.String()
}

type SiteTotalFlowChartRequestHeader struct {
}

func (s SiteTotalFlowChartRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartRequestHeader) GoString() string {
  return s.String()
}

type SiteTotalFlowChartPaths struct {
}

func (s SiteTotalFlowChartPaths) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartPaths) GoString() string {
  return s.String()
}

type SiteTotalFlowChartParameters struct {
  // {"en":"Site name\nWhen this parameter must be specified, query the specified site data","zh_CN":"站点名称\n该参数必须指定时，查询指定的站点数据"}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {"en":"1.the start time of the query, in the format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm\n2.must be less than the current time and endtime\n3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)\n4.can only query the last six months of data","zh_CN":"1、查询的开始时间，格式：yyyy-mm-ddThh:mm:ss 或 yyyy-mm-ddThh:mm\n2、必须小于当前时间和endtime\n3、startTime和endTime相差不能超过31天（可联系技术支持调整）\n4、只能查询最近半年内数据"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"1.the end time of the query, format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm\n2.must be less than the current time, more than starttime\n3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)\n4.can only query the last six months of data","zh_CN":"1、查询的结束时间，格式：yyyy-mm-ddThh:mm:ss或 yyyy-mm-ddThh:mm\n2、必须小于当前时间，大于starttime\n3、startTime和endTime相差不能超过31天（可联系技术支持调整）\n4、只能查询最近半年内数据"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time interval, enumerated\n1 = 1 minute\n2 = 5 minutes\nNo pass default 5 minutes","zh_CN":"时间间隔，枚举\n1 = 1分钟\n2 = 5分钟\n不传默认5分钟"}
  ChartAccuracy *int `json:"chartAccuracy,omitempty" xml:"chartAccuracy,omitempty"`
}

func (s SiteTotalFlowChartParameters) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartParameters) GoString() string {
  return s.String()
}

func (s *SiteTotalFlowChartParameters) SetSiteName(v string) *SiteTotalFlowChartParameters {
  s.SiteName = &v
  return s
}

func (s *SiteTotalFlowChartParameters) SetStartTime(v string) *SiteTotalFlowChartParameters {
  s.StartTime = &v
  return s
}

func (s *SiteTotalFlowChartParameters) SetEndTime(v string) *SiteTotalFlowChartParameters {
  s.EndTime = &v
  return s
}

func (s *SiteTotalFlowChartParameters) SetChartAccuracy(v int) *SiteTotalFlowChartParameters {
  s.ChartAccuracy = &v
  return s
}

type SiteTotalFlowChartResponse struct {
  // {"en":"Interface status codes","zh_CN":"接口状态码"}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Interface information","zh_CN":"接口信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"response content","zh_CN":"响应内容"}
  Content *SiteTotalFlowChartResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s SiteTotalFlowChartResponse) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponse) GoString() string {
  return s.String()
}

func (s *SiteTotalFlowChartResponse) SetReturnCode(v int) *SiteTotalFlowChartResponse {
  s.ReturnCode = &v
  return s
}

func (s *SiteTotalFlowChartResponse) SetReturnMsg(v string) *SiteTotalFlowChartResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SiteTotalFlowChartResponse) SetContent(v *SiteTotalFlowChartResponseContent) *SiteTotalFlowChartResponse {
  s.Content = v
  return s
}

type SiteTotalFlowChartResponseContent struct {
  // {"en":"site name","zh_CN":"站点名称"}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {"en":"site Id","zh_CN":"站点 ID"}
  SiteId *int64 `json:"siteId,omitempty" xml:"siteId,omitempty" require:"true"`
  // {"en":"Allocated bandwidth in Mbps","zh_CN":"分配带宽，单位：Mbps"}
  AllocateBandwidth *SiteTotalFlowChartResponseContentAllocateBandwidth `json:"allocateBandwidth,omitempty" xml:"allocateBandwidth,omitempty" require:"true" type:"Struct"`
  // {"en":"Bandwidth trend data","zh_CN":"带宽趋势数据"}
  Bandwidths *SiteTotalFlowChartResponseContentBandwidths `json:"bandwidths,omitempty" xml:"bandwidths,omitempty" require:"true" type:"Struct"`
}

func (s SiteTotalFlowChartResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponseContent) GoString() string {
  return s.String()
}

func (s *SiteTotalFlowChartResponseContent) SetSiteName(v string) *SiteTotalFlowChartResponseContent {
  s.SiteName = &v
  return s
}

func (s *SiteTotalFlowChartResponseContent) SetSiteId(v int64) *SiteTotalFlowChartResponseContent {
  s.SiteId = &v
  return s
}

func (s *SiteTotalFlowChartResponseContent) SetAllocateBandwidth(v *SiteTotalFlowChartResponseContentAllocateBandwidth) *SiteTotalFlowChartResponseContent {
  s.AllocateBandwidth = v
  return s
}

func (s *SiteTotalFlowChartResponseContent) SetBandwidths(v *SiteTotalFlowChartResponseContentBandwidths) *SiteTotalFlowChartResponseContent {
  s.Bandwidths = v
  return s
}

type SiteTotalFlowChartResponseContentAllocateBandwidth struct {
}

func (s SiteTotalFlowChartResponseContentAllocateBandwidth) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponseContentAllocateBandwidth) GoString() string {
  return s.String()
}

type SiteTotalFlowChartResponseContentBandwidths struct {
  // {"en":"timeStamp","zh_CN":"时间戳"}
  TimeStamp []*int64 `json:"timeStamp,omitempty" xml:"timeStamp,omitempty" require:"true" type:"Repeated"`
  // {"en":"Time slice","zh_CN":"时间分片"}
  TimeData []*string `json:"timeData,omitempty" xml:"timeData,omitempty" require:"true" type:"Repeated"`
  // {"en":"1min granularity, showing direct uplink bandwidth, 5min granularity, showing average uplink bandwidth over 5min, unit is byte","zh_CN":"单位为byte，1min粒度，直接显示上行流量，5min粒度，显示5min内的平均上行流量，带宽(Mbps)=流量/1000/1000/60"}
  UpStreamBandwidth []*SiteTotalFlowChartResponseContentBandwidthsUpStreamBandwidth `json:"upStreamBandwidth,omitempty" xml:"upStreamBandwidth,omitempty" require:"true" type:"Repeated"`
  // {"en":"1min granularity, showing direct downlink bandwidth, 5min granularity, showing average downlink bandwidth over 5min, unit is byte","zh_CN":"单位为byte，1min粒度，直接显示下行流量，5min粒度，显示5min内的平均下行流量，带宽（Mbps）=流量/1000/1000/60"}
  DownStreamBandwidth []*SiteTotalFlowChartResponseContentBandwidthsDownStreamBandwidth `json:"downStreamBandwidth,omitempty" xml:"downStreamBandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s SiteTotalFlowChartResponseContentBandwidths) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponseContentBandwidths) GoString() string {
  return s.String()
}

func (s *SiteTotalFlowChartResponseContentBandwidths) SetTimeStamp(v []*int64) *SiteTotalFlowChartResponseContentBandwidths {
  s.TimeStamp = v
  return s
}

func (s *SiteTotalFlowChartResponseContentBandwidths) SetTimeData(v []*string) *SiteTotalFlowChartResponseContentBandwidths {
  s.TimeData = v
  return s
}

func (s *SiteTotalFlowChartResponseContentBandwidths) SetUpStreamBandwidth(v []*SiteTotalFlowChartResponseContentBandwidthsUpStreamBandwidth) *SiteTotalFlowChartResponseContentBandwidths {
  s.UpStreamBandwidth = v
  return s
}

func (s *SiteTotalFlowChartResponseContentBandwidths) SetDownStreamBandwidth(v []*SiteTotalFlowChartResponseContentBandwidthsDownStreamBandwidth) *SiteTotalFlowChartResponseContentBandwidths {
  s.DownStreamBandwidth = v
  return s
}

type SiteTotalFlowChartResponseContentBandwidthsUpStreamBandwidth struct     {
}

func (s SiteTotalFlowChartResponseContentBandwidthsUpStreamBandwidth) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponseContentBandwidthsUpStreamBandwidth) GoString() string {
  return s.String()
}

type SiteTotalFlowChartResponseContentBandwidthsDownStreamBandwidth struct     {
}

func (s SiteTotalFlowChartResponseContentBandwidthsDownStreamBandwidth) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponseContentBandwidthsDownStreamBandwidth) GoString() string {
  return s.String()
}

type SiteTotalFlowChartResponseHeader struct {
}

func (s SiteTotalFlowChartResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteTotalFlowChartResponseHeader) GoString() string {
  return s.String()
}




type SaasSessionRequest struct {
}

func (s SaasSessionRequest) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionRequest) GoString() string {
  return s.String()
}

type SaasSessionResponse struct {
  // {'en':'Interface status codes', 'zh_CN':'接口状态码'}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {'en':'Interface information', 'zh_CN':'接口信息'}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {'en':'response content', 'zh_CN':'响应内容'}
  Content *SaasSessionResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s SaasSessionResponse) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionResponse) GoString() string {
  return s.String()
}

func (s *SaasSessionResponse) SetReturnCode(v int) *SaasSessionResponse {
  s.ReturnCode = &v
  return s
}

func (s *SaasSessionResponse) SetReturnMsg(v string) *SaasSessionResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SaasSessionResponse) SetContent(v *SaasSessionResponseContent) *SaasSessionResponse {
  s.Content = v
  return s
}

type SaasSessionResponseContent struct {
  // {'en':'site name', 'zh_CN':'站点名称'}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {'en':'site Id', 'zh_CN':'站点 ID'}
  SiteId *int64 `json:"siteId,omitempty" xml:"siteId,omitempty" require:"true"`
  // {'en':'Session data', 'zh_CN':'会话数据'}
  Content []*SaasSessionResponseContentContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s SaasSessionResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionResponseContent) GoString() string {
  return s.String()
}

func (s *SaasSessionResponseContent) SetSiteName(v string) *SaasSessionResponseContent {
  s.SiteName = &v
  return s
}

func (s *SaasSessionResponseContent) SetSiteId(v int64) *SaasSessionResponseContent {
  s.SiteId = &v
  return s
}

func (s *SaasSessionResponseContent) SetContent(v []*SaasSessionResponseContentContent) *SaasSessionResponseContent {
  s.Content = v
  return s
}

type SaasSessionResponseContentContent struct     {
  // {'en':'Original IP', 'zh_CN':'原始IP'}
  SrcIp *string `json:"srcIp,omitempty" xml:"srcIp,omitempty" require:"true"`
  // {'en':'Original Port', 'zh_CN':'原始端口'}
  SrcPort *int `json:"srcPort,omitempty" xml:"srcPort,omitempty" require:"true"`
  // {'en':'Destination IP', 'zh_CN':'目的IP'}
  DestIp *string `json:"destIp,omitempty" xml:"destIp,omitempty" require:"true"`
  // {'en':'Destination Port', 'zh_CN':'目的端口'}
  DestPort *int `json:"destPort,omitempty" xml:"destPort,omitempty" require:"true"`
  // {'en':'protocol', 'zh_CN':'协议'}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {'en':'inFlow size in byte', 'zh_CN':'流入流量大小，单位byte'}
  Rx *int `json:"rx,omitempty" xml:"rx,omitempty" require:"true"`
  // {'en':'outFlow size in byte', 'zh_CN':'流出流量大小，单位byte'}
  Tx *int `json:"tx,omitempty" xml:"tx,omitempty" require:"true"`
  // {'en':'sumFlow size in byte', 'zh_CN':'流入流出流量大小，单位byte'}
  Sumflow *int `json:"sumflow,omitempty" xml:"sumflow,omitempty" require:"true"`
}

func (s SaasSessionResponseContentContent) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionResponseContentContent) GoString() string {
  return s.String()
}

func (s *SaasSessionResponseContentContent) SetSrcIp(v string) *SaasSessionResponseContentContent {
  s.SrcIp = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetSrcPort(v int) *SaasSessionResponseContentContent {
  s.SrcPort = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetDestIp(v string) *SaasSessionResponseContentContent {
  s.DestIp = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetDestPort(v int) *SaasSessionResponseContentContent {
  s.DestPort = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetProtocol(v string) *SaasSessionResponseContentContent {
  s.Protocol = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetRx(v int) *SaasSessionResponseContentContent {
  s.Rx = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetTx(v int) *SaasSessionResponseContentContent {
  s.Tx = &v
  return s
}

func (s *SaasSessionResponseContentContent) SetSumflow(v int) *SaasSessionResponseContentContent {
  s.Sumflow = &v
  return s
}

type SaasSessionPaths struct {
}

func (s SaasSessionPaths) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionPaths) GoString() string {
  return s.String()
}

type SaasSessionParameters struct {
  // {'en':'Site name
  // When this parameter must be specified, query the specified site data', 'zh_CN':'站点名称
  // 该参数必须指定时，查询指定的站点数据'}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {'en':'1.the start time of the query, in the format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm
  // 2.must be less than the current time and endtime
  // 3.the difference between startTime and endTime cannot exceed 6 hours (can be adjusted by contacting technical support)
  // 4.can only query the last six months of data', 'zh_CN':'1、查询的开始时间，格式：yyyy-mm-ddThh:mm:ss 或 yyyy-mm-ddThh:mm
  // 2、必须小于当前时间和endtime
  // 3、startTime和endTime相差不能超过6小时（可联系技术支持调整）
  // 4、只能查询最近半年内数据'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'1.the end time of the query, format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm
  // 2.must be less than the current time, more than starttime
  // 3.the difference between startTime and endTime cannot exceed 6 hours (can be adjusted by contacting technical support)
  // 4.can only query the last six months of data', 'zh_CN':'1、查询的结束时间，格式：yyyy-mm-ddThh:mm:ss或 yyyy-mm-ddThh:mm
  // 2、必须小于当前时间，大于starttime
  // 3、startTime和endTime相差不能超过6小时（可联系技术支持调整）
  // 4、只能查询最近半年内数据'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'zh_CN':'查询IP地址（不区分源IP、目的IP）','en':'Source IP And Destination IP'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {'zh_CN':'查询端口（不区分源端口、目的端口）','en':'Source Port And Destination Port'}
  Port *int `json:"port,omitempty" xml:"port,omitempty"`
}

func (s SaasSessionParameters) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionParameters) GoString() string {
  return s.String()
}

func (s *SaasSessionParameters) SetSiteName(v string) *SaasSessionParameters {
  s.SiteName = &v
  return s
}

func (s *SaasSessionParameters) SetStartTime(v string) *SaasSessionParameters {
  s.StartTime = &v
  return s
}

func (s *SaasSessionParameters) SetEndTime(v string) *SaasSessionParameters {
  s.EndTime = &v
  return s
}

func (s *SaasSessionParameters) SetIp(v string) *SaasSessionParameters {
  s.Ip = &v
  return s
}

func (s *SaasSessionParameters) SetPort(v int) *SaasSessionParameters {
  s.Port = &v
  return s
}

type SaasSessionRequestHeader struct {
}

func (s SaasSessionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionRequestHeader) GoString() string {
  return s.String()
}

type SaasSessionResponseHeader struct {
}

func (s SaasSessionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SaasSessionResponseHeader) GoString() string {
  return s.String()
}




type SiteLogicChartRequest struct {
}

func (s SiteLogicChartRequest) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartRequest) GoString() string {
  return s.String()
}

type SiteLogicChartResponse struct {
  // {'en':'Interface status codes', 'zh_CN':'接口状态码'}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {'en':'Interface information', 'zh_CN':'接口信息'}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {'en':'response content', 'zh_CN':'响应内容'}
  Content *SiteLogicChartResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s SiteLogicChartResponse) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartResponse) GoString() string {
  return s.String()
}

func (s *SiteLogicChartResponse) SetReturnCode(v int) *SiteLogicChartResponse {
  s.ReturnCode = &v
  return s
}

func (s *SiteLogicChartResponse) SetReturnMsg(v string) *SiteLogicChartResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SiteLogicChartResponse) SetContent(v *SiteLogicChartResponseContent) *SiteLogicChartResponse {
  s.Content = v
  return s
}

type SiteLogicChartResponseContent struct {
  // {'en':'site A name', 'zh_CN':'站点A名称'}
  SiteAName *string `json:"siteAName,omitempty" xml:"siteAName,omitempty" require:"true"`
  // {'en':'site B name', 'zh_CN':'站点B名称'}
  SiteBName *string `json:"siteBName,omitempty" xml:"siteBName,omitempty" require:"true"`
  // {'en':'site A Id', 'zh_CN':'站点A ID'}
  SiteAId *int64 `json:"siteAId,omitempty" xml:"siteAId,omitempty" require:"true"`
  // {'en':'site B Id', 'zh_CN':'站点B ID'}
  SiteBId *int64 `json:"siteBId,omitempty" xml:"siteBId,omitempty" require:"true"`
  // {'en':'tunnel status 0-unknown, 1-online, 2-offline, 3-not reported, 4-suspend', 'zh_CN':'链路状态0–未知，1–在线，2-离线，3-未上报，4-挂起'}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'Quality trend data', 'zh_CN':'质量趋势数据'}
  Quality *SiteLogicChartResponseContentQuality `json:"quality,omitempty" xml:"quality,omitempty" require:"true" type:"Struct"`
}

func (s SiteLogicChartResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartResponseContent) GoString() string {
  return s.String()
}

func (s *SiteLogicChartResponseContent) SetSiteAName(v string) *SiteLogicChartResponseContent {
  s.SiteAName = &v
  return s
}

func (s *SiteLogicChartResponseContent) SetSiteBName(v string) *SiteLogicChartResponseContent {
  s.SiteBName = &v
  return s
}

func (s *SiteLogicChartResponseContent) SetSiteAId(v int64) *SiteLogicChartResponseContent {
  s.SiteAId = &v
  return s
}

func (s *SiteLogicChartResponseContent) SetSiteBId(v int64) *SiteLogicChartResponseContent {
  s.SiteBId = &v
  return s
}

func (s *SiteLogicChartResponseContent) SetStatus(v int) *SiteLogicChartResponseContent {
  s.Status = &v
  return s
}

func (s *SiteLogicChartResponseContent) SetQuality(v *SiteLogicChartResponseContentQuality) *SiteLogicChartResponseContent {
  s.Quality = v
  return s
}

type SiteLogicChartResponseContentQuality struct {
  // {'en':'timeStamp', 'zh_CN':'时间戳'}
  TimeStamp []*int64 `json:"timeStamp,omitempty" xml:"timeStamp,omitempty" require:"true" type:"Repeated"`
  // {'en':'Time slice', 'zh_CN':'时间分片'}
  TimeData []*string `json:"timeData,omitempty" xml:"timeData,omitempty" require:"true" type:"Repeated"`
  // {'en':'Time delay', 'zh_CN':'时延'}
  Rtt []*int `json:"rtt,omitempty" xml:"rtt,omitempty" require:"true" type:"Repeated"`
  // {'en':'Packet loss rate', 'zh_CN':'丢包率'}
  Loss []*int `json:"loss,omitempty" xml:"loss,omitempty" require:"true" type:"Repeated"`
  // {'en':'jitter', 'zh_CN':'抖动'}
  Mdev []*int `json:"mdev,omitempty" xml:"mdev,omitempty" require:"true" type:"Repeated"`
}

func (s SiteLogicChartResponseContentQuality) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartResponseContentQuality) GoString() string {
  return s.String()
}

func (s *SiteLogicChartResponseContentQuality) SetTimeStamp(v []*int64) *SiteLogicChartResponseContentQuality {
  s.TimeStamp = v
  return s
}

func (s *SiteLogicChartResponseContentQuality) SetTimeData(v []*string) *SiteLogicChartResponseContentQuality {
  s.TimeData = v
  return s
}

func (s *SiteLogicChartResponseContentQuality) SetRtt(v []*int) *SiteLogicChartResponseContentQuality {
  s.Rtt = v
  return s
}

func (s *SiteLogicChartResponseContentQuality) SetLoss(v []*int) *SiteLogicChartResponseContentQuality {
  s.Loss = v
  return s
}

func (s *SiteLogicChartResponseContentQuality) SetMdev(v []*int) *SiteLogicChartResponseContentQuality {
  s.Mdev = v
  return s
}

type SiteLogicChartPaths struct {
}

func (s SiteLogicChartPaths) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartPaths) GoString() string {
  return s.String()
}

type SiteLogicChartParameters struct {
  // {'en':'Site A name
  // When this parameter must be specified, query the specified site data', 'zh_CN':'站点A名称
  // 该参数必须指定时，查询指定的站点数据'}
  SiteAName *string `json:"siteAName,omitempty" xml:"siteAName,omitempty" require:"true"`
  // {'en':'Site B name
  // When this parameter must be specified, query the specified site data', 'zh_CN':'站点B名称
  // 该参数必须指定时，查询指定的站点数据'}
  SiteBName *string `json:"siteBName,omitempty" xml:"siteBName,omitempty" require:"true"`
  // {'en':'1.the start time of the query, in the format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm
  // 2.must be less than the current time and endtime
  // 3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)
  // 4.can only query the last six months of data', 'zh_CN':'1、查询的开始时间，格式：yyyy-mm-ddThh:mm:ss 或 yyyy-mm-ddThh:mm
  // 2、必须小于当前时间和endtime
  // 3、startTime和endTime相差不能超过31天（可联系技术支持调整）
  // 4、只能查询最近半年内数据'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'1.the end time of the query, format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm
  // 2.must be less than the current time, more than starttime
  // 3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)
  // 4.can only query the last six months of data', 'zh_CN':'1、查询的结束时间，格式：yyyy-mm-ddThh:mm:ss或 yyyy-mm-ddThh:mm
  // 2、必须小于当前时间，大于starttime
  // 3、startTime和endTime相差不能超过31天（可联系技术支持调整）
  // 4、只能查询最近半年内数据'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s SiteLogicChartParameters) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartParameters) GoString() string {
  return s.String()
}

func (s *SiteLogicChartParameters) SetSiteAName(v string) *SiteLogicChartParameters {
  s.SiteAName = &v
  return s
}

func (s *SiteLogicChartParameters) SetSiteBName(v string) *SiteLogicChartParameters {
  s.SiteBName = &v
  return s
}

func (s *SiteLogicChartParameters) SetStartTime(v string) *SiteLogicChartParameters {
  s.StartTime = &v
  return s
}

func (s *SiteLogicChartParameters) SetEndTime(v string) *SiteLogicChartParameters {
  s.EndTime = &v
  return s
}

type SiteLogicChartRequestHeader struct {
}

func (s SiteLogicChartRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartRequestHeader) GoString() string {
  return s.String()
}

type SiteLogicChartResponseHeader struct {
}

func (s SiteLogicChartResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteLogicChartResponseHeader) GoString() string {
  return s.String()
}




type SiteQualityChartRequest struct {
}

func (s SiteQualityChartRequest) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartRequest) GoString() string {
  return s.String()
}

type SiteQualityChartRequestHeader struct {
}

func (s SiteQualityChartRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartRequestHeader) GoString() string {
  return s.String()
}

type SiteQualityChartPaths struct {
}

func (s SiteQualityChartPaths) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartPaths) GoString() string {
  return s.String()
}

type SiteQualityChartParameters struct {
  // {"en":"Site name\nWhen this parameter must be specified, query the specified site data","zh_CN":"站点名称\n该参数必须指定时，查询指定的站点数据"}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {"en":"1.the start time of the query, in the format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm\n2.must be less than the current time and endtime\n3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)\n4.can only query the last six months of data","zh_CN":"1、查询的开始时间，格式：yyyy-mm-ddThh:mm:ss 或 yyyy-mm-ddThh:mm\n2、必须小于当前时间和endtime\n3、startTime和endTime相差不能超过31天（可联系技术支持调整）\n4、只能查询最近半年内数据"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"1.the end time of the query, format: yyyy-mm-ddThh:mm:ss or yyyy-mm-ddThh:mm\n2.must be less than the current time, more than starttime\n3.the difference between startTime and endTime cannot exceed 31 days (can be adjusted by contacting technical support)\n4.can only query the last six months of data","zh_CN":"1、查询的结束时间，格式：yyyy-mm-ddThh:mm:ss或 yyyy-mm-ddThh:mm\n2、必须小于当前时间，大于starttime\n3、startTime和endTime相差不能超过31天（可联系技术支持调整）\n4、只能查询最近半年内数据"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s SiteQualityChartParameters) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartParameters) GoString() string {
  return s.String()
}

func (s *SiteQualityChartParameters) SetSiteName(v string) *SiteQualityChartParameters {
  s.SiteName = &v
  return s
}

func (s *SiteQualityChartParameters) SetStartTime(v string) *SiteQualityChartParameters {
  s.StartTime = &v
  return s
}

func (s *SiteQualityChartParameters) SetEndTime(v string) *SiteQualityChartParameters {
  s.EndTime = &v
  return s
}

type SiteQualityChartResponse struct {
  // {"en":"Interface status codes","zh_CN":"接口状态码"}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Interface information","zh_CN":"接口信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"response content","zh_CN":"响应内容"}
  Content *SiteQualityChartResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s SiteQualityChartResponse) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponse) GoString() string {
  return s.String()
}

func (s *SiteQualityChartResponse) SetReturnCode(v int) *SiteQualityChartResponse {
  s.ReturnCode = &v
  return s
}

func (s *SiteQualityChartResponse) SetReturnMsg(v string) *SiteQualityChartResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SiteQualityChartResponse) SetContent(v *SiteQualityChartResponseContent) *SiteQualityChartResponse {
  s.Content = v
  return s
}

type SiteQualityChartResponseContent struct {
  // {"en":"site name","zh_CN":"站点名称"}
  SiteName *string `json:"siteName,omitempty" xml:"siteName,omitempty" require:"true"`
  // {"en":"site Id","zh_CN":"站点 ID"}
  SiteId *int64 `json:"siteId,omitempty" xml:"siteId,omitempty" require:"true"`
  // {"en":"Tunnel quality data under site","zh_CN":"站点下的隧道质量数据"}
  Tunnel []*SiteQualityChartResponseContentTunnel `json:"tunnel,omitempty" xml:"tunnel,omitempty" require:"true" type:"Repeated"`
}

func (s SiteQualityChartResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponseContent) GoString() string {
  return s.String()
}

func (s *SiteQualityChartResponseContent) SetSiteName(v string) *SiteQualityChartResponseContent {
  s.SiteName = &v
  return s
}

func (s *SiteQualityChartResponseContent) SetSiteId(v int64) *SiteQualityChartResponseContent {
  s.SiteId = &v
  return s
}

func (s *SiteQualityChartResponseContent) SetTunnel(v []*SiteQualityChartResponseContentTunnel) *SiteQualityChartResponseContent {
  s.Tunnel = v
  return s
}

type SiteQualityChartResponseContentTunnel struct     {
  // {"en":"tunnel name","zh_CN":"链路名称"}
  TunnelName *string `json:"tunnelName,omitempty" xml:"tunnelName,omitempty" require:"true"`
  // {"en":"tunnel Id","zh_CN":"链路 ID"}
  TunnelId *int64 `json:"tunnelId,omitempty" xml:"tunnelId,omitempty" require:"true"`
  // {"en":"tunnel status 0-unknown, 1-online, 2-offline, 3-not reported, 4-suspend","zh_CN":"链路状态0–未知，1–在线，2-离线，3-未上报，4-挂起"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"timeStamp","zh_CN":"时间戳"}
  TimeStamp []*int64 `json:"timeStamp,omitempty" xml:"timeStamp,omitempty" require:"true" type:"Repeated"`
  // {"en":"Time slice","zh_CN":"时间分片"}
  TimeData []*string `json:"timeData,omitempty" xml:"timeData,omitempty" require:"true" type:"Repeated"`
  // {"en":"Time delay","zh_CN":"时延"}
  Rtt []*SiteQualityChartResponseContentTunnelRtt `json:"rtt,omitempty" xml:"rtt,omitempty" require:"true" type:"Repeated"`
  // {"en":"Packet loss rate","zh_CN":"丢包率"}
  Loss []*SiteQualityChartResponseContentTunnelLoss `json:"loss,omitempty" xml:"loss,omitempty" require:"true" type:"Repeated"`
  // {"en":"jitter","zh_CN":"抖动"}
  Mdev []*SiteQualityChartResponseContentTunnelMdev `json:"mdev,omitempty" xml:"mdev,omitempty" require:"true" type:"Repeated"`
}

func (s SiteQualityChartResponseContentTunnel) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponseContentTunnel) GoString() string {
  return s.String()
}

func (s *SiteQualityChartResponseContentTunnel) SetTunnelName(v string) *SiteQualityChartResponseContentTunnel {
  s.TunnelName = &v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetTunnelId(v int64) *SiteQualityChartResponseContentTunnel {
  s.TunnelId = &v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetStatus(v int) *SiteQualityChartResponseContentTunnel {
  s.Status = &v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetTimeStamp(v []*int64) *SiteQualityChartResponseContentTunnel {
  s.TimeStamp = v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetTimeData(v []*string) *SiteQualityChartResponseContentTunnel {
  s.TimeData = v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetRtt(v []*SiteQualityChartResponseContentTunnelRtt) *SiteQualityChartResponseContentTunnel {
  s.Rtt = v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetLoss(v []*SiteQualityChartResponseContentTunnelLoss) *SiteQualityChartResponseContentTunnel {
  s.Loss = v
  return s
}

func (s *SiteQualityChartResponseContentTunnel) SetMdev(v []*SiteQualityChartResponseContentTunnelMdev) *SiteQualityChartResponseContentTunnel {
  s.Mdev = v
  return s
}

type SiteQualityChartResponseContentTunnelRtt struct     {
}

func (s SiteQualityChartResponseContentTunnelRtt) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponseContentTunnelRtt) GoString() string {
  return s.String()
}

type SiteQualityChartResponseContentTunnelLoss struct     {
}

func (s SiteQualityChartResponseContentTunnelLoss) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponseContentTunnelLoss) GoString() string {
  return s.String()
}

type SiteQualityChartResponseContentTunnelMdev struct     {
}

func (s SiteQualityChartResponseContentTunnelMdev) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponseContentTunnelMdev) GoString() string {
  return s.String()
}

type SiteQualityChartResponseHeader struct {
}

func (s SiteQualityChartResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SiteQualityChartResponseHeader) GoString() string {
  return s.String()
}




