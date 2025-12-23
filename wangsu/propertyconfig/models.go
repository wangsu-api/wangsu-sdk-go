package propertyconfig

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreatePropertyForAkamaiMigrationRequest struct {
  // {"en":"Product Service Type related to your contract.","zh_CN":"产品服务类型。请根据您的合同产品服务类型填写。","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
  // {"en":"Property name. The length must not exceed 128 characters.","zh_CN":"项目的名称。长度不超过128个字符。"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"Akamai Property's RuleTree JSON.","zh_CN":"Akamai项目规则树JSON配置。"}
  AkProperty *string `json:"akProperty,omitempty" xml:"akProperty,omitempty" require:"true"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*CreatePropertyForAkamaiMigrationRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
  // {"en":"Add new version to this propertyId","zh_CN":"新增版本到已有的propertyId"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"Is filter config by hostname. Default value is '0'.","zh_CN":"是否按域名过滤配置。默认值为'0'","exampleValue":"0,1"}
  FilterConfigByHostname *string `json:"filterConfigByHostname,omitempty" xml:"filterConfigByHostname,omitempty"`
}

func (s CreatePropertyForAkamaiMigrationRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequest) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetServiceType(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.ServiceType = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetPropertyComment(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.PropertyComment = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetPropertyName(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetVersionComment(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.VersionComment = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetAkProperty(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.AkProperty = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetHostnames(v []*CreatePropertyForAkamaiMigrationRequestHostnames) *CreatePropertyForAkamaiMigrationRequest {
  s.Hostnames = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetContractId(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.ContractId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetItemId(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.ItemId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetPropertyId(v int) *CreatePropertyForAkamaiMigrationRequest {
  s.PropertyId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetFilterConfigByHostname(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.FilterConfigByHostname = &v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnames struct     {
  // {"en":"hostnames. The length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*). Multiple hostnames are separated by English semicolons(;).","zh_CN":"加速域名。长度不超过128个字符。泛域名需要以“*”开头，例如：`*.example.com`。多个域名以英文分号(;)分隔。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Associated edge certificate configurations.","zh_CN":"关联的边缘证书"}
  Certificates []*CreatePropertyForAkamaiMigrationRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"Edge Hostname","zh_CN":"调度域名"}
  EdgeHostname *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnames) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetHostname(v string) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetCertificates(v []*CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.Certificates = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetEdgeHostname(v *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.EdgeHostname = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetDefaultOrigin(v *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.DefaultOrigin = v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname struct {
  // {"en":"Edge Hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"Edge Hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"Edge Hostname description","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetScheme(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHeader struct {
  // {"en":"Property's Contract ID in Akamai.","zh_CN":"Akamai项目所属合同编号(Contract ID)。"}
  XAkAkamaiContractId *string `json:"x-ak-akamai-contract-id,omitempty" xml:"x-ak-akamai-contract-id,omitempty"`
  // {"en":"Customer account in Akamai.","zh_CN":"Akamai账号(Customer account)。"}
  XAkCustomerMainAccount *string `json:"x-ak-customer-main-account,omitempty" xml:"x-ak-customer-main-account,omitempty"`
  // {"en":"Property's Parent Group ID in Akamai.","zh_CN":"Akamai项目所属用户组的父组编号(Parent Group ID)。"}
  XAkParentGroupId *string `json:"x-ak-parent-group-id,omitempty" xml:"x-ak-parent-group-id,omitempty"`
  // {"en":"Property's Parent Group name in Akamai.","zh_CN":"Akamai项目所属用户组的父组名称(Parent Group name)。"}
  XAkParentGroupName *string `json:"x-ak-parent-group-name,omitempty" xml:"x-ak-parent-group-name,omitempty"`
  // {"en":"Property's group ID of the group.","zh_CN":"Akamai项目所属用户组编号(Group ID)。"}
  XAkGroupId *string `json:"x-ak-group-id,omitempty" xml:"x-ak-group-id,omitempty"`
  // {"en":"Property's Group name in Akamai.","zh_CN":"Akamai项目所属用户组名称(Group name)。"}
  XAkGroupName *string `json:"x-ak-group-name,omitempty" xml:"x-ak-group-name,omitempty"`
}

func (s CreatePropertyForAkamaiMigrationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHeader) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkAkamaiContractId(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkAkamaiContractId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkCustomerMainAccount(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkCustomerMainAccount = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkParentGroupId(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkParentGroupId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkParentGroupName(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkParentGroupName = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkGroupId(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkGroupId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkGroupName(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkGroupName = &v
  return s
}

type CreatePropertyForAkamaiMigrationPaths struct {
}

func (s CreatePropertyForAkamaiMigrationPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationPaths) GoString() string {
  return s.String()
}

type CreatePropertyForAkamaiMigrationParameters struct {
}

func (s CreatePropertyForAkamaiMigrationParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationParameters) GoString() string {
  return s.String()
}

type CreatePropertyForAkamaiMigrationResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreatePropertyForAkamaiMigrationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Metadate transformation details. Metadate include Akamai Behavior and Criteria directives.","zh_CN":"来源元数据转换详情。元数据包含了 Akamai 的 Behaviors 和 Criterias 指令。"}
  MetaDataMsg *CreatePropertyForAkamaiMigrationResponseMetaDataMsg `json:"metaDataMsg,omitempty" xml:"metaDataMsg,omitempty" require:"true" type:"Struct"`
}

func (s CreatePropertyForAkamaiMigrationResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponse) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetCode(v string) *CreatePropertyForAkamaiMigrationResponse {
  s.Code = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetData(v *CreatePropertyForAkamaiMigrationResponseData) *CreatePropertyForAkamaiMigrationResponse {
  s.Data = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetMessage(v string) *CreatePropertyForAkamaiMigrationResponse {
  s.Message = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetMetaDataMsg(v *CreatePropertyForAkamaiMigrationResponseMetaDataMsg) *CreatePropertyForAkamaiMigrationResponse {
  s.MetaDataMsg = v
  return s
}

type CreatePropertyForAkamaiMigrationResponseData struct {
  // {"en":"Property version.","zh_CN":"项目的版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"Property name.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyForAkamaiMigrationResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponseData) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationResponseData) SetPropertyVersion(v int) *CreatePropertyForAkamaiMigrationResponseData {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponseData) SetPropertyName(v string) *CreatePropertyForAkamaiMigrationResponseData {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponseData) SetPropertyId(v int64) *CreatePropertyForAkamaiMigrationResponseData {
  s.PropertyId = &v
  return s
}

type CreatePropertyForAkamaiMigrationResponseMetaDataMsg struct {
  // {"en":"Metadate transformation details. For Behavior/Criteria metadata in Akamai JSON, classifies transformation log level into 4 message types (INFO, NOTICE, WARN, ERR ) and 9 error code (100001, 100002, 200001 etc.)","zh_CN":"元数据转换结果详情。对Akamai JSON中的Behavior/Criteria元数据，打印了配置转换反馈日志，划分了4类消息类型（INFO, NOTICE, WARN, ERR）和9类错误码（100001，100002，200001等）。"}
  MsgData *string `json:"msgData,omitempty" xml:"msgData,omitempty" require:"true"`
  // {"en":"Statistical information of metadate transformation.  The Transformation success rate can be calculated based on statistical information. Calculation formula:success rate = (1- ERR/COUNT)100%.","zh_CN":"元数据转换结果统计信息。可根据统计信息计算转换成功率。计算公式：转换成功率 = （1- ERR/COUNT）*100%。"}
  MsgStat *int64 `json:"msgStat,omitempty" xml:"msgStat,omitempty" require:"true"`
}

func (s CreatePropertyForAkamaiMigrationResponseMetaDataMsg) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponseMetaDataMsg) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationResponseMetaDataMsg) SetMsgData(v string) *CreatePropertyForAkamaiMigrationResponseMetaDataMsg {
  s.MsgData = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponseMetaDataMsg) SetMsgStat(v int64) *CreatePropertyForAkamaiMigrationResponseMetaDataMsg {
  s.MsgStat = &v
  return s
}

type CreatePropertyForAkamaiMigrationResponseHeader struct {
}

func (s CreatePropertyForAkamaiMigrationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponseHeader) GoString() string {
  return s.String()
}




type GetPropertyVersionRequest struct {
}

func (s GetPropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionRequest) GoString() string {
  return s.String()
}

type GetPropertyVersionRequestHeader struct {
}

func (s GetPropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type GetPropertyVersionPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s GetPropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionPaths) SetPropertyId(v int) *GetPropertyVersionPaths {
  s.PropertyId = &v
  return s
}

func (s *GetPropertyVersionPaths) SetVersion(v int) *GetPropertyVersionPaths {
  s.Version = &v
  return s
}

type GetPropertyVersionParameters struct {
}

func (s GetPropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionParameters) GoString() string {
  return s.String()
}

type GetPropertyVersionResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetPropertyVersionResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetPropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponse) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponse) SetCode(v string) *GetPropertyVersionResponse {
  s.Code = &v
  return s
}

func (s *GetPropertyVersionResponse) SetData(v *GetPropertyVersionResponseData) *GetPropertyVersionResponse {
  s.Data = v
  return s
}

func (s *GetPropertyVersionResponse) SetMessage(v string) *GetPropertyVersionResponse {
  s.Message = &v
  return s
}

type GetPropertyVersionResponseData struct {
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables []*GetPropertyVersionResponseDataVariables `json:"variables,omitempty" xml:"variables,omitempty" require:"true" type:"Repeated"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"A description of the version.","zh_CN":"版本描述。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty" require:"true"`
  // {"en":"Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.","zh_CN":"该项目版本是否处于冻结状态。当加速项目版本部署后即进入冻结状态，不可再更新该版本。"}
  Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins []*GetPropertyVersionResponseDataOrigins `json:"origins,omitempty" xml:"origins,omitempty" require:"true" type:"Repeated"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*GetPropertyVersionResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *GetPropertyVersionResponseDataRules `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Struct"`
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"A property version. It must be an integer value >=1.","zh_CN":"项目的版本，必须是大于0的整数。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseData) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseData) SetVariables(v []*GetPropertyVersionResponseDataVariables) *GetPropertyVersionResponseData {
  s.Variables = v
  return s
}

func (s *GetPropertyVersionResponseData) SetCreationTime(v string) *GetPropertyVersionResponseData {
  s.CreationTime = &v
  return s
}

func (s *GetPropertyVersionResponseData) SetVersionComment(v string) *GetPropertyVersionResponseData {
  s.VersionComment = &v
  return s
}

func (s *GetPropertyVersionResponseData) SetFrozen(v bool) *GetPropertyVersionResponseData {
  s.Frozen = &v
  return s
}

func (s *GetPropertyVersionResponseData) SetOrigins(v []*GetPropertyVersionResponseDataOrigins) *GetPropertyVersionResponseData {
  s.Origins = v
  return s
}

func (s *GetPropertyVersionResponseData) SetHostnames(v []*GetPropertyVersionResponseDataHostnames) *GetPropertyVersionResponseData {
  s.Hostnames = v
  return s
}

func (s *GetPropertyVersionResponseData) SetRules(v *GetPropertyVersionResponseDataRules) *GetPropertyVersionResponseData {
  s.Rules = v
  return s
}

func (s *GetPropertyVersionResponseData) SetPropertyId(v int) *GetPropertyVersionResponseData {
  s.PropertyId = &v
  return s
}

func (s *GetPropertyVersionResponseData) SetVersion(v int) *GetPropertyVersionResponseData {
  s.Version = &v
  return s
}

func (s *GetPropertyVersionResponseData) SetLastUpdateTime(v string) *GetPropertyVersionResponseData {
  s.LastUpdateTime = &v
  return s
}

type GetPropertyVersionResponseDataVariables struct     {
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"action","zh_CN":"动作"}
  Action *GetPropertyVersionResponseDataVariablesAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"variate ID","zh_CN":"变量id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataVariables) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataVariables) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataVariables) SetCondition(v string) *GetPropertyVersionResponseDataVariables {
  s.Condition = &v
  return s
}

func (s *GetPropertyVersionResponseDataVariables) SetAction(v *GetPropertyVersionResponseDataVariablesAction) *GetPropertyVersionResponseDataVariables {
  s.Action = v
  return s
}

func (s *GetPropertyVersionResponseDataVariables) SetId(v int) *GetPropertyVersionResponseDataVariables {
  s.Id = &v
  return s
}

type GetPropertyVersionResponseDataVariablesAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *GetPropertyVersionResponseDataVariablesActionOptions `json:"options,omitempty" xml:"options,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataVariablesAction) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataVariablesAction) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataVariablesAction) SetName(v string) *GetPropertyVersionResponseDataVariablesAction {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataVariablesAction) SetOptions(v *GetPropertyVersionResponseDataVariablesActionOptions) *GetPropertyVersionResponseDataVariablesAction {
  s.Options = v
  return s
}

type GetPropertyVersionResponseDataVariablesActionOptions struct {
  // {"en":"variate name","zh_CN":"变量名称"}
  VarName *string `json:"varName,omitempty" xml:"varName,omitempty" require:"true"`
  // {"en":"variate value","zh_CN":"变量值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataVariablesActionOptions) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataVariablesActionOptions) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataVariablesActionOptions) SetVarName(v string) *GetPropertyVersionResponseDataVariablesActionOptions {
  s.VarName = &v
  return s
}

func (s *GetPropertyVersionResponseDataVariablesActionOptions) SetValue(v string) *GetPropertyVersionResponseDataVariablesActionOptions {
  s.Value = &v
  return s
}

type GetPropertyVersionResponseDataOrigins struct     {
  // {"en":"Origin servers list","zh_CN":"源站服务器列表"}
  Servers []*GetPropertyVersionResponseDataOriginsServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty" require:"true"`
  // {"en":"Origin name","zh_CN":"源站名称。支持英文字符大小写（a-z，A-Z），点号(.)、下划线（_）、横杠（-）。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Origin ID","zh_CN":"源站id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty" require:"true"`
  // {"en":"Enable SNI. Whether to enable SNI. If enabled, the origin TLS handshake will carry SNI (Server Name Indication).","zh_CN":"是否启用SNI。开启后，回源TLS握手将携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty" require:"true"`
  // {"en":"SNI Server Name. After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. Default: ${http.request.host}, which follows the client request host. If left empty, the default value will be used.","zh_CN":"SNI服务器。启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。默认值：${http.request.host}，即跟随客户端请求host。空值：置空则与默认值保持一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty" require:"true"`
  // {"defaultValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3","en":"TLS Versions. List of origin TLS protocol versions, supporting configuring multiple versions.","zh_CN":"TLS版本。回源TLS协议版本列表，支持配置多个版本。","exampleValue":"SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" require:"true" type:"Repeated"`
  // {"en":"Verify Origin Certificate. Whether to enable origin certificate verification. If enabled, the back to origin node will verify the certificate offer by origin.","zh_CN":"是否开启节点验证源站证书。开启后，节点访问源站将会验证源站发过来的证书。","exampleValue":"true,false"}
  ProxySSLVerifyEnabled *bool `json:"proxySSLVerifyEnabled,omitempty" xml:"proxySSLVerifyEnabled,omitempty" require:"true"`
  // {"en":"Trust CA Certificate ID. Trusted CA Certificate for origin certificate verification.","zh_CN":"信任CA证书ID。配置节点验证源站时使用的信任证书。"}
  ProxySSLTrustedCertificate *string `json:"proxySSLTrustedCertificate,omitempty" xml:"proxySSLTrustedCertificate,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataOrigins) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataOrigins) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataOrigins) SetServers(v []*GetPropertyVersionResponseDataOriginsServers) *GetPropertyVersionResponseDataOrigins {
  s.Servers = v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetHttpPort(v int) *GetPropertyVersionResponseDataOrigins {
  s.HttpPort = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetName(v string) *GetPropertyVersionResponseDataOrigins {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetId(v int) *GetPropertyVersionResponseDataOrigins {
  s.Id = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetHttpsPort(v int) *GetPropertyVersionResponseDataOrigins {
  s.HttpsPort = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetProxySSLSNIEnabled(v bool) *GetPropertyVersionResponseDataOrigins {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetProxySSLSNIServer(v string) *GetPropertyVersionResponseDataOrigins {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetProxySSLVersion(v []*string) *GetPropertyVersionResponseDataOrigins {
  s.ProxySSLVersion = v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetProxySSLVerifyEnabled(v bool) *GetPropertyVersionResponseDataOrigins {
  s.ProxySSLVerifyEnabled = &v
  return s
}

func (s *GetPropertyVersionResponseDataOrigins) SetProxySSLTrustedCertificate(v string) *GetPropertyVersionResponseDataOrigins {
  s.ProxySSLTrustedCertificate = &v
  return s
}

type GetPropertyVersionResponseDataOriginsServers struct     {
  // {"en":"Origin server. Supports upper and lower case of English characters(a-z, A-Z), and dots(.) Underline(_), horizontal bar(-).","zh_CN":"源站服务器"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataOriginsServers) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataOriginsServers) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataOriginsServers) SetServer(v string) *GetPropertyVersionResponseDataOriginsServers {
  s.Server = &v
  return s
}

func (s *GetPropertyVersionResponseDataOriginsServers) SetWeight(v int) *GetPropertyVersionResponseDataOriginsServers {
  s.Weight = &v
  return s
}

func (s *GetPropertyVersionResponseDataOriginsServers) SetPriority(v int) *GetPropertyVersionResponseDataOriginsServers {
  s.Priority = &v
  return s
}

type GetPropertyVersionResponseDataHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *GetPropertyVersionResponseDataHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" require:"true" type:"Struct"`
  // {"en":"hostname","zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"hostname association certificate configuration","zh_CN":"关联证书配置"}
  Certificates []*GetPropertyVersionResponseDataHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
  // {"en":"icp","zh_CN":"域名备案号"}
  Icp *string `json:"icp,omitempty" xml:"icp,omitempty" require:"true"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *GetPropertyVersionResponseDataHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataHostnames) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataHostnames) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataHostnames) SetDefaultOrigin(v *GetPropertyVersionResponseDataHostnamesDefaultOrigin) *GetPropertyVersionResponseDataHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *GetPropertyVersionResponseDataHostnames) SetHostname(v string) *GetPropertyVersionResponseDataHostnames {
  s.Hostname = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnames) SetCertificates(v []*GetPropertyVersionResponseDataHostnamesCertificates) *GetPropertyVersionResponseDataHostnames {
  s.Certificates = v
  return s
}

func (s *GetPropertyVersionResponseDataHostnames) SetIcp(v string) *GetPropertyVersionResponseDataHostnames {
  s.Icp = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnames) SetEdgeHostname(v *GetPropertyVersionResponseDataHostnamesEdgeHostname) *GetPropertyVersionResponseDataHostnames {
  s.EdgeHostname = v
  return s
}

type GetPropertyVersionResponseDataHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" require:"true"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty" require:"true"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty" require:"true"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty" require:"true"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty" require:"true"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyVersionResponseDataHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetServers(v []*string) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetIpVersion(v string) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetHttpPort(v int) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetHost(v string) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetHttpsPort(v int) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetScheme(v string) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *GetPropertyVersionResponseDataHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type GetPropertyVersionResponseDataHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataHostnamesCertificates) SetCertificateId(v int64) *GetPropertyVersionResponseDataHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesCertificates) SetCertificateUsage(v string) *GetPropertyVersionResponseDataHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type GetPropertyVersionResponseDataHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"dns service status.","zh_CN":"DNS服务状态。备注：inactive：挂起，active：生效","exampleValue":"inactive, active"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *GetPropertyVersionResponseDataHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesEdgeHostname) SetDnsServiceStatus(v string) *GetPropertyVersionResponseDataHostnamesEdgeHostname {
  s.DnsServiceStatus = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *GetPropertyVersionResponseDataHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *GetPropertyVersionResponseDataHostnamesEdgeHostname) SetEdgeHostname(v string) *GetPropertyVersionResponseDataHostnamesEdgeHostname {
  s.EdgeHostname = &v
  return s
}

type GetPropertyVersionResponseDataRules struct {
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ResponsePhase []*GetPropertyVersionResponseDataRulesResponsePhase `json:"responsePhase,omitempty" xml:"responsePhase,omitempty" require:"true" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  OriginPhase []*GetPropertyVersionResponseDataRulesOriginPhase `json:"originPhase,omitempty" xml:"originPhase,omitempty" require:"true" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  CachePhase []*GetPropertyVersionResponseDataRulesCachePhase `json:"cachePhase,omitempty" xml:"cachePhase,omitempty" require:"true" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ConnectPhase []*GetPropertyVersionResponseDataRulesConnectPhase `json:"connectPhase,omitempty" xml:"connectPhase,omitempty" require:"true" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ReqeustPhase []*GetPropertyVersionResponseDataRulesReqeustPhase `json:"reqeustPhase,omitempty" xml:"reqeustPhase,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyVersionResponseDataRules) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRules) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRules) SetResponsePhase(v []*GetPropertyVersionResponseDataRulesResponsePhase) *GetPropertyVersionResponseDataRules {
  s.ResponsePhase = v
  return s
}

