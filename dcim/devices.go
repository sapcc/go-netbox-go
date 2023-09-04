package dcim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListDevices(opts models.ListDevicesRequest) (*models.ListDevicesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"devices/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListDevicesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	var resObj = models.ListDevicesResponse{}
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

func setListDevicesParams(req *http.Request, opts models.ListDevicesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)

	if opts.Id != 0 {
		q.Set("id", strconv.Itoa(opts.Id))
	}
	if opts.ClusterId != 0 {
		q.Set("cluster_id", strconv.Itoa(opts.ClusterId))
	}
	if opts.DeviceTypeId != 0 {
		q.Set("device_type_id", strconv.Itoa(opts.DeviceTypeId))
	}
	if opts.Site != "" {
		q.Set("site", opts.Site)
	}
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	if opts.Name != "" {
		q.Set("name", opts.Name)
	}
	if opts.RackId != 0 {
		q.Set("rack_id", strconv.Itoa(opts.RackId))
	}
	if opts.Serial != "" {
		q.Set("serial", opts.Serial)
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetDevice(id int) (*models.Device, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"devices/"+strconv.Itoa(id)+"/", nil)
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
	var resObj = models.Device{}
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

func (c *Client) CreateDevice(dev models.WritableDeviceWithConfigContext) (*models.Device, error) {
	body, err := json.Marshal(dev)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String()+basePath+"devices/", bytes.NewBuffer(body))
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
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.Device{}
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

func (c *Client) DeleteDevice(id int) error {
	request, err := http.NewRequest("DELETE", c.BaseUrl.String()+basePath+"devices/"+strconv.Itoa(id)+"/", nil)
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

/*
func (c *Client) ListDevicesByCluster(id int) (*models.ListDevicesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "devices/?cluster_id=" + strconv.Itoa(id), nil )
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
	var resObj = models.ListDevicesResponse{}
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
*/
