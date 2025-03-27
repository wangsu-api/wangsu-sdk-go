package clusterinfor

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ListClusterRequest struct {
}

func (s ListClusterRequest) String() string {
  return tea.Prettify(s)
}

func (s ListClusterRequest) GoString() string {
  return s.String()
}

type ListClusterResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"cluster", "zh_CN":"集群列表"}
  Data []*ListClusterCluster `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListClusterResponse) String() string {
  return tea.Prettify(s)
}

func (s ListClusterResponse) GoString() string {
  return s.String()
}

func (s *ListClusterResponse) SetCode(v int64) *ListClusterResponse {
  s.Code = &v
  return s
}

func (s *ListClusterResponse) SetMsg(v string) *ListClusterResponse {
  s.Msg = &v
  return s
}

func (s *ListClusterResponse) SetRequestId(v string) *ListClusterResponse {
  s.RequestId = &v
  return s
}

func (s *ListClusterResponse) SetData(v []*ListClusterCluster) *ListClusterResponse {
  s.Data = v
  return s
}

type ListClusterPaths struct {
}

func (s ListClusterPaths) String() string {
  return tea.Prettify(s)
}

func (s ListClusterPaths) GoString() string {
  return s.String()
}

type ListClusterParameters struct {
  // {"en":"labelSelector", "zh_CN":"labelSelector"}
  LabelSelector *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListClusterParameters) String() string {
  return tea.Prettify(s)
}

func (s ListClusterParameters) GoString() string {
  return s.String()
}

func (s *ListClusterParameters) SetLabelSelector(v string) *ListClusterParameters {
  s.LabelSelector = &v
  return s
}

type ListClusterRequestHeader struct {
}

func (s ListClusterRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListClusterRequestHeader) GoString() string {
  return s.String()
}

type ListClusterResponseHeader struct {
}

func (s ListClusterResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListClusterResponseHeader) GoString() string {
  return s.String()
}

type ListClusterCluster struct {
  // {"en":"chinese name of the clustere", "zh_CN":"集群中文名"}
  ClusterCnName *string `json:"clusterCnName,omitempty" xml:"clusterCnName,omitempty" require:"true"`
  // {"en":"cluster name", "zh_CN":"集群名"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty" require:"true"`
  // {"en":"cluster alias", "zh_CN":"集群别名"}
  ClusterAlias *string `json:"clusterAlias,omitempty" xml:"clusterAlias,omitempty" require:"true"`
  // {"en":"country name", "zh_CN":"集群所在国家名称"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"area", "zh_CN":"集群所在区域"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"province name", "zh_CN":"集群所在省份名称"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"city name", "zh_CN":"集群所在城市名称"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"集群支持的运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"country name", "zh_CN":"集群所在国家英文名"}
  CountryEn *string `json:"countryEn,omitempty" xml:"countryEn,omitempty" require:"true"`
  // {"en":"the English name of area", "zh_CN":"集群所在区域英文名"}
  AreaEn *string `json:"areaEn,omitempty" xml:"areaEn,omitempty" require:"true"`
  // {"en":"the English name of province", "zh_CN":"集群所在省份英文名"}
  ProvinceEn *string `json:"provinceEn,omitempty" xml:"provinceEn,omitempty" require:"true"`
  // {"en":"the English name of city", "zh_CN":"集群所在城市英文名"}
  CityEn *string `json:"cityEn,omitempty" xml:"cityEn,omitempty" require:"true"`
  // {"en":"cluster pool type(0,1:Universal,2: GPU,3: Storage)", "zh_CN":"集群业务类型(0,1:通用型,2: gpu型,3: 存储型)"}
  PoolType *int64 `json:"poolType,omitempty" xml:"poolType,omitempty" require:"true"`
}

func (s ListClusterCluster) String() string {
  return tea.Prettify(s)
}

func (s ListClusterCluster) GoString() string {
  return s.String()
}

func (s *ListClusterCluster) SetClusterCnName(v string) *ListClusterCluster {
  s.ClusterCnName = &v
  return s
}

func (s *ListClusterCluster) SetClusterName(v string) *ListClusterCluster {
  s.ClusterName = &v
  return s
}

func (s *ListClusterCluster) SetClusterAlias(v string) *ListClusterCluster {
  s.ClusterAlias = &v
  return s
}

func (s *ListClusterCluster) SetCountry(v string) *ListClusterCluster {
  s.Country = &v
  return s
}

func (s *ListClusterCluster) SetArea(v string) *ListClusterCluster {
  s.Area = &v
  return s
}

func (s *ListClusterCluster) SetProvince(v string) *ListClusterCluster {
  s.Province = &v
  return s
}

func (s *ListClusterCluster) SetCity(v string) *ListClusterCluster {
  s.City = &v
  return s
}

func (s *ListClusterCluster) SetIsp(v string) *ListClusterCluster {
  s.Isp = &v
  return s
}

