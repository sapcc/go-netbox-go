package ipam

import (
	"os"
	"testing"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListIpAddresses(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListIpAddresses", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListIpAddressesRequest{}
	res, err := client.ListIpAddresses(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	opts.Role = "vip"
	res, err = client.ListIpAddresses(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_CreateDeleteIpAddress(t *testing.T) {
	wIp := models.WriteableIpAddress{}
	wIp.Address = "199.199.199.199/32"
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateDeleteIpAddress", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.CreateIpAddress(wIp)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Id)
	err = client.DeleteIpAddress(res.Id)
	if err != nil {
		t.Fatal(err)
	}
}