func (s *GetPropertyVersionResponseDataRules) SetOriginPhase(v []*GetPropertyVersionResponseDataRulesOriginPhase) *GetPropertyVersionResponseDataRules {
  s.OriginPhase = v
  return s
}

func (s *GetPropertyVersionResponseDataRules) SetCachePhase(v []*GetPropertyVersionResponseDataRulesCachePhase) *GetPropertyVersionResponseDataRules {
  s.CachePhase = v
  return s
}

func (s *GetPropertyVersionResponseDataRules) SetConnectPhase(v []*GetPropertyVersionResponseDataRulesConnectPhase) *GetPropertyVersionResponseDataRules {
  s.ConnectPhase = v
  return s
}

func (s *GetPropertyVersionResponseDataRules) SetReqeustPhase(v []*GetPropertyVersionResponseDataRulesReqeustPhase) *GetPropertyVersionResponseDataRules {
  s.ReqeustPhase = v
  return s
}

type GetPropertyVersionResponseDataRulesResponsePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"action","zh_CN":"动作"}
  Action *GetPropertyVersionResponseDataRulesResponsePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataRulesResponsePhase) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesResponsePhase) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetName(v string) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetCondition(v string) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Condition = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetAction(v *GetPropertyVersionResponseDataRulesResponsePhaseAction) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Action = v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetDescription(v string) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Description = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetId(v int) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Id = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetPriority(v int) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Priority = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhase) SetEnabled(v string) *GetPropertyVersionResponseDataRulesResponsePhase {
  s.Enabled = &v
  return s
}

type GetPropertyVersionResponseDataRulesResponsePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *GetPropertyVersionResponseDataRulesResponsePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataRulesResponsePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesResponsePhaseAction) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesResponsePhaseAction) SetName(v string) *GetPropertyVersionResponseDataRulesResponsePhaseAction {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesResponsePhaseAction) SetOptions(v *GetPropertyVersionResponseDataRulesResponsePhaseActionOptions) *GetPropertyVersionResponseDataRulesResponsePhaseAction {
  s.Options = v
  return s
}

type GetPropertyVersionResponseDataRulesResponsePhaseActionOptions struct {
}

func (s GetPropertyVersionResponseDataRulesResponsePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesResponsePhaseActionOptions) GoString() string {
  return s.String()
}

type GetPropertyVersionResponseDataRulesOriginPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"action","zh_CN":"动作"}
  Action *GetPropertyVersionResponseDataRulesOriginPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataRulesOriginPhase) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesOriginPhase) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetName(v string) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetCondition(v string) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Condition = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetAction(v *GetPropertyVersionResponseDataRulesOriginPhaseAction) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Action = v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetDescription(v string) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Description = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetId(v int) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Id = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetPriority(v int) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Priority = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhase) SetEnabled(v string) *GetPropertyVersionResponseDataRulesOriginPhase {
  s.Enabled = &v
  return s
}

type GetPropertyVersionResponseDataRulesOriginPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *GetPropertyVersionResponseDataRulesOriginPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataRulesOriginPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesOriginPhaseAction) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesOriginPhaseAction) SetName(v string) *GetPropertyVersionResponseDataRulesOriginPhaseAction {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesOriginPhaseAction) SetOptions(v *GetPropertyVersionResponseDataRulesOriginPhaseActionOptions) *GetPropertyVersionResponseDataRulesOriginPhaseAction {
  s.Options = v
  return s
}

type GetPropertyVersionResponseDataRulesOriginPhaseActionOptions struct {
}

func (s GetPropertyVersionResponseDataRulesOriginPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesOriginPhaseActionOptions) GoString() string {
  return s.String()
}

type GetPropertyVersionResponseDataRulesCachePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"action","zh_CN":"动作"}
  Action *GetPropertyVersionResponseDataRulesCachePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataRulesCachePhase) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesCachePhase) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetName(v string) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetCondition(v string) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Condition = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetAction(v *GetPropertyVersionResponseDataRulesCachePhaseAction) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Action = v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetDescription(v string) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Description = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetId(v int) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Id = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetPriority(v int) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Priority = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhase) SetEnabled(v string) *GetPropertyVersionResponseDataRulesCachePhase {
  s.Enabled = &v
  return s
}

type GetPropertyVersionResponseDataRulesCachePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *GetPropertyVersionResponseDataRulesCachePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataRulesCachePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesCachePhaseAction) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesCachePhaseAction) SetName(v string) *GetPropertyVersionResponseDataRulesCachePhaseAction {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesCachePhaseAction) SetOptions(v *GetPropertyVersionResponseDataRulesCachePhaseActionOptions) *GetPropertyVersionResponseDataRulesCachePhaseAction {
  s.Options = v
  return s
}

type GetPropertyVersionResponseDataRulesCachePhaseActionOptions struct {
}

func (s GetPropertyVersionResponseDataRulesCachePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesCachePhaseActionOptions) GoString() string {
  return s.String()
}

type GetPropertyVersionResponseDataRulesConnectPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"action","zh_CN":"动作"}
  Action *GetPropertyVersionResponseDataRulesConnectPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataRulesConnectPhase) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesConnectPhase) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetName(v string) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetCondition(v string) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Condition = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetAction(v *GetPropertyVersionResponseDataRulesConnectPhaseAction) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Action = v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetDescription(v string) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Description = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetId(v int) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Id = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetPriority(v int) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Priority = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhase) SetEnabled(v string) *GetPropertyVersionResponseDataRulesConnectPhase {
  s.Enabled = &v
  return s
}

type GetPropertyVersionResponseDataRulesConnectPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *GetPropertyVersionResponseDataRulesConnectPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataRulesConnectPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesConnectPhaseAction) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesConnectPhaseAction) SetName(v string) *GetPropertyVersionResponseDataRulesConnectPhaseAction {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesConnectPhaseAction) SetOptions(v *GetPropertyVersionResponseDataRulesConnectPhaseActionOptions) *GetPropertyVersionResponseDataRulesConnectPhaseAction {
  s.Options = v
  return s
}

type GetPropertyVersionResponseDataRulesConnectPhaseActionOptions struct {
}

func (s GetPropertyVersionResponseDataRulesConnectPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesConnectPhaseActionOptions) GoString() string {
  return s.String()
}

type GetPropertyVersionResponseDataRulesReqeustPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"action","zh_CN":"动作"}
  Action *GetPropertyVersionResponseDataRulesReqeustPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPropertyVersionResponseDataRulesReqeustPhase) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesReqeustPhase) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetName(v string) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetCondition(v string) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Condition = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetAction(v *GetPropertyVersionResponseDataRulesReqeustPhaseAction) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Action = v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetDescription(v string) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Description = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetId(v int) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Id = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetPriority(v int) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Priority = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhase) SetEnabled(v string) *GetPropertyVersionResponseDataRulesReqeustPhase {
  s.Enabled = &v
  return s
}

type GetPropertyVersionResponseDataRulesReqeustPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *GetPropertyVersionResponseDataRulesReqeustPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyVersionResponseDataRulesReqeustPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesReqeustPhaseAction) GoString() string {
  return s.String()
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhaseAction) SetName(v string) *GetPropertyVersionResponseDataRulesReqeustPhaseAction {
  s.Name = &v
  return s
}

func (s *GetPropertyVersionResponseDataRulesReqeustPhaseAction) SetOptions(v *GetPropertyVersionResponseDataRulesReqeustPhaseActionOptions) *GetPropertyVersionResponseDataRulesReqeustPhaseAction {
  s.Options = v
  return s
}

type GetPropertyVersionResponseDataRulesReqeustPhaseActionOptions struct {
}

func (s GetPropertyVersionResponseDataRulesReqeustPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseDataRulesReqeustPhaseActionOptions) GoString() string {
  return s.String()
}

type GetPropertyVersionResponseHeader struct {
}

func (s GetPropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyVersionResponseHeader) GoString() string {
  return s.String()
}




type QueryDeploymentForTerraformRequest struct {
}

func (s QueryDeploymentForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformRequest) GoString() string {
  return s.String()
}

type QueryDeploymentForTerraformRequestHeader struct {
}

func (s QueryDeploymentForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryDeploymentForTerraformPaths struct {
  // {"en":"ID of the deployment task","zh_CN":"部署任务ID"}
  DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
}

func (s QueryDeploymentForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformPaths) GoString() string {
  return s.String()
}

func (s *QueryDeploymentForTerraformPaths) SetDeploymentId(v int) *QueryDeploymentForTerraformPaths {
  s.DeploymentId = &v
  return s
}

type QueryDeploymentForTerraformParameters struct {
}

func (s QueryDeploymentForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformParameters) GoString() string {
  return s.String()
}

type QueryDeploymentForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryDeploymentForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryDeploymentForTerraformResponse) SetCode(v string) *QueryDeploymentForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryDeploymentForTerraformResponse) SetMessage(v string) *QueryDeploymentForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryDeploymentForTerraformResponse) SetData(v *QueryDeploymentForTerraformResponseData) *QueryDeploymentForTerraformResponse {
  s.Data = v
  return s
}

type QueryDeploymentForTerraformResponseData struct {
  // {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
  DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
  // {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
  DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty" require:"true"`
  // {"en":"Status of Deployment. Enum:PENDING,IN_PROCESS,SUCCESS,FAIL","zh_CN":"任务状态。取值范围:PENDING,IN_PROCESS,SUCCESS,FAIL"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"项目的部署环境。取值范围: staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en":"An RFC3339 date indicates when the task was last updated.","zh_CN":"RFC 3339格式的日期，表示任务的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties to the staging or production environments.","zh_CN":"指部署任务所要执行的操作，可以包括项目的部署或卸载操作。"}
  Actions []*QueryDeploymentForTerraformResponseDataActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDeploymentForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryDeploymentForTerraformResponseData) SetDeploymentId(v int) *QueryDeploymentForTerraformResponseData {
  s.DeploymentId = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetDeploymentName(v string) *QueryDeploymentForTerraformResponseData {
  s.DeploymentName = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetStatus(v string) *QueryDeploymentForTerraformResponseData {
  s.Status = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetTarget(v string) *QueryDeploymentForTerraformResponseData {
  s.Target = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetSubmissionTime(v string) *QueryDeploymentForTerraformResponseData {
  s.SubmissionTime = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetLastUpdateTime(v string) *QueryDeploymentForTerraformResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetFinishTime(v string) *QueryDeploymentForTerraformResponseData {
  s.FinishTime = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseData) SetActions(v []*QueryDeploymentForTerraformResponseDataActions) *QueryDeploymentForTerraformResponseData {
  s.Actions = v
  return s
}

type QueryDeploymentForTerraformResponseDataActions struct     {
  // {"en":"Describe an action to take. You can deploy a property, remove a property. Enum: deploy_property,remove_property","zh_CN":"指定操作类型，包括部署加速项目、卸载项目。取值范围: deploy_property,remove_property"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的项目ID。"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Indicates the version of the property to deploy or remove.","zh_CN":"指定要部署或卸载项目的版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s QueryDeploymentForTerraformResponseDataActions) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponseDataActions) GoString() string {
  return s.String()
}

func (s *QueryDeploymentForTerraformResponseDataActions) SetAction(v string) *QueryDeploymentForTerraformResponseDataActions {
  s.Action = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseDataActions) SetPropertyId(v int) *QueryDeploymentForTerraformResponseDataActions {
  s.PropertyId = &v
  return s
}

func (s *QueryDeploymentForTerraformResponseDataActions) SetVersion(v int) *QueryDeploymentForTerraformResponseDataActions {
  s.Version = &v
  return s
}

type QueryDeploymentForTerraformResponseHeader struct {
}

func (s QueryDeploymentForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponseHeader) GoString() string {
  return s.String()
}




type CreatePropertyRequest struct {
  // {"en":"Product Service Type related to your contract.","zh_CN":"产品服务类型。请根据您的合同产品服务类型填写。","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables []*CreatePropertyRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
  // {"en":"Clone param","zh_CN":"克隆版本参数"}
  CloneParam *CreatePropertyRequestCloneParam `json:"cloneParam,omitempty" xml:"cloneParam,omitempty" type:"Struct"`
  // {"en":"The name of the property. The length must not exceed 128 characters.","zh_CN":"项目的名称。长度不超过128个字符。"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
  // {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins []*CreatePropertyRequestOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*CreatePropertyRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *CreatePropertyRequestRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Struct"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}

func (s CreatePropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequest) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequest) SetServiceType(v string) *CreatePropertyRequest {
  s.ServiceType = &v
  return s
}

func (s *CreatePropertyRequest) SetPropertyComment(v string) *CreatePropertyRequest {
  s.PropertyComment = &v
  return s
}

func (s *CreatePropertyRequest) SetVariables(v []*CreatePropertyRequestVariables) *CreatePropertyRequest {
  s.Variables = v
  return s
}

func (s *CreatePropertyRequest) SetCloneParam(v *CreatePropertyRequestCloneParam) *CreatePropertyRequest {
  s.CloneParam = v
  return s
}

func (s *CreatePropertyRequest) SetPropertyName(v string) *CreatePropertyRequest {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyRequest) SetVersionComment(v string) *CreatePropertyRequest {
  s.VersionComment = &v
  return s
}

func (s *CreatePropertyRequest) SetOrigins(v []*CreatePropertyRequestOrigins) *CreatePropertyRequest {
  s.Origins = v
  return s
}

func (s *CreatePropertyRequest) SetHostnames(v []*CreatePropertyRequestHostnames) *CreatePropertyRequest {
  s.Hostnames = v
  return s
}

func (s *CreatePropertyRequest) SetRules(v *CreatePropertyRequestRules) *CreatePropertyRequest {
  s.Rules = v
  return s
}

func (s *CreatePropertyRequest) SetContractId(v string) *CreatePropertyRequest {
  s.ContractId = &v
  return s
}

func (s *CreatePropertyRequest) SetItemId(v string) *CreatePropertyRequest {
  s.ItemId = &v
  return s
}

type CreatePropertyRequestVariables struct     {
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestVariablesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en":"The name of variable","zh_CN":"变量名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreatePropertyRequestVariables) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestVariables) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestVariables) SetCondition(v string) *CreatePropertyRequestVariables {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestVariables) SetAction(v *CreatePropertyRequestVariablesAction) *CreatePropertyRequestVariables {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestVariables) SetName(v string) *CreatePropertyRequestVariables {
  s.Name = &v
  return s
}

type CreatePropertyRequestVariablesAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestVariablesActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestVariablesAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestVariablesAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestVariablesAction) SetName(v string) *CreatePropertyRequestVariablesAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestVariablesAction) SetOptions(v *CreatePropertyRequestVariablesActionOptions) *CreatePropertyRequestVariablesAction {
  s.Options = v
  return s
}

type CreatePropertyRequestVariablesActionOptions struct {
  // {"en":"variate name","zh_CN":"变量名称"}
  VarName *string `json:"varName,omitempty" xml:"varName,omitempty"`
  // {"en":"variate value","zh_CN":"变量值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreatePropertyRequestVariablesActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestVariablesActionOptions) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestVariablesActionOptions) SetVarName(v string) *CreatePropertyRequestVariablesActionOptions {
  s.VarName = &v
  return s
}

func (s *CreatePropertyRequestVariablesActionOptions) SetValue(v string) *CreatePropertyRequestVariablesActionOptions {
  s.Value = &v
  return s
}

type CreatePropertyRequestCloneParam struct {
  // {"en":"The property version you need to clone","zh_CN":"克隆项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"The property Id you need to clone","zh_CN":"克隆项目id"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyRequestCloneParam) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestCloneParam) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestCloneParam) SetPropertyVersion(v int) *CreatePropertyRequestCloneParam {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyRequestCloneParam) SetPropertyId(v int) *CreatePropertyRequestCloneParam {
  s.PropertyId = &v
  return s
}

type CreatePropertyRequestOrigins struct     {
  // {"en":"Origin servers list","zh_CN":"源站服务器列表"}
  Servers []*CreatePropertyRequestOriginsServers `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"Http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"Origin name. Supports upper and lower case of English characters(a-z, A-Z), and dots(.) Underline(_), horizontal bar(-).","zh_CN":"源站名称。支持英文字符大小写（a-z，A-Z），点号(.)、下划线（_）、横杠（-）。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Enable SNI. Whether to enable SNI. If enabled, the origin TLS handshake will carry SNI (Server Name Indication).","zh_CN":"是否启用SNI。开启后，回源TLS握手将携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"SNI Server Name. After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. Default: ${http.request.host}, which follows the client request host. If left empty, the default value will be used.","zh_CN":"SNI服务器。启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。默认值：${http.request.host}，即跟随客户端请求host。空值：置空则与默认值保持一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"defaultValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3","en":"TLS Versions. List of origin TLS protocol versions, supporting configuring multiple versions.","zh_CN":"TLS版本。回源TLS协议版本列表，支持配置多个版本。","exampleValue":"SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
  // {"en":"Verify Origin Certificate. Whether to enable origin certificate verification. If enabled, the back to origin node will verify the certificate offer by origin.","zh_CN":"是否开启节点验证源站证书。开启后，节点访问源站将会验证源站发过来的证书。","exampleValue":"true,false"}
  ProxySSLVerifyEnabled *bool `json:"proxySSLVerifyEnabled,omitempty" xml:"proxySSLVerifyEnabled,omitempty"`
  // {"en":"Trust CA Certificate ID. Trusted CA Certificate for origin certificate verification.","zh_CN":"信任CA证书ID。配置节点验证源站时使用的信任证书。"}
  ProxySSLTrustedCertificate *string `json:"proxySSLTrustedCertificate,omitempty" xml:"proxySSLTrustedCertificate,omitempty"`
}

func (s CreatePropertyRequestOrigins) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestOrigins) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestOrigins) SetServers(v []*CreatePropertyRequestOriginsServers) *CreatePropertyRequestOrigins {
  s.Servers = v
  return s
}

func (s *CreatePropertyRequestOrigins) SetHttpPort(v int) *CreatePropertyRequestOrigins {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetName(v string) *CreatePropertyRequestOrigins {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetHttpsPort(v int) *CreatePropertyRequestOrigins {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLSNIEnabled(v bool) *CreatePropertyRequestOrigins {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLSNIServer(v string) *CreatePropertyRequestOrigins {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLVersion(v []*string) *CreatePropertyRequestOrigins {
  s.ProxySSLVersion = v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLVerifyEnabled(v bool) *CreatePropertyRequestOrigins {
  s.ProxySSLVerifyEnabled = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLTrustedCertificate(v string) *CreatePropertyRequestOrigins {
  s.ProxySSLTrustedCertificate = &v
  return s
}

type CreatePropertyRequestOriginsServers struct     {
  // {"en":"Origin server","zh_CN":"源站服务器"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en":"priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s CreatePropertyRequestOriginsServers) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestOriginsServers) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestOriginsServers) SetServer(v string) *CreatePropertyRequestOriginsServers {
  s.Server = &v
  return s
}

func (s *CreatePropertyRequestOriginsServers) SetWeight(v int) *CreatePropertyRequestOriginsServers {
  s.Weight = &v
  return s
}

func (s *CreatePropertyRequestOriginsServers) SetPriority(v int) *CreatePropertyRequestOriginsServers {
  s.Priority = &v
  return s
}

type CreatePropertyRequestHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *CreatePropertyRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
  // {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
  Certificates []*CreatePropertyRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *CreatePropertyRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnames) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnames) SetDefaultOrigin(v *CreatePropertyRequestHostnamesDefaultOrigin) *CreatePropertyRequestHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *CreatePropertyRequestHostnames) SetHostname(v string) *CreatePropertyRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *CreatePropertyRequestHostnames) SetCertificates(v []*CreatePropertyRequestHostnamesCertificates) *CreatePropertyRequestHostnames {
  s.Certificates = v
  return s
}

func (s *CreatePropertyRequestHostnames) SetEdgeHostname(v *CreatePropertyRequestHostnamesEdgeHostname) *CreatePropertyRequestHostnames {
  s.EdgeHostname = v
  return s
}

type CreatePropertyRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:\n- http: Fixedly use HTTP.\n- https: Fixedly use HTTPS.\n- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： \n- http：固定使用HTTP协议回源。\n- https：固定使用HTTPS协议回源。\n- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:\n- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.\n- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：\n- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。\n- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.\n- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3\n- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。\n- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3\n- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetScheme(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type CreatePropertyRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;\n- origin_mtls: Origin mTLS Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书；\n- origin_mtls：表示回源mTLS证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty"`
}

func (s CreatePropertyRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *CreatePropertyRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type CreatePropertyRequestHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
}

func (s CreatePropertyRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *CreatePropertyRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

func (s *CreatePropertyRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

type CreatePropertyRequestRules struct {
  // {"en":"response phase","zh_CN":"响应阶段"}
  ResponsePhase []*CreatePropertyRequestRulesResponsePhase `json:"responsePhase,omitempty" xml:"responsePhase,omitempty" type:"Repeated"`
  // {"en":"origin phase","zh_CN":"回源阶段"}
  OriginPhase []*CreatePropertyRequestRulesOriginPhase `json:"originPhase,omitempty" xml:"originPhase,omitempty" type:"Repeated"`
  // {"en":"cache phase","zh_CN":"缓存阶段"}
  CachePhase []*CreatePropertyRequestRulesCachePhase `json:"cachePhase,omitempty" xml:"cachePhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ConnectPhase []*CreatePropertyRequestRulesConnectPhase `json:"connectPhase,omitempty" xml:"connectPhase,omitempty" type:"Repeated"`
  // {"en":"request phase","zh_CN":"请求阶段"}
  ReqeustPhase []*CreatePropertyRequestRulesReqeustPhase `json:"reqeustPhase,omitempty" xml:"reqeustPhase,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequestRules) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRules) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRules) SetResponsePhase(v []*CreatePropertyRequestRulesResponsePhase) *CreatePropertyRequestRules {
  s.ResponsePhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetOriginPhase(v []*CreatePropertyRequestRulesOriginPhase) *CreatePropertyRequestRules {
  s.OriginPhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetCachePhase(v []*CreatePropertyRequestRulesCachePhase) *CreatePropertyRequestRules {
  s.CachePhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetConnectPhase(v []*CreatePropertyRequestRulesConnectPhase) *CreatePropertyRequestRules {
  s.ConnectPhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetReqeustPhase(v []*CreatePropertyRequestRulesReqeustPhase) *CreatePropertyRequestRules {
  s.ReqeustPhase = v
  return s
}

type CreatePropertyRequestRulesResponsePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesResponsePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesResponsePhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesResponsePhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesResponsePhase) SetName(v string) *CreatePropertyRequestRulesResponsePhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetCondition(v string) *CreatePropertyRequestRulesResponsePhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetAction(v *CreatePropertyRequestRulesResponsePhaseAction) *CreatePropertyRequestRulesResponsePhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetDescription(v string) *CreatePropertyRequestRulesResponsePhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetPriority(v int) *CreatePropertyRequestRulesResponsePhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetEnabled(v bool) *CreatePropertyRequestRulesResponsePhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesResponsePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesResponsePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesResponsePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesResponsePhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesResponsePhaseAction) SetName(v string) *CreatePropertyRequestRulesResponsePhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhaseAction) SetOptions(v *CreatePropertyRequestRulesResponsePhaseActionOptions) *CreatePropertyRequestRulesResponsePhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesResponsePhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesResponsePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesResponsePhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesOriginPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesOriginPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesOriginPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesOriginPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesOriginPhase) SetName(v string) *CreatePropertyRequestRulesOriginPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetCondition(v string) *CreatePropertyRequestRulesOriginPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetAction(v *CreatePropertyRequestRulesOriginPhaseAction) *CreatePropertyRequestRulesOriginPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetDescription(v string) *CreatePropertyRequestRulesOriginPhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetPriority(v int) *CreatePropertyRequestRulesOriginPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetEnabled(v bool) *CreatePropertyRequestRulesOriginPhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesOriginPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesOriginPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesOriginPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesOriginPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesOriginPhaseAction) SetName(v string) *CreatePropertyRequestRulesOriginPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhaseAction) SetOptions(v *CreatePropertyRequestRulesOriginPhaseActionOptions) *CreatePropertyRequestRulesOriginPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesOriginPhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesOriginPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesOriginPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesCachePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesCachePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesCachePhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesCachePhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesCachePhase) SetName(v string) *CreatePropertyRequestRulesCachePhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetCondition(v string) *CreatePropertyRequestRulesCachePhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetAction(v *CreatePropertyRequestRulesCachePhaseAction) *CreatePropertyRequestRulesCachePhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetDescription(v string) *CreatePropertyRequestRulesCachePhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetPriority(v int) *CreatePropertyRequestRulesCachePhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetEnabled(v bool) *CreatePropertyRequestRulesCachePhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesCachePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesCachePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesCachePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesCachePhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesCachePhaseAction) SetName(v string) *CreatePropertyRequestRulesCachePhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhaseAction) SetOptions(v *CreatePropertyRequestRulesCachePhaseActionOptions) *CreatePropertyRequestRulesCachePhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesCachePhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesCachePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesCachePhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesConnectPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesConnectPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesConnectPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesConnectPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesConnectPhase) SetName(v string) *CreatePropertyRequestRulesConnectPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetCondition(v string) *CreatePropertyRequestRulesConnectPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetAction(v *CreatePropertyRequestRulesConnectPhaseAction) *CreatePropertyRequestRulesConnectPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetDescription(v string) *CreatePropertyRequestRulesConnectPhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetPriority(v int) *CreatePropertyRequestRulesConnectPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetEnabled(v bool) *CreatePropertyRequestRulesConnectPhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesConnectPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesConnectPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesConnectPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesConnectPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesConnectPhaseAction) SetName(v string) *CreatePropertyRequestRulesConnectPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhaseAction) SetOptions(v *CreatePropertyRequestRulesConnectPhaseActionOptions) *CreatePropertyRequestRulesConnectPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesConnectPhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesConnectPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesConnectPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesReqeustPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesReqeustPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesReqeustPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesReqeustPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetName(v string) *CreatePropertyRequestRulesReqeustPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetCondition(v string) *CreatePropertyRequestRulesReqeustPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetAction(v *CreatePropertyRequestRulesReqeustPhaseAction) *CreatePropertyRequestRulesReqeustPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetDescription(v string) *CreatePropertyRequestRulesReqeustPhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetPriority(v int) *CreatePropertyRequestRulesReqeustPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetEnabled(v bool) *CreatePropertyRequestRulesReqeustPhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesReqeustPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesReqeustPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesReqeustPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesReqeustPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesReqeustPhaseAction) SetName(v string) *CreatePropertyRequestRulesReqeustPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhaseAction) SetOptions(v *CreatePropertyRequestRulesReqeustPhaseActionOptions) *CreatePropertyRequestRulesReqeustPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesReqeustPhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesReqeustPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesReqeustPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestHeader struct {
}

