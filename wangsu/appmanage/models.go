package appmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreateAppRequest struct {
  // {"en":"Debug fingerprint, only applicable to Android. If it matches the production fingerprint, it can be left blank.","zh_CN":"调试指纹，仅对Android生效，如果与正式指纹一致可不填写。"}
  DebugFingerprintList []*CreateAppRequestDebugFingerprintList `json:"debugFingerprintList,omitempty" xml:"debugFingerprintList,omitempty" type:"Repeated"`
  // {"en":"Application name should not exceed 60 characters and does not support characters ',\",<,>,&,/.","zh_CN":"应用名称，长度不超过60，不支持字符',\",<,>,&,/。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Official fingerprint, mandatory for Android applications. Please ensure this is the fingerprint for the officially launched application, otherwise authentication will fail and the service cannot be activated. You can check with technical support for the method to obtain the fingerprint. The total length should not exceed 100 characters and must include colons, uppercase and lowercase letters from A to F, and numbers from 0 to 9. After removing the colons, the length should be between 32 and 64.","zh_CN":"正式指纹，Android应用必填，请确保该指纹为正式上线应用的指纹，否则将鉴权失败无法启用服务。获取指纹方法可找技术支持确认。总长度不超过100，仅包含冒号、A~F大小写字母和0~9的数字，去除冒号后长度介于32~64之间。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
  // {"en":"Please fill in the package name carefully; otherwise, it will not be usable. The length must be between 3 and 100 characters, consisting of letters, numbers, underscores, and hyphens, with sections of the package name separated by periods. Format example: com.maa.test.","zh_CN":"包名，请认真填写包名，否则将无法使用。长度必须在 3 到 100 个字符之间，并且由字母、数字、下划线和连字符组成，包名的各部分之间用点号分隔。格式如：com.maa.test。"}
  PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
  // {"dictionary":"belong=MAA-masp-portal-console|dict=AppType","en":"Application Type.","zh_CN":"应用类型。"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Application Platform, Android or iOS.","zh_CN":"应用平台，Android、iOS。"}
  Platform *string `json:"platform,omitempty" xml:"platform,omitempty" require:"true"`
}

func (s CreateAppRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
  return s.String()
}

func (s *CreateAppRequest) SetDebugFingerprintList(v []*CreateAppRequestDebugFingerprintList) *CreateAppRequest {
  s.DebugFingerprintList = v
  return s
}

func (s *CreateAppRequest) SetName(v string) *CreateAppRequest {
  s.Name = &v
  return s
}

func (s *CreateAppRequest) SetFingerprint(v string) *CreateAppRequest {
  s.Fingerprint = &v
  return s
}

func (s *CreateAppRequest) SetPackageName(v string) *CreateAppRequest {
  s.PackageName = &v
  return s
}

func (s *CreateAppRequest) SetType(v int) *CreateAppRequest {
  s.Type = &v
  return s
}

func (s *CreateAppRequest) SetPlatform(v string) *CreateAppRequest {
  s.Platform = &v
  return s
}

type CreateAppRequestDebugFingerprintList struct     {
  // {"en":"Fingerprint, with a total length not exceeding 100, containing only colons, letters A to F in both uppercase and lowercase, and digits 0 to 9. After removing colons, the length should be between 32 and 64.","zh_CN":"指纹，总长度不超过100，仅包含冒号、A~F大小写字母和0~9的数字，去除冒号后长度介于32~64之间。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
  // {"en":"Description, no more than 60 characters.","zh_CN":"描述，长度不超过60。"}
  Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s CreateAppRequestDebugFingerprintList) String() string {
  return tea.Prettify(s)
}

func (s CreateAppRequestDebugFingerprintList) GoString() string {
  return s.String()
}

func (s *CreateAppRequestDebugFingerprintList) SetFingerprint(v string) *CreateAppRequestDebugFingerprintList {
  s.Fingerprint = &v
  return s
}

func (s *CreateAppRequestDebugFingerprintList) SetDesc(v string) *CreateAppRequestDebugFingerprintList {
  s.Desc = &v
  return s
}

type CreateAppRequestHeader struct {
}

func (s CreateAppRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppRequestHeader) GoString() string {
  return s.String()
}

type CreateAppPaths struct {
}

func (s CreateAppPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAppPaths) GoString() string {
  return s.String()
}

type CreateAppParameters struct {
}

func (s CreateAppParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAppParameters) GoString() string {
  return s.String()
}

type CreateAppResponse struct {
  // {"dictionary":"belong=MAA-masp-portal-console|dict=wplus_code","en":"Response Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return data, Application ID","zh_CN":"返回数据，应用ID"}
  Data *int `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Response Description","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
  return s.String()
}

func (s *CreateAppResponse) SetCode(v string) *CreateAppResponse {
  s.Code = &v
  return s
}

func (s *CreateAppResponse) SetData(v int) *CreateAppResponse {
  s.Data = &v
  return s
}

func (s *CreateAppResponse) SetMessage(v string) *CreateAppResponse {
  s.Message = &v
  return s
}

type CreateAppResponseHeader struct {
}

func (s CreateAppResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppResponseHeader) GoString() string {
  return s.String()
}




type UpdateTunnelApplicationsRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true" maxLength:"128"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*UpdateTunnelApplicationsContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" type:"Repeated"`
}

func (s UpdateTunnelApplicationsRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsRequest) GoString() string {
  return s.String()
}

func (s *UpdateTunnelApplicationsRequest) SetName(v string) *UpdateTunnelApplicationsRequest {
  s.Name = &v
  return s
}

func (s *UpdateTunnelApplicationsRequest) SetRemark(v string) *UpdateTunnelApplicationsRequest {
  s.Remark = &v
  return s
}

