package main

import (
  // "fmt"
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
  "strings"
)

func TestGetPort(t *testing.T) {
  port := getPort()
  if port != ":3000" {
    t.Error("Expected :3000, got ", port)
  }
}

func TestGetFoods(t *testing.T) { // change for db
  r := mux.NewRouter()
  setRoutes(r)
  req, _ := http.NewRequest("GET", "/api/v1/foods/", nil)
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)
  actual := rr.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Pizza","calories":400},{"id":2,"name":"Cat","calories":800}]`
  if(actual != expected) {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}
