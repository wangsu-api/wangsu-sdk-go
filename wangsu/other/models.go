package other

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type MtrTestRequest struct {
  // {"en":"需要检测的域名或 IP", "zh_CN":"需要检测的域名或 IP"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"监控机所属地区中文名称，支持中国大陆省份、港澳台以及海外国家。
  // 可选值：
  // anhui: 安徽
  // beijing: 北京
  // chongqing: 重庆
  // fujian: 福建
  // gansu: 甘肃
  // guangdong: 广东
  // guangxi: 广西
  // guizhou: 贵州
  // hainan: 海南
  // hebei: 河北
  // heilongjiang: 黑龙江
  // henan: 河南
  // hubei: 湖北
  // hunan: 湖南
  // jiangsu: 江苏
  // jiangxi: 江西
  // jilin: 吉林
  // liaoning: 辽宁
  // neimenggu: 内蒙古
  // ningxia: 宁夏
  // qinghai: 青海
  // shaanxi: 陕西
  // shandong: 山东
  // shanghai: 上海
  // shanxi: 山西
  // sichuan: 四川
  // tianjin: 天津
  // xinjiang: 新疆
  // xizang: 西藏
  // yunnan: 云南
  // zhejiang: 浙江
  // TW: 台湾
  // HK: 香港
  // MO: 澳门
  // AE: 阿联酋
  // AU: 澳大利亚
  // BD: 孟加拉
  // BN: 文莱
  // BR: 巴西
  // CA: 加拿大
  // CL: 智利
  // CO: 哥伦比亚
  // DJ: 吉布提
  // ID: 印度尼西亚
  // IT: 意大利
  // JP: 日本
  // KG: 吉尔吉斯斯坦
  // KH: 柬埔寨
  // KR: 韩国
  // KW: 科威特
  // LA: 老挝
  // MG: 马达加斯加
  // MM: 缅甸
  // MU: 毛里求斯
  // MY: 马来西亚
  // NP: 尼泊尔
  // OM: 阿曼
  // PE: 秘鲁
  // PH: 菲律宾
  // PK: 巴基斯坦
  // QA: 卡塔尔
  // RO: 罗马尼亚
  // RU: 俄罗斯
  // SA: 沙特阿拉伯
  // SE: 瑞典
  // SG: 新加坡
  // TH: 泰国
  // TR: 土耳其
  // US: 美国
  // VN: 越南", "zh_CN":"监控机所属地区中文名称，支持中国大陆省份、港澳台以及海外国家。
  // 可选值：
  // anhui: 安徽
  // beijing: 北京
  // chongqing: 重庆
  // fujian: 福建
  // gansu: 甘肃
  // guangdong: 广东
  // guangxi: 广西
  // guizhou: 贵州
  // hainan: 海南
  // hebei: 河北
  // heilongjiang: 黑龙江
  // henan: 河南
  // hubei: 湖北
  // hunan: 湖南
  // jiangsu: 江苏
  // jiangxi: 江西
  // jilin: 吉林
  // liaoning: 辽宁
  // neimenggu: 内蒙古
  // ningxia: 宁夏
  // qinghai: 青海
  // shaanxi: 陕西
  // shandong: 山东
  // shanghai: 上海
  // shanxi: 山西
  // sichuan: 四川
  // tianjin: 天津
  // xinjiang: 新疆
  // xizang: 西藏
  // yunnan: 云南
  // zhejiang: 浙江
  // TW: 台湾
  // HK: 香港
  // MO: 澳门
  // AE: 阿联酋
  // AU: 澳大利亚
  // BD: 孟加拉
  // BN: 文莱
  // BR: 巴西
  // CA: 加拿大
  // CL: 智利
  // CO: 哥伦比亚
  // DJ: 吉布提
  // ID: 印度尼西亚
  // IT: 意大利
  // JP: 日本
  // KG: 吉尔吉斯斯坦
  // KH: 柬埔寨
  // KR: 韩国
  // KW: 科威特
  // LA: 老挝
  // MG: 马达加斯加
  // MM: 缅甸
  // MU: 毛里求斯
  // MY: 马来西亚
  // NP: 尼泊尔
  // OM: 阿曼
  // PE: 秘鲁
  // PH: 菲律宾
  // PK: 巴基斯坦
  // QA: 卡塔尔
  // RO: 罗马尼亚
  // RU: 俄罗斯
  // SA: 沙特阿拉伯
  // SE: 瑞典
  // SG: 新加坡
  // TH: 泰国
  // TR: 土耳其
  // US: 美国
  // VN: 越南"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"监控机所属运营商中文名。
  // 可选值：
  // 0: 中国电信
  // 1: 中国联通
  // 2: 中国铁通
  // 4: 中国移动
  // 5: 中国教育网
  // 9: 中国广电
  // 10: 长城宽带", "zh_CN":"监控机所属运营商中文名。
  // 可选值：
  // 0: 中国电信
  // 1: 中国联通
  // 2: 中国铁通
  // 4: 中国移动
  // 5: 中国教育网
  // 9: 中国广电
  // 10: 长城宽带"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s MtrTestRequest) String() string {
  return tea.Prettify(s)
}

func (s MtrTestRequest) GoString() string {
  return s.String()
}

func (s *MtrTestRequest) SetHost(v string) *MtrTestRequest {
  s.Host = &v
  return s
}

func (s *MtrTestRequest) SetArea(v string) *MtrTestRequest {
  s.Area = &v
  return s
}

func (s *MtrTestRequest) SetIsp(v string) *MtrTestRequest {
  s.Isp = &v
  return s
}

type MtrTestResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result *MtrTestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s MtrTestResponse) String() string {
  return tea.Prettify(s)
}

func (s MtrTestResponse) GoString() string {
  return s.String()
}

func (s *MtrTestResponse) SetResult(v *MtrTestResponseResult) *MtrTestResponse {
  s.Result = v
  return s
}

type MtrTestResponseResult struct {
  // {'en':'', 'zh_CN':''}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'', 'zh_CN':''}
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'目标域名'}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {'en':'data','zh_CN':'探测数据'}
  Data []*MtrTestResponseResultData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s MtrTestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s MtrTestResponseResult) GoString() string {
  return s.String()
}

func (s *MtrTestResponseResult) SetStatus(v string) *MtrTestResponseResult {
  s.Status = &v
  return s
}

func (s *MtrTestResponseResult) SetErrorMsg(v string) *MtrTestResponseResult {
  s.ErrorMsg = &v
  return s
}

func (s *MtrTestResponseResult) SetHost(v string) *MtrTestResponseResult {
  s.Host = &v
  return s
}

func (s *MtrTestResponseResult) SetData(v []*MtrTestResponseResultData) *MtrTestResponseResult {
  s.Data = v
  return s
}

type MtrTestResponseResultData struct     {
  // {'en':'', 'zh_CN':'任务 ID'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'任务 ID'}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'监控机 IP'}
  DetectIp *string `json:"detectIp,omitempty" xml:"detectIp,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'监控机运营商'}
  DetectIpIsp *string `json:"detectIpIsp,omitempty" xml:"detectIpIsp,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'监控机运营商编码'}
  DetectIpIspCode *string `json:"detectIpIspCode,omitempty" xml:"detectIpIspCode,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'监控机所属省份'}
  DetectIpPro *string `json:"detectIpPro,omitempty" xml:"detectIpPro,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'监控机所属省份编码'}
  DetectIpProCode *string `json:"detectIpProCode,omitempty" xml:"detectIpProCode,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'探测结果'}
  DetectResult *string `json:"detectResult,omitempty" xml:"detectResult,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'探测结束时间戳'}
  DoneTime *int64 `json:"doneTime,omitempty" xml:"doneTime,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'探测状态'}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'','zh_CN':'每跳记录'}
  Items []*MtrTestResponseResultDataItems `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s MtrTestResponseResultData) String() string {
  return tea.Prettify(s)
}

func (s MtrTestResponseResultData) GoString() string {
  return s.String()
}

func (s *MtrTestResponseResultData) SetId(v string) *MtrTestResponseResultData {
  s.Id = &v
  return s
}

func (s *MtrTestResponseResultData) SetTaskId(v string) *MtrTestResponseResultData {
  s.TaskId = &v
  return s
}

func (s *MtrTestResponseResultData) SetDetectIp(v string) *MtrTestResponseResultData {
  s.DetectIp = &v
  return s
}

func (s *MtrTestResponseResultData) SetDetectIpIsp(v string) *MtrTestResponseResultData {
  s.DetectIpIsp = &v
  return s
}

func (s *MtrTestResponseResultData) SetDetectIpIspCode(v string) *MtrTestResponseResultData {
  s.DetectIpIspCode = &v
  return s
}

func (s *MtrTestResponseResultData) SetDetectIpPro(v string) *MtrTestResponseResultData {
  s.DetectIpPro = &v
  return s
}

func (s *MtrTestResponseResultData) SetDetectIpProCode(v string) *MtrTestResponseResultData {
  s.DetectIpProCode = &v
  return s
}

func (s *MtrTestResponseResultData) SetDetectResult(v string) *MtrTestResponseResultData {
  s.DetectResult = &v
  return s
}

func (s *MtrTestResponseResultData) SetDoneTime(v int64) *MtrTestResponseResultData {
  s.DoneTime = &v
  return s
}

func (s *MtrTestResponseResultData) SetStatus(v int32) *MtrTestResponseResultData {
  s.Status = &v
  return s
}

func (s *MtrTestResponseResultData) SetItems(v []*MtrTestResponseResultDataItems) *MtrTestResponseResultData {
  s.Items = v
  return s
}

type MtrTestResponseResultDataItems struct     {
  // {'en':'','zh_CN':'序号'}
  Num *int32 `json:"num,omitempty" xml:"num,omitempty" require:"true"`
  // {'en':'','zh_CN':'IP'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {'en':'','zh_CN':'所属地区运营商'}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {'en':'The average round-trip time of all traceroute probes, in milliseconds.','zh_CN':'平均时延'}
  Avg *string `json:"avg,omitempty" xml:"avg,omitempty" require:"true"`
  // {'en':'The shortest round-trip time of all traceroute probes, in milliseconds.','zh_CN':'最优时延'}
  Best *string `json:"best,omitempty" xml:"best,omitempty" require:"true"`
  // {'en':'The longest round-trip time of all traceroute probes, in milliseconds.','zh_CN':'最差时延'}
  Wrst *string `json:"wrst,omitempty" xml:"wrst,omitempty" require:"true"`
  // {'en':'','zh_CN':'最近一次时延'}
  Last *string `json:"last,omitempty" xml:"last,omitempty" require:"true"`
  // {'en':'The percentage of packets for which an ICMP reply was not received.','zh_CN':'丢包率'}
  Loss *string `json:"loss,omitempty" xml:"loss,omitempty" require:"true"`
  // {'en':'The number of packets sent to each hop.','zh_CN':'每跳发送的数据包数'}
  Snt *string `json:"snt,omitempty" xml:"snt,omitempty" require:"true"`
}

func (s MtrTestResponseResultDataItems) String() string {
  return tea.Prettify(s)
}

func (s MtrTestResponseResultDataItems) GoString() string {
  return s.String()
}

