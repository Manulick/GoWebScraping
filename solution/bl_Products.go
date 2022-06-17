package main

import (
	"GoWebScraping/solution/openapi"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

var urlList = []string{
	"https://laptops.mercadolibre.com.mx/laptops-accesorios/#menu=categories",
	"https://listado.mercadolibre.com.mx/supermercado/bebidas/",
	"https://listado.mercadolibre.com.mx/_Deal_deportes-y-fitness-accesorios",
	"https://www.mercadolibre.com.mx/mas-vendidos/MLM1144",
}

func getUrlList(urls []string) interface{} {
	products := openapi.Products{}
	product := openapi.Product{}

	c := colly.NewCollector(
		colly.Async(),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting ", r.URL)
	})
	c.OnHTML("ol.ui-search-layout a.ui-search-item__group__element", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Data) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "div.ui-pdp-products")
			imgUrl := getDataAttribute(link, "img.ui-pdp-image.ui-pdp-gallery__figure__image", "data-src")
			if err != nil {
				log.Fatal(err)
				return
			}
			product.Description = name
			product.Price = price
			product.OldPrice = oldPrice
			product.ImageURL = imgUrl
			products.Data = append(products.Data, product)
		}
	})
	for _, url := range urls {
		c.Visit(url)
	}
	c.Wait()

	return products
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
