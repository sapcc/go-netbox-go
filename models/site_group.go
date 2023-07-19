package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/sapcc/go-netbox-go/common"
)

type NestedSiteGroup struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type SiteGroup struct {
	Created      strfmt.DateTime `json:"created"`
	CustomFields interface{}     `json:"custom_fields"`
	Depth        int             `json:"_depth"`
	Description  string          `json:"description"`
	Display      string          `json:"display,omitempty"`
	Id           int             `json:"id"`
	LastUpdated  strfmt.DateTime `json:"last_updated"`
	Name         string          `json:"name"`
	Parent       NestedSiteGroup `json:"parent,omitempty"`
	SiteCount    int             `json:"site_count"`
	Slug         string          `json:"slug"`
	Tags         []NestedTag     `json:"tags"`
	Url          strfmt.URI      `json:"url"`
}

type ListSiteGroupsRequest struct {
	common.ListParams
	Region string
}

type ListSiteGroupsResponse struct {
	common.ReturnValues
	Results []SiteGroup `json:"results"`
}
