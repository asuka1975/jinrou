package jinrou

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	// Conn is a websocket connection of Client
	conn *websocket.Conn
	// Dead is a channel when
	dead *chan bool
}

func newClient(conn *websocket.Conn, dead *chan bool) Client {
	return Client{
		conn,
		dead,
	}
}

func NewClient() Client {
	return Client{}
}

type MatchingServer struct {
	Matcher *Matcher
	mu      sync.Mutex
}

func (s *MatchingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
	}
	dead := make(chan bool)
	client := newClient(conn, &dead)
	s.mu.Lock()
	s.Matcher.EnQueue(&client)
	s.mu.Unlock()
}

func (s *MatchingServer) Start() {
	defer func() {
		fmt.Println("matching server end")
	}()
	ticker := time.NewTicker(s.Matcher.MatchingInterval)
	for {
		select {
		case <-ticker.C:
			s.Matcher.Match()
		}
	}
}

func NewMatchingServer() *MatchingServer {
	matcher := NewMatcher()
	server := MatchingServer{
		&matcher,
		sync.Mutex{},
	}
	go server.Start()
	return &server
}
