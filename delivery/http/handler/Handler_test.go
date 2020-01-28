package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EmailConfig struct {
	Username   string
	Password   string
	ServerHost string
	ServerPort string
	SenderAddr string
}

type EmailSender interface {
	Send(to []string, body []byte) error
}

func NewEmailSender(conf EmailConfig) EmailSender {
	return &emailSender{conf, smtp.SendMail}
}

type emailSender struct {
	conf EmailConfig
	send func(string, smtp.Auth, string, []string, []byte) error
}

func (e *emailSender) Send(to []string, body []byte) error {
	addr := e.conf.ServerHost + ":" + e.conf.ServerPort
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, to, body)
}

/****** testing ******/

func TestEmail_SendSuccessful(t *testing.T) {
	f, r := mockSend(nil)
	sender := &emailSender{send: f}

	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/articles/:id", nil)
	if err != nil {
		t.Fatal(err)
	}

	ArticleHandler.PostArticle(httprr, req)
	respBody := httprr.Body()

	err := sender.Send([]string{"me@example.com"}, respBody)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(r.msg) != body {
		t.Errorf("wrong message body.")
	}
}

func mockSend(errToReturn error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		*r = emailRecorder{addr, a, from, to, msg}
		return errToReturn
	}, r
}

type emailRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}


func TestPostArticle(t *testing.T) {

	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/articles/:id", nil)
	if err != nil {
		t.Fatal(err)
	}

	ArticleHandler.PostArticle(httprr, req)
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestGetArticle(t *testing.T) {

	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	ArticleHandler.GetArticle(httprr, req)
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

}