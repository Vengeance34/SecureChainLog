package db

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    connStr := "host=localhost port=5432 user=postgres password=123450 dbname=securechainlog sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    fmt.Println("✅ Connected to PostgreSQL")
    return db, nil
}
