package jinrou

import (
	"fmt"
	"time"
)

type Group struct {
	// Clients is all Client of Group
}

func NewGroup(Clients []Client) Group {
	return Group{}
}

// Matcher is a service which divide active Clients into Group
type Matcher struct {
	// Queue is a queue of Client
	Queue []Client
	// GroupMemberNum is a number of Client in each Group
	GroupMemberNum int
	// MatchingInterval is a interval of matching
	MatchingInterval time.Duration
}

func NewMatcher() Matcher {
	return Matcher{
		make([]Client, 0),
		5,
		5 * time.Second,
	}
}

func (matcher *Matcher) Match() *[]Group {
	fmt.Println("match")
	client := NewClient()
	Clients := []Client{client}
	group := NewGroup(Clients)
	Groups := []Group{group}
	return &Groups
}

func (matcher *Matcher) EnQueue(client *Client) {
	fmt.Println("enqueue")
}
