package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Config struct {
	BlockedApps []string `json:"blocked_apps"`
}

var config Config
var focusMode bool
var distractionCount int
var startTime time.Time

// Safe apps (won’t be blocked)
var safeApps = []string{
	"visual studio",
	"code",
	"cmd",
	"powershell",
}

// Load config
func loadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("❌ Error reading config:", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &config)
}

// Get active window
func getActiveWindowTitle() string {
	psScript := `
Add-Type @"
using System;
using System.Runtime.InteropServices;
using System.Text;

public class WinAPI {
    [DllImport("user32.dll")]
    public static extern IntPtr GetForegroundWindow();

    [DllImport("user32.dll")]
    public static extern int GetWindowText(IntPtr hWnd, StringBuilder text, int count);
}
"@

$hwnd = [WinAPI]::GetForegroundWindow()
$text = New-Object System.Text.StringBuilder 256
[WinAPI]::GetWindowText($hwnd, $text, $text.Capacity) | Out-Null
$text.ToString()
`
	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command", psScript)
	output, _ := cmd.Output()
	return strings.TrimSpace(string(output))
}

// Check if blocked
func isBlocked(title string) bool {
	title = strings.ToLower(title)

	for _, safe := range safeApps {
		if strings.Contains(title, safe) {
			return false
		}
	}

	for _, app := range config.BlockedApps {
		if strings.Contains(title, app) {
			return true
		}
	}
	return false
}

// Map title → process
func getProcessFromTitle(title string) string {
	title = strings.ToLower(title)

	if strings.Contains(title, "chrome") {
		return "chrome.exe"
	}
	if strings.Contains(title, "edge") {
		return "msedge.exe"
	}

	if strings.Contains(title, "youtube") ||
		strings.Contains(title, "instagram") ||
		strings.Contains(title, "facebook") ||
		strings.Contains(title, "netflix") {
		return "browser"
	}

	return ""
}

// Kill app
func killApp(process string) {
	if process == "" {
		return
	}

	if process == "browser" {
		exec.Command("taskkill", "/IM", "chrome.exe", "/F").Run()
		exec.Command("taskkill", "/IM", "msedge.exe", "/F").Run()
		return
	}

	exec.Command("taskkill", "/IM", process, "/F").Run()
}

// Focus loop
func startFocusLoop() {
	startTime = time.Now()
	fmt.Println("🎯 Focus Mode Started")

	for focusMode {
		title := getActiveWindowTitle()

		if title != "" {
			fmt.Println("🖥️ Active:", title)

			if isBlocked(title) {
				fmt.Print("\a")
				fmt.Println("⚠️ Distraction detected! Closing in 2 seconds...")
				time.Sleep(2 * time.Second)

				process := getProcessFromTitle(title)
				killApp(process)

				distractionCount++
				fmt.Println("❌ Closed:", process)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// Stats
func showStats() {
	if startTime.IsZero() {
		fmt.Println("No session yet.")
		return
	}

	duration := time.Since(startTime).Minutes()

	fmt.Println("\n📊 STATS")
	fmt.Println("-------------------")
	fmt.Printf("Focus Time: %.2f minutes\n", duration)
	fmt.Println("Distractions Blocked:", distractionCount)
}
