package edgefunc

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ListFunctionTemplatesRequest struct {
}

func (s ListFunctionTemplatesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesRequest) GoString() string {
  return s.String()
}

type ListFunctionTemplatesRequestHeader struct {
}

func (s ListFunctionTemplatesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesRequestHeader) GoString() string {
  return s.String()
}

type ListFunctionTemplatesPaths struct {
}

func (s ListFunctionTemplatesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesPaths) GoString() string {
  return s.String()
}

type ListFunctionTemplatesParameters struct {
  // {"defaultValue":"1","en":"Page Number","zh_CN":"页码，默认1"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"defaultValue":"10","en":"Page Size","zh_CN":"每页大小，默认10"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Template name, which supports fuzzy query","zh_CN":"模板名称，支持模糊查询"}
  Name *int `json:"name,omitempty" xml:"name,omitempty"`
  // {"defaultValue":"false","en":"is english","zh_CN":"是否英文"}
  IsEnglish *string `json:"isEnglish,omitempty" xml:"isEnglish,omitempty"`
}

func (s ListFunctionTemplatesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesParameters) GoString() string {
  return s.String()
}

func (s *ListFunctionTemplatesParameters) SetPageNo(v int) *ListFunctionTemplatesParameters {
  s.PageNo = &v
  return s
}

func (s *ListFunctionTemplatesParameters) SetPageSize(v int) *ListFunctionTemplatesParameters {
  s.PageSize = &v
  return s
}

func (s *ListFunctionTemplatesParameters) SetName(v int) *ListFunctionTemplatesParameters {
  s.Name = &v
  return s
}

func (s *ListFunctionTemplatesParameters) SetIsEnglish(v string) *ListFunctionTemplatesParameters {
  s.IsEnglish = &v
  return s
}

type ListFunctionTemplatesResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data body","zh_CN":"响应数据体"}
  Data *ListFunctionTemplatesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListFunctionTemplatesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesResponse) GoString() string {
  return s.String()
}

func (s *ListFunctionTemplatesResponse) SetCode(v int) *ListFunctionTemplatesResponse {
  s.Code = &v
  return s
}

func (s *ListFunctionTemplatesResponse) SetMessage(v string) *ListFunctionTemplatesResponse {
  s.Message = &v
  return s
}

func (s *ListFunctionTemplatesResponse) SetData(v *ListFunctionTemplatesResponseData) *ListFunctionTemplatesResponse {
  s.Data = v
  return s
}

type ListFunctionTemplatesResponseData struct {
  // {"en":"Current page number","zh_CN":"当前页码"}
  PageNum *int `json:"pageNum,omitempty" xml:"pageNum,omitempty" require:"true"`
  // {"en":"Page size","zh_CN":"每页大小"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total number of records","zh_CN":"总记录数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of function templates","zh_CN":"函数模板列表"}
  Templates []*ListFunctionTemplatesResponseDataTemplates `json:"templates,omitempty" xml:"templates,omitempty" require:"true" type:"Repeated"`
}

func (s ListFunctionTemplatesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesResponseData) GoString() string {
  return s.String()
}

func (s *ListFunctionTemplatesResponseData) SetPageNum(v int) *ListFunctionTemplatesResponseData {
  s.PageNum = &v
  return s
}

func (s *ListFunctionTemplatesResponseData) SetPageSize(v int) *ListFunctionTemplatesResponseData {
  s.PageSize = &v
  return s
}

func (s *ListFunctionTemplatesResponseData) SetTotal(v int) *ListFunctionTemplatesResponseData {
  s.Total = &v
  return s
}

func (s *ListFunctionTemplatesResponseData) SetTemplates(v []*ListFunctionTemplatesResponseDataTemplates) *ListFunctionTemplatesResponseData {
  s.Templates = v
  return s
}

type ListFunctionTemplatesResponseDataTemplates struct     {
  // {"en":"Template ID","zh_CN":"模板ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Template code","zh_CN":"模板编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Template Name","zh_CN":"模板名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Template desc","zh_CN":"模板描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ListFunctionTemplatesResponseDataTemplates) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesResponseDataTemplates) GoString() string {
  return s.String()
}

func (s *ListFunctionTemplatesResponseDataTemplates) SetId(v int) *ListFunctionTemplatesResponseDataTemplates {
  s.Id = &v
  return s
}

func (s *ListFunctionTemplatesResponseDataTemplates) SetCode(v string) *ListFunctionTemplatesResponseDataTemplates {
  s.Code = &v
  return s
}

func (s *ListFunctionTemplatesResponseDataTemplates) SetName(v string) *ListFunctionTemplatesResponseDataTemplates {
  s.Name = &v
  return s
}

