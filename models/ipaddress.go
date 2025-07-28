// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"fmt"

	"github.com/sapcc/go-netbox-go/common"
)

type NestedIPAddress struct {
	ID      int         `json:"id,omitempty"`
	URL     string      `json:"url,omitempty"`
	Family  interface{} `json:"family,omitempty"`
	Address string      `json:"address"`
}

type AvailableIP struct {
	Family  int         `json:"family"`
	Address string      `json:"address"`
	Vrf     interface{} `json:"vrf"`
}

type IpamRole struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type IPAddressStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type WriteableIPAddress struct {
	NestedIPAddress
	Vrf                int         `json:"vrf,omitempty"`
	Tenant             int         `json:"tenant,omitempty"`
	Status             string      `json:"status,omitempty"`
	Role               string      `json:"role,omitempty"`
	AssignedObjectType string      `json:"assigned_object_type,omitempty"`
	AssignedObjectID   int         `json:"assigned_object_id,omitempty"`
	NatInside          int         `json:"nat_inside,omitempty"`
	NatOutside         int         `json:"nat_outside,omitempty"`
	DNSName            string      `json:"dns_name,omitempty"`
	Description        string      `json:"description,omitempty"`
	Tags               []NestedTag `json:"tags,omitempty"`
	CustomFields       interface{} `json:"custom_fields,omitempty"`
	Created            string      `json:"created,omitempty"`
	LastUpdated        string      `json:"last_updated,omitempty"`
}

type IPAddress struct {
	NestedIPAddress
	Vrf                 interface{}     `json:"vrf"`
	Tenant              Tenant          `json:"tenant"`
	Status              IPAddressStatus `json:"status"`
	Role                IpamRole        `json:"role"`
	NatInside           interface{}     `json:"nat_inside"`
	NatOutside          interface{}     `json:"nat_outside"`
	DNSName             string          `json:"dns_name"`
	Description         string          `json:"description"`
	Tags                []NestedTag     `json:"tags"`
	CustomFields        interface{}     `json:"custom_fields"`
	Created             string          `json:"created"`
	LastUpdated         string          `json:"last_updated"`
	AssignedInterface   NestedInterface
	AssignedVMInterface NestedVMInterface
	AssignedObjectType  string `json:"assigned_object_type"`
	AssignedObjectID    int    `json:"assigned_object_id"`
}

type ListIPAddressesRequest struct {
	common.ListParams
	InterfaceID   int
	VMInterfaceID int
	DeviceID      int
	Role          string
	Address       string
	VrfID         int
	Parent        string
}

type ListIPAddressesResponse struct {
	common.ReturnValues
	Results []IPAddress `json:"results"`
}

func (ip *IPAddress) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	var i int
	if err := json.Unmarshal(tmp["assigned_object_id"], &i); err != nil {
		return err
	}
	ip.AssignedObjectID = i
	var s string
	if err := json.Unmarshal(tmp["assigned_object_type"], &s); err != nil {
		return err
	}
	switch s {
	case "dcim.interface":
		var inter NestedInterface
		if err := json.Unmarshal(tmp["assigned_object"], &inter); err != nil {
			return err
		}
		ip.AssignedInterface = inter
		ip.AssignedObjectType = s
	case "virtualization.vminterface":
		var inter NestedVMInterface
		if err := json.Unmarshal(tmp["assigned_object"], &inter); err != nil {
			fmt.Println("unable to unmarshal vminterface")
			fmt.Println(string(tmp["assigned_object"]))
			return err
		}
		ip.AssignedVMInterface = inter
		ip.AssignedObjectType = s
	case "":

	default:
		return fmt.Errorf("unknown assigned object type %v", s)
	}
	var role IpamRole
	if err := json.Unmarshal(tmp["role"], &role); err != nil {
		return err
	}
	ip.Role = role
	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	ip.ID = id
	var url string
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	ip.URL = url
	var addr string
	if err := json.Unmarshal(tmp["address"], &addr); err != nil {
		return err
	}
	ip.Address = addr
	var tenant Tenant
	if err := json.Unmarshal(tmp["tenant"], &tenant); err != nil {
		return err
	}
	ip.Tenant = tenant
	var status IPAddressStatus
	if err := json.Unmarshal(tmp["status"], &status); err != nil {
		return err
	}
	ip.Status = status
	var dnsName string
	if err := json.Unmarshal(tmp["dns_name"], &dnsName); err != nil {
		return err
	}
	ip.DNSName = dnsName
	var descr string
	if err := json.Unmarshal(tmp["description"], &descr); err != nil {
		return err
	}
	ip.Description = descr
	var created string
	if err := json.Unmarshal(tmp["created"], &created); err != nil {
		return err
	}
	ip.Created = created
	var lastUpdated string
	if err := json.Unmarshal(tmp["last_updated"], &lastUpdated); err != nil {
		return err
	}
	ip.LastUpdated = lastUpdated
	if _, ok := tmp["tags"]; ok {
		var tags []NestedTag
		if err := json.Unmarshal(tmp["tags"], &tags); err != nil {
			return err
		}
		ip.Tags = tags
	}
	return nil
}
