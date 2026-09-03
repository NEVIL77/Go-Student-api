package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/NEVIL77/students-api/internal/config"
	"github.com/NEVIL77/students-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.Storage_path)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	return &Sqlite{
		DB: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	// step to create a database

	stmt, err := s.DB.Prepare("INSERT INTO students(name, email, age) VALUES(?,?,?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {

	stmt, err := s.DB.Prepare("SELECT id,name,email,age FROM students WHERE id=?")

	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("No student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("failed to find student")
	}
	return student, nil
}
