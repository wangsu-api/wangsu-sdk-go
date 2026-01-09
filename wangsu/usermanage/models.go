package usermanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type AddAccountIdentRequest struct {
  // {"en":"login name","zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
}

func (s AddAccountIdentRequest) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentRequest) GoString() string {
  return s.String()
}

func (s *AddAccountIdentRequest) SetLoginName(v string) *AddAccountIdentRequest {
  s.LoginName = &v
  return s
}

type AddAccountIdentRequestHeader struct {
}

func (s AddAccountIdentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentRequestHeader) GoString() string {
  return s.String()
}

type AddAccountIdentPaths struct {
}

func (s AddAccountIdentPaths) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentPaths) GoString() string {
  return s.String()
}

type AddAccountIdentParameters struct {
}

func (s AddAccountIdentParameters) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentParameters) GoString() string {
  return s.String()
}

type AddAccountIdentResponse struct {
  // {"en":"response code","zh_CN":"响应编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *AddAccountIdentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddAccountIdentResponse) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentResponse) GoString() string {
  return s.String()
}

func (s *AddAccountIdentResponse) SetCode(v string) *AddAccountIdentResponse {
  s.Code = &v
  return s
}

func (s *AddAccountIdentResponse) SetMessage(v string) *AddAccountIdentResponse {
  s.Message = &v
  return s
}

func (s *AddAccountIdentResponse) SetData(v *AddAccountIdentResponseData) *AddAccountIdentResponse {
  s.Data = v
  return s
}

type AddAccountIdentResponseData struct {
  // {"en":"account accessKey","zh_CN":"账号的accessKey"}
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
  // {"en":"account secretKey","zh_CN":"账号的secretKey"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty" require:"true"`
}

func (s AddAccountIdentResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentResponseData) GoString() string {
  return s.String()
}

func (s *AddAccountIdentResponseData) SetAccessKey(v string) *AddAccountIdentResponseData {
  s.AccessKey = &v
  return s
}

func (s *AddAccountIdentResponseData) SetSecretKey(v string) *AddAccountIdentResponseData {
  s.SecretKey = &v
  return s
}

type AddAccountIdentResponseHeader struct {
}

func (s AddAccountIdentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddAccountIdentResponseHeader) GoString() string {
  return s.String()
}




type DeleteSubAccountRequest struct {
}

func (s DeleteSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountRequest) GoString() string {
  return s.String()
}

type DeleteSubAccountRequestHeader struct {
}

func (s DeleteSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountRequestHeader) GoString() string {
  return s.String()
}

type DeleteSubAccountPaths struct {
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s DeleteSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountPaths) GoString() string {
  return s.String()
}

func (s *DeleteSubAccountPaths) SetLoginName(v string) *DeleteSubAccountPaths {
  s.LoginName = &v
  return s
}

type DeleteSubAccountParameters struct {
}

func (s DeleteSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountParameters) GoString() string {
  return s.String()
}

type DeleteSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountResponse) GoString() string {
  return s.String()
}

func (s *DeleteSubAccountResponse) SetCode(v string) *DeleteSubAccountResponse {
  s.Code = &v
  return s
}

func (s *DeleteSubAccountResponse) SetMessage(v string) *DeleteSubAccountResponse {
  s.Message = &v
  return s
}

type DeleteSubAccountResponseHeader struct {
}

func (s DeleteSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountResponseHeader) GoString() string {
  return s.String()
}




type CheckLoginNameLegalRequest struct {
  // {"en":"login name", "zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s CheckLoginNameLegalRequest) String() string {
  return tea.Prettify(s)
}

func (s CheckLoginNameLegalRequest) GoString() string {
  return s.String()
}

func (s *CheckLoginNameLegalRequest) SetLoginName(v string) *CheckLoginNameLegalRequest {
  s.LoginName = &v
  return s
}

type CheckLoginNameLegalResponse struct {
  // {"en":"Message", "zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Status Code", "zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CheckLoginNameLegalResponse) String() string {
  return tea.Prettify(s)
}

func (s CheckLoginNameLegalResponse) GoString() string {
  return s.String()
}

