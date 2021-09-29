package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
	// *** next three lines no longer used, and can be removed

	//PLAYERWINS   = 1
	//COMPUTERWINS = 2
	//DRAW         = 3
)

type Round struct {
	// *** changed winner to Message
	//Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

// *** created three clices of strings, each with exactly three entries.

var winMessages = []string{
	"Good Job",
	"Nice Work",
	"You should buy a lottery Ticket!",
}

var loseMessages = []string{
	"Too bad, you lost",
	"Try again",
	"Today is not your lucky day",
}

var drawMessages = []string{
	"Great minds think alike",
	"Uh Oh! Try again.",
	"Nobody wins, buy you can try one more time",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)

	computerChoice := ""

	roundResult := ""
	//*** var no longer used
	//winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}
	// **** generate a random number from 0-2, which I use to pick a random message
	messageInt := rand.Intn(3)
	// *** declare a car to hold a message
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		//winner = DRAW
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		//winner = PLAYERWINS
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins!"
		//winner = COMPUTERWINS
		message = loseMessages[messageInt]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result
}
