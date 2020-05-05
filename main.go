package main

import (
	"fmt"
	"jinrou/jinrou"
)

func main() {
	fmt.Println("Game Start!")
	
	manager := jinrou.NewGameManager()
	manager.Run()
}