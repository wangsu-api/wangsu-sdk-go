package dnsrecordmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type AddRecordRequest struct {
  // {"en":"The domain whose
  // DNS records to be
  // added", "zh_CN":"需要添加解析记录的域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Host name records", "zh_CN":"主机记录
  // CloudDNS支持的记录类型：A、AAAA、CNAME、TXT、SRV和MX"}
  DcName *string `json:"dcName,omitempty" xml:"dcName,omitempty" require:"true"`
  // {"en":"Record type.
  // 
  // Record types supported by CloudDNS: A, AAAA, CNAME, TXT, SRV, MX", "zh_CN":"记录类型
  // CloudDNS支持的记录类型：A、AAAA、CNAME、TXT、SRV、MX"}
  DcType *string `json:"dcType,omitempty" xml:"dcType,omitempty" require:"true"`
  // {"en":"Line
  // Please refer to the
  // Appendix for pair
  // details of lines", "zh_CN":"线路，线路代码对应表请参考附录"}
  DcView *string `json:"dcView,omitempty" xml:"dcView,omitempty" require:"true"`
  // {"en":"Record value
  // (Special formats of
  // SRV type: priority,
  // space, weight, port
  // number, space, target
  // address)", "zh_CN":"记录值(SRV类型特殊格式：优先级、空格、权重、空格、端口号、空格、目标地址）"}
  DcValue *string `json:"dcValue,omitempty" xml:"dcValue,omitempty" require:"true"`
  // {"en":"MX priority
  // This parameter is
  // required if 'MX' is
  // selected for 'Record
  // type'.
  // Value range is 1-50;
  // the default value is 5.
  // The smaller a number
  // is, the higher priority it
  // has.", "zh_CN":"MX优先级：
  // 如果“记录类型”选择“MX”，则需配置该参数
  // 取值范围为1~50，默认为5。数值越小，则优先级越高"}
  MxPri *int `json:"mxPri,omitempty" xml:"mxPri,omitempty"`
  // {"en":"The survival time for
  // cache. The default", "zh_CN":"指缓存的生存时间。默认可配置为600s。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"Return Chinese results for null (default)
  // En: Return the English prompt result", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s AddRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s AddRecordRequest) GoString() string {
  return s.String()
}

func (s *AddRecordRequest) SetDomainName(v string) *AddRecordRequest {
  s.DomainName = &v
  return s
}

func (s *AddRecordRequest) SetDcName(v string) *AddRecordRequest {
  s.DcName = &v
  return s
}

func (s *AddRecordRequest) SetDcType(v string) *AddRecordRequest {
  s.DcType = &v
  return s
}

func (s *AddRecordRequest) SetDcView(v string) *AddRecordRequest {
  s.DcView = &v
  return s
}

func (s *AddRecordRequest) SetDcValue(v string) *AddRecordRequest {
  s.DcValue = &v
  return s
}

func (s *AddRecordRequest) SetMxPri(v int) *AddRecordRequest {
  s.MxPri = &v
  return s
}

func (s *AddRecordRequest) SetTtl(v int) *AddRecordRequest {
  s.Ttl = &v
  return s
}

func (s *AddRecordRequest) SetLanguage(v string) *AddRecordRequest {
  s.Language = &v
  return s
}

type AddRecordResponse struct {
  // {"en":"Status code. For detailed description of resCode, please refer to 'Status Codes of Dispatch Business'.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"recordId, ID of host name record, used to identify this record.", "zh_CN":"主机记录的ID，用于标识该记录。"}
  RecordId []*string `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true" type:"Repeated"`
}

func (s AddRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s AddRecordResponse) GoString() string {
  return s.String()
}

func (s *AddRecordResponse) SetResCode(v string) *AddRecordResponse {
  s.ResCode = &v
  return s
}

func (s *AddRecordResponse) SetMsg(v string) *AddRecordResponse {
  s.Msg = &v
  return s
}

func (s *AddRecordResponse) SetRecordId(v []*string) *AddRecordResponse {
  s.RecordId = v
  return s
}

type AddRecordPaths struct {
}

func (s AddRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s AddRecordPaths) GoString() string {
  return s.String()
}

type AddRecordParameters struct {
}

func (s AddRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s AddRecordParameters) GoString() string {
  return s.String()
}

type AddRecordRequestHeader struct {
}

func (s AddRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddRecordRequestHeader) GoString() string {
  return s.String()
}

type AddRecordResponseHeader struct {
}

func (s AddRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddRecordResponseHeader) GoString() string {
  return s.String()
}




