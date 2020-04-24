package message

type Type int

const (
	TypeOpen        = Type(0) // Message with connection options
	TypeClose       = Type(1) // Close connection and destroy all handle routines
	TypePing        = Type(2) // Ping request message
	TypePong        = Type(3) // Pong response message
	TypeEmpty       = Type(4) // Empty message
	TypeEmit        = Type(5) // Emit request, no response
	TypeAckRequest  = Type(6) // Emit request, wait for response (ack)
	TypeAckResponse = Type(7) // ack response
)

type Message struct {
	Type   Type
	AckId  int
	Method string
	Args   string
	Source string
}