func (s *ListFunctionTemplatesResponseDataTemplates) SetDescription(v string) *ListFunctionTemplatesResponseDataTemplates {
  s.Description = &v
  return s
}

type ListFunctionTemplatesResponseHeader struct {
}

func (s ListFunctionTemplatesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListFunctionTemplatesResponseHeader) GoString() string {
  return s.String()
}




type GetEdgeFunctionInfoRequest struct {
}

func (s GetEdgeFunctionInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoRequest) GoString() string {
  return s.String()
}

type GetEdgeFunctionInfoRequestHeader struct {
}

func (s GetEdgeFunctionInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeFunctionInfoPaths struct {
  // {"en":"Function ID","zh_CN":"函数ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetEdgeFunctionInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoPaths) GoString() string {
  return s.String()
}

func (s *GetEdgeFunctionInfoPaths) SetId(v string) *GetEdgeFunctionInfoPaths {
  s.Id = &v
  return s
}

type GetEdgeFunctionInfoParameters struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetEdgeFunctionInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeFunctionInfoParameters) SetId(v string) *GetEdgeFunctionInfoParameters {
  s.Id = &v
  return s
}

type GetEdgeFunctionInfoResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"状态信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data body","zh_CN":"响应数据体"}
  Data *GetEdgeFunctionInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetEdgeFunctionInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeFunctionInfoResponse) SetCode(v int) *GetEdgeFunctionInfoResponse {
  s.Code = &v
  return s
}

func (s *GetEdgeFunctionInfoResponse) SetMessage(v string) *GetEdgeFunctionInfoResponse {
  s.Message = &v
  return s
}

func (s *GetEdgeFunctionInfoResponse) SetData(v *GetEdgeFunctionInfoResponseData) *GetEdgeFunctionInfoResponse {
  s.Data = v
  return s
}

type GetEdgeFunctionInfoResponseData struct {
  // {"en":"Function info body","zh_CN":"函数信息结构体"}
  Func *GetEdgeFunctionInfoResponseDataFunc `json:"func,omitempty" xml:"func,omitempty" require:"true" type:"Struct"`
}

func (s GetEdgeFunctionInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoResponseData) GoString() string {
  return s.String()
}

func (s *GetEdgeFunctionInfoResponseData) SetFunc(v *GetEdgeFunctionInfoResponseDataFunc) *GetEdgeFunctionInfoResponseData {
  s.Func = v
  return s
}

