package dbresourcemedia

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbresource"
)

type ResourceMedia struct {
	ResourceMediaID int64               `db:"resource_media_id"`
	ResourceID      int64               `db:"resource_id"`
	MediaType       database.NullString `db:"media_type"`
	MediaFormat     database.NullString `db:"media_format"`
	MediaURL        database.NullString `db:"media_url"`
	Data            database.NullString `db:"data"`
	dbresource.Resource
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*ResourceMedia, error) {
	u := ResourceMedia{}
	err := database.Get(&u, "SELECT * "+
		" FROM resource_media rm "+
		" INNER JOIN resource r ON r.resource_id = rm.resource_id "+
		" WHERE resource_media_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.ResourceID = u.Resource.ResourceID
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]ResourceMedia, error) {
	var users []ResourceMedia
	err := database.Select(&users, "SELECT * "+
		" FROM resource_media rm "+
		" INNER JOIN resource r ON r.resource_id = rm.resource_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ResourceMedia) error {
	_, err := database.Exec("DELETE FROM resource_media WHERE resource_media_id = ?", u.ResourceMediaID)
	if err != nil {
		return fmt.Errorf("resource_media: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ResourceMedia) error {
	res, err := database.Exec(database.BuildInsert("resource_media", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("resource_media: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("resource_media: couldn't get last inserted ID %S", err)
	}

	u.ResourceMediaID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ResourceMedia) error {
	sql := database.BuildUpdate("resource_media", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("resource_media: couldn't update %s", err)
	}

	return nil
}

func Save(s *ResourceMedia) error {
	if s.ResourceMediaID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
