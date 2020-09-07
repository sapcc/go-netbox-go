package dcim

import (
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/tenancy"
)

type Interface struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Device Device `json:"device"`
	Name string `json:"name"`
	Label string `json:"label"`
	Type InterfaceType `json:"type"`
	Enabled bool `json:"enabled"`
	Lag interface{} `json:"lag"`
	MTU int `json:"mtu"`
	MacAddress string `json:"mac_address"`
	ManagementOnly bool `json:"mgmt_only"`
	Description string `json:"description"`
	ConnectionStatus interface{} `json:"connection_status"`
	Cable interface{} `json:"cable"`
	Mode interface{} `json:"mode"`
	UntaggedVlan interface{} `json:"untagged_vlan"`
	TaggedVlans interface{} `json:"tagged_vlans"`
	Tags interface{} `json:"tags"`
	CountIpAddresses int `json:"count_ipaddresses"`
}

type InterfaceType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Device struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	DeviceType interface{} `json:"device_type"`
	Site Site `json:"site"`
}

type ListInterfacesRequest struct {
	common.ListParams
	Type string
	DeviceId int
}

type ListInterfacesResponse struct {
	common.ReturnValues
	Results []Interface `json:"results"`
}

type Site struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Slug string `json:"slug"`
	Status interface{} `json:"status"`
	Region interface{} `json:"region"`
	Tenant tenancy.Tenant `json:"tenant"`
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

type Rack struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	FacilityId string `json:"facility_id"`
	DisplayName string `json:"display_name"`
	Site Site `json:"site"`
	Group interface{} `json:"group"`
	Tenant tenancy.Tenant `json:"tenant"`
	Status interface{} `json:"status"`
	Role interface{} `json:"role"`
	Serial string `json:"serial"`
	AssetTag string `json:"asset_tag"`
	Type interface{} `json:"type"`
	Width interface{} `json:width`
	UHeight int `json:"u_height"`
	DescUnits bool `json:"desc_units"`
	OuterWidth int `json:"outer_width"`
	OuterDepth int `json:"outer_depth"`
	OuterUnit interface{} `json:"outer_unit"`
	Comments string `json:"comments"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	DeviceCount int `json:"device_count"`
	PowerfeedCount int `json:"powerfeed_count"`
}

type ListRacksRequest struct {
	common.ListParams
}

type ListRacksResponse struct {
	common.ReturnValues
	Results []Rack `json:"results"`
}