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

package models

import (
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/go-netbox-go/common"
)

type NestedSiteGroup struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type SiteGroup struct {
	Created      strfmt.DateTime `json:"created"`
	CustomFields interface{}     `json:"custom_fields"`
	Depth        int             `json:"_depth"`
	Description  string          `json:"description"`
	Display      string          `json:"display,omitempty"`
	ID           int             `json:"id"`
	LastUpdated  strfmt.DateTime `json:"last_updated"`
	Name         string          `json:"name"`
	Parent       NestedSiteGroup `json:"parent,omitempty"`
	SiteCount    int             `json:"site_count"`
	Slug         string          `json:"slug"`
	Tags         []NestedTag     `json:"tags"`
	URL          strfmt.URI      `json:"url"`
}

type ListSiteGroupsRequest struct {
	common.ListParams
	Region string
}

type ListSiteGroupsResponse struct {
	common.ReturnValues
	Results []SiteGroup `json:"results"`
}
