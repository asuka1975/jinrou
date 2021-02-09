package jinrou

type Action int

const (
	None Action = iota
	Kill
	Protect
	Predict
)

var priorities = map[Action]int{
	None:    0,
	Kill:    2,
	Protect: 1,
	Predict: 1,
}

type IActiveCommand interface {
	Execute()
	GetAction() Action
	GetPriority() int
	GetSelf() *Player
	GetOther() *Player
}

type activeCommandImpl struct {
	self     *Player
	other    *Player
	action   Action
	priority int
}

func (c activeCommandImpl) GetAction() Action {
	return c.action
}

func (c activeCommandImpl) GetPriority() int {
	return c.priority
}

func (c activeCommandImpl) GetSelf() *Player {
	return c.self
}

func (c activeCommandImpl) GetOther() *Player {
	return c.other
}

type NoneCommand struct {
	activeCommandImpl
}

func (c NoneCommand) Execute() {

}

type KillCommand struct {
	activeCommandImpl
}

func (c KillCommand) Execute() {
	c.other.Status = dead
}

type ProtectCommand struct {
	activeCommandImpl
}

func (c ProtectCommand) Execute() {
	c.other.command = &PassiveCommand{
		Cancel: func(command IActiveCommand) bool {
			return command.GetAction() == Kill
		},
		Execute: func(self *Player, other *Player) {

		},
	}
}

func NewActiveCommand(self *Player, other *Player) IActiveCommand {
	action := self.role.GetAction()
	command := activeCommandImpl{self: self, other: other, action: action, priority: priorities[action]}
	switch action {
	case Kill:
		return KillCommand{command}
	case Protect:
		return ProtectCommand{command}
	default:
		return NoneCommand{command}
	}
}

type CommandList []IActiveCommand

func (c CommandList) Len() int {
	return len(c)
}

func (c CommandList) Less(i int, j int) bool {
	return c[i].GetPriority() < c[j].GetPriority()
}

func (c CommandList) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

type PassiveCommand struct {
	Cancel  func(command IActiveCommand) bool
	Execute func(self *Player, other *Player)
}
