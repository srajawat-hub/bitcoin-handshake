package main

import (
	"log"
	"time"

	"bitcoin-handshake/handshake"
)

func createServer(listenAddr string, nodes ...string) *Server {
	tcptransportOpts := handshake.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: handshake.Handshake,
	}
	tcpTransport := handshake.NewTCPTransport(tcptransportOpts)

	fileServerOpts := ServerOpts{
		Transport:      tcpTransport,
		BootstrapNodes: nodes,
	}

	server := NewServer(fileServerOpts)

	return server
}

func main() {
	server := createServer(":3000", ":8333")

	go func() { log.Fatal(server.Start()) }()
	time.Sleep(500 * time.Millisecond)
}
