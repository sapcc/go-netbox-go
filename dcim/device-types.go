package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) ListDeviceTypes(opts models.ListDeviceTypesRequest) (*models.ListDeviceTypesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"device-types/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListDeviceTypesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		errBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errBody)
	}
	resObj := models.ListDeviceTypesResponse{}
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

func setListDeviceTypesParams(req *http.Request, opts models.ListDeviceTypesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}
