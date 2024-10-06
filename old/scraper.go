package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Name, Price string
}

func main() {

	link := flag.String("link", "", "a link")
	flag.Parse()

	fmt.Println(*link)

	c := colly.NewCollector(
		colly.AllowedDomains("www.ozon.ru"),
	)

	/* 	c.OnResponse(func(r *colly.Response) {
		err := os.WriteFile("ozon_page.html", r.Body, 0644)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("HTML content saved to ozon_page.html")
	}) */

	c.OnHTML("[data-state]", func(e *colly.HTMLElement) {
		var product Product

		dataState := e.Attr("data-state")

		if strings.Contains(dataState, "cardPrice") {
			var data map[string]interface{}

			json.Unmarshal([]byte(dataState), &data)

			if price, ok := data["cardPrice"].(string); ok {
				product.Price = strings.TrimSpace(strings.TrimSuffix(price, " ₽"))
				fmt.Println("Price found:", product.Price)
			}
		}
	})

	c.Visit(*link)
}
