package toolservice

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryDeviceOperationalInfoServiceRequest struct {
  // {"en":"Device serial number (SN)","zh_CN":"设备序列号"}
  Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
  // {"en":"Query type. Dictionary `srAgentQueryType`. `0` indicates querying for qualifiable status.","zh_CN":"字典：srAgentQueryType  0:查询可质检状态"}
  QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
}

func (s QueryDeviceOperationalInfoServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDeviceOperationalInfoServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryDeviceOperationalInfoServiceRequest) SetSn(v string) *QueryDeviceOperationalInfoServiceRequest {
  s.Sn = &v
  return s
}

func (s *QueryDeviceOperationalInfoServiceRequest) SetQueryType(v string) *QueryDeviceOperationalInfoServiceRequest {
  s.QueryType = &v
  return s
}

type QueryDeviceOperationalInfoServiceRequestHeader struct {
}

func (s QueryDeviceOperationalInfoServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeviceOperationalInfoServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryDeviceOperationalInfoServicePaths struct {
}

func (s QueryDeviceOperationalInfoServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDeviceOperationalInfoServicePaths) GoString() string {
  return s.String()
}

type QueryDeviceOperationalInfoServiceParameters struct {
}

func (s QueryDeviceOperationalInfoServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDeviceOperationalInfoServiceParameters) GoString() string {
  return s.String()
}

type QueryDeviceOperationalInfoServiceResponse struct {
  // {"en":"Service request ID","zh_CN":"服务请求ID"}
  SrId *string `json:"srId,omitempty" xml:"srId,omitempty" require:"true"`
  // {"en":"Device serial number","zh_CN":"设备序列号"}
  Sn *string `json:"sn,omitempty" xml:"sn,omitempty" require:"true"`
  // {"en":"Query result code. `000`: Indicates SN conflict (an identical SN exists in the database but is retired) or existing device does not meet quality inspection conditions (device is in use or a fault order for quality inspection has not been applied). `001`: Existing device, awaiting quality inspection. `002`: SN does not exist. `999`: Query content is currently not supported.","zh_CN":"查询结果,返回Code（000/001/002/999）  000：对应两种情况  1、SN冲突，库内存在同名已退用的SN。  2、存量设备但不满足质检条件（设备在用或者故障单未申请质检）  001：存量设备，等待质检  002：SN不存在  999：要查询的内容暂未支持"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
  // {"en":"Reference message","zh_CN":"参考信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QueryDeviceOperationalInfoServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDeviceOperationalInfoServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryDeviceOperationalInfoServiceResponse) SetSrId(v string) *QueryDeviceOperationalInfoServiceResponse {
  s.SrId = &v
  return s
}

func (s *QueryDeviceOperationalInfoServiceResponse) SetSn(v string) *QueryDeviceOperationalInfoServiceResponse {
  s.Sn = &v
  return s
}

func (s *QueryDeviceOperationalInfoServiceResponse) SetResult(v string) *QueryDeviceOperationalInfoServiceResponse {
  s.Result = &v
  return s
}

func (s *QueryDeviceOperationalInfoServiceResponse) SetMessage(v string) *QueryDeviceOperationalInfoServiceResponse {
  s.Message = &v
  return s
}

type QueryDeviceOperationalInfoServiceResponseHeader struct {
}

func (s QueryDeviceOperationalInfoServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDeviceOperationalInfoServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryConversionTaskListRequest struct {
  // {"en":"Start Time. Format: yyyy-MM-dd HH:mm:ss","zh_CN":"开始时间。格式为yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"End Time. Format: yyyy-MM-dd HH:mm:ss","zh_CN":"结束时间。格式为yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"Task ID or domain name. Supports fuzzy search","zh_CN":"任务ID或者域名。支持模糊查询"}
  SearchText *string `json:"searchText,omitempty" xml:"searchText,omitempty"`
}

func (s QueryConversionTaskListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListRequest) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskListRequest) SetStartTime(v string) *QueryConversionTaskListRequest {
  s.StartTime = &v
  return s
}

func (s *QueryConversionTaskListRequest) SetEndTime(v string) *QueryConversionTaskListRequest {
  s.EndTime = &v
  return s
}

func (s *QueryConversionTaskListRequest) SetSearchText(v string) *QueryConversionTaskListRequest {
  s.SearchText = &v
  return s
}

type QueryConversionTaskListRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:\n1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.\n2. If the parameter is unspecified,results will be in the GMT timezone.","zh_CN":"报表数据时区： \n1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8\n2.若参数未传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryConversionTaskListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskListRequestHeader) SetServiceType(v string) *QueryConversionTaskListRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *QueryConversionTaskListRequestHeader) SetTimezone(v string) *QueryConversionTaskListRequestHeader {
  s.Timezone = &v
  return s
}

