package test

import (
	"bitcoin-handshake/handshake"
	"encoding/hex"
	"testing"
	"time"
)

func TestMessageSerialization(t *testing.T) {
	version := handshake.VersionMessage{
		Version:   handshake.Version,
		Services:  handshake.SrvNodeNetwork,
		Timestamp: time.Date(2024, 11, 11, 0, 0, 0, 0, time.UTC).Unix(),
		AddrRecv: handshake.NetAddr{
			Services: handshake.SrvNodeNetwork,
			IP:       handshake.NewIPv4(127, 0, 0, 1),
			Port:     8333,
		},
		AddrFrom: handshake.NetAddr{
			Services: handshake.SrvNodeNetwork,
			IP:       handshake.NewIPv4(127, 0, 0, 1),
			Port:     3000,
		},
		Nonce:       31337,
		UserAgent:   handshake.NewUserAgent(),
		StartHeight: -1,
		Relay:       true,
	}
	msg, err := handshake.NewBTCMessage("version", version)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
		return
	}

	msgSerialized, err := msg.Serialize()
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
		return
	}

	actual := hex.EncodeToString(msgSerialized)
	expected := "f9beb4d976657273696f6e00000000006d0000005a97f5587f11010001000000000000000049316700000000010000000000000000000000000000000000ffff7f000001208d010000000000000000000000000000000000ffff7f0000010bb8697a000000000000172f5361746f7368693a352e36342f74623a302e302e312fffffffff01"
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}

}
