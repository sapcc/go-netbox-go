package virtualization

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) ListClusters (opts models.ListClusterRequest) (*models.ListClusterResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "clusters/", bytes2.NewBuffer([]byte{'a'}))
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListClusterParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil{
		return nil, err
	}
	if response.StatusCode != 200 {
		errorBody,_ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("unexpected return code of %d: %s", response.StatusCode, errorBody)
	}
	resObj := models.ListClusterResponse{}
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

func setListClusterParams(req *http.Request, opts models.ListClusterRequest) {
	q:= req.URL.Query()
	opts.SetListParams(&q)
	if opts.Id != 0 {
		q.Set("id", strconv.Itoa(opts.Id))
	}
	if opts.Region != "" {
		q.Set("region", opts.Region)
	}
	if opts.Type != "" {
		q.Set("type", opts.Type)
	}
	req.URL.RawQuery = q.Encode()
}
