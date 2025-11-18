package ipcheck

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryHwCdnInInfoRequest struct {
}

func (s QueryHwCdnInInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryHwCdnInInfoRequest) GoString() string {
  return s.String()
}

type QueryHwCdnInInfoRequestHeader struct {
}

func (s QueryHwCdnInInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHwCdnInInfoRequestHeader) GoString() string {
  return s.String()
}

type QueryHwCdnInInfoPaths struct {
}

func (s QueryHwCdnInInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryHwCdnInInfoPaths) GoString() string {
  return s.String()
}

type QueryHwCdnInInfoParameters struct {
  // {"en":"IP list, split with , for multiple IP.","zh_CN":"IP列表，多个以,分隔"}
  Ips *string `json:"ips,omitempty" xml:"ips,omitempty"`
}

func (s QueryHwCdnInInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryHwCdnInInfoParameters) GoString() string {
  return s.String()
}

func (s *QueryHwCdnInInfoParameters) SetIps(v string) *QueryHwCdnInInfoParameters {
  s.Ips = &v
  return s
}

type QueryHwCdnInInfoResponse struct {
  // {"en":"IP addresses","zh_CN":"IP地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Whether to network the our IP, 1.true is the node IP of our CDN, 2.false is not the node IP of the CDN","zh_CN":"是否是华为云CDN节点IP：1.true 是，2.false 不是"}
  Belongs *string `json:"belongs,omitempty" xml:"belongs,omitempty" require:"true"`
  // {"en":"IP address location, province. If it is not a node of the CDN, it will not return.","zh_CN":"ip归属地，省份，使用全小写拼音；不是CDN的节点，不返回"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.","zh_CN":"归属城市；不是我司CDN的节点，不返回；如未规划的则返回未知；"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"If it is not a node of the CDN, it will not return.","zh_CN":"归属运营商；不是CDN的节点，不返回。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
}

func (s QueryHwCdnInInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryHwCdnInInfoResponse) GoString() string {
  return s.String()
}

func (s *QueryHwCdnInInfoResponse) SetIp(v string) *QueryHwCdnInInfoResponse {
  s.Ip = &v
  return s
}

func (s *QueryHwCdnInInfoResponse) SetBelongs(v string) *QueryHwCdnInInfoResponse {
  s.Belongs = &v
  return s
}

func (s *QueryHwCdnInInfoResponse) SetRegion(v string) *QueryHwCdnInInfoResponse {
  s.Region = &v
  return s
}

func (s *QueryHwCdnInInfoResponse) SetStatus(v string) *QueryHwCdnInInfoResponse {
  s.Status = &v
  return s
}

func (s *QueryHwCdnInInfoResponse) SetIsp(v string) *QueryHwCdnInInfoResponse {
  s.Isp = &v
  return s
}

type QueryHwCdnInInfoResponseHeader struct {
}

func (s QueryHwCdnInInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHwCdnInInfoResponseHeader) GoString() string {
  return s.String()
}




type QueryCDNServiceRealIPRequest struct {
}

func (s QueryCDNServiceRealIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPRequest) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPRequestHeader struct {
}

func (s QueryCDNServiceRealIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPRequestHeader) GoString() string {
  return s.String()
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
  // {"en":"Domain names. Which are separated by semicolons, and it supports 20 domains at max.","zh_CN":"域名，以英文分号分隔，最多20个域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Node type. Default value is all.\nOptional values:\ndyfu: dynamic relay;\nstfu: static relay;\nfu: dynamic relays and static relays;\nedge: edge node;\nall: dynamic relays, static relaysand edge nodes.","zh_CN":"节点类型，不传默认all。\ndyfu：动态父； stfu:静态父； fu：动态父+静态父；&nbsp; edge ：边缘机器；&nbsp; all：动静+边缘机器。"}
  Viewtype *string `json:"viewtype,omitempty" xml:"viewtype,omitempty"`
  // {"en":"IP form. Default value is ipseg.\nOptional values:\nrealip: real IP;\nipseg: IP segment.","zh_CN":"ip形式，不传默认 ipseg。realip：真实IP ； ipseg：ip段。"}
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

type QueryCDNServiceRealIPResponse struct {
  // {"en":"result","zh_CN":"结果"}
  Result *QueryCDNServiceRealIPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s QueryCDNServiceRealIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponse) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponse) SetResult(v *QueryCDNServiceRealIPResponseResult) *QueryCDNServiceRealIPResponse {
  s.Result = v
  return s
}

