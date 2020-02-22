package topic

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbtopic"
	"github.com/gkeele21/ldsmediaAPI/internal/pkg/topic"
	"github.com/gkeele21/ldsmediaAPI/pkg/log"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// RegisterRoutes sets up routes on a given nova.Server instance
func RegisterRoutes(g *echo.Group) {
	g.POST("/topics/addDefaultsToUser/:userId", addDefaultsToUser)
	g.GET("/topics/:userId", getUserTopics)
	g.PUT("/topics/:topicId", updateTopic)
	g.DELETE("/topics/:topicId", deleteTopic)
	g.POST("/topics/:topicId", addTopic)
}

type TopicData struct {
	TopicID   int64
	TopicName string
	UserID int64
}

// Response is the json representation of a user
//type Response struct {
//	User dbuser.User
//}

// getUserTopics retrieves all the topics for the user with the route parameter :userId
func getUserTopics(req echo.Context) error {
	var err error

	log.LogRequestData(req)
	searchID := req.Param("userId")
	userId, err := strconv.ParseInt(searchID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad user ID given")
	}

	// Build the object in a hierarchy with children
	newTopics := PrepareHierarchy(userId)
	fmt.Printf("# of UserTopics : %s", len(newTopics))

	return req.JSON(http.StatusOK, newTopics)
}

// addDefaultsToUser adds all the default topics to the user for the route parameter :userId
func addDefaultsToUser(req echo.Context) error {
	var err error

	log.LogRequestData(req)
	searchID := req.Param("userId")
	num, err := strconv.ParseInt(searchID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad user ID given")
	}

	userTopics, err := topic.AddDefaultToUser(num)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldn't add defaults", err)
	}

	fmt.Printf("# of UserTopics added : %s", len(userTopics))

	return req.JSON(http.StatusOK, userTopics)
}

func updateTopic(req echo.Context) error {
	var err error

	log.LogRequestData(req)
	//num := req.Param("topicId")
	//fmt.Printf("Num sent: %s\n", num)
	//topicId, err := strconv.ParseInt(num, 10, 64)
	//fmt.Printf("TopicID sent: %s\n", topicId)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "bad topic ID given")
	//}

	//tempId := req.QueryParam("TopicID")
	//topicId, _ := strconv.Atoi(tempId)
	//topicObj, err := dbtopic.ReadByID(int64(topicId))
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "bad topic ID given")
	//}

	tempTopic := new(TopicData)
	if err = req.Bind(tempTopic); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Please send a request body", 400)
	}
	if err != nil {
		req.Logger().Errorf("Error populating tempTopic struct : %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), 400)
	}

	topicObj, err := dbtopic.ReadByID(tempTopic.TopicID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad topic ID given")
	}

	fmt.Printf("TempTopic : %#v\n", tempTopic)

	if tempTopic.TopicName != "" {
		topicObj.TopicName = tempTopic.TopicName
	}

	ret := dbtopic.Update(topicObj)
	if ret != nil {
		req.Logger().Errorf("Error updating topic record : %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, ret.Error())
	}

	return req.JSON(http.StatusOK, topicObj)

}

func deleteTopic(req echo.Context) error {
	var err error

	log.LogRequestData(req)

	num := req.Param("userId")
	fmt.Printf("Num sent: %s\n", num)
	topicId, err := strconv.ParseInt(num, 10, 64)
	fmt.Printf("TopicID sent: %s\n", topicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad topic ID given")
	}

	topicObj, err := dbtopic.ReadByID(topicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad topic ID given")
	}

	userId := topicObj.UserID
	ret := dbtopic.Delete(topicObj)
	if ret != nil {
		req.Logger().Errorf("Error deleting topic record : %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, ret.Error())
	}

	newTopics := PrepareHierarchy(userId.Int64)
	return req.JSON(http.StatusOK, newTopics)
}

func addTopic(req echo.Context) error {
	var err error

	log.LogRequestData(req)
	num := req.Param("userId")
	fmt.Printf("Num sent: %s\n", num)
	topicId, err := strconv.ParseInt(num, 10, 64)
	fmt.Printf("TopicID sent: %s\n", topicId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad topic ID given")
	}

	topicName := req.QueryParam("topicName")
	if topicName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "new topic name not given")
	}

	tempTopic := new(TopicData)
	if err = req.Bind(tempTopic); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Please send a request body", 400)
	}
	if err != nil {
		req.Logger().Errorf("Error populating tempTopic struct : %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), 400)
	}

	fmt.Printf("ParentTopic : %#v\n", tempTopic)

	topicObj := dbtopic.Topic{
		TopicID:       0,
		UserID:        database.ToNullInt(tempTopic.UserID, true),
		TopicName:     topicName,
		ParentTopicID: database.ToNullInt(tempTopic.TopicID, true),
		Children:      nil,
	}

	ret := dbtopic.Save(&topicObj)
	if ret != nil {
		req.Logger().Errorf("Error creating topic record : %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, ret.Error())
	}
	newTopics := PrepareHierarchy(tempTopic.UserID)
	return req.JSON(http.StatusOK, newTopics)
}

func PrepareHierarchy(userId int64) []*dbtopic.Topic {

	all := &dbtopic.Topic{
		TopicID:       0,
		UserID:        database.ToNullInt(userId, false),
		TopicName:     "All",
		ParentTopicID: database.ToNullInt(int64(0), false),
		Children:      nil,
	}
	getChildren(all)
	return all.Children
}

func getChildren(topic *dbtopic.Topic) {
	var query string
	var children []dbtopic.Topic
	var err error
	if topic.TopicID == 0 {
		query = "select * from topic where user_id = ? and parent_topic_id IS NULL"
		err = database.Select(&children, query, topic.UserID)
	} else {
		query = "select * from topic where user_id = ? and parent_topic_id = ?"
		err = database.Select(&children, query, topic.UserID, topic.TopicID)
	}

	if err != nil {
		fmt.Printf("Error getting children : %s\n", err)
	}

	for _, child := range children {
		newChild := child
		topic.Children = append(topic.Children, &newChild)
		getChildren(&newChild)
	}
}
