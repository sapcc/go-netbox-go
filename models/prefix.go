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
	"github.com/sapcc/go-netbox-go/common"
)

type Status struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Prefix struct {
	ID           int         `json:"id"`
	URL          string      `json:"url"`
	Family       interface{} `json:"family"`
	Prefix       string      `json:"prefix"`
	Site         Site        `json:"site"`
	Vrf          NestedVRF   `json:"vrf"`
	Tenant       Tenant      `json:"tenant"`
	Vlan         NestedVLAN  `json:"vlan"`
	Status       Status      `json:"status"`
	Role         Role        `json:"role"`
	IsPool       bool        `json:"is_pool"`
	Description  string      `json:"description"`
	Tags         []NestedTag `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
}

type WriteablePrefix struct {
	ID           int         `json:"id,omitempty"`
	URL          string      `json:"url"`
	Prefix       string      `json:"prefix"`
	Site         int         `json:"site,omitempty"`
	Vrf          int         `json:"vrf,omitempty"`
	Tenant       int         `json:"tenant,omitempty"`
	Vlan         int         `json:"vlan,omitempty"`
	Status       string      `json:"status,omitempty"`
	Role         int         `json:"role,omitempty"`
	IsPool       bool        `json:"is_pool"`
	Description  string      `json:"description"`
	Tags         []NestedTag `json:"tags,omitempty"`
	CustomFields interface{} `json:"custom_fields,omitempty"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
}

type ListPrefixesRequest struct {
	common.ListParams
	Role          string
	Region        string
	Site          string
	Tag           string
	TenantID      int
	VrfID         int
	Prefix        string
	MaskLength    int
	MaskLengthGte int
	MaskLengthLte int
	Status        string
	Within        string
	Contains      string
	Children      *int
}

type ListPrefixesReponse struct {
	common.ReturnValues
	Results []Prefix `json:"results"`
}

type CreateAvailablePrefixRequest struct {
	PrefixLength int `json:"prefix_length"`
}
