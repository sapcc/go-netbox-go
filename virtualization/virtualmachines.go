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

func (c *Client) CreateVirtualMachine (vm models.WriteableVirtualMachine) (*models.VirtualMachine, error) {
	body, err := json.Marshal(vm)
	if err != nil {
		return nil,err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String() + basePath + "virtual-machines/", bytes.NewBuffer(body))
	if err != nil {
		return nil,err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil,err
	}
	if response.StatusCode != 201 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil,fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.VirtualMachine{}
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj,nil
}

func (c *Client) DeleteVirtualMachine (id int) error {
	request, err := http.NewRequest("DELETE", c.BaseUrl.String() + basePath + "virtual-machines/" + strconv.Itoa(id) + "/", nil)
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

func (c *Client) GetVirtualMachine(id int) (*models.VirtualMachine, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "virtual-machines/" + strconv.Itoa(id) + "/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errorBody,_ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errorBody)
	}
	resObj := models.VirtualMachine{}
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

func (c *Client) ListVirtualMachines (opts models.ListVirtualMachinesRequest) (*models.ListVirtualMachinesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "virtual-machines/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListVirtualMachinesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.ListVirtualMachinesResponse{}
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

func setListVirtualMachinesParams(req *http.Request, opts models.ListVirtualMachinesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.ClusterId != 0 {
		q.Set("cluster_id", strconv.Itoa(opts.ClusterId))
	}
	if opts.RoleId != 0 {
		q.Set("role_id", strconv.Itoa(opts.RoleId))
	}
	req.URL.RawQuery = q.Encode()
}
