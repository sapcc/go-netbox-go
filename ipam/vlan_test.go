package ipam

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListVlans(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"))
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
	assert.NotEqual(t, 0, res.Count)
}
