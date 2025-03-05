package handshake

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

const (
	checksumLength = 4
	magicLength    = 4
)

// MessagePayload ...
type MessagePayload interface {
	Serialize() ([]byte, error)
}

var (
	magicMainnet = [magicLength]byte{0xf9, 0xbe, 0xb4, 0xd9}
)

// BTCMessage ...
type BTCMessage struct {
	Magic    [magicLength]byte
	Command  [commandLength]byte
	Length   uint32
	Checksum [checksumLength]byte
	Payload  []byte
}

// NewBTCMessage is used to create BTCMessage
func NewBTCMessage(cmd string, payload MessagePayload) (*BTCMessage, error) {
	serializedPayload, err := payload.Serialize()
	if err != nil {
		return nil, err
	}

	command, ok := commands[cmd]
	if !ok {
		return nil, fmt.Errorf("unsupported command %s", cmd)
	}

	msg := BTCMessage{
		Magic:    magicMainnet,
		Command:  command,
		Length:   uint32(len(serializedPayload)),
		Checksum: checksum(serializedPayload),
		Payload:  serializedPayload,
	}

	return &msg, nil
}

// Serialize is used for serializing BTCMessage
func (m BTCMessage) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if _, err := buf.Write(m.Magic[:]); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Command[:]); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, m.Length); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Checksum[:]); err != nil {
		return nil, err
	}

	if _, err := buf.Write(m.Payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v VarStr) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, v.Length); err != nil {
		return nil, err
	}

	if _, err := buf.Write([]byte(v.String)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func checksum(data []byte) [checksumLength]byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	var hashArr [checksumLength]byte
	copy(hashArr[:], hash[0:checksumLength])

	return hashArr
}