func (s CreatePropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHeader) GoString() string {
  return s.String()
}

type CreatePropertyPaths struct {
}

func (s CreatePropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyPaths) GoString() string {
  return s.String()
}

type CreatePropertyParameters struct {
}

func (s CreatePropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyParameters) GoString() string {
  return s.String()
}

type CreatePropertyResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreatePropertyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreatePropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyResponse) GoString() string {
  return s.String()
}

func (s *CreatePropertyResponse) SetCode(v string) *CreatePropertyResponse {
  s.Code = &v
  return s
}

func (s *CreatePropertyResponse) SetData(v *CreatePropertyResponseData) *CreatePropertyResponse {
  s.Data = v
  return s
}

func (s *CreatePropertyResponse) SetMessage(v string) *CreatePropertyResponse {
  s.Message = &v
  return s
}

type CreatePropertyResponseData struct {
  // {"en":"Property version.","zh_CN":"项目的版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"Property name.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyResponseData) GoString() string {
  return s.String()
}

func (s *CreatePropertyResponseData) SetPropertyVersion(v int) *CreatePropertyResponseData {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyResponseData) SetPropertyName(v string) *CreatePropertyResponseData {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyResponseData) SetPropertyId(v int64) *CreatePropertyResponseData {
  s.PropertyId = &v
  return s
}

type CreatePropertyResponseHeader struct {
}

func (s CreatePropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyResponseHeader) GoString() string {
  return s.String()
}




type CreateDeploymentTaskForTerraformRequest struct {
  // {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
  DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
  // {"en":"Indicates whether to deploy to staging or production.","zh_CN":"指定部署任务的目标环境，即演练或生产环境。","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties to the staging or production environments.","zh_CN":"指部署任务所要执行的操作，可以包括项目的部署或卸载操作。"}
  Actions []*CreateDeploymentTaskForTerraformRequestActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentTaskForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformRequest) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskForTerraformRequest) SetDeploymentName(v string) *CreateDeploymentTaskForTerraformRequest {
  s.DeploymentName = &v
  return s
}

func (s *CreateDeploymentTaskForTerraformRequest) SetTarget(v string) *CreateDeploymentTaskForTerraformRequest {
  s.Target = &v
  return s
}

func (s *CreateDeploymentTaskForTerraformRequest) SetActions(v []*CreateDeploymentTaskForTerraformRequestActions) *CreateDeploymentTaskForTerraformRequest {
  s.Actions = v
  return s
}

type CreateDeploymentTaskForTerraformRequestActions struct     {
  // {"en":"Describe an action to take. You can deploy a property, remove a property.","zh_CN":"指定操作类型，包括部署项目、卸载项目。","exampleValue":"deploy_property,remove_property"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的项目ID。"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Indicates the version of the property to deploy or remove.","zh_CN":"指定要部署或卸载项目的版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s CreateDeploymentTaskForTerraformRequestActions) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformRequestActions) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskForTerraformRequestActions) SetAction(v string) *CreateDeploymentTaskForTerraformRequestActions {
  s.Action = &v
  return s
}

func (s *CreateDeploymentTaskForTerraformRequestActions) SetPropertyId(v int) *CreateDeploymentTaskForTerraformRequestActions {
  s.PropertyId = &v
  return s
}

func (s *CreateDeploymentTaskForTerraformRequestActions) SetVersion(v int) *CreateDeploymentTaskForTerraformRequestActions {
  s.Version = &v
  return s
}

type CreateDeploymentTaskForTerraformRequestHeader struct {
}

func (s CreateDeploymentTaskForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformRequestHeader) GoString() string {
  return s.String()
}

type CreateDeploymentTaskForTerraformPaths struct {
}

func (s CreateDeploymentTaskForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformPaths) GoString() string {
  return s.String()
}

type CreateDeploymentTaskForTerraformParameters struct {
}

func (s CreateDeploymentTaskForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformParameters) GoString() string {
  return s.String()
}

type CreateDeploymentTaskForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreateDeploymentTaskForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateDeploymentTaskForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformResponse) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskForTerraformResponse) SetCode(v string) *CreateDeploymentTaskForTerraformResponse {
  s.Code = &v
  return s
}

func (s *CreateDeploymentTaskForTerraformResponse) SetMessage(v string) *CreateDeploymentTaskForTerraformResponse {
  s.Message = &v
  return s
}

func (s *CreateDeploymentTaskForTerraformResponse) SetData(v *CreateDeploymentTaskForTerraformResponseData) *CreateDeploymentTaskForTerraformResponse {
  s.Data = v
  return s
}

type CreateDeploymentTaskForTerraformResponseData struct {
  // {"en":"ID of the deployment task.","zh_CN":"部署任务标识"}
  DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
}

func (s CreateDeploymentTaskForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskForTerraformResponseData) SetDeploymentId(v int) *CreateDeploymentTaskForTerraformResponseData {
  s.DeploymentId = &v
  return s
}

type CreateDeploymentTaskForTerraformResponseHeader struct {
}

func (s CreateDeploymentTaskForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformResponseHeader) GoString() string {
  return s.String()
}




type CreatePropertyVersionRequest struct {
  // {"en":"Property version description. The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*CreatePropertyVersionRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins []*CreatePropertyVersionRequestOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables []*CreatePropertyVersionRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *CreatePropertyVersionRequestRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Struct"`
  // {"en":"Clone param","zh_CN":"克隆版本参数"}
  CloneParam *CreatePropertyVersionRequestCloneParam `json:"cloneParam,omitempty" xml:"cloneParam,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequest) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequest) SetVersionComment(v string) *CreatePropertyVersionRequest {
  s.VersionComment = &v
  return s
}

func (s *CreatePropertyVersionRequest) SetHostnames(v []*CreatePropertyVersionRequestHostnames) *CreatePropertyVersionRequest {
  s.Hostnames = v
  return s
}

func (s *CreatePropertyVersionRequest) SetOrigins(v []*CreatePropertyVersionRequestOrigins) *CreatePropertyVersionRequest {
  s.Origins = v
  return s
}

func (s *CreatePropertyVersionRequest) SetVariables(v []*CreatePropertyVersionRequestVariables) *CreatePropertyVersionRequest {
  s.Variables = v
  return s
}

func (s *CreatePropertyVersionRequest) SetRules(v *CreatePropertyVersionRequestRules) *CreatePropertyVersionRequest {
  s.Rules = v
  return s
}

func (s *CreatePropertyVersionRequest) SetCloneParam(v *CreatePropertyVersionRequestCloneParam) *CreatePropertyVersionRequest {
  s.CloneParam = v
  return s
}

type CreatePropertyVersionRequestHostnames struct     {
  // {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *CreatePropertyVersionRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
  // {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
  Certificates []*CreatePropertyVersionRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *CreatePropertyVersionRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestHostnames) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestHostnames) SetHostname(v string) *CreatePropertyVersionRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnames) SetDefaultOrigin(v *CreatePropertyVersionRequestHostnamesDefaultOrigin) *CreatePropertyVersionRequestHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *CreatePropertyVersionRequestHostnames) SetCertificates(v []*CreatePropertyVersionRequestHostnamesCertificates) *CreatePropertyVersionRequestHostnames {
  s.Certificates = v
  return s
}

func (s *CreatePropertyVersionRequestHostnames) SetEdgeHostname(v *CreatePropertyVersionRequestHostnamesEdgeHostname) *CreatePropertyVersionRequestHostnames {
  s.EdgeHostname = v
  return s
}

type CreatePropertyVersionRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s CreatePropertyVersionRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetScheme(v string) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *CreatePropertyVersionRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type CreatePropertyVersionRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;\n- origin_mtls: Origin mTLS Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书；\n- origin_mtls：表示回源mTLS证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty"`
}

