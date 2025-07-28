// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package tenancy_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/sapcc/go-netbox-go/tenancy"
)

func TestClient_GetTenant(t *testing.T) {
	client, err := tenancy.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListTenants", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetTenant(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	// assert.NotEqual(t, 0, res.Count)
}
func TestClient_ListTenants(t *testing.T) {
	client, err := tenancy.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListTenants", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListTenantsRequest{}
	res, err := client.ListTenants(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}
