package ipam

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListPrefixes(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListPrefixes", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListPrefixesRequest{}
	res, err := client.ListPrefixes(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	opts.Role = "cc-transit"
	res, err = client.ListPrefixes(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	opts.Region = "qa-de-1"
	res, err = client.ListPrefixes(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_ListAvailableIps(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListAvailableIps", vcrConf)
	client.HttpClient = vcr.Client
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
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateDeletePrefix", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.CreatePrefix(wPre)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Id)
	err = client.DeletePrefix(res.Id)
	if err != nil {
		t.Fatal(err)
	}
}