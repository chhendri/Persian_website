package main


import (
  "net/http"
  "github.com/go-sql-driver/mysql"
  "database/sql"
  "fmt"
  "log"
  //"html/template"
  //"encoding/json"
  //"bytes"
)

var db *sql.DB


func connectDB(w http.ResponseWriter, r *http.Request) {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   "Persian_dvl",
        Passwd: "dvl",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "Persian",
    }
    // Get a database handle.
    var err error

    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    // Ping the connection
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected to the database Persian!")
}
