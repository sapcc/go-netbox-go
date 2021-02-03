package models

import "github.com/sapcc/go-netbox-go/common"

type DeviceRole struct {
	NestedDeviceRole
	Color 		string 	`json:"color"`
	VMRole 		bool 	`json:"vm_role"`
	Description string 	`json:"description"`
}

type NestedDeviceRole struct {
	Id 					int 	`json:"id"`
	Url 				string 	`json:"url"`
	Name 				string 	`json:"name"`
	Slug 				string 	`json:"slug"`
	DeviceCount 		int 	`json:"device_count"`
	VirtualMachineCount	int 	`json:"virtualmachine_count"`
}

type ListDeviceRolesRequest struct {
	common.ListParams
}

type ListDeviceRolesResponse struct {
	common.ReturnValues
	Results []DeviceRole `json:"results"`
}

