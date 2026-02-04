package loadbalancer

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteLoadBalancerRequest struct {
}

func (s DeleteLoadBalancerRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
  return s.String()
}

type DeleteLoadBalancerRequestHeader struct {
}

func (s DeleteLoadBalancerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequestHeader) GoString() string {
  return s.String()
}

type DeleteLoadBalancerPaths struct {
  // {"en":"The ID of the load balancer to be deleted.","zh_CN":"待删除的负载均衡实例的唯一标识ID"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty" require:"true"`
}

func (s DeleteLoadBalancerPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerPaths) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerPaths) SetLbId(v string) *DeleteLoadBalancerPaths {
  s.LbId = &v
  return s
}

type DeleteLoadBalancerParameters struct {
}

func (s DeleteLoadBalancerParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerParameters) GoString() string {
  return s.String()
}

type DeleteLoadBalancerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"No specific business data returned for this API","zh_CN":"该接口无特定业务数据返回"}
  Data *DeleteLoadBalancerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteLoadBalancerResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponse) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerResponse) SetCode(v string) *DeleteLoadBalancerResponse {
  s.Code = &v
  return s
}

func (s *DeleteLoadBalancerResponse) SetData(v *DeleteLoadBalancerResponseData) *DeleteLoadBalancerResponse {
  s.Data = v
  return s
}

func (s *DeleteLoadBalancerResponse) SetMessage(v string) *DeleteLoadBalancerResponse {
  s.Message = &v
  return s
}

type DeleteLoadBalancerResponseData struct {
}

func (s DeleteLoadBalancerResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseData) GoString() string {
  return s.String()
}

type DeleteLoadBalancerResponseHeader struct {
}

func (s DeleteLoadBalancerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseHeader) GoString() string {
  return s.String()
}




type UpdateLoadBalancerServerPoolRequest struct {
  // {"en":"","zh_CN":""}
  Serverpool *UpdateLoadBalancerServerPoolRequestServerpool `json:"serverpool,omitempty" xml:"serverpool,omitempty" require:"true" type:"Struct"`
}

func (s UpdateLoadBalancerServerPoolRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolRequest) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerServerPoolRequest) SetServerpool(v *UpdateLoadBalancerServerPoolRequestServerpool) *UpdateLoadBalancerServerPoolRequest {
  s.Serverpool = v
  return s
}

type UpdateLoadBalancerServerPoolRequestServerpool struct {
  // {"en":"Server pool name","zh_CN":"服务器池名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark information","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // {"en":"RS information","zh_CN":"RS信息"}
  Members []*UpdateLoadBalancerServerPoolRequestServerpoolMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerServerPoolRequestServerpool) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolRequestServerpool) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerServerPoolRequestServerpool) SetName(v string) *UpdateLoadBalancerServerPoolRequestServerpool {
  s.Name = &v
  return s
}

func (s *UpdateLoadBalancerServerPoolRequestServerpool) SetRemark(v string) *UpdateLoadBalancerServerPoolRequestServerpool {
  s.Remark = &v
  return s
}

func (s *UpdateLoadBalancerServerPoolRequestServerpool) SetMembers(v []*UpdateLoadBalancerServerPoolRequestServerpoolMembers) *UpdateLoadBalancerServerPoolRequestServerpool {
  s.Members = v
  return s
}

type UpdateLoadBalancerServerPoolRequestServerpoolMembers struct     {
  // {"en":"RS instance ID","zh_CN":"RS实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"RS service port, port range: 1-65535","zh_CN":"RS服务端口，端口范围：1-65535"}
  ServicePort *string `json:"servicePort,omitempty" xml:"servicePort,omitempty" require:"true"`
  // {"en":"Weight, range: 0-65535","zh_CN":"权重，范围：0-65535"}
  Weight *string `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s UpdateLoadBalancerServerPoolRequestServerpoolMembers) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolRequestServerpoolMembers) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerServerPoolRequestServerpoolMembers) SetInstanceId(v string) *UpdateLoadBalancerServerPoolRequestServerpoolMembers {
  s.InstanceId = &v
  return s
}

func (s *UpdateLoadBalancerServerPoolRequestServerpoolMembers) SetServicePort(v string) *UpdateLoadBalancerServerPoolRequestServerpoolMembers {
  s.ServicePort = &v
  return s
}

func (s *UpdateLoadBalancerServerPoolRequestServerpoolMembers) SetWeight(v string) *UpdateLoadBalancerServerPoolRequestServerpoolMembers {
  s.Weight = &v
  return s
}

type UpdateLoadBalancerServerPoolRequestHeader struct {
}

func (s UpdateLoadBalancerServerPoolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolRequestHeader) GoString() string {
  return s.String()
}

type UpdateLoadBalancerServerPoolPaths struct {
  // {"en":"The ID of the server pool to be updated.","zh_CN":"需要更新的服务器池ID"}
  PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty" require:"true"`
}

func (s UpdateLoadBalancerServerPoolPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolPaths) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerServerPoolPaths) SetPoolId(v string) *UpdateLoadBalancerServerPoolPaths {
  s.PoolId = &v
  return s
}

type UpdateLoadBalancerServerPoolParameters struct {
}

func (s UpdateLoadBalancerServerPoolParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolParameters) GoString() string {
  return s.String()
}

type UpdateLoadBalancerServerPoolResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *UpdateLoadBalancerServerPoolResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateLoadBalancerServerPoolResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolResponse) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerServerPoolResponse) SetCode(v string) *UpdateLoadBalancerServerPoolResponse {
  s.Code = &v
  return s
}

func (s *UpdateLoadBalancerServerPoolResponse) SetData(v *UpdateLoadBalancerServerPoolResponseData) *UpdateLoadBalancerServerPoolResponse {
  s.Data = v
  return s
}

func (s *UpdateLoadBalancerServerPoolResponse) SetMessage(v string) *UpdateLoadBalancerServerPoolResponse {
  s.Message = &v
  return s
}

type UpdateLoadBalancerServerPoolResponseData struct {
}

func (s UpdateLoadBalancerServerPoolResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolResponseData) GoString() string {
  return s.String()
}

type UpdateLoadBalancerServerPoolResponseHeader struct {
}

func (s UpdateLoadBalancerServerPoolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerServerPoolResponseHeader) GoString() string {
  return s.String()
}




type UpdateLoadBalancerListenerRequest struct {
  // {"en":"","zh_CN":""}
  Listener *UpdateLoadBalancerListenerRequestListener `json:"listener,omitempty" xml:"listener,omitempty" require:"true" type:"Struct"`
}

func (s UpdateLoadBalancerListenerRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerRequest) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerListenerRequest) SetListener(v *UpdateLoadBalancerListenerRequestListener) *UpdateLoadBalancerListenerRequest {
  s.Listener = v
  return s
}

type UpdateLoadBalancerListenerRequestListener struct {
  // {"en":"Listener Name","zh_CN":"监听名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Listener port, port range: 1-65535, the same protocol listener port cannot be duplicated under the same load balancer","zh_CN":"监听端口，端口范围：1-65535，同一负载均衡下，同样协议监听端口不能重复"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Protocol, optional values: TCP/UDP","zh_CN":"协议，可选值：TCP/UDP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"Load balancing algorithm\n- rr: Round Robin\n- wrr: Weighted Round Robin\n- dh: Destination Hash\n- sh: Source Hash\n- lc: Least Connections\n- wlc: Weighted Least Connections","zh_CN":"负载均衡算法\n- rr：轮询\n- wrr：加权轮询\n- dh：目标地址散列\n- sh：源地址散列\n- lc：最少连接\n- wlc：加权最少连接"}
  Algo *string `json:"algo,omitempty" xml:"algo,omitempty" require:"true"`
  // {"en":"Whether to enable session persistence","zh_CN":"是否开启会话保持"}
  PersistenceSession *int `json:"persistenceSession,omitempty" xml:"persistenceSession,omitempty" require:"true"`
  // {"en":"Session persistence timeout\n- Required when persistenceSession=1,\n- Cannot be specified when persistenceSession=-1\n- Value range: [1, 2678400]\n- Unit: seconds","zh_CN":"会话保持超时时间\n- persistenceSession=1时，必填，\n- persistenceSession=-1时，不能指定\n- 取值： [1, 2678400] \n- 单位：秒"}
  PersistenceTimeout *int `json:"persistenceTimeout,omitempty" xml:"persistenceTimeout,omitempty"`
  // {"en":"Server Pool ID","zh_CN":"服务器池ID"}
  PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty" require:"true"`
  // {"en":"Whether to enable health check. Values: 1 (Yes), -1 (No)","zh_CN":"是否开启健康检查。取值：1（是），-1（否）"}
  HealthCheck *int `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" require:"true"`
  // {"en":"Health check configuration, must be configured when healthCheck=1, cannot be configured when healthCheck=-1","zh_CN":"健康检查配置，healthCheck=1时必须配置，healthCheck=-1时不能配置"}
  HealthMonitor *UpdateLoadBalancerListenerRequestListenerHealthMonitor `json:"healthMonitor,omitempty" xml:"healthMonitor,omitempty" type:"Struct"`
}