type GetEdgeFunctionInfoResponseDataFunc struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Function name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Function alias","zh_CN":"函数别名"}
  FuncAlias *string `json:"funcAlias,omitempty" xml:"funcAlias,omitempty" require:"true"`
  // {"en":"Test domain","zh_CN":"测试域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"comment","zh_CN":"备注"}
  Memo *string `json:"memo,omitempty" xml:"memo,omitempty" require:"true"`
  // {"en":"Create time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update time","zh_CN":"更新时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s GetEdgeFunctionInfoResponseDataFunc) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoResponseDataFunc) GoString() string {
  return s.String()
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetId(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.Id = &v
  return s
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetFuncName(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.FuncName = &v
  return s
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetFuncAlias(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.FuncAlias = &v
  return s
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetDomain(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.Domain = &v
  return s
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetMemo(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.Memo = &v
  return s
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetCreateTime(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.CreateTime = &v
  return s
}

func (s *GetEdgeFunctionInfoResponseDataFunc) SetUpdateTime(v string) *GetEdgeFunctionInfoResponseDataFunc {
  s.UpdateTime = &v
  return s
}

type GetEdgeFunctionInfoResponseHeader struct {
}

func (s GetEdgeFunctionInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeFunctionInfoResponseHeader) GoString() string {
  return s.String()
}




type DeleteFuncDomainTriggerRequest struct {
}

func (s DeleteFuncDomainTriggerRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerRequest) GoString() string {
  return s.String()
}

type DeleteFuncDomainTriggerRequestHeader struct {
}

func (s DeleteFuncDomainTriggerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerRequestHeader) GoString() string {
  return s.String()
}

type DeleteFuncDomainTriggerPaths struct {
  // {"en":"Domain ID","zh_CN":"域名的 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteFuncDomainTriggerPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerPaths) GoString() string {
  return s.String()
}

func (s *DeleteFuncDomainTriggerPaths) SetId(v string) *DeleteFuncDomainTriggerPaths {
  s.Id = &v
  return s
}

type DeleteFuncDomainTriggerParameters struct {
}

func (s DeleteFuncDomainTriggerParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerParameters) GoString() string {
  return s.String()
}

type DeleteFuncDomainTriggerResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status description","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteFuncDomainTriggerResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerResponse) GoString() string {
  return s.String()
}

func (s *DeleteFuncDomainTriggerResponse) SetCode(v int) *DeleteFuncDomainTriggerResponse {
  s.Code = &v
  return s
}

func (s *DeleteFuncDomainTriggerResponse) SetMessage(v string) *DeleteFuncDomainTriggerResponse {
  s.Message = &v
  return s
}

type DeleteFuncDomainTriggerResponseHeader struct {
}

func (s DeleteFuncDomainTriggerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerResponseHeader) GoString() string {
  return s.String()
}




type QueryFunctionCodeRequest struct {
}

func (s QueryFunctionCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionCodeRequest) GoString() string {
  return s.String()
}

type QueryFunctionCodeRequestHeader struct {
}

func (s QueryFunctionCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryFunctionCodePaths struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryFunctionCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionCodePaths) GoString() string {
  return s.String()
}

func (s *QueryFunctionCodePaths) SetId(v string) *QueryFunctionCodePaths {
  s.Id = &v
  return s
}

type QueryFunctionCodeParameters struct {
  // {"en":"Function ID","zh_CN":"函数ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s QueryFunctionCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionCodeParameters) GoString() string {
  return s.String()
}

func (s *QueryFunctionCodeParameters) SetId(v string) *QueryFunctionCodeParameters {
  s.Id = &v
  return s
}

type QueryFunctionCodeResponse struct {
}

func (s QueryFunctionCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionCodeResponse) GoString() string {
  return s.String()
}

type QueryFunctionCodeResponseHeader struct {
}

func (s QueryFunctionCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionCodeResponseHeader) GoString() string {
  return s.String()
}




type DeleteEdgeFunctionRequest struct {
}

func (s DeleteEdgeFunctionRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeFunctionRequest) GoString() string {
  return s.String()
}

type DeleteEdgeFunctionRequestHeader struct {
}

func (s DeleteEdgeFunctionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeFunctionRequestHeader) GoString() string {
  return s.String()
}

type DeleteEdgeFunctionPaths struct {
  // {"en":"The function ID to be removed","zh_CN":"要删除的函数ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteEdgeFunctionPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeFunctionPaths) GoString() string {
  return s.String()
}

func (s *DeleteEdgeFunctionPaths) SetId(v string) *DeleteEdgeFunctionPaths {
  s.Id = &v
  return s
}

type DeleteEdgeFunctionParameters struct {
}

func (s DeleteEdgeFunctionParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeFunctionParameters) GoString() string {
  return s.String()
}

type DeleteEdgeFunctionResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"状态信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteEdgeFunctionResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeFunctionResponse) GoString() string {
  return s.String()
}

func (s *DeleteEdgeFunctionResponse) SetCode(v int) *DeleteEdgeFunctionResponse {
  s.Code = &v
  return s
}

func (s *DeleteEdgeFunctionResponse) SetMessage(v string) *DeleteEdgeFunctionResponse {
  s.Message = &v
  return s
}

type DeleteEdgeFunctionResponseHeader struct {
}

func (s DeleteEdgeFunctionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEdgeFunctionResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeFuncTriggerRequest struct {
}

func (s QueryEdgeFuncTriggerRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerRequest) GoString() string {
  return s.String()
}

type QueryEdgeFuncTriggerRequestHeader struct {
}

func (s QueryEdgeFuncTriggerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeFuncTriggerPaths struct {
}

func (s QueryEdgeFuncTriggerPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerPaths) GoString() string {
  return s.String()
}

type QueryEdgeFuncTriggerParameters struct {
  // {"en":"Domain, Either domain or function name is required.","zh_CN":"域名，与函数任填其一"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Function name, Either domain or function name is required.","zh_CN":"函数名称，与域名任填其一"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty"`
  // {"en":"The page number for the query, starting from 1.","zh_CN":"查询的页码，从1开始"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"The number of items to return per page.","zh_CN":"每页返回的条目数量"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryEdgeFuncTriggerParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerParameters) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerParameters) SetDomain(v string) *QueryEdgeFuncTriggerParameters {
  s.Domain = &v
  return s
}

func (s *QueryEdgeFuncTriggerParameters) SetFuncName(v string) *QueryEdgeFuncTriggerParameters {
  s.FuncName = &v
  return s
}

func (s *QueryEdgeFuncTriggerParameters) SetPageNo(v int) *QueryEdgeFuncTriggerParameters {
  s.PageNo = &v
  return s
}

func (s *QueryEdgeFuncTriggerParameters) SetPageSize(v int) *QueryEdgeFuncTriggerParameters {
  s.PageSize = &v
  return s
}

type QueryEdgeFuncTriggerResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status description","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response Data","zh_CN":"响应数据"}
  Data *QueryEdgeFuncTriggerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeFuncTriggerResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponse) SetCode(v string) *QueryEdgeFuncTriggerResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponse) SetMessage(v string) *QueryEdgeFuncTriggerResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponse) SetData(v *QueryEdgeFuncTriggerResponseData) *QueryEdgeFuncTriggerResponse {
  s.Data = v
  return s
}

type QueryEdgeFuncTriggerResponseData struct {
  // {"en":"Trigger List","zh_CN":"触发器列表"}
  Result []*QueryEdgeFuncTriggerResponseDataResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeFuncTriggerResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponseData) SetResult(v []*QueryEdgeFuncTriggerResponseDataResult) *QueryEdgeFuncTriggerResponseData {
  s.Result = v
  return s
}

type QueryEdgeFuncTriggerResponseDataResult struct     {
  // {"en":"Domain ID","zh_CN":"域名 iD"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Router List","zh_CN":"路由列表"}
  Routers *QueryEdgeFuncTriggerResponseDataResultRouters `json:"routers,omitempty" xml:"routers,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeFuncTriggerResponseDataResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseDataResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponseDataResult) SetId(v string) *QueryEdgeFuncTriggerResponseDataResult {
  s.Id = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponseDataResult) SetDomain(v string) *QueryEdgeFuncTriggerResponseDataResult {
  s.Domain = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponseDataResult) SetRouters(v *QueryEdgeFuncTriggerResponseDataResultRouters) *QueryEdgeFuncTriggerResponseDataResult {
  s.Routers = v
  return s
}

type QueryEdgeFuncTriggerResponseDataResultRouters struct {
  // {"en":"Func name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Match path","zh_CN":"匹配路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s QueryEdgeFuncTriggerResponseDataResultRouters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseDataResultRouters) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponseDataResultRouters) SetFuncName(v string) *QueryEdgeFuncTriggerResponseDataResultRouters {
  s.FuncName = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponseDataResultRouters) SetPath(v string) *QueryEdgeFuncTriggerResponseDataResultRouters {
  s.Path = &v
  return s
}

type QueryEdgeFuncTriggerResponseHeader struct {
}

func (s QueryEdgeFuncTriggerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseHeader) GoString() string {
  return s.String()
}




type SaveFunctionCodeRequest struct {
}

func (s SaveFunctionCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s SaveFunctionCodeRequest) GoString() string {
  return s.String()
}

type SaveFunctionCodeRequestHeader struct {
}

func (s SaveFunctionCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SaveFunctionCodeRequestHeader) GoString() string {
  return s.String()
}

type SaveFunctionCodePaths struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s SaveFunctionCodePaths) String() string {
  return tea.Prettify(s)
}

func (s SaveFunctionCodePaths) GoString() string {
  return s.String()
}

func (s *SaveFunctionCodePaths) SetId(v string) *SaveFunctionCodePaths {
  s.Id = &v
  return s
}

type SaveFunctionCodeParameters struct {
}

func (s SaveFunctionCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s SaveFunctionCodeParameters) GoString() string {
  return s.String()
}

type SaveFunctionCodeResponse struct {
}

func (s SaveFunctionCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s SaveFunctionCodeResponse) GoString() string {
  return s.String()
}

type SaveFunctionCodeResponseHeader struct {
}

func (s SaveFunctionCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SaveFunctionCodeResponseHeader) GoString() string {
  return s.String()
}




type CreateEdgeFuncTriggerRequest struct {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Router list","zh_CN":"路由列表"}
  Routers []*CreateEdgeFuncTriggerRequestRouters `json:"routers,omitempty" xml:"routers,omitempty" require:"true" type:"Repeated"`
}

func (s CreateEdgeFuncTriggerRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerRequest) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerRequest) SetDomain(v string) *CreateEdgeFuncTriggerRequest {
  s.Domain = &v
  return s
}

func (s *CreateEdgeFuncTriggerRequest) SetRouters(v []*CreateEdgeFuncTriggerRequestRouters) *CreateEdgeFuncTriggerRequest {
  s.Routers = v
  return s
}

type CreateEdgeFuncTriggerRequestRouters struct     {
  // {"en":"Func name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Match path","zh_CN":"匹配路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s CreateEdgeFuncTriggerRequestRouters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerRequestRouters) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerRequestRouters) SetFuncName(v string) *CreateEdgeFuncTriggerRequestRouters {
  s.FuncName = &v
  return s
}

func (s *CreateEdgeFuncTriggerRequestRouters) SetPath(v string) *CreateEdgeFuncTriggerRequestRouters {
  s.Path = &v
  return s
}

type CreateEdgeFuncTriggerRequestHeader struct {
}

func (s CreateEdgeFuncTriggerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerRequestHeader) GoString() string {
  return s.String()
}

type CreateEdgeFuncTriggerPaths struct {
}

func (s CreateEdgeFuncTriggerPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerPaths) GoString() string {
  return s.String()
}

type CreateEdgeFuncTriggerParameters struct {
}

func (s CreateEdgeFuncTriggerParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerParameters) GoString() string {
  return s.String()
}

type CreateEdgeFuncTriggerResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status description","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response Data","zh_CN":"响应数据"}
  Data *CreateEdgeFuncTriggerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateEdgeFuncTriggerResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponse) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerResponse) SetCode(v string) *CreateEdgeFuncTriggerResponse {
  s.Code = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponse) SetMessage(v string) *CreateEdgeFuncTriggerResponse {
  s.Message = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponse) SetData(v *CreateEdgeFuncTriggerResponseData) *CreateEdgeFuncTriggerResponse {
  s.Data = v
  return s
}

type CreateEdgeFuncTriggerResponseData struct {
  // {"en":"Domain ID","zh_CN":"域名 iD"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Router List","zh_CN":"路由列表"}
  Routers *CreateEdgeFuncTriggerResponseDataRouters `json:"routers,omitempty" xml:"routers,omitempty" require:"true" type:"Struct"`
}

func (s CreateEdgeFuncTriggerResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponseData) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerResponseData) SetId(v string) *CreateEdgeFuncTriggerResponseData {
  s.Id = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponseData) SetDomain(v string) *CreateEdgeFuncTriggerResponseData {
  s.Domain = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponseData) SetRouters(v *CreateEdgeFuncTriggerResponseDataRouters) *CreateEdgeFuncTriggerResponseData {
  s.Routers = v
  return s
}

type CreateEdgeFuncTriggerResponseDataRouters struct {
  // {"en":"Func name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Match path","zh_CN":"匹配路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s CreateEdgeFuncTriggerResponseDataRouters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponseDataRouters) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerResponseDataRouters) SetFuncName(v string) *CreateEdgeFuncTriggerResponseDataRouters {
  s.FuncName = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponseDataRouters) SetPath(v string) *CreateEdgeFuncTriggerResponseDataRouters {
  s.Path = &v
  return s
}

type CreateEdgeFuncTriggerResponseHeader struct {
}

func (s CreateEdgeFuncTriggerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponseHeader) GoString() string {
  return s.String()
}




type EdgeFuncGetDebugLogRequest struct {
}

func (s EdgeFuncGetDebugLogRequest) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogRequest) GoString() string {
  return s.String()
}

type EdgeFuncGetDebugLogRequestHeader struct {
}

func (s EdgeFuncGetDebugLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogRequestHeader) GoString() string {
  return s.String()
}

type EdgeFuncGetDebugLogPaths struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Debug session ID","zh_CN":"调试会话ID"}
  DebugId *string `json:"debugId,omitempty" xml:"debugId,omitempty" require:"true"`
}

func (s EdgeFuncGetDebugLogPaths) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogPaths) GoString() string {
  return s.String()
}

func (s *EdgeFuncGetDebugLogPaths) SetId(v string) *EdgeFuncGetDebugLogPaths {
  s.Id = &v
  return s
}

func (s *EdgeFuncGetDebugLogPaths) SetDebugId(v string) *EdgeFuncGetDebugLogPaths {
  s.DebugId = &v
  return s
}

type EdgeFuncGetDebugLogParameters struct {
}

func (s EdgeFuncGetDebugLogParameters) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogParameters) GoString() string {
  return s.String()
}

type EdgeFuncGetDebugLogResponse struct {
  // {"en":"Code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"状态信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data body","zh_CN":"响应数据结构体"}
  Data *EdgeFuncGetDebugLogResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s EdgeFuncGetDebugLogResponse) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogResponse) GoString() string {
  return s.String()
}

func (s *EdgeFuncGetDebugLogResponse) SetCode(v int) *EdgeFuncGetDebugLogResponse {
  s.Code = &v
  return s
}

func (s *EdgeFuncGetDebugLogResponse) SetMessage(v string) *EdgeFuncGetDebugLogResponse {
  s.Message = &v
  return s
}

func (s *EdgeFuncGetDebugLogResponse) SetData(v *EdgeFuncGetDebugLogResponseData) *EdgeFuncGetDebugLogResponse {
  s.Data = v
  return s
}

type EdgeFuncGetDebugLogResponseData struct {
  // {"en":"Log content generated by function execution","zh_CN":"函数执行产生的日志内容"}
  Logs *string `json:"logs,omitempty" xml:"logs,omitempty" require:"true"`
  // {"en":"Timestamp","zh_CN":"日志获取时间戳"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
}

func (s EdgeFuncGetDebugLogResponseData) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogResponseData) GoString() string {
  return s.String()
}

func (s *EdgeFuncGetDebugLogResponseData) SetLogs(v string) *EdgeFuncGetDebugLogResponseData {
  s.Logs = &v
  return s
}

func (s *EdgeFuncGetDebugLogResponseData) SetTimestamp(v string) *EdgeFuncGetDebugLogResponseData {
  s.Timestamp = &v
  return s
}

type EdgeFuncGetDebugLogResponseHeader struct {
}

func (s EdgeFuncGetDebugLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugLogResponseHeader) GoString() string {
  return s.String()
}




type UploadFuncCodeRequest struct {
  // {"en":"func name","zh_CN":"函数名称"}
  FuncName *string `json:"FuncName,omitempty" xml:"FuncName,omitempty"`
  // {"en":"commit message","zh_CN":"提交信息"}
  CommitMessage *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
  // {"en":"func file list","zh_CN":"函数文件列表"}
  Files []*UploadFuncCodeRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s UploadFuncCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequest) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeRequest) SetFuncName(v string) *UploadFuncCodeRequest {
  s.FuncName = &v
  return s
}

func (s *UploadFuncCodeRequest) SetCommitMessage(v string) *UploadFuncCodeRequest {
  s.CommitMessage = &v
  return s
}

func (s *UploadFuncCodeRequest) SetFiles(v []*UploadFuncCodeRequestFiles) *UploadFuncCodeRequest {
  s.Files = v
  return s
}

type UploadFuncCodeRequestFiles struct     {
  // {"en":"FileName","zh_CN":"文件名称"}
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // {"en":"File code (base64 encoding)","zh_CN":"文件所含代码（Base64 编码）"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UploadFuncCodeRequestFiles) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequestFiles) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeRequestFiles) SetFileName(v string) *UploadFuncCodeRequestFiles {
  s.FileName = &v
  return s
}

func (s *UploadFuncCodeRequestFiles) SetCode(v string) *UploadFuncCodeRequestFiles {
  s.Code = &v
  return s
}

type UploadFuncCodeRequestHeader struct {
}

func (s UploadFuncCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequestHeader) GoString() string {
  return s.String()
}

type UploadFuncCodePaths struct {
}

func (s UploadFuncCodePaths) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodePaths) GoString() string {
  return s.String()
}

type UploadFuncCodeParameters struct {
}

func (s UploadFuncCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeParameters) GoString() string {
  return s.String()
}

type UploadFuncCodeResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"状态码描述"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s UploadFuncCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeResponse) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeResponse) SetCode(v string) *UploadFuncCodeResponse {
  s.Code = &v
  return s
}

func (s *UploadFuncCodeResponse) SetMessage(v string) *UploadFuncCodeResponse {
  s.Message = &v
  return s
}

type UploadFuncCodeResponseHeader struct {
}

func (s UploadFuncCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeResponseHeader) GoString() string {
  return s.String()
}




type EdgeFuncGetDebugUrlRequest struct {
}

func (s EdgeFuncGetDebugUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlRequest) GoString() string {
  return s.String()
}

type EdgeFuncGetDebugUrlRequestHeader struct {
}

func (s EdgeFuncGetDebugUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlRequestHeader) GoString() string {
  return s.String()
}

type EdgeFuncGetDebugUrlPaths struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s EdgeFuncGetDebugUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlPaths) GoString() string {
  return s.String()
}

func (s *EdgeFuncGetDebugUrlPaths) SetId(v string) *EdgeFuncGetDebugUrlPaths {
  s.Id = &v
  return s
}

type EdgeFuncGetDebugUrlParameters struct {
}

func (s EdgeFuncGetDebugUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlParameters) GoString() string {
  return s.String()
}

type EdgeFuncGetDebugUrlResponse struct {
  // {"en":"Code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"状态信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data body","zh_CN":"响应数据结构体"}
  Data *EdgeFuncGetDebugUrlResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s EdgeFuncGetDebugUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlResponse) GoString() string {
  return s.String()
}

func (s *EdgeFuncGetDebugUrlResponse) SetCode(v int) *EdgeFuncGetDebugUrlResponse {
  s.Code = &v
  return s
}

func (s *EdgeFuncGetDebugUrlResponse) SetMessage(v string) *EdgeFuncGetDebugUrlResponse {
  s.Message = &v
  return s
}

func (s *EdgeFuncGetDebugUrlResponse) SetData(v *EdgeFuncGetDebugUrlResponseData) *EdgeFuncGetDebugUrlResponse {
  s.Data = v
  return s
}

type EdgeFuncGetDebugUrlResponseData struct {
  // {"en":"Debug URL","zh_CN":"调试 URL"}
  DebugUrl *string `json:"debugUrl,omitempty" xml:"debugUrl,omitempty" require:"true"`
  // {"en":"Debug ID","zh_CN":"调试 ID"}
  DebugId *string `json:"debugId,omitempty" xml:"debugId,omitempty" require:"true"`
}

func (s EdgeFuncGetDebugUrlResponseData) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlResponseData) GoString() string {
  return s.String()
}

func (s *EdgeFuncGetDebugUrlResponseData) SetDebugUrl(v string) *EdgeFuncGetDebugUrlResponseData {
  s.DebugUrl = &v
  return s
}

func (s *EdgeFuncGetDebugUrlResponseData) SetDebugId(v string) *EdgeFuncGetDebugUrlResponseData {
  s.DebugId = &v
  return s
}

type EdgeFuncGetDebugUrlResponseHeader struct {
}

func (s EdgeFuncGetDebugUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EdgeFuncGetDebugUrlResponseHeader) GoString() string {
  return s.String()
}




type CreateEdgeFuncRequest struct {
  // {"en":"function name","zh_CN":"函数名称","exampleValue":"calculateTotal"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty"`
  // {"en":"function alias","zh_CN":"函数别名","exampleValue":"calcTotal"}
  FuncAlias *string `json:"funcAlias,omitempty" xml:"funcAlias,omitempty"`
  // {"en":"test domain","zh_CN":"测试域名","exampleValue":"test.example.com"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"template name, specifies the function template to use","zh_CN":"模板名称，指定使用的函数模板","exampleValue":"SumCalculatorTemplate"}
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // {"en":"remarks","zh_CN":"备注","exampleValue":"optional notes"}
  Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s CreateEdgeFuncRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncRequest) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncRequest) SetFuncName(v string) *CreateEdgeFuncRequest {
  s.FuncName = &v
  return s
}

func (s *CreateEdgeFuncRequest) SetFuncAlias(v string) *CreateEdgeFuncRequest {
  s.FuncAlias = &v
  return s
}

func (s *CreateEdgeFuncRequest) SetDomain(v string) *CreateEdgeFuncRequest {
  s.Domain = &v
  return s
}

func (s *CreateEdgeFuncRequest) SetTemplateName(v string) *CreateEdgeFuncRequest {
  s.TemplateName = &v
  return s
}

func (s *CreateEdgeFuncRequest) SetMemo(v string) *CreateEdgeFuncRequest {
  s.Memo = &v
  return s
}

type CreateEdgeFuncRequestHeader struct {
}

func (s CreateEdgeFuncRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncRequestHeader) GoString() string {
  return s.String()
}

type CreateEdgeFuncPaths struct {
}

func (s CreateEdgeFuncPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncPaths) GoString() string {
  return s.String()
}

type CreateEdgeFuncParameters struct {
}

func (s CreateEdgeFuncParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncParameters) GoString() string {
  return s.String()
}

type CreateEdgeFuncResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"状态信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Create a successful function ID","zh_CN":"创建成功的函数ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateEdgeFuncResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncResponse) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncResponse) SetCode(v int) *CreateEdgeFuncResponse {
  s.Code = &v
  return s
}

