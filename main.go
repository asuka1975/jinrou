package main

import (
	"fmt"
	"jinrou/jinrou"
)

type J jinrou.Jinrou

func (j J) String() string {
	s := ""
	for _, p := range j.Players {
		s += fmt.Sprintf("name:%s role:%s state:%d\n", p.GetName(), p.GetRole().GetName(), p.Status)
	}
	return s
}

func main() {
	fmt.Println("Game Start!")

	printNames := func(names []string) {
		for i, name := range names {
			fmt.Printf("%d: %s\n", i, name)
		}
	}

	names := []string{"John", "Alice", "Bob", "Jay", "Shawn"}
	j := jinrou.NewJinrou(
		names,
		[]string{"Werewolf", "Knight", "Villager", "Villager", "Villager"})
	var commands []jinrou.IActiveCommand
	for _, player := range j.Players {
		fmt.Printf("%s, your turn.\n", player.GetName())
		switch player.GetRole().GetName() {
		case "Knight":
			printNames(names)
			fmt.Printf("Who do you protect?: ")
			var i int
			_, _ = fmt.Scanf("%d", &i)
			commands = append(commands, jinrou.NewActiveCommand(player, j.Players[i]))
		case "Werewolf":
			printNames(names)
			fmt.Printf("Who do you kill?: ")
			var i int
			_, _ = fmt.Scanf("%d", &i)
			commands = append(commands, jinrou.NewActiveCommand(player, j.Players[i]))
		case "Villager":
			fmt.Printf("Sleep right now!\n")
		}
	}
	fmt.Printf("the number of commands: %d\n", len(commands))
	j.Execute(commands)
	fmt.Println(J(*j))
}