func (s UpdateLoadBalancerListenerRequestListener) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerRequestListener) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerListenerRequestListener) SetName(v string) *UpdateLoadBalancerListenerRequestListener {
  s.Name = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetPort(v int) *UpdateLoadBalancerListenerRequestListener {
  s.Port = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetProtocol(v string) *UpdateLoadBalancerListenerRequestListener {
  s.Protocol = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetAlgo(v string) *UpdateLoadBalancerListenerRequestListener {
  s.Algo = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetPersistenceSession(v int) *UpdateLoadBalancerListenerRequestListener {
  s.PersistenceSession = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetPersistenceTimeout(v int) *UpdateLoadBalancerListenerRequestListener {
  s.PersistenceTimeout = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetPoolId(v string) *UpdateLoadBalancerListenerRequestListener {
  s.PoolId = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetHealthCheck(v int) *UpdateLoadBalancerListenerRequestListener {
  s.HealthCheck = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListener) SetHealthMonitor(v *UpdateLoadBalancerListenerRequestListenerHealthMonitor) *UpdateLoadBalancerListenerRequestListener {
  s.HealthMonitor = v
  return s
}

type UpdateLoadBalancerListenerRequestListenerHealthMonitor struct {
  // {"en":"Check protocol, optional values: TCP/UDP/HTTP","zh_CN":"检查协议，可选值：TCP/UDP/HTTP"}
  HealthProtocol *string `json:"healthProtocol,omitempty" xml:"healthProtocol,omitempty" require:"true"`
  // {"en":"Check port\n- Will use the backend RS instance port for health check\n- Range: 1-65535","zh_CN":"检查端口\n- 将使用后端RS实例端口进行健康检查\n- 范围：1-65535"}
  HealthPort *int `json:"healthPort,omitempty" xml:"healthPort,omitempty"`
  // {"en":"Check timeout\n- Default is 5 seconds if not specified\n- Unit: seconds","zh_CN":"检查超时时间\n- 未指定时默认5秒\n- 单位：秒"}
  HealthTimeout *int `json:"healthTimeout,omitempty" xml:"healthTimeout,omitempty" require:"true"`
  // {"en":"Check interval\n- Default is 2 seconds if not specified\n- Unit: seconds","zh_CN":"检查间隔时间\n- 未指定时默认2秒\n- 单位：秒"}
  HealthInterval *int `json:"healthInterval,omitempty" xml:"healthInterval,omitempty" require:"true"`
  // {"en":"Check failure retry count\n- Default is 3 times if not specified","zh_CN":"检查失败重试次数\n- 未指定时默认3次"}
  HealthRetry *int `json:"healthRetry,omitempty" xml:"healthRetry,omitempty" require:"true"`
  // {"en":"Check path. Required when healthProtocol=HTTP, cannot be specified for other protocols","zh_CN":"检查路径。healthProtocol=HTTP时，必填，其他协议时不能指定"}
  HealthPath *string `json:"healthPath,omitempty" xml:"healthPath,omitempty"`
  // {"en":"Check normal response\n- Default configuration is '200-299' when healthProtocol=HTTP and not specified, cannot be specified for other protocols\n\nParameter value is an array, example: \"healthStatusCodes\":[\"400\",\"200-299\"]","zh_CN":"检查正常响应\n- healthProtocol=HTTP时，未指定时，默认配置“200-299”，其他协议时不能指定\n\n参数值为数组，示例：\"healthStatusCodes\":[\"400\",\"200-299\"]"}
  HealthStatusCodes []*string `json:"healthStatusCodes,omitempty" xml:"healthStatusCodes,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerListenerRequestListenerHealthMonitor) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerRequestListenerHealthMonitor) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthProtocol(v string) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthProtocol = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthPort(v int) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthPort = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthTimeout(v int) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthTimeout = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthInterval(v int) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthInterval = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthRetry(v int) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthRetry = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthPath(v string) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthPath = &v
  return s
}

func (s *UpdateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthStatusCodes(v []*string) *UpdateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthStatusCodes = v
  return s
}

type UpdateLoadBalancerListenerRequestHeader struct {
}

func (s UpdateLoadBalancerListenerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerRequestHeader) GoString() string {
  return s.String()
}

type UpdateLoadBalancerListenerPaths struct {
  // {"en":"The ID of the load balancer listener to be updated.","zh_CN":"需要更新的负载均衡监听ID"}
  ListenerId *string `json:"listenerId,omitempty" xml:"listenerId,omitempty" require:"true"`
}

func (s UpdateLoadBalancerListenerPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerPaths) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerListenerPaths) SetListenerId(v string) *UpdateLoadBalancerListenerPaths {
  s.ListenerId = &v
  return s
}

type UpdateLoadBalancerListenerParameters struct {
}

func (s UpdateLoadBalancerListenerParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerParameters) GoString() string {
  return s.String()
}

type UpdateLoadBalancerListenerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *UpdateLoadBalancerListenerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateLoadBalancerListenerResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerResponse) GoString() string {
  return s.String()
}

func (s *UpdateLoadBalancerListenerResponse) SetCode(v string) *UpdateLoadBalancerListenerResponse {
  s.Code = &v
  return s
}

func (s *UpdateLoadBalancerListenerResponse) SetData(v *UpdateLoadBalancerListenerResponseData) *UpdateLoadBalancerListenerResponse {
  s.Data = v
  return s
}

func (s *UpdateLoadBalancerListenerResponse) SetMessage(v string) *UpdateLoadBalancerListenerResponse {
  s.Message = &v
  return s
}

type UpdateLoadBalancerListenerResponseData struct {
}

func (s UpdateLoadBalancerListenerResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerResponseData) GoString() string {
  return s.String()
}

type UpdateLoadBalancerListenerResponseHeader struct {
}

func (s UpdateLoadBalancerListenerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateLoadBalancerListenerResponseHeader) GoString() string {
  return s.String()
}




type CreateLoadBalancerRequest struct {
  // {"en":"","zh_CN":""}
  LoadBalancer *CreateLoadBalancerRequestLoadBalancer `json:"loadBalancer,omitempty" xml:"loadBalancer,omitempty" require:"true" type:"Struct"`
}

func (s CreateLoadBalancerRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerRequest) SetLoadBalancer(v *CreateLoadBalancerRequestLoadBalancer) *CreateLoadBalancerRequest {
  s.LoadBalancer = v
  return s
}

type CreateLoadBalancerRequestLoadBalancer struct {
  // {"en":"Load Balancer Name, must be unique","zh_CN":"负载均衡名称，名称不可重复"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Load Balancer Instance Specification, options:\n- lb.small\n- lb.medium\n- lb.large\n- lb.xlarge","zh_CN":"负载均衡实例规格，可选值：\n- lb.small\n- lb.medium\n- lb.large\n- lb.xlarge"}
  Spec *string `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
  // {"en":"Network Type\n- Options: PUBLIC-Public Load Balancer, PRIVATE-Private Load Balancer (not supported yet)\n- Default is Public Load Balancer","zh_CN":"网络类型\n- 取值：PUBLIC-公网负载均衡，PRIVATE-内网负载均衡（暂不支持）\n- 默认为公网负载均衡"}
  NetType *string `json:"netType,omitempty" xml:"netType,omitempty"`
  // {"en":"Network Protocol\n- Options: 4-IPv4 Load Balancer, 6-IPv6 Load Balancer (not supported yet)\n- Default is 4 (assign IPv4 VIP)","zh_CN":"网络协议\n- 取值：4-IPv4负载均衡，6-IPv6负载均衡（暂不支持）\n- 默认 4（分配IPv4的VIP）"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
  // {"en":"Specifies the carrier line for the load balancer. This parameter is only applicable to multi-line nodes. If not specified, a line will be randomly selected.","zh_CN":"指定负载均衡线路运营商\n- 仅多线节点可指定线路\n- 若未指定则随机选择一个线路"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node Name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
}

func (s CreateLoadBalancerRequestLoadBalancer) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestLoadBalancer) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerRequestLoadBalancer) SetName(v string) *CreateLoadBalancerRequestLoadBalancer {
  s.Name = &v
  return s
}

func (s *CreateLoadBalancerRequestLoadBalancer) SetSpec(v string) *CreateLoadBalancerRequestLoadBalancer {
  s.Spec = &v
  return s
}

func (s *CreateLoadBalancerRequestLoadBalancer) SetNetType(v string) *CreateLoadBalancerRequestLoadBalancer {
  s.NetType = &v
  return s
}

func (s *CreateLoadBalancerRequestLoadBalancer) SetIpProtocol(v string) *CreateLoadBalancerRequestLoadBalancer {
  s.IpProtocol = &v
  return s
}

func (s *CreateLoadBalancerRequestLoadBalancer) SetCarrier(v string) *CreateLoadBalancerRequestLoadBalancer {
  s.Carrier = &v
  return s
}

func (s *CreateLoadBalancerRequestLoadBalancer) SetNodeName(v string) *CreateLoadBalancerRequestLoadBalancer {
  s.NodeName = &v
  return s
}

type CreateLoadBalancerRequestHeader struct {
}

func (s CreateLoadBalancerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestHeader) GoString() string {
  return s.String()
}

type CreateLoadBalancerPaths struct {
}

func (s CreateLoadBalancerPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerPaths) GoString() string {
  return s.String()
}

type CreateLoadBalancerParameters struct {
}

func (s CreateLoadBalancerParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerParameters) GoString() string {
  return s.String()
}

type CreateLoadBalancerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *CreateLoadBalancerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateLoadBalancerResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerResponse) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerResponse) SetCode(v string) *CreateLoadBalancerResponse {
  s.Code = &v
  return s
}

func (s *CreateLoadBalancerResponse) SetData(v *CreateLoadBalancerResponseData) *CreateLoadBalancerResponse {
  s.Data = v
  return s
}

func (s *CreateLoadBalancerResponse) SetMessage(v string) *CreateLoadBalancerResponse {
  s.Message = &v
  return s
}

type CreateLoadBalancerResponseData struct {
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateLoadBalancerResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseData) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerResponseData) SetId(v string) *CreateLoadBalancerResponseData {
  s.Id = &v
  return s
}

type CreateLoadBalancerResponseHeader struct {
}

func (s CreateLoadBalancerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseHeader) GoString() string {
  return s.String()
}




type DisableLoadBalancerRequest struct {
  // {"en":"The list of IDs for the load balancer instances to be disabled.","zh_CN":"需要停用的负载均衡实例的ID列表"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
}

func (s DisableLoadBalancerRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerRequest) GoString() string {
  return s.String()
}

func (s *DisableLoadBalancerRequest) SetIds(v []*string) *DisableLoadBalancerRequest {
  s.Ids = v
  return s
}

type DisableLoadBalancerRequestHeader struct {
}

func (s DisableLoadBalancerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerRequestHeader) GoString() string {
  return s.String()
}

type DisableLoadBalancerPaths struct {
}

func (s DisableLoadBalancerPaths) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerPaths) GoString() string {
  return s.String()
}

type DisableLoadBalancerParameters struct {
}

func (s DisableLoadBalancerParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerParameters) GoString() string {
  return s.String()
}

type DisableLoadBalancerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *DisableLoadBalancerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DisableLoadBalancerResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerResponse) GoString() string {
  return s.String()
}

