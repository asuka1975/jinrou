package jinrou

type IRole interface {
	GetAction() Action
	GetName() string
}

type Villager struct {
	roleName string
}

func (v *Villager) GetAction() Action {
	return None
}

func (v *Villager) GetName() string {
	return "Villager"
}

func newVillager(name string) *Villager {
	return &Villager{roleName: name}
}

type Werewolf struct {
	roleName string
	self     *Player
}

func (w Werewolf) GetAction() Action {
	return Kill
}

func (w Werewolf) GetName() string {
	return "Werewolf"
}

func newWerewolf(name string, self *Player) *Werewolf {
	return &Werewolf{roleName: name, self: self}
}

type Knight struct {
	name string
	self *Player
}

func (k Knight) GetAction() Action {
	return Protect
}

func (k Knight) GetName() string {
	return "Knight"
}

func newKnight(name string, self *Player) *Knight {
	return &Knight{name: name, self: self}
}

func newRole(name string, self *Player) IRole {
	switch name {
	case "Werewolf":
		return newWerewolf(name, self)
	case "Knight":
		return newKnight(name, self)
	default:
		return newVillager(name)
	}
}
