package challenge

import (
	"encoding/json"
	"fmt"
	"net/http"

	"firebase.google.com/go/messaging"
	"github.com/freshteapot/learnalist-api/server/api/utils"
	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/sirupsen/logrus"
)

// removeUser when a user is deleted
// Currently we only remove the users entries, not any entries they created.
func (s ChallengeService) removeUser(entry event.Eventlog) {
	if entry.Kind != event.ApiUserDelete {
		return
	}

	b, err := json.Marshal(entry.Data)
	if err != nil {
		return
	}

	var moment event.EventUser
	json.Unmarshal(b, &moment)
	s.repo.DeleteUser(moment.UUID)
	s.logContext.WithFields(logrus.Fields{
		"user_uuid": moment.UUID,
		"event":     event.UserDeleted,
	}).Info("user removed")
}

func (s ChallengeService) eventChallengeDone(entry event.Eventlog) {
	if entry.Kind != EventChallengeDone {
		return
	}

	var moment EventChallengeDoneEntry
	b, _ := json.Marshal(entry.Data)
	json.Unmarshal(b, &moment)

	challengeUUID := moment.UUID
	if moment.Kind != EventKindPlank {
		s.logContext.WithFields(logrus.Fields{
			"kind":           moment.Kind,
			"challenge_uuid": challengeUUID,
			"user_uuid":      moment.UserUUID,
		}).Info("kind not supported, yet!")
		return
	}

	b, _ = json.Marshal(moment.Data)
	var record ChallengePlankRecord
	json.Unmarshal(b, &record)

	// Add the record
	// If it is a new entry, send a event that it was new.
	status, err := s.repo.AddRecord(challengeUUID, record.UUID, moment.UserUUID)
	if status == http.StatusInternalServerError {
		s.logContext.WithFields(logrus.Fields{
			"error":  err,
			"record": entry,
		}).Error("Failed to add record")
		return
	}

	if status != http.StatusCreated {
		s.logContext.WithFields(logrus.Fields{
			"error":  "duplicate entry",
			"record": entry,
		}).Error("Failed to add record")
		return
	}

	event.GetBus().Publish(event.TopicMonolog, event.Eventlog{
		Kind: EventChallengeNewRecord,
		Data: moment,
	})
}

func (s ChallengeService) eventChallengePushNotification(entry event.Eventlog) {
	allowed := []string{
		EventChallengeNewRecord,
		EventChallengeJoined,
		EventChallengeLeft,
	}

	if !utils.StringArrayContains(allowed, entry.Kind) {
		return
	}

	var (
		challengeUUID string
		userUUID      string
		title         string
		body          string
	)

	// TODO
	challengeName := "Join Tine Rehab"
	// TODO
	userDisplayName := "Tine"

	switch entry.Kind {
	case EventChallengeNewRecord:
		var moment EventChallengeDoneEntry
		b, _ := json.Marshal(entry.Data)
		json.Unmarshal(b, &moment)
		challengeUUID = moment.UUID
		userUUID = moment.UserUUID

		title = "Challenge update"
		body = fmt.Sprintf("%s added a record to %s", userDisplayName, challengeName)
	case EventChallengeJoined:
		var moment ChallengeJoined
		b, _ := json.Marshal(entry.Data)
		json.Unmarshal(b, &moment)
		challengeUUID = moment.UUID
		userUUID = moment.UserUUID

		title = "Challenge update"
		body = fmt.Sprintf("%s joined %s", userDisplayName, challengeName)
	case EventChallengeLeft:
		var moment ChallengeLeft
		b, _ := json.Marshal(entry.Data)
		json.Unmarshal(b, &moment)
		challengeUUID = moment.UUID
		userUUID = moment.UserUUID

		title = "Challenge update"
		body = fmt.Sprintf("%s left %s", userDisplayName, challengeName)
	}

	users, _ := s.challengeNotificationRepository.GetUsersInfo(challengeUUID)
	for _, user := range users {
		// Ignore the user who created the moment
		if user.UserUUID == userUUID {
			continue
		}
		fmt.Println("write notification for user", user.DisplayName, user)
		// TODO need nats option with other subject
		// TODO do we drop memory?
		// TODO mobile_device table needs a file
		// I now have enough informaiton to send to the topic to build the message
		// Or do I build the message here?
		// Note, it doesnt need to have setup the other channel, to publish to it
		data2 := map[string]string{
			"uuid":   challengeUUID,
			"who":    userDisplayName,
			"name":   challengeName,
			"action": entry.Kind,
		}

		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Data:  data2,
			Token: user.Token,
		}

		event.GetBus().Publish("challenges", event.Eventlog{
			Kind: "push-notification",
			Data: message,
		})

	}
}
