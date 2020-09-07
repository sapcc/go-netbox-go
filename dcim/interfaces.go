package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListInterfaces (opts models.ListInterfacesRequest) (*models.ListInterfacesResponse, error) {
	request, err := http.NewRequest("Get", c.BaseUrl.String() + basePath + "interfaces/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListInterfacesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil{
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListInterfacesResponse{}
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

func setListInterfacesParams(req *http.Request, opts models.ListInterfacesRequest) {
	q:= req.URL.Query()
	opts.SetListParams(&q)
	if opts.Type != "" {
		q.Set("type", opts.Type)
	}
	if opts.DeviceId != 0 {
		q.Set("device_id", strconv.Itoa(opts.DeviceId))
	}
	req.URL.RawQuery = q.Encode()
}