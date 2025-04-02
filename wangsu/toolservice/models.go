package toolservice

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type IcpQueryServiceRequest struct {
}

func (s IcpQueryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceRequest) GoString() string {
  return s.String()
}

type IcpQueryServiceResponse struct {
  // {'en':'domainIcpData', 'zh_CN':'域名备案信息'}
  DomainIcpDataList []*IcpQueryServiceResponseDomainIcpDataList `json:"domain-icp-data,omitempty" xml:"domain-icp-data,omitempty" require:"true" type:"Repeated"`
}

func (s IcpQueryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponse) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceResponse) SetDomainIcpDataList(v []*IcpQueryServiceResponseDomainIcpDataList) *IcpQueryServiceResponse {
  s.DomainIcpDataList = v
  return s
}

type IcpQueryServiceResponseDomainIcpDataList struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Registration No.", "zh_CN":"备案号"}
  IcpNumber *string `json:"icp-number,omitempty" xml:"icp-number,omitempty" require:"true"`
}

func (s IcpQueryServiceResponseDomainIcpDataList) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponseDomainIcpDataList) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceResponseDomainIcpDataList) SetDomain(v string) *IcpQueryServiceResponseDomainIcpDataList {
  s.Domain = &v
  return s
}

func (s *IcpQueryServiceResponseDomainIcpDataList) SetIcpNumber(v string) *IcpQueryServiceResponseDomainIcpDataList {
  s.IcpNumber = &v
  return s
}

type IcpQueryServicePaths struct {
}

func (s IcpQueryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServicePaths) GoString() string {
  return s.String()
}

type IcpQueryServiceParameters struct {
  // {"en":"Domain names, multiple domain names are separated by English semicolons. The maximum number of domain names is 20.", "zh_CN":"域名，多个以英文分号分隔。域名数上限为20个。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s IcpQueryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceParameters) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceParameters) SetDomain(v string) *IcpQueryServiceParameters {
  s.Domain = &v
  return s
}

type IcpQueryServiceRequestHeader struct {
}

func (s IcpQueryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceRequestHeader) GoString() string {
  return s.String()
}

type IcpQueryServiceResponseHeader struct {
}

func (s IcpQueryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthLimitTaskListServiceRequest struct {
}

func (s QueryBandwidthLimitTaskListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceRequest) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*QueryBandwidthLimitTaskListServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthLimitTaskListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthLimitTaskListServiceResponse) SetResult(v []*QueryBandwidthLimitTaskListServiceResponseResult) *QueryBandwidthLimitTaskListServiceResponse {
  s.Result = v
  return s
}

type QueryBandwidthLimitTaskListServiceResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {'en':'Maximum bandwidth set', 'zh_CN':'设置的最大带宽值'}
  BandwidthLimit *int32 `json:"bandwidthLimit,omitempty" xml:"bandwidthLimit,omitempty" require:"true"`
  // {'en':'Task name', 'zh_CN':'任务名称'}
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty" require:"true"`
}

func (s QueryBandwidthLimitTaskListServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetDomainName(v string) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.DomainName = &v
  return s
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetBandwidthLimit(v int32) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.BandwidthLimit = &v
  return s
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetTaskName(v string) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.TaskName = &v
  return s
}

type QueryBandwidthLimitTaskListServicePaths struct {
}

func (s QueryBandwidthLimitTaskListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServicePaths) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceParameters struct {
}

func (s QueryBandwidthLimitTaskListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceParameters) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceRequestHeader struct {
}

func (s QueryBandwidthLimitTaskListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceResponseHeader struct {
}

func (s QueryBandwidthLimitTaskListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryForbiddingVisitorIPsByLabelCodeServiceRequest struct {
  // {"en":"List of forbidding Label Code", "zh_CN":"封禁标签列表"}
  LabelCodeList []*string `json:"labelCodeList,omitempty" xml:"labelCodeList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of forbidding IP, leave it empty to query all ", "zh_CN":"封禁IP列表,放空则查询全部"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"Current page number,the first page starts from 0,default 0 ", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"Page size,must be greater than 0,default 100 ", "zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageNo(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageSize(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageSize = &v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code", "zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetData(v *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData struct {
  // {"en":"Total count ", "zh_CN":"总数据条数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Current page number,the first page starts from 0,default 0 ", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"Page size,must be greater than 0,default 100 ", "zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Query results", "zh_CN":"查询结果"}
  Result []*QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetTotal(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.Total = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetPageNo(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetPageSize(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetResult(v []*QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.Result = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData struct {
  // {"en":"Label Code", "zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"Label Name", "zh_CN":"标签名称"}
  LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty" require:"true"`
  // {"en":"IP forbidding", "zh_CN":"封禁的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Start time of forbidden", "zh_CN":"封禁开始时间"}
  StartTime *int `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time of forbidden", "zh_CN":"封禁结束时间"}
  EndTime *int `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetLabelCode(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.LabelCode = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetLabelName(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.LabelName = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetIp(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.Ip = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetStartTime(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.StartTime = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetEndTime(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.EndTime = &v
  return s
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

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader) GoString() string {
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
  // {"en":"List of operating objects ", "zh_CN":"操作对象列表"}
  OperationObjectList []*ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject `json:"operationObjectList,omitempty" xml:"operationObjectList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operation Type, 1: forbid; 2: resume ", "zh_CN":"操作类型， 1: 封禁； 2: 解禁 "}
  OperationType *int `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Forbid Duration(minutes),The maximum value is 2628000 minutes(five years), and it will automatically be set to 2628000 if exceeded. Required for forbidding operation, non-required for resuming operation.", "zh_CN":"封禁时长（分钟），最大值为2628000分钟（即五年），超过自动设置为2628000。封禁操作时，必填，解禁时非必填。"}
  ForbidTime *int `json:"forbidTime,omitempty" xml:"forbidTime,omitempty"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetOperationObjectList(v []*ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.OperationObjectList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetOperationType(v int) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.OperationType = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetForbidTime(v int) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.ForbidTime = &v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject struct {
  // {"en":"Label Code (Please contact technical support for assistance)", "zh_CN":"标签编码（请联系专属技术支持获取）"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"List of IPs to forbid or resume. The maximum number of IPs is 10,000. The IP address can be v4 or v6,only supports IPV4 segment, does not support IPV6 segment.", "zh_CN":"待封禁或解禁的IP列表。IP个数上限为10000个。支持IPV4、IPV6格式，仅支持IPV4段，不支持IPV6段"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject) SetLabelCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject {
  s.LabelCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject) SetIpList(v []*string) *ForbidOrResumeVisitorIPsByLabelCodeServiceOperationObject {
  s.IpList = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code.If it shows `PartialSuccess`, please pay attention to the details of partial failures in errCode ", "zh_CN":"响应码，如果为“PartialSuccess”，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetData(v *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData struct {
  // {"en":" Error Code", "zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message", "zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Label Code", "zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"List of Failed IPs", "zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) SetErrCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData {
  s.ErrCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) SetErrMessage(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData {
  s.ErrMessage = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) SetLabelCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData {
  s.LabelCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData) SetFailedIpList(v []*string) *ForbidOrResumeVisitorIPsByLabelCodeServiceForbidIpOperateResponseData {
  s.FailedIpList = v
  return s
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

type ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader) GoString() string {
  return s.String()
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
  // {"en":"List of Domains", "zh_CN":"域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of IPs to forbid or resume.The IP address can be v4 or v6,only supports IPV4 segment, does not support IPV6 segment.", "zh_CN":"待封禁或解禁的IP列表。支持IPV4、IPV6格式，仅支持IPV4段，不支持IPV6段"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operate Type(1: forbid; 2: resume;)", "zh_CN":"操作类型(1: 封禁； 2: 解禁；)"}
  OperationType *int `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Forbid duration(minute), default 43200(means 30 days),max 2628000(means 5 years),it is recommended to specify a number,such as 30", "zh_CN":"封禁时长（分钟），不传默认为43200（30天），最大支持2628000（5年），建议指定具体的数值，比如 30"}
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

type ForbidOrResumeVisitorIPsByDomainServiceResponse struct {
  // {"en":"Result Code.If it shows `36010032`,means `partial success`, please pay attention to the details of partial failures in errCode, adjust and try again", "zh_CN":"响应码，如果为“36010032”，意味着部分成功，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Error Code", "zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message", "zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Failed Domain", "zh_CN":"失败的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"List of Failed IPs", "zh_CN":"失败的IP列表"}
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

type ForbidOrResumeVisitorIPsByDomainServiceRequestHeader struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServiceResponseHeader struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportServerIpCountryCodeServiceRequest struct {
  // {'en':'Domain name:
  // 1.The maximum number of transferable domain names is 50 ;
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)', 'zh_CN':'域名：
  // 1.可传递域名数量上限为50个；
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'countryCode
  // 
  // 1Up to 10 States consulted on a one-time basis;
  // 2.Transitive values:
  // 
  // ae:Arab Emirates
  // 
  // ar:Argentina
  // 
  // at:Austria
  // 
  // au:Australia
  // 
  // bd:Bangladesh
  // 
  // be:Belgium
  // 
  // bnBrunei
  // 
  // br:Brazil
  // 
  // by:Belarus
  // 
  // ca:Canada
  // 
  // ch:Switzerland
  // 
  // cl:Chile
  // 
  // cn:China
  // 
  // co:Colombia
  // 
  // cz:Czech-Republic
  // 
  // de:Germany
  // 
  // dk:Denmark
  // 
  // dz:Algeria
  // 
  // eg:Egypt
  // 
  // es:Spain
  // 
  // fi:Finland
  // 
  // fr:France
  // 
  // gb:UK
  // 
  // gr:Greece
  // 
  // hk:Hong Kong
  // 
  // hu:Hungary
  // 
  // id:Indonesia
  // 
  // ie:Ireland
  // 
  // il:Israel
  // 
  // in:India
  // 
  // iq:Iraq
  // 
  // ir:Iran
  // 
  // it:Italy
  // 
  // jp:Japan
  // 
  // ke:Kenya
  // 
  // kg:Kyrgyzstan
  // 
  // kh:Cambodia
  // 
  // kr:South Korea
  // 
  // kw:Kuwait
  // 
  // kz:Kazakhstan
  // 
  // la:Laos
  // 
  // lt:Lithuania
  // 
  // ma:Morocco
  // 
  // mm:Myanmar
  // 
  // mn:Mongolia
  // 
  // mo:Macau
  // 
  // mx:Mexico
  // 
  // my:Malaysia
  // 
  // ng:Nigeria
  // 
  // nl:Netherlands
  // 
  // no:Norway
  // 
  // np:Nepal
  // 
  // nz:New Zealand
  // 
  // pe:Peru
  // 
  // ph:Philippines
  // 
  // pk:Pakistan
  // 
  // pl:Poland
  // 
  // pt:Portugal
  // 
  // qa:Qatar
  // 
  // ro:Romania
  // 
  // rs:Serbia
  // 
  // ru:Russian Federation
  // 
  // sa:Saudia Arabia
  // 
  // se:Sweden
  // 
  // sg:Singapore
  // 
  // sk:Slovak-Republic
  // 
  // th:Thailand
  // 
  // tr:Turkey
  // 
  // tw:Taiwan
  // 
  // tz:Tanzania
  // 
  // ua:Ukraine
  // 
  // us:USA
  // 
  // uy:Uruguay
  // 
  // uz:Uzbekistan
  // 
  // ve:Venezuela
  // 
  // vn:Vietnam
  // 
  // za:South Africa', 'zh_CN':'国家地区代号：
  // 1.单次最多查询 10 个国家；
  // 2.可传递的值：
  // 
  // ae：阿联酋
  // 
  // ar：阿根廷
  // 
  // at：奥地利
  // 
  // au：澳大利亚
  // 
  // bd：孟加拉
  // 
  // be：比利时
  // 
  // bn：文莱
  // 
  // br：巴西
  // 
  // by：白俄罗斯
  // 
  // ca：加拿大
  // 
  // ch：瑞士
  // 
  // cl：智利
  // 
  // cn：中国大陆
  // 
  // co：哥伦比亚
  // 
  // cz：捷克
  // 
  // de：德国
  // 
  // dk：丹麦
  // 
  // dz：阿尔及利亚
  // 
  // eg：埃及
  // 
  // es：西班牙
  // 
  // fi：芬兰
  // 
  // fr：法国
  // 
  // gb：英国
  // 
  // gr：希腊
  // 
  // hk：香港
  // 
  // hu：匈牙利
  // 
  // id：印度尼西亚
  // 
  // ie：爱尔兰
  // 
  // il：以色列
  // 
  // in：印度
  // 
  // iq：伊拉克
  // 
  // ir：伊朗
  // 
  // it：意大利
  // 
  // jp：日本
  // 
  // ke：肯尼亚
  // 
  // kg：吉尔吉斯斯坦
  // 
  // kh：柬埔寨
  // 
  // kr：韩国
  // 
  // kw：科威特
  // 
  // kz：哈萨克斯坦
  // 
  // la：老挝
  // 
  // lt：立陶宛
  // 
  // ma：摩洛哥
  // 
  // mm：缅甸
  // 
  // mn：蒙古
  // 
  // mo：澳门
  // 
  // mx：墨西哥
  // 
  // my：马来西亚
  // 
  // ng：尼日利亚
  // 
  // nl：荷兰
  // 
  // no：挪威
  // 
  // np：尼泊尔
  // 
  // nz：新西兰
  // 
  // pe：秘鲁
  // 
  // ph：菲律宾
  // 
  // pk：巴基斯坦
  // 
  // pl：波兰
  // 
  // pt：葡萄牙
  // 
  // qa：卡塔尔
  // 
  // ro：罗马尼亚
  // 
  // rs：塞尔维亚
  // 
  // ru：俄罗斯
  // 
  // sa：沙特阿拉伯
  // 
  // se：瑞典
  // 
  // sg：新加坡
  // 
  // sk：斯洛伐克
  // 
  // th：泰国
  // 
  // tr：土耳其
  // 
  // tw：台湾
  // 
  // tz：坦桑尼亚
  // 
  // ua：乌克兰
  // 
  // us：美国
  // 
  // uy：乌拉圭
  // 
  // uz：乌兹别克斯坦
  // 
  // ve：委内瑞拉
  // 
  // vn：越南
  // 
  // za：南非'}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpCountryCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportServerIpCountryCodeServiceRequest) SetDomain(v []*string) *ReportServerIpCountryCodeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportServerIpCountryCodeServiceRequest) SetCountryCode(v []*string) *ReportServerIpCountryCodeServiceRequest {
  s.CountryCode = v
  return s
}

type ReportServerIpCountryCodeServiceResponse struct {
  // {'en':'request result status code', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Request result information', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'domain', 'zh_CN':'-'}
  Data []*ReportServerIpCountryCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpCountryCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportServerIpCountryCodeServiceResponse) SetCode(v string) *ReportServerIpCountryCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportServerIpCountryCodeServiceResponse) SetMessage(v string) *ReportServerIpCountryCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportServerIpCountryCodeServiceResponse) SetData(v []*ReportServerIpCountryCodeServiceResponseData) *ReportServerIpCountryCodeServiceResponse {
  s.Data = v
  return s
}

type ReportServerIpCountryCodeServiceResponseData struct     {
  // {'en':'domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Detailed data on the results of the request', 'zh_CN':'请求结果的详细数据'}
  CountryData []*ReportServerIpCountryCodeServiceResponseDataCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpCountryCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportServerIpCountryCodeServiceResponseData) SetDomain(v string) *ReportServerIpCountryCodeServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportServerIpCountryCodeServiceResponseData) SetCountryData(v []*ReportServerIpCountryCodeServiceResponseDataCountryData) *ReportServerIpCountryCodeServiceResponseData {
  s.CountryData = v
  return s
}

type ReportServerIpCountryCodeServiceResponseDataCountryData struct     {
  // {'en':'Chinese country name', 'zh_CN':'中文国家名'}
  CountryNameZH *string `json:"countryNameZH,omitempty" xml:"countryNameZH,omitempty" require:"true"`
  // {'en':'English country name', 'zh_CN':'英文国家名'}
  CountryNameEN *string `json:"countryNameEN,omitempty" xml:"countryNameEN,omitempty" require:"true"`
  // {'en':'server IP list', 'zh_CN':'覆盖节点IP列表'}
  ServerIp []*string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true" type:"Repeated"`
}

func (s ReportServerIpCountryCodeServiceResponseDataCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceResponseDataCountryData) GoString() string {
  return s.String()
}

func (s *ReportServerIpCountryCodeServiceResponseDataCountryData) SetCountryNameZH(v string) *ReportServerIpCountryCodeServiceResponseDataCountryData {
  s.CountryNameZH = &v
  return s
}

func (s *ReportServerIpCountryCodeServiceResponseDataCountryData) SetCountryNameEN(v string) *ReportServerIpCountryCodeServiceResponseDataCountryData {
  s.CountryNameEN = &v
  return s
}

func (s *ReportServerIpCountryCodeServiceResponseDataCountryData) SetServerIp(v []*string) *ReportServerIpCountryCodeServiceResponseDataCountryData {
  s.ServerIp = v
  return s
}

type ReportServerIpCountryCodeServicePaths struct {
}

func (s ReportServerIpCountryCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServicePaths) GoString() string {
  return s.String()
}

type ReportServerIpCountryCodeServiceParameters struct {
}

func (s ReportServerIpCountryCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceParameters) GoString() string {
  return s.String()
}

type ReportServerIpCountryCodeServiceRequestHeader struct {
}

func (s ReportServerIpCountryCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportServerIpCountryCodeServiceResponseHeader struct {
}

func (s ReportServerIpCountryCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportServerIpCountryCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type IpDomainServiceRequest struct {
  // {"en":"IP", "zh_CN":"IP"}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *IpDomainServiceRequest) SetIp(v []*string) *IpDomainServiceRequest {
  s.Ip = v
  return s
}

type IpDomainServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*IpDomainServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *IpDomainServiceResponse) SetCode(v string) *IpDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *IpDomainServiceResponse) SetMessage(v string) *IpDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *IpDomainServiceResponse) SetData(v []*IpDomainServiceResponseData) *IpDomainServiceResponse {
  s.Data = v
  return s
}

type IpDomainServiceResponseData struct     {
  // {"en":"ip", "zh_CN":"IP名称"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Whether to use:
  // 
  //   idle --IP not used yet;
  //   runing -- IP in use;
  //   out of range -- IP is not in a queryable range", "zh_CN":"是否使用:
  //   idle -- 暂未使用;
  //   runing -- 使用中;
  //   out of range -- 不在查询范围内的ip"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"List of domain using this IP.The domain list of the IP that was idle or out of range was empty", "zh_CN":"用该IP的域名列表,未使用的ip/不在查询范围内的ip,域名列表为空"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponseData) GoString() string {
  return s.String()
}

func (s *IpDomainServiceResponseData) SetIp(v string) *IpDomainServiceResponseData {
  s.Ip = &v
  return s
}

func (s *IpDomainServiceResponseData) SetStatus(v string) *IpDomainServiceResponseData {
  s.Status = &v
  return s
}

func (s *IpDomainServiceResponseData) SetDomainList(v []*string) *IpDomainServiceResponseData {
  s.DomainList = v
  return s
}

type IpDomainServicePaths struct {
}

func (s IpDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServicePaths) GoString() string {
  return s.String()
}

type IpDomainServiceParameters struct {
}

func (s IpDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceParameters) GoString() string {
  return s.String()
}

type IpDomainServiceRequestHeader struct {
}

func (s IpDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type IpDomainServiceResponseHeader struct {
}

func (s IpDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type AkamaiIpPermitServiceRequest struct {
  // {"en":"accessToken", "zh_CN":"accessToken"}
  AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty" require:"true"`
  // {"en":"clientSecret", "zh_CN":"clientSecret"}
  ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty" require:"true"`
  // {"en":"clientToken", "zh_CN":"clientToken"}
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty" require:"true"`
  // {"en":"host", "zh_CN":"host"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"Unique identifier for each network list.", "zh_CN":"每个网络列表的唯一标识符。"}
  NetworkListId *string `json:"networkListId,omitempty" xml:"networkListId,omitempty" require:"true"`
  // {"en":"Support single IP, Support IP segment, example 8.7.6.0/24", "zh_CN":"支持单个ip, 支持IP段, 示例8.7.6.0/24"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s AkamaiIpPermitServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpPermitServiceRequest) GoString() string {
  return s.String()
}

func (s *AkamaiIpPermitServiceRequest) SetAccessToken(v string) *AkamaiIpPermitServiceRequest {
  s.AccessToken = &v
  return s
}

func (s *AkamaiIpPermitServiceRequest) SetClientSecret(v string) *AkamaiIpPermitServiceRequest {
  s.ClientSecret = &v
  return s
}

func (s *AkamaiIpPermitServiceRequest) SetClientToken(v string) *AkamaiIpPermitServiceRequest {
  s.ClientToken = &v
  return s
}

func (s *AkamaiIpPermitServiceRequest) SetHost(v string) *AkamaiIpPermitServiceRequest {
  s.Host = &v
  return s
}

func (s *AkamaiIpPermitServiceRequest) SetNetworkListId(v string) *AkamaiIpPermitServiceRequest {
  s.NetworkListId = &v
  return s
}

func (s *AkamaiIpPermitServiceRequest) SetIp(v string) *AkamaiIpPermitServiceRequest {
  s.Ip = &v
  return s
}

type AkamaiIpPermitServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s AkamaiIpPermitServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpPermitServiceResponse) GoString() string {
  return s.String()
}

func (s *AkamaiIpPermitServiceResponse) SetCode(v string) *AkamaiIpPermitServiceResponse {
  s.Code = &v
  return s
}

func (s *AkamaiIpPermitServiceResponse) SetMessage(v string) *AkamaiIpPermitServiceResponse {
  s.Message = &v
  return s
}

func (s *AkamaiIpPermitServiceResponse) SetData(v []*string) *AkamaiIpPermitServiceResponse {
  s.Data = v
  return s
}

type AkamaiIpPermitServicePaths struct {
}

func (s AkamaiIpPermitServicePaths) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpPermitServicePaths) GoString() string {
  return s.String()
}

type AkamaiIpPermitServiceParameters struct {
}

func (s AkamaiIpPermitServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpPermitServiceParameters) GoString() string {
  return s.String()
}

type AkamaiIpPermitServiceRequestHeader struct {
}

func (s AkamaiIpPermitServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpPermitServiceRequestHeader) GoString() string {
  return s.String()
}

type AkamaiIpPermitServiceResponseHeader struct {
}

func (s AkamaiIpPermitServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpPermitServiceResponseHeader) GoString() string {
  return s.String()
}




type AkamaiIpForbiddenServiceRequest struct {
  // {"en":"accessToken", "zh_CN":"accessToken"}
  AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty" require:"true"`
  // {"en":"clientSecret", "zh_CN":"clientSecret"}
  ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty" require:"true"`
  // {"en":"clientToken", "zh_CN":"clientToken"}
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty" require:"true"`
  // {"en":"host", "zh_CN":"host"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"Unique identifier for each network list.", "zh_CN":"每个网络列表的唯一标识符。"}
  NetworkListId *string `json:"networkListId,omitempty" xml:"networkListId,omitempty" require:"true"`
  // {"en":"1.The number of IP addresses to block is adjustable according to the account number.The default is 10000;
  // 2.Support IP segment, example 8.7.6.0/24", "zh_CN":"1. 要封禁IP地址，IP个数限制根据账号可调，默认为10000个
  // 2.支持IP段，示例 8.7.6.0/24"}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
  // {"en":"Mail receiver activated networkListId status after blocking successfully", "zh_CN":"封禁成功后激活networkListId 状态的邮件接收人"}
  Email []*string `json:"email,omitempty" xml:"email,omitempty" type:"Repeated"`
}

func (s AkamaiIpForbiddenServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpForbiddenServiceRequest) GoString() string {
  return s.String()
}

func (s *AkamaiIpForbiddenServiceRequest) SetAccessToken(v string) *AkamaiIpForbiddenServiceRequest {
  s.AccessToken = &v
  return s
}

func (s *AkamaiIpForbiddenServiceRequest) SetClientSecret(v string) *AkamaiIpForbiddenServiceRequest {
  s.ClientSecret = &v
  return s
}

func (s *AkamaiIpForbiddenServiceRequest) SetClientToken(v string) *AkamaiIpForbiddenServiceRequest {
  s.ClientToken = &v
  return s
}

func (s *AkamaiIpForbiddenServiceRequest) SetHost(v string) *AkamaiIpForbiddenServiceRequest {
  s.Host = &v
  return s
}

func (s *AkamaiIpForbiddenServiceRequest) SetNetworkListId(v string) *AkamaiIpForbiddenServiceRequest {
  s.NetworkListId = &v
  return s
}

func (s *AkamaiIpForbiddenServiceRequest) SetIp(v []*string) *AkamaiIpForbiddenServiceRequest {
  s.Ip = v
  return s
}

func (s *AkamaiIpForbiddenServiceRequest) SetEmail(v []*string) *AkamaiIpForbiddenServiceRequest {
  s.Email = v
  return s
}

type AkamaiIpForbiddenServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s AkamaiIpForbiddenServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpForbiddenServiceResponse) GoString() string {
  return s.String()
}

func (s *AkamaiIpForbiddenServiceResponse) SetCode(v string) *AkamaiIpForbiddenServiceResponse {
  s.Code = &v
  return s
}

func (s *AkamaiIpForbiddenServiceResponse) SetMessage(v string) *AkamaiIpForbiddenServiceResponse {
  s.Message = &v
  return s
}

func (s *AkamaiIpForbiddenServiceResponse) SetData(v []*string) *AkamaiIpForbiddenServiceResponse {
  s.Data = v
  return s
}

type AkamaiIpForbiddenServicePaths struct {
}

func (s AkamaiIpForbiddenServicePaths) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpForbiddenServicePaths) GoString() string {
  return s.String()
}

type AkamaiIpForbiddenServiceParameters struct {
}

func (s AkamaiIpForbiddenServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpForbiddenServiceParameters) GoString() string {
  return s.String()
}

type AkamaiIpForbiddenServiceRequestHeader struct {
}

func (s AkamaiIpForbiddenServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpForbiddenServiceRequestHeader) GoString() string {
  return s.String()
}

type AkamaiIpForbiddenServiceResponseHeader struct {
}

func (s AkamaiIpForbiddenServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AkamaiIpForbiddenServiceResponseHeader) GoString() string {
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




type QueryAllBandwidthLimitTaskListServiceRequest struct {
  // {"en":"Domain:
  // 1. The maximum number of domain is 100 by default (you can contact technical support for adjustment);
  // 2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为100个(可联系技术支持调整);
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether to ignore task status.
  // 0: No, only return the configuration whose task status is enabled; default: 0;
  // 1: Yes, return whether the task status is open or closed.", "zh_CN":"是否忽略任务状态。0:否,只返回任务状态为开启的配置;默认:0; 1:是,不论任务状态是开启还是关闭都返回。"}
  IgnoreTaskStatus *string `json:"ignoreTaskStatus,omitempty" xml:"ignoreTaskStatus,omitempty"`
  // {"en":"Whether the returned data contains all customer domain names involved in the task, the default is 0;
  // 0: no;
  // 1: Yes.", "zh_CN":"返回数据是否包含任务涉及的所有客户域名。0:否;默认:0;1:是。"}
  ContainDomain *string `json:"containDomain,omitempty" xml:"containDomain,omitempty"`
}

func (s QueryAllBandwidthLimitTaskListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceRequest) SetDomain(v []*string) *QueryAllBandwidthLimitTaskListServiceRequest {
  s.Domain = v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceRequest) SetIgnoreTaskStatus(v string) *QueryAllBandwidthLimitTaskListServiceRequest {
  s.IgnoreTaskStatus = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceRequest) SetContainDomain(v string) *QueryAllBandwidthLimitTaskListServiceRequest {
  s.ContainDomain = &v
  return s
}

type QueryAllBandwidthLimitTaskListServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryAllBandwidthLimitTaskListServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAllBandwidthLimitTaskListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceResponse) SetCode(v string) *QueryAllBandwidthLimitTaskListServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponse) SetMessage(v string) *QueryAllBandwidthLimitTaskListServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponse) SetData(v []*QueryAllBandwidthLimitTaskListServiceResponseData) *QueryAllBandwidthLimitTaskListServiceResponse {
  s.Data = v
  return s
}

