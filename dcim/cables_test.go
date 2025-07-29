// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package dcim_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/seborama/govcr"

	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_GetCable(t *testing.T) {
	client, err := dcim.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("GetCable", vcrConf)
	client.SetHTTPClient(vcr.Client)
	res, err := client.GetCable(2026) // 2026
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cable ID:", res.ID)
}

func TestClient_CreateDeleteCable(t *testing.T) {
	client, err := dcim.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("CreateDeleteCable", vcrConf)
	client.SetHTTPClient(vcr.Client)

	// test cable creation/deletion against test server TestServer-Go-Netbox-Go in Netbox
	wCable := models.WriteableCable{}
	wCable.Type = "cat6"
	aterm := models.Termination{}
	aterm.ObjectID = 1040790
	aterm.ObjectType = "dcim.interface"
	bterm := models.Termination{}
	bterm.ObjectID = 1040791
	bterm.ObjectType = "dcim.interface"

	wCable.Aterminations = append(wCable.Aterminations, aterm)
	wCable.Bterminations = append(wCable.Bterminations, bterm)

	wCable.Label = "GNG Test Cable"

	wTag := models.NestedTag{}
	wTag.Name = "apod"
	wTag.Slug = "apod"

	wCable.Tags = append(wCable.Tags, wTag)

	fmt.Println("--->", wCable.Tags)

	cab, err := client.CreateCable(wCable)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cab)

	err = client.DeleteCable(cab.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateCable(t *testing.T) {
	client, err := dcim.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("CreateDeleteCable", vcrConf)
	client.SetHTTPClient(vcr.Client)

	wCable := models.WriteableCable{}
	wCable.ID = 117102
	aterm := models.Termination{}
	aterm.ObjectID = 598436
	aterm.ObjectType = "dcim.interface"

	wCable.Aterminations = append(wCable.Aterminations, aterm)

	cab, err := client.UpdateCable(wCable)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cab)
}