func (s *MtrTestResponseResultDataItems) SetNum(v int32) *MtrTestResponseResultDataItems {
  s.Num = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetIp(v string) *MtrTestResponseResultDataItems {
  s.Ip = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetLocation(v string) *MtrTestResponseResultDataItems {
  s.Location = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetAvg(v string) *MtrTestResponseResultDataItems {
  s.Avg = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetBest(v string) *MtrTestResponseResultDataItems {
  s.Best = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetWrst(v string) *MtrTestResponseResultDataItems {
  s.Wrst = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetLast(v string) *MtrTestResponseResultDataItems {
  s.Last = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetLoss(v string) *MtrTestResponseResultDataItems {
  s.Loss = &v
  return s
}

func (s *MtrTestResponseResultDataItems) SetSnt(v string) *MtrTestResponseResultDataItems {
  s.Snt = &v
  return s
}

type MtrTestPaths struct {
}

func (s MtrTestPaths) String() string {
  return tea.Prettify(s)
}

func (s MtrTestPaths) GoString() string {
  return s.String()
}

type MtrTestParameters struct {
}

func (s MtrTestParameters) String() string {
  return tea.Prettify(s)
}

func (s MtrTestParameters) GoString() string {
  return s.String()
}

type MtrTestRequestHeader struct {
}

func (s MtrTestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s MtrTestRequestHeader) GoString() string {
  return s.String()
}

type MtrTestResponseHeader struct {
}

func (s MtrTestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s MtrTestResponseHeader) GoString() string {
  return s.String()
}




type GetCdnOriginIpRequestRequest struct {
  // {"en":"Start Time:\n\n1.The Time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n\n2.Cannot exceed the current time\n\n3.Up to the past 183 days of data can be obtained","zh_CN":"开始时间：\n\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n\n2.不能大于当前时间\n\n3.最多可获取最近183天的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:\n\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n\n2. The end time is greater than the start time\n\n3. If the end time is greater than the current time, the current time is taken\n\n4. Maximum query interval allowed: 1 day, that is, the difference between dateFrom and dateTo can not exceed 1 day","zh_CN":"结束时间：  \n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.结束时间需大于开始时间   \n3. 结束时间需大于开始时间，结束时间如果大于当前时间,取当前时间。\n4.允许查询最大间隔：1天，即dateFrom和dateTo相差不能超过1天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Rounds up or down a timestamp by a given time period. Only valid when the granularity is less than 1d(not including 1day)\n1. up-rounds up, e.g.: granularity=5m, 00:00:00-00:04:49 will be displayed as 00:05:00\n2. down-rounds down, e.g.: granularity=5m, 00:00:00-00:04:49 will be displayed as 00:00:00\n3. If not specified, the result will be rounded up (up).","zh_CN":"根据指定的时间周期对时间戳进行向上或向下取整。仅当粒度小于1天（不包含1天）时有效。 1. up – 向上取整。例如：granularity=5m，00:00:00-00:04:49 将会显示为 00:05:00 2. down-向下取整，例如：granularity=5m时，00:00:00-00:04:49将显示为00:00:00 3. 如未传值，结果将进行向上取整(up)"}
  TimeRounding *string `json:"timeRounding,omitempty" xml:"timeRounding,omitempty"`
  // {"en":"Domain Name: The maximum number of domain names that can be submitted is 20 by default.","zh_CN":"域名：可传递域名数量上限默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Origin IP: The maximum number of origin IPs that can be submitted is 50 by default.","zh_CN":"源站IP：可传递源站IP数量上限默认为50个"}
  OriginIp []*string `json:"originIp,omitempty" xml:"originIp,omitempty" type:"Repeated"`
}

func (s GetCdnOriginIpRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestRequest) GoString() string {
  return s.String()
}

func (s *GetCdnOriginIpRequestRequest) SetDateFrom(v string) *GetCdnOriginIpRequestRequest {
  s.DateFrom = &v
  return s
}

func (s *GetCdnOriginIpRequestRequest) SetDateTo(v string) *GetCdnOriginIpRequestRequest {
  s.DateTo = &v
  return s
}

func (s *GetCdnOriginIpRequestRequest) SetTimeRounding(v string) *GetCdnOriginIpRequestRequest {
  s.TimeRounding = &v
  return s
}

func (s *GetCdnOriginIpRequestRequest) SetDomain(v []*string) *GetCdnOriginIpRequestRequest {
  s.Domain = v
  return s
}

func (s *GetCdnOriginIpRequestRequest) SetOriginIp(v []*string) *GetCdnOriginIpRequestRequest {
  s.OriginIp = v
  return s
}

type GetCdnOriginIpRequestRequestHeader struct {
}

func (s GetCdnOriginIpRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestRequestHeader) GoString() string {
  return s.String()
}

type GetCdnOriginIpRequestPaths struct {
}

func (s GetCdnOriginIpRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestPaths) GoString() string {
  return s.String()
}

type GetCdnOriginIpRequestParameters struct {
}

func (s GetCdnOriginIpRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestParameters) GoString() string {
  return s.String()
}

type GetCdnOriginIpRequestResponse struct {
  // {"en":"Status Code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result","zh_CN":"结果"}
  Data []*GetCdnOriginIpRequestResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnOriginIpRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestResponse) GoString() string {
  return s.String()
}

func (s *GetCdnOriginIpRequestResponse) SetCode(v string) *GetCdnOriginIpRequestResponse {
  s.Code = &v
  return s
}

func (s *GetCdnOriginIpRequestResponse) SetMessage(v string) *GetCdnOriginIpRequestResponse {
  s.Message = &v
  return s
}

func (s *GetCdnOriginIpRequestResponse) SetData(v []*GetCdnOriginIpRequestResponseData) *GetCdnOriginIpRequestResponse {
  s.Data = v
  return s
}

type GetCdnOriginIpRequestResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Results per minute","zh_CN":"每一分钟的结果数据"}
  PerMinuteStats []*GetCdnOriginIpRequestResponseDataPerMinuteStats `json:"perMinuteStats,omitempty" xml:"perMinuteStats,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnOriginIpRequestResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestResponseData) GoString() string {
  return s.String()
}

func (s *GetCdnOriginIpRequestResponseData) SetDomain(v string) *GetCdnOriginIpRequestResponseData {
  s.Domain = &v
  return s
}

func (s *GetCdnOriginIpRequestResponseData) SetPerMinuteStats(v []*GetCdnOriginIpRequestResponseDataPerMinuteStats) *GetCdnOriginIpRequestResponseData {
  s.PerMinuteStats = v
  return s
}

type GetCdnOriginIpRequestResponseDataPerMinuteStats struct     {
  // {"en":"Time","zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"List of origin IPs","zh_CN":"源站ip列表"}
  OriginIp []*GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp `json:"originIp,omitempty" xml:"originIp,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnOriginIpRequestResponseDataPerMinuteStats) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestResponseDataPerMinuteStats) GoString() string {
  return s.String()
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStats) SetTimestamp(v string) *GetCdnOriginIpRequestResponseDataPerMinuteStats {
  s.Timestamp = &v
  return s
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStats) SetOriginIp(v []*GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp) *GetCdnOriginIpRequestResponseDataPerMinuteStats {
  s.OriginIp = v
  return s
}

type GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp struct     {
  // {"en":"Origin IP","zh_CN":"源站ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Total number of requests under this IP","zh_CN":"该ip下的总请求数"}
  TotalRequests *int `json:"totalRequests,omitempty" xml:"totalRequests,omitempty" require:"true"`
  // {"en":"List of status codes","zh_CN":"状态码列表"}
  StatusCodes []*GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes `json:"statusCodes,omitempty" xml:"statusCodes,omitempty" require:"true" type:"Repeated"`
}

func (s GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp) GoString() string {
  return s.String()
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp) SetIp(v string) *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp {
  s.Ip = &v
  return s
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp) SetTotalRequests(v int) *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp {
  s.TotalRequests = &v
  return s
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp) SetStatusCodes(v []*GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes) *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIp {
  s.StatusCodes = v
  return s
}

type GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes struct     {
  // {"en":"Status code, summary status code: provide summarized status code, 0, 1XX-6XX, others categorized as others","zh_CN":"状态码，汇总状态码:提供汇总后的状态码，0,1XX-6XX,其他归类到others"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Number of requests","zh_CN":"请求数"}
  Requests *int `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
}

func (s GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes) GoString() string {
  return s.String()
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes) SetStatusCode(v string) *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes {
  s.StatusCode = &v
  return s
}

func (s *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes) SetRequests(v int) *GetCdnOriginIpRequestResponseDataPerMinuteStatsOriginIpStatusCodes {
  s.Requests = &v
  return s
}

type GetCdnOriginIpRequestResponseHeader struct {
}

func (s GetCdnOriginIpRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCdnOriginIpRequestResponseHeader) GoString() string {
  return s.String()
}




type ReportAvgSpeedDomainIspProvinceServiceRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  //         1. 时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  //         2. 不能大于当前时间
  //         3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 1 hour; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 1 hour (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 1 hour.", "zh_CN":"结束时间:
  //         1. 时间格式2016-12-02T10:00:00+08:00
  //         2. 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  //         3. dateFrom,dateTo二者都未传,默认查询过去的1小时;如仅有一个未传,抛异常
  //         4. 允许查询最大时间间隔:1小时,即dateFrom和dateTo相差不能超过1小时(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of domains that can be delivered is 20 by default (contact technical support adjustment);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3. Domain name exceeding limit, misstatement", "zh_CN":"域名:
  //         1. 可传递域名数量上限默认为20个(可联系技术支持调整)。未传递该入参时查询账号下所有域名,但当账号下域名数量超过限制时不可查询(报错)。
  //         2. 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. default 5m(only support 5 minutes)", "zh_CN":"数据粒度,默认为5m(5分钟,当前只支持5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province chinese name:
  // 1. Query all provinces by default without transferring;
  // 2. Please refer to the description of the province list on the overview page for the values.", "zh_CN":"省份中文名称:
  //         1. 不传默认查询全部省份;
  //         2. 可传递的值详见概览页省份列表说明。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP chinese name:
  // 1. Query all ISP by default without transferring;
  // 2. Please refer to the description of the ISP list on the overview page for the values.", "zh_CN":"运营商中文名称:
  //         1. 不传默认查询全部运营商;
  //         2. 可传递的值详见概览页运营商列表说明。"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension:
  // 1. Optional values:domain,province,isp,you can pass in single or multiple values  
  // 2. The detailed data will be displayed according to the dimension.
  //        ", "zh_CN":"分组维度:
  //         1. 可选值为domain、province、isp,可传入单个或多个值;
  //         2. 有传入则按照该维度展示明细数据;
  //         3. 返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:'groupBy': ['domain','province']与'groupBy': ['province','domain']返回结果一样。
  //         例如:传递'groupBy': ['domain','province'],则ispData下的isp节点无需返回。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetDateFrom(v string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetDateTo(v string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetDomain(v []*string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetDataInterval(v string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetProvince(v []*string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetIsp(v []*string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportAvgSpeedDomainIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportAvgSpeedDomainIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponse) SetResult(v []*ReportAvgSpeedDomainIspProvinceServiceResponseResult) *ReportAvgSpeedDomainIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ispData", "zh_CN":"ISP数据"}
  IspData []*ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResult) SetDomain(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResult) SetIspData(v []*ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) *ReportAvgSpeedDomainIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"provinceData", "zh_CN":"省份数据"}
  ProvinceData []*ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"ISP chinese name", "zh_CN":"省份中文名称"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"details", "zh_CN":"详情数据"}
  Details []*ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) SetDetails(v []*ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData {
  s.Details = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails struct     {
  // {"en":"Time:
  //         1. When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  //         2. Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间
  //         1. 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。
  //         2. 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Average download speed: the average download speed of all requests that meet the statistical rules in the time slice.
  // 1. unit: KB/s;
  // 2. Two decimal places are reserved.", "zh_CN":"平均下载速度:时间片(timestamp)内符合统计规则的所有请求的平均下载速度。
  //         1. 计量单位:KB/s;
  //         2. 保留两位小数。"}
  AvgSpeed *string `json:"avgSpeed,omitempty" xml:"avgSpeed,omitempty" require:"true"`
  // {"en":"Average response time: the average response time of a single request in a time stamp. Unit : millisecond.", "zh_CN":"平均响应时间:时间片(timestamp)内单条请求平均响应时间。计量单位:毫秒。"}
  AvgResponseTime *string `json:"avgResponseTime,omitempty" xml:"avgResponseTime,omitempty" require:"true"`
  // {"en":"Average first packet response time: the average first packet response time of a single request in a time slice. Unit: millisecond.", "zh_CN":"平均首包响应时间:时间片(timestamp)内单条请求平均首包响应时间。计量单位:毫秒。"}
  AvgFirstPacketTime *string `json:"avgFirstPacketTime,omitempty" xml:"avgFirstPacketTime,omitempty" require:"true"`
  // {"en":"Total response time: the response time summary of all requests that meet the statistical rules in the time slice. Unit: millisecond.", "zh_CN":"总响应时间:时间片(timestamp)内符合统计规则的所有请求响应时间汇总。计量单位:毫秒。"}
  TotalResponseTime *string `json:"totalResponseTime,omitempty" xml:"totalResponseTime,omitempty" require:"true"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetTimestamp(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetAvgSpeed(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.AvgSpeed = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetAvgResponseTime(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.AvgResponseTime = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetAvgFirstPacketTime(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.AvgFirstPacketTime = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetTotalResponseTime(v string) *ReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.TotalResponseTime = &v
  return s
}

type ReportAvgSpeedDomainIspProvinceServicePaths struct {
}

func (s ReportAvgSpeedDomainIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportAvgSpeedDomainIspProvinceServiceParameters struct {
}

func (s ReportAvgSpeedDomainIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportAvgSpeedDomainIspProvinceServiceRequestHeader struct {
}

func (s ReportAvgSpeedDomainIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportAvgSpeedDomainIspProvinceServiceResponseHeader struct {
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type GetLiveStreamPushingStatusRequest struct {
}

func (s GetLiveStreamPushingStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusRequest) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusRequestHeader struct {
}

func (s GetLiveStreamPushingStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusRequestHeader) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusPaths struct {
}

func (s GetLiveStreamPushingStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusPaths) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusParameters struct {
  // {"en":"Push domain(multiple domains supported, separated by commas)(    All parameters are passed via HTTP GET requests.)","zh_CN":"推流域名（所有参数以HTTP GET请求方式传参)"}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {"en":"Time, eg: 20160527152300, if not filled in, the current time -5 minutes","zh_CN":"时间，eg：20160527152300，不填时为当前时间-5分钟"}
  T *int `json:"t,omitempty" xml:"t,omitempty"`
  // {"en":"Push channel URL (multiple push channel URLs are supported, separated by commas)","zh_CN":"推流流名(支持多个推流流名，以逗号分隔)"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Query interval. The value can be 10 or 60. The default value is 60\nWhen g is 10, query the data of the nearest whole 10 seconds to time t\nWhen g is 60, query the data of the nearest whole minute to time t","zh_CN":"查询间隔，可选值10、60，默认60\n当g为10时，查询距离时间t最近的整10秒点数据\n当g为60时，查询距离时间t最近的整分钟点数据"}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
}

func (s GetLiveStreamPushingStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusParameters) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusParameters) SetU(v string) *GetLiveStreamPushingStatusParameters {
  s.U = &v
  return s
}

func (s *GetLiveStreamPushingStatusParameters) SetT(v int) *GetLiveStreamPushingStatusParameters {
  s.T = &v
  return s
}

func (s *GetLiveStreamPushingStatusParameters) SetChannel(v string) *GetLiveStreamPushingStatusParameters {
  s.Channel = &v
  return s
}

func (s *GetLiveStreamPushingStatusParameters) SetG(v string) *GetLiveStreamPushingStatusParameters {
  s.G = &v
  return s
}

type GetLiveStreamPushingStatusResponse struct {
  // {"en":"The time of the data returned","zh_CN":"返回的数据的时间"}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {"en":"Number of data items. 0 is returned if there is no data","zh_CN":"数据条数，无数据返回0"}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {"en":"","zh_CN":"推流信息数据集合"}
  DataValue []*GetLiveStreamPushingStatusResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
}

func (s GetLiveStreamPushingStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponse) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusResponse) SetRettime(v string) *GetLiveStreamPushingStatusResponse {
  s.Rettime = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponse) SetRetcode(v int64) *GetLiveStreamPushingStatusResponse {
  s.Retcode = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponse) SetDataValue(v []*GetLiveStreamPushingStatusResponseDataValue) *GetLiveStreamPushingStatusResponse {
  s.DataValue = v
  return s
}

type GetLiveStreamPushingStatusResponseDataValue struct     {
  // {"en":"Push stream CDN node, that is, the IP address of edge node of the received data source, separated by multiple commas","zh_CN":"推流cdn节点，即收到数据源的edge节点IP地址，多个逗号分隔"}
  Deployaddress *string `json:"deployaddress,omitempty" xml:"deployaddress,omitempty" require:"true"`
  // {"en":"Push user IP (data source) address, separated by multiple commas","zh_CN":"推流用户ip（数据源）地址，多个逗号分隔"}
  Inaddress *string `json:"inaddress,omitempty" xml:"inaddress,omitempty" require:"true"`
  // {"en":"The channel of anchor","zh_CN":"主播流名"}
  Streamname *string `json:"streamname,omitempty" xml:"streamname,omitempty" require:"true"`
  // {"en":"Anchor Current encoding frame rate","zh_CN":"主播当前编码帧率"}
  Fps *int64 `json:"fps,omitempty" xml:"fps,omitempty" require:"true"`
  // {"en":"Current frame loss rate of anchor","zh_CN":"主播当前丢帧率"}
  Lfr *GetLiveStreamPushingStatusResponseDataValueLfr `json:"lfr,omitempty" xml:"lfr,omitempty" require:"true" type:"Struct"`
  // {"en":"Anchor Current bit rate","zh_CN":"主播当前码率"}
  Inbandwidth *int64 `json:"inbandwidth,omitempty" xml:"inbandwidth,omitempty" require:"true"`
  // {"en":"Video timestamps are separated by multiple commas","zh_CN":"视频时间戳 多个逗号分隔"}
  Videotmstmp *string `json:"videotmstmp,omitempty" xml:"videotmstmp,omitempty" require:"true"`
  // {"en":"Audio timestamp, separated by multiple commas","zh_CN":"音频时间戳，多个逗号分隔"}
  Audiotmstmp *string `json:"audiotmstmp,omitempty" xml:"audiotmstmp,omitempty" require:"true"`
}

func (s GetLiveStreamPushingStatusResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponseDataValue) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetDeployaddress(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Deployaddress = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetInaddress(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Inaddress = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetStreamname(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Streamname = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetFps(v int64) *GetLiveStreamPushingStatusResponseDataValue {
  s.Fps = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetLfr(v *GetLiveStreamPushingStatusResponseDataValueLfr) *GetLiveStreamPushingStatusResponseDataValue {
  s.Lfr = v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetInbandwidth(v int64) *GetLiveStreamPushingStatusResponseDataValue {
  s.Inbandwidth = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetVideotmstmp(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Videotmstmp = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetAudiotmstmp(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Audiotmstmp = &v
  return s
}

type GetLiveStreamPushingStatusResponseDataValueLfr struct {
}

func (s GetLiveStreamPushingStatusResponseDataValueLfr) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponseDataValueLfr) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusResponseHeader struct {
}

func (s GetLiveStreamPushingStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponseHeader) GoString() string {
  return s.String()
}




type ReportOnlineNumIspProvinceServiceRequest struct {
  // {"en":"Start time:
  // 	1.The format is yyyy-MM-ddTHH:mm:ss+08:00; 
  // 	2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo; 
  // 	3.Period between dataFrom and dateTo cannot be longer than 1 hour; 
  // 	4.dateFrom and dateTo can be either both are specified or neither is specifies; 
  // 	5.If neither dateFrom nor dateTo is specified, then by default, data in the last 30 minutes is queried.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过1小时(可联系技术支持调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去30分钟的数据;"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00.
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  // 	1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support).
  // 	2.Domain is not uploaded: Query all domain names of the account", "zh_CN":"域名:
  // 可传递域名数量上限默认为20个(可联系技术支持调整);
  // 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 域名超过上限,报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"data granularity:
  // 1m: 1 minute
  // 5m: 5 minutes
  // Do not pass default query 5m", "zh_CN":"数据粒度:
  // 1m: 1分钟粒度
  // 5m: 5分钟粒度
  // 不传默认查询 5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.", "zh_CN":"省份
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商：
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group by:
  // 	1.Optional value: domain,province,isp; Multiple values can be selected.
  // 	2.If there is an input, the detailed data will be displayed according to this dimension ", "zh_CN":"分组维度:
  // 可选值为domain,province,isp,可传入多个值;
  // 有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetDateFrom(v string) *ReportOnlineNumIspProvinceServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetDateTo(v string) *ReportOnlineNumIspProvinceServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetDomain(v []*string) *ReportOnlineNumIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetDataInterval(v string) *ReportOnlineNumIspProvinceServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetProvince(v []*string) *ReportOnlineNumIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetIsp(v []*string) *ReportOnlineNumIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportOnlineNumIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportOnlineNumIspProvinceServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportOnlineNumIspProvinceServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceResponse) SetCode(v string) *ReportOnlineNumIspProvinceServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceResponse) SetMessage(v string) *ReportOnlineNumIspProvinceServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceResponse) SetData(v []*ReportOnlineNumIspProvinceServiceResponseData) *ReportOnlineNumIspProvinceServiceResponse {
  s.Data = v
  return s
}

type ReportOnlineNumIspProvinceServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportOnlineNumIspProvinceServiceResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceResponseData) SetDomain(v string) *ReportOnlineNumIspProvinceServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceResponseData) SetIspData(v []*ReportOnlineNumIspProvinceServiceResponseDataIspData) *ReportOnlineNumIspProvinceServiceResponseData {
  s.IspData = v
  return s
}

type ReportOnlineNumIspProvinceServiceResponseDataIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceResponseDataIspData) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceResponseDataIspData) SetIsp(v string) *ReportOnlineNumIspProvinceServiceResponseDataIspData {
  s.Isp = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceResponseDataIspData) SetProvinceData(v []*ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) *ReportOnlineNumIspProvinceServiceResponseDataIspData {
  s.ProvinceData = v
  return s
}

type ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  Details []*ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) SetProvince(v string) *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) SetDetails(v []*ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData {
  s.Details = v
  return s
}

type ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails struct     {
  // {"en":"Time, the format is yyyy-MM-dd HH:mm", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"online users", "zh_CN":"在线人数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetTimestamp(v string) *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetValue(v string) *ReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Value = &v
  return s
}

type ReportOnlineNumIspProvinceServicePaths struct {
}

func (s ReportOnlineNumIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportOnlineNumIspProvinceServiceParameters struct {
}

func (s ReportOnlineNumIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportOnlineNumIspProvinceServiceRequestHeader struct {
}

func (s ReportOnlineNumIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportOnlineNumIspProvinceServiceResponseHeader struct {
}

func (s ReportOnlineNumIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type ConcurrentSessionRequest struct {
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not specified,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope.\n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope.\n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1)If there are multiple inputs,use  ';' as separator.\n2)If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"dictionary":"belong=BCS-CC-API|dict=flowRegionCode","en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.\n1)If there are multiple inputs,use ';' as separator.\n2)If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Greenwich Mean Time zone. The format GMT+09:00 represents UTC+9 (East 9), while GMT-09:00 represents UTC-9 (West 9). If not specified, the default is the local time zone (UTC+8).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Data granularity: 1-minute granularity by default\n5m: 5-minute granularity\n1m: 1-minute granularity","zh_CN":"数据粒度：默认1分钟粒度\n5m: 5分钟粒度\n1m: 1分钟粒度","exampleValue":"1m,5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Group by keyword:  \n1. If not provided, the results will be aggregated and displayed by default.\n2. If a keyword is provided, the details will be displayed based on the dimension corresponding to that keyword (for example, if 'channel' is provided, the details will be returned and expanded by domain). If an unsupported keyword is provided, an error message \"invalid groupby\" will be returned.\n3. Multiple values are supported; separate multiple values with an English semicolon ';'. Currently, only 'channel' is supported.","zh_CN":"分组关键词：  \n1.未传递时，默认聚合展示\n2.传入关键词则代表需要按照关键词维度分组对应的值展示明细（例如传channel，则代表返回按照domain明细展开。），如传入不支持的关键词，返回错误提示\"invalid groupby\"\n3.支持传多个值，多个值以英文分号';'分隔，当前只支持channel","exampleValue":"channel"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s ConcurrentSessionRequest) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionRequest) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionRequest) SetDate(v string) *ConcurrentSessionRequest {
  s.Date = &v
  return s
}

func (s *ConcurrentSessionRequest) SetStartdate(v string) *ConcurrentSessionRequest {
  s.Startdate = &v
  return s
}

func (s *ConcurrentSessionRequest) SetEnddate(v string) *ConcurrentSessionRequest {
  s.Enddate = &v
  return s
}

func (s *ConcurrentSessionRequest) SetChannel(v string) *ConcurrentSessionRequest {
  s.Channel = &v
  return s
}

func (s *ConcurrentSessionRequest) SetRegion(v string) *ConcurrentSessionRequest {
  s.Region = &v
  return s
}

func (s *ConcurrentSessionRequest) SetAccetype(v string) *ConcurrentSessionRequest {
  s.Accetype = &v
  return s
}

func (s *ConcurrentSessionRequest) SetDataformat(v string) *ConcurrentSessionRequest {
  s.Dataformat = &v
  return s
}

func (s *ConcurrentSessionRequest) SetTimezone(v string) *ConcurrentSessionRequest {
  s.Timezone = &v
  return s
}

func (s *ConcurrentSessionRequest) SetGranularity(v string) *ConcurrentSessionRequest {
  s.Granularity = &v
  return s
}

func (s *ConcurrentSessionRequest) SetGroupBy(v string) *ConcurrentSessionRequest {
  s.GroupBy = &v
  return s
}

type ConcurrentSessionRequestHeader struct {
}

func (s ConcurrentSessionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionRequestHeader) GoString() string {
  return s.String()
}

type ConcurrentSessionPaths struct {
}

func (s ConcurrentSessionPaths) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionPaths) GoString() string {
  return s.String()
}

type ConcurrentSessionParameters struct {
}

func (s ConcurrentSessionParameters) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionParameters) GoString() string {
  return s.String()
}

type ConcurrentSessionResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *ConcurrentSessionResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s ConcurrentSessionResponse) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionResponse) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionResponse) SetProvider(v *ConcurrentSessionResponseProvider) *ConcurrentSessionResponse {
  s.Provider = v
  return s
}

type ConcurrentSessionResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"明细数据"}
  Date *ConcurrentSessionResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s ConcurrentSessionResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionResponseProvider) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionResponseProvider) SetName(v string) *ConcurrentSessionResponseProvider {
  s.Name = &v
  return s
}

func (s *ConcurrentSessionResponseProvider) SetType(v string) *ConcurrentSessionResponseProvider {
  s.Type = &v
  return s
}

func (s *ConcurrentSessionResponseProvider) SetDate(v *ConcurrentSessionResponseProviderDate) *ConcurrentSessionResponseProvider {
  s.Date = v
  return s
}

type ConcurrentSessionResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"concurrent","zh_CN":"明细数据"}
  Concurrent []*ConcurrentSessionResponseProviderDateConcurrent `json:"concurrent,omitempty" xml:"concurrent,omitempty" require:"true" type:"Repeated"`
}

func (s ConcurrentSessionResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionResponseProviderDate) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionResponseProviderDate) SetStartdate(v string) *ConcurrentSessionResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *ConcurrentSessionResponseProviderDate) SetEnddate(v string) *ConcurrentSessionResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *ConcurrentSessionResponseProviderDate) SetConcurrent(v []*ConcurrentSessionResponseProviderDateConcurrent) *ConcurrentSessionResponseProviderDate {
  s.Concurrent = v
  return s
}

type ConcurrentSessionResponseProviderDateConcurrent struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"hit count","zh_CN":"明细数据"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s ConcurrentSessionResponseProviderDateConcurrent) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionResponseProviderDateConcurrent) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionResponseProviderDateConcurrent) SetTime(v string) *ConcurrentSessionResponseProviderDateConcurrent {
  s.Time = &v
  return s
}

func (s *ConcurrentSessionResponseProviderDateConcurrent) SetText(v string) *ConcurrentSessionResponseProviderDateConcurrent {
  s.Text = &v
  return s
}

type ConcurrentSessionResponseHeader struct {
}

func (s ConcurrentSessionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionResponseHeader) GoString() string {
  return s.String()
}




type BandwidthAppaRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"为需要查询带宽数据的起始日期,日期格式为yyyy-mm-dd ；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"为需要查询带宽数据的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"bandwidth type of appa.
  // 1)optional values:appaEdgeUp, appaEdgeDown, appaEdgeSend,appaInterimReceiver,appaInterimSend,appaOriginReceiver,appaSourceUpSae,appaSourceUpNode,appaSourceUpTotal,appaSoruceDownSae,appaSourceDownNod,appaSourceDownTotal.
  // 2)If there are multiple inputs,use ';'as separator.
  // 3)Four type at most at a  time.", "zh_CN":"appaEdgeUp(边缘上行), appaEdgeDown(边缘下行), appaEdgeSend(中转下行),appaInterimReceiver(中转上行),appaInterimSend(回源下行),appaOriginReceiver(回源上行),appaSourceUpSae(源站上行（SAE）),appaSourceUpNode(源站上行（节点）),appaSourceUpTotal(源站下行总带宽（SAE+节点）),appaSoruceDownSae(源站下行（SAE）),appaSourceDownNode(源站下行（节点）),appaSourceDownTotal(源站下行总带宽（SAE+节点）), 多种类型之间用英文分号进行分隔，最多同时只能选择四种类型,"}
  AppaBandwidthType *string `json:"appaBandwidthType,omitempty" xml:"appaBandwidthType,omitempty" require:"true"`
  // {"en":"Whether to return traffic details, 1: yes; 0: no. The default is 0.", "zh_CN":"是否需要返回流量明细，1：需要；0：不需要。默认为0."}
  NeedFlow *string `json:"needFlow,omitempty" xml:"needFlow,omitempty"`
}

func (s BandwidthAppaRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaRequest) GoString() string {
  return s.String()
}

func (s *BandwidthAppaRequest) SetCust(v string) *BandwidthAppaRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthAppaRequest) SetDate(v string) *BandwidthAppaRequest {
  s.Date = &v
  return s
}

func (s *BandwidthAppaRequest) SetStartdate(v string) *BandwidthAppaRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthAppaRequest) SetEnddate(v string) *BandwidthAppaRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthAppaRequest) SetChannel(v string) *BandwidthAppaRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthAppaRequest) SetTimezone(v string) *BandwidthAppaRequest {
  s.Timezone = &v
  return s
}

func (s *BandwidthAppaRequest) SetIsExactMatch(v string) *BandwidthAppaRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthAppaRequest) SetRegion(v string) *BandwidthAppaRequest {
  s.Region = &v
  return s
}

func (s *BandwidthAppaRequest) SetAccetype(v string) *BandwidthAppaRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthAppaRequest) SetDataformat(v string) *BandwidthAppaRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthAppaRequest) SetAppaBandwidthType(v string) *BandwidthAppaRequest {
  s.AppaBandwidthType = &v
  return s
}

func (s *BandwidthAppaRequest) SetNeedFlow(v string) *BandwidthAppaRequest {
  s.NeedFlow = &v
  return s
}

type BandwidthAppaResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthAppaResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthAppaResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponse) GoString() string {
  return s.String()
}

func (s *BandwidthAppaResponse) SetProvider(v *BandwidthAppaResponseProvider) *BandwidthAppaResponse {
  s.Provider = v
  return s
}

type BandwidthAppaResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'APPA带宽数据'}
  Date *BandwidthAppaResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthAppaResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthAppaResponseProvider) SetName(v string) *BandwidthAppaResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthAppaResponseProvider) SetType(v string) *BandwidthAppaResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthAppaResponseProvider) SetDate(v *BandwidthAppaResponseProviderDate) *BandwidthAppaResponseProvider {
  s.Date = v
  return s
}

type BandwidthAppaResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'appa', 'zh_CN':'appa带宽数据'}
  Appa []*BandwidthAppaResponseProviderDateAppa `json:"appa,omitempty" xml:"appa,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthAppaResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthAppaResponseProviderDate) SetStartdate(v string) *BandwidthAppaResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthAppaResponseProviderDate) SetEnddate(v string) *BandwidthAppaResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthAppaResponseProviderDate) SetAppa(v []*BandwidthAppaResponseProviderDateAppa) *BandwidthAppaResponseProviderDate {
  s.Appa = v
  return s
}

type BandwidthAppaResponseProviderDateAppa struct     {
  // {'en':'type', 'zh_CN':'appa数据类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道名称'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*BandwidthAppaResponseProviderDateAppaBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthAppaResponseProviderDateAppa) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponseProviderDateAppa) GoString() string {
  return s.String()
}

func (s *BandwidthAppaResponseProviderDateAppa) SetType(v string) *BandwidthAppaResponseProviderDateAppa {
  s.Type = &v
  return s
}

func (s *BandwidthAppaResponseProviderDateAppa) SetChannel(v string) *BandwidthAppaResponseProviderDateAppa {
  s.Channel = &v
  return s
}

func (s *BandwidthAppaResponseProviderDateAppa) SetBandwidth(v []*BandwidthAppaResponseProviderDateAppaBandwidth) *BandwidthAppaResponseProviderDateAppa {
  s.Bandwidth = v
  return s
}

type BandwidthAppaResponseProviderDateAppaBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点，格式 yyyy-MM-dd hh:mm:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽，单位Mbps'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthAppaResponseProviderDateAppaBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponseProviderDateAppaBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthAppaResponseProviderDateAppaBandwidth) SetTime(v string) *BandwidthAppaResponseProviderDateAppaBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthAppaResponseProviderDateAppaBandwidth) SetText(v string) *BandwidthAppaResponseProviderDateAppaBandwidth {
  s.Text = &v
  return s
}

type BandwidthAppaPaths struct {
}

func (s BandwidthAppaPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaPaths) GoString() string {
  return s.String()
}

type BandwidthAppaParameters struct {
}

func (s BandwidthAppaParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaParameters) GoString() string {
  return s.String()
}

type BandwidthAppaRequestHeader struct {
}

func (s BandwidthAppaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaRequestHeader) GoString() string {
  return s.String()
}

type BandwidthAppaResponseHeader struct {
}

func (s BandwidthAppaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainUserAgentServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days.", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error,you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration area:
  // Acceleration areaCode is not uploaded, query all acceleration areas by default.", "zh_CN":"加速区域:未传递areaCode时,默认查询所有加速区域。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {"en":"Query type:
  // 1.The optional values are browser, brand, os;The defalut value is browser.
  // 2.When the value is browser, query according to the browser.
  // 3.When the value is brand, query based on the mobile device brand.
  // 4.When the value is os, query according to the operating system.", "zh_CN":"查询类型
  // 1.可选值为browser,brand,os,不传默认按照browser查询;
  // 2.值为browser时,则根据浏览器查询;
  // 3.值为brand时,则根据移动设备品牌查询;
  // 4.值为os时,则根据操作系统查询。"}
  DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
  // {"en":"Order by:
  // 	1.The optional values are flow and request;
  // 	2.If not passed, the default is request;
  // 	3.When the value is flow, the query results are sorted in descending order of traffic, and when the value is request, they are sorted in descending order of the number of requests.", "zh_CN":"排序:
  // 1.可选值为flow、request;
  // 2.不传默认为request;
  // 3.值为flow时查询结果按流量降序排列,值为request时按请求数降序排列"}
  Orderby *string `json:"orderby,omitempty" xml:"orderby,omitempty"`
}

func (s ReportDomainUserAgentServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainUserAgentServiceRequest) SetDateFrom(v string) *ReportDomainUserAgentServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainUserAgentServiceRequest) SetDateTo(v string) *ReportDomainUserAgentServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainUserAgentServiceRequest) SetDomain(v []*string) *ReportDomainUserAgentServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportDomainUserAgentServiceRequest) SetAreaCode(v []*string) *ReportDomainUserAgentServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportDomainUserAgentServiceRequest) SetDeviceType(v string) *ReportDomainUserAgentServiceRequest {
  s.DeviceType = &v
  return s
}

func (s *ReportDomainUserAgentServiceRequest) SetOrderby(v string) *ReportDomainUserAgentServiceRequest {
  s.Orderby = &v
  return s
}

type ReportDomainUserAgentServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*ReportDomainUserAgentServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainUserAgentServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainUserAgentServiceResponse) SetCode(v string) *ReportDomainUserAgentServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportDomainUserAgentServiceResponse) SetMessage(v string) *ReportDomainUserAgentServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportDomainUserAgentServiceResponse) SetData(v []*ReportDomainUserAgentServiceResponseData) *ReportDomainUserAgentServiceResponse {
  s.Data = v
  return s
}

type ReportDomainUserAgentServiceResponseData struct     {
  // {"en":"UA", "zh_CN":"UA"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {"en":"Flow value. Unit is MB and 2 digits of decimals are allowed.", "zh_CN":"流量,保留2位小数,单位MB"}
  Flow *int `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {"en":"Number of requests.", "zh_CN":"请求数"}
  Request *int `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s ReportDomainUserAgentServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainUserAgentServiceResponseData) SetUserAgent(v string) *ReportDomainUserAgentServiceResponseData {
  s.UserAgent = &v
  return s
}

func (s *ReportDomainUserAgentServiceResponseData) SetFlow(v int) *ReportDomainUserAgentServiceResponseData {
  s.Flow = &v
  return s
}

func (s *ReportDomainUserAgentServiceResponseData) SetRequest(v int) *ReportDomainUserAgentServiceResponseData {
  s.Request = &v
  return s
}

type ReportDomainUserAgentServicePaths struct {
}

func (s ReportDomainUserAgentServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServicePaths) GoString() string {
  return s.String()
}

type ReportDomainUserAgentServiceParameters struct {
}

func (s ReportDomainUserAgentServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainUserAgentServiceRequestHeader struct {
}

func (s ReportDomainUserAgentServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainUserAgentServiceResponseHeader struct {
}

func (s ReportDomainUserAgentServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDailyLiveTranscodingDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"This parameter specifies if the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"Transcoding type, values can be h264, h265, zdgq_264, zdgq_265, cf_264, cf_265, or other. Multiple transcoding types should be separated by a semicolon. If some of the transcoding types are incorrect, the system will return data for the correct types; if all transcoding types are incorrect, it will return an error 'invalid transcodeType.' If not provided or left empty, it defaults to all types.", "zh_CN":"转码类型,值为h264、h265、zdgq_264、zdgq_265、cf_264、cf_265，other 多个转码类型用英文分号;分隔开。当传入转码类型部分错误时，返回正确的类型的数据；当传入转码类型全部错误时，返回错误invalid transcodeType. 不填或为空，默认为所有类型."}
  TranscodeType *string `json:"transcodeType,omitempty" xml:"transcodeType,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Resolution types include LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240. Multiple resolutions are separated by a semicolon. When isAudio=1, this parameter is invalid and will return an error. Param definition must be empty when querying audio data.", "zh_CN":"清晰度类型,值为LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240，多个清晰度用英文分号;分隔开, 当isAudio=1时，此入参无效返回错误 param definition must be empty when query audio data."}
  Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
  // {"en":"Audio/Video Type, 1: Audio 2: Video. Defaults to 2 if not selected or empty. Only a single value is allowed.", "zh_CN":"音视频类型, 1:音频   2:视频. 不选或者为空时默认为2. 只能输入单个值."}
  IsAudio *string `json:"isAudio,omitempty" xml:"isAudio,omitempty"`
}

func (s QueryDailyLiveTranscodingDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationRequest) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetCust(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Cust = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetDate(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Date = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetStartdate(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Startdate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetEnddate(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Enddate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetChannel(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Channel = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetIsExactMatch(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.IsExactMatch = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetAccetype(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Accetype = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetDataformat(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetResultType(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.ResultType = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetTranscodeType(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.TranscodeType = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetTimezone(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Timezone = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetDefinition(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Definition = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetIsAudio(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.IsAudio = &v
  return s
}

type QueryDailyLiveTranscodingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryDailyLiveTranscodingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponse) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponse) SetProvider(v *QueryDailyLiveTranscodingDurationResponseProvider) *QueryDailyLiveTranscodingDurationResponse {
  s.Provider = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'直播转码时长每日统计数据'}
  Date *QueryDailyLiveTranscodingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProvider) SetName(v string) *QueryDailyLiveTranscodingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProvider) SetType(v string) *QueryDailyLiveTranscodingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProvider) SetDate(v *QueryDailyLiveTranscodingDurationResponseProviderDate) *QueryDailyLiveTranscodingDurationResponseProvider {
  s.Date = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDate) SetStartdate(v string) *QueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDate) SetEnddate(v string) *QueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDate) SetTranscoding(v *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) *QueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Transcoding = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'live', 'zh_CN':'直播转码时长数据'}
  Live []*QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) SetName(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Name = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) SetLive(v []*QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Live = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'h264 transcoding time', 'zh_CN':'h264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H264 *string `json:"h264,omitempty" xml:"h264,omitempty" require:"true"`
  // {'en':'h265 transcoding time', 'zh_CN':'h265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H265 *string `json:"h265,omitempty" xml:"h265,omitempty" require:"true"`
  // {'en':'zdgq_264 transcoding time', 'zh_CN':'zdgq_264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_264 *string `json:"zdgq_264,omitempty" xml:"zdgq_264,omitempty" require:"true"`
  // {'en':'zdgq_265 transcoding time', 'zh_CN':'zdgq_265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_265 *string `json:"zdgq_265,omitempty" xml:"zdgq_265,omitempty" require:"true"`
  // {'en':'voice transcoding time', 'zh_CN':'voice转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty" require:"true"`
  // {'en':'total transcoding time', 'zh_CN':'总转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTime(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Time = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH264(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H264 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH265(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H265 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_264(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_264 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_265(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_265 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetVoice(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Voice = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTotal(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Total = &v
  return s
}

type QueryDailyLiveTranscodingDurationPaths struct {
}

func (s QueryDailyLiveTranscodingDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationPaths) GoString() string {
  return s.String()
}

type QueryDailyLiveTranscodingDurationParameters struct {
}

func (s QueryDailyLiveTranscodingDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationParameters) GoString() string {
  return s.String()
}

type QueryDailyLiveTranscodingDurationRequestHeader struct {
}

func (s QueryDailyLiveTranscodingDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationRequestHeader) GoString() string {
  return s.String()
}

type QueryDailyLiveTranscodingDurationResponseHeader struct {
}

func (s QueryDailyLiveTranscodingDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseHeader) GoString() string {
  return s.String()
}




type ChannelValueRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
}

func (s ChannelValueRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueRequest) GoString() string {
  return s.String()
}

func (s *ChannelValueRequest) SetCust(v string) *ChannelValueRequest {
  s.Cust = &v
  return s
}

func (s *ChannelValueRequest) SetDate(v string) *ChannelValueRequest {
  s.Date = &v
  return s
}

func (s *ChannelValueRequest) SetStartdate(v string) *ChannelValueRequest {
  s.Startdate = &v
  return s
}

func (s *ChannelValueRequest) SetEnddate(v string) *ChannelValueRequest {
  s.Enddate = &v
  return s
}

func (s *ChannelValueRequest) SetChannel(v string) *ChannelValueRequest {
  s.Channel = &v
  return s
}

func (s *ChannelValueRequest) SetRegion(v string) *ChannelValueRequest {
  s.Region = &v
  return s
}

func (s *ChannelValueRequest) SetDataformat(v string) *ChannelValueRequest {
  s.Dataformat = &v
  return s
}

func (s *ChannelValueRequest) SetIsExactMatch(v string) *ChannelValueRequest {
  s.IsExactMatch = &v
  return s
}

type ChannelValueResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *ChannelValueResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueResponse) GoString() string {
  return s.String()
}

func (s *ChannelValueResponse) SetProvider(v *ChannelValueResponseProvider) *ChannelValueResponse {
  s.Provider = v
  return s
}

type ChannelValueResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道流量数据'}
  Date *ChannelValueResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueResponseProvider) GoString() string {
  return s.String()
}

func (s *ChannelValueResponseProvider) SetName(v string) *ChannelValueResponseProvider {
  s.Name = &v
  return s
}

func (s *ChannelValueResponseProvider) SetType(v string) *ChannelValueResponseProvider {
  s.Type = &v
  return s
}

func (s *ChannelValueResponseProvider) SetDate(v *ChannelValueResponseProviderDate) *ChannelValueResponseProvider {
  s.Date = v
  return s
}

type ChannelValueResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'information', 'zh_CN':'信息集'}
  Information *ChannelValueResponseProviderDateInformation `json:"information,omitempty" xml:"information,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueResponseProviderDate) GoString() string {
  return s.String()
}

func (s *ChannelValueResponseProviderDate) SetStartdate(v string) *ChannelValueResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *ChannelValueResponseProviderDate) SetEnddate(v string) *ChannelValueResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *ChannelValueResponseProviderDate) SetInformation(v *ChannelValueResponseProviderDateInformation) *ChannelValueResponseProviderDate {
  s.Information = v
  return s
}

type ChannelValueResponseProviderDateInformation struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'charge method', 'zh_CN':'计费方式'}
  ChargeMethod *string `json:"chargeMethod,omitempty" xml:"chargeMethod,omitempty" require:"true"`
  // {'en':'acce type', 'zh_CN':'加速类型'}
  AcceType *string `json:"acceType,omitempty" xml:"acceType,omitempty" require:"true"`
  // {'en':'value', 'zh_CN':'计费值'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'unit', 'zh_CN':'单位'}
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty" require:"true"`
  // {'en':'peak value', 'zh_CN':'峰值'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'peak time', 'zh_CN':'峰值时间'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
}

func (s ChannelValueResponseProviderDateInformation) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueResponseProviderDateInformation) GoString() string {
  return s.String()
}

func (s *ChannelValueResponseProviderDateInformation) SetChannel(v string) *ChannelValueResponseProviderDateInformation {
  s.Channel = &v
  return s
}

func (s *ChannelValueResponseProviderDateInformation) SetChargeMethod(v string) *ChannelValueResponseProviderDateInformation {
  s.ChargeMethod = &v
  return s
}

func (s *ChannelValueResponseProviderDateInformation) SetAcceType(v string) *ChannelValueResponseProviderDateInformation {
  s.AcceType = &v
  return s
}

func (s *ChannelValueResponseProviderDateInformation) SetValue(v string) *ChannelValueResponseProviderDateInformation {
  s.Value = &v
  return s
}

func (s *ChannelValueResponseProviderDateInformation) SetUnit(v string) *ChannelValueResponseProviderDateInformation {
  s.Unit = &v
  return s
}

func (s *ChannelValueResponseProviderDateInformation) SetPeakValue(v string) *ChannelValueResponseProviderDateInformation {
  s.PeakValue = &v
  return s
}

func (s *ChannelValueResponseProviderDateInformation) SetPeakTime(v string) *ChannelValueResponseProviderDateInformation {
  s.PeakTime = &v
  return s
}

type ChannelValuePaths struct {
}

func (s ChannelValuePaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelValuePaths) GoString() string {
  return s.String()
}

type ChannelValueParameters struct {
}

func (s ChannelValueParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueParameters) GoString() string {
  return s.String()
}

type ChannelValueRequestHeader struct {
}

func (s ChannelValueRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueRequestHeader) GoString() string {
  return s.String()
}

type ChannelValueResponseHeader struct {
}

func (s ChannelValueResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainStreamDurationServiceRequest struct {
  // {"en":"Start time:\n1.Format is yyyyMMdd;\n2.Must be smaller than the current system time;\n3.Default value of current time is used if the field is not specified;\n4.You can only query data for the last 6 months.","zh_CN":"开始时间:\n1.时间格式为yyyy-MM-dd,例如,2021-08-10\n2.不能大于当前时间\n3.只能查询最近半年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:\n1. The time format is yyyy-MM-dd\n2. The end time must be greater than the start time. If the end time is greater than the current time, take the current time\n3. Both dateFrom and dateTo have not been passed, and the past 7 days are queried by default; if only one has not been passed, an exception will be thrown\n4. The maximum time interval allowed for query: 31 days, that is, the difference between dateFrom and dateTo cannot exceed 31 days","zh_CN":"结束时间:\n1.时间格式为yyyy-MM-dd\n2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间\n3.dateFrom,dateTo二者都未传,默认查询过去的7天;如仅有一个未传,抛异常\n4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"-","zh_CN":"-"}
  DomainStream []*ReportDomainStreamDurationServiceRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceRequest) SetDateFrom(v string) *ReportDomainStreamDurationServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainStreamDurationServiceRequest) SetDateTo(v string) *ReportDomainStreamDurationServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainStreamDurationServiceRequest) SetDomainStream(v []*ReportDomainStreamDurationServiceRequestDomainStream) *ReportDomainStreamDurationServiceRequest {
  s.DomainStream = v
  return s
}

type ReportDomainStreamDurationServiceRequestDomainStream struct     {
  // {"en":"Domain name:\n1. The maximum number of domain names that can be transferred is 1 by default;\n2. Automatically filter out invalid domain names (if an illegal domain name is transferred, it will be filtered out, and the query results will only return the data of valid domain names).","zh_CN":"域名:\n1、可传递域名数量上限默认为1个;\n2、自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"stream name:\nJust pass the publishing point + stream name, Example: live/test-20180101-test where live is a publishing point and test-20180101-test is a stream name","zh_CN":"流名:\n只需要传发布点+流名,例如:live/test-20180101-test ,其中live是发布点，test-20180101-test是流名"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequestDomainStream) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceRequestDomainStream) SetDomain(v string) *ReportDomainStreamDurationServiceRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamDurationServiceRequestDomainStream) SetStream(v []*string) *ReportDomainStreamDurationServiceRequestDomainStream {
  s.Stream = v
  return s
}

type ReportDomainStreamDurationServiceRequestHeader struct {
}

func (s ReportDomainStreamDurationServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServicePaths struct {
}

func (s ReportDomainStreamDurationServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServicePaths) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceParameters struct {
}

func (s ReportDomainStreamDurationServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the result of the request","zh_CN":"请求结果的详细数据"}
  Data []*ReportDomainStreamDurationServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponse) SetCode(v string) *ReportDomainStreamDurationServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponse) SetMessage(v string) *ReportDomainStreamDurationServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponse) SetData(v []*ReportDomainStreamDurationServiceResponseData) *ReportDomainStreamDurationServiceResponse {
  s.Data = v
  return s
}

type ReportDomainStreamDurationServiceResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"-","zh_CN":"-"}
  StreamList []*ReportDomainStreamDurationServiceResponseDataStreamList `json:"streamList,omitempty" xml:"streamList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponseData) SetDomain(v string) *ReportDomainStreamDurationServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseData) SetStreamList(v []*ReportDomainStreamDurationServiceResponseDataStreamList) *ReportDomainStreamDurationServiceResponseData {
  s.StreamList = v
  return s
}

type ReportDomainStreamDurationServiceResponseDataStreamList struct     {
  // {"en":"domain + publishing point + stream name","zh_CN":"流名(域名+发布点+流名)"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"The sum of the flow duration of the flow name in the corresponding time period, in milliseconds","zh_CN":"对应时间段内流名推流时长之和,单位为毫秒"}
  SumTime *ReportDomainStreamDurationServiceResponseDataStreamListSumTime `json:"sumTime,omitempty" xml:"sumTime,omitempty" require:"true" type:"Struct"`
  // {"en":"-","zh_CN":"-"}
  DurationDetailList []*ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList `json:"durationDetailList,omitempty" xml:"durationDetailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceResponseDataStreamList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamList) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamList) SetStream(v string) *ReportDomainStreamDurationServiceResponseDataStreamList {
  s.Stream = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamList) SetSumTime(v *ReportDomainStreamDurationServiceResponseDataStreamListSumTime) *ReportDomainStreamDurationServiceResponseDataStreamList {
  s.SumTime = v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamList) SetDurationDetailList(v []*ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) *ReportDomainStreamDurationServiceResponseDataStreamList {
  s.DurationDetailList = v
  return s
}

type ReportDomainStreamDurationServiceResponseDataStreamListSumTime struct {
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListSumTime) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListSumTime) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList struct     {
  // {"en":"Stream start time","zh_CN":"推流起始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Stream end time","zh_CN":"推流终止时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Streaming duration, in milliseconds","zh_CN":"推流时长,单位为毫秒"}
  Duration *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration `json:"duration,omitempty" xml:"duration,omitempty" require:"true" type:"Struct"`
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetStartTime(v string) *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.StartTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetEndTime(v string) *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.EndTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetDuration(v *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration) *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.Duration = v
  return s
}

type ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration struct {
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceResponseHeader struct {
}

func (s ReportDomainStreamDurationServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseHeader) GoString() string {
  return s.String()
}




type FlowAppaChannelRequest struct {
  // {"en":"cust_en_name of sub-client.\nWhen a merged-account wants to  view the information of the subclient,the cust_en_name is required.","zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not specified,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope.\n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope.\n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:\n1)If there are multiple inputs,use  ';' as separator.\n2)If not specified, it means all the domains of the account .","zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.\n1)If there are multiple inputs,use ';' as separator.\n2)If not specified or specified as 'all', it means all the accetypes.","zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"If TopN is empty, the system will return the top 200 by default. Otherwise, it returns data based on the TopN parameter value. If TopN is set to -1, all domains will be returned (subject to the domain quantity limits required by the API).","zh_CN":"TopN为空,则默认返回top200, 否则按topN参数值返回, TopN为-1，则返回全部域名(接口限制要求的域名个数)"}
  TopN *string `json:"topN,omitempty" xml:"topN,omitempty"`
  // {"en":"Ranking method: 1: Total Traffic Ranking; 2: Total Bandwidth Peak Value Ranking; 3: Edge Downstream Traffic Ranking; 4: Edge Downstream Bandwidth Peak Value Ranking; 5: Edge Upstream Traffic Ranking; 6: Edge Upstream Bandwidth Peak Value Ranking","zh_CN":"排行方式: 1:总流量排行;2:总带宽峰值排行; 3:边缘下行流量排行; 4:边缘下行带宽峰值排行; 5:边缘上行流量排行; 6:边缘上行带宽峰值排行"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s FlowAppaChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelRequest) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelRequest) SetCust(v string) *FlowAppaChannelRequest {
  s.Cust = &v
  return s
}

func (s *FlowAppaChannelRequest) SetDate(v string) *FlowAppaChannelRequest {
  s.Date = &v
  return s
}

func (s *FlowAppaChannelRequest) SetStartdate(v string) *FlowAppaChannelRequest {
  s.Startdate = &v
  return s
}

func (s *FlowAppaChannelRequest) SetEnddate(v string) *FlowAppaChannelRequest {
  s.Enddate = &v
  return s
}

func (s *FlowAppaChannelRequest) SetChannel(v string) *FlowAppaChannelRequest {
  s.Channel = &v
  return s
}

func (s *FlowAppaChannelRequest) SetTimezone(v string) *FlowAppaChannelRequest {
  s.Timezone = &v
  return s
}

func (s *FlowAppaChannelRequest) SetRegion(v string) *FlowAppaChannelRequest {
  s.Region = &v
  return s
}

func (s *FlowAppaChannelRequest) SetAccetype(v string) *FlowAppaChannelRequest {
  s.Accetype = &v
  return s
}

func (s *FlowAppaChannelRequest) SetDataformat(v string) *FlowAppaChannelRequest {
  s.Dataformat = &v
  return s
}

func (s *FlowAppaChannelRequest) SetTopN(v string) *FlowAppaChannelRequest {
  s.TopN = &v
  return s
}

func (s *FlowAppaChannelRequest) SetSortBy(v string) *FlowAppaChannelRequest {
  s.SortBy = &v
  return s
}

type FlowAppaChannelRequestHeader struct {
}

func (s FlowAppaChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelRequestHeader) GoString() string {
  return s.String()
}

type FlowAppaChannelPaths struct {
}

func (s FlowAppaChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelPaths) GoString() string {
  return s.String()
}

type FlowAppaChannelParameters struct {
}

func (s FlowAppaChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelParameters) GoString() string {
  return s.String()
}

type FlowAppaChannelResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *FlowAppaChannelResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s FlowAppaChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponse) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelResponse) SetProvider(v *FlowAppaChannelResponseProvider) *FlowAppaChannelResponse {
  s.Provider = v
  return s
}

type FlowAppaChannelResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"请求数数据"}
  Date *FlowAppaChannelResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s FlowAppaChannelResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponseProvider) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelResponseProvider) SetName(v string) *FlowAppaChannelResponseProvider {
  s.Name = &v
  return s
}

func (s *FlowAppaChannelResponseProvider) SetType(v string) *FlowAppaChannelResponseProvider {
  s.Type = &v
  return s
}

func (s *FlowAppaChannelResponseProvider) SetDate(v *FlowAppaChannelResponseProviderDate) *FlowAppaChannelResponseProvider {
  s.Date = v
  return s
}

type FlowAppaChannelResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"total","zh_CN":"汇总"}
  Total *FlowAppaChannelResponseProviderDateTotal `json:"total,omitempty" xml:"total,omitempty" require:"true" type:"Struct"`
  // {"en":"channel","zh_CN":"频道"}
  Channel *FlowAppaChannelResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s FlowAppaChannelResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponseProviderDate) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelResponseProviderDate) SetStartdate(v string) *FlowAppaChannelResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDate) SetEnddate(v string) *FlowAppaChannelResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDate) SetTotal(v *FlowAppaChannelResponseProviderDateTotal) *FlowAppaChannelResponseProviderDate {
  s.Total = v
  return s
}

func (s *FlowAppaChannelResponseProviderDate) SetChannel(v *FlowAppaChannelResponseProviderDateChannel) *FlowAppaChannelResponseProviderDate {
  s.Channel = v
  return s
}

type FlowAppaChannelResponseProviderDateTotal struct {
  // {"en":"channel","zh_CN":"频道"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"edgeup","zh_CN":"边缘上行总流量,单位Mbps"}
  Edgeup *string `json:"edgeup,omitempty" xml:"edgeup,omitempty" require:"true"`
  // {"en":"edgedown","zh_CN":"边缘下行总流量,单位Mbps"}
  Edgedown *string `json:"edgedown,omitempty" xml:"edgedown,omitempty" require:"true"`
  // {"en":"total","zh_CN":"汇总"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"edgeupPeakValue","zh_CN":"边缘上行带宽峰值(MBps)"}
  EdgeupPeakValue *string `json:"edgeupPeakValue,omitempty" xml:"edgeupPeakValue,omitempty" require:"true"`
  // {"en":"edgedownPeakValue","zh_CN":"边缘下行带宽峰值(MBps)"}
  EdgedownPeakValue *string `json:"edgedownPeakValue,omitempty" xml:"edgedownPeakValue,omitempty" require:"true"`
  // {"en":"totalPeakValue","zh_CN":"总带宽峰值:边缘上行流量和边缘下行流量叠加后取带宽峰值(MBps)"}
  TotalPeakValue *string `json:"totalPeakValue,omitempty" xml:"totalPeakValue,omitempty" require:"true"`
  // {"en":"The moment corresponding to the total bandwidth peak (sum of edge upstream traffic and edge downstream traffic)","zh_CN":"总带宽峰值对应的时刻 （边缘上行流量和边缘下行流量叠加）"}
  TotalPeakTime *string `json:"totalPeakTime,omitempty" xml:"totalPeakTime,omitempty" require:"true"`
}

