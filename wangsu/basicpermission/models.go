package basicpermission

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteTerminalAuthRequest struct {
  // {"en":"terminal auth name ", "zh_CN":"要删除的基础权限名称"}
  TerminalAuthName []*int64 `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" type:"Repeated"`
}

func (s DeleteTerminalAuthRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthRequest) GoString() string {
  return s.String()
}

func (s *DeleteTerminalAuthRequest) SetTerminalAuthName(v []*int64) *DeleteTerminalAuthRequest {
  s.TerminalAuthName = v
  return s
}

type DeleteTerminalAuthResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *DeleteTerminalAuthContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s DeleteTerminalAuthResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthResponse) GoString() string {
  return s.String()
}

func (s *DeleteTerminalAuthResponse) SetReturnCode(v string) *DeleteTerminalAuthResponse {
  s.ReturnCode = &v
  return s
}

func (s *DeleteTerminalAuthResponse) SetReturnMsg(v string) *DeleteTerminalAuthResponse {
  s.ReturnMsg = &v
  return s
}

func (s *DeleteTerminalAuthResponse) SetContent(v *DeleteTerminalAuthContentEntity) *DeleteTerminalAuthResponse {
  s.Content = v
  return s
}

type DeleteTerminalAuthContentEntity struct {
  // {"en":"terminal auth name ", "zh_CN":"要删除的基础权限名称"}
  TerminalAuthName []*int64 `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" type:"Repeated"`
}

func (s DeleteTerminalAuthContentEntity) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthContentEntity) GoString() string {
  return s.String()
}

func (s *DeleteTerminalAuthContentEntity) SetTerminalAuthName(v []*int64) *DeleteTerminalAuthContentEntity {
  s.TerminalAuthName = v
  return s
}

type DeleteTerminalAuthPaths struct {
}

func (s DeleteTerminalAuthPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthPaths) GoString() string {
  return s.String()
}

type DeleteTerminalAuthParameters struct {
}

func (s DeleteTerminalAuthParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthParameters) GoString() string {
  return s.String()
}

type DeleteTerminalAuthRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s DeleteTerminalAuthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteTerminalAuthRequestHeader) SetAuthUser(v string) *DeleteTerminalAuthRequestHeader {
  s.AuthUser = &v
  return s
}

type DeleteTerminalAuthResponseHeader struct {
}

func (s DeleteTerminalAuthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteTerminalAuthResponseHeader) GoString() string {
  return s.String()
}




type AssociateRightsGroupsToUserOrUserGroupRequest struct {
  // {"en":"List of base permission names","zh_CN":"基础权限名称列表"}
  TerminalAuthNames []*string `json:"terminalAuthNames,omitempty" xml:"terminalAuthNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of authorized users","zh_CN":"授权的用户列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"List of authorized user group IDs","zh_CN":"授权的用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
  // {"en":"Action Type, 0: Append, 1: Overwrite","zh_CN":"操作类型，0：追加，1：覆盖"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
}

func (s AssociateRightsGroupsToUserOrUserGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateRightsGroupsToUserOrUserGroupRequest) GoString() string {
  return s.String()
}

func (s *AssociateRightsGroupsToUserOrUserGroupRequest) SetTerminalAuthNames(v []*string) *AssociateRightsGroupsToUserOrUserGroupRequest {
  s.TerminalAuthNames = v
  return s
}

func (s *AssociateRightsGroupsToUserOrUserGroupRequest) SetAuthorizedUsers(v []*string) *AssociateRightsGroupsToUserOrUserGroupRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *AssociateRightsGroupsToUserOrUserGroupRequest) SetAuthorizedUserGroupIds(v []*int64) *AssociateRightsGroupsToUserOrUserGroupRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

func (s *AssociateRightsGroupsToUserOrUserGroupRequest) SetActionType(v int) *AssociateRightsGroupsToUserOrUserGroupRequest {
  s.ActionType = &v
  return s
}

type AssociateRightsGroupsToUserOrUserGroupRequestHeader struct {
}

func (s AssociateRightsGroupsToUserOrUserGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateRightsGroupsToUserOrUserGroupRequestHeader) GoString() string {
  return s.String()
}

type AssociateRightsGroupsToUserOrUserGroupPaths struct {
}

func (s AssociateRightsGroupsToUserOrUserGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateRightsGroupsToUserOrUserGroupPaths) GoString() string {
  return s.String()
}

type AssociateRightsGroupsToUserOrUserGroupParameters struct {
}

func (s AssociateRightsGroupsToUserOrUserGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateRightsGroupsToUserOrUserGroupParameters) GoString() string {
  return s.String()
}

type AssociateRightsGroupsToUserOrUserGroupResponse struct {
  // {"en":"List of base permission names","zh_CN":"基础权限名称列表"}
  TerminalAuthNames []*string `json:"terminalAuthNames,omitempty" xml:"terminalAuthNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of authorized users","zh_CN":"授权的用户列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of authorized user group IDs","zh_CN":"授权的用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" require:"true" type:"Repeated"`
  // {"en":"Action Type, 0: Append, 1: Overwrite","zh_CN":"操作类型，0：追加，1：覆盖"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
}

func (s AssociateRightsGroupsToUserOrUserGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateRightsGroupsToUserOrUserGroupResponse) GoString() string {
  return s.String()
}

func (s *AssociateRightsGroupsToUserOrUserGroupResponse) SetTerminalAuthNames(v []*string) *AssociateRightsGroupsToUserOrUserGroupResponse {
  s.TerminalAuthNames = v
  return s
}

func (s *AssociateRightsGroupsToUserOrUserGroupResponse) SetAuthorizedUsers(v []*string) *AssociateRightsGroupsToUserOrUserGroupResponse {
  s.AuthorizedUsers = v
  return s
}

func (s *AssociateRightsGroupsToUserOrUserGroupResponse) SetAuthorizedUserGroupIds(v []*int64) *AssociateRightsGroupsToUserOrUserGroupResponse {
  s.AuthorizedUserGroupIds = v
  return s
}

func (s *AssociateRightsGroupsToUserOrUserGroupResponse) SetActionType(v int) *AssociateRightsGroupsToUserOrUserGroupResponse {
  s.ActionType = &v
  return s
}

type AssociateRightsGroupsToUserOrUserGroupResponseHeader struct {
}

func (s AssociateRightsGroupsToUserOrUserGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateRightsGroupsToUserOrUserGroupResponseHeader) GoString() string {
  return s.String()
}




type UpdateTerminalAuthRequest struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Terminal auth name of the specific auth.", "zh_CN":"权限名称"}
  TerminalAuthName *string `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty"`
  // {"en":"Whether to enable the auth,1: enable 0: disable", "zh_CN":"是否启用该权限，1:启用0:禁用"}
  EnableStatus *int `json:"enableStatus,omitempty" xml:"enableStatus,omitempty"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"Associated application list", "zh_CN":"关联的应用列表"}
  AuthorizedResources []*int64 `json:"authorizedResources,omitempty" xml:"authorizedResources,omitempty" type:"Repeated"`
  // {"en":"Associated user list", "zh_CN":"关联的用户列表"}
  AuthorizedUsers []*int64 `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"关联的用户组列表ID"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
  // {"en":"New terminal auth name of the specific auth.", "zh_CN":"新权限名称"}
  NewTerminalAuthName *string `json:"newTerminalAuthName,omitempty" xml:"newTerminalAuthName,omitempty"`
}

