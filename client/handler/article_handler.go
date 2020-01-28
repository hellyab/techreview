package handler

import "net/http"

func (uh *UserHandler)PostArticle(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		uh.tmpl.ExecuteTemplate(w,"editor.html", http.StatusSeeOther)

	}
}