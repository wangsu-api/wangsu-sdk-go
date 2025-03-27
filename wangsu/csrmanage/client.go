package csrmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCsrServiceRequest struct {
}

func (s QueryCsrServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceRequest) GoString() string {
  return s.String()
}

type QueryCsrServiceResponse struct {
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Request result data", "zh_CN":"请求结果数据"}
  Data *QueryCsrServiceQueryCsrServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"CSR ID", "zh_CN":"CSR ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"CSR name, which cannot be repeated", "zh_CN":"csr名称，不能重复"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Key algorithm: RSA/EC", "zh_CN":"密钥算法：RSA/EC 二选一"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"csr content", "zh_CN":"CSR内容"}
  CsrFile *string `json:"csr-file,omitempty" xml:"csr-file,omitempty" require:"true"`
  // {"en":"The main domain name", "zh_CN":"主域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Backup domain name list", "zh_CN":"备份域名列表"}
  Sans []*string `json:"sans,omitempty" xml:"sans,omitempty" require:"true" type:"Repeated"`
  // {"en":"comment", "zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Country or region", "zh_CN":"国家地区"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"state", "zh_CN":"州"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"city", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"company", "zh_CN":"公司"}
  Company *string `json:"company,omitempty" xml:"company,omitempty" require:"true"`
  // {"en":"department", "zh_CN":"部门"}
  Department *string `json:"department,omitempty" xml:"department,omitempty" require:"true"`
}

func (s QueryCsrServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryCsrServiceResponse) SetCode(v string) *QueryCsrServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryCsrServiceResponse) SetMessage(v string) *QueryCsrServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryCsrServiceResponse) SetData(v *QueryCsrServiceQueryCsrServiceResponseData) *QueryCsrServiceResponse {
  s.Data = v
  return s
}

func (s *QueryCsrServiceResponse) SetId(v string) *QueryCsrServiceResponse {
  s.Id = &v
  return s
}

func (s *QueryCsrServiceResponse) SetName(v string) *QueryCsrServiceResponse {
  s.Name = &v
  return s
}

func (s *QueryCsrServiceResponse) SetAlgorithm(v string) *QueryCsrServiceResponse {
  s.Algorithm = &v
  return s
}

func (s *QueryCsrServiceResponse) SetCsrFile(v string) *QueryCsrServiceResponse {
  s.CsrFile = &v
  return s
}

func (s *QueryCsrServiceResponse) SetDomain(v string) *QueryCsrServiceResponse {
  s.Domain = &v
  return s
}

func (s *QueryCsrServiceResponse) SetSans(v []*string) *QueryCsrServiceResponse {
  s.Sans = v
  return s
}

func (s *QueryCsrServiceResponse) SetComment(v string) *QueryCsrServiceResponse {
  s.Comment = &v
  return s
}

func (s *QueryCsrServiceResponse) SetCountry(v string) *QueryCsrServiceResponse {
  s.Country = &v
  return s
}

func (s *QueryCsrServiceResponse) SetState(v string) *QueryCsrServiceResponse {
  s.State = &v
  return s
}

func (s *QueryCsrServiceResponse) SetCity(v string) *QueryCsrServiceResponse {
  s.City = &v
  return s
}

func (s *QueryCsrServiceResponse) SetCompany(v string) *QueryCsrServiceResponse {
  s.Company = &v
  return s
}

func (s *QueryCsrServiceResponse) SetDepartment(v string) *QueryCsrServiceResponse {
  s.Department = &v
  return s
}

type QueryCsrServiceQueryCsrServiceResponseData struct {
}

func (s QueryCsrServiceQueryCsrServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceQueryCsrServiceResponseData) GoString() string {
  return s.String()
}

type QueryCsrServicePaths struct {
  // {"en":"CSR ID", "zh_CN":"CSR ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryCsrServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServicePaths) GoString() string {
  return s.String()
}

func (s *QueryCsrServicePaths) SetId(v string) *QueryCsrServicePaths {
  s.Id = &v
  return s
}

type QueryCsrServiceParameters struct {
}

func (s QueryCsrServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceParameters) GoString() string {
  return s.String()
}

type QueryCsrServiceRequestHeader struct {
}