func (s UpdateTerminalAuthRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRequest) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRequest) SetTerminalAuthId(v int64) *UpdateTerminalAuthRequest {
  s.TerminalAuthId = &v
  return s
}

func (s *UpdateTerminalAuthRequest) SetTerminalAuthName(v string) *UpdateTerminalAuthRequest {
  s.TerminalAuthName = &v
  return s
}

func (s *UpdateTerminalAuthRequest) SetEnableStatus(v int) *UpdateTerminalAuthRequest {
  s.EnableStatus = &v
  return s
}

func (s *UpdateTerminalAuthRequest) SetRemark(v string) *UpdateTerminalAuthRequest {
  s.Remark = &v
  return s
}

func (s *UpdateTerminalAuthRequest) SetAuthorizedResources(v []*int64) *UpdateTerminalAuthRequest {
  s.AuthorizedResources = v
  return s
}

func (s *UpdateTerminalAuthRequest) SetAuthorizedUsers(v []*int64) *UpdateTerminalAuthRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *UpdateTerminalAuthRequest) SetAuthorizedUserGroupIds(v []*int64) *UpdateTerminalAuthRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

func (s *UpdateTerminalAuthRequest) SetNewTerminalAuthName(v string) *UpdateTerminalAuthRequest {
  s.NewTerminalAuthName = &v
  return s
}

type UpdateTerminalAuthResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *UpdateTerminalAuthContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s UpdateTerminalAuthResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthResponse) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthResponse) SetReturnCode(v string) *UpdateTerminalAuthResponse {
  s.ReturnCode = &v
  return s
}

func (s *UpdateTerminalAuthResponse) SetReturnMsg(v string) *UpdateTerminalAuthResponse {
  s.ReturnMsg = &v
  return s
}

func (s *UpdateTerminalAuthResponse) SetContent(v *UpdateTerminalAuthContentEntity) *UpdateTerminalAuthResponse {
  s.Content = v
  return s
}

type UpdateTerminalAuthContentEntity struct {
  // {"en":"Username of the specific user.", "zh_CN":"权限名称"}
  TerminalAuthName *string `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" require:"true"`
  // {"en":"Whether to enable the user,1: enable 0: disable", "zh_CN":"是否启用用户，1:启用0:禁用"}
  EnableStatus *int `json:"enableStatus,omitempty" xml:"enableStatus,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"Associated application list ", "zh_CN":"关联的应用列表"}
  AuthorizedResources []*int64 `json:"authorizedResources,omitempty" xml:"authorizedResources,omitempty" type:"Repeated"`
  // {"en":"Associated user list", "zh_CN":"关联的用户列表"}
  AuthorizedUsers []*int64 `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"关联的用户组列表ID"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateTerminalAuthContentEntity) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthContentEntity) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthContentEntity) SetTerminalAuthName(v string) *UpdateTerminalAuthContentEntity {
  s.TerminalAuthName = &v
  return s
}

func (s *UpdateTerminalAuthContentEntity) SetEnableStatus(v int) *UpdateTerminalAuthContentEntity {
  s.EnableStatus = &v
  return s
}

func (s *UpdateTerminalAuthContentEntity) SetRemark(v string) *UpdateTerminalAuthContentEntity {
  s.Remark = &v
  return s
}

func (s *UpdateTerminalAuthContentEntity) SetAuthorizedResources(v []*int64) *UpdateTerminalAuthContentEntity {
  s.AuthorizedResources = v
  return s
}

func (s *UpdateTerminalAuthContentEntity) SetAuthorizedUsers(v []*int64) *UpdateTerminalAuthContentEntity {
  s.AuthorizedUsers = v
  return s
}

func (s *UpdateTerminalAuthContentEntity) SetAuthorizedUserGroupIds(v []*int64) *UpdateTerminalAuthContentEntity {
  s.AuthorizedUserGroupIds = v
  return s
}

type UpdateTerminalAuthPaths struct {
}

func (s UpdateTerminalAuthPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthPaths) GoString() string {
  return s.String()
}

type UpdateTerminalAuthParameters struct {
}

func (s UpdateTerminalAuthParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthParameters) GoString() string {
  return s.String()
}

type UpdateTerminalAuthRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s UpdateTerminalAuthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRequestHeader) SetAuthUser(v string) *UpdateTerminalAuthRequestHeader {
  s.AuthUser = &v
  return s
}

type UpdateTerminalAuthResponseHeader struct {
}

func (s UpdateTerminalAuthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthResponseHeader) GoString() string {
  return s.String()
}




type QueryTerminalAuthInfoRequest struct {
}

func (s QueryTerminalAuthInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoRequest) GoString() string {
  return s.String()
}

type QueryTerminalAuthInfoResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *QueryTerminalAuthInfoContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s QueryTerminalAuthInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoResponse) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoResponse) SetReturnCode(v string) *QueryTerminalAuthInfoResponse {
  s.ReturnCode = &v
  return s
}

func (s *QueryTerminalAuthInfoResponse) SetReturnMsg(v string) *QueryTerminalAuthInfoResponse {
  s.ReturnMsg = &v
  return s
}

func (s *QueryTerminalAuthInfoResponse) SetContent(v *QueryTerminalAuthInfoContentEntity) *QueryTerminalAuthInfoResponse {
  s.Content = v
  return s
}

type QueryTerminalAuthInfoContentEntity struct {
  // {"en":"Username of the specific user.", "zh_CN":"权限名称"}
  TerminalAuthName *string `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" require:"true"`
  // {"en":"Associated application list ", "zh_CN":"关联的应用列表"}
  RelevantResources []*QueryTerminalAuthInfoResourceEntity `json:"relevantResources,omitempty" xml:"relevantResources,omitempty" require:"true" type:"Repeated"`
  // {'en':'remark', 'zh_CN':'备注'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"The status of basic permissions", "zh_CN":"基础权限的状态 0 关闭 1 开启"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated user list ", "zh_CN":"终端用户列表"}
  AuthorizedUsers []*QueryTerminalAuthInfoUserEntity `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ", "zh_CN":"终端用户组列表"}
  AuthorizedUserGroups []*QueryTerminalAuthInfoGroupEntity `json:"authorizedUserGroups,omitempty" xml:"authorizedUserGroups,omitempty" type:"Repeated"`
}

func (s QueryTerminalAuthInfoContentEntity) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoContentEntity) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoContentEntity) SetTerminalAuthName(v string) *QueryTerminalAuthInfoContentEntity {
  s.TerminalAuthName = &v
  return s
}

func (s *QueryTerminalAuthInfoContentEntity) SetRelevantResources(v []*QueryTerminalAuthInfoResourceEntity) *QueryTerminalAuthInfoContentEntity {
  s.RelevantResources = v
  return s
}

func (s *QueryTerminalAuthInfoContentEntity) SetRemark(v string) *QueryTerminalAuthInfoContentEntity {
  s.Remark = &v
  return s
}

func (s *QueryTerminalAuthInfoContentEntity) SetStatus(v int) *QueryTerminalAuthInfoContentEntity {
  s.Status = &v
  return s
}

func (s *QueryTerminalAuthInfoContentEntity) SetTerminalAuthId(v int64) *QueryTerminalAuthInfoContentEntity {
  s.TerminalAuthId = &v
  return s
}

func (s *QueryTerminalAuthInfoContentEntity) SetAuthorizedUsers(v []*QueryTerminalAuthInfoUserEntity) *QueryTerminalAuthInfoContentEntity {
  s.AuthorizedUsers = v
  return s
}

func (s *QueryTerminalAuthInfoContentEntity) SetAuthorizedUserGroups(v []*QueryTerminalAuthInfoGroupEntity) *QueryTerminalAuthInfoContentEntity {
  s.AuthorizedUserGroups = v
  return s
}

type QueryTerminalAuthInfoResourceEntity struct {
  // {"en":"resource name", "zh_CN":"应用名称/资源名称"}
  ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty" require:"true"`
  // {"en":"the resource id, "zh_CN":"应用ID"}
  ResourceId []*int64 `json:"resourceId,omitempty" xml:"resourceId,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTerminalAuthInfoResourceEntity) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoResourceEntity) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoResourceEntity) SetResourceName(v string) *QueryTerminalAuthInfoResourceEntity {
  s.ResourceName = &v
  return s
}

func (s *QueryTerminalAuthInfoResourceEntity) SetResourceId(v []*int64) *QueryTerminalAuthInfoResourceEntity {
  s.ResourceId = v
  return s
}

type QueryTerminalAuthInfoUserEntity struct {
  // {"en":"terminal username", "zh_CN":"用户名称"}
  TerminalUserName *string `json:"terminalUserName,omitempty" xml:"terminalUserName,omitempty" require:"true"`
  // {"en":"the terminal user id, "zh_CN":"终端用户id"}
  TerminalUserId []*int64 `json:"terminalUserId,omitempty" xml:"terminalUserId,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTerminalAuthInfoUserEntity) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoUserEntity) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoUserEntity) SetTerminalUserName(v string) *QueryTerminalAuthInfoUserEntity {
  s.TerminalUserName = &v
  return s
}

func (s *QueryTerminalAuthInfoUserEntity) SetTerminalUserId(v []*int64) *QueryTerminalAuthInfoUserEntity {
  s.TerminalUserId = v
  return s
}

type QueryTerminalAuthInfoGroupEntity struct {
  // {"en":"group name", "zh_CN":"用户组名称"}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {"en":"the group id, "zh_CN":"用户组id"}
  ResourgroupIdceId []*int64 `json:"resourgroupIdceId,omitempty" xml:"resourgroupIdceId,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTerminalAuthInfoGroupEntity) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoGroupEntity) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoGroupEntity) SetGroupName(v string) *QueryTerminalAuthInfoGroupEntity {
  s.GroupName = &v
  return s
}

func (s *QueryTerminalAuthInfoGroupEntity) SetResourgroupIdceId(v []*int64) *QueryTerminalAuthInfoGroupEntity {
  s.ResourgroupIdceId = v
  return s
}

type QueryTerminalAuthInfoPaths struct {
}

func (s QueryTerminalAuthInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoPaths) GoString() string {
  return s.String()
}

type QueryTerminalAuthInfoParameters struct {
  // {"en":"terminal auth name.", "zh_CN":"基础权限的名称"}
  TerminalAuthName *int64 `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" require:"true"`
}

