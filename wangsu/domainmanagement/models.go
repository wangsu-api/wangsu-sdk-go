package domainmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type AddDomainRequest struct {
  // {"en":"Added domain
  // Use English half-width
  // semicolon between
  // two domains if there
  // are multiple to be
  // added.", "zh_CN":"添加的域名
  // 如果添加多个域名，用英文半角分号分隔。"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Return Chinese results for null (default)
  // En: Return the English prompt result", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s AddDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s AddDomainRequest) GoString() string {
  return s.String()
}

func (s *AddDomainRequest) SetDomainName(v string) *AddDomainRequest {
  s.DomainName = &v
  return s
}

func (s *AddDomainRequest) SetLanguage(v string) *AddDomainRequest {
  s.Language = &v
  return s
}

type AddDomainResponse struct {
  // {"en":"Status code. For detailed description of resCode, please refer to 'Status Codes of Dispatch Business'.", "zh_CN":"状态码，详细说明请参见“业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"recordId, ID of host name record, used to identify this record.", "zh_CN":"域名的详细说明。"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s AddDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s AddDomainResponse) GoString() string {
  return s.String()
}

func (s *AddDomainResponse) SetResCode(v string) *AddDomainResponse {
  s.ResCode = &v
  return s
}

func (s *AddDomainResponse) SetMsg(v string) *AddDomainResponse {
  s.Msg = &v
  return s
}

func (s *AddDomainResponse) SetContent(v []*string) *AddDomainResponse {
  s.Content = v
  return s
}

type AddDomainPaths struct {
}

func (s AddDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s AddDomainPaths) GoString() string {
  return s.String()
}

type AddDomainParameters struct {
}

func (s AddDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s AddDomainParameters) GoString() string {
  return s.String()
}

type AddDomainRequestHeader struct {
}

func (s AddDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDomainRequestHeader) GoString() string {
  return s.String()
}

type AddDomainResponseHeader struct {
}

func (s AddDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDomainResponseHeader) GoString() string {
  return s.String()
}




type UpdateCustomerAnycastIPRecordStatusRequest struct {
  // {"en":"Ip list", "zh_CN":"ip 列表"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" require:"true" type:"Repeated"`
  // {"en":"Record Status, For example, lock or unlock; data of length 1 or 2 can be passed.", "zh_CN":"记录状态，例如锁定、非锁定等，可以传长度为1或2的数据"}
  RecordStatus *string `json:"recordStatus,omitempty" xml:"recordStatus,omitempty" require:"true"`
}

func (s UpdateCustomerAnycastIPRecordStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusRequest) GoString() string {
  return s.String()
}

func (s *UpdateCustomerAnycastIPRecordStatusRequest) SetIps(v []*string) *UpdateCustomerAnycastIPRecordStatusRequest {
  s.Ips = v
  return s
}

func (s *UpdateCustomerAnycastIPRecordStatusRequest) SetRecordStatus(v string) *UpdateCustomerAnycastIPRecordStatusRequest {
  s.RecordStatus = &v
  return s
}

type UpdateCustomerAnycastIPRecordStatusResponse struct {
  // {"en":"The error code that appears when the HTTP status is not 202, indicating the type of error for the current request.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateCustomerAnycastIPRecordStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusResponse) GoString() string {
  return s.String()
}

func (s *UpdateCustomerAnycastIPRecordStatusResponse) SetCode(v string) *UpdateCustomerAnycastIPRecordStatusResponse {
  s.Code = &v
  return s
}

func (s *UpdateCustomerAnycastIPRecordStatusResponse) SetMessage(v string) *UpdateCustomerAnycastIPRecordStatusResponse {
  s.Message = &v
  return s
}

type UpdateCustomerAnycastIPRecordStatusPaths struct {
}

func (s UpdateCustomerAnycastIPRecordStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusPaths) GoString() string {
  return s.String()
}

type UpdateCustomerAnycastIPRecordStatusParameters struct {
}

func (s UpdateCustomerAnycastIPRecordStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusParameters) GoString() string {
  return s.String()
}

type UpdateCustomerAnycastIPRecordStatusRequestHeader struct {
}

func (s UpdateCustomerAnycastIPRecordStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusRequestHeader) GoString() string {
  return s.String()
}

type UpdateCustomerAnycastIPRecordStatusResponseHeader struct {
}

func (s UpdateCustomerAnycastIPRecordStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusResponseHeader) GoString() string {
  return s.String()
}




type DeleteCdnDomainServiceRequest struct {
  // {"en":"Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domainId
  // 2. After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domainId of the domain name; You can also query domainId through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名在系统中对应的ID
  // 注意：
  // 1、参看请求示例中的url，123344对应的就是domainId
  // 2、创建域名成功提交后，返回参数中的location访问url中，能够查询到域名对应的domainId；也可以通过【获取域名配置】和【获取域名列表】接口查询到domainId"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DeleteCdnDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCdnDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *DeleteCdnDomainServiceRequest) SetDomainId(v string) *DeleteCdnDomainServiceRequest {
  s.DomainId = &v
  return s
}

type DeleteCdnDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DeleteCdnDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCdnDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DeleteCdnDomainServiceResponse) SetCode(v string) *DeleteCdnDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DeleteCdnDomainServiceResponse) SetMessage(v string) *DeleteCdnDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DeleteCdnDomainServiceResponse) SetHttpStatus(v int) *DeleteCdnDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DeleteCdnDomainServiceResponse) SetXCncRequestId(v string) *DeleteCdnDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type DeleteCdnDomainServicePaths struct {
  // {"en":"", "zh_CN":"域名名称或域名id，在请求的url后面"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s DeleteCdnDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCdnDomainServicePaths) GoString() string {
  return s.String()
}

func (s *DeleteCdnDomainServicePaths) SetDomainName(v string) *DeleteCdnDomainServicePaths {
  s.DomainName = &v
  return s
}

type DeleteCdnDomainServiceParameters struct {
}

func (s DeleteCdnDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCdnDomainServiceParameters) GoString() string {
  return s.String()
}

type DeleteCdnDomainServiceRequestHeader struct {
}

