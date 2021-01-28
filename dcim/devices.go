package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListDevices(opts models.ListDevicesRequest) (*models.ListDevicesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "devices/", nil )
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
	req.URL.RawQuery = q.Encode()
	if opts.Id != 0 {
		q.Set("id", strconv.Itoa(opts.Id))
	}
	if opts.ClusterId != 0 {
		q.Set("cluster_id", strconv.Itoa(opts.ClusterId))
	}
}

func (c *Client) GetDevice(id int) (*models.Device, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "devices/" + strconv.Itoa(id) + "/", nil )
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