package woo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/integrmais/woo"
)

var customer = []byte(`
{
    "id": 5,
    "email": "joao.silva@example.com",
    "first_name": "joão",
    "last_name": "silva",
    "role": "author"
}
`)

var customers = []byte(`[
  {
    "id": 26,
    "date_created": "2017-03-21T16:11:14",
    "date_created_gmt": "2017-03-21T19:11:14",
    "date_modified": "2017-03-21T16:11:16",
    "date_modified_gmt": "2017-03-21T19:11:16",
    "email": "joao.silva@example.com",
    "first_name": "João",
    "last_name": "Silva",
    "role": "customer",
    "username": "joao.silva",
    "billing": {
      "first_name": "João",
      "last_name": "Silva",
      "company": "",
      "address_1": "Av. Brasil, 432",
      "address_2": "",
      "city": "Rio de Janeiro",
      "state": "RJ",
      "postcode": "12345-000",
      "country": "BR",
      "email": "joao.silva@example.com",
      "phone": "(55) 5555-5555"
    },
    "shipping": {
      "first_name": "João",
      "last_name": "Silva",
      "company": "",
      "address_1": "Av. Brasil, 432",
      "address_2": "",
      "city": "Rio de Janeiro",
      "state": "RJ",
      "postcode": "12345-000",
      "country": "BR"
    },
    "is_paying_customer": false,
    "avatar_url": "https://secure.gravatar.com/avatar/be7b5febff88a2d947c3289e90cdf017?s=96",
    "meta_data": [],
    "_links": {
      "self": [
        {
          "href": "https://example.com/wp-json/wc/v3/customers/26"
        }
      ],
      "collection": [
        {
          "href": "https://example.com/wp-json/wc/v3/customers"
        }
      ]
    }
  },
  {
    "id": 25,
    "date_created": "2017-03-21T16:09:28",
    "date_created_gmt": "2017-03-21T19:09:28",
    "date_modified": "2017-03-21T16:09:30",
    "date_modified_gmt": "2017-03-21T19:09:30",
    "email": "john.doe@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "customer",
    "username": "john.doe",
    "billing": {
      "first_name": "John",
      "last_name": "Doe",
      "company": "",
      "address_1": "969 Market",
      "address_2": "",
      "city": "San Francisco",
      "state": "CA",
      "postcode": "94103",
      "country": "US",
      "email": "john.doe@example.com",
      "phone": "(555) 555-5555"
    },
    "shipping": {
      "first_name": "John",
      "last_name": "Doe",
      "company": "",
      "address_1": "969 Market",
      "address_2": "147",
      "city": "San Francisco",
      "state": "CA",
      "postcode": "94103",
      "country": "US"
    },
    "is_paying_customer": false,
    "avatar_url": "https://secure.gravatar.com/avatar/8eb1b522f60d11fa897de1dc6351b7e8?s=96",
    "meta_data": [],
    "_links": {
      "self": [
        {
          "href": "https://example.com/wp-json/wc/v3/customers/25"
        }
      ],
      "collection": [
        {
          "href": "https://example.com/wp-json/wc/v3/customers"
        }
      ]
    }
  }
]`)

func TestListCustomers(t *testing.T) {
	serverMock := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(customers)
	}))

	c := woo.NewClient(serverMock.URL, versionMock, consumerKeyMock, consumerSecretMock)

	customers, err := c.Customer.List()
	if err != nil {
		t.Fatalf("Expected empty error, got %v", err.Error())
	}

	if customers[0].Id != 26 {
		t.Fatalf("Expected first customer to be 26, got %v", customers[0].Id)
	}
}

func TestCreateCustomer(t *testing.T) {
    serverMock := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        res.WriteHeader(http.StatusOK)
        res.Write(customer)
    }))

    c := woo.NewClient(serverMock.URL, versionMock, consumerKeyMock, consumerSecretMock)

    cm := woo.Customer{
		FirstName: "joão",
		LastName: "silva",
		Email: "joao.silva@example.com",
		Role: woo.CustomerAuthorRole,
	}

	customerCreated, err := c.Customer.Create(cm)
    if err != nil {
        t.Fatalf("Expected empty error, got %v", err.Error())
    }

	if customerCreated.FirstName != cm.FirstName {
		t.Fatalf("Expected customer first name %s, got %s", cm.FirstName, customerCreated.FirstName)
	}

	if customerCreated.Role != woo.CustomerAuthorRole {
		t.Fatalf("Expected customer role %s, got %s", woo.CustomerAuthorRole, customerCreated.Role)
	}
}