func (s *DisableLoadBalancerResponse) SetCode(v string) *DisableLoadBalancerResponse {
  s.Code = &v
  return s
}

func (s *DisableLoadBalancerResponse) SetData(v *DisableLoadBalancerResponseData) *DisableLoadBalancerResponse {
  s.Data = v
  return s
}

func (s *DisableLoadBalancerResponse) SetMessage(v string) *DisableLoadBalancerResponse {
  s.Message = &v
  return s
}

type DisableLoadBalancerResponseData struct {
  // {"en":"Batch failure details","zh_CN":"批量失败详情"}
  BatchErrorMsg []*DisableLoadBalancerResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s DisableLoadBalancerResponseData) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerResponseData) GoString() string {
  return s.String()
}

func (s *DisableLoadBalancerResponseData) SetBatchErrorMsg(v []*DisableLoadBalancerResponseDataBatchErrorMsg) *DisableLoadBalancerResponseData {
  s.BatchErrorMsg = v
  return s
}

type DisableLoadBalancerResponseDataBatchErrorMsg struct     {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Failure details","zh_CN":"失败详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DisableLoadBalancerResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *DisableLoadBalancerResponseDataBatchErrorMsg) SetCode(v string) *DisableLoadBalancerResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *DisableLoadBalancerResponseDataBatchErrorMsg) SetKey(v string) *DisableLoadBalancerResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *DisableLoadBalancerResponseDataBatchErrorMsg) SetMsg(v string) *DisableLoadBalancerResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type DisableLoadBalancerResponseHeader struct {
}

func (s DisableLoadBalancerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableLoadBalancerResponseHeader) GoString() string {
  return s.String()
}




type DeleteLoadBalancerServerPoolsRequest struct {
}

func (s DeleteLoadBalancerServerPoolsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsRequest) GoString() string {
  return s.String()
}

type DeleteLoadBalancerServerPoolsRequestHeader struct {
}

func (s DeleteLoadBalancerServerPoolsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsRequestHeader) GoString() string {
  return s.String()
}

type DeleteLoadBalancerServerPoolsPaths struct {
  // {"en":"The ID of the load balancer server pool to be deleted. Multiple IDs can be provided, separated by commas, to support batch deletion.","zh_CN":"负载均衡服务器池ID，支持批量删除（多个ID逗号分隔）"}
  PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty" require:"true"`
}

func (s DeleteLoadBalancerServerPoolsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsPaths) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerServerPoolsPaths) SetPoolId(v string) *DeleteLoadBalancerServerPoolsPaths {
  s.PoolId = &v
  return s
}

type DeleteLoadBalancerServerPoolsParameters struct {
}

func (s DeleteLoadBalancerServerPoolsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsParameters) GoString() string {
  return s.String()
}

type DeleteLoadBalancerServerPoolsResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *DeleteLoadBalancerServerPoolsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteLoadBalancerServerPoolsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsResponse) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerServerPoolsResponse) SetCode(v string) *DeleteLoadBalancerServerPoolsResponse {
  s.Code = &v
  return s
}

func (s *DeleteLoadBalancerServerPoolsResponse) SetData(v *DeleteLoadBalancerServerPoolsResponseData) *DeleteLoadBalancerServerPoolsResponse {
  s.Data = v
  return s
}

func (s *DeleteLoadBalancerServerPoolsResponse) SetMessage(v string) *DeleteLoadBalancerServerPoolsResponse {
  s.Message = &v
  return s
}

type DeleteLoadBalancerServerPoolsResponseData struct {
  // {"en":"Batch failure details","zh_CN":"批量失败详情"}
  BatchErrorMsg []*DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteLoadBalancerServerPoolsResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsResponseData) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerServerPoolsResponseData) SetBatchErrorMsg(v []*DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg) *DeleteLoadBalancerServerPoolsResponseData {
  s.BatchErrorMsg = v
  return s
}

type DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg struct     {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Server Pool ID","zh_CN":"服务器池ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Failure details","zh_CN":"失败详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg) SetCode(v string) *DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg) SetKey(v string) *DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg) SetMsg(v string) *DeleteLoadBalancerServerPoolsResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type DeleteLoadBalancerServerPoolsResponseHeader struct {
}

func (s DeleteLoadBalancerServerPoolsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerServerPoolsResponseHeader) GoString() string {
  return s.String()
}




type QueryLoadBalancerListenersRequest struct {
}

func (s QueryLoadBalancerListenersRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersRequest) GoString() string {
  return s.String()
}

