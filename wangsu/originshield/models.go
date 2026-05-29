package originshield

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryOriginShieldsRequest struct {
}

func (s QueryOriginShieldsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsRequest) GoString() string {
  return s.String()
}

type QueryOriginShieldsRequestHeader struct {
}

func (s QueryOriginShieldsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginShieldsPaths struct {
}

func (s QueryOriginShieldsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsPaths) GoString() string {
  return s.String()
}

type QueryOriginShieldsParameters struct {
  // {"en":"page number","zh_CN":"当前页数"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"page size","zh_CN":"每页记录数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Origin shield code list, separated by commas.","zh_CN":"源站防护盾编码列表,多个以英文逗号分隔"}
  RouteMapCodes *string `json:"routeMapCodes,omitempty" xml:"routeMapCodes,omitempty"`
  // {"en":"Origin shield name list, separated by commas. Supports fuzzy search for a single input, and exact search for multiple inputs.","zh_CN":"源站防护盾名称列表,多个以英文逗号分隔。单个入参时支持模糊搜索，多个入参时为精确搜索。"}
  OriginShieldNames *string `json:"originShieldNames,omitempty" xml:"originShieldNames,omitempty"`
  // {"en":"Origin shield ID list, separated by commas. Supports fuzzy search for a single input, and exact search for multiple inputs.","zh_CN":"源站防护盾ID列表,多个以英文逗号分隔。单个入参时支持模糊搜索，多个入参时为精确搜索。"}
  OriginShieldIds *string `json:"originShieldIds,omitempty" xml:"originShieldIds,omitempty"`
  // {"en":"Origin shield status","zh_CN":"源站防护盾状态","exampleValue":"enable,disable"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"IP confirmation status","zh_CN":"IP确认状态","exampleValue":"pending,confirmed"}
  IpConfirmationStatus *string `json:"ipConfirmationStatus,omitempty" xml:"ipConfirmationStatus,omitempty"`
}

func (s QueryOriginShieldsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsParameters) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldsParameters) SetPageNumber(v int) *QueryOriginShieldsParameters {
  s.PageNumber = &v
  return s
}

func (s *QueryOriginShieldsParameters) SetPageSize(v int) *QueryOriginShieldsParameters {
  s.PageSize = &v
  return s
}

func (s *QueryOriginShieldsParameters) SetRouteMapCodes(v string) *QueryOriginShieldsParameters {
  s.RouteMapCodes = &v
  return s
}

func (s *QueryOriginShieldsParameters) SetOriginShieldNames(v string) *QueryOriginShieldsParameters {
  s.OriginShieldNames = &v
  return s
}

func (s *QueryOriginShieldsParameters) SetOriginShieldIds(v string) *QueryOriginShieldsParameters {
  s.OriginShieldIds = &v
  return s
}

func (s *QueryOriginShieldsParameters) SetStatus(v string) *QueryOriginShieldsParameters {
  s.Status = &v
  return s
}

func (s *QueryOriginShieldsParameters) SetIpConfirmationStatus(v string) *QueryOriginShieldsParameters {
  s.IpConfirmationStatus = &v
  return s
}

type QueryOriginShieldsResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Query result","zh_CN":"查询结果"}
  Data *QueryOriginShieldsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryOriginShieldsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldsResponse) SetCode(v string) *QueryOriginShieldsResponse {
  s.Code = &v
  return s
}

func (s *QueryOriginShieldsResponse) SetMessage(v string) *QueryOriginShieldsResponse {
  s.Message = &v
  return s
}

func (s *QueryOriginShieldsResponse) SetData(v *QueryOriginShieldsResponseData) *QueryOriginShieldsResponse {
  s.Data = v
  return s
}

