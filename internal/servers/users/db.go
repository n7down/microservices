package users

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        string
	Username  string
	Firstname string
	Lastname  string
}

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

// FIXME: use stored procedure that checks if the username exists - if it does return it
// otherwise insert the data
func (u *UsersDB) Create(id string, username string, password string, firstname string, lastname string) error {
	query := `INSERT INTO users(id, username, password, firstname, lastname, is_active) VALUES (?, ?, ?, ?, ?, 1)`
	//query := `INSERT INTO users(id, username, password, firstname, lastname, is_active) VALUES (?, ?, ?, ?, ?, 1) WHERE NOT EXISTS (SELECT username FROM users WHERE username = ?) LIMIT 1`
	_, err := u.db.Exec(query, id, username, password, firstname, lastname)

	//me, ok := err.(*mysql.MySQLError)
	//if !ok {
	//return err
	//}
	//if me.Number == 1062 {
	//return errors.New("username already exists")
	//}

	if err != nil {
		return err
	}

	return nil
}

func (u *UsersDB) Update(id string, username string, firstname string, lastname string) (User, error) {
	query := `UPDATE users SET username = ?, firstname = ?, lastname = ? WHERE id = ? AND is_active = 1`
	_, err := u.db.Exec(query, username, firstname, lastname, id)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:        id,
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
	}, nil
}

func (u *UsersDB) GetAll() ([]User, error) {
	users := []User{}
	query := `SELECT id, username, firstname, lastname FROM users WHERE is_active = 1`
	rows, err := u.db.Query(query)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UsersDB) ByID(id string) (User, error) {
	var user User
	query := `SELECT id, username, firstname, lastname FROM users WHERE is_active = 1 AND id = ?`
	row := u.db.QueryRow(query, id)
	switch err := row.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname); err {
	case sql.ErrNoRows:
		return user, errors.New("id invalid")
	case nil:
		return user, nil
	default:
		return user, err
	}
	return User{}, nil
}

func (u *UsersDB) Delete(id string) error {
	query := `UPDATE users SET is_active = 0 WHERE id = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
