package devicemanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type InvitePushRequest struct {
  // {"en":"The national standard id of the device", "zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"channelId", "zh_CN":"通道Id"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"Callback notification URL", "zh_CN":"回调通知URL"}
  NoticeUrl *string `json:"noticeUrl,omitempty" xml:"noticeUrl,omitempty"`
  // {"en":"Custom request id, customers need to ensure uniqueness", "zh_CN":"自定义请求id, 客户需要保证唯一性"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
}

func (s InvitePushRequest) String() string {
  return tea.Prettify(s)
}

func (s InvitePushRequest) GoString() string {
  return s.String()
}

func (s *InvitePushRequest) SetDeviceId(v string) *InvitePushRequest {
  s.DeviceId = &v
  return s
}

func (s *InvitePushRequest) SetChannelId(v string) *InvitePushRequest {
  s.ChannelId = &v
  return s
}

func (s *InvitePushRequest) SetNoticeUrl(v string) *InvitePushRequest {
  s.NoticeUrl = &v
  return s
}

func (s *InvitePushRequest) SetTransNo(v string) *InvitePushRequest {
  s.TransNo = &v
  return s
}

type InvitePushResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s InvitePushResponse) String() string {
  return tea.Prettify(s)
}

func (s InvitePushResponse) GoString() string {
  return s.String()
}

func (s *InvitePushResponse) SetCode(v int32) *InvitePushResponse {
  s.Code = &v
  return s
}

func (s *InvitePushResponse) SetMessage(v string) *InvitePushResponse {
  s.Message = &v
  return s
}

func (s *InvitePushResponse) SetData(v string) *InvitePushResponse {
  s.Data = &v
  return s
}

type InvitePushPaths struct {
}

func (s InvitePushPaths) String() string {
  return tea.Prettify(s)
}

func (s InvitePushPaths) GoString() string {
  return s.String()
}

type InvitePushParameters struct {
}

func (s InvitePushParameters) String() string {
  return tea.Prettify(s)
}

func (s InvitePushParameters) GoString() string {
  return s.String()
}

type InvitePushRequestHeader struct {
}

func (s InvitePushRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InvitePushRequestHeader) GoString() string {
  return s.String()
}

type InvitePushResponseHeader struct {
}

func (s InvitePushResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InvitePushResponseHeader) GoString() string {
  return s.String()
}




type StopPushRequest struct {
  // {"en":"The national standard id of the device", "zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"channelId", "zh_CN":"通道Id"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"Callback notification URL", "zh_CN":"回调通知URL"}
  NoticeUrl *string `json:"noticeUrl,omitempty" xml:"noticeUrl,omitempty"`
  // {"en":"Custom request id, customers need to ensure uniqueness", "zh_CN":"自定义请求id, 客户需要保证唯一性"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
}

func (s StopPushRequest) String() string {
  return tea.Prettify(s)
}

func (s StopPushRequest) GoString() string {
  return s.String()
}

func (s *StopPushRequest) SetDeviceId(v string) *StopPushRequest {
  s.DeviceId = &v
  return s
}

func (s *StopPushRequest) SetChannelId(v string) *StopPushRequest {
  s.ChannelId = &v
  return s
}

func (s *StopPushRequest) SetNoticeUrl(v string) *StopPushRequest {
  s.NoticeUrl = &v
  return s
}

func (s *StopPushRequest) SetTransNo(v string) *StopPushRequest {
  s.TransNo = &v
  return s
}

type StopPushResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s StopPushResponse) String() string {
  return tea.Prettify(s)
}

func (s StopPushResponse) GoString() string {
  return s.String()
}

func (s *StopPushResponse) SetCode(v int32) *StopPushResponse {
  s.Code = &v
  return s
}

func (s *StopPushResponse) SetMessage(v string) *StopPushResponse {
  s.Message = &v
  return s
}

func (s *StopPushResponse) SetData(v string) *StopPushResponse {
  s.Data = &v
  return s
}

type StopPushPaths struct {
}

func (s StopPushPaths) String() string {
  return tea.Prettify(s)
}

func (s StopPushPaths) GoString() string {
  return s.String()
}

type StopPushParameters struct {
}

func (s StopPushParameters) String() string {
  return tea.Prettify(s)
}

func (s StopPushParameters) GoString() string {
  return s.String()
}

type StopPushRequestHeader struct {
}

func (s StopPushRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StopPushRequestHeader) GoString() string {
  return s.String()
}

type StopPushResponseHeader struct {
}

func (s StopPushResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StopPushResponseHeader) GoString() string {
  return s.String()
}




type GetDevicesListRequest struct {
  // {"en":"Space name. Fuzzy search is not supported.","zh_CN":"空间名称。不支持模糊搜索"}
  SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
  // {"en":"Page number, default is first page","zh_CN":"第几页，默认第一页"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Device name. Supports fuzzy search","zh_CN":"设备名称。支持模糊搜索"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Specify the paging size. The default value is 10, and the maximum value is 50.","zh_CN":"分页大小。默认10，最大50"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Relay node ID","zh_CN":"父节点ID"}
  ParentNodeId *string `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty"`
  // {"en":"Device type: 1: IPC 2: NVR","zh_CN":"设备类型：1：IPC   2：NVR"}
  Type *int `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"The national standard id of the device. Fuzzy search is not supported.","zh_CN":"设备国标id。不支持模糊搜索。"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
  // {"en":"Device status: 0: Offline 1: Online","zh_CN":"设备状态： 0:离线    1：在线"}
  Status *int `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDevicesListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListRequest) GoString() string {
  return s.String()
}

func (s *GetDevicesListRequest) SetSpaceName(v string) *GetDevicesListRequest {
  s.SpaceName = &v
  return s
}

func (s *GetDevicesListRequest) SetPageIndex(v int) *GetDevicesListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetDevicesListRequest) SetName(v string) *GetDevicesListRequest {
  s.Name = &v
  return s
}

func (s *GetDevicesListRequest) SetPageSize(v int) *GetDevicesListRequest {
  s.PageSize = &v
  return s
}

