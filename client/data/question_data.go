package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

//FetchQuestions fetchs all the questions in the database
func FetchQuestions() ([]entities.Question, error) {
	qdata := []entities.Question{}
	res, err := http.Get(baseURL)

	if err != nil {
		return nil, err
	}

	// qdata := &CollectionData{}
	// fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return nil, err
	}

	// fmt.Println(body)
	err = json.Unmarshal(body, &qdata)

	if err != nil {
		fmt.Println("Unmarshling error")
		panic(err)
		// return nil, err
	}

	return qdata, nil
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
