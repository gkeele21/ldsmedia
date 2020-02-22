package dbstandardwookbook

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbstandardwork"
)

type StandardWorkBook struct {
	StandardWorkBookID int64               `db:"standard_work_book_id"`
	StandardWorkID     int64               `db:"standard_work_id"`
	BookName           string              `db:"book_name"`
	DisplayOrder       database.NullInt64  `db:"display_order"`
	ChapterTitle       database.NullString `db:"chapter_title"`
	dbstandardwork.StandardWork
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*StandardWorkBook, error) {
	u := StandardWorkBook{}
	err := database.Get(&u, "SELECT * "+
		" FROM standard_work_book swb "+
		" INNER JOIN standard_work sw ON sw.standard_work_id = swb.standard_work_id "+
		" WHERE standard_work_book_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.StandardWorkID = u.StandardWork.StandardWorkID
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]StandardWorkBook, error) {
	var users []StandardWorkBook
	err := database.Select(&users, "SELECT * "+
		" FROM standard_work_book swb "+
		" INNER JOIN standard_work sw ON sw.standard_work_id = swb.standard_work_id ")

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *StandardWorkBook) error {
	_, err := database.Exec("DELETE FROM standard_work_book WHERE standard_work_book_id = ?", u.StandardWorkBookID)
	if err != nil {
		return fmt.Errorf("standard_work_book: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *StandardWorkBook) error {
	res, err := database.Exec(database.BuildInsert("standard_work_book", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("standard_work_book: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("standard_work_book: couldn't get last inserted ID %S", err)
	}

	u.StandardWorkBookID = ID

	return nil
}

// Update will update a record in the database
func Update(s *StandardWorkBook) error {
	sql := database.BuildUpdate("standard_work_book", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("standard_work_book: couldn't update %s", err)
	}

	return nil
}

func Save(s *StandardWorkBook) error {
	if s.StandardWorkBookID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