func (s CreatePropertyVersionRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyVersionRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyVersionRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type CreatePropertyVersionRequestHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s CreatePropertyVersionRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyVersionRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyVersionRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *CreatePropertyVersionRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyVersionRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

type CreatePropertyVersionRequestOrigins struct     {
  // {"en":"Origin name. Supports upper and lower case of English characters(a-z, A-Z), and dots(.) Underline(_), horizontal bar(-).","zh_CN":"源站名称。支持英文字符大小写（a-z，A-Z），点号(.)、下划线（_）、横杠（-）。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Origin servers list","zh_CN":"源站服务器列表"}
  Servers []*CreatePropertyVersionRequestOriginsServers `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"Http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"Https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Enable SNI. Whether to enable SNI. If enabled, the origin TLS handshake will carry SNI (Server Name Indication).","zh_CN":"是否启用SNI。开启后，回源TLS握手将携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"SNI Server Name. After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. Default: ${http.request.host}, which follows the client request host. If left empty, the default value will be used.","zh_CN":"SNI服务器。启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。默认值：${http.request.host}，即跟随客户端请求host。空值：置空则与默认值保持一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"TLS Versions. List of origin TLS protocol versions, supporting configuring multiple versions.","zh_CN":"TLS版本。回源TLS协议版本列表，支持配置多个版本。"}
  ProxySSLVersion *string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty"`
  // {"en":"Verify Origin Certificate. Whether to enable origin certificate verification. If enabled, the back to origin node will verify the certificate offer by origin.","zh_CN":"是否开启节点验证源站证书。开启后，节点访问源站将会验证源站发过来的证书。","exampleValue":"true,false"}
  ProxySSLVerifyEnabled *bool `json:"proxySSLVerifyEnabled,omitempty" xml:"proxySSLVerifyEnabled,omitempty"`
  // {"en":"Trust CA Certificate ID. Trusted CA Certificate for origin certificate verification.","zh_CN":"信任CA证书ID。配置节点验证源站时使用的信任证书。"}
  ProxySSLTrustedCertificate *string `json:"proxySSLTrustedCertificate,omitempty" xml:"proxySSLTrustedCertificate,omitempty"`
}

func (s CreatePropertyVersionRequestOrigins) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestOrigins) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestOrigins) SetName(v string) *CreatePropertyVersionRequestOrigins {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetServers(v []*CreatePropertyVersionRequestOriginsServers) *CreatePropertyVersionRequestOrigins {
  s.Servers = v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetHttpPort(v int) *CreatePropertyVersionRequestOrigins {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetHttpsPort(v int) *CreatePropertyVersionRequestOrigins {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetProxySSLSNIEnabled(v bool) *CreatePropertyVersionRequestOrigins {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetProxySSLSNIServer(v string) *CreatePropertyVersionRequestOrigins {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetProxySSLVersion(v string) *CreatePropertyVersionRequestOrigins {
  s.ProxySSLVersion = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetProxySSLVerifyEnabled(v bool) *CreatePropertyVersionRequestOrigins {
  s.ProxySSLVerifyEnabled = &v
  return s
}

func (s *CreatePropertyVersionRequestOrigins) SetProxySSLTrustedCertificate(v string) *CreatePropertyVersionRequestOrigins {
  s.ProxySSLTrustedCertificate = &v
  return s
}

type CreatePropertyVersionRequestOriginsServers struct     {
  // {"en":"Origin server","zh_CN":"源站服务器"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en":"priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s CreatePropertyVersionRequestOriginsServers) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestOriginsServers) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestOriginsServers) SetServer(v string) *CreatePropertyVersionRequestOriginsServers {
  s.Server = &v
  return s
}

func (s *CreatePropertyVersionRequestOriginsServers) SetWeight(v int) *CreatePropertyVersionRequestOriginsServers {
  s.Weight = &v
  return s
}

func (s *CreatePropertyVersionRequestOriginsServers) SetPriority(v int) *CreatePropertyVersionRequestOriginsServers {
  s.Priority = &v
  return s
}

type CreatePropertyVersionRequestVariables struct     {
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyVersionRequestVariablesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestVariables) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestVariables) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestVariables) SetCondition(v string) *CreatePropertyVersionRequestVariables {
  s.Condition = &v
  return s
}

func (s *CreatePropertyVersionRequestVariables) SetAction(v *CreatePropertyVersionRequestVariablesAction) *CreatePropertyVersionRequestVariables {
  s.Action = v
  return s
}

type CreatePropertyVersionRequestVariablesAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyVersionRequestVariablesActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestVariablesAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestVariablesAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestVariablesAction) SetName(v string) *CreatePropertyVersionRequestVariablesAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestVariablesAction) SetOptions(v *CreatePropertyVersionRequestVariablesActionOptions) *CreatePropertyVersionRequestVariablesAction {
  s.Options = v
  return s
}

type CreatePropertyVersionRequestVariablesActionOptions struct {
  // {"en":"variate name","zh_CN":"变量名称"}
  VarName *string `json:"varName,omitempty" xml:"varName,omitempty"`
  // {"en":"variate value","zh_CN":"变量值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreatePropertyVersionRequestVariablesActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestVariablesActionOptions) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestVariablesActionOptions) SetVarName(v string) *CreatePropertyVersionRequestVariablesActionOptions {
  s.VarName = &v
  return s
}

func (s *CreatePropertyVersionRequestVariablesActionOptions) SetValue(v string) *CreatePropertyVersionRequestVariablesActionOptions {
  s.Value = &v
  return s
}

type CreatePropertyVersionRequestRules struct {
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ConnectPhase []*CreatePropertyVersionRequestRulesConnectPhase `json:"connectPhase,omitempty" xml:"connectPhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ReqeustPhase []*CreatePropertyVersionRequestRulesReqeustPhase `json:"reqeustPhase,omitempty" xml:"reqeustPhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  CachePhase []*CreatePropertyVersionRequestRulesCachePhase `json:"cachePhase,omitempty" xml:"cachePhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  OriginPhase []*CreatePropertyVersionRequestRulesOriginPhase `json:"originPhase,omitempty" xml:"originPhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ResponsePhase []*CreatePropertyVersionRequestRulesResponsePhase `json:"responsePhase,omitempty" xml:"responsePhase,omitempty" type:"Repeated"`
}

func (s CreatePropertyVersionRequestRules) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRules) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRules) SetConnectPhase(v []*CreatePropertyVersionRequestRulesConnectPhase) *CreatePropertyVersionRequestRules {
  s.ConnectPhase = v
  return s
}

func (s *CreatePropertyVersionRequestRules) SetReqeustPhase(v []*CreatePropertyVersionRequestRulesReqeustPhase) *CreatePropertyVersionRequestRules {
  s.ReqeustPhase = v
  return s
}

func (s *CreatePropertyVersionRequestRules) SetCachePhase(v []*CreatePropertyVersionRequestRulesCachePhase) *CreatePropertyVersionRequestRules {
  s.CachePhase = v
  return s
}

func (s *CreatePropertyVersionRequestRules) SetOriginPhase(v []*CreatePropertyVersionRequestRulesOriginPhase) *CreatePropertyVersionRequestRules {
  s.OriginPhase = v
  return s
}

func (s *CreatePropertyVersionRequestRules) SetResponsePhase(v []*CreatePropertyVersionRequestRulesResponsePhase) *CreatePropertyVersionRequestRules {
  s.ResponsePhase = v
  return s
}

type CreatePropertyVersionRequestRulesConnectPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyVersionRequestRulesConnectPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreatePropertyVersionRequestRulesConnectPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesConnectPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesConnectPhase) SetName(v string) *CreatePropertyVersionRequestRulesConnectPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesConnectPhase) SetCondition(v string) *CreatePropertyVersionRequestRulesConnectPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesConnectPhase) SetAction(v *CreatePropertyVersionRequestRulesConnectPhaseAction) *CreatePropertyVersionRequestRulesConnectPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyVersionRequestRulesConnectPhase) SetPriority(v int) *CreatePropertyVersionRequestRulesConnectPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesConnectPhase) SetEnabled(v bool) *CreatePropertyVersionRequestRulesConnectPhase {
  s.Enabled = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesConnectPhase) SetDescription(v string) *CreatePropertyVersionRequestRulesConnectPhase {
  s.Description = &v
  return s
}

type CreatePropertyVersionRequestRulesConnectPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyVersionRequestRulesConnectPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestRulesConnectPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesConnectPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesConnectPhaseAction) SetName(v string) *CreatePropertyVersionRequestRulesConnectPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesConnectPhaseAction) SetOptions(v *CreatePropertyVersionRequestRulesConnectPhaseActionOptions) *CreatePropertyVersionRequestRulesConnectPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyVersionRequestRulesConnectPhaseActionOptions struct {
}

func (s CreatePropertyVersionRequestRulesConnectPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesConnectPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyVersionRequestRulesReqeustPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyVersionRequestRulesReqeustPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreatePropertyVersionRequestRulesReqeustPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesReqeustPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesReqeustPhase) SetName(v string) *CreatePropertyVersionRequestRulesReqeustPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesReqeustPhase) SetCondition(v string) *CreatePropertyVersionRequestRulesReqeustPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesReqeustPhase) SetAction(v *CreatePropertyVersionRequestRulesReqeustPhaseAction) *CreatePropertyVersionRequestRulesReqeustPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyVersionRequestRulesReqeustPhase) SetPriority(v int) *CreatePropertyVersionRequestRulesReqeustPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesReqeustPhase) SetEnabled(v bool) *CreatePropertyVersionRequestRulesReqeustPhase {
  s.Enabled = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesReqeustPhase) SetDescription(v string) *CreatePropertyVersionRequestRulesReqeustPhase {
  s.Description = &v
  return s
}

type CreatePropertyVersionRequestRulesReqeustPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyVersionRequestRulesReqeustPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestRulesReqeustPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesReqeustPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesReqeustPhaseAction) SetName(v string) *CreatePropertyVersionRequestRulesReqeustPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesReqeustPhaseAction) SetOptions(v *CreatePropertyVersionRequestRulesReqeustPhaseActionOptions) *CreatePropertyVersionRequestRulesReqeustPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyVersionRequestRulesReqeustPhaseActionOptions struct {
}

func (s CreatePropertyVersionRequestRulesReqeustPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesReqeustPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyVersionRequestRulesCachePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyVersionRequestRulesCachePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreatePropertyVersionRequestRulesCachePhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesCachePhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesCachePhase) SetName(v string) *CreatePropertyVersionRequestRulesCachePhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesCachePhase) SetCondition(v string) *CreatePropertyVersionRequestRulesCachePhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesCachePhase) SetAction(v *CreatePropertyVersionRequestRulesCachePhaseAction) *CreatePropertyVersionRequestRulesCachePhase {
  s.Action = v
  return s
}

func (s *CreatePropertyVersionRequestRulesCachePhase) SetPriority(v int) *CreatePropertyVersionRequestRulesCachePhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesCachePhase) SetEnabled(v bool) *CreatePropertyVersionRequestRulesCachePhase {
  s.Enabled = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesCachePhase) SetDescription(v string) *CreatePropertyVersionRequestRulesCachePhase {
  s.Description = &v
  return s
}

type CreatePropertyVersionRequestRulesCachePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyVersionRequestRulesCachePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestRulesCachePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesCachePhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesCachePhaseAction) SetName(v string) *CreatePropertyVersionRequestRulesCachePhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesCachePhaseAction) SetOptions(v *CreatePropertyVersionRequestRulesCachePhaseActionOptions) *CreatePropertyVersionRequestRulesCachePhaseAction {
  s.Options = v
  return s
}

type CreatePropertyVersionRequestRulesCachePhaseActionOptions struct {
}

func (s CreatePropertyVersionRequestRulesCachePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesCachePhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyVersionRequestRulesOriginPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyVersionRequestRulesOriginPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreatePropertyVersionRequestRulesOriginPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesOriginPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesOriginPhase) SetName(v string) *CreatePropertyVersionRequestRulesOriginPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesOriginPhase) SetCondition(v string) *CreatePropertyVersionRequestRulesOriginPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesOriginPhase) SetAction(v *CreatePropertyVersionRequestRulesOriginPhaseAction) *CreatePropertyVersionRequestRulesOriginPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyVersionRequestRulesOriginPhase) SetPriority(v int) *CreatePropertyVersionRequestRulesOriginPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesOriginPhase) SetEnabled(v bool) *CreatePropertyVersionRequestRulesOriginPhase {
  s.Enabled = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesOriginPhase) SetDescription(v string) *CreatePropertyVersionRequestRulesOriginPhase {
  s.Description = &v
  return s
}

type CreatePropertyVersionRequestRulesOriginPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyVersionRequestRulesOriginPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestRulesOriginPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesOriginPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesOriginPhaseAction) SetName(v string) *CreatePropertyVersionRequestRulesOriginPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesOriginPhaseAction) SetOptions(v *CreatePropertyVersionRequestRulesOriginPhaseActionOptions) *CreatePropertyVersionRequestRulesOriginPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyVersionRequestRulesOriginPhaseActionOptions struct {
}

func (s CreatePropertyVersionRequestRulesOriginPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesOriginPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyVersionRequestRulesResponsePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyVersionRequestRulesResponsePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreatePropertyVersionRequestRulesResponsePhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesResponsePhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesResponsePhase) SetName(v string) *CreatePropertyVersionRequestRulesResponsePhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesResponsePhase) SetCondition(v string) *CreatePropertyVersionRequestRulesResponsePhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesResponsePhase) SetAction(v *CreatePropertyVersionRequestRulesResponsePhaseAction) *CreatePropertyVersionRequestRulesResponsePhase {
  s.Action = v
  return s
}

func (s *CreatePropertyVersionRequestRulesResponsePhase) SetPriority(v int) *CreatePropertyVersionRequestRulesResponsePhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesResponsePhase) SetEnabled(v bool) *CreatePropertyVersionRequestRulesResponsePhase {
  s.Enabled = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesResponsePhase) SetDescription(v string) *CreatePropertyVersionRequestRulesResponsePhase {
  s.Description = &v
  return s
}

type CreatePropertyVersionRequestRulesResponsePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyVersionRequestRulesResponsePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyVersionRequestRulesResponsePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesResponsePhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestRulesResponsePhaseAction) SetName(v string) *CreatePropertyVersionRequestRulesResponsePhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyVersionRequestRulesResponsePhaseAction) SetOptions(v *CreatePropertyVersionRequestRulesResponsePhaseActionOptions) *CreatePropertyVersionRequestRulesResponsePhaseAction {
  s.Options = v
  return s
}

type CreatePropertyVersionRequestRulesResponsePhaseActionOptions struct {
}

func (s CreatePropertyVersionRequestRulesResponsePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestRulesResponsePhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyVersionRequestCloneParam struct {
  // {"en":"The property Id you need to clone","zh_CN":"克隆项目id"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"The property version you need to clone","zh_CN":"克隆项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
}

func (s CreatePropertyVersionRequestCloneParam) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestCloneParam) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionRequestCloneParam) SetPropertyId(v int) *CreatePropertyVersionRequestCloneParam {
  s.PropertyId = &v
  return s
}

func (s *CreatePropertyVersionRequestCloneParam) SetPropertyVersion(v int) *CreatePropertyVersionRequestCloneParam {
  s.PropertyVersion = &v
  return s
}

type CreatePropertyVersionRequestHeader struct {
}

func (s CreatePropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type CreatePropertyVersionPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionPaths) SetPropertyId(v int) *CreatePropertyVersionPaths {
  s.PropertyId = &v
  return s
}

type CreatePropertyVersionParameters struct {
}

func (s CreatePropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionParameters) GoString() string {
  return s.String()
}

type CreatePropertyVersionResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreatePropertyVersionResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreatePropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionResponse) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionResponse) SetCode(v string) *CreatePropertyVersionResponse {
  s.Code = &v
  return s
}

func (s *CreatePropertyVersionResponse) SetData(v *CreatePropertyVersionResponseData) *CreatePropertyVersionResponse {
  s.Data = v
  return s
}

func (s *CreatePropertyVersionResponse) SetMessage(v string) *CreatePropertyVersionResponse {
  s.Message = &v
  return s
}

type CreatePropertyVersionResponseData struct {
  // {"en":"Property version.","zh_CN":"项目的版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"Property name.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyVersionResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionResponseData) GoString() string {
  return s.String()
}

func (s *CreatePropertyVersionResponseData) SetPropertyVersion(v int) *CreatePropertyVersionResponseData {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyVersionResponseData) SetPropertyName(v string) *CreatePropertyVersionResponseData {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyVersionResponseData) SetPropertyId(v int64) *CreatePropertyVersionResponseData {
  s.PropertyId = &v
  return s
}

type CreatePropertyVersionResponseHeader struct {
}

func (s CreatePropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyVersionResponseHeader) GoString() string {
  return s.String()
}




type QueryTierRouteMapsRequest struct {
}

func (s QueryTierRouteMapsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsRequest) GoString() string {
  return s.String()
}

type QueryTierRouteMapsRequestHeader struct {
}

func (s QueryTierRouteMapsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsRequestHeader) GoString() string {
  return s.String()
}

type QueryTierRouteMapsPaths struct {
}

func (s QueryTierRouteMapsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsPaths) GoString() string {
  return s.String()
}

type QueryTierRouteMapsParameters struct {
  // {"en":"Product Service Type related to your contract.","zh_CN":"产品服务类型。","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Property ID. You should specify 'serviceType' or 'propertyId'.","zh_CN":"项目ID。‘serviceType’与'propertyId'必须指定一个"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"route map type","zh_CN":"回源路由类型","exampleValue":"static,dynamic"}
  TierType *string `json:"tierType,omitempty" xml:"tierType,omitempty"`
}

func (s QueryTierRouteMapsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsParameters) GoString() string {
  return s.String()
}

func (s *QueryTierRouteMapsParameters) SetServiceType(v string) *QueryTierRouteMapsParameters {
  s.ServiceType = &v
  return s
}

func (s *QueryTierRouteMapsParameters) SetPropertyId(v int) *QueryTierRouteMapsParameters {
  s.PropertyId = &v
  return s
}

func (s *QueryTierRouteMapsParameters) SetTierType(v string) *QueryTierRouteMapsParameters {
  s.TierType = &v
  return s
}

type QueryTierRouteMapsResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"List of available route maps","zh_CN":"可用回源路由列表"}
  Data []*QueryTierRouteMapsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTierRouteMapsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsResponse) GoString() string {
  return s.String()
}

func (s *QueryTierRouteMapsResponse) SetCode(v string) *QueryTierRouteMapsResponse {
  s.Code = &v
  return s
}

func (s *QueryTierRouteMapsResponse) SetMessage(v string) *QueryTierRouteMapsResponse {
  s.Message = &v
  return s
}

func (s *QueryTierRouteMapsResponse) SetData(v []*QueryTierRouteMapsResponseData) *QueryTierRouteMapsResponse {
  s.Data = v
  return s
}

type QueryTierRouteMapsResponseData struct     {
  // {"en":"Route map code.","zh_CN":"回源路由标识"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
  // {"en":"Route map type, optional values: [static, dynamic]","zh_CN":"回源路由类型，可选值: [static,dynamic]","exampleValue":"static,dynamic"}
  TierType *string `json:"tierType,omitempty" xml:"tierType,omitempty" require:"true"`
  // {"en":"Map type, optional values:[public,custom]","zh_CN":"回源路由分类，可选值:[public,custom]","exampleValue":"public,custom"}
  MapType *string `json:"mapType,omitempty" xml:"mapType,omitempty" require:"true"`
}

func (s QueryTierRouteMapsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsResponseData) GoString() string {
  return s.String()
}

func (s *QueryTierRouteMapsResponseData) SetRouteMapCode(v string) *QueryTierRouteMapsResponseData {
  s.RouteMapCode = &v
  return s
}

func (s *QueryTierRouteMapsResponseData) SetTierType(v string) *QueryTierRouteMapsResponseData {
  s.TierType = &v
  return s
}

func (s *QueryTierRouteMapsResponseData) SetMapType(v string) *QueryTierRouteMapsResponseData {
  s.MapType = &v
  return s
}

type QueryTierRouteMapsResponseHeader struct {
}

func (s QueryTierRouteMapsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTierRouteMapsResponseHeader) GoString() string {
  return s.String()
}




type GetDeploymentTaskRequest struct {
}

func (s GetDeploymentTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskRequest) GoString() string {
  return s.String()
}

type GetDeploymentTaskRequestHeader struct {
}

func (s GetDeploymentTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskRequestHeader) GoString() string {
  return s.String()
}

type GetDeploymentTaskPaths struct {
  // {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
  DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
}

func (s GetDeploymentTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskPaths) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskPaths) SetDeploymentId(v int) *GetDeploymentTaskPaths {
  s.DeploymentId = &v
  return s
}

type GetDeploymentTaskParameters struct {
}

func (s GetDeploymentTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskParameters) GoString() string {
  return s.String()
}

type GetDeploymentTaskResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetDeploymentTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetDeploymentTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponse) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskResponse) SetCode(v string) *GetDeploymentTaskResponse {
  s.Code = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetMessage(v string) *GetDeploymentTaskResponse {
  s.Message = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetData(v *GetDeploymentTaskResponseData) *GetDeploymentTaskResponse {
  s.Data = v
  return s
}

type GetDeploymentTaskResponseData struct {
  // {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
  DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
  // {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
  DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty" require:"true"`
  // {"en":"Status of Deployment.","zh_CN":"任务状态","exampleValue":"PENDING,IN_PROCESS,SUCCESS,FAIL"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Deployment environment.","zh_CN":"项目的部署环境","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en":"An RFC3339 date indicates when the task was last updated.","zh_CN":"RFC 3339格式的日期，表示任务的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties to the staging or production environments.","zh_CN":"指部署任务所要执行的操作，可以包括项目的部署或卸载操作。"}
  Actions []*GetDeploymentTaskResponseDataActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s GetDeploymentTaskResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponseData) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskResponseData) SetDeploymentId(v int64) *GetDeploymentTaskResponseData {
  s.DeploymentId = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetDeploymentName(v string) *GetDeploymentTaskResponseData {
  s.DeploymentName = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetStatus(v string) *GetDeploymentTaskResponseData {
  s.Status = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetTarget(v string) *GetDeploymentTaskResponseData {
  s.Target = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetSubmissionTime(v string) *GetDeploymentTaskResponseData {
  s.SubmissionTime = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetLastUpdateTime(v string) *GetDeploymentTaskResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetFinishTime(v string) *GetDeploymentTaskResponseData {
  s.FinishTime = &v
  return s
}

func (s *GetDeploymentTaskResponseData) SetActions(v []*GetDeploymentTaskResponseDataActions) *GetDeploymentTaskResponseData {
  s.Actions = v
  return s
}

type GetDeploymentTaskResponseDataActions struct     {
  // {"en":"Describe an action to take. You can deploy a property, remove a property.","zh_CN":"指定操作类型，包括部署加速项目、卸载项目。","exampleValue":"deploy_property,remove_property"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的项目ID。"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Indicates the version of the property to deploy or remove.","zh_CN":"指定要部署或卸载项目的版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s GetDeploymentTaskResponseDataActions) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponseDataActions) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskResponseDataActions) SetAction(v string) *GetDeploymentTaskResponseDataActions {
  s.Action = &v
  return s
}

func (s *GetDeploymentTaskResponseDataActions) SetPropertyId(v int64) *GetDeploymentTaskResponseDataActions {
  s.PropertyId = &v
  return s
}

func (s *GetDeploymentTaskResponseDataActions) SetVersion(v int) *GetDeploymentTaskResponseDataActions {
  s.Version = &v
  return s
}

type GetDeploymentTaskResponseHeader struct {
}

func (s GetDeploymentTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponseHeader) GoString() string {
  return s.String()
}




type QueryDeploymentsForTerraformRequest struct {
}

func (s QueryDeploymentsForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformRequest) GoString() string {
  return s.String()
}

type QueryDeploymentsForTerraformRequestHeader struct {
}

func (s QueryDeploymentsForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryDeploymentsForTerraformPaths struct {
}

func (s QueryDeploymentsForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformPaths) GoString() string {
  return s.String()
}

type QueryDeploymentsForTerraformParameters struct {
  // {"en":"Property ID","zh_CN":"加速项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"Status of Deployment. Enum:PENDING,PENDING_REVIEW,IN_PROCESS,SUCCESS,FAIL","zh_CN":"任务状态。取值范围:PENDING,PENDING_REVIEW,IN_PROCESS,SUCCESS,FAIL"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"加速项目的部署环境。取值范围: staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en":"Indicates the first item to return. The default is '0'.","zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200","zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Order of properties to return. Enum: asc,desc Default: desc","zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en":"Returns results in sorted order. Enum: submissionTime,lastUpdateTime Default: submissionTime","zh_CN":"返回结果的排序依据。取值范围: submissionTime,lastUpdateTime 默认值: submissionTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s QueryDeploymentsForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformParameters) GoString() string {
  return s.String()
}

func (s *QueryDeploymentsForTerraformParameters) SetPropertyId(v int) *QueryDeploymentsForTerraformParameters {
  s.PropertyId = &v
  return s
}

func (s *QueryDeploymentsForTerraformParameters) SetStatus(v string) *QueryDeploymentsForTerraformParameters {
  s.Status = &v
  return s
}

func (s *QueryDeploymentsForTerraformParameters) SetTarget(v string) *QueryDeploymentsForTerraformParameters {
  s.Target = &v
  return s
}

func (s *QueryDeploymentsForTerraformParameters) SetOffset(v int) *QueryDeploymentsForTerraformParameters {
  s.Offset = &v
  return s
}

func (s *QueryDeploymentsForTerraformParameters) SetLimit(v int) *QueryDeploymentsForTerraformParameters {
  s.Limit = &v
  return s
}

func (s *QueryDeploymentsForTerraformParameters) SetSortOrder(v string) *QueryDeploymentsForTerraformParameters {
  s.SortOrder = &v
  return s
}

func (s *QueryDeploymentsForTerraformParameters) SetSortBy(v string) *QueryDeploymentsForTerraformParameters {
  s.SortBy = &v
  return s
}

type QueryDeploymentsForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryDeploymentsForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentsForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryDeploymentsForTerraformResponse) SetCode(v string) *QueryDeploymentsForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponse) SetMessage(v string) *QueryDeploymentsForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponse) SetData(v *QueryDeploymentsForTerraformResponseData) *QueryDeploymentsForTerraformResponse {
  s.Data = v
  return s
}

type QueryDeploymentsForTerraformResponseData struct {
  // {"en":"Number of deployment task.","zh_CN":"部署任务的总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of deployment task summaries.","zh_CN":"部署任务列表。"}
  Deployments []*QueryDeploymentsForTerraformResponseDataDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDeploymentsForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryDeploymentsForTerraformResponseData) SetCount(v int) *QueryDeploymentsForTerraformResponseData {
  s.Count = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseData) SetDeployments(v []*QueryDeploymentsForTerraformResponseDataDeployments) *QueryDeploymentsForTerraformResponseData {
  s.Deployments = v
  return s
}

type QueryDeploymentsForTerraformResponseDataDeployments struct     {
  // {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
  DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
  // {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
  DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty" require:"true"`
  // {"en":"Status of Deployment. Enum:PENDING,IN_PROCESS,SUCCESS,FAIL","zh_CN":"任务状态。取值范围:PENDING,IN_PROCESS,SUCCESS,FAIL"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"项目的部署环境。取值范围: staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en":"An RFC3339 date indicates when the task was last updated.","zh_CN":"RFC 3339格式的日期，表示任务的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
}

func (s QueryDeploymentsForTerraformResponseDataDeployments) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponseDataDeployments) GoString() string {
  return s.String()
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetDeploymentId(v int) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.DeploymentId = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetDeploymentName(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.DeploymentName = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetStatus(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.Status = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetTarget(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.Target = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetSubmissionTime(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.SubmissionTime = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetLastUpdateTime(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetFinishTime(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
  s.FinishTime = &v
  return s
}

type QueryDeploymentsForTerraformResponseHeader struct {
}

func (s QueryDeploymentsForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponseHeader) GoString() string {
  return s.String()
}




type CreateDeploymentTaskRequest struct {
  // {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
  DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
  // {"en":"Indicates whether to deploy to staging or production.","zh_CN":"指定部署任务的目标环境，即演练或生产环境。","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties to the staging or production environments.","zh_CN":"指部署任务所要执行的操作，可以包括项目的部署或卸载操作。"}
  Actions []*CreateDeploymentTaskRequestActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskRequest) SetDeploymentName(v string) *CreateDeploymentTaskRequest {
  s.DeploymentName = &v
  return s
}

func (s *CreateDeploymentTaskRequest) SetTarget(v string) *CreateDeploymentTaskRequest {
  s.Target = &v
  return s
}

func (s *CreateDeploymentTaskRequest) SetActions(v []*CreateDeploymentTaskRequestActions) *CreateDeploymentTaskRequest {
  s.Actions = v
  return s
}

type CreateDeploymentTaskRequestActions struct     {
  // {"en":"Describe an action to take. You can deploy a property, remove a property.","zh_CN":"指定操作类型，包括部署项目、卸载项目。","exampleValue":"deploy_property,remove_property"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的项目ID。"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Indicates the version of the property to deploy or remove.","zh_CN":"指定要部署或卸载项目的版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s CreateDeploymentTaskRequestActions) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskRequestActions) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskRequestActions) SetAction(v string) *CreateDeploymentTaskRequestActions {
  s.Action = &v
  return s
}

func (s *CreateDeploymentTaskRequestActions) SetPropertyId(v int64) *CreateDeploymentTaskRequestActions {
  s.PropertyId = &v
  return s
}

func (s *CreateDeploymentTaskRequestActions) SetVersion(v int) *CreateDeploymentTaskRequestActions {
  s.Version = &v
  return s
}

type CreateDeploymentTaskRequestHeader struct {
}

func (s CreateDeploymentTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskRequestHeader) GoString() string {
  return s.String()
}

type CreateDeploymentTaskPaths struct {
}

func (s CreateDeploymentTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskPaths) GoString() string {
  return s.String()
}

type CreateDeploymentTaskParameters struct {
}

func (s CreateDeploymentTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskParameters) GoString() string {
  return s.String()
}

type CreateDeploymentTaskResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreateDeploymentTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateDeploymentTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskResponse) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskResponse) SetCode(v string) *CreateDeploymentTaskResponse {
  s.Code = &v
  return s
}

func (s *CreateDeploymentTaskResponse) SetMessage(v string) *CreateDeploymentTaskResponse {
  s.Message = &v
  return s
}

func (s *CreateDeploymentTaskResponse) SetData(v *CreateDeploymentTaskResponseData) *CreateDeploymentTaskResponse {
  s.Data = v
  return s
}

type CreateDeploymentTaskResponseData struct {
  // {"en":"ID of the deployment task.","zh_CN":"部署任务标识"}
  DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
}

func (s CreateDeploymentTaskResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskResponseData) GoString() string {
  return s.String()
}

func (s *CreateDeploymentTaskResponseData) SetDeploymentId(v int64) *CreateDeploymentTaskResponseData {
  s.DeploymentId = &v
  return s
}

type CreateDeploymentTaskResponseHeader struct {
}

func (s CreateDeploymentTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeploymentTaskResponseHeader) GoString() string {
  return s.String()
}




type DeletePropertyRequest struct {
}

func (s DeletePropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyRequest) GoString() string {
  return s.String()
}

type DeletePropertyRequestHeader struct {
}

func (s DeletePropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyRequestHeader) GoString() string {
  return s.String()
}

type DeletePropertyPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s DeletePropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyPaths) GoString() string {
  return s.String()
}

func (s *DeletePropertyPaths) SetPropertyId(v int) *DeletePropertyPaths {
  s.PropertyId = &v
  return s
}

type DeletePropertyParameters struct {
}

func (s DeletePropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyParameters) GoString() string {
  return s.String()
}

type DeletePropertyResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeletePropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyResponse) GoString() string {
  return s.String()
}

func (s *DeletePropertyResponse) SetCode(v string) *DeletePropertyResponse {
  s.Code = &v
  return s
}

func (s *DeletePropertyResponse) SetMessage(v string) *DeletePropertyResponse {
  s.Message = &v
  return s
}

type DeletePropertyResponseHeader struct {
}

func (s DeletePropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyResponseHeader) GoString() string {
  return s.String()
}




type UpdatePropertyVersionRequest struct {
  // {"en":"A descriptive comment for the property version. The length must not exceed 256 characters.","zh_CN":"版本描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*UpdatePropertyVersionRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins []*UpdatePropertyVersionRequestOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables []*UpdatePropertyVersionRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *UpdatePropertyVersionRequestRules `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Struct"`
}

func (s UpdatePropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequest) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequest) SetVersionComment(v string) *UpdatePropertyVersionRequest {
  s.VersionComment = &v
  return s
}

func (s *UpdatePropertyVersionRequest) SetHostnames(v []*UpdatePropertyVersionRequestHostnames) *UpdatePropertyVersionRequest {
  s.Hostnames = v
  return s
}

func (s *UpdatePropertyVersionRequest) SetOrigins(v []*UpdatePropertyVersionRequestOrigins) *UpdatePropertyVersionRequest {
  s.Origins = v
  return s
}

func (s *UpdatePropertyVersionRequest) SetVariables(v []*UpdatePropertyVersionRequestVariables) *UpdatePropertyVersionRequest {
  s.Variables = v
  return s
}

func (s *UpdatePropertyVersionRequest) SetRules(v *UpdatePropertyVersionRequestRules) *UpdatePropertyVersionRequest {
  s.Rules = v
  return s
}

type UpdatePropertyVersionRequestHostnames struct     {
  // {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *UpdatePropertyVersionRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
  // {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
  Certificates []*UpdatePropertyVersionRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *UpdatePropertyVersionRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestHostnames) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestHostnames) SetHostname(v string) *UpdatePropertyVersionRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnames) SetDefaultOrigin(v *UpdatePropertyVersionRequestHostnamesDefaultOrigin) *UpdatePropertyVersionRequestHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *UpdatePropertyVersionRequestHostnames) SetCertificates(v []*UpdatePropertyVersionRequestHostnamesCertificates) *UpdatePropertyVersionRequestHostnames {
  s.Certificates = v
  return s
}

func (s *UpdatePropertyVersionRequestHostnames) SetEdgeHostname(v *UpdatePropertyVersionRequestHostnamesEdgeHostname) *UpdatePropertyVersionRequestHostnames {
  s.EdgeHostname = v
  return s
}

type UpdatePropertyVersionRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s UpdatePropertyVersionRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetServers(v []*string) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetHost(v string) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetHttpPort(v int) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetIpVersion(v string) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetScheme(v string) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *UpdatePropertyVersionRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type UpdatePropertyVersionRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;\n- origin_mtls: Origin mTLS Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书；\n- origin_mtls：表示回源mTLS证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s UpdatePropertyVersionRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestHostnamesCertificates) SetCertificateId(v int) *UpdatePropertyVersionRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesCertificates) SetCertificateUsage(v string) *UpdatePropertyVersionRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type UpdatePropertyVersionRequestHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s UpdatePropertyVersionRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *UpdatePropertyVersionRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *UpdatePropertyVersionRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *UpdatePropertyVersionRequestHostnamesEdgeHostname) SetComment(v string) *UpdatePropertyVersionRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

