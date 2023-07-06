package models

import "github.com/sapcc/go-netbox-go/common"

type NestedInterface struct {
	Id     int          `json:"id"`
	Url    string       `json:"url"`
	Device NestedDevice `json:"device"`
	Name   string       `json:"name"`
	Cable  int          `json:"cable"`
	//ConnectionStatus interface{}  `json:"connection_status"`
}

type InterfaceMode struct {
	Label *string `json:"label"`
	Value *string `json:"value"`
}

type Interface struct {
	NestedInterface
	Cable              *NestedCable      `json:"cable,omitempty"`
	ConnectedEndpoints []NestedInterface `json:"connected_endpoints"`
	CountIpaddresses   int64             `json:"count_ipaddresses,omitempty"`
	Description        string            `json:"description,omitempty"`
	Enabled            bool              `json:"enabled,omitempty"`
	Label              string            `json:"label,omitempty"`
	Lag                *NestedInterface  `json:"lag,omitempty"`
	MacAddress         *string           `json:"mac_address,omitempty"`
	ManagementOnly     bool              `json:"mgmt_only"`
	Mode               *InterfaceMode    `json:"mode,omitempty"`
	Mtu                *int64            `json:"mtu,omitempty"`
	Name               *string           `json:"name"`
	Occupied           *bool             `json:"_occupied,omitempty"`
	TaggedVlans        []*NestedVLAN     `json:"tagged_vlans"`
	Tags               []*NestedTag      `json:"tags,omitempty"`
	Type               *InterfaceType    `json:"type"`
	UntaggedVlan       *NestedVLAN       `json:"untagged_vlan,omitempty"`
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
	Device        int         `json:"device"`
	Name          string      `json:"name"`
	Label         string      `json:"label,omitempty"`
	Type          string      `json:"type"`
	Enabled       bool        `json:"enabled"`
	Parent        int         `json:"parent,omitempty"`
	Lag           int         `json:"lag,omitempty"`
	Mtu           int         `json:"mtu,omitempty"`
	MacAddress    string      `json:"mac_address,omitempty"`
	MgmtOnly      string      `json:"mgmt_only,omitempty"`
	Description   string      `json:"description,omitempty"`
	Mode          string      `json:"mode,omitempty"`
	UntaggedVlan  int         `json:"untagged_vlan,omitempty"`
	TaggedVlan    []int       `json:"tagged_vlan,omitempty"`
	MarkConnected bool        `json:"mark_connected,omitempty"`
	Cable         NestedCable `json:"cable,omitempty"`
	Tags          []NestedTag `json:"tags,omitempty"`
	CustomFields  interface{} `json:"custom_fields,omitempty"`
}
