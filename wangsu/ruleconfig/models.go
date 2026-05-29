package ruleconfig

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPropertyJsonSchemaRequest struct {
}

func (s GetPropertyJsonSchemaRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaRequest) GoString() string {
  return s.String()
}

type GetPropertyJsonSchemaRequestHeader struct {
}

func (s GetPropertyJsonSchemaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaRequestHeader) GoString() string {
  return s.String()
}

type GetPropertyJsonSchemaPaths struct {
}

func (s GetPropertyJsonSchemaPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaPaths) GoString() string {
  return s.String()
}

type GetPropertyJsonSchemaParameters struct {
}

func (s GetPropertyJsonSchemaParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaParameters) GoString() string {
  return s.String()
}

type GetPropertyJsonSchemaResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetPropertyJsonSchemaResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetPropertyJsonSchemaResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaResponse) GoString() string {
  return s.String()
}

func (s *GetPropertyJsonSchemaResponse) SetCode(v string) *GetPropertyJsonSchemaResponse {
  s.Code = &v
  return s
}

func (s *GetPropertyJsonSchemaResponse) SetData(v *GetPropertyJsonSchemaResponseData) *GetPropertyJsonSchemaResponse {
  s.Data = v
  return s
}

func (s *GetPropertyJsonSchemaResponse) SetMessage(v string) *GetPropertyJsonSchemaResponse {
  s.Message = &v
  return s
}

type GetPropertyJsonSchemaResponseData struct {
  // {"en":"property configuration json schema.","zh_CN":"property配置的json schema"}
  Schema *string `json:"schema,omitempty" xml:"schema,omitempty" require:"true"`
}

func (s GetPropertyJsonSchemaResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaResponseData) GoString() string {
  return s.String()
}

func (s *GetPropertyJsonSchemaResponseData) SetSchema(v string) *GetPropertyJsonSchemaResponseData {
  s.Schema = &v
  return s
}

type GetPropertyJsonSchemaResponseHeader struct {
}

func (s GetPropertyJsonSchemaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyJsonSchemaResponseHeader) GoString() string {
  return s.String()
}




type DeleteRuleRequest struct {
}

func (s DeleteRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
  return s.String()
}

type DeleteRuleRequestHeader struct {
}

func (s DeleteRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRuleRequestHeader) GoString() string {
  return s.String()
}

type DeleteRulePaths struct {
  // {"en":"rule ID","zh_CN":"规则ID"}
  RuleID *int `json:"ruleID,omitempty" xml:"ruleID,omitempty" require:"true"`
}

func (s DeleteRulePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteRulePaths) GoString() string {
  return s.String()
}

func (s *DeleteRulePaths) SetRuleID(v int) *DeleteRulePaths {
  s.RuleID = &v
  return s
}

type DeleteRuleParameters struct {
}

func (s DeleteRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteRuleParameters) GoString() string {
  return s.String()
}

type DeleteRuleResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
  return s.String()
}

func (s *DeleteRuleResponse) SetCode(v string) *DeleteRuleResponse {
  s.Code = &v
  return s
}

func (s *DeleteRuleResponse) SetMessage(v string) *DeleteRuleResponse {
  s.Message = &v
  return s
}

type DeleteRuleResponseHeader struct {
}

