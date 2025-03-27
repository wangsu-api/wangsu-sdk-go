package clientsdk

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type SecureLinkSdkAuthRequest struct {
}

func (s SecureLinkSdkAuthRequest) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthRequest) GoString() string {
  return s.String()
}

type SecureLinkSdkAuthResponse struct {
  // {'en':'Interface error code, 0-fail,1-success', 'zh_CN':'接口错误码，0-代表失败，1-代表成功'}
  ReturnCode *string `json:"returnCode,omitempty" xml:"returnCode,omitempty" require:"true"`
  // {'en':'Error message', 'zh_CN':'错误信息'}
  ReturnMsg *string `json:"returnMsg,omitempty" xml:"returnMsg,omitempty" require:"true"`
  // {'en':'content', 'zh_CN':'数据，下面全是数据的内容'}
  Content *SecureLinkSdkAuthSecureLinkSdkAuthResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
}

func (s SecureLinkSdkAuthResponse) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthResponse) GoString() string {
  return s.String()
}

func (s *SecureLinkSdkAuthResponse) SetReturnCode(v string) *SecureLinkSdkAuthResponse {
  s.ReturnCode = &v
  return s
}

func (s *SecureLinkSdkAuthResponse) SetReturnMsg(v string) *SecureLinkSdkAuthResponse {
  s.ReturnMsg = &v
  return s
}

func (s *SecureLinkSdkAuthResponse) SetContent(v *SecureLinkSdkAuthSecureLinkSdkAuthResponseContent) *SecureLinkSdkAuthResponse {
  s.Content = v
  return s
}

type SecureLinkSdkAuthSecureLinkSdkAuthResponseContent struct {
  // {'en':'sdkToken', 'zh_CN':'sdk token'}
  SdkToken *string `json:"sdkToken,omitempty" xml:"sdkToken,omitempty" require:"true"`
  // {'en':'issueTime', 'zh_CN':'下发时间'}
  IssueTime *int64 `json:"issueTime,omitempty" xml:"issueTime,omitempty" require:"true"`
}

func (s SecureLinkSdkAuthSecureLinkSdkAuthResponseContent) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthSecureLinkSdkAuthResponseContent) GoString() string {
  return s.String()
}

func (s *SecureLinkSdkAuthSecureLinkSdkAuthResponseContent) SetSdkToken(v string) *SecureLinkSdkAuthSecureLinkSdkAuthResponseContent {
  s.SdkToken = &v
  return s
}

func (s *SecureLinkSdkAuthSecureLinkSdkAuthResponseContent) SetIssueTime(v int64) *SecureLinkSdkAuthSecureLinkSdkAuthResponseContent {
  s.IssueTime = &v
  return s
}

type SecureLinkSdkAuthPaths struct {
}

func (s SecureLinkSdkAuthPaths) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthPaths) GoString() string {
  return s.String()
}

type SecureLinkSdkAuthParameters struct {
}

func (s SecureLinkSdkAuthParameters) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthParameters) GoString() string {
  return s.String()
}

type SecureLinkSdkAuthRequestHeader struct {
}

func (s SecureLinkSdkAuthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthRequestHeader) GoString() string {
  return s.String()
}

type SecureLinkSdkAuthResponseHeader struct {
}

func (s SecureLinkSdkAuthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SecureLinkSdkAuthResponseHeader) GoString() string {
  return s.String()
}




