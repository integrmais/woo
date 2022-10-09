package woo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/integrmais/woo"
)

const versionMock = "v3"
const consumerKeyMock = "ck_api_key"
const consumerSecretMock = "ck_api_secret"

func TestNewClient(t *testing.T) {
	serverMock := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(""))
	}))

	defer func() { serverMock.Close() }()

	c := woo.NewClient(serverMock.URL, versionMock, consumerKeyMock, consumerSecretMock)

	err := c.Ping()
	if err != nil {
		t.Fatalf("New Client failed: %v", err.Error())
	}
}
