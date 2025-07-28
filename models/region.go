// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

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
