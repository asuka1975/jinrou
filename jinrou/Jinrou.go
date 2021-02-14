package jinrou

import (
	"sort"
)

type Jinrou struct {
	Players []*Player
	Session ISession
}

func NewJinrou(player []*Player) *Jinrou {
	session := NightSession{commands: CommandList{}}
	j := &Jinrou{Players: player}
	session.jinrou = j
	j.Session = &session
	return j
}

func (j *Jinrou) Execute(commands CommandList) {
	ctx := newContext(j.Players)
	sort.Sort(commands)
	commands = append(commands, newCommandQueue([]iBasicCommand{ElectCommand{}, KillCommand{}}, 0, Night, nil))
	for _, command := range commands {
		command.Execute(ctx)
	}
	for _, player := range j.Players {
		player.command = nil
	}
}
