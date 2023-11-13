package dcim

import (
	"os"
	"testing"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListRegions(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListRegions", vcrConf)
	client.HttpClient = vcr.Client
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
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("GetRegion", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetRegion(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Name)
}
