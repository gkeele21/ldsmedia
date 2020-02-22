package dbstandardwork

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type StandardWork struct {
	StandardWorkID int64  `db:"standard_work_id"`
	Name           string `db:"_name"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*StandardWork, error) {
	u := StandardWork{}
	err := database.Get(&u, "SELECT * FROM standard_work where standard_work_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]StandardWork, error) {
	var users []StandardWork
	err := database.Select(&users, "SELECT * FROM standard_work")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *StandardWork) error {
	_, err := database.Exec("DELETE FROM standard_work WHERE standard_work_id = ?", u.StandardWorkID)
	if err != nil {
		return fmt.Errorf("standard_work: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *StandardWork) error {
	res, err := database.Exec(database.BuildInsert("standard_work", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("standard_work: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("standard_work: couldn't get last inserted ID %S", err)
	}

	u.StandardWorkID = ID

	return nil
}

// Update will update a record in the database
func Update(s *StandardWork) error {
	sql := database.BuildUpdate("standard_work", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("standard_work: couldn't update %s", err)
	}

	return nil
}

func Save(s *StandardWork) error {
	if s.StandardWorkID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
