package jinrou

import (
	"math/rand"
	"sync/atomic"
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
	d         *Jinrou
	vote      map[string]int
	wait      chan string
	waitCount int32
}

func (e *Evening) Next() ISession {
	return &Night{d: e.d}
}

func (e *Evening) String() string {
	return "Evening"
}

func (e *Evening) Act(actor string, target string) {
	if e.wait == nil {
		e.wait = make(chan string, len(e.d.Players))
	}
	go func() {
		e.wait <- target
	}()
	c := &(e.waitCount)
	atomic.AddInt32(c, 1)
}

func (e *Evening) Done() {
	for i := 0; i < int(e.waitCount); i++ {
		target := <-e.wait
		_, ok := e.vote[target]
		if ok {
			e.vote[target]++
		} else {
			e.vote[target] = 1
		}
	}
	max := 0
	for _, v := range e.vote {
		if max < v {
			max = v
		}
	}
	var names []string
	for k, v := range e.vote {
		if max == v {
			names = append(names, k)
		}
	}
	if len(names) == 0 {
		e.d.Execute(e.d.Players[rand.Intn(len(e.d.Players))].name)
	} else {
		e.d.Execute(names[rand.Intn(len(names))])
	}
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
