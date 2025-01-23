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

import "github.com/go-openapi/strfmt"

type VirtualChassis struct {
	Comments     string          `json:"comments,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Description  string          `json:"description,omitempty"`
	Display      string          `json:"display,omitempty"`
	Domain       string          `json:"domain,omitempty"`
	ID           int             `json:"id,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	Master       NestedDevice    `json:"master,omitempty"`
	MemberCount  int             `json:"member_count,omitempty"`
	Name         string          `json:"name"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	URL          strfmt.URI      `json:"url,omitempty"`
}

type NestedVirtualChassis struct {
	Display     string       `json:"display,omitempty"`
	ID          int          `json:"id,omitempty"`
	Master      NestedDevice `json:"master"`
	MemberCount int          `json:"member_count,omitempty"`
	Name        string       `json:"name"`
	URL         strfmt.URI   `json:"url,omitempty"`
}
