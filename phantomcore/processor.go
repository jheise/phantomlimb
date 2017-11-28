package main

import (
	// standard
	"fmt"
	// local
	"github.com/jheise/phantomlimb/phantommsg"
	// external
	"github.com/jheise/phantomjs"
	// "github.com/benbjohnson/phantomjs"
)

type PhantomProcessor struct {
	Incoming chan *phantommsg.PhantomRequest
}

func NewPhantomProcessor(incoming chan *phantommsg.PhantomRequest) *PhantomProcessor {
	processor := new(PhantomProcessor)
	processor.Incoming = incoming

	return processor
}

func (self *PhantomProcessor) Run() {
	fmt.Println("Processor starting")
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		panic(err)
	}
	defer phantomjs.DefaultProcess.Close()

	for request := range self.Incoming {
		fmt.Printf("Processor: received %s request\n", request.PluginName)
		page, err := phantomjs.DefaultProcess.CreateWebPage()
		if err != nil {
			panic(err)
		}

		switch plugin := request.PluginName; plugin {
		case "screenshot":
			// process screenshot here
			response, err := takeScreenShot(request, page)
			if err != nil {
				panic(err)
			}
			request.Response <- response

		// allow for other hardcode tests

		default:
			// handle lookup here
		}
	}

}
