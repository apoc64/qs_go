package models

import (
  // "qs_go"
  "fmt"
  "log"
  "strings"
)

type Meal struct {
  ID     int        `json:"id"`
  Name   string     `json:"name"`
  Foods  []Food     `json:"foods"`
}

func GetMealsFromDB() []Meal {
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
  for i := 0; i <= 3; i++ {
    meals[i].Foods = GetMealFoods(i + 1)
  }
  return meals
}

func GetMealFoods(meal_id int) []Food {
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

func GetMealFromDB(id int) Meal {
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

func PostMealFoodToDB(foodID int, mealID int) Message {
  food := GetFoodFromDB(foodID)
  meal := GetMealFromDB(mealID)
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

func DeleteMealFoodFromDB(foodID int, mealID int) Message {
  food := GetFoodFromDB(foodID)
  meal := GetMealFromDB(mealID)
  queryString := "DELETE FROM meal_foods WHERE meal_id=$1 AND food_id=$2 RETURNING id"
  fmt.Println("Preparing to add meal_food:", queryString, mealID, foodID)
  id := 0
  err := db().QueryRow(queryString, mealID, foodID).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }
  if id > 0 {
    message := fmt.Sprintf("Successfully removed %v from %v", strings.ToUpper(food.Name), strings.ToUpper(meal.Name))
    return Message{Message: message}
  } else {
    return Message{Message: "Failure"}
  }
}
