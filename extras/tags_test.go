/*
 *   Copyright 2020 SAP SE
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package extras

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/go-netbox-go/models"
)

func TestClient_ListTags(t *testing.T) {
	client, err := NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.GetHTTPClient()
	vcr := govcr.NewVCR("ListTags", vcrConf)
	client.SetHTTPClient(vcr.Client)
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
// 	client, err := NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	vcrConf := &govcr.VCRConfig{}
// 	vcrConf.Client = client.GetHTTPClient()
// 	vcr := govcr.NewVCR("CreateTag", vcrConf)
// 	client.SetHTTPClient(vcr.Client)
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
