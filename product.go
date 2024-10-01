package main

import (
	"strings"

	"golang.org/x/net/html"
)

type Order struct {
	Number string
	Price  string
}

func extractOrderInfo(n *html.Node) (string, string) {
	var price, orderNumber string
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "div":
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "f9w_14" {
						price = extractPrice(n)
					}

					if a.Key == "class" && a.Val == "w8f_14" {
						orderNumber = extractOrderNumber(n)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return price, orderNumber
}

func extractPrice(n *html.Node) string {
	var price string
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "w9f_14" {
					price = strings.TrimSpace(n.FirstChild.Data)
					price = strings.TrimSuffix(price, " ₽")
					return
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return price
}

func extractOrderNumber(n *html.Node) string {
	var orderNumber string
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "fw9_14" {
					orderNumber = n.FirstChild.Data
					return
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return orderNumber
}
