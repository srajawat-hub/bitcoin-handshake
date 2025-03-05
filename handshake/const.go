package handshake

import "math/rand"

const (
	UserAgent      = "/Satoshi:5.64/tb:0.0.1/"
	Network        = "mainnet"
	NodeAddress    = "127.0.0.1:8333"
	Version        = 70015
	SrvNodeNetwork = 1
)

// VarStr
type VarStr struct {
	Length uint8
	String string
}

func newVarStr(str string) VarStr {
	return VarStr{
		Length: uint8(len(str)),
		String: str,
	}
}

func NewUserAgent() VarStr {
	return newVarStr(UserAgent)
}

func nonce() uint64 {
	return rand.Uint64()
}
