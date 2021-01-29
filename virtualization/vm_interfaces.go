package virtualization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) CreateVMInterface (vmni models.WritableVMInterface) (*models.VMInterface, error) {
	body, err := json.Marshal(vmni)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String() + basePath + "interfaces/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201{
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected response code of %d:%s", response.StatusCode, errBody)
	}
	resObj := models.VMInterface{}
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func (c *Client) DeleteVMInterface(id int) error {
	request, err := http.NewRequest("DELETE", c.BaseUrl.String() + basePath + "interfaces/" + strconv.Itoa(id) + "/", nil)
	if err != nil {
		return err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	return nil
}

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
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func setListVMInterfacesParams(req *http.Request, opts models.ListVMInterfacesRequest) {
	q:= req.URL.Query()
	opts.SetListParams(&q)
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
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}
