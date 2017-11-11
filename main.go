package main

import (
	// standard
	"fmt"
	"os"

	// external
	"github.com/benbjohnson/phantomjs"
)

func main() {
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer phantomjs.DefaultProcess.Close()

	page, err := phantomjs.DefaultProcess.CreateWebPage()
	if err != nil {
		panic(err)
	}
	defer page.Close()

	if err := page.Open("https://www.google.com"); err != nil {
		panic(err)
	}

	if err := page.SetViewportSize(1024, 800); err != nil {
		panic(err)
	}

	if err := page.Render("google.png", "png", 100); err != nil {
		panic(err)
	}

}