func (s *UpdateTunnelApplicationsRequest) SetContentList(v []*UpdateTunnelApplicationsContentItem) *UpdateTunnelApplicationsRequest {
  s.ContentList = v
  return s
}

type UpdateTunnelApplicationsContentItem struct {
  // {"en":"protocol,support for tcp,udp,icmp,all", "zh_CN":"协议,支持tcp、udp、icmp、all.  all表示不限制协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"ip or domain", "zh_CN":"ip或者域名"}
  IpOrDomain *string `json:"ipOrDomain,omitempty" xml:"ipOrDomain,omitempty" require:"true"`
  // {"en":"application ports", "zh_CN":"端口,支持IP地址(IP、IP/掩码、IPX-IPy)和域名(精确域名、泛域),多个以英文分号间隔,最大支持2万个"}
  Ports *string `json:"ports,omitempty" xml:"ports,omitempty" require:"true"`
}

func (s UpdateTunnelApplicationsContentItem) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsContentItem) GoString() string {
  return s.String()
}

func (s *UpdateTunnelApplicationsContentItem) SetProtocol(v string) *UpdateTunnelApplicationsContentItem {
  s.Protocol = &v
  return s
}

func (s *UpdateTunnelApplicationsContentItem) SetIpOrDomain(v string) *UpdateTunnelApplicationsContentItem {
  s.IpOrDomain = &v
  return s
}

func (s *UpdateTunnelApplicationsContentItem) SetPorts(v string) *UpdateTunnelApplicationsContentItem {
  s.Ports = &v
  return s
}

type UpdateTunnelApplicationsResponse struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*UpdateTunnelApplicationsContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" type:"Repeated"`
}

func (s UpdateTunnelApplicationsResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsResponse) GoString() string {
  return s.String()
}

func (s *UpdateTunnelApplicationsResponse) SetName(v string) *UpdateTunnelApplicationsResponse {
  s.Name = &v
  return s
}

func (s *UpdateTunnelApplicationsResponse) SetRemark(v string) *UpdateTunnelApplicationsResponse {
  s.Remark = &v
  return s
}

func (s *UpdateTunnelApplicationsResponse) SetContentList(v []*UpdateTunnelApplicationsContentItem) *UpdateTunnelApplicationsResponse {
  s.ContentList = v
  return s
}

type UpdateTunnelApplicationsPaths struct {
}

func (s UpdateTunnelApplicationsPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsPaths) GoString() string {
  return s.String()
}

type UpdateTunnelApplicationsParameters struct {
}

func (s UpdateTunnelApplicationsParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsParameters) GoString() string {
  return s.String()
}

type UpdateTunnelApplicationsRequestHeader struct {
}

func (s UpdateTunnelApplicationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsRequestHeader) GoString() string {
  return s.String()
}

type UpdateTunnelApplicationsResponseHeader struct {
}

func (s UpdateTunnelApplicationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateTunnelApplicationsResponseHeader) GoString() string {
  return s.String()
}




type AddDebugFingerprintRequest struct {
  // {"en":"Application ID","zh_CN":"应用ID"}
  AppId *int `json:"appId,omitempty" xml:"appId,omitempty" require:"true"`
  // {"en":"Debug fingerprint, only applicable to Android.","zh_CN":"调试指纹，仅对Android生效。"}
  DebugFingerprintList []*AddDebugFingerprintRequestDebugFingerprintList `json:"debugFingerprintList,omitempty" xml:"debugFingerprintList,omitempty" require:"true" type:"Repeated"`
}

func (s AddDebugFingerprintRequest) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintRequest) GoString() string {
  return s.String()
}

func (s *AddDebugFingerprintRequest) SetAppId(v int) *AddDebugFingerprintRequest {
  s.AppId = &v
  return s
}

func (s *AddDebugFingerprintRequest) SetDebugFingerprintList(v []*AddDebugFingerprintRequestDebugFingerprintList) *AddDebugFingerprintRequest {
  s.DebugFingerprintList = v
  return s
}

type AddDebugFingerprintRequestDebugFingerprintList struct     {
  // {"en":"Fingerprint, with a total length not exceeding 100, containing only colons, letters A to F in both uppercase and lowercase, and digits 0 to 9. After removing colons, the length should be between 32 and 64.","zh_CN":"指纹，总长度不超过100，仅包含冒号、A~F大小写字母和0~9的数字，去除冒号后长度介于32~64之间。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty" require:"true"`
  // {"en":"Description, no more than 60 characters.","zh_CN":"描述，长度不超过60。"}
  Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s AddDebugFingerprintRequestDebugFingerprintList) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintRequestDebugFingerprintList) GoString() string {
  return s.String()
}

func (s *AddDebugFingerprintRequestDebugFingerprintList) SetFingerprint(v string) *AddDebugFingerprintRequestDebugFingerprintList {
  s.Fingerprint = &v
  return s
}

func (s *AddDebugFingerprintRequestDebugFingerprintList) SetDesc(v string) *AddDebugFingerprintRequestDebugFingerprintList {
  s.Desc = &v
  return s
}

type AddDebugFingerprintRequestHeader struct {
}

func (s AddDebugFingerprintRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintRequestHeader) GoString() string {
  return s.String()
}

type AddDebugFingerprintPaths struct {
}

func (s AddDebugFingerprintPaths) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintPaths) GoString() string {
  return s.String()
}

type AddDebugFingerprintParameters struct {
}

func (s AddDebugFingerprintParameters) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintParameters) GoString() string {
  return s.String()
}

type AddDebugFingerprintResponse struct {
  // {"dictionary":"belong=MAA-masp-portal-console|dict=wplus_code","en":"Response Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response Description","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddDebugFingerprintResponse) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintResponse) GoString() string {
  return s.String()
}

