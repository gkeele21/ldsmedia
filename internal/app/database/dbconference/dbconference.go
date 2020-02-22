package dbconference

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"time"
)

type Conference struct {
	ConferenceID int64     `db:"conference_id"`
	StartingDate time.Time `db:"starting_date"`
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*Conference, error) {
	u := Conference{}
	fmt.Println("Getting db record")
	err := database.Get(&u, "SELECT * FROM conference where conference_id = ?", ID)
	fmt.Println("Done Getting db record")
	if err != nil {
		fmt.Printf("Error : %s", err)
		return nil, err
	}

	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]Conference, error) {
	var users []Conference
	err := database.Select(&users, "SELECT * FROM conference")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *Conference) error {
	_, err := database.Exec("DELETE FROM conference WHERE conference_id = ?", u.ConferenceID)
	if err != nil {
		return fmt.Errorf("conference: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *Conference) error {
	res, err := database.Exec(database.BuildInsert("conference", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("conference: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("conference: couldn't get last inserted ID %S", err)
	}

	u.ConferenceID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Conference) error {
	sql := database.BuildUpdate("conference", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("conference: couldn't update %s", err)
	}

	return nil
}

func Save(s *Conference) error {
	if s.ConferenceID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
