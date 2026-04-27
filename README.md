# 🎯 FocusGuard — Distraction Control System

FocusGuard is a real-time CLI-based productivity tool built in Go that automatically detects and blocks distracting applications like YouTube, Instagram, and social media platforms.

---

## 🚀 Features

- 🔍 Real-time active window detection
- 🚫 Automatic blocking of distractions
- ⚡ Fast and lightweight (built in Go)
- 📊 Focus time tracking and stats
- ⚙️ Configurable blocked apps 

---

## 🧠 How It Works

FocusGuard continuously monitors the active window using Windows API.

If a distracting app is detected:
1. It alerts the user
2. Waits 2 seconds
3. Automatically closes the application

---


## 🛠️ Tech Stack

- Language: Go (Golang)
- OS Integration: Windows API (via PowerShell)
- Process Control: taskkill

---

## ▶️ How to Run

```bash
git clone https://github.com/your-username/FocusGuard.git
cd FocusGuard
go run .
