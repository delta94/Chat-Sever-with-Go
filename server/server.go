package server

// 서버의 동작을 명확하게 정의하기 위해 정의한 인터페이스
type ChatServer interface {
	// Listen 메서드는 외부에서 들어오는 연결을 수신한다.
	Listen(address string) error
	// BroadCast 메서드는 다른 모든 클라이언트에게 명령을 보낸다.
	BroadCast(command interface{}) error
	// StartServer와 CloseServer 메서드는 서버를 시작하고 종료시킨다.
	StartServer()
	CloseServer()
}