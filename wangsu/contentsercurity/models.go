package contentsercurity

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ImageAuditRequest struct {
  // {"en":"userName", "zh_CN":"客户登陆账号"}
  UserName *string `json:"userName,omitempty" xml:"userName,omitempty" require:"true"`
  // {"en":"Non-mandatory field. If left blank, all categories are assumed. This field pertains to image audit category and supports multiple configurations. The categories include POLITY: political identification; VIOLENT: violent, terrorist, and prohibited identification; EROTIC: pornography identification; ADVERT: advertising identification; IMGTEXTRISK: image and text OCR identification. Multiple underscores are used to separate them. For example, all: POLITY_VIOLENT_EROTIC_ADVERT_IMGTEXTRISK", "zh_CN":"非必填项，为空默认是全部。图片审计类别，支持配置多个。POLITY：涉政识别； VIOLENT：暴恐违禁识别；EROTIC：色情识别；ADVERT：广告识别；IMGTEXTRISK：图文OCR识别；多个下划线分隔。例如全部：POLITY_VIOLENT_EROTIC_ADVERT_IMGTEXTRISK"}
  AuditType *string `json:"auditType,omitempty" xml:"auditType,omitempty"`
  // {"en":"This field is optional, and the default value is set to 'default'. If the customer has specific requirements for detection rules, Wangsu will need to provide them separately.", "zh_CN":"非必填项，默认值为“default”。若客户有特殊的检测规则要求，则需要网宿单独分配提供。"}
  AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
  // {"en":"Optional, precision requirement, default value is 'balance'. Optional values are: balance (balanced recall), recall (high recall), precise (high accuracy)", "zh_CN":"非必填项，精度要求，默认值为“balance”。可选值为：balance（准召平衡）、recall（高召回）、precise（高准确）"}
  Precision *string `json:"precision,omitempty" xml:"precision,omitempty"`
  // {"en":"The picture array contains picture content and picture IDs. Each picture ID cannot be repeated in a single request. A maximum of 12 pictures can be uploaded in one request.", "zh_CN":"图片数组：里面含图片内容，图片Id。图片Id一次请求内不能重复，一次请求最多可上传12张图片。"}
  Images []*ImageAuditImgDto `json:"images,omitempty" xml:"images,omitempty" require:"true" type:"Repeated"`
}

func (s ImageAuditRequest) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditRequest) GoString() string {
  return s.String()
}

func (s *ImageAuditRequest) SetUserName(v string) *ImageAuditRequest {
  s.UserName = &v
  return s
}

func (s *ImageAuditRequest) SetAuditType(v string) *ImageAuditRequest {
  s.AuditType = &v
  return s
}

func (s *ImageAuditRequest) SetAppType(v string) *ImageAuditRequest {
  s.AppType = &v
  return s
}

func (s *ImageAuditRequest) SetPrecision(v string) *ImageAuditRequest {
  s.Precision = &v
  return s
}

func (s *ImageAuditRequest) SetImages(v []*ImageAuditImgDto) *ImageAuditRequest {
  s.Images = v
  return s
}

type ImageAuditResponse struct {
  // {"en":"code", "zh_CN":"请求返回码，1100成功，其他是8位失败码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"错误码具体信息描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"网宿生成的请求唯一id"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"imgs", "zh_CN":"查询结果"}
  Imgs []*ImageAuditImgQueryResp `json:"imgs,omitempty" xml:"imgs,omitempty" require:"true" type:"Repeated"`
}

func (s ImageAuditResponse) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditResponse) GoString() string {
  return s.String()
}

func (s *ImageAuditResponse) SetCode(v string) *ImageAuditResponse {
  s.Code = &v
  return s
}

func (s *ImageAuditResponse) SetMessage(v string) *ImageAuditResponse {
  s.Message = &v
  return s
}

func (s *ImageAuditResponse) SetRequestId(v string) *ImageAuditResponse {
  s.RequestId = &v
  return s
}

func (s *ImageAuditResponse) SetImgs(v []*ImageAuditImgQueryResp) *ImageAuditResponse {
  s.Imgs = v
  return s
}

type ImageAuditPaths struct {
}