func (s *AddDebugFingerprintResponse) SetCode(v string) *AddDebugFingerprintResponse {
  s.Code = &v
  return s
}

func (s *AddDebugFingerprintResponse) SetMessage(v string) *AddDebugFingerprintResponse {
  s.Message = &v
  return s
}

type AddDebugFingerprintResponseHeader struct {
}

func (s AddDebugFingerprintResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintResponseHeader) GoString() string {
  return s.String()
}




type AuthorizeUserApplicationRequest struct {
  // {"en":"Resource/Application type, default is tunnel. tunnel: tunnelApp, web: webApp, link: linkApp, saas: saasApp", "zh_CN":"资源/应用类型，默认是隧道应用. tunnel: 隧道应用, web: WEB应用, link: 快捷链接, saas: saas应用"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
  // {"en":"List of resource/application names", "zh_CN":"资源/应用名称列表"}
  ResourceNames []*string `json:"resourceNames,omitempty" xml:"resourceNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Action Type, 0: Append, 1: Overwrite", "zh_CN":"操作类型，0：追加，1：覆盖"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"List of authorized users", "zh_CN":"授权的用户列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"List of authorized user group IDs", "zh_CN":"授权的用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s AuthorizeUserApplicationRequest) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeUserApplicationRequest) GoString() string {
  return s.String()
}

func (s *AuthorizeUserApplicationRequest) SetResourceType(v string) *AuthorizeUserApplicationRequest {
  s.ResourceType = &v
  return s
}

func (s *AuthorizeUserApplicationRequest) SetResourceNames(v []*string) *AuthorizeUserApplicationRequest {
  s.ResourceNames = v
  return s
}

func (s *AuthorizeUserApplicationRequest) SetActionType(v int) *AuthorizeUserApplicationRequest {
  s.ActionType = &v
  return s
}

func (s *AuthorizeUserApplicationRequest) SetAuthorizedUsers(v []*string) *AuthorizeUserApplicationRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *AuthorizeUserApplicationRequest) SetAuthorizedUserGroupIds(v []*int64) *AuthorizeUserApplicationRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

type AuthorizeUserApplicationResponse struct {
  // {"en":"Resource/Application type.tunnel: tunnelApp, web: webApp, link: linkApp, saas: saasApp", "zh_CN":"资源/应用类型 tunnel: 隧道应用, web: WEB应用, link: 快捷链接, saas: saas应用"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
  // {"en":"List of resource/application names", "zh_CN":"资源/应用名称列表"}
  ResourceNames []*string `json:"resourceNames,omitempty" xml:"resourceNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Action Type, 0: Append, 1: Overwrite", "zh_CN":"操作类型，0：追加，1：覆盖"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"List of authorized users", "zh_CN":"授权的用户列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"List of authorized user group IDs", "zh_CN":"授权的用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s AuthorizeUserApplicationResponse) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeUserApplicationResponse) GoString() string {
  return s.String()
}

func (s *AuthorizeUserApplicationResponse) SetResourceType(v string) *AuthorizeUserApplicationResponse {
  s.ResourceType = &v
  return s
}

func (s *AuthorizeUserApplicationResponse) SetResourceNames(v []*string) *AuthorizeUserApplicationResponse {
  s.ResourceNames = v
  return s
}

func (s *AuthorizeUserApplicationResponse) SetActionType(v int) *AuthorizeUserApplicationResponse {
  s.ActionType = &v
  return s
}

func (s *AuthorizeUserApplicationResponse) SetAuthorizedUsers(v []*string) *AuthorizeUserApplicationResponse {
  s.AuthorizedUsers = v
  return s
}

func (s *AuthorizeUserApplicationResponse) SetAuthorizedUserGroupIds(v []*int64) *AuthorizeUserApplicationResponse {
  s.AuthorizedUserGroupIds = v
  return s
}

type AuthorizeUserApplicationPaths struct {
}

func (s AuthorizeUserApplicationPaths) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeUserApplicationPaths) GoString() string {
  return s.String()
}

type AuthorizeUserApplicationParameters struct {
}

func (s AuthorizeUserApplicationParameters) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeUserApplicationParameters) GoString() string {
  return s.String()
}

type AuthorizeUserApplicationRequestHeader struct {
}

func (s AuthorizeUserApplicationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeUserApplicationRequestHeader) GoString() string {
  return s.String()
}

type AuthorizeUserApplicationResponseHeader struct {
}

func (s AuthorizeUserApplicationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeUserApplicationResponseHeader) GoString() string {
  return s.String()
}




type AddWebApplicationsRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true" maxLength:"128"`
  // {"en":"protocol", "zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"ip or domain", "zh_CN":"ip或者域名"}
  IpOrDomain *string `json:"ipOrDomain,omitempty" xml:"ipOrDomain,omitempty" require:"true"`
  // {"en":"application ports", "zh_CN":"端口:0-65535"}
  Port *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
}

func (s AddWebApplicationsRequest) String() string {
  return tea.Prettify(s)
}

func (s AddWebApplicationsRequest) GoString() string {
  return s.String()
}

func (s *AddWebApplicationsRequest) SetName(v string) *AddWebApplicationsRequest {
  s.Name = &v
  return s
}

func (s *AddWebApplicationsRequest) SetProtocol(v string) *AddWebApplicationsRequest {
  s.Protocol = &v
  return s
}

func (s *AddWebApplicationsRequest) SetIpOrDomain(v string) *AddWebApplicationsRequest {
  s.IpOrDomain = &v
  return s
}

func (s *AddWebApplicationsRequest) SetPort(v string) *AddWebApplicationsRequest {
  s.Port = &v
  return s
}

func (s *AddWebApplicationsRequest) SetRemark(v string) *AddWebApplicationsRequest {
  s.Remark = &v
  return s
}

