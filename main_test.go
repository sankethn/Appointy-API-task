package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnById(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code, got %v expected %v", status, http.StatusOK)
	}

	expected := `{
		"_id": "616148873a2829c036e71fe8",
		"uid": "1",
		"name": "sanket",
		"email": "sanket@abc.com",
		"password": "$2a$08$lWsE6dOSOg0Xn1w21L7FFOPTmvZ/aplYLHuye4iABqwPVFlhjPmRy"
	}`
	if rr.Body.String() == expected {
		t.Errorf("handler returned unexpected body, got %v expected %v", rr.Body.String(), expected)
	}
}
