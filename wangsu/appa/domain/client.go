package appadomain

import (
	"errors"
	"github.com/wangsu-api/wangsu-sdk-go/common"
	"github.com/wangsu-api/wangsu-sdk-go/common/auth"
)

type Client struct {
	common.Client
}

func NewClient(credential common.CredentialIface, httpProfile common.HttpProfileIface) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential)
	client.WithHttpProfile(httpProfile)
	return
}

func (c *Client) AddAppaDomain(req *AddAppaDomainForTerraformRequest) (requestId string, response *AddAppaDomainForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp AddAppaDomainForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/appa", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) UpdateAppaDomain(req *UpdateAppaDomainForTerraformRequest, domainName string) (requestId string, response *UpdateAppaDomainForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateAppaDomainForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/appa/"+domainName, "PUT")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryAppaDomain(domainName string) (requestId string, response *QueryAppaDomainForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	//invoke auth.Invoke
	var resp QueryAppaDomainForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/appa/"+domainName, "GET")

	requestId, err = auth.Invoke(config, nil, &resp)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, &resp, nil
}
