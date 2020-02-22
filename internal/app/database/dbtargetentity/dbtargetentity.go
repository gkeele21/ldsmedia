package dbtargetentity

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type TargetEntity struct {
	TargetEntityID int64  `db:"target_entity_id"`
	Type           string `db:"type"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*TargetEntity, error) {
	u := TargetEntity{}
	err := database.Get(&u, "SELECT * FROM target_entity where target_entity_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]TargetEntity, error) {
	var users []TargetEntity
	err := database.Select(&users, "SELECT * FROM target_entity")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *TargetEntity) error {
	_, err := database.Exec("DELETE FROM target_entity WHERE target_entity_id = ?", u.TargetEntityID)
	if err != nil {
		return fmt.Errorf("target_entity: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *TargetEntity) error {
	res, err := database.Exec(database.BuildInsert("target_entity", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("target_entity: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("target_entity: couldn't get last inserted ID %S", err)
	}

	u.TargetEntityID = ID

	return nil
}

// Update will update a record in the database
func Update(s *TargetEntity) error {
	sql := database.BuildUpdate("target_entity", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("target_entity: couldn't update %s", err)
	}

	return nil
}

func Save(s *TargetEntity) error {
	if s.TargetEntityID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