func (s FlowAppaChannelResponseProviderDateTotal) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponseProviderDateTotal) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetName(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.Name = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetEdgeup(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.Edgeup = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetEdgedown(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.Edgedown = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetTotal(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.Total = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetEdgeupPeakValue(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.EdgeupPeakValue = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetEdgedownPeakValue(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.EdgedownPeakValue = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetTotalPeakValue(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.TotalPeakValue = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateTotal) SetTotalPeakTime(v string) *FlowAppaChannelResponseProviderDateTotal {
  s.TotalPeakTime = &v
  return s
}

type FlowAppaChannelResponseProviderDateChannel struct {
  // {"en":"channel","zh_CN":"频道"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"edgeup","zh_CN":"边缘上行总流量,单位Mbps"}
  Edgeup *string `json:"edgeup,omitempty" xml:"edgeup,omitempty" require:"true"`
  // {"en":"edgedown","zh_CN":"边缘下行总流量,单位Mbps"}
  Edgedown *string `json:"edgedown,omitempty" xml:"edgedown,omitempty" require:"true"`
  // {"en":"total","zh_CN":"汇总"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"edgeupPeakValue","zh_CN":"边缘上行带宽峰值(MBps)"}
  EdgeupPeakValue *string `json:"edgeupPeakValue,omitempty" xml:"edgeupPeakValue,omitempty" require:"true"`
  // {"en":"edgedownPeakValue","zh_CN":"边缘下行带宽峰值(MBps)"}
  EdgedownPeakValue *string `json:"edgedownPeakValue,omitempty" xml:"edgedownPeakValue,omitempty" require:"true"`
  // {"en":"totalPeakValue","zh_CN":"总带宽峰值:边缘上行流量和边缘下行流量叠加后取带宽峰值(MBps)"}
  TotalPeakValue *string `json:"totalPeakValue,omitempty" xml:"totalPeakValue,omitempty" require:"true"`
  // {"en":"The moment when the channel bandwidth peak occurs (the sum of edge upstream traffic and edge downstream traffic)","zh_CN":"频道带宽峰值对应的时刻 （边缘上行流量和边缘下行流量叠加）"}
  TotalPeakTime *string `json:"totalPeakTime,omitempty" xml:"totalPeakTime,omitempty" require:"true"`
}

func (s FlowAppaChannelResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetName(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetEdgeup(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.Edgeup = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetEdgedown(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.Edgedown = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetTotal(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.Total = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetEdgeupPeakValue(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.EdgeupPeakValue = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetEdgedownPeakValue(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.EdgedownPeakValue = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetTotalPeakValue(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.TotalPeakValue = &v
  return s
}

func (s *FlowAppaChannelResponseProviderDateChannel) SetTotalPeakTime(v string) *FlowAppaChannelResponseProviderDateChannel {
  s.TotalPeakTime = &v
  return s
}

type FlowAppaChannelResponseHeader struct {
}

func (s FlowAppaChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponseHeader) GoString() string {
  return s.String()
}




type ReportUrlDlFinishServiceRequest struct {
  // {'en':'Start time
  // 
  // 1.         The   format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.         Must   be a time that is 183 days earlier than the current time, and the time must   be earlier than the current time and dateTo;
  // 
  // 3.         Period   between dataFrom and dateTo cannot be longer than 7 days;
  // 
  // 4.         dateFrom   and dateTo can be either both are specified or neither is specifies;
  // 
  // 5.         If   neither dateFrom nor dateTo is specified, then by default, data in the last   24 hour is queried', 'zh_CN':'开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于当前时间-183天，并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过7天；（可联系技术支持调整）
  // 4.dateFrom和dateTo要么都传递，要么都不传递；
  // 5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time
  // 
  // 1.         The   format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2.         Must   be greater than dateFrom; if it&rsquo;s greater than the current time, then the   current time is assigned as the value;', 'zh_CN':'结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；如果大于当前时间，则重新赋值为当前时间；'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain names, domain number   limits can be adjusted depending on different accounts. The default value is   20', 'zh_CN':'域名，域名个数限制根据账号可调，默认为20个'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'url：
  // 1. fuzzy matching is supported.
  // 2.Several are separated by ','.', 'zh_CN':'url：
  // 1.支持模糊匹配。
  // 2.多个用 '，'隔开。'}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ReportUrlDlFinishServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUrlDlFinishServiceRequest) SetDateFrom(v string) *ReportUrlDlFinishServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUrlDlFinishServiceRequest) SetDateTo(v string) *ReportUrlDlFinishServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUrlDlFinishServiceRequest) SetDomain(v []*string) *ReportUrlDlFinishServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUrlDlFinishServiceRequest) SetUrl(v string) *ReportUrlDlFinishServiceRequest {
  s.Url = &v
  return s
}

type ReportUrlDlFinishServiceResponse struct {
  // {'en':'-', 'zh_CN':'请求结果的详细数据'}
  Result []*ReportUrlDlFinishServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlDlFinishServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUrlDlFinishServiceResponse) SetResult(v []*ReportUrlDlFinishServiceResponseResult) *ReportUrlDlFinishServiceResponse {
  s.Result = v
  return s
}

type ReportUrlDlFinishServiceResponseResult struct     {
  // {'en':'The top500 of url', 'zh_CN':'top500的url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'Number of successful downloads', 'zh_CN':'下载成功数'}
  Num *string `json:"num,omitempty" xml:"num,omitempty" require:"true"`
}

func (s ReportUrlDlFinishServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportUrlDlFinishServiceResponseResult) SetUrl(v string) *ReportUrlDlFinishServiceResponseResult {
  s.Url = &v
  return s
}

func (s *ReportUrlDlFinishServiceResponseResult) SetNum(v string) *ReportUrlDlFinishServiceResponseResult {
  s.Num = &v
  return s
}

type ReportUrlDlFinishServicePaths struct {
}

func (s ReportUrlDlFinishServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServicePaths) GoString() string {
  return s.String()
}

type ReportUrlDlFinishServiceParameters struct {
}

func (s ReportUrlDlFinishServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceParameters) GoString() string {
  return s.String()
}

type ReportUrlDlFinishServiceRequestHeader struct {
}

func (s ReportUrlDlFinishServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUrlDlFinishServiceResponseHeader struct {
}

func (s ReportUrlDlFinishServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDDoSMitigatedBandwidthBySuiteOrProductRequest struct {
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"PackageId or acctype should be selected at least one.", "zh_CN":"套餐ID: packageId和acctype至少传一个,但不能同时传。"}
  PackageId *string `json:"packageId,omitempty" xml:"packageId,omitempty"`
  // {"en":"need Detail : 1
  // no need Detail: 0
  // default : 1.", "zh_CN":"是否需要查看域名或是转发规则带宽的详细信息：0：不需要；1：需要，默认需要。"}
  NeedDetail *int `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"Customer English Name.", "zh_CN":"客户英文名。"}
  CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty" require:"true"`
  // {"en":"PackageId or acctype should be selected at least one.
  // acctype( Only One can be selected): gess, fsa, app-s, dms-https, wss, dms, wss-https, s-appa, esa, wsa-https, 1551.", "zh_CN":"PackageId和acctype不能同时传且至少传一个；产品外部服务类型,只支持传1个:gess，fsa，app-s，dms-https，wss, dms， wss-https，s-appa，esa，wsa-https，1551。"}
  Acctype *string `json:"acctype,omitempty" xml:"acctype,omitempty"`
  // {"en":"ContractId", "zh_CN":"合同号  
  // 支持按Contract# item#粒度查询
  // 当ContractId不为空时，会替换掉packageId"}
  ContractId *string `json:"ContractId,omitempty" xml:"ContractId,omitempty"`
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) GoString() string {
  return s.String()
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetStartdate(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.Startdate = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetPackageId(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.PackageId = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetNeedDetail(v int) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.NeedDetail = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetEnddate(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.Enddate = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetCustomCode(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.CustomCode = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetAcctype(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.Acctype = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetContractId(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.ContractId = &v
  return s
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductResponse struct {
  // {"en":"错误信息或Success。", "zh_CN":"Error message or Success."}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"CleanBandwidth:
  //     time: "2018-08-21 00:00:00"
  //    atFlow: cleaned bandwidth Mbps", "zh_CN":"带宽信息： 
  // Time："2018-08-21 00:00:00"
  // atFlow：已清洗的带宽　Mbps"}
  At_bw *string `json:"at_bw,omitempty" xml:"at_bw,omitempty" require:"true"`
  // {"en":"peak information:
  //    peakTime: "2018-08-21 00:00:00"
  //    peakValue:  peak of Clean Bandwidth (Mbps)
  //    mitigatedTraffic: Cleaned flow (GB)", "zh_CN":"峰值统计信息：
  // peakTime：峰值时间 "2018-08-21 00:00:00"
  // peakValue：DDoS攻击峰值带宽Mbps
  // mitigatedTraffic：已清洗的流量 GB"}
  PeakStat *string `json:"peakStat,omitempty" xml:"peakStat,omitempty" require:"true"`
  // {"en":"Return 200 means success, please see <Error code> to check other status code.", "zh_CN":"200状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) GoString() string {
  return s.String()
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetMsg(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.Msg = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetAt_bw(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.At_bw = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetPeakStat(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.PeakStat = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetCode(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.Code = &v
  return s
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductPaths struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductPaths) GoString() string {
  return s.String()
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductParameters struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductParameters) GoString() string {
  return s.String()
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductRequestHeader struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequestHeader) GoString() string {
  return s.String()
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductResponseHeader struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponseHeader) GoString() string {
  return s.String()
}




type CloudDirectDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
}

func (s CloudDirectDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationRequest) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationRequest) SetCust(v string) *CloudDirectDurationRequest {
  s.Cust = &v
  return s
}

func (s *CloudDirectDurationRequest) SetDate(v string) *CloudDirectDurationRequest {
  s.Date = &v
  return s
}

func (s *CloudDirectDurationRequest) SetStartdate(v string) *CloudDirectDurationRequest {
  s.Startdate = &v
  return s
}

func (s *CloudDirectDurationRequest) SetEnddate(v string) *CloudDirectDurationRequest {
  s.Enddate = &v
  return s
}

func (s *CloudDirectDurationRequest) SetTimezone(v string) *CloudDirectDurationRequest {
  s.Timezone = &v
  return s
}

func (s *CloudDirectDurationRequest) SetRegion(v string) *CloudDirectDurationRequest {
  s.Region = &v
  return s
}

func (s *CloudDirectDurationRequest) SetDataformat(v string) *CloudDirectDurationRequest {
  s.Dataformat = &v
  return s
}

type CloudDirectDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *CloudDirectDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s CloudDirectDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponse) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationResponse) SetProvider(v *CloudDirectDurationResponseProvider) *CloudDirectDurationResponse {
  s.Provider = v
  return s
}

type CloudDirectDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Date *CloudDirectDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s CloudDirectDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationResponseProvider) SetName(v string) *CloudDirectDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *CloudDirectDurationResponseProvider) SetType(v string) *CloudDirectDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *CloudDirectDurationResponseProvider) SetDate(v *CloudDirectDurationResponseProviderDate) *CloudDirectDurationResponseProvider {
  s.Date = v
  return s
}

type CloudDirectDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'totalDuration', 'zh_CN':'汇总'}
  TotalDuration *string `json:"totalDuration,omitempty" xml:"totalDuration,omitempty" require:"true"`
  // {'en':'result', 'zh_CN':'明细数据'}
  Result *CloudDirectDurationResponseProviderDateResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s CloudDirectDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationResponseProviderDate) SetStartdate(v string) *CloudDirectDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *CloudDirectDurationResponseProviderDate) SetEnddate(v string) *CloudDirectDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *CloudDirectDurationResponseProviderDate) SetTotalDuration(v string) *CloudDirectDurationResponseProviderDate {
  s.TotalDuration = &v
  return s
}

func (s *CloudDirectDurationResponseProviderDate) SetResult(v *CloudDirectDurationResponseProviderDateResult) *CloudDirectDurationResponseProviderDate {
  s.Result = v
  return s
}

type CloudDirectDurationResponseProviderDateResult struct {
  // {'en':'duration', 'zh_CN':'明细数据'}
  Duration []*CloudDirectDurationResponseProviderDateResultDuration `json:"duration,omitempty" xml:"duration,omitempty" require:"true" type:"Repeated"`
}

func (s CloudDirectDurationResponseProviderDateResult) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponseProviderDateResult) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationResponseProviderDateResult) SetDuration(v []*CloudDirectDurationResponseProviderDateResultDuration) *CloudDirectDurationResponseProviderDateResult {
  s.Duration = v
  return s
}

type CloudDirectDurationResponseProviderDateResultDuration struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数据'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s CloudDirectDurationResponseProviderDateResultDuration) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponseProviderDateResultDuration) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationResponseProviderDateResultDuration) SetTime(v string) *CloudDirectDurationResponseProviderDateResultDuration {
  s.Time = &v
  return s
}

func (s *CloudDirectDurationResponseProviderDateResultDuration) SetText(v string) *CloudDirectDurationResponseProviderDateResultDuration {
  s.Text = &v
  return s
}

type CloudDirectDurationPaths struct {
}

func (s CloudDirectDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationPaths) GoString() string {
  return s.String()
}

type CloudDirectDurationParameters struct {
}

func (s CloudDirectDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationParameters) GoString() string {
  return s.String()
}

type CloudDirectDurationRequestHeader struct {
}

func (s CloudDirectDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationRequestHeader) GoString() string {
  return s.String()
}

type CloudDirectDurationResponseHeader struct {
}

func (s CloudDirectDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponseHeader) GoString() string {
  return s.String()
}




type BandwidthUploadRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Bandwidth types.
  // 1)optional values:edgeUp,midUp.
  // 2)If there are multiple inpus,use ';' as delemeter.", "zh_CN":"带宽类型：edgeUp：边缘上行, midUp：中间上行, 多个类型以英文分号';'分隔. 空默认为边缘上行(edgeUp)"}
  FlowType *string `json:"flowType,omitempty" xml:"flowType,omitempty"`
}

func (s BandwidthUploadRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadRequest) GoString() string {
  return s.String()
}

func (s *BandwidthUploadRequest) SetCust(v string) *BandwidthUploadRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthUploadRequest) SetDate(v string) *BandwidthUploadRequest {
  s.Date = &v
  return s
}

func (s *BandwidthUploadRequest) SetStartdate(v string) *BandwidthUploadRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthUploadRequest) SetEnddate(v string) *BandwidthUploadRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthUploadRequest) SetChannel(v string) *BandwidthUploadRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthUploadRequest) SetIsExactMatch(v string) *BandwidthUploadRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthUploadRequest) SetRegion(v string) *BandwidthUploadRequest {
  s.Region = &v
  return s
}

func (s *BandwidthUploadRequest) SetAccetype(v string) *BandwidthUploadRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthUploadRequest) SetDataformat(v string) *BandwidthUploadRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthUploadRequest) SetFlowType(v string) *BandwidthUploadRequest {
  s.FlowType = &v
  return s
}

type BandwidthUploadResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthUploadResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthUploadResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponse) GoString() string {
  return s.String()
}

func (s *BandwidthUploadResponse) SetProvider(v *BandwidthUploadResponseProvider) *BandwidthUploadResponse {
  s.Provider = v
  return s
}

type BandwidthUploadResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *BandwidthUploadResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthUploadResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthUploadResponseProvider) SetName(v string) *BandwidthUploadResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthUploadResponseProvider) SetType(v string) *BandwidthUploadResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthUploadResponseProvider) SetDate(v *BandwidthUploadResponseProviderDate) *BandwidthUploadResponseProvider {
  s.Date = v
  return s
}

type BandwidthUploadResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthUploadResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthUploadResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthUploadResponseProviderDate) SetStartdate(v string) *BandwidthUploadResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthUploadResponseProviderDate) SetEnddate(v string) *BandwidthUploadResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthUploadResponseProviderDate) SetChannel(v *BandwidthUploadResponseProviderDateChannel) *BandwidthUploadResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthUploadResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*BandwidthUploadResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthUploadResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthUploadResponseProviderDateChannel) SetName(v string) *BandwidthUploadResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthUploadResponseProviderDateChannel) SetBandwidth(v []*BandwidthUploadResponseProviderDateChannelBandwidth) *BandwidthUploadResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthUploadResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthUploadResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthUploadResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthUploadResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthUploadResponseProviderDateChannelBandwidth) SetText(v string) *BandwidthUploadResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type BandwidthUploadPaths struct {
}

func (s BandwidthUploadPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadPaths) GoString() string {
  return s.String()
}

type BandwidthUploadParameters struct {
}

func (s BandwidthUploadParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadParameters) GoString() string {
  return s.String()
}

type BandwidthUploadRequestHeader struct {
}

func (s BandwidthUploadRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadRequestHeader) GoString() string {
  return s.String()
}

type BandwidthUploadResponseHeader struct {
}

func (s BandwidthUploadResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponseHeader) GoString() string {
  return s.String()
}




type HttpTestRequest struct {
  // {"en":"URL", "zh_CN":"指定检测的 URL"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"area", "zh_CN":"监控机所属地区中文名称，支持中国大陆省份、港澳台以及海外国家。
  // 可选值：
  // anhui: 安徽
  // beijing: 北京
  // chongqing: 重庆
  // fujian: 福建
  // gansu: 甘肃
  // guangdong: 广东
  // guangxi: 广西
  // guizhou: 贵州
  // hainan: 海南
  // hebei: 河北
  // heilongjiang: 黑龙江
  // henan: 河南
  // hubei: 湖北
  // hunan: 湖南
  // jiangsu: 江苏
  // jiangxi: 江西
  // jilin: 吉林
  // liaoning: 辽宁
  // neimenggu: 内蒙古
  // ningxia: 宁夏
  // qinghai: 青海
  // shaanxi: 陕西
  // shandong: 山东
  // shanghai: 上海
  // shanxi: 山西
  // sichuan: 四川
  // tianjin: 天津
  // xinjiang: 新疆
  // xizang: 西藏
  // yunnan: 云南
  // zhejiang: 浙江
  // TW: 台湾
  // HK: 香港
  // MO: 澳门
  // AE: 阿联酋
  // AU: 澳大利亚
  // BD: 孟加拉
  // BN: 文莱
  // BR: 巴西
  // CA: 加拿大
  // CL: 智利
  // CO: 哥伦比亚
  // DJ: 吉布提
  // ID: 印度尼西亚
  // IT: 意大利
  // JP: 日本
  // KG: 吉尔吉斯斯坦
  // KH: 柬埔寨
  // KR: 韩国
  // KW: 科威特
  // LA: 老挝
  // MG: 马达加斯加
  // MM: 缅甸
  // MU: 毛里求斯
  // MY: 马来西亚
  // NP: 尼泊尔
  // OM: 阿曼
  // PE: 秘鲁
  // PH: 菲律宾
  // PK: 巴基斯坦
  // QA: 卡塔尔
  // RO: 罗马尼亚
  // RU: 俄罗斯
  // SA: 沙特阿拉伯
  // SE: 瑞典
  // SG: 新加坡
  // TH: 泰国
  // TR: 土耳其
  // US: 美国
  // VN: 越南"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"监控机所属运营商中文名。
  // 可选值：
  // 0: 中国电信
  // 1: 中国联通
  // 2: 中国铁通
  // 4: 中国移动
  // 5: 中国教育网
  // 9: 中国广电
  // 10: 长城宽带"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s HttpTestRequest) String() string {
  return tea.Prettify(s)
}

func (s HttpTestRequest) GoString() string {
  return s.String()
}

func (s *HttpTestRequest) SetUrl(v string) *HttpTestRequest {
  s.Url = &v
  return s
}

func (s *HttpTestRequest) SetArea(v string) *HttpTestRequest {
  s.Area = &v
  return s
}

func (s *HttpTestRequest) SetIsp(v string) *HttpTestRequest {
  s.Isp = &v
  return s
}

type HttpTestResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result *HttpTestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s HttpTestResponse) String() string {
  return tea.Prettify(s)
}

func (s HttpTestResponse) GoString() string {
  return s.String()
}

func (s *HttpTestResponse) SetResult(v *HttpTestResponseResult) *HttpTestResponse {
  s.Result = v
  return s
}

type HttpTestResponseResult struct {
  // {'en':'status', 'zh_CN':'状态码'}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'error message', 'zh_CN':'异常信息'}
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty" require:"true"`
  // {'en':'test id', 'zh_CN':'任务 ID'}
  TestId *string `json:"testId,omitempty" xml:"testId,omitempty" require:"true"`
  // {'en':'url', 'zh_CN':'目标 URL'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'data','zh_CN':'探测数据'}
  Data []*HttpTestResponseResultData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s HttpTestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s HttpTestResponseResult) GoString() string {
  return s.String()
}

func (s *HttpTestResponseResult) SetStatus(v string) *HttpTestResponseResult {
  s.Status = &v
  return s
}

func (s *HttpTestResponseResult) SetErrorMsg(v string) *HttpTestResponseResult {
  s.ErrorMsg = &v
  return s
}

func (s *HttpTestResponseResult) SetTestId(v string) *HttpTestResponseResult {
  s.TestId = &v
  return s
}

func (s *HttpTestResponseResult) SetUrl(v string) *HttpTestResponseResult {
  s.Url = &v
  return s
}

func (s *HttpTestResponseResult) SetData(v []*HttpTestResponseResultData) *HttpTestResponseResult {
  s.Data = v
  return s
}