type QueryRecordsRequest struct {
  // {"en":"The domain whose DNS records to be viewed", "zh_CN":"需要查询解析记录的域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Fuzzy query(no need to enter this parameter if accurate query is used). Used this parameter to fuzzy query 'Host name records', 'Record type', 'Record value'.", "zh_CN":"模糊查询 (如果需要精确查询，则这个参数不填)。根据此查询参数对“主机记录”，“记录类型”，“记录值”进行模糊查询"}
  Param *string `json:"param,omitempty" xml:"param,omitempty"`
  // {"en":"Host name records(accurate query)", "zh_CN":"主机记录 （精确查询）"}
  DcName *string `json:"dcName,omitempty" xml:"dcName,omitempty"`
  // {"en":"Record type(accurate query) CloudDNS supports the following record types: A, AAAA, CNAME and MX", "zh_CN":"记录类型（精确查询）
  // CloudDNS支持的记录类型：A、AAAA、CNAME、TXT和MX"}
  DcType *string `json:"dcType,omitempty" xml:"dcType,omitempty"`
  // {"en":"Chinese name of the line(accurate query), for example: China Telecom", "zh_CN":"线路中文名称（精确查询）例如：中国电信"}
  DcView *string `json:"dcView,omitempty" xml:"dcView,omitempty"`
  // {"en":"Line ID (accurate query)", "zh_CN":"线路ID(精确查询)"}
  DcViewId *string `json:"dcViewId,omitempty" xml:"dcViewId,omitempty"`
  // {"en":"Record value(accurate query)", "zh_CN":"记录值（精确查询）"}
  DcValue *string `json:"dcValue,omitempty" xml:"dcValue,omitempty"`
  // {"en":"Status of DNS record. (Accurate query)Normal: 2; Suspend:1; All:0", "zh_CN":"解析记录状态。（精确查询）正常：2 停用：1 全部：0"}
  State *int `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Status of DNS record. (Accurate query)Normal: 2; Suspend:1; All:0", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s QueryRecordsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsRequest) GoString() string {
  return s.String()
}

func (s *QueryRecordsRequest) SetDomainName(v string) *QueryRecordsRequest {
  s.DomainName = &v
  return s
}

func (s *QueryRecordsRequest) SetParam(v string) *QueryRecordsRequest {
  s.Param = &v
  return s
}

func (s *QueryRecordsRequest) SetDcName(v string) *QueryRecordsRequest {
  s.DcName = &v
  return s
}

func (s *QueryRecordsRequest) SetDcType(v string) *QueryRecordsRequest {
  s.DcType = &v
  return s
}

func (s *QueryRecordsRequest) SetDcView(v string) *QueryRecordsRequest {
  s.DcView = &v
  return s
}

func (s *QueryRecordsRequest) SetDcViewId(v string) *QueryRecordsRequest {
  s.DcViewId = &v
  return s
}

func (s *QueryRecordsRequest) SetDcValue(v string) *QueryRecordsRequest {
  s.DcValue = &v
  return s
}

func (s *QueryRecordsRequest) SetState(v int) *QueryRecordsRequest {
  s.State = &v
  return s
}

func (s *QueryRecordsRequest) SetLanguage(v string) *QueryRecordsRequest {
  s.Language = &v
  return s
}

type QueryRecordsResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the
  // status code.", "zh_CN":"状态码详细说明"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  Content []*QueryRecordsQueryRecordsResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRecordsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsResponse) GoString() string {
  return s.String()
}

func (s *QueryRecordsResponse) SetResCode(v int) *QueryRecordsResponse {
  s.ResCode = &v
  return s
}

func (s *QueryRecordsResponse) SetMsg(v string) *QueryRecordsResponse {
  s.Msg = &v
  return s
}

func (s *QueryRecordsResponse) SetContent(v []*QueryRecordsQueryRecordsResponseContent) *QueryRecordsResponse {
  s.Content = v
  return s
}

type QueryRecordsQueryRecordsResponseContent struct     {
  // {"en":"ID of host name record", "zh_CN":"主机记录ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"Host name records", "zh_CN":"主机记录"}
  DcName *string `json:"dcName,omitempty" xml:"dcName,omitempty" require:"true"`
  // {"en":"Record type
  // CloudDNS supports the
  // following record types: A,
  // AAAA, CNAME and MX", "zh_CN":"记录类型
  // CloudDNS支持的记录类型：A、AAAA、CNAME、TXT和MX"}
  DcType *string `json:"dcType,omitempty" xml:"dcType,omitempty" require:"true"`
  // {"en":"Line  Please refer to the Appendix
  // for pair details of lines", "zh_CN":"线路，线路代码对应表请参考附录"}
  DcView *int `json:"dcView,omitempty" xml:"dcView,omitempty" require:"true"`
  // {"en":"Record value", "zh_CN":"记录值"}
  DcValue *string `json:"dcValue,omitempty" xml:"dcValue,omitempty" require:"true"`
  // {"en":"Status of DNS record. Normal:
  // 2; Suspend:1", "zh_CN":"解析记录状态。正常：2 停用：1"}
  State *int `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"MX priority
  // This parameter is required if
  // 'MX' is selected for 'Record
  // type'.
  // Value range is 1-50; the
  // default value is 5. The smaller
  // a number is, the higher priority
  // it has.", "zh_CN":"MX优先级，注意点：
  // 1.如果“记录类型”选择“MX”，则需配置该参数
  // 2.取值范围为1~50，默认为5。数值越小，则优先级越高"}
  MxPri *int `json:"mxPri,omitempty" xml:"mxPri,omitempty" require:"true"`
  // {"en":"The survival time for cache.
  // The default configured value is
  // 600s.", "zh_CN":"指缓存的生存时间。默认可配置为600s。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"Operation permissions: 1, read-only 2, read and write", "zh_CN":"操作权限：1、只读 2、读写"}
  Auth *int `json:"auth,omitempty" xml:"auth,omitempty" require:"true"`
}