type QueryConversionTaskListPaths struct {
}

func (s QueryConversionTaskListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListPaths) GoString() string {
  return s.String()
}

type QueryConversionTaskListParameters struct {
}

func (s QueryConversionTaskListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListParameters) GoString() string {
  return s.String()
}

type QueryConversionTaskListResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Status code.","zh_CN":"状态码。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Detailed information of the interface response.","zh_CN":"接口响应的详细信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"The specific business data list returned by the interface, containing detailed conversion task information.","zh_CN":"接口返回的具体业务数据列表，包含转换任务的详细信息。"}
  Data []*QueryConversionTaskListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryConversionTaskListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListResponse) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskListResponse) SetCode(v int) *QueryConversionTaskListResponse {
  s.Code = &v
  return s
}

func (s *QueryConversionTaskListResponse) SetMsg(v string) *QueryConversionTaskListResponse {
  s.Msg = &v
  return s
}

func (s *QueryConversionTaskListResponse) SetData(v []*QueryConversionTaskListResponseData) *QueryConversionTaskListResponse {
  s.Data = v
  return s
}

type QueryConversionTaskListResponseData struct     {
  // {"en":"Task creation time.","zh_CN":"任务创建时间。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Task ID.","zh_CN":"任务ID。"}
  TaskNo *string `json:"taskNo,omitempty" xml:"taskNo,omitempty" require:"true"`
  // {"en":"Configuration file type. \nDOMAIN: Akamai Web security configuration \nCLIENT_LIST: Akamai client/network list","zh_CN":"配置文件类型。\nDOMAIN：Akamai Web安全配置\nCLIENT_LIST：Akamai 客户端/网络列表"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Task status. \nPROCESSING: Processing \nSUCCESS: Success FAIL: Fail","zh_CN":"任务状态。\nPROCESSING：处理中\nSUCCESS：成功\nFAIL：失败"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s QueryConversionTaskListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListResponseData) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskListResponseData) SetCreateTime(v string) *QueryConversionTaskListResponseData {
  s.CreateTime = &v
  return s
}

func (s *QueryConversionTaskListResponseData) SetTaskNo(v string) *QueryConversionTaskListResponseData {
  s.TaskNo = &v
  return s
}

func (s *QueryConversionTaskListResponseData) SetType(v string) *QueryConversionTaskListResponseData {
  s.Type = &v
  return s
}

func (s *QueryConversionTaskListResponseData) SetStatus(v string) *QueryConversionTaskListResponseData {
  s.Status = &v
  return s
}

type QueryConversionTaskListResponseHeader struct {
}

func (s QueryConversionTaskListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskListResponseHeader) GoString() string {
  return s.String()
}




type PublishConvertedDomainRequest struct {
  // {"en":"List of domains to be published.","zh_CN":"待发布的域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s PublishConvertedDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s PublishConvertedDomainRequest) GoString() string {
  return s.String()
}

func (s *PublishConvertedDomainRequest) SetDomainList(v []*string) *PublishConvertedDomainRequest {
  s.DomainList = v
  return s
}

type PublishConvertedDomainRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s PublishConvertedDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PublishConvertedDomainRequestHeader) GoString() string {
  return s.String()
}

func (s *PublishConvertedDomainRequestHeader) SetServiceType(v string) *PublishConvertedDomainRequestHeader {
  s.ServiceType = &v
  return s
}

type PublishConvertedDomainPaths struct {
}

func (s PublishConvertedDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s PublishConvertedDomainPaths) GoString() string {
  return s.String()
}

type PublishConvertedDomainParameters struct {
}

func (s PublishConvertedDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s PublishConvertedDomainParameters) GoString() string {
  return s.String()
}

type PublishConvertedDomainResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Status code.","zh_CN":"状态码。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message.","zh_CN":"响应信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s PublishConvertedDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s PublishConvertedDomainResponse) GoString() string {
  return s.String()
}

func (s *PublishConvertedDomainResponse) SetCode(v int) *PublishConvertedDomainResponse {
  s.Code = &v
  return s
}

func (s *PublishConvertedDomainResponse) SetMsg(v string) *PublishConvertedDomainResponse {
  s.Msg = &v
  return s
}

type PublishConvertedDomainResponseHeader struct {
}

func (s PublishConvertedDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PublishConvertedDomainResponseHeader) GoString() string {
  return s.String()
}




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

type QueryBandwidthLimitTaskListServiceRequestHeader struct {
}

func (s QueryBandwidthLimitTaskListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceRequestHeader) GoString() string {
  return s.String()
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

type QueryBandwidthLimitTaskListServiceResponse struct {
  // {"en":"result","zh_CN":"结果"}
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
  // {"en":"Domain","zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Maximum bandwidth set","zh_CN":"设置的最大带宽值"}
  BandwidthLimit *int `json:"bandwidthLimit,omitempty" xml:"bandwidthLimit,omitempty" require:"true"`
  // {"en":"Task name","zh_CN":"任务名称"}
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

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetBandwidthLimit(v int) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.BandwidthLimit = &v
  return s
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetTaskName(v string) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.TaskName = &v
  return s
}

type QueryBandwidthLimitTaskListServiceResponseHeader struct {
}

func (s QueryBandwidthLimitTaskListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryConversionTaskDetailRequest struct {
  // {"en":"Task ID","zh_CN":"任务ID"}
  TaskNo *string `json:"taskNo,omitempty" xml:"taskNo,omitempty" require:"true"`
  // {"en":"Configuration file type. \nDOMAIN: Akamai Web security configuration\nCLIENT_LIST: Akamai client/network list","zh_CN":"配置文件类型。\nDOMAIN：Akamai Web安全配置\nCLIENT_LIST：Akamai 客户端/网络列表"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s QueryConversionTaskDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskDetailRequest) SetTaskNo(v string) *QueryConversionTaskDetailRequest {
  s.TaskNo = &v
  return s
}

func (s *QueryConversionTaskDetailRequest) SetType(v string) *QueryConversionTaskDetailRequest {
  s.Type = &v
  return s
}

type QueryConversionTaskDetailRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s QueryConversionTaskDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskDetailRequestHeader) SetServiceType(v string) *QueryConversionTaskDetailRequestHeader {
  s.ServiceType = &v
  return s
}

type QueryConversionTaskDetailPaths struct {
}

func (s QueryConversionTaskDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailPaths) GoString() string {
  return s.String()
}

type QueryConversionTaskDetailParameters struct {
}

func (s QueryConversionTaskDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailParameters) GoString() string {
  return s.String()
}

type QueryConversionTaskDetailResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Status code.","zh_CN":"状态码。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message.","zh_CN":"响应信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"响应数据。"}
  Data *QueryConversionTaskDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryConversionTaskDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskDetailResponse) SetCode(v int) *QueryConversionTaskDetailResponse {
  s.Code = &v
  return s
}

func (s *QueryConversionTaskDetailResponse) SetMsg(v string) *QueryConversionTaskDetailResponse {
  s.Msg = &v
  return s
}

func (s *QueryConversionTaskDetailResponse) SetData(v *QueryConversionTaskDetailResponseData) *QueryConversionTaskDetailResponse {
  s.Data = v
  return s
}

type QueryConversionTaskDetailResponseData struct {
  // {"en":"Task details.","zh_CN":"任务详情。"}
  TaskDetail *QueryConversionTaskDetailResponseDataTaskDetail `json:"taskDetail,omitempty" xml:"taskDetail,omitempty" require:"true" type:"Struct"`
  // {"en":"File name. Multiple files are separated by semicolons.","zh_CN":"文件名称。多个文件使用;分隔。"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"Task creation time.","zh_CN":"任务创建时间。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Domain list.","zh_CN":"域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Task ID.","zh_CN":"任务ID。"}
  TaskNo *string `json:"taskNo,omitempty" xml:"taskNo,omitempty" require:"true"`
  // {"en":"Configuration file type. \nDOMAIN: Akamai Web security configuration\nCLIENT_LIST: Akamai client/network list","zh_CN":"配置文件类型。\nDOMAIN：Akamai Web安全配置\nCLIENT_LIST：Akamai 客户端/网络列表"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Task status. \nPROCESSING: Processing \nSUCCESS: Success \nFAIL: Fail","zh_CN":"任务状态。\nPROCESSING：处理中\nSUCCESS：成功\nFAIL：失败"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Verification environment information.","zh_CN":"验证环境信息。"}
  DeployObject *string `json:"deployObject,omitempty" xml:"deployObject,omitempty" require:"true"`
}

