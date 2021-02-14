package jinrou

import (
	"log"
)

type Action int

type context struct {
	target  Stack
	station pollingStation
}

func newContext(players []*Player) *context {
	return &context{
		target:  []*Player{},
		station: newPollingStation(players),
	}
}

type iBasicCommand interface {
	execute(ctx *context)
	target(ctx *context) *Player
}

type NoneCommand struct{}

type PushCommand struct {
	player *Player
}

type PopCommand struct{}

type KillCommand struct{}

type ReviveCommand struct{}

type InformCommand struct{}

type KnowCommand struct{}

type VoteCommand struct{}

type ElectCommand struct{}

type SetRoleCommand struct {
	role IRole
}

type SetPassiveCommand struct {
	command *PassiveCommand
}

func (c NoneCommand) execute(*context)        {}
func (c NoneCommand) target(*context) *Player { return nil }

func (c PushCommand) execute(ctx *context) {
	ctx.target.Push(c.player)
}
func (c PushCommand) target(*context) *Player { return nil }

func (c PopCommand) execute(ctx *context) {
	_, _ = ctx.target.Pop()
}
func (c PopCommand) target(*context) *Player { return nil }

func (c KillCommand) execute(ctx *context) {
	p, err := ctx.target.Pop()
	if err != nil {
		log.Fatalf("invalid command call: %s\n", err.Error())
		return
	}
	p.Status = dead
}
func (c KillCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

func (c ReviveCommand) execute(ctx *context) {
	p, err := ctx.target.Pop()
	if err != nil {
		log.Fatalf("invalid command call: %s\n", err.Error())
		return
	}
	p.Status = alive
}
func (c ReviveCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

func (c InformCommand) execute(ctx *context) {
	other, err1 := ctx.target.Pop()
	if err1 != nil {
		log.Fatalf("invalid command call: %s\n", err1.Error())
		return
	}
	self, err2 := ctx.target.Pop()
	if err1 != nil {
		log.Fatalf("invalid command call: %s\n", err2.Error())
		return
	}
	other.knowledge.Emplace(self)
}
func (c InformCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

func (c KnowCommand) execute(ctx *context) {
	other, err1 := ctx.target.Pop()
	if err1 != nil {
		log.Fatalf("invalid command call: %s\n", err1.Error())
		return
	}
	self, err2 := ctx.target.Pop()
	if err1 != nil {
		log.Fatalf("invalid command call: %s\n", err2.Error())
		return
	}
	self.knowledge.Emplace(other)
}
func (c KnowCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

func (c VoteCommand) execute(ctx *context) {
	other, err := ctx.target.Pop()
	if err != nil {
		log.Fatalf("invalid command call: %s\n", err.Error())
		return
	}
	ctx.station.vote(other.name)
}
func (c VoteCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

func (c ElectCommand) execute(ctx *context) {
	ctx.target.Push(ctx.station.voted())
}
func (c ElectCommand) target(*context) *Player { return nil }

func (c SetRoleCommand) execute(ctx *context) {
	p, err := ctx.target.Pop()
	if err != nil {
		log.Fatalf("invalid command call: %s\n", err.Error())
		return
	}
	p.role = c.role
}
func (c SetRoleCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

func (c SetPassiveCommand) execute(ctx *context) {
	p, err := ctx.target.Pop()
	if err != nil {
		log.Fatalf("invalid command call: %s\n", err.Error())
		return
	}
	p.command = c.command
}
func (c SetPassiveCommand) target(ctx *context) *Player {
	p, _ := ctx.target.Top()
	return p
}

type CommandQueue struct {
	commands      []iBasicCommand
	priority      int
	enableSession SessionID
	tag           interface{}
	targets       []*Player
}

func (c CommandQueue) Execute(ctx *context) {
	for _, command := range c.commands {
		trg := command.target(ctx)
		if trg != nil && trg.command != nil {
			if !trg.command.Cancel(command) {
				command.execute(ctx)
			}
			trg.command.Command.Execute(ctx)
		} else {
			command.execute(ctx)
		}
	}
}

func (c CommandQueue) GetPriority() int {
	return c.priority
}

func (c CommandQueue) GetTag() interface{} {
	return c.tag
}

func (c CommandQueue) GetSelf() *Player {
	if len(c.targets) == 0 {
		return nil
	} else {
		return c.targets[0]
	}
}

func (c CommandQueue) GetOther() *Player {
	if len(c.targets) < 2 {
		return nil
	} else {
		return c.targets[1]
	}
}

func newCommandQueue(commands []iBasicCommand, priority int, enableSession SessionID, tag interface{}, initialTargets ...*Player) CommandQueue {
	preCommands := make([]iBasicCommand, len(initialTargets))
	for i := 0; i < len(initialTargets); i++ {
		preCommands[i] = PushCommand{initialTargets[i]}
	}
	return CommandQueue{
		commands:      append(preCommands, commands...),
		priority:      priority,
		enableSession: enableSession,
		tag:           tag,
		targets:       initialTargets,
	}
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
	Cancel  func(command iBasicCommand) bool
	Command CommandQueue
}
