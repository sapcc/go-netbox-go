package dcim

import (
	"github.com/seborama/govcr"
	"os"
	"testing"
)

func TestClient_GetDevice(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"))
	if err != nil {
		t.Error(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("GetDevice", vcrConf)
	client.HttpClient = vcr.Client
	res, err := client.GetDevice(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
