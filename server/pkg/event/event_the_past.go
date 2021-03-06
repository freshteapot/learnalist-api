package event

import (
	"github.com/freshteapot/learnalist-api/server/api/uuid"
	log "github.com/sirupsen/logrus"
)

type eventAction struct {
	Todo     string
	Happened string
}

// Action A list of possible event options.
var Action = eventAction{
	Todo:     "todo",
	Happened: "happened",
}

type eventType struct {
	UserNew     string
	UserSaved   string
	UserDeleted string

	ListDeleted string
	ListNew     string
	ListSaved   string

	LinkitParse string
}

// Type A list of possible event types.
var Type = eventType{
	UserNew:     "user.new",
	UserSaved:   "user.saved",
	UserDeleted: "user.deleted",

	ListDeleted: "list.deleted",
	ListNew:     "list.new",
	ListSaved:   "list.saved",

	LinkitParse: "linkit.parse",
}

// This could be changed for something else entirely
var eventLog *log.Logger

func init() {
	eventLog = log.New()
	eventLog.Formatter = new(log.JSONFormatter)
}

// New Allow to override the default way of saving events
func New(logger *log.Logger) {
	eventLog = logger
}

// Todo A record to trigger a todo, implying that an action should take this uuid and do something.
func Todo(uuid uuid.Info, event string) {
	record(uuid, event, Action.Todo)
}

// Happened A record that it happened, implying that it is a historical reference of something happening.
func Happened(uuid uuid.Info, event string) {
	record(uuid, event, Action.Happened)
}

func record(uuid uuid.Info, eventType string, eventAction string) {
	eventLog.WithFields(log.Fields{
		"uuid":   uuid.ToString(),
		"action": eventAction,
	}).Info(eventType)
}
