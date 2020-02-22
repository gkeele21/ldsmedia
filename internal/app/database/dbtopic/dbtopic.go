package dbtopic

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
)

const (
	DEFAULT_USER = 1
)

type Topic struct {
	TopicID       int64              `db:"topic_id"`
	UserID        database.NullInt64 `db:"user_id"`
	TopicName     string             `db:"topic_name"`
	ParentTopicID database.NullInt64 `db:"parent_topic_id"`
	Children      []*Topic
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*Topic, error) {
	u := Topic{}
	err := database.Get(&u, "SELECT * FROM topic where topic_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]Topic, error) {
	var users []Topic
	err := database.Select(&users, "SELECT * FROM topic")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *Topic) error {
	_, err := database.Exec("DELETE FROM topic WHERE topic_id = ?", u.TopicID)
	if err != nil {
		return fmt.Errorf("topic: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *Topic) error {
	res, err := database.Exec(database.BuildInsert("topic", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("topic: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("topic: couldn't get last inserted ID %S", err)
	}

	u.TopicID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Topic) error {
	sql := database.BuildUpdate("topic", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("topic: couldn't update %s", err)
	}

	return nil
}

func Save(s *Topic) error {
	if s.TopicID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}

// ReadByUserID reads all records in the database for the given userId
func ReadByUserID(userId int64) ([]Topic, error) {
	var users []Topic
	err := database.Select(&users, "SELECT * FROM topic WHERE user_id =?", userId)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// ReadDefaults reads all default records in the database
func ReadDefaults() ([]Topic, error) {
	var users []Topic
	err := database.Select(&users, "SELECT * FROM topic WHERE user_id = ?", DEFAULT_USER)
	if err != nil {
		return nil, err
	}

	return users, nil
}
