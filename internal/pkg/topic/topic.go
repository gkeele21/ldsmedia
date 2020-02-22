package topic

import (
	"github.com/gkeele21/ldsmediaAPI/internal/app/database"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbtopic"
	"log"
)

func AddDefaultToUser(userId int64) ([]dbtopic.Topic, error) {
	var userTopics []dbtopic.Topic

	log.Printf("Adding the default topics to userId %s", userId)

	defaultTopics, err := dbtopic.ReadDefaults()
	if err != nil {
		log.Fatalf("Error getting default topics : %s", err)
		return userTopics, err
	}

	for _, defaultTopic := range defaultTopics {
		// add this default topic to the user
		topic := new(dbtopic.Topic)
		topic.TopicName = defaultTopic.TopicName
		topic.UserID = database.ToNullInt(userId, true)
		topic.ParentTopicID = defaultTopic.ParentTopicID

		dbtopic.Save(topic)

		userTopics = append(userTopics, *topic)
	}

	return userTopics, nil
}
