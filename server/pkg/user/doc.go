package user

import (
	"errors"
	"time"

	"github.com/freshteapot/learnalist-api/server/pkg/event"
)

type ManagementStorage interface {
	FindUserUUID(search string) ([]string, error)
	GetLists(userUUID string) ([]string, error)
	DeleteUser(userUUID string) error
	DeleteList(listUUID string) error
	SaveInfo(userUUID string, info []byte) error
	GetInfo(userUUID string) ([]byte, error)
}

type ManagementSite interface {
	DeleteList(listUUID string) error
	DeleteUser(userUUID string) error
}

type Management interface {
	FindUser(search string) ([]string, error)
	DeleteUser(userUUID string) error
	SaveInfo(userUUID string, info []byte) error
	GetInfo(userUUID string) ([]byte, error)
}

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
	Register(idp string, identifier string, info []byte) (userUUID string, err error)
	Lookup(idp string, identifier string) (userUUID string, err error)
}

var ErrNotFound = errors.New("user-not-found")
