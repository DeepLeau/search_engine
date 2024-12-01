package main

import (
 "fmt"
 "net/http"
 "github.com/gorilla/mux"
)

type search_query struct {
    url string;
	title string;
	brief_description string;
}

func main() {
	r := mux.NewRouter()
   
	r.HandleFunc("/users/{search_query}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		text := vars["search_query"]
		fmt.Fprintf(w, "Search query: %s", text)
	   }).Methods("GET")
   
	http.ListenAndServe(":8080", r)
   }