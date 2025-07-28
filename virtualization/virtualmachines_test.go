// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package virtualization_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/ipam"
	"github.com/sapcc/go-netbox-go/models"
	"github.com/sapcc/go-netbox-go/tenancy"
	"github.com/sapcc/go-netbox-go/virtualization"
)

func TestClient_GetVirtualMachine(t *testing.T) {
	client, err := virtualization.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("GetVirtualMachine", vcrConf)
	client.SetHTTPClient(vcr.Client)
	vm, err := client.GetVirtualMachine(4773)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vm.Name)
}

func TestClient_ListVirtualMachines(t *testing.T) {
	client, err := virtualization.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListVirtualMachines", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListVirtualMachinesRequest{}
	opts.ID = 4773
	res, err := client.ListVirtualMachines(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0].Name)
	t.Log(res.Results[0].Cluster.Name)
	t.Log(res.Results[0].Status)
}

func TestClient_CreateDeleteVirtualMachine(t *testing.T) {
	client, err := virtualization.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	ipamClient, err := ipam.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	tenantClient, err := tenancy.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	dcimClient, err := dcim.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("CreateVirtualMachine", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListClusterRequest{}
	res, err := client.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	rOpts := models.ListRolesRequest{}
	roles, err := ipamClient.ListRoles(rOpts)
	if err != nil {
		t.Fatal(err)
	}
	tOpts := models.ListTenantsRequest{}
	tenants, err := tenantClient.ListTenants(tOpts)
	if err != nil {
		t.Fatal(err)
	}
	pOpts := models.ListPlatformsRequest{}
	platforms, err := dcimClient.ListPlatforms(pOpts)
	if err != nil {
		t.Fatal(err)
	}
	vm := models.WriteableVirtualMachine{
		Name:     "test.cc.qa-de-1.cloud.sap",
		Cluster:  res.Results[0].ID,
		Status:   "active",
		Role:     roles.Results[0].ID,
		Tenant:   tenants.Results[0].ID,
		Platform: platforms.Results[0].ID,
	}
	vm2, err := client.CreateVirtualMachine(vm)
	if err != nil {
		t.Fatal(err)
	}
	err = client.DeleteVirtualMachine(vm2.ID)
	if err != nil {
		t.Fatal(err)
	}
}
