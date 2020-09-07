package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type Site struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Status interface{} `json:"status"`
	Region interface{} `json:"region"`
	Tenant Tenant `json:"tenant"`
	Facility string `json:"facility"`
	ASN int `json:"asn"`
	TimeZone string `json:"time_zone"`
	Description string `json:"description"`
	PhysicalAddress string `json:"physical_address"`
	ShippingAddress string `json:"shipping_address"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	ContactName string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
	Comments string `json:"comments"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	CircuitCount int `json:"circuit_count"`
	DeviceCount int `json:"device_count"`
	PrefixCount int `json:"prefix_count"`
	RackCount int `json:"rack_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
	VlanCount int `json:"vlan_count"`
}

type ListSitesRequest struct {
	common.ListParams
}

type ListSitesResponse struct {
	common.ReturnValues
	Results []Site `json:"results"`
}
