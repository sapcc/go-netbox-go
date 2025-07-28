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

func TestClient_ListVRFs(t *testing.T) {
	client, err := ipam.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListVRFs", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListVRFsRequest{}
	res, err := client.ListVRFs(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
}
