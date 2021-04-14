package ipam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListPrefixes (opts models.ListPrefixesRequest) (*models.ListPrefixesReponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "prefixes/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListPrefixesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.ListPrefixesReponse{}
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

func (c *Client) CreatePrefix(prefix models.WriteablePrefix) (*models.Prefix, error) {
	body, err := json.Marshal(prefix)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String() + basePath + "prefixes/", bytes.NewBuffer(body))
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
	resObj := models.Prefix{}
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

func (c *Client) UpdatePrefix(prefix models.WriteablePrefix) (*models.Prefix, error) {
	body, err := json.Marshal(prefix)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", c.BaseUrl.String() + basePath + "prefixes/" + strconv.Itoa(prefix.Id) + "/", bytes.NewBuffer(body))
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
	resObj := models.Prefix{}
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

func (c *Client) DeletePrefix (id int) error {
	request, err := http.NewRequest("DELETE", c.BaseUrl.String() + basePath + "prefixes/" + strconv.Itoa(id) + "/", nil)
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

func setListPrefixesParams(req *http.Request, opts models.ListPrefixesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Role != "" {
		q.Set("role", opts.Role)
	}
	if opts.Tag != "" {
		q.Set("tag", opts.Tag)
	}
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	req.URL.RawQuery = q.Encode()
}