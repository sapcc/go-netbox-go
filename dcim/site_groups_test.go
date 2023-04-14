package dcim

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListSiteGroups(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListSiteGroups", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListSiteGroupsRequest{}
	res, err := client.ListSiteGroups(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_GetSiteGroup(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("GetSiteGroup", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetSiteGroup(71)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Name)
}
