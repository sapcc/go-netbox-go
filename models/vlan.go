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

import "github.com/sapcc/go-netbox-go/common"

type NestedVLAN struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	VID         int    `json:"vid"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type Vlan struct {
	NestedVLAN
	Group        interface{} `json:"group"`
	Description  string      `json:"description"`
	Tags         interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
	PrefixCount  int         `json:"prefix_count"`
	Role         Role        `json:"role"`
	Site         Site        `json:"site"`
	Status       interface{} `json:"status"`
	Tenant       Tenant      `json:"tenant"`
}

type ListVlanRequest struct {
	common.ListParams
	Group string `json:"group"`
}

type ListVlanResponse struct {
	common.ReturnValues
	Results []Vlan `json:"results"`
}
