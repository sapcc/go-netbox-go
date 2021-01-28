package models

import "github.com/sapcc/go-netbox-go/common"

type Platform struct {
	Id 					int `json:"id"`
	Url 				string `json:"url"`
	Name 				string `json:"name"`
	Slug 				string `json:"slug"`
	Manufacturer 		interface{} `json:"manufacturer"`
	NapalmDriver 		string `json:"napalm_driver"`
	NapalmArgs   		string `json:"napalm_args"`
	Description 		string `json:"description"`
	DeviceCount 		int `json:"device_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
}

type ListPlatformsRequest struct {
	common.ListParams
}

type ListPlatformsResponse struct {
	common.ReturnValues
	Results []Platform `json:"results"`
}

