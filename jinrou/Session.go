package jinrou

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ISession interface {
	Next() ISession
	String() string
	Act(actor string, target string)
	Done()
}

type Noon struct {
	d *Jinrou
}

func (n *Noon) Next() ISession {
	return &Evening{d: n.d, vote: map[string]int{}}
}

func (n Noon) String() string {
	return "Noon"
}

func (n Noon) Act(actor string, target string) {

}

func (n Noon) Done() {

}

type Evening struct {
	d    *Jinrou
	vote map[string]int
	mtx  sync.Mutex
}

func (e *Evening) Next() ISession {
	return &Night{d: e.d}
}

func (e *Evening) String() string {
	return "Evening"
}

func (e *Evening) Act(actor string, target string) {
	go func() {
		e.mtx.Lock()
		defer e.mtx.Unlock()
		_, ok := e.vote[target]
		if ok {
			e.vote[target]++
		} else {
			e.vote[target] = 1
		}
	}()
}

func (e *Evening) Done() {
	count := 0
	for _, v := range e.vote {
		count += v
	}
	rand.Seed(time.Now().UnixNano())
	l := len(e.d.Players)
	for i := 0; i < l-count; i++ {
		e.Act("", e.d.Players[rand.Intn(l)].name)
	}
	max := 0
	maxName := ""
	fmt.Println(e.vote)
	for k, v := range e.vote {
		if max < v {
			max = v
			maxName = k
		}
	}
	e.d.Execute(maxName)
}

type Night struct {
	d *Jinrou
}

func (n *Night) Next() ISession {
	return &Morning{d: n.d}
}

func (n *Night) String() string {
	return "Night"
}

func (n *Night) Act(actor string, target string) {
	n.d.actionPlayer(actor, target)
}

func (n *Night) Done() {
	n.d.HandleCommand()
}

type Morning struct {
	d *Jinrou
}

func (n *Morning) Next() ISession {
	return &Noon{d: n.d}
}

func (n Morning) String() string {
	return "Morning"
}

func (n Morning) Act(actor string, target string) {

}

func (n Morning) Done() {

}
