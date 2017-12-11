package main

import (
	// standard
	"fmt"
	"net/http"

	// local
	"github.com/jheise/phantomlimb/phantommsg"

	// external
	"github.com/gorilla/mux"
)

func ScreenShotHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	website := vars["website"]
	info.Printf("Screenshot request for %s\n", website)

	shotrequest := phantommsg.NewPhantomRequest("screenshot",[]string{website})
	requests <- shotrequest

	info.Printf("Waiting for response\n")
	response := <- shotrequest.Response
	info.Printf("Response received\n")

	fmt.Fprintf(w, string(response.ResponseCode))

}