type QueryOriginShieldsResponseData struct {
  // {"en":"","zh_CN":""}
  PageInfo *QueryOriginShieldsResponseDataPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
  // {"en":"origin shield list","zh_CN":"源站防护盾列表"}
  OriginShields []*QueryOriginShieldsResponseDataOriginShields `json:"originShields,omitempty" xml:"originShields,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsResponseData) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldsResponseData) SetPageInfo(v *QueryOriginShieldsResponseDataPageInfo) *QueryOriginShieldsResponseData {
  s.PageInfo = v
  return s
}

func (s *QueryOriginShieldsResponseData) SetOriginShields(v []*QueryOriginShieldsResponseDataOriginShields) *QueryOriginShieldsResponseData {
  s.OriginShields = v
  return s
}

type QueryOriginShieldsResponseDataPageInfo struct {
  // {"en":"total record count","zh_CN":"总记录数"}
  TotalNumber *int `json:"totalNumber,omitempty" xml:"totalNumber,omitempty" require:"true"`
  // {"en":"total page count","zh_CN":"总页数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
  // {"en":"current page number","zh_CN":"当前页数"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"page size","zh_CN":"每页记录数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
}

func (s QueryOriginShieldsResponseDataPageInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsResponseDataPageInfo) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldsResponseDataPageInfo) SetTotalNumber(v int) *QueryOriginShieldsResponseDataPageInfo {
  s.TotalNumber = &v
  return s
}

func (s *QueryOriginShieldsResponseDataPageInfo) SetTotalPageNumber(v int) *QueryOriginShieldsResponseDataPageInfo {
  s.TotalPageNumber = &v
  return s
}

func (s *QueryOriginShieldsResponseDataPageInfo) SetPageNumber(v int) *QueryOriginShieldsResponseDataPageInfo {
  s.PageNumber = &v
  return s
}

func (s *QueryOriginShieldsResponseDataPageInfo) SetPageSize(v int) *QueryOriginShieldsResponseDataPageInfo {
  s.PageSize = &v
  return s
}

type QueryOriginShieldsResponseDataOriginShields struct     {
  // {"en":"origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
  // {"en":"origin shield name","zh_CN":"源站防护盾名称"}
  OriginShieldName *string `json:"originShieldName,omitempty" xml:"originShieldName,omitempty" require:"true"`
  // {"en":"origin shield code","zh_CN":"源站防护盾编码"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
  // {"en":"origin shield status","zh_CN":"源站防护盾编状态","exampleValue":"enable,disable"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"origin shield description","zh_CN":"源站防护盾编描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"create time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"update time","zh_CN":"修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"IP confirmation status","zh_CN":"IP确认状态","exampleValue":"pending,confirmed"}
  IpConfirmationStatus *string `json:"ipConfirmationStatus,omitempty" xml:"ipConfirmationStatus,omitempty" require:"true"`
  // {"en":"related origin and hostnames info list","zh_CN":"关联源站防护盾的源站与域名信息列表"}
  RelatedInfos []*QueryOriginShieldsResponseDataOriginShieldsRelatedInfos `json:"relatedInfos,omitempty" xml:"relatedInfos,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldsResponseDataOriginShields) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsResponseDataOriginShields) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetOriginShieldId(v int) *QueryOriginShieldsResponseDataOriginShields {
  s.OriginShieldId = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetOriginShieldName(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.OriginShieldName = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetRouteMapCode(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.RouteMapCode = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetStatus(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.Status = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetDescription(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.Description = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetCreateTime(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.CreateTime = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetUpdateTime(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.UpdateTime = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetIpConfirmationStatus(v string) *QueryOriginShieldsResponseDataOriginShields {
  s.IpConfirmationStatus = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShields) SetRelatedInfos(v []*QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) *QueryOriginShieldsResponseDataOriginShields {
  s.RelatedInfos = v
  return s
}

type QueryOriginShieldsResponseDataOriginShieldsRelatedInfos struct     {
  // {"en":"origin name","zh_CN":"源站名称"}
  OriginName *string `json:"originName,omitempty" xml:"originName,omitempty" require:"true"`
  // {"en":"property ID","zh_CN":"项目ID"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"property name","zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"property version","zh_CN":"项目版本"}
  PropertyVersion *string `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"deploy environment","zh_CN":"部署环境","exampleValue":"production,staging"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"hostname list","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) SetOriginName(v string) *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos {
  s.OriginName = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) SetPropertyId(v string) *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos {
  s.PropertyId = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) SetPropertyName(v string) *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos {
  s.PropertyName = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) SetPropertyVersion(v string) *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos {
  s.PropertyVersion = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) SetTarget(v string) *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos {
  s.Target = &v
  return s
}

func (s *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos) SetHostnames(v []*string) *QueryOriginShieldsResponseDataOriginShieldsRelatedInfos {
  s.Hostnames = v
  return s
}

type QueryOriginShieldsResponseHeader struct {
}

func (s QueryOriginShieldsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldsResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginShieldPendingIpSegmentsRequest struct {
}

func (s QueryOriginShieldPendingIpSegmentsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsRequest) GoString() string {
  return s.String()
}

type QueryOriginShieldPendingIpSegmentsRequestHeader struct {
}

func (s QueryOriginShieldPendingIpSegmentsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginShieldPendingIpSegmentsPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s QueryOriginShieldPendingIpSegmentsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsPaths) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPendingIpSegmentsPaths) SetOriginShieldId(v int) *QueryOriginShieldPendingIpSegmentsPaths {
  s.OriginShieldId = &v
  return s
}

type QueryOriginShieldPendingIpSegmentsParameters struct {
  // {"en":"action","zh_CN":"操作","exampleValue":"add,delete"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en":"IP type","zh_CN":"IP类型","exampleValue":"IPv4,IPv6"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
  // {"en":"page number","zh_CN":"当前页数"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"page size","zh_CN":"每页记录数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
}

func (s QueryOriginShieldPendingIpSegmentsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsParameters) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPendingIpSegmentsParameters) SetAction(v string) *QueryOriginShieldPendingIpSegmentsParameters {
  s.Action = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsParameters) SetIpType(v string) *QueryOriginShieldPendingIpSegmentsParameters {
  s.IpType = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsParameters) SetPageNumber(v int) *QueryOriginShieldPendingIpSegmentsParameters {
  s.PageNumber = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsParameters) SetPageSize(v int) *QueryOriginShieldPendingIpSegmentsParameters {
  s.PageSize = &v
  return s
}

type QueryOriginShieldPendingIpSegmentsResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Query result","zh_CN":"查询结果"}
  Data *QueryOriginShieldPendingIpSegmentsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryOriginShieldPendingIpSegmentsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPendingIpSegmentsResponse) SetCode(v string) *QueryOriginShieldPendingIpSegmentsResponse {
  s.Code = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponse) SetMessage(v string) *QueryOriginShieldPendingIpSegmentsResponse {
  s.Message = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponse) SetData(v *QueryOriginShieldPendingIpSegmentsResponseData) *QueryOriginShieldPendingIpSegmentsResponse {
  s.Data = v
  return s
}

type QueryOriginShieldPendingIpSegmentsResponseData struct {
  // {"en":"","zh_CN":""}
  PageInfo *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
  // {"en":"IP segment list","zh_CN":"IP段列表"}
  IpSegments []*QueryOriginShieldPendingIpSegmentsResponseDataIpSegments `json:"ipSegments,omitempty" xml:"ipSegments,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldPendingIpSegmentsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsResponseData) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPendingIpSegmentsResponseData) SetPageInfo(v *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) *QueryOriginShieldPendingIpSegmentsResponseData {
  s.PageInfo = v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseData) SetIpSegments(v []*QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) *QueryOriginShieldPendingIpSegmentsResponseData {
  s.IpSegments = v
  return s
}

type QueryOriginShieldPendingIpSegmentsResponseDataPageInfo struct {
  // {"en":"total record count","zh_CN":"总记录数"}
  TotalNumber *int `json:"totalNumber,omitempty" xml:"totalNumber,omitempty" require:"true"`
  // {"en":"total page count","zh_CN":"总页数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
  // {"en":"current page number","zh_CN":"当前页数"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"page size","zh_CN":"每页记录数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
}

func (s QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) SetTotalNumber(v int) *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo {
  s.TotalNumber = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) SetTotalPageNumber(v int) *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo {
  s.TotalPageNumber = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) SetPageNumber(v int) *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo {
  s.PageNumber = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo) SetPageSize(v int) *QueryOriginShieldPendingIpSegmentsResponseDataPageInfo {
  s.PageSize = &v
  return s
}

type QueryOriginShieldPendingIpSegmentsResponseDataIpSegments struct     {
  // {"en":"record ID","zh_CN":"记录ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"action","zh_CN":"操作","exampleValue":"add,delete"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"IP type","zh_CN":"IP类型","exampleValue":"IPv4,IPv6"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty" require:"true"`
  // {"en":"IP segment","zh_CN":"IP段"}
  IpSegment *string `json:"ipSegment,omitempty" xml:"ipSegment,omitempty" require:"true"`
  // {"en":"notify time","zh_CN":"通知时间"}
  NotifyTime *string `json:"notifyTime,omitempty" xml:"notifyTime,omitempty" require:"true"`
}

func (s QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) SetRecordId(v int) *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments {
  s.RecordId = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) SetAction(v string) *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments {
  s.Action = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) SetIpType(v string) *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments {
  s.IpType = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) SetIpSegment(v string) *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments {
  s.IpSegment = &v
  return s
}

func (s *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments) SetNotifyTime(v string) *QueryOriginShieldPendingIpSegmentsResponseDataIpSegments {
  s.NotifyTime = &v
  return s
}

type QueryOriginShieldPendingIpSegmentsResponseHeader struct {
}

func (s QueryOriginShieldPendingIpSegmentsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPendingIpSegmentsResponseHeader) GoString() string {
  return s.String()
}




type QueryAvailableOriginShieldsRequest struct {
}

func (s QueryAvailableOriginShieldsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsRequest) GoString() string {
  return s.String()
}

type QueryAvailableOriginShieldsRequestHeader struct {
}

func (s QueryAvailableOriginShieldsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsRequestHeader) GoString() string {
  return s.String()
}

type QueryAvailableOriginShieldsPaths struct {
}

func (s QueryAvailableOriginShieldsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsPaths) GoString() string {
  return s.String()
}

type QueryAvailableOriginShieldsParameters struct {
  // {"en":"service type","zh_CN":"商品","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
}

func (s QueryAvailableOriginShieldsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsParameters) GoString() string {
  return s.String()
}

func (s *QueryAvailableOriginShieldsParameters) SetServiceType(v string) *QueryAvailableOriginShieldsParameters {
  s.ServiceType = &v
  return s
}

func (s *QueryAvailableOriginShieldsParameters) SetPropertyId(v int) *QueryAvailableOriginShieldsParameters {
  s.PropertyId = &v
  return s
}

type QueryAvailableOriginShieldsResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Query result","zh_CN":"查询结果"}
  Data []*QueryAvailableOriginShieldsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAvailableOriginShieldsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsResponse) GoString() string {
  return s.String()
}

func (s *QueryAvailableOriginShieldsResponse) SetCode(v string) *QueryAvailableOriginShieldsResponse {
  s.Code = &v
  return s
}

func (s *QueryAvailableOriginShieldsResponse) SetMessage(v string) *QueryAvailableOriginShieldsResponse {
  s.Message = &v
  return s
}

func (s *QueryAvailableOriginShieldsResponse) SetData(v []*QueryAvailableOriginShieldsResponseData) *QueryAvailableOriginShieldsResponse {
  s.Data = v
  return s
}

type QueryAvailableOriginShieldsResponseData struct     {
  // {"en":"origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
  // {"en":"origin shield name","zh_CN":"源站防护盾名称"}
  OriginShieldName *string `json:"originShieldName,omitempty" xml:"originShieldName,omitempty" require:"true"`
  // {"en":"origin shield code","zh_CN":"源站防护盾编码"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
}

func (s QueryAvailableOriginShieldsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsResponseData) GoString() string {
  return s.String()
}

func (s *QueryAvailableOriginShieldsResponseData) SetOriginShieldId(v int) *QueryAvailableOriginShieldsResponseData {
  s.OriginShieldId = &v
  return s
}

func (s *QueryAvailableOriginShieldsResponseData) SetOriginShieldName(v string) *QueryAvailableOriginShieldsResponseData {
  s.OriginShieldName = &v
  return s
}

func (s *QueryAvailableOriginShieldsResponseData) SetRouteMapCode(v string) *QueryAvailableOriginShieldsResponseData {
  s.RouteMapCode = &v
  return s
}

type QueryAvailableOriginShieldsResponseHeader struct {
}

func (s QueryAvailableOriginShieldsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAvailableOriginShieldsResponseHeader) GoString() string {
  return s.String()
}




type EnableOriginShieldRequest struct {
}

func (s EnableOriginShieldRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableOriginShieldRequest) GoString() string {
  return s.String()
}

type EnableOriginShieldRequestHeader struct {
}

func (s EnableOriginShieldRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOriginShieldRequestHeader) GoString() string {
  return s.String()
}

type EnableOriginShieldPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s EnableOriginShieldPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableOriginShieldPaths) GoString() string {
  return s.String()
}

func (s *EnableOriginShieldPaths) SetOriginShieldId(v int) *EnableOriginShieldPaths {
  s.OriginShieldId = &v
  return s
}

type EnableOriginShieldParameters struct {
}

func (s EnableOriginShieldParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableOriginShieldParameters) GoString() string {
  return s.String()
}

type EnableOriginShieldResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EnableOriginShieldResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableOriginShieldResponse) GoString() string {
  return s.String()
}

func (s *EnableOriginShieldResponse) SetCode(v string) *EnableOriginShieldResponse {
  s.Code = &v
  return s
}

func (s *EnableOriginShieldResponse) SetMessage(v string) *EnableOriginShieldResponse {
  s.Message = &v
  return s
}

type EnableOriginShieldResponseHeader struct {
}

func (s EnableOriginShieldResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOriginShieldResponseHeader) GoString() string {
  return s.String()
}




type DisableOriginShieldRequest struct {
}

func (s DisableOriginShieldRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableOriginShieldRequest) GoString() string {
  return s.String()
}

type DisableOriginShieldRequestHeader struct {
}

func (s DisableOriginShieldRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableOriginShieldRequestHeader) GoString() string {
  return s.String()
}

type DisableOriginShieldPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s DisableOriginShieldPaths) String() string {
  return tea.Prettify(s)
}

func (s DisableOriginShieldPaths) GoString() string {
  return s.String()
}

func (s *DisableOriginShieldPaths) SetOriginShieldId(v int) *DisableOriginShieldPaths {
  s.OriginShieldId = &v
  return s
}

type DisableOriginShieldParameters struct {
}

func (s DisableOriginShieldParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableOriginShieldParameters) GoString() string {
  return s.String()
}

type DisableOriginShieldResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DisableOriginShieldResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableOriginShieldResponse) GoString() string {
  return s.String()
}

func (s *DisableOriginShieldResponse) SetCode(v string) *DisableOriginShieldResponse {
  s.Code = &v
  return s
}

func (s *DisableOriginShieldResponse) SetMessage(v string) *DisableOriginShieldResponse {
  s.Message = &v
  return s
}

type DisableOriginShieldResponseHeader struct {
}

func (s DisableOriginShieldResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableOriginShieldResponseHeader) GoString() string {
  return s.String()
}




type QueryShieldRouteMapsRequest struct {
}

func (s QueryShieldRouteMapsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsRequest) GoString() string {
  return s.String()
}

type QueryShieldRouteMapsRequestHeader struct {
}

func (s QueryShieldRouteMapsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsRequestHeader) GoString() string {
  return s.String()
}

type QueryShieldRouteMapsPaths struct {
}

func (s QueryShieldRouteMapsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsPaths) GoString() string {
  return s.String()
}

type QueryShieldRouteMapsParameters struct {
}

func (s QueryShieldRouteMapsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsParameters) GoString() string {
  return s.String()
}

type QueryShieldRouteMapsResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"List of available route maps","zh_CN":"可用回源路由列表"}
  Data []*QueryShieldRouteMapsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryShieldRouteMapsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsResponse) GoString() string {
  return s.String()
}

