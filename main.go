package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/atotto/clipboard"
	hook "github.com/robotn/gohook"
)

const (
	maxHistorySize = 100
	pollInterval   = 500 * time.Millisecond
)

var (
	clipboardHistory []string
	lastClipboard    string
)

func main() {
	fmt.Println("Clipboard Manager is running...")

	// Handle graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go monitorClipboard()
	go registerHotkeys()

	<-stopChan
	fmt.Println("\nShutting down clipboard manager...")
}

func monitorClipboard() {
	for {
		current, err := clipboard.ReadAll()
		if err != nil {
			fmt.Printf("Error reading clipboard: %v\n", err)
			continue
		}

		if current != lastClipboard && current != "" {
			// Add new entry to beginning of history
			clipboardHistory = append([]string{current}, clipboardHistory...)

			// Trim history if it exceeds maximum size
			if len(clipboardHistory) > maxHistorySize {
				clipboardHistory = clipboardHistory[:maxHistorySize]
			}

			lastClipboard = current
		}

		time.Sleep(pollInterval)
		
	}
}

func registerHotkeys() {
	fmt.Println("--- Press Shift+alt+1-9 to paste from history ---")

	for i := 1; i <= 9; i++ {
		key := fmt.Sprintf("%d", i)
		hook.Register(hook.KeyDown, []string{key, "shift", "alt"}, func(e hook.Event) {
			keyStr := string(e.Rawcode)
			pasteFromHistory(keyStr)
		})
	}

	s := hook.Start()
	<-hook.Process(s)
}

func pasteFromHistory(keyStr string) {
	if len(clipboardHistory) == 0 {
		fmt.Println("Clipboard history is empty")
		return
	}

	if keyStr >= "1" && keyStr <= "9" {
		index := int(keyStr[0] - '1')
		if index < len(clipboardHistory) {
			err := clipboard.WriteAll(clipboardHistory[index])
			if err != nil {
				fmt.Printf("Error writing to clipboard: %v\n", err)
				return
			}
			lastClipboard = clipboardHistory[index]
			fmt.Printf("Pasted item %d from history\n", index+1)
		}
	}
}