func (s QueryTerminalAuthInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoParameters) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoParameters) SetTerminalAuthName(v int64) *QueryTerminalAuthInfoParameters {
  s.TerminalAuthName = &v
  return s
}

type QueryTerminalAuthInfoRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s QueryTerminalAuthInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthInfoRequestHeader) SetAuthUser(v string) *QueryTerminalAuthInfoRequestHeader {
  s.AuthUser = &v
  return s
}

type QueryTerminalAuthInfoResponseHeader struct {
}

func (s QueryTerminalAuthInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthInfoResponseHeader) GoString() string {
  return s.String()
}




type RemoveTerminalAuthUserOrGroupRequest struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated user list ", "zh_CN":"终端用户名称列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"终端用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s RemoveTerminalAuthUserOrGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupRequest) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthUserOrGroupRequest) SetTerminalAuthId(v int64) *RemoveTerminalAuthUserOrGroupRequest {
  s.TerminalAuthId = &v
  return s
}

func (s *RemoveTerminalAuthUserOrGroupRequest) SetAuthorizedUsers(v []*string) *RemoveTerminalAuthUserOrGroupRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *RemoveTerminalAuthUserOrGroupRequest) SetAuthorizedUserGroupIds(v []*int64) *RemoveTerminalAuthUserOrGroupRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

type RemoveTerminalAuthUserOrGroupResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *RemoveTerminalAuthUserOrGroupContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s RemoveTerminalAuthUserOrGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupResponse) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthUserOrGroupResponse) SetReturnCode(v string) *RemoveTerminalAuthUserOrGroupResponse {
  s.ReturnCode = &v
  return s
}

func (s *RemoveTerminalAuthUserOrGroupResponse) SetReturnMsg(v string) *RemoveTerminalAuthUserOrGroupResponse {
  s.ReturnMsg = &v
  return s
}

func (s *RemoveTerminalAuthUserOrGroupResponse) SetContent(v *RemoveTerminalAuthUserOrGroupContentEntity) *RemoveTerminalAuthUserOrGroupResponse {
  s.Content = v
  return s
}

type RemoveTerminalAuthUserOrGroupContentEntity struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated user list ", "zh_CN":"终端用户名称列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"终端用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s RemoveTerminalAuthUserOrGroupContentEntity) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupContentEntity) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthUserOrGroupContentEntity) SetTerminalAuthId(v int64) *RemoveTerminalAuthUserOrGroupContentEntity {
  s.TerminalAuthId = &v
  return s
}

func (s *RemoveTerminalAuthUserOrGroupContentEntity) SetAuthorizedUsers(v []*string) *RemoveTerminalAuthUserOrGroupContentEntity {
  s.AuthorizedUsers = v
  return s
}

func (s *RemoveTerminalAuthUserOrGroupContentEntity) SetAuthorizedUserGroupIds(v []*int64) *RemoveTerminalAuthUserOrGroupContentEntity {
  s.AuthorizedUserGroupIds = v
  return s
}

type RemoveTerminalAuthUserOrGroupPaths struct {
}

func (s RemoveTerminalAuthUserOrGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupPaths) GoString() string {
  return s.String()
}

type RemoveTerminalAuthUserOrGroupParameters struct {
}

func (s RemoveTerminalAuthUserOrGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupParameters) GoString() string {
  return s.String()
}

type RemoveTerminalAuthUserOrGroupRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s RemoveTerminalAuthUserOrGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupRequestHeader) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthUserOrGroupRequestHeader) SetAuthUser(v string) *RemoveTerminalAuthUserOrGroupRequestHeader {
  s.AuthUser = &v
  return s
}

type RemoveTerminalAuthUserOrGroupResponseHeader struct {
}

func (s RemoveTerminalAuthUserOrGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthUserOrGroupResponseHeader) GoString() string {
  return s.String()
}




type QueryTerminalAuthListRequest struct {
}

func (s QueryTerminalAuthListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthListRequest) GoString() string {
  return s.String()
}

type QueryTerminalAuthListResponse struct {
}

func (s QueryTerminalAuthListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthListResponse) GoString() string {
  return s.String()
}

type QueryTerminalAuthListPaths struct {
}

func (s QueryTerminalAuthListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthListPaths) GoString() string {
  return s.String()
}

type QueryTerminalAuthListParameters struct {
}

func (s QueryTerminalAuthListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthListParameters) GoString() string {
  return s.String()
}

type QueryTerminalAuthListRequestHeader struct {
  // {"en":"UAC account name", "zh_CN":"UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s QueryTerminalAuthListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthListRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthListRequestHeader) SetAuthUser(v string) *QueryTerminalAuthListRequestHeader {
  s.AuthUser = &v
  return s
}

type QueryTerminalAuthListResponseHeader struct {
}

func (s QueryTerminalAuthListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthListResponseHeader) GoString() string {
  return s.String()
}




type AddTerminalAuthUserOrGroupRequest struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated user list ", "zh_CN":"终端用户名称列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"终端用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s AddTerminalAuthUserOrGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupRequest) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthUserOrGroupRequest) SetTerminalAuthId(v int64) *AddTerminalAuthUserOrGroupRequest {
  s.TerminalAuthId = &v
  return s
}

