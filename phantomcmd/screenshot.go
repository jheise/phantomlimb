package main

import (
	// standard
	"fmt"
	"os"

	// local
	// "github.com/jheise/phantomlimb/phantomapi"
)

func processScreenShot(target string) {
	fmt.Printf("Capturing %s\n", target)
	response, err := api.ScreenShot(target)
	if err != nil {
		fmt.Printf("unable to capture, %s\n", target)
		os.Exit(1)
	}

	fmt.Printf("Captured %s\n", target)
	fmt.Println(string(response[:]))
}
