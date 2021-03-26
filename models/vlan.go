package models

import "github.com/sapcc/go-netbox-go/common"

type NestedVLAN struct {
	Id int `json:"id"`
	Url string `json:"url"`
	VId int `json:"vid"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
}

type Vlan struct {
	NestedVLAN
	Group interface{} `json:"group"`
	Description string `json:"description"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	PrefixCount int `json:"prefix_count"`
	Role interface{} `json:"role"`
	Site NestedSite `json:"site"`
	Status interface{} `json:"status"`
	Tenant Tenant `json:"tenant"`
}

type ListVlanRequest struct {
	common.ListParams
	Group string `json:"group"`
}

type ListVlanResponse struct {
	common.ReturnValues
	Results []Vlan `json:"results"`
}
