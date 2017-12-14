package main

import (
	// standard
	b64 "encoding/base64"
	"fmt"
	"net/http"

	// local
	"github.com/jheise/phantomlimb/phantommsg"

	// external
	"github.com/gorilla/mux"
)

func ScreenShotHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b64web := vars["website"]
	bwebsite, err := b64.StdEncoding.DecodeString(b64web)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	website := string(bwebsite)
	info.Printf("Screenshot request for %s\n", website)

	shotrequest := phantommsg.NewPhantomRequest("screenshot",[]string{website})
	requests <- shotrequest

	info.Printf("Waiting for response\n")
	response := <- shotrequest.Response
	info.Printf("Response received\n")

	fmt.Fprintf(w, string(response.ResponseCode))

}
