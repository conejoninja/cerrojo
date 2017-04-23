package transport

type Transport interface {
	Write([]byte)
	Read() ([]byte, uint16, int, error)
	Close()
}

const (
	TimeoutError = iota + 999
	ProtocolError
	EndpointError
	DisconnectedError
)