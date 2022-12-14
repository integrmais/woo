package woo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	CustomerAllRole        string = "all"
	CustomerAdminRole             = "administrator"
	CustomerAuthorRole            = "author"
	CustomerSellerRole            = "seller"
	CustomerShopManageRole        = "shop_manager"
)

type CustomerService Client

type Customer struct {
	Message          string `json:"message,omitempty"`
	Code             string `json:"code,omitempty"`
	Id               int64  `json:"id,omitempty"`
	Username         string `json:"username,omitempty"`
	AvatarURL        string `json:"avatar_url,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	IsPayingCustomer bool   `json:"is_paying_customer,omitempty"`
	DateCreated      string `json:"date_created,omitempty"`
	DateModified     string `json:"date_modified,omitempty"`
	Role             string `json:"role,omitempty"`
	ShippingAddress  `json:"shipping,omitempty"`
	BillingAddress   `json:"billing,omitempty"`
}

type ShippingAddress struct {
	Address1     string `json:"address_1"`
	Address2     string `json:"address_2"`
	City         string `json:"city"`
	Company      string `json:"company"`
	Country      string `json:"country"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Neighborhood string `json:"neighborhood"`
	Number       string `json:"number"`
	Phone        string `json:"phone"`
	Postcode     string `json:"postcode"`
	State        string `json:"state"`
}

type BillingAddress struct {
	Address1            string `json:"address_1"`
	Address2            string `json:"address_2"`
	Birthdate           string `json:"birthdate"`
	Cellphone           string `json:"cellphone"`
	City                string `json:"city"`
	FederalCompanyTaxId string `json:"cnpj"`
	Company             string `json:"company"`
	Country             string `json:"country"`
	FederalTaxId        string `json:"cpf"`
	Email               string `json:"email"`
	FirstName           string `json:"first_name"`
	IE                  string `json:"ie"`
	LastName            string `json:"last_name"`
	Neighborhood        string `json:"neighborhood"`
	Number              string `json:"number"`
	Persontype          string `json:"persontype"`
	Phone               string `json:"phone"`
	Postcode            string `json:"postcode"`
	FederalGeneralId    string `json:"rg"`
	Sex                 string `json:"sex"`
	State               string `json:"state"`
}

func (c *CustomerService) DoRequest(req *http.Request) (io.ReadCloser, error) {
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.ConsumerKey, c.ConsumerSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func (c *CustomerService) List() ([]Customer, error) {
	return c.Customer.ListByPage(1)
}

func (c *CustomerService) ListByPage(page int) ([]Customer, error) {
	if page <= 0 {
		page = 1
	}

	serviceUrl := fmt.Sprintf("%s/customers", c.apiUrl)

	req, err := http.NewRequest(http.MethodGet, serviceUrl, nil)
	if err != nil {
		return []Customer{}, err
	}

	q := url.Values{}
	q.Add("page", fmt.Sprintf("%d", page))

	req.URL.RawQuery = q.Encode()

	body, err := c.DoRequest(req)
	if err != nil {
		return []Customer{}, err
	}

	var customers []Customer
	if err := json.NewDecoder(body).Decode(&customers); err != nil {
		return []Customer{}, err
	}

	return customers, nil
}

func (c *CustomerService) Create(cm Customer) (Customer, error) {
	serviceUrl := fmt.Sprintf("%s/customers", c.apiUrl)

	bcm, err := json.Marshal(cm)
	if err != nil {
		return Customer{}, err
	}

	req, err := http.NewRequest(http.MethodPost, serviceUrl, bytes.NewReader(bcm))
	if err != nil {
		return Customer{}, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return Customer{}, nil
	}

	var customer Customer
	if err := json.NewDecoder(body).Decode(&customer); err != nil {
		return Customer{}, err
	}

	return customer, nil
}
