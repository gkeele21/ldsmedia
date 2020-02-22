package main

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/cmd/server/api/handlers/topic"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbconference"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbconferencesession"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbtopic"
)

func main() {
	//testUpdatingConferenceSesssion()
	testHierarchy()
}

func testHierarchy() {
	userId := 2
	userTopics, err := dbtopic.ReadByUserID(int64(userId))
	if err != nil {
		fmt.Println("Error : %#v", err)
	}

	newTopics := topic.PrepareHierarchy(userTopics, int64(userId))

	fmt.Printf("UserTopics : %#v\n", newTopics)
}

func testUpdatingConferenceSesssion() {
	conf, err := dbconference.ReadByID(12)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	fmt.Printf("Conference : %#v\n", conf)

	confsession, err := dbconferencesession.ReadByID(13)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	fmt.Printf("ConferenceSession : %#v\n", confsession)
	confsession.Name = "Testing"
	err = dbconferencesession.Save(confsession)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
}
