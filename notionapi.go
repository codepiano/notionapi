package notionapi

import (
	"fmt"
	"net/http"
)

const (
	ApiURL        = "https://api.notion.com"
	ApiVersion    = "v1"
	NotionVersion = "2021-05-13"
)

type ApiClient interface {
}

type Client struct {
	httpClient *http.Client

	Token IntegrationToken
}

func NewClient(httpClient *http.Client, token IntegrationToken) *Client {
	return &Client{
		httpClient: httpClient,
		Token:      token,
	}
}

type IntegrationToken string

func (it IntegrationToken) String() string {
	return string(it)
}

type Color string

func (c Color) String() string {
	return string(c)
}

func (c *Client) addRequestHeaders(req *http.Request) *http.Request {
	req.Header.Add("application/json", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token.String()))
	req.Header.Add("Notion-Version", NotionVersion)

	return req
}
