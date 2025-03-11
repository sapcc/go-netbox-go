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

package dcim

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListRegions(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
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
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("GetRegion", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetRegion(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Name)
}