type UpdatePropertyVersionRequestOrigins struct     {
  // {"en":"Origin ID","zh_CN":"源站id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Origin name. Supports upper and lower case of English characters(a-z, A-Z), and dots(.) Underline(_), horizontal bar(-).","zh_CN":"源站名称。支持英文字符大小写（a-z，A-Z），点号(.)、下划线（_）、横杠（-）。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Origin servers list","zh_CN":"源站服务器列表"}
  Servers []*UpdatePropertyVersionRequestOriginsServers `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"Http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"Https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Enable SNI. Whether to enable SNI. If enabled, the origin TLS handshake will carry SNI (Server Name Indication).","zh_CN":"是否启用SNI。开启后，回源TLS握手将携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"SNI Server Name. After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. Default: ${http.request.host}, which follows the client request host. If left empty, the default value will be used.","zh_CN":"SNI服务器。启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。默认值：${http.request.host}，即跟随客户端请求host。空值：置空则与默认值保持一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"defaultValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3","en":"TLS Versions. List of origin TLS protocol versions, supporting configuring multiple versions.","zh_CN":"TLS版本。回源TLS协议版本列表，支持配置多个版本。","exampleValue":"SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
  // {"en":"Verify Origin Certificate. Whether to enable origin certificate verification. If enabled, the back to origin node will verify the certificate offer by origin.","zh_CN":"是否开启节点验证源站证书。开启后，节点访问源站将会验证源站发过来的证书。","exampleValue":"true,false"}
  ProxySSLVerifyEnabled *bool `json:"proxySSLVerifyEnabled,omitempty" xml:"proxySSLVerifyEnabled,omitempty"`
  // {"en":"Trust CA Certificate ID. Trusted CA Certificate for origin certificate verification.","zh_CN":"信任CA证书ID。配置节点验证源站时使用的信任证书。"}
  ProxySSLTrustedCertificate *string `json:"proxySSLTrustedCertificate,omitempty" xml:"proxySSLTrustedCertificate,omitempty"`
}

func (s UpdatePropertyVersionRequestOrigins) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestOrigins) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestOrigins) SetId(v int) *UpdatePropertyVersionRequestOrigins {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetName(v string) *UpdatePropertyVersionRequestOrigins {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetServers(v []*UpdatePropertyVersionRequestOriginsServers) *UpdatePropertyVersionRequestOrigins {
  s.Servers = v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetHttpPort(v int) *UpdatePropertyVersionRequestOrigins {
  s.HttpPort = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetHttpsPort(v int) *UpdatePropertyVersionRequestOrigins {
  s.HttpsPort = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetProxySSLSNIEnabled(v bool) *UpdatePropertyVersionRequestOrigins {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetProxySSLSNIServer(v string) *UpdatePropertyVersionRequestOrigins {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetProxySSLVersion(v []*string) *UpdatePropertyVersionRequestOrigins {
  s.ProxySSLVersion = v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetProxySSLVerifyEnabled(v bool) *UpdatePropertyVersionRequestOrigins {
  s.ProxySSLVerifyEnabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestOrigins) SetProxySSLTrustedCertificate(v string) *UpdatePropertyVersionRequestOrigins {
  s.ProxySSLTrustedCertificate = &v
  return s
}

type UpdatePropertyVersionRequestOriginsServers struct     {
  // {"en":"Origin server","zh_CN":"源站服务器"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en":"priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s UpdatePropertyVersionRequestOriginsServers) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestOriginsServers) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestOriginsServers) SetServer(v string) *UpdatePropertyVersionRequestOriginsServers {
  s.Server = &v
  return s
}

func (s *UpdatePropertyVersionRequestOriginsServers) SetWeight(v int) *UpdatePropertyVersionRequestOriginsServers {
  s.Weight = &v
  return s
}

func (s *UpdatePropertyVersionRequestOriginsServers) SetPriority(v int) *UpdatePropertyVersionRequestOriginsServers {
  s.Priority = &v
  return s
}

type UpdatePropertyVersionRequestVariables struct     {
  // {"en":"variate ID","zh_CN":"变量id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *UpdatePropertyVersionRequestVariablesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestVariables) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestVariables) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestVariables) SetId(v int) *UpdatePropertyVersionRequestVariables {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestVariables) SetCondition(v string) *UpdatePropertyVersionRequestVariables {
  s.Condition = &v
  return s
}

func (s *UpdatePropertyVersionRequestVariables) SetAction(v *UpdatePropertyVersionRequestVariablesAction) *UpdatePropertyVersionRequestVariables {
  s.Action = v
  return s
}

type UpdatePropertyVersionRequestVariablesAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *UpdatePropertyVersionRequestVariablesActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestVariablesAction) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestVariablesAction) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestVariablesAction) SetName(v string) *UpdatePropertyVersionRequestVariablesAction {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestVariablesAction) SetOptions(v *UpdatePropertyVersionRequestVariablesActionOptions) *UpdatePropertyVersionRequestVariablesAction {
  s.Options = v
  return s
}

type UpdatePropertyVersionRequestVariablesActionOptions struct {
  // {"en":"variate name","zh_CN":"变量名称"}
  VarName *string `json:"varName,omitempty" xml:"varName,omitempty"`
  // {"en":"variate value","zh_CN":"变量值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdatePropertyVersionRequestVariablesActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestVariablesActionOptions) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestVariablesActionOptions) SetVarName(v string) *UpdatePropertyVersionRequestVariablesActionOptions {
  s.VarName = &v
  return s
}

func (s *UpdatePropertyVersionRequestVariablesActionOptions) SetValue(v string) *UpdatePropertyVersionRequestVariablesActionOptions {
  s.Value = &v
  return s
}

type UpdatePropertyVersionRequestRules struct {
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ConnectPhase []*UpdatePropertyVersionRequestRulesConnectPhase `json:"connectPhase,omitempty" xml:"connectPhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ReqeustPhase []*UpdatePropertyVersionRequestRulesReqeustPhase `json:"reqeustPhase,omitempty" xml:"reqeustPhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  CachePhase []*UpdatePropertyVersionRequestRulesCachePhase `json:"cachePhase,omitempty" xml:"cachePhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  OriginPhase []*UpdatePropertyVersionRequestRulesOriginPhase `json:"originPhase,omitempty" xml:"originPhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ResponsePhase []*UpdatePropertyVersionRequestRulesResponsePhase `json:"responsePhase,omitempty" xml:"responsePhase,omitempty" type:"Repeated"`
}

func (s UpdatePropertyVersionRequestRules) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRules) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRules) SetConnectPhase(v []*UpdatePropertyVersionRequestRulesConnectPhase) *UpdatePropertyVersionRequestRules {
  s.ConnectPhase = v
  return s
}

func (s *UpdatePropertyVersionRequestRules) SetReqeustPhase(v []*UpdatePropertyVersionRequestRulesReqeustPhase) *UpdatePropertyVersionRequestRules {
  s.ReqeustPhase = v
  return s
}

func (s *UpdatePropertyVersionRequestRules) SetCachePhase(v []*UpdatePropertyVersionRequestRulesCachePhase) *UpdatePropertyVersionRequestRules {
  s.CachePhase = v
  return s
}

func (s *UpdatePropertyVersionRequestRules) SetOriginPhase(v []*UpdatePropertyVersionRequestRulesOriginPhase) *UpdatePropertyVersionRequestRules {
  s.OriginPhase = v
  return s
}

func (s *UpdatePropertyVersionRequestRules) SetResponsePhase(v []*UpdatePropertyVersionRequestRulesResponsePhase) *UpdatePropertyVersionRequestRules {
  s.ResponsePhase = v
  return s
}

type UpdatePropertyVersionRequestRulesConnectPhase struct     {
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *UpdatePropertyVersionRequestRulesConnectPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdatePropertyVersionRequestRulesConnectPhase) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesConnectPhase) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetId(v int) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetName(v string) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetCondition(v string) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Condition = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetAction(v *UpdatePropertyVersionRequestRulesConnectPhaseAction) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Action = v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetPriority(v int) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Priority = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetEnabled(v bool) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Enabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhase) SetDescription(v string) *UpdatePropertyVersionRequestRulesConnectPhase {
  s.Description = &v
  return s
}

