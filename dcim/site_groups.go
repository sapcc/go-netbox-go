// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package dcim

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListSiteGroups(opts models.ListSiteGroupsRequest) (*models.ListSiteGroupsResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"site-groups/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListSiteGroupsParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListSiteGroupsResponse{}
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

func setListSiteGroupsParams(req *http.Request, opts models.ListSiteGroupsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetSiteGroup(id int) (*models.SiteGroup, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"site-groups/"+strconv.Itoa(id)+"/", http.NoBody)
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
	resObj := models.SiteGroup{}
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
