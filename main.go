package main

import (
	"fmt"
	"jinrou/jinrou"
)

type J jinrou.Jinrou

func (j J) String() string {
	s := ""
	for _, p := range j.Players {
		s += fmt.Sprintf("name:%s role:%s state:%d knowledge:%+v\n", p.GetName(), p.GetRole().GetName(), p.Status, p)
	}
	return s
}

func main() {
	fmt.Println("Game Start!")

	printNames := func(player *jinrou.Player, players []*jinrou.Player) {
		for i, p := range players {
			if player.GetRole().FilterTarget(p) && player != p {
				fmt.Printf("%d: %s\n", i, p.GetName())
			}
		}
	}

	//jinrou.NewMatchingServer()
	//http.ListenAndServe(":8080", server)

	names := []string{"John", "Alice", "Bob", "Jay", "Shawn", "Elen", "Charlse"}
	roles := []string{"Werewolf", "Knight", "Villager", "Villager", "Villager", "Diviner", "Lupin"}
	var players []*jinrou.Player
	for i := 0; i < len(names); i++ {
		players = append(players, jinrou.NewPlayer(names[i], roles[i]))
	}
	j := jinrou.NewJinrou(players)
	var commands jinrou.CommandList
	for _, player := range j.Players {
		fmt.Printf("%s, your turn. You are %s\n", player.GetName(), player.GetRole().GetName())
		printNames(player, j.Players)
		fmt.Printf("Who do you choose?: ")
		var i int
		_, _ = fmt.Scanf("%d", &i)
		if i < len(players) {
			commands = append(commands, player.GetRole().GetCommand()(player, players[i]))
		}
	}
	fmt.Printf("the number of commands: %d\n", len(commands))
	j.Execute(commands)
	fmt.Println(J(*j))
}
