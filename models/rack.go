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

type NestedRack struct {
	DeviceCount int        `json:"device_count,omitempty"`
	Display     string     `json:"display,omitempty"`
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name"`
	URL         strfmt.URI `json:"url,omitempty"`
}

type Rack struct {
	ID             int         `json:"id"`
	URL            string      `json:"url"`
	Name           string      `json:"name"`
	FacilityID     string      `json:"facility_id"`
	DisplayName    string      `json:"display_name"`
	Site           Site        `json:"site"`
	Group          interface{} `json:"group"`
	Tenant         Tenant      `json:"tenant"`
	Status         interface{} `json:"status"`
	Role           RackRole    `json:"role"`
	Serial         string      `json:"serial"`
	AssetTag       string      `json:"asset_tag"`
	Type           interface{} `json:"type"`
	Width          interface{} `json:"width"`
	UHeight        int         `json:"u_height"`
	DescUnits      bool        `json:"desc_units"`
	OuterWidth     int         `json:"outer_width"`
	OuterDepth     int         `json:"outer_depth"`
	OuterUnit      interface{} `json:"outer_unit"`
	Comments       string      `json:"comments"`
	Tags           interface{} `json:"tags"`
	CustomFields   interface{} `json:"custom_fields"`
	Created        string      `json:"created"`
	LastUpdated    string      `json:"last_updated"`
	DeviceCount    int         `json:"device_count"`
	PowerfeedCount int         `json:"powerfeed_count"`
}

type RackRole struct {
	ID           int         `json:"id"`
	URL          string      `json:"url"`
	Display      string      `json:"display"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Color        string      `json:"color"`
	Description  string      `json:"description"`
	CustomFields interface{} `json:"custom_fields"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
	RackCount    int         `json:"rack_count"`
}

type ListRacksRequest struct {
	common.ListParams
}

type ListRacksResponse struct {
	common.ReturnValues
	Results []Rack `json:"results"`
}
