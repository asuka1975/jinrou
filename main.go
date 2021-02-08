package main

import (
	"fmt"
	"jinrou/jinrou"
	"math/rand"
	"net/http"
	"time"
)

type J jinrou.Jinrou

func main() {
	fmt.Println("Game Start!")
	server := jinrou.NewMatchingServer()
	http.ListenAndServe(":8080", server)

	j := jinrou.NewJinrou(
		[]string{"John", "Alice", "Bob", "Jay", "Shawn"},
		[]string{"Knight", "Werewolf", "Villager", "Villager", "Villager"})
	rand.Seed(time.Now().UnixNano())
	j.Run()
}