type AddWebApplicationsResponse struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"protocol", "zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"ip or domain", "zh_CN":"ip或者域名"}
  IpOrDomain *string `json:"ipOrDomain,omitempty" xml:"ipOrDomain,omitempty" require:"true"`
  // {"en":"application ports", "zh_CN":"端口:0-65535"}
  Port *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
}

func (s AddWebApplicationsResponse) String() string {
  return tea.Prettify(s)
}

func (s AddWebApplicationsResponse) GoString() string {
  return s.String()
}

func (s *AddWebApplicationsResponse) SetName(v string) *AddWebApplicationsResponse {
  s.Name = &v
  return s
}

func (s *AddWebApplicationsResponse) SetProtocol(v string) *AddWebApplicationsResponse {
  s.Protocol = &v
  return s
}

func (s *AddWebApplicationsResponse) SetIpOrDomain(v string) *AddWebApplicationsResponse {
  s.IpOrDomain = &v
  return s
}

func (s *AddWebApplicationsResponse) SetPort(v string) *AddWebApplicationsResponse {
  s.Port = &v
  return s
}

func (s *AddWebApplicationsResponse) SetRemark(v string) *AddWebApplicationsResponse {
  s.Remark = &v
  return s
}

type AddWebApplicationsPaths struct {
}

func (s AddWebApplicationsPaths) String() string {
  return tea.Prettify(s)
}

func (s AddWebApplicationsPaths) GoString() string {
  return s.String()
}

type AddWebApplicationsParameters struct {
}

func (s AddWebApplicationsParameters) String() string {
  return tea.Prettify(s)
}

func (s AddWebApplicationsParameters) GoString() string {
  return s.String()
}

type AddWebApplicationsRequestHeader struct {
}

func (s AddWebApplicationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddWebApplicationsRequestHeader) GoString() string {
  return s.String()
}

type AddWebApplicationsResponseHeader struct {
}

func (s AddWebApplicationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddWebApplicationsResponseHeader) GoString() string {
  return s.String()
}




type QueryAppListRequest struct {
}

func (s QueryAppListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListRequest) GoString() string {
  return s.String()
}

type QueryAppListResponse struct {
  // {"en":"app list", "zh_CN":"应用列表"}
  Apps []*QueryAppListAppInfo `json:"apps,omitempty" xml:"apps,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s QueryAppListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListResponse) GoString() string {
  return s.String()
}

func (s *QueryAppListResponse) SetApps(v []*QueryAppListAppInfo) *QueryAppListResponse {
  s.Apps = v
  return s
}

func (s *QueryAppListResponse) SetStatus(v int) *QueryAppListResponse {
  s.Status = &v
  return s
}

func (s *QueryAppListResponse) SetResult(v string) *QueryAppListResponse {
  s.Result = &v
  return s
}

type QueryAppListAppInfo struct {
  // {"en":"app name", "zh_CN":"应用名称"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {"en":"app package name", "zh_CN":"应用包名"}
  AppPackage *string `json:"appPackage,omitempty" xml:"appPackage,omitempty" require:"true"`
  // {"en":"app version", "zh_CN":"应用版本"}
  AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty" require:"true"`
  // {"en":"app size", "zh_CN":"应用大小"}
  AppSize *string `json:"appSize,omitempty" xml:"appSize,omitempty" require:"true"`
  // {"en":"app update/install time", "zh_CN":"应用更新/安装时间"}
  AppUpdateTime *string `json:"appUpdateTime,omitempty" xml:"appUpdateTime,omitempty" require:"true"`
}

func (s QueryAppListAppInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListAppInfo) GoString() string {
  return s.String()
}

func (s *QueryAppListAppInfo) SetAppName(v string) *QueryAppListAppInfo {
  s.AppName = &v
  return s
}

func (s *QueryAppListAppInfo) SetAppPackage(v string) *QueryAppListAppInfo {
  s.AppPackage = &v
  return s
}

func (s *QueryAppListAppInfo) SetAppVersion(v string) *QueryAppListAppInfo {
  s.AppVersion = &v
  return s
}

func (s *QueryAppListAppInfo) SetAppSize(v string) *QueryAppListAppInfo {
  s.AppSize = &v
  return s
}

func (s *QueryAppListAppInfo) SetAppUpdateTime(v string) *QueryAppListAppInfo {
  s.AppUpdateTime = &v
  return s
}

type QueryAppListPaths struct {
}

func (s QueryAppListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListPaths) GoString() string {
  return s.String()
}

type QueryAppListParameters struct {
  // {"en":"no", "zh_CN":"云手机id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryAppListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListParameters) GoString() string {
  return s.String()
}

func (s *QueryAppListParameters) SetId(v string) *QueryAppListParameters {
  s.Id = &v
  return s
}

type QueryAppListRequestHeader struct {
}

func (s QueryAppListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListRequestHeader) GoString() string {
  return s.String()
}

type QueryAppListResponseHeader struct {
}

func (s QueryAppListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppListResponseHeader) GoString() string {
  return s.String()
}




type DescribeAuthorizedApplicationsOfUserRequest struct {
  // {"en":"User Name", "zh_CN":"用户名"}
  Username *string `json:"username,omitempty" xml:"username,omitempty" require:"true"`
}

func (s DescribeAuthorizedApplicationsOfUserRequest) String() string {
  return tea.Prettify(s)
}

func (s DescribeAuthorizedApplicationsOfUserRequest) GoString() string {
  return s.String()
}

func (s *DescribeAuthorizedApplicationsOfUserRequest) SetUsername(v string) *DescribeAuthorizedApplicationsOfUserRequest {
  s.Username = &v
  return s
}

