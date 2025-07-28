// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ipam_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/ipam"
	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_GetVlan(t *testing.T) {
	client, err := ipam.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("GetVlan", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetVlan(318)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.ID)
}

func TestClient_ListVlans(t *testing.T) {
	client, err := ipam.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListVlans", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListVlanRequest{}
	opts.ID = 1661
	res, err := client.ListVlans(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	t.Log(res.Results[0].Name)
	t.Log(res.Results[0].Role)
	assert.NotEqual(t, 0, res.Count)
}
