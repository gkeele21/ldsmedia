package dbresourceview

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbresource"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbsiteuser"
	"time"
)

type ResourceView struct {
	ResourceViewID int64     `db:"resource_view_id"`
	ResourceID     int64     `db:"resource_id"`
	UserID         int64     `db:"user_id"`
	ViewedDate     time.Time `db:"viewed_date"`
	dbresource.Resource
	dbsiteuser.User
}

// ReadByID reads resource_view by id column
func ReadByID(ID int64) (*ResourceView, error) {
	u := ResourceView{}
	err := database.Get(&u, "SELECT * "+
		"FROM resource_view rv "+
		"INNER JOIN resource r ON r.resource_id = rv.resource_id "+
		"INNER JOIN site_user u ON u.user_id = rv.user_id "+
		"WHERE ur.resource_note_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.ResourceID = u.Resource.ResourceID
	u.UserID = u.User.UserID

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]ResourceView, error) {
	var users []ResourceView
	err := database.Select(&users, "SELECT * "+
		"FROM resource_view rv "+
		"INNER JOIN resource r ON r.resource_id = rv.resource_id"+
		"INNER JOIN site_user u ON u.user_id = rv.user_id")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ResourceView) error {
	_, err := database.Exec("DELETE FROM resource_view WHERE resource_view_id = ?", u.ResourceViewID)
	if err != nil {
		return fmt.Errorf("resource_view: couldn't delete resource_view %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ResourceView) error {
	res, err := database.Exec(database.BuildInsert("resource_view", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("resource_view: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("resource_view: couldn't get last inserted ID %S", err)
	}

	u.ResourceViewID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ResourceView) error {
	sql := database.BuildUpdate("resource_view", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("resource_view: couldn't update %s", err)
	}

	return nil
}

func Save(s *ResourceView) error {
	if s.ResourceViewID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