type HttpTestResponseResultData struct     {
  // {'en':'id', 'zh_CN':'任务 ID'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'task id', 'zh_CN':'任务 ID'}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {'en':'url', 'zh_CN':'目标 URL'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'monitor ip', 'zh_CN':'监控机 IP'}
  DetectIp *string `json:"detectIp,omitempty" xml:"detectIp,omitempty" require:"true"`
  // {'en':'monitor isp', 'zh_CN':'监控机运营商'}
  DetectIpIsp *string `json:"detectIpIsp,omitempty" xml:"detectIpIsp,omitempty" require:"true"`
  // {'en':'monitor isp code', 'zh_CN':'监控机运营商编码'}
  DetectIpIspCode *string `json:"detectIpIspCode,omitempty" xml:"detectIpIspCode,omitempty" require:"true"`
  // {'en':'monitor province', 'zh_CN':'监控机所属省份'}
  DetectIpPro *string `json:"detectIpPro,omitempty" xml:"detectIpPro,omitempty" require:"true"`
  // {'en':'monitor province code', 'zh_CN':'监控机所属省份编码'}
  DetectIpProCode *string `json:"detectIpProCode,omitempty" xml:"detectIpProCode,omitempty" require:"true"`
  // {'en':'target ip', 'zh_CN':'探测目标 IP'}
  TargetIp *string `json:"targetIp,omitempty" xml:"targetIp,omitempty" require:"true"`
  // {'en':'target isp', 'zh_CN':'探测目标运营商'}
  TargetIpIsp *string `json:"targetIpIsp,omitempty" xml:"targetIpIsp,omitempty" require:"true"`
  // {'en':'tartget province', 'zh_CN':'探测目标所属省份'}
  TargetIpPro *string `json:"targetIpPro,omitempty" xml:"targetIpPro,omitempty" require:"true"`
  // {'en':'target province and isp', 'zh_CN':'探测目标省份运营商'}
  TargetLocation *string `json:"targetLocation,omitempty" xml:"targetLocation,omitempty" require:"true"`
  // {'en':'dns', 'zh_CN':'DNS'}
  DnsIp *string `json:"dnsIp,omitempty" xml:"dnsIp,omitempty" require:"true"`
  // {'en':'connect time', 'zh_CN':'建联耗时，单位：毫秒'}
  ConnectTime *int32 `json:"connectTime,omitempty" xml:"connectTime,omitempty" require:"true"`
  // {'en':'dns time', 'zh_CN':'DNS 查询耗时，单位：毫秒'}
  DnsTime *int32 `json:"dnsTime,omitempty" xml:"dnsTime,omitempty" require:"true"`
  // {'en':'http redirect time', 'zh_CN':'重定向耗时，单位：毫秒'}
  RedirectTime *int32 `json:"redirectTime,omitempty" xml:"redirectTime,omitempty" require:"true"`
  // {'en':'response time', 'zh_CN':'响应耗时，单位：毫秒'}
  ResponseTime *int32 `json:"responseTime,omitempty" xml:"responseTime,omitempty" require:"true"`
  // {'en':'all time', 'zh_CN':'探测总耗时，单位：毫秒'}
  UseTime *int32 `json:"useTime,omitempty" xml:"useTime,omitempty" require:"true"`
  // {'en':'end time', 'zh_CN':'探测结束时间戳'}
  DoneTime *int64 `json:"doneTime,omitempty" xml:"doneTime,omitempty" require:"true"`
  // {'en':'error code', 'zh_CN':'错误码'}
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty" require:"true"`
  // {'en':'file size', 'zh_CN':'下载文件大小，单位：Byte'}
  FileSize *int32 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
  // {'en':'download rate', 'zh_CN':'下载速率，单位：KBytes/s'}
  Rate *float32 `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {'en':'http status code', 'zh_CN':'HTTP 状态码'}
  HttpCode *string `json:"httpCode,omitempty" xml:"httpCode,omitempty" require:"true"`
  // {'en':'http request header', 'zh_CN':'请求头'}
  HttpTestRequestHeader *string `json:"requestHeader,omitempty" xml:"requestHeader,omitempty" require:"true"`
  // {'en':'http response header', 'zh_CN':'响应头'}
  HttpTestResponseHeader *string `json:"responseHeader,omitempty" xml:"responseHeader,omitempty" require:"true"`
  // {'en':'status', 'zh_CN':'探测状态'}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s HttpTestResponseResultData) String() string {
  return tea.Prettify(s)
}

func (s HttpTestResponseResultData) GoString() string {
  return s.String()
}

func (s *HttpTestResponseResultData) SetId(v string) *HttpTestResponseResultData {
  s.Id = &v
  return s
}

func (s *HttpTestResponseResultData) SetTaskId(v string) *HttpTestResponseResultData {
  s.TaskId = &v
  return s
}

func (s *HttpTestResponseResultData) SetUrl(v string) *HttpTestResponseResultData {
  s.Url = &v
  return s
}

func (s *HttpTestResponseResultData) SetDetectIp(v string) *HttpTestResponseResultData {
  s.DetectIp = &v
  return s
}

func (s *HttpTestResponseResultData) SetDetectIpIsp(v string) *HttpTestResponseResultData {
  s.DetectIpIsp = &v
  return s
}

func (s *HttpTestResponseResultData) SetDetectIpIspCode(v string) *HttpTestResponseResultData {
  s.DetectIpIspCode = &v
  return s
}

func (s *HttpTestResponseResultData) SetDetectIpPro(v string) *HttpTestResponseResultData {
  s.DetectIpPro = &v
  return s
}

func (s *HttpTestResponseResultData) SetDetectIpProCode(v string) *HttpTestResponseResultData {
  s.DetectIpProCode = &v
  return s
}

func (s *HttpTestResponseResultData) SetTargetIp(v string) *HttpTestResponseResultData {
  s.TargetIp = &v
  return s
}

func (s *HttpTestResponseResultData) SetTargetIpIsp(v string) *HttpTestResponseResultData {
  s.TargetIpIsp = &v
  return s
}

func (s *HttpTestResponseResultData) SetTargetIpPro(v string) *HttpTestResponseResultData {
  s.TargetIpPro = &v
  return s
}

func (s *HttpTestResponseResultData) SetTargetLocation(v string) *HttpTestResponseResultData {
  s.TargetLocation = &v
  return s
}

func (s *HttpTestResponseResultData) SetDnsIp(v string) *HttpTestResponseResultData {
  s.DnsIp = &v
  return s
}

func (s *HttpTestResponseResultData) SetConnectTime(v int32) *HttpTestResponseResultData {
  s.ConnectTime = &v
  return s
}

func (s *HttpTestResponseResultData) SetDnsTime(v int32) *HttpTestResponseResultData {
  s.DnsTime = &v
  return s
}

func (s *HttpTestResponseResultData) SetRedirectTime(v int32) *HttpTestResponseResultData {
  s.RedirectTime = &v
  return s
}

func (s *HttpTestResponseResultData) SetResponseTime(v int32) *HttpTestResponseResultData {
  s.ResponseTime = &v
  return s
}

func (s *HttpTestResponseResultData) SetUseTime(v int32) *HttpTestResponseResultData {
  s.UseTime = &v
  return s
}

func (s *HttpTestResponseResultData) SetDoneTime(v int64) *HttpTestResponseResultData {
  s.DoneTime = &v
  return s
}

func (s *HttpTestResponseResultData) SetErrorCode(v string) *HttpTestResponseResultData {
  s.ErrorCode = &v
  return s
}

func (s *HttpTestResponseResultData) SetFileSize(v int32) *HttpTestResponseResultData {
  s.FileSize = &v
  return s
}

func (s *HttpTestResponseResultData) SetRate(v float32) *HttpTestResponseResultData {
  s.Rate = &v
  return s
}

func (s *HttpTestResponseResultData) SetHttpCode(v string) *HttpTestResponseResultData {
  s.HttpCode = &v
  return s
}

func (s *HttpTestResponseResultData) SetHttpTestRequestHeader(v string) *HttpTestResponseResultData {
  s.HttpTestRequestHeader = &v
  return s
}

func (s *HttpTestResponseResultData) SetHttpTestResponseHeader(v string) *HttpTestResponseResultData {
  s.HttpTestResponseHeader = &v
  return s
}

func (s *HttpTestResponseResultData) SetStatus(v int32) *HttpTestResponseResultData {
  s.Status = &v
  return s
}

type HttpTestPaths struct {
}

func (s HttpTestPaths) String() string {
  return tea.Prettify(s)
}

func (s HttpTestPaths) GoString() string {
  return s.String()
}

type HttpTestParameters struct {
}

func (s HttpTestParameters) String() string {
  return tea.Prettify(s)
}

func (s HttpTestParameters) GoString() string {
  return s.String()
}

type HttpTestRequestHeader struct {
}

func (s HttpTestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s HttpTestRequestHeader) GoString() string {
  return s.String()
}

type HttpTestResponseHeader struct {
}

func (s HttpTestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s HttpTestResponseHeader) GoString() string {
  return s.String()
}




type Query5minLiveTranscodingDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"This parameter specifies if the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"acceleration type:
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way.
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"Transcoding type, values can be h264, h265, zdgq_264, zdgq_265, cf_264, cf_265, or other. Multiple transcoding types should be separated by a semicolon. If some of the transcoding types are incorrect, the system will return data for the correct types; if all transcoding types are incorrect, it will return an error 'invalid transcodeType.' If not provided or left empty, it defaults to all types.", "zh_CN":"转码类型,值为h264、h265、zdgq_264、zdgq_265、cf_264、cf_265，other 多个转码类型用英文分号;分隔开。当传入转码类型部分错误时，返回正确的类型的数据；当传入转码类型全部错误时，返回错误invalid transcodeType. 不填或为空，默认为所有类型."}
  TranscodeType *string `json:"transcodeType,omitempty" xml:"transcodeType,omitempty"`
  // {"en":"Resolution types include LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240. Multiple resolutions are separated by a semicolon. When isAudio=1, this parameter is invalid and will return an error. Param definition must be empty when querying audio data.", "zh_CN":"清晰度类型,值为LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240，多个清晰度用英文分号;分隔开, 当isAudio=1时，此入参无效返回错误 param definition must be empty when query audio data."}
  Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
  // {"en":"Audio/Video Type, 1: Audio 2: Video. Defaults to 2 if not selected or empty. Only a single value is allowed.", "zh_CN":"音视频类型, 1:音频   2:视频. 不选或者为空时默认为2. 只能输入单个值."}
  IsAudio *string `json:"isAudio,omitempty" xml:"isAudio,omitempty"`
}

func (s Query5minLiveTranscodingDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationRequest) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationRequest) SetCust(v string) *Query5minLiveTranscodingDurationRequest {
  s.Cust = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetDate(v string) *Query5minLiveTranscodingDurationRequest {
  s.Date = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetStartdate(v string) *Query5minLiveTranscodingDurationRequest {
  s.Startdate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetEnddate(v string) *Query5minLiveTranscodingDurationRequest {
  s.Enddate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetTimezone(v string) *Query5minLiveTranscodingDurationRequest {
  s.Timezone = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetChannel(v string) *Query5minLiveTranscodingDurationRequest {
  s.Channel = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetIsExactMatch(v string) *Query5minLiveTranscodingDurationRequest {
  s.IsExactMatch = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetAccetype(v string) *Query5minLiveTranscodingDurationRequest {
  s.Accetype = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetDataformat(v string) *Query5minLiveTranscodingDurationRequest {
  s.Dataformat = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetResultType(v string) *Query5minLiveTranscodingDurationRequest {
  s.ResultType = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetTranscodeType(v string) *Query5minLiveTranscodingDurationRequest {
  s.TranscodeType = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetDefinition(v string) *Query5minLiveTranscodingDurationRequest {
  s.Definition = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetIsAudio(v string) *Query5minLiveTranscodingDurationRequest {
  s.IsAudio = &v
  return s
}

type Query5minLiveTranscodingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *Query5minLiveTranscodingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponse) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponse) SetProvider(v *Query5minLiveTranscodingDurationResponseProvider) *Query5minLiveTranscodingDurationResponse {
  s.Provider = v
  return s
}

type Query5minLiveTranscodingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'直播转码时长数据'}
  Date *Query5minLiveTranscodingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProvider) SetName(v string) *Query5minLiveTranscodingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProvider) SetType(v string) *Query5minLiveTranscodingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProvider) SetDate(v *Query5minLiveTranscodingDurationResponseProviderDate) *Query5minLiveTranscodingDurationResponseProvider {
  s.Date = v
  return s
}

type Query5minLiveTranscodingDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *Query5minLiveTranscodingDurationResponseProviderDateTranscoding `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProviderDate) SetStartdate(v string) *Query5minLiveTranscodingDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDate) SetEnddate(v string) *Query5minLiveTranscodingDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDate) SetTranscoding(v *Query5minLiveTranscodingDurationResponseProviderDateTranscoding) *Query5minLiveTranscodingDurationResponseProviderDate {
  s.Transcoding = v
  return s
}

type Query5minLiveTranscodingDurationResponseProviderDateTranscoding struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'live', 'zh_CN':'直播转码时长数据'}
  Live []*Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscoding) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscoding) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscoding) SetName(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Name = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscoding) SetLive(v []*Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) *Query5minLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Live = v
  return s
}

type Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'h264 transcoding time', 'zh_CN':'h264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H264 *string `json:"h264,omitempty" xml:"h264,omitempty" require:"true"`
  // {'en':'h265 transcoding time', 'zh_CN':'h265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H265 *string `json:"h265,omitempty" xml:"h265,omitempty" require:"true"`
  // {'en':'zdgq_264 transcoding time', 'zh_CN':'zdgq_264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_264 *string `json:"zdgq_264,omitempty" xml:"zdgq_264,omitempty" require:"true"`
  // {'en':'zdgq_265 transcoding time', 'zh_CN':'zdgq_265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_265 *string `json:"zdgq_265,omitempty" xml:"zdgq_265,omitempty" require:"true"`
  // {'en':'voice transcoding time', 'zh_CN':'voice转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty" require:"true"`
  // {'en':'cf_264 transcoding time', 'zh_CN':'cf_264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Cf_264 *string `json:"cf_264,omitempty" xml:"cf_264,omitempty" require:"true"`
  // {'en':'cf_265 transcoding time', 'zh_CN':'cf_265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Cf_265 *string `json:"cf_265,omitempty" xml:"cf_265,omitempty" require:"true"`
  // {'en':'definition_2K transcoding time', 'zh_CN':'definition_2K转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_2K *string `json:"definition_2K,omitempty" xml:"definition_2K,omitempty" require:"true"`
  // {'en':'definition_4K transcoding time', 'zh_CN':'definition_4K转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_4K *string `json:"definition_4K,omitempty" xml:"definition_4K,omitempty" require:"true"`
  // {'en':'definition_8K transcoding time', 'zh_CN':'definition_8K转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_8K *string `json:"definition_8K,omitempty" xml:"definition_8K,omitempty" require:"true"`
  // {'en':'definition_LD480 transcoding time', 'zh_CN':'definition_LD480转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_LD480 *string `json:"definition_LD480,omitempty" xml:"definition_LD480,omitempty" require:"true"`
  // {'en':'definition_SD720 transcoding time', 'zh_CN':'definition_SD720转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_SD720 *string `json:"definition_SD720,omitempty" xml:"definition_SD720,omitempty" require:"true"`
  // {'en':'definition_HD1080 transcoding time', 'zh_CN':'definition_HD1080转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_HD1080 *string `json:"definition_HD1080,omitempty" xml:"definition_HD1080,omitempty" require:"true"`
  // {'en':'definition_SD576 transcoding time', 'zh_CN':'definition_SD576转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_SD576 *string `json:"definition_SD576,omitempty" xml:"definition_SD576,omitempty" require:"true"`
  // {'en':'total transcoding time', 'zh_CN':'总转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTime(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Time = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH264(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH265(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_264(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_265(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetVoice(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Voice = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetCf_264(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Cf_264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetCf_265(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Cf_265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_2K(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_2K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_4K(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_4K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_8K(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_8K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_LD480(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_LD480 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_SD720(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_SD720 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_HD1080(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_HD1080 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_SD576(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_SD576 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTotal(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Total = &v
  return s
}

type Query5minLiveTranscodingDurationPaths struct {
}

func (s Query5minLiveTranscodingDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationPaths) GoString() string {
  return s.String()
}

type Query5minLiveTranscodingDurationParameters struct {
}

func (s Query5minLiveTranscodingDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationParameters) GoString() string {
  return s.String()
}

type Query5minLiveTranscodingDurationRequestHeader struct {
}

func (s Query5minLiveTranscodingDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationRequestHeader) GoString() string {
  return s.String()
}

type Query5minLiveTranscodingDurationResponseHeader struct {
}

func (s Query5minLiveTranscodingDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseHeader) GoString() string {
  return s.String()
}




type PicProcessStatisticsRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s PicProcessStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsRequest) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsRequest) SetCust(v string) *PicProcessStatisticsRequest {
  s.Cust = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetDate(v string) *PicProcessStatisticsRequest {
  s.Date = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetStartdate(v string) *PicProcessStatisticsRequest {
  s.Startdate = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetEnddate(v string) *PicProcessStatisticsRequest {
  s.Enddate = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetChannel(v string) *PicProcessStatisticsRequest {
  s.Channel = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetRegion(v string) *PicProcessStatisticsRequest {
  s.Region = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetDataformat(v string) *PicProcessStatisticsRequest {
  s.Dataformat = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetResultType(v string) *PicProcessStatisticsRequest {
  s.ResultType = &v
  return s
}

type PicProcessStatisticsResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *PicProcessStatisticsResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponse) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponse) SetProvider(v *PicProcessStatisticsResponseProvider) *PicProcessStatisticsResponse {
  s.Provider = v
  return s
}

type PicProcessStatisticsResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Date *PicProcessStatisticsResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProvider) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProvider) SetName(v string) *PicProcessStatisticsResponseProvider {
  s.Name = &v
  return s
}

func (s *PicProcessStatisticsResponseProvider) SetType(v string) *PicProcessStatisticsResponseProvider {
  s.Type = &v
  return s
}

func (s *PicProcessStatisticsResponseProvider) SetResultType(v string) *PicProcessStatisticsResponseProvider {
  s.ResultType = &v
  return s
}

func (s *PicProcessStatisticsResponseProvider) SetDate(v *PicProcessStatisticsResponseProviderDate) *PicProcessStatisticsResponseProvider {
  s.Date = v
  return s
}

type PicProcessStatisticsResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *PicProcessStatisticsResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProviderDate) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProviderDate) SetStartdate(v string) *PicProcessStatisticsResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDate) SetEnddate(v string) *PicProcessStatisticsResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDate) SetChannel(v *PicProcessStatisticsResponseProviderDateChannel) *PicProcessStatisticsResponseProviderDate {
  s.Channel = v
  return s
}

type PicProcessStatisticsResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'picflowhit', 'zh_CN':'请求数数据'}
  Picflowhit []*PicProcessStatisticsResponseProviderDateChannelPicflowhit `json:"picflowhit,omitempty" xml:"picflowhit,omitempty" require:"true" type:"Repeated"`
}

func (s PicProcessStatisticsResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProviderDateChannel) SetName(v string) *PicProcessStatisticsResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDateChannel) SetPicflowhit(v []*PicProcessStatisticsResponseProviderDateChannelPicflowhit) *PicProcessStatisticsResponseProviderDateChannel {
  s.Picflowhit = v
  return s
}

type PicProcessStatisticsResponseProviderDateChannelPicflowhit struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'请求数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s PicProcessStatisticsResponseProviderDateChannelPicflowhit) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProviderDateChannelPicflowhit) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProviderDateChannelPicflowhit) SetTime(v string) *PicProcessStatisticsResponseProviderDateChannelPicflowhit {
  s.Time = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDateChannelPicflowhit) SetText(v string) *PicProcessStatisticsResponseProviderDateChannelPicflowhit {
  s.Text = &v
  return s
}

type PicProcessStatisticsPaths struct {
}

func (s PicProcessStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsPaths) GoString() string {
  return s.String()
}

type PicProcessStatisticsParameters struct {
}

func (s PicProcessStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsParameters) GoString() string {
  return s.String()
}

type PicProcessStatisticsRequestHeader struct {
}

func (s PicProcessStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsRequestHeader) GoString() string {
  return s.String()
}

type PicProcessStatisticsResponseHeader struct {
}

func (s PicProcessStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainStreamHlsOnlineServiceRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 10minutes will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 1 hour, means the period between dateFrom to dateTo should not exceed 1 hour (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的10分钟;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔1小时:即dateFrom和dateTo相差不能超过1小时。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain&stream group:
  // 1.Allowable maximum number of domain&stream group is 20 (can be adjusted by contacting technical support).
  // 2.Domain: a group of domain&stream only can include one domain,and the domain is required
  // 3.Stream:There is no limit on the number of stream names under a group of domain&stream. If the stream is not transmitted, all stream under the domain will be queried by default", "zh_CN":"域名和流名组:
  // 1.可传递的域名和流名组数量上限默认为20组(可联系技术支持调整);
  // 2.域名domain:一组域名和流名组中只能传递单个域名,且域名必须传递;
  // 3.流名stream:一组域名和流名组下(即单个域名下)不限制流名个数,流名未传递时默认查询域名下所有流名"}
  DomainStream []*ReportDomainStreamHlsOnlineServiceRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamHlsOnlineServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceRequest) SetDateFrom(v string) *ReportDomainStreamHlsOnlineServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceRequest) SetDateTo(v string) *ReportDomainStreamHlsOnlineServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceRequest) SetDomainStream(v []*ReportDomainStreamHlsOnlineServiceRequestDomainStream) *ReportDomainStreamHlsOnlineServiceRequest {
  s.DomainStream = v
  return s
}

type ReportDomainStreamHlsOnlineServiceRequestDomainStream struct     {
  // {"en":"domian", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"stream: 'publishing point'/'stream'; 
  // 					For example: live/test-20200101-test.flv ,'live' is the publishing point, 'test-20200101-test' is stream;
  // 					If the stream is not transmitted, all stream under the domain will be queried by default", "zh_CN":"流名:发布点/流名。
  // 					例如:live/test-20200101-test.flv ,其中live是发布点, test-20200101-test是流名;
  // 					不传,默认查询指定域名下的所有流的数据"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamHlsOnlineServiceRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceRequestDomainStream) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceRequestDomainStream) SetDomain(v string) *ReportDomainStreamHlsOnlineServiceRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceRequestDomainStream) SetStream(v []*string) *ReportDomainStreamHlsOnlineServiceRequestDomainStream {
  s.Stream = v
  return s
}

type ReportDomainStreamHlsOnlineServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportDomainStreamHlsOnlineServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamHlsOnlineServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceResponse) SetCode(v string) *ReportDomainStreamHlsOnlineServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceResponse) SetMessage(v string) *ReportDomainStreamHlsOnlineServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceResponse) SetData(v []*ReportDomainStreamHlsOnlineServiceResponseData) *ReportDomainStreamHlsOnlineServiceResponse {
  s.Data = v
  return s
}

type ReportDomainStreamHlsOnlineServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Number of streams under domain name", "zh_CN":"域名下的流个数"}
  StreamCount *string `json:"streamCount,omitempty" xml:"streamCount,omitempty" require:"true"`
  // {"en":"The number of online people corresponding to the domain.The value is the cumulative number of online people of all stream under the domain", "zh_CN":"该频道下总的在线人数,值为频道下所有流名的在线人数累加"}
  TotalOnlineCount *string `json:"totalOnlineCount,omitempty" xml:"totalOnlineCount,omitempty" require:"true"`
  StreamDetails []*ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails `json:"streamDetails,omitempty" xml:"streamDetails,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamHlsOnlineServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceResponseData) SetDomain(v string) *ReportDomainStreamHlsOnlineServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceResponseData) SetStreamCount(v string) *ReportDomainStreamHlsOnlineServiceResponseData {
  s.StreamCount = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceResponseData) SetTotalOnlineCount(v string) *ReportDomainStreamHlsOnlineServiceResponseData {
  s.TotalOnlineCount = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceResponseData) SetStreamDetails(v []*ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) *ReportDomainStreamHlsOnlineServiceResponseData {
  s.StreamDetails = v
  return s
}

type ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails struct     {
  // {"en":"stream", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"The number of online people corresponding to the stream", "zh_CN":"该流名对应的在线人数"}
  OnlineCount *string `json:"onlineCount,omitempty" xml:"onlineCount,omitempty" require:"true"`
}

func (s ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) SetStream(v string) *ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails {
  s.Stream = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) SetOnlineCount(v string) *ReportDomainStreamHlsOnlineServiceResponseDataStreamDetails {
  s.OnlineCount = &v
  return s
}

type ReportDomainStreamHlsOnlineServicePaths struct {
}

func (s ReportDomainStreamHlsOnlineServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServicePaths) GoString() string {
  return s.String()
}

type ReportDomainStreamHlsOnlineServiceParameters struct {
}

func (s ReportDomainStreamHlsOnlineServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainStreamHlsOnlineServiceRequestHeader struct {
}

func (s ReportDomainStreamHlsOnlineServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainStreamHlsOnlineServiceResponseHeader struct {
}

func (s ReportDomainStreamHlsOnlineServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceResponseHeader) GoString() string {
  return s.String()
}




type StreamTrafficServiceRequest struct {
  // {"en":"Domain(s),  multiple domains are separated by commas,up to a maximum of 5 supported","zh_CN":"域名，可多个，英文逗号分隔，最多可支持5个"}
  DomainList *string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true"`
  // {"en":"Region(s), multiple regions are separated by commas","zh_CN":"区域，可多个，英文逗号分隔"}
  RegionList *string `json:"regionList,omitempty" xml:"regionList,omitempty"`
  // {"en":"Stream(s), multiple allowed, separated by commas","zh_CN":"流名，可多个，英文逗号分隔"}
  StreamNameList *string `json:"streamNameList,omitempty" xml:"streamNameList,omitempty"`
  // {"en":"App(s), multiple, separated by commas","zh_CN":"发布点，可多个，英文逗号分隔"}
  AppList *string `json:"appList,omitempty" xml:"appList,omitempty"`
  // {"en":"Query start time. The format is yyyy-MM-ddTHH:mm:ss+08:00; for example, 2024-12-12T10:00:00+08:00 (which is 10:00 AM Beijing time on December 12, 2024). You can query data for up to 1 day.","zh_CN":"查询开始时间。格式为yyyy-MM-ddTHH:mm:ss+08:00；例如，2024-12-12T10:00:00+08:00（为北京时间2024年12月12日10点0分0秒）;最多查1天数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"Query end time. The format is yyyy-MM-ddTHH:mm:ss+08:00","zh_CN":"查询结束时间。格式为yyyy-MM-ddTHH:mm:ss+08:00"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Value: rtmp, hdl, hls, rtc, srt, other; multiple values allowed, separated by commas.","zh_CN":"值：rtmp，hdl，hls，rtc，srt，other  可多个，英文逗号分隔"}
  ProtocolList *string `json:"protocolList,omitempty" xml:"protocolList,omitempty"`
}