func (s QueryCsrServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryCsrServiceResponseHeader struct {
}

func (s QueryCsrServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceResponseHeader) GoString() string {
  return s.String()
}




type CreateTheCsrRequest struct {
  // {"en":"Enter a unique name representing your CSR. This is shown in the SSL List in the portal but is not shared with the certificate authority. This field is required.", "zh_CN":"csr名称，不能重复"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Key algorithm: RSA/EC
  // Choose RSA if you will use the RSA algorithm for your certificate. Choose EC if you will be using Elliptic Curve Cryptography instead.", "zh_CN":"密钥算法：RSA/EC 二选一"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"The main domain name", "zh_CN":"主域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"If you wish to specify subject alternative names for your certificate.Wildcards are permitted. Specify '*' in front of the domain name. For example, *.domain.com.You must enter like this
  // sans': [
  // 		'aaa.com',
  // 		'bbb.com'
  // 	]", "zh_CN":"备份域名列表"}
  Sans []*string `json:"sans,omitempty" xml:"sans,omitempty" type:"Repeated"`
  // {"en":"Choose the country where your organization is located.", "zh_CN":"国家地区"}
  Country *string `json:"country,omitempty" xml:"country,omitempty"`
  // {"en":"Enter the state/region where your organization is located. This shouldn't be abbreviated.", "zh_CN":"州"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Enter the city where your organization is located.", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty"`
  // {"en":"Enter the legal name for your company. Do not abbreviate.", "zh_CN":"公司"}
  Company *string `json:"company,omitempty" xml:"company,omitempty"`
  // {"en":"department", "zh_CN":"部门"}
  Department *string `json:"department,omitempty" xml:"department,omitempty"`
  // {"en":"Enter a department name if appropriate.", "zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s CreateTheCsrRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrRequest) GoString() string {
  return s.String()
}

func (s *CreateTheCsrRequest) SetName(v string) *CreateTheCsrRequest {
  s.Name = &v
  return s
}

func (s *CreateTheCsrRequest) SetAlgorithm(v string) *CreateTheCsrRequest {
  s.Algorithm = &v
  return s
}

func (s *CreateTheCsrRequest) SetDomain(v string) *CreateTheCsrRequest {
  s.Domain = &v
  return s
}

func (s *CreateTheCsrRequest) SetSans(v []*string) *CreateTheCsrRequest {
  s.Sans = v
  return s
}

func (s *CreateTheCsrRequest) SetCountry(v string) *CreateTheCsrRequest {
  s.Country = &v
  return s
}

func (s *CreateTheCsrRequest) SetState(v string) *CreateTheCsrRequest {
  s.State = &v
  return s
}

func (s *CreateTheCsrRequest) SetCity(v string) *CreateTheCsrRequest {
  s.City = &v
  return s
}

func (s *CreateTheCsrRequest) SetCompany(v string) *CreateTheCsrRequest {
  s.Company = &v
  return s
}

func (s *CreateTheCsrRequest) SetDepartment(v string) *CreateTheCsrRequest {
  s.Department = &v
  return s
}

func (s *CreateTheCsrRequest) SetComment(v string) *CreateTheCsrRequest {
  s.Comment = &v
  return s
}

type CreateTheCsrResponse struct {
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Request result data", "zh_CN":"请求结果数据"}
  Data *CreateTheCsrCreateTheCsrResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Added CSR ID after success, child node of data node", "zh_CN":"新增成功后的csr id ，&ldquo;data&rdquo;节点的子节点"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateTheCsrResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrResponse) GoString() string {
  return s.String()
}

func (s *CreateTheCsrResponse) SetCode(v string) *CreateTheCsrResponse {
  s.Code = &v
  return s
}

func (s *CreateTheCsrResponse) SetMessage(v string) *CreateTheCsrResponse {
  s.Message = &v
  return s
}

func (s *CreateTheCsrResponse) SetData(v *CreateTheCsrCreateTheCsrResponseData) *CreateTheCsrResponse {
  s.Data = v
  return s
}

func (s *CreateTheCsrResponse) SetId(v string) *CreateTheCsrResponse {
  s.Id = &v
  return s
}

type CreateTheCsrCreateTheCsrResponseData struct {
}

func (s CreateTheCsrCreateTheCsrResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrCreateTheCsrResponseData) GoString() string {
  return s.String()
}

type CreateTheCsrPaths struct {
}

func (s CreateTheCsrPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrPaths) GoString() string {
  return s.String()
}

type CreateTheCsrParameters struct {
}

func (s CreateTheCsrParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrParameters) GoString() string {
  return s.String()
}

type CreateTheCsrRequestHeader struct {
}

func (s CreateTheCsrRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrRequestHeader) GoString() string {
  return s.String()
}

type CreateTheCsrResponseHeader struct {
}

func (s CreateTheCsrResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrResponseHeader) GoString() string {
  return s.String()
}




type QueryCsrListRequest struct {
}

func (s QueryCsrListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListRequest) GoString() string {
  return s.String()
}

type QueryCsrListResponse struct {
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Request result data", "zh_CN":"请求结果数据"}
  CsrRecords []*QueryCsrListQueryCsrListResponseCsrRecords `json:"csr-records,omitempty" xml:"csr-records,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCsrListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListResponse) GoString() string {
  return s.String()
}

func (s *QueryCsrListResponse) SetCode(v string) *QueryCsrListResponse {
  s.Code = &v
  return s
}

func (s *QueryCsrListResponse) SetMessage(v string) *QueryCsrListResponse {
  s.Message = &v
  return s
}

func (s *QueryCsrListResponse) SetCsrRecords(v []*QueryCsrListQueryCsrListResponseCsrRecords) *QueryCsrListResponse {
  s.CsrRecords = v
  return s
}

type QueryCsrListQueryCsrListResponseCsrRecords struct     {
  // {"en":"CSR ID", "zh_CN":"CSR ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"CSR name", "zh_CN":"csr名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Key algorithm", "zh_CN":"密钥算法"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"The main domain name", "zh_CN":"主域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"comment", "zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
}

func (s QueryCsrListQueryCsrListResponseCsrRecords) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListQueryCsrListResponseCsrRecords) GoString() string {
  return s.String()
}

func (s *QueryCsrListQueryCsrListResponseCsrRecords) SetId(v string) *QueryCsrListQueryCsrListResponseCsrRecords {
  s.Id = &v
  return s
}

func (s *QueryCsrListQueryCsrListResponseCsrRecords) SetName(v string) *QueryCsrListQueryCsrListResponseCsrRecords {
  s.Name = &v
  return s
}

func (s *QueryCsrListQueryCsrListResponseCsrRecords) SetAlgorithm(v string) *QueryCsrListQueryCsrListResponseCsrRecords {
  s.Algorithm = &v
  return s
}

func (s *QueryCsrListQueryCsrListResponseCsrRecords) SetDomain(v string) *QueryCsrListQueryCsrListResponseCsrRecords {
  s.Domain = &v
  return s
}

func (s *QueryCsrListQueryCsrListResponseCsrRecords) SetComment(v string) *QueryCsrListQueryCsrListResponseCsrRecords {
  s.Comment = &v
  return s
}

type QueryCsrListPaths struct {
}

func (s QueryCsrListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListPaths) GoString() string {
  return s.String()
}

type QueryCsrListParameters struct {
}

func (s QueryCsrListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListParameters) GoString() string {
  return s.String()
}

type QueryCsrListRequestHeader struct {
}

func (s QueryCsrListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListRequestHeader) GoString() string {
  return s.String()
}

type QueryCsrListResponseHeader struct {
}

func (s QueryCsrListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListResponseHeader) GoString() string {
  return s.String()
}




type DeleteCsrRecordRequest struct {
}

func (s DeleteCsrRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCsrRecordRequest) GoString() string {
  return s.String()
}

type DeleteCsrRecordResponse struct {
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Request result data", "zh_CN":"请求结果数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteCsrRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCsrRecordResponse) GoString() string {
  return s.String()
}

func (s *DeleteCsrRecordResponse) SetCode(v string) *DeleteCsrRecordResponse {
  s.Code = &v
  return s
}

func (s *DeleteCsrRecordResponse) SetMessage(v string) *DeleteCsrRecordResponse {
  s.Message = &v
  return s
}

func (s *DeleteCsrRecordResponse) SetData(v string) *DeleteCsrRecordResponse {
  s.Data = &v
  return s
}

type DeleteCsrRecordPaths struct {
  // {"en":"Request id", "zh_CN":"请求id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteCsrRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCsrRecordPaths) GoString() string {
  return s.String()
}

func (s *DeleteCsrRecordPaths) SetId(v string) *DeleteCsrRecordPaths {
  s.Id = &v
  return s
}

type DeleteCsrRecordParameters struct {
}

func (s DeleteCsrRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCsrRecordParameters) GoString() string {
  return s.String()
}

type DeleteCsrRecordRequestHeader struct {
}

func (s DeleteCsrRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCsrRecordRequestHeader) GoString() string {
  return s.String()
}

type DeleteCsrRecordResponseHeader struct {
}

func (s DeleteCsrRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCsrRecordResponseHeader) GoString() string {
  return s.String()
}




