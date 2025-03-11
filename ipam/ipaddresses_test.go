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

package ipam

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_GetIpAddress(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("GetIpAddresses", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetIPAdress(41797)
	t.Log(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
func TestClient_ListIpAddresses(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("ListIpAddresses", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListIPAddressesRequest{}
	res, err := client.ListIPAddresses(opts)
	t.Log(res)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
	opts.Role = "vip"
	res, err = client.ListIPAddresses(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_CreateDeleteIpAddress(t *testing.T) {
	wIP := models.WriteableIPAddress{}
	wIP.Address = "199.199.199.199/32"
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("CreateDeleteIpAddress", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.CreateIPAddress(wIP)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.ID)
	err = client.DeleteIPAddress(res.ID)
	if err != nil {
		t.Fatal(err)
	}
}
