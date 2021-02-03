package models

import "github.com/sapcc/go-netbox-go/common"

type Tenant struct {
	NestedTenant
	Group interface{} `json:"group"`
	Description string `json:"description"`
	Comments string `json:"comments"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	CircuitCount int `json:"circuit_count"`
	DeviceCount int `json:"device_count"`
	IpAddressCount int `json:"ipaddress_count"`
	PrefixCount int `json:"prefix_count"`
	RackCount int `json:"rack_count"`
	SiteCount int `json:"site_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
	VlanCount int `json:"vlan_count"`
	VrfCount int `json:"vrf_count"`
	ClusterCount int `json:"cluster_count"`
}

type NestedTenant struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug string	`json:"slug"`
}

type ListTenantsRequest struct {
	common.ListParams
}

type ListTenantsResponse struct {
	common.ReturnValues
	Results []Tenant `json:"results"`
}