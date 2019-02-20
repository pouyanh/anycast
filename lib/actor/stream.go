package actor

type Stream interface {
	Send(b []byte) error
	Receive() ([]byte, error)
}
