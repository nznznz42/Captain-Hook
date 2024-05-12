/*
Copyright © 2024 nznznz42
*/
package cmd

import (
	"fmt"
	hookcore "hooktest/hook-core"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

var ctestCmd = &cobra.Command{
	Use:   "ctest [domain: wss://example.com/ws] [port]",
	Short: "runs cloud test system",
	Long:  `This command uses the cloud component to live test your webhook.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		port := args[1]
		portInt, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalf("invalid Port")
		}

		//wg := sync.WaitGroup{}
		//wg.Add(1)
		c := hookcore.Newclient(domain)
		defer c.Conn.Close()
		fmt.Printf("link : %s", c.URL)

		fields := []string{"Header", "Method", "Body"}
		go c.Stream(os.Stdout, fields, portInt)
		//wg.Wait()
		//go c.Stream(os.Stdout, fields, portInt)

		// Wait for interrupt signal (Ctrl+C)
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
	},
}

func init() {
	rootCmd.AddCommand(ctestCmd)
}
