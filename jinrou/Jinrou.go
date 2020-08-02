package jinrou

import (
	"sort"
)
import "sync"

type behavior int

const (
	kill behavior = iota
	protect
	reveal
	strongReveal
	none
)

type command struct {
	b        behavior
	self     *Player
	other    *Player
	priority int
}
type commands []command

func (c commands) Len() int {
	return len(c)
}

func (c commands) Less(i int, j int) bool {
	return c[i].priority < c[j].priority
}

func (c commands) Swap(i int, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

type Jinrou struct {
	Players  []*Player
	commands commands
	session  ISession
	mtx      sync.Mutex
}

func NewJinrou(name []string, role []string) *Jinrou {
	j := &Jinrou{Players: make([]*Player, len(name)), commands: []command{}}
	j.session = &Noon{d: j}
	for i := 0; i < len(name); i++ {
		j.Players[i] = NewPlayer(name[i], role[i])
	}
	return j
}

func (j *Jinrou) PushCommand(c command) {
	go func() {
		j.mtx.Lock()
		defer j.mtx.Unlock()
		j.commands = append(j.commands, c)
	}()
}

func (j *Jinrou) HandleCommand() {
	sort.Sort(j.commands)
	for _, v := range j.commands {
		switch v.b {
		case kill:
			v.other.State = dead
		case protect:
			for i, w := range j.commands {
				if w.b == kill && w.other != nil && v.other.GetName() == w.other.GetName() {
					j.commands[i].b = none
				}
			}
		case reveal:

		case strongReveal:

		default:

		}
	}
}

func (j *Jinrou) GetPlayersName() []string {
	names := make([]string, len(j.Players))
	for i, v := range j.Players {
		names[i] = v.name
	}
	return names
}

func (j *Jinrou) actionPlayer(actorName string, targetName string) {
	var actor, target *Player
	for _, v := range j.Players {
		if actorName == v.GetName() {
			actor = v
		}
		if targetName == v.GetName() {
			target = v
		}
	}
	if actor != nil && target != nil {
		actor.role.Behave(j, target)
	}
}

func (j *Jinrou) Execute(name string) {
	for i, v := range j.Players {
		if v.name == name {
			j.Players[i].State = dead
		}
	}
}

func (j *Jinrou) NextSession() ISession {
	j.session = j.session.Next()
	return j.session
}

func (j *Jinrou) GetSession() ISession {
	return j.session
}

func (j *Jinrou) IsEnd() bool {
	return false
}