func (s *CreateEdgeFuncResponse) SetMessage(v string) *CreateEdgeFuncResponse {
  s.Message = &v
  return s
}

func (s *CreateEdgeFuncResponse) SetId(v string) *CreateEdgeFuncResponse {
  s.Id = &v
  return s
}

type CreateEdgeFuncResponseHeader struct {
}

func (s CreateEdgeFuncResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncResponseHeader) GoString() string {
  return s.String()
}




type UpdateEdgeFuncRequest struct {
  // {"en":"Function name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Function alias","zh_CN":"函数别名"}
  FuncAlias *string `json:"funcAlias,omitempty" xml:"funcAlias,omitempty" require:"true"`
  // {"en":"Test domain","zh_CN":"测试域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Comment","zh_CN":"备注"}
  Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
}

func (s UpdateEdgeFuncRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeFuncRequest) GoString() string {
  return s.String()
}

func (s *UpdateEdgeFuncRequest) SetFuncName(v string) *UpdateEdgeFuncRequest {
  s.FuncName = &v
  return s
}

func (s *UpdateEdgeFuncRequest) SetFuncAlias(v string) *UpdateEdgeFuncRequest {
  s.FuncAlias = &v
  return s
}

func (s *UpdateEdgeFuncRequest) SetDomain(v string) *UpdateEdgeFuncRequest {
  s.Domain = &v
  return s
}

