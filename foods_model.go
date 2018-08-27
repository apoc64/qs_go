package main

import (
  "fmt"
  "log"
)

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
  queryString := "INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING id"
  fmt.Println("Preparing to add food:", queryString, food.Name, food.Calories)
  id := 0
  err := db().QueryRow(queryString, food.Name, food.Calories).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }
  return id
}

func deleteFoodFromDB(id int) bool {
  queryString := "DELETE FROM foods WHERE id=$1 RETURNING id, name"
  fmt.Println("Preparing to delete food:", queryString, id)
  deleted_id := 0
  name := ""
  err := db().QueryRow(queryString, id).Scan(&deleted_id, &name)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Deleted", name)
  return (id == deleted_id)
}
