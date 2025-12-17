package sqllite

import (
	"database/sql"
	"fmt"

	"github.com/gunjanghate/learning-go/internal/types"

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


func (s *Sqllite) GetStudentById(id int64) (types.Student, error){
	stmt, err := s.Db.Prepare("SELECT * FROM students where id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows{
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}
		return types.Student{}, fmt.Errorf("query err : %q", err)
	}
	return student, nil
}


func (s *Sqllite) GetStudents() ([]types.Student, error){
	stmt, err := s.Db.Prepare("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []types.Student
	
	for rows.Next(){
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}