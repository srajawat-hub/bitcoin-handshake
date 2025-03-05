package handshake

import (
	"io"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

// HandshakeFunc represents a function to perform handshake with a peer.
type HandshakeFunc func(Peer) error

func Handshake(peer Peer) error {
	versionMsg := VersionMessage{
		Version:   Version,
		Services:  SrvNodeNetwork,
		Timestamp: time.Now().UTC().Unix(),
		AddrRecv: NetAddr{
			Services: SrvNodeNetwork,
			IP:       NewIPv4(127, 0, 0, 1),
			Port:     8333,
		},
		AddrFrom: NetAddr{
			Services: SrvNodeNetwork,
			IP:       NewIPv4(127, 0, 0, 1),
			Port:     3000,
		},
		Nonce:       nonce(),
		UserAgent:   NewUserAgent(),
		StartHeight: -1,
		Relay:       true,
	}

	btcMessage, err := NewBTCMessage("version", versionMsg)
	if err != nil {
		return err
	}

	msgSerialized, err := btcMessage.Serialize()
	if err != nil {
		logrus.Fatalln(err)
	}

	conn, err := net.Dial("tcp", NodeAddress)
	if err != nil {
		logrus.Fatalln(err)
	}

	_, err = conn.Write(msgSerialized)
	if err != nil {
		logrus.Fatalln(err)
	}

	tmp := make([]byte, 256)

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				logrus.Fatalln(err)
			}
			return err
		}

		logrus.Infof("received response: %x", tmp[:n])
	}
}
