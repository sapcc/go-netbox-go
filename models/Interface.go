// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import "github.com/sapcc/go-netbox-go/common"

type NestedInterface struct {
	ID     int          `json:"id"`
	URL    string       `json:"url"`
	Device NestedDevice `json:"device"`
	Name   string       `json:"name"`
	Cable  NestedCable  `json:"cable"`
	// ConnectionStatus interface{}  `json:"connection_status"`
}

type InterfaceMode struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Interface struct {
	NestedInterface
	Cable              NestedCable       `json:"cable,omitempty"`
	ConnectedEndpoints []NestedInterface `json:"connected_endpoints"`
	CountIpaddresses   int               `json:"count_ipaddresses,omitempty"`
	Description        string            `json:"description,omitempty"`
	Enabled            bool              `json:"enabled,omitempty"`
	Label              string            `json:"label,omitempty"`
	Lag                NestedInterface   `json:"lag,omitempty"`
	MacAddress         string            `json:"mac_address,omitempty"`
	MgmtOnly           bool              `json:"mgmt_only,omitempty"`
	Mode               InterfaceMode     `json:"mode,omitempty"`
	Mtu                int               `json:"mtu,omitempty"`
	Name               string            `json:"name"`
	Occupied           bool              `json:"_occupied,omitempty"`
	TaggedVlans        []NestedVLAN      `json:"tagged_vlans"`
	Tags               []NestedTag       `json:"tags,omitempty"`
	Type               InterfaceType     `json:"type"`
	UntaggedVlan       NestedVLAN        `json:"untagged_vlan,omitempty"`
}

type InterfaceType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type ListInterfacesRequest struct {
	common.ListParams
	Type       string
	DeviceID   int
	MacAddress string
	LagID      int
}

type ListInterfacesResponse struct {
	common.ReturnValues
	Results []Interface `json:"results"`
}

type WritableInterface struct {
	Cable         NestedCable `json:"cable,omitempty"`
	CustomFields  interface{} `json:"custom_fields,omitempty"`
	Description   string      `json:"description,omitempty"`
	Device        int         `json:"device"`
	Enabled       bool        `json:"enabled,omitempty"`
	Label         string      `json:"label,omitempty"`
	Lag           int         `json:"lag,omitempty"`
	MacAddress    string      `json:"mac_address,omitempty"`
	MarkConnected bool        `json:"mark_connected,omitempty"`
	MgmtOnly      bool        `json:"mgmt_only,omitempty"`
	Mode          string      `json:"mode,omitempty"`
	Mtu           int         `json:"mtu,omitempty"`
	Name          string      `json:"name"`
	Parent        int         `json:"parent,omitempty"`
	TaggedVlans   []int       `json:"tagged_vlans,omitempty"`
	Tags          []NestedTag `json:"tags,omitempty"`
	Type          string      `json:"type"`
	UntaggedVlan  int         `json:"untagged_vlan,omitempty"`
}