func (s *QueryShieldRouteMapsResponse) SetCode(v string) *QueryShieldRouteMapsResponse {
  s.Code = &v
  return s
}

func (s *QueryShieldRouteMapsResponse) SetMessage(v string) *QueryShieldRouteMapsResponse {
  s.Message = &v
  return s
}

func (s *QueryShieldRouteMapsResponse) SetData(v []*QueryShieldRouteMapsResponseData) *QueryShieldRouteMapsResponse {
  s.Data = v
  return s
}

type QueryShieldRouteMapsResponseData struct     {
  // {"en":"Route map code.","zh_CN":"回源路由标识"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
  // {"en":"Map type. optional values:[public,custom]","zh_CN":"回源路由分类，可选值:[public,custom]","exampleValue":"public,custom"}
  MapType *string `json:"mapType,omitempty" xml:"mapType,omitempty" require:"true"`
}

func (s QueryShieldRouteMapsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsResponseData) GoString() string {
  return s.String()
}

func (s *QueryShieldRouteMapsResponseData) SetRouteMapCode(v string) *QueryShieldRouteMapsResponseData {
  s.RouteMapCode = &v
  return s
}

func (s *QueryShieldRouteMapsResponseData) SetMapType(v string) *QueryShieldRouteMapsResponseData {
  s.MapType = &v
  return s
}

