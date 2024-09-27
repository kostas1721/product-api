package database

import (
    "database/sql"
    _ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./products.db")
    if err != nil {
        panic(err)
    }

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
