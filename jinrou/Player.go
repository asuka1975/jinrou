package jinrou

type playerStatus int

const (
	dead playerStatus = iota
	alive
)

type Player struct {
	name    string
	role    IRole
	Status  playerStatus
	command *PassiveCommand
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetRole() IRole {
	return p.role
}

func NewPlayer(name string, role string) *Player {
	p := &Player{name: name, Status: alive, command: nil}
	p.role = newRole(role, p)
	return p
}