func (s *UpdateEdgeFuncRequest) SetMemo(v string) *UpdateEdgeFuncRequest {
  s.Memo = &v
  return s
}

type UpdateEdgeFuncRequestHeader struct {
}

func (s UpdateEdgeFuncRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeFuncRequestHeader) GoString() string {
  return s.String()
}

type UpdateEdgeFuncPaths struct {
  // {"en":"Function ID","zh_CN":"函数ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateEdgeFuncPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeFuncPaths) GoString() string {
  return s.String()
}

func (s *UpdateEdgeFuncPaths) SetId(v string) *UpdateEdgeFuncPaths {
  s.Id = &v
  return s
}

type UpdateEdgeFuncParameters struct {
  // {"en":"Function ID","zh_CN":"函数 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateEdgeFuncParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeFuncParameters) GoString() string {
  return s.String()
}

func (s *UpdateEdgeFuncParameters) SetId(v string) *UpdateEdgeFuncParameters {
  s.Id = &v
  return s
}

type UpdateEdgeFuncResponse struct {
  // {"en":"code","zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateEdgeFuncResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeFuncResponse) GoString() string {
  return s.String()
}

func (s *UpdateEdgeFuncResponse) SetCode(v string) *UpdateEdgeFuncResponse {
  s.Code = &v
  return s
}

