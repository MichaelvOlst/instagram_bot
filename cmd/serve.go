package cmd

import (
	"instagram_bot/database"
	"instagram_bot/server"
	"log"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command runs the api and webserver",

	Run: func(cmd *cobra.Command, args []string) {

		db, err := database.New()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		s := server.New(db)

		s.Start()

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