func (s *GetDevicesListRequest) SetParentNodeId(v string) *GetDevicesListRequest {
  s.ParentNodeId = &v
  return s
}

func (s *GetDevicesListRequest) SetType(v int) *GetDevicesListRequest {
  s.Type = &v
  return s
}

func (s *GetDevicesListRequest) SetDeviceId(v string) *GetDevicesListRequest {
  s.DeviceId = &v
  return s
}

func (s *GetDevicesListRequest) SetStatus(v int) *GetDevicesListRequest {
  s.Status = &v
  return s
}

type GetDevicesListRequestHeader struct {
}

func (s GetDevicesListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListRequestHeader) GoString() string {
  return s.String()
}

type GetDevicesListPaths struct {
}

func (s GetDevicesListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListPaths) GoString() string {
  return s.String()
}

type GetDevicesListParameters struct {
}

func (s GetDevicesListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListParameters) GoString() string {
  return s.String()
}

type GetDevicesListResponse struct {
  // {"en":"Result status code, 0 indicates success","zh_CN":"结果状态码，0为成功"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return data","zh_CN":"返回数据"}
  Data *GetDevicesListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Return message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetDevicesListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListResponse) GoString() string {
  return s.String()
}

func (s *GetDevicesListResponse) SetCode(v int) *GetDevicesListResponse {
  s.Code = &v
  return s
}

func (s *GetDevicesListResponse) SetData(v *GetDevicesListResponseData) *GetDevicesListResponse {
  s.Data = v
  return s
}

func (s *GetDevicesListResponse) SetMessage(v string) *GetDevicesListResponse {
  s.Message = &v
  return s
}

type GetDevicesListResponseData struct {
  // {"en":"Total number of devices","zh_CN":"总设备数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Page number,","zh_CN":"第几页"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"Specify the paging size","zh_CN":"分页大小"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Device List","zh_CN":"设备列表"}
  Rows []*GetDevicesListResponseDataRows `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
}

func (s GetDevicesListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListResponseData) GoString() string {
  return s.String()
}

func (s *GetDevicesListResponseData) SetTotal(v int) *GetDevicesListResponseData {
  s.Total = &v
  return s
}

func (s *GetDevicesListResponseData) SetPageIndex(v int) *GetDevicesListResponseData {
  s.PageIndex = &v
  return s
}

func (s *GetDevicesListResponseData) SetPageSize(v int) *GetDevicesListResponseData {
  s.PageSize = &v
  return s
}

func (s *GetDevicesListResponseData) SetRows(v []*GetDevicesListResponseDataRows) *GetDevicesListResponseData {
  s.Rows = v
  return s
}

type GetDevicesListResponseDataRows struct     {
  // {"en":"Device Address","zh_CN":"设备地址"}
  Address *string `json:"address,omitempty" xml:"address,omitempty" require:"true"`
  // {"en":"Whether to enable automatic invitation streaming. true: Enable automatic invitation streaming, false: Disable automatic invitation streaming","zh_CN":"是否启动自动邀请推流。1:启动自动邀请推流，0:不启动自动邀请推流"}
  IsAutoPush *int `json:"isAutoPush,omitempty" xml:"isAutoPush,omitempty" require:"true"`
  // {"en":"latitude","zh_CN":"纬度"}
  Latitude *GetDevicesListResponseDataRowsLatitude `json:"latitude,omitempty" xml:"latitude,omitempty" require:"true" type:"Struct"`
  // {"en":"Device Description","zh_CN":"设备描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Number of channels","zh_CN":"通道数量"}
  ChannelNum *int `json:"channelNum,omitempty" xml:"channelNum,omitempty" require:"true"`
  // {"en":"Relay node ID","zh_CN":"父节点ID"}
  ParentNodeId *string `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty" require:"true"`
  // {"en":"Streaming protocols","zh_CN":"流协议"}
  StreamProtocol *string `json:"streamProtocol,omitempty" xml:"streamProtocol,omitempty" require:"true"`
  // {"en":"Device type: 1: IPC 2: NVR","zh_CN":"设备类型：1：IPC   2：NVR"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"National standard id of the device","zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Equipment Manufacturer","zh_CN":"设备厂商"}
  Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty" require:"true"`
  // {"en":"SIP server ID","zh_CN":"sip服务器id"}
  SipId *string `json:"sipId,omitempty" xml:"sipId,omitempty" require:"true"`
  // {"en":"Space Name","zh_CN":"所属空间名称"}
  SpaceName *int `json:"spaceName,omitempty" xml:"spaceName,omitempty" require:"true"`
  // {"en":"SIP Server Address","zh_CN":"sip服务器地址"}
  SipServerAddress *string `json:"sipServerAddress,omitempty" xml:"sipServerAddress,omitempty" require:"true"`
  // {"en":"Stream Type:\ndefault\nstream:0\nstream:1\nstreamnumber:0\nstreamnumber:1\nstreamprofile:0\nstreamprofile:1\nstreamMode:MAIN\nstreamMode:SUB","zh_CN":"码流类型:\ndefault\nstream:0\nstream:1\nstreamnumber:0\nstreamnumber:1\nstreamprofile:0\nstreamprofile:1\nstreamMode:MAIN\nstreamMode:SUB"}
  StreamType *string `json:"streamType,omitempty" xml:"streamType,omitempty" require:"true"`
  // {"en":"Device creation time (timestamp)","zh_CN":"设备创建时间（时间戳）"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Device name","zh_CN":"设备名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Access Agreement 1: GB28181-2016","zh_CN":"接入协议 1：GB28181-2016"}
  AccessProtocol *int `json:"accessProtocol,omitempty" xml:"accessProtocol,omitempty" require:"true"`
  // {"en":"SIP server port","zh_CN":"sip服务器端口"}
  SipServerPort *int `json:"sipServerPort,omitempty" xml:"sipServerPort,omitempty" require:"true"`
  // {"en":"The status of the device.","zh_CN":"设备状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"longitude","zh_CN":"经度"}
  Longitude *GetDevicesListResponseDataRowsLongitude `json:"longitude,omitempty" xml:"longitude,omitempty" require:"true" type:"Struct"`
}