type QueryLoadBalancerListenersRequestHeader struct {
}

func (s QueryLoadBalancerListenersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersRequestHeader) GoString() string {
  return s.String()
}

type QueryLoadBalancerListenersPaths struct {
}

func (s QueryLoadBalancerListenersPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersPaths) GoString() string {
  return s.String()
}

type QueryLoadBalancerListenersParameters struct {
  // {"en":"Listener ID, supports multiple values (comma-separated)","zh_CN":"监听ID，支持多值（逗号分隔）"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Load Balancer ID, supports multiple values (comma-separated)","zh_CN":"负载均衡ID，支持多值（逗号分隔）"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty"`
  // {"en":"Listener status, supports multiple values (comma-separated)","zh_CN":"监听状态，支持多值（逗号分隔）"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Server Pool ID, supports multiple values (comma-separated)","zh_CN":"服务器池ID，支持多值（逗号分隔）"}
  PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty"`
}

func (s QueryLoadBalancerListenersParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersParameters) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerListenersParameters) SetId(v string) *QueryLoadBalancerListenersParameters {
  s.Id = &v
  return s
}

func (s *QueryLoadBalancerListenersParameters) SetLbId(v string) *QueryLoadBalancerListenersParameters {
  s.LbId = &v
  return s
}

func (s *QueryLoadBalancerListenersParameters) SetState(v string) *QueryLoadBalancerListenersParameters {
  s.State = &v
  return s
}

func (s *QueryLoadBalancerListenersParameters) SetPoolId(v string) *QueryLoadBalancerListenersParameters {
  s.PoolId = &v
  return s
}

type QueryLoadBalancerListenersResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *QueryLoadBalancerListenersResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QueryLoadBalancerListenersResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersResponse) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerListenersResponse) SetCode(v string) *QueryLoadBalancerListenersResponse {
  s.Code = &v
  return s
}

func (s *QueryLoadBalancerListenersResponse) SetData(v *QueryLoadBalancerListenersResponseData) *QueryLoadBalancerListenersResponse {
  s.Data = v
  return s
}

func (s *QueryLoadBalancerListenersResponse) SetMessage(v string) *QueryLoadBalancerListenersResponse {
  s.Message = &v
  return s
}

