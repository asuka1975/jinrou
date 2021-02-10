package jinrou

type playerStatus int

const (
	dead playerStatus = iota
	alive
)

type Player struct {
	name       string
	role       IRole
	Status     playerStatus
	command    *PassiveCommand
	Connection *Connection
	knowledge  Knowledge
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

type KnowledgeFragment struct {
	playerName string
	status     playerStatus
	roleName   string
}

type Knowledge []KnowledgeFragment

func (k *Knowledge) IndexOf(player string) int {
	for i, frag := range *k {
		if frag.playerName == player {
			return i
		}
	}
	return -1
}

func (k *Knowledge) Emplace(player *Player) {
	*k = append(*k, KnowledgeFragment{
		playerName: player.name,
		status:     player.Status,
		roleName:   player.role.GetName(),
	})
}
