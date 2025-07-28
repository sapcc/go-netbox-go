// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type Tag struct {
	NestedTag
	Description string `json:"description"`
	TaggedItems int    `json:"tagged_items"`
}

type NestedTag struct {
	ID int `json:"id,omitempty"`
	// URL   strfmt.URI `json:"url,omitempty"` //commented this aout as we get error when updating devices with tag: "Error: unexpected return code of 400: {"tags":[["Cannot resolve keyword 'url' into field. Choices are:"
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Color string `json:"color,omitempty"`
}

type ListTagsRequest struct {
	common.ListParams
}

type ListTagsResponse struct {
	common.ReturnValues
	Results []Tag `json:"results"`
}
