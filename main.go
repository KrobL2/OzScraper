package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	htmlContent, err := os.ReadFile("OZON.html")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	doc, err := html.Parse(strings.NewReader(string(htmlContent)))
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return
	}

	orders := extractAllOrders(doc)

	for _, order := range orders {
		fmt.Printf("Order Number: %s, Price: %s\n", order.Number, order.Price)
	}
}

func extractAllOrders(n *html.Node) []Order {
	orderMap := make(map[string]Order)

	var f func(*html.Node)
	f = func(n *html.Node) {
		price, orderNumber := extractOrderInfo(n)

		if price != "" && orderNumber != "" {
			orderMap[orderNumber] = Order{Number: orderNumber, Price: price}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	// Convert map to slice
	var uniqueOrders []Order
	for _, order := range orderMap {
		uniqueOrders = append(uniqueOrders, order)
	}

	return uniqueOrders
}
