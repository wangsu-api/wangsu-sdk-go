package ipcheck

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCDNServiceRealIPRequest struct {
}

func (s QueryCDNServiceRealIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPRequest) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s QueryCDNServiceRealIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponse) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponse) SetResult(v *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult) *QueryCDNServiceRealIPResponse {
  s.Result = v
  return s
}

type QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Real service IP list of domains', 'zh_CN':'域名对应的真实服务IP列表'}
  WhiteipList *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList `json:"whiteipList,omitempty" xml:"whiteipList,omitempty" require:"true" type:"Struct"`
}

func (s QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult) SetCode(v string) *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult {
  s.Code = &v
  return s
}

func (s *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult) SetWhiteipList(v *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList) *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResult {
  s.WhiteipList = v
  return s
}

type QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList struct {
  // {'en':'Ip List', 'zh_CN':'真实服务IP列表'}
  Whiteiplist []*string `json:"whiteiplist,omitempty" xml:"whiteiplist,omitempty" require:"true" type:"Repeated"`
  // {'en':'Domain List', 'zh_CN':'域名列表'}
  DomainName []*string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList) SetWhiteiplist(v []*string) *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList {
  s.Whiteiplist = v
  return s
}

func (s *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList) SetDomainName(v []*string) *QueryCDNServiceRealIPQueryCDNServiceRealIPResponseResultWhiteipList {
  s.DomainName = v
  return s
}

type QueryCDNServiceRealIPPaths struct {
}

func (s QueryCDNServiceRealIPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPPaths) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPParameters struct {
  // {'en':'Domain names. Which are separated by semicolons, and it supports 20 domains at max.', 'zh_CN':'域名，以英文分号分隔，最多20个域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Node type. Default value is all. 
  //     Optional values:
  //     dyfu: dynamic relay;
  //     stfu: static relay; 
  //     fu: dynamic relays and static relays;
  //     edge: edge node;
  //     all: dynamic relays, static relaysand edge nodes.', 'zh_CN':'节点类型，不传默认all。
  // 	dyfu：动态父； stfu:静态父； fu：动态父+静态父；&nbsp; edge ：边缘机器；&nbsp; all：动静+边缘机器。'}
  Viewtype *string `json:"viewtype,omitempty" xml:"viewtype,omitempty"`
  // {'en':'IP form. Default value is ipseg. 
  //     Optional values:
  //     realip: real IP; 
  //     ipseg: IP segment.', 'zh_CN':'ip形式，不传默认 ipseg。	realip：真实IP ； ipseg：ip段。'}
  Iptype *string `json:"iptype,omitempty" xml:"iptype,omitempty"`
}

func (s QueryCDNServiceRealIPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPParameters) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPParameters) SetDomain(v string) *QueryCDNServiceRealIPParameters {
  s.Domain = &v
  return s
}

func (s *QueryCDNServiceRealIPParameters) SetViewtype(v string) *QueryCDNServiceRealIPParameters {
  s.Viewtype = &v
  return s
}

func (s *QueryCDNServiceRealIPParameters) SetIptype(v string) *QueryCDNServiceRealIPParameters {
  s.Iptype = &v
  return s
}

type QueryCDNServiceRealIPRequestHeader struct {
}

func (s QueryCDNServiceRealIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPRequestHeader) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPResponseHeader struct {
}

func (s QueryCDNServiceRealIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponseHeader) GoString() string {
  return s.String()
}




type CheckIsCusWhiteIpRequest struct {
  // {"en":"entername", "zh_CN":"客户英文名"}
  White_name *string `json:"white_name,omitempty" xml:"white_name,omitempty" require:"true"`
  // {"en":"resource pool type: comm_white relay_white", "zh_CN":"资源池类型【可选】 普通：comm_white；仅中转：relay_white"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"ip white hash verify value format hash1:hash2", "zh_CN":"资源池 hash值 hash1:hash2， 如果ipv6独立 则：hash1:hash2:y"}
  White_hash *string `json:"white_hash,omitempty" xml:"white_hash,omitempty" require:"true"`
  // {"en":"check ips,split by ,", "zh_CN":"检查的ip，多个,隔开"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s CheckIsCusWhiteIpRequest) String() string {
  return tea.Prettify(s)
}

