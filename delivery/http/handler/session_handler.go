package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/user"
	"github.com/julienschmidt/httprouter"
)

//SessionHandler handles user related reuqests
type SessionHandler struct {
	userService user.SessionService
}

//NewSessionHandler returns new SessionHandler
func NewSessionHandler(usrService user.SessionService) *SessionHandler {
	return &SessionHandler{userService: usrService}
}

//GetSession handles GET users/:id request
func (uh *SessionHandler) GetSession(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	usr, errs := uh.userService.Session(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(usr, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//AddSession handles POST user request
func (uh *SessionHandler) AddSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	usr := &entities.Session{}

	err := json.Unmarshal(body, usr)

	// fmt.Println(user.Interests)

	if err != nil {
		fmt.Printf("Error %s", err.Error())
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotExtended), http.StatusNotExtended)
		return
	}

	usr, errs := uh.userService.StoreSession(usr)
	fmt.Println(usr)
	if len(errs) > 0 {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(usr, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}


//DeleteSession handles DELETE users/:id requests
func (uh *SessionHandler) DeleteSession(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	deletedSession, errs := uh.userService.DeleteSession(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(deletedSession, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	fmt.Println("session deleted")
	return

}
