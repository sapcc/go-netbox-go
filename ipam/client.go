/*
 *   Copyright 2020 SAP SE
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

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
