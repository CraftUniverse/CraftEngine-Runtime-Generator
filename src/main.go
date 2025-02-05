package main

import (
	"flag"
	"log"
)

var InputLanguage string
var SrcRoot string
var OutputPackage string
var ServerMode bool
var ServerPort int

func main() {
	flag.StringVar(&InputLanguage, "input", "", "The language to convert from (server = false)")
	flag.StringVar(&SrcRoot, "src", "", "The input project path (server = false)")
	flag.StringVar(&OutputPackage, "output", "", "The output package path (server = false)")
	flag.IntVar(&ServerPort, "port", 49181, "The port to start the server on (server = true)")
	flag.BoolVar(&ServerMode, "server", false, "If enabled a websocket will be started and files can be converted on the fly")

	flag.Parse()

	if ServerMode == false && (InputLanguage == "" || SrcRoot == "" || OutputPackage == "") {
		flag.PrintDefaults()
		return
	}

	if ServerMode == true {
		log.Println("Starting Server...")
	}
}
