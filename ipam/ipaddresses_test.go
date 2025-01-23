package ipam

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListIpAddresses(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("ListIpAddresses", vcrConf)
	client.HTTPClient = vcr.Client
	opts := models.ListIPAddressesRequest{}
	res, err := client.ListIPAddresses(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	opts.Role = "vip"
	res, err = client.ListIPAddresses(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_CreateDeleteIpAddress(t *testing.T) {
	wIP := models.WriteableIPAddress{}
	wIP.Address = "199.199.199.199/32"
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("CreateDeleteIpAddress", vcrConf)
	client.HTTPClient = vcr.Client
	res, err := client.CreateIPAddress(wIP)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.ID)
	err = client.DeleteIPAddress(res.ID)
	if err != nil {
		t.Fatal(err)
	}
}
