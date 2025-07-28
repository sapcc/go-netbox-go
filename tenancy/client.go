// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package tenancy

import (
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
)

const basePath = "/api/tenancy/"

type NetboxAPI interface {
	common.HTTPConnectable

	// tenant group
	ListTenantGroups(opts models.ListTenantGroupsRequest) (*models.ListTenantGroupsResponse, error)

	// tenants
	GetTenant(id int) (*models.Tenant, error)
	ListTenants(opts models.ListTenantsRequest) (*models.ListTenantsResponse, error)
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
