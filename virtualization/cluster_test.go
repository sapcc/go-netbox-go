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
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListClusters(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("ListClusters", vcrConf)
	client.HTTPClient = vcr.Client
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
	clint, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = clint.HTTPClient
	vcr := govcr.NewVCR("ListClusterByType", vcrConf)
	clint.HTTPClient = vcr.Client
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
