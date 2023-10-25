package common

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	HttpClient *http.Client
	BaseUrl    url.URL
	AuthToken  string
}

type ListParams struct {
	Id     int
	Name   string
	Q      string
	Limit  int
	OffSet int
}

type ReturnValues struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func (p *ListParams) SetListParams(values *url.Values) {
	if p.Id != 0 {
		values.Set("id", strconv.Itoa(p.Id))
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
	values.Set("exclude", "config_context")
}

func (c *Client) SetAuthToken(header *http.Header) {
	header.Add("Authorization", "Token "+c.AuthToken)
	header.Set("Content-Type", "application/json")
}
