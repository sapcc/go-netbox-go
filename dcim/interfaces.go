package dcim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListInterfaces(opts models.ListInterfacesRequest) (*models.ListInterfacesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"interfaces/", bytes.NewBuffer([]byte{'a'}))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListInterfacesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListInterfacesResponse{}
	byts, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byts, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func setListInterfacesParams(req *http.Request, opts models.ListInterfacesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Type != "" {
		q.Set("type", opts.Type)
	}
	if opts.DeviceId != 0 {
		q.Set("device_id", strconv.Itoa(opts.DeviceId))
	}
	if opts.MacAddress != "" {
		q.Set("mac_address", opts.MacAddress)
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) UpdateInterface (interf models.WritableInterface, id int) (*models.Interface, error) {
	body, err := json.Marshal(interf)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", c.BaseUrl.String() + basePath + "interfaces/" + strconv.Itoa(id) + "/", bytes.NewBuffer(body) )
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
	resObj := models.Interface{}
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

func (c *Client) CreateInterface (interf models.WritableInterface) (*models.Interface, error) {
	body, err := json.Marshal(interf)
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
	if response.StatusCode != 201 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.Interface{}
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