type QueryAllBandwidthLimitTaskListServiceResponseData struct     {
  // {"en":"domain", "zh_CN":"客户域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Whether there is a configuration control task, 0: no, 1: yes", "zh_CN":"是否有配置控制任务,0表示没有,1表示有"}
  IsExist *string `json:"isExist,omitempty" xml:"isExist,omitempty" require:"true"`
  Content []*QueryAllBandwidthLimitTaskListServiceResponseDataContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAllBandwidthLimitTaskListServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseData) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseData) SetDomain(v string) *QueryAllBandwidthLimitTaskListServiceResponseData {
  s.Domain = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseData) SetIsExist(v string) *QueryAllBandwidthLimitTaskListServiceResponseData {
  s.IsExist = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseData) SetContent(v []*QueryAllBandwidthLimitTaskListServiceResponseDataContent) *QueryAllBandwidthLimitTaskListServiceResponseData {
  s.Content = v
  return s
}

type QueryAllBandwidthLimitTaskListServiceResponseDataContent struct     {
  // {"en":"taskName", "zh_CN":"任务名称"}
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty" require:"true"`
  // {"en":"Domain name configuration task types: 1. Static bandwidth control task, 2. Bandwidth buyout task, 3. Flow buyout task, 4. Request number buyout task, 5. Redundant pool speed limit task, 6. Back-to-source task, 7. POP running high scheduling task (this kind of task is quite special, the domain name is a global quantity, so as long as there is configuration, it will be enabled by default), 8. IP ban task", "zh_CN":"域名配置的任务类型:1静态带宽控制任务,2带宽买断任务,3流量买断任务,4请求数买断任务,5冗余池限速任务,6回源任务,7POP跑高调度任务(此种任务比较特殊,域名为全局量,所以只要有配置,就默认开启),8IP封禁任务"}
  TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty" require:"true"`
  // {"en":"Task status, 0: task is off, 1: task is on.", "zh_CN":"任务状态,0表示任务关闭,1表示任务开启。"}
  TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty" require:"true"`
  // {"en":"Control strategy. 0 means squid default, 1 rejects, 2 when taskType=8, is ladder blocked, the rest is redirected ip, 3 when taskType=6, is backup source, the rest is speed limit = rejection + maximum download rate, 4 speed limit = rejection + timeout to disconnect, 5 does not process, controls each connection at minimum speed, does not process the excess part, 6 backups, 7 redirects domain name, 8 redirects URL", "zh_CN":"控制策略。0表示squid默认,1拒绝,2当taskType=8时,为阶梯封禁,其余为重定向ip,3当taskType=6时,为主备回源,其余为限速=拒绝+最大下载速率,4限速=拒绝+超时断开连接,5不处理,按最小速率控制每个连接,超出部分不处理,6回源,7重定向域名,8重定向URL"}
  CtrlMode *string `json:"ctrlMode,omitempty" xml:"ctrlMode,omitempty" require:"true"`
  // {"en":"Bandwidth limit value, when taskType=1,2,5,6, in Mbps, when taskType=3, in G, when taskType=4, in MH (millions), when taskType=7, in -1, there is no bandwidth limit value, when taskType=8, in seconds.", "zh_CN":"带宽限制值,当taskType=1,2,5,6时,单位为Mbps,当taskType=3,单位为G,当taskType=4,单位为MH(百万个),当taskType=7,为-1,表示没有带宽限制值,当taskType=8,单位为次。"}
  CtrlValue *string `json:"ctrlValue,omitempty" xml:"ctrlValue,omitempty" require:"true"`
  // {"en":"List of customer domains involved under the task. Values are only available when entering ContainOrichannelName=1.", "zh_CN":"任务下涉及的客户域名列表。只有在入参containOrichannelName=1的时候有值。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAllBandwidthLimitTaskListServiceResponseDataContent) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseDataContent) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetTaskName(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.TaskName = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetTaskType(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.TaskType = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetTaskStatus(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.TaskStatus = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetCtrlMode(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.CtrlMode = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetCtrlValue(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.CtrlValue = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetDomainList(v []*string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.DomainList = v
  return s
}

