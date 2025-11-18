package security

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type LECHDeleteSecurityGroupRequest struct {
}

func (s LECHDeleteSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHDeleteSecurityGroupRequest) GoString() string {
  return s.String()
}

type LECHDeleteSecurityGroupRequestHeader struct {
}

func (s LECHDeleteSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHDeleteSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type LECHDeleteSecurityGroupPaths struct {
  // {"en":"Security Group id,multiple values separated by commas","zh_CN":"安全组id，多个用逗号分隔"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHDeleteSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHDeleteSecurityGroupPaths) GoString() string {
  return s.String()
}

func (s *LECHDeleteSecurityGroupPaths) SetId(v string) *LECHDeleteSecurityGroupPaths {
  s.Id = &v
  return s
}

type LECHDeleteSecurityGroupParameters struct {
}

func (s LECHDeleteSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHDeleteSecurityGroupParameters) GoString() string {
  return s.String()
}

type LECHDeleteSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LECHDeleteSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHDeleteSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *LECHDeleteSecurityGroupResponse) SetCode(v string) *LECHDeleteSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *LECHDeleteSecurityGroupResponse) SetMessage(v string) *LECHDeleteSecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *LECHDeleteSecurityGroupResponse) SetData(v string) *LECHDeleteSecurityGroupResponse {
  s.Data = &v
  return s
}

type LECHDeleteSecurityGroupResponseHeader struct {
}

