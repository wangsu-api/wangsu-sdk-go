package reportiplist

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCdnIpListRequest struct {
  // {"en":"Domain list.
  // 1. All domains are queried if this field is not specified;
  // 2. Number of domains can be adjusted depending on different accounts. The default value is 20(this limit applies to the empty value);", "zh_CN":"域名列表
  // 域名个数限制根据账号可调，默认为20个（不传递时同样受此限制）；"}
  QueryCdnIpListDomainList *QueryCdnIpListDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryCdnIpListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListRequest) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListRequest) SetDomainList(v *QueryCdnIpListDomainList) *QueryCdnIpListRequest {
  s.QueryCdnIpListDomainList = v
  return s
}

type QueryCdnIpListDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnIpListDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListDomainList) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListDomainList) SetDomainName(v []*string) *QueryCdnIpListDomainList {
  s.DomainName = v
  return s
}

type QueryCdnIpListResponse struct {
  // {"en":"domainServerList", "zh_CN":"CDN服务IP数据"}
  Result []*QueryCdnIpListQueryCdnIpListResponseResult `json:"domain-server-list,omitempty" xml:"domain-server-list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnIpListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListResponse) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListResponse) SetResult(v []*QueryCdnIpListQueryCdnIpListResponseResult) *QueryCdnIpListResponse {
  s.Result = v
  return s
}

type QueryCdnIpListQueryCdnIpListResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"serverList", "zh_CN":"服务数据"}
  ServerList []*QueryCdnIpListQueryCdnIpListResponseResultServerList `json:"server-list,omitempty" xml:"server-list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnIpListQueryCdnIpListResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListQueryCdnIpListResponseResult) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListQueryCdnIpListResponseResult) SetDomainName(v string) *QueryCdnIpListQueryCdnIpListResponseResult {
  s.DomainName = &v
  return s
}

func (s *QueryCdnIpListQueryCdnIpListResponseResult) SetServerList(v []*QueryCdnIpListQueryCdnIpListResponseResultServerList) *QueryCdnIpListQueryCdnIpListResponseResult {
  s.ServerList = v
  return s
}

type QueryCdnIpListQueryCdnIpListResponseResultServerList struct     {
  // {"en":"Server node IP", "zh_CN":"覆盖节点IP"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
}

func (s QueryCdnIpListQueryCdnIpListResponseResultServerList) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListQueryCdnIpListResponseResultServerList) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListQueryCdnIpListResponseResultServerList) SetServer(v string) *QueryCdnIpListQueryCdnIpListResponseResultServerList {
  s.Server = &v
  return s
}

type QueryCdnIpListPaths struct {
}

func (s QueryCdnIpListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListPaths) GoString() string {
  return s.String()
}

type QueryCdnIpListParameters struct {
}

func (s QueryCdnIpListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListParameters) GoString() string {
  return s.String()
}

type QueryCdnIpListRequestHeader struct {
}

func (s QueryCdnIpListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListRequestHeader) GoString() string {
  return s.String()
}

type QueryCdnIpListResponseHeader struct {
}

func (s QueryCdnIpListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListResponseHeader) GoString() string {
  return s.String()
}




type ReportServerIpIspProvinceServiceRequest struct {
  // {'en':'Domain:
  // 1.the maximum number of transitive domain names is 20 by default.
  // 2.Automatically filter out illegal domain names (e.g. passing illegal domain names will be filtered out, and the search results will only return the legal domain name data)', 'zh_CN':'域名列表
  // 1. 可传递域名数量上限默认为20个（可联系技术支持调整）。未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）;
  // 2. 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）;'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'ISP.Default Query for All ISPs', 'zh_CN':'运营商，不传默认查询全部运营商；'}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {'en':'Province, do not pass the default query for all provinces; optional province information:
  // anhui,beijing,chongqing,fujian,guangdong,gansu,guangxi,guizhou,henan,hubei,hebei,hainan,heilongjiang,
  // hunan,jilin,jiangsu,jiangxi,liaoning,neimenggu,ningxia,qinghai,sichuan,shandong,shanghai,shanxi2 (Shaanxi),shanxi1 (Shanxi)
  // ,tianjin,xinjiang,xizang,yunnan,zhejiang,qita (other)', 'zh_CN':'省份，不传默认查询全部省份；可选省份信息:
  // anhui,beijing,chongqing,fujian,guangdong,gansu,guangxi,guizhou,henan,hubei,hebei,hainan,heilongjiang,
  // hunan,jilin,jiangsu,jiangxi,liaoning,neimenggu,ningxia,qinghai,sichuan,shandong,shanghai,shanxi2（陕西）,shanxi1（山西）
  // ,tianjin,xinjiang,xizang,yunnan,zhejiang,qita（其它）'}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"Group dimension:
  // 1.Optional values:domain,province,isp,you can pass in single or multiple values
  // 2.The detailed data will be displayed according to the dimension.
  // 3.The order of entering parameters does not affect the order of returned results", "zh_CN":"分组维度
  // 1. 可选值为domain、province、isp，可传入单个或多个值；
  // 2. 有传入则按照该维度展示明细数据；
  // 3. 返回结果层级顺序固定，入参顺序不影响返回结果顺序。例如：'groupBy': ['domain','province']与'groupBy': ['province','domain']返回结果一样。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportServerIpIspProvinceServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportServerIpIspProvinceServiceRequest) SetDomain(v []*string) *ReportServerIpIspProvinceServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportServerIpIspProvinceServiceRequest) SetIsp(v []*string) *ReportServerIpIspProvinceServiceRequest {
  s.Isp = v
  return s
}

