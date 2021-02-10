package jinrou

import "sort"

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
		self := command.GetSelf()
		other := command.GetOther()
		if other.command != nil {
			if !other.command.Cancel(command) {
				command.Execute()
			}
			other.command.Execute(other, self)
		} else {
			command.Execute()
		}
	}
	for _, player := range j.Players {
		player.command = nil
	}
}
