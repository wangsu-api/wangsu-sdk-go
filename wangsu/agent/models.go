package agent

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type AgencyQuerySaleOrderDetailRequest struct {
  // {"en":"account", "zh_CN":"账号"}
  Account *string `json:"account,omitempty" xml:"account,omitempty" require:"true"`
  // {"en":"Item code list. If it is blank, it means all, and multiple means that any one condition is satisfied.", "zh_CN":"商品编码列表。为空表示全部，多个表示满足任意一个条件即可。"}
  OfferCodes []*string `json:"offerCodes,omitempty" xml:"offerCodes,omitempty" type:"Repeated"`
  // {"en":"List of resource IDs. If it is blank, it means all, and multiple means that any one condition is satisfied.", "zh_CN":"资源ID列表。为空表示全部，多个表示满足任意一个条件即可。"}
  ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
}

func (s AgencyQuerySaleOrderDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailRequest) GoString() string {
  return s.String()
}

func (s *AgencyQuerySaleOrderDetailRequest) SetAccount(v string) *AgencyQuerySaleOrderDetailRequest {
  s.Account = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailRequest) SetOfferCodes(v []*string) *AgencyQuerySaleOrderDetailRequest {
  s.OfferCodes = v
  return s
}

func (s *AgencyQuerySaleOrderDetailRequest) SetResourceIds(v []*string) *AgencyQuerySaleOrderDetailRequest {
  s.ResourceIds = v
  return s
}

type AgencyQuerySaleOrderDetailResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data []*AgencyQuerySaleOrderDetailOrderDetailInfo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyQuerySaleOrderDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailResponse) GoString() string {
  return s.String()
}

func (s *AgencyQuerySaleOrderDetailResponse) SetCode(v int) *AgencyQuerySaleOrderDetailResponse {
  s.Code = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailResponse) SetMessage(v string) *AgencyQuerySaleOrderDetailResponse {
  s.Message = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailResponse) SetData(v []*AgencyQuerySaleOrderDetailOrderDetailInfo) *AgencyQuerySaleOrderDetailResponse {
  s.Data = v
  return s
}

