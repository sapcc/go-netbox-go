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

type NestedCluster struct {
	Display             string     `json:"display,omitempty"`
	ID                  int        `json:"id,omitempty"`
	Name                string     `json:"name"`
	URL                 strfmt.URI `json:"url,omitempty"`
	VirtualmachineCount int        `json:"virtualmachine_count,omitempty"`
}

type NestedClusterGroup struct {
	ClusterCount int        `json:"cluster_count,omitempty"`
	Display      string     `json:"display,omitempty"`
	ID           int        `json:"id,omitempty"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug"`
	URL          strfmt.URI `json:"url,omitempty"`
}

type NestedClusterType struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ClusterCount int    `json:"cluster_count"`
}
type ClusterStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Cluster struct {
	Comments            string             `json:"comments,omitempty"`
	Created             strfmt.DateTime    `json:"created,omitempty"`
	CustomFields        interface{}        `json:"custom_fields,omitempty"`
	Description         string             `json:"description,omitempty"`
	DeviceCount         int                `json:"device_count,omitempty"`
	Display             string             `json:"display,omitempty"`
	Group               NestedClusterGroup `json:"group,omitempty"`
	ID                  int                `json:"id,omitempty"`
	LastUpdated         strfmt.DateTime    `json:"last_updated,omitempty"`
	Name                string             `json:"name"`
	Site                NestedSite         `json:"site,omitempty"`
	Status              ClusterStatus      `json:"status,omitempty"`
	Tags                []NestedTag        `json:"tags,omitempty"`
	Tenant              NestedTenant       `json:"tenant,omitempty"`
	Type                NestedClusterType  `json:"type"`
	URL                 strfmt.URI         `json:"url,omitempty"`
	VirtualmachineCount int                `json:"virtualmachine_count,omitempty"`
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
