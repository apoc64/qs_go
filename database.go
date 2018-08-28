package main

import (
  "fmt"
  "log"
  "os"
  "database/sql"
  "github.com/lib/pq"
)

var database *sql.DB
var dbInitialized bool

func db() *sql.DB {
  if !dbInitialized {
    initializeDB()
  }
  return database
}

func initializeDB() {
  fmt.Println("Initializing database with name:", getDBName())
  newDB, err := sql.Open("postgres", getDBName())
  if err != nil {
    log.Fatal(err)
  }
  dbInitialized = true
  database = newDB
  fmt.Println("Database Initialized:", newDB)
  migrateDB()
  seedDB()
}

func getDBName() string {
  url := os.Getenv("DATABASE_URL")
  if url != "" {
    connection, _ := pq.ParseURL(url)
    connection += " sslmode=require"
    return connection
  }
  return "dbname=qs_go sslmode=disable"
}

func migrateDB() {
  runSQL(foodsTableCreation)
  runSQL(mealsTableCreation)
  runSQL(mealFoodsTableCreation)
}

func seedDB() {
  count := 0
  if err := db().QueryRow("SELECT COUNT(id) FROM meals").Scan(&count); err != nil {
    log.Fatal(err)
  }
  fmt.Println("Number of meals:", count)
  if count != 4 {
    seedMeals()
  }
}

func seedMeals() {
  fmt.Println("Seeding meals")
  runSQL("INSERT INTO meals (id, name) VALUES (1, 'Breakfast')")
  runSQL("INSERT INTO meals (id, name) VALUES (2, 'Snack')")
  runSQL("INSERT INTO meals (id, name) VALUES (3, 'Lunch')")
  runSQL("INSERT INTO meals (id, name) VALUES (4, 'Dinner')")
}

func runSQL(sqlQuery string) {
  fmt.Println("Preparing to execute:", sqlQuery)
  if _, err := db().Exec(sqlQuery); err != nil {
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
