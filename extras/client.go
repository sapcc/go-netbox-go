package extras

import (
	"crypto/tls"
	"github.com/sapcc/go-netbox-go/common"
	"net/http"
	"net/url"
)

const basePath = "/api/extras/"

type Client struct {
	common.Client
}

func New(baseUrl string, authToken string, insecureSkipVerify bool) (*Client, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	res := &Client{}
	res.BaseUrl = *u
	res.HttpClient = &http.Client{
		Transport: tr,
	}
	res.AuthToken = authToken
	return res, nil
}