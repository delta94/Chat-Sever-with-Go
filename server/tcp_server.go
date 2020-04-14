package server

import (
	"chat.server.com/protocol"
	"io"
	"log"
	"net"
	"sync"
)

type TcpChatServer struct {
	listener net.Listener	// 연결을 수신하기 위해 서버를 대기시키 위한 필드
	clients []*client		// 연결되어있는 모든 클라이언트에 대한 정보를 담기 위한 필드
	mutex *sync.Mutex		// race condition을 다루기 위한 mutex 필드
}

type client struct {
	conn net.Conn						// 클라이언트와 연결에 대한 정보를 담아두기 위한 필드
	name string							// 클라이언트의 사용 이름을 저장시키기 위한 필드
	writer *protocol.CommandWriter		// 클라이언트 측에 명령을 보내기 위한 필드
}

func (ts *TcpChatServer) Listen(address string) (err error) {
	var l net.Listener
	// net.Listen 함수를 이용하여 해당 주소에 서버를 대기시키기 위한 객체를 얻을 수 있다.
	if l, err = net.Listen("tcp", address); err != nil {
		return
	}

	log.Printf("Listening on %v", address)
	ts.listener = l
	return
}

func (ts *TcpChatServer) StartServer() {
	for {
		conn, err := ts.listener.Accept()
		if err != nil {
			log.Printf("Unable to Accept err: %v\n", err)
			return
		}

		client := ts.accept(conn)
		go ts.serve(client)
	}
}

func (ts *TcpChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting new connection from %v... (current clients: %v)", conn.RemoteAddr().String(), len(ts.clients) + 1)

	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}

	ts.clients = append(ts.clients, client)
	return client
}