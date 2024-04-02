/*
Copyright © 2024 nznznz42
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ctestCmd represents the ctest command
var ctestCmd = &cobra.Command{
	Use:   "ctest",
	Short: "runs cloud test system",
	Long:  `This command uses the cloud component to live test your webhook.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ctest called")
	},
}

func init() {
	rootCmd.AddCommand(ctestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
