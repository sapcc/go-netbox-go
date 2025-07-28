// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package virtualization

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

func (c *Client) CreateVMInterface(vmni models.WritableVMInterface) (*models.VMInterface, error) {
	body, err := json.Marshal(vmni)
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
		return nil, fmt.Errorf("unexpected response code of %d:%s", response.StatusCode, errBody)
	}
	resObj := models.VMInterface{}
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

func (c *Client) UpdateVMInterface(vmi models.WritableVMInterface) (*models.VMInterface, error) {
	body, err := json.Marshal(vmi)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPatch, c.BaseURL().String()+basePath+"interfaces/"+strconv.Itoa(vmi.ID)+"/", bytes.NewBuffer(body))
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
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.VMInterface{}
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

func (c *Client) DeleteVMInterface(id int) error {
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

func (c *Client) ListVMInterfaces(opts models.ListVMInterfacesRequest) (*models.ListVMInterfacesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"interfaces/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListVMInterfacesParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListVMInterfacesResponse{}
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

func setListVMInterfacesParams(req *http.Request, opts models.ListVMInterfacesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.VMID != 0 {
		q.Set("virtual_machine_id", strconv.Itoa(opts.VMID))
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetVMInterface(id int) (*models.VMInterface, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"interfaces/"+strconv.Itoa(id)+"/", http.NoBody)
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
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.VMInterface{}
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
