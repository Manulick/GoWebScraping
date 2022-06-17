package main

import (
	"GoWebScraping/solution/openapi"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"sync"
)

var urlList = "https://laptops.mercadolibre.com.mx/laptops-accesorios/#menu=categories"

func getUrlList(url string) interface{} {
	products := openapi.Products{}
	product := openapi.Product{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting ", r.URL)
	})
	c.OnHTML("ol.ui-search-layout", func(e *colly.HTMLElement) {
		e.ForEach("a.ui-search-item__group__element", func(i int, e *colly.HTMLElement) {
			link := e.Attr("href")

			log.Println(link)

			if i < 10 {

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
	})
	c.Visit(url)

	return products
}

func getData(url string, element string) (string, error) {
	var response string
	var productError error
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := colly.NewCollector()

	defer wg.Done()
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
