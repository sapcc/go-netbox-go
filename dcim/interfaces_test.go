package dcim

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
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
	opts := models.ListInterfacesRequest{}
	res, err := client.ListInterfaces(opts)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}
