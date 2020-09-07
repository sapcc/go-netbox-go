package ipam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListRoles (opts ListRolesRequest) (ListRolesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "roles/", nil)
	if err != nil {
		return ListRolesResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListRolesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListRolesResponse{}, err
	}
	if response.StatusCode != 200 {
		return ListRolesResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListRolesResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListRolesResponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return ListRolesResponse{}, err
	}
	return resObj, nil
}

func setListRolesParams(req *http.Request, opts ListRolesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}