package sqllite

import (
	"database/sql"

	"github.com/gunjanghate/learning-go/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqllite struct {
	Db *sql.DB
}

func New(cfg config.Config) (*Sqllite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}


	_ , err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students(
	id Integer Primary Key Autoincrement,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	age INTEGER NOT NULL
	);
	`)

	if err != nil {
		return nil, err
	}

	return &Sqllite{
		Db: db,
	}, nil
}

func (s *Sqllite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT into students (name, email, age) VALUES (?,?,?)") /// placeholder for sql statement to prevent sql injection
    
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	
	
	return lastid, nil

}