func (s LECHDeleteSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHDeleteSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type EditSecurityGroupRulesRequest struct {
  // {"en":"","zh_CN":""}
  SecurityGroupRule *EditSecurityGroupRulesRequestSecurityGroupRule `json:"securityGroupRule,omitempty" xml:"securityGroupRule,omitempty" type:"Struct"`
}

func (s EditSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRulesRequest) SetSecurityGroupRule(v *EditSecurityGroupRulesRequestSecurityGroupRule) *EditSecurityGroupRulesRequest {
  s.SecurityGroupRule = v
  return s
}

type EditSecurityGroupRulesRequestSecurityGroupRule struct {
  // {"en":"Security Group id","zh_CN":"规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS","zh_CN":"规则方向： INGRESS-流入 EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP","zh_CN":"授权策略： ACCEPT-允许 DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.When selecting ICMP, the maximum/minimum port numbers below cannot be specified.","zh_CN":"协议类型：TCP/UDP/ICMP 选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:4-IPv4  6-IPv6","zh_CN":"IP协议： 4-IPv4 6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority: 1-100, higher the value, higher the priority","zh_CN":"优先级：取值范围1-100，值越大优先级越高"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The maximum port number.1) Value 1-65535 and greater than or equal to the minimum port number.2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最大的端口号1）取值1-65535且大于等于最小端口号2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMax *string `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty"`
  // {"en":"The minimum port number.1) Value 1-65535 and less than or equal to the maximum port number.2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最小的端口号1）取值1-65535且小于等于最大端口号2）当要指定某个特定端口时，最大和最小端口号取值相等即可'"}
  PortRangeMin *string `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty"`
  // {"en":"Authorization object: The IP or CIDR that matches the security group rules.1) Specify multiple IPs or CIDR2, separated by commas.2) The specified IP or CIDR must match the IP protocol type mentioned above.","zh_CN":"授权对象：匹配该安全组规则的ip或cidr1）指定多个IP或CIDR,逗号分割 2）指定的IP或CIDR必须和上述的IP协议类型一致"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes,cannot exceed 255 characters","zh_CN":"备注信息（可选），不能超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s EditSecurityGroupRulesRequestSecurityGroupRule) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesRequestSecurityGroupRule) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetId(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Id = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetDirect(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Direct = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetPolicy(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Policy = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetProtocol(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Protocol = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetEthertype(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Ethertype = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetPriority(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Priority = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetPortRangeMax(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.PortRangeMax = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetPortRangeMin(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.PortRangeMin = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetRemoteIpPerfix(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.RemoteIpPerfix = &v
  return s
}

func (s *EditSecurityGroupRulesRequestSecurityGroupRule) SetRemark(v string) *EditSecurityGroupRulesRequestSecurityGroupRule {
  s.Remark = &v
  return s
}

type EditSecurityGroupRulesRequestHeader struct {
}

func (s EditSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesPaths struct {
}

func (s EditSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesParameters struct {
}

func (s EditSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRulesResponse) SetCode(v string) *EditSecurityGroupRulesResponse {
  s.Code = &v
  return s
}

func (s *EditSecurityGroupRulesResponse) SetMessage(v string) *EditSecurityGroupRulesResponse {
  s.Message = &v
  return s
}

func (s *EditSecurityGroupRulesResponse) SetData(v string) *EditSecurityGroupRulesResponse {
  s.Data = &v
  return s
}

type EditSecurityGroupRulesResponseHeader struct {
}

func (s EditSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type LECHQuerySecurityGroupRequest struct {
}

func (s LECHQuerySecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupRequest) GoString() string {
  return s.String()
}

type LECHQuerySecurityGroupRequestHeader struct {
}

func (s LECHQuerySecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type LECHQuerySecurityGroupPaths struct {
}

func (s LECHQuerySecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupPaths) GoString() string {
  return s.String()
}

type LECHQuerySecurityGroupParameters struct {
  // {"en":"Security Group ID","zh_CN":"根据安全组id查询"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Security Group name","zh_CN":"根据安全组名称查询"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s LECHQuerySecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupParameters) GoString() string {
  return s.String()
}

func (s *LECHQuerySecurityGroupParameters) SetId(v string) *LECHQuerySecurityGroupParameters {
  s.Id = &v
  return s
}

func (s *LECHQuerySecurityGroupParameters) SetName(v string) *LECHQuerySecurityGroupParameters {
  s.Name = &v
  return s
}

type LECHQuerySecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *LECHQuerySecurityGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHQuerySecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *LECHQuerySecurityGroupResponse) SetCode(v string) *LECHQuerySecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *LECHQuerySecurityGroupResponse) SetMessage(v string) *LECHQuerySecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *LECHQuerySecurityGroupResponse) SetData(v *LECHQuerySecurityGroupResponseData) *LECHQuerySecurityGroupResponse {
  s.Data = v
  return s
}

type LECHQuerySecurityGroupResponseData struct {
  // {"en":"","zh_CN":""}
  SecurityGroups []*LECHQuerySecurityGroupResponseDataSecurityGroups `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQuerySecurityGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupResponseData) GoString() string {
  return s.String()
}

func (s *LECHQuerySecurityGroupResponseData) SetSecurityGroups(v []*LECHQuerySecurityGroupResponseDataSecurityGroups) *LECHQuerySecurityGroupResponseData {
  s.SecurityGroups = v
  return s
}

type LECHQuerySecurityGroupResponseDataSecurityGroups struct     {
  // {"en":"Security Group ID","zh_CN":"安全组id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Security Group name","zh_CN":"安全组名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time:yyyy-MM-dd HH:mm:ss","zh_CN":"创建时间（yyyy-MM-dd HH:mm:ss）"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time:yyyy-MM-dd HH:mm:ss","zh_CN":"修改时间（yyyy-MM-dd HH:mm:ss）"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Rules []*LECHQuerySecurityGroupResponseDataSecurityGroupsRules `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Repeated"`
}

func (s LECHQuerySecurityGroupResponseDataSecurityGroups) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupResponseDataSecurityGroups) GoString() string {
  return s.String()
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroups) SetId(v string) *LECHQuerySecurityGroupResponseDataSecurityGroups {
  s.Id = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroups) SetName(v string) *LECHQuerySecurityGroupResponseDataSecurityGroups {
  s.Name = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroups) SetRemark(v string) *LECHQuerySecurityGroupResponseDataSecurityGroups {
  s.Remark = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroups) SetCreateTime(v string) *LECHQuerySecurityGroupResponseDataSecurityGroups {
  s.CreateTime = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroups) SetModifyTime(v string) *LECHQuerySecurityGroupResponseDataSecurityGroups {
  s.ModifyTime = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroups) SetRules(v []*LECHQuerySecurityGroupResponseDataSecurityGroupsRules) *LECHQuerySecurityGroupResponseDataSecurityGroups {
  s.Rules = v
  return s
}

type LECHQuerySecurityGroupResponseDataSecurityGroupsRules struct     {
  // {"en":"Rule Id","zh_CN":"安全组规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction","zh_CN":"规则方向"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy","zh_CN":"授权策略"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.When selecting ICMP, the maximum/minimum port numbers below cannot be specified.","zh_CN":"协议类型：TCP/UDP/ICMP 选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol","zh_CN":"ip协议"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Maximum Port Number","zh_CN":"最大端口号"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty" require:"true"`
  // {"en":"Minimum Port Number","zh_CN":"最小端口号"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty" require:"true"`
  // {"en":"IP or CIDR matching rules","zh_CN":"匹配规则的ip或cidr"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time:yyyy-MM-dd HH:mm:ss","zh_CN":"创建时间（yyyy-MM-dd HH:mm:ss）"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time:yyyy-MM-dd HH:mm:ss","zh_CN":"修改时间（yyyy-MM-dd HH:mm:ss）"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s LECHQuerySecurityGroupResponseDataSecurityGroupsRules) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupResponseDataSecurityGroupsRules) GoString() string {
  return s.String()
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetId(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Id = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetDirect(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Direct = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetPolicy(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Policy = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetProtocol(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Protocol = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetEthertype(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Ethertype = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetPriority(v int) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Priority = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetPortRangeMax(v int) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.PortRangeMax = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetPortRangeMin(v int) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.PortRangeMin = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetRemoteIpPerfix(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.RemoteIpPerfix = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetRemark(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Remark = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetCreateTime(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.CreateTime = &v
  return s
}

func (s *LECHQuerySecurityGroupResponseDataSecurityGroupsRules) SetModifyTime(v string) *LECHQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.ModifyTime = &v
  return s
}

type LECHQuerySecurityGroupResponseHeader struct {
}

func (s LECHQuerySecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type LECHCreateSSHKeyRequest struct {
  // {"en":"Creating array objects for SSH key","zh_CN":"创建SSH密钥对的数组对象"}
  Keypair *LECHCreateSSHKeyRequestKeypair `json:"keypair,omitempty" xml:"keypair,omitempty" require:"true" type:"Struct"`
}

func (s LECHCreateSSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyRequest) GoString() string {
  return s.String()
}

func (s *LECHCreateSSHKeyRequest) SetKeypair(v *LECHCreateSSHKeyRequestKeypair) *LECHCreateSSHKeyRequest {
  s.Keypair = v
  return s
}

type LECHCreateSSHKeyRequestKeypair struct {
  // {"en":"Key pair name, must be unique, otherwise creation will fail","zh_CN":"密钥对名称，必须保持唯一，否则会创建失败"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Imported public key information. The maximum length of the import public key is 1024 bytes. If this parameter is not specified, it will be generated automatically.","zh_CN":"导入的公钥信息。导入公钥最大长度为1024字节。如果未指定该参数，系统将自动生成。"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s LECHCreateSSHKeyRequestKeypair) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyRequestKeypair) GoString() string {
  return s.String()
}

func (s *LECHCreateSSHKeyRequestKeypair) SetName(v string) *LECHCreateSSHKeyRequestKeypair {
  s.Name = &v
  return s
}

func (s *LECHCreateSSHKeyRequestKeypair) SetPublicKey(v string) *LECHCreateSSHKeyRequestKeypair {
  s.PublicKey = &v
  return s
}

type LECHCreateSSHKeyRequestHeader struct {
}

func (s LECHCreateSSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyRequestHeader) GoString() string {
  return s.String()
}

type LECHCreateSSHKeyPaths struct {
}

func (s LECHCreateSSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyPaths) GoString() string {
  return s.String()
}

type LECHCreateSSHKeyParameters struct {
}

func (s LECHCreateSSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyParameters) GoString() string {
  return s.String()
}

type LECHCreateSSHKeyResponse struct {
  // {"en":"","zh_CN":""}
  Keypair []*LECHCreateSSHKeyResponseKeypair `json:"keypair,omitempty" xml:"keypair,omitempty" require:"true" type:"Repeated"`
}

func (s LECHCreateSSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyResponse) GoString() string {
  return s.String()
}

func (s *LECHCreateSSHKeyResponse) SetKeypair(v []*LECHCreateSSHKeyResponseKeypair) *LECHCreateSSHKeyResponse {
  s.Keypair = v
  return s
}

type LECHCreateSSHKeyResponseKeypair struct     {
  // {"en":"Key pair name, must be unique, otherwise creation will fail","zh_CN":"密钥对名称，必须保持唯一，否则会创建失败"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Imported public key information. The maximum length of the import public key is 1024 bytes. If this parameter is not specified, it will be generated automatically.","zh_CN":"导入的公钥信息。导入公钥最大长度为1024字节。如果未指定该参数，系统将自动生成。"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty" require:"true"`
  // {"en":"The key corresponds to the privatekey information. If publickey is given in the request parameter, the privatekey will not be returned. Otherwise, the system generated private key will be returned.","zh_CN":"密钥对应privateKey信息，如果请求参数中给了publicKey，则不会返回privateKey，否则返回系统生成的私钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
}

func (s LECHCreateSSHKeyResponseKeypair) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyResponseKeypair) GoString() string {
  return s.String()
}

func (s *LECHCreateSSHKeyResponseKeypair) SetName(v string) *LECHCreateSSHKeyResponseKeypair {
  s.Name = &v
  return s
}

func (s *LECHCreateSSHKeyResponseKeypair) SetPublicKey(v string) *LECHCreateSSHKeyResponseKeypair {
  s.PublicKey = &v
  return s
}

func (s *LECHCreateSSHKeyResponseKeypair) SetPrivateKey(v string) *LECHCreateSSHKeyResponseKeypair {
  s.PrivateKey = &v
  return s
}

type LECHCreateSSHKeyResponseHeader struct {
}

func (s LECHCreateSSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHCreateSSHKeyResponseHeader) GoString() string {
  return s.String()
}




type LECHDeletionSecurityGroupRulesRequest struct {
}

func (s LECHDeletionSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHDeletionSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

type LECHDeletionSecurityGroupRulesRequestHeader struct {
}

func (s LECHDeletionSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHDeletionSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type LECHDeletionSecurityGroupRulesPaths struct {
  // {"en":"Security group rule id, multiple separated by commas","zh_CN":"安全组规则id，多个用逗号分隔"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHDeletionSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHDeletionSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

func (s *LECHDeletionSecurityGroupRulesPaths) SetId(v string) *LECHDeletionSecurityGroupRulesPaths {
  s.Id = &v
  return s
}

type LECHDeletionSecurityGroupRulesParameters struct {
}

func (s LECHDeletionSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHDeletionSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type LECHDeletionSecurityGroupRulesResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LECHDeletionSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHDeletionSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *LECHDeletionSecurityGroupRulesResponse) SetCode(v string) *LECHDeletionSecurityGroupRulesResponse {
  s.Code = &v
  return s
}

func (s *LECHDeletionSecurityGroupRulesResponse) SetMessage(v string) *LECHDeletionSecurityGroupRulesResponse {
  s.Message = &v
  return s
}

func (s *LECHDeletionSecurityGroupRulesResponse) SetData(v string) *LECHDeletionSecurityGroupRulesResponse {
  s.Data = &v
  return s
}

type LECHDeletionSecurityGroupRulesResponseHeader struct {
}

func (s LECHDeletionSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHDeletionSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type AddSecurityGroupRulesRequest struct {
  // {"en":"Security Group Rule","zh_CN":"安全组规则"}
  SecurityGroupRules []*AddSecurityGroupRulesRequestSecurityGroupRules `json:"securityGroupRules,omitempty" xml:"securityGroupRules,omitempty" require:"true" type:"Repeated"`
}

func (s AddSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesRequest) SetSecurityGroupRules(v []*AddSecurityGroupRulesRequestSecurityGroupRules) *AddSecurityGroupRulesRequest {
  s.SecurityGroupRules = v
  return s
}

type AddSecurityGroupRulesRequestSecurityGroupRules struct     {
  // {"en":"Security Group id","zh_CN":"安全组id"}
  SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS","zh_CN":"规则方向：\nINGRESS-流入\nEGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP","zh_CN":"授权策略：\nACCEPT-允许\nDROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.\nWhen selecting ICMP, the maximum/minimum port numbers below cannot be specified.","zh_CN":"协议类型：TCP/UDP/ICMP\n选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:\n4-IPv4\n6-IPv6","zh_CN":"IP协议：\n4-IPv4\n6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority: 1-100, higher the value, higher the priority","zh_CN":"优先级：取值范围1-100，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The maximum port number.\n1) Value 1-65535 and greater than or equal to the minimum port number.\n2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最大的端口号\n1）取值1-65535且大于等于最小端口号\n2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty"`
  // {"en":"The minimum port number.\n1) Value 1-65535 and less than or equal to the maximum port number.\n2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最小的端口号\n1）取值1-65535且小于等于最大端口号\n2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty"`
  // {"en":"Authorization object: The IP or CIDR that matches the security group rules.1) Specify multiple IPs or CIDR2, separated by commas.2) The specified IP or CIDR must match the IP protocol type mentioned above.","zh_CN":"授权对象：匹配该安全组规则的ip或cidr1）指定多个IP或CIDR,逗号分割 2）指定的IP或CIDR必须和上述的IP协议类型一致"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes,cannot exceed 255 characters","zh_CN":"备注信息，不能超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s AddSecurityGroupRulesRequestSecurityGroupRules) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesRequestSecurityGroupRules) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetSecurityGroupId(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.SecurityGroupId = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetDirect(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.Direct = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetPolicy(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.Policy = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetProtocol(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.Protocol = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetEthertype(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.Ethertype = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetPriority(v int) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.Priority = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetPortRangeMax(v int) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.PortRangeMax = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetPortRangeMin(v int) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.PortRangeMin = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetRemoteIpPerfix(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.RemoteIpPerfix = &v
  return s
}

func (s *AddSecurityGroupRulesRequestSecurityGroupRules) SetRemark(v string) *AddSecurityGroupRulesRequestSecurityGroupRules {
  s.Remark = &v
  return s
}

type AddSecurityGroupRulesRequestHeader struct {
}

func (s AddSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type AddSecurityGroupRulesPaths struct {
}

func (s AddSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

type AddSecurityGroupRulesParameters struct {
}

func (s AddSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type AddSecurityGroupRulesResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *AddSecurityGroupRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesResponse) SetCode(v string) *AddSecurityGroupRulesResponse {
  s.Code = &v
  return s
}

func (s *AddSecurityGroupRulesResponse) SetMessage(v string) *AddSecurityGroupRulesResponse {
  s.Message = &v
  return s
}

func (s *AddSecurityGroupRulesResponse) SetData(v *AddSecurityGroupRulesResponseData) *AddSecurityGroupRulesResponse {
  s.Data = v
  return s
}

type AddSecurityGroupRulesResponseData struct {
  // {"en":"","zh_CN":""}
  SecurityGroupRules []*AddSecurityGroupRulesResponseDataSecurityGroupRules `json:"securityGroupRules,omitempty" xml:"securityGroupRules,omitempty" require:"true" type:"Repeated"`
}

func (s AddSecurityGroupRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesResponseData) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesResponseData) SetSecurityGroupRules(v []*AddSecurityGroupRulesResponseDataSecurityGroupRules) *AddSecurityGroupRulesResponseData {
  s.SecurityGroupRules = v
  return s
}

type AddSecurityGroupRulesResponseDataSecurityGroupRules struct     {
  // {"en":"Rule id","zh_CN":"规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS","zh_CN":"规则方向：INGRESS-流入EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP","zh_CN":"授权策略：ACCEPT-允许DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP","zh_CN":"协议类型：TCP/UDP/ICMP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:4-IPv46-IPv6","zh_CN":"IP协议：4-IPv46-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority, the higher the value, the higher the priority","zh_CN":"优先级，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Maximum Port Number","zh_CN":"最大的端口号"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty" require:"true"`
  // {"en":"Minimum Port Number","zh_CN":"最小的端口号"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty" require:"true"`
  // {"en":"Authorization object: IP or CIDR that matches the security group rules","zh_CN":"授权对象：匹配该安全组规则的IP或CIDR"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time","zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s AddSecurityGroupRulesResponseDataSecurityGroupRules) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesResponseDataSecurityGroupRules) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetId(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Id = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetDirect(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Direct = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetPolicy(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Policy = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetProtocol(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Protocol = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetEthertype(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Ethertype = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetPriority(v int) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Priority = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetPortRangeMax(v int) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.PortRangeMax = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetPortRangeMin(v int) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.PortRangeMin = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetRemoteIpPerfix(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.RemoteIpPerfix = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetRemark(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Remark = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetCreateTime(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.CreateTime = &v
  return s
}

func (s *AddSecurityGroupRulesResponseDataSecurityGroupRules) SetModifyTime(v string) *AddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.ModifyTime = &v
  return s
}

type AddSecurityGroupRulesResponseHeader struct {
}

func (s AddSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type LECHRemoveSSHKeyRequest struct {
}

func (s LECHRemoveSSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveSSHKeyRequest) GoString() string {
  return s.String()
}

type LECHRemoveSSHKeyRequestHeader struct {
}

func (s LECHRemoveSSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveSSHKeyRequestHeader) GoString() string {
  return s.String()
}

type LECHRemoveSSHKeyPaths struct {
  // {"en":"ssh key name","zh_CN":"ssh key 名称"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty" require:"true"`
}

func (s LECHRemoveSSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveSSHKeyPaths) GoString() string {
  return s.String()
}

func (s *LECHRemoveSSHKeyPaths) SetKeyName(v string) *LECHRemoveSSHKeyPaths {
  s.KeyName = &v
  return s
}

type LECHRemoveSSHKeyParameters struct {
}

func (s LECHRemoveSSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveSSHKeyParameters) GoString() string {
  return s.String()
}

type LECHRemoveSSHKeyResponse struct {
}

func (s LECHRemoveSSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveSSHKeyResponse) GoString() string {
  return s.String()
}

type LECHRemoveSSHKeyResponseHeader struct {
}

func (s LECHRemoveSSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHRemoveSSHKeyResponseHeader) GoString() string {
  return s.String()
}




type LECHEditSecurityGroupRequest struct {
  // {"en":"","zh_CN":""}
  SecurityGroup *LECHEditSecurityGroupRequestSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Struct"`
}

func (s LECHEditSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *LECHEditSecurityGroupRequest) SetSecurityGroup(v *LECHEditSecurityGroupRequestSecurityGroup) *LECHEditSecurityGroupRequest {
  s.SecurityGroup = v
  return s
}

type LECHEditSecurityGroupRequestSecurityGroup struct {
  // {"en":"Security Group id","zh_CN":"安全组id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Security Group name.Constraint: can only consist of letters, numbers, and underscores, with a length of no more than 64 characters.It cannot be named default and must be unique.","zh_CN":"安全组名称 约束：只能由字母/数字/下划线组成，长度不超过64个字符，不能命名为default且名称唯一"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes,not exceeding 255 characters in length.","zh_CN":"备注信息，长度不超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s LECHEditSecurityGroupRequestSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRequestSecurityGroup) GoString() string {
  return s.String()
}

func (s *LECHEditSecurityGroupRequestSecurityGroup) SetId(v string) *LECHEditSecurityGroupRequestSecurityGroup {
  s.Id = &v
  return s
}

func (s *LECHEditSecurityGroupRequestSecurityGroup) SetName(v string) *LECHEditSecurityGroupRequestSecurityGroup {
  s.Name = &v
  return s
}

func (s *LECHEditSecurityGroupRequestSecurityGroup) SetRemark(v string) *LECHEditSecurityGroupRequestSecurityGroup {
  s.Remark = &v
  return s
}

type LECHEditSecurityGroupRequestHeader struct {
}

func (s LECHEditSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type LECHEditSecurityGroupPaths struct {
}

func (s LECHEditSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupPaths) GoString() string {
  return s.String()
}

type LECHEditSecurityGroupParameters struct {
}

func (s LECHEditSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupParameters) GoString() string {
  return s.String()
}

type LECHEditSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LECHEditSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *LECHEditSecurityGroupResponse) SetCode(v string) *LECHEditSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *LECHEditSecurityGroupResponse) SetMessage(v string) *LECHEditSecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *LECHEditSecurityGroupResponse) SetData(v string) *LECHEditSecurityGroupResponse {
  s.Data = &v
  return s
}

type LECHEditSecurityGroupResponseHeader struct {
}

func (s LECHEditSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type LECHQuerySSHKeyRequest struct {
  // {"en":"The number of items displayed on each page is 20 by default","zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the virtual machine ID specified by the marker","zh_CN":"从marker指定的实例id开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s LECHQuerySSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySSHKeyRequest) GoString() string {
  return s.String()
}

func (s *LECHQuerySSHKeyRequest) SetLimit(v int) *LECHQuerySSHKeyRequest {
  s.Limit = &v
  return s
}

func (s *LECHQuerySSHKeyRequest) SetMarker(v string) *LECHQuerySSHKeyRequest {
  s.Marker = &v
  return s
}

type LECHQuerySSHKeyRequestHeader struct {
}

func (s LECHQuerySSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySSHKeyRequestHeader) GoString() string {
  return s.String()
}

type LECHQuerySSHKeyPaths struct {
}

func (s LECHQuerySSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySSHKeyPaths) GoString() string {
  return s.String()
}

type LECHQuerySSHKeyParameters struct {
}

func (s LECHQuerySSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySSHKeyParameters) GoString() string {
  return s.String()
}

type LECHQuerySSHKeyResponse struct {
  // {"en":"Key pair","zh_CN":"密钥对"}
  Keypairs []*string `json:"keypairs,omitempty" xml:"keypairs,omitempty" require:"true" type:"Repeated"`
  // {"en":"Key pair name","zh_CN":"密钥对名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Public key information corresponding to key","zh_CN":"密钥对应publicKey信息"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty" require:"true"`
  // {"en":"Key pair id","zh_CN":"密钥对id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s LECHQuerySSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySSHKeyResponse) GoString() string {
  return s.String()
}

func (s *LECHQuerySSHKeyResponse) SetKeypairs(v []*string) *LECHQuerySSHKeyResponse {
  s.Keypairs = v
  return s
}

func (s *LECHQuerySSHKeyResponse) SetName(v string) *LECHQuerySSHKeyResponse {
  s.Name = &v
  return s
}

func (s *LECHQuerySSHKeyResponse) SetPublicKey(v string) *LECHQuerySSHKeyResponse {
  s.PublicKey = &v
  return s
}

func (s *LECHQuerySSHKeyResponse) SetId(v string) *LECHQuerySSHKeyResponse {
  s.Id = &v
  return s
}

type LECHQuerySSHKeyResponseHeader struct {
}

func (s LECHQuerySSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHQuerySSHKeyResponseHeader) GoString() string {
  return s.String()
}




type VMPBindSecurityGroupRequest struct {
  // {"en":"","zh_CN":""}
  BindInfo []*VMPBindSecurityGroupRequestBindInfo `json:"bindInfo,omitempty" xml:"bindInfo,omitempty" require:"true" type:"Repeated"`
}

func (s VMPBindSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupRequest) SetBindInfo(v []*VMPBindSecurityGroupRequestBindInfo) *VMPBindSecurityGroupRequest {
  s.BindInfo = v
  return s
}

type VMPBindSecurityGroupRequestBindInfo struct     {
  // {"en":"Instance ID","zh_CN":"实例id"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"Security Group ID","zh_CN":"安全组id"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
}

func (s VMPBindSecurityGroupRequestBindInfo) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupRequestBindInfo) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupRequestBindInfo) SetServerId(v string) *VMPBindSecurityGroupRequestBindInfo {
  s.ServerId = &v
  return s
}

func (s *VMPBindSecurityGroupRequestBindInfo) SetSecurityGroupIds(v []*string) *VMPBindSecurityGroupRequestBindInfo {
  s.SecurityGroupIds = v
  return s
}

type VMPBindSecurityGroupRequestHeader struct {
}

func (s VMPBindSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupPaths struct {
}

func (s VMPBindSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupPaths) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupParameters struct {
}

func (s VMPBindSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupParameters) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *VMPBindSecurityGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s VMPBindSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupResponse) SetCode(v string) *VMPBindSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *VMPBindSecurityGroupResponse) SetData(v *VMPBindSecurityGroupResponseData) *VMPBindSecurityGroupResponse {
  s.Data = v
  return s
}

func (s *VMPBindSecurityGroupResponse) SetMessage(v string) *VMPBindSecurityGroupResponse {
  s.Message = &v
  return s
}

type VMPBindSecurityGroupResponseData struct {
  // {"en":"","zh_CN":""}
  BatchErrorMsg []*VMPBindSecurityGroupResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s VMPBindSecurityGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupResponseData) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupResponseData) SetBatchErrorMsg(v []*VMPBindSecurityGroupResponseDataBatchErrorMsg) *VMPBindSecurityGroupResponseData {
  s.BatchErrorMsg = v
  return s
}

type VMPBindSecurityGroupResponseDataBatchErrorMsg struct     {
  // {"en":"error code","zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error key","zh_CN":"错误关键字"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"error message","zh_CN":"错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s VMPBindSecurityGroupResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupResponseDataBatchErrorMsg) SetCode(v string) *VMPBindSecurityGroupResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *VMPBindSecurityGroupResponseDataBatchErrorMsg) SetKey(v string) *VMPBindSecurityGroupResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *VMPBindSecurityGroupResponseDataBatchErrorMsg) SetMsg(v string) *VMPBindSecurityGroupResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type VMPBindSecurityGroupResponseHeader struct {
}

func (s VMPBindSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type LECHEditSecurityGroupRulesRequest struct {
  // {"en":"","zh_CN":""}
  SecurityGroupRule *LECHEditSecurityGroupRulesRequestSecurityGroupRule `json:"securityGroupRule,omitempty" xml:"securityGroupRule,omitempty" type:"Struct"`
}

func (s LECHEditSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

func (s *LECHEditSecurityGroupRulesRequest) SetSecurityGroupRule(v *LECHEditSecurityGroupRulesRequestSecurityGroupRule) *LECHEditSecurityGroupRulesRequest {
  s.SecurityGroupRule = v
  return s
}

type LECHEditSecurityGroupRulesRequestSecurityGroupRule struct {
  // {"en":"Security Group id","zh_CN":"规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS","zh_CN":"规则方向： INGRESS-流入 EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP","zh_CN":"授权策略： ACCEPT-允许 DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.When selecting ICMP, the maximum/minimum port numbers below cannot be specified.","zh_CN":"协议类型：TCP/UDP/ICMP 选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:4-IPv4  6-IPv6","zh_CN":"IP协议： 4-IPv4 6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority: 1-100, higher the value, higher the priority","zh_CN":"优先级：取值范围1-100，值越大优先级越高"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The maximum port number.1) Value 1-65535 and greater than or equal to the minimum port number.2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最大的端口号1）取值1-65535且大于等于最小端口号2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMax *string `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty"`
  // {"en":"The minimum port number.1) Value 1-65535 and less than or equal to the maximum port number.2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最小的端口号1）取值1-65535且小于等于最大端口号2）当要指定某个特定端口时，最大和最小端口号取值相等即可'"}
  PortRangeMin *string `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty"`
  // {"en":"Authorization object: The IP or CIDR that matches the security group rules.1) Specify multiple IPs or CIDR2, separated by commas.2) The specified IP or CIDR must match the IP protocol type mentioned above.","zh_CN":"授权对象：匹配该安全组规则的ip或cidr1）指定多个IP或CIDR,逗号分割 2）指定的IP或CIDR必须和上述的IP协议类型一致"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes,cannot exceed 255 characters","zh_CN":"备注信息（可选），不能超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s LECHEditSecurityGroupRulesRequestSecurityGroupRule) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesRequestSecurityGroupRule) GoString() string {
  return s.String()
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetId(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Id = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetDirect(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Direct = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetPolicy(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Policy = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetProtocol(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Protocol = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetEthertype(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Ethertype = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetPriority(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Priority = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetPortRangeMax(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.PortRangeMax = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetPortRangeMin(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.PortRangeMin = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetRemoteIpPerfix(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.RemoteIpPerfix = &v
  return s
}

func (s *LECHEditSecurityGroupRulesRequestSecurityGroupRule) SetRemark(v string) *LECHEditSecurityGroupRulesRequestSecurityGroupRule {
  s.Remark = &v
  return s
}

type LECHEditSecurityGroupRulesRequestHeader struct {
}

func (s LECHEditSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type LECHEditSecurityGroupRulesPaths struct {
}

func (s LECHEditSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

type LECHEditSecurityGroupRulesParameters struct {
}

func (s LECHEditSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type LECHEditSecurityGroupRulesResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LECHEditSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *LECHEditSecurityGroupRulesResponse) SetCode(v string) *LECHEditSecurityGroupRulesResponse {
  s.Code = &v
  return s
}

func (s *LECHEditSecurityGroupRulesResponse) SetMessage(v string) *LECHEditSecurityGroupRulesResponse {
  s.Message = &v
  return s
}

func (s *LECHEditSecurityGroupRulesResponse) SetData(v string) *LECHEditSecurityGroupRulesResponse {
  s.Data = &v
  return s
}

type LECHEditSecurityGroupRulesResponseHeader struct {
}

func (s LECHEditSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHEditSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type DeleteSecurityGroupRequest struct {
}

func (s DeleteSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
  return s.String()
}

type DeleteSecurityGroupRequestHeader struct {
}

func (s DeleteSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type DeleteSecurityGroupPaths struct {
  // {"en":"Security Group id,multiple values separated by commas","zh_CN":"安全组id，多个用逗号分隔"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupPaths) GoString() string {
  return s.String()
}

func (s *DeleteSecurityGroupPaths) SetId(v string) *DeleteSecurityGroupPaths {
  s.Id = &v
  return s
}

type DeleteSecurityGroupParameters struct {
}

func (s DeleteSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupParameters) GoString() string {
  return s.String()
}

type DeleteSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *DeleteSecurityGroupResponse) SetCode(v string) *DeleteSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *DeleteSecurityGroupResponse) SetMessage(v string) *DeleteSecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *DeleteSecurityGroupResponse) SetData(v string) *DeleteSecurityGroupResponse {
  s.Data = &v
  return s
}

type DeleteSecurityGroupResponseHeader struct {
}

func (s DeleteSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type LECHBindSecurityGroupRequest struct {
  // {"en":"","zh_CN":""}
  BindInfo []*LECHBindSecurityGroupRequestBindInfo `json:"bindInfo,omitempty" xml:"bindInfo,omitempty" require:"true" type:"Repeated"`
}

func (s LECHBindSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *LECHBindSecurityGroupRequest) SetBindInfo(v []*LECHBindSecurityGroupRequestBindInfo) *LECHBindSecurityGroupRequest {
  s.BindInfo = v
  return s
}

type LECHBindSecurityGroupRequestBindInfo struct     {
  // {"en":"Instance ID","zh_CN":"实例id"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"Security Group ID","zh_CN":"安全组id"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
}

func (s LECHBindSecurityGroupRequestBindInfo) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupRequestBindInfo) GoString() string {
  return s.String()
}

func (s *LECHBindSecurityGroupRequestBindInfo) SetServerId(v string) *LECHBindSecurityGroupRequestBindInfo {
  s.ServerId = &v
  return s
}

func (s *LECHBindSecurityGroupRequestBindInfo) SetSecurityGroupIds(v []*string) *LECHBindSecurityGroupRequestBindInfo {
  s.SecurityGroupIds = v
  return s
}

type LECHBindSecurityGroupRequestHeader struct {
}

func (s LECHBindSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type LECHBindSecurityGroupPaths struct {
}

func (s LECHBindSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupPaths) GoString() string {
  return s.String()
}

type LECHBindSecurityGroupParameters struct {
}

func (s LECHBindSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupParameters) GoString() string {
  return s.String()
}

type LECHBindSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *LECHBindSecurityGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LECHBindSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *LECHBindSecurityGroupResponse) SetCode(v string) *LECHBindSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *LECHBindSecurityGroupResponse) SetData(v *LECHBindSecurityGroupResponseData) *LECHBindSecurityGroupResponse {
  s.Data = v
  return s
}

func (s *LECHBindSecurityGroupResponse) SetMessage(v string) *LECHBindSecurityGroupResponse {
  s.Message = &v
  return s
}

type LECHBindSecurityGroupResponseData struct {
  // {"en":"","zh_CN":""}
  BatchErrorMsg []*LECHBindSecurityGroupResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s LECHBindSecurityGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupResponseData) GoString() string {
  return s.String()
}

func (s *LECHBindSecurityGroupResponseData) SetBatchErrorMsg(v []*LECHBindSecurityGroupResponseDataBatchErrorMsg) *LECHBindSecurityGroupResponseData {
  s.BatchErrorMsg = v
  return s
}

type LECHBindSecurityGroupResponseDataBatchErrorMsg struct     {
  // {"en":"error code","zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error key","zh_CN":"错误关键字"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"error message","zh_CN":"错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s LECHBindSecurityGroupResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *LECHBindSecurityGroupResponseDataBatchErrorMsg) SetCode(v string) *LECHBindSecurityGroupResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *LECHBindSecurityGroupResponseDataBatchErrorMsg) SetKey(v string) *LECHBindSecurityGroupResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *LECHBindSecurityGroupResponseDataBatchErrorMsg) SetMsg(v string) *LECHBindSecurityGroupResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type LECHBindSecurityGroupResponseHeader struct {
}

func (s LECHBindSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHBindSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type DeletionSecurityGroupRulesRequest struct {
}

func (s DeletionSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesRequestHeader struct {
}

func (s DeletionSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesPaths struct {
  // {"en":"Security group rule id, multiple separated by commas","zh_CN":"安全组规则id，多个用逗号分隔"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeletionSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

func (s *DeletionSecurityGroupRulesPaths) SetId(v string) *DeletionSecurityGroupRulesPaths {
  s.Id = &v
  return s
}

type DeletionSecurityGroupRulesParameters struct {
}

func (s DeletionSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeletionSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *DeletionSecurityGroupRulesResponse) SetCode(v string) *DeletionSecurityGroupRulesResponse {
  s.Code = &v
  return s
}

func (s *DeletionSecurityGroupRulesResponse) SetMessage(v string) *DeletionSecurityGroupRulesResponse {
  s.Message = &v
  return s
}

func (s *DeletionSecurityGroupRulesResponse) SetData(v string) *DeletionSecurityGroupRulesResponse {
  s.Data = &v
  return s
}

type DeletionSecurityGroupRulesResponseHeader struct {
}

func (s DeletionSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type LECHAddSecurityGroupRequest struct {
  // {"en":"","zh_CN":""}
  SecurityGroup *LECHAddSecurityGroupRequestSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
}

func (s LECHAddSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRequest) SetSecurityGroup(v *LECHAddSecurityGroupRequestSecurityGroup) *LECHAddSecurityGroupRequest {
  s.SecurityGroup = v
  return s
}

type LECHAddSecurityGroupRequestSecurityGroup struct {
  // {"en":"Security group name.Constraint: can only consist of letters, numbers, and underscores, with a length of no more than 64 characters.It cannot be named default and must be unique.","zh_CN":"安全组名称。约束：只能由字母/数字/下划线组成，长度不超过64个字符，不能命名为default且名称唯一。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes (optional), no more than 255 characters in length","zh_CN":"备注（可选），长度不超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s LECHAddSecurityGroupRequestSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRequestSecurityGroup) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRequestSecurityGroup) SetName(v string) *LECHAddSecurityGroupRequestSecurityGroup {
  s.Name = &v
  return s
}

func (s *LECHAddSecurityGroupRequestSecurityGroup) SetRemark(v string) *LECHAddSecurityGroupRequestSecurityGroup {
  s.Remark = &v
  return s
}

type LECHAddSecurityGroupRequestHeader struct {
}

func (s LECHAddSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type LECHAddSecurityGroupPaths struct {
}

func (s LECHAddSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupPaths) GoString() string {
  return s.String()
}

type LECHAddSecurityGroupParameters struct {
}

func (s LECHAddSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupParameters) GoString() string {
  return s.String()
}

type LECHAddSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *LECHAddSecurityGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHAddSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupResponse) SetCode(v string) *LECHAddSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *LECHAddSecurityGroupResponse) SetMessage(v string) *LECHAddSecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *LECHAddSecurityGroupResponse) SetData(v *LECHAddSecurityGroupResponseData) *LECHAddSecurityGroupResponse {
  s.Data = v
  return s
}

type LECHAddSecurityGroupResponseData struct {
  // {"en":"","zh_CN":""}
  SecurityGroup *LECHAddSecurityGroupResponseDataSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Struct"`
}

func (s LECHAddSecurityGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupResponseData) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupResponseData) SetSecurityGroup(v *LECHAddSecurityGroupResponseDataSecurityGroup) *LECHAddSecurityGroupResponseData {
  s.SecurityGroup = v
  return s
}

type LECHAddSecurityGroupResponseDataSecurityGroup struct {
  // {"en":"Security Group ID","zh_CN":"安全组id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Security Group name","zh_CN":"安全组名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time","zh_CN":"创建时间（yyyy-MM-dd HH:mm:ss）"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time","zh_CN":"修改时间（yyyy-MM-dd HH:mm:ss）"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s LECHAddSecurityGroupResponseDataSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupResponseDataSecurityGroup) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupResponseDataSecurityGroup) SetId(v string) *LECHAddSecurityGroupResponseDataSecurityGroup {
  s.Id = &v
  return s
}

func (s *LECHAddSecurityGroupResponseDataSecurityGroup) SetName(v string) *LECHAddSecurityGroupResponseDataSecurityGroup {
  s.Name = &v
  return s
}

func (s *LECHAddSecurityGroupResponseDataSecurityGroup) SetRemark(v string) *LECHAddSecurityGroupResponseDataSecurityGroup {
  s.Remark = &v
  return s
}

func (s *LECHAddSecurityGroupResponseDataSecurityGroup) SetCreateTime(v string) *LECHAddSecurityGroupResponseDataSecurityGroup {
  s.CreateTime = &v
  return s
}

func (s *LECHAddSecurityGroupResponseDataSecurityGroup) SetModifyTime(v string) *LECHAddSecurityGroupResponseDataSecurityGroup {
  s.ModifyTime = &v
  return s
}

type LECHAddSecurityGroupResponseHeader struct {
}

func (s LECHAddSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type VMPCreateSSHKeyRequest struct {
  // {"en":"Creating array objects for SSH key", "zh_CN":"创建SSH密钥对的数组对象"}
  Keypair *VMPCreateSSHKeyKeypairCreate `json:"keypair,omitempty" xml:"keypair,omitempty" require:"true"`
}

func (s VMPCreateSSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyRequest) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyRequest) SetKeypair(v *VMPCreateSSHKeyKeypairCreate) *VMPCreateSSHKeyRequest {
  s.Keypair = v
  return s
}

type VMPCreateSSHKeyKeypairCreate struct {
  // {"en":"Key pair name, must be unique, otherwise creation will fail", "zh_CN":"密钥对名称，必须保持唯一，否则会创建失败"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Imported public key information. The maximum length of the import public key is 1024 bytes. If this parameter is not specified, it will be generated automatically.", "zh_CN":"导入的公钥信息。导入公钥最大长度为1024字节。如果未指定该参数，系统将自动生成。"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s VMPCreateSSHKeyKeypairCreate) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyKeypairCreate) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyKeypairCreate) SetName(v string) *VMPCreateSSHKeyKeypairCreate {
  s.Name = &v
  return s
}

func (s *VMPCreateSSHKeyKeypairCreate) SetPublicKey(v string) *VMPCreateSSHKeyKeypairCreate {
  s.PublicKey = &v
  return s
}

type VMPCreateSSHKeyKeypairResult struct {
  // {"en":"Key pair name, must be unique, otherwise creation will fail", "zh_CN":"密钥对名称，必须保持唯一，否则会创建失败"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Imported public key information. The maximum length of the import public key is 1024 bytes. If this parameter is not specified, it will be generated automatically.", "zh_CN":"导入的公钥信息。导入公钥最大长度为1024字节。如果未指定该参数，系统将自动生成。"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty" require:"true"`
  // {"en":"The key corresponds to the privatekey information. If publickey is given in the request parameter, the privatekey will not be returned. Otherwise, the system generated private key will be returned.", "zh_CN":"密钥对应privateKey信息，如果请求参数中给了publicKey，则不会返回privateKey，否则返回系统生成的私钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
}

func (s VMPCreateSSHKeyKeypairResult) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyKeypairResult) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyKeypairResult) SetName(v string) *VMPCreateSSHKeyKeypairResult {
  s.Name = &v
  return s
}

func (s *VMPCreateSSHKeyKeypairResult) SetPublicKey(v string) *VMPCreateSSHKeyKeypairResult {
  s.PublicKey = &v
  return s
}

func (s *VMPCreateSSHKeyKeypairResult) SetPrivateKey(v string) *VMPCreateSSHKeyKeypairResult {
  s.PrivateKey = &v
  return s
}

type VMPCreateSSHKeyResponse struct {
  Keypair []*VMPCreateSSHKeyKeypairResult `json:"keypair,omitempty" xml:"keypair,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateSSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyResponse) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyResponse) SetKeypair(v []*VMPCreateSSHKeyKeypairResult) *VMPCreateSSHKeyResponse {
  s.Keypair = v
  return s
}

type VMPCreateSSHKeyPaths struct {
}

func (s VMPCreateSSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyPaths) GoString() string {
  return s.String()
}

type VMPCreateSSHKeyParameters struct {
}

func (s VMPCreateSSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyParameters) GoString() string {
  return s.String()
}

type VMPCreateSSHKeyRequestHeader struct {
}

func (s VMPCreateSSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyRequestHeader) GoString() string {
  return s.String()
}

type VMPCreateSSHKeyResponseHeader struct {
}

func (s VMPCreateSSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyResponseHeader) GoString() string {
  return s.String()
}




type LECHAddSecurityGroupRulesRequest struct {
  // {"en":"Security Group Rule","zh_CN":"安全组规则"}
  SecurityGroupRules []*LECHAddSecurityGroupRulesRequestSecurityGroupRules `json:"securityGroupRules,omitempty" xml:"securityGroupRules,omitempty" require:"true" type:"Repeated"`
}

func (s LECHAddSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRulesRequest) SetSecurityGroupRules(v []*LECHAddSecurityGroupRulesRequestSecurityGroupRules) *LECHAddSecurityGroupRulesRequest {
  s.SecurityGroupRules = v
  return s
}

type LECHAddSecurityGroupRulesRequestSecurityGroupRules struct     {
  // {"en":"Security Group id","zh_CN":"安全组id"}
  SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS","zh_CN":"规则方向：\nINGRESS-流入\nEGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP","zh_CN":"授权策略：\nACCEPT-允许\nDROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.\nWhen selecting ICMP, the maximum/minimum port numbers below cannot be specified.","zh_CN":"协议类型：TCP/UDP/ICMP\n选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:\n4-IPv4\n6-IPv6","zh_CN":"IP协议：\n4-IPv4\n6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority: 1-100, higher the value, higher the priority","zh_CN":"优先级：取值范围1-100，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The maximum port number.\n1) Value 1-65535 and greater than or equal to the minimum port number.\n2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最大的端口号\n1）取值1-65535且大于等于最小端口号\n2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty"`
  // {"en":"The minimum port number.\n1) Value 1-65535 and less than or equal to the maximum port number.\n2) When specifying a specific port, the maximum and minimum port numbers should be equal.","zh_CN":"最小的端口号\n1）取值1-65535且小于等于最大端口号\n2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty"`
  // {"en":"Authorization object: The IP or CIDR that matches the security group rules.1) Specify multiple IPs or CIDR2, separated by commas.2) The specified IP or CIDR must match the IP protocol type mentioned above.","zh_CN":"授权对象：匹配该安全组规则的ip或cidr1）指定多个IP或CIDR,逗号分割 2）指定的IP或CIDR必须和上述的IP协议类型一致"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes,cannot exceed 255 characters","zh_CN":"备注信息，不能超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s LECHAddSecurityGroupRulesRequestSecurityGroupRules) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesRequestSecurityGroupRules) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetSecurityGroupId(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.SecurityGroupId = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetDirect(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.Direct = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetPolicy(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.Policy = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetProtocol(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.Protocol = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetEthertype(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.Ethertype = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetPriority(v int) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.Priority = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetPortRangeMax(v int) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.PortRangeMax = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetPortRangeMin(v int) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.PortRangeMin = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetRemoteIpPerfix(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.RemoteIpPerfix = &v
  return s
}

func (s *LECHAddSecurityGroupRulesRequestSecurityGroupRules) SetRemark(v string) *LECHAddSecurityGroupRulesRequestSecurityGroupRules {
  s.Remark = &v
  return s
}

type LECHAddSecurityGroupRulesRequestHeader struct {
}

func (s LECHAddSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type LECHAddSecurityGroupRulesPaths struct {
}

func (s LECHAddSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

type LECHAddSecurityGroupRulesParameters struct {
}

func (s LECHAddSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type LECHAddSecurityGroupRulesResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *LECHAddSecurityGroupRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s LECHAddSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRulesResponse) SetCode(v string) *LECHAddSecurityGroupRulesResponse {
  s.Code = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponse) SetMessage(v string) *LECHAddSecurityGroupRulesResponse {
  s.Message = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponse) SetData(v *LECHAddSecurityGroupRulesResponseData) *LECHAddSecurityGroupRulesResponse {
  s.Data = v
  return s
}

type LECHAddSecurityGroupRulesResponseData struct {
  // {"en":"","zh_CN":""}
  SecurityGroupRules []*LECHAddSecurityGroupRulesResponseDataSecurityGroupRules `json:"securityGroupRules,omitempty" xml:"securityGroupRules,omitempty" require:"true" type:"Repeated"`
}

func (s LECHAddSecurityGroupRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesResponseData) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRulesResponseData) SetSecurityGroupRules(v []*LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) *LECHAddSecurityGroupRulesResponseData {
  s.SecurityGroupRules = v
  return s
}

type LECHAddSecurityGroupRulesResponseDataSecurityGroupRules struct     {
  // {"en":"Rule id","zh_CN":"规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS","zh_CN":"规则方向：INGRESS-流入EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP","zh_CN":"授权策略：ACCEPT-允许DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP","zh_CN":"协议类型：TCP/UDP/ICMP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:4-IPv46-IPv6","zh_CN":"IP协议：4-IPv46-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority, the higher the value, the higher the priority","zh_CN":"优先级，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Maximum Port Number","zh_CN":"最大的端口号"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty" require:"true"`
  // {"en":"Minimum Port Number","zh_CN":"最小的端口号"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty" require:"true"`
  // {"en":"Authorization object: IP or CIDR that matches the security group rules","zh_CN":"授权对象：匹配该安全组规则的IP或CIDR"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time","zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) GoString() string {
  return s.String()
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetId(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Id = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetDirect(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Direct = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetPolicy(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Policy = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetProtocol(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Protocol = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetEthertype(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Ethertype = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetPriority(v int) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Priority = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetPortRangeMax(v int) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.PortRangeMax = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetPortRangeMin(v int) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.PortRangeMin = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetRemoteIpPerfix(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.RemoteIpPerfix = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetRemark(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.Remark = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetCreateTime(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.CreateTime = &v
  return s
}

func (s *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules) SetModifyTime(v string) *LECHAddSecurityGroupRulesResponseDataSecurityGroupRules {
  s.ModifyTime = &v
  return s
}

type LECHAddSecurityGroupRulesResponseHeader struct {
}

func (s LECHAddSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LECHAddSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type EditSecurityGroupRequest struct {
  // {"en":"","zh_CN":""}
  SecurityGroup *EditSecurityGroupRequestSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Struct"`
}

func (s EditSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRequest) SetSecurityGroup(v *EditSecurityGroupRequestSecurityGroup) *EditSecurityGroupRequest {
  s.SecurityGroup = v
  return s
}

type EditSecurityGroupRequestSecurityGroup struct {
  // {"en":"Security Group id","zh_CN":"安全组id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Security Group name.Constraint: can only consist of letters, numbers, and underscores, with a length of no more than 64 characters.It cannot be named default and must be unique.","zh_CN":"安全组名称 约束：只能由字母/数字/下划线组成，长度不超过64个字符，不能命名为default且名称唯一"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes,not exceeding 255 characters in length.","zh_CN":"备注信息，长度不超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s EditSecurityGroupRequestSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRequestSecurityGroup) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRequestSecurityGroup) SetId(v string) *EditSecurityGroupRequestSecurityGroup {
  s.Id = &v
  return s
}

func (s *EditSecurityGroupRequestSecurityGroup) SetName(v string) *EditSecurityGroupRequestSecurityGroup {
  s.Name = &v
  return s
}

func (s *EditSecurityGroupRequestSecurityGroup) SetRemark(v string) *EditSecurityGroupRequestSecurityGroup {
  s.Remark = &v
  return s
}

type EditSecurityGroupRequestHeader struct {
}

func (s EditSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type EditSecurityGroupPaths struct {
}

func (s EditSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupPaths) GoString() string {
  return s.String()
}

type EditSecurityGroupParameters struct {
}

func (s EditSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupParameters) GoString() string {
  return s.String()
}

type EditSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupResponse) SetCode(v string) *EditSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *EditSecurityGroupResponse) SetMessage(v string) *EditSecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *EditSecurityGroupResponse) SetData(v string) *EditSecurityGroupResponse {
  s.Data = &v
  return s
}

type EditSecurityGroupResponseHeader struct {
}

func (s EditSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type VMPRemoveSSHKeyRequest struct {
}

func (s VMPRemoveSSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyRequest) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyRequestHeader struct {
}

func (s VMPRemoveSSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyRequestHeader) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyPaths struct {
  // {"en":"ssh key name","zh_CN":"ssh key 名称"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty" require:"true"`
}

func (s VMPRemoveSSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyPaths) GoString() string {
  return s.String()
}

func (s *VMPRemoveSSHKeyPaths) SetKeyName(v string) *VMPRemoveSSHKeyPaths {
  s.KeyName = &v
  return s
}

type VMPRemoveSSHKeyParameters struct {
}

func (s VMPRemoveSSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyParameters) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyResponse struct {
}

func (s VMPRemoveSSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyResponse) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyResponseHeader struct {
}

func (s VMPRemoveSSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyResponseHeader) GoString() string {
  return s.String()
}




type VMPQuerySecurityGroupRequest struct {
}

func (s VMPQuerySecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupRequest) GoString() string {
  return s.String()
}

type VMPQuerySecurityGroupRequestHeader struct {
}

func (s VMPQuerySecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type VMPQuerySecurityGroupPaths struct {
}

func (s VMPQuerySecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupPaths) GoString() string {
  return s.String()
}

type VMPQuerySecurityGroupParameters struct {
  // {"en":"Security Group ID","zh_CN":"根据安全组id查询"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Security Group name","zh_CN":"根据安全组名称查询"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s VMPQuerySecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupParameters) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupParameters) SetId(v string) *VMPQuerySecurityGroupParameters {
  s.Id = &v
  return s
}

func (s *VMPQuerySecurityGroupParameters) SetName(v string) *VMPQuerySecurityGroupParameters {
  s.Name = &v
  return s
}

type VMPQuerySecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *VMPQuerySecurityGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s VMPQuerySecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupResponse) SetCode(v string) *VMPQuerySecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *VMPQuerySecurityGroupResponse) SetMessage(v string) *VMPQuerySecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *VMPQuerySecurityGroupResponse) SetData(v *VMPQuerySecurityGroupResponseData) *VMPQuerySecurityGroupResponse {
  s.Data = v
  return s
}

type VMPQuerySecurityGroupResponseData struct {
  // {"en":"","zh_CN":""}
  SecurityGroups []*VMPQuerySecurityGroupResponseDataSecurityGroups `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQuerySecurityGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponseData) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupResponseData) SetSecurityGroups(v []*VMPQuerySecurityGroupResponseDataSecurityGroups) *VMPQuerySecurityGroupResponseData {
  s.SecurityGroups = v
  return s
}

type VMPQuerySecurityGroupResponseDataSecurityGroups struct     {
  // {"en":"Security Group ID","zh_CN":"安全组id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Security Group name","zh_CN":"安全组名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time:yyyy-MM-dd HH:mm:ss","zh_CN":"创建时间（yyyy-MM-dd HH:mm:ss）"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time:yyyy-MM-dd HH:mm:ss","zh_CN":"修改时间（yyyy-MM-dd HH:mm:ss）"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Rules []*VMPQuerySecurityGroupResponseDataSecurityGroupsRules `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQuerySecurityGroupResponseDataSecurityGroups) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponseDataSecurityGroups) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroups) SetId(v string) *VMPQuerySecurityGroupResponseDataSecurityGroups {
  s.Id = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroups) SetName(v string) *VMPQuerySecurityGroupResponseDataSecurityGroups {
  s.Name = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroups) SetRemark(v string) *VMPQuerySecurityGroupResponseDataSecurityGroups {
  s.Remark = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroups) SetCreateTime(v string) *VMPQuerySecurityGroupResponseDataSecurityGroups {
  s.CreateTime = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroups) SetModifyTime(v string) *VMPQuerySecurityGroupResponseDataSecurityGroups {
  s.ModifyTime = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroups) SetRules(v []*VMPQuerySecurityGroupResponseDataSecurityGroupsRules) *VMPQuerySecurityGroupResponseDataSecurityGroups {
  s.Rules = v
  return s
}

type VMPQuerySecurityGroupResponseDataSecurityGroupsRules struct     {
  // {"en":"Rule Id","zh_CN":"安全组规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction","zh_CN":"规则方向"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy","zh_CN":"授权策略"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.When selecting ICMP, the maximum/minimum port numbers below cannot be specified.","zh_CN":"协议类型：TCP/UDP/ICMP 选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol","zh_CN":"ip协议"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Maximum Port Number","zh_CN":"最大端口号"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty" require:"true"`
  // {"en":"Minimum Port Number","zh_CN":"最小端口号"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty" require:"true"`
  // {"en":"IP or CIDR matching rules","zh_CN":"匹配规则的ip或cidr"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time:yyyy-MM-dd HH:mm:ss","zh_CN":"创建时间（yyyy-MM-dd HH:mm:ss）"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time:yyyy-MM-dd HH:mm:ss","zh_CN":"修改时间（yyyy-MM-dd HH:mm:ss）"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s VMPQuerySecurityGroupResponseDataSecurityGroupsRules) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponseDataSecurityGroupsRules) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetId(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Id = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetDirect(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Direct = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetPolicy(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Policy = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetProtocol(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Protocol = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetEthertype(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Ethertype = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetPriority(v int) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Priority = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetPortRangeMax(v int) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.PortRangeMax = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetPortRangeMin(v int) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.PortRangeMin = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetRemoteIpPerfix(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.RemoteIpPerfix = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetRemark(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.Remark = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetCreateTime(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.CreateTime = &v
  return s
}

func (s *VMPQuerySecurityGroupResponseDataSecurityGroupsRules) SetModifyTime(v string) *VMPQuerySecurityGroupResponseDataSecurityGroupsRules {
  s.ModifyTime = &v
  return s
}

type VMPQuerySecurityGroupResponseHeader struct {
}

func (s VMPQuerySecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type AddSecurityGroupRequest struct {
  // {"en":"","zh_CN":""}
  SecurityGroup *AddSecurityGroupRequestSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
}

func (s AddSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRequest) SetSecurityGroup(v *AddSecurityGroupRequestSecurityGroup) *AddSecurityGroupRequest {
  s.SecurityGroup = v
  return s
}

type AddSecurityGroupRequestSecurityGroup struct {
  // {"en":"Security group name.Constraint: can only consist of letters, numbers, and underscores, with a length of no more than 64 characters.It cannot be named default and must be unique.","zh_CN":"安全组名称。约束：只能由字母/数字/下划线组成，长度不超过64个字符，不能命名为default且名称唯一。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes (optional), no more than 255 characters in length","zh_CN":"备注（可选），长度不超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s AddSecurityGroupRequestSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRequestSecurityGroup) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRequestSecurityGroup) SetName(v string) *AddSecurityGroupRequestSecurityGroup {
  s.Name = &v
  return s
}

func (s *AddSecurityGroupRequestSecurityGroup) SetRemark(v string) *AddSecurityGroupRequestSecurityGroup {
  s.Remark = &v
  return s
}

type AddSecurityGroupRequestHeader struct {
}

func (s AddSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type AddSecurityGroupPaths struct {
}

func (s AddSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupPaths) GoString() string {
  return s.String()
}

type AddSecurityGroupParameters struct {
}

func (s AddSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupParameters) GoString() string {
  return s.String()
}

type AddSecurityGroupResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *AddSecurityGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupResponse) SetCode(v string) *AddSecurityGroupResponse {
  s.Code = &v
  return s
}

func (s *AddSecurityGroupResponse) SetMessage(v string) *AddSecurityGroupResponse {
  s.Message = &v
  return s
}

func (s *AddSecurityGroupResponse) SetData(v *AddSecurityGroupResponseData) *AddSecurityGroupResponse {
  s.Data = v
  return s
}

type AddSecurityGroupResponseData struct {
  // {"en":"","zh_CN":""}
  SecurityGroup *AddSecurityGroupResponseDataSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Struct"`
}

func (s AddSecurityGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupResponseData) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupResponseData) SetSecurityGroup(v *AddSecurityGroupResponseDataSecurityGroup) *AddSecurityGroupResponseData {
  s.SecurityGroup = v
  return s
}

type AddSecurityGroupResponseDataSecurityGroup struct {
  // {"en":"Security Group ID","zh_CN":"安全组id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Security Group name","zh_CN":"安全组名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Notes","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time","zh_CN":"创建时间（yyyy-MM-dd HH:mm:ss）"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time","zh_CN":"修改时间（yyyy-MM-dd HH:mm:ss）"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s AddSecurityGroupResponseDataSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupResponseDataSecurityGroup) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupResponseDataSecurityGroup) SetId(v string) *AddSecurityGroupResponseDataSecurityGroup {
  s.Id = &v
  return s
}

func (s *AddSecurityGroupResponseDataSecurityGroup) SetName(v string) *AddSecurityGroupResponseDataSecurityGroup {
  s.Name = &v
  return s
}

func (s *AddSecurityGroupResponseDataSecurityGroup) SetRemark(v string) *AddSecurityGroupResponseDataSecurityGroup {
  s.Remark = &v
  return s
}

func (s *AddSecurityGroupResponseDataSecurityGroup) SetCreateTime(v string) *AddSecurityGroupResponseDataSecurityGroup {
  s.CreateTime = &v
  return s
}

func (s *AddSecurityGroupResponseDataSecurityGroup) SetModifyTime(v string) *AddSecurityGroupResponseDataSecurityGroup {
  s.ModifyTime = &v
  return s
}

type AddSecurityGroupResponseHeader struct {
}

func (s AddSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type VMPQuerySSHKeyRequest struct {
  // {"en":"The number of items displayed on each page is 20 by default","zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the virtual machine ID specified by the marker","zh_CN":"从marker指定的实例id开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s VMPQuerySSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyRequest) GoString() string {
  return s.String()
}

func (s *VMPQuerySSHKeyRequest) SetLimit(v int) *VMPQuerySSHKeyRequest {
  s.Limit = &v
  return s
}

func (s *VMPQuerySSHKeyRequest) SetMarker(v string) *VMPQuerySSHKeyRequest {
  s.Marker = &v
  return s
}

type VMPQuerySSHKeyRequestHeader struct {
}

func (s VMPQuerySSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyRequestHeader) GoString() string {
  return s.String()
}

type VMPQuerySSHKeyPaths struct {
}

func (s VMPQuerySSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyPaths) GoString() string {
  return s.String()
}

type VMPQuerySSHKeyParameters struct {
}

func (s VMPQuerySSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyParameters) GoString() string {
  return s.String()
}

type VMPQuerySSHKeyResponse struct {
  // {"en":"Key pair","zh_CN":"密钥对"}
  Keypairs []*string `json:"keypairs,omitempty" xml:"keypairs,omitempty" require:"true" type:"Repeated"`
  // {"en":"Key pair name","zh_CN":"密钥对名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Public key information corresponding to key","zh_CN":"密钥对应publicKey信息"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty" require:"true"`
  // {"en":"Key pair id","zh_CN":"密钥对id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s VMPQuerySSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyResponse) GoString() string {
  return s.String()
}

func (s *VMPQuerySSHKeyResponse) SetKeypairs(v []*string) *VMPQuerySSHKeyResponse {
  s.Keypairs = v
  return s
}

func (s *VMPQuerySSHKeyResponse) SetName(v string) *VMPQuerySSHKeyResponse {
  s.Name = &v
  return s
}

func (s *VMPQuerySSHKeyResponse) SetPublicKey(v string) *VMPQuerySSHKeyResponse {
  s.PublicKey = &v
  return s
}

func (s *VMPQuerySSHKeyResponse) SetId(v string) *VMPQuerySSHKeyResponse {
  s.Id = &v
  return s
}

type VMPQuerySSHKeyResponseHeader struct {
}

func (s VMPQuerySSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyResponseHeader) GoString() string {
  return s.String()
}




