package jinrou

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	// Conn is a websocket connection of Client
	conn *websocket.Conn
	// Dead is a channel when
	dead *chan bool
}

func NewConnection(conn *websocket.Conn, dead *chan bool) Connection {
	return Connection{
		conn,
		dead,
	}
}
