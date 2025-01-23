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

type Tenant struct {
	NestedTenant
	Group               NestedTenantGroup `json:"group"`
	Description         string            `json:"description"`
	Comments            string            `json:"comments"`
	Tags                interface{}       `json:"tags"`
	CustomFields        interface{}       `json:"custom_fields"`
	Created             string            `json:"created"`
	LastUpdated         string            `json:"last_updated"`
	CircuitCount        int               `json:"circuit_count"`
	DeviceCount         int               `json:"device_count"`
	IPAddressCount      int               `json:"ipaddress_count"`
	PrefixCount         int               `json:"prefix_count"`
	RackCount           int               `json:"rack_count"`
	SiteCount           int               `json:"site_count"`
	VirtualMachineCount int               `json:"virtualmachine_count"`
	VlanCount           int               `json:"vlan_count"`
	VrfCount            int               `json:"vrf_count"`
	ClusterCount        int               `json:"cluster_count"`
}

type NestedTenant struct {
	ID      int        `json:"id"`
	URL     strfmt.URI `json:"url,omitempty"`
	Name    string     `json:"name"`
	Slug    string     `json:"slug"`
	Display string     `json:"display,omitempty"`
}

type ListTenantsRequest struct {
	common.ListParams
}

type ListTenantsResponse struct {
	common.ReturnValues
	Results []Tenant `json:"results"`
}