type QueryAllBandwidthLimitTaskListServicePaths struct {
}

func (s QueryAllBandwidthLimitTaskListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServicePaths) GoString() string {
  return s.String()
}

type QueryAllBandwidthLimitTaskListServiceParameters struct {
}

func (s QueryAllBandwidthLimitTaskListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceParameters) GoString() string {
  return s.String()
}

type QueryAllBandwidthLimitTaskListServiceRequestHeader struct {
}

func (s QueryAllBandwidthLimitTaskListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryAllBandwidthLimitTaskListServiceResponseHeader struct {
}

func (s QueryAllBandwidthLimitTaskListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseHeader) GoString() string {
  return s.String()
}




type AddOrRemoveForbiddingIPWhitelistServiceRequest struct {
  // {"en":"Whitelist Objects", "zh_CN":"白名单操作对象列表 "}
  AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList []*AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList `json:"whitelistObjectList,omitempty" xml:"whitelistObjectList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operation Type, 1:add, 2:remove ", "zh_CN":"操作类型，1:新增，2:删除 "}
  OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Whitelist Type, 0:'Customer Granularity', 1:'Domain Granularity', 2:'Label Granularity' ", "zh_CN":"白名单类型，0:客户粒度，1:域名粒度，2:标签粒度 "}
  WhitelistType *string `json:"whitelistType,omitempty" xml:"whitelistType,omitempty" require:"true"`
  // {"en":"Force resume,default 'false'. When adding, simultaneously resume IPs that are still forbidding at the corresponding granularity ", "zh_CN":"是否强制解禁，默认否。添加时，同时解禁对应粒度下还处于封禁中的IP。"}
  ForceResume *bool `json:"forceResume,omitempty" xml:"forceResume,omitempty"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequest) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceRequest) SetWhitelistObjectList(v []*AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList) *AddOrRemoveForbiddingIPWhitelistServiceRequest {
  s.AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList = v
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

type AddOrRemoveForbiddingIPWhitelistServiceResponse struct {
  // {"en":"Result Code.If it shows `36010032`,means `partial success`, please pay attention to the details of partial failures in errCode.", "zh_CN":"响应码，成功为0，如果为“36010032”部分成功，请关注errCode中部分失败的详情 "}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data []*AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *AddOrRemoveForbiddingIPWhitelistServiceResponse) SetData(v []*AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) *AddOrRemoveForbiddingIPWhitelistServiceResponse {
  s.Data = v
  return s
}

type AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList struct {
  // {"en":"Whitelist Object,must be Domain or Label Code.When the whitelist type is `customer granularity`, this field must be empty ", "zh_CN":"白名单操作对象，白名单类型“域名粒度”时，请填写加速域名，白名单类型“标签粒度”时，请填写标签编码，白名单类型为“客户粒度”时，此值必须为空 "}
  WhitelistObject *string `json:"whitelistObject,omitempty" xml:"whitelistObject,omitempty"`
  // {"en":"List of Whitelist IPs ", "zh_CN":"白名单IP列表"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList) SetWhitelistObject(v string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList {
  s.WhitelistObject = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList) SetIpList(v []*string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectList {
  s.IpList = v
  return s
}

type AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult struct {
  // {"en":"Error Code", "zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message", "zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Whitelist Object", "zh_CN":"白名单操作对象"}
  WhitelistObject *string `json:"whitelistObject,omitempty" xml:"whitelistObject,omitempty" require:"true"`
  // {"en":"Whitelist Type, 0:'Customer Granularity', 1:'Domain Granularity', 2:'Label Granularity' ", "zh_CN":"白名单类型，0:客户粒度，1:域名粒度，2:标签粒度 "}
  WhitelistType *string `json:"whitelistType,omitempty" xml:"whitelistType,omitempty" require:"true"`
  // {"en":"List of Failed IPs", "zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) GoString() string {
  return s.String()
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) SetErrCode(v string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult {
  s.ErrCode = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) SetErrMessage(v string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult {
  s.ErrMessage = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) SetWhitelistObject(v string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult {
  s.WhitelistObject = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) SetWhitelistType(v string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult {
  s.WhitelistType = &v
  return s
}

func (s *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult) SetFailedIpList(v []*string) *AddOrRemoveForbiddingIPWhitelistServiceWhitelistObjectOperateResult {
  s.FailedIpList = v
  return s
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

type AddOrRemoveForbiddingIPWhitelistServiceRequestHeader struct {
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceRequestHeader) GoString() string {
  return s.String()
}

type AddOrRemoveForbiddingIPWhitelistServiceResponseHeader struct {
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddOrRemoveForbiddingIPWhitelistServiceResponseHeader) GoString() string {
  return s.String()
}




