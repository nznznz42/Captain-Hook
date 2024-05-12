/*
Copyright © 2024 nznznz42
*/
package cmd

import (
	"fmt"
	hookcore "hooktest/hook-core"
	"log"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "This command just runs the cached ltest command",
	Long:  `Checks the cache for a stored ltest command and runs it.`,
	Run: func(cmd *cobra.Command, args []string) {
		cacheCmd, err := checkCache()
		if err != nil {
			log.Fatalf("Cache Corrupted")
		}
		if cacheCmd == "" {
			fmt.Print("Cache is empty")
		} else if cacheCmd == "ltest" {
			lcmd, err := hookcore.DeserializeLcmd()
			if err != nil {
				log.Fatalf("cache corrupted")
			}
			hookcore.SendPayload(lcmd)
		} else if cacheCmd == "ctest" {
			ccmd, err := hookcore.DeserializeCcmd()
			if err != nil {
				log.Fatalf("cache corrupted")
			}
			ExecuteCtest(ccmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func checkCache() (string, error) {
	cacheStatus, err := hookcore.IsFileEmpty("Cache/LastCmdCache.json")
	if err != nil {
		panic("noooo")
	}

	if !cacheStatus {
		cmd, err := hookcore.ReadCmdCache()
		if err != nil {
			panic("nooo")
		}
		return cmd, nil
	}

	return "", err
}
