package jinrou

import "sort"

type Jinrou struct {
	Players []*Player
	session ISession
}

func NewJinrou(name []string, role []string) *Jinrou {
	j := &Jinrou{Players: make([]*Player, len(name))}
	for i := 0; i < len(name); i++ {
		j.Players[i] = NewPlayer(name[i], role[i])
	}
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
