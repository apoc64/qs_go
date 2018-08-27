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

func TestGetFoods(t *testing.T) { // change for db
  r := setup()
  // runSQL("INSERT INTO foods (name, calories) VALUES ('pizza', 500)")
  // runSQL("INSERT INTO foods (name, calories) VALUES ('cat', 700)")
  addFoodToDB(Food{ID: 1, Name: "Pizza", Calories: 500})
  addFoodToDB(Food{ID: 1, Name: "Cat", Calories: 700})
  req, _ := http.NewRequest("GET", "/api/v1/foods/", nil)
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)
  actual := rr.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `[{"id":1,"name":"Pizza","calories":500},{"id":2,"name":"Cat","calories":700}]`
  if(actual != expected) {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}

func TestAddFood(t *testing.T) {
  r := setup()
  payload := []byte(`{"name":"sushi","calories":300}`)
  req, _ := http.NewRequest("POST", "/api/v1/foods/", bytes.NewBuffer(payload))
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)
  actual := rr.Body.String()
  actual = strings.TrimRight(actual, "\r\n ")
  expected := `{"id":1,"name":"sushi","calories":300}`
  if(actual != expected) {
    t.Error("Get Foods - Expected:", expected, "Got:", actual)
  }
}
