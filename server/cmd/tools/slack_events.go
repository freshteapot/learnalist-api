package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/freshteapot/learnalist-api/server/pkg/challenge"
	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/freshteapot/learnalist-api/server/pkg/logging"
	"github.com/freshteapot/learnalist-api/server/pkg/plank"
	"github.com/freshteapot/learnalist-api/server/pkg/spaced_repetition"
)

var slackEventsCMD = &cobra.Command{
	Use:   "slack-events",
	Short: "Read events and write to slack",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger()
		logger.Info("Read events")
		event.SetDefaultSettingsForCMD()

		webhook := viper.GetString("server.events.slack.webhook")
		if webhook == "" {
			panic("Webhook shouldnt be empty")
		}

		event.SetupEventBus(logger.WithField("context", "event-bus-setup"))

		reader := NewSlackEvents(webhook, logger.WithField("context", "slack-events"))
		event.GetBus().Subscribe("slack-listener", reader.Read)

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		select {
		case <-signals:
		}
		event.GetBus().Close()
	},
}

type SlackEvents struct {
	webhook    string
	logContext logrus.FieldLogger
}

func NewSlackEvents(webhook string, logContext logrus.FieldLogger) SlackEvents {
	return SlackEvents{
		webhook:    webhook,
		logContext: logContext,
	}
}

func (s SlackEvents) Read(entry event.Eventlog) {
	var msg slack.WebhookMessage

	switch entry.Kind {
	case event.ApiUserRegister:
		b, _ := json.Marshal(entry.Data)
		var moment event.EventUser
		json.Unmarshal(b, &moment)
		msg.Text = fmt.Sprintf("%s: user %s registered via %s", entry.Kind, moment.UUID, moment.Kind)
	case event.ApiUserLogin:
		b, _ := json.Marshal(entry.Data)
		var moment event.EventUser
		json.Unmarshal(b, &moment)
		msg.Text = fmt.Sprintf("%s: user %s logged in via %s", entry.Kind, moment.UUID, moment.Kind)
	case event.ApiUserLogout:
	case event.BrowserUserLogout:
		b, _ := json.Marshal(entry.Data)
		var moment event.EventUser
		json.Unmarshal(b, &moment)
		via := "api"
		if entry.Kind == event.BrowserUserLogout {
			via = "browser"
		}

		clearing := "current session"
		if moment.Kind == event.KindUserLogoutSessions {
			clearing = "all sessions"
		}

		msg.Text = fmt.Sprintf("%s: user %s logged out via %s, clearing %s", entry.Kind, moment.UUID, via, clearing)
	case event.ApiUserDelete:
		b, _ := json.Marshal(entry.Data)
		var moment event.EventUser
		json.Unmarshal(b, &moment)
		msg.Text = fmt.Sprintf("%s: user %s should be deleted", entry.Kind, moment.UUID)
	case event.ApiListSaved:
		b, _ := json.Marshal(entry.Data)
		var moment event.EventList
		json.Unmarshal(b, &moment)
		msg.Text = fmt.Sprintf(`list:%s (%s) %s by user:%s`, moment.UUID, moment.Data.Info.SharedWith, moment.Action, moment.UserUUID)
	case event.ApiListDelete:
		b, _ := json.Marshal(entry.Data)
		var moment event.EventList
		json.Unmarshal(b, &moment)
		msg.Text = fmt.Sprintf("list:%s deleted by user:%s", moment.UUID, moment.UserUUID)
	case spaced_repetition.EventApiSpacedRepetition:
		b, _ := json.Marshal(entry.Data)
		var moment spaced_repetition.EventSpacedRepetition
		json.Unmarshal(b, &moment)

		if moment.Kind == spaced_repetition.EventKindNew {
			msg.Text = fmt.Sprintf("User:%s added a new entry for spaced based learning", moment.Data.UserUUID)
		}

		if moment.Kind == spaced_repetition.EventKindViewed {
			when := "na"
			if moment.Action == "incr" {
				when = "later"
			}

			if moment.Action == "decr" {
				when = "sooner"
			}
			msg.Text = fmt.Sprintf("User:%s will be reminded %s of entry:%s", moment.Data.UserUUID, when, moment.Data.UUID)
		}

		if moment.Kind == spaced_repetition.EventKindDeleted {
			msg.Text = fmt.Sprintf("User:%s removed entry:%s from spaced based learning", moment.Data.UserUUID, moment.Data.UUID)
		}
	case plank.EventApiPlank:
		b, _ := json.Marshal(entry.Data)
		var moment plank.EventPlank
		json.Unmarshal(b, &moment)
		if moment.Kind == plank.EventKindNew {
			msg.Text = fmt.Sprintf("User:%s added a plank:%s", moment.UserUUID, moment.Data.UUID)
		}

		if moment.Kind == plank.EventKindDeleted {
			msg.Text = fmt.Sprintf("User:%s deleted a plank:%s", moment.UserUUID, moment.Data.UUID)
		}
	case challenge.EventChallengeDone:
		b, _ := json.Marshal(entry.Data)
		var moment challenge.EventChallengeDoneEntry
		json.Unmarshal(b, &moment)
		if moment.Kind == challenge.EventKindPlank {
			b, _ = json.Marshal(moment.Data)
			var record plank.HttpRequestInput
			json.Unmarshal(b, &record)
			msg.Text = fmt.Sprintf("User:%s added a plank:%s to challenge:%s", moment.UserUUID, record.UUID, moment.UUID)
		} else {
			return
		}
		// TODO Add challene notification

	default:
		b, _ := json.Marshal(entry)
		fmt.Println(string(b))
		msg.Text = entry.Kind
	}

	err := slack.PostWebhook(s.webhook, &msg)
	if err != nil {
		s.logContext.Panic(err)
	}
}

func init() {
	viper.SetDefault("server.events.slack.webhook", "")
	viper.BindEnv("server.events.slack.webhook", "EVENTS_SLACK_WEBHOOK")
}
