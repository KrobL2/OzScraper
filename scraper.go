package main

import (
	"encoding/csv"
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
	var products []Product

	// OnHTML callback
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// fmt.Print("Scraping: ", e)

		// initialize a new Product instance
		product := Product{}

		// Find the price using a more general selector
		e.ForEach("span", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			if strings.Contains(text, "₽") && strings.Count(text, " ") <= 2 {
				product.Price = strings.TrimSpace(strings.TrimSuffix(text, " ₽"))
				return
			}
		})

		// scrape the target data
		product.Url = e.ChildAttr("a", "href")
		product.Name = e.ChildText(".product-name")

		// add the product instance with scraped data to the list of products
		products = append(products, product)
	})

	// store the data to a CSV after extraction
	c.OnScraped(func(r *colly.Response) {

		fmt.Println("output", products)

		return

		// open the CSV file
		file, err := os.Create("products.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		// initialize a file writer
		writer := csv.NewWriter(file)

		// write the CSV headers
		headers := []string{
			"Url",
			"Name",
			"Price",
		}

		writer.Write(headers)

		// write each product as a CSV row
		for _, product := range products {
			// convert a Product to an array of strings
			record := []string{
				product.Url,
				product.Name,
				product.Price,
			}

			// add a CSV record to the output file
			writer.Write(record)
		}
		defer writer.Flush()
	})

	// open the target URL
	c.Visit("https://www.ozon.ru/product/odeyalo-ikea-stjarnbracka-teploe-150x200-sm-1648654068/?avtc=1&avte=2&avts=1727299349")

}
