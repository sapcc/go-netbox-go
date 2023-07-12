package models

import "github.com/go-openapi/strfmt"

type NestedRegion struct {
	Depth     int        `json:"_depth,omitempty"`
	Display   string     `json:"display,omitempty"`
	Id        int        `json:"id,omitempty"`
	Name      string     `json:"name"`
	SiteCount int        `json:"site_count,omitempty"`
	Slug      string     `json:"slug"`
	Url       strfmt.URI `json:"url,omitempty"`
}

type Region struct {
	Depth        int             `json:"_depth,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Description  string          `json:"description,omitempty"`
	Display      string          `json:"display,omitempty"`
	Id           int             `json:"id,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	Name         string          `json:"name"`
	Parent       NestedRegion    `json:"parent,omitempty"`
	SiteCount    int             `json:"site_count,omitempty"`
	Slug         string          `json:"slug"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	Url          strfmt.URI      `json:"url,omitempty"`
}
