package handler

import (
	"bytes"
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"html/template"
	"net/http"
	"net/url"
	_ "strconv"
	"strings"
	_ "strings"

	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/form"
	"github.com/hellyab/techreview/permission"
	"github.com/hellyab/techreview/rtoken"
	"github.com/hellyab/techreview/session"
	"github.com/hellyab/techreview/user"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handler handles user related requests
type UserHandler struct {
	tmpl           *template.Template
	userService    user.UserService //TODO remove this
	sessionService user.SessionService //TODO remove this
	userSess       *entities.Session
	loggedInUser   *entities.User
	userRole       user.RoleService //TODO remove this
	csrfSignKey    []byte
}

type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

// NewUserHandler returns new UserHandler object
func NewUserHandler(t *template.Template, usrServ user.UserService,
	sessServ user.SessionService, uRole user.RoleService,
	usrSess *entities.Session, csKey []byte) *UserHandler {
	return &UserHandler{tmpl: t, userService: usrServ, sessionService: sessServ,
		userRole: uRole, userSess: usrSess, csrfSignKey: csKey}
}

// Authenticated checks if a user is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			fmt.Println("user is not logged in")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		fmt.Println("context stored",ctx, uh.userSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Authorized checks if a user has proper authority to access a give route
func (uh *UserHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInUser == nil {

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		roles, errs := uh.userService.UserRoles(uh.loggedInUser)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// Login hanldes the GET/POST /login requests
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	userByUsernameDest := "http://localhost:8181/users/username/"
	sessionDest := "http://localhost:8181/sessions"
	userInQuestion := entities.User{}
	userSession := entities.Session{}
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	fmt.Println("Login called successfully")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		fmt.Println("Login get called successfully")

		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
		return
	}

	if r.Method == http.MethodPost {
		fmt.Println("Login post called successfully")

		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		fmt.Println("Form Parsed successfully")

		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		//usr, errs := uh.userService.UserByUsername(r.FormValue("loginUsername"))
		username := r.FormValue("loginUsername")
		//fmt.Println(username)
		userResp, err := http.Get(userByUsernameDest + username)
		//fmt.Println(userResp)
		if err!=nil {
			fmt.Println("Error here: not found")
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
			return
		}
		size:= userResp.ContentLength
		body :=  make ([]byte, size)
		//fmt.Println(body)
		userResp.Body.Read(body)
		//fmt.Println(body)

		errJson := json.Unmarshal(body, &userInQuestion)

		if errJson!=nil{
			fmt.Println(errJson.Error())
			loginForm.VErrors.Add("generic", "Username or password")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
			fmt.Println("Unable to get body from client")
			return
		}
		

		err = bcrypt.CompareHashAndPassword([]byte(userInQuestion.Password), []byte(r.FormValue("loginPassword")))
		if err == bcrypt.ErrMismatchedHashAndPassword {

			fmt.Println("passwords do not match")
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
			return
		}
		fmt.Println("PASS N USER CHECKED")
		uh.loggedInUser = &userInQuestion
		claims := rtoken.Claims(userInQuestion.Username, uh.userSess.Expires)
		session.Create(claims, uh.userSess.UUID, uh.userSess.SigningKey, w)
		sessionJson, err := json.Marshal(uh.userSess)
		//fmt.Println("session json", sessionJson)
		if err!=nil{
			fmt.Println("user session error")
			loginForm.VErrors.Add("generic", "Username or password")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
			fmt.Println("Marshal err")
			return
		}
		sessionResp, err := http.Post(sessionDest,"application/json", bytes.NewBuffer(sessionJson))
		if err!=nil{
			loginForm.VErrors.Add("generic", "Username or password")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
			fmt.Println("Session resp err")
			return
		}

		sessionRespSize:= sessionResp.ContentLength
		sessionBody :=  make ([]byte, sessionRespSize)
		sessionResp.Body.Read(sessionBody)
		//fmt.Println("session body", sessionBody)

		sessErrJson := json.Unmarshal(sessionBody, &userSession)

		if sessErrJson!=nil{
			fmt.Println(sessErrJson.Error())
			loginForm.VErrors.Add("generic", "Username or password")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
			fmt.Println("Session json err")
			return
		}

		uh.userSess = &userSession
		roles, _ := uh.userService.UserRoles(&userInQuestion)
		if uh.checkAdmin(roles) {
			fmt.Println("ADMIN")
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/questions", http.StatusSeeOther)

	}
}

// Logout hanldes the POST /logout requests
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	deleteSessionDest := "http://localhost:8181/sessions/"
	userSess, err2:= r.Context().Value(ctxUserSessionKey).(*entities.Session)
	fmt.Println("check user session", userSess)
	if err2{
		fmt.Println("error fetching user session")
	}

	fmt.Println("about to call session.REmove")
	session.Remove(userSess.UUID, w)
	uuidJson, err := json.Marshal(userSess.UUID)
	fmt.Println("session json", uuidJson)
	if err!=nil{
		fmt.Println("Marshal err in uuidJson")
		return
	}
	req, err := http.NewRequest("DELETE", deleteSessionDest+userSess.UUID, bytes.NewBuffer(uuidJson))
	if err!=nil{
		fmt.Println("DEL request error err in uuidJson")
		return
	}

	client := http.Client{}

	resp, err := client.Do(req)

	//resp, err := http.DEL(deleteSessionDest + userSess.UUID, "application/json", bytes.NewBuffer(uuidJson))
	fmt.Println("deleted via ui", resp)
	//uh.sessionService.DeleteSession(userSess.UUID)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Signup hanldes the GET/POST /signup requests
func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	dest := "http://localhost:8181/users/"
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "user-entry.html", signUpForm)
		return
	} else if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		singnUpForm.Required("firstName", "lastName", "email", "password", "cPassword")
		singnUpForm.MatchesPattern("email", form.EmailRX)
		//singnUpForm.MatchesPattern("phone", form.PhoneRX)
		singnUpForm.MinLength("password", 8)
		singnUpForm.PasswordMatches("password", "cPassword")
		//singnUpForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !singnUpForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", singnUpForm)
			return
		}

		//pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		//if pExists {
		//	singnUpForm.VErrors.Add("phone", "Phone Already Exists")
		//	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		//	return
		//}
		eExists := uh.userService.EmailExists(r.FormValue("email"))
		if eExists {
			singnUpForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", singnUpForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			singnUpForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "user-entry.html", singnUpForm)
			return
		}

		// role, errs := uh.userRole.RoleByName("USER")

		// if len(errs) > 0 {
		// 	fmt.Println("error: role", errs)
		// 	singnUpForm.VErrors.Add("role", "could not assign role to the user")
		// 	uh.tmpl.ExecuteTemplate(w, "user-entry.html", singnUpForm)
		// 	return
		// }

		user := entities.User {
			Username   :r.FormValue("username"),
			FirstName  :r.FormValue("firstName"),
			LastName   :r.FormValue("lastName"),
			Email      :r.FormValue("email"),
			Password   :string(hashedPassword),
			RoleID     :"UserRole1",
			Interests  :[]byte("[]"),
		}

		newUser, err := json.MarshalIndent(user, "", "\t")

		// fmt.Println(newUser)

		if err != nil {
			fmt.Println("this happened")
			http.Redirect(w, r, "/userentry", http.StatusNotAcceptable)
		}

		// req, err := http.NewRequest("POST", dest, bytes.NewBuffer(newUser))
		// if err != nil {
		// 	fmt.Println("this happened to")
		// 	http.Redirect(w, r, "/userentry", http.StatusNotAcceptable)
		// }
		// req.Header.Set("Content-Type", "application/json")

		// client := &http.Client{}
		// resp, err := client.Do(req)
		resp, err := http.Post(dest, "application/json", bytes.NewBuffer(newUser))
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Println("response Statusp:", resp.Status)
			http.Redirect(w, r, "/questions", http.StatusSeeOther)

		} else {
			fmt.Println("response Status:", resp.Status)
			http.Redirect(w, r, "/userentry", http.StatusSeeOther)
		}

		// fmt.Println("response Headers:", resp.Header)
		// body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println("response Body:", string(body))
		//localhost:8181/users POST userjson
		//if len(errs) > 0 {
		//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		//	return
		//}
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		fmt.Println("this is no user user session")
		return false
	}
	userSess := uh.userSess
	fmt.Println("there is a user session", userSess)
	if r.Context().Value(ctxUserSessionKey) != nil {
		fmt.Println(r.Context().Value(ctxUserSessionKey))
	}
		c, err := r.Cookie(userSess.UUID)

		// cookie is related with session id
	fmt.Println("got the cookie ", c)
	if err != nil {
		fmt.Println("didn't get the cookie ", err)
		return false
	}
	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		fmt.Println("there is error while checking session validity")
		return false
	}
	fmt.Println("user is logged in")
	return true
}
func (uh *UserHandler) checkAdmin(rs []entities.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("Admin") {
			return true
		}
	}
	return false
}
