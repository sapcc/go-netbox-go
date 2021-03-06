package virtualization

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListClusters(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListClusters", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListClusterRequest{}
	//opts.VmId = 1060
	opts.Id = 632
	res, err := client.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	t.Log(res.Results[0].Name)
	assert.NotEqual(t, 0, res.Count)
}
