package models

type NestedVLAN struct {
	Id int `json:"id"`
	Url string `json:"url"`
	VId int `json:"vid"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
}
