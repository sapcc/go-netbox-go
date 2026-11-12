// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import "github.com/go-openapi/strfmt"

type Location struct {
	Depth        int             `json:"_depth,omitempty"`
	Created      strfmt.DateTime `json:"created"`
	CustomFields any             `json:"custom_fields,omitempty"`
	Description  string          `json:"description,omitempty"`
	DeviceCount  int             `json:"device_count,omitempty"`
	Display      string          `json:"display,omitempty"`
	ID           int             `json:"id,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated"`
	Name         string          `json:"name"`
	Parent       NestedLocation  `json:"parent"`
	RackCount    int             `json:"rack_count,omitempty"`
	Site         NestedSite      `json:"site"`
	Slug         string          `json:"slug"`
	Status       LocationStatus  `json:"status"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	Tenant       NestedTenant    `json:"tenant"`
	URL          strfmt.URI      `json:"url,omitempty"`
}

type NestedLocation struct {
	Depth     int        `json:"_depth,omitempty"`
	Display   string     `json:"display,omitempty"`
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name"`
	RackCount int        `json:"rack_count,omitempty"`
	Slug      string     `json:"slug"`
	URL       strfmt.URI `json:"url,omitempty"`
}

type LocationStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
