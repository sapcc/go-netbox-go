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
