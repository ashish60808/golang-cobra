package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "cobraintro",
	Short: "This tools gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Hello! ", username)
	},
}
 var username, password string

func main() {
	cmd.PersistentFlags().StringVarP(&username,"username","u","","Username")
	cmd.PersistentFlags().StringVarP(&username,"password","p","","Password")
	cmd.Execute()
}