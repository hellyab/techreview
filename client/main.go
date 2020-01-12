package main

import (
	"fmt"
	"html/template"

	// "io/ioutil"
	// "log"
	// "encoding/json"
	"net/http"

	"github.com/hellyab/techreview/client/data"
)

var questionTemplate = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/questions", allQuestions)

	http.ListenAndServe(":8080", mux)

}

func allQuestions(w http.ResponseWriter, r *http.Request) {
	Questions, err := data.FetchQuestions()

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Println("Error Occured")
		//TODO Add error layout
		//tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	questionTemplate.ExecuteTemplate(w, "question.html", Questions)

}
