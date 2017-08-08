package userRepo

import (
	"database/sql"
	user "github.com/danilojunS/widgets-spa-api/business/entities/user"
	database "github.com/danilojunS/widgets-spa-api/infra/database"
	_ "github.com/lib/pq"
)

// Create a user
func Create(u user.User) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var lastInsertID int
	err = db.QueryRow("INSERT INTO users (name, gravatar) VALUES($1,$2) returning id;", u.Name, u.Gravatar).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Read users
func Read(id int) ([]user.User, error) {
	db, err := database.Connect()
	if err != nil {
		return []user.User{}, err
	}
	defer db.Close()

	var rows *sql.Rows
	if id == 0 {
		rows, err = db.Query("SELECT * FROM users")
		if err != nil {
			return []user.User{}, err
		}
	} else {
		rows, err = db.Query("SELECT * FROM users WHERE id=$1", id)
		if err != nil {
			return []user.User{}, err
		}
	}

	var users []user.User

	for rows.Next() {
		var id int
		var name string
		var gravatar string

		err = rows.Scan(&id, &name, &gravatar)
		if err != nil {
			return []user.User{}, err
		}

		user := user.User{ID: id, Name: name, Gravatar: gravatar}
		users = append(users, user)
	}

	return users, nil
}

// Update user
func Update(u user.User) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET name=$1, gravatar=$2 WHERE id=$3")
	if err != nil {
		return 0, err
	}

	_, err = stmt.Exec(u.Name, u.Gravatar, u.ID)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Delete user
func Delete(u user.User) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return 0, err
	}

	_, err = stmt.Exec(u.ID)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Clear all users
func Clear() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
