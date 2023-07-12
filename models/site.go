package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type NestedSite struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Display string `json:"display,omitempty"`
}

type Site struct {
	Id                  int             `json:"id"`
	Url                 string          `json:"url"`
	Name                string          `json:"name"`
	Slug                string          `json:"slug"`
	Status              Status          `json:"status"`
	Region              NestedRegion    `json:"region"`
	Tenant              Tenant          `json:"tenant"`
	Facility            string          `json:"facility"`
	Group               NestedSiteGroup `json:"group"`
	Asns                interface{}     `json:"asns"`
	TimeZone            string          `json:"time_zone"`
	Description         string          `json:"description"`
	PhysicalAddress     string          `json:"physical_address"`
	ShippingAddress     string          `json:"shipping_address"`
	Latitude            float64         `json:"latitude"`
	Longitude           float64         `json:"longitude"`
	Comments            string          `json:"comments"`
	Tags                []NestedTag     `json:"tags"`
	CustomFields        interface{}     `json:"custom_fields"`
	Created             string          `json:"created"`
	LastUpdated         string          `json:"last_updated"`
	CircuitCount        int             `json:"circuit_count"`
	DeviceCount         int             `json:"device_count"`
	PrefixCount         int             `json:"prefix_count"`
	RackCount           int             `json:"rack_count"`
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
