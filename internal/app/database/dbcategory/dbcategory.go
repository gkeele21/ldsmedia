package dbcategory

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type Category struct {
	CategoryID   int64  `db:"category_id"`
	CategoryName string `db:"category_name"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*Category, error) {
	u := Category{}
	err := database.Get(&u, "SELECT * FROM category where category_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]Category, error) {
	var users []Category
	err := database.Select(&users, "SELECT * FROM category")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *Category) error {
	_, err := database.Exec("DELETE FROM category WHERE category_id = ?", u.CategoryID)
	if err != nil {
		return fmt.Errorf("category: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *Category) error {
	res, err := database.Exec(database.BuildInsert("category", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("category: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("category: couldn't get last inserted ID %S", err)
	}

	u.CategoryID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Category) error {
	sql := database.BuildUpdate("category", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("category: couldn't update %s", err)
	}

	return nil
}

func Save(s *Category) error {
	if s.CategoryID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
