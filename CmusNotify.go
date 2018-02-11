package main

import (
	_ "fmt"
	"os/exec"
	"strings"
	"fmt"
	"os"
)

func displayClipboard(data []string) {
	displayData := strings.Join(data, "\n")
	fmt.Println(displayData)
	exec.Command("notify-send", "-i", "/", "CmusNotify", displayData).Run()
}

func main() {
	status := os.Args[2]
	artist := os.Args[6]
	album := os.Args[8]
	track := os.Args[12]

	if status == "playing" {
		var cleData = make([]string, 3, 3)
		cleData[0] = "<b>Artist:</b> " + artist
		cleData[1] = "<b>Title:</b> " + track
		cleData[2] = "<b>Album:</b> " + album

		displayClipboard(cleData)
	}
}
