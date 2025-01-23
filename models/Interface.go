/*
 *   Copyright 2020 SAP SE
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package models

import "github.com/sapcc/go-netbox-go/common"

type NestedInterface struct {
	ID     int          `json:"id"`
	URL    string       `json:"url"`
	Device NestedDevice `json:"device"`
	Name   string       `json:"name"`
	Cable  int          `json:"cable"`
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
