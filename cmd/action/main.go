package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	scriptTemplate = `tell application "System Events" to tell process "Rectangle"
	ignoring application responses
		click menu bar item 1 of menu bar 2
	end ignoring
end tell

delay 0.5
do shell script "killall System\\ Events"
tell application "System Events" to tell process "Rectangle"
	tell menu bar item 1 of menu bar 2
		click menu item "%s" of menu 1
	end tell
end tell
`
)

func main() {
	option := os.Args[1]
	script := fmt.Sprintf(scriptTemplate, option)
	cmd := exec.Command("osascript", "-e", script)
	_ = cmd.Run()
}