type QueryLoadBalancerListenersResponseData struct {
  // {"en":"Load Balancer Listener List","zh_CN":"负载均衡监听列表"}
  Listeners []*QueryLoadBalancerListenersResponseDataListeners `json:"listeners,omitempty" xml:"listeners,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLoadBalancerListenersResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersResponseData) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerListenersResponseData) SetListeners(v []*QueryLoadBalancerListenersResponseDataListeners) *QueryLoadBalancerListenersResponseData {
  s.Listeners = v
  return s
}

type QueryLoadBalancerListenersResponseDataListeners struct     {
  // {"en":"Load Balancer Listener ID","zh_CN":"负载均衡监听ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Associated Load Balancer ID","zh_CN":"所属负载均衡ID"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty" require:"true"`
  // {"en":"Name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"State","zh_CN":"状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Port","zh_CN":"端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Protocol","zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"Load Balancing Algorithm","zh_CN":"负载均衡算法"}
  Algo *string `json:"algo,omitempty" xml:"algo,omitempty" require:"true"`
  // {"en":"Session Persistence Enabled","zh_CN":"是否开启会话保持"}
  PersistenceSession *string `json:"persistenceSession,omitempty" xml:"persistenceSession,omitempty" require:"true"`
  // {"en":"Session Persistence Timeout","zh_CN":"会话保持超时时间"}
  PersistenceTimeout *string `json:"persistenceTimeout,omitempty" xml:"persistenceTimeout,omitempty" require:"true"`
  // {"en":"Server Pool ID","zh_CN":"服务器池ID"}
  PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty" require:"true"`
  // {"en":"Health Check Enabled","zh_CN":"是否开启健康检查"}
  HealthCheck *string `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" require:"true"`
  // {"en":"Health Check Protocol","zh_CN":"健康检查协议"}
  HealthProtocol *string `json:"healthProtocol,omitempty" xml:"healthProtocol,omitempty" require:"true"`
  // {"en":"Health Check Port","zh_CN":"健康检查端口"}
  HealthPort *string `json:"healthPort,omitempty" xml:"healthPort,omitempty" require:"true"`
  // {"en":"Health Check Timeout","zh_CN":"健康检查超时时间"}
  HealthTimeout *string `json:"healthTimeout,omitempty" xml:"healthTimeout,omitempty" require:"true"`
  // {"en":"Health Check Interval","zh_CN":"健康检查间隔"}
  HealthInterval *string `json:"healthInterval,omitempty" xml:"healthInterval,omitempty" require:"true"`
  // {"en":"Health Check Retry Count","zh_CN":"健康检查失败重试次数"}
  HealthRetry *string `json:"healthRetry,omitempty" xml:"healthRetry,omitempty" require:"true"`
  // {"en":"Health Check Path","zh_CN":"健康检查路径"}
  HealthPath *string `json:"healthPath,omitempty" xml:"healthPath,omitempty" require:"true"`
  // {"en":"Normal Response Code","zh_CN":"检查正常响应码"}
  HealthStatusCodes []*string `json:"healthStatusCodes,omitempty" xml:"healthStatusCodes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Creation Time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modification Time","zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s QueryLoadBalancerListenersResponseDataListeners) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersResponseDataListeners) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetId(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.Id = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetLbId(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.LbId = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetName(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.Name = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetState(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.State = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetPort(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.Port = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetProtocol(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.Protocol = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetAlgo(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.Algo = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetPersistenceSession(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.PersistenceSession = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetPersistenceTimeout(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.PersistenceTimeout = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetPoolId(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.PoolId = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthCheck(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthCheck = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthProtocol(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthProtocol = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthPort(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthPort = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthTimeout(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthTimeout = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthInterval(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthInterval = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthRetry(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthRetry = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthPath(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthPath = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetHealthStatusCodes(v []*string) *QueryLoadBalancerListenersResponseDataListeners {
  s.HealthStatusCodes = v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetCreateTime(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.CreateTime = &v
  return s
}

func (s *QueryLoadBalancerListenersResponseDataListeners) SetModifyTime(v string) *QueryLoadBalancerListenersResponseDataListeners {
  s.ModifyTime = &v
  return s
}

type QueryLoadBalancerListenersResponseHeader struct {
}

func (s QueryLoadBalancerListenersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerListenersResponseHeader) GoString() string {
  return s.String()
}




type StartLoadBalancerListenerRequest struct {
  // {"en":"The list of IDs for the load balancer listeners to be enabled.","zh_CN":"需要启用的负载均衡监听的ID列表"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
}

func (s StartLoadBalancerListenerRequest) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerRequest) GoString() string {
  return s.String()
}

func (s *StartLoadBalancerListenerRequest) SetIds(v []*string) *StartLoadBalancerListenerRequest {
  s.Ids = v
  return s
}

type StartLoadBalancerListenerRequestHeader struct {
}

func (s StartLoadBalancerListenerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerRequestHeader) GoString() string {
  return s.String()
}

type StartLoadBalancerListenerPaths struct {
}

func (s StartLoadBalancerListenerPaths) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerPaths) GoString() string {
  return s.String()
}

type StartLoadBalancerListenerParameters struct {
}

func (s StartLoadBalancerListenerParameters) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerParameters) GoString() string {
  return s.String()
}

type StartLoadBalancerListenerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *StartLoadBalancerListenerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s StartLoadBalancerListenerResponse) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerResponse) GoString() string {
  return s.String()
}

func (s *StartLoadBalancerListenerResponse) SetCode(v string) *StartLoadBalancerListenerResponse {
  s.Code = &v
  return s
}

func (s *StartLoadBalancerListenerResponse) SetData(v *StartLoadBalancerListenerResponseData) *StartLoadBalancerListenerResponse {
  s.Data = v
  return s
}

func (s *StartLoadBalancerListenerResponse) SetMessage(v string) *StartLoadBalancerListenerResponse {
  s.Message = &v
  return s
}

type StartLoadBalancerListenerResponseData struct {
  // {"en":"Batch failure details","zh_CN":"批量失败详情"}
  BatchErrorMsg []*StartLoadBalancerListenerResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s StartLoadBalancerListenerResponseData) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerResponseData) GoString() string {
  return s.String()
}

func (s *StartLoadBalancerListenerResponseData) SetBatchErrorMsg(v []*StartLoadBalancerListenerResponseDataBatchErrorMsg) *StartLoadBalancerListenerResponseData {
  s.BatchErrorMsg = v
  return s
}

type StartLoadBalancerListenerResponseDataBatchErrorMsg struct     {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Listener ID","zh_CN":"监听ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Failure details","zh_CN":"失败详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s StartLoadBalancerListenerResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *StartLoadBalancerListenerResponseDataBatchErrorMsg) SetCode(v string) *StartLoadBalancerListenerResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *StartLoadBalancerListenerResponseDataBatchErrorMsg) SetKey(v string) *StartLoadBalancerListenerResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *StartLoadBalancerListenerResponseDataBatchErrorMsg) SetMsg(v string) *StartLoadBalancerListenerResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type StartLoadBalancerListenerResponseHeader struct {
}

func (s StartLoadBalancerListenerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StartLoadBalancerListenerResponseHeader) GoString() string {
  return s.String()
}




type StopLoadBalancerListenerRequest struct {
  // {"en":"List of load balancer listener IDs to be stopped","zh_CN":"需要停用的负载均衡监听的ID列表"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
}

func (s StopLoadBalancerListenerRequest) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerRequest) GoString() string {
  return s.String()
}

func (s *StopLoadBalancerListenerRequest) SetIds(v []*string) *StopLoadBalancerListenerRequest {
  s.Ids = v
  return s
}

type StopLoadBalancerListenerRequestHeader struct {
}

func (s StopLoadBalancerListenerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerRequestHeader) GoString() string {
  return s.String()
}

type StopLoadBalancerListenerPaths struct {
}

func (s StopLoadBalancerListenerPaths) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerPaths) GoString() string {
  return s.String()
}

type StopLoadBalancerListenerParameters struct {
}

func (s StopLoadBalancerListenerParameters) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerParameters) GoString() string {
  return s.String()
}

type StopLoadBalancerListenerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *StopLoadBalancerListenerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s StopLoadBalancerListenerResponse) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerResponse) GoString() string {
  return s.String()
}

func (s *StopLoadBalancerListenerResponse) SetCode(v string) *StopLoadBalancerListenerResponse {
  s.Code = &v
  return s
}

func (s *StopLoadBalancerListenerResponse) SetData(v *StopLoadBalancerListenerResponseData) *StopLoadBalancerListenerResponse {
  s.Data = v
  return s
}

func (s *StopLoadBalancerListenerResponse) SetMessage(v string) *StopLoadBalancerListenerResponse {
  s.Message = &v
  return s
}

type StopLoadBalancerListenerResponseData struct {
  // {"en":"Batch failure details","zh_CN":"批量失败详情"}
  BatchErrorMsg []*StopLoadBalancerListenerResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s StopLoadBalancerListenerResponseData) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerResponseData) GoString() string {
  return s.String()
}

func (s *StopLoadBalancerListenerResponseData) SetBatchErrorMsg(v []*StopLoadBalancerListenerResponseDataBatchErrorMsg) *StopLoadBalancerListenerResponseData {
  s.BatchErrorMsg = v
  return s
}

type StopLoadBalancerListenerResponseDataBatchErrorMsg struct     {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Listener ID","zh_CN":"监听ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Failure details","zh_CN":"失败详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s StopLoadBalancerListenerResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *StopLoadBalancerListenerResponseDataBatchErrorMsg) SetCode(v string) *StopLoadBalancerListenerResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *StopLoadBalancerListenerResponseDataBatchErrorMsg) SetKey(v string) *StopLoadBalancerListenerResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *StopLoadBalancerListenerResponseDataBatchErrorMsg) SetMsg(v string) *StopLoadBalancerListenerResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type StopLoadBalancerListenerResponseHeader struct {
}

func (s StopLoadBalancerListenerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StopLoadBalancerListenerResponseHeader) GoString() string {
  return s.String()
}




type QueryLoadBalancerRequest struct {
}

func (s QueryLoadBalancerRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerRequest) GoString() string {
  return s.String()
}

type QueryLoadBalancerRequestHeader struct {
}

func (s QueryLoadBalancerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerRequestHeader) GoString() string {
  return s.String()
}

type QueryLoadBalancerPaths struct {
}

func (s QueryLoadBalancerPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerPaths) GoString() string {
  return s.String()
}

type QueryLoadBalancerParameters struct {
  // {"en":"Load Balancer ID, supports multiple values (comma separated)","zh_CN":"负载均衡ID，支持多值（逗号分隔）"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Load Balancer specification, supports multiple values (comma separated)","zh_CN":"负载均衡规格，支持多值（逗号分隔）"}
  Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"Load Balancer state, supports multiple values (comma separated)","zh_CN":"负载均衡状态，支持多值（逗号分隔）"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Node name, supports multiple values (comma separated)","zh_CN":"节点名称，支持多值（逗号分隔）"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
}

func (s QueryLoadBalancerParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerParameters) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerParameters) SetId(v string) *QueryLoadBalancerParameters {
  s.Id = &v
  return s
}

func (s *QueryLoadBalancerParameters) SetSpec(v string) *QueryLoadBalancerParameters {
  s.Spec = &v
  return s
}

func (s *QueryLoadBalancerParameters) SetState(v string) *QueryLoadBalancerParameters {
  s.State = &v
  return s
}

func (s *QueryLoadBalancerParameters) SetNodeName(v string) *QueryLoadBalancerParameters {
  s.NodeName = &v
  return s
}

type QueryLoadBalancerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *QueryLoadBalancerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QueryLoadBalancerResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerResponse) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerResponse) SetCode(v string) *QueryLoadBalancerResponse {
  s.Code = &v
  return s
}

func (s *QueryLoadBalancerResponse) SetData(v *QueryLoadBalancerResponseData) *QueryLoadBalancerResponse {
  s.Data = v
  return s
}

func (s *QueryLoadBalancerResponse) SetMessage(v string) *QueryLoadBalancerResponse {
  s.Message = &v
  return s
}

type QueryLoadBalancerResponseData struct {
  // {"en":"Load Balancer List","zh_CN":"负载均衡列表"}
  LoadBalancers []*QueryLoadBalancerResponseDataLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLoadBalancerResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerResponseData) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerResponseData) SetLoadBalancers(v []*QueryLoadBalancerResponseDataLoadBalancers) *QueryLoadBalancerResponseData {
  s.LoadBalancers = v
  return s
}

type QueryLoadBalancerResponseDataLoadBalancers struct     {
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Load Balancer Name","zh_CN":"负载均衡名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"State","zh_CN":"状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"State Description","zh_CN":"状态描述"}
  StateNote *string `json:"stateNote,omitempty" xml:"stateNote,omitempty" require:"true"`
  // {"en":"Specification Type","zh_CN":"规格类型"}
  Spec *string `json:"spec,omitempty" xml:"spec,omitempty" require:"true"`
  // {"en":"Node Name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Creation Time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modification Time","zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  // {"en":"VIP Information","zh_CN":"vip信息"}
  VipInfo *QueryLoadBalancerResponseDataLoadBalancersVipInfo `json:"vipInfo,omitempty" xml:"vipInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryLoadBalancerResponseDataLoadBalancers) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerResponseDataLoadBalancers) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetId(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.Id = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetName(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.Name = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetState(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.State = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetStateNote(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.StateNote = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetSpec(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.Spec = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetNodeName(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.NodeName = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetCreateTime(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.CreateTime = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetModifyTime(v string) *QueryLoadBalancerResponseDataLoadBalancers {
  s.ModifyTime = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancers) SetVipInfo(v *QueryLoadBalancerResponseDataLoadBalancersVipInfo) *QueryLoadBalancerResponseDataLoadBalancers {
  s.VipInfo = v
  return s
}

type QueryLoadBalancerResponseDataLoadBalancersVipInfo struct {
  // {"en":"VIP Address","zh_CN":"vip地址"}
  Vip *string `json:"vip,omitempty" xml:"vip,omitempty" require:"true"`
  // {"en":"Network Type","zh_CN":"网络类型"}
  NetType *string `json:"netType,omitempty" xml:"netType,omitempty" require:"true"`
  // {"en":"IP Protocol","zh_CN":"IP协议"}
  IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty" require:"true"`
}

