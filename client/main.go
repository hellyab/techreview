package main

import (
	"bytes"
	"fmt"
	"html/template"

	// "io/ioutil"
	// "log"

	"encoding/json"
	"net/http"
	"net/url"

	"github.com/hellyab/techreview/client/data"
	"github.com/hellyab/techreview/delivery/http/handler"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/questions", allQuestions)
	mux.HandleFunc("/userentry", userEntry)
	mux.HandleFunc("/newuser", registerUser)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/article", articleHandler)

	http.ListenAndServe(":8080", mux)

}

func allQuestions(w http.ResponseWriter, _ *http.Request) {
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
	dest := "http://localhost:8181/user/"
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
		req, err := http.NewRequest("POST", dest, bytes.NewBuffer(newUser))
		if err != nil {
			http.Redirect(w, r, "/userentry", http.StatusNotAcceptable)
		}
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}

		fileName, err := handler.UploadImage(w, r, file, header)
		URLSafeFileName := url.PathEscape(fileName)

		errorResponseJSON := fmt.Sprintf(`{
			"success" : 0,
			"file": {
				"url" : ""
			}
			}`)
		successResponseJSON := fmt.Sprintf(`{
			"success" : 1,
			"file": {
				"url" : "http://localhost:8080/assets/images/%s"
			}
			}`, URLSafeFileName)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(errorResponseJSON))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(successResponseJSON))

	} else if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "upload-test.html", nil)
	}
}

func userEntry(w http.ResponseWriter, _ *http.Request) {
	templates.ExecuteTemplate(w, "user-entry.html", nil)
}

func articleHandler(w http.ResponseWriter, _ *http.Request) {
	templates.ExecuteTemplate(w, "editor.html", nil)
}
