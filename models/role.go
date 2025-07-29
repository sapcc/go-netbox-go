// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

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
