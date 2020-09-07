package ipam

import (
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListIpAddresses(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"))
	if err != nil {
		t.Error(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListIpAddresses", vcrConf)
	client.HttpClient = vcr.Client
	opts := ListIpAddressesRequest{}
	res, err := client.ListIpAddresses(opts)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}