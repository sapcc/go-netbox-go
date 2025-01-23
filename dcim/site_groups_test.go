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
