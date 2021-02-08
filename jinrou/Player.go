package jinrou

type playerState int

const (
	dead playerState = iota
	alive
)

type Player struct {
	name  string
	role  IRole
	State playerState
	Conn  *Connection
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetRole() IRole {
	return p.role
}

func NewPlayer(name string, role string) *Player {
	p := &Player{name: name, State: alive}
	p.role = newRole(role, p)

	return p
}

func newPlayer(name string, role string, connection *Connection) *Player {
	p := &Player{name: name, State: alive, Conn: connection}
	p.role = newRole(role, p)

	return p
}
