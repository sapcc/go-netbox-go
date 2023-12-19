package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type DeviceType struct {
	NestedDeviceType
}

type NestedDeviceType struct {
	Id           int                `json:"id"`
	Url          strfmt.URI         `json:"url"`
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
