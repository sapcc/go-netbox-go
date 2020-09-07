package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListInterfaces (opts ListInterfacesRequest) (ListInterfacesResponse, error) {
	request, err := http.NewRequest("Get", c.BaseUrl.String() + basePath + "interfaces/", nil)
	if err != nil {
		return ListInterfacesResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListInterfacesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil{
		return ListInterfacesResponse{}, err
	}
	if response.StatusCode != 200 {
		return ListInterfacesResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListInterfacesResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListInterfacesResponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return ListInterfacesResponse{}, err
	}
	return resObj, nil
}

func setListInterfacesParams(req *http.Request, opts ListInterfacesRequest) {
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