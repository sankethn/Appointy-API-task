package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserByID(t *testing.T) {
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

func TestCreateUser(t *testing.T) {

	var jsonStr = []byte(`{
		"uid": "5",
		"name": "vishwa",
		"email": "vishwa@abc.com",
		"password": "vishwa"
	}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		"uid": "5",
		"name": "vishwa",
		"email": "vishwa@abc.com",
		"password": "vishwa"
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetPostByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code, got %v expected %v", status, http.StatusOK)
	}

	expected := `{
		"_id": "61614440dbc7ad2ba637938a",
		"pid": "1",
		"uid": "1",
		"caption": "tired",
		"imageurl": "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/2022-chevrolet-corvette-z06-1607016574.jpg?crop=0.737xw:0.738xh;0.181xw,0.218xh&resize=640:*",
		"postedtimestamp": {
			"T": 1633764416,
			"I": 1
		}
	}`
	if rr.Body.String() == expected {
		t.Errorf("handler returned unexpected body, got %v expected %v", rr.Body.String(), expected)
	}
}

func TestListPostsByUID(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ListUserPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code, got %v expected %v", status, http.StatusOK)
	}

	expected := `[
		{
			"_id": "61614e12dea47b29f5ce8b82",
			"pid": "4",
			"uid": "3",
			"caption": "have a nice day!",
			"imageurl": "https://cdn.motor1.com/images/mgl/8e8Mo/s1/most-expensive-new-cars-ever.webp",
			"postedtimestamp": {
				"T": 1633766930,
				"I": 1
			}
		}
	]`
	if rr.Body.String() == expected {
		t.Errorf("handler returned unexpected body, got %v expected %v", rr.Body.String(), expected)
	}
}

func TestCreatePost(t *testing.T) {

	var jsonStr = []byte(`{
		"_id": "61614440dbc7ad2ba637938a",
		"pid": "1",
		"uid": "1",
		"caption": "tired",
		"imageurl": "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/2022-chevrolet-corvette-z06-1607016574.jpg?crop=0.737xw:0.738xh;0.181xw,0.218xh&resize=640:*",
		"postedtimestamp": {
			"T": 1633764416,
			"I": 1
		}
	}`)

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		"_id": "61614440dbc7ad2ba637938a",
		"pid": "1",
		"uid": "1",
		"caption": "tired",
		"imageurl": "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/2022-chevrolet-corvette-z06-1607016574.jpg?crop=0.737xw:0.738xh;0.181xw,0.218xh&resize=640:*",
		"postedtimestamp": {
			"T": 1633764416,
			"I": 1
		}
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
