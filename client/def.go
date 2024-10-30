package client

import "errors"

const (
	// ClientVersion is the version of the client.
	ClientVersion int32 = 100
)

type futuHeader struct {
	HeaderFlag   [2]byte  // Packet header start flag, fixed as "FT"
	ProtoID      uint32   // Protocol ID
	ProtoFmtType uint8    // Protocol type, 0 for Protobuf, 1 for Json
	ProtoVer     uint8    // Protocol version, used for iterative compatibility, currently 0
	SerialNo     uint32   // Packet serial number, used to correspond to the request packet and return packet, and it is required to be incremented
	BodyLen      uint32   // Body length
	BodySHA1     [20]byte // SHA1 hash value of the original data of the packet body (after decryption)
	Reserved     [8]byte  // Reserved 8-byte extension
}

type response struct {
	ProtoID  uint32
	SerialNo uint32
	Body     []byte
}

var (
	ErrChannelClosed = errors.New("channel is closed")
	ErrInterrupted   = errors.New("process is interrupted")
	ErrTimeout       = errors.New("timeout")
)