func (s StreamTrafficServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceRequest) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceRequest) SetDomainList(v string) *StreamTrafficServiceRequest {
  s.DomainList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetRegionList(v string) *StreamTrafficServiceRequest {
  s.RegionList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetStreamNameList(v string) *StreamTrafficServiceRequest {
  s.StreamNameList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetAppList(v string) *StreamTrafficServiceRequest {
  s.AppList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetDateFrom(v string) *StreamTrafficServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetDateTo(v string) *StreamTrafficServiceRequest {
  s.DateTo = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetProtocolList(v string) *StreamTrafficServiceRequest {
  s.ProtocolList = &v
  return s
}

type StreamTrafficServiceRequestHeader struct {
}

func (s StreamTrafficServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceRequestHeader) GoString() string {
  return s.String()
}

type StreamTrafficServicePaths struct {
}

func (s StreamTrafficServicePaths) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServicePaths) GoString() string {
  return s.String()
}

type StreamTrafficServiceParameters struct {
}

func (s StreamTrafficServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceParameters) GoString() string {
  return s.String()
}

type StreamTrafficServiceResponse struct {
  // {"en":"requestId","zh_CN":"请求id"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"Query start time","zh_CN":"查询开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Query end time","zh_CN":"查询结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"regions","zh_CN":"区域"}
  Regions []*StreamTrafficServiceResponseRegions `json:"regions,omitempty" xml:"regions,omitempty" require:"true" type:"Repeated"`
}

func (s StreamTrafficServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponse) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponse) SetRequestId(v string) *StreamTrafficServiceResponse {
  s.RequestId = &v
  return s
}

func (s *StreamTrafficServiceResponse) SetStartTime(v string) *StreamTrafficServiceResponse {
  s.StartTime = &v
  return s
}

func (s *StreamTrafficServiceResponse) SetEndTime(v string) *StreamTrafficServiceResponse {
  s.EndTime = &v
  return s
}

func (s *StreamTrafficServiceResponse) SetRegions(v []*StreamTrafficServiceResponseRegions) *StreamTrafficServiceResponse {
  s.Regions = v
  return s
}

type StreamTrafficServiceResponseRegions struct     {
  // {"en":"regionCode","zh_CN":"区域编码"}
  RegionCode *string `json:"regionCode,omitempty" xml:"regionCode,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Domains []*StreamTrafficServiceResponseRegionsDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s StreamTrafficServiceResponseRegions) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseRegions) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponseRegions) SetRegionCode(v string) *StreamTrafficServiceResponseRegions {
  s.RegionCode = &v
  return s
}

func (s *StreamTrafficServiceResponseRegions) SetDomains(v []*StreamTrafficServiceResponseRegionsDomains) *StreamTrafficServiceResponseRegions {
  s.Domains = v
  return s
}

type StreamTrafficServiceResponseRegionsDomains struct     {
  // {"en":"domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Streams []*StreamTrafficServiceResponseRegionsDomainsStreams `json:"streams,omitempty" xml:"streams,omitempty" require:"true" type:"Repeated"`
}

func (s StreamTrafficServiceResponseRegionsDomains) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseRegionsDomains) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponseRegionsDomains) SetDomain(v string) *StreamTrafficServiceResponseRegionsDomains {
  s.Domain = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomains) SetStreams(v []*StreamTrafficServiceResponseRegionsDomainsStreams) *StreamTrafficServiceResponseRegionsDomains {
  s.Streams = v
  return s
}

type StreamTrafficServiceResponseRegionsDomainsStreams struct     {
  // {"en":"streamName","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty" require:"true"`
  // {"en":"Application Name","zh_CN":"发布点"}
  App *string `json:"app,omitempty" xml:"app,omitempty" require:"true"`
  // {"en":"protocol","zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"flow","zh_CN":"流量"}
  Traffic *int64 `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s StreamTrafficServiceResponseRegionsDomainsStreams) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseRegionsDomainsStreams) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetStreamName(v string) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.StreamName = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetApp(v string) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.App = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetProtocol(v string) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.Protocol = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetTraffic(v int64) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.Traffic = &v
  return s
}

type StreamTrafficServiceResponseHeader struct {
}

func (s StreamTrafficServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainListExistFlowServiceRequest struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 1 days;4.You can only query data for the last 2 years.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过1天;
  // 4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
}

func (s ReportDomainListExistFlowServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainListExistFlowServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainListExistFlowServiceRequest) SetDateFrom(v string) *ReportDomainListExistFlowServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainListExistFlowServiceRequest) SetDateTo(v string) *ReportDomainListExistFlowServiceRequest {
  s.DateTo = &v
  return s
}

type ReportDomainListExistFlowServiceResponse struct {
  // {"en":"Number of domains that have generated traffic", "zh_CN":"有量的域名数量"}
  DomainNum *int `json:"domainNum,omitempty" xml:"domainNum,omitempty" require:"true"`
  // {"en":"List of the domains that have generated traffic", "zh_CN":"有量的域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainListExistFlowServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainListExistFlowServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainListExistFlowServiceResponse) SetDomainNum(v int) *ReportDomainListExistFlowServiceResponse {
  s.DomainNum = &v
  return s
}

func (s *ReportDomainListExistFlowServiceResponse) SetDomainList(v []*string) *ReportDomainListExistFlowServiceResponse {
  s.DomainList = v
  return s
}

type ReportDomainListExistFlowServicePaths struct {
}

func (s ReportDomainListExistFlowServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainListExistFlowServicePaths) GoString() string {
  return s.String()
}

type ReportDomainListExistFlowServiceParameters struct {
}

func (s ReportDomainListExistFlowServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainListExistFlowServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainListExistFlowServiceRequestHeader struct {
}

func (s ReportDomainListExistFlowServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainListExistFlowServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainListExistFlowServiceResponseHeader struct {
}

func (s ReportDomainListExistFlowServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainListExistFlowServiceResponseHeader) GoString() string {
  return s.String()
}




type SubmitXlwInjectTaskRequest struct {
  // {"en":"CDN manufacturer authorized username","zh_CN":"cdn厂家授权用户名"}
  Username *string `json:"username,omitempty" xml:"username,omitempty" require:"true"`
  // {"en":"Password authorized by the CDN manufacturer","zh_CN":"cdn厂家授权用户名密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty" require:"true"`
  // {"en":"The injected content supports submitting multiple injection tasks simultaneously.","zh_CN":"注入的内容 支持同时提交多条注入任务"}
  FcSub []*SubmitXlwInjectTaskRequestFcSub `json:"fc_sub,omitempty" xml:"fc_sub,omitempty" require:"true" type:"Repeated"`
}

func (s SubmitXlwInjectTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskRequest) GoString() string {
  return s.String()
}

func (s *SubmitXlwInjectTaskRequest) SetUsername(v string) *SubmitXlwInjectTaskRequest {
  s.Username = &v
  return s
}

func (s *SubmitXlwInjectTaskRequest) SetPassword(v string) *SubmitXlwInjectTaskRequest {
  s.Password = &v
  return s
}

func (s *SubmitXlwInjectTaskRequest) SetFcSub(v []*SubmitXlwInjectTaskRequestFcSub) *SubmitXlwInjectTaskRequest {
  s.FcSub = v
  return s
}

type SubmitXlwInjectTaskRequestFcSub struct     {
  // {"en":"The unique identifier of the video","zh_CN":"视频的唯一标识"}
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
  // {"en":"Operation instructions","zh_CN":"操作指令"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty" require:"true"`
  // {"en":"File source URL","zh_CN":"文件源 url"}
  SourcePath *string `json:"source_path,omitempty" xml:"source_path,omitempty" require:"true"`
  // {"en":"File distribution path","zh_CN":"文件发布路径"}
  PublishPath *string `json:"publish_path,omitempty" xml:"publish_path,omitempty" require:"true"`
  // {"en":"1 indicates sectioning, 0 indicates no sectioning","zh_CN":"1标示切片，0标示不切片"}
  Slice *string `json:"slice,omitempty" xml:"slice,omitempty" require:"true"`
}

func (s SubmitXlwInjectTaskRequestFcSub) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskRequestFcSub) GoString() string {
  return s.String()
}

func (s *SubmitXlwInjectTaskRequestFcSub) SetItemId(v string) *SubmitXlwInjectTaskRequestFcSub {
  s.ItemId = &v
  return s
}

func (s *SubmitXlwInjectTaskRequestFcSub) SetOperation(v string) *SubmitXlwInjectTaskRequestFcSub {
  s.Operation = &v
  return s
}

func (s *SubmitXlwInjectTaskRequestFcSub) SetSourcePath(v string) *SubmitXlwInjectTaskRequestFcSub {
  s.SourcePath = &v
  return s
}

func (s *SubmitXlwInjectTaskRequestFcSub) SetPublishPath(v string) *SubmitXlwInjectTaskRequestFcSub {
  s.PublishPath = &v
  return s
}

func (s *SubmitXlwInjectTaskRequestFcSub) SetSlice(v string) *SubmitXlwInjectTaskRequestFcSub {
  s.Slice = &v
  return s
}

type SubmitXlwInjectTaskRequestHeader struct {
}

func (s SubmitXlwInjectTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskRequestHeader) GoString() string {
  return s.String()
}

type SubmitXlwInjectTaskPaths struct {
}

func (s SubmitXlwInjectTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskPaths) GoString() string {
  return s.String()
}

type SubmitXlwInjectTaskParameters struct {
}

func (s SubmitXlwInjectTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskParameters) GoString() string {
  return s.String()
}

type SubmitXlwInjectTaskResponse struct {
  // {"en":"Operation success or failure status: 0 for Failure, 1 for Success","zh_CN":"操作是否成功 0 失败 1成功"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Feedback information","zh_CN":"反馈信息"}
  Info *string `json:"info,omitempty" xml:"info,omitempty" require:"true"`
  // {"en":"Returns the result; if all results are successful, it returns true.","zh_CN":"返回结果内容，如果全都成功了 直接返回true"}
  Result *SubmitXlwInjectTaskResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s SubmitXlwInjectTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskResponse) GoString() string {
  return s.String()
}

func (s *SubmitXlwInjectTaskResponse) SetStatus(v int) *SubmitXlwInjectTaskResponse {
  s.Status = &v
  return s
}

func (s *SubmitXlwInjectTaskResponse) SetInfo(v string) *SubmitXlwInjectTaskResponse {
  s.Info = &v
  return s
}

func (s *SubmitXlwInjectTaskResponse) SetResult(v *SubmitXlwInjectTaskResponseResult) *SubmitXlwInjectTaskResponse {
  s.Result = v
  return s
}

type SubmitXlwInjectTaskResponseResult struct {
  // {"en":"Success data","zh_CN":"成功数据"}
  Finish []*SubmitXlwInjectTaskResponseResultFinish `json:"finish,omitempty" xml:"finish,omitempty" require:"true" type:"Repeated"`
  // {"en":"Failure data","zh_CN":"失败数据"}
  Failed []*SubmitXlwInjectTaskResponseResultFailed `json:"failed,omitempty" xml:"failed,omitempty" require:"true" type:"Repeated"`
}

func (s SubmitXlwInjectTaskResponseResult) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskResponseResult) GoString() string {
  return s.String()
}

func (s *SubmitXlwInjectTaskResponseResult) SetFinish(v []*SubmitXlwInjectTaskResponseResultFinish) *SubmitXlwInjectTaskResponseResult {
  s.Finish = v
  return s
}

func (s *SubmitXlwInjectTaskResponseResult) SetFailed(v []*SubmitXlwInjectTaskResponseResultFailed) *SubmitXlwInjectTaskResponseResult {
  s.Failed = v
  return s
}

type SubmitXlwInjectTaskResponseResultFinish struct     {
  // {"en":"Video ID","zh_CN":"视频id"}
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
}

func (s SubmitXlwInjectTaskResponseResultFinish) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskResponseResultFinish) GoString() string {
  return s.String()
}

func (s *SubmitXlwInjectTaskResponseResultFinish) SetItemId(v string) *SubmitXlwInjectTaskResponseResultFinish {
  s.ItemId = &v
  return s
}

type SubmitXlwInjectTaskResponseResultFailed struct     {
  // {"en":"Video ID","zh_CN":"视频id"}
  ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
  // {"en":"Failure message","zh_CN":"失败信息"}
  FailedInfo *string `json:"failed_info,omitempty" xml:"failed_info,omitempty" require:"true"`
}

func (s SubmitXlwInjectTaskResponseResultFailed) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskResponseResultFailed) GoString() string {
  return s.String()
}

func (s *SubmitXlwInjectTaskResponseResultFailed) SetItemId(v string) *SubmitXlwInjectTaskResponseResultFailed {
  s.ItemId = &v
  return s
}

func (s *SubmitXlwInjectTaskResponseResultFailed) SetFailedInfo(v string) *SubmitXlwInjectTaskResponseResultFailed {
  s.FailedInfo = &v
  return s
}

type SubmitXlwInjectTaskResponseHeader struct {
}

func (s SubmitXlwInjectTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SubmitXlwInjectTaskResponseHeader) GoString() string {
  return s.String()
}




type ReportStreamListServiceRequest struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 3 days;4.You can only query data for the last 2 years.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过3天；
  // 4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain, must follow regular expression rule of (([\w-]{1,62})?(\.[\w-]{1,62})+)", "zh_CN":"域名，需要符合正则(([\w-]{1,62})?(\.[\w-]{1,62})+)"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s ReportStreamListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStreamListServiceRequest) SetDateFrom(v string) *ReportStreamListServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStreamListServiceRequest) SetDateTo(v string) *ReportStreamListServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStreamListServiceRequest) SetDomain(v string) *ReportStreamListServiceRequest {
  s.Domain = &v
  return s
}

type ReportStreamListServiceResponse struct {
  // {"en":"", "zh_CN":""}
  Result []*ReportStreamListServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStreamListServiceResponse) SetResult(v []*ReportStreamListServiceResponseResult) *ReportStreamListServiceResponse {
  s.Result = v
  return s
}

type ReportStreamListServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Domain list", "zh_CN":"流名列表"}
  StreamList []*string `json:"streamList,omitempty" xml:"streamList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamListServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportStreamListServiceResponseResult) SetDomain(v string) *ReportStreamListServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportStreamListServiceResponseResult) SetStreamList(v []*string) *ReportStreamListServiceResponseResult {
  s.StreamList = v
  return s
}

type ReportStreamListServicePaths struct {
}

func (s ReportStreamListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServicePaths) GoString() string {
  return s.String()
}

type ReportStreamListServiceParameters struct {
}

func (s ReportStreamListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceParameters) GoString() string {
  return s.String()
}

type ReportStreamListServiceRequestHeader struct {
}

func (s ReportStreamListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStreamListServiceResponseHeader struct {
}

func (s ReportStreamListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainResourceGroupRequest struct {
  // {"en":"Domain list, up to 100 domains","zh_CN":"域名列表，最多100个域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainResourceGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupRequest) SetDomainList(v []*string) *QueryDomainResourceGroupRequest {
  s.DomainList = v
  return s
}

type QueryDomainResourceGroupRequestHeader struct {
}

func (s QueryDomainResourceGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainResourceGroupPaths struct {
}

func (s QueryDomainResourceGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupPaths) GoString() string {
  return s.String()
}

type QueryDomainResourceGroupParameters struct {
}

func (s QueryDomainResourceGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupParameters) GoString() string {
  return s.String()
}

type QueryDomainResourceGroupResponse struct {
  // {"en":"Data","zh_CN":"数据"}
  Data *QueryDomainResourceGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Status code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s QueryDomainResourceGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupResponse) SetData(v *QueryDomainResourceGroupResponseData) *QueryDomainResourceGroupResponse {
  s.Data = v
  return s
}

func (s *QueryDomainResourceGroupResponse) SetCode(v int) *QueryDomainResourceGroupResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainResourceGroupResponse) SetMsg(v string) *QueryDomainResourceGroupResponse {
  s.Msg = &v
  return s
}

type QueryDomainResourceGroupResponseData struct {
  // {"en":"List of domain associated resource group information","zh_CN":"域名关联资源组信息列表"}
  InfoList []*QueryDomainResourceGroupResponseDataInfoList `json:"infoList,omitempty" xml:"infoList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainResourceGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponseData) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupResponseData) SetInfoList(v []*QueryDomainResourceGroupResponseDataInfoList) *QueryDomainResourceGroupResponseData {
  s.InfoList = v
  return s
}

type QueryDomainResourceGroupResponseDataInfoList struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Resource group name","zh_CN":"资源组名称"}
  Resource *string `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
  // {"en":"Resource group ID","zh_CN":"资源组Id"}
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty" require:"true"`
}

func (s QueryDomainResourceGroupResponseDataInfoList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponseDataInfoList) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupResponseDataInfoList) SetDomain(v string) *QueryDomainResourceGroupResponseDataInfoList {
  s.Domain = &v
  return s
}

func (s *QueryDomainResourceGroupResponseDataInfoList) SetResource(v string) *QueryDomainResourceGroupResponseDataInfoList {
  s.Resource = &v
  return s
}

func (s *QueryDomainResourceGroupResponseDataInfoList) SetResourceGroupId(v string) *QueryDomainResourceGroupResponseDataInfoList {
  s.ResourceGroupId = &v
  return s
}

type QueryDomainResourceGroupResponseHeader struct {
}

func (s QueryDomainResourceGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponseHeader) GoString() string {
  return s.String()
}




type QueryLiveRecordingDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way
  // 1.If specified 1, get the merged result.
  // 2.If  specified 2,get the separate result.
  // 3.If specifed 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s QueryLiveRecordingDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationRequest) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationRequest) SetCust(v string) *QueryLiveRecordingDurationRequest {
  s.Cust = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetDate(v string) *QueryLiveRecordingDurationRequest {
  s.Date = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetStartdate(v string) *QueryLiveRecordingDurationRequest {
  s.Startdate = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetEnddate(v string) *QueryLiveRecordingDurationRequest {
  s.Enddate = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetTimezone(v string) *QueryLiveRecordingDurationRequest {
  s.Timezone = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetChannel(v string) *QueryLiveRecordingDurationRequest {
  s.Channel = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetRegion(v string) *QueryLiveRecordingDurationRequest {
  s.Region = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetAccetype(v string) *QueryLiveRecordingDurationRequest {
  s.Accetype = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetDataformat(v string) *QueryLiveRecordingDurationRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetResultType(v string) *QueryLiveRecordingDurationRequest {
  s.ResultType = &v
  return s
}

type QueryLiveRecordingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryLiveRecordingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponse) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponse) SetProvider(v *QueryLiveRecordingDurationResponseProvider) *QueryLiveRecordingDurationResponse {
  s.Provider = v
  return s
}

type QueryLiveRecordingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'请明细数据'}
  Date *QueryLiveRecordingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProvider) SetName(v string) *QueryLiveRecordingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProvider) SetType(v string) *QueryLiveRecordingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProvider) SetResultType(v string) *QueryLiveRecordingDurationResponseProvider {
  s.ResultType = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProvider) SetDate(v *QueryLiveRecordingDurationResponseProviderDate) *QueryLiveRecordingDurationResponseProvider {
  s.Date = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDate struct {
  // {'en':'start', 'zh_CN':'开始时间'}
  Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {'en':'end', 'zh_CN':'结束时间'}
  End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  RecordingTime *QueryLiveRecordingDurationResponseProviderDateRecordingTime `json:"recordingTime,omitempty" xml:"recordingTime,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDate) SetStart(v string) *QueryLiveRecordingDurationResponseProviderDate {
  s.Start = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDate) SetEnd(v string) *QueryLiveRecordingDurationResponseProviderDate {
  s.End = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDate) SetRecordingTime(v *QueryLiveRecordingDurationResponseProviderDateRecordingTime) *QueryLiveRecordingDurationResponseProviderDate {
  s.RecordingTime = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDateRecordingTime struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Channel *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTime) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTime) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTime) SetName(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTime {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTime) SetChannel(v *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) *QueryLiveRecordingDurationResponseProviderDateRecordingTime {
  s.Channel = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'明细数据'}
  Live []*QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) SetName(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) SetLive(v []*QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel {
  s.Live = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数据'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) SetTime(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive {
  s.Time = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) SetText(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive {
  s.Text = &v
  return s
}

type QueryLiveRecordingDurationPaths struct {
}

func (s QueryLiveRecordingDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationPaths) GoString() string {
  return s.String()
}

type QueryLiveRecordingDurationParameters struct {
}

func (s QueryLiveRecordingDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationParameters) GoString() string {
  return s.String()
}

type QueryLiveRecordingDurationRequestHeader struct {
}

func (s QueryLiveRecordingDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationRequestHeader) GoString() string {
  return s.String()
}

type QueryLiveRecordingDurationResponseHeader struct {
}

func (s QueryLiveRecordingDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseHeader) GoString() string {
  return s.String()
}




type HttpDnsStatisticsRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Querying date, the date format is yyyy-mm-dd, and the current date will be adopted by default if the field is not selected or left empty;", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"Querying start date, and the date format is yyyy-mm-dd.
  // The parameter needs to work in concert with enddate, and it will be invalid if there exists date parameter", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"Querying end date, and the date format is yyyy-mm-dd.
  // The parameter should work in concert with the parameter startdate. It will be invalid if there exists date parameter", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"Querying domain, and use English semicolon ;  as separator if there are multiple domains; all domains of the user will be queried if the field is not selected or left empty.", "zh_CN":"查询的频道，多个频道值请用英文分号;，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"查询域名所属的加速类型，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"Response result format, and the supporting format is xml and json. Default as xml.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Whether the channel matches exactly: when true, the full domain name must be provided (at this point, any invalid or duplicate channels entered by the user will be filtered out, and a 403 error will be returned if all input channels are invalid). When not true, all channels ending with the user-entered channel are displayed. The default is true.", "zh_CN":" 频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"Whether the display of results includes merged values: Enter 1 to provide merged results; enter 2 to provide only split values; if not selected or left empty, the default is 1.", "zh_CN":"结果的显示是否提供合并值。填写1时：提供合并结果；填写2时：只提供拆分值；不选或者为空时默认为1。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s HttpDnsStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsRequest) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsRequest) SetCust(v string) *HttpDnsStatisticsRequest {
  s.Cust = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetDate(v string) *HttpDnsStatisticsRequest {
  s.Date = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetStartdate(v string) *HttpDnsStatisticsRequest {
  s.Startdate = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetEnddate(v string) *HttpDnsStatisticsRequest {
  s.Enddate = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetChannel(v string) *HttpDnsStatisticsRequest {
  s.Channel = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetAccetype(v string) *HttpDnsStatisticsRequest {
  s.Accetype = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetDataformat(v string) *HttpDnsStatisticsRequest {
  s.Dataformat = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetIsExactMatch(v string) *HttpDnsStatisticsRequest {
  s.IsExactMatch = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetResultType(v string) *HttpDnsStatisticsRequest {
  s.ResultType = &v
  return s
}

type HttpDnsStatisticsResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *HttpDnsStatisticsResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponse) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponse) SetProvider(v *HttpDnsStatisticsResponseProvider) *HttpDnsStatisticsResponse {
  s.Provider = v
  return s
}

type HttpDnsStatisticsResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'解析量数据'}
  Date *HttpDnsStatisticsResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProvider) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProvider) SetName(v string) *HttpDnsStatisticsResponseProvider {
  s.Name = &v
  return s
}

func (s *HttpDnsStatisticsResponseProvider) SetType(v string) *HttpDnsStatisticsResponseProvider {
  s.Type = &v
  return s
}

func (s *HttpDnsStatisticsResponseProvider) SetResultType(v string) *HttpDnsStatisticsResponseProvider {
  s.ResultType = &v
  return s
}

func (s *HttpDnsStatisticsResponseProvider) SetDate(v *HttpDnsStatisticsResponseProviderDate) *HttpDnsStatisticsResponseProvider {
  s.Date = v
  return s
}

type HttpDnsStatisticsResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *HttpDnsStatisticsResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProviderDate) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProviderDate) SetStartdate(v string) *HttpDnsStatisticsResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDate) SetEnddate(v string) *HttpDnsStatisticsResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDate) SetChannel(v *HttpDnsStatisticsResponseProviderDateChannel) *HttpDnsStatisticsResponseProviderDate {
  s.Channel = v
  return s
}

type HttpDnsStatisticsResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'总解析量'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'httpdns', 'zh_CN':'解析量数据'}
  Httpdns []*HttpDnsStatisticsResponseProviderDateChannelHttpdns `json:"httpdns,omitempty" xml:"httpdns,omitempty" require:"true" type:"Repeated"`
}

func (s HttpDnsStatisticsResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProviderDateChannel) SetName(v string) *HttpDnsStatisticsResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDateChannel) SetTotal(v string) *HttpDnsStatisticsResponseProviderDateChannel {
  s.Total = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDateChannel) SetHttpdns(v []*HttpDnsStatisticsResponseProviderDateChannelHttpdns) *HttpDnsStatisticsResponseProviderDateChannel {
  s.Httpdns = v
  return s
}

type HttpDnsStatisticsResponseProviderDateChannelHttpdns struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'解析量'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s HttpDnsStatisticsResponseProviderDateChannelHttpdns) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProviderDateChannelHttpdns) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProviderDateChannelHttpdns) SetTime(v string) *HttpDnsStatisticsResponseProviderDateChannelHttpdns {
  s.Time = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDateChannelHttpdns) SetText(v string) *HttpDnsStatisticsResponseProviderDateChannelHttpdns {
  s.Text = &v
  return s
}

type HttpDnsStatisticsPaths struct {
}

func (s HttpDnsStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsPaths) GoString() string {
  return s.String()
}

type HttpDnsStatisticsParameters struct {
}

func (s HttpDnsStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsParameters) GoString() string {
  return s.String()
}

type HttpDnsStatisticsRequestHeader struct {
}

func (s HttpDnsStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsRequestHeader) GoString() string {
  return s.String()
}

type HttpDnsStatisticsResponseHeader struct {
}

func (s HttpDnsStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseHeader) GoString() string {
  return s.String()
}




type WafHitRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1)With format yyyy-mm-dd.
  // 2)If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"This parameter specifies if the 'channel' parameter should be exactly matched:
  // 1)'true' as default.
  // 2) If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2)If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1)If there are multiple inputs,use ';' as separator.
  // 2)If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
}

func (s WafHitRequest) String() string {
  return tea.Prettify(s)
}

func (s WafHitRequest) GoString() string {
  return s.String()
}

func (s *WafHitRequest) SetCust(v string) *WafHitRequest {
  s.Cust = &v
  return s
}

func (s *WafHitRequest) SetDate(v string) *WafHitRequest {
  s.Date = &v
  return s
}

func (s *WafHitRequest) SetStartdate(v string) *WafHitRequest {
  s.Startdate = &v
  return s
}

func (s *WafHitRequest) SetEnddate(v string) *WafHitRequest {
  s.Enddate = &v
  return s
}

func (s *WafHitRequest) SetChannel(v string) *WafHitRequest {
  s.Channel = &v
  return s
}

func (s *WafHitRequest) SetIsExactMatch(v string) *WafHitRequest {
  s.IsExactMatch = &v
  return s
}

func (s *WafHitRequest) SetRegion(v string) *WafHitRequest {
  s.Region = &v
  return s
}

func (s *WafHitRequest) SetAccetype(v string) *WafHitRequest {
  s.Accetype = &v
  return s
}

func (s *WafHitRequest) SetDataformat(v string) *WafHitRequest {
  s.Dataformat = &v
  return s
}

type WafHitResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *WafHitResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s WafHitResponse) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponse) GoString() string {
  return s.String()
}

func (s *WafHitResponse) SetProvider(v *WafHitResponseProvider) *WafHitResponse {
  s.Provider = v
  return s
}

type WafHitResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'waf 请求数数据'}
  Date *WafHitResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s WafHitResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponseProvider) GoString() string {
  return s.String()
}

func (s *WafHitResponseProvider) SetName(v string) *WafHitResponseProvider {
  s.Name = &v
  return s
}

func (s *WafHitResponseProvider) SetType(v string) *WafHitResponseProvider {
  s.Type = &v
  return s
}

func (s *WafHitResponseProvider) SetDate(v *WafHitResponseProviderDate) *WafHitResponseProvider {
  s.Date = v
  return s
}

type WafHitResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *WafHitResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s WafHitResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponseProviderDate) GoString() string {
  return s.String()
}

func (s *WafHitResponseProviderDate) SetStartdate(v string) *WafHitResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *WafHitResponseProviderDate) SetEnddate(v string) *WafHitResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *WafHitResponseProviderDate) SetChannel(v *WafHitResponseProviderDateChannel) *WafHitResponseProviderDate {
  s.Channel = v
  return s
}

type WafHitResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'waf 请求数数据'}
  Wafhit []*WafHitResponseProviderDateChannelWafhit `json:"wafhit,omitempty" xml:"wafhit,omitempty" require:"true" type:"Repeated"`
}

func (s WafHitResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *WafHitResponseProviderDateChannel) SetName(v string) *WafHitResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *WafHitResponseProviderDateChannel) SetWafhit(v []*WafHitResponseProviderDateChannelWafhit) *WafHitResponseProviderDateChannel {
  s.Wafhit = v
  return s
}

type WafHitResponseProviderDateChannelWafhit struct     {
  // {'en':'time of every 5 duration,with format yyyy-mmm-dd hh:MM:ss', 'zh_CN':'waf请求数5分钟粒度时间，格式yyyy-mm-dd hh:MM:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'displaying the waf request.', 'zh_CN':'waf请求数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s WafHitResponseProviderDateChannelWafhit) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponseProviderDateChannelWafhit) GoString() string {
  return s.String()
}

func (s *WafHitResponseProviderDateChannelWafhit) SetTime(v string) *WafHitResponseProviderDateChannelWafhit {
  s.Time = &v
  return s
}

func (s *WafHitResponseProviderDateChannelWafhit) SetText(v string) *WafHitResponseProviderDateChannelWafhit {
  s.Text = &v
  return s
}

type WafHitPaths struct {
}

func (s WafHitPaths) String() string {
  return tea.Prettify(s)
}

func (s WafHitPaths) GoString() string {
  return s.String()
}

type WafHitParameters struct {
}

func (s WafHitParameters) String() string {
  return tea.Prettify(s)
}

func (s WafHitParameters) GoString() string {
  return s.String()
}

type WafHitRequestHeader struct {
}

func (s WafHitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s WafHitRequestHeader) GoString() string {
  return s.String()
}

type WafHitResponseHeader struct {
}

func (s WafHitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponseHeader) GoString() string {
  return s.String()
}




type QueryLiveStreamStatusRequest struct {
}

func (s QueryLiveStreamStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusRequest) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusRequestHeader struct {
}

func (s QueryLiveStreamStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusRequestHeader) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusPaths struct {
}

func (s QueryLiveStreamStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusPaths) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusParameters struct {
  // {"en":"Domain (multiple domains supported, separated by commas)(    All parameters are passed via HTTP GET requests.)","zh_CN":"域名（支持多个域名，以逗号分隔）(所有参数以HTTP GET请求方式传参)"}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {"en":"20160527152300, non-real-time indicates the current time -5 minutes, real-time indicates the current time -30 seconds\nA:\n(t is not transmitted, the current system time is obtained and rounded to second (for example, 2017/3/28 14:38:55 rounded to second 2017/3/28 14:38:50 rounded to second);) If the query interval g=10, the final time (dateFrom) is rounded in seconds (e.g. 2017/3/28 14:38:55 after rounded in seconds 2017/3/28 14:38:50); If the g! =10, rounded minutes: (e.g. 2017/3/28 14:38:55 rounded seconds 2017/3/28 14:38:00); DateTo = dataFrom + 9 (9 seconds)","zh_CN":"20160527152300，不填为当前时间-5分钟\n详解：\n（t不传，获取当前系统时间，对秒取整(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:50)；）如果查询间隔g=10，最后获得的时间对秒取整(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:50)；如果g!=10，对分钟取整：(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:00);"}
  T *string `json:"t,omitempty" xml:"t,omitempty"`
  // {"en":"Channel URL (simple channel url are supported,eg:push1.test.com/test/test1)","zh_CN":"流名(支持单流名，如：push1.test.com/test/test1)"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Optional value: push or pull. The default value is push\nPush stands for watershed name\nPull indicates the name of a pull basin","zh_CN":"域名类型，可选值：push、pull，默认push\npush代表推流域名\npull代表拉流域名"}
  D *string `json:"d,omitempty" xml:"d,omitempty"`
  // {"en":"The value can be true or false. The default value is false\nOnly realtime data is returned when realtime=true","zh_CN":"是否返回端口，可选值：true、false，默认false"}
  Showport *string `json:"showport,omitempty" xml:"showport,omitempty"`
  // {"en":"Query interval. The value can be 10 or 60. The default value is 60\nWhen g is 10, query the data of the nearest whole 10 seconds to time t\nWhen g is 60, query the data of the nearest whole minute to time t","zh_CN":"查询间隔，可选值10、60，默认60\n当g为10时，查询距离时间t最近的整10秒点数据\n当g为60时，查询距离时间t最近的整分钟点数据"}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
  // {"en":"it is query realtime datas,default false","zh_CN":"是否返回实时数据，默认false"}
  Realtime *string `json:"realtime,omitempty" xml:"realtime,omitempty"`
  // {"en":"Data extension fields(Only supports non-real-time query), providing optional return fields:gop","zh_CN":"数据扩展字段(仅支持按非实时查)，提供可选返回字段：gop"}
  Expand *string `json:"expand,omitempty" xml:"expand,omitempty"`
}

func (s QueryLiveStreamStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusParameters) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusParameters) SetU(v string) *QueryLiveStreamStatusParameters {
  s.U = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetT(v string) *QueryLiveStreamStatusParameters {
  s.T = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetChannel(v string) *QueryLiveStreamStatusParameters {
  s.Channel = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetD(v string) *QueryLiveStreamStatusParameters {
  s.D = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetShowport(v string) *QueryLiveStreamStatusParameters {
  s.Showport = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetG(v string) *QueryLiveStreamStatusParameters {
  s.G = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetRealtime(v string) *QueryLiveStreamStatusParameters {
  s.Realtime = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetExpand(v string) *QueryLiveStreamStatusParameters {
  s.Expand = &v
  return s
}

type QueryLiveStreamStatusResponse struct {
  // {"en":"The time of the data returned","zh_CN":"返回的数据的时间"}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {"en":"Number of data items. 0 is returned if there is no data","zh_CN":"数据条数，无数据返回0"}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {"en":"Total onlines","zh_CN":"总在线人数"}
  Histscount *int64 `json:"histscount,omitempty" xml:"histscount,omitempty" require:"true"`
  // {"en":"Total channel bandwidth","zh_CN":"总频道带宽"}
  Bandwidthcount *int64 `json:"bandwidthcount,omitempty" xml:"bandwidthcount,omitempty" require:"true"`
  // {"en":"Timestamp","zh_CN":"数据的时间戳,如果有传t,则等于t"}
  Datetime *int64 `json:"datetime,omitempty" xml:"datetime,omitempty" require:"true"`
  // {"en":"","zh_CN":"数据集合"}
  DataValue []*QueryLiveStreamStatusResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLiveStreamStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponse) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusResponse) SetRettime(v string) *QueryLiveStreamStatusResponse {
  s.Rettime = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetRetcode(v int64) *QueryLiveStreamStatusResponse {
  s.Retcode = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetHistscount(v int64) *QueryLiveStreamStatusResponse {
  s.Histscount = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetBandwidthcount(v int64) *QueryLiveStreamStatusResponse {
  s.Bandwidthcount = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetDatetime(v int64) *QueryLiveStreamStatusResponse {
  s.Datetime = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetDataValue(v []*QueryLiveStreamStatusResponseDataValue) *QueryLiveStreamStatusResponse {
  s.DataValue = v
  return s
}

type QueryLiveStreamStatusResponseDataValue struct     {
  // {"en":"The channel of anchor","zh_CN":"主播流名"}
  Streamname *string `json:"streamname,omitempty" xml:"streamname,omitempty" require:"true"`
  // {"en":"IP address of the CDN node","zh_CN":"推流cdn节点IP"}
  Deployaddress *string `json:"deployaddress,omitempty" xml:"deployaddress,omitempty" require:"true"`
  // {"en":"Anchor exit Address","zh_CN":"主播出口地址"}
  Inaddress *string `json:"inaddress,omitempty" xml:"inaddress,omitempty" require:"true"`
  // {"en":"The number of online","zh_CN":"在线人数"}
  Hists *int64 `json:"hists,omitempty" xml:"hists,omitempty" require:"true"`
  // {"en":"Anchor Current bit rate (transcoding stream has no bit rate data) unit: BPS","zh_CN":"主播当前码率(转码流没有码率数据) 单位：bps"}
  Inbandwidth *int64 `json:"inbandwidth,omitempty" xml:"inbandwidth,omitempty" require:"true"`
  // {"en":"Channel current viewing bandwidth unit: BPS","zh_CN":"频道当前观看带宽 单位：bps"}
  Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Anchor delay (MS)","zh_CN":"主播延迟(ms)"}
  Delay *int64 `json:"delay,omitempty" xml:"delay,omitempty" require:"true"`
  // {"en":"Anchor Current encoding frame rate","zh_CN":"主播当前编码帧率"}
  Fps *int64 `json:"fps,omitempty" xml:"fps,omitempty" require:"true"`
  // {"en":"Current frame loss rate of anchor","zh_CN":"主播当前丢帧率"}
  Lfr *QueryLiveStreamStatusResponseDataValueLfr `json:"lfr,omitempty" xml:"lfr,omitempty" require:"true" type:"Struct"`
  // {"en":"Anchor raw frame rate","zh_CN":"主播原始帧率"}
  Ofr *int64 `json:"ofr,omitempty" xml:"ofr,omitempty" require:"true"`
  // {"en":"The resolution of the","zh_CN":"分辨率"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {"en":"Video coding","zh_CN":"视频编码"}
  VideoCodec *string `json:"video_codec,omitempty" xml:"video_codec,omitempty" require:"true"`
  // {"en":"Audio coding","zh_CN":"音频编码"}
  AudioCodec *string `json:"audio_codec,omitempty" xml:"audio_codec,omitempty" require:"true"`
  // {"en":"Keyframe interval","zh_CN":"关键帧间隔,有传expand才会返回"}
  Gop *int64 `json:"gop,omitempty" xml:"gop,omitempty" require:"true"`
}

func (s QueryLiveStreamStatusResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponseDataValue) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusResponseDataValue) SetStreamname(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Streamname = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetDeployaddress(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Deployaddress = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetInaddress(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Inaddress = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetHists(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Hists = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetInbandwidth(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Inbandwidth = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetBandwidth(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Bandwidth = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetDelay(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Delay = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetFps(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Fps = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetLfr(v *QueryLiveStreamStatusResponseDataValueLfr) *QueryLiveStreamStatusResponseDataValue {
  s.Lfr = v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetOfr(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Ofr = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetResolution(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Resolution = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetVideoCodec(v string) *QueryLiveStreamStatusResponseDataValue {
  s.VideoCodec = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetAudioCodec(v string) *QueryLiveStreamStatusResponseDataValue {
  s.AudioCodec = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetGop(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Gop = &v
  return s
}

type QueryLiveStreamStatusResponseDataValueLfr struct {
}

func (s QueryLiveStreamStatusResponseDataValueLfr) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponseDataValueLfr) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusResponseHeader struct {
}

func (s QueryLiveStreamStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponseHeader) GoString() string {
  return s.String()
}




type AttOverViewRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way.
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"attack Type.
  // 1.Including DDOS,WAF and BOT.
  // 2.inputs 'DDOS' matches 'DDOS' data,'BOT' matches 'BOT' data,and 'WAF' matches  all types started with 'WAF'.
  // 3.If there are multiple  inputss,use english ';' as separator.
  // 4.If not specified,means all the types.", "zh_CN":"攻击类型DDOS、WAF或BOT。填写DDOS时匹配DDOS数据，填写WAF时匹配WAF_xxx数据，填写BOT时匹配BOT数据，多个类型之间用英文分号;隔开;不填或放空默认为全部类型"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty"`
}

func (s AttOverViewRequest) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewRequest) GoString() string {
  return s.String()
}

func (s *AttOverViewRequest) SetCust(v string) *AttOverViewRequest {
  s.Cust = &v
  return s
}

func (s *AttOverViewRequest) SetDate(v string) *AttOverViewRequest {
  s.Date = &v
  return s
}

func (s *AttOverViewRequest) SetStartdate(v string) *AttOverViewRequest {
  s.Startdate = &v
  return s
}

func (s *AttOverViewRequest) SetEnddate(v string) *AttOverViewRequest {
  s.Enddate = &v
  return s
}

func (s *AttOverViewRequest) SetChannel(v string) *AttOverViewRequest {
  s.Channel = &v
  return s
}

func (s *AttOverViewRequest) SetIsExactMatch(v string) *AttOverViewRequest {
  s.IsExactMatch = &v
  return s
}

func (s *AttOverViewRequest) SetRegion(v string) *AttOverViewRequest {
  s.Region = &v
  return s
}

func (s *AttOverViewRequest) SetAccetype(v string) *AttOverViewRequest {
  s.Accetype = &v
  return s
}

func (s *AttOverViewRequest) SetDataformat(v string) *AttOverViewRequest {
  s.Dataformat = &v
  return s
}

func (s *AttOverViewRequest) SetResultType(v string) *AttOverViewRequest {
  s.ResultType = &v
  return s
}

func (s *AttOverViewRequest) SetAttackType(v string) *AttOverViewRequest {
  s.AttackType = &v
  return s
}

type AttOverViewResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *AttOverViewResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s AttOverViewResponse) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponse) GoString() string {
  return s.String()
}

func (s *AttOverViewResponse) SetProvider(v *AttOverViewResponseProvider) *AttOverViewResponse {
  s.Provider = v
  return s
}

type AttOverViewResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'安全防护安全状况汇总数据'}
  Date *AttOverViewResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s AttOverViewResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponseProvider) GoString() string {
  return s.String()
}

func (s *AttOverViewResponseProvider) SetName(v string) *AttOverViewResponseProvider {
  s.Name = &v
  return s
}

func (s *AttOverViewResponseProvider) SetType(v string) *AttOverViewResponseProvider {
  s.Type = &v
  return s
}

func (s *AttOverViewResponseProvider) SetResultType(v string) *AttOverViewResponseProvider {
  s.ResultType = &v
  return s
}

func (s *AttOverViewResponseProvider) SetDate(v *AttOverViewResponseProviderDate) *AttOverViewResponseProvider {
  s.Date = v
  return s
}

type AttOverViewResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *AttOverViewResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s AttOverViewResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponseProviderDate) GoString() string {
  return s.String()
}

func (s *AttOverViewResponseProviderDate) SetStartdate(v string) *AttOverViewResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *AttOverViewResponseProviderDate) SetEnddate(v string) *AttOverViewResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *AttOverViewResponseProviderDate) SetChannel(v *AttOverViewResponseProviderDateChannel) *AttOverViewResponseProviderDate {
  s.Channel = v
  return s
}

type AttOverViewResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'overview', 'zh_CN':'安全防护安全状况汇总数据'}
  OverView []*AttOverViewResponseProviderDateChannelOverView `json:"over-view,omitempty" xml:"over-view,omitempty" require:"true" type:"Repeated"`
}

func (s AttOverViewResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *AttOverViewResponseProviderDateChannel) SetName(v string) *AttOverViewResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *AttOverViewResponseProviderDateChannel) SetOverView(v []*AttOverViewResponseProviderDateChannelOverView) *AttOverViewResponseProviderDateChannel {
  s.OverView = v
  return s
}

