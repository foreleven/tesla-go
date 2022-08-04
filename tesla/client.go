package tesla

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

var DEFAULT_SSO_URL = ""
var DEFAULT_API_URL = "https://owner-api.teslamotors.com"

var DEFAULT_CHINA_SSO_URL = ""
var DEFAULT_CHINA_API_URL = "https://owner-api.vn.cloud.tesla.cn"

type AuthToken struct {
	AccessToken string
}

type Client struct {
	Token  *AuthToken
	SsoURL string ``
	ApiURL string
}

type Region string

const (
	China Region = "China"
	Other        = "Other"
)

type ClientOptions struct {
	AuthToken
	Region
}

var http_client *http.Client = &http.Client{}

func NewClient(options ClientOptions) (*Client, error) {
	var ssoURL string
	var apiURL string
	if options.Region == China {
		ssoURL = DEFAULT_CHINA_SSO_URL
		apiURL = DEFAULT_CHINA_API_URL
	} else {
		ssoURL = DEFAULT_SSO_URL
		apiURL = DEFAULT_API_URL
	}
	client := &Client{
		Token: &AuthToken{
			AccessToken: options.AccessToken,
		},
		SsoURL: ssoURL,
		ApiURL: apiURL,
	}

	return client, nil
}

// Calls an GET http request
func (c Client) get(path string) ([]byte, error) {
	req, _ := http.NewRequest("GET", c.ApiURL+path, nil)
	return c.request(req)
}

func (c Client) post(path string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(body))
	return c.request(req)
}

func (c Client) request(req *http.Request) ([]byte, error) {
	c.setHeaders(req)
	res, err := http_client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c Client) setHeaders(req *http.Request) {
	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}
