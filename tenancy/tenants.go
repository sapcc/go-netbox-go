package tenancy

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListTenants (opts models.ListTenantsRequest) (*models.ListTenantsResponse,error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "tenants/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListTenantsParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil,err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListTenantsResponse{}
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

func setListTenantsParams(req *http.Request, opts models.ListTenantsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}