func (s QueryConversionTaskDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailResponseData) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskDetailResponseData) SetTaskDetail(v *QueryConversionTaskDetailResponseDataTaskDetail) *QueryConversionTaskDetailResponseData {
  s.TaskDetail = v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetFileName(v string) *QueryConversionTaskDetailResponseData {
  s.FileName = &v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetCreateTime(v string) *QueryConversionTaskDetailResponseData {
  s.CreateTime = &v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetDomainList(v []*string) *QueryConversionTaskDetailResponseData {
  s.DomainList = v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetTaskNo(v string) *QueryConversionTaskDetailResponseData {
  s.TaskNo = &v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetType(v string) *QueryConversionTaskDetailResponseData {
  s.Type = &v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetStatus(v string) *QueryConversionTaskDetailResponseData {
  s.Status = &v
  return s
}

func (s *QueryConversionTaskDetailResponseData) SetDeployObject(v string) *QueryConversionTaskDetailResponseData {
  s.DeployObject = &v
  return s
}

type QueryConversionTaskDetailResponseDataTaskDetail struct {
  // {"en":"Conversion module.","zh_CN":"转换模块。"}
  CLIENTLIST []*QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST `json:"CLIENT_LIST,omitempty" xml:"CLIENT_LIST,omitempty" require:"true" type:"Repeated"`
}

func (s QueryConversionTaskDetailResponseDataTaskDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailResponseDataTaskDetail) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskDetailResponseDataTaskDetail) SetCLIENTLIST(v []*QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST) *QueryConversionTaskDetailResponseDataTaskDetail {
  s.CLIENTLIST = v
  return s
}

type QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST struct     {
  // {"en":"Conversion information.","zh_CN":"转换信息。"}
  ParseMessage *string `json:"parseMessage,omitempty" xml:"parseMessage,omitempty" require:"true"`
  // {"en":"Conversion result.","zh_CN":"转换结果。"}
  ParseResult *string `json:"parseResult,omitempty" xml:"parseResult,omitempty" require:"true"`
}

func (s QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST) GoString() string {
  return s.String()
}

func (s *QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST) SetParseMessage(v string) *QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST {
  s.ParseMessage = &v
  return s
}

func (s *QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST) SetParseResult(v string) *QueryConversionTaskDetailResponseDataTaskDetailCLIENTLIST {
  s.ParseResult = &v
  return s
}

type QueryConversionTaskDetailResponseHeader struct {
}

func (s QueryConversionTaskDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryConversionTaskDetailResponseHeader) GoString() string {
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




type BandwidthLimitServiceRequest struct {
  // {"en":"Task: Limits to the number of tasks can be adjusted depending on different accounts. The default value is 3","zh_CN":"任务：任务个数限制根据账号可调，默认为3个"}
  Task *string `json:"task,omitempty" xml:"task,omitempty"`
  // {"en":"Operation type: \nenable: set, update or enable bandwidth limit; \ndisable: to disable bandwidth limit","zh_CN":"操作类型：\nenable 设置、更新或开启带宽限制，\ndisable 关闭带宽限制"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en":"Domain list.","zh_CN":"域名列表"}
  DomainList *string `json:"domain-list,omitempty" xml:"domain-list,omitempty"`
  // {"en":"Domain, must follow regular expression rule of (([\w-]{1,62})?(\.[\w-]{1,62})+)","zh_CN":"域名，必须符合正则(([\w-]{1,62})?(\.[\w-]{1,62})+)"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
  // {"en":"1. Cancel the control ratio (the proportion to which the threshold is reduced, and cancel the control bandwidth), calculated as a percentage;\n\n2. The input value is a positive integer. When no parameters are passed, the default is 60.","zh_CN":"1.取消控制比例（阈值降到多少比例，取消控制带宽），按百分比计算；\n\n2.输入值为正整数。未传参时，默认为40。"}
  CtrlMinRatio *int `json:"ctrlMinRatio,omitempty" xml:"ctrlMinRatio,omitempty"`
  // {"en":"1. Control the effective ratio (what ratio the threshold reaches and start to control the bandwidth), calculated as a percentage;\n2. The input value is a positive integer. When no parameter is passed, the default is 60;\n3. The effective ratio of control is greater than the ratio of canceled control;\n4. The effective control ratio and the cancel control ratio need to be paired and configured.","zh_CN":"1.控制生效比例（ 阈值达到多少比例，开始控制带宽），按百分比计算；\n\n2.输入值为正整数。未传参时，默认为60；\n\n3.控制生效比例 要大于 取消控制比例；\n\n4.控制生效比例 和 取消控制比例 需配对配置。"}
  CtrlInitRatio *int `json:"ctrlInitRatio,omitempty" xml:"ctrlInitRatio,omitempty"`
  // {"en":"Bandwidth limit, positive integer, unit is Mbps. This filed is required when action is enable","zh_CN":"带宽限制值，为正整数，单位为Mbps，当action为enable时为必选项"}
  BandwidthLimit *int `json:"bandwidth-limit,omitempty" xml:"bandwidth-limit,omitempty"`
}

func (s BandwidthLimitServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceRequest) GoString() string {
  return s.String()
}

func (s *BandwidthLimitServiceRequest) SetTask(v string) *BandwidthLimitServiceRequest {
  s.Task = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetAction(v string) *BandwidthLimitServiceRequest {
  s.Action = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetDomainList(v string) *BandwidthLimitServiceRequest {
  s.DomainList = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetDomainName(v string) *BandwidthLimitServiceRequest {
  s.DomainName = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetCtrlMinRatio(v int) *BandwidthLimitServiceRequest {
  s.CtrlMinRatio = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetCtrlInitRatio(v int) *BandwidthLimitServiceRequest {
  s.CtrlInitRatio = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetBandwidthLimit(v int) *BandwidthLimitServiceRequest {
  s.BandwidthLimit = &v
  return s
}

type BandwidthLimitServiceRequestHeader struct {
}

func (s BandwidthLimitServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceRequestHeader) GoString() string {
  return s.String()
}

type BandwidthLimitServicePaths struct {
}

func (s BandwidthLimitServicePaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServicePaths) GoString() string {
  return s.String()
}

type BandwidthLimitServiceParameters struct {
}

func (s BandwidthLimitServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceParameters) GoString() string {
  return s.String()
}

type BandwidthLimitServiceResponse struct {
  // {"en":"Task ID","zh_CN":"任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"Domain","zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Status codes, 1 for success and 0   means failed","zh_CN":"状态码，1：成功，0：失败"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description of results","zh_CN":"结果描述信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s BandwidthLimitServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceResponse) GoString() string {
  return s.String()
}

func (s *BandwidthLimitServiceResponse) SetTaskId(v string) *BandwidthLimitServiceResponse {
  s.TaskId = &v
  return s
}

func (s *BandwidthLimitServiceResponse) SetDomainName(v string) *BandwidthLimitServiceResponse {
  s.DomainName = &v
  return s
}

func (s *BandwidthLimitServiceResponse) SetCode(v int) *BandwidthLimitServiceResponse {
  s.Code = &v
  return s
}

func (s *BandwidthLimitServiceResponse) SetMsg(v string) *BandwidthLimitServiceResponse {
  s.Msg = &v
  return s
}

type BandwidthLimitServiceResponseHeader struct {
}

func (s BandwidthLimitServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceResponseHeader) GoString() string {
  return s.String()
}




type CreateWebSecurityConfigurationTaskRequest struct {
  // {"en":"Domain list. Limit of 100 domains per request.","zh_CN":"域名列表。单次限制100个域名。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Name of the configuration file to be uploaded. Must include file extension.","zh_CN":"需要上传的配置文件的名称。需要包含文件后缀。"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"File content, a single line string obtained by standard base64 encoding of the XML configuration file.","zh_CN":"文件内容，对XML配置文件进行标准base64编码后得到的单行字符串"}
  FileContentBase64 *string `json:"fileContentBase64,omitempty" xml:"fileContentBase64,omitempty" require:"true"`
}

func (s CreateWebSecurityConfigurationTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateWebSecurityConfigurationTaskRequest) SetDomainList(v []*string) *CreateWebSecurityConfigurationTaskRequest {
  s.DomainList = v
  return s
}

func (s *CreateWebSecurityConfigurationTaskRequest) SetFileName(v string) *CreateWebSecurityConfigurationTaskRequest {
  s.FileName = &v
  return s
}

func (s *CreateWebSecurityConfigurationTaskRequest) SetFileContentBase64(v string) *CreateWebSecurityConfigurationTaskRequest {
  s.FileContentBase64 = &v
  return s
}

type CreateWebSecurityConfigurationTaskRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateWebSecurityConfigurationTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateWebSecurityConfigurationTaskRequestHeader) SetServiceType(v string) *CreateWebSecurityConfigurationTaskRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateWebSecurityConfigurationTaskPaths struct {
}

func (s CreateWebSecurityConfigurationTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskPaths) GoString() string {
  return s.String()
}

type CreateWebSecurityConfigurationTaskParameters struct {
}

func (s CreateWebSecurityConfigurationTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskParameters) GoString() string {
  return s.String()
}

type CreateWebSecurityConfigurationTaskResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code.","zh_CN":"请参照错误码。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description information.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Specific business data returned by the interface, including the task ID.","zh_CN":"接口返回的具体业务数据，包含任务ID。"}
  Data *CreateWebSecurityConfigurationTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateWebSecurityConfigurationTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskResponse) GoString() string {
  return s.String()
}

func (s *CreateWebSecurityConfigurationTaskResponse) SetCode(v int) *CreateWebSecurityConfigurationTaskResponse {
  s.Code = &v
  return s
}

func (s *CreateWebSecurityConfigurationTaskResponse) SetMsg(v string) *CreateWebSecurityConfigurationTaskResponse {
  s.Msg = &v
  return s
}

func (s *CreateWebSecurityConfigurationTaskResponse) SetData(v *CreateWebSecurityConfigurationTaskResponseData) *CreateWebSecurityConfigurationTaskResponse {
  s.Data = v
  return s
}

type CreateWebSecurityConfigurationTaskResponseData struct {
  // {"en":"Task ID.","zh_CN":"任务ID。"}
  TaskNo *string `json:"taskNo,omitempty" xml:"taskNo,omitempty" require:"true"`
  // {"en":"Client/network lists that are referenced but not uploaded in the Web Security Configuration.Data format:{policyId1: list name1, policyId2: list name2}","zh_CN":"Web安全配置中有引用但未上传的客户端/网络列表列表信息。返回数据格式：{policyId1: 列表名称1, policyId2: 列表名称2}"}
  Errors *CreateWebSecurityConfigurationTaskResponseDataErrors `json:"errors,omitempty" xml:"errors,omitempty" require:"true" type:"Struct"`
}

func (s CreateWebSecurityConfigurationTaskResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskResponseData) GoString() string {
  return s.String()
}

func (s *CreateWebSecurityConfigurationTaskResponseData) SetTaskNo(v string) *CreateWebSecurityConfigurationTaskResponseData {
  s.TaskNo = &v
  return s
}

func (s *CreateWebSecurityConfigurationTaskResponseData) SetErrors(v *CreateWebSecurityConfigurationTaskResponseDataErrors) *CreateWebSecurityConfigurationTaskResponseData {
  s.Errors = v
  return s
}

type CreateWebSecurityConfigurationTaskResponseDataErrors struct {
}

func (s CreateWebSecurityConfigurationTaskResponseDataErrors) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskResponseDataErrors) GoString() string {
  return s.String()
}

type CreateWebSecurityConfigurationTaskResponseHeader struct {
}

func (s CreateWebSecurityConfigurationTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWebSecurityConfigurationTaskResponseHeader) GoString() string {
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




type CreateClientNetworkListTaskRequest struct {
  // {"en":"List of configuration transformations.","zh_CN":"转换配置列表。"}
  MoveList []*CreateClientNetworkListTaskRequestMoveList `json:"moveList,omitempty" xml:"moveList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClientNetworkListTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateClientNetworkListTaskRequest) SetMoveList(v []*CreateClientNetworkListTaskRequestMoveList) *CreateClientNetworkListTaskRequest {
  s.MoveList = v
  return s
}

type CreateClientNetworkListTaskRequestMoveList struct     {
  // {"en":"The name of the configuration file to be uploaded. It must include the file extension.","zh_CN":"需要上传的配置文件的名称。需要包含文件后缀。"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"File type.\nCLIENT_LIST_JSON: Client list in JSON file format\nCLIENT_LIST_CSV: Client list in CSV file format\nNETWORK_LIST_JSON: Network list in JSON file format","zh_CN":"文件类型。\nCLIENT_LIST_JSON：客户端列表JSON文件格式\nCLIENT_LIST_CSV：客户端列表CSV文件格式\nNETWORK_LIST_JSON：网络列表JSON文件格式"}
  MoveType *string `json:"moveType,omitempty" xml:"moveType,omitempty" require:"true"`
  // {"en":"Base64 encoded file content.","zh_CN":"base64加密的文件内容。"}
  FileContentBase64 *string `json:"fileContentBase64,omitempty" xml:"fileContentBase64,omitempty" require:"true"`
}

func (s CreateClientNetworkListTaskRequestMoveList) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskRequestMoveList) GoString() string {
  return s.String()
}

func (s *CreateClientNetworkListTaskRequestMoveList) SetFileName(v string) *CreateClientNetworkListTaskRequestMoveList {
  s.FileName = &v
  return s
}

func (s *CreateClientNetworkListTaskRequestMoveList) SetMoveType(v string) *CreateClientNetworkListTaskRequestMoveList {
  s.MoveType = &v
  return s
}

func (s *CreateClientNetworkListTaskRequestMoveList) SetFileContentBase64(v string) *CreateClientNetworkListTaskRequestMoveList {
  s.FileContentBase64 = &v
  return s
}

type CreateClientNetworkListTaskRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateClientNetworkListTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateClientNetworkListTaskRequestHeader) SetServiceType(v string) *CreateClientNetworkListTaskRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateClientNetworkListTaskPaths struct {
}

func (s CreateClientNetworkListTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskPaths) GoString() string {
  return s.String()
}

type CreateClientNetworkListTaskParameters struct {
}

func (s CreateClientNetworkListTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskParameters) GoString() string {
  return s.String()
}

type CreateClientNetworkListTaskResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Status code.","zh_CN":"状态码。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message.","zh_CN":"响应信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"响应数据。"}
  Data *CreateClientNetworkListTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateClientNetworkListTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskResponse) GoString() string {
  return s.String()
}

func (s *CreateClientNetworkListTaskResponse) SetCode(v int) *CreateClientNetworkListTaskResponse {
  s.Code = &v
  return s
}

func (s *CreateClientNetworkListTaskResponse) SetMsg(v string) *CreateClientNetworkListTaskResponse {
  s.Msg = &v
  return s
}

func (s *CreateClientNetworkListTaskResponse) SetData(v *CreateClientNetworkListTaskResponseData) *CreateClientNetworkListTaskResponse {
  s.Data = v
  return s
}

type CreateClientNetworkListTaskResponseData struct {
  // {"en":"Task ID.","zh_CN":"任务ID。"}
  TaskNo *string `json:"taskNo,omitempty" xml:"taskNo,omitempty" require:"true"`
}

func (s CreateClientNetworkListTaskResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskResponseData) GoString() string {
  return s.String()
}

func (s *CreateClientNetworkListTaskResponseData) SetTaskNo(v string) *CreateClientNetworkListTaskResponseData {
  s.TaskNo = &v
  return s
}

type CreateClientNetworkListTaskResponseHeader struct {
}

func (s CreateClientNetworkListTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateClientNetworkListTaskResponseHeader) GoString() string {
  return s.String()
}




type QueryAllBandwidthLimitTaskListServiceRequest struct {
  // {"en":"Domain:\n1. The maximum number of domain is 100 by default (you can contact technical support for adjustment);\n2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).","zh_CN":"域名:\n1.可传递域名数量上限默认为100个(可联系技术支持调整);\n2.自动过滤掉无效域名(如传递非法域名,会被过滤,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether to ignore task status.\n0: No, only return the configuration whose task status is enabled; default: 0;\n1: Yes, return whether the task status is open or closed.","zh_CN":"是否忽略任务状态。0:否,只返回任务状态为开启的配置;默认:0; 1:是,不论任务状态是开启还是关闭都返回。"}
  IgnoreTaskStatus *string `json:"ignoreTaskStatus,omitempty" xml:"ignoreTaskStatus,omitempty"`
  // {"en":"Whether the returned data contains all customer domain names involved in the task, the default is 0;\n0: no;\n1: Yes.","zh_CN":"返回数据是否包含任务涉及的所有客户域名。0:否;默认:0;1:是。"}
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

type QueryAllBandwidthLimitTaskListServiceRequestHeader struct {
}

func (s QueryAllBandwidthLimitTaskListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceRequestHeader) GoString() string {
  return s.String()
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

type QueryAllBandwidthLimitTaskListServiceResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
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
  // {"en":"domain","zh_CN":"客户域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Whether there is a configuration control task, 0: no, 1: yes","zh_CN":"是否有配置控制任务,0表示没有,1表示有"}
  IsExist *string `json:"isExist,omitempty" xml:"isExist,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
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
  // {"en":"taskName","zh_CN":"任务名称"}
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty" require:"true"`
  // {"en":"Domain name configuration task types: 1. Static bandwidth control task, 2. Bandwidth buyout task, 3. Flow buyout task, 4. Request number buyout task, 5. Redundant pool speed limit task, 6. Back-to-source task, 7. POP running high scheduling task (this kind of task is quite special, the domain name is a global quantity, so as long as there is configuration, it will be enabled by default), 8. IP ban task","zh_CN":"域名配置的任务类型:1静态带宽控制任务,2带宽买断任务,3流量买断任务,4请求数买断任务,5冗余池限速任务,6回源任务,7POP跑高调度任务(此种任务比较特殊,域名为全局量,所以只要有配置,就默认开启),8IP封禁任务"}
  TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty" require:"true"`
  // {"en":"Task status, 0: task is off, 1: task is on.","zh_CN":"任务状态,0表示任务关闭,1表示任务开启。"}
  TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty" require:"true"`
  // {"en":"Control strategy. 0 means squid default, 1 rejects, 2 when taskType=8, is ladder blocked, the rest is redirected ip, 3 when taskType=6, is backup source, the rest is speed limit = rejection + maximum download rate, 4 speed limit = rejection + timeout to disconnect, 5 does not process, controls each connection at minimum speed, does not process the excess part, 6 backups, 7 redirects domain name, 8 redirects URL","zh_CN":"控制策略。0表示squid默认,1拒绝,2当taskType=8时,为阶梯封禁,其余为重定向ip,3当taskType=6时,为主备回源,其余为限速=拒绝+最大下载速率,4限速=拒绝+超时断开连接,5不处理,按最小速率控制每个连接,超出部分不处理,6回源,7重定向域名,8重定向URL"}
  CtrlMode *string `json:"ctrlMode,omitempty" xml:"ctrlMode,omitempty" require:"true"`
  // {"en":"Bandwidth limit value, when taskType=1,2,5,6, in Mbps, when taskType=3, in G, when taskType=4, in MH (millions), when taskType=7, in -1, there is no bandwidth limit value, when taskType=8, in seconds.","zh_CN":"带宽限制值,当taskType=1,2,5,6时,单位为Mbps,当taskType=3,单位为G,当taskType=4,单位为MH(百万个),当taskType=7,为-1,表示没有带宽限制值,当taskType=8,单位为次。"}
  CtrlValue *string `json:"ctrlValue,omitempty" xml:"ctrlValue,omitempty" require:"true"`
  // {"en":"List of customer domains involved under the task. Values are only available when entering ContainOrichannelName=1.","zh_CN":"任务下涉及的客户域名列表。只有在入参containOrichannelName=1的时候有值。"}
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

type QueryAllBandwidthLimitTaskListServiceResponseHeader struct {
}

func (s QueryAllBandwidthLimitTaskListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseHeader) GoString() string {
  return s.String()
}




