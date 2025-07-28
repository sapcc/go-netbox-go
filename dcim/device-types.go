// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package dcim

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListDeviceTypes(opts models.ListDeviceTypesRequest) (*models.ListDeviceTypesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"device-types/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListDeviceTypesParams(request, opts)
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
	resObj := models.ListDeviceTypesResponse{}
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

func setListDeviceTypesParams(req *http.Request, opts models.ListDeviceTypesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}