func (s DeleteCdnDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCdnDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type DeleteCdnDomainServiceResponseHeader struct {
}

func (s DeleteCdnDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCdnDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type PreDeployChangeServerConfigRequest struct {
  // {"en":"Change servers configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <change-servers/>", "zh_CN":"【接入域名跳转】
  // 注意：
  // 1、需要取消【接入域名跳转】时，可以传入空节点<change-servers></change-servers>。
  // 2、表示需要设置【接入域名跳转】，此项必填"}
  ChangeServers []*PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers `json:"change-servers,omitempty" xml:"change-servers,omitempty" require:"true" type:"Repeated"`
}

func (s PreDeployChangeServerConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigRequest) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigRequest) SetChangeServers(v []*PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers) *PreDeployChangeServerConfigRequest {
  s.ChangeServers = v
  return s
}

type PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers struct     {
  // {"en":"If it is a universal domain name, set it to a universal domain name, for example, *.56.com.", "zh_CN":"如果是泛域名，需要填写为泛域名，例如：*.56.com"}
  TargetServer *string `json:"target-server,omitempty" xml:"target-server,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: 
  // A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. 
  // B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. 
  // C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. 
  // D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. 
  // E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", 
  //       "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： 
  // a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； 
  // b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； 
  // c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； 
  // e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。
  // "}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers) SetTargetServer(v string) *PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers {
  s.TargetServer = &v
  return s
}

func (s *PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers) SetDataId(v int) *PreDeployChangeServerConfigPreDeployChangeServerConfigRequestChangeServers {
  s.DataId = &v
  return s
}

type PreDeployChangeServerConfigResponse struct {
  // {"en":"The error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data.", "zh_CN":"返回数据体"}
  Data *PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s PreDeployChangeServerConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigResponse) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigResponse) SetCode(v string) *PreDeployChangeServerConfigResponse {
  s.Code = &v
  return s
}

func (s *PreDeployChangeServerConfigResponse) SetMessage(v string) *PreDeployChangeServerConfigResponse {
  s.Message = &v
  return s
}

func (s *PreDeployChangeServerConfigResponse) SetData(v *PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData) *PreDeployChangeServerConfigResponse {
  s.Data = v
  return s
}

func (s *PreDeployChangeServerConfigResponse) SetXCncRequestId(v string) *PreDeployChangeServerConfigResponse {
  s.XCncRequestId = &v
  return s
}

type PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData struct {
  // {"en":"The preliminary deployment id", "zh_CN":"预部署id"}
  PreDeployId *string `json:"preDeployId,omitempty" xml:"preDeployId,omitempty" require:"true"`
}

func (s PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData) SetPreDeployId(v string) *PreDeployChangeServerConfigPreDeployChangeServerConfigResponseData {
  s.PreDeployId = &v
  return s
}

type PreDeployChangeServerConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s PreDeployChangeServerConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigPaths) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigPaths) SetDomain(v string) *PreDeployChangeServerConfigPaths {
  s.Domain = &v
  return s
}

type PreDeployChangeServerConfigParameters struct {
}

func (s PreDeployChangeServerConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigParameters) GoString() string {
  return s.String()
}

type PreDeployChangeServerConfigRequestHeader struct {
}

func (s PreDeployChangeServerConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigRequestHeader) GoString() string {
  return s.String()
}

type PreDeployChangeServerConfigResponseHeader struct {
}

func (s PreDeployChangeServerConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigResponseHeader) GoString() string {
  return s.String()
}




type AddDomainGroupRequest struct {
  // {"en":"	Domain group name. Only Chinese characters/English alphabet/numbers/underscores are supported, and case insensitive.Up to 32 characters.", "zh_CN":"域名组名称，仅支持输入中文/英文/数字/下划线，不区分大小写，最多可传32个字符。"}
  DomainGroupName *string `json:"domainGroupName,omitempty" xml:"domainGroupName,omitempty" require:"true"`
  // {"en":"The serviceType that the domain group belongs to.", "zh_CN":"域名组所属加速服务类型"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"Domains associated with domain group.The domain needs to be under the serviceType.", "zh_CN":"域名组关联的域名，域名需要是传入加速服务类型下的域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s AddDomainGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s AddDomainGroupRequest) GoString() string {
  return s.String()
}

func (s *AddDomainGroupRequest) SetDomainGroupName(v string) *AddDomainGroupRequest {
  s.DomainGroupName = &v
  return s
}

func (s *AddDomainGroupRequest) SetServiceType(v string) *AddDomainGroupRequest {
  s.ServiceType = &v
  return s
}

func (s *AddDomainGroupRequest) SetDomainList(v []*string) *AddDomainGroupRequest {
  s.DomainList = v
  return s
}

type AddDomainGroupResponse struct {
  // {"en":"Status Code", "zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddDomainGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s AddDomainGroupResponse) GoString() string {
  return s.String()
}

func (s *AddDomainGroupResponse) SetCode(v string) *AddDomainGroupResponse {
  s.Code = &v
  return s
}

func (s *AddDomainGroupResponse) SetMessage(v string) *AddDomainGroupResponse {
  s.Message = &v
  return s
}

type AddDomainGroupPaths struct {
}

func (s AddDomainGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s AddDomainGroupPaths) GoString() string {
  return s.String()
}

type AddDomainGroupParameters struct {
}

func (s AddDomainGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s AddDomainGroupParameters) GoString() string {
  return s.String()
}

type AddDomainGroupRequestHeader struct {
}

func (s AddDomainGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDomainGroupRequestHeader) GoString() string {
  return s.String()
}

type AddDomainGroupResponseHeader struct {
}

func (s AddDomainGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDomainGroupResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainByOriginIPRequest struct {
}

func (s QueryDomainByOriginIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPRequest) GoString() string {
  return s.String()
}

type QueryDomainByOriginIPResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Code =200 indicates that relevant data was returned successfully", "zh_CN":"code=200，表示成功返回相关数据"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Query the source station IP", "zh_CN":"查询的源站IP"}
  Originip *string `json:"originip,omitempty" xml:"originip,omitempty" require:"true"`
  // {"en":"Returns a list of domain name names corresponding to each source station IP", "zh_CN":"返回各个源站IP对应的域名名称列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainByOriginIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainByOriginIPResponse) SetHttpStatus(v int) *QueryDomainByOriginIPResponse {
  s.HttpStatus = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetXCncRequestId(v string) *QueryDomainByOriginIPResponse {
  s.XCncRequestId = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetCode(v int) *QueryDomainByOriginIPResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetOriginip(v string) *QueryDomainByOriginIPResponse {
  s.Originip = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetDomainList(v []*string) *QueryDomainByOriginIPResponse {
  s.DomainList = v
  return s
}

type QueryDomainByOriginIPPaths struct {
}

func (s QueryDomainByOriginIPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPPaths) GoString() string {
  return s.String()
}

type QueryDomainByOriginIPParameters struct {
  // {"en":"Source station IP, multiple IPs separated by semicolons", "zh_CN":"源站IP，多个IP用分号隔开
  // 注意：
  // 1、每次查询最多只能传入10个源站IP
  // 2、不支持源站域名的查询
  // 3、高级源匹配到对应IP时也能查到对应域名"}
  Originip *string `json:"originip,omitempty" xml:"originip,omitempty" require:"true"`
}

func (s QueryDomainByOriginIPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainByOriginIPParameters) SetOriginip(v string) *QueryDomainByOriginIPParameters {
  s.Originip = &v
  return s
}

type QueryDomainByOriginIPRequestHeader struct {
}

func (s QueryDomainByOriginIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainByOriginIPResponseHeader struct {
}

func (s QueryDomainByOriginIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPResponseHeader) GoString() string {
  return s.String()
}




type DeleteDomainRequest struct {
  // {"en":"Return Chinese results for null (default)
  // En: Return the English prompt result", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"en":"Deleted domains
  // Use English half-width
  // semicolon between
  // two domains if there
  // are multiple to be
  // deleted.", "zh_CN":"删除的域名
  // 如果需要删除多个域名，用英文半角分号分隔。"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
}

func (s DeleteDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
  return s.String()
}

func (s *DeleteDomainRequest) SetLanguage(v string) *DeleteDomainRequest {
  s.Language = &v
  return s
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
  s.DomainName = &v
  return s
}

type DeleteDomainResponse struct {
  // {"en":"For more details of status
  // codes resCode, please refer to
  // “Status Codes of Dispatch
  // Business”.", "zh_CN":"状态码，resCode的详细说明请参见附录6： 业务状态码。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the
  // status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Detailed description of the
  // domain.
  // domainId Domain ID
  // domainName Domain name
  // ret Succeed or Fail tag
  // msg Succeed or Fail
  // messages
  // msgEn Succeed or Fail
  // messages", "zh_CN":"域名的详细说明。
  //   domainId 域名ID
  //   domainName 域名名称
  //   ret 成功或失败标识
  //   msg 成功或失败提示
  //   msgEn 成功或失败提示"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
  return s.String()
}

func (s *DeleteDomainResponse) SetResCode(v string) *DeleteDomainResponse {
  s.ResCode = &v
  return s
}

func (s *DeleteDomainResponse) SetMsg(v string) *DeleteDomainResponse {
  s.Msg = &v
  return s
}

func (s *DeleteDomainResponse) SetContent(v []*string) *DeleteDomainResponse {
  s.Content = v
  return s
}

type DeleteDomainPaths struct {
}

func (s DeleteDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainPaths) GoString() string {
  return s.String()
}

type DeleteDomainParameters struct {
}

func (s DeleteDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainParameters) GoString() string {
  return s.String()
}

type DeleteDomainRequestHeader struct {
}

func (s DeleteDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainRequestHeader) GoString() string {
  return s.String()
}

type DeleteDomainResponseHeader struct {
}

func (s DeleteDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainResponseHeader) GoString() string {
  return s.String()
}




type EditDomainGroupRequest struct {
  // {"en":"domain group ID", "zh_CN":"域名组ID"}
  DomainGroupId *string `json:"domainGroupId,omitempty" xml:"domainGroupId,omitempty" require:"true"`
  // {"en":"Domain group name. Only Chinese characters/English alphabet/numbers/underscores are supported, and case insensitive.Up to 32 characters.", "zh_CN":"域名组名称，仅支持输入中文/英文/数字/下划线，不区分大小写，最多可传32个字符。"}
  DomainGroupName *string `json:"domainGroupName,omitempty" xml:"domainGroupName,omitempty" require:"true"`
  // {"en":"The serviceType that the domain group belongs to", "zh_CN":"域名组所属加速服务类型"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"Domains associated with domain group The domain needs to be under the serviceType", "zh_CN":"域名组关联的域名,域名需要是传入加速服务类型下的域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s EditDomainGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s EditDomainGroupRequest) GoString() string {
  return s.String()
}

func (s *EditDomainGroupRequest) SetDomainGroupId(v string) *EditDomainGroupRequest {
  s.DomainGroupId = &v
  return s
}

func (s *EditDomainGroupRequest) SetDomainGroupName(v string) *EditDomainGroupRequest {
  s.DomainGroupName = &v
  return s
}

func (s *EditDomainGroupRequest) SetServiceType(v string) *EditDomainGroupRequest {
  s.ServiceType = &v
  return s
}

func (s *EditDomainGroupRequest) SetDomainList(v []*string) *EditDomainGroupRequest {
  s.DomainList = v
  return s
}

type EditDomainGroupResponse struct {
  // {"en":"Status Code", "zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditDomainGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s EditDomainGroupResponse) GoString() string {
  return s.String()
}

func (s *EditDomainGroupResponse) SetCode(v string) *EditDomainGroupResponse {
  s.Code = &v
  return s
}

func (s *EditDomainGroupResponse) SetMessage(v string) *EditDomainGroupResponse {
  s.Message = &v
  return s
}

type EditDomainGroupPaths struct {
}

func (s EditDomainGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s EditDomainGroupPaths) GoString() string {
  return s.String()
}

type EditDomainGroupParameters struct {
}

func (s EditDomainGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s EditDomainGroupParameters) GoString() string {
  return s.String()
}

type EditDomainGroupRequestHeader struct {
}

func (s EditDomainGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainGroupRequestHeader) GoString() string {
  return s.String()
}

type EditDomainGroupResponseHeader struct {
}

func (s EditDomainGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainGroupResponseHeader) GoString() string {
  return s.String()
}




type ChannelAcceTypeRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope. 
  // 2)With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3)If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"The response format:
  // 1)optional values:xml, json.
  // 2)'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
}

func (s ChannelAcceTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeRequest) GoString() string {
  return s.String()
}

func (s *ChannelAcceTypeRequest) SetCust(v string) *ChannelAcceTypeRequest {
  s.Cust = &v
  return s
}

func (s *ChannelAcceTypeRequest) SetStartdate(v string) *ChannelAcceTypeRequest {
  s.Startdate = &v
  return s
}

func (s *ChannelAcceTypeRequest) SetEnddate(v string) *ChannelAcceTypeRequest {
  s.Enddate = &v
  return s
}

func (s *ChannelAcceTypeRequest) SetDataformat(v string) *ChannelAcceTypeRequest {
  s.Dataformat = &v
  return s
}

type ChannelAcceTypeResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *ChannelAcceTypeChannelAcceTypeResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s ChannelAcceTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeResponse) GoString() string {
  return s.String()
}

func (s *ChannelAcceTypeResponse) SetProvider(v *ChannelAcceTypeChannelAcceTypeResponseProvider) *ChannelAcceTypeResponse {
  s.Provider = v
  return s
}

type ChannelAcceTypeChannelAcceTypeResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'account', 'zh_CN':'账号数据'}
  Account *ChannelAcceTypeChannelAcceTypeResponseProviderAccount `json:"account,omitempty" xml:"account,omitempty" require:"true" type:"Struct"`
}

func (s ChannelAcceTypeChannelAcceTypeResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeChannelAcceTypeResponseProvider) GoString() string {
  return s.String()
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProvider) SetName(v string) *ChannelAcceTypeChannelAcceTypeResponseProvider {
  s.Name = &v
  return s
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProvider) SetType(v string) *ChannelAcceTypeChannelAcceTypeResponseProvider {
  s.Type = &v
  return s
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProvider) SetAccount(v *ChannelAcceTypeChannelAcceTypeResponseProviderAccount) *ChannelAcceTypeChannelAcceTypeResponseProvider {
  s.Account = v
  return s
}

type ChannelAcceTypeChannelAcceTypeResponseProviderAccount struct {
  // {'en':'login-name', 'zh_CN':'日期'}
  LoginName *string `json:"login-name,omitempty" xml:"login-name,omitempty" require:"true"`
  // {'en':'acce-type', 'zh_CN':'频道'}
  AcceType *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType `json:"acce-type,omitempty" xml:"acce-type,omitempty" require:"true" type:"Struct"`
}

func (s ChannelAcceTypeChannelAcceTypeResponseProviderAccount) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeChannelAcceTypeResponseProviderAccount) GoString() string {
  return s.String()
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProviderAccount) SetLoginName(v string) *ChannelAcceTypeChannelAcceTypeResponseProviderAccount {
  s.LoginName = &v
  return s
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProviderAccount) SetAcceType(v *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType) *ChannelAcceTypeChannelAcceTypeResponseProviderAccount {
  s.AcceType = v
  return s
}

type ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType struct {
  // {'en':'name', 'zh_CN':'加速类型'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'value', 'zh_CN':'加速类型值'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel []*ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Repeated"`
}

func (s ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType) GoString() string {
  return s.String()
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType) SetName(v string) *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType {
  s.Name = &v
  return s
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType) SetValue(v string) *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType {
  s.Value = &v
  return s
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType) SetChannel(v []*ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel) *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceType {
  s.Channel = v
  return s
}

type ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel struct     {
  // {'en':'channel', 'zh_CN':'频道'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel) GoString() string {
  return s.String()
}

func (s *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel) SetText(v string) *ChannelAcceTypeChannelAcceTypeResponseProviderAccountAcceTypeChannel {
  s.Text = &v
  return s
}

type ChannelAcceTypePaths struct {
}

func (s ChannelAcceTypePaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypePaths) GoString() string {
  return s.String()
}

type ChannelAcceTypeParameters struct {
}

func (s ChannelAcceTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeParameters) GoString() string {
  return s.String()
}

type ChannelAcceTypeRequestHeader struct {
}

func (s ChannelAcceTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeRequestHeader) GoString() string {
  return s.String()
}

type ChannelAcceTypeResponseHeader struct {
}

func (s ChannelAcceTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelAcceTypeResponseHeader) GoString() string {
  return s.String()
}




type CancelApiDomainServiceRequest struct {
}

func (s CancelApiDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s CancelApiDomainServiceRequest) GoString() string {
  return s.String()
}

type CancelApiDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s CancelApiDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s CancelApiDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *CancelApiDomainServiceResponse) SetCode(v string) *CancelApiDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *CancelApiDomainServiceResponse) SetMessage(v string) *CancelApiDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *CancelApiDomainServiceResponse) SetHttpStatus(v int) *CancelApiDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *CancelApiDomainServiceResponse) SetXCncRequestId(v string) *CancelApiDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type CancelApiDomainServicePaths struct {
  // {"en":"", "zh_CN":"加速域名在系统中对应的ID
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s CancelApiDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s CancelApiDomainServicePaths) GoString() string {
  return s.String()
}

func (s *CancelApiDomainServicePaths) SetDomainId(v int) *CancelApiDomainServicePaths {
  s.DomainId = &v
  return s
}

type CancelApiDomainServiceParameters struct {
}

func (s CancelApiDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s CancelApiDomainServiceParameters) GoString() string {
  return s.String()
}

type CancelApiDomainServiceRequestHeader struct {
}

func (s CancelApiDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CancelApiDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type CancelApiDomainServiceResponseHeader struct {
}

func (s CancelApiDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CancelApiDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type EnableSingleDomainServiceRequest struct {
}

func (s EnableSingleDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceRequest) GoString() string {
  return s.String()
}

type EnableSingleDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s EnableSingleDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *EnableSingleDomainServiceResponse) SetCode(v string) *EnableSingleDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *EnableSingleDomainServiceResponse) SetMessage(v string) *EnableSingleDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *EnableSingleDomainServiceResponse) SetHttpStatus(v int) *EnableSingleDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *EnableSingleDomainServiceResponse) SetXCncRequestId(v string) *EnableSingleDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type EnableSingleDomainServicePaths struct {
  // {"en":"Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domain-id
  // 2. After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domain-id of the domain name; You can also query domain-id through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名在系统中对应的ID
  // 注意：
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s EnableSingleDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServicePaths) GoString() string {
  return s.String()
}

func (s *EnableSingleDomainServicePaths) SetDomainId(v int) *EnableSingleDomainServicePaths {
  s.DomainId = &v
  return s
}

type EnableSingleDomainServiceParameters struct {
}

func (s EnableSingleDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceParameters) GoString() string {
  return s.String()
}

type EnableSingleDomainServiceRequestHeader struct {
}

func (s EnableSingleDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type EnableSingleDomainServiceResponseHeader struct {
}

func (s EnableSingleDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type EnableDisableHwDomainRequest struct {
  // {"en":"state.online/offline", "zh_CN":"状态.online：启用.offline：停用。"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s EnableDisableHwDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainRequest) GoString() string {
  return s.String()
}

func (s *EnableDisableHwDomainRequest) SetState(v string) *EnableDisableHwDomainRequest {
  s.State = &v
  return s
}

type EnableDisableHwDomainResponse struct {
  // {"en":"The error code, when HTTPStatus is not 201, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为201时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The response data", "zh_CN":"响应数据"}
  Data *EnableDisableHwDomainEnableDisableHwDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id,omitempty" require:"true"`
}

func (s EnableDisableHwDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainResponse) GoString() string {
  return s.String()
}

func (s *EnableDisableHwDomainResponse) SetCode(v string) *EnableDisableHwDomainResponse {
  s.Code = &v
  return s
}

func (s *EnableDisableHwDomainResponse) SetMessage(v string) *EnableDisableHwDomainResponse {
  s.Message = &v
  return s
}

func (s *EnableDisableHwDomainResponse) SetData(v *EnableDisableHwDomainEnableDisableHwDomainResponseData) *EnableDisableHwDomainResponse {
  s.Data = v
  return s
}

func (s *EnableDisableHwDomainResponse) SetXRequestId(v string) *EnableDisableHwDomainResponse {
  s.XRequestId = &v
  return s
}

type EnableDisableHwDomainEnableDisableHwDomainResponseData struct {
  // {"en":"task id.", "zh_CN":"任务id"}
  TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s EnableDisableHwDomainEnableDisableHwDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainEnableDisableHwDomainResponseData) GoString() string {
  return s.String()
}

func (s *EnableDisableHwDomainEnableDisableHwDomainResponseData) SetTaskId(v string) *EnableDisableHwDomainEnableDisableHwDomainResponseData {
  s.TaskId = &v
  return s
}

type EnableDisableHwDomainPaths struct {
  // {"en":"The domain name for the acceleration domain to be deleted", "zh_CN":"要删除的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s EnableDisableHwDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainPaths) GoString() string {
  return s.String()
}

func (s *EnableDisableHwDomainPaths) SetDomain(v string) *EnableDisableHwDomainPaths {
  s.Domain = &v
  return s
}

type EnableDisableHwDomainParameters struct {
}

func (s EnableDisableHwDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainParameters) GoString() string {
  return s.String()
}

type EnableDisableHwDomainRequestHeader struct {
}

func (s EnableDisableHwDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainRequestHeader) GoString() string {
  return s.String()
}

type EnableDisableHwDomainResponseHeader struct {
}

func (s EnableDisableHwDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableDisableHwDomainResponseHeader) GoString() string {
  return s.String()
}




type DisableCdnDomainServiceRequest struct {
}

func (s DisableCdnDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableCdnDomainServiceRequest) GoString() string {
  return s.String()
}

type DisableCdnDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DisableCdnDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableCdnDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DisableCdnDomainServiceResponse) SetCode(v string) *DisableCdnDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DisableCdnDomainServiceResponse) SetMessage(v string) *DisableCdnDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DisableCdnDomainServiceResponse) SetHttpStatus(v int) *DisableCdnDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DisableCdnDomainServiceResponse) SetXCncRequestId(v string) *DisableCdnDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type DisableCdnDomainServicePaths struct {
  // {"en":"", "zh_CN":"加速域名在系统中对应的ID
  // 1. 参看请求示例中的url，123344对应的就是domainId
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domainId"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DisableCdnDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DisableCdnDomainServicePaths) GoString() string {
  return s.String()
}

