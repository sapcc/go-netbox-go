package common

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    url.URL
	AuthToken  string
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

func (c *Client) SetAuthToken(header *http.Header) {
	header.Add("Authorization", "Token "+c.AuthToken)
	header.Set("Content-Type", "application/json")
}
