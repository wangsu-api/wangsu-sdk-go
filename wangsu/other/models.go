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
  Result *MtrTestMtrTestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s MtrTestResponse) String() string {
  return tea.Prettify(s)
}

func (s MtrTestResponse) GoString() string {
  return s.String()
}

func (s *MtrTestResponse) SetResult(v *MtrTestMtrTestResponseResult) *MtrTestResponse {
  s.Result = v
  return s
}

type MtrTestMtrTestResponseResult struct {
  // {'en':'', 'zh_CN':''}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'', 'zh_CN':''}
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'目标域名'}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {'en':'data','zh_CN':'探测数据'}
  Data []*MtrTestMtrTestResponseResultData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s MtrTestMtrTestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s MtrTestMtrTestResponseResult) GoString() string {
  return s.String()
}

func (s *MtrTestMtrTestResponseResult) SetStatus(v string) *MtrTestMtrTestResponseResult {
  s.Status = &v
  return s
}

func (s *MtrTestMtrTestResponseResult) SetErrorMsg(v string) *MtrTestMtrTestResponseResult {
  s.ErrorMsg = &v
  return s
}

func (s *MtrTestMtrTestResponseResult) SetHost(v string) *MtrTestMtrTestResponseResult {
  s.Host = &v
  return s
}

func (s *MtrTestMtrTestResponseResult) SetData(v []*MtrTestMtrTestResponseResultData) *MtrTestMtrTestResponseResult {
  s.Data = v
  return s
}

type MtrTestMtrTestResponseResultData struct     {
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
  Items []*MtrTestMtrTestResponseResultDataItems `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s MtrTestMtrTestResponseResultData) String() string {
  return tea.Prettify(s)
}

func (s MtrTestMtrTestResponseResultData) GoString() string {
  return s.String()
}

func (s *MtrTestMtrTestResponseResultData) SetId(v string) *MtrTestMtrTestResponseResultData {
  s.Id = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetTaskId(v string) *MtrTestMtrTestResponseResultData {
  s.TaskId = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDetectIp(v string) *MtrTestMtrTestResponseResultData {
  s.DetectIp = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDetectIpIsp(v string) *MtrTestMtrTestResponseResultData {
  s.DetectIpIsp = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDetectIpIspCode(v string) *MtrTestMtrTestResponseResultData {
  s.DetectIpIspCode = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDetectIpPro(v string) *MtrTestMtrTestResponseResultData {
  s.DetectIpPro = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDetectIpProCode(v string) *MtrTestMtrTestResponseResultData {
  s.DetectIpProCode = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDetectResult(v string) *MtrTestMtrTestResponseResultData {
  s.DetectResult = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetDoneTime(v int64) *MtrTestMtrTestResponseResultData {
  s.DoneTime = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetStatus(v int32) *MtrTestMtrTestResponseResultData {
  s.Status = &v
  return s
}

func (s *MtrTestMtrTestResponseResultData) SetItems(v []*MtrTestMtrTestResponseResultDataItems) *MtrTestMtrTestResponseResultData {
  s.Items = v
  return s
}

type MtrTestMtrTestResponseResultDataItems struct     {
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

func (s MtrTestMtrTestResponseResultDataItems) String() string {
  return tea.Prettify(s)
}

func (s MtrTestMtrTestResponseResultDataItems) GoString() string {
  return s.String()
}

func (s *MtrTestMtrTestResponseResultDataItems) SetNum(v int32) *MtrTestMtrTestResponseResultDataItems {
  s.Num = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetIp(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Ip = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetLocation(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Location = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetAvg(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Avg = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetBest(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Best = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetWrst(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Wrst = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetLast(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Last = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetLoss(v string) *MtrTestMtrTestResponseResultDataItems {
  s.Loss = &v
  return s
}

func (s *MtrTestMtrTestResponseResultDataItems) SetSnt(v string) *MtrTestMtrTestResponseResultDataItems {
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
  Result []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceResponse) SetResult(v []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult) *ReportAvgSpeedDomainIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"ispData", "zh_CN":"ISP数据"}
  IspData []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult) SetDomain(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult) SetIspData(v []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData struct     {
  // {"en":"Internet service providers", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"provinceData", "zh_CN":"省份数据"}
  ProvinceData []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {"en":"ISP chinese name", "zh_CN":"省份中文名称"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"details", "zh_CN":"详情数据"}
  Details []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData) SetDetails(v []*ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceData {
  s.Details = v
  return s
}

type ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails struct     {
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

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetTimestamp(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetAvgSpeed(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.AvgSpeed = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetAvgResponseTime(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.AvgResponseTime = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetAvgFirstPacketTime(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
  s.AvgFirstPacketTime = &v
  return s
}

func (s *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails) SetTotalResponseTime(v string) *ReportAvgSpeedDomainIspProvinceServiceReportAvgSpeedDomainIspProvinceServiceResponseResultIspDataProvinceDataDetails {
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

type GetLiveStreamPushingStatusResponse struct {
  // {'en':'The time of the data returned', 'zh_CN':'返回的数据的时间'}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {'en':'Number of data items. 0 is returned if there is no data', 'zh_CN':'数据条数，无数据返回0'}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'推流信息数据集合'}
  DataValue []*GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
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

func (s *GetLiveStreamPushingStatusResponse) SetDataValue(v []*GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) *GetLiveStreamPushingStatusResponse {
  s.DataValue = v
  return s
}

type GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue struct     {
  // {'en':'Push stream CDN node, that is, the IP address of edge node of the received data source, separated by multiple commas', 'zh_CN':'推流cdn节点，即收到数据源的edge节点IP地址，多个逗号分隔'}
  Deployaddress *string `json:"deployaddress,omitempty" xml:"deployaddress,omitempty" require:"true"`
  // {'en':'Push user IP (data source) address, separated by multiple commas', 'zh_CN':'推流用户ip（数据源）地址，多个逗号分隔'}
  Inaddress *string `json:"inaddress,omitempty" xml:"inaddress,omitempty" require:"true"`
  // {'en':'The channel of anchor', 'zh_CN':'主播流名'}
  Streamname *string `json:"streamname,omitempty" xml:"streamname,omitempty" require:"true"`
  // {'en':'Anchor Current encoding frame rate.Unit: fps', 'zh_CN':'主播当前编码帧率，单位:fps'}
  Fps *int64 `json:"fps,omitempty" xml:"fps,omitempty" require:"true"`
  // {'en':'Current frame loss rate of anchor.Unit: fps', 'zh_CN':'主播当前丢帧率，单位:fps'}
  Lfr *float64 `json:"lfr,omitempty" xml:"lfr,omitempty" require:"true"`
  // {'en':'Anchor Current bit rate.Unit: bps', 'zh_CN':'主播当前码率，单位:bps'}
  Inbandwidth *int64 `json:"inbandwidth,omitempty" xml:"inbandwidth,omitempty" require:"true"`
  // {'en':'Video timestamps are separated by multiple commas.Unit: ms', 'zh_CN':'视频时间戳 多个逗号分隔，单位:ms'}
  Videotmstmp *string `json:"videotmstmp,omitempty" xml:"videotmstmp,omitempty" require:"true"`
  // {'en':'Audio timestamp, separated by multiple commas.Unit: ms', 'zh_CN':'音频时间戳，多个逗号分隔，单位:ms'}
  Audiotmstmp *string `json:"audiotmstmp,omitempty" xml:"audiotmstmp,omitempty" require:"true"`
}

func (s GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetDeployaddress(v string) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Deployaddress = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetInaddress(v string) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Inaddress = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetStreamname(v string) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Streamname = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetFps(v int64) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Fps = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetLfr(v float64) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Lfr = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetInbandwidth(v int64) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Inbandwidth = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetVideotmstmp(v string) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Videotmstmp = &v
  return s
}

func (s *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue) SetAudiotmstmp(v string) *GetLiveStreamPushingStatusGetLiveStreamPushingStatusResponseDataValue {
  s.Audiotmstmp = &v
  return s
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
  // {'en':'Push domain(multiple domains supported, separated by commas)', 'zh_CN':'推流域名（支持多个域名，以逗号分隔）'}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {'en':'Time, eg: 20160527152300, if not filled in, the current time -5 minutes', 'zh_CN':'时间，eg：20160527152300，不填时为当前时间-5分钟'}
  T *int64 `json:"t,omitempty" xml:"t,omitempty"`
  // {'en':'Push channel URL (multiple push channel URLs are supported, separated by commas)', 'zh_CN':'推流流名(支持多个推流流名，以逗号分隔)'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {'en':'Query interval. The value can be 10 or 60. The default value is 60
  // When g is 10, query the data of the nearest whole 10 seconds to time t
  // When g is 60, query the data of the nearest whole minute to time t', 'zh_CN':'查询间隔，可选值10、60，默认60
  // 当g为10时，查询距离时间t最近的整10秒点数据
  // 当g为60时，查询距离时间t最近的整分钟点数据'}
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

func (s *GetLiveStreamPushingStatusParameters) SetT(v int64) *GetLiveStreamPushingStatusParameters {
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

type GetLiveStreamPushingStatusRequestHeader struct {
}

func (s GetLiveStreamPushingStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusRequestHeader) GoString() string {
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
  Data []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportOnlineNumIspProvinceServiceResponse) SetData(v []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData) *ReportOnlineNumIspProvinceServiceResponse {
  s.Data = v
  return s
}

type ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData) SetDomain(v string) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData) SetIspData(v []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseData {
  s.IspData = v
  return s
}

type ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData) SetIsp(v string) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData {
  s.Isp = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData) SetProvinceData(v []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspData {
  s.ProvinceData = v
  return s
}

type ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  Details []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) SetProvince(v string) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData) SetDetails(v []*ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceData {
  s.Details = v
  return s
}

type ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails struct     {
  // {"en":"Time, the format is yyyy-MM-dd HH:mm", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"online users", "zh_CN":"在线人数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) GoString() string {
  return s.String()
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetTimestamp(v string) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails {
  s.Timestamp = &v
  return s
}

func (s *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails) SetValue(v string) *ReportOnlineNumIspProvinceServiceReportOnlineNumIspProvinceServiceResponseDataIspDataProvinceDataDetails {
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
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
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

func (s ConcurrentSessionRequest) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionRequest) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionRequest) SetCust(v string) *ConcurrentSessionRequest {
  s.Cust = &v
  return s
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

type ConcurrentSessionResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *ConcurrentSessionConcurrentSessionResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s ConcurrentSessionResponse) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionResponse) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionResponse) SetProvider(v *ConcurrentSessionConcurrentSessionResponseProvider) *ConcurrentSessionResponse {
  s.Provider = v
  return s
}

type ConcurrentSessionConcurrentSessionResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Date *ConcurrentSessionConcurrentSessionResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s ConcurrentSessionConcurrentSessionResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionConcurrentSessionResponseProvider) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionConcurrentSessionResponseProvider) SetName(v string) *ConcurrentSessionConcurrentSessionResponseProvider {
  s.Name = &v
  return s
}

func (s *ConcurrentSessionConcurrentSessionResponseProvider) SetType(v string) *ConcurrentSessionConcurrentSessionResponseProvider {
  s.Type = &v
  return s
}

func (s *ConcurrentSessionConcurrentSessionResponseProvider) SetDate(v *ConcurrentSessionConcurrentSessionResponseProviderDate) *ConcurrentSessionConcurrentSessionResponseProvider {
  s.Date = v
  return s
}

type ConcurrentSessionConcurrentSessionResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'concurrent', 'zh_CN':'明细数据'}
  Concurrent []*ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent `json:"concurrent,omitempty" xml:"concurrent,omitempty" require:"true" type:"Repeated"`
}

func (s ConcurrentSessionConcurrentSessionResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionConcurrentSessionResponseProviderDate) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionConcurrentSessionResponseProviderDate) SetStartdate(v string) *ConcurrentSessionConcurrentSessionResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *ConcurrentSessionConcurrentSessionResponseProviderDate) SetEnddate(v string) *ConcurrentSessionConcurrentSessionResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *ConcurrentSessionConcurrentSessionResponseProviderDate) SetConcurrent(v []*ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent) *ConcurrentSessionConcurrentSessionResponseProviderDate {
  s.Concurrent = v
  return s
}

type ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数据'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent) GoString() string {
  return s.String()
}

func (s *ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent) SetTime(v string) *ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent {
  s.Time = &v
  return s
}

func (s *ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent) SetText(v string) *ConcurrentSessionConcurrentSessionResponseProviderDateConcurrent {
  s.Text = &v
  return s
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

type ConcurrentSessionRequestHeader struct {
}

func (s ConcurrentSessionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ConcurrentSessionRequestHeader) GoString() string {
  return s.String()
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
  Provider *BandwidthAppaBandwidthAppaResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthAppaResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaResponse) GoString() string {
  return s.String()
}

func (s *BandwidthAppaResponse) SetProvider(v *BandwidthAppaBandwidthAppaResponseProvider) *BandwidthAppaResponse {
  s.Provider = v
  return s
}

type BandwidthAppaBandwidthAppaResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'APPA带宽数据'}
  Date *BandwidthAppaBandwidthAppaResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthAppaBandwidthAppaResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaBandwidthAppaResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthAppaBandwidthAppaResponseProvider) SetName(v string) *BandwidthAppaBandwidthAppaResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProvider) SetType(v string) *BandwidthAppaBandwidthAppaResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProvider) SetDate(v *BandwidthAppaBandwidthAppaResponseProviderDate) *BandwidthAppaBandwidthAppaResponseProvider {
  s.Date = v
  return s
}

type BandwidthAppaBandwidthAppaResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'appa', 'zh_CN':'appa带宽数据'}
  Appa []*BandwidthAppaBandwidthAppaResponseProviderDateAppa `json:"appa,omitempty" xml:"appa,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthAppaBandwidthAppaResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaBandwidthAppaResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDate) SetStartdate(v string) *BandwidthAppaBandwidthAppaResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDate) SetEnddate(v string) *BandwidthAppaBandwidthAppaResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDate) SetAppa(v []*BandwidthAppaBandwidthAppaResponseProviderDateAppa) *BandwidthAppaBandwidthAppaResponseProviderDate {
  s.Appa = v
  return s
}

