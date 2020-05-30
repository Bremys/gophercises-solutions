package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func arcHandler(arc ArcDescription, parsed *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := parsed.Execute(w, arc)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	adventure, err := ParseJSONAdventureFile(os.Args[1])

	if err != nil {
		log.Fatalln("Couldn't parse json to adventure: ", err)
	}

	parsed, err := template.ParseFiles("view/template.html")
	if err != nil {
		log.Fatal("Error parsing html", err)
	}

	for name, arc := range adventure {
		http.HandleFunc("/"+name, arcHandler(arc, parsed))
	}

	http.Handle("/", http.RedirectHandler("/intro", http.StatusFound))
	http.ListenAndServe(":8080", nil)
}
