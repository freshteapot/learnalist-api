package cron_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CronSuite struct {
	suite.Suite
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(CronSuite))
}