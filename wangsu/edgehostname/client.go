package edgehostname

import (
	"errors"
	"github.com/google/go-querystring/query"
	common2 "github.com/wangsu-api/wangsu-sdk-go/wangsu/common"
	"github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
	"log"
)

type Client struct {
	common2.Client
}

func NewClient(credential common2.CredentialIface, httpProfile common2.HttpProfileIface) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential)
	client.WithHttpProfile(httpProfile)
	return
}

func (c *Client) UpdateEdgeHostname(edgeHostname string, req *UpdateEdgeHostnameForTerraformRequest) (response *UpdateEdgeHostnameForTerraformResponse, err error) {
	if edgeHostname == "" {
		return nil, errors.New("edgeHostname is required")
	}

	if req == nil {
		return nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp UpdateEdgeHostnameForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/edge-hostnames/"+edgeHostname, "PUT")
	var requestId string
	requestId, err = auth.Invoke(config, req, &resp)
	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteEdgeHostname(edgeHostname string) (response *DeleteEdgeHostnameForTerraformResponse, err error) {
	if edgeHostname == "" {
		return nil, errors.New("edgeHostname is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp DeleteEdgeHostnameForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/edge-hostnames/"+edgeHostname, "DELETE")
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)
	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeployEdgeHostname(edgeHostname string) (response *DeployEdgeHostnameForTerraformResponse, err error) {
	if edgeHostname == "" {
		return nil, errors.New("edgeHostname is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp DeployEdgeHostnameForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/edge-hostnames/"+edgeHostname+"/deploy", "POST")
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)
	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) QueryEdgeHostname(edgeHostname string) (response *QueryEdgeHostnameForTerraformResponse, err error) {
	if edgeHostname == "" {
		return nil, errors.New("edgeHostname is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp QueryEdgeHostnameForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/edge-hostnames/"+edgeHostname, "GET")
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)
	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) QueryEdgeHostnames(parameters *QueryEdgeHostnamesForTerraformParameters) (requestId string, response *QueryEdgeHostnamesForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryEdgeHostnamesForTerraformResponse
	var uri = "/api/terraform/edge-hostnames"
	if parameters != nil {
		values, err := query.Values(parameters)
		if err != nil {
			return "", nil, errors.New("failed to encode query parameters: " + err.Error())
		}
		uri = uri + "?" + values.Encode()
	}
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), uri, "GET")

	requestId, err = auth.Invoke(config, nil, &resp)
	log.Printf("requestId: %s", requestId)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, &resp, nil
}
