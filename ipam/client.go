// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
)

const basePath = "/api/ipam/"

type NetboxAPI interface {
	common.HTTPConnectable

	// ip-addresses
	ListIPAddresses(opts models.ListIPAddressesRequest) (*models.ListIPAddressesResponse, error)
	GetIPAdress(id int) (*models.IPAddress, error)
	CreateIPAddress(ip models.WriteableIPAddress) (*models.IPAddress, error)
	UpdateIPAddress(address models.WriteableIPAddress) (*models.IPAddress, error)
	DeleteIPAddress(id int) error

	// prefixes
	ListPrefixes(opts models.ListPrefixesRequest) (*models.ListPrefixesReponse, error)
	CreatePrefix(prefix models.WriteablePrefix) (*models.Prefix, error)
	ListAvailableIps(id int) ([]models.AvailableIP, error)
	CreateAvailablePrefix(id int, opts models.CreateAvailablePrefixRequest) (*models.Prefix, error)
	UpdatePrefix(prefix models.WriteablePrefix) (*models.Prefix, error)
	DeletePrefix(id int) error

	// roles
	ListRoles(opts models.ListRolesRequest) (*models.ListRolesResponse, error)

	// vlan
	ListVlans(opts models.ListVlanRequest) (*models.ListVlanResponse, error)
	GetVlan(id int) (*models.Vlan, error)

	// vrfs
	ListVRFs(opts models.ListVRFsRequest) (*models.ListVRFsResponse, error)
}

type Client struct {
	common.Client
}

// Deprecated: Use NewClient() function instead.
func New(baseURL, authToken string, insecureSkipVerify bool) (*Client, error) {
	client := &Client{}
	return client, common.Initialize(client, baseURL, authToken, insecureSkipVerify)
}

func NewClient(baseURL, authToken string, insecureSkipVerify bool) (NetboxAPI, error) {
	client := &Client{}
	return client, common.Initialize(client, baseURL, authToken, insecureSkipVerify)
}
