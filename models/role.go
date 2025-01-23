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

type Role struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	PrefixCount int    `json:"prefix_count"`
	VlanCount   int    `json:"vlan_count"`
}

type ListRolesRequest struct {
	common.ListParams
	Name string
	Slug string
}

type ListRolesResponse struct {
	common.ReturnValues
	Results []Role `json:"results"`
}
