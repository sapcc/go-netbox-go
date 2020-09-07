package ipam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListIpAddresses(opts ListIpAddressesRequest) (ListIpAddressesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "ip-addresses/", nil)
	if err != nil {
		return ListIpAddressesResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListIpAddressesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListIpAddressesResponse{}, err
	}
	if response.StatusCode != 200 {
		return ListIpAddressesResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListIpAddressesResponse{}
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListIpAddressesResponse{}, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return ListIpAddressesResponse{}, err
	}
	return resObj, nil
}

func setListIpAddressesParams(req *http.Request, opts ListIpAddressesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.InterfaceId != 0 {
		q.Set("interface_id", strconv.Itoa(opts.InterfaceId))
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) CreateIpAddress(address IpAddress) error {
	body, err := json.Marshal(address)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String() + basePath + "ip-addresses/", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 201 {
		return fmt.Errorf("unexpected response code of %d", response.StatusCode)
	}
	return nil
}