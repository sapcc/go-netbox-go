package virtualization

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListVMInterfaces (opts models.ListVMInterfacesRequest) (*models.ListVMInterfacesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "interfaces/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListVMInterfacesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil{
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListVMInterfacesResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func setListVMInterfacesParams(req *http.Request, opts models.ListVMInterfacesRequest) {
	q:= req.URL.Query()
	opts.SetListParams(&q)
	if opts.Id != 0 {
		q.Set("id", strconv.Itoa(opts.Id))
	}
	if opts.VmId != 0 {
		q.Set("virtual_machine_id", strconv.Itoa(opts.VmId))
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetVMInterface (id int) (*models.VMInterface, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "interfaces/" + strconv.Itoa(id) + "/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.VMInterface{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}
