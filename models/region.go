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

type NestedRegion struct {
	Depth     int        `json:"_depth,omitempty"`
	Display   string     `json:"display,omitempty"`
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name"`
	SiteCount int        `json:"site_count,omitempty"`
	Slug      string     `json:"slug"`
	URL       strfmt.URI `json:"url,omitempty"`
}

type Region struct {
	Depth        int             `json:"_depth,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Description  string          `json:"description,omitempty"`
	Display      string          `json:"display,omitempty"`
	ID           int             `json:"id,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	Name         string          `json:"name"`
	Parent       NestedRegion    `json:"parent,omitempty"`
	SiteCount    int             `json:"site_count,omitempty"`
	Slug         string          `json:"slug"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	URL          strfmt.URI      `json:"url,omitempty"`
}

type ListRegionsRequest struct {
	common.ListParams
	Region string
	Slug   string
}

type ListRegionsResponse struct {
	common.ReturnValues
	Results []Region `json:"results"`
}
