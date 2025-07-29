// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/go-netbox-go/common"
)

type NestedManufacturer struct {
	ID              int    `json:"id"`
	URL             string `json:"url"`
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
	ID          int    `json:"id"`
	URL         string `json:"url"`
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
	DeviceRole       NestedDeviceRole     `json:"role"`
	DeviceType       NestedDeviceType     `json:"device_type"`
	Display          string               `json:"display,omitempty"`
	Face             DeviceFace           `json:"face"`
	ID               int                  `json:"id,omitempty"`
	LastUpdated      strfmt.DateTime      `json:"last_updated,omitempty"`
	LocalContextData interface{}          `json:"local_context_data"`
	Location         NestedLocation       `json:"location,omitempty"`
	Name             string               `json:"name,omitempty"`
	ParentDevice     NestedDevice         `json:"parent_device"`
	Platform         NestedPlatform       `json:"platform"`
	Position         float64              `json:"position"`
	PrimaryIP        NestedIPAddress      `json:"primary_ip"`
	PrimaryIP4       NestedIPAddress      `json:"primary_ip4"`
	PrimaryIP6       NestedIPAddress      `json:"primary_ip6"`
	OOBIp            NestedIPAddress      `json:"oob_ip"`
	Rack             NestedRack           `json:"rack"`
	Serial           string               `json:"serial"`
	Site             NestedSite           `json:"site"`
	Status           DeviceStatus         `json:"status"`
	Tags             []NestedTag          `json:"tags,omitempty"`
	Tenant           NestedTenant         `json:"tenant"`
	VCPosition       int                  `json:"vc_position"`
	VCPriority       int                  `json:"vc_priority"`
	VirtualChassis   NestedVirtualChassis `json:"virtual_chassis"`
}

type ListDevicesRequest struct {
	common.ListParams
	ClusterID    int
	DeviceTypeID int
	Region       string
	Site         string
	SiteID       int
	RackID       int
	Serial       string
	RoleID       int
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
	DeviceRole       int             `json:"role"`
	DeviceType       int             `json:"device_type"`
	Display          string          `json:"display,omitempty"`
	Face             string          `json:"face,omitempty"`
	ID               int             `json:"id,omitempty"`
	LastUpdated      strfmt.DateTime `json:"last_updated,omitempty"`
	LocalContextData interface{}     `json:"local_context_data,omitempty"`
	Location         int             `json:"location,omitempty"`
	Name             string          `json:"name,omitempty"`
	ParentDevice     *NestedDevice   `json:"parent_device,omitempty"`
	Platform         int             `json:"platform,omitempty"`
	Position         float64         `json:"position,omitempty"`
	PrimaryIP        string          `json:"primary_ip,omitempty"`
	PrimaryIP4       int             `json:"primary_ip4,omitempty"`
	PrimaryIP6       int             `json:"primary_ip6,omitempty"`
	OOBIp            int             `json:"oob_ip,omitempty"`
	Rack             int             `json:"rack,omitempty"`
	Serial           string          `json:"serial,omitempty"`
	Site             int             `json:"site"`
	Status           string          `json:"status,omitempty"`
	Tags             []NestedTag     `json:"tags,omitempty"`
	Tenant           int             `json:"tenant,omitempty"`
	URL              strfmt.URI      `json:"url,omitempty"`
	VcPosition       int             `json:"vc_position,omitempty"`
	VcPriority       int             `json:"vc_priority,omitempty"`
	VirtualChassis   int             `json:"virtual_chassis,omitempty"`
}

func (dev *Device) Writeable() WritableDeviceWithConfigContext {
	var clusterID int
	if dev.Cluster == nil {
		clusterID = 0
	} else {
		clusterID = dev.Cluster.ID
	}
	res := WritableDeviceWithConfigContext{
		ID:         dev.ID,
		Cluster:    clusterID,
		DeviceRole: dev.DeviceRole.ID,
		DeviceType: dev.DeviceType.ID,
		PrimaryIP4: dev.PrimaryIP4.ID,
		OOBIp:      dev.OOBIp.ID,
		Tags:       dev.Tags,
		Site:       dev.Site.ID,
	}
	return res
}
