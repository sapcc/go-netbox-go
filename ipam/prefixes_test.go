package ipam

import (
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListPrefixes(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"))
	if err != nil {
		t.Error(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListPrefixes", vcrConf)
	client.HttpClient = vcr.Client
	opts := ListPrefixesRequest{}
	res, err := client.ListPrefixes(opts)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}