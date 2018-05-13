// Author: Sebastian Janko
// Mail: janko.sebastian@gmail.com
// Github link: https://github.com/sebojanko/CmusNotify
// License: GNU GPL v3

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func displayNotification(data []string) {
	displayData := strings.Join(data, "\n")
	exec.Command("notify-send", "-i", "/", "CmusNotify", displayData).Run()
}

func main() {
	var argsMap = make(map[string]interface{})

	for i := 1; i < len(os.Args); i += 2 {
		fmt.Println(os.Args[i])
		fmt.Println(os.Args[i+1])
		argsMap[os.Args[i]] = os.Args[i+1]
		fmt.Println(argsMap)
	}

	if argsMap["status"] == "playing" {
		var notificationData = make([]string, 3, 3)

		notificationData[0] = "<b>Artist:</b> " + argsMap["artist"].(string)
		notificationData[1] = "<b>Title:</b> " + argsMap["title"].(string)
		notificationData[2] = "<b>Album:</b> " + argsMap["album"].(string)

		displayNotification(notificationData)
	}
}
