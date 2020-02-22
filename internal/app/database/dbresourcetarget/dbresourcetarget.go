package dbresourcetarget

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbresource"
)

type ResourceTarget struct {
	ResourceTargetID    int64              `db:"resource_target_id"`
	ResourceID          int64              `db:"resource_id"`
	TargetEntityIDStart int64              `db:"target_entity_id_start"`
	TargetEntityIDEnd   database.NullInt64 `db:"target_entity_id_end"`
	dbresource.Resource
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*ResourceTarget, error) {
	u := ResourceTarget{}
	err := database.Get(&u, "SELECT * "+
		"FROM resource_target rt "+
		"INNER JOIN resource r ON r.resource_id = rt.resource_id "+
		"WHERE rt.resource_target_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]ResourceTarget, error) {
	var users []ResourceTarget
	err := database.Select(&users, "SELECT * "+
		"FROM resource_target rt "+
		"INNER JOIN resource r ON r.resource_id = rt.resource_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ResourceTarget) error {
	_, err := database.Exec("DELETE FROM resource_target WHERE resource_target_id = ?", u.ResourceTargetID)
	if err != nil {
		return fmt.Errorf("resource_target: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ResourceTarget) error {
	res, err := database.Exec(database.BuildInsert("resource_target", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("resource_target: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("resource_target: couldn't get last inserted ID %S", err)
	}

	u.ResourceTargetID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ResourceTarget) error {
	sql := database.BuildUpdate("resource_target", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("resource_target: couldn't update %s", err)
	}

	return nil
}

func Save(s *ResourceTarget) error {
	if s.ResourceTargetID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
