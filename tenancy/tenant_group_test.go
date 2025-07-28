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

func TestClient_ListTenantGroups(t *testing.T) {
	client, err := tenancy.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListTenantGroups", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListTenantGroupsRequest{}
	res, err := client.ListTenantGroups(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}
