// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

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
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"tags/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListTagsParams(request, opts)
	response, err := c.HTTPClient().Do(request)
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
// 	request, err := http.NewRequestWithContext(context.TODO(), "POST", c.BaseURL().String()+basePath+"tags/", bytes.NewBuffer(body))
// 	fmt.Println(request)
// 	if err != nil {
// 		return err
// 	}
// 	c.ApplyAuthTokenToHeader(&request.Header)
// 	response, err := c.HTTPClient().Do(request)
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
