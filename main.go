package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	hook "github.com/robotn/gohook"
)


var (
	history       []string
	lastClipboard string
)

func main() {
	fmt.Print("program is running")
	go monitorClipboard()
	go add()
	select {}
}

func monitorClipboard() {
	for {
		current, err := clipboard.ReadAll()
		if err == nil && current != lastClipboard {
			history = append([]string{current}, history...)
			lastClipboard = current
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func add() {
	fmt.Println("--- Please press ctrl + shift + 1-9 multiple times ---")
	hook.Register(hook.KeyDown, []string{"1", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})

	hook.Register(hook.KeyDown, []string{"2", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})
	hook.Register(hook.KeyDown, []string{"3", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})

	hook.Register(hook.KeyDown, []string{"4", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})
	hook.Register(hook.KeyDown, []string{"5", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})
	hook.Register(hook.KeyDown, []string{"6", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})
	hook.Register(hook.KeyDown, []string{"7", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})
	hook.Register(hook.KeyDown, []string{"8", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})
	hook.Register(hook.KeyDown, []string{"9", "ctrl", "shift"}, func(e hook.Event) {
		keyStr := string(e.Rawcode)
		monitorKeyboard(keyStr)
	})

	s := hook.Start()
	<-hook.Process(s)
}

func monitorKeyboard(keyStr string) {
	if len(history) == 0 {
		return
	}
	fmt.Print(history)

	if keyStr >= "0" && keyStr <= "9" {
		index := int(keyStr[0] - '1')
		if index < len(history) {
			clipboard.WriteAll(history[index])
			lastClipboard = history[index]
		}
	}
}
