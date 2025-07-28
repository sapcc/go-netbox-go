// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import "github.com/go-openapi/strfmt"

type NestedASN struct {
	Asn     int        `json:"asn"`
	Display string     `json:"display,omitempty"`
	ID      int        `json:"id,omitempty"`
	URL     strfmt.URI `json:"url,omitempty"`
}

type ASN struct {
	Asn           int             `json:"asn"`
	Comments      string          `json:"comments,omitempty"`
	Created       strfmt.DateTime `json:"created,omitempty"`
	CustomFields  interface{}     `json:"custom_fields,omitempty"`
	Description   string          `json:"description,omitempty"`
	Display       string          `json:"display,omitempty"`
	ID            int             `json:"id,omitempty"`
	LastUpdated   strfmt.DateTime `json:"last_updated,omitempty"`
	ProviderCount int             `json:"provider_count,omitempty"`
	Rir           int             `json:"rir"`
	SiteCount     int             `json:"site_count,omitempty"`
	Tags          []NestedTag     `json:"tags,omitempty"`
	Tenant        NestedTenant    `json:"tenant,omitempty"`
	URL           strfmt.URI      `json:"url,omitempty"`
}
