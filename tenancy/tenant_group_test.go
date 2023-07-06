package tenancy

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListTenantGroups(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListTenantGroups", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListTenantGroupsRequest{}
	res, err := client.ListTenantGroups(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}
