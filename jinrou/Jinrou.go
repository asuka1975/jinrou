package jinrou

import (
	"fmt"
	"math/rand"
	"sort"
	"sync/atomic"
	"time"
)

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
	Players      []*Player
	commands     commands
	session      ISession
	commandFence chan command
	commandCount int32
}

func NewJinrou(name []string, role []string) *Jinrou {
	j := &Jinrou{Players: make([]*Player, len(name)), commands: []command{}, commandCount: 0}
	j.session = &Noon{d: j}
	j.commandFence = make(chan command, len(j.Players))
	for i := 0; i < len(name); i++ {
		j.Players[i] = NewPlayer(name[i], role[i])
	}
	return j
}

func newJinrou(name []string, role []string, players []*Player) *Jinrou {
	j := &Jinrou{Players: make([]*Player, len(name)), commands: []command{}, commandCount: 0}
	j.session = &Noon{d: j}
	j.commandFence = make(chan command, len(j.Players))
	j.Players = players
	for i := 0; i < len(name); i++ {
		j.Players[i] = NewPlayer(name[i], role[i])
	}
	return j
}

func (j *Jinrou) PushCommand(c command) {
	go func() {
		j.commandFence <- c
	}()
	atomic.AddInt32(&(j.commandCount), 1)
}

func (j *Jinrou) HandleCommand() {
	for i := 0; i < int(j.commandCount); i++ {
		j.commands = append(j.commands, <-j.commandFence)
	}
	j.commandCount = 0
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

func (j *Jinrou) String() string {
	s := ""
	for _, p := range j.Players {
		s += fmt.Sprintf("name:%s role:%s state:%d\n", p.GetName(), p.GetRole().GetName(), p.State)
	}
	return s
}

func (j *Jinrou) Run() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		session := j.GetSession()
		fmt.Println(session.String())
		switch session.String() {
		case "Morning":

		case "Noon":

		case "Evening":
			for i := 0; i < len(j.Players); i++ {
				k := rand.Intn(len(j.Players) - 1)
				if k >= i {
					k++
				}
				session.Act(j.Players[i].GetName(), j.Players[k].GetName())
			}
		case "Night":
			for i := 0; i < len(j.Players); i++ {
				k := rand.Intn(len(j.Players) - 1)
				if k >= i {
					k++
				}
				session.Act(j.Players[i].GetName(), j.Players[k].GetName())
			}
		}
		j_ := Jinrou(*j)
		fmt.Println(j_.String())
		session.Done()
		j.NextSession()
	}
}
