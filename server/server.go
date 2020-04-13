package server

// 서버의 동작을 명확하게 정의하기 위해 정의한 인터페이스
type ChatServer interface {
	Listen(address string) error
	BroadCast(command interface{}) error
	StartServer()
	Close()
}