package extras

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListTags(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient
	vcr := govcr.NewVCR("ListTags", vcrConf)
	client.HTTPClient = vcr.Client
	opts := models.ListTagsRequest{}
	res, err := client.ListTags(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, res.Count)
	t.Log(res.Results[0].Name)
}

// Permission issue in netbox - not allwed to create tags

// func TestClient_CreateTag(t *testing.T) {
// 	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	vcrConf := &govcr.VCRConfig{}
// 	vcrConf.Client = client.HTTPClient
// 	vcr := govcr.NewVCR("CreateTag", vcrConf)
// 	client.HTTPClient = vcr.Client
// 	opts := models.Tag{
// 		NestedTag: models.NestedTag{
// 			Name: "testTAG",
// 			Slug: "testtag",
// 		},
// 		Description: "test tag",
// 	}
// 	err = client.CreateTag(opts)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
