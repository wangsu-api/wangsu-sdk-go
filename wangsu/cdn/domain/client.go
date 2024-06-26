package domain

import (
	"errors"
	"github.com/wangsu/wangsu-sdk-go/common"
	"github.com/wangsu/wangsu-sdk-go/common/auth"
	"log"
)

type Client struct {
	common.Client
}

func NewClient(credential common.CredentialIface) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential)
	return
}

func (c *Client) AddCdnDomain(req *CreateDomainRequest) (requestId string, response *CreateDomainResponse, err error) {
	if req == nil {
		req = &CreateDomainRequest{}
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp CreateDomainResponse
	config := auth.NewAkskConfig(c.GetCredential(), "/api/domain", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetCdnDomainStatus(domainName string) (response *GetBasicConfigurationOfDomainResponse, err error) {
	req := &GetBasicConfigurationOfDomainRequest{}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}
	//invoke auth.Invoke
	var resp GetBasicConfigurationOfDomainResponse
	var requestId string
	config := auth.NewAkskConfig(c.GetCredential(), "/api/domain"+"/"+domainName, "GET")

	requestId, err = auth.Invoke(config, req, &resp)
	if err != nil {
		return nil, err
	}
	log.Printf("requestId: %s", requestId)
	return &resp, nil
}

func (c *Client) DeleteCdnDomain(domainName string, domainId int) (requestId string, response *DeleteApiDomainServiceResponse, err error) {
	req := &DeleteApiDomainServiceRequest{}
	req.SetDomainName(domainName)
	req.SetDomainId(string(rune(domainId)))

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	config := auth.NewAkskConfig(c.GetCredential(), "/api/domain"+"/"+domainName, "DELETE")
	requestId, err = auth.Invoke(config, req, &response)

	if err != nil {
		return "", nil, err
	}

	return requestId, response, nil
}

func (c *Client) QueryDomainDeployStatus(xCncRequestId string) (response *QueryApiDeployServiceResponse, err error) {
	req := &QueryApiDeployServiceRequest{}
	req.SetCncCequestId(xCncRequestId)

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}
	var resp QueryApiDeployServiceResponse
	config := auth.NewAkskConfig(c.GetCredential(), "/api/request/"+xCncRequestId, "GET")

	auth.Invoke(config, req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) EditDomainConfig(request *EditDomainConfigRequest, domainName string) (requestId string, response *EditDomainConfigResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp EditDomainConfigResponse
	config := auth.NewAkskConfig(c.GetCredential(), "/api/domain/"+domainName, "PUT")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetCdnDomainList(request *GetFuzzyPagingDomainListRequest) (requestId string, response *GetFuzzyPagingDomainListResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp GetFuzzyPagingDomainListResponse
	config := auth.NewAkskConfig(c.GetCredential(), "/api/domain/domainList", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
