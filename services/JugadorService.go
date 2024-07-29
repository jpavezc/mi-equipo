package services

import (
	"fmt"
	"math/rand"
	"time"
)

func ElegirRandom() {

	rand.Seed(time.Now().UnixNano())

	players := []string{"jugador1", "jugador2", "jugador3", "jugador4", "jugador5", "jugador6", "jugador7", "jugador8", "jugador9", "jugador10"}

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	team1 := players[:5]
	team2 := players[5:]

	fmt.Printf("Equipo 1: %v\n", team1)
	fmt.Printf("Equipo 2: %v\n", team2)

}
