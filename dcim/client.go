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

package dcim

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
)

const basePath = "/api/dcim/"

type API interface {
	common.API

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

func New(baseURL, authToken string, insecureSkipVerify bool) (API, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify, // #nosec
	}
	res := &Client{}
	res.BaseURL = *u
	res.HTTPClient = &http.Client{
		Transport: tr,
	}
	res.AuthToken = authToken
	return res, nil
}
