package user

import (
	"time"

	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/freshteapot/learnalist-api/server/pkg/openapi"
)

const (
	IDPKindEmail  = "email"
	IDPKindUserID = "id"
)

type management struct {
	storage  ManagementStorage
	site     ManagementSite
	insights event.Insights
}

type UserInfoFromUsernameAndPassword struct {
	UserUUID string
	Username string
	Hash     string
}

// TODO what is this
type UserInfo struct {
	UserUUID  string
	Challenge string
	Created   time.Time
}

type UserSession struct {
	Token     string
	UserUUID  string
	Challenge string
	// TODO I want to know what client it is, web, mobile, chrome-extension, so I can handle different responses.
	Created time.Time
}

type UserPreference struct {
	UserUUID         string                       `json:"user_uuid,omitempty"`
	DisplayName      string                       `json:"display_name,omitempty"`
	CreatedVia       string                       `json:"created_via,omitempty"`
	DailyReminder    *UserPreferenceDailyReminder `json:"daily_reminder,omitempty"`
	Apps             *UserPreferenceApps          `json:"app_settings,omitempty"` // TODO good to know, but lets not run with it yet
	LastActive       *LastActive                  `json:"last_active,omitempty"`
	SpacedRepetition *SpacedRepetition            `json:"spaced_repetition,omitempty"`
	Acl              ACL                          `json:"acl"`
}

type UserPreferenceDailyReminder struct {
	RemindV1 *openapi.RemindDailySettings `json:"remind_v1,omitempty"` // Needed first :D
	PlankV1  *openapi.RemindDailySettings `json:"plank_v1,omitempty"`
}

type UserPreferenceApps struct {
	PlankV1  *openapi.MobilePlankAppV1Settings `json:"plank_v1,omitempty"` // Only nice to sync between app and web, not needed yet
	RemindV1 *openapi.AppSettingsRemindV1      `json:"remind_v1,omitempty"`
}

// TODO actually use
type LastActive struct {
	Plank            string `json:"plank,omitempty"`             // UTC int64? or string time.RFC3339Nano
	SpacedRepetition string `json:"spaced_repetition,omitempty"` // UTC int64? or string time.RFC3339Nano
}

type SpacedRepetition struct {
	ListsOvertime []string `json:"lists_overtime"` // UTC int64? or string time.RFC3339Nano
}

type ACL struct {
	PublicListWrite int `json:"list_public_write"`
}

// TODO is this the correct name
type UserInfoRepository interface {
	Get(userUUID string) (UserPreference, error)
	Save(userUUID string, pref UserPreference) error
}

type ManagementStorage interface {
	UserExists(userUUID string) bool
	FindUserUUID(search string) ([]string, error)
	GetLists(userUUID string) ([]string, error)
	DeleteUser(userUUID string) error
	DeleteList(listUUID string) error
}

type ManagementSite interface {
	DeleteList(listUUID string) error
	DeleteUser(userUUID string) error
}

type Management interface {
	UserExists(userUUID string) bool
	FindUser(search string) ([]string, error)
	DeleteUser(userUUID string) error
}

// TODO rename to UserSession
type Session interface {
	NewSession(userUUID string) (session UserSession, err error)
	// Create create a session with a unique challenge, send the challenge in the oauth2 flow
	// The string returned is the actual challenge
	CreateWithChallenge() (string, error)
	// Activate update the challenge with the userUUID and token
	Activate(session UserSession) error
	GetUserUUIDByToken(token string) (userUUID string, err error)
	IsChallengeValid(challenge string) (bool, error)

	RemoveSessionForUser(userUUID string, token string) error
	// RemoveSessionsForUser remove all sessions for a user
	RemoveSessionsForUser(userUUID string) error
	// RemoveExpiredChallenges remove challenges that were never activated
	RemoveExpiredChallenges() error
}

type UserWithUsernameAndPassword interface {
	Register(username string, hash string) (info UserInfoFromUsernameAndPassword, err error)
	Lookup(username string, hash string) (userUUID string, err error)
}

type UserFromIDP interface {
	Register(idp string, kind string, identifier string, info []byte) (userUUID string, err error)
	Lookup(idp string, kind string, identifier string) (userUUID string, err error)
}
