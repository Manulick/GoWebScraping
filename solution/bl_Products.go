package main

import (
	"GoWebScraping/solution/openapi"
	"github.com/gocolly/colly/v2"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"sync"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var urlList = []string{
	"https://www.mercadolibre.com.mx/mas-vendidos/MLM1144",
	"https://laptops.mercadolibre.com.mx/laptops-accesorios/#menu=categories",
	"https://listado.mercadolibre.com.mx/supermercado/bebidas/",
	"https://listado.mercadolibre.com.mx/_Deal_deportes-y-fitness-accesorios",
}

var products openapi.Products

func getUrlList(url string, response *openapi.Products) {

	var product openapi.Product

	var mu sync.Mutex

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Default().Println("visiting ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Fatalf("connection with %s LOST", url)
		}
	})

	c.OnHTML("ol.ui-search-layout a.ui-search-item__group__element", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Products) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "div.ui-pdp-price.mt-16.ui-pdp-price--size-large div.ui-pdp-price__second-line span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "s.andes-money-amount.ui-pdp-price__part.ui-pdp-price__original-value.andes-money-amount--previous.andes-money-amount--cents-superscript.andes-money-amount--compact span.andes-money-amount__fraction")
			brand, err := getDataBrand(link, "td.andes-table__column")
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
			product.Brand = brand
			product.ImageURL = imgUrl
			products.Products = append(products.Products, product)
		}
	})

	c.OnHTML("ol.items_container a.promotion-item__link-container", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Products) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "div.ui-pdp-price.mt-16.ui-pdp-price--size-large div.ui-pdp-price__second-line span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "s.andes-money-amount.ui-pdp-price__part.ui-pdp-price__original-value.andes-money-amount--previous.andes-money-amount--cents-superscript.andes-money-amount--compact span.andes-money-amount__fraction")
			brand, err := getDataBrand(link, "td.andes-table__column")
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
			product.Brand = brand
			product.ImageURL = imgUrl
			products.Products = append(products.Products, product)
		}
	})

	c.OnHTML("ol.ui-search-layout.ui-search-layout--grid a.ui-search-result__content.ui-search-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(products.Products) < 10 {
			name, err := getData(link, "h1.ui-pdp-title")
			price, err := getData(link, "div.ui-pdp-price.mt-16.ui-pdp-price--size-large div.ui-pdp-price__second-line span.andes-money-amount__fraction")
			oldPrice, err := getData(link, "s.andes-money-amount.ui-pdp-price__part.ui-pdp-price__original-value.andes-money-amount--previous.andes-money-amount--cents-superscript.andes-money-amount--compact span.andes-money-amount__fraction")
			brand, err := getDataBrand(link, "td.andes-table__column")
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
			product.Brand = brand
			product.ImageURL = imgUrl
			products.Products = append(products.Products, product)
		}
	})

	c.Visit(url)
	for _, prod := range products.Products {
		response.Products = append(response.Products, prod)
	}

}

func getData(url string, element string) (string, error) {
	response := ""
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

func getDataBrand(url string, element string) (string, error) {
	response := ""
	i := 0
	var productError error
	c := colly.NewCollector()

	c.OnHTML(element, func(e *colly.HTMLElement) {
		e.ForEach("span.andes-table__column--value", func(_ int, element *colly.HTMLElement) {
			if i == 0 {
				response = e.Text
				i++
			}
		})
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

func createFile(productsList openapi.Products) {
	log.Default().Println("creating file...")
	file, _ := json.Marshal(productsList)
	err := ioutil.WriteFile("ProductList.json", file, 0644)
	if err != nil {
		log.Fatalf(" something got wrong with %s creating", file)
	}
}
