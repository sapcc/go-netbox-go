package dcim

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListInterfaces(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("ListInterfaces", vcrConf)
	client.HTTPClient = vcr.Client
	opts := models.ListInterfacesRequest{}
	opts.DeviceID = 10867
	opts.Name = "bond2"
	res, err := client.ListInterfaces(opts)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(res)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_CreateDeleteInterface(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("CreateDeleteInterface", vcrConf)
	client.HTTPClient = vcr.Client
	wInt := models.WritableInterface{}
	wInt.MacAddress = "aaaa:bbbb:cccc"
	wInt.Name = "Test Interface"
	wInt.Device = 10867
	wInt.Type = "1000base-t"
	interf, err := client.CreateInterface(wInt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(interf)

	err = client.DeleteInterface(interf.ID)
	if err != nil {
		t.Fatal(err)
	}
}
