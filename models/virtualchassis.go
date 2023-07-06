package models

import "github.com/go-openapi/strfmt"

type VirtualChassis struct {
	Comments     string          `json:"comments,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Description  string          `json:"description,omitempty"`
	Display      string          `json:"display,omitempty"`
	Domain       string          `json:"domain,omitempty"`
	Id           int             `json:"id,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	Master       NestedDevice    `json:"master,omitempty"`
	MemberCount  int             `json:"member_count,omitempty"`
	Name         string          `json:"name"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	URL          strfmt.URI      `json:"url,omitempty"`
}

type NestedVirtualChassis struct {
	Display     string       `json:"display,omitempty"`
	Id          int          `json:"id,omitempty"`
	Master      NestedDevice `json:"master"`
	MemberCount int          `json:"member_count,omitempty"`
	Name        string       `json:"name"`
	URL         strfmt.URI   `json:"url,omitempty"`
}