func (s QueryLoadBalancerResponseDataLoadBalancersVipInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerResponseDataLoadBalancersVipInfo) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerResponseDataLoadBalancersVipInfo) SetVip(v string) *QueryLoadBalancerResponseDataLoadBalancersVipInfo {
  s.Vip = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancersVipInfo) SetNetType(v string) *QueryLoadBalancerResponseDataLoadBalancersVipInfo {
  s.NetType = &v
  return s
}

func (s *QueryLoadBalancerResponseDataLoadBalancersVipInfo) SetIpProtocol(v string) *QueryLoadBalancerResponseDataLoadBalancersVipInfo {
  s.IpProtocol = &v
  return s
}

type QueryLoadBalancerResponseHeader struct {
}

func (s QueryLoadBalancerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerResponseHeader) GoString() string {
  return s.String()
}




type QueryLoadBalancerServerPoolsRequest struct {
}

func (s QueryLoadBalancerServerPoolsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsRequest) GoString() string {
  return s.String()
}

type QueryLoadBalancerServerPoolsRequestHeader struct {
}

func (s QueryLoadBalancerServerPoolsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsRequestHeader) GoString() string {
  return s.String()
}

type QueryLoadBalancerServerPoolsPaths struct {
}

func (s QueryLoadBalancerServerPoolsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsPaths) GoString() string {
  return s.String()
}

type QueryLoadBalancerServerPoolsParameters struct {
  // {"en":"Load Balancer Server Pool ID","zh_CN":"负载均衡服务器池ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty"`
}

func (s QueryLoadBalancerServerPoolsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsParameters) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerServerPoolsParameters) SetId(v string) *QueryLoadBalancerServerPoolsParameters {
  s.Id = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsParameters) SetLbId(v string) *QueryLoadBalancerServerPoolsParameters {
  s.LbId = &v
  return s
}

type QueryLoadBalancerServerPoolsResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *QueryLoadBalancerServerPoolsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QueryLoadBalancerServerPoolsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsResponse) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerServerPoolsResponse) SetCode(v string) *QueryLoadBalancerServerPoolsResponse {
  s.Code = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponse) SetData(v *QueryLoadBalancerServerPoolsResponseData) *QueryLoadBalancerServerPoolsResponse {
  s.Data = v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponse) SetMessage(v string) *QueryLoadBalancerServerPoolsResponse {
  s.Message = &v
  return s
}

type QueryLoadBalancerServerPoolsResponseData struct {
  // {"en":"Server Pool ID","zh_CN":"服务器池ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty" require:"true"`
  // {"en":"Name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark","zh_CN":"备注"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Creation Time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modification Time","zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  // {"en":"List of associated RS instances","zh_CN":"关联的RS实例列表"}
  Members []*QueryLoadBalancerServerPoolsResponseDataMembers `json:"members,omitempty" xml:"members,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLoadBalancerServerPoolsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsResponseData) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetId(v string) *QueryLoadBalancerServerPoolsResponseData {
  s.Id = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetLbId(v string) *QueryLoadBalancerServerPoolsResponseData {
  s.LbId = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetName(v string) *QueryLoadBalancerServerPoolsResponseData {
  s.Name = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetRemark(v string) *QueryLoadBalancerServerPoolsResponseData {
  s.Remark = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetCreateTime(v string) *QueryLoadBalancerServerPoolsResponseData {
  s.CreateTime = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetModifyTime(v string) *QueryLoadBalancerServerPoolsResponseData {
  s.ModifyTime = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseData) SetMembers(v []*QueryLoadBalancerServerPoolsResponseDataMembers) *QueryLoadBalancerServerPoolsResponseData {
  s.Members = v
  return s
}

type QueryLoadBalancerServerPoolsResponseDataMembers struct     {
  // {"en":"RS Instance ID","zh_CN":"RS实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"Service Port","zh_CN":"服务端口"}
  ServicePort *int `json:"servicePort,omitempty" xml:"servicePort,omitempty" require:"true"`
  // {"en":"Weight","zh_CN":"权重"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"Creation Time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modification Time","zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s QueryLoadBalancerServerPoolsResponseDataMembers) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsResponseDataMembers) GoString() string {
  return s.String()
}

func (s *QueryLoadBalancerServerPoolsResponseDataMembers) SetInstanceId(v string) *QueryLoadBalancerServerPoolsResponseDataMembers {
  s.InstanceId = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseDataMembers) SetServicePort(v int) *QueryLoadBalancerServerPoolsResponseDataMembers {
  s.ServicePort = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseDataMembers) SetWeight(v int) *QueryLoadBalancerServerPoolsResponseDataMembers {
  s.Weight = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseDataMembers) SetCreateTime(v string) *QueryLoadBalancerServerPoolsResponseDataMembers {
  s.CreateTime = &v
  return s
}

func (s *QueryLoadBalancerServerPoolsResponseDataMembers) SetModifyTime(v string) *QueryLoadBalancerServerPoolsResponseDataMembers {
  s.ModifyTime = &v
  return s
}

type QueryLoadBalancerServerPoolsResponseHeader struct {
}

func (s QueryLoadBalancerServerPoolsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLoadBalancerServerPoolsResponseHeader) GoString() string {
  return s.String()
}




type EnableLoadBalancerRequest struct {
  // {"en":"The list of IDs for the load balancer instances to be enabled.","zh_CN":"需要启用的负载均衡实例的ID列表"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
}

func (s EnableLoadBalancerRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerRequest) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerRequest) SetIds(v []*string) *EnableLoadBalancerRequest {
  s.Ids = v
  return s
}

type EnableLoadBalancerRequestHeader struct {
}

func (s EnableLoadBalancerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerRequestHeader) GoString() string {
  return s.String()
}

type EnableLoadBalancerPaths struct {
}

func (s EnableLoadBalancerPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerPaths) GoString() string {
  return s.String()
}

type EnableLoadBalancerParameters struct {
}

func (s EnableLoadBalancerParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerParameters) GoString() string {
  return s.String()
}

type EnableLoadBalancerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *EnableLoadBalancerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EnableLoadBalancerResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerResponse) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerResponse) SetCode(v string) *EnableLoadBalancerResponse {
  s.Code = &v
  return s
}

func (s *EnableLoadBalancerResponse) SetData(v *EnableLoadBalancerResponseData) *EnableLoadBalancerResponse {
  s.Data = v
  return s
}

func (s *EnableLoadBalancerResponse) SetMessage(v string) *EnableLoadBalancerResponse {
  s.Message = &v
  return s
}

type EnableLoadBalancerResponseData struct {
  // {"en":"Batch failure details","zh_CN":"批量失败详情"}
  BatchErrorMsg []*EnableLoadBalancerResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s EnableLoadBalancerResponseData) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerResponseData) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerResponseData) SetBatchErrorMsg(v []*EnableLoadBalancerResponseDataBatchErrorMsg) *EnableLoadBalancerResponseData {
  s.BatchErrorMsg = v
  return s
}

type EnableLoadBalancerResponseDataBatchErrorMsg struct     {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Failure details","zh_CN":"失败详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s EnableLoadBalancerResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerResponseDataBatchErrorMsg) SetCode(v string) *EnableLoadBalancerResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *EnableLoadBalancerResponseDataBatchErrorMsg) SetKey(v string) *EnableLoadBalancerResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *EnableLoadBalancerResponseDataBatchErrorMsg) SetMsg(v string) *EnableLoadBalancerResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type EnableLoadBalancerResponseHeader struct {
}

func (s EnableLoadBalancerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableLoadBalancerResponseHeader) GoString() string {
  return s.String()
}




type DeleteLoadBalancerListenersRequest struct {
}

func (s DeleteLoadBalancerListenersRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersRequest) GoString() string {
  return s.String()
}

type DeleteLoadBalancerListenersRequestHeader struct {
}

func (s DeleteLoadBalancerListenersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersRequestHeader) GoString() string {
  return s.String()
}

type DeleteLoadBalancerListenersPaths struct {
  // {"en":"The load balancer listener ID. Multiple IDs can be separated by commas for batch deletion.","zh_CN":"负载均衡监听ID，支持批量删除（多个ID逗号分隔）"}
  ListenerId *string `json:"listenerId,omitempty" xml:"listenerId,omitempty" require:"true"`
}

func (s DeleteLoadBalancerListenersPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersPaths) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerListenersPaths) SetListenerId(v string) *DeleteLoadBalancerListenersPaths {
  s.ListenerId = &v
  return s
}

type DeleteLoadBalancerListenersParameters struct {
}

func (s DeleteLoadBalancerListenersParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersParameters) GoString() string {
  return s.String()
}

type DeleteLoadBalancerListenersResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *DeleteLoadBalancerListenersResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteLoadBalancerListenersResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersResponse) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerListenersResponse) SetCode(v string) *DeleteLoadBalancerListenersResponse {
  s.Code = &v
  return s
}

func (s *DeleteLoadBalancerListenersResponse) SetData(v *DeleteLoadBalancerListenersResponseData) *DeleteLoadBalancerListenersResponse {
  s.Data = v
  return s
}

func (s *DeleteLoadBalancerListenersResponse) SetMessage(v string) *DeleteLoadBalancerListenersResponse {
  s.Message = &v
  return s
}

type DeleteLoadBalancerListenersResponseData struct {
  // {"en":"Batch failure details","zh_CN":"批量失败详情"}
  BatchErrorMsg []*DeleteLoadBalancerListenersResponseDataBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteLoadBalancerListenersResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersResponseData) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerListenersResponseData) SetBatchErrorMsg(v []*DeleteLoadBalancerListenersResponseDataBatchErrorMsg) *DeleteLoadBalancerListenersResponseData {
  s.BatchErrorMsg = v
  return s
}

type DeleteLoadBalancerListenersResponseDataBatchErrorMsg struct     {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Listener ID","zh_CN":"监听ID"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Failure details","zh_CN":"失败详情"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteLoadBalancerListenersResponseDataBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersResponseDataBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *DeleteLoadBalancerListenersResponseDataBatchErrorMsg) SetCode(v string) *DeleteLoadBalancerListenersResponseDataBatchErrorMsg {
  s.Code = &v
  return s
}

func (s *DeleteLoadBalancerListenersResponseDataBatchErrorMsg) SetKey(v string) *DeleteLoadBalancerListenersResponseDataBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *DeleteLoadBalancerListenersResponseDataBatchErrorMsg) SetMsg(v string) *DeleteLoadBalancerListenersResponseDataBatchErrorMsg {
  s.Msg = &v
  return s
}

type DeleteLoadBalancerListenersResponseHeader struct {
}

func (s DeleteLoadBalancerListenersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenersResponseHeader) GoString() string {
  return s.String()
}




type CreateLoadBalancerServerPoolRequest struct {
  // {"en":"","zh_CN":""}
  Serverpool *CreateLoadBalancerServerPoolRequestServerpool `json:"serverpool,omitempty" xml:"serverpool,omitempty" require:"true" type:"Struct"`
}

func (s CreateLoadBalancerServerPoolRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolRequest) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerServerPoolRequest) SetServerpool(v *CreateLoadBalancerServerPoolRequestServerpool) *CreateLoadBalancerServerPoolRequest {
  s.Serverpool = v
  return s
}

type CreateLoadBalancerServerPoolRequestServerpool struct {
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty" require:"true"`
  // {"en":"Server Pool Name","zh_CN":"服务器池名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Remark Information","zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // {"en":"Server pool member information, including instance ID, weight, and port.","zh_CN":"服务器池成员信息，包含实例ID、权重和端口等。"}
  Members []*CreateLoadBalancerServerPoolRequestServerpoolMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerServerPoolRequestServerpool) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolRequestServerpool) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerServerPoolRequestServerpool) SetLbId(v string) *CreateLoadBalancerServerPoolRequestServerpool {
  s.LbId = &v
  return s
}

