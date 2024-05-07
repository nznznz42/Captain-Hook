/*
Copyright © 2024 nznznz42
*/
package cmd

import (
	hookcore "hooktest/hook-core"

	"github.com/spf13/cobra"
)

var ltestCmd = &cobra.Command{
	Use:   "ltest",
	Short: "Runs Local Test System",
	Long:  `This command creates and sends a payload to your webhook locally`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]
		logPath := args[1]
		rflag, err := cmd.Flags().GetBool("randomize")
		if err != nil {
			panic("no bool")
		}
		lcmd := hookcore.NewCmd(configPath, logPath, rflag)

		hookcore.SendPayload(&lcmd)
	},
}

func init() {
	rootCmd.AddCommand(ltestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ltestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ltestCmd.Flags().BoolP("randomize", "r", false, "randomize payload")
}