func (s DeleteRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateRuleStatusRequest struct {
  // {"en":"The status of the rule.","zh_CN":"规则的状态","exampleValue":"enable,disable"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s UpdateRuleStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusRequest) GoString() string {
  return s.String()
}

func (s *UpdateRuleStatusRequest) SetStatus(v string) *UpdateRuleStatusRequest {
  s.Status = &v
  return s
}

type UpdateRuleStatusRequestHeader struct {
}

func (s UpdateRuleStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusRequestHeader) GoString() string {
  return s.String()
}

type UpdateRuleStatusPaths struct {
  // {"en":"rule ID","zh_CN":"规则ID"}
  RuleId *int `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s UpdateRuleStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusPaths) GoString() string {
  return s.String()
}

func (s *UpdateRuleStatusPaths) SetRuleId(v int) *UpdateRuleStatusPaths {
  s.RuleId = &v
  return s
}

type UpdateRuleStatusParameters struct {
}

func (s UpdateRuleStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusParameters) GoString() string {
  return s.String()
}

type UpdateRuleStatusResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *UpdateRuleStatusResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateRuleStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusResponse) GoString() string {
  return s.String()
}

func (s *UpdateRuleStatusResponse) SetCode(v string) *UpdateRuleStatusResponse {
  s.Code = &v
  return s
}

func (s *UpdateRuleStatusResponse) SetData(v *UpdateRuleStatusResponseData) *UpdateRuleStatusResponse {
  s.Data = v
  return s
}

func (s *UpdateRuleStatusResponse) SetMessage(v string) *UpdateRuleStatusResponse {
  s.Message = &v
  return s
}

type UpdateRuleStatusResponseData struct {
  // {"en":"Rule ID","zh_CN":"规则标识"}
  RuleId *int64 `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s UpdateRuleStatusResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusResponseData) GoString() string {
  return s.String()
}

func (s *UpdateRuleStatusResponseData) SetRuleId(v int64) *UpdateRuleStatusResponseData {
  s.RuleId = &v
  return s
}

type UpdateRuleStatusResponseHeader struct {
}

func (s UpdateRuleStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleStatusResponseHeader) GoString() string {
  return s.String()
}




type GetRuleRequest struct {
}

func (s GetRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
  return s.String()
}

type GetRuleRequestHeader struct {
}

func (s GetRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRuleRequestHeader) GoString() string {
  return s.String()
}

type GetRulePaths struct {
  // {"en":"rule ID","zh_CN":"规则ID"}
  RuleId *int `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s GetRulePaths) String() string {
  return tea.Prettify(s)
}

func (s GetRulePaths) GoString() string {
  return s.String()
}

func (s *GetRulePaths) SetRuleId(v int) *GetRulePaths {
  s.RuleId = &v
  return s
}

type GetRuleParameters struct {
}

func (s GetRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRuleParameters) GoString() string {
  return s.String()
}

type GetRuleResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
  return s.String()
}

func (s *GetRuleResponse) SetCode(v string) *GetRuleResponse {
  s.Code = &v
  return s
}

func (s *GetRuleResponse) SetData(v *GetRuleResponseData) *GetRuleResponse {
  s.Data = v
  return s
}

func (s *GetRuleResponse) SetMessage(v string) *GetRuleResponse {
  s.Message = &v
  return s
}

type GetRuleResponseData struct {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Phase.","zh_CN":"阶段","exampleValue":"connectPhase,requestPhase,cachePhase,originPhase,responsePhase"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty" require:"true"`
  // {"en":"The condition in the rule","zh_CN":"规则中的条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Rule action","zh_CN":"规则动作"}
  Action *GetRuleResponseDataAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"Comment","zh_CN":"规则的备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"The ID of the rule","zh_CN":"规则的标识"}
  RuleId *int64 `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"The priority of rule.","zh_CN":"规则排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The state of rule.","zh_CN":"规则状态。false为禁用，true为启用。","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRuleResponseData) GoString() string {
  return s.String()
}

func (s *GetRuleResponseData) SetName(v string) *GetRuleResponseData {
  s.Name = &v
  return s
}

func (s *GetRuleResponseData) SetPhase(v string) *GetRuleResponseData {
  s.Phase = &v
  return s
}

func (s *GetRuleResponseData) SetCondition(v string) *GetRuleResponseData {
  s.Condition = &v
  return s
}

func (s *GetRuleResponseData) SetAction(v *GetRuleResponseDataAction) *GetRuleResponseData {
  s.Action = v
  return s
}

func (s *GetRuleResponseData) SetDescription(v string) *GetRuleResponseData {
  s.Description = &v
  return s
}

func (s *GetRuleResponseData) SetRuleId(v int64) *GetRuleResponseData {
  s.RuleId = &v
  return s
}

func (s *GetRuleResponseData) SetPriority(v int) *GetRuleResponseData {
  s.Priority = &v
  return s
}

func (s *GetRuleResponseData) SetEnabled(v bool) *GetRuleResponseData {
  s.Enabled = &v
  return s
}

type GetRuleResponseDataAction struct {
  // {"en":"Action name","zh_CN":"动作名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Action options.","zh_CN":"动作参数。"}
  Options *string `json:"options,omitempty" xml:"options,omitempty" require:"true"`
}

func (s GetRuleResponseDataAction) String() string {
  return tea.Prettify(s)
}

func (s GetRuleResponseDataAction) GoString() string {
  return s.String()
}

func (s *GetRuleResponseDataAction) SetName(v string) *GetRuleResponseDataAction {
  s.Name = &v
  return s
}

func (s *GetRuleResponseDataAction) SetOptions(v string) *GetRuleResponseDataAction {
  s.Options = &v
  return s
}

type GetRuleResponseHeader struct {
}

func (s GetRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateRulePriorityRequest struct {
  // {"en":"The priority of the rule","zh_CN":"规则的顺序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
}

func (s UpdateRulePriorityRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityRequest) GoString() string {
  return s.String()
}

func (s *UpdateRulePriorityRequest) SetPriority(v int) *UpdateRulePriorityRequest {
  s.Priority = &v
  return s
}

type UpdateRulePriorityRequestHeader struct {
}

func (s UpdateRulePriorityRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityRequestHeader) GoString() string {
  return s.String()
}

type UpdateRulePriorityPaths struct {
  // {"en":"rule ID","zh_CN":"规则ID"}
  RuleId *int `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s UpdateRulePriorityPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityPaths) GoString() string {
  return s.String()
}

func (s *UpdateRulePriorityPaths) SetRuleId(v int) *UpdateRulePriorityPaths {
  s.RuleId = &v
  return s
}

type UpdateRulePriorityParameters struct {
}

func (s UpdateRulePriorityParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityParameters) GoString() string {
  return s.String()
}

type UpdateRulePriorityResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *UpdateRulePriorityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateRulePriorityResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityResponse) GoString() string {
  return s.String()
}

