package models

import "github.com/sapcc/go-netbox-go/common"

type TenantGroup struct {
	NestedTenantGroup
	Description  string             `json:"description"`
	Comments     string             `json:"comments"`
	CustomFields interface{}        `json:"custom_fields"`
	Created      string             `json:"created"`
	LastUpdated  string             `json:"last_updated"`
	Name         string             `json:"name"`
	Parent       *NestedTenantGroup `json:"parent,omitempty"`
	Slug         string             `json:"slug"`
	Tags         []*NestedTag       `json:"tags"`
	TenantCount  int                `json:"tenant_count,omitempty"`
	Url          string             `json:"url,omitempty"`
}

type NestedTenantGroup struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	TenantCount string `json:"tenant_count,omitempty"`
}

type ListTenantGroupsRequest struct {
	common.ListParams
}

type ListTenantGroupsResponse struct {
	common.ReturnValues
	Results []TenantGroup `json:"results"`
}
