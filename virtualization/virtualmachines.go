package virtualization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListVirtualMachines (opts ListVirtualMachinesRequest) (ListVirtualMachinesResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "virtual-machines/", nil)
	if err != nil {
		return ListVirtualMachinesResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListVirtualMachinesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListVirtualMachinesResponse{}, err
	}
	if response.StatusCode != 200 {
		return ListVirtualMachinesResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListVirtualMachinesResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListVirtualMachinesResponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return ListVirtualMachinesResponse{}, err
	}
	return resObj, nil
}

func setListVirtualMachinesParams(req *http.Request, opts ListVirtualMachinesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}