func (s *UpdateRulePriorityResponse) SetCode(v string) *UpdateRulePriorityResponse {
  s.Code = &v
  return s
}

func (s *UpdateRulePriorityResponse) SetData(v *UpdateRulePriorityResponseData) *UpdateRulePriorityResponse {
  s.Data = v
  return s
}

func (s *UpdateRulePriorityResponse) SetMessage(v string) *UpdateRulePriorityResponse {
  s.Message = &v
  return s
}

type UpdateRulePriorityResponseData struct {
  // {"en":"Rule ID","zh_CN":"规则标识"}
  RuleId *int64 `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s UpdateRulePriorityResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityResponseData) GoString() string {
  return s.String()
}

func (s *UpdateRulePriorityResponseData) SetRuleId(v int64) *UpdateRulePriorityResponseData {
  s.RuleId = &v
  return s
}

type UpdateRulePriorityResponseHeader struct {
}

func (s UpdateRulePriorityResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePriorityResponseHeader) GoString() string {
  return s.String()
}




type UpdateRuleRequest struct {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Phase","zh_CN":"阶段","exampleValue":"connectPhase,requestPhase,cachePhase,originPhase,responsePhase"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty" require:"true"`
  // {"en":"The condition in the rule","zh_CN":"规则中的条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"Rule action","zh_CN":"规则动作"}
  Action *UpdateRuleRequestAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"Comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"defaultValue":"1","en":"The priority of rule.","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"The status of rule.","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
  return s.String()
}

func (s *UpdateRuleRequest) SetName(v string) *UpdateRuleRequest {
  s.Name = &v
  return s
}

func (s *UpdateRuleRequest) SetPhase(v string) *UpdateRuleRequest {
  s.Phase = &v
  return s
}

func (s *UpdateRuleRequest) SetCondition(v string) *UpdateRuleRequest {
  s.Condition = &v
  return s
}

func (s *UpdateRuleRequest) SetAction(v *UpdateRuleRequestAction) *UpdateRuleRequest {
  s.Action = v
  return s
}

func (s *UpdateRuleRequest) SetDescription(v string) *UpdateRuleRequest {
  s.Description = &v
  return s
}

func (s *UpdateRuleRequest) SetPriority(v int) *UpdateRuleRequest {
  s.Priority = &v
  return s
}

func (s *UpdateRuleRequest) SetEnabled(v bool) *UpdateRuleRequest {
  s.Enabled = &v
  return s
}

type UpdateRuleRequestAction struct {
  // {"en":"Action name","zh_CN":"动作名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Action options.","zh_CN":"动作参数,详情查看schema接口"}
  Options *UpdateRuleRequestActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s UpdateRuleRequestAction) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleRequestAction) GoString() string {
  return s.String()
}

func (s *UpdateRuleRequestAction) SetName(v string) *UpdateRuleRequestAction {
  s.Name = &v
  return s
}

func (s *UpdateRuleRequestAction) SetOptions(v *UpdateRuleRequestActionOptions) *UpdateRuleRequestAction {
  s.Options = v
  return s
}

type UpdateRuleRequestActionOptions struct {
}

func (s UpdateRuleRequestActionOptions) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleRequestActionOptions) GoString() string {
  return s.String()
}

type UpdateRuleRequestHeader struct {
}

func (s UpdateRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleRequestHeader) GoString() string {
  return s.String()
}

