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

package dcim_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_GetDevice(t *testing.T) {
	client, err := dcim.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("GetDevice", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetDevice(27572)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Name)
}

func TestClient_GetDeviceWithContext(t *testing.T) {
	client, err := dcim.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("GetDevice", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetDeviceWithContext(27572)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Name)
}
func TestClient_ListDevices(t *testing.T) {
	client, err := dcim.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("ListDevices", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListDevicesRequest{}
	// opts.ID = 12509
	res, err := client.ListDevices(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Results[0].Name)
}

func TestClient_ListDevicesByCluster(t *testing.T) {
	client, err := dcim.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("ListDevicesByCluster", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListDevicesRequest{
		ClusterID: 831,
	}
	res, err := client.ListDevices(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0])
}

func TestClient_CreateDeleteDevice(t *testing.T) {
	client, err := dcim.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("CreateDeleteDevice", vcrConf)
	client.SetHTTPClient(vcr.Client)

	wDev := models.WritableDeviceWithConfigContext{}
	wDev.Name = "GNG Test Device"
	wDev.DeviceRole = 8   // Server
	wDev.DeviceType = 132 // PowerEdge R640
	wDev.Site = 21        // ap-sa-2a
	wDev.Status = "staged"
	wDev.Rack = 1126 // AP003-01 (sa-2-ET14-Cage2-AY9-4)
	wDev.Tenant = 1  // Converged Cloud
	dev, err := client.CreateDevice(wDev)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dev)

	err = client.DeleteDevice(dev.ID)
	if err != nil {
		t.Fatal(err)
	}
}
