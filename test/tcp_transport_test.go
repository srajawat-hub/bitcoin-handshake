package test

import (
	"bitcoin-handshake/handshake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := handshake.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: handshake.Handshake,
		Decoder:       handshake.DefaultDecoder{},
	}
	tr := handshake.NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr, ":3000")

	assert.Nil(t, tr.ListenAndAccept())
}
