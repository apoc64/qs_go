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
  // fmt.Println(meals[0])
  for i := 0; i <= 3; i++ {
    meals[i].Foods = getMealFoods(i + 1)
  }
  return meals
}

func getMealFoods(meal_id int) []Food {
  queryString := "SELECT foods.id, foods.name, foods.calories FROM foods INNER JOIN meal_foods ON foods.id=meal_foods.food_id WHERE meal_foods.meal_id=$1"
  rows, err := db().Query(queryString, meal_id)
  if err != nil {
    log.Fatal(err)
  }
  var (
    food Food
    foods []Food
  )
  defer rows.Close() // async - closes rows when func finishes
  for rows.Next() {
    if err := rows.Scan(&food.ID, &food.Name, &food.Calories); err != nil {
      log.Fatal(err)
    }
    foods = append(foods, food)
  }
  return foods
}
