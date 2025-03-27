package taskmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QuerysubtasklistRequest struct {
}

func (s QuerysubtasklistRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerysubtasklistRequest) GoString() string {
  return s.String()
}

type QuerysubtasklistResponse struct {
  // {"en":"list of subtasks", "zh_CN":"云手机子任务单列表"}
  Subtasks []*string `json:"subtasks,omitempty" xml:"subtasks,omitempty" require:"true" type:"Repeated"`
}

func (s QuerysubtasklistResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerysubtasklistResponse) GoString() string {
  return s.String()
}

func (s *QuerysubtasklistResponse) SetSubtasks(v []*string) *QuerysubtasklistResponse {
  s.Subtasks = v
  return s
}

type QuerysubtasklistPaths struct {
}

func (s QuerysubtasklistPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerysubtasklistPaths) GoString() string {
  return s.String()
}

type QuerysubtasklistParameters struct {
  // {"en":"subtask ID to be queried", "zh_CN":"要查询的任务单id，多个id使用英文逗号进行分隔"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
}

func (s QuerysubtasklistParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerysubtasklistParameters) GoString() string {
  return s.String()
}

func (s *QuerysubtasklistParameters) SetIds(v string) *QuerysubtasklistParameters {
  s.Ids = &v
  return s
}

type QuerysubtasklistRequestHeader struct {
}

func (s QuerysubtasklistRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerysubtasklistRequestHeader) GoString() string {
  return s.String()
}

type QuerysubtasklistResponseHeader struct {
}

func (s QuerysubtasklistResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerysubtasklistResponseHeader) GoString() string {
  return s.String()
}




type QueryTaskListRequest struct {
}

func (s QueryTaskListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListRequest) GoString() string {
  return s.String()
}

type QueryTaskListResponse struct {
  // {"en":"task list", "zh_CN":"任务单列表"}
  Tasks []*QueryTaskListTask `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
  // {"en":"status", "zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"消息"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s QueryTaskListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListResponse) GoString() string {
  return s.String()
}

func (s *QueryTaskListResponse) SetTasks(v []*QueryTaskListTask) *QueryTaskListResponse {
  s.Tasks = v
  return s
}

func (s *QueryTaskListResponse) SetStatus(v int) *QueryTaskListResponse {
  s.Status = &v
  return s
}

func (s *QueryTaskListResponse) SetResult(v string) *QueryTaskListResponse {
  s.Result = &v
  return s
}

type QueryTaskListTask struct {
  // {"en":"task id", "zh_CN":"任务单id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"任务单消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QueryTaskListTask) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListTask) GoString() string {
  return s.String()
}

func (s *QueryTaskListTask) SetId(v string) *QueryTaskListTask {
  s.Id = &v
  return s
}

func (s *QueryTaskListTask) SetMessage(v string) *QueryTaskListTask {
  s.Message = &v
  return s
}

type QueryTaskListPaths struct {
}

func (s QueryTaskListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListPaths) GoString() string {
  return s.String()
}

type QueryTaskListParameters struct {
  // {"en":"task ID to be queried", "zh_CN":"要查询的任务单id，多个id使用英文逗号进行分隔"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
}

func (s QueryTaskListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListParameters) GoString() string {
  return s.String()
}

func (s *QueryTaskListParameters) SetIds(v string) *QueryTaskListParameters {
  s.Ids = &v
  return s
}

type QueryTaskListRequestHeader struct {
}

func (s QueryTaskListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListRequestHeader) GoString() string {
  return s.String()
}

type QueryTaskListResponseHeader struct {
}

func (s QueryTaskListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTaskListResponseHeader) GoString() string {
  return s.String()
}




