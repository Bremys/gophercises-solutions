package main

import (
	"html/template"
	"log"
	"net/http"
)

type SiteState struct {
	Adjective string
	Gender    string
}

func (state *SiteState) AdjLength() int {
	return len(state.Adjective)
}

func (state *SiteState) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := template.ParseFiles("view/template.html")
	if err != nil {
		log.Fatal("Error parsing html", err)
	}
	err = p.Execute(w, state)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	state := SiteState{
		Adjective: "baaaaaaaaaaad",
		Gender:    "guy",
	}
	http.Handle("/", &state)
	http.ListenAndServe(":8080", nil)
}