func (s CheckIsCusWhiteIpRequest) GoString() string {
  return s.String()
}

func (s *CheckIsCusWhiteIpRequest) SetWhite_name(v string) *CheckIsCusWhiteIpRequest {
  s.White_name = &v
  return s
}

func (s *CheckIsCusWhiteIpRequest) SetType(v string) *CheckIsCusWhiteIpRequest {
  s.Type = &v
  return s
}

func (s *CheckIsCusWhiteIpRequest) SetWhite_hash(v string) *CheckIsCusWhiteIpRequest {
  s.White_hash = &v
  return s
}

func (s *CheckIsCusWhiteIpRequest) SetIp(v string) *CheckIsCusWhiteIpRequest {
  s.Ip = &v
  return s
}

type CheckIsCusWhiteIpResponse struct {
  // {"en":"Return status code success represents normal fail represents abnormality", "zh_CN":"返回状态码 success 代表正常 fail 代表异常"}
  Ret_code *string `json:"ret_code,omitempty" xml:"ret_code,omitempty" require:"true"`
  // {"en":"return is every ip in use with yes or no", "zh_CN":"返回每个ip是否是在用服务资源池IP，格式为yes或者no"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"return fail message", "zh_CN":"返回报错提示消息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s CheckIsCusWhiteIpResponse) String() string {
  return tea.Prettify(s)
}

func (s CheckIsCusWhiteIpResponse) GoString() string {
  return s.String()
}

func (s *CheckIsCusWhiteIpResponse) SetRet_code(v string) *CheckIsCusWhiteIpResponse {
  s.Ret_code = &v
  return s
}

func (s *CheckIsCusWhiteIpResponse) SetData(v []*string) *CheckIsCusWhiteIpResponse {
  s.Data = v
  return s
}

func (s *CheckIsCusWhiteIpResponse) SetMsg(v string) *CheckIsCusWhiteIpResponse {
  s.Msg = &v
  return s
}

type CheckIsCusWhiteIpPaths struct {
}

func (s CheckIsCusWhiteIpPaths) String() string {
  return tea.Prettify(s)
}

func (s CheckIsCusWhiteIpPaths) GoString() string {
  return s.String()
}

type CheckIsCusWhiteIpParameters struct {
}

func (s CheckIsCusWhiteIpParameters) String() string {
  return tea.Prettify(s)
}

func (s CheckIsCusWhiteIpParameters) GoString() string {
  return s.String()
}

type CheckIsCusWhiteIpRequestHeader struct {
}

func (s CheckIsCusWhiteIpRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIsCusWhiteIpRequestHeader) GoString() string {
  return s.String()
}

type CheckIsCusWhiteIpResponseHeader struct {
}

func (s CheckIsCusWhiteIpResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIsCusWhiteIpResponseHeader) GoString() string {
  return s.String()
}




type QuerySpecificIPBelongRequest struct {
  // {'en':'IP address, use English comma to separate two items. Every IP address needs to following regular expression rule of   ((2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)\\.){3}(2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?).   The default number of IPs cannot exceed 20 (you can contact technical support to adjust) .', 'zh_CN':'ip地址，以英文逗号分隔，每个ip都需要符合正则((2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)\\.){3}(2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)，ip个数默认不能超过20（可联系技术支持调整）'}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificIPBelongRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongRequest) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongRequest) SetIp(v []*string) *QuerySpecificIPBelongRequest {
  s.Ip = v
  return s
}

