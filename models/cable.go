package models

import (
	"encoding/json"
	"fmt"
)

type NestedCable struct {
	Id      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Label   string `json:"label,omitempty"`
}

type Cable struct {
	NestedCable
	TerminationAType string      `json:"termination_a_type"`
	TerminationAId   int         `json:"termination_a_id"`
	TerminationA     string      `json:"termination_a"`
	TerminationBType string      `json:"termination_b_type"`
	TerminationBId   int         `json:"termination_b_id"`
	//TerminationB     string      `json:"termination_b"`
	TerminationBInterface NestedInterface
	TerminationBConsolePort NesterConsolePort
	TerminationBConsoleServerPort NestedConsoleServerPort
	Type             string      `json:"type"`
	Status           CableStatus `json:"status"`
	Color            string      `json:"color"`
	Length           int         `json:"length"`
	LengthUnit       LengthUnit  `json:"length_unit"`
	Tags             interface{} `json:"tags"`
	CustomFields     interface{} `json:"custom_fields"`
}

type WriteableCable struct {
	TerminationAType string      `json:"termination_a_type"`
	TerminationAId   int         `json:"termination_a_id"`
	TerminationBType string      `json:"termination_b_type"`
	TerminationBId   int         `json:"termination_b_id"`
	Type             string      `json:"type,omitempty"`
	Status           string      `json:"status,omitempty"`
	Label            string      `json:"label,omitempty"`
	Color            string      `json:"color,omitempty"`
	Length           int         `json:"length,omitempty"`
	LengthUnit       string      `json:"length_unit,omitempty"`
	Tags             interface{} `json:"tags"`
	CustomFields     interface{} `json:"custom_fields"`
}


type LengthUnit struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type CableStatus struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (cable *Cable) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	var termAType string
	if err := json.Unmarshal(tmp["termination_a_type"], &termAType); err != nil {
		fmt.Println("unable to unmarshal cable")
		fmt.Println(string(tmp["termination_a_type"]))
		return err
	}
	switch termAType {
	case "dcim.consoleport":
	case "dcim.consoleserverport":
	case "dcim.interface":
	case "":
	default:
		return fmt.Errorf("unknown termination type %s", termAType)

	}
	var termBType string
	if err := json.Unmarshal(tmp["termination_b_type"], &termBType); err != nil {
		fmt.Println("unable to unmarshal cable")
		fmt.Println(string(tmp["termination_b_type"]))
		return err
	}
	switch termBType {
	case "dcim.consoleport":
	case "dcim.consoleserverport":
	case "dcim.interface":
		var inter NestedInterface{}
		if err := json.Unmarshal(tmp["termination_b"], &inter); err != nil {
			fmt.Println("unable to unmarshal termination b")
			fmt.Println(string(tmp["termination_b"]))
			return err
		}
		cable.TerminationBInterface = inter
	case "":
	default:
		return fmt.Errorf("unknown termination type %s", termBType)

	}
	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	cable.Id = id
	var url string
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
	var typ string
	if err := json.Unmarshal(tmp["type"], &typ); err != nil {
		return err
	}
	cable.Type = typ
	var status CableStatus
	if err := json.Unmarshal(tmp["status"], &status); err != nil {
		return err
	}
	cable.Status = status
	var color string
	if err := json.Unmarshal(tmp["color"], &color); err != nil {
		return err
	}
	var length int
	if err := json.Unmarshal(tmp["length"], &length); err != nil {
		return err
	}
	cable.Length = length
	var lengthUnit LengthUnit
	if err := json.Unmarshal(tmp["length_unit"], &lengthUnit); err != nil {
		return err
	}
	cable.LengthUnit = lengthUnit
	return nil
}

