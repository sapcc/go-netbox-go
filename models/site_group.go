package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type NestedSiteGroup struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type SiteGroup struct {
	Id           int         `json:"id"`
	Url          string      `json:"url"`
	Depth        int         `json:"_depth"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	Tags         interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created      string      `json:"created"`
	SiteCount    int         `json:"site_count"`
	LastUpdated  string      `json:"last_updated"`
}

type ListSiteGroupsRequest struct {
	common.ListParams
	Region string
}

type ListSiteGroupsResponse struct {
	common.ReturnValues
	Results []SiteGroup `json:"results"`
}
