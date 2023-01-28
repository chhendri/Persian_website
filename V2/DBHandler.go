package main

import (
  "database/sql"
  "github.com/go-sql-driver/mysql"
  "os"
  "log"
  "fmt"
)


var db *sql.DB

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func connectDB() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   getEnv("MYSQL_USER", "Persian_dvl"),
        Passwd: getEnv("MYSQL_PASSWORD", "password"),
        Net:    "tcp",
        Addr:   getEnv("MYSQL_HOST", "127.0.0.1:3306"),
        DBName: getEnv("MYSQL_DATABASE", "Persian_V2"),
        AllowNativePasswords: true,
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
