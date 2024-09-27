package database

import (
    "database/sql"
    "os" // Import the os package
    _ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
    var err error

    // Ensure the directory exists
    err = os.MkdirAll("/tmp/data/database", os.ModePerm) // Use = instead of :=
    if err != nil {
        panic(err)
    }

    // Open the database
    DB, err = sql.Open("sqlite", "/tmp/data/database/products.db")
    if err != nil {
        panic(err)
    }

    // Create the products table if it doesn't exist
    createTable := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        description TEXT,
        price REAL
    );`
    
    _, err = DB.Exec(createTable)
    if err != nil {
        panic(err)
    }
}
