package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListSites (opts ListSitesRequest) (ListSitesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "sites/", nil)
	if err != nil {
		return ListSitesResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListSitesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListSitesResponse{}, err
	}
	if response.StatusCode != 200 {
		return ListSitesResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListSitesResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListSitesResponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	return resObj, nil
}

func setListSitesParams(req *http.Request, opts ListSitesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}

func (c *Client) GetSite (id int) (Site, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "sites/" + strconv.Itoa(id) + "/", nil)
	if err != nil {
		return Site{}, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return Site{}, err
	}
	if response.StatusCode != 200 {
		return Site{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := Site{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Site{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return Site{}, err
	}
	return resObj, nil
}