func (s *DisableCdnDomainServicePaths) SetDomainId(v int) *DisableCdnDomainServicePaths {
  s.DomainId = &v
  return s
}

type DisableCdnDomainServiceParameters struct {
}

func (s DisableCdnDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableCdnDomainServiceParameters) GoString() string {
  return s.String()
}

type DisableCdnDomainServiceRequestHeader struct {
}

func (s DisableCdnDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableCdnDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type DisableCdnDomainServiceResponseHeader struct {
}

func (s DisableCdnDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableCdnDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type DisableSingleDomainServiceRequest struct {
}

func (s DisableSingleDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceRequest) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DisableSingleDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DisableSingleDomainServiceResponse) SetCode(v string) *DisableSingleDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetMessage(v string) *DisableSingleDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetHttpStatus(v int) *DisableSingleDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetXCncRequestId(v string) *DisableSingleDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type DisableSingleDomainServicePaths struct {
  // {"en":"", "zh_CN":"加速域名在系统中对应的ID
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DisableSingleDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServicePaths) GoString() string {
  return s.String()
}

func (s *DisableSingleDomainServicePaths) SetDomainId(v int) *DisableSingleDomainServicePaths {
  s.DomainId = &v
  return s
}

type DisableSingleDomainServiceParameters struct {
}

func (s DisableSingleDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceParameters) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceRequestHeader struct {
}

func (s DisableSingleDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceResponseHeader struct {
}

func (s DisableSingleDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type DeleteApiDomainServiceRequest struct {
  // {"en":"Accelerated domain name, choose one from domain-id. Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domain-id
  // 2、After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domain-id of the domain name; You can also query domain-id through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名与domain-id二选一。
  // domain-id：加速域名在系统中对应的ID
  // domain-name：加速的域名
  // 注意：
  // 1、参看请求示例中的url，123344对应的就是domain-id
  // 2、创建域名成功提交后，返回参数中的location访问url中，能够查询到域名对应的domain-id；也可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Accelerated domain name, choose one from domain-id. Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domain-id
  // 2、After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domain-id of the domain name; You can also query domain-id through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名与domain-id二选一。
  // domain-id：加速域名在系统中对应的ID
  // domain-name：加速的域名
  // 注意：
  // 1、参看请求示例中的url，123344对应的就是domain-id
  // 2、创建域名成功提交后，返回参数中的location访问url中，能够查询到域名对应的domain-id；也可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DeleteApiDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *DeleteApiDomainServiceRequest) SetDomainName(v string) *DeleteApiDomainServiceRequest {
  s.DomainName = &v
  return s
}

func (s *DeleteApiDomainServiceRequest) SetDomainId(v string) *DeleteApiDomainServiceRequest {
  s.DomainId = &v
  return s
}

type DeleteApiDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DeleteApiDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DeleteApiDomainServiceResponse) SetCode(v string) *DeleteApiDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DeleteApiDomainServiceResponse) SetMessage(v string) *DeleteApiDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DeleteApiDomainServiceResponse) SetHttpStatus(v int) *DeleteApiDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DeleteApiDomainServiceResponse) SetXCncRequestId(v string) *DeleteApiDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type DeleteApiDomainServicePaths struct {
  // {"en":"", "zh_CN":"域名名称或域名id，在请求的url后面"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s DeleteApiDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServicePaths) GoString() string {
  return s.String()
}

func (s *DeleteApiDomainServicePaths) SetDomainName(v string) *DeleteApiDomainServicePaths {
  s.DomainName = &v
  return s
}

type DeleteApiDomainServiceParameters struct {
}

func (s DeleteApiDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceParameters) GoString() string {
  return s.String()
}

type DeleteApiDomainServiceRequestHeader struct {
}

func (s DeleteApiDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type DeleteApiDomainServiceResponseHeader struct {
}

func (s DeleteApiDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type DeleteHwDomainRequest struct {
}

func (s DeleteHwDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainRequest) GoString() string {
  return s.String()
}

type DeleteHwDomainResponse struct {
  // {"en":"The error code, when HTTPStatus is not 201, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为201时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The response data", "zh_CN":"响应数据"}
  Data *DeleteHwDomainDeleteHwDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id,omitempty" require:"true"`
}

func (s DeleteHwDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainResponse) GoString() string {
  return s.String()
}

func (s *DeleteHwDomainResponse) SetCode(v string) *DeleteHwDomainResponse {
  s.Code = &v
  return s
}

func (s *DeleteHwDomainResponse) SetMessage(v string) *DeleteHwDomainResponse {
  s.Message = &v
  return s
}

func (s *DeleteHwDomainResponse) SetData(v *DeleteHwDomainDeleteHwDomainResponseData) *DeleteHwDomainResponse {
  s.Data = v
  return s
}

func (s *DeleteHwDomainResponse) SetXRequestId(v string) *DeleteHwDomainResponse {
  s.XRequestId = &v
  return s
}

type DeleteHwDomainDeleteHwDomainResponseData struct {
  // {"en":"task id.", "zh_CN":"任务id"}
  TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s DeleteHwDomainDeleteHwDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainDeleteHwDomainResponseData) GoString() string {
  return s.String()
}

func (s *DeleteHwDomainDeleteHwDomainResponseData) SetTaskId(v string) *DeleteHwDomainDeleteHwDomainResponseData {
  s.TaskId = &v
  return s
}

type DeleteHwDomainPaths struct {
  // {"en":"The domain name for the acceleration domain to be deleted", "zh_CN":"要删除的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s DeleteHwDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainPaths) GoString() string {
  return s.String()
}

func (s *DeleteHwDomainPaths) SetDomain(v string) *DeleteHwDomainPaths {
  s.Domain = &v
  return s
}

type DeleteHwDomainParameters struct {
}

func (s DeleteHwDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainParameters) GoString() string {
  return s.String()
}

type DeleteHwDomainRequestHeader struct {
}

func (s DeleteHwDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainRequestHeader) GoString() string {
  return s.String()
}

type DeleteHwDomainResponseHeader struct {
}

func (s DeleteHwDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteHwDomainResponseHeader) GoString() string {
  return s.String()
}




type EnableCdnDomainServiceRequest struct {
}

func (s EnableCdnDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableCdnDomainServiceRequest) GoString() string {
  return s.String()
}

type EnableCdnDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s EnableCdnDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableCdnDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *EnableCdnDomainServiceResponse) SetCode(v string) *EnableCdnDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *EnableCdnDomainServiceResponse) SetMessage(v string) *EnableCdnDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *EnableCdnDomainServiceResponse) SetHttpStatus(v int) *EnableCdnDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *EnableCdnDomainServiceResponse) SetXCncRequestId(v string) *EnableCdnDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type EnableCdnDomainServicePaths struct {
  // {"en":"Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domainId
  // 2. After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domainId of the domain name; You can also query domainId through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名在系统中对应的ID
  // 注意：
  // 1. 参看请求示例中的url，123344对应的就是domainId
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domainId"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s EnableCdnDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s EnableCdnDomainServicePaths) GoString() string {
  return s.String()
}

func (s *EnableCdnDomainServicePaths) SetDomainId(v int) *EnableCdnDomainServicePaths {
  s.DomainId = &v
  return s
}

type EnableCdnDomainServiceParameters struct {
}

func (s EnableCdnDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableCdnDomainServiceParameters) GoString() string {
  return s.String()
}

type EnableCdnDomainServiceRequestHeader struct {
}

func (s EnableCdnDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableCdnDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type EnableCdnDomainServiceResponseHeader struct {
}

func (s EnableCdnDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableCdnDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type UpdateChangeServerRequest struct {
  // {"en":"Change servers configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <change-servers/>", "zh_CN":"【接入域名跳转】
  // 注意：
  // 1、需要取消【接入域名跳转】时，可以传入空节点<change-servers></change-servers>。
  // 2、表示需要设置【接入域名跳转】，此项必填"}
  ChangeServers []*UpdateChangeServerUpdateChangeServerRequestChangeServers `json:"change-servers,omitempty" xml:"change-servers,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateChangeServerRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerRequest) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerRequest) SetChangeServers(v []*UpdateChangeServerUpdateChangeServerRequestChangeServers) *UpdateChangeServerRequest {
  s.ChangeServers = v
  return s
}

type UpdateChangeServerUpdateChangeServerRequestChangeServers struct     {
  // {"en":"If it is a universal domain name, set it to a universal domain name, for example, *.56.com.", "zh_CN":"如果是泛域名，需要填写为泛域名，例如：*.56.com"}
  TargetServer *string `json:"target-server,omitempty" xml:"target-server,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: 
  // A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. 
  // B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. 
  // C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. 
  // D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. 
  // E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", 
  //       "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： 
  // a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； 
  // b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； 
  // c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； 
  // d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； 
  // e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。
  // "}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s UpdateChangeServerUpdateChangeServerRequestChangeServers) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerUpdateChangeServerRequestChangeServers) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerUpdateChangeServerRequestChangeServers) SetTargetServer(v string) *UpdateChangeServerUpdateChangeServerRequestChangeServers {
  s.TargetServer = &v
  return s
}

func (s *UpdateChangeServerUpdateChangeServerRequestChangeServers) SetDataId(v int) *UpdateChangeServerUpdateChangeServerRequestChangeServers {
  s.DataId = &v
  return s
}

type UpdateChangeServerResponse struct {
  // {"en":"The error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data.", "zh_CN":"返回数据体。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s UpdateChangeServerResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerResponse) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerResponse) SetCode(v string) *UpdateChangeServerResponse {
  s.Code = &v
  return s
}

func (s *UpdateChangeServerResponse) SetMessage(v string) *UpdateChangeServerResponse {
  s.Message = &v
  return s
}

func (s *UpdateChangeServerResponse) SetData(v string) *UpdateChangeServerResponse {
  s.Data = &v
  return s
}

func (s *UpdateChangeServerResponse) SetXCncRequestId(v string) *UpdateChangeServerResponse {
  s.XCncRequestId = &v
  return s
}

type UpdateChangeServerPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s UpdateChangeServerPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerPaths) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerPaths) SetDomainName(v string) *UpdateChangeServerPaths {
  s.DomainName = &v
  return s
}

type UpdateChangeServerParameters struct {
}

func (s UpdateChangeServerParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerParameters) GoString() string {
  return s.String()
}

type UpdateChangeServerRequestHeader struct {
}

func (s UpdateChangeServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerRequestHeader) GoString() string {
  return s.String()
}

type UpdateChangeServerResponseHeader struct {
}

func (s UpdateChangeServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerResponseHeader) GoString() string {
  return s.String()
}




type UpdateDispatchDomainRequest struct {
  // {"en":"domainName", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"domainId", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"dispatchZone", "zh_CN":"可选，国内：cdngtm.cn， 海外：cdngtm.com；默认值为cdngtm.com"}
  DispatchZone *string `json:"dispatchZone,omitempty" xml:"dispatchZone,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"默认值为10，不得低于10，否则返回解析记录TTL错误提示。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty"`
  // {"en":"Desc", "zh_CN":"备注"}
  DomainDesc *string `json:"domainDesc,omitempty" xml:"domainDesc,omitempty"`
  // {"en":"sessionHold", "zh_CN":"1：开启  空为不开启"}
  SessionHold *int `json:"sessionHold,omitempty" xml:"sessionHold,omitempty"`
  // {"en":"sessionInterval", "zh_CN":"保持时间"}
  SessionInterval *int `json:"sessionInterval,omitempty" xml:"sessionInterval,omitempty"`
  // {"en":"language", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s UpdateDispatchDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateDispatchDomainRequest) GoString() string {
  return s.String()
}