func (s *ListClusterCluster) SetCountryEn(v string) *ListClusterCluster {
  s.CountryEn = &v
  return s
}

func (s *ListClusterCluster) SetAreaEn(v string) *ListClusterCluster {
  s.AreaEn = &v
  return s
}

func (s *ListClusterCluster) SetProvinceEn(v string) *ListClusterCluster {
  s.ProvinceEn = &v
  return s
}

func (s *ListClusterCluster) SetCityEn(v string) *ListClusterCluster {
  s.CityEn = &v
  return s
}

func (s *ListClusterCluster) SetPoolType(v int64) *ListClusterCluster {
  s.PoolType = &v
  return s
}




type ListPoolsRequest struct {
  // {"en":"resource request", "zh_CN":"集群拥有的资源"}
  Resources []*ListPoolsPoolResource `json:"resources,omitempty" xml:"resources,omitempty" require:"true" type:"Repeated"`
}

func (s ListPoolsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsRequest) GoString() string {
  return s.String()
}

func (s *ListPoolsRequest) SetResources(v []*ListPoolsPoolResource) *ListPoolsRequest {
  s.Resources = v
  return s
}

type ListPoolsPoolResource struct {
  // {"en":"resource group", "zh_CN":"资源分组"}
  Group *string `json:"group,omitempty" xml:"group,omitempty" require:"true"`
  // {"en":"resource type", "zh_CN":"资源类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s ListPoolsPoolResource) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsPoolResource) GoString() string {
  return s.String()
}

func (s *ListPoolsPoolResource) SetGroup(v string) *ListPoolsPoolResource {
  s.Group = &v
  return s
}

func (s *ListPoolsPoolResource) SetType(v string) *ListPoolsPoolResource {
  s.Type = &v
  return s
}

type ListPoolsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"ListPoolsPoolClusterListResp", "zh_CN":"ListPoolsPoolClusterListResp"}
  Data *ListPoolsPoolClusterListResp `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListPoolsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsResponse) GoString() string {
  return s.String()
}

func (s *ListPoolsResponse) SetCode(v int64) *ListPoolsResponse {
  s.Code = &v
  return s
}

func (s *ListPoolsResponse) SetMsg(v string) *ListPoolsResponse {
  s.Msg = &v
  return s
}

func (s *ListPoolsResponse) SetRequestId(v string) *ListPoolsResponse {
  s.RequestId = &v
  return s
}

func (s *ListPoolsResponse) SetData(v *ListPoolsPoolClusterListResp) *ListPoolsResponse {
  s.Data = v
  return s
}

type ListPoolsPaths struct {
}

func (s ListPoolsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsPaths) GoString() string {
  return s.String()
}

type ListPoolsParameters struct {
}

func (s ListPoolsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsParameters) GoString() string {
  return s.String()
}

type ListPoolsRequestHeader struct {
}

func (s ListPoolsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsRequestHeader) GoString() string {
  return s.String()
}

type ListPoolsResponseHeader struct {
}

func (s ListPoolsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsResponseHeader) GoString() string {
  return s.String()
}

type ListPoolsPoolClusterListResp struct {
  // {"en":"poolCluster total", "zh_CN":"调度池总数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"poolCluster list", "zh_CN":"调度池列表"}
  Infos []*ListPoolsPoolClusterInfo `json:"infos,omitempty" xml:"infos,omitempty" require:"true" type:"Repeated"`
}

func (s ListPoolsPoolClusterListResp) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsPoolClusterListResp) GoString() string {
  return s.String()
}

func (s *ListPoolsPoolClusterListResp) SetTotal(v int64) *ListPoolsPoolClusterListResp {
  s.Total = &v
  return s
}

func (s *ListPoolsPoolClusterListResp) SetInfos(v []*ListPoolsPoolClusterInfo) *ListPoolsPoolClusterListResp {
  s.Infos = v
  return s
}

type ListPoolsPoolClusterInfo struct {
  // {"en":"poolCluster name", "zh_CN":"调度池名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"poolCluster cn name", "zh_CN":"调度池中文名"}
  NameCn *string `json:"nameCn,omitempty" xml:"nameCn,omitempty" require:"true"`
  // {"en":"poolCluster alias", "zh_CN":"调度池别名"}
  Alias *string `json:"alias,omitempty" xml:"alias,omitempty" require:"true"`
  // {"en":"poolCluster type", "zh_CN":"调度池类别"}
  Type *int64 `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"poolCluster id", "zh_CN":"调度池ID"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"country", "zh_CN":"国家"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"area", "zh_CN":"地区"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"city", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"isp", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
}

func (s ListPoolsPoolClusterInfo) String() string {
  return tea.Prettify(s)
}

func (s ListPoolsPoolClusterInfo) GoString() string {
  return s.String()
}

