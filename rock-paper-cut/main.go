package main

import (
	"myapp/game"
)

//constants
//capitalize constant variables.. and they can't be changed

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}
	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	//clears terminal window

	for {
		game.RoundChan <- 1
		<-game.RoundChan //wait's for something to happens
		if game.Round.RoundNumber > 3 {
			break
		}
		if !game.PlayerRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}
	game.PrintSummary()

}

// clearScreen clears the screen