type UpdatePropertyVersionRequestRulesConnectPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *UpdatePropertyVersionRequestRulesConnectPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestRulesConnectPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesConnectPhaseAction) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesConnectPhaseAction) SetName(v string) *UpdatePropertyVersionRequestRulesConnectPhaseAction {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesConnectPhaseAction) SetOptions(v *UpdatePropertyVersionRequestRulesConnectPhaseActionOptions) *UpdatePropertyVersionRequestRulesConnectPhaseAction {
  s.Options = v
  return s
}

type UpdatePropertyVersionRequestRulesConnectPhaseActionOptions struct {
}

func (s UpdatePropertyVersionRequestRulesConnectPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesConnectPhaseActionOptions) GoString() string {
  return s.String()
}

type UpdatePropertyVersionRequestRulesReqeustPhase struct     {
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *UpdatePropertyVersionRequestRulesReqeustPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdatePropertyVersionRequestRulesReqeustPhase) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesReqeustPhase) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetId(v int) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetName(v string) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetCondition(v string) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Condition = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetAction(v *UpdatePropertyVersionRequestRulesReqeustPhaseAction) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Action = v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetPriority(v int) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Priority = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetEnabled(v bool) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Enabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhase) SetDescription(v string) *UpdatePropertyVersionRequestRulesReqeustPhase {
  s.Description = &v
  return s
}

type UpdatePropertyVersionRequestRulesReqeustPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *UpdatePropertyVersionRequestRulesReqeustPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestRulesReqeustPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesReqeustPhaseAction) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhaseAction) SetName(v string) *UpdatePropertyVersionRequestRulesReqeustPhaseAction {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesReqeustPhaseAction) SetOptions(v *UpdatePropertyVersionRequestRulesReqeustPhaseActionOptions) *UpdatePropertyVersionRequestRulesReqeustPhaseAction {
  s.Options = v
  return s
}

type UpdatePropertyVersionRequestRulesReqeustPhaseActionOptions struct {
}

func (s UpdatePropertyVersionRequestRulesReqeustPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesReqeustPhaseActionOptions) GoString() string {
  return s.String()
}

type UpdatePropertyVersionRequestRulesCachePhase struct     {
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *UpdatePropertyVersionRequestRulesCachePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdatePropertyVersionRequestRulesCachePhase) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesCachePhase) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetId(v int) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetName(v string) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetCondition(v string) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Condition = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetAction(v *UpdatePropertyVersionRequestRulesCachePhaseAction) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Action = v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetPriority(v int) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Priority = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetEnabled(v bool) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Enabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhase) SetDescription(v string) *UpdatePropertyVersionRequestRulesCachePhase {
  s.Description = &v
  return s
}

type UpdatePropertyVersionRequestRulesCachePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *UpdatePropertyVersionRequestRulesCachePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestRulesCachePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesCachePhaseAction) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesCachePhaseAction) SetName(v string) *UpdatePropertyVersionRequestRulesCachePhaseAction {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesCachePhaseAction) SetOptions(v *UpdatePropertyVersionRequestRulesCachePhaseActionOptions) *UpdatePropertyVersionRequestRulesCachePhaseAction {
  s.Options = v
  return s
}

type UpdatePropertyVersionRequestRulesCachePhaseActionOptions struct {
}

func (s UpdatePropertyVersionRequestRulesCachePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesCachePhaseActionOptions) GoString() string {
  return s.String()
}

type UpdatePropertyVersionRequestRulesOriginPhase struct     {
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *UpdatePropertyVersionRequestRulesOriginPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdatePropertyVersionRequestRulesOriginPhase) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesOriginPhase) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetId(v int) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetName(v string) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetCondition(v string) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Condition = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetAction(v *UpdatePropertyVersionRequestRulesOriginPhaseAction) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Action = v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetPriority(v int) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Priority = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetEnabled(v bool) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Enabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhase) SetDescription(v string) *UpdatePropertyVersionRequestRulesOriginPhase {
  s.Description = &v
  return s
}

type UpdatePropertyVersionRequestRulesOriginPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *UpdatePropertyVersionRequestRulesOriginPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestRulesOriginPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesOriginPhaseAction) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesOriginPhaseAction) SetName(v string) *UpdatePropertyVersionRequestRulesOriginPhaseAction {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesOriginPhaseAction) SetOptions(v *UpdatePropertyVersionRequestRulesOriginPhaseActionOptions) *UpdatePropertyVersionRequestRulesOriginPhaseAction {
  s.Options = v
  return s
}

type UpdatePropertyVersionRequestRulesOriginPhaseActionOptions struct {
}

func (s UpdatePropertyVersionRequestRulesOriginPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesOriginPhaseActionOptions) GoString() string {
  return s.String()
}

