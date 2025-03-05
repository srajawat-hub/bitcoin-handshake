package main

import (
	"context"
	"fmt"
	"log"

	"bitcoin-handshake/handshake"
)

// ServerOpts contains options for configuring the server.
type ServerOpts struct {
	Transport      handshake.Transport // Transport interface for communication
	BootstrapNodes []string            // List of bootstrap nodes to connect to
}

// Server represents a server instance.
type Server struct {
	ServerOpts                           // Embedding ServerOpts to inherit its fields
	peers      map[string]handshake.Peer // Map of connected peers
	ctx        context.Context           // Context for graceful shutdown
	cancel     context.CancelFunc        // Function to cancel context
}

// NewServer creates a new Server instance with the provided options.
func NewServer(opts ServerOpts) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		ServerOpts: opts,
		peers:      make(map[string]handshake.Peer), // Initialize peers map
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Message represents a message received from peers.
type Message struct {
	Payload any
}

// loop is the main event loop of the server.
func (s *Server) loop() {
	defer func() {
		log.Println("file server stopped due to error or user quit action")
		s.Transport.Close()
	}()

	for range s.ctx.Done() {
		return
	}
}

// bootstrapNetwork connects to bootstrap nodes on startup.
func (s *Server) bootstrapNetwork() error {
	for _, addr := range s.BootstrapNodes {
		if len(addr) == 0 {
			continue
		}

		go func(addr string) {
			fmt.Printf("[%s] attemping to connect with remote %s\n", s.Transport.Addr(), addr)
			if err := s.Transport.Dial(addr); err != nil {
				log.Println("dial error: ", err)
			}
		}(addr)
	}

	return nil
}

// Start starts the server.
func (s *Server) Start() error {
	fmt.Printf("[%s] starting fileserver...\n", s.Transport.Addr())

	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}

	s.bootstrapNetwork()

	s.loop()

	return nil
}

// Stop signals the server to stop gracefully.
func (s *Server) Stop() {
	s.cancel()
}
