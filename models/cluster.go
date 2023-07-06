package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type NestedCluster struct {
	Display             string     `json:"display,omitempty"`
	Id                  int        `json:"id,omitempty"`
	Name                string     `json:"name"`
	Url                 strfmt.URI `json:"url,omitempty"`
	VirtualmachineCount int        `json:"virtualmachine_count,omitempty"`
}

type NestedClusterType struct {
	Id           int    `json:"id"`
	Url          string `json:"url"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ClusterCount int    `json:"cluster_count"`
}

type Cluster struct {
	Id                  int               `json:"id"`
	Url                 string            `json:"url"`
	Name                string            `json:"name"`
	Group               interface{}       `json:"group"`
	Comments            string            `json:"comments"`
	Tags                interface{}       `json:"tags"`
	CustomFields        interface{}       `json:"custom_fields"`
	Created             string            `json:"created"`
	LastUpdated         string            `json:"last_updated"`
	DeviceCount         int               `json:"device_count"`
	Tenant              Tenant            `json:"tenant"`
	Site                Site              `json:"site"`
	Type                NestedClusterType `json:"type"`
	VirtualMachineCount int               `json:"virtualmachine_count"`
	VlanCount           int               `json:"vlan_count"`
}

type ListClusterRequest struct {
	common.ListParams
	Type   string
	Region string
}

type ListClusterResponse struct {
	common.ReturnValues
	Results []Cluster `json:"results"`
}
