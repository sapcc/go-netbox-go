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

package virtualization

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListClusters(opts models.ListClusterRequest) (*models.ListClusterResponse, error) {
	// request, err := http.NewRequestWithContext(context.TODO(), "GET", c.BaseURL.String() + basePath + "clusters/", bytes2.NewBuffer([]byte{'a'}))
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL.String()+basePath+"clusters/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListClusterParams(request, opts)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errorBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errorBody)
	}
	resObj := models.ListClusterResponse{}
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

func setListClusterParams(req *http.Request, opts models.ListClusterRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.ID != 0 {
		q.Set("id", strconv.Itoa(opts.ID))
	}
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	if opts.Type != "" {
		q.Set("type", opts.Type)
	}
	req.URL.RawQuery = q.Encode()
}
