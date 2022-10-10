package woo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProductService Client

type Dimensions struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type Categories []struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductImages []struct {
	Id        int64  `json:"id"`
	SourceUrl string `json:"src"`
	Name      string `json:"name"`
	Alt       string `json:"alt"`
}

type Product struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name"`
	Slug              string  `json:"slug"`
	Permalink         string  `json:"permalink"`
	DateCreated       string  `json:"date_created"`
	DateCreatedGmt    string  `json:"date_created_gmt"`
	DateModified      string  `json:"date_modified"`
	DateModifiedGmt   string  `json:"date_modified_gmt"`
	Type              string  `json:"type"`
	Status            string  `json:"status"`
	Featured          bool    `json:"featured"`
	CatalogVisibility string  `json:"catalog_visibility"`
	Description       string  `json:"description"`
	SKU               string  `json:"sku"`
	Price             string  `json:"price"`
	RegularPrice      string  `json:"regular_price"`
	SalePrice         string  `json:"sale_price"`
	DateOnSaleFrom    string  `json:"date_on_sale_from"`
	DateOnSaleFromGmt string  `json:"date_on_sale_from_gmt"`
	PriceHTML         string  `json:"price_html"`
	OnSale            bool    `json:"on_sale"`
	Purchasale        bool    `json:"purchasale"`
	TotalSales        int     `json:"total_sales"`
	Virtual           bool    `json:"virtual"`
	Downloadable      bool    `json:"downloadable"`
	TaxStatus         string  `json:"tax_status"`
	TaxClass          string  `json:"tax_class"`
	ManageStock       bool    `json:"manage_stock"`
	StockQuantity     int     `json:"stock_quantity"`
	StockStatus       string  `json:"stock_status"`
	BackOrders        string  `json:"back_orders"`
	BackOrdersAllowed bool    `json:"back_orders_allowed"`
	SoldIndividually  bool    `json:"sold_individually"`
	Weight            string  `json:"weight"`
	RelatedIds        []int64 `json:"related_ids"`
	ParentId          int64   `json:"parent_id"`
	Dimensions        `json:"dimensions"`
	Categories        `json:"categories"`
	ProductImages     `json:"images"`
}

func (s *ProductService) List() ([]Product, error) {
	apiUrl := fmt.Sprintf("%s/products", s.apiUrl)

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return []Product{}, err
	}

	req.SetBasicAuth(s.ConsumerKey, s.ConsumerSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Product{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Product{}, err
	}

	var products []Product
	if err := json.Unmarshal(body, &products); err != nil {
		return []Product{}, err
	}

	return products, nil
}
