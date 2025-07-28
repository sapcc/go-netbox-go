// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListRoles(opts models.ListRolesRequest) (*models.ListRolesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"roles/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListRolesParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListRolesResponse{}
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

func setListRolesParams(req *http.Request, opts models.ListRolesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Name != "" {
		q.Set("name", opts.Name)
	}
	if opts.Slug != "" {
		q.Set("slug", opts.Slug)
	}
	req.URL.RawQuery = q.Encode()
}
