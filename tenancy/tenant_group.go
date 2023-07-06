package tenancy

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListTenantGroups(opts models.ListTenantGroupsRequest) (*models.ListTenantGroupsResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"tenant-groups/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListTenantGroupParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListTenantGroupsResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func setListTenantGroupParams(req *http.Request, opts models.ListTenantGroupsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}
