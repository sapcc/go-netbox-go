package models

import (
	"encoding/json"

	"github.com/sapcc/go-netbox-go/common"
)

type NestedVirtualMachine struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

type VirtualMachine struct {
	NestedVirtualMachine
	Status           VirtualMachineStatus `json:"status"`
	Role             NestedDeviceRole     `json:"role"`
	Tenant           NestedTenant         `json:"tenant"`
	Platform         NestedPlatform       `json:"platform"`
	PrimaryIp        NestedIpAddress      `json:"primary_ip"`
	PrimaryIp4       NestedIpAddress      `json:"primary_ip4"`
	PrimaryIp6       interface{}          `json:"primary_ip6"`
	VCPUs            float64              `json:"vcpus"`
	Memory           int                  `json:"memory"`
	Disk             int                  `json:"disk"`
	Comments         string               `json:"comments"`
	LocalContextData string               `json:"local_context_data"`
	Tags             interface{}          `json:"tags"`
	CustomFields     interface{}          `json:"custom_fields"`
	ConfigContext    interface{}          `json:"config_context"`
	Created          string               `json:"created"`
	LastUpdated      string               `json:"last_updated"`
	Cluster          NestedCluster
	Site             NestedSite
}

type VirtualMachineStatus struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type WriteableVirtualMachine struct {
	Id         int     `json:"id,omitempty"`
	Url        string  `json:"url,omitempty"`
	Name       string  `json:"name"`
	Status     string  `json:"status,omitempty"`
	Site       string  `json:"site,omitempty"`
	Cluster    int     `json:"cluster"`
	Role       int     `json:"role,omitempty"`
	Tenant     int     `json:"tenant,omitempty"`
	Platform   int     `json:"platform,omitempty"`
	PrimaryIp  string  `json:"primary_ip,omitempty"`
	PrimaryIp4 string  `json:"primary_ipv4,omitempty"`
	Comments   string  `json:"comments,omitempty"`
	VCPUs      float64 `json:"vcpus,omitempty"`
	Memory     int     `json:"memory,omitempty"`
	Disk       int     `json:"disk,omitempty"`
}

func (vm *VirtualMachine) Writeable() WriteableVirtualMachine {
	res := WriteableVirtualMachine{
		Id:         vm.Id,
		Url:        vm.Url,
		Name:       vm.Name,
		Status:     vm.Status.Value,
		Site:       vm.Site.Slug,
		Cluster:    vm.Cluster.Id,
		Role:       vm.Role.Id,
		Tenant:     vm.Tenant.Id,
		Platform:   vm.Platform.Id,
		PrimaryIp:  vm.PrimaryIp.Address,
		PrimaryIp4: vm.PrimaryIp4.Address,
		Comments:   vm.Comments,
		VCPUs:      vm.VCPUs,
		Memory:     vm.Memory,
		Disk:       vm.Disk,
	}
	return res
}

type ListVirtualMachinesRequest struct {
	common.ListParams
	ClusterId int `json:"cluster_id"`
	RoleId    int `json:"role_id"`
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
	var platform NestedPlatform
	if err := json.Unmarshal(tmp["platform"], &platform); err != nil {
		return err
	}
	vm.Platform = platform
	var role NestedDeviceRole
	if err := json.Unmarshal(tmp["role"], &role); err != nil {
		return err
	}
	vm.Role = role
	var pIp NestedIpAddress
	if err := json.Unmarshal(tmp["primary_ip"], &pIp); err != nil {
		return err
	}
	vm.PrimaryIp = pIp
	var pIp4 NestedIpAddress
	if err := json.Unmarshal(tmp["primary_ip4"], &pIp4); err != nil {
		return err
	}
	vm.PrimaryIp4 = pIp4
	var status VirtualMachineStatus
	if err := json.Unmarshal(tmp["status"], &status); err != nil {
		return err
	}
	vm.Status = status
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
	var tenant NestedTenant
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
	var vcpus float64
	if err := json.Unmarshal(tmp["vcpus"], &vcpus); err != nil {
		return err
	}
	vm.VCPUs = vcpus
	var memory int
	if err := json.Unmarshal(tmp["memory"], &memory); err != nil {
		return err
	}
	vm.Memory = memory
	var disk int
	if err := json.Unmarshal(tmp["disk"], &disk); err != nil {
		return err
	}
	vm.Disk = disk
	return nil
}
