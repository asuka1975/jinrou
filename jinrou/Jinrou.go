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
	sort.Sort(commands)
	for _, command := range commands {
		passiveCalled := false
		for _, basicCommand := range command.commands {
			other := basicCommand.GetOther()
			if !passiveCalled && other.command != nil {
				if other.command.Cancel(command) {
					other.command.Command.Execute()
					break
				}
				basicCommand.Execute()
				passiveCalled = true
			} else {
				basicCommand.Execute()
			}
		}
	}
	for _, player := range j.Players {
		player.command = nil
	}
}