type DescribeAuthorizedApplicationsOfUserResponse struct {
  // {"en":"Resource/Application type, default is tunnel. tunnel: tunnelApp, web: webApp, link: linkApp, saas: saasApp", "zh_CN":"资源/应用类型，默认是隧道应用. tunnel: 隧道应用, web: WEB应用, link: 快捷链接, saas: saas应用"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
  // {"en":"Resource/Application name", "zh_CN":"资源/应用名称"}
  ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty" require:"true"`
}

func (s DescribeAuthorizedApplicationsOfUserResponse) String() string {
  return tea.Prettify(s)
}

func (s DescribeAuthorizedApplicationsOfUserResponse) GoString() string {
  return s.String()
}

func (s *DescribeAuthorizedApplicationsOfUserResponse) SetResourceType(v string) *DescribeAuthorizedApplicationsOfUserResponse {
  s.ResourceType = &v
  return s
}

func (s *DescribeAuthorizedApplicationsOfUserResponse) SetResourceName(v string) *DescribeAuthorizedApplicationsOfUserResponse {
  s.ResourceName = &v
  return s
}

type DescribeAuthorizedApplicationsOfUserPaths struct {
}

func (s DescribeAuthorizedApplicationsOfUserPaths) String() string {
  return tea.Prettify(s)
}

func (s DescribeAuthorizedApplicationsOfUserPaths) GoString() string {
  return s.String()
}

type DescribeAuthorizedApplicationsOfUserParameters struct {
}

func (s DescribeAuthorizedApplicationsOfUserParameters) String() string {
  return tea.Prettify(s)
}

func (s DescribeAuthorizedApplicationsOfUserParameters) GoString() string {
  return s.String()
}

type DescribeAuthorizedApplicationsOfUserRequestHeader struct {
}

func (s DescribeAuthorizedApplicationsOfUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DescribeAuthorizedApplicationsOfUserRequestHeader) GoString() string {
  return s.String()
}

type DescribeAuthorizedApplicationsOfUserResponseHeader struct {
}

func (s DescribeAuthorizedApplicationsOfUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DescribeAuthorizedApplicationsOfUserResponseHeader) GoString() string {
  return s.String()
}




type DeleteAppRequest struct {
  // {"en":"Application ID","zh_CN":"应用ID"}
  AppId *int `json:"appId,omitempty" xml:"appId,omitempty" require:"true"`
}

func (s DeleteAppRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
  return s.String()
}

func (s *DeleteAppRequest) SetAppId(v int) *DeleteAppRequest {
  s.AppId = &v
  return s
}

type DeleteAppRequestHeader struct {
}

func (s DeleteAppRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppRequestHeader) GoString() string {
  return s.String()
}

type DeleteAppPaths struct {
}

func (s DeleteAppPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppPaths) GoString() string {
  return s.String()
}

type DeleteAppParameters struct {
}

func (s DeleteAppParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppParameters) GoString() string {
  return s.String()
}

type DeleteAppResponse struct {
  // {"dictionary":"belong=MAA-masp-portal-console|dict=wplus_code","en":"Response Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response Description","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
  return s.String()
}

func (s *DeleteAppResponse) SetCode(v string) *DeleteAppResponse {
  s.Code = &v
  return s
}

func (s *DeleteAppResponse) SetMessage(v string) *DeleteAppResponse {
  s.Message = &v
  return s
}

type DeleteAppResponseHeader struct {
}

func (s DeleteAppResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppResponseHeader) GoString() string {
  return s.String()
}




type QueryApplicationDetailRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*QueryApplicationDetailContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" type:"Repeated"`
}

func (s QueryApplicationDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryApplicationDetailRequest) SetName(v string) *QueryApplicationDetailRequest {
  s.Name = &v
  return s
}

func (s *QueryApplicationDetailRequest) SetRemark(v string) *QueryApplicationDetailRequest {
  s.Remark = &v
  return s
}

func (s *QueryApplicationDetailRequest) SetContentList(v []*QueryApplicationDetailContentItem) *QueryApplicationDetailRequest {
  s.ContentList = v
  return s
}

type QueryApplicationDetailContentItem struct {
  // {"en":"application name", "zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"ip or domain", "zh_CN":"ip或者域名"}
  IpOrDomain *string `json:"ipOrDomain,omitempty" xml:"ipOrDomain,omitempty" require:"true"`
  // {"en":"application ports", "zh_CN":"端口,支持IP地址(IP、IP/掩码、IPX-IPy)和域名(精确域名、泛域),多个以英文分号间隔,最大支持2万个"}
  Ports *string `json:"ports,omitempty" xml:"ports,omitempty" require:"true"`
}

func (s QueryApplicationDetailContentItem) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailContentItem) GoString() string {
  return s.String()
}

func (s *QueryApplicationDetailContentItem) SetProtocol(v string) *QueryApplicationDetailContentItem {
  s.Protocol = &v
  return s
}

func (s *QueryApplicationDetailContentItem) SetIpOrDomain(v string) *QueryApplicationDetailContentItem {
  s.IpOrDomain = &v
  return s
}

func (s *QueryApplicationDetailContentItem) SetPorts(v string) *QueryApplicationDetailContentItem {
  s.Ports = &v
  return s
}

type QueryApplicationDetailResponse struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*QueryApplicationDetailContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" type:"Repeated"`
}

func (s QueryApplicationDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryApplicationDetailResponse) SetName(v string) *QueryApplicationDetailResponse {
  s.Name = &v
  return s
}

func (s *QueryApplicationDetailResponse) SetRemark(v string) *QueryApplicationDetailResponse {
  s.Remark = &v
  return s
}

