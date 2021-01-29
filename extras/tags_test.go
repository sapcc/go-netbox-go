package extras

import (
	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListTags(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListTags", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListTagsRequest{}
	res, err := client.ListTags(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0].Name)
}

/*
permission error ...

func TestClient_CreateTag(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateTag", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.Tag{
		Name: "testTag",
		Slug: "test-tag",
		Description: "test tag description",
	}
	err = client.CreateTag(opts)
	if err != nil {
		t.Fatal(err)
	}
}
 */