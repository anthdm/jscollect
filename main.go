package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var domain string
	flag.StringVar(&domain, "d", "", "The target domain to collect js files from <proto:port>")
	flag.Parse()

	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	urls := []string{}
	doc.Find("script").Each(func(i int, e *goquery.Selection) {
		src, ok := e.Attr("src")
		if ok {
			urls = append(urls, src)
		}
	})

	for _, s := range urls {
		fmt.Println(s)
	}
}
