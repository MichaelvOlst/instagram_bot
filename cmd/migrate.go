package cmd

import (
	"instagram_bot/database"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Will create and migrate the database",
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := os.Stat(viper.GetString("database_file")); os.IsNotExist(err) {
			var file, err = os.Create(viper.GetString("database_file"))
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}

		db, err := database.New()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
