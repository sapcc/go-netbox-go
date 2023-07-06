package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type DeviceType struct {
	Id           int          `json:"id"`
	Url          string       `json:"url"`
	Manufacturer Manufacturer `json:"manufacturer"`
	Model        string       `json:"model"`
	Slug         string       `json:"slug"`
	DisplayName  string       `json:"display_name"`
	DeviceCount  int          `json:"device_count"`
}

type Manufacturer struct {
	Id               int    `json:"id"`
	Url              string `json:"url"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Devicetype_count int    `json:"devicetype_count"`
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

type Device struct {
	NestedDevice
	DeviceType       DeviceType       `json:"device_type"`
	DeviceRole       NestedDeviceRole `json:"device_role"`
	Tenant           NestedTenant     `json:"tenant"`
	Platform         interface{}      `json:"platform"`
	Serial           string           `json:"serial"`
	AssetTag         string           `json:"asset_tag"`
	Site             Site             `json:"site"`
	Rack             Rack             `json:"rack"`
	Position         float64          `json:"position"`
	Face             interface{}      `json:"face"`
	ParentDevice     interface{}      `json:"parent_device"`
	Status           DeviceStatus     `json:"status"`
	PrimaryIp        NestedIpAddress  `json:"primary_ip"`
	PrimaryIp4       NestedIpAddress  `json:"primary_ip4"`
	PrimaryIp6       NestedIpAddress  `json:"primary_ip6"`
	Cluster          Cluster          `json:"cluster"`
	VirtualChassis   interface{}      `json:"virtual_chassis"`
	VCPosition       int              `json:"vc_position"`
	VCPriority       int              `json:"vc_priority"`
	Comments         string           `json:"comments"`
	LocalContextData interface{}      `json:"local_context_data"`
	Tags             interface{}      `json:"tags"`
	CustomFields     interface{}      `json:"custom_fields"`
	ConfigContext    interface{}      `json:"config_context"`
	Created          string           `json:"created"`
	LastUpdated      string           `json:"last_updated"`
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
