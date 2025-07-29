// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package virtualization_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/common"
	"github.com/sapcc/go-netbox-go/models"
	"github.com/sapcc/go-netbox-go/virtualization"
)

func TestClient_ListClusters(t *testing.T) {
	client, err := virtualization.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListClusters", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListClusterRequest{}
	// opts.VmID = 1060
	opts.ID = 632
	res, err := client.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	t.Log(res.Results[0].Name)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_ListClusterByType(t *testing.T) {
	clint, err := virtualization.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = clint.HTTPClient()
	vcr := govcr.NewVCR("ListClusterByType", vcrConf)
	clint.SetHTTPClient(vcr.Client)
	opts := models.ListClusterRequest{
		Region: "ap-ae-1",
		Type:   "cc-k8s-controlplane-swift",
	}
	res, err := clint.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0].Name)
}

func TestClient_ListClusterByName(t *testing.T) {
	clint, err := virtualization.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = clint.HTTPClient()
	vcr := govcr.NewVCR("ListClusterByType", vcrConf)
	clint.SetHTTPClient(vcr.Client)
	opts := models.ListClusterRequest{
		ListParams: common.ListParams{
			Name: "ap-ae-1a-controlplane",
		},
	}
	res, err := clint.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0].Name)
}