type QueryCDNServiceRealIPResponseResult struct {
  // {"en":"Result status code","zh_CN":"结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Real service IP list of domains","zh_CN":"域名对应的真实服务IP列表"}
  WhiteipList *QueryCDNServiceRealIPResponseResultWhiteipList `json:"whiteipList,omitempty" xml:"whiteipList,omitempty" require:"true" type:"Struct"`
}

func (s QueryCDNServiceRealIPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponseResult) SetCode(v string) *QueryCDNServiceRealIPResponseResult {
  s.Code = &v
  return s
}

func (s *QueryCDNServiceRealIPResponseResult) SetWhiteipList(v *QueryCDNServiceRealIPResponseResultWhiteipList) *QueryCDNServiceRealIPResponseResult {
  s.WhiteipList = v
  return s
}

type QueryCDNServiceRealIPResponseResultWhiteipList struct {
  // {"en":"Ip List","zh_CN":"真实服务IP列表"}
  Whiteiplist []*string `json:"whiteiplist,omitempty" xml:"whiteiplist,omitempty" require:"true" type:"Repeated"`
  // {"en":"Domain List","zh_CN":"域名列表"}
  DomainName []*string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCDNServiceRealIPResponseResultWhiteipList) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponseResultWhiteipList) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponseResultWhiteipList) SetWhiteiplist(v []*string) *QueryCDNServiceRealIPResponseResultWhiteipList {
  s.Whiteiplist = v
  return s
}