type QueryShieldRouteMapsResponseHeader struct {
}

func (s QueryShieldRouteMapsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryShieldRouteMapsResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginShieldActiveIpSegmentsRequest struct {
}

func (s QueryOriginShieldActiveIpSegmentsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsRequest) GoString() string {
  return s.String()
}

type QueryOriginShieldActiveIpSegmentsRequestHeader struct {
}

func (s QueryOriginShieldActiveIpSegmentsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginShieldActiveIpSegmentsPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s QueryOriginShieldActiveIpSegmentsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsPaths) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldActiveIpSegmentsPaths) SetOriginShieldId(v int) *QueryOriginShieldActiveIpSegmentsPaths {
  s.OriginShieldId = &v
  return s
}

type QueryOriginShieldActiveIpSegmentsParameters struct {
}

func (s QueryOriginShieldActiveIpSegmentsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsParameters) GoString() string {
  return s.String()
}

type QueryOriginShieldActiveIpSegmentsResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Query result","zh_CN":"查询结果"}
  Data []*QueryOriginShieldActiveIpSegmentsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldActiveIpSegmentsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldActiveIpSegmentsResponse) SetCode(v string) *QueryOriginShieldActiveIpSegmentsResponse {
  s.Code = &v
  return s
}

func (s *QueryOriginShieldActiveIpSegmentsResponse) SetMessage(v string) *QueryOriginShieldActiveIpSegmentsResponse {
  s.Message = &v
  return s
}

