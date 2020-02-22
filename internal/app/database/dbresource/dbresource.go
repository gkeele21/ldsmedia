package dbresource

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbcategory"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbperson"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbsiteuser"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbsource"
	"time"
)

type Resource struct {
	ResourceID       int64               `db:"resource_id"`
	ResUserID        int64               `db:"res_user_id"`
	ResCategoryID    database.NullInt64  `db:"res_category_id"`
	ResSourceID      database.NullInt64  `db:"res_source_id"`
	ResPersonID      database.NullInt64  `db:"res_person_id"`
	ResPersonTitleID database.NullInt64  `db:"res_person_title_id"`
	AddedDate        time.Time           `db:"added_date"`
	ResourceDate     time.Time           `db:"resource_date"`
	Title            string              `db:"title"`
	Description      database.NullString `db:"description"`
	URL              database.NullString `db:"url"`
	Tags             database.NullString `db:"tags"`
	//dbsiteuser.User
	//dbcategory.Category
	//dbsource.Source
	//dbperson.Person
	//dbpersontitle.PersonTitle
}

type ResourceFull struct {
	Resource
	dbsiteuser.User
	dbcategory.Category
	dbsource.Source
	dbperson.Person
	//dbpersontitle.PersonTitle
}

// ReadByID reads record by id column
func ReadByID(ID int64) (*Resource, error) {
	u := Resource{}
	err := database.Get(&u, "SELECT * "+
		" FROM resource r "+
		" INNER JOIN site_user su ON su.user_id = r.user_id "+
		" LEFT JOIN category c ON c.category_id = r.category_id "+
		" LEFT JOIN source s ON s.source_id = r.source_id "+
		" LEFT JOIN person p ON p.person_id = r.person_id "+
		" LEFT JOIN person_title pt ON pt.person_title_id = r.person_title_id "+
		" WHERE r.resource_id = ?", ID)
	if err != nil {
		return nil, err
	}

	//u.UserID = u.User.UserID
	//u.PersonTitleID = database.ToNullInt(u.PersonTitle.PersonTitleID, true)
	//u.PersonID = database.ToNullInt(u.Person.PersonID, true)
	//u.SourceID = database.ToNullInt(u.Source.SourceID, true)
	//u.CategoryID = database.ToNullInt(u.Category.CategoryID, true)
	return &u, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]Resource, error) {
	var users []Resource
	err := database.Select(&users, "SELECT * "+
		" FROM resource r ")
	//" INNER JOIN site_user su ON su.user_id = r.user_id ")
	//" LEFT JOIN category c ON c.category_id = r.category_id "+
	//" LEFT JOIN source s ON s.source_id = r.source_id "+
	//" LEFT JOIN person p ON p.person_id = r.person_id "+
	//" LEFT JOIN person_title pt ON pt.person_title_id = r.person_title_id ")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Delete deletes a record from the database
func Delete(u *Resource) error {
	_, err := database.Exec("DELETE FROM resource WHERE resource_id = ?", u.ResourceID)
	if err != nil {
		return fmt.Errorf("resource: couldn't delete %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(u *Resource) error {
	res, err := database.Exec(database.BuildInsert("resource", u), database.GetArguments(*u)...)

	if err != nil {
		return fmt.Errorf("resource: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("resource: couldn't get last inserted ID %S", err)
	}

	u.ResourceID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Resource) error {
	sql := database.BuildUpdate("resource", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("resource: couldn't update %s", err)
	}

	return nil
}

func Save(s *Resource) error {
	if s.ResourceID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}

// ReadByUserID reads all records in the database for the given userId
func ReadByUserID(userId int64) ([]ResourceFull, error) {
	//var users []Resource
	var rFull []ResourceFull
	query := "SELECT * " +
		" FROM resource r " +
		" INNER JOIN site_user su ON su.user_id = r.res_user_id " +
		" LEFT JOIN category c ON c.category_id = r.res_category_id " +
		" LEFT JOIN source s ON s.source_id = r.res_source_id " +
		" LEFT JOIN person p ON p.person_id = r.res_person_id " +
		//" LEFT JOIN person_title pt ON pt.person_person_title_id = r.res_person_title_id "+
		" WHERE r.res_user_id = ?"
	// fetch all places from the db
	err := database.Select(&rFull, query, userId)

	if err != nil {
		fmt.Printf("Error getting resource %#v", err)
		return nil, err
	}

	//for i, v := range rFull {
	//	//v.Resource.UserID = v.User.UserID
	//	fmt.Printf("i : %#v\n", i)
	//	fmt.Printf("v : %#v\n", v)
	//	fmt.Printf("User Id : %s\n", v.User.UserID)
	//}
	// iterate over each row
	//for rows.Next() {
	//
	//	err = rows.StructScan(rFull)
	//	fmt.Printf("ResourceFull : %#v\n", rFull)
	//}
	//err := database.Select(&users, , userId)

	return rFull, nil
}
