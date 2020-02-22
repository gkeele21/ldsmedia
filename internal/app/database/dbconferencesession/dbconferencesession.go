package dbconferencesession

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbconference"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbconferencesessiontype"
)

type ConferenceSession struct {
	ConferenceSessionID     int64              `db:"conference_session_id"`
	ConferenceID            int64              `db:"conference_id"`
	Name                    string             `db:"name"`
	ConferenceSessionTypeID database.NullInt64 `db:"conference_session_type_id"`
	DisplayOrder            database.NullInt64 `db:"display_order"`
	SessionDate             database.NullTime  `db:"session_date"`
	dbconference.Conference
	dbconferencesessiontype.ConferenceSessionType
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*ConferenceSession, error) {
	u := ConferenceSession{}
	err := database.Get(&u, "SELECT * "+
		" FROM conference_session cs "+
		" INNER JOIN conference c ON c.conference_id = cs.conference_id "+
		" INNER JOIN conference_session_type cst ON cst.conference_session_type_id = cs.conference_session_type_id "+
		" WHERE conference_session_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.ConferenceID = u.Conference.ConferenceID
	u.ConferenceSessionTypeID = database.ToNullInt(u.ConferenceSessionType.ConferenceSessionTypeID, true)
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]ConferenceSession, error) {
	var users []ConferenceSession
	err := database.Select(&users, "SELECT * "+
		" FROM conference_session cs "+
		" INNER JOIN conference c ON c.conference_id = cs.conference_id "+
		" INNER JOIN conference_session_type cst ON cst.conference_session_type_id = cs.conference_session_type_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ConferenceSession) error {
	_, err := database.Exec("DELETE FROM conference_session WHERE conference_session_id = ?", u.ConferenceSessionID)
	if err != nil {
		return fmt.Errorf("conference_session: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ConferenceSession) error {
	res, err := database.Exec(database.BuildInsert("conference_session", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("conference_session: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("conference_session: couldn't get last inserted ID %S", err)
	}

	u.ConferenceSessionID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ConferenceSession) error {
	sql := database.BuildUpdate("conference_session", s)
	fmt.Printf("SQL : %s\n", sql)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("conference_session: couldn't update %s", err)
	}

	return nil
}

func Save(s *ConferenceSession) error {
	if s.ConferenceSessionID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