type AttOverViewResponseProviderDateChannelOverView struct     {
  // {'en':'timestamp', 'zh_CN':'时间点，格式 yyyy-MM-dd hh:mm:ss'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'total request number.', 'zh_CN':'总请求数'}
  Hit *string `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
  // {'en':'total denied request number.', 'zh_CN':'总访问拒绝数'}
  Deny *string `json:"deny,omitempty" xml:"deny,omitempty" require:"true"`
}

func (s AttOverViewResponseProviderDateChannelOverView) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponseProviderDateChannelOverView) GoString() string {
  return s.String()
}

func (s *AttOverViewResponseProviderDateChannelOverView) SetChannel(v string) *AttOverViewResponseProviderDateChannelOverView {
  s.Channel = &v
  return s
}

func (s *AttOverViewResponseProviderDateChannelOverView) SetHit(v string) *AttOverViewResponseProviderDateChannelOverView {
  s.Hit = &v
  return s
}

func (s *AttOverViewResponseProviderDateChannelOverView) SetDeny(v string) *AttOverViewResponseProviderDateChannelOverView {
  s.Deny = &v
  return s
}

type AttOverViewPaths struct {
}

func (s AttOverViewPaths) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewPaths) GoString() string {
  return s.String()
}

type AttOverViewParameters struct {
}

func (s AttOverViewParameters) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewParameters) GoString() string {
  return s.String()
}

type AttOverViewRequestHeader struct {
}

func (s AttOverViewRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewRequestHeader) GoString() string {
  return s.String()
}

type AttOverViewResponseHeader struct {
}

func (s AttOverViewResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponseHeader) GoString() string {
  return s.String()
}




type ReportBandwidthLowDelayP2pServiceRequest struct {
  // {'en':'-', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2021-05-19T10:00:00+08:00（为北京时间2021年5月19日10点0分0秒）
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'-', 'zh_CN':'结束时间：
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间
  // 3.dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天。（可联系技术支持调整）'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'-', 'zh_CN':'域名：
  // 1、可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 2、自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'-', 'zh_CN':'可选值：domain
  // 不传默认聚合所有频道数据'}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportBandwidthLowDelayP2pServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportBandwidthLowDelayP2pServiceRequest) SetDateFrom(v string) *ReportBandwidthLowDelayP2pServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceRequest) SetDateTo(v string) *ReportBandwidthLowDelayP2pServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceRequest) SetDomain(v []*string) *ReportBandwidthLowDelayP2pServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceRequest) SetGroupBy(v []*string) *ReportBandwidthLowDelayP2pServiceRequest {
  s.GroupBy = v
  return s
}

type ReportBandwidthLowDelayP2pServiceResponse struct {
  // {'en':'-', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'-', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportBandwidthLowDelayP2pServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthLowDelayP2pServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportBandwidthLowDelayP2pServiceResponse) SetCode(v string) *ReportBandwidthLowDelayP2pServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceResponse) SetMessage(v string) *ReportBandwidthLowDelayP2pServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceResponse) SetData(v []*ReportBandwidthLowDelayP2pServiceResponseData) *ReportBandwidthLowDelayP2pServiceResponse {
  s.Data = v
  return s
}

type ReportBandwidthLowDelayP2pServiceResponseData struct     {
  // {'en':'-', 'zh_CN':'域名，聚合全部域名数据不返回该字段'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'-', 'zh_CN':'请求结果的详细数据'}
  DetailList []*ReportBandwidthLowDelayP2pServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthLowDelayP2pServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthLowDelayP2pServiceResponseData) SetDomain(v string) *ReportBandwidthLowDelayP2pServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceResponseData) SetDetailList(v []*ReportBandwidthLowDelayP2pServiceResponseDataDetailList) *ReportBandwidthLowDelayP2pServiceResponseData {
  s.DetailList = v
  return s
}

type ReportBandwidthLowDelayP2pServiceResponseDataDetailList struct     {
  // {'en':'-', 'zh_CN':'时间片,返回开始时间和结束时间包含的时间片。时间格式：yyyy-MM-dd HH:mm'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'-', 'zh_CN':'P2P带宽'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportBandwidthLowDelayP2pServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportBandwidthLowDelayP2pServiceResponseDataDetailList) SetTimestamp(v string) *ReportBandwidthLowDelayP2pServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceResponseDataDetailList) SetValue(v string) *ReportBandwidthLowDelayP2pServiceResponseDataDetailList {
  s.Value = &v
  return s
}

type ReportBandwidthLowDelayP2pServicePaths struct {
}

func (s ReportBandwidthLowDelayP2pServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServicePaths) GoString() string {
  return s.String()
}

type ReportBandwidthLowDelayP2pServiceParameters struct {
}

func (s ReportBandwidthLowDelayP2pServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceParameters) GoString() string {
  return s.String()
}

type ReportBandwidthLowDelayP2pServiceRequestHeader struct {
}

func (s ReportBandwidthLowDelayP2pServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportBandwidthLowDelayP2pServiceResponseHeader struct {
}

func (s ReportBandwidthLowDelayP2pServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceResponseHeader) GoString() string {
  return s.String()
}




type DnsTestRequest struct {
  // {"en":"Domain name or IP to be detected", "zh_CN":"需要检测的域名或 IP"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"DNS record type, optional values: A, CNAME, SOA, TXT, MX, NS", "zh_CN":"DNS 记录类型，可选值：A、CNAME、SOA、TXT、MX、NS"}
  DnsType *string `json:"dnsType,omitempty" xml:"dnsType,omitempty" require:"true"`
  // {"en":"Specify a DNS server, such as 8.8.8.8", "zh_CN":"指定 DNS 服务器，如 114.114.114.114"}
  DnsServer *string `json:"dnsServer,omitempty" xml:"dnsServer,omitempty" require:"true"`
  // {"en":"The region of the monitor belongs to, supports mainland China provinces, Hong Kong, Macao and Taiwan, as well as overseas countries.
  // optional values:, anhui, beijing, chongqing, fujian, gansu, guangdong, guangxi, guizhou, hainan, hebei, heilongjiang, henan, hubei, hunan, jiangsu, jiangxi, jilin, liaoning, neimenggu, ningxia, qinghai, shaanxi, shandong, shanghai, shanxi, sichuan, tianjin, xinjiang, xizang, yunnan, zhejiang, TW, HK, MO, AE, AU, BD, BN, BR, CA, CL, CO, DJ, ID, IT, JP, KG, KH, KR, KW, LA, MG, MM, MU, MY, NP, OM, PE, PH, PK, QA, RO, RU, SA, SE, SG, TH, TR, US, VN", "zh_CN":"监控机所属地区，支持中国大陆省份、港澳台以及海外国家。
  // 可选值：
  // anhui: 安徽
  // beijing: 北京
  // chongqing: 重庆
  // fujian: 福建
  // gansu: 甘肃
  // guangdong: 广东
  // guangxi: 广西
  // guizhou: 贵州
  // hainan: 海南
  // hebei: 河北
  // heilongjiang: 黑龙江
  // henan: 河南
  // hubei: 湖北
  // hunan: 湖南
  // jiangsu: 江苏
  // jiangxi: 江西
  // jilin: 吉林
  // liaoning: 辽宁
  // neimenggu: 内蒙古
  // ningxia: 宁夏
  // qinghai: 青海
  // shaanxi: 陕西
  // shandong: 山东
  // shanghai: 上海
  // shanxi: 山西
  // sichuan: 四川
  // tianjin: 天津
  // xinjiang: 新疆
  // xizang: 西藏
  // yunnan: 云南
  // zhejiang: 浙江
  // TW: 台湾
  // HK: 香港
  // MO: 澳门
  // AE: 阿联酋
  // AU: 澳大利亚
  // BD: 孟加拉
  // BN: 文莱
  // BR: 巴西
  // CA: 加拿大
  // CL: 智利
  // CO: 哥伦比亚
  // DJ: 吉布提
  // ID: 印度尼西亚
  // IT: 意大利
  // JP: 日本
  // KG: 吉尔吉斯斯坦
  // KH: 柬埔寨
  // KR: 韩国
  // KW: 科威特
  // LA: 老挝
  // MG: 马达加斯加
  // MM: 缅甸
  // MU: 毛里求斯
  // MY: 马来西亚
  // NP: 尼泊尔
  // OM: 阿曼
  // PE: 秘鲁
  // PH: 菲律宾
  // PK: 巴基斯坦
  // QA: 卡塔尔
  // RO: 罗马尼亚
  // RU: 俄罗斯
  // SA: 沙特阿拉伯
  // SE: 瑞典
  // SG: 新加坡
  // TH: 泰国
  // TR: 土耳其
  // US: 美国
  // VN: 越南"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"ISP
  // optional values: 
  // 0: China-Telecom
  // 1: China-Unicom
  // 2: China-Tietong
  // 4: China-Mobile
  // 5: China-Education-and-Research
  // 9: China-Cable-Television
  // 10: GreatWall-Broadband", "zh_CN":"监控机所属运营商中文名。
  // 可选值：
  // 0: 中国电信
  // 1: 中国联通
  // 2: 中国铁通
  // 4: 中国移动
  // 5: 中国教育网
  // 9: 中国广电
  // 10: 长城宽带"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s DnsTestRequest) String() string {
  return tea.Prettify(s)
}

func (s DnsTestRequest) GoString() string {
  return s.String()
}

func (s *DnsTestRequest) SetHost(v string) *DnsTestRequest {
  s.Host = &v
  return s
}

func (s *DnsTestRequest) SetDnsType(v string) *DnsTestRequest {
  s.DnsType = &v
  return s
}

func (s *DnsTestRequest) SetDnsServer(v string) *DnsTestRequest {
  s.DnsServer = &v
  return s
}

func (s *DnsTestRequest) SetArea(v string) *DnsTestRequest {
  s.Area = &v
  return s
}

func (s *DnsTestRequest) SetIsp(v string) *DnsTestRequest {
  s.Isp = &v
  return s
}

type DnsTestResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result *DnsTestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s DnsTestResponse) String() string {
  return tea.Prettify(s)
}

func (s DnsTestResponse) GoString() string {
  return s.String()
}

func (s *DnsTestResponse) SetResult(v *DnsTestResponseResult) *DnsTestResponse {
  s.Result = v
  return s
}

type DnsTestResponseResult struct {
  // {'en':'status', 'zh_CN':'状态码'}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'error message', 'zh_CN':'错误信息'}
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty" require:"true"`
  // {'en':'test id', 'zh_CN':'任务 ID'}
  TestId *string `json:"testId,omitempty" xml:"testId,omitempty" require:"true"`
  // {'en':'domain', 'zh_CN':'目标域名'}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {'en':'data','zh_CN':'探测数据'}
  Data []*DnsTestResponseResultData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s DnsTestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s DnsTestResponseResult) GoString() string {
  return s.String()
}

func (s *DnsTestResponseResult) SetStatus(v string) *DnsTestResponseResult {
  s.Status = &v
  return s
}

func (s *DnsTestResponseResult) SetErrorMsg(v string) *DnsTestResponseResult {
  s.ErrorMsg = &v
  return s
}

func (s *DnsTestResponseResult) SetTestId(v string) *DnsTestResponseResult {
  s.TestId = &v
  return s
}

func (s *DnsTestResponseResult) SetHost(v string) *DnsTestResponseResult {
  s.Host = &v
  return s
}

func (s *DnsTestResponseResult) SetData(v []*DnsTestResponseResultData) *DnsTestResponseResult {
  s.Data = v
  return s
}

type DnsTestResponseResultData struct     {
  // {'en':'task id', 'zh_CN':'任务 ID'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'task id', 'zh_CN':'任务 ID'}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {'en':'domain', 'zh_CN':'目标域名'}
  QuestionUrl *string `json:"questionUrl,omitempty" xml:"questionUrl,omitempty" require:"true"`
  // {'en':'monitor IP', 'zh_CN':'监控机 IP'}
  DetectIp *string `json:"detectIp,omitempty" xml:"detectIp,omitempty" require:"true"`
  // {'en':'monitor ISP', 'zh_CN':'监控机运营商'}
  DetectIpIsp *string `json:"detectIpIsp,omitempty" xml:"detectIpIsp,omitempty" require:"true"`
  // {'en':'monitor ISP', 'zh_CN':'监控机运营商编码'}
  DetectIpIspCode *string `json:"detectIpIspCode,omitempty" xml:"detectIpIspCode,omitempty" require:"true"`
  // {'en':'monitor province', 'zh_CN':'监控机所属省份'}
  DetectIpPro *string `json:"detectIpPro,omitempty" xml:"detectIpPro,omitempty" require:"true"`
  // {'en':'monitor province code', 'zh_CN':'监控机所属省份编码'}
  DetectIpProCode *string `json:"detectIpProCode,omitempty" xml:"detectIpProCode,omitempty" require:"true"`
  // {'en':'dns server', 'zh_CN':'DNS'}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
  // {'en':'time', 'zh_CN':'探测时间'}
  DetectTime *string `json:"detectTime,omitempty" xml:"detectTime,omitempty" require:"true"`
  // {'en':'end time', 'zh_CN':'探测结束时间戳'}
  DoneTime *int64 `json:"doneTime,omitempty" xml:"doneTime,omitempty" require:"true"`
  // {'en':'dns query time', 'zh_CN':'查询耗时，单位：毫秒'}
  QueryTime *int32 `json:"queryTime,omitempty" xml:"queryTime,omitempty" require:"true"`
  // {'en':'status', 'zh_CN':'探测状态'}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'status code', 'zh_CN':'HTTP 状态码'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'dns answer', 'zh_CN':'DNS'}
  Answer *string `json:"answer,omitempty" xml:"answer,omitempty" require:"true"`
  // {'en':'dns result','zh_CN':'解析结果'}
  AnswerList []*DnsTestResponseResultDataAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" require:"true" type:"Repeated"`
}

func (s DnsTestResponseResultData) String() string {
  return tea.Prettify(s)
}

func (s DnsTestResponseResultData) GoString() string {
  return s.String()
}

func (s *DnsTestResponseResultData) SetId(v string) *DnsTestResponseResultData {
  s.Id = &v
  return s
}

func (s *DnsTestResponseResultData) SetTaskId(v string) *DnsTestResponseResultData {
  s.TaskId = &v
  return s
}

func (s *DnsTestResponseResultData) SetQuestionUrl(v string) *DnsTestResponseResultData {
  s.QuestionUrl = &v
  return s
}

func (s *DnsTestResponseResultData) SetDetectIp(v string) *DnsTestResponseResultData {
  s.DetectIp = &v
  return s
}

func (s *DnsTestResponseResultData) SetDetectIpIsp(v string) *DnsTestResponseResultData {
  s.DetectIpIsp = &v
  return s
}

func (s *DnsTestResponseResultData) SetDetectIpIspCode(v string) *DnsTestResponseResultData {
  s.DetectIpIspCode = &v
  return s
}

func (s *DnsTestResponseResultData) SetDetectIpPro(v string) *DnsTestResponseResultData {
  s.DetectIpPro = &v
  return s
}

func (s *DnsTestResponseResultData) SetDetectIpProCode(v string) *DnsTestResponseResultData {
  s.DetectIpProCode = &v
  return s
}

func (s *DnsTestResponseResultData) SetServer(v string) *DnsTestResponseResultData {
  s.Server = &v
  return s
}

func (s *DnsTestResponseResultData) SetDetectTime(v string) *DnsTestResponseResultData {
  s.DetectTime = &v
  return s
}

func (s *DnsTestResponseResultData) SetDoneTime(v int64) *DnsTestResponseResultData {
  s.DoneTime = &v
  return s
}

func (s *DnsTestResponseResultData) SetQueryTime(v int32) *DnsTestResponseResultData {
  s.QueryTime = &v
  return s
}

func (s *DnsTestResponseResultData) SetStatus(v int32) *DnsTestResponseResultData {
  s.Status = &v
  return s
}

func (s *DnsTestResponseResultData) SetStatusCode(v string) *DnsTestResponseResultData {
  s.StatusCode = &v
  return s
}

func (s *DnsTestResponseResultData) SetAnswer(v string) *DnsTestResponseResultData {
  s.Answer = &v
  return s
}

func (s *DnsTestResponseResultData) SetAnswerList(v []*DnsTestResponseResultDataAnswerList) *DnsTestResponseResultData {
  s.AnswerList = v
  return s
}

type DnsTestResponseResultDataAnswerList struct     {
  // {'en':'id','zh_CN':'ID'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'tess id','zh_CN':'任务 ID'}
  TestId *string `json:"testId,omitempty" xml:"testId,omitempty" require:"true"`
  // {'en':'isp','zh_CN':'记录所属地区运营商'}
  Address *string `json:"address,omitempty" xml:"address,omitempty" require:"true"`
  // {'en':'dns class','zh_CN':'DNS 记录'}
  DnsClass *string `json:"dnsClass,omitempty" xml:"dnsClass,omitempty" require:"true"`
  // {'en':'dns type','zh_CN':'DNS 记录类型'}
  DnsType *string `json:"dnsType,omitempty" xml:"dnsType,omitempty" require:"true"`
  // {'en':'dns record','zh_CN':'DNS 记录值'}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
  // {'en':'time to live', 'zh_CN':'TTL'}
  Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s DnsTestResponseResultDataAnswerList) String() string {
  return tea.Prettify(s)
}

func (s DnsTestResponseResultDataAnswerList) GoString() string {
  return s.String()
}

func (s *DnsTestResponseResultDataAnswerList) SetId(v string) *DnsTestResponseResultDataAnswerList {
  s.Id = &v
  return s
}

func (s *DnsTestResponseResultDataAnswerList) SetTestId(v string) *DnsTestResponseResultDataAnswerList {
  s.TestId = &v
  return s
}

func (s *DnsTestResponseResultDataAnswerList) SetAddress(v string) *DnsTestResponseResultDataAnswerList {
  s.Address = &v
  return s
}

func (s *DnsTestResponseResultDataAnswerList) SetDnsClass(v string) *DnsTestResponseResultDataAnswerList {
  s.DnsClass = &v
  return s
}

func (s *DnsTestResponseResultDataAnswerList) SetDnsType(v string) *DnsTestResponseResultDataAnswerList {
  s.DnsType = &v
  return s
}

func (s *DnsTestResponseResultDataAnswerList) SetResult(v string) *DnsTestResponseResultDataAnswerList {
  s.Result = &v
  return s
}

func (s *DnsTestResponseResultDataAnswerList) SetTtl(v int32) *DnsTestResponseResultDataAnswerList {
  s.Ttl = &v
  return s
}

type DnsTestPaths struct {
}

func (s DnsTestPaths) String() string {
  return tea.Prettify(s)
}

func (s DnsTestPaths) GoString() string {
  return s.String()
}

type DnsTestParameters struct {
}

func (s DnsTestParameters) String() string {
  return tea.Prettify(s)
}

func (s DnsTestParameters) GoString() string {
  return s.String()
}

type DnsTestRequestHeader struct {
}

func (s DnsTestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DnsTestRequestHeader) GoString() string {
  return s.String()
}

type DnsTestResponseHeader struct {
}

func (s DnsTestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DnsTestResponseHeader) GoString() string {
  return s.String()
}




type QueryOnlineViewerCountRequest struct {
}

func (s QueryOnlineViewerCountRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountRequest) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountRequestHeader struct {
}

func (s QueryOnlineViewerCountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountRequestHeader) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountPaths struct {
}

func (s QueryOnlineViewerCountPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountPaths) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountParameters struct {
  // {"en":"Domain (multiple domains supported, separated by commas)(    All parameters are passed via HTTP GET requests.)","zh_CN":"域名（支持多个域名，以逗号分隔）(所有参数以HTTP GET请求方式传参)"}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {"en":"Time,eg:20160527152300.If the parameter is not specified, the value is 3 minutes","zh_CN":"时间，eg：20160527152300，不填为当前时间-3分钟"}
  T *string `json:"t,omitempty" xml:"t,omitempty"`
  // {"en":"The domain type can be pull or push. If this parameter is not specified, the default value is pull","zh_CN":"域名类型，pull或push，不填时默认为pull"}
  D *string `json:"d,omitempty" xml:"d,omitempty"`
  // {"en":"Channel URL(single channel query only),It is not recommended to query with this parameter, and the performance of range query is poor.","zh_CN":"频道URL(仅支持单频道查询),不建议带该参数查询，范围查询性能较差"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Start time, eg: 20160803103500, when both from and to are filled or left blank, channel parameter is mandatory. The query time span is two hours at most. If the query time exceeds two hours, the system queries the data within two hours from the start time and the number of online users within the last 7 days","zh_CN":"开始时间，eg: 20160803103500，from和to都填或都不填，都填时channel参数必填，查询时间跨度最大为两个小时，如果超过两个小时，将查询开始时间两个小时内的数据，可查近7天内在线人数数据"}
  From *string `json:"from,omitempty" xml:"from,omitempty"`
  // {"en":"End time, eg: 20160803103900, when both from and to are filled or left blank, channel parameter is mandatory. The query time span is two hours at most. If the query time exceeds two hours, the system queries the data within two hours from the start time and the number of online users in the last 7 days","zh_CN":"结束时间，eg: 20160803103900，from和to都填或都不填，都填时channel参数必填，查询时间跨度最大为两个小时，如果超过两个小时，将查询开始时间两个小时内的数据，可查近7天内在线人数数据"}
  To *string `json:"to,omitempty" xml:"to,omitempty"`
  // {"en":"Query interval, optional value: 10, 60s.\nWhen g is 10, the number of online users every 10 seconds is queried;\nWhen g is 60, the number of online users at the whole minute within the time range is queried.","zh_CN":"查询间隔，可选值10、60s，\n当g为10时，查询时间范围内每10秒的在线人数\n当g为60时，查询时间范围内整分钟点对应的在线人数"}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
  // {"en":"it is query realtime datas,default false","zh_CN":"是否返回实时数据，默认false"}
  Realtime *string `json:"realtime,omitempty" xml:"realtime,omitempty"`
  // {"en":"The default value is false. If the value is true, the domain data is split. If the value is false, the domain data is merged","zh_CN":"域名拆分控制，默认为false，为true时，拆分域名数据，为false时，合并域名数据"}
  Unpack *string `json:"unpack,omitempty" xml:"unpack,omitempty"`
}

func (s QueryOnlineViewerCountParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountParameters) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountParameters) SetU(v string) *QueryOnlineViewerCountParameters {
  s.U = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetT(v string) *QueryOnlineViewerCountParameters {
  s.T = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetD(v string) *QueryOnlineViewerCountParameters {
  s.D = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetChannel(v string) *QueryOnlineViewerCountParameters {
  s.Channel = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetFrom(v string) *QueryOnlineViewerCountParameters {
  s.From = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetTo(v string) *QueryOnlineViewerCountParameters {
  s.To = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetG(v string) *QueryOnlineViewerCountParameters {
  s.G = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetRealtime(v string) *QueryOnlineViewerCountParameters {
  s.Realtime = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetUnpack(v string) *QueryOnlineViewerCountParameters {
  s.Unpack = &v
  return s
}

type QueryOnlineViewerCountResponse struct {
  // {"en":"Total number of online users. This parameter is displayed only when from and to are empty","zh_CN":"在线总人数，仅当查询时间点，即from和to为空时才显示"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Total abnormal online number, only need customers return","zh_CN":"异常总在线人数，只对有需要客户进行返回"}
  ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty" require:"true"`
  // {"en":"The number of data items is displayed only when from and to are empty.","zh_CN":"数据条数，仅当查询时间点，即from和to为空时才显示"}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {"en":"The time of the data returned","zh_CN":"返回的数据的时间"}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {"en":"","zh_CN":"数据集合"}
  DataValue []*QueryOnlineViewerCountResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOnlineViewerCountResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponse) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountResponse) SetCount(v int64) *QueryOnlineViewerCountResponse {
  s.Count = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetErrorCount(v int64) *QueryOnlineViewerCountResponse {
  s.ErrorCount = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetRetcode(v int64) *QueryOnlineViewerCountResponse {
  s.Retcode = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetRettime(v string) *QueryOnlineViewerCountResponse {
  s.Rettime = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetDataValue(v []*QueryOnlineViewerCountResponseDataValue) *QueryOnlineViewerCountResponse {
  s.DataValue = v
  return s
}

type QueryOnlineViewerCountResponseDataValue struct     {
  // {"en":"Stream name","zh_CN":"流名"}
  Prog *string `json:"prog,omitempty" xml:"prog,omitempty" require:"true"`
  // {"en":"Time is displayed only if the query time range (channel, FROM, and to parameters) is not empty","zh_CN":"时间，仅当查询时间范围，即channel，from和to参数不为空时才显示"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"The number of online","zh_CN":"在线人数"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Abnormal online number, only need customers return","zh_CN":"异常在线人数，只对需要客户进行返回"}
  ErrorValue *int64 `json:"errorValue,omitempty" xml:"errorValue,omitempty" require:"true"`
}

func (s QueryOnlineViewerCountResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponseDataValue) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountResponseDataValue) SetProg(v string) *QueryOnlineViewerCountResponseDataValue {
  s.Prog = &v
  return s
}

func (s *QueryOnlineViewerCountResponseDataValue) SetTime(v string) *QueryOnlineViewerCountResponseDataValue {
  s.Time = &v
  return s
}

func (s *QueryOnlineViewerCountResponseDataValue) SetValue(v int64) *QueryOnlineViewerCountResponseDataValue {
  s.Value = &v
  return s
}

func (s *QueryOnlineViewerCountResponseDataValue) SetErrorValue(v int64) *QueryOnlineViewerCountResponseDataValue {
  s.ErrorValue = &v
  return s
}

type QueryOnlineViewerCountResponseHeader struct {
}

func (s QueryOnlineViewerCountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponseHeader) GoString() string {
  return s.String()
}




