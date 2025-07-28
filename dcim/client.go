// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package dcim

import (
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
)

const basePath = "/api/dcim/"

type NetboxAPI interface {
	common.HTTPConnectable

	// cables
	GetCable(id int) (*models.Cable, error)
	CreateCable(cable models.WriteableCable) (*models.Cable, error)
	DeleteCable(id int) error
	UpdateCable(cable models.WriteableCable) (*models.Cable, error)

	// device-roles
	ListDeviceRoles(opts models.ListDeviceRolesRequest) (*models.ListDeviceRolesResponse, error)

	// device-types
	ListDeviceTypes(opts models.ListDeviceTypesRequest) (*models.ListDeviceTypesResponse, error)

	// devices
	ListDevices(opts models.ListDevicesRequest) (*models.ListDevicesResponse, error)
	GetDevice(id int) (*models.Device, error)
	GetDeviceWithContext(id int) (*models.Device, error)
	CreateDevice(dev models.WritableDeviceWithConfigContext) (*models.Device, error)
	DeleteDevice(id int) error
	UpdateDevice(dev models.WritableDeviceWithConfigContext) (*models.Device, error)

	// interfaces
	ListInterfaces(opts models.ListInterfacesRequest) (*models.ListInterfacesResponse, error)
	UpdateInterface(interf models.WritableInterface, id int) (*models.Interface, error)
	CreateInterface(interf models.WritableInterface) (*models.Interface, error)
	DeleteInterface(id int) error

	// platforms
	ListPlatforms(opts models.ListPlatformsRequest) (*models.ListPlatformsResponse, error)

	// racks
	ListRacks(opts models.ListRacksRequest) (*models.ListRacksResponse, error)

	// regions
	ListRegions(opts models.ListRegionsRequest) (*models.ListRegionsResponse, error)
	GetRegion(id int) (*models.Region, error)

	// site groups
	ListSiteGroups(opts models.ListSiteGroupsRequest) (*models.ListSiteGroupsResponse, error)
	GetSiteGroup(id int) (*models.SiteGroup, error)

	// sites
	ListSites(opts models.ListSitesRequest) (*models.ListSitesResponse, error)
	GetSite(id int) (*models.Site, error)
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
