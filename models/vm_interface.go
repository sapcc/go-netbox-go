package models

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type NestedVMInterface struct {
	Display        string                `json:"display,omitempty"`
	Id             int                   `json:"id,omitempty"`
	Name           string                `json:"name"`
	URL            strfmt.URI            `json:"url,omitempty"`
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine,omitempty"`
}
type VMInterface struct {
	Enabled        *bool                `json:"enabled,omitempty"`
	Url            string               `json:"url"`
	Id             int                  `json:"id"`
	Name           string               `json:"name"`
	MTU            *int                 `json:"mtu,omitempty"`
	MacAddress     *string              `json:"mac_address,omitempty"`
	Description    *string              `json:"description,omitempty"`
	Mode           *string              `json:"mode,omitempty"`
	Tags           []NestedTag          `json:"tags,omitempty"`
	VirtualMachine NestedVirtualMachine `json:"virtual_machine"`
	UntaggedVlan   *NestedVLAN          `json:"untagged_vlan,omitempty"`
	TaggedVlans    []NestedVLAN         `json:"tagged_vlans,omitempty"`
}

type WritableVMInterface struct {
	Id             int         `json:"id,omitempty"`
	Url            string      `json:"url,omitempty"`
	VirtualMachine int         `json:"virtual_machine,omitempty"`
	Name           string      `json:"name"`
	Enabled        bool        `json:"enabled,omitempty"`
	MTU            int         `json:"mtu,omitempty"`
	MacAddress     *string     `json:"mac_address,omitempty"`
	Description    string      `json:"description,omitempty"`
	Mode           string      `json:"mode,omitempty"`
	UntaggedVlan   *int        `json:"untagged_vlan,omitempty"`
	TaggedVlans    []int       `json:"tagged_vlans,omitempty"`
	Tags           []NestedTag `json:"tags,omitempty"`
}

func (vmi *VMInterface) Writeable() WritableVMInterface {

	res := WritableVMInterface{
		Id:             vmi.Id,
		Url:            vmi.Url,
		VirtualMachine: vmi.VirtualMachine.Id,
		Name:           vmi.Name,
	}
	if vmi.MacAddress != nil {
		res.MacAddress = vmi.MacAddress
	}
	if vmi.Enabled != nil {
		res.Enabled = *vmi.Enabled
	}
	if vmi.MTU != nil {
		res.MTU = *vmi.MTU
	}
	if vmi.Description != nil {
		res.Description = *vmi.Description
	}
	if vmi.Mode != nil {
		res.Mode = *vmi.Mode
	}
	if vmi.UntaggedVlan != nil {
		res.UntaggedVlan = &vmi.UntaggedVlan.Id
	}
	var taggedVlans = make([]int, len(vmi.TaggedVlans))
	for _, vlan := range vmi.TaggedVlans {
		taggedVlans = append(taggedVlans, vlan.Id)
	}
	res.TaggedVlans = taggedVlans
	res.Tags = vmi.Tags

	return res
}

type ListVMInterfacesRequest struct {
	common.ListParams
	VmId int
}

type ListVMInterfacesResponse struct {
	common.ReturnValues
	Results []VMInterface `json:"results"`
}

func (vmIf *VMInterface) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	if _, ok := tmp["tags"]; ok {
		var tags []NestedTag
		if err := json.Unmarshal(tmp["tags"], &tags); err != nil {
			return err
		}
		vmIf.Tags = tags
	}
	var vm NestedVirtualMachine
	if err := json.Unmarshal(tmp["virtual_machine"], &vm); err != nil {
		return err
	}
	vmIf.VirtualMachine = vm
	if _, ok := tmp["tagged_vlans"]; ok {
		var tagVlan []NestedVLAN
		if err := json.Unmarshal(tmp["tagged_vlans"], &tagVlan); err != nil {
			return err
		}
		vmIf.TaggedVlans = tagVlan
	}
	if _, ok := tmp["untagged_vlan"]; ok {
		var untagVlan NestedVLAN
		if err := json.Unmarshal(tmp["untagged_vlan"], &untagVlan); err != nil {
			return err
		}
		vmIf.UntaggedVlan = &untagVlan
	}
	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	vmIf.Id = id
	var url string
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	vmIf.Url = url
	var Name string
	if err := json.Unmarshal(tmp["name"], &Name); err != nil {
		return err
	}
	vmIf.Name = Name
	if _, ok := tmp["description"]; ok {
		var descr string
		if err := json.Unmarshal(tmp["description"], &descr); err != nil {
			return err
		}
		vmIf.Description = &descr
	}
	if _, ok := tmp["mac_address"]; ok {
		var mac string
		if err := json.Unmarshal(tmp["mac_address"], &mac); err != nil {
			return err
		}
		vmIf.MacAddress = &mac
	}
	if _, ok := tmp["mtu"]; ok {
		var mtu int
		if err := json.Unmarshal(tmp["mtu"], &mtu); err != nil {
			return err
		}
		vmIf.MTU = &mtu
	}
	if _, ok := tmp["enabled"]; ok {
		var en bool
		if err := json.Unmarshal(tmp["enabled"], &en); err != nil {
			return err
		}
		vmIf.Enabled = &en
	}
	return nil
}
