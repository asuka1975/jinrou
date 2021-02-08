package jinrou

import (
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
	connection := NewConnection(conn, &dead)
	s.mu.Lock()
	s.Matcher.EnQueue(&connection)
	s.mu.Unlock()
}

func (s *MatchingServer) Start() {
	defer func() {
		log.Println("matching server end")
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
