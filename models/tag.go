package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type Tag struct {
	NestedTag
	Description string `json:"description"`
	TaggedItems int    `json:"tagged_items"`
}

type NestedTag struct {
	ID    int64      `json:"id,omitempty"`
	Url   strfmt.URI `json:"url,omitempty"`
	Name  string     `json:"name"`
	Slug  string     `json:"slug"`
	Color string     `json:"color,omitempty"`
}

type ListTagsRequest struct {
	common.ListParams
}

type ListTagsResponse struct {
	common.ReturnValues
	Results []Tag `json:"results"`
}
