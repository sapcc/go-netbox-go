package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListPlatforms(opts models.ListPlatformsRequest) (*models.ListPlatformsResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "platforms/", nil)
	if err != nil {
		return nil,err
	}
	c.SetAuthToken(&request.Header)
	setListPlatformParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListPlatformsResponse{}
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

func setListPlatformParams(req *http.Request, opts models.ListPlatformsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}