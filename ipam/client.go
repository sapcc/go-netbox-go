package ipam

import (
	"net/http"
	"net/url"

	"github.com/sapcc/go-netbox-go/common"
)

const basePath = "/api/ipam/"

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
	// leave here for staging tests
	//	tr := &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	}
	//	res.HttpClient = &http.Client{Transport: tr}
	res.AuthToken = authToken
	return res, nil
}