type BandwidthAppaBandwidthAppaResponseProviderDateAppa struct     {
  // {'en':'type', 'zh_CN':'appa数据类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道名称'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthAppaBandwidthAppaResponseProviderDateAppa) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaBandwidthAppaResponseProviderDateAppa) GoString() string {
  return s.String()
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDateAppa) SetType(v string) *BandwidthAppaBandwidthAppaResponseProviderDateAppa {
  s.Type = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDateAppa) SetChannel(v string) *BandwidthAppaBandwidthAppaResponseProviderDateAppa {
  s.Channel = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDateAppa) SetBandwidth(v []*BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth) *BandwidthAppaBandwidthAppaResponseProviderDateAppa {
  s.Bandwidth = v
  return s
}

type BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点，格式 yyyy-MM-dd hh:mm:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽，单位Mbps'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth) SetTime(v string) *BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth) SetText(v string) *BandwidthAppaBandwidthAppaResponseProviderDateAppaBandwidth {
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
  Data []*ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportDomainUserAgentServiceResponse) SetData(v []*ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData) *ReportDomainUserAgentServiceResponse {
  s.Data = v
  return s
}

type ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData struct     {
  // {"en":"UA", "zh_CN":"UA"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {"en":"Flow value. Unit is MB and 2 digits of decimals are allowed.", "zh_CN":"流量,保留2位小数,单位MB"}
  Flow *int `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {"en":"Number of requests.", "zh_CN":"请求数"}
  Request *int `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData) SetUserAgent(v string) *ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData {
  s.UserAgent = &v
  return s
}

func (s *ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData) SetFlow(v int) *ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData {
  s.Flow = &v
  return s
}

func (s *ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData) SetRequest(v int) *ReportDomainUserAgentServiceReportDomainUserAgentServiceResponseData {
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
  // {"en":"Specifies the  transcoding types:
  // 1.optional values:h264,h265,zdgq_264,zdgq_265.
  // 2.If there are multiple inputs,use ';' as delemeter.
  // 3.If not specified,output data of  'h264' and 'h265'.
  // 4.If none of the input values is right,error message 'invalid transcodeType'  will be returned.", "zh_CN":"转码类型，值为h264、h265、zdgq_264或zdgq_265，多个转码类型用英文分号';'分隔开，不选或者为空时默认提供h264、h265的内容。当传入转码类型部分错误时，返回正确的类型的数据；当传入转码类型全部错误时，返回错误invalid transcodeType 。"}
  TranscodeType *string `json:"transcodeType,omitempty" xml:"transcodeType,omitempty"`
  // {"en":"Definition type, the value is LD480, SD720, HD1080, 2K, 4K, multiple definitions are separated by English semicolon ';'", "zh_CN":"清晰度类型，值为LD480、SD720、HD1080、2K、4K，多个清晰度用英文分号';'分隔开"}
  Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
  // {"en":"Greenwich time zone, the parameter format GMT+09:00 means East Nine District, GMT-09:00 means West Nine District, if not passed, the default is the local time zone (East Eight District)", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
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

func (s *QueryDailyLiveTranscodingDurationRequest) SetDefinition(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Definition = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetTimezone(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Timezone = &v
  return s
}

type QueryDailyLiveTranscodingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponse) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponse) SetProvider(v *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider) *QueryDailyLiveTranscodingDurationResponse {
  s.Provider = v
  return s
}

type QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'直播转码时长每日统计数据'}
  Date *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider) SetName(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider) SetType(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider) SetDate(v *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProvider {
  s.Date = v
  return s
}

type QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate) SetStartdate(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate) SetEnddate(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate) SetTranscoding(v *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Transcoding = v
  return s
}

type QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'live', 'zh_CN':'直播转码时长数据'}
  Live []*QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) SetName(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Name = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) SetLive(v []*QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Live = v
  return s
}

type QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive struct     {
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

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTime(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Time = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH264(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H264 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH265(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H265 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_264(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_264 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_265(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_265 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetVoice(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Voice = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTotal(v string) *QueryDailyLiveTranscodingDurationQueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
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
  Provider *ChannelValueChannelValueResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueResponse) GoString() string {
  return s.String()
}

func (s *ChannelValueResponse) SetProvider(v *ChannelValueChannelValueResponseProvider) *ChannelValueResponse {
  s.Provider = v
  return s
}

type ChannelValueChannelValueResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道流量数据'}
  Date *ChannelValueChannelValueResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueChannelValueResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueChannelValueResponseProvider) GoString() string {
  return s.String()
}

func (s *ChannelValueChannelValueResponseProvider) SetName(v string) *ChannelValueChannelValueResponseProvider {
  s.Name = &v
  return s
}

func (s *ChannelValueChannelValueResponseProvider) SetType(v string) *ChannelValueChannelValueResponseProvider {
  s.Type = &v
  return s
}

func (s *ChannelValueChannelValueResponseProvider) SetDate(v *ChannelValueChannelValueResponseProviderDate) *ChannelValueChannelValueResponseProvider {
  s.Date = v
  return s
}

type ChannelValueChannelValueResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'information', 'zh_CN':'信息集'}
  Information *ChannelValueChannelValueResponseProviderDateInformation `json:"information,omitempty" xml:"information,omitempty" require:"true" type:"Struct"`
}

func (s ChannelValueChannelValueResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueChannelValueResponseProviderDate) GoString() string {
  return s.String()
}

func (s *ChannelValueChannelValueResponseProviderDate) SetStartdate(v string) *ChannelValueChannelValueResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDate) SetEnddate(v string) *ChannelValueChannelValueResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDate) SetInformation(v *ChannelValueChannelValueResponseProviderDateInformation) *ChannelValueChannelValueResponseProviderDate {
  s.Information = v
  return s
}

type ChannelValueChannelValueResponseProviderDateInformation struct {
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

func (s ChannelValueChannelValueResponseProviderDateInformation) String() string {
  return tea.Prettify(s)
}

func (s ChannelValueChannelValueResponseProviderDateInformation) GoString() string {
  return s.String()
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetChannel(v string) *ChannelValueChannelValueResponseProviderDateInformation {
  s.Channel = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetChargeMethod(v string) *ChannelValueChannelValueResponseProviderDateInformation {
  s.ChargeMethod = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetAcceType(v string) *ChannelValueChannelValueResponseProviderDateInformation {
  s.AcceType = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetValue(v string) *ChannelValueChannelValueResponseProviderDateInformation {
  s.Value = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetUnit(v string) *ChannelValueChannelValueResponseProviderDateInformation {
  s.Unit = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetPeakValue(v string) *ChannelValueChannelValueResponseProviderDateInformation {
  s.PeakValue = &v
  return s
}

func (s *ChannelValueChannelValueResponseProviderDateInformation) SetPeakTime(v string) *ChannelValueChannelValueResponseProviderDateInformation {
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
  // {"en":"Start time: 
  // 1.Format is yyyyMMdd; 
  // 2.Must be smaller than the current system time; 
  // 3.Default value of current time is used if the field is not specified; 
  // 4.You can only query data for the last 6 months.
  // ", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-dd,例如,2021-08-10
  // 2.不能大于当前时间
  // 3.只能查询最近半年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 1. The time format is yyyy-MM-dd
  // 2. The end time must be greater than the start time. If the end time is greater than the current time, take the current time
  // 3. Both dateFrom and dateTo have not been passed, and the past 7 days are queried by default; if only one has not been passed, an exception will be thrown
  // 4. The maximum time interval allowed for query: 31 days, that is, the difference between dateFrom and dateTo cannot exceed 31 days", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-dd
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的7天;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"-", "zh_CN":"-"}
  DomainStream []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" type:"Repeated"`
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

func (s *ReportDomainStreamDurationServiceRequest) SetDomainStream(v []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream) *ReportDomainStreamDurationServiceRequest {
  s.DomainStream = v
  return s
}

type ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream struct     {
  // {"en":"Domain name: 
  // 1. The maximum number of domain names that can be transferred is 1 by default; 
  // 2. Automatically filter out invalid domain names (if an illegal domain name is transferred, it will be filtered out, and the query results will only return the data of valid domain names).", "zh_CN":"域名:
  // 1、可传递域名数量上限默认为1个;
  // 2、自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"stream name:
  // Just pass the publishing point + stream name, Example: live/test-20180101-test where live is a publishing point and test-20180101-test is a stream name", "zh_CN":"流名:
  // 只需要传发布点+流名,例如:live/test-20180101-test ,其中live是发布点，test-20180101-test是流名"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream) SetDomain(v string) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream) SetStream(v []*string) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceRequestDomainStream {
  s.Stream = v
  return s
}

type ReportDomainStreamDurationServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the result of the request", "zh_CN":"请求结果的详细数据"}
  Data []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportDomainStreamDurationServiceResponse) SetData(v []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData) *ReportDomainStreamDurationServiceResponse {
  s.Data = v
  return s
}

type ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"-", "zh_CN":"-"}
  StreamList []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList `json:"streamList,omitempty" xml:"streamList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData) SetDomain(v string) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData) SetStreamList(v []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseData {
  s.StreamList = v
  return s
}

type ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList struct     {
  // {"en":"domain + publishing point + stream name", "zh_CN":"流名(域名+发布点+流名)"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"The sum of the flow duration of the flow name in the corresponding time period, in milliseconds", "zh_CN":"对应时间段内流名推流时长之和,单位为毫秒"}
  SumTime *int `json:"sumTime,omitempty" xml:"sumTime,omitempty" require:"true"`
  // {"en":"-", "zh_CN":"-"}
  DurationDetailList []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList `json:"durationDetailList,omitempty" xml:"durationDetailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList) SetStream(v string) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList {
  s.Stream = &v
  return s
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList) SetSumTime(v int) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList {
  s.SumTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList) SetDurationDetailList(v []*ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamList {
  s.DurationDetailList = v
  return s
}

type ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList struct     {
  // {"en":"Stream start time", "zh_CN":"推流起始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Stream end time", "zh_CN":"推流终止时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Streaming duration, in milliseconds", "zh_CN":"推流时长,单位为毫秒"}
  Duration *int `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetStartTime(v string) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.StartTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetEndTime(v string) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.EndTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetDuration(v int) *ReportDomainStreamDurationServiceReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.Duration = &v
  return s
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

type ReportDomainStreamDurationServiceRequestHeader struct {
}

func (s ReportDomainStreamDurationServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequestHeader) GoString() string {
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
  // {"en":"domains that been queried:
  // 1)If there are multiple inputs,use  ';' as separator.
  // 2)If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
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

type FlowAppaChannelResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *FlowAppaChannelFlowAppaChannelResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s FlowAppaChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelResponse) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelResponse) SetProvider(v *FlowAppaChannelFlowAppaChannelResponseProvider) *FlowAppaChannelResponse {
  s.Provider = v
  return s
}

type FlowAppaChannelFlowAppaChannelResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'请求数数据'}
  Date *FlowAppaChannelFlowAppaChannelResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s FlowAppaChannelFlowAppaChannelResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelFlowAppaChannelResponseProvider) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelFlowAppaChannelResponseProvider) SetName(v string) *FlowAppaChannelFlowAppaChannelResponseProvider {
  s.Name = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProvider) SetType(v string) *FlowAppaChannelFlowAppaChannelResponseProvider {
  s.Type = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProvider) SetDate(v *FlowAppaChannelFlowAppaChannelResponseProviderDate) *FlowAppaChannelFlowAppaChannelResponseProvider {
  s.Date = v
  return s
}

type FlowAppaChannelFlowAppaChannelResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'汇总'}
  Total *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal `json:"total,omitempty" xml:"total,omitempty" require:"true" type:"Struct"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s FlowAppaChannelFlowAppaChannelResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelFlowAppaChannelResponseProviderDate) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDate) SetStartdate(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDate) SetEnddate(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDate) SetTotal(v *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) *FlowAppaChannelFlowAppaChannelResponseProviderDate {
  s.Total = v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDate) SetChannel(v *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) *FlowAppaChannelFlowAppaChannelResponseProviderDate {
  s.Channel = v
  return s
}

type FlowAppaChannelFlowAppaChannelResponseProviderDateTotal struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'edgeup', 'zh_CN':'边缘上行总流量,单位Mbps'}
  Edgeup *string `json:"edgeup,omitempty" xml:"edgeup,omitempty" require:"true"`
  // {'en':'edgedown', 'zh_CN':'边缘下行总流量,单位Mbps'}
  Edgedown *string `json:"edgedown,omitempty" xml:"edgedown,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'汇总'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) SetName(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal {
  s.Name = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) SetEdgeup(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal {
  s.Edgeup = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) SetEdgedown(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal {
  s.Edgedown = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal) SetTotal(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateTotal {
  s.Total = &v
  return s
}

type FlowAppaChannelFlowAppaChannelResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'edgeup', 'zh_CN':'边缘上行总流量,单位Mbps'}
  Edgeup *string `json:"edgeup,omitempty" xml:"edgeup,omitempty" require:"true"`
  // {'en':'edgedown', 'zh_CN':'边缘下行总流量,单位Mbps'}
  Edgedown *string `json:"edgedown,omitempty" xml:"edgedown,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'汇总'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) SetName(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) SetEdgeup(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel {
  s.Edgeup = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) SetEdgedown(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel {
  s.Edgedown = &v
  return s
}

func (s *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel) SetTotal(v string) *FlowAppaChannelFlowAppaChannelResponseProviderDateChannel {
  s.Total = &v
  return s
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

type FlowAppaChannelRequestHeader struct {
}

func (s FlowAppaChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FlowAppaChannelRequestHeader) GoString() string {
  return s.String()
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
  Result []*ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUrlDlFinishServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUrlDlFinishServiceResponse) SetResult(v []*ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult) *ReportUrlDlFinishServiceResponse {
  s.Result = v
  return s
}

type ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult struct     {
  // {'en':'The top500 of url', 'zh_CN':'top500的url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'Number of successful downloads', 'zh_CN':'下载成功数'}
  Num *string `json:"num,omitempty" xml:"num,omitempty" require:"true"`
}

func (s ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult) SetUrl(v string) *ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult {
  s.Url = &v
  return s
}

func (s *ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult) SetNum(v string) *ReportUrlDlFinishServiceReportUrlDlFinishServiceResponseResult {
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
  Provider *CloudDirectDurationCloudDirectDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s CloudDirectDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationResponse) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationResponse) SetProvider(v *CloudDirectDurationCloudDirectDurationResponseProvider) *CloudDirectDurationResponse {
  s.Provider = v
  return s
}

type CloudDirectDurationCloudDirectDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Date *CloudDirectDurationCloudDirectDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s CloudDirectDurationCloudDirectDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationCloudDirectDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationCloudDirectDurationResponseProvider) SetName(v string) *CloudDirectDurationCloudDirectDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *CloudDirectDurationCloudDirectDurationResponseProvider) SetType(v string) *CloudDirectDurationCloudDirectDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *CloudDirectDurationCloudDirectDurationResponseProvider) SetDate(v *CloudDirectDurationCloudDirectDurationResponseProviderDate) *CloudDirectDurationCloudDirectDurationResponseProvider {
  s.Date = v
  return s
}

type CloudDirectDurationCloudDirectDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'totalDuration', 'zh_CN':'汇总'}
  TotalDuration *string `json:"totalDuration,omitempty" xml:"totalDuration,omitempty" require:"true"`
  // {'en':'result', 'zh_CN':'明细数据'}
  Result *CloudDirectDurationCloudDirectDurationResponseProviderDateResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s CloudDirectDurationCloudDirectDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationCloudDirectDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDate) SetStartdate(v string) *CloudDirectDurationCloudDirectDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDate) SetEnddate(v string) *CloudDirectDurationCloudDirectDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDate) SetTotalDuration(v string) *CloudDirectDurationCloudDirectDurationResponseProviderDate {
  s.TotalDuration = &v
  return s
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDate) SetResult(v *CloudDirectDurationCloudDirectDurationResponseProviderDateResult) *CloudDirectDurationCloudDirectDurationResponseProviderDate {
  s.Result = v
  return s
}

