package jinrou

type IGameState interface {
	Execute()
	NextState() IGameState
}

type LoginState struct {

}

func (l LoginState) Execute() {

}

func (l LoginState) NextState() IGameState {

}

type RoleConfirmState struct {

}

func (r RoleConfirmState) Execute() {

}

func (r RoleConfirmState) NextState() IGameState {

}

type GameState struct {

}

func (g GameState) Execute() {

}

func (g GameState) NextState() IGameState {

}

type VotingState struct {

}

func (v VotingState) Execute() {

}

func (v VotingState) NextState() IGameState {

}

type FinishState struct {

}

func (f FinishState) Execute() {

}

func (f FinishState) NextState() IGameState {

}