func (s *QueryOriginShieldActiveIpSegmentsResponse) SetData(v []*QueryOriginShieldActiveIpSegmentsResponseData) *QueryOriginShieldActiveIpSegmentsResponse {
  s.Data = v
  return s
}

type QueryOriginShieldActiveIpSegmentsResponseData struct     {
  // {"en":"record ID","zh_CN":"记录ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"IP type","zh_CN":"IP类型","exampleValue":"IPv4,IPv6"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty" require:"true"`
  // {"en":"IP segment","zh_CN":"IP段"}
  IpSegment *string `json:"ipSegment,omitempty" xml:"ipSegment,omitempty" require:"true"`
  // {"en":"active time","zh_CN":"生效时间"}
  ActiveTime *string `json:"activeTime,omitempty" xml:"activeTime,omitempty" require:"true"`
}

func (s QueryOriginShieldActiveIpSegmentsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsResponseData) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldActiveIpSegmentsResponseData) SetRecordId(v int) *QueryOriginShieldActiveIpSegmentsResponseData {
  s.RecordId = &v
  return s
}

func (s *QueryOriginShieldActiveIpSegmentsResponseData) SetIpType(v string) *QueryOriginShieldActiveIpSegmentsResponseData {
  s.IpType = &v
  return s
}

func (s *QueryOriginShieldActiveIpSegmentsResponseData) SetIpSegment(v string) *QueryOriginShieldActiveIpSegmentsResponseData {
  s.IpSegment = &v
  return s
}

func (s *QueryOriginShieldActiveIpSegmentsResponseData) SetActiveTime(v string) *QueryOriginShieldActiveIpSegmentsResponseData {
  s.ActiveTime = &v
  return s
}

type QueryOriginShieldActiveIpSegmentsResponseHeader struct {
}

func (s QueryOriginShieldActiveIpSegmentsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldActiveIpSegmentsResponseHeader) GoString() string {
  return s.String()
}




type CreateOriginShieldRequest struct {
  // {"en":"Route map code","zh_CN":"源站防护盾编码"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
  // {"en":"Origin shield name","zh_CN":"源站防护盾名称"}
  OriginShieldName *string `json:"originShieldName,omitempty" xml:"originShieldName,omitempty" require:"true"`
  // {"en":"Origin shield description","zh_CN":"源站防护盾描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"IP Confirmation status","zh_CN":"IP确认状态","exampleValue":"pending,confirmed"}
  IpConfirmationStatus *string `json:"ipConfirmationStatus,omitempty" xml:"ipConfirmationStatus,omitempty" require:"true"`
  // {"en":"Service type","zh_CN":"商品","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
}

func (s CreateOriginShieldRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldRequest) GoString() string {
  return s.String()
}

