package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, " tcp port for blockchain server")
	flag.Parse()

	fmt.Println(*port)
	// 5000

	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