func (s GetDevicesListResponseDataRows) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListResponseDataRows) GoString() string {
  return s.String()
}

func (s *GetDevicesListResponseDataRows) SetAddress(v string) *GetDevicesListResponseDataRows {
  s.Address = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetIsAutoPush(v int) *GetDevicesListResponseDataRows {
  s.IsAutoPush = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetLatitude(v *GetDevicesListResponseDataRowsLatitude) *GetDevicesListResponseDataRows {
  s.Latitude = v
  return s
}

func (s *GetDevicesListResponseDataRows) SetDescription(v string) *GetDevicesListResponseDataRows {
  s.Description = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetChannelNum(v int) *GetDevicesListResponseDataRows {
  s.ChannelNum = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetParentNodeId(v string) *GetDevicesListResponseDataRows {
  s.ParentNodeId = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetStreamProtocol(v string) *GetDevicesListResponseDataRows {
  s.StreamProtocol = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetType(v int) *GetDevicesListResponseDataRows {
  s.Type = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetDeviceId(v string) *GetDevicesListResponseDataRows {
  s.DeviceId = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetManufacturer(v string) *GetDevicesListResponseDataRows {
  s.Manufacturer = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetSipId(v string) *GetDevicesListResponseDataRows {
  s.SipId = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetSpaceName(v int) *GetDevicesListResponseDataRows {
  s.SpaceName = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetSipServerAddress(v string) *GetDevicesListResponseDataRows {
  s.SipServerAddress = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetStreamType(v string) *GetDevicesListResponseDataRows {
  s.StreamType = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetCreateTime(v int64) *GetDevicesListResponseDataRows {
  s.CreateTime = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetName(v string) *GetDevicesListResponseDataRows {
  s.Name = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetAccessProtocol(v int) *GetDevicesListResponseDataRows {
  s.AccessProtocol = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetSipServerPort(v int) *GetDevicesListResponseDataRows {
  s.SipServerPort = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetStatus(v int) *GetDevicesListResponseDataRows {
  s.Status = &v
  return s
}

func (s *GetDevicesListResponseDataRows) SetLongitude(v *GetDevicesListResponseDataRowsLongitude) *GetDevicesListResponseDataRows {
  s.Longitude = v
  return s
}

type GetDevicesListResponseDataRowsLatitude struct {
}

func (s GetDevicesListResponseDataRowsLatitude) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListResponseDataRowsLatitude) GoString() string {
  return s.String()
}

type GetDevicesListResponseDataRowsLongitude struct {
}

func (s GetDevicesListResponseDataRowsLongitude) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListResponseDataRowsLongitude) GoString() string {
  return s.String()
}

type GetDevicesListResponseHeader struct {
}

func (s GetDevicesListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDevicesListResponseHeader) GoString() string {
  return s.String()
}




type ChannelControlRequest struct {
  // {"en":"Device GB28181 Id", "zh_CN":"设备国标ID"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Channel GB28181 Id", "zh_CN":"通道国标ID"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"true: channel is disabled
  // false: channel is enabled", "zh_CN":"true:通道禁用
  // 
  // false：通道启用"}
  Disable *bool `json:"disable,omitempty" xml:"disable,omitempty" require:"true"`
}

func (s ChannelControlRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelControlRequest) GoString() string {
  return s.String()
}

func (s *ChannelControlRequest) SetDeviceId(v string) *ChannelControlRequest {
  s.DeviceId = &v
  return s
}

func (s *ChannelControlRequest) SetChannelId(v string) *ChannelControlRequest {
  s.ChannelId = &v
  return s
}

func (s *ChannelControlRequest) SetDisable(v bool) *ChannelControlRequest {
  s.Disable = &v
  return s
}

type ChannelControlResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ChannelControlResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelControlResponse) GoString() string {
  return s.String()
}

func (s *ChannelControlResponse) SetCode(v int32) *ChannelControlResponse {
  s.Code = &v
  return s
}

func (s *ChannelControlResponse) SetMessage(v string) *ChannelControlResponse {
  s.Message = &v
  return s
}

type ChannelControlPaths struct {
}

func (s ChannelControlPaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelControlPaths) GoString() string {
  return s.String()
}

type ChannelControlParameters struct {
}

func (s ChannelControlParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelControlParameters) GoString() string {
  return s.String()
}

type ChannelControlRequestHeader struct {
}

func (s ChannelControlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelControlRequestHeader) GoString() string {
  return s.String()
}

type ChannelControlResponseHeader struct {
}

func (s ChannelControlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelControlResponseHeader) GoString() string {
  return s.String()
}




type GetStreamUrlRequest struct {
  // {"en":"deviceId", "zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"channelId", "zh_CN":"通道Id"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
}

func (s GetStreamUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlRequest) GoString() string {
  return s.String()
}

func (s *GetStreamUrlRequest) SetDeviceId(v string) *GetStreamUrlRequest {
  s.DeviceId = &v
  return s
}

func (s *GetStreamUrlRequest) SetChannelId(v string) *GetStreamUrlRequest {
  s.ChannelId = &v
  return s
}

type GetStreamUrlResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetStreamUrlData []*GetStreamUrlData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetStreamUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlResponse) GoString() string {
  return s.String()
}

func (s *GetStreamUrlResponse) SetCode(v int32) *GetStreamUrlResponse {
  s.Code = &v
  return s
}

func (s *GetStreamUrlResponse) SetMessage(v string) *GetStreamUrlResponse {
  s.Message = &v
  return s
}

func (s *GetStreamUrlResponse) SetData(v []*GetStreamUrlData) *GetStreamUrlResponse {
  s.GetStreamUrlData = v
  return s
}

type GetStreamUrlData struct {
  // {"en":"flv url", "zh_CN":"flv拉流地址"}
  Flv *string `json:"flv,omitempty" xml:"flv,omitempty" require:"true"`
  // {"en":"rtmp url", "zh_CN":"rtmp拉流地址"}
  Rtmp *string `json:"rtmp,omitempty" xml:"rtmp,omitempty" require:"true"`
  // {"en":"hls url", "zh_CN":"hls拉流地址"}
  Hls *string `json:"hls,omitempty" xml:"hls,omitempty" require:"true"`
  // {"en":"webrtc url", "zh_CN":"webrtc拉流地址"}
  Webrtc *string `json:"webrtc,omitempty" xml:"webrtc,omitempty" require:"true"`
}

func (s GetStreamUrlData) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlData) GoString() string {
  return s.String()
}

func (s *GetStreamUrlData) SetFlv(v string) *GetStreamUrlData {
  s.Flv = &v
  return s
}

func (s *GetStreamUrlData) SetRtmp(v string) *GetStreamUrlData {
  s.Rtmp = &v
  return s
}

func (s *GetStreamUrlData) SetHls(v string) *GetStreamUrlData {
  s.Hls = &v
  return s
}

func (s *GetStreamUrlData) SetWebrtc(v string) *GetStreamUrlData {
  s.Webrtc = &v
  return s
}

type GetStreamUrlPaths struct {
}

func (s GetStreamUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlPaths) GoString() string {
  return s.String()
}

type GetStreamUrlParameters struct {
}

func (s GetStreamUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlParameters) GoString() string {
  return s.String()
}

type GetStreamUrlRequestHeader struct {
}

func (s GetStreamUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlRequestHeader) GoString() string {
  return s.String()
}

type GetStreamUrlResponseHeader struct {
}

func (s GetStreamUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetStreamUrlResponseHeader) GoString() string {
  return s.String()
}




type CreateDeviceRequest struct {
  // {"en":"Device Description","zh_CN":"设备地址"}
  Address *string `json:"address,omitempty" xml:"address,omitempty"`
  // {"en":"Whether to enable automatic invitation streaming. true: Enable automatic invitation streaming, false: Disable automatic invitation streaming","zh_CN":"是否启动自动邀请推流。1:启动自动邀请推流，0:不启动自动邀请推流"}
  IsAutoPush *int `json:"isAutoPush,omitempty" xml:"isAutoPush,omitempty"`
  // {"en":"Latitude","zh_CN":"纬度"}
  Latitude *int64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
  // {"en":"Device Address","zh_CN":"设备描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Organization tree directory node id. If not filled in, it will be added to the root directory by default","zh_CN":"组织树目录节点id。不填默认新增到根目录下"}
  ParentNodeId *string `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty"`
  // {"en":"Streaming protocols, Do not pass the default TCP mode to stream pushing. If you need to use UDP to stream pushing, carry this field The values are as follows: TCP UDP","zh_CN":"流协议，  不传默认TCP方式推流，如果需要使用UDP推流则携带这个字段  取值如下：  TCP  UDP"}
  StreamProtocol *string `json:"streamProtocol,omitempty" xml:"streamProtocol,omitempty"`
  // {"en":"Device type: 1: IPC 2: NVR","zh_CN":"设备类型：1：IPC   2：NVR"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"The national standard id of the device","zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Equipment Manufacturer","zh_CN":"设备厂商"}
  Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
  // {"en":"Space name","zh_CN":"所属空间名称"}
  SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty" require:"true"`
  // {"en":"Device registration password","zh_CN":"设备注册密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty" require:"true"`
  // {"en":"When not transmitting, the code stream type is default\nStream Type:\ndefault\nstream:0\nstream:1\nstreamnumber:0\nstreamnumber:1\nstreamprofile:0\nstreamprofile:1\nstreamMode:MAIN\nstreamMode:SUB","zh_CN":"不传时码流类型为default\n码流类型:\ndefault\nstream:0\nstream:1\nstreamnumber:0\nstreamnumber:1\nstreamprofile:0\nstreamprofile:1\nstreamMode:MAIN\nstreamMode:SUB"}
  StreamType *string `json:"streamType,omitempty" xml:"streamType,omitempty"`
  // {"en":"Device Name","zh_CN":"设备名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Longituder","zh_CN":"经度"}
  Longitude *int64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s CreateDeviceRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateDeviceRequest) GoString() string {
  return s.String()
}

func (s *CreateDeviceRequest) SetAddress(v string) *CreateDeviceRequest {
  s.Address = &v
  return s
}

func (s *CreateDeviceRequest) SetIsAutoPush(v int) *CreateDeviceRequest {
  s.IsAutoPush = &v
  return s
}

func (s *CreateDeviceRequest) SetLatitude(v int64) *CreateDeviceRequest {
  s.Latitude = &v
  return s
}

func (s *CreateDeviceRequest) SetDescription(v string) *CreateDeviceRequest {
  s.Description = &v
  return s
}

func (s *CreateDeviceRequest) SetParentNodeId(v string) *CreateDeviceRequest {
  s.ParentNodeId = &v
  return s
}

func (s *CreateDeviceRequest) SetStreamProtocol(v string) *CreateDeviceRequest {
  s.StreamProtocol = &v
  return s
}

func (s *CreateDeviceRequest) SetType(v int) *CreateDeviceRequest {
  s.Type = &v
  return s
}

func (s *CreateDeviceRequest) SetDeviceId(v string) *CreateDeviceRequest {
  s.DeviceId = &v
  return s
}

func (s *CreateDeviceRequest) SetManufacturer(v string) *CreateDeviceRequest {
  s.Manufacturer = &v
  return s
}

func (s *CreateDeviceRequest) SetSpaceName(v string) *CreateDeviceRequest {
  s.SpaceName = &v
  return s
}

func (s *CreateDeviceRequest) SetPassword(v string) *CreateDeviceRequest {
  s.Password = &v
  return s
}

func (s *CreateDeviceRequest) SetStreamType(v string) *CreateDeviceRequest {
  s.StreamType = &v
  return s
}

func (s *CreateDeviceRequest) SetName(v string) *CreateDeviceRequest {
  s.Name = &v
  return s
}

func (s *CreateDeviceRequest) SetLongitude(v int64) *CreateDeviceRequest {
  s.Longitude = &v
  return s
}

type CreateDeviceRequestHeader struct {
}

func (s CreateDeviceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeviceRequestHeader) GoString() string {
  return s.String()
}

type CreateDevicePaths struct {
}

func (s CreateDevicePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateDevicePaths) GoString() string {
  return s.String()
}

type CreateDeviceParameters struct {
}

func (s CreateDeviceParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateDeviceParameters) GoString() string {
  return s.String()
}

type CreateDeviceResponse struct {
  // {"en":"Result status code, 0 indicates success","zh_CN":"结果状态码，0为成功"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateDeviceResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateDeviceResponse) GoString() string {
  return s.String()
}

func (s *CreateDeviceResponse) SetCode(v int) *CreateDeviceResponse {
  s.Code = &v
  return s
}

func (s *CreateDeviceResponse) SetMessage(v string) *CreateDeviceResponse {
  s.Message = &v
  return s
}

type CreateDeviceResponseHeader struct {
}

func (s CreateDeviceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateDeviceResponseHeader) GoString() string {
  return s.String()
}




type DeleteDeviceRequest struct {
  // {"en":"The national standard id of the device", "zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
}

func (s DeleteDeviceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
  return s.String()
}

func (s *DeleteDeviceRequest) SetDeviceId(v string) *DeleteDeviceRequest {
  s.DeviceId = &v
  return s
}

type DeleteDeviceResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteDeviceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
  return s.String()
}

func (s *DeleteDeviceResponse) SetCode(v int32) *DeleteDeviceResponse {
  s.Code = &v
  return s
}

func (s *DeleteDeviceResponse) SetMessage(v string) *DeleteDeviceResponse {
  s.Message = &v
  return s
}

type DeleteDevicePaths struct {
}

func (s DeleteDevicePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteDevicePaths) GoString() string {
  return s.String()
}

type DeleteDeviceParameters struct {
}

func (s DeleteDeviceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteDeviceParameters) GoString() string {
  return s.String()
}

type DeleteDeviceRequestHeader struct {
}

func (s DeleteDeviceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteDeviceRequestHeader) GoString() string {
  return s.String()
}

type DeleteDeviceResponseHeader struct {
}

func (s DeleteDeviceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteDeviceResponseHeader) GoString() string {
  return s.String()
}




type PtzControlRequest struct {
  // {"en":"Device GB28181 Id", "zh_CN":"设备国标ID"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Channel GB28181 Id", "zh_CN":"通道国标ID"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"PTZ control commands
  // ptzUp: The Device turns upward
  // ptzDown: The Device turns down
  // ptzLeft: The Device turns left
  // ptzRight: The Device turns right
  // ptzLeftUp: The Device turns to the upper left
  // ptzRightUp: The Device turns right up
  // ptzLeftDown: The Device turns left and down
  // ptzRightDown: The Device turns right and down
  // ptzZoomIn: Device lens zoom in
  // ptzZoomOut: The Device lens is zoomed out
  // ptzStop: Stop ptz operation
  // fiIrisIn: Device aperture magnification
  // fiIrisOut: The Device aperture is reduced
  // fiFocusIn: Device lens far focus
  // fiFocusOut: The Device lens is in close focus
  // fiStop: Stop", "zh_CN":"云台控制指令
  // 
  // ptzUp：设备向上转
  // 
  // ptzDown：设备向下转
  // 
  // ptzLeft：设备向左转
  // 
  // ptzRight：设备向右转
  // 
  // ptzLeftUp：设备向左上转
  // 
  // ptzRightUp：设备向右上转
  // 
  // ptzLeftDown：设备向左下转
  // 
  // ptzRightDown：设备向右下转
  // 
  // ptzZoomIn：设备镜头放大
  // 
  // ptzZoomOut：设备镜头缩小
  // 
  // ptzStop：停止ptz操作
  // 
  // fiIrisIn：设备光圈放大
  // 
  // fiIrisOut：设备光圈缩小
  // 
  // fiFocusIn：设备镜头远焦
  // 
  // fiFocusOut：设备镜头近焦
  // 
  // fiStop：停止"}
  Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty" require:"true"`
  // {"en":"Step Length
  // 1. Set the rotation speed when the control Device rotates up, down, left, and right. The value range is 0-255, and the default is 10
  // 2. Set the value when the control Device lens is zoomed in or out. The value range is 0-15, and the default value is 1
  // 3. When controlling the aperture or focal length of the Device, set the value, the range is 0-255, the default is 10", "zh_CN":"步长
  // 
  // 1、当控制设备向上下左右旋转时设置旋转的速度，取值范围0-255，默认10
  // 
  // 2、当控制设备镜头放大或缩小时设置数值，取值范围0-15，默认1
  // 
  // 3、当控制设备光圈或焦距时设置数值，取值范围0-255，默认10"}
  PtzParam *int32 `json:"ptzParam,omitempty" xml:"ptzParam,omitempty"`
}

func (s PtzControlRequest) String() string {
  return tea.Prettify(s)
}

func (s PtzControlRequest) GoString() string {
  return s.String()
}

func (s *PtzControlRequest) SetDeviceId(v string) *PtzControlRequest {
  s.DeviceId = &v
  return s
}

func (s *PtzControlRequest) SetChannelId(v string) *PtzControlRequest {
  s.ChannelId = &v
  return s
}

func (s *PtzControlRequest) SetCmd(v string) *PtzControlRequest {
  s.Cmd = &v
  return s
}

func (s *PtzControlRequest) SetPtzParam(v int32) *PtzControlRequest {
  s.PtzParam = &v
  return s
}

type PtzControlResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s PtzControlResponse) String() string {
  return tea.Prettify(s)
}

func (s PtzControlResponse) GoString() string {
  return s.String()
}

func (s *PtzControlResponse) SetCode(v int32) *PtzControlResponse {
  s.Code = &v
  return s
}

func (s *PtzControlResponse) SetMessage(v string) *PtzControlResponse {
  s.Message = &v
  return s
}

type PtzControlPaths struct {
}

func (s PtzControlPaths) String() string {
  return tea.Prettify(s)
}

func (s PtzControlPaths) GoString() string {
  return s.String()
}

type PtzControlParameters struct {
}

func (s PtzControlParameters) String() string {
  return tea.Prettify(s)
}

func (s PtzControlParameters) GoString() string {
  return s.String()
}

type PtzControlRequestHeader struct {
}

func (s PtzControlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PtzControlRequestHeader) GoString() string {
  return s.String()
}

type PtzControlResponseHeader struct {
}

func (s PtzControlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PtzControlResponseHeader) GoString() string {
  return s.String()
}




type EditDeviceRequest struct {
  // {"en":"Device registration password","zh_CN":"设备注册密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"Device Address","zh_CN":"设备地址"}
  Address *string `json:"address,omitempty" xml:"address,omitempty"`
  // {"en":"Stream Type:\ndefault\nstream:0\nstream:1\nstreamnumber:0\nstreamnumber:1\nstreamprofile:0\nstreamprofile:1\nstreamMode:MAIN\nstreamMode:SUB","zh_CN":"码流类型:\ndefault\nstream:0\nstream:1\nstreamnumber:0\nstreamnumber:1\nstreamprofile:0\nstreamprofile:1\nstreamMode:MAIN\nstreamMode:SUB"}
  StreamType *string `json:"streamType,omitempty" xml:"streamType,omitempty"`
  // {"en":"Whether to enable automatic invitation streaming. true: Enable automatic invitation streaming, false: Disable automatic invitation streaming","zh_CN":"是否启动自动邀请推流\n1.是，0.否\n默认：01:启动自动邀请推流0:不启动自动邀请推流"}
  IsAutoPush *int `json:"isAutoPush,omitempty" xml:"isAutoPush,omitempty"`
  // {"en":"Laitude","zh_CN":"纬度 仅支持正负数，整数部分最多3位，小数部分最多8位"}
  Latitude *int64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
  // {"en":"Device Name","zh_CN":"设备名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Device Description","zh_CN":"设备描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Streaming protocols, Do not pass the default TCP mode to stream pushing. If you need to use UDP to stream pushing, carry this field The values are as follows: TCP UDP","zh_CN":"流协议，  不传默认TCP方式推流，如果需要使用UDP推流则携带这个字段  取值如下：  TCP  UDP"}
  StreamProtocol *string `json:"streamProtocol,omitempty" xml:"streamProtocol,omitempty"`
  // {"en":"The national standard id of the device","zh_CN":"设备国标id"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Equipment Manufacturer","zh_CN":"设备厂商"}
  Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
  // {"en":"Longituder","zh_CN":"经度 仅支持正负数，整数部分最多3位，小数部分最多8位"}
  Longitude *int64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s EditDeviceRequest) String() string {
  return tea.Prettify(s)
}

func (s EditDeviceRequest) GoString() string {
  return s.String()
}

func (s *EditDeviceRequest) SetPassword(v string) *EditDeviceRequest {
  s.Password = &v
  return s
}

func (s *EditDeviceRequest) SetAddress(v string) *EditDeviceRequest {
  s.Address = &v
  return s
}

func (s *EditDeviceRequest) SetStreamType(v string) *EditDeviceRequest {
  s.StreamType = &v
  return s
}

func (s *EditDeviceRequest) SetIsAutoPush(v int) *EditDeviceRequest {
  s.IsAutoPush = &v
  return s
}

func (s *EditDeviceRequest) SetLatitude(v int64) *EditDeviceRequest {
  s.Latitude = &v
  return s
}

func (s *EditDeviceRequest) SetName(v string) *EditDeviceRequest {
  s.Name = &v
  return s
}

func (s *EditDeviceRequest) SetDescription(v string) *EditDeviceRequest {
  s.Description = &v
  return s
}

func (s *EditDeviceRequest) SetStreamProtocol(v string) *EditDeviceRequest {
  s.StreamProtocol = &v
  return s
}

func (s *EditDeviceRequest) SetDeviceId(v string) *EditDeviceRequest {
  s.DeviceId = &v
  return s
}

func (s *EditDeviceRequest) SetManufacturer(v string) *EditDeviceRequest {
  s.Manufacturer = &v
  return s
}

func (s *EditDeviceRequest) SetLongitude(v int64) *EditDeviceRequest {
  s.Longitude = &v
  return s
}

type EditDeviceRequestHeader struct {
}

func (s EditDeviceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDeviceRequestHeader) GoString() string {
  return s.String()
}

type EditDevicePaths struct {
}

func (s EditDevicePaths) String() string {
  return tea.Prettify(s)
}

func (s EditDevicePaths) GoString() string {
  return s.String()
}

type EditDeviceParameters struct {
}

func (s EditDeviceParameters) String() string {
  return tea.Prettify(s)
}

func (s EditDeviceParameters) GoString() string {
  return s.String()
}

type EditDeviceResponse struct {
  // {"en":"Result status code, 0 indicates success","zh_CN":"结果状态码，0为成功"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return data","zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Return message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditDeviceResponse) String() string {
  return tea.Prettify(s)
}

func (s EditDeviceResponse) GoString() string {
  return s.String()
}

func (s *EditDeviceResponse) SetCode(v int) *EditDeviceResponse {
  s.Code = &v
  return s
}

func (s *EditDeviceResponse) SetData(v string) *EditDeviceResponse {
  s.Data = &v
  return s
}

func (s *EditDeviceResponse) SetMessage(v string) *EditDeviceResponse {
  s.Message = &v
  return s
}

type EditDeviceResponseHeader struct {
}

func (s EditDeviceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDeviceResponseHeader) GoString() string {
  return s.String()
}




type EditChannelInfoRequest struct {
  // {"en":"Device GB28181 Id.", "zh_CN":"设备国标id。不支持模糊搜索。"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Channel GB28181 Id", "zh_CN":"通道国标ID"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"Channel name", "zh_CN":"通道名称"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
  // {"en":"Cloud Recording enabled,0: Disabled (default) 1: Enabled", "zh_CN":"云端录制开关状态，0：未开启（默认）  1：开启"}
  RecordSwitch *int32 `json:"recordSwitch,omitempty" xml:"recordSwitch,omitempty"`
  // {"en":"Immediate record enabled,0: Disabled (default),1: Enabled,If the Device is Push streaming, it is forbidden to turn on or off the push-to-record function. When the push-to-record function is turned on, the Recording template Filed recordTemplate is invalid.", "zh_CN":"即推即录开关状态，0：未开启（默认）  1：开启，如果设备在推流中，禁止开启关闭即推即录。当开启即推即录是录制模板字段“recordTemplate”无效。"}
  ImmediateRecord *int32 `json:"immediateRecord,omitempty" xml:"immediateRecord,omitempty"`
  // {"en":"The Recording template associated with the channel.", "zh_CN":"通道关联的录制模板名称。当前录制模板名称是唯一"}
  RecordTemplate *string `json:"recordTemplate,omitempty" xml:"recordTemplate,omitempty"`
  // {"en":"Immediate record duration of file retention.Unit: Day
  // 
  // Default: 30 days", "zh_CN":"即推即录文件保存时长。单位：天
  // 
  // 默认30天"}
  SavePeriod *int32 `json:"savePeriod,omitempty" xml:"savePeriod,omitempty"`
}

func (s EditChannelInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s EditChannelInfoRequest) GoString() string {
  return s.String()
}

func (s *EditChannelInfoRequest) SetDeviceId(v string) *EditChannelInfoRequest {
  s.DeviceId = &v
  return s
}

func (s *EditChannelInfoRequest) SetChannelId(v string) *EditChannelInfoRequest {
  s.ChannelId = &v
  return s
}

func (s *EditChannelInfoRequest) SetChannelName(v string) *EditChannelInfoRequest {
  s.ChannelName = &v
  return s
}

func (s *EditChannelInfoRequest) SetRecordSwitch(v int32) *EditChannelInfoRequest {
  s.RecordSwitch = &v
  return s
}

func (s *EditChannelInfoRequest) SetImmediateRecord(v int32) *EditChannelInfoRequest {
  s.ImmediateRecord = &v
  return s
}

func (s *EditChannelInfoRequest) SetRecordTemplate(v string) *EditChannelInfoRequest {
  s.RecordTemplate = &v
  return s
}

func (s *EditChannelInfoRequest) SetSavePeriod(v int32) *EditChannelInfoRequest {
  s.SavePeriod = &v
  return s
}

type EditChannelInfoResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditChannelInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s EditChannelInfoResponse) GoString() string {
  return s.String()
}

func (s *EditChannelInfoResponse) SetCode(v int32) *EditChannelInfoResponse {
  s.Code = &v
  return s
}

func (s *EditChannelInfoResponse) SetMessage(v string) *EditChannelInfoResponse {
  s.Message = &v
  return s
}

type EditChannelInfoPaths struct {
}

func (s EditChannelInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s EditChannelInfoPaths) GoString() string {
  return s.String()
}

type EditChannelInfoParameters struct {
}

func (s EditChannelInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s EditChannelInfoParameters) GoString() string {
  return s.String()
}

type EditChannelInfoRequestHeader struct {
}

func (s EditChannelInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditChannelInfoRequestHeader) GoString() string {
  return s.String()
}

type EditChannelInfoResponseHeader struct {
}

func (s EditChannelInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditChannelInfoResponseHeader) GoString() string {
  return s.String()
}




type ChannelStatusUpdateRequest struct {
  // {"en":"Channel GB28181 Id", "zh_CN":"通道国标ID"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
}

func (s ChannelStatusUpdateRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelStatusUpdateRequest) GoString() string {
  return s.String()
}

func (s *ChannelStatusUpdateRequest) SetDeviceId(v string) *ChannelStatusUpdateRequest {
  s.DeviceId = &v
  return s
}

type ChannelStatusUpdateResponse struct {
  // {"en":"Results status code, 0 means success", "zh_CN":"结果状态码，0 为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operate successfully", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ChannelStatusUpdateResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelStatusUpdateResponse) GoString() string {
  return s.String()
}

func (s *ChannelStatusUpdateResponse) SetCode(v int32) *ChannelStatusUpdateResponse {
  s.Code = &v
  return s
}

func (s *ChannelStatusUpdateResponse) SetMessage(v string) *ChannelStatusUpdateResponse {
  s.Message = &v
  return s
}

type ChannelStatusUpdatePaths struct {
}

func (s ChannelStatusUpdatePaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelStatusUpdatePaths) GoString() string {
  return s.String()
}

type ChannelStatusUpdateParameters struct {
}

func (s ChannelStatusUpdateParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelStatusUpdateParameters) GoString() string {
  return s.String()
}

type ChannelStatusUpdateRequestHeader struct {
}

func (s ChannelStatusUpdateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelStatusUpdateRequestHeader) GoString() string {
  return s.String()
}

type ChannelStatusUpdateResponseHeader struct {
}

func (s ChannelStatusUpdateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelStatusUpdateResponseHeader) GoString() string {
  return s.String()
}




type QueryChannelListRequest struct {
  // {"en":"The national standard id of the device. Fuzzy search is not supported.", "zh_CN":"设备国标id。不支持模糊搜索。"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"Specify the paging size. The default value is 10, and the maximum value is 50.", "zh_CN":"分页大小。默认10，最大50"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Page number, default is first page", "zh_CN":"第几页，默认第一页"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
}

func (s QueryChannelListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListRequest) GoString() string {
  return s.String()
}

func (s *QueryChannelListRequest) SetDeviceId(v string) *QueryChannelListRequest {
  s.DeviceId = &v
  return s
}

func (s *QueryChannelListRequest) SetPageSize(v int32) *QueryChannelListRequest {
  s.PageSize = &v
  return s
}

func (s *QueryChannelListRequest) SetPageIndex(v int32) *QueryChannelListRequest {
  s.PageIndex = &v
  return s
}

type QueryChannelListResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  QueryChannelListData *QueryChannelListData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryChannelListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListResponse) GoString() string {
  return s.String()
}

func (s *QueryChannelListResponse) SetCode(v int32) *QueryChannelListResponse {
  s.Code = &v
  return s
}

func (s *QueryChannelListResponse) SetMessage(v string) *QueryChannelListResponse {
  s.Message = &v
  return s
}

func (s *QueryChannelListResponse) SetData(v *QueryChannelListData) *QueryChannelListResponse {
  s.QueryChannelListData = v
  return s
}

type QueryChannelListData struct {
  // {"en":"QueryChannelListChannel List", "zh_CN":"通道列表"}
  Rows []*QueryChannelListChannel `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
  // {"en":"Page number,", "zh_CN":"第几页"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"Specify the paging size", "zh_CN":"分页大小"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total number of devices", "zh_CN":"总设备数"}
  Total *int32 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryChannelListData) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListData) GoString() string {
  return s.String()
}

func (s *QueryChannelListData) SetRows(v []*QueryChannelListChannel) *QueryChannelListData {
  s.Rows = v
  return s
}

func (s *QueryChannelListData) SetPageIndex(v int32) *QueryChannelListData {
  s.PageIndex = &v
  return s
}

func (s *QueryChannelListData) SetPageSize(v int32) *QueryChannelListData {
  s.PageSize = &v
  return s
}

func (s *QueryChannelListData) SetTotal(v int32) *QueryChannelListData {
  s.Total = &v
  return s
}

type QueryChannelListChannel struct {
  // {"en":"QueryChannelListChannel name", "zh_CN":"通道名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"QueryChannelListChannel national standard ID", "zh_CN":"通道国标ID"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"QueryChannelListChannel status: 0. Offline, 1. Online", "zh_CN":"通道状态：0.离线，1.在线"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Streaming status: 0: no streaming, 1: streaming", "zh_CN":"流状态：0：未推流，1：有推流"}
  HasStream *int32 `json:"hasStream,omitempty" xml:"hasStream,omitempty" require:"true"`
  // {"en":"QueryChannelListChannel type: 1. Video channel, 2. Alarm channel", "zh_CN":"通道类型：1.视频通道，2.告警通道"}
  Type *int32 `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Cloud Recording enabled,0: Disabled (default) 1: Enabled", "zh_CN":"云端录制开关状态，0：未开启（默认）  1：开启"}
  RecordSwitch *int32 `json:"recordSwitch,omitempty" xml:"recordSwitch,omitempty" require:"true"`
  // {"en":"Recording Status,NOT_RECORD: Not Recording(default) RECCORDING: Recording", "zh_CN":"录制状态，NOT_RECORD：未录制（默认） RECCORDING：录制中"}
  RecordState *string `json:"recordState,omitempty" xml:"recordState,omitempty" require:"true"`
  // {"en":"Immediate record enabled,0: Disabled (default) 1: Enabled", "zh_CN":"即推即录开关状态，0：未开启（默认）  1：开启"}
  ImmediateRecord *int32 `json:"immediateRecord,omitempty" xml:"immediateRecord,omitempty" require:"true"`
  // {"en":"Immediate record duration of file retention.Unit: Day", "zh_CN":"即推即录文件保存时长。单位：天"}
  SavePeriod *int32 `json:"savePeriod,omitempty" xml:"savePeriod,omitempty" require:"true"`
  // {"en":"QueryChannelListChannel resolution,If the Platform fails to obtain the resolution, it is unified to -,Example: 720*1280.", "zh_CN":"通道分辨率，如果平台未能获取到分辨率，统一为“-”，示列：720*1280。"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {"en":"Template for recording associated with a specific channel.", "zh_CN":"通道关联的录制模板名称"}
  RecordTemplate *string `json:"recordTemplate,omitempty" xml:"recordTemplate,omitempty" require:"true"`
  // {"en":"QueryChannelListChannel enable status. 1: Enable. 0: Disable", "zh_CN":"通道启用状态。1：启动。0：禁用"}
  IsEnabled *int32 `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s QueryChannelListChannel) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListChannel) GoString() string {
  return s.String()
}

func (s *QueryChannelListChannel) SetName(v string) *QueryChannelListChannel {
  s.Name = &v
  return s
}

func (s *QueryChannelListChannel) SetChannelId(v string) *QueryChannelListChannel {
  s.ChannelId = &v
  return s
}

func (s *QueryChannelListChannel) SetStatus(v int32) *QueryChannelListChannel {
  s.Status = &v
  return s
}

func (s *QueryChannelListChannel) SetHasStream(v int32) *QueryChannelListChannel {
  s.HasStream = &v
  return s
}

func (s *QueryChannelListChannel) SetType(v int32) *QueryChannelListChannel {
  s.Type = &v
  return s
}

func (s *QueryChannelListChannel) SetRecordSwitch(v int32) *QueryChannelListChannel {
  s.RecordSwitch = &v
  return s
}

func (s *QueryChannelListChannel) SetRecordState(v string) *QueryChannelListChannel {
  s.RecordState = &v
  return s
}

func (s *QueryChannelListChannel) SetImmediateRecord(v int32) *QueryChannelListChannel {
  s.ImmediateRecord = &v
  return s
}

func (s *QueryChannelListChannel) SetSavePeriod(v int32) *QueryChannelListChannel {
  s.SavePeriod = &v
  return s
}

func (s *QueryChannelListChannel) SetResolution(v string) *QueryChannelListChannel {
  s.Resolution = &v
  return s
}

func (s *QueryChannelListChannel) SetRecordTemplate(v string) *QueryChannelListChannel {
  s.RecordTemplate = &v
  return s
}

func (s *QueryChannelListChannel) SetIsEnabled(v int32) *QueryChannelListChannel {
  s.IsEnabled = &v
  return s
}

type QueryChannelListPaths struct {
}

func (s QueryChannelListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListPaths) GoString() string {
  return s.String()
}

type QueryChannelListParameters struct {
}

func (s QueryChannelListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListParameters) GoString() string {
  return s.String()
}

type QueryChannelListRequestHeader struct {
}

func (s QueryChannelListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListRequestHeader) GoString() string {
  return s.String()
}

type QueryChannelListResponseHeader struct {
}

func (s QueryChannelListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelListResponseHeader) GoString() string {
  return s.String()
}




