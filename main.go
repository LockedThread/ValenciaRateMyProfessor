package main

import (
	"fmt"
	"github.com/LockedThread/ValenciaRateMyProfessor/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strconv"
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
		if i%20 == 0 {
			return
		}

		if s.HasClass("a") {
			return
		}

		fmt.Printf("i=%d", i)
		fmt.Printf("\nlength= %d", s.Length)
		if s.Length() > 1 {
			return
		}

		fmt.Printf("\ntext=%s", s.Text())

		var attrCount = 0

		find := s.Find("a")
		if find != nil {
			if len(find.Nodes) > 0 {
				if len(find.Nodes[0].Attr) > 0 {
					attrCount = len(find.Nodes[0].Attr)
				}
			}
		}

		fmt.Printf("\nattrCount=%d", attrCount)

		if attrCount > 1 {
			return
		}

		fmt.Printf("\n")
		credits, err := strconv.Atoi(s.Nodes[5].Data)
		if err != nil {
			log.Fatal(err)
		}

		course := model.Course{
			CRN:        s.Nodes[0].Data,
			Subject:    s.Nodes[1].Data,
			Course:     s.Nodes[2].Data,
			CampusID:   s.Nodes[4].Data,
			Title:      s.Nodes[6].Data,
			Credits:    credits,
			Honors:     false,
			Instructor: s.Nodes[15].Data,
		}
		fmt.Printf("%s", course)
	})
}
