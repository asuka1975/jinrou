package jinrou

type IRole interface {
	GetCommand() CommandCreator
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

func (v Villager) GetCommand() CommandCreator {
	return func(self *Player, other *Player) CommandQueue {
		return CommandQueue{
			commands:      nil,
			priority:      0,
			enableSession: Night,
			tag:           nil,
		}
	}
}

func (v Villager) GetName() string {
	return "Villager"
}

func (v Villager) FilterTarget(_ *Player) bool {
	return false
}

func (w Werewolf) GetCommand() CommandCreator {
	return func(self *Player, other *Player) CommandQueue {
		return CommandQueue{
			commands:      []IBasicCommand{&KillCommand{commandImpl{self, other}}},
			priority:      2,
			enableSession: Night,
			tag:           Kill,
		}
	}
}

func (w Werewolf) GetName() string {
	return "Werewolf"
}

func (w Werewolf) FilterTarget(player *Player) bool {
	return player.role.GetName() != "Werewolf" && player.Status == alive
}

func (k Knight) GetCommand() CommandCreator {
	return func(self *Player, other *Player) CommandQueue {
		return CommandQueue{
			commands: []IBasicCommand{&SetPassiveCommand{
				commandImpl: commandImpl{self, other},
				command: &PassiveCommand{
					Cancel: func(command CommandQueue) bool {
						action, ok := command.tag.(Action)
						return ok && action == Kill
					},
					Execute: func(self *Player, other *Player) {},
				},
			}},
			priority:      1,
			enableSession: 0,
			tag:           nil,
		}
	}
}

func (k Knight) GetName() string {
	return "Knight"
}

func (k Knight) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func (d Diviner) GetCommand() CommandCreator {
	return func(self *Player, other *Player) CommandQueue {
		return CommandQueue{
			commands:      []IBasicCommand{&KnowCommand{commandImpl{self, other}}},
			priority:      1,
			enableSession: 0,
			tag:           nil,
		}
	}
}

func (d Diviner) GetName() string {
	return "Diviner"
}

func (d Diviner) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func (l Lupin) GetCommand() CommandCreator {
	return func(self *Player, other *Player) CommandQueue {
		return CommandQueue{
			commands: []IBasicCommand{
				&SetRoleCommand{
					commandImpl: commandImpl{self, other},
					role:        other.role,
				},
				&SetRoleCommand{
					commandImpl: commandImpl{other, self},
					role:        villager,
				},
			},
			priority:      3,
			enableSession: 0,
			tag:           nil,
		}
	}
}

func (l Lupin) GetName() string {
	return "Lupin"
}

func (l Lupin) FilterTarget(player *Player) bool {
	return player.Status == alive
}

func (s Shaman) GetCommand() CommandCreator {
	return func(self *Player, other *Player) CommandQueue {
		return CommandQueue{
			commands:      []IBasicCommand{&KnowCommand{commandImpl{self, other}}},
			priority:      1,
			enableSession: 0,
			tag:           nil,
		}
	}
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
