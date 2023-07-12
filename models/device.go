package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type NestedDeviceType struct {
	Id           int                `json:"id"`
	Url          strfmt.URI         `json:"url"`
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	Slug         string             `json:"slug"`
	Display      string             `json:"display"`
	DeviceCount  int                `json:"device_count"`
}

type NestedManufacturer struct {
	Id              int    `json:"id"`
	Url             string `json:"url"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	DevicetypeCount int    `json:"devicetype_count"`
	Display         string `json:"display,omitempty"`
}

type DeviceStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type NestedDevice struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}
type DeviceAirflow struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type DeviceFace struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Device struct {
	NestedDevice
	Airflow          *DeviceAirflow       `json:"airflow,omitempty"`
	AssetTag         string               `json:"asset_tag,omitempty"`
	Cluster          *NestedCluster       `json:"cluster,omitempty"`
	Comments         string               `json:"comments,omitempty"`
	ConfigContext    interface{}          `json:"config_context"`
	Created          strfmt.DateTime      `json:"created,omitempty"`
	CustomFields     interface{}          `json:"custom_fields,omitempty"`
	Description      string               `json:"description,omitempty"`
	DeviceRole       NestedDeviceRole     `json:"device_role"`
	DeviceType       NestedDeviceType     `json:"device_type"`
	Display          string               `json:"display,omitempty"`
	Face             DeviceFace           `json:"face"`
	Id               int                  `json:"id,omitempty"`
	LastUpdated      strfmt.DateTime      `json:"last_updated,omitempty"`
	LocalContextData interface{}          `json:"local_context_data"`
	Location         NestedLocation       `json:"location,omitempty"`
	Name             string               `json:"name,omitempty"`
	ParentDevice     NestedDevice         `json:"parent_device"`
	Platform         NestedPlatform       `json:"platform"`
	Position         float64              `json:"position"`
	PrimaryIp        NestedIpAddress      `json:"primary_ip"`
	PrimaryIp4       NestedIpAddress      `json:"primary_ip4"`
	PrimaryIp6       NestedIpAddress      `json:"primary_ip6"`
	Rack             NestedRack           `json:"rack"`
	Serial           string               `json:"serial"`
	Site             NestedSite           `json:"site"`
	Status           DeviceStatus         `json:"status"`
	Tags             []NestedTag          `json:"tags"`
	Tenant           NestedTenant         `json:"tenant"`
	VCPosition       int                  `json:"vc_position"`
	VCPriority       int                  `json:"vc_priority"`
	VirtualChassis   NestedVirtualChassis `json:"virtual_chassis"`
}

type ListDevicesRequest struct {
	common.ListParams
	ClusterId    int
	DeviceTypeId int
	Region       string
	Site         string
	RackId       int
	Serial       string
}

type ListDevicesResponse struct {
	common.ReturnValues
	Results []Device `json:"results"`
}

type WritableDeviceWithConfigContext struct {
	Airflow          string          `json:"airflow,omitempty"`
	AssetTag         string          `json:"asset_tag,omitempty"`
	Cluster          int             `json:"cluster,omitempty"`
	Comments         string          `json:"comments,omitempty"`
	ConfigContext    interface{}     `json:"config_context,omitempty"`
	Created          strfmt.DateTime `json:"created,omitempty"`
	CustomFields     interface{}     `json:"custom_fields,omitempty"`
	Description      string          `json:"description,omitempty"`
	DeviceRole       int             `json:"device_role"`
	DeviceType       int             `json:"device_type"`
	Display          string          `json:"display,omitempty"`
	Face             string          `json:"face,omitempty"`
	Id               int             `json:"id,omitempty"`
	LastUpdated      strfmt.DateTime `json:"last_updated,omitempty"`
	LocalContextData interface{}     `json:"local_context_data,omitempty"`
	Location         int             `json:"location,omitempty"`
	Name             string          `json:"name,omitempty"`
	ParentDevice     *NestedDevice   `json:"parent_device,omitempty"`
	Platform         int             `json:"platform,omitempty"`
	Position         float64         `json:"position,omitempty"`
	PrimaryIP        string          `json:"primary_ip,omitempty"`
	PrimaryIp4       int             `json:"primary_ip4,omitempty"`
	PrimaryIp6       int             `json:"primary_ip6,omitempty"`
	Rack             int             `json:"rack,omitempty"`
	Serial           string          `json:"serial,omitempty"`
	Site             int             `json:"site"`
	Status           string          `json:"status,omitempty"`
	Tags             []NestedTag     `json:"tags,omitempty"`
	Tenant           int             `json:"tenant,omitempty"`
	Url              strfmt.URI      `json:"url,omitempty"`
	VcPosition       int             `json:"vc_position,omitempty"`
	VcPriority       int             `json:"vc_priority,omitempty"`
	VirtualChassis   int             `json:"virtual_chassis,omitempty"`
}
