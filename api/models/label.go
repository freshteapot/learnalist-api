package models

import (
	"errors"
	"net/http"
	"strings"

	"github.com/freshteapot/learnalist-api/api/i18n"
)

type UserLabel struct {
	Label    string `db:"label"`
	UserUuid string `db:"user_uuid"`
}

type AlistLabel struct {
	Label     string `db:"label"`
	UserUuid  string `db:"user_uuid"`
	AlistUuid string `db:"alist_uuid"`
}

func NewUserLabel(label string, user string) *UserLabel {
	userLabel := &UserLabel{
		Label:    label,
		UserUuid: user,
	}
	return userLabel
}

func NewAlistLabel(label string, user string, alist string) *AlistLabel {
	alistLabel := &AlistLabel{
		Label:     label,
		UserUuid:  user,
		AlistUuid: alist,
	}
	return alistLabel
}

func ValidateLabel(label string) error {
	if label == "" {
		return errors.New(i18n.ValidationWarningLabelNotEmpty)
	}

	if len(label) > 20 {
		return errors.New(i18n.ValidationWarningLabelToLong)
	}
	return nil
}

func (dal *DAL) PostUserLabel(label *UserLabel) (int, error) {
	statusCode := http.StatusBadRequest
	err := ValidateLabel(label.Label)
	if err != nil {
		return statusCode, err
	}

	query := "INSERT INTO user_labels(label, user_uuid) VALUES (:label, :user_uuid);"

	_, err = dal.Db.NamedExec(query, label)
	statusCode = http.StatusCreated
	if err != nil {
		statusCode = http.StatusBadRequest
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			statusCode = http.StatusOK
		}
	}
	return statusCode, err
}

// Parse in the user uuid and get back their labels
func (dal *DAL) PostAlistLabel(label *AlistLabel) (int, error) {
	statusCode := http.StatusBadRequest
	err := ValidateLabel(label.Label)
	if err != nil {
		return statusCode, err
	}

	query := "INSERT INTO alist_labels(label, user_uuid, alist_uuid) VALUES (:label, :user_uuid, :alist_uuid);"

	_, err = dal.Db.NamedExec(query, label)
	statusCode = http.StatusCreated
	if err != nil {
		statusCode = http.StatusBadRequest
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			statusCode = http.StatusOK
		}
	}
	return statusCode, err
}

func (dal *DAL) GetUserLabels(uuid string) ([]string, error) {
	var labels = []string{}

	query := `
SELECT label
FROM user_labels
WHERE user_uuid=$1
UNION
SELECT label
FROM alist_labels
WHERE user_uuid=$1

`
	err := dal.Db.Select(&labels, query, uuid)
	if err != nil {
		return labels, err
	}
	return labels, err
}

// Pass in the label and the user (uuid) to remove them from the tables
func (dal *DAL) RemoveUserLabel(label string, user string) error {
	query1 := "DELETE FROM user_labels WHERE user_uuid=$1 AND label=$2"
	query2 := "DELETE FROM alist_labels WHERE user_uuid=$1 AND label=$2"

	tx := dal.Db.MustBegin()
	tx.MustExec(query1, user, label)
	tx.MustExec(query2, user, label)
	err := tx.Commit()
	return err
}

func (dal *DAL) RemoveLabelsForAlist(uuid string) error {
	if uuid == "" {
		return nil
	}

	query := "DELETE FROM alist_labels WHERE alist_uuid=$1"

	tx := dal.Db.MustBegin()
	tx.MustExec(query, uuid)
	err := tx.Commit()
	return err
}
