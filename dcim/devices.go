package dcim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) GetDevice(id int) (Device, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "devices/" + strconv.Itoa(id) + "/", nil )
	if err != nil {
		return Device{}, err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return Device{}, err
	}
	if response.StatusCode != 200 {
		return Device{}, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	var resObj = Device{}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Device{}, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return Device{}, err
	}
	return resObj, nil
}
