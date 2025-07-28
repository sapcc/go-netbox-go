// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/sapcc/go-netbox-go/common"
)

type NestedCable struct {
	ID      int        `json:"id,omitempty"`
	URL     strfmt.URI `json:"url,omitempty"`
	Display string     `json:"display,omitempty"`
	Label   string     `json:"label,omitempty"`
}

type CableLengthUnit struct {
	Label *string `json:"label"`
	Value *string `json:"value"`
}

type Cable struct {
	Color              string          `json:"color,omitempty"`
	Comments           string          `json:"comments,omitempty"`
	Created            strfmt.DateTime `json:"created,omitempty"`
	CustomFields       interface{}     `json:"custom_fields,omitempty"`
	Description        string          `json:"description,omitempty"`
	Display            string          `json:"display,omitempty"`
	ID                 int             `json:"id,omitempty"`
	Label              string          `json:"label,omitempty"`
	LastUpdated        strfmt.DateTime `json:"last_updated,omitempty"`
	Length             float64         `json:"length,omitempty"`
	LengthUnit         CableLengthUnit `json:"length_unit,omitempty"`
	Status             CableStatus     `json:"status,omitempty"`
	Tags               []NestedTag     `json:"tags,omitempty"`
	Tenant             NestedTenant    `json:"tenant,omitempty"`
	Type               string          `json:"type,omitempty"`
	URL                strfmt.URI      `json:"url,omitempty"`
	Aterminations      []Termination
	Bterminations      []Termination
	AssignedObjectType string
}

type WriteableCable struct {
	Aterminations []Termination   `json:"a_terminations,omitempty"`
	Bterminations []Termination   `json:"b_terminations,omitempty"`
	Color         string          `json:"color,omitempty"`
	Comments      string          `json:"comments,omitempty"`
	Created       strfmt.DateTime `json:"created,omitempty"`
	CustomFields  interface{}     `json:"custom_fields,omitempty"`
	Description   string          `json:"description,omitempty"`
	Display       string          `json:"display,omitempty"`
	ID            int64           `json:"id,omitempty"`
	Label         string          `json:"label,omitempty"`
	LastUpdated   strfmt.DateTime `json:"last_updated,omitempty"`
	Length        float64         `json:"length,omitempty"`
	LengthUnit    string          `json:"length_unit,omitempty"`
	Status        string          `json:"status,omitempty"`
	Tags          []NestedTag     `json:"tags,omitempty"`
	Tenant        int64           `json:"tenant,omitempty"`
	Type          string          `json:"type,omitempty"`
	URL           strfmt.URI      `json:"url,omitempty"`
}

type LengthUnit struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type CableStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Termination struct {
	Object     TerminationObject `json:"object,omitempty"`
	ObjectID   int               `json:"object_id"`
	ObjectType string            `json:"object_type"`
}

type TerminationObject struct {
	ID     int          `json:"id"`
	URL    string       `json:"url"`
	Device NestedDevice `json:"device"`
	Name   string       `json:"name"`
	Cable  NestedCable  `json:"cable"`
}

type ListCablesRequest struct {
	common.ListParams
	CableID   int
	CableType string
}

type ListCablesResponse struct {
	common.ReturnValues
	Results []Cable `json:"results"`
}

func (cable *Cable) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	var aTerminations []Termination
	if err := json.Unmarshal(tmp["a_terminations"], &aTerminations); err != nil {
		return err
	}

	var aterm []Termination
	for term := range aTerminations {
		if aTerminations[term].ObjectType != "dcim.interface" {
			return fmt.Errorf("unknown assigned object type %v", aTerminations[term].ObjectType)
		}
		aterm = append(aterm, aTerminations[term])
	}
	cable.Aterminations = aterm

	var bTerminations []Termination
	if err := json.Unmarshal(tmp["b_terminations"], &bTerminations); err != nil {
		return err
	}

	var bterm []Termination
	for term := range bTerminations {
		if bTerminations[term].ObjectType != "dcim.interface" {
			return fmt.Errorf("unknown assigned object type %v", bTerminations[term].ObjectType)
		}
		bterm = append(bterm, bTerminations[term])
	}
	cable.Bterminations = bterm

	cable.AssignedObjectType = "dcim.interface"

	var cabletype string
	if err := json.Unmarshal(tmp["type"], &cabletype); err != nil {
		return err
	}
	cable.Type = cabletype

	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	cable.ID = id

	var url strfmt.URI
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	cable.URL = url

	var display string
	if err := json.Unmarshal(tmp["display"], &display); err != nil {
		return err
	}
	cable.Display = display

	var label string
	if err := json.Unmarshal(tmp["label"], &label); err != nil {
		return err
	}
	cable.Label = label

	var status CableStatus
	if err := json.Unmarshal(tmp["status"], &status); err != nil {
		return err
	}
	cable.Status = status

	var color string
	if err := json.Unmarshal(tmp["color"], &color); err != nil {
		return err
	}
	cable.Color = color

	var length float64
	if err := json.Unmarshal(tmp["length"], &length); err != nil {
		return err
	}
	cable.Length = length

	var lengthUnit CableLengthUnit
	if err := json.Unmarshal(tmp["length_unit"], &lengthUnit); err != nil {
		return err
	}
	cable.LengthUnit = lengthUnit
	return nil
}
