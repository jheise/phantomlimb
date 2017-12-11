package main

import (
	"fmt"
	// local
	"github.com/jheise/phantomlimb/phantommsg"

	// external
	"github.com/jheise/phantomjs"
)

func takeScreenShot(request *phantommsg.PhantomRequest, page *phantomjs.WebPage) (*phantommsg.PhantomResponse, error) {
	response := new(phantommsg.PhantomResponse)

	// page, err := phantomjs.DefaultProcess.CreateWebPage()
	// if err != nil {
	// 	fmt.Println("Screenshot: failed to create page")
	// 	response.ResponseCode = 500
	// 	response.Data = ""
	// 	return response, err
	// }
	defer page.Close()

	target := request.Arguments[0]

	// if err := page.Open(target); err != nil {
	// 	fmt.Println("Screenshot: failed to open page")
	// 	response.ResponseCode = 500
	// 	response.Data = ""
	// 	return response, err
	// }

	if err := page.SetViewportSize(1024, 800); err != nil {
		fmt.Println("Screenshot: failed to viewport")
		response.ResponseCode = 500
		response.Data = ""
		return response, err
	}

	if err := page.Render(fmt.Sprintf("%s.png", target), "png", 100); err != nil {
		response.ResponseCode = 500
		response.Data = ""
		return response, err
	}

	return response, nil

}
