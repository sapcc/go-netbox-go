package virtualization

import (
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/tenancy"
)

type VirtualMachine struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Status interface{} `json:"status"`
	Site dcim.Site `json:"site"`
	Cluster interface{} `json:"cluster"`
	Role interface{} `json:"role"`
	Tenant tenancy.Tenant `json:"tenant"`
	Platform interface{} `json:"platform"`
	PrimaryIp interface{} `json:"primary_ip"`
	PrimaryIp4 interface{} `json:"primary_ip4"`
	PrimaryIp6 interface{} `json:"primary_ip6"`
	VCPUs int `json:"vcpus"`
	Memory int `json:"memory"`
	Disk int `json:"disk"`
	Comments string `json:"comments"`
	LocalContextData string `json:"local_context_data"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	ConfigContext interface{} `json:"config_context"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

type ListVirtualMachinesRequest struct {
	common.ListParams
}

type ListVirtualMachinesResponse struct {
	common.ReturnValues
	Results []VirtualMachine `json:"results"`
}

