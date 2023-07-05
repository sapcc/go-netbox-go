package dcim

import (
	"os"
	"testing"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetDevice(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("GetDevice", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetDevice(10867)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	t.Log(res.Cluster)
}

func TestClient_ListDevices(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListDevices", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListDevicesRequest{}
	//opts.Id = 12509
	res, err := client.ListDevices(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Results[0].Name)
}

func TestClient_ListDevicesByCluster(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListDevicesByCluster", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListDevicesRequest{
		ClusterId: 831,
	}
	res, err := client.ListDevices(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0])
}
