package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/user"
	"net/http"


	"github.com/julienschmidt/httprouter"
)

//RoleHandler handles user related reuqests
type RoleHandler struct {
	userService user.RoleService
}

//NewRoleHandler returns new RoleHandler
func NewRoleHandler(usrService user.RoleService) *RoleHandler {
	return &RoleHandler{userService: usrService}
}

//GetRoles handles GET /users requests
func (uh *RoleHandler) GetRoles(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	users, errs := uh.userService.Roles()

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

//GetRole handles GET users/:id request
func (uh *RoleHandler) GetRole(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	rol, errs := uh.userService.Role(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(rol, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//GetRole handles GET users/byName/:name request
func (uh *RoleHandler) GetRoleByName(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")

	rol, errs := uh.userService.RoleByName(name)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(rol, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//AddRole handles POST user request
func (uh *RoleHandler) AddRole(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	rol := &entities.Role{}

	err := json.Unmarshal(body, rol)

	// fmt.Println(user.Interests)

	if err != nil {
		fmt.Printf("Error %s", err.Error())
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotExtended), http.StatusNotExtended)
		return
	}

	rol, errs := uh.userService.StoreRole(rol)
	fmt.Println(rol)
	if len(errs) > 0 {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(rol, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//UpdateRole handles PUT users/:id requests
func (uh *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	rol, errs := uh.userService.Role(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	_ = json.Unmarshal(body, &rol)

	rol, errs = uh.userService.UpdateRole(rol)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(rol, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

//DeleteRole handles DELETE users/:id requests
func (uh *RoleHandler) DeleteRole(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	deletedRole, errs := uh.userService.DeleteRole(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(deletedRole, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}


