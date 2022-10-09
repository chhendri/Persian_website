package main


import (
  "net/http"
  "github.com/go-sql-driver/mysql"
  "database/sql"
  "fmt"
  "log"
  "os"
  //"html/template"
  //"encoding/json"
  //"bytes"
)

var db *sql.DB

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func connectDB(w http.ResponseWriter, r *http.Request) {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   getEnv("MYSQL_USER", "Persian_dvl"),
        Passwd: getEnv("MYSQL_ROOT_PASSWORD", "password"),
        Net:    "tcp",
        Addr:   getEnv("MYSQL_HOST", "127.0.0.1:3306"),
        DBName: getEnv("MYSQL_DATABASE", "Persian"),
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
