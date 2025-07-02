package propertyconfig

import (
	"errors"
	"github.com/google/go-querystring/query"
	common2 "github.com/wangsu-api/wangsu-sdk-go/wangsu/common"
	"github.com/wangsu-api/wangsu-sdk-go/wangsu/common/auth"
	"log"
	"strconv"
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

func (c *Client) CreateProperty(req *CreatePropertyForTerraformRequest) (response *CreatePropertyForTerraformResponse, err error) {
	if req == nil {
		return nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp CreatePropertyForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties", "POST")
	var requestId string
	requestId, err = auth.Invoke(config, req, &resp)

	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) UpdateProperty(propertyId int, req *UpdatePropertyForTerraformRequest) (response *UpdatePropertyForTerraformResponse, err error) {
	if req == nil {
		return nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}
	var propertyIdStr = strconv.Itoa(propertyId)
	var resp UpdatePropertyForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties/"+propertyIdStr, "PUT")
	var requestId string
	requestId, err = auth.Invoke(config, req, &resp)

	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) QueryProperty(propertyId int) (response *QueryPropertyConfigForTerrformResponse, err error) {
	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp QueryPropertyConfigForTerrformResponse
	var propertyIdStr = strconv.Itoa(propertyId)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties/"+propertyIdStr, "GET")
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)
	if err != nil {
		return nil, err
	}
	log.Printf("requestId: %s", requestId)
	return &resp, nil
}

func (c *Client) QueryPropertyVersion(propertyId int, version int) (response *QueryPropertyVersionConfigForTerrformResponse, err error) {
	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp QueryPropertyVersionConfigForTerrformResponse
	var propertyIdStr = strconv.Itoa(propertyId)
	var versionStr = strconv.Itoa(version)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties/"+propertyIdStr+"/versions/"+versionStr, "GET")
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)

	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) QueryProperties(parameters *QueryPropertiesForTerraformParameters) (requestId string, response *QueryPropertiesForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryPropertiesForTerraformResponse
	var uri = "/api/terraform/properties"
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

func (c *Client) DeleteProperty(propertyId int) (response *DeletePropertyForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var propertyIdStr = strconv.Itoa(propertyId)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties/"+propertyIdStr, "DELETE")
	var resp DeletePropertyForTerraformResponse
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)

	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CreateDeployment(req *CreateDeploymentTaskForTerraformRequest) (response *CreateDeploymentTaskForTerraformResponse, err error) {
	if req == nil {
		return nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp CreateDeploymentTaskForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties/deployments", "POST")
	var requestId string
	requestId, err = auth.Invoke(config, req, &resp)

	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) QueryDeployment(deploymentId int) (response *QueryDeploymentForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}

	var resp QueryDeploymentForTerraformResponse
	var deploymentIdStr = strconv.Itoa(deploymentId)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/properties/deployments/"+deploymentIdStr, "GET")
	var requestId string
	requestId, err = auth.Invoke(config, nil, &resp)

	log.Printf("requestId: %s", requestId)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) QueryDeployments(parameters *QueryDeploymentsForTerraformParameters) (requestId string, response *QueryDeploymentsForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryDeploymentsForTerraformResponse
	var uri = "/api/terraform/properties/deployments"
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