func (s *CheckLoginNameLegalResponse) SetMessage(v string) *CheckLoginNameLegalResponse {
  s.Message = &v
  return s
}

func (s *CheckLoginNameLegalResponse) SetCode(v string) *CheckLoginNameLegalResponse {
  s.Code = &v
  return s
}

type CheckLoginNameLegalPaths struct {
}

func (s CheckLoginNameLegalPaths) String() string {
  return tea.Prettify(s)
}

func (s CheckLoginNameLegalPaths) GoString() string {
  return s.String()
}

type CheckLoginNameLegalParameters struct {
}

func (s CheckLoginNameLegalParameters) String() string {
  return tea.Prettify(s)
}

func (s CheckLoginNameLegalParameters) GoString() string {
  return s.String()
}

type CheckLoginNameLegalRequestHeader struct {
}

func (s CheckLoginNameLegalRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckLoginNameLegalRequestHeader) GoString() string {
  return s.String()
}

type CheckLoginNameLegalResponseHeader struct {
}

func (s CheckLoginNameLegalResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckLoginNameLegalResponseHeader) GoString() string {
  return s.String()
}




type QuerySubAccountInfoRequest struct {
}

func (s QuerySubAccountInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoRequest) GoString() string {
  return s.String()
}

type QuerySubAccountInfoRequestHeader struct {
}

func (s QuerySubAccountInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoRequestHeader) GoString() string {
  return s.String()
}

type QuerySubAccountInfoPaths struct {
  // {"en":"Login Name","zh_CN":"子账号登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s QuerySubAccountInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoPaths) GoString() string {
  return s.String()
}

func (s *QuerySubAccountInfoPaths) SetLoginName(v string) *QuerySubAccountInfoPaths {
  s.LoginName = &v
  return s
}

type QuerySubAccountInfoParameters struct {
}

func (s QuerySubAccountInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoParameters) GoString() string {
  return s.String()
}

type QuerySubAccountInfoResponse struct {
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回数据"}
  Data *QuerySubAccountInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QuerySubAccountInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoResponse) GoString() string {
  return s.String()
}

func (s *QuerySubAccountInfoResponse) SetCode(v string) *QuerySubAccountInfoResponse {
  s.Code = &v
  return s
}

func (s *QuerySubAccountInfoResponse) SetMessage(v string) *QuerySubAccountInfoResponse {
  s.Message = &v
  return s
}

func (s *QuerySubAccountInfoResponse) SetData(v *QuerySubAccountInfoResponseData) *QuerySubAccountInfoResponse {
  s.Data = v
  return s
}

type QuerySubAccountInfoResponseData struct {
  // {"en":"login name","zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"display name","zh_CN":"称呼"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
  // {"en":"mobile","zh_CN":"手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty" require:"true"`
  // {"en":"create time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"consoleEnable","zh_CN":"是否允许登录控制台：1是 0 否"}
  ConsoleEnable *int `json:"consoleEnable,omitempty" xml:"consoleEnable,omitempty" require:"true"`
  // {"en":"status","zh_CN":"状态： 1 启用 0 停用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s QuerySubAccountInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoResponseData) GoString() string {
  return s.String()
}

