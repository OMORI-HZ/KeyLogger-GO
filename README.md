# KeyLogger-GO
its just educational and pls do not use it for unethical jobs

keylogger, even for educational or legitimate purposes, must be approached with caution and responsibility. Always ensure you have explicit permission and legal justification for doing so. Below, I'll guide you through creating a simple keylogger in Go, using the github.com/go-vgo/robotgo library. This example will focus on capturing key events and logging them to a file.

Prerequisites
Go Environment: Ensure Go is installed on your machine.
Install robotgo: You can install the robotgo library with:
bash

#go get github.com/go-vgo/robotgo
#go get github.com/AllenDang/w32

Imports:

robotgo: Library to capture key events.
w32: Used internally by robotgo for Windows-specific functions.
Standard libraries fmt, os, and time for formatting, file operations, and time management.
Open File for Logging:

os.OpenFile("keylog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644): Opens or creates a file named keylog.txt for appending logs.
Infinite Loop:

robotgo.AddEvent("k"): Captures the 'k' key event. You can change "k" to any key you want to monitor or use other functions to capture all key events.
time.Sleep(100 * time.Millisecond): Pauses the loop to reduce CPU usage.
Log Keystrokes:

logKeystroke: Function to write the keystroke and timestamp to the file.
Running the Keylogger
Build and Run:

#go build -o keylogger keylogger.go./keylogger


also dont forget to star&follow 👾

🤍🩶💛🤎💜🩵💙💚🧡🩷❤️🖤
