package main

import (
  // "fmt"
  // "log"
  // "os"
  // "database/sql"
  // _ "github.com/lib/pq"
)

// var database *sql.DB
// var dbInitialized bool
//
// func db() *sql.DB {
//   if !dbInitialized {
//     fmt.Println("Initializing database")
//     newDB, err := sql.Open("postgres", getDBName())
//     if err != nil {
//       log.Fatal(err)
//     }
//     migrateDB(newDB)
//     dbInitialized = true
//     database = newDB
//   }
//   return database
// }
//
// func getDBName() string {
//   d := os.Getenv("DATABASE_URL")
//   if d != "" {
//     return d + "?ssl=true"
//   }
//   return "postgres://localhost/qs_go"
// }
//
// func migrateDB(db *sql.DB) {
//
// }
//
// Food Struct (model)
type Food struct {
  ID        int    `json:"id"`
  Name      string  `json:"name"`
  Calories  int     `json:"calories"`
}

func getFoodsFromDB() {
  db()
}

func addFoodToDB(name string, calories int) {


}
