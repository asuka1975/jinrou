package jinrou

type IRole interface {
	GetAction() Action
	GetName() string
	FilterTarget(player *Player) bool
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

func (v *Villager) FilterTarget(_ *Player) bool {
	return false
}

type Werewolf struct {
	name string
	self *Player
}

func (w Werewolf) GetAction() Action {
	return Kill
}

func (w Werewolf) GetName() string {
	return "Werewolf"
}

func (w Werewolf) FilterTarget(player *Player) bool {
	return player.role.GetName() != "Werewolf" && player.Status == alive
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

func (k Knight) FilterTarget(player *Player) bool {
	return player.Status == alive
}

type Diviner struct {
	name string
	self *Player
}

func (d Diviner) GetAction() Action {
	return Predict
}

func (d Diviner) GetName() string {
	return "Diviner"
}

func (d Diviner) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func newRole(name string, self *Player) IRole {
	switch name {
	case "Werewolf":
		return &Werewolf{name: name, self: self}
	case "Knight":
		return &Knight{name: name, self: self}
	case "Diviner":
		return &Diviner{self: self, name: name}
	default:
		return &Villager{roleName: name}
	}
}
