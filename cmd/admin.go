package cmd

import (
	"fmt"
	"instagram_bot/database"
	"instagram_bot/models"
	"log"

	"github.com/spf13/cobra"
)

var (
	email    string
	password string
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "This will create an admin user",
	Run: func(cmd *cobra.Command, args []string) {

		u := models.NewUser(email, password)

		db, err := database.New()
		if err != nil {
			log.Fatal(err)
		}
		err = db.CreateUser(&u)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Created admin %s\n", email)

	},
}

func init() {
	rootCmd.AddCommand(adminCmd)

	adminCmd.PersistentFlags().StringVar(&email, "email", "", "Email address")
	adminCmd.PersistentFlags().StringVar(&password, "password", "", "Password")

	adminCmd.MarkFlagRequired("email")
	adminCmd.MarkFlagRequired("password")
}
