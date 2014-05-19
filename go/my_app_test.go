package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostPerson(t *testing.T) {
	database := Open()
	defer database.Close()

	p := Person{
		id:      123456,
		name:    "Billy Smith",
		email:   "r@example.com",
		updates: 0}

	database.Insert(p)
	database.Insert(p) // do twice to verify update counts

	req, err := http.NewRequest("GET", "http://localhost:4567/list_people", nil)
	if err != nil {
		panic("Can not get /list_peoplee")
	}

	w := httptest.NewRecorder()
	httpGetHandler(w, req)

	expected_title := "ID    Name            Email                Updates"
	expected_row := "Billy Smith     r@example.com        2"

	actual_body := w.Body.String()
	if !strings.Contains(actual_body, expected_title) {
		t.Errorf("expected title %s, actual body = %s", expected_title, actual_body)
	}

	if !strings.Contains(actual_body, expected_row) {
		t.Errorf("expected row %s, actual body = %s", expected_title, actual_body)
	}
}
