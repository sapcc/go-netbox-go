// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package dcim

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

func (c *Client) ListInterfaces(opts models.ListInterfacesRequest) (*models.ListInterfacesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"interfaces/", bytes.NewBuffer([]byte{'a'}))
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListInterfacesParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListInterfacesResponse{}
	byts, err := io.ReadAll(response.Body)
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
	if opts.DeviceID != 0 {
		q.Set("device_id", strconv.Itoa(opts.DeviceID))
	}
	if opts.MacAddress != "" {
		q.Set("mac_address", opts.MacAddress)
	}
	if opts.LagID != 0 {
		q.Set("lag_id", strconv.Itoa(opts.LagID))
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) UpdateInterface(interf models.WritableInterface, id int) (*models.Interface, error) {
	body, err := json.Marshal(interf)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, c.BaseURL().String()+basePath+"interfaces/"+strconv.Itoa(id)+"/", bytes.NewBuffer(body))
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
	resObj := models.Interface{}
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

func (c *Client) CreateInterface(interf models.WritableInterface) (*models.Interface, error) {
	body, err := json.Marshal(interf)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, c.BaseURL().String()+basePath+"interfaces/", bytes.NewBuffer(body))
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
	resObj := models.Interface{}
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

func (c *Client) DeleteInterface(id int) error {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, c.BaseURL().String()+basePath+"interfaces/"+strconv.Itoa(id)+"/", http.NoBody)
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
