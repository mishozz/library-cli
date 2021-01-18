package cli

import (
	"fmt"

	"github.com/mishozz/library-cli/client"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with username and password",
	Long:  "Login with username and password",
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")

		token, err := client.User.Login(email, "")
		if err != nil {
			fmt.Printf("Unable to login. Check your username and password")
		} else {
			fmt.Println("Login succesful.")
			fmt.Printf("Your token is: %s", token)
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  "Logout from your account",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		msg, err := client.User.Logout(token)
		if err != nil {
			fmt.Printf("Unable to logout. Check you token!")
		} else {
			fmt.Printf(msg)
		}
	},
}

var takeBookCmd = &cobra.Command{
	Use:   "take",
	Short: "Take book",
	Long:  "Take book from the library",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		email, _ := cmd.Flags().GetString("email")
		isbn, _ := cmd.Flags().GetString("isbn")

		msg, err := client.User.TakeBook(token, email, isbn)
		if err != nil {
			fmt.Printf("Unable to take book from the library")
		} else {
			fmt.Printf(msg)
		}
	},
}

var returnBookCmd = &cobra.Command{
	Use:   "return",
	Short: "Return book",
	Long:  "Return book in the library",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		email, _ := cmd.Flags().GetString("email")
		isbn, _ := cmd.Flags().GetString("isbn")

		err := client.User.ReturnBook(token, email, isbn)
		if err != nil {
			if err == client.UnauthorizedErr {
				fmt.Printf("You need to be authorized to access this route")
			} else {
				fmt.Printf("Unable to return your book")
			}
		} else {
			fmt.Printf("Successfully returned you book")
		}
	},
}

var getUsersCmd = &cobra.Command{
	Use:   "get-all-users",
	Short: "Get all users",
	Long:  "Get all users of the library",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")

		respString, err := client.User.GetAllUsers(token)
		if err != nil {
			fmt.Printf("Unable to fetch users")
		} else {
			fmt.Printf(respString)
		}
	},
}

var getUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Get user of the library",
	Long:  "Get user of the library",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		email, _ := cmd.Flags().GetString("email")

		respString, err := client.User.GetUser(token, email)
		if err != nil {
			fmt.Printf("Unable to fetch user with email %s", email)
		} else {
			fmt.Printf(respString)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(takeBookCmd)
	rootCmd.AddCommand(returnBookCmd)
	rootCmd.AddCommand(getUsersCmd)
	rootCmd.AddCommand(getUserCmd)

	loginCmd.Flags().StringP("email", "e", "", "Set your email")
	loginCmd.MarkFlagRequired("email")

	logoutCmd.Flags().StringP("token", "t", "", "Your jwt token")
	logoutCmd.MarkFlagRequired("token")

	takeBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	takeBookCmd.Flags().StringP("email", "e", "", "Set your email")
	takeBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	takeBookCmd.MarkFlagRequired("token")
	takeBookCmd.MarkFlagRequired("email")
	takeBookCmd.MarkFlagRequired("isbn")

	returnBookCmd.Flags().StringP("token", "t", "", "Your jwt token")
	returnBookCmd.Flags().StringP("email", "e", "", "Set your email")
	returnBookCmd.Flags().StringP("isbn", "i", "", "Isbn of the book")
	returnBookCmd.MarkFlagRequired("token")
	returnBookCmd.MarkFlagRequired("email")
	returnBookCmd.MarkFlagRequired("isbn")

	getUsersCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getUsersCmd.MarkFlagRequired("token")

	getUserCmd.Flags().StringP("token", "t", "", "Your jwt token")
	getUserCmd.Flags().StringP("email", "e", "", "Set your email")
	getUserCmd.MarkFlagRequired("token")
	getUserCmd.MarkFlagRequired("email")
}