func (s *CreateLoadBalancerServerPoolRequestServerpool) SetName(v string) *CreateLoadBalancerServerPoolRequestServerpool {
  s.Name = &v
  return s
}

func (s *CreateLoadBalancerServerPoolRequestServerpool) SetRemark(v string) *CreateLoadBalancerServerPoolRequestServerpool {
  s.Remark = &v
  return s
}

func (s *CreateLoadBalancerServerPoolRequestServerpool) SetMembers(v []*CreateLoadBalancerServerPoolRequestServerpoolMembers) *CreateLoadBalancerServerPoolRequestServerpool {
  s.Members = v
  return s
}

type CreateLoadBalancerServerPoolRequestServerpoolMembers struct     {
  // {"en":"RS Instance ID","zh_CN":"RS实例ID"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"RS Service Port, port range: 1-65535","zh_CN":"RS服务端口，端口范围：1-65535"}
  ServicePort *string `json:"servicePort,omitempty" xml:"servicePort,omitempty" require:"true"`
  // {"en":"Weight, range: 0-65535","zh_CN":"权重，范围：0-65535"}
  Weight *string `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s CreateLoadBalancerServerPoolRequestServerpoolMembers) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolRequestServerpoolMembers) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerServerPoolRequestServerpoolMembers) SetInstanceId(v string) *CreateLoadBalancerServerPoolRequestServerpoolMembers {
  s.InstanceId = &v
  return s
}

func (s *CreateLoadBalancerServerPoolRequestServerpoolMembers) SetServicePort(v string) *CreateLoadBalancerServerPoolRequestServerpoolMembers {
  s.ServicePort = &v
  return s
}

func (s *CreateLoadBalancerServerPoolRequestServerpoolMembers) SetWeight(v string) *CreateLoadBalancerServerPoolRequestServerpoolMembers {
  s.Weight = &v
  return s
}

type CreateLoadBalancerServerPoolRequestHeader struct {
}

func (s CreateLoadBalancerServerPoolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolRequestHeader) GoString() string {
  return s.String()
}

type CreateLoadBalancerServerPoolPaths struct {
}

func (s CreateLoadBalancerServerPoolPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolPaths) GoString() string {
  return s.String()
}

type CreateLoadBalancerServerPoolParameters struct {
}

func (s CreateLoadBalancerServerPoolParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolParameters) GoString() string {
  return s.String()
}

type CreateLoadBalancerServerPoolResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *CreateLoadBalancerServerPoolResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateLoadBalancerServerPoolResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolResponse) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerServerPoolResponse) SetCode(v string) *CreateLoadBalancerServerPoolResponse {
  s.Code = &v
  return s
}

func (s *CreateLoadBalancerServerPoolResponse) SetData(v *CreateLoadBalancerServerPoolResponseData) *CreateLoadBalancerServerPoolResponse {
  s.Data = v
  return s
}

func (s *CreateLoadBalancerServerPoolResponse) SetMessage(v string) *CreateLoadBalancerServerPoolResponse {
  s.Message = &v
  return s
}

type CreateLoadBalancerServerPoolResponseData struct {
  // {"en":"Server Pool ID","zh_CN":"服务器池ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateLoadBalancerServerPoolResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolResponseData) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerServerPoolResponseData) SetId(v string) *CreateLoadBalancerServerPoolResponseData {
  s.Id = &v
  return s
}

type CreateLoadBalancerServerPoolResponseHeader struct {
}

func (s CreateLoadBalancerServerPoolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerServerPoolResponseHeader) GoString() string {
  return s.String()
}




type CreateLoadBalancerListenerRequest struct {
  // {"en":"","zh_CN":""}
  Listener *CreateLoadBalancerListenerRequestListener `json:"listener,omitempty" xml:"listener,omitempty" type:"Struct"`
}

func (s CreateLoadBalancerListenerRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerRequest) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerListenerRequest) SetListener(v *CreateLoadBalancerListenerRequestListener) *CreateLoadBalancerListenerRequest {
  s.Listener = v
  return s
}

type CreateLoadBalancerListenerRequestListener struct {
  // {"en":"Load Balancer ID","zh_CN":"负载均衡ID"}
  LbId *string `json:"lbId,omitempty" xml:"lbId,omitempty" require:"true"`
  // {"en":"Listener Name","zh_CN":"监听名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Listener port, port range: 1-65535, the same protocol listener port cannot be duplicated under the same load balancer","zh_CN":"监听端口，端口范围：1-65535，同一负载均衡下，同样协议监听端口不能重复"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"Protocol, optional values: TCP/UDP","zh_CN":"协议，可选值：TCP/UDP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"Load balancing algorithm\n- rr: Round Robin\n- wrr: Weighted Round Robin\n- dh: Destination Hash\n- sh: Source Hash\n- lc: Least Connections\n- wlc: Weighted Least Connections","zh_CN":"负载均衡算法\n- rr：轮询\n- wrr：加权轮询\n- dh：目标地址散列\n- sh：源地址散列\n- lc：最少连接\n- wlc：加权最少连接"}
  Algo *string `json:"algo,omitempty" xml:"algo,omitempty" require:"true"`
  // {"en":"Enable session persistence","zh_CN":"是否开启会话保持"}
  PersistenceSession *int `json:"persistenceSession,omitempty" xml:"persistenceSession,omitempty" require:"true"`
  // {"en":"Session persistence timeout\n- Required when persistenceSession=1,\n- Cannot be specified when persistenceSession=-1\n- Value range: [1, 2678400]\n- Unit: seconds","zh_CN":"会话保持超时时间\n- persistenceSession=1时，必填，\n- persistenceSession=-1时，不能指定\n- 取值： [1, 2678400] \n- 单位：秒"}
  PersistenceTimeout *int `json:"persistenceTimeout,omitempty" xml:"persistenceTimeout,omitempty"`
  // {"en":"Auto-enable. Values: 1 (Yes), -1 (No)","zh_CN":"是否自动启用。取值：1（是），-1（否）"}
  AutoEnable *int `json:"autoEnable,omitempty" xml:"autoEnable,omitempty"`
  // {"en":"Server Pool ID","zh_CN":"服务器池ID"}
  PoolId *string `json:"poolId,omitempty" xml:"poolId,omitempty" require:"true"`
  // {"en":"Enable health check. Values: 1 (Yes), -1 (No)","zh_CN":"是否开启健康检查。取值：1（是），-1（否）"}
  HealthCheck *int `json:"healthCheck,omitempty" xml:"healthCheck,omitempty" require:"true"`
  // {"en":"Health check configuration, must be configured when healthCheck=1, cannot be configured when healthCheck=-1","zh_CN":"健康检查配置，healthCheck=1时必须配置，healthCheck=-1时不能配置"}
  HealthMonitor *CreateLoadBalancerListenerRequestListenerHealthMonitor `json:"healthMonitor,omitempty" xml:"healthMonitor,omitempty" type:"Struct"`
}

func (s CreateLoadBalancerListenerRequestListener) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerRequestListener) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerListenerRequestListener) SetLbId(v string) *CreateLoadBalancerListenerRequestListener {
  s.LbId = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetName(v string) *CreateLoadBalancerListenerRequestListener {
  s.Name = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetPort(v int) *CreateLoadBalancerListenerRequestListener {
  s.Port = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetProtocol(v string) *CreateLoadBalancerListenerRequestListener {
  s.Protocol = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetAlgo(v string) *CreateLoadBalancerListenerRequestListener {
  s.Algo = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetPersistenceSession(v int) *CreateLoadBalancerListenerRequestListener {
  s.PersistenceSession = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetPersistenceTimeout(v int) *CreateLoadBalancerListenerRequestListener {
  s.PersistenceTimeout = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetAutoEnable(v int) *CreateLoadBalancerListenerRequestListener {
  s.AutoEnable = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetPoolId(v string) *CreateLoadBalancerListenerRequestListener {
  s.PoolId = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetHealthCheck(v int) *CreateLoadBalancerListenerRequestListener {
  s.HealthCheck = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListener) SetHealthMonitor(v *CreateLoadBalancerListenerRequestListenerHealthMonitor) *CreateLoadBalancerListenerRequestListener {
  s.HealthMonitor = v
  return s
}

type CreateLoadBalancerListenerRequestListenerHealthMonitor struct {
  // {"en":"Check protocol, optional values: TCP/UDP/HTTP","zh_CN":"检查协议，可选值：TCP/UDP/HTTP"}
  HealthProtocol *string `json:"healthProtocol,omitempty" xml:"healthProtocol,omitempty" require:"true"`
  // {"en":"Check port\n- Will use the backend RS instance port for health check\n- Range: 1-65535","zh_CN":"检查端口\n- 将使用后端RS实例端口进行健康检查\n- 范围：1-65535"}
  HealthPort *int `json:"healthPort,omitempty" xml:"healthPort,omitempty"`
  // {"en":"Check timeout\n- Default is 5 seconds if not specified\n- Unit: seconds","zh_CN":"检查超时时间\n- 未指定时默认5秒\n- 单位：秒"}
  HealthTimeout *int `json:"healthTimeout,omitempty" xml:"healthTimeout,omitempty"`
  // {"en":"Check interval\n- Default is 2 seconds if not specified\n- Unit: seconds","zh_CN":"检查间隔时间\n- 未指定时默认2秒\n- 单位：秒"}
  HealthInterval *int `json:"healthInterval,omitempty" xml:"healthInterval,omitempty"`
  // {"en":"Check failure retry count\n- Default is 3 times if not specified","zh_CN":"检查失败重试次数\n- 未指定时默认3次"}
  HealthRetry *int `json:"healthRetry,omitempty" xml:"healthRetry,omitempty"`
  // {"en":"Check path. Required when healthProtocol=HTTP, cannot be specified for other protocols","zh_CN":"检查路径。healthProtocol=HTTP时，必填，其他协议时不能指定"}
  HealthPath *string `json:"healthPath,omitempty" xml:"healthPath,omitempty"`
  // {"en":"Check normal response\n- Default configuration is '200-299' when healthProtocol=HTTP if not specified, cannot be specified for other protocols\n\nParameter value is an array, example: \"healthStatusCodes\":[\"400\",\"200-299\"]","zh_CN":"检查正常响应\n- healthProtocol=HTTP时，未指定时，默认配置“200-299”，其他协议时不能指定\n\n参数值为数组，示例：\"healthStatusCodes\":[\"400\",\"200-299\"]"}
  HealthStatusCodes []*string `json:"healthStatusCodes,omitempty" xml:"healthStatusCodes,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerListenerRequestListenerHealthMonitor) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerRequestListenerHealthMonitor) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthProtocol(v string) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthProtocol = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthPort(v int) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthPort = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthTimeout(v int) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthTimeout = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthInterval(v int) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthInterval = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthRetry(v int) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthRetry = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthPath(v string) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthPath = &v
  return s
}

