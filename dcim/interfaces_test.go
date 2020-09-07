package dcim

import (
	"github.com/stretchr/testify/assert"
	"github.com/seborama/govcr"
	"os"
	"testing"
)

func TestClient_ListInterfaces(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"))
	if err != nil {
		t.Error(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListInterfaces", vcrConf)
	client.HttpClient = vcr.Client
	opts := ListInterfacesRequest{}
	res, err := client.ListInterfaces(opts)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}
