package models

import "github.com/sapcc/go-netbox-go/common"

type Tag struct {
	NestedTag
	Description string	`json:"description"`
	TaggedItems	int		`json:"tagged_items"`
}

type NestedTag struct {
	Id 		int		`json:"id"'`
	Url 	string	`json:"url"`
	Name	string	`json:"name"`
	Slug 	string	`json:"slug"`
	Color 	string	`json:"color"`
}

type ListTagsRequest struct {
	common.ListParams
}

type ListTagsResponse struct {
	common.ReturnValues
	Results []Tag `json:"results"`
}