func (s *QuerySubAccountInfoResponseData) SetLoginName(v string) *QuerySubAccountInfoResponseData {
  s.LoginName = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetDisplayName(v string) *QuerySubAccountInfoResponseData {
  s.DisplayName = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetEmail(v string) *QuerySubAccountInfoResponseData {
  s.Email = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetMobile(v string) *QuerySubAccountInfoResponseData {
  s.Mobile = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetCreateTime(v string) *QuerySubAccountInfoResponseData {
  s.CreateTime = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetConsoleEnable(v int) *QuerySubAccountInfoResponseData {
  s.ConsoleEnable = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetStatus(v int) *QuerySubAccountInfoResponseData {
  s.Status = &v
  return s
}

type QuerySubAccountInfoResponseHeader struct {
}

func (s QuerySubAccountInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoResponseHeader) GoString() string {
  return s.String()
}




type BatchAddOrRevokePolicyToSubAccountRequest struct {
  // {"en":"Specify policy ID","zh_CN":"指定权限策略id"}
  PolicyId []*int64 `json:"policyId,omitempty" xml:"policyId,omitempty" type:"Repeated"`
  // {"en":"Policy name","zh_CN":"策略名称"}
  PolicyName []*string `json:"policyName,omitempty" xml:"policyName,omitempty" type:"Repeated"`
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"Select you want to add or revoke policy for sub account.\n\n0:add policy\n\n1:revoke policy","zh_CN":"选择需要为子用户添加或撤销权限策略\n\n0：添加权限\n\n1：撤销权限"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s BatchAddOrRevokePolicyToSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountRequest) GoString() string {
  return s.String()
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetPolicyId(v []*int64) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.PolicyId = v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetPolicyName(v []*string) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.PolicyName = v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetLoginName(v string) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetType(v int) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.Type = &v
  return s
}

type BatchAddOrRevokePolicyToSubAccountRequestHeader struct {
}

func (s BatchAddOrRevokePolicyToSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountRequestHeader) GoString() string {
  return s.String()
}

type BatchAddOrRevokePolicyToSubAccountPaths struct {
}

func (s BatchAddOrRevokePolicyToSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountPaths) GoString() string {
  return s.String()
}

type BatchAddOrRevokePolicyToSubAccountParameters struct {
}

func (s BatchAddOrRevokePolicyToSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountParameters) GoString() string {
  return s.String()
}

type BatchAddOrRevokePolicyToSubAccountResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s BatchAddOrRevokePolicyToSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountResponse) GoString() string {
  return s.String()
}

func (s *BatchAddOrRevokePolicyToSubAccountResponse) SetCode(v string) *BatchAddOrRevokePolicyToSubAccountResponse {
  s.Code = &v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountResponse) SetMsg(v string) *BatchAddOrRevokePolicyToSubAccountResponse {
  s.Msg = &v
  return s
}

type BatchAddOrRevokePolicyToSubAccountResponseHeader struct {
}

func (s BatchAddOrRevokePolicyToSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountResponseHeader) GoString() string {
  return s.String()
}




type UpdateAccountIdentRequest struct {
  // {"en":"accessKey", "zh_CN":"accessKey"}
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
  // {"en":"status", "zh_CN":"状态 disabled表示禁用，activated表示启用"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateAccountIdentRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAccountIdentRequest) GoString() string {
  return s.String()
}

func (s *UpdateAccountIdentRequest) SetAccessKey(v string) *UpdateAccountIdentRequest {
  s.AccessKey = &v
  return s
}

func (s *UpdateAccountIdentRequest) SetStatus(v string) *UpdateAccountIdentRequest {
  s.Status = &v
  return s
}

type UpdateAccountIdentResponse struct {
  // {"en":"code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateAccountIdentResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAccountIdentResponse) GoString() string {
  return s.String()
}

func (s *UpdateAccountIdentResponse) SetCode(v string) *UpdateAccountIdentResponse {
  s.Code = &v
  return s
}

func (s *UpdateAccountIdentResponse) SetMessage(v string) *UpdateAccountIdentResponse {
  s.Message = &v
  return s
}

type UpdateAccountIdentPaths struct {
}

func (s UpdateAccountIdentPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAccountIdentPaths) GoString() string {
  return s.String()
}

type UpdateAccountIdentParameters struct {
}

func (s UpdateAccountIdentParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAccountIdentParameters) GoString() string {
  return s.String()
}

type UpdateAccountIdentRequestHeader struct {
}

func (s UpdateAccountIdentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAccountIdentRequestHeader) GoString() string {
  return s.String()
}

type UpdateAccountIdentResponseHeader struct {
}

func (s UpdateAccountIdentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAccountIdentResponseHeader) GoString() string {
  return s.String()
}




type QueryAgentAssociatedMainAccountServiceRequest struct {
}

func (s QueryAgentAssociatedMainAccountServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServiceRequest) GoString() string {
  return s.String()
}

type QueryAgentAssociatedMainAccountServiceRequestHeader struct {
}

func (s QueryAgentAssociatedMainAccountServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryAgentAssociatedMainAccountServicePaths struct {
}

func (s QueryAgentAssociatedMainAccountServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServicePaths) GoString() string {
  return s.String()
}

type QueryAgentAssociatedMainAccountServiceParameters struct {
}

func (s QueryAgentAssociatedMainAccountServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServiceParameters) GoString() string {
  return s.String()
}

type QueryAgentAssociatedMainAccountServiceResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryAgentAssociatedMainAccountServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryAgentAssociatedMainAccountServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryAgentAssociatedMainAccountServiceResponse) SetCode(v string) *QueryAgentAssociatedMainAccountServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryAgentAssociatedMainAccountServiceResponse) SetMsg(v string) *QueryAgentAssociatedMainAccountServiceResponse {
  s.Msg = &v
  return s
}

func (s *QueryAgentAssociatedMainAccountServiceResponse) SetData(v *QueryAgentAssociatedMainAccountServiceResponseData) *QueryAgentAssociatedMainAccountServiceResponse {
  s.Data = v
  return s
}

type QueryAgentAssociatedMainAccountServiceResponseData struct {
  // {"en":"Service type of the Acceleration domain:1. If not specified, it is considered as 'no restriction on service type.'2. For multiple Application server types, please separate them with an English semicolon \";\"","zh_CN":"加速域名的服务类型：1.未传递视为不限服务类型2.多个服务类型请使用英文分号\";\"分隔"}
  MainAccount *string `json:"mainAccount,omitempty" xml:"mainAccount,omitempty" require:"true"`
  // {"en":"Primary account display name","zh_CN":"主账号显示名"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"The parent account of the primary account","zh_CN":"主账号对应的父主账号登录名"}
  ParentMainAccount *string `json:"parentMainAccount,omitempty" xml:"parentMainAccount,omitempty" require:"true"`
}

func (s QueryAgentAssociatedMainAccountServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServiceResponseData) GoString() string {
  return s.String()
}

func (s *QueryAgentAssociatedMainAccountServiceResponseData) SetMainAccount(v string) *QueryAgentAssociatedMainAccountServiceResponseData {
  s.MainAccount = &v
  return s
}

func (s *QueryAgentAssociatedMainAccountServiceResponseData) SetDisplayName(v string) *QueryAgentAssociatedMainAccountServiceResponseData {
  s.DisplayName = &v
  return s
}

func (s *QueryAgentAssociatedMainAccountServiceResponseData) SetParentMainAccount(v string) *QueryAgentAssociatedMainAccountServiceResponseData {
  s.ParentMainAccount = &v
  return s
}

type QueryAgentAssociatedMainAccountServiceResponseHeader struct {
}

func (s QueryAgentAssociatedMainAccountServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAgentAssociatedMainAccountServiceResponseHeader) GoString() string {
  return s.String()
}




type DeleteAccountIdentRequest struct {
}

func (s DeleteAccountIdentRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAccountIdentRequest) GoString() string {
  return s.String()
}

type DeleteAccountIdentResponse struct {
  // {"en":"Status Code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteAccountIdentResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAccountIdentResponse) GoString() string {
  return s.String()
}

func (s *DeleteAccountIdentResponse) SetCode(v string) *DeleteAccountIdentResponse {
  s.Code = &v
  return s
}

func (s *DeleteAccountIdentResponse) SetMessage(v string) *DeleteAccountIdentResponse {
  s.Message = &v
  return s
}

type DeleteAccountIdentPaths struct {
  // {"en":"accessKey", "zh_CN":"accessKey"}
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
}

func (s DeleteAccountIdentPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAccountIdentPaths) GoString() string {
  return s.String()
}

func (s *DeleteAccountIdentPaths) SetAccessKey(v string) *DeleteAccountIdentPaths {
  s.AccessKey = &v
  return s
}

type DeleteAccountIdentParameters struct {
}

func (s DeleteAccountIdentParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAccountIdentParameters) GoString() string {
  return s.String()
}

type DeleteAccountIdentRequestHeader struct {
}

func (s DeleteAccountIdentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAccountIdentRequestHeader) GoString() string {
  return s.String()
}

type DeleteAccountIdentResponseHeader struct {
}

func (s DeleteAccountIdentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAccountIdentResponseHeader) GoString() string {
  return s.String()
}




type AddSubAccountRequest struct {
  // {"en":"password","zh_CN":"密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty" require:"true"`
  // {"en":"parentLoginName","zh_CN":"父账号登录名"}
  ParentLoginName *string `json:"parentLoginName,omitempty" xml:"parentLoginName,omitempty" require:"true"`
  // {"en":"apiKey","zh_CN":"apiKey"}
  ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
  // {"en":"phone","zh_CN":"电话"}
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  // {"en":"consoleEnable","zh_CN":"是否允许登录控制台：1是 0 否"}
  ConsoleEnable *int `json:"consoleEnable,omitempty" xml:"consoleEnable,omitempty"`
  // {"en":"display name","zh_CN":"称呼"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"login name","zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"mobile","zh_CN":"手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
  // {"en":"programmaticEnable","zh_CN":"是否允许编程访问：1是 0 否"}
  ProgrammaticEnable *int `json:"programmaticEnable,omitempty" xml:"programmaticEnable,omitempty"`
  // {"en":"loginResetPassword","zh_CN":"登录是否需重置密码：1是 0 否"}
  LoginResetPassword *int `json:"loginResetPassword,omitempty" xml:"loginResetPassword,omitempty"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en":"status","zh_CN":"状态： 1 启用 0 停用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s AddSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountRequest) GoString() string {
  return s.String()
}

func (s *AddSubAccountRequest) SetPassword(v string) *AddSubAccountRequest {
  s.Password = &v
  return s
}

func (s *AddSubAccountRequest) SetParentLoginName(v string) *AddSubAccountRequest {
  s.ParentLoginName = &v
  return s
}

func (s *AddSubAccountRequest) SetApiKey(v string) *AddSubAccountRequest {
  s.ApiKey = &v
  return s
}

func (s *AddSubAccountRequest) SetPhone(v string) *AddSubAccountRequest {
  s.Phone = &v
  return s
}

func (s *AddSubAccountRequest) SetConsoleEnable(v int) *AddSubAccountRequest {
  s.ConsoleEnable = &v
  return s
}

func (s *AddSubAccountRequest) SetDisplayName(v string) *AddSubAccountRequest {
  s.DisplayName = &v
  return s
}

func (s *AddSubAccountRequest) SetLoginName(v string) *AddSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *AddSubAccountRequest) SetMobile(v string) *AddSubAccountRequest {
  s.Mobile = &v
  return s
}

func (s *AddSubAccountRequest) SetProgrammaticEnable(v int) *AddSubAccountRequest {
  s.ProgrammaticEnable = &v
  return s
}

func (s *AddSubAccountRequest) SetLoginResetPassword(v int) *AddSubAccountRequest {
  s.LoginResetPassword = &v
  return s
}

func (s *AddSubAccountRequest) SetEmail(v string) *AddSubAccountRequest {
  s.Email = &v
  return s
}

func (s *AddSubAccountRequest) SetStatus(v int) *AddSubAccountRequest {
  s.Status = &v
  return s
}

type AddSubAccountRequestHeader struct {
}

func (s AddSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountRequestHeader) GoString() string {
  return s.String()
}

type AddSubAccountPaths struct {
}

func (s AddSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountPaths) GoString() string {
  return s.String()
}

type AddSubAccountParameters struct {
}

func (s AddSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountParameters) GoString() string {
  return s.String()
}

type AddSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountResponse) GoString() string {
  return s.String()
}

func (s *AddSubAccountResponse) SetCode(v string) *AddSubAccountResponse {
  s.Code = &v
  return s
}

func (s *AddSubAccountResponse) SetMessage(v string) *AddSubAccountResponse {
  s.Message = &v
  return s
}

type AddSubAccountResponseHeader struct {
}

func (s AddSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountResponseHeader) GoString() string {
  return s.String()
}




type GetSubAccountListRequest struct {
  // {"en":"Get a list of sub accounts under the main account","zh_CN":"主用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"Sub account loginName or displayName fuzzy query","zh_CN":"子账号loginName或displayName模糊查询"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Page Number of Current Page.If it was empty, it would be treated as not being divided into pages. The contents filled in the pageSize field would not be effective","zh_CN":"指定分页查询时,当前页的页码。为空则不分页处理全部返回,pageSize字段填写内容不生效"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"The maximum number of data displayed on each page.\nThe maximum value of the PageSize is 100. The number of data bar displayed on each page was 20 by default. When the PageSize value is empty, and pageIndex is not empty,20 data would be returned by default.","zh_CN":"指定分页查询时,每页显示的数据最大条数。\nPageSize参数最大取值为100。每页默认显示的数据条数为20条,PageSize参数值为空时,将默认返回20条数据。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetSubAccountListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListRequest) GoString() string {
  return s.String()
}

func (s *GetSubAccountListRequest) SetLoginName(v string) *GetSubAccountListRequest {
  s.LoginName = &v
  return s
}

func (s *GetSubAccountListRequest) SetName(v string) *GetSubAccountListRequest {
  s.Name = &v
  return s
}

func (s *GetSubAccountListRequest) SetPageIndex(v int) *GetSubAccountListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetSubAccountListRequest) SetPageSize(v int) *GetSubAccountListRequest {
  s.PageSize = &v
  return s
}

type GetSubAccountListRequestHeader struct {
}

func (s GetSubAccountListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListRequestHeader) GoString() string {
  return s.String()
}

type GetSubAccountListPaths struct {
}

func (s GetSubAccountListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListPaths) GoString() string {
  return s.String()
}

type GetSubAccountListParameters struct {
}

func (s GetSubAccountListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListParameters) GoString() string {
  return s.String()
}

type GetSubAccountListResponse struct {
  // {"en":"code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回值"}
  Data *GetSubAccountListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetSubAccountListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponse) GoString() string {
  return s.String()
}

func (s *GetSubAccountListResponse) SetCode(v string) *GetSubAccountListResponse {
  s.Code = &v
  return s
}

func (s *GetSubAccountListResponse) SetMessage(v string) *GetSubAccountListResponse {
  s.Message = &v
  return s
}

func (s *GetSubAccountListResponse) SetData(v *GetSubAccountListResponseData) *GetSubAccountListResponse {
  s.Data = v
  return s
}

type GetSubAccountListResponseData struct {
  // {"en":"page Index","zh_CN":"页码"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"page Size","zh_CN":"每页个数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"total","zh_CN":"总条数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"rows","zh_CN":"每页数据"}
  Rows []*GetSubAccountListResponseDataRows `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
}

func (s GetSubAccountListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponseData) GoString() string {
  return s.String()
}

func (s *GetSubAccountListResponseData) SetPageIndex(v int) *GetSubAccountListResponseData {
  s.PageIndex = &v
  return s
}

func (s *GetSubAccountListResponseData) SetPageSize(v int) *GetSubAccountListResponseData {
  s.PageSize = &v
  return s
}

func (s *GetSubAccountListResponseData) SetTotal(v int64) *GetSubAccountListResponseData {
  s.Total = &v
  return s
}

func (s *GetSubAccountListResponseData) SetRows(v []*GetSubAccountListResponseDataRows) *GetSubAccountListResponseData {
  s.Rows = v
  return s
}

type GetSubAccountListResponseDataRows struct     {
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"Sub account display name","zh_CN":"子用户显示名"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"Sub accout's E-mail","zh_CN":"绑定的邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
  // {"en":"Sub accout's mobile phone","zh_CN":"绑定的手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty" require:"true"`
  // {"en":"CreateTime","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"The status of accout1:Activate0:Disable","zh_CN":"账号启用/禁用状态,1代表启用,0代表禁用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s GetSubAccountListResponseDataRows) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponseDataRows) GoString() string {
  return s.String()
}

func (s *GetSubAccountListResponseDataRows) SetLoginName(v string) *GetSubAccountListResponseDataRows {
  s.LoginName = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetDisplayName(v string) *GetSubAccountListResponseDataRows {
  s.DisplayName = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetEmail(v string) *GetSubAccountListResponseDataRows {
  s.Email = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetMobile(v string) *GetSubAccountListResponseDataRows {
  s.Mobile = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetCreateTime(v string) *GetSubAccountListResponseDataRows {
  s.CreateTime = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetStatus(v int) *GetSubAccountListResponseDataRows {
  s.Status = &v
  return s
}

type GetSubAccountListResponseHeader struct {
}

func (s GetSubAccountListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponseHeader) GoString() string {
  return s.String()
}




type QueryPolicyAttachedMainAccountOrSubAccountRequest struct {
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequest) GoString() string {
  return s.String()
}

type QueryPolicyAttachedMainAccountOrSubAccountRequestHeader struct {
  // {"en":"Select the specified language and return the policy description of the corresponding language. Optional values:zh_CN，en，ko_KR，ja_JP；Default language is en","zh_CN":"选择指定语言返回对应语言的策略描述，可选值：zh_CN，en，ko_KR，ja_JP；未选择默认en"}
  AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountRequestHeader) SetAcceptLanguage(v string) *QueryPolicyAttachedMainAccountOrSubAccountRequestHeader {
  s.AcceptLanguage = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountPaths struct {
  // {"en":"loginName(Main or ordinary subAccount are available)","zh_CN":"用户登录名（可传主子用户）"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountPaths) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountPaths) SetLoginName(v string) *QueryPolicyAttachedMainAccountOrSubAccountPaths {
  s.LoginName = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountParameters struct {
}

func (s QueryPolicyAttachedMainAccountOrSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountParameters) GoString() string {
  return s.String()
}

type QueryPolicyAttachedMainAccountOrSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*QueryPolicyAttachedMainAccountOrSubAccountResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponse) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetCode(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.Code = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetMsg(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.Msg = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetData(v []*QueryPolicyAttachedMainAccountOrSubAccountResponseData) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.Data = v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetRequestId(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.RequestId = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountResponseData struct     {
  // {"en":"policyId","zh_CN":"策略id"}
  PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"policy name","zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Policy description","zh_CN":"策略描述内容"}
  PolicyDescribe *string `json:"policyDescribe,omitempty" xml:"policyDescribe,omitempty" require:"true"`
  // {"en":"policy type","zh_CN":"策略类型：system：系统策略、custom自定义策略"}
  PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseData) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyId(v int64) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyId = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyName(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyName = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyDescribe(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyDescribe = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyType(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyType = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountResponseHeader struct {
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseHeader) GoString() string {
  return s.String()
}




type ListAccountIdentRequest struct {
  // {"en":"loginName", "zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
}

func (s ListAccountIdentRequest) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentRequest) GoString() string {
  return s.String()
}

func (s *ListAccountIdentRequest) SetLoginName(v string) *ListAccountIdentRequest {
  s.LoginName = &v
  return s
}

type ListAccountIdentResponse struct {
  // {"en":"code", "zh_CN":"codec"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"响应数据"}
  Data []*ListAccountIdentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListAccountIdentResponse) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentResponse) GoString() string {
  return s.String()
}

func (s *ListAccountIdentResponse) SetCode(v string) *ListAccountIdentResponse {
  s.Code = &v
  return s
}

func (s *ListAccountIdentResponse) SetMessage(v string) *ListAccountIdentResponse {
  s.Message = &v
  return s
}

func (s *ListAccountIdentResponse) SetData(v []*ListAccountIdentResponseData) *ListAccountIdentResponse {
  s.Data = v
  return s
}

type ListAccountIdentResponseData struct     {
  // {"en":"accessKey", "zh_CN":"accessKey"}
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
  // {"en":"secretKey", "zh_CN":"secretKey"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"状态 disabled表示禁用，activated表示启用"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s ListAccountIdentResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentResponseData) GoString() string {
  return s.String()
}

func (s *ListAccountIdentResponseData) SetAccessKey(v string) *ListAccountIdentResponseData {
  s.AccessKey = &v
  return s
}

func (s *ListAccountIdentResponseData) SetSecretKey(v string) *ListAccountIdentResponseData {
  s.SecretKey = &v
  return s
}

func (s *ListAccountIdentResponseData) SetStatus(v string) *ListAccountIdentResponseData {
  s.Status = &v
  return s
}

type ListAccountIdentPaths struct {
}

func (s ListAccountIdentPaths) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentPaths) GoString() string {
  return s.String()
}

type ListAccountIdentParameters struct {
}

func (s ListAccountIdentParameters) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentParameters) GoString() string {
  return s.String()
}

type ListAccountIdentRequestHeader struct {
}

func (s ListAccountIdentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentRequestHeader) GoString() string {
  return s.String()
}

type ListAccountIdentResponseHeader struct {
}

func (s ListAccountIdentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAccountIdentResponseHeader) GoString() string {
  return s.String()
}




type UpdateSubAccountRequest struct {
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"display name","zh_CN":"子用户显示名称"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en":"mobile","zh_CN":"手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
  // {"en":"Programmatic Access or not 1:Yes0:No","zh_CN":"子用户是否允许登录控制台：1是 0 否"}
  ConsoleEnable *int `json:"consoleEnable,omitempty" xml:"consoleEnable,omitempty"`
  // {"en":"openApiStatus","zh_CN":"是否开启OpenAPI 0否，1是"}
  OpenApiStatus *int `json:"openApiStatus,omitempty" xml:"openApiStatus,omitempty"`
  // {"en":"area code","zh_CN":"手机号区号"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
  // {"en":"loginResetPassword","zh_CN":"登录是否重置密码 1 是 0  否"}
  LoginResetPassword *int `json:"loginResetPassword,omitempty" xml:"loginResetPassword,omitempty"`
  // {"en":"password","zh_CN":"密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"Account status: 0: Disabled, 1: Enabled","zh_CN":"账号状态：0：禁用，1：启用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountRequest) GoString() string {
  return s.String()
}

func (s *UpdateSubAccountRequest) SetLoginName(v string) *UpdateSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *UpdateSubAccountRequest) SetDisplayName(v string) *UpdateSubAccountRequest {
  s.DisplayName = &v
  return s
}

func (s *UpdateSubAccountRequest) SetEmail(v string) *UpdateSubAccountRequest {
  s.Email = &v
  return s
}

func (s *UpdateSubAccountRequest) SetMobile(v string) *UpdateSubAccountRequest {
  s.Mobile = &v
  return s
}

func (s *UpdateSubAccountRequest) SetConsoleEnable(v int) *UpdateSubAccountRequest {
  s.ConsoleEnable = &v
  return s
}

func (s *UpdateSubAccountRequest) SetOpenApiStatus(v int) *UpdateSubAccountRequest {
  s.OpenApiStatus = &v
  return s
}

func (s *UpdateSubAccountRequest) SetAreaCode(v string) *UpdateSubAccountRequest {
  s.AreaCode = &v
  return s
}

func (s *UpdateSubAccountRequest) SetLoginResetPassword(v int) *UpdateSubAccountRequest {
  s.LoginResetPassword = &v
  return s
}

func (s *UpdateSubAccountRequest) SetPassword(v string) *UpdateSubAccountRequest {
  s.Password = &v
  return s
}

func (s *UpdateSubAccountRequest) SetStatus(v int) *UpdateSubAccountRequest {
  s.Status = &v
  return s
}

type UpdateSubAccountRequestHeader struct {
}

func (s UpdateSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountRequestHeader) GoString() string {
  return s.String()
}

type UpdateSubAccountPaths struct {
}

func (s UpdateSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountPaths) GoString() string {
  return s.String()
}

type UpdateSubAccountParameters struct {
}

func (s UpdateSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountParameters) GoString() string {
  return s.String()
}

type UpdateSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountResponse) GoString() string {
  return s.String()
}

func (s *UpdateSubAccountResponse) SetCode(v string) *UpdateSubAccountResponse {
  s.Code = &v
  return s
}

func (s *UpdateSubAccountResponse) SetMessage(v string) *UpdateSubAccountResponse {
  s.Message = &v
  return s
}

type UpdateSubAccountResponseHeader struct {
}

func (s UpdateSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountResponseHeader) GoString() string {
  return s.String()
}




