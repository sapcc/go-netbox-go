// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

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
