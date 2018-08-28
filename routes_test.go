package main

import (
  // "fmt"
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
  "strings"
  "bytes"
)

func setup() *mux.Router {
  runSQL("TRUNCATE TABLE meal_foods RESTART IDENTITY")
  runSQL("TRUNCATE TABLE foods RESTART IDENTITY")
  runSQL("TRUNCATE TABLE meals RESTART IDENTITY")
  seedMeals()

  r := mux.NewRouter()
  setRoutes(r)
  return r
}

func TestGetPort(t *testing.T) {
  port := getPort()
  if port != ":3000" {
    t.Error("Expected :3000, got ", port)
  }
}

func TestGetFoods(t *testing.T) {
  r := setup()
  runSQL("INSERT INTO foods (name, calories) VALUES ('Pizza', 500)")
  runSQL("INSERT INTO foods (name, calories) VALUES ('Cat', 700)")
  req, _ := http.NewRequest("GET", "/api/v1/foods/", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Pizza","calories":500},{"id":2,"name":"Cat","calories":700}]`
  if(actual != expected) {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestAddFood(t *testing.T) {
  r := setup()
  payload := []byte(`{"food":{"name":"sushi","calories":"300"}}`)
  req, _ := http.NewRequest("POST", "/api/v1/foods/", bytes.NewBuffer(payload))
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `{"id":1,"name":"sushi","calories":300}`
  if(actual != expected) {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestGetOneFood(t *testing.T) {
  r := setup()
  runSQL("INSERT INTO foods (name, calories) VALUES ('Pizza', 500)")
  runSQL("INSERT INTO foods (name, calories) VALUES ('Cat', 700)")
  req, _ := http.NewRequest("GET", "/api/v1/foods/2/", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `{"id":2,"name":"Cat","calories":700}`
  if(actual != expected) {
    t.Error("Get One Food - Expected:", expected, "Got:", actual)
  }
}

func TestDeleteFood(t *testing.T) {
  r := setup()
  runSQL("INSERT INTO foods (name, calories) VALUES ('Pizza', 500)")
  runSQL("INSERT INTO foods (name, calories) VALUES ('Cat', 700)")
  req, _ := http.NewRequest("DELETE", "/api/v1/foods/2", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  status := response.Code
  if(status != 204) {
    t.Error("Delete Food - Expected 204, Got:", status)
  }
}

func TestGetMeals(t *testing.T) {
  r := setup()
  runSQL("INSERT INTO foods (name, calories) VALUES ('Pizza', 500)")
  runSQL("INSERT INTO foods (name, calories) VALUES ('Cat', 700)")
  runSQL("INSERT INTO meal_foods (food_id, meal_id) VALUES (1, 1)")
  runSQL("INSERT INTO meal_foods (food_id, meal_id) VALUES (2, 1)")
  req, _ := http.NewRequest("GET", "/api/v1/meals/", nil)
  response := httptest.NewRecorder()
  r.ServeHTTP(response, req)
  actual := response.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Breakfast","foods":[{"id":1,"name":"Pizza","calories":500},{"id":2,"name":"Cat","calories":700}]},{"id":2,"name":"Snack","foods":null},{"id":3,"name":"Lunch","foods":null},{"id":4,"name":"Dinner","foods":null}]`
  if(actual != expected) {
    t.Error("Get One Food - Expected:\n", expected, "Got:\n", actual)
  }
}
