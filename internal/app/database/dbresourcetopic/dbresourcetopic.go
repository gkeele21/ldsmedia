package dbresourcetopic

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

type ResourceTopic struct {
	ResourceTopicID int64 `db:"resource_topic_id"`
	ResourceID      int64 `db:"resource_id"`
	TopicID         int64 `db:"topic_id"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*ResourceTopic, error) {
	u := ResourceTopic{}
	err := database.Get(&u, "SELECT * FROM resource_topic where resource_topic_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]ResourceTopic, error) {
	var users []ResourceTopic
	err := database.Select(&users, "SELECT * FROM resource_topic")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *ResourceTopic) error {
	_, err := database.Exec("DELETE FROM resource_topic WHERE resource_topic_id = ?", u.ResourceTopicID)
	if err != nil {
		return fmt.Errorf("resource_topic: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *ResourceTopic) error {
	res, err := database.Exec(database.BuildInsert("resource_topic", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("resource_topic: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("resource_topic: couldn't get last inserted ID %S", err)
	}

	u.ResourceTopicID = ID

	return nil
}

// Update will update a record in the database
func Update(s *ResourceTopic) error {
	sql := database.BuildUpdate("resource_topic", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("resource_topic: couldn't update %s", err)
	}

	return nil
}

func Save(s *ResourceTopic) error {
	if s.ResourceTopicID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
