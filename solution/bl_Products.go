package main

import (
	"GoWebScraping/solution/openapi"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"sync"
)

var urlList = []string{
	"https://laptops.mercadolibre.com.mx/laptops-accesorios/#menu=categories",
	"https://listado.mercadolibre.com.mx/supermercado/bebidas/",
	"https://listado.mercadolibre.com.mx/_Deal_deportes-y-fitness-accesorios",
	"https://www.mercadolibre.com.mx/mas-vendidos/MLM1144",
}

func getUrlList(url string, response *openapi.Response) {

	var products openapi.Products
	var product openapi.Product

	var mu sync.Mutex

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting ", r.URL)
	})

	c.OnHTML("ol.ui-search-layout a.ui-search-item__group__element", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Products) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "div.ui-pdp-products")
			imgUrl := getDataAttribute(link, "img.ui-pdp-image.ui-pdp-gallery__figure__image", "data-src")
			if err != nil {
				mu.Lock()
				log.Fatal(err)
				mu.Unlock()
				return
			}
			product.Description = name
			product.Price = price
			product.OldPrice = oldPrice
			product.ImageURL = imgUrl
			products.Products = append(products.Products, product)
		}
	})

	c.OnHTML("ol.items_container a.promotion-item__link-container", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Products) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "div.ui-pdp-products")
			imgUrl := getDataAttribute(link, "img.ui-pdp-image.ui-pdp-gallery__figure__image", "data-src")
			if err != nil {
				mu.Lock()
				log.Fatal(err)
				mu.Unlock()
				return
			}
			product.Description = name
			product.Price = price
			product.OldPrice = oldPrice
			product.ImageURL = imgUrl
			products.Products = append(products.Products, product)
		}
	})

	c.OnHTML("ol.ui-search-layout.ui-search-layout--grid a.ui-search-result__content.ui-search-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Products) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "div.ui-pdp-products")
			imgUrl := getDataAttribute(link, "img.ui-pdp-image.ui-pdp-gallery__figure__image", "data-src")
			if err != nil {
				mu.Lock()
				log.Fatal(err)
				mu.Unlock()
				return
			}
			product.Description = name
			product.Price = price
			product.OldPrice = oldPrice
			product.ImageURL = imgUrl
			products.Products = append(products.Products, product)
		}
	})

	c.Visit(url)

	response.Data = append(response.Data, products)
}

func getData(url string, element string) (string, error) {
	var response string
	var productError error
	c := colly.NewCollector()

	c.OnHTML(element, func(e *colly.HTMLElement) {
		attribute, err := e.DOM.Html()
		if err != nil {
			productError = err
			return
		}
		response = attribute
	})
	c.Visit(url)

	return response, productError
}

func getDataAttribute(url string, element string, attr string) string {
	var response string
	c := colly.NewCollector()

	c.OnHTML(element, func(e *colly.HTMLElement) {
		attribute := e.Attr(attr)
		response = attribute
	})
	c.Visit(url)
	return response
}
