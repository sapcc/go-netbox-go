package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sapcc/go-netbox-go/models"
)

func (c *Client) GetCable(id int) (*models.Cable, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"cables/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.Cable{}
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

// func (c *Client) ListCables(opts models.ListCablesRequest) (*models.ListCablesResponse, error) {
// 	request, err := http.NewRequest("GET", c.BaseUrl.String()+basePath+"cables/", nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	c.SetAuthToken(&request.Header)
// 	setListCablesParams(request, opts)
// 	response, err := c.HttpClient.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if response.StatusCode != 200 {
// 		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
// 	}
// 	resObj := models.ListCablesResponse{}
// 	bytes, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(bytes, &resObj)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &resObj, nil
// }
// func setListCablesParams(req *http.Request, opts models.ListCablesRequest) {
// 	q := req.URL.Query()
// 	opts.SetListParams(&q)
// 	if opts.CableID != 0 {
// 		q.Set("id", strconv.Itoa(opts.CableID))
// 	}
// 	if opts.CableType != "" {
// 		q.Set("a_terminations.object_type", opts.CableType)
// 		q.Set("b_terminations.object_type", opts.CableType)
// 	}
// 	req.URL.RawQuery = q.Encode()
// }
