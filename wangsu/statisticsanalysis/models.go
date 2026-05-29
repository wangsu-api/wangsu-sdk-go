package statisticsanalysis

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCloudVODStorageVolumeRequest struct {
  // {"en":"Start Time:\n1. The format is yyyy-MM-dd, for example, 2024-01-23 (which is January 23, 2024, Beijing Time);\n2. Cannot be greater than the current time;","zh_CN":"开始时间：\n1.格式为yyyy-MM-dd，例如，2024-01-23（为北京时间2024年01月23日）；\n2.不能大于当前时间；"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:\n1. The format is yyyy-MM-dd;\n2. The end time must be greater than the start time;\n3. If the end time is greater than the current time, use the current time;\n4. If both dateFrom and dateTo are not passed, today data will be queried by default; if only one is not passed, an exception will be thrown;\n5. The maximum query interval allowed is 61 days, that is, the difference between dateFrom and dateTo cannot exceed 61 days.","zh_CN":"结束时间：\n1.格式为yyyy-MM-dd；\n2.结束时间需大于开始时间；\n3.结束时间如果大于当前时间，取当前时间；\n4.dateFrom，dateTo二者都未传，默认查询今天数据；如仅有一个未传，抛异常；\n5.允许查询最大间隔：61天，即dateFrom和dateTo相差不能超过61天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Whether to return details. Default: 0\n0: Do not return the peak storage peak value details at the daily granularity\n1: Returns the peak storage peak value details at the daily granularity","zh_CN":"是否返回明细。默认：0\n0:不返回天粒度存储量峰值明细\n1:返回天粒度存储量峰值明细"}
  IsDetails *int `json:"isDetails,omitempty" xml:"isDetails,omitempty"`
}

func (s QueryCloudVODStorageVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeRequest) GoString() string {
  return s.String()
}

func (s *QueryCloudVODStorageVolumeRequest) SetDateFrom(v string) *QueryCloudVODStorageVolumeRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryCloudVODStorageVolumeRequest) SetDateTo(v string) *QueryCloudVODStorageVolumeRequest {
  s.DateTo = &v
  return s
}

func (s *QueryCloudVODStorageVolumeRequest) SetIsDetails(v int) *QueryCloudVODStorageVolumeRequest {
  s.IsDetails = &v
  return s
}

type QueryCloudVODStorageVolumeRequestHeader struct {
}

func (s QueryCloudVODStorageVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeRequestHeader) GoString() string {
  return s.String()
}

type QueryCloudVODStorageVolumePaths struct {
}

func (s QueryCloudVODStorageVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumePaths) GoString() string {
  return s.String()
}

type QueryCloudVODStorageVolumeParameters struct {
}

func (s QueryCloudVODStorageVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeParameters) GoString() string {
  return s.String()
}

