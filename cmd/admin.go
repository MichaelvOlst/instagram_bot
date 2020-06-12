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
	create   bool
	delete   bool
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "This will create or delete an admin user",
	Run: func(cmd *cobra.Command, args []string) {

		if !create && !delete {
			log.Fatal("You need to provide the create or delete flag")
		}

		if create {
			if email == "" || password == "" {
				log.Fatal("You need to fill in an email and password")
			}

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
		}

		if delete {
			if email == "" {
				log.Fatal("You need to specify an email")
			}

			u := models.NewUser(email, password)

			db, err := database.New()
			if err != nil {
				log.Fatal(err)
			}
			err = db.DeleteUserByEmail(&u)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Admin deleted %s\n", email)

		}

	},
}

func init() {
	rootCmd.AddCommand(adminCmd)

	adminCmd.PersistentFlags().StringVar(&email, "email", "", "Email address")
	adminCmd.PersistentFlags().StringVar(&password, "password", "", "Password")
	adminCmd.Flags().BoolVar(&create, "create", false, "Create an user")
	adminCmd.Flags().BoolVar(&delete, "delete", false, "Delete an user")

	adminCmd.MarkFlagRequired("email")
	adminCmd.MarkFlagRequired("password")
}
