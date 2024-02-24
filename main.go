package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := NewServer("Logs/reqlog.log")

	reader := bufio.NewReader(os.Stdin)

	commandList := map[string]string{
		"help":  "prints commands",
		"start": "starts server",
		"stop":  "stops server",
		"exit":  "Exit Program",
	}

	for s.isRunning {
		fmt.Print("Enter command (type help to see list of available commands): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "help":
			count := 1
			for i, v := range commandList {
				fmt.Printf("\t%d. %s: %s\n", count, i, v)
				count += 1
			}
		case "start":
			s.Start()
		case "stop":
			s.Stop()
		case "exit":
			os.Exit(0)
		}
	}
}
