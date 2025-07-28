// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package virtualization

import (
	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
)

const basePath = "/api/virtualization/"

type NetboxAPI interface {
	common.HTTPConnectable

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

// Deprecated: Use NewClient() function instead.
func New(baseURL, authToken string, insecureSkipVerify bool) (*Client, error) {
	client := &Client{}
	return client, common.Initialize(client, baseURL, authToken, insecureSkipVerify)
}

func NewClient(baseURL, authToken string, insecureSkipVerify bool) (NetboxAPI, error) {
	client := &Client{}
	return client, common.Initialize(client, baseURL, authToken, insecureSkipVerify)
}
