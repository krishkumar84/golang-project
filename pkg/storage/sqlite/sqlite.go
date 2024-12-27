package sqlite

import (
	"database/sql"

	"github.com/krishkumar84/golang-project/pkg/config"
	_"github.com/mattn/go-sqlite3"
	"github.com/krishkumar84/golang-project/pkg/types"
	"fmt"

)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_,err =db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL 
	)`)
    if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil 
}

func (s *Sqlite) CreateUser(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO users (name, email, age) VALUES (?, ?, ?)")
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

func (s *Sqlite) GetUserById(id int64)(types.User, error){
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM users WHERE id = ? LIMIT 1")
	if err != nil {
		return types.User{}, err
	}
	defer stmt.Close()
	var user types.User

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fmt.Errorf("user not found %s",fmt.Sprint(id))
		}
		return types.User{}, err
	}
	return user, nil
}

func (s *Sqlite) GetAllUsers()([]types.User, error){
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []types.User
	for rows.Next(){
		var user types.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}