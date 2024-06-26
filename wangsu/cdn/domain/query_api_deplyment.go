// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryApiDeployServiceRequest struct {
	// {"en":"or each account request record, a unique cnc-request-id (for all API) is generated through which the execution results of each asynchronous request task can be queried", "zh_CN":"对于账号每一次请求记录，都会生成唯一的cnc-request-id（适用全部接口），通过该id可以查询每次异步请求任务的执行结果"}
	CncCequestId *string `json:"cncCequestId,omitempty" xml:"cncCequestId,omitempty" require:"true"`
}

func (s QueryApiDeployServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApiDeployServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryApiDeployServiceRequest) SetCncCequestId(v string) *QueryApiDeployServiceRequest {
	s.CncCequestId = &v
	return s
}

type QueryApiDeployServiceResponse struct {
	// {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
	HttpStatus *int `json:"http-status-code,omitempty" xml:"http-status-code,omitempty" require:"true"`
	// {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
	XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
	// {"en":"The cnc-request-id corresponding to the request record you want to query this time", "zh_CN":"本次想要查询的请求记录对应的cnc-request-id"}
	CncRequestId *string `json:"cnc-request-id,omitempty" xml:"cnc-request-id,omitempty" require:"true"`
	// {"en":"The submission time for the request record that you want to query, for example: Thu, 09 Nov 2017 22:37:53 CST", "zh_CN":"本次想要查询的请求记录的提交时间，例如：Thu,09 Nov 2017 22:37:53 CST"}
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
	// {"en":"Task execution results for asynchronous requests, including the following four states: WAIT: Indicates that the request is pending execution INPROGRESS: indicates that the request is in execution SUCCESS: Indicates that the request has been successfully executed FAIL: Indicates that the request execution failed Note: 1. If the query result is a failure of execution, there will be a mechanism or manual intervention in the backend of the system until the deployment is successful.If the result is the asynchronous request task of the new domain name, the system will have a resubmission mechanism or human intervention. 2. Modify the asynchronous request task for the domain name configuration, as well as add the domain name, there will be the mechanism of reintroduction and human intervention.", "zh_CN":"异步请求的任务执行结果，包括以下4种状态： WAIT：表示该请求等待执行 INPROGRESS：表示该请求执行中 SUCCESS：表示该请求已经执行成功 FAIL：表示该请求执行失败 注意： 1.新建域名的异步请求任务，通过本接口如果查询到结果是执行失败，系统后台会会有重提机制或者人工干预，直到部署成功。 2.修改域名配置的异步请求任务，和新增域名一样，也会有重提机制和人工干预。"}
	AsyncResult *string `json:"async-result,omitempty" xml:"async-result,omitempty" require:"true"`
	// {"en":"More information about task execution results for asynchronous requests", "zh_CN":"异步请求的任务执行结果的更多相关信息"}
	AsyncMessage *string `json:"async-message,omitempty" xml:"async-message,omitempty" require:"true"`
}

func (s QueryApiDeployServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApiDeployServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryApiDeployServiceResponse) SetHttpStatus(v int) *QueryApiDeployServiceResponse {
	s.HttpStatus = &v
	return s
}

func (s *QueryApiDeployServiceResponse) SetXCncRequestId(v string) *QueryApiDeployServiceResponse {
	s.XCncRequestId = &v
	return s
}

func (s *QueryApiDeployServiceResponse) SetCncRequestId(v string) *QueryApiDeployServiceResponse {
	s.CncRequestId = &v
	return s
}

func (s *QueryApiDeployServiceResponse) SetTimestamp(v string) *QueryApiDeployServiceResponse {
	s.Timestamp = &v
	return s
}

func (s *QueryApiDeployServiceResponse) SetAsyncResult(v string) *QueryApiDeployServiceResponse {
	s.AsyncResult = &v
	return s
}

func (s *QueryApiDeployServiceResponse) SetAsyncMessage(v string) *QueryApiDeployServiceResponse {
	s.AsyncMessage = &v
	return s
}
