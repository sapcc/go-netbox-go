package ipam

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/sapcc/go-netbox-go/common"
)

const basePath = "/api/ipam/"

type Client struct {
	common.Client
}

func New(baseURL, authToken string, insecureSkipVerify bool) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify, // #nosec
	}
	res := &Client{}
	res.BaseURL = *u
	res.HTTPClient = &http.Client{
		Transport: tr,
	}
	res.AuthToken = authToken
	return res, nil
}
