package ipam

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListPrefixes(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("ListPrefixes", vcrConf)
	client.HTTPClient = vcr.Client
	opts := models.ListPrefixesRequest{}
	res, err := client.ListPrefixes(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	opts.Role = "cc-transit"
	res, err = client.ListPrefixes(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	opts.Region = "ap-sa-1"
	res, err = client.ListPrefixes(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_ListAvailableIps(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("ListAvailableIps", vcrConf)
	client.HTTPClient = vcr.Client
	res, err := client.ListAvailableIps(299)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, len(res))
}

func TestClient_CreateDeletePrefix(t *testing.T) {
	wPre := models.WriteablePrefix{}
	wPre.Prefix = "10.0.0.0/8"
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("CreateDeletePrefix", vcrConf)
	client.HTTPClient = vcr.Client
	res, err := client.CreatePrefix(wPre)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.ID)
	err = client.DeletePrefix(res.ID)
	if err != nil {
		t.Fatal(err)
	}
}
