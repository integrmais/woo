package woo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/integrmais/woo"
)

func TestListProducts(t *testing.T) {
	var consumerKey, consumerSecret string

	serverMock := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		consumerKey, consumerSecret, _ = req.BasicAuth()
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`[
			{
				"id": 799,
				"name": "Ship Your Idea",
				"slug": "ship-your-idea-22",
				"permalink": "https://example.com/product/ship-your-idea-22/",
				"date_created": "2017-03-23T17:03:12",
				"date_created_gmt": "2017-03-23T20:03:12",
				"date_modified": "2017-03-23T17:03:12",
				"date_modified_gmt": "2017-03-23T20:03:12",
				"type": "variable",
				"status": "publish",
				"featured": false,
				"catalog_visibility": "visible",
				"description": "<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>\n",
				"short_description": "<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.</p>\n",
				"sku": "",
				"price": "",
				"regular_price": "",
				"sale_price": "",
				"date_on_sale_from": null,
				"date_on_sale_from_gmt": null,
				"date_on_sale_to": null,
				"date_on_sale_to_gmt": null,
				"price_html": "",
				"on_sale": false,
				"purchasable": false,
				"total_sales": 0,
				"virtual": false,
				"downloadable": false,
				"downloads": [],
				"download_limit": -1,
				"download_expiry": -1,
				"external_url": "",
				"button_text": "",
				"tax_status": "taxable",
				"tax_class": "",
				"manage_stock": false,
				"stock_quantity": null,
				"stock_status": "instock",
				"backorders": "no",
				"backorders_allowed": false,
				"backordered": false,
				"sold_individually": false,
				"weight": "",
				"dimensions": {
					"length": "",
					"width": "",
					"height": ""
				},
				"shipping_required": true,
				"shipping_taxable": true,
				"shipping_class": "",
				"shipping_class_id": 0,
				"reviews_allowed": true,
				"average_rating": "0.00",
				"rating_count": 0,
				"related_ids": [
					31,
					22,
					369,
					414,
					56
				],
				"upsell_ids": [],
				"cross_sell_ids": [],
				"parent_id": 0,
				"purchase_note": "",
				"categories": [
					{
						"id": 9,
						"name": "Clothing",
						"slug": "clothing"
					},
					{
						"id": 14,
						"name": "T-shirts",
						"slug": "t-shirts"
					}
				],
				"tags": [],
				"images": [
					{
						"id": 795,
						"date_created": "2017-03-23T14:03:08",
						"date_created_gmt": "2017-03-23T20:03:08",
						"date_modified": "2017-03-23T14:03:08",
						"date_modified_gmt": "2017-03-23T20:03:08",
						"src": "https://example.com/wp-content/uploads/2017/03/T_4_front-11.jpg",
						"name": "",
						"alt": ""
					},
					{
						"id": 796,
						"date_created": "2017-03-23T14:03:09",
						"date_created_gmt": "2017-03-23T20:03:09",
						"date_modified": "2017-03-23T14:03:09",
						"date_modified_gmt": "2017-03-23T20:03:09",
						"src": "https://example.com/wp-content/uploads/2017/03/T_4_back-10.jpg",
						"name": "",
						"alt": ""
					},
					{
						"id": 797,
						"date_created": "2017-03-23T14:03:10",
						"date_created_gmt": "2017-03-23T20:03:10",
						"date_modified": "2017-03-23T14:03:10",
						"date_modified_gmt": "2017-03-23T20:03:10",
						"src": "https://example.com/wp-content/uploads/2017/03/T_3_front-10.jpg",
						"name": "",
						"alt": ""
					},
					{
						"id": 798,
						"date_created": "2017-03-23T14:03:11",
						"date_created_gmt": "2017-03-23T20:03:11",
						"date_modified": "2017-03-23T14:03:11",
						"date_modified_gmt": "2017-03-23T20:03:11",
						"src": "https://example.com/wp-content/uploads/2017/03/T_3_back-10.jpg",
						"name": "",
						"alt": ""
					}
				],
				"attributes": [
					{
						"id": 6,
						"name": "Color",
						"position": 0,
						"visible": false,
						"variation": true,
						"options": [
							"Black",
							"Green"
						]
					},
					{
						"id": 0,
						"name": "Size",
						"position": 0,
						"visible": true,
						"variation": true,
						"options": [
							"S",
							"M"
						]
					}
				],
				"default_attributes": [],
				"variations": [],
				"grouped_products": [],
				"menu_order": 0,
				"meta_data": [],
				"_links": {
					"self": [
						{
							"href": "https://example.com/wp-json/wc/v3/products/799"
						}
					],
					"collection": [
						{
							"href": "https://example.com/wp-json/wc/v3/products"
						}
					]
				}
			}]`))
	}))

	defer func() { serverMock.Close() }()

	c := woo.NewClient(serverMock.URL, versionMock, consumerKeyMock, consumerSecretMock)

	products, err := c.Product.List()

	if err != nil {
		t.Fatalf("Expected empty error, got %v", err.Error())
	}

	if consumerKey != consumerKeyMock || consumerSecret != consumerSecretMock {
		t.Fatalf("Expected correct basic auth, got %s:%s", consumerKey, consumerSecret)
	}

	if len(products) == 0 {
		t.Fatalf("Expected 1 products, got %v", len(products))
	}

	if products[0].Id == 0 {
		t.Fatalf("Expected 799 product id, got %v", products[0].Id)
	}
}
