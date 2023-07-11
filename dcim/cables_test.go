package dcim

import (
	"os"
	"testing"

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

// func TestClient_ListCables(t *testing.T) {
// 	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	vcrConf := &govcr.VCRConfig{}
// 	vcrConf.Client = client.HttpClient
// 	vcr := govcr.NewVCR("ListCables", vcrConf)
// 	client.HttpClient = vcr.Client
// 	opts := models.ListCablesRequest{}
// 	opts.CableType = "dcim.interface"
// 	res, err := client.ListCables(opts)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(res)
// 	assert.NotEqual(t, 0, res.Count)
// }
