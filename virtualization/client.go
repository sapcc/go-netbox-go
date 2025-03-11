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

package virtualization

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
)

const basePath = "/api/virtualization/"

type API interface {
	common.Base

	// cluster
	ListClusters(opts models.ListClusterRequest) (*models.ListClusterResponse, error)

	// virtual machines
	CreateVirtualMachine(vm models.WriteableVirtualMachine) (*models.VirtualMachine, error)
	UpdateVirtualMachine(vm models.WriteableVirtualMachine) (*models.VirtualMachine, error)
	DeleteVirtualMachine(id int) error
	GetVirtualMachine(id int) (*models.VirtualMachine, error)
	ListVirtualMachines(opts models.ListVirtualMachinesRequest) (*models.ListVirtualMachinesResponse, error)

	// vm interfaces
	CreateVMInterface(vmni models.WritableVMInterface) (*models.VMInterface, error)
	UpdateVMInterface(vmi models.WritableVMInterface) (*models.VMInterface, error)
	DeleteVMInterface(id int) error
	ListVMInterfaces(opts models.ListVMInterfacesRequest) (*models.ListVMInterfacesResponse, error)
	GetVMInterface(id int) (*models.VMInterface, error)
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
