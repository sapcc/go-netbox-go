package ipam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListIpAddresses(opts models.ListIpAddressesRequest) (*models.ListIpAddressesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"ip-addresses/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListIpAddressesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.ListIpAddressesResponse{}
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

func (c *Client) GetIpAdress(id int) (*models.IpAddress, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"ip-addresses/"+strconv.Itoa(id)+"/", nil)
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
	resObj := models.IpAddress{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, nil
	}
	return &resObj, nil
}

func setListIpAddressesParams(req *http.Request, opts models.ListIpAddressesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.InterfaceId != 0 {
		q.Set("interface_id", strconv.Itoa(opts.InterfaceId))
	}
	if opts.VmInterfaceId != 0 {
		q.Set("vminterface_id", strconv.Itoa(opts.VmInterfaceId))
	}
	if opts.DeviceId != 0 {
		q.Set("device_id", strconv.Itoa(opts.DeviceId))
	}
	if opts.Role != "" {
		q.Set("role", opts.Role)
	}
	if opts.Address != "" {
		q.Set("address", opts.Address)
	}
	if opts.VrfId != 0 {
		q.Set("vrf_id", strconv.Itoa(opts.VrfId))
	}
	if opts.Parent != "" {
		q.Set("parent", opts.Parent)
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) CreateIpAddress(address models.WriteableIpAddress) (*models.IpAddress, error) {
	body, err := json.Marshal(address)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String()+basePath+"ip-addresses/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected response code of %d:%s", response.StatusCode, errBody)
	}
	resObj := models.IpAddress{}
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

func (c *Client) UpdateIpAddress(address models.WriteableIpAddress) (*models.IpAddress, error) {
	body, err := json.Marshal(address)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", c.BaseUrl.String()+basePath+"ip-addresses/"+strconv.Itoa(address.Id)+"/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.IpAddress{}
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
