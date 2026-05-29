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

type QueryCsrServiceRequestHeader struct {
}

func (s QueryCsrServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryCsrServicePaths struct {
  // {"en":"CSR ID","zh_CN":"CSR ID"}
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

type QueryCsrServiceResponse struct {
  // {"en":"Request result code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Request result data","zh_CN":"请求结果数据"}
  Data *QueryCsrServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"CSR ID","zh_CN":"CSR ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"CSR name, which cannot be repeated","zh_CN":"csr名称，不能重复"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Key algorithm","zh_CN":"密钥算法","exampleValue":"RSA,EC"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"CSR content","zh_CN":"CSR内容"}
  CsrFile *string `json:"csr-file,omitempty" xml:"csr-file,omitempty" require:"true"`
  // {"en":"The main domain name","zh_CN":"主域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Backup domain name list","zh_CN":"备份域名列表"}
  Sans []*string `json:"sans,omitempty" xml:"sans,omitempty" require:"true" type:"Repeated"`
  // {"en":"Comment","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Country or region","zh_CN":"国家地区"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"State","zh_CN":"州"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"City","zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"Company","zh_CN":"公司"}
  Company *string `json:"company,omitempty" xml:"company,omitempty" require:"true"`
  // {"en":"Department","zh_CN":"部门"}
  Department *string `json:"department,omitempty" xml:"department,omitempty" require:"true"`
  // {"en":"Key Length.\nRSA: 2048 | 3072 | 4096 (Default: 2048)\nECC: 256 | 384 | 521 (Default: 256)\nSM2: 256 (Default: 256)","zh_CN":"秘钥强度，RSA：2048\3072\4096， 默认2048 EC：256\384\521， 默认256 SM2：256， 默认256"}
  KeyLength *int `json:"keyLength,omitempty" xml:"keyLength,omitempty" require:"true"`
  // {"en":"Street","zh_CN":"街道1"}
  Street *string `json:"street,omitempty" xml:"street,omitempty" require:"true"`
  // {"en":"Street1","zh_CN":"街道2"}
  Street1 *string `json:"street1,omitempty" xml:"street1,omitempty" require:"true"`
  // {"en":"Postal code","zh_CN":"邮编"}
  PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty" require:"true"`
  // {"en":"Organization contact phone number","zh_CN":"组织联系电话"}
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
  // {"en":"Modification time, format: yyyy-MM-dd HH:mm:ss","zh_CN":"修改时间，时间格式：yyyy-MM-dd HH:mm:ss"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  // {"en":"Creation time, format: yyyy-MM-dd HH:mm:ss","zh_CN":"创建时间，时间格式：yyyy-MM-dd HH:mm:ss"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
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

func (s *QueryCsrServiceResponse) SetData(v *QueryCsrServiceResponseData) *QueryCsrServiceResponse {
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

func (s *QueryCsrServiceResponse) SetKeyLength(v int) *QueryCsrServiceResponse {
  s.KeyLength = &v
  return s
}

func (s *QueryCsrServiceResponse) SetStreet(v string) *QueryCsrServiceResponse {
  s.Street = &v
  return s
}

func (s *QueryCsrServiceResponse) SetStreet1(v string) *QueryCsrServiceResponse {
  s.Street1 = &v
  return s
}

func (s *QueryCsrServiceResponse) SetPostalCode(v string) *QueryCsrServiceResponse {
  s.PostalCode = &v
  return s
}

func (s *QueryCsrServiceResponse) SetPhone(v string) *QueryCsrServiceResponse {
  s.Phone = &v
  return s
}

func (s *QueryCsrServiceResponse) SetModifyTime(v string) *QueryCsrServiceResponse {
  s.ModifyTime = &v
  return s
}

func (s *QueryCsrServiceResponse) SetCreateTime(v string) *QueryCsrServiceResponse {
  s.CreateTime = &v
  return s
}

type QueryCsrServiceResponseData struct {
}

func (s QueryCsrServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrServiceResponseData) GoString() string {
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
  // {"en":"Enter a unique name representing your CSR. This is shown in the SSL List in the portal but is not shared with the certificate authority. This field is required.","zh_CN":"csr名称，不能重复"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Key algorithm.\nChoose RSA if you will use the RSA algorithm for your certificate. Choose EC if you will be using Elliptic Curve Cryptography instead.","zh_CN":"密钥算法","exampleValue":"RSA,EC"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"Common name","zh_CN":"主域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Subject Alternative Names (SANs). \nSupports wildcards (e.g., *.example.com). Underscores (_) are not allowed in domain names.","zh_CN":"证书备用名称（SAN）。支持泛域名，以“*”开头（例如：*.example.com）；域名不能包含“_”。"}
  Sans []*string `json:"sans,omitempty" xml:"sans,omitempty" type:"Repeated"`
  // {"en":"Choose the country where your organization is located.","zh_CN":"国家地区"}
  Country *string `json:"country,omitempty" xml:"country,omitempty"`
  // {"en":"Enter the state/region where your organization is located. This shouldn't be abbreviated.","zh_CN":"州"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Enter the city where your organization is located.","zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty"`
  // {"en":"Enter the legal name for your company. Do not abbreviate.","zh_CN":"公司"}
  Company *string `json:"company,omitempty" xml:"company,omitempty"`
  // {"en":"department","zh_CN":"部门"}
  Department *string `json:"department,omitempty" xml:"department,omitempty"`
  // {"en":"Enter a department name if appropriate.","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Key Length.\nRSA: 2048 | 3072 | 4096 (Default: 2048)\nECC: 256 | 384 | 521 (Default: 256)\nSM2: 256 (Default: 256)","zh_CN":"秘钥强度，RSA：2048\3072\4096， 默认2048 EC：256\384\521， 默认256 SM2：256， 默认256"}
  KeyLength *int `json:"keyLength,omitempty" xml:"keyLength,omitempty"`
  // {"en":"street","zh_CN":"街道1"}
  Street *string `json:"street,omitempty" xml:"street,omitempty"`
  // {"en":"street1","zh_CN":"街道2"}
  Street1 *string `json:"street1,omitempty" xml:"street1,omitempty"`
  // {"en":"postalCode","zh_CN":"邮编"}
  PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty"`
  // {"en":"phone","zh_CN":"组织联系电话"}
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
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

func (s *CreateTheCsrRequest) SetKeyLength(v int) *CreateTheCsrRequest {
  s.KeyLength = &v
  return s
}

func (s *CreateTheCsrRequest) SetStreet(v string) *CreateTheCsrRequest {
  s.Street = &v
  return s
}

func (s *CreateTheCsrRequest) SetStreet1(v string) *CreateTheCsrRequest {
  s.Street1 = &v
  return s
}

func (s *CreateTheCsrRequest) SetPostalCode(v string) *CreateTheCsrRequest {
  s.PostalCode = &v
  return s
}

func (s *CreateTheCsrRequest) SetPhone(v string) *CreateTheCsrRequest {
  s.Phone = &v
  return s
}

func (s *CreateTheCsrRequest) SetEmail(v string) *CreateTheCsrRequest {
  s.Email = &v
  return s
}

type CreateTheCsrRequestHeader struct {
}

func (s CreateTheCsrRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrRequestHeader) GoString() string {
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

type CreateTheCsrResponse struct {
  // {"en":"Request result code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Request result data","zh_CN":"请求结果数据"}
  Data *CreateTheCsrResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Added CSR ID after success, child node of data node","zh_CN":"新增成功后的csr id ，&ldquo;data&rdquo;节点的子节点"}
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

func (s *CreateTheCsrResponse) SetData(v *CreateTheCsrResponseData) *CreateTheCsrResponse {
  s.Data = v
  return s
}

func (s *CreateTheCsrResponse) SetId(v string) *CreateTheCsrResponse {
  s.Id = &v
  return s
}

type CreateTheCsrResponseData struct {
}

func (s CreateTheCsrResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateTheCsrResponseData) GoString() string {
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




type UpdateCsrRequest struct {
  // {"en":"The name of the CSR.","zh_CN":"CSR名称"}
  RecordName *string `json:"recordName,omitempty" xml:"recordName,omitempty" require:"true"`
  // {"en":"Additional comments for the CSR.","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s UpdateCsrRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCsrRequest) GoString() string {
  return s.String()
}

func (s *UpdateCsrRequest) SetRecordName(v string) *UpdateCsrRequest {
  s.RecordName = &v
  return s
}

func (s *UpdateCsrRequest) SetComment(v string) *UpdateCsrRequest {
  s.Comment = &v
  return s
}

type UpdateCsrRequestHeader struct {
}

func (s UpdateCsrRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCsrRequestHeader) GoString() string {
  return s.String()
}

type UpdateCsrPaths struct {
  // {"en":"The csr record id","zh_CN":"CSR记录ID"}
  CsrId *int `json:"csrId,omitempty" xml:"csrId,omitempty" require:"true"`
}

func (s UpdateCsrPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCsrPaths) GoString() string {
  return s.String()
}

func (s *UpdateCsrPaths) SetCsrId(v int) *UpdateCsrPaths {
  s.CsrId = &v
  return s
}

type UpdateCsrParameters struct {
}

func (s UpdateCsrParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCsrParameters) GoString() string {
  return s.String()
}

type UpdateCsrResponse struct {
  // {"en":"Request result code","zh_CN":"响应代码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateCsrResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCsrResponse) GoString() string {
  return s.String()
}

func (s *UpdateCsrResponse) SetCode(v string) *UpdateCsrResponse {
  s.Code = &v
  return s
}

func (s *UpdateCsrResponse) SetMessage(v string) *UpdateCsrResponse {
  s.Message = &v
  return s
}

type UpdateCsrResponseHeader struct {
}

func (s UpdateCsrResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCsrResponseHeader) GoString() string {
  return s.String()
}




type QueryCsrListRequest struct {
  // {"en":"Csr record name","zh_CN":"CSR名称"}
  RecordName *string `json:"recordName,omitempty" xml:"recordName,omitempty"`
  // {"en":"Domain name","zh_CN":"CSR授权域名"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Company Name","zh_CN":"组织名称"}
  Companies []*string `json:"companies,omitempty" xml:"companies,omitempty" type:"Repeated"`
  // {"en":"Page index","zh_CN":"分页页码"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Page size","zh_CN":"每页个数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Sort order, default: desc, range: asc,desc","zh_CN":"排序方式，默认值: desc，取值范围: asc,desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
}

func (s QueryCsrListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListRequest) GoString() string {
  return s.String()
}

func (s *QueryCsrListRequest) SetRecordName(v string) *QueryCsrListRequest {
  s.RecordName = &v
  return s
}

func (s *QueryCsrListRequest) SetDomains(v []*string) *QueryCsrListRequest {
  s.Domains = v
  return s
}

func (s *QueryCsrListRequest) SetCompanies(v []*string) *QueryCsrListRequest {
  s.Companies = v
  return s
}

func (s *QueryCsrListRequest) SetPageIndex(v int) *QueryCsrListRequest {
  s.PageIndex = &v
  return s
}

func (s *QueryCsrListRequest) SetPageSize(v int) *QueryCsrListRequest {
  s.PageSize = &v
  return s
}

func (s *QueryCsrListRequest) SetSortOrder(v string) *QueryCsrListRequest {
  s.SortOrder = &v
  return s
}

type QueryCsrListRequestHeader struct {
}

func (s QueryCsrListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListRequestHeader) GoString() string {
  return s.String()
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

type QueryCsrListResponse struct {
  // {"en":"Request result code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryCsrListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *QueryCsrListResponse) SetData(v *QueryCsrListResponseData) *QueryCsrListResponse {
  s.Data = v
  return s
}

type QueryCsrListResponseData struct {
  // {"en":"Page size","zh_CN":"分页页码"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"Page size","zh_CN":"每页个数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Record total","zh_CN":"记录总数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  CsrRecords []*QueryCsrListResponseDataCsrRecords `json:"csrRecords,omitempty" xml:"csrRecords,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCsrListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListResponseData) GoString() string {
  return s.String()
}

func (s *QueryCsrListResponseData) SetPageIndex(v int) *QueryCsrListResponseData {
  s.PageIndex = &v
  return s
}

func (s *QueryCsrListResponseData) SetPageSize(v int) *QueryCsrListResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryCsrListResponseData) SetTotal(v int) *QueryCsrListResponseData {
  s.Total = &v
  return s
}

func (s *QueryCsrListResponseData) SetCsrRecords(v []*QueryCsrListResponseDataCsrRecords) *QueryCsrListResponseData {
  s.CsrRecords = v
  return s
}

type QueryCsrListResponseDataCsrRecords struct     {
  // {"en":"CSR ID","zh_CN":"CSR ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"CSR name","zh_CN":"csr名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Key algorithm","zh_CN":"密钥算法"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"The main domain name","zh_CN":"主域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Comment","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Sans","zh_CN":"备用域名"}
  Sans []*string `json:"sans,omitempty" xml:"sans,omitempty" require:"true" type:"Repeated"`
  // {"en":"Modify time","zh_CN":"修改时间；时间格式：yyyy-MM-dd HH:mm:ss"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  // {"en":"Company","zh_CN":"组织名称"}
  Company *string `json:"company,omitempty" xml:"company,omitempty" require:"true"`
}

func (s QueryCsrListResponseDataCsrRecords) String() string {
  return tea.Prettify(s)
}

func (s QueryCsrListResponseDataCsrRecords) GoString() string {
  return s.String()
}

func (s *QueryCsrListResponseDataCsrRecords) SetId(v string) *QueryCsrListResponseDataCsrRecords {
  s.Id = &v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetName(v string) *QueryCsrListResponseDataCsrRecords {
  s.Name = &v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetAlgorithm(v string) *QueryCsrListResponseDataCsrRecords {
  s.Algorithm = &v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetDomain(v string) *QueryCsrListResponseDataCsrRecords {
  s.Domain = &v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetComment(v string) *QueryCsrListResponseDataCsrRecords {
  s.Comment = &v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetSans(v []*string) *QueryCsrListResponseDataCsrRecords {
  s.Sans = v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetModifyTime(v string) *QueryCsrListResponseDataCsrRecords {
  s.ModifyTime = &v
  return s
}

func (s *QueryCsrListResponseDataCsrRecords) SetCompany(v string) *QueryCsrListResponseDataCsrRecords {
  s.Company = &v
  return s
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




