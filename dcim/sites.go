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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListSites(opts models.ListSitesRequest) (*models.ListSitesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL.String()+basePath+"sites/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListSitesParams(request, opts)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListSitesResponse{}
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

func setListSitesParams(req *http.Request, opts models.ListSitesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetSite(id int) (*models.Site, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL.String()+basePath+"sites/"+strconv.Itoa(id)+"/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.Site{}
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
