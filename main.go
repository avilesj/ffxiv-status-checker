package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Realm struct {
	name     string
	category string
	status   string
}

func getServerData(s *goquery.Selection, wg *sync.WaitGroup, r []Realm, i int) {
	defer wg.Done()
	name, err := s.Find(".world-list__world_name > p").First().Html()
	if err != nil {
		log.Fatal(err)
	}
	category, err := s.Find(".world-list__world_category > p").First().Html()
	if err != nil {
		log.Fatal(err)
	}
	status, _ := s.Find(".world-list__create_character > i").First().Attr("data-tooltip")
	r[i] = Realm{
		name:     name,
		category: category,
		status:   status,
	}
}

func main() {
	var wg sync.WaitGroup
	res, err := http.Get("https://na.finalfantasyxiv.com/lodestone/worldstatus/")
	if err != nil {
		fmt.Print(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	servers := doc.Find(".world-list__item")
	r := make([]Realm, servers.Length())
	servers.Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go getServerData(s, &wg, r, i)
	})
	wg.Wait()
}