func (s *UpdateDispatchDomainRequest) SetDomainName(v string) *UpdateDispatchDomainRequest {
  s.DomainName = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetDomainId(v int) *UpdateDispatchDomainRequest {
  s.DomainId = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetDispatchZone(v string) *UpdateDispatchDomainRequest {
  s.DispatchZone = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetTtl(v int) *UpdateDispatchDomainRequest {
  s.Ttl = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetDomainDesc(v string) *UpdateDispatchDomainRequest {
  s.DomainDesc = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetSessionHold(v int) *UpdateDispatchDomainRequest {
  s.SessionHold = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetSessionInterval(v int) *UpdateDispatchDomainRequest {
  s.SessionInterval = &v
  return s
}

func (s *UpdateDispatchDomainRequest) SetLanguage(v string) *UpdateDispatchDomainRequest {
  s.Language = &v
  return s
}

type UpdateDispatchDomainResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"detail", "zh_CN":"详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateDispatchDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateDispatchDomainResponse) GoString() string {
  return s.String()
}

func (s *UpdateDispatchDomainResponse) SetResCode(v string) *UpdateDispatchDomainResponse {
  s.ResCode = &v
  return s
}

func (s *UpdateDispatchDomainResponse) SetMsg(v string) *UpdateDispatchDomainResponse {
  s.Msg = &v
  return s
}

type UpdateDispatchDomainPaths struct {
}

func (s UpdateDispatchDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateDispatchDomainPaths) GoString() string {
  return s.String()
}

type UpdateDispatchDomainParameters struct {
}

func (s UpdateDispatchDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateDispatchDomainParameters) GoString() string {
  return s.String()
}

type UpdateDispatchDomainRequestHeader struct {
}

func (s UpdateDispatchDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDispatchDomainRequestHeader) GoString() string {
  return s.String()
}

type UpdateDispatchDomainResponseHeader struct {
}

func (s UpdateDispatchDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDispatchDomainResponseHeader) GoString() string {
  return s.String()
}




type GetFuzzyPagingDomainListRequest struct {
  // {"en":"Page number must be a positive integer greater than 0", "zh_CN":"分页的页码，必须为大于0的正整数"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Number of domain name data items for paging, must be a positive integer greater than 0", "zh_CN":"分页的域名数据条数，必须大于0的正整数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Specifies the service type of the query, only one type per query, and no default lookup for all types", "zh_CN":"指定查询的服务类型，每次查询只能传一个类型，不传默认查全部类型"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Specifies the accelerated domain name for the query, allows multiple domains, commas delimited, and no default lookup of all domain names", "zh_CN":"指定查询的加速域名，允许多个域名，逗号分隔，不传默认查全部域名"}
  DomainName []*string `json:"domainName,omitempty" xml:"domainName,omitempty" type:"Repeated"`
  // {"en":"Query to accelerated domain name, optional value is: fuzzy_match for fuzzy query; Full_match represents an exact query
  // No fuzzy_match by default, for accelerated domain name only", "zh_CN":"查询加速域名的方式，可选值为：fuzzy_match表示模糊查询；full_match表示精确查询
  // 不传默认为fuzzy_match，仅针对加速域名"}
  QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
  // {"en":"Query start time, support for years, months, days, hours, minutes, and seconds, for example: 20170101.09 million. Time equals", "zh_CN":"查询开始时间，支持范围为年月日时分秒，例如：20170101090000。时间含等于"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"Query end time, query time within the existence of the accelerated domain name, time is equal to, do not pass the default query all", "zh_CN":"查询结束时间，查询时间段内存在的加速域名，时间含等于，不传默认查询所有"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"Accelerate the status of the domain name, enabled indicates that it is in effect; Disabled indicates that it is Disabled; Deploying means in the process of deployment; Checking indicates that the audit is in progress; Disabling: Indicates disabled, no default lookup for all", "zh_CN":"加速域名的状态，enabled表示已生效；disabled表示已禁用；deploying表示部署中；checking表示审核中；disabling:表示禁用中，不传默认查全部"}
  DomainStatus *string `json:"domainStatus,omitempty" xml:"domainStatus,omitempty"`
}

func (s GetFuzzyPagingDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListRequest) GoString() string {
  return s.String()
}

func (s *GetFuzzyPagingDomainListRequest) SetPageNumber(v int) *GetFuzzyPagingDomainListRequest {
  s.PageNumber = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetPageSize(v int) *GetFuzzyPagingDomainListRequest {
  s.PageSize = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetServiceType(v string) *GetFuzzyPagingDomainListRequest {
  s.ServiceType = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetDomainName(v []*string) *GetFuzzyPagingDomainListRequest {
  s.DomainName = v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetQueryType(v string) *GetFuzzyPagingDomainListRequest {
  s.QueryType = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetStartTime(v string) *GetFuzzyPagingDomainListRequest {
  s.StartTime = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetEndTime(v string) *GetFuzzyPagingDomainListRequest {
  s.EndTime = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetDomainStatus(v string) *GetFuzzyPagingDomainListRequest {
  s.DomainStatus = &v
  return s
}

type GetFuzzyPagingDomainListResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Responses the page number of the data", "zh_CN":"所有满足条件的数据总条数"}
  TotalCount *int `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"total pages", "zh_CN":"总页数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
  // {"en":"Responses the page number of the data", "zh_CN":"返回数据的页码"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Number of data page", "zh_CN":"每个页面的数据条数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Responses status information for the accelerated domain name", "zh_CN":"返回加速域名的状态信息"}
  ResultList []*GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList `json:"resultList,omitempty" xml:"resultList,omitempty" require:"true" type:"Repeated"`
}

func (s GetFuzzyPagingDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponse) GoString() string {
  return s.String()
}

func (s *GetFuzzyPagingDomainListResponse) SetCode(v int) *GetFuzzyPagingDomainListResponse {
  s.Code = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetXCncRequestId(v string) *GetFuzzyPagingDomainListResponse {
  s.XCncRequestId = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetTotalCount(v int) *GetFuzzyPagingDomainListResponse {
  s.TotalCount = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetTotalPageNumber(v int) *GetFuzzyPagingDomainListResponse {
  s.TotalPageNumber = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetPageNumber(v int) *GetFuzzyPagingDomainListResponse {
  s.PageNumber = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetPageSize(v int) *GetFuzzyPagingDomainListResponse {
  s.PageSize = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetResultList(v []*GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) *GetFuzzyPagingDomainListResponse {
  s.ResultList = v
  return s
}

type GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList struct     {
  // {"en":"Accelerated domain CNAME corresponding to CNAME, for example: 7nt6mrh7sdkslj.cdn30.com", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Configuration name", "zh_CN":"配置单名称"}
  ConfigFormName *string `json:"configFormName,omitempty" xml:"configFormName,omitempty" require:"true"`
  // {"en":"The time format is: 20160323112310", "zh_CN":"时间格式为：20160323112310"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Corresponding domain ID", "zh_CN":"对应的域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Operator of this query", "zh_CN":"本次查询的操作者"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
  // {"en":"Accelerate the origin IP of a domain name", "zh_CN":"加速域名的回源IP"}
  OriginIps *string `json:"originIps,omitempty" xml:"originIps,omitempty" require:"true"`
  // {"en":"Service type for accelerated domain name", "zh_CN":"加速域名的服务类型"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"Status of accelerated domain name.", "zh_CN":"加速域名的状态"}
  DomainStatus *string `json:"domainStatus,omitempty" xml:"domainStatus,omitempty" require:"true"`
  // {"en":"Deployment version code", "zh_CN":"部署版本号"}
  DeployVersion *string `json:"deployVersion,omitempty" xml:"deployVersion,omitempty" require:"true"`
  // {"en":"Does the domain name enable CDN acceleration services, Y and N?", "zh_CN":"域名是否启用CDN加速服务，Y和N"}
  CdnServiceStatus *string `json:"cdnServiceStatus,omitempty" xml:"cdnServiceStatus,omitempty" require:"true"`
  // {"en":"Whether the accelerated domain name is enabled, Y and N?", "zh_CN":"加速域名是否启用，Y和N"}
  IsEnabled *string `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) GoString() string {
  return s.String()
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetCname(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.Cname = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetConfigFormName(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.ConfigFormName = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetCreateTime(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.CreateTime = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetDomainId(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.DomainId = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetDomainName(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.DomainName = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetOperator(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.Operator = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetOriginIps(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.OriginIps = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetServiceType(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.ServiceType = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetDomainStatus(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.DomainStatus = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetDeployVersion(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.DeployVersion = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetCdnServiceStatus(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.CdnServiceStatus = &v
  return s
}

func (s *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList) SetIsEnabled(v string) *GetFuzzyPagingDomainListGetFuzzyPagingDomainListResponseResultList {
  s.IsEnabled = &v
  return s
}

type GetFuzzyPagingDomainListPaths struct {
}

func (s GetFuzzyPagingDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListPaths) GoString() string {
  return s.String()
}

type GetFuzzyPagingDomainListParameters struct {
}

func (s GetFuzzyPagingDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListParameters) GoString() string {
  return s.String()
}

type GetFuzzyPagingDomainListRequestHeader struct {
}

func (s GetFuzzyPagingDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListRequestHeader) GoString() string {
  return s.String()
}

type GetFuzzyPagingDomainListResponseHeader struct {
}

func (s GetFuzzyPagingDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponseHeader) GoString() string {
  return s.String()
}




type QueryCustomerAnycastIPForWplusRequest struct {
}

func (s QueryCustomerAnycastIPForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusRequest) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusResponse struct {
  // {"en":"code", "zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"anycastIp详细"}
  Data []*QueryCustomerAnycastIPForWplusAnycastIPDetail `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCustomerAnycastIPForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusResponse) GoString() string {
  return s.String()
}

func (s *QueryCustomerAnycastIPForWplusResponse) SetCode(v string) *QueryCustomerAnycastIPForWplusResponse {
  s.Code = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusResponse) SetMessage(v string) *QueryCustomerAnycastIPForWplusResponse {
  s.Message = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusResponse) SetData(v []*QueryCustomerAnycastIPForWplusAnycastIPDetail) *QueryCustomerAnycastIPForWplusResponse {
  s.Data = v
  return s
}

type QueryCustomerAnycastIPForWplusAnycastIPDetail struct {
  // {"en":"IP", "zh_CN":"IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"域名加速状态: USED, AVAILABLE"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Record Status, For example, lock or unlock; data of length 1 or 2 can be passed.", "zh_CN":"记录状态，例如锁定、非锁定等，可以传长度为1或2的数据"}
  RecordStatus *string `json:"recordStatus,omitempty" xml:"recordStatus,omitempty" require:"true"`
  // {"en":"Is china mainland, 0:NO,1: YES", "zh_CN":"是否属于大陆，0：否，1：是"}
  IsChinaMainland *string `json:"isChinaMainland,omitempty" xml:"isChinaMainland,omitempty" require:"true"`
}

func (s QueryCustomerAnycastIPForWplusAnycastIPDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusAnycastIPDetail) GoString() string {
  return s.String()
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetIp(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.Ip = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetStatus(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.Status = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetRecordStatus(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.RecordStatus = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetIsChinaMainland(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.IsChinaMainland = &v
  return s
}

type QueryCustomerAnycastIPForWplusPaths struct {
}

func (s QueryCustomerAnycastIPForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusPaths) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusParameters struct {
}

func (s QueryCustomerAnycastIPForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusParameters) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusRequestHeader struct {
}

func (s QueryCustomerAnycastIPForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusRequestHeader) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusResponseHeader struct {
}

func (s QueryCustomerAnycastIPForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainsRequest struct {
  // {"en":"If no domain entered, it
  // means all domains and
  // domains details under
  // the user are queried
  // and returned.", "zh_CN":"如果没有填写域名，则返回该用户的所有域名及域名相应信息。"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
  // {"en":"If no domain entered, it
  // means all domains and
  // domains details under
  // the user are queried
  // and returned.", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"en":"Paging query page number (from 1) . If it is empty, the default is 1.", "zh_CN":"分页查询页码（从1开始），为空则默认为1"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Paging query number.If it is empty, the default is 10 thousand.", "zh_CN":"分页查询条数，为空默认为1万条"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainsRequest) SetDomainName(v string) *QueryDomainsRequest {
  s.DomainName = &v
  return s
}

func (s *QueryDomainsRequest) SetLanguage(v string) *QueryDomainsRequest {
  s.Language = &v
  return s
}

func (s *QueryDomainsRequest) SetPageIndex(v int) *QueryDomainsRequest {
  s.PageIndex = &v
  return s
}

func (s *QueryDomainsRequest) SetPageSize(v int) *QueryDomainsRequest {
  s.PageSize = &v
  return s
}

type QueryDomainsResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"domain Id", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Take-over status. Options: 1
  // means Already in take-over, 0
  // means not in take-over.", "zh_CN":"接管状态。其中：1表示已接管，0表示未接管。"}
  AdoptState []*string `json:"adoptState,omitempty" xml:"adoptState,omitempty" require:"true" type:"Repeated"`
  // {"en":"remark", "zh_CN":"备注"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"domain state: 2: normal, 3: stop, 8: lock, 9: stop and lock", "zh_CN":"域名状态: 2/正常,3/停止,8/正常锁定,9/停止锁定"}
  State *int `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"soa ttl: A 32-bit signed integer that specifies the time interval that the resource record may be cached before the source of the information should be consulted again. Zero values are interpreted to mean that the RR can only be used for the transaction in progress, and should not be cached. For example, SOA records are always distributed with a zero TTL to prohibit caching. Zero values can also be used for extremely volatile data.", "zh_CN":"soa ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"soa Serial: The unsigned 32-bit version number of the original copy of the zone.", "zh_CN":"soa序列号"}
  SerialNumber *int `json:"serialNumber,omitempty" xml:"serialNumber,omitempty" require:"true"`
  // {"en":"soa Contact: A domain name that specifies the mailbox of the person responsible for this zone.", "zh_CN":"soa邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
}

func (s QueryDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainsResponse) SetResCode(v string) *QueryDomainsResponse {
  s.ResCode = &v
  return s
}

func (s *QueryDomainsResponse) SetDomainId(v int) *QueryDomainsResponse {
  s.DomainId = &v
  return s
}

func (s *QueryDomainsResponse) SetDomainName(v string) *QueryDomainsResponse {
  s.DomainName = &v
  return s
}

func (s *QueryDomainsResponse) SetAdoptState(v []*string) *QueryDomainsResponse {
  s.AdoptState = v
  return s
}

func (s *QueryDomainsResponse) SetRemark(v string) *QueryDomainsResponse {
  s.Remark = &v
  return s
}

func (s *QueryDomainsResponse) SetState(v int) *QueryDomainsResponse {
  s.State = &v
  return s
}

func (s *QueryDomainsResponse) SetTtl(v int) *QueryDomainsResponse {
  s.Ttl = &v
  return s
}

func (s *QueryDomainsResponse) SetSerialNumber(v int) *QueryDomainsResponse {
  s.SerialNumber = &v
  return s
}

func (s *QueryDomainsResponse) SetEmail(v string) *QueryDomainsResponse {
  s.Email = &v
  return s
}

type QueryDomainsPaths struct {
}

func (s QueryDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsPaths) GoString() string {
  return s.String()
}

type QueryDomainsParameters struct {
}

func (s QueryDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsParameters) GoString() string {
  return s.String()
}

type QueryDomainsRequestHeader struct {
}

func (s QueryDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainsResponseHeader struct {
}

func (s QueryDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsResponseHeader) GoString() string {
  return s.String()
}




type QueryCustomerDomainNameGroupServiceRequest struct {
  // {'en':'Domain group name: the default value is all domain group', 'zh_CN':'域名组名称
  // 不传递则默认查询账号下全部域名组；'}
  DomainGroupNameList []*string `json:"domainGroupNameList,omitempty" xml:"domainGroupNameList,omitempty" type:"Repeated"`
}

func (s QueryCustomerDomainNameGroupServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryCustomerDomainNameGroupServiceRequest) SetDomainGroupNameList(v []*string) *QueryCustomerDomainNameGroupServiceRequest {
  s.DomainGroupNameList = v
  return s
}

type QueryCustomerDomainNameGroupServiceResponse struct {
  Result []*QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCustomerDomainNameGroupServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryCustomerDomainNameGroupServiceResponse) SetResult(v []*QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult) *QueryCustomerDomainNameGroupServiceResponse {
  s.Result = v
  return s
}

type QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult struct     {
  // {'en':'Domain group name', 'zh_CN':'域名组名称'}
  DomainGroupName *string `json:"domainGroupName,omitempty" xml:"domainGroupName,omitempty" require:"true"`
  // {'en':'Domain list', 'zh_CN':'域名列表'}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {'en':'Domain group ID', 'zh_CN':'域名组ID'}
  DomainGroupId *string `json:"domainGroupId,omitempty" xml:"domainGroupId,omitempty" require:"true"`
}

func (s QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult) SetDomainGroupName(v string) *QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult {
  s.DomainGroupName = &v
  return s
}

func (s *QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult) SetDomainList(v []*string) *QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult {
  s.DomainList = v
  return s
}

func (s *QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult) SetDomainGroupId(v string) *QueryCustomerDomainNameGroupServiceQueryCustomerDomainNameGroupServiceResponseResult {
  s.DomainGroupId = &v
  return s
}

type QueryCustomerDomainNameGroupServicePaths struct {
}

func (s QueryCustomerDomainNameGroupServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServicePaths) GoString() string {
  return s.String()
}

type QueryCustomerDomainNameGroupServiceParameters struct {
}

func (s QueryCustomerDomainNameGroupServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServiceParameters) GoString() string {
  return s.String()
}

type QueryCustomerDomainNameGroupServiceRequestHeader struct {
}

func (s QueryCustomerDomainNameGroupServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryCustomerDomainNameGroupServiceResponseHeader struct {
}

func (s QueryCustomerDomainNameGroupServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerDomainNameGroupServiceResponseHeader) GoString() string {
  return s.String()
}




type DelDispatchDomainRequest struct {
  // {"en":"Ids of the domains to be deleted Use English half-width semicolon between two domains if there are multiple to be deleted.", "zh_CN":"要删除的域名的Id
  // 如果需要删除多个域名，用英文半角分号分隔。"}
  DomainIds *string `json:"domainIds,omitempty" xml:"domainIds,omitempty"`
  // {"en":"Ids of the domains to be deleted Use English half-width semicolon between two domains if there are multiple to be deleted.", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s DelDispatchDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchDomainRequest) GoString() string {
  return s.String()
}

func (s *DelDispatchDomainRequest) SetDomainIds(v string) *DelDispatchDomainRequest {
  s.DomainIds = &v
  return s
}

func (s *DelDispatchDomainRequest) SetLanguage(v string) *DelDispatchDomainRequest {
  s.Language = &v
  return s
}

type DelDispatchDomainResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Detailed description of the domain. 
  // domainId:Domain ID 
  // domainName:Domain name 
  // dispatchCname:Dispatch CNAME", "zh_CN":"域名的详细说明。
  // domainId 域名ID标识
  // 
  // domainName 域名
  // 
  // dispatchCname 调度CNAME"}
  Content map[string]interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s DelDispatchDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchDomainResponse) GoString() string {
  return s.String()
}

func (s *DelDispatchDomainResponse) SetResCode(v string) *DelDispatchDomainResponse {
  s.ResCode = &v
  return s
}

func (s *DelDispatchDomainResponse) SetMsg(v string) *DelDispatchDomainResponse {
  s.Msg = &v
  return s
}

func (s *DelDispatchDomainResponse) SetContent(v map[string]interface{}) *DelDispatchDomainResponse {
  s.Content = v
  return s
}

type DelDispatchDomainPaths struct {
}

func (s DelDispatchDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchDomainPaths) GoString() string {
  return s.String()
}

type DelDispatchDomainParameters struct {
}

func (s DelDispatchDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchDomainParameters) GoString() string {
  return s.String()
}

type DelDispatchDomainRequestHeader struct {
}

func (s DelDispatchDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchDomainRequestHeader) GoString() string {
  return s.String()
}

type DelDispatchDomainResponseHeader struct {
}

func (s DelDispatchDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DelDispatchDomainResponseHeader) GoString() string {
  return s.String()
}




type ControlDispatchDomainRequest struct {
  // {"en":"domainIds", "zh_CN":"域名id 多个用英文分号隔开 ;"}
  DomainIds *string `json:"domainIds,omitempty" xml:"domainIds,omitempty" require:"true"`
  // {"en":"controlType", "zh_CN":"类型: 1 ：启动  2 ：停用"}
  ControlType *int `json:"controlType,omitempty" xml:"controlType,omitempty" require:"true"`
  // {"en":"language", "zh_CN":"为空返回中文结果(默认)"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ControlDispatchDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchDomainRequest) GoString() string {
  return s.String()
}

func (s *ControlDispatchDomainRequest) SetDomainIds(v string) *ControlDispatchDomainRequest {
  s.DomainIds = &v
  return s
}

func (s *ControlDispatchDomainRequest) SetControlType(v int) *ControlDispatchDomainRequest {
  s.ControlType = &v
  return s
}

func (s *ControlDispatchDomainRequest) SetLanguage(v string) *ControlDispatchDomainRequest {
  s.Language = &v
  return s
}

type ControlDispatchDomainResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"detail", "zh_CN":"详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"domainId调度域名id
  // 
  // code 授权处理结果代码请参见“附录1业务状态码”
  // status 0启用 1停用"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s ControlDispatchDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchDomainResponse) GoString() string {
  return s.String()
}

func (s *ControlDispatchDomainResponse) SetResCode(v string) *ControlDispatchDomainResponse {
  s.ResCode = &v
  return s
}

func (s *ControlDispatchDomainResponse) SetMsg(v string) *ControlDispatchDomainResponse {
  s.Msg = &v
  return s
}

func (s *ControlDispatchDomainResponse) SetContent(v []*string) *ControlDispatchDomainResponse {
  s.Content = v
  return s
}

type ControlDispatchDomainPaths struct {
}

func (s ControlDispatchDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchDomainPaths) GoString() string {
  return s.String()
}

type ControlDispatchDomainParameters struct {
}

func (s ControlDispatchDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchDomainParameters) GoString() string {
  return s.String()
}

type ControlDispatchDomainRequestHeader struct {
}

func (s ControlDispatchDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchDomainRequestHeader) GoString() string {
  return s.String()
}

type ControlDispatchDomainResponseHeader struct {
}

func (s ControlDispatchDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDispatchDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryApiDomainListServiceRequest struct {
}

func (s QueryApiDomainListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceRequest) GoString() string {
  return s.String()
}

type QueryApiDomainListServiceResponse struct {
  // {"en":"domain list", "zh_CN":"域名列表"}
  DomainList []*QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApiDomainListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceResponse) SetDomainList(v []*QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) *QueryApiDomainListServiceResponse {
  s.DomainList = v
  return s
}

type QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList struct     {
  // {"en":"Name of accelerated domain name", "zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"The corresponding domain name ID: the domain name ID, used to perform the query and modification operations of the related domain name.", "zh_CN":"对应的域名ID：域名ID，用于执行相关域名的查询、修改操作等。"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Accelerated domain CNAME corresponding to CNAME, for example: 7nt6mrh7sdkslj.cdn30.com", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Speed up the service type of the domain name, the value is:
  // Web/web-https: web acceleration / web acceleration - https
  // Wsa/wsa-https: Total Station Acceleration / Total Station Acceleration - https
  // Vodstream/vod-https: on-demand acceleration/on-demand acceleration - https
  // Download/dl-https: Download acceleration/download acceleration - https
  // livestream/live-https/cloudv-live: Live acceleration/Live acceleration - https/Cloud vedio for live
  // 1028: Content Acceleration;
  // 1115: Dynamic Web Acceleration;
  // 1369: Media Acceleration - RTMP
  // 1391: Download Acceleration
  // 1348: Media Acceleration Live Broadcast
  // 1551: Floodshield", "zh_CN":"加速域名的服务类型，取值：
  // web/web-https：网页加速/网页加速-https
  // wsa/wsa-https：全站加速/全站加速-https
  // vodstream/vod-https：点播加速/点播加速-https
  // download/dl-https：下载加速/下载加速-https
  // livestream/live-https/cloudv-live：直播加速/直播加速-https/云直播
  // appa/s-appa：应用加速/应用安全加速解决方案
  // 1028 : Content Acceleration;
  // 1115 : Dynamic Web Acceleration;
  // 1369 : Media Acceleration - RTMP
  // 1391 : Download Acceleration
  // 1348 : Media Acceleration Live Broadcast
  // 1551 : Floodshield"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
  // {"en":"The deployment status of the accelerated domain name: Deployed indicates that the accelerated domain name configuration is complete; InProgress indicates that the deployment task for this accelerated domain name configuration is still in InProgress and may be in a queue, deploy, or fail in any one of the states", "zh_CN":"加速域名的部署状态：Deployed表示该加速域名配置完成部署；InProgress表示该加速域名配置的部署任务还在进行中，可能处于排队、部署中或失败任意一种状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Accelerate the CDN service status of the domain name: This is false when the accelerated domain name CDN service is canceled; this is true when the accelerated domain name CDN service is restored.", "zh_CN":"加速域名的CDN服务状态：当取消加速域名CDN服务后，此项为false；当恢复加速域名CDN服务后，此项为true"}
  CdnServiceStatus *string `json:"cdn-service-status,omitempty" xml:"cdn-service-status,omitempty" require:"true"`
  // {"en":"Accelerated domain activation: This is false when the accelerated domain name service is disabled; true when the accelerated domain name service is enabled", "zh_CN":"加速域名的启用状态：当禁用加速域名服务后，此项为false；当启用加速域名服务后，此项为true"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
  // {"en":"Domain name last modified time,
  // Format: 2024-01-01T22:30:00+08:00", "zh_CN":"域名最近修改时间，格式: 2024-01-01T22:30:00+08:00"}
  LastModified *string `json:"last-modified,omitempty" xml:"last-modified,omitempty" require:"true"`
  // {"en":"Billing areas.", "zh_CN":"计费区域"}
  BillingArea *string `json:"billing-areas,omitempty" xml:"billing-areas,omitempty" require:"true"`
}

func (s QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetDomainName(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.DomainName = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetDomainId(v int) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.DomainId = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetCname(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.Cname = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetServiceType(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.ServiceType = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetStatus(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.Status = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetCdnServiceStatus(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.CdnServiceStatus = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetEnabled(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.Enabled = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetLastModified(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.LastModified = &v
  return s
}

func (s *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList) SetBillingArea(v string) *QueryApiDomainListServiceQueryApiDomainListServiceResponseDomainList {
  s.BillingArea = &v
  return s
}

type QueryApiDomainListServicePaths struct {
}

func (s QueryApiDomainListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServicePaths) GoString() string {
  return s.String()
}

type QueryApiDomainListServiceParameters struct {
  // {"en":"Public CNAME alias, optional entry, does not indicate all domain names under the query account number
  // The customer has the demand that the domain cname share more than one level, so we introduce the cname-label identifier in the interface, which is a set of domain cname with the same cname-label, and share the first level of cname.", "zh_CN":"共用一级别名标示，可选入参，不选表示查询账号下所有域名
  // 客户存在较多一级域名共用的需求，因此在接口中引入cname-label标识，即拥有相同cname-label的一组域名，共用一级cname。关于cname-label的具体使用方式和注意事项，请参看【创建加速域名】和【修改域名配置】接口"}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
}

func (s QueryApiDomainListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceParameters) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceParameters) SetCnameLabel(v string) *QueryApiDomainListServiceParameters {
  s.CnameLabel = &v
  return s
}

type QueryApiDomainListServiceRequestHeader struct {
}

func (s QueryApiDomainListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryApiDomainListServiceResponseHeader struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http-status-code,omitempty" xml:"http-status-code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s QueryApiDomainListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceResponseHeader) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceResponseHeader) SetHttpStatus(v int) *QueryApiDomainListServiceResponseHeader {
  s.HttpStatus = &v
  return s
}

func (s *QueryApiDomainListServiceResponseHeader) SetXCncRequestId(v string) *QueryApiDomainListServiceResponseHeader {
  s.XCncRequestId = &v
  return s
}




type CreateDomainRequest struct {
  // {"en":"Version code , the current version is 1.0.0", "zh_CN":"版本号，当前版本号1.0.0"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Need to access the domain name of the CDN. a generic domain name is supported, starting with the symbol '.', such as.example.com, which also contains a multilevel 'a.b.example.com'.If example.com is filed, the domain name xx.example.com does not need to be filed.", "zh_CN":"需要接入CDN的域名。支持泛域名，以符号“.”开头，如：.example.com，泛域名也包含多级“a.b.example.com”。
  // 如果example.com已备案，那么域名xx.example.com则不需要备案。"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"The service type of the accelerated domain name (only one service type can be submitted at a time):
  //   web/web-https: Web page acceleration/Web page acceleration-https
  // wsa/Wsa-https: Full-station acceleration/full-station acceleration-https
  // vodstream/vod-https: on-demand acceleration/on-demand acceleration-https
  // download/dl-https: Download Acceleration/Download Acceleration-https
  // livestream/live-https/cloudv-live: livestream acceleration
  // v6sa/osv6: IPv6 Security&Acceleration Solution/IPv6 One-stop Solution
  // Note:
  // 1. the https in the code, such as web-https does not represent immediate support for https access, you need to upload the certificate to support https.
  //   ", "zh_CN":"加速域名的服务类型（一次只能提交一个服务类型）：
  // web/web-https：网页加速/网页加速-https
  // wsa/wsa-https：全站加速/全站加速-https
  // vodstream/vod-https：点播加速/点播加速-https
  // download/dl-https：下载加速/下载加速-https
  // livestream/live-https/cloudv-live：直播加速
  // v6sa/osv6：ipv6安全加速解决方案/IPv6一体化解决方案
  // 注意：
  // 1、service-type中的https不代表立即开启https，比如web-https中的https并不代表立刻支持https访问，需上传完证书后才可以支持https，切记！"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty"`
  // {"en":"The acceleration area of the acceleration domain, if the resource coverage needs to be limited according to the area, the acceleration area needs to be specified.
  //     When no acceleration area is specified, we will provide acceleration services with optimal resource coverage according to the service area opened by the customer. 
  //     Multiple regions are separated by semicolons, and the supported regions are as follows: cn (Mainland China), am (Americas), 
  //     emea (Europe, Middle East, Africa), apac (Asia-Pacific region).",
  // "zh_CN":"加速域名的加速区域，如果有需要根据区域限定资源覆盖时，才需要指定加速区域。未指定加速区域时，我们将按照客户开通的服务区域，以最优的资源覆盖提供加速服务。多个区域以分号分隔，支持配置的区域如下：cn（中国大陆）、am（美洲）、emea（欧洲、中东、非洲）、apac（亚太地区）"}
  ServiceAreas *string `json:"service-areas,omitempty" xml:"service-areas,omitempty"`
  // {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Configuration template, if you want to add the a domain using some specified configuration by default, you can specify the template id. For more detail, please contract the technical support.",
  //     "zh_CN":"配置单模板，特定的使用场景下，如果希望新增的加速域名参照某些指定配置时，可以指定配置单模板，具体使用请咨询对应的客户负责人。"}
  ConfigFormId *int `json:"config-form-id,omitempty" xml:"config-form-id,omitempty"`
  // {"en":"Refer to the configuration of the specified domain.
  // Note:
  // 1. If the referenced domain uses a certificate, the new domain should be in the 'DNS name' of the certificate.
  // 2. If the referenced domain has no China ICP, while the new domain name has, it may affect the cover resources and service quality.
  // 3. If the referenced domain has China ICP, while the new domain name doesn't, then the cover resources may be re-selected if it does not meet the policy requirements.
  // 4. It is not allowed to reference a domain which is traffic-free.", "zh_CN":"参照指定域名的配置，来创建加速域名。
  // 注意：
  // 1.参照域名如果有使用证书，新增域名也要在对应证书授权范围内。
  // 2.参照未备案域名，新增的域名如果已备案，可能影响资源使用和服务质量。
  // 3.参照备案域名，新增的域名如果未备案，若资源不满足政策要求，可能重选。
  // 4.不允许参照免流域名创建新域名。"}
  ReferencedDomainName *string `json:"referenced-domain-name,omitempty" xml:"referenced-domain-name,omitempty"`
  // {"en":"If you need to share a CNAME between domains, you can use this parameter. This parameter is a unique label for a public CNAME. Domains with the same cname-label will have the same CNAME. 
  // Note:
  // 1. Domains with the same cname-label have the same coverage.
  // 2. Constraints of sharing a CNAME: consistent service-type, consistent certificate-id (if there is a certificate), consistent service-areas
  // 3. Multiple http domains can share a CNAME, multiple sni https domains can share a CNAME too.
  // 4. When a cname-label is used by a single domain, then the domain can be canceled acceleration. While a cname-label using by more then one domains, they can not be canceled acceleration.
  // 5. Support the purpose of modifying cname by modifying cname-label. )
  // ", "zh_CN":"共用一级标签，若有多个加速域名需要共用一级域名，则可以使用该参数。即拥有相同cname-label的一组域名，共用一级cname。
  // 注意：
  // 1、拥有相同cname-label的域名共用一级cname，且有完全一致的dns覆盖
  // 2、共用一级的约束：加速类型一致(service-type)、证书id一致（certificate-id,如果有证书）、加速区域一致(service-areas)
  // 3、多个http域名可共用一级，多个sni https域名可共用一级
  // 4、单个域名使用cname-label时，域名可cancel；多个域名共用一级时，不允许cancel这些域名
  // 5、支持通过修改cname-label达到修改cname的目的。）
  // "}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
  // {"en":"The first level of cname prefix, true, indicates that the domain cname is used as the cname prefix, otherwise the 14-bit random string (number + letter) is used as the cname prefix.
  // Note: When the prefix is a generic domain name, a wsall is added as a prefix. Such as... Baidu.com.wscloudcdn.com, which will generate wsall.Baidu.com.wscloudcdn.com", "zh_CN":"一级cname前缀，true表示使用域名名称作为cname前缀，否则，使用14位随机串（数字+字母）作为cname前缀。
  //   注意：当前缀是泛域名时，则再增加wsall作为前缀。如.baidu.com.wscloudcdn.com，会生成wsall.baidu.com.wscloudcdn.com"}
  CnameWithCustomizedPrefix *string `json:"cname-with-customized-prefix,omitempty" xml:"cname-with-customized-prefix,omitempty"`
  // {"en":"Back-to-origin policy setting, which is used to set the origin site information and the back-to-origin policy of the none-live accelerated domain", "zh_CN":"回源策略设置(非直播域名使用)，用于设置加速域名的源站信息和回源策略。"}
  OriginConfig *CreateDomainCreateDomainRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" type:"Struct"`
  // {"en":"Live domain configuration, used to set the push flow of rtmp live acceleration domain (use required)
  // Note: In addition to the API call permission, you need to contact the dedicated customer service to apply for the corresponding API client template.", "zh_CN":"直播域名配置，用于设置rtmp直播加速域名的推拉流（使用需申请）
  // 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
  LiveConfig *CreateDomainCreateDomainRequestLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
  // {"en":"Identifies whether a domain name is fully overseas accelerated.
  // Whether the default is false
  // True: indicates that the client domain name is a pure overseas acceleration
  // False: Indicates that the client domain name has accelerated in China", "zh_CN":"标识域名是否是纯海外加速的。
  // 默认是否（false）
  // true ：表示客户域名纯海外加速
  // false：表示客户域名有在中国加速"}
  AccelerateNoChina *string `json:"accelerate-no-china,omitempty" xml:"accelerate-no-china,omitempty"`
  // {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip, X-Forwarded-For and ori_X-Forwarded-For.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip、X-Forwarded-For、ori_X-Forwarded-For
  // 1） Cdn-Src-Ip： 回源头部名称为Cdn-Src-Ip，获取与节点进行建联的IP作为客户端IP传递回源。
  // 2） X-Forwarded-For： 回源头部名称为X-Forwarded-For，携带的客户端IP值是Cdn-Src-Ip获取到的建联IP。
  // 3） ori_X-Forwarded-For：客户端请求CDN节点时会自带X-Forwarded-For，则CDN透传此头部和值回源。"}
  HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
  // {"en":"The live streaming domain which is pull domian ,and  directly returned to the source to verify the configuration.
  // which can be an IP or a domain name.
  // Can be IP or domain name. Ip and domain names can only be one. Multiple input parameters are not supported.", "zh_CN":"直播拉流域名，直接回源校验配置。
  // 可以是IP或域名。ip和域名只能一种。不支持多个入参。"}
  UpstreamHost *string `json:"upstream-host,omitempty" xml:"upstream-host,omitempty"`
  // {"en":"Set the publishing point of the live push-pull domain name
  // note:
  // 1. Pull flow and corresponding push flow domain name must be configured with the same publishing point.
  // 2. do not want to modify the publishing point, do not pass the node and the following parameters
  // 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
  // 注意：
  // 1、拉流和对应的推流域名，必须配置相同的发布点；
  // 2、不想修改发布点时，不要传入该节点及以下入参；
  // 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
  PublishPoints []*CreateDomainCreateDomainRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
  // {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your  certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate]", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】"}
  Ssl *CreateDomainCreateDomainRequestSsl `json:"ssl,omitempty" xml:"ssl,omitempty" type:"Struct"`
}

func (s CreateDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
  return s.String()
}

func (s *CreateDomainRequest) SetVersion(v string) *CreateDomainRequest {
  s.Version = &v
  return s
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
  s.DomainName = &v
  return s
}

func (s *CreateDomainRequest) SetServiceType(v string) *CreateDomainRequest {
  s.ServiceType = &v
  return s
}

func (s *CreateDomainRequest) SetServiceAreas(v string) *CreateDomainRequest {
  s.ServiceAreas = &v
  return s
}

func (s *CreateDomainRequest) SetComment(v string) *CreateDomainRequest {
  s.Comment = &v
  return s
}

func (s *CreateDomainRequest) SetConfigFormId(v int) *CreateDomainRequest {
  s.ConfigFormId = &v
  return s
}

func (s *CreateDomainRequest) SetReferencedDomainName(v string) *CreateDomainRequest {
  s.ReferencedDomainName = &v
  return s
}

func (s *CreateDomainRequest) SetCnameLabel(v string) *CreateDomainRequest {
  s.CnameLabel = &v
  return s
}

func (s *CreateDomainRequest) SetCnameWithCustomizedPrefix(v string) *CreateDomainRequest {
  s.CnameWithCustomizedPrefix = &v
  return s
}

func (s *CreateDomainRequest) SetOriginConfig(v *CreateDomainCreateDomainRequestOriginConfig) *CreateDomainRequest {
  s.OriginConfig = v
  return s
}

func (s *CreateDomainRequest) SetLiveConfig(v *CreateDomainCreateDomainRequestLiveConfig) *CreateDomainRequest {
  s.LiveConfig = v
  return s
}

func (s *CreateDomainRequest) SetAccelerateNoChina(v string) *CreateDomainRequest {
  s.AccelerateNoChina = &v
  return s
}

func (s *CreateDomainRequest) SetHeaderOfClientip(v string) *CreateDomainRequest {
  s.HeaderOfClientip = &v
  return s
}

func (s *CreateDomainRequest) SetUpstreamHost(v string) *CreateDomainRequest {
  s.UpstreamHost = &v
  return s
}

func (s *CreateDomainRequest) SetPublishPoints(v []*CreateDomainCreateDomainRequestPublishPoints) *CreateDomainRequest {
  s.PublishPoints = v
  return s
}

func (s *CreateDomainRequest) SetSsl(v *CreateDomainCreateDomainRequestSsl) *CreateDomainRequest {
  s.Ssl = v
  return s
}

type CreateDomainCreateDomainRequestOriginConfig struct {
  // {"en":"Origin address, which can be an IP or domain name.
  // 1. Multiple IPs are supported, separated by semicolons.
  // 2. Only one domain name is allowed. IP and domain name cannot exist at the same time.
  // 3. The length cannot exceed 500 characters.
  // 4. The number of IPs cannot exceed 15.", "zh_CN":"回源地址，可以是IP或域名。
  // 1、IP以分号分隔，支持多个。
  // 2、域名只能输入一个。IP与域名不能同时输入。
  // 3、限制最大不能超过500个字符长度。
  // 4、源IP个数不能超过15个。"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
  // {"en":"The Origin HOST for changing the HOST field in the return source HTTP request header. The supported domain name formats, each segement separated by a dot, does not exceed 62 characters, the total length should not exceed 128 characters.
  // .", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: 域名，每段（点号分隔）长度小于等于62，域名总长度小于等于128。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
}

func (s CreateDomainCreateDomainRequestOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainCreateDomainRequestOriginConfig) GoString() string {
  return s.String()
}

func (s *CreateDomainCreateDomainRequestOriginConfig) SetOriginIps(v string) *CreateDomainCreateDomainRequestOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *CreateDomainCreateDomainRequestOriginConfig) SetDefaultOriginHostHeader(v string) *CreateDomainCreateDomainRequestOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

type CreateDomainCreateDomainRequestLiveConfig struct {
  // {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
  StreamType *string `json:"stream-type,omitempty" xml:"stream-type,omitempty"`
  // {"en":"The push-pull domain name is used to set the push-flow domain name corresponding to the rtmp live streaming domain name. When the stream-type is pull, at least one of the source IP address and the corresponding push-stream domain name is not empty. When the stream-type is push, Incoming.", "zh_CN":"配套推流域名，用于设置rtmp直播拉流域名对应的推流域名，当stream-type为pull时，源站IP和配套推流域名至少一个不为空；当stream-type为push时，无需传入。"}
  OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty"`
  // {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.
  // 1. If it is a push-pull flow package, fill in 127.0.0.1, and the system will also default to 127.0.0.1.
  // 2. If it is directly returning to the source, fill in the source IP of the source pull stream.
  // ", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
  // 1、如果是推拉流配套，则填写127.0.0.1，不传系统也默认为127.0.0.1
  // 2、如果是直接回源拉流，则填写回源拉流的源站IP"}
  LiveConfigOriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
}

func (s CreateDomainCreateDomainRequestLiveConfig) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainCreateDomainRequestLiveConfig) GoString() string {
  return s.String()
}

func (s *CreateDomainCreateDomainRequestLiveConfig) SetStreamType(v string) *CreateDomainCreateDomainRequestLiveConfig {
  s.StreamType = &v
  return s
}

func (s *CreateDomainCreateDomainRequestLiveConfig) SetOriginPushHost(v string) *CreateDomainCreateDomainRequestLiveConfig {
  s.OriginPushHost = &v
  return s
}

func (s *CreateDomainCreateDomainRequestLiveConfig) SetLiveConfigOriginIps(v string) *CreateDomainCreateDomainRequestLiveConfig {
  s.LiveConfigOriginIps = &v
  return s
}

type CreateDomainCreateDomainRequestPublishPoints struct     {
  // {"en":"Livestream domain settings. Publish point, support multiple, do not pass the system by default to generate a publishing point uri for [/]", "zh_CN":"发布点，支持多个，不传系统默认生成一条发布点uri为“/”"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s CreateDomainCreateDomainRequestPublishPoints) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainCreateDomainRequestPublishPoints) GoString() string {
  return s.String()
}

func (s *CreateDomainCreateDomainRequestPublishPoints) SetUri(v string) *CreateDomainCreateDomainRequestPublishPoints {
  s.Uri = &v
  return s
}

type CreateDomainCreateDomainRequestSsl struct {
  // {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
  UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用合用证书（暂不支持）"}
  UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID，use-ssl为true时，才能传ssl-certificate-id。"}
  SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty"`
}

func (s CreateDomainCreateDomainRequestSsl) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainCreateDomainRequestSsl) GoString() string {
  return s.String()
}

func (s *CreateDomainCreateDomainRequestSsl) SetUseSsl(v string) *CreateDomainCreateDomainRequestSsl {
  s.UseSsl = &v
  return s
}

func (s *CreateDomainCreateDomainRequestSsl) SetUseForSni(v string) *CreateDomainCreateDomainRequestSsl {
  s.UseForSni = &v
  return s
}

func (s *CreateDomainCreateDomainRequestSsl) SetSslCertificateId(v int) *CreateDomainCreateDomainRequestSsl {
  s.SslCertificateId = &v
  return s
}

type CreateDomainResponse struct {
  // {"en":"The name of the service domain automatically generated by the My company, for example: xxxx.cdn30.com", "zh_CN":"由我司自动生成的服务域名名称，例如：xxxx.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
  return s.String()
}

func (s *CreateDomainResponse) SetCname(v string) *CreateDomainResponse {
  s.Cname = &v
  return s
}

func (s *CreateDomainResponse) SetCode(v string) *CreateDomainResponse {
  s.Code = &v
  return s
}

func (s *CreateDomainResponse) SetMessage(v string) *CreateDomainResponse {
  s.Message = &v
  return s
}

type CreateDomainPaths struct {
}

func (s CreateDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainPaths) GoString() string {
  return s.String()
}

type CreateDomainParameters struct {
}

func (s CreateDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainParameters) GoString() string {
  return s.String()
}

type CreateDomainRequestHeader struct {
}

func (s CreateDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainRequestHeader) GoString() string {
  return s.String()
}

type CreateDomainResponseHeader struct {
  // {"en":"If httpstatus=202, the interface is successfully invoked. And the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况。"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）。"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"A URL used to access the domain name information, where the domain-id is a unique identifier for the domain name.", "zh_CN":"用于访问该域名信息的URL，其中domain-id为该域名的唯一标识。"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDomainResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateDomainResponseHeader) SetHttpStatus(v int) *CreateDomainResponseHeader {
  s.HttpStatus = &v
  return s
}

func (s *CreateDomainResponseHeader) SetXCncRequestId(v string) *CreateDomainResponseHeader {
  s.XCncRequestId = &v
  return s
}

func (s *CreateDomainResponseHeader) SetLocation(v string) *CreateDomainResponseHeader {
  s.Location = &v
  return s
}




type DispatchWarnLogRequest struct {
  // {"en":"Language type The default language: Chinese 'zh_CN': The English 'en'", "zh_CN":"语言类型 默认 中文, 中文  'zh_CN'英文  'en'"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"en":"domainId", "zh_CN":"调度域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"policyId", "zh_CN":"策略id"}
  PolicyId *int `json:"policyId,omitempty" xml:"policyId,omitempty"`
  // {"en":"view name", "zh_CN":"线路中文名，支持模糊搜索，例如‘中国电信’"}
  View *string `json:"view,omitempty" xml:"view,omitempty"`
  // {"en":"warnType", "zh_CN":"告警类型 0：调度告警 1：故障告警 2：恢复通知"}
  WarnType *int `json:"warnType,omitempty" xml:"warnType,omitempty"`
  // {"en":"Query start time format yyyy-mm-dd HH: MM :ss for example '2021-04-26 18:00:00'", "zh_CN":"查询开始时间 格式yyyy-MM-dd HH:mm:ss  例如‘2021-04-26 18:00:00’"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"End of query format yyyy-mm-dd HH: MM :ss for example '2021-04-26 18:00:00'", "zh_CN":"查询结束时间 格式yyyy-MM-dd HH:mm:ss  例如‘2021-04-26 18:00:00’"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s DispatchWarnLogRequest) String() string {
  return tea.Prettify(s)
}

func (s DispatchWarnLogRequest) GoString() string {
  return s.String()
}

func (s *DispatchWarnLogRequest) SetLanguage(v string) *DispatchWarnLogRequest {
  s.Language = &v
  return s
}

func (s *DispatchWarnLogRequest) SetDomainId(v int) *DispatchWarnLogRequest {
  s.DomainId = &v
  return s
}

func (s *DispatchWarnLogRequest) SetPolicyId(v int) *DispatchWarnLogRequest {
  s.PolicyId = &v
  return s
}

func (s *DispatchWarnLogRequest) SetView(v string) *DispatchWarnLogRequest {
  s.View = &v
  return s
}

func (s *DispatchWarnLogRequest) SetWarnType(v int) *DispatchWarnLogRequest {
  s.WarnType = &v
  return s
}

func (s *DispatchWarnLogRequest) SetStartTime(v string) *DispatchWarnLogRequest {
  s.StartTime = &v
  return s
}

func (s *DispatchWarnLogRequest) SetEndTime(v string) *DispatchWarnLogRequest {
  s.EndTime = &v
  return s
}

type DispatchWarnLogResponse struct {
  // {"en":"Status code. See "Scheduling Business Status Codes" for a detailed description of RESCODE.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code", "zh_CN":"状态码的详细说明"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"业务数据"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s DispatchWarnLogResponse) String() string {
  return tea.Prettify(s)
}

func (s DispatchWarnLogResponse) GoString() string {
  return s.String()
}

func (s *DispatchWarnLogResponse) SetResCode(v int) *DispatchWarnLogResponse {
  s.ResCode = &v
  return s
}

func (s *DispatchWarnLogResponse) SetMsg(v string) *DispatchWarnLogResponse {
  s.Msg = &v
  return s
}

func (s *DispatchWarnLogResponse) SetContent(v []*string) *DispatchWarnLogResponse {
  s.Content = v
  return s
}

type DispatchWarnLogPaths struct {
}

func (s DispatchWarnLogPaths) String() string {
  return tea.Prettify(s)
}

func (s DispatchWarnLogPaths) GoString() string {
  return s.String()
}

type DispatchWarnLogParameters struct {
}

func (s DispatchWarnLogParameters) String() string {
  return tea.Prettify(s)
}

func (s DispatchWarnLogParameters) GoString() string {
  return s.String()
}

type DispatchWarnLogRequestHeader struct {
}

func (s DispatchWarnLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DispatchWarnLogRequestHeader) GoString() string {
  return s.String()
}

type DispatchWarnLogResponseHeader struct {
}

func (s DispatchWarnLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DispatchWarnLogResponseHeader) GoString() string {
  return s.String()
}




type RestoreApiDomainServiceRequest struct {
}

func (s RestoreApiDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s RestoreApiDomainServiceRequest) GoString() string {
  return s.String()
}

type RestoreApiDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http-status-code,omitempty" xml:"http-status-code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s RestoreApiDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s RestoreApiDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *RestoreApiDomainServiceResponse) SetCode(v string) *RestoreApiDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *RestoreApiDomainServiceResponse) SetMessage(v string) *RestoreApiDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *RestoreApiDomainServiceResponse) SetHttpStatus(v int) *RestoreApiDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *RestoreApiDomainServiceResponse) SetXCncRequestId(v string) *RestoreApiDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type RestoreApiDomainServicePaths struct {
  // {"en":"", "zh_CN":"加速域名在系统中对应的ID
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s RestoreApiDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s RestoreApiDomainServicePaths) GoString() string {
  return s.String()
}

func (s *RestoreApiDomainServicePaths) SetDomainId(v int) *RestoreApiDomainServicePaths {
  s.DomainId = &v
  return s
}

type RestoreApiDomainServiceParameters struct {
}

func (s RestoreApiDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s RestoreApiDomainServiceParameters) GoString() string {
  return s.String()
}

type RestoreApiDomainServiceRequestHeader struct {
}

func (s RestoreApiDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RestoreApiDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type RestoreApiDomainServiceResponseHeader struct {
}

func (s RestoreApiDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RestoreApiDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type ControlDomainRequest struct {
  // {"en":"Domain names are separated by English symbols (e.g. aaa.com; bb.com).", "zh_CN":"域名 多个域名用英文符号;隔开（如： aaa.com;bb.com）"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Operation: 1 is enabled; 0 is disabled", "zh_CN":"操作：1 启用；0 停用"}
  Operate *int `json:"operate,omitempty" xml:"operate,omitempty" require:"true"`
  // {"en":"Return Chinese results for null (default)
  // En: Return the English prompt result", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ControlDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s ControlDomainRequest) GoString() string {
  return s.String()
}

func (s *ControlDomainRequest) SetDomainName(v string) *ControlDomainRequest {
  s.DomainName = &v
  return s
}

func (s *ControlDomainRequest) SetOperate(v int) *ControlDomainRequest {
  s.Operate = &v
  return s
}

func (s *ControlDomainRequest) SetLanguage(v string) *ControlDomainRequest {
  s.Language = &v
  return s
}

type ControlDomainResponse struct {
  // {"en":"Status code. For detailed description of resCode, please refer to 'Status Codes of Dispatch Business'.", "zh_CN":"状态码，详细说明请参见“业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"recordId, ID of host name record, used to identify this record.", "zh_CN":"域名的详细说明。"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s ControlDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s ControlDomainResponse) GoString() string {
  return s.String()
}

func (s *ControlDomainResponse) SetResCode(v string) *ControlDomainResponse {
  s.ResCode = &v
  return s
}

func (s *ControlDomainResponse) SetMsg(v string) *ControlDomainResponse {
  s.Msg = &v
  return s
}

func (s *ControlDomainResponse) SetContent(v []*string) *ControlDomainResponse {
  s.Content = v
  return s
}

type ControlDomainPaths struct {
}

func (s ControlDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s ControlDomainPaths) GoString() string {
  return s.String()
}

type ControlDomainParameters struct {
}

func (s ControlDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s ControlDomainParameters) GoString() string {
  return s.String()
}

type ControlDomainRequestHeader struct {
}

func (s ControlDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDomainRequestHeader) GoString() string {
  return s.String()
}

type ControlDomainResponseHeader struct {
}

func (s ControlDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlDomainResponseHeader) GoString() string {
  return s.String()
}




type GetPagingDomainListRequest struct {
  // {"en":"Page number must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.", "zh_CN":"分页的页码，必须为大于0的正整数。不传默认不分页，若传参则pageSize必填.。"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
  // {"en":"Number of domain name data items for paging, must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.", "zh_CN":"分页的域名数据条数，必须大于0的正整数。不传默认不分页，若传参则pageSize必填.。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Specify the service type to be queried. Multiple services are allowed. Data will be returned if any one service is satisfied. If not passed, all services will be checked by default. For example: [wsa,waf], returns all domains whose services include wsa or include waf.", "zh_CN":"指定查询的服务，允许多个服务，任意一个服务满足就返回数据，不传默认查全部服务。如：[wsa,waf], 则返回服务包含wsa或包含waf的所有域名。"}
  ServiceTypes []*string `json:"serviceTypes,omitempty" xml:"serviceTypes,omitempty" type:"Repeated"`
  // {"en":"Specify the accelerated domain name for the query. Multiple domain names are allowed. If not specified, all domain names will be searched by default.", "zh_CN":"指定查询的加速域名，允许多个域名，不传默认查全部域名。"}
  DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" type:"Repeated"`
  // {"en":"RFC3339 formatted date indicating the starting date. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"	查询开始时间，支持时间格式如：2024-01-01T22:30:00+08:00"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"RFC3339 formatted date indicating the ending date. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"查询结束时间，支持时间格式如：2024-01-01T22:30:00+08:00"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":" Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling, deployFailed, disableFailed.", "zh_CN":"加速域名的状态：enabled表示已启用；disabled表示已禁用；deploying表示配置部署中；checking表示审核中；disabling表示禁用中；deployFailed表示配置部署失败；disableFailed表示禁用失败。不传默认查全部"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPagingDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListRequest) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListRequest) SetPageNumber(v int) *GetPagingDomainListRequest {
  s.PageNumber = &v
  return s
}

func (s *GetPagingDomainListRequest) SetPageSize(v int) *GetPagingDomainListRequest {
  s.PageSize = &v
  return s
}

func (s *GetPagingDomainListRequest) SetServiceTypes(v []*string) *GetPagingDomainListRequest {
  s.ServiceTypes = v
  return s
}

func (s *GetPagingDomainListRequest) SetDomainNames(v []*string) *GetPagingDomainListRequest {
  s.DomainNames = v
  return s
}

func (s *GetPagingDomainListRequest) SetStartTime(v string) *GetPagingDomainListRequest {
  s.StartTime = &v
  return s
}

func (s *GetPagingDomainListRequest) SetEndTime(v string) *GetPagingDomainListRequest {
  s.EndTime = &v
  return s
}

func (s *GetPagingDomainListRequest) SetStatus(v string) *GetPagingDomainListRequest {
  s.Status = &v
  return s
}

type GetPagingDomainListResponse struct {
  // {"en":"Responses the page number of the data", "zh_CN":"所有满足条件的数据总条数"}
  TotalCount *int `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"total pages", "zh_CN":"总页数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
  // {"en":"Responses the page number of the data", "zh_CN":"返回数据的页码"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Number of data page", "zh_CN":"每个页面的数据条数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Domain list.", "zh_CN":"域名列表"}
  ResultList []*GetPagingDomainListGetPagingDomainListResponseResultList `json:"resultList,omitempty" xml:"resultList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPagingDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListResponse) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListResponse) SetTotalCount(v int) *GetPagingDomainListResponse {
  s.TotalCount = &v
  return s
}

func (s *GetPagingDomainListResponse) SetTotalPageNumber(v int) *GetPagingDomainListResponse {
  s.TotalPageNumber = &v
  return s
}

func (s *GetPagingDomainListResponse) SetPageNumber(v int) *GetPagingDomainListResponse {
  s.PageNumber = &v
  return s
}

func (s *GetPagingDomainListResponse) SetPageSize(v int) *GetPagingDomainListResponse {
  s.PageSize = &v
  return s
}

func (s *GetPagingDomainListResponse) SetResultList(v []*GetPagingDomainListGetPagingDomainListResponseResultList) *GetPagingDomainListResponse {
  s.ResultList = v
  return s
}

type GetPagingDomainListGetPagingDomainListResponseResultList struct     {
  // {"en":"Cname of the accelerated domain", "zh_CN":"加速域名cname，如：a1.example.com.wscdns.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Create time of the accelerated domain. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"	域名创建时间，时间格式如：2024-01-01T22:30:00+08:00"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Corresponding domain ID", "zh_CN":"对应的域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Service type for accelerated domain name", "zh_CN":"加速域名的服务，如[wsa,waf]。"}
  ServiceTypes []*string `json:"serviceTypes,omitempty" xml:"serviceTypes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling.", "zh_CN":"	加速域名的状态：enabled表示已启用；disabled表示已禁用；deploying表示部署中；checking表示审核中；disabling表示禁用中；deployFailed表示配置部署失败；disableFailed表示禁用失败"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Accelerated domain enabling status: true indicates that it is enabled, false indicates that it is disabled.", "zh_CN":"加速域名启用状态：true为启用，false为禁用。"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPagingDomainListGetPagingDomainListResponseResultList) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListGetPagingDomainListResponseResultList) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetCname(v string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.Cname = &v
  return s
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetCreateTime(v string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.CreateTime = &v
  return s
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetDomainId(v string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.DomainId = &v
  return s
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetDomainName(v string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.DomainName = &v
  return s
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetServiceTypes(v []*string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.ServiceTypes = v
  return s
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetStatus(v string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.Status = &v
  return s
}

func (s *GetPagingDomainListGetPagingDomainListResponseResultList) SetEnabled(v string) *GetPagingDomainListGetPagingDomainListResponseResultList {
  s.Enabled = &v
  return s
}

type GetPagingDomainListPaths struct {
}

func (s GetPagingDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListPaths) GoString() string {
  return s.String()
}

type GetPagingDomainListParameters struct {
}

func (s GetPagingDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListParameters) GoString() string {
  return s.String()
}

type GetPagingDomainListRequestHeader struct {
}

func (s GetPagingDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListRequestHeader) GoString() string {
  return s.String()
}

type GetPagingDomainListResponseHeader struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s GetPagingDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListResponseHeader) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListResponseHeader) SetCode(v int) *GetPagingDomainListResponseHeader {
  s.Code = &v
  return s
}

func (s *GetPagingDomainListResponseHeader) SetXCncRequestId(v string) *GetPagingDomainListResponseHeader {
  s.XCncRequestId = &v
  return s
}




type QueryDispatchDomainsRequest struct {
  // {"en":"Query inital records by page", "zh_CN":"分页查询起始记录"}
  Start *int `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {"en":"Number if quried items by page", "zh_CN":"分页查询条数"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"Parameter of domain fuzzy search Domain fuzzy search. If no domain entered, it means all domains and domains details under the user are queried and returned.", "zh_CN":"域名模糊搜索参数
  // 域名模糊搜索，如果没有填写域名，则返回该用户的所有域名及域名相应信息。"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
  // {"en":"Parameter of domain fuzzy search Domain fuzzy search. If no domain entered, it means all domains and domains details under the user are queried and returned.", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s QueryDispatchDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsRequest) SetStart(v int) *QueryDispatchDomainsRequest {
  s.Start = &v
  return s
}

func (s *QueryDispatchDomainsRequest) SetLimit(v int) *QueryDispatchDomainsRequest {
  s.Limit = &v
  return s
}

func (s *QueryDispatchDomainsRequest) SetDomainName(v string) *QueryDispatchDomainsRequest {
  s.DomainName = &v
  return s
}

func (s *QueryDispatchDomainsRequest) SetLanguage(v string) *QueryDispatchDomainsRequest {
  s.Language = &v
  return s
}

type QueryDispatchDomainsResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Detailed description of the domain. count The total number of user's domains rows The queried results of domains", "zh_CN":"域名的详细说明。count 用户域名总数量"}
  Content *QueryDispatchDomainsQueryDispatchDomainsResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s QueryDispatchDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsResponse) SetResCode(v string) *QueryDispatchDomainsResponse {
  s.ResCode = &v
  return s
}

func (s *QueryDispatchDomainsResponse) SetMsg(v string) *QueryDispatchDomainsResponse {
  s.Msg = &v
  return s
}

func (s *QueryDispatchDomainsResponse) SetContent(v *QueryDispatchDomainsQueryDispatchDomainsResponseContent) *QueryDispatchDomainsResponse {
  s.Content = v
  return s
}

type QueryDispatchDomainsQueryDispatchDomainsResponseContent struct {
  Rows []*QueryDispatchDomainsQueryDispatchDomainsResponseContentRows `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDispatchDomainsQueryDispatchDomainsResponseContent) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsQueryDispatchDomainsResponseContent) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsQueryDispatchDomainsResponseContent) SetRows(v []*QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) *QueryDispatchDomainsQueryDispatchDomainsResponseContent {
  s.Rows = v
  return s
}

type QueryDispatchDomainsQueryDispatchDomainsResponseContentRows struct     {
  // {"en":"domainId", "zh_CN":"域名ID标识"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domainName", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Dispatch CNAME", "zh_CN":"调度CNAME"}
  DispatchCname *string `json:"dispatchCname,omitempty" xml:"dispatchCname,omitempty" require:"true"`
  // {"en":"The number of domain policies", "zh_CN":"域名策略数量"}
  PolicyCount *string `json:"policyCount,omitempty" xml:"policyCount,omitempty" require:"true"`
}

func (s QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) SetDomainId(v int) *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows {
  s.DomainId = &v
  return s
}

func (s *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) SetDomainName(v string) *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows {
  s.DomainName = &v
  return s
}

func (s *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) SetDispatchCname(v string) *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows {
  s.DispatchCname = &v
  return s
}

func (s *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows) SetPolicyCount(v string) *QueryDispatchDomainsQueryDispatchDomainsResponseContentRows {
  s.PolicyCount = &v
  return s
}

type QueryDispatchDomainsPaths struct {
}

func (s QueryDispatchDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsPaths) GoString() string {
  return s.String()
}

type QueryDispatchDomainsParameters struct {
}

func (s QueryDispatchDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsParameters) GoString() string {
  return s.String()
}

type QueryDispatchDomainsRequestHeader struct {
}

func (s QueryDispatchDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsRequestHeader) GoString() string {
  return s.String()
}

type QueryDispatchDomainsResponseHeader struct {
}

func (s QueryDispatchDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsResponseHeader) GoString() string {
  return s.String()
}




type AddDispatchDomainRequest struct {
  // {"en":"The added dispatch domain", "zh_CN":"添加的调度域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
  // {"en":"Dispatch CName suffix Options, domestic: cdngtm.cn, Overseas: ; the default setting is cdngtm.com", "zh_CN":"调度CName后缀
  // 可选，国内：cdngtm.cn， 海外：cdngtm.com；默认值为cdngtm.com"}
  DispatchZone *string `json:"dispatchZone,omitempty" xml:"dispatchZone,omitempty"`
  // {"en":"Dispatch CName suffix Options, domestic: cdngtm.cn, Overseas: ; the default setting is cdngtm.com", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s AddDispatchDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s AddDispatchDomainRequest) GoString() string {
  return s.String()
}

func (s *AddDispatchDomainRequest) SetDomainName(v string) *AddDispatchDomainRequest {
  s.DomainName = &v
  return s
}

func (s *AddDispatchDomainRequest) SetDispatchZone(v string) *AddDispatchDomainRequest {
  s.DispatchZone = &v
  return s
}

func (s *AddDispatchDomainRequest) SetLanguage(v string) *AddDispatchDomainRequest {
  s.Language = &v
  return s
}

type AddDispatchDomainResponse struct {
  // {"en":"Status code.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Detailed description of the domain. 
  // domainId:Domain ID 
  // domainName:Domain name 
  // dispatchCname:Dispatch CNAME", "zh_CN":"域名的详细说明。
  // domainId 域名ID标识
  // 
  // domainName 域名
  // 
  // dispatchCname 调度CNAME"}
  Content map[string]interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s AddDispatchDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s AddDispatchDomainResponse) GoString() string {
  return s.String()
}

func (s *AddDispatchDomainResponse) SetResCode(v string) *AddDispatchDomainResponse {
  s.ResCode = &v
  return s
}

func (s *AddDispatchDomainResponse) SetMsg(v string) *AddDispatchDomainResponse {
  s.Msg = &v
  return s
}

func (s *AddDispatchDomainResponse) SetContent(v map[string]interface{}) *AddDispatchDomainResponse {
  s.Content = v
  return s
}

type AddDispatchDomainPaths struct {
}

func (s AddDispatchDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s AddDispatchDomainPaths) GoString() string {
  return s.String()
}

type AddDispatchDomainParameters struct {
}

func (s AddDispatchDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s AddDispatchDomainParameters) GoString() string {
  return s.String()
}

type AddDispatchDomainRequestHeader struct {
}

func (s AddDispatchDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDispatchDomainRequestHeader) GoString() string {
  return s.String()
}

type AddDispatchDomainResponseHeader struct {
}

func (s AddDispatchDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDispatchDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryApiDomainListRequest struct {
}

func (s QueryApiDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListRequest) GoString() string {
  return s.String()
}

type QueryApiDomainListResponse struct {
  // {"en":"", "zh_CN":"httpstatus=202; 表示成功调用新增域名接口"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  DomainSummary []*QueryApiDomainListQueryApiDomainListResponseDomainSummary `json:"domain-summary,omitempty" xml:"domain-summary,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApiDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListResponse) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListResponse) SetCode(v int) *QueryApiDomainListResponse {
  s.Code = &v
  return s
}

func (s *QueryApiDomainListResponse) SetXCncRequestId(v string) *QueryApiDomainListResponse {
  s.XCncRequestId = &v
  return s
}

func (s *QueryApiDomainListResponse) SetDomainSummary(v []*QueryApiDomainListQueryApiDomainListResponseDomainSummary) *QueryApiDomainListResponse {
  s.DomainSummary = v
  return s
}

type QueryApiDomainListQueryApiDomainListResponseDomainSummary struct     {
  // {"en":"", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"", "zh_CN":"对应的域名ID"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名的回源IP"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名的服务类型"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名的部署状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名是否启用，true和false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s QueryApiDomainListQueryApiDomainListResponseDomainSummary) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListQueryApiDomainListResponseDomainSummary) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetCname(v string) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.Cname = &v
  return s
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetDomainId(v int) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.DomainId = &v
  return s
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetDomainName(v string) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.DomainName = &v
  return s
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetOriginIps(v string) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.OriginIps = &v
  return s
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetServiceType(v string) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.ServiceType = &v
  return s
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetStatus(v string) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.Status = &v
  return s
}

func (s *QueryApiDomainListQueryApiDomainListResponseDomainSummary) SetEnabled(v string) *QueryApiDomainListQueryApiDomainListResponseDomainSummary {
  s.Enabled = &v
  return s
}

type QueryApiDomainListPaths struct {
}

func (s QueryApiDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListPaths) GoString() string {
  return s.String()
}

type QueryApiDomainListParameters struct {
}

func (s QueryApiDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListParameters) GoString() string {
  return s.String()
}

type QueryApiDomainListRequestHeader struct {
}

func (s QueryApiDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListRequestHeader) GoString() string {
  return s.String()
}

type QueryApiDomainListResponseHeader struct {
}

func (s QueryApiDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListResponseHeader) GoString() string {
  return s.String()
}




type QueryChangeServerRequest struct {
}

func (s QueryChangeServerRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerRequest) GoString() string {
  return s.String()
}

type QueryChangeServerResponse struct {
  // {"en":"The error code, when HTTPStatus is not 200, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为200时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The response data", "zh_CN":"响应数据"}
  Data *QueryChangeServerQueryChangeServerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s QueryChangeServerResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerResponse) GoString() string {
  return s.String()
}

func (s *QueryChangeServerResponse) SetCode(v string) *QueryChangeServerResponse {
  s.Code = &v
  return s
}

func (s *QueryChangeServerResponse) SetMessage(v string) *QueryChangeServerResponse {
  s.Message = &v
  return s
}

func (s *QueryChangeServerResponse) SetData(v *QueryChangeServerQueryChangeServerResponseData) *QueryChangeServerResponse {
  s.Data = v
  return s
}

func (s *QueryChangeServerResponse) SetXCncRequestId(v string) *QueryChangeServerResponse {
  s.XCncRequestId = &v
  return s
}

type QueryChangeServerQueryChangeServerResponseData struct {
  // {"en":"domain id.", "zh_CN":"域名id"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domain name.", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Change servers configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <change-servers/>", "zh_CN":"【接入域名跳转】
  // 注意：
  // 1、需要取消【接入域名跳转】时，可以传入空节点<change-servers></change-servers>。
  // 2、表示需要设置【接入域名跳转】，此项必填"}
  ChangeServers []*QueryChangeServerQueryChangeServerResponseDataChangeServers `json:"change-servers,omitempty" xml:"change-servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryChangeServerQueryChangeServerResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerQueryChangeServerResponseData) GoString() string {
  return s.String()
}

func (s *QueryChangeServerQueryChangeServerResponseData) SetDomainId(v string) *QueryChangeServerQueryChangeServerResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryChangeServerQueryChangeServerResponseData) SetDomainName(v string) *QueryChangeServerQueryChangeServerResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryChangeServerQueryChangeServerResponseData) SetChangeServers(v []*QueryChangeServerQueryChangeServerResponseDataChangeServers) *QueryChangeServerQueryChangeServerResponseData {
  s.ChangeServers = v
  return s
}

type QueryChangeServerQueryChangeServerResponseDataChangeServers struct     {
  // {"en":"If it is a universal domain name, set it to a universal domain name, for example, *.56.com.", "zh_CN":"如果是泛域名，需要填写为泛域名，例如：*.56.com"}
  TargetServer *string `json:"change-server,omitempty" xml:"change-server,omitempty" require:"true"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryChangeServerQueryChangeServerResponseDataChangeServers) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerQueryChangeServerResponseDataChangeServers) GoString() string {
  return s.String()
}

func (s *QueryChangeServerQueryChangeServerResponseDataChangeServers) SetTargetServer(v string) *QueryChangeServerQueryChangeServerResponseDataChangeServers {
  s.TargetServer = &v
  return s
}

func (s *QueryChangeServerQueryChangeServerResponseDataChangeServers) SetDataId(v int) *QueryChangeServerQueryChangeServerResponseDataChangeServers {
  s.DataId = &v
  return s
}

type QueryChangeServerPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryChangeServerPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerPaths) GoString() string {
  return s.String()
}

func (s *QueryChangeServerPaths) SetDomain(v string) *QueryChangeServerPaths {
  s.Domain = &v
  return s
}

type QueryChangeServerParameters struct {
}

func (s QueryChangeServerParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerParameters) GoString() string {
  return s.String()
}

type QueryChangeServerRequestHeader struct {
}

func (s QueryChangeServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerRequestHeader) GoString() string {
  return s.String()
}

type QueryChangeServerResponseHeader struct {
}

func (s QueryChangeServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerResponseHeader) GoString() string {
  return s.String()
}




