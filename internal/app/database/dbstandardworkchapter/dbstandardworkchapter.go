package dbstandardworkchapter

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbstandardwookbook"
)

type StandardWorkChapter struct {
	StandardWorkChapterID int64 `db:"standard_work_chapter_id"`
	StandardWorkBookID    int64 `db:"standard_work_book_id"`
	ChapterNumber         int64 `db:"chapter_number"`
	dbstandardwookbook.StandardWorkBook
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*StandardWorkChapter, error) {
	u := StandardWorkChapter{}
	err := database.Get(&u, "SELECT * "+
		" FROM standard_work_chapter swc "+
		" INNER JOIN standard_wook_book swb ON swb.standard_work_book_id = swc.standard_work_book_id "+
		" WHERE standard_work_chapter_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.StandardWorkBookID = u.StandardWorkBook.StandardWorkBookID
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]StandardWorkChapter, error) {
	var users []StandardWorkChapter
	err := database.Select(&users, "SELECT * "+
		" FROM standard_work_chapter swc "+
		" INNER JOIN standard_wook_book swb ON swb.standard_work_book_id = swc.standard_work_book_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *StandardWorkChapter) error {
	_, err := database.Exec("DELETE FROM standard_work_chapter WHERE standard_work_chapter_id = ?", u.StandardWorkChapterID)
	if err != nil {
		return fmt.Errorf("standard_work_chapter: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *StandardWorkChapter) error {
	fmt.Println("Calling Insert on StandardWorkChapter...")
	res, err := database.Exec(database.BuildInsert("standard_work_chapter", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("standard_work_chapter: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("standard_work_chapter: couldn't get last inserted ID %S", err)
	}

	fmt.Printf("Inserted StandardWorkChapter %#v\n", u)
	u.StandardWorkChapterID = ID

	return nil
}

// Update will update a record in the database
func Update(s *StandardWorkChapter) error {
	sql := database.BuildUpdate("standard_work_chapter", s)
	fmt.Printf("SQL for update : %s\n", sql)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)
	if err != nil {
		return fmt.Errorf("standard_work_chapter: couldn't update %s", err)
	}

	return nil
}

func Save(s *StandardWorkChapter) error {
	if s.StandardWorkChapterID > 0 {
		fmt.Println("Calling Update")
		return Update(s)
	} else {
		return Insert(s)
	}
}
