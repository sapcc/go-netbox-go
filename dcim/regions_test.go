// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package dcim_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListRegions(t *testing.T) {
	client, err := dcim.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListRegions", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListRegionsRequest{}
	opts.Slug = "qa-de-1"
	res, err := client.ListRegions(opts)
	if err != nil {
		t.Fatal(err)
	}
	for _, x := range res.Results {
		t.Log(x.Name)
	}

	assert.NotEqual(t, 0, res.Count)
}

func TestClient_GetRegion(t *testing.T) {
	client, err := dcim.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("GetRegion", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetRegion(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Name)
}