func (s *CreateLoadBalancerListenerRequestListenerHealthMonitor) SetHealthStatusCodes(v []*string) *CreateLoadBalancerListenerRequestListenerHealthMonitor {
  s.HealthStatusCodes = v
  return s
}

type CreateLoadBalancerListenerRequestHeader struct {
}

func (s CreateLoadBalancerListenerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerRequestHeader) GoString() string {
  return s.String()
}

type CreateLoadBalancerListenerPaths struct {
}

func (s CreateLoadBalancerListenerPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerPaths) GoString() string {
  return s.String()
}

type CreateLoadBalancerListenerParameters struct {
}

func (s CreateLoadBalancerListenerParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerParameters) GoString() string {
  return s.String()
}

type CreateLoadBalancerListenerResponse struct {
  // {"en":"Function code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data","zh_CN":"数据"}
  Data *CreateLoadBalancerListenerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateLoadBalancerListenerResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerResponse) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerListenerResponse) SetCode(v string) *CreateLoadBalancerListenerResponse {
  s.Code = &v
  return s
}

func (s *CreateLoadBalancerListenerResponse) SetData(v *CreateLoadBalancerListenerResponseData) *CreateLoadBalancerListenerResponse {
  s.Data = v
  return s
}

func (s *CreateLoadBalancerListenerResponse) SetMessage(v string) *CreateLoadBalancerListenerResponse {
  s.Message = &v
  return s
}

type CreateLoadBalancerListenerResponseData struct {
  // {"en":"Listener ID","zh_CN":"监听ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateLoadBalancerListenerResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerResponseData) GoString() string {
  return s.String()
}

func (s *CreateLoadBalancerListenerResponseData) SetId(v string) *CreateLoadBalancerListenerResponseData {
  s.Id = &v
  return s
}

type CreateLoadBalancerListenerResponseHeader struct {
}

func (s CreateLoadBalancerListenerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateLoadBalancerListenerResponseHeader) GoString() string {
  return s.String()
}




