// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import "github.com/sapcc/go-netbox-go/common"

type NestedVRF struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Rd          string `json:"rd"` // Route Distinguisher
	DisplayName string `json:"display_name"`
	PrefixCount int    `json:"prefix_count"`
}

type VRF struct {
	NestedVRF
	Tenant         NestedTenant `json:"tenant"`
	EnforceUnique  bool         `json:"enforce_unique"`
	Description    string       `json:"description"`
	Tags           []NestedTag  `json:"tags"`
	CustomFields   interface{}  `json:"custom_fields"`
	Created        string       `json:"created"`
	LastUpdated    string       `json:"last_updated"`
	IPAddressCount int          `json:"ipaddress_count"`
	PrefixCount    int          `json:"prefix_count"`
}

type ListVRFsRequest struct {
	common.ListParams
	Name string
}

type ListVRFsResponse struct {
	common.ReturnValues
	Results []VRF `json:"results"`
}