type QueryCloudVODStorageVolumeResponse struct {
  // {"en":"Result status code, 200 indicates success","zh_CN":"结果状态码，200为成功"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data","zh_CN":"返回数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Peak value of storage, unit MB","zh_CN":"存储量峰值，单位MB"}
  PeakValue *QueryCloudVODStorageVolumeResponsePeakValue `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true" type:"Struct"`
  // {"en":"Storage Volume Detail List","zh_CN":"存储量明细列表"}
  Details []*string `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
  // {"en":"Date","zh_CN":"日期"}
  Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
  // {"en":"Storage volume, unit: MB","zh_CN":"每日存储量峰值，单位MB"}
  StorageVolume *QueryCloudVODStorageVolumeResponseStorageVolume `json:"storageVolume,omitempty" xml:"storageVolume,omitempty" require:"true" type:"Struct"`
}

func (s QueryCloudVODStorageVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeResponse) GoString() string {
  return s.String()
}

func (s *QueryCloudVODStorageVolumeResponse) SetCode(v int) *QueryCloudVODStorageVolumeResponse {
  s.Code = &v
  return s
}

func (s *QueryCloudVODStorageVolumeResponse) SetMessage(v string) *QueryCloudVODStorageVolumeResponse {
  s.Message = &v
  return s
}

func (s *QueryCloudVODStorageVolumeResponse) SetData(v []*string) *QueryCloudVODStorageVolumeResponse {
  s.Data = v
  return s
}

func (s *QueryCloudVODStorageVolumeResponse) SetPeakValue(v *QueryCloudVODStorageVolumeResponsePeakValue) *QueryCloudVODStorageVolumeResponse {
  s.PeakValue = v
  return s
}

func (s *QueryCloudVODStorageVolumeResponse) SetDetails(v []*string) *QueryCloudVODStorageVolumeResponse {
  s.Details = v
  return s
}

func (s *QueryCloudVODStorageVolumeResponse) SetDate(v string) *QueryCloudVODStorageVolumeResponse {
  s.Date = &v
  return s
}

func (s *QueryCloudVODStorageVolumeResponse) SetStorageVolume(v *QueryCloudVODStorageVolumeResponseStorageVolume) *QueryCloudVODStorageVolumeResponse {
  s.StorageVolume = v
  return s
}

type QueryCloudVODStorageVolumeResponsePeakValue struct {
}

func (s QueryCloudVODStorageVolumeResponsePeakValue) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeResponsePeakValue) GoString() string {
  return s.String()
}

type QueryCloudVODStorageVolumeResponseStorageVolume struct {
}

func (s QueryCloudVODStorageVolumeResponseStorageVolume) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeResponseStorageVolume) GoString() string {
  return s.String()
}

type QueryCloudVODStorageVolumeResponseHeader struct {
}

func (s QueryCloudVODStorageVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudVODStorageVolumeResponseHeader) GoString() string {
  return s.String()
}




type GetRealTimeChannelOnlineNumberRequest struct {
  // {"en":"Pull id, multiple values separated by \",\"", "zh_CN":"拉流 id，多个值通过\",\"隔开"}
  PullIds *string `json:"pullIds,omitempty" xml:"pullIds,omitempty"`
}

func (s GetRealTimeChannelOnlineNumberRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberRequest) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberRequest) SetPullIds(v string) *GetRealTimeChannelOnlineNumberRequest {
  s.PullIds = &v
  return s
}

type GetRealTimeChannelOnlineNumberResponse struct {
  // {"en":"200 success", "zh_CN":"200，操作成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetRealTimeChannelOnlineNumberData []*GetRealTimeChannelOnlineNumberData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetRealTimeChannelOnlineNumberResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberResponse) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberResponse) SetCode(v int32) *GetRealTimeChannelOnlineNumberResponse {
  s.Code = &v
  return s
}

func (s *GetRealTimeChannelOnlineNumberResponse) SetMessage(v string) *GetRealTimeChannelOnlineNumberResponse {
  s.Message = &v
  return s
}

func (s *GetRealTimeChannelOnlineNumberResponse) SetData(v []*GetRealTimeChannelOnlineNumberData) *GetRealTimeChannelOnlineNumberResponse {
  s.GetRealTimeChannelOnlineNumberData = v
  return s
}

type GetRealTimeChannelOnlineNumberData struct {
  // {"en":"Online list of people", "zh_CN":"在线人数列表"}
  OnlineNumberList []*GetRealTimeChannelOnlineNumberOnlineNumberItem `json:"onlineNumberList,omitempty" xml:"onlineNumberList,omitempty" require:"true" type:"Repeated"`
}

func (s GetRealTimeChannelOnlineNumberData) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberData) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberData) SetOnlineNumberList(v []*GetRealTimeChannelOnlineNumberOnlineNumberItem) *GetRealTimeChannelOnlineNumberData {
  s.OnlineNumberList = v
  return s
}

type GetRealTimeChannelOnlineNumberOnlineNumberItem struct {
  // {"en":"Online population", "zh_CN":"在线人数"}
  OnlineNumber *int32 `json:"onlineNumber,omitempty" xml:"onlineNumber,omitempty" require:"true"`
  // {"en":"pullId", "zh_CN":"拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s GetRealTimeChannelOnlineNumberOnlineNumberItem) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberOnlineNumberItem) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberOnlineNumberItem) SetOnlineNumber(v int32) *GetRealTimeChannelOnlineNumberOnlineNumberItem {
  s.OnlineNumber = &v
  return s
}

func (s *GetRealTimeChannelOnlineNumberOnlineNumberItem) SetPullId(v string) *GetRealTimeChannelOnlineNumberOnlineNumberItem {
  s.PullId = &v
  return s
}

type GetRealTimeChannelOnlineNumberPaths struct {
}

func (s GetRealTimeChannelOnlineNumberPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberPaths) GoString() string {
  return s.String()
}

type GetRealTimeChannelOnlineNumberParameters struct {
}

func (s GetRealTimeChannelOnlineNumberParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberParameters) GoString() string {
  return s.String()
}

type GetRealTimeChannelOnlineNumberRequestHeader struct {
}

func (s GetRealTimeChannelOnlineNumberRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberRequestHeader) GoString() string {
  return s.String()
}

type GetRealTimeChannelOnlineNumberResponseHeader struct {
}

func (s GetRealTimeChannelOnlineNumberResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberResponseHeader) GoString() string {
  return s.String()
}




