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

func (c *Client) ListPrefixes(opts models.ListPrefixesRequest) (*models.ListPrefixesReponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"prefixes/", nil)
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
	request, err := http.NewRequest("POST", c.BaseUrl.String()+basePath+"prefixes/", bytes.NewBuffer(body))
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

func (c *Client) ListAvailableIps(id int) ([]models.AvailableIp, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"prefixes/"+strconv.Itoa(id)+"/available-ips/", nil)
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
		return nil, fmt.Errorf("unexpected reponse code of %d, %s", response.StatusCode, errBody)
	}
	var resObj []models.AvailableIp
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return resObj, nil
}

func (c *Client) CreateAvailablePrefix(id int, opts models.CreateAvailablePrefixRequest) (*models.Prefix, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String()+basePath+"prefixes/"+strconv.Itoa(id)+"/available-prefixes/", bytes.NewBuffer(body))
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
	request, err := http.NewRequest("PUT", c.BaseUrl.String()+basePath+"prefixes/"+strconv.Itoa(prefix.Id)+"/", bytes.NewBuffer(body))
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

func (c *Client) DeletePrefix(id int) error {
	request, err := http.NewRequest("DELETE", c.BaseUrl.String()+basePath+"prefixes/"+strconv.Itoa(id)+"/", nil)
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
	if opts.Site != "" {
		q.Set("site", opts.Site)
	}
	if opts.TenantId != 0 {
		q.Set("tenant_id", strconv.Itoa(opts.TenantId))
	}
	if opts.VrfId != 0 {
		q.Set("vrf_id", strconv.Itoa(opts.VrfId))
	}
	if opts.Prefix != "" {
		q.Set("prefix", opts.Prefix)
	}
	if opts.MaskLength != 0 {
		q.Set("mask_length", strconv.Itoa(opts.MaskLength))
	}
	if opts.MaskLengthGte != 0 {
		q.Set("mask_length__gte", strconv.Itoa(opts.MaskLengthGte))
	}
	if opts.MaskLengthLte != 0 {
		q.Set("mask_length__lte", strconv.Itoa(opts.MaskLengthLte))
	}
	if opts.Status != "" {
		q.Set("status", opts.Status)
	}
	if opts.Within != "" {
		q.Set("within", opts.Within)
	}
	if opts.Contains != "" {
		q.Set("contains", opts.Contains)
	}
	if opts.Children != nil {
		q.Set("children", strconv.Itoa(*opts.Children))
	}
	req.URL.RawQuery = q.Encode()
}
