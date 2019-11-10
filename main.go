package main

import (
	"flag"
	"log"
)

var (
	address = flag.String("address", "0.0.0.0", "")
	port    = flag.String("port", "8080", "")

	configuration ServerConfiguration
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	flag.Parse()

	configuration = Parse("config.json")

	Serve(*address, *port)
}