func (s *CreateOriginShieldRequest) SetRouteMapCode(v string) *CreateOriginShieldRequest {
  s.RouteMapCode = &v
  return s
}

func (s *CreateOriginShieldRequest) SetOriginShieldName(v string) *CreateOriginShieldRequest {
  s.OriginShieldName = &v
  return s
}

func (s *CreateOriginShieldRequest) SetDescription(v string) *CreateOriginShieldRequest {
  s.Description = &v
  return s
}

func (s *CreateOriginShieldRequest) SetIpConfirmationStatus(v string) *CreateOriginShieldRequest {
  s.IpConfirmationStatus = &v
  return s
}

func (s *CreateOriginShieldRequest) SetServiceType(v string) *CreateOriginShieldRequest {
  s.ServiceType = &v
  return s
}

type CreateOriginShieldRequestHeader struct {
}

func (s CreateOriginShieldRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldRequestHeader) GoString() string {
  return s.String()
}

type CreateOriginShieldPaths struct {
}

func (s CreateOriginShieldPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldPaths) GoString() string {
  return s.String()
}

type CreateOriginShieldParameters struct {
}

func (s CreateOriginShieldParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldParameters) GoString() string {
  return s.String()
}

type CreateOriginShieldResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Origin shield information","zh_CN":"源站防护盾信息"}
  Data *CreateOriginShieldResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateOriginShieldResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldResponse) GoString() string {
  return s.String()
}

func (s *CreateOriginShieldResponse) SetCode(v string) *CreateOriginShieldResponse {
  s.Code = &v
  return s
}

func (s *CreateOriginShieldResponse) SetMessage(v string) *CreateOriginShieldResponse {
  s.Message = &v
  return s
}

func (s *CreateOriginShieldResponse) SetData(v *CreateOriginShieldResponseData) *CreateOriginShieldResponse {
  s.Data = v
  return s
}

type CreateOriginShieldResponseData struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
  // {"en":"Origin shield name","zh_CN":"源站防护盾名称"}
  OriginShieldName *string `json:"originShieldName,omitempty" xml:"originShieldName,omitempty" require:"true"`
  // {"en":"Origin shield code","zh_CN":"源站防护盾编码"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
}

func (s CreateOriginShieldResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldResponseData) GoString() string {
  return s.String()
}

func (s *CreateOriginShieldResponseData) SetOriginShieldId(v int) *CreateOriginShieldResponseData {
  s.OriginShieldId = &v
  return s
}

func (s *CreateOriginShieldResponseData) SetOriginShieldName(v string) *CreateOriginShieldResponseData {
  s.OriginShieldName = &v
  return s
}

func (s *CreateOriginShieldResponseData) SetRouteMapCode(v string) *CreateOriginShieldResponseData {
  s.RouteMapCode = &v
  return s
}

type CreateOriginShieldResponseHeader struct {
}

func (s CreateOriginShieldResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateOriginShieldResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginShieldRequest struct {
}

func (s QueryOriginShieldRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRequest) GoString() string {
  return s.String()
}

type QueryOriginShieldRequestHeader struct {
}

func (s QueryOriginShieldRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginShieldPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s QueryOriginShieldPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldPaths) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldPaths) SetOriginShieldId(v int) *QueryOriginShieldPaths {
  s.OriginShieldId = &v
  return s
}

type QueryOriginShieldParameters struct {
}

func (s QueryOriginShieldParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldParameters) GoString() string {
  return s.String()
}

type QueryOriginShieldResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Origin shield information","zh_CN":"源站防护盾信息"}
  Data *QueryOriginShieldResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryOriginShieldResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldResponse) SetCode(v string) *QueryOriginShieldResponse {
  s.Code = &v
  return s
}

func (s *QueryOriginShieldResponse) SetMessage(v string) *QueryOriginShieldResponse {
  s.Message = &v
  return s
}

func (s *QueryOriginShieldResponse) SetData(v *QueryOriginShieldResponseData) *QueryOriginShieldResponse {
  s.Data = v
  return s
}

type QueryOriginShieldResponseData struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
  // {"en":"Origin shield name","zh_CN":"源站防护盾名称"}
  OriginShieldName *string `json:"originShieldName,omitempty" xml:"originShieldName,omitempty" require:"true"`
  // {"en":"Origin shield code","zh_CN":"源站防护盾编码"}
  RouteMapCode *string `json:"routeMapCode,omitempty" xml:"routeMapCode,omitempty" require:"true"`
  // {"en":"Status","zh_CN":"状态","exampleValue":"enable,disable"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Origin shield description","zh_CN":"源站防护盾描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Create time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update time","zh_CN":"创建时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s QueryOriginShieldResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldResponseData) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldResponseData) SetOriginShieldId(v int) *QueryOriginShieldResponseData {
  s.OriginShieldId = &v
  return s
}

func (s *QueryOriginShieldResponseData) SetOriginShieldName(v string) *QueryOriginShieldResponseData {
  s.OriginShieldName = &v
  return s
}

