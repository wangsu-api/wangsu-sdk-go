package gtmmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ControlResourceClusterRequest struct {
  Param []*ControlResourceClusterControlResourceClusterRequestParam `json:"param,omitempty" xml:"param,omitempty" require:"true" type:"Repeated"`
}

func (s ControlResourceClusterRequest) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterRequest) GoString() string {
  return s.String()
}

func (s *ControlResourceClusterRequest) SetParam(v []*ControlResourceClusterControlResourceClusterRequestParam) *ControlResourceClusterRequest {
  s.Param = v
  return s
}

type ControlResourceClusterControlResourceClusterRequestParam struct     {
  // {"en":"The policy id", "zh_CN":"策略id"}
  PolicyId *int `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"The start stop code 1111 consists of four digits of 0 or 1, representing the primary source, primary backup source, secondary backup source, and tertiary backup source. 0 is disabled and 1 is enabled. For example, if only the primary backup source needs to be disabled, code=1011, both the primary source and the tertiary backup source need to be disabled. code=0110. If the backup source does not exist, it is defaulted to 1", "zh_CN":"启停代码 1111 四位0或1的数字代表主源 一级备源 二级备源 三级备源  0 停用 1 启用  例如 需要只停用一级备 code=1011  需要同时停用主源和三级备 code=0110  如果备源不存在则默认补1"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s ControlResourceClusterControlResourceClusterRequestParam) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterControlResourceClusterRequestParam) GoString() string {
  return s.String()
}

func (s *ControlResourceClusterControlResourceClusterRequestParam) SetPolicyId(v int) *ControlResourceClusterControlResourceClusterRequestParam {
  s.PolicyId = &v
  return s
}

func (s *ControlResourceClusterControlResourceClusterRequestParam) SetCode(v string) *ControlResourceClusterControlResourceClusterRequestParam {
  s.Code = &v
  return s
}

type ControlResourceClusterResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Response message", "zh_CN":"返回说明"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回值"}
  Content []*ControlResourceClusterControlResourceClusterResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s ControlResourceClusterResponse) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterResponse) GoString() string {
  return s.String()
}

func (s *ControlResourceClusterResponse) SetResCode(v int) *ControlResourceClusterResponse {
  s.ResCode = &v
  return s
}

func (s *ControlResourceClusterResponse) SetMsg(v string) *ControlResourceClusterResponse {
  s.Msg = &v
  return s
}

func (s *ControlResourceClusterResponse) SetContent(v []*ControlResourceClusterControlResourceClusterResponseContent) *ControlResourceClusterResponse {
  s.Content = v
  return s
}

type ControlResourceClusterControlResourceClusterResponseContent struct     {
}

func (s ControlResourceClusterControlResourceClusterResponseContent) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterControlResourceClusterResponseContent) GoString() string {
  return s.String()
}

type ControlResourceClusterPaths struct {
}

func (s ControlResourceClusterPaths) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterPaths) GoString() string {
  return s.String()
}

type ControlResourceClusterParameters struct {
}

func (s ControlResourceClusterParameters) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterParameters) GoString() string {
  return s.String()
}

type ControlResourceClusterRequestHeader struct {
}

func (s ControlResourceClusterRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterRequestHeader) GoString() string {
  return s.String()
}

type ControlResourceClusterResponseHeader struct {
}

func (s ControlResourceClusterResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlResourceClusterResponseHeader) GoString() string {
  return s.String()
}




type ControlDispatchPolicyRequest struct {
  // {"en":"Domain that the dispatch policy to be operated on belongs to", "zh_CN":"要操作调度策略所属的域名"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"ID of the dispatch policy to be operated on Use English half-width semicolon between two policies if there are multiple to be operated on.", "zh_CN":"要删除的调度策略ID
  // 如果需要删除多个策略，用英文半角分号分隔。"}
  PolicyIds *string `json:"policyIds,omitempty" xml:"policyIds,omitempty" require:"true"`
  // {"en":"Operation type 0 means to enable policies; 1 means to disable policies", "zh_CN":"操作类型：0表示启用策略；1表示停用策略"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Operation type 0 means to enable policies; 1 means to disable policies", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ControlDispatchPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchPolicyRequest) GoString() string {
  return s.String()
}

func (s *ControlDispatchPolicyRequest) SetDomainId(v int) *ControlDispatchPolicyRequest {
  s.DomainId = &v
  return s
}

func (s *ControlDispatchPolicyRequest) SetPolicyIds(v string) *ControlDispatchPolicyRequest {
  s.PolicyIds = &v
  return s
}

func (s *ControlDispatchPolicyRequest) SetType(v int) *ControlDispatchPolicyRequest {
  s.Type = &v
  return s
}

func (s *ControlDispatchPolicyRequest) SetLanguage(v string) *ControlDispatchPolicyRequest {
  s.Language = &v
  return s
}

type ControlDispatchPolicyResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"policyId:The dispatch policy ID, used to identify the newly added dispatch policy", "zh_CN":"policyId调度策略ID，用于标识调度策略"}
  Content map[string]interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s ControlDispatchPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchPolicyResponse) GoString() string {
  return s.String()
}

func (s *ControlDispatchPolicyResponse) SetResCode(v int) *ControlDispatchPolicyResponse {
  s.ResCode = &v
  return s
}

func (s *ControlDispatchPolicyResponse) SetMsg(v string) *ControlDispatchPolicyResponse {
  s.Msg = &v
  return s
}

func (s *ControlDispatchPolicyResponse) SetContent(v map[string]interface{}) *ControlDispatchPolicyResponse {
  s.Content = v
  return s
}

type ControlDispatchPolicyPaths struct {
}

func (s ControlDispatchPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchPolicyPaths) GoString() string {
  return s.String()
}

type ControlDispatchPolicyParameters struct {
}

func (s ControlDispatchPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchPolicyParameters) GoString() string {
  return s.String()
}

type ControlDispatchPolicyRequestHeader struct {
}

func (s ControlDispatchPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchPolicyRequestHeader) GoString() string {
  return s.String()
}

type ControlDispatchPolicyResponseHeader struct {
}

func (s ControlDispatchPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchPolicyResponseHeader) GoString() string {
  return s.String()
}




type QueryDispatchPoliciesRequest struct {
  // {"en":"Domain that the dispatch policy to be edited belongs to", "zh_CN":"要查询调度策略所属的域名"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Query inital records by page", "zh_CN":"分页查询起始记录"}
  Start *int `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {"en":"Number if quried items by page", "zh_CN":"分页查询条数"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"Number if quried items by page", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s QueryDispatchPoliciesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesRequest) GoString() string {
  return s.String()
}

func (s *QueryDispatchPoliciesRequest) SetDomainId(v int) *QueryDispatchPoliciesRequest {
  s.DomainId = &v
  return s
}

func (s *QueryDispatchPoliciesRequest) SetStart(v int) *QueryDispatchPoliciesRequest {
  s.Start = &v
  return s
}

func (s *QueryDispatchPoliciesRequest) SetLimit(v int) *QueryDispatchPoliciesRequest {
  s.Limit = &v
  return s
}

func (s *QueryDispatchPoliciesRequest) SetLanguage(v string) *QueryDispatchPoliciesRequest {
  s.Language = &v
  return s
}

type QueryDispatchPoliciesResponse struct {
  // {"en":"Policy description", "zh_CN":"策略描述"}
  PolicyDesc *string `json:"policyDesc,omitempty" xml:"policyDesc,omitempty" require:"true"`
  // {"en":"Policy type, 0: Load balance+primary and redundant, 1: Load balance", "zh_CN":"策略类型，0:负载均衡+主备,1:负载均衡"}
  PolicyType *int `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {"en":"Domain ID", "zh_CN":"域名ID标识"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Line information type Line type, 0: Standard line, 1: Custom Line viewId  Line ID, field required when the line type is 0 viewCn  Line in Chinese, field required when the line type is 0 userView Custom line, field required when the line type is 1 viewTag Custom line tag viewMembers Actual lines contained in the custom line viewId Line ID viewCn Line in Chinese", "zh_CN":"线路信息：
  // type 线路类型， 0:标准线路,1:自定义线路
  // 
  // viewId 线路ID  线路类型为0时选项
  // 
  // viewCn线路中文  线路类型为0时选项
  // 
  // viewEn线路英文  线路类型为0时选项
  // 
  // userView 自定义线路  线路类型为1时选项
  // 
  // viewTag自定义线路标签
  // 
  // viewMembers 自定义线路包含的实际线路
  // 
  // viewId线路ID
  // 
  // viewCn 线路中文"}
  View *QueryDispatchPoliciesQueryDispatchPoliciesResponseView `json:"view,omitempty" xml:"view,omitempty" require:"true" type:"Struct"`
  // {"en":"Call frequency Unit is minute(1,2,5,10,30,60)", "zh_CN":"调度频率
  // 以分钟为单位， 可选值1、2、5、10、30、60"}
  Rate *int `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {"en":"Monitor type monitorType Monitor type  0 http, 1 https, 2 udp(not supported presently), 3 tcp, 4 ping monitorNodes Monitor nodes isp ISP of monitor nodes area Area of monitor nodes path Monitor path, options are available when the monitor types are http and https.  port Monitor port, options are available when the monitor types are http, https and tcp.  responseTimeout Reponse timeout time, unit is second. Options are available when the monitor types are http and https.  excludedCodes Excluded status codes. Options are available when the monitor types are http and https. Use semicolon to separate when there are multiple status codes packetLossLimit Packet loss ratio, available when the monitor type is ping.  delayLimit Time delay, unit is millisecond. Available when the monitor type is ping.", "zh_CN":"监控配置：
  // monitorType 监控类型，0 http 1 https 2 udp(暂不支持) 3 tcp 4 ping
  // 
  // monitorNodes 监控节点
  // 
  // isp 监控节点运营商
  // 
  // area 监控节点区域
  // 
  // path 监控路径，当监控方式为http,https 时选项
  // 
  // port监控端口，当监控方式为http,https,tcp 时选项
  // 
  // responseTimeout响应超时时间，单位：秒，当监控方式为http,https时选项
  // 
  // excludedCodes状态排除码，当监控方式为http,https 时选项，多个状态码用英文分号分隔
  // 
  // httpMethod 监控方式为http/https时支持，可选值：0 默认请求方法 1 post请求
  // 
  // requestData 监控方式为http/https时支持，httpMethod为1时 必填(可为空串)
  // 
  // packetLossLimit 丢包率，当监控方式为ping时选项
  // 
  // delayLimit时延，单位：毫秒，当监控方式为ping时选项"}
  Monitor *QueryDispatchPoliciesQueryDispatchPoliciesResponseMonitor `json:"monitor,omitempty" xml:"monitor,omitempty" require:"true" type:"Struct"`
  // {"en":"Warning configurations warnMethod Warning type,1 Warn with email warnInterval How long will the warning last, unit: minute warnEmail The Email box to receive warning messages. Use English semicolon to separate two if there are multiple email boxes exist", "zh_CN":"告警配置：
  // warnMethod告警方式， 1 邮件告警
  // 
  // warnInterval连续告警提醒周期，单位：分钟
  // 
  // warnEmail 告警邮箱，多个邮箱以英文分号分隔"}
  Warning *QueryDispatchPoliciesQueryDispatchPoliciesResponseWarning `json:"warning,omitempty" xml:"warning,omitempty" require:"true" type:"Struct"`
  // {"en":"Policy resources partType Resource type, 0 Primary DNS 1 Level-one redundancy 2 Level-two redundancy 3 Level-three redundancy type Resource record type, 0 A record 1 CNAME value Resource record value loadRatio Ratio", "zh_CN":"策略资源：
  // partType资源类型， 0 主解析资源 1 一级备 2 二级备 3 三级备
  // 
  // type 资源记录类型， 0 A记录 1 CNAME
  // 
  // value 资源记录值
  // 
  // loadRatio 比例"}
  PolicyResource []*string `json:"policyResource,omitempty" xml:"policyResource,omitempty" require:"true" type:"Repeated"`
  // {"en":"DNS details(if release node does not exist or is empty, it means the current policy has not been released) spType Type of the service provider  0 Primary DNS 1 Level-one redundancy 2 Leveltwo redundancy 3 Level-three redundancy spList Resource list of service providers load Load ratio value Resource record value.", "zh_CN":"解析情况(release节点不存在或为空，表示当前策略尚未发布解析 )
  // spType 服务提供者类型 0 主解析资源 1 一级备 2 二级备 3 三级备
  // 
  // spList 服务提供资源列表
  // 
  // load 负载比例
  // 
  // value 资源记录值"}
  Release []*string `json:"release,omitempty" xml:"release,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDispatchPoliciesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesResponse) GoString() string {
  return s.String()
}

func (s *QueryDispatchPoliciesResponse) SetPolicyDesc(v string) *QueryDispatchPoliciesResponse {
  s.PolicyDesc = &v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetPolicyType(v int) *QueryDispatchPoliciesResponse {
  s.PolicyType = &v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetDomainId(v int) *QueryDispatchPoliciesResponse {
  s.DomainId = &v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetView(v *QueryDispatchPoliciesQueryDispatchPoliciesResponseView) *QueryDispatchPoliciesResponse {
  s.View = v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetRate(v int) *QueryDispatchPoliciesResponse {
  s.Rate = &v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetMonitor(v *QueryDispatchPoliciesQueryDispatchPoliciesResponseMonitor) *QueryDispatchPoliciesResponse {
  s.Monitor = v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetWarning(v *QueryDispatchPoliciesQueryDispatchPoliciesResponseWarning) *QueryDispatchPoliciesResponse {
  s.Warning = v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetPolicyResource(v []*string) *QueryDispatchPoliciesResponse {
  s.PolicyResource = v
  return s
}

func (s *QueryDispatchPoliciesResponse) SetRelease(v []*string) *QueryDispatchPoliciesResponse {
  s.Release = v
  return s
}

type QueryDispatchPoliciesQueryDispatchPoliciesResponseView struct {
}

func (s QueryDispatchPoliciesQueryDispatchPoliciesResponseView) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesQueryDispatchPoliciesResponseView) GoString() string {
  return s.String()
}

type QueryDispatchPoliciesQueryDispatchPoliciesResponseMonitor struct {
}

func (s QueryDispatchPoliciesQueryDispatchPoliciesResponseMonitor) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesQueryDispatchPoliciesResponseMonitor) GoString() string {
  return s.String()
}

type QueryDispatchPoliciesQueryDispatchPoliciesResponseWarning struct {
}

func (s QueryDispatchPoliciesQueryDispatchPoliciesResponseWarning) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesQueryDispatchPoliciesResponseWarning) GoString() string {
  return s.String()
}

type QueryDispatchPoliciesPaths struct {
}

func (s QueryDispatchPoliciesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesPaths) GoString() string {
  return s.String()
}

type QueryDispatchPoliciesParameters struct {
}

func (s QueryDispatchPoliciesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesParameters) GoString() string {
  return s.String()
}

type QueryDispatchPoliciesRequestHeader struct {
}

func (s QueryDispatchPoliciesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesRequestHeader) GoString() string {
  return s.String()
}

type QueryDispatchPoliciesResponseHeader struct {
}

func (s QueryDispatchPoliciesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPoliciesResponseHeader) GoString() string {
  return s.String()
}




type SaveDispatchPolicyRequest struct {
  // {"en":"Policy description Less than 200 characters", "zh_CN":"策略描述
  // 小于200字符"}
  PolicyDesc *string `json:"policyDesc,omitempty" xml:"policyDesc,omitempty"`
  // {"en":"Policy type 0: Load balance+primary and redundant, 1: Load balance", "zh_CN":"策略类型
  // 0:负载均衡+主备,1:负载均衡"}
  PolicyType *int `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {"en":"Domain ID", "zh_CN":"域名ID标识"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Line information. Entry rules: type Line type, optional field. Options: 0: Standard line, 1: Custom Line. The default value is Standard line customId ID of the custom line, available and required when the type value is 1 More details about the related table of lines can be found in Appendix II Relations of ViewID and Lines viewId Line ID, available and required when the type value is 0 viewCn Line in Chinese, available and required when the type value is 0", "zh_CN":"线路信息。
  // 填写规则：
  // 
  // type 线路类型，可选，可选值0 标准线路  1 自定义线路。默认值为标准线路
  // 
  // customId 自定义线路ID , type值为1时选项，必填
  // 
  // 线路的对应表请参考附录“附录2 ViewID与线路的对应关系”
  // 
  // viewId 线路ID，type值为0时选项，必填
  // 
  // viewCn线路中文，type值为0时选项，必填"}
  View map[string]interface{} `json:"view,omitempty" xml:"view,omitempty" require:"true"`
  // {"en":"Call frequency Unit is minute(1,2,5,10,30,60)", "zh_CN":"调度频率
  // 以分钟为单位， 可选值1、2、5、10、30、60"}
  Rate *int `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {"en":"Monitor configurations. Entry rules: monitorType Monitor type, required. Optional values: 0 http, 1 https, 2 udp(not supported presently), 3 tcp, 4 ping monitorNodes Monitor nodes, required isp ISP of monitor nodes area Area of monitor nodes path Monitor path, options are available and required when the monitor types are http and https. Length should be less than 255 port Monitor port, options are available and required when the monitor types are http, https and tcp. Value range is 1-65535. responseTimeout Reponse timeout time, unit is second. Options are available and required when the monitor types are http and https. Value range is 1-20 excludedCodes Excluded status codes. Options are available and optional when the monitor types are http and https. Use semicolon to separate when there are multiple status codes packetLossLimit Packet loss ratio, available when the monitor type is ping. Optional field. Either packet loss ratio or time delay has to be entered delayLimit Time delay, unit is millisecond. Available when the monitor type is ping. Optional field. Either packet loss ratio or time delay has to be entered", "zh_CN":"监控配置。
  // 填写规则：
  // 
  // monitorType 监控类型，必填，可选值0 http 1 https 2 udp(暂不支持) 3 tcp 4 ping
  // 
  // monitorNodes 监控节点，必填
  // 
  // isp 监控节点运营商
  // 
  // area 监控节点区域
  // 
  // path 监控路径，当监控方式为http,https 时选项，必填， 长度不超过255
  // 
  // port监控端口，当监控方式为http,https,tcp 时选项，必填，取值范围 1~65535
  // 
  // responseTimeout响应超时时间，单位：秒，当监控方式为http,https时选项，必填，1 ~ 20
  // 
  // excludedCodes状态排除码，当监控方式为http,https 时选项，选填，多个状态码用英文分号分隔
  // 
  // httpMethod 监控方式为httptps时支持，可选值：0 默认请求方法 1 post请求
  // 
  // requestData 监控方式为httptps时支持，httpMethod为1时 必填(可为空串)
  // 
  // packetLossLimit 丢包率，当监控方式为ping时选项，选填，但丢包率和时延至少填一项
  // 
  // delayLimit时延，单位：毫秒，当监控方式为ping时选项，选填，但丢包率和时延至少填一项"}
  Monitor map[string]interface{} `json:"monitor,omitempty" xml:"monitor,omitempty" require:"true"`
  // {"en":"Warning configurations. Entry rules: warnMethod Warning type, required filed, 1 Warn with email warnInterval How long will the warning last, unit: minutes. This filed is complusory when the warn type is 0. Warning duration>=call frequency duration warnEmail The email box to receive warning messages, and this field is required if the warn type is 1. Use English semicolon to separate two if there are multiple email boxes exist", "zh_CN":"告警配置。
  // 填写规则：
  // 
  // warnMethod告警方式，必填，1 邮件告警
  // 
  // warnInterval连续告警提醒周期，单位：分钟，告警方式不为0时必填，告警提醒周期>=调度频率周期
  // 
  // warnEmail 告警邮箱，告警方式为1时必填，多个邮箱以英文分号分隔"}
  Warning map[string]interface{} `json:"warning,omitempty" xml:"warning,omitempty" require:"true"`
  // {"en":"Policy resources. Entry rules: partType Resource type, field required, optional values: 0 Primary DNS 1 Levelone redundancy 2 Level-two redundancy 3 Level-three redundancy type Resource record type, field required; optional values: 0 A record 1 CNAME value Resource record value, field required  loadRatio Ratio, field required Primary resource has to be entered. When lower levels of redundancies exist, higher levels cannot be empty. When the policy type is load balance+primary and redundancy, the redundant resource cannot be empty. When the policy type is load balance, the entered redundant resource is invalid. No duplicate values of policy resources are allowed", "zh_CN":"策略资源。
  // 填写规则：
  // 
  // partType资源类型，必填，可选值： 0 主解析资源 1 一级备 2 二级备 3 三级备
  // 
  // type 资源记录类型，必填，可选值：0 A记录 1 CNAME
  // 
  // value 资源记录值，必填
  // 
  // loadRatio 比例，必填
  // 
  // 必须填写主资源。
  // 
  // 当有更低级备资源时，较高级备资源不能为空。
  // 
  // 当策略类型为负载均衡+主备时，备资源不能为空。
  // 
  // 当策略类型为负载均衡时，填写的备资源无效。
  // 
  // 策略资源值不能重复"}
  PolicyResource []*string `json:"policyResource,omitempty" xml:"policyResource,omitempty" require:"true" type:"Repeated"`
  // {"en":"Policy resources. Entry rules: partType Resource type, field required, optional values: 0 Primary DNS 1 Levelone redundancy 2 Level-two redundancy 3 Level-three redundancy type Resource record type, field required; optional values: 0 A record 1 CNAME value Resource record value, field required  loadRatio Ratio, field required Primary resource has to be entered. When lower levels of redundancies exist, higher levels cannot be empty. When the policy type is load balance+primary and redundancy, the redundant resource cannot be empty. When the policy type is load balance, the entered redundant resource is invalid. No duplicate values of policy resources are allowed", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s SaveDispatchPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s SaveDispatchPolicyRequest) GoString() string {
  return s.String()
}

func (s *SaveDispatchPolicyRequest) SetPolicyDesc(v string) *SaveDispatchPolicyRequest {
  s.PolicyDesc = &v
  return s
}

func (s *SaveDispatchPolicyRequest) SetPolicyType(v int) *SaveDispatchPolicyRequest {
  s.PolicyType = &v
  return s
}

func (s *SaveDispatchPolicyRequest) SetDomainId(v int) *SaveDispatchPolicyRequest {
  s.DomainId = &v
  return s
}

func (s *SaveDispatchPolicyRequest) SetView(v map[string]interface{}) *SaveDispatchPolicyRequest {
  s.View = v
  return s
}

func (s *SaveDispatchPolicyRequest) SetRate(v int) *SaveDispatchPolicyRequest {
  s.Rate = &v
  return s
}

func (s *SaveDispatchPolicyRequest) SetMonitor(v map[string]interface{}) *SaveDispatchPolicyRequest {
  s.Monitor = v
  return s
}

func (s *SaveDispatchPolicyRequest) SetWarning(v map[string]interface{}) *SaveDispatchPolicyRequest {
  s.Warning = v
  return s
}

func (s *SaveDispatchPolicyRequest) SetPolicyResource(v []*string) *SaveDispatchPolicyRequest {
  s.PolicyResource = v
  return s
}

func (s *SaveDispatchPolicyRequest) SetLanguage(v string) *SaveDispatchPolicyRequest {
  s.Language = &v
  return s
}

type SaveDispatchPolicyResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"policyId scheduling policy ID", "zh_CN":"policyId调度策略ID，用于标识新增的调度策略"}
  Content map[string]interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s SaveDispatchPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s SaveDispatchPolicyResponse) GoString() string {
  return s.String()
}

func (s *SaveDispatchPolicyResponse) SetResCode(v int) *SaveDispatchPolicyResponse {
  s.ResCode = &v
  return s
}

func (s *SaveDispatchPolicyResponse) SetMsg(v string) *SaveDispatchPolicyResponse {
  s.Msg = &v
  return s
}

func (s *SaveDispatchPolicyResponse) SetContent(v map[string]interface{}) *SaveDispatchPolicyResponse {
  s.Content = v
  return s
}

type SaveDispatchPolicyPaths struct {
}

func (s SaveDispatchPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s SaveDispatchPolicyPaths) GoString() string {
  return s.String()
}

type SaveDispatchPolicyParameters struct {
}

func (s SaveDispatchPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s SaveDispatchPolicyParameters) GoString() string {
  return s.String()
}

type SaveDispatchPolicyRequestHeader struct {
}

func (s SaveDispatchPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SaveDispatchPolicyRequestHeader) GoString() string {
  return s.String()
}

type SaveDispatchPolicyResponseHeader struct {
}

func (s SaveDispatchPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SaveDispatchPolicyResponseHeader) GoString() string {
  return s.String()
}




type QueryDispatchPolicyDetailRequest struct {
  // {"en":"Domain ID", "zh_CN":"域名ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"PolicyID", "zh_CN":"策略ID"}
  PolicyId *int `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"PolicyID", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s QueryDispatchPolicyDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPolicyDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryDispatchPolicyDetailRequest) SetDomainId(v int) *QueryDispatchPolicyDetailRequest {
  s.DomainId = &v
  return s
}

func (s *QueryDispatchPolicyDetailRequest) SetPolicyId(v int) *QueryDispatchPolicyDetailRequest {
  s.PolicyId = &v
  return s
}

func (s *QueryDispatchPolicyDetailRequest) SetLanguage(v string) *QueryDispatchPolicyDetailRequest {
  s.Language = &v
  return s
}

type QueryDispatchPolicyDetailResponse struct {
  // {"en":"Policy description", "zh_CN":"策略描述"}
  PolicyDesc *string `json:"policyDesc,omitempty" xml:"policyDesc,omitempty" require:"true"`
  // {"en":"Policy type, 0: Load balance+primary and redundant, 1: Load balance", "zh_CN":"策略类型，0:负载均衡+主备,1:负载均衡"}
  PolicyType *int `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {"en":"Domain ID", "zh_CN":"域名ID标识"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Line information type Line type, 0: Standard line, 1: Custom Line viewId  Line ID, field required when the line type is 0 viewCn  Line in Chinese, field required when the line type is 0 userView Custom line, field required when the line type is 1 viewTag Custom line tag viewMembers Actual lines contained in the custom line viewId Line ID viewCn Line in Chinese", "zh_CN":"线路信息：
  // type 线路类型， 0:标准线路,1:自定义线路
  // 
  // viewId 线路ID  线路类型为0时选项
  // 
  // viewCn线路中文  线路类型为0时选项
  // 
  // viewEn线路英文  线路类型为0时选项
  // 
  // userView 自定义线路  线路类型为1时选项
  // 
  // viewTag自定义线路标签
  // 
  // viewMembers 自定义线路包含的实际线路
  // 
  // viewId线路ID
  // 
  // viewCn 线路中文"}
  View map[string]interface{} `json:"view,omitempty" xml:"view,omitempty" require:"true"`
  // {"en":"Call frequency Unit is minute(1,2,5,10,30,60)", "zh_CN":"调度频率
  // 以分钟为单位， 可选值1、2、5、10、30、60"}
  Rate *int `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {"en":"Monitor type monitorType Monitor type  0 http, 1 https, 2 udp(not supported presently), 3 tcp, 4 ping monitorNodes Monitor nodes isp ISP of monitor nodes area Area of monitor nodes path Monitor path, options are available when the monitor types are http and https.  port Monitor port, options are available when the monitor types are http, https and tcp.  responseTimeout Reponse timeout time, unit is second. Options are available when the monitor types are http and https.  excludedCodes Excluded status codes. Options are available when the monitor types are http and https. Use semicolon to separate when there are multiple status codes packetLossLimit Packet loss ratio, available when the monitor type is ping.  delayLimit Time delay, unit is millisecond. Available when the monitor type is ping.", "zh_CN":"监控配置：
  // monitorType 监控类型，0 http 1 https 2 udp(暂不支持) 3 tcp 4 ping
  // 
  // monitorNodes 监控节点
  // 
  // isp 监控节点运营商
  // 
  // area 监控节点区域
  // 
  // path 监控路径，当监控方式为http,https 时选项
  // 
  // port监控端口，当监控方式为http,https,tcp 时选项
  // 
  // responseTimeout响应超时时间，单位：秒，当监控方式为http,https时选项
  // 
  // excludedCodes状态排除码，当监控方式为http,https 时选项，多个状态码用英文分号分隔
  // 
  // httpMethod 监控方式为http/https时支持，可选值：0 默认请求方法 1 post请求
  // 
  // requestData 监控方式为http/https时支持，httpMethod为1时 必填(可为空串)
  // 
  // packetLossLimit 丢包率，当监控方式为ping时选项
  // 
  // delayLimit时延，单位：毫秒，当监控方式为ping时选项"}
  Monitor map[string]interface{} `json:"monitor,omitempty" xml:"monitor,omitempty" require:"true"`
  // {"en":"Warning configurations warnMethod Warning type,1 Warn with email warnInterval How long will the warning last, unit: minute warnEmail The Email box to receive warning messages. Use English semicolon to separate two if there are multiple email boxes exist", "zh_CN":"告警配置：
  // warnMethod告警方式， 1 邮件告警
  // 
  // warnInterval连续告警提醒周期，单位：分钟
  // 
  // warnEmail 告警邮箱，多个邮箱以英文分号分隔"}
  Warning map[string]interface{} `json:"warning,omitempty" xml:"warning,omitempty" require:"true"`
  // {"en":"Policy resources partType Resource type, 0 Primary DNS 1 Level-one redundancy 2 Level-two redundancy 3 Level-three redundancy type Resource record type, 0 A record 1 CNAME value Resource record value loadRatio Ratio", "zh_CN":"策略资源：
  // partType资源类型， 0 主解析资源 1 一级备 2 二级备 3 三级备
  // 
  // type 资源记录类型， 0 A记录 1 CNAME
  // 
  // value 资源记录值
  // 
  // loadRatio 比例"}
  PolicyResource []*string `json:"policyResource,omitempty" xml:"policyResource,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDispatchPolicyDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPolicyDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryDispatchPolicyDetailResponse) SetPolicyDesc(v string) *QueryDispatchPolicyDetailResponse {
  s.PolicyDesc = &v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetPolicyType(v int) *QueryDispatchPolicyDetailResponse {
  s.PolicyType = &v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetDomainId(v int) *QueryDispatchPolicyDetailResponse {
  s.DomainId = &v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetView(v map[string]interface{}) *QueryDispatchPolicyDetailResponse {
  s.View = v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetRate(v int) *QueryDispatchPolicyDetailResponse {
  s.Rate = &v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetMonitor(v map[string]interface{}) *QueryDispatchPolicyDetailResponse {
  s.Monitor = v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetWarning(v map[string]interface{}) *QueryDispatchPolicyDetailResponse {
  s.Warning = v
  return s
}

func (s *QueryDispatchPolicyDetailResponse) SetPolicyResource(v []*string) *QueryDispatchPolicyDetailResponse {
  s.PolicyResource = v
  return s
}

type QueryDispatchPolicyDetailPaths struct {
}

func (s QueryDispatchPolicyDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPolicyDetailPaths) GoString() string {
  return s.String()
}

type QueryDispatchPolicyDetailParameters struct {
}

func (s QueryDispatchPolicyDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPolicyDetailParameters) GoString() string {
  return s.String()
}

type QueryDispatchPolicyDetailRequestHeader struct {
}

func (s QueryDispatchPolicyDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPolicyDetailRequestHeader) GoString() string {
  return s.String()
}

type QueryDispatchPolicyDetailResponseHeader struct {
}

func (s QueryDispatchPolicyDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchPolicyDetailResponseHeader) GoString() string {
  return s.String()
}




type DelDispatchPolicyRequest struct {
  // {"en":"Domain that the dispatch policy to be deleted belongs to", "zh_CN":"要删除调度策略所属的域名"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"ID of the dispatch policy to be deleted Use English half-width semicolon between two policies if there are multiple to be deleted.", "zh_CN":"要删除的调度策略ID
  // 如果需要删除多个策略，用英文半角分号分隔。"}
  PolicyIds *string `json:"policyIds,omitempty" xml:"policyIds,omitempty" require:"true"`
  // {"en":"ID of the dispatch policy to be deleted Use English half-width semicolon between two policies if there are multiple to be deleted.", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s DelDispatchPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchPolicyRequest) GoString() string {
  return s.String()
}

func (s *DelDispatchPolicyRequest) SetDomainId(v int) *DelDispatchPolicyRequest {
  s.DomainId = &v
  return s
}

func (s *DelDispatchPolicyRequest) SetPolicyIds(v string) *DelDispatchPolicyRequest {
  s.PolicyIds = &v
  return s
}

func (s *DelDispatchPolicyRequest) SetLanguage(v string) *DelDispatchPolicyRequest {
  s.Language = &v
  return s
}

type DelDispatchPolicyResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Return PolicyID", "zh_CN":"返回策略ID"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s DelDispatchPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchPolicyResponse) GoString() string {
  return s.String()
}

func (s *DelDispatchPolicyResponse) SetResCode(v int) *DelDispatchPolicyResponse {
  s.ResCode = &v
  return s
}

func (s *DelDispatchPolicyResponse) SetMsg(v string) *DelDispatchPolicyResponse {
  s.Msg = &v
  return s
}

func (s *DelDispatchPolicyResponse) SetContent(v []*string) *DelDispatchPolicyResponse {
  s.Content = v
  return s
}

type DelDispatchPolicyPaths struct {
}

func (s DelDispatchPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchPolicyPaths) GoString() string {
  return s.String()
}

type DelDispatchPolicyParameters struct {
}

func (s DelDispatchPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchPolicyParameters) GoString() string {
  return s.String()
}

type DelDispatchPolicyRequestHeader struct {
}

func (s DelDispatchPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchPolicyRequestHeader) GoString() string {
  return s.String()
}

type DelDispatchPolicyResponseHeader struct {
}

func (s DelDispatchPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchPolicyResponseHeader) GoString() string {
  return s.String()
}




type ControlDispatchResourceRequest struct {
  // {"en":"The domain id", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"The policy id", "zh_CN":"策略id"}
  PolicyId *int `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"The resource value to operate (for example: 192.168.0.1 or xxxx.com)", "zh_CN":"要操作的资源值（例如:192.168.0.1或者xxxx.com）"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"The operation type, where 1 indicates the activation of resources and 0 indicates the deactivation of resources.", "zh_CN":"操作类型，1表示启用资源；0表示停用资源"}
  Operate *int `json:"operate,omitempty" xml:"operate,omitempty" require:"true"`
  // {"en":"If it is empty, the Chinese result will be returned by default; otherwise, the English prompt result will be returned.", "zh_CN":"为空返回中文结果(默认)，en:返回英文提示结果"}
  Language *int `json:"language,omitempty" xml:"language,omitempty" require:"true"`
}

func (s ControlDispatchResourceRequest) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchResourceRequest) GoString() string {
  return s.String()
}

func (s *ControlDispatchResourceRequest) SetDomainId(v int) *ControlDispatchResourceRequest {
  s.DomainId = &v
  return s
}

func (s *ControlDispatchResourceRequest) SetPolicyId(v int) *ControlDispatchResourceRequest {
  s.PolicyId = &v
  return s
}

func (s *ControlDispatchResourceRequest) SetValue(v string) *ControlDispatchResourceRequest {
  s.Value = &v
  return s
}

func (s *ControlDispatchResourceRequest) SetOperate(v int) *ControlDispatchResourceRequest {
  s.Operate = &v
  return s
}

func (s *ControlDispatchResourceRequest) SetLanguage(v int) *ControlDispatchResourceRequest {
  s.Language = &v
  return s
}

type ControlDispatchResourceResponse struct {
  // {"en":"The status code", "zh_CN":"状态码"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"详细说明"}
  Msg *int `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ControlDispatchResourceResponse) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchResourceResponse) GoString() string {
  return s.String()
}

func (s *ControlDispatchResourceResponse) SetResCode(v int) *ControlDispatchResourceResponse {
  s.ResCode = &v
  return s
}

func (s *ControlDispatchResourceResponse) SetMsg(v int) *ControlDispatchResourceResponse {
  s.Msg = &v
  return s
}

type ControlDispatchResourcePaths struct {
}

func (s ControlDispatchResourcePaths) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchResourcePaths) GoString() string {
  return s.String()
}

type ControlDispatchResourceParameters struct {
}

func (s ControlDispatchResourceParameters) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchResourceParameters) GoString() string {
  return s.String()
}

type ControlDispatchResourceRequestHeader struct {
}

func (s ControlDispatchResourceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchResourceRequestHeader) GoString() string {
  return s.String()
}

type ControlDispatchResourceResponseHeader struct {
}

func (s ControlDispatchResourceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchResourceResponseHeader) GoString() string {
  return s.String()
}




