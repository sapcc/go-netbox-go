package ipam

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListVlans(opts models.ListVlanRequest) (*models.ListVlanResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"vlans/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setVlanParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListVlanResponse{}
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

func (c *Client) GetVlan(id int) (*models.Vlan, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"vlans/"+strconv.Itoa(id)+"/", nil)
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
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.Vlan{}
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

func setVlanParams(req *http.Request, opts models.ListVlanRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Group != "" {
		q.Set("group", opts.Group)
	}
	req.URL.RawQuery = q.Encode()
}