func (s QueryRecordsQueryRecordsResponseContent) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsQueryRecordsResponseContent) GoString() string {
  return s.String()
}

func (s *QueryRecordsQueryRecordsResponseContent) SetRecordId(v int) *QueryRecordsQueryRecordsResponseContent {
  s.RecordId = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetDcName(v string) *QueryRecordsQueryRecordsResponseContent {
  s.DcName = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetDcType(v string) *QueryRecordsQueryRecordsResponseContent {
  s.DcType = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetDcView(v int) *QueryRecordsQueryRecordsResponseContent {
  s.DcView = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetDcValue(v string) *QueryRecordsQueryRecordsResponseContent {
  s.DcValue = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetState(v int) *QueryRecordsQueryRecordsResponseContent {
  s.State = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetMxPri(v int) *QueryRecordsQueryRecordsResponseContent {
  s.MxPri = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetTtl(v int) *QueryRecordsQueryRecordsResponseContent {
  s.Ttl = &v
  return s
}

func (s *QueryRecordsQueryRecordsResponseContent) SetAuth(v int) *QueryRecordsQueryRecordsResponseContent {
  s.Auth = &v
  return s
}

type QueryRecordsPaths struct {
}

func (s QueryRecordsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsPaths) GoString() string {
  return s.String()
}

type QueryRecordsParameters struct {
}

func (s QueryRecordsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsParameters) GoString() string {
  return s.String()
}

type QueryRecordsRequestHeader struct {
}

func (s QueryRecordsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsRequestHeader) GoString() string {
  return s.String()
}

type QueryRecordsResponseHeader struct {
}

func (s QueryRecordsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRecordsResponseHeader) GoString() string {
  return s.String()
}




type ControlRecordRequest struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"ID of host name
  // record, use ; to
  // separate two", "zh_CN":"主机记录ID,多个用英文;分隔"}
  RecordIds *string `json:"recordIds,omitempty" xml:"recordIds,omitempty" require:"true"`
  // {"en":"Operation: 1 Disable; 2
  // Enable", "zh_CN":"操作：1停用；2启用"}
  Operate *int `json:"operate,omitempty" xml:"operate,omitempty" require:"true"`
  // {"en":"Operation: 1 Disable; 2
  // Enable", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ControlRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s ControlRecordRequest) GoString() string {
  return s.String()
}

func (s *ControlRecordRequest) SetDomainName(v string) *ControlRecordRequest {
  s.DomainName = &v
  return s
}

func (s *ControlRecordRequest) SetRecordIds(v string) *ControlRecordRequest {
  s.RecordIds = &v
  return s
}

func (s *ControlRecordRequest) SetOperate(v int) *ControlRecordRequest {
  s.Operate = &v
  return s
}

func (s *ControlRecordRequest) SetLanguage(v string) *ControlRecordRequest {
  s.Language = &v
  return s
}

type ControlRecordResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *int `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the
  // status code.", "zh_CN":"状态码详细说明"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"recordId, ID of host name
  // record
  // code msg of status code,
  // detailed description of the
  // status code.", "zh_CN":"recordId 主机记录ID
  // code 状态码
  // msg 状态码详细说明"}
  Content []*string `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s ControlRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s ControlRecordResponse) GoString() string {
  return s.String()
}

func (s *ControlRecordResponse) SetResCode(v int) *ControlRecordResponse {
  s.ResCode = &v
  return s
}

func (s *ControlRecordResponse) SetMsg(v string) *ControlRecordResponse {
  s.Msg = &v
  return s
}

func (s *ControlRecordResponse) SetContent(v []*string) *ControlRecordResponse {
  s.Content = v
  return s
}

type ControlRecordPaths struct {
}

func (s ControlRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s ControlRecordPaths) GoString() string {
  return s.String()
}

type ControlRecordParameters struct {
}

func (s ControlRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s ControlRecordParameters) GoString() string {
  return s.String()
}

type ControlRecordRequestHeader struct {
}

func (s ControlRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlRecordRequestHeader) GoString() string {
  return s.String()
}

type ControlRecordResponseHeader struct {
}

func (s ControlRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ControlRecordResponseHeader) GoString() string {
  return s.String()
}




type ConfigShareOperRequest struct {
  // {"en":"operate: 1-add", "zh_CN":"操作：1 新增"}
  Operate *int `json:"operate,omitempty" xml:"operate,omitempty" require:"true"`
  // {"en":"share object", "zh_CN":"授权对象"}
  ConfigShare []*string `json:"configShare,omitempty" xml:"configShare,omitempty" require:"true" type:"Repeated"`
}

func (s ConfigShareOperRequest) String() string {
  return tea.Prettify(s)
}

func (s ConfigShareOperRequest) GoString() string {
  return s.String()
}

func (s *ConfigShareOperRequest) SetOperate(v int) *ConfigShareOperRequest {
  s.Operate = &v
  return s
}

func (s *ConfigShareOperRequest) SetConfigShare(v []*string) *ConfigShareOperRequest {
  s.ConfigShare = v
  return s
}

type ConfigShareOperResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"detail", "zh_CN":"详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ConfigShareOperResponse) String() string {
  return tea.Prettify(s)
}

func (s ConfigShareOperResponse) GoString() string {
  return s.String()
}

func (s *ConfigShareOperResponse) SetResCode(v string) *ConfigShareOperResponse {
  s.ResCode = &v
  return s
}

func (s *ConfigShareOperResponse) SetMsg(v string) *ConfigShareOperResponse {
  s.Msg = &v
  return s
}

type ConfigShareOperPaths struct {
}

func (s ConfigShareOperPaths) String() string {
  return tea.Prettify(s)
}

func (s ConfigShareOperPaths) GoString() string {
  return s.String()
}

type ConfigShareOperParameters struct {
}

func (s ConfigShareOperParameters) String() string {
  return tea.Prettify(s)
}

func (s ConfigShareOperParameters) GoString() string {
  return s.String()
}

type ConfigShareOperRequestHeader struct {
}

func (s ConfigShareOperRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ConfigShareOperRequestHeader) GoString() string {
  return s.String()
}

type ConfigShareOperResponseHeader struct {
}

func (s ConfigShareOperResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ConfigShareOperResponseHeader) GoString() string {
  return s.String()
}




type ModifyRecordRequest struct {
  // {"en":"ID of host name record", "zh_CN":"主机记录ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"The domain whose DNS records to be added", "zh_CN":"需要添加解析记录的域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Host name records", "zh_CN":"主机记录"}
  DcName *string `json:"dcName,omitempty" xml:"dcName,omitempty" require:"true"`
  // {"en":"Record type CloudDNS supports the following record types: A, AAAA, CNAME, TXT, SRV and MX", "zh_CN":"主机记录
  // CloudDNS支持的记录类型：A、AAAA、CNAME、TXT、SRV和MX"}
  DcType *string `json:"dcType,omitempty" xml:"dcType,omitempty" require:"true"`
  // {"en":"Line Please refer to the Appendix for pair details of lines", "zh_CN":"线路，线路代码对应表请参考附录"}
  DcView *int `json:"dcView,omitempty" xml:"dcView,omitempty" require:"true"`
  // {"en":"Record value (Special formats of SRV type: priority, space, weight, port number, space, target address)", "zh_CN":"记录值(SRV类型特殊格式：优先级、空格、权重、空格、端口号、空格、目标地址）"}
  DcValue *string `json:"dcValue,omitempty" xml:"dcValue,omitempty" require:"true"`
  // {"en":"MX priority This parameter is required if 'MX' is selected for 'Record type'. Value range is 1-50; the default value is 5. The smaller a number is, the higher priority it has.", "zh_CN":"MX优先级：
  // 如果“记录类型”选择“MX”，则需配置该参数
  // 取值范围为1~50，默认为5。数值越小，则优先级越高"}
  MxPri *int `json:"mxPri,omitempty" xml:"mxPri,omitempty"`
  // {"en":"The survival time for cache. The default configured value is 600s.", "zh_CN":"指缓存的生存时间。默认可配置为600s。"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"The survival time for cache. The default configured value is 600s.", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ModifyRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordRequest) GoString() string {
  return s.String()
}

func (s *ModifyRecordRequest) SetRecordId(v int) *ModifyRecordRequest {
  s.RecordId = &v
  return s
}

func (s *ModifyRecordRequest) SetDomainName(v string) *ModifyRecordRequest {
  s.DomainName = &v
  return s
}

func (s *ModifyRecordRequest) SetDcName(v string) *ModifyRecordRequest {
  s.DcName = &v
  return s
}

func (s *ModifyRecordRequest) SetDcType(v string) *ModifyRecordRequest {
  s.DcType = &v
  return s
}

func (s *ModifyRecordRequest) SetDcView(v int) *ModifyRecordRequest {
  s.DcView = &v
  return s
}

func (s *ModifyRecordRequest) SetDcValue(v string) *ModifyRecordRequest {
  s.DcValue = &v
  return s
}

func (s *ModifyRecordRequest) SetMxPri(v int) *ModifyRecordRequest {
  s.MxPri = &v
  return s
}

func (s *ModifyRecordRequest) SetTtl(v int) *ModifyRecordRequest {
  s.Ttl = &v
  return s
}

func (s *ModifyRecordRequest) SetLanguage(v string) *ModifyRecordRequest {
  s.Language = &v
  return s
}

type ModifyRecordResponse struct {
  // {"en":"Status code. For detailed description of resCode, please refer to "Status Codes of Dispatch Business".", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"recordId, ID of host name record, used to identify this record.", "zh_CN":"recordId 主机记录的ID，用于标识该记录。"}
  Content map[string]interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s ModifyRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordResponse) GoString() string {
  return s.String()
}

func (s *ModifyRecordResponse) SetResCode(v string) *ModifyRecordResponse {
  s.ResCode = &v
  return s
}

func (s *ModifyRecordResponse) SetMsg(v string) *ModifyRecordResponse {
  s.Msg = &v
  return s
}

func (s *ModifyRecordResponse) SetContent(v map[string]interface{}) *ModifyRecordResponse {
  s.Content = v
  return s
}

type ModifyRecordPaths struct {
}

func (s ModifyRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordPaths) GoString() string {
  return s.String()
}

type ModifyRecordParameters struct {
}

func (s ModifyRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordParameters) GoString() string {
  return s.String()
}

type ModifyRecordRequestHeader struct {
}

func (s ModifyRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordRequestHeader) GoString() string {
  return s.String()
}

type ModifyRecordResponseHeader struct {
}

func (s ModifyRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ModifyRecordResponseHeader) GoString() string {
  return s.String()
}




type DelRecordRequest struct {
  // {"en":"ID of host name record", "zh_CN":"主机记录ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"ID of host name record", "zh_CN":"为空返回中文结果(默认)
  // en:返回英文提示结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s DelRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s DelRecordRequest) GoString() string {
  return s.String()
}

func (s *DelRecordRequest) SetRecordId(v int) *DelRecordRequest {
  s.RecordId = &v
  return s
}

func (s *DelRecordRequest) SetLanguage(v string) *DelRecordRequest {
  s.Language = &v
  return s
}

type DelRecordResponse struct {
  // {"en":"Status code. For detailed
  // description of resCode, please
  // refer to 'Status Codes of
  // Dispatch Business'.", "zh_CN":"状态码。resCode的详细说明请参见“调度业务状态码”。"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
  // {"en":"Detailed description of the
  // status code.", "zh_CN":"状态码的详细说明。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DelRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s DelRecordResponse) GoString() string {
  return s.String()
}

func (s *DelRecordResponse) SetResCode(v string) *DelRecordResponse {
  s.ResCode = &v
  return s
}

func (s *DelRecordResponse) SetMsg(v string) *DelRecordResponse {
  s.Msg = &v
  return s
}

type DelRecordPaths struct {
}

func (s DelRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s DelRecordPaths) GoString() string {
  return s.String()
}

type DelRecordParameters struct {
}

func (s DelRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s DelRecordParameters) GoString() string {
  return s.String()
}

type DelRecordRequestHeader struct {
}

func (s DelRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DelRecordRequestHeader) GoString() string {
  return s.String()
}

type DelRecordResponseHeader struct {
}

func (s DelRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DelRecordResponseHeader) GoString() string {
  return s.String()
}




