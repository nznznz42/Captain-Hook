/*
Copyright © 2024 nznznz42
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hooktest",
	Short: "A Simple Local First Webhook Tester",
	Long:  `hooktest is a simple webhook tester that allows you to test a webhook locally by generating random payloads or if you want, you can deploy the cloud component and test your webhook the traditional way.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
