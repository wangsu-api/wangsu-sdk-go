package ipforbid

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryForbiddingVisitorIPsByLabelCodeServiceRequest struct {
  // {"en":"List of forbidding Label Code","zh_CN":"封禁标签列表"}
  LabelCodeList []*string `json:"labelCodeList,omitempty" xml:"labelCodeList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of forbidding IP, leave it empty to query all","zh_CN":"封禁IP列表,放空则查询全部"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"Current page number,the first page starts from 0,default 0","zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo `json:"pageNo,omitempty" xml:"pageNo,omitempty" type:"Struct"`
  // {"en":"Page size,must be greater than 0,default 100","zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize `json:"pageSize,omitempty" xml:"pageSize,omitempty" type:"Struct"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetLabelCodeList(v []*string) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.LabelCodeList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetIpList(v []*string) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.IpList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageNo(v *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageNo = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageSize(v *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageSize = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServicePaths struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServicePaths) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceParameters struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceParameters) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetCode(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetMessage(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetData(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseData struct {
  // {"en":"Total count","zh_CN":"总数据条数"}
  Total *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal `json:"total,omitempty" xml:"total,omitempty" require:"true" type:"Struct"`
  // {"en":"Current page number,the first page starts from 0,default 0","zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true" type:"Struct"`
  // {"en":"Page size,must be greater than 0,default 100","zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true" type:"Struct"`
  // {"en":"Query results","zh_CN":"查询结果"}
  Result []*QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetTotal(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.Total = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetPageNo(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.PageNo = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetPageSize(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.PageSize = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetResult(v []*QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.Result = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult struct     {
  // {"en":"Label Code","zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"Label Name","zh_CN":"标签名称"}
  LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty" require:"true"`
  // {"en":"IP forbidding","zh_CN":"封禁的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Start time of forbidden","zh_CN":"封禁开始时间"}
  StartTime *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true" type:"Struct"`
  // {"en":"End time of forbidden","zh_CN":"封禁结束时间"}
  EndTime *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true" type:"Struct"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetLabelCode(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.LabelCode = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetLabelName(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.LabelName = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetIp(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.Ip = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetStartTime(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.StartTime = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetEndTime(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.EndTime = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseHeader struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type ForbidOrResumeVisitorIPsByLabelCodeServiceRequest struct {
  // {"en":"List of operating objects","zh_CN":"操作对象列表"}
  OperationObjectList []*ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList `json:"operationObjectList,omitempty" xml:"operationObjectList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operation Type, 1: forbid; 2: resume","zh_CN":"操作类型， 1: 封禁； 2: 解禁"}
  OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Forbid Duration(minutes),The maximum value is 2628000 minutes(five years), and it will automatically be set to 2628000 if exceeded. Required for forbidding operation, non-required for resuming operation.","zh_CN":"封禁时长（分钟），最大值为2628000分钟（即五年），超过自动设置为2628000。封禁操作时，必填，解禁时非必填。"}
  ForbidTime *int `json:"forbidTime,omitempty" xml:"forbidTime,omitempty"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetOperationObjectList(v []*ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.OperationObjectList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetOperationType(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.OperationType = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetForbidTime(v int) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.ForbidTime = &v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList struct     {
  // {"en":"Label Code (Please contact technical support for assistance)","zh_CN":"标签编码（请联系专属技术支持获取）"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"List of IPs to forbid or resume. The maximum number of IPs is 10,000. The IP address can be v4 or v6,only supports IPV4 segment, does not support IPV6 segment.","zh_CN":"待封禁或解禁的IP列表。IP个数上限为10000个。支持IPV4、IPV6格式，仅支持IPV4段，不支持IPV6段"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) SetLabelCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList {
  s.LabelCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) SetIpList(v []*string) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList {
  s.IpList = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByLabelCodeServicePaths struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServicePaths) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceParameters struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceParameters) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code.If it shows `PartialSuccess`, please pay attention to the details of partial failures in errCode","zh_CN":"响应码，如果为“PartialSuccess”，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetMessage(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetData(v *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData struct {
  // {"en":"Error Code","zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message","zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Label Code","zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"List of Failed IPs","zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetErrCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.ErrCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetErrMessage(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.ErrMessage = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetLabelCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.LabelCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetFailedIpList(v []*string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.FailedIpList = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponseHeader struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type ForbidOrResumeVisitorIPsByDomainServiceRequest struct {
  // {"en":"List of Domains","zh_CN":"域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of IPs to forbid or resume.The IP address can be v4 or v6,only supports IPV4 segment, does not support IPV6 segment.","zh_CN":"待封禁或解禁的IP列表。支持IPV4、IPV6格式，仅支持IPV4段，不支持IPV6段"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operate Type(1: forbid; 2: resume;)","zh_CN":"操作类型(1: 封禁； 2: 解禁；)"}
  OperationType *int `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Forbid duration(minute), default 43200(means 30 days),max 2628000(means 5 years),it is recommended to specify a number,such as 30","zh_CN":"封禁时长（分钟），不传默认为43200（30天），最大支持2628000（5年），建议指定具体的数值，比如 30"}
  ForbidTime *int `json:"forbidTime,omitempty" xml:"forbidTime,omitempty"`
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetDomainList(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.DomainList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetIpList(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.IpList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetOperationType(v int) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.OperationType = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetForbidTime(v int) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.ForbidTime = &v
  return s
}

type ForbidOrResumeVisitorIPsByDomainServiceRequestHeader struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServicePaths struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServicePaths) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServiceParameters struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceParameters) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServiceResponse struct {
  // {"en":"Result Code.If it shows `36010032`,means `partial success`, please pay attention to the details of partial failures in errCode, adjust and try again","zh_CN":"响应码，如果为“36010032”，意味着部分成功，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Error Code","zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message","zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Failed Domain","zh_CN":"失败的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"List of Failed IPs","zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetCode(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetMessage(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetData(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Data = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetErrCode(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.ErrCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetErrMessage(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.ErrMessage = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetDomain(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Domain = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetFailedIpList(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.FailedIpList = v
  return s
}

type ForbidOrResumeVisitorIPsByDomainServiceResponseHeader struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryForbiddingVisitorIPsByDomainServiceRequest struct {
  // {"en":"List of forbidding domain", "zh_CN":"封禁域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of forbidding IP", "zh_CN":"封禁IP列表"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"Current page number,the first page starts from 0,default 0", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"Page size,must be greater than 0,default 100, the maximum is 1000 ", "zh_CN":"每页大小，必须大于0，默认100，最大1000"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetDomainList(v []*string) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.DomainList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetIpList(v []*string) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.IpList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetPageNo(v int) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetPageSize(v int) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.PageSize = &v
  return s
}

type QueryForbiddingVisitorIPsByDomainServiceResponse struct {
  // {"en":"Result Code", "zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Total count", "zh_CN":"总数据条数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Current page number,the first page starts from 0,default 0", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"Page size,must be greater than 0,default 100, the maximum is 1000", "zh_CN":"每页大小，必须大于0，默认100，最大1000"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Query results", "zh_CN":"查询结果"}
  Result []*string `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
  // {"en":"Domain forbidding", "zh_CN":"封禁的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"IP forbidding", "zh_CN":"封禁的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Start time of forbidden", "zh_CN":"封禁开始时间"}
  StartTime *int `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time of forbidden", "zh_CN":"封禁结束时间"}
  EndTime *int `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetCode(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetMessage(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetData(v map[string]interface{}) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Data = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetTotal(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Total = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetPageNo(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetPageSize(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.PageSize = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetResult(v []*string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Result = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetDomain(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Domain = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetIp(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Ip = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetStartTime(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.StartTime = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetEndTime(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.EndTime = &v
  return s
}

type QueryForbiddingVisitorIPsByDomainServicePaths struct {
}

func (s QueryForbiddingVisitorIPsByDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServicePaths) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByDomainServiceParameters struct {
}

func (s QueryForbiddingVisitorIPsByDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceParameters) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByDomainServiceRequestHeader struct {
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByDomainServiceResponseHeader struct {
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryForbiddingIPWhitelistServiceRequest struct {
  // {"en":"List of Domains, leave it empty to query all ", "zh_CN":"域名列表，为空时查全部 "}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"List of Label Code, leave it empty to query all ", "zh_CN":"标签列表，为空时查全部 "}
  LabelCodeList []*string `json:"labelCodeList,omitempty" xml:"labelCodeList,omitempty" type:"Repeated"`
  // {"en":"List of Whitelist IPs, leave it empty to query all ", "zh_CN":"IP列表，为空时查全部 "}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"List of Whitelist Types, 0:'Customer Granularity', 1:'Domain Granularity', 2:'Label Granularity', leave it empty to query all ", "zh_CN":"白名单类型列表，0:客户粒度，1:域名粒度，2:标签粒度，为空时查全部 "}
  WhitelistTypes []*string `json:"whitelistTypes,omitempty" xml:"whitelistTypes,omitempty" type:"Repeated"`
  // {"en":"Page size,must be greater than 0,default 100 ", "zh_CN":"每页大小，必须大于0，默认100 "}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Current page number,the first page starts from 0,default 0  ", "zh_CN":"分页，当前页，第一页从0开始，默认0 "}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
}

func (s QueryForbiddingIPWhitelistServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryForbiddingIPWhitelistServiceRequest) SetDomainList(v []*string) *QueryForbiddingIPWhitelistServiceRequest {
  s.DomainList = v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceRequest) SetLabelCodeList(v []*string) *QueryForbiddingIPWhitelistServiceRequest {
  s.LabelCodeList = v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceRequest) SetIpList(v []*string) *QueryForbiddingIPWhitelistServiceRequest {
  s.IpList = v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceRequest) SetWhitelistTypes(v []*string) *QueryForbiddingIPWhitelistServiceRequest {
  s.WhitelistTypes = v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceRequest) SetPageSize(v int) *QueryForbiddingIPWhitelistServiceRequest {
  s.PageSize = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceRequest) SetPageNo(v int) *QueryForbiddingIPWhitelistServiceRequest {
  s.PageNo = &v
  return s
}

type QueryForbiddingIPWhitelistServiceResponse struct {
  // {"en":"Result Code", "zh_CN":"响应码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result QueryForbiddingIPWhitelistServiceData", "zh_CN":"响应数据"}
  QueryForbiddingIPWhitelistServiceData *QueryForbiddingIPWhitelistServiceData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryForbiddingIPWhitelistServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryForbiddingIPWhitelistServiceResponse) SetCode(v string) *QueryForbiddingIPWhitelistServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceResponse) SetMessage(v string) *QueryForbiddingIPWhitelistServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceResponse) SetData(v *QueryForbiddingIPWhitelistServiceData) *QueryForbiddingIPWhitelistServiceResponse {
  s.QueryForbiddingIPWhitelistServiceData = v
  return s
}

type QueryForbiddingIPWhitelistServiceData struct {
  // {"en":"Total count", "zh_CN":"总数据条数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Page size,must be greater than 0,default 100", "zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Current page number,the first page starts from 0,default 0 ", "zh_CN":"分页，当前页，第一页从0开始，默认0 "}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"Query results", "zh_CN":"查询结果"}
  Results []*QueryForbiddingIPWhitelistServiceIpWhitelist `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s QueryForbiddingIPWhitelistServiceData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingIPWhitelistServiceData) SetTotal(v int) *QueryForbiddingIPWhitelistServiceData {
  s.Total = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceData) SetPageSize(v int) *QueryForbiddingIPWhitelistServiceData {
  s.PageSize = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceData) SetPageNo(v int) *QueryForbiddingIPWhitelistServiceData {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceData) SetResults(v []*QueryForbiddingIPWhitelistServiceIpWhitelist) *QueryForbiddingIPWhitelistServiceData {
  s.Results = v
  return s
}

type QueryForbiddingIPWhitelistServiceIpWhitelist struct {
  // {"en":"Whitelist Type, 0:'Customer Granularity', 1:'Domain Granularity', 2:'Label Granularity' ", "zh_CN":"白名单类型，0:客户粒度，1:域名粒度，2:标签粒度 "}
  WhitelistType *string `json:"whitelistType,omitempty" xml:"whitelistType,omitempty" require:"true"`
  // {"en":"Whitelist Object, must be CustomerCode or Domain or Label Code,It depends on the `Whitelist Type` ", "zh_CN":"白名单对象，客户编码或加速域名或标签编码，取决于白名单类型 "}
  WhitelistObject *string `json:"whitelistObject,omitempty" xml:"whitelistObject,omitempty" require:"true"`
  // {"en":"Whitelist IP", "zh_CN":"白名单IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s QueryForbiddingIPWhitelistServiceIpWhitelist) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceIpWhitelist) GoString() string {
  return s.String()
}

func (s *QueryForbiddingIPWhitelistServiceIpWhitelist) SetWhitelistType(v string) *QueryForbiddingIPWhitelistServiceIpWhitelist {
  s.WhitelistType = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceIpWhitelist) SetWhitelistObject(v string) *QueryForbiddingIPWhitelistServiceIpWhitelist {
  s.WhitelistObject = &v
  return s
}

func (s *QueryForbiddingIPWhitelistServiceIpWhitelist) SetIp(v string) *QueryForbiddingIPWhitelistServiceIpWhitelist {
  s.Ip = &v
  return s
}

type QueryForbiddingIPWhitelistServicePaths struct {
}

func (s QueryForbiddingIPWhitelistServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServicePaths) GoString() string {
  return s.String()
}

type QueryForbiddingIPWhitelistServiceParameters struct {
}

func (s QueryForbiddingIPWhitelistServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceParameters) GoString() string {
  return s.String()
}

type QueryForbiddingIPWhitelistServiceRequestHeader struct {
}

func (s QueryForbiddingIPWhitelistServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryForbiddingIPWhitelistServiceResponseHeader struct {
}

func (s QueryForbiddingIPWhitelistServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingIPWhitelistServiceResponseHeader) GoString() string {
  return s.String()
}




type AddOrRemoveForbiddingIPWhitelistServiceRequest struct {
  // {"en":"Whitelist Objects","zh_CN":"白名单操作对象列表"}
  WhitelistObjectList []*AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList `json:"whitelistObjectList,omitempty" xml:"whitelistObjectList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operation Type, 1:add, 2:remove","zh_CN":"操作类型，1:新增，2:删除"}
  OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Whitelist Type, 0:'Customer Granularity', 1:'Domain Granularity', 2:'Label Granularity'","zh_CN":"白名单类型，0:客户粒度，1:域名粒度，2:标签粒度"}
  WhitelistType *string `json:"whitelistType,omitempty" xml:"whitelistType,omitempty" require:"true"`
  // {"en":"Force resume,default 'false'. When adding, simultaneously resume IPs that are still forbidding at the corresponding granularity","zh_CN":"是否强制解禁，默认否。添加时，同时解禁对应粒度下还处于封禁中的IP。"}
  ForceResume *bool `json:"forceResume,omitempty" xml:"forceResume,omitempty"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequest) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequest) SetWhitelistObjectList(v []*AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList) *AddOrRemoveForbiddingIPWhitelistServiceRequest {
  s.WhitelistObjectList = v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequest) SetOperationType(v string) *AddOrRemoveForbiddingIPWhitelistServiceRequest {
  s.OperationType = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequest) SetWhitelistType(v string) *AddOrRemoveForbiddingIPWhitelistServiceRequest {
  s.WhitelistType = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequest) SetForceResume(v bool) *AddOrRemoveForbiddingIPWhitelistServiceRequest {
  s.ForceResume = &v
  return s
}

type AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList struct     {
  // {"en":"Whitelist Object,must be Domain or Label Code.When the whitelist type is `customer granularity`, this field must be empty","zh_CN":"白名单操作对象，白名单类型“域名粒度”时，请填写加速域名，白名单类型“标签粒度”时，请填写标签编码，白名单类型为“客户粒度”时，此值必须为空"}
  WhitelistObject *string `json:"whitelistObject,omitempty" xml:"whitelistObject,omitempty"`
  // {"en":"List of Whitelist IPs","zh_CN":"白名单IP列表"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList) SetWhitelistObject(v string) *AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList {
  s.WhitelistObject = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList) SetIpList(v []*string) *AddOrRemoveForbiddingIPWhitelistServiceRequestWhitelistObjectList {
  s.IpList = v
  return s
}

type AddOrRemoveForbiddingIPWhitelistServiceRequestHeader struct {
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequestHeader) GoString() string {
  return s.String()
}

type AddOrRemoveForbiddingIPWhitelistServicePaths struct {
}

func (s AddOrRemoveForbiddingIPWhitelistServicePaths) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServicePaths) GoString() string {
  return s.String()
}

type AddOrRemoveForbiddingIPWhitelistServiceParameters struct {
}

func (s AddOrRemoveForbiddingIPWhitelistServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceParameters) GoString() string {
  return s.String()
}

type AddOrRemoveForbiddingIPWhitelistServiceResponse struct {
  // {"en":"Result Code.If it shows `36010032`,means `partial success`, please pay attention to the details of partial failures in errCode.","zh_CN":"响应码，成功为0，如果为“36010032”部分成功，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data []*AddOrRemoveForbiddingIPWhitelistServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponse) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponse) SetCode(v string) *AddOrRemoveForbiddingIPWhitelistServiceResponse {
  s.Code = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponse) SetMessage(v string) *AddOrRemoveForbiddingIPWhitelistServiceResponse {
  s.Message = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponse) SetData(v []*AddOrRemoveForbiddingIPWhitelistServiceResponseData) *AddOrRemoveForbiddingIPWhitelistServiceResponse {
  s.Data = v
  return s
}

type AddOrRemoveForbiddingIPWhitelistServiceResponseData struct     {
  // {"en":"Error Code","zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message","zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Whitelist Object","zh_CN":"白名单操作对象"}
  WhitelistObject *string `json:"whitelistObject,omitempty" xml:"whitelistObject,omitempty" require:"true"`
  // {"en":"Whitelist Type, 0:'Customer Granularity', 1:'Domain Granularity', 2:'Label Granularity'","zh_CN":"白名单类型，0:客户粒度，1:域名粒度，2:标签粒度"}
  WhitelistType *string `json:"whitelistType,omitempty" xml:"whitelistType,omitempty" require:"true"`
  // {"en":"List of Failed IPs","zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponseData) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponseData) SetErrCode(v string) *AddOrRemoveForbiddingIPWhitelistServiceResponseData {
  s.ErrCode = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponseData) SetErrMessage(v string) *AddOrRemoveForbiddingIPWhitelistServiceResponseData {
  s.ErrMessage = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponseData) SetWhitelistObject(v string) *AddOrRemoveForbiddingIPWhitelistServiceResponseData {
  s.WhitelistObject = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponseData) SetWhitelistType(v string) *AddOrRemoveForbiddingIPWhitelistServiceResponseData {
  s.WhitelistType = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponseData) SetFailedIpList(v []*string) *AddOrRemoveForbiddingIPWhitelistServiceResponseData {
  s.FailedIpList = v
  return s
}

type AddOrRemoveForbiddingIPWhitelistServiceResponseHeader struct {
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponseHeader) GoString() string {
  return s.String()
}