func (s *UpdateEdgeFuncResponse) SetMessage(v string) *UpdateEdgeFuncResponse {
  s.Message = &v
  return s
}

type UpdateEdgeFuncResponseHeader struct {
}

func (s UpdateEdgeFuncResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEdgeFuncResponseHeader) GoString() string {
  return s.String()
}




type QueryFunctionListRequest struct {
}

func (s QueryFunctionListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListRequest) GoString() string {
  return s.String()
}

type QueryFunctionListRequestHeader struct {
}

func (s QueryFunctionListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListRequestHeader) GoString() string {
  return s.String()
}

type QueryFunctionListPaths struct {
}

func (s QueryFunctionListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListPaths) GoString() string {
  return s.String()
}

type QueryFunctionListParameters struct {
  // {"en":"Function ID (for exact query)","zh_CN":"函数ID（精确查询）"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Function name (for fuzzy query)","zh_CN":"函数名称（模糊查询）"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty"`
  // {"en":"Page number","zh_CN":"页码"}
  PageNo *string `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"Page size","zh_CN":"每页大小"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryFunctionListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListParameters) GoString() string {
  return s.String()
}

func (s *QueryFunctionListParameters) SetId(v string) *QueryFunctionListParameters {
  s.Id = &v
  return s
}

func (s *QueryFunctionListParameters) SetFuncName(v string) *QueryFunctionListParameters {
  s.FuncName = &v
  return s
}

func (s *QueryFunctionListParameters) SetPageNo(v string) *QueryFunctionListParameters {
  s.PageNo = &v
  return s
}

func (s *QueryFunctionListParameters) SetPageSize(v string) *QueryFunctionListParameters {
  s.PageSize = &v
  return s
}

type QueryFunctionListResponse struct {
  // {"en":"Status code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status message","zh_CN":"状态信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data body","zh_CN":"响应数据体"}
  Data *QueryFunctionListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryFunctionListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListResponse) GoString() string {
  return s.String()
}

func (s *QueryFunctionListResponse) SetCode(v int) *QueryFunctionListResponse {
  s.Code = &v
  return s
}

func (s *QueryFunctionListResponse) SetMessage(v string) *QueryFunctionListResponse {
  s.Message = &v
  return s
}

func (s *QueryFunctionListResponse) SetData(v *QueryFunctionListResponseData) *QueryFunctionListResponse {
  s.Data = v
  return s
}

type QueryFunctionListResponseData struct {
  // {"en":"Total","zh_CN":"总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of functions","zh_CN":"函数列表"}
  Funcs []*string `json:"funcs,omitempty" xml:"funcs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryFunctionListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListResponseData) GoString() string {
  return s.String()
}

func (s *QueryFunctionListResponseData) SetTotal(v int) *QueryFunctionListResponseData {
  s.Total = &v
  return s
}

func (s *QueryFunctionListResponseData) SetFuncs(v []*string) *QueryFunctionListResponseData {
  s.Funcs = v
  return s
}

type QueryFunctionListResponseHeader struct {
}

func (s QueryFunctionListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryFunctionListResponseHeader) GoString() string {
  return s.String()
}