func (s *AddTerminalAuthUserOrGroupRequest) SetAuthorizedUsers(v []*string) *AddTerminalAuthUserOrGroupRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *AddTerminalAuthUserOrGroupRequest) SetAuthorizedUserGroupIds(v []*int64) *AddTerminalAuthUserOrGroupRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

type AddTerminalAuthUserOrGroupResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *AddTerminalAuthUserOrGroupContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s AddTerminalAuthUserOrGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupResponse) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthUserOrGroupResponse) SetReturnCode(v string) *AddTerminalAuthUserOrGroupResponse {
  s.ReturnCode = &v
  return s
}

func (s *AddTerminalAuthUserOrGroupResponse) SetReturnMsg(v string) *AddTerminalAuthUserOrGroupResponse {
  s.ReturnMsg = &v
  return s
}

func (s *AddTerminalAuthUserOrGroupResponse) SetContent(v *AddTerminalAuthUserOrGroupContentEntity) *AddTerminalAuthUserOrGroupResponse {
  s.Content = v
  return s
}

type AddTerminalAuthUserOrGroupContentEntity struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated user list ", "zh_CN":"终端用户名称列表"}
  AuthorizedUsers []*string `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"终端用户组ID列表"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s AddTerminalAuthUserOrGroupContentEntity) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupContentEntity) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthUserOrGroupContentEntity) SetTerminalAuthId(v int64) *AddTerminalAuthUserOrGroupContentEntity {
  s.TerminalAuthId = &v
  return s
}

func (s *AddTerminalAuthUserOrGroupContentEntity) SetAuthorizedUsers(v []*string) *AddTerminalAuthUserOrGroupContentEntity {
  s.AuthorizedUsers = v
  return s
}

func (s *AddTerminalAuthUserOrGroupContentEntity) SetAuthorizedUserGroupIds(v []*int64) *AddTerminalAuthUserOrGroupContentEntity {
  s.AuthorizedUserGroupIds = v
  return s
}

type AddTerminalAuthUserOrGroupPaths struct {
}

func (s AddTerminalAuthUserOrGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupPaths) GoString() string {
  return s.String()
}

type AddTerminalAuthUserOrGroupParameters struct {
}

func (s AddTerminalAuthUserOrGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupParameters) GoString() string {
  return s.String()
}

type AddTerminalAuthUserOrGroupRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s AddTerminalAuthUserOrGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupRequestHeader) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthUserOrGroupRequestHeader) SetAuthUser(v string) *AddTerminalAuthUserOrGroupRequestHeader {
  s.AuthUser = &v
  return s
}

type AddTerminalAuthUserOrGroupResponseHeader struct {
}

func (s AddTerminalAuthUserOrGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthUserOrGroupResponseHeader) GoString() string {
  return s.String()
}




type RemoveTerminalAuthResourceRequest struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated application list ID", "zh_CN":"关联的应用名称列表"}
  RelevantResources []*string `json:"relevantResources,omitempty" xml:"relevantResources,omitempty" type:"Repeated"`
}

func (s RemoveTerminalAuthResourceRequest) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourceRequest) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthResourceRequest) SetTerminalAuthId(v int64) *RemoveTerminalAuthResourceRequest {
  s.TerminalAuthId = &v
  return s
}

func (s *RemoveTerminalAuthResourceRequest) SetRelevantResources(v []*string) *RemoveTerminalAuthResourceRequest {
  s.RelevantResources = v
  return s
}

type RemoveTerminalAuthResourceResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *RemoveTerminalAuthResourceContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s RemoveTerminalAuthResourceResponse) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourceResponse) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthResourceResponse) SetReturnCode(v string) *RemoveTerminalAuthResourceResponse {
  s.ReturnCode = &v
  return s
}

func (s *RemoveTerminalAuthResourceResponse) SetReturnMsg(v string) *RemoveTerminalAuthResourceResponse {
  s.ReturnMsg = &v
  return s
}

func (s *RemoveTerminalAuthResourceResponse) SetContent(v *RemoveTerminalAuthResourceContentEntity) *RemoveTerminalAuthResourceResponse {
  s.Content = v
  return s
}

type RemoveTerminalAuthResourceContentEntity struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated application list ID", "zh_CN":"关联的应用名称列表}
  RelevantResources []*string `json:"relevantResources,omitempty" xml:"relevantResources,omitempty" type:"Repeated"`
}

func (s RemoveTerminalAuthResourceContentEntity) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourceContentEntity) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthResourceContentEntity) SetTerminalAuthId(v int64) *RemoveTerminalAuthResourceContentEntity {
  s.TerminalAuthId = &v
  return s
}

func (s *RemoveTerminalAuthResourceContentEntity) SetRelevantResources(v []*string) *RemoveTerminalAuthResourceContentEntity {
  s.RelevantResources = v
  return s
}

type RemoveTerminalAuthResourcePaths struct {
}

func (s RemoveTerminalAuthResourcePaths) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourcePaths) GoString() string {
  return s.String()
}

type RemoveTerminalAuthResourceParameters struct {
}

func (s RemoveTerminalAuthResourceParameters) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourceParameters) GoString() string {
  return s.String()
}

type RemoveTerminalAuthResourceRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s RemoveTerminalAuthResourceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourceRequestHeader) GoString() string {
  return s.String()
}

func (s *RemoveTerminalAuthResourceRequestHeader) SetAuthUser(v string) *RemoveTerminalAuthResourceRequestHeader {
  s.AuthUser = &v
  return s
}

type RemoveTerminalAuthResourceResponseHeader struct {
}

func (s RemoveTerminalAuthResourceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RemoveTerminalAuthResourceResponseHeader) GoString() string {
  return s.String()
}




type QueryTerminalUserGroupByAuthConfigRequest struct {
  // {"en":"the auth config name", "zh_CN":"身份源名称"}
  AuthConfigName *string `json:"authConfigName,omitempty" xml:"authConfigName,omitempty"`
  // {"en":"terminal group name.", "zh_CN":"用户组名称"}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s QueryTerminalUserGroupByAuthConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigRequest) GoString() string {
  return s.String()
}

func (s *QueryTerminalUserGroupByAuthConfigRequest) SetAuthConfigName(v string) *QueryTerminalUserGroupByAuthConfigRequest {
  s.AuthConfigName = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigRequest) SetGroupName(v string) *QueryTerminalUserGroupByAuthConfigRequest {
  s.GroupName = &v
  return s
}

type QueryTerminalUserGroupByAuthConfigResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *QueryTerminalUserGroupByAuthConfigContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s QueryTerminalUserGroupByAuthConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryTerminalUserGroupByAuthConfigResponse) SetReturnCode(v string) *QueryTerminalUserGroupByAuthConfigResponse {
  s.ReturnCode = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigResponse) SetReturnMsg(v string) *QueryTerminalUserGroupByAuthConfigResponse {
  s.ReturnMsg = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigResponse) SetContent(v *QueryTerminalUserGroupByAuthConfigContentEntity) *QueryTerminalUserGroupByAuthConfigResponse {
  s.Content = v
  return s
}

type QueryTerminalUserGroupByAuthConfigContentEntity struct {
  // {'en':'groupList', 'zh_CN':'group集合，下面为字段'}
  GroupList []*QueryTerminalUserGroupByAuthConfigContentEntityGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTerminalUserGroupByAuthConfigContentEntity) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigContentEntity) GoString() string {
  return s.String()
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntity) SetGroupList(v []*QueryTerminalUserGroupByAuthConfigContentEntityGroupList) *QueryTerminalUserGroupByAuthConfigContentEntity {
  s.GroupList = v
  return s
}

type QueryTerminalUserGroupByAuthConfigContentEntityGroupList struct     {
  // {'en':'groupName', 'zh_CN':'用户组名称'}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {'en':'groupId', 'zh_CN':'用户组ID'}
  GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {'en':'parentGroupName', 'zh_CN':'父组名称'}
  ParentGroupName *string `json:"parentGroupName,omitempty" xml:"parentGroupName,omitempty" require:"true"`
  // {'en':'parentGroupID', 'zh_CN':'父组ID'}
  ParentGroupID *string `json:"parentGroupID,omitempty" xml:"parentGroupID,omitempty" require:"true"`
  // {'en':'remark', 'zh_CN':'备注'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {'en':'createTime', 'zh_CN':'创建时间'}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s QueryTerminalUserGroupByAuthConfigContentEntityGroupList) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigContentEntityGroupList) GoString() string {
  return s.String()
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntityGroupList) SetGroupName(v string) *QueryTerminalUserGroupByAuthConfigContentEntityGroupList {
  s.GroupName = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntityGroupList) SetGroupId(v int64) *QueryTerminalUserGroupByAuthConfigContentEntityGroupList {
  s.GroupId = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntityGroupList) SetParentGroupName(v string) *QueryTerminalUserGroupByAuthConfigContentEntityGroupList {
  s.ParentGroupName = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntityGroupList) SetParentGroupID(v string) *QueryTerminalUserGroupByAuthConfigContentEntityGroupList {
  s.ParentGroupID = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntityGroupList) SetRemark(v string) *QueryTerminalUserGroupByAuthConfigContentEntityGroupList {
  s.Remark = &v
  return s
}

func (s *QueryTerminalUserGroupByAuthConfigContentEntityGroupList) SetCreateTime(v int64) *QueryTerminalUserGroupByAuthConfigContentEntityGroupList {
  s.CreateTime = &v
  return s
}

type QueryTerminalUserGroupByAuthConfigPaths struct {
}

func (s QueryTerminalUserGroupByAuthConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigPaths) GoString() string {
  return s.String()
}

type QueryTerminalUserGroupByAuthConfigParameters struct {
}

func (s QueryTerminalUserGroupByAuthConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigParameters) GoString() string {
  return s.String()
}

type QueryTerminalUserGroupByAuthConfigRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s QueryTerminalUserGroupByAuthConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryTerminalUserGroupByAuthConfigRequestHeader) SetAuthUser(v string) *QueryTerminalUserGroupByAuthConfigRequestHeader {
  s.AuthUser = &v
  return s
}

type QueryTerminalUserGroupByAuthConfigResponseHeader struct {
}

func (s QueryTerminalUserGroupByAuthConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalUserGroupByAuthConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateTerminalAuthRelatedResourcesRequest struct {
  // {"en":"the identify of terminal auth.", "zh_CN":"基础权限的ID"}
  TerminalAuthId *int64 `json:"terminalAuthId,omitempty" xml:"terminalAuthId,omitempty" require:"true"`
  // {"en":"Associated application list ID", "zh_CN":"关联的应用列表ID"}
  RelevantResources []*string `json:"relevantResources,omitempty" xml:"relevantResources,omitempty" type:"Repeated"`
}

func (s UpdateTerminalAuthRelatedResourcesRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesRequest) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRelatedResourcesRequest) SetTerminalAuthId(v int64) *UpdateTerminalAuthRelatedResourcesRequest {
  s.TerminalAuthId = &v
  return s
}

func (s *UpdateTerminalAuthRelatedResourcesRequest) SetRelevantResources(v []*string) *UpdateTerminalAuthRelatedResourcesRequest {
  s.RelevantResources = v
  return s
}

type UpdateTerminalAuthRelatedResourcesResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *UpdateTerminalAuthRelatedResourcesContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s UpdateTerminalAuthRelatedResourcesResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesResponse) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRelatedResourcesResponse) SetReturnCode(v string) *UpdateTerminalAuthRelatedResourcesResponse {
  s.ReturnCode = &v
  return s
}

func (s *UpdateTerminalAuthRelatedResourcesResponse) SetReturnMsg(v string) *UpdateTerminalAuthRelatedResourcesResponse {
  s.ReturnMsg = &v
  return s
}

func (s *UpdateTerminalAuthRelatedResourcesResponse) SetContent(v *UpdateTerminalAuthRelatedResourcesContentEntity) *UpdateTerminalAuthRelatedResourcesResponse {
  s.Content = v
  return s
}

type UpdateTerminalAuthRelatedResourcesContentEntity struct {
  // {"en":"Username of the specific user.", "zh_CN":"权限名称"}
  TerminalAuthName *string `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" require:"true"`
  // {"en":"Associated application list ID", "zh_CN":"关联的应用列表ID"}
  RelevantResources []*string `json:"relevantResources,omitempty" xml:"relevantResources,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateTerminalAuthRelatedResourcesContentEntity) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesContentEntity) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRelatedResourcesContentEntity) SetTerminalAuthName(v string) *UpdateTerminalAuthRelatedResourcesContentEntity {
  s.TerminalAuthName = &v
  return s
}

func (s *UpdateTerminalAuthRelatedResourcesContentEntity) SetRelevantResources(v []*string) *UpdateTerminalAuthRelatedResourcesContentEntity {
  s.RelevantResources = v
  return s
}

type UpdateTerminalAuthRelatedResourcesPaths struct {
}

func (s UpdateTerminalAuthRelatedResourcesPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesPaths) GoString() string {
  return s.String()
}

type UpdateTerminalAuthRelatedResourcesParameters struct {
  // {"en":"Username of the specific user.", "zh_CN":"权限名称"}
  TerminalAuthName *string `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" require:"true"`
  // {"en":"Associated application list ID", "zh_CN":"关联的应用列表ID"}
  RelevantResources []*string `json:"relevantResources,omitempty" xml:"relevantResources,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateTerminalAuthRelatedResourcesParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesParameters) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRelatedResourcesParameters) SetTerminalAuthName(v string) *UpdateTerminalAuthRelatedResourcesParameters {
  s.TerminalAuthName = &v
  return s
}

func (s *UpdateTerminalAuthRelatedResourcesParameters) SetRelevantResources(v []*string) *UpdateTerminalAuthRelatedResourcesParameters {
  s.RelevantResources = v
  return s
}

type UpdateTerminalAuthRelatedResourcesRequestHeader struct {
  // {"en":"uac account name", "zh_CN":"当前企业的UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s UpdateTerminalAuthRelatedResourcesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateTerminalAuthRelatedResourcesRequestHeader) SetAuthUser(v string) *UpdateTerminalAuthRelatedResourcesRequestHeader {
  s.AuthUser = &v
  return s
}

type UpdateTerminalAuthRelatedResourcesResponseHeader struct {
}

func (s UpdateTerminalAuthRelatedResourcesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateTerminalAuthRelatedResourcesResponseHeader) GoString() string {
  return s.String()
}




type QueryTerminalAuthResourceListRequest struct {
}

func (s QueryTerminalAuthResourceListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListRequest) GoString() string {
  return s.String()
}

type QueryTerminalAuthResourceListResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *QueryTerminalAuthResourceListContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s QueryTerminalAuthResourceListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListResponse) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthResourceListResponse) SetReturnCode(v string) *QueryTerminalAuthResourceListResponse {
  s.ReturnCode = &v
  return s
}

func (s *QueryTerminalAuthResourceListResponse) SetReturnMsg(v string) *QueryTerminalAuthResourceListResponse {
  s.ReturnMsg = &v
  return s
}

func (s *QueryTerminalAuthResourceListResponse) SetContent(v *QueryTerminalAuthResourceListContentEntity) *QueryTerminalAuthResourceListResponse {
  s.Content = v
  return s
}

type QueryTerminalAuthResourceListContentEntity struct {
  // {"en":"resource name", "zh_CN":"应用名称/资源名称"}
  ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty" require:"true"`
  // {"en":"Description of the resource", "zh_CN":"备注"}
  Remark []*string `json:"remark,omitempty" xml:"remark,omitempty" require:"true" type:"Repeated"`
  // {"en":"the resource id, "zh_CN":"应用ID"}
  ResourceId []*int64 `json:"resourceId,omitempty" xml:"resourceId,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTerminalAuthResourceListContentEntity) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListContentEntity) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthResourceListContentEntity) SetResourceName(v string) *QueryTerminalAuthResourceListContentEntity {
  s.ResourceName = &v
  return s
}

func (s *QueryTerminalAuthResourceListContentEntity) SetRemark(v []*string) *QueryTerminalAuthResourceListContentEntity {
  s.Remark = v
  return s
}

func (s *QueryTerminalAuthResourceListContentEntity) SetResourceId(v []*int64) *QueryTerminalAuthResourceListContentEntity {
  s.ResourceId = v
  return s
}

type QueryTerminalAuthResourceListPaths struct {
}

func (s QueryTerminalAuthResourceListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListPaths) GoString() string {
  return s.String()
}

type QueryTerminalAuthResourceListParameters struct {
}

func (s QueryTerminalAuthResourceListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListParameters) GoString() string {
  return s.String()
}

type QueryTerminalAuthResourceListRequestHeader struct {
  // {"en":"UAC account name", "zh_CN":"UAC主账号名"}
  AuthUser *string `json:"authUser,omitempty" xml:"authUser,omitempty" require:"true"`
}

func (s QueryTerminalAuthResourceListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryTerminalAuthResourceListRequestHeader) SetAuthUser(v string) *QueryTerminalAuthResourceListRequestHeader {
  s.AuthUser = &v
  return s
}

type QueryTerminalAuthResourceListResponseHeader struct {
}

func (s QueryTerminalAuthResourceListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTerminalAuthResourceListResponseHeader) GoString() string {
  return s.String()
}




type AddTerminalAuthRequest struct {
  // {"en":"Username of the specific user.", "zh_CN":"权限名称"}
  TerminalAuthName *string `json:"terminalAuthName,omitempty" xml:"terminalAuthName,omitempty" require:"true"`
  // {"en":"Whether to enable the user,1: enable 0: disable", "zh_CN":"是否启用该权限，1:启用0:禁用"}
  EnableStatus *int `json:"enableStatus,omitempty" xml:"enableStatus,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"Associated application list ", "zh_CN":"关联的应用列表"}
  AuthResources []*int64 `json:"authResources,omitempty" xml:"authResources,omitempty" type:"Repeated"`
  // {"en":"Associated user list", "zh_CN":"关联的用户列表"}
  AuthorizedUsers []*int64 `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"关联的用户组列表ID"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s AddTerminalAuthRequest) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthRequest) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthRequest) SetTerminalAuthName(v string) *AddTerminalAuthRequest {
  s.TerminalAuthName = &v
  return s
}

func (s *AddTerminalAuthRequest) SetEnableStatus(v int) *AddTerminalAuthRequest {
  s.EnableStatus = &v
  return s
}

func (s *AddTerminalAuthRequest) SetRemark(v string) *AddTerminalAuthRequest {
  s.Remark = &v
  return s
}

func (s *AddTerminalAuthRequest) SetAuthResources(v []*int64) *AddTerminalAuthRequest {
  s.AuthResources = v
  return s
}

func (s *AddTerminalAuthRequest) SetAuthorizedUsers(v []*int64) *AddTerminalAuthRequest {
  s.AuthorizedUsers = v
  return s
}

func (s *AddTerminalAuthRequest) SetAuthorizedUserGroupIds(v []*int64) *AddTerminalAuthRequest {
  s.AuthorizedUserGroupIds = v
  return s
}

type AddTerminalAuthResponse struct {
  // {"en":"Interface error code, 0-fail,1-success", "zh_CN":"接口错误码，0-代表失败，1-代表成功"}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message", "zh_CN":"错误信息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {"en":"content", "zh_CN":"数据，下面全是数据的内容"}
  Content *AddTerminalAuthContentEntity `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s AddTerminalAuthResponse) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthResponse) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthResponse) SetReturnCode(v string) *AddTerminalAuthResponse {
  s.ReturnCode = &v
  return s
}

