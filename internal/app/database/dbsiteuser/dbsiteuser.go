package dbsiteuser

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"time"
)

type User struct {
	UserID        int64               `db:"user_id"`
	FirstName     string              `db:"first_name"`
	LastName      database.NullString `db:"last_name"`
	Email         string              `db:"email"`
	CreatedDate   time.Time           `db:"created_date"`
	Username      database.NullString `db:"username"`
	UserPassword  database.NullString `db:"user_password"`
	Cell          database.NullString `db:"cell"`
	IsActive      int64               `db:"is_active"`
	LastLoginDate database.NullTime   `db:"last_login_date"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*User, error) {
	u := User{}
	err := database.Get(&u, "SELECT * FROM site_user where user_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]User, error) {
	var users []User
	err := database.Select(&users, "SELECT * FROM site_user")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *User) error {
	_, err := database.Exec("DELETE FROM site_user WHERE user_id = ?", u.UserID)
	if err != nil {
		return fmt.Errorf("user: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *User) error {
	res, err := database.Exec(database.BuildInsert("site_user", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("user: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("user: couldn't get last inserted ID %S", err)
	}

	u.UserID = ID

	return nil
}

// Update will update a record in the database
func Update(s *User) error {
	sql := database.BuildUpdate("site_user", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("user: couldn't update %s", err)
	}

	return nil
}

func Save(s *User) error {
	if s.UserID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}

// ReadByUsername reads user by username column
func ReadByUsername(username string) (*User, error) {
	u := User{}
	err := database.Get(&u, "SELECT * FROM site_user WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadActiveByUsername reads an active user by username column
func ReadActiveByUsername(username string) (*User, error) {
	u := User{}
	err := database.Get(&u, "SELECT * FROM site_user WHERE is_active = 1 AND username = ?", username)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
