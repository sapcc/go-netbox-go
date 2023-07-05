package dcim

import (
	"os"
	"testing"

	"github.com/sapcc/go-netbox-go/models"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListCables(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListCables", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListCablesRequest{}
	opts.Id = 2026
	res, err := client.ListCables(opts) //2026
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println("RES:", res.Results)
	// var cab models.Cable
	// //cab = res.Results
	// fmt.Println("RESULTS:")
	// for _, cab = range res.Results {
	// 	fmt.Println()
	// 	fmt.Println("Cable ID:", cab.ID)
	// 	fmt.Println("Cable Type:", cab.AssignedObjectType)
	// 	fmt.Println()
	// 	fmt.Println("A terminations:", cab.Aterminations)
	// 	fmt.Println()
	// 	for _, term := range cab.Aterminations {
	// 		fmt.Println()
	// 		fmt.Println("Nested interface ID:", term.Id)
	// 		fmt.Println("Device ID:", term.Device.Id)
	// 		fmt.Println("Device name:", term.Device.Name)
	// 		fmt.Println("Nested interface name:", term.Name)
	// 	}
	// 	fmt.Println()
	// 	fmt.Println("B terminations:", cab.Bterminations)
	// 	fmt.Println()
	// 	for _, term := range cab.Bterminations {
	// 		fmt.Println()

	// 		fmt.Println("Nested interface ID:", term.Id)
	// 		fmt.Println("Device ID:", term.Device.Id)
	// 		fmt.Println("Device name:", term.Device.Name)
	// 		fmt.Println("Nested interface name:", term.Name)
	// 	}
	// 	//t.Log(res)
	// }
	assert.NotEqual(t, 0, res.Count)
	t.Log("Cable ID:", res.Results[0].ID)
}
