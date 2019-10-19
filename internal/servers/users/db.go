package users

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type UsersDB struct {
	db *sql.DB
}

func NewUsersDB() (*UsersDB, error) {
	usersDB := &UsersDB{}
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/users")
	if err != nil {
		return usersDB, err
	}

	err = db.Ping()
	if err != nil {
		return usersDB, err
	}
	return &UsersDB{db: db}, nil
}

func (u *UsersDB) GetPassword(username string) (string, string, error) {
	var (
		id       string
		password string
	)
	query := `SELECT id, password FROM users WHERE username = ?`
	row := u.db.QueryRow(query, username)
	switch err := row.Scan(&id, &password); err {
	case sql.ErrNoRows:
		return "", "", errors.New("id invalid")
	case nil:
		return id, password, nil
	default:
		return "", "", err
	}
}

func (u *UsersDB) Create(id string, username string, password string, firstname string, lastname string) error {
	query := `INSERT INTO users(id, username, password, firstname, lastname, is_active) VALUES (?, ?, ?, ?, ?, 1)`
	_, err := u.db.Exec(query, id, username, password, firstname, lastname)
	if err != nil {
		return err
	}
	return nil
}

// TODO: implement
func (u *UsersDB) Update(id string, username string, password string, firstname string, lastname string) error {
	return nil
}

// TODO: list only the is_active = 1 users
func (u *UsersDB) List() error {
	return nil
}

// TODO: list only the is_active = 1 users
func (u *UsersDB) ByID(id string) error {
	return nil
}

func (u *UsersDB) Delete(id string) error {
	query := `UPDATE users SET is_active = 0 WHERE id = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
