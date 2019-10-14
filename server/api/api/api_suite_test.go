package api_test

import (
	"testing"

	"github.com/freshteapot/learnalist-api/server/alists/pkg/hugo/mocks"
	"github.com/freshteapot/learnalist-api/server/api/api"
	"github.com/freshteapot/learnalist-api/server/api/database"
	"github.com/freshteapot/learnalist-api/server/api/models"
	aclSqlite "github.com/freshteapot/learnalist-api/server/pkg/acl/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Test Suite")
}

var dal *models.DAL
var m api.Manager

var _ = BeforeSuite(func() {
	db := database.NewTestDB()
	acl := aclSqlite.NewAcl(db)
	dal = models.NewDAL(db, acl)
	hugoHelper := new(mocks.HugoSiteBuilder)

	m = api.Manager{
		Datastore:  dal,
		Acl:        acl,
		HugoHelper: hugoHelper,
	}
})