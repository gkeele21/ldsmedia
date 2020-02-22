package dbsource

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type Source struct {
	SourceID int64               `db:"source_id"`
	Name     string              `db:"name" json:"SourceName"`
	BaseURL  database.NullString `db:"base_url"`
	Platform database.NullString `db:"platform"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*Source, error) {
	u := Source{}
	err := database.Get(&u, "SELECT * FROM source where source_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]Source, error) {
	var users []Source
	err := database.Select(&users, "SELECT * FROM source")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *Source) error {
	_, err := database.Exec("DELETE FROM source WHERE source_id = ?", u.SourceID)
	if err != nil {
		return fmt.Errorf("source: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *Source) error {
	res, err := database.Exec(database.BuildInsert("source", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("source: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("source: couldn't get last inserted ID %S", err)
	}

	u.SourceID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Source) error {
	sql := database.BuildUpdate("source", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("user: couldn't update %s", err)
	}

	return nil
}

func Save(s *Source) error {
	if s.SourceID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
