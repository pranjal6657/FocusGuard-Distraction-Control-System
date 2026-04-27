package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func showMenu() {
	fmt.Println("\n===== FocusGuard Menu =====")
	fmt.Println("1. Start Focus Mode")
	fmt.Println("2. Stop Focus Mode")
	fmt.Println("3. Show Stats")
	fmt.Println("4. Exit")
	fmt.Print("Enter choice: ")
}

func main() {
	loadConfig()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===================================")
	fmt.Println(" FocusGuard — Distraction Controller")
	fmt.Println(" Version 1.0 | Built in Go")
	fmt.Println("===================================")

	for {
		showMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			if !focusMode {
				focusMode = true
				go startFocusLoop()
			} else {
				fmt.Println("⚠️ Already running")
			}

		case "2":
			if focusMode {
				focusMode = false
				fmt.Println("🛑 Focus Mode Stopped")
			} else {
				fmt.Println("⚠️ Not running")
			}

		case "3":
			showStats()

		case "4":
			fmt.Println("👋 Exiting...")
			return

		default:
			fmt.Println("❌ Invalid choice")
		}
	}
}
