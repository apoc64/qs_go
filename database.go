package main

import (
  "fmt"
  "log"
  "os"
  "database/sql"
  "github.com/lib/pq"
  // "strings"
)

var database *sql.DB
var dbInitialized bool

func db() *sql.DB {
  if !dbInitialized {
    fmt.Println("Initializing database")
    fmt.Println(getDBName())
    newDB, err := sql.Open("postgres", getDBName())
    if err != nil {
      log.Fatal(err)
    }
    migrateDB(newDB)
    fmt.Println(getDBName())
    fmt.Println(newDB)
    dbInitialized = true
    database = newDB
  }
  return database
}

func getDBName() string {
  url := os.Getenv("DATABASE_URL")
  if url != "" {
    connection, _ := pq.ParseURL(url)
    connection += " sslmode=require"

    // url = strings.TrimPrefix(url, "postgres://")
    // url = fmt.Sprintf("dbname=%s", url)
    return connection
  }
  return "dbname=qs_go sslmode=disable"
}

func migrateDB(db *sql.DB) {
  if _, err := db.Exec(foodsTableCreation); err != nil {
    log.Fatal(err)
  }
  if _, err := db.Exec(mealsTableCreation); err != nil {
    log.Fatal(err)
  }
  if _, err := db.Exec(mealFoodsTableCreation); err != nil {
    log.Fatal(err)
  }
}

const foodsTableCreation = `CREATE TABLE IF NOT EXISTS foods
(
  id SERIAL,
  name TEXT,
  calories INT,
  CONSTRAINT foods_pkey PRIMARY KEY (id)
);`

const mealsTableCreation = `CREATE TABLE IF NOT EXISTS meals
(
  id SERIAL,
  name TEXT,
  CONSTRAINT meals_pkey PRIMARY KEY (id)
);`

const mealFoodsTableCreation = `CREATE TABLE IF NOT EXISTS meal_foods
(
  id SERIAL,
  food_id INT,
  meal_id INT,
  CONSTRAINT meal_foods_pkey PRIMARY KEY (id)
);`
