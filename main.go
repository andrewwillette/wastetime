package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/andrewwillette/gofzf"
)

var websites = map[string]string{
	"instagram": "https://www.instagram.com",
	"reddit":    "https://www.reddit.com",
	"twitter":   "https://www.twitter.com",
	"espn":      "https://www.espn.com",
}

func main() {
	var fzfInput []string
	for k := range websites {
		fzfInput = append(fzfInput, k)
	}
	selected, err := gofzf.Select(fzfInput)
	if err != nil {
		panic(err)
	}

	openbrowser(websites[selected])
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
