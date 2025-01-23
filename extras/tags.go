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

package extras

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListTags(opts models.ListTagsRequest) (*models.ListTagsResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL.String()+basePath+"tags/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListTagsParams(request, opts)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListTagsResponse{}
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

// permission issue in netbox - not allowed to create tags

// func (c *Client) CreateTag(tag models.Tag) error {
// 	body, err := json.Marshal(tag)
// 	if err != nil {
// 		return err
// 	}
// 	request, err := http.NewRequestWithContext(context.TODO(), "POST", c.BaseURL.String()+basePath+"tags/", bytes.NewBuffer(body))
// 	fmt.Println(request)
// 	if err != nil {
// 		return err
// 	}
// 	c.SetAuthToken(&request.Header)
// 	response, err := c.HTTPClient.Do(request)
// 	fmt.Println("****")
// 	fmt.Println(response)
// 	fmt.Println("****")
// 	if err != nil {
// 		return err
// 	}
// 	if response.StatusCode != 201 {
// 		return fmt.Errorf("unexpected response code of %d", response.StatusCode)
// 	}

// 	return nil
// }

func setListTagsParams(req *http.Request, opts models.ListTagsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}
