// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListPrefixes(opts models.ListPrefixesRequest) (*models.ListPrefixesReponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"prefixes/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListPrefixesParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.ListPrefixesReponse{}
	bytes, err := io.ReadAll(response.Body)
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
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, c.BaseURL().String()+basePath+"prefixes/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		errBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.Prefix{}
	byteses, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func (c *Client) ListAvailableIps(id int) ([]models.AvailableIP, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"prefixes/"+strconv.Itoa(id)+"/available-ips/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected response code of %d, %s", response.StatusCode, errBody)
	}
	var resObj []models.AvailableIP
	byteses, err := io.ReadAll(response.Body)
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
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, c.BaseURL().String()+basePath+"prefixes/"+strconv.Itoa(id)+"/available-prefixes/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		errBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.Prefix{}
	byteses, err := io.ReadAll(response.Body)
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
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, c.BaseURL().String()+basePath+"prefixes/"+strconv.Itoa(prefix.ID)+"/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected response code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.Prefix{}
	byteses, err := io.ReadAll(response.Body)
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
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, c.BaseURL().String()+basePath+"prefixes/"+strconv.Itoa(id)+"/", http.NoBody)
	if err != nil {
		return err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusNoContent {
		errBody, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
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
	if opts.TenantID != 0 {
		q.Set("tenant_id", strconv.Itoa(opts.TenantID))
	}
	if opts.VrfID != 0 {
		q.Set("vrf_id", strconv.Itoa(opts.VrfID))
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