func (s *AddTerminalAuthResponse) SetReturnMsg(v string) *AddTerminalAuthResponse {
  s.ReturnMsg = &v
  return s
}

func (s *AddTerminalAuthResponse) SetContent(v *AddTerminalAuthContentEntity) *AddTerminalAuthResponse {
  s.Content = v
  return s
}

type AddTerminalAuthContentEntity struct {
  // {"en":"Username of the specific user.", "zh_CN":"权限名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Whether to enable the user,1: enable 0: disable", "zh_CN":"是否启用用户，1:启用0:禁用"}
  EnableStatus *int `json:"enableStatus,omitempty" xml:"enableStatus,omitempty" require:"true"`
  // {"en":"Remark,Maximum length is 255 characters.", "zh_CN":"备注最大长度255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" maxLength:"255"`
  // {"en":"Associated application list", "zh_CN":"关联的应用列表"}
  AuthorizedResources []*int64 `json:"authorizedResources,omitempty" xml:"authorizedResources,omitempty" type:"Repeated"`
  // {"en":"Associated user list ", "zh_CN":"关联的用户列表"}
  AuthorizedUsers []*int64 `json:"authorizedUsers,omitempty" xml:"authorizedUsers,omitempty" type:"Repeated"`
  // {"en":"Associated user group list ID", "zh_CN":"关联的用户组列表ID"}
  AuthorizedUserGroupIds []*int64 `json:"authorizedUserGroupIds,omitempty" xml:"authorizedUserGroupIds,omitempty" type:"Repeated"`
}

