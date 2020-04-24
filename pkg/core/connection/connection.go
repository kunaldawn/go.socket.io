package connection

import "time"

type Connection interface {
	GetMessage() (message string, err error)
	WriteMessage(message string) error
	Close()
	PingParams() (interval, timeout time.Duration)
}
