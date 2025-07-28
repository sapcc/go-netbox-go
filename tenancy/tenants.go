// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package tenancy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) GetTenant(id int) (*models.Tenant, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"tenants/"+strconv.Itoa(id), http.NoBody)
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
	var resObj = models.Tenant{}
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

func (c *Client) ListTenants(opts models.ListTenantsRequest) (*models.ListTenantsResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"tenants/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListTenantsParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListTenantsResponse{}
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

func setListTenantsParams(req *http.Request, opts models.ListTenantsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}