func (s AddTerminalAuthContentEntity) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthContentEntity) GoString() string {
  return s.String()
}

func (s *AddTerminalAuthContentEntity) SetName(v string) *AddTerminalAuthContentEntity {
  s.Name = &v
  return s
}

func (s *AddTerminalAuthContentEntity) SetEnableStatus(v int) *AddTerminalAuthContentEntity {
  s.EnableStatus = &v
  return s
}

func (s *AddTerminalAuthContentEntity) SetRemark(v string) *AddTerminalAuthContentEntity {
  s.Remark = &v
  return s
}

func (s *AddTerminalAuthContentEntity) SetAuthorizedResources(v []*int64) *AddTerminalAuthContentEntity {
  s.AuthorizedResources = v
  return s
}

func (s *AddTerminalAuthContentEntity) SetAuthorizedUsers(v []*int64) *AddTerminalAuthContentEntity {
  s.AuthorizedUsers = v
  return s
}

func (s *AddTerminalAuthContentEntity) SetAuthorizedUserGroupIds(v []*int64) *AddTerminalAuthContentEntity {
  s.AuthorizedUserGroupIds = v
  return s
}

type AddTerminalAuthPaths struct {
}

func (s AddTerminalAuthPaths) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthPaths) GoString() string {
  return s.String()
}

type AddTerminalAuthParameters struct {
}

func (s AddTerminalAuthParameters) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthParameters) GoString() string {
  return s.String()
}

type AddTerminalAuthRequestHeader struct {
}

func (s AddTerminalAuthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthRequestHeader) GoString() string {
  return s.String()
}

type AddTerminalAuthResponseHeader struct {
}

func (s AddTerminalAuthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddTerminalAuthResponseHeader) GoString() string {
  return s.String()
}




