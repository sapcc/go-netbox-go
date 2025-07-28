// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/go-netbox-go/common"
)

type NestedSite struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Display string `json:"display,omitempty"`
}

type SiteStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Site struct {
	Asns                []NestedASN     `json:"asns"`
	CircuitCount        int             `json:"circuit_count"`
	Comments            string          `json:"comments"`
	Created             strfmt.DateTime `json:"created"`
	CustomFields        interface{}     `json:"custom_fields"`
	Description         string          `json:"description"`
	DeviceCount         int             `json:"device_count"`
	Facility            string          `json:"facility"`
	Group               NestedSiteGroup `json:"group"`
	ID                  int             `json:"id"`
	LastUpdated         strfmt.DateTime `json:"last_updated"`
	Latitude            float64         `json:"latitude"`
	Longitude           float64         `json:"longitude"`
	Name                string          `json:"name"`
	PhysicalAddress     string          `json:"physical_address"`
	PrefixCount         int             `json:"prefix_count"`
	RackCount           int             `json:"rack_count"`
	Region              NestedRegion    `json:"region"`
	ShippingAddress     string          `json:"shipping_address"`
	Slug                string          `json:"slug"`
	Status              SiteStatus      `json:"status"`
	Tags                []NestedTag     `json:"tags"`
	Tenant              Tenant          `json:"tenant"`
	TimeZone            string          `json:"time_zone"`
	URL                 strfmt.URI      `json:"url"`
	VirtualMachineCount int             `json:"virtualmachine_count"`
	VlanCount           int             `json:"vlan_count"`
}

type ListSitesRequest struct {
	common.ListParams
	Region string
}

type ListSitesResponse struct {
	common.ReturnValues
	Results []Site `json:"results"`
}
