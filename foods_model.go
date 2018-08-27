package main

import (
  "fmt"
  "log"
)

// Food Struct (model)
type Food struct {
  ID        int    `json:"id"`
  Name      string  `json:"name"`
  Calories  int     `json:"calories"`
}

func getFoodsFromDB() []Food {
  queryString := "SELECT * FROM foods"
  rows, err := db().Query(queryString)
  if err != nil {
    log.Fatal(err)
  }
  var (
    id int
    name string
    calories int
    foods []Food
  )
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&id, &name, &calories)
    if err != nil {
      log.Fatal(err)
    }
    food := Food{ID: id, Name: name, Calories: calories}
    foods = append(foods, food)
  }
  return foods
}

func getFoodFromDB(id int) Food {
  queryString := "SELECT * FROM foods WHERE id=$1"
  fmt.Println("Preparing to get food:", queryString, id)
  var food Food
  err := db().QueryRow(queryString, id).Scan(&food.ID, &food.Name, &food.Calories)
  if err != nil {
    log.Fatal(err)
  }
  return food
}

func addFoodToDB(food Food) int {
  // db := db()
  queryString := "INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING id"
  fmt.Println("Preparing to add food:", queryString, food.Name, food.Calories)
  id := 0
  err := db().QueryRow(queryString, food.Name, food.Calories).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }
  return id
}
