package jinrou

type IRole interface {
	GetName() string
	Behave(d *Jinrou, other *Player)
	Reveal() string
	StrongReveal() string
}

type Villager struct {
	name string
}

func newVillager(name string) *Villager {
	return &Villager{name: name}
}

func (v *Villager) GetName() string {
	return v.name
}

func (v *Villager) Behave(d *Jinrou, other *Player) {

}

func (v *Villager) Reveal() string {
	return "Not werewolf"
}

func (v *Villager) StrongReveal() string {
	return v.name
}

type Werewolf struct {
	name string
	self *Player
}

func newWerewolf(name string, self *Player) *Werewolf {
	return &Werewolf{name: name, self: self}
}

func (w *Werewolf) GetName() string {
	return w.name
}

func (w *Werewolf) Behave(d *Jinrou, other *Player) {
	d.PushCommand(command{b: kill, self: w.self, other: other, priority: 1})
}

func (w *Werewolf) Reveal() string {
	return "Werewolf"
}

func (w *Werewolf) StrongReveal() string {
	return w.name
}

type Knight struct {
	name string
	self *Player
}

func newKnight(name string, self *Player) *Knight {
	return &Knight{name: name, self: self}
}

func (k *Knight) GetName() string {
	return k.name
}

func (k *Knight) Behave(d *Jinrou, other *Player) {
	d.PushCommand(command{b: protect, self: k.self, other: other, priority: 0})
}

func (k *Knight) Reveal() string {
	return "Not werewolf"
}

func (k *Knight) StrongReveal() string {
	return k.name
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
