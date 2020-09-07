package ipam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListPrefixes (opts ListPrefixesRequest) (ListPrefixesReponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "prefixes/", nil)
	if err != nil {
		return ListPrefixesReponse{}, err
	}
	c.SetAuthToken(&request.Header)
	setListPrefixesParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return ListPrefixesReponse{}, err
	}
	if response.StatusCode != 200 {
		return ListPrefixesReponse{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := ListPrefixesReponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ListPrefixesReponse{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return ListPrefixesReponse{}, err
	}
	return resObj, nil
}

func setListPrefixesParams(req *http.Request, opts ListPrefixesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}