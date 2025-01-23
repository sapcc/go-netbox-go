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

type Platform struct {
	NestedPlatform
	Manufacturer interface{} `json:"manufacturer"`
	NapalmDriver string      `json:"napalm_driver"`
	NapalmArgs   string      `json:"napalm_args"`
	Description  string      `json:"description"`
}

type NestedPlatform struct {
	ID                  int    `json:"id"`
	URL                 string `json:"url"`
	Name                string `json:"name"`
	Slug                string `json:"slug"`
	DeviceCount         int    `json:"device_count"`
	VirtualMachineCount int    `json:"virtualmachine_count"`
	Display             string `json:"display,omitempty"`
}

type ListPlatformsRequest struct {
	common.ListParams
}

type ListPlatformsResponse struct {
	common.ReturnValues
	Results []Platform `json:"results"`
}