func (s *QueryOriginShieldResponseData) SetRouteMapCode(v string) *QueryOriginShieldResponseData {
  s.RouteMapCode = &v
  return s
}

func (s *QueryOriginShieldResponseData) SetStatus(v string) *QueryOriginShieldResponseData {
  s.Status = &v
  return s
}

func (s *QueryOriginShieldResponseData) SetDescription(v string) *QueryOriginShieldResponseData {
  s.Description = &v
  return s
}

func (s *QueryOriginShieldResponseData) SetCreateTime(v string) *QueryOriginShieldResponseData {
  s.CreateTime = &v
  return s
}

func (s *QueryOriginShieldResponseData) SetUpdateTime(v string) *QueryOriginShieldResponseData {
  s.UpdateTime = &v
  return s
}

type QueryOriginShieldResponseHeader struct {
}

func (s QueryOriginShieldResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldResponseHeader) GoString() string {
  return s.String()
}




type UpdateOriginShieldRequest struct {
  // {"en":"Origin shield name","zh_CN":"源站防护盾名称"}
  OriginShieldName *string `json:"originShieldName,omitempty" xml:"originShieldName,omitempty" require:"true"`
  // {"en":"Origin shield description","zh_CN":"源站防护盾描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateOriginShieldRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateOriginShieldRequest) GoString() string {
  return s.String()
}

func (s *UpdateOriginShieldRequest) SetOriginShieldName(v string) *UpdateOriginShieldRequest {
  s.OriginShieldName = &v
  return s
}

func (s *UpdateOriginShieldRequest) SetDescription(v string) *UpdateOriginShieldRequest {
  s.Description = &v
  return s
}

type UpdateOriginShieldRequestHeader struct {
}

func (s UpdateOriginShieldRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateOriginShieldRequestHeader) GoString() string {
  return s.String()
}

type UpdateOriginShieldPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s UpdateOriginShieldPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateOriginShieldPaths) GoString() string {
  return s.String()
}

func (s *UpdateOriginShieldPaths) SetOriginShieldId(v int) *UpdateOriginShieldPaths {
  s.OriginShieldId = &v
  return s
}

type UpdateOriginShieldParameters struct {
}

func (s UpdateOriginShieldParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateOriginShieldParameters) GoString() string {
  return s.String()
}

type UpdateOriginShieldResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateOriginShieldResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateOriginShieldResponse) GoString() string {
  return s.String()
}

func (s *UpdateOriginShieldResponse) SetCode(v string) *UpdateOriginShieldResponse {
  s.Code = &v
  return s
}

func (s *UpdateOriginShieldResponse) SetMessage(v string) *UpdateOriginShieldResponse {
  s.Message = &v
  return s
}

type UpdateOriginShieldResponseHeader struct {
}

func (s UpdateOriginShieldResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateOriginShieldResponseHeader) GoString() string {
  return s.String()
}




type ConfirmOriginShieldIpSegmentChangeRequest struct {
  // {"en":"Pending IP segment record ID list","zh_CN":"待确认IP段记录ID列表"}
  RecordIds []*int `json:"recordIds,omitempty" xml:"recordIds,omitempty" require:"true" type:"Repeated"`
}

func (s ConfirmOriginShieldIpSegmentChangeRequest) String() string {
  return tea.Prettify(s)
}

func (s ConfirmOriginShieldIpSegmentChangeRequest) GoString() string {
  return s.String()
}

func (s *ConfirmOriginShieldIpSegmentChangeRequest) SetRecordIds(v []*int) *ConfirmOriginShieldIpSegmentChangeRequest {
  s.RecordIds = v
  return s
}

type ConfirmOriginShieldIpSegmentChangeRequestHeader struct {
}

func (s ConfirmOriginShieldIpSegmentChangeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ConfirmOriginShieldIpSegmentChangeRequestHeader) GoString() string {
  return s.String()
}

type ConfirmOriginShieldIpSegmentChangePaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s ConfirmOriginShieldIpSegmentChangePaths) String() string {
  return tea.Prettify(s)
}

func (s ConfirmOriginShieldIpSegmentChangePaths) GoString() string {
  return s.String()
}

func (s *ConfirmOriginShieldIpSegmentChangePaths) SetOriginShieldId(v int) *ConfirmOriginShieldIpSegmentChangePaths {
  s.OriginShieldId = &v
  return s
}

type ConfirmOriginShieldIpSegmentChangeParameters struct {
}

func (s ConfirmOriginShieldIpSegmentChangeParameters) String() string {
  return tea.Prettify(s)
}

func (s ConfirmOriginShieldIpSegmentChangeParameters) GoString() string {
  return s.String()
}

type ConfirmOriginShieldIpSegmentChangeResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ConfirmOriginShieldIpSegmentChangeResponse) String() string {
  return tea.Prettify(s)
}

func (s ConfirmOriginShieldIpSegmentChangeResponse) GoString() string {
  return s.String()
}

func (s *ConfirmOriginShieldIpSegmentChangeResponse) SetCode(v string) *ConfirmOriginShieldIpSegmentChangeResponse {
  s.Code = &v
  return s
}

