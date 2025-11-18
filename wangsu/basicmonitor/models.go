package basicmonitor

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportFlowDomainIspProvinceIaasServiceRequest struct {
  // {'en':'Starting time:
  // 
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00) );
  // 
  // 2. Cannot be greater than the current time
  // 
  // 3. Get up to the last six months (183 days) of data.', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 
  // 1. Time format 2016-12-02T10:00:00+08:00
  // 
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 
  // 3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception
  // 
  // 4. Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo can't exceed 7 days (can contact technical support adjustment).", "zh_CN":"结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'domain:
  // 
  // 1. The maximum number of passable domain names is 20 by default (you can contact technical support adjustment). All domain names under the account are not queried when the entry is not passed, but cannot be queried (error) when the number of domain names under the account exceeds the limit.
  // 
  // 2. Automatically filter out illegal domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of legitimate domain names).', 'zh_CN':'域名：
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）。未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'Data granularity, the default is 5m (5 minutes, currently only supports 5m)', 'zh_CN':'数据粒度，默认为5m（5分钟，当前只支持5m）'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Province
  // 
  // 1.Province is not upload: Query all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.', 'zh_CN':'省份
  // 
  // 1.未传递province时：查询所有省份
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。'}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {'en':'ISP:
  // 1.ISP is not upload: Query all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.', 'zh_CN':'运营商：
  // 1.未传递isp时：查询所有isp；
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节'}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {'en':'Traffic type:
  // 
  // 1. Passable value range: trafficOut (outflow), trafficIn (inflow)
  // 
  // 2. Multiple deliverables, no default outgoing, and merged traffic when multiple passes.', 'zh_CN':'流量类型：
  // 1.可传递值范围：trafficOut（流出）、trafficIn（流入）
  // 2.可传递多个，不传默认流出，传递多个时返回合并流量。'}
  Type []*string `json:"type,omitempty" xml:"type,omitempty" type:"Repeated"`
  // {"en":"Grouping dimensions:
  // 
  // 1. The optional values are domain,province, and isp, which can pass in single or multiple values.
  // 
  // 2. If there is an incoming, the detailed data will be displayed according to the dimension;
  // 
  // 3. The result hierarchy is fixed in order, and the order of the parameters does not affect the order of the returned results. For example: 'groupBy': ['domain','province'] is the same as 'groupBy': ['province', 'domain'].
  // 
  // For example: pass 'groupBy': ['domain', 'province'], then the isp node under ispData does not need to return.", "zh_CN":"分组维度：
  // 1.可选值为domain、province、isp，可传入单个或多个值；
  // 2.有传入则按照该维度展示明细数据；
  // 3.返回结果层级顺序固定，入参顺序不影响返回结果顺序。例如：'groupBy': ['domain','province']与'groupBy': ['province','domain']返回结果一样。
  // 例如：传递'groupBy': ['domain','province']，则ispData下的isp节点无需返回。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowDomainIspProvinceIaasServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetDateFrom(v string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetDateTo(v string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetDomain(v []*string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetDataInterval(v string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetProvince(v []*string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.Province = v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetIsp(v []*string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetType(v []*string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.Type = v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceRequest) SetGroupBy(v []*string) *ReportFlowDomainIspProvinceIaasServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowDomainIspProvinceIaasServiceResponse struct {
  Result []*ReportFlowDomainIspProvinceIaasServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainIspProvinceIaasServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponse) SetResult(v []*ReportFlowDomainIspProvinceIaasServiceResponseResult) *ReportFlowDomainIspProvinceIaasServiceResponse {
  s.Result = v
  return s
}

type ReportFlowDomainIspProvinceIaasServiceResponseResult struct     {
  // {'en':'domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportFlowDomainIspProvinceIaasServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResult) SetDomain(v string) *ReportFlowDomainIspProvinceIaasServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResult) SetIspData(v []*ReportFlowDomainIspProvinceIaasServiceResponseResultIspData) *ReportFlowDomainIspProvinceIaasServiceResponseResult {
  s.IspData = v
  return s
}

type ReportFlowDomainIspProvinceIaasServiceResponseResultIspData struct     {
  // {'en':'Service provider&rsquo;s Chinese name', 'zh_CN':'运营商中文名称'}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResultIspData) SetIsp(v string) *ReportFlowDomainIspProvinceIaasServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResultIspData) SetProvinceData(v []*ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData) *ReportFlowDomainIspProvinceIaasServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData struct     {
  // {'en':'Chinese name of the province', 'zh_CN':'省份中文名称'}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  FlowData []*ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData) SetFlowData(v []*ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData) *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceData {
  s.FlowData = v
  return s
}

type ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData struct     {
  // {'en':'time
  // 1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00.
  // 3. Returns the time slice contained in the start time and end time.', 'zh_CN':'时间
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd &nbsp; HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00。
  // 2.查询的数据粒度为1h时，格式为yyyy-MM-dd &nbsp; HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） &nbsp; 00。
  // 3.返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'The value of the flow, the unit is MB, and the two decimal digits are retained', 'zh_CN':'流量值，单位为MB，保留两位小数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData) SetTimestamp(v string) *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData) SetValue(v string) *ReportFlowDomainIspProvinceIaasServiceResponseResultIspDataProvinceDataFlowData {
  s.Value = &v
  return s
}

type ReportFlowDomainIspProvinceIaasServicePaths struct {
}

func (s ReportFlowDomainIspProvinceIaasServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServicePaths) GoString() string {
  return s.String()
}

type ReportFlowDomainIspProvinceIaasServiceParameters struct {
}

func (s ReportFlowDomainIspProvinceIaasServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowDomainIspProvinceIaasServiceRequestHeader struct {
}

func (s ReportFlowDomainIspProvinceIaasServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowDomainIspProvinceIaasServiceResponseHeader struct {
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainIspProvinceIaasServiceResponseHeader) GoString() string {
  return s.String()
}




type LECHQueryServerMetricRequest struct {
}

func (s LECHQueryServerMetricRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricRequest) GoString() string {
  return s.String()
}

type LECHQueryServerMetricRequestHeader struct {
}

func (s LECHQueryServerMetricRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricRequestHeader) GoString() string {
  return s.String()
}

type LECHQueryServerMetricPaths struct {
}

func (s LECHQueryServerMetricPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricPaths) GoString() string {
  return s.String()
}

type LECHQueryServerMetricParameters struct {
  // {"en":"The ID of the virtual machine to query. At most 30 queries can be made at a time, ids\nare separated by character  ','","zh_CN":"云主机ID。单次最多查询 20 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Monitoring data type of query, value:\nCPU: CPU utilization\nMem: memory\nBandwidth: Traffic data","zh_CN":"查询的监控数据类型，取值：\ncpu： cpu使用率\nmem：内存\nbandwidth：流量数据"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Query time range in format: yyyymmddhhmm-yyyymmddhhmm\nQuery data for nearly 90 days, a single query no more than 3 days.\nFor example: 202001201730-202001201930 means to query 2020-01-20 17:30 to 19:30 monitoring data.","zh_CN":"查询时间范围，格式：yyyyMMddHHmm- yyyyMMddHHmm\n查询近90天的数据，单次查询不超过10天。\n例如：202001201730-202001201930表示查询2020-01-20 17:30到19:30的监控数据。"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"The ISP code","zh_CN":"当查询流量数据时可填。实例所属运营商：dx-电信；wt-网通；yd-移动。一次仅允许传入一个运营商参数"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
}

func (s LECHQueryServerMetricParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricParameters) GoString() string {
  return s.String()
}

func (s *LECHQueryServerMetricParameters) SetIds(v string) *LECHQueryServerMetricParameters {
  s.Ids = &v
  return s
}

func (s *LECHQueryServerMetricParameters) SetType(v string) *LECHQueryServerMetricParameters {
  s.Type = &v
  return s
}

func (s *LECHQueryServerMetricParameters) SetStatTime(v string) *LECHQueryServerMetricParameters {
  s.StatTime = &v
  return s
}

func (s *LECHQueryServerMetricParameters) SetCarrier(v string) *LECHQueryServerMetricParameters {
  s.Carrier = &v
  return s
}

type LECHQueryServerMetricResponse struct {
  // {"en":"Instance information array","zh_CN":"实例信息数组"}
  Servers []*LECHQueryServerMetricResponseServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryServerMetricResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricResponse) GoString() string {
  return s.String()
}

func (s *LECHQueryServerMetricResponse) SetServers(v []*LECHQueryServerMetricResponseServers) *LECHQueryServerMetricResponse {
  s.Servers = v
  return s
}

type LECHQueryServerMetricResponseServers struct     {
  // {"en":"instance id","zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"instance ip","zh_CN":"实例公网IP"}
  InstanceIp *string `json:"instanceIp,omitempty" xml:"instanceIp,omitempty" require:"true"`
  // {"en":"Bandwidth information","zh_CN":"带宽信息"}
  Bandwidths []*LECHQueryServerMetricResponseServersBandwidths `json:"bandwidths,omitempty" xml:"bandwidths,omitempty" require:"true" type:"Repeated"`
  // {"en":"Memory information","zh_CN":"内存信息"}
  Mems []*LECHQueryServerMetricResponseServersMems `json:"mems,omitempty" xml:"mems,omitempty" require:"true" type:"Repeated"`
  // {"en":"CPU information","zh_CN":"CPU信息"}
  Cpus []*LECHQueryServerMetricResponseServersCpus `json:"cpus,omitempty" xml:"cpus,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQueryServerMetricResponseServers) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricResponseServers) GoString() string {
  return s.String()
}

func (s *LECHQueryServerMetricResponseServers) SetId(v string) *LECHQueryServerMetricResponseServers {
  s.Id = &v
  return s
}

func (s *LECHQueryServerMetricResponseServers) SetInstanceIp(v string) *LECHQueryServerMetricResponseServers {
  s.InstanceIp = &v
  return s
}

func (s *LECHQueryServerMetricResponseServers) SetBandwidths(v []*LECHQueryServerMetricResponseServersBandwidths) *LECHQueryServerMetricResponseServers {
  s.Bandwidths = v
  return s
}

func (s *LECHQueryServerMetricResponseServers) SetMems(v []*LECHQueryServerMetricResponseServersMems) *LECHQueryServerMetricResponseServers {
  s.Mems = v
  return s
}

func (s *LECHQueryServerMetricResponseServers) SetCpus(v []*LECHQueryServerMetricResponseServersCpus) *LECHQueryServerMetricResponseServers {
  s.Cpus = v
  return s
}

type LECHQueryServerMetricResponseServersBandwidths struct     {
  // {"en":"The bandwidth collection time, formatted as YYYYMMDDHHMM, is 5 minutes granularity","zh_CN":"带宽采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"Total outflow bandwidth (including intra-node flow) in Mbps","zh_CN":"总流出带宽（含节点内流量），单位Mbps"}
  Out *string `json:"out,omitempty" xml:"out,omitempty" require:"true"`
  // {"en":"Total inflow bandwidth (including intra-node traffic) in Mbps","zh_CN":"总流入带宽（含节点内流量），单位Mbps"}
  In *string `json:"in,omitempty" xml:"in,omitempty" require:"true"`
  // {"en":"Inflow bandwidth of external network (not including intra-node traffic), unit of Mbps","zh_CN":"外网流入带宽（不含节点内流量），单位Mbps"}
  ExtIn *string `json:"extIn,omitempty" xml:"extIn,omitempty" require:"true"`
  // {"en":"Outflow bandwidth of the external network (excluding intra-node traffic), unit of Mbps","zh_CN":"外网流出带宽（不含节点内流量），单位Mbps"}
  ExtOut *string `json:"extOut,omitempty" xml:"extOut,omitempty" require:"true"`
}

func (s LECHQueryServerMetricResponseServersBandwidths) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricResponseServersBandwidths) GoString() string {
  return s.String()
}

func (s *LECHQueryServerMetricResponseServersBandwidths) SetStatTime(v string) *LECHQueryServerMetricResponseServersBandwidths {
  s.StatTime = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersBandwidths) SetOut(v string) *LECHQueryServerMetricResponseServersBandwidths {
  s.Out = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersBandwidths) SetIn(v string) *LECHQueryServerMetricResponseServersBandwidths {
  s.In = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersBandwidths) SetExtIn(v string) *LECHQueryServerMetricResponseServersBandwidths {
  s.ExtIn = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersBandwidths) SetExtOut(v string) *LECHQueryServerMetricResponseServersBandwidths {
  s.ExtOut = &v
  return s
}

type LECHQueryServerMetricResponseServersMems struct     {
  // {"en":"The memory data collection time, in the format YYYYMMDDHHMM, is 5 minutes granularity","zh_CN":"内存数据采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"Memory usage in MB","zh_CN":"内存使用量，单位为MB"}
  Used *string `json:"used,omitempty" xml:"used,omitempty" require:"true"`
  // {"en":"Total memory, in MB","zh_CN":"总内存量，单位为MB"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s LECHQueryServerMetricResponseServersMems) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricResponseServersMems) GoString() string {
  return s.String()
}

func (s *LECHQueryServerMetricResponseServersMems) SetStatTime(v string) *LECHQueryServerMetricResponseServersMems {
  s.StatTime = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersMems) SetUsed(v string) *LECHQueryServerMetricResponseServersMems {
  s.Used = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersMems) SetTotal(v string) *LECHQueryServerMetricResponseServersMems {
  s.Total = &v
  return s
}

type LECHQueryServerMetricResponseServersCpus struct     {
  // {"en":"CPU usage collection time in the format YYYYMMDDHHMM, with a granularity of 5 minutes","zh_CN":"cpu使用率采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"CPU utilization, with a value of 0-1, such as 0.25, represents 25% utilization","zh_CN":"cpu使用率，取值0-1，如0.25，该值则表示使用率为25%"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
}

func (s LECHQueryServerMetricResponseServersCpus) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricResponseServersCpus) GoString() string {
  return s.String()
}

func (s *LECHQueryServerMetricResponseServersCpus) SetStatTime(v string) *LECHQueryServerMetricResponseServersCpus {
  s.StatTime = &v
  return s
}

func (s *LECHQueryServerMetricResponseServersCpus) SetUsage(v string) *LECHQueryServerMetricResponseServersCpus {
  s.Usage = &v
  return s
}

type LECHQueryServerMetricResponseHeader struct {
}

func (s LECHQueryServerMetricResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQueryServerMetricResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryServerMetricRequest struct {
}

func (s VMPQueryServerMetricRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricRequest) GoString() string {
  return s.String()
}

type VMPQueryServerMetricResponse struct {
  // {"en":"Instance information array", "zh_CN":"实例信息数组"}
  Servers []*VMPQueryServerMetricServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryServerMetricResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricResponse) SetServers(v []*VMPQueryServerMetricServer) *VMPQueryServerMetricResponse {
  s.Servers = v
  return s
}

type VMPQueryServerMetricServer struct {
  // {"en":"instance id", "zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"instance ip", "zh_CN":"实例公网IP"}
  InstanceIp *string `json:"instanceIp,omitempty" xml:"instanceIp,omitempty" require:"true"`
  // {"en":"VMPQueryServerMetricBandwidth information", "zh_CN":"带宽信息"}
  Bandwidths []*VMPQueryServerMetricBandwidth `json:"bandwidths,omitempty" xml:"bandwidths,omitempty" require:"true" type:"Repeated"`
  // {"en":"Memory information", "zh_CN":"内存信息"}
  Mems []*VMPQueryServerMetricMem `json:"mems,omitempty" xml:"mems,omitempty" require:"true" type:"Repeated"`
  // {"en":"CPU information", "zh_CN":"CPU信息"}
  Cpus []*VMPQueryServerMetricCpu `json:"cpus,omitempty" xml:"cpus,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryServerMetricServer) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricServer) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricServer) SetId(v string) *VMPQueryServerMetricServer {
  s.Id = &v
  return s
}

func (s *VMPQueryServerMetricServer) SetInstanceIp(v string) *VMPQueryServerMetricServer {
  s.InstanceIp = &v
  return s
}

func (s *VMPQueryServerMetricServer) SetBandwidths(v []*VMPQueryServerMetricBandwidth) *VMPQueryServerMetricServer {
  s.Bandwidths = v
  return s
}

func (s *VMPQueryServerMetricServer) SetMems(v []*VMPQueryServerMetricMem) *VMPQueryServerMetricServer {
  s.Mems = v
  return s
}

func (s *VMPQueryServerMetricServer) SetCpus(v []*VMPQueryServerMetricCpu) *VMPQueryServerMetricServer {
  s.Cpus = v
  return s
}

type VMPQueryServerMetricBandwidth struct {
  // {"en":"The bandwidth collection time, formatted as YYYYMMDDHHMM, is 5 minutes granularity", "zh_CN":"带宽采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"Total outflow bandwidth (including intra-node flow) in Mbps", "zh_CN":"总流出带宽（含节点内流量），单位Mbps"}
  Out *string `json:"out,omitempty" xml:"out,omitempty" require:"true"`
  // {"en":"Total inflow bandwidth (including intra-node traffic) in Mbps", "zh_CN":"总流入带宽（含节点内流量），单位Mbps"}
  In *string `json:"in,omitempty" xml:"in,omitempty" require:"true"`
  // {"en":"Inflow bandwidth of external network (not including intra-node traffic), unit of Mbps", "zh_CN":"外网流入带宽（不含节点内流量），单位Mbps"}
  ExtIn *string `json:"extIn,omitempty" xml:"extIn,omitempty" require:"true"`
  // {"en":"Outflow bandwidth of the external network (excluding intra-node traffic), unit of Mbps", "zh_CN":"外网流出带宽（不含节点内流量），单位Mbps"}
  ExtOut *string `json:"extOut,omitempty" xml:"extOut,omitempty" require:"true"`
}

func (s VMPQueryServerMetricBandwidth) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricBandwidth) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricBandwidth) SetStatTime(v string) *VMPQueryServerMetricBandwidth {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetOut(v string) *VMPQueryServerMetricBandwidth {
  s.Out = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetIn(v string) *VMPQueryServerMetricBandwidth {
  s.In = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetExtIn(v string) *VMPQueryServerMetricBandwidth {
  s.ExtIn = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetExtOut(v string) *VMPQueryServerMetricBandwidth {
  s.ExtOut = &v
  return s
}

type VMPQueryServerMetricMem struct {
  // {"en":"The memory data collection time, in the format YYYYMMDDHHMM, is 5 minutes granularity", "zh_CN":"内存数据采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"Memory usage in MB", "zh_CN":"内存使用量，单位为MB"}
  Used *string `json:"used,omitempty" xml:"used,omitempty" require:"true"`
  // {"en":"Total memory, in MB", "zh_CN":"总内存量，单位为MB"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s VMPQueryServerMetricMem) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricMem) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricMem) SetStatTime(v string) *VMPQueryServerMetricMem {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricMem) SetUsed(v string) *VMPQueryServerMetricMem {
  s.Used = &v
  return s
}

func (s *VMPQueryServerMetricMem) SetTotal(v string) *VMPQueryServerMetricMem {
  s.Total = &v
  return s
}

type VMPQueryServerMetricCpu struct {
  // {"en":"CPU usage collection time in the format YYYYMMDDHHMM, with a granularity of 5 minutes", "zh_CN":"cpu使用率采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"CPU utilization, with a value of 0-1, such as 0.25, represents 25% utilization", "zh_CN":"cpu使用率，取值0-1，如0.25，该值则表示使用率为25%"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
}

func (s VMPQueryServerMetricCpu) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricCpu) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricCpu) SetStatTime(v string) *VMPQueryServerMetricCpu {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricCpu) SetUsage(v string) *VMPQueryServerMetricCpu {
  s.Usage = &v
  return s
}

type VMPQueryServerMetricPaths struct {
}

func (s VMPQueryServerMetricPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricPaths) GoString() string {
  return s.String()
}

type VMPQueryServerMetricParameters struct {
  // {"en":"The ID of the virtual machine to query. At most 30 queries can be made at a time, ids 
  //  are separated by character  ','", "zh_CN":"云主机ID。单次最多查询 20 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Monitoring data type of query, value:
  // CPU: CPU utilization
  // VMPQueryServerMetricMem: memory
  // VMPQueryServerMetricBandwidth: Traffic data", "zh_CN":"查询的监控数据类型，取值：
  // cpu： cpu使用率
  // mem：内存
  // bandwidth：流量数据"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Query time range in format: yyyymmddhhmm-yyyymmddhhmm
  // Query data for nearly 90 days, a single query no more than 3 days.
  // For example: 202001201730-202001201930 means to query 2020-01-20 17:30 to 19:30 monitoring data.", "zh_CN":"查询时间范围，格式：yyyyMMddHHmm- yyyyMMddHHmm
  // 查询近90天的数据，单次查询不超过10天。
  // 例如：202001201730-202001201930表示查询2020-01-20 17:30到19:30的监控数据。"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"The ISP code  ", "zh_CN":"当查询流量数据时可填。实例所属运营商：dx-电信；wt-网通；yd-移动。一次仅允许传入一个运营商参数"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
}

func (s VMPQueryServerMetricParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricParameters) SetIds(v string) *VMPQueryServerMetricParameters {
  s.Ids = &v
  return s
}

func (s *VMPQueryServerMetricParameters) SetType(v string) *VMPQueryServerMetricParameters {
  s.Type = &v
  return s
}

func (s *VMPQueryServerMetricParameters) SetStatTime(v string) *VMPQueryServerMetricParameters {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricParameters) SetCarrier(v string) *VMPQueryServerMetricParameters {
  s.Carrier = &v
  return s
}

type VMPQueryServerMetricRequestHeader struct {
}

func (s VMPQueryServerMetricRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryServerMetricResponseHeader struct {
}

func (s VMPQueryServerMetricResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricResponseHeader) GoString() string {
  return s.String()
}




