/*
 *   Copyright 2020 SAP SE
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

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
