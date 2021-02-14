package jinrou

type Jinrou struct {
	Players []*Player
	session ISession
}

func NewJinrou(player []*Player) *Jinrou {
	session := NightSession{commands: CommandList{}}
	j := &Jinrou{Players: player}
	session.jinrou = j
	j.session = &session
	return j
}

func (j *Jinrou) GetSession() ISession {
	return j.session
}

func (j *Jinrou) Next() {
	j.session = j.session.Next()
}