func (s ImageAuditPaths) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditPaths) GoString() string {
  return s.String()
}

type ImageAuditParameters struct {
}

func (s ImageAuditParameters) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditParameters) GoString() string {
  return s.String()
}

type ImageAuditRequestHeader struct {
}

func (s ImageAuditRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditRequestHeader) GoString() string {
  return s.String()
}

type ImageAuditResponseHeader struct {
}

func (s ImageAuditResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditResponseHeader) GoString() string {
  return s.String()
}

type ImageAuditImgDto struct {
  // {"en":"content", "zh_CN":"上传要检测的图片，支持两种方式：
  // 1、上传图片，图片必须是base64格式编码；
  // 2、上传图片的URL；
  // 上传的图片只支持如下格式：jpg，jpeg，png，webp，gif，tiff，tif, hief，建议图片像素不小于256*256，目前最低支持20*20分辨率的图片。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"btId", "zh_CN":"图⽚唯⼀标识,同⼀次请求中不可重复，btId⻓度 在30以内"}
  BtId *string `json:"btId,omitempty" xml:"btId,omitempty" require:"true"`
}

func (s ImageAuditImgDto) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditImgDto) GoString() string {
  return s.String()
}

func (s *ImageAuditImgDto) SetContent(v string) *ImageAuditImgDto {
  s.Content = &v
  return s
}

func (s *ImageAuditImgDto) SetBtId(v string) *ImageAuditImgDto {
  s.BtId = &v
  return s
}

type ImageAuditImgQueryResp struct {
  // {"en":"requestId", "zh_CN":"网宿返回的请求id"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"btId", "zh_CN":"图⽚唯⼀标识,同⼀次请求中不可重复"}
  BtId *string `json:"btId,omitempty" xml:"btId,omitempty" require:"true"`
  // {"en":"code", "zh_CN":"此图片审计结果的返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"此图片审计结果的返回信息，和code对应"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"riskLevel", "zh_CN":"审计结果，PASS ：正常，建议直接放行；REVIEW ：可疑，建议人工审核； REJECT ：违规，建议直接拦截"}
  RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty" require:"true"`
  // {"en":"riskLabel", "zh_CN":"风险标签 涉政：politics,暴恐:violence,色情:porn,违禁:ban,辱骂:abuse,广告法:ad_law,广告:ad,黑名单:blacklist,无意义:meaningless,隐私:privacy,网络诈骗:fraud,未成年人:minor"}
  RiskLabel *string `json:"riskLabel,omitempty" xml:"riskLabel,omitempty" require:"true"`
  // {"en":"riskDescription", "zh_CN":"风险原因"}
  RiskDescription *string `json:"riskDescription,omitempty" xml:"riskDescription,omitempty" require:"true"`
  // {"en":"ocrText", "zh_CN":"返回图⽚中违规⽂字相关信息，当请求参数type字段包含OCR时存在"}
  OcrText *string `json:"ocrText,omitempty" xml:"ocrText,omitempty" require:"true"`
}

func (s ImageAuditImgQueryResp) String() string {
  return tea.Prettify(s)
}

func (s ImageAuditImgQueryResp) GoString() string {
  return s.String()
}