func (s *QueryApplicationDetailResponse) SetContentList(v []*QueryApplicationDetailContentItem) *QueryApplicationDetailResponse {
  s.ContentList = v
  return s
}

type QueryApplicationDetailPaths struct {
}

func (s QueryApplicationDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailPaths) GoString() string {
  return s.String()
}

type QueryApplicationDetailParameters struct {
}

func (s QueryApplicationDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailParameters) GoString() string {
  return s.String()
}

type QueryApplicationDetailRequestHeader struct {
}

func (s QueryApplicationDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailRequestHeader) GoString() string {
  return s.String()
}

type QueryApplicationDetailResponseHeader struct {
}

func (s QueryApplicationDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApplicationDetailResponseHeader) GoString() string {
  return s.String()
}




type AuthorizeApplicationForUserRequest struct {
  // {"en":"Resource/Application Name", "zh_CN":"资源/应用名称"}
  ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty" require:"true"`
  // {"en":"Action Type, 0: Append, 1: Overwrite", "zh_CN":"操作类型，0：追加，1：覆盖"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"Resource Type, e.g., tunnel, web, link, saas", "zh_CN":"资源类型，比如：隧道应用、WEB应用、快捷链接、saas应用"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
  // {"en":"List of authorized users", "zh_CN":"授权的用户列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"List of authorized user group IDs", "zh_CN":"授权的用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
  // {"en":"List of excluded users", "zh_CN":"例外用户列表"}
  ExcludeUsers []*string `json:"excludeUsers,omitempty" xml:"excludeUsers,omitempty" type:"Repeated"`
  // {"en":"List of excluded user group IDs", "zh_CN":"例外用户组ID列表"}
  ExcludeUserGroupIds []*int64 `json:"excludeUserGroupIds,omitempty" xml:"excludeUserGroupIds,omitempty" type:"Repeated"`
}

func (s AuthorizeApplicationForUserRequest) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeApplicationForUserRequest) GoString() string {
  return s.String()
}

func (s *AuthorizeApplicationForUserRequest) SetResourceName(v string) *AuthorizeApplicationForUserRequest {
  s.ResourceName = &v
  return s
}

func (s *AuthorizeApplicationForUserRequest) SetActionType(v int) *AuthorizeApplicationForUserRequest {
  s.ActionType = &v
  return s
}

func (s *AuthorizeApplicationForUserRequest) SetResourceType(v string) *AuthorizeApplicationForUserRequest {
  s.ResourceType = &v
  return s
}

func (s *AuthorizeApplicationForUserRequest) SetAuthorizedUsers(v []*string) *AuthorizeApplicationForUserRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *AuthorizeApplicationForUserRequest) SetAuthorizedUserGroupIds(v []*int64) *AuthorizeApplicationForUserRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

func (s *AuthorizeApplicationForUserRequest) SetExcludeUsers(v []*string) *AuthorizeApplicationForUserRequest {
  s.ExcludeUsers = v
  return s
}

func (s *AuthorizeApplicationForUserRequest) SetExcludeUserGroupIds(v []*int64) *AuthorizeApplicationForUserRequest {
  s.ExcludeUserGroupIds = v
  return s
}

type AuthorizeApplicationForUserResponse struct {
}

func (s AuthorizeApplicationForUserResponse) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeApplicationForUserResponse) GoString() string {
  return s.String()
}

type AuthorizeApplicationForUserPaths struct {
}

func (s AuthorizeApplicationForUserPaths) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeApplicationForUserPaths) GoString() string {
  return s.String()
}

type AuthorizeApplicationForUserParameters struct {
}

func (s AuthorizeApplicationForUserParameters) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeApplicationForUserParameters) GoString() string {
  return s.String()
}

type AuthorizeApplicationForUserRequestHeader struct {
}

func (s AuthorizeApplicationForUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeApplicationForUserRequestHeader) GoString() string {
  return s.String()
}

type AuthorizeApplicationForUserResponseHeader struct {
}

func (s AuthorizeApplicationForUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AuthorizeApplicationForUserResponseHeader) GoString() string {
  return s.String()
}




type DeleteWebApplicationsRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteWebApplicationsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebApplicationsRequest) GoString() string {
  return s.String()
}

func (s *DeleteWebApplicationsRequest) SetName(v string) *DeleteWebApplicationsRequest {
  s.Name = &v
  return s
}

type DeleteWebApplicationsResponse struct {
}

func (s DeleteWebApplicationsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebApplicationsResponse) GoString() string {
  return s.String()
}

type DeleteWebApplicationsPaths struct {
}

func (s DeleteWebApplicationsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebApplicationsPaths) GoString() string {
  return s.String()
}

type DeleteWebApplicationsParameters struct {
}

func (s DeleteWebApplicationsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebApplicationsParameters) GoString() string {
  return s.String()
}

type DeleteWebApplicationsRequestHeader struct {
}

func (s DeleteWebApplicationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebApplicationsRequestHeader) GoString() string {
  return s.String()
}

type DeleteWebApplicationsResponseHeader struct {
}

func (s DeleteWebApplicationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebApplicationsResponseHeader) GoString() string {
  return s.String()
}




type AddTunnelApplicationsRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true" maxLength:"128"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*AddTunnelApplicationsContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
}

func (s AddTunnelApplicationsRequest) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsRequest) GoString() string {
  return s.String()
}

func (s *AddTunnelApplicationsRequest) SetName(v string) *AddTunnelApplicationsRequest {
  s.Name = &v
  return s
}

func (s *AddTunnelApplicationsRequest) SetRemark(v string) *AddTunnelApplicationsRequest {
  s.Remark = &v
  return s
}

func (s *AddTunnelApplicationsRequest) SetContentList(v []*AddTunnelApplicationsContentItem) *AddTunnelApplicationsRequest {
  s.ContentList = v
  return s
}

