package jinrou

type Action int

const (
	None Action = iota
	Kill
	Protect
	Predict
	Steal
	Trance
)

var priorities = map[Action]int{
	None:    0,
	Kill:    2,
	Protect: 1,
	Predict: 1,
	Steal:   3,
	Trance:  1,
}

type IBasicCommand interface {
	Execute()
	GetSelf() *Player
	GetOther() *Player
}

type commandImpl struct {
	self  *Player
	other *Player
}

func (c *commandImpl) GetSelf() *Player {
	return c.self
}

func (c *commandImpl) GetOther() *Player {
	return c.other
}

type NoneCommand struct {
	commandImpl
}

type KillCommand struct {
	commandImpl
}

type ReviveCommand struct {
	commandImpl
}

type KnowCommand struct {
	commandImpl
}

type InformCommand struct {
	commandImpl
}

type SetRoleCommand struct {
	commandImpl
	role IRole
}

type SetPassiveCommand struct {
	commandImpl
	command *PassiveCommand
}

func (c NoneCommand) Execute() {}

func (c KillCommand) Execute() {
	c.other.Status = dead
}

func (c ReviveCommand) Execute() {
	c.other.Status = alive
}

func (c KnowCommand) Execute() {
	c.self.knowledge.Emplace(c.other)
}

func (c InformCommand) Execute() {
	c.other.knowledge.Emplace(c.self)
}

func (c SetRoleCommand) Execute() {
	c.self.role = c.role
}

func (c SetPassiveCommand) Execute() {
	c.other.command = c.command
}

type CommandQueue struct {
	commands      []IBasicCommand
	priority      int
	enableSession SessionID
	tag           interface{}
}

type CommandCreator func(*Player, *Player) CommandQueue

func (c CommandQueue) Execute() {
	for _, command := range c.commands {
		command.Execute()
	}
}

func (c CommandQueue) GetPriority() int {
	return c.priority
}

func (c CommandQueue) GetTag() interface{} {
	return c.tag
}

type CommandList []CommandQueue

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
	Cancel  func(command CommandQueue) bool
	Command CommandQueue
}
