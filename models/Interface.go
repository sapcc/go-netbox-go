package models

import "github.com/sapcc/go-netbox-go/common"

type NestedInterface struct {
	Id               int          `json:"id"`
	Url              string       `json:"url"`
	Device           NestedDevice `json:"device"`
	Name             string       `json:"name"`
	Cable            int          `json:"cable"`
	//ConnectionStatus interface{}  `json:"connection_status"`
}

type ConnectedEndpoint struct {
	Id               int         `json:"id"`
	Url              string      `json:"url"`
	Device           Device      `json:"device"`
	Name             string      `json:"name"`
	Cable            interface{} `json:"cable"`
	ConnectionStatus interface{} `json:"connection_status"`
}

type Interface struct {
	NestedInterface
	Label             string            `json:"label"`
	Type              InterfaceType     `json:"type"`
	Enabled           bool              `json:"enabled"`
	Lag               NestedInterface   `json:"lag"`
	MTU               int               `json:"mtu"`
	MacAddress        string            `json:"mac_address"`
	ManagementOnly    bool              `json:"mgmt_only"`
	Description       string            `json:"description"`
	Mode              interface{}       `json:"mode"`
	UntaggedVlan      NestedVLAN        `json:"untagged_vlan"`
	TaggedVlans       []NestedVLAN      `json:"tagged_vlans"`
	Tags              interface{}       `json:"tags"`
	CountIpAddresses  int               `json:"count_ipaddresses"`
	ConnectedEndpoint ConnectedEndpoint `json:"connected_endpoint"`
}

type InterfaceType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type ListInterfacesRequest struct {
	common.ListParams
	Type       string
	DeviceId   int
	MacAddress string
}

type ListInterfacesResponse struct {
	common.ReturnValues
	Results []Interface `json:"results"`
}

type WritableInterface struct {
	Device int `json:"device"`
	Name string `json:"name"`
	Label string `json:"label,omitempty"`
	Type string `json:"type"`
	Enabled bool `json:"enabled"`
	Parent int `json:"parent,omitempty"`
	Lag int `json:"lag,omitempty"`
	Mtu int `json:"mtu,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	MgmtOnly string `json:"mgmt_only,omitempty"`
	Description string `json:"description,omitempty"`
	Mode string `json:"mode,omitempty"`
	UntaggedVlan int `json:"untagged_vlan,omitempty"`
	TaggedVlan []int `json:"tagged_vlan,omitempty"`
	MarkConnected bool `json:"mark_connected,omitempty"`
	Cable NestedCable `json:"cable,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields interface{} `json:"custom_fields,omitempty"`
}