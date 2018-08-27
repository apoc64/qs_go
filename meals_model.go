package main

import (
  // "fmt"
  "log"
)

type Meal struct {
  ID     int        `json:"id"`
  Name   string     `json:"name"`
  Foods  []Food     `json:"foods"`
}

func getMealsFromDB() []Meal {
  queryString := "SELECT * FROM meals"
  rows, err := db().Query(queryString)
  if err != nil {
    log.Fatal(err)
  }
  var (
    meal Meal
    meals []Meal
  )
  defer rows.Close() // async - closes rows when func finishes
  for rows.Next() {
    if err := rows.Scan(&meal.ID, &meal.Name); err != nil {
      log.Fatal(err)
    }
    meals = append(meals, meal)
  }
  return meals
}
