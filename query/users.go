package query

import (
	"database/sql"
	"errors"
	"github.com/ekkapob/buybits/model"
)

func UserExist(db *sql.DB, username string) bool {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&id)
	return err != sql.ErrNoRows
}

func AddUser(db *sql.DB, user model.User) error {
	stmt, err := db.Prepare("INSERT INTO users (username, password, firstname, lastname) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.Username,
		user.Password,
		user.Firstname,
		user.Lastname,
	)
	return err
}

func GetUser(db *sql.DB, username string) (user model.User, err error) {

	if !UserExist(db, username) {
		return user, errors.New("Username doesn't exist.")
	}

	user.Username = username
	err = db.QueryRow("SELECT id, password, firstname, lastname FROM users WHERE username = $1", username).Scan(&user.Id, &user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
