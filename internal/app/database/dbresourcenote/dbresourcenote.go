package dbresourcenote

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbresource"
	"time"
)

type ResourceNote struct {
	ResourceNoteID int64               `db:"resource_note_id"`
	ResourceID     int64               `db:"resource_id"`
	Notes          database.NullString `db:"notes"`
	CreatedDate    time.Time           `db:"created_date"`
	dbresource.Resource
}

// ReadByID reads resource_note by id column
func ReadByID(ID int64) (*ResourceNote, error) {
	u := ResourceNote{}
	err := database.Get(&u, "SELECT * "+
		"FROM resource_note rn "+
		"INNER JOIN resource r ON r.resource_id = rn.resource_id "+
		"WHERE ur.resource_note_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.ResourceID = u.Resource.ResourceID

	return &u, nil
}

// ReadAll reads all resourcenotes in the database
func ReadAll() ([]ResourceNote, error) {
	var users []ResourceNote
	err := database.Select(&users, "SELECT * "+
		"FROM resource_note ur "+
		"INNER JOIN resource r ON r.resource_id = ur.resource_id")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ResourceNote) error {
	_, err := database.Exec("DELETE FROM resource_note WHERE resource_note_id = ?", u.ResourceNoteID)
	if err != nil {
		return fmt.Errorf("resource_note: couldn't delete resource_note %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ResourceNote) error {
	res, err := database.Exec(database.BuildInsert("resource_note", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("resource_note: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("resource_note: couldn't get last inserted ID %S", err)
	}

	u.ResourceNoteID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ResourceNote) error {
	sql := database.BuildUpdate("resource_note", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("resource_note: couldn't update %s", err)
	}

	return nil
}

func Save(s *ResourceNote) error {
	if s.ResourceNoteID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
