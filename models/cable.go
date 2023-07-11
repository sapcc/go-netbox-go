package models

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type NestedCable struct {
	Id      int        `json:"id,omitempty"`
	Url     strfmt.URI `json:"url,omitempty"`
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
	Id                 int64           `json:"id,omitempty"`
	Label              string          `json:"label,omitempty"`
	LastUpdated        strfmt.DateTime `json:"last_updated,omitempty"`
	Length             float64         `json:"length,omitempty"`
	LengthUnit         CableLengthUnit `json:"length_unit,omitempty"`
	Status             CableStatus     `json:"status,omitempty"`
	Tags               []NestedTag     `json:"tags,omitempty"`
	Tenant             NestedTenant    `json:"tenant,omitempty"`
	Type               string          `json:"type,omitempty"`
	Url                strfmt.URI      `json:"url,omitempty"`
	Aterminations      []NestedInterface
	Bterminations      []NestedInterface
	AssignedObjectType string
}

type WriteableCable struct {
	Color        string          `json:"color,omitempty"`
	Comments     string          `json:"comments,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Description  string          `json:"description,omitempty"`
	Display      string          `json:"display,omitempty"`
	Id           int64           `json:"id,omitempty"`
	Label        string          `json:"label,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	Length       float64         `json:"length,omitempty"`
	LengthUnit   string          `json:"length_unit,omitempty"`
	Status       string          `json:"status,omitempty"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	Tenant       int64           `json:"tenant,omitempty"`
	Type         string          `json:"type,omitempty"`
	Url          strfmt.URI      `json:"url,omitempty"`
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
	Interface  NestedInterface `json:"object,omitempty"`
	ObjectId   int             `json:"object_id"`
	ObjectType string          `json:"object_type"`
}
type ListCablesRequest struct {
	common.ListParams
	CableId   int
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

	var a_terminations []Termination
	if err := json.Unmarshal(tmp["a_terminations"], &a_terminations); err != nil {
		return err
	}

	var aterm []NestedInterface
	for term := range a_terminations {
		if a_terminations[term].ObjectType != "dcim.interface" {
			return fmt.Errorf("unknown assigned object type %v", a_terminations[term].ObjectType)
		}
		aterm = append(aterm, a_terminations[term].Interface)
	}
	cable.Aterminations = aterm

	var b_terminations []Termination
	if err := json.Unmarshal(tmp["b_terminations"], &b_terminations); err != nil {
		return err
	}

	var bterm []NestedInterface
	for term := range b_terminations {
		if b_terminations[term].ObjectType != "dcim.interface" {
			return fmt.Errorf("unknown assigned object type %v", b_terminations[term].ObjectType)
		}
		bterm = append(bterm, b_terminations[term].Interface)
	}
	cable.Bterminations = bterm

	cable.AssignedObjectType = "dcim.interface"

	var cabletype string
	if err := json.Unmarshal(tmp["type"], &cabletype); err != nil {
		return err
	}
	cable.Type = cabletype

	var id int64
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	cable.Id = id

	var url strfmt.URI
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	cable.Url = url

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
