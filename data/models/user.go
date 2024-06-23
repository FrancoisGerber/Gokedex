package models

import (
	"Gokedex/data"
	"Gokedex/utils"
)

// Implements IModel
type User struct {
	Id       int64
	Username string
	Email    string
	Password string
}

func (u *User) GetAll() (users []User, err error) {
	selectAllQuery := `SELECT Id, Username, Email, Password FROM Users`
	rows, err := data.DB.Query(selectAllQuery)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) Get(id int64) error {
	selectSingleQuery := `SELECT Id, Username, Email, Password FROM Users WHERE Id = ?`
	row := data.DB.QueryRow(selectSingleQuery, id)

	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Password)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) Create() error {
	insertQuery := `
	INSERT INTO Users(Username, Email, Password)
	VALUES(?,?,?)
	`

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := data.DB.Exec(insertQuery, u.Username, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	u.Id, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update(id int64) (int64, error) {
	updateQuery := `
	UPDATE Users
	SET Username=?,Email=?
	WHERE Id = ?`

	result, err := data.DB.Exec(updateQuery, u.Username, u.Email, id)

	if err != nil {
		return 0, err
	}

	rowsChanged, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsChanged, nil
}

func (u *User) Delete(id int64) (bool, error) {
	deleteQuery := `
	DELETE FROM Users
	WHERE Id = ?`

	result, err := data.DB.Exec(deleteQuery, id)

	if err != nil {
		return false, err
	}

	rowsChanged, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rowsChanged > 0, nil
}

func (u *User) ValidateCredentials() (bool, string, error) {
	validateQuery := `SELECT Id, Password FROM Users WHERE Username = ? OR Email = ?`

	row := data.DB.QueryRow(validateQuery, u.Username, u.Username)

	var dbPassword string
	err := row.Scan(&u.Id, &dbPassword)

	if err != nil {
		return false, "", err
	}

	valid, err := utils.ComparePassword(dbPassword, u.Password)

	if err != nil {
		return false, "", err
	}

	if valid {
		token, err := utils.GenerateJWT(u.Id, u.Username)

		if err != nil {
			return false, "", err
		}

		return valid, token, nil
	}

	return valid, "", nil
}
