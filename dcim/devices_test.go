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

func TestClient_CreateDeleteDevice(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateDeleteDevice", vcrConf)
	client.HttpClient = vcr.Client

	wDev := models.WritableDeviceWithConfigContext{}
	wDev.Name = "GNG Test Device"
	wDev.DeviceRole = 8   //Server
	wDev.DeviceType = 132 //PowerEdge R640
	wDev.Site = 21        //ap-sa-2a
	wDev.Status = "staged"
	wDev.Rack = 1126 //AP003-01 (sa-2-ET14-Cage2-AY9-4)
	wDev.Tenant = 1  //Converged Cloud
	dev, err := client.CreateDevice(wDev)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dev)

	err = client.DeleteDevice(dev.Id)
	if err != nil {
		t.Fatal(err)
	}
}
