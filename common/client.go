package common

import "net/http"

type Client struct {
	httpClient    *http.Client
	credential    CredentialIface
	signMethod    string
	debug         bool
	requestClient string
}

func (c *Client) GetCredential() CredentialIface {
	return c.credential
}

func (c *Client) WithCredential(cred CredentialIface) *Client {
	c.credential = cred
	return c
}