func (s *ConfirmOriginShieldIpSegmentChangeResponse) SetMessage(v string) *ConfirmOriginShieldIpSegmentChangeResponse {
  s.Message = &v
  return s
}

type ConfirmOriginShieldIpSegmentChangeResponseHeader struct {
}

func (s ConfirmOriginShieldIpSegmentChangeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ConfirmOriginShieldIpSegmentChangeResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginShieldRelatedInfosRequest struct {
}

func (s QueryOriginShieldRelatedInfosRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosRequest) GoString() string {
  return s.String()
}

type QueryOriginShieldRelatedInfosRequestHeader struct {
}

func (s QueryOriginShieldRelatedInfosRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginShieldRelatedInfosPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s QueryOriginShieldRelatedInfosPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosPaths) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldRelatedInfosPaths) SetOriginShieldId(v int) *QueryOriginShieldRelatedInfosPaths {
  s.OriginShieldId = &v
  return s
}

type QueryOriginShieldRelatedInfosParameters struct {
  // {"en":"deploy environment","zh_CN":"部署环境","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s QueryOriginShieldRelatedInfosParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosParameters) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldRelatedInfosParameters) SetTarget(v string) *QueryOriginShieldRelatedInfosParameters {
  s.Target = &v
  return s
}

type QueryOriginShieldRelatedInfosResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Related property and origin info list","zh_CN":"关联的项目与源站信息列表"}
  Data []*QueryOriginShieldRelatedInfosResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldRelatedInfosResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldRelatedInfosResponse) SetCode(v string) *QueryOriginShieldRelatedInfosResponse {
  s.Code = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponse) SetMessage(v string) *QueryOriginShieldRelatedInfosResponse {
  s.Message = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponse) SetData(v []*QueryOriginShieldRelatedInfosResponseData) *QueryOriginShieldRelatedInfosResponse {
  s.Data = v
  return s
}

type QueryOriginShieldRelatedInfosResponseData struct     {
  // {"en":"origin name","zh_CN":"源站名称"}
  OriginName *string `json:"originName,omitempty" xml:"originName,omitempty" require:"true"`
  // {"en":"property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"property name","zh_CN":"项目名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"property version","zh_CN":"项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"deploy environment","zh_CN":"部署环境","exampleValue":"production,staging"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"hostname list","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginShieldRelatedInfosResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosResponseData) GoString() string {
  return s.String()
}

func (s *QueryOriginShieldRelatedInfosResponseData) SetOriginName(v string) *QueryOriginShieldRelatedInfosResponseData {
  s.OriginName = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponseData) SetPropertyId(v int) *QueryOriginShieldRelatedInfosResponseData {
  s.PropertyId = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponseData) SetPropertyName(v string) *QueryOriginShieldRelatedInfosResponseData {
  s.PropertyName = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponseData) SetPropertyVersion(v int) *QueryOriginShieldRelatedInfosResponseData {
  s.PropertyVersion = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponseData) SetTarget(v string) *QueryOriginShieldRelatedInfosResponseData {
  s.Target = &v
  return s
}

func (s *QueryOriginShieldRelatedInfosResponseData) SetHostnames(v []*string) *QueryOriginShieldRelatedInfosResponseData {
  s.Hostnames = v
  return s
}

type QueryOriginShieldRelatedInfosResponseHeader struct {
}

func (s QueryOriginShieldRelatedInfosResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginShieldRelatedInfosResponseHeader) GoString() string {
  return s.String()
}




type DeleteOriginShieldRequest struct {
}

func (s DeleteOriginShieldRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteOriginShieldRequest) GoString() string {
  return s.String()
}

type DeleteOriginShieldRequestHeader struct {
}

func (s DeleteOriginShieldRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteOriginShieldRequestHeader) GoString() string {
  return s.String()
}

type DeleteOriginShieldPaths struct {
  // {"en":"Origin shield ID","zh_CN":"源站防护盾ID"}
  OriginShieldId *int `json:"originShieldId,omitempty" xml:"originShieldId,omitempty" require:"true"`
}

func (s DeleteOriginShieldPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteOriginShieldPaths) GoString() string {
  return s.String()
}

func (s *DeleteOriginShieldPaths) SetOriginShieldId(v int) *DeleteOriginShieldPaths {
  s.OriginShieldId = &v
  return s
}

type DeleteOriginShieldParameters struct {
}

func (s DeleteOriginShieldParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteOriginShieldParameters) GoString() string {
  return s.String()
}

type DeleteOriginShieldResponse struct {
  // {"en":"Response code, 0 represents success.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, 'success' represents success, otherwise provides failure information.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteOriginShieldResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteOriginShieldResponse) GoString() string {
  return s.String()
}

func (s *DeleteOriginShieldResponse) SetCode(v string) *DeleteOriginShieldResponse {
  s.Code = &v
  return s
}

func (s *DeleteOriginShieldResponse) SetMessage(v string) *DeleteOriginShieldResponse {
  s.Message = &v
  return s
}

type DeleteOriginShieldResponseHeader struct {
}

func (s DeleteOriginShieldResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteOriginShieldResponseHeader) GoString() string {
  return s.String()
}




