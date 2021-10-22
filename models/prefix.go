package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type Prefix struct {
	Id           int         `json:"id"`
	Url          string      `json:"url"`
	Family       interface{} `json:"family"`
	Prefix       string      `json:"prefix"`
	Site         Site        `json:"site"`
	Vrf          NestedVRF   `json:"vrf"`
	Tenant       Tenant      `json:"tenant"`
	Vlan         NestedVLAN  `json:"vlan"`
	Status       interface{} `json:"status"`
	Role         Role        `json:"role"`
	IsPool       bool        `json:"is_pool"`
	Description  string      `json:"description"`
	Tags         interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
}

type WriteablePrefix struct {
	Id           int         `json:"id,omitempty"`
	Url          string      `json:"url"`
	Prefix       string      `json:"prefix"`
	Site         int         `json:"site,omitempty"`
	Vrf          int         `json:"vrf,omitempty"`
	Tenant       int         `json:"tenant,omitempty"`
	Vlan         int         `json:"vlan,omitempty"`
	Status       string      `json:"status,omitempty"`
	Role         int         `json:"role,omitempty"`
	IsPool       bool        `json:"is_pool"`
	Description  string      `json:"description"`
	Tags         []NestedTag `json:"tags,omitempty"`
	CustomFields interface{} `json:"custom_fields,omitempty"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
}

type ListPrefixesRequest struct {
	common.ListParams
	Role   string
	Region string
	Tag    string
	VrfId  int
	Prefix string
	MaskLength int
	Status string
}

type ListPrefixesReponse struct {
	common.ReturnValues
	Results []Prefix `json:"results"`
}
