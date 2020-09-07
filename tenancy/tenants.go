package tenancy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListTenants (opts ListTenantsRequest) (ListTenantsResponse,error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "tenants/", nil)
	if err != nil {
		return ListTenantsResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListTenantsParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListTenantsResponse{},err
	}
	if response.StatusCode != 200 {
		return ListTenantsResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListTenantsResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListTenantsResponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return ListTenantsResponse{}, err
	}
	return resObj, nil
}

func setListTenantsParams(req *http.Request, opts ListTenantsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}