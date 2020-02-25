package scrapper

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type extractedJob struct {
	id string
	jobTitle string
	location string
	salary string
	summary string
}

func Scrape(term string) {
	var baseURL  = "https://www.indeed.com/jobs?q="+term+"&l=Brooklyn%2C+NY"
	var mainC = make(chan []extractedJob)
	var totalJobs []extractedJob
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, mainC)
	}

	for i := 0; i < totalPages; i++ {
		jobs := <-mainC
		totalJobs = append(totalJobs, jobs...)
	}
	//fmt.Println(totalJobs)
	fmt.Println("Done, Extracted all jobs")
}

// Gather jobs in each page
func getPage(page int, baseURL string, mainC chan<-[]extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting: ", pageURL)
	res, _ := http.Get(pageURL)

	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection){
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

//Read each card
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := ClearText(card.Find(".title>a").Text()) // <a> tag 안에 있는 "title" 가져오기
	loc := ClearText(card.Find(".location").Text())  //.을 붙여야해! - class일때...
	salary := ClearText(card.Find(".salaryText").Text())
	summary := ClearText(card.Find(".summary").Text())

	c <- extractedJob{id, title, loc, salary,summary}
}

//Check how many pages in Indeed
func getPages(baseURL string) int {
	pages := 0
	res, _ := http.Get(baseURL)

	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() // how many page links are in there
	})
	fmt.Println(pages)
	return pages

}

func ClearText(text string) string{
	return strings.Join(strings.Fields(strings.TrimSpace(text))," ")
}