type CloudDirectDurationCloudDirectDurationResponseProviderDateResult struct {
  // {'en':'duration', 'zh_CN':'明细数据'}
  Duration []*CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration `json:"duration,omitempty" xml:"duration,omitempty" require:"true" type:"Repeated"`
}

func (s CloudDirectDurationCloudDirectDurationResponseProviderDateResult) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationCloudDirectDurationResponseProviderDateResult) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDateResult) SetDuration(v []*CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration) *CloudDirectDurationCloudDirectDurationResponseProviderDateResult {
  s.Duration = v
  return s
}

type CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数据'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration) String() string {
  return tea.Prettify(s)
}

func (s CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration) GoString() string {
  return s.String()
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration) SetTime(v string) *CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration {
  s.Time = &v
  return s
}

func (s *CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration) SetText(v string) *CloudDirectDurationCloudDirectDurationResponseProviderDateResultDuration {
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
  Provider *BandwidthUploadBandwidthUploadResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthUploadResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadResponse) GoString() string {
  return s.String()
}

func (s *BandwidthUploadResponse) SetProvider(v *BandwidthUploadBandwidthUploadResponseProvider) *BandwidthUploadResponse {
  s.Provider = v
  return s
}

type BandwidthUploadBandwidthUploadResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *BandwidthUploadBandwidthUploadResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthUploadBandwidthUploadResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadBandwidthUploadResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthUploadBandwidthUploadResponseProvider) SetName(v string) *BandwidthUploadBandwidthUploadResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthUploadBandwidthUploadResponseProvider) SetType(v string) *BandwidthUploadBandwidthUploadResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthUploadBandwidthUploadResponseProvider) SetDate(v *BandwidthUploadBandwidthUploadResponseProviderDate) *BandwidthUploadBandwidthUploadResponseProvider {
  s.Date = v
  return s
}

type BandwidthUploadBandwidthUploadResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthUploadBandwidthUploadResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthUploadBandwidthUploadResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadBandwidthUploadResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDate) SetStartdate(v string) *BandwidthUploadBandwidthUploadResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDate) SetEnddate(v string) *BandwidthUploadBandwidthUploadResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDate) SetChannel(v *BandwidthUploadBandwidthUploadResponseProviderDateChannel) *BandwidthUploadBandwidthUploadResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthUploadBandwidthUploadResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthUploadBandwidthUploadResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadBandwidthUploadResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDateChannel) SetName(v string) *BandwidthUploadBandwidthUploadResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDateChannel) SetBandwidth(v []*BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth) *BandwidthUploadBandwidthUploadResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth) SetText(v string) *BandwidthUploadBandwidthUploadResponseProviderDateChannelBandwidth {
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
  Result *HttpTestHttpTestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s HttpTestResponse) String() string {
  return tea.Prettify(s)
}

func (s HttpTestResponse) GoString() string {
  return s.String()
}

func (s *HttpTestResponse) SetResult(v *HttpTestHttpTestResponseResult) *HttpTestResponse {
  s.Result = v
  return s
}

type HttpTestHttpTestResponseResult struct {
  // {'en':'status', 'zh_CN':'状态码'}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'error message', 'zh_CN':'异常信息'}
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty" require:"true"`
  // {'en':'test id', 'zh_CN':'任务 ID'}
  TestId *string `json:"testId,omitempty" xml:"testId,omitempty" require:"true"`
  // {'en':'url', 'zh_CN':'目标 URL'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'data','zh_CN':'探测数据'}
  Data []*HttpTestHttpTestResponseResultData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s HttpTestHttpTestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s HttpTestHttpTestResponseResult) GoString() string {
  return s.String()
}

func (s *HttpTestHttpTestResponseResult) SetStatus(v string) *HttpTestHttpTestResponseResult {
  s.Status = &v
  return s
}

func (s *HttpTestHttpTestResponseResult) SetErrorMsg(v string) *HttpTestHttpTestResponseResult {
  s.ErrorMsg = &v
  return s
}

func (s *HttpTestHttpTestResponseResult) SetTestId(v string) *HttpTestHttpTestResponseResult {
  s.TestId = &v
  return s
}

func (s *HttpTestHttpTestResponseResult) SetUrl(v string) *HttpTestHttpTestResponseResult {
  s.Url = &v
  return s
}

func (s *HttpTestHttpTestResponseResult) SetData(v []*HttpTestHttpTestResponseResultData) *HttpTestHttpTestResponseResult {
  s.Data = v
  return s
}

type HttpTestHttpTestResponseResultData struct     {
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

func (s HttpTestHttpTestResponseResultData) String() string {
  return tea.Prettify(s)
}

func (s HttpTestHttpTestResponseResultData) GoString() string {
  return s.String()
}

func (s *HttpTestHttpTestResponseResultData) SetId(v string) *HttpTestHttpTestResponseResultData {
  s.Id = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetTaskId(v string) *HttpTestHttpTestResponseResultData {
  s.TaskId = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetUrl(v string) *HttpTestHttpTestResponseResultData {
  s.Url = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDetectIp(v string) *HttpTestHttpTestResponseResultData {
  s.DetectIp = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDetectIpIsp(v string) *HttpTestHttpTestResponseResultData {
  s.DetectIpIsp = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDetectIpIspCode(v string) *HttpTestHttpTestResponseResultData {
  s.DetectIpIspCode = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDetectIpPro(v string) *HttpTestHttpTestResponseResultData {
  s.DetectIpPro = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDetectIpProCode(v string) *HttpTestHttpTestResponseResultData {
  s.DetectIpProCode = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetTargetIp(v string) *HttpTestHttpTestResponseResultData {
  s.TargetIp = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetTargetIpIsp(v string) *HttpTestHttpTestResponseResultData {
  s.TargetIpIsp = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetTargetIpPro(v string) *HttpTestHttpTestResponseResultData {
  s.TargetIpPro = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetTargetLocation(v string) *HttpTestHttpTestResponseResultData {
  s.TargetLocation = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDnsIp(v string) *HttpTestHttpTestResponseResultData {
  s.DnsIp = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetConnectTime(v int32) *HttpTestHttpTestResponseResultData {
  s.ConnectTime = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDnsTime(v int32) *HttpTestHttpTestResponseResultData {
  s.DnsTime = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetRedirectTime(v int32) *HttpTestHttpTestResponseResultData {
  s.RedirectTime = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetResponseTime(v int32) *HttpTestHttpTestResponseResultData {
  s.ResponseTime = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetUseTime(v int32) *HttpTestHttpTestResponseResultData {
  s.UseTime = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetDoneTime(v int64) *HttpTestHttpTestResponseResultData {
  s.DoneTime = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetErrorCode(v string) *HttpTestHttpTestResponseResultData {
  s.ErrorCode = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetFileSize(v int32) *HttpTestHttpTestResponseResultData {
  s.FileSize = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetRate(v float32) *HttpTestHttpTestResponseResultData {
  s.Rate = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetHttpCode(v string) *HttpTestHttpTestResponseResultData {
  s.HttpCode = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetHttpTestRequestHeader(v string) *HttpTestHttpTestResponseResultData {
  s.HttpTestRequestHeader = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetHttpTestResponseHeader(v string) *HttpTestHttpTestResponseResultData {
  s.HttpTestResponseHeader = &v
  return s
}

func (s *HttpTestHttpTestResponseResultData) SetStatus(v int32) *HttpTestHttpTestResponseResultData {
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
  // {"en":"Specifies the  transcoding types:
  // 1.optional values:h264,h265,zdgq_264,zdgq_265.
  // 2.If there are multiple inputs,use ';' as delemeter.
  // 3.If not specified,output data of  'h264' and 'h265'.
  // 4.If none of the input values is right,error message 'invalid transcodeType'  will be returned.", "zh_CN":"转码类型，值为h264、h265、zdgq_264或zdgq_265，多个转码类型用英文分号';'分隔开，不选或者为空时默认提供h264、h265的内容。当传入转码类型部分错误时，返回正确的类型的数据；当传入转码类型全部错误时，返回错误invalid transcodeType 。"}
  TranscodeType *string `json:"transcodeType,omitempty" xml:"transcodeType,omitempty"`
  // {"en":"Definition type,value is LD480,SD720,HD1080,2K,4K,8K,SD576,,multiple resolutions are separated by English semicolon.", "zh_CN":"清晰度类型，值为LD480、SD720、HD1080、2K、4K、8K、SD576，多个清晰度用英文分号“;”分隔开
  // "}
  Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
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

type Query5minLiveTranscodingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponse) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponse) SetProvider(v *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider) *Query5minLiveTranscodingDurationResponse {
  s.Provider = v
  return s
}

type Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'直播转码时长数据'}
  Date *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider) SetName(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider) SetType(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider) SetDate(v *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProvider {
  s.Date = v
  return s
}

type Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate) SetStartdate(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate) SetEnddate(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate) SetTranscoding(v *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDate {
  s.Transcoding = v
  return s
}

type Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'live', 'zh_CN':'直播转码时长数据'}
  Live []*Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding) SetName(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Name = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding) SetLive(v []*Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Live = v
  return s
}

type Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive struct     {
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

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTime(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Time = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH264(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH265(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_264(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_265(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetVoice(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Voice = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetCf_264(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Cf_264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetCf_265(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Cf_265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_2K(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_2K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_4K(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_4K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_8K(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_8K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_LD480(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_LD480 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_SD720(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_SD720 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_HD1080(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_HD1080 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_SD576(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_SD576 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTotal(v string) *Query5minLiveTranscodingDurationQuery5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
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
  Provider *PicProcessStatisticsPicProcessStatisticsResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponse) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponse) SetProvider(v *PicProcessStatisticsPicProcessStatisticsResponseProvider) *PicProcessStatisticsResponse {
  s.Provider = v
  return s
}

type PicProcessStatisticsPicProcessStatisticsResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Date *PicProcessStatisticsPicProcessStatisticsResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProvider) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProvider) SetName(v string) *PicProcessStatisticsPicProcessStatisticsResponseProvider {
  s.Name = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProvider) SetType(v string) *PicProcessStatisticsPicProcessStatisticsResponseProvider {
  s.Type = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProvider) SetResultType(v string) *PicProcessStatisticsPicProcessStatisticsResponseProvider {
  s.ResultType = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProvider) SetDate(v *PicProcessStatisticsPicProcessStatisticsResponseProviderDate) *PicProcessStatisticsPicProcessStatisticsResponseProvider {
  s.Date = v
  return s
}

type PicProcessStatisticsPicProcessStatisticsResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProviderDate) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDate) SetStartdate(v string) *PicProcessStatisticsPicProcessStatisticsResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDate) SetEnddate(v string) *PicProcessStatisticsPicProcessStatisticsResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDate) SetChannel(v *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel) *PicProcessStatisticsPicProcessStatisticsResponseProviderDate {
  s.Channel = v
  return s
}

type PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'picflowhit', 'zh_CN':'请求数数据'}
  Picflowhit []*PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit `json:"picflowhit,omitempty" xml:"picflowhit,omitempty" require:"true" type:"Repeated"`
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel) SetName(v string) *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel) SetPicflowhit(v []*PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit) *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannel {
  s.Picflowhit = v
  return s
}

type PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'请求数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit) SetTime(v string) *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit {
  s.Time = &v
  return s
}

func (s *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit) SetText(v string) *PicProcessStatisticsPicProcessStatisticsResponseProviderDateChannelPicflowhit {
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
  DomainStream []*ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportDomainStreamHlsOnlineServiceRequest) SetDomainStream(v []*ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream) *ReportDomainStreamHlsOnlineServiceRequest {
  s.DomainStream = v
  return s
}

type ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream struct     {
  // {"en":"domian", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"stream: 'publishing point'/'stream'; 
  // 					For example: live/test-20200101-test.flv ,'live' is the publishing point, 'test-20200101-test' is stream;
  // 					If the stream is not transmitted, all stream under the domain will be queried by default", "zh_CN":"流名:发布点/流名。
  // 					例如:live/test-20200101-test.flv ,其中live是发布点, test-20200101-test是流名;
  // 					不传,默认查询指定域名下的所有流的数据"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream) SetDomain(v string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream) SetStream(v []*string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceRequestDomainStream {
  s.Stream = v
  return s
}

type ReportDomainStreamHlsOnlineServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportDomainStreamHlsOnlineServiceResponse) SetData(v []*ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) *ReportDomainStreamHlsOnlineServiceResponse {
  s.Data = v
  return s
}

type ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Number of streams under domain name", "zh_CN":"域名下的流个数"}
  StreamCount *string `json:"streamCount,omitempty" xml:"streamCount,omitempty" require:"true"`
  // {"en":"The number of online people corresponding to the domain.The value is the cumulative number of online people of all stream under the domain", "zh_CN":"该频道下总的在线人数,值为频道下所有流名的在线人数累加"}
  TotalOnlineCount *string `json:"totalOnlineCount,omitempty" xml:"totalOnlineCount,omitempty" require:"true"`
  StreamDetails []*ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails `json:"streamDetails,omitempty" xml:"streamDetails,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) SetDomain(v string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) SetStreamCount(v string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData {
  s.StreamCount = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) SetTotalOnlineCount(v string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData {
  s.TotalOnlineCount = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData) SetStreamDetails(v []*ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseData {
  s.StreamDetails = v
  return s
}

type ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails struct     {
  // {"en":"stream", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"The number of online people corresponding to the stream", "zh_CN":"该流名对应的在线人数"}
  OnlineCount *string `json:"onlineCount,omitempty" xml:"onlineCount,omitempty" require:"true"`
}

func (s ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) SetStream(v string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails {
  s.Stream = &v
  return s
}

func (s *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails) SetOnlineCount(v string) *ReportDomainStreamHlsOnlineServiceReportDomainStreamHlsOnlineServiceResponseDataStreamDetails {
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
  Result []*ReportStreamListServiceReportStreamListServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStreamListServiceResponse) SetResult(v []*ReportStreamListServiceReportStreamListServiceResponseResult) *ReportStreamListServiceResponse {
  s.Result = v
  return s
}

type ReportStreamListServiceReportStreamListServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Domain list", "zh_CN":"流名列表"}
  StreamList []*string `json:"streamList,omitempty" xml:"streamList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStreamListServiceReportStreamListServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportStreamListServiceReportStreamListServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportStreamListServiceReportStreamListServiceResponseResult) SetDomain(v string) *ReportStreamListServiceReportStreamListServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportStreamListServiceReportStreamListServiceResponseResult) SetStreamList(v []*string) *ReportStreamListServiceReportStreamListServiceResponseResult {
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
  Provider *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponse) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponse) SetProvider(v *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) *QueryLiveRecordingDurationResponse {
  s.Provider = v
  return s
}

type QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'请明细数据'}
  Date *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) SetName(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) SetType(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) SetResultType(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider {
  s.ResultType = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider) SetDate(v *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProvider {
  s.Date = v
  return s
}

type QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate struct {
  // {'en':'start', 'zh_CN':'开始时间'}
  Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {'en':'end', 'zh_CN':'结束时间'}
  End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  RecordingTime *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime `json:"recordingTime,omitempty" xml:"recordingTime,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate) SetStart(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate {
  s.Start = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate) SetEnd(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate {
  s.End = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate) SetRecordingTime(v *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDate {
  s.RecordingTime = v
  return s
}

type QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Channel *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime) SetName(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime) SetChannel(v *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTime {
  s.Channel = v
  return s
}

type QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'明细数据'}
  Live []*QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) SetName(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) SetLive(v []*QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel {
  s.Live = v
  return s
}

type QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数据'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) SetTime(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive {
  s.Time = &v
  return s
}

func (s *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) SetText(v string) *QueryLiveRecordingDurationQueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive {
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
  Provider *HttpDnsStatisticsHttpDnsStatisticsResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponse) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponse) SetProvider(v *HttpDnsStatisticsHttpDnsStatisticsResponseProvider) *HttpDnsStatisticsResponse {
  s.Provider = v
  return s
}

type HttpDnsStatisticsHttpDnsStatisticsResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'解析量数据'}
  Date *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProvider) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProvider) SetName(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProvider {
  s.Name = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProvider) SetType(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProvider {
  s.Type = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProvider) SetResultType(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProvider {
  s.ResultType = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProvider) SetDate(v *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate) *HttpDnsStatisticsHttpDnsStatisticsResponseProvider {
  s.Date = v
  return s
}

type HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate) SetStartdate(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate) SetEnddate(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate) SetChannel(v *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDate {
  s.Channel = v
  return s
}

type HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'总解析量'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'httpdns', 'zh_CN':'解析量数据'}
  Httpdns []*HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns `json:"httpdns,omitempty" xml:"httpdns,omitempty" require:"true" type:"Repeated"`
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel) SetName(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel) SetTotal(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel {
  s.Total = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel) SetHttpdns(v []*HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannel {
  s.Httpdns = v
  return s
}

type HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'解析量'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns) SetTime(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns {
  s.Time = &v
  return s
}

func (s *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns) SetText(v string) *HttpDnsStatisticsHttpDnsStatisticsResponseProviderDateChannelHttpdns {
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
  Provider *WafHitWafHitResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s WafHitResponse) String() string {
  return tea.Prettify(s)
}

func (s WafHitResponse) GoString() string {
  return s.String()
}

func (s *WafHitResponse) SetProvider(v *WafHitWafHitResponseProvider) *WafHitResponse {
  s.Provider = v
  return s
}

type WafHitWafHitResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'waf 请求数数据'}
  Date *WafHitWafHitResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s WafHitWafHitResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s WafHitWafHitResponseProvider) GoString() string {
  return s.String()
}

func (s *WafHitWafHitResponseProvider) SetName(v string) *WafHitWafHitResponseProvider {
  s.Name = &v
  return s
}

func (s *WafHitWafHitResponseProvider) SetType(v string) *WafHitWafHitResponseProvider {
  s.Type = &v
  return s
}

func (s *WafHitWafHitResponseProvider) SetDate(v *WafHitWafHitResponseProviderDate) *WafHitWafHitResponseProvider {
  s.Date = v
  return s
}

type WafHitWafHitResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *WafHitWafHitResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s WafHitWafHitResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s WafHitWafHitResponseProviderDate) GoString() string {
  return s.String()
}

func (s *WafHitWafHitResponseProviderDate) SetStartdate(v string) *WafHitWafHitResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *WafHitWafHitResponseProviderDate) SetEnddate(v string) *WafHitWafHitResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *WafHitWafHitResponseProviderDate) SetChannel(v *WafHitWafHitResponseProviderDateChannel) *WafHitWafHitResponseProviderDate {
  s.Channel = v
  return s
}

type WafHitWafHitResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'waf 请求数数据'}
  Wafhit []*WafHitWafHitResponseProviderDateChannelWafhit `json:"wafhit,omitempty" xml:"wafhit,omitempty" require:"true" type:"Repeated"`
}

func (s WafHitWafHitResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s WafHitWafHitResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *WafHitWafHitResponseProviderDateChannel) SetName(v string) *WafHitWafHitResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *WafHitWafHitResponseProviderDateChannel) SetWafhit(v []*WafHitWafHitResponseProviderDateChannelWafhit) *WafHitWafHitResponseProviderDateChannel {
  s.Wafhit = v
  return s
}

type WafHitWafHitResponseProviderDateChannelWafhit struct     {
  // {'en':'time of every 5 duration,with format yyyy-mmm-dd hh:MM:ss', 'zh_CN':'waf请求数5分钟粒度时间，格式yyyy-mm-dd hh:MM:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'displaying the waf request.', 'zh_CN':'waf请求数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s WafHitWafHitResponseProviderDateChannelWafhit) String() string {
  return tea.Prettify(s)
}

func (s WafHitWafHitResponseProviderDateChannelWafhit) GoString() string {
  return s.String()
}

func (s *WafHitWafHitResponseProviderDateChannelWafhit) SetTime(v string) *WafHitWafHitResponseProviderDateChannelWafhit {
  s.Time = &v
  return s
}

func (s *WafHitWafHitResponseProviderDateChannelWafhit) SetText(v string) *WafHitWafHitResponseProviderDateChannelWafhit {
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

type QueryLiveStreamStatusResponse struct {
  // {'en':'The time of the data returned', 'zh_CN':'返回的数据的时间'}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {'en':'Number of data items. 0 is returned if there is no data', 'zh_CN':'数据条数，无数据返回0'}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {'en':'Total onlines', 'zh_CN':'总在线人数'}
  Histscount *int64 `json:"histscount,omitempty" xml:"histscount,omitempty" require:"true"`
  // {'en':'Total channel bandwidth', 'zh_CN':'总频道带宽'}
  Bandwidthcount *int64 `json:"bandwidthcount,omitempty" xml:"bandwidthcount,omitempty" require:"true"`
  // {'en':'Timestamp', 'zh_CN':'数据的时间戳,如果有传t,则等于t'}
  Datetime *int64 `json:"datetime,omitempty" xml:"datetime,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'数据集合'}
  DataValue []*QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryLiveStreamStatusResponse) SetDataValue(v []*QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) *QueryLiveStreamStatusResponse {
  s.DataValue = v
  return s
}

type QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue struct     {
  // {'en':'The channel of anchor', 'zh_CN':'主播流名'}
  Streamname *string `json:"streamname,omitempty" xml:"streamname,omitempty" require:"true"`
  // {'en':'IP address of the CDN node', 'zh_CN':'推流cdn节点IP'}
  Deployaddress *string `json:"deployaddress,omitempty" xml:"deployaddress,omitempty" require:"true"`
  // {'en':'Anchor exit Address', 'zh_CN':'主播出口地址'}
  Inaddress *string `json:"inaddress,omitempty" xml:"inaddress,omitempty" require:"true"`
  // {'en':'The number of online', 'zh_CN':'在线人数'}
  Hists *int64 `json:"hists,omitempty" xml:"hists,omitempty" require:"true"`
  // {'en':'Anchor Current bit rate (transcoding stream has no bit rate data) unit: BPS', 'zh_CN':'主播当前码率(转码流没有码率数据) 单位：bps'}
  Inbandwidth *int64 `json:"inbandwidth,omitempty" xml:"inbandwidth,omitempty" require:"true"`
  // {'en':'Channel current viewing bandwidth unit: BPS ', 'zh_CN':'频道当前观看带宽 单位：bps'}
  Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {'en':'Anchor delay (MS)', 'zh_CN':'主播延迟(ms)'}
  Delay *int64 `json:"delay,omitempty" xml:"delay,omitempty" require:"true"`
  // {'en':'Anchor Current encoding frame rate(fps)', 'zh_CN':'主播当前编码帧率(fps)'}
  Fps *int64 `json:"fps,omitempty" xml:"fps,omitempty" require:"true"`
  // {'en':'Current frame loss rate of anchor(fps)', 'zh_CN':'主播当前丢帧率(fps)'}
  Lfr *float64 `json:"lfr,omitempty" xml:"lfr,omitempty" require:"true"`
  // {'en':'Anchor raw frame rate(fps)', 'zh_CN':'主播原始帧率(fps)'}
  Ofr *int64 `json:"ofr,omitempty" xml:"ofr,omitempty" require:"true"`
  // {'en':'The resolution of the', 'zh_CN':'分辨率'}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {'en':'Video coding', 'zh_CN':'视频编码'}
  Video_codec *string `json:"video_codec,omitempty" xml:"video_codec,omitempty" require:"true"`
  // {'en':'Audio coding', 'zh_CN':'音频编码'}
  Audio_codec *string `json:"audio_codec,omitempty" xml:"audio_codec,omitempty" require:"true"`
  // {'en':'Keyframe interval', 'zh_CN':'关键帧间隔,有传expand才会返回'}
  Gop *int64 `json:"gop,omitempty" xml:"gop,omitempty" require:"true"`
}

func (s QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetStreamname(v string) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Streamname = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetDeployaddress(v string) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Deployaddress = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetInaddress(v string) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Inaddress = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetHists(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Hists = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetInbandwidth(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Inbandwidth = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetBandwidth(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Bandwidth = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetDelay(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Delay = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetFps(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Fps = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetLfr(v float64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Lfr = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetOfr(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Ofr = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetResolution(v string) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Resolution = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetVideo_codec(v string) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Video_codec = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetAudio_codec(v string) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Audio_codec = &v
  return s
}

func (s *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue) SetGop(v int64) *QueryLiveStreamStatusQueryLiveStreamStatusResponseDataValue {
  s.Gop = &v
  return s
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
  // {'en':'Domain (multiple domains supported, separated by commas)', 'zh_CN':'域名（支持多个域名，以逗号分隔）'}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {'en':'20160527152300, non-real-time indicates the current time -5 minutes, real-time indicates the current time -30 seconds
  // A:
  // (t is not transmitted, the current system time is obtained and rounded to second (for example, 2017/3/28 14:38:55 rounded to second 2017/3/28 14:38:50 rounded to second);) If the query interval g=10, the final time (dateFrom) is rounded in seconds (e.g. 2017/3/28 14:38:55 after rounded in seconds 2017/3/28 14:38:50); If the g! =10, rounded minutes: (e.g. 2017/3/28 14:38:55 rounded seconds 2017/3/28 14:38:00); DateTo = dataFrom + 9 (9 seconds)', 'zh_CN':'20160527152300，不填为当前时间-5分钟
  // 详解：
  // （t不传，获取当前系统时间，对秒取整(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:50)；）如果查询间隔g=10，最后获得的时间对秒取整(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:50)；如果g!=10，对分钟取整：(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:00);'}
  T *string `json:"t,omitempty" xml:"t,omitempty"`
  // {'en':'Channel URL (simple channel url are supported,eg:push1.test.com/test/test1)', 'zh_CN':'流名(支持单流名，如：push1.test.com/test/test1)'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {'en':'Optional value: push or pull. The default value is push
  // Push stands for watershed name
  // Pull indicates the name of a pull basin', 'zh_CN':'域名类型，可选值：push、pull，默认push
  // push代表推流域名
  // pull代表拉流域名'}
  D *string `json:"d,omitempty" xml:"d,omitempty"`
  // {'en':'The value can be true or false. The default value is false
  // Only realtime data is returned when realtime=true', 'zh_CN':'是否返回端口，可选值：true、false，默认false
  // '}
  Showport *string `json:"showport,omitempty" xml:"showport,omitempty"`
  // {'en':'Query interval. The value can be 10 or 60. The default value is 60
  // When g is 10, query the data of the nearest whole 10 seconds to time t
  // When g is 60, query the data of the nearest whole minute to time t', 'zh_CN':'查询间隔，可选值10、60，默认60
  // 当g为10时，查询距离时间t最近的整10秒点数据
  // 当g为60时，查询距离时间t最近的整分钟点数据'}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
  // {'en':'it is query realtime datas,default false', 'zh_CN':'是否返回实时数据，默认false'}
  Realtime *string `json:"realtime,omitempty" xml:"realtime,omitempty"`
  // {'en':'Data extension fields(Only supports non-real-time query), providing optional return fields:gop', 'zh_CN':'数据扩展字段(仅支持按非实时查)，提供可选返回字段：gop'}
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

type QueryLiveStreamStatusRequestHeader struct {
}

func (s QueryLiveStreamStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusRequestHeader) GoString() string {
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
  Provider *AttOverViewAttOverViewResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s AttOverViewResponse) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewResponse) GoString() string {
  return s.String()
}

func (s *AttOverViewResponse) SetProvider(v *AttOverViewAttOverViewResponseProvider) *AttOverViewResponse {
  s.Provider = v
  return s
}

type AttOverViewAttOverViewResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'安全防护安全状况汇总数据'}
  Date *AttOverViewAttOverViewResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s AttOverViewAttOverViewResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewAttOverViewResponseProvider) GoString() string {
  return s.String()
}

func (s *AttOverViewAttOverViewResponseProvider) SetName(v string) *AttOverViewAttOverViewResponseProvider {
  s.Name = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProvider) SetType(v string) *AttOverViewAttOverViewResponseProvider {
  s.Type = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProvider) SetResultType(v string) *AttOverViewAttOverViewResponseProvider {
  s.ResultType = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProvider) SetDate(v *AttOverViewAttOverViewResponseProviderDate) *AttOverViewAttOverViewResponseProvider {
  s.Date = v
  return s
}

type AttOverViewAttOverViewResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *AttOverViewAttOverViewResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s AttOverViewAttOverViewResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewAttOverViewResponseProviderDate) GoString() string {
  return s.String()
}

func (s *AttOverViewAttOverViewResponseProviderDate) SetStartdate(v string) *AttOverViewAttOverViewResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProviderDate) SetEnddate(v string) *AttOverViewAttOverViewResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProviderDate) SetChannel(v *AttOverViewAttOverViewResponseProviderDateChannel) *AttOverViewAttOverViewResponseProviderDate {
  s.Channel = v
  return s
}

type AttOverViewAttOverViewResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'overview', 'zh_CN':'安全防护安全状况汇总数据'}
  OverView []*AttOverViewAttOverViewResponseProviderDateChannelOverView `json:"over-view,omitempty" xml:"over-view,omitempty" require:"true" type:"Repeated"`
}

func (s AttOverViewAttOverViewResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewAttOverViewResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *AttOverViewAttOverViewResponseProviderDateChannel) SetName(v string) *AttOverViewAttOverViewResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProviderDateChannel) SetOverView(v []*AttOverViewAttOverViewResponseProviderDateChannelOverView) *AttOverViewAttOverViewResponseProviderDateChannel {
  s.OverView = v
  return s
}

type AttOverViewAttOverViewResponseProviderDateChannelOverView struct     {
  // {'en':'timestamp', 'zh_CN':'时间点，格式 yyyy-MM-dd hh:mm:ss'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'total request number.', 'zh_CN':'总请求数'}
  Hit *string `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
  // {'en':'total denied request number.', 'zh_CN':'总访问拒绝数'}
  Deny *string `json:"deny,omitempty" xml:"deny,omitempty" require:"true"`
}

func (s AttOverViewAttOverViewResponseProviderDateChannelOverView) String() string {
  return tea.Prettify(s)
}

func (s AttOverViewAttOverViewResponseProviderDateChannelOverView) GoString() string {
  return s.String()
}

func (s *AttOverViewAttOverViewResponseProviderDateChannelOverView) SetChannel(v string) *AttOverViewAttOverViewResponseProviderDateChannelOverView {
  s.Channel = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProviderDateChannelOverView) SetHit(v string) *AttOverViewAttOverViewResponseProviderDateChannelOverView {
  s.Hit = &v
  return s
}

func (s *AttOverViewAttOverViewResponseProviderDateChannelOverView) SetDeny(v string) *AttOverViewAttOverViewResponseProviderDateChannelOverView {
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
  Data []*ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ReportBandwidthLowDelayP2pServiceResponse) SetData(v []*ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData) *ReportBandwidthLowDelayP2pServiceResponse {
  s.Data = v
  return s
}

type ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData struct     {
  // {'en':'-', 'zh_CN':'域名，聚合全部域名数据不返回该字段'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'-', 'zh_CN':'请求结果的详细数据'}
  DetailList []*ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData) SetDomain(v string) *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData) SetDetailList(v []*ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList) *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseData {
  s.DetailList = v
  return s
}

type ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList struct     {
  // {'en':'-', 'zh_CN':'时间片,返回开始时间和结束时间包含的时间片。时间格式：yyyy-MM-dd HH:mm'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'-', 'zh_CN':'P2P带宽'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList) SetTimestamp(v string) *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList) SetValue(v string) *ReportBandwidthLowDelayP2pServiceReportBandwidthLowDelayP2pServiceResponseDataDetailList {
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
  Result *DnsTestDnsTestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s DnsTestResponse) String() string {
  return tea.Prettify(s)
}

func (s DnsTestResponse) GoString() string {
  return s.String()
}

func (s *DnsTestResponse) SetResult(v *DnsTestDnsTestResponseResult) *DnsTestResponse {
  s.Result = v
  return s
}

type DnsTestDnsTestResponseResult struct {
  // {'en':'status', 'zh_CN':'状态码'}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'error message', 'zh_CN':'错误信息'}
  ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty" require:"true"`
  // {'en':'test id', 'zh_CN':'任务 ID'}
  TestId *string `json:"testId,omitempty" xml:"testId,omitempty" require:"true"`
  // {'en':'domain', 'zh_CN':'目标域名'}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {'en':'data','zh_CN':'探测数据'}
  Data []*DnsTestDnsTestResponseResultData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s DnsTestDnsTestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s DnsTestDnsTestResponseResult) GoString() string {
  return s.String()
}

func (s *DnsTestDnsTestResponseResult) SetStatus(v string) *DnsTestDnsTestResponseResult {
  s.Status = &v
  return s
}

func (s *DnsTestDnsTestResponseResult) SetErrorMsg(v string) *DnsTestDnsTestResponseResult {
  s.ErrorMsg = &v
  return s
}

func (s *DnsTestDnsTestResponseResult) SetTestId(v string) *DnsTestDnsTestResponseResult {
  s.TestId = &v
  return s
}

func (s *DnsTestDnsTestResponseResult) SetHost(v string) *DnsTestDnsTestResponseResult {
  s.Host = &v
  return s
}

func (s *DnsTestDnsTestResponseResult) SetData(v []*DnsTestDnsTestResponseResultData) *DnsTestDnsTestResponseResult {
  s.Data = v
  return s
}

type DnsTestDnsTestResponseResultData struct     {
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
  AnswerList []*DnsTestDnsTestResponseResultDataAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" require:"true" type:"Repeated"`
}

func (s DnsTestDnsTestResponseResultData) String() string {
  return tea.Prettify(s)
}

func (s DnsTestDnsTestResponseResultData) GoString() string {
  return s.String()
}

func (s *DnsTestDnsTestResponseResultData) SetId(v string) *DnsTestDnsTestResponseResultData {
  s.Id = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetTaskId(v string) *DnsTestDnsTestResponseResultData {
  s.TaskId = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetQuestionUrl(v string) *DnsTestDnsTestResponseResultData {
  s.QuestionUrl = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDetectIp(v string) *DnsTestDnsTestResponseResultData {
  s.DetectIp = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDetectIpIsp(v string) *DnsTestDnsTestResponseResultData {
  s.DetectIpIsp = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDetectIpIspCode(v string) *DnsTestDnsTestResponseResultData {
  s.DetectIpIspCode = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDetectIpPro(v string) *DnsTestDnsTestResponseResultData {
  s.DetectIpPro = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDetectIpProCode(v string) *DnsTestDnsTestResponseResultData {
  s.DetectIpProCode = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetServer(v string) *DnsTestDnsTestResponseResultData {
  s.Server = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDetectTime(v string) *DnsTestDnsTestResponseResultData {
  s.DetectTime = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetDoneTime(v int64) *DnsTestDnsTestResponseResultData {
  s.DoneTime = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetQueryTime(v int32) *DnsTestDnsTestResponseResultData {
  s.QueryTime = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetStatus(v int32) *DnsTestDnsTestResponseResultData {
  s.Status = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetStatusCode(v string) *DnsTestDnsTestResponseResultData {
  s.StatusCode = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetAnswer(v string) *DnsTestDnsTestResponseResultData {
  s.Answer = &v
  return s
}

func (s *DnsTestDnsTestResponseResultData) SetAnswerList(v []*DnsTestDnsTestResponseResultDataAnswerList) *DnsTestDnsTestResponseResultData {
  s.AnswerList = v
  return s
}

type DnsTestDnsTestResponseResultDataAnswerList struct     {
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

func (s DnsTestDnsTestResponseResultDataAnswerList) String() string {
  return tea.Prettify(s)
}

func (s DnsTestDnsTestResponseResultDataAnswerList) GoString() string {
  return s.String()
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetId(v string) *DnsTestDnsTestResponseResultDataAnswerList {
  s.Id = &v
  return s
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetTestId(v string) *DnsTestDnsTestResponseResultDataAnswerList {
  s.TestId = &v
  return s
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetAddress(v string) *DnsTestDnsTestResponseResultDataAnswerList {
  s.Address = &v
  return s
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetDnsClass(v string) *DnsTestDnsTestResponseResultDataAnswerList {
  s.DnsClass = &v
  return s
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetDnsType(v string) *DnsTestDnsTestResponseResultDataAnswerList {
  s.DnsType = &v
  return s
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetResult(v string) *DnsTestDnsTestResponseResultDataAnswerList {
  s.Result = &v
  return s
}

func (s *DnsTestDnsTestResponseResultDataAnswerList) SetTtl(v int32) *DnsTestDnsTestResponseResultDataAnswerList {
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

type QueryOnlineViewerCountResponse struct {
  // {'en':'Total number of online users. This parameter is displayed only when from and to are empty', 'zh_CN':'在线总人数，仅当查询时间点，即from和to为空时才显示'}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {'en':'Total abnormal online number, only need customers return', 'zh_CN':'异常总在线人数，只对有需要客户进行返回'}
  ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty" require:"true"`
  // {'en':'The number of data items is displayed only when from and to are empty.', 'zh_CN':'数据条数，仅当查询时间点，即from和to为空时才显示'}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {'en':'The time of the data returned', 'zh_CN':'返回的数据的时间'}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {'en':'', 'zh_CN':'数据集合'}
  DataValue []*QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryOnlineViewerCountResponse) SetDataValue(v []*QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) *QueryOnlineViewerCountResponse {
  s.DataValue = v
  return s
}

type QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue struct     {
  // {'en':'Stream name', 'zh_CN':'流名'}
  Prog *string `json:"prog,omitempty" xml:"prog,omitempty" require:"true"`
  // {'en':'Time is displayed only if the query time range (channel, FROM, and to parameters) is not empty', 'zh_CN':'时间，仅当查询时间范围，即channel，from和to参数不为空时才显示'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'The number of online', 'zh_CN':'在线人数'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'Abnormal online number, only need customers return', 'zh_CN':'异常在线人数，只对需要客户进行返回'}
  ErrorValue *int64 `json:"errorValue,omitempty" xml:"errorValue,omitempty" require:"true"`
}

func (s QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) SetProg(v string) *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue {
  s.Prog = &v
  return s
}

func (s *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) SetTime(v string) *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue {
  s.Time = &v
  return s
}

func (s *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) SetValue(v int64) *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue {
  s.Value = &v
  return s
}

func (s *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue) SetErrorValue(v int64) *QueryOnlineViewerCountQueryOnlineViewerCountResponseDataValue {
  s.ErrorValue = &v
  return s
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
  // {'en':'Domain (multiple domains supported, separated by commas)', 'zh_CN':'域名（支持多个域名，以逗号分隔）'}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {'en':'Time,eg:20160527152300.If the parameter is not specified, the value is 3 minutes', 'zh_CN':'时间，eg：20160527152300，不填为当前时间-3分钟'}
  T *string `json:"t,omitempty" xml:"t,omitempty"`
  // {'en':'The domain type can be pull or push. If this parameter is not specified, the default value is pull', 'zh_CN':'域名类型，pull或push，不填时默认为pull'}
  D *string `json:"d,omitempty" xml:"d,omitempty"`
  // {'en':'Channel URL(single channel query only),It is not recommended to query with this parameter, and the performance of range query is poor.', 'zh_CN':'频道URL(仅支持单频道查询),不建议带该参数查询，范围查询性能较差'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {'en':'Start time, eg: 20160803103500, when both from and to are filled or left blank, channel parameter is mandatory. The query time span is two hours at most. If the query time exceeds two hours, the system queries the data within two hours from the start time and the number of online users within the last 7 days', 'zh_CN':'开始时间，eg: 20160803103500，from和to都填或都不填，都填时channel参数必填，查询时间跨度最大为两个小时，如果超过两个小时，将查询开始时间两个小时内的数据，可查近7天内在线人数数据'}
  From *string `json:"from,omitempty" xml:"from,omitempty"`
  // {'en':'End time, eg: 20160803103900, when both from and to are filled or left blank, channel parameter is mandatory. The query time span is two hours at most. If the query time exceeds two hours, the system queries the data within two hours from the start time and the number of online users in the last 7 days', 'zh_CN':'结束时间，eg: 20160803103900，from和to都填或都不填，都填时channel参数必填，查询时间跨度最大为两个小时，如果超过两个小时，将查询开始时间两个小时内的数据，可查近7天内在线人数数据'}
  To *string `json:"to,omitempty" xml:"to,omitempty"`
  // {'en':'Query interval, optional value: 10, 60s.
  // When g is 10, the number of online users every 10 seconds is queried;
  // When g is 60, the number of online users at the whole minute within the time range is queried.', 'zh_CN':'查询间隔，可选值10、60s，
  // 当g为10时，查询时间范围内每10秒的在线人数
  // 当g为60时，查询时间范围内整分钟点对应的在线人数'}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
  // {'en':'it is query realtime datas,default false', 'zh_CN':'是否返回实时数据，默认false'}
  Realtime *string `json:"realtime,omitempty" xml:"realtime,omitempty"`
  // {'en':'The default value is false. If the value is true, the domain data is split. If the value is false, the domain data is merged', 'zh_CN':'域名拆分控制，默认为false，为true时，拆分域名数据，为false时，合并域名数据'}
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

type QueryOnlineViewerCountRequestHeader struct {
}

func (s QueryOnlineViewerCountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountRequestHeader) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountResponseHeader struct {
}

func (s QueryOnlineViewerCountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponseHeader) GoString() string {
  return s.String()
}