type UpdateRulePaths struct {
  // {"en":"rule ID","zh_CN":"规则ID"}
  RuleId *int `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s UpdateRulePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateRulePaths) GoString() string {
  return s.String()
}

func (s *UpdateRulePaths) SetRuleId(v int) *UpdateRulePaths {
  s.RuleId = &v
  return s
}

type UpdateRuleParameters struct {
}

func (s UpdateRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleParameters) GoString() string {
  return s.String()
}

type UpdateRuleResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *UpdateRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
  return s.String()
}

func (s *UpdateRuleResponse) SetCode(v string) *UpdateRuleResponse {
  s.Code = &v
  return s
}

func (s *UpdateRuleResponse) SetData(v *UpdateRuleResponseData) *UpdateRuleResponse {
  s.Data = v
  return s
}

func (s *UpdateRuleResponse) SetMessage(v string) *UpdateRuleResponse {
  s.Message = &v
  return s
}

type UpdateRuleResponseData struct {
  // {"en":"Rule ID","zh_CN":"规则标识"}
  RuleId *int64 `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s UpdateRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleResponseData) GoString() string {
  return s.String()
}

func (s *UpdateRuleResponseData) SetRuleId(v int64) *UpdateRuleResponseData {
  s.RuleId = &v
  return s
}

type UpdateRuleResponseHeader struct {
}

func (s UpdateRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateRuleRequest struct {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Phase","zh_CN":"阶段","exampleValue":"connectPhase,requestPhase,cachePhase,originPhase,responsePhase"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty" require:"true"`
  // {"en":"The condition in the rule","zh_CN":"规则中的条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"Rule action","zh_CN":"规则动作"}
  Action *CreateRuleRequestAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"Comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"defaultValue":"1","en":"The priority of rule.","zh_CN":"规则排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"The state of rule.","zh_CN":"规则状态。false为禁用，true为启用。","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
  return s.String()
}

func (s *CreateRuleRequest) SetName(v string) *CreateRuleRequest {
  s.Name = &v
  return s
}

func (s *CreateRuleRequest) SetPhase(v string) *CreateRuleRequest {
  s.Phase = &v
  return s
}

func (s *CreateRuleRequest) SetCondition(v string) *CreateRuleRequest {
  s.Condition = &v
  return s
}

func (s *CreateRuleRequest) SetAction(v *CreateRuleRequestAction) *CreateRuleRequest {
  s.Action = v
  return s
}

func (s *CreateRuleRequest) SetDescription(v string) *CreateRuleRequest {
  s.Description = &v
  return s
}

func (s *CreateRuleRequest) SetPriority(v int) *CreateRuleRequest {
  s.Priority = &v
  return s
}

func (s *CreateRuleRequest) SetEnabled(v bool) *CreateRuleRequest {
  s.Enabled = &v
  return s
}

type CreateRuleRequestAction struct {
  // {"en":"Action name","zh_CN":"动作名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Action options.","zh_CN":"动作参数,详情查看schema接口"}
  Options *CreateRuleRequestActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreateRuleRequestAction) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleRequestAction) GoString() string {
  return s.String()
}

func (s *CreateRuleRequestAction) SetName(v string) *CreateRuleRequestAction {
  s.Name = &v
  return s
}

func (s *CreateRuleRequestAction) SetOptions(v *CreateRuleRequestActionOptions) *CreateRuleRequestAction {
  s.Options = v
  return s
}

type CreateRuleRequestActionOptions struct {
}

func (s CreateRuleRequestActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleRequestActionOptions) GoString() string {
  return s.String()
}

type CreateRuleRequestHeader struct {
}

func (s CreateRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleRequestHeader) GoString() string {
  return s.String()
}

type CreateRulePaths struct {
  // {"en":"property ID","zh_CN":"项目ID"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"property version","zh_CN":"项目版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s CreateRulePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateRulePaths) GoString() string {
  return s.String()
}

func (s *CreateRulePaths) SetPropertyId(v int) *CreateRulePaths {
  s.PropertyId = &v
  return s
}

func (s *CreateRulePaths) SetVersion(v int) *CreateRulePaths {
  s.Version = &v
  return s
}

type CreateRuleParameters struct {
}

func (s CreateRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleParameters) GoString() string {
  return s.String()
}

type CreateRuleResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreateRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
  return s.String()
}

func (s *CreateRuleResponse) SetCode(v string) *CreateRuleResponse {
  s.Code = &v
  return s
}

func (s *CreateRuleResponse) SetData(v *CreateRuleResponseData) *CreateRuleResponse {
  s.Data = v
  return s
}

func (s *CreateRuleResponse) SetMessage(v string) *CreateRuleResponse {
  s.Message = &v
  return s
}

type CreateRuleResponseData struct {
  // {"en":"Rule ID","zh_CN":"规则标识"}
  RuleId *int64 `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s CreateRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleResponseData) GoString() string {
  return s.String()
}

func (s *CreateRuleResponseData) SetRuleId(v int64) *CreateRuleResponseData {
  s.RuleId = &v
  return s
}

type CreateRuleResponseHeader struct {
}

func (s CreateRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateRuleResponseHeader) GoString() string {
  return s.String()
}




