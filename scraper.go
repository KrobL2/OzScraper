package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Name, Price string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ozon.ru"),
	)

	// initialize the slice of structs that will contain the scraped data

	c.OnResponse(func(r *colly.Response) {
		err := os.WriteFile("ozon_page.html", r.Body, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("HTML content saved to ozon_page.html")
	})

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

	// open the target URL
	c.Visit("https://www.ozon.ru/product/odeyalo-ikea-stjarnbracka-teploe-150x200-sm-1648654068/")

}
