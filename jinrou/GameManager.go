package jinrou

type GameManager struct {
	state IGameState
}

func newGameManager() *GameManager {
	return &GameManager{ state: (IGameState)(LoginState{})}
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

