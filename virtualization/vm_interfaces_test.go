package virtualization

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_CreateDeleteTaggedVMInterface(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateTaggedVMInterface", vcrConf)
	client.HttpClient = vcr.Client
	tag1 := models.NestedTag{
		Name: "101",
		Slug: "101",
	}
	tag2 := models.NestedTag{
		Name: "128",
		Slug: "128",
	}
	tags := []models.NestedTag{
		tag1, tag2,
	}
	vmi := models.WritableVMInterface{
		VirtualMachine: 1107,
		Name: "test-tagged-interface",
		Enabled: true,
		MTU: 9000,
		Description: "this is a test tagged interface",
		Tags: tags,
		TaggedVlans: []models.NestedVLAN{},
	}
	vmi2, err := client.CreateVMInterface(vmi)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vmi2)
	assert.Equal(t, 2, len(vmi2.Tags))
	err = client.DeleteVMInterface(vmi2.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateDeleteVMInterface(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateVMInterface", vcrConf)
	client.HttpClient = vcr.Client
	vmi := models.WritableVMInterface{
		VirtualMachine: 1107,
		Name: "test-interface",
		Enabled: true,
		MTU: 9000,
		Description: "this is a test interface",
	}
	vmi2, err := client.CreateVMInterface(vmi)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "test-interface", vmi2.Name)
	assert.Equal(t, 9000, vmi2.MTU)
	assert.Equal(t, true, vmi2.Enabled)
	assert.Equal(t, 1107, vmi2.VirtualMachine.Id)
	t.Log(vmi2)
	err = client.DeleteVMInterface(vmi2.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListVMInterfaces(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListVMInterfaces", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListVMInterfacesRequest{}
	//opts.VmId = 804
	//opts.Id = 971
	res, err := client.ListVMInterfaces(opts)
	//t.Log(res)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	t.Log(res.Results[0].Name)
	assert.NotEqual(t, 0, res.Count)
}
