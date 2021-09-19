package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

const (
	script1 = `tell application "System Events" to tell process "Rectangle"
	click menu bar item 1 of menu bar 2
end tell
`
	script2Template = `do shell script "killall System\\ Events"
tell application "System Events" to tell process "Rectangle"
	tell menu bar item 1 of menu bar 2
		click menu item "%s" of menu 1
	end tell
end tell
`
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	option := os.Args[1]
	script2 := fmt.Sprintf(script2Template, option)
	cmd1 := exec.Command("osascript", "-e", script1)
	cmd2 := exec.Command("osascript", "-e", script2)
	go func() {
		_ = cmd1.Run()
		wg.Done()
	}()
	go func() {
		time.Sleep(250 * time.Millisecond)
		_ = cmd2.Run()
		wg.Done()
	}()
	wg.Wait()
}
