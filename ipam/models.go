package ipam

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/tenancy"
)

// TODO: Refactor out NestedIP for VM
type IpAddress struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Family interface{} `json:"family"`
	Address string `json:"address"`
	Vrf interface{} `json:"vrf"`
	Tenant tenancy.Tenant `json:"tenant"`
	Status interface{} `json:"status"`
	Role interface{} `json:"role"`
	NatInside interface{} `json:"nat_inside"`
	NatOutside interface{} `json:"nat_outside"`
	DnsName string `json:"dns_name"`
	Description string `json:"description"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
	AssignedInterface dcim.Interface
}

type ListIpAddressesRequest struct {
	common.ListParams
	InterfaceId int
	DeviceId int
}

type ListIpAddressesResponse struct {
	common.ReturnValues
	Results []IpAddress `json:"results"`
}

func (ip *IpAddress) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	var s string
	if err := json.Unmarshal(tmp["assigned_object_type"], &s); err != nil {
		return err
	}
	switch s {
	case "dcim.interface":
		var inter dcim.Interface
		if err := json.Unmarshal(tmp["assigned_object"], &inter); err != nil {
			return err
		}
		ip.AssignedInterface = inter
	default:
		_ = fmt.Errorf("unknown assigned object type %v", s)
	}
	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	ip.Id = id
	var url string
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	ip.Url = url
	var addr string
	if err := json.Unmarshal(tmp["address"], &addr); err != nil {
		return err
	}
	ip.Address = addr
	var tenant tenancy.Tenant
	if err := json.Unmarshal(tmp["tenant"], &tenant); err != nil {
		return err
	}
	ip.Tenant = tenant
	var dnsName string
	if err := json.Unmarshal(tmp["dns_name"], &dnsName); err != nil {
		return err
	}
	ip.DnsName = dnsName
	var descr string
	if err := json.Unmarshal(tmp["description"],&descr); err != nil {
		return err
	}
	ip.Description = descr
	var created string
	if err := json.Unmarshal(tmp["created"], &created); err != nil {
		return err
	}
	ip.Created = created
	var lastUpdated string
	if err := json.Unmarshal(tmp["last_updated"], &lastUpdated); err != nil {
		return err
	}
	ip.LastUpdated = lastUpdated
	return nil
}

type Role struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Weight int `json:"weight"`
	Description string `json:"description"`
	PrefixCount int `json:"prefix_count"`
	VlanCount int `json:"vlan_count"`
}

type ListRolesRequest struct {
	common.ListParams
}

type ListRolesResponse struct {
	common.ReturnValues
	Results []Role `json:"results"`
}


type Prefix struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Family interface{} `json:"family"`
	Prefix string `json:"prefix"`
	Site dcim.Site `json:"site"`
	Vrf interface{} `json:"vrf"`
	Tenant tenancy.Tenant `json:"tenant"`
	Vlan interface{} `json:"vlan"`
	Status interface{} `json:"status"`
	Role Role `json:"role"`
	IsPool bool `json:"is_pool"`
	Description string `json:"description"`
	Tags interface{} `json:"tags"`
	CustomFields interface{} `json:"custom_fields"`
	Created string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

type ListPrefixesRequest struct {
	common.ListParams
}

type ListPrefixesReponse struct {
	common.ReturnValues
	Results []Prefix `json:"results"`
}