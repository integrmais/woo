package woo

import (
	"errors"
	"fmt"
	"net/http"
)

const DefaultDateFormat = "2006-01-02T15:04:05"

func NewClient(baseUrl, versionApi, consumerKey, consumerSecret string) *Client {
	c := &Client{
		BaseUrl:        baseUrl,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		VersionApi:     versionApi,
	}

	c.apiUrl = fmt.Sprintf("%s/wp-json/wc/%s", baseUrl, versionApi)

	c.Product = (*ProductService)(c)

	return c
}

func (c *Client) Ping() error {
	res, err := http.DefaultClient.Get(c.apiUrl)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New(fmt.Sprintf("Failed API Ping to %s", c.BaseUrl))
}