type AddTunnelApplicationsContentItem struct {
  // {"en":"protocol,support for tcp,udp,imcp,all.", "zh_CN":"协议,支持tcp,udp,icmp,all. all表示不限制协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"ip or domain", "zh_CN":"ip或者域名"}
  IpOrDomain *string `json:"ipOrDomain,omitempty" xml:"ipOrDomain,omitempty" require:"true"`
  // {"en":"application ports", "zh_CN":"端口,支持IP地址(IP、IP/掩码、IPX-IPy)和域名(精确域名、泛域),多个以英文分号间隔,最大支持2万个"}
  Ports *string `json:"ports,omitempty" xml:"ports,omitempty" require:"true"`
}

func (s AddTunnelApplicationsContentItem) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsContentItem) GoString() string {
  return s.String()
}

func (s *AddTunnelApplicationsContentItem) SetProtocol(v string) *AddTunnelApplicationsContentItem {
  s.Protocol = &v
  return s
}

func (s *AddTunnelApplicationsContentItem) SetIpOrDomain(v string) *AddTunnelApplicationsContentItem {
  s.IpOrDomain = &v
  return s
}

func (s *AddTunnelApplicationsContentItem) SetPorts(v string) *AddTunnelApplicationsContentItem {
  s.Ports = &v
  return s
}

type AddTunnelApplicationsResponse struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*AddTunnelApplicationsContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
}

func (s AddTunnelApplicationsResponse) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsResponse) GoString() string {
  return s.String()
}

func (s *AddTunnelApplicationsResponse) SetName(v string) *AddTunnelApplicationsResponse {
  s.Name = &v
  return s
}

func (s *AddTunnelApplicationsResponse) SetRemark(v string) *AddTunnelApplicationsResponse {
  s.Remark = &v
  return s
}

func (s *AddTunnelApplicationsResponse) SetContentList(v []*AddTunnelApplicationsContentItem) *AddTunnelApplicationsResponse {
  s.ContentList = v
  return s
}

type AddTunnelApplicationsPaths struct {
}

func (s AddTunnelApplicationsPaths) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsPaths) GoString() string {
  return s.String()
}

type AddTunnelApplicationsParameters struct {
}

func (s AddTunnelApplicationsParameters) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsParameters) GoString() string {
  return s.String()
}

type AddTunnelApplicationsRequestHeader struct {
}

func (s AddTunnelApplicationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsRequestHeader) GoString() string {
  return s.String()
}

type AddTunnelApplicationsResponseHeader struct {
}

func (s AddTunnelApplicationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTunnelApplicationsResponseHeader) GoString() string {
  return s.String()
}




type UpdateWebApplicationsRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"new  application name", "zh_CN":"新的应用名称"}
  NewName *string `json:"newName,omitempty" xml:"newName,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*UpdateWebApplicationsContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" type:"Repeated"`
}

func (s UpdateWebApplicationsRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsRequest) GoString() string {
  return s.String()
}

func (s *UpdateWebApplicationsRequest) SetName(v string) *UpdateWebApplicationsRequest {
  s.Name = &v
  return s
}

func (s *UpdateWebApplicationsRequest) SetNewName(v string) *UpdateWebApplicationsRequest {
  s.NewName = &v
  return s
}

func (s *UpdateWebApplicationsRequest) SetRemark(v string) *UpdateWebApplicationsRequest {
  s.Remark = &v
  return s
}

func (s *UpdateWebApplicationsRequest) SetContentList(v []*UpdateWebApplicationsContentItem) *UpdateWebApplicationsRequest {
  s.ContentList = v
  return s
}

type UpdateWebApplicationsContentItem struct {
  // {"en":"protocol", "zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"ip or domain", "zh_CN":"ip或者域名"}
  IpOrDomain *string `json:"ipOrDomain,omitempty" xml:"ipOrDomain,omitempty" require:"true"`
  // {"en":"application ports", "zh_CN":"端口,支持IP地址(IP、IP/掩码、IPX-IPy)和域名(精确域名、泛域),多个以英文分号间隔,最大支持2万个"}
  Ports *string `json:"ports,omitempty" xml:"ports,omitempty" require:"true"`
}

func (s UpdateWebApplicationsContentItem) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsContentItem) GoString() string {
  return s.String()
}

func (s *UpdateWebApplicationsContentItem) SetProtocol(v string) *UpdateWebApplicationsContentItem {
  s.Protocol = &v
  return s
}

func (s *UpdateWebApplicationsContentItem) SetIpOrDomain(v string) *UpdateWebApplicationsContentItem {
  s.IpOrDomain = &v
  return s
}

func (s *UpdateWebApplicationsContentItem) SetPorts(v string) *UpdateWebApplicationsContentItem {
  s.Ports = &v
  return s
}

type UpdateWebApplicationsResponse struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"application contents ", "zh_CN":"应用内容"}
  ContentList []*UpdateWebApplicationsContentItem `json:"contentList,omitempty" xml:"contentList,omitempty" type:"Repeated"`
}

func (s UpdateWebApplicationsResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsResponse) GoString() string {
  return s.String()
}

func (s *UpdateWebApplicationsResponse) SetName(v string) *UpdateWebApplicationsResponse {
  s.Name = &v
  return s
}

func (s *UpdateWebApplicationsResponse) SetRemark(v string) *UpdateWebApplicationsResponse {
  s.Remark = &v
  return s
}

func (s *UpdateWebApplicationsResponse) SetContentList(v []*UpdateWebApplicationsContentItem) *UpdateWebApplicationsResponse {
  s.ContentList = v
  return s
}

type UpdateWebApplicationsPaths struct {
}

func (s UpdateWebApplicationsPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsPaths) GoString() string {
  return s.String()
}

