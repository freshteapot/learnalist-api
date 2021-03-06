package user

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/freshteapot/learnalist-api/server/api/database"
	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/freshteapot/learnalist-api/server/pkg/event/staticsite"
	"github.com/freshteapot/learnalist-api/server/pkg/logging"
	"github.com/freshteapot/learnalist-api/server/pkg/user"
	userStorage "github.com/freshteapot/learnalist-api/server/pkg/user/sqlite"
	"github.com/freshteapot/learnalist-api/server/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a user based on a username or email",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger()
		event.SetDefaultSettingsForCMD()
		os.Setenv("EVENTS_STAN_CLIENT_ID", "tools-user-mangement")
		event.SetupEventBus(logger.WithField("context", "tools-user-find"))

		dsn := viper.GetString("server.sqlite.database")
		search := args[0]
		if search == "" {
			fmt.Println("Nothing to search for, means nothing to find")
			return
		}

		db := database.NewDB(dsn)
		userManagement := user.NewManagement(
			userStorage.NewSqliteManagementStorage(db),
			staticsite.NewSiteManagementViaEvents(),
			event.NewInsights(logger),
		)

		userUUIDs, err := userManagement.FindUser(search)

		if err != nil {
			fmt.Println("Something went wrong")
			fmt.Println(err)
			// Printing this, as it might contain 2 results
			fmt.Println(userUUIDs)
			os.Exit(1)
		}

		b, _ := json.Marshal(userUUIDs)
		fmt.Println(utils.PrettyPrintJSON(b))
	},
}
