package jinrou

import (
	"log"
	"time"
)

// Matcher is a service which divide active Clients into Group
type Matcher struct {
	// Queue is a queue of Client
	Queue []*Connection
	// GroupMemberNum is a number of Client in each Group
	GroupMemberNum int
	// MatchingInterval is a interval of matching
	MatchingInterval time.Duration
}

func NewMatcher() Matcher {
	return Matcher{
		make([]*Connection, 0),
		2,
		5 * time.Second,
	}
}

func (matcher *Matcher) Match() {
	log.Printf("%v clients in Queue\n", len(matcher.Queue))
	for len(matcher.Queue) > matcher.GroupMemberNum {
		log.Println("match")
		var players []*Player
		for j := 0; j < matcher.GroupMemberNum; j++ {
			connection := matcher.Queue[j]
			player := NewPlayer("time", "Werewolf")
			player.Connection = connection
			players = append(players, player)
			matcher.Queue = matcher.Queue[1:]
		}
	}
}

func (matcher *Matcher) EnQueue(connection *Connection) {
	log.Println("enqueue")
	matcher.Queue = append(matcher.Queue, connection)
}
