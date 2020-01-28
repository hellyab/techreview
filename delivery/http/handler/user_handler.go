package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/user"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//UserHandler handles user related reuqests
type UserHandler struct {
	userService user.UserService
}

//NewUserHandler returns new UserHandler
func NewUserHandler(usrService user.UserService) *UserHandler {
	return &UserHandler{userService: usrService}
}

//GetUsers handles GET /users requests
func (uh *UserHandler) GetUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	users, errs := uh.userService.Users()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

//GetUser handles GET users/:id request
func (uh *UserHandler) GetUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	usr, errs := uh.userService.User(id)
	usr.Password = ""	// don expose password

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
	_, _ = w.Write(output)
	return

}

//GetUserByUsername handles GET userbyusername/:username requests
func (uh *UserHandler) GetUserByUsername(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")

	usr, errs := uh.userService.UserByUsername(username)

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
	_, _ = w.Write(output)
	return

}

//AddUser handles POST user request
func (uh *UserHandler) AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	usr := &entities.User{}

	err := json.Unmarshal(body, usr)

	// fmt.Println(user.Interests)

	if err != nil {
		fmt.Printf("Error %s", err.Error())
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotExtended), http.StatusNotExtended)
		return
	}

	usr, errs := uh.userService.StoreUser(usr)
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

//UpdateUser handles PUT users/:id requests
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	usr, errs := uh.userService.User(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	_ = json.Unmarshal(body, &usr)

	usr, errs = uh.userService.UpdateUser(usr)

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
	_, _ = w.Write(output)
	return
}

//DeleteUser handles DELETE users/:id requests
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	deletedUser, errs := uh.userService.DeleteUser(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(deletedUser, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