type UpdateWebApplicationsParameters struct {
}

func (s UpdateWebApplicationsParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsParameters) GoString() string {
  return s.String()
}

type UpdateWebApplicationsRequestHeader struct {
}

func (s UpdateWebApplicationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsRequestHeader) GoString() string {
  return s.String()
}

type UpdateWebApplicationsResponseHeader struct {
}

func (s UpdateWebApplicationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebApplicationsResponseHeader) GoString() string {
  return s.String()
}




type ManageEphoneAppRequest struct {
  // {"en":"list of apps to operate on", "zh_CN":"操作的云手机应用数组"}
  Apps []*ManageEphoneAppOperateObject `json:"apps,omitempty" xml:"apps,omitempty" require:"true" type:"Repeated"`
}

func (s ManageEphoneAppRequest) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppRequest) GoString() string {
  return s.String()
}

func (s *ManageEphoneAppRequest) SetApps(v []*ManageEphoneAppOperateObject) *ManageEphoneAppRequest {
  s.Apps = v
  return s
}

type ManageEphoneAppOperateObject struct {
  // {"en":"ephone instance id", "zh_CN":"云手机实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"operate type", "zh_CN":"操作类型，可选值：install, uninstall, start, stop, clean"}
  Op *string `json:"op,omitempty" xml:"op,omitempty" require:"true"`
  // {"en":"app operaate params", "zh_CN":"app 操作的参数"}
  Params *ManageEphoneAppAppParamsObject `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ManageEphoneAppOperateObject) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppOperateObject) GoString() string {
  return s.String()
}

func (s *ManageEphoneAppOperateObject) SetId(v string) *ManageEphoneAppOperateObject {
  s.Id = &v
  return s
}

func (s *ManageEphoneAppOperateObject) SetOp(v string) *ManageEphoneAppOperateObject {
  s.Op = &v
  return s
}

func (s *ManageEphoneAppOperateObject) SetParams(v *ManageEphoneAppAppParamsObject) *ManageEphoneAppOperateObject {
  s.Params = v
  return s
}

type ManageEphoneAppAppParamsObject struct {
  // {"en":"apk url", "zh_CN":"apk文件下载地址"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en":"papckage name", "zh_CN":"app包名"}
  PkgName *string `json:"pkgName,omitempty" xml:"pkgName,omitempty"`
}

func (s ManageEphoneAppAppParamsObject) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppAppParamsObject) GoString() string {
  return s.String()
}

func (s *ManageEphoneAppAppParamsObject) SetUrl(v string) *ManageEphoneAppAppParamsObject {
  s.Url = &v
  return s
}

func (s *ManageEphoneAppAppParamsObject) SetPkgName(v string) *ManageEphoneAppAppParamsObject {
  s.PkgName = &v
  return s
}

type ManageEphoneAppResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Tasks []*ManageEphoneAppTask `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"创建消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s ManageEphoneAppResponse) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppResponse) GoString() string {
  return s.String()
}

func (s *ManageEphoneAppResponse) SetTasks(v []*ManageEphoneAppTask) *ManageEphoneAppResponse {
  s.Tasks = v
  return s
}

func (s *ManageEphoneAppResponse) SetStatus(v int) *ManageEphoneAppResponse {
  s.Status = &v
  return s
}

func (s *ManageEphoneAppResponse) SetResult(v string) *ManageEphoneAppResponse {
  s.Result = &v
  return s
}

type ManageEphoneAppTask struct {
  // {"en":"task id", "zh_CN":"任务单id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"任务单消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ManageEphoneAppTask) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppTask) GoString() string {
  return s.String()
}

func (s *ManageEphoneAppTask) SetId(v string) *ManageEphoneAppTask {
  s.Id = &v
  return s
}

func (s *ManageEphoneAppTask) SetMessage(v string) *ManageEphoneAppTask {
  s.Message = &v
  return s
}

type ManageEphoneAppPaths struct {
}

func (s ManageEphoneAppPaths) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppPaths) GoString() string {
  return s.String()
}

type ManageEphoneAppParameters struct {
}

func (s ManageEphoneAppParameters) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppParameters) GoString() string {
  return s.String()
}

type ManageEphoneAppRequestHeader struct {
}

func (s ManageEphoneAppRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppRequestHeader) GoString() string {
  return s.String()
}

type ManageEphoneAppResponseHeader struct {
}

func (s ManageEphoneAppResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageEphoneAppResponseHeader) GoString() string {
  return s.String()
}




type DeleteTunnelApplicationsRequest struct {
  // {"en":"application name", "zh_CN":"应用名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeleteTunnelApplicationsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteTunnelApplicationsRequest) GoString() string {
  return s.String()
}

func (s *DeleteTunnelApplicationsRequest) SetName(v string) *DeleteTunnelApplicationsRequest {
  s.Name = &v
  return s
}

type DeleteTunnelApplicationsResponse struct {
}

func (s DeleteTunnelApplicationsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteTunnelApplicationsResponse) GoString() string {
  return s.String()
}

type DeleteTunnelApplicationsPaths struct {
}

func (s DeleteTunnelApplicationsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteTunnelApplicationsPaths) GoString() string {
  return s.String()
}

type DeleteTunnelApplicationsParameters struct {
}

func (s DeleteTunnelApplicationsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteTunnelApplicationsParameters) GoString() string {
  return s.String()
}

type DeleteTunnelApplicationsRequestHeader struct {
}

func (s DeleteTunnelApplicationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTunnelApplicationsRequestHeader) GoString() string {
  return s.String()
}

type DeleteTunnelApplicationsResponseHeader struct {
}

func (s DeleteTunnelApplicationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTunnelApplicationsResponseHeader) GoString() string {
  return s.String()
}