type QuerySpecificIPBelongResponse struct {
  // {'en':'checkList', 'zh_CN':'结果数据'}
  CheckList []*QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList `json:"checkList,omitempty" xml:"checkList,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificIPBelongResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponse) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongResponse) SetCheckList(v []*QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList) *QuerySpecificIPBelongResponse {
  s.CheckList = v
  return s
}

type QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList struct     {
  // {'en':'yes: the IP belongs to Our system,
  //         no: the IP does not belong to Our system', 'zh_CN':'yes：ip属于我司，no：ip不属于我司'}
  Response *string `json:"response,omitempty" xml:"response,omitempty" require:"true"`
  // {'en':'IP addresses', 'zh_CN':'ip地址'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList) SetResponse(v string) *QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList {
  s.Response = &v
  return s
}

func (s *QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList) SetIp(v string) *QuerySpecificIPBelongQuerySpecificIPBelongResponseCheckList {
  s.Ip = &v
  return s
}

type QuerySpecificIPBelongPaths struct {
}

func (s QuerySpecificIPBelongPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongPaths) GoString() string {
  return s.String()
}

type QuerySpecificIPBelongParameters struct {
}

func (s QuerySpecificIPBelongParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongParameters) GoString() string {
  return s.String()
}

type QuerySpecificIPBelongRequestHeader struct {
}

func (s QuerySpecificIPBelongRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongRequestHeader) GoString() string {
  return s.String()
}

type QuerySpecificIPBelongResponseHeader struct {
}

func (s QuerySpecificIPBelongResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponseHeader) GoString() string {
  return s.String()
}




type IpInfoServiceRequest struct {
  // {'en':'The list of IP that needs to be querying is 20 times a single time.', 'zh_CN':'需要查询的IP列表，单次最大20个（联系技术支持可调上限）'}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s IpInfoServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceRequest) GoString() string {
  return s.String()
}

func (s *IpInfoServiceRequest) SetIp(v []*string) *IpInfoServiceRequest {
  s.Ip = v
  return s
}

type IpInfoServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*IpInfoServiceIpInfoServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s IpInfoServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponse) GoString() string {
  return s.String()
}

func (s *IpInfoServiceResponse) SetResult(v []*IpInfoServiceIpInfoServiceResponseResult) *IpInfoServiceResponse {
  s.Result = v
  return s
}

type IpInfoServiceIpInfoServiceResponseResult struct     {
  // {'en':'IP addresses', 'zh_CN':'IP地址'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {'en':'Whether to network the our IP
  // 
  //         1.true is the node IP of our CDN
  // 
  //         2.false is not the node IP of the CDN', 'zh_CN':'是否我司CDN的IP
  //         1.true 是我司CDN的节点IP
  //         2.false &nbsp; 不是我司CDN的节点IP'}
  IsCdnIp *bool `json:"isCdnIp,omitempty" xml:"isCdnIp,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属国家地区；不是我司CDN的节点，不返回；如未规划的则返回未知。'}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属省份；不是我司CDN的节点，不返回；如未规划的则返回未知；'}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属城市；不是我司CDN的节点，不返回；如未规划的则返回未知；'}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属运营商；不是我司CDN的节点，不返回；如未规划的则返回未知。'}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
}

func (s IpInfoServiceIpInfoServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceIpInfoServiceResponseResult) GoString() string {
  return s.String()
}

func (s *IpInfoServiceIpInfoServiceResponseResult) SetIp(v string) *IpInfoServiceIpInfoServiceResponseResult {
  s.Ip = &v
  return s
}

func (s *IpInfoServiceIpInfoServiceResponseResult) SetIsCdnIp(v bool) *IpInfoServiceIpInfoServiceResponseResult {
  s.IsCdnIp = &v
  return s
}

func (s *IpInfoServiceIpInfoServiceResponseResult) SetCountry(v string) *IpInfoServiceIpInfoServiceResponseResult {
  s.Country = &v
  return s
}

func (s *IpInfoServiceIpInfoServiceResponseResult) SetProvince(v string) *IpInfoServiceIpInfoServiceResponseResult {
  s.Province = &v
  return s
}

func (s *IpInfoServiceIpInfoServiceResponseResult) SetCity(v string) *IpInfoServiceIpInfoServiceResponseResult {
  s.City = &v
  return s
}

func (s *IpInfoServiceIpInfoServiceResponseResult) SetIsp(v string) *IpInfoServiceIpInfoServiceResponseResult {
  s.Isp = &v
  return s
}

type IpInfoServicePaths struct {
}

func (s IpInfoServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServicePaths) GoString() string {
  return s.String()
}

type IpInfoServiceParameters struct {
}

func (s IpInfoServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceParameters) GoString() string {
  return s.String()
}

type IpInfoServiceRequestHeader struct {
}

func (s IpInfoServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceRequestHeader) GoString() string {
  return s.String()
}

type IpInfoServiceResponseHeader struct {
}

func (s IpInfoServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponseHeader) GoString() string {
  return s.String()
}




