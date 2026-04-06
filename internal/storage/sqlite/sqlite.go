package sqlite

import (
	"database/sql"

	"github.com/iShinzoo/studentApi/internal/config"
	_ "modernc.org/sqlite" // to use this indirectly we have to use underscore (_) otherwise will show error
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stat, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stat.Close()

	result, err := stat.Exec(name, email, age)

	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return lastId, nil
}