type AgencyQuerySaleOrderDetailOrderDetailInfo struct {
  // {"en":"order ID", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
  // {"en":"resource ID", "zh_CN":"资源ID"}
  ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty" require:"true"`
  // {"en":"service order list", "zh_CN":"服务订单列表"}
  ServiceOrders []*AgencyQuerySaleOrderDetailServiceOrderDto `json:"serviceOrders,omitempty" xml:"serviceOrders,omitempty" require:"true" type:"Repeated"`
  // {"en":"Effective time (timestamp, millisecond)", "zh_CN":"生效时间（时间戳，毫秒）"}
  EffDate *string `json:"effDate,omitempty" xml:"effDate,omitempty" require:"true"`
  // {"en":"Expiration time: (timestamp, millisecond) time of expiration or destruction. May be empty. After the resource configuration is changed, the last order will be set with an expiration time equal to the effective time of the new order minus one second", "zh_CN":"失效时间：（时间戳，毫秒）失效或销毁时间。可能为空。资源变更配置后，上一个订单会被设置失效时间，并等于新订单的生效时间减一秒"}
  ExpDate *string `json:"expDate,omitempty" xml:"expDate,omitempty" require:"true"`
  // {"en":"Order status. (after changing the configuration, the status of the previous order will automatically change to destroyed.)
  //     OSA pretreatment
  //     OSC pending
  //     OSD destroy
  //     OSH in process
  //     OSN OK", 
  //     "zh_CN":"订单状态。（变更配置后，前一个订单状态也会自动变成销毁）
  //     OSA-预处理
  //     OSC-挂起
  //     OSD-销毁 
  //     OSH-处理中
  //     OSN正常"}
  OrderState *string `json:"orderState,omitempty" xml:"orderState,omitempty" require:"true"`
  // {"en":"Offer code, corresponding to the offers->name transmitted when generating resources", "zh_CN":"商品编码，对应生成资源时传过来的offers->name"}
  OfferCode *string `json:"offerCode,omitempty" xml:"offerCode,omitempty" require:"true"`
  // {"en":"List of product codes. Solution offers will one offer corresponds to multiple products", "zh_CN":"产品编码列表。解决方案一个商品会对应多个产品"}
  ProductCodeList []*string `json:"productCodeList,omitempty" xml:"productCodeList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data center code", "zh_CN":"数据中心编码"}
  DatacenterCode *string `json:"datacenterCode,omitempty" xml:"datacenterCode,omitempty" require:"true"`
  // {"en":"Commodity specifications Show only CBSs internal interface_ Specification attribute of display='T' ",   "zh_CN":"商品规格.只展示CBSS内部interface_display= 'T' 的规格属性"}
  Specifications []*AgencyQuerySaleOrderDetailQueryOrderSpecDto `json:"specifications,omitempty" xml:"specifications,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyQuerySaleOrderDetailOrderDetailInfo) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailOrderDetailInfo) GoString() string {
  return s.String()
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetOrderId(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.OrderId = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetResourceId(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.ResourceId = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetServiceOrders(v []*AgencyQuerySaleOrderDetailServiceOrderDto) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.ServiceOrders = v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetEffDate(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.EffDate = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetExpDate(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.ExpDate = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetOrderState(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.OrderState = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetOfferCode(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.OfferCode = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetProductCodeList(v []*string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.ProductCodeList = v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetDatacenterCode(v string) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.DatacenterCode = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailOrderDetailInfo) SetSpecifications(v []*AgencyQuerySaleOrderDetailQueryOrderSpecDto) *AgencyQuerySaleOrderDetailOrderDetailInfo {
  s.Specifications = v
  return s
}

type AgencyQuerySaleOrderDetailServiceOrderDto struct {
  // {"en":"Service order ID", "zh_CN":"服务订单ID"}
  ServiceOrderId *string `json:"serviceOrderId,omitempty" xml:"serviceOrderId,omitempty" require:"true"`
  // {"en":"Product code", "zh_CN":"产品编码"}
  ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty" require:"true"`
  // {"en":"Billing area", "zh_CN":"计费区域"}
  PayArea *string `json:"payArea,omitempty" xml:"payArea,omitempty" require:"true"`
}

func (s AgencyQuerySaleOrderDetailServiceOrderDto) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailServiceOrderDto) GoString() string {
  return s.String()
}

func (s *AgencyQuerySaleOrderDetailServiceOrderDto) SetServiceOrderId(v string) *AgencyQuerySaleOrderDetailServiceOrderDto {
  s.ServiceOrderId = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailServiceOrderDto) SetProductCode(v string) *AgencyQuerySaleOrderDetailServiceOrderDto {
  s.ProductCode = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailServiceOrderDto) SetPayArea(v string) *AgencyQuerySaleOrderDetailServiceOrderDto {
  s.PayArea = &v
  return s
}

type AgencyQuerySaleOrderDetailQueryOrderSpecDto struct {
  // {"en":"Attribute code", "zh_CN":"属性编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Attribute value", "zh_CN":"属性值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Attribute unit code", "zh_CN":"属性单位编码"}
  UnitCode *string `json:"unitCode,omitempty" xml:"unitCode,omitempty" require:"true"`
}

func (s AgencyQuerySaleOrderDetailQueryOrderSpecDto) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailQueryOrderSpecDto) GoString() string {
  return s.String()
}

func (s *AgencyQuerySaleOrderDetailQueryOrderSpecDto) SetCode(v string) *AgencyQuerySaleOrderDetailQueryOrderSpecDto {
  s.Code = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailQueryOrderSpecDto) SetValue(v string) *AgencyQuerySaleOrderDetailQueryOrderSpecDto {
  s.Value = &v
  return s
}

func (s *AgencyQuerySaleOrderDetailQueryOrderSpecDto) SetUnitCode(v string) *AgencyQuerySaleOrderDetailQueryOrderSpecDto {
  s.UnitCode = &v
  return s
}

type AgencyQuerySaleOrderDetailPaths struct {
}

func (s AgencyQuerySaleOrderDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailPaths) GoString() string {
  return s.String()
}

type AgencyQuerySaleOrderDetailParameters struct {
}

func (s AgencyQuerySaleOrderDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailParameters) GoString() string {
  return s.String()
}

type AgencyQuerySaleOrderDetailRequestHeader struct {
}

func (s AgencyQuerySaleOrderDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailRequestHeader) GoString() string {
  return s.String()
}

type AgencyQuerySaleOrderDetailResponseHeader struct {
}

func (s AgencyQuerySaleOrderDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AgencyQuerySaleOrderDetailResponseHeader) GoString() string {
  return s.String()
}




type AgencyCreateResourceRequest struct {
  // {"en":"requestId", "zh_CN":"请求Id，同一个请求的唯一标志，可用于查询该请求"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"account", "zh_CN":"用户主账号"}
  Account *string `json:"account,omitempty" xml:"account,omitempty" require:"true"`
  // {"en":"operator", "zh_CN":"操作员"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
  // {"en":"Operation timestamp (order opening time), accurate to milliseconds,Suggested vacancies, cbss will fill in the system time", 
  // 	"zh_CN":"操作时间戳（订单开通时间），精确到毫秒，建议空缺，cbss会以系统时间填入"}
  OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
  // {"en":"offer", "zh_CN":"商品"}
  AgencyCreateResourceOffer []*AgencyCreateResourceOffer `json:"offer,omitempty" xml:"offer,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyCreateResourceRequest) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceRequest) GoString() string {
  return s.String()
}

func (s *AgencyCreateResourceRequest) SetRequestId(v string) *AgencyCreateResourceRequest {
  s.RequestId = &v
  return s
}

func (s *AgencyCreateResourceRequest) SetAccount(v string) *AgencyCreateResourceRequest {
  s.Account = &v
  return s
}

func (s *AgencyCreateResourceRequest) SetOperator(v string) *AgencyCreateResourceRequest {
  s.Operator = &v
  return s
}

func (s *AgencyCreateResourceRequest) SetOperateTime(v string) *AgencyCreateResourceRequest {
  s.OperateTime = &v
  return s
}

func (s *AgencyCreateResourceRequest) SetOffer(v []*AgencyCreateResourceOffer) *AgencyCreateResourceRequest {
  s.AgencyCreateResourceOffer = v
  return s
}

type AgencyCreateResourceSpecification struct {
  // {"en":"product specification code", "zh_CN":"产品规格编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"product specification value", "zh_CN":"产品规格值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s AgencyCreateResourceSpecification) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceSpecification) GoString() string {
  return s.String()
}

func (s *AgencyCreateResourceSpecification) SetCode(v string) *AgencyCreateResourceSpecification {
  s.Code = &v
  return s
}

func (s *AgencyCreateResourceSpecification) SetValue(v string) *AgencyCreateResourceSpecification {
  s.Value = &v
  return s
}

type AgencyCreateResourceOffer struct {
  // {"en":"offer code", "zh_CN":"商品编码"}
  OfferCode *string `json:"offerCode,omitempty" xml:"offerCode,omitempty" require:"true"`
  // {"en":"The number of resources to be activated in batches (1 by default). Each resource will return an orderId",
  //    "zh_CN":"批量开通资源数量（默认为1）。每个资源会返回一个orderId"}
  Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
  // {"en":"List of resource IDs. Each quantity needs to be assigned a resource ID. If empty, CBSS generates UUID by itself. If non-null, the number of resource IDs must be the same as quantity", 
  //    "zh_CN":"资源ID列表。每个quantity需要配一个资源ID。如果为空，CBSS自行生成UUID。如果非空，资源ID个数必须和quantity一样"}
  ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
  // {"en":"Expiration timestamp. The expiration time of this order. If it is empty, cbss interprets it as 2099",
  //    "zh_CN":"到期时间戳。此次订单的到期时间。为空的话，cbss解释为2099年"}
  DueTime *string `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
  // {"en":"list of configuration items", "zh_CN":"配置项列表"}
  Specifications []*AgencyCreateResourceSpecification `json:"specifications,omitempty" xml:"specifications,omitempty" type:"Repeated"`
}

func (s AgencyCreateResourceOffer) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceOffer) GoString() string {
  return s.String()
}

func (s *AgencyCreateResourceOffer) SetOfferCode(v string) *AgencyCreateResourceOffer {
  s.OfferCode = &v
  return s
}

func (s *AgencyCreateResourceOffer) SetQuantity(v string) *AgencyCreateResourceOffer {
  s.Quantity = &v
  return s
}

func (s *AgencyCreateResourceOffer) SetResourceIds(v []*string) *AgencyCreateResourceOffer {
  s.ResourceIds = v
  return s
}

func (s *AgencyCreateResourceOffer) SetDueTime(v string) *AgencyCreateResourceOffer {
  s.DueTime = &v
  return s
}

func (s *AgencyCreateResourceOffer) SetSpecifications(v []*AgencyCreateResourceSpecification) *AgencyCreateResourceOffer {
  s.Specifications = v
  return s
}

type AgencyCreateResourceOrder struct {
  // {"en":"AgencyCreateResourceOrder ID. Each business operation generates an orderId", "zh_CN":"订单ID。每次业务操作产生一个orderId"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
  // {"en":"Resource ID. The same resource, the resourceId is always the same", "zh_CN":"资源ID。同一个资源，resourceId总是不变"}
  ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty" require:"true"`
  // {"en":"List of service order IDs", "zh_CN":"服务订单ID列表"}
  ServiceOrderIds []*string `json:"serviceOrderIds,omitempty" xml:"serviceOrderIds,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyCreateResourceOrder) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceOrder) GoString() string {
  return s.String()
}

func (s *AgencyCreateResourceOrder) SetOrderId(v string) *AgencyCreateResourceOrder {
  s.OrderId = &v
  return s
}

func (s *AgencyCreateResourceOrder) SetResourceId(v string) *AgencyCreateResourceOrder {
  s.ResourceId = &v
  return s
}

func (s *AgencyCreateResourceOrder) SetServiceOrderIds(v []*string) *AgencyCreateResourceOrder {
  s.ServiceOrderIds = v
  return s
}

type AgencyCreateResourceResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"数据"}
  Data *AgencyCreateResourceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AgencyCreateResourceResponse) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceResponse) GoString() string {
  return s.String()
}

func (s *AgencyCreateResourceResponse) SetCode(v int) *AgencyCreateResourceResponse {
  s.Code = &v
  return s
}

func (s *AgencyCreateResourceResponse) SetMessage(v string) *AgencyCreateResourceResponse {
  s.Message = &v
  return s
}

func (s *AgencyCreateResourceResponse) SetData(v *AgencyCreateResourceResponseData) *AgencyCreateResourceResponse {
  s.Data = v
  return s
}

type AgencyCreateResourceResponseData struct {
  // {"en":"Effective time (timestamp, milliseconds)", "zh_CN":"生效时间（时间戳，毫秒）"}
  EffDate *string `json:"effDate,omitempty" xml:"effDate,omitempty" require:"true"`
  // {"en":"AgencyCreateResourceOrder List. Usually 1, when the input parameter qunatity>1, each resource returns 1 order", 
  // 	  "zh_CN":"订单列表。一般是1个，当入参qunatity>1时，每个资源返回1个订单"}
  Orders []*AgencyCreateResourceOrder `json:"orders,omitempty" xml:"orders,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyCreateResourceResponseData) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceResponseData) GoString() string {
  return s.String()
}

func (s *AgencyCreateResourceResponseData) SetEffDate(v string) *AgencyCreateResourceResponseData {
  s.EffDate = &v
  return s
}

func (s *AgencyCreateResourceResponseData) SetOrders(v []*AgencyCreateResourceOrder) *AgencyCreateResourceResponseData {
  s.Orders = v
  return s
}

type AgencyCreateResourcePaths struct {
}

func (s AgencyCreateResourcePaths) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourcePaths) GoString() string {
  return s.String()
}

type AgencyCreateResourceParameters struct {
}

func (s AgencyCreateResourceParameters) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceParameters) GoString() string {
  return s.String()
}

type AgencyCreateResourceRequestHeader struct {
}

func (s AgencyCreateResourceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceRequestHeader) GoString() string {
  return s.String()
}

type AgencyCreateResourceResponseHeader struct {
}

func (s AgencyCreateResourceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AgencyCreateResourceResponseHeader) GoString() string {
  return s.String()
}




type AgencyDestroyResourceRequest struct {
  // {"en":"account", "zh_CN":"账号"}
  Account *string `json:"account,omitempty" xml:"account,omitempty" require:"true"`
  // {"en":"operator", "zh_CN":"操作员"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
  // {"en":"Destruction time (timestamp, accurate to milliseconds)Suggested vacancies, cbss will fill in the system time", "zh_CN":"销毁时间(时间戳，精确到毫秒)，建议空缺，cbss会以系统时间填入"}
  OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
  // {"en":"resource id list", "zh_CN":"资源id列表"}
  ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyDestroyResourceRequest) String() string {
  return tea.Prettify(s)
}

func (s AgencyDestroyResourceRequest) GoString() string {
  return s.String()
}

func (s *AgencyDestroyResourceRequest) SetAccount(v string) *AgencyDestroyResourceRequest {
  s.Account = &v
  return s
}

func (s *AgencyDestroyResourceRequest) SetOperator(v string) *AgencyDestroyResourceRequest {
  s.Operator = &v
  return s
}

func (s *AgencyDestroyResourceRequest) SetOperateTime(v string) *AgencyDestroyResourceRequest {
  s.OperateTime = &v
  return s
}

func (s *AgencyDestroyResourceRequest) SetResourceIds(v []*string) *AgencyDestroyResourceRequest {
  s.ResourceIds = v
  return s
}

type AgencyDestroyResourceResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AgencyDestroyResourceResponse) String() string {
  return tea.Prettify(s)
}

func (s AgencyDestroyResourceResponse) GoString() string {
  return s.String()
}

func (s *AgencyDestroyResourceResponse) SetCode(v int) *AgencyDestroyResourceResponse {
  s.Code = &v
  return s
}

func (s *AgencyDestroyResourceResponse) SetMessage(v string) *AgencyDestroyResourceResponse {
  s.Message = &v
  return s
}

type AgencyDestroyResourcePaths struct {
}

func (s AgencyDestroyResourcePaths) String() string {
  return tea.Prettify(s)
}

func (s AgencyDestroyResourcePaths) GoString() string {
  return s.String()
}

type AgencyDestroyResourceParameters struct {
}

func (s AgencyDestroyResourceParameters) String() string {
  return tea.Prettify(s)
}

func (s AgencyDestroyResourceParameters) GoString() string {
  return s.String()
}

type AgencyDestroyResourceRequestHeader struct {
}

func (s AgencyDestroyResourceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AgencyDestroyResourceRequestHeader) GoString() string {
  return s.String()
}

type AgencyDestroyResourceResponseHeader struct {
}

func (s AgencyDestroyResourceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AgencyDestroyResourceResponseHeader) GoString() string {
  return s.String()
}