func (s *ImageAuditImgQueryResp) SetRequestId(v string) *ImageAuditImgQueryResp {
  s.RequestId = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetBtId(v string) *ImageAuditImgQueryResp {
  s.BtId = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetCode(v string) *ImageAuditImgQueryResp {
  s.Code = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetMessage(v string) *ImageAuditImgQueryResp {
  s.Message = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetRiskLevel(v string) *ImageAuditImgQueryResp {
  s.RiskLevel = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetRiskLabel(v string) *ImageAuditImgQueryResp {
  s.RiskLabel = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetRiskDescription(v string) *ImageAuditImgQueryResp {
  s.RiskDescription = &v
  return s
}

func (s *ImageAuditImgQueryResp) SetOcrText(v string) *ImageAuditImgQueryResp {
  s.OcrText = &v
  return s
}




type TextAuditRequest struct {
  // {"en":"userName", "zh_CN":"客户登陆账号"}
  UserName *string `json:"userName,omitempty" xml:"userName,omitempty" require:"true"`
  // {"en":"lang", "zh_CN":"业务类别，当前支持： zh-中文文本审计；en-英文文本审计；ar-阿语文本审计；传值zh或en或ar"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty" require:"true"`
  // {"en":"text", "zh_CN":"要审计的文本内容"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
  // {"en":"businessType", "zh_CN":"非必填项，配置是否针对“未成年人”进行审计,若需要 则填写固定值为：MINOR"}
  BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
  // {"en":"appType", "zh_CN":"应用类型：非必填项，默认值为“default”。
  // 若客户有特殊的检测规则要求，则需要网宿单独分配提供。"}
  AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
}

func (s TextAuditRequest) String() string {
  return tea.Prettify(s)
}

func (s TextAuditRequest) GoString() string {
  return s.String()
}

func (s *TextAuditRequest) SetUserName(v string) *TextAuditRequest {
  s.UserName = &v
  return s
}

func (s *TextAuditRequest) SetLang(v string) *TextAuditRequest {
  s.Lang = &v
  return s
}

func (s *TextAuditRequest) SetText(v string) *TextAuditRequest {
  s.Text = &v
  return s
}

func (s *TextAuditRequest) SetBusinessType(v string) *TextAuditRequest {
  s.BusinessType = &v
  return s
}

func (s *TextAuditRequest) SetAppType(v string) *TextAuditRequest {
  s.AppType = &v
  return s
}

type TextAuditResponse struct {
  // {"en":"code", "zh_CN":"请求返回码，1100成功，其他是8位失败码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"错误码具体信息描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"ws生成的请求唯一id"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"riskLevel", "zh_CN":"审计结果，PASS ：正常，建议直接放行；REVIEW ：可疑，建议人工审核； REJECT ：违规，建议直接拦截"}
  RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty" require:"true"`
  // {"en":"riskLabel", "zh_CN":"一级风险标签：涉政:politics,暴恐:violence,色情:porn,违禁:ban,辱骂:abuse,广告法:ad_law,广告:ad,黑名单:blacklist,无意义:meaningless,隐私:privacy,网络诈骗:fraud,未成年人:minor"}
  RiskLabel *string `json:"riskLabel,omitempty" xml:"riskLabel,omitempty" require:"true"`
  // {"en":"riskDescription", "zh_CN":"风险原因"}
  RiskDescription *string `json:"riskDescription,omitempty" xml:"riskDescription,omitempty" require:"true"`
}

func (s TextAuditResponse) String() string {
  return tea.Prettify(s)
}

func (s TextAuditResponse) GoString() string {
  return s.String()
}

func (s *TextAuditResponse) SetCode(v string) *TextAuditResponse {
  s.Code = &v
  return s
}

func (s *TextAuditResponse) SetMessage(v string) *TextAuditResponse {
  s.Message = &v
  return s
}

func (s *TextAuditResponse) SetRequestId(v string) *TextAuditResponse {
  s.RequestId = &v
  return s
}

func (s *TextAuditResponse) SetRiskLevel(v string) *TextAuditResponse {
  s.RiskLevel = &v
  return s
}

func (s *TextAuditResponse) SetRiskLabel(v string) *TextAuditResponse {
  s.RiskLabel = &v
  return s
}

func (s *TextAuditResponse) SetRiskDescription(v string) *TextAuditResponse {
  s.RiskDescription = &v
  return s
}

type TextAuditPaths struct {
}

func (s TextAuditPaths) String() string {
  return tea.Prettify(s)
}

func (s TextAuditPaths) GoString() string {
  return s.String()
}

type TextAuditParameters struct {
}

func (s TextAuditParameters) String() string {
  return tea.Prettify(s)
}

func (s TextAuditParameters) GoString() string {
  return s.String()
}

type TextAuditRequestHeader struct {
}

func (s TextAuditRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s TextAuditRequestHeader) GoString() string {
  return s.String()
}

type TextAuditResponseHeader struct {
}

func (s TextAuditResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s TextAuditResponseHeader) GoString() string {
  return s.String()
}




