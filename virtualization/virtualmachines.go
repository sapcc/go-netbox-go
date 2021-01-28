package virtualization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
)

func (c *Client) CreateVirtualMachine (vm models.WriteableVirtualMachine) error {
	body, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String() + basePath + "virtual-machines/", bytes.NewBuffer(body))
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
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListVirtualMachinesResponse{}
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

	func setListVirtualMachinesParams(req *http.Request, opts models.ListVirtualMachinesRequest) {
		q := req.URL.Query()
		opts.SetListParams(&q)
		req.URL.RawQuery = q.Encode()
	}
