package sqlite

import (
	"database/sql"

	"github.com/Vivek09Chahal/studentsAPI/internal/config"
)

type Sqlite struct {
    Db *sql.DB
}

func New(cfg  *config.Config) (*Sqlite, error) {
    db, err := sql.Open("sqlite3", cfg.StoragePath)

    if err != nil {
        return nil, err
    }

    _, err = db.Exec(`CREATE TABLE IF NOT EXIST students(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            age INTEGER,
            email TEXT
        )`)

    if err != nil {
        return nil, err
    }

    return &Sqlite{
        Db: db,
    }, nil 
    
}