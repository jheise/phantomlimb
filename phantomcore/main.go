package main

import (
	// standard
	"fmt"
	"os"

	// local
	// "github.com/jheise/phantomlimb/phantommsg"
	// external
	"github.com/jheise/phantomjs"
)

func main() {
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer phantomjs.DefaultProcess.Close()

	// requests := make(chan *phantommsg.PhantomRequest)

	// fmt.Println("Starting processor")
	// processor := NewPhantomProcessor(requests)
	// go processor.Run()

	// fmt.Println("generating request")
	// screenshot := phantommsg.NewPhantomRequest("screenshot", []string{"www.reddit.com"})
	// requests <- screenshot

	// fmt.Println("Waiting for response")
	// response := <-screenshot.Response
	// fmt.Println(response.ResponseCode)

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
