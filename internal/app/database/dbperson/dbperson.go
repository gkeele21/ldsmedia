package dbperson

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type Person struct {
	PersonID      int64               `db:"person_id"`
	PersonTitleID database.NullInt64  `db:"person_person_title_id" json:"PersonPersonTitleID"`
	Gender        database.NullString `db:"gender"`
	FirstName     string              `db:"first_name" json:"PersonFirstName"`
	MiddleName    database.NullString `db:"middle_name" json:"PersonMiddleName"`
	LastName      database.NullString `db:"last_name" json:"PersonLastName"`
	//dbpersontitle.PersonTitle
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*Person, error) {
	u := Person{}
	err := database.Get(&u, "SELECT * "+
		" FROM person p "+
		//" INNER JOIN person_title pt ON pt.person_title_id = p.person_title_id "+
		" WHERE person_id = ?", ID)
	if err != nil {
		return nil, err
	}

	//u.PersonTitleID = database.ToNullInt(u.PersonTitle.PersonTitleID, true)
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]Person, error) {
	var users []Person
	err := database.Select(&users, "SELECT * "+
		" FROM person p "+
		" INNER JOIN person_title pt ON pt.person_title_id = p.person_title_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *Person) error {
	_, err := database.Exec("DELETE FROM person WHERE person_id = ?", u.PersonID)
	if err != nil {
		return fmt.Errorf("person: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *Person) error {
	res, err := database.Exec(database.BuildInsert("person", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("person: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("person: couldn't get last inserted ID %S", err)
	}

	u.PersonID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Person) error {
	sql := database.BuildUpdate("person", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("person: couldn't update %s", err)
	}

	return nil
}

func Save(s *Person) error {
	if s.PersonID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
