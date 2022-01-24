package adapter

type WSClient interface {
	SubscribeExecutions() error
	ReadMessage() ([]byte, error)
	Close()
}
