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

type TenantGroup struct {
	NestedTenantGroup
	Description  string             `json:"description"`
	Comments     string             `json:"comments"`
	CustomFields interface{}        `json:"custom_fields"`
	Created      string             `json:"created"`
	LastUpdated  string             `json:"last_updated"`
	Name         string             `json:"name"`
	Parent       *NestedTenantGroup `json:"parent,omitempty"`
	Slug         string             `json:"slug"`
	Tags         []*NestedTag       `json:"tags"`
	TenantCount  int                `json:"tenant_count,omitempty"`
	URL          string             `json:"url,omitempty"`
}

type NestedTenantGroup struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	TenantCount string `json:"tenant_count,omitempty"`
}

type ListTenantGroupsRequest struct {
	common.ListParams
}

type ListTenantGroupsResponse struct {
	common.ReturnValues
	Results []TenantGroup `json:"results"`
}
