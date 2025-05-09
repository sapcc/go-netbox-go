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

type NestedVRF struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Rd          string `json:"rd"` // Route Distinguisher
	DisplayName string `json:"display_name"`
	PrefixCount int    `json:"prefix_count"`
}

type VRF struct {
	NestedVRF
	Tenant         NestedTenant `json:"tenant"`
	EnforceUnique  bool         `json:"enforce_unique"`
	Description    string       `json:"description"`
	Tags           []NestedTag  `json:"tags"`
	CustomFields   interface{}  `json:"custom_fields"`
	Created        string       `json:"created"`
	LastUpdated    string       `json:"last_updated"`
	IPAddressCount int          `json:"ipaddress_count"`
	PrefixCount    int          `json:"prefix_count"`
}

type ListVRFsRequest struct {
	common.ListParams
	Name string
}

type ListVRFsResponse struct {
	common.ReturnValues
	Results []VRF `json:"results"`
}
