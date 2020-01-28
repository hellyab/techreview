package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hellyab/techreview/session"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hellyab/techreview/entities"
)

var baseURL = "http://localhost:8181/questions"

//SingleData represents a single user
type SingleData struct {
	Question entities.Question
}

//CollectionData represents a collection of users
type CollectionData struct {
	Questions []entities.Question
}

func (uh *UserHandler) FollowQuestion( w http.ResponseWriter, r *http.Request){
	fmt.Println("are you even here?" )
	dest := "http://localhost:8181/questions/follow"
	userID := uh.loggedInUser.ID
	questionID :=r.FormValue("question-id")
	questionFollow := entities.QuestionFollow{QuestionID:questionID, UserID: userID}
	qFollowJSON, err :=json.MarshalIndent(questionFollow, "", "\t")
	if err!=nil{
		fmt.Println("marshalling error", err.Error())
		return
	}
	resp, err:=http.Post(dest, "application/json", bytes.NewBuffer(qFollowJSON))
	fmt.Println(resp.Status)
	if err!=nil{
		fmt.Println("follow-error", err.Error())
		return
	}
	if resp.StatusCode==200{
		http.Redirect(w,r,"/questions", http.StatusSeeOther)
	}




}

//FetchQuestions fetchs all the questions in the database
func (uh *UserHandler) FetchQuestions( w http.ResponseWriter, _ *http.Request) {
	qdata := []entities.Question{}
	res, err := http.Get(baseURL)

	if err != nil {
		return
	}

	// qdata := &CollectionData{}
	// fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return
	}

	// fmt.Println(body)
	err = json.Unmarshal(body, &qdata)

	if err != nil {
		fmt.Println("Unmarshling error")
		panic(err)
		// return nil, err
	}
	fmt.Println(qdata[0].InquirerID)

	templateData := struct {
		LoggedIn bool
		Questions []entities.Question
		UserID string

	}{
		true,
		qdata,
		uh.loggedInUser.ID,
	}

	uh.tmpl.ExecuteTemplate(w, "questions.html", templateData)

}


//StoreQuestion sends a question to the database through the rest api
func StoreQuestion(w http.ResponseWriter, r *http.Request){
	dest:= "http://localhost:8181/questions"
	inquiry := r.FormValue("inquiryBody")
	userSess, errBool:= r.Context().Value(ctxUserSessionKey).(*entities.Session)
	if !errBool{
		fmt.Println("COntext error")
	}
	//userSess := UserHandler{}.userSess
	c, err := r.Cookie(userSess.UUID)

	if err!=nil{
		fmt.Println(err.Error())
	}

	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		fmt.Println("Expired session")
		return
	}

	hmacSecret := []byte(userSess.SigningKey)
	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})
	if err!=nil{
		fmt.Println(err.Error())
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"]
		userIDString := fmt.Sprintf("%v", userID)
	//fmt.Println("the found user ID is ", userID)
		questionStruct := struct {
			InquirerID      string
			Inquiry         string
			AskedAt	time.Time
			Follows         uint
			NumberOfAnswers uint
			Topics          json.RawMessage
		}{
			userIDString,
			inquiry,
			time.Now(),
			0,
			0,
			json.RawMessage("[]"),
		}

		questionJSON, err := json.MarshalIndent(questionStruct, "", "\t")
		//questionJSONString := fmt.Sprintf(`"InquirerID": "%s", "Inquiry": "%s", "Follows":%d, "NumberOfAnswers":%d, "Topics":[]`,userID, inquiry, 0, 0)
		//question := []byte(questionJSONString)
		resp, err := http.Post(dest,"application/json", bytes.NewBuffer(questionJSON))

		if err!=nil{
			fmt.Println("question post error - ", err.Error())
		}
		fmt.Println(resp.StatusCode)
		if resp.StatusCode ==200{
			http.Redirect(w,r,"/questions", http.StatusSeeOther)
		}

	} else {
		fmt.Println("Invalid JWT Token")
	}
		//fmt.Println("value---", token["userID"])
	//fmt.Println("cookie", c)
	//fmt.Println("asked by", inquirer)
	fmt.Println("question", inquiry)

}

func (uh *UserHandler) AnswerAQuestion (w http.ResponseWriter, r *http.Request){
	dest := "http://localhost:8181/answers"
	userID := uh.loggedInUser.ID
	questionID := r.FormValue("questionID")
	answer := r.FormValue("answer")

	newAnswer := entities.Answer{QuestionID:questionID, ReplierID:userID,Answer:answer, Votes:0}

	ansJSON, err := json.MarshalIndent(newAnswer,"", "\t", )

	if err!=nil{
		fmt.Println("Error Marshaling Answer")
	}


	_, err2 := http.Post(dest, "application/json", bytes.NewBuffer(ansJSON))

	if err2!=nil{
		fmt.Println("error posting answer")
	}


		http.Redirect(w, r, "http://localhost:8080/questions/read?questionID="+questionID, http.StatusSeeOther)
		return



}

func (uh *UserHandler) SingleQuestion (w http.ResponseWriter, r *http.Request){
	dest := "http://localhost:8181/questions/"
	questionID := r.FormValue("questionID")
	fmt.Println("The question we are looking for is ", questionID)
	question := entities.Question{}

	questResp, err := http.Get(dest + questionID)

	//fmt.Println(userResp)
	if err!=nil {
		fmt.Println("Error here: not found")
		//loginForm.VErrors.Add("generic", "Your email address or password is wrong")
		//uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
		return
	}
	size:= questResp.ContentLength
	body :=  make ([]byte, size)
	//fmt.Println(body)
	questResp.Body.Read(body)
	//fmt.Println(body)

	errJson := json.Unmarshal(body, &question)
	fmt.Println(body, "ufff")
	fmt.Println(question)

	if errJson!=nil{
		fmt.Println(errJson.Error())
		fmt.Println("Some error has occured")
	}

	ansDest := "http://localhost:8181/answers/byquestion/"+ questionID
	ansResp, err := http.Get(ansDest)
	answers := []entities.AnswersByQuesId{}

	//fmt.Println(userResp)
	if err!=nil {
		fmt.Println("Error here: not found")
		//loginForm.VErrors.Add("generic", "Your email address or password is wrong")
		//uh.tmpl.ExecuteTemplate(w, "user-entry.html", loginForm)
		return
	}
	size2 := ansResp.ContentLength
	body2 :=  make ([]byte, size2)
	//fmt.Println(body)
	ansResp.Body.Read(body2)
	//fmt.Println(body)

	errJson = json.Unmarshal(body2, &answers)

	if errJson!=nil{
		fmt.Println(errJson.Error())
		fmt.Println("Some error has occured")
	}


	fmt.Println("QUestion\n", question, "\nAnswer\n", answers)


	PayLoad := struct{
		LoggedIn bool
		Question entities.Question
		Answers []entities.AnswersByQuesId
	}{
		LoggedIn:true,
		Question:question,
		Answers:answers,

	}
		 uh.tmpl.ExecuteTemplate(w, "question.html", PayLoad)
}

//DeleteQuestion deletes a question from the database
func DeleteQuestion(id string) (*entities.Question, error) {
	qstn := entities.Question{}
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s", baseURL, id)
	req, _ := http.NewRequest("DELETE", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// qstn := &CollectionData{}
	// fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return nil, err
	}

	// fmt.Println(body)
	err = json.Unmarshal(body, &qstn)

	if err != nil {
		fmt.Println("Unmarshling error")
		panic(err)
		// return nil, err
	}

	return &qstn, nil
}
