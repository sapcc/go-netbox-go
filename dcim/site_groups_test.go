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

package dcim

// func TestClient_ListSiteGroups(t *testing.T) {
// 	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	vcrConf := &govcr.VCRConfig{}
// 	vcrConf.Client = client.HTTPClient
// 	vcr := govcr.NewVCR("ListSiteGroups", vcrConf)
// 	client.HTTPClient = vcr.Client
// 	opts := models.ListSiteGroupsRequest{}
// 	res, err := client.ListSiteGroups(opts)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// t.Log(res)
// 	assert.NotEqual(t, 0, res.Count)
// }

// func TestClient_GetSiteGroup(t *testing.T) {
// 	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	vcrConf := &govcr.VCRConfig{}
// 	vcrConf.Client = client.HTTPClient
// 	vcr := govcr.NewVCR("GetSiteGroup", vcrConf)
// 	client.HTTPClient = vcr.Client
// 	res, err := client.GetSiteGroup(69)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(res.Name)
// }
