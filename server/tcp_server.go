package server

import (
	"chat.server.com/protocol"
	"log"
	"net"
	"sync"
)

type TcpChatServer struct {
	listener net.Listener
	clients []*client
	mutex *sync.Mutex
}

type client struct {
	conn net.Conn
	name string
	writer *protocol.CommandWriter
}

func (ts *TcpChatServer) Listen(address string) (err error) {
	var l net.Listener
	if l, err = net.Listen("tcp", address); err != nil {
		return
	}

	log.Printf("Listening on %v", address)
	ts.listener = l
	return
}
