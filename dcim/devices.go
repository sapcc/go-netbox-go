/*
 *   Copyright 2020 SAP SE
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

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

func (c *Client) ListDevices(opts models.ListDevicesRequest) (*models.ListDevicesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"devices/?exclude=config_context", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListDevicesParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	var resObj = models.ListDevicesResponse{}
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

func setListDevicesParams(req *http.Request, opts models.ListDevicesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)

	if opts.ID != 0 {
		q.Set("id", strconv.Itoa(opts.ID))
	}
	if opts.ClusterID != 0 {
		q.Set("cluster_id", strconv.Itoa(opts.ClusterID))
	}
	if opts.DeviceTypeID != 0 {
		q.Set("device_type_id", strconv.Itoa(opts.DeviceTypeID))
	}
	if opts.Site != "" {
		q.Set("site", opts.Site)
	}
	if opts.SiteId != 0 {
		q.Set("site_id", strconv.Itoa(opts.SiteId))
	}
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	if opts.Name != "" {
		q.Set("name", opts.Name)
	}
	if opts.RackID != 0 {
		q.Set("rack_id", strconv.Itoa(opts.RackID))
	}
	if opts.Serial != "" {
		q.Set("serial", opts.Serial)
	}
	if opts.RoleID != 0 {
		q.Set("role_id", strconv.Itoa(opts.RoleID))
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetDevice(id int) (*models.Device, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"devices/"+strconv.Itoa(id)+"/?exclude=config_context", http.NoBody)
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
	var resObj = models.Device{}
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

func (c *Client) GetDeviceWithContext(id int) (*models.Device, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"devices/"+strconv.Itoa(id)+"/", http.NoBody)
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
	var resObj = models.Device{}
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

func (c *Client) CreateDevice(dev models.WritableDeviceWithConfigContext) (*models.Device, error) {
	body, err := json.Marshal(dev)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, c.BaseURL().String()+basePath+"devices/", bytes.NewBuffer(body))
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
	resObj := models.Device{}
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

func (c *Client) DeleteDevice(id int) error {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, c.BaseURL().String()+basePath+"devices/"+strconv.Itoa(id)+"/", http.NoBody)
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

func (c *Client) UpdateDevice(dev models.WritableDeviceWithConfigContext) (*models.Device, error) {
	body, err := json.Marshal(dev)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPatch, c.BaseURL().String()+basePath+"devices/"+strconv.Itoa(dev.ID)+"/", bytes.NewBuffer(body))
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
	resObj := models.Device{}
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

/*
func (c *Client) ListDevicesByCluster(id int) (*models.ListDevicesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), "GET", c.BaseURL().String() + basePath + "devices/?cluster_id=" + strconv.Itoa(id), nil )
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	var resObj = models.ListDevicesResponse{}
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
*/
