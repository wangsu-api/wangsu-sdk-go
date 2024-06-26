// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type GetFuzzyPagingDomainListRequest struct {
	// {"en":"Page number must be a positive integer greater than 0", "zh_CN":"分页的页码，必须为大于0的正整数"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
	// {"en":"Number of domain name data items for paging, must be a positive integer greater than 0", "zh_CN":"分页的域名数据条数，必须大于0的正整数"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	// {"en":"Specifies the service type of the query, only one type per query, and no default lookup for all types", "zh_CN":"指定查询的服务类型，每次查询只能传一个类型，不传默认查全部类型"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// {"en":"Specifies the accelerated domain name for the query, allows multiple domains, commas delimited, and no default lookup of all domain names", "zh_CN":"指定查询的加速域名，允许多个域名，逗号分隔，不传默认查全部域名"}
	DomainName []*string `json:"domainName,omitempty" xml:"domainName,omitempty" type:"Repeated"`
	// {"en":"Query to accelerated domain name, optional value is: fuzzy_match for fuzzy query; Full_match represents an exact query
	// No fuzzy_match by default, for accelerated domain name only", "zh_CN":"查询加速域名的方式，可选值为：fuzzy_match表示模糊查询；full_match表示精确查询
	// 不传默认为fuzzy_match，仅针对加速域名"}
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// {"en":"Query start time, support for years, months, days, hours, minutes, and seconds, for example: 20170101.09 million. Time equals", "zh_CN":"查询开始时间，支持范围为年月日时分秒，例如：20170101090000。时间含等于"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// {"en":"Query end time, query time within the existence of the accelerated domain name, time is equal to, do not pass the default query all", "zh_CN":"查询结束时间，查询时间段内存在的加速域名，时间含等于，不传默认查询所有"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// {"en":"Accelerate the status of the domain name, enabled indicates that it is in effect; Disabled indicates that it is Disabled; Deploying means in the process of deployment; Checking indicates that the audit is in progress; Disabling: Indicates disabled, no default lookup for all", "zh_CN":"加速域名的状态，enabled表示已生效；disabled表示已禁用；deploying表示部署中；checking表示审核中；disabling:表示禁用中，不传默认查全部"}
	DomainStatus *string `json:"domainStatus,omitempty" xml:"domainStatus,omitempty"`
}

func (s GetFuzzyPagingDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListRequest) GoString() string {
	return s.String()
}

func (s *GetFuzzyPagingDomainListRequest) SetPageNumber(v int) *GetFuzzyPagingDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetPageSize(v int) *GetFuzzyPagingDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetServiceType(v string) *GetFuzzyPagingDomainListRequest {
	s.ServiceType = &v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetDomainName(v []*string) *GetFuzzyPagingDomainListRequest {
	s.DomainName = v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetQueryType(v string) *GetFuzzyPagingDomainListRequest {
	s.QueryType = &v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetStartTime(v string) *GetFuzzyPagingDomainListRequest {
	s.StartTime = &v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetEndTime(v string) *GetFuzzyPagingDomainListRequest {
	s.EndTime = &v
	return s
}

func (s *GetFuzzyPagingDomainListRequest) SetDomainStatus(v string) *GetFuzzyPagingDomainListRequest {
	s.DomainStatus = &v
	return s
}

type GetFuzzyPagingDomainListResponse struct {
	// {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
	Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
	XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
	// {"en":"Responses the page number of the data", "zh_CN":"所有满足条件的数据总条数"}
	TotalCount *int `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
	// {"en":"total pages", "zh_CN":"总页数"}
	TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
	// {"en":"Responses the page number of the data", "zh_CN":"返回数据的页码"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
	// {"en":"Number of data page", "zh_CN":"每个页面的数据条数"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	// {"en":"Responses status information for the accelerated domain name", "zh_CN":"返回加速域名的状态信息"}
	ResultList []*GetFuzzyPagingDomainListResponseResultList `json:"resultList,omitempty" xml:"resultList,omitempty" require:"true" type:"Repeated"`
}

func (s GetFuzzyPagingDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponse) GoString() string {
	return s.String()
}

func (s *GetFuzzyPagingDomainListResponse) SetCode(v int) *GetFuzzyPagingDomainListResponse {
	s.Code = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponse) SetXCncRequestId(v string) *GetFuzzyPagingDomainListResponse {
	s.XCncRequestId = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponse) SetTotalCount(v int) *GetFuzzyPagingDomainListResponse {
	s.TotalCount = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponse) SetTotalPageNumber(v int) *GetFuzzyPagingDomainListResponse {
	s.TotalPageNumber = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponse) SetPageNumber(v int) *GetFuzzyPagingDomainListResponse {
	s.PageNumber = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponse) SetPageSize(v int) *GetFuzzyPagingDomainListResponse {
	s.PageSize = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponse) SetResultList(v []*GetFuzzyPagingDomainListResponseResultList) *GetFuzzyPagingDomainListResponse {
	s.ResultList = v
	return s
}

type GetFuzzyPagingDomainListResponseResultList struct {
	// {"en":"Accelerated domain CNAME corresponding to CNAME, for example: 7nt6mrh7sdkslj.cdn30.com", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
	// {"en":"Configuration name", "zh_CN":"配置单名称"}
	ConfigFormName *string `json:"configFormName,omitempty" xml:"configFormName,omitempty" require:"true"`
	// {"en":"The time format is: 20160323112310", "zh_CN":"时间格式为：20160323112310"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"Corresponding domain ID", "zh_CN":"对应的域名ID"}
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
	// {"en":"Operator of this query", "zh_CN":"本次查询的操作者"}
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
	// {"en":"Accelerate the origin IP of a domain name", "zh_CN":"加速域名的回源IP"}
	OriginIps *string `json:"originIps,omitempty" xml:"originIps,omitempty" require:"true"`
	// {"en":"Service type for accelerated domain name", "zh_CN":"加速域名的服务类型"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"Status of accelerated domain name.", "zh_CN":"加速域名的状态"}
	DomainStatus *string `json:"domainStatus,omitempty" xml:"domainStatus,omitempty" require:"true"`
	// {"en":"Deployment version code", "zh_CN":"部署版本号"}
	DeployVersion *string `json:"deployVersion,omitempty" xml:"deployVersion,omitempty" require:"true"`
	// {"en":"Does the domain name enable CDN acceleration services, Y and N?", "zh_CN":"域名是否启用CDN加速服务，Y和N"}
	CdnServiceStatus *string `json:"cdnServiceStatus,omitempty" xml:"cdnServiceStatus,omitempty" require:"true"`
	// {"en":"Whether the accelerated domain name is enabled, Y and N?", "zh_CN":"加速域名是否启用，Y和N"}
	IsEnabled *string `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s GetFuzzyPagingDomainListResponseResultList) String() string {
	return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponseResultList) GoString() string {
	return s.String()
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetCname(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.Cname = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetConfigFormName(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.ConfigFormName = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetCreateTime(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.CreateTime = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDomainId(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.DomainId = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDomainName(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.DomainName = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetOperator(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.Operator = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetOriginIps(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.OriginIps = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetServiceType(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.ServiceType = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDomainStatus(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.DomainStatus = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDeployVersion(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.DeployVersion = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetCdnServiceStatus(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.CdnServiceStatus = &v
	return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetIsEnabled(v string) *GetFuzzyPagingDomainListResponseResultList {
	s.IsEnabled = &v
	return s
}
