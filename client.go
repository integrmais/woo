package woo

type Client struct {
	BaseUrl        string
	apiUrl         string
	VersionApi     string
	ConsumerKey    string
	ConsumerSecret string

	Product  *ProductService
	Customer *CustomerService
}