func (s *ReportServerIpIspProvinceServiceRequest) SetProvince(v []*string) *ReportServerIpIspProvinceServiceRequest {
  s.Province = v
  return s
}

func (s *ReportServerIpIspProvinceServiceRequest) SetGroupBy(v []*string) *ReportServerIpIspProvinceServiceRequest {
  s.GroupBy = v
  return s
}

type ReportServerIpIspProvinceServiceResponse struct {
  Result []*ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpIspProvinceServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportServerIpIspProvinceServiceResponse) SetResult(v []*ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult) *ReportServerIpIspProvinceServiceResponse {
  s.Result = v
  return s
}

type ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult) SetDomain(v string) *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult) SetIspData(v []*ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData) *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResult {
  s.IspData = v
  return s
}

type ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData struct     {
  // {'en':'ISP', 'zh_CN':'运营商'}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData) SetIsp(v string) *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData) SetProvinceData(v []*ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData) *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData struct     {
  // {'en':'Province', 'zh_CN':'省份'}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {'en':'IP list of the covered node', 'zh_CN':'覆盖节点IP列表'}
  ServerIpData []*string `json:"serverIpData,omitempty" xml:"serverIpData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData) SetProvince(v string) *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData) SetServerIpData(v []*string) *ReportServerIpIspProvinceServiceReportServerIpIspProvinceServiceResponseResultIspDataProvinceData {
  s.ServerIpData = v
  return s
}

type ReportServerIpIspProvinceServicePaths struct {
}

func (s ReportServerIpIspProvinceServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServicePaths) GoString() string {
  return s.String()
}

type ReportServerIpIspProvinceServiceParameters struct {
}

func (s ReportServerIpIspProvinceServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceParameters) GoString() string {
  return s.String()
}

type ReportServerIpIspProvinceServiceRequestHeader struct {
}

func (s ReportServerIpIspProvinceServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportServerIpIspProvinceServiceResponseHeader struct {
}

func (s ReportServerIpIspProvinceServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpIspProvinceServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportServerIpExistFlowServiceRequest struct {
  // {'en':'domain:
  // 
  // 1.The maximum number of passable domains is 20 by default (you can contact technical support tp adjust it).
  // 2.Automatically filter out illegal domains (e.g. passing illegal domains will be filtered out, and the search results will only return the legal domain name data)', 'zh_CN':'域名：
  // 
  // 可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpExistFlowServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportServerIpExistFlowServiceRequest) SetDomain(v []*string) *ReportServerIpExistFlowServiceRequest {
  s.Domain = v
  return s
}

type ReportServerIpExistFlowServiceResponse struct {
  // {'en':'Result', 'zh_CN':'结果'}
  Result []*ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpExistFlowServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportServerIpExistFlowServiceResponse) SetResult(v []*ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult) *ReportServerIpExistFlowServiceResponse {
  s.Result = v
  return s
}

type ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Service IP List of domains that have traffic', 'zh_CN':'域名对应的有流量的服务IP列表'}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult) SetDomain(v string) *ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult) SetIpList(v []*string) *ReportServerIpExistFlowServiceReportServerIpExistFlowServiceResponseResult {
  s.IpList = v
  return s
}

type ReportServerIpExistFlowServicePaths struct {
}

func (s ReportServerIpExistFlowServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServicePaths) GoString() string {
  return s.String()
}

type ReportServerIpExistFlowServiceParameters struct {
}

func (s ReportServerIpExistFlowServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServiceParameters) GoString() string {
  return s.String()
}

type ReportServerIpExistFlowServiceRequestHeader struct {
}

func (s ReportServerIpExistFlowServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportServerIpExistFlowServiceResponseHeader struct {
}

func (s ReportServerIpExistFlowServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpExistFlowServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportServerListServiceRequest struct {
}

func (s ReportServerListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportServerListServiceRequest) GoString() string {
  return s.String()
}

type ReportServerListServiceResponse struct {
  // {"en":"IP", "zh_CN":"节点 IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"Area", "zh_CN":"区域"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
}

func (s ReportServerListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportServerListServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportServerListServiceResponse) SetIp(v string) *ReportServerListServiceResponse {
  s.Ip = &v
  return s
}

func (s *ReportServerListServiceResponse) SetIsp(v string) *ReportServerListServiceResponse {
  s.Isp = &v
  return s
}

func (s *ReportServerListServiceResponse) SetArea(v string) *ReportServerListServiceResponse {
  s.Area = &v
  return s
}

type ReportServerListServicePaths struct {
}

func (s ReportServerListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportServerListServicePaths) GoString() string {
  return s.String()
}

type ReportServerListServiceParameters struct {
  // {"en":"Domain:
  // 1.The maximum number of deliverable domain names is 20 by default;
  // 2.Note that the domain_name of the input parameter is separated by a semicolon, and the semicolon needs to be escaped %3b;
  // 3.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names);
  // ", "zh_CN":"域名：
  // 
  // 可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 注意入参的domain_name分号隔开，分号需要转义%3b；
  // 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。"}
  Domain_name *string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
}

func (s ReportServerListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportServerListServiceParameters) GoString() string {
  return s.String()
}

func (s *ReportServerListServiceParameters) SetDomain_name(v string) *ReportServerListServiceParameters {
  s.Domain_name = &v
  return s
}

type ReportServerListServiceRequestHeader struct {
}

func (s ReportServerListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerListServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportServerListServiceResponseHeader struct {
}

func (s ReportServerListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerListServiceResponseHeader) GoString() string {
  return s.String()
}




