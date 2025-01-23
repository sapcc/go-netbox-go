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

import (
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/go-netbox-go/common"
)

type DeviceType struct {
	NestedDeviceType
}

type NestedDeviceType struct {
	ID           int                `json:"id"`
	URL          strfmt.URI         `json:"url"`
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	Slug         string             `json:"slug"`
	Display      string             `json:"display"`
	DeviceCount  int                `json:"device_count"`
}

type ListDeviceTypesRequest struct {
	common.ListParams
}

type ListDeviceTypesResponse struct {
	common.ReturnValues
	Results []DeviceType `json:"results"`
}
