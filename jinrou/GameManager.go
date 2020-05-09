package jinrou

type GameManager struct {
	state IGameState
	jinrou *Jinrou
}

func NewGameManager() *GameManager {
	return &GameManager{ state: FirstState()}
}

func (g *GameManager) Run() {
	for {
		g.state.Execute()
		_, ok := g.state.(FinishState)
		if ok {
			break
		}
		g.state = g.state.NextState()
	}
}

