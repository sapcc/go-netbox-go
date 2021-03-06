package models

import (
	"github.com/sapcc/go-netbox-go/common"
)

type Rack struct {
	Id          int            `json:"id"`
	Url         string         `json:"url"`
	Name        string         `json:"name"`
	FacilityId  string         `json:"facility_id"`
	DisplayName string         `json:"display_name"`
	Site        Site           `json:"site"`
	Group       interface{}    `json:"group"`
	Tenant      Tenant `json:"tenant"`
	Status      interface{}    `json:"status"`
	Role        interface{}    `json:"role"`
	Serial      string         `json:"serial"`
	AssetTag    string         `json:"asset_tag"`
	Type        interface{}    `json:"type"`
	Width       interface{}    `json:width`
	UHeight     int            `json:"u_height"`
	DescUnits   bool           `json:"desc_units"`
	OuterWidth int `json:"outer_width"`
	OuterDepth int `json:"outer_depth"`
	OuterUnit interface{} `json:"outer_unit"`
	Comments string `json:"comments"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	DeviceCount int `json:"device_count"`
	PowerfeedCount int `json:"powerfeed_count"`
}

type ListRacksRequest struct {
	common.ListParams
}

type ListRacksResponse struct {
	common.ReturnValues
	Results []Rack `json:"results"`
}