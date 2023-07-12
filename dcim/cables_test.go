package dcim

import (
	"os"
	"testing"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
)

func TestClient_GetCable(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("GetCable", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetCable(2026) //2026
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Cable ID:", res.Id)
}

func TestClient_CreateDeleteCable(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateDeleteCable", vcrConf)
	client.HttpClient = vcr.Client

	wCable := models.WriteableCable{}
	wCable.Type = "cat6"
	aterm := models.Termination{}
	aterm.ObjectId = 376515
	aterm.ObjectType = "dcim.interface"
	bterm := models.Termination{}
	bterm.ObjectId = 376517
	bterm.ObjectType = "dcim.interface"

	wCable.Aterminations = append(wCable.Aterminations, aterm)
	wCable.Bterminations = append(wCable.Bterminations, bterm)

	wCable.Label = "GNG Test Cable"

	wTag := models.NestedTag{}
	wTag.Name = "apod"
	wTag.Slug = "apod"

	wCable.Tags = append(wCable.Tags, wTag)

	cab, err := client.CreateCable(wCable)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cab)

	err = client.DeleteCable(cab.Id)
	if err != nil {
		t.Fatal(err)
	}
}
