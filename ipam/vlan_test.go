package ipam

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_GetVlan(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("GetVlan", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetVlan(318)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Id)
}

func TestClient_ListVlans(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListVlans", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListVlanRequest{}
	opts.Id = 1661
	res, err := client.ListVlans(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	t.Log(res.Results[0].Name)
	t.Log(res.Results[0].Role)
	assert.NotEqual(t, 0, res.Count)
}
