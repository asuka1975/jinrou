package jinrou

import "fmt"

type IGameState interface {
	Execute()
	NextState() IGameState
}

func FirstState() IGameState {
	return LoginState{}
}

type LoginState struct {

}

func (l LoginState) Execute() {
	fmt.Println("login")
}

func (l LoginState) NextState() IGameState {
	return RoleConfirmState{}
}

type RoleConfirmState struct {

}

func (r RoleConfirmState) Execute() {
	fmt.Println("Bob is Villager")
	fmt.Println("Alice is Werewolf")
}

func (r RoleConfirmState) NextState() IGameState {
	return GameState{}
}

type GameState struct {

}

func (g GameState) Execute() {

	user1 := CreateUser("Bob", "Villager")
	user2 := CreateUser("Alice", "Werewolf")
	num := 1
	user1.AddOnDied(func() {
		num++
		fmt.Println("Bob was killed")
	})
	fmt.Println(user2.Action(user1))
	fmt.Printf("num is %d\n", num)
}

func (g GameState) NextState() IGameState {
	return VotingState{}
}

type VotingState struct {

}

func (v VotingState) Execute() {
	fmt.Println("Voting")
}

func (v VotingState) NextState() IGameState {
	return FinishState{}
}

type FinishState struct {

}

func (f FinishState) Execute() {
	fmt.Println("Finish")
}

func (f FinishState) NextState() IGameState {
	return RoleConfirmState{}
}