package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("example.html")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", "example.html")
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("tr td").Each(func(i int, s *goquery.Selection) {
		s.Nodes[0]
		fmt.Printf("%s", s.Text())
	})
}
