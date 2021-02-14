package jinrou

type IRole interface {
	GetCommand(self *Player, other *Player) CommandQueue
	GetName() string
	FilterTarget(player *Player) bool
}

type Villager int
type Werewolf int
type Knight int
type Diviner int
type Lupin int
type Shaman int

const (
	villager Villager = 0
	werewolf Werewolf = 1
	knight   Knight   = 2
	diviner  Diviner  = 3
	lupin    Lupin    = 4
	shaman   Shaman   = 5
)

func (v Villager) GetCommand(self *Player, other *Player) CommandQueue {
	return newCommandQueue(nil, 0, Night, nil)
}

func (v Villager) GetName() string {
	return "Villager"
}

func (v Villager) FilterTarget(_ *Player) bool {
	return false
}

func (w Werewolf) GetCommand(self *Player, other *Player) CommandQueue {
	return newCommandQueue(
		[]iBasicCommand{VoteCommand{}},
		1, Night, nil, self, other)
}

func (w Werewolf) GetName() string {
	return "Werewolf"
}

func (w Werewolf) FilterTarget(player *Player) bool {
	return player.role.GetName() != "Werewolf" && player.Status == alive
}

func (k Knight) GetCommand(self *Player, other *Player) CommandQueue {
	return newCommandQueue(
		[]iBasicCommand{SetPassiveCommand{
			command: &PassiveCommand{
				Cancel: func(command iBasicCommand) bool {
					switch command.(type) {
					case KillCommand:
						return true
					default:
						return false
					}
				},
				Command: newCommandQueue(
					[]iBasicCommand{NoneCommand{}},
					1, Night, nil, other),
			},
		}},
		1, Night, nil, self, other)
}

func (k Knight) GetName() string {
	return "Knight"
}

func (k Knight) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func (d Diviner) GetCommand(self *Player, other *Player) CommandQueue {
	return newCommandQueue(
		[]iBasicCommand{KnowCommand{}},
		1, Night, nil, self, other)
}

func (d Diviner) GetName() string {
	return "Diviner"
}

func (d Diviner) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func (l Lupin) GetCommand(self *Player, other *Player) CommandQueue {
	return newCommandQueue(
		[]iBasicCommand{SetRoleCommand{villager}, SetRoleCommand{other.role}},
		3, Night, nil, self, other)
}

func (l Lupin) GetName() string {
	return "Lupin"
}

func (l Lupin) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func (s Shaman) GetCommand(self *Player, other *Player) CommandQueue {
	return newCommandQueue([]iBasicCommand{InformCommand{}}, 1, Night, nil, other, self)
}

func (s Shaman) GetName() string {
	return "Shaman"
}

func (s Shaman) FilterTarget(player *Player) bool {
	return player.Status == dead
}

func newRole(name string, self *Player) IRole {
	switch name {
	case "Werewolf":
		return werewolf
	case "Knight":
		return knight
	case "Diviner":
		return diviner
	case "Lupin":
		return lupin
	case "Shaman":
		return shaman
	default:
		return villager
	}
}
