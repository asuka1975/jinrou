package main

import (
	"fmt"
	"jinrou/jinrou"
	"math/rand"
	"time"
)

type J jinrou.Jinrou

func (j *J) String() string {
	s := ""
	for _, p := range j.Players {
		s += fmt.Sprintf("name:%s role:%s state:%d\n", p.GetName(), p.GetRole().GetName(), p.State)
	}
	return s
}

func main() {
	fmt.Println("Game Start!")

	j := jinrou.NewJinrou(
		[]string{"John", "Alice", "Bob", "Jay", "Shohn"},
		[]string{"Werewolf", "Knight", "Villager", "Villager", "Villager"})
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		session := j.GetSession()
		fmt.Println(session.String())
		switch session.String() {
		case "Morning":

		case "Noon":

		case "Evening":
			for i := 0; i < len(j.Players); i++ {
				k := rand.Intn(len(j.Players) - 1)
				if k >= i {
					k++
				}
				session.Act(j.Players[i].GetName(), j.Players[k].GetName())
			}
		case "Night":
			for i := 0; i < len(j.Players); i++ {
				k := rand.Intn(len(j.Players) - 1)
				if k >= i {
					k++
				}
				session.Act(j.Players[i].GetName(), j.Players[k].GetName())
			}
		}
		j_ := J(*j)
		fmt.Println(j_.String())
		session.Done()
		j.NextSession()
	}
}
