package main

import (
  "fmt"
  "log"
  "strings"
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

func getMealFromDB(id int) Meal {
  queryString := "SELECT * FROM meals WHERE id=$1"
  fmt.Println("Preparing to get meal:", queryString, id)
  var meal Meal
  err := db().QueryRow(queryString, id).Scan(&meal.ID, &meal.Name)
  if err != nil {
    log.Fatal(err)
  }
  return meal
}

type Message struct {
  Message string `json:"message"`
}

func postMealFoodToDB(foodID int, mealID int) Message {
  food := getFoodFromDB(foodID)
  meal := getMealFromDB(mealID)
  queryString := "INSERT INTO meal_foods (meal_id, food_id) VALUES ($1, $2) RETURNING id"
  fmt.Println("Preparing to add meal_food:", queryString, mealID, foodID)
  id := 0
  err := db().QueryRow(queryString, mealID, foodID).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }
  if id > 0 {
    message := fmt.Sprintf("Successfully added %v to %v", strings.ToUpper(food.Name), strings.ToUpper(meal.Name))
    return Message{Message: message}
  } else {
    return Message{Message: "Failure"}
  }
}
