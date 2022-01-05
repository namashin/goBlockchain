package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port: port}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.port)), nil))
}