func (s *QueryCDNServiceRealIPResponseResultWhiteipList) SetDomainName(v []*string) *QueryCDNServiceRealIPResponseResultWhiteipList {
  s.DomainName = v
  return s
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




type QueryStageFlowIpRequest struct {
}

func (s QueryStageFlowIpRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpRequest) GoString() string {
  return s.String()
}

type QueryStageFlowIpRequestHeader struct {
}

func (s QueryStageFlowIpRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpRequestHeader) GoString() string {
  return s.String()
}

type QueryStageFlowIpPaths struct {
  // {"en":"appGroupName","zh_CN":"应用服务组名称"}
  AppGroupName *string `json:"appGroupName,omitempty" xml:"appGroupName,omitempty" require:"true"`
}

func (s QueryStageFlowIpPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpPaths) GoString() string {
  return s.String()
}

func (s *QueryStageFlowIpPaths) SetAppGroupName(v string) *QueryStageFlowIpPaths {
  s.AppGroupName = &v
  return s
}

type QueryStageFlowIpParameters struct {
}

func (s QueryStageFlowIpParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpParameters) GoString() string {
  return s.String()
}

type QueryStageFlowIpResponse struct {
  // {"en":"ip list","zh_CN":"staging ip列表"}
  Result []*QueryStageFlowIpResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStageFlowIpResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpResponse) GoString() string {
  return s.String()
}

func (s *QueryStageFlowIpResponse) SetResult(v []*QueryStageFlowIpResponseResult) *QueryStageFlowIpResponse {
  s.Result = v
  return s
}

type QueryStageFlowIpResponseResult struct     {
  // {"en":"flow ip","zh_CN":"上流量ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"code","zh_CN":"区域"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"location","zh_CN":"国家"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"ipVersion","zh_CN":"ip协议"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" require:"true"`
}

func (s QueryStageFlowIpResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStageFlowIpResponseResult) SetIp(v string) *QueryStageFlowIpResponseResult {
  s.Ip = &v
  return s
}

func (s *QueryStageFlowIpResponseResult) SetCode(v string) *QueryStageFlowIpResponseResult {
  s.Code = &v
  return s
}

func (s *QueryStageFlowIpResponseResult) SetLocation(v string) *QueryStageFlowIpResponseResult {
  s.Location = &v
  return s
}

func (s *QueryStageFlowIpResponseResult) SetIpVersion(v string) *QueryStageFlowIpResponseResult {
  s.IpVersion = &v
  return s
}

type QueryStageFlowIpResponseHeader struct {
}

func (s QueryStageFlowIpResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStageFlowIpResponseHeader) GoString() string {
  return s.String()
}




type QuerySpecificIPBelongRequest struct {
  // {"en":"IP address, use English comma to separate two items. Every IP address needs to following regular expression rule of   ((2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)\\.){3}(2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?).   The default number of IPs cannot exceed 20 (you can contact technical support to adjust) .","zh_CN":"ip地址，每个ip都需要符合正则((2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)\\.){3}(2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)，ip个数默认不能超过20（可联系技术支持调整）"}
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

type QuerySpecificIPBelongRequestHeader struct {
}

func (s QuerySpecificIPBelongRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongRequestHeader) GoString() string {
  return s.String()
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

type QuerySpecificIPBelongResponse struct {
  // {"en":"checkList","zh_CN":"结果数据"}
  CheckList []*QuerySpecificIPBelongResponseCheckList `json:"checkList,omitempty" xml:"checkList,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificIPBelongResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponse) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongResponse) SetCheckList(v []*QuerySpecificIPBelongResponseCheckList) *QuerySpecificIPBelongResponse {
  s.CheckList = v
  return s
}

type QuerySpecificIPBelongResponseCheckList struct     {
  // {"en":"yes: the IP belongs to Our system,
  // no: the IP does not belong to Our system","zh_CN":"yes：ip属于我司，no：ip不属于我司"}
  Response *string `json:"response,omitempty" xml:"response,omitempty" require:"true"`
  // {"en":"IP addresses","zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s QuerySpecificIPBelongResponseCheckList) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponseCheckList) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongResponseCheckList) SetResponse(v string) *QuerySpecificIPBelongResponseCheckList {
  s.Response = &v
  return s
}

func (s *QuerySpecificIPBelongResponseCheckList) SetIp(v string) *QuerySpecificIPBelongResponseCheckList {
  s.Ip = &v
  return s
}

type QuerySpecificIPBelongResponseHeader struct {
}

func (s QuerySpecificIPBelongResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponseHeader) GoString() string {
  return s.String()
}




type CheckIpOwnerRequest struct {
  // {"en":"The list of IP that needs to be querying is 20 times a single time.","zh_CN":"需要查询的IP列表，单次最大20个（联系技术支持可调上限）"}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s CheckIpOwnerRequest) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerRequest) GoString() string {
  return s.String()
}

func (s *CheckIpOwnerRequest) SetIp(v []*string) *CheckIpOwnerRequest {
  s.Ip = v
  return s
}

type CheckIpOwnerRequestHeader struct {
}

func (s CheckIpOwnerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerRequestHeader) GoString() string {
  return s.String()
}

type CheckIpOwnerPaths struct {
}

func (s CheckIpOwnerPaths) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerPaths) GoString() string {
  return s.String()
}

type CheckIpOwnerParameters struct {
}

func (s CheckIpOwnerParameters) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerParameters) GoString() string {
  return s.String()
}

type CheckIpOwnerResponse struct {
  // {"en":"result","zh_CN":"结果"}
  Result []*CheckIpOwnerResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s CheckIpOwnerResponse) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerResponse) GoString() string {
  return s.String()
}

func (s *CheckIpOwnerResponse) SetResult(v []*CheckIpOwnerResponseResult) *CheckIpOwnerResponse {
  s.Result = v
  return s
}

type CheckIpOwnerResponseResult struct     {
  // {"en":"IP addresses","zh_CN":"IP地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Whether the IP belongs to our CDN network.\n- true: Our CDN node IP.\n- false: Not our CDN node IP.","zh_CN":"是否我司CDN的IP：\n1.true：表示是我司CDN的节点IP\n2.false：表示不是我司CDN的节点IP","exampleValue":"true.false"}
  IsCdnIp *bool `json:"isCdnIp,omitempty" xml:"isCdnIp,omitempty" require:"true"`
  // {"en":"If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.","zh_CN":"归属国家地区；不是我司CDN的节点，不返回；如未规划的则返回未知。"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.","zh_CN":"归属省份；不是我司CDN的节点，不返回；如未规划的则返回未知；"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.","zh_CN":"归属城市；不是我司CDN的节点，不返回；如未规划的则返回未知；"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.","zh_CN":"归属运营商；不是我司CDN的节点，不返回；如未规划的则返回未知。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
}

func (s CheckIpOwnerResponseResult) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerResponseResult) GoString() string {
  return s.String()
}

func (s *CheckIpOwnerResponseResult) SetIp(v string) *CheckIpOwnerResponseResult {
  s.Ip = &v
  return s
}

func (s *CheckIpOwnerResponseResult) SetIsCdnIp(v bool) *CheckIpOwnerResponseResult {
  s.IsCdnIp = &v
  return s
}

func (s *CheckIpOwnerResponseResult) SetCountry(v string) *CheckIpOwnerResponseResult {
  s.Country = &v
  return s
}

func (s *CheckIpOwnerResponseResult) SetProvince(v string) *CheckIpOwnerResponseResult {
  s.Province = &v
  return s
}

func (s *CheckIpOwnerResponseResult) SetCity(v string) *CheckIpOwnerResponseResult {
  s.City = &v
  return s
}

func (s *CheckIpOwnerResponseResult) SetIsp(v string) *CheckIpOwnerResponseResult {
  s.Isp = &v
  return s
}

type CheckIpOwnerResponseHeader struct {
}

func (s CheckIpOwnerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIpOwnerResponseHeader) GoString() string {
  return s.String()
}




