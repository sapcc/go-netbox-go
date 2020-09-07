package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListRacks (opts ListRacksRequest) (ListRacksResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "racks/", nil)
	if err != nil {
		return ListRacksResponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListRacksParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListRacksResponse{}, err
	}
	if response.StatusCode != 200 {
		return ListRacksResponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListRacksResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListRacksResponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return ListRacksResponse{}, err
	}
	return resObj, nil
}

func setListRacksParams(req *http.Request, opts ListRacksRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}
