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

func TestClient_CreateDeleteVLANVMInterface(t *testing.T) {
	client, err := NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("CreateVLANVMInterface", vcrConf)
	client.SetHTTPClient(vcr.Client)
	vlans := []int{1678, 1679, 1680}
	vmi := models.WritableVMInterface{
		VirtualMachine: 740,
		Name:           "test-vlan-interface",
		Enabled:        true,
		MTU:            9000,
		Description:    "this is a test vlan interface",
		TaggedVlans:    vlans,
	}
	vmi2, err := client.CreateVMInterface(vmi)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vmi2)
	assert.Equal(t, 3, len(vmi2.TaggedVlans))
	err = client.DeleteVMInterface(vmi2.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateDeleteTaggedVMInterface(t *testing.T) {
	client, err := NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("CreateTaggedVMInterface", vcrConf)
	client.SetHTTPClient(vcr.Client)
	tag1 := models.NestedTag{
		Name: "CC-APOD",
		Slug: "cc-apod",
	}
	tag2 := models.NestedTag{
		Name: "CC-VPOD",
		Slug: "cc-vpod",
	}
	tags := []models.NestedTag{
		tag1, tag2,
	}
	vmi := models.WritableVMInterface{
		VirtualMachine: 1107,
		Name:           "test-tagged-interface",
		Enabled:        true,
		MTU:            9000,
		Description:    "this is a test tagged interface",
		Tags:           tags,
	}
	vmi2, err := client.CreateVMInterface(vmi)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vmi2)
	assert.Equal(t, 2, len(vmi2.Tags))
	err = client.DeleteVMInterface(vmi2.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateDeleteVMInterface(t *testing.T) {
	client, err := NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("CreateVMInterface", vcrConf)
	client.SetHTTPClient(vcr.Client)
	vmi := models.WritableVMInterface{
		VirtualMachine: 1107,
		Name:           "test-interface",
		Enabled:        true,
		MTU:            9000,
		Description:    "this is a test interface",
	}
	vmi2, err := client.CreateVMInterface(vmi)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "test-interface", vmi2.Name)
	mtu := *vmi2.MTU
	en := *vmi2.Enabled
	assert.Equal(t, 9000, mtu)
	assert.Equal(t, true, en)
	assert.Equal(t, 1107, vmi2.VirtualMachine.ID)
	t.Log(vmi2)
	err = client.DeleteVMInterface(vmi2.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListVMInterfaces(t *testing.T) {
	client, err := NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("ListVMInterfaces", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListVMInterfacesRequest{}
	// opts.VmID = 804
	// opts.ID = 971
	res, err := client.ListVMInterfaces(opts)
	// t.Log(res)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	t.Log(res.Results[0].Name)
	assert.NotEqual(t, 0, res.Count)
}
