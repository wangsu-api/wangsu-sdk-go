package nac

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type EnableNacUserRequest struct {
  // {"en":"User ID. The value will be used to search user if not null","zh_CN":"用户ID. 如果该项非空,则使用该值查询用户, 否则使用userName"}
  UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
  // {"en":"User name. The field is required when userId is null","zh_CN":"用户名称. 如果userId为空, 必须要传userName"}
  UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s EnableNacUserRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableNacUserRequest) GoString() string {
  return s.String()
}

func (s *EnableNacUserRequest) SetUserId(v int64) *EnableNacUserRequest {
  s.UserId = &v
  return s
}

func (s *EnableNacUserRequest) SetUserName(v string) *EnableNacUserRequest {
  s.UserName = &v
  return s
}

type EnableNacUserRequestHeader struct {
}

func (s EnableNacUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableNacUserRequestHeader) GoString() string {
  return s.String()
}

type EnableNacUserPaths struct {
}

func (s EnableNacUserPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableNacUserPaths) GoString() string {
  return s.String()
}

type EnableNacUserParameters struct {
}

func (s EnableNacUserParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableNacUserParameters) GoString() string {
  return s.String()
}

type EnableNacUserResponse struct {
  // {"en":"Interface error code. 1-success,The other number-fail","zh_CN":"接口错误码. 1代表成功. 非1代表失败"}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误消息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
}

func (s EnableNacUserResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableNacUserResponse) GoString() string {
  return s.String()
}

func (s *EnableNacUserResponse) SetReturnCode(v int) *EnableNacUserResponse {
  s.ReturnCode = &v
  return s
}

func (s *EnableNacUserResponse) SetReturnMsg(v string) *EnableNacUserResponse {
  s.ReturnMsg = &v
  return s
}

type EnableNacUserResponseHeader struct {
}

func (s EnableNacUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableNacUserResponseHeader) GoString() string {
  return s.String()
}




type DisableNacUserRequest struct {
  // {"en":"User ID. The value will be used to search user if not null","zh_CN":"用户ID. 如果该项非空,则使用该值查询用户, 否则使用userName"}
  UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
  // {"en":"User name. The field is required when userId is null","zh_CN":"用户名称. 如果userId为空, 必须要传userName"}
  UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DisableNacUserRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableNacUserRequest) GoString() string {
  return s.String()
}

func (s *DisableNacUserRequest) SetUserId(v int64) *DisableNacUserRequest {
  s.UserId = &v
  return s
}

func (s *DisableNacUserRequest) SetUserName(v string) *DisableNacUserRequest {
  s.UserName = &v
  return s
}

type DisableNacUserRequestHeader struct {
}

func (s DisableNacUserRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableNacUserRequestHeader) GoString() string {
  return s.String()
}

type DisableNacUserPaths struct {
}

func (s DisableNacUserPaths) String() string {
  return tea.Prettify(s)
}

func (s DisableNacUserPaths) GoString() string {
  return s.String()
}

type DisableNacUserParameters struct {
}

func (s DisableNacUserParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableNacUserParameters) GoString() string {
  return s.String()
}

type DisableNacUserResponse struct {
  // {"en":"Interface error code. 1-success,The other number-fail","zh_CN":"接口错误码. 1代表成功. 非1代表失败"}
  ReturnCode *int `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {"en":"Error message","zh_CN":"错误消息"}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
}

func (s DisableNacUserResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableNacUserResponse) GoString() string {
  return s.String()
}

func (s *DisableNacUserResponse) SetReturnCode(v int) *DisableNacUserResponse {
  s.ReturnCode = &v
  return s
}

func (s *DisableNacUserResponse) SetReturnMsg(v string) *DisableNacUserResponse {
  s.ReturnMsg = &v
  return s
}

type DisableNacUserResponseHeader struct {
}

func (s DisableNacUserResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableNacUserResponseHeader) GoString() string {
  return s.String()
}




