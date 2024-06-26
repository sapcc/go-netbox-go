package tenancy

import (
	"os"
	"testing"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetTenant(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListTenants", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetTenant(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	//assert.NotEqual(t, 0, res.Count)
}
func TestClient_ListTenants(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListTenants", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListTenantsRequest{}
	res, err := client.ListTenants(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}
