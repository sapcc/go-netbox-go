// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strconv"
)

type HTTPConnectable interface {
	HTTPClient() *http.Client
	SetHTTPClient(httpClient *http.Client)

	BaseURL() *url.URL
	SetBaseURL(url *url.URL)

	AuthToken() string
	SetAuthToken(authToken string)
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	authToken  string
}

func Initialize(httpConnect HTTPConnectable, baseURL, authToken string, insecureSkipVerify bool) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	httpConnect.SetBaseURL(u)
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify, // #nosec
	}
	httpConnect.SetHTTPClient(&http.Client{Transport: tr})
	httpConnect.SetAuthToken(authToken)
	return nil
}

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) BaseURL() *url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(url *url.URL) {
	c.baseURL = url
}

func (c *Client) AuthToken() string {
	return c.authToken
}

func (c *Client) SetAuthToken(authToken string) {
	c.authToken = authToken
}

func (c *Client) ApplyAuthTokenToHeader(header *http.Header) {
	header.Add("Authorization", "Token "+c.authToken)
	header.Set("Content-Type", "application/json")
}

type ListParams struct {
	ID                   int
	Name                 string
	Q                    string
	Limit                int
	OffSet               int
	ExcludeConfigContext bool
}

type ReturnValues struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func (p *ListParams) SetListParams(values *url.Values) {
	if p.ID != 0 {
		values.Set("id", strconv.Itoa(p.ID))
	}
	if p.Name != "" {
		values.Set("name", p.Name)
	}
	if p.Q != "" {
		values.Set("q", p.Q)
	}
	if p.Limit != 0 {
		values.Set("limit", strconv.Itoa(p.Limit))
	}
	if p.OffSet != 0 {
		values.Set("offset", strconv.Itoa(p.OffSet))
	}
	if p.ExcludeConfigContext {
		values.Set("exclude", "config_context")
	}
}
