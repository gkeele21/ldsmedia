package dbconferencesessiontype

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type ConferenceSessionType struct {
	ConferenceSessionTypeID int64  `db:"conference_session_type_id"`
	Name                    string `db:"name"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*ConferenceSessionType, error) {
	u := ConferenceSessionType{}
	err := database.Get(&u, "SELECT * FROM conference_session_type where conference_session_type_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]ConferenceSessionType, error) {
	var users []ConferenceSessionType
	err := database.Select(&users, "SELECT * FROM conference_session_type")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ConferenceSessionType) error {
	_, err := database.Exec("DELETE FROM conference_session_type WHERE conference_session_type_id = ?", u.ConferenceSessionTypeID)
	if err != nil {
		return fmt.Errorf("conference_session_type: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ConferenceSessionType) error {
	res, err := database.Exec(database.BuildInsert("conference_session_type", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("conference_session_type: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("conference_session_type: couldn't get last inserted ID %S", err)
	}

	u.ConferenceSessionTypeID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ConferenceSessionType) error {
	sql := database.BuildUpdate("conference_session_type", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("conference_session_type: couldn't update %s", err)
	}

	return nil
}

func Save(s *ConferenceSessionType) error {
	if s.ConferenceSessionTypeID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
