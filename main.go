
package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "CobraDemo",
	Long:  `Demo Program`,
}

var (
	greeting string
	subject string
)

func init() {
	helloCmd := &cobra.Command{
		Use: "hello",
		Short: "Say hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(greeting, subject)
		},
	}

	helloCmd.Flags().StringVar(&subject, "subject", "World", "to whom do you say it?")
	rootCmd.AddCommand(helloCmd)
}

func main()  {
	rootCmd.PersistentFlags().StringVar(&greeting, "greeting", "Hello", "what say you?")
	rootCmd.PersistentFlags().StringVar(&greeting, "greeting1", "Hello1", "what say you1?")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
	