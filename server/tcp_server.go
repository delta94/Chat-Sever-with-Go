package server

import (
	"chat.server.com/protocol"
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
