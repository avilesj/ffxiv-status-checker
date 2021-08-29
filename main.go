package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://na.finalfantasyxiv.com/lodestone/worldstatus/")
	if err != nil {
		fmt.Print(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".world-list__item").Each(func(i int, s *goquery.Selection) {
		name, err := s.Find(".world-list__world_name > p").First().Html()
		if err != nil {
			log.Fatal(err)
		}
		category, err := s.Find(".world-list__world_category > p").First().Html()
		if err != nil {
			log.Fatal(err)
		}
		status, _ := s.Find(".world-list__create_character > i").First().Attr("data-tooltip")
		fmt.Printf("Realm: %s\n", name)
		fmt.Printf("Status: %s\n", status)
		fmt.Printf("Category: %s\n", category)
	})
}
