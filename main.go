package main

import (
	"flag"

	"./server"
)

func main() {
	var conc int
	var addr string
	flag.IntVar(&conc, "concurency", 1, "Maximum simultanious requests")
	flag.StringVar(&addr, "address", ":8080", "Server's address and port")
	flag.Parse()

	server.Start(conc, addr)
}
