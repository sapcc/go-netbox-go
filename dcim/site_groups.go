package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListSiteGroups(opts models.ListSiteGroupsRequest) (*models.ListSiteGroupsResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"site-groups/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListSiteGroupsParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListSiteGroupsResponse{}
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

func setListSiteGroupsParams(req *http.Request, opts models.ListSiteGroupsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetSiteGroup(id int) (*models.SiteGroup, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"site-groups/"+strconv.Itoa(id)+"/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.SiteGroup{}
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
