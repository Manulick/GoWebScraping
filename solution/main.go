package main

import (
    "fmt"

    "github.com/gocolly/colly"
)

var url string = "https://www.liverpool.com.mx/tienda/pdp/apple-iphone-8-retina-4.7-pulgadas-desbloqueado-reacondicionado/1083798323"

type Product struct{
    name string
    price string
}

func main(){
	c := colly.NewCollector()

	c.OnRequest( func(r *colly.Request){
        fmt.Println("visiting ", r.URL)
    })

	c.OnHTML("p", func(e *colly.HTMLElement) {
		name := e.ChildText(".a-product__paragraphDiscountPrice m-0 d-inline")
        fmt.Println(name)
	})


	c.Visit(url)
}