package cmd

import (
	"fmt"
	"syscall"

	"github.com/popshootjapan/getho-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with your account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil || email == "" {
			fmt.Print("email: ")
			fmt.Scan(&email)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil || password == "" {
			fmt.Print("password: ")
			passwordBytes, err := terminal.ReadPassword(syscall.Stdin)
			if err != nil {
				fmt.Println(err)
				return
			}
			password = fmt.Sprintf("%s", passwordBytes)
			fmt.Println("")
		}

		data, err := api.Singin(email, password)
		if err != nil {
			fmt.Println(err)
			return
		}

		viper.Set("token", data.Token)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("login completed")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "email")
	loginCmd.Flags().StringP("password", "p", "", "password")
}
