// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

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
