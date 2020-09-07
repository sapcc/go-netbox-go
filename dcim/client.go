package dcim

import (
	"github.com/sapcc/go-netbox-go/common"
	"net/http"
	"net/url"
)

const basePath = "/api/dcim/"

type Client struct {
	common.Client
}

func New(baseUrl string, authToken string) (*Client, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	res := &Client{}
	res.BaseUrl = *u
	res.HttpClient = &http.Client{}
	res.AuthToken = authToken
	return res, nil
}