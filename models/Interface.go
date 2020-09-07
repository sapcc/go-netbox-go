package models

import "github.com/sapcc/go-netbox-go/common"

type Interface struct {
	Id               int           `json:"id"`
	Url              string        `json:"url"`
	Device           Device        `json:"device"`
	Name             string        `json:"name"`
	Label            string        `json:"label"`
	Type             InterfaceType `json:"type"`
	Enabled          bool          `json:"enabled"`
	Lag              interface{}   `json:"lag"`
	MTU              int           `json:"mtu"`
	MacAddress       string        `json:"mac_address"`
	ManagementOnly   bool          `json:"mgmt_only"`
	Description      string        `json:"description"`
	ConnectionStatus interface{}   `json:"connection_status"`
	Cable            interface{}   `json:"cable"`
	Mode             interface{}   `json:"mode"`
	UntaggedVlan     interface{}   `json:"untagged_vlan"`
	TaggedVlans interface{} `json:"tagged_vlans"`
	Tags interface{} `json:"tags"`
	CountIpAddresses int `json:"count_ipaddresses"`
}

type InterfaceType struct {
	Value string `json:"value"`
	Label string `json:"label"`
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
