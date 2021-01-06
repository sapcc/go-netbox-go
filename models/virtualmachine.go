package models

import (
	"encoding/json"
	"github.com/sapcc/go-netbox-go/common"
)

type NestedVirtualMachine struct {
	Id         int            `json:"id"`
	Url        string         `json:"url"`
	Name       string         `json:"name"`
}

type VirtualMachine struct {
	Id         int            `json:"id"`
	Url        string         `json:"url"`
	Name       string         `json:"name"`
	Status     interface{}    `json:"status"`
	Role       interface{}    `json:"role"`
	Tenant     Tenant `json:"tenant"`
	Platform   interface{}    `json:"platform"`
	PrimaryIp  interface{}    `json:"primary_ip"`
	PrimaryIp4 interface{}    `json:"primary_ip4"`
	PrimaryIp6 interface{}    `json:"primary_ip6"`
	VCPUs      int            `json:"vcpus"`
	Memory     int            `json:"memory"`
	Disk       int            `json:"disk"`
	Comments string `json:"comments"`
	LocalContextData string `json:"local_context_data"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	ConfigContext interface{} `json:"config_context"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	Cluster    NestedCluster
	Site	   NestedSite
}

type ListVirtualMachinesRequest struct {
	common.ListParams
}

type ListVirtualMachinesResponse struct {
	common.ReturnValues
	Results []VirtualMachine `json:"results"`
}

func (vm *VirtualMachine) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	var cl NestedCluster
	if err := json.Unmarshal(tmp["cluster"], &cl); err != nil {
		return err
	}
	vm.Cluster = cl
	var st NestedSite
	if err := json.Unmarshal(tmp["site"], &st); err != nil {
		return err
	}
	vm.Site = st
	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	vm.Id = id
	var url string
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	vm.Url = url
	var Name string
	if err := json.Unmarshal(tmp["name"], &Name); err != nil {
		return err
	}
	vm.Name = Name
	var tenant Tenant
	if err := json.Unmarshal(tmp["tenant"], &tenant); err != nil {
		return err
	}
	vm.Tenant = tenant
	var created string
	if err := json.Unmarshal(tmp["created"], &created); err != nil {
		return err
	}
	vm.Created = created
	var lastUpdated string
	if err := json.Unmarshal(tmp["last_updated"], &lastUpdated); err != nil {
		return err
	}
	vm.LastUpdated = lastUpdated
	return nil
}
