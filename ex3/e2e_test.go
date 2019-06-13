package main

import(
	"testing"
	"net/http"
	//"io/ioutil"
	"net/http/httptest"
)

func TestHelloHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil{
		panic(err)
	}

	response := httptest.NewRecorder()

	http.HandlerFunc(handler).
		ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, response.Code)
	}

	expected := "Hello World"
	if response.Body.String() != expected {
		t.Errorf("body differs expected:%v,actual:%v",expected,response.Body.String())
	}
}