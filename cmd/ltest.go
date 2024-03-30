/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	hookcore "hooktest/hook-core"
	"io"
	"net/http"
	"time"

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

		sendExamplePayload(configPath, logPath)

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
	// ltestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sendExamplePayload(configFileName string, logFilenName string) {
	s := hookcore.NewServer(logFilenName)

	c := hookcore.ReadConfigFile(configFileName)

	req, err := c.ConstructRequest()
	if err != nil {
		panic(err)
	}

	responseChan := make(chan interface{})

	go func() {
		response, err := s.SendRequest(req)
		if err != nil {
			panic(err)
		}
		responseChan <- response
	}()

	select {
	case received := <-responseChan:
		httpResponse, ok := received.(*http.Response)
		if !ok {
			fmt.Println("Error: received value is not of type *http.Response")
			return
		}
		//defer httpResponse.Body.Close()

		body, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		fmt.Println("Response received:", string(body))
	case <-time.After(time.Second * 30):
		fmt.Println("Timeout: No response received within 30 seconds")
	}
}
