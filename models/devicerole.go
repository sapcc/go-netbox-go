// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import "github.com/sapcc/go-netbox-go/common"

type DeviceRole struct {
	NestedDeviceRole
	Color       string `json:"color"`
	VMRole      bool   `json:"vm_role"`
	Description string `json:"description"`
}

type NestedDeviceRole struct {
	ID                  int    `json:"id"`
	URL                 string `json:"url"`
	Name                string `json:"name"`
	Slug                string `json:"slug"`
	DeviceCount         int    `json:"device_count"`
	Display             string `json:"display"`
	VirtualMachineCount int    `json:"virtualmachine_count"`
}

type ListDeviceRolesRequest struct {
	common.ListParams
}

type ListDeviceRolesResponse struct {
	common.ReturnValues
	Results []DeviceRole `json:"results"`
}