func (s *ListPoolsPoolClusterInfo) SetName(v string) *ListPoolsPoolClusterInfo {
  s.Name = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetNameCn(v string) *ListPoolsPoolClusterInfo {
  s.NameCn = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetAlias(v string) *ListPoolsPoolClusterInfo {
  s.Alias = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetType(v int64) *ListPoolsPoolClusterInfo {
  s.Type = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetId(v int64) *ListPoolsPoolClusterInfo {
  s.Id = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetCountry(v string) *ListPoolsPoolClusterInfo {
  s.Country = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetArea(v string) *ListPoolsPoolClusterInfo {
  s.Area = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetProvince(v string) *ListPoolsPoolClusterInfo {
  s.Province = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetCity(v string) *ListPoolsPoolClusterInfo {
  s.City = &v
  return s
}

func (s *ListPoolsPoolClusterInfo) SetIsp(v string) *ListPoolsPoolClusterInfo {
  s.Isp = &v
  return s
}




type QueryEdgeContainerBandwidthServiceRequest struct {
  // {"en": "start_time", "zh_CN": "开始时间戳(毫秒)"}
  Start_time *int64 `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
  // {"en": "end_time", "zh_CN": "结束时间戳(毫秒)"}
  End_time *int64 `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
  // {"en": "cluster", "zh_CN": "集群"}
  Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
}

func (s QueryEdgeContainerBandwidthServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeContainerBandwidthServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryEdgeContainerBandwidthServiceRequest) SetStart_time(v int64) *QueryEdgeContainerBandwidthServiceRequest {
  s.Start_time = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceRequest) SetEnd_time(v int64) *QueryEdgeContainerBandwidthServiceRequest {
  s.End_time = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceRequest) SetCluster(v string) *QueryEdgeContainerBandwidthServiceRequest {
  s.Cluster = &v
  return s
}

type QueryEdgeContainerBandwidthServiceResponse struct {
  // {"en": "cluster", "zh_CN": "集群名称"}
  Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty" require:"true"`
  Data_points []*string `json:"data_points,omitempty" xml:"data_points,omitempty" require:"true" type:"Repeated"`
  // {"en": "upstream or downstream", "zh_CN": "请求方向,流入/流出"}
  Bw_direction *string `json:"bw_direction,omitempty" xml:"bw_direction,omitempty" require:"true"`
  Data_point []*string `json:"data_point,omitempty" xml:"data_point,omitempty" require:"true" type:"Repeated"`
  // {"en": "bandwidth", "zh_CN": "每5分钟的带宽平均值"}
  Max *float64 `json:"max,omitempty" xml:"max,omitempty" require:"true"`
  // {"en": "timestamp", "zh_CN": "毫秒级别时间戳"}
  Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en": "cpu_size", "zh_CN": "时间范围内CPU使⽤最⼤值"}
  Cpu *int `json:"cpu,omitempty" xml:"cpu,omitempty" require:"true"`
  // {"en": "memory_size", "zh_CN": "时间范围内内存使⽤最⼤值"}
  Memory *int `json:"memory,omitempty" xml:"memory,omitempty" require:"true"`
  // {"en": "pvc", "zh_CN": "时间范围内使⽤最⼤值"}
  Pvc *int `json:"pvc,omitempty" xml:"pvc,omitempty" require:"true"`
}

func (s QueryEdgeContainerBandwidthServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeContainerBandwidthServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetCluster(v string) *QueryEdgeContainerBandwidthServiceResponse {
  s.Cluster = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetData_points(v []*string) *QueryEdgeContainerBandwidthServiceResponse {
  s.Data_points = v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetBw_direction(v string) *QueryEdgeContainerBandwidthServiceResponse {
  s.Bw_direction = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetData_point(v []*string) *QueryEdgeContainerBandwidthServiceResponse {
  s.Data_point = v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetMax(v float64) *QueryEdgeContainerBandwidthServiceResponse {
  s.Max = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetTimestamp(v int64) *QueryEdgeContainerBandwidthServiceResponse {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetCpu(v int) *QueryEdgeContainerBandwidthServiceResponse {
  s.Cpu = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetMemory(v int) *QueryEdgeContainerBandwidthServiceResponse {
  s.Memory = &v
  return s
}

func (s *QueryEdgeContainerBandwidthServiceResponse) SetPvc(v int) *QueryEdgeContainerBandwidthServiceResponse {
  s.Pvc = &v
  return s
}

type QueryEdgeContainerBandwidthServicePaths struct {
}

func (s QueryEdgeContainerBandwidthServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeContainerBandwidthServicePaths) GoString() string {
  return s.String()
}

type QueryEdgeContainerBandwidthServiceParameters struct {
}

func (s QueryEdgeContainerBandwidthServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeContainerBandwidthServiceParameters) GoString() string {
  return s.String()
}

type QueryEdgeContainerBandwidthServiceRequestHeader struct {
}

func (s QueryEdgeContainerBandwidthServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeContainerBandwidthServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeContainerBandwidthServiceResponseHeader struct {
}

func (s QueryEdgeContainerBandwidthServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeContainerBandwidthServiceResponseHeader) GoString() string {
  return s.String()
}




