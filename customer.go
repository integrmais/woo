package woo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CustomerService Client

type Customer struct {
	Id               int64  `json:"id,omitempty"`
	AvatarURL        string `json:"avatar_url"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	IsPayingCustomer bool   `json:"is_paying_customer"`
	DateCreated      string `json:"date_created"`
	DateModified     string `json:"date_modified"`
	ShippingAddress  `json:"shipping"`
	BillingAddress   `json:"billing"`
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
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func (c *CustomerService) List() ([]Customer, error) {
	serviceUrl := fmt.Sprintf("%s/customers", c.apiUrl)

	req, err := http.NewRequest(http.MethodGet, serviceUrl, nil)
	if err != nil {
		return []Customer{}, err
	}

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

func (c *CustomerService) Create() error {
	return nil
}