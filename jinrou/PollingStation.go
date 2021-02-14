package jinrou

import (
	"math/rand"
)

type pair func() (*Player, int)

func (p *pair) modify(s *Player, i int) {
	*p = func() (*Player, int) { return s, i }
}

type pollingStation []pair

func newPollingStation(players []*Player) []pair {
	l := len(players)
	station := make([]pair, l)
	for i := 0; i < l; i++ {
		player := players[i]
		station[i] = func() (*Player, int) { return player, 0 }
	}
	return station
}

func (p *pollingStation) vote(name string) {
	for i := 0; i < len(*p); i++ {
		player, n := (*p)[i]()
		if player.name == name {
			(*p)[i].modify(player, n+1)
			break
		}
	}
}

func (p *pollingStation) voted() *Player {
	l := len(*p)
	for i := 0; i < l; i++ {
		j := rand.Intn(l)
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
	maxi := 0
	for i := 0; i < l; i++ {
		_, n1 := (*p)[maxi]()
		_, n2 := (*p)[i]()
		if n1 < n2 {
			maxi = i
		}
	}
	player, _ := (*p)[maxi]()
	return player
}
