package playback

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPublishCodeRequest struct {
  // {"en":"Video ID", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"
  // Play code type
  // Value range:
  // 0(all)
  // 2(swf code)
  // 4(Video URL)
  // 5(Adaptive code)
  // 6(Try watch video URL)
  // 7(Try watch adaptive code)
  // 8(The customized encrypted play code)
  // The default is 0;
  // General licensed video only has adaptive/video URL.
  // Normal encrypted video only swf/ customer custom/video URL.
  // Non-encrypted video only has swf/ Adaptive/Custom/video URL", "zh_CN":"播放代码类型
  // 取值范围 ：
  // 0(全部)
  // 2(swf代码)
  // 4(视频URL)
  // 5(自适应代码)
  // 6(试看视频URL)
  // 7(试看自适应代码)
  // 8(加密客户定制的播放代码)
  // 默认为0；
  // 通用授权视频只有自适应/视频URL。
  // 普通加密视频只swf/客户自定义/视频URL。
  // 非加密视频只有swf/自适应/客户自定义/视频URL"}
  CodeType *int32 `json:"codeType,omitempty" xml:"codeType,omitempty"`
}

func (s GetPublishCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeRequest) GoString() string {
  return s.String()
}

func (s *GetPublishCodeRequest) SetVideoId(v string) *GetPublishCodeRequest {
  s.VideoId = &v
  return s
}

func (s *GetPublishCodeRequest) SetCodeType(v int32) *GetPublishCodeRequest {
  s.CodeType = &v
  return s
}

type GetPublishCodeResponse struct {
  // {'en':'Status code', 'zh_CN':'返回状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Data *GetPublishCodeGetPublishCodeData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetPublishCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeResponse) GoString() string {
  return s.String()
}

func (s *GetPublishCodeResponse) SetCode(v int32) *GetPublishCodeResponse {
  s.Code = &v
  return s
}

func (s *GetPublishCodeResponse) SetData(v *GetPublishCodeGetPublishCodeData) *GetPublishCodeResponse {
  s.Data = v
  return s
}

func (s *GetPublishCodeResponse) SetMessage(v string) *GetPublishCodeResponse {
  s.Message = &v
  return s
}

type GetPublishCodeGetPublishCodeData struct {
  // {'en':'videoId', 'zh_CN':'视频ID'}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {'en':'
  // Whether the video is encrypted
  // Value range:
  // 0(unencrypted)
  // 1(encryption)', 'zh_CN':'视频是否加密
  // 取值范围 ：
  // 0(不加密)
  // 1(加密)'}
  Encrypt *int32 `json:"encrypt,omitempty" xml:"encrypt,omitempty" require:"true"`
  // {'en':'Video adaptive code, encrypted video is empty', 'zh_CN':'视频自适应代码，加密视频为空'}
  AutoCode *string `json:"autoCode,omitempty" xml:"autoCode,omitempty" require:"true"`
  // {'en':'Video swf code', 'zh_CN':'视频swf代码'}
  SwfCode *string `json:"swfCode,omitempty" xml:"swfCode,omitempty" require:"true"`
  // {'en':'Custom play code. The default value is null. If you have personalized needs, please contact customer service.', 'zh_CN':'客户定制的播放代码。默认为空。如有个性化需求，请与客服联系。'}
  CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty" require:"true"`
  VideoUrl []*GetPublishCodeGetPublishCodeDataVideoUrl `json:"videoUrl,omitempty" xml:"videoUrl,omitempty" require:"true" type:"Repeated"`
}

func (s GetPublishCodeGetPublishCodeData) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeGetPublishCodeData) GoString() string {
  return s.String()
}

func (s *GetPublishCodeGetPublishCodeData) SetVideoId(v string) *GetPublishCodeGetPublishCodeData {
  s.VideoId = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeData) SetEncrypt(v int32) *GetPublishCodeGetPublishCodeData {
  s.Encrypt = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeData) SetAutoCode(v string) *GetPublishCodeGetPublishCodeData {
  s.AutoCode = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeData) SetSwfCode(v string) *GetPublishCodeGetPublishCodeData {
  s.SwfCode = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeData) SetCustomCode(v string) *GetPublishCodeGetPublishCodeData {
  s.CustomCode = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeData) SetVideoUrl(v []*GetPublishCodeGetPublishCodeDataVideoUrl) *GetPublishCodeGetPublishCodeData {
  s.VideoUrl = v
  return s
}

type GetPublishCodeGetPublishCodeDataVideoUrl struct     {
  // {'en':'Smooth bit rate video url', 'zh_CN':'流畅码率视频url'}
  FluentUrl *string `json:"fluentUrl,omitempty" xml:"fluentUrl,omitempty" require:"true"`
  // {'en':'Ultra clear bit rate video url', 'zh_CN':'超清码率视频url'}
  HdPullUrl *string `json:"hdPullUrl,omitempty" xml:"hdPullUrl,omitempty" require:"true"`
  // {'en':'Hd bit rate video url', 'zh_CN':'高清码率视频url'}
  HighUrl *string `json:"highUrl,omitempty" xml:"highUrl,omitempty" require:"true"`
  // {'en':'Original video url', 'zh_CN':'原画视频url'}
  OriginUrl *string `json:"originUrl,omitempty" xml:"originUrl,omitempty" require:"true"`
  // {'en':'Standard definition bit rate video url', 'zh_CN':'标清码率视频url'}
  SdUrl *string `json:"sdUrl,omitempty" xml:"sdUrl,omitempty" require:"true"`
  // {'en':'PC/mobile terminal', 'zh_CN':'PC端/移动端'}
  UrlType *string `json:"urlType,omitempty" xml:"urlType,omitempty" require:"true"`
}

func (s GetPublishCodeGetPublishCodeDataVideoUrl) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeGetPublishCodeDataVideoUrl) GoString() string {
  return s.String()
}

func (s *GetPublishCodeGetPublishCodeDataVideoUrl) SetFluentUrl(v string) *GetPublishCodeGetPublishCodeDataVideoUrl {
  s.FluentUrl = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeDataVideoUrl) SetHdPullUrl(v string) *GetPublishCodeGetPublishCodeDataVideoUrl {
  s.HdPullUrl = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeDataVideoUrl) SetHighUrl(v string) *GetPublishCodeGetPublishCodeDataVideoUrl {
  s.HighUrl = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeDataVideoUrl) SetOriginUrl(v string) *GetPublishCodeGetPublishCodeDataVideoUrl {
  s.OriginUrl = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeDataVideoUrl) SetSdUrl(v string) *GetPublishCodeGetPublishCodeDataVideoUrl {
  s.SdUrl = &v
  return s
}

func (s *GetPublishCodeGetPublishCodeDataVideoUrl) SetUrlType(v string) *GetPublishCodeGetPublishCodeDataVideoUrl {
  s.UrlType = &v
  return s
}

type GetPublishCodePaths struct {
}

func (s GetPublishCodePaths) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodePaths) GoString() string {
  return s.String()
}

type GetPublishCodeParameters struct {
}

func (s GetPublishCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeParameters) GoString() string {
  return s.String()
}

type GetPublishCodeRequestHeader struct {
}

func (s GetPublishCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeRequestHeader) GoString() string {
  return s.String()
}

type GetPublishCodeResponseHeader struct {
}

func (s GetPublishCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPublishCodeResponseHeader) GoString() string {
  return s.String()
}




