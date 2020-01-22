package main

import (
	"bytes"
	"fmt"
	"html/template"

	// "io/ioutil"
	// "log"
	"encoding/json"
	"net/http"

	"github.com/hellyab/techreview/client/data"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/questions", allQuestions)
	mux.HandleFunc("/userentry", userEntry)
	mux.HandleFunc("/newuser", registerUser)

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

	templates.ExecuteTemplate(w, "question.html", Questions)

}

func registerUser(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8181/user/"
	if r.Method != http.MethodPost {
		fmt.Println("Not post!")
	} else {
		user := struct {
			Username   string
			FirstName  string
			MiddleName string
			LastName   string
			Email      string
			Password   string
			Interests  string
		}{
			r.FormValue("username"),
			r.FormValue("firstName"),
			"",
			r.FormValue("lastName"),
			r.FormValue("email"),
			r.FormValue("password"),
			"[]",
		}
		newUser, err := json.Marshal(user)

		if err != nil {
			http.Redirect(w, r, "/userentry", http.StatusNotAcceptable)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(newUser))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {

			http.Redirect(w, r, "/questions", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/userentry", http.StatusSeeOther)
		}

		fmt.Println("response Status:", resp.Status)
		// fmt.Println("response Headers:", resp.Header)
		// body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println("response Body:", string(body))

	}
}

func userEntry(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "user-entry.html", nil)
}
