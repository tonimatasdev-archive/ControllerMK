package main

import (
	"github.com/TonimatasDEV/controller/src"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	config, _ := src.LoadConfig()
	startMilli := time.Now().UnixMilli()
	port := strconv.Itoa(config.Port)
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalln("Error starting the listener:", err)
	}

	log.Println("Listener opened successfully in the port " + port + ".")

	go src.AcceptLoop(listener, config)

	msToStart := time.Now().UnixMilli() - startMilli
	msToStartStr := strconv.Itoa(int(msToStart))
	log.Println("Done! (" + msToStartStr + "ms)")

	src.ConsoleLoop()
}
