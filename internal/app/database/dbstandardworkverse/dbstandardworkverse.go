package dbstandardworkverse

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbstandardworkchapter"
)

type StandardWorkVerse struct {
	StandardWorkVerseID   int64 `db:"standard_work_verse_id"`
	StandardWorkChapterID int64 `db:"standard_work_chapter_id"`
	VerseNumber           int64 `db:"verse_number"`
	dbstandardworkchapter.StandardWorkChapter
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*StandardWorkVerse, error) {
	u := StandardWorkVerse{}
	err := database.Get(&u, "SELECT * "+
		" FROM standard_work_verse swv "+
		" INNER JOIN standard_work_chapter swc ON swc.standard_work_chapter_id = swv.standard_work_chapter_id "+
		" WHERE standard_work_verse_id = ?", ID)
	if err != nil {
		return nil, err
	}

	u.StandardWorkChapterID = u.StandardWorkChapter.StandardWorkChapterID
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]StandardWorkVerse, error) {
	var users []StandardWorkVerse
	err := database.Select(&users, "SELECT * "+
		" FROM standard_work_verse swv "+
		" INNER JOIN standard_work_chapter swc ON swc.standard_work_chapter_id = swv.standard_work_chapter_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *StandardWorkVerse) error {
	_, err := database.Exec("DELETE FROM standard_work_verse WHERE standard_work_verse_id = ?", u.StandardWorkVerseID)
	if err != nil {
		return fmt.Errorf("standard_work_verse: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *StandardWorkVerse) error {
	res, err := database.Exec(database.BuildInsert("standard_work_verse", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("standard_work_verse: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("standard_work_verse: couldn't get last inserted ID %S", err)
	}

	u.StandardWorkVerseID = ID

	return nil
}

// Update will update a record in the database
func Update(s *StandardWorkVerse) error {
	sql := database.BuildUpdate("standard_work_verse", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("standard_work_verse: couldn't update %s", err)
	}

	return nil
}

func Save(s *StandardWorkVerse) error {
	if s.StandardWorkVerseID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}
