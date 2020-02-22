package dbpersontitle

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type PersonTitle struct {
	PersonTitleID      int64               `db:"person_title_id"`
	Title              string              `db:"title" json:"PersonTitleTitle"`
	IsApostle          bool                `db:"is_apostle"`
	IsGeneralAuthority bool                `db:"is_general_authority"`
	TitlePrefix        database.NullString `db:"title_prefix"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*PersonTitle, error) {
	u := PersonTitle{}
	err := database.Get(&u, "SELECT * FROM person_title where person_title_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]PersonTitle, error) {
	var users []PersonTitle
	err := database.Select(&users, "SELECT * FROM person_title")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *PersonTitle) error {
	_, err := database.Exec("DELETE FROM person_title WHERE person_title_id = ?", u.PersonTitleID)
	if err != nil {
		return fmt.Errorf("person_title: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *PersonTitle) error {
	res, err := database.Exec(database.BuildInsert("person_title", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("person_title: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("person_title: couldn't get last inserted ID %S", err)
	}

	u.PersonTitleID = ID

	return nil
}

// Update will update a record in the database
func Update(s *PersonTitle) error {
	sql := database.BuildUpdate("person_title", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("person_title: couldn't update %s", err)
	}

	return nil
}

func Save(s *PersonTitle) error {
	if s.PersonTitleID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
