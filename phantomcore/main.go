package main

import (
	// standard
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	// local
	"github.com/jheise/phantomlimb/phantommsg"

	// external
	// "github.com/jheise/phantomjs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// config options
	address string
	port string

	// loggers
	info *log.Logger
	elog *log.Logger

	// channels
	requests chan *phantommsg.PhantomRequest
)

func init() {
	flag.StringVar(&address, "Address", "0.0.0.0", "IP address to bind to")
	flag.StringVar(&port, "Port", "9797", "port to bind to ")
	flag.Parse()

	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {

	// create channels
	info.Println("Initializing channels")
	requests = make(chan *phantommsg.PhantomRequest)

	// create and launch scanner
	info.Println("Starting processor")
	processor := NewPhantomProcessor(requests)
	go processor.Run()

	// prep address and port
	addrport := fmt.Sprintf("%s:%s", address, port)

	// setup http server and begin serving traffic
	r := mux.NewRouter()
	r.HandleFunc("/phantom/v1/screenshot/{website}", ScreenShotHandler).Methods("GET")
	http.Handle("/", r)
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
	http.ListenAndServe(addrport, loggedRouter)

	// setup http server and begin serving traffic
	// fmt.Println("generating request")
	// screenshot := phantommsg.NewPhantomRequest("screenshot", []string{"https://www.google.com"})
	// screenshot := phantommsg.NewPhantomRequest("screenshot", []string{"https://www.blockexplorer.com"})
	// requests <- screenshot

	// fmt.Println("Waiting for response")
	// response := <-screenshot.Response
	// fmt.Println(response.ResponseCode)

}
