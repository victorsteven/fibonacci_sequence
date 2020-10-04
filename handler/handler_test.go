package handler_test

import (
	"encoding/json"
	"fibonacci_sequence/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//To achieve unit test on the handler methods, we need to fake the domain layer:
type fakeFibService struct {
	FibonacciGeneratorFn func(request string) int
}

func (u *fakeFibService) FibonacciGenerator(request string) int {
	return u.FibonacciGeneratorFn(request)
}

var (
	fakeFib fakeFibService
	h = handler.NewHandlerService(&fakeFib)
)

type Response struct {
	Result int `json:"result"`
}

type NegativeResponse struct {
	Result string `json:"result"`
}

func TestService_CurrentNumber(t *testing.T) {

	//We assume that the current number returns zero
	fakeFib.FibonacciGeneratorFn = func(request string) int {
		return 0
	}

	req, err := http.NewRequest(http.MethodGet, "/current", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/current", h.CurrentNumber)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = Response{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Result, 0)
}

func TestService_NextNumber(t *testing.T) {

	//We assume that the next number returns 1
	fakeFib.FibonacciGeneratorFn = func(request string) int {
		return 1
	}

	req, err := http.NewRequest(http.MethodGet, "/next", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/next", h.NextNumber)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = Response{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Result, 1)
}

func TestService_PreviousNumber_Negative(t *testing.T) {

	//We assume that the previous number returns -1, which we wont permit permit in the sequence, so we return a message
	fakeFib.FibonacciGeneratorFn = func(request string) int {
		return -1
	}

	req, err := http.NewRequest(http.MethodGet, "/previous", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/previous", h.PreviousNumber)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = NegativeResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Result, "This cannot go below 0")
}


func TestService_PreviousNumber_Positive(t *testing.T) {

	//We assume that the previous number returns 1
	fakeFib.FibonacciGeneratorFn = func(request string) int {
		return 1
	}

	req, err := http.NewRequest(http.MethodGet, "/previous", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	r := gin.Default()
	r.GET("/previous", h.PreviousNumber)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response = Response{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.EqualValues(t, rr.Code, 200)
	assert.EqualValues(t, response.Result, 1)
}