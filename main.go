package main

import (
	"log"
	sc_server "github.com/achievermina/IndeedWebScrapping-API/scrapper"
	"google.golang.org/grpc"
	"net"
	"context"
	"strings"
	"github.com/achievermina/cloneIndeed/scrapper"
)


type Server struct {
	sc_server.UnimplementedJobServiceServer
}

// proto file 에는 첫문장 capital 아니더라도 여기서는 Must 첫 letter capital
func (s *Server) Search(ctx context.Context, req *sc_server.SearchRequest) (*sc_server.SearchResponse, error) {
	log.Printf("get tasks search term: %v", req.Term)
	var jobs []*sc_server.JobObject
	//job1 := &sc_server.JobObject{
	//	Id: "1",
	//	Title: "programmer",
	//	Location: "NY",
	//	Salary: "$1000000",
	//	Summary: "ASAP",
	//}
	//jobs = append(jobs, job1)
	//job2 := &sc_server.JobObject{
	//	Id: "2",
	//	Title: "programmer2",
	//	Location: "NY",
	//	Salary: "$1000000",
	//	Summary: "ASAP",
	//}
	//jobs = append(jobs, job2)

	scrapperReqTerm := strings.ToLower(scrapper.ClearText(req.Term))
	extractedJobs := sc_server.Scrape(scrapperReqTerm)

	for  i := 0; i < len(extractedJobs); i++ {
		cur := extractedJobs[i]
		job := &sc_server.JobObject{
			Id: cur.Id,
			Title: cur.JobTitle,
			Location: cur.Location,
			Salary: cur.Salary,
			Summary: cur.Summary,
		}
		jobs = append(jobs,job)
	}
	return &sc_server.SearchResponse{Jobs: jobs}, nil
}

func main() {
	//gRPC
	grpcServer := grpc.NewServer()
	sc_server.RegisterJobServiceServer(grpcServer, &Server{})
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
	}
	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}


//https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
//
//func home(w http.ResponseWriter, r *http.Request) {
//
//	term, ok := r.URL.Query()["term"]  //query string --> GET
//
//	if !ok || len(term[0]) < 1 {
//		log.Println("Url Param 'term' is missing")
//		return
//	}
//
//	term := "python"
//	fmt.Println(term)
//
//	searchTerm := strings.ToLower(scrapper.ClearText(term))
//	res := scrapper.Scrape(searchTerm)
//	fmt.Println(res)
//
//}
//
//
//
//func main() {
//	term := "python"
//	fmt.Println(term)
//	searchTerm := strings.ToLower(scrapper.ClearText(term))
//	res := scrapper.Scrape(searchTerm)
//	// http response --> excel download
//	//http api response excel
//
//	r := mux.NewRouter()
//	r.HandleFunc("/", home)
//	log.Fatalln(http.ListenAndServe(":8080", r))
//
//	err = json.NewDecoder(r.Body).Decode(&res)
//	newsJson, err := json.Marshal(news)
//	if err != nil {
//		panic(err)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusAccepted)
//	w.Write(newsJson)
//
//
//	//http.HandleFunc("/", home)
//	//http.ListenAndServe(":8080", nil)
//}
//
//type event struct {
//	userID          string `json:"userID"`
//	Title       string `json:"Title"`
//}
//
//type allEvents []event
//
//var events = allEvents{
//	{
//		userID:          "1",
//		Title:       "Introduction to Golang",
//	},
//}
//
//var searchterm string
//var allsearch []string
//
//func home(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(`{"message": "get called"}`))
//}
//
//func createEvent(w http.ResponseWriter, r *http.Request) {
//	var newEvent event
//	//var newSearch string
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "please type the word you are searching for")
//	}
//
//	json.Unmarshal(reqBody, &newEvent)
//	//allsearch = append(allsearch, newSearch)
//	events = append(events, newEvent)
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(newEvent)
//}
//
//func main() {
//	//term :="python"
////	//scrapper.Scrape(term)
////
//	r := mux.NewRouter()
//	r.HandleFunc("/", home)
//	r.HandleFunc("/event", createEvent).Methods("GET")
//	log.Fatalln(http.ListenAndServe(":8080", r))
//
//}
