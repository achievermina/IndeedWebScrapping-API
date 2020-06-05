package scrapper

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/achievermina/IndeedWebScrapping-API/scrapper/scrapperpb"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type extractedJob struct {
	Id string  `json:"ID"`
	JobTitle string  `json:"Title"`
	Location string  `json:"Location"`
	Salary string  `json:"Salary"`
	Summary string  `json:"Description"`
}

// proto file 에는 첫문장 capital 아니더라도 여기서는 Must 첫 letter capital
func (s *Server) Search(ctx context.Context, req *scrapperpb.SearchRequest) (*scrapperpb.SearchResponse, error) {
	log.Printf("get tasks search term: %v", req.Term)
	var jobs []*scrapperpb.JobObject

	scrapperReqTerm := strings.ToLower(ClearText(req.Term))
	extractedJobs := Scrape(scrapperReqTerm)

	for  i := 0; i < len(extractedJobs); i++ {
		cur := extractedJobs[i]
		job := &scrapperpb.JobObject{
			Id: cur.Id,
			Title: cur.JobTitle,
			Location: cur.Location,
			Salary: cur.Salary,
			Summary: cur.Summary,
		}
		jobs = append(jobs,job)
	}
	return &scrapperpb.SearchResponse{Jobs: jobs}, nil
}

func Scrape(term string) []extractedJob{
	var baseURL  = "https://www.indeed.com/jobs?q="+term+"&l=Brooklyn%2C+NY"
	var mainC = make(chan []extractedJob)
	var totalJobs []extractedJob
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, mainC)
	}

	for i := 0; i < 1; i++ {
		jobs := <-mainC
		totalJobs = append(totalJobs, jobs...)
	}

	log.Println("Done, Extracted all jobs")
	return totalJobs
}

// Gather jobs in each page
func getPage(page int, baseURL string, mainC chan<-[]extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	log.Println("Requesting: ", pageURL)
	res, err := http.Get(pageURL)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
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
	res, err := http.Get(baseURL)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() // how many page links are in there
	})
	log.Println(pages)
	return pages

}

func ClearText(text string) string{
	return strings.Join(strings.Fields(strings.TrimSpace(text))," ")
}

func checkError(err error){
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response){
	if res.StatusCode != 200 {
		log.Fatalln("Request failed: ", res.StatusCode)
	}
}
