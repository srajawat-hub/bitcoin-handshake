package handshake

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

// IPv4 represents an IPv4 address.
type IPv4 [4]byte

// NetAddr represents a network address.
type NetAddr struct {
	Time     uint32
	Services uint64
	IP       *IPv4
	Port     uint16
}

// VersionMessage represents a version message.
type VersionMessage struct {
	Version     int32
	Services    uint64
	Timestamp   int64
	AddrRecv    NetAddr
	AddrFrom    NetAddr
	Nonce       uint64
	UserAgent   VarStr
	StartHeight int32
	Relay       bool
}

// Checksum represents a checksum
type Checksum interface {
	CalculateSHA256() []byte
}

// CalculateSHA256 calculates the SHA-256 checksum of the given payload.
func CalculateSHA256(payload []byte) [32]byte {
	h1 := sha256.New()
	h1.Write(payload)
	intermediateHash := h1.Sum(nil)

	h2 := sha256.New()
	h2.Write(intermediateHash)
	finalHash := h2.Sum(nil)

	var result [32]byte
	copy(result[:], finalHash[:32])
	return result
}

func NewIPv4(a, b, c, d uint8) *IPv4 {
	return &IPv4{a, b, c, d}
}

// Serialize serializes a network address.
func (na NetAddr) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if na.Time != 0 {
		if err := binary.Write(&buf, binary.LittleEndian, na.Time); err != nil {
			return nil, err
		}
	}

	if err := binary.Write(&buf, binary.LittleEndian, na.Services); err != nil {
		return nil, err
	}

	if _, err := buf.Write(na.IP.ToIPv6()); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.BigEndian, na.Port); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToIPv6 converts IPv4 address to IPv6 format.
func (ip IPv4) ToIPv6() []byte {
	return append([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF}, ip[:]...)
}

// Serialize serializes a version message.
func (v VersionMessage) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, v.Version); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.Services); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.Timestamp); err != nil {
		return nil, err
	}

	serializedAddrRecv, err := v.AddrRecv.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(serializedAddrRecv); err != nil {
		return nil, err
	}

	serializedAddrFrom, err := v.AddrFrom.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(serializedAddrFrom); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.Nonce); err != nil {
		return nil, err
	}

	serializedUserAgent, err := v.UserAgent.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(serializedUserAgent); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.StartHeight); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, v.Relay); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
