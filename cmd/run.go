/*
Copyright © 2024 nznznz42
*/
package cmd

import (
	"fmt"
	hookcore "hooktest/hook-core"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "This command just runs the cached ltest command",
	Long:  `Checks the cache for a stored ltest command and runs it.`,
	Run: func(cmd *cobra.Command, args []string) {
		cacheCmd := checkCache()
		if cacheCmd == nil {
			fmt.Print("Cache is empty. specify args using the ltest command.")
		} else {
			hookcore.SendPayload(cacheCmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func checkCache() *hookcore.Ltestcmd {
	cacheStatus, err := hookcore.IsFileEmpty()
	if err != nil {
		panic("noooo")
	}

	if !cacheStatus {
		cmd, err := hookcore.Deserialize()
		if err != nil {
			panic("nooo")
		}
		return cmd
	}
	return nil
}
