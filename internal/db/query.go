package db

import (
	"database/sql"

	"github.com/govnocods/JWT-Authorization/models"
)

func (d *DataBase) GetUser(username string) (*models.User, error) {
	user := &models.User{}

	query := `SELECT username, password FROM users WHERE username = ?`
	row := d.db.QueryRow(query, username)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