type UpdatePropertyVersionRequestRulesResponsePhase struct     {
  // {"en":"rule ID","zh_CN":"规则id"}
  Id *int `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *UpdatePropertyVersionRequestRulesResponsePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdatePropertyVersionRequestRulesResponsePhase) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesResponsePhase) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetId(v int) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Id = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetName(v string) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetCondition(v string) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Condition = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetAction(v *UpdatePropertyVersionRequestRulesResponsePhaseAction) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Action = v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetPriority(v int) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Priority = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetEnabled(v bool) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Enabled = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhase) SetDescription(v string) *UpdatePropertyVersionRequestRulesResponsePhase {
  s.Description = &v
  return s
}

type UpdatePropertyVersionRequestRulesResponsePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *UpdatePropertyVersionRequestRulesResponsePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdatePropertyVersionRequestRulesResponsePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesResponsePhaseAction) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionRequestRulesResponsePhaseAction) SetName(v string) *UpdatePropertyVersionRequestRulesResponsePhaseAction {
  s.Name = &v
  return s
}

func (s *UpdatePropertyVersionRequestRulesResponsePhaseAction) SetOptions(v *UpdatePropertyVersionRequestRulesResponsePhaseActionOptions) *UpdatePropertyVersionRequestRulesResponsePhaseAction {
  s.Options = v
  return s
}

type UpdatePropertyVersionRequestRulesResponsePhaseActionOptions struct {
}

func (s UpdatePropertyVersionRequestRulesResponsePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestRulesResponsePhaseActionOptions) GoString() string {
  return s.String()
}

type UpdatePropertyVersionRequestHeader struct {
}

func (s UpdatePropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type UpdatePropertyVersionPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s UpdatePropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionPaths) SetPropertyId(v int) *UpdatePropertyVersionPaths {
  s.PropertyId = &v
  return s
}

func (s *UpdatePropertyVersionPaths) SetVersion(v int) *UpdatePropertyVersionPaths {
  s.Version = &v
  return s
}

type UpdatePropertyVersionParameters struct {
}

func (s UpdatePropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionParameters) GoString() string {
  return s.String()
}

type UpdatePropertyVersionResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdatePropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionResponse) GoString() string {
  return s.String()
}

func (s *UpdatePropertyVersionResponse) SetCode(v string) *UpdatePropertyVersionResponse {
  s.Code = &v
  return s
}

func (s *UpdatePropertyVersionResponse) SetMessage(v string) *UpdatePropertyVersionResponse {
  s.Message = &v
  return s
}

type UpdatePropertyVersionResponseHeader struct {
}

func (s UpdatePropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyVersionResponseHeader) GoString() string {
  return s.String()
}




type QueryPropertyConfigForTerrformRequest struct {
}

func (s QueryPropertyConfigForTerrformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformRequest) GoString() string {
  return s.String()
}

type QueryPropertyConfigForTerrformRequestHeader struct {
}

func (s QueryPropertyConfigForTerrformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformRequestHeader) GoString() string {
  return s.String()
}

type QueryPropertyConfigForTerrformPaths struct {
  // {"en":"Property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformPaths) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformPaths) SetPropertyId(v int) *QueryPropertyConfigForTerrformPaths {
  s.PropertyId = &v
  return s
}

type QueryPropertyConfigForTerrformParameters struct {
}

func (s QueryPropertyConfigForTerrformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformParameters) GoString() string {
  return s.String()
}

type QueryPropertyConfigForTerrformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryPropertyConfigForTerrformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertyConfigForTerrformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponse) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponse) SetCode(v string) *QueryPropertyConfigForTerrformResponse {
  s.Code = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponse) SetMessage(v string) *QueryPropertyConfigForTerrformResponse {
  s.Message = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponse) SetData(v *QueryPropertyConfigForTerrformResponseData) *QueryPropertyConfigForTerrformResponse {
  s.Data = v
  return s
}

type QueryPropertyConfigForTerrformResponseData struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  PropertyCreationTime *string `json:"propertyCreationTime,omitempty" xml:"propertyCreationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  PropertyLastUpdateTime *string `json:"propertyLastUpdateTime,omitempty" xml:"propertyLastUpdateTime,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *QueryPropertyConfigForTerrformResponseDataStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *QueryPropertyConfigForTerrformResponseDataProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"A property version. It must be an integer value >=1.","zh_CN":"项目的版本，必须是大于0的整数。"}
  CurrentVersion *int `json:"currentVersion,omitempty" xml:"currentVersion,omitempty" require:"true"`
  // {"en":"A description of the version.","zh_CN":"版本描述。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty" require:"true"`
  // {"en":"Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.","zh_CN":"该项目版本是否处于冻结状态。当加速项目版本部署后即进入冻结状态，不可再更新该版本。"}
  Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the version was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  VersionCreationTime *string `json:"versionCreationTime,omitempty" xml:"versionCreationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the version was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  VersionLastUpdateTime *string `json:"versionLastUpdateTime,omitempty" xml:"versionLastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*QueryPropertyConfigForTerrformResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *string `json:"rules,omitempty" xml:"rules,omitempty" require:"true"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins *string `json:"origins,omitempty" xml:"origins,omitempty" require:"true"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables *string `json:"variables,omitempty" xml:"variables,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseData) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyId(v int) *QueryPropertyConfigForTerrformResponseData {
  s.PropertyId = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyName(v string) *QueryPropertyConfigForTerrformResponseData {
  s.PropertyName = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyComment(v string) *QueryPropertyConfigForTerrformResponseData {
  s.PropertyComment = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetServiceType(v string) *QueryPropertyConfigForTerrformResponseData {
  s.ServiceType = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyCreationTime(v string) *QueryPropertyConfigForTerrformResponseData {
  s.PropertyCreationTime = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyLastUpdateTime(v string) *QueryPropertyConfigForTerrformResponseData {
  s.PropertyLastUpdateTime = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetStagingVersion(v *QueryPropertyConfigForTerrformResponseDataStagingVersion) *QueryPropertyConfigForTerrformResponseData {
  s.StagingVersion = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetProductionVersion(v *QueryPropertyConfigForTerrformResponseDataProductionVersion) *QueryPropertyConfigForTerrformResponseData {
  s.ProductionVersion = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetStagingDeployingVersion(v *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) *QueryPropertyConfigForTerrformResponseData {
  s.StagingDeployingVersion = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetProductionDeployingVersion(v *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) *QueryPropertyConfigForTerrformResponseData {
  s.ProductionDeployingVersion = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetLatestVersion(v int) *QueryPropertyConfigForTerrformResponseData {
  s.LatestVersion = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetCurrentVersion(v int) *QueryPropertyConfigForTerrformResponseData {
  s.CurrentVersion = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVersionComment(v string) *QueryPropertyConfigForTerrformResponseData {
  s.VersionComment = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetFrozen(v bool) *QueryPropertyConfigForTerrformResponseData {
  s.Frozen = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVersionCreationTime(v string) *QueryPropertyConfigForTerrformResponseData {
  s.VersionCreationTime = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVersionLastUpdateTime(v string) *QueryPropertyConfigForTerrformResponseData {
  s.VersionLastUpdateTime = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetHostnames(v []*QueryPropertyConfigForTerrformResponseDataHostnames) *QueryPropertyConfigForTerrformResponseData {
  s.Hostnames = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetRules(v string) *QueryPropertyConfigForTerrformResponseData {
  s.Rules = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetOrigins(v string) *QueryPropertyConfigForTerrformResponseData {
  s.Origins = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVariables(v string) *QueryPropertyConfigForTerrformResponseData {
  s.Variables = &v
  return s
}

type QueryPropertyConfigForTerrformResponseDataStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataStagingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataStagingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataStagingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyConfigForTerrformResponseDataProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataProductionVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataProductionVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataProductionVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyConfigForTerrformResponseDataHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" require:"true" type:"Struct"`
  // {"en":"hostname","zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true" type:"Struct"`
  // {"en":"hostname association certificate configuration","zh_CN":"关联证书配置"}
  Certificates []*QueryPropertyConfigForTerrformResponseDataHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
  // {"en":"icp","zh_CN":"域名备案号"}
  Icp *string `json:"icp,omitempty" xml:"icp,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnames) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetDefaultOrigin(v *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) *QueryPropertyConfigForTerrformResponseDataHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetHostname(v string) *QueryPropertyConfigForTerrformResponseDataHostnames {
  s.Hostname = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetEdgeHostname(v *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) *QueryPropertyConfigForTerrformResponseDataHostnames {
  s.EdgeHostname = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetCertificates(v []*QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) *QueryPropertyConfigForTerrformResponseDataHostnames {
  s.Certificates = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetIcp(v string) *QueryPropertyConfigForTerrformResponseDataHostnames {
  s.Icp = &v
  return s
}

type QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" require:"true"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty" require:"true"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty" require:"true"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty" require:"true"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty" require:"true"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetServers(v []*string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetIpVersion(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpPort(v int) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHost(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpsPort(v int) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetScheme(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"dns service status.data range: inactive,active","zh_CN":"DNS服务状态。取值范围：inactive, active。备注：inactive：挂起，active：生效"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetDnsServiceStatus(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostname(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.EdgeHostname = &v
  return s
}

type QueryPropertyConfigForTerrformResponseDataHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. data range: default_sni, ssl_bk, gm_sm2_enc, gm_sm2_sign, client_mtls","zh_CN":"证书用途。取值范围：default_sni, dual_sni, gm_sm2_enc, gm_sm2_sign, client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) SetCertificateId(v int) *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) SetCertificateUsage(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type QueryPropertyConfigForTerrformResponseHeader struct {
}

func (s QueryPropertyConfigForTerrformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseHeader) GoString() string {
  return s.String()
}




type UpdatePropertyForTerraformRequest struct {
  // {"en":"The name of the property. The length must not exceed 256 characters.","zh_CN":"项目的名称。长度不超过256个字符。"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
  // {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*UpdatePropertyForTerraformRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins *string `json:"origins,omitempty" xml:"origins,omitempty"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
  // {"en":"Rules","zh_CN":"规则配置"}
  Rules *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s UpdatePropertyForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequest) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformRequest) SetPropertyName(v string) *UpdatePropertyForTerraformRequest {
  s.PropertyName = &v
  return s
}

func (s *UpdatePropertyForTerraformRequest) SetPropertyComment(v string) *UpdatePropertyForTerraformRequest {
  s.PropertyComment = &v
  return s
}

func (s *UpdatePropertyForTerraformRequest) SetVersionComment(v string) *UpdatePropertyForTerraformRequest {
  s.VersionComment = &v
  return s
}

func (s *UpdatePropertyForTerraformRequest) SetHostnames(v []*UpdatePropertyForTerraformRequestHostnames) *UpdatePropertyForTerraformRequest {
  s.Hostnames = v
  return s
}

func (s *UpdatePropertyForTerraformRequest) SetOrigins(v string) *UpdatePropertyForTerraformRequest {
  s.Origins = &v
  return s
}

func (s *UpdatePropertyForTerraformRequest) SetVariables(v string) *UpdatePropertyForTerraformRequest {
  s.Variables = &v
  return s
}

func (s *UpdatePropertyForTerraformRequest) SetRules(v string) *UpdatePropertyForTerraformRequest {
  s.Rules = &v
  return s
}

type UpdatePropertyForTerraformRequestHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
  // {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
  Certificates []*UpdatePropertyForTerraformRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *UpdatePropertyForTerraformRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s UpdatePropertyForTerraformRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnames) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetDefaultOrigin(v *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) *UpdatePropertyForTerraformRequestHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetHostname(v string) *UpdatePropertyForTerraformRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetCertificates(v []*UpdatePropertyForTerraformRequestHostnamesCertificates) *UpdatePropertyForTerraformRequestHostnames {
  s.Certificates = v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetEdgeHostname(v *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) *UpdatePropertyForTerraformRequestHostnames {
  s.EdgeHostname = v
  return s
}

type UpdatePropertyForTerraformRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetServers(v []*string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetIpVersion(v string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpPort(v int) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHost(v string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetScheme(v string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type UpdatePropertyForTerraformRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;\n- origin_mtls: Origin mTLS Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书；\n- origin_mtls：表示回源mTLS证书。"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnamesCertificates) SetCertificateId(v int) *UpdatePropertyForTerraformRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesCertificates) SetCertificateUsage(v string) *UpdatePropertyForTerraformRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type UpdatePropertyForTerraformRequestHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *UpdatePropertyForTerraformRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) SetComment(v string) *UpdatePropertyForTerraformRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *UpdatePropertyForTerraformRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

type UpdatePropertyForTerraformRequestHeader struct {
}

func (s UpdatePropertyForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHeader) GoString() string {
  return s.String()
}

type UpdatePropertyForTerraformPaths struct {
  // {"en":"Property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformPaths) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformPaths) SetPropertyId(v int) *UpdatePropertyForTerraformPaths {
  s.PropertyId = &v
  return s
}

type UpdatePropertyForTerraformParameters struct {
}

func (s UpdatePropertyForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformParameters) GoString() string {
  return s.String()
}

type UpdatePropertyForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *UpdatePropertyForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s UpdatePropertyForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformResponse) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformResponse) SetCode(v string) *UpdatePropertyForTerraformResponse {
  s.Code = &v
  return s
}

func (s *UpdatePropertyForTerraformResponse) SetMessage(v string) *UpdatePropertyForTerraformResponse {
  s.Message = &v
  return s
}

func (s *UpdatePropertyForTerraformResponse) SetData(v *UpdatePropertyForTerraformResponseData) *UpdatePropertyForTerraformResponse {
  s.Data = v
  return s
}

type UpdatePropertyForTerraformResponseData struct {
  // {"en":"Property ID","zh_CN":"项目ID"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Property Name","zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property Version","zh_CN":"项目版本"}
  PropertyVersion *int64 `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *UpdatePropertyForTerraformResponseData) SetPropertyId(v int64) *UpdatePropertyForTerraformResponseData {
  s.PropertyId = &v
  return s
}

func (s *UpdatePropertyForTerraformResponseData) SetPropertyName(v string) *UpdatePropertyForTerraformResponseData {
  s.PropertyName = &v
  return s
}

func (s *UpdatePropertyForTerraformResponseData) SetPropertyVersion(v int64) *UpdatePropertyForTerraformResponseData {
  s.PropertyVersion = &v
  return s
}

type UpdatePropertyForTerraformResponseHeader struct {
}

func (s UpdatePropertyForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformResponseHeader) GoString() string {
  return s.String()
}




type ListPropertiesRequest struct {
}

func (s ListPropertiesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesRequest) GoString() string {
  return s.String()
}

type ListPropertiesRequestHeader struct {
}

func (s ListPropertiesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesRequestHeader) GoString() string {
  return s.String()
}

type ListPropertiesPaths struct {
}

func (s ListPropertiesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesPaths) GoString() string {
  return s.String()
}

type ListPropertiesParameters struct {
  // {"en":"Unique identifier for the product","zh_CN":"服务类型","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Deployment Environment","zh_CN":"加速项目的部署环境","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"defaultValue":"0","en":"Indicates the first item to return.","zh_CN":"查询起始位置，取值范围：>= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"defaultValue":"100","en":"Maximum number of properties to return.  Range: <= 200","zh_CN":"每次查询的最大条数。取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"defaultValue":"desc","en":"Order of properties to return.","zh_CN":"返回结果的顺序。默认按最后更新时间降序。","exampleValue":"asc,desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"defaultValue":"lastUpdateTime","en":"Returns results in sorted order.","zh_CN":"返回结果的排序依据。","exampleValue":"creationTime,lastUpdateTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}

func (s ListPropertiesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesParameters) GoString() string {
  return s.String()
}

func (s *ListPropertiesParameters) SetServiceType(v string) *ListPropertiesParameters {
  s.ServiceType = &v
  return s
}

func (s *ListPropertiesParameters) SetTarget(v string) *ListPropertiesParameters {
  s.Target = &v
  return s
}

func (s *ListPropertiesParameters) SetOffset(v int) *ListPropertiesParameters {
  s.Offset = &v
  return s
}

func (s *ListPropertiesParameters) SetLimit(v int) *ListPropertiesParameters {
  s.Limit = &v
  return s
}

func (s *ListPropertiesParameters) SetSortOrder(v string) *ListPropertiesParameters {
  s.SortOrder = &v
  return s
}

func (s *ListPropertiesParameters) SetSortBy(v string) *ListPropertiesParameters {
  s.SortBy = &v
  return s
}

func (s *ListPropertiesParameters) SetContractId(v string) *ListPropertiesParameters {
  s.ContractId = &v
  return s
}

func (s *ListPropertiesParameters) SetItemId(v string) *ListPropertiesParameters {
  s.ItemId = &v
  return s
}

type ListPropertiesResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *ListPropertiesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListPropertiesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponse) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponse) SetCode(v string) *ListPropertiesResponse {
  s.Code = &v
  return s
}

func (s *ListPropertiesResponse) SetMessage(v string) *ListPropertiesResponse {
  s.Message = &v
  return s
}

func (s *ListPropertiesResponse) SetData(v *ListPropertiesResponseData) *ListPropertiesResponse {
  s.Data = v
  return s
}

type ListPropertiesResponseData struct {
  // {"en":"Number of properties.","zh_CN":"项目数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of properties.","zh_CN":"项目列表。"}
  Properties []*ListPropertiesResponseDataProperties `json:"properties,omitempty" xml:"properties,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseData) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseData) SetCount(v int) *ListPropertiesResponseData {
  s.Count = &v
  return s
}

func (s *ListPropertiesResponseData) SetProperties(v []*ListPropertiesResponseDataProperties) *ListPropertiesResponseData {
  s.Properties = v
  return s
}

type ListPropertiesResponseDataProperties struct     {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product.","zh_CN":"服务类型","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *ListPropertiesResponseDataPropertiesStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *ListPropertiesResponseDataPropertiesProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *ListPropertiesResponseDataPropertiesStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *ListPropertiesResponseDataPropertiesProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"The id of contract","zh_CN":"合同号"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty" require:"true"`
  // {"en":"The id of product","zh_CN":"产品号"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s ListPropertiesResponseDataProperties) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataProperties) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataProperties) SetPropertyId(v int64) *ListPropertiesResponseDataProperties {
  s.PropertyId = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetPropertyName(v string) *ListPropertiesResponseDataProperties {
  s.PropertyName = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetPropertyComment(v string) *ListPropertiesResponseDataProperties {
  s.PropertyComment = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetServiceType(v string) *ListPropertiesResponseDataProperties {
  s.ServiceType = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetCreationTime(v string) *ListPropertiesResponseDataProperties {
  s.CreationTime = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetLastUpdateTime(v string) *ListPropertiesResponseDataProperties {
  s.LastUpdateTime = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetLatestVersion(v int) *ListPropertiesResponseDataProperties {
  s.LatestVersion = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetStagingVersion(v *ListPropertiesResponseDataPropertiesStagingVersion) *ListPropertiesResponseDataProperties {
  s.StagingVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetProductionVersion(v *ListPropertiesResponseDataPropertiesProductionVersion) *ListPropertiesResponseDataProperties {
  s.ProductionVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetStagingDeployingVersion(v *ListPropertiesResponseDataPropertiesStagingDeployingVersion) *ListPropertiesResponseDataProperties {
  s.StagingDeployingVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetProductionDeployingVersion(v *ListPropertiesResponseDataPropertiesProductionDeployingVersion) *ListPropertiesResponseDataProperties {
  s.ProductionDeployingVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetContractId(v string) *ListPropertiesResponseDataProperties {
  s.ContractId = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetItemId(v string) *ListPropertiesResponseDataProperties {
  s.ItemId = &v
  return s
}

type ListPropertiesResponseDataPropertiesStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesStagingVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesStagingVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesStagingVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesStagingVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesStagingVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseDataPropertiesProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesProductionVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesProductionVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesProductionVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesProductionVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesProductionVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseDataPropertiesStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesStagingDeployingVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesStagingDeployingVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseDataPropertiesProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesProductionDeployingVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesProductionDeployingVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseHeader struct {
}

func (s ListPropertiesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseHeader) GoString() string {
  return s.String()
}




type ListDeploymentTasksRequest struct {
}

func (s ListDeploymentTasksRequest) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksRequest) GoString() string {
  return s.String()
}

type ListDeploymentTasksRequestHeader struct {
}

func (s ListDeploymentTasksRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksRequestHeader) GoString() string {
  return s.String()
}

type ListDeploymentTasksPaths struct {
}

func (s ListDeploymentTasksPaths) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksPaths) GoString() string {
  return s.String()
}

type ListDeploymentTasksParameters struct {
  // {"en":"Property ID","zh_CN":"加速项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"Status of Deployment","zh_CN":"任务状态","exampleValue":"PENDING,IN_PROCESS,SUCCESS,FAIL"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"Deployment environment","zh_CN":"加速项目的部署环境","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"defaultValue":"0","en":"Indicates the first item to return.","zh_CN":"查询起始位置。取值范围: >= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"defaultValue":"100","en":"Maximum number of tasks to return per query. Range: <= 200","zh_CN":"每次查询的最大条数。取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"defaultValue":"desc","en":"The order in which results are returned. asc: ascending order; desc: descending order.","zh_CN":"返回结果的顺序。asc：升序；desc：降序。”","exampleValue":"asc,desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"defaultValue":"submissionTime","en":"The field by which to sort the results. submissionTime: Task submission time; lastUpdateTime: Task last update time.","zh_CN":"返回结果的排序依据。submissionTime：任务提交时间；lastUpdateTime：任务更新时间。","exampleValue":"submissionTime,lastUpdateTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s ListDeploymentTasksParameters) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksParameters) GoString() string {
  return s.String()
}

func (s *ListDeploymentTasksParameters) SetPropertyId(v int) *ListDeploymentTasksParameters {
  s.PropertyId = &v
  return s
}

func (s *ListDeploymentTasksParameters) SetStatus(v string) *ListDeploymentTasksParameters {
  s.Status = &v
  return s
}

func (s *ListDeploymentTasksParameters) SetTarget(v string) *ListDeploymentTasksParameters {
  s.Target = &v
  return s
}

func (s *ListDeploymentTasksParameters) SetOffset(v int) *ListDeploymentTasksParameters {
  s.Offset = &v
  return s
}

func (s *ListDeploymentTasksParameters) SetLimit(v int) *ListDeploymentTasksParameters {
  s.Limit = &v
  return s
}

func (s *ListDeploymentTasksParameters) SetSortOrder(v string) *ListDeploymentTasksParameters {
  s.SortOrder = &v
  return s
}

func (s *ListDeploymentTasksParameters) SetSortBy(v string) *ListDeploymentTasksParameters {
  s.SortBy = &v
  return s
}

type ListDeploymentTasksResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' indicates success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *ListDeploymentTasksResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListDeploymentTasksResponse) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksResponse) GoString() string {
  return s.String()
}

func (s *ListDeploymentTasksResponse) SetCode(v string) *ListDeploymentTasksResponse {
  s.Code = &v
  return s
}

func (s *ListDeploymentTasksResponse) SetMessage(v string) *ListDeploymentTasksResponse {
  s.Message = &v
  return s
}

func (s *ListDeploymentTasksResponse) SetData(v *ListDeploymentTasksResponseData) *ListDeploymentTasksResponse {
  s.Data = v
  return s
}

type ListDeploymentTasksResponseData struct {
  // {"en":"Total number of deployment tasks.","zh_CN":"部署任务的总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of deployment task summaries.","zh_CN":"部署任务列表。"}
  Deployments []*ListDeploymentTasksResponseDataDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" require:"true" type:"Repeated"`
}

func (s ListDeploymentTasksResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksResponseData) GoString() string {
  return s.String()
}

func (s *ListDeploymentTasksResponseData) SetCount(v int) *ListDeploymentTasksResponseData {
  s.Count = &v
  return s
}

func (s *ListDeploymentTasksResponseData) SetDeployments(v []*ListDeploymentTasksResponseDataDeployments) *ListDeploymentTasksResponseData {
  s.Deployments = v
  return s
}

type ListDeploymentTasksResponseDataDeployments struct     {
  // {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
  DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
  // {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
  DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty" require:"true"`
  // {"en":"Status of Deployment","zh_CN":"任务状态","exampleValue":"PENDING,IN_PROCESS,SUCCESS,FAIL"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"项目的部署环境。取值范围: staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en":"An RFC3339 date indicates when the task was last updated.","zh_CN":"RFC 3339格式的日期，表示任务的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
}

func (s ListDeploymentTasksResponseDataDeployments) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksResponseDataDeployments) GoString() string {
  return s.String()
}

func (s *ListDeploymentTasksResponseDataDeployments) SetDeploymentId(v int64) *ListDeploymentTasksResponseDataDeployments {
  s.DeploymentId = &v
  return s
}

func (s *ListDeploymentTasksResponseDataDeployments) SetDeploymentName(v string) *ListDeploymentTasksResponseDataDeployments {
  s.DeploymentName = &v
  return s
}

func (s *ListDeploymentTasksResponseDataDeployments) SetStatus(v string) *ListDeploymentTasksResponseDataDeployments {
  s.Status = &v
  return s
}

func (s *ListDeploymentTasksResponseDataDeployments) SetTarget(v string) *ListDeploymentTasksResponseDataDeployments {
  s.Target = &v
  return s
}

func (s *ListDeploymentTasksResponseDataDeployments) SetSubmissionTime(v string) *ListDeploymentTasksResponseDataDeployments {
  s.SubmissionTime = &v
  return s
}

func (s *ListDeploymentTasksResponseDataDeployments) SetLastUpdateTime(v string) *ListDeploymentTasksResponseDataDeployments {
  s.LastUpdateTime = &v
  return s
}

func (s *ListDeploymentTasksResponseDataDeployments) SetFinishTime(v string) *ListDeploymentTasksResponseDataDeployments {
  s.FinishTime = &v
  return s
}

type ListDeploymentTasksResponseHeader struct {
}

func (s ListDeploymentTasksResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListDeploymentTasksResponseHeader) GoString() string {
  return s.String()
}




type GetPropertyRequest struct {
}

func (s GetPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyRequest) GoString() string {
  return s.String()
}

type GetPropertyRequestHeader struct {
}

func (s GetPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyRequestHeader) GoString() string {
  return s.String()
}

type GetPropertyPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s GetPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyPaths) GoString() string {
  return s.String()
}

func (s *GetPropertyPaths) SetPropertyId(v int) *GetPropertyPaths {
  s.PropertyId = &v
  return s
}

type GetPropertyParameters struct {
}

func (s GetPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyParameters) GoString() string {
  return s.String()
}

type GetPropertyResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetPropertyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponse) GoString() string {
  return s.String()
}

func (s *GetPropertyResponse) SetCode(v string) *GetPropertyResponse {
  s.Code = &v
  return s
}

func (s *GetPropertyResponse) SetMessage(v string) *GetPropertyResponse {
  s.Message = &v
  return s
}

func (s *GetPropertyResponse) SetData(v *GetPropertyResponseData) *GetPropertyResponse {
  s.Data = v
  return s
}

type GetPropertyResponseData struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product.","zh_CN":"服务类型","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *GetPropertyResponseDataStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *GetPropertyResponseDataProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *GetPropertyResponseDataStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *GetPropertyResponseDataProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty" require:"true"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s GetPropertyResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseData) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseData) SetPropertyId(v int64) *GetPropertyResponseData {
  s.PropertyId = &v
  return s
}

func (s *GetPropertyResponseData) SetPropertyName(v string) *GetPropertyResponseData {
  s.PropertyName = &v
  return s
}

func (s *GetPropertyResponseData) SetPropertyComment(v string) *GetPropertyResponseData {
  s.PropertyComment = &v
  return s
}

func (s *GetPropertyResponseData) SetServiceType(v string) *GetPropertyResponseData {
  s.ServiceType = &v
  return s
}

func (s *GetPropertyResponseData) SetCreationTime(v string) *GetPropertyResponseData {
  s.CreationTime = &v
  return s
}

func (s *GetPropertyResponseData) SetLastUpdateTime(v string) *GetPropertyResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *GetPropertyResponseData) SetLatestVersion(v int) *GetPropertyResponseData {
  s.LatestVersion = &v
  return s
}

func (s *GetPropertyResponseData) SetStagingVersion(v *GetPropertyResponseDataStagingVersion) *GetPropertyResponseData {
  s.StagingVersion = v
  return s
}

func (s *GetPropertyResponseData) SetProductionVersion(v *GetPropertyResponseDataProductionVersion) *GetPropertyResponseData {
  s.ProductionVersion = v
  return s
}

func (s *GetPropertyResponseData) SetStagingDeployingVersion(v *GetPropertyResponseDataStagingDeployingVersion) *GetPropertyResponseData {
  s.StagingDeployingVersion = v
  return s
}

func (s *GetPropertyResponseData) SetProductionDeployingVersion(v *GetPropertyResponseDataProductionDeployingVersion) *GetPropertyResponseData {
  s.ProductionDeployingVersion = v
  return s
}

func (s *GetPropertyResponseData) SetContractId(v string) *GetPropertyResponseData {
  s.ContractId = &v
  return s
}

func (s *GetPropertyResponseData) SetItemId(v string) *GetPropertyResponseData {
  s.ItemId = &v
  return s
}

type GetPropertyResponseDataStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataStagingVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataStagingVersion) SetVersion(v int) *GetPropertyResponseDataStagingVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataStagingVersion) SetHostnames(v []*string) *GetPropertyResponseDataStagingVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseDataProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataProductionVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataProductionVersion) SetVersion(v int) *GetPropertyResponseDataProductionVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataProductionVersion) SetHostnames(v []*string) *GetPropertyResponseDataProductionVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseDataStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataStagingDeployingVersion) SetVersion(v int) *GetPropertyResponseDataStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataStagingDeployingVersion) SetHostnames(v []*string) *GetPropertyResponseDataStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseDataProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataProductionDeployingVersion) SetVersion(v int) *GetPropertyResponseDataProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataProductionDeployingVersion) SetHostnames(v []*string) *GetPropertyResponseDataProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseHeader struct {
}

func (s GetPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseHeader) GoString() string {
  return s.String()
}




type QueryPropertiesForTerraformRequest struct {
}

func (s QueryPropertiesForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformRequest) GoString() string {
  return s.String()
}

type QueryPropertiesForTerraformRequestHeader struct {
}

func (s QueryPropertiesForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryPropertiesForTerraformPaths struct {
}

func (s QueryPropertiesForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformPaths) GoString() string {
  return s.String()
}

type QueryPropertiesForTerraformParameters struct {
  // {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"加速项目的部署环境。取值范围: staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en":"Indicates the first item to return. The default is '0'.","zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200","zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Order of properties to return. Enum: asc,desc Default: desc","zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en":"Returns results in sorted order. Enum: creationTime,lastUpdateTime Default: lastUpdateTime","zh_CN":"返回结果的排序依据。取值范围: creationTime,lastUpdateTime 默认值: lastUpdateTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en":"Hostname","zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
}

func (s QueryPropertiesForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformParameters) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformParameters) SetServiceType(v string) *QueryPropertiesForTerraformParameters {
  s.ServiceType = &v
  return s
}

func (s *QueryPropertiesForTerraformParameters) SetTarget(v string) *QueryPropertiesForTerraformParameters {
  s.Target = &v
  return s
}

func (s *QueryPropertiesForTerraformParameters) SetOffset(v int) *QueryPropertiesForTerraformParameters {
  s.Offset = &v
  return s
}

func (s *QueryPropertiesForTerraformParameters) SetLimit(v int) *QueryPropertiesForTerraformParameters {
  s.Limit = &v
  return s
}

func (s *QueryPropertiesForTerraformParameters) SetSortOrder(v string) *QueryPropertiesForTerraformParameters {
  s.SortOrder = &v
  return s
}

func (s *QueryPropertiesForTerraformParameters) SetSortBy(v string) *QueryPropertiesForTerraformParameters {
  s.SortBy = &v
  return s
}

func (s *QueryPropertiesForTerraformParameters) SetHostname(v string) *QueryPropertiesForTerraformParameters {
  s.Hostname = &v
  return s
}

type QueryPropertiesForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryPropertiesForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertiesForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponse) SetCode(v string) *QueryPropertiesForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryPropertiesForTerraformResponse) SetMessage(v string) *QueryPropertiesForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryPropertiesForTerraformResponse) SetData(v *QueryPropertiesForTerraformResponseData) *QueryPropertiesForTerraformResponse {
  s.Data = v
  return s
}

type QueryPropertiesForTerraformResponseData struct {
  // {"en":"Number of properties.","zh_CN":"项目数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of properties.","zh_CN":"项目列表。"}
  Properties []*QueryPropertiesForTerraformResponseDataProperties `json:"properties,omitempty" xml:"properties,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponseData) SetCount(v int) *QueryPropertiesForTerraformResponseData {
  s.Count = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseData) SetProperties(v []*QueryPropertiesForTerraformResponseDataProperties) *QueryPropertiesForTerraformResponseData {
  s.Properties = v
  return s
}

type QueryPropertiesForTerraformResponseDataProperties struct     {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertiesForTerraformResponseDataProperties) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataProperties) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetPropertyId(v int) *QueryPropertiesForTerraformResponseDataProperties {
  s.PropertyId = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetPropertyName(v string) *QueryPropertiesForTerraformResponseDataProperties {
  s.PropertyName = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetPropertyComment(v string) *QueryPropertiesForTerraformResponseDataProperties {
  s.PropertyComment = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetServiceType(v string) *QueryPropertiesForTerraformResponseDataProperties {
  s.ServiceType = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetCreationTime(v string) *QueryPropertiesForTerraformResponseDataProperties {
  s.CreationTime = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetLastUpdateTime(v string) *QueryPropertiesForTerraformResponseDataProperties {
  s.LastUpdateTime = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetLatestVersion(v int) *QueryPropertiesForTerraformResponseDataProperties {
  s.LatestVersion = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetStagingVersion(v *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) *QueryPropertiesForTerraformResponseDataProperties {
  s.StagingVersion = v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetProductionVersion(v *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) *QueryPropertiesForTerraformResponseDataProperties {
  s.ProductionVersion = v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetStagingDeployingVersion(v *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) *QueryPropertiesForTerraformResponseDataProperties {
  s.StagingDeployingVersion = v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetProductionDeployingVersion(v *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) *QueryPropertiesForTerraformResponseDataProperties {
  s.ProductionDeployingVersion = v
  return s
}

type QueryPropertiesForTerraformResponseDataPropertiesStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertiesForTerraformResponseDataPropertiesProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion {
  s.Hostnames = v
  return s
}

type QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertiesForTerraformResponseHeader struct {
}

func (s QueryPropertiesForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseHeader) GoString() string {
  return s.String()
}




type ListPropertyVersionsRequest struct {
}

func (s ListPropertyVersionsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsRequest) GoString() string {
  return s.String()
}

type ListPropertyVersionsRequestHeader struct {
}

func (s ListPropertyVersionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsRequestHeader) GoString() string {
  return s.String()
}

type ListPropertyVersionsPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s ListPropertyVersionsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsPaths) GoString() string {
  return s.String()
}

func (s *ListPropertyVersionsPaths) SetPropertyId(v int) *ListPropertyVersionsPaths {
  s.PropertyId = &v
  return s
}

type ListPropertyVersionsParameters struct {
  // {"defaultValue":"0","en":"Indicates the first item to return.","zh_CN":"查询起始位置。取值范围: >= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"defaultValue":"100","en":"Maximum number of properties to return. Range: <= 200","zh_CN":"每次查询的最大条数。取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"defaultValue":"desc","en":"Order of properties to return.","zh_CN":"返回结果的顺序","exampleValue":"asc,desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"defaultValue":"version","en":"Returns results in sorted order.","zh_CN":"返回结果的排序依据","exampleValue":"version,lastUpdateTime "}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s ListPropertyVersionsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsParameters) GoString() string {
  return s.String()
}

func (s *ListPropertyVersionsParameters) SetOffset(v int) *ListPropertyVersionsParameters {
  s.Offset = &v
  return s
}

func (s *ListPropertyVersionsParameters) SetLimit(v int) *ListPropertyVersionsParameters {
  s.Limit = &v
  return s
}

func (s *ListPropertyVersionsParameters) SetSortOrder(v string) *ListPropertyVersionsParameters {
  s.SortOrder = &v
  return s
}

func (s *ListPropertyVersionsParameters) SetSortBy(v string) *ListPropertyVersionsParameters {
  s.SortBy = &v
  return s
}

type ListPropertyVersionsResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *ListPropertyVersionsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListPropertyVersionsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsResponse) GoString() string {
  return s.String()
}

func (s *ListPropertyVersionsResponse) SetCode(v string) *ListPropertyVersionsResponse {
  s.Code = &v
  return s
}

func (s *ListPropertyVersionsResponse) SetMessage(v string) *ListPropertyVersionsResponse {
  s.Message = &v
  return s
}

func (s *ListPropertyVersionsResponse) SetData(v *ListPropertyVersionsResponseData) *ListPropertyVersionsResponse {
  s.Data = v
  return s
}

type ListPropertyVersionsResponseData struct {
  // {"en":"Number of properties.","zh_CN":"项目数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"A summary of each version of the property.","zh_CN":"项目每个版本的摘要。"}
  PropertyVersions []*ListPropertyVersionsResponseDataPropertyVersions `json:"propertyVersions,omitempty" xml:"propertyVersions,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertyVersionsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsResponseData) GoString() string {
  return s.String()
}

func (s *ListPropertyVersionsResponseData) SetCount(v int) *ListPropertyVersionsResponseData {
  s.Count = &v
  return s
}

func (s *ListPropertyVersionsResponseData) SetPropertyVersions(v []*ListPropertyVersionsResponseDataPropertyVersions) *ListPropertyVersionsResponseData {
  s.PropertyVersions = v
  return s
}

type ListPropertyVersionsResponseDataPropertyVersions struct     {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.","zh_CN":"该项目版本是否处于冻结状态。当项目版本部署后即进入冻结状态，不可再更新该版本。"}
  Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
  // {"en":"Show this version deploy or deploying target info.","zh_CN":"环境部署信息","exampleValue":"roduction,staging,productionDeploying,stagingDeploying"}
  TargetDeployInfo []*string `json:"targetDeployInfo,omitempty" xml:"targetDeployInfo,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertyVersionsResponseDataPropertyVersions) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsResponseDataPropertyVersions) GoString() string {
  return s.String()
}

func (s *ListPropertyVersionsResponseDataPropertyVersions) SetVersion(v int) *ListPropertyVersionsResponseDataPropertyVersions {
  s.Version = &v
  return s
}

func (s *ListPropertyVersionsResponseDataPropertyVersions) SetVersionComment(v string) *ListPropertyVersionsResponseDataPropertyVersions {
  s.VersionComment = &v
  return s
}

func (s *ListPropertyVersionsResponseDataPropertyVersions) SetCreationTime(v string) *ListPropertyVersionsResponseDataPropertyVersions {
  s.CreationTime = &v
  return s
}

func (s *ListPropertyVersionsResponseDataPropertyVersions) SetLastUpdateTime(v string) *ListPropertyVersionsResponseDataPropertyVersions {
  s.LastUpdateTime = &v
  return s
}

func (s *ListPropertyVersionsResponseDataPropertyVersions) SetFrozen(v bool) *ListPropertyVersionsResponseDataPropertyVersions {
  s.Frozen = &v
  return s
}

func (s *ListPropertyVersionsResponseDataPropertyVersions) SetTargetDeployInfo(v []*string) *ListPropertyVersionsResponseDataPropertyVersions {
  s.TargetDeployInfo = v
  return s
}

type ListPropertyVersionsResponseHeader struct {
}

func (s ListPropertyVersionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPropertyVersionsResponseHeader) GoString() string {
  return s.String()
}




type QueryPropertyVersionConfigForTerrformRequest struct {
}

func (s QueryPropertyVersionConfigForTerrformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformRequest) GoString() string {
  return s.String()
}

type QueryPropertyVersionConfigForTerrformRequestHeader struct {
}

func (s QueryPropertyVersionConfigForTerrformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformRequestHeader) GoString() string {
  return s.String()
}

type QueryPropertyVersionConfigForTerrformPaths struct {
  // {"en":"Property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Property Version","zh_CN":"项目版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformPaths) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformPaths) SetPropertyId(v int) *QueryPropertyVersionConfigForTerrformPaths {
  s.PropertyId = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformPaths) SetVersion(v int) *QueryPropertyVersionConfigForTerrformPaths {
  s.Version = &v
  return s
}

type QueryPropertyVersionConfigForTerrformParameters struct {
}

func (s QueryPropertyVersionConfigForTerrformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformParameters) GoString() string {
  return s.String()
}

type QueryPropertyVersionConfigForTerrformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *QueryPropertyVersionConfigForTerrformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertyVersionConfigForTerrformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponse) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponse) SetCode(v string) *QueryPropertyVersionConfigForTerrformResponse {
  s.Code = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponse) SetMessage(v string) *QueryPropertyVersionConfigForTerrformResponse {
  s.Message = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponse) SetData(v *QueryPropertyVersionConfigForTerrformResponseData) *QueryPropertyVersionConfigForTerrformResponse {
  s.Data = v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseData struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  PropertyCreationTime *string `json:"propertyCreationTime,omitempty" xml:"propertyCreationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  PropertyLastUpdateTime *string `json:"propertyLastUpdateTime,omitempty" xml:"propertyLastUpdateTime,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"A property version. It must be an integer value >=1.","zh_CN":"项目的版本，必须是大于0的整数。"}
  CurrentVersion *int `json:"currentVersion,omitempty" xml:"currentVersion,omitempty" require:"true"`
  // {"en":"A description of the version.","zh_CN":"版本描述。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty" require:"true"`
  // {"en":"Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.","zh_CN":"该项目版本是否处于冻结状态。当加速项目版本部署后即进入冻结状态，不可再更新该版本。"}
  Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the version was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  VersionCreationTime *string `json:"versionCreationTime,omitempty" xml:"versionCreationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the version was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  VersionLastUpdateTime *string `json:"versionLastUpdateTime,omitempty" xml:"versionLastUpdateTime,omitempty" require:"true"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*QueryPropertyVersionConfigForTerrformResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *string `json:"rules,omitempty" xml:"rules,omitempty" require:"true"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins *string `json:"origins,omitempty" xml:"origins,omitempty" require:"true"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables *string `json:"variables,omitempty" xml:"variables,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseData) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyId(v int) *QueryPropertyVersionConfigForTerrformResponseData {
  s.PropertyId = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyName(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.PropertyName = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyComment(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.PropertyComment = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetServiceType(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.ServiceType = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyCreationTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.PropertyCreationTime = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyLastUpdateTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.PropertyLastUpdateTime = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetStagingVersion(v *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) *QueryPropertyVersionConfigForTerrformResponseData {
  s.StagingVersion = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetProductionVersion(v *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) *QueryPropertyVersionConfigForTerrformResponseData {
  s.ProductionVersion = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetStagingDeployingVersion(v *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) *QueryPropertyVersionConfigForTerrformResponseData {
  s.StagingDeployingVersion = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetProductionDeployingVersion(v *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) *QueryPropertyVersionConfigForTerrformResponseData {
  s.ProductionDeployingVersion = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetLatestVersion(v int) *QueryPropertyVersionConfigForTerrformResponseData {
  s.LatestVersion = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetCurrentVersion(v int) *QueryPropertyVersionConfigForTerrformResponseData {
  s.CurrentVersion = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVersionComment(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.VersionComment = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetFrozen(v bool) *QueryPropertyVersionConfigForTerrformResponseData {
  s.Frozen = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVersionCreationTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.VersionCreationTime = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVersionLastUpdateTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.VersionLastUpdateTime = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetHostnames(v []*QueryPropertyVersionConfigForTerrformResponseDataHostnames) *QueryPropertyVersionConfigForTerrformResponseData {
  s.Hostnames = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetRules(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.Rules = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetOrigins(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.Origins = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVariables(v string) *QueryPropertyVersionConfigForTerrformResponseData {
  s.Variables = &v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" require:"true" type:"Struct"`
  // {"en":"hostname","zh_CN":"域名"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true" type:"Struct"`
  // {"en":"hostname association certificate configuration","zh_CN":"关联证书配置"}
  Certificates []*QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
  // {"en":"icp","zh_CN":"域名备案号"}
  Icp *string `json:"icp,omitempty" xml:"icp,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnames) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnames) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetDefaultOrigin(v *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetHostname(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
  s.Hostname = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetEdgeHostname(v *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
  s.EdgeHostname = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetCertificates(v []*QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
  s.Certificates = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetIcp(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
  s.Icp = &v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" require:"true"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty" require:"true"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty" require:"true"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty" require:"true"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty" require:"true"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetServers(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetIpVersion(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpPort(v int) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHost(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpsPort(v int) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetScheme(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"dns service status.data range: inactive,active","zh_CN":"DNS服务状态。取值范围：inactive, active。备注：inactive：挂起，active：生效"}
  DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"edge-hostname","zh_CN":"调度域名"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetDnsServiceStatus(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.DnsServiceStatus = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostname(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
  s.EdgeHostname = &v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. data range: default_sni, ssl_bk, gm_sm2_enc, gm_sm2_sign, client_mtls","zh_CN":"证书用途。取值范围：default_sni, dual_sni, gm_sm2_enc, gm_sm2_sign, client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) SetCertificateId(v int) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) SetCertificateUsage(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type QueryPropertyVersionConfigForTerrformResponseHeader struct {
}

func (s QueryPropertyVersionConfigForTerrformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseHeader) GoString() string {
  return s.String()
}




type DeletePropertyForTerraformRequest struct {
}

func (s DeletePropertyForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyForTerraformRequest) GoString() string {
  return s.String()
}

type DeletePropertyForTerraformRequestHeader struct {
}

func (s DeletePropertyForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyForTerraformRequestHeader) GoString() string {
  return s.String()
}

type DeletePropertyForTerraformPaths struct {
  // {"en":"Property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s DeletePropertyForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyForTerraformPaths) GoString() string {
  return s.String()
}

func (s *DeletePropertyForTerraformPaths) SetPropertyId(v int) *DeletePropertyForTerraformPaths {
  s.PropertyId = &v
  return s
}

type DeletePropertyForTerraformParameters struct {
}

func (s DeletePropertyForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyForTerraformParameters) GoString() string {
  return s.String()
}

type DeletePropertyForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeletePropertyForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyForTerraformResponse) GoString() string {
  return s.String()
}

func (s *DeletePropertyForTerraformResponse) SetCode(v string) *DeletePropertyForTerraformResponse {
  s.Code = &v
  return s
}

func (s *DeletePropertyForTerraformResponse) SetMessage(v string) *DeletePropertyForTerraformResponse {
  s.Message = &v
  return s
}

type DeletePropertyForTerraformResponseHeader struct {
}

func (s DeletePropertyForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePropertyForTerraformResponseHeader) GoString() string {
  return s.String()
}




type CreatePropertyForTerraformRequest struct {
  // {"en":"Product Service Type related to your contract.","zh_CN":"产品服务类型。请根据您的合同产品服务类型填写。","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"The name of the property. The length must not exceed 256 characters.","zh_CN":"项目的名称。长度不超过256个字符。"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
  // {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*CreatePropertyForTerraformRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins *string `json:"origins,omitempty" xml:"origins,omitempty"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
  // {"en":"Rules","zh_CN":"规则配置，详情查看schema接口。"}
  Rules *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s CreatePropertyForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequest) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformRequest) SetServiceType(v string) *CreatePropertyForTerraformRequest {
  s.ServiceType = &v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetPropertyName(v string) *CreatePropertyForTerraformRequest {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetPropertyComment(v string) *CreatePropertyForTerraformRequest {
  s.PropertyComment = &v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetVersionComment(v string) *CreatePropertyForTerraformRequest {
  s.VersionComment = &v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetHostnames(v []*CreatePropertyForTerraformRequestHostnames) *CreatePropertyForTerraformRequest {
  s.Hostnames = v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetOrigins(v string) *CreatePropertyForTerraformRequest {
  s.Origins = &v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetVariables(v string) *CreatePropertyForTerraformRequest {
  s.Variables = &v
  return s
}

func (s *CreatePropertyForTerraformRequest) SetRules(v string) *CreatePropertyForTerraformRequest {
  s.Rules = &v
  return s
}

type CreatePropertyForTerraformRequestHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *CreatePropertyForTerraformRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
  // {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
  Certificates []*CreatePropertyForTerraformRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *CreatePropertyForTerraformRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s CreatePropertyForTerraformRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnames) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnames) SetDefaultOrigin(v *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) *CreatePropertyForTerraformRequestHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnames) SetHostname(v string) *CreatePropertyForTerraformRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnames) SetCertificates(v []*CreatePropertyForTerraformRequestHostnamesCertificates) *CreatePropertyForTerraformRequestHostnames {
  s.Certificates = v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnames) SetEdgeHostname(v *CreatePropertyForTerraformRequestHostnamesEdgeHostname) *CreatePropertyForTerraformRequestHostnames {
  s.EdgeHostname = v
  return s
}

type CreatePropertyForTerraformRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s CreatePropertyForTerraformRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetScheme(v string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type CreatePropertyForTerraformRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;\n- origin_mtls: Origin mTLS Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书；\n- origin_mtls：表示回源mTLS证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s CreatePropertyForTerraformRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyForTerraformRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyForTerraformRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type CreatePropertyForTerraformRequestHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
}

func (s CreatePropertyForTerraformRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyForTerraformRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyForTerraformRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

func (s *CreatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyForTerraformRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

type CreatePropertyForTerraformRequestHeader struct {
}

func (s CreatePropertyForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHeader) GoString() string {
  return s.String()
}

type CreatePropertyForTerraformPaths struct {
}

func (s CreatePropertyForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformPaths) GoString() string {
  return s.String()
}

type CreatePropertyForTerraformParameters struct {
}

func (s CreatePropertyForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformParameters) GoString() string {
  return s.String()
}

type CreatePropertyForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreatePropertyForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreatePropertyForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformResponse) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformResponse) SetCode(v string) *CreatePropertyForTerraformResponse {
  s.Code = &v
  return s
}

func (s *CreatePropertyForTerraformResponse) SetMessage(v string) *CreatePropertyForTerraformResponse {
  s.Message = &v
  return s
}

func (s *CreatePropertyForTerraformResponse) SetData(v *CreatePropertyForTerraformResponseData) *CreatePropertyForTerraformResponse {
  s.Data = v
  return s
}

type CreatePropertyForTerraformResponseData struct {
  // {"en":"Property ID","zh_CN":"项目ID"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Property Name","zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property Version","zh_CN":"项目版本"}
  PropertyVersion *int64 `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
}

func (s CreatePropertyForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *CreatePropertyForTerraformResponseData) SetPropertyId(v int64) *CreatePropertyForTerraformResponseData {
  s.PropertyId = &v
  return s
}

func (s *CreatePropertyForTerraformResponseData) SetPropertyName(v string) *CreatePropertyForTerraformResponseData {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyForTerraformResponseData) SetPropertyVersion(v int64) *CreatePropertyForTerraformResponseData {
  s.PropertyVersion = &v
  return s
}

type CreatePropertyForTerraformResponseHeader struct {
}

func (s CreatePropertyForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForTerraformResponseHeader) GoString() string {
  return s.String()
}




type QueryIpSegmentByRouteMapCodeRequest struct {
}

func (s QueryIpSegmentByRouteMapCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodeRequest) GoString() string {
  return s.String()
}

type QueryIpSegmentByRouteMapCodeRequestHeader struct {
}

func (s QueryIpSegmentByRouteMapCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryIpSegmentByRouteMapCodePaths struct {
}

func (s QueryIpSegmentByRouteMapCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodePaths) GoString() string {
  return s.String()
}

type QueryIpSegmentByRouteMapCodeParameters struct {
  // {"en":"Route map code","zh_CN":"回源路由标识"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
}

func (s QueryIpSegmentByRouteMapCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodeParameters) GoString() string {
  return s.String()
}

func (s *QueryIpSegmentByRouteMapCodeParameters) SetRouteMapCode(v string) *QueryIpSegmentByRouteMapCodeParameters {
  s.RouteMapCode = &v
  return s
}

type QueryIpSegmentByRouteMapCodeResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"IP segment list","zh_CN":"IP段列表"}
  Data []*QueryIpSegmentByRouteMapCodeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIpSegmentByRouteMapCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodeResponse) GoString() string {
  return s.String()
}

func (s *QueryIpSegmentByRouteMapCodeResponse) SetCode(v string) *QueryIpSegmentByRouteMapCodeResponse {
  s.Code = &v
  return s
}

func (s *QueryIpSegmentByRouteMapCodeResponse) SetMessage(v string) *QueryIpSegmentByRouteMapCodeResponse {
  s.Message = &v
  return s
}

func (s *QueryIpSegmentByRouteMapCodeResponse) SetData(v []*QueryIpSegmentByRouteMapCodeResponseData) *QueryIpSegmentByRouteMapCodeResponse {
  s.Data = v
  return s
}

type QueryIpSegmentByRouteMapCodeResponseData struct     {
  // {"en":"IP type","zh_CN":"IP类型","exampleValue":"IPv4,IPv6"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty" require:"true"`
  // {"en":"IP segment","zh_CN":"IP段"}
  IpSegment *string `json:"ipSegment,omitempty" xml:"ipSegment,omitempty" require:"true"`
}

func (s QueryIpSegmentByRouteMapCodeResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodeResponseData) GoString() string {
  return s.String()
}

func (s *QueryIpSegmentByRouteMapCodeResponseData) SetIpType(v string) *QueryIpSegmentByRouteMapCodeResponseData {
  s.IpType = &v
  return s
}

func (s *QueryIpSegmentByRouteMapCodeResponseData) SetIpSegment(v string) *QueryIpSegmentByRouteMapCodeResponseData {
  s.IpSegment = &v
  return s
}

type QueryIpSegmentByRouteMapCodeResponseHeader struct {
}

func (s QueryIpSegmentByRouteMapCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIpSegmentByRouteMapCodeResponseHeader) GoString() string {
  return s.String()
}




