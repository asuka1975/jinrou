package jinrou

import (
	"sort"
)

type Jinrou struct {
	Players []*Player
	session ISession
}

func NewJinrou(player []*Player) *Jinrou {
	j := &Jinrou{Players